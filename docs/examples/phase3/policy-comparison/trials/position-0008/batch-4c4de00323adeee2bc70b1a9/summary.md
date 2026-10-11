# Redis Reliability Lab — repeated trials

Batch: `batch-4c4de00323adeee2bc70b1a9`  
Outcome: **recovered**  
Started: 2026-10-11T09:16:05.416536Z  
Finished: 2026-10-11T09:16:42.753194Z

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
| failure_detected | 1 | 0 | 0 | 0 | 0 | 0 |
| replica_promoted | 1 | 0 | 0 | 0 | 0 | 0 |
| full_slot_coverage | 1 | 0 | 0 | 0 | 0 | 0 |
| topology_converged | 1 | 0 | 0 | 0 | 0 | 0 |
| client_first_all_shard_success | 1 | 0 | 0 | 0 | 0 | 0 |
| client_stable | 1 | 0 | 0 | 0 | 0 | 0 |
| combined_stable | 1 | 0 | 0 | 0 | 0 | 0 |
| write_policy_stable | 1 | 0 | 0 | 0 | 0 | 0 |
| replication_interrupted | 1 | 0 | 0 | 0 | 0 | 0 |
| target_reachable_after_repair | 1 | 0 | 0 | 0 | 0 | 0 |
| membership_rejoined | 1 | 0 | 0 | 0 | 0 | 0 |
| replica_link_up | 1 | 0 | 0 | 0 | 0 | 0 |
| replica_sync_complete | 1 | 0 | 0 | 0 | 0 | 0 |
| stable_redundancy | 1 | 0 | 0 | 0 | 0 | 0 |
| original_topology_stable | 0 | 0 | 0 | 1 | 0 | 0 |
| persistence_recovered | 0 | 0 | 0 | 1 | 0 | 0 |
| all_processes_stopped | 0 | 0 | 0 | 1 | 0 | 0 |

Descriptive timing statistics use milliseconds. **P90/P95 are omitted for N<20**; small samples do not support reliable tail estimates. Sample standard deviation uses N-1 and is undefined for N<2. These are descriptive statistics, not confidence intervals.

| Timing / anchor | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| fault action duration | 1 | 235.555 | 235.555 | 235.555 | undefined | omitted | omitted | 235.555 |
| failure_detected / orchestration start | 1 | 1797.670 | 1797.670 | 1797.670 | undefined | omitted | omitted | 1797.670 |
| failure_detected / action completion | 1 | 1562.115 | 1562.115 | 1562.115 | undefined | omitted | omitted | 1562.115 |
| replica_promoted / orchestration start | 1 | 2572.051 | 2572.051 | 2572.051 | undefined | omitted | omitted | 2572.051 |
| replica_promoted / action completion | 1 | 2336.497 | 2336.497 | 2336.497 | undefined | omitted | omitted | 2336.497 |
| full_slot_coverage / orchestration start | 1 | 2572.051 | 2572.051 | 2572.051 | undefined | omitted | omitted | 2572.051 |
| full_slot_coverage / action completion | 1 | 2336.497 | 2336.497 | 2336.497 | undefined | omitted | omitted | 2336.497 |
| topology_converged / orchestration start | 1 | 2572.051 | 2572.051 | 2572.051 | undefined | omitted | omitted | 2572.051 |
| topology_converged / action completion | 1 | 2336.497 | 2336.497 | 2336.497 | undefined | omitted | omitted | 2336.497 |
| client_first_all_shard_success / orchestration start | 1 | 3035.827 | 3035.827 | 3035.827 | undefined | omitted | omitted | 3035.827 |
| client_first_all_shard_success / action completion | 1 | 2800.273 | 2800.273 | 2800.273 | undefined | omitted | omitted | 2800.273 |
| client_stable / orchestration start | 1 | 4230.164 | 4230.164 | 4230.164 | undefined | omitted | omitted | 4230.164 |
| client_stable / action completion | 1 | 3994.609 | 3994.609 | 3994.609 | undefined | omitted | omitted | 3994.609 |
| combined_stable / orchestration start | 1 | 4230.164 | 4230.164 | 4230.164 | undefined | omitted | omitted | 4230.164 |
| combined_stable / action completion | 1 | 3994.609 | 3994.609 | 3994.609 | undefined | omitted | omitted | 3994.609 |
| write_policy_stable / orchestration start | 1 | 10212.858 | 10212.858 | 10212.858 | undefined | omitted | omitted | 10212.858 |
| write_policy_stable / action completion | 1 | 9977.303 | 9977.303 | 9977.303 | undefined | omitted | omitted | 9977.303 |
| replication_interrupted / orchestration start | 1 | 2572.051 | 2572.051 | 2572.051 | undefined | omitted | omitted | 2572.051 |
| replication_interrupted / action completion | 1 | 2336.497 | 2336.497 | 2336.497 | undefined | omitted | omitted | 2336.497 |
| target_reachable_after_repair / orchestration start | 1 | 5882.436 | 5882.436 | 5882.436 | undefined | omitted | omitted | 5882.436 |
| target_reachable_after_repair / action completion | 1 | 5646.882 | 5646.882 | 5646.882 | undefined | omitted | omitted | 5646.882 |
| membership_rejoined / orchestration start | 1 | 5882.436 | 5882.436 | 5882.436 | undefined | omitted | omitted | 5882.436 |
| membership_rejoined / action completion | 1 | 5646.882 | 5646.882 | 5646.882 | undefined | omitted | omitted | 5646.882 |
| replica_link_up / orchestration start | 1 | 9247.574 | 9247.574 | 9247.574 | undefined | omitted | omitted | 9247.574 |
| replica_link_up / action completion | 1 | 9012.019 | 9012.019 | 9012.019 | undefined | omitted | omitted | 9012.019 |
| replica_sync_complete / orchestration start | 1 | 9247.574 | 9247.574 | 9247.574 | undefined | omitted | omitted | 9247.574 |
| replica_sync_complete / action completion | 1 | 9012.019 | 9012.019 | 9012.019 | undefined | omitted | omitted | 9012.019 |
| stable_redundancy / orchestration start | 1 | 10523.834 | 10523.834 | 10523.834 | undefined | omitted | omitted | 10523.834 |
| stable_redundancy / action completion | 1 | 10288.280 | 10288.280 | 10288.280 | undefined | omitted | omitted | 10288.280 |
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
| baseline | 1 | 803 | 803 | 0 | 100.00% | 199.775 | 199.775 |
| fault | 1 | 494 | 378 | 116 | 76.52% | 113.778 | 87.061 |
| post_recovery | 1 | 401 | 401 | 0 | 100.00% | 200.462 | 200.462 |
| reconciliation | 1 | 936 | 834 | 102 | 89.10% | 199.901 | 178.117 |
| repair | 1 | 307 | 259 | 48 | 84.36% | 199.699 | 168.476 |

Whole-workload schedule counters measured in 1 trials; legacy or incomplete counters unmeasured in 0. Offered: 3316; started: 2941; completed: 2941; dropped: 375. Started fraction of offered: 88.69%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| baseline / completed ops/s | 1 | 199.775 | 199.775 | 199.775 | undefined | omitted | omitted | 199.775 |
| baseline / successful request p50 ms | 1 | 1.252 | 1.252 | 1.252 | undefined | omitted | omitted | 1.252 |
| baseline / successful request p95 ms | 1 | 2.732 | 2.732 | 2.732 | undefined | omitted | omitted | 2.732 |
| baseline / successful request p99 ms | 1 | 3.321 | 3.321 | 3.321 | undefined | omitted | omitted | 3.321 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 1 | 0.765 | 0.765 | 0.765 | undefined | omitted | omitted | 0.765 |
| fault / completed ops/s | 1 | 113.778 | 113.778 | 113.778 | undefined | omitted | omitted | 113.778 |
| fault / successful request p50 ms | 1 | 0.768 | 0.768 | 0.768 | undefined | omitted | omitted | 0.768 |
| fault / successful request p95 ms | 1 | 3.163 | 3.163 | 3.163 | undefined | omitted | omitted | 3.163 |
| fault / successful request p99 ms | 1 | 436.746 | 436.746 | 436.746 | undefined | omitted | omitted | 436.746 |
| fault / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| fault / failed request p95 ms | 1 | 500.792 | 500.792 | 500.792 | undefined | omitted | omitted | 500.792 |
| fault / failed request p99 ms | 1 | 500.792 | 500.792 | 500.792 | undefined | omitted | omitted | 500.792 |
| post_recovery / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| post_recovery / completed ops/s | 1 | 200.462 | 200.462 | 200.462 | undefined | omitted | omitted | 200.462 |
| post_recovery / successful request p50 ms | 1 | 1.192 | 1.192 | 1.192 | undefined | omitted | omitted | 1.192 |
| post_recovery / successful request p95 ms | 1 | 3.012 | 3.012 | 3.012 | undefined | omitted | omitted | 3.012 |
| post_recovery / successful request p99 ms | 1 | 6.263 | 6.263 | 6.263 | undefined | omitted | omitted | 6.263 |
| post_recovery / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / trial success fraction | 1 | 0.891 | 0.891 | 0.891 | undefined | omitted | omitted | 0.891 |
| reconciliation / completed ops/s | 1 | 199.901 | 199.901 | 199.901 | undefined | omitted | omitted | 199.901 |
| reconciliation / successful request p50 ms | 1 | 0.520 | 0.520 | 0.520 | undefined | omitted | omitted | 0.520 |
| reconciliation / successful request p95 ms | 1 | 2.478 | 2.478 | 2.478 | undefined | omitted | omitted | 2.478 |
| reconciliation / successful request p99 ms | 1 | 4.239 | 4.239 | 4.239 | undefined | omitted | omitted | 4.239 |
| reconciliation / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| reconciliation / failed request p95 ms | 1 | 128.972 | 128.972 | 128.972 | undefined | omitted | omitted | 128.972 |
| reconciliation / failed request p99 ms | 1 | 137.102 | 137.102 | 137.102 | undefined | omitted | omitted | 137.102 |
| repair / trial success fraction | 1 | 0.844 | 0.844 | 0.844 | undefined | omitted | omitted | 0.844 |
| repair / completed ops/s | 1 | 199.699 | 199.699 | 199.699 | undefined | omitted | omitted | 199.699 |
| repair / successful request p50 ms | 1 | 0.449 | 0.449 | 0.449 | undefined | omitted | omitted | 0.449 |
| repair / successful request p95 ms | 1 | 2.360 | 2.360 | 2.360 | undefined | omitted | omitted | 2.360 |
| repair / successful request p99 ms | 1 | 3.321 | 3.321 | 3.321 | undefined | omitted | omitted | 3.321 |
| repair / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| repair / failed request p95 ms | 1 | 135.421 | 135.421 | 135.421 | undefined | omitted | omitted | 135.421 |
| repair / failed request p99 ms | 1 | 135.591 | 135.591 | 135.591 | undefined | omitted | omitted | 135.591 |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 1 trials, unmeasured in 0. Every successful SET remains in the main validation ledger when the later acknowledgement command is short, errors, or is not attempted. Diagnostic SET/WAIT/WAITAOF stage latency remains in individual reports and does not increase logical request or pacing counts. Errors and unattempted commands supply unknown acknowledgement counts, never fabricated zero replies.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 417 | 0 | 417 | 417 | 0 | 0 | 0 |
| fault | 209 | 35 | 209 | 165 | 44 | 0 | 0 |
| post_recovery | 211 | 0 | 211 | 211 | 0 | 0 | 0 |
| reconciliation | 479 | 0 | 479 | 377 | 102 | 0 | 0 |
| repair | 146 | 0 | 146 | 98 | 48 | 0 | 0 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. Attempted WAIT client-call timeouts are counted separately above.

| Observed WAIT replica count | Replies |
|---:|---:|
| 0 | 194 |
| 1 | 1268 |

WAIT-satisfied subset validation: expected 1268; checked 1268; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Validation attempted in 1 trials. Complete means every subset record was checked; preservation also requires zero missing, corrupt and unreadable records.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 4 |
| fault | network | 31 |
| fault | timeout | 37 |
| fault | wait_short | 44 |
| reconciliation | wait_short | 102 |
| repair | wait_short | 48 |

## Acknowledged-write accounting

Validation attempted in 1 trials and complete in 1. SET attempts: 1497; acknowledged writes: 1462; ambiguous SET outcomes: 35.

Checked: 1462; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 1 | 1462.000 | 1462.000 | 1462.000 | undefined | omitted | omitted | 1462.000 |
| ambiguous writes | 1 | 35.000 | 35.000 | 35.000 | undefined | omitted | omitted | 35.000 |
| missing writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | 73a8d11d94ba9a7f56fb11b1 | recovered | true | true | [JSON](runs/73a8d11d94ba9a7f56fb11b1/report.json), [Markdown](runs/73a8d11d94ba9a7f56fb11b1/report.md) |

## Measurement methodology

- Trials execute sequentially with a fresh environment per callback; the next callback starts only after the previous callback returns with successful cleanup evidence (or an authoritative report that no environment was provisioned).
- Recovery success fraction uses all requested trials as the denominator. Outcomes, started/unstarted counts and individual reports preserve failed, partial and aborted trials; the started-trial fraction is reported separately.
- Recovery timings are conditional on observed milestones, including observed milestones in partial/failed trials. Unobserved post-fault events are censored at the accepted sampled observation horizon with deadline/abort/error reasons, never zero-duration samples; this does not prove the physical event occurred after that horizon. Trials without a requested fault are separately not at risk. No survival distribution, independence assumption or population reliability guarantee is inferred.
- Event durations use runtime SinceFaultMS and fault_action_duration_ms when present, preserving monotonic elapsed-time measurements. Schema 1 action-completion-relative durations fall back to wall-clock timestamp differences and can be affected by clock adjustments. First sampled observations include polling uncertainty; action duration includes ownership checks and fault command overhead.
- Min, median, mean, N-1 sample standard deviation and maximum describe available trial observations. P90/P95 use nearest rank only when N>=20; smaller samples have an explicit tail-estimation caveat.
- Completed-request success fractions weight actual completed counts; aggregate throughput divides summed completed/successful counts by summed phase wall-clock duration. Offered load and worker saturation are not inferred from completed requests.
- Whole-workload offered, started, completed and dropped schedule-slot counters are aggregated only for reports carrying explicit pacing methodology. Legacy absent counters are unmeasured, not zero. Offered=started+dropped and started=completed after workload quiescence; service probes and validation are excluded.
- Acknowledged SET writes remain distinct from ambiguous SET outcomes and final read-back validation. Skipped/incomplete validation contributes acknowledged-not-checked counts; no observed loss is not a durability guarantee.
- WAIT replica-memory acknowledgements and WAITAOF local/replica AOF acknowledgements are distinct policies. Each successful SET is accounted for by an attempted acknowledgement command or an explicit not-attempted count. Valid WAITAOF replies preserve paired local/replica counts and independently overlapping satisfaction counts; errors provide no acknowledgement observation. Diagnostic command stages do not increase logical request or pacing counts. A recovered policy trial additionally requires sustained policy availability and complete preservation of its acknowledged subset.
- Recovered schema 2 reports require the explicit combined-recovery flag. Legacy schema 1 reports can establish it through healthy topology/client flags and an observed combined-stability milestone, together with complete acknowledged-write validation and verified cleanup.
- Recovered primary-isolate trials additionally require sampled unchanged process survival while disconnected, restored connectivity, observed stable reconciliation and restored replication redundancy. Pre-restoration recovery alone does not establish a recovered network experiment.

Effective experiment configuration SHA-256: `d7e106ea5cb6011eb4736e886b40e76b334dd77ec754a2cae619035e090e1ef9`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
