# Redis Reliability Lab — repeated trials

Batch: `batch-e3a8aef71917461d7de5322c`  
Outcome: **failed**  
Started: 2026-10-11T07:17:52.022606Z  
Finished: 2026-10-11T07:18:28.201987Z

Requested: 2; started: 2; callbacks completed: 2; unstarted: 0. Cleanup failures: 0; incomplete individual evidence: 0.

| Trial outcome | Count |
|---|---:|
| recovered | 0 |
| partial | 0 |
| failed | 2 |
| aborted | 0 |
| unstarted | 0 |

Recovery success fraction: 0.00% of **all 2 requested trials**. Started-trial recovery fraction: 0.00%.

## Recovery observations

Timing statistics below are conditional on observed milestones, including observed milestones in partial or failed trials. Unobserved events are not zero-duration recoveries. Censoring means unobserved by the accepted sampled horizon; it does not prove an event physically occurred after that horizon. No-requested-fault trials are not at risk. Aborted/error horizons are separate from deadline horizons in JSON.

| Event | Observed | Censored | Deadline censored | Not at risk | Missing | Unstarted |
|---|---:|---:|---:|---:|---:|---:|
| failure_detected | 0 | 2 | 2 | 0 | 0 | 0 |
| replica_promoted | 0 | 2 | 2 | 0 | 0 | 0 |
| full_slot_coverage | 0 | 2 | 2 | 0 | 0 | 0 |
| topology_converged | 0 | 2 | 2 | 0 | 0 | 0 |
| client_first_all_shard_success | 0 | 2 | 2 | 0 | 0 | 0 |
| client_stable | 0 | 2 | 2 | 0 | 0 | 0 |
| combined_stable | 0 | 2 | 2 | 0 | 0 | 0 |

Descriptive timing statistics use milliseconds. **P90/P95 are omitted for N<20**; small samples do not support reliable tail estimates. Sample standard deviation uses N-1 and is undefined for N<2. These are descriptive statistics, not confidence intervals.

| Timing / anchor | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| fault action duration | 2 | 204.361 | 208.765 | 208.765 | 6.227 | omitted | omitted | 213.168 |
| failure_detected / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| failure_detected / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_promoted / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| replica_promoted / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| full_slot_coverage / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| full_slot_coverage / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| topology_converged / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| topology_converged / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| client_first_all_shard_success / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| client_first_all_shard_success / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| client_stable / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| client_stable / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| combined_stable / orchestration start | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| combined_stable / action completion | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |

## Client operations

Completed request fractions weight actual request counts. Throughput divides summed completed/successful counts by summed trial phase wall-clock seconds; it does not imply workers achieved the configured offered rate. Zero-attempt phases have an unobserved success fraction.

| Phase | Trials with phase | Completed | Successful | Failed | Request success fraction | Completed ops/s | Successful ops/s |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 2 | 837 | 837 | 0 | 100.00% | 199.283 | 199.283 |
| fault | 2 | 99 | 83 | 16 | 83.84% | 66.342 | 55.620 |

Whole-workload schedule counters measured in 2 trials; legacy or incomplete counters unmeasured in 0. Offered: 975; started: 936; completed: 936; dropped: 39. Started fraction of offered: 96.00%. These counters exclude service probes and validation.

### Across-trial performance observations

| Phase / metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline / trial success fraction | 2 | 1.000 | 1.000 | 1.000 | 0.000 | omitted | omitted | 1.000 |
| baseline / completed ops/s | 2 | 199.173 | 199.283 | 199.283 | 0.155 | omitted | omitted | 199.393 |
| baseline / successful request p50 ms | 2 | 0.573 | 0.588 | 0.588 | 0.020 | omitted | omitted | 0.602 |
| baseline / successful request p95 ms | 2 | 0.807 | 0.918 | 0.918 | 0.158 | omitted | omitted | 1.030 |
| baseline / successful request p99 ms | 2 | 1.677 | 1.763 | 1.763 | 0.122 | omitted | omitted | 1.849 |
| baseline / failed request p50 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p95 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| baseline / failed request p99 ms | 0 | unobserved | unobserved | unobserved | unobserved | omitted | omitted | unobserved |
| fault / trial success fraction | 2 | 0.826 | 0.838 | 0.838 | 0.016 | omitted | omitted | 0.849 |
| fault / completed ops/s | 2 | 63.257 | 66.266 | 66.266 | 4.255 | omitted | omitted | 69.274 |
| fault / successful request p50 ms | 2 | 0.546 | 0.546 | 0.546 | 0.000 | omitted | omitted | 0.546 |
| fault / successful request p95 ms | 2 | 0.768 | 0.952 | 0.952 | 0.259 | omitted | omitted | 1.135 |
| fault / successful request p99 ms | 2 | 0.940 | 1.885 | 1.885 | 1.337 | omitted | omitted | 2.831 |
| fault / failed request p50 ms | 2 | 501.090 | 501.129 | 501.129 | 0.054 | omitted | omitted | 501.167 |
| fault / failed request p95 ms | 2 | 501.090 | 501.129 | 501.129 | 0.054 | omitted | omitted | 501.167 |
| fault / failed request p99 ms | 2 | 501.090 | 501.129 | 501.129 | 0.054 | omitted | omitted | 501.167 |

No pooled request latency quantiles are computed: individual report schemas 1 and 2 contain histogram quantile summaries, not mergeable bucket counts. Latency statistics describe distributions of per-trial successful/failed p50/p95/p99 values, conditional on trials with those observations; their means are not pooled request quantiles.

### Terminal errors

| Phase | Class | Count |
|---|---|---:|
| fault | timeout | 16 |

## Acknowledged-write accounting

Validation attempted in 2 trials and complete in 2. SET attempts: 481; acknowledged writes: 476; ambiguous SET outcomes: 5.

Checked: 476; missing: 0; corrupt: 0; unreadable: 8; explicitly unverified: 0; acknowledged writes not checked (including skipped validation): 0. During-traffic missing checks: 0; mismatched checks: 0.

Totals and descriptive measurements use verified individual report evidence; failed and partial trials with valid evidence remain included. Incomplete-evidence trials remain in outcomes and are explicitly missing from measurements. Redis asynchronous replication and client ambiguity prevent a universal durability guarantee.

| Trial data metric | N | Min | Median | Mean | Sample SD | P90 | P95 | Max |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| acknowledged writes | 2 | 236.000 | 238.000 | 238.000 | 2.828 | omitted | omitted | 240.000 |
| ambiguous writes | 2 | 2.000 | 2.500 | 2.500 | 0.707 | omitted | omitted | 3.000 |
| missing writes / validation attempted | 2 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| corrupt writes / validation attempted | 2 | 0.000 | 0.000 | 0.000 | 0.000 | omitted | omitted | 0.000 |
| unreadable writes / validation attempted | 2 | 3.000 | 4.000 | 4.000 | 1.414 | omitted | omitted | 5.000 |

## Individual trials

| Trial | Run ID | Outcome | Individual evidence complete | Cleanup verified | Report |
|---:|---|---|---|---|---|
| 1 | 948d28b9474a42d4b514a332 | failed | true | true | [JSON](runs/948d28b9474a42d4b514a332/report.json), [Markdown](runs/948d28b9474a42d4b514a332/report.md) |
| 2 | 184cb060cd5efa1e64511cac | failed | true | true | [JSON](runs/184cb060cd5efa1e64511cac/report.json), [Markdown](runs/184cb060cd5efa1e64511cac/report.md) |

## Execution errors

- trial 1: experiment outcome: failed; trial 2: experiment outcome: failed

## Measurement methodology

- Trials execute sequentially with a fresh environment per callback; the next callback starts only after the previous callback returns with successful cleanup evidence (or an authoritative report that no environment was provisioned).
- Recovery success fraction uses all requested trials as the denominator. Outcomes, started/unstarted counts and individual reports preserve failed, partial and aborted trials; the started-trial fraction is reported separately.
- Recovery timings are conditional on observed milestones, including observed milestones in partial/failed trials. Unobserved post-fault events are censored at the accepted sampled observation horizon with deadline/abort/error reasons, never zero-duration samples; this does not prove the physical event occurred after that horizon. Trials without a requested fault are separately not at risk. No survival distribution, independence assumption or population reliability guarantee is inferred.
- Event durations use runtime SinceFaultMS and fault_action_duration_ms when present, preserving monotonic elapsed-time measurements. Schema 1 action-completion-relative durations fall back to wall-clock timestamp differences and can be affected by clock adjustments. First sampled observations include polling uncertainty; action duration includes ownership checks and fault command overhead.
- Min, median, mean, N-1 sample standard deviation and maximum describe available trial observations. P90/P95 use nearest rank only when N>=20; smaller samples have an explicit tail-estimation caveat.
- Completed-request success fractions weight actual completed counts; aggregate throughput divides summed completed/successful counts by summed phase wall-clock duration. Offered load and worker saturation are not inferred from completed requests.
- Whole-workload offered, started, completed and dropped schedule-slot counters are aggregated only for reports carrying explicit pacing methodology. Legacy absent counters are unmeasured, not zero. Offered=started+dropped and started=completed after workload quiescence; service probes and validation are excluded.
- Acknowledged SET writes remain distinct from ambiguous SET outcomes and final read-back validation. Skipped/incomplete validation contributes acknowledged-not-checked counts; no observed loss is not a durability guarantee.

Effective experiment configuration SHA-256: `3eb0f14865498b0318d98c5fd9edbe4b7292ae905ad1bd167a69dd0d5a947545`. Individual JSON reports retain configuration, host/environment metadata, raw topology observations and their own measurement methodology.
