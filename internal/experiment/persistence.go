package experiment

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/cluster"
	"github.com/Alaneel/redis-reliability-lab/internal/config"
	"github.com/Alaneel/redis-reliability-lab/internal/environment"
	"github.com/Alaneel/redis-reliability-lab/internal/metrics"
	"github.com/Alaneel/redis-reliability-lab/internal/report"
	"github.com/Alaneel/redis-reliability-lab/internal/workload"
)

const PersistenceFailureModel = report.PersistenceFailureModel

type persistenceTraffic interface {
	Stop()
	Stats() workload.Stats
	Snapshot() metrics.Snapshot
	Ledger() []workload.Write
	Validate(context.Context, time.Duration) workload.Validation
}

type persistenceObserver interface {
	Observe(context.Context) cluster.Snapshot
}

// Administrative preparation is deliberately before the workload: its Docker
// inspection/CONFIG/INFO/file-list latency must not extend the drain-to-kill gap.
func preparePersistence(ctx context.Context, c config.Config, r *report.Report, prober ProcessProber) error {
	if _, _, err := persistenceNodeOrder(r.Environment, r.Initial); err != nil {
		return err
	}
	budget, cancel := context.WithTimeout(ctx, c.Timeouts.Readiness.Value())
	defer cancel()
	return observePersistenceProcesses(budget, c, r, prober, false)
}

func observePersistenceProcesses(ctx context.Context, c config.Config, r *report.Report, prober ProcessProber, restored bool) error {
	for _, n := range r.Environment.Nodes {
		var evidence environment.ProcessEvidence
		var err error
		for {
			op, done := context.WithTimeout(ctx, 5*time.Second)
			evidence, err = prober.ProbeProcess(op, r.Environment, n)
			if err == nil && op.Err() != nil {
				err = op.Err()
			}
			done()
			if err != nil {
				break
			}
			if persistencePreparationBusy(evidence) {
				if err = delay(ctx, c.Experiment.SampleInterval.Value()); err != nil {
					break
				}
				continue
			}
			err = persistenceRuntimeEvidence(evidence, c.Environment.Persistence)
			if err == nil {
				err = persistencePreparationQuiet(evidence)
			}
			break
		}
		if err == nil && (evidence.RedisClusterID != n.RedisClusterID || !evidence.NetworkConnected || evidence.NetworkCount != 1) {
			err = errors.New("prepared process identity or managed network does not match inventory")
		}
		if err == nil && restored && !persistenceRestoredProcess(r, n, evidence) {
			err = errors.New("restored runtime is not the same observed restarted original process")
		}
		sample := report.NodeProcessSample{Node: n, OriginalRedisID: n.RedisClusterID, Evidence: evidence}
		if err != nil {
			sample.Error = err.Error()
		}
		stage := "preparation"
		if restored {
			r.PersistenceFault.Restored = append(r.PersistenceFault.Restored, sample)
			stage = "restored runtime"
		} else if r.PersistenceFault != nil {
			r.PersistenceFault.Prepared = append(r.PersistenceFault.Prepared, sample)
		} else {
			r.PersistencePrepared = append(r.PersistencePrepared, sample)
		}
		if err != nil {
			return persistenceError(r, fmt.Errorf("persistence %s %s: %w", stage, n.ID, err))
		}
	}
	return nil
}

// The original role inventory determines action order. Original replicas stop
// first so a surviving replica cannot promote while primaries are being killed.
func persistenceNodeOrder(env *environment.Environment, initial *cluster.Snapshot) (kills, restarts []environment.Node, err error) {
	if env == nil || initial == nil || !env.RetainData || len(env.Nodes) != 6 || len(env.Volumes) != 6 {
		return nil, nil, errors.New("persistence restart requires six original nodes and retained owned volumes")
	}
	if ready, why := cluster.Ready(*initial); !ready {
		return nil, nil, fmt.Errorf("persistence original topology is not ready: %v", why)
	}
	views, volumes := map[string]cluster.View{}, map[string]bool{}
	for _, v := range initial.Views {
		views[v.Node.ID] = v
	}
	for _, v := range env.Volumes {
		if v.NodeID == "" || v.Name == "" || volumes[v.NodeID] {
			return nil, nil, errors.New("duplicate or incomplete owned volume inventory")
		}
		volumes[v.NodeID] = true
	}
	seen, redisIDs := map[string]bool{}, map[string]bool{}
	var primaries, replicas []environment.Node
	for _, n := range env.Nodes {
		v, exists := views[n.ID]
		if !exists || seen[n.ID] || n.RedisClusterID == "" || redisIDs[n.RedisClusterID] || v.RedisID != n.RedisClusterID || n.Addr != v.Node.Addr || n.InternalAddr != v.Node.InternalAddr || !volumes[n.ID] {
			return nil, nil, errors.New("managed node, Redis identity or volume does not match original topology")
		}
		seen[n.ID], redisIDs[n.RedisClusterID] = true, true
		switch v.Replication["role"] {
		case "master":
			primaries = append(primaries, n)
		case "slave":
			replicas = append(replicas, n)
		default:
			return nil, nil, errors.New("original role is unobserved")
		}
	}
	if len(primaries) != 3 || len(replicas) != 3 {
		return nil, nil, errors.New("original topology must contain three primaries and three replicas")
	}
	sort.Slice(primaries, func(i, j int) bool { return primaries[i].ID < primaries[j].ID })
	sort.Slice(replicas, func(i, j int) bool { return replicas[i].ID < replicas[j].ID })
	kills = append(append([]environment.Node{}, replicas...), primaries...)
	restarts = append(append([]environment.Node{}, primaries...), replicas...)
	return kills, restarts, nil
}

func runPersistenceExperiment(ctx context.Context, c config.Config, r *report.Report, manager EnvironmentManager, prober ProcessProber, restart RestartManager, observer persistenceObserver, traffic persistenceTraffic, client workload.Client, writer workload.WriteExecutor, keys []string, move func(Stage)) error {
	p := r.PersistenceFault
	kills, restarts, err := persistenceNodeOrder(r.Environment, r.Initial)
	if err != nil {
		return persistenceError(r, err)
	}
	move(Quiescing)
	traffic.Stop() // joins in-flight SET acknowledgements before freezing anything
	drained := time.Now()
	p.TrafficDrainedAt = &drained
	r.WorkloadStats, r.Metrics = traffic.Stats(), traffic.Snapshot()
	ledger := traffic.Ledger()
	frozen := time.Now()
	p.LedgerFrozenAt, p.AcknowledgedFrozen, p.AmbiguousFrozen = &frozen, r.WorkloadStats.AcknowledgedWrites, r.WorkloadStats.AmbiguousWrites
	p.PreCrashSnapshot = r.Initial
	if p.AcknowledgedFrozen == 0 || uint64(len(ledger)) != p.AcknowledgedFrozen || r.WorkloadStats.StartedRequests != r.WorkloadStats.CompletedRequests || r.WorkloadStats.CompletedRequests != r.Metrics.Totals.Attempts || r.WorkloadStats.OfferedRequests != r.WorkloadStats.StartedRequests+r.WorkloadStats.DroppedRequests {
		return persistenceError(r, errors.New("drained workload counters do not reconcile with the nonempty frozen ledger"))
	}
	baseline := r.Metrics.Phases["baseline"]
	if baseline.Attempts == 0 || baseline.SuccessRate < c.Experiment.MinimumBaselineSuccess {
		return persistenceError(r, errors.New("fully drained baseline did not meet the success criterion"))
	}
	if ctx.Err() != nil {
		return ctx.Err()
	}
	move(Injecting)
	faultAt := time.Now()
	r.FaultRequestedAt = &faultAt
	actionCtx, actionCancel := context.WithTimeout(ctx, c.Timeouts.Recovery.Value())
	err = stopPersistenceNodes(actionCtx, r, manager, prober, kills)
	actionCancel()
	if err != nil {
		return err
	}
	move(Observing)
	if err := delay(ctx, c.Experiment.DownDuration.Value()); err != nil {
		return err
	}
	move(Repairing)
	repairCtx, repairCancel := context.WithTimeout(ctx, c.Timeouts.Repair.Value())
	if err := restartPersistenceNodes(repairCtx, r, restart, restarts); err != nil {
		repairCancel()
		return err
	}
	move(Reconciling)
	observePersistenceTopology(repairCtx, c, r, observer)
	if !p.RestoreTimedOut && repairCtx.Err() == nil {
		if err := observePersistenceProcesses(repairCtx, c, r, prober, true); err != nil {
			p.RestoreTimedOut = errors.Is(err, context.DeadlineExceeded) || errors.Is(repairCtx.Err(), context.DeadlineExceeded)
			repairCancel()
			return err
		}
	}
	repairCancel()
	if ctx.Err() != nil {
		return ctx.Err()
	}
	// Even a topology deadline can leave useful missing/unreadable evidence. A
	// fresh validation budget records its coverage without erasing the deadline.
	move(Validating)
	validationCtx, validationCancel := context.WithTimeout(ctx, c.Timeouts.Validation.Value())
	r.Validation = traffic.Validate(validationCtx, c.Workload.OperationTimeout.Value())
	validationCancel()
	validation := r.Validation
	p.FrozenValidation, p.DataCheckedAt = &validation, &validation.CompletedAt
	if validation.StartedAt.IsZero() || validation.CompletedAt.IsZero() || validation.StartedAt.Before(*p.RestartCompletedAt) || validation.CompletedAt.Before(validation.StartedAt) {
		return persistenceError(r, errors.New("frozen validation did not provide an actual post-restart check interval; diagnostic writes are withheld"))
	}
	r.Recovery.DataVerified = validation.Complete && validation.Expected > 0 && uint64(validation.Expected) == p.AcknowledgedFrozen && validation.Checked == validation.Expected && validation.Missing+validation.Corrupt+validation.Unreadable+validation.Unverified == 0 && r.WorkloadStats.MissingReads+r.WorkloadStats.MismatchedReads == 0
	dimensions := report.RecoveryDimensions{}
	if r.Dimensions != nil {
		dimensions = *r.Dimensions
	}
	dataVerified := r.Recovery.DataVerified
	dimensions.DataVerified = &dataVerified
	r.Dimensions = &dimensions
	updatePersistenceVerdict(r)
	if ctx.Err() != nil {
		return ctx.Err()
	}
	move(Diagnosing)
	diagnosticsAt := time.Now()
	p.DiagnosticsRequestedAt = &diagnosticsAt
	diagnosticCtx, diagnosticCancel := context.WithTimeout(ctx, c.Timeouts.Recovery.Value())
	collect := func(sampleCtx context.Context) report.Sample {
		snapshot := observer.Observe(sampleCtx)
		probe := ProbeAll(sampleCtx, client, keys, r.RunID+":diagnostic:"+fmt.Sprint(time.Now().UnixNano()), c.Workload.OperationTimeout.Value())
		s := report.Sample{Snapshot: snapshot, Probe: probe}
		if requiresAcknowledgementPolicy(c.Workload.WritePolicy) {
			policy := ProbeAcknowledgementPolicy(sampleCtx, writer, client, keys, r.RunID, c.Workload.WritePolicy, c.Workload.WaitReplicas, c.Workload.OperationTimeout.Value())
			s.WritePolicyProbe = &policy
		}
		return s
	}
	observePersistenceDiagnostics(diagnosticCtx, c, r, collect)
	diagnosticCancel()
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if p.RestoreTimedOut || p.StableRedundancy == nil {
		r.Recovery.CombinedRecovered = false
	}
	r.Outcome = Outcome(r.Recovery)
	if r.Outcome != "recovered" {
		return persistenceError(r, fmt.Errorf("persistence experiment ended %s: frozen data verified=%t, original redundancy stable=%t, restoration deadline=%t", r.Outcome, r.Recovery.DataVerified, p.StableRedundancy != nil, p.RestoreTimedOut))
	}
	return nil
}

func stopPersistenceNodes(ctx context.Context, r *report.Report, manager EnvironmentManager, prober ProcessProber, nodes []environment.Node) error {
	p := r.PersistenceFault
	for _, n := range nodes {
		if ctx.Err() != nil {
			return persistenceError(r, ctx.Err())
		}
		at := time.Now()
		if len(p.KillActions) == 0 {
			elapsed := milliseconds(at.Sub(*p.TrafficDrainedAt))
			p.DrainToFirstKillMS = &elapsed
		}
		action := report.NodeAction{Node: n, RequestedAt: at}
		op, cancel := context.WithTimeout(ctx, 10*time.Second)
		err := manager.Kill(op, r.Environment, n)
		if err == nil && op.Err() != nil {
			err = op.Err()
		}
		cancel()
		if err == nil {
			completed := time.Now()
			action.CompletedAt = &completed
		} else {
			action.Error = err.Error()
		}
		p.KillActions = append(p.KillActions, action)
		if err != nil {
			return persistenceError(r, fmt.Errorf("kill %s; restart barrier not established: %w", n.ID, err))
		}
	}
	lastKill := *p.KillActions[len(p.KillActions)-1].CompletedAt
	elapsed := milliseconds(lastKill.Sub(*p.TrafficDrainedAt))
	p.DrainToLastKillMS = &elapsed
	for _, n := range nodes {
		op, cancel := context.WithTimeout(ctx, 5*time.Second)
		evidence, err := prober.ProbeProcess(op, r.Environment, n)
		accepted := err == environment.ErrProcessStopped && op.Err() == nil && !evidence.StartedAt.IsZero() && !evidence.ObservedAt.IsZero() && !evidence.ObservedAt.Before(evidence.StartedAt) && !evidence.ObservedAt.Before(lastKill) && evidence.ContainerStartedAt != "" && !evidence.ContainerRunning && !evidence.RedisPONG && !evidence.NetworkConnected && evidence.ContainerRestartPolicy == "no"
		cancel()
		sample := report.NodeProcessSample{Node: n, OriginalRedisID: n.RedisClusterID, Evidence: evidence}
		if err != nil {
			sample.Error = err.Error()
		}
		p.Stopped = append(p.Stopped, sample)
		if !accepted {
			return persistenceError(r, fmt.Errorf("stopped barrier %s requires authoritative owned stopped-process proof, got %v", n.ID, err))
		}
		for i := range p.KillActions {
			if p.KillActions[i].Node.ID == n.ID {
				copy := evidence
				p.KillActions[i].Stopped = &copy
			}
		}
	}
	completed := time.Now()
	p.AllStoppedAt, r.FaultCompletedAt = &completed, &completed
	duration := milliseconds(completed.Sub(*r.FaultRequestedAt))
	r.FaultActionDurationMS = &duration
	return nil
}

func restartPersistenceNodes(ctx context.Context, r *report.Report, manager RestartManager, nodes []environment.Node) error {
	p := r.PersistenceFault
	if len(p.Stopped) != 6 || p.AllStoppedAt == nil {
		return persistenceError(r, errors.New("all-six stopped barrier is absent; refusing any restart"))
	}
	for _, n := range nodes {
		if ctx.Err() != nil {
			p.RestoreTimedOut = errors.Is(ctx.Err(), context.DeadlineExceeded)
			return persistenceError(r, ctx.Err())
		}
		at := time.Now()
		if p.RestartRequestedAt == nil {
			p.RestartRequestedAt = &at
			elapsed := milliseconds(at.Sub(*p.AllStoppedAt))
			p.AllStoppedToFirstRestartMS = &elapsed
		}
		action := report.NodeAction{Node: n, RequestedAt: at}
		x, err := manager.Restart(ctx, r.Environment, n)
		if x != nil {
			p.Restarts = append(p.Restarts, x)
		}
		if err == nil && ctx.Err() != nil {
			err = ctx.Err()
		}
		if err == nil && !validRestartReceipt(x, r.RunID, n, n.RedisClusterID) {
			err = errors.New("restart returned incomplete or contradictory original-node process identity evidence")
		}
		if err == nil && (x.RequestedAt.Before(at) || x.Before.ObservedAt.Before(x.RequestedAt)) {
			err = errors.New("restart receipt observation precedes the actual requested action")
		}
		if err == nil && !persistenceNewProcess(r, n, x) {
			err = errors.New("restart did not prove replacement of the prepared original Redis process")
		}
		if err == nil {
			completed := time.Now()
			action.CompletedAt = &completed
		} else {
			action.Error = err.Error()
		}
		p.RestartActions = append(p.RestartActions, action)
		if err != nil {
			p.RestoreTimedOut = errors.Is(err, context.DeadlineExceeded) || errors.Is(ctx.Err(), context.DeadlineExceeded)
			return persistenceError(r, fmt.Errorf("restart %s: %w", n.ID, err))
		}
	}
	completed := time.Now()
	p.RestartCompletedAt = &completed
	r.Recovery.ProcessRestarted = true
	return nil
}

// No workload or service/policy probe is invoked by this observation loop.
func observePersistenceTopology(ctx context.Context, c config.Config, r *report.Report, observer persistenceObserver) {
	p := r.PersistenceFault
	stability := Stability{}
	for ctx.Err() == nil {
		snapshot := observer.Observe(ctx)
		ok, why := cluster.ReplicaRecoveryWithBaseline(snapshot, *r.Selection, *r.Initial)
		s := report.Sample{Phase: "persistence_recovery", Snapshot: snapshot, TopologyOK: ok, Reasons: why}
		r.Final = &snapshot
		if persistenceSampleExpired(ctx, snapshot.Completed) {
			s.Excluded = true
			s.Reasons = append(s.Reasons, "sample crossed the restoration deadline; excluded from persistence restoration decisions")
			appendSample(r, s)
			break
		}
		updatePersistenceDimensions(r, s)
		copy := *r.Dimensions
		s.Dimensions = &copy
		appendSample(r, s)
		r.Recovery.TopologyRecovered, r.Recovery.RedundancyRestored = ok, ok
		if ok && p.ReplicaReady == nil {
			p.ReplicaReady = milestone(snapshot.Completed, *r.FaultRequestedAt)
			r.Recovery.FullCoverage = p.ReplicaReady
		}
		if stability.Update(snapshot.Completed, ok, c.Experiment.StableFor.Value()) {
			p.TopologyStable = milestone(snapshot.Completed, *r.FaultRequestedAt)
			p.StableRedundancy = p.TopologyStable
			r.Recovery.TopologyConverged = p.TopologyStable
			return
		}
		if delay(ctx, c.Experiment.SampleInterval.Value()) != nil {
			break
		}
	}
	p.RestoreTimedOut = errors.Is(ctx.Err(), context.DeadlineExceeded)
	r.Recovery.TopologyRecovered, r.Recovery.RedundancyRestored = false, false
	if p.RestoreTimedOut {
		r.Errors = append(r.Errors, "original six-node topology did not sustain the stability interval within the restoration budget")
	}
}

func observePersistenceDiagnostics(ctx context.Context, c config.Config, r *report.Report, collect func(context.Context) report.Sample) {
	combined, service, policy := Stability{}, Stability{}, Stability{}
	var postUntil time.Time
	for ctx.Err() == nil {
		s := collect(ctx)
		s.Phase = "persistence_diagnostics"
		s.TopologyOK, s.Reasons = cluster.ReplicaRecoveryWithBaseline(s.Snapshot, *r.Selection, *r.Initial)
		r.Final = &s.Snapshot
		lastObservation := s.Probe.At
		if s.Snapshot.Completed.After(lastObservation) {
			lastObservation = s.Snapshot.Completed
		}
		if s.WritePolicyProbe != nil && s.WritePolicyProbe.At.After(lastObservation) {
			lastObservation = s.WritePolicyProbe.At
		}
		if persistenceSampleExpired(ctx, lastObservation) {
			s.Excluded = true
			s.Reasons = append(s.Reasons, "sample crossed the diagnostic deadline; excluded from client recovery decisions")
			appendSample(r, s)
			break
		}
		at := s.Probe.At
		updatePersistenceDimensions(r, s)
		copy := *r.Dimensions
		s.Dimensions = &copy
		appendSample(r, s)
		updatePolicyRecovery(r, s, &policy, c.Experiment.StableFor.Value())
		r.Recovery.TopologyRecovered, r.Recovery.RedundancyRestored = s.TopologyOK, s.TopologyOK
		r.Recovery.ClientRecovered = service.Update(at, s.Probe.Successful, c.Experiment.StableFor.Value())
		r.Recovery.CombinedRecovered = combined.Update(at, s.TopologyOK && s.Probe.Successful, c.Experiment.StableFor.Value())
		if s.Probe.Successful && r.Recovery.ClientFirstSuccess == nil {
			r.Recovery.ClientFirstSuccess = milestone(at, *r.FaultRequestedAt)
		}
		if r.Recovery.ClientRecovered && r.Recovery.ClientStable == nil {
			r.Recovery.ClientStable = milestone(at, *r.FaultRequestedAt)
		}
		policyOK := !requiresAcknowledgementPolicy(c.Workload.WritePolicy) || (r.Recovery.WritePolicyRecovered != nil && *r.Recovery.WritePolicyRecovered)
		if r.Recovery.CombinedRecovered && policyOK && postUntil.IsZero() {
			r.Recovery.CombinedStable = milestone(at, *r.FaultRequestedAt)
			r.PersistenceFault.DiagnosticServiceStable = r.Recovery.CombinedStable
			postUntil = at.Add(c.Experiment.PostRecovery.Value())
		}
		if !postUntil.IsZero() && !at.Before(postUntil) {
			return // final current flags preserve any post-stability regression
		}
		if delay(ctx, c.Experiment.SampleInterval.Value()) != nil {
			break
		}
	}
	r.Recovery.TimedOut = errors.Is(ctx.Err(), context.DeadlineExceeded)
	r.Recovery.ClientRecovered, r.Recovery.CombinedRecovered = false, false
	if requiresAcknowledgementPolicy(c.Workload.WritePolicy) {
		unavailable := false
		r.Recovery.WritePolicyRecovered = &unavailable
	}
}

func updatePersistenceDimensions(r *report.Report, s report.Sample) {
	// Share the independently checked membership/usable-owner logic without
	// introducing a single-node lifecycle fault into this all-process report.
	local := report.Report{Config: r.Config, LifecycleFault: &report.LifecycleFault{Target: r.Selection.Target, TargetRedisID: r.Selection.TargetID}, Recovery: r.Recovery}
	updateLifecycleDimensions(&local, s, *r.Initial)
	d := *local.Dimensions
	d.TargetRole = ""
	allLive := len(s.Snapshot.Views) == 6
	for _, v := range s.Snapshot.Views {
		allLive = allLive && v.Reachable && len(v.Errors) == 0
	}
	d.TargetProcessLive = &allLive
	ok, _ := cluster.ReplicaRecoveryWithBaseline(s.Snapshot, *r.Selection, *r.Initial)
	d.RedundancyRestored = &ok
	if r.Dimensions != nil {
		d.DataVerified, d.PersistenceRecovered = r.Dimensions.DataVerified, r.Dimensions.PersistenceRecovered
	}
	r.Dimensions = &d
}

func updatePersistenceVerdict(r *report.Report) {
	p := r.PersistenceFault
	if r.Dimensions == nil {
		r.Dimensions = &report.RecoveryDimensions{}
	}
	if len(p.Prepared) != 6 || len(p.Restarts) != 6 || len(p.Restored) != 6 {
		return
	}
	for _, sample := range p.Prepared {
		if sample.Error != "" || persistenceRuntimeEvidence(sample.Evidence, p.Mode) != nil {
			return
		}
	}
	for _, sample := range p.Restored {
		if sample.Error != "" || persistenceRuntimeEvidence(sample.Evidence, p.Mode) != nil || !persistenceRestoredProcess(r, sample.Node, sample.Evidence) {
			return
		}
	}
	if p.Mode == "memory" {
		disabled := false
		r.Dimensions.PersistenceRecovered = &disabled
		p.PersistenceRecoveryObserved, p.PersistenceRecoveryAt = nil, nil
		return
	}
	if r.Recovery.DataVerified && !p.RestoreTimedOut && p.StableRedundancy != nil {
		preserved := true
		p.PersistenceRecoveryObserved, r.Dimensions.PersistenceRecovered = &preserved, &preserved
		p.PersistenceRecoveryAt = milestone(r.Validation.CompletedAt, *r.FaultRequestedAt)
	} else if r.Validation.Complete && r.Validation.Expected > 0 && r.Validation.Checked == r.Validation.Expected && r.Validation.Unreadable+r.Validation.Unverified == 0 && r.Validation.Missing+r.Validation.Corrupt > 0 {
		preserved := false
		p.PersistenceRecoveryObserved, r.Dimensions.PersistenceRecovered = &preserved, &preserved
	}
}

func persistenceNewProcess(r *report.Report, n environment.Node, x *environment.RestartReceipt) bool {
	for _, prepared := range r.PersistenceFault.Prepared {
		if prepared.Node == n && prepared.Error == "" && prepared.OriginalRedisID == n.RedisClusterID {
			return prepared.Evidence.RedisRunID != "" && x.After.RedisRunID != prepared.Evidence.RedisRunID && x.Before.ContainerStartedAt == prepared.Evidence.ContainerStartedAt && x.Before.ContainerRestartPolicy == "no"
		}
	}
	return false
}

func persistenceRestoredProcess(r *report.Report, n environment.Node, e environment.ProcessEvidence) bool {
	if r.PersistenceFault.RestartCompletedAt == nil || e.StartedAt.Before(*r.PersistenceFault.RestartCompletedAt) || e.RedisClusterID != n.RedisClusterID {
		return false
	}
	for _, x := range r.PersistenceFault.Restarts {
		if x != nil && x.Node == n && x.After.RedisRunID == e.RedisRunID && x.After.ContainerStartedAt == e.ContainerStartedAt && x.After.ContainerPID == e.ContainerPID && x.After.RedisPID == e.RedisPID {
			return true
		}
	}
	return false
}

func persistencePreparationBusy(e environment.ProcessEvidence) bool {
	p, repl := cluster.ParseInfo(e.RawPersistence), cluster.ParseInfo(e.RawReplication)
	pending, _ := strconv.ParseUint(p["aof_pending_bio_fsync"], 10, 64)
	return p["loading"] == "1" || p["rdb_bgsave_in_progress"] == "1" || p["aof_rewrite_in_progress"] == "1" || p["aof_rewrite_scheduled"] == "1" || pending > 0 || p["aof_pending_rewrite"] == "1" || repl["master_sync_in_progress"] == "1"
}

func persistencePreparationQuiet(e environment.ProcessEvidence) error {
	p, repl := cluster.ParseInfo(e.RawPersistence), cluster.ParseInfo(e.RawReplication)
	for _, key := range []string{"loading", "rdb_bgsave_in_progress", "aof_rewrite_in_progress", "aof_rewrite_scheduled"} {
		if p[key] != "0" {
			return fmt.Errorf("quiet preparation indicator %s is unavailable or not zero", key)
		}
	}
	if p["aof_enabled"] == "1" && (p["aof_pending_bio_fsync"] != "0" || p["aof_pending_rewrite"] != "0") {
		return errors.New("quiet preparation lacks zero pending AOF fsync/rewrite observations")
	}
	switch repl["role"] {
	case "master":
		if repl["connected_slaves"] != "1" {
			return errors.New("prepared primary has no observed synchronized replica")
		}
	case "slave":
		if repl["master_link_status"] != "up" || repl["master_sync_in_progress"] != "0" {
			return errors.New("prepared replica is not observed linked and synchronized")
		}
	default:
		return errors.New("prepared replication role is unobserved")
	}
	return nil
}

// CONFIG GET is retained as alternating raw key/value lines. Empty save values
// must survive parsing; INFO persistence uses its independent key:value form.
func persistenceRuntimeEvidence(e environment.ProcessEvidence, mode string) error {
	if e.ObservedAt.IsZero() || !e.ContainerRunning || e.ContainerPID <= 0 || !e.RedisPONG || e.RedisPID <= 0 || e.RedisRunID == "" || e.RedisClusterID == "" || e.ContainerRestartPolicy != "no" {
		return errors.New("live managed process evidence is incomplete")
	}
	lines := strings.Split(strings.ReplaceAll(e.RawPersistenceConfig, "\r\n", "\n"), "\n")
	if len(lines)%2 != 0 {
		return errors.New("persistence CONFIG GET is incomplete")
	}
	values := map[string]string{}
	for i := 0; i < len(lines); i += 2 {
		if lines[i] == "" {
			return errors.New("persistence CONFIG GET has an empty key")
		}
		if _, duplicate := values[lines[i]]; duplicate {
			return errors.New("persistence CONFIG GET has duplicate keys")
		}
		values[lines[i]] = lines[i+1]
	}
	expected := map[string]string{"save": "", "repl-diskless-load": "swapdb", "repl-diskless-sync": "yes", "auto-aof-rewrite-percentage": "0", "maxmemory": "134217728", "maxmemory-policy": "noeviction", "aof-use-rdb-preamble": "no", "no-appendfsync-on-rewrite": "no", "aof-load-truncated": "yes", "rdb-del-sync-files": "yes"}
	appendonly, aofEnabled := "yes", "1"
	switch mode {
	case "memory":
		appendonly, aofEnabled = "no", "0"
	case "aof-everysec":
		expected["appendfsync"] = "everysec"
		expected["aof-use-rdb-preamble"] = "yes"
	case "aof-always":
		expected["appendfsync"] = "always"
		expected["aof-use-rdb-preamble"] = "yes"
	default:
		return errors.New("unrecognized persistence mode")
	}
	expected["appendonly"] = appendonly
	if len(values) != 12 || (values["appendfsync"] != "no" && values["appendfsync"] != "everysec" && values["appendfsync"] != "always") {
		return errors.New("persistence CONFIG GET contains unexpected fields or fsync policy")
	}
	for key, value := range expected {
		if observed, exists := values[key]; !exists || observed != value {
			return fmt.Errorf("runtime persistence %s=%q does not match expected %q", key, observed, value)
		}
	}
	info := cluster.ParseInfo(e.RawPersistence)
	if info["loading"] != "0" || info["aof_enabled"] != aofEnabled || (aofEnabled == "1" && (info["aof_last_write_status"] != "ok" || info["aof_last_bgrewrite_status"] != "ok")) {
		return errors.New("runtime persistence is loading, unhealthy, or does not match configured mode")
	}
	files := strings.ToLower(e.RawDataFiles)
	if !strings.Contains(files, "/data/nodes.conf") {
		return errors.New("managed data-file inventory is unobserved")
	}
	if mode == "memory" && (strings.Contains(files, ".rdb") || strings.Contains(files, ".aof") || strings.Contains(files, "appendonly")) {
		return errors.New("memory control unexpectedly contains RDB or AOF value persistence files")
	}
	if mode != "memory" && (!strings.Contains(files, ".aof") || !strings.Contains(files, ".manifest")) {
		return errors.New("AOF runtime has no observed AOF data and manifest files")
	}
	return nil
}

func persistenceError(r *report.Report, err error) error {
	if r.PersistenceFault != nil {
		r.PersistenceFault.Errors = append(r.PersistenceFault.Errors, err.Error())
	}
	return err
}

func milliseconds(d time.Duration) float64 { return float64(d) / float64(time.Millisecond) }

func persistenceSampleExpired(ctx context.Context, at time.Time) bool {
	deadline, bounded := ctx.Deadline()
	return ctx.Err() != nil || at.IsZero() || (bounded && !at.Before(deadline))
}

func persistenceMethodology() []string {
	return []string{
		"Quiet administrative process, CONFIG, persistence INFO, keyspace and bounded data-file preparation occurs before baseline traffic; it issues no SAVE, BGSAVE, BGREWRITEAOF or explicit durability barrier.",
		"The workload fully stops and joins in-flight acknowledgements before freezing its immutable ledger and counters. All later topology observations are read-only until every frozen acknowledged value has been checked. SET-error outcomes remain ambiguous, excluded from acknowledgement guarantees.",
		"Original replicas are SIGKILLed before original primaries. Every fixed owned container must produce ErrProcessStopped administrative evidence with restart policy no before any restart. Original primaries restart before replicas in the same containers and retained volumes; original Redis Cluster identities and slot ownership must survive.",
		"The model is sequential process SIGKILL while Docker VM, host and storage remain running, followed by same-container starts. It does not simulate power loss, host/page-cache loss or disk loss. Drain, kill and inspection intervals are reported because they permit background fsync progress.",
		"Topology restoration requires all six original identities, endpoints, primary slot owners and replica assignments synchronized for stable_for, without forced Redis topology commands. Physical receipts retain early raw evidence, which can be transient during full sync. Six fresh quiet administrative runtime/configuration/file observations after topology stability prove the final persistence mode on those same restarted processes, within the shared repair budget and before frozen-value validation. Samples crossing restoration or diagnostic deadlines are retained but excluded from successful milestones.",
		"Frozen-ledger GET validation happens exactly once before diagnostic SET/GET or configured replica-acknowledgement probes. Process/slot readiness and later diagnostic service are independent from frozen-value preservation; missing or corrupt values retain a partial data verdict even after service recovers.",
		"Memory mode has RDB/AOF disabled, rdb-del-sync-files=yes to remove replication-created RDB files, and observed absence of value persistence files. Its persistence-recovery event is not applicable and remains absent; its independent persistence dimension is false. AOF modes use the RDB base preamble required by the pinned image's bundled modules plus incremental AOF records. AOF recovery requires actual matching healthy runtime configuration, loading=0 and complete frozen acknowledged-value preservation; this observation establishes no loss-free or power-loss guarantee.",
		"No workload demand is offered after draining, so absent fault/recovery traffic does not measure zero availability or throughput. Post-validation probes are separate diagnostics and never enter workload metrics or the frozen ledger. Poll intervals follow sample completion; observations are sampled and non-atomic.",
	}
}
