# Redis Reliability Lab — repeated trials

Batch: `comparison-321638985ff49f9040049bf4-condition-02`  
Outcome: **recovered**  
Started: 2026-10-11T09:11:38.844832Z  
Finished: 2026-10-11T09:17:22.385436Z

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
| write_policy_stable | 3 | 0 | 0 | 0 | 0 | 0 |
| replication_interrupted | 3 | 0 | 0 | 0 | 0 | 0 |
| target_reachable_after_repair | 3 | 0 | 0 | 0 | 0 | 0 |
| membership_rejoined | 3 | 0 | 0 | 0 | 0 | 0 |
| replica_link_up | 3 | 0 | 0 | 0 | 0 | 0 |
| replica_sync_complete | 3 | 0 | 0 | 0 | 0 | 0 |
| stable_redundancy | 3 | 0 | 0 | 0 | 0 | 0 |
| original_topology_stable | 0 | 0 | 0 | 3 | 0 | 0 |
| persistence_recovered | 0 | 0 | 0 | 3 | 0 | 0 |
| all_processes_stopped | 0 | 0 | 0 | 3 | 0 | 0 |

Descriptive timing statistics use milliseconds. **P90/P95 are omitted for N<20**; small samples do not support reliable tail estimates. Sample standard deviation uses N-1 and is undefined for N<2. These are descriptive statistics, not confidence intervals.

| Timing / anchor | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| fault action duration | 3 | 235.555 | 236.903 | 242.432 | 10.766 | omitted | omitted | 254.839 |
| failure_detected / orchestration start | 3 | 1797.670 | 1825.250 | 1819.461 | 19.551 | omitted | omitted | 1835.464 |
| failure_detected / action completion | 3 | 1562.115 | 1580.625 | 1577.029 | 13.481 | omitted | omitted | 1588.348 |
| replica_promoted / orchestration start | 3 | 2572.051 | 2615.063 | 2604.097 | 28.210 | omitted | omitted | 2625.178 |
| replica_promoted / action completion | 3 | 2336.497 | 2370.339 | 2361.665 | 22.145 | omitted | omitted | 2378.161 |
| full_slot_coverage / orchestration start | 3 | 2572.051 | 2625.178 | 2902.730 | 527.413 | omitted | omitted | 3510.961 |
| full_slot_coverage / action completion | 3 | 2336.497 | 2370.339 | 2660.298 | 531.801 | omitted | omitted | 3274.058 |
| topology_converged / orchestration start | 3 | 2572.051 | 2625.178 | 2902.730 | 527.413 | omitted | omitted | 3510.961 |
| topology_converged / action completion | 3 | 2336.497 | 2370.339 | 2660.298 | 531.801 | omitted | omitted | 3274.058 |
| client_first_all_shard_success / orchestration start | 3 | 3035.827 | 3128.115 | 3231.844 | 263.656 | omitted | omitted | 3531.589 |
| client_first_all_shard_success / action completion | 3 | 2800.273 | 2891.212 | 2989.412 | 252.962 | omitted | omitted | 3276.750 |
| client_stable / orchestration start | 3 | 4230.164 | 4333.921 | 4427.034 | 256.435 | omitted | omitted | 4717.017 |
| client_stable / action completion | 3 | 3994.609 | 4097.019 | 4184.602 | 245.781 | omitted | omitted | 4462.178 |
| combined_stable / orchestration start | 3 | 4230.164 | 4717.017 | 4558.631 | 284.522 | omitted | omitted | 4728.713 |
| combined_stable / action completion | 3 | 3994.609 | 4462.178 | 4316.199 | 278.899 | omitted | omitted | 4491.811 |
| write_policy_stable / orchestration start | 3 | 10212.858 | 11080.822 | 11089.841 | 881.528 | omitted | omitted | 11975.844 |
| write_policy_stable / action completion | 3 | 9977.303 | 10843.919 | 10847.409 | 871.856 | omitted | omitted | 11721.005 |
| replication_interrupted / orchestration start | 3 | 2572.051 | 2615.063 | 2604.097 | 28.210 | omitted | omitted | 2625.178 |
| replication_interrupted / action completion | 3 | 2336.497 | 2370.339 | 2361.665 | 22.145 | omitted | omitted | 2378.161 |
| target_reachable_after_repair / orchestration start | 3 | 5882.436 | 6407.196 | 6245.066 | 314.632 | omitted | omitted | 6445.564 |
| target_reachable_after_repair / action completion | 3 | 5646.882 | 6170.294 | 6002.634 | 308.259 | omitted | omitted | 6190.725 |
| membership_rejoined / orchestration start | 3 | 5882.436 | 6407.196 | 6245.066 | 314.632 | omitted | omitted | 6445.564 |
| membership_rejoined / action completion | 3 | 5646.882 | 6170.294 | 6002.634 | 308.259 | omitted | omitted | 6190.725 |
| replica_link_up / orchestration start | 3 | 9247.574 | 10118.887 | 10022.315 | 731.254 | omitted | omitted | 10700.484 |
| replica_link_up / action completion | 3 | 9012.019 | 9881.984 | 9779.883 | 722.246 | omitted | omitted | 10445.645 |
| replica_sync_complete / orchestration start | 3 | 9247.574 | 10118.887 | 10022.315 | 731.254 | omitted | omitted | 10700.484 |
| replica_sync_complete / action completion | 3 | 9012.019 | 9881.984 | 9779.883 | 722.246 | omitted | omitted | 10445.645 |
| stable_redundancy / orchestration start | 3 | 10523.834 | 11388.004 | 11294.763 | 728.795 | omitted | omitted | 11972.450 |
| stable_redundancy / action completion | 3 | 10288.280 | 11151.102 | 11052.331 | 719.766 | omitted | omitted | 11717.611 |
| original_topology_stable / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| original_topology_stable / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| persistence_recovered / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| persistence_recovered / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| all_processes_stopped / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| all_processes_stopped / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

## Client operations

Completed request fractions weight actual request counts. Throughput divides summed completed/successful counts by summed trial phase wall-clock seconds; it does not imply workers achieved the configured offered rate. Zero-attempt phases have an unobserved success fraction.

| Phase | Trials with phase | Completed | Successful | Failed | Request success fraction | Completed ops/s | Successful ops/s |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 3 | 2410 | 2410 | 0 | 100.00% | 199.848 | 199.848 |
| fault | 3 | 1679 | 1290 | 389 | 76.83% | 119.434 | 91.763 |
| post_recovery | 3 | 1201 | 1201 | 0 | 100.00% | 200.144 | 200.144 |
| reconciliation | 3 | 3054 | 2699 | 355 | 88.38% | 200.028 | 176.777 |
| repair | 3 | 932 | 775 | 157 | 83.15% | 199.637 | 166.007 |

Whole-workload schedule counters measured in 3 trials; legacy or incomplete counters unmeasured in 0. Offered: 10410; started: 9276; completed: 9276; dropped: 1134. Started fraction of offered: 89.11%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| baseline / completed ops/s | 3 | 199.775 | 199.841 | 199.848 | 0.077 | omitted | omitted | 199.929 |
| baseline / successful request p50 ms | 3 | 1.192 | 1.252 | 1.232 | 0.034 | omitted | omitted | 1.252 |
| baseline / successful request p95 ms | 3 | 2.732 | 2.732 | 2.732 | 0.000 | omitted | omitted | 2.732 |
| baseline / successful request p99 ms | 3 | 3.163 | 3.321 | 3.719 | 0.830 | omitted | omitted | 4.673 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 3 | 0.765 | 0.768 | 0.768 | 0.003 | omitted | omitted | 0.771 |
| fault / completed ops/s | 3 | 113.778 | 118.882 | 119.231 | 5.636 | omitted | omitted | 125.034 |
| fault / successful request p50 ms | 3 | 0.546 | 0.664 | 0.659 | 0.111 | omitted | omitted | 0.768 |
| fault / successful request p95 ms | 3 | 3.163 | 3.321 | 3.443 | 0.357 | omitted | omitted | 3.845 |
| fault / successful request p99 ms | 3 | 200.078 | 436.746 | 357.856 | 136.640 | omitted | omitted | 436.746 |
| fault / failed request p50 ms | 3 | 111.411 | 111.411 | 111.411 | 0.000 | omitted | omitted | 111.411 |
| fault / failed request p95 ms | 3 | 500.720 | 500.792 | 500.886 | 0.227 | omitted | omitted | 501.145 |
| fault / failed request p99 ms | 3 | 500.720 | 500.792 | 500.886 | 0.227 | omitted | omitted | 501.145 |
| post_recovery / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| post_recovery / completed ops/s | 3 | 199.980 | 199.990 | 200.144 | 0.276 | omitted | omitted | 200.462 |
| post_recovery / successful request p50 ms | 3 | 0.520 | 0.546 | 0.753 | 0.381 | omitted | omitted | 1.192 |
| post_recovery / successful request p95 ms | 3 | 2.360 | 2.732 | 2.702 | 0.327 | omitted | omitted | 3.012 |
| post_recovery / successful request p99 ms | 3 | 2.869 | 3.321 | 4.151 | 1.843 | omitted | omitted | 6.263 |
| post_recovery / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / trial success fraction | 3 | 0.875 | 0.886 | 0.884 | 0.008 | omitted | omitted | 0.891 |
| reconciliation / completed ops/s | 3 | 199.901 | 200.043 | 200.022 | 0.111 | omitted | omitted | 200.121 |
| reconciliation / successful request p50 ms | 3 | 0.520 | 0.520 | 0.529 | 0.015 | omitted | omitted | 0.546 |
| reconciliation / successful request p95 ms | 3 | 2.478 | 2.732 | 2.648 | 0.147 | omitted | omitted | 2.732 |
| reconciliation / successful request p99 ms | 3 | 4.037 | 4.239 | 4.242 | 0.207 | omitted | omitted | 4.451 |
| reconciliation / failed request p50 ms | 3 | 111.411 | 111.411 | 111.411 | 0.000 | omitted | omitted | 111.411 |
| reconciliation / failed request p95 ms | 3 | 122.831 | 128.972 | 126.925 | 3.546 | omitted | omitted | 128.972 |
| reconciliation / failed request p99 ms | 3 | 128.972 | 137.102 | 137.386 | 8.560 | omitted | omitted | 146.084 |
| repair / trial success fraction | 3 | 0.825 | 0.826 | 0.832 | 0.010 | omitted | omitted | 0.844 |
| repair / completed ops/s | 3 | 199.450 | 199.699 | 199.639 | 0.168 | omitted | omitted | 199.769 |
| repair / successful request p50 ms | 3 | 0.449 | 0.520 | 0.505 | 0.050 | omitted | omitted | 0.546 |
| repair / successful request p95 ms | 3 | 2.360 | 2.602 | 2.565 | 0.189 | omitted | omitted | 2.732 |
| repair / successful request p99 ms | 3 | 3.321 | 3.321 | 3.560 | 0.413 | omitted | omitted | 4.037 |
| repair / failed request p50 ms | 3 | 111.411 | 111.411 | 111.411 | 0.000 | omitted | omitted | 111.411 |
| repair / failed request p95 ms | 3 | 128.972 | 128.972 | 131.122 | 3.723 | omitted | omitted | 135.421 |
| repair / failed request p99 ms | 3 | 135.104 | 135.591 | 135.503 | 0.363 | omitted | omitted | 135.814 |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 3 trials, unmeasured in 0. Every successful SET remains in the main validation ledger when the later acknowledgement command is short, errors, or is not attempted. Diagnostic SET/WAIT/WAITAOF stage latency remains in individual reports and does not increase logical request or pacing counts. Errors and unattempted commands supply unknown acknowledgement counts, never fabricated zero replies.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 1250 | 0 | 1250 | 1250 | 0 | 0 | 0 |
| fault | 678 | 124 | 678 | 522 | 156 | 0 | 0 |
| post_recovery | 587 | 0 | 587 | 587 | 0 | 0 | 0 |
| reconciliation | 1587 | 0 | 1587 | 1232 | 355 | 0 | 0 |
| repair | 473 | 0 | 473 | 316 | 157 | 0 | 0 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. Attempted WAIT client-call timeouts are counted separately above.

| Observed WAIT replica count | Replies |
|---:|---:|
| 0 | 668 |
| 1 | 3907 |

WAIT-satisfied subset validation: expected 3907; checked 3907; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Validation attempted in 3 trials. Complete means every subset record was checked; preservation also requires zero missing, corrupt and unreadable records.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 30 |
| fault | network | 94 |
| fault | timeout | 109 |
| fault | wait_short | 156 |
| reconciliation | wait_short | 355 |
| repair | wait_short | 157 |

## Acknowledged-write accounting

Validation attempted in 3 trials and complete in 3. SET attempts: 4699; acknowledged writes: 4575; ambiguous SET outcomes: 124.

Checked: 4575; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 3 | 1462.000 | 1530.000 | 1525.000 | 60.655 | omitted | omitted | 1583.000 |
| ambiguous writes | 3 | 35.000 | 42.000 | 41.333 | 6.028 | omitted | omitted | 47.000 |
| missing writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | dc0464195769627e90802e83 | recovered | true | true | [JSON](../../trials/position-0001/batch-28685cb04eeec1fdcdf25b49/runs/dc0464195769627e90802e83/report.json), [Markdown](../../trials/position-0001/batch-28685cb04eeec1fdcdf25b49/runs/dc0464195769627e90802e83/report.md) |
| 2 | a41f7d40e0b182eec96b39bb | recovered | true | true | [JSON](../../trials/position-0006/batch-a4333dacfdfe4f1ff20e1e7a/runs/a41f7d40e0b182eec96b39bb/report.json), [Markdown](../../trials/position-0006/batch-a4333dacfdfe4f1ff20e1e7a/runs/a41f7d40e0b182eec96b39bb/report.md) |
| 3 | 73a8d11d94ba9a7f56fb11b1 | recovered | true | true | [JSON](../../trials/position-0008/batch-4c4de00323adeee2bc70b1a9/runs/73a8d11d94ba9a7f56fb11b1/report.json), [Markdown](../../trials/position-0008/batch-4c4de00323adeee2bc70b1a9/runs/73a8d11d94ba9a7f56fb11b1/report.md) |

## Measurement methodology

- A seeded permutation of condition indices is rotated by one position each round. Within each complete block of N rounds, every condition occupies every round position once. An incomplete block is only partially balanced; this rotating Latin order does not balance all predecessor/successor pairs.
- The order seed controls only the condition schedule. Effective workload seeds and all other configuration values remain unchanged. Each entry executes one fresh sequential batch trial with persisted-report integrity and cleanup checks.
- Further trials require verified individual evidence and successful cleanup, or verified evidence that no environment was provisioned. Unsafe cleanup, missing evidence, cancellation and summary-write failures stop execution; all remaining planned entries stay explicitly unstarted.
- Failed and partial trials with valid evidence and safe cleanup remain in their condition statistics. Timing statistics are conditional on observed milestones; missing and censored observations are not zero-duration events. Request latency quantiles are not pooled across trials or conditions.
- Condition summaries reuse batch aggregation. Condition-summary report paths are relative to their own summary directory; schedule report paths and batch references are relative to the comparison directory. Host and environment references point to individual JSON report sections.
- Ordering reduces systematic position confounding but does not prove causality, independent samples, equal host conditions or population reliability. Compare individual environmental metadata and actual order before interpreting differences.

Effective experiment configuration SHA-256: `d7e106ea5cb6011eb4736e886b40e76b334dd77ec754a2cae619035e090e1ef9`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
