package batch

import (
	"context"
	"errors"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/metrics"
	"github.com/alanwang/redis-reliability-lab/internal/report"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

func waitAOFConfig() config.Config {
	c := config.Default()
	c.Workload.WritePolicy = "waitaof"
	c.Environment.Persistence = "aof-everysec"
	return c
}

func waitAOFFixture(c config.Config, index int) *report.Report {
	r := fixture(c, index)
	r.Outcome = "partial"
	unavailable := false
	r.Recovery.WritePolicyRecovered = &unavailable
	w := &r.WorkloadStats
	w.WritePolicy = "waitaof"
	w.WritesAttempted, w.AcknowledgedWrites, w.AmbiguousWrites = 7, 6, 1
	w.PolicyTotals = workload.PolicyCounts{
		SetAcknowledged: 6, SetErrors: 1, PolicySatisfied: 1, HandledRedirects: 2,
		WaitAOFAttempted: 5, WaitAOFSatisfied: 1, WaitAOFShort: 3, WaitAOFErrors: 1, WaitAOFClientTimeouts: 1,
		WaitAOFNotAttempted: 1, WaitAOFNotAttemptedTimeouts: 1, WaitAOFLocalSatisfied: 2, WaitAOFReplicaSatisfied: 2,
		WaitAOFReplies: map[string]uint64{"1/1": 1, "1/0": 1, "0/1": 1, "0/0": 1},
	}
	w.PolicyByPhase = map[string]workload.PolicyCounts{"fault": sumPolicyCounts(workload.PolicyCounts{}, w.PolicyTotals)}
	stage := metrics.Summary{Attempts: 12, Successes: 7, Errors: 5, DurationSeconds: 2,
		ByOperation: map[string]metrics.OperationSummary{"set": {Attempts: 7, Successes: 6, Errors: 1}, "waitaof": {Attempts: 5, Successes: 1, Errors: 4}}}
	w.WriteStages = metrics.Snapshot{Totals: stage, Phases: map[string]metrics.PhaseSummary{"fault": stage}}
	m := r.Metrics.Phases["fault"]
	m.Successes, m.Errors, m.SuccessfulLatency.Count, m.FailedLatency.Count = 4, 6, 4, 6
	m.ByOperation = map[string]metrics.OperationSummary{"set_waitaof": {Attempts: 7, Successes: 1, Errors: 6}, "get": {Attempts: 3, Successes: 3}}
	r.Metrics.Totals = m
	r.Metrics.Phases["fault"] = m
	r.Validation.Expected, r.Validation.Checked = 6, 6
	r.Validation.WaitAOFSatisfiedSubset = workload.ValidationSubset{Expected: 1, Checked: 1, Complete: true}
	return r
}

func TestWAITAOFIndependentPairedRepliesAndStageAccounting(t *testing.T) {
	r := waitAOFFixture(waitAOFConfig(), 1)
	if err := validPolicyNumbers(r); err != nil {
		t.Fatalf("valid paired local/replica observations rejected: %v", err)
	}
	if r.WorkloadStats.WriteStages.Totals.Attempts != 12 || r.Metrics.Totals.Attempts != 10 {
		t.Fatal("WAITAOF diagnostic stages became additional logical requests")
	}
	p := r.WorkloadStats.PolicyTotals
	if p.WaitAOFLocalSatisfied != 2 || p.WaitAOFReplicaSatisfied != 2 || p.WaitAOFSatisfied != 1 || p.WaitAOFShort != 3 || len(p.WaitAOFReplies) != 4 {
		t.Fatal("overlapping independent acknowledgements were flattened")
	}
}

func TestWAITAOFMalformedEvidenceRejected(t *testing.T) {
	for _, test := range []struct {
		name   string
		change func(*report.Report)
		want   string
	}{
		{"missing_policy", func(r *report.Report) { r.WorkloadStats.WritePolicy = "" }, "lacks policy evidence"},
		{"unaccounted_set", func(r *report.Report) {
			r.WorkloadStats.PolicyTotals.WaitAOFNotAttempted, r.WorkloadStats.PolicyTotals.WaitAOFNotAttemptedTimeouts = 0, 0
		}, "every successful SET"},
		{"extra_attempt", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAOFAttempted++ }, "do not reconcile"},
		{"false_success", func(r *report.Report) { r.WorkloadStats.PolicyTotals.PolicySatisfied++ }, "policy outcome"},
		{"invented_local", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAOFLocalSatisfied++ }, "paired reply histogram"},
		{"invented_replica", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAOFReplicaSatisfied++ }, "paired reply histogram"},
		{"timeout_not_error", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAOFClientTimeouts++ }, "do not reconcile"},
		{"not_attempted_timeout", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAOFNotAttemptedTimeouts++ }, "do not reconcile"},
		{"unknown_as_zero", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAOFReplies["0/0"]++ }, "reply histogram"},
		{"histogram_wrap", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAOFReplies["0/0"] = math.MaxUint64 }, "excess observations"},
		{"zero_frequency", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAOFReplies["1/2"] = 0 }, "zero-frequency"},
		{"phase_missing", func(r *report.Report) { r.WorkloadStats.PolicyByPhase = nil }, "summed phase policy"},
		{"phase_histogram_missing", func(r *report.Report) {
			p := r.WorkloadStats.PolicyByPhase["fault"]
			delete(p.WaitAOFReplies, "0/1")
		}, "reply histogram"},
		{"stage_missing", func(r *report.Report) {
			p := r.WorkloadStats.WriteStages.Totals.ByOperation["waitaof"]
			p.Attempts--
			r.WorkloadStats.WriteStages.Totals.ByOperation["waitaof"] = p
		}, "stage attempts"},
		{"foreign_stage", func(r *report.Report) {
			r.WorkloadStats.WriteStages.Totals.ByOperation["wait"] = metrics.OperationSummary{Attempts: 1}
		}, "stage attempts"},
		{"unknown_stage", func(r *report.Report) {
			r.WorkloadStats.WriteStages.Totals.ByOperation["extra"] = metrics.OperationSummary{Attempts: 1}
		}, "stage attempts"},
		{"extra_stage_phase", func(r *report.Report) {
			r.WorkloadStats.WriteStages.Phases["extra"] = metrics.PhaseSummary{Attempts: 1}
		}, "summed diagnostic stage"},
		{"foreign_wait_counter", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitErrors = 1 }, "different acknowledgement policy"},
		{"foreign_wait_histogram", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAcknowledgements = map[int64]uint64{0: 1} }, "different acknowledgement policy"},
		{"subset_expected", func(r *report.Report) { r.Validation.WaitAOFSatisfiedSubset.Expected++ }, "does not reconcile"},
		{"subset_unverified", func(r *report.Report) { r.Validation.WaitAOFSatisfiedSubset.Unverified++ }, "does not reconcile"},
		{"subset_negative", func(r *report.Report) { r.Validation.WaitAOFSatisfiedSubset.Missing = -1 }, "does not reconcile"},
		{"subset_more_missing", func(r *report.Report) { r.Validation.WaitAOFSatisfiedSubset.Missing++ }, "exceeds the full SET ledger"},
		{"subset_more_corrupt", func(r *report.Report) { r.Validation.WaitAOFSatisfiedSubset.Corrupt++ }, "exceeds the full SET ledger"},
		{"subset_more_unreadable", func(r *report.Report) { r.Validation.WaitAOFSatisfiedSubset.Unreadable++ }, "exceeds the full SET ledger"},
		{"foreign_subset", func(r *report.Report) { r.Validation.WaitSatisfiedSubset.Expected = 1 }, "different acknowledgement policy"},
	} {
		t.Run(test.name, func(t *testing.T) {
			r := waitAOFFixture(waitAOFConfig(), 1)
			test.change(r)
			if err := validPolicyNumbers(r); err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("malformed evidence accepted or misclassified: %v", err)
			}
		})
	}
	for _, reply := range []string{"", "1", "1/1/1", "-1/1", "2/1", "1/-1", "+1/1", "01/1", "1/01", "1/+1", "1/ 1", "1/9223372036854775808"} {
		t.Run("noncanonical_"+reply, func(t *testing.T) {
			r := waitAOFFixture(waitAOFConfig(), 1)
			delete(r.WorkloadStats.PolicyTotals.WaitAOFReplies, "1/1")
			r.WorkloadStats.PolicyTotals.WaitAOFReplies[reply] = 1
			if err := validPolicyNumbers(r); err == nil || !strings.Contains(err.Error(), "noncanonical") {
				t.Fatalf("invalid pair %q accepted: %v", reply, err)
			}
		})
	}
}

func TestWAITAOFRecoveryRequiresStablePolicyAndPreservedSubset(t *testing.T) {
	recovered := func() *report.Report {
		r := waitAOFFixture(waitAOFConfig(), 1)
		r.Outcome = "recovered"
		ok := true
		r.Recovery.WritePolicyRecovered = &ok
		r.Recovery.WritePolicyStable = &report.Milestone{At: r.FinishedAt.Add(-time.Millisecond), SinceFaultMS: 999}
		return r
	}
	if err := validPolicyNumbers(recovered()); err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		name   string
		change func(*report.Report)
	}{
		{"ordinary_only", func(r *report.Report) { r.Recovery.WritePolicyRecovered = nil }},
		{"not_stable", func(r *report.Report) { r.Recovery.WritePolicyStable = nil }},
		{"zero_stable_time", func(r *report.Report) { r.Recovery.WritePolicyStable.At = time.Time{} }},
		{"validation_skipped", func(r *report.Report) { r.Validation.StartedAt = time.Time{} }},
		{"validation_not_finished", func(r *report.Report) { r.Validation.CompletedAt = time.Time{} }},
		{"lost_satisfied_record", func(r *report.Report) {
			r.Validation.Missing = 1
			r.Validation.WaitAOFSatisfiedSubset.Missing = 1
		}},
		{"incomplete_subset", func(r *report.Report) {
			r.Validation.Checked, r.Validation.Unverified, r.Validation.Complete = 5, 1, false
			r.Validation.WaitAOFSatisfiedSubset.Checked, r.Validation.WaitAOFSatisfiedSubset.Unverified, r.Validation.WaitAOFSatisfiedSubset.Complete = 0, 1, false
		}},
	} {
		t.Run(test.name, func(t *testing.T) {
			r := recovered()
			test.change(r)
			if err := validPolicyNumbers(r); err == nil {
				t.Fatal("recovered policy accepted without stable complete preservation")
			}
		})
	}
}

func TestWAITAOFPartialEvidencePersistsAndCensorsUnknownRecovery(t *testing.T) {
	c := waitAOFConfig()
	calls := 0
	s, err := Run(context.Background(), c, 2, t.TempDir(), func(_ context.Context, c config.Config, root string) (*report.Report, error) {
		calls++
		r := waitAOFFixture(c, calls)
		r.Errors = []string{"AOF acknowledgement policy unavailable"}
		save(t, root, r)
		return r, errors.New(r.Errors[0])
	})
	if err == nil || calls != 2 || s.IncompleteEvidenceTrials != 0 || s.Outcomes["partial"] != 2 || s.WritePolicy.MeasuredTrials != 2 {
		t.Fatalf("safe partial trials omitted: %+v, %v", s, err)
	}
	p, v := s.WritePolicy.Totals, s.WritePolicy.WaitAOFSatisfiedSubset
	if p.SetAcknowledged != 12 || p.WaitAOFAttempted != 10 || p.WaitAOFSatisfied != 2 || p.WaitAOFShort != 6 || p.WaitAOFErrors != 2 || p.WaitAOFLocalSatisfied != 4 || p.WaitAOFReplicaSatisfied != 4 || p.WaitAOFNotAttempted != 2 || p.WaitAOFReplies["0/0"] != 2 || v.Checked != 2 || !v.Complete || s.WritePolicy.WaitAOFSubsetValidationTrials != 2 || s.Data.Checked != 12 {
		t.Fatalf("independent AOF evidence lost: %+v", s.WritePolicy)
	}
	if e := s.Events["write_policy_stable"]; e.CensoredError != 2 || e.NotAtRisk != 0 || e.SinceFaultRequestedMS.N != 0 {
		t.Fatalf("unknown policy became zero-time or not-at-risk: %+v", e)
	}
	if s.Phases["fault"].Attempts != 20 {
		t.Fatal("diagnostic WAITAOF calls inflated logical attempts")
	}
	markdown, readErr := os.ReadFile(filepath.Join(s.Directory, "summary.md"))
	if readErr != nil {
		t.Fatal(readErr)
	}
	for _, want := range []string{"WAITAOF attempted", "Local satisfied", "Replica satisfied", "| 0/0 | 2 |", "unknown acknowledgement counts", "WAITAOF-satisfied subset validation: expected 2; checked 2", "Validation attempted in 2 trials"} {
		if !strings.Contains(string(markdown), want) {
			t.Fatalf("readable evidence missing %q", want)
		}
	}
	if strings.Contains(string(markdown), "| WAIT attempted |") {
		t.Fatal("WAITAOF replies presented as replica-memory WAIT")
	}
}

func TestWAITAOFMalformedEvidenceCannotContaminateAggregate(t *testing.T) {
	s, err := Run(context.Background(), waitAOFConfig(), 1, t.TempDir(), func(_ context.Context, c config.Config, root string) (*report.Report, error) {
		r := waitAOFFixture(c, 1)
		r.WorkloadStats.PolicyTotals.WaitAOFReplies["0/0"]++
		save(t, root, r)
		return r, errors.New("partial policy")
	})
	if err == nil || s.IncompleteEvidenceTrials != 1 || s.Outcomes["failed"] != 1 || s.WritePolicy.MeasuredTrials != 0 || s.WritePolicy.Totals.WaitAOFAttempted != 0 || s.Data.Checked != 0 || len(s.Phases) != 0 {
		t.Fatalf("malformed policy contaminated aggregate: %+v err=%v", s, err)
	}
}

func TestWAITAOFEarlyFailedTrialRemainsVerifiedAndUnmeasured(t *testing.T) {
	calls := 0
	s, err := Run(context.Background(), waitAOFConfig(), 2, t.TempDir(), func(_ context.Context, c config.Config, root string) (*report.Report, error) {
		calls++
		if calls == 1 {
			r := fixture(c, calls)
			r.Outcome = "failed"
			r.WorkloadStats, r.Validation, r.Metrics, r.Recovery = workload.Stats{}, workload.Validation{}, metrics.Snapshot{}, report.Recovery{}
			r.Errors = []string{"runtime compatibility unavailable before workload"}
			save(t, root, r)
			return r, errors.New(r.Errors[0])
		}
		r := waitAOFFixture(c, calls)
		save(t, root, r)
		return r, errors.New("partial policy")
	})
	if err == nil || calls != 2 || s.IncompleteEvidenceTrials != 0 || s.Outcomes["failed"] != 1 || s.Outcomes["partial"] != 1 || s.WritePolicy.MeasuredTrials != 1 || s.WritePolicy.UnmeasuredTrials != 1 || s.Data.AcknowledgedWrites != 6 || !s.Trials[0].EvidenceComplete || !s.Trials[0].Cleanup.Successful {
		t.Fatalf("pretraffic failure became corrupt evidence or stopped safe trials: %+v, %v", s, err)
	}
}

func TestSkippedWAITAOFValidationRemainsUnverified(t *testing.T) {
	c := waitAOFConfig()
	r := waitAOFFixture(c, 1)
	s := Summary{Config: c, RequestedTrials: 1, Trials: []Trial{{Started: true, EvidenceComplete: true, Outcome: "aborted", WorkloadStats: r.WorkloadStats}}}
	Aggregate(&s)
	v := s.WritePolicy.WaitAOFSatisfiedSubset
	if v.Expected != 1 || v.Unverified != 1 || v.Complete || s.WritePolicy.WaitAOFSubsetValidationTrials != 0 || s.Data.AcknowledgedNotChecked != 6 {
		t.Fatal("skipped validation erased the acknowledged subset")
	}
}

func TestMeasuredLogicalWritesMustReconcileWithPolicyOutcomes(t *testing.T) {
	for _, policy := range []string{"set", "wait", "waitaof"} {
		for _, change := range []string{"extra_attempt", "invented_success", "missing_error", "diagnostic_as_logical", "phase_missing", "mixed_summary_success"} {
			t.Run(policy+"_"+change, func(t *testing.T) {
				c := config.Default()
				c.Workload.WritePolicy = policy
				var r *report.Report
				if policy == "waitaof" {
					r = waitAOFFixture(waitAOFConfig(), 1)
				} else {
					r = waitFixture(c, 1)
					if policy == "set" {
						r.WorkloadStats.WritePolicy = "set"
						r.WorkloadStats.PolicyTotals = workload.PolicyCounts{SetAcknowledged: 4, SetErrors: 1, PolicySatisfied: 4}
						r.WorkloadStats.PolicyByPhase = map[string]workload.PolicyCounts{"fault": r.WorkloadStats.PolicyTotals}
						stage := metrics.Summary{Attempts: 5, ByOperation: map[string]metrics.OperationSummary{"set": {Attempts: 5}}}
						r.WorkloadStats.WriteStages = metrics.Snapshot{Totals: stage, Phases: map[string]metrics.PhaseSummary{"fault": stage}}
						r.Validation.WaitSatisfiedSubset = workload.ValidationSubset{}
						m := r.Metrics.Phases["fault"]
						m.Successes, m.Errors = 9, 1
						m.ByOperation = map[string]metrics.OperationSummary{"set": {Attempts: 5, Successes: 4, Errors: 1}, "get": {Attempts: 5, Successes: 5}}
						r.Metrics.Totals, r.Metrics.Phases["fault"] = m, m
					}
				}
				name := logicalWriteName(policy)
				m := r.Metrics.Phases["fault"]
				op := m.ByOperation[name]
				switch change {
				case "extra_attempt":
					op.Attempts++
				case "invented_success":
					op.Successes++
				case "missing_error":
					op.Errors--
				case "diagnostic_as_logical":
					m.ByOperation["waitaof"] = metrics.OperationSummary{Attempts: 1, Successes: 1}
				case "phase_missing":
					r.Metrics.Phases = nil
				case "mixed_summary_success":
					m.Successes++
					m.Errors--
					r.Metrics.Phases["fault"] = m
				}
				if change != "phase_missing" && change != "mixed_summary_success" {
					m.ByOperation[name] = op
					r.Metrics.Totals, r.Metrics.Phases["fault"] = m, m
				}
				if err := validPolicyNumbers(r); err == nil || !strings.Contains(err.Error(), "logical write") {
					t.Fatalf("inconsistent logical population accepted: %v", err)
				}
			})
		}
	}
}

func TestPolicyHistogramSummationDoesNotMutateSourceAndNormalizesNil(t *testing.T) {
	a := workload.PolicyCounts{WaitAcknowledgements: map[int64]uint64{1: 2}, WaitAOFReplies: map[string]uint64{"1/1": 3}}
	b := workload.PolicyCounts{WaitAcknowledgements: map[int64]uint64{1: 4}, WaitAOFReplies: map[string]uint64{"1/1": 5}, WaitAOFAttempted: 5, WaitAOFNotAttempted: 2, WaitAOFNotAttemptedTimeouts: 1, WaitAOFSatisfied: 3, WaitAOFShort: 1, WaitAOFErrors: 1, WaitAOFClientTimeouts: 1, WaitAOFLocalSatisfied: 4, WaitAOFReplicaSatisfied: 3}
	sum := sumPolicyCounts(a, b)
	sum.WaitAcknowledgements[1]++
	sum.WaitAOFReplies["1/1"]++
	if a.WaitAcknowledgements[1] != 2 || b.WaitAcknowledgements[1] != 4 || a.WaitAOFReplies["1/1"] != 3 || b.WaitAOFReplies["1/1"] != 5 {
		t.Fatal("aggregation mutated individual evidence histograms")
	}
	if sum.WaitAOFAttempted != 5 || sum.WaitAOFNotAttempted != 2 || sum.WaitAOFNotAttemptedTimeouts != 1 || sum.WaitAOFSatisfied != 3 || sum.WaitAOFShort != 1 || sum.WaitAOFErrors != 1 || sum.WaitAOFClientTimeouts != 1 || sum.WaitAOFLocalSatisfied != 4 || sum.WaitAOFReplicaSatisfied != 3 || !equalPolicyCounts(workload.PolicyCounts{}, sumPolicyCounts(workload.PolicyCounts{}, workload.PolicyCounts{})) {
		t.Fatal("AOF counters or nil histogram normalization omitted")
	}
	for _, policy := range []string{"set", "wait"} {
		t.Run("reject_aof_counts_under_"+policy, func(t *testing.T) {
			p := workload.PolicyCounts{WaitAOFLocalSatisfied: 1}
			if err := validPolicyCounts(p, policy, 1); err == nil {
				t.Fatal("cross-policy AOF counter accepted")
			}
		})
	}
}
