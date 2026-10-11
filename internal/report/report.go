// Package report writes experiment evidence as versioned JSON and Markdown.
package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/metrics"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

type Event struct {
	Stage string    `json:"stage"`
	At    time.Time `json:"at"`
}
type Milestone struct {
	At           time.Time `json:"at"`
	SinceFaultMS float64   `json:"since_fault_ms"`
}
type Recovery struct {
	FailureDetected      *Milestone `json:"failure_detected,omitempty"`
	ReplicaPromoted      *Milestone `json:"replica_promoted,omitempty"`
	FullCoverage         *Milestone `json:"full_slot_coverage,omitempty"`
	TopologyConverged    *Milestone `json:"topology_converged,omitempty"`
	ClientFirstSuccess   *Milestone `json:"client_first_all_shard_success,omitempty"`
	ClientStable         *Milestone `json:"client_stable,omitempty"`
	CombinedStable       *Milestone `json:"combined_stable,omitempty"`
	TopologyRecovered    bool       `json:"topology_recovered"`
	ClientRecovered      bool       `json:"client_recovered"`
	CombinedRecovered    bool       `json:"combined_recovered"`
	DataVerified         bool       `json:"data_verified"`
	ProcessRestarted     bool       `json:"process_restarted"`
	RedundancyRestored   bool       `json:"redundancy_restored"`
	TimedOut             bool       `json:"timed_out"`
	WritePolicyRecovered *bool      `json:"write_policy_recovered,omitempty"`
	WritePolicyStable    *Milestone `json:"write_policy_stable,omitempty"`
}
type Probe struct {
	At              time.Time `json:"at"`
	Successful      bool      `json:"successful"`
	ReadSuccessful  *bool     `json:"read_successful,omitempty"`
	WriteSuccessful *bool     `json:"write_successful,omitempty"`
	Errors          []string  `json:"errors,omitempty"`
}
type Sample struct {
	Phase            string              `json:"phase,omitempty"`
	Snapshot         cluster.Snapshot    `json:"snapshot"`
	Probe            Probe               `json:"probe"`
	TopologyOK       bool                `json:"topology_ok"`
	Reasons          []string            `json:"reasons,omitempty"`
	Excluded         bool                `json:"excluded_from_recovery,omitempty"`
	WritePolicyProbe *PolicyProbe        `json:"write_policy_probe,omitempty"`
	Dimensions       *RecoveryDimensions `json:"dimensions,omitempty"`
}

type PolicyProbe struct {
	At                             time.Time `json:"at"`
	Successful                     bool      `json:"successful"`
	Policy                         string    `json:"policy,omitempty"`
	RequestedReplicas              int       `json:"requested_replicas"`
	SetAcknowledged                []bool    `json:"set_acknowledged"`
	Acknowledgements               []int64   `json:"acknowledgements"`
	WaitAttempted                  []bool    `json:"wait_attempted"`
	AcknowledgementObserved        []bool    `json:"acknowledgement_observed"`
	WaitAOFRequestedLocal          int       `json:"waitaof_requested_local,omitempty"`
	WaitAOFRequestedReplicas       int       `json:"waitaof_requested_replicas,omitempty"`
	WaitAOFAttempted               []bool    `json:"waitaof_attempted,omitempty"`
	WaitAOFLocalAcknowledgements   []int64   `json:"waitaof_local_acknowledgements,omitempty"`
	WaitAOFReplicaAcknowledgements []int64   `json:"waitaof_replica_acknowledgements,omitempty"`
	WaitAOFAcknowledgementObserved []bool    `json:"waitaof_acknowledgement_observed,omitempty"`
	Errors                         []string  `json:"errors,omitempty"`
}

type ProcessSample struct {
	Evidence environment.ProcessEvidence `json:"evidence"`
	Error    string                      `json:"error,omitempty"`
}
type NetworkFault struct {
	Receipt                *environment.IsolationReceipt `json:"receipt,omitempty"`
	ProcessSamples         []ProcessSample               `json:"process_samples"`
	RestoreRequestedAt     *time.Time                    `json:"restore_requested_at,omitempty"`
	RestoreCompletedAt     *time.Time                    `json:"restore_completed_at,omitempty"`
	RestoreDurationMS      *float64                      `json:"restore_duration_ms,omitempty"`
	ConnectivityRestored   bool                          `json:"connectivity_restored"`
	Reconciled             bool                          `json:"reconciled"`
	ReconciledAt           *Milestone                    `json:"reconciled_at,omitempty"`
	RedundancyRestoredAt   *Milestone                    `json:"redundancy_restored_at,omitempty"`
	StableReconciliationMS *float64                      `json:"stable_reconciliation_since_restore_completed_ms,omitempty"`
	ReconciliationTimedOut bool                          `json:"reconciliation_timed_out"`
	Errors                 []string                      `json:"errors,omitempty"`
}
type Cleanup struct {
	Attempted  bool   `json:"attempted"`
	Successful bool   `json:"successful"`
	Error      string `json:"error,omitempty"`
}
type Report struct {
	SchemaVersion            int                      `json:"schema_version"`
	RunID                    string                   `json:"run_id"`
	Outcome                  string                   `json:"outcome"`
	StartedAt                time.Time                `json:"started_at"`
	FinishedAt               time.Time                `json:"finished_at"`
	Config                   config.Config            `json:"config"`
	ConfigSHA256             string                   `json:"effective_config_sha256"`
	Host                     map[string]string        `json:"host"`
	Environment              *environment.Environment `json:"environment,omitempty"`
	Events                   []Event                  `json:"events"`
	Selection                *cluster.Selection       `json:"fault_selection,omitempty"`
	FaultRequestedAt         *time.Time               `json:"fault_requested_at,omitempty"`
	FaultCompletedAt         *time.Time               `json:"fault_completed_at,omitempty"`
	FaultActionDurationMS    *float64                 `json:"fault_action_duration_ms,omitempty"`
	NetworkFault             *NetworkFault            `json:"network_fault,omitempty"`
	LifecycleFault           *LifecycleFault          `json:"lifecycle_fault,omitempty"`
	PersistenceFault         *PersistenceFault        `json:"persistence_fault,omitempty"`
	PersistencePrepared      []NodeProcessSample      `json:"persistence_prepared,omitempty"`
	BaselineWritePolicyProbe *PolicyProbe             `json:"baseline_write_policy_probe,omitempty"`
	Dimensions               *RecoveryDimensions      `json:"recovery_dimensions,omitempty"`
	Initial                  *cluster.Snapshot        `json:"initial,omitempty"`
	Final                    *cluster.Snapshot        `json:"final,omitempty"`
	Samples                  []Sample                 `json:"samples"`
	DroppedSamples           int                      `json:"dropped_samples"`
	Recovery                 Recovery                 `json:"recovery"`
	Metrics                  metrics.Snapshot         `json:"metrics"`
	WorkloadStats            workload.Stats           `json:"workload_stats"`
	Validation               workload.Validation      `json:"validation"`
	Cleanup                  Cleanup                  `json:"cleanup"`
	Errors                   []string                 `json:"errors,omitempty"`
	Methodology              []string                 `json:"methodology"`
}

func (r Report) Markdown() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "# Redis Reliability Lab — %s\n\nRun: `%s`  \nOutcome: **%s**  \nStarted: %s  \nFinished: %s\n\n", r.Config.Name, r.RunID, r.Outcome, r.StartedAt.Format(time.RFC3339Nano), r.FinishedAt.Format(time.RFC3339Nano))
	if r.PersistenceFault != nil {
		fmt.Fprintf(&b, "## Recovery\n\n| Condition | Observed |\n|---|---|\n| Original six-node topology recovered | %t |\n| Post-validation diagnostic service sustained | %t |\n| Topology and diagnostic service jointly sustained | %t |\n| Frozen acknowledged writes verified | %t |\n| All managed processes restarted | %t |\n| Original replication redundancy restored | %t |\n| Restoration deadline expired | %t |\n| Diagnostic deadline expired | %t |\n| Cleanup verified | %t |\n\n", r.Recovery.TopologyRecovered, r.Recovery.ClientRecovered, r.Recovery.CombinedRecovered, r.Recovery.DataVerified, r.Recovery.ProcessRestarted, r.Recovery.RedundancyRestored, r.PersistenceFault.RestoreTimedOut, r.Recovery.TimedOut, r.Cleanup.Successful)
	} else {
		fmt.Fprintf(&b, "## Recovery\n\n| Condition | Observed |\n|---|---|\n| Surviving topology converged | %t |\n| All-shard client recovery sustained | %t |\n| Topology and client recovery jointly sustained | %t |\n| All acknowledged writes verified | %t |\n| Target process restarted | %t |\n| Original replication redundancy restored | %t |\n| Recovery deadline expired | %t |\n| Cleanup verified | %t |\n\n", r.Recovery.TopologyRecovered, r.Recovery.ClientRecovered, r.Recovery.CombinedRecovered, r.Recovery.DataVerified, r.Recovery.ProcessRestarted, r.Recovery.RedundancyRestored, r.Recovery.TimedOut, r.Cleanup.Successful)
	}
	if r.LifecycleFault == nil && r.NetworkFault == nil && r.PersistenceFault == nil {
		fmt.Fprintln(&b, "The primary stays stopped. Service recovery can succeed with reduced redundancy. Redis replication is asynchronous; an acknowledged SET is not a durability guarantee. Missing acknowledged writes cause a partial outcome.")
	} else if r.NetworkFault != nil {
		kind := "primary"
		if r.LifecycleFault != nil {
			kind = r.LifecycleFault.TargetKind
		}
		fmt.Fprintf(&b, "The %s remains running during whole-node bidirectional network isolation. Connectivity restoration, topology reconciliation and replication redundancy are evaluated separately. Redis replication is asynchronous; acknowledged SETs are not durability guarantees.\n", kind)
		fmt.Fprintf(&b, "\n## Network restoration\n\n| Condition | Observed |\n|---|---|\n| Original endpoints and process restored | %t |\n| Topology and clients jointly reconciled | %t |\n| Reconciliation deadline expired | %t |\n\n", r.NetworkFault.ConnectivityRestored, r.NetworkFault.Reconciled, r.NetworkFault.ReconciliationTimedOut)
		if r.NetworkFault.RestoreDurationMS != nil {
			fmt.Fprintf(&b, "Restoration action duration: %.3f ms.\n", *r.NetworkFault.RestoreDurationMS)
		}
		if r.NetworkFault.StableReconciliationMS != nil {
			fmt.Fprintf(&b, "Stable reconciliation first observed %.3f ms after restoration completed.\n", *r.NetworkFault.StableReconciliationMS)
		}
		fmt.Fprint(&b, "\n| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |\n|---|---|---|---|---:|---|---|\n")
		for _, s := range r.NetworkFault.ProcessSamples {
			p := s.Evidence
			fmt.Fprintf(&b, "| %s | %t | %t | %t | %d | %s | %s |\n", p.ObservedAt.Format(time.RFC3339Nano), p.RedisPONG, p.ContainerRunning, p.NetworkConnected, p.ContainerPID, p.RedisRunID, s.Error)
		}
		for _, err := range r.NetworkFault.Errors {
			fmt.Fprintf(&b, "\nNetwork fault error: %s\n", err)
		}
	}
	if r.LifecycleFault != nil {
		r.LifecycleFault.markdown(&b)
	}
	if r.PersistenceFault != nil {
		r.PersistenceFault.markdown(&b)
	}
	if r.Dimensions != nil {
		r.Dimensions.markdown(&b)
	} else if r.LifecycleFault != nil && r.LifecycleFault.Dimensions != nil {
		r.LifecycleFault.Dimensions.markdown(&b)
	}
	fmt.Fprintln(&b, "\n## Timing\n\nMilestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.")
	if r.FaultActionDurationMS != nil {
		fmt.Fprintf(&b, "\nFault action completed %.3f ms after orchestration began (measured with the monotonic runtime clock).\n", *r.FaultActionDurationMS)
	}
	if r.PersistenceFault == nil {
		fmt.Fprintln(&b, "\n| Observation | Since orchestration began, ms | Since action completed, ms |\n|---|---:|---:|")
		for _, m := range []struct {
			name  string
			value *Milestone
		}{
			{"Failure flagged in gossip", r.Recovery.FailureDetected}, {"Original replica owns failed primary slots", r.Recovery.ReplicaPromoted}, {"Full usable slot coverage", r.Recovery.FullCoverage}, {"Surviving views converge", r.Recovery.TopologyConverged}, {"First all-shard SET/GET probe succeeds", r.Recovery.ClientFirstSuccess}, {"Client probes sustain recovery", r.Recovery.ClientStable}, {"Topology and clients sustain recovery", r.Recovery.CombinedStable},
		} {
			if m.value != nil {
				afterAction := "unavailable"
				if r.FaultActionDurationMS != nil {
					afterAction = fmt.Sprintf("%.2f", m.value.SinceFaultMS-*r.FaultActionDurationMS)
				}
				fmt.Fprintf(&b, "| %s | %.2f | %s |\n", m.name, m.value.SinceFaultMS, afterAction)
			} else {
				fmt.Fprintf(&b, "| %s | unobserved | unobserved |\n", m.name)
			}
		}
	} else {
		fmt.Fprintln(&b, "\nNo primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.")
	}
	fmt.Fprintln(&b, "\n## Workload\n\n| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |\n|---|---:|---:|---:|---:|---|")
	for _, phase := range []string{"baseline", "fault", "repair", "reconciliation", "post_recovery"} {
		if s, ok := r.Metrics.Phases[phase]; ok {
			if r.PersistenceFault != nil && s.Attempts == 0 {
				continue
			}
			latency := "unobserved"
			if s.SuccessfulLatency.Count > 0 {
				latency = fmt.Sprintf("%.3f / %.3f / %.3f", s.SuccessfulLatency.P50MS, s.SuccessfulLatency.P95MS, s.SuccessfulLatency.P99MS)
			}
			fmt.Fprintf(&b, "| %s | %d | %.2f%% | %d | %.2f | %s |\n", phase, s.Attempts, s.SuccessRate*100, s.Errors, s.Throughput, latency)
		}
	}
	fmt.Fprintf(&b, "\nAcknowledged writes: %d. Ambiguous SET outcomes: %d.\n", r.WorkloadStats.AcknowledgedWrites, r.WorkloadStats.AmbiguousWrites)
	r.persistencePreparationMarkdown(&b)
	r.policyMarkdown(&b)
	if r.WorkloadStats.PacingMethodology != "" {
		fmt.Fprintf(&b, "\nGlobal pacing: offered %d; started %d; completed %d; dropped %d. These counts include all workload phases and exclude validation and probes.\n\n%s\n", r.WorkloadStats.OfferedRequests, r.WorkloadStats.StartedRequests, r.WorkloadStats.CompletedRequests, r.WorkloadStats.DroppedRequests, r.WorkloadStats.PacingMethodology)
	}
	validationAttempted := !r.Validation.StartedAt.IsZero()
	fmt.Fprintf(&b, "\nValidation attempted: %t; expected: %d; checked: %d; missing: %d; corrupt: %d; unreadable: %d; unverified: %d. Validation complete: %t.\n", validationAttempted, r.Validation.Expected, r.Validation.Checked, r.Validation.Missing, r.Validation.Corrupt, r.Validation.Unreadable, r.Validation.Unverified, r.Validation.Complete)
	if !validationAttempted {
		fmt.Fprintf(&b, "Validation was not attempted; %d acknowledged writes remain unchecked.\n", r.WorkloadStats.AcknowledgedWrites)
	}
	r.policyValidationMarkdown(&b)
	fmt.Fprint(&b, "\n### Terminal workload errors\n\n| Phase | Error class | Count |\n|---|---|---:|\n")
	phases := make([]string, 0, len(r.Metrics.Phases))
	for phase := range r.Metrics.Phases {
		phases = append(phases, phase)
	}
	sort.Strings(phases)
	for _, phase := range phases {
		classes := make([]string, 0, len(r.Metrics.Phases[phase].ErrorClasses))
		for class, count := range r.Metrics.Phases[phase].ErrorClasses {
			if count > 0 {
				classes = append(classes, class)
			}
		}
		sort.Strings(classes)
		for _, class := range classes {
			fmt.Fprintf(&b, "| %s | %s | %d |\n", phase, class, r.Metrics.Phases[phase].ErrorClasses[class])
		}
	}
	if r.Cleanup.Error != "" {
		fmt.Fprintf(&b, "\nCleanup error: %s\n", r.Cleanup.Error)
	}
	fmt.Fprintf(&b, "\n%s\n", r.Metrics.Methodology)
	if len(r.Errors) > 0 {
		fmt.Fprint(&b, "\n## Errors\n\n")
		for _, err := range r.Errors {
			fmt.Fprintf(&b, "- %s\n", err)
		}
	}
	fmt.Fprint(&b, "\n## Measurement methodology\n\n")
	for _, m := range r.Methodology {
		fmt.Fprintf(&b, "- %s\n", m)
	}
	fmt.Fprintf(&b, "\nThe JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `%s`.\n", r.ConfigSHA256)
	return b.String()
}

// Write atomically replaces each file after fully writing its content. Each run
// uses its own directory. A failed second write is returned, never hidden.
func Write(dir string, r Report) error {
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}
	if err := atomicWrite(filepath.Join(dir, "report.json"), append(data, '\n')); err != nil {
		return err
	}
	return atomicWrite(filepath.Join(dir, "report.md"), []byte(r.Markdown()))
}
func atomicWrite(path string, data []byte) (err error) {
	f, err := os.CreateTemp(filepath.Dir(path), ".report-*")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if _, err = f.Write(data); err != nil {
		f.Close()
		return err
	}
	if err = f.Sync(); err != nil {
		f.Close()
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), path)
}
