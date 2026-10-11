package cluster

import (
	"fmt"
	"net"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// Ready requires evidence from every managed node. Link-up and a completed
// initial synchronization are readiness conditions; they do not imply a zero
// replication lag or a durability guarantee for subsequent writes.
func Ready(s Snapshot) (bool, []string) {
	var reasons []string
	if len(s.Views) != 6 {
		reasons = append(reasons, fmt.Sprintf("expected 6 managed observations, got %d", len(s.Views)))
	}
	managed, direct := make(map[string]bool), make(map[string]View)
	for _, v := range s.Views {
		if managed[v.Node.ID] || v.Node.ID == "" {
			reasons = append(reasons, "missing or duplicate managed node identity")
		}
		managed[v.Node.ID] = true
		reasons = append(reasons, healthyView(v)...)
		if v.RedisID == "" || direct[v.RedisID].RedisID != "" {
			reasons = append(reasons, "missing or duplicate directly observed Redis identity")
		}
		direct[v.RedisID] = v
		primaries, replicas := 0, 0
		links := make(map[string]int)
		for _, n := range v.Nodes {
			if !n.Healthy() {
				reasons = append(reasons, fmt.Sprintf("%s sees unhealthy node %s", v.Node.ID, n.ID))
			}
			if len(n.SlotTransitions) != 0 {
				reasons = append(reasons, fmt.Sprintf("%s sees slot migration on %s", v.Node.ID, n.ID))
			}
			if n.Primary() && (n.PrimaryID != "" || len(n.Slots) == 0) {
				reasons = append(reasons, fmt.Sprintf("%s sees primary %s without its own slot-bearing shard", v.Node.ID, n.ID))
			}
			switch {
			case n.Primary():
				primaries++
			case n.Replica():
				replicas++
				links[n.PrimaryID]++
			default:
				reasons = append(reasons, fmt.Sprintf("%s sees node %s without a primary/replica role", v.Node.ID, n.ID))
			}
		}
		if len(v.Nodes) != 6 || primaries != 3 || replicas != 3 {
			reasons = append(reasons, fmt.Sprintf("%s topology has %d nodes, %d primaries and %d replicas", v.Node.ID, len(v.Nodes), primaries, replicas))
		}
		for _, n := range v.Nodes {
			if n.Primary() && links[n.ID] != 1 {
				reasons = append(reasons, fmt.Sprintf("%s sees primary %s with %d replicas", v.Node.ID, n.ID, links[n.ID]))
			}
		}
	}
	reasons = append(reasons, convergence(s.Views)...)
	for _, v := range s.Views {
		self, ok := findNode(v.Nodes, v.RedisID)
		if !ok {
			continue
		}
		if self.Primary() {
			if v.Replication["role"] != "master" || v.Replication["connected_slaves"] != "1" {
				reasons = append(reasons, fmt.Sprintf("%s does not report a primary with one connected replica", v.Node.ID))
			}
		} else if self.Replica() {
			reasons = append(reasons, replicaReady(v, self)...)
			if _, exists := direct[self.PrimaryID]; !exists {
				reasons = append(reasons, fmt.Sprintf("%s follows an unmanaged primary %s", v.Node.ID, self.PrimaryID))
			}
		}
		for _, n := range v.Nodes {
			if _, exists := direct[n.ID]; !exists {
				reasons = append(reasons, fmt.Sprintf("%s topology contains unmanaged node %s", v.Node.ID, n.ID))
			}
		}
	}
	reasons = unique(reasons)
	return len(reasons) == 0, reasons
}

// SelectPrimary deterministically chooses the primary owning the lowest slot,
// so the injected fault is reproducible despite randomized Redis node IDs.
func SelectPrimary(s Snapshot) (Selection, error) {
	if ready, reasons := Ready(s); !ready {
		return Selection{}, fmt.Errorf("cluster not ready: %s", strings.Join(reasons, "; "))
	}
	view := s.Views[0]
	owners := ownerMap(view.Nodes)
	targetID := owners[0]
	var selected Selection
	selected.TargetID = targetID
	for _, n := range view.Nodes {
		if n.Replica() && n.PrimaryID == targetID {
			selected.ReplicaID = n.ID
		}
	}
	for _, v := range s.Views {
		if v.RedisID == selected.TargetID {
			selected.Target = v.Node
		}
		if v.RedisID == selected.ReplicaID {
			selected.Replica = v.Node
		}
	}
	for slot, owner := range owners {
		if owner == targetID {
			selected.Slots = append(selected.Slots, slot)
		}
	}
	if selected.Target.ID == "" || selected.Replica.ID == "" {
		return Selection{}, fmt.Errorf("selected shard is not mapped to managed nodes")
	}
	return selected, nil
}

// Recovery requires all five surviving managed nodes to agree on a healthy,
// fully covered slot map and on promotion of the designated original replica.
// Call RecoveryWithBaseline to additionally verify unaffected slot owners.
func Recovery(s Snapshot, targetID, replicaID string, targetSlots []int) (bool, []string) {
	var reasons []string
	if len(s.Views) != 6 {
		reasons = append(reasons, "recovery requires all 6 managed observations, including the killed node")
	}
	if targetID == "" || replicaID == "" || targetID == replicaID || len(targetSlots) == 0 {
		reasons = append(reasons, "invalid failover selection")
	}
	slots := make(map[int]bool)
	for _, slot := range targetSlots {
		if slot < 0 || slot >= SlotCount || slots[slot] {
			reasons = append(reasons, "invalid or duplicate target slot")
			continue
		}
		slots[slot] = true
	}
	var survivors []View
	managed, redisIDs := make(map[string]bool), make(map[string]bool)
	for _, v := range s.Views {
		if v.Node.ID == "" || managed[v.Node.ID] {
			reasons = append(reasons, "missing or duplicate managed node identity")
		}
		managed[v.Node.ID] = true
		if !v.Reachable {
			continue
		}
		survivors = append(survivors, v)
		reasons = append(reasons, healthyView(v)...)
		if v.RedisID == targetID {
			reasons = append(reasons, "killed original primary remains directly reachable")
		}
		if v.RedisID == "" || redisIDs[v.RedisID] {
			reasons = append(reasons, "missing or duplicate survivor Redis identity")
		}
		redisIDs[v.RedisID] = true
		owners := ownerMap(v.Nodes)
		for slot := range slots {
			if owners[slot] != replicaID {
				reasons = append(reasons, fmt.Sprintf("%s has target slots not owned by promoted replica", v.Node.ID))
				break
			}
		}
		promoted, exists := findNode(v.Nodes, replicaID)
		if !exists || !promoted.Primary() || !promoted.Healthy() {
			reasons = append(reasons, fmt.Sprintf("%s does not see healthy replica promotion", v.Node.ID))
		}
		primaries, replicas := 0, 0
		for _, n := range v.Nodes {
			if len(n.SlotTransitions) != 0 {
				reasons = append(reasons, fmt.Sprintf("%s sees active slot migration", v.Node.ID))
			}
			if n.ID == targetID {
				if len(n.Slots) > 0 || !n.HasFlag("fail") {
					reasons = append(reasons, fmt.Sprintf("%s still sees original primary owning slots or without FAIL", v.Node.ID))
				}
				continue
			}
			if !n.Healthy() {
				reasons = append(reasons, fmt.Sprintf("%s sees unhealthy surviving node %s", v.Node.ID, n.ID))
			}
			if n.Primary() {
				primaries++
			} else if n.Replica() {
				replicas++
			} else {
				reasons = append(reasons, fmt.Sprintf("%s sees an unknown surviving node role", v.Node.ID))
			}
		}
		if primaries != 3 || replicas != 2 {
			reasons = append(reasons, fmt.Sprintf("%s sees %d surviving primaries and %d replicas, expected 3 and 2", v.Node.ID, primaries, replicas))
		}
		self, exists := findNode(v.Nodes, v.RedisID)
		if exists {
			if self.Replica() {
				reasons = append(reasons, replicaReady(v, self)...)
			}
			if self.Primary() && v.Replication["role"] != "master" {
				reasons = append(reasons, fmt.Sprintf("%s INFO role disagrees with primary topology", v.Node.ID))
			}
			if self.Primary() {
				linkedReplicas := 0
				for _, n := range v.Nodes {
					if n.Replica() && n.PrimaryID == self.ID && n.Healthy() {
						linkedReplicas++
					}
				}
				if v.Replication["connected_slaves"] != strconv.Itoa(linkedReplicas) {
					reasons = append(reasons, fmt.Sprintf("%s INFO replica count disagrees with surviving topology", v.Node.ID))
				}
			}
		}
	}
	if len(survivors) != 5 {
		reasons = append(reasons, fmt.Sprintf("expected 5 reachable surviving nodes, got %d", len(survivors)))
	}
	if !redisIDs[replicaID] {
		reasons = append(reasons, "promoted replica is not directly reachable")
	}
	for _, v := range survivors {
		for _, n := range v.Nodes {
			if n.ID != targetID && !redisIDs[n.ID] {
				reasons = append(reasons, fmt.Sprintf("%s sees unmanaged surviving node %s", v.Node.ID, n.ID))
			}
		}
	}
	reasons = append(reasons, convergence(survivors)...)
	reasons = unique(reasons)
	return len(reasons) == 0, reasons
}

func RecoveryWithBaseline(s Snapshot, selected Selection, baseline Snapshot) (bool, []string) {
	_, reasons := Recovery(s, selected.TargetID, selected.ReplicaID, selected.Slots)
	if ready, why := Ready(baseline); !ready {
		reasons = append(reasons, "initial snapshot was not ready: "+strings.Join(why, "; "))
	} else {
		original := ownerMap(baseline.Views[0].Nodes)
		target := make(map[int]bool)
		for _, slot := range selected.Slots {
			target[slot] = true
		}
		for slot, owner := range original {
			if (owner == selected.TargetID) != target[slot] {
				reasons = append(reasons, "selected slots do not exactly match original primary ownership")
				break
			}
		}
		known := make(map[string]string)
		for _, v := range baseline.Views {
			known[v.Node.ID] = v.RedisID
		}
		if known[selected.Target.ID] != selected.TargetID || known[selected.Replica.ID] != selected.ReplicaID {
			reasons = append(reasons, "failover selection does not map to initial managed identities")
		}
		originalReplica, exists := findNode(baseline.Views[0].Nodes, selected.ReplicaID)
		if !exists || !originalReplica.Replica() || originalReplica.PrimaryID != selected.TargetID {
			reasons = append(reasons, "selected replacement was not the original primary's replica")
		}
		seen := make(map[string]bool)
		for _, v := range s.Views {
			seen[v.Node.ID] = true
			if _, exists := known[v.Node.ID]; !exists {
				reasons = append(reasons, "observation contains an unknown managed node")
			}
			if v.Node.ID == selected.Target.ID {
				if v.Reachable {
					reasons = append(reasons, "killed managed target is still reachable")
				}
				continue
			}
			if !v.Reachable {
				reasons = append(reasons, fmt.Sprintf("surviving managed node %s is unreachable", v.Node.ID))
				continue
			}
			if v.RedisID != known[v.Node.ID] {
				reasons = append(reasons, fmt.Sprintf("managed node %s changed its Redis identity", v.Node.ID))
			}
			for _, n := range v.Nodes {
				if n.ID == selected.TargetID || n.ID == selected.ReplicaID {
					continue
				}
				initial, found := findNode(baseline.Views[0].Nodes, n.ID)
				if !found || n.Primary() != initial.Primary() || n.Replica() != initial.Replica() || n.PrimaryID != initial.PrimaryID {
					reasons = append(reasons, fmt.Sprintf("%s sees an unaffected node change role or replication primary", v.Node.ID))
				}
			}
			current := ownerMap(v.Nodes)
			for slot, owner := range original {
				if !target[slot] && current[slot] != owner {
					reasons = append(reasons, fmt.Sprintf("%s sees an unaffected slot change owner", v.Node.ID))
					break
				}
			}
		}
		for id := range known {
			if !seen[id] {
				reasons = append(reasons, "missing managed node observation: "+id)
			}
		}
	}
	reasons = unique(reasons)
	return len(reasons) == 0, reasons
}

// ReconciledWithBaseline requires the original six managed nodes to agree on
// completed failover after connectivity is restored. The original replica must
// own the selected shard, and the former primary must follow it with its link up
// and initial synchronization complete. All shards again have one connected
// replica. This proves sampled topology and replication readiness, not zero
// replication lag, write preservation, or a durability guarantee.
func ReconciledWithBaseline(s Snapshot, selected Selection, baseline Snapshot) (bool, []string) {
	_, reasons := Ready(s)
	if ready, why := Ready(baseline); !ready {
		reasons = append(reasons, "initial snapshot was not ready: "+strings.Join(why, "; "))
		return false, unique(reasons)
	}
	if selected.TargetID == "" || selected.ReplicaID == "" || selected.TargetID == selected.ReplicaID || len(selected.Slots) == 0 {
		reasons = append(reasons, "invalid reconciliation selection")
	}
	targetSlots := make(map[int]bool)
	for _, slot := range selected.Slots {
		if slot < 0 || slot >= SlotCount || targetSlots[slot] {
			reasons = append(reasons, "invalid or duplicate target slot")
			continue
		}
		targetSlots[slot] = true
	}
	initialNodes := baseline.Views[0].Nodes
	initialOwners := ownerMap(initialNodes)
	for slot, owner := range initialOwners {
		if (owner == selected.TargetID) != targetSlots[slot] {
			reasons = append(reasons, "selected slots do not exactly match original primary ownership")
			break
		}
	}
	known := make(map[string]View)
	for _, v := range baseline.Views {
		known[v.Node.ID] = v
	}
	if target := known[selected.Target.ID]; target.RedisID != selected.TargetID || target.Node != selected.Target {
		reasons = append(reasons, "selected primary does not map to the initial managed identity and endpoints")
	}
	if replica := known[selected.Replica.ID]; replica.RedisID != selected.ReplicaID || replica.Node != selected.Replica {
		reasons = append(reasons, "selected replacement does not map to the initial managed identity and endpoints")
	}
	if original, exists := findNode(initialNodes, selected.TargetID); !exists || !original.Primary() {
		reasons = append(reasons, "selected target was not an original primary")
	}
	if original, exists := findNode(initialNodes, selected.ReplicaID); !exists || !original.Replica() || original.PrimaryID != selected.TargetID {
		reasons = append(reasons, "selected replacement was not the original primary's replica")
	}
	seen := make(map[string]bool)
	for _, v := range s.Views {
		seen[v.Node.ID] = true
		originalView, exists := known[v.Node.ID]
		if !exists {
			reasons = append(reasons, "observation contains an unknown managed node: "+v.Node.ID)
		} else {
			if v.RedisID != originalView.RedisID {
				reasons = append(reasons, fmt.Sprintf("managed node %s changed its Redis identity", v.Node.ID))
			}
			if v.Node != originalView.Node {
				reasons = append(reasons, fmt.Sprintf("managed node %s changed its managed endpoints", v.Node.ID))
			}
		}
		promoted, exists := findNode(v.Nodes, selected.ReplicaID)
		if !exists || !promoted.Primary() || promoted.Replica() || promoted.PrimaryID != "" || !promoted.Healthy() {
			reasons = append(reasons, fmt.Sprintf("%s does not see the original replica as a healthy primary", v.Node.ID))
		}
		former, exists := findNode(v.Nodes, selected.TargetID)
		if !exists || !former.Replica() || former.Primary() || former.PrimaryID != selected.ReplicaID || len(former.Slots) != 0 || !former.Healthy() {
			reasons = append(reasons, fmt.Sprintf("%s does not see the former primary as a healthy replica of the promoted owner", v.Node.ID))
		}
		// Ready checks direct INFO replication for every node, including the
		// former primary's link, synchronization state, and primary endpoint.
		for _, n := range v.Nodes {
			if n.ID == selected.TargetID || n.ID == selected.ReplicaID {
				continue
			}
			original, found := findNode(initialNodes, n.ID)
			if !found || n.Primary() != original.Primary() || n.Replica() != original.Replica() || n.PrimaryID != original.PrimaryID {
				reasons = append(reasons, fmt.Sprintf("%s sees an unaffected node change role or replication primary", v.Node.ID))
			}
		}
		current := ownerMap(v.Nodes)
		for slot, owner := range initialOwners {
			expected := owner
			if owner == selected.TargetID {
				expected = selected.ReplicaID
			}
			if current[slot] != expected {
				if owner == selected.TargetID {
					reasons = append(reasons, fmt.Sprintf("%s has target slots not owned by the promoted replica", v.Node.ID))
				} else {
					reasons = append(reasons, fmt.Sprintf("%s sees an unaffected slot change owner", v.Node.ID))
				}
				break
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

func healthyView(v View) []string {
	var reasons []string
	if !v.Reachable {
		reasons = append(reasons, v.Node.ID+" is unreachable")
	}
	for _, err := range v.Errors {
		reasons = append(reasons, v.Node.ID+": "+err)
	}
	if v.RedisID == "" {
		reasons = append(reasons, v.Node.ID+" has no observed Redis identity")
	}
	self, exists := findNode(v.Nodes, v.RedisID)
	if !exists || !self.HasFlag("myself") || !self.Healthy() {
		reasons = append(reasons, v.Node.ID+" has no healthy myself topology entry")
	}
	if v.ClusterInfo["cluster_state"] != "ok" {
		reasons = append(reasons, v.Node.ID+" cluster_state is not ok")
	}
	for key, expected := range map[string]string{"cluster_slots_assigned": "16384", "cluster_slots_ok": "16384", "cluster_slots_fail": "0", "cluster_slots_pfail": "0"} {
		if v.ClusterInfo[key] != expected {
			reasons = append(reasons, fmt.Sprintf("%s %s=%q, expected %s", v.Node.ID, key, v.ClusterInfo[key], expected))
		}
	}
	if !AnalyzeCoverage(v.Nodes).Full() {
		reasons = append(reasons, v.Node.ID+" does not have unique, healthy full slot coverage")
	}
	if v.RedisVersion == "" {
		reasons = append(reasons, v.Node.ID+" Redis version is unavailable")
	}
	return reasons
}

func replicaReady(v View, replica NodeState) []string {
	var reasons []string
	if v.Replication["role"] != "slave" || v.Replication["master_link_status"] != "up" || v.Replication["master_sync_in_progress"] != "0" {
		reasons = append(reasons, v.Node.ID+" replica is not linked with initial synchronization complete")
	}
	primary, ok := findNode(v.Nodes, replica.PrimaryID)
	if !ok || !primary.Primary() || !primary.Healthy() {
		reasons = append(reasons, v.Node.ID+" replica does not follow a healthy primary")
		return reasons
	}
	host, port, err := net.SplitHostPort(primary.Addr)
	if err != nil || v.Replication["master_host"] != host || v.Replication["master_port"] != port {
		reasons = append(reasons, v.Node.ID+" replication endpoint does not match its topology primary")
	}
	return reasons
}

func convergence(views []View) []string {
	if len(views) == 0 {
		return []string{"no live topology observations"}
	}
	reference := ownerMap(views[0].Nodes)
	var reasons []string
	for _, v := range views[1:] {
		if !slices.Equal(reference, ownerMap(v.Nodes)) {
			reasons = append(reasons, "slot ownership views have not converged")
			break
		}
	}
	// Replication links and roles must also agree; gossip timestamps/epochs and
	// myself flags are observer-local and deliberately excluded from comparison.
	roles := func(nodes []NodeState) []string {
		var entries []string
		for _, n := range nodes {
			if n.HasFlag("fail") {
				continue
			}
			entries = append(entries, fmt.Sprintf("%s:%t:%t:%s", n.ID, n.Primary(), n.Replica(), n.PrimaryID))
		}
		sort.Strings(entries)
		return entries
	}
	referenceRoles := roles(views[0].Nodes)
	for _, v := range views[1:] {
		if !slices.Equal(referenceRoles, roles(v.Nodes)) {
			reasons = append(reasons, "node role and replication views have not converged")
			break
		}
	}
	return reasons
}

func findNode(nodes []NodeState, id string) (NodeState, bool) {
	for _, n := range nodes {
		if n.ID == id {
			return n, true
		}
	}
	return NodeState{}, false
}

func unique(reasons []string) []string {
	sort.Strings(reasons)
	return slices.Compact(reasons)
}
