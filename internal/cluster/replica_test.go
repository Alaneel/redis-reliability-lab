package cluster

import (
	"slices"
	"strings"
	"testing"
)

func replicaFaultSnapshot(t *testing.T, omitTarget bool) Snapshot {
	t.Helper()
	s := healthySnapshot(t)
	s.Views[3].Reachable, s.Views[3].RedisID, s.Views[3].Nodes = false, "", nil
	s.Views[3].Errors = []string{"connection refused"}
	s.Views[0].Replication["connected_slaves"] = "0"
	for i := range s.Views {
		if i == 3 {
			continue
		}
		s.Views[i].Nodes[3].Flags = []string{"slave", "fail"}
		s.Views[i].Nodes[3].LinkState = "disconnected"
		if omitTarget {
			s.Views[i].Nodes = slices.Delete(s.Views[i].Nodes, 3, 4)
		}
		s.Views[i].Coverage = AnalyzeCoverage(s.Views[i].Nodes)
	}
	return s
}

func TestReplicaSelectionRetainsDeterministicPrimaryContext(t *testing.T) {
	s := healthySnapshot(t)
	selected, err := SelectReplica(s)
	if err != nil {
		t.Fatal(err)
	}
	if selected.TargetID != fixtureID(0) || selected.ReplicaID != fixtureID(3) || selected.Replica.ID != "node-4" || len(selected.Slots) != 5461 {
		t.Fatalf("wrong selected replica/shard context: %+v", selected)
	}
	slices.Reverse(s.Views)
	again, err := SelectReplica(s)
	if err != nil || again.Target != selected.Target || again.Replica != selected.Replica || !slices.Equal(again.Slots, selected.Slots) {
		t.Fatal("replica selection depends on observer order")
	}
}

func TestReplicaFaultPreservesServiceWithoutPromotionOrHealthyTargetGossip(t *testing.T) {
	initial := healthySnapshot(t)
	selected, err := SelectReplica(initial)
	if err != nil {
		t.Fatal(err)
	}
	for _, missing := range []bool{false, true} {
		s := replicaFaultSnapshot(t, missing)
		if ok, why := DuringReplicaFault(s, selected, initial); !ok {
			t.Fatalf("replica fault degraded otherwise healthy service: %v", why)
		}
		if ok, why := ReplicaServiceWithBaseline(s, selected, initial); !ok {
			t.Fatalf("runner alias differs: %v", why)
		}
		if ok, _ := Ready(s); ok {
			t.Fatal("five-node service topology was mistaken for full redundancy")
		}
		if ok, _ := ReplicaRecoveryWithBaseline(s, selected, initial); ok {
			t.Fatal("missing replica was mistaken for restored redundancy")
		}
	}
	// Healthy/stale target gossip and connected count can precede failure
	// detection. They must not turn continued service into an artificial outage.
	s := replicaFaultSnapshot(t, false)
	s.Views[0].Replication["connected_slaves"] = "1"
	for i := range s.Views {
		if i != 3 {
			s.Views[i].Nodes[3].Flags = []string{"slave"}
			s.Views[i].Nodes[3].LinkState = "connected"
		}
	}
	if ok, why := DuringReplicaFault(s, selected, initial); !ok {
		t.Fatalf("stale target failure gossip prevented service observation: %v", why)
	}
	if ok, why := DuringReplicaFault(initial, selected, initial); !ok {
		t.Fatalf("healthy service cannot be verified independently of action evidence: %v", why)
	}
}

func TestReplicaFaultRejectsLeaderChangesOrOtherReplicationFailures(t *testing.T) {
	initial := healthySnapshot(t)
	selected, _ := SelectReplica(initial)
	cases := map[string]func(*Snapshot){
		"primary unreachable":                           func(s *Snapshot) { s.Views[0].Reachable = false },
		"unaffected replica unreachable":                func(s *Snapshot) { s.Views[5].Reachable = false },
		"missing managed observation":                   func(s *Snapshot) { s.Views = s.Views[:5] },
		"duplicate managed identity":                    func(s *Snapshot) { s.Views[5].Node.ID = s.Views[4].Node.ID },
		"target endpoints changed":                      func(s *Snapshot) { s.Views[3].Node.Addr = "127.0.0.1:7999" },
		"survivor Redis identity changed":               func(s *Snapshot) { s.Views[5].RedisID = fixtureID(4) },
		"survivor INFO error":                           func(s *Snapshot) { s.Views[1].Errors = []string{"INFO stats: timeout"} },
		"original primary INFO role stale":              func(s *Snapshot) { s.Views[0].Replication["role"] = "slave" },
		"unexpected targeted connected replica count":   func(s *Snapshot) { s.Views[0].Replication["connected_slaves"] = "2" },
		"unaffected primary loses replica":              func(s *Snapshot) { s.Views[1].Replication["connected_slaves"] = "0" },
		"unaffected replica link down":                  func(s *Snapshot) { s.Views[4].Replication["master_link_status"] = "down" },
		"unaffected replica still syncing":              func(s *Snapshot) { s.Views[4].Replication["master_sync_in_progress"] = "1" },
		"missing primary in gossip":                     func(s *Snapshot) { s.Views[2].Nodes = slices.Delete(s.Views[2].Nodes, 0, 1) },
		"unknown topology identity":                     func(s *Snapshot) { s.Views[2].Nodes[5].ID = "unknown" },
		"target replica promoted":                       func(s *Snapshot) { s.Views[2].Nodes[3].Flags = []string{"master"} },
		"replica changed primary":                       func(s *Snapshot) { s.Views[2].Nodes[3].PrimaryID = fixtureID(1) },
		"target directly reports primary":               func(s *Snapshot) { s.Views[3].Replication = map[string]string{"role": "master"} },
		"original leader's advertised endpoint changed": func(s *Snapshot) { s.Views[2].Nodes[0].Addr = "172.22.0.99:6379" },
		"unaffected owner change": func(s *Snapshot) {
			for i := range s.Views {
				if i == 3 {
					continue
				}
				s.Views[i].Nodes[0].Slots = []SlotRange{{0, 5459}}
				s.Views[i].Nodes[1].Slots = []SlotRange{{5460, 10922}}
			}
		},
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			s := replicaFaultSnapshot(t, false)
			mutate(&s)
			if ok, why := DuringReplicaFault(s, selected, initial); ok || len(why) == 0 {
				t.Fatal("unhealthy or changed service topology accepted")
			}
		})
	}
	if ok, _ := DuringReplicaFault(reconciledSnapshot(t), selected, initial); ok {
		t.Fatal("primary failover was accepted as unchanged replica fault topology")
	}
}

func TestReplicaRecoveryRequiresOriginalHealthySixNodeTopology(t *testing.T) {
	initial := healthySnapshot(t)
	selected, _ := SelectReplica(initial)
	if ok, why := ReplicaRecoveryWithBaseline(healthySnapshot(t), selected, initial); !ok {
		t.Fatalf("restored replica rejected: %v", why)
	}
	if ok, why := ReplicaReconciledWithBaseline(healthySnapshot(t), selected, initial); !ok {
		t.Fatalf("runner alias differs: %v", why)
	}
	cases := map[string]func(*Snapshot){
		"target unreachable":           func(s *Snapshot) { s.Views[3].Reachable = false },
		"target link down":             func(s *Snapshot) { s.Views[3].Replication["master_link_status"] = "down" },
		"target syncing":               func(s *Snapshot) { s.Views[3].Replication["master_sync_in_progress"] = "1" },
		"primary no connected replica": func(s *Snapshot) { s.Views[0].Replication["connected_slaves"] = "0" },
		"managed endpoint changed":     func(s *Snapshot) { s.Views[3].Node.InternalAddr = "172.22.0.99:6379" },
		"healthy substituted Redis identity": func(s *Snapshot) {
			id := strings.Repeat("z", 40)
			s.Views[3].RedisID = id
			for i := range s.Views {
				s.Views[i].Nodes[3].ID = id
			}
		},
		"healthy changed advertised endpoint": func(s *Snapshot) {
			for i := range s.Views {
				s.Views[i].Nodes[0].Addr = "172.22.0.99:6379"
			}
			s.Views[3].Replication["master_host"] = "172.22.0.99"
		},
		"healthy swapped replica assignments": func(s *Snapshot) {
			for i := range s.Views {
				s.Views[i].Nodes[3].PrimaryID = fixtureID(1)
				s.Views[i].Nodes[4].PrimaryID = fixtureID(0)
			}
			s.Views[3].Replication["master_host"] = "172.22.0.3"
			s.Views[4].Replication["master_host"] = "172.22.0.2"
		},
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			s := healthySnapshot(t)
			mutate(&s)
			if ok, why := ReplicaRecoveryWithBaseline(s, selected, initial); ok || len(why) == 0 {
				t.Fatal("incomplete or changed replica restoration accepted")
			}
		})
	}
	if ok, _ := ReplicaRecoveryWithBaseline(reconciledSnapshot(t), selected, initial); ok {
		t.Fatal("healthy promoted shard accepted as unchanged replica restoration")
	}
	// This predicate proves sampled synchronization readiness, not zero offset lag.
	s := healthySnapshot(t)
	s.Views[3].Replication["slave_repl_offset"] = "100"
	s.Views[0].Replication["master_repl_offset"] = "200"
	if ok, why := ReplicaRecoveryWithBaseline(s, selected, initial); !ok {
		t.Fatalf("replication progress confused with durability: %v", why)
	}
}

func TestReplicaPredicatesRejectInvalidSelectionAndBaseline(t *testing.T) {
	for _, kind := range []string{"unready baseline", "unknown target", "wrong replica", "missing slot", "duplicate slot", "invalid slot"} {
		t.Run(kind, func(t *testing.T) {
			initial := healthySnapshot(t)
			selected, _ := SelectReplica(initial)
			switch kind {
			case "unready baseline":
				initial.Views[3].Replication["master_link_status"] = "down"
			case "unknown target":
				selected.Target.ID = "unmanaged"
			case "wrong replica":
				selected.Replica, selected.ReplicaID = initial.Views[4].Node, fixtureID(4)
			case "missing slot":
				selected.Slots = selected.Slots[1:]
			case "duplicate slot":
				selected.Slots = append(selected.Slots, 0)
			case "invalid slot":
				selected.Slots = append(selected.Slots, -1)
			}
			for _, predicate := range []func(Snapshot, Selection, Snapshot) (bool, []string){DuringReplicaFault, ReplicaRecoveryWithBaseline} {
				if ok, why := predicate(healthySnapshot(t), selected, initial); ok || len(why) == 0 {
					t.Fatal("invalid input accepted")
				}
			}
		})
	}
}
