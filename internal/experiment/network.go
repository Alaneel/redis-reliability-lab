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

func validateIsolationReceipt(x *environment.IsolationReceipt, env *environment.Environment, target environment.Node) error {
	if x == nil {
		return errors.New("isolation manager returned no restoration receipt")
	}
	host, _, err := net.SplitHostPort(target.InternalAddr)
	if env == nil || err != nil || x.RunID != env.RunID || x.NetworkID != env.NetworkID || x.Node != target || x.OriginalIP != host || x.HostAddr != target.Addr {
		return errors.New("receipt does not identify the selected owned node and its original endpoints")
	}
	return nil
}

func restoreNetwork(ctx context.Context, manager NetworkFaultManager, r *report.Report) error {
	n := r.NetworkFault
	at := time.Now()
	n.RestoreRequestedAt = &at
	if err := manager.Restore(ctx, r.Environment, n.Receipt); err != nil {
		n.Errors = append(n.Errors, err.Error())
		return err
	}
	completed := time.Now()
	duration := float64(completed.Sub(at)) / float64(time.Millisecond)
	n.RestoreCompletedAt, n.RestoreDurationMS = &completed, &duration
	n.ConnectivityRestored = true
	return nil
}

func captureProcess(ctx context.Context, manager NetworkFaultManager, r *report.Report, target environment.Node) {
	probeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	evidence, err := manager.ProbeProcess(probeCtx, r.Environment, target)
	cancel()
	s := report.ProcessSample{Evidence: evidence}
	if err != nil {
		s.Error = err.Error()
	}
	// Preserve first evidence and latest state while keeping reports bounded.
	if len(r.NetworkFault.ProcessSamples) < 64 {
		r.NetworkFault.ProcessSamples = append(r.NetworkFault.ProcessSamples, s)
	} else {
		r.NetworkFault.ProcessSamples[63] = s
	}
}

func processSurvivedIsolation(n *report.NetworkFault) bool {
	if n == nil || n.Receipt == nil {
		return false
	}
	before := n.Receipt.BeforeProcess
	for _, s := range n.ProcessSamples {
		p := s.Evidence
		if s.Error == "" && p.ContainerRunning && p.RedisPONG && !p.NetworkConnected && p.NetworkCount == 0 &&
			p.ContainerPID == before.ContainerPID && p.ContainerStartedAt == before.ContainerStartedAt &&
			p.RedisRunID != "" && p.RedisRunID == before.RedisRunID && p.RedisPID == before.RedisPID {
			return true
		}
	}
	return false
}

func observeReconciliation(ctx context.Context, c config.Config, r *report.Report, observer *cluster.Observer, client workload.Client, writer workload.WriteExecutor, keys []string, selected cluster.Selection, initial cluster.Snapshot, combined, service, policy *Stability) {
	n := r.NetworkFault
	reconcileCtx, cancel := context.WithTimeout(ctx, c.Timeouts.Reconciliation.Value())
	defer cancel()
	// A prior majority-side milestone cannot establish post-restoration health.
	r.Recovery.CombinedRecovered, r.Recovery.TopologyRecovered, r.Recovery.ClientRecovered = false, false, false
	if r.Recovery.WritePolicyRecovered != nil {
		unavailable := false
		r.Recovery.WritePolicyRecovered = &unavailable
	}
	for reconcileCtx.Err() == nil {
		sample := sampleRecovery(reconcileCtx, observer, client, keys, r.RunID, c.Workload.OperationTimeout.Value(), selected, initial)
		sample.Phase = "reconciliation"
		sample.TopologyOK, sample.Reasons = cluster.ReconciledWithBaseline(sample.Snapshot, selected, initial)
		if requiresAcknowledgementPolicy(c.Workload.WritePolicy) {
			p := ProbeAcknowledgementPolicy(reconcileCtx, writer, client, keys, r.RunID, c.Workload.WritePolicy, c.Workload.WaitReplicas, c.Workload.OperationTimeout.Value())
			sample.WritePolicyProbe = &p
		}
		if recoverySampleExpired(reconcileCtx, sample) {
			sample.Excluded = true
			sample.Reasons = append(sample.Reasons, "sample crossed the reconciliation deadline; excluded from recovery decisions")
			appendSample(r, sample)
			r.Final = &sample.Snapshot
			break
		}
		appendSample(r, sample)
		r.Final = &sample.Snapshot
		updatePolicyRecovery(r, sample, policy, c.Experiment.StableFor.Value())
		at := sample.Probe.At
		r.Recovery.RedundancyRestored, _ = cluster.Ready(sample.Snapshot)
		if r.Recovery.RedundancyRestored && n.RedundancyRestoredAt == nil {
			n.RedundancyRestoredAt = milestone(sample.Snapshot.Completed, *r.FaultRequestedAt)
		}
		if sample.TopologyOK && n.ReconciledAt == nil {
			n.ReconciledAt = milestone(sample.Snapshot.Completed, *r.FaultRequestedAt)
		}
		r.Recovery.TopologyRecovered = sample.TopologyOK
		r.Recovery.ClientRecovered = service.Update(at, sample.Probe.Successful, c.Experiment.StableFor.Value())
		r.Recovery.CombinedRecovered = combined.Update(at, sample.TopologyOK && sample.Probe.Successful, c.Experiment.StableFor.Value())
		n.Reconciled = r.Recovery.CombinedRecovered
		if n.Reconciled {
			duration := float64(at.Sub(*n.RestoreCompletedAt)) / float64(time.Millisecond)
			n.StableReconciliationMS = &duration
			return
		}
		if delay(reconcileCtx, c.Experiment.SampleInterval.Value()) != nil {
			break
		}
	}
	n.ReconciliationTimedOut = reconcileCtx.Err() == context.DeadlineExceeded
	n.Errors = append(n.Errors, fmt.Sprintf("topology and client reconciliation did not stabilize within %s", c.Timeouts.Reconciliation.Value()))
	// Do not retain an earlier positive flag after the decision deadline expires.
	n.Reconciled, r.Recovery.CombinedRecovered = false, false
	if r.Recovery.WritePolicyRecovered != nil {
		unavailable := false
		r.Recovery.WritePolicyRecovered = &unavailable
	}
}

func networkMethodology() []string {
	return []string{
		"Fault is a complete bidirectional Docker bridge disconnection of the deterministic primary owning the lowest initial slot. Redis remains running. All peer and published client connectivity is removed; this is not a selective or one-way partition.",
		"The orchestration and action-completion anchors are recorded separately using monotonic runtime durations. isolation_duration is a minimum hold from successful action completion, bounded by the recovery deadline; samples and bounded process probes can extend the hold. Actual restoration timestamps record the observed duration.",
		"During isolation, recovery requires the five surviving nodes to converge on healthy usable slot ownership, the original replica to own the selected primary slots, unchanged unaffected ownership, and successful all-shard SET/GET probes over stable_for. A stable milestone is retained while observation continues until restoration.",
		"Docker exec runs fixed Redis liveness and topology commands independent of the disconnected data plane. PONG and unchanged container PID, start time and Redis run ID prove sampled process survival, not cluster health. Isolation success requires at least one such disconnected observation.",
		"Restoration reconnects only the same owned node to the same owned network with its original private IP and loopback port. Process identity and endpoint identity must remain unchanged. Cancellation uses a fresh bounded restoration attempt before independent cleanup; safe cleanup can directly remove an owned detached node.",
		"After restoration, fresh stability tracking requires all six original identities and endpoints, three healthy primaries with synchronized replicas, the promoted replica retaining shard ownership, former-primary demotion, unchanged unaffected shards, and all-shard client probes. Service recovery, connectivity restoration and replication redundancy have separate evidence.",
		"Poll intervals are delays after sample completion; concurrent views are not atomic snapshots. Probes do not prove continuous availability. Workload metrics include terminal client errors separately from probe results. At most 1200 topology samples and 64 process samples are retained.",
		"Traffic drains before validating every successfully acknowledged immutable SET. Missing, corrupt, unreadable or unverified records prevent full recovery. Error-returning SET outcomes remain ambiguous. No client can write to the disconnected minority, so this experiment does not exercise minority-side write conflicts or establish loss-free asynchronous replication.",
	}
}
