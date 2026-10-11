package batch

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/report"
)

const MaxRuns = 100

var runIDPattern = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{7,63}$`)

// Run preserves every requested trial and executes fresh trials sequentially.
// It refuses to proceed without authoritative cleanup evidence after a callback
// failure. Summary writes use a fresh operation, independent of a canceled ctx.
func Run(ctx context.Context, c config.Config, runs int, outputRoot string, trial TrialFunc) (*Summary, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	if trial == nil {
		return nil, errors.New("batch trial callback is required")
	}
	if runs < 1 || runs > MaxRuns {
		return nil, fmt.Errorf("batch runs must be between 1 and %d", MaxRuns)
	}
	var random [12]byte
	if _, err := rand.Read(random[:]); err != nil {
		return nil, err
	}
	root, err := filepath.Abs(outputRoot)
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(root, 0755); err != nil {
		return nil, err
	}
	id := "batch-" + hex.EncodeToString(random[:])
	dir := filepath.Join(root, id)
	if err := os.Mkdir(dir, 0755); err != nil {
		return nil, err
	}
	runsDir := filepath.Join(dir, "runs")
	if err := os.Mkdir(runsDir, 0755); err != nil {
		return nil, err
	}
	effective, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(effective)
	s := &Summary{SchemaVersion: 1, BatchID: id, Directory: dir, StartedAt: time.Now().UTC(), Config: c, ConfigSHA256: hex.EncodeToString(hash[:]), RequestedTrials: runs,
		LatencyAggregation: "No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.",
		Methodology: []string{
			"Trials execute sequentially with a fresh environment per callback; the next callback starts only after the previous callback returns with successful cleanup evidence (or an authoritative report that no environment was provisioned).",
			"Recovery success fraction uses all requested trials as the denominator. Outcomes, started/unstarted counts and individual reports preserve failed, partial and aborted trials; the started-trial fraction is reported separately.",
			"Recovery timings are conditional on observed milestones, including observed milestones in partial/failed trials. Unobserved post-fault events are censored at the accepted sampled observation horizon with deadline/abort/error reasons, never zero-duration samples; this does not prove the physical event occurred after that horizon. Trials without a requested fault are separately not at risk. No survival distribution, independence assumption or population reliability guarantee is inferred.",
			"Event durations use runtime SinceFaultMS and fault_action_duration_ms when present, preserving monotonic elapsed-time measurements. Schema 1 action-completion-relative durations fall back to wall-clock timestamp differences and can be affected by clock adjustments. First sampled observations include polling uncertainty; action duration includes ownership checks and fault command overhead.",
			"Min, median, mean, N-1 sample standard deviation and maximum describe available trial observations. P90/P95 use nearest rank only when N>=20; smaller samples have an explicit tail-estimation caveat.",
			"Completed-request success fractions weight actual completed counts; aggregate throughput divides summed completed/successful counts by summed phase wall-clock duration. Offered load and worker saturation are not inferred from completed requests.",
			"Whole-workload offered, started, completed and dropped schedule-slot counters are aggregated only for reports carrying explicit pacing methodology. Legacy absent counters are unmeasured, not zero. Offered=started+dropped and started=completed after workload quiescence; service probes and validation are excluded.",
			"Acknowledged SET writes remain distinct from ambiguous SET outcomes and final read-back validation. Skipped/incomplete validation contributes acknowledged-not-checked counts; no observed loss is not a durability guarantee.",
			"WAIT replica-memory acknowledgements and WAITAOF local/replica AOF acknowledgements are distinct policies. Each successful SET is accounted for by an attempted acknowledgement command or an explicit not-attempted count. Valid WAITAOF replies preserve paired local/replica counts and independently overlapping satisfaction counts; errors provide no acknowledgement observation. Diagnostic command stages do not increase logical request or pacing counts. A recovered policy trial additionally requires sustained policy availability and complete preservation of its acknowledged subset.",
			"Recovered schema 2 reports require the explicit combined-recovery flag. Legacy schema 1 reports can establish it through healthy topology/client flags and an observed combined-stability milestone, together with complete acknowledged-write validation and verified cleanup.",
			"Recovered primary-isolate trials additionally require sampled unchanged process survival while disconnected, restored connectivity, observed stable reconciliation and restored replication redundancy. Pre-restoration recovery alone does not establish a recovered network experiment.",
		}}
	for i := 1; i <= runs; i++ {
		s.Trials = append(s.Trials, Trial{Index: i, Outcome: "unstarted", Observations: map[string]Observation{}})
	}
	var runErr error
	seen := make(map[string]bool)
	for index := range s.Trials {
		if ctx.Err() != nil {
			s.StopReason = "cancellation"
			runErr = errors.Join(runErr, ctx.Err())
			break
		}
		started := time.Now().UTC()
		r, callbackErr := trial(ctx, c, runsDir)
		t := Trial{Index: index + 1, Started: true, StartedAt: started, FinishedAt: time.Now().UTC(), Outcome: "failed", Observations: map[string]Observation{}}
		if r == nil {
			callbackErr = errors.Join(callbackErr, errors.New("trial callback returned no report; resource cleanup cannot be established"))
			t.RunID = fmt.Sprintf("trial-%04d-error", index+1)
			t.SyntheticErrorReport = true
			t.Errors = []string{callbackErr.Error()}
			if ctx.Err() != nil {
				t.Outcome = "aborted"
			}
			synthetic := report.Report{SchemaVersion: 1, RunID: t.RunID, Outcome: t.Outcome, StartedAt: t.StartedAt, FinishedAt: t.FinishedAt, Config: c, ConfigSHA256: s.ConfigSHA256, Errors: t.Errors}
			path := filepath.Join(runsDir, t.RunID)
			if err := os.Mkdir(path, 0755); err != nil {
				callbackErr = errors.Join(callbackErr, fmt.Errorf("save synthetic error report: %w", err))
			} else if err := report.Write(path, synthetic); err != nil {
				callbackErr = errors.Join(callbackErr, fmt.Errorf("save synthetic error report: %w", err))
			} else {
				t.ReportPath = filepath.Join("runs", t.RunID, "report.json")
				t.ReportJSONAvailable, t.ReportMarkdownAvailable = true, true
			}
			s.StopReason = "missing_cleanup_evidence"
		} else {
			t = trialEvidence(index+1, r)
			if t.StartedAt.IsZero() {
				t.StartedAt = started
			}
			if t.FinishedAt.IsZero() {
				t.FinishedAt = time.Now().UTC()
			}
			verificationErr := verifyEvidence(r, runsDir, c, s.ConfigSHA256, seen)
			if verificationErr != nil {
				callbackErr = errors.Join(callbackErr, verificationErr)
				t.ExperimentOutcome = t.Outcome
				t.Outcome = "failed"
				t.Errors = append(t.Errors, verificationErr.Error())
				t.Phases = nil
				t.Observations = map[string]Observation{}
				t.FaultActionDurationMS = nil
				if runIDPattern.MatchString(r.RunID) {
					if info, err := os.Lstat(filepath.Join(runsDir, r.RunID, "report.json")); err == nil && info.Mode().IsRegular() {
						t.ReportJSONAvailable = true
						t.ReportPath = filepath.Join("runs", r.RunID, "report.json")
					}
					if info, err := os.Lstat(filepath.Join(runsDir, r.RunID, "report.md")); err == nil && info.Mode().IsRegular() {
						t.ReportMarkdownAvailable = true
					}
				}
			} else {
				t.EvidenceComplete = true
				t.ReportPath = filepath.Join("runs", r.RunID, "report.json")
				t.ReportJSONAvailable, t.ReportMarkdownAvailable = true, true
			}
			if callbackErr != nil {
				t.Errors = appendUnique(t.Errors, callbackErr.Error())
				if t.Outcome == "recovered" {
					t.ExperimentOutcome = t.Outcome
					t.Outcome = "failed"
				}
			}
			if t.Outcome == "failed" || t.Outcome == "partial" || t.Outcome == "aborted" {
				if callbackErr == nil {
					callbackErr = fmt.Errorf("trial outcome: %s", t.Outcome)
				}
			}
			if !runIDPattern.MatchString(r.RunID) || seen[r.RunID] {
				s.StopReason = "invalid_or_duplicate_run_id"
			}
			seen[r.RunID] = true
			if r.Environment != nil && (!r.Cleanup.Attempted || !r.Cleanup.Successful) {
				s.StopReason = "cleanup_failure"
				callbackErr = errors.Join(callbackErr, errors.New("trial environment cleanup was not verified successful"))
				if t.Outcome != "aborted" {
					t.ExperimentOutcome = r.Outcome
					t.Outcome = "failed"
				}
			}
		}
		s.Trials[index] = t
		if callbackErr != nil {
			runErr = errors.Join(runErr, fmt.Errorf("trial %d: %w", index+1, callbackErr))
		}
		if ctx.Err() != nil {
			s.StopReason = "cancellation"
			runErr = errors.Join(runErr, ctx.Err())
		}
		s.FinishedAt = time.Now().UTC()
		Aggregate(s)
		if runErr != nil {
			s.Errors = []string{runErr.Error()}
		}
		if err := Write(dir, *s); err != nil {
			runErr = errors.Join(runErr, fmt.Errorf("write batch summary: %w", err))
			s.StopReason = "summary_write_failure"
			break
		}
		if s.StopReason != "" {
			break
		}
	}
	s.FinishedAt = time.Now().UTC()
	Aggregate(s)
	if runErr != nil {
		s.Errors = []string{runErr.Error()}
	}
	if err := Write(dir, *s); err != nil {
		runErr = errors.Join(runErr, fmt.Errorf("write batch summary: %w", err))
		s.StopReason = "summary_write_failure"
		s.Errors = []string{runErr.Error()}
		Aggregate(s)
	}
	return s, runErr
}

func trialEvidence(index int, r *report.Report) Trial {
	t := Trial{Index: index, Started: true, RunID: r.RunID, Outcome: r.Outcome, StartedAt: r.StartedAt, FinishedAt: r.FinishedAt, ConfigSHA256: r.ConfigSHA256,
		EnvironmentPresent: r.Environment != nil, Cleanup: r.Cleanup, Errors: append([]string(nil), r.Errors...), Observations: observations(r), Phases: r.Metrics.Phases, WorkloadStats: r.WorkloadStats, Validation: r.Validation}
	if r.FaultActionDurationMS != nil {
		t.FaultActionDurationMS = number(*r.FaultActionDurationMS)
	} else if r.FaultRequestedAt != nil && r.FaultCompletedAt != nil {
		t.FaultActionDurationMS = number(float64(r.FaultCompletedAt.Sub(*r.FaultRequestedAt)) / float64(time.Millisecond))
	}
	return t
}

func verifyEvidence(r *report.Report, root string, c config.Config, hash string, seen map[string]bool) error {
	if r.SchemaVersion != 1 && r.SchemaVersion != 2 && r.SchemaVersion != report.LifecycleSchemaVersion {
		return fmt.Errorf("unsupported individual report schema %d", r.SchemaVersion)
	}
	if !runIDPattern.MatchString(r.RunID) {
		return fmt.Errorf("invalid trial run ID %q", r.RunID)
	}
	if seen[r.RunID] {
		return fmt.Errorf("duplicate trial run ID %q", r.RunID)
	}
	switch r.Outcome {
	case "recovered", "partial", "failed", "aborted":
	default:
		return fmt.Errorf("unknown trial outcome %q", r.Outcome)
	}
	actual, err := json.Marshal(r.Config)
	if err != nil {
		return err
	}
	expected, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if !bytes.Equal(actual, expected) || r.ConfigSHA256 != hash {
		return errors.New("trial effective configuration or its SHA-256 differs from the requested experiment")
	}
	path := filepath.Join(root, r.RunID, "report.json")
	for _, name := range []string{"report.json", "report.md"} {
		info, err := os.Lstat(filepath.Join(root, r.RunID, name))
		if err != nil {
			return fmt.Errorf("trial %s unavailable: %w", name, err)
		}
		if !info.Mode().IsRegular() || info.Size() == 0 || info.Size() > 64<<20 {
			return fmt.Errorf("trial %s must be a nonempty regular file of at most 64 MiB", name)
		}
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	var persisted report.Report
	if err := json.Unmarshal(data, &persisted); err != nil {
		return fmt.Errorf("read individual JSON report: %w", err)
	}
	persistedData, err := json.Marshal(persisted)
	if err != nil {
		return err
	}
	memoryData, err := json.Marshal(r)
	if err != nil {
		return err
	}
	if !bytes.Equal(persistedData, memoryData) {
		return errors.New("individual JSON report does not match returned trial evidence")
	}
	markdown, err := os.ReadFile(filepath.Join(root, r.RunID, "report.md"))
	if err != nil {
		return err
	}
	if string(markdown) != r.Markdown() {
		return errors.New("individual Markdown report does not match returned trial evidence")
	}
	if err := validNumbers(r); err != nil {
		return err
	}
	if r.Outcome == "recovered" {
		v, w, recovery := r.Validation, r.WorkloadStats, r.Recovery
		if !recovery.TopologyRecovered || !recovery.ClientRecovered || (r.SchemaVersion >= 2 && !recovery.CombinedRecovered) || recovery.CombinedStable == nil || !recovery.DataVerified || recovery.TimedOut || !v.Complete || v.Expected == 0 || v.Missing != 0 || v.Corrupt != 0 || v.Unreadable != 0 || v.Unverified != 0 || v.StartedAt.IsZero() || v.CompletedAt.IsZero() || w.MissingReads != 0 || w.MismatchedReads != 0 || (r.Environment != nil && (!r.Cleanup.Attempted || !r.Cleanup.Successful)) {
			return errors.New("recovered trial outcome lacks successful recovery, acknowledged-write validation or cleanup evidence")
		}
		if err := recoveredLifecycleEvidence(r); err != nil {
			return err
		}
		if err := recoveredPersistenceEvidence(r); err != nil {
			return err
		}
		if err := recoveredPolicyRuntimeEvidence(r); err != nil {
			return err
		}
		if r.Config.Experiment.Fault == "primary-isolate" && !recoveredNetworkEvidence(r) {
			return errors.New("recovered network trial lacks unchanged disconnected process survival, restoration, reconciliation or redundancy evidence")
		}
	}
	return nil
}

func validNumbers(r *report.Report) error {
	if err := validPolicyNumbers(r); err != nil {
		return err
	}
	w := r.WorkloadStats
	if w.PacingMethodology != "" && (w.StartedRequests > w.OfferedRequests || w.DroppedRequests != w.OfferedRequests-w.StartedRequests || w.StartedRequests != w.CompletedRequests || w.CompletedRequests != r.Metrics.Totals.Attempts) {
		return errors.New("trial offered/started/completed/dropped accounting does not reconcile")
	}
	for _, m := range r.Metrics.Phases {
		if !finite(m.DurationSeconds) || m.DurationSeconds < 0 || (m.Attempts > 0 && m.DurationSeconds == 0) || m.Successes > m.Attempts || m.Errors != m.Attempts-m.Successes {
			return errors.New("invalid trial phase operation counts or duration")
		}
		for _, h := range []struct {
			count  uint64
			values []float64
		}{{m.SuccessfulLatency.Count, []float64{m.SuccessfulLatency.P50MS, m.SuccessfulLatency.P95MS, m.SuccessfulLatency.P99MS}}, {m.FailedLatency.Count, []float64{m.FailedLatency.P50MS, m.FailedLatency.P95MS, m.FailedLatency.P99MS}}} {
			for _, value := range h.values {
				if !finite(value) || value < 0 {
					return errors.New("invalid trial latency summary")
				}
			}
		}
	}
	v := r.Validation
	if v.Checked < 0 || v.Expected < 0 || v.Missing < 0 || v.Corrupt < 0 || v.Unreadable < 0 || v.Unverified < 0 || v.Checked > v.Expected || v.Missing+v.Corrupt+v.Unreadable > v.Checked {
		return errors.New("invalid trial validation counts")
	}
	if !v.StartedAt.IsZero() && (uint64(v.Expected) != r.WorkloadStats.AcknowledgedWrites || v.Unverified != v.Expected-v.Checked || v.Complete != (v.Unverified == 0)) {
		return errors.New("trial validation does not account for the acknowledged ledger")
	}
	return nil
}

func appendUnique(values []string, value string) []string {
	for _, existing := range values {
		if existing == value {
			return values
		}
	}
	return append(values, value)
}
