package cluster

import (
	"strings"
	"testing"
)

func reconciledSnapshot(t *testing.T) Snapshot {
	t.Helper()
	s := healthySnapshot(t)
	for i := range s.Views {
		former := &s.Views[i].Nodes[0]
		former.Flags = []string{"slave"}
		if i == 0 {
			former.Flags = append(former.Flags, "myself")
		}
		former.PrimaryID = fixtureID(3)
		former.Slots = nil
		promoted := &s.Views[i].Nodes[3]
		promoted.Flags = []string{"master"}
		if i == 3 {
			promoted.Flags = append(promoted.Flags, "myself")
		}
		promoted.PrimaryID = ""
		promoted.Slots = []SlotRange{{0, 5460}}
		s.Views[i].Coverage = AnalyzeCoverage(s.Views[i].Nodes)
	}
	s.Views[0].Replication = map[string]string{"role": "slave", "master_link_status": "up", "master_sync_in_progress": "0", "master_host": "172.22.0.5", "master_port": "6379"}
	s.Views[3].Replication = map[string]string{"role": "master", "connected_slaves": "1"}
	return s
}

func TestReconciledWithBaselineRequiresPromotionAndRestoredRedundancy(t *testing.T) {
	baseline := healthySnapshot(t)
	selected, err := SelectPrimary(baseline)
	if err != nil {
		t.Fatal(err)
	}
	s := reconciledSnapshot(t)
	if ready, why := Ready(s); !ready {
		t.Fatalf("fixture is not healthy after reconciliation: %v", why)
	}
	if ok, why := ReconciledWithBaseline(s, selected, baseline); !ok {
		t.Fatalf("reconciled topology rejected: %v", why)
	}
	if ok, why := ReconciledWithBaseline(healthySnapshot(t), selected, baseline); ok || len(why) == 0 {
		t.Fatal("unchanged original topology was accepted as promoted reconciliation")
	}
	if ok, why := ReconciledWithBaseline(recoveredSnapshot(t), selected, baseline); ok || len(why) == 0 {
		t.Fatal("five-node service recovery was accepted as restored redundancy")
	}
	// Observer order, local gossip clocks and epochs do not identify a different
	// topology. They can legitimately differ after failover and reconnection.
	for i := range s.Views {
		for j := range s.Views[i].Nodes {
			s.Views[i].Nodes[j].ConfigEpoch += uint64(i + 10)
			s.Views[i].Nodes[j].PingSent += int64(j + 100)
			s.Views[i].Nodes[j].PongReceived += int64(i + 200)
		}
	}
	for i, j := 0, len(s.Views)-1; i < j; i, j = i+1, j-1 {
		s.Views[i], s.Views[j] = s.Views[j], s.Views[i]
	}
	if ok, why := ReconciledWithBaseline(s, selected, baseline); !ok {
		t.Fatalf("observer-local timing/epoch changes affected reconciliation: %v", why)
	}
}

func TestReconciliationRejectsMissingOrUnhealthyReplicationEvidence(t *testing.T) {
	baseline := healthySnapshot(t)
	selected, err := SelectPrimary(baseline)
	if err != nil {
		t.Fatal(err)
	}
	cases := map[string]func(*Snapshot){
		"former primary unreachable": func(s *Snapshot) { s.Views[0].Reachable = false },
		"survivor unreachable":       func(s *Snapshot) { s.Views[5].Reachable = false },
		"missing observation":        func(s *Snapshot) { s.Views = s.Views[:5] },
		"processes running without cluster evidence": func(s *Snapshot) {
			for i := range s.Views {
				s.Views[i].ClusterInfo = nil
			}
		},
		"former primary link down":     func(s *Snapshot) { s.Views[0].Replication["master_link_status"] = "down" },
		"former primary synchronizing": func(s *Snapshot) { s.Views[0].Replication["master_sync_in_progress"] = "1" },
		"wrong replication host":       func(s *Snapshot) { s.Views[0].Replication["master_host"] = "172.22.0.3" },
		"wrong replication port":       func(s *Snapshot) { s.Views[0].Replication["master_port"] = "6380" },
		"missing replication link":     func(s *Snapshot) { delete(s.Views[0].Replication, "master_link_status") },
		"promoted primary INFO stale":  func(s *Snapshot) { s.Views[3].Replication["role"] = "slave" },
		"promoted primary lacks connected replica": func(s *Snapshot) {
			s.Views[3].Replication["connected_slaves"] = "0"
		},
		"stale former primary role in gossip": func(s *Snapshot) {
			s.Views[2].Nodes[0].Flags = []string{"master"}
			s.Views[2].Nodes[0].PrimaryID = ""
		},
		"stale promoted replica role in gossip": func(s *Snapshot) {
			s.Views[2].Nodes[3].Flags = []string{"slave"}
			s.Views[2].Nodes[3].PrimaryID = fixtureID(0)
		},
		"partial target slot ownership": func(s *Snapshot) { s.Views[2].Nodes[3].Slots = []SlotRange{{0, 5459}} },
		"former primary still owns slots": func(s *Snapshot) {
			s.Views[2].Nodes[0].Slots = []SlotRange{{0, 5460}}
		},
		"former primary still FAIL": func(s *Snapshot) { s.Views[2].Nodes[0].Flags = append(s.Views[2].Nodes[0].Flags, "fail") },
		"multiple primary roles": func(s *Snapshot) {
			s.Views[2].Nodes[5].Flags = []string{"master"}
			s.Views[2].Nodes[5].PrimaryID = ""
		},
		"active slot migration":                   func(s *Snapshot) { s.Views[2].Nodes[3].SlotTransitions = []string{"[42->-" + fixtureID(1) + "]"} },
		"INFO error despite reachable connection": func(s *Snapshot) { s.Views[0].Errors = []string{"INFO replication: timeout"} },
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			s := reconciledSnapshot(t)
			mutate(&s)
			if ok, why := ReconciledWithBaseline(s, selected, baseline); ok || len(why) == 0 {
				t.Fatal("incomplete reconciliation accepted")
			}
		})
	}
}

func TestReconciliationRejectsHealthyButDifferentIdentitiesShardsOrReplication(t *testing.T) {
	baseline := healthySnapshot(t)
	selected, err := SelectPrimary(baseline)
	if err != nil {
		t.Fatal(err)
	}
	cases := map[string]func(*Snapshot){
		"unknown managed identity":     func(s *Snapshot) { s.Views[5].Node.ID = "node-substitution" },
		"different published endpoint": func(s *Snapshot) { s.Views[0].Node.Addr = "127.0.0.1:7999" },
		"different internal endpoint":  func(s *Snapshot) { s.Views[0].Node.InternalAddr = "172.22.0.99:6379" },
		"replacement Redis identity": func(s *Snapshot) {
			id := strings.Repeat("z", 40)
			s.Views[5].RedisID = id
			for i := range s.Views {
				s.Views[i].Nodes[5].ID = id
			}
		},
		"unaffected slot moves to promoted primary": func(s *Snapshot) {
			for i := range s.Views {
				s.Views[i].Nodes[1].Slots = []SlotRange{{5462, 10922}}
				s.Views[i].Nodes[3].Slots = []SlotRange{{0, 5461}}
			}
		},
		"unaffected replicas swap primaries": func(s *Snapshot) {
			for i := range s.Views {
				s.Views[i].Nodes[4].PrimaryID = fixtureID(2)
				s.Views[i].Nodes[5].PrimaryID = fixtureID(1)
			}
			s.Views[4].Replication["master_host"] = "172.22.0.4"
			s.Views[5].Replication["master_host"] = "172.22.0.3"
		},
		"wrong replica promoted": func(s *Snapshot) {
			for self := range s.Views {
				for i := range s.Views[self].Nodes {
					n := &s.Views[self].Nodes[i]
					if i != 0 && i != 3 && i != 4 {
						continue
					}
					n.Flags = []string{"slave"}
					n.Slots = nil
					switch i {
					case 0:
						n.PrimaryID = fixtureID(4)
					case 3:
						n.PrimaryID = fixtureID(1)
					case 4:
						n.Flags = []string{"master"}
						n.PrimaryID = ""
						n.Slots = []SlotRange{{0, 5460}}
					}
					if i == self {
						n.Flags = append(n.Flags, "myself")
					}
				}
			}
			s.Views[0].Replication["master_host"] = "172.22.0.6"
			s.Views[3].Replication = map[string]string{"role": "slave", "master_link_status": "up", "master_sync_in_progress": "0", "master_host": "172.22.0.3", "master_port": "6379"}
			s.Views[4].Replication = map[string]string{"role": "master", "connected_slaves": "1"}
		},
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			s := reconciledSnapshot(t)
			mutate(&s)
			if ready, why := Ready(s); !ready {
				t.Fatalf("test must preserve a generally healthy six-node cluster: %v", why)
			}
			if ok, why := ReconciledWithBaseline(s, selected, baseline); ok || len(why) == 0 {
				t.Fatal("healthy but different topology was accepted as reconciliation")
			}
		})
	}
}

func TestReconciliationRejectsInvalidBaselineOrSelection(t *testing.T) {
	for _, kind := range []string{"baseline unhealthy", "empty target", "same IDs", "wrong original replica", "wrong managed target", "wrong target endpoint", "missing target slot", "extra target slot", "duplicate target slot", "negative target slot", "out of range target slot"} {
		t.Run(kind, func(t *testing.T) {
			baseline := healthySnapshot(t)
			selected, err := SelectPrimary(baseline)
			if err != nil {
				t.Fatal(err)
			}
			switch kind {
			case "baseline unhealthy":
				baseline.Views[4].Replication["master_link_status"] = "down"
			case "empty target":
				selected.TargetID = ""
			case "same IDs":
				selected.ReplicaID = selected.TargetID
			case "wrong original replica":
				selected.ReplicaID, selected.Replica = fixtureID(4), baseline.Views[4].Node
			case "wrong managed target":
				selected.Target.ID = "unknown"
			case "wrong target endpoint":
				selected.Target.Addr = "127.0.0.1:7999"
			case "missing target slot":
				selected.Slots = selected.Slots[1:]
			case "extra target slot":
				selected.Slots = append(selected.Slots, 5461)
			case "duplicate target slot":
				selected.Slots = append(selected.Slots, 0)
			case "negative target slot":
				selected.Slots = append(selected.Slots, -1)
			case "out of range target slot":
				selected.Slots = append(selected.Slots, SlotCount)
			default:
				t.Fatalf("unhandled case %s", kind)
			}
			if ok, why := ReconciledWithBaseline(reconciledSnapshot(t), selected, baseline); ok || len(why) == 0 {
				t.Fatal("invalid reconciliation inputs accepted")
			}
		})
	}
}
