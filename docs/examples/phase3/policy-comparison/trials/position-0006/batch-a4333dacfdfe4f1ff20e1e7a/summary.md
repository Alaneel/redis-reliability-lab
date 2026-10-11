# Redis Reliability Lab — repeated trials

Batch: `batch-a4333dacfdfe4f1ff20e1e7a`  
Outcome: **recovered**  
Started: 2026-10-11T09:14:49.130591Z  
Finished: 2026-10-11T09:15:27.078957Z

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
| fault action duration | 1 | 254.839 | 254.839 | 254.839 | undefined | omitted | omitted | 254.839 |
| failure_detected / orchestration start | 1 | 1835.464 | 1835.464 | 1835.464 | undefined | omitted | omitted | 1835.464 |
| failure_detected / action completion | 1 | 1580.625 | 1580.625 | 1580.625 | undefined | omitted | omitted | 1580.625 |
| replica_promoted / orchestration start | 1 | 2625.178 | 2625.178 | 2625.178 | undefined | omitted | omitted | 2625.178 |
| replica_promoted / action completion | 1 | 2370.339 | 2370.339 | 2370.339 | undefined | omitted | omitted | 2370.339 |
| full_slot_coverage / orchestration start | 1 | 2625.178 | 2625.178 | 2625.178 | undefined | omitted | omitted | 2625.178 |
| full_slot_coverage / action completion | 1 | 2370.339 | 2370.339 | 2370.339 | undefined | omitted | omitted | 2370.339 |
| topology_converged / orchestration start | 1 | 2625.178 | 2625.178 | 2625.178 | undefined | omitted | omitted | 2625.178 |
| topology_converged / action completion | 1 | 2370.339 | 2370.339 | 2370.339 | undefined | omitted | omitted | 2370.339 |
| client_first_all_shard_success / orchestration start | 1 | 3531.589 | 3531.589 | 3531.589 | undefined | omitted | omitted | 3531.589 |
| client_first_all_shard_success / action completion | 1 | 3276.750 | 3276.750 | 3276.750 | undefined | omitted | omitted | 3276.750 |
| client_stable / orchestration start | 1 | 4717.017 | 4717.017 | 4717.017 | undefined | omitted | omitted | 4717.017 |
| client_stable / action completion | 1 | 4462.178 | 4462.178 | 4462.178 | undefined | omitted | omitted | 4462.178 |
| combined_stable / orchestration start | 1 | 4717.017 | 4717.017 | 4717.017 | undefined | omitted | omitted | 4717.017 |
| combined_stable / action completion | 1 | 4462.178 | 4462.178 | 4462.178 | undefined | omitted | omitted | 4462.178 |
| write_policy_stable / orchestration start | 1 | 11975.844 | 11975.844 | 11975.844 | undefined | omitted | omitted | 11975.844 |
| write_policy_stable / action completion | 1 | 11721.005 | 11721.005 | 11721.005 | undefined | omitted | omitted | 11721.005 |
| replication_interrupted / orchestration start | 1 | 2625.178 | 2625.178 | 2625.178 | undefined | omitted | omitted | 2625.178 |
| replication_interrupted / action completion | 1 | 2370.339 | 2370.339 | 2370.339 | undefined | omitted | omitted | 2370.339 |
| target_reachable_after_repair / orchestration start | 1 | 6445.564 | 6445.564 | 6445.564 | undefined | omitted | omitted | 6445.564 |
| target_reachable_after_repair / action completion | 1 | 6190.725 | 6190.725 | 6190.725 | undefined | omitted | omitted | 6190.725 |
| membership_rejoined / orchestration start | 1 | 6445.564 | 6445.564 | 6445.564 | undefined | omitted | omitted | 6445.564 |
| membership_rejoined / action completion | 1 | 6190.725 | 6190.725 | 6190.725 | undefined | omitted | omitted | 6190.725 |
| replica_link_up / orchestration start | 1 | 10700.484 | 10700.484 | 10700.484 | undefined | omitted | omitted | 10700.484 |
| replica_link_up / action completion | 1 | 10445.645 | 10445.645 | 10445.645 | undefined | omitted | omitted | 10445.645 |
| replica_sync_complete / orchestration start | 1 | 10700.484 | 10700.484 | 10700.484 | undefined | omitted | omitted | 10700.484 |
| replica_sync_complete / action completion | 1 | 10445.645 | 10445.645 | 10445.645 | undefined | omitted | omitted | 10445.645 |
| stable_redundancy / orchestration start | 1 | 11972.450 | 11972.450 | 11972.450 | undefined | omitted | omitted | 11972.450 |
| stable_redundancy / action completion | 1 | 11717.611 | 11717.611 | 11717.611 | undefined | omitted | omitted | 11717.611 |
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
| baseline | 1 | 804 | 804 | 0 | 100.00% | 199.929 | 199.929 |
| fault | 1 | 577 | 443 | 134 | 76.78% | 118.882 | 91.273 |
| post_recovery | 1 | 400 | 400 | 0 | 100.00% | 199.990 | 199.990 |
| reconciliation | 1 | 1114 | 975 | 139 | 87.52% | 200.121 | 175.151 |
| repair | 1 | 317 | 262 | 55 | 82.65% | 199.450 | 164.845 |

Whole-workload schedule counters measured in 1 trials; legacy or incomplete counters unmeasured in 0. Offered: 3606; started: 3212; completed: 3212; dropped: 394. Started fraction of offered: 89.07%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| baseline / completed ops/s | 1 | 199.929 | 199.929 | 199.929 | undefined | omitted | omitted | 199.929 |
| baseline / successful request p50 ms | 1 | 1.252 | 1.252 | 1.252 | undefined | omitted | omitted | 1.252 |
| baseline / successful request p95 ms | 1 | 2.732 | 2.732 | 2.732 | undefined | omitted | omitted | 2.732 |
| baseline / successful request p99 ms | 1 | 3.163 | 3.163 | 3.163 | undefined | omitted | omitted | 3.163 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 1 | 0.768 | 0.768 | 0.768 | undefined | omitted | omitted | 0.768 |
| fault / completed ops/s | 1 | 118.882 | 118.882 | 118.882 | undefined | omitted | omitted | 118.882 |
| fault / successful request p50 ms | 1 | 0.664 | 0.664 | 0.664 | undefined | omitted | omitted | 0.664 |
| fault / successful request p95 ms | 1 | 3.845 | 3.845 | 3.845 | undefined | omitted | omitted | 3.845 |
| fault / successful request p99 ms | 1 | 200.078 | 200.078 | 200.078 | undefined | omitted | omitted | 200.078 |
| fault / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| fault / failed request p95 ms | 1 | 501.145 | 501.145 | 501.145 | undefined | omitted | omitted | 501.145 |
| fault / failed request p99 ms | 1 | 501.145 | 501.145 | 501.145 | undefined | omitted | omitted | 501.145 |
| post_recovery / trial success fraction | 1 | 1.000 | 1.000 | 1.000 | undefined | omitted | omitted | 1.000 |
| post_recovery / completed ops/s | 1 | 199.990 | 199.990 | 199.990 | undefined | omitted | omitted | 199.990 |
| post_recovery / successful request p50 ms | 1 | 0.520 | 0.520 | 0.520 | undefined | omitted | omitted | 0.520 |
| post_recovery / successful request p95 ms | 1 | 2.360 | 2.360 | 2.360 | undefined | omitted | omitted | 2.360 |
| post_recovery / successful request p99 ms | 1 | 2.869 | 2.869 | 2.869 | undefined | omitted | omitted | 2.869 |
| post_recovery / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| post_recovery / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| reconciliation / trial success fraction | 1 | 0.875 | 0.875 | 0.875 | undefined | omitted | omitted | 0.875 |
| reconciliation / completed ops/s | 1 | 200.121 | 200.121 | 200.121 | undefined | omitted | omitted | 200.121 |
| reconciliation / successful request p50 ms | 1 | 0.520 | 0.520 | 0.520 | undefined | omitted | omitted | 0.520 |
| reconciliation / successful request p95 ms | 1 | 2.732 | 2.732 | 2.732 | undefined | omitted | omitted | 2.732 |
| reconciliation / successful request p99 ms | 1 | 4.451 | 4.451 | 4.451 | undefined | omitted | omitted | 4.451 |
| reconciliation / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| reconciliation / failed request p95 ms | 1 | 128.972 | 128.972 | 128.972 | undefined | omitted | omitted | 128.972 |
| reconciliation / failed request p99 ms | 1 | 146.084 | 146.084 | 146.084 | undefined | omitted | omitted | 146.084 |
| repair / trial success fraction | 1 | 0.826 | 0.826 | 0.826 | undefined | omitted | omitted | 0.826 |
| repair / completed ops/s | 1 | 199.450 | 199.450 | 199.450 | undefined | omitted | omitted | 199.450 |
| repair / successful request p50 ms | 1 | 0.546 | 0.546 | 0.546 | undefined | omitted | omitted | 0.546 |
| repair / successful request p95 ms | 1 | 2.602 | 2.602 | 2.602 | undefined | omitted | omitted | 2.602 |
| repair / successful request p99 ms | 1 | 3.321 | 3.321 | 3.321 | undefined | omitted | omitted | 3.321 |
| repair / failed request p50 ms | 1 | 111.411 | 111.411 | 111.411 | undefined | omitted | omitted | 111.411 |
| repair / failed request p95 ms | 1 | 128.972 | 128.972 | 128.972 | undefined | omitted | omitted | 128.972 |
| repair / failed request p99 ms | 1 | 135.104 | 135.104 | 135.104 | undefined | omitted | omitted | 135.104 |

No pooled request latency quantiles are computed: individual report schemas 1, 2 and 3 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

## Write acknowledgement policy

Explicit policy counters measured in 1 trials, unmeasured in 0. Every successful SET remains in the main validation ledger when the later acknowledgement command is short, errors, or is not attempted. Diagnostic SET/WAIT/WAITAOF stage latency remains in individual reports and does not increase logical request or pacing counts. Errors and unattempted commands supply unknown acknowledgement counts, never fabricated zero replies.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 417 | 0 | 417 | 417 | 0 | 0 | 0 |
| fault | 229 | 42 | 229 | 174 | 55 | 0 | 0 |
| post_recovery | 187 | 0 | 187 | 187 | 0 | 0 | 0 |
| reconciliation | 585 | 0 | 585 | 446 | 139 | 0 | 0 |
| repair | 165 | 0 | 165 | 110 | 55 | 0 | 0 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. Attempted WAIT client-call timeouts are counted separately above.

| Observed WAIT replica count | Replies |
|---:|---:|
| 0 | 249 |
| 1 | 1334 |

WAIT-satisfied subset validation: expected 1334; checked 1334; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Validation attempted in 1 trials. Complete means every subset record was checked; preservation also requires zero missing, corrupt and unreadable records.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 12 |
| fault | network | 30 |
| fault | timeout | 37 |
| fault | wait_short | 55 |
| reconciliation | wait_short | 139 |
| repair | wait_short | 55 |

## Acknowledged-write accounting

Validation attempted in 1 trials and complete in 1. SET attempts: 1625; acknowledged writes: 1583; ambiguous SET outcomes: 42.

Checked: 1583; missing: 0; corrupt: 0; unreadable: 0; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 1 | 1583.000 | 1583.000 | 1583.000 | undefined | omitted | omitted | 1583.000 |
| ambiguous writes | 1 | 42.000 | 42.000 | 42.000 | undefined | omitted | omitted | 42.000 |
| missing writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 1 | 0.000 | 0.000 | 0.000 | undefined | omitted | omitted | 0.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | a41f7d40e0b182eec96b39bb | recovered | true | true | [JSON](runs/a41f7d40e0b182eec96b39bb/report.json), [Markdown](runs/a41f7d40e0b182eec96b39bb/report.md) |

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
