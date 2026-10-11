//go:build integration

package experiment

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/report"
)

func lifecycleIntegrationConfig(fault string) config.Config {
	c := integrationConfig()
	c.Experiment.Fault = fault
	c.Experiment.Repair = fault == "primary-kill"
	c.Experiment.DownDuration = config.Duration(3 * time.Second)
	c.Experiment.PostRecovery = config.Duration(time.Second)
	c.Timeouts.Repair = config.Duration(30 * time.Second)
	c.Workload.RatePerSecond = 200
	return c
}

func assertRestoredLifecycle(t *testing.T, r *report.Report, replica bool) {
	t.Helper()
	l := r.LifecycleFault
	if l == nil || r.SchemaVersion != report.LifecycleSchemaVersion || r.Initial == nil || r.Final == nil || r.Selection == nil || !l.RepairCompleted || l.RepairTimedOut || l.RepairRequestedAt == nil || l.RepairCompletedAt == nil || l.RepairActionDurationMS == nil || l.StableRedundancy == nil || l.StableRedundancySinceRepairMS == nil || l.FaultSnapshot == nil || l.RepairSnapshot == nil || l.TargetReachable == nil || l.MembershipRejoined == nil || l.LinkUp == nil || l.SyncComplete == nil {
		t.Fatalf("missing independently measured lifecycle evidence: %+v", l)
	}
	if !r.Recovery.RedundancyRestored || !r.Recovery.CombinedRecovered || !r.Recovery.TopologyRecovered || !r.Recovery.ClientRecovered || r.Recovery.TimedOut || *l.StableRedundancySinceRepairMS < float64(r.Config.Experiment.StableFor.Value())/float64(time.Millisecond) {
		t.Fatalf("physical action was mistaken for sustained six-node recovery: %+v", r.Recovery)
	}
	if ok, why := lifecycleRestored(*r.Final, *r.Selection, *r.Initial, replica); !ok {
		t.Fatalf("final topology contradicts lifecycle recovery: %v", why)
	}
	if r.Dimensions == nil || l.Dimensions == nil || r.Dimensions.DataVerified == nil || *r.Dimensions.DataVerified != r.Recovery.DataVerified || r.Dimensions.RedundancyRestored == nil || !*r.Dimensions.RedundancyRestored || r.Dimensions.PersistenceRecovered != nil || r.Dimensions.TargetRole != "slave" || r.Dimensions.ReadAvailable == nil || !*r.Dimensions.ReadAvailable || r.Dimensions.WriteAvailable == nil || !*r.Dimensions.WriteAvailable {
		t.Fatalf("final independent dimensions contradict measured SET lifecycle: %+v", r.Dimensions)
	}
	for _, sample := range r.Samples {
		if sample.Dimensions != nil && sample.Dimensions.DataVerified != nil {
			t.Fatal("post-validation data verdict was retroactively assigned to a pre-validation topology sample")
		}
	}
	if !r.Validation.Complete || r.Validation.Expected == 0 || r.Validation.Expected != r.Validation.Checked || r.Validation.Unreadable != 0 || r.Validation.Unverified != 0 || r.Validation.Corrupt != 0 {
		t.Fatalf("acknowledged ledger validation incomplete or corrupt: %+v", r.Validation)
	}
	if l.Resynchronization == nil {
		t.Fatal("resynchronization counter evidence/limitations absent")
	}
	t.Logf("repair_action_ms=%.3f stable_redundancy_since_repair_ms=%.3f sync_evidence=%s acknowledged=%d missing=%d", *l.RepairActionDurationMS, *l.StableRedundancySinceRepairMS, l.Resynchronization.Inference, r.Validation.Expected, r.Validation.Missing)
}

func assertOwnedRestart(t *testing.T, r *report.Report) {
	t.Helper()
	l := r.LifecycleFault
	x := l.RestartReceipt
	if !r.Environment.RetainData || len(r.Environment.Volumes) != 6 || x == nil || x.RunID != r.RunID || x.NodeID != l.Target.ID || x.OriginalRedisID != l.TargetRedisID || x.After.RedisClusterID != l.TargetRedisID || !x.After.ContainerRunning || !x.After.RedisPONG || x.After.RedisRunID == "" || x.Before.ContainerRunning || x.Before.ContainerStartedAt == x.After.ContainerStartedAt || !r.Recovery.ProcessRestarted {
		t.Fatalf("restart did not preserve owned Redis identity across process replacement: receipt=%+v environment=%+v", x, r.Environment)
	}
}

func TestDockerReplicaCrashRestart(t *testing.T) {
	r, err := runIntegration(t, lifecycleIntegrationConfig("replica-kill"), environment.New())
	assertRestoredLifecycle(t, r, true)
	assertOwnedRestart(t, r)
	if err != nil || r.Outcome != "recovered" || r.Recovery.ReplicaPromoted != nil || r.LifecycleFault.Target.ID != r.Selection.Replica.ID || r.LifecycleFault.TargetRedisID != r.Selection.ReplicaID || r.LifecycleFault.ReplicationInterrupted == nil || r.Validation.Missing != 0 || r.WorkloadStats.MissingReads != 0 {
		t.Fatalf("replica-only failure changed primary ownership or hid correctness failure: outcome=%s lifecycle=%+v err=%v", r.Outcome, r.LifecycleFault, err)
	}
}

func TestDockerReplicaNetworkIsolation(t *testing.T) {
	r, err := runIntegration(t, lifecycleIntegrationConfig("replica-isolate"), environment.New())
	assertRestoredLifecycle(t, r, true)
	assertNetworkRestorationEvidence(t, r)
	if err != nil || r.Outcome != "recovered" || !processSurvivedIsolation(r.NetworkFault) || r.Recovery.ProcessRestarted || r.Recovery.ReplicaPromoted != nil || r.LifecycleFault.RestartReceipt != nil || r.LifecycleFault.Target.ID != r.Selection.Replica.ID || r.Validation.Missing != 0 || r.WorkloadStats.MissingReads != 0 {
		t.Fatalf("replica network fault conflated process restart/promotion or verification: outcome=%s err=%v", r.Outcome, err)
	}
	// A brief administrative disconnection can end before Redis's existing TCP
	// links time out. Only INFO evidence establishes an observed link/count
	// interruption; detached endpoint proof must not fabricate that milestone.
	if r.LifecycleFault.ReplicationInterrupted == nil {
		t.Log("owned interface was detached while Redis survived; INFO-level replication interruption was unobserved")
	}
}

func TestDockerPrimaryRepairRestoresRedundancy(t *testing.T) {
	c := lifecycleIntegrationConfig("primary-kill")
	c.Workload.WritePolicy = "wait"
	r, err := runIntegration(t, c, environment.New())
	assertRestoredLifecycle(t, r, false)
	assertOwnedRestart(t, r)
	l := r.LifecycleFault
	if r.Recovery.ReplicaPromoted == nil || r.Recovery.CombinedStable == nil || !r.Recovery.CombinedStable.At.Before(*l.RepairRequestedAt) || !l.RepairCompletedAt.Before(l.StableRedundancy.At) || r.Recovery.WritePolicyRecovered == nil || !*r.Recovery.WritePolicyRecovered || r.Recovery.WritePolicyStable == nil {
		t.Fatalf("former primary repair bypassed failover/service or WAIT recovery: %+v", r.Recovery)
	}
	if r.Validation.Missing > 0 || r.WorkloadStats.MissingReads > 0 {
		if r.Outcome != "partial" || err == nil || r.Recovery.DataVerified {
			t.Fatal("observed asynchronous data loss became full recovery")
		}
	} else if err != nil || r.Outcome != "recovered" {
		t.Fatalf("verified full repair misclassified: outcome=%s err=%v", r.Outcome, err)
	}
	if r.Validation.WaitSatisfiedSubset.Expected == 0 || !r.Validation.WaitSatisfiedSubset.Complete || r.Validation.WaitSatisfiedSubset.Checked != r.Validation.WaitSatisfiedSubset.Expected {
		t.Fatal("WAIT-satisfied subset was not retained and validated after redundancy repair")
	}
}

func TestDockerReplicaRepairTimeout(t *testing.T) {
	c := lifecycleIntegrationConfig("replica-kill")
	c.Experiment.StableFor = config.Duration(2 * time.Second)
	c.Experiment.DownDuration = config.Duration(4 * time.Second)
	c.Timeouts.Repair = config.Duration(100 * time.Millisecond)
	r, err := runIntegration(t, c, environment.New())
	if err == nil || r.Outcome == "recovered" || r.LifecycleFault == nil || !r.LifecycleFault.RepairTimedOut || r.LifecycleFault.StableRedundancy != nil || r.Recovery.CombinedRecovered || r.Recovery.RedundancyRestored {
		t.Fatalf("repair action/synchronization deadline became full restoration: outcome=%s lifecycle=%+v err=%v", r.Outcome, r.LifecycleFault, err)
	}
}

type interruptedRestartManager struct {
	*environment.Manager
	err             error
	cancel          context.CancelFunc
	started         bool
	cleanupCalls    int
	cleanupErr      error
	cleanupDeadline bool
}

func (m *interruptedRestartManager) Restart(ctx context.Context, env *environment.Environment, n environment.Node) (*environment.RestartReceipt, error) {
	x, err := m.Manager.Restart(ctx, env, n)
	if err != nil {
		return x, err
	}
	m.started = true
	if m.cancel != nil {
		m.cancel()
	}
	return x, m.err
}

func (m *interruptedRestartManager) Cleanup(ctx context.Context, env *environment.Environment) error {
	if m.cleanupCalls == 0 {
		m.cleanupErr = ctx.Err()
		_, m.cleanupDeadline = ctx.Deadline()
	}
	m.cleanupCalls++
	return m.Manager.Cleanup(ctx, env)
}

func TestDockerReplicaRepairAmbiguousFailureCleansVolumes(t *testing.T) {
	injected := errors.New("injected ambiguous restart response")
	m := &interruptedRestartManager{Manager: environment.New(), err: injected}
	r, err := runIntegration(t, lifecycleIntegrationConfig("replica-kill"), m)
	if !errors.Is(err, injected) || r.Outcome != "failed" || !m.started || r.LifecycleFault.RestartReceipt == nil || !r.LifecycleFault.RestartReceipt.After.ContainerRunning || r.LifecycleFault.RepairCompleted || r.LifecycleFault.StableRedundancy != nil || r.Recovery.CombinedRecovered || !r.Validation.StartedAt.IsZero() || m.cleanupCalls != 1 || m.cleanupErr != nil || !m.cleanupDeadline {
		t.Fatalf("ambiguous applied restart lost failure or skipped independent persistent cleanup: outcome=%s lifecycle=%+v err=%v", r.Outcome, r.LifecycleFault, err)
	}
}

func TestDockerCancellationDuringReplicaSynchronization(t *testing.T) {
	bounded, boundCancel := context.WithTimeout(context.Background(), 4*time.Minute)
	defer boundCancel()
	ctx, cancel := context.WithCancel(bounded)
	defer cancel()
	m := &interruptedRestartManager{Manager: environment.New(), cancel: cancel}
	root, err := filepath.Abs("../../reports/ci")
	if err != nil {
		t.Fatal(err)
	}
	r, err := (Runner{Manager: m}).Run(ctx, lifecycleIntegrationConfig("replica-kill"), root)
	if r == nil {
		t.Fatalf("missing cancellation evidence: %v", err)
	}
	t.Cleanup(func() {
		ctx, cleanupCancel := context.WithTimeout(context.Background(), time.Minute)
		defer cleanupCancel()
		if err := m.Manager.Cleanup(ctx, r.Environment); err != nil {
			t.Error(err)
		}
	})
	if !errors.Is(err, context.Canceled) || r.Outcome != "aborted" || !m.started || !r.Cleanup.Successful || m.cleanupCalls != 1 || m.cleanupErr != nil || !m.cleanupDeadline || r.LifecycleFault.StableRedundancy != nil || r.Recovery.CombinedRecovered || !r.Validation.StartedAt.IsZero() {
		t.Fatalf("synchronization cancellation claimed restoration or lost independent cleanup: outcome=%s err=%v", r.Outcome, err)
	}
	found, err := environment.New().Discover(bounded, r.RunID)
	if err != nil || len(found.Nodes) != 0 || found.NetworkID != "" || len(found.Volumes) != 0 {
		t.Fatalf("canceled repair leaked owned persistent resources: inventory=%+v err=%v", found, err)
	}
	assertReportFiles(t, root, r.RunID, "aborted")
	t.Logf("run_id=%s outcome=aborted restart_applied=true cleanup_verified=true", r.RunID)
}
