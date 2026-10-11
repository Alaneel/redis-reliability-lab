# Redis Reliability Lab — repeated trials

Batch: `comparison-321638985ff49f9040049bf4-condition-03`  
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
| fault action duration | 3 | 236.692 | 241.023 | 242.533 | 6.723 | omitted | omitted | 249.882 |
| failure_detected / orchestration start | 3 | 1801.791 | 1813.686 | 1813.059 | 10.969 | omitted | omitted | 1823.702 |
| failure_detected / action completion | 3 | 1565.098 | 1572.663 | 1570.527 | 4.737 | omitted | omitted | 1573.819 |
| replica_promoted / orchestration start | 3 | 2586.054 | 2611.642 | 2882.591 | 491.624 | omitted | omitted | 3450.078 |
| replica_promoted / action completion | 3 | 2349.361 | 2361.760 | 2640.059 | 492.804 | omitted | omitted | 3209.055 |
| full_slot_coverage / orchestration start | 3 | 2586.054 | 2611.642 | 2882.591 | 491.624 | omitted | omitted | 3450.078 |
| full_slot_coverage / action completion | 3 | 2349.361 | 2361.760 | 2640.059 | 492.804 | omitted | omitted | 3209.055 |
| topology_converged / orchestration start | 3 | 2586.054 | 2611.642 | 2882.591 | 491.624 | omitted | omitted | 3450.078 |
| topology_converged / action completion | 3 | 2349.361 | 2361.760 | 2640.059 | 492.804 | omitted | omitted | 3209.055 |
| client_first_all_shard_success / orchestration start | 3 | 3077.328 | 3497.771 | 3368.783 | 252.962 | omitted | omitted | 3531.250 |
| client_first_all_shard_success / action completion | 3 | 2836.305 | 3261.078 | 3126.251 | 251.305 | omitted | omitted | 3281.368 |
| client_stable / orchestration start | 3 | 4254.828 | 4656.215 | 4541.256 | 249.659 | omitted | omitted | 4712.726 |
| client_stable / action completion | 3 | 4013.804 | 4419.523 | 4298.724 | 247.696 | omitted | omitted | 4462.844 |
| combined_stable / orchestration start | 3 | 4638.861 | 4656.215 | 4669.268 | 38.624 | omitted | omitted | 4712.726 |
| combined_stable / action completion | 3 | 4397.838 | 4419.523 | 4426.735 | 33.098 | omitted | omitted | 4462.844 |
| write_policy_stable / orchestration start | 3 | 11426.722 | 11530.280 | 11615.215 | 242.391 | omitted | omitted | 11888.643 |
| write_policy_stable / action completion | 3 | 11190.030 | 11280.398 | 11372.683 | 242.352 | omitted | omitted | 11647.620 |
| replication_interrupted / orchestration start | 3 | 2586.054 | 2611.642 | 2882.591 | 491.624 | omitted | omitted | 3450.078 |
| replication_interrupted / action completion | 3 | 2349.361 | 2361.760 | 2640.059 | 492.804 | omitted | omitted | 3209.055 |
| target_reachable_after_repair / orchestration start | 3 | 6253.872 | 6329.500 | 6321.547 | 64.069 | omitted | omitted | 6381.267 |
| target_reachable_after_repair / action completion | 3 | 6012.849 | 6092.808 | 6079.014 | 60.460 | omitted | omitted | 6131.385 |
| membership_rejoined / orchestration start | 3 | 6253.872 | 6329.500 | 6321.547 | 64.069 | omitted | omitted | 6381.267 |
| membership_rejoined / action completion | 3 | 6012.849 | 6092.808 | 6079.014 | 60.460 | omitted | omitted | 6131.385 |
| replica_link_up / orchestration start | 3 | 10142.302 | 10572.642 | 10542.963 | 386.677 | omitted | omitted | 10913.946 |
| replica_link_up / action completion | 3 | 9905.610 | 10322.760 | 10300.431 | 384.144 | omitted | omitted | 10672.922 |
| replica_sync_complete / orchestration start | 3 | 10142.302 | 10572.642 | 10542.963 | 386.677 | omitted | omitted | 10913.946 |
| replica_sync_complete / action completion | 3 | 9905.610 | 10322.760 | 10300.431 | 384.144 | omitted | omitted | 10672.922 |
| stable_redundancy / orchestration start | 3 | 11424.893 | 11837.719 | 11821.038 | 388.073 | omitted | omitted | 12200.501 |
| stable_redundancy / action completion | 3 | 11188.201 | 11587.837 | 11578.505 | 385.723 | omitted | omitted | 11959.478 |
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
| baseline | 3 | 2409 | 2409 | 0 | 100.00% | 199.857 | 199.857 |
| fault | 3 | 1745 | 1352 | 393 | 77.48% | 121.627 | 94.235 |
| post_recovery | 3 | 1201 | 1201 | 0 | 100.00% | 200.123 | 200.123 |
| reconciliation | 3 | 3324 | 2921 | 403 | 87.88% | 200.061 | 175.806 |
| repair | 3 | 921 | 759 | 162 | 82.41% | 199.819 | 164.672 |

Whole-workload schedule counters measured in 3 trials; legacy or incomplete counters unmeasured in 0. Offered: 10724; started: 9600; completed: 9600; dropped: 1124. Started fraction of offered: 89.52%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| baseline / completed ops/s | 3 | 199.810 | 199.829 | 199.857 | 0.065 | omitted | omitted | 199.931 |
| baseline / successful request p50 ms | 3 | 1.135 | 1.252 | 1.234 | 0.091 | omitted | omitted | 1.314 |
| baseline / successful request p95 ms | 3 | 2.360 | 2.478 | 2.480 | 0.121 | omitted | omitted | 2.602 |
| baseline / successful request p99 ms | 3 | 2.869 | 3.163 | 3.065 | 0.170 | omitted | omitted | 3.163 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 3 | 0.769 | 0.776 | 0.775 | 0.005 | omitted | omitted | 0.779 |
| fault / completed ops/s | 3 | 121.016 | 121.361 | 121.627 | 0.779 | omitted | omitted | 122.505 |
| fault / successful request p50 ms | 3 | 0.495 | 0.546 | 0.579 | 0.105 | omitted | omitted | 0.697 |
| fault / successful request p95 ms | 3 | 2.869 | 3.012 | 3.015 | 0.147 | omitted | omitted | 3.163 |
| fault / successful request p99 ms | 3 | 41.990 | 231.616 | 176.320 | 116.938 | omitted | omitted | 255.356 |
| fault / failed request p50 ms | 3 | 111.411 | 111.411 | 111.411 | 0.000 | omitted | omitted | 111.411 |
| fault / failed request p95 ms | 3 | 500.586 | 500.606 | 500.898 | 0.523 | omitted | omitted | 501.501 |
| fault / failed request p99 ms | 3 | 500.586 | 500.606 | 500.898 | 0.523 | omitted | omitted | 501.501 |
| post_recovery / trial success fraction | 3 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| post_recovery / completed ops/s | 3 | 199.965 | 199.985 | 200.123 | 0.257 | omitted | omitted | 200.420 |
| post_recovery / successful request p50 ms | 3 | 0.573 | 0.934 | 0.829 | 0.223 | omitted | omitted | 0.981 |
| post_recovery / successful request p95 ms | 3 | 1.942 | 2.602 | 2.471 | 0.477 | omitted | omitted | 2.869 |
| post_recovery / successful request p99 ms | 3 | 2.869 | 3.163 | 3.292 | 0.501 | omitted | omitted | 3.845 |
| post_recovery / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / trial success fraction | 3 | 0.875 | 0.877 | 0.879 | 0.004 | omitted | omitted | 0.884 |
| reconciliation / completed ops/s | 3 | 200.001 | 200.042 | 200.059 | 0.067 | omitted | omitted | 200.133 |
| reconciliation / successful request p50 ms | 3 | 0.472 | 0.472 | 0.480 | 0.014 | omitted | omitted | 0.495 |
| reconciliation / successful request p95 ms | 3 | 2.478 | 2.478 | 2.563 | 0.147 | omitted | omitted | 2.732 |
| reconciliation / successful request p99 ms | 3 | 3.163 | 3.487 | 3.630 | 0.552 | omitted | omitted | 4.239 |
| reconciliation / failed request p50 ms | 3 | 111.411 | 111.411 | 111.411 | 0.000 | omitted | omitted | 111.411 |
| reconciliation / failed request p95 ms | 3 | 122.831 | 122.831 | 124.878 | 3.546 | omitted | omitted | 128.972 |
| reconciliation / failed request p99 ms | 3 | 135.421 | 140.727 | 139.446 | 3.562 | omitted | omitted | 142.192 |
| repair / trial success fraction | 3 | 0.819 | 0.823 | 0.824 | 0.005 | omitted | omitted | 0.830 |
| repair / completed ops/s | 3 | 199.441 | 199.989 | 199.822 | 0.331 | omitted | omitted | 200.037 |
| repair / successful request p50 ms | 3 | 0.495 | 0.495 | 0.512 | 0.029 | omitted | omitted | 0.546 |
| repair / successful request p95 ms | 3 | 2.248 | 2.602 | 2.573 | 0.312 | omitted | omitted | 2.869 |
| repair / successful request p99 ms | 3 | 3.163 | 3.662 | 4.467 | 1.843 | omitted | omitted | 6.576 |
| repair / failed request p50 ms | 3 | 111.411 | 111.411 | 111.411 | 0.000 | omitted | omitted | 111.411 |
| repair / failed request p95 ms | 3 | 128.972 | 128.972 | 128.972 | 0.000 | omitted | omitted | 128.972 |
| repair / failed request p99 ms | 3 | 130.613 | 130.697 | 132.178 | 2.637 | omitted | omitted | 135.223 |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 3 trials, unmeasured in 0. Every successful SET remains in the main validation ledger when the later acknowledgement command is short, errors, or is not attempted. Diagnostic SET/WAIT/WAITAOF stage latency remains in individual reports and does not increase logical request or pacing counts. Errors and unattempted commands supply unknown acknowledgement counts, never fabricated zero replies.

WAITAOF requires local AOF acknowledgement 1 and replica AOF acknowledgement 1. Local and replica satisfied counts overlap; they are independently reconciled against paired valid replies.

| Phase | SET acknowledged | SET errors | WAITAOF attempted | Both satisfied | Short | Errors | Client timeouts | Local satisfied | Replica satisfied |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline | 1248 | 0 | 1248 | 1248 | 0 | 0 | 0 | 1248 | 1248 |
| fault | 710 | 111 | 710 | 545 | 165 | 0 | 0 | 710 | 545 |
| post_recovery | 586 | 0 | 586 | 586 | 0 | 0 | 0 | 586 | 586 |
| reconciliation | 1729 | 0 | 1729 | 1326 | 403 | 0 | 0 | 1729 | 1326 |
| repair | 486 | 0 | 486 | 324 | 162 | 0 | 0 | 486 | 324 |

SET acknowledgements without a subsequent WAITAOF attempt: 0; client timeouts among those: 0. Attempted WAITAOF client-call timeouts are counted separately above. Total local satisfied: 4759; replica satisfied: 4029; both satisfied: 4029.

| Observed WAITAOF local/replicas pair | Replies |
|---|---:|
| 1/0 | 730 |
| 1/1 | 4029 |

WAITAOF-satisfied subset validation: expected 4029; checked 4029; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Validation attempted in 3 trials. Complete means every subset record was checked; preservation also requires zero missing, corrupt and unreadable records.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 21 |
| fault | network | 90 |
| fault | timeout | 117 |
| fault | waitaof_short | 165 |
| reconciliation | waitaof_short | 403 |
| repair | waitaof_short | 162 |

## Acknowledged-write accounting

Validation attempted in 3 trials and complete in 3. SET attempts: 4870; acknowledged writes: 4759; ambiguous SET outcomes: 111.

Checked: 4759; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 3 | 1544.000 | 1586.000 | 1586.333 | 42.501 | omitted | omitted | 1629.000 |
| ambiguous writes | 3 | 33.000 | 36.000 | 37.000 | 4.583 | omitted | omitted | 42.000 |
| missing writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 3 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | e519c821277d8791d7e6796e | recovered | true | true | [JSON](../../trials/position-0003/batch-bb6f098573c5605a2fe65248/runs/e519c821277d8791d7e6796e/report.json), [Markdown](../../trials/position-0003/batch-bb6f098573c5605a2fe65248/runs/e519c821277d8791d7e6796e/report.md) |
| 2 | 688a7dc24b138ea24d1b7e0d | recovered | true | true | [JSON](../../trials/position-0005/batch-05a52c53dbeff938cd618ed7/runs/688a7dc24b138ea24d1b7e0d/report.json), [Markdown](../../trials/position-0005/batch-05a52c53dbeff938cd618ed7/runs/688a7dc24b138ea24d1b7e0d/report.md) |
| 3 | 37475e6b0cf7b098669c0663 | recovered | true | true | [JSON](../../trials/position-0007/batch-82ffccc231329f6e4d7dae7b/runs/37475e6b0cf7b098669c0663/report.json), [Markdown](../../trials/position-0007/batch-82ffccc231329f6e4d7dae7b/runs/37475e6b0cf7b098669c0663/report.md) |

## Measurement methodology

- A seeded permutation of condition indices is rotated by one position each round. Within each complete block of N rounds, every condition occupies every round position once. An incomplete block is only partially balanced; this rotating Latin order does not balance all predecessor/successor pairs.
- The order seed controls only the condition schedule. Effective workload seeds and all other configuration values remain unchanged. Each entry executes one fresh sequential batch trial with persisted-report integrity and cleanup checks.
- Further trials require verified individual evidence and successful cleanup, or verified evidence that no environment was provisioned. Unsafe cleanup, missing evidence, cancellation and summary-write failures stop execution; all remaining planned entries stay explicitly unstarted.
- Failed and partial trials with valid evidence and safe cleanup remain in their condition statistics. Timing statistics are conditional on observed milestones; missing and censored observations are not zero-duration events. Request latency quantiles are not pooled across trials or conditions.
- Condition summaries reuse batch aggregation. Condition-summary report paths are relative to their own summary directory; schedule report paths and batch references are relative to the comparison directory. Host and environment references point to individual JSON report sections.
- Ordering reduces systematic position confounding but does not prove causality, independent samples, equal host conditions or population reliability. Compare individual environmental metadata and actual order before interpreting differences.

Effective experiment configuration SHA-256: `975080021092fa9a75ee5da95f13abfb2ac4db2ce23d3e320a4867ae224e9637`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
