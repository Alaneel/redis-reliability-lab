package cluster

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestParseNodesRealisticFormat(t *testing.T) {
	primary, replica := strings.Repeat("a", 40), strings.Repeat("b", 40)
	raw := fmt.Sprintf("%s 172.22.0.2:6379@16379,node-1 myself,master - 0 1728763200123 1 connected 0-5460 [42->-%s]\n%s 172.22.0.5:6379@16379 slave %s 0 1728763200110 4 connected [43-<-%s]\n", primary, replica, replica, primary, primary)
	nodes, err := ParseNodes(raw)
	if err != nil {
		t.Fatal(err)
	}
	if len(nodes) != 2 || nodes[0].Addr != "172.22.0.2:6379" || nodes[1].PrimaryID != primary {
		t.Fatalf("unexpected parsed nodes: %+v", nodes)
	}
	if !nodes[0].Primary() || !nodes[1].Replica() || !nodes[1].Healthy() {
		t.Fatal("role or link flags lost")
	}
	if len(nodes[0].SlotTransitions) != 1 || len(nodes[1].Slots) != 0 {
		t.Fatal("transition markers must not create slot assignments")
	}
	c := AnalyzeCoverage(nodes)
	if c.Assigned != 5461 || len(c.Missing) != 1 || c.Missing[0] != (SlotRange{5461, 16383}) {
		t.Fatalf("wrong coverage: %+v", c)
	}
}

func TestParseNodesRejectsMalformedInput(t *testing.T) {
	for _, suffix := range []string{"16384", "-1", "12-10", "0-16384", "1-x", "[42->-]", "[42-43->-abc]", "[42->-abc", "wat"} {
		t.Run(suffix, func(t *testing.T) {
			_, err := ParseNodes("a 172.22.0.2:6379@16379 myself,master - 0 1 1 connected " + suffix)
			if err == nil {
				t.Fatal("malformed slots accepted")
			}
		})
	}
	for _, raw := range []string{"", "a addr master - 0", "a addr master - -1 0 0 connected", "a addr master - 0 0 invalid connected", "a addr master,slave - 0 0 0 connected", "a addr master - 0 0 0 unknown", "a addr master - 0 0 0 connected\na addr slave b 0 0 0 connected"} {
		if _, err := ParseNodes(raw); err == nil {
			t.Errorf("malformed nodes accepted: %q", raw)
		}
	}
}

func TestCoverageRejectsOverlapAndFailedOwners(t *testing.T) {
	nodes := []NodeState{
		{ID: "a", Flags: []string{"master"}, LinkState: "connected", Slots: []SlotRange{{0, 100}, {50, 60}}},
		{ID: "b", Flags: []string{"master", "fail"}, LinkState: "disconnected", Slots: []SlotRange{{100, 16383}}},
	}
	c := AnalyzeCoverage(nodes)
	if c.Full() || c.Assigned != SlotCount || len(c.Missing) != 0 {
		t.Fatalf("wrong assigned slot counting: %+v", c)
	}
	if len(c.Duplicate) != 2 || c.Duplicate[0] != (SlotRange{50, 60}) || c.Duplicate[1] != (SlotRange{100, 100}) {
		t.Fatalf("overlapping claims lost: %+v", c.Duplicate)
	}
	if len(c.UnhealthyOwners) != 1 || c.UnhealthyOwners[0] != (SlotRange{100, 16383}) {
		t.Fatal("failed owner treated as healthy coverage")
	}
}

func TestReadyAndDeterministicSelection(t *testing.T) {
	s := healthySnapshot(t)
	if ok, reasons := Ready(s); !ok {
		t.Fatalf("healthy cluster rejected: %v", reasons)
	}
	selected, err := SelectPrimary(s)
	if err != nil {
		t.Fatal(err)
	}
	if selected.TargetID != fixtureID(0) || selected.ReplicaID != fixtureID(3) || selected.Target.ID != "node-1" || selected.Replica.ID != "node-4" {
		t.Fatalf("wrong selection: %+v", selected)
	}
	if len(selected.Slots) != 5461 || selected.Slots[0] != 0 || selected.Slots[5460] != 5460 {
		t.Fatal("wrong target slots")
	}
	for i, j := 0, len(s.Views)-1; i < j; i, j = i+1, j-1 {
		s.Views[i], s.Views[j] = s.Views[j], s.Views[i]
	}
	reordered, err := SelectPrimary(s)
	if err != nil || selected.TargetID != reordered.TargetID {
		t.Fatal("selection depends on observer order")
	}
}

func TestReadyRequiresAllObserversAndSynchronizedReplicas(t *testing.T) {
	cases := map[string]func(*Snapshot){
		"unreachable node":           func(s *Snapshot) { s.Views[5].Reachable = false; s.Views[5].Errors = []string{"connection refused"} },
		"replica link down":          func(s *Snapshot) { s.Views[3].Replication["master_link_status"] = "down" },
		"replica synchronizing":      func(s *Snapshot) { s.Views[3].Replication["master_sync_in_progress"] = "1" },
		"wrong replication endpoint": func(s *Snapshot) { s.Views[3].Replication["master_host"] = "172.22.0.99" },
		"master has no replica":      func(s *Snapshot) { s.Views[0].Replication["connected_slaves"] = "0" },
		"missing slots":              func(s *Snapshot) { s.Views[0].Nodes[0].Slots = []SlotRange{{0, 5459}} },
		"failed slot owner":          func(s *Snapshot) { s.Views[0].Nodes[0].Flags = append(s.Views[0].Nodes[0].Flags, "fail") },
		"unavailable cluster":        func(s *Snapshot) { s.Views[0].ClusterInfo["cluster_state"] = "fail" },
		"command error":              func(s *Snapshot) { s.Views[0].Errors = []string{"INFO server: read timeout"} },
		"divergent owner map": func(s *Snapshot) {
			s.Views[1].Nodes[0].Slots = []SlotRange{{0, 5459}}
			s.Views[1].Nodes[1].Slots = []SlotRange{{5460, 10922}}
		},
		"duplicate managed identity": func(s *Snapshot) { s.Views[5].Node.ID = s.Views[4].Node.ID },
		"missing observation":        func(s *Snapshot) { s.Views = s.Views[:5] },
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			s := healthySnapshot(t)
			mutate(&s)
			if ok, why := Ready(s); ok || len(why) == 0 {
				t.Fatal("invalid readiness accepted")
			}
		})
	}
}

func TestRecoveryRequiresPromotionAndFullSurvivorConvergence(t *testing.T) {
	baseline := healthySnapshot(t)
	selected, err := SelectPrimary(baseline)
	if err != nil {
		t.Fatal(err)
	}
	recovered := recoveredSnapshot(t)
	if ok, why := RecoveryWithBaseline(recovered, selected, baseline); !ok {
		t.Fatalf("full recovery rejected: %v", why)
	}
	cases := map[string]func(*Snapshot){
		"survivor unreachable":          func(s *Snapshot) { s.Views[5].Reachable = false },
		"old target reachable":          func(s *Snapshot) { s.Views[0].Reachable = true },
		"dead primary still owns slots": func(s *Snapshot) { s.Views[1].Nodes[0].Slots = []SlotRange{{0, 5460}} },
		"old primary not flagged FAIL":  func(s *Snapshot) { s.Views[1].Nodes[0].Flags = []string{"master", "fail?"} },
		"promotion not observed": func(s *Snapshot) {
			s.Views[1].Nodes[3].Flags = []string{"slave"}
			s.Views[1].Nodes[3].PrimaryID = fixtureID(0)
		},
		"partial slot transfer":       func(s *Snapshot) { s.Views[1].Nodes[3].Slots = []SlotRange{{0, 5459}} },
		"cluster state fail":          func(s *Snapshot) { s.Views[1].ClusterInfo["cluster_state"] = "fail" },
		"promoted replica INFO stale": func(s *Snapshot) { s.Views[3].Replication["role"] = "slave" },
		"unaffected owner change": func(s *Snapshot) {
			for i := 1; i < 6; i++ {
				s.Views[i].Nodes[1].Slots = []SlotRange{{5462, 10922}}
				s.Views[i].Nodes[3].Slots = []SlotRange{{0, 5461}}
			}
		},
		"survivor identity changed": func(s *Snapshot) { s.Views[5].RedisID = fixtureID(2) },
		"unknown managed identity":  func(s *Snapshot) { s.Views[5].Node.ID = "unmanaged" },
		"target substitution":       func(s *Snapshot) { s.Views[0].Node.ID = "unmanaged" },
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			s := recoveredSnapshot(t)
			mutate(&s)
			if ok, why := RecoveryWithBaseline(s, selected, baseline); ok || len(why) == 0 {
				t.Fatal("partial recovery accepted")
			}
		})
	}
}

func TestObserveCanceledContextPreservesEveryFailure(t *testing.T) {
	var nodes []Node
	for i := 0; i < 6; i++ {
		nodes = append(nodes, Node{ID: fmt.Sprintf("node-%d", i), Addr: "127.0.0.1:1"})
	}
	o := New(nodes, 100*time.Millisecond)
	defer func() {
		if err := o.Close(); err != nil {
			t.Error(err)
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	start := time.Now()
	s := o.Observe(ctx)
	if time.Since(start) > time.Second {
		t.Fatal("canceled observation did not terminate promptly")
	}
	if len(s.Views) != 6 || s.Completed.IsZero() {
		t.Fatal("failed observations not retained")
	}
	for _, v := range s.Views {
		if v.Reachable || len(v.Errors) == 0 || v.Node.ID == "" {
			t.Fatalf("failure omitted: %+v", v)
		}
	}
	if ok, _ := Ready(s); ok {
		t.Fatal("failed observation accepted as ready")
	}
}

func TestParseInfoPreservesOffsetsAndColonValues(t *testing.T) {
	info := ParseInfo("# Replication\r\nrole:slave\r\nmaster_host:::1\r\nslave_repl_offset:1234\r\nmalformed\r\n")
	if info["role"] != "slave" || info["master_host"] != "::1" || info["slave_repl_offset"] != "1234" {
		t.Fatalf("INFO parse lost fields: %v", info)
	}
}

func fixtureID(i int) string { return strings.Repeat(string(rune('a'+i)), 40) }

func healthySnapshot(t *testing.T) Snapshot {
	t.Helper()
	s := Snapshot{At: time.Now().UTC()}
	for self := 0; self < 6; self++ {
		var raw strings.Builder
		for i := 0; i < 6; i++ {
			flags, parent, slots := "master", "-", ""
			if i >= 3 {
				flags, parent = "slave", fixtureID(i-3)
			} else {
				slots = []string{"0-5460", "5461-10922", "10923-16383"}[i]
			}
			if i == self {
				flags = "myself," + flags
			}
			fmt.Fprintf(&raw, "%s 172.22.0.%d:6379@16379 %s %s 0 1728763200123 %d connected %s\n", fixtureID(i), i+2, flags, parent, i+1, slots)
		}
		nodes, err := ParseNodes(raw.String())
		if err != nil {
			t.Fatal(err)
		}
		replication := map[string]string{"role": "master", "connected_slaves": "1"}
		if self >= 3 {
			replication = map[string]string{"role": "slave", "master_link_status": "up", "master_sync_in_progress": "0", "master_host": fmt.Sprintf("172.22.0.%d", self-1), "master_port": "6379"}
		}
		s.Views = append(s.Views, View{Node: Node{ID: fmt.Sprintf("node-%d", self+1), Addr: fmt.Sprintf("127.0.0.1:%d", 7000+self), InternalAddr: fmt.Sprintf("172.22.0.%d:6379", self+2)}, At: s.At, RedisID: fixtureID(self), Reachable: true, ClusterInfo: map[string]string{"cluster_state": "ok", "cluster_slots_assigned": "16384", "cluster_slots_ok": "16384", "cluster_slots_fail": "0", "cluster_slots_pfail": "0"}, Replication: replication, RedisVersion: "8.2.1", Nodes: nodes, Coverage: AnalyzeCoverage(nodes), RawNodes: raw.String()})
	}
	s.Completed = s.At
	return s
}

func recoveredSnapshot(t *testing.T) Snapshot {
	t.Helper()
	s := healthySnapshot(t)
	s.Views[0].Reachable = false
	s.Views[0].Errors = []string{"connection refused"}
	s.Views[0].Nodes = nil
	s.Views[0].RedisID = ""
	for i := 1; i < 6; i++ {
		s.Views[i].Nodes[0].Slots = nil
		s.Views[i].Nodes[0].Flags = []string{"master", "fail"}
		s.Views[i].Nodes[0].LinkState = "disconnected"
		s.Views[i].Nodes[3].Flags = []string{"master"}
		if i == 3 {
			s.Views[i].Nodes[3].Flags = append(s.Views[i].Nodes[3].Flags, "myself")
		}
		s.Views[i].Nodes[3].PrimaryID = ""
		s.Views[i].Nodes[3].Slots = []SlotRange{{0, 5460}}
		s.Views[i].Coverage = AnalyzeCoverage(s.Views[i].Nodes)
	}
	s.Views[3].Replication = map[string]string{"role": "master", "connected_slaves": "0"}
	return s
}
