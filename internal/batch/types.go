// Package batch executes independent sequential trials and preserves conditional,
// censored and missing evidence in aggregate experiment reports.
package batch

import (
	"context"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/metrics"
	"github.com/alanwang/redis-reliability-lab/internal/report"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

// TrialFunc must create a fresh isolated environment and return only after
// bounded cleanup finishes. experiment.Runner.Run satisfies this contract.
// Its outputRoot argument is the batch's runs/ directory.
type TrialFunc func(context.Context, config.Config, string) (*report.Report, error)

type Observation struct {
	Status                       string     `json:"status"`
	SinceFaultRequestedMS        *float64   `json:"since_fault_requested_ms,omitempty"`
	SinceFaultCompletedMS        *float64   `json:"since_fault_completed_ms,omitempty"`
	ObservationHorizonAt         *time.Time `json:"observation_horizon_at,omitempty"`
	HorizonSinceFaultRequestedMS *float64   `json:"horizon_since_fault_requested_ms,omitempty"`
	HorizonSinceFaultCompletedMS *float64   `json:"horizon_since_fault_completed_ms,omitempty"`
}

type Trial struct {
	Index                   int                             `json:"index"`
	Started                 bool                            `json:"started"`
	RunID                   string                          `json:"run_id,omitempty"`
	Outcome                 string                          `json:"outcome"`
	ExperimentOutcome       string                          `json:"experiment_outcome,omitempty"`
	StartedAt               time.Time                       `json:"started_at"`
	FinishedAt              time.Time                       `json:"finished_at"`
	ConfigSHA256            string                          `json:"effective_config_sha256,omitempty"`
	ReportPath              string                          `json:"report_path,omitempty"`
	ReportJSONAvailable     bool                            `json:"report_json_available"`
	ReportMarkdownAvailable bool                            `json:"report_markdown_available"`
	EvidenceComplete        bool                            `json:"evidence_complete"`
	SyntheticErrorReport    bool                            `json:"synthetic_error_report,omitempty"`
	EnvironmentPresent      bool                            `json:"environment_present"`
	Cleanup                 report.Cleanup                  `json:"cleanup"`
	Errors                  []string                        `json:"errors,omitempty"`
	FaultActionDurationMS   *float64                        `json:"fault_action_duration_ms,omitempty"`
	Observations            map[string]Observation          `json:"observations"`
	Phases                  map[string]metrics.PhaseSummary `json:"phases,omitempty"`
	WorkloadStats           workload.Stats                  `json:"workload_stats"`
	Validation              workload.Validation             `json:"validation"`
}

type EventStatistics struct {
	Observed               int            `json:"observed"`
	NotAtRisk              int            `json:"not_at_risk"`
	Censored               int            `json:"censored"`
	CensoredDeadline       int            `json:"censored_deadline"`
	CensoredAborted        int            `json:"censored_aborted"`
	CensoredError          int            `json:"censored_error"`
	CensoredOther          int            `json:"censored_other"`
	Missing                int            `json:"missing"`
	Unstarted              int            `json:"unstarted"`
	MissingReasons         map[string]int `json:"missing_reasons"`
	RequestedAnchorMissing int            `json:"requested_anchor_missing"`
	CompletedAnchorMissing int            `json:"completed_anchor_missing"`
	SinceFaultRequestedMS  Statistics     `json:"since_fault_requested_ms"`
	SinceFaultCompletedMS  Statistics     `json:"since_fault_completed_ms"`
}

type PhaseStatistics struct {
	TrialsWithPhase          int               `json:"trials_with_phase"`
	TrialsWithAttempts       int               `json:"trials_with_attempts"`
	Attempts                 uint64            `json:"attempts"`
	Successes                uint64            `json:"successes"`
	Errors                   uint64            `json:"errors"`
	ErrorClasses             map[string]uint64 `json:"error_classes"`
	DurationSeconds          float64           `json:"sum_trial_phase_duration_seconds"`
	SuccessFraction          *float64          `json:"completed_request_success_fraction,omitempty"`
	CompletedThroughput      *float64          `json:"completed_throughput_over_sum_phase_seconds,omitempty"`
	SuccessfulThroughput     *float64          `json:"successful_throughput_over_sum_phase_seconds,omitempty"`
	TrialSuccessFraction     Statistics        `json:"trial_success_fraction"`
	TrialCompletedThroughput Statistics        `json:"trial_completed_throughput_ops_per_second"`
	TrialSuccessfulP50MS     Statistics        `json:"trial_successful_latency_p50_ms"`
	TrialSuccessfulP95MS     Statistics        `json:"trial_successful_latency_p95_ms"`
	TrialSuccessfulP99MS     Statistics        `json:"trial_successful_latency_p99_ms"`
	TrialFailedP50MS         Statistics        `json:"trial_failed_latency_p50_ms"`
	TrialFailedP95MS         Statistics        `json:"trial_failed_latency_p95_ms"`
	TrialFailedP99MS         Statistics        `json:"trial_failed_latency_p99_ms"`
}

type DataStatistics struct {
	ValidationAttemptedTrials int        `json:"validation_attempted_trials"`
	ValidationCompleteTrials  int        `json:"validation_complete_trials"`
	AcknowledgedWrites        uint64     `json:"acknowledged_writes"`
	WritesAttempted           uint64     `json:"writes_attempted"`
	AmbiguousWrites           uint64     `json:"ambiguous_writes"`
	Checked                   uint64     `json:"checked"`
	Missing                   uint64     `json:"missing"`
	Corrupt                   uint64     `json:"corrupt"`
	Unreadable                uint64     `json:"unreadable"`
	Unverified                uint64     `json:"unverified"`
	AcknowledgedNotChecked    uint64     `json:"acknowledged_not_checked"`
	MissingReadChecks         uint64     `json:"missing_read_checks"`
	MismatchedReadChecks      uint64     `json:"mismatched_read_checks"`
	TrialAcknowledgedWrites   Statistics `json:"trial_acknowledged_writes"`
	TrialAmbiguousWrites      Statistics `json:"trial_ambiguous_writes"`
	TrialMissingWrites        Statistics `json:"trial_missing_writes"`
	TrialCorruptWrites        Statistics `json:"trial_corrupt_writes"`
	TrialUnreadableWrites     Statistics `json:"trial_unreadable_writes"`
}

// LoadStatistics uses whole-workload counters. Legacy reports lacking the pacing
// methodology are counted as unmeasured, never treated as measured zero demand.
type LoadStatistics struct {
	MeasuredTrials       int        `json:"measured_trials"`
	UnmeasuredTrials     int        `json:"unmeasured_trials"`
	OfferedRequests      uint64     `json:"offered_requests"`
	StartedRequests      uint64     `json:"started_requests"`
	CompletedRequests    uint64     `json:"completed_requests"`
	DroppedRequests      uint64     `json:"dropped_requests"`
	StartedFraction      *float64   `json:"started_fraction_of_offered,omitempty"`
	TrialOfferedRequests Statistics `json:"trial_offered_requests"`
	TrialStartedRequests Statistics `json:"trial_started_requests"`
	TrialDroppedRequests Statistics `json:"trial_dropped_requests"`
}

type PolicyStatistics struct {
	MeasuredTrials                int                              `json:"measured_trials"`
	UnmeasuredTrials              int                              `json:"unmeasured_trials"`
	Totals                        workload.PolicyCounts            `json:"totals"`
	ByPhase                       map[string]workload.PolicyCounts `json:"by_phase"`
	WaitSatisfiedSubset           workload.ValidationSubset        `json:"wait_satisfied_subset"`
	WaitSubsetValidationTrials    int                              `json:"wait_subset_validation_trials"`
	WaitAOFSatisfiedSubset        workload.ValidationSubset        `json:"waitaof_satisfied_subset"`
	WaitAOFSubsetValidationTrials int                              `json:"waitaof_subset_validation_trials"`
}

type Summary struct {
	SchemaVersion                   int                        `json:"schema_version"`
	BatchID                         string                     `json:"batch_id"`
	Directory                       string                     `json:"directory"`
	Outcome                         string                     `json:"outcome"`
	StartedAt                       time.Time                  `json:"started_at"`
	FinishedAt                      time.Time                  `json:"finished_at"`
	Config                          config.Config              `json:"config"`
	ConfigSHA256                    string                     `json:"effective_config_sha256"`
	RequestedTrials                 int                        `json:"requested_trials"`
	StartedTrials                   int                        `json:"started_trials"`
	CompletedTrials                 int                        `json:"completed_trials"`
	UnstartedTrials                 int                        `json:"unstarted_trials"`
	Outcomes                        map[string]int             `json:"outcomes"`
	RecoverySuccessFraction         float64                    `json:"recovery_success_fraction_all_requested_trials"`
	ExecutedRecoverySuccessFraction *float64                   `json:"recovery_success_fraction_started_trials,omitempty"`
	CleanupFailures                 int                        `json:"cleanup_failures"`
	IncompleteEvidenceTrials        int                        `json:"incomplete_evidence_trials"`
	StopReason                      string                     `json:"stop_reason,omitempty"`
	Trials                          []Trial                    `json:"trials"`
	Events                          map[string]EventStatistics `json:"events"`
	FaultActionDurationMS           Statistics                 `json:"fault_action_duration_ms"`
	Phases                          map[string]PhaseStatistics `json:"phases"`
	Data                            DataStatistics             `json:"data"`
	Load                            LoadStatistics             `json:"load"`
	WritePolicy                     PolicyStatistics           `json:"write_policy"`
	LatencyAggregation              string                     `json:"latency_aggregation"`
	Methodology                     []string                   `json:"methodology"`
	Errors                          []string                   `json:"errors,omitempty"`
}
