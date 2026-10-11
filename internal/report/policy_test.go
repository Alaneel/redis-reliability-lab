package report

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/metrics"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

func TestWAITAOFMarkdownSeparatesPairCountsAndFullSETLedger(t *testing.T) {
	r := sampleReport()
	r.SchemaVersion = 3
	r.WorkloadStats.WritePolicy = "waitaof"
	counts := workload.PolicyCounts{SetAcknowledged: 8, WaitAOFAttempted: 7, WaitAOFNotAttempted: 1, WaitAOFNotAttemptedTimeouts: 1,
		WaitAOFSatisfied: 3, WaitAOFShort: 3, WaitAOFErrors: 1, WaitAOFClientTimeouts: 1, WaitAOFLocalSatisfied: 4, WaitAOFReplicaSatisfied: 4,
		WaitAOFReplies: map[string]uint64{"1/1": 3, "0/1": 1, "1/0": 1, "0/0": 1}}
	r.WorkloadStats.PolicyByPhase = map[string]workload.PolicyCounts{"repair": counts}
	r.WorkloadStats.PolicyTotals = counts
	r.WorkloadStats.AcknowledgedWrites = 8
	r.Metrics.Phases["repair"] = metrics.PhaseSummary{DurationSeconds: 2, ByOperation: map[string]metrics.OperationSummary{
		"set_waitaof": {Attempts: 8, Successes: 3, Errors: 5, Latency: metrics.LatencySummary{Count: 8, P50MS: 101, P95MS: 120, P99MS: 150}},
		"set_wait":    {Attempts: 99, Successes: 99, Latency: metrics.LatencySummary{Count: 99, P50MS: 999}},
	}}
	r.WorkloadStats.WriteStages.Phases = map[string]metrics.PhaseSummary{"repair": {ByOperation: map[string]metrics.OperationSummary{
		"set":     {Latency: metrics.LatencySummary{Count: 8, P50MS: 1}},
		"waitaof": {Latency: metrics.LatencySummary{Count: 7, P50MS: 100}},
		"wait":    {Latency: metrics.LatencySummary{Count: 99, P50MS: 999}},
	}}}
	r.Validation.Expected, r.Validation.Checked, r.Validation.Missing = 8, 8, 2
	r.Validation.WaitAOFSatisfiedSubset = workload.ValidationSubset{Expected: 3, Checked: 3, Missing: 1, Complete: true}
	r.Validation.WaitSatisfiedSubset = workload.ValidationSubset{Expected: 99, Checked: 99, Complete: true}
	policy := false
	r.Recovery.WritePolicyRecovered = &policy
	markdown := r.Markdown()
	for _, want := range []string{
		"Write policy: `waitaof`", "same physical connection", "| repair | 8 | 0 | 7 | 3 | 3 | 1 | 1 | 4 | 4 |",
		"| repair | 0/0 | 1 |", "| repair | 0/1 | 1 |", "| repair | 1/0 | 1 |", "| repair | 1/1 | 3 |",
		"| repair | Logical write completion | 8 | 101.000", "| repair | set | 8 | 1.000", "| repair | waitaof | 7 | 100.000", "| repair | 3 | 1.500 |",
		"without a subsequent WAITAOF attempt: 1; client timeouts among those: 1", "Local/replica satisfied counts can overlap and must not be added",
		"WAITAOF-satisfied subset: expected 3; checked 3; missing 1", "main ledger also includes every successful SET followed by incomplete WAITAOF",
		"expected: 8; checked: 8; missing: 2", "policy recovered for the stable interval: false",
	} {
		if !strings.Contains(markdown, want) {
			t.Fatalf("missing WAITAOF evidence %q:\n%s", want, markdown)
		}
	}
	for _, unwanted := range []string{"| WAIT attempted |", "WAIT-satisfied subset:", "| repair | wait |", "999.000"} {
		if strings.Contains(markdown, unwanted) {
			t.Fatalf("WAITAOF evidence mixed with WAIT: %q", unwanted)
		}
	}
	previous := -1
	for _, pair := range []string{"0/0", "0/1", "1/0", "1/1"} {
		position := strings.Index(markdown, "| repair | "+pair+" |")
		if position <= previous {
			t.Fatal("reply pair distribution is not deterministically sorted")
		}
		previous = position
	}
	for i := 0; i < 10; i++ {
		if r.Markdown() != markdown {
			t.Fatal("pair distribution changes between renders")
		}
	}
}

func TestWAITAOFErrorsDoNotInventZeroReplyOrStageMeasurement(t *testing.T) {
	r := sampleReport()
	r.WorkloadStats.WritePolicy = "waitaof"
	r.WorkloadStats.PolicyByPhase = map[string]workload.PolicyCounts{"fault": {SetAcknowledged: 2, WaitAOFAttempted: 1, WaitAOFErrors: 1, WaitAOFNotAttempted: 1}}
	markdown := r.Markdown()
	if strings.Contains(markdown, "| fault | 0/0 |") || strings.Contains(markdown, "| fault | waitaof |") {
		t.Fatal("missing reply or stage metrics were rendered as measured zeros")
	}
}

func TestPolicyProbeJSONPreservesKnownZeroAndUnknownWAITAOFPair(t *testing.T) {
	probe := PolicyProbe{Policy: "waitaof", WaitAOFRequestedLocal: 1, WaitAOFRequestedReplicas: 1,
		SetAcknowledged: []bool{true, true, true}, WaitAOFAttempted: []bool{true, true, true},
		WaitAOFLocalAcknowledgements: []int64{0, 1, 0}, WaitAOFReplicaAcknowledgements: []int64{0, 1, 0},
		WaitAOFAcknowledgementObserved: []bool{true, true, false}, Errors: []string{"shard 2: deadline exceeded"}}
	data, err := json.Marshal(probe)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"policy":"waitaof"`, `"waitaof_local_acknowledgements":[0,1,0]`, `"waitaof_replica_acknowledgements":[0,1,0]`, `"waitaof_acknowledgement_observed":[true,true,false]`} {
		if !strings.Contains(string(data), want) {
			t.Fatalf("observed-zero/unknown distinction lost: %s", data)
		}
	}
	var decoded PolicyProbe
	if err := json.Unmarshal(data, &decoded); err != nil || decoded.WaitAOFAcknowledgementObserved[2] || !decoded.WaitAOFAcknowledgementObserved[0] || decoded.WaitAOFLocalAcknowledgements[0] != 0 {
		t.Fatalf("WAITAOF probe round trip lost observations: %v", err)
	}
	data, err = json.Marshal(PolicyProbe{RequestedReplicas: 1, WaitAttempted: []bool{true}, Acknowledgements: []int64{1}, AcknowledgementObserved: []bool{true}})
	if err != nil || strings.Contains(string(data), "waitaof_") || strings.Contains(string(data), `"policy"`) {
		t.Fatalf("legacy WAIT probe invented WAITAOF evidence: %v %s", err, data)
	}
}

func TestWAITAOFPreparationRendersShortAndUnknownEvidence(t *testing.T) {
	r := sampleReport()
	r.WorkloadStats.WritePolicy = "waitaof"
	r.BaselineWritePolicyProbe = &PolicyProbe{At: time.Now(), Policy: "waitaof", WaitAOFRequestedLocal: 1, WaitAOFRequestedReplicas: 1,
		SetAcknowledged: []bool{true, true, true}, WaitAOFAttempted: []bool{true, true, true},
		WaitAOFLocalAcknowledgements: []int64{0, 1, 0}, WaitAOFReplicaAcknowledgements: []int64{1, 1, 0},
		WaitAOFAcknowledgementObserved: []bool{true, true, false}, Errors: []string{"shard 2: deadline exceeded"}}
	markdown := r.Markdown()
	for _, want := range []string{"Administrative persistence runtime preparation: unobserved", "all shards successful: false", "requested local/replica AOF counts: 1/1", "| 0 | true | true | 0/1 |", "| 1 | true | true | 1/1 |", "| 2 | true | true | unobserved |", "shard 2: deadline exceeded"} {
		if !strings.Contains(markdown, want) {
			t.Fatalf("pre-fault evidence missing %q", want)
		}
	}
	r.BaselineWritePolicyProbe = nil
	if !strings.Contains(r.Markdown(), "All-shard before-baseline acknowledgement-policy probe: unobserved") {
		t.Fatal("missing pre-fault probe became a healthy observation")
	}
	data, err := json.Marshal(sampleReport())
	if err != nil || strings.Contains(string(data), "persistence_prepared") || strings.Contains(string(data), "baseline_write_policy_probe") {
		t.Fatalf("legacy report invented policy preparation: %v %s", err, data)
	}
}

func TestBeforeBaselineProbeRetainsSelectedPolicy(t *testing.T) {
	for _, policy := range []string{"set", "wait"} {
		t.Run(policy, func(t *testing.T) {
			r := sampleReport()
			r.WorkloadStats.WritePolicy = policy
			r.BaselineWritePolicyProbe = &PolicyProbe{At: time.Now(), Policy: policy, Successful: true, SetAcknowledged: []bool{true, true, true}}
			if policy == "wait" {
				r.BaselineWritePolicyProbe.RequestedReplicas = 1
				r.BaselineWritePolicyProbe.WaitAttempted = []bool{true, true, true}
				r.BaselineWritePolicyProbe.Acknowledgements = []int64{1, 1, 1}
				r.BaselineWritePolicyProbe.AcknowledgementObserved = []bool{true, true, true}
			}
			markdown := r.Markdown()
			if !strings.Contains(markdown, "Before-baseline `"+policy+"` policy probe") || strings.Contains(markdown, "WAITAOF attempted") || strings.Contains(markdown, "local/replica AOF counts") {
				t.Fatalf("compatibility probe misidentified policy:\n%s", markdown)
			}
			if policy == "wait" && !strings.Contains(markdown, "| 0 | true | true | 1 |") {
				t.Fatal("WAIT compatibility reply omitted")
			}
		})
	}
}
