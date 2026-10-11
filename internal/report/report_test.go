package report

import (
	"encoding/json"
	"errors"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/config"
	"github.com/Alaneel/redis-reliability-lab/internal/metrics"
	"github.com/Alaneel/redis-reliability-lab/internal/workload"
)

func sampleReport() Report {
	start := time.Date(2026, 10, 11, 1, 2, 3, 0, time.UTC)
	c := metrics.NewCollector(start)
	c.RecordAt("baseline", "set", 2*time.Millisecond, nil, start.Add(time.Second))
	c.RecordAt("baseline", "get", 4*time.Millisecond, nil, start.Add(time.Second))
	c.SetPhase("fault", start.Add(2*time.Second))
	c.RecordAt("fault", "set", 30*time.Millisecond, errors.New("CLUSTERDOWN no slot"), start.Add(3*time.Second))
	c.Close(start.Add(4 * time.Second))
	return Report{SchemaVersion: 1, RunID: "report-test", Outcome: "partial", StartedAt: start, FinishedAt: start.Add(5 * time.Second),
		Config: config.Default(), ConfigSHA256: strings.Repeat("a", 64), Host: map[string]string{"os": "test"},
		Events: []Event{{"created", start}, {"complete", start.Add(5 * time.Second)}}, Samples: []Sample{},
		Recovery: Recovery{ClientFirstSuccess: &Milestone{At: start, SinceFaultMS: 0}, ClientRecovered: true},
		Metrics:  c.Snapshot(start.Add(5 * time.Second)), WorkloadStats: workload.Stats{WritesAttempted: 22, AcknowledgedWrites: 20, AmbiguousWrites: 2, Guarantee: workload.Guarantee},
		Validation: workload.Validation{StartedAt: start.Add(4 * time.Second), CompletedAt: start.Add(5 * time.Second), Expected: 20, Checked: 17, Missing: 1, Corrupt: 2, Unreadable: 3, Unverified: 3, Complete: false, Guarantee: workload.Guarantee},
		Cleanup:    Cleanup{Attempted: true, Successful: false, Error: "container still running"}, Errors: []string{"recovery deadline elapsed", "cleanup failed: container still running"},
		Methodology: []string{"Measured from real client attempts.", "Milestones are sampled observations."}}
}

func TestJSONPreservesAbsenceZeroMetricsAndEvidence(t *testing.T) {
	r := sampleReport()
	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	var document map[string]json.RawMessage
	if err := json.Unmarshal(data, &document); err != nil {
		t.Fatal(err)
	}
	var recovery map[string]json.RawMessage
	if err := json.Unmarshal(document["recovery"], &recovery); err != nil {
		t.Fatal(err)
	}
	if _, present := recovery["replica_promoted"]; present {
		t.Fatal("unobserved milestone fabricated")
	}
	if _, present := recovery["client_first_all_shard_success"]; !present {
		t.Fatal("observed zero-delay milestone omitted")
	}
	var decoded struct {
		Recovery      Recovery            `json:"recovery"`
		Metrics       metrics.Snapshot    `json:"metrics"`
		WorkloadStats workload.Stats      `json:"workload_stats"`
		Validation    workload.Validation `json:"validation"`
		Cleanup       Cleanup             `json:"cleanup"`
	}
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.Recovery.ClientFirstSuccess == nil || decoded.Recovery.ClientFirstSuccess.SinceFaultMS != 0 || decoded.Recovery.ReplicaPromoted != nil {
		t.Fatalf("milestones changed: %+v", decoded.Recovery)
	}
	var encodedConfig struct {
		Workload struct {
			OperationTimeout string `json:"operation_timeout"`
		} `json:"workload"`
	}
	if err := json.Unmarshal(document["config"], &encodedConfig); err != nil {
		t.Fatal(err)
	}
	if encodedConfig.Workload.OperationTimeout != r.Config.Workload.OperationTimeout.String() {
		t.Fatal("configuration duration changed")
	}
	baseline := decoded.Metrics.Phases["baseline"]
	if baseline.Attempts != 2 || baseline.SuccessRate != 1 || baseline.Throughput != 1 || baseline.SuccessfulLatency.P99MS != 4 {
		t.Fatalf("metrics changed: %+v", baseline)
	}
	if decoded.Metrics.Phases["fault"].ErrorClasses["CLUSTERDOWN"] != 1 || decoded.WorkloadStats.AmbiguousWrites != 2 {
		t.Fatal("error evidence lost")
	}
	if decoded.Validation.Unverified != 3 || decoded.Validation.Missing != 1 || decoded.Validation.Corrupt != 2 || decoded.Validation.Unreadable != 3 || decoded.Cleanup.Error != "container still running" {
		t.Fatalf("validation/cleanup evidence changed: %+v %+v", decoded.Validation, decoded.Cleanup)
	}
}

func TestMarkdownUsesMeasuredMetricsAndUnobservedMilestones(t *testing.T) {
	markdown := sampleReport().Markdown()
	for _, want := range []string{
		"Outcome: **partial**",
		"| Original replica owns failed primary slots | unobserved |",
		"| First all-shard SET/GET probe succeeds | 0.00 |",
		"| baseline | 2 | 100.00% | 0 | 1.00 |",
		"| fault | 1 | 0.00% | 1 | 0.50 | unobserved |",
		"Acknowledged writes: 20.",
		"Validation attempted: true; expected: 20; checked: 17; missing: 1; corrupt: 2; unreadable: 3; unverified: 3.",
		"| fault | CLUSTERDOWN | 1 |",
		"| Topology and client recovery jointly sustained | false |",
		"Validation complete: false", "Ambiguous SET outcomes: 2.",
		"container still running", "recovery deadline elapsed", "Cleanup verified | false",
		"Redis replication is asynchronous", "successful", "terminal", strings.Repeat("a", 64),
	} {
		if !strings.Contains(markdown, want) {
			t.Errorf("markdown missing %q:\n%s", want, markdown)
		}
	}
	if strings.Contains(markdown, "| post_recovery | 0 |") {
		t.Fatal("absent workload phase fabricated as zero measurements")
	}
}

func TestFaultActionClockSeparatesOrchestrationFromRecovery(t *testing.T) {
	r := sampleReport()
	r.SchemaVersion = 2
	overhead := 125.0
	r.FaultActionDurationMS = &overhead
	r.Recovery.ClientFirstSuccess.SinceFaultMS = 1125
	markdown := r.Markdown()
	if !strings.Contains(markdown, "action completed 125.000 ms") ||
		!strings.Contains(markdown, "| First all-shard SET/GET probe succeeds | 1125.00 | 1000.00 |") {
		t.Fatalf("fault overhead was conflated with recovery:\n%s", markdown)
	}
	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}
	var saved Report
	if err := json.Unmarshal(data, &saved); err != nil {
		t.Fatal(err)
	}
	if saved.FaultActionDurationMS == nil || *saved.FaultActionDurationMS != overhead {
		t.Fatal("measured fault action duration lost")
	}
	r.FaultActionDurationMS = nil
	if !strings.Contains(r.Markdown(), "1125.00 | unavailable |") {
		t.Fatal("legacy report invented an action-relative measurement")
	}
}

func TestMarkdownSeparatesAcknowledgedLedgerFromSkippedValidation(t *testing.T) {
	r := sampleReport()
	r.Outcome = "aborted"
	r.Validation = workload.Validation{}
	markdown := r.Markdown()
	for _, want := range []string{"Acknowledged writes: 20.", "Validation attempted: false; expected: 0; checked: 0;", "Validation was not attempted; 20 acknowledged writes remain unchecked."} {
		if !strings.Contains(markdown, want) {
			t.Errorf("skipped validation missing %q:\n%s", want, markdown)
		}
	}
	if strings.Contains(markdown, "Acknowledged writes: 0.") {
		t.Fatal("acknowledged ledger discarded when validation skipped")
	}
}

func TestMarkdownErrorDistributionIsActualAndSorted(t *testing.T) {
	r := sampleReport()
	r.Metrics.Phases["baseline"] = metrics.PhaseSummary{ErrorClasses: map[string]uint64{"timeout": 2, "ASK": 1, "absent": 0}}
	r.Metrics.Phases["fault"] = metrics.PhaseSummary{ErrorClasses: map[string]uint64{"network": 4, "CLUSTERDOWN": 3}}
	r.Recovery.CombinedRecovered = true
	markdown := r.Markdown()
	rows := []string{"| baseline | ASK | 1 |", "| baseline | timeout | 2 |", "| fault | CLUSTERDOWN | 3 |", "| fault | network | 4 |"}
	previous := -1
	for _, row := range rows {
		at := strings.Index(markdown, row)
		if at <= previous {
			t.Fatalf("error rows absent or unsorted: %q\n%s", row, markdown)
		}
		previous = at
	}
	if strings.Contains(markdown, "| baseline | absent |") {
		t.Fatal("zero-count error fabricated")
	}
	if !strings.Contains(markdown, "| Topology and client recovery jointly sustained | true |") {
		t.Fatal("combined recovery condition omitted")
	}
	for i := 0; i < 10; i++ {
		if got := r.Markdown(); got != markdown {
			t.Fatal("error distribution is nondeterministic")
		}
	}
}

func TestMarkdownDisplaysCleanupErrorIndependently(t *testing.T) {
	r := sampleReport()
	r.Errors = nil
	if markdown := r.Markdown(); !strings.Contains(markdown, r.Cleanup.Error) {
		t.Fatalf("cleanup failure detail absent:\n%s", markdown)
	}
}

func TestMarkdownSeparatesSETAndWAITAndNetworkReconciliation(t *testing.T) {
	r := sampleReport()
	r.WorkloadStats.WritePolicy = "wait"
	r.WorkloadStats.PolicyByPhase = map[string]workload.PolicyCounts{"fault": {SetAcknowledged: 20, WaitAttempted: 20, WaitSatisfied: 15, WaitShort: 4, WaitErrors: 1, WaitClientTimeouts: 1}}
	r.Validation.WaitSatisfiedSubset = workload.ValidationSubset{Expected: 15, Checked: 15, Complete: true}
	policy := false
	r.Recovery.WritePolicyRecovered = &policy
	r.NetworkFault = &NetworkFault{ConnectivityRestored: true, Reconciled: false, ReconciliationTimedOut: true}
	r.Metrics.Phases["reconciliation"] = metrics.PhaseSummary{Attempts: 2, Errors: 1}
	markdown := r.Markdown()
	for _, want := range []string{"| fault | 20 | 0 | 20 | 15 | 4 | 1 | 1 |", "WAIT-satisfied subset: expected 15; checked 15", "policy recovered for the stable interval: false", "| reconciliation | 2 |", "whole-node bidirectional network isolation", "| Reconciliation deadline expired | true |"} {
		if !strings.Contains(markdown, want) {
			t.Fatalf("evidence omitted %q: %s", want, markdown)
		}
	}
	if strings.Contains(markdown, "primary stays stopped") {
		t.Fatal("network isolation described as process termination")
	}
}

func TestWriteReplacesValidReportsAndLeavesNoTemporaryFiles(t *testing.T) {
	dir := t.TempDir()
	r := sampleReport()
	for _, name := range []string{"report.json", "report.md"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("old report"), 0600); err != nil {
			t.Fatal(err)
		}
	}
	if err := Write(dir, r); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "report.json"))
	if err != nil {
		t.Fatal(err)
	}
	var decoded struct {
		RunID   string `json:"run_id"`
		Outcome string `json:"outcome"`
	}
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.RunID != r.RunID || decoded.Outcome != r.Outcome || len(data) == 0 || data[len(data)-1] != '\n' {
		t.Fatal("written JSON does not match report")
	}
	markdown, err := os.ReadFile(filepath.Join(dir, "report.md"))
	if err != nil {
		t.Fatal(err)
	}
	if string(markdown) != r.Markdown() {
		t.Fatal("written markdown differs")
	}
	assertNoTemporaryReports(t, dir)
}

func TestEncodingFailureDoesNotReplaceExistingReports(t *testing.T) {
	dir := t.TempDir()
	r := sampleReport()
	r.Config.Workload.RatePerSecond = math.NaN()
	for _, name := range []string{"report.json", "report.md"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("preserved"), 0600); err != nil {
			t.Fatal(err)
		}
	}
	if err := Write(dir, r); err == nil || !strings.Contains(err.Error(), "encode JSON") {
		t.Fatalf("expected JSON encoding error, got %v", err)
	}
	for _, name := range []string{"report.json", "report.md"} {
		data, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != "preserved" {
			t.Fatal("encoding failure replaced existing report")
		}
	}
	assertNoTemporaryReports(t, dir)
}

func TestAtomicWriteRenameFailurePreservesTargetAndRemovesTemporary(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "report.json")
	if err := os.Mkdir(target, 0700); err != nil {
		t.Fatal(err)
	}
	if err := Write(dir, sampleReport()); err == nil {
		t.Fatal("write over directory unexpectedly succeeded")
	}
	info, err := os.Stat(target)
	if err != nil || !info.IsDir() {
		t.Fatalf("target changed: info=%v err=%v", info, err)
	}
	if _, err := os.Stat(filepath.Join(dir, "report.md")); !os.IsNotExist(err) {
		t.Fatalf("markdown written after first file failed: %v", err)
	}
	assertNoTemporaryReports(t, dir)
}

func TestSecondFileFailureIsReturnedWithoutClaimingPairAtomicity(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "report.md")
	if err := os.Mkdir(target, 0700); err != nil {
		t.Fatal(err)
	}
	if err := Write(dir, sampleReport()); err == nil {
		t.Fatal("second-file failure was hidden")
	}
	jsonData, err := os.ReadFile(filepath.Join(dir, "report.json"))
	if err != nil || !json.Valid(jsonData) {
		t.Fatalf("successful first JSON write lost: %v", err)
	}
	info, err := os.Stat(target)
	if err != nil || !info.IsDir() {
		t.Fatalf("markdown target replaced: %v", err)
	}
	assertNoTemporaryReports(t, dir)
}

func TestWriteMissingDirectoryReturnsError(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")
	if err := Write(missing, sampleReport()); err == nil {
		t.Fatal("write with missing directory succeeded")
	}
	if _, err := os.Stat(missing); !os.IsNotExist(err) {
		t.Fatalf("missing report directory unexpectedly created: %v", err)
	}
}

func assertNoTemporaryReports(t *testing.T, dir string) {
	t.Helper()
	matches, err := filepath.Glob(filepath.Join(dir, ".report-*"))
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) > 0 {
		t.Fatalf("temporary report files leaked: %v", matches)
	}
}
