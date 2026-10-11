package comparison

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/batch"
	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/metrics"
	"github.com/alanwang/redis-reliability-lab/internal/report"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

func configurations(n int) []config.Config {
	var out []config.Config
	for i := 1; i <= n; i++ {
		c := config.Default()
		c.Name = fmt.Sprintf("condition-%d", i)
		c.Workload.Seed = int64(i * 101)
		out = append(out, c)
	}
	return out
}

func fixture(c config.Config, index int) *report.Report {
	start := time.Unix(int64(10000+index*10), 0).UTC()
	requested := start.Add(time.Second)
	completed := requested.Add(100 * time.Millisecond)
	at := requested.Add(time.Duration(c.Workload.Seed) * time.Millisecond)
	data, _ := json.Marshal(c)
	hash := sha256.Sum256(data)
	m := &report.Milestone{At: at, SinceFaultMS: float64(c.Workload.Seed)}
	actionDuration := 100.0
	return &report.Report{SchemaVersion: 2, RunID: fmt.Sprintf("compare-%04d-valid", index), Outcome: "recovered", StartedAt: start, FinishedAt: at.Add(time.Second), Config: c, ConfigSHA256: hex.EncodeToString(hash[:]),
		Host: map[string]string{"os": "fixture-os"}, Environment: &environment.Environment{RunID: fmt.Sprintf("compare-%04d-valid", index), Metadata: map[string]string{"redis_image_id": "fixture-image"}},
		FaultRequestedAt: &requested, FaultCompletedAt: &completed, FaultActionDurationMS: &actionDuration,
		Recovery:      report.Recovery{FailureDetected: m, ReplicaPromoted: m, FullCoverage: m, TopologyConverged: m, ClientFirstSuccess: m, ClientStable: m, CombinedStable: m, TopologyRecovered: true, ClientRecovered: true, CombinedRecovered: true, DataVerified: true},
		Samples:       []report.Sample{{Snapshot: cluster.Snapshot{At: at.Add(-time.Millisecond), Completed: at}, Probe: report.Probe{At: at, Successful: true}}},
		Metrics:       metrics.Snapshot{Phases: map[string]metrics.PhaseSummary{"fault": {Attempts: 10, Successes: 9, Errors: 1, DurationSeconds: 2, SuccessfulLatency: metrics.LatencySummary{Count: 9, P50MS: float64(c.Workload.Seed), P95MS: float64(c.Workload.Seed * 2), P99MS: float64(c.Workload.Seed * 3)}}}},
		WorkloadStats: workload.Stats{WritesAttempted: 5, AcknowledgedWrites: 4, AmbiguousWrites: 1}, Validation: workload.Validation{StartedAt: at, CompletedAt: at.Add(time.Millisecond), Expected: 4, Checked: 4, Complete: true}, Cleanup: report.Cleanup{Attempted: true, Successful: true}}
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

func TestCounterbalancedOrderIsDeterministicAndPreservesWorkloadSeeds(t *testing.T) {
	configs := configurations(3)
	original := append([]config.Config(nil), configs...)
	var schedules [][]int
	for repeat := 0; repeat < 2; repeat++ {
		calls := 0
		var actual []string
		s, err := Run(context.Background(), configs, 3, 20261011, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
			calls++
			actual = append(actual, c.Name)
			r := fixture(c, calls)
			save(t, dir, r)
			return r, nil
		})
		if err != nil {
			t.Fatal(err)
		}
		if calls != 9 || s.Outcome != "recovered" || s.StartedTrials != 9 || s.CompletedTrials != 9 || s.UnstartedTrials != 0 || s.OrderSeed != 20261011 {
			t.Fatalf("accounting: %+v", s)
		}
		balance := make([][]int, 3)
		for i := range balance {
			balance[i] = make([]int, 3)
		}
		var order []int
		for i, entry := range s.Schedule {
			order = append(order, entry.ConditionIndex)
			balance[entry.ConditionIndex-1][entry.RoundPosition-1]++
			if entry.Position != i+1 || entry.Round != i/3+1 || entry.RoundPosition != i%3+1 || actual[i] != configs[entry.ConditionIndex-1].Name {
				t.Fatalf("actual schedule: %+v", entry)
			}
			if entry.Trial.ConfigSHA256 != s.Conditions[entry.ConditionIndex-1].ConfigHash || !entry.Trial.EvidenceComplete {
				t.Fatal("condition evidence not verified")
			}
		}
		for _, positions := range balance {
			if !reflect.DeepEqual(positions, []int{1, 1, 1}) {
				t.Fatalf("unbalanced: %v", balance)
			}
		}
		for i, c := range s.Conditions {
			if c.Config != original[i] || c.Summary.Config.Workload.Seed != original[i].Workload.Seed || c.Summary.RequestedTrials != 3 || c.Summary.Outcomes["recovered"] != 3 {
				t.Fatalf("config changed: %+v", c)
			}
			p := c.Summary.Phases["fault"].TrialSuccessfulP95MS
			if p.N != 3 || p.Mean == nil || *p.Mean != float64(c.Config.Workload.Seed*2) || p.P90 != nil || p.P95 != nil {
				t.Fatalf("condition latency statistics pooled or tails invented: %+v", p)
			}
		}
		if !reflect.DeepEqual(configs, original) {
			t.Fatal("input configs mutated")
		}
		schedules = append(schedules, order)
	}
	if !reflect.DeepEqual(schedules[0], schedules[1]) {
		t.Fatalf("order seed nondeterministic: %v", schedules)
	}
}

func TestIncompleteBlockRecordsPartialPositionBalance(t *testing.T) {
	calls := 0
	s, err := Run(context.Background(), configurations(3), 2, 7, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		calls++
		r := fixture(c, calls)
		save(t, dir, r)
		return r, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for c := 1; c <= 3; c++ {
		positions := map[int]bool{}
		for _, entry := range s.Schedule {
			if entry.ConditionIndex == c {
				if positions[entry.RoundPosition] {
					t.Fatal("position reused in incomplete block")
				}
				positions[entry.RoundPosition] = true
			}
		}
		if len(positions) != 2 {
			t.Fatal("requested round lost")
		}
	}
	if !strings.Contains(strings.Join(s.Methodology, " "), "incomplete block") || !strings.Contains(s.Markdown(), "does not balance all predecessor/successor pairs") {
		t.Fatal("balance/carryover limitation missing")
	}
}

func TestFailedPartialAndNoEnvironmentTrialsContinueWithoutCensoredZeros(t *testing.T) {
	calls := 0
	s, err := Run(context.Background(), configurations(3), 3, 29, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		calls++
		r := fixture(c, calls)
		switch c.Name {
		case "condition-2":
			r.Outcome = "partial"
			r.Recovery.CombinedStable, r.Recovery.ClientStable = nil, nil
			r.Recovery.CombinedRecovered, r.Recovery.ClientRecovered, r.Recovery.TimedOut = false, false, true
		case "condition-3":
			r.Outcome = "failed"
			r.Environment = nil
			r.Cleanup = report.Cleanup{}
			r.FaultRequestedAt, r.FaultCompletedAt, r.FaultActionDurationMS = nil, nil, nil
			r.Recovery = report.Recovery{}
			r.Samples = nil
		}
		save(t, dir, r)
		if r.Outcome != "recovered" {
			return r, fmt.Errorf("fixture outcome: %s", r.Outcome)
		}
		return r, nil
	})
	if err == nil || calls != 9 || s.Outcome != "partial" || s.StopReason != "" || s.UnstartedTrials != 0 {
		t.Fatalf("safe failed/partial trials dropped: %+v %v", s, err)
	}
	if s.Conditions[0].Summary.Outcomes["recovered"] != 3 || s.Conditions[1].Summary.Outcomes["partial"] != 3 || s.Conditions[2].Summary.Outcomes["failed"] != 3 {
		t.Fatal("condition outcomes mixed")
	}
	e := s.Conditions[1].Summary.Events["combined_stable"]
	if e.CensoredDeadline != 3 || e.SinceFaultRequestedMS.N != 0 || e.SinceFaultRequestedMS.Mean != nil {
		t.Fatalf("censoring became zero: %+v", e)
	}
	if e := s.Conditions[2].Summary.Events["combined_stable"]; e.NotAtRisk != 3 || e.SinceFaultRequestedMS.N != 0 {
		t.Fatalf("setup failure invented recovery: %+v", e)
	}
	for _, entry := range s.Schedule {
		if entry.ConditionIndex == 3 && entry.EnvironmentReference != "" {
			t.Fatal("no-environment report received a nonexistent environment reference")
		}
	}
}

func TestUnsafeMissingOrUnverifiedEvidenceStopsRemainingSchedule(t *testing.T) {
	for _, fault := range []string{"cleanup", "nil-report", "markdown-mismatch", "duplicate-run-id"} {
		t.Run(fault, func(t *testing.T) {
			calls := 0
			s, err := Run(context.Background(), configurations(2), 3, 43, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
				calls++
				if fault == "nil-report" {
					return nil, errors.New("callback lost result")
				}
				r := fixture(c, calls)
				if fault == "cleanup" {
					r.Cleanup.Successful = false
				}
				if fault == "duplicate-run-id" {
					r.RunID = "reused-valid-run"
				}
				save(t, dir, r)
				if fault == "markdown-mismatch" {
					if err := os.WriteFile(filepath.Join(dir, r.RunID, "report.md"), []byte("altered report"), 0600); err != nil {
						t.Fatal(err)
					}
				}
				return r, nil
			})
			want := 1
			if fault == "duplicate-run-id" {
				want = 2
			}
			if err == nil || calls != want || s.StartedTrials != want || s.UnstartedTrials != 6-want || s.StopReason == "" || s.Outcome != "failed" {
				t.Fatalf("unsafe evidence did not stop: calls=%d %+v %v", calls, s, err)
			}
			for _, entry := range s.Schedule[want:] {
				if entry.Trial.Started || entry.Trial.Outcome != "unstarted" || entry.BatchID != "" || entry.Trial.RunID != "" {
					t.Fatalf("unstarted evidence invented: %+v", entry)
				}
			}
		})
	}
}

func TestCancellationPreservesWholeScheduleAndNoFurtherCallbacks(t *testing.T) {
	for _, before := range []bool{false, true} {
		t.Run(fmt.Sprintf("before=%t", before), func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if before {
				cancel()
			}
			calls := 0
			s, err := Run(ctx, configurations(2), 2, 11, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
				calls++
				r := fixture(c, calls)
				save(t, dir, r)
				cancel()
				return r, nil
			})
			want := 1
			if before {
				want = 0
			}
			if !errors.Is(err, context.Canceled) || calls != want || s.Outcome != "aborted" || s.StopReason != "cancellation" || len(s.Schedule) != 4 || s.UnstartedTrials != 4-want {
				t.Fatalf("cancellation lost schedule: %+v %v", s, err)
			}
			data, err := os.ReadFile(filepath.Join(s.Directory, "comparison.json"))
			if err != nil {
				t.Fatal(err)
			}
			var persisted Summary
			if err := json.Unmarshal(data, &persisted); err != nil || persisted.Outcome != "aborted" || len(persisted.Schedule) != 4 {
				t.Fatalf("cancellation not saved: %+v %v", persisted, err)
			}
		})
	}
}

func TestComparisonAndConditionFilesHaveResolvableRelativeReferences(t *testing.T) {
	calls := 0
	s, err := Run(context.Background(), configurations(2), 2, 91, t.TempDir(), func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		calls++
		r := fixture(c, calls)
		save(t, dir, r)
		return r, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range s.Schedule {
		if filepath.IsAbs(entry.ReportPath) || filepath.IsAbs(entry.BatchSummaryPath) || entry.Host["os"] != "fixture-os" || entry.EnvironmentMetadata["redis_image_id"] != "fixture-image" || entry.HostEvidenceReference != entry.ReportPath+"#/host" {
			t.Fatalf("relative/environment refs lost: %+v", entry)
		}
		for _, path := range []string{entry.ReportPath, entry.BatchSummaryPath} {
			if _, err := os.Stat(filepath.Join(s.Directory, path)); err != nil {
				t.Fatal(err)
			}
		}
	}
	for _, c := range s.Conditions {
		data, err := os.ReadFile(filepath.Join(s.Directory, c.SummaryPath))
		if err != nil {
			t.Fatal(err)
		}
		var persisted batch.Summary
		if err := json.Unmarshal(data, &persisted); err != nil {
			t.Fatal(err)
		}
		if persisted.RequestedTrials != 2 || persisted.StartedTrials != 2 || persisted.ConfigSHA256 != c.ConfigHash {
			t.Fatal("condition summary mismatches verified grouped evidence")
		}
		for _, trial := range persisted.Trials {
			if filepath.IsAbs(trial.ReportPath) {
				t.Fatal("absolute condition reference")
			}
			if _, err := os.Stat(filepath.Join(c.Summary.Directory, trial.ReportPath)); err != nil {
				t.Fatal(err)
			}
		}
	}
	data, err := os.ReadFile(filepath.Join(s.Directory, "comparison.json"))
	if err != nil {
		t.Fatal(err)
	}
	var persisted Summary
	if err := json.Unmarshal(data, &persisted); err != nil || persisted.ComparisonID != s.ComparisonID || persisted.Conditions[0].ConfigHash != s.Conditions[0].ConfigHash {
		t.Fatalf("JSON roundtrip: %v", err)
	}
	md, err := os.ReadFile(filepath.Join(s.Directory, "comparison.md"))
	if err != nil || string(md) != s.Markdown() || !strings.Contains(string(md), "No pooled request latency quantiles or causal inference") {
		t.Fatalf("Markdown mismatch: %v", err)
	}
	if err := filepath.WalkDir(s.Directory, func(path string, entry os.DirEntry, err error) error {
		if err == nil && strings.HasPrefix(entry.Name(), ".comparison-") {
			t.Errorf("temporary summary remained: %s", path)
		}
		return err
	}); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidComparisonRejectedBeforeCallbacksOrFiles(t *testing.T) {
	for _, test := range []struct{ conditions, rounds int }{{0, 1}, {1, 1}, {9, 1}, {2, 0}, {2, 101}, {2, 51}, {8, 13}} {
		t.Run(fmt.Sprintf("%dx%d", test.conditions, test.rounds), func(t *testing.T) {
			root := filepath.Join(t.TempDir(), "not-created")
			calls := 0
			s, err := Run(context.Background(), configurations(test.conditions), test.rounds, 1, root, func(context.Context, config.Config, string) (*report.Report, error) { calls++; return nil, nil })
			if err == nil || s != nil || calls != 0 {
				t.Fatal("invalid comparison accepted")
			}
			if _, err := os.Stat(root); !os.IsNotExist(err) {
				t.Fatal("invalid comparison created output")
			}
		})
	}
	cs := configurations(2)
	cs[1].Environment.RedisImage = "redis:latest"
	if s, err := Run(context.Background(), cs, 1, 1, t.TempDir(), func(context.Context, config.Config, string) (*report.Report, error) {
		t.Fatal("invalid condition executed")
		return nil, nil
	}); err == nil || s != nil {
		t.Fatal("invalid condition accepted")
	}
	if s, err := Run(context.Background(), configurations(2), 1, 1, t.TempDir(), nil); err == nil || s != nil {
		t.Fatal("nil callback accepted")
	}
}

func TestOrderSeedChangesOnlyScheduleAndFullBlocksBalanceForAllAllowedWidths(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var previous []int
	for _, seed := range []int64{7, 777} {
		s, err := Run(ctx, configurations(8), 1, seed, t.TempDir(), func(context.Context, config.Config, string) (*report.Report, error) {
			t.Fatal("canceled comparison started callback")
			return nil, nil
		})
		if !errors.Is(err, context.Canceled) {
			t.Fatal(err)
		}
		if previous != nil && reflect.DeepEqual(previous, s.BaseOrder) {
			t.Fatal("distinct order seeds did not change this schedule")
		}
		previous = s.BaseOrder
		for i, c := range s.Conditions {
			if c.Config.Workload.Seed != int64((i+1)*101) {
				t.Fatal("schedule seed overwrote workload seed")
			}
		}
	}
	for n := 2; n <= 8; n++ {
		s, err := Run(ctx, configurations(n), n, 17, t.TempDir(), func(context.Context, config.Config, string) (*report.Report, error) {
			t.Fatal("canceled comparison started callback")
			return nil, nil
		})
		if !errors.Is(err, context.Canceled) {
			t.Fatal(err)
		}
		balance := make([][]int, n)
		for i := range balance {
			balance[i] = make([]int, n)
		}
		for _, entry := range s.Schedule {
			balance[entry.ConditionIndex-1][entry.RoundPosition-1]++
		}
		for _, positions := range balance {
			for _, count := range positions {
				if count != 1 {
					t.Fatalf("width=%d is not balanced: %v", n, balance)
				}
			}
		}
	}
}

func TestSummaryWriteFailureStopsBeforeAnotherTrial(t *testing.T) {
	root := t.TempDir()
	calls := 0
	s, err := Run(context.Background(), configurations(2), 2, 59, root, func(_ context.Context, c config.Config, dir string) (*report.Report, error) {
		calls++
		r := fixture(c, calls)
		save(t, dir, r)
		// Only the test's generated comparison tree is changed. Make an atomic
		// rename fail after the callback has returned complete cleanup evidence.
		entries, err := os.ReadDir(root)
		if err != nil || len(entries) != 1 {
			t.Fatalf("comparison tree unavailable: %v", err)
		}
		path := filepath.Join(root, entries[0].Name(), "comparison.json")
		if err := os.Remove(path); err != nil {
			t.Fatal(err)
		}
		if err := os.Mkdir(path, 0755); err != nil {
			t.Fatal(err)
		}
		return r, nil
	})
	if err == nil || calls != 1 || s.StopReason != "summary_write_failure" || s.Outcome != "failed" || s.UnstartedTrials != 3 {
		t.Fatalf("write failure proceeded or claimed success: %+v %v", s, err)
	}
}
