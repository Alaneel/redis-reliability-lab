package batch

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/config"
	"github.com/Alaneel/redis-reliability-lab/internal/report"
)

func waitAOFRuntimeFixture(t *testing.T, persistence bool) *report.Report {
	t.Helper()
	r := persistenceFixture(t)
	r.Config.Workload.WritePolicy = "waitaof"
	base := r.StartedAt
	r.Events = []report.Event{{Stage: "baseline", At: base.Add(time.Second)}}
	for i := range r.PersistenceFault.Prepared {
		e := &r.PersistenceFault.Prepared[i].Evidence
		e.RawPersistence += "rdb_bgsave_in_progress:0\naof_rewrite_in_progress:0\naof_rewrite_scheduled:0\naof_pending_bio_fsync:0\naof_pending_rewrite:0\n"
		e.RawReplication = "role:master\nconnected_slaves:1\n"
		if i >= 3 {
			e.RawReplication = "role:slave\nmaster_link_status:up\nmaster_sync_in_progress:0\n"
		}
	}
	if !persistence {
		r.Config.Experiment.Fault = "primary-kill"
		r.PersistencePrepared = r.PersistenceFault.Prepared
		r.PersistenceFault = nil
	}
	r.BaselineWritePolicyProbe = &report.PolicyProbe{Policy: "waitaof", At: base.Add(500 * time.Millisecond), Successful: true, WaitAOFRequestedLocal: 1, WaitAOFRequestedReplicas: 1,
		SetAcknowledged: []bool{true, true, true}, WaitAOFAttempted: []bool{true, true, true}, WaitAOFAcknowledgementObserved: []bool{true, true, true},
		WaitAOFLocalAcknowledgements: []int64{1, 1, 1}, WaitAOFReplicaAcknowledgements: []int64{1, 1, 1}}
	return r
}

func TestRecoveredWAITAOFRequiresRuntimeAndBeforeBaselineCommandEvidence(t *testing.T) {
	for _, persistence := range []bool{false, true} {
		r := waitAOFRuntimeFixture(t, persistence)
		if err := recoveredPolicyRuntimeEvidence(r); err != nil {
			t.Fatalf("valid recovered runtime rejected (persistence=%t): %v", persistence, err)
		}
	}
	for _, test := range []struct {
		name   string
		change func(*report.Report)
	}{
		{"no_environment", func(r *report.Report) { r.Environment = nil }},
		{"not_retained", func(r *report.Report) { r.Environment.RetainData = false }},
		{"wrong_aof_mode", func(r *report.Report) { r.Environment.Persistence = "aof-everysec" }},
		{"missing_original_views", func(r *report.Report) { r.Initial.Views = nil }},
		{"foreign_container", func(r *report.Report) { r.PersistencePrepared[0].Node.ContainerID = "foreign" }},
		{"duplicate_node", func(r *report.Report) { r.PersistencePrepared[1] = r.PersistencePrepared[0] }},
		{"missing_node", func(r *report.Report) { r.PersistencePrepared = r.PersistencePrepared[:5] }},
		{"wrong_identity", func(r *report.Report) { r.PersistencePrepared[0].Evidence.RedisClusterID = "foreign" }},
		{"wrong_restart_policy", func(r *report.Report) { r.PersistencePrepared[0].Evidence.ContainerRestartPolicy = "always" }},
		{"runtime_not_aof", func(r *report.Report) {
			r.PersistencePrepared[0].Evidence.RawPersistenceConfig = strings.ReplaceAll(r.PersistencePrepared[0].Evidence.RawPersistenceConfig, "appendonly\nyes", "appendonly\nno")
		}},
		{"fsync_suppressed", func(r *report.Report) {
			r.PersistencePrepared[0].Evidence.RawPersistenceConfig = strings.ReplaceAll(r.PersistencePrepared[0].Evidence.RawPersistenceConfig, "no-appendfsync-on-rewrite\nno", "no-appendfsync-on-rewrite\nyes")
		}},
		{"pending_fsync", func(r *report.Report) {
			r.PersistencePrepared[0].Evidence.RawPersistence = strings.ReplaceAll(r.PersistencePrepared[0].Evidence.RawPersistence, "aof_pending_bio_fsync:0", "aof_pending_bio_fsync:1")
		}},
		{"rewrite_active", func(r *report.Report) {
			r.PersistencePrepared[0].Evidence.RawPersistence = strings.ReplaceAll(r.PersistencePrepared[0].Evidence.RawPersistence, "aof_rewrite_in_progress:0", "aof_rewrite_in_progress:1")
		}},
		{"replica_sync_active", func(r *report.Report) {
			r.PersistencePrepared[3].Evidence.RawReplication = strings.ReplaceAll(r.PersistencePrepared[3].Evidence.RawReplication, "master_sync_in_progress:0", "master_sync_in_progress:1")
		}},
		{"stale_before_run", func(r *report.Report) { r.PersistencePrepared[0].Evidence.ObservedAt = r.StartedAt.Add(-time.Second) }},
		{"after_baseline", func(r *report.Report) { r.PersistencePrepared[0].Evidence.ObservedAt = r.Events[0].At.Add(time.Second) }},
		{"no_baseline_anchor", func(r *report.Report) { r.Events = nil }},
		{"baseline_after_fault", func(r *report.Report) { r.Events[0].At = r.FaultRequestedAt.Add(time.Second) }},
		{"baseline_after_finish", func(r *report.Report) { r.Events[0].At = r.FinishedAt.Add(time.Second) }},
		{"fault_after_finish", func(r *report.Report) { at := r.FinishedAt.Add(time.Second); r.FaultRequestedAt = &at }},
		{"fault_before_run", func(r *report.Report) { at := r.StartedAt.Add(-time.Second); r.FaultRequestedAt = &at }},
		{"missing_fault", func(r *report.Report) { r.FaultRequestedAt = nil }},
		{"no_command_probe", func(r *report.Report) { r.BaselineWritePolicyProbe = nil }},
		{"probe_before_runtime", func(r *report.Report) { r.BaselineWritePolicyProbe.At = r.StartedAt }},
		{"probe_after_baseline", func(r *report.Report) { r.BaselineWritePolicyProbe.At = r.Events[0].At.Add(time.Second) }},
		{"probe_failed", func(r *report.Report) { r.BaselineWritePolicyProbe.Successful = false }},
		{"probe_error", func(r *report.Report) { r.BaselineWritePolicyProbe.Errors = []string{"timeout"} }},
		{"probe_wrong_policy", func(r *report.Report) { r.BaselineWritePolicyProbe.Policy = "wait" }},
		{"probe_wrong_local_request", func(r *report.Report) { r.BaselineWritePolicyProbe.WaitAOFRequestedLocal = 0 }},
		{"probe_wrong_replica_request", func(r *report.Report) { r.BaselineWritePolicyProbe.WaitAOFRequestedReplicas = 2 }},
		{"probe_missing_shard", func(r *report.Report) { r.BaselineWritePolicyProbe.WaitAOFAttempted = []bool{true, true} }},
		{"probe_unacknowledged_set", func(r *report.Report) { r.BaselineWritePolicyProbe.SetAcknowledged[0] = false }},
		{"probe_unattempted", func(r *report.Report) { r.BaselineWritePolicyProbe.WaitAOFAttempted[0] = false }},
		{"probe_unknown", func(r *report.Report) { r.BaselineWritePolicyProbe.WaitAOFAcknowledgementObserved[0] = false }},
		{"probe_local_short", func(r *report.Report) { r.BaselineWritePolicyProbe.WaitAOFLocalAcknowledgements[0] = 0 }},
		{"probe_replica_short", func(r *report.Report) { r.BaselineWritePolicyProbe.WaitAOFReplicaAcknowledgements[0] = 0 }},
		{"probe_malformed", func(r *report.Report) { r.BaselineWritePolicyProbe.WaitAOFLocalAcknowledgements[0] = 2 }},
		{"probe_cross_policy", func(r *report.Report) { r.BaselineWritePolicyProbe.WaitAttempted = []bool{true, true, true} }},
	} {
		t.Run(test.name, func(t *testing.T) {
			r := waitAOFRuntimeFixture(t, false)
			test.change(r)
			if err := recoveredPolicyRuntimeEvidence(r); err == nil {
				t.Fatal("recovered WAITAOF accepted without managed compatible AOF runtime")
			}
		})
	}
}

func TestWAITAOFRecoveredWithoutRuntimeCannotEnterBatch(t *testing.T) {
	s, err := Run(context.Background(), waitAOFConfig(), 1, t.TempDir(), func(_ context.Context, c config.Config, root string) (*report.Report, error) {
		r := waitAOFFixture(c, 1)
		r.Outcome = "recovered"
		ok := true
		r.Recovery.WritePolicyRecovered = &ok
		r.Recovery.WritePolicyStable = r.Recovery.CombinedStable
		save(t, root, r)
		return r, nil
	})
	if err == nil || s.Outcomes["failed"] != 1 || s.IncompleteEvidenceTrials != 1 || s.WritePolicy.MeasuredTrials != 0 {
		t.Fatalf("unproven recovered policy contaminated batch: %+v %v", s, err)
	}
}
