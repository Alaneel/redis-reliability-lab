package cluster

import (
	"fmt"
	"slices"
	"strings"
)

// SelectReplica deterministically selects the replica of the primary owning
// slot zero. Selection retains the original primary in Target and the actual
// replica fault target in Replica, preserving shard context for all predicates.
func SelectReplica(s Snapshot) (Selection, error) { return SelectPrimary(s) }

func replicaBaseline(selected Selection, baseline Snapshot) (map[string]View, []NodeState, []string, []string) {
	if ready, why := Ready(baseline); !ready {
		return nil, nil, nil, []string{"initial snapshot was not ready: " + strings.Join(why, "; ")}
	}
	known := make(map[string]View)
	for _, v := range baseline.Views {
		known[v.Node.ID] = v
	}
	nodes := baseline.Views[0].Nodes
	owners := ownerMap(nodes)
	var reasons []string
	if selected.TargetID == "" || selected.ReplicaID == "" || selected.TargetID == selected.ReplicaID || len(selected.Slots) == 0 {
		reasons = append(reasons, "invalid replica lifecycle selection")
	}
	if v := known[selected.Target.ID]; v.RedisID != selected.TargetID || v.Node != selected.Target {
		reasons = append(reasons, "selected primary does not match initial managed identity and endpoints")
	}
	if v := known[selected.Replica.ID]; v.RedisID != selected.ReplicaID || v.Node != selected.Replica {
		reasons = append(reasons, "selected replica does not match initial managed identity and endpoints")
	}
	if n, exists := findNode(nodes, selected.TargetID); !exists || !n.Primary() || n.Replica() {
		reasons = append(reasons, "selected primary was not an original primary")
	}
	if n, exists := findNode(nodes, selected.ReplicaID); !exists || !n.Replica() || n.Primary() || n.PrimaryID != selected.TargetID {
		reasons = append(reasons, "selected replica did not follow the selected original primary")
	}
	slots := make(map[int]bool)
	for _, slot := range selected.Slots {
		if slot < 0 || slot >= SlotCount || slots[slot] {
			reasons = append(reasons, "invalid or duplicate selected slot")
		}
		slots[slot] = true
	}
	for slot, owner := range owners {
		if (owner == selected.TargetID) != slots[slot] {
			reasons = append(reasons, "selected slots do not exactly match original primary ownership")
			break
		}
	}
	return known, nodes, owners, unique(reasons)
}

func originalReplicaTopology(v View, initial []NodeState, owners []string, missingAllowed string) []string {
	var reasons []string
	seen := make(map[string]bool)
	for _, n := range v.Nodes {
		if seen[n.ID] {
			reasons = append(reasons, v.Node.ID+" sees a duplicate topology identity")
		}
		seen[n.ID] = true
		original, found := findNode(initial, n.ID)
		if !found {
			reasons = append(reasons, v.Node.ID+" sees an unknown topology identity")
			continue
		}
		if n.Addr != original.Addr || n.Primary() != original.Primary() || n.Replica() != original.Replica() || n.PrimaryID != original.PrimaryID {
			reasons = append(reasons, fmt.Sprintf("%s sees node %s change endpoint, role or replication primary", v.Node.ID, n.ID))
		}
		if len(n.SlotTransitions) != 0 || n.Replica() && len(n.Slots) != 0 {
			reasons = append(reasons, v.Node.ID+" sees slot migration or a replica owning slots")
		}
		if n.ID != missingAllowed && !n.Healthy() {
			reasons = append(reasons, fmt.Sprintf("%s sees unhealthy unaffected node %s", v.Node.ID, n.ID))
		}
	}
	for _, original := range initial {
		if original.ID != missingAllowed && !seen[original.ID] {
			reasons = append(reasons, fmt.Sprintf("%s is missing unaffected topology identity %s", v.Node.ID, original.ID))
		}
	}
	if !slices.Equal(owners, ownerMap(v.Nodes)) {
		reasons = append(reasons, v.Node.ID+" sees original slot ownership change")
	}
	return reasons
}

// DuringReplicaFault evaluates service topology independently of replica
// redundancy. All other five nodes, original primary ownership and unaffected
// replication links must remain healthy. The selected replica may be missing or
// marked failed in gossip, without requiring replica FAIL convergence. A
// healthy pre-fault topology can satisfy this predicate; action/process evidence
// must independently establish that the requested replica fault occurred.
func DuringReplicaFault(s Snapshot, selected Selection, baseline Snapshot) (bool, []string) {
	known, initial, owners, reasons := replicaBaseline(selected, baseline)
	if len(reasons) != 0 {
		return false, reasons
	}
	if len(s.Views) != 6 {
		reasons = append(reasons, "replica fault requires all 6 managed observations including the target")
	}
	seen := make(map[string]bool)
	for _, v := range s.Views {
		original, exists := known[v.Node.ID]
		if !exists || seen[v.Node.ID] || v.Node != original.Node {
			reasons = append(reasons, "missing, duplicate or changed managed identity/endpoints")
		}
		seen[v.Node.ID] = true
		if v.Node.ID == selected.Replica.ID {
			if v.RedisID != "" && v.RedisID != selected.ReplicaID {
				reasons = append(reasons, "target replica changed Redis identity")
			}
			if self, found := findNode(v.Nodes, selected.ReplicaID); found && (!self.Replica() || self.Primary() || self.PrimaryID != selected.TargetID || len(self.Slots) != 0) || v.Replication["role"] == "master" {
				reasons = append(reasons, "target replica was promoted or changed replication primary")
			}
			continue
		}
		reasons = append(reasons, healthyView(v)...)
		if v.RedisID != original.RedisID {
			reasons = append(reasons, v.Node.ID+" changed Redis identity")
		}
		reasons = append(reasons, originalReplicaTopology(v, initial, owners, selected.ReplicaID)...)
		self, found := findNode(v.Nodes, v.RedisID)
		if !found {
			continue
		}
		if self.Replica() {
			reasons = append(reasons, replicaReady(v, self)...)
		} else if self.Primary() {
			if v.Replication["role"] != "master" {
				reasons = append(reasons, v.Node.ID+" INFO role disagrees with original primary")
			}
			count := v.Replication["connected_slaves"]
			if self.ID == selected.TargetID {
				// A stale connected count is valid shortly after isolation. The
				// interruption and lost redundancy are measured separately.
				if count != "0" && count != "1" {
					reasons = append(reasons, v.Node.ID+" has an unexpected targeted-shard replica count")
				}
			} else if count != "1" {
				reasons = append(reasons, v.Node.ID+" lost an unaffected replica link")
			}
		}
	}
	for id := range known {
		if !seen[id] {
			reasons = append(reasons, "missing managed node observation: "+id)
		}
	}
	reasons = unique(reasons)
	return len(reasons) == 0, reasons
}

// ReplicaRecoveryWithBaseline requires the complete original topology after
// replica restart or reconnect: all original identities/endpoints, unchanged
// owners and replica assignments, link-up and completed initial synchronization.
func ReplicaRecoveryWithBaseline(s Snapshot, selected Selection, baseline Snapshot) (bool, []string) {
	known, initial, owners, reasons := replicaBaseline(selected, baseline)
	if len(reasons) != 0 {
		return false, reasons
	}
	_, why := Ready(s)
	reasons = append(reasons, why...)
	for _, v := range s.Views {
		original, exists := known[v.Node.ID]
		if !exists || v.Node != original.Node || v.RedisID != original.RedisID {
			reasons = append(reasons, "restored managed/Redis identity or endpoints differ from the baseline")
		}
		reasons = append(reasons, originalReplicaTopology(v, initial, owners, "")...)
	}
	reasons = unique(reasons)
	return len(reasons) == 0, reasons
}

// Runner-facing names retain the same service/restoration semantics.
func ReplicaServiceWithBaseline(s Snapshot, selected Selection, baseline Snapshot) (bool, []string) {
	return DuringReplicaFault(s, selected, baseline)
}

func ReplicaReconciledWithBaseline(s Snapshot, selected Selection, baseline Snapshot) (bool, []string) {
	return ReplicaRecoveryWithBaseline(s, selected, baseline)
}
