package diagnosis

import (
	"context"
	"fmt"
	"sort"
	"strings"
)

// Analyze reasons over observations only. Assessments are deterministic logical
// support levels, not probabilities. Root cause remains withheld unless a future
// independent administrative observation channel can disambiguate causes.
func Analyze(ctx context.Context, in Incident) (Assessment, error) {
	if err := ctx.Err(); err != nil {
		return Assessment{}, err
	}
	if in.SchemaVersion != SchemaVersion || len(in.Evidence) > MaxEvidence {
		return Assessment{}, fmt.Errorf("invalid incident schema or size")
	}
	a := Assessment{SchemaVersion: SchemaVersion, EngineVersion: EngineVersion, Incident: in, Detection: "insufficient_evidence", RootCause: "abstained", Findings: []Finding{}, Timeline: []Event{}, Recommendations: []Recommendation{}, Uncertainty: []string{"Insufficient evidence to identify a unique root cause. Process failure, network disruption and observation-channel failure can be indistinguishable."}, Recovery: []string{}}
	latencyBaseline := Evidence{}
	replicaBaseline := map[string]Evidence{}
	roles := map[string]Evidence{}
	lastReach := map[string]Evidence{}
	lastRole := map[string]Evidence{}
	latest := map[string]Evidence{}
	byRef := map[string]Evidence{}
	coverageBySample := map[string][]Evidence{}
	findings := map[string]*Finding{}
	add := func(category, description, rule string, e Evidence, extra ...Evidence) {
		f := findings[category]
		if f == nil {
			f = &Finding{Category: category, Description: description, Status: "strongly_supported", Rule: rule, Support: []string{}, Against: []string{}, Nodes: []string{}, Missing: []string{}, Alternatives: []string{}, Verification: []string{}}
			findings[category] = f
		}
		for _, v := range append([]Evidence{e}, extra...) {
			if v.ID == "" {
				continue
			}
			if v.Status != Valid {
				f.Missing = append(f.Missing, v.Signal+": "+string(v.Status))
				continue
			}
			if !contains(f.Support, v.ID) && len(f.Support) < 128 {
				f.Support = append(f.Support, v.ID)
			}
			if v.Node != "" && !contains(f.Nodes, v.Node) {
				f.Nodes = append(f.Nodes, v.Node)
			}
		}
	}
	event := func(kind string, e Evidence, prev Evidence) {
		ev := Event{Kind: kind, At: e.At, Timing: "first_sampled_observation; exact event time unknown", Evidence: []string{e.ID}}
		if prev.ID != "" && !prev.At.After(e.At) {
			t := prev.At
			ev.After = &t
			ev.Evidence = append(ev.Evidence, prev.ID)
		}
		a.Timeline = append(a.Timeline, ev)
	}
	for _, e := range in.Evidence {
		if err := ctx.Err(); err != nil {
			return Assessment{}, err
		}
		byRef[e.Reference] = e
		if e.Signal == "slot_coverage" {
			if prefix, _, ok := strings.Cut(e.Reference, "/views/"); ok {
				coverageBySample[prefix] = append(coverageBySample[prefix], e)
			}
		}
	}
	for _, e := range in.Evidence {
		if err := ctx.Err(); err != nil {
			return Assessment{}, err
		}
		key := e.Node + "|" + e.Signal + "|" + e.Source
		latest[key] = e
		if e.Status != Valid {
			continue
		}
		switch e.Signal {
		case "replication.role":
			role := stringValue(e)
			if role != "master" && role != "slave" {
				continue
			}
			if _, ok := roles[e.Node]; !ok {
				roles[e.Node] = e
			}
			prev := lastRole[e.Node]
			if stringValue(prev) == "slave" && role == "master" {
				add("primary_unavailability", "An observed replica became primary; failover or an administrative role change is supported, process crash is not established.", "R1-role-transition", e, prev)
				event("replica_promotion_observed", e, prev)
			}
			lastRole[e.Node] = e
		case "reachable":
			b, ok := boolValue(e)
			if !ok {
				continue
			}
			prev := lastReach[e.Node]
			if !b {
				baseline := roles[e.Node]
				category := "node_unavailability"
				if stringValue(baseline) == "master" {
					category = "primary_unavailability"
				}
				if stringValue(baseline) == "slave" {
					category = "replica_unavailability"
				}
				add(category, "Node observation endpoint unavailable; the cause remains unresolved.", "R2-endpoint-unavailable-with-prior-role", e, baseline)
				if p, ok := boolValue(prev); prev.ID == "" || (ok && p) {
					event("component_unavailability_observed", e, prev)
				}
			} else if p, ok := boolValue(prev); ok && !p {
				event("node_reachability_restored", e, prev)
			}
			lastReach[e.Node] = e
		case "replication.connected_slaves":
			n, ok := number(e)
			if ok && n > 0 {
				replicaBaseline[e.Node] = e
			}
			if ok && n == 0 && replicaBaseline[e.Node].ID != "" {
				add("redundancy_degradation", "A directly observed primary reports no connected replicas; service availability and synchronization remain separate questions.", "R3-primary-zero-connected-replicas-after-positive-baseline", e, replicaBaseline[e.Node])
			}
		case "replication.master_link_status":
			if stringValue(e) == "down" {
				add("replica_unavailability", "Replica INFO reports its primary link down; process and network causes remain alternatives.", "R4-replica-link-down", e)
			}
		case "client_success":
			b, ok := boolValue(e)
			if !ok {
				continue
			}
			if !b {
				add("client_disruption", "The all-shard client probe failed.", "R5-observed-client-probe-failure", e)
				prefix := strings.TrimSuffix(e.Reference, "/probe/successful") + "/snapshot"
				coverage := []Evidence{}
				all := true
				for _, v := range coverageBySample[prefix] {
					{
						coverage = append(coverage, v)
						good, valid := boolValue(v)
						if !valid || !good {
							all = false
						}
					}
				}
				if all && len(coverage) >= 2 {
					add("client_recovery_lag", "Multiple contemporaneous node views report full healthy slot coverage while the client probe fails. This does not establish a specific cached-routing or connection cause.", "R6-full-coverage-client-probe-failure", e, coverage...)
				}
			}
		case "client.p95_ms":
			n, ok := number(e)
			if !ok || n <= 0 {
				continue
			}
			prefix := strings.TrimSuffix(e.Reference, "p95_ms")
			attempts, validAttempts := number(byRef[prefix+"attempts"])
			errors, validErrors := number(byRef[prefix+"errors"])
			if latencyBaseline.ID == "" {
				if validAttempts && validErrors && attempts >= 20 && errors == 0 {
					latencyBaseline = e
				}
				continue
			}
			baseline, _ := number(latencyBaseline)
			if validAttempts && attempts >= 20 && n >= 20 && n >= 5*baseline {
				add("client_latency_change", "Terminal request-window p95 is at least five times an earlier error-free measured window and at least 20 ms. Workload mix and acknowledgement waits can explain this; it is not a calibrated anomaly probability.", "R15-measured-window-p95-5x-and-20ms-with-20-attempts", e, latencyBaseline, byRef[prefix+"attempts"])
			}
		case "client.errors":
			n, ok := number(e)
			if ok && n > 0 {
				add("client_disruption", "Client request window contains terminal errors; policy and SET outcomes must be inspected separately.", "R7-terminal-client-errors", e)
			}
		case "ack.successful":
			good, ok := boolValue(e)
			if !ok || good {
				continue
			}
			prefix := strings.TrimSuffix(e.Reference, "successful")
			set := byRef[prefix+"set_acknowledged"]
			var sets []bool
			if !value(set, &sets) {
				continue
			}
			anySet := false
			for _, b := range sets {
				anySet = anySet || b
			}
			if !anySet {
				continue
			}
			policy := byRef[prefix+"policy"]
			support := []Evidence{set, policy}
			short := false
			if stringValue(policy) == "wait" {
				var counts []int64
				var observed []bool
				value(byRef[prefix+"acknowledgements"], &counts)
				value(byRef[prefix+"acknowledgement_observed"], &observed)
				requested, valid := number(byRef[prefix+"requested_replicas"])
				for i, n := range counts {
					if valid && i < len(sets) && sets[i] && i < len(observed) && observed[i] && float64(n) < requested {
						short = true
					}
				}
				support = append(support, byRef[prefix+"acknowledgements"], byRef[prefix+"acknowledgement_observed"], byRef[prefix+"requested_replicas"])
			} else if stringValue(policy) == "waitaof" {
				var local, replicas []int64
				var observed []bool
				value(byRef[prefix+"waitaof_local_acknowledgements"], &local)
				value(byRef[prefix+"waitaof_replica_acknowledgements"], &replicas)
				value(byRef[prefix+"waitaof_acknowledgement_observed"], &observed)
				reqL, okL := number(byRef[prefix+"waitaof_requested_local"])
				reqR, okR := number(byRef[prefix+"waitaof_requested_replicas"])
				for i, n := range local {
					if okL && okR && i < len(sets) && sets[i] && i < len(observed) && observed[i] && i < len(replicas) && (float64(n) < reqL || float64(replicas[i]) < reqR) {
						short = true
					}
				}
				for _, k := range []string{"waitaof_local_acknowledgements", "waitaof_replica_acknowledgements", "waitaof_acknowledgement_observed", "waitaof_requested_local", "waitaof_requested_replicas"} {
					support = append(support, byRef[prefix+k])
				}
			}
			if short {
				add("acknowledgement_degradation", "At least one acknowledged SET received a valid short WAIT/WAITAOF reply. The SET succeeded; the additional acknowledgement policy was not satisfied.", "R8-acknowledged-set-valid-short-policy-reply", e, support...)
			} else {
				add("acknowledgement_command_failure", "An acknowledged SET did not obtain a successful policy probe; no valid short reply is established.", "R9-set-success-policy-unavailable", e, set, policy)
			}
		case "validation.missing", "validation.corrupt":
			n, ok := number(e)
			if ok && n > 0 {
				add("data_validation_failure", "Final acknowledged-value validation observed missing or corrupt values. Serving requests does not establish data preservation.", "R10-positive-validation-loss-counter", e)
			}
		case "validation.unreadable", "validation.unverified":
			n, ok := number(e)
			if ok && n > 0 {
				add("validation_incomplete", "Some acknowledged values could not be verified; preservation remains unknown for that population.", "R11-positive-unchecked-counter", e)
			}
		case "persistence.loading":
			if stringValue(e) == "1" {
				add("persistence_recovery", "Redis INFO reports loading; readiness and final acknowledged-value preservation require separate evidence.", "R12-loading-observed", e)
			}
		case "persistence.aof_last_write_status":
			if stringValue(e) == "err" {
				add("persistence_recovery", "Redis INFO reports an AOF write error; storage health requires investigation.", "R13-aof-write-error", e)
			}
		}
	}
	categories := make([]string, 0, len(findings))
	for k := range findings {
		categories = append(categories, k)
	}
	sort.Strings(categories)
	for _, k := range categories {
		f := findings[k]
		sort.Strings(f.Nodes)
		f.Alternatives = []string{"Process or endpoint failure", "Network or observer connectivity problem", "Planned administrative change"}
		f.Missing = append(f.Missing, "Independent process liveness and network-path observations", "Per-request routing/cache/redirect traces")
		f.Verification = []string{"Inspect independently collected process and network status", "Compare all node views and client outcomes at aligned timestamps"}
		// Later contrary observations are retained as recovery evidence, not used to
		// erase a real earlier incident or present conflicting views as consensus.
		for _, e := range in.Evidence {
			if e.Status != Valid {
				continue
			}
			b, ok := boolValue(e)
			if !ok || !b {
				continue
			}
			if (k == "client_disruption" || k == "client_recovery_lag") && e.Signal == "client_success" {
				if len(f.Against) < 16 {
					f.Against = append(f.Against, e.ID)
				}
			}
			if (k == "primary_unavailability" || k == "replica_unavailability") && e.Signal == "reachable" && contains(f.Nodes, e.Node) {
				if len(f.Against) < 16 {
					f.Against = append(f.Against, e.ID)
				}
			}
		}
		a.Findings = append(a.Findings, *f)
		// Timeline event is supported by the first cited abnormal observation; it is
		// deliberately not labelled as the injected fault's start time.
		first := findEvidence(in.Evidence, f.Support[0])
		event("first_observed_"+k, first, Evidence{})
		a.Recommendations = append(a.Recommendations, plan(*f))
	}
	for _, f := range peerDisagreements(in) {
		a.Findings = append(a.Findings, f)
		a.Recommendations = append(a.Recommendations, plan(f))
		event("topology_disagreement_observed", findEvidence(in.Evidence, f.Support[0]), Evidence{})
	}
	if len(a.Findings) > 0 {
		a.Detection = "incident"
	} else {
		reachable, healthy := 0, true
		for _, e := range latest {
			if e.Signal == "reachable" {
				b, ok := boolValue(e)
				if !ok || !b {
					healthy = false
				}
				reachable++
			}
			if e.Signal == "slot_coverage" {
				b, ok := boolValue(e)
				if !ok || !b {
					healthy = false
				}
			}
		}
		if reachable >= 2 && healthy {
			a.Detection = "healthy_observed"
			a.Uncertainty = append(a.Uncertainty, "Healthy at the recorded observation points only; gaps and unobserved metrics can hide incidents.")
		}
	}
	for _, e := range latest {
		if e.Signal == "client_success" {
			b, ok := boolValue(e)
			if ok {
				a.Recovery = append(a.Recovery, fmt.Sprintf("Last client probe successful=%t (%s)", b, e.ID))
			}
		}
	}
	validationComplete := byRef["/validation/complete"]
	if b, ok := boolValue(validationComplete); ok && b {
		checked, valid := number(byRef["/validation/checked"])
		expected, validE := number(byRef["/validation/expected"])
		missing, okM := number(byRef["/validation/missing"])
		corrupt, okC := number(byRef["/validation/corrupt"])
		unreadable, okU := number(byRef["/validation/unreadable"])
		unverified, okV := number(byRef["/validation/unverified"])
		if valid && validE && okM && okC && okU && okV && expected > 0 && checked == expected && missing == 0 && corrupt == 0 && unreadable == 0 && unverified == 0 {
			a.Recovery = append(a.Recovery, "Final checked acknowledged-value population preserved; see /validation evidence. This is not a durability guarantee.")
			event("final_correctness_validation_observed", validationComplete, Evidence{})
		}
	} else {
		a.Uncertainty = append(a.Uncertainty, "Final value validation not established; zero default counters do not prove preservation.")
	}
	a.Timeline = append(a.Timeline, observedTransitions(in)...)
	sort.SliceStable(a.Timeline, func(i, j int) bool { return a.Timeline[i].At.Before(a.Timeline[j].At) })
	return a, a.Validate()
}

func contains(s []string, v string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}
func findEvidence(es []Evidence, id string) Evidence {
	for _, e := range es {
		if e.ID == id {
			return e
		}
	}
	return Evidence{}
}
func plan(f Finding) Recommendation {
	r := Recommendation{Category: f.Category, Evidence: f.Support, Objective: "Collect independent evidence before changing system state", Preconditions: []string{"Confirm affected node identity and current topology", "Obtain operator approval before any state-changing action", "Preserve logs and verify available backups"}, Risks: []string{"Stale observations may identify the wrong current owner", "Changing topology can lose data or interrupt availability"}, ExpectedEffects: "Resolve alternative explanations without executing remediation", Verification: f.Verification, HumanRequired: true, Execution: "read_only_plan"}
	switch f.Category {
	case "redundancy_degradation", "replica_unavailability":
		r.Objective = "Assess restoration of a synchronized replica"
		r.Preconditions = append(r.Preconditions, "Confirm an authoritative healthy primary and compatible former-node identity and data")
		r.Verification = append(r.Verification, "Observe replica assignment, link up, synchronization complete and repeated healthy service probes")
	case "data_validation_failure", "validation_incomplete":
		r.Objective = "Preserve evidence and investigate acknowledged-value validation failures"
		r.Risks = append(r.Risks, "Rejoining or overwriting a node may destroy recoverable values")
		r.Verification = append(r.Verification, "Read the frozen acknowledged ledger and compare expected values; retain unreadable and unchecked counts")
	case "acknowledgement_degradation", "acknowledgement_command_failure":
		r.Objective = "Inspect configured acknowledgement thresholds and replica/AOF health"
		r.Risks = append(r.Risks, "Reducing acknowledgement thresholds changes durability expectations")
		r.Verification = append(r.Verification, "Observe successful SET and explicit sufficient WAIT/WAITAOF counts on the same connection")
	case "client_recovery_lag":
		r.Objective = "Inspect client routing refresh, redirect and connection traces"
		r.Verification = append(r.Verification, "Compare direct-node service and client request outcomes without assuming a routing-cache fault")
	}
	return r
}
