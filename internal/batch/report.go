package batch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func (s Summary) Markdown() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "# Redis Reliability Lab — repeated trials\n\nBatch: `%s`  \nOutcome: **%s**  \nStarted: %s  \nFinished: %s\n\n", s.BatchID, s.Outcome, s.StartedAt.Format(time.RFC3339Nano), s.FinishedAt.Format(time.RFC3339Nano))
	fmt.Fprintf(&b, "Requested: %d; started: %d; callbacks completed: %d; unstarted: %d. Cleanup failures: %d; incomplete individual evidence: %d.\n\n", s.RequestedTrials, s.StartedTrials, s.CompletedTrials, s.UnstartedTrials, s.CleanupFailures, s.IncompleteEvidenceTrials)
	fmt.Fprint(&b, "| Trial outcome | Count |\n|---|---:|\n")
	for _, outcome := range []string{"recovered", "partial", "failed", "aborted", "unstarted"} {
		fmt.Fprintf(&b, "| %s | %d |\n", outcome, s.Outcomes[outcome])
	}
	fmt.Fprintf(&b, "\nRecovery success fraction: %.2f%% of **all %d requested trials**. Started-trial recovery fraction: %s.\n", 100*s.RecoverySuccessFraction, s.RequestedTrials, fraction(s.ExecutedRecoverySuccessFraction))
	if s.StopReason != "" {
		fmt.Fprintf(&b, "\nStopped before further trials: `%s`.\n", s.StopReason)
	}
	fmt.Fprint(&b, "\n## Recovery observations\n\nTiming statistics below are conditional on observed milestones, including observed milestones in partial or failed trials. Unobserved events are not zero-duration recoveries. Censoring means unobserved by the accepted sampled horizon; it does not prove an event physically occurred after that horizon. No-requested-fault trials are not at risk. Aborted/error horizons are separate from deadline horizons in JSON.\n\n| Event | Observed | Censored | Deadline censored | Not at risk | Missing | Unstarted |\n|---|---:|---:|---:|---:|---:|---:|\n")
	for _, name := range eventNames {
		e := s.Events[name]
		fmt.Fprintf(&b, "| %s | %d | %d | %d | %d | %d | %d |\n", name, e.Observed, e.Censored, e.CensoredDeadline, e.NotAtRisk, e.Missing, e.Unstarted)
	}
	fmt.Fprint(&b, "\nDescriptive timing statistics use milliseconds. **P90/P95 are omitted for N<20**; small samples do not support reliable tail estimates. Sample standard deviation uses N-1 and is undefined for N<2. These are descriptive statistics, not confidence intervals.\n\n")
	statisticsHeader(&b, "Timing / anchor")
	statisticsRow(&b, "fault action duration", s.FaultActionDurationMS)
	for _, name := range eventNames {
		e := s.Events[name]
		statisticsRow(&b, name+" / orchestration start", e.SinceFaultRequestedMS)
		statisticsRow(&b, name+" / action completion", e.SinceFaultCompletedMS)
	}
	fmt.Fprint(&b, "\n## Client operations\n\nCompleted request fractions weight actual request counts. Throughput divides summed completed/successful counts by summed trial phase wall-clock seconds; it does not imply workers achieved the configured offered rate. Zero-attempt phases have an unobserved success fraction.\n\n| Phase | Trials with phase | Completed | Successful | Failed | Request success fraction | Completed ops/s | Successful ops/s |\n|---|---:|---:|---:|---:|---:|---:|---:|\n")
	var phases []string
	for phase := range s.Phases {
		phases = append(phases, phase)
	}
	sort.Strings(phases)
	for _, phase := range phases {
		p := s.Phases[phase]
		fmt.Fprintf(&b, "| %s | %d | %d | %d | %d | %s | %s | %s |\n", phase, p.TrialsWithPhase, p.Attempts, p.Successes, p.Errors, fraction(p.SuccessFraction), statValue(p.CompletedThroughput), statValue(p.SuccessfulThroughput))
	}
	fmt.Fprintf(&b, "\nWhole-workload schedule counters measured in %d trials; legacy or incomplete counters unmeasured in %d. Offered: %d; started: %d; completed: %d; dropped: %d. Started fraction of offered: %s. These counters exclude service probes and validation.\n", s.Load.MeasuredTrials, s.Load.UnmeasuredTrials, s.Load.OfferedRequests, s.Load.StartedRequests, s.Load.CompletedRequests, s.Load.DroppedRequests, fraction(s.Load.StartedFraction))
	fmt.Fprint(&b, "\n### Across-trial performance observations\n\n")
	statisticsHeader(&b, "Phase / metric")
	for _, phase := range phases {
		p := s.Phases[phase]
		statisticsRow(&b, phase+" / trial success fraction", p.TrialSuccessFraction)
		statisticsRow(&b, phase+" / completed ops/s", p.TrialCompletedThroughput)
		statisticsRow(&b, phase+" / successful request p50 ms", p.TrialSuccessfulP50MS)
		statisticsRow(&b, phase+" / successful request p95 ms", p.TrialSuccessfulP95MS)
		statisticsRow(&b, phase+" / successful request p99 ms", p.TrialSuccessfulP99MS)
		statisticsRow(&b, phase+" / failed request p50 ms", p.TrialFailedP50MS)
		statisticsRow(&b, phase+" / failed request p95 ms", p.TrialFailedP95MS)
		statisticsRow(&b, phase+" / failed request p99 ms", p.TrialFailedP99MS)
	}
	fmt.Fprintf(&b, "\n%s\n", s.LatencyAggregation)
	if s.WritePolicy.MeasuredTrials > 0 {
		s.policyMarkdown(&b)
	}
	fmt.Fprint(&b, "\n### Terminal errors\n\n| Phase | Class | Count |\n|---|---|---:|\n")
	for _, phase := range phases {
		var classes []string
		for class, count := range s.Phases[phase].ErrorClasses {
			if count > 0 {
				classes = append(classes, class)
			}
		}
		sort.Strings(classes)
		for _, class := range classes {
			fmt.Fprintf(&b, "| %s | %s | %d |\n", phase, class, s.Phases[phase].ErrorClasses[class])
		}
	}
	d := s.Data
	fmt.Fprintf(&b, "\n## Acknowledged-write accounting\n\nValidation attempted in %d trials and complete in %d. SET attempts: %d; acknowledged writes: %d; ambiguous SET outcomes: %d.\n\nChecked: %d; missing: %d; corrupt: %d; unreadable: %d; explicitly unverified: %d; acknowledged writes not checked (including skipped validation): %d. During-traffic missing checks: %d; mismatched checks: %d.\n\nTotals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.\n\n", d.ValidationAttemptedTrials, d.ValidationCompleteTrials, d.WritesAttempted, d.AcknowledgedWrites, d.AmbiguousWrites, d.Checked, d.Missing, d.Corrupt, d.Unreadable, d.Unverified, d.AcknowledgedNotChecked, d.MissingReadChecks, d.MismatchedReadChecks)
	statisticsHeader(&b, "Trial data metric")
	statisticsRow(&b, "acknowledged writes", d.TrialAcknowledgedWrites)
	statisticsRow(&b, "ambiguous writes", d.TrialAmbiguousWrites)
	statisticsRow(&b, "missing writes / validation attempted", d.TrialMissingWrites)
	statisticsRow(&b, "corrupt writes / validation attempted", d.TrialCorruptWrites)
	statisticsRow(&b, "unreadable writes / validation attempted", d.TrialUnreadableWrites)
	fmt.Fprint(&b, "\n## Individual trials\n\n| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |\n|---:|---|---|---|---|---|\n")
	for _, trial := range s.Trials {
		link := "unavailable"
		if trial.ReportPath != "" {
			link = fmt.Sprintf("[JSON](%s)", filepath.ToSlash(trial.ReportPath))
			if trial.ReportMarkdownAvailable {
				link += fmt.Sprintf(", [Markdown](%s)", filepath.ToSlash(filepath.Join(filepath.Dir(trial.ReportPath), "report.md")))
			} else {
				link += ", Markdown unavailable"
			}
		} else if !trial.Started {
			link = "unstarted"
		}
		fmt.Fprintf(&b, "| %d | %s | %s | %t | %t | %s |\n", trial.Index, trial.RunID, trial.Outcome, trial.EvidenceComplete, trial.Cleanup.Successful, link)
	}
	if len(s.Errors) > 0 {
		fmt.Fprint(&b, "\n## Execution errors\n\n")
		for _, err := range s.Errors {
			fmt.Fprintf(&b, "- %s\n", strings.ReplaceAll(err, "\n", "; "))
		}
	}
	fmt.Fprint(&b, "\n## Measurement methodology\n\n")
	for _, method := range s.Methodology {
		fmt.Fprintf(&b, "- %s\n", method)
	}
	fmt.Fprintf(&b, "\nEffective experiment configuration SHA-256: `%s`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.\n", s.ConfigSHA256)
	return b.String()
}

func statisticsHeader(b *bytes.Buffer, label string) {
	fmt.Fprintf(b, "| %s | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |\n|---|---:|---:|---:|---:|---:|---:|---:|---:|\n", label)
}
func statisticsRow(b *bytes.Buffer, label string, s Statistics) {
	std := statValue(s.SampleStdDev)
	if s.N == 1 {
		std = "undefined"
	}
	p90, p95 := statValue(s.P90), statValue(s.P95)
	if s.N < minimumPercentileSamples {
		p90, p95 = "omitted", "omitted"
	}
	fmt.Fprintf(b, "| %s | %d | %s | %s | %s | %s | %s | %s | %s |\n", label, s.N, statValue(s.Min), statValue(s.Median), statValue(s.Mean), std, p90, p95, statValue(s.Max))
}
func statValue(value *float64) string {
	if value == nil {
		return "unobserved"
	}
	return fmt.Sprintf("%.3f", *value)
}
func fraction(value *float64) string {
	if value == nil {
		return "unobserved"
	}
	return fmt.Sprintf("%.2f%%", *value*100)
}

// Write replaces each completed summary file atomically, with write errors
// returned. Like individual reports, the pair is not a transactional filesystem.
func Write(dir string, s Summary) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("encode batch summary: %w", err)
	}
	if err := atomicWrite(filepath.Join(dir, "summary.json"), append(data, '\n')); err != nil {
		return err
	}
	return atomicWrite(filepath.Join(dir, "summary.md"), []byte(s.Markdown()))
}

func atomicWrite(path string, data []byte) error {
	f, err := os.CreateTemp(filepath.Dir(path), ".summary-*")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	if err := f.Sync(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), path)
}
