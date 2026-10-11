//go:build integration

package experiment

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/cluster"
	"github.com/Alaneel/redis-reliability-lab/internal/config"
	"github.com/Alaneel/redis-reliability-lab/internal/environment"
	"github.com/Alaneel/redis-reliability-lab/internal/report"
)

func assertWAITAOFAccounting(t *testing.T, r *report.Report) {
	t.Helper()
	s, p := r.WorkloadStats, r.WorkloadStats.PolicyTotals
	if s.WritePolicy != "waitaof" || p.SetAcknowledged != s.AcknowledgedWrites || p.SetAcknowledged+p.SetErrors != s.WritesAttempted || p.WaitAOFAttempted+p.WaitAOFNotAttempted != p.SetAcknowledged || p.WaitAOFSatisfied+p.WaitAOFShort+p.WaitAOFErrors != p.WaitAOFAttempted || p.WaitAOFClientTimeouts > p.WaitAOFErrors || p.WaitAOFNotAttemptedTimeouts > p.WaitAOFNotAttempted || p.WaitAttempted != 0 || p.WaitSatisfied != 0 || len(p.WaitAcknowledgements) != 0 {
		t.Fatalf("separate SET/WAITAOF/WAIT accounting: %+v", s)
	}
	var replies uint64
	for pair, count := range p.WaitAOFReplies {
		if !strings.Contains(pair, "/") {
			t.Fatalf("reply pair lost: %q", pair)
		}
		replies += count
	}
	if replies != p.WaitAOFSatisfied+p.WaitAOFShort || p.WaitAOFLocalSatisfied < p.WaitAOFSatisfied || p.WaitAOFReplicaSatisfied < p.WaitAOFSatisfied || p.PolicySatisfied != p.WaitAOFSatisfied {
		t.Fatal("actual valid-pair histogram/components do not reconcile with policy counts")
	}
	if s.WriteStages.Totals.ByOperation["set"].Attempts != s.WritesAttempted || s.WriteStages.Totals.ByOperation["waitaof"].Attempts != p.WaitAOFAttempted || s.WriteStages.Totals.ByOperation["wait"].Attempts != 0 || r.Metrics.Totals.ByOperation["set_waitaof"].Attempts != s.WritesAttempted || s.OfferedRequests != s.StartedRequests+s.DroppedRequests || s.StartedRequests != s.CompletedRequests || s.CompletedRequests != r.Metrics.Totals.Attempts {
		t.Fatal("AOF stage commands leaked into logical operation or pacing counts")
	}
	v := r.Validation
	if !v.Complete || v.Expected != int(s.AcknowledgedWrites) || v.Checked != v.Expected || !v.WaitAOFSatisfiedSubset.Complete || v.WaitAOFSatisfiedSubset.Expected != int(p.WaitAOFSatisfied) || v.WaitAOFSatisfiedSubset.Checked != v.WaitAOFSatisfiedSubset.Expected || v.WaitSatisfiedSubset.Expected != 0 {
		t.Fatalf("SET ledger or both-required AOF subset disappeared: %+v", v)
	}
	b := r.BaselineWritePolicyProbe
	if b == nil || b.Policy != "waitaof" || !b.Successful || b.WaitAOFRequestedLocal != 1 || b.WaitAOFRequestedReplicas != 1 || len(b.WaitAOFAcknowledgementObserved) != 3 || len(b.WaitAttempted) != 0 || r.FaultRequestedAt == nil || !b.At.Before(*r.FaultRequestedAt) {
		t.Fatal("compatible actual pre-baseline WAITAOF command evidence absent")
	}
	for i, observed := range b.WaitAOFAcknowledgementObserved {
		if !observed || !b.SetAcknowledged[i] || !b.WaitAOFAttempted[i] || b.WaitAOFLocalAcknowledgements[i] != 1 || b.WaitAOFReplicaAcknowledgements[i] < 1 {
			t.Fatal("pre-baseline command did not observe both local and replica AOF acknowledgement")
		}
	}
	for _, event := range r.Events {
		if event.Stage == string(Baseline) && !b.At.Before(event.At) {
			t.Fatal("preflight WAITAOF fsynced workload writes after baseline began")
		}
	}
	t.Logf("set_ack=%d both_aof=%d short=%d errors=%d ambiguous=%d all_checked=%d subset_checked=%d subset_missing=%d", p.SetAcknowledged, p.WaitAOFSatisfied, p.WaitAOFShort, p.WaitAOFErrors, s.AmbiguousWrites, v.Checked, v.WaitAOFSatisfiedSubset.Checked, v.WaitAOFSatisfiedSubset.Missing)
}

func TestDockerWAITAOFPrimaryRepairRestoresPolicy(t *testing.T) {
	c := lifecycleIntegrationConfig("primary-kill")
	c.Name, c.Environment.Persistence, c.Workload.WritePolicy = "waitaof-primary-repair", "aof-always", "waitaof"
	c.Experiment.Baseline, c.Experiment.PostRecovery = config.Duration(4*time.Second), config.Duration(2*time.Second)
	c.Timeouts.Repair, c.Workload.MaxWrites, c.Workload.ValueBytes = config.Duration(time.Minute), 20000, 128
	r, err := runIntegration(t, c, environment.New())
	assertRestoredLifecycle(t, r, false)
	assertOwnedRestart(t, r)
	assertWAITAOFAccounting(t, r)
	if len(r.PersistencePrepared) != 6 || r.WorkloadStats.PolicyByPhase["baseline"].WaitAOFSatisfied == 0 || r.WorkloadStats.PolicyByPhase["fault"].WaitAOFShort == 0 || r.Recovery.WritePolicyRecovered == nil || !*r.Recovery.WritePolicyRecovered || r.Recovery.WritePolicyStable == nil || r.Dimensions.WritePolicyAvailable == nil || !*r.Dimensions.WritePolicyAvailable {
		t.Fatal("former primary repair did not restore independently measured local+replica AOF policy")
	}
	if r.Validation.Missing > 0 || r.WorkloadStats.MissingReads > 0 {
		if err == nil || r.Outcome != "partial" || r.Recovery.DataVerified {
			t.Fatal("measured loss hidden behind successful AOF policy")
		}
	} else if err != nil || r.Outcome != "recovered" {
		t.Fatalf("actual full repair misclassified: outcome=%s err=%v", r.Outcome, err)
	}
}

func TestDockerWAITAOFFrozenPersistenceValidation(t *testing.T) {
	c := persistenceIntegrationConfig("aof-always")
	c.Name, c.Workload.WritePolicy = "waitaof-frozen-persistence", "waitaof"
	r, err := runIntegration(t, c, environment.New())
	assertPersistenceRestartEvidence(t, r)
	assertWAITAOFAccounting(t, r)
	if r.Recovery.WritePolicyRecovered == nil || !*r.Recovery.WritePolicyRecovered || r.Dimensions.WritePolicyAvailable == nil || !*r.Dimensions.WritePolicyAvailable || r.WorkloadStats.PolicyByPhase["baseline"].WaitAOFSatisfied == 0 {
		t.Fatal("post-validation AOF policy diagnostics absent or unavailable")
	}
	if r.Validation.Missing > 0 {
		if err == nil || r.Outcome != "partial" || r.Recovery.DataVerified {
			t.Fatal("frozen acknowledged loss became full recovery")
		}
	} else if err != nil || r.Outcome != "recovered" {
		t.Fatalf("actual frozen AOF validation misclassified: outcome=%s err=%v", r.Outcome, err)
	}
}

func TestDockerWAITAOFDisabledRejectsCommandRetainsSET(t *testing.T) {
	// Config and ManagedWriter correctly reject memory WAITAOF experiments.
	// This command compatibility negative test uses only a fixed owned Redis
	// node and one sticky SDK connection; no runtime configuration is changed.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	m := environment.New()
	id, err := RunID()
	if err != nil {
		t.Fatal(err)
	}
	c := config.Default()
	e, err := m.Provision(ctx, id, environment.Config{Image: c.Environment.RedisImage, NodeTimeout: c.Environment.NodeTimeout.Value(), Persistence: "memory", RetainData: true})
	t.Cleanup(func() {
		if e != nil {
			bounded, done := context.WithTimeout(context.Background(), time.Minute)
			defer done()
			if err := m.Cleanup(bounded, e); err != nil {
				t.Error(err)
			}
		}
	})
	if err != nil {
		t.Fatal(err)
	}
	nodes := make([]cluster.Node, len(e.Nodes))
	for i, n := range e.Nodes {
		nodes[i] = cluster.Node{ID: n.ID, Addr: n.Addr, InternalAddr: n.InternalAddr}
	}
	o := cluster.New(nodes, time.Second)
	defer o.Close()
	ready, err := waitReady(ctx, o, 200*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	client, err := NewClient(e, time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	keys, err := ProbeKeys(ready)
	if err != nil {
		t.Fatal(err)
	}
	for _, key := range keys {
		master, err := client.MasterForKey(ctx, key)
		if err != nil {
			t.Fatal(err)
		}
		conn := master.Conn()
		op, done := context.WithTimeout(ctx, time.Second)
		value := id + ":aof-disabled"
		setErr := conn.Set(op, key, value, 0).Err()
		reply, aofErr := conn.WaitAOF(op, 1, 1, 100*time.Millisecond).Result()
		done()
		conn.Close()
		if setErr != nil || aofErr == nil || !strings.Contains(strings.ToLower(aofErr.Error()), "appendonly") {
			t.Fatalf("actual memory-mode WAITAOF did not reject local durability: SET=%v reply=%v error=%v", setErr, reply, aofErr)
		}
		read, done := context.WithTimeout(ctx, time.Second)
		actual, getErr := client.Get(read, key).Result()
		done()
		if getErr != nil || actual != value {
			t.Fatalf("SET acknowledgement erased by failed AOF policy: value=%q err=%v", actual, getErr)
		}
		t.Logf("run_id=%s shard_key=%s SET acknowledged; WAITAOF error=%v; GET retained value", id, key, aofErr)
	}
	// A cluster routing helper may fall back to a primary. Bind each negative
	// command to the exact managed replica endpoint, then observe its live role
	// on that same physical connection before issuing the keyless command.
	fixed, err := NewWriter(client, e, time.Second, "set", 1, 100*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	defer fixed.Close()
	replicas := 0
	for _, v := range ready.Views {
		if v.Replication["role"] != "slave" {
			continue
		}
		var owned environment.Node
		for _, n := range e.Nodes {
			if n.ID == v.Node.ID && n.Addr == v.Node.Addr && n.InternalAddr == v.Node.InternalAddr && n.RedisClusterID == v.RedisID {
				owned = n
			}
		}
		if owned.ID == "" || fixed.nodes[owned.InternalAddr] == nil {
			t.Fatal("replica observation is not an exact managed endpoint")
		}
		conn := fixed.nodes[owned.InternalAddr].Conn()
		op, done := context.WithTimeout(ctx, time.Second)
		raw, infoErr := conn.Info(op, "replication").Result()
		role := cluster.ParseInfo(raw)["role"]
		if infoErr != nil || role != "slave" {
			done()
			conn.Close()
			t.Fatalf("fixed managed endpoint %s is not an actual replica: role=%q err=%v", owned.ID, role, infoErr)
		}
		reply, commandErr := conn.WaitAOF(op, 1, 1, 100*time.Millisecond).Result()
		done()
		conn.Close()
		if commandErr == nil || !strings.Contains(strings.ToLower(commandErr.Error()), "replica") {
			t.Fatalf("WAITAOF actual replica %s was not rejected: role=%s reply=%v err=%v", owned.ID, role, reply, commandErr)
		}
		replicas++
		t.Logf("run_id=%s managed_node=%s endpoint=%s observed_role=%s WAITAOF replica rejection without SET: %v", id, owned.ID, owned.InternalAddr, role, commandErr)
	}
	if replicas != 3 {
		t.Fatalf("expected three actual managed replica command rejections, got %d", replicas)
	}
}
