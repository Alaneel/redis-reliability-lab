# Redis Reliability Lab — policy-wait-primary-crash

Run: `2f20da5016af35f23348d0de`  
Outcome: **partial**  
Started: 2026-10-11T15:42:48.549653+08:00  
Finished: 2026-10-11T15:43:17.513247+08:00

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

Fault action completed 201.039 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1762.09 | 1561.06 |
| Original replica owns failed primary slots | 2536.22 | 2335.18 |
| Full usable slot coverage | 2536.22 | 2335.18 |
| Surviving views converge | 2536.22 | 2335.18 |
| First all-shard SET/GET probe succeeds | 3343.15 | 3142.12 |
| Client probes sustain recovery | 4495.40 | 4294.36 |
| Topology and clients sustain recovery | 4495.40 | 4294.36 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 1603 | 100.00% | 0 | 199.90 | 0.847 / 1.380 / 2.732 |
| fault | 529 | 75.61% | 129 | 114.69 | 0.732 / 1.598 / 181.477 |
| post_recovery | 600 | 81.83% | 109 | 194.25 | 0.602 / 1.135 / 1.380 |

Acknowledged writes: 1358. Ambiguous SET outcomes: 37.

Write policy: `wait`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 810 | 0 | 810 | 810 | 0 | 0 | 0 |
| fault | 221 | 37 | 221 | 169 | 52 | 0 | 0 |
| post_recovery | 327 | 0 | 327 | 218 | 109 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 810 | 1.081 / 1.521 / 2.869 |
| baseline | set | 810 | 0.546 / 0.807 / 1.761 |
| baseline | wait | 810 | 0.495 / 0.732 / 1.081 |
| fault | Logical write completion | 258 | 1.030 / 116.982 / 122.831 |
| fault | set | 258 | 0.520 / 1.030 / 2.141 |
| fault | wait | 221 | 0.520 / 116.982 / 122.831 |
| post_recovery | Logical write completion | 327 | 1.030 / 122.831 / 135.421 |
| post_recovery | set | 327 | 0.472 / 0.664 / 0.890 |
| post_recovery | wait | 327 | 0.495 / 122.831 / 135.421 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 810 | 101.009 |
| fault | 169 | 36.639 |
| post_recovery | 218 | 70.578 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Replica-acknowledgement policy recovered for the stable interval: false. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 3126; started 2732; completed 2732; dropped 394. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1358; checked: 1358; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAIT-satisfied subset: expected 1197; checked 1197; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAIT.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 8 |
| fault | network | 29 |
| fault | timeout | 40 |
| fault | wait_short | 52 |
| post_recovery | wait_short | 109 |

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
