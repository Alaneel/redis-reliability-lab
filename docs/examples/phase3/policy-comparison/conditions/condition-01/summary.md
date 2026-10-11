# Redis Reliability Lab — repeated trials

Batch: `comparison-321638985ff49f9040049bf4-condition-01`  
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
| write_policy_stable | 0 | 0 | 0 | 3 | 0 | 0 |
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
| fault action duration | 3 | 233.343 | 236.912 | 237.679 | 4.766 | omitted | omitted | 242.783 |
| failure_detected / orchestration start | 3 | 1804.405 | 1807.338 | 1808.536 | 4.843 | omitted | omitted | 1813.866 |
| failure_detected / action completion | 3 | 1567.493 | 1571.084 | 1570.857 | 3.256 | omitted | omitted | 1573.994 |
| replica_promoted / orchestration start | 3 | 2588.466 | 2590.549 | 2854.013 | 458.139 | omitted | omitted | 3383.025 |
| replica_promoted / action completion | 3 | 2351.554 | 2357.205 | 2616.334 | 453.727 | omitted | omitted | 3140.242 |
| full_slot_coverage / orchestration start | 3 | 2588.466 | 2590.549 | 2854.013 | 458.139 | omitted | omitted | 3383.025 |
| full_slot_coverage / action completion | 3 | 2351.554 | 2357.205 | 2616.334 | 453.727 | omitted | omitted | 3140.242 |
| topology_converged / orchestration start | 3 | 2588.466 | 2590.549 | 2854.013 | 458.139 | omitted | omitted | 3383.025 |
| topology_converged / action completion | 3 | 2351.554 | 2357.205 | 2616.334 | 453.727 | omitted | omitted | 3140.242 |
| client_first_all_shard_success / orchestration start | 3 | 3398.085 | 3405.680 | 3403.340 | 4.560 | omitted | omitted | 3406.256 |
| client_first_all_shard_success / action completion | 3 | 3161.173 | 3162.897 | 3165.661 | 6.339 | omitted | omitted | 3172.912 |
| client_stable / orchestration start | 3 | 4522.179 | 4523.550 | 4524.233 | 2.468 | omitted | omitted | 4526.970 |
| client_stable / action completion | 3 | 4280.768 | 4285.267 | 4286.554 | 6.525 | omitted | omitted | 4293.627 |
| combined_stable / orchestration start | 3 | 4522.179 | 4523.550 | 4524.233 | 2.468 | omitted | omitted | 4526.970 |
| combined_stable / action completion | 3 | 4280.768 | 4285.267 | 4286.554 | 6.525 | omitted | omitted | 4293.627 |
| write_policy_stable / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| write_policy_stable / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replication_interrupted / orchestration start | 3 | 2588.466 | 2590.549 | 2854.013 | 458.139 | omitted | omitted | 3383.025 |
| replication_interrupted / action completion | 3 | 2351.554 | 2357.205 | 2616.334 | 453.727 | omitted | omitted | 3140.242 |
| target_reachable_after_repair / orchestration start | 3 | 6038.904 | 6039.703 | 6060.829 | 37.285 | omitted | omitted | 6103.880 |
| target_reachable_after_repair / action completion | 3 | 5796.121 | 5802.791 | 5823.150 | 41.173 | omitted | omitted | 5870.537 |
| membership_rejoined / orchestration start | 3 | 6038.904 | 6039.703 | 6060.829 | 37.285 | omitted | omitted | 6103.880 |
| membership_rejoined / action completion | 3 | 5796.121 | 5802.791 | 5823.150 | 41.173 | omitted | omitted | 5870.537 |
| replica_link_up / orchestration start | 3 | 9549.942 | 10105.573 | 10129.698 | 592.187 | omitted | omitted | 10733.578 |
| replica_link_up / action completion | 3 | 9316.599 | 9862.790 | 9892.018 | 590.576 | omitted | omitted | 10496.666 |
| replica_sync_complete / orchestration start | 3 | 9549.942 | 10105.573 | 10129.698 | 592.187 | omitted | omitted | 10733.578 |
| replica_sync_complete / action completion | 3 | 9316.599 | 9862.790 | 9892.018 | 590.576 | omitted | omitted | 10496.666 |
| stable_redundancy / orchestration start | 3 | 10826.177 | 11386.981 | 11408.768 | 593.785 | omitted | omitted | 12013.147 |
| stable_redundancy / action completion | 3 | 10592.833 | 11144.198 | 11171.089 | 592.159 | omitted | omitted | 11776.236 |
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
| baseline | 3 | 2410 | 2410 | 0 | 100.00% | 199.815 | 199.815 |
| fault | 3 | 1626 | 1388 | 238 | 85.36% | 119.706 | 102.185 |
| post_recovery | 3 | 1200 | 1200 | 0 | 100.00% | 199.968 | 199.968 |
| reconciliation | 3 | 3231 | 3231 | 0 | 100.00% | 199.955 | 199.955 |
| repair | 3 | 918 | 918 | 0 | 100.00% | 200.005 | 200.005 |

Whole-workload schedule counters measured in 3 trials; legacy or incomplete counters unmeasured in 0. Offered: 10477; started: 9385; completed: 9385; dropped: 1092. Started fraction of offered: 89.58%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| baseline / completed ops/s | 3 | 199.733 | 199.764 | 199.815 | 0.116 | omitted | omitted | 199.947 |
| baseline / successful request p50 ms | 3 | 0.664 | 0.807 | 0.759 | 0.083 | omitted | omitted | 0.807 |
| baseline / successful request p95 ms | 3 | 1.252 | 1.380 | 1.337 | 0.074 | omitted | omitted | 1.380 |
| baseline / successful request p99 ms | 3 | 2.248 | 2.248 | 2.285 | 0.065 | omitted | omitted | 2.360 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 3 | 0.842 | 0.856 | 0.854 | 0.011 | omitted | omitted | 0.864 |
| fault / completed ops/s | 3 | 111.912 | 119.974 | 119.708 | 7.667 | omitted | omitted | 127.239 |
| fault / successful request p50 ms | 3 | 0.697 | 0.732 | 0.720 | 0.020 | omitted | omitted | 0.732 |
| fault / successful request p95 ms | 3 | 1.942 | 2.360 | 2.260 | 0.282 | omitted | omitted | 2.478 |
| fault / successful request p99 ms | 3 | 135.421 | 149.301 | 203.543 | 106.197 | omitted | omitted | 325.906 |
| fault / failed request p50 ms | 3 | 0.732 | 0.732 | 167.382 | 288.647 | omitted | omitted | 500.684 |
| fault / failed request p95 ms | 3 | 500.539 | 500.684 | 500.781 | 0.301 | omitted | omitted | 501.118 |
| fault / failed request p99 ms | 3 | 500.539 | 500.684 | 500.781 | 0.301 | omitted | omitted | 501.118 |
| post_recovery / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| post_recovery / completed ops/s | 3 | 199.944 | 199.976 | 199.968 | 0.020 | omitted | omitted | 199.983 |
| post_recovery / successful request p50 ms | 3 | 0.573 | 0.697 | 0.656 | 0.071 | omitted | omitted | 0.697 |
| post_recovery / successful request p95 ms | 3 | 1.252 | 1.252 | 1.514 | 0.454 | omitted | omitted | 2.039 |
| post_recovery / successful request p99 ms | 3 | 2.141 | 2.360 | 2.327 | 0.171 | omitted | omitted | 2.478 |
| post_recovery / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| reconciliation / completed ops/s | 3 | 199.954 | 199.954 | 199.955 | 0.001 | omitted | omitted | 199.956 |
| reconciliation / successful request p50 ms | 3 | 0.632 | 0.632 | 0.632 | 0.000 | omitted | omitted | 0.632 |
| reconciliation / successful request p95 ms | 3 | 1.192 | 1.192 | 1.212 | 0.034 | omitted | omitted | 1.252 |
| reconciliation / successful request p99 ms | 3 | 1.677 | 2.141 | 2.022 | 0.303 | omitted | omitted | 2.248 |
| reconciliation / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| repair / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| repair / completed ops/s | 3 | 199.943 | 199.958 | 200.006 | 0.096 | omitted | omitted | 200.117 |
| repair / successful request p50 ms | 3 | 0.602 | 0.664 | 0.666 | 0.065 | omitted | omitted | 0.732 |
| repair / successful request p95 ms | 3 | 1.314 | 1.314 | 1.409 | 0.164 | omitted | omitted | 1.598 |
| repair / successful request p99 ms | 3 | 1.849 | 2.602 | 2.646 | 0.820 | omitted | omitted | 3.487 |
| repair / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| repair / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| repair / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 3 trials, unmeasured in 0. Every successful SET remains in the main validation ledger when the later acknowledgement command is short, errors, or is not attempted. Diagnostic SET/WAIT/WAITAOF stage latency remains in individual reports and does not increase logical request or pacing counts. Errors and unattempted commands supply unknown acknowledgement counts, never fabricated zero replies.

| Phase | SET acknowledged | SET errors | Policy satisfied |
|---|---:|---:|---:|
| baseline | 1249 | 0 | 1249 |
| fault | 658 | 124 | 658 |
| post_recovery | 596 | 0 | 596 |
| reconciliation | 1685 | 0 | 1685 |
| repair | 458 | 0 | 458 |

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 32 |
| fault | network | 92 |
| fault | timeout | 114 |

## Acknowledged-write accounting

Validation attempted in 3 trials and complete in 3. SET attempts: 4770; acknowledged writes: 4646; ambiguous SET outcomes: 124.

Checked: 4646; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 3 | 1494.000 | 1545.000 | 1548.667 | 56.589 | omitted | omitted | 1607.000 |
| ambiguous writes | 3 | 28.000 | 44.000 | 41.333 | 12.220 | omitted | omitted | 52.000 |
| missing writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | dab1232736a2b59ad498e657 | recovered | true | true | [JSON](../../trials/position-0002/batch-bb39bfe07e9868b38d98ea5d/runs/dab1232736a2b59ad498e657/report.json), [Markdown](../../trials/position-0002/batch-bb39bfe07e9868b38d98ea5d/runs/dab1232736a2b59ad498e657/report.md) |
| 2 | 60141689fee92e0e4cd9bc77 | recovered | true | true | [JSON](../../trials/position-0004/batch-c7b388dcb147ac10049fac7b/runs/60141689fee92e0e4cd9bc77/report.json), [Markdown](../../trials/position-0004/batch-c7b388dcb147ac10049fac7b/runs/60141689fee92e0e4cd9bc77/report.md) |
| 3 | 8d7edf6681a3f591879efeda | recovered | true | true | [JSON](../../trials/position-0009/batch-d0662642c7cf1047aa35a9ed/runs/8d7edf6681a3f591879efeda/report.json), [Markdown](../../trials/position-0009/batch-d0662642c7cf1047aa35a9ed/runs/8d7edf6681a3f591879efeda/report.md) |

## Measurement methodology

- A seeded permutation of condition indices is rotated by one position each round. Within each complete block of N rounds, every condition occupies every round position once. An incomplete block is only partially balanced; this rotating Latin order does not balance all predecessor/successor pairs.
- The order seed controls only the condition schedule. Effective workload seeds and all other configuration values remain unchanged. Each entry executes one fresh sequential batch trial with persisted-report integrity and cleanup checks.
- Further trials require verified individual evidence and successful cleanup, or verified evidence that no environment was provisioned. Unsafe cleanup, missing evidence, cancellation and summary-write failures stop execution; all remaining planned entries stay explicitly unstarted.
- Failed and partial trials with valid evidence and safe cleanup remain in their condition statistics. Timing statistics are conditional on observed milestones; missing and censored observations are not zero-duration events. Request latency quantiles are not pooled across trials or conditions.
- Condition summaries reuse batch aggregation. Condition-summary report paths are relative to their own summary directory; schedule report paths and batch references are relative to the comparison directory. Host and environment references point to individual JSON report sections.
- Ordering reduces systematic position confounding but does not prove causality, independent samples, equal host conditions or population reliability. Compare individual environmental metadata and actual order before interpreting differences.

Effective experiment configuration SHA-256: `de05e21df2ed5b0a82dc0af8fa135f9173da7c1fdf43019ed7cd93af6d12f3cf`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
