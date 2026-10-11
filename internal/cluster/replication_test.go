package cluster

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestDeriveReplicationRetainsZeroAndUnknownIndicators(t *testing.T) {
	info := map[string]string{
		"role": "slave", "master_replid": "lineage", "master_repl_offset": "120",
		"slave_repl_offset": "110", "slave_read_repl_offset": "115", "connected_slaves": "0",
		"master_link_status": "down", "master_sync_in_progress": "0", "master_last_io_seconds_ago": "0",
		"master_link_down_since_seconds": "-1", "repl_backlog_histlen": "invalid",
		"slave10": "ip=172.22.0.12,port=6379,state=online,offset=120,lag=0",
		"slave2":  "ip=172.22.0.4,port=6379,state=online,offset=110,lag=2",
		"slave0":  "ip=172.22.0.2,port=6379,state=wait_bgsave,offset=125,lag=-1",
		"slave-1": "offset=1", "slave_bad": "offset=1",
	}
	m := DeriveReplication(info)
	if m.Role != "slave" || m.ReplicationID != "lineage" || m.ReplicaOffset == nil || *m.ReplicaOffset != 110 || m.ReplicaReadOffset == nil || *m.ReplicaReadOffset != 115 {
		t.Fatalf("offset evidence lost: %+v", m)
	}
	if m.ConnectedReplicas == nil || *m.ConnectedReplicas != 0 || m.LastIOSeconds == nil || *m.LastIOSeconds != 0 || m.LinkUp == nil || *m.LinkUp || m.SyncInProgress == nil || *m.SyncInProgress {
		t.Fatalf("measured zero/false was confused with unknown: %+v", m)
	}
	if m.LinkDownSeconds != nil || m.BacklogBytes != nil {
		t.Fatal("invalid progress was fabricated as zero")
	}
	if len(m.Replicas) != 3 || m.Replicas[0].Name != "slave0" || m.Replicas[1].Name != "slave2" || m.Replicas[2].Name != "slave10" {
		t.Fatalf("replica progress not ordered by numeric index: %+v", m.Replicas)
	}
	if m.Replicas[0].OffsetGapBytes != nil || m.Replicas[0].LagSeconds != nil || m.Replicas[1].OffsetGapBytes == nil || *m.Replicas[1].OffsetGapBytes != 10 || m.Replicas[2].OffsetGapBytes == nil || *m.Replicas[2].OffsetGapBytes != 0 {
		t.Fatal("offset gaps were clamped or invalid lag accepted")
	}
	unknown := DeriveReplication(map[string]string{"master_link_status": "unknown", "master_sync_in_progress": "2", "master_repl_offset": "9223372036854775808"})
	if unknown.LinkUp != nil || unknown.SyncInProgress != nil || unknown.MasterOffset != nil || unknown.ConnectedReplicas != nil {
		t.Fatal("absent or invalid indicator accepted")
	}
}

func comparablePrimaryViews(t *testing.T) (View, View) {
	t.Helper()
	a, b := healthySnapshot(t).Views[0], healthySnapshot(t).Views[0]
	a.At = time.Date(2026, 10, 11, 0, 0, 0, 0, time.UTC)
	b.At = a.At.Add(time.Second)
	for _, v := range []*View{&a, &b} {
		v.Server = map[string]string{"run_id": "same-primary-process"}
		v.Replication["master_replid"] = "same-replication-lineage"
		v.Stats = map[string]string{"sync_full": "2", "sync_partial_ok": "3", "sync_partial_err": "1"}
	}
	return a, b
}

func TestResynchronizationUsesComparablePrimaryCounters(t *testing.T) {
	cases := []struct {
		name, full, partial, rejected, inference string
	}{
		{"unchanged", "2", "3", "1", "no_counter_increase_observed"},
		{"full", "3", "3", "1", "full_sync_counter_increased"},
		{"partial", "2", "4", "1", "partial_sync_counter_increased"},
		{"both", "3", "4", "2", "full_and_partial_sync_counters_increased"},
		{"rejected only", "2", "3", "2", "partial_sync_request_rejected"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a, b := comparablePrimaryViews(t)
			b.Stats = map[string]string{"sync_full": c.full, "sync_partial_ok": c.partial, "sync_partial_err": c.rejected}
			e := ResynchronizationDelta(a, b)
			if e.Inference != c.inference || len(e.Reasons) != 0 || e.FullSyncDelta == nil || e.PartialAcceptedDelta == nil || e.PartialRejectedDelta == nil || e.PrimaryRedisID != b.RedisID {
				t.Fatalf("wrong counter evidence: %+v", e)
			}
			if !strings.Contains(e.Methodology, "server-wide") || !strings.Contains(e.Methodology, "never exact byte-loss") {
				t.Fatal("counter limitations not preserved")
			}
			if c.name == "unchanged" {
				data, err := json.Marshal(e)
				if err != nil || !strings.Contains(string(data), `"full_sync_delta":0`) || !strings.Contains(string(data), `"partial_accepted_delta":0`) {
					t.Fatal("observed zero counter increase omitted")
				}
			}
		})
	}
	// Older JSON artifacts stored the server reply without a parsed map.
	a, b := comparablePrimaryViews(t)
	a.Server, b.Server = nil, nil
	a.RawServer, b.RawServer = "# Server\r\nrun_id:same-process\r\n", "# Server\r\nrun_id:same-process\r\n"
	if e := ResynchronizationDelta(a, b); e.Inference == "unknown" {
		t.Fatalf("raw INFO server identity was ignored: %+v", e)
	}
}

func TestResynchronizationRemainsUnknownAcrossInvalidCounterIntervals(t *testing.T) {
	cases := map[string]func(*View, *View){
		"primary unavailable":                            func(a, b *View) { b.Reachable = false },
		"command error":                                  func(a, b *View) { b.Errors = []string{"INFO persistence: timeout"} },
		"different Redis identity":                       func(a, b *View) { b.RedisID = fixtureID(1) },
		"different managed endpoint":                     func(a, b *View) { b.Node.Addr = "127.0.0.1:7999" },
		"process restart":                                func(a, b *View) { b.Server["run_id"] = "new-primary-process" },
		"missing process identity":                       func(a, b *View) { delete(a.Server, "run_id") },
		"replication lineage changed":                    func(a, b *View) { b.Replication["master_replid"] = "new-lineage" },
		"missing replication lineage":                    func(a, b *View) { delete(a.Replication, "master_replid") },
		"stale INFO role":                                func(a, b *View) { b.Replication["role"] = "slave" },
		"missing primary topology":                       func(a, b *View) { b.Nodes = nil },
		"failed primary topology":                        func(a, b *View) { b.Nodes[0].Flags = append(b.Nodes[0].Flags, "fail") },
		"promoted replica cannot compare local counters": func(a, b *View) { a.Nodes[0].Flags = []string{"slave"} },
		"missing time":                                   func(a, b *View) { a.At = time.Time{} },
		"reversed interval":                              func(a, b *View) { b.At = a.At.Add(-time.Second) },
		"missing counter":                                func(a, b *View) { delete(b.Stats, "sync_partial_err") },
		"invalid counter":                                func(a, b *View) { b.Stats["sync_partial_ok"] = "NaN" },
		"negative counter":                               func(a, b *View) { b.Stats["sync_full"] = "-1" },
		"overflow counter":                               func(a, b *View) { b.Stats["sync_full"] = "18446744073709551616" },
		"reset counter invalidates all deltas":           func(a, b *View) { b.Stats["sync_full"], b.Stats["sync_partial_ok"] = "0", "9" },
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			a, b := comparablePrimaryViews(t)
			mutate(&a, &b)
			e := ResynchronizationDelta(a, b)
			if e.Inference != "unknown" || len(e.Reasons) == 0 || e.FullSyncDelta != nil || e.PartialAcceptedDelta != nil || e.PartialRejectedDelta != nil {
				t.Fatalf("invalid interval fabricated synchronization: %+v", e)
			}
		})
	}
}
