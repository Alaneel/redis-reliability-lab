package diagnosis

import (
	"encoding/json"
	"sort"
	"strings"
	"time"
)

// Reconstruction records observed transitions, including recovery regressions.
// It does not import laboratory recovery milestones or fault start times.
func observedTransitions(in Incident) []Event {
	last := map[string]Evidence{}
	events := []Event{}
	for _, e := range in.Evidence {
		if e.Status != Valid {
			continue
		}
		key := e.Source + "|" + e.Node + "|" + e.Signal
		prev := last[key]
		last[key] = e
		kind := ""
		switch e.Signal {
		case "peer.flags":
			var flags, old []string
			value(e, &flags)
			value(prev, &old)
			if (contains(flags, "fail") || contains(flags, "fail?")) && !contains(old, "fail") && !contains(old, "fail?") {
				kind = "cluster_failure_flag_observed"
			}
		case "peer.slots":
			if prev.ID != "" && prev.Status == Valid {
				var before, after any
				json.Unmarshal(prev.Value, &before)
				json.Unmarshal(e.Value, &after)
				b, _ := json.Marshal(before)
				c, _ := json.Marshal(after)
				if string(b) != string(c) {
					kind = "slot_ownership_change_observed"
				}
			}
		case "client_success":
			b, ok := boolValue(e)
			p, pok := boolValue(prev)
			if ok && pok && b && !p {
				kind = "client_visible_recovery_observed"
			}
			if ok && pok && !b && p {
				kind = "client_recovery_regression_observed"
			}
		case "ack.successful":
			b, ok := boolValue(e)
			p, pok := boolValue(prev)
			if ok && pok && b && !p {
				kind = "acknowledgement_policy_restoration_observed"
			}
		case "replication.connected_slaves":
			n, ok := number(e)
			p, pok := number(prev)
			if ok && pok && n > 0 && p == 0 {
				kind = "primary_connected_replica_count_restored"
			}
		case "replication.master_sync_in_progress":
			n, ok := number(e)
			p, pok := number(prev)
			if ok && pok && n == 0 && p == 1 {
				kind = "replica_synchronization_completion_observed"
			}
		case "process_live":
			b, ok := boolValue(e)
			p, pok := boolValue(prev)
			if ok && pok && b && !p {
				kind = "process_liveness_restored"
			}
		}
		if kind != "" {
			ev := Event{Kind: kind, At: e.At, Timing: "first sampled observation; exact event time unknown", Evidence: []string{e.ID}}
			if prev.ID != "" {
				at := prev.At
				ev.After = &at
				ev.Evidence = append(ev.Evidence, prev.ID)
				ev.Timing = "transition sampled within observation interval; exact time unknown"
			}
			events = append(events, ev)
		}
	}
	type cohort struct {
		at      time.Time
		refs    []string
		healthy bool
		views   int
	}
	cohorts := map[string]*cohort{}
	for _, e := range in.Evidence {
		if e.Signal != "slot_coverage" {
			continue
		}
		prefix, _, ok := strings.Cut(e.Reference, "/views/")
		if !ok {
			continue
		}
		c := cohorts[prefix]
		if c == nil {
			c = &cohort{at: e.At, healthy: true}
			cohorts[prefix] = c
		}
		c.views++
		b, valid := boolValue(e)
		c.healthy = c.healthy && valid && b
		if valid {
			c.refs = append(c.refs, e.ID)
		}
		if c.at.Before(e.At) {
			c.at = e.At
		}
	}
	cs := []*cohort{}
	for _, c := range cohorts {
		cs = append(cs, c)
	}
	sort.Slice(cs, func(i, j int) bool { return cs[i].at.Before(cs[j].at) })
	var prior *cohort
	for _, c := range cs {
		if c.views < 2 {
			continue
		}
		kind := ""
		if prior == nil && c.healthy {
			kind = "baseline_full_slot_coverage_observed"
		} else if prior != nil && !prior.healthy && c.healthy {
			kind = "slot_coverage_convergence_observed"
		}
		if kind != "" {
			event := Event{Kind: kind, At: c.at, Timing: "all recorded views in this non-atomic sample report full healthy coverage", Evidence: c.refs}
			if prior != nil {
				at := prior.at
				event.After = &at
			}
			events = append(events, event)
		}
		prior = c
	}
	return events
}
