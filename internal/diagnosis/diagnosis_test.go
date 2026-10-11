package diagnosis

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func normalizeTest(t *testing.T, report string) Incident {
	t.Helper()
	in, err := Normalize(context.Background(), strings.NewReader(report))
	if err != nil {
		t.Fatal(err)
	}
	return in
}
func testEvidence(i int, node, signal string, v any) Evidence {
	b, _ := json.Marshal(v)
	return Evidence{ID: fmtID(i), At: time.Unix(int64(i), 0).UTC(), Node: node, Source: "test", Signal: signal, Value: b, Status: Valid, Reference: "/observations/" + fmtID(i), Quality: "direct_recorded_observation"}
}
func fmtID(i int) string { return time.Unix(int64(i), 0).UTC().Format("150405") }
func analyzeTest(t *testing.T, es ...Evidence) Assessment {
	t.Helper()
	a, err := Analyze(context.Background(), Incident{SchemaVersion: 1, ID: "opaque", Evidence: es})
	if err != nil {
		t.Fatal(err)
	}
	return a
}
func hasCategory(a Assessment, k string) bool {
	for _, f := range a.Findings {
		if f.Category == k {
			return true
		}
	}
	return false
}

func TestObservableCategoriesAndAmbiguousCauses(t *testing.T) {
	for _, tc := range []struct {
		name string
		es   []Evidence
		want string
	}{
		{"primary", []Evidence{testEvidence(1, "a", "replication.role", "master"), testEvidence(2, "a", "reachable", false)}, "primary_unavailability"},
		{"replica", []Evidence{testEvidence(1, "a", "replication.role", "slave"), testEvidence(2, "a", "reachable", false)}, "replica_unavailability"},
		{"promotion", []Evidence{testEvidence(1, "a", "replication.role", "slave"), testEvidence(2, "a", "replication.role", "master")}, "primary_unavailability"},
		{"redundancy", []Evidence{testEvidence(1, "a", "replication.connected_slaves", 1), testEvidence(2, "a", "replication.connected_slaves", 0)}, "redundancy_degradation"},
		{"loading", []Evidence{testEvidence(1, "a", "persistence.loading", "1")}, "persistence_recovery"},
		{"loss", []Evidence{testEvidence(1, "", "validation.missing", 20)}, "data_validation_failure"},
		{"incomplete", []Evidence{testEvidence(1, "", "validation.unverified", 3)}, "validation_incomplete"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a := analyzeTest(t, tc.es...)
			if !hasCategory(a, tc.want) || a.RootCause != "abstained" {
				t.Fatalf("%+v", a)
			}
			if len(a.Recommendations) != 1 || !a.Recommendations[0].HumanRequired {
				t.Fatal("unsafe plan")
			}
		})
	}
	a := analyzeTest(t, testEvidence(0, "a", "replication.connected_slaves", 1), testEvidence(1, "a", "replication.connected_slaves", 0), testEvidence(2, "", "client_success", false), testEvidence(3, "", "validation.missing", 1))
	if len(a.Findings) != 3 {
		t.Fatal("simultaneous symptoms lost")
	}
}

func TestHealthyMissingAndFalsePositivePrevention(t *testing.T) {
	a := analyzeTest(t, testEvidence(1, "a", "reachable", true), testEvidence(2, "b", "reachable", true), testEvidence(3, "a", "slot_coverage", true))
	if a.Detection != "healthy_observed" || len(a.Findings) != 0 {
		t.Fatal(a.Detection)
	}
	a = analyzeTest(t)
	if a.Detection != "insufficient_evidence" {
		t.Fatal(a.Detection)
	}
	e := testEvidence(1, "a", "reachable", false)
	e.Status = Failed
	a = analyzeTest(t, e)
	if a.Detection == "incident" {
		t.Fatal("collection failure invented a negative observation")
	}
	a = analyzeTest(t, testEvidence(1, "", "validation.complete", false), testEvidence(2, "", "validation.missing", 0))
	if hasCategory(a, "data_validation_failure") {
		t.Fatal("zero defaults classified as loss")
	}
}

func TestNormalizationCompatibilityAndLabelIsolation(t *testing.T) {
	for _, schema := range []string{"1", "2", "3"} {
		raw := `{"schema_version":` + schema + `,"initial":{"at":"2026-10-11T00:00:00Z","views":[{"at":"2026-10-11T00:00:00Z","managed_node":{"id":"n1"},"reachable":true,"replication":{"role":"master","connected_slaves":"1"},"coverage":{"assigned":16384,"owners":[]}}]}}`
		in := normalizeTest(t, raw)
		var root map[string]any
		json.Unmarshal([]byte(raw), &root)
		root["run_id"] = "PRIMARY-KILL"
		root["config"] = map[string]any{"name": "replica-isolation", "experiment": map[string]any{"fault": "replica-isolate"}}
		root["fault_selection"] = map[string]any{"target_redis_id": "label"}
		root["lifecycle_fault"] = map[string]any{"target_kind": "primary", "process_samples": []any{map[string]any{"fault": "kill"}}}
		root["outcome"] = "failed"
		root["effective_config_sha256"] = "secret-label"
		root["events"] = []any{map[string]any{"stage": "injected-kill"}}
		b, _ := json.Marshal(root)
		other := normalizeTest(t, string(b))
		if !reflect.DeepEqual(in.Evidence, other.Evidence) || in.ID != other.ID {
			t.Fatal("labels leaked into normalized evidence")
		}
		a := analyzeTest(t, in.Evidence...)
		bA := analyzeTest(t, other.Evidence...)
		if !reflect.DeepEqual(a.Findings, bA.Findings) {
			t.Fatal("labels changed inference")
		}
		missing := false
		na := false
		for _, e := range in.Evidence {
			missing = missing || e.Signal == "process_live" && e.Status == Missing
			na = na || e.Signal == "replication.master_link_status" && e.Status == NotApplicable
		}
		if !missing || !na {
			t.Fatal("compatibility status lost")
		}
	}
}

func TestHistoricalReportsOffline(t *testing.T) {
	for _, path := range []string{"../../docs/examples/phase3/primary-repair-wait/report.json", "../../docs/examples/phase3/replica-crash/report.json", "../../docs/examples/phase3/persistence-memory/report.json"} {
		f, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		in, err := Normalize(context.Background(), f)
		f.Close()
		if err != nil {
			t.Fatal(err)
		}
		a, err := Analyze(context.Background(), in)
		if err != nil {
			t.Fatal(err)
		}
		if a.Detection != "incident" || len(a.Findings) == 0 {
			t.Fatalf("no diagnosis for %s", path)
		}
		if path == "../../docs/examples/phase3/persistence-memory/report.json" && !hasCategory(a, "data_validation_failure") {
			t.Fatal("missed data loss")
		}
		if !strings.Contains(a.Markdown(), "Evidence provenance") {
			t.Fatal("missing provenance")
		}
	}
}

func TestProvenanceAndSafetyValidation(t *testing.T) {
	a := analyzeTest(t, testEvidence(1, "a", "reachable", false))
	a.Findings[0].Support = append(a.Findings[0].Support, "invented")
	if a.Validate() == nil {
		t.Fatal("accepted fabricated citation")
	}
	a = analyzeTest(t, testEvidence(1, "a", "reachable", false))
	a.Recommendations[0].Execution = "force_failover"
	if a.Validate() == nil {
		t.Fatal("accepted execution")
	}
	a = analyzeTest(t, testEvidence(1, "a", "reachable", false))
	a.Recommendations[0].Preconditions = nil
	if a.Validate() == nil {
		t.Fatal("accepted missing preconditions")
	}
}

func TestInvalidReportsAndCancellation(t *testing.T) {
	for _, raw := range []string{`{`, `null`, `{"schema_version":4}`, `{"schema_version":3} {}`, `{"schema_version":3,"initial":{"views":[{"reachable":false}]}}`} {
		if _, err := Normalize(context.Background(), strings.NewReader(raw)); err == nil {
			t.Fatalf("accepted %s", raw)
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := Normalize(ctx, strings.NewReader(`{"schema_version":3}`)); err == nil {
		t.Fatal("ignored cancellation")
	}
	if _, err := Analyze(ctx, Incident{SchemaVersion: 1, Evidence: []Evidence{testEvidence(1, "a", "reachable", false)}}); err == nil {
		t.Fatal("ignored analysis cancellation")
	}
}

func TestTimestampOrderingAndNoExecutionSideEffects(t *testing.T) {
	in := normalizeTest(t, `{"schema_version":3,"initial":{"views":[{"at":"2026-10-11T00:00:02Z","managed_node":{"id":"a"},"reachable":true},{"at":"2026-10-11T00:00:01Z","managed_node":{"id":"b"},"reachable":true}]}}`)
	for i := 1; i < len(in.Evidence); i++ {
		if in.Evidence[i].At.Before(in.Evidence[i-1].At) {
			t.Fatal("unordered")
		}
	}
	dir := t.TempDir()
	a := analyzeTest(t, testEvidence(1, "a", "reachable", false))
	if err := Write(dir, a); err != nil {
		t.Fatal(err)
	}
	if err := Write(dir, a); err == nil {
		t.Fatal("overwrote report")
	}
	files, err := os.ReadDir(dir)
	if err != nil || len(files) != 2 {
		t.Fatal("unexpected output")
	}
}

func TestStaleAndContradictoryTelemetry(t *testing.T) {
	stale := normalizeTest(t, `{"schema_version":3,"initial":{"at":"2026-10-11T00:01:00Z","completed":"2026-10-11T00:01:01Z","views":[{"at":"2026-10-11T00:00:00Z","managed_node":{"id":"a"},"reachable":false}]}}`)
	for _, e := range stale.Evidence {
		if e.Signal == "reachable" && e.Status != Invalid {
			t.Fatal("stale observation remained valid")
		}
	}
	a := analyzeTest(t, stale.Evidence...)
	if a.Detection == "incident" {
		t.Fatal("stale sample fabricated current incident")
	}
	e1 := testEvidence(1, "peer", "peer.flags", []string{"master"})
	e1.Source = "peer_view:a"
	e1.Reference = "/samples/0/snapshot/views/0/nodes/0/flags"
	e2 := testEvidence(2, "peer", "peer.flags", []string{"master", "fail"})
	e2.Source = "peer_view:b"
	e2.Reference = "/samples/0/snapshot/views/1/nodes/0/flags"
	a = analyzeTest(t, e1, e2)
	if !hasCategory(a, "topology_disagreement") {
		t.Fatal("conflicting peer evidence collapsed")
	}
	e2.Reference = "/samples/1/snapshot/views/1/nodes/0/flags"
	a = analyzeTest(t, e1, e2)
	if hasCategory(a, "topology_disagreement") {
		t.Fatal("different sample times called contradictory")
	}
}

func TestAcknowledgementCountsPreserveSuccessfulSET(t *testing.T) {
	for _, policy := range []string{"wait", "waitaof"} {
		signals := map[string]any{"policy": policy, "successful": false, "set_acknowledged": []bool{true}, "requested_replicas": 1, "acknowledgements": []int64{0}, "acknowledgement_observed": []bool{true}, "waitaof_requested_local": 1, "waitaof_requested_replicas": 1, "waitaof_local_acknowledgements": []int64{1}, "waitaof_replica_acknowledgements": []int64{0}, "waitaof_acknowledgement_observed": []bool{true}}
		es := []Evidence{}
		i := 1
		for _, k := range []string{"policy", "successful", "set_acknowledged", "requested_replicas", "acknowledgements", "acknowledgement_observed", "waitaof_requested_local", "waitaof_requested_replicas", "waitaof_local_acknowledgements", "waitaof_replica_acknowledgements", "waitaof_acknowledgement_observed"} {
			e := testEvidence(i, "", "ack."+k, signals[k])
			e.Reference = "/samples/0/write_policy_probe/" + k
			es = append(es, e)
			i++
		}
		a := analyzeTest(t, es...)
		if !hasCategory(a, "acknowledgement_degradation") {
			t.Fatal("missed short " + policy)
		}
		if hasCategory(a, "client_disruption") {
			t.Fatal("short ack treated as failed SET")
		}
		for i := range es {
			if es[i].Signal == "ack.set_acknowledged" {
				es[i].Value = []byte(`[false]`)
			}
		}
		a = analyzeTest(t, es...)
		if hasCategory(a, "acknowledgement_degradation") {
			t.Fatal("unsuccessful SET cited as ack degradation")
		}
	}
}

func TestPeerSelfMarkerIsNotTopologyConflict(t *testing.T) {
	e1 := testEvidence(1, "peer", "peer.flags", []string{"master", "myself"})
	e1.Source = "peer_view:a"
	e1.Reference = "/initial/views/0/nodes/0/flags"
	e2 := testEvidence(2, "peer", "peer.flags", []string{"master"})
	e2.Source = "peer_view:b"
	e2.Reference = "/initial/views/1/nodes/0/flags"
	a := analyzeTest(t, e1, e2)
	if hasCategory(a, "topology_disagreement") {
		t.Fatal("node-local self marker treated as membership disagreement")
	}
}

func TestLatencyChangeRequiresMeasuredBaseline(t *testing.T) {
	es := []Evidence{}
	for i, p95 := range []float64{1, 100} {
		for j, k := range []string{"attempts", "errors", "p95_ms"} {
			v := any(30)
			if k == "errors" {
				v = 0
			}
			if k == "p95_ms" {
				v = p95
			}
			e := testEvidence(1+i*3+j, "", "client."+k, v)
			e.Reference = fmt.Sprintf("/metrics/windows/%d/%s", i, k)
			es = append(es, e)
		}
	}
	a := analyzeTest(t, es...)
	if !hasCategory(a, "client_latency_change") {
		t.Fatal("missed measured latency change")
	}
	a = analyzeTest(t, es[3:]...)
	if hasCategory(a, "client_latency_change") {
		t.Fatal("fabricated a missing baseline")
	}
}

func TestMissingReachabilityIsNotFailedCollection(t *testing.T) {
	in := normalizeTest(t, `{"schema_version":3,"initial":{"views":[{"at":"2026-10-11T00:00:00Z","managed_node":{"id":"a"}}]}}`)
	for _, e := range in.Evidence {
		if e.Signal == "replication.role" && e.Status != Missing {
			t.Fatalf("missing reachability became collection failure: %+v", e)
		}
	}
}

func TestUnknownSignalsAndCancelledEmptyAnalysis(t *testing.T) {
	a := analyzeTest(t, testEvidence(1, "a", "unknown.future_signal", true))
	if a.Detection != "insufficient_evidence" || len(a.Findings) != 0 {
		t.Fatal("invented diagnosis for unknown signal")
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := Analyze(ctx, Incident{SchemaVersion: 1}); err == nil {
		t.Fatal("ignored cancellation on empty incident")
	}
}

func TestInitialZeroReplicaCountDoesNotInventDegradation(t *testing.T) {
	a := analyzeTest(t, testEvidence(1, "a", "replication.connected_slaves", 0), testEvidence(2, "a", "replication.connected_slaves", 1))
	if hasCategory(a, "redundancy_degradation") {
		t.Fatal("startup unknown replica baseline became degradation")
	}
	a = analyzeTest(t, testEvidence(1, "a", "replication.connected_slaves", 1), testEvidence(2, "a", "replication.connected_slaves", 0))
	if !hasCategory(a, "redundancy_degradation") {
		t.Fatal("missed decline from measured baseline")
	}
}

func TestRealTimeoutCancellationAndAppliedRepairFailureKeepUncertainty(t *testing.T) {
	for _, id := range []string{"f5a5a102abc9f92056580064", "488a31734f8e918c27af7ca0", "ab7af306ada2be3182c8226c"} {
		f, err := os.Open("../../docs/examples/phase3/repair-failure-paths/" + id + "/report.json")
		if err != nil {
			t.Fatal(err)
		}
		in, err := Normalize(context.Background(), f)
		f.Close()
		if err != nil {
			t.Fatal(err)
		}
		a, err := Analyze(context.Background(), in)
		if err != nil {
			t.Fatal(err)
		}
		if a.RootCause != "abstained" {
			t.Fatal("invented physical cause")
		}
		for _, s := range a.Recovery {
			if strings.Contains(s, "population preserved") {
				t.Fatal("early exit falsely claimed value preservation")
			}
		}
		if len(a.Uncertainty) == 0 {
			t.Fatal("lost early-exit uncertainty")
		}
	}
}

func TestClientLagRequiresContemporaneousCoverageAgreement(t *testing.T) {
	client := testEvidence(3, "", "client_success", false)
	client.Reference = "/samples/0/probe/successful"
	c1 := testEvidence(1, "a", "slot_coverage", true)
	c1.Reference = "/samples/0/snapshot/views/0/coverage"
	c2 := testEvidence(2, "b", "slot_coverage", true)
	c2.Reference = "/samples/0/snapshot/views/1/coverage"
	a := analyzeTest(t, c1, c2, client)
	if !hasCategory(a, "client_recovery_lag") {
		t.Fatal("missed contemporaneous topology/client divergence")
	}
	c2.Reference = "/samples/1/snapshot/views/1/coverage"
	a = analyzeTest(t, c1, c2, client)
	if hasCategory(a, "client_recovery_lag") {
		t.Fatal("correlated different samples")
	}
	c2.Reference = "/samples/0/snapshot/views/1/coverage"
	c2.Value = []byte(`false`)
	a = analyzeTest(t, c1, c2, client)
	if hasCategory(a, "client_recovery_lag") {
		t.Fatal("conflicting coverage represented as consensus")
	}
}

func TestDeterministicReplay(t *testing.T) {
	f, err := os.Open("../../docs/examples/phase3/primary-repair-wait/report.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	in, err := Normalize(context.Background(), f)
	if err != nil {
		t.Fatal(err)
	}
	first, err := Analyze(context.Background(), in)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		again, err := Analyze(context.Background(), in)
		if err != nil || !reflect.DeepEqual(first, again) {
			t.Fatalf("nondeterministic replay: %v", err)
		}
	}
}

func TestProjectionBoundsRejectExpansionAbuse(t *testing.T) {
	cases := []map[string]any{
		{"schema_version": 3, "initial": map[string]any{"views": []any{map[string]any{"managed_node": map[string]any{"id": strings.Repeat("a", 257)}, "reachable": true}}}},
		{"schema_version": 3, "initial": map[string]any{"views": []any{map[string]any{"managed_node": map[string]any{"id": "a"}, "reachable": true, "nodes": make([]any, 65)}}}},
		{"schema_version": 3, "initial": map[string]any{"views": []any{map[string]any{"managed_node": map[string]any{"id": "a"}, "reachable": true, "server": map[string]any{}, "replication": map[string]any{"role": strings.Repeat("a", 65537)}}}}},
	}
	for _, doc := range cases {
		raw, _ := json.Marshal(doc)
		if _, err := Normalize(context.Background(), strings.NewReader(string(raw))); err == nil {
			t.Fatal("accepted expansion-prone input")
		}
	}
}
