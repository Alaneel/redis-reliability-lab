package batch

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

func (s Summary) policyMarkdown(b *bytes.Buffer) {
	fmt.Fprintf(b, "\n## Write acknowledgement policy\n\nExplicit policy counters measured in %d trials, unmeasured in %d. Every successful SET remains in the main validation ledger when the later acknowledgement command is short, errors, or is not attempted. Diagnostic SET/WAIT/WAITAOF stage latency remains in individual reports and does not increase logical request or pacing counts. Errors and unattempted commands supply unknown acknowledgement counts, never fabricated zero replies.\n\n", s.WritePolicy.MeasuredTrials, s.WritePolicy.UnmeasuredTrials)
	var phases []string
	for phase := range s.WritePolicy.ByPhase {
		phases = append(phases, phase)
	}
	sort.Strings(phases)
	p := s.WritePolicy.Totals
	switch s.Config.Workload.WritePolicy {
	case "wait":
		fmt.Fprint(b, "| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |\n|---|---:|---:|---:|---:|---:|---:|---:|\n")
		for _, phase := range phases {
			p := s.WritePolicy.ByPhase[phase]
			fmt.Fprintf(b, "| %s | %d | %d | %d | %d | %d | %d | %d |\n", phase, p.SetAcknowledged, p.SetErrors, p.WaitAttempted, p.WaitSatisfied, p.WaitShort, p.WaitErrors, p.WaitClientTimeouts)
		}
		fmt.Fprintf(b, "\nSET acknowledgements without a subsequent WAIT attempt: %d; client timeouts among those: %d. Attempted WAIT client-call timeouts are counted separately above.\n", p.WaitNotAttempted, p.WaitNotAttemptedTimeouts)
		fmt.Fprint(b, "\n| Observed WAIT replica count | Replies |\n|---:|---:|\n")
		var replies []int64
		for reply := range p.WaitAcknowledgements {
			replies = append(replies, reply)
		}
		sort.Slice(replies, func(i, j int) bool { return replies[i] < replies[j] })
		for _, reply := range replies {
			fmt.Fprintf(b, "| %d | %d |\n", reply, p.WaitAcknowledgements[reply])
		}
		subsetMarkdown(b, "WAIT", s.WritePolicy.WaitSatisfiedSubset, s.WritePolicy.WaitSubsetValidationTrials)
	case "waitaof":
		fmt.Fprintf(b, "WAITAOF requires local AOF acknowledgement 1 and replica AOF acknowledgement %d. Local and replica satisfied counts overlap; they are independently reconciled against paired valid replies.\n\n| Phase | SET acknowledged | SET errors | WAITAOF attempted | Both satisfied | Short | Errors | Client timeouts | Local satisfied | Replica satisfied |\n|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n", s.Config.Workload.WaitReplicas)
		for _, phase := range phases {
			p := s.WritePolicy.ByPhase[phase]
			fmt.Fprintf(b, "| %s | %d | %d | %d | %d | %d | %d | %d | %d | %d |\n", phase, p.SetAcknowledged, p.SetErrors, p.WaitAOFAttempted, p.WaitAOFSatisfied, p.WaitAOFShort, p.WaitAOFErrors, p.WaitAOFClientTimeouts, p.WaitAOFLocalSatisfied, p.WaitAOFReplicaSatisfied)
		}
		fmt.Fprintf(b, "\nSET acknowledgements without a subsequent WAITAOF attempt: %d; client timeouts among those: %d. Attempted WAITAOF client-call timeouts are counted separately above. Total local satisfied: %d; replica satisfied: %d; both satisfied: %d.\n\n| Observed WAITAOF local/replicas pair | Replies |\n|---|---:|\n", p.WaitAOFNotAttempted, p.WaitAOFNotAttemptedTimeouts, p.WaitAOFLocalSatisfied, p.WaitAOFReplicaSatisfied, p.WaitAOFSatisfied)
		var replies []string
		for reply := range p.WaitAOFReplies {
			replies = append(replies, reply)
		}
		sort.Slice(replies, func(i, j int) bool {
			li, ri, _ := parseWaitAOFReply(replies[i])
			lj, rj, _ := parseWaitAOFReply(replies[j])
			return li < lj || li == lj && ri < rj
		})
		for _, reply := range replies {
			fmt.Fprintf(b, "| %s | %d |\n", reply, p.WaitAOFReplies[reply])
		}
		subsetMarkdown(b, "WAITAOF", s.WritePolicy.WaitAOFSatisfiedSubset, s.WritePolicy.WaitAOFSubsetValidationTrials)
	default:
		fmt.Fprint(b, "| Phase | SET acknowledged | SET errors | Policy satisfied |\n|---|---:|---:|---:|\n")
		for _, phase := range phases {
			p := s.WritePolicy.ByPhase[phase]
			fmt.Fprintf(b, "| %s | %d | %d | %d |\n", phase, p.SetAcknowledged, p.SetErrors, p.PolicySatisfied)
		}
	}
}

func subsetMarkdown(b *bytes.Buffer, name string, v workload.ValidationSubset, trials int) {
	fmt.Fprintf(b, "\n%s-satisfied subset validation: expected %d; checked %d; missing %d; corrupt %d; unreadable %d; unverified %d; complete %t. Validation attempted in %d trials. Complete means every subset record was checked; preservation also requires zero missing, corrupt and unreadable records.\n", name, v.Expected, v.Checked, v.Missing, v.Corrupt, v.Unreadable, v.Unverified, v.Complete, trials)
}
