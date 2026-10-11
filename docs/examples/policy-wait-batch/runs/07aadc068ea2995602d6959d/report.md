# Redis Reliability Lab — policy-wait-primary-crash

Run: `07aadc068ea2995602d6959d`  
Outcome: **partial**  
Started: 2026-10-11T15:43:17.559611+08:00  
Finished: 2026-10-11T15:43:46.689815+08:00

## Recovery

| Condition | Observed |
|---|---|
| Surviving topology converged | true |
| All-shard client recovery sustained | true |
| Topology and client recovery jointly sustained | true |
| All acknowledged writes verified | true |
| Killed process restarted | false |
| Original replication redundancy restored | false |
| Recovery deadline expired | false |
| Cleanup verified | true |

The primary stays stopped. Service recovery can succeed with reduced redundancy. Redis replication is asynchronous; an acknowledged SET is not a durability guarantee. Missing acknowledged writes cause a partial outcome.

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 210.659 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1773.32 | 1562.66 |
| Original replica owns failed primary slots | 2555.98 | 2345.32 |
| Full usable slot coverage | 2555.98 | 2345.32 |
| Surviving views converge | 2555.98 | 2345.32 |
| First all-shard SET/GET probe succeeds | 3471.67 | 3261.02 |
| Client probes sustain recovery | 4620.01 | 4409.35 |
| Topology and clients sustain recovery | 4620.01 | 4409.35 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 1603 | 100.00% | 0 | 199.97 | 0.768 / 1.192 / 1.521 |
| fault | 542 | 74.54% | 138 | 114.36 | 0.697 / 1.677 / 359.312 |
| post_recovery | 600 | 82.67% | 104 | 190.77 | 0.546 / 1.081 / 1.677 |

Acknowledged writes: 1355. Ambiguous SET outcomes: 42.

Write policy: `wait`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 809 | 0 | 809 | 809 | 0 | 0 | 0 |
| fault | 233 | 42 | 233 | 175 | 58 | 0 | 0 |
| post_recovery | 313 | 0 | 313 | 209 | 104 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 809 | 0.981 / 1.314 / 1.849 |
| baseline | set | 809 | 0.495 / 0.664 / 0.981 |
| baseline | wait | 809 | 0.495 / 0.664 / 0.847 |
| fault | Logical write completion | 275 | 1.081 / 116.982 / 128.972 |
| fault | set | 275 | 0.520 / 1.192 / 2.248 |
| fault | wait | 233 | 0.546 / 116.982 / 128.972 |
| post_recovery | Logical write completion | 313 | 1.030 / 122.831 / 142.192 |
| post_recovery | set | 313 | 0.472 / 0.602 / 0.697 |
| post_recovery | wait | 313 | 0.495 / 122.831 / 142.192 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 809 | 100.918 |
| fault | 175 | 36.924 |
| post_recovery | 209 | 66.452 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Replica-acknowledgement policy recovered for the stable interval: false. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 3151; started 2745; completed 2745; dropped 406. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1355; checked: 1355; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAIT-satisfied subset: expected 1193; checked 1193; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAIT.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 13 |
| fault | network | 29 |
| fault | timeout | 38 |
| fault | wait_short | 58 |
| post_recovery | wait_short | 104 |

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Errors

- experiment outcome: partial

## Measurement methodology

- Fault is SIGKILL of the deterministic primary owning the lowest initial slot; the process remains stopped. Initial topology is three primaries and one synchronized replica each.
- Time origin is the start of validated fault injection, including Docker ownership inspections before SIGKILL; successful command completion is also recorded. Redis event times are bounded by polling and request deadlines; views within a snapshot are concurrent and not an atomic Redis snapshot.
- A sample collects direct CLUSTER INFO/NODES and replication/server INFO from all six managed nodes, then parallel SET/GET probes to all three initial shards. Poll interval is a delay after sample completion; actual timestamps define resolution.
- Recovery requires five surviving nodes healthy with exactly one owner per 16384 slots, equal ownership maps, the original replica owning the failed primary slots, unaffected ownership preserved, and all-shard SET/GET probes succeeding for stable_for. Post-recovery probes continue to detect regressions.
- Probes are separate from workload metrics and do not guarantee every request succeeded between samples. The workload success rate, latency distributions and error counts report client-visible operations, including terminal errors after cluster redirections.
- Workload stops and waits for in-flight requests before verifying every acknowledged immutable write. Missing/corrupt reads during traffic also prevent a recovered outcome. Error-returning writes have unknown outcomes, not proven lost writes.
- At most 1200 topology samples are retained (first 1199 and latest); dropped_samples records omitted intermediates. Milestones are retained independently. No process restart or redundancy restoration is attempted.
- SET+WAIT uses one physical connection, records every successful SET even when WAIT is incomplete, and separately tracks all-shard replica-acknowledgement policy probes. Basic service recovery can succeed while that policy remains unavailable. WAIT is not a durability or strong-consistency guarantee.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `e2a6d0c505f6d1bac0abfa27b1e134e13352da9e819b7560d276fe8d7080b863`.
