package batch

import (
	"errors"
	"reflect"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/report"
)

// A recovered WAITAOF claim additionally requires the original managed AOF
// runtime and same-policy all-shard command compatibility before traffic begins.
func recoveredPolicyRuntimeEvidence(r *report.Report) error {
	if r.Config.Workload.WritePolicy != "waitaof" {
		return nil
	}
	env, mode := r.Environment, r.Config.Environment.Persistence
	if env == nil || env.RunID != r.RunID || !env.RetainData || env.Persistence != mode || mode != "aof-everysec" && mode != "aof-always" || len(env.Nodes) != 6 || len(env.Volumes) != 6 || r.Initial == nil || r.StartedAt.IsZero() {
		return errors.New("recovered WAITAOF trial lacks original managed AOF runtime")
	}
	var baseline time.Time
	for _, event := range r.Events {
		if event.Stage == "baseline" {
			if !baseline.IsZero() || event.At.IsZero() || event.At.Before(r.StartedAt) {
				return errors.New("WAITAOF baseline chronology is unavailable or contradictory")
			}
			baseline = event.At
		}
	}
	if baseline.IsZero() || r.FinishedAt.IsZero() || !validTime(r.FaultRequestedAt) || r.FaultRequestedAt.Before(r.StartedAt) || r.FaultRequestedAt.After(r.FinishedAt) || !baseline.Before(*r.FaultRequestedAt) || baseline.After(r.FinishedAt) {
		return errors.New("WAITAOF baseline and fault chronology is unavailable or outside the run")
	}
	prepared := r.PersistencePrepared
	if r.PersistenceFault != nil {
		if len(prepared) > 0 && !reflect.DeepEqual(prepared, r.PersistenceFault.Prepared) {
			return errors.New("WAITAOF preparation differs between runtime and persistence evidence")
		}
		prepared = r.PersistenceFault.Prepared
	}
	if len(prepared) != 6 {
		return errors.New("recovered WAITAOF trial lacks six before-baseline runtime samples")
	}
	owned := map[string]environment.Node{}
	volumes := map[string]bool{}
	for _, volume := range env.Volumes {
		if volume.Name == "" || volume.NodeID == "" || volumes[volume.NodeID] {
			return errors.New("WAITAOF retained volume inventory is invalid")
		}
		volumes[volume.NodeID] = true
	}
	for _, node := range env.Nodes {
		if node.ID == "" || node.ContainerID == "" || node.RedisClusterID == "" || owned[node.ID].ID != "" || !volumes[node.ID] {
			return errors.New("WAITAOF original managed node inventory is invalid")
		}
		matches := 0
		for _, view := range r.Initial.Views {
			if view.Node.ID == node.ID && view.Node.Addr == node.Addr && view.Node.InternalAddr == node.InternalAddr && view.RedisID == node.RedisClusterID {
				matches++
			}
		}
		if matches != 1 {
			return errors.New("WAITAOF prepared inventory differs from direct initial node observations")
		}
		owned[node.ID] = node
	}
	seen := map[string]bool{}
	var latestPrepared time.Time
	for _, sample := range prepared {
		e := sample.Evidence
		if sample.Node != owned[sample.Node.ID] || sample.OriginalRedisID != sample.Node.RedisClusterID || sample.Error != "" || seen[sample.Node.ID] || !liveProcess(e, sample.OriginalRedisID) || e.ObservedAt.Before(r.StartedAt) || !e.ObservedAt.Before(baseline) || !persistenceRuntime(e, mode) || !quietAOFPreparation(e) {
			return errors.New("WAITAOF before-baseline sample lacks matching quiet AOF runtime")
		}
		seen[sample.Node.ID] = true
		if e.ObservedAt.After(latestPrepared) {
			latestPrepared = e.ObservedAt
		}
	}
	p := r.BaselineWritePolicyProbe
	if p == nil || p.Policy != "waitaof" || !p.Successful || len(p.Errors) != 0 || !p.At.After(latestPrepared) || !p.At.Before(baseline) || p.WaitAOFRequestedLocal != 1 || p.WaitAOFRequestedReplicas != r.Config.Workload.WaitReplicas || len(p.SetAcknowledged) != 3 || len(p.WaitAOFAttempted) != 3 || len(p.WaitAOFAcknowledgementObserved) != 3 || len(p.WaitAOFLocalAcknowledgements) != 3 || len(p.WaitAOFReplicaAcknowledgements) != 3 || len(p.WaitAttempted) != 0 || len(p.Acknowledgements) != 0 || len(p.AcknowledgementObserved) != 0 {
		return errors.New("WAITAOF before-baseline all-shard command compatibility is unavailable")
	}
	for i := range 3 {
		if !p.SetAcknowledged[i] || !p.WaitAOFAttempted[i] || !p.WaitAOFAcknowledgementObserved[i] || p.WaitAOFLocalAcknowledgements[i] != 1 || p.WaitAOFReplicaAcknowledgements[i] < int64(r.Config.Workload.WaitReplicas) {
			return errors.New("WAITAOF before-baseline all-shard acknowledgement requirements were not observed")
		}
	}
	return nil
}

func quietAOFPreparation(e environment.ProcessEvidence) bool {
	p := cluster.ParseInfo(e.RawPersistence)
	for _, key := range []string{"loading", "rdb_bgsave_in_progress", "aof_rewrite_in_progress", "aof_rewrite_scheduled", "aof_pending_bio_fsync", "aof_pending_rewrite"} {
		if p[key] != "0" {
			return false
		}
	}
	repl := cluster.ParseInfo(e.RawReplication)
	switch repl["role"] {
	case "master":
		return repl["connected_slaves"] == "1"
	case "slave":
		return repl["master_link_status"] == "up" && repl["master_sync_in_progress"] == "0"
	default:
		return false
	}
}
