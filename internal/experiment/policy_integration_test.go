//go:build integration

package experiment

import (
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
)

func TestDockerWaitPolicyPreservesSETLedgerAndReportsDegradedAcknowledgements(t *testing.T) {
	c := integrationConfig()
	c.Name = "wait-policy-primary-crash"
	c.Workload.WritePolicy = "wait"
	c.Experiment.Baseline = config.Duration(4 * time.Second)
	c.Experiment.PostRecovery = config.Duration(3 * time.Second)
	r, err := runIntegration(t, c, environment.New())
	if err == nil || r.Outcome != "partial" || r.Recovery.WritePolicyRecovered == nil || *r.Recovery.WritePolicyRecovered || !r.Recovery.CombinedRecovered || !r.Recovery.TopologyRecovered || !r.Recovery.ClientRecovered {
		t.Fatalf("service recovery confused with replica acknowledgement policy: outcome=%s recovery=%+v err=%v", r.Outcome, r.Recovery, err)
	}
	s := r.WorkloadStats
	p := s.PolicyTotals
	if s.WritePolicy != "wait" || s.PolicyByPhase["baseline"].WaitSatisfied == 0 || s.PolicyByPhase["post_recovery"].WaitShort == 0 || p.WaitAttempted == 0 || p.WaitSatisfied == 0 || p.WaitShort == 0 || p.SetAcknowledged != s.AcknowledgedWrites || p.SetAcknowledged+p.SetErrors != s.WritesAttempted || p.WaitSatisfied+p.WaitShort+p.WaitErrors != p.WaitAttempted || p.WaitClientTimeouts > p.WaitErrors {
		t.Fatalf("incorrect SET/WAIT accounting: %+v", s)
	}
	if !r.Validation.Complete || r.Validation.Expected != int(s.AcknowledgedWrites) || r.Validation.Checked != r.Validation.Expected || !r.Validation.WaitSatisfiedSubset.Complete || r.Validation.WaitSatisfiedSubset.Expected != int(p.WaitSatisfied) || r.Validation.WaitSatisfiedSubset.Checked != int(p.WaitSatisfied) || r.Validation.WaitSatisfiedSubset.Expected >= r.Validation.Expected {
		t.Fatalf("incomplete WAIT disappeared from SET validation: %+v", r.Validation)
	}
	if s.OfferedRequests != s.StartedRequests+s.DroppedRequests || s.StartedRequests != s.CompletedRequests || s.CompletedRequests != r.Metrics.Totals.Attempts {
		t.Fatal("diagnostic SET/WAIT stage commands changed logical pacing totals")
	}
	if s.WriteStages.Totals.ByOperation["set"].Attempts != s.WritesAttempted || s.WriteStages.Totals.ByOperation["wait"].Attempts != p.WaitAttempted || r.Metrics.Totals.ByOperation["set_wait"].Attempts != s.WritesAttempted {
		t.Fatal("write-stage diagnostics and logical operation counts disagree")
	}
	for _, sample := range r.Samples {
		if sample.Phase == "post_recovery" && !sample.Excluded && sample.WritePolicyProbe != nil && sample.WritePolicyProbe.Successful {
			t.Fatal("all-shard WAIT policy succeeded despite a promoted shard lacking a replica")
		}
	}
	t.Logf("set_acknowledged=%d wait_satisfied=%d wait_short=%d wait_errors=%d set_ledger_checked=%d wait_subset_checked=%d missing=%d", p.SetAcknowledged, p.WaitSatisfied, p.WaitShort, p.WaitErrors, r.Validation.Checked, r.Validation.WaitSatisfiedSubset.Checked, r.Validation.Missing)
}
