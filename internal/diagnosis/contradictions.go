package diagnosis

import (
	"encoding/json"
	"sort"
	"strings"
)

// peerDisagreements compares simultaneous views by subject, retaining both
// witnesses. Different observations at different sample times are not conflicts.
func peerDisagreements(in Incident) []Finding {
	groups := map[string][]Evidence{}
	for _, e := range in.Evidence {
		if e.Status != Valid || !strings.HasPrefix(e.Source, "peer_view:") || (e.Signal != "peer.flags" && e.Signal != "peer.slots" && e.Signal != "peer.primary_id") {
			continue
		}
		snapshot, _, ok := strings.Cut(e.Reference, "/views/")
		if !ok {
			continue
		}
		key := snapshot + "|" + e.Node + "|" + e.Signal
		groups[key] = append(groups[key], e)
	}
	keys := []string{}
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	out := []Finding{}
	for _, k := range keys {
		es := groups[k]
		if len(es) < 2 {
			continue
		}
		first := canonical(es[0])
		for _, e := range es[1:] {
			if canonical(e) != first {
				out = append(out, Finding{Category: "topology_disagreement", Description: "Contemporaneous peer views disagree about the same node; no single view is chosen as authoritative.", Nodes: []string{e.Node}, Support: []string{es[0].ID, e.ID}, Against: []string{}, Missing: []string{"Atomic topology snapshot is unavailable"}, Status: "strongly_supported", Rule: "R14-conflicting-contemporaneous-peer-views", Alternatives: []string{"Normal gossip convergence", "Stale peer view", "Partitioned membership"}, Verification: []string{"Collect repeated independent views and compare per-node epochs and membership"}})
				break
			}
		}
	}
	if len(out) > 1 {
		combined := out[0]
		for _, f := range out[1:] {
			if !contains(combined.Nodes, f.Nodes[0]) {
				combined.Nodes = append(combined.Nodes, f.Nodes[0])
			}
			for _, id := range f.Support {
				if !contains(combined.Support, id) && len(combined.Support) < 128 {
					combined.Support = append(combined.Support, id)
				}
			}
		}
		sort.Strings(combined.Nodes)
		return []Finding{combined}
	}
	return out
}
func canonical(e Evidence) string {
	if e.Signal == "peer.flags" {
		var flags []string
		if json.Unmarshal(e.Value, &flags) == nil {
			filtered := flags[:0]
			for _, flag := range flags {
				if flag != "myself" {
					filtered = append(filtered, flag)
				}
			}
			flags = filtered
			sort.Strings(flags)
			b, _ := json.Marshal(flags)
			return string(b)
		}
	}
	return string(e.Value)
}
