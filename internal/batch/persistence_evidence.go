package batch

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/cluster"
	"github.com/Alaneel/redis-reliability-lab/internal/environment"
	"github.com/Alaneel/redis-reliability-lab/internal/report"
)

// Aggregate recovery claims require the stopped barrier and frozen read-back
// evidence. A newly healthy topology alone cannot establish persistence.
func recoveredPersistenceEvidence(r *report.Report) error {
	if r.Config.Experiment.Fault != "cluster-restart" {
		return nil
	}
	p, env := r.PersistenceFault, r.Environment
	if r.SchemaVersion != report.LifecycleSchemaVersion || p == nil || env == nil || env.RunID != r.RunID || !env.RetainData || len(env.Nodes) != 6 || len(env.Volumes) != 6 || p.Mode != r.Config.Environment.Persistence || p.FailureModel != report.PersistenceFailureModel || p.RestoreTimedOut || r.Initial == nil || r.Final == nil || r.Selection == nil {
		return errors.New("recovered persistence trial lacks managed all-process restart evidence")
	}
	if !validTime(p.TrafficDrainedAt) || !validTime(p.LedgerFrozenAt) || !validTime(p.AllStoppedAt) || !validTime(p.RestartRequestedAt) || !validTime(p.RestartCompletedAt) || !validTime(p.DataCheckedAt) || !validTime(p.DiagnosticsRequestedAt) || p.LedgerFrozenAt.Before(*p.TrafficDrainedAt) || p.RestartRequestedAt.Before(*p.AllStoppedAt) || p.RestartCompletedAt.Before(*p.RestartRequestedAt) || p.DataCheckedAt.Before(*p.RestartCompletedAt) || p.DiagnosticsRequestedAt.Before(*p.DataCheckedAt) {
		return errors.New("recovered persistence trial lacks ordered drain, stopped barrier, validation and diagnostics")
	}
	if p.PreCrashSnapshot == nil || p.PreCrashSnapshot.At.IsZero() || p.PreCrashSnapshot.Completed.IsZero() || !reflect.DeepEqual(*p.PreCrashSnapshot, *r.Initial) || p.PreCrashSnapshot.Completed.After(*p.LedgerFrozenAt) || p.AcknowledgedFrozen != r.WorkloadStats.AcknowledgedWrites || p.AmbiguousFrozen != r.WorkloadStats.AmbiguousWrites || p.FrozenValidation == nil || !reflect.DeepEqual(*p.FrozenValidation, r.Validation) || uint64(r.Validation.Expected) != p.AcknowledgedFrozen || r.Validation.StartedAt.Before(*p.RestartCompletedAt) || r.Validation.CompletedAt.After(*p.DataCheckedAt) {
		return errors.New("recovered persistence trial frozen ledger or validation differs from final accounting")
	}
	if r.FinishedAt.IsZero() || r.Final.At.Before(*p.RestartCompletedAt) || r.Final.Completed.Before(r.Final.At) || r.Final.Completed.After(r.FinishedAt) || p.DiagnosticsRequestedAt.After(r.FinishedAt) || p.DiagnosticServiceStable == nil || p.DiagnosticServiceStable.At.After(r.FinishedAt) {
		return errors.New("recovered persistence final observation is stale or outside the run")
	}
	for _, view := range r.Final.Views {
		if view.At.Before(*p.RestartCompletedAt) || view.At.After(r.Final.Completed) {
			return errors.New("persistence final node observation predates restart or postdates the snapshot")
		}
	}
	if !nonnegativeObservation(p.DrainToFirstKillMS) || !nonnegativeObservation(p.DrainToLastKillMS) || !nonnegativeObservation(p.AllStoppedToFirstRestartMS) || !observedMilestone(p.TopologyStable) || !observedMilestone(p.StableRedundancy) || !observedMilestone(p.DiagnosticServiceStable) || !r.Recovery.RedundancyRestored {
		return errors.New("recovered persistence trial lacks sustained restoration observations")
	}
	if p.TopologyStable.At.Before(*p.RestartCompletedAt) || p.StableRedundancy.At.Before(*p.RestartCompletedAt) || p.DiagnosticServiceStable.At.Before(*p.DiagnosticsRequestedAt) || !observedMilestone(p.PersistenceRecoveryAt) || p.PersistenceRecoveryAt.At.Before(*p.DataCheckedAt) {
		return errors.New("persistence recovery milestones precede their required evidence")
	}
	if !reflect.DeepEqual(r.Recovery.CombinedStable, p.DiagnosticServiceStable) {
		return errors.New("persistence combined recovery differs from post-validation diagnostic evidence")
	}
	for _, milestone := range []*report.Milestone{p.TopologyStable, p.StableRedundancy, p.PersistenceRecoveryAt, p.DiagnosticServiceStable, r.Recovery.ClientFirstSuccess, r.Recovery.ClientStable} {
		if !observedMilestone(milestone) || milestone.At.After(r.FinishedAt) || !nonnegativeObservation(r.FaultActionDurationMS) || milestone.SinceFaultMS < *r.FaultActionDurationMS {
			return errors.New("persistence milestone lacks observed post-crash timing within the run")
		}
	}
	if r.Recovery.ClientFirstSuccess.At.Before(*p.DiagnosticsRequestedAt) || r.Recovery.ClientStable.At.Before(*p.DiagnosticsRequestedAt) || r.WorkloadStats.PacingMethodology == "" || r.Metrics.Phases["baseline"].Attempts == 0 {
		return errors.New("persistence trial lacks measured baseline or fresh diagnostic client observations")
	}
	for phase, summary := range r.Metrics.Phases {
		if phase != "baseline" && summary.Attempts != 0 {
			return errors.New("persistence trial offered workload requests after its frozen baseline")
		}
	}
	if ok, _ := cluster.ReplicaRecoveryWithBaseline(*r.Final, *r.Selection, *r.Initial); !ok {
		return errors.New("recovered persistence trial changed original topology or lacks synchronized replicas")
	}
	if len(p.Prepared) != 6 || len(p.Stopped) != 6 || len(p.KillActions) != 6 || len(p.Restarts) != 6 || len(p.RestartActions) != 6 || len(p.Restored) != 6 {
		return errors.New("recovered persistence trial does not account for all six processes")
	}
	prepared := map[string]report.NodeProcessSample{}
	stopped := map[string]report.NodeProcessSample{}
	owned := map[string]environment.Node{}
	volumes := map[string]bool{}
	for _, volume := range env.Volumes {
		if volume.Name == "" || volume.NodeID == "" || volumes[volume.NodeID] {
			return errors.New("invalid retained volume inventory")
		}
		volumes[volume.NodeID] = true
	}
	for _, node := range env.Nodes {
		if node.ID == "" || node.RedisClusterID == "" || owned[node.ID].ID != "" || !volumes[node.ID] {
			return errors.New("invalid original persistence node inventory")
		}
		matches := 0
		for _, view := range r.Initial.Views {
			if view.Node.ID == node.ID && view.Node.Addr == node.Addr && view.Node.InternalAddr == node.InternalAddr && view.RedisID == node.RedisClusterID {
				matches++
			}
		}
		if matches != 1 {
			return errors.New("persistence inventory differs from original observed identities or endpoints")
		}
		owned[node.ID] = node
	}
	for _, sample := range p.Prepared {
		if sample.Node != owned[sample.Node.ID] || sample.OriginalRedisID != sample.Node.RedisClusterID || prepared[sample.Node.ID].Node.ID != "" || sample.Error != "" || !liveProcess(sample.Evidence, sample.OriginalRedisID) || sample.Evidence.ObservedAt.After(*p.TrafficDrainedAt) || !persistenceRuntime(sample.Evidence, p.Mode) {
			return errors.New("invalid pre-crash persistence process evidence")
		}
		prepared[sample.Node.ID] = sample
	}
	for _, sample := range p.Stopped {
		before := prepared[sample.Node.ID]
		if sample.Node != owned[sample.Node.ID] || sample.OriginalRedisID != sample.Node.RedisClusterID || stopped[sample.Node.ID].Node.ID != "" || sample.Error != environment.ErrProcessStopped.Error() || !stoppedProcess(sample.Evidence) || sample.Evidence.ObservedAt.After(*p.AllStoppedAt) || sample.Evidence.ObservedAt.Before(*p.LedgerFrozenAt) || sample.Evidence.ContainerStartedAt != before.Evidence.ContainerStartedAt {
			return errors.New("invalid stopped process barrier evidence")
		}
		stopped[sample.Node.ID] = sample
	}
	seen := map[string]bool{}
	var lastKill time.Time
	roles := map[string]string{}
	for _, view := range r.Initial.Views {
		roles[view.Node.ID] = view.Replication["role"]
	}
	for i, action := range p.KillActions {
		if action.Node != owned[action.Node.ID] || seen[action.Node.ID] || action.RequestedAt.IsZero() || action.RequestedAt.Before(*p.LedgerFrozenAt) || !validTime(action.CompletedAt) || action.CompletedAt.Before(action.RequestedAt) || action.CompletedAt.After(*p.AllStoppedAt) || action.Error != "" {
			return errors.New("invalid six-process kill action evidence")
		}
		expectedRole := "slave"
		if i >= 3 {
			expectedRole = "master"
		}
		if roles[action.Node.ID] != expectedRole || action.RequestedAt.Before(lastKill) {
			return errors.New("persistence kills did not stop original replicas before primaries sequentially")
		}
		lastKill = *action.CompletedAt
		seen[action.Node.ID] = true
	}
	for _, sample := range p.Stopped {
		if sample.Evidence.ObservedAt.Before(lastKill) {
			return errors.New("stopped barrier sample predates the last kill")
		}
	}
	seen = map[string]bool{}
	afterRestart := map[string]environment.ProcessEvidence{}
	var lastRestart time.Time
	for i, x := range p.Restarts {
		if x == nil || seen[x.NodeID] || x.Node != owned[x.NodeID] || x.NodeID != x.Node.ID || x.RunID != r.RunID || x.OriginalRedisID != x.Node.RedisClusterID || x.RequestedAt.Before(*p.AllStoppedAt) || x.CompletedAt.IsZero() || x.CompletedAt.Before(x.RequestedAt) || x.CompletedAt.After(*p.RestartCompletedAt) || !stoppedProcess(x.Before) || !liveProcess(x.After, x.OriginalRedisID) || x.Before.ContainerStartedAt != prepared[x.NodeID].Evidence.ContainerStartedAt || x.After.ContainerStartedAt == x.Before.ContainerStartedAt || x.After.RedisRunID == prepared[x.NodeID].Evidence.RedisRunID {
			return errors.New("invalid persistence restart receipt or runtime configuration")
		}
		action := p.RestartActions[i]
		if action.Node != x.Node || action.RequestedAt.IsZero() || action.RequestedAt.Before(*p.AllStoppedAt) || !validTime(action.CompletedAt) || action.CompletedAt.Before(action.RequestedAt) || action.CompletedAt.After(*p.RestartCompletedAt) || action.Error != "" {
			return errors.New("invalid persistence restart action order")
		}
		expectedRole := "master"
		if i >= 3 {
			expectedRole = "slave"
		}
		if roles[x.NodeID] != expectedRole || action.RequestedAt.Before(lastRestart) || x.RequestedAt.Before(action.RequestedAt) || x.Before.ObservedAt.Before(x.RequestedAt) || x.After.ObservedAt.Before(x.Before.ObservedAt) || x.After.ObservedAt.After(x.CompletedAt) || x.CompletedAt.After(*action.CompletedAt) {
			return errors.New("persistence restart order or receipt observation timing is contradictory")
		}
		lastRestart = *action.CompletedAt
		seen[x.NodeID] = true
		afterRestart[x.NodeID] = x.After
	}
	seen = map[string]bool{}
	for _, sample := range p.Restored {
		after := afterRestart[sample.Node.ID]
		if sample.Node != owned[sample.Node.ID] || seen[sample.Node.ID] || sample.Error != "" || sample.OriginalRedisID != sample.Node.RedisClusterID || !liveProcess(sample.Evidence, sample.OriginalRedisID) || sample.Evidence.ObservedAt.Before(p.TopologyStable.At) || sample.Evidence.ObservedAt.After(r.Validation.StartedAt) || sample.Evidence.RedisRunID != after.RedisRunID || sample.Evidence.ContainerStartedAt != after.ContainerStartedAt || !persistenceRuntime(sample.Evidence, p.Mode) {
			return errors.New("persistence recovery lacks fresh matching post-synchronization runtime evidence")
		}
		seen[sample.Node.ID] = true
	}
	if p.Mode == "memory" || p.PersistenceRecoveryObserved == nil || !*p.PersistenceRecoveryObserved || !observedMilestone(p.PersistenceRecoveryAt) || r.Dimensions == nil || r.Dimensions.PersistenceRecovered == nil || !*r.Dimensions.PersistenceRecovered || r.Dimensions.DataVerified == nil || !*r.Dimensions.DataVerified {
		return errors.New("recovered persistence outcome lacks observed AOF frozen-value preservation")
	}
	return nil
}

func validTime(at *time.Time) bool { return at != nil && !at.IsZero() }

func stoppedProcess(e environment.ProcessEvidence) bool {
	return !e.ObservedAt.IsZero() && !e.ContainerRunning && e.ContainerPID == 0 && !e.RedisPONG && !e.NetworkConnected && e.ContainerStartedAt != "" && e.ContainerRestartPolicy == "no"
}

func liveProcess(e environment.ProcessEvidence, redisID string) bool {
	return !e.ObservedAt.IsZero() && e.ContainerRunning && e.ContainerPID > 0 && e.RedisPONG && e.RedisPID > 0 && e.RedisRunID != "" && e.RedisClusterID == redisID && e.NetworkConnected && e.NetworkCount == 1 && e.ContainerStartedAt != "" && e.ContainerRestartPolicy == "no"
}

func persistenceRuntime(e environment.ProcessEvidence, mode string) bool {
	lines := strings.Split(strings.ReplaceAll(e.RawPersistenceConfig, "\r\n", "\n"), "\n")
	if len(lines)%2 != 0 {
		return false
	}
	values := map[string]string{}
	for i := 0; i < len(lines); i += 2 {
		if _, duplicate := values[lines[i]]; duplicate || lines[i] == "" {
			return false
		}
		values[lines[i]] = lines[i+1]
	}
	info := cluster.ParseInfo(e.RawPersistence)
	preamble := "no"
	if mode != "memory" {
		preamble = "yes"
	}
	expected := map[string]string{"save": "", "repl-diskless-load": "swapdb", "repl-diskless-sync": "yes", "auto-aof-rewrite-percentage": "0", "maxmemory": "134217728", "maxmemory-policy": "noeviction", "aof-use-rdb-preamble": preamble, "no-appendfsync-on-rewrite": "no", "aof-load-truncated": "yes", "rdb-del-sync-files": "yes"}
	if len(values) != len(expected)+2 {
		return false
	}
	for key, want := range expected {
		if got, exists := values[key]; !exists || got != want {
			return false
		}
	}
	if info["loading"] != "0" || !strings.Contains(e.RawDataFiles, "/data/nodes.conf") {
		return false
	}
	if mode == "memory" {
		return values["appendonly"] == "no" && info["aof_enabled"] == "0" && !strings.Contains(e.RawDataFiles, ".rdb") && !strings.Contains(e.RawDataFiles, ".aof")
	}
	return values["appendonly"] == "yes" && values["appendfsync"] == strings.TrimPrefix(mode, "aof-") && info["aof_enabled"] == "1" && info["aof_last_write_status"] == "ok" && info["aof_last_bgrewrite_status"] == "ok" && strings.Contains(e.RawDataFiles, "/data/appendonlydir/") && strings.Contains(e.RawDataFiles, ".aof") && strings.Contains(e.RawDataFiles, ".manifest")
}
