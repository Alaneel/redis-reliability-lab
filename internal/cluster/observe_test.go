package cluster

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

	"github.com/redis/go-redis/v9"
)

func readObservationCommand(reader *bufio.Reader) ([]string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	if len(line) < 3 || line[0] != '*' {
		return nil, errors.New("expected RESP command array")
	}
	count, err := strconv.Atoi(strings.TrimSpace(line[1:]))
	if err != nil || count < 1 || count > 16 {
		return nil, errors.New("invalid RESP command count")
	}
	args := make([]string, count)
	for i := range args {
		line, err = reader.ReadString('\n')
		if err != nil || len(line) < 3 || line[0] != '$' {
			return nil, errors.New("expected RESP argument")
		}
		length, err := strconv.Atoi(strings.TrimSpace(line[1:]))
		if err != nil || length < 0 || length > 1<<20 {
			return nil, errors.New("invalid RESP argument length")
		}
		value := make([]byte, length+2)
		if _, err = io.ReadFull(reader, value); err != nil {
			return nil, err
		}
		args[i] = string(value[:length])
	}
	return args, nil
}

func TestObserverCapturesPersistenceStatsAndIndependentCommandErrors(t *testing.T) {
	for _, failed := range []string{"", "INFO PERSISTENCE", "INFO STATS"} {
		t.Run(failed, func(t *testing.T) {
			rawNodes := fmt.Sprintf("%s 172.22.0.2:6379@16379 myself,master - 0 1 1 connected 0-16383\n", fixtureID(0))
			replies := map[string]string{
				"CLUSTER INFO":     "cluster_state:ok\r\ncluster_slots_assigned:16384\r\ncluster_slots_ok:16384\r\ncluster_slots_fail:0\r\ncluster_slots_pfail:0\r\n",
				"CLUSTER NODES":    rawNodes,
				"INFO REPLICATION": "# Replication\r\nrole:master\r\nconnected_slaves:0\r\nmaster_repl_offset:0\r\nmaster_replid:lineage\r\n",
				"INFO SERVER":      "# Server\r\nredis_version:8.2.1\r\nrun_id:server-process\r\n",
				"INFO PERSISTENCE": "# Persistence\r\nloading:0\r\naof_enabled:1\r\naof_last_write_status:ok\r\naof_pending_bio_fsync:0\r\n",
				"INFO STATS":       "# Stats\r\nsync_full:2\r\nsync_partial_ok:3\r\nsync_partial_err:0\r\n",
			}
			observer := New([]Node{{ID: "node-1", Addr: "127.0.0.1:7000", InternalAddr: "172.22.0.2:6379"}}, time.Second)
			options := *observer.clients[0].Options()
			_ = observer.clients[0].Close()
			options.DisableIdentity, options.MaxRetries = true, -1
			var sockets []net.Conn
			var socketMu sync.Mutex
			var wg sync.WaitGroup
			options.Dialer = func(ctx context.Context, network, addr string) (net.Conn, error) {
				client, server := net.Pipe()
				socketMu.Lock()
				sockets = append(sockets, server)
				socketMu.Unlock()
				wg.Add(1)
				go func() {
					defer wg.Done()
					defer server.Close()
					reader := bufio.NewReader(server)
					for {
						args, err := readObservationCommand(reader)
						if err != nil {
							return
						}
						command := strings.ToUpper(strings.Join(args, " "))
						var reply string
						if strings.HasPrefix(command, "HELLO ") {
							reply = "-ERR unknown command HELLO\r\n"
						} else if command == failed {
							reply = "-ERR deliberately rejected observation\r\n"
						} else if value, exists := replies[command]; exists {
							reply = fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)
						} else {
							reply = "-ERR unexpected test command\r\n"
						}
						if _, err = io.WriteString(server, reply); err != nil {
							return
						}
					}
				}()
				return client, nil
			}
			observer.clients[0] = redis.NewClient(&options)
			t.Cleanup(func() {
				_ = observer.Close()
				socketMu.Lock()
				for _, socket := range sockets {
					_ = socket.Close()
				}
				socketMu.Unlock()
				wg.Wait()
			})
			s := observer.Observe(context.Background())
			if len(s.Views) != 1 || s.Completed.Before(s.At) {
				t.Fatal("observation evidence interval absent")
			}
			v := s.Views[0]
			if !v.Reachable || v.RedisID != fixtureID(0) || v.RedisVersion != "8.2.1" || v.Server["run_id"] != "server-process" || v.RawNodes != rawNodes || !v.Coverage.Full() || v.ReplicationMetrics == nil || v.ReplicationMetrics.MasterOffset == nil || *v.ReplicationMetrics.MasterOffset != 0 {
				t.Fatalf("core or derived evidence lost: %+v", v)
			}
			if failed == "" && len(v.Errors) != 0 || failed != "" && (len(v.Errors) != 1 || !strings.Contains(strings.ToUpper(v.Errors[0]), failed)) {
				t.Fatalf("command error lost or invented: %v", v.Errors)
			}
			if failed != "INFO PERSISTENCE" && (v.RawPersistence != replies["INFO PERSISTENCE"] || v.Persistence["aof_last_write_status"] != "ok") {
				t.Fatal("persistence evidence lost")
			}
			if failed != "INFO STATS" && (v.RawStats != replies["INFO STATS"] || v.Stats["sync_partial_ok"] != "3") {
				t.Fatal("stats reply after previous command error lost")
			}
			if failed == "INFO PERSISTENCE" && (v.RawPersistence != "" || len(v.Persistence) != 0) || failed == "INFO STATS" && (v.RawStats != "" || len(v.Stats) != 0) {
				t.Fatal("failed reply fabricated as observed values")
			}
		})
	}
}
