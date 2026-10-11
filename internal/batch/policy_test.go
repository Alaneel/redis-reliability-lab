package batch

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/Alaneel/redis-reliability-lab/internal/config"
	"github.com/Alaneel/redis-reliability-lab/internal/metrics"
	"github.com/Alaneel/redis-reliability-lab/internal/report"
	"github.com/Alaneel/redis-reliability-lab/internal/workload"
)

func waitFixture(c config.Config, index int) *report.Report {
	r := fixture(c, index)
	r.Outcome = "partial"
	unavailable := false
	r.Recovery.WritePolicyRecovered = &unavailable
	r.WorkloadStats.WritePolicy = "wait"
	r.WorkloadStats.PolicyTotals = workload.PolicyCounts{SetAcknowledged: 4, SetErrors: 1, WaitAttempted: 4, WaitSatisfied: 2, WaitShort: 1, WaitErrors: 1, WaitClientTimeouts: 1, PolicySatisfied: 2, WaitAcknowledgements: map[int64]uint64{1: 2, 0: 1}}
	r.WorkloadStats.PolicyByPhase = map[string]workload.PolicyCounts{"fault": r.WorkloadStats.PolicyTotals}
	stage := metrics.Summary{Attempts: 9, Successes: 6, Errors: 3, DurationSeconds: 2,
		ByOperation: map[string]metrics.OperationSummary{"set": {Attempts: 5, Successes: 4, Errors: 1}, "wait": {Attempts: 4, Successes: 2, Errors: 2}}}
	r.WorkloadStats.WriteStages = metrics.Snapshot{Totals: stage, Phases: map[string]metrics.PhaseSummary{"fault": stage}}
	r.Validation.WaitSatisfiedSubset = workload.ValidationSubset{Expected: 2, Checked: 2, Complete: true}
	m := r.Metrics.Phases["fault"]
	m.Successes, m.Errors, m.SuccessfulLatency.Count, m.FailedLatency.Count = 7, 3, 7, 3
	m.ByOperation = map[string]metrics.OperationSummary{"set_wait": {Attempts: 5, Successes: 2, Errors: 3}, "get": {Attempts: 5, Successes: 5}}
	r.Metrics.Totals = m
	r.Metrics.Phases["fault"] = m
	return r
}

func TestMeasuredPolicyEvidenceReconcilesPhasesHistogramsAndStages(t *testing.T) {
	for _, test := range []struct {
		name   string
		mutate func(*report.Report)
		want   string
	}{
		{"phase_missing", func(r *report.Report) { r.WorkloadStats.PolicyByPhase = nil }, "summed phase policy"},
		{"phase_redirect_missing", func(r *report.Report) {
			r.WorkloadStats.PolicyTotals.HandledRedirects = 2
		}, "summed phase policy"},
		{"phase_count_missing", func(r *report.Report) {
			p := r.WorkloadStats.PolicyByPhase["fault"]
			p.SetErrors = 0
			r.WorkloadStats.PolicyByPhase["fault"] = p
		}, "stage attempts"},
		{"reply_missing", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAcknowledgements[1]-- }, "reply histogram"},
		{"unknown_reply_invented", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAcknowledgements[0]++ }, "reply histogram"},
		{"negative_reply", func(r *report.Report) { r.WorkloadStats.PolicyTotals.WaitAcknowledgements[-1] = 1 }, "negative reply"},
		{"total_set_stage_missing", func(r *report.Report) {
			p := r.WorkloadStats.WriteStages.Totals.ByOperation["set"]
			p.Attempts--
			r.WorkloadStats.WriteStages.Totals.ByOperation["set"] = p
		}, "stage attempts"},
		{"total_wait_stage_missing", func(r *report.Report) {
			p := r.WorkloadStats.WriteStages.Totals.ByOperation["wait"]
			p.Attempts--
			r.WorkloadStats.WriteStages.Totals.ByOperation["wait"] = p
		}, "stage attempts"},
		{"stage_phase_missing", func(r *report.Report) { r.WorkloadStats.WriteStages.Phases = nil }, "stage attempts"},
		{"unexpected_stage_phase", func(r *report.Report) {
			r.WorkloadStats.WriteStages.Phases["extra"] = metrics.PhaseSummary{Attempts: 1}
		}, "summed diagnostic stage"},
		{"subset_more_missing_than_ledger", func(r *report.Report) { r.Validation.WaitSatisfiedSubset.Missing = 1 }, "exceeds the full SET ledger"},
		{"subset_more_corrupt_than_ledger", func(r *report.Report) { r.Validation.WaitSatisfiedSubset.Corrupt = 1 }, "exceeds the full SET ledger"},
		{"subset_more_unreadable_than_ledger", func(r *report.Report) { r.Validation.WaitSatisfiedSubset.Unreadable = 1 }, "exceeds the full SET ledger"},
	} {
		t.Run(test.name, func(t *testing.T) {
			c := config.Default()
			c.Workload.WritePolicy = "wait"
			r := waitFixture(c, 1)
			test.mutate(r)
			err := validPolicyNumbers(r)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("invalid evidence accepted or misclassified: %v", err)
			}
		})
	}
}

func TestSETPolicyLegacyAndMeasuredEvidenceRemainDistinct(t *testing.T) {
	c := config.Default()
	r := fixture(c, 1)
	// Absence predates the policy fields. It is unmeasured, never fabricated as
	// zero stage counts, and remains accepted for existing SET reports.
	if err := validPolicyNumbers(r); err != nil {
		t.Fatal(err)
	}
	r.WorkloadStats.WritePolicy = "set"
	r.WorkloadStats.PolicyTotals = workload.PolicyCounts{SetAcknowledged: 4, SetErrors: 1, PolicySatisfied: 4}
	r.WorkloadStats.PolicyByPhase = map[string]workload.PolicyCounts{"fault": r.WorkloadStats.PolicyTotals}
	stage := metrics.Summary{Attempts: 5, Successes: 4, Errors: 1, DurationSeconds: 2,
		ByOperation: map[string]metrics.OperationSummary{"set": {Attempts: 5, Successes: 4, Errors: 1}}}
	r.WorkloadStats.WriteStages = metrics.Snapshot{Totals: stage, Phases: map[string]metrics.PhaseSummary{"fault": stage}}
	m := r.Metrics.Phases["fault"]
	m.ByOperation = map[string]metrics.OperationSummary{"set": {Attempts: 5, Successes: 4, Errors: 1}, "get": {Attempts: 5, Successes: 5}}
	r.Metrics.Totals = m
	r.Metrics.Phases["fault"] = m
	if err := validPolicyNumbers(r); err != nil {
		t.Fatalf("valid measured SET evidence rejected: %v", err)
	}
	r.WorkloadStats.WriteStages.Totals.Attempts++
	if err := validPolicyNumbers(r); err == nil {
		t.Fatal("extra stage attempt accepted under measured SET policy")
	}
}

func TestMalformedMeasuredPolicyEvidenceDoesNotEnterBatchMeasurements(t *testing.T) {
	c := config.Default()
	c.Workload.WritePolicy = "wait"
	s, err := Run(context.Background(), c, 1, t.TempDir(), func(_ context.Context, c config.Config, root string) (*report.Report, error) {
		r := waitFixture(c, 1)
		r.WorkloadStats.PolicyByPhase = nil
		save(t, root, r)
		return r, errors.New("policy unavailable")
	})
	if err == nil || s.IncompleteEvidenceTrials != 1 || s.Outcomes["failed"] != 1 || s.WritePolicy.MeasuredTrials != 0 || s.Data.AcknowledgedWrites != 0 || s.WritePolicy.Totals.SetAcknowledged != 0 {
		t.Fatalf("unreconciled measured counters contaminated batch evidence: %+v err=%v", s, err)
	}
}

func TestPartialWAITTrialsRetainSETAndSubsetEvidence(t *testing.T) {
	c := config.Default()
	c.Workload.WritePolicy = "wait"
	calls := 0
	s, err := Run(context.Background(), c, 2, t.TempDir(), func(_ context.Context, c config.Config, root string) (*report.Report, error) {
		calls++
		r := waitFixture(c, calls)
		save(t, root, r)
		return r, errors.New("replica acknowledgement policy unavailable")
	})
	if err == nil || s.Outcomes["partial"] != 2 || s.IncompleteEvidenceTrials != 0 || s.WritePolicy.MeasuredTrials != 2 || s.WritePolicy.Totals.SetAcknowledged != 8 || s.WritePolicy.Totals.WaitSatisfied != 4 || s.WritePolicy.Totals.WaitShort != 2 || s.WritePolicy.Totals.WaitErrors != 2 || s.WritePolicy.Totals.WaitAcknowledgements[1] != 4 || s.WritePolicy.WaitSatisfiedSubset.Checked != 4 || s.Data.Checked != 8 {
		t.Fatalf("partial policy evidence omitted: %+v, %v", s, err)
	}
	if e := s.Events["write_policy_stable"]; e.Observed != 0 || e.Censored != 2 || e.SinceFaultRequestedMS.N != 0 {
		t.Fatalf("unavailable policy became zero-time recovery: %+v", e)
	}
	if !strings.Contains(s.Markdown(), "WAIT-satisfied subset validation: expected 4; checked 4") {
		t.Fatal("WAIT subset missing from readable batch evidence")
	}
}

func TestPolicyGuardRejectsFalseRecoveryAndInvalidLedger(t *testing.T) {
	c := config.Default()
	c.Workload.WritePolicy = "wait"
	r := waitFixture(c, 1)
	r.Outcome = "recovered"
	if err := validPolicyNumbers(r); err == nil {
		t.Fatal("ordinary service recovery accepted as WAIT policy recovery")
	}
	r.Outcome = "partial"
	r.Validation.WaitSatisfiedSubset.Expected = 3
	if err := validPolicyNumbers(r); err == nil {
		t.Fatal("WAIT subset inconsistent with acknowledgement counts accepted")
	}
	r = waitFixture(c, 1)
	r.WorkloadStats.PolicyTotals.WaitAttempted--
	if err := validPolicyNumbers(r); err == nil {
		t.Fatal("missing WAIT outcome accepted")
	}
}

func TestSkippedWAITSubsetValidationRemainsUnverified(t *testing.T) {
	c := config.Default()
	c.Workload.WritePolicy = "wait"
	r := waitFixture(c, 1)
	r.Validation = workload.Validation{}
	s := Summary{Config: c, RequestedTrials: 1, Trials: []Trial{{Started: true, EvidenceComplete: true, Outcome: "aborted", WorkloadStats: r.WorkloadStats}}}
	Aggregate(&s)
	if s.WritePolicy.WaitSatisfiedSubset.Expected != 2 || s.WritePolicy.WaitSatisfiedSubset.Unverified != 2 || s.WritePolicy.WaitSatisfiedSubset.Complete || s.Data.AcknowledgedNotChecked != 4 {
		t.Fatal("skipped validation erased acknowledged SET or WAIT subset")
	}
}
