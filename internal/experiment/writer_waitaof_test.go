package experiment

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/workload"
	"github.com/redis/go-redis/v9"
)

func aofWriterSuccess(command writerCommand) string {
	if strings.EqualFold(command.Args[0], "waitaof") {
		return "*2\r\n:1\r\n:1\r\n"
	}
	return "+OK\r\n"
}

func TestManagedWriterWAITAOFUsesSuccessfulSETPhysicalConnection(t *testing.T) {
	for _, redirect := range []string{"", "MOVED", "ASK"} {
		t.Run(redirect, func(t *testing.T) {
			env := testManagedEnvironment()
			key := "aof-affinity"
			writer, _, wire := writerFixture(t, "waitaof", func(command writerCommand) string {
				if redirect != "" && command.Node == env.Nodes[0].InternalAddr && command.Args[0] == "set" {
					return fmt.Sprintf("-%s %d %s\r\n", redirect, workload.HashSlot(key), env.Nodes[1].InternalAddr)
				}
				return aofWriterSuccess(command)
			})
			result := writer.Write(context.Background(), key, "immutable")
			if !result.SetAcknowledged || !result.PolicySatisfied || !result.WaitAOFAttempted || result.WaitAOFError != nil || result.WaitAttempted || result.WaitAOFRequestedLocal != 1 || result.WaitAOFRequestedReplicas != 1 || result.WaitAOFLocalAcknowledgements != 1 || result.WaitAOFReplicaAcknowledgements != 1 || result.WaitAOFLatency <= 0 {
				t.Fatalf("wrong independent acknowledgement result: %#v", result)
			}
			commands := wire.snapshot()
			last := commands[len(commands)-2:]
			if last[0].Args[0] != "set" || !strings.EqualFold(strings.Join(last[1].Args, " "), "waitaof 1 1 100") || last[0].Connection != last[1].Connection || last[0].Node != last[1].Node {
				t.Fatalf("lost physical connection or bounded parameters: %#v", commands)
			}
			if redirect == "ASK" && (commands[1].Args[0] != "ASKING" || commands[1].Connection != last[0].Connection) {
				t.Fatal("ASKING did not precede SET and WAITAOF on one connection")
			}
		})
	}
}

func TestManagedWriterWAITAOFShortErrorsAndMalformedReplyKeepSETAck(t *testing.T) {
	for _, tc := range []struct {
		name, reply     string
		local, replicas int64
		wantError       bool
	}{
		{"local_short", "*2\r\n:0\r\n:1\r\n", 0, 1, false},
		{"replica_short", "*2\r\n:1\r\n:0\r\n", 1, 0, false},
		{"both_short", "*2\r\n:0\r\n:0\r\n", 0, 0, false},
		{"unknown_command", "-ERR unknown command WAITAOF\r\n", 0, 0, true},
		{"aof_disabled", "-ERR WAITAOF requires appendonly to be enabled\r\n", 0, 0, true},
		{"disconnect", "", 0, 0, true},
		{"wrong_length", "*1\r\n:1\r\n", 0, 0, true},
		{"negative_local", "*2\r\n:-1\r\n:1\r\n", 0, 0, true},
		{"invalid_local", "*2\r\n:2\r\n:1\r\n", 0, 0, true},
		{"negative_replica", "*2\r\n:1\r\n:-1\r\n", 0, 0, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			writer, _, wire := writerFixture(t, "waitaof", func(command writerCommand) string {
				if strings.EqualFold(command.Args[0], "waitaof") {
					return tc.reply
				}
				return "+OK\r\n"
			})
			result := writer.Write(context.Background(), "key", "value")
			if !result.SetAcknowledged || result.SetError != nil || result.SetAcknowledgedAt.IsZero() || !result.WaitAOFAttempted || result.PolicySatisfied || result.WaitAttempted || (result.WaitAOFError != nil) != tc.wantError || result.WaitAOFLocalAcknowledgements != tc.local || result.WaitAOFReplicaAcknowledgements != tc.replicas || len(wire.snapshot()) != 2 {
				t.Fatalf("SET ack lost, reply conflated or command replayed: %#v %#v", result, wire.snapshot())
			}
		})
	}
}

func TestManagedWriterWAITAOFDeadlineAndUnacknowledgedSET(t *testing.T) {
	for _, delayed := range []string{"set", "waitaof"} {
		t.Run(delayed, func(t *testing.T) {
			writer, _, wire := writerFixture(t, "waitaof", func(command writerCommand) string {
				if strings.EqualFold(command.Args[0], delayed) {
					time.Sleep(100 * time.Millisecond)
				}
				return aofWriterSuccess(command)
			})
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
			defer cancel()
			result := writer.Write(ctx, "key", "value")
			if delayed == "set" {
				if result.SetAcknowledged || result.SetError == nil || result.WaitAOFAttempted || len(wire.snapshot()) != 1 {
					t.Fatalf("unacknowledged SET became policy ack: %#v", result)
				}
			} else if !result.SetAcknowledged || result.WaitAOFError == nil || !result.WaitAOFAttempted || result.PolicySatisfied || len(wire.snapshot()) != 2 || result.SetLatency >= result.TotalLatency {
				t.Fatalf("deadline hid earlier SET or replayed: %#v", result)
			}
		})
	}
}

func TestManagedWriterConcurrentWAITAOFDoesNotInterleaveWrites(t *testing.T) {
	writer, _, wire := writerFixture(t, "waitaof", aofWriterSuccess)
	var wg sync.WaitGroup
	for i := range 32 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			result := writer.Write(context.Background(), fmt.Sprintf("key-%d", i), "value")
			if !result.PolicySatisfied {
				t.Errorf("policy failed: %#v", result)
			}
		}(i)
	}
	wg.Wait()
	byConn := map[int][]writerCommand{}
	for _, command := range wire.snapshot() {
		byConn[command.Connection] = append(byConn[command.Connection], command)
	}
	for _, commands := range byConn {
		if len(commands)%2 != 0 {
			t.Fatal("unpaired command")
		}
		for i := 0; i < len(commands); i += 2 {
			if commands[i].Args[0] != "set" || !strings.EqualFold(commands[i+1].Args[0], "waitaof") {
				t.Fatalf("interleaved commands: %#v", commands)
			}
		}
	}
}

func TestManagedWriterWAITAOFRequiresManagedPersistence(t *testing.T) {
	env := testManagedEnvironment()
	client := redis.NewClusterClient(&redis.ClusterOptions{Addrs: []string{env.Nodes[0].Addr}})
	defer client.Close()
	for _, persistence := range []string{"memory", "", "aof-always"} {
		env.Persistence, env.RetainData = persistence, false
		if writer, err := NewWriter(client, env, time.Second, "waitaof", 1, 100*time.Millisecond); err == nil || writer != nil {
			t.Fatal("accepted unavailable AOF persistence")
		}
	}
	env.Persistence, env.RetainData = "aof-always", true
	for _, invalid := range []struct {
		replicas int
		wait     time.Duration
	}{{0, time.Millisecond}, {2, time.Millisecond}, {1, 0}, {1, 1500 * time.Microsecond}, {1, 2 * time.Second}} {
		if writer, err := NewWriter(client, env, 3*time.Second, "waitaof", invalid.replicas, invalid.wait); err == nil || writer != nil {
			t.Fatal("accepted invalid WAITAOF bounds")
		}
	}
}
