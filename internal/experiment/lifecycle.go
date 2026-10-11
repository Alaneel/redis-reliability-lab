package experiment

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/report"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

func lifecycleRestored(s cluster.Snapshot, selected cluster.Selection, baseline cluster.Snapshot, replica bool) (bool, []string) {
	if replica {
		return cluster.ReplicaReconciledWithBaseline(s, selected, baseline)
	}
	return cluster.ReconciledWithBaseline(s, selected, baseline)
}

// Repair uses one deadline covering both the owned action and synchronization.
// No CLUSTER REPLICATE or ownership mutation is issued: accepted observations
// must prove Redis reconciled its current membership and intended assignment.
func repairLifecycle(ctx context.Context, c config.Config, r *report.Report, restart RestartManager, network NetworkFaultManager, target environment.Node, observer *cluster.Observer, client workload.Client, writer workload.WriteExecutor, keys []string, selected cluster.Selection, initial cluster.Snapshot, combined, service, policy *Stability, beginSynchronization func()) error {
	repairCtx, cancel := context.WithTimeout(ctx, c.Timeouts.Repair.Value())
	defer cancel()
	if err := performLifecycleRepair(repairCtx, r, restart, network, target); err != nil {
		return err
	}
	beginSynchronization()
	collect := func(sampleCtx context.Context) report.Sample {
		s := sampleRecovery(sampleCtx, observer, client, keys, r.RunID, c.Workload.OperationTimeout.Value(), selected, initial)
		if requiresAcknowledgementPolicy(c.Workload.WritePolicy) {
			p := ProbeAcknowledgementPolicy(sampleCtx, writer, client, keys, r.RunID, c.Workload.WritePolicy, c.Workload.WaitReplicas, c.Workload.OperationTimeout.Value())
			s.WritePolicyProbe = &p
		}
		return s
	}
	observeLifecycleSynchronization(repairCtx, c, r, selected, initial, combined, service, policy, collect)
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return nil
}

func performLifecycleRepair(ctx context.Context, r *report.Report, restart RestartManager, network NetworkFaultManager, target environment.Node) error {
	l := r.LifecycleFault
	at := time.Now()
	l.RepairRequestedAt = &at
	var err error
	if r.NetworkFault != nil {
		l.RepairKind = "reconnect"
		err = restoreNetwork(ctx, network, r)
	} else {
		l.RepairKind = "restart"
		l.RestartReceipt, err = restart.Restart(ctx, r.Environment, target)
		if x := l.RestartReceipt; x != nil && x.RunID == r.RunID && x.NodeID == target.ID && x.Node == target && x.OriginalRedisID == l.TargetRedisID {
			for _, p := range []environment.ProcessEvidence{x.Before, x.After} {
				if !p.ObservedAt.IsZero() {
					l.ProcessSamples = append(l.ProcessSamples, report.ProcessSample{Evidence: p})
				}
			}
			if !x.After.ObservedAt.IsZero() && x.After.RedisClusterID == l.TargetRedisID {
				dimensions := report.RecoveryDimensions{}
				if l.Dimensions != nil {
					dimensions = *l.Dimensions
				}
				if x.After.ContainerRunning && x.After.RedisPONG {
					live := true
					dimensions.TargetProcessLive = &live
				} else if !x.After.ContainerRunning {
					live := false
					dimensions.TargetProcessLive = &live
				} else {
					dimensions.TargetProcessLive = nil
				}
				dimensions.TargetRole = cluster.ParseInfo(x.After.RawReplication)["role"]
				l.Dimensions, r.Dimensions = &dimensions, &dimensions
			}
		}
		if err == nil && !validRestartReceipt(l.RestartReceipt, r.RunID, target, l.TargetRedisID) {
			err = errors.New("restart returned incomplete or contradictory owned-node identity evidence")
		}
	}
	if err != nil {
		l.RepairTimedOut = errors.Is(err, context.DeadlineExceeded) || errors.Is(ctx.Err(), context.DeadlineExceeded)
		r.Recovery.CombinedRecovered, r.Recovery.RedundancyRestored = false, false
		l.Errors = append(l.Errors, err.Error())
		return fmt.Errorf("repair %s: %w", target.ID, err)
	}
	completed := time.Now()
	duration := float64(completed.Sub(at)) / float64(time.Millisecond)
	l.RepairCompletedAt, l.RepairActionDurationMS, l.RepairCompleted = &completed, &duration, true
	if r.NetworkFault == nil {
		r.Recovery.ProcessRestarted = true
	}
	return nil
}

func validRestartReceipt(x *environment.RestartReceipt, runID string, target environment.Node, redisID string) bool {
	if x == nil || x.RunID != runID || x.NodeID != target.ID || x.Node != target || redisID == "" || target.RedisClusterID != redisID || x.OriginalRedisID != redisID || x.RequestedAt.IsZero() || x.CompletedAt.IsZero() || x.CompletedAt.Before(x.RequestedAt) {
		return false
	}
	b, a := x.Before, x.After
	return !b.ContainerRunning && b.ContainerStartedAt != "" && !b.ObservedAt.IsZero() &&
		!a.ObservedAt.IsZero() && !a.ObservedAt.Before(b.ObservedAt) && !x.CompletedAt.Before(a.ObservedAt) &&
		a.ContainerRunning && a.ContainerPID > 0 && a.ContainerStartedAt != "" && a.ContainerStartedAt != b.ContainerStartedAt &&
		a.RedisPONG && a.RedisPID > 0 && a.RedisRunID != "" && a.RedisRunID != b.RedisRunID && a.RedisClusterID == redisID &&
		a.NetworkConnected && a.NetworkCount == 1
}

// The collection function lets unit tests exercise real deadline, stability and
// regression decisions without substituting fake evidence for integration tests.
func observeLifecycleSynchronization(ctx context.Context, c config.Config, r *report.Report, selected cluster.Selection, initial cluster.Snapshot, combined, service, policy *Stability, collect func(context.Context) report.Sample) {
	l := r.LifecycleFault
	r.Recovery.CombinedRecovered, r.Recovery.TopologyRecovered, r.Recovery.ClientRecovered, r.Recovery.RedundancyRestored = false, false, false, false
	if r.Recovery.WritePolicyRecovered != nil {
		unavailable := false
		r.Recovery.WritePolicyRecovered = &unavailable
	}
	for ctx.Err() == nil {
		s := collect(ctx)
		s.Phase = "reconciliation"
		s.TopologyOK, s.Reasons = lifecycleRestored(s.Snapshot, selected, initial, l.TargetKind == "replica")
		if recoverySampleExpired(ctx, s) {
			s.Excluded = true
			s.Reasons = append(s.Reasons, "sample crossed the repair deadline; excluded from restoration decisions")
			appendSample(r, s)
			r.Final = &s.Snapshot
			break
		}
		r.Final = &s.Snapshot
		if l.RepairSnapshot == nil {
			l.RepairSnapshot = &s.Snapshot
		}
		updateLifecycleMilestones(r, s, selected, initial)
		dimensions := *l.Dimensions
		s.Dimensions = &dimensions
		appendSample(r, s)
		updatePolicyRecovery(r, s, policy, c.Experiment.StableFor.Value())
		at := s.Probe.At
		r.Recovery.TopologyRecovered = s.TopologyOK
		r.Recovery.ClientRecovered = service.Update(at, s.Probe.Successful, c.Experiment.StableFor.Value())
		r.Recovery.CombinedRecovered = combined.Update(at, s.TopologyOK && s.Probe.Successful, c.Experiment.StableFor.Value())
		r.Recovery.RedundancyRestored, _ = cluster.Ready(s.Snapshot)
		if r.NetworkFault != nil {
			n := r.NetworkFault
			n.Reconciled = r.Recovery.CombinedRecovered
			if s.TopologyOK && n.ReconciledAt == nil {
				n.ReconciledAt = milestone(s.Snapshot.Completed, *r.FaultRequestedAt)
			}
			if r.Recovery.RedundancyRestored && n.RedundancyRestoredAt == nil {
				n.RedundancyRestoredAt = milestone(s.Snapshot.Completed, *r.FaultRequestedAt)
			}
		}
		if r.Recovery.CombinedRecovered {
			if r.Recovery.ClientStable == nil {
				r.Recovery.ClientStable = milestone(at, *r.FaultRequestedAt)
			}
			if r.Recovery.CombinedStable == nil {
				r.Recovery.CombinedStable = milestone(at, *r.FaultRequestedAt)
			}
			l.StableRedundancy = milestone(at, *r.FaultRequestedAt)
			duration := float64(at.Sub(*l.RepairCompletedAt)) / float64(time.Millisecond)
			l.StableRedundancySinceRepairMS = &duration
			if r.NetworkFault != nil {
				duration := float64(at.Sub(*r.NetworkFault.RestoreCompletedAt)) / float64(time.Millisecond)
				r.NetworkFault.StableReconciliationMS = &duration
			}
			return
		}
		if delay(ctx, c.Experiment.SampleInterval.Value()) != nil {
			break
		}
	}
	l.RepairTimedOut = errors.Is(ctx.Err(), context.DeadlineExceeded)
	r.Recovery.CombinedRecovered, r.Recovery.RedundancyRestored = false, false
	if r.NetworkFault != nil {
		r.NetworkFault.Reconciled = false
		r.NetworkFault.ReconciliationTimedOut = l.RepairTimedOut
	}
	if r.Recovery.WritePolicyRecovered != nil {
		unavailable := false
		r.Recovery.WritePolicyRecovered = &unavailable
	}
	reason := fmt.Sprintf("original six-node redundancy and clients did not stabilize within the %s repair budget", c.Timeouts.Repair.Value())
	l.Errors = append(l.Errors, reason)
	r.Errors = append(r.Errors, reason)
}

func updateReplicaMilestones(r *report.Report, s report.Sample, selected cluster.Selection, fault time.Time) {
	for _, v := range s.Snapshot.Views {
		if !v.Reachable || len(v.Errors) != 0 {
			continue
		}
		for _, n := range v.Nodes {
			if n.ID == selected.ReplicaID && (n.HasFlag("fail") || n.HasFlag("fail?")) && r.Recovery.FailureDetected == nil {
				r.Recovery.FailureDetected = milestone(s.Snapshot.Completed, fault)
			}
		}
	}
	if s.TopologyOK {
		if r.Recovery.FullCoverage == nil {
			r.Recovery.FullCoverage = milestone(s.Snapshot.Completed, fault)
		}
		if r.Recovery.TopologyConverged == nil {
			r.Recovery.TopologyConverged = milestone(s.Snapshot.Completed, fault)
		}
	}
	if s.Probe.Successful && r.Recovery.ClientFirstSuccess == nil {
		r.Recovery.ClientFirstSuccess = milestone(s.Probe.At, fault)
	}
}

func updateLifecycleMilestones(r *report.Report, s report.Sample, selected cluster.Selection, initial cluster.Snapshot) {
	l := r.LifecycleFault
	if l == nil || s.Excluded || r.FaultRequestedAt == nil {
		return
	}
	updateLifecycleDimensions(r, s, initial)
	at, fault := s.Snapshot.Completed, *r.FaultRequestedAt
	if s.Phase == "fault" {
		l.FaultSnapshot = &s.Snapshot
	}
	if l.FailureDetected == nil && r.Recovery.FailureDetected != nil {
		l.FailureDetected = r.Recovery.FailureDetected
	}
	for _, v := range s.Snapshot.Views {
		if !v.Reachable || len(v.Errors) != 0 {
			continue
		}
		if (v.RedisID == selected.TargetID || v.RedisID == selected.ReplicaID) && v.Replication["role"] == "master" && v.Replication["connected_slaves"] == "0" && l.ReplicationInterrupted == nil {
			l.ReplicationInterrupted = milestone(at, fault)
		}
		if v.Node.ID != l.Target.ID || v.RedisID != l.TargetRedisID || !l.RepairCompleted {
			continue
		}
		if l.TargetReachable == nil {
			l.TargetReachable = milestone(at, fault)
		}
		if v.Replication["role"] == "slave" && v.Replication["master_sync_in_progress"] == "1" && l.SyncStarted == nil {
			l.SyncStarted = milestone(at, fault)
		}
		expectedPrimary := selected.TargetID
		if l.TargetKind == "primary" {
			expectedPrimary = selected.ReplicaID
		}
		if v.Replication["role"] == "slave" && v.Replication["master_link_status"] == "up" && lifecycleReplicaAssignment(v, expectedPrimary, initial) {
			if l.LinkUp == nil {
				l.LinkUp = milestone(at, fault)
			}
			if v.Replication["master_sync_in_progress"] == "0" && l.SyncComplete == nil {
				l.SyncComplete = milestone(at, fault)
			}
		}
	}
	if l.RepairCompleted && l.MembershipRejoined == nil && lifecycleMembership(s.Snapshot, l.TargetRedisID, initial) {
		l.MembershipRejoined = milestone(at, fault)
	}
	if l.RepairCompleted && l.FaultSnapshot != nil {
		primaryID := selected.TargetID
		if l.TargetKind == "primary" {
			primaryID = selected.ReplicaID
		}
		beforeSnapshot := l.FaultSnapshot
		if l.TargetKind == "replica" {
			beforeSnapshot = &initial
		}
		for _, before := range beforeSnapshot.Views {
			for _, after := range s.Snapshot.Views {
				if before.RedisID == primaryID && after.RedisID == primaryID && after.Reachable && len(after.Errors) == 0 {
					evidence := cluster.ResynchronizationDelta(before, after)
					l.Resynchronization = &evidence
					if evidence.FullSyncDelta != nil && *evidence.FullSyncDelta > 0 && l.FullSyncObserved == nil {
						l.FullSyncObserved = milestone(at, fault)
					}
					if evidence.PartialAcceptedDelta != nil && *evidence.PartialAcceptedDelta > 0 && l.PartialSyncObserved == nil {
						l.PartialSyncObserved = milestone(at, fault)
					}
				}
			}
		}
	}
}

func lifecycleReplicaAssignment(v cluster.View, primaryID string, initial cluster.Snapshot) bool {
	assigned := false
	for _, n := range v.Nodes {
		if n.ID == v.RedisID && n.Replica() && !n.Primary() && n.PrimaryID == primaryID {
			assigned = true
		}
	}
	for _, primary := range initial.Views {
		if primary.RedisID == primaryID {
			host, port, err := net.SplitHostPort(primary.Node.InternalAddr)
			return assigned && err == nil && v.Replication["master_host"] == host && v.Replication["master_port"] == port
		}
	}
	return false
}

func updateLifecycleDimensions(r *report.Report, s report.Sample, initial cluster.Snapshot) {
	l := r.LifecycleFault
	d := &report.RecoveryDimensions{}
	membership := lifecycleMembership(s.Snapshot, l.TargetRedisID, initial)
	d.MembershipHealthy = &membership
	coverage, leaders, observations := len(s.Snapshot.Views) == 6, true, 0
	for _, v := range s.Snapshot.Views {
		if !v.Reachable || len(v.Errors) != 0 {
			continue
		}
		observations++
		coverage = coverage && v.Coverage.Full()
		if v.Node.ID == l.Target.ID && v.RedisID == l.TargetRedisID {
			live := true
			d.TargetProcessLive = &live
			d.TargetRole = v.Replication["role"]
		}
		for _, owner := range v.Coverage.Owners {
			ownerLive := false
			for _, direct := range s.Snapshot.Views {
				if direct.RedisID != owner.RedisID || !direct.Reachable || len(direct.Errors) != 0 || direct.Replication["role"] != "master" {
					continue
				}
				primary := false
				for _, self := range direct.Nodes {
					if self.ID == direct.RedisID && self.Primary() && !self.Replica() && self.Healthy() {
						primary = true
					}
				}
				for _, own := range direct.Coverage.Owners {
					if primary && own.RedisID == direct.RedisID && own.Start <= owner.Start && own.End >= owner.End {
						ownerLive = true
					}
				}
			}
			leaders = leaders && ownerLive
		}
	}
	coverage = coverage && observations > 0
	leaders = leaders && coverage
	coverage = coverage && leaders
	d.SlotCoverage, d.LeadersAvailable = &coverage, &leaders
	d.ReadAvailable, d.WriteAvailable = s.Probe.ReadSuccessful, s.Probe.WriteSuccessful
	// Older schemas do not distinguish read and write failures.
	if s.Probe.Successful && d.ReadAvailable == nil && d.WriteAvailable == nil {
		available := true
		d.ReadAvailable, d.WriteAvailable = &available, &available
	}
	if s.WritePolicyProbe != nil {
		available := s.WritePolicyProbe.Successful
		d.WritePolicyAvailable = &available
	} else if r.Config.Workload.WritePolicy == "set" && s.Probe.Successful {
		available := true
		d.WritePolicyAvailable = &available
	}
	ready, _ := cluster.Ready(s.Snapshot)
	d.RedundancyRestored = &ready
	if r.NetworkFault != nil {
		l.ProcessSamples = append([]report.ProcessSample(nil), r.NetworkFault.ProcessSamples...)
		if len(l.ProcessSamples) > 0 {
			last := l.ProcessSamples[len(l.ProcessSamples)-1]
			if last.Error == "" {
				live := last.Evidence.ContainerRunning && last.Evidence.RedisPONG
				d.TargetProcessLive = &live
			}
		}
	} else if !l.RepairCompleted && r.FaultCompletedAt != nil {
		live := false // a successfully completed SIGKILL, not a reachability guess
		d.TargetProcessLive = &live
	}
	// Data and persistence are established by separate experiments/validation.
	// An ordinary six-node replication sample establishes neither.
	if l.Dimensions != nil {
		d.DataVerified = l.Dimensions.DataVerified
		d.PersistenceRecovered = l.Dimensions.PersistenceRecovered
	}
	l.Dimensions = d
	r.Dimensions = d
}

func lifecycleMembership(s cluster.Snapshot, targetID string, initial cluster.Snapshot) bool {
	if len(s.Views) != 6 || len(initial.Views) != 6 {
		return false
	}
	known, redisIDs, seen := make(map[string]cluster.View), make(map[string]bool), make(map[string]bool)
	for _, v := range initial.Views {
		if v.Node.ID == "" || v.RedisID == "" || known[v.Node.ID].RedisID != "" || redisIDs[v.RedisID] {
			return false
		}
		known[v.Node.ID], redisIDs[v.RedisID] = v, true
	}
	for _, v := range s.Views {
		original, exists := known[v.Node.ID]
		if !exists || seen[v.Node.ID] || v.Node != original.Node || v.RedisID != original.RedisID || !v.Reachable || len(v.Errors) != 0 || len(v.Nodes) != 6 {
			return false
		}
		seen[v.Node.ID] = true
		found := false
		nodeIDs := make(map[string]bool)
		for _, n := range v.Nodes {
			if !n.Healthy() || !redisIDs[n.ID] || nodeIDs[n.ID] {
				return false
			}
			nodeIDs[n.ID] = true
			if n.ID == targetID && n.Healthy() {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func lifecycleMethodology(c config.Config) []string {
	kind := "selected replica"
	if c.Experiment.Fault == "primary-kill" {
		kind = "selected primary after stable five-survivor service recovery"
	}
	return []string{
		"Fault selection preserves the initial primary and its replica. The destructive target is the " + kind + "; only strictly owned local Docker resources are affected.",
		"Replica faults preserve all original slot owners and do not require promotion. down_duration is a minimum hold from successful fault action completion, bounded by the recovery observation budget. Client service, failure gossip and replication interruption are separately sampled.",
		"Replica network isolation removes its complete bidirectional Docker bridge connectivity while Redis remains running. It is not a selective replication-only partition. Restoration uses the same owned endpoint and unchanged live process identity.",
		"Administrative endpoint-disconnection evidence is separate from Redis INFO link or connected-replica indicators. A short disconnection can finish before TCP/link failure detection; absent interruption or synchronization-counter milestones remain unobserved and are not invented from the Docker interface state.",
		"Crash repair starts the same stopped owned container with its bounded disposable data volume. Redis Cluster identity must match the pre-fault identity. Restart changes the process; it does not simulate power loss, host/storage failure or persistence durability.",
		"The repair deadline covers the restart/reconnect action and subsequent synchronization. Redis's current membership is inspected without issuing CLUSTER REPLICATE or forcing slot ownership. Replica repair preserves original assignments; former-primary repair requires the original replica to retain promoted ownership and the old primary to rejoin as its replica.",
		"Fresh stability after repair requires all six original identities and endpoints, healthy equal slot maps, exactly one synchronized replica per primary, and all-shard client probes. Historical service recovery remains separate from final redundancy and configured acknowledgement-policy availability.",
		"Replication link-up and master_sync_in_progress=0 are sampled indicators, not exact zero lag, bytes preserved, or identification of partial versus full resynchronization. Raw replication INFO is retained for inspection; unobserved milestones remain absent.",
		"Traffic drains before validating every acknowledged immutable SET, including short/error WAIT/WAITAOF outcomes. Ambiguous SET errors remain separate. Cancellation skips unsupported recovery/data claims and still uses fresh bounded restoration and cleanup contexts, including managed volumes.",
	}
}
