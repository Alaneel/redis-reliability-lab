package experiment

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
	"github.com/redis/go-redis/v9"
)

type writerCommand struct {
	Node       string
	Connection int
	Args       []string
}

type writerWire struct {
	mu       sync.Mutex
	commands []writerCommand
	next     int
	sockets  []net.Conn
	wg       sync.WaitGroup
	handler  func(writerCommand) string
}

func readWriterCommand(reader *bufio.Reader) ([]string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	if len(line) < 3 || line[0] != '*' {
		return nil, errors.New("expected RESP array")
	}
	n, err := strconv.Atoi(strings.TrimSpace(line[1:]))
	if err != nil || n < 1 || n > 32 {
		return nil, errors.New("invalid RESP command length")
	}
	args := make([]string, n)
	for i := range args {
		line, err = reader.ReadString('\n')
		if err != nil || len(line) < 3 || line[0] != '$' {
			return nil, errors.New("expected RESP bulk string")
		}
		length, parseErr := strconv.Atoi(strings.TrimSpace(line[1:]))
		if parseErr != nil || length < 0 || length > 1<<20 {
			return nil, errors.New("invalid RESP bulk length")
		}
		payload := make([]byte, length+2)
		if _, err = io.ReadFull(reader, payload); err != nil {
			return nil, err
		}
		args[i] = string(payload[:length])
	}
	return args, nil
}

func (wire *writerWire) dial(node string) func(context.Context, string, string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		client, server := net.Pipe()
		wire.mu.Lock()
		wire.next++
		id := wire.next
		wire.sockets = append(wire.sockets, server)
		wire.wg.Add(1)
		wire.mu.Unlock()
		go func() {
			defer wire.wg.Done()
			defer server.Close()
			reader := bufio.NewReader(server)
			for {
				args, err := readWriterCommand(reader)
				if err != nil {
					return
				}
				if strings.EqualFold(args[0], "hello") {
					if _, err = io.WriteString(server, "-ERR unknown command HELLO\r\n"); err != nil {
						return
					}
					continue
				}
				command := writerCommand{Node: node, Connection: id, Args: args}
				wire.mu.Lock()
				wire.commands = append(wire.commands, command)
				wire.mu.Unlock()
				reply := wire.handler(command)
				if reply == "" {
					return
				}
				if _, err = io.WriteString(server, reply); err != nil {
					return
				}
			}
		}()
		return client, nil
	}
}

func (wire *writerWire) snapshot() []writerCommand {
	wire.mu.Lock()
	defer wire.mu.Unlock()
	return append([]writerCommand(nil), wire.commands...)
}

func (wire *writerWire) close() {
	wire.mu.Lock()
	for _, socket := range wire.sockets {
		_ = socket.Close()
	}
	wire.mu.Unlock()
	wire.wg.Wait()
}

func writerFixture(t *testing.T, policy string, handler func(writerCommand) string) (*ManagedWriter, *redis.ClusterClient, *writerWire) {
	t.Helper()
	env := testManagedEnvironment()
	if policy == "waitaof" {
		env.Persistence, env.RetainData = "aof-always", true
	}
	first := env.Nodes[0].InternalAddr
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{first}, Protocol: 2, MaxRetries: -1,
		ClusterSlots: func(context.Context) ([]redis.ClusterSlot, error) {
			return []redis.ClusterSlot{{Start: 0, End: 16383, Nodes: []redis.ClusterNode{{Addr: first}}}}, nil
		},
		Dialer: func(context.Context, string, string) (net.Conn, error) {
			return nil, errors.New("topology-only test client must not dial")
		},
	})
	writer, err := NewWriter(clusterClient, env, time.Second, policy, 1, 100*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	wire := &writerWire{handler: handler}
	for address, client := range writer.nodes {
		options := *client.Options()
		_ = client.Close()
		options.Dialer = wire.dial(address)
		// Options have already normalized -1 to zero. NewClient normalizes zero
		// to default retries, so set -1 again for the actual test socket client.
		options.MaxRetries = -1
		writer.nodes[address] = redis.NewClient(&options)
	}
	t.Cleanup(func() { _ = writer.Close(); _ = clusterClient.Close(); wire.close() })
	return writer, clusterClient, wire
}

func writerSuccess(command writerCommand) string {
	if strings.EqualFold(command.Args[0], "wait") {
		return ":1\r\n"
	}
	return "+OK\r\n"
}

func TestManagedWriterSamePhysicalConnectionAndComparableSETPath(t *testing.T) {
	for _, policy := range []string{"set", "wait"} {
		t.Run(policy, func(t *testing.T) {
			writer, _, wire := writerFixture(t, policy, writerSuccess)
			result := writer.Write(context.Background(), "immutable-key", "payload")
			if !result.SetAcknowledged || !result.PolicySatisfied || result.SetError != nil || result.WaitError != nil || result.TotalLatency <= 0 || result.SetLatency <= 0 {
				t.Fatalf("write result: %#v", result)
			}
			commands := wire.snapshot()
			if policy == "set" {
				if len(commands) != 1 || commands[0].Args[0] != "set" || result.WaitAttempted {
					t.Fatalf("normal policy issued WAIT: %#v", commands)
				}
			} else {
				if len(commands) != 2 || commands[0].Args[0] != "set" || commands[1].Args[0] != "wait" || commands[0].Connection != commands[1].Connection || commands[0].Node != commands[1].Node || !result.WaitAttempted || result.WaitAcknowledgements != 1 || result.WaitLatency <= 0 {
					t.Fatalf("SET/WAIT physical affinity violated: %#v %#v", commands, result)
				}
				if strings.Join(commands[1].Args, " ") != "wait 1 100" {
					t.Fatalf("wrong bounded WAIT: %#v", commands[1])
				}
			}
		})
	}
}

func TestManagedWriterRetainsSETAcknowledgementAfterShortOrUnknownWAIT(t *testing.T) {
	for _, reply := range []string{":0\r\n", "-ERR WAIT unavailable\r\n", ""} {
		t.Run(fmt.Sprintf("reply-%q", reply), func(t *testing.T) {
			writer, _, wire := writerFixture(t, "wait", func(command writerCommand) string {
				if command.Args[0] == "wait" {
					return reply
				}
				return "+OK\r\n"
			})
			result := writer.Write(context.Background(), "key", "value")
			if !result.SetAcknowledged || result.SetError != nil || !result.WaitAttempted || result.PolicySatisfied {
				t.Fatalf("SET acknowledgement disappeared: %#v", result)
			}
			if (reply == ":0\r\n") != (result.WaitError == nil) {
				t.Fatalf("short versus unknown WAIT misclassified: %#v", result)
			}
			if len(wire.snapshot()) != 2 {
				t.Fatalf("SET or WAIT replayed: %#v", wire.snapshot())
			}
		})
	}
}

func TestManagedWriterDeadlineDuringWAITRetainsEarlierSETAckTimestamp(t *testing.T) {
	waitObserved := make(chan time.Time, 1)
	writer, _, wire := writerFixture(t, "wait", func(command writerCommand) string {
		if command.Args[0] == "wait" {
			waitObserved <- time.Now()
			time.Sleep(100 * time.Millisecond)
			return ":1\r\n"
		}
		return "+OK\r\n"
	})
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	result := writer.Write(ctx, "key", "value")
	if !result.SetAcknowledged || result.SetError != nil || result.SetAcknowledgedAt.IsZero() || !result.WaitAttempted || result.WaitError == nil || result.PolicySatisfied || result.TotalLatency > 500*time.Millisecond {
		t.Fatalf("deadline hid SET or escaped operation bound: %#v", result)
	}
	// The command is observed before its delayed reply. That observation gives
	// a wire-level boundary for the acknowledgement timestamp, independent of
	// the final logical operation completion.
	waitBegan := <-waitObserved
	if waitBegan.IsZero() || result.SetAcknowledgedAt.After(waitBegan) || result.SetLatency >= result.TotalLatency {
		t.Fatalf("SET acknowledgement recorded after WAIT: %#v, wait=%v", result, waitBegan)
	}
	if len(wire.snapshot()) != 2 {
		t.Fatal("deadline replayed SET or WAIT")
	}
}

func TestManagedWriterDeadlineDuringSETNeverIssuesWAIT(t *testing.T) {
	writer, _, wire := writerFixture(t, "wait", func(writerCommand) string {
		time.Sleep(100 * time.Millisecond)
		return "+OK\r\n"
	})
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	result := writer.Write(ctx, "key", "value")
	if result.SetAcknowledged || result.SetError == nil || result.WaitAttempted || !result.SetAcknowledgedAt.IsZero() || len(wire.snapshot()) != 1 {
		t.Fatalf("interrupted SET gained an ack or WAIT: %#v %#v", result, wire.snapshot())
	}
}

func TestManagedWriterConcurrentOperationsKeepEachWAITAfterItsOwnSET(t *testing.T) {
	writer, _, wire := writerFixture(t, "wait", writerSuccess)
	var wg sync.WaitGroup
	results := make(chan workload.WriteResult, 32)
	for i := range 32 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results <- writer.Write(context.Background(), fmt.Sprintf("key-%d", i), "immutable")
		}(i)
	}
	wg.Wait()
	close(results)
	for result := range results {
		if !result.SetAcknowledged || !result.PolicySatisfied || result.SetError != nil || result.WaitError != nil {
			t.Fatalf("concurrent write failed: %#v", result)
		}
	}
	perConnection := map[int][]writerCommand{}
	for _, command := range wire.snapshot() {
		perConnection[command.Connection] = append(perConnection[command.Connection], command)
	}
	for id, commands := range perConnection {
		if len(commands)%2 != 0 {
			t.Fatalf("connection %d carried unmatched command: %#v", id, commands)
		}
		for i := 0; i < len(commands); i += 2 {
			if commands[i].Args[0] != "set" || commands[i+1].Args[0] != "wait" {
				t.Fatalf("another write interleaved into sticky connection %d: %#v", id, commands)
			}
		}
	}
}

func TestManagedWriterNeverReplaysInterruptedSETOrIssuesWAIT(t *testing.T) {
	writer, _, wire := writerFixture(t, "wait", func(writerCommand) string { return "" })
	result := writer.Write(context.Background(), "key", "value")
	if result.SetAcknowledged || result.SetError == nil || result.WaitAttempted || len(wire.snapshot()) != 1 {
		t.Fatalf("interrupted SET replayed or WAIT issued: %#v %#v", result, wire.snapshot())
	}
}

func TestManagedWriterUnexpectedSETStatusIsNotAcknowledged(t *testing.T) {
	writer, _, wire := writerFixture(t, "wait", func(writerCommand) string { return "+UNEXPECTED\r\n" })
	result := writer.Write(context.Background(), "key", "value")
	if result.SetAcknowledged || result.SetError == nil || result.WaitAttempted || len(wire.snapshot()) != 1 {
		t.Fatalf("unexpected SET status admitted to ledger: %#v", result)
	}
}

func TestManagedWriterMOVEDAndASKOnlyUseAllowlistedTargets(t *testing.T) {
	for _, kind := range []string{"MOVED", "ASK"} {
		t.Run(kind, func(t *testing.T) {
			env := testManagedEnvironment()
			key := "redirect-key"
			writer, _, wire := writerFixture(t, "wait", func(command writerCommand) string {
				if command.Node == env.Nodes[0].InternalAddr && command.Args[0] == "set" {
					return fmt.Sprintf("-%s %d %s\r\n", kind, workload.HashSlot(key), env.Nodes[1].InternalAddr)
				}
				return writerSuccess(command)
			})
			result := writer.Write(context.Background(), key, "value")
			if !result.SetAcknowledged || !result.PolicySatisfied || result.Redirects != 1 {
				t.Fatalf("safe redirect failed: %#v", result)
			}
			commands := wire.snapshot()
			last := commands[len(commands)-2:]
			if last[0].Node != env.Nodes[1].InternalAddr || last[0].Connection != last[1].Connection || last[0].Args[0] != "set" || last[1].Args[0] != "wait" {
				t.Fatalf("redirect lost affinity: %#v", commands)
			}
			if kind == "ASK" && (len(commands) != 4 || commands[1].Args[0] != "ASKING" || commands[1].Connection != last[0].Connection) {
				t.Fatalf("ASKING not on SET connection: %#v", commands)
			}
		})
	}
	for _, target := range []string{"203.0.113.1:6379", "127.0.0.1:9999", "example.com:6379"} {
		writer, _, wire := writerFixture(t, "wait", func(writerCommand) string { return fmt.Sprintf("-MOVED %d %s\r\n", workload.HashSlot("key"), target) })
		result := writer.Write(context.Background(), "key", "value")
		if result.SetAcknowledged || result.SetError == nil || result.Redirects != 0 || len(wire.snapshot()) != 1 {
			t.Fatalf("unsafe redirect accepted: %#v %#v", result, wire.snapshot())
		}
	}
}

func TestManagedWriterBoundsRedirectLoopsAndRejectsWrongSlot(t *testing.T) {
	for _, wrongSlot := range []bool{false, true} {
		env := testManagedEnvironment()
		slot := workload.HashSlot("key")
		if wrongSlot {
			slot = (slot + 1) % 16384
		}
		writer, _, wire := writerFixture(t, "wait", func(writerCommand) string { return fmt.Sprintf("-MOVED %d %s\r\n", slot, env.Nodes[0].InternalAddr) })
		result := writer.Write(context.Background(), "key", "value")
		want := 9
		if wrongSlot {
			want = 1
		}
		if result.SetAcknowledged || result.SetError == nil || len(wire.snapshot()) != want || result.WaitAttempted {
			t.Fatalf("redirect not bounded: %#v %#v", result, wire.snapshot())
		}
	}
}

func TestManagedWriterASKINGFailureCannotIssueSETOrWAITAtTarget(t *testing.T) {
	env := testManagedEnvironment()
	writer, _, wire := writerFixture(t, "wait", func(command writerCommand) string {
		if command.Node == env.Nodes[0].InternalAddr {
			return fmt.Sprintf("-ASK %d %s\r\n", workload.HashSlot("key"), env.Nodes[1].InternalAddr)
		}
		return "-ERR ASKING denied\r\n"
	})
	result := writer.Write(context.Background(), "key", "value")
	commands := wire.snapshot()
	if result.SetAcknowledged || result.SetError == nil || result.WaitAttempted || len(commands) != 2 || commands[1].Args[0] != "ASKING" {
		t.Fatalf("failed ASKING continued writing: %#v %#v", result, commands)
	}
}

func TestManagedWriterRefusesUnmanagedDiscoveredOwnerBeforeDial(t *testing.T) {
	env := testManagedEnvironment()
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{env.Nodes[0].Addr},
		ClusterSlots: func(context.Context) ([]redis.ClusterSlot, error) {
			return []redis.ClusterSlot{{Start: 0, End: 16383, Nodes: []redis.ClusterNode{{Addr: "203.0.113.1:6379"}}}}, nil
		},
	})
	defer clusterClient.Close()
	writer, err := NewWriter(clusterClient, env, time.Second, "wait", 1, 100*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	defer writer.Close()
	result := writer.Write(context.Background(), "key", "value")
	if result.SetAcknowledged || result.WaitAttempted || result.SetError == nil || !strings.Contains(result.SetError.Error(), "unmanaged slot owner") {
		t.Fatalf("unmanaged topology owner was not rejected: %#v", result)
	}
}

func TestManagedRedirectCannotMistakeTransportTextForServerRejection(t *testing.T) {
	env := testManagedEnvironment()
	err := fmt.Errorf("MOVED %d %s", workload.HashSlot("key"), env.Nodes[0].InternalAddr)
	if _, _, ok := managedRedirect(err, "key", map[string]string{env.Nodes[0].InternalAddr: env.Nodes[0].InternalAddr}); ok {
		t.Fatal("transport error text permitted SET replay")
	}
}

func TestManagedWriterCloseDoesNotCloseClusterAndRejectsLaterWrites(t *testing.T) {
	writer, clusterClient, _ := writerFixture(t, "wait", writerSuccess)
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := clusterClient.MasterForKey(context.Background(), "still-open"); err != nil {
		t.Fatalf("writer closed borrowed cluster: %v", err)
	}
	result := writer.Write(context.Background(), "key", "value")
	if result.SetError == nil || result.SetAcknowledged {
		t.Fatalf("closed writer used: %#v", result)
	}
}

func TestManagedWriterValidatesWAITMillisecondsAndOwnedInventory(t *testing.T) {
	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{Addrs: []string{testManagedEnvironment().Nodes[0].Addr}})
	defer clusterClient.Close()
	for _, invalid := range []struct {
		policy          string
		replicas        int
		wait, timeLimit time.Duration
	}{
		{"unknown", 1, time.Millisecond, time.Second}, {"wait", 0, time.Millisecond, time.Second}, {"wait", 2, time.Millisecond, time.Second},
		{"wait", 1, 0, time.Second}, {"wait", 1, time.Nanosecond, time.Second}, {"wait", 1, 1500 * time.Microsecond, time.Second},
		{"wait", 1, time.Second, time.Second}, {"wait", 1, time.Second, 500 * time.Millisecond},
	} {
		if writer, err := NewWriter(clusterClient, testManagedEnvironment(), invalid.timeLimit, invalid.policy, invalid.replicas, invalid.wait); err == nil || writer != nil {
			t.Fatalf("unsafe policy accepted: %#v", invalid)
		}
	}
	for _, mutate := range []func(*environment.Environment){
		func(env *environment.Environment) { env.Nodes[0].Addr = "203.0.113.1:6379" },
		func(env *environment.Environment) { env.Nodes[0].InternalAddr = "203.0.113.1:6379" },
		func(env *environment.Environment) { env.Nodes[1].InternalAddr = env.Nodes[0].InternalAddr },
		func(env *environment.Environment) { env.Nodes = env.Nodes[:5] },
	} {
		env := testManagedEnvironment()
		mutate(env)
		if writer, err := NewWriter(clusterClient, env, time.Second, "set", 1, 100*time.Millisecond); err == nil || writer != nil {
			t.Fatal("unsafe writer inventory accepted")
		}
	}
}
