package batch

import (
	"net"

	"github.com/alanwang/redis-reliability-lab/internal/report"
)

// recoveredNetworkEvidence validates the network experiment's evidence without
// importing its orchestration package. Majority-side recovery while disconnected
// is distinct from restored connectivity and fresh six-node reconciliation.
func recoveredNetworkEvidence(r *report.Report) bool {
	n := r.NetworkFault
	if (r.SchemaVersion != 2 && r.SchemaVersion != report.LifecycleSchemaVersion) || n == nil || n.Receipt == nil || r.Environment == nil ||
		!n.ConnectivityRestored || !n.Reconciled || n.ReconciliationTimedOut || !r.Recovery.RedundancyRestored ||
		n.RestoreRequestedAt == nil || n.RestoreRequestedAt.IsZero() || n.RestoreCompletedAt == nil || n.RestoreCompletedAt.IsZero() ||
		!nonnegativeObservation(n.RestoreDurationMS) || !nonnegativeObservation(n.StableReconciliationMS) ||
		!observedMilestone(n.ReconciledAt) || !observedMilestone(n.RedundancyRestoredAt) {
		return false
	}
	receipt := n.Receipt
	if r.Selection == nil {
		return false
	}
	target := r.Selection.Target
	if r.Config.Experiment.Fault == "replica-isolate" {
		target = r.Selection.Replica
	}
	ownedTarget := false
	for _, node := range r.Environment.Nodes {
		if node.ID == target.ID && node.Addr == target.Addr && node.InternalAddr == target.InternalAddr && receipt.Node == node {
			ownedTarget = true
		}
	}
	originalIP, _, addressErr := net.SplitHostPort(target.InternalAddr)
	if !ownedTarget || addressErr != nil || receipt.OriginalIP != originalIP || receipt.HostAddr != target.Addr {
		return false
	}
	if receipt.RunID != r.RunID || receipt.RunID != r.Environment.RunID ||
		receipt.NetworkID == "" || receipt.NetworkID != r.Environment.NetworkID || receipt.DisconnectedAt.IsZero() {
		return false
	}
	before := receipt.BeforeProcess
	if !before.ContainerRunning || !before.RedisPONG || !before.NetworkConnected || before.NetworkCount != 1 ||
		before.ContainerPID <= 0 || before.ContainerStartedAt == "" || before.RedisPID <= 0 || before.RedisRunID == "" {
		return false
	}
	for _, sample := range n.ProcessSamples {
		p := sample.Evidence
		if sample.Error == "" && !p.ObservedAt.IsZero() && p.ContainerRunning && p.RedisPONG && !p.NetworkConnected && p.NetworkCount == 0 &&
			p.ContainerPID == before.ContainerPID && p.ContainerStartedAt == before.ContainerStartedAt &&
			p.RedisPID == before.RedisPID && p.RedisRunID == before.RedisRunID {
			return true
		}
	}
	return false
}

func nonnegativeObservation(value *float64) bool {
	return value != nil && finite(*value) && *value >= 0
}

func observedMilestone(value *report.Milestone) bool {
	return value != nil && !value.At.IsZero() && finite(value.SinceFaultMS) && value.SinceFaultMS >= 0
}
