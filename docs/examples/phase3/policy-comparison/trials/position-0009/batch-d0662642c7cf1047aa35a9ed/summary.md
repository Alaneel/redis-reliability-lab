# Redis Reliability Lab — repeated trials

Batch: `batch-d0662642c7cf1047aa35a9ed`  
Outcome: **recovered**  
Started: 2026-10-11T09:16:42.829065Z  
Finished: 2026-10-11T09:17:22.331474Z

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
| fault action duration | 1 | 236.912 | 236.912 | 236.912 | undefined | omitted | omitted | 236.912 |
| failure_detected / orchestration start | 1 | 1804.405 | 1804.405 | 1804.405 | undefined | omitted | omitted | 1804.405 |
| failure_detected / action completion | 1 | 1567.493 | 1567.493 | 1567.493 | undefined | omitted | omitted | 1567.493 |
| replica_promoted / orchestration start | 1 | 2588.466 | 2588.466 | 2588.466 | undefined | omitted | omitted | 2588.466 |
| replica_promoted / action completion | 1 | 2351.554 | 2351.554 | 2351.554 | undefined | omitted | omitted | 2351.554 |
| full_slot_coverage / orchestration start | 1 | 2588.466 | 2588.466 | 2588.466 | undefined | omitted | omitted | 2588.466 |
| full_slot_coverage / action completion | 1 | 2351.554 | 2351.554 | 2351.554 | undefined | omitted | omitted | 2351.554 |
| topology_converged / orchestration start | 1 | 2588.466 | 2588.466 | 2588.466 | undefined | omitted | omitted | 2588.466 |
| topology_converged / action completion | 1 | 2351.554 | 2351.554 | 2351.554 | undefined | omitted | omitted | 2351.554 |
| client_first_all_shard_success / orchestration start | 1 | 3398.085 | 3398.085 | 3398.085 | undefined | omitted | omitted | 3398.085 |
| client_first_all_shard_success / action completion | 1 | 3161.173 | 3161.173 | 3161.173 | undefined | omitted | omitted | 3161.173 |
| client_stable / orchestration start | 1 | 4522.179 | 4522.179 | 4522.179 | undefined | omitted | omitted | 4522.179 |
| client_stable / action completion | 1 | 4285.267 | 4285.267 | 4285.267 | undefined | omitted | omitted | 4285.267 |
| combined_stable / orchestration start | 1 | 4522.179 | 4522.179 | 4522.179 | undefined | omitted | omitted | 4522.179 |
| combined_stable / action completion | 1 | 4285.267 | 4285.267 | 4285.267 | undefined | omitted | omitted | 4285.267 |
| write_policy_stable / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| write_policy_stable / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replication_interrupted / orchestration start | 1 | 2588.466 | 2588.466 | 2588.466 | undefined | omitted | omitted | 2588.466 |
| replication_interrupted / action completion | 1 | 2351.554 | 2351.554 | 2351.554 | undefined | omitted | omitted | 2351.554 |
| target_reachable_after_repair / orchestration start | 1 | 6039.703 | 6039.703 | 6039.703 | undefined | omitted | omitted | 6039.703 |
| target_reachable_after_repair / action completion | 1 | 5802.791 | 5802.791 | 5802.791 | undefined | omitted | omitted | 5802.791 |
| membership_rejoined / orchestration start | 1 | 6039.703 | 6039.703 | 6039.703 | undefined | omitted | omitted | 6039.703 |
| membership_rejoined / action completion | 1 | 5802.791 | 5802.791 | 5802.791 | undefined | omitted | omitted | 5802.791 |
| replica_link_up / orchestration start | 1 | 10733.578 | 10733.578 | 10733.578 | undefined | omitted | omitted | 10733.578 |
| replica_link_up / action completion | 1 | 10496.666 | 10496.666 | 10496.666 | undefined | omitted | omitted | 10496.666 |
| replica_sync_complete / orchestration start | 1 | 10733.578 | 10733.578 | 10733.578 | undefined | omitted | omitted | 10733.578 |
| replica_sync_complete / action completion | 1 | 10496.666 | 10496.666 | 10496.666 | undefined | omitted | omitted | 10496.666 |
| stable_redundancy / orchestration start | 1 | 12013.147 | 12013.147 | 12013.147 | undefined | omitted | omitted | 12013.147 |
| stable_redundancy / action completion | 1 | 11776.236 | 11776.236 | 11776.236 | undefined | omitted | omitted | 11776.236 |
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
| baseline | 1 | 804 | 804 | 0 | 100.00% | 199.947 | 199.947 |
| fault | 1 | 543 | 465 | 78 | 85.64% | 119.974 | 102.740 |
| post_recovery | 1 | 400 | 400 | 0 | 100.00% | 199.944 | 199.944 |
| reconciliation | 1 | 1202 | 1202 | 0 | 100.00% | 199.956 | 199.956 |
| repair | 1 | 302 | 302 | 0 | 100.00% | 199.943 | 199.943 |

Whole-workload schedule counters measured in 1 trials; legacy or incomplete counters unmeasured in 0. Offered: 3613; started: 3251; completed: 3251; dropped: 362. Started fraction of offered: 89.98%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| baseline / completed ops/s | 1 | 199.947 | 199.947 | 199.947 | undefined | omitted | omitted | 199.947 |
| baseline / successful request p50 ms | 1 | 0.807 | 0.807 | 0.807 | undefined | omitted | omitted | 0.807 |
| baseline / successful request p95 ms | 1 | 1.380 | 1.380 | 1.380 | undefined | omitted | omitted | 1.380 |
| baseline / successful request p99 ms | 1 | 2.360 | 2.360 | 2.360 | undefined | omitted | omitted | 2.360 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 1 | 0.856 | 0.856 | 0.856 | undefined | omitted | omitted | 0.856 |
| fault / completed ops/s | 1 | 119.974 | 119.974 | 119.974 | undefined | omitted | omitted | 119.974 |
| fault / successful request p50 ms | 1 | 0.697 | 0.697 | 0.697 | undefined | omitted | omitted | 0.697 |
| fault / successful request p95 ms | 1 | 2.360 | 2.360 | 2.360 | undefined | omitted | omitted | 2.360 |
| fault / successful request p99 ms | 1 | 325.906 | 325.906 | 325.906 | undefined | omitted | omitted | 325.906 |
| fault / failed request p50 ms | 1 | 0.732 | 0.732 | 0.732 | undefined | omitted | omitted | 0.732 |
| fault / failed request p95 ms | 1 | 501.118 | 501.118 | 501.118 | undefined | omitted | omitted | 501.118 |
| fault / failed request p99 ms | 1 | 501.118 | 501.118 | 501.118 | undefined | omitted | omitted | 501.118 |
| post_recovery / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| post_recovery / completed ops/s | 1 | 199.944 | 199.944 | 199.944 | undefined | omitted | omitted | 199.944 |
| post_recovery / successful request p50 ms | 1 | 0.697 | 0.697 | 0.697 | undefined | omitted | omitted | 0.697 |
| post_recovery / successful request p95 ms | 1 | 2.039 | 2.039 | 2.039 | undefined | omitted | omitted | 2.039 |
| post_recovery / successful request p99 ms | 1 | 2.478 | 2.478 | 2.478 | undefined | omitted | omitted | 2.478 |
| post_recovery / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| reconciliation / completed ops/s | 1 | 199.956 | 199.956 | 199.956 | undefined | omitted | omitted | 199.956 |
| reconciliation / successful request p50 ms | 1 | 0.632 | 0.632 | 0.632 | undefined | omitted | omitted | 0.632 |
| reconciliation / successful request p95 ms | 1 | 1.192 | 1.192 | 1.192 | undefined | omitted | omitted | 1.192 |
| reconciliation / successful request p99 ms | 1 | 1.677 | 1.677 | 1.677 | undefined | omitted | omitted | 1.677 |
| reconciliation / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| repair / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| repair / completed ops/s | 1 | 199.943 | 199.943 | 199.943 | undefined | omitted | omitted | 199.943 |
| repair / successful request p50 ms | 1 | 0.602 | 0.602 | 0.602 | undefined | omitted | omitted | 0.602 |
| repair / successful request p95 ms | 1 | 1.314 | 1.314 | 1.314 | undefined | omitted | omitted | 1.314 |
| repair / successful request p99 ms | 1 | 2.602 | 2.602 | 2.602 | undefined | omitted | omitted | 2.602 |
| repair / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| repair / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| repair / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 1 trials, unmeasured in 0. Every successful SET remains in the main validation ledger when the later acknowledgement command is short, errors, or is not attempted. Diagnostic SET/WAIT/WAITAOF stage latency remains in individual reports and does not increase logical request or pacing counts. Errors and unattempted commands supply unknown acknowledgement counts, never fabricated zero replies.

| Phase | SET acknowledged | SET errors | Policy satisfied |
|---|---:|---:|---:|
| baseline | 417 | 0 | 417 |
| fault | 217 | 44 | 217 |
| post_recovery | 195 | 0 | 195 |
| reconciliation | 634 | 0 | 634 |
| repair | 144 | 0 | 144 |

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 14 |
| fault | network | 30 |
| fault | timeout | 34 |

## Acknowledged-write accounting

Validation attempted in 1 trials and complete in 1. SET attempts: 1651; acknowledged writes: 1607; ambiguous SET outcomes: 44.

Checked: 1607; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 1 | 1607.000 | 1607.000 | 1607.000 | undefined | omitted | omitted | 1607.000 |
| ambiguous writes | 1 | 44.000 | 44.000 | 44.000 | undefined | omitted | omitted | 44.000 |
| missing writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | 8d7edf6681a3f591879efeda | recovered | true | true | [JSON](runs/8d7edf6681a3f591879efeda/report.json), [Markdown](runs/8d7edf6681a3f591879efeda/report.md) |

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
