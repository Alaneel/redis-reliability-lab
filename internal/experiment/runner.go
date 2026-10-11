// Package experiment coordinates the MVP lifecycle using independently testable
// environment, topology, workload, metrics and reporting components.
package experiment

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/cluster"
	"github.com/Alaneel/redis-reliability-lab/internal/config"
	"github.com/Alaneel/redis-reliability-lab/internal/environment"
	"github.com/Alaneel/redis-reliability-lab/internal/report"
	"github.com/Alaneel/redis-reliability-lab/internal/workload"
	"github.com/redis/go-redis/v9"
)

type EnvironmentManager interface {
	Provision(context.Context, string, environment.Config) (*environment.Environment, error)
	Kill(context.Context, *environment.Environment, environment.Node) error
	Cleanup(context.Context, *environment.Environment) error
}
type NetworkFaultManager interface {
	Isolate(context.Context, *environment.Environment, environment.Node) (*environment.IsolationReceipt, error)
	Restore(context.Context, *environment.Environment, *environment.IsolationReceipt) error
	ProbeProcess(context.Context, *environment.Environment, environment.Node) (environment.ProcessEvidence, error)
}
type RestartManager interface {
	Restart(context.Context, *environment.Environment, environment.Node) (*environment.RestartReceipt, error)
}
type ProcessProber interface {
	ProbeProcess(context.Context, *environment.Environment, environment.Node) (environment.ProcessEvidence, error)
}
type Runner struct {
	Manager EnvironmentManager
	Logger  *slog.Logger
}

func RunID() (string, error) {
	var data [12]byte
	if _, err := rand.Read(data[:]); err != nil {
		return "", err
	}
	return hex.EncodeToString(data[:]), nil
}

// Run always attempts independent, bounded cleanup before writing reports.
// Report errors and cleanup failures are returned even if recovery succeeded.
func (runner Runner) Run(ctx context.Context, c config.Config, outputRoot string) (r *report.Report, runErr error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	if runner.Manager == nil {
		return nil, errors.New("environment manager is required")
	}
	networkManager, networkCapable := runner.Manager.(NetworkFaultManager)
	networkIsolation := c.Experiment.Fault == "primary-isolate" || c.Experiment.Fault == "replica-isolate"
	replicaFault := c.Experiment.Fault == "replica-kill" || c.Experiment.Fault == "replica-isolate"
	lifecycleRepair := replicaFault || c.Experiment.Repair
	persistenceFault := c.Experiment.Fault == "cluster-restart"
	restartManager, restartCapable := runner.Manager.(RestartManager)
	processProber, processCapable := runner.Manager.(ProcessProber)
	if networkIsolation && !networkCapable {
		return nil, errors.New("network isolation requires a network fault manager")
	}
	if lifecycleRepair && !networkIsolation && !restartCapable {
		return nil, errors.New("replica crash or primary repair requires a restart manager")
	}
	if persistenceFault && (!restartCapable || !processCapable) {
		return nil, errors.New("cluster-restart requires owned restart and process observation capabilities")
	}
	if c.Environment.Persistence != "memory" && !processCapable {
		return nil, errors.New("AOF experiments require owned process observation for actual quiet runtime preparation")
	}
	id, err := RunID()
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(outputRoot, 0755); err != nil {
		return nil, err
	}
	dir := filepath.Join(outputRoot, id)
	if err := os.Mkdir(dir, 0755); err != nil {
		return nil, err
	}
	started := time.Now()
	effective, _ := json.Marshal(c)
	hash := sha256.Sum256(effective)
	r = &report.Report{SchemaVersion: 2, RunID: id, Outcome: "failed", StartedAt: started, Config: c, ConfigSHA256: hex.EncodeToString(hash[:]), Host: map[string]string{"os": runtime.GOOS, "architecture": runtime.GOARCH, "go_version": runtime.Version(), "logical_cpus": fmt.Sprint(runtime.NumCPU())}, Samples: []report.Sample{}, Methodology: methodology()}
	if networkIsolation {
		r.NetworkFault = &report.NetworkFault{ProcessSamples: []report.ProcessSample{}}
		r.Methodology = networkMethodology()
	}
	if lifecycleRepair {
		r.LifecycleFault = &report.LifecycleFault{ProcessSamples: []report.ProcessSample{}}
		r.Methodology = lifecycleMethodology(c)
	}
	if lifecycleRepair || c.Environment.Persistence != "memory" {
		r.SchemaVersion = report.LifecycleSchemaVersion
	}
	if persistenceFault {
		r.SchemaVersion = report.LifecycleSchemaVersion
		r.PersistenceFault = &report.PersistenceFault{Mode: c.Environment.Persistence, FailureModel: PersistenceFailureModel, Prepared: []report.NodeProcessSample{}, Restored: []report.NodeProcessSample{}, Stopped: []report.NodeProcessSample{}, KillActions: []report.NodeAction{}, Restarts: []*environment.RestartReceipt{}, RestartActions: []report.NodeAction{}}
		r.Methodology = persistenceMethodology()
	}
	lifecycle := NewLifecycle(started)
	logger := runner.Logger
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	}
	logger = logger.With("run_id", id)
	move := func(next Stage) {
		if err := lifecycle.Move(next, time.Now()); err != nil {
			panic(err)
		}
		logger.Info("experiment stage", "stage", next)
	}
	var observer *cluster.Observer
	var client *redis.ClusterClient
	var traffic *workload.Runner
	var writer *ManagedWriter
	defer func() {
		if traffic != nil {
			traffic.Stop()
			r.Metrics = traffic.Snapshot()
			r.WorkloadStats = traffic.Stats()
		}
		// Cancellation or an ambiguous disconnect result must not skip restoration.
		// Cleanup independently accepts the owned detached container if reconnect fails.
		if networkIsolation && r.NetworkFault.Receipt != nil && !r.NetworkFault.ConnectivityRestored {
			restoreCtx, cancel := context.WithTimeout(context.Background(), c.Timeouts.Restoration.Value())
			err := restoreNetwork(restoreCtx, networkManager, r)
			cancel()
			if err != nil {
				runErr = errors.Join(runErr, fmt.Errorf("deferred network restoration: %w", err))
				r.Outcome = "failed"
			}
		}
		if writer != nil {
			if err := writer.Close(); err != nil {
				runErr = errors.Join(runErr, fmt.Errorf("close write executor: %w", err))
			}
		}
		if client != nil {
			if err := client.Close(); err != nil {
				runErr = errors.Join(runErr, fmt.Errorf("close workload client: %w", err))
			}
		}
		if observer != nil {
			if err := observer.Close(); err != nil {
				runErr = errors.Join(runErr, err)
			}
		}
		move(Cleaning)
		if r.Environment != nil {
			r.Cleanup.Attempted = true
			cleanupCtx, cancel := context.WithTimeout(context.Background(), c.Timeouts.Cleanup.Value())
			err := runner.Manager.Cleanup(cleanupCtx, r.Environment)
			cancel()
			if err != nil {
				r.Cleanup.Error = err.Error()
				runErr = errors.Join(runErr, fmt.Errorf("cleanup failed; use redis-lab cleanup %s: %w", id, err))
				r.Outcome = "failed"
			} else {
				r.Cleanup.Successful = true
			}
		}
		if ctx.Err() != nil {
			r.Outcome = "aborted"
			if !errors.Is(runErr, ctx.Err()) {
				runErr = errors.Join(runErr, ctx.Err())
			}
		}
		if runErr != nil {
			r.Errors = append(r.Errors, runErr.Error())
		}
		move(Complete)
		r.FinishedAt = time.Now()
		for _, e := range lifecycle.Transitions {
			r.Events = append(r.Events, report.Event{Stage: string(e.Stage), At: e.At})
		}
		if err := report.Write(dir, *r); err != nil {
			runErr = errors.Join(runErr, fmt.Errorf("write report: %w", err))
		}
		logger.Info("experiment finished", "outcome", r.Outcome, "reports", dir, "cleanup_verified", r.Cleanup.Successful)
	}()
	move(Provisioning)
	provisionCtx, cancel := context.WithTimeout(ctx, c.Timeouts.Provision.Value())
	r.Environment, err = runner.Manager.Provision(provisionCtx, id, environment.Config{Image: c.Environment.RedisImage, NodeTimeout: c.Environment.NodeTimeout.Value(), NetworkIsolation: networkIsolation, Persistence: c.Environment.Persistence, RetainData: c.RetainData()})
	cancel()
	if err != nil {
		return r, fmt.Errorf("provision: %w", err)
	}
	if r.Environment == nil || len(r.Environment.Nodes) != 6 {
		return r, errors.New("provisioner did not return six managed nodes")
	}
	nodes := make([]cluster.Node, len(r.Environment.Nodes))
	for i, n := range r.Environment.Nodes {
		nodes[i] = cluster.Node{ID: n.ID, Addr: n.Addr, InternalAddr: n.InternalAddr}
	}
	observer = cluster.New(nodes, min(c.Workload.OperationTimeout.Value(), time.Second))
	move(Readiness)
	readinessCtx, cancel := context.WithTimeout(ctx, c.Timeouts.Readiness.Value())
	initial, err := waitReady(readinessCtx, observer, c.Experiment.SampleInterval.Value())
	cancel()
	r.Initial = &initial
	if err != nil {
		return r, err
	}
	selected, err := cluster.SelectPrimary(initial)
	if err != nil {
		return r, err
	}
	r.Selection = &selected
	if persistenceFault || c.Environment.Persistence != "memory" {
		if err := preparePersistence(ctx, c, r, processProber); err != nil {
			return r, err
		}
	}
	client, err = NewClient(r.Environment, c.Workload.OperationTimeout.Value())
	if err != nil {
		return r, err
	}
	r.Environment.Metadata["client_protocol"] = "RESP2"
	r.Environment.Metadata["client_topology_refresh_interval"] = "1s"
	r.Environment.Metadata["client_dial_attempts"] = "1"
	r.Environment.Metadata["client_max_redirects"] = "8"
	r.Environment.Metadata["client_node_retries"] = "disabled"
	r.Environment.Metadata["client_pool_size_per_node"] = "16"
	w := c.Workload
	writer, err = NewWriter(client, r.Environment, w.OperationTimeout.Value(), w.WritePolicy, w.WaitReplicas, w.WaitTimeout.Value())
	if err != nil {
		return r, err
	}
	r.Environment.Metadata["write_connection_affinity"] = "SET and optional WAIT/WAITAOF on one sticky physical connection"
	r.Environment.Metadata["write_retries"] = "only explicit MOVED/ASK before SET acknowledgement, bounded to 8; no transport or WAIT/WAITAOF retries"
	traffic, err = workload.NewWithWriter(client, writer, workload.Config{WritePolicy: w.WritePolicy, Workers: w.Workers, RatePerSecond: w.RatePerSecond, MaxWrites: w.MaxWrites, ReadPercent: w.ReadPercent, ValueBytes: w.ValueBytes, Seed: w.Seed, OperationTimeout: w.OperationTimeout.Value()}, id)
	if err != nil {
		return r, err
	}
	// Compatibility is exercised on an empty workload ledger before baseline.
	// In particular, WAITAOF here cannot fsync writes from the durability trial.
	if c.Environment.Persistence != "memory" {
		r.Methodology = append(r.Methodology, "All AOF modes and write policies use six-node quiet runtime preparation followed by exactly three all-shard policy SET/GET warmup writes before baseline starts. Warmup writes are excluded from the workload ledger and metrics; no WAITAOF compatibility command is issued after baseline to alter crash-trial durability.")
		keys, probeErr := ProbeKeys(initial)
		if probeErr != nil {
			return r, probeErr
		}
		policy := ProbeAcknowledgementPolicy(ctx, writer, client, keys, id, w.WritePolicy, w.WaitReplicas, w.OperationTimeout.Value())
		r.BaselineWritePolicyProbe = &policy
		if !policy.Successful {
			return r, fmt.Errorf("pre-baseline write-policy compatibility probe rejected: %v", policy.Errors)
		}
	}
	if requiresAcknowledgementPolicy(w.WritePolicy) {
		notRecovered := false
		r.Recovery.WritePolicyRecovered = &notRecovered
		r.Methodology = append(r.Methodology, "Configured acknowledgement policy is a separate current condition, measured through all-three-shard SET plus WAIT/WAITAOF on one physical connection. A stable policy milestone is historical; a later unsuccessful probe resets the current condition.")
	}
	move(Baseline)
	if err = traffic.Start(ctx); err != nil {
		return r, err
	}
	traffic.SetPhase("baseline", time.Now())
	if err = delay(ctx, c.Experiment.Baseline.Value()); err != nil {
		return r, err
	}
	baseline := traffic.Snapshot().Phases["baseline"]
	if baseline.Attempts == 0 || baseline.SuccessRate < c.Experiment.MinimumBaselineSuccess {
		return r, fmt.Errorf("baseline rejected: %d attempts, %.4f success rate (required %.4f)", baseline.Attempts, baseline.SuccessRate, c.Experiment.MinimumBaselineSuccess)
	}
	// Readiness is checked again immediately before selecting the destructive target.
	initial = observer.Observe(ctx)
	r.Initial = &initial
	if ready, reasons := cluster.Ready(initial); !ready {
		return r, fmt.Errorf("pre-fault health check: %v", reasons)
	}
	selected, err = cluster.SelectPrimary(initial)
	if err != nil {
		return r, err
	}
	r.Selection = &selected
	probeKeys, err := ProbeKeys(initial)
	if err != nil {
		return r, err
	}
	if persistenceFault {
		return r, runPersistenceExperiment(ctx, c, r, runner.Manager, processProber, restartManager, observer, traffic, client, writer, probeKeys, move)
	}
	move(Injecting)
	faultAt := time.Now()
	r.FaultRequestedAt = &faultAt
	traffic.SetPhase("fault", faultAt)
	var target environment.Node
	selectedTarget := selected.Target
	if replicaFault {
		selectedTarget = selected.Replica
	}
	for _, n := range r.Environment.Nodes {
		if n.ID == selectedTarget.ID {
			target = n
		}
	}
	if target.ID == "" {
		return r, errors.New("selected Redis fault target is not in the managed environment")
	}
	if lifecycleRepair {
		l := r.LifecycleFault
		l.Target, l.TargetRedisID, l.TargetKind = selectedTarget, selected.TargetID, "primary"
		if replicaFault {
			l.TargetRedisID, l.TargetKind = selected.ReplicaID, "replica"
		}
	}
	if networkIsolation {
		faultCtx, faultCancel := context.WithTimeout(ctx, c.Timeouts.Restoration.Value())
		r.NetworkFault.Receipt, err = networkManager.Isolate(faultCtx, r.Environment, target)
		faultCancel()
		if err != nil {
			return r, fmt.Errorf("isolate %s network: %w", target.ID, err)
		}
		if err := validateIsolationReceipt(r.NetworkFault.Receipt, r.Environment, target); err != nil {
			return r, fmt.Errorf("isolation receipt: %w", err)
		}
	} else {
		faultCtx, faultCancel := context.WithTimeout(ctx, c.Timeouts.Restoration.Value())
		err = runner.Manager.Kill(faultCtx, r.Environment, target)
		faultCancel()
		if err != nil {
			return r, fmt.Errorf("inject %s SIGKILL: %w", target.ID, err)
		}
	}
	completed := time.Now()
	r.FaultCompletedAt = &completed
	actionDurationMS := float64(completed.Sub(faultAt)) / float64(time.Millisecond)
	r.FaultActionDurationMS = &actionDurationMS
	move(Observing)
	var combined, service Stability
	var policy Stability

	holdUntil := completed.Add(c.Experiment.IsolationDuration.Value())
	if replicaFault {
		holdUntil = completed.Add(c.Experiment.DownDuration.Value())
	}
	var lastProcessProbe time.Time
	recoveryCtx, cancel := context.WithTimeout(ctx, c.Timeouts.Recovery.Value())
	for {
		if recoveryCtx.Err() != nil {
			break
		}
		if replicaFault && !time.Now().Before(holdUntil) {
			break
		}
		if networkIsolation {
			if !lastProcessProbe.IsZero() && !time.Now().Before(holdUntil) {
				break
			}
			if lastProcessProbe.IsZero() || time.Since(lastProcessProbe) >= 2*time.Second {
				captureProcess(recoveryCtx, networkManager, r, target)
				lastProcessProbe = time.Now()
			}
		}
		sample := sampleRecovery(recoveryCtx, observer, client, probeKeys, id, c.Workload.OperationTimeout.Value(), selected, initial)
		if replicaFault {
			sample.TopologyOK, sample.Reasons = cluster.ReplicaServiceWithBaseline(sample.Snapshot, selected, initial)
		}
		sample.Phase = "fault"
		if requiresAcknowledgementPolicy(w.WritePolicy) {
			p := ProbeAcknowledgementPolicy(recoveryCtx, writer, client, probeKeys, id, w.WritePolicy, w.WaitReplicas, w.OperationTimeout.Value())
			sample.WritePolicyProbe = &p
		}
		if recoverySampleExpired(recoveryCtx, sample) {
			sample.Excluded = true
			sample.Reasons = append(sample.Reasons, "sample crossed the recovery deadline; excluded from recovery decisions")
			appendSample(r, sample)
			r.Final = &sample.Snapshot
			break
		}
		r.Final = &sample.Snapshot
		if replicaFault {
			updateReplicaMilestones(r, sample, selected, faultAt)
		} else {
			updateMilestones(r, sample, selected, faultAt)
		}
		if lifecycleRepair {
			updateLifecycleMilestones(r, sample, selected, initial)
			dimensions := *r.LifecycleFault.Dimensions
			sample.Dimensions = &dimensions
		}
		appendSample(r, sample)
		updatePolicyRecovery(r, sample, &policy, c.Experiment.StableFor.Value())
		at := sample.Probe.At
		serviceStable := service.Update(at, sample.Probe.Successful, c.Experiment.StableFor.Value())
		if serviceStable && r.Recovery.ClientStable == nil {
			r.Recovery.ClientStable = milestone(at, faultAt)
		}
		stable := combined.Update(at, sample.TopologyOK && sample.Probe.Successful, c.Experiment.StableFor.Value())
		r.Recovery.CombinedRecovered = stable
		r.Recovery.TopologyRecovered = sample.TopologyOK
		r.Recovery.ClientRecovered = serviceStable
		if stable {
			if r.Recovery.CombinedStable == nil {
				r.Recovery.CombinedStable = milestone(at, faultAt)
			}
			if !networkIsolation && !replicaFault {
				break
			}
		}
		if delay(recoveryCtx, c.Experiment.SampleInterval.Value()) != nil {
			break
		}
	}
	recoveryExpired := recoveryCtx.Err()
	cancel()
	if ctx.Err() != nil {
		return r, ctx.Err()
	}
	if r.Recovery.CombinedStable == nil {
		r.Recovery.TimedOut = errors.Is(recoveryExpired, context.DeadlineExceeded)
		r.Errors = append(r.Errors, "recovery conditions did not hold for the required stable interval within the fault observation window")
	}
	if lifecycleRepair && (replicaFault || r.Recovery.CombinedStable != nil) {
		move(Repairing)
		traffic.SetPhase("repair", time.Now())
		combined, service, policy = Stability{}, Stability{}, Stability{}
		beginSynchronization := func() {
			move(Reconciling)
			traffic.SetPhase("reconciliation", time.Now())
		}
		err = repairLifecycle(ctx, c, r, restartManager, networkManager, target, observer, client, writer, probeKeys, selected, initial, &combined, &service, &policy, beginSynchronization)
		if err != nil {
			return r, err
		}
		if ctx.Err() != nil {
			return r, ctx.Err()
		}
	} else if networkIsolation {
		move(Restoring)
		restoreCtx, restoreCancel := context.WithTimeout(ctx, c.Timeouts.Restoration.Value())
		err = restoreNetwork(restoreCtx, networkManager, r)
		restoreCancel()
		if err != nil {
			return r, fmt.Errorf("restore primary network: %w", err)
		}
		move(Reconciling)
		traffic.SetPhase("reconciliation", time.Now())
		combined, service, policy = Stability{}, Stability{}, Stability{}
		observeReconciliation(ctx, c, r, observer, client, writer, probeKeys, selected, initial, &combined, &service, &policy)
		captureProcess(ctx, networkManager, r, target)
		if ctx.Err() != nil {
			return r, ctx.Err()
		}
	}
	if r.Recovery.CombinedStable != nil && (!networkIsolation || r.NetworkFault.Reconciled) && (!lifecycleRepair || r.LifecycleFault.StableRedundancy != nil) {
		move(PostRecovery)
		traffic.SetPhase("post_recovery", time.Now())
		postCtx, postCancel := context.WithTimeout(ctx, c.Experiment.PostRecovery.Value())
		for postCtx.Err() == nil {
			sample := sampleRecovery(postCtx, observer, client, probeKeys, id, c.Workload.OperationTimeout.Value(), selected, initial)
			if lifecycleRepair {
				sample.TopologyOK, sample.Reasons = lifecycleRestored(sample.Snapshot, selected, initial, replicaFault)
			} else if networkIsolation {
				sample.TopologyOK, sample.Reasons = cluster.ReconciledWithBaseline(sample.Snapshot, selected, initial)
			}
			sample.Phase = "post_recovery"
			if requiresAcknowledgementPolicy(w.WritePolicy) {
				p := ProbeAcknowledgementPolicy(postCtx, writer, client, probeKeys, id, w.WritePolicy, w.WaitReplicas, w.OperationTimeout.Value())
				sample.WritePolicyProbe = &p
			}
			// A sample cut short by the phase deadline is not a service failure.
			if recoverySampleExpired(postCtx, sample) {
				sample.Excluded = true
				sample.Reasons = append(sample.Reasons, "sample crossed the post-recovery deadline; excluded from recovery decisions")
				appendSample(r, sample)
				break
			}
			r.Final = &sample.Snapshot
			if lifecycleRepair {
				updateLifecycleMilestones(r, sample, selected, initial)
				dimensions := *r.LifecycleFault.Dimensions
				sample.Dimensions = &dimensions
				r.Recovery.RedundancyRestored, _ = cluster.Ready(sample.Snapshot)
			}
			appendSample(r, sample)
			updatePolicyRecovery(r, sample, &policy, c.Experiment.StableFor.Value())
			r.Recovery.TopologyRecovered = sample.TopologyOK
			r.Recovery.ClientRecovered = service.Update(sample.Probe.At, sample.Probe.Successful, c.Experiment.StableFor.Value())
			r.Recovery.CombinedRecovered = combined.Update(sample.Probe.At, sample.TopologyOK && sample.Probe.Successful, c.Experiment.StableFor.Value())
			if networkIsolation {
				r.NetworkFault.Reconciled = sample.TopologyOK && r.Recovery.CombinedRecovered
				r.Recovery.RedundancyRestored, _ = cluster.Ready(sample.Snapshot)
			}
			if delay(postCtx, c.Experiment.SampleInterval.Value()) != nil {
				break
			}
		}
		postCancel()
		if ctx.Err() != nil {
			return r, ctx.Err()
		}
	}
	move(Validating)
	traffic.Stop()
	validationCtx, validationCancel := context.WithTimeout(ctx, c.Timeouts.Validation.Value())
	r.Validation = traffic.Validate(validationCtx, c.Workload.OperationTimeout.Value())
	validationCancel()
	r.WorkloadStats = traffic.Stats()
	finalBaseline := traffic.Snapshot().Phases["baseline"]
	if finalBaseline.Attempts == 0 || finalBaseline.SuccessRate < c.Experiment.MinimumBaselineSuccess {
		r.Recovery.CombinedRecovered = false
		r.Errors = append(r.Errors, fmt.Sprintf("final baseline failed threshold after in-flight operations settled: %.4f success rate", finalBaseline.SuccessRate))
	}
	r.Recovery.DataVerified = r.Validation.Complete && r.Validation.Expected > 0 && r.Validation.Missing == 0 && r.Validation.Corrupt == 0 && r.Validation.Unreadable == 0 && r.WorkloadStats.MissingReads == 0 && r.WorkloadStats.MismatchedReads == 0
	if lifecycleRepair {
		if r.LifecycleFault.Dimensions != nil {
			dimensions := *r.LifecycleFault.Dimensions
			verified := r.Recovery.DataVerified
			dimensions.DataVerified = &verified
			r.LifecycleFault.Dimensions, r.Dimensions = &dimensions, &dimensions
		}
		if !r.LifecycleFault.RepairCompleted || r.LifecycleFault.RepairTimedOut || r.LifecycleFault.StableRedundancy == nil || !r.Recovery.RedundancyRestored {
			r.Recovery.CombinedRecovered = false
		}
	}
	if networkIsolation && !processSurvivedIsolation(r.NetworkFault) {
		r.Recovery.CombinedRecovered = false
		r.Errors = append(r.Errors, "unchanged live Redis process was not verified while its network was disconnected")
	}
	r.Outcome = Outcome(r.Recovery)
	if r.Outcome != "recovered" {
		return r, fmt.Errorf("experiment outcome: %s", r.Outcome)
	}
	return r, nil
}

func Outcome(r report.Recovery) string {
	policyOK := r.WritePolicyRecovered == nil || *r.WritePolicyRecovered
	if r.TopologyRecovered && r.ClientRecovered && r.DataVerified && r.CombinedRecovered && r.CombinedStable != nil && !r.TimedOut && policyOK {
		return "recovered"
	}
	if r.TopologyRecovered || r.ClientRecovered || r.CombinedStable != nil {
		return "partial"
	}
	return "failed"
}

// NewClient refuses all addresses not belonging to this provisioned inventory,
// including unexpected MOVED/ASK targets. Private Docker addresses are translated
// to the exact published loopback ports; Redis's cluster client handles routing.
func NewClient(env *environment.Environment, timeout time.Duration) (*redis.ClusterClient, error) {
	if env == nil || len(env.Nodes) != 6 {
		return nil, errors.New("client requires six managed nodes")
	}
	mapping := map[string]string{}
	ids := map[string]bool{}
	endpoints := map[string]bool{}
	var seeds []string
	for _, n := range env.Nodes {
		host, port, err := net.SplitHostPort(n.Addr)
		if err != nil || host != "127.0.0.1" || port == "" {
			return nil, fmt.Errorf("refusing non-loopback managed address %q", n.Addr)
		}
		portNumber, err := strconv.Atoi(port)
		if err != nil || portNumber < 1 || portNumber > 65535 {
			return nil, errors.New("invalid managed port")
		}
		if n.ID == "" || ids[n.ID] || endpoints[n.Addr] {
			return nil, errors.New("duplicate or empty managed node identity or endpoint")
		}
		ids[n.ID] = true
		endpoints[n.Addr] = true
		if _, _, err := net.SplitHostPort(n.InternalAddr); err != nil {
			return nil, err
		}
		if _, duplicate := mapping[n.InternalAddr]; duplicate {
			return nil, errors.New("duplicate managed internal address")
		}
		mapping[n.InternalAddr] = n.Addr
		mapping[n.Addr] = n.Addr
		seeds = append(seeds, n.Addr)
	}
	dialer := net.Dialer{Timeout: timeout}
	return redis.NewClusterClient(&redis.ClusterOptions{Addrs: seeds, Protocol: 2, MaxRedirects: 8, MaxRetries: -1, DialerRetries: 1, ClusterStateReloadInterval: time.Second, ContextTimeoutEnabled: true, DialTimeout: timeout, ReadTimeout: timeout, WriteTimeout: timeout, PoolTimeout: timeout, PoolSize: 16,
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			mapped, ok := mapping[addr]
			if !ok {
				return nil, fmt.Errorf("refusing unmanaged Redis redirection target %q", addr)
			}
			return dialer.DialContext(ctx, "tcp", mapped)
		},
	}), nil
}

func ProbeKeys(initial cluster.Snapshot) ([]string, error) {
	if len(initial.Views) == 0 {
		return nil, errors.New("no initial topology")
	}
	owners := initial.Views[0].Coverage.Owners
	ids := map[string]bool{}
	for _, o := range owners {
		ids[o.RedisID] = true
	}
	var sorted []string
	for id := range ids {
		sorted = append(sorted, id)
	}
	sort.Strings(sorted)
	var keys []string
	for _, id := range sorted {
		found := false
		for candidate := 0; candidate < 100000; candidate++ {
			key := fmt.Sprintf("redis-lab-probe-{%d}", candidate)
			slot := int(workload.HashSlot(key))
			for _, o := range owners {
				if o.RedisID == id && slot >= o.Start && slot <= o.End {
					keys = append(keys, key)
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("cannot find probe key for primary %s", id)
		}
	}
	if len(keys) != 3 {
		return nil, errors.New("probes require three initial primary owners")
	}
	return keys, nil
}

func ProbeAll(ctx context.Context, client workload.Client, keys []string, value string, timeout time.Duration) report.Probe {
	p := report.Probe{}
	results := make([]error, len(keys))
	writes, reads, readAttempted := make([]bool, len(keys)), make([]bool, len(keys)), make([]bool, len(keys))
	var wg sync.WaitGroup
	for i, key := range keys {
		wg.Add(1)
		go func(i int, key string) {
			defer wg.Done()
			opCtx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			if err := client.Set(opCtx, key, value, 0).Err(); err != nil {
				results[i] = err
				return
			}
			writes[i], readAttempted[i] = true, true
			got, err := client.Get(opCtx, key).Result()
			if err != nil {
				results[i] = err
			} else if got != value {
				results[i] = errors.New("probe value mismatch")
			} else {
				reads[i] = true
			}
		}(i, key)
	}
	wg.Wait()
	p.At = time.Now()
	p.Successful = len(keys) == 3
	writeOK, readOK, allReadAttempted := len(keys) == 3, len(keys) == 3, len(keys) == 3
	for i := range keys {
		writeOK, readOK = writeOK && writes[i], readOK && reads[i]
		allReadAttempted = allReadAttempted && readAttempted[i]
	}
	if len(keys) == 3 {
		p.WriteSuccessful = &writeOK
		if allReadAttempted {
			p.ReadSuccessful = &readOK
		}
	}
	for i, err := range results {
		if err != nil {
			p.Successful = false
			p.Errors = append(p.Errors, fmt.Sprintf("shard %d: %v", i, err))
		}
	}
	return p
}

func sampleRecovery(ctx context.Context, o *cluster.Observer, client workload.Client, keys []string, id string, timeout time.Duration, selection cluster.Selection, initial cluster.Snapshot) report.Sample {
	s := o.Observe(ctx)
	ok, reasons := cluster.RecoveryWithBaseline(s, selection, initial)
	p := ProbeAll(ctx, client, keys, id+":"+fmt.Sprint(time.Now().UnixNano()), timeout)
	return report.Sample{Snapshot: s, Probe: p, TopologyOK: ok, Reasons: reasons}
}
func waitReady(ctx context.Context, o *cluster.Observer, interval time.Duration) (cluster.Snapshot, error) {
	var s cluster.Snapshot
	var reasons []string
	for ctx.Err() == nil {
		s = o.Observe(ctx)
		if ok, why := cluster.Ready(s); ok {
			return s, nil
		} else {
			reasons = why
		}
		if delay(ctx, interval) != nil {
			break
		}
	}
	return s, fmt.Errorf("readiness deadline: %w; last observations: %v", ctx.Err(), reasons)
}
func delay(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
func milestone(at, fault time.Time) *report.Milestone {
	return &report.Milestone{At: at, SinceFaultMS: float64(at.Sub(fault)) / float64(time.Millisecond)}
}
func appendSample(r *report.Report, s report.Sample) {
	const limit = 1200
	if len(r.Samples) < limit {
		r.Samples = append(r.Samples, s)
	} else {
		r.Samples[limit-1] = s
		r.DroppedSamples++
	}
}
func updateMilestones(r *report.Report, s report.Sample, selection cluster.Selection, fault time.Time) {
	at := s.Snapshot.Completed
	for _, v := range s.Snapshot.Views {
		for _, n := range v.Nodes {
			if n.ID == selection.TargetID && (n.HasFlag("fail?") || n.HasFlag("fail")) && r.Recovery.FailureDetected == nil {
				r.Recovery.FailureDetected = milestone(at, fault)
			}
		}
	}
	if r.Recovery.ReplicaPromoted == nil {
		for _, v := range s.Snapshot.Views {
			if !v.Reachable || len(v.Errors) > 0 {
				continue
			}
			for _, n := range v.Nodes {
				if n.ID != selection.ReplicaID || !n.Primary() || !n.Healthy() {
					continue
				}
				ownsAll := true
				for _, slot := range selection.Slots {
					owns := false
					for _, o := range v.Coverage.Owners {
						if slot >= o.Start && slot <= o.End && o.RedisID == selection.ReplicaID {
							owns = true
							break
						}
					}
					if !owns {
						ownsAll = false
						break
					}
				}
				if ownsAll {
					r.Recovery.ReplicaPromoted = milestone(at, fault)
					break
				}
			}
		}
	}
	full := true
	live := 0
	reachableOwners := map[string]bool{}
	for _, v := range s.Snapshot.Views {
		if v.Node.ID != selection.Target.ID && v.Reachable && len(v.Errors) == 0 {
			reachableOwners[v.RedisID] = true
		}
	}
	for _, v := range s.Snapshot.Views {
		if v.Node.ID == selection.Target.ID {
			continue
		}
		live++
		if !v.Reachable || len(v.Errors) > 0 || v.ClusterInfo["cluster_state"] != "ok" || !v.Coverage.Full() {
			full = false
		}
		// Immediately after SIGKILL, gossip can still advertise the dead primary
		// as healthy. Assigned slots are not usable coverage when their observed
		// owner is the known killed node or lacks a successful direct observation.
		for _, owner := range v.Coverage.Owners {
			if owner.RedisID == selection.TargetID || !reachableOwners[owner.RedisID] {
				full = false
			}
		}
	}
	if full && live == 5 && r.Recovery.FullCoverage == nil {
		r.Recovery.FullCoverage = milestone(at, fault)
	}
	if s.TopologyOK && r.Recovery.TopologyConverged == nil {
		r.Recovery.TopologyConverged = milestone(at, fault)
	}
	if s.Probe.Successful && r.Recovery.ClientFirstSuccess == nil {
		r.Recovery.ClientFirstSuccess = milestone(s.Probe.At, fault)
	}
}
func methodology() []string {
	return []string{
		"Fault is SIGKILL of the deterministic primary owning the lowest initial slot; the process remains stopped. Initial topology is three primaries and one synchronized replica each.",
		"Time origin is the start of validated fault injection, including Docker ownership inspections before SIGKILL; successful command completion is also recorded. Redis event times are bounded by polling and request deadlines; views within a snapshot are concurrent and not an atomic Redis snapshot.",
		"A sample collects direct CLUSTER INFO/NODES and replication/server INFO from all six managed nodes, then parallel SET/GET probes to all three initial shards. Poll interval is a delay after sample completion; actual timestamps define resolution.",
		"Recovery requires five surviving nodes healthy with exactly one owner per 16384 slots, equal ownership maps, the original replica owning the failed primary slots, unaffected ownership preserved, and all-shard SET/GET probes succeeding for stable_for. Post-recovery probes continue to detect regressions.",
		"Probes are separate from workload metrics and do not guarantee every request succeeded between samples. The workload success rate, latency distributions and error counts report client-visible operations, including terminal errors after cluster redirections.",
		"Workload stops and waits for in-flight requests before verifying every acknowledged immutable write. Missing/corrupt reads during traffic also prevent a recovered outcome. Error-returning writes have unknown outcomes, not proven lost writes.",
		"At most 1200 topology samples are retained (first 1199 and latest); dropped_samples records omitted intermediates. Milestones are retained independently. No process restart or redundancy restoration is attempted.",
	}
}
