# Redis Reliability Lab — dataset-trial

Run: `a272b5d9e6773d9c5fa6c493`  
Outcome: **recovered**  
Started: 2026-10-11T19:55:35.680393+08:00  
Finished: 2026-10-11T19:55:57.60083+08:00

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

Fault action completed 223.769 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 2112.01 | 1888.24 |
| Original replica owns failed primary slots | 3049.83 | 2826.06 |
| Full usable slot coverage | 3049.83 | 2826.06 |
| Surviving views converge | 3049.83 | 2826.06 |
| First all-shard SET/GET probe succeeds | 3176.03 | 2952.26 |
| Client probes sustain recovery | 4511.00 | 4287.23 |
| Topology and clients sustain recovery | 4511.00 | 4287.23 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 356 | 100.00% | 0 | 169.74 | 0.546 / 0.847 / 2.360 |
| fault | 520 | 84.62% | 80 | 115.27 | 0.573 / 0.807 / 325.906 |
| post_recovery | 178 | 100.00% | 0 | 169.89 | 0.495 / 1.030 / 2.039 |

Acknowledged writes: 504. Ambiguous SET outcomes: 47.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 173 | 0 |
| fault | 233 | 47 |
| post_recovery | 98 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 173 | 0.573 / 0.890 / 3.487 |
| baseline | set | 173 | 0.546 / 0.890 / 3.487 |
| fault | Logical write completion | 280 | 0.573 / 0.807 / 1.314 |
| fault | set | 280 | 0.546 / 0.768 / 1.252 |
| post_recovery | Logical write completion | 98 | 0.495 / 1.314 / 1.994 |
| post_recovery | set | 98 | 0.495 / 1.252 / 1.971 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 173 | 82.484 |
| fault | 233 | 51.651 |
| post_recovery | 98 | 93.536 |

Global pacing: offered 1301; started 1054; completed 1054; dropped 247. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 504; checked: 504; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 10 |
| fault | network | 37 |
| fault | timeout | 33 |

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Measurement methodology

- Fault is SIGKILL of the deterministic primary owning the lowest initial slot; the process remains stopped. Initial topology is three primaries and one synchronized replica each.
- Time origin is the start of validated fault injection, including Docker ownership inspections before SIGKILL; successful command completion is also recorded. Redis event times are bounded by polling and request deadlines; views within a snapshot are concurrent and not an atomic Redis snapshot.
- A sample collects direct CLUSTER INFO/NODES and replication/server INFO from all six managed nodes, then parallel SET/GET probes to all three initial shards. Poll interval is a delay after sample completion; actual timestamps define resolution.
- Recovery requires five surviving nodes healthy with exactly one owner per 16384 slots, equal ownership maps, the original replica owning the failed primary slots, unaffected ownership preserved, and all-shard SET/GET probes succeeding for stable_for. Post-recovery probes continue to detect regressions.
- Probes are separate from workload metrics and do not guarantee every request succeeded between samples. The workload success rate, latency distributions and error counts report client-visible operations, including terminal errors after cluster redirections.
- Workload stops and waits for in-flight requests before verifying every acknowledged immutable write. Missing/corrupt reads during traffic also prevent a recovered outcome. Error-returning writes have unknown outcomes, not proven lost writes.
- At most 1200 topology samples are retained (first 1199 and latest); dropped_samples records omitted intermediates. Milestones are retained independently. No process restart or redundancy restoration is attempted.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `db6e82bfd40c14f7dfc88987e54c6bac89af794f0e24547210fdd29ce10fe2c1`.
