# Redis Reliability Lab — wait-policy-primary-crash

Run: `470c45df26b8848ae11e929a`  
Outcome: **partial**  
Started: 2026-10-11T15:38:33.879145+08:00  
Finished: 2026-10-11T15:38:58.438879+08:00

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

Fault action completed 205.142 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1886.14 | 1681.00 |
| Original replica owns failed primary slots | 2737.88 | 2532.74 |
| Full usable slot coverage | 2737.88 | 2532.74 |
| Surviving views converge | 2737.88 | 2532.74 |
| First all-shard SET/GET probe succeeds | 3012.00 | 2806.85 |
| Client probes sustain recovery | 4346.14 | 4141.00 |
| Topology and clients sustain recovery | 4346.14 | 4141.00 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 821 | 100.00% | 0 | 199.93 | 0.847 / 1.192 / 1.677 |
| fault | 582 | 77.66% | 130 | 130.66 | 0.664 / 1.521 / 268.124 |
| post_recovery | 600 | 82.83% | 103 | 191.57 | 0.495 / 1.030 / 1.314 |

Acknowledged writes: 959. Ambiguous SET outcomes: 48.

Write policy: `wait`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 426 | 0 | 426 | 426 | 0 | 0 | 0 |
| fault | 225 | 48 | 225 | 179 | 46 | 0 | 0 |
| post_recovery | 308 | 0 | 308 | 205 | 103 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 426 | 1.030 / 1.314 / 1.761 |
| baseline | set | 426 | 0.495 / 0.664 / 1.252 |
| baseline | wait | 426 | 0.520 / 0.632 / 0.890 |
| fault | Logical write completion | 273 | 1.030 / 116.982 / 128.972 |
| fault | set | 273 | 0.472 / 0.934 / 1.449 |
| fault | wait | 225 | 0.520 / 116.982 / 128.972 |
| post_recovery | Logical write completion | 308 | 0.981 / 122.831 / 142.192 |
| post_recovery | set | 308 | 0.449 / 0.664 / 1.081 |
| post_recovery | wait | 308 | 0.472 / 122.831 / 135.421 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 426 | 103.739 |
| fault | 179 | 40.185 |
| post_recovery | 205 | 65.455 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count issued WAIT commands only.

Replica-acknowledgement policy recovered for the stable interval: false. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 2312; started 2003; completed 2003; dropped 309. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual workload SET/GET client calls; completed requests count their terminal metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and service probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 959; checked: 959; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAIT-satisfied subset: expected 810; checked 810; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAIT.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 7 |
| fault | network | 41 |
| fault | timeout | 36 |
| fault | wait_short | 46 |
| post_recovery | wait_short | 103 |

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `632cea087a0ebdf23596008773cec25cc59ce13fa4f69380d6bfc1a1389cfe16`.
