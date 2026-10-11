package report

import (
	"fmt"
	"io"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/cluster"
	"github.com/Alaneel/redis-reliability-lab/internal/environment"
	"github.com/Alaneel/redis-reliability-lab/internal/workload"
)

// PersistenceFailureModel names the measured process failure model. It is shared
// with the runner and saved-evidence gate so reports use the same exact contract.
const PersistenceFailureModel = "sequential SIGKILL of all six Redis processes, same-container restart; host and storage remain running"

// NodeProcessSample binds an administrative observation to the original managed
// node and Redis Cluster identity. Process evidence is distinct from client or
// peer reachability. Expected stopped-process observations retain their status.
type NodeProcessSample struct {
	Node            environment.Node            `json:"node"`
	OriginalRedisID string                      `json:"original_redis_id"`
	Evidence        environment.ProcessEvidence `json:"evidence"`
	Error           string                      `json:"error,omitempty"`
}

// NodeAction records attempted mutations even when a receipt is unavailable.
// CompletedAt is absent unless completion was established. Slice order records
// the actual sequential action order, never an inferred or sorted order.
type NodeAction struct {
	Node        environment.Node             `json:"node"`
	RequestedAt time.Time                    `json:"requested_at"`
	CompletedAt *time.Time                   `json:"completed_at,omitempty"`
	Stopped     *environment.ProcessEvidence `json:"stopped,omitempty"`
	Error       string                       `json:"error,omitempty"`
}

// PersistenceFault describes a drained, frozen-ledger all-process crash and
// restart. Topology readiness does not imply frozen data or persistence recovery.
// FrozenValidation precedes diagnostics that can write unrelated probe keys.
type PersistenceFault struct {
	Mode                        string                        `json:"mode"`
	FailureModel                string                        `json:"failure_model"`
	Prepared                    []NodeProcessSample           `json:"prepared"`
	LedgerFrozenAt              *time.Time                    `json:"ledger_frozen_at,omitempty"`
	TrafficDrainedAt            *time.Time                    `json:"traffic_drained_at,omitempty"`
	AcknowledgedFrozen          uint64                        `json:"acknowledged_frozen"`
	AmbiguousFrozen             uint64                        `json:"ambiguous_frozen"`
	PreCrashSnapshot            *cluster.Snapshot             `json:"pre_crash_snapshot,omitempty"`
	KillActions                 []NodeAction                  `json:"kill_actions"`
	Stopped                     []NodeProcessSample           `json:"stopped"`
	AllStoppedAt                *time.Time                    `json:"all_stopped_at,omitempty"`
	Restarts                    []*environment.RestartReceipt `json:"restarts"`
	RestartActions              []NodeAction                  `json:"restart_actions"`
	Restored                    []NodeProcessSample           `json:"restored"`
	RestartRequestedAt          *time.Time                    `json:"restart_requested_at,omitempty"`
	RestartCompletedAt          *time.Time                    `json:"restart_completed_at,omitempty"`
	DataCheckedAt               *time.Time                    `json:"data_checked_at,omitempty"`
	FrozenValidation            *workload.Validation          `json:"frozen_validation,omitempty"`
	DiagnosticsRequestedAt      *time.Time                    `json:"diagnostics_requested_at,omitempty"`
	ReplicaReady                *Milestone                    `json:"replica_ready,omitempty"`
	TopologyStable              *Milestone                    `json:"topology_stable,omitempty"`
	StableRedundancy            *Milestone                    `json:"stable_redundancy,omitempty"`
	DiagnosticServiceStable     *Milestone                    `json:"diagnostic_service_stable,omitempty"`
	RestoreTimedOut             bool                          `json:"restore_timed_out"`
	DrainToFirstKillMS          *float64                      `json:"drain_to_first_kill_ms,omitempty"`
	DrainToLastKillMS           *float64                      `json:"drain_to_last_kill_ms,omitempty"`
	AllStoppedToFirstRestartMS  *float64                      `json:"all_stopped_to_first_restart_ms,omitempty"`
	PersistenceRecoveryObserved *bool                         `json:"persistence_recovery_observed,omitempty"`
	PersistenceRecoveryAt       *Milestone                    `json:"persistence_recovery_at,omitempty"`
	Errors                      []string                      `json:"errors,omitempty"`
}

func optionalTime(at *time.Time) string {
	if at == nil || at.IsZero() {
		return "unobserved"
	}
	return at.Format(time.RFC3339Nano)
}

func (p PersistenceFault) markdown(b io.Writer) {
	fmt.Fprintf(b, "\n## Frozen-ledger process restart\n\nConfigured persistence mode: `%s`. Failure model: `%s`.\n", p.Mode, p.FailureModel)
	fmt.Fprintln(b, "\nWorkload traffic stops and drains before the sequential SIGKILL actions. No workload requests are offered during fault or restart observation; these phases do not measure zero availability or recovery throughput. Read-only frozen-ledger validation precedes separate diagnostic writes.")
	if p.LedgerFrozenAt == nil {
		fmt.Fprintln(b, "\nFrozen ledger counts: unobserved.")
	} else {
		fmt.Fprintf(b, "\nFrozen acknowledged SETs: %d. Frozen ambiguous SET outcomes: %d.\n", p.AcknowledgedFrozen, p.AmbiguousFrozen)
	}
	fmt.Fprintf(b, "\nRestoration deadline expired: %t.\n", p.RestoreTimedOut)
	fmt.Fprint(b, "\n| Evidence anchor | Observed time |\n|---|---|\n")
	for _, row := range []struct {
		name string
		at   *time.Time
	}{
		{"Traffic fully drained", p.TrafficDrainedAt}, {"Immutable ledger frozen", p.LedgerFrozenAt},
		{"All six stopped verified before restart", p.AllStoppedAt}, {"First restart requested", p.RestartRequestedAt},
		{"Restart sequence completed", p.RestartCompletedAt}, {"Frozen-ledger validation completed", p.DataCheckedAt},
		{"Post-validation diagnostics requested", p.DiagnosticsRequestedAt},
	} {
		fmt.Fprintf(b, "| %s | %s |\n", row.name, optionalTime(row.at))
	}
	for _, duration := range []struct {
		name  string
		value *float64
	}{
		{"Drain to first kill request", p.DrainToFirstKillMS}, {"Drain to last kill completion", p.DrainToLastKillMS},
		{"All-stopped barrier to first restart request", p.AllStoppedToFirstRestartMS},
	} {
		if duration.value != nil {
			fmt.Fprintf(b, "\n%s: %.3f ms (runtime monotonic clock).\n", duration.name, *duration.value)
		}
	}
	fmt.Fprint(b, "\n| Restart observation | Since fault orchestration began, ms |\n|---|---:|\n")
	for _, entry := range []struct {
		name string
		at   *Milestone
	}{
		{"Original six-node replication ready", p.ReplicaReady}, {"Original topology sustained", p.TopologyStable},
		{"Replication redundancy sustained", p.StableRedundancy}, {"Post-validation diagnostic service sustained", p.DiagnosticServiceStable},
		{"AOF frozen-value recovery observed", p.PersistenceRecoveryAt},
	} {
		if entry.at == nil {
			fmt.Fprintf(b, "| %s | unobserved |\n", entry.name)
		} else {
			fmt.Fprintf(b, "| %s | %.3f |\n", entry.name, entry.at.SinceFaultMS)
		}
	}
	if p.Mode == "memory" {
		fmt.Fprintln(b, "\nConfigured memory-only control disables RDB/AOF value persistence. Persistence recovery is not applicable; restored processes and topology can coexist with lost frozen values. The independent persistence dimension is false when the disabled runtime configuration was observed, while the persistence-recovery event remains absent.")
	} else {
		fmt.Fprintf(b, "\nAOF frozen-value recovery observed: %s. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.\n", observedFlag(p.PersistenceRecoveryObserved))
	}
	if p.FrozenValidation == nil || p.FrozenValidation.StartedAt.IsZero() {
		fmt.Fprintf(b, "\nFrozen-ledger validation unobserved; %d acknowledged values remain unchecked.\n", p.AcknowledgedFrozen)
	} else {
		v := p.FrozenValidation
		fmt.Fprintf(b, "\nFrozen-ledger validation: expected %d; checked %d; missing %d; corrupt %d; unreadable %d; unverified %d; complete %t. Completion describes verification coverage; observed loss remains a data failure.\n", v.Expected, v.Checked, v.Missing, v.Corrupt, v.Unreadable, v.Unverified, v.Complete)
	}
	for _, group := range []struct {
		name    string
		actions []NodeAction
	}{
		{"Kill", p.KillActions}, {"Restart", p.RestartActions},
	} {
		fmt.Fprintf(b, "\n### %s action order\n\n| Order | Managed node | Requested | Completed | Error |\n|---:|---|---|---|---|\n", group.name)
		for i, action := range group.actions {
			fmt.Fprintf(b, "| %d | %s | %s | %s | %s |\n", i+1, action.Node.ID, action.RequestedAt.Format(time.RFC3339Nano), optionalTime(action.CompletedAt), action.Error)
		}
	}
	fmt.Fprint(b, "\n### Stopped administrative observations\n\n| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |\n|---|---|---|---|---|---|\n")
	for _, sample := range p.Stopped {
		running := "unobserved"
		if !sample.Evidence.ObservedAt.IsZero() {
			running = fmt.Sprint(sample.Evidence.ContainerRunning)
		}
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s |\n", sample.Node.ID, sample.OriginalRedisID, optionalTime(&sample.Evidence.ObservedAt), running, sample.Evidence.ContainerRestartPolicy, sample.Error)
	}
	for _, sample := range p.Prepared {
		fmt.Fprintf(b, "\n### Prepared %s\n\nOriginal Redis Cluster ID: `%s`. Process run ID: `%s`. Observed: %s.\n", sample.Node.ID, sample.OriginalRedisID, sample.Evidence.RedisRunID, optionalTime(&sample.Evidence.ObservedAt))
		persistenceProcessMarkdown(b, sample.Evidence)
		if sample.Error != "" {
			fmt.Fprintf(b, "\nObservation error: %s\n", sample.Error)
		}
	}
	for i, receipt := range p.Restarts {
		if receipt == nil {
			continue
		}
		fmt.Fprintf(b, "\n### Restart receipt %d: %s\n\nOriginal Redis Cluster ID: `%s`. Observed after restart: %s.\n", i+1, receipt.NodeID, receipt.OriginalRedisID, optionalTime(&receipt.After.ObservedAt))
		persistenceProcessMarkdown(b, receipt.After)
	}
	if len(p.Restored) == 0 {
		fmt.Fprintln(b, "\nPost-synchronization administrative runtime observations: unobserved.")
	} else {
		fmt.Fprintln(b, "\n### Post-synchronization runtime evidence\n\nThese fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.")
		for _, sample := range p.Restored {
			fmt.Fprintf(b, "\n#### Restored %s\n\nOriginal Redis Cluster ID: `%s`. Process run ID: `%s`. Observed: %s.\n", sample.Node.ID, sample.OriginalRedisID, sample.Evidence.RedisRunID, optionalTime(&sample.Evidence.ObservedAt))
			persistenceProcessMarkdown(b, sample.Evidence)
			if sample.Error != "" {
				fmt.Fprintf(b, "\nObservation error: %s\n", sample.Error)
			}
		}
	}
	fmt.Fprintln(b, "\nThis is a sequential process crash and same-container restart with retained managed volumes. It does not remove Docker VM/host page cache or simulate power loss, storage loss, or a graceful shutdown. Draining and sequential kill/inspection time allow background fsync to progress, so everysec preservation does not establish a one-second-loss guarantee or equivalence to always.")
	for _, err := range p.Errors {
		fmt.Fprintf(b, "\nPersistence experiment error: %s\n", err)
	}
}

func persistenceProcessMarkdown(b io.Writer, evidence environment.ProcessEvidence) {
	if evidence.ObservedAt.IsZero() {
		fmt.Fprintln(b, "\nAdministrative process observation: unobserved; process and data-file status are not established.")
		return
	}
	usage := "unobserved"
	if evidence.DataSizeKiB > 0 {
		usage = fmt.Sprintf("%d KiB", evidence.DataSizeKiB)
	}
	fmt.Fprintf(b, "\nContainer running: %t. Redis PONG: %t. Redis Cluster ID: `%s`. Restart policy: `%s`. `/data` usage: %s.\n", evidence.ContainerRunning, evidence.RedisPONG, evidence.RedisClusterID, evidence.ContainerRestartPolicy, usage)
	for _, reply := range []struct{ name, raw string }{
		{"Observed persistence configuration", evidence.RawPersistenceConfig}, {"Observed INFO persistence", evidence.RawPersistence},
		{"Observed INFO keyspace", evidence.RawKeyspace}, {"Observed managed data files", evidence.RawDataFiles},
	} {
		if reply.raw == "" {
			continue
		}
		fmt.Fprintf(b, "\n%s:\n\n```text\n%s\n```\n", reply.name, reply.raw)
	}
}
