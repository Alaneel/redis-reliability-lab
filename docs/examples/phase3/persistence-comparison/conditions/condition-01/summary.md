# Redis Reliability Lab — repeated trials

Batch: `comparison-dde55f9ada692afef5838308-condition-01`  
Outcome: **partial**  
Started: 2026-10-11T08:50:18.083444Z  
Finished: 2026-10-11T08:58:02.56308Z

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
| persistence_recovered | 0 | 0 | 0 | 3 | 0 | 0 |
| all_processes_stopped | 3 | 0 | 0 | 0 | 0 | 0 |

Descriptive timing statistics use milliseconds. **P90/P95 are omitted for N<20**; small samples do not support reliable tail estimates. Sample standard deviation uses N-1 and is undefined for N<2. These are descriptive statistics, not confidence intervals.

| Timing / anchor | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| fault action duration | 3 | 2197.910 | 2219.856 | 2271.373 | 108.791 | omitted | omitted | 2396.354 |
| failure_detected / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| failure_detected / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_promoted / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_promoted / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| full_slot_coverage / orchestration start | 3 | 16375.412 | 16394.367 | 16467.893 | 144.078 | omitted | omitted | 16633.900 |
| full_slot_coverage / action completion | 3 | 14155.557 | 14196.457 | 14196.520 | 40.995 | omitted | omitted | 14237.547 |
| topology_converged / orchestration start | 3 | 17602.125 | 17654.187 | 17691.128 | 112.134 | omitted | omitted | 17817.071 |
| topology_converged / action completion | 3 | 15382.270 | 15420.718 | 15419.755 | 37.013 | omitted | omitted | 15456.277 |
| client_first_all_shard_success / orchestration start | 3 | 23684.565 | 23803.056 | 23816.565 | 139.247 | omitted | omitted | 23962.074 |
| client_first_all_shard_success / action completion | 3 | 21486.655 | 21565.721 | 21545.192 | 51.443 | omitted | omitted | 21583.201 |
| client_stable / orchestration start | 3 | 24887.334 | 25047.236 | 25046.449 | 158.722 | omitted | omitted | 25204.776 |
| client_stable / action completion | 3 | 22689.425 | 22808.422 | 22775.076 | 74.779 | omitted | omitted | 22827.380 |
| combined_stable / orchestration start | 3 | 24887.334 | 25047.236 | 25046.449 | 158.722 | omitted | omitted | 25204.776 |
| combined_stable / action completion | 3 | 22689.425 | 22808.422 | 22775.076 | 74.779 | omitted | omitted | 22827.380 |
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
| stable_redundancy / orchestration start | 3 | 17602.125 | 17654.187 | 17691.128 | 112.134 | omitted | omitted | 17817.071 |
| stable_redundancy / action completion | 3 | 15382.270 | 15420.718 | 15419.755 | 37.013 | omitted | omitted | 15456.277 |
| original_topology_stable / orchestration start | 3 | 17602.125 | 17654.187 | 17691.128 | 112.134 | omitted | omitted | 17817.071 |
| original_topology_stable / action completion | 3 | 15382.270 | 15420.718 | 15419.755 | 37.013 | omitted | omitted | 15456.277 |
| persistence_recovered / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| persistence_recovered / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| all_processes_stopped / orchestration start | 3 | 2197.910 | 2219.856 | 2271.373 | 108.791 | omitted | omitted | 2396.354 |
| all_processes_stopped / action completion | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |

## Client operations

Completed request fractions weight actual request counts. Throughput divides summed completed/successful counts by summed trial phase wall-clock seconds; it does not imply workers achieved the configured offered rate. Zero-attempt phases have an unobserved success fraction.

| Phase | Trials with phase | Completed | Successful | Failed | Request success fraction | Completed ops/s | Successful ops/s |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 3 | 2414 | 2414 | 0 | 100.00% | 199.796 | 199.796 |

Whole-workload schedule counters measured in 3 trials; legacy or incomplete counters unmeasured in 0. Offered: 2415; started: 2414; completed: 2414; dropped: 1. Started fraction of offered: 99.96%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| baseline / completed ops/s | 3 | 199.575 | 199.851 | 199.796 | 0.199 | omitted | omitted | 199.962 |
| baseline / successful request p50 ms | 3 | 0.428 | 0.472 | 0.482 | 0.060 | omitted | omitted | 0.546 |
| baseline / successful request p95 ms | 3 | 0.664 | 0.697 | 0.698 | 0.034 | omitted | omitted | 0.732 |
| baseline / successful request p99 ms | 3 | 0.934 | 1.449 | 1.327 | 0.348 | omitted | omitted | 1.598 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 3 trials, unmeasured in 0. Every successful SET remains in the main validation ledger even when WAIT is short, errors, or is not attempted. Stage command latency remains in individual reports; it does not increase logical request counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 1251 | 0 | 0 | 0 | 0 | 0 | 0 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. Attempted WAIT client-call timeouts are counted separately above.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|

## Acknowledged-write accounting

Validation attempted in 3 trials and complete in 3. SET attempts: 1251; acknowledged writes: 1251; ambiguous SET outcomes: 0.

Checked: 1251; missing: 1251; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 3 | 416.000 | 417.000 | 417.000 | 1.000 | omitted | omitted | 418.000 |
| ambiguous writes | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| missing writes / validation attempted | 3 | 416.000 | 417.000 | 417.000 | 1.000 | omitted | omitted | 418.000 |
| corrupt writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | aa2a96b0899c2885b2aa3a32 | partial | true | true | [JSON](../../trials/position-0002/batch-aedea818cdd8ecb1335a5885/runs/aa2a96b0899c2885b2aa3a32/report.json), [Markdown](../../trials/position-0002/batch-aedea818cdd8ecb1335a5885/runs/aa2a96b0899c2885b2aa3a32/report.md) |
| 2 | fa1b57f987b8c265178b1829 | partial | true | true | [JSON](../../trials/position-0004/batch-9b1b272f35ff29ce65565565/runs/fa1b57f987b8c265178b1829/report.json), [Markdown](../../trials/position-0004/batch-9b1b272f35ff29ce65565565/runs/fa1b57f987b8c265178b1829/report.md) |
| 3 | 4a3cd1c67b9ff55137ab5115 | partial | true | true | [JSON](../../trials/position-0009/batch-61bb928a8f8d405beb4b630e/runs/4a3cd1c67b9ff55137ab5115/report.json), [Markdown](../../trials/position-0009/batch-61bb928a8f8d405beb4b630e/runs/4a3cd1c67b9ff55137ab5115/report.md) |

## Execution errors

- trial 1: persistence experiment ended partial: frozen data verified=false, original redundancy stable=true, restoration deadline=false
- trial 1: persistence experiment ended partial: frozen data verified=false, original redundancy stable=true, restoration deadline=false
- trial 1: persistence experiment ended partial: frozen data verified=false, original redundancy stable=true, restoration deadline=false

## Measurement methodology

- A seeded permutation of condition indices is rotated by one position each round. Within each complete block of N rounds, every condition occupies every round position once. An incomplete block is only partially balanced; this rotating Latin order does not balance all predecessor/successor pairs.
- The order seed controls only the condition schedule. Effective workload seeds and all other configuration values remain unchanged. Each entry executes one fresh sequential batch trial with persisted-report integrity and cleanup checks.
- Further trials require verified individual evidence and successful cleanup, or verified evidence that no environment was provisioned. Unsafe cleanup, missing evidence, cancellation and summary-write failures stop execution; all remaining planned entries stay explicitly unstarted.
- Failed and partial trials with valid evidence and safe cleanup remain in their condition statistics. Timing statistics are conditional on observed milestones; missing and censored observations are not zero-duration events. Request latency quantiles are not pooled across trials or conditions.
- Condition summaries reuse batch aggregation. Condition-summary report paths are relative to their own summary directory; schedule report paths and batch references are relative to the comparison directory. Host and environment references point to individual JSON report sections.
- Ordering reduces systematic position confounding but does not prove causality, independent samples, equal host conditions or population reliability. Compare individual environmental metadata and actual order before interpreting differences.

Effective experiment configuration SHA-256: `359df9c4ebd051da3e82059877c38c43334b012d0ec20f156604a3c2447ffe2a`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
