# Redis Reliability Lab — repeated trials

Batch: `batch-05a52c53dbeff938cd618ed7`  
Outcome: **recovered**  
Started: 2026-10-11T09:14:10.864501Z  
Finished: 2026-10-11T09:14:49.076364Z

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
| fault action duration | 1 | 249.882 | 249.882 | 249.882 | undefined | omitted | omitted | 249.882 |
| failure_detected / orchestration start | 1 | 1823.702 | 1823.702 | 1823.702 | undefined | omitted | omitted | 1823.702 |
| failure_detected / action completion | 1 | 1573.819 | 1573.819 | 1573.819 | undefined | omitted | omitted | 1573.819 |
| replica_promoted / orchestration start | 1 | 2611.642 | 2611.642 | 2611.642 | undefined | omitted | omitted | 2611.642 |
| replica_promoted / action completion | 1 | 2361.760 | 2361.760 | 2361.760 | undefined | omitted | omitted | 2361.760 |
| full_slot_coverage / orchestration start | 1 | 2611.642 | 2611.642 | 2611.642 | undefined | omitted | omitted | 2611.642 |
| full_slot_coverage / action completion | 1 | 2361.760 | 2361.760 | 2361.760 | undefined | omitted | omitted | 2361.760 |
| topology_converged / orchestration start | 1 | 2611.642 | 2611.642 | 2611.642 | undefined | omitted | omitted | 2611.642 |
| topology_converged / action completion | 1 | 2361.760 | 2361.760 | 2361.760 | undefined | omitted | omitted | 2361.760 |
| client_first_all_shard_success / orchestration start | 1 | 3531.250 | 3531.250 | 3531.250 | undefined | omitted | omitted | 3531.250 |
| client_first_all_shard_success / action completion | 1 | 3281.368 | 3281.368 | 3281.368 | undefined | omitted | omitted | 3281.368 |
| client_stable / orchestration start | 1 | 4712.726 | 4712.726 | 4712.726 | undefined | omitted | omitted | 4712.726 |
| client_stable / action completion | 1 | 4462.844 | 4462.844 | 4462.844 | undefined | omitted | omitted | 4462.844 |
| combined_stable / orchestration start | 1 | 4712.726 | 4712.726 | 4712.726 | undefined | omitted | omitted | 4712.726 |
| combined_stable / action completion | 1 | 4462.844 | 4462.844 | 4462.844 | undefined | omitted | omitted | 4462.844 |
| write_policy_stable / orchestration start | 1 | 11530.280 | 11530.280 | 11530.280 | undefined | omitted | omitted | 11530.280 |
| write_policy_stable / action completion | 1 | 11280.398 | 11280.398 | 11280.398 | undefined | omitted | omitted | 11280.398 |
| replication_interrupted / orchestration start | 1 | 2611.642 | 2611.642 | 2611.642 | undefined | omitted | omitted | 2611.642 |
| replication_interrupted / action completion | 1 | 2361.760 | 2361.760 | 2361.760 | undefined | omitted | omitted | 2361.760 |
| target_reachable_after_repair / orchestration start | 1 | 6381.267 | 6381.267 | 6381.267 | undefined | omitted | omitted | 6381.267 |
| target_reachable_after_repair / action completion | 1 | 6131.385 | 6131.385 | 6131.385 | undefined | omitted | omitted | 6131.385 |
| membership_rejoined / orchestration start | 1 | 6381.267 | 6381.267 | 6381.267 | undefined | omitted | omitted | 6381.267 |
| membership_rejoined / action completion | 1 | 6131.385 | 6131.385 | 6131.385 | undefined | omitted | omitted | 6131.385 |
| replica_link_up / orchestration start | 1 | 10572.642 | 10572.642 | 10572.642 | undefined | omitted | omitted | 10572.642 |
| replica_link_up / action completion | 1 | 10322.760 | 10322.760 | 10322.760 | undefined | omitted | omitted | 10322.760 |
| replica_sync_complete / orchestration start | 1 | 10572.642 | 10572.642 | 10572.642 | undefined | omitted | omitted | 10572.642 |
| replica_sync_complete / action completion | 1 | 10322.760 | 10322.760 | 10322.760 | undefined | omitted | omitted | 10322.760 |
| stable_redundancy / orchestration start | 1 | 11837.719 | 11837.719 | 11837.719 | undefined | omitted | omitted | 11837.719 |
| stable_redundancy / action completion | 1 | 11587.837 | 11587.837 | 11587.837 | undefined | omitted | omitted | 11587.837 |
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
| baseline | 1 | 803 | 803 | 0 | 100.00% | 199.931 | 199.931 |
| fault | 1 | 586 | 455 | 131 | 77.65% | 121.361 | 94.231 |
| post_recovery | 1 | 401 | 401 | 0 | 100.00% | 200.420 | 200.420 |
| reconciliation | 1 | 1099 | 971 | 128 | 88.35% | 200.001 | 176.707 |
| repair | 1 | 310 | 254 | 56 | 81.94% | 199.989 | 163.862 |

Whole-workload schedule counters measured in 1 trials; legacy or incomplete counters unmeasured in 0. Offered: 3578; started: 3199; completed: 3199; dropped: 379. Started fraction of offered: 89.41%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| baseline / completed ops/s | 1 | 199.931 | 199.931 | 199.931 | undefined | omitted | omitted | 199.931 |
| baseline / successful request p50 ms | 1 | 1.135 | 1.135 | 1.135 | undefined | omitted | omitted | 1.135 |
| baseline / successful request p95 ms | 1 | 2.602 | 2.602 | 2.602 | undefined | omitted | omitted | 2.602 |
| baseline / successful request p99 ms | 1 | 3.163 | 3.163 | 3.163 | undefined | omitted | omitted | 3.163 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 1 | 0.776 | 0.776 | 0.776 | undefined | omitted | omitted | 0.776 |
| fault / completed ops/s | 1 | 121.361 | 121.361 | 121.361 | undefined | omitted | omitted | 121.361 |
| fault / successful request p50 ms | 1 | 0.495 | 0.495 | 0.495 | undefined | omitted | omitted | 0.495 |
| fault / successful request p95 ms | 1 | 2.869 | 2.869 | 2.869 | undefined | omitted | omitted | 2.869 |
| fault / successful request p99 ms | 1 | 41.990 | 41.990 | 41.990 | undefined | omitted | omitted | 41.990 |
| fault / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| fault / failed request p95 ms | 1 | 500.586 | 500.586 | 500.586 | undefined | omitted | omitted | 500.586 |
| fault / failed request p99 ms | 1 | 500.586 | 500.586 | 500.586 | undefined | omitted | omitted | 500.586 |
| post_recovery / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| post_recovery / completed ops/s | 1 | 200.420 | 200.420 | 200.420 | undefined | omitted | omitted | 200.420 |
| post_recovery / successful request p50 ms | 1 | 0.981 | 0.981 | 0.981 | undefined | omitted | omitted | 0.981 |
| post_recovery / successful request p95 ms | 1 | 2.602 | 2.602 | 2.602 | undefined | omitted | omitted | 2.602 |
| post_recovery / successful request p99 ms | 1 | 3.163 | 3.163 | 3.163 | undefined | omitted | omitted | 3.163 |
| post_recovery / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / trial success fraction | 1 | 0.884 | 0.884 | 0.884 | undefined | omitted | omitted | 0.884 |
| reconciliation / completed ops/s | 1 | 200.001 | 200.001 | 200.001 | undefined | omitted | omitted | 200.001 |
| reconciliation / successful request p50 ms | 1 | 0.472 | 0.472 | 0.472 | undefined | omitted | omitted | 0.472 |
| reconciliation / successful request p95 ms | 1 | 2.478 | 2.478 | 2.478 | undefined | omitted | omitted | 2.478 |
| reconciliation / successful request p99 ms | 1 | 3.163 | 3.163 | 3.163 | undefined | omitted | omitted | 3.163 |
| reconciliation / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| reconciliation / failed request p95 ms | 1 | 128.972 | 128.972 | 128.972 | undefined | omitted | omitted | 128.972 |
| reconciliation / failed request p99 ms | 1 | 140.727 | 140.727 | 140.727 | undefined | omitted | omitted | 140.727 |
| repair / trial success fraction | 1 | 0.819 | 0.819 | 0.819 | undefined | omitted | omitted | 0.819 |
| repair / completed ops/s | 1 | 199.989 | 199.989 | 199.989 | undefined | omitted | omitted | 199.989 |
| repair / successful request p50 ms | 1 | 0.546 | 0.546 | 0.546 | undefined | omitted | omitted | 0.546 |
| repair / successful request p95 ms | 1 | 2.602 | 2.602 | 2.602 | undefined | omitted | omitted | 2.602 |
| repair / successful request p99 ms | 1 | 3.163 | 3.163 | 3.163 | undefined | omitted | omitted | 3.163 |
| repair / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| repair / failed request p95 ms | 1 | 128.972 | 128.972 | 128.972 | undefined | omitted | omitted | 128.972 |
| repair / failed request p99 ms | 1 | 130.697 | 130.697 | 130.697 | undefined | omitted | omitted | 130.697 |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 1 trials, unmeasured in 0. Every successful SET remains in the main validation ledger when the later acknowledgement command is short, errors, or is not attempted. Diagnostic SET/WAIT/WAITAOF stage latency remains in individual reports and does not increase logical request or pacing counts. Errors and unattempted commands supply unknown acknowledgement counts, never fabricated zero replies.

WAITAOF requires local AOF acknowledgement 1 and replica AOF acknowledgement 1. Local and replica satisfied counts overlap; they are independently reconciled against paired valid replies.

| Phase | SET acknowledged | SET errors | WAITAOF attempted | Both satisfied | Short | Errors | Client timeouts | Local satisfied | Replica satisfied |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline | 416 | 0 | 416 | 416 | 0 | 0 | 0 | 416 | 416 |
| fault | 236 | 36 | 236 | 181 | 55 | 0 | 0 | 236 | 181 |
| post_recovery | 199 | 0 | 199 | 199 | 0 | 0 | 0 | 199 | 199 |
| reconciliation | 565 | 0 | 565 | 437 | 128 | 0 | 0 | 565 | 437 |
| repair | 170 | 0 | 170 | 114 | 56 | 0 | 0 | 170 | 114 |

SET acknowledgements without a subsequent WAITAOF attempt: 0; client timeouts among those: 0. Attempted WAITAOF client-call timeouts are counted separately above. Total local satisfied: 1586; replica satisfied: 1347; both satisfied: 1347.

| Observed WAITAOF local/replicas pair | Replies |
|---|---:|
| 1/0 | 239 |
| 1/1 | 1347 |

WAITAOF-satisfied subset validation: expected 1347; checked 1347; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Validation attempted in 1 trials. Complete means every subset record was checked; preservation also requires zero missing, corrupt and unreadable records.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 6 |
| fault | network | 30 |
| fault | timeout | 40 |
| fault | waitaof_short | 55 |
| reconciliation | waitaof_short | 128 |
| repair | waitaof_short | 56 |

## Acknowledged-write accounting

Validation attempted in 1 trials and complete in 1. SET attempts: 1622; acknowledged writes: 1586; ambiguous SET outcomes: 36.

Checked: 1586; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 1 | 1586.000 | 1586.000 | 1586.000 | undefined | omitted | omitted | 1586.000 |
| ambiguous writes | 1 | 36.000 | 36.000 | 36.000 | undefined | omitted | omitted | 36.000 |
| missing writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | 688a7dc24b138ea24d1b7e0d | recovered | true | true | [JSON](runs/688a7dc24b138ea24d1b7e0d/report.json), [Markdown](runs/688a7dc24b138ea24d1b7e0d/report.md) |

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
