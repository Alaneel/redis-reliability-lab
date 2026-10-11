package report

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/metrics"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

func lifecycleReport() Report {
	r := sampleReport()
	r.SchemaVersion = LifecycleSchemaVersion
	good, bad := true, false
	actionDuration := 14.125
	requested, completed := r.StartedAt.Add(time.Second), r.StartedAt.Add(time.Second+14*time.Millisecond)
	target := cluster.Node{ID: "node-4", Addr: "127.0.0.1:7003", InternalAddr: "172.22.0.5:6379"}
	view := cluster.View{Node: target, RedisID: strings.Repeat("d", 40), Reachable: true, At: completed,
		RawPersistence: "# Persistence\r\nloading:0\r\naof_enabled:0\r\n", Persistence: map[string]string{"loading": "0", "aof_enabled": "0"},
		RawStats: "# Stats\r\nsync_full:0\r\nsync_partial_ok:1\r\nsync_partial_err:0\r\n", Stats: map[string]string{"sync_full": "0", "sync_partial_ok": "1", "sync_partial_err": "0"},
		RawReplication: "role:slave\r\nmaster_link_status:down\r\nmaster_sync_in_progress:1\r\n", Replication: map[string]string{"role": "slave", "master_link_status": "down", "master_sync_in_progress": "1"}}
	derived := cluster.DeriveReplication(view.Replication)
	view.ReplicationMetrics = &derived
	snapshot := &cluster.Snapshot{At: completed, Completed: completed, Views: []cluster.View{view}}
	dimensions := &RecoveryDimensions{TargetRole: "replica", TargetProcessLive: &good, MembershipHealthy: &bad, SlotCoverage: &good, LeadersAvailable: &good, ReadAvailable: &good, WriteAvailable: &good, RedundancyRestored: &bad}
	r.Dimensions = dimensions
	r.Final = snapshot
	r.LifecycleFault = &LifecycleFault{TargetKind: "replica", Target: target, TargetRedisID: view.RedisID, RepairKind: "restart",
		RepairRequestedAt: &requested, RepairCompletedAt: &completed, RepairActionDurationMS: &actionDuration, RepairCompleted: true, RepairTimedOut: true,
		RestartReceipt: &environment.RestartReceipt{RunID: r.RunID, NodeID: target.ID, RequestedAt: requested, CompletedAt: completed, OriginalRedisID: view.RedisID,
			Before: environment.ProcessEvidence{ContainerRunning: false}, After: environment.ProcessEvidence{ContainerRunning: true, RedisPONG: true, RedisClusterID: view.RedisID, RawPersistence: view.RawPersistence}},
		FaultSnapshot: snapshot, RepairSnapshot: snapshot, Dimensions: dimensions,
		TargetReachable:   &Milestone{At: completed, SinceFaultMS: 0},
		Resynchronization: &cluster.ResynchronizationEvidence{Inference: "unknown", Reasons: []string{"primary process identity changed"}, Methodology: cluster.ResynchronizationMethodology},
		Errors:            []string{"synchronization deadline elapsed"}}
	r.Samples = []Sample{{Snapshot: *snapshot, Dimensions: dimensions, Probe: Probe{At: completed, Successful: true, ReadSuccessful: &good, WriteSuccessful: &good}}}
	return r
}

func TestLifecycleJSONPreservesIndependentUnknownAndFailedConditions(t *testing.T) {
	r := lifecycleReport()
	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	var doc map[string]json.RawMessage
	if err = json.Unmarshal(data, &doc); err != nil {
		t.Fatal(err)
	}
	var dimensions map[string]json.RawMessage
	if err = json.Unmarshal(doc["recovery_dimensions"], &dimensions); err != nil {
		t.Fatal(err)
	}
	if string(dimensions["redundancy_restored"]) != "false" || string(dimensions["target_process_live"]) != "true" {
		t.Fatal("actual false and true dimensions were omitted")
	}
	for _, unobserved := range []string{"persistence_recovered", "data_verified", "write_policy_available"} {
		if _, present := dimensions[unobserved]; present {
			t.Fatalf("unobserved %s fabricated", unobserved)
		}
	}
	var decoded Report
	if err = json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	l := decoded.LifecycleFault
	if decoded.SchemaVersion != 3 || l == nil || !l.RepairCompleted || !l.RepairTimedOut || l.StableRedundancy != nil || l.SyncComplete != nil || l.TargetReachable == nil || l.TargetReachable.SinceFaultMS != 0 {
		t.Fatalf("repair action completion was confused with observed recovery: %+v", l)
	}
	if l.RestartReceipt == nil || l.RestartReceipt.Before.ContainerRunning || l.RestartReceipt.After.RedisClusterID != l.TargetRedisID || l.RestartReceipt.After.RawPersistence != r.Final.Views[0].RawPersistence {
		t.Fatal("physical restart evidence lost")
	}
	v := decoded.Final.Views[0]
	if v.RawPersistence != r.Final.Views[0].RawPersistence || v.Persistence["aof_enabled"] != "0" || v.RawStats != r.Final.Views[0].RawStats || v.Stats["sync_partial_ok"] != "1" || v.ReplicationMetrics == nil || v.ReplicationMetrics.LinkUp == nil || *v.ReplicationMetrics.LinkUp {
		t.Fatal("raw or derived synchronization evidence lost")
	}
	if decoded.Dimensions == nil || decoded.Dimensions.PersistenceRecovered != nil || decoded.Dimensions.RedundancyRestored == nil || *decoded.Dimensions.RedundancyRestored || decoded.Samples[0].Probe.ReadSuccessful == nil || !*decoded.Samples[0].Probe.ReadSuccessful || decoded.Samples[0].Dimensions == nil {
		t.Fatal("sample/final dimensions changed")
	}
}

func TestLifecycleMarkdownSeparatesRepairFromSynchronizedRedundancy(t *testing.T) {
	r := lifecycleReport()
	markdown := r.Markdown()
	for _, expected := range []string{
		"Target: replica `node-4`", "Repair action completed: true", "deadline expired: true", "Repair action duration: 14.125 ms",
		"| Cluster membership healthy | false |", "| Client reads available | true |", "| Client writes available | true |",
		"| Replication redundancy restored | false |", "| Persistence recovery verified | unobserved |", "| Acknowledged data verified | unobserved |",
		"| Target reachable after repair | 0.000 |", "| Synchronization complete | unobserved |", "| Replication redundancy sustained | unobserved |",
		"Resynchronization evidence: `unknown`", "primary process identity changed", "synchronization deadline elapsed", "not zero lag or durability guarantees",
	} {
		if !strings.Contains(markdown, expected) {
			t.Errorf("missing lifecycle evidence %q:\n%s", expected, markdown)
		}
	}
	if strings.Contains(markdown, "The primary stays stopped") || strings.Count(markdown, "## Independent recovery dimensions") != 1 {
		t.Fatal("replica lifecycle was described as a primary kill or dimensions duplicated")
	}
	// Independent final dimensions also apply to persistence experiments that
	// have no replica lifecycle section. Lifecycle-only saved reports retain a
	// readable fallback when the additive top-level pointer is absent.
	r.LifecycleFault = nil
	if !strings.Contains(r.Markdown(), "| Persistence recovery verified | unobserved |") {
		t.Fatal("top-level dimensions depend on a lifecycle fault")
	}
	r = lifecycleReport()
	r.Dimensions = nil
	if !strings.Contains(r.Markdown(), "| Replication redundancy restored | false |") {
		t.Fatal("lifecycle dimensions fallback omitted")
	}
}

func TestLifecycleReportPreservesRepairPhaseMetrics(t *testing.T) {
	r := lifecycleReport()
	r.WorkloadStats.WritePolicy = "wait"
	r.WorkloadStats.PolicyByPhase = map[string]workload.PolicyCounts{"repair": {SetAcknowledged: 3, WaitAttempted: 3, WaitSatisfied: 1, WaitShort: 2}}
	r.WorkloadStats.WriteStages.Phases = map[string]metrics.PhaseSummary{"repair": {ByOperation: map[string]metrics.OperationSummary{"set": {Latency: metrics.LatencySummary{Count: 3, P50MS: 2}}}}}
	r.Metrics.Phases["repair"] = metrics.PhaseSummary{Attempts: 4, Successes: 2, Errors: 2, SuccessRate: .5, Throughput: 2, DurationSeconds: 2,
		ByOperation: map[string]metrics.OperationSummary{"set_wait": {Successes: 1, Latency: metrics.LatencySummary{Count: 3, P50MS: 100}}}}
	markdown := r.Markdown()
	for _, expected := range []string{"| repair | 4 | 50.00% | 2 | 2.00 |", "| repair | 3 | 0 | 3 | 1 | 2 | 0 | 0 |", "| repair | Logical write completion | 3 | 100.000", "| repair | set | 3 | 2.000", "| repair | 1 | 0.500 |"} {
		if !strings.Contains(markdown, expected) {
			t.Errorf("repair metrics missing %q", expected)
		}
	}
}

func TestLegacyReportsDoNotInventLifecycleEvidence(t *testing.T) {
	for _, version := range []int{1, 2} {
		r := sampleReport()
		r.SchemaVersion = version
		data, err := json.Marshal(r)
		if err != nil {
			t.Fatal(err)
		}
		for _, field := range []string{`"lifecycle_fault"`, `"recovery_dimensions"`, `"read_successful"`, `"write_successful"`} {
			if strings.Contains(string(data), field) {
				t.Fatalf("schema %d invented %s", version, field)
			}
		}
		var decoded Report
		if err = json.Unmarshal(data, &decoded); err != nil || decoded.SchemaVersion != version || decoded.LifecycleFault != nil || decoded.Dimensions != nil {
			t.Fatalf("legacy report decoding changed: %v", err)
		}
	}
}

func TestProbeJSONSeparatesMeasuredWriteFailureFromUnattemptedRead(t *testing.T) {
	failed := false
	data, err := json.Marshal(Probe{Successful: false, WriteSuccessful: &failed})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"write_successful":false`) || strings.Contains(string(data), `"read_successful"`) {
		t.Fatalf("unattempted read inferred from failed write: %s", data)
	}
}
