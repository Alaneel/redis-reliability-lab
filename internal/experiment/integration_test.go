//go:build integration

package experiment

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/report"
)

func integrationConfig() config.Config {
	c := config.Default()
	c.Experiment.Baseline = config.Duration(2 * time.Second)
	c.Experiment.PostRecovery = config.Duration(2 * time.Second)
	c.Environment.NodeTimeout = config.Duration(time.Second)
	c.Experiment.StableFor = config.Duration(time.Second)
	c.Experiment.SampleInterval = config.Duration(200 * time.Millisecond)
	c.Timeouts.Recovery = config.Duration(20 * time.Second)
	return c
}
func runIntegration(t *testing.T, c config.Config, m EnvironmentManager) (*report.Report, error) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Minute)
	defer cancel()
	root, err := filepath.Abs("../../reports/ci")
	if err != nil {
		t.Fatal(err)
	}
	r, runErr := (Runner{Manager: m}).Run(ctx, c, root)
	if r == nil {
		t.Fatalf("no report: %v", runErr)
	}
	// In case an assertion fails or regular cleanup failed, independently retry.
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		if err := m.Cleanup(ctx, r.Environment); err != nil {
			t.Error(err)
		}
	})
	if !r.Cleanup.Attempted || !r.Cleanup.Successful {
		t.Fatalf("cleanup not verified: %+v", r.Cleanup)
	}
	found, err := environment.New().Discover(context.Background(), r.RunID)
	if err != nil {
		t.Fatal(err)
	}
	if found.NetworkID != "" || len(found.Nodes) != 0 || len(found.Volumes) != 0 {
		t.Fatalf("run leaked resources: %+v", found)
	}
	data, err := os.ReadFile(filepath.Join(root, r.RunID, "report.json"))
	if err != nil {
		t.Fatal(err)
	}
	var saved report.Report
	if err := json.Unmarshal(data, &saved); err != nil {
		t.Fatalf("invalid report round-trip: %v", err)
	}
	if saved.RunID != r.RunID || saved.Outcome != r.Outcome || saved.Cleanup.Successful != true {
		t.Fatal("report lost results")
	}
	if _, err := os.Stat(filepath.Join(root, r.RunID, "report.md")); err != nil {
		t.Fatal(err)
	}
	t.Logf("run_id=%s outcome=%s report=%s", r.RunID, r.Outcome, filepath.Join(root, r.RunID, "report.json"))
	return r, runErr
}
func TestDockerPrimaryFailure(t *testing.T) {
	r, err := runIntegration(t, integrationConfig(), environment.New())
	if !r.Recovery.CombinedRecovered || !r.Recovery.TopologyRecovered || !r.Recovery.ClientRecovered || r.Recovery.ReplicaPromoted == nil || r.Recovery.FailureDetected == nil {
		t.Fatalf("incomplete recovery: %+v", r.Recovery)
	}
	if !r.Validation.Complete || r.Validation.Expected == 0 || r.Validation.Checked != r.Validation.Expected || r.Validation.Unreadable != 0 || r.Validation.Corrupt != 0 {
		t.Fatalf("validation failed: %+v", r.Validation)
	}
	// Asynchronous replication can legitimately lose an acknowledged write.
	// Test the accuracy of classification, not a zero-loss Redis guarantee.
	if r.Validation.Missing > 0 || r.WorkloadStats.MissingReads > 0 {
		if r.Outcome != "partial" || err == nil || r.Recovery.DataVerified {
			t.Fatalf("acknowledged data loss falsely succeeded: %+v, %v", r.Recovery, err)
		}
	} else if err != nil || r.Outcome != "recovered" || !r.Recovery.DataVerified {
		t.Fatalf("verified recovery misclassified: %+v, %v", r.Recovery, err)
	}
	if r.Recovery.ProcessRestarted || r.Recovery.RedundancyRestored {
		t.Fatal("claimed unperformed recovery")
	}
	t.Logf("promoted_ms=%.2f stable_ms=%.2f acknowledged=%d", r.Recovery.ReplicaPromoted.SinceFaultMS, r.Recovery.CombinedStable.SinceFaultMS, r.Validation.Expected)
}
func TestDockerRecoveryTimeout(t *testing.T) {
	c := integrationConfig()
	c.Timeouts.Recovery = config.Duration(100 * time.Millisecond)
	r, err := runIntegration(t, c, environment.New())
	if err == nil || r.Outcome == "recovered" || !r.Recovery.TimedOut || r.Recovery.CombinedRecovered {
		t.Fatalf("deadline falsely succeeded: err=%v recovery=%+v", err, r.Recovery)
	}
}

type cancelAfterKill struct {
	*environment.Manager
	cancel context.CancelFunc
}

func (m cancelAfterKill) Kill(ctx context.Context, env *environment.Environment, n environment.Node) error {
	err := m.Manager.Kill(ctx, env, n)
	m.cancel()
	return err
}
func TestDockerCancellationCleanup(t *testing.T) {
	c := integrationConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	root, err := filepath.Abs("../../reports/ci")
	if err != nil {
		t.Fatal(err)
	}
	manager := cancelAfterKill{environment.New(), cancel}
	r, err := (Runner{Manager: manager}).Run(ctx, c, root)
	if r == nil || err == nil || r.Outcome != "aborted" || !r.Cleanup.Successful {
		t.Fatalf("cancel/cleanup result: %+v, %v", r, err)
	}
	found, err := environment.New().Discover(context.Background(), r.RunID)
	if err != nil {
		t.Fatal(err)
	}
	if found.NetworkID != "" || len(found.Nodes) > 0 {
		t.Fatal("canceled run leaked resources")
	}
}
