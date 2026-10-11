# Redis Reliability Lab — repeated trials

Batch: `comparison-dde55f9ada692afef5838308-condition-03`  
Outcome: **recovered**  
Started: 2026-10-11T08:50:18.083444Z  
Finished: 2026-10-11T08:58:02.56308Z

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
| failure_detected | 0 | 0 | 0 | 3 | 0 | 0 |
| replica_promoted | 0 | 0 | 0 | 3 | 0 | 0 |
| full_slot_coverage | 3 | 0 | 0 | 0 | 0 | 0 |
| topology_converged | 3 | 0 | 0 | 0 | 0 | 0 |
| client_first_all_shard_success | 3 | 0 | 0 | 0 | 0 | 0 |
| client_stable | 3 | 0 | 0 | 0 | 0 | 0 |
| combined_stable | 3 | 0 | 0 | 0 | 0 | 0 |
| write_policy_stable | 0 | 0 | 0 | 3 | 0 | 0 |
| replication_interrupted | 0 | 0 | 0 | 3 | 0 | 0 |
| target_reachable_after_repair | 0 | 0 | 0 | 3 | 0 | 0 |
| membership_rejoined | 0 | 0 | 0 | 3 | 0 | 0 |
| replica_link_up | 0 | 0 | 0 | 3 | 0 | 0 |
| replica_sync_complete | 0 | 0 | 0 | 3 | 0 | 0 |
| stable_redundancy | 3 | 0 | 0 | 0 | 0 | 0 |
| original_topology_stable | 3 | 0 | 0 | 0 | 0 | 0 |
| persistence_recovered | 3 | 0 | 0 | 0 | 0 | 0 |
| all_processes_stopped | 3 | 0 | 0 | 0 | 0 | 0 |

Descriptive timing statistics use milliseconds. **P90/P95 are omitted for N<20**; small samples do not support reliable tail estimates. Sample standard deviation uses N-1 and is undefined for N<2. These are descriptive statistics, not confidence intervals.

| Timing / anchor | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| fault action duration | 3 | 2158.851 | 2229.842 | 2217.609 | 53.698 | omitted | omitted | 2264.135 |
| failure_detected / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| failure_detected / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_promoted / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_promoted / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| full_slot_coverage / orchestration start | 3 | 16323.188 | 16454.906 | 16418.033 | 82.813 | omitted | omitted | 16476.006 |
| full_slot_coverage / action completion | 3 | 14164.337 | 14190.771 | 14200.424 | 41.759 | omitted | omitted | 14246.165 |
| topology_converged / orchestration start | 3 | 17565.176 | 17633.924 | 17647.922 | 90.560 | omitted | omitted | 17744.665 |
| topology_converged / action completion | 3 | 15369.789 | 15406.325 | 15430.313 | 75.434 | omitted | omitted | 15514.824 |
| client_first_all_shard_success / orchestration start | 3 | 23697.381 | 23787.545 | 23826.605 | 152.551 | omitted | omitted | 23994.888 |
| client_first_all_shard_success / action completion | 3 | 21523.410 | 21538.530 | 21608.996 | 135.355 | omitted | omitted | 21765.046 |
| client_stable / orchestration start | 3 | 24954.929 | 25028.941 | 25061.706 | 126.386 | omitted | omitted | 25201.248 |
| client_stable / action completion | 3 | 22764.806 | 22796.079 | 22844.097 | 111.356 | omitted | omitted | 22971.406 |
| combined_stable / orchestration start | 3 | 24954.929 | 25028.941 | 25061.706 | 126.386 | omitted | omitted | 25201.248 |
| combined_stable / action completion | 3 | 22764.806 | 22796.079 | 22844.097 | 111.356 | omitted | omitted | 22971.406 |
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
| stable_redundancy / orchestration start | 3 | 17565.176 | 17633.924 | 17647.922 | 90.560 | omitted | omitted | 17744.665 |
| stable_redundancy / action completion | 3 | 15369.789 | 15406.325 | 15430.313 | 75.434 | omitted | omitted | 15514.824 |
| original_topology_stable / orchestration start | 3 | 17565.176 | 17633.924 | 17647.922 | 90.560 | omitted | omitted | 17744.665 |
| original_topology_stable / action completion | 3 | 15369.789 | 15406.325 | 15430.313 | 75.434 | omitted | omitted | 15514.824 |
| persistence_recovered / orchestration start | 3 | 23677.718 | 23781.657 | 23815.827 | 157.990 | omitted | omitted | 23988.105 |
| persistence_recovered / action completion | 3 | 21517.522 | 21518.867 | 21598.217 | 138.605 | omitted | omitted | 21758.263 |
| all_processes_stopped / orchestration start | 3 | 2158.851 | 2229.842 | 2217.609 | 53.698 | omitted | omitted | 2264.135 |
| all_processes_stopped / action completion | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |

## Client operations

Completed request fractions weight actual request counts. Throughput divides summed completed/successful counts by summed trial phase wall-clock seconds; it does not imply workers achieved the configured offered rate. Zero-attempt phases have an unobserved success fraction.

| Phase | Trials with phase | Completed | Successful | Failed | Request success fraction | Completed ops/s | Successful ops/s |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 3 | 2415 | 2415 | 0 | 100.00% | 199.901 | 199.901 |

Whole-workload schedule counters measured in 3 trials; legacy or incomplete counters unmeasured in 0. Offered: 2415; started: 2415; completed: 2415; dropped: 0. Started fraction of offered: 100.00%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| baseline / completed ops/s | 3 | 199.864 | 199.917 | 199.901 | 0.033 | omitted | omitted | 199.923 |
| baseline / successful request p50 ms | 3 | 0.664 | 0.732 | 0.709 | 0.039 | omitted | omitted | 0.732 |
| baseline / successful request p95 ms | 3 | 1.380 | 1.380 | 1.403 | 0.040 | omitted | omitted | 1.449 |
| baseline / successful request p99 ms | 3 | 2.248 | 2.360 | 2.362 | 0.115 | omitted | omitted | 2.478 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 3 trials, unmeasured in 0. Every successful SET remains in the main validation ledger even when WAIT is short, errors, or is not attempted. Stage command latency remains in individual reports; it does not increase logical request counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 1252 | 0 | 0 | 0 | 0 | 0 | 0 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. Attempted WAIT client-call timeouts are counted separately above.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|

## Acknowledged-write accounting

Validation attempted in 3 trials and complete in 3. SET attempts: 1252; acknowledged writes: 1252; ambiguous SET outcomes: 0.

Checked: 1252; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 3 | 417.000 | 417.000 | 417.333 | 0.577 | omitted | omitted | 418.000 |
| ambiguous writes | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| missing writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | bc5ed8626797ee3fa6013069 | recovered | true | true | [JSON](../../trials/position-0003/batch-b36014fbb9afa7b331dd7462/runs/bc5ed8626797ee3fa6013069/report.json), [Markdown](../../trials/position-0003/batch-b36014fbb9afa7b331dd7462/runs/bc5ed8626797ee3fa6013069/report.md) |
| 2 | a50433e55ee7e3e72db5414c | recovered | true | true | [JSON](../../trials/position-0005/batch-ea3add7d0827fb08c748b425/runs/a50433e55ee7e3e72db5414c/report.json), [Markdown](../../trials/position-0005/batch-ea3add7d0827fb08c748b425/runs/a50433e55ee7e3e72db5414c/report.md) |
| 3 | 516b2a2428d8d408b3fc1b4f | recovered | true | true | [JSON](../../trials/position-0007/batch-b1104f16b924e36fcac1a492/runs/516b2a2428d8d408b3fc1b4f/report.json), [Markdown](../../trials/position-0007/batch-b1104f16b924e36fcac1a492/runs/516b2a2428d8d408b3fc1b4f/report.md) |

## Measurement methodology

- A seeded permutation of condition indices is rotated by one position each round. Within each complete block of N rounds, every condition occupies every round position once. An incomplete block is only partially balanced; this rotating Latin order does not balance all predecessor/successor pairs.
- The order seed controls only the condition schedule. Effective workload seeds and all other configuration values remain unchanged. Each entry executes one fresh sequential batch trial with persisted-report integrity and cleanup checks.
- Further trials require verified individual evidence and successful cleanup, or verified evidence that no environment was provisioned. Unsafe cleanup, missing evidence, cancellation and summary-write failures stop execution; all remaining planned entries stay explicitly unstarted.
- Failed and partial trials with valid evidence and safe cleanup remain in their condition statistics. Timing statistics are conditional on observed milestones; missing and censored observations are not zero-duration events. Request latency quantiles are not pooled across trials or conditions.
- Condition summaries reuse batch aggregation. Condition-summary report paths are relative to their own summary directory; schedule report paths and batch references are relative to the comparison directory. Host and environment references point to individual JSON report sections.
- Ordering reduces systematic position confounding but does not prove causality, independent samples, equal host conditions or population reliability. Compare individual environmental metadata and actual order before interpreting differences.

Effective experiment configuration SHA-256: `4d4fd037cd04b6f208194358a86091403dc4075d14d2149be6a06dc3067b954e`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
