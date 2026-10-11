package batch

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/alanwang/redis-reliability-lab/internal/metrics"
	"github.com/alanwang/redis-reliability-lab/internal/report"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

func sumPolicyCounts(a, b workload.PolicyCounts) workload.PolicyCounts {
	a.SetAcknowledged += b.SetAcknowledged
	a.SetErrors += b.SetErrors
	a.WaitAttempted += b.WaitAttempted
	a.WaitSatisfied += b.WaitSatisfied
	a.WaitShort += b.WaitShort
	a.WaitErrors += b.WaitErrors
	a.WaitClientTimeouts += b.WaitClientTimeouts
	a.WaitNotAttempted += b.WaitNotAttempted
	a.WaitNotAttemptedTimeouts += b.WaitNotAttemptedTimeouts
	a.WaitAOFAttempted += b.WaitAOFAttempted
	a.WaitAOFSatisfied += b.WaitAOFSatisfied
	a.WaitAOFShort += b.WaitAOFShort
	a.WaitAOFErrors += b.WaitAOFErrors
	a.WaitAOFClientTimeouts += b.WaitAOFClientTimeouts
	a.WaitAOFNotAttempted += b.WaitAOFNotAttempted
	a.WaitAOFNotAttemptedTimeouts += b.WaitAOFNotAttemptedTimeouts
	a.WaitAOFLocalSatisfied += b.WaitAOFLocalSatisfied
	a.WaitAOFReplicaSatisfied += b.WaitAOFReplicaSatisfied
	a.PolicySatisfied += b.PolicySatisfied
	a.HandledRedirects += b.HandledRedirects
	wait := make(map[int64]uint64, len(a.WaitAcknowledgements)+len(b.WaitAcknowledgements))
	for count, n := range a.WaitAcknowledgements {
		wait[count] = n
	}
	for count, n := range b.WaitAcknowledgements {
		wait[count] += n
	}
	a.WaitAcknowledgements = wait
	waitaof := make(map[string]uint64, len(a.WaitAOFReplies)+len(b.WaitAOFReplies))
	for reply, n := range a.WaitAOFReplies {
		waitaof[reply] += n
	}
	for reply, n := range b.WaitAOFReplies {
		waitaof[reply] += n
	}
	a.WaitAOFReplies = waitaof
	return a
}

func validPolicyNumbers(r *report.Report) error {
	w := r.WorkloadStats
	if w.WritePolicy == "" {
		measuredTraffic := w.WritesAttempted != 0 || w.AcknowledgedWrites != 0 || w.AmbiguousWrites != 0 || w.StartedRequests != 0 || w.PacingMethodology != ""
		if r.Config.Workload.WritePolicy == "waitaof" && (r.Outcome == "recovered" || measuredTraffic) || r.Config.Workload.WritePolicy == "wait" && r.Outcome == "recovered" {
			return fmt.Errorf("%s trial lacks policy evidence", strings.ToUpper(r.Config.Workload.WritePolicy))
		}
		return nil // Legacy SET reports and failures before traffic are unmeasured.
	}
	p := w.PolicyTotals
	if w.WritePolicy != r.Config.Workload.WritePolicy {
		return errors.New("workload write policy differs from effective configuration")
	}
	if p.SetAcknowledged != w.AcknowledgedWrites || p.SetAcknowledged > w.WritesAttempted || p.SetErrors != w.WritesAttempted-p.SetAcknowledged || p.SetErrors != w.AmbiguousWrites {
		return errors.New("SET acknowledgement policy counters do not reconcile")
	}
	if err := validPolicyCounts(p, w.WritePolicy, r.Config.Workload.WaitReplicas); err != nil {
		return err
	}
	var summed workload.PolicyCounts
	for phase, counts := range w.PolicyByPhase {
		if err := validPolicyCounts(counts, w.WritePolicy, r.Config.Workload.WaitReplicas); err != nil {
			return fmt.Errorf("phase %s: %w", phase, err)
		}
		summed = sumPolicyCounts(summed, counts)
		if !validStageAttempts(w.WriteStages.Phases[phase], counts.SetAcknowledged+counts.SetErrors, counts.WaitAttempted, counts.WaitAOFAttempted) {
			return fmt.Errorf("phase %s diagnostic SET/WAIT/WAITAOF stage attempts differ from policy counters", phase)
		}
	}
	if !equalPolicyCounts(p, summed) {
		return errors.New("summed phase policy counters or reply histogram differ from policy totals")
	}
	stages := w.WriteStages.Totals
	if !validStageAttempts(stages, w.WritesAttempted, p.WaitAttempted, p.WaitAOFAttempted) {
		return errors.New("diagnostic SET/WAIT/WAITAOF stage attempts differ from policy totals")
	}
	var stageSet, stageWait, stageWaitAOF, stageAttempts uint64
	for _, phase := range w.WriteStages.Phases {
		stageSet += phase.ByOperation["set"].Attempts
		stageWait += phase.ByOperation["wait"].Attempts
		stageWaitAOF += phase.ByOperation["waitaof"].Attempts
		stageAttempts += phase.Attempts
	}
	if stageSet != w.WritesAttempted || stageWait != p.WaitAttempted || stageWaitAOF != p.WaitAOFAttempted || stageAttempts != stages.Attempts {
		return errors.New("summed diagnostic stage phases differ from their totals")
	}
	if !validLogicalWriteCounts(r.Metrics.Totals, p, w.WritePolicy) {
		return errors.New("logical write totals differ from SET attempts or policy success")
	}
	for phase, counts := range w.PolicyByPhase {
		if !validLogicalWriteCounts(r.Metrics.Phases[phase], counts, w.WritePolicy) {
			return fmt.Errorf("phase %s logical write counts differ from SET attempts or policy success", phase)
		}
	}
	var logicalAttempts, logicalSuccesses uint64
	for phase, summary := range r.Metrics.Phases {
		if !validLogicalWriteCounts(summary, w.PolicyByPhase[phase], w.WritePolicy) {
			return fmt.Errorf("phase %s logical write counts differ from SET attempts or policy success", phase)
		}
		writes := summary.ByOperation[logicalWriteName(w.WritePolicy)]
		logicalAttempts += writes.Attempts
		logicalSuccesses += writes.Successes
	}
	if logicalAttempts != w.WritesAttempted || logicalSuccesses != p.PolicySatisfied {
		return errors.New("summed phase logical writes differ from policy totals")
	}
	if w.WritePolicy != "wait" && subsetHasRecords(r.Validation.WaitSatisfiedSubset) || w.WritePolicy != "waitaof" && subsetHasRecords(r.Validation.WaitAOFSatisfiedSubset) {
		return errors.New("validation includes records for a different acknowledgement policy")
	}
	var v workload.ValidationSubset
	var satisfied uint64
	switch w.WritePolicy {
	case "wait":
		v, satisfied = r.Validation.WaitSatisfiedSubset, p.WaitSatisfied
	case "waitaof":
		v, satisfied = r.Validation.WaitAOFSatisfiedSubset, p.WaitAOFSatisfied
	default:
		return nil
	}
	label := strings.ToUpper(w.WritePolicy)
	if !r.Validation.StartedAt.IsZero() {
		if v.Expected < 0 || uint64(v.Expected) != satisfied || v.Checked < 0 || v.Missing < 0 || v.Corrupt < 0 || v.Unreadable < 0 || v.Unverified < 0 || v.Checked > v.Expected || v.Unverified != v.Expected-v.Checked || v.Missing > v.Checked || v.Corrupt > v.Checked-v.Missing || v.Unreadable > v.Checked-v.Missing-v.Corrupt || v.Complete != (v.Unverified == 0) {
			return fmt.Errorf("%s-satisfied subset validation does not reconcile", label)
		}
		if v.Expected > r.Validation.Expected || v.Checked > r.Validation.Checked || v.Missing > r.Validation.Missing || v.Corrupt > r.Validation.Corrupt || v.Unreadable > r.Validation.Unreadable || v.Unverified > r.Validation.Unverified {
			return fmt.Errorf("%s-satisfied subset validation exceeds the full SET ledger", label)
		}
	}
	if r.Outcome == "recovered" && (r.Recovery.WritePolicyRecovered == nil || !*r.Recovery.WritePolicyRecovered || !observedMilestone(r.Recovery.WritePolicyStable) || r.Recovery.WritePolicyStable.At.After(r.FinishedAt) || r.Validation.StartedAt.IsZero() || r.Validation.CompletedAt.IsZero() || !v.Complete || v.Missing != 0 || v.Corrupt != 0 || v.Unreadable != 0 || v.Unverified != 0) {
		return fmt.Errorf("recovered %s trial lacks sustained acknowledgement policy and subset validation", label)
	}
	return nil
}

func logicalWriteName(policy string) string {
	switch policy {
	case "wait":
		return "set_wait"
	case "waitaof":
		return "set_waitaof"
	default:
		return "set"
	}
}

func validLogicalWriteCounts(summary metrics.Summary, p workload.PolicyCounts, policy string) bool {
	name := logicalWriteName(policy)
	writes := summary.ByOperation[name]
	attempts := p.SetAcknowledged + p.SetErrors
	if p.PolicySatisfied > attempts || writes.Attempts != attempts || writes.Successes != p.PolicySatisfied || writes.Errors != attempts-p.PolicySatisfied {
		return false
	}
	for _, other := range []string{"set", "set_wait", "set_waitaof", "wait", "waitaof"} {
		op := summary.ByOperation[other]
		if other != name && (op.Attempts != 0 || op.Successes != 0 || op.Errors != 0) {
			return false
		}
	}
	var measuredAttempts, measuredSuccesses, measuredErrors uint64
	for operation, counts := range summary.ByOperation {
		if counts.Successes > counts.Attempts || counts.Errors != counts.Attempts-counts.Successes || counts.Attempts > summary.Attempts-measuredAttempts || counts.Successes > summary.Successes-measuredSuccesses || counts.Errors > summary.Errors-measuredErrors {
			return false
		}
		if operation != name && operation != "get" && counts.Attempts != 0 {
			return false
		}
		measuredAttempts += counts.Attempts
		measuredSuccesses += counts.Successes
		measuredErrors += counts.Errors
	}
	if measuredAttempts != summary.Attempts || measuredSuccesses != summary.Successes || measuredErrors != summary.Errors {
		return false
	}
	return true
}

func validStageAttempts(stages metrics.Summary, set, wait, waitaof uint64) bool {
	if stages.ByOperation["set"].Attempts != set || stages.ByOperation["wait"].Attempts != wait || stages.ByOperation["waitaof"].Attempts != waitaof || stages.Attempts != set+wait+waitaof {
		return false
	}
	for name, operation := range stages.ByOperation {
		if name != "set" && name != "wait" && name != "waitaof" && operation.Attempts != 0 {
			return false
		}
	}
	return true
}

func subsetHasRecords(v workload.ValidationSubset) bool {
	return v.Expected != 0 || v.Checked != 0 || v.Missing != 0 || v.Corrupt != 0 || v.Unreadable != 0 || v.Unverified != 0
}

func waitHasCounts(p workload.PolicyCounts) bool {
	return p.WaitAttempted != 0 || p.WaitSatisfied != 0 || p.WaitShort != 0 || p.WaitErrors != 0 || p.WaitClientTimeouts != 0 || p.WaitNotAttempted != 0 || p.WaitNotAttemptedTimeouts != 0 || len(p.WaitAcknowledgements) != 0
}

func waitAOFHasCounts(p workload.PolicyCounts) bool {
	return p.WaitAOFAttempted != 0 || p.WaitAOFSatisfied != 0 || p.WaitAOFShort != 0 || p.WaitAOFErrors != 0 || p.WaitAOFClientTimeouts != 0 || p.WaitAOFNotAttempted != 0 || p.WaitAOFNotAttemptedTimeouts != 0 || p.WaitAOFLocalSatisfied != 0 || p.WaitAOFReplicaSatisfied != 0 || len(p.WaitAOFReplies) != 0
}

func validPolicyCounts(p workload.PolicyCounts, policy string, requested int) error {
	if policy != "wait" && waitHasCounts(p) || policy != "waitaof" && waitAOFHasCounts(p) {
		return errors.New("write policy includes commands or counters for a different acknowledgement policy")
	}
	switch policy {
	case "set":
		if p.PolicySatisfied != p.SetAcknowledged {
			return errors.New("SET policy includes incorrect policy success")
		}
	case "wait":
		if p.WaitSatisfied > p.WaitAttempted || p.WaitShort > p.WaitAttempted-p.WaitSatisfied || p.WaitErrors != p.WaitAttempted-p.WaitSatisfied-p.WaitShort || p.WaitClientTimeouts > p.WaitErrors || p.WaitNotAttemptedTimeouts > p.WaitNotAttempted {
			return errors.New("SET/WAIT policy counters do not reconcile")
		}
		if requested <= 0 || p.WaitAttempted > p.SetAcknowledged || p.WaitNotAttempted != p.SetAcknowledged-p.WaitAttempted || p.PolicySatisfied != p.WaitSatisfied {
			return errors.New("WAIT policy does not account for every successful SET and policy outcome")
		}
		var satisfied, short uint64
		for acknowledgements, count := range p.WaitAcknowledgements {
			if acknowledgements < 0 {
				return errors.New("WAIT acknowledgement histogram contains a negative reply")
			}
			if count == 0 || count > p.WaitAttempted-satisfied-short {
				return errors.New("WAIT reply histogram has zero-frequency or excess observations")
			}
			if acknowledgements >= int64(requested) {
				satisfied += count
			} else {
				short += count
			}
		}
		if satisfied != p.WaitSatisfied || short != p.WaitShort {
			return errors.New("WAIT reply histogram does not reconcile with full and short replies")
		}
	case "waitaof":
		if p.WaitAOFSatisfied > p.WaitAOFAttempted || p.WaitAOFShort > p.WaitAOFAttempted-p.WaitAOFSatisfied || p.WaitAOFErrors != p.WaitAOFAttempted-p.WaitAOFSatisfied-p.WaitAOFShort || p.WaitAOFClientTimeouts > p.WaitAOFErrors || p.WaitAOFNotAttemptedTimeouts > p.WaitAOFNotAttempted || p.WaitAOFLocalSatisfied > p.WaitAOFAttempted || p.WaitAOFReplicaSatisfied > p.WaitAOFAttempted {
			return errors.New("SET/WAITAOF policy counters do not reconcile")
		}
		if requested <= 0 || p.WaitAOFAttempted > p.SetAcknowledged || p.WaitAOFNotAttempted != p.SetAcknowledged-p.WaitAOFAttempted || p.PolicySatisfied != p.WaitAOFSatisfied {
			return errors.New("WAITAOF policy does not account for every successful SET and policy outcome")
		}
		var satisfied, short, localSatisfied, replicaSatisfied uint64
		for reply, count := range p.WaitAOFReplies {
			local, replicas, err := parseWaitAOFReply(reply)
			if err != nil {
				return err
			}
			if count == 0 || count > p.WaitAOFAttempted-satisfied-short {
				return errors.New("WAITAOF reply histogram has zero-frequency or excess observations")
			}
			localOK, replicaOK := local == 1, replicas >= int64(requested)
			if localOK {
				localSatisfied += count
			}
			if replicaOK {
				replicaSatisfied += count
			}
			if localOK && replicaOK {
				satisfied += count
			} else {
				short += count
			}
		}
		if satisfied != p.WaitAOFSatisfied || short != p.WaitAOFShort || localSatisfied != p.WaitAOFLocalSatisfied || replicaSatisfied != p.WaitAOFReplicaSatisfied {
			return errors.New("WAITAOF paired reply histogram does not reconcile with both, short, local and replica satisfaction")
		}
	default:
		return errors.New("unsupported measured write policy")
	}
	return nil
}

func parseWaitAOFReply(reply string) (int64, int64, error) {
	localText, replicaText, ok := strings.Cut(reply, "/")
	local, localErr := strconv.ParseInt(localText, 10, 64)
	replicas, replicaErr := strconv.ParseInt(replicaText, 10, 64)
	if !ok || localErr != nil || replicaErr != nil || local < 0 || local > 1 || replicas < 0 || fmt.Sprintf("%d/%d", local, replicas) != reply {
		return 0, 0, errors.New("WAITAOF reply histogram contains a malformed or noncanonical local/replicas pair")
	}
	return local, replicas, nil
}

func equalPolicyCounts(a, b workload.PolicyCounts) bool {
	if a.WaitAcknowledgements == nil {
		a.WaitAcknowledgements = map[int64]uint64{}
	}
	if b.WaitAcknowledgements == nil {
		b.WaitAcknowledgements = map[int64]uint64{}
	}
	if a.WaitAOFReplies == nil {
		a.WaitAOFReplies = map[string]uint64{}
	}
	if b.WaitAOFReplies == nil {
		b.WaitAOFReplies = map[string]uint64{}
	}
	return reflect.DeepEqual(a, b)
}
