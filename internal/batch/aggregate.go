package batch

import (
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/report"
	"github.com/Alaneel/redis-reliability-lab/internal/workload"
)

var eventNames = []string{"failure_detected", "replica_promoted", "full_slot_coverage", "topology_converged", "client_first_all_shard_success", "client_stable", "combined_stable", "write_policy_stable", "replication_interrupted", "target_reachable_after_repair", "membership_rejoined", "replica_link_up", "replica_sync_complete", "stable_redundancy", "original_topology_stable", "persistence_recovered", "all_processes_stopped"}

func observations(r *report.Report) map[string]Observation {
	result := make(map[string]Observation, len(eventNames))
	milestones := []*report.Milestone{r.Recovery.FailureDetected, r.Recovery.ReplicaPromoted, r.Recovery.FullCoverage, r.Recovery.TopologyConverged, r.Recovery.ClientFirstSuccess, r.Recovery.ClientStable, r.Recovery.CombinedStable, r.Recovery.WritePolicyStable, nil, nil, nil, nil, nil, nil, nil, nil, nil}
	if l := r.LifecycleFault; l != nil {
		milestones[8], milestones[9], milestones[10], milestones[11], milestones[12], milestones[13] = l.ReplicationInterrupted, l.TargetReachable, l.MembershipRejoined, l.LinkUp, l.SyncComplete, l.StableRedundancy
	}
	if p := r.PersistenceFault; p != nil {
		milestones[13], milestones[14], milestones[15] = p.StableRedundancy, p.TopologyStable, p.PersistenceRecoveryAt
		if validTime(p.AllStoppedAt) && nonnegativeObservation(r.FaultActionDurationMS) {
			milestones[16] = &report.Milestone{At: *p.AllStoppedAt, SinceFaultMS: *r.FaultActionDurationMS}
		}
	}
	var horizon time.Time
	for _, sample := range r.Samples {
		if sample.Excluded {
			continue
		}
		at := sample.Snapshot.Completed
		if sample.Probe.At.After(at) {
			at = sample.Probe.At
		}
		if sample.WritePolicyProbe != nil && sample.WritePolicyProbe.At.After(at) {
			at = sample.WritePolicyProbe.At
		}
		if at.After(horizon) {
			horizon = at
		}
	}
	for i, name := range eventNames {
		lifecycleOnly := i >= 8 && i <= 12
		persistenceOnly := i >= 14
		notApplicable := lifecycleOnly && r.LifecycleFault == nil || persistenceOnly && r.PersistenceFault == nil || name == "stable_redundancy" && r.LifecycleFault == nil && r.PersistenceFault == nil || name == "persistence_recovered" && r.Config.Environment.Persistence == "memory"
		if notApplicable || name == "failure_detected" && r.Config.Experiment.Fault == "cluster-restart" || (name == "replica_promoted" && (r.Config.Experiment.Fault == "replica-kill" || r.Config.Experiment.Fault == "replica-isolate" || r.Config.Experiment.Fault == "cluster-restart")) {
			result[name] = Observation{Status: "not_at_risk"}
			continue
		}
		if name == "write_policy_stable" && r.Config.Workload.WritePolicy != "wait" && r.Config.Workload.WritePolicy != "waitaof" {
			result[name] = Observation{Status: "not_at_risk"}
			continue
		}
		value := milestones[i]
		eventHorizon := horizon
		restoreDeadline := false
		if p := r.PersistenceFault; p != nil {
			switch name {
			case "stable_redundancy", "original_topology_stable", "full_slot_coverage", "topology_converged":
				eventHorizon = time.Time{}
				for _, sample := range r.Samples {
					if !sample.Excluded && sample.Phase == "persistence_recovery" && sample.Snapshot.Completed.After(eventHorizon) {
						eventHorizon = sample.Snapshot.Completed
					}
				}
				restoreDeadline = p.RestoreTimedOut
			case "persistence_recovered":
				eventHorizon = time.Time{}
				if validTime(p.DataCheckedAt) {
					eventHorizon = *p.DataCheckedAt
				}
				restoreDeadline = p.RestoreTimedOut || r.Validation.Error == "context deadline exceeded"
			case "all_processes_stopped":
				eventHorizon = time.Time{}
				if validTime(p.AllStoppedAt) {
					eventHorizon = *p.AllStoppedAt
				}
				for _, sample := range p.Stopped {
					if sample.Evidence.ObservedAt.After(eventHorizon) {
						eventHorizon = sample.Evidence.ObservedAt
					}
				}
			}
		}
		o := Observation{Status: "missing"}
		if !eventHorizon.IsZero() {
			at := eventHorizon
			o.ObservationHorizonAt = &at
			if r.FaultRequestedAt != nil {
				o.HorizonSinceFaultRequestedMS = number(float64(eventHorizon.Sub(*r.FaultRequestedAt)) / float64(time.Millisecond))
			}
			if r.FaultCompletedAt != nil {
				o.HorizonSinceFaultCompletedMS = number(float64(eventHorizon.Sub(*r.FaultCompletedAt)) / float64(time.Millisecond))
			}
		}
		if value != nil && !value.At.IsZero() && finite(value.SinceFaultMS) {
			o.Status = "observed"
			if r.FaultRequestedAt != nil {
				o.SinceFaultRequestedMS = number(value.SinceFaultMS)
			}
			if r.FaultCompletedAt != nil {
				if r.FaultActionDurationMS != nil {
					o.SinceFaultCompletedMS = number(value.SinceFaultMS - *r.FaultActionDurationMS)
				} else {
					o.SinceFaultCompletedMS = number(float64(value.At.Sub(*r.FaultCompletedAt)) / float64(time.Millisecond))
				}
			}
		} else if r.FaultRequestedAt == nil {
			o.Status = "not_at_risk"
		} else if r.FaultCompletedAt == nil {
			o.Status = "missing_fault_not_completed"
		} else if (r.Recovery.TimedOut && (r.PersistenceFault == nil || i >= 4 && i <= 7)) || (i >= 9 && r.LifecycleFault != nil && r.LifecycleFault.RepairTimedOut) || restoreDeadline {
			o.Status = "censored_deadline"
		} else if r.Outcome == "aborted" {
			o.Status = "censored_aborted"
		} else if len(r.Errors) > 0 {
			o.Status = "censored_error"
		} else {
			o.Status = "censored_other"
		}
		result[name] = o
	}
	return result
}

// Aggregate recomputes summary statistics from compact trial evidence. It does
// not invent observations for failed/unstarted trials or treat them as zero.
func Aggregate(s *Summary) {
	s.StartedTrials, s.CompletedTrials, s.CleanupFailures, s.IncompleteEvidenceTrials = 0, 0, 0, 0
	s.Outcomes = map[string]int{"recovered": 0, "partial": 0, "failed": 0, "aborted": 0, "unstarted": 0}
	s.Events = make(map[string]EventStatistics, len(eventNames))
	s.Phases = make(map[string]PhaseStatistics)
	s.Data = DataStatistics{}
	s.Load = LoadStatistics{}
	s.WritePolicy = PolicyStatistics{ByPhase: map[string]workload.PolicyCounts{}}
	var faultDurations, acknowledged, ambiguous, missing, corrupt, unreadable []float64
	var offered, started, dropped []float64
	for _, trial := range s.Trials {
		s.Outcomes[trial.Outcome]++
		if !trial.Started {
			continue
		}
		s.StartedTrials++
		if !trial.FinishedAt.IsZero() {
			s.CompletedTrials++
		}
		if !trial.EvidenceComplete {
			s.IncompleteEvidenceTrials++
		}
		if trial.EnvironmentPresent && (!trial.Cleanup.Attempted || !trial.Cleanup.Successful) {
			s.CleanupFailures++
		}
		if !trial.EvidenceComplete {
			s.Load.UnmeasuredTrials++
			s.WritePolicy.UnmeasuredTrials++
			continue
		}
		if trial.FaultActionDurationMS != nil {
			faultDurations = append(faultDurations, *trial.FaultActionDurationMS)
		}
		w, v := trial.WorkloadStats, trial.Validation
		if w.WritePolicy == "" {
			s.WritePolicy.UnmeasuredTrials++
		} else {
			s.WritePolicy.MeasuredTrials++
			s.WritePolicy.Totals = sumPolicyCounts(s.WritePolicy.Totals, w.PolicyTotals)
			for phase, counts := range w.PolicyByPhase {
				s.WritePolicy.ByPhase[phase] = sumPolicyCounts(s.WritePolicy.ByPhase[phase], counts)
			}
			if w.WritePolicy == "wait" && !v.StartedAt.IsZero() {
				s.WritePolicy.WaitSubsetValidationTrials++
				d := &s.WritePolicy.WaitSatisfiedSubset
				v := v.WaitSatisfiedSubset
				d.Expected += v.Expected
				d.Checked += v.Checked
				d.Missing += v.Missing
				d.Corrupt += v.Corrupt
				d.Unreadable += v.Unreadable
				d.Unverified += v.Unverified
			} else if w.WritePolicy == "wait" {
				n := int(w.PolicyTotals.WaitSatisfied)
				s.WritePolicy.WaitSatisfiedSubset.Expected += n
				s.WritePolicy.WaitSatisfiedSubset.Unverified += n
			}
			if w.WritePolicy == "waitaof" && !v.StartedAt.IsZero() {
				s.WritePolicy.WaitAOFSubsetValidationTrials++
				d := &s.WritePolicy.WaitAOFSatisfiedSubset
				v := v.WaitAOFSatisfiedSubset
				d.Expected += v.Expected
				d.Checked += v.Checked
				d.Missing += v.Missing
				d.Corrupt += v.Corrupt
				d.Unreadable += v.Unreadable
				d.Unverified += v.Unverified
			} else if w.WritePolicy == "waitaof" {
				n := int(w.PolicyTotals.WaitAOFSatisfied)
				s.WritePolicy.WaitAOFSatisfiedSubset.Expected += n
				s.WritePolicy.WaitAOFSatisfiedSubset.Unverified += n
			}
		}
		if w.PacingMethodology != "" {
			s.Load.MeasuredTrials++
			s.Load.OfferedRequests += w.OfferedRequests
			s.Load.StartedRequests += w.StartedRequests
			s.Load.CompletedRequests += w.CompletedRequests
			s.Load.DroppedRequests += w.DroppedRequests
			offered = append(offered, float64(w.OfferedRequests))
			started = append(started, float64(w.StartedRequests))
			dropped = append(dropped, float64(w.DroppedRequests))
		} else {
			s.Load.UnmeasuredTrials++
		}
		s.Data.WritesAttempted += w.WritesAttempted
		s.Data.AcknowledgedWrites += w.AcknowledgedWrites
		s.Data.AmbiguousWrites += w.AmbiguousWrites
		s.Data.MissingReadChecks += w.MissingReads
		s.Data.MismatchedReadChecks += w.MismatchedReads
		if !trial.SyntheticErrorReport {
			acknowledged = append(acknowledged, float64(w.AcknowledgedWrites))
			ambiguous = append(ambiguous, float64(w.AmbiguousWrites))
		}
		if v.StartedAt.IsZero() {
			s.Data.AcknowledgedNotChecked += w.AcknowledgedWrites
			continue
		}
		s.Data.ValidationAttemptedTrials++
		if v.Complete {
			s.Data.ValidationCompleteTrials++
		}
		s.Data.Checked += uint64(v.Checked)
		s.Data.Missing += uint64(v.Missing)
		s.Data.Corrupt += uint64(v.Corrupt)
		s.Data.Unreadable += uint64(v.Unreadable)
		s.Data.Unverified += uint64(v.Unverified)
		if w.AcknowledgedWrites > uint64(v.Checked) {
			s.Data.AcknowledgedNotChecked += w.AcknowledgedWrites - uint64(v.Checked)
		}
		missing = append(missing, float64(v.Missing))
		corrupt = append(corrupt, float64(v.Corrupt))
		unreadable = append(unreadable, float64(v.Unreadable))
	}
	s.UnstartedTrials = s.RequestedTrials - s.StartedTrials
	s.WritePolicy.WaitSatisfiedSubset.Complete = s.WritePolicy.WaitSubsetValidationTrials > 0 && s.WritePolicy.WaitSatisfiedSubset.Unverified == 0
	s.WritePolicy.WaitAOFSatisfiedSubset.Complete = s.WritePolicy.WaitAOFSubsetValidationTrials > 0 && s.WritePolicy.WaitAOFSatisfiedSubset.Unverified == 0
	if s.RequestedTrials > 0 {
		s.RecoverySuccessFraction = float64(s.Outcomes["recovered"]) / float64(s.RequestedTrials)
	}
	s.ExecutedRecoverySuccessFraction = nil
	if s.StartedTrials > 0 {
		s.ExecutedRecoverySuccessFraction = number(float64(s.Outcomes["recovered"]) / float64(s.StartedTrials))
	}
	s.FaultActionDurationMS = describe(faultDurations)
	s.Load.TrialOfferedRequests, s.Load.TrialStartedRequests, s.Load.TrialDroppedRequests = describe(offered), describe(started), describe(dropped)
	if s.Load.OfferedRequests > 0 {
		s.Load.StartedFraction = number(float64(s.Load.StartedRequests) / float64(s.Load.OfferedRequests))
	}
	s.Data.TrialAcknowledgedWrites = describe(acknowledged)
	s.Data.TrialAmbiguousWrites = describe(ambiguous)
	s.Data.TrialMissingWrites = describe(missing)
	s.Data.TrialCorruptWrites = describe(corrupt)
	s.Data.TrialUnreadableWrites = describe(unreadable)
	for _, name := range eventNames {
		stat := EventStatistics{MissingReasons: make(map[string]int)}
		var requested, completed []float64
		for _, trial := range s.Trials {
			if !trial.Started {
				stat.Unstarted++
				continue
			}
			if !trial.EvidenceComplete {
				stat.Missing++
				stat.MissingReasons["incomplete_trial_evidence"]++
				continue
			}
			o := trial.Observations[name]
			switch o.Status {
			case "observed":
				stat.Observed++
				if o.SinceFaultRequestedMS != nil {
					requested = append(requested, *o.SinceFaultRequestedMS)
				} else {
					stat.RequestedAnchorMissing++
				}
				if o.SinceFaultCompletedMS != nil {
					completed = append(completed, *o.SinceFaultCompletedMS)
				} else {
					stat.CompletedAnchorMissing++
				}
			case "not_at_risk":
				stat.NotAtRisk++
			case "censored_deadline":
				stat.Censored++
				stat.CensoredDeadline++
			case "censored_aborted":
				stat.Censored++
				stat.CensoredAborted++
			case "censored_error":
				stat.Censored++
				stat.CensoredError++
			case "censored_other":
				stat.Censored++
				stat.CensoredOther++
			default:
				stat.Missing++
				reason := o.Status
				if reason == "" {
					reason = "missing_report"
				}
				stat.MissingReasons[reason]++
			}
		}
		stat.SinceFaultRequestedMS, stat.SinceFaultCompletedMS = describe(requested), describe(completed)
		s.Events[name] = stat
	}
	for _, trial := range s.Trials {
		if !trial.Started || !trial.EvidenceComplete {
			continue
		}
		for phase := range trial.Phases {
			if _, found := s.Phases[phase]; !found {
				s.Phases[phase] = aggregatePhase(s.Trials, phase)
			}
		}
	}
	if s.StopReason == "summary_write_failure" {
		s.Outcome = "failed"
	} else if s.StopReason == "cancellation" {
		s.Outcome = "aborted"
	} else if s.Outcomes["recovered"] == s.RequestedTrials {
		s.Outcome = "recovered"
	} else if s.Outcomes["recovered"]+s.Outcomes["partial"] > 0 {
		s.Outcome = "partial"
	} else {
		s.Outcome = "failed"
	}
}

func aggregatePhase(trials []Trial, phase string) PhaseStatistics {
	s := PhaseStatistics{ErrorClasses: make(map[string]uint64)}
	var success, throughput, p50, p95, p99, f50, f95, f99 []float64
	for _, trial := range trials {
		m, found := trial.Phases[phase]
		if !trial.Started || !trial.EvidenceComplete || !found {
			continue
		}
		s.TrialsWithPhase++
		s.Attempts += m.Attempts
		s.Successes += m.Successes
		s.Errors += m.Errors
		s.DurationSeconds += m.DurationSeconds
		for class, count := range m.ErrorClasses {
			s.ErrorClasses[class] += count
		}
		if m.Attempts > 0 {
			s.TrialsWithAttempts++
			success = append(success, float64(m.Successes)/float64(m.Attempts))
		}
		if m.DurationSeconds > 0 {
			throughput = append(throughput, float64(m.Attempts)/m.DurationSeconds)
		}
		if m.SuccessfulLatency.Count > 0 {
			p50 = append(p50, m.SuccessfulLatency.P50MS)
			p95 = append(p95, m.SuccessfulLatency.P95MS)
			p99 = append(p99, m.SuccessfulLatency.P99MS)
		}
		if m.FailedLatency.Count > 0 {
			f50 = append(f50, m.FailedLatency.P50MS)
			f95 = append(f95, m.FailedLatency.P95MS)
			f99 = append(f99, m.FailedLatency.P99MS)
		}
	}
	if s.Attempts > 0 {
		s.SuccessFraction = number(float64(s.Successes) / float64(s.Attempts))
	}
	if s.DurationSeconds > 0 {
		s.CompletedThroughput = number(float64(s.Attempts) / s.DurationSeconds)
		s.SuccessfulThroughput = number(float64(s.Successes) / s.DurationSeconds)
	}
	s.TrialSuccessFraction = describe(success)
	s.TrialCompletedThroughput = describe(throughput)
	s.TrialSuccessfulP50MS, s.TrialSuccessfulP95MS, s.TrialSuccessfulP99MS = describe(p50), describe(p95), describe(p99)
	s.TrialFailedP50MS, s.TrialFailedP95MS, s.TrialFailedP99MS = describe(f50), describe(f95), describe(f99)
	return s
}
