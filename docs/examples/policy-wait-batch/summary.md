# Redis Reliability Lab — repeated trials

Batch: `batch-26adf7924f9e85799ff28453`  
Outcome: **partial**  
Started: 2026-10-11T07:42:19.514247Z  
Finished: 2026-10-11T07:43:46.73678Z

Requested: 3; started: 3; callbacks completed: 3; unstarted: 0. Cleanup failures: 0; incomplete individual evidence: 0.

| Trial outcome | Count |
|---|---:|
| recovered | 0 |
| partial | 3 |
| failed | 0 |
| aborted | 0 |
| unstarted | 0 |

Recovery success fraction: 0.00% of **all 3 requested trials**. Started-trial recovery fraction: 0.00%.

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
| write_policy_stable | 0 | 3 | 0 | 0 | 0 | 0 |

Descriptive timing statistics use milliseconds. **P90/P95 are omitted for N<20**; small samples do not support reliable tail estimates. Sample standard deviation uses N-1 and is undefined for N<2. These are descriptive statistics, not confidence intervals.

| Timing / anchor | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| fault action duration | 3 | 201.039 | 207.776 | 206.491 | 4.937 | omitted | omitted | 210.659 |
| failure_detected / orchestration start | 3 | 1760.862 | 1762.094 | 1765.426 | 6.866 | omitted | omitted | 1773.322 |
| failure_detected / action completion | 3 | 1553.086 | 1561.055 | 1558.935 | 5.128 | omitted | omitted | 1562.663 |
| replica_promoted / orchestration start | 3 | 2532.826 | 2536.222 | 2541.677 | 12.506 | omitted | omitted | 2555.984 |
| replica_promoted / action completion | 3 | 2325.050 | 2335.183 | 2335.186 | 10.138 | omitted | omitted | 2345.325 |
| full_slot_coverage / orchestration start | 3 | 2536.222 | 2555.984 | 2800.390 | 440.549 | omitted | omitted | 3308.964 |
| full_slot_coverage / action completion | 3 | 2335.183 | 2345.325 | 2593.899 | 439.355 | omitted | omitted | 3101.188 |
| topology_converged / orchestration start | 3 | 2536.222 | 2555.984 | 2800.390 | 440.549 | omitted | omitted | 3308.964 |
| topology_converged / action completion | 3 | 2335.183 | 2345.325 | 2593.899 | 439.355 | omitted | omitted | 3101.188 |
| client_first_all_shard_success / orchestration start | 3 | 3330.535 | 3343.155 | 3381.788 | 78.099 | omitted | omitted | 3471.674 |
| client_first_all_shard_success / action completion | 3 | 3122.759 | 3142.116 | 3175.297 | 74.862 | omitted | omitted | 3261.015 |
| client_stable / orchestration start | 3 | 4474.964 | 4495.402 | 4530.125 | 78.510 | omitted | omitted | 4620.010 |
| client_stable / action completion | 3 | 4267.188 | 4294.363 | 4323.634 | 75.466 | omitted | omitted | 4409.351 |
| combined_stable / orchestration start | 3 | 4474.964 | 4495.402 | 4530.125 | 78.510 | omitted | omitted | 4620.010 |
| combined_stable / action completion | 3 | 4267.188 | 4294.363 | 4323.634 | 75.466 | omitted | omitted | 4409.351 |
| write_policy_stable / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| write_policy_stable / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

## Client operations

Completed request fractions weight actual request counts. Throughput divides summed completed/successful counts by summed trial phase wall-clock seconds; it does not imply workers achieved the configured offered rate. Zero-attempt phases have an unobserved success fraction.

| Phase | Trials with phase | Completed | Successful | Failed | Request success fraction | Completed ops/s | Successful ops/s |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 3 | 4809 | 4809 | 0 | 100.00% | 199.931 | 199.931 |
| fault | 3 | 1611 | 1210 | 401 | 75.11% | 115.606 | 86.830 |
| post_recovery | 3 | 1800 | 1478 | 322 | 82.11% | 192.131 | 157.761 |

Whole-workload schedule counters measured in 3 trials; legacy or incomplete counters unmeasured in 0. Offered: 9397; started: 8220; completed: 8220; dropped: 1177. Started fraction of offered: 87.47%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| baseline / completed ops/s | 3 | 199.898 | 199.929 | 199.931 | 0.034 | omitted | omitted | 199.966 |
| baseline / successful request p50 ms | 3 | 0.768 | 0.807 | 0.808 | 0.039 | omitted | omitted | 0.847 |
| baseline / successful request p95 ms | 3 | 1.192 | 1.314 | 1.295 | 0.095 | omitted | omitted | 1.380 |
| baseline / successful request p99 ms | 3 | 1.521 | 1.849 | 2.034 | 0.626 | omitted | omitted | 2.732 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 3 | 0.745 | 0.752 | 0.751 | 0.005 | omitted | omitted | 0.756 |
| fault / completed ops/s | 3 | 114.360 | 114.688 | 115.623 | 1.910 | omitted | omitted | 117.820 |
| fault / successful request p50 ms | 3 | 0.697 | 0.697 | 0.709 | 0.020 | omitted | omitted | 0.732 |
| fault / successful request p95 ms | 3 | 1.598 | 1.677 | 1.708 | 0.129 | omitted | omitted | 1.849 |
| fault / successful request p99 ms | 3 | 181.477 | 295.607 | 278.798 | 90.101 | omitted | omitted | 359.312 |
| fault / failed request p50 ms | 3 | 111.411 | 111.411 | 111.411 | 0.000 | omitted | omitted | 111.411 |
| fault / failed request p95 ms | 3 | 500.560 | 500.562 | 500.567 | 0.010 | omitted | omitted | 500.578 |
| fault / failed request p99 ms | 3 | 500.560 | 500.562 | 500.567 | 0.010 | omitted | omitted | 500.578 |
| post_recovery / trial success fraction | 3 | 0.818 | 0.818 | 0.821 | 0.005 | omitted | omitted | 0.827 |
| post_recovery / completed ops/s | 3 | 190.770 | 191.408 | 192.143 | 1.853 | omitted | omitted | 194.251 |
| post_recovery / successful request p50 ms | 3 | 0.546 | 0.602 | 0.593 | 0.044 | omitted | omitted | 0.632 |
| post_recovery / successful request p95 ms | 3 | 1.081 | 1.135 | 1.177 | 0.122 | omitted | omitted | 1.314 |
| post_recovery / successful request p99 ms | 3 | 1.380 | 1.677 | 1.975 | 0.788 | omitted | omitted | 2.869 |
| post_recovery / failed request p50 ms | 3 | 111.411 | 111.411 | 111.411 | 0.000 | omitted | omitted | 111.411 |
| post_recovery / failed request p95 ms | 3 | 128.972 | 135.421 | 135.528 | 6.610 | omitted | omitted | 142.192 |
| post_recovery / failed request p99 ms | 3 | 142.192 | 156.767 | 154.521 | 11.374 | omitted | omitted | 164.605 |

No pooled request latency quantiles are computed: individual report schemas 1 and 2 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 3 trials, unmeasured in 0. Every successful SET remains in the main validation ledger even when WAIT is short, errors, or is not attempted. Stage command latency remains in individual reports; it does not increase logical request counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 2430 | 0 | 2430 | 2430 | 0 | 0 | 0 |
| fault | 673 | 124 | 673 | 513 | 160 | 0 | 0 |
| post_recovery | 967 | 0 | 967 | 645 | 322 | 0 | 0 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. Attempted WAIT client-call timeouts are counted separately above.

WAIT-satisfied subset validation: expected 3588; checked 3588; missing 0; corrupt 0; unreadable 0; unverified 0; complete true.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 32 |
| fault | network | 92 |
| fault | timeout | 117 |
| fault | wait_short | 160 |
| post_recovery | wait_short | 322 |

## Acknowledged-write accounting

Validation attempted in 3 trials and complete in 3. SET attempts: 4194; acknowledged writes: 4070; ambiguous SET outcomes: 124.

Checked: 4070; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 3 | 1355.000 | 1357.000 | 1356.667 | 1.528 | omitted | omitted | 1358.000 |
| ambiguous writes | 3 | 37.000 | 42.000 | 41.333 | 4.041 | omitted | omitted | 45.000 |
| missing writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | 92d758d529f45b4bb1836474 | partial | true | true | [JSON](runs/92d758d529f45b4bb1836474/report.json), [Markdown](runs/92d758d529f45b4bb1836474/report.md) |
| 2 | 2f20da5016af35f23348d0de | partial | true | true | [JSON](runs/2f20da5016af35f23348d0de/report.json), [Markdown](runs/2f20da5016af35f23348d0de/report.md) |
| 3 | 07aadc068ea2995602d6959d | partial | true | true | [JSON](runs/07aadc068ea2995602d6959d/report.json), [Markdown](runs/07aadc068ea2995602d6959d/report.md) |

## Execution errors

- trial 1: experiment outcome: partial; trial 2: experiment outcome: partial; trial 3: experiment outcome: partial

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

Effective experiment configuration SHA-256: `e2a6d0c505f6d1bac0abfa27b1e134e13352da9e819b7560d276fe8d7080b863`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
