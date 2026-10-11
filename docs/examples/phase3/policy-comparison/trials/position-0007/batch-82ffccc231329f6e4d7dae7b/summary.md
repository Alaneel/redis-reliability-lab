# Redis Reliability Lab — repeated trials

Batch: `batch-82ffccc231329f6e4d7dae7b`  
Outcome: **recovered**  
Started: 2026-10-11T09:15:27.137083Z  
Finished: 2026-10-11T09:16:05.357127Z

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
| fault action duration | 1 | 241.023 | 241.023 | 241.023 | undefined | omitted | omitted | 241.023 |
| failure_detected / orchestration start | 1 | 1813.686 | 1813.686 | 1813.686 | undefined | omitted | omitted | 1813.686 |
| failure_detected / action completion | 1 | 1572.663 | 1572.663 | 1572.663 | undefined | omitted | omitted | 1572.663 |
| replica_promoted / orchestration start | 1 | 3450.078 | 3450.078 | 3450.078 | undefined | omitted | omitted | 3450.078 |
| replica_promoted / action completion | 1 | 3209.055 | 3209.055 | 3209.055 | undefined | omitted | omitted | 3209.055 |
| full_slot_coverage / orchestration start | 1 | 3450.078 | 3450.078 | 3450.078 | undefined | omitted | omitted | 3450.078 |
| full_slot_coverage / action completion | 1 | 3209.055 | 3209.055 | 3209.055 | undefined | omitted | omitted | 3209.055 |
| topology_converged / orchestration start | 1 | 3450.078 | 3450.078 | 3450.078 | undefined | omitted | omitted | 3450.078 |
| topology_converged / action completion | 1 | 3209.055 | 3209.055 | 3209.055 | undefined | omitted | omitted | 3209.055 |
| client_first_all_shard_success / orchestration start | 1 | 3077.328 | 3077.328 | 3077.328 | undefined | omitted | omitted | 3077.328 |
| client_first_all_shard_success / action completion | 1 | 2836.305 | 2836.305 | 2836.305 | undefined | omitted | omitted | 2836.305 |
| client_stable / orchestration start | 1 | 4254.828 | 4254.828 | 4254.828 | undefined | omitted | omitted | 4254.828 |
| client_stable / action completion | 1 | 4013.804 | 4013.804 | 4013.804 | undefined | omitted | omitted | 4013.804 |
| combined_stable / orchestration start | 1 | 4638.861 | 4638.861 | 4638.861 | undefined | omitted | omitted | 4638.861 |
| combined_stable / action completion | 1 | 4397.838 | 4397.838 | 4397.838 | undefined | omitted | omitted | 4397.838 |
| write_policy_stable / orchestration start | 1 | 11888.643 | 11888.643 | 11888.643 | undefined | omitted | omitted | 11888.643 |
| write_policy_stable / action completion | 1 | 11647.620 | 11647.620 | 11647.620 | undefined | omitted | omitted | 11647.620 |
| replication_interrupted / orchestration start | 1 | 3450.078 | 3450.078 | 3450.078 | undefined | omitted | omitted | 3450.078 |
| replication_interrupted / action completion | 1 | 3209.055 | 3209.055 | 3209.055 | undefined | omitted | omitted | 3209.055 |
| target_reachable_after_repair / orchestration start | 1 | 6253.872 | 6253.872 | 6253.872 | undefined | omitted | omitted | 6253.872 |
| target_reachable_after_repair / action completion | 1 | 6012.849 | 6012.849 | 6012.849 | undefined | omitted | omitted | 6012.849 |
| membership_rejoined / orchestration start | 1 | 6253.872 | 6253.872 | 6253.872 | undefined | omitted | omitted | 6253.872 |
| membership_rejoined / action completion | 1 | 6012.849 | 6012.849 | 6012.849 | undefined | omitted | omitted | 6012.849 |
| replica_link_up / orchestration start | 1 | 10913.946 | 10913.946 | 10913.946 | undefined | omitted | omitted | 10913.946 |
| replica_link_up / action completion | 1 | 10672.922 | 10672.922 | 10672.922 | undefined | omitted | omitted | 10672.922 |
| replica_sync_complete / orchestration start | 1 | 10913.946 | 10913.946 | 10913.946 | undefined | omitted | omitted | 10913.946 |
| replica_sync_complete / action completion | 1 | 10672.922 | 10672.922 | 10672.922 | undefined | omitted | omitted | 10672.922 |
| stable_redundancy / orchestration start | 1 | 12200.501 | 12200.501 | 12200.501 | undefined | omitted | omitted | 12200.501 |
| stable_redundancy / action completion | 1 | 11959.478 | 11959.478 | 11959.478 | undefined | omitted | omitted | 11959.478 |
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
| baseline | 1 | 803 | 803 | 0 | 100.00% | 199.810 | 199.810 |
| fault | 1 | 575 | 448 | 127 | 77.91% | 121.016 | 94.287 |
| post_recovery | 1 | 400 | 400 | 0 | 100.00% | 199.965 | 199.965 |
| reconciliation | 1 | 1198 | 1051 | 147 | 87.73% | 200.133 | 175.575 |
| repair | 1 | 300 | 247 | 53 | 82.33% | 200.037 | 164.697 |

Whole-workload schedule counters measured in 1 trials; legacy or incomplete counters unmeasured in 0. Offered: 3651; started: 3276; completed: 3276; dropped: 375. Started fraction of offered: 89.73%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| baseline / completed ops/s | 1 | 199.810 | 199.810 | 199.810 | undefined | omitted | omitted | 199.810 |
| baseline / successful request p50 ms | 1 | 1.314 | 1.314 | 1.314 | undefined | omitted | omitted | 1.314 |
| baseline / successful request p95 ms | 1 | 2.478 | 2.478 | 2.478 | undefined | omitted | omitted | 2.478 |
| baseline / successful request p99 ms | 1 | 3.163 | 3.163 | 3.163 | undefined | omitted | omitted | 3.163 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 1 | 0.779 | 0.779 | 0.779 | undefined | omitted | omitted | 0.779 |
| fault / completed ops/s | 1 | 121.016 | 121.016 | 121.016 | undefined | omitted | omitted | 121.016 |
| fault / successful request p50 ms | 1 | 0.697 | 0.697 | 0.697 | undefined | omitted | omitted | 0.697 |
| fault / successful request p95 ms | 1 | 3.012 | 3.012 | 3.012 | undefined | omitted | omitted | 3.012 |
| fault / successful request p99 ms | 1 | 231.616 | 231.616 | 231.616 | undefined | omitted | omitted | 231.616 |
| fault / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| fault / failed request p95 ms | 1 | 501.501 | 501.501 | 501.501 | undefined | omitted | omitted | 501.501 |
| fault / failed request p99 ms | 1 | 501.501 | 501.501 | 501.501 | undefined | omitted | omitted | 501.501 |
| post_recovery / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| post_recovery / completed ops/s | 1 | 199.965 | 199.965 | 199.965 | undefined | omitted | omitted | 199.965 |
| post_recovery / successful request p50 ms | 1 | 0.934 | 0.934 | 0.934 | undefined | omitted | omitted | 0.934 |
| post_recovery / successful request p95 ms | 1 | 2.869 | 2.869 | 2.869 | undefined | omitted | omitted | 2.869 |
| post_recovery / successful request p99 ms | 1 | 3.845 | 3.845 | 3.845 | undefined | omitted | omitted | 3.845 |
| post_recovery / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / trial success fraction | 1 | 0.877 | 0.877 | 0.877 | undefined | omitted | omitted | 0.877 |
| reconciliation / completed ops/s | 1 | 200.133 | 200.133 | 200.133 | undefined | omitted | omitted | 200.133 |
| reconciliation / successful request p50 ms | 1 | 0.472 | 0.472 | 0.472 | undefined | omitted | omitted | 0.472 |
| reconciliation / successful request p95 ms | 1 | 2.732 | 2.732 | 2.732 | undefined | omitted | omitted | 2.732 |
| reconciliation / successful request p99 ms | 1 | 3.487 | 3.487 | 3.487 | undefined | omitted | omitted | 3.487 |
| reconciliation / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| reconciliation / failed request p95 ms | 1 | 122.831 | 122.831 | 122.831 | undefined | omitted | omitted | 122.831 |
| reconciliation / failed request p99 ms | 1 | 135.421 | 135.421 | 135.421 | undefined | omitted | omitted | 135.421 |
| repair / trial success fraction | 1 | 0.823 | 0.823 | 0.823 | undefined | omitted | omitted | 0.823 |
| repair / completed ops/s | 1 | 200.037 | 200.037 | 200.037 | undefined | omitted | omitted | 200.037 |
| repair / successful request p50 ms | 1 | 0.495 | 0.495 | 0.495 | undefined | omitted | omitted | 0.495 |
| repair / successful request p95 ms | 1 | 2.248 | 2.248 | 2.248 | undefined | omitted | omitted | 2.248 |
| repair / successful request p99 ms | 1 | 3.662 | 3.662 | 3.662 | undefined | omitted | omitted | 3.662 |
| repair / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| repair / failed request p95 ms | 1 | 128.972 | 128.972 | 128.972 | undefined | omitted | omitted | 128.972 |
| repair / failed request p99 ms | 1 | 135.223 | 135.223 | 135.223 | undefined | omitted | omitted | 135.223 |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 1 trials, unmeasured in 0. Every successful SET remains in the main validation ledger when the later acknowledgement command is short, errors, or is not attempted. Diagnostic SET/WAIT/WAITAOF stage latency remains in individual reports and does not increase logical request or pacing counts. Errors and unattempted commands supply unknown acknowledgement counts, never fabricated zero replies.

WAITAOF requires local AOF acknowledgement 1 and replica AOF acknowledgement 1. Local and replica satisfied counts overlap; they are independently reconciled against paired valid replies.

| Phase | SET acknowledged | SET errors | WAITAOF attempted | Both satisfied | Short | Errors | Client timeouts | Local satisfied | Replica satisfied |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline | 416 | 0 | 416 | 416 | 0 | 0 | 0 | 416 | 416 |
| fault | 240 | 33 | 240 | 185 | 55 | 0 | 0 | 240 | 185 |
| post_recovery | 196 | 0 | 196 | 196 | 0 | 0 | 0 | 196 | 196 |
| reconciliation | 619 | 0 | 619 | 472 | 147 | 0 | 0 | 619 | 472 |
| repair | 158 | 0 | 158 | 105 | 53 | 0 | 0 | 158 | 105 |

SET acknowledgements without a subsequent WAITAOF attempt: 0; client timeouts among those: 0. Attempted WAITAOF client-call timeouts are counted separately above. Total local satisfied: 1629; replica satisfied: 1374; both satisfied: 1374.

| Observed WAITAOF local/replicas pair | Replies |
|---|---:|
| 1/0 | 255 |
| 1/1 | 1374 |

WAITAOF-satisfied subset validation: expected 1374; checked 1374; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Validation attempted in 1 trials. Complete means every subset record was checked; preservation also requires zero missing, corrupt and unreadable records.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 4 |
| fault | network | 29 |
| fault | timeout | 39 |
| fault | waitaof_short | 55 |
| reconciliation | waitaof_short | 147 |
| repair | waitaof_short | 53 |

## Acknowledged-write accounting

Validation attempted in 1 trials and complete in 1. SET attempts: 1662; acknowledged writes: 1629; ambiguous SET outcomes: 33.

Checked: 1629; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 1 | 1629.000 | 1629.000 | 1629.000 | undefined | omitted | omitted | 1629.000 |
| ambiguous writes | 1 | 33.000 | 33.000 | 33.000 | undefined | omitted | omitted | 33.000 |
| missing writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | 37475e6b0cf7b098669c0663 | recovered | true | true | [JSON](runs/37475e6b0cf7b098669c0663/report.json), [Markdown](runs/37475e6b0cf7b098669c0663/report.md) |

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

Effective experiment configuration SHA-256: `975080021092fa9a75ee5da95f13abfb2ac4db2ce23d3e320a4867ae224e9637`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
