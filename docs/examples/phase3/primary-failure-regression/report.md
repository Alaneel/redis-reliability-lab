# Redis Reliability Lab — primary-failure

Run: `5a38321b7d2a351e9600b705`  
Outcome: **recovered**  
Started: 2026-10-11T17:31:45.972347+08:00  
Finished: 2026-10-11T17:32:23.366263+08:00

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

Fault action completed 227.526 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 4341.98 | 4114.45 |
| Original replica owns failed primary slots | 5370.30 | 5142.77 |
| Full usable slot coverage | 5370.30 | 5142.77 |
| Surviving views converge | 5370.30 | 5142.77 |
| First all-shard SET/GET probe succeeds | 5390.86 | 5163.34 |
| Client probes sustain recovery | 8567.05 | 8339.53 |
| Topology and clients sustain recovery | 8567.05 | 8339.53 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 2001 | 100.00% | 0 | 199.70 | 0.546 / 0.768 / 1.761 |
| fault | 1169 | 86.74% | 155 | 136.45 | 0.546 / 1.030 / 2.141 |
| post_recovery | 1000 | 100.00% | 0 | 199.99 | 0.546 / 0.732 / 1.081 |

Acknowledged writes: 2030. Ambiguous SET outcomes: 87.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 1002 | 0 |
| fault | 517 | 87 |
| post_recovery | 511 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 1002 | 0.573 / 0.807 / 1.677 |
| baseline | set | 1002 | 0.573 / 0.807 / 1.677 |
| fault | Logical write completion | 604 | 0.546 / 0.934 / 1.380 |
| fault | set | 604 | 0.546 / 0.934 / 1.380 |
| post_recovery | Logical write completion | 511 | 0.546 / 0.732 / 1.081 |
| post_recovery | set | 511 | 0.546 / 0.732 / 1.030 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 1002 | 100.001 |
| fault | 517 | 60.347 |
| post_recovery | 511 | 102.193 |

Global pacing: offered 4717; started 4170; completed 4170; dropped 547. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 2030; checked: 2030; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 9 |
| fault | network | 78 |
| fault | timeout | 68 |

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
