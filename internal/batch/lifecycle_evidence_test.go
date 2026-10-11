package batch

import (
	"context"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/report"
	"testing"
)

func TestLifecycleBatchRejectsServiceOnlyRecoveredClaim(t *testing.T) {
	c := config.Default()
	c.Experiment.Fault = "replica-kill"
	s, err := Run(context.Background(), c, 1, t.TempDir(), func(_ context.Context, c config.Config, root string) (*report.Report, error) {
		r := fixture(c, 1)
		r.SchemaVersion = report.LifecycleSchemaVersion
		save(t, root, r)
		return r, nil
	})
	if err == nil || s.Outcomes["recovered"] != 0 || s.IncompleteEvidenceTrials != 1 {
		t.Fatalf("service-only claim accepted: %+v %v", s, err)
	}
}
func TestRepairTimeoutCensoredWithoutInventedPromotion(t *testing.T) {
	c := config.Default()
	c.Experiment.Fault = "replica-kill"
	s, err := Run(context.Background(), c, 1, t.TempDir(), func(_ context.Context, c config.Config, root string) (*report.Report, error) {
		r := fixture(c, 1)
		r.SchemaVersion = report.LifecycleSchemaVersion
		r.Outcome = "partial"
		r.Recovery.ReplicaPromoted = nil
		r.LifecycleFault = &report.LifecycleFault{RepairTimedOut: true}
		save(t, root, r)
		return r, nil
	})
	if err == nil || s.Outcomes["partial"] != 1 || s.Events["stable_redundancy"].CensoredDeadline != 1 || s.Events["stable_redundancy"].SinceFaultRequestedMS.N != 0 || s.Events["replica_promoted"].NotAtRisk != 1 {
		t.Fatalf("bad lifecycle censoring: %+v %v", s, err)
	}
}
