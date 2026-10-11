# Redis Reliability Lab — repeated trials

Batch: `batch-bbc2fe8f6b545c67219b3727`  
Outcome: **recovered**  
Started: 2026-10-11T07:40:37.843336Z  
Finished: 2026-10-11T07:42:04.063905Z

Requested: 3; started: 3; callbacks completed: 3; unstarted: 0. Cleanup failures: 0; incomplete individual evidence: 0.

| Trial outcome | Count |
|---|---:|
| recovered | 3 |
| partial | 0 |
| failed | 0 |
| aborted | 0 |
| unstarted | 0 |

Recovery success fraction: 100.00% of **all 3 requested trials**. Started-trial recovery fraction: 100.00%.

## Recovery observations

Timing statistics below are conditional on observed milestones, including observed milestones in partial or failed trials. Unobserved events are not zero-duration recoveries. Censoring means unobserved by the accepted sampled horizon; it does not prove an event physically occurred after that horizon. No-requested-fault trials are not at risk. Aborted/error horizons are separate from deadline horizons in JSON.

| Event | Observed | Censored | Deadline censored | Not at risk | Missing | Unstarted |
|---|---:|---:|---:|---:|---:|---:|
| failure_detected | 3 | 0 | 0 | 0 | 0 | 0 |
| replica_promoted | 3 | 0 | 0 | 0 | 0 | 0 |
| full_slot_coverage | 3 | 0 | 0 | 0 | 0 | 0 |
| topology_converged | 3 | 0 | 0 | 0 | 0 | 0 |
| client_first_all_shard_success | 3 | 0 | 0 | 0 | 0 | 0 |
| client_stable | 3 | 0 | 0 | 0 | 0 | 0 |
| combined_stable | 3 | 0 | 0 | 0 | 0 | 0 |
| write_policy_stable | 0 | 0 | 0 | 3 | 0 | 0 |

Descriptive timing statistics use milliseconds. **P90/P95 are omitted for N<20**; small samples do not support reliable tail estimates. Sample standard deviation uses N-1 and is undefined for N<2. These are descriptive statistics, not confidence intervals.

| Timing / anchor | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| fault action duration | 3 | 216.483 | 217.703 | 268.255 | 88.618 | omitted | omitted | 370.580 |
| failure_detected / orchestration start | 3 | 1778.134 | 1785.533 | 1832.045 | 87.046 | omitted | omitted | 1932.466 |
| failure_detected / action completion | 3 | 1561.651 | 1561.886 | 1563.789 | 3.502 | omitted | omitted | 1567.831 |
| replica_promoted / orchestration start | 3 | 2561.823 | 2711.090 | 2874.476 | 418.964 | omitted | omitted | 3350.516 |
| replica_promoted / action completion | 3 | 2340.510 | 2345.340 | 2606.221 | 456.049 | omitted | omitted | 3132.813 |
| full_slot_coverage / orchestration start | 3 | 2711.090 | 3340.834 | 3134.146 | 366.410 | omitted | omitted | 3350.516 |
| full_slot_coverage / action completion | 3 | 2340.510 | 3124.350 | 2865.891 | 455.013 | omitted | omitted | 3132.813 |
| topology_converged / orchestration start | 3 | 2711.090 | 3340.834 | 3134.146 | 366.410 | omitted | omitted | 3350.516 |
| topology_converged / action completion | 3 | 2340.510 | 3124.350 | 2865.891 | 455.013 | omitted | omitted | 3132.813 |
| client_first_all_shard_success / orchestration start | 3 | 3125.271 | 3368.109 | 3290.263 | 142.963 | omitted | omitted | 3377.408 |
| client_first_all_shard_success / action completion | 3 | 2754.691 | 3151.626 | 3022.007 | 231.538 | omitted | omitted | 3159.706 |
| client_stable / orchestration start | 3 | 4239.585 | 4475.079 | 4397.653 | 136.901 | omitted | omitted | 4478.296 |
| client_stable / action completion | 3 | 3869.005 | 4257.376 | 4129.398 | 225.518 | omitted | omitted | 4261.813 |
| combined_stable / orchestration start | 3 | 4239.585 | 4475.079 | 4397.653 | 136.901 | omitted | omitted | 4478.296 |
| combined_stable / action completion | 3 | 3869.005 | 4257.376 | 4129.398 | 225.518 | omitted | omitted | 4261.813 |
| write_policy_stable / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| write_policy_stable / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

## Client operations

Completed request fractions weight actual request counts. Throughput divides summed completed/successful counts by summed trial phase wall-clock seconds; it does not imply workers achieved the configured offered rate. Zero-attempt phases have an unobserved success fraction.

| Phase | Trials with phase | Completed | Successful | Failed | Request success fraction | Completed ops/s | Successful ops/s |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 3 | 4799 | 4799 | 0 | 100.00% | 199.429 | 199.429 |
| fault | 3 | 1527 | 1279 | 248 | 83.76% | 115.742 | 96.945 |
| post_recovery | 3 | 1800 | 1800 | 0 | 100.00% | 199.970 | 199.970 |

Whole-workload schedule counters measured in 3 trials; legacy or incomplete counters unmeasured in 0. Offered: 9251; started: 8126; completed: 8126; dropped: 1125. Started fraction of offered: 87.84%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| baseline / completed ops/s | 3 | 198.356 | 199.953 | 199.430 | 0.930 | omitted | omitted | 199.980 |
| baseline / successful request p50 ms | 3 | 0.495 | 0.546 | 0.568 | 0.086 | omitted | omitted | 0.664 |
| baseline / successful request p95 ms | 3 | 0.768 | 0.807 | 1.481 | 1.202 | omitted | omitted | 2.869 |
| baseline / successful request p99 ms | 3 | 1.314 | 2.869 | 4.332 | 3.957 | omitted | omitted | 8.812 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 3 | 0.830 | 0.837 | 0.837 | 0.007 | omitted | omitted | 0.845 |
| fault / completed ops/s | 3 | 107.930 | 113.658 | 115.926 | 9.339 | omitted | omitted | 126.191 |
| fault / successful request p50 ms | 3 | 0.495 | 0.520 | 0.607 | 0.173 | omitted | omitted | 0.807 |
| fault / successful request p95 ms | 3 | 1.380 | 1.521 | 2.380 | 1.611 | omitted | omitted | 4.239 |
| fault / successful request p99 ms | 3 | 149.301 | 149.301 | 219.305 | 121.249 | omitted | omitted | 359.312 |
| fault / failed request p50 ms | 3 | 0.981 | 1.252 | 1.182 | 0.177 | omitted | omitted | 1.314 |
| fault / failed request p95 ms | 3 | 500.533 | 500.572 | 500.582 | 0.055 | omitted | omitted | 500.641 |
| fault / failed request p99 ms | 3 | 500.533 | 500.572 | 500.582 | 0.055 | omitted | omitted | 500.641 |
| post_recovery / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| post_recovery / completed ops/s | 3 | 199.945 | 199.978 | 199.970 | 0.023 | omitted | omitted | 199.989 |
| post_recovery / successful request p50 ms | 3 | 0.449 | 0.472 | 0.508 | 0.083 | omitted | omitted | 0.602 |
| post_recovery / successful request p95 ms | 3 | 0.632 | 0.664 | 1.048 | 0.694 | omitted | omitted | 1.849 |
| post_recovery / successful request p99 ms | 3 | 0.807 | 0.890 | 2.123 | 2.209 | omitted | omitted | 4.673 |
| post_recovery / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

No pooled request latency quantiles are computed: individual report schemas 1 and 2 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 3 trials, unmeasured in 0. Every successful SET remains in the main validation ledger even when WAIT is short, errors, or is not attempted. Stage command latency remains in individual reports; it does not increase logical request counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 2426 | 0 | 0 | 0 | 0 | 0 | 0 |
| fault | 624 | 132 | 0 | 0 | 0 | 0 | 0 |
| post_recovery | 965 | 0 | 0 | 0 | 0 | 0 | 0 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. Attempted WAIT client-call timeouts are counted separately above.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 31 |
| fault | network | 101 |
| fault | timeout | 116 |

## Acknowledged-write accounting

Validation attempted in 3 trials and complete in 3. SET attempts: 4147; acknowledged writes: 4015; ambiguous SET outcomes: 132.

Checked: 4015; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 3 | 1330.000 | 1342.000 | 1338.333 | 7.234 | omitted | omitted | 1343.000 |
| ambiguous writes | 3 | 41.000 | 42.000 | 44.000 | 4.359 | omitted | omitted | 49.000 |
| missing writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | 0207534b36545671d4c6bcb3 | recovered | true | true | [JSON](runs/0207534b36545671d4c6bcb3/report.json), [Markdown](runs/0207534b36545671d4c6bcb3/report.md) |
| 2 | fc9ec7fdcee8870aaa67e2ce | recovered | true | true | [JSON](runs/fc9ec7fdcee8870aaa67e2ce/report.json), [Markdown](runs/fc9ec7fdcee8870aaa67e2ce/report.md) |
| 3 | 461f750112eb15258eeb6892 | recovered | true | true | [JSON](runs/461f750112eb15258eeb6892/report.json), [Markdown](runs/461f750112eb15258eeb6892/report.md) |

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

Effective experiment configuration SHA-256: `4f909e9e757b339b59f088d76a570370528942c29f9612b1d1b964f9f0b11aea`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
