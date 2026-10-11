# Redis Reliability Lab — repeated trials

Batch: `batch-8f7c142e2101fdf1c2cf419c`  
Outcome: **recovered**  
Started: 2026-10-11T08:54:36.841101Z  
Finished: 2026-10-11T08:55:27.866916Z

Requested: 1; started: 1; callbacks completed: 1; unstarted: 0. Cleanup failures: 0; incomplete individual evidence: 0.

| Trial outcome | Count |
|---|---:|
| recovered | 1 |
| partial | 0 |
| failed | 0 |
| aborted | 0 |
| unstarted | 0 |

Recovery success fraction: 100.00% of **all 1 requested trials**. Started-trial recovery fraction: 100.00%.

## Recovery observations

Timing statistics below are conditional on observed milestones, including observed milestones in partial or failed trials. Unobserved events are not zero-duration recoveries. Censoring means unobserved by the accepted sampled horizon; it does not prove an event physically occurred after that horizon. No-requested-fault trials are not at risk. Aborted/error horizons are separate from deadline horizons in JSON.

| Event | Observed | Censored | Deadline censored | Not at risk | Missing | Unstarted |
|---|---:|---:|---:|---:|---:|---:|
| failure_detected | 0 | 0 | 0 | 1 | 0 | 0 |
| replica_promoted | 0 | 0 | 0 | 1 | 0 | 0 |
| full_slot_coverage | 1 | 0 | 0 | 0 | 0 | 0 |
| topology_converged | 1 | 0 | 0 | 0 | 0 | 0 |
| client_first_all_shard_success | 1 | 0 | 0 | 0 | 0 | 0 |
| client_stable | 1 | 0 | 0 | 0 | 0 | 0 |
| combined_stable | 1 | 0 | 0 | 0 | 0 | 0 |
| write_policy_stable | 0 | 0 | 0 | 1 | 0 | 0 |
| replication_interrupted | 0 | 0 | 0 | 1 | 0 | 0 |
| target_reachable_after_repair | 0 | 0 | 0 | 1 | 0 | 0 |
| membership_rejoined | 0 | 0 | 0 | 1 | 0 | 0 |
| replica_link_up | 0 | 0 | 0 | 1 | 0 | 0 |
| replica_sync_complete | 0 | 0 | 0 | 1 | 0 | 0 |
| stable_redundancy | 1 | 0 | 0 | 0 | 0 | 0 |
| original_topology_stable | 1 | 0 | 0 | 0 | 0 | 0 |
| persistence_recovered | 1 | 0 | 0 | 0 | 0 | 0 |
| all_processes_stopped | 1 | 0 | 0 | 0 | 0 | 0 |

Descriptive timing statistics use milliseconds. **P90/P95 are omitted for N<20**; small samples do not support reliable tail estimates. Sample standard deviation uses N-1 and is undefined for N<2. These are descriptive statistics, not confidence intervals.

| Timing / anchor | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| fault action duration | 1 | 2186.243 | 2186.243 | 2186.243 | undefined | omitted | omitted | 2186.243 |
| failure_detected / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| failure_detected / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_promoted / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_promoted / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| full_slot_coverage / orchestration start | 1 | 16340.912 | 16340.912 | 16340.912 | undefined | omitted | omitted | 16340.912 |
| full_slot_coverage / action completion | 1 | 14154.669 | 14154.669 | 14154.669 | undefined | omitted | omitted | 14154.669 |
| topology_converged / orchestration start | 1 | 17585.182 | 17585.182 | 17585.182 | undefined | omitted | omitted | 17585.182 |
| topology_converged / action completion | 1 | 15398.939 | 15398.939 | 15398.939 | undefined | omitted | omitted | 15398.939 |
| client_first_all_shard_success / orchestration start | 1 | 23604.988 | 23604.988 | 23604.988 | undefined | omitted | omitted | 23604.988 |
| client_first_all_shard_success / action completion | 1 | 21418.746 | 21418.746 | 21418.746 | undefined | omitted | omitted | 21418.746 |
| client_stable / orchestration start | 1 | 24855.190 | 24855.190 | 24855.190 | undefined | omitted | omitted | 24855.190 |
| client_stable / action completion | 1 | 22668.948 | 22668.948 | 22668.948 | undefined | omitted | omitted | 22668.948 |
| combined_stable / orchestration start | 1 | 24855.190 | 24855.190 | 24855.190 | undefined | omitted | omitted | 24855.190 |
| combined_stable / action completion | 1 | 22668.948 | 22668.948 | 22668.948 | undefined | omitted | omitted | 22668.948 |
| write_policy_stable / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| write_policy_stable / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replication_interrupted / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replication_interrupted / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| target_reachable_after_repair / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| target_reachable_after_repair / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| membership_rejoined / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| membership_rejoined / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_link_up / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_link_up / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_sync_complete / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_sync_complete / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| stable_redundancy / orchestration start | 1 | 17585.182 | 17585.182 | 17585.182 | undefined | omitted | omitted | 17585.182 |
| stable_redundancy / action completion | 1 | 15398.939 | 15398.939 | 15398.939 | undefined | omitted | omitted | 15398.939 |
| original_topology_stable / orchestration start | 1 | 17585.182 | 17585.182 | 17585.182 | undefined | omitted | omitted | 17585.182 |
| original_topology_stable / action completion | 1 | 15398.939 | 15398.939 | 15398.939 | undefined | omitted | omitted | 15398.939 |
| persistence_recovered / orchestration start | 1 | 23600.379 | 23600.379 | 23600.379 | undefined | omitted | omitted | 23600.379 |
| persistence_recovered / action completion | 1 | 21414.136 | 21414.136 | 21414.136 | undefined | omitted | omitted | 21414.136 |
| all_processes_stopped / orchestration start | 1 | 2186.243 | 2186.243 | 2186.243 | undefined | omitted | omitted | 2186.243 |
| all_processes_stopped / action completion | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |

## Client operations

Completed request fractions weight actual request counts. Throughput divides summed completed/successful counts by summed trial phase wall-clock seconds; it does not imply workers achieved the configured offered rate. Zero-attempt phases have an unobserved success fraction.

| Phase | Trials with phase | Completed | Successful | Failed | Request success fraction | Completed ops/s | Successful ops/s |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 1 | 804 | 804 | 0 | 100.00% | 199.806 | 199.806 |

Whole-workload schedule counters measured in 1 trials; legacy or incomplete counters unmeasured in 0. Offered: 804; started: 804; completed: 804; dropped: 0. Started fraction of offered: 100.00%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| baseline / completed ops/s | 1 | 199.806 | 199.806 | 199.806 | undefined | omitted | omitted | 199.806 |
| baseline / successful request p50 ms | 1 | 0.520 | 0.520 | 0.520 | undefined | omitted | omitted | 0.520 |
| baseline / successful request p95 ms | 1 | 0.732 | 0.732 | 0.732 | undefined | omitted | omitted | 0.732 |
| baseline / successful request p99 ms | 1 | 0.981 | 0.981 | 0.981 | undefined | omitted | omitted | 0.981 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 1 trials, unmeasured in 0. Every successful SET remains in the main validation ledger even when WAIT is short, errors, or is not attempted. Stage command latency remains in individual reports; it does not increase logical request counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 417 | 0 | 0 | 0 | 0 | 0 | 0 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. Attempted WAIT client-call timeouts are counted separately above.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|

## Acknowledged-write accounting

Validation attempted in 1 trials and complete in 1. SET attempts: 417; acknowledged writes: 417; ambiguous SET outcomes: 0.

Checked: 417; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 1 | 417.000 | 417.000 | 417.000 | undefined | omitted | omitted | 417.000 |
| ambiguous writes | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| missing writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | 29bae03a4d4db021a7868f0f | recovered | true | true | [JSON](runs/29bae03a4d4db021a7868f0f/report.json), [Markdown](runs/29bae03a4d4db021a7868f0f/report.md) |

## Measurement methodology

- Trials execute sequentially with a fresh environment per callback; the next callback starts only after the previous callback returns with successful cleanup evidence (or an authoritative report that no environment was provisioned).
- Recovery success fraction uses all requested trials as the denominator. Outcomes, started/unstarted counts and individual reports preserve failed, partial and aborted trials; the started-trial fraction is reported separately.
- Recovery timings are conditional on observed milestones, including observed milestones in partial/failed trials. Unobserved post-fault events are censored at the accepted sampled observation horizon with deadline/abort/error reasons, never zero-duration samples; this does not prove the physical event occurred after that horizon. Trials without a requested fault are separately not at risk. No survival distribution, independence assumption or population reliability guarantee is inferred.
- Event durations use runtime SinceFaultMS and fault_action_duration_ms when present, preserving monotonic elapsed-time measurements. Schema 1 action-completion-relative durations fall back to wall-clock timestamp differences and can be affected by clock adjustments. First sampled observations include polling uncertainty; action duration includes ownership checks and fault command overhead.
- Min, median, mean, N-1 sample standard deviation and maximum describe available trial observations. P90/P95 use nearest rank only when N>=20; smaller samples have an explicit tail-estimation caveat.
- Completed-request success fractions weight actual completed counts; aggregate throughput divides summed completed/successful counts by summed phase wall-clock duration. Offered load and worker saturation are not inferred from completed requests.
- Whole-workload offered, started, completed and dropped schedule-slot counters are aggregated only for reports carrying explicit pacing methodology. Legacy absent counters are unmeasured, not zero. Offered=started+dropped and started=completed after workload quiescence; service probes and validation are excluded.
- Acknowledged SET writes remain distinct from ambiguous SET outcomes and final read-back validation. Skipped/incomplete validation contributes acknowledged-not-checked counts; no observed loss is not a durability guarantee.
- Recovered schema 2 reports require the explicit combined-recovery flag. Legacy schema 1 reports can establish it through healthy topology/client flags and an observed combined-stability milestone, together with complete acknowledged-write validation and verified cleanup.
- Recovered primary-isolate trials additionally require sampled unchanged process survival while disconnected, restored connectivity, observed stable reconciliation and restored replication redundancy. Pre-restoration recovery alone does not establish a recovered network experiment.

Effective experiment configuration SHA-256: `216e98ee4d66a4cd497a054a7ad6776564916f52380dbc7586e7f3e402dca046`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
