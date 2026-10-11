//go:build integration

package experiment

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/batch"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/report"
)

// An unsuccessful trial is evidence, not a reason to omit it or reuse its
// cluster. Two real deadline failures exercise sequential teardown and fresh
// provisioning; larger successful batches are also executed through the CLI.
func TestDockerRepeatedRecoveryTimeout(t *testing.T) {
	c := integrationConfig()
	c.Timeouts.Recovery = config.Duration(100 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	root, err := filepath.Abs("../../reports/ci")
	if err != nil {
		t.Fatal(err)
	}
	manager := environment.New()
	summary, runErr := batch.Run(ctx, c, 2, root, (Runner{Manager: manager}).Run)
	if summary == nil || runErr == nil {
		t.Fatalf("failed trials must return a summary and error: %+v %v", summary, runErr)
	}
	if summary.StartedTrials != 2 || summary.CompletedTrials != 2 || summary.UnstartedTrials != 0 || summary.Outcomes["failed"] != 2 {
		t.Fatalf("failed trials disappeared: %+v", summary)
	}
	dir := summary.Directory
	paths, err := filepath.Glob(filepath.Join(dir, "runs", "*", "report.json"))
	if err != nil || len(paths) != 2 {
		t.Fatalf("reports=%v error=%v", paths, err)
	}
	seen := map[string]bool{}
	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		var r report.Report
		if err := json.Unmarshal(data, &r); err != nil {
			t.Fatal(err)
		}
		if seen[r.RunID] || !r.Recovery.TimedOut || r.Outcome == "recovered" || !r.Cleanup.Successful {
			t.Fatalf("inaccurate trial evidence: %+v", r)
		}
		seen[r.RunID] = true
		found, err := manager.Discover(ctx, r.RunID)
		if err != nil || found.NetworkID != "" || len(found.Nodes) != 0 {
			t.Fatalf("trial resources remain: %+v %v", found, err)
		}
	}
	for _, name := range []string{"summary.json", "summary.md"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Fatal(err)
		}
	}
	t.Logf("batch_id=%s failed_trials=%d cleanup_verified=true", summary.BatchID, summary.Outcomes["failed"])
}
