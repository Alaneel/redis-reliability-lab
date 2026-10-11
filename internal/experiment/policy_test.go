package experiment

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/report"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

type policyProbeWriter struct {
	client *probeClient
	result workload.WriteResult
}

func (w policyProbeWriter) Write(ctx context.Context, key, value string) workload.WriteResult {
	if w.result.SetAcknowledged {
		w.client.Set(ctx, key, value, 0)
	}
	return w.result
}

func TestPolicyProbeDistinguishesShortReplyFromUnknownAcknowledgement(t *testing.T) {
	for _, tc := range []struct {
		name     string
		result   workload.WriteResult
		observed bool
		success  bool
	}{
		{"full", workload.WriteResult{SetAcknowledged: true, WaitAttempted: true, WaitRequested: 1, WaitAcknowledgements: 1, PolicySatisfied: true}, true, true},
		{"short integer zero", workload.WriteResult{SetAcknowledged: true, WaitAttempted: true, WaitRequested: 1}, true, false},
		{"WAIT error unknown", workload.WriteResult{SetAcknowledged: true, WaitAttempted: true, WaitRequested: 1, WaitError: errors.New("lost reply")}, false, false},
		{"WAIT not issued", workload.WriteResult{SetAcknowledged: true, WaitRequested: 1, WaitError: context.DeadlineExceeded}, false, false},
		{"malformed success with zero acknowledgement", workload.WriteResult{SetAcknowledged: true, WaitAttempted: true, WaitRequested: 1, PolicySatisfied: true}, true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			client := newProbeClient()
			p := ProbeWritePolicy(context.Background(), policyProbeWriter{client, tc.result}, client, []string{"a", "b", "c"}, "test", 1, time.Second)
			if p.Successful != tc.success || p.At.IsZero() || len(p.AcknowledgementObserved) != 3 {
				t.Fatalf("incorrect policy probe: %+v", p)
			}
			for i := range p.Acknowledgements {
				if p.AcknowledgementObserved[i] != tc.observed || p.WaitAttempted[i] != tc.result.WaitAttempted {
					t.Fatal("unknown WAIT outcome fabricated as observed zero")
				}
			}
		})
	}
}

func TestWAITAOFPolicyProbeRequiresIndependentLocalAndReplicaCounts(t *testing.T) {
	full := workload.WriteResult{SetAcknowledged: true, WaitAOFAttempted: true, WaitAOFRequestedLocal: 1, WaitAOFRequestedReplicas: 1, WaitAOFLocalAcknowledgements: 1, WaitAOFReplicaAcknowledgements: 1, PolicySatisfied: true}
	for _, tc := range []struct {
		name              string
		change            func(*workload.WriteResult)
		observed, success bool
	}{
		{"both", func(*workload.WriteResult) {}, true, true},
		{"local only", func(r *workload.WriteResult) { r.WaitAOFReplicaAcknowledgements = 0 }, true, false},
		{"replica only", func(r *workload.WriteResult) { r.WaitAOFLocalAcknowledgements = 0 }, true, false},
		{"observed zero pair", func(r *workload.WriteResult) {
			r.WaitAOFLocalAcknowledgements = 0
			r.WaitAOFReplicaAcknowledgements = 0
		}, true, false},
		{"command error unknown pair", func(r *workload.WriteResult) { r.WaitAOFError = errors.New("lost reply") }, false, false},
		{"not attempted", func(r *workload.WriteResult) { r.WaitAOFAttempted = false; r.WaitAOFError = context.DeadlineExceeded }, false, false},
		{"invalid local", func(r *workload.WriteResult) { r.WaitAOFLocalAcknowledgements = 2 }, false, false},
		{"negative replica", func(r *workload.WriteResult) { r.WaitAOFReplicaAcknowledgements = -1 }, false, false},
		{"different requested replicas", func(r *workload.WriteResult) { r.WaitAOFRequestedReplicas = 2 }, true, false},
		{"SET error contradicts full pair", func(r *workload.WriteResult) { r.SetError = errors.New("lost SET reply") }, true, false},
		{"writer policy false", func(r *workload.WriteResult) { r.PolicySatisfied = false }, true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			result := full
			tc.change(&result)
			client := newProbeClient()
			p := ProbeAcknowledgementPolicy(context.Background(), policyProbeWriter{client, result}, client, []string{"a", "b", "c"}, "aof-test", "waitaof", 1, time.Second)
			if p.Policy != "waitaof" || p.Successful != tc.success || p.WaitAOFRequestedLocal != 1 || p.WaitAOFRequestedReplicas != 1 || p.At.IsZero() || len(p.WaitAOFAcknowledgementObserved) != 3 || len(p.Acknowledgements) != 0 || len(p.WaitAttempted) != 0 || p.RequestedReplicas != 0 {
				t.Fatalf("separate AOF probe populations: %+v", p)
			}
			for i, observed := range p.WaitAOFAcknowledgementObserved {
				if observed != tc.observed || p.WaitAOFAttempted[i] != result.WaitAOFAttempted || p.WaitAOFLocalAcknowledgements[i] != result.WaitAOFLocalAcknowledgements || p.WaitAOFReplicaAcknowledgements[i] != result.WaitAOFReplicaAcknowledgements {
					t.Fatal("known short pair and unknown command outcome conflated")
				}
			}
			if !tc.success && len(p.Errors) != 3 {
				t.Fatal("failed shards omitted from diagnostic errors")
			}
		})
	}
}

func TestUnsupportedAcknowledgementProbeDoesNotCallWriter(t *testing.T) {
	client := newProbeClient()
	p := ProbeAcknowledgementPolicy(context.Background(), nil, client, []string{"a", "b", "c"}, "invalid", "unknown", 1, time.Second)
	if p.Successful || p.At.IsZero() || len(p.Errors) == 0 {
		t.Fatal("unsupported policy silently accepted")
	}
}

func TestSETPreparationProbeHasNoReplicationAcknowledgementPopulation(t *testing.T) {
	client := newProbeClient()
	r := workload.WriteResult{SetAcknowledged: true, PolicySatisfied: true}
	p := ProbeAcknowledgementPolicy(context.Background(), policyProbeWriter{client, r}, client, []string{"a", "b", "c"}, "set-preparation", "set", 0, time.Second)
	if !p.Successful || p.Policy != "set" || len(p.SetAcknowledged) != 3 || p.RequestedReplicas != 0 || p.WaitAOFRequestedLocal != 0 || len(p.Acknowledgements) != 0 || len(p.WaitAOFAcknowledgementObserved) != 0 {
		t.Fatalf("SET prewarming fabricated WAIT/WAITAOF data: %+v", p)
	}
}

func TestAOFExperimentRequiresProcessProofBeforeProvisioning(t *testing.T) {
	c := config.Default()
	c.Environment.Persistence, c.Workload.WritePolicy = "aof-always", "waitaof"
	c.Workload.MaxWrites, c.Workload.ValueBytes = 20000, 128
	m := &fakeManager{}
	called := false
	m.afterProvision = func() { called = true }
	r, err := quietRunner(m).Run(context.Background(), c, t.TempDir())
	if err == nil || r != nil || called {
		t.Fatal("AOF runtime observations were optional or rejected after provisioning")
	}
}

func TestWAITAOFCompletionAtDeadlineCannotRaiseRecovery(t *testing.T) {
	c, r, _, _, _ := persistenceFixture(t, "aof-always")
	c.Workload.WritePolicy = "waitaof"
	at := time.Now()
	r.FaultRequestedAt = &at
	deadline := at.Add(time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	snapshot := *r.Initial
	snapshot.At, snapshot.Completed = at, at
	probe := report.Sample{Snapshot: snapshot, Probe: report.Probe{At: at, Successful: true}, WritePolicyProbe: &report.PolicyProbe{At: deadline, Policy: "waitaof", Successful: true}}
	if ctx.Err() != nil || !recoverySampleExpired(ctx, probe) {
		t.Fatal("explicit policy-reply deadline was not enforced before the context timer fired")
	}
	observePersistenceDiagnostics(ctx, c, r, func(context.Context) report.Sample { return probe })
	if len(r.Samples) != 1 || !r.Samples[0].Excluded || (r.Recovery.WritePolicyRecovered != nil && *r.Recovery.WritePolicyRecovered) || r.Recovery.WritePolicyStable != nil || r.Recovery.CombinedStable != nil || r.PersistenceFault.DiagnosticServiceStable != nil {
		t.Fatal("late AOF command promoted basic or policy recovery")
	}
}
