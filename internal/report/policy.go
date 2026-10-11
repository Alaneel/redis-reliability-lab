package report

import (
	"bytes"
	"fmt"
	"sort"
)

var policyPhases = []string{"baseline", "fault", "repair", "reconciliation", "post_recovery"}

func (r Report) policyMarkdown(b *bytes.Buffer) {
	policy := r.WorkloadStats.WritePolicy
	if policy == "" {
		return
	}
	fmt.Fprintf(b, "\nWrite policy: `%s`. Diagnostic stage requests are excluded from logical workload counts.\n", policy)
	logical := "set"
	stages := []string{"set"}
	switch policy {
	case "wait":
		logical, stages = "set_wait", []string{"set", "wait"}
		fmt.Fprintln(b, "Logical write latency includes SET and WAIT on the same physical connection; policy success requires the requested replica-memory acknowledgement count.")
		fmt.Fprint(b, "\n| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |\n|---|---:|---:|---:|---:|---:|---:|---:|\n")
	case "waitaof":
		logical, stages = "set_waitaof", []string{"set", "waitaof"}
		fmt.Fprintln(b, "Logical write latency includes SET and WAITAOF on the same physical connection. WAITAOF policy success requires both requested local and replica AOF acknowledgement counts. These disk acknowledgements are separate from WAIT replica-memory acknowledgements; they do not establish strong consistency or a power-loss guarantee from this process-crash experiment.")
		fmt.Fprint(b, "\n| Phase | SET acknowledged | SET errors | WAITAOF attempted | Both satisfied | Short pair | Command errors | Client timeouts | Local satisfied | Replica satisfied |\n|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	default:
		fmt.Fprint(b, "\n| Phase | SET acknowledged | SET errors |\n|---|---:|---:|\n")
	}
	for _, phase := range policyPhases {
		s, ok := r.WorkloadStats.PolicyByPhase[phase]
		if !ok {
			continue
		}
		switch policy {
		case "wait":
			fmt.Fprintf(b, "| %s | %d | %d | %d | %d | %d | %d | %d |\n", phase, s.SetAcknowledged, s.SetErrors, s.WaitAttempted, s.WaitSatisfied, s.WaitShort, s.WaitErrors, s.WaitClientTimeouts)
		case "waitaof":
			fmt.Fprintf(b, "| %s | %d | %d | %d | %d | %d | %d | %d | %d | %d |\n", phase, s.SetAcknowledged, s.SetErrors, s.WaitAOFAttempted, s.WaitAOFSatisfied, s.WaitAOFShort, s.WaitAOFErrors, s.WaitAOFClientTimeouts, s.WaitAOFLocalSatisfied, s.WaitAOFReplicaSatisfied)
		default:
			fmt.Fprintf(b, "| %s | %d | %d |\n", phase, s.SetAcknowledged, s.SetErrors)
		}
	}
	if policy == "waitaof" {
		fmt.Fprint(b, "\nWAITAOF reply distribution preserves the local/replica pair. Only valid two-integer replies are counted; command errors and unattempted calls have no observed pair. `0/0` is an observed short reply. Local/replica satisfied counts can overlap and must not be added.\n\n| Phase | Local / replica AOF acknowledgements | Replies |\n|---|---|---:|\n")
		for _, phase := range policyPhases {
			replies := r.WorkloadStats.PolicyByPhase[phase].WaitAOFReplies
			pairs := make([]string, 0, len(replies))
			for pair, count := range replies {
				if count > 0 {
					pairs = append(pairs, pair)
				}
			}
			sort.Strings(pairs)
			for _, pair := range pairs {
				fmt.Fprintf(b, "| %s | %s | %d |\n", phase, pair, replies[pair])
			}
		}
	}
	fmt.Fprint(b, "\n| Phase | Measured path | Count | p50 / p95 / p99, ms |\n|---|---|---:|---|\n")
	for _, phase := range policyPhases {
		if s, ok := r.Metrics.Phases[phase].ByOperation[logical]; ok && s.Latency.Count > 0 {
			fmt.Fprintf(b, "| %s | Logical write completion | %d | %.3f / %.3f / %.3f |\n", phase, s.Latency.Count, s.Latency.P50MS, s.Latency.P95MS, s.Latency.P99MS)
		}
		for _, operation := range stages {
			s, ok := r.WorkloadStats.WriteStages.Phases[phase].ByOperation[operation]
			if ok && s.Latency.Count > 0 {
				fmt.Fprintf(b, "| %s | %s | %d | %.3f / %.3f / %.3f |\n", phase, operation, s.Latency.Count, s.Latency.P50MS, s.Latency.P95MS, s.Latency.P99MS)
			}
		}
	}
	fmt.Fprint(b, "\n| Phase | Successful logical writes | Successful writes/s over phase wall time |\n|---|---:|---:|\n")
	for _, phase := range policyPhases {
		p := r.Metrics.Phases[phase]
		if s, ok := p.ByOperation[logical]; ok && p.DurationSeconds > 0 {
			if r.PersistenceFault != nil && s.Attempts == 0 {
				continue
			}
			fmt.Fprintf(b, "| %s | %d | %.3f |\n", phase, s.Successes, float64(s.Successes)/p.DurationSeconds)
		}
	}
	switch policy {
	case "wait":
		fmt.Fprintf(b, "\nSET acknowledgements without a subsequent WAIT attempt: %d; client timeouts among those: %d. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.\n", r.WorkloadStats.PolicyTotals.WaitNotAttempted, r.WorkloadStats.PolicyTotals.WaitNotAttemptedTimeouts)
	case "waitaof":
		fmt.Fprintf(b, "\nSET acknowledgements without a subsequent WAITAOF attempt: %d; client timeouts among those: %d. WAITAOF command errors and WAITAOF client timeouts above count attempted WAITAOF client calls only. A short integer pair is separate from a client timeout.\n", r.WorkloadStats.PolicyTotals.WaitAOFNotAttempted, r.WorkloadStats.PolicyTotals.WaitAOFNotAttemptedTimeouts)
	}
	if r.Recovery.WritePolicyRecovered != nil {
		fmt.Fprintf(b, "\nConfigured write policy recovered for the stable interval: %t. This criterion is independent of ordinary SET/GET service recovery.\n", *r.Recovery.WritePolicyRecovered)
	}
}

func (r Report) persistencePreparationMarkdown(b *bytes.Buffer) {
	if len(r.PersistencePrepared) == 0 && r.BaselineWritePolicyProbe == nil && r.WorkloadStats.WritePolicy != "waitaof" {
		return
	}
	fmt.Fprintln(b, "\n### Before-baseline persistence and command evidence")
	prepared := r.PersistencePrepared
	if len(prepared) == 0 && r.PersistenceFault != nil {
		prepared = r.PersistenceFault.Prepared
	}
	if len(prepared) == 0 {
		fmt.Fprintln(b, "\nAdministrative persistence runtime preparation: unobserved.")
	} else {
		fmt.Fprintf(b, "\nPersistence preparation contains %d node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.\n", len(prepared))
		fmt.Fprint(b, "\n| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |\n|---|---|---|---|---|---|\n")
		for _, sample := range prepared {
			pong := "unobserved"
			if !sample.Evidence.ObservedAt.IsZero() {
				pong = fmt.Sprint(sample.Evidence.RedisPONG)
			}
			fmt.Fprintf(b, "| %s | %s | %s | %s | %s | %s |\n", sample.Node.ID, sample.OriginalRedisID, optionalTime(&sample.Evidence.ObservedAt), sample.Evidence.RedisClusterID, pong, sample.Error)
		}
	}
	p := r.BaselineWritePolicyProbe
	if p == nil && r.WorkloadStats.WritePolicy != "waitaof" {
		return
	}
	if p == nil || p.At.IsZero() {
		fmt.Fprintln(b, "\nAll-shard before-baseline acknowledgement-policy probe: unobserved.")
		return
	}
	policy := p.Policy
	if policy == "" {
		policy = r.WorkloadStats.WritePolicy
	}
	fmt.Fprintf(b, "\nBefore-baseline `%s` policy probe at %s; all shards successful: %t.\n", policy, optionalTime(&p.At), p.Successful)
	switch policy {
	case "wait":
		fmt.Fprintf(b, "Requested replica-memory count: %d.\n", p.RequestedReplicas)
		fmt.Fprint(b, "\n| Original shard index | SET acknowledged | WAIT attempted | Replica-memory reply |\n|---|---|---|---|\n")
	case "waitaof":
		fmt.Fprintf(b, "requested local/replica AOF counts: %d/%d.\n", p.WaitAOFRequestedLocal, p.WaitAOFRequestedReplicas)
		fmt.Fprint(b, "\n| Original shard index | SET acknowledged | WAITAOF attempted | Local / replica AOF reply |\n|---|---|---|---|\n")
	default:
		fmt.Fprint(b, "\n| Original shard index | SET acknowledged |\n|---|---|\n")
	}
	n := max(len(p.SetAcknowledged), len(p.WaitAttempted), len(p.Acknowledgements), len(p.AcknowledgementObserved), len(p.WaitAOFAttempted), len(p.WaitAOFLocalAcknowledgements), len(p.WaitAOFReplicaAcknowledgements), len(p.WaitAOFAcknowledgementObserved))
	for i := 0; i < n; i++ {
		set, attempted, pair := "unobserved", "unobserved", "unobserved"
		if i < len(p.SetAcknowledged) {
			set = fmt.Sprint(p.SetAcknowledged[i])
		}
		if policy == "wait" {
			if i < len(p.WaitAttempted) {
				attempted = fmt.Sprint(p.WaitAttempted[i])
			}
			if i < len(p.AcknowledgementObserved) && p.AcknowledgementObserved[i] && i < len(p.WaitAttempted) && p.WaitAttempted[i] && i < len(p.Acknowledgements) && p.Acknowledgements[i] >= 0 {
				pair = fmt.Sprint(p.Acknowledgements[i])
			}
		}
		if policy == "waitaof" && i < len(p.WaitAOFAttempted) {
			attempted = fmt.Sprint(p.WaitAOFAttempted[i])
		}
		if policy == "waitaof" && i < len(p.WaitAOFAcknowledgementObserved) && p.WaitAOFAcknowledgementObserved[i] && i < len(p.WaitAOFAttempted) && p.WaitAOFAttempted[i] && i < len(p.WaitAOFLocalAcknowledgements) && i < len(p.WaitAOFReplicaAcknowledgements) {
			local, replica := p.WaitAOFLocalAcknowledgements[i], p.WaitAOFReplicaAcknowledgements[i]
			if local >= 0 && local <= 1 && replica >= 0 {
				pair = fmt.Sprintf("%d/%d", local, replica)
			}
		}
		if policy == "wait" || policy == "waitaof" {
			fmt.Fprintf(b, "| %d | %s | %s | %s |\n", i, set, attempted, pair)
		} else {
			fmt.Fprintf(b, "| %d | %s |\n", i, set)
		}
	}
	for _, err := range p.Errors {
		fmt.Fprintf(b, "\nBefore-baseline policy probe error: %s\n", err)
	}
}

func (r Report) policyValidationMarkdown(b *bytes.Buffer) {
	policy := r.WorkloadStats.WritePolicy
	if policy != "wait" && policy != "waitaof" {
		return
	}
	s := r.Validation.WaitSatisfiedSubset
	command := "WAIT"
	if policy == "waitaof" {
		s, command = r.Validation.WaitAOFSatisfiedSubset, "WAITAOF"
	}
	fmt.Fprintf(b, "\n%s-satisfied subset: expected %d; checked %d; missing %d; corrupt %d; unreadable %d; unverified %d; complete %t. The main ledger also includes every successful SET followed by incomplete %s.\n", command, s.Expected, s.Checked, s.Missing, s.Corrupt, s.Unreadable, s.Unverified, s.Complete, command)
}
