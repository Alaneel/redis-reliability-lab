package batch

import (
	"context"
	"math"
	"strings"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/report"
)

func networkFixture(c config.Config, index int) *report.Report {
	r := fixture(c, index)
	r.Environment.NetworkID = "owned-network"
	node := environment.Node{ID: "node-1", ContainerID: "owned-container", Name: "owned-name", Addr: "127.0.0.1:16379", InternalAddr: "10.224.1.2:6379"}
	r.Environment.Nodes = []environment.Node{node}
	r.Selection = &cluster.Selection{Target: cluster.Node{ID: node.ID, Addr: node.Addr, InternalAddr: node.InternalAddr}}
	disconnected := r.FaultCompletedAt.Add(-time.Millisecond)
	restored := r.FaultCompletedAt.Add(8 * time.Second)
	restoredComplete := restored.Add(100 * time.Millisecond)
	reconciled := restoredComplete.Add(time.Second)
	before := environment.ProcessEvidence{
		ContainerRunning: true, ContainerPID: 123, ContainerStartedAt: "original-start",
		RedisPONG: true, RedisPID: 1, RedisRunID: "original-redis-id", NetworkConnected: true, NetworkCount: 1,
	}
	during := before
	during.ObservedAt = r.FaultCompletedAt.Add(time.Second)
	during.NetworkConnected, during.NetworkCount = false, 0
	r.NetworkFault = &report.NetworkFault{
		Receipt:        &environment.IsolationReceipt{RunID: r.RunID, NetworkID: r.Environment.NetworkID, Node: node, OriginalIP: "10.224.1.2", HostAddr: node.Addr, BeforeProcess: before, DisconnectedAt: disconnected},
		ProcessSamples: []report.ProcessSample{{Evidence: during}}, ConnectivityRestored: true, Reconciled: true,
		RestoreRequestedAt: &restored, RestoreCompletedAt: &restoredComplete, RestoreDurationMS: number(100), StableReconciliationMS: number(1000),
		ReconciledAt:         &report.Milestone{At: reconciled, SinceFaultMS: 9200},
		RedundancyRestoredAt: &report.Milestone{At: reconciled, SinceFaultMS: 9200},
	}
	r.Recovery.RedundancyRestored = true
	r.FinishedAt = reconciled.Add(time.Second)
	return r
}

func TestRecoveredNetworkTrialRequiresRestorationAndProcessEvidence(t *testing.T) {
	for _, test := range []struct {
		name   string
		mutate func(*report.Report)
	}{
		{"schema_one", func(r *report.Report) { r.SchemaVersion = 1 }},
		{"no_network_report", func(r *report.Report) { r.NetworkFault = nil }},
		{"no_receipt", func(r *report.Report) { r.NetworkFault.Receipt = nil }},
		{"different_receipt_run", func(r *report.Report) { r.NetworkFault.Receipt.RunID = "other-run" }},
		{"different_receipt_network", func(r *report.Report) { r.NetworkFault.Receipt.NetworkID = "other-network" }},
		{"no_selection", func(r *report.Report) { r.Selection = nil }},
		{"different_receipt_target", func(r *report.Report) { r.NetworkFault.Receipt.Node.ID = "node-2" }},
		{"different_container", func(r *report.Report) { r.NetworkFault.Receipt.Node.ContainerID = "other-container" }},
		{"different_original_ip", func(r *report.Report) { r.NetworkFault.Receipt.OriginalIP = "10.224.1.3" }},
		{"different_host_endpoint", func(r *report.Report) { r.NetworkFault.Receipt.HostAddr = "127.0.0.1:16380" }},
		{"no_disconnection_observed", func(r *report.Report) { r.NetworkFault.Receipt.DisconnectedAt = time.Time{} }},
		{"connectivity_not_restored", func(r *report.Report) { r.NetworkFault.ConnectivityRestored = false }},
		{"no_restore_start", func(r *report.Report) { r.NetworkFault.RestoreRequestedAt = nil }},
		{"no_restore_completion", func(r *report.Report) { r.NetworkFault.RestoreCompletedAt = nil }},
		{"no_restore_duration", func(r *report.Report) { r.NetworkFault.RestoreDurationMS = nil }},
		{"negative_restore_duration", func(r *report.Report) { r.NetworkFault.RestoreDurationMS = number(-1) }},
		{"no_reconciliation", func(r *report.Report) { r.NetworkFault.Reconciled = false }},
		{"reconciliation_timed_out", func(r *report.Report) { r.NetworkFault.ReconciliationTimedOut = true }},
		{"no_reconciliation_milestone", func(r *report.Report) { r.NetworkFault.ReconciledAt = nil }},
		{"no_stable_reconciliation_duration", func(r *report.Report) { r.NetworkFault.StableReconciliationMS = nil }},
		{"no_redundancy_milestone", func(r *report.Report) { r.NetworkFault.RedundancyRestoredAt = nil }},
		{"redundancy_not_restored", func(r *report.Report) { r.Recovery.RedundancyRestored = false }},
		{"no_process_sample", func(r *report.Report) { r.NetworkFault.ProcessSamples = nil }},
		{"process_probe_failed", func(r *report.Report) { r.NetworkFault.ProcessSamples[0].Error = "exec failed" }},
		{"process_not_running", func(r *report.Report) { r.NetworkFault.ProcessSamples[0].Evidence.ContainerRunning = false }},
		{"redis_not_live", func(r *report.Report) { r.NetworkFault.ProcessSamples[0].Evidence.RedisPONG = false }},
		{"still_connected", func(r *report.Report) { r.NetworkFault.ProcessSamples[0].Evidence.NetworkConnected = true }},
		{"another_network", func(r *report.Report) { r.NetworkFault.ProcessSamples[0].Evidence.NetworkCount = 1 }},
		{"container_pid_changed", func(r *report.Report) { r.NetworkFault.ProcessSamples[0].Evidence.ContainerPID++ }},
		{"container_start_changed", func(r *report.Report) { r.NetworkFault.ProcessSamples[0].Evidence.ContainerStartedAt = "changed" }},
		{"redis_pid_changed", func(r *report.Report) { r.NetworkFault.ProcessSamples[0].Evidence.RedisPID++ }},
		{"redis_identity_changed", func(r *report.Report) { r.NetworkFault.ProcessSamples[0].Evidence.RedisRunID = "changed" }},
		{"zero_identity", func(r *report.Report) {
			r.NetworkFault.Receipt.BeforeProcess.RedisRunID = ""
			r.NetworkFault.ProcessSamples[0].Evidence.RedisRunID = ""
		}},
	} {
		t.Run(test.name, func(t *testing.T) {
			c := config.Default()
			c.Experiment.Fault = "primary-isolate"
			s, err := Run(context.Background(), c, 1, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
				r := networkFixture(c, 1)
				test.mutate(r)
				save(t, dir, r)
				return r, nil
			})
			if err == nil || !strings.Contains(err.Error(), "recovered network trial") || s.Outcomes["recovered"] != 0 || s.Outcomes["failed"] != 1 || s.RecoverySuccessFraction != 0 || s.IncompleteEvidenceTrials != 1 {
				t.Fatalf("invalid network recovery counted: %+v err=%v", s, err)
			}
		})
	}
}

func TestRecoveredNetworkEvidenceAcceptsOneValidDisconnectedObservation(t *testing.T) {
	c := config.Default()
	c.Experiment.Fault = "primary-isolate"
	s, err := Run(context.Background(), c, 1, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		r := networkFixture(c, 1)
		failed := r.NetworkFault.ProcessSamples[0]
		failed.Error = "earlier probe failed"
		r.NetworkFault.ProcessSamples = append([]report.ProcessSample{failed}, r.NetworkFault.ProcessSamples...)
		save(t, dir, r)
		return r, nil
	})
	if err != nil || s.Outcomes["recovered"] != 1 || !s.Trials[0].EvidenceComplete {
		t.Fatalf("valid network evidence rejected: %+v err=%v", s, err)
	}
}

func TestPartialNetworkTrialsKeepObservedEventsWithoutClaimingReconciliation(t *testing.T) {
	c := config.Default()
	c.Experiment.Fault = "primary-isolate"
	s, err := Run(context.Background(), c, 1, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		r := networkFixture(c, 1)
		r.Outcome = "partial"
		r.NetworkFault.Reconciled = false
		r.NetworkFault.ReconciliationTimedOut = true
		r.NetworkFault.StableReconciliationMS = nil
		r.Recovery.RedundancyRestored = false
		save(t, dir, r)
		return r, nil
	})
	if err == nil || s.Outcomes["partial"] != 1 || s.IncompleteEvidenceTrials != 0 || s.Events["replica_promoted"].Observed != 1 || s.RecoverySuccessFraction != 0 {
		t.Fatalf("partial network evidence lost: %+v err=%v", s, err)
	}
}

func TestNetworkEvidenceRejectsNonfiniteMeasuredDurations(t *testing.T) {
	c := config.Default()
	c.Experiment.Fault = "primary-isolate"
	r := networkFixture(c, 1)
	r.NetworkFault.StableReconciliationMS = number(math.Inf(1))
	if recoveredNetworkEvidence(r) {
		t.Fatal("nonfinite reconciliation duration accepted")
	}
}
