//go:build integration

package experiment

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/report"
)

func networkIntegrationConfig() config.Config {
	c := integrationConfig()
	c.Name = "primary-network-isolation"
	c.Experiment.Fault = "primary-isolate"
	c.Experiment.IsolationDuration = config.Duration(8 * time.Second)
	c.Timeouts.Restoration = config.Duration(30 * time.Second)
	c.Timeouts.Reconciliation = config.Duration(20 * time.Second)
	return c
}

func assertNetworkRestorationEvidence(t *testing.T, r *report.Report) {
	t.Helper()
	n := r.NetworkFault
	if n == nil || n.Receipt == nil || n.Receipt.DisconnectedAt.IsZero() || n.Receipt.RestoredAt.IsZero() || !n.ConnectivityRestored || n.RestoreRequestedAt == nil || n.RestoreCompletedAt == nil || n.RestoreDurationMS == nil || *n.RestoreDurationMS < 0 {
		t.Fatalf("verified disconnect/restoration evidence missing: %+v", n)
	}
	if r.Environment == nil || !r.Environment.NetworkIsolation || n.Receipt.RunID != r.RunID || n.Receipt.NetworkID != r.Environment.NetworkID {
		t.Fatal("network receipt does not identify the isolated experiment")
	}
	if n.Receipt.HostAddr != n.Receipt.Node.Addr || n.Receipt.Node.InternalAddr != n.Receipt.OriginalIP+":6379" {
		t.Fatal("network receipt lost the original published/private endpoints")
	}
	before := n.Receipt.BeforeProcess
	if !before.ContainerRunning || !before.RedisPONG || before.RedisRunID == "" || before.ContainerPID < 1 || before.RedisPID < 1 || !before.NetworkConnected || before.NetworkCount != 1 {
		t.Fatal("receipt lacks a live connected pre-fault Redis identity")
	}
	root, err := filepath.Abs("../../reports/ci")
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(root, r.RunID, "report.json"))
	if err != nil {
		t.Fatal(err)
	}
	var saved report.Report
	if err := json.Unmarshal(data, &saved); err != nil {
		t.Fatal(err)
	}
	if saved.NetworkFault == nil || saved.NetworkFault.Receipt == nil || !saved.NetworkFault.ConnectivityRestored || saved.NetworkFault.RestoreCompletedAt == nil || !saved.NetworkFault.Receipt.RestoredAt.Equal(n.Receipt.RestoredAt) || saved.NetworkFault.Reconciled != n.Reconciled || saved.NetworkFault.ReconciliationTimedOut != n.ReconciliationTimedOut || len(saved.NetworkFault.ProcessSamples) != len(n.ProcessSamples) {
		t.Fatal("saved report lost network restoration/process evidence")
	}
}

func assertHistoricalMajorityRecovery(t *testing.T, r *report.Report) {
	t.Helper()
	if r.Initial == nil || r.Selection == nil || r.NetworkFault == nil || r.NetworkFault.RestoreRequestedAt == nil || r.Recovery.ReplicaPromoted == nil || r.Recovery.CombinedStable == nil || r.Recovery.TopologyConverged == nil || r.Recovery.FullCoverage == nil || r.Recovery.FailureDetected == nil {
		t.Fatalf("majority-side historical recovery missing: %+v", r.Recovery)
	}
	if !r.Recovery.CombinedStable.At.Before(*r.NetworkFault.RestoreRequestedAt) {
		t.Fatal("restored connectivity was mistaken for recovery during isolation")
	}
	if !processSurvivedIsolation(r.NetworkFault) {
		t.Fatal("Redis was not observed alive with unchanged identity while detached")
	}
	found := false
	for _, sample := range r.Samples {
		if sample.Phase != "fault" || sample.Excluded || !sample.TopologyOK || !sample.Probe.Successful {
			continue
		}
		if ok, why := cluster.RecoveryWithBaseline(sample.Snapshot, *r.Selection, *r.Initial); !ok {
			t.Fatalf("claimed majority recovery contradicts saved topology: %v", why)
		}
		found = true
	}
	if !found {
		t.Fatal("historical majority recovery has no supporting all-shard topology/client sample")
	}
}

func TestDockerPrimaryNetworkIsolation(t *testing.T) {
	r, err := runIntegration(t, networkIntegrationConfig(), environment.New())
	assertNetworkRestorationEvidence(t, r)
	assertHistoricalMajorityRecovery(t, r)
	n := r.NetworkFault
	if !n.Reconciled || n.ReconciliationTimedOut || n.ReconciledAt == nil || n.RedundancyRestoredAt == nil || n.StableReconciliationMS == nil || !r.Recovery.RedundancyRestored || !r.Recovery.CombinedRecovered || !r.Recovery.ClientRecovered || !r.Recovery.TopologyRecovered || r.Recovery.ProcessRestarted {
		t.Fatalf("post-restoration reconciliation incomplete: network=%+v recovery=%+v", n, r.Recovery)
	}
	if r.Final == nil {
		t.Fatal("final six-node topology missing")
	}
	if ok, why := cluster.ReconciledWithBaseline(*r.Final, *r.Selection, *r.Initial); !ok {
		t.Fatalf("final topology does not preserve original identities and expected failover: %v", why)
	}
	restoredProcessObserved := false
	before := n.Receipt.BeforeProcess
	for _, sample := range n.ProcessSamples {
		p := sample.Evidence
		if sample.Error == "" && p.NetworkConnected && p.NetworkCount == 1 && p.ContainerRunning && p.RedisPONG && p.ContainerPID == before.ContainerPID && p.ContainerStartedAt == before.ContainerStartedAt && p.RedisPID == before.RedisPID && p.RedisRunID == before.RedisRunID {
			restoredProcessObserved = true
		}
	}
	if !restoredProcessObserved {
		t.Fatal("unchanged Redis process not observed after network restoration")
	}
	if !r.Validation.Complete || r.Validation.Expected == 0 || r.Validation.Checked != r.Validation.Expected || r.Validation.Unreadable != 0 || r.Validation.Corrupt != 0 || r.WorkloadStats.MismatchedReads != 0 {
		t.Fatalf("validation incomplete or corrupted: %+v", r.Validation)
	}
	// Redis asynchronous replication can lose acknowledged SETs during failover.
	// Permit partial only for observed missing records, never hidden topology,
	// process, timeout, or restoration failures.
	if r.Validation.Missing > 0 || r.WorkloadStats.MissingReads > 0 {
		if r.Outcome != "partial" || err == nil || r.Recovery.DataVerified {
			t.Fatalf("observed asynchronous write loss misclassified: outcome=%s err=%v", r.Outcome, err)
		}
	} else if err != nil || r.Outcome != "recovered" || !r.Recovery.DataVerified {
		t.Fatalf("verified reconciliation misclassified: outcome=%s err=%v", r.Outcome, err)
	}
	t.Logf("promoted_ms=%.2f majority_stable_ms=%.2f stable_reconciliation_ms=%.2f acknowledged=%d missing=%d", r.Recovery.ReplicaPromoted.SinceFaultMS, r.Recovery.CombinedStable.SinceFaultMS, *n.StableReconciliationMS, r.Validation.Expected, r.Validation.Missing)
}

func TestDockerNetworkReconciliationTimeout(t *testing.T) {
	c := networkIntegrationConfig()
	c.Timeouts.Reconciliation = config.Duration(100 * time.Millisecond)
	r, err := runIntegration(t, c, environment.New())
	assertNetworkRestorationEvidence(t, r)
	assertHistoricalMajorityRecovery(t, r)
	if err == nil || r.Outcome != "partial" || !r.NetworkFault.ReconciliationTimedOut || r.NetworkFault.Reconciled || r.NetworkFault.StableReconciliationMS != nil || r.Recovery.CombinedRecovered {
		t.Fatalf("reconciliation deadline became full recovery: outcome=%s network=%+v err=%v", r.Outcome, r.NetworkFault, err)
	}
}

type interruptedIsolationManager struct {
	*environment.Manager
	cancel                context.CancelFunc
	ambiguousErr          error
	isolationCompleted    bool
	restoreCalls          int
	restoreContextErr     error
	restoreDeadline       bool
	cleanupCalls          int
	cleanupContextErr     error
	cleanupDeadline       bool
	restoredBeforeCleanup bool
	killCalls             int
}

func (m *interruptedIsolationManager) Isolate(ctx context.Context, env *environment.Environment, node environment.Node) (*environment.IsolationReceipt, error) {
	receipt, err := m.Manager.Isolate(ctx, env, node)
	if err != nil {
		return receipt, err
	}
	m.isolationCompleted = true
	if m.cancel != nil {
		m.cancel()
	}
	return receipt, m.ambiguousErr
}
func (m *interruptedIsolationManager) Restore(ctx context.Context, env *environment.Environment, receipt *environment.IsolationReceipt) error {
	m.restoreCalls++
	m.restoreContextErr = ctx.Err()
	_, m.restoreDeadline = ctx.Deadline()
	return m.Manager.Restore(ctx, env, receipt)
}
func (m *interruptedIsolationManager) Cleanup(ctx context.Context, env *environment.Environment) error {
	if m.cleanupCalls == 0 {
		m.cleanupContextErr = ctx.Err()
		_, m.cleanupDeadline = ctx.Deadline()
		m.restoredBeforeCleanup = m.restoreCalls > 0
	}
	m.cleanupCalls++
	return m.Manager.Cleanup(ctx, env)
}
func (m *interruptedIsolationManager) Kill(context.Context, *environment.Environment, environment.Node) error {
	m.killCalls++
	return errors.New("isolation experiment unexpectedly attempted SIGKILL")
}

func assertFreshInterruptedNetworkTeardown(t *testing.T, m *interruptedIsolationManager, r *report.Report) {
	t.Helper()
	if !m.isolationCompleted || m.killCalls != 0 || m.restoreCalls != 1 || m.restoreContextErr != nil || !m.restoreDeadline || m.cleanupCalls != 1 || m.cleanupContextErr != nil || !m.cleanupDeadline || !m.restoredBeforeCleanup {
		t.Fatalf("interrupted isolation did not restore before cleanup with fresh bounded contexts: %+v", m)
	}
	if r == nil || !r.Cleanup.Attempted || !r.Cleanup.Successful || r.NetworkFault == nil || r.NetworkFault.Reconciled {
		t.Fatalf("interrupted isolation has inaccurate recovery/cleanup evidence: %+v", r)
	}
	assertNetworkRestorationEvidence(t, r)
}

func TestDockerAmbiguousNetworkIsolationRestoresAndCleans(t *testing.T) {
	injected := errors.New("injected ambiguous disconnect response")
	m := &interruptedIsolationManager{Manager: environment.New(), ambiguousErr: injected}
	r, err := runIntegration(t, networkIntegrationConfig(), m)
	if !errors.Is(err, injected) || r.Outcome != "failed" || r.FaultCompletedAt != nil {
		t.Fatalf("ambiguous isolation became a completed successful fault: outcome=%s err=%v", r.Outcome, err)
	}
	assertFreshInterruptedNetworkTeardown(t, m, r)
}

func TestDockerNetworkCancellationRestoresAndCleans(t *testing.T) {
	bounded, boundCancel := context.WithTimeout(context.Background(), 4*time.Minute)
	defer boundCancel()
	ctx, cancel := context.WithCancel(bounded)
	defer cancel()
	m := &interruptedIsolationManager{Manager: environment.New(), cancel: cancel}
	root, err := filepath.Abs("../../reports/ci")
	if err != nil {
		t.Fatal(err)
	}
	r, runErr := (Runner{Manager: m}).Run(ctx, networkIntegrationConfig(), root)
	if r != nil {
		t.Cleanup(func() {
			cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), time.Minute)
			defer cleanupCancel()
			if err := m.Manager.Cleanup(cleanupCtx, r.Environment); err != nil {
				t.Error(err)
			}
		})
	}
	if r == nil || !errors.Is(runErr, context.Canceled) || r.Outcome != "aborted" {
		t.Fatalf("cancelled isolation result inaccurate: report=%+v err=%v", r, runErr)
	}
	assertFreshInterruptedNetworkTeardown(t, m, r)
	found, err := environment.New().Discover(bounded, r.RunID)
	if err != nil || found.NetworkID != "" || len(found.Nodes) != 0 {
		t.Fatalf("cancelled isolation leaked owned resources: inventory=%+v err=%v", found, err)
	}
	assertReportFiles(t, root, r.RunID, "aborted")
	t.Logf("run_id=%s outcome=aborted connectivity_restored=true cleanup_verified=true", r.RunID)
}
