package experiment

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/report"
	"github.com/Alaneel/redis-reliability-lab/internal/workload"
)

// ProbeWritePolicy preserves the WAIT probe API for existing callers.
func ProbeWritePolicy(ctx context.Context, writer workload.WriteExecutor, reader workload.Client, keys []string, runID string, replicas int, timeout time.Duration) report.PolicyProbe {
	return ProbeAcknowledgementPolicy(ctx, writer, reader, keys, runID, "wait", replicas, timeout)
}

func requiresAcknowledgementPolicy(policy string) bool {
	return policy == "wait" || policy == "waitaof"
}

// ProbeAcknowledgementPolicy tests the configured acknowledgement on each initial
// shard independently of ordinary SET/GET service recovery and workload metrics.
// A valid zero reply is an observed short result; errors never fabricate a pair.
func ProbeAcknowledgementPolicy(ctx context.Context, writer workload.WriteExecutor, reader workload.Client, keys []string, runID, policy string, replicas int, timeout time.Duration) report.PolicyProbe {
	p := report.PolicyProbe{Policy: policy, SetAcknowledged: make([]bool, len(keys))}
	if policy != "set" && !requiresAcknowledgementPolicy(policy) || (requiresAcknowledgementPolicy(policy) && replicas <= 0) || timeout <= 0 || writer == nil || reader == nil {
		p.At = time.Now()
		p.Errors = []string{"policy probe requires SET, WAIT or WAITAOF, positive requested replicas for acknowledgement commands, positive timeout, and a writer and reader"}
		return p
	}
	if policy == "wait" {
		p.RequestedReplicas = replicas
		p.Acknowledgements = make([]int64, len(keys))
		p.WaitAttempted = make([]bool, len(keys))
		p.AcknowledgementObserved = make([]bool, len(keys))
	} else if policy == "waitaof" {
		p.WaitAOFRequestedLocal, p.WaitAOFRequestedReplicas = 1, replicas
		p.WaitAOFAttempted = make([]bool, len(keys))
		p.WaitAOFLocalAcknowledgements = make([]int64, len(keys))
		p.WaitAOFReplicaAcknowledgements = make([]int64, len(keys))
		p.WaitAOFAcknowledgementObserved = make([]bool, len(keys))
	}
	results := make([]string, len(keys))
	value := runID + ":policy:" + fmt.Sprint(time.Now().UnixNano())
	var wg sync.WaitGroup
	for i, key := range keys {
		wg.Add(1)
		go func(i int, key string) {
			defer wg.Done()
			opCtx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			// Distinct keys prevent concurrent ordinary probes overwriting values.
			key += ":policy"
			result := writer.Write(opCtx, key, value)
			p.SetAcknowledged[i] = result.SetAcknowledged
			policyOK := result.SetAcknowledged && result.SetError == nil && result.PolicySatisfied
			if policy == "wait" {
				p.Acknowledgements[i], p.WaitAttempted[i] = result.WaitAcknowledgements, result.WaitAttempted
				p.AcknowledgementObserved[i] = result.WaitAttempted && result.WaitError == nil && result.WaitAcknowledgements >= 0
				policyOK = policyOK && p.AcknowledgementObserved[i] && result.WaitRequested == replicas && result.WaitAcknowledgements >= int64(replicas)
				if !policyOK {
					results[i] = fmt.Sprintf("shard %d: SET acknowledged=%t; WAIT=%d/%d; SET error=%v; WAIT error=%v", i, result.SetAcknowledged, result.WaitAcknowledgements, replicas, result.SetError, result.WaitError)
				}
			} else if policy == "waitaof" {
				p.WaitAOFAttempted[i] = result.WaitAOFAttempted
				p.WaitAOFLocalAcknowledgements[i], p.WaitAOFReplicaAcknowledgements[i] = result.WaitAOFLocalAcknowledgements, result.WaitAOFReplicaAcknowledgements
				p.WaitAOFAcknowledgementObserved[i] = result.WaitAOFAttempted && result.WaitAOFError == nil && result.WaitAOFLocalAcknowledgements >= 0 && result.WaitAOFLocalAcknowledgements <= 1 && result.WaitAOFReplicaAcknowledgements >= 0
				policyOK = policyOK && p.WaitAOFAcknowledgementObserved[i] && result.WaitAOFRequestedLocal == 1 && result.WaitAOFRequestedReplicas == replicas && result.WaitAOFLocalAcknowledgements == 1 && result.WaitAOFReplicaAcknowledgements >= int64(replicas)
				if !policyOK {
					results[i] = fmt.Sprintf("shard %d: SET acknowledged=%t; WAITAOF local=%d/1 replicas=%d/%d; SET error=%v; WAITAOF error=%v", i, result.SetAcknowledged, result.WaitAOFLocalAcknowledgements, result.WaitAOFReplicaAcknowledgements, replicas, result.SetError, result.WaitAOFError)
				}
			} else if !policyOK {
				results[i] = fmt.Sprintf("shard %d: SET acknowledged=%t; SET error=%v", i, result.SetAcknowledged, result.SetError)
			}
			if !policyOK {
				return
			}
			actual, err := reader.Get(opCtx, key).Result()
			if err != nil || actual != value || opCtx.Err() != nil {
				results[i] = fmt.Sprintf("shard %d: policy-probe read mismatch or error: %v", i, err)
			}
		}(i, key)
	}
	wg.Wait()
	p.At = time.Now()
	p.Successful = len(keys) == 3
	for _, err := range results {
		if err != "" {
			p.Successful = false
			p.Errors = append(p.Errors, err)
		}
	}
	if persistenceSampleExpired(ctx, p.At) {
		p.Successful = false
		p.Errors = append(p.Errors, "policy probe crossed its parent deadline or cancellation")
	}
	return p
}

// A policy command can finish after the ordinary probe. Recovery decisions use
// the last actual observation, including its command reply, before the deadline.
func recoverySampleExpired(ctx context.Context, s report.Sample) bool {
	at := s.Probe.At
	if s.Snapshot.Completed.After(at) {
		at = s.Snapshot.Completed
	}
	if s.WritePolicyProbe != nil && s.WritePolicyProbe.At.After(at) {
		at = s.WritePolicyProbe.At
	}
	return persistenceSampleExpired(ctx, at)
}

func updatePolicyRecovery(r *report.Report, s report.Sample, stability *Stability, duration time.Duration) {
	if s.WritePolicyProbe == nil || s.Excluded {
		return
	}
	p := s.WritePolicyProbe
	stable := stability.Update(p.At, p.Successful, duration)
	r.Recovery.WritePolicyRecovered = &stable
	if stable && r.Recovery.WritePolicyStable == nil {
		r.Recovery.WritePolicyStable = milestone(p.At, *r.FaultRequestedAt)
	}
}
