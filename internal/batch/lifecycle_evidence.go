package batch

import (
	"errors"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/report"
)

// Recovered lifecycle outcomes need a concrete repaired node and a fresh full
// topology predicate; a historical service milestone alone is insufficient.
func recoveredLifecycleEvidence(r *report.Report) error {
	replica := r.Config.Experiment.Fault == "replica-kill" || r.Config.Experiment.Fault == "replica-isolate"
	if !replica && !r.Config.Experiment.Repair {
		return nil
	}
	l := r.LifecycleFault
	if r.SchemaVersion != report.LifecycleSchemaVersion || l == nil || !l.RepairCompleted || l.RepairTimedOut || !r.Recovery.RedundancyRestored || !observedMilestone(l.StableRedundancy) || !nonnegativeObservation(l.StableRedundancySinceRepairMS) || l.RepairRequestedAt == nil || l.RepairRequestedAt.IsZero() || l.RepairCompletedAt == nil || l.RepairCompletedAt.IsZero() || !nonnegativeObservation(l.RepairActionDurationMS) || r.Initial == nil || r.Final == nil || r.Selection == nil {
		return errors.New("recovered lifecycle trial lacks completed repair and sustained redundancy evidence")
	}
	expectedNode, expectedID := r.Selection.Target, r.Selection.TargetID
	if replica {
		expectedNode, expectedID = r.Selection.Replica, r.Selection.ReplicaID
	}
	if l.Target != expectedNode || l.TargetRedisID != expectedID {
		return errors.New("recovered lifecycle target differs from its baseline selection")
	}
	if r.Config.Experiment.Fault == "replica-isolate" {
		if l.RepairKind != "reconnect" || !recoveredNetworkEvidence(r) {
			return errors.New("recovered replica isolation lacks unchanged process and reconnection evidence")
		}
	} else {
		x := l.RestartReceipt
		ownedNodeMatches := false
		if x != nil && r.Environment != nil {
			for _, node := range r.Environment.Nodes {
				if node.ID == expectedNode.ID && node.Addr == expectedNode.Addr && node.InternalAddr == expectedNode.InternalAddr && x.Node == node {
					ownedNodeMatches = true
				}
			}
		}
		if l.RepairKind != "restart" || x == nil || r.Environment == nil || !ownedNodeMatches || x.RunID != r.RunID || x.RunID != r.Environment.RunID || x.NodeID != expectedNode.ID || x.OriginalRedisID != expectedID || x.CompletedAt.IsZero() || x.Before.ContainerRunning || !x.After.ContainerRunning || !x.After.RedisPONG || !x.After.NetworkConnected || x.After.RedisClusterID != expectedID || x.After.ContainerStartedAt == "" || x.After.ContainerStartedAt == x.Before.ContainerStartedAt {
			return errors.New("recovered lifecycle lacks an acknowledged same-node restart with preserved membership identity")
		}
	}
	var ok bool
	if replica {
		ok, _ = cluster.ReplicaRecoveryWithBaseline(*r.Final, *r.Selection, *r.Initial)
	} else {
		ok, _ = cluster.ReconciledWithBaseline(*r.Final, *r.Selection, *r.Initial)
	}
	if !ok {
		return errors.New("recovered lifecycle trial lacks final baseline-relative topology reconciliation")
	}
	if r.Config.RetainData() && (r.Environment == nil || !r.Environment.RetainData || len(r.Environment.Volumes) != 6) {
		return errors.New("recovered restart trial lacks disposable retained-resource inventory")
	}
	return nil
}
