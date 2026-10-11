package report

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/metrics"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

func persistenceReport(mode string) Report {
	r := sampleReport()
	r.SchemaVersion = 3
	r.Outcome = "partial"
	r.Config.Name = "frozen-ledger-control"
	drained := r.StartedAt.Add(time.Second)
	stopped := drained.Add(time.Second)
	firstRestart, restartsComplete := stopped.Add(time.Millisecond), stopped.Add(time.Second)
	dataChecked, diagnostics := restartsComplete.Add(time.Second), restartsComplete.Add(2*time.Second)
	interval := 1000.0
	falseFlag := false
	r.PersistenceFault = &PersistenceFault{Mode: mode, FailureModel: PersistenceFailureModel,
		TrafficDrainedAt: &drained, LedgerFrozenAt: &drained, AcknowledgedFrozen: 20, AmbiguousFrozen: 2,
		AllStoppedAt: &stopped, RestartRequestedAt: &firstRestart, RestartCompletedAt: &restartsComplete, DataCheckedAt: &dataChecked, DiagnosticsRequestedAt: &diagnostics,
		DrainToFirstKillMS: &interval, DrainToLastKillMS: &interval, AllStoppedToFirstRestartMS: &interval,
		ReplicaReady: &Milestone{At: restartsComplete, SinceFaultMS: 1000}, TopologyStable: &Milestone{At: restartsComplete, SinceFaultMS: 2000},
		FrozenValidation: &workload.Validation{StartedAt: restartsComplete.Add(750 * time.Millisecond), CompletedAt: dataChecked, Expected: 20, Checked: 20, Missing: 20, Complete: true}}
	r.Validation = *r.PersistenceFault.FrozenValidation
	r.Dimensions = &RecoveryDimensions{PersistenceRecovered: &falseFlag, DataVerified: &falseFlag}
	r.Recovery.TopologyRecovered, r.Recovery.RedundancyRestored, r.Recovery.ProcessRestarted = true, true, true
	r.Metrics.Phases["fault"] = metrics.PhaseSummary{Attempts: 0, DurationSeconds: 2}
	r.Metrics.Phases["repair"] = metrics.PhaseSummary{Attempts: 0, DurationSeconds: 2}
	for i := range 6 {
		node := environment.Node{ID: fmt.Sprintf("node-%d", i+1), RedisClusterID: strings.Repeat(string(rune('a'+i)), 40)}
		process := environment.ProcessEvidence{ObservedAt: drained, ContainerRunning: true, ContainerRestartPolicy: "no", RedisPONG: true, RedisClusterID: node.RedisClusterID, RedisRunID: "old-process", DataSizeKiB: 4,
			RawPersistence: "loading:0\r\naof_enabled:0\r\n", RawPersistenceConfig: "appendonly\nno\nsave\n", RawKeyspace: "db0:keys=20,expires=0,avg_ttl=0\r\n", RawDataFiles: "nodes.conf 1000 bytes\nappendonlydir/appendonly.aof.1.incr.aof 5000 bytes"}
		r.PersistenceFault.Prepared = append(r.PersistenceFault.Prepared, NodeProcessSample{Node: node, OriginalRedisID: node.RedisClusterID, Evidence: process})
		stoppedProof := environment.ProcessEvidence{ObservedAt: stopped, ContainerRunning: false}
		r.PersistenceFault.KillActions = append(r.PersistenceFault.KillActions, NodeAction{Node: node, RequestedAt: drained, CompletedAt: &stopped, Stopped: &stoppedProof})
		r.PersistenceFault.Stopped = append(r.PersistenceFault.Stopped, NodeProcessSample{Node: node, OriginalRedisID: node.RedisClusterID, Evidence: stoppedProof, Error: environment.ErrProcessStopped.Error()})
		r.PersistenceFault.RestartActions = append(r.PersistenceFault.RestartActions, NodeAction{Node: node, RequestedAt: firstRestart, CompletedAt: &restartsComplete})
		process.ObservedAt, process.RedisRunID = restartsComplete, "new-process"
		r.PersistenceFault.Restarts = append(r.PersistenceFault.Restarts, &environment.RestartReceipt{NodeID: node.ID, Node: node, OriginalRedisID: node.RedisClusterID, RequestedAt: firstRestart, CompletedAt: restartsComplete, Before: stoppedProof, After: process})
		process.ObservedAt = restartsComplete.Add(500 * time.Millisecond)
		r.PersistenceFault.Restored = append(r.PersistenceFault.Restored, NodeProcessSample{Node: node, OriginalRedisID: node.RedisClusterID, Evidence: process})
	}
	return r
}

func TestPersistenceJSONPreservesFrozenEvidenceAndUnknownRecovery(t *testing.T) {
	r := persistenceReport("aof-everysec")
	r.PersistenceFault.RestartActions[0].CompletedAt = nil
	r.PersistenceFault.RestartActions[0].Error = "start command returned no receipt"
	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	var decoded Report
	if err = json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	p := decoded.PersistenceFault
	if p == nil || p.AcknowledgedFrozen != 20 || p.AmbiguousFrozen != 2 || len(p.Prepared) != 6 || len(p.KillActions) != 6 || len(p.Stopped) != 6 || len(p.Restarts) != 6 || len(p.RestartActions) != 6 || len(p.Restored) != 6 {
		t.Fatalf("frozen process lifecycle evidence lost: %+v", p)
	}
	if p.PersistenceRecoveryObserved != nil || p.PersistenceRecoveryAt != nil || p.DiagnosticServiceStable != nil || p.RestartActions[0].CompletedAt != nil || p.RestartActions[0].Error == "" {
		t.Fatal("restart/topology readiness fabricated persistence or diagnostic completion")
	}
	if p.Prepared[0].Node.ID != "node-1" || p.Prepared[0].OriginalRedisID != p.Restarts[0].OriginalRedisID || p.Restarts[0].After.RedisRunID != "new-process" || p.KillActions[0].Stopped == nil || p.KillActions[0].Stopped.ContainerRunning || p.AllStoppedAt == nil {
		t.Fatal("original identity, stopped barrier or changed process lost")
	}
	if p.FrozenValidation == nil || p.FrozenValidation.Missing != 20 || !p.FrozenValidation.Complete || p.DataCheckedAt == nil || p.DiagnosticsRequestedAt == nil || !p.DiagnosticsRequestedAt.After(*p.DataCheckedAt) {
		t.Fatal("complete-but-missing validation or pre-diagnostic chronology lost")
	}
	if !p.Restored[0].Evidence.ObservedAt.After(p.Restarts[0].After.ObservedAt) || p.Restored[0].Evidence.ObservedAt.After(p.FrozenValidation.StartedAt) || p.Restored[0].Evidence.RedisRunID != p.Restarts[0].After.RedisRunID {
		t.Fatal("later post-sync runtime evidence overwrote the physical receipt or crossed validation")
	}
	var doc map[string]json.RawMessage
	if err = json.Unmarshal(data, &doc); err != nil {
		t.Fatal(err)
	}
	var fault map[string]json.RawMessage
	if err = json.Unmarshal(doc["persistence_fault"], &fault); err != nil {
		t.Fatal(err)
	}
	for _, unknown := range []string{"persistence_recovery_observed", "persistence_recovery_at", "diagnostic_service_stable"} {
		if _, exists := fault[unknown]; exists {
			t.Fatalf("unknown %s did not remain absent", unknown)
		}
	}
	falseFlag := false
	r.PersistenceFault.PersistenceRecoveryObserved = &falseFlag
	data, err = json.Marshal(r)
	if err != nil || !strings.Contains(string(data), `"persistence_recovery_observed":false`) {
		t.Fatal("measured failed AOF recovery omitted")
	}
}

func TestPersistenceMarkdownDoesNotTurnNoDemandOrRestartsIntoRecovery(t *testing.T) {
	r := persistenceReport("aof-everysec")
	r.WorkloadStats.WritePolicy = "set"
	r.Metrics.Phases["fault"] = metrics.PhaseSummary{DurationSeconds: 2, ByOperation: map[string]metrics.OperationSummary{"set": {}}}
	r.Metrics.Phases["repair"] = metrics.PhaseSummary{DurationSeconds: 2, ByOperation: map[string]metrics.OperationSummary{"set": {}}}
	markdown := r.Markdown()
	for _, want := range []string{
		"Configured persistence mode: `aof-everysec`", "Frozen acknowledged SETs: 20. Frozen ambiguous SET outcomes: 2.",
		"No workload requests are offered", "do not measure zero availability or recovery throughput", "Read-only frozen-ledger validation precedes",
		"| Frozen acknowledged writes verified | false |", "AOF frozen-value recovery observed: unobserved",
		"missing 20; corrupt 0; unreadable 0; unverified 0; complete true", "observed loss remains a data failure",
		"| Post-validation diagnostic service sustained | unobserved |", "No primary promotion is required",
		"same-container restart with retained managed volumes", "does not remove Docker VM/host page cache", "background fsync to progress",
		"Observed persistence configuration:", "appendonly\nno", "Observed INFO persistence:", "aof_enabled:0", "Original Redis Cluster ID:", "Observed INFO keyspace:", "Observed managed data files:", "appendonly.aof.1.incr.aof 5000 bytes", "Restart policy: `no`",
		"Post-synchronization runtime evidence", "follow topology stability and precede frozen-ledger validation", "earlier physical restart receipts remain unchanged",
	} {
		if !strings.Contains(markdown, want) {
			t.Errorf("truthful evidence missing %q:\n%s", want, markdown)
		}
	}
	for _, falseClaim := range []string{"The primary stays stopped", "| fault | 0 | 0.00%", "| repair | 0 | 0.00%", "| fault | 0 | 0.000 |", "| repair | 0 | 0.000 |", "Original replica owns failed primary slots", "AOF frozen-value recovery observed: true"} {
		if strings.Contains(markdown, falseClaim) {
			t.Errorf("misleading claim %q", falseClaim)
		}
	}
	if !strings.Contains(markdown, "| baseline | 2 | 100.00%") {
		t.Fatal("measured baseline was discarded")
	}
}

func TestPersistenceMemoryControlReportsObservedLossAndNoPersistenceEvent(t *testing.T) {
	r := persistenceReport("memory")
	markdown := r.Markdown()
	for _, want := range []string{"Persistence recovery is not applicable", "lost frozen values", "| Persistence recovery verified | false |", "| AOF frozen-value recovery observed | unobserved |", "missing 20"} {
		if !strings.Contains(markdown, want) {
			t.Errorf("memory control semantics missing %q", want)
		}
	}
	if strings.Contains(markdown, "AOF frozen-value recovery observed: false") || r.PersistenceFault.PersistenceRecoveryObserved != nil {
		t.Fatal("memory control treated as failed AOF event")
	}
}

func TestPersistenceIncompleteActionsAndUnverifiedValuesRemainExplicit(t *testing.T) {
	r := persistenceReport("aof-always")
	p := r.PersistenceFault
	p.FrozenValidation, p.DataCheckedAt, p.DiagnosticsRequestedAt, p.LedgerFrozenAt = nil, nil, nil, nil
	p.RestartActions = []NodeAction{{Node: environment.Node{ID: "node-6"}, RequestedAt: r.StartedAt, Error: "ambiguous restart"}, {Node: environment.Node{ID: "node-1"}, RequestedAt: r.StartedAt}}
	p.Restarts = []*environment.RestartReceipt{nil, {NodeID: "node-1"}}
	p.Errors = []string{"six-node restoration timed out"}
	p.RestoreTimedOut = true
	markdown := r.Markdown()
	for _, want := range []string{"Frozen ledger counts: unobserved", "Frozen-ledger validation unobserved; 20 acknowledged values remain unchecked", "| Restart sequence completed |", "| Post-validation diagnostics requested | unobserved |", "ambiguous restart", "Administrative process observation: unobserved", "six-node restoration timed out", "Restoration deadline expired: true"} {
		if !strings.Contains(markdown, want) {
			t.Errorf("missing incomplete evidence %q", want)
		}
	}
	if strings.Index(markdown, "| 1 | node-6 |") > strings.Index(markdown, "| 2 | node-1 |") {
		t.Fatal("actual restart order was sorted")
	}
}

func TestLegacyReportOmitsPersistenceFault(t *testing.T) {
	for _, schema := range []int{1, 2, 3} {
		r := sampleReport()
		r.SchemaVersion = schema
		data, err := json.Marshal(r)
		if err != nil || strings.Contains(string(data), `"persistence_fault"`) {
			t.Fatalf("schema%d invented persistence evidence: %v", schema, err)
		}
	}
}

func TestPersistenceMarkdownSeparatesRestorationAndDiagnosticTimeouts(t *testing.T) {
	r := persistenceReport("aof-always")
	r.PersistenceFault.RestoreTimedOut = true
	r.Recovery.TimedOut = false
	markdown := r.Markdown()
	if !strings.Contains(markdown, "| Restoration deadline expired | true |") || !strings.Contains(markdown, "| Diagnostic deadline expired | false |") {
		t.Fatal("later diagnostic success concealed the restoration deadline")
	}
	r.PersistenceFault.RestoreTimedOut = false
	r.Recovery.TimedOut = true
	markdown = r.Markdown()
	if !strings.Contains(markdown, "| Restoration deadline expired | false |") || !strings.Contains(markdown, "| Diagnostic deadline expired | true |") {
		t.Fatal("diagnostic deadline was mislabeled as a topology restoration timeout")
	}
}
