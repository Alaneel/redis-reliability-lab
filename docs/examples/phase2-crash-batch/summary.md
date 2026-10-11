# Redis Reliability Lab — repeated trials

Batch: `batch-e4b81963e956ed40e22b4a27`  
Outcome: **recovered**  
Started: 2026-10-11T07:15:35.336483Z  
Finished: 2026-10-11T07:17:15.963626Z

Requested: 5; started: 5; callbacks completed: 5; unstarted: 0. Cleanup failures: 0; incomplete individual evidence: 0.

| Trial outcome | Count |
|---|---:|
| recovered | 5 |
| partial | 0 |
| failed | 0 |
| aborted | 0 |
| unstarted | 0 |

Recovery success fraction: 100.00% of **all 5 requested trials**. Started-trial recovery fraction: 100.00%.

## Recovery observations

Timing statistics below are conditional on observed milestones, including observed milestones in partial or failed trials. Unobserved events are not zero-duration recoveries. Censoring means unobserved by the accepted sampled horizon; it does not prove an event physically occurred after that horizon. No-requested-fault trials are not at risk. Aborted/error horizons are separate from deadline horizons in JSON.

| Event | Observed | Censored | Deadline censored | Not at risk | Missing | Unstarted |
|---|---:|---:|---:|---:|---:|---:|
| failure_detected | 5 | 0 | 0 | 0 | 0 | 0 |
| replica_promoted | 5 | 0 | 0 | 0 | 0 | 0 |
| full_slot_coverage | 5 | 0 | 0 | 0 | 0 | 0 |
| topology_converged | 5 | 0 | 0 | 0 | 0 | 0 |
| client_first_all_shard_success | 5 | 0 | 0 | 0 | 0 | 0 |
| client_stable | 5 | 0 | 0 | 0 | 0 | 0 |
| combined_stable | 5 | 0 | 0 | 0 | 0 | 0 |

Descriptive timing statistics use milliseconds. **P90/P95 are omitted for N<20**; small samples do not support reliable tail estimates. Sample standard deviation uses N-1 and is undefined for N<2. These are descriptive statistics, not confidence intervals.

| Timing / anchor | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| fault action duration | 5 | 199.701 | 206.691 | 208.304 | 6.703 | omitted | omitted | 217.500 |
| failure_detected / orchestration start | 5 | 1754.993 | 1774.005 | 1768.292 | 12.241 | omitted | omitted | 1779.714 |
| failure_detected / action completion | 5 | 1549.338 | 1562.214 | 1559.988 | 7.545 | omitted | omitted | 1567.314 |
| replica_promoted / orchestration start | 5 | 2538.939 | 2558.434 | 2705.453 | 339.388 | omitted | omitted | 3312.375 |
| replica_promoted / action completion | 5 | 2333.127 | 2347.920 | 2497.149 | 344.163 | omitted | omitted | 3112.673 |
| full_slot_coverage / orchestration start | 5 | 2538.939 | 2558.434 | 2705.453 | 339.388 | omitted | omitted | 3312.375 |
| full_slot_coverage / action completion | 5 | 2333.127 | 2347.920 | 2497.149 | 344.163 | omitted | omitted | 3112.673 |
| topology_converged / orchestration start | 5 | 2538.939 | 2558.434 | 2705.453 | 339.388 | omitted | omitted | 3312.375 |
| topology_converged / action completion | 5 | 2333.127 | 2347.920 | 2497.149 | 344.163 | omitted | omitted | 3112.673 |
| client_first_all_shard_success / orchestration start | 5 | 3058.486 | 3353.813 | 3243.029 | 165.055 | omitted | omitted | 3369.663 |
| client_first_all_shard_success / action completion | 5 | 2858.784 | 3136.313 | 3034.724 | 160.153 | omitted | omitted | 3162.973 |
| client_stable / orchestration start | 5 | 4172.748 | 4441.869 | 4347.898 | 159.720 | omitted | omitted | 4482.438 |
| client_stable / action completion | 5 | 3966.936 | 4224.369 | 4139.593 | 155.104 | omitted | omitted | 4275.747 |
| combined_stable / orchestration start | 5 | 4172.748 | 4454.121 | 4403.807 | 130.052 | omitted | omitted | 4482.438 |
| combined_stable / action completion | 5 | 3966.936 | 4254.419 | 4195.503 | 129.085 | omitted | omitted | 4275.747 |

## Client operations

Completed request fractions weight actual request counts. Throughput divides summed completed/successful counts by summed trial phase wall-clock seconds; it does not imply workers achieved the configured offered rate. Zero-attempt phases have an unobserved success fraction.

| Phase | Trials with phase | Completed | Successful | Failed | Request success fraction | Completed ops/s | Successful ops/s |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 5 | 2019 | 2019 | 0 | 100.00% | 199.838 | 199.838 |
| fault | 5 | 2052 | 1845 | 207 | 89.91% | 93.191 | 83.790 |
| post_recovery | 5 | 1000 | 1000 | 0 | 100.00% | 199.927 | 199.927 |

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 5 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| baseline / completed ops/s | 5 | 199.604 | 199.868 | 199.838 | 0.136 | omitted | omitted | 199.937 |
| baseline / successful request p50 ms | 5 | 0.428 | 0.495 | 0.492 | 0.038 | omitted | omitted | 0.520 |
| baseline / successful request p95 ms | 5 | 0.573 | 0.697 | 0.694 | 0.073 | omitted | omitted | 0.768 |
| baseline / successful request p99 ms | 5 | 0.732 | 1.252 | 1.223 | 0.436 | omitted | omitted | 1.761 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 5 | 0.880 | 0.901 | 0.898 | 0.012 | omitted | omitted | 0.909 |
| fault / completed ops/s | 5 | 83.877 | 95.455 | 93.074 | 5.488 | omitted | omitted | 97.661 |
| fault / successful request p50 ms | 5 | 0.428 | 0.520 | 0.492 | 0.042 | omitted | omitted | 0.520 |
| fault / successful request p95 ms | 5 | 0.981 | 1.030 | 1.151 | 0.217 | omitted | omitted | 1.449 |
| fault / successful request p99 ms | 5 | 255.356 | 396.141 | 353.937 | 62.943 | omitted | omitted | 396.141 |
| fault / failed request p50 ms | 5 | 500.456 | 500.566 | 500.605 | 0.139 | omitted | omitted | 500.832 |
| fault / failed request p95 ms | 5 | 500.456 | 500.566 | 500.605 | 0.139 | omitted | omitted | 500.832 |
| fault / failed request p99 ms | 5 | 500.456 | 500.566 | 500.605 | 0.139 | omitted | omitted | 500.832 |
| post_recovery / trial success fraction | 5 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| post_recovery / completed ops/s | 5 | 199.903 | 199.923 | 199.927 | 0.028 | omitted | omitted | 199.974 |
| post_recovery / successful request p50 ms | 5 | 0.449 | 0.472 | 0.472 | 0.016 | omitted | omitted | 0.495 |
| post_recovery / successful request p95 ms | 5 | 0.602 | 0.632 | 0.646 | 0.054 | omitted | omitted | 0.732 |
| post_recovery / successful request p99 ms | 5 | 0.697 | 0.732 | 0.835 | 0.166 | omitted | omitted | 1.081 |
| post_recovery / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

No pooled request latency quantiles are computed: individual report schemas 1 and 2 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | timeout | 207 |

## Acknowledged-write accounting

Validation attempted in 5 trials and complete in 5. SET attempts: 2564; acknowledged writes: 2459; ambiguous SET outcomes: 105.

Checked: 2459; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 5 | 466.000 | 496.000 | 491.800 | 16.006 | omitted | omitted | 506.000 |
| ambiguous writes | 5 | 19.000 | 21.000 | 21.000 | 1.581 | omitted | omitted | 23.000 |
| missing writes / validation attempted | 5 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 5 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 5 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | c71517bdf2d2fd24649843c9 | recovered | true | true | [JSON](runs/c71517bdf2d2fd24649843c9/report.json), [Markdown](runs/c71517bdf2d2fd24649843c9/report.md) |
| 2 | 447f02dcf25ca6d8a1931941 | recovered | true | true | [JSON](runs/447f02dcf25ca6d8a1931941/report.json), [Markdown](runs/447f02dcf25ca6d8a1931941/report.md) |
| 3 | 4052b45ef729eee33f8ba3f8 | recovered | true | true | [JSON](runs/4052b45ef729eee33f8ba3f8/report.json), [Markdown](runs/4052b45ef729eee33f8ba3f8/report.md) |
| 4 | 3a9679dd0e8a675a20d49be6 | recovered | true | true | [JSON](runs/3a9679dd0e8a675a20d49be6/report.json), [Markdown](runs/3a9679dd0e8a675a20d49be6/report.md) |
| 5 | 9c81a869a4e04108f3c73780 | recovered | true | true | [JSON](runs/9c81a869a4e04108f3c73780/report.json), [Markdown](runs/9c81a869a4e04108f3c73780/report.md) |

## Measurement methodology

- Trials execute sequentially with a fresh environment per callback; the next callback starts only after the previous callback returns with successful cleanup evidence (or an authoritative report that no environment was provisioned).
- Recovery success fraction uses all requested trials as the denominator. Outcomes, started/unstarted counts and individual reports preserve failed, partial and aborted trials; the started-trial fraction is reported separately.
- Recovery timings are conditional on observed milestones, including observed milestones in partial/failed trials. Unobserved post-fault events are censored at the accepted sampled observation horizon with deadline/abort/error reasons, never zero-duration samples; this does not prove the physical event occurred after that horizon. Trials without a requested fault are separately not at risk. No survival distribution, independence assumption or population reliability guarantee is inferred.
- Event durations use runtime SinceFaultMS and fault_action_duration_ms when present, preserving monotonic elapsed-time measurements. Schema 1 action-completion-relative durations fall back to wall-clock timestamp differences and can be affected by clock adjustments. First sampled observations include polling uncertainty; action duration includes ownership checks and fault command overhead.
- Min, median, mean, N-1 sample standard deviation and maximum describe available trial observations. P90/P95 use nearest rank only when N>=20; smaller samples have an explicit tail-estimation caveat.
- Completed-request success fractions weight actual completed counts; aggregate throughput divides summed completed/successful counts by summed phase wall-clock duration. Offered load and worker saturation are not inferred from completed requests.
- Acknowledged SET writes remain distinct from ambiguous SET outcomes and final read-back validation. Skipped/incomplete validation contributes acknowledged-not-checked counts; no observed loss is not a durability guarantee.

Effective experiment configuration SHA-256: `8e959199c0f2dff505399e403a8b7f25bfe40c39ec4dbc9ab905198f7d8d36c3`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
