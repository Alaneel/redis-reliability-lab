package workload

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/metrics"
)

// WriteExecutor owns routing and physical-connection affinity for SET and its
// optional WAIT/WAITAOF command. Later errors cannot erase a SET acknowledgement.
type WriteExecutor interface {
	Write(context.Context, string, string) WriteResult
}

type WriteResult struct {
	SetAcknowledged                bool
	SetAcknowledgedAt              time.Time
	SetError                       error
	WaitAttempted                  bool
	WaitRequested                  int
	WaitAcknowledgements           int64
	WaitError                      error
	WaitAOFAttempted               bool
	WaitAOFRequestedLocal          int
	WaitAOFRequestedReplicas       int
	WaitAOFLocalAcknowledgements   int64
	WaitAOFReplicaAcknowledgements int64
	WaitAOFError                   error
	PolicySatisfied                bool
	Redirects                      int
	SetLatency                     time.Duration
	WaitLatency                    time.Duration
	WaitAOFLatency                 time.Duration
	TotalLatency                   time.Duration
}

// PolicyCounts keeps WAIT memory and WAITAOF disk acknowledgements separate.
// Command errors/client timeouts count attempted calls only; not-attempted
// timeout counts are subsets of not-attempted calls. Reply histograms exclude
// errors and invalid replies. WAITAOF component successes overlap; both requested
// counts must be satisfied for full policy success.
type PolicyCounts struct {
	SetAcknowledged             uint64            `json:"set_acknowledged"`
	SetErrors                   uint64            `json:"set_errors"`
	WaitAttempted               uint64            `json:"wait_attempted"`
	WaitNotAttempted            uint64            `json:"wait_not_attempted"`
	WaitNotAttemptedTimeouts    uint64            `json:"wait_not_attempted_timeouts"`
	WaitSatisfied               uint64            `json:"wait_satisfied"`
	WaitShort                   uint64            `json:"wait_short"`
	WaitErrors                  uint64            `json:"wait_errors"`
	WaitClientTimeouts          uint64            `json:"wait_client_timeouts"`
	PolicySatisfied             uint64            `json:"policy_satisfied"`
	HandledRedirects            uint64            `json:"handled_redirects"`
	WaitAcknowledgements        map[int64]uint64  `json:"wait_acknowledgements"`
	WaitAOFAttempted            uint64            `json:"waitaof_attempted"`
	WaitAOFNotAttempted         uint64            `json:"waitaof_not_attempted"`
	WaitAOFNotAttemptedTimeouts uint64            `json:"waitaof_not_attempted_timeouts"`
	WaitAOFSatisfied            uint64            `json:"waitaof_satisfied"`
	WaitAOFShort                uint64            `json:"waitaof_short"`
	WaitAOFErrors               uint64            `json:"waitaof_errors"`
	WaitAOFClientTimeouts       uint64            `json:"waitaof_client_timeouts"`
	WaitAOFLocalSatisfied       uint64            `json:"waitaof_local_satisfied"`
	WaitAOFReplicaSatisfied     uint64            `json:"waitaof_replica_satisfied"`
	WaitAOFReplies              map[string]uint64 `json:"waitaof_replies"`
}

type ValidationSubset struct {
	Expected   int  `json:"expected"`
	Checked    int  `json:"checked"`
	Missing    int  `json:"missing"`
	Corrupt    int  `json:"corrupt"`
	Unreadable int  `json:"unreadable"`
	Unverified int  `json:"unverified"`
	Complete   bool `json:"complete"`
}

const WriteStagesMethodology = "Diagnostic SET, WAIT and WAITAOF stages are excluded from logical request and pacing counts. SET latency includes routing, pool/connection initialization and handled redirections before its terminal result; WAIT/WAITAOF latency covers only that attempted command on the same physical connection. Logical set_wait/set_waitaof latency covers the complete write executor operation. Short WAIT or WAITAOF replies have distinct wait_short/waitaof_short classes, separate from client timeout. Unissued commands have no stage latency. Reply histograms exclude errors and malformed replies; WAITAOF histogram keys are canonical local/replicas integer pairs. WAIT replica-memory acknowledgements and WAITAOF local/replica AOF acknowledgements remain separate."

var (
	errSetUnacknowledged   = errors.New("write executor returned no SET acknowledgement or error")
	errWaitNotAttempted    = errors.New("WAIT policy was not attempted after SET acknowledgement")
	errWaitAOFNotAttempted = errors.New("WAITAOF policy was not attempted after SET acknowledgement")
	errWaitAOFInvalidReply = errors.New("WAITAOF returned an invalid local or replica acknowledgement count")
	errPolicyUnsatisfied   = errors.New("requested write policy was not satisfied")
)

type normalWriter struct{ client Client }

func (w normalWriter) Write(ctx context.Context, key, value string) WriteResult {
	started := time.Now()
	err := w.client.Set(ctx, key, value, 0).Err()
	completed := time.Now()
	elapsed := completed.Sub(started)
	result := WriteResult{SetAcknowledged: err == nil, SetError: err, PolicySatisfied: err == nil, SetLatency: elapsed, TotalLatency: elapsed}
	if err == nil {
		result.SetAcknowledgedAt = completed
	}
	return result
}

func newPolicyCounts() PolicyCounts {
	return PolicyCounts{WaitAcknowledgements: make(map[int64]uint64), WaitAOFReplies: make(map[string]uint64)}
}

func clonePolicyCounts(counts PolicyCounts) PolicyCounts {
	out := counts
	out.WaitAcknowledgements = make(map[int64]uint64, len(counts.WaitAcknowledgements))
	for acknowledgements, count := range counts.WaitAcknowledgements {
		out.WaitAcknowledgements[acknowledgements] = count
	}
	out.WaitAOFReplies = make(map[string]uint64, len(counts.WaitAOFReplies))
	for reply, count := range counts.WaitAOFReplies {
		out.WaitAOFReplies[reply] = count
	}
	return out
}

func logicalWriteError(policy string, result WriteResult) error {
	if result.SetError != nil {
		return result.SetError
	}
	if !result.SetAcknowledged {
		return errSetUnacknowledged
	}
	if policy == "wait" {
		if result.WaitError != nil {
			return result.WaitError
		}
		if !result.WaitAttempted {
			return errWaitNotAttempted
		}
		if result.WaitRequested <= 0 || result.WaitAcknowledgements < int64(result.WaitRequested) {
			return metrics.ErrWaitShort
		}
	}
	if policy == "waitaof" {
		if result.WaitAOFError != nil {
			return result.WaitAOFError
		}
		if !result.WaitAOFAttempted {
			return errWaitAOFNotAttempted
		}
		if !validWaitAOFReply(result.WaitAOFLocalAcknowledgements, result.WaitAOFReplicaAcknowledgements) {
			return errWaitAOFInvalidReply
		}
		if !waitAOFCountsSatisfied(result.WaitAOFRequestedLocal, result.WaitAOFRequestedReplicas, result.WaitAOFLocalAcknowledgements, result.WaitAOFReplicaAcknowledgements) {
			return metrics.ErrWaitAOFShort
		}
	}
	if !result.PolicySatisfied {
		return errPolicyUnsatisfied
	}
	return nil
}

func waitSatisfiedWrite(write Write) bool {
	return write.WaitAttempted && write.WaitRequested > 0 && write.WaitAcknowledgements >= int64(write.WaitRequested) && write.PolicySatisfied
}

func validWaitAOFReply(local, replicas int64) bool { return local >= 0 && local <= 1 && replicas >= 0 }

func waitAOFCountsSatisfied(requestLocal, requestReplicas int, local, replicas int64) bool {
	return requestLocal > 0 && requestReplicas > 0 && validWaitAOFReply(local, replicas) && local >= int64(requestLocal) && replicas >= int64(requestReplicas)
}

func waitAOFSatisfiedWrite(write Write) bool {
	return write.WaitAOFAttempted && write.PolicySatisfied && waitAOFCountsSatisfied(write.WaitAOFRequestedLocal, write.WaitAOFRequestedReplicas, write.WaitAOFLocalAcknowledgements, write.WaitAOFReplicaAcknowledgements)
}

func (r *Runner) recordWriteResult(phase string, result WriteResult) {
	logicalErr := logicalWriteError(r.config.WritePolicy, result)
	setErr := result.SetError
	if !result.SetAcknowledged && setErr == nil {
		setErr = errSetUnacknowledged
	}
	r.stages.Record(phase, "set", result.SetLatency, setErr)
	if result.WaitAttempted {
		waitErr := result.WaitError
		if waitErr == nil && (result.WaitRequested <= 0 || result.WaitAcknowledgements < int64(result.WaitRequested)) {
			waitErr = metrics.ErrWaitShort
		}
		r.stages.Record(phase, "wait", result.WaitLatency, waitErr)
	}
	if result.WaitAOFAttempted {
		waitAOFErr := result.WaitAOFError
		if waitAOFErr == nil && !validWaitAOFReply(result.WaitAOFLocalAcknowledgements, result.WaitAOFReplicaAcknowledgements) {
			waitAOFErr = errWaitAOFInvalidReply
		}
		if waitAOFErr == nil && !waitAOFCountsSatisfied(result.WaitAOFRequestedLocal, result.WaitAOFRequestedReplicas, result.WaitAOFLocalAcknowledgements, result.WaitAOFReplicaAcknowledgements) {
			waitAOFErr = metrics.ErrWaitAOFShort
		}
		r.stages.Record(phase, "waitaof", result.WaitAOFLatency, waitAOFErr)
	}
	r.policyMu.Lock()
	defer r.policyMu.Unlock()
	phaseCounts, present := r.policyByPhase[phase]
	if !present {
		phaseCounts = newPolicyCounts()
	}
	for _, counts := range []*PolicyCounts{&r.policyTotals, &phaseCounts} {
		if result.SetAcknowledged {
			counts.SetAcknowledged++
		} else {
			counts.SetErrors++
		}
		if logicalErr == nil {
			counts.PolicySatisfied++
		}
		if result.Redirects > 0 {
			counts.HandledRedirects += uint64(result.Redirects)
		}
		if r.config.WritePolicy == "wait" && result.SetAcknowledged && !result.WaitAttempted {
			counts.WaitNotAttempted++
			if metrics.ClassifyError(result.WaitError) == "timeout" {
				counts.WaitNotAttemptedTimeouts++
			}
		}
		if result.WaitAttempted {
			counts.WaitAttempted++
			if result.WaitError != nil {
				counts.WaitErrors++
				if metrics.ClassifyError(result.WaitError) == "timeout" {
					counts.WaitClientTimeouts++
				}
			} else {
				counts.WaitAcknowledgements[result.WaitAcknowledgements]++
				if result.WaitRequested > 0 && result.WaitAcknowledgements >= int64(result.WaitRequested) {
					counts.WaitSatisfied++
				} else {
					counts.WaitShort++
				}
			}
		}
		if r.config.WritePolicy == "waitaof" && result.SetAcknowledged && !result.WaitAOFAttempted {
			counts.WaitAOFNotAttempted++
			if metrics.ClassifyError(result.WaitAOFError) == "timeout" {
				counts.WaitAOFNotAttemptedTimeouts++
			}
		}
		if result.WaitAOFAttempted {
			counts.WaitAOFAttempted++
			if result.WaitAOFError != nil || !validWaitAOFReply(result.WaitAOFLocalAcknowledgements, result.WaitAOFReplicaAcknowledgements) {
				counts.WaitAOFErrors++
				if metrics.ClassifyError(result.WaitAOFError) == "timeout" {
					counts.WaitAOFClientTimeouts++
				}
			} else {
				counts.WaitAOFReplies[fmt.Sprintf("%d/%d", result.WaitAOFLocalAcknowledgements, result.WaitAOFReplicaAcknowledgements)]++
				localOK := result.WaitAOFRequestedLocal > 0 && result.WaitAOFLocalAcknowledgements >= int64(result.WaitAOFRequestedLocal)
				replicaOK := result.WaitAOFRequestedReplicas > 0 && result.WaitAOFReplicaAcknowledgements >= int64(result.WaitAOFRequestedReplicas)
				if localOK {
					counts.WaitAOFLocalSatisfied++
				}
				if replicaOK {
					counts.WaitAOFReplicaSatisfied++
				}
				if localOK && replicaOK {
					counts.WaitAOFSatisfied++
				} else {
					counts.WaitAOFShort++
				}
			}
		}
	}
	r.policyByPhase[phase] = phaseCounts
}
