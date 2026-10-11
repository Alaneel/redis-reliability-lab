package batch

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/metrics"
	"github.com/alanwang/redis-reliability-lab/internal/report"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

func fixture(c config.Config, index int) *report.Report {
	start := time.Unix(int64(10000+index*10), 0).UTC()
	requested := start.Add(time.Second)
	completed := requested.Add(100 * time.Millisecond)
	at := requested.Add(time.Duration(index) * time.Second)
	data, _ := json.Marshal(c)
	hash := sha256.Sum256(data)
	m := &report.Milestone{At: at, SinceFaultMS: float64(index * 1000)}
	return &report.Report{SchemaVersion: 2, RunID: fmt.Sprintf("trial-%04d-valid", index), Outcome: "recovered", StartedAt: start, FinishedAt: at.Add(time.Second),
		Config: c, ConfigSHA256: hex.EncodeToString(hash[:]), Environment: &environment.Environment{RunID: fmt.Sprintf("trial-%04d-valid", index)},
		FaultRequestedAt: &requested, FaultCompletedAt: &completed, FaultActionDurationMS: number(100),
		Recovery:      report.Recovery{FailureDetected: m, ReplicaPromoted: m, FullCoverage: m, TopologyConverged: m, ClientFirstSuccess: m, ClientStable: m, CombinedStable: m, TopologyRecovered: true, ClientRecovered: true, CombinedRecovered: true, DataVerified: true},
		Samples:       []report.Sample{{Snapshot: cluster.Snapshot{At: at.Add(-time.Millisecond), Completed: at}, Probe: report.Probe{At: at, Successful: true}}},
		Metrics:       metrics.Snapshot{Phases: map[string]metrics.PhaseSummary{"fault": {Attempts: 10, Successes: 9, Errors: 1, DurationSeconds: 2, SuccessRate: .9, Throughput: 5, ErrorClasses: map[string]uint64{"network": 1}, SuccessfulLatency: metrics.LatencySummary{Count: 9, P50MS: float64(index), P95MS: float64(index * 2), P99MS: float64(index * 3)}, FailedLatency: metrics.LatencySummary{Count: 1, P50MS: 10, P95MS: 10, P99MS: 10}}}},
		WorkloadStats: workload.Stats{WritesAttempted: 5, AcknowledgedWrites: 4, AmbiguousWrites: 1, Guarantee: workload.Guarantee},
		Validation:    workload.Validation{StartedAt: at, CompletedAt: at.Add(time.Millisecond), Expected: 4, Checked: 4, Complete: true, Guarantee: workload.Guarantee},
		Cleanup:       report.Cleanup{Attempted: true, Successful: true}}
}

func save(t *testing.T, root string, r *report.Report) {
	t.Helper()
	dir := filepath.Join(root, r.RunID)
	if err := os.MkdirAll(dir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := report.Write(dir, *r); err != nil {
		t.Fatal(err)
	}
}

func TestFiveSequentialTrialsAgreeWithIndividualReports(t *testing.T) {
	root := t.TempDir()
	calls := 0
	cleaned := true
	s, err := Run(context.Background(), config.Default(), 5, root, func(_ context.Context, c config.Config, runsDir string) (*report.Report, error) {
		if !cleaned {
			t.Fatal("next trial started before cleanup")
		}
		cleaned = false
		calls++
		r := fixture(c, calls)
		save(t, runsDir, r)
		cleaned = true
		return r, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if calls != 5 || s.StartedTrials != 5 || s.CompletedTrials != 5 || s.UnstartedTrials != 0 || s.Outcomes["recovered"] != 5 || s.RecoverySuccessFraction != 1 || s.Outcome != "recovered" {
		t.Fatalf("trial accounting: %+v", s)
	}
	stats := s.Events["replica_promoted"]
	if stats.Observed != 5 || stats.Censored != 0 || stats.SinceFaultRequestedMS.N != 5 || *stats.SinceFaultRequestedMS.Mean != 3000 || *stats.SinceFaultRequestedMS.Median != 3000 || *stats.SinceFaultCompletedMS.Mean != 2900 || stats.SinceFaultRequestedMS.P90 != nil || stats.SinceFaultRequestedMS.P95 != nil {
		t.Fatalf("event statistics: %+v", stats)
	}
	if math.Abs(*stats.SinceFaultRequestedMS.SampleStdDev-math.Sqrt(2.5)*1000) > 1e-9 {
		t.Fatalf("sample standard deviation: %v", *stats.SinceFaultRequestedMS.SampleStdDev)
	}
	p := s.Phases["fault"]
	if p.Attempts != 50 || p.Successes != 45 || p.Errors != 5 || *p.SuccessFraction != .9 || *p.CompletedThroughput != 5 || *p.SuccessfulThroughput != 4.5 || p.TrialSuccessfulP95MS.N != 5 || *p.TrialSuccessfulP95MS.Mean != 6 {
		t.Fatalf("phase aggregation: %+v", p)
	}
	if s.Data.AcknowledgedWrites != 20 || s.Data.Checked != 20 || s.Data.AmbiguousWrites != 5 || s.Data.AcknowledgedNotChecked != 0 {
		t.Fatalf("data aggregation: %+v", s.Data)
	}
	if !filepath.IsAbs(s.Directory) || filepath.Base(s.Directory) != s.BatchID || !strings.HasPrefix(s.BatchID, "batch-") {
		t.Fatal("incorrect batch layout")
	}
	var persisted Summary
	data, err := os.ReadFile(filepath.Join(s.Directory, "summary.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &persisted); err != nil {
		t.Fatal(err)
	}
	if persisted.RequestedTrials != 5 || persisted.Data.AcknowledgedWrites != 20 {
		t.Fatal("summary JSON lost measurements")
	}
	for _, trial := range s.Trials {
		if !trial.EvidenceComplete || trial.RunID == "" || !trial.Cleanup.Successful {
			t.Fatalf("trial evidence: %+v", trial)
		}
		var raw report.Report
		data, err := os.ReadFile(filepath.Join(s.Directory, trial.ReportPath))
		if err != nil {
			t.Fatal(err)
		}
		if err := json.Unmarshal(data, &raw); err != nil {
			t.Fatal(err)
		}
		if raw.RunID != trial.RunID || raw.ConfigSHA256 != s.ConfigSHA256 {
			t.Fatal("individual report mismatch")
		}
	}
	markdown, err := os.ReadFile(filepath.Join(s.Directory, "summary.md"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"all 5 requested trials", "P90/P95 are omitted for N<20", "not pooled request quantiles", "runs/trial-0001-valid/report.json"} {
		if !strings.Contains(string(markdown), want) {
			t.Errorf("summary markdown missing %q", want)
		}
	}
}

func TestFailedPartialAndDeadlineCensoredTrialsRemain(t *testing.T) {
	calls := 0
	s, err := Run(context.Background(), config.Default(), 3, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		calls++
		r := fixture(c, calls)
		switch calls {
		case 2:
			r.Outcome = "partial"
			r.Recovery.ClientStable = nil
			r.Recovery.CombinedStable = nil
			r.Recovery.TimedOut = true
			r.Validation.Missing = 1
			r.Recovery.DataVerified = false
			r.Errors = []string{"recovery deadline"}
		case 3:
			r.Outcome = "failed"
			r.FaultRequestedAt = nil
			r.FaultCompletedAt = nil
			r.FaultActionDurationMS = nil
			r.Recovery = report.Recovery{}
			r.Metrics.Phases = nil
			r.WorkloadStats = workload.Stats{}
			r.Validation = workload.Validation{}
			r.Errors = []string{"provision failed"}
		}
		save(t, dir, r)
		if r.Outcome != "recovered" {
			return r, errors.New(r.Outcome)
		}
		return r, nil
	})
	if err == nil {
		t.Fatal("trial failures hidden")
	}
	if s.Outcomes["recovered"] != 1 || s.Outcomes["partial"] != 1 || s.Outcomes["failed"] != 1 || s.RecoverySuccessFraction != 1.0/3 || s.StartedTrials != 3 {
		t.Fatalf("outcome denominator: %+v", s)
	}
	stable := s.Events["client_stable"]
	if stable.Observed != 1 || stable.CensoredDeadline != 1 || stable.NotAtRisk != 1 || stable.SinceFaultRequestedMS.N != 1 || *stable.SinceFaultRequestedMS.Mean != 1000 || stable.SinceFaultRequestedMS.SampleStdDev != nil {
		t.Fatalf("censored event: %+v", stable)
	}
	promoted := s.Events["replica_promoted"]
	if promoted.Observed != 2 || *promoted.SinceFaultRequestedMS.Mean != 1500 {
		t.Fatal("observed promotion in partial trial dropped")
	}
	if s.Data.Missing != 1 || s.Data.TrialMissingWrites.N != 2 {
		t.Fatalf("loss evidence dropped: %+v", s.Data)
	}
	if s.Trials[1].Observations["client_stable"].ObservationHorizonAt == nil {
		t.Fatal("censor horizon absent")
	}
}

func TestCleanupFailureStopsBeforeNextTrialEvenIfNotAttempted(t *testing.T) {
	calls := 0
	s, err := Run(context.Background(), config.Default(), 5, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		calls++
		r := fixture(c, calls)
		r.Cleanup = report.Cleanup{}
		save(t, dir, r)
		return r, nil
	})
	if err == nil || calls != 1 || s.CleanupFailures != 1 || s.UnstartedTrials != 4 || s.Outcomes["failed"] != 1 || s.Outcomes["unstarted"] != 4 || s.StopReason != "cleanup_failure" {
		t.Fatalf("cleanup failure lost: calls=%d summary=%+v err=%v", calls, s, err)
	}
}

func TestNilReportFailureCreatesEvidenceAndStops(t *testing.T) {
	calls := 0
	s, err := Run(context.Background(), config.Default(), 3, t.TempDir(), func(context.Context, config.Config, string) (*report.Report, error) {
		calls++
		return nil, errors.New("cannot create run directory")
	})
	if err == nil || calls != 1 || s.IncompleteEvidenceTrials != 1 || s.UnstartedTrials != 2 || s.Trials[0].Outcome != "failed" || !s.Trials[0].SyntheticErrorReport || s.StopReason != "missing_cleanup_evidence" {
		t.Fatalf("nil report accounting: %+v, err=%v", s, err)
	}
	if _, err := os.Stat(filepath.Join(s.Directory, s.Trials[0].ReportPath)); err != nil {
		t.Fatal("nil report failure not inspectable")
	}
	if s.Events["replica_promoted"].Missing != 1 || s.Events["replica_promoted"].Unstarted != 2 || s.Events["replica_promoted"].SinceFaultRequestedMS.N != 0 {
		t.Fatal("nil report fabricated recovery samples")
	}
}

func TestCancellationPreservesAbortedAndUnstartedCounts(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	calls := 0
	s, err := Run(ctx, config.Default(), 4, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		calls++
		r := fixture(c, calls)
		r.Outcome = "aborted"
		r.Recovery.CombinedStable = nil
		r.Recovery.ClientStable = nil
		r.Errors = []string{"canceled"}
		save(t, dir, r)
		cancel()
		return r, context.Canceled
	})
	if !errors.Is(err, context.Canceled) || calls != 1 || s.Outcome != "aborted" || s.Outcomes["aborted"] != 1 || s.UnstartedTrials != 3 || s.Outcomes["unstarted"] != 3 {
		t.Fatalf("canceled accounting: %+v err=%v", s, err)
	}
	if s.Events["client_stable"].CensoredAborted != 1 {
		t.Fatal("aborted horizon confused with deadline")
	}
	preCanceled, preCancel := context.WithCancel(context.Background())
	preCancel()
	before, err := Run(preCanceled, config.Default(), 2, t.TempDir(), func(context.Context, config.Config, string) (*report.Report, error) {
		t.Fatal("callback invoked after cancellation")
		return nil, nil
	})
	if !errors.Is(err, context.Canceled) || before.StartedTrials != 0 || before.UnstartedTrials != 2 || before.Outcome != "aborted" {
		t.Fatalf("pre-canceled summary: %+v err=%v", before, err)
	}
	if _, err := os.Stat(filepath.Join(before.Directory, "summary.json")); err != nil {
		t.Fatal("canceled summary not saved")
	}
}

func TestMissingOrMismatchedReportFailsVisiblyButCleanTrialsCanContinue(t *testing.T) {
	for _, kind := range []string{"missing_markdown", "mismatch_json", "mismatch_markdown", "config_mismatch", "unsupported_schema", "negative_validation"} {
		t.Run(kind, func(t *testing.T) {
			calls := 0
			s, err := Run(context.Background(), config.Default(), 2, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
				calls++
				r := fixture(c, calls)
				if calls == 1 {
					switch kind {
					case "config_mismatch":
						r.Config.Workload.Seed++
					case "unsupported_schema":
						r.SchemaVersion = 99
					case "negative_validation":
						r.Validation.Checked = -1
					}
				}
				save(t, dir, r)
				if calls == 1 {
					switch kind {
					case "missing_markdown":
						if err := os.Remove(filepath.Join(dir, r.RunID, "report.md")); err != nil {
							t.Fatal(err)
						}
					case "mismatch_json":
						if err := os.WriteFile(filepath.Join(dir, r.RunID, "report.json"), []byte("{}"), 0600); err != nil {
							t.Fatal(err)
						}
					case "mismatch_markdown":
						if err := os.WriteFile(filepath.Join(dir, r.RunID, "report.md"), []byte("stale report"), 0600); err != nil {
							t.Fatal(err)
						}
					}
				}
				return r, nil
			})
			if err == nil || calls != 2 || s.Outcomes["failed"] != 1 || s.Outcomes["recovered"] != 1 || s.IncompleteEvidenceTrials != 1 || s.Trials[0].EvidenceComplete {
				t.Fatalf("invalid evidence hidden: %+v err=%v", s, err)
			}
			if s.Phases["fault"].Attempts != 10 || s.Data.Checked != 4 {
				t.Fatal("invalid evidence contaminated metrics")
			}
		})
	}
}

func TestDuplicateRunIDStopsBatch(t *testing.T) {
	calls := 0
	s, err := Run(context.Background(), config.Default(), 3, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		calls++
		r := fixture(c, 1)
		if calls == 1 {
			save(t, dir, r)
		}
		return r, nil
	})
	if err == nil || calls != 2 || s.Outcomes["failed"] != 1 || s.UnstartedTrials != 1 || s.StopReason != "invalid_or_duplicate_run_id" {
		t.Fatalf("duplicate ID unhandled: %+v err=%v", s, err)
	}
}

func TestSummaryWriteFailureStopsAndPreservesIndividualReport(t *testing.T) {
	calls := 0
	s, err := Run(context.Background(), config.Default(), 3, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		calls++
		r := fixture(c, calls)
		save(t, dir, r)
		if err := os.Mkdir(filepath.Join(filepath.Dir(dir), "summary.md"), 0700); err != nil {
			t.Fatal(err)
		}
		return r, nil
	})
	if err == nil || calls != 1 || s.UnstartedTrials != 2 || s.StopReason != "summary_write_failure" {
		t.Fatalf("summary failure hidden: %+v err=%v", s, err)
	}
	if _, err := os.Stat(filepath.Join(s.Directory, s.Trials[0].ReportPath)); err != nil {
		t.Fatal("summary failure destroyed individual report")
	}
	if _, err := os.Stat(filepath.Join(s.Directory, "summary.json")); err != nil {
		t.Fatal("successful first summary JSON write lost")
	}
	matches, _ := filepath.Glob(filepath.Join(s.Directory, ".summary-*"))
	if len(matches) > 0 {
		t.Fatalf("temporary summaries leaked: %v", matches)
	}
}

func TestMonotonicElapsedFieldsAndAcceptedHorizonUsed(t *testing.T) {
	r := fixture(config.Default(), 1)
	// Deliberately inconsistent wall-clock differences emulate a clock adjustment:
	// runtime elapsed fields must retain their independent measured durations.
	r.FaultCompletedAt = numberTime(r.FaultRequestedAt.Add(500 * time.Millisecond))
	r.FaultActionDurationMS = number(100)
	late := r.Samples[0]
	late.Excluded = true
	late.Probe.At = late.Probe.At.Add(time.Hour)
	r.Samples = append(r.Samples, late)
	tEvidence := trialEvidence(1, r)
	tEvidence.EvidenceComplete = true
	if *tEvidence.FaultActionDurationMS != 100 || *tEvidence.Observations["replica_promoted"].SinceFaultCompletedMS != 900 || !tEvidence.Observations["replica_promoted"].ObservationHorizonAt.Equal(r.Samples[0].Probe.At) {
		t.Fatalf("runtime timing or horizon lost: %+v", tEvidence)
	}
	r.SchemaVersion = 1
	r.FaultActionDurationMS = nil
	legacy := trialEvidence(1, r)
	if *legacy.FaultActionDurationMS != 500 || *legacy.Observations["replica_promoted"].SinceFaultCompletedMS != 500 {
		t.Fatal("schema1 fallback unavailable")
	}
}

func TestRequestWeightedAvailabilityDiffersFromTrialMean(t *testing.T) {
	a, b := fixture(config.Default(), 1), fixture(config.Default(), 2)
	a.Metrics.Phases["fault"] = metrics.PhaseSummary{Attempts: 1, Successes: 1, Errors: 0, DurationSeconds: 1}
	b.Metrics.Phases["fault"] = metrics.PhaseSummary{Attempts: 99, Successes: 0, Errors: 99, DurationSeconds: 3}
	t1, t2 := trialEvidence(1, a), trialEvidence(2, b)
	t1.EvidenceComplete = true
	t2.EvidenceComplete = true
	s := Summary{RequestedTrials: 2, Trials: []Trial{t1, t2}}
	Aggregate(&s)
	p := s.Phases["fault"]
	if *p.SuccessFraction != .01 || *p.TrialSuccessFraction.Mean != .5 || *p.CompletedThroughput != 25 || *p.SuccessfulThroughput != .25 {
		t.Fatalf("incorrect weighted availability: %+v", p)
	}
	if p.TrialSuccessfulP95MS.N != 0 || p.TrialSuccessfulP95MS.Mean != nil {
		t.Fatal("zero-count successful latency treated as observation")
	}
}

func TestLoadCountersAreValidatedAndLegacyAbsenceIsNotZero(t *testing.T) {
	calls := 0
	s, err := Run(context.Background(), config.Default(), 2, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		calls++
		r := fixture(c, calls)
		if calls == 1 {
			r.WorkloadStats.PacingMethodology = "measured no-catchup pacing"
			r.WorkloadStats.OfferedRequests = 15
			r.WorkloadStats.StartedRequests = 10
			r.WorkloadStats.CompletedRequests = 10
			r.WorkloadStats.DroppedRequests = 5
			r.Metrics.Totals.Attempts = 10
		}
		save(t, dir, r)
		return r, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if s.Load.MeasuredTrials != 1 || s.Load.UnmeasuredTrials != 1 || s.Load.OfferedRequests != 15 || s.Load.StartedRequests != 10 || s.Load.CompletedRequests != 10 || s.Load.DroppedRequests != 5 || s.Load.TrialOfferedRequests.N != 1 || *s.Load.StartedFraction != 2.0/3 {
		t.Fatalf("load accounting: %+v", s.Load)
	}
	r := fixture(config.Default(), 1)
	r.WorkloadStats.PacingMethodology = "measured"
	r.WorkloadStats.OfferedRequests = 1
	r.WorkloadStats.StartedRequests = 2
	if err := validNumbers(r); err == nil {
		t.Fatal("invalid load counters accepted")
	}
}

func TestRecoveredOutcomeNeedsSemanticEvidence(t *testing.T) {
	for _, test := range []struct {
		name   string
		mutate func(*report.Report)
	}{
		{"no_combined_recovery", func(r *report.Report) { r.Recovery.CombinedRecovered = false }},
		{"no_stable_milestone", func(r *report.Report) { r.Recovery.CombinedStable = nil }},
		{"topology_failed", func(r *report.Report) { r.Recovery.TopologyRecovered = false }},
		{"client_failed", func(r *report.Report) { r.Recovery.ClientRecovered = false }},
		{"data_unverified", func(r *report.Report) { r.Recovery.DataVerified = false }},
		{"recovery_timed_out", func(r *report.Report) { r.Recovery.TimedOut = true }},
		{"missing_write", func(r *report.Report) { r.Validation.Missing = 1 }},
		{"corrupt_write", func(r *report.Report) { r.Validation.Corrupt = 1 }},
		{"unreadable_write", func(r *report.Report) { r.Validation.Unreadable = 1 }},
		{"unchecked_write", func(r *report.Report) {
			r.Validation.Checked = 3
			r.Validation.Unverified = 1
			r.Validation.Complete = false
		}},
		{"no_validation_timestamp", func(r *report.Report) { r.Validation.StartedAt = time.Time{} }},
		{"traffic_missing", func(r *report.Report) { r.WorkloadStats.MissingReads = 1 }},
		{"traffic_mismatch", func(r *report.Report) { r.WorkloadStats.MismatchedReads = 1 }},
		{"cleanup_unverified", func(r *report.Report) { r.Cleanup.Successful = false }},
		{"cleanup_success_without_attempt", func(r *report.Report) { r.Cleanup.Attempted = false }},
	} {
		t.Run(test.name, func(t *testing.T) {
			s, err := Run(context.Background(), config.Default(), 1, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
				r := fixture(c, 1)
				test.mutate(r)
				save(t, dir, r)
				return r, nil
			})
			if err == nil || s.Outcomes["recovered"] != 0 || s.Outcomes["failed"] != 1 || s.RecoverySuccessFraction != 0 || s.IncompleteEvidenceTrials != 1 {
				t.Fatalf("malformed recovered outcome accepted: %+v err=%v", s, err)
			}
		})
	}
}

func TestSchemaOneExplicitRecoveryFallbackRemainsSupported(t *testing.T) {
	s, err := Run(context.Background(), config.Default(), 1, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		r := fixture(c, 1)
		r.SchemaVersion = 1
		r.Recovery.CombinedRecovered = false
		save(t, dir, r)
		return r, nil
	})
	if err != nil || s.Outcomes["recovered"] != 1 || s.IncompleteEvidenceTrials != 0 {
		t.Fatalf("legacy explicit evidence rejected: %+v err=%v", s, err)
	}
}

func TestCleanupSuccessWithoutAnAttemptStopsTheBatch(t *testing.T) {
	calls := 0
	s, err := Run(context.Background(), config.Default(), 2, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		calls++
		r := fixture(c, calls)
		r.Cleanup.Attempted = false
		save(t, dir, r)
		return r, nil
	})
	if err == nil || calls != 1 || s.CleanupFailures != 1 || s.Outcomes["failed"] != 1 || s.UnstartedTrials != 1 || s.StopReason != "cleanup_failure" {
		t.Fatalf("contradictory cleanup allowed next trial: %+v err=%v", s, err)
	}
}

func TestPositiveOperationCountRequiresPositivePhaseWallDuration(t *testing.T) {
	r := fixture(config.Default(), 1)
	phase := r.Metrics.Phases["fault"]
	phase.DurationSeconds = 0
	r.Metrics.Phases["fault"] = phase
	if err := validNumbers(r); err == nil {
		t.Fatal("positive completed count with zero duration accepted")
	}
	s, err := Run(context.Background(), config.Default(), 1, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		r.Config = c
		save(t, dir, r)
		return r, nil
	})
	if err == nil || s.Outcomes["failed"] != 1 || len(s.Phases) != 0 {
		t.Fatalf("zero-duration trial contaminated throughput: %+v err=%v", s, err)
	}
}

func TestStatisticsKnownValuesAndSmallSampleTailSuppression(t *testing.T) {
	empty := describe(nil)
	if empty.N != 0 || empty.Mean != nil || empty.SampleStdDev != nil || empty.P95 != nil {
		t.Fatal("empty statistics fabricated")
	}
	one := describe([]float64{0})
	if *one.Min != 0 || *one.Mean != 0 || one.SampleStdDev != nil || one.P90 != nil {
		t.Fatal("single zero observation lost")
	}
	s := describe([]float64{1, 2, 3, 4})
	if *s.Median != 2.5 || *s.Mean != 2.5 || math.Abs(*s.SampleStdDev-math.Sqrt(5.0/3)) > 1e-12 || s.P90 != nil || s.P95 != nil {
		t.Fatalf("known sample stats: %+v", s)
	}
	var twenty []float64
	for i := 1; i <= 20; i++ {
		twenty = append(twenty, float64(i))
	}
	tail := describe(twenty)
	if *tail.P90 != 18 || *tail.P95 != 19 || *tail.Median != 10.5 || tail.SmallSample {
		t.Fatalf("tail statistics: %+v", tail)
	}
}

func TestInvalidInputsDoNotInvokeTrial(t *testing.T) {
	for _, runs := range []int{0, -1, MaxRuns + 1} {
		if _, err := Run(context.Background(), config.Default(), runs, t.TempDir(), func(context.Context, config.Config, string) (*report.Report, error) {
			t.Fatal("invalid batch invoked callback")
			return nil, nil
		}); err == nil {
			t.Fatal("invalid run count accepted")
		}
	}
	if _, err := Run(context.Background(), config.Default(), 1, t.TempDir(), nil); err == nil {
		t.Fatal("nil callback accepted")
	}
	c := config.Default()
	c.Version = 99
	if _, err := Run(context.Background(), c, 1, t.TempDir(), func(context.Context, config.Config, string) (*report.Report, error) {
		t.Fatal("invalid config invoked callback")
		return nil, nil
	}); err == nil {
		t.Fatal("invalid experiment config accepted")
	}
}

func numberTime(value time.Time) *time.Time { return &value }
