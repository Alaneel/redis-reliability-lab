# Redis Reliability Lab — repeated trials

Batch: `batch-bb39bfe07e9868b38d98ea5d`  
Outcome: **recovered**  
Started: 2026-10-11T09:12:17.765962Z  
Finished: 2026-10-11T09:12:55.638533Z

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
| write_policy_stable | 0 | 0 | 0 | 1 | 0 | 0 |
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
| fault action duration | 1 | 242.783 | 242.783 | 242.783 | undefined | omitted | omitted | 242.783 |
| failure_detected / orchestration start | 1 | 1813.866 | 1813.866 | 1813.866 | undefined | omitted | omitted | 1813.866 |
| failure_detected / action completion | 1 | 1571.084 | 1571.084 | 1571.084 | undefined | omitted | omitted | 1571.084 |
| replica_promoted / orchestration start | 1 | 3383.025 | 3383.025 | 3383.025 | undefined | omitted | omitted | 3383.025 |
| replica_promoted / action completion | 1 | 3140.242 | 3140.242 | 3140.242 | undefined | omitted | omitted | 3140.242 |
| full_slot_coverage / orchestration start | 1 | 3383.025 | 3383.025 | 3383.025 | undefined | omitted | omitted | 3383.025 |
| full_slot_coverage / action completion | 1 | 3140.242 | 3140.242 | 3140.242 | undefined | omitted | omitted | 3140.242 |
| topology_converged / orchestration start | 1 | 3383.025 | 3383.025 | 3383.025 | undefined | omitted | omitted | 3383.025 |
| topology_converged / action completion | 1 | 3140.242 | 3140.242 | 3140.242 | undefined | omitted | omitted | 3140.242 |
| client_first_all_shard_success / orchestration start | 1 | 3405.680 | 3405.680 | 3405.680 | undefined | omitted | omitted | 3405.680 |
| client_first_all_shard_success / action completion | 1 | 3162.897 | 3162.897 | 3162.897 | undefined | omitted | omitted | 3162.897 |
| client_stable / orchestration start | 1 | 4523.550 | 4523.550 | 4523.550 | undefined | omitted | omitted | 4523.550 |
| client_stable / action completion | 1 | 4280.768 | 4280.768 | 4280.768 | undefined | omitted | omitted | 4280.768 |
| combined_stable / orchestration start | 1 | 4523.550 | 4523.550 | 4523.550 | undefined | omitted | omitted | 4523.550 |
| combined_stable / action completion | 1 | 4280.768 | 4280.768 | 4280.768 | undefined | omitted | omitted | 4280.768 |
| write_policy_stable / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| write_policy_stable / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replication_interrupted / orchestration start | 1 | 3383.025 | 3383.025 | 3383.025 | undefined | omitted | omitted | 3383.025 |
| replication_interrupted / action completion | 1 | 3140.242 | 3140.242 | 3140.242 | undefined | omitted | omitted | 3140.242 |
| target_reachable_after_repair / orchestration start | 1 | 6038.904 | 6038.904 | 6038.904 | undefined | omitted | omitted | 6038.904 |
| target_reachable_after_repair / action completion | 1 | 5796.121 | 5796.121 | 5796.121 | undefined | omitted | omitted | 5796.121 |
| membership_rejoined / orchestration start | 1 | 6038.904 | 6038.904 | 6038.904 | undefined | omitted | omitted | 6038.904 |
| membership_rejoined / action completion | 1 | 5796.121 | 5796.121 | 5796.121 | undefined | omitted | omitted | 5796.121 |
| replica_link_up / orchestration start | 1 | 10105.573 | 10105.573 | 10105.573 | undefined | omitted | omitted | 10105.573 |
| replica_link_up / action completion | 1 | 9862.790 | 9862.790 | 9862.790 | undefined | omitted | omitted | 9862.790 |
| replica_sync_complete / orchestration start | 1 | 10105.573 | 10105.573 | 10105.573 | undefined | omitted | omitted | 10105.573 |
| replica_sync_complete / action completion | 1 | 9862.790 | 9862.790 | 9862.790 | undefined | omitted | omitted | 9862.790 |
| stable_redundancy / orchestration start | 1 | 11386.981 | 11386.981 | 11386.981 | undefined | omitted | omitted | 11386.981 |
| stable_redundancy / action completion | 1 | 11144.198 | 11144.198 | 11144.198 | undefined | omitted | omitted | 11144.198 |
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
| baseline | 1 | 803 | 803 | 0 | 100.00% | 199.764 | 199.764 |
| fault | 1 | 576 | 485 | 91 | 84.20% | 127.239 | 107.137 |
| post_recovery | 1 | 400 | 400 | 0 | 100.00% | 199.983 | 199.983 |
| reconciliation | 1 | 1077 | 1077 | 0 | 100.00% | 199.954 | 199.954 |
| repair | 1 | 302 | 302 | 0 | 100.00% | 200.117 | 200.117 |

Whole-workload schedule counters measured in 1 trials; legacy or incomplete counters unmeasured in 0. Offered: 3488; started: 3158; completed: 3158; dropped: 330. Started fraction of offered: 90.54%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| baseline / completed ops/s | 1 | 199.764 | 199.764 | 199.764 | undefined | omitted | omitted | 199.764 |
| baseline / successful request p50 ms | 1 | 0.664 | 0.664 | 0.664 | undefined | omitted | omitted | 0.664 |
| baseline / successful request p95 ms | 1 | 1.252 | 1.252 | 1.252 | undefined | omitted | omitted | 1.252 |
| baseline / successful request p99 ms | 1 | 2.248 | 2.248 | 2.248 | undefined | omitted | omitted | 2.248 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 1 | 0.842 | 0.842 | 0.842 | undefined | omitted | omitted | 0.842 |
| fault / completed ops/s | 1 | 127.239 | 127.239 | 127.239 | undefined | omitted | omitted | 127.239 |
| fault / successful request p50 ms | 1 | 0.732 | 0.732 | 0.732 | undefined | omitted | omitted | 0.732 |
| fault / successful request p95 ms | 1 | 1.942 | 1.942 | 1.942 | undefined | omitted | omitted | 1.942 |
| fault / successful request p99 ms | 1 | 149.301 | 149.301 | 149.301 | undefined | omitted | omitted | 149.301 |
| fault / failed request p50 ms | 1 | 0.732 | 0.732 | 0.732 | undefined | omitted | omitted | 0.732 |
| fault / failed request p95 ms | 1 | 500.539 | 500.539 | 500.539 | undefined | omitted | omitted | 500.539 |
| fault / failed request p99 ms | 1 | 500.539 | 500.539 | 500.539 | undefined | omitted | omitted | 500.539 |
| post_recovery / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| post_recovery / completed ops/s | 1 | 199.983 | 199.983 | 199.983 | undefined | omitted | omitted | 199.983 |
| post_recovery / successful request p50 ms | 1 | 0.573 | 0.573 | 0.573 | undefined | omitted | omitted | 0.573 |
| post_recovery / successful request p95 ms | 1 | 1.252 | 1.252 | 1.252 | undefined | omitted | omitted | 1.252 |
| post_recovery / successful request p99 ms | 1 | 2.141 | 2.141 | 2.141 | undefined | omitted | omitted | 2.141 |
| post_recovery / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| reconciliation / completed ops/s | 1 | 199.954 | 199.954 | 199.954 | undefined | omitted | omitted | 199.954 |
| reconciliation / successful request p50 ms | 1 | 0.632 | 0.632 | 0.632 | undefined | omitted | omitted | 0.632 |
| reconciliation / successful request p95 ms | 1 | 1.192 | 1.192 | 1.192 | undefined | omitted | omitted | 1.192 |
| reconciliation / successful request p99 ms | 1 | 2.141 | 2.141 | 2.141 | undefined | omitted | omitted | 2.141 |
| reconciliation / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| repair / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| repair / completed ops/s | 1 | 200.117 | 200.117 | 200.117 | undefined | omitted | omitted | 200.117 |
| repair / successful request p50 ms | 1 | 0.664 | 0.664 | 0.664 | undefined | omitted | omitted | 0.664 |
| repair / successful request p95 ms | 1 | 1.314 | 1.314 | 1.314 | undefined | omitted | omitted | 1.314 |
| repair / successful request p99 ms | 1 | 1.849 | 1.849 | 1.849 | undefined | omitted | omitted | 1.849 |
| repair / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| repair / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| repair / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 1 trials, unmeasured in 0. Every successful SET remains in the main validation ledger when the later acknowledgement command is short, errors, or is not attempted. Diagnostic SET/WAIT/WAITAOF stage latency remains in individual reports and does not increase logical request or pacing counts. Errors and unattempted commands supply unknown acknowledgement counts, never fabricated zero replies.

| Phase | SET acknowledged | SET errors | Policy satisfied |
|---|---:|---:|---:|
| baseline | 416 | 0 | 416 |
| fault | 230 | 52 | 230 |
| post_recovery | 190 | 0 | 190 |
| reconciliation | 556 | 0 | 556 |
| repair | 153 | 0 | 153 |

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 15 |
| fault | network | 37 |
| fault | timeout | 39 |

## Acknowledged-write accounting

Validation attempted in 1 trials and complete in 1. SET attempts: 1597; acknowledged writes: 1545; ambiguous SET outcomes: 52.

Checked: 1545; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 1 | 1545.000 | 1545.000 | 1545.000 | undefined | omitted | omitted | 1545.000 |
| ambiguous writes | 1 | 52.000 | 52.000 | 52.000 | undefined | omitted | omitted | 52.000 |
| missing writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | dab1232736a2b59ad498e657 | recovered | true | true | [JSON](runs/dab1232736a2b59ad498e657/report.json), [Markdown](runs/dab1232736a2b59ad498e657/report.md) |

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

Effective experiment configuration SHA-256: `de05e21df2ed5b0a82dc0af8fa135f9173da7c1fdf43019ed7cd93af6d12f3cf`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
