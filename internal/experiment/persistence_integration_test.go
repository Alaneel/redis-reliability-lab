//go:build integration

package experiment

import (
	"context"
	"errors"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/report"
)

func persistenceIntegrationConfig(mode string) config.Config {
	c := integrationConfig()
	c.Experiment.Fault = "cluster-restart"
	c.Environment.Persistence = mode
	c.Experiment.Baseline = config.Duration(4 * time.Second)
	c.Experiment.DownDuration = config.Duration(time.Second)
	c.Timeouts.Repair = config.Duration(time.Minute)
	c.Timeouts.Recovery = config.Duration(30 * time.Second)
	c.Workload.RatePerSecond = 200
	c.Workload.MaxWrites = 20000
	c.Workload.ValueBytes = 128
	return c
}

func assertPersistenceRestartEvidence(t *testing.T, r *report.Report) {
	t.Helper()
	p := r.PersistenceFault
	if p == nil || r.SchemaVersion != report.LifecycleSchemaVersion || p.Mode != r.Config.Environment.Persistence || p.FailureModel != PersistenceFailureModel || len(p.Prepared) != 6 || len(p.Restored) != 6 || len(p.KillActions) != 6 || len(p.Stopped) != 6 || len(p.RestartActions) != 6 || len(p.Restarts) != 6 || !r.Environment.RetainData || len(r.Environment.Volumes) != 6 {
		t.Fatalf("six-node retained crash/restart evidence incomplete: %+v", p)
	}
	if p.TrafficDrainedAt == nil || p.LedgerFrozenAt == nil || p.PreCrashSnapshot == nil || p.AllStoppedAt == nil || p.RestartRequestedAt == nil || p.RestartCompletedAt == nil || p.DataCheckedAt == nil || p.DiagnosticsRequestedAt == nil || p.FrozenValidation == nil || p.AcknowledgedFrozen == 0 || p.AcknowledgedFrozen != r.WorkloadStats.AcknowledgedWrites || p.AmbiguousFrozen != r.WorkloadStats.AmbiguousWrites || !reflect.DeepEqual(*p.FrozenValidation, r.Validation) {
		t.Fatal("actual frozen ledger, validation or action anchors missing")
	}
	if p.PreCrashSnapshot.Completed.After(*p.TrafficDrainedAt) || p.TrafficDrainedAt.After(*p.LedgerFrozenAt) || p.LedgerFrozenAt.After(p.KillActions[0].RequestedAt) || p.AllStoppedAt.After(*p.RestartRequestedAt) || p.RestartCompletedAt.After(r.Validation.StartedAt) || r.Validation.CompletedAt.After(*p.DataCheckedAt) || p.DataCheckedAt.After(*p.DiagnosticsRequestedAt) || p.DrainToFirstKillMS == nil || p.DrainToLastKillMS == nil || p.AllStoppedToFirstRestartMS == nil {
		t.Fatal("drain/barrier/restart/read-only-validation/diagnostic ordering contradicted")
	}
	kills, restarts, err := persistenceNodeOrder(r.Environment, r.Initial)
	if err != nil {
		t.Fatal(err)
	}
	for i, n := range kills {
		x, stopped := p.Restarts[i], p.Stopped[i]
		if p.KillActions[i].Node != n || p.RestartActions[i].Node != restarts[i] || p.KillActions[i].CompletedAt == nil || p.RestartActions[i].CompletedAt == nil || p.KillActions[i].Error != "" || p.RestartActions[i].Error != "" || p.KillActions[i].Stopped == nil || stopped.Node != n || stopped.Error != environment.ErrProcessStopped.Error() || stopped.Evidence.ContainerRunning || stopped.Evidence.RedisPONG || stopped.Evidence.NetworkConnected || stopped.Evidence.ContainerRestartPolicy != "no" || stopped.Evidence.ObservedAt.IsZero() || !validRestartReceipt(x, r.RunID, restarts[i], restarts[i].RedisClusterID) || !persistenceNewProcess(r, restarts[i], x) {
			t.Fatalf("original process/action/barrier identity mismatch at %d", i)
		}
		if !p.KillActions[5].CompletedAt.Before(stopped.Evidence.ObservedAt) || stopped.Evidence.ObservedAt.After(*p.AllStoppedAt) || x.RequestedAt.Before(*p.AllStoppedAt) {
			t.Fatal("a process restarted before all six stopped proofs")
		}
	}
	for _, prepared := range p.Prepared {
		if prepared.Error != "" || persistenceRuntimeEvidence(prepared.Evidence, p.Mode) != nil || persistencePreparationBusy(prepared.Evidence) {
			t.Fatalf("prepared runtime configuration was not observed quiet and healthy: %+v", prepared)
		}
	}
	for _, restored := range p.Restored {
		if restored.Error != "" || persistenceRuntimeEvidence(restored.Evidence, p.Mode) != nil || persistencePreparationQuiet(restored.Evidence) != nil || !persistenceRestoredProcess(r, restored.Node, restored.Evidence) || restored.Evidence.ObservedAt.After(r.Validation.StartedAt) {
			t.Fatalf("final runtime/config/files were not actually observed on the same synchronized new process before validation: %+v", restored)
		}
	}
	for _, s := range r.Samples {
		if s.Phase == "persistence_recovery" && (!s.Probe.At.IsZero() || s.WritePolicyProbe != nil || (s.Dimensions != nil && (s.Dimensions.ReadAvailable != nil || s.Dimensions.WriteAvailable != nil || s.Dimensions.DataVerified != nil))) {
			t.Fatal("client or value availability invented before read-only frozen validation")
		}
		if s.Phase == "persistence_diagnostics" && s.Probe.At.Before(*p.DataCheckedAt) {
			t.Fatal("diagnostic writes preceded frozen validation")
		}
	}
	for phase, counts := range r.Metrics.Phases {
		if phase != "baseline" && counts.Attempts > 0 {
			t.Fatalf("workload traffic resumed after crash: %s %d", phase, counts.Attempts)
		}
	}
	if r.WorkloadStats.OfferedRequests != r.WorkloadStats.StartedRequests+r.WorkloadStats.DroppedRequests || r.WorkloadStats.StartedRequests != r.WorkloadStats.CompletedRequests || r.WorkloadStats.CompletedRequests != r.Metrics.Totals.Attempts {
		t.Fatal("drained workload pacing or terminal counts do not reconcile")
	}
	if p.ReplicaReady == nil || p.TopologyStable == nil || p.StableRedundancy == nil || p.DiagnosticServiceStable == nil || p.RestoreTimedOut || r.Recovery.TimedOut || !r.Recovery.ProcessRestarted || !r.Recovery.RedundancyRestored || !r.Recovery.CombinedRecovered || !r.Recovery.TopologyRecovered || !r.Recovery.ClientRecovered || r.Recovery.ReplicaPromoted != nil {
		t.Fatalf("restored service/original topology conflated with data verdict: %+v", r.Recovery)
	}
	if ok, why := cluster.ReplicaRecoveryWithBaseline(*r.Final, *r.Selection, *r.Initial); !ok {
		t.Fatalf("original six-node ownership/assignment not restored: %v", why)
	}
	if r.Dimensions == nil || r.Dimensions.DataVerified == nil || *r.Dimensions.DataVerified != r.Recovery.DataVerified || r.Dimensions.ReadAvailable == nil || !*r.Dimensions.ReadAvailable || r.Dimensions.WriteAvailable == nil || !*r.Dimensions.WriteAvailable || r.Dimensions.PersistenceRecovered == nil {
		t.Fatalf("independent measured dimensions missing: %+v", r.Dimensions)
	}
	if !r.Validation.Complete || r.Validation.Expected != int(p.AcknowledgedFrozen) || r.Validation.Checked != r.Validation.Expected || r.Validation.Corrupt+r.Validation.Unreadable+r.Validation.Unverified != 0 {
		t.Fatalf("frozen value coverage incomplete/corrupt: %+v", r.Validation)
	}
	t.Logf("mode=%s frozen_ack=%d ambiguous=%d missing=%d drain_first_kill_ms=%.3f drain_last_kill_ms=%.3f all_stopped_first_restart_ms=%.3f", p.Mode, p.AcknowledgedFrozen, p.AmbiguousFrozen, r.Validation.Missing, *p.DrainToFirstKillMS, *p.DrainToLastKillMS, *p.AllStoppedToFirstRestartMS)
}

func TestDockerPersistenceModes(t *testing.T) {
	for _, mode := range []string{"memory", "aof-everysec", "aof-always"} {
		t.Run(mode, func(t *testing.T) {
			r, err := runIntegration(t, persistenceIntegrationConfig(mode), environment.New())
			assertPersistenceRestartEvidence(t, r)
			p := r.PersistenceFault
			if mode == "memory" {
				if err == nil || r.Outcome != "partial" || r.Recovery.DataVerified || r.Validation.Missing != r.Validation.Expected || p.PersistenceRecoveryObserved != nil || p.PersistenceRecoveryAt != nil || *r.Dimensions.PersistenceRecovered {
					t.Fatal("memory-only value loss or absent persistence mechanism falsely became complete recovery")
				}
			} else if r.Validation.Missing > 0 {
				// A measured process crash can lose unsynced everysec writes. Preserve
				// the evidence and fail its data criterion without promising zero loss.
				if err == nil || r.Outcome != "partial" || r.Recovery.DataVerified || p.PersistenceRecoveryObserved == nil || *p.PersistenceRecoveryObserved || *r.Dimensions.PersistenceRecovered {
					t.Fatal("observed AOF loss was hidden by service recovery")
				}
			} else if err != nil || r.Outcome != "recovered" || !r.Recovery.DataVerified || p.PersistenceRecoveryObserved == nil || !*p.PersistenceRecoveryObserved || p.PersistenceRecoveryAt == nil || !*r.Dimensions.PersistenceRecovered {
				t.Fatalf("observed frozen AOF preservation did not meet the explicit recovery criterion: outcome=%s err=%v", r.Outcome, err)
			}
		})
	}
}

func TestDockerPersistenceRestartFailureCleansVolumes(t *testing.T) {
	injected := errors.New("injected ambiguous all-process restart response")
	m := &interruptedRestartManager{Manager: environment.New(), err: injected}
	r, err := runIntegration(t, persistenceIntegrationConfig("aof-always"), m)
	p := r.PersistenceFault
	if !errors.Is(err, injected) || r.Outcome != "failed" || p == nil || len(p.Stopped) != 6 || p.AllStoppedAt == nil || !m.started || len(p.Restarts) != 1 || !p.Restarts[0].After.ContainerRunning || len(p.RestartActions) != 1 || p.RestartActions[0].Error == "" || p.RestartActions[0].CompletedAt != nil || p.RestartCompletedAt != nil || r.Recovery.ProcessRestarted || p.StableRedundancy != nil || p.FrozenValidation != nil || p.DiagnosticsRequestedAt != nil || m.cleanupCalls != 1 || m.cleanupErr != nil || !m.cleanupDeadline {
		t.Fatalf("applied restart failure crossed data/recovery boundary or skipped fresh cleanup: persistence=%+v err=%v", p, err)
	}
}

func TestDockerPersistenceRestartTimeout(t *testing.T) {
	c := persistenceIntegrationConfig("aof-everysec")
	c.Timeouts.Repair = config.Duration(100 * time.Millisecond)
	r, err := runIntegration(t, c, environment.New())
	if err == nil || r.Outcome == "recovered" || r.PersistenceFault == nil || r.PersistenceFault.AllStoppedAt == nil || len(r.PersistenceFault.Stopped) != 6 || !r.PersistenceFault.RestoreTimedOut || r.PersistenceFault.StableRedundancy != nil || r.Recovery.CombinedRecovered || r.Recovery.ProcessRestarted || r.PersistenceFault.FrozenValidation != nil || r.PersistenceFault.DiagnosticsRequestedAt != nil {
		t.Fatalf("expired physical mutation budget became validated restart: persistence=%+v err=%v", r.PersistenceFault, err)
	}
}

func TestDockerPersistenceCancellationAfterRestart(t *testing.T) {
	bounded, boundCancel := context.WithTimeout(context.Background(), 4*time.Minute)
	defer boundCancel()
	ctx, cancel := context.WithCancel(bounded)
	defer cancel()
	m := &interruptedRestartManager{Manager: environment.New(), cancel: cancel}
	root, err := filepath.Abs("../../reports/ci")
	if err != nil {
		t.Fatal(err)
	}
	r, err := (Runner{Manager: m}).Run(ctx, persistenceIntegrationConfig("aof-always"), root)
	if r == nil {
		t.Fatalf("missing cancellation evidence: %v", err)
	}
	t.Cleanup(func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), time.Minute)
		defer cleanupCancel()
		if err := m.Manager.Cleanup(cleanupCtx, r.Environment); err != nil {
			t.Error(err)
		}
	})
	p := r.PersistenceFault
	if !errors.Is(err, context.Canceled) || r.Outcome != "aborted" || !m.started || !r.Cleanup.Successful || m.cleanupCalls != 1 || m.cleanupErr != nil || !m.cleanupDeadline || p == nil || p.AllStoppedAt == nil || len(p.Stopped) != 6 || len(p.Restarts) != 1 || p.StableRedundancy != nil || p.FrozenValidation != nil || p.DiagnosticsRequestedAt != nil || r.Recovery.CombinedRecovered || r.Recovery.ProcessRestarted {
		t.Fatalf("canceled applied restart fabricated validation or skipped independent cleanup: outcome=%s persistence=%+v err=%v", r.Outcome, p, err)
	}
	found, err := environment.New().Discover(bounded, r.RunID)
	if err != nil || len(found.Nodes) != 0 || found.NetworkID != "" || len(found.Volumes) != 0 {
		t.Fatalf("canceled retained restart leaked owned resources: inventory=%+v err=%v", found, err)
	}
	assertReportFiles(t, root, r.RunID, "aborted")
	t.Logf("run_id=%s outcome=aborted stopped_barrier=true first_restart_applied=true cleanup_verified=true", r.RunID)
}
