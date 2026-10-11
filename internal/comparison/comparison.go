// Package comparison counterbalances sequential, independently verified trials.
// It groups measurements by condition without pooling request latency quantiles.
package comparison

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	mathrand "math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/batch"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/report"
)

type Condition struct {
	Index       int           `json:"index"`
	ID          string        `json:"condition_id"`
	Config      config.Config `json:"config"`
	ConfigHash  string        `json:"effective_config_sha256"`
	SummaryPath string        `json:"summary_path"`
	Summary     batch.Summary `json:"summary"`
}

type ScheduledTrial struct {
	Position              int               `json:"position"`
	Round                 int               `json:"round"`
	RoundPosition         int               `json:"round_position"`
	ConditionIndex        int               `json:"condition_index"`
	ConditionID           string            `json:"condition_id"`
	ConfigHash            string            `json:"effective_config_sha256"`
	BatchID               string            `json:"batch_id,omitempty"`
	BatchSummaryPath      string            `json:"batch_summary_path,omitempty"`
	ReportPath            string            `json:"report_path,omitempty"`
	HostEvidenceReference string            `json:"host_evidence_reference,omitempty"`
	EnvironmentReference  string            `json:"environment_evidence_reference,omitempty"`
	Host                  map[string]string `json:"host,omitempty"`
	EnvironmentMetadata   map[string]string `json:"environment_metadata,omitempty"`
	Trial                 batch.Trial       `json:"trial"`
}

type Summary struct {
	SchemaVersion   int              `json:"schema_version"`
	ComparisonID    string           `json:"comparison_id"`
	Directory       string           `json:"directory"`
	Outcome         string           `json:"outcome"`
	StartedAt       time.Time        `json:"started_at"`
	FinishedAt      time.Time        `json:"finished_at"`
	OrderSeed       int64            `json:"order_seed"`
	Rounds          int              `json:"rounds"`
	BaseOrder       []int            `json:"base_condition_order"`
	RequestedTrials int              `json:"requested_trials"`
	StartedTrials   int              `json:"started_trials"`
	CompletedTrials int              `json:"completed_trials"`
	UnstartedTrials int              `json:"unstarted_trials"`
	Outcomes        map[string]int   `json:"outcomes"`
	Conditions      []Condition      `json:"conditions"`
	Schedule        []ScheduledTrial `json:"schedule"`
	StopReason      string           `json:"stop_reason,omitempty"`
	Methodology     []string         `json:"methodology"`
	Errors          []string         `json:"errors,omitempty"`
}

// Run executes one independent batch trial per schedule entry. A seeded base
// permutation rotates once per round: each condition occupies every position
// once in a full block of N rounds. This does not balance all carryover pairs.
func Run(ctx context.Context, configs []config.Config, rounds int, orderSeed int64, outputRoot string, trial batch.TrialFunc) (*Summary, error) {
	if len(configs) < 2 || len(configs) > 8 || rounds < 1 || rounds > 100 || rounds*len(configs) > batch.MaxRuns {
		return nil, errors.New("comparison requires 2–8 conditions, 1–100 rounds and at most 100 total trials")
	}
	if trial == nil {
		return nil, errors.New("comparison trial callback is required")
	}
	for i, c := range configs {
		if err := c.Validate(); err != nil {
			return nil, fmt.Errorf("condition %d: %w", i+1, err)
		}
	}
	var random [12]byte
	if _, err := rand.Read(random[:]); err != nil {
		return nil, err
	}
	root, err := filepath.Abs(outputRoot)
	if err != nil {
		return nil, err
	}
	id := "comparison-" + hex.EncodeToString(random[:])
	dir := filepath.Join(root, id)
	if err := os.MkdirAll(root, 0755); err != nil {
		return nil, err
	}
	if err := os.Mkdir(dir, 0755); err != nil {
		return nil, err
	}
	s := &Summary{SchemaVersion: 1, ComparisonID: id, Directory: dir, StartedAt: time.Now().UTC(), OrderSeed: orderSeed, Rounds: rounds, RequestedTrials: len(configs) * rounds,
		Methodology: []string{
			"A seeded permutation of condition indices is rotated by one position each round. Within each complete block of N rounds, every condition occupies every round position once. An incomplete block is only partially balanced; this rotating Latin order does not balance all predecessor/successor pairs.",
			"The order seed controls only the condition schedule. Effective workload seeds and all other configuration values remain unchanged. Each entry executes one fresh sequential batch trial with persisted-report integrity and cleanup checks.",
			"Further trials require verified individual evidence and successful cleanup, or verified evidence that no environment was provisioned. Unsafe cleanup, missing evidence, cancellation and summary-write failures stop execution; all remaining planned entries stay explicitly unstarted.",
			"Failed and partial trials with valid evidence and safe cleanup remain in their condition statistics. Timing statistics are conditional on observed milestones; missing and censored observations are not zero-duration events. Request latency quantiles are not pooled across trials or conditions.",
			"Condition summaries reuse batch aggregation. Condition-summary report paths are relative to their own summary directory; schedule report paths and batch references are relative to the comparison directory. Host and environment references point to individual JSON report sections.",
			"Ordering reduces systematic position confounding but does not prove causality, independent samples, equal host conditions or population reliability. Compare individual environmental metadata and actual order before interpreting differences.",
		}}
	for i, c := range configs {
		data, err := json.Marshal(c)
		if err != nil {
			return nil, err
		}
		hash := sha256.Sum256(data)
		cid := fmt.Sprintf("condition-%02d", i+1)
		conditionDir := filepath.Join(dir, "conditions", cid)
		if err := os.MkdirAll(conditionDir, 0755); err != nil {
			return nil, err
		}
		condition := Condition{Index: i + 1, ID: cid, Config: c, ConfigHash: hex.EncodeToString(hash[:]), SummaryPath: filepath.ToSlash(filepath.Join("conditions", cid, "summary.json")),
			Summary: batch.Summary{SchemaVersion: 1, BatchID: id + "-" + cid, Directory: conditionDir, StartedAt: s.StartedAt, Config: c, ConfigSHA256: hex.EncodeToString(hash[:]), RequestedTrials: rounds, Methodology: append([]string(nil), s.Methodology...), LatencyAggregation: "No pooled request latency quantiles are computed; per-trial quantile distributions are conditional on available observations."}}
		for round := 1; round <= rounds; round++ {
			condition.Summary.Trials = append(condition.Summary.Trials, batch.Trial{Index: round, Outcome: "unstarted", Observations: map[string]batch.Observation{}})
		}
		s.Conditions = append(s.Conditions, condition)
	}
	order := mathrand.New(mathrand.NewSource(orderSeed)).Perm(len(configs))
	for _, index := range order {
		s.BaseOrder = append(s.BaseOrder, index+1)
	}
	for round := 0; round < rounds; round++ {
		for position := range order {
			index := order[(position+round)%len(order)]
			condition := &s.Conditions[index]
			s.Schedule = append(s.Schedule, ScheduledTrial{Position: len(s.Schedule) + 1, Round: round + 1, RoundPosition: position + 1, ConditionIndex: index + 1, ConditionID: condition.ID, ConfigHash: condition.ConfigHash,
				Trial: batch.Trial{Index: round + 1, Outcome: "unstarted", Observations: map[string]batch.Observation{}}})
		}
	}
	update(s)
	if err := Write(dir, *s); err != nil {
		s.StopReason = "summary_write_failure"
		s.FinishedAt = time.Now().UTC()
		s.Errors = []string{err.Error()}
		update(s)
		return s, err
	}
	var runErr error
	seenRuns := map[string]bool{}
	for i := range s.Schedule {
		if ctx.Err() != nil {
			s.StopReason = "cancellation"
			runErr = errors.Join(runErr, ctx.Err())
			break
		}
		entry := &s.Schedule[i]
		condition := &s.Conditions[entry.ConditionIndex-1]
		trialRoot := filepath.Join(dir, "trials", fmt.Sprintf("position-%04d", entry.Position))
		var returned *report.Report
		one, trialErr := batch.Run(ctx, condition.Config, 1, trialRoot, func(ctx context.Context, c config.Config, output string) (*report.Report, error) {
			var err error
			returned, err = trial(ctx, c, output)
			return returned, err
		})
		if trialErr != nil {
			runErr = errors.Join(runErr, fmt.Errorf("schedule position %d (%s): %w", entry.Position, entry.ConditionID, trialErr))
		}
		if one == nil || len(one.Trials) != 1 {
			s.StopReason = "trial_setup_failure"
			runErr = errors.Join(runErr, errors.New("single-trial batch returned no authoritative trial"))
			break
		}
		entry.BatchID = one.BatchID
		entry.BatchSummaryPath = relative(dir, filepath.Join(one.Directory, "summary.json"))
		entry.Trial = one.Trials[0]
		entry.Trial.Index = entry.Round
		if entry.Trial.ReportPath != "" {
			entry.ReportPath = relative(dir, filepath.Join(one.Directory, entry.Trial.ReportPath))
			entry.Trial.ReportPath = entry.ReportPath
			entry.HostEvidenceReference = entry.ReportPath + "#/host"
			if entry.Trial.EnvironmentPresent {
				entry.EnvironmentReference = entry.ReportPath + "#/environment"
			}
		}
		if entry.Trial.EvidenceComplete && returned != nil {
			entry.Host = cloneMetadata(returned.Host)
			if returned.Environment != nil {
				entry.EnvironmentMetadata = cloneMetadata(returned.Environment.Metadata)
			}
		}
		if entry.Trial.Started && entry.Trial.RunID != "" {
			if seenRuns[entry.Trial.RunID] {
				s.StopReason = "duplicate_run_id"
				entry.Trial.ExperimentOutcome, entry.Trial.Outcome = entry.Trial.Outcome, "failed"
				entry.Trial.EvidenceComplete = false
				runErr = errors.Join(runErr, fmt.Errorf("run ID %q reused across comparison trials", entry.Trial.RunID))
			}
			seenRuns[entry.Trial.RunID] = true
		}
		grouped := entry.Trial
		if entry.ReportPath != "" {
			grouped.ReportPath = relative(condition.Summary.Directory, filepath.Join(dir, entry.ReportPath))
		}
		condition.Summary.Trials[entry.Round-1] = grouped
		condition.Summary.LatencyAggregation = one.LatencyAggregation
		if len(one.Errors) > 0 {
			condition.Summary.Errors = append(condition.Summary.Errors, one.Errors...)
		}
		if s.StopReason == "" && one.StopReason != "" {
			s.StopReason = one.StopReason
		}
		if s.StopReason == "" && (!entry.Trial.EvidenceComplete || entry.Trial.SyntheticErrorReport) {
			s.StopReason = "missing_verified_evidence"
			runErr = errors.Join(runErr, errors.New("individual trial evidence was not verified; cleanup authority cannot be established"))
		}
		if s.StopReason == "" && entry.Trial.EnvironmentPresent && (!entry.Trial.Cleanup.Attempted || !entry.Trial.Cleanup.Successful) {
			s.StopReason = "cleanup_failure"
			runErr = errors.Join(runErr, errors.New("individual environment cleanup was not verified"))
		}
		if ctx.Err() != nil {
			s.StopReason = "cancellation"
			runErr = errors.Join(runErr, ctx.Err())
		}
		s.FinishedAt = time.Now().UTC()
		update(s)
		if runErr != nil {
			s.Errors = []string{runErr.Error()}
		}
		if err := Write(dir, *s); err != nil {
			s.StopReason = "summary_write_failure"
			runErr = errors.Join(runErr, err)
			break
		}
		if s.StopReason != "" {
			break
		}
	}
	s.FinishedAt = time.Now().UTC()
	update(s)
	if runErr != nil {
		s.Errors = []string{runErr.Error()}
	}
	if err := Write(dir, *s); err != nil {
		s.StopReason = "summary_write_failure"
		runErr = errors.Join(runErr, err)
		s.Errors = []string{runErr.Error()}
		update(s)
	}
	return s, runErr
}

func cloneMetadata(values map[string]string) map[string]string {
	if values == nil {
		return nil
	}
	out := make(map[string]string, len(values))
	for key, value := range values {
		out[key] = value
	}
	return out
}

func relative(root, target string) string {
	p, err := filepath.Rel(root, target)
	if err != nil {
		panic(err) // All paths are constructed under one absolute comparison root.
	}
	return filepath.ToSlash(p)
}

func update(s *Summary) {
	s.StartedTrials, s.CompletedTrials = 0, 0
	s.Outcomes = map[string]int{"recovered": 0, "partial": 0, "failed": 0, "aborted": 0, "unstarted": 0}
	for _, entry := range s.Schedule {
		s.Outcomes[entry.Trial.Outcome]++
		if entry.Trial.Started {
			s.StartedTrials++
			if !entry.Trial.FinishedAt.IsZero() {
				s.CompletedTrials++
			}
		}
	}
	s.UnstartedTrials = s.RequestedTrials - s.StartedTrials
	for i := range s.Conditions {
		condition := &s.Conditions[i]
		condition.Summary.FinishedAt = s.FinishedAt
		condition.Summary.StopReason = s.StopReason
		batch.Aggregate(&condition.Summary)
	}
	switch {
	case s.StopReason == "cancellation":
		s.Outcome = "aborted"
	case s.StopReason != "":
		s.Outcome = "failed"
	case s.Outcomes["recovered"] == s.RequestedTrials:
		s.Outcome = "recovered"
	case s.Outcomes["recovered"]+s.Outcomes["partial"] > 0:
		s.Outcome = "partial"
	default:
		s.Outcome = "failed"
	}
}
