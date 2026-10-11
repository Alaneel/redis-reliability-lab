# Redis Reliability Lab — primary-failure

Run: `4f98886ad19defb3c16e8744`  
Outcome: **recovered**  
Started: 2026-10-11T18:53:33.87183+08:00  
Finished: 2026-10-11T18:54:10.498915+08:00

## Recovery

| Condition | Observed |
|---|---|
| Surviving topology converged | true |
| All-shard client recovery sustained | true |
| Topology and client recovery jointly sustained | true |
| All acknowledged writes verified | true |
| Target process restarted | false |
| Original replication redundancy restored | false |
| Recovery deadline expired | false |
| Cleanup verified | true |

The primary stays stopped. Service recovery can succeed with reduced redundancy. Redis replication is asynchronous; an acknowledged SET is not a durability guarantee. Missing acknowledged writes cause a partial outcome.

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 205.206 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 3275.62 | 3070.41 |
| Original replica owns failed primary slots | 5333.40 | 5128.19 |
| Full usable slot coverage | 5333.40 | 5128.19 |
| Surviving views converge | 5333.40 | 5128.19 |
| First all-shard SET/GET probe succeeds | 5358.62 | 5153.42 |
| Client probes sustain recovery | 8535.12 | 8329.92 |
| Topology and clients sustain recovery | 8535.12 | 8329.92 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 2003 | 100.00% | 0 | 199.97 | 0.546 / 0.807 / 1.135 |
| fault | 1156 | 86.07% | 161 | 135.44 | 0.520 / 1.030 / 2.602 |
| post_recovery | 1000 | 100.00% | 0 | 199.99 | 0.546 / 0.934 / 1.942 |

Acknowledged writes: 2020. Ambiguous SET outcomes: 91.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 1004 | 0 |
| fault | 508 | 91 |
| post_recovery | 508 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 1004 | 0.573 / 0.807 / 1.252 |
| baseline | set | 1004 | 0.573 / 0.807 / 1.192 |
| fault | Logical write completion | 599 | 0.546 / 1.030 / 1.942 |
| fault | set | 599 | 0.520 / 1.030 / 1.849 |
| post_recovery | Logical write completion | 508 | 0.546 / 0.890 / 1.942 |
| post_recovery | set | 508 | 0.546 / 0.890 / 1.942 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 1004 | 100.234 |
| fault | 508 | 59.518 |
| post_recovery | 508 | 101.596 |

Global pacing: offered 4710; started 4159; completed 4159; dropped 551. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 2020; checked: 2020; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 11 |
| fault | network | 80 |
| fault | timeout | 70 |

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Measurement methodology

- Fault is SIGKILL of the deterministic primary owning the lowest initial slot; the process remains stopped. Initial topology is three primaries and one synchronized replica each.
- Time origin is the start of validated fault injection, including Docker ownership inspections before SIGKILL; successful command completion is also recorded. Redis event times are bounded by polling and request deadlines; views within a snapshot are concurrent and not an atomic Redis snapshot.
- A sample collects direct CLUSTER INFO/NODES and replication/server INFO from all six managed nodes, then parallel SET/GET probes to all three initial shards. Poll interval is a delay after sample completion; actual timestamps define resolution.
- Recovery requires five surviving nodes healthy with exactly one owner per 16384 slots, equal ownership maps, the original replica owning the failed primary slots, unaffected ownership preserved, and all-shard SET/GET probes succeeding for stable_for. Post-recovery probes continue to detect regressions.
- Probes are separate from workload metrics and do not guarantee every request succeeded between samples. The workload success rate, latency distributions and error counts report client-visible operations, including terminal errors after cluster redirections.
- Workload stops and waits for in-flight requests before verifying every acknowledged immutable write. Missing/corrupt reads during traffic also prevent a recovered outcome. Error-returning writes have unknown outcomes, not proven lost writes.
- At most 1200 topology samples are retained (first 1199 and latest); dropped_samples records omitted intermediates. Milestones are retained independently. No process restart or redundancy restoration is attempted.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `95499751a2c5fdb488ba027b42df495f9965b435e446bf61186f6ce4443b7c52`.
