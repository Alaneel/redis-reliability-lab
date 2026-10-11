# Redis Reliability Lab — primary-failure-repeat

Run: `c71517bdf2d2fd24649843c9`  
Outcome: **recovered**  
Started: 2026-10-11T15:15:35.336563+08:00  
Finished: 2026-10-11T15:15:55.888905+08:00

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

Fault action completed 211.818 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1777.60 | 1565.78 |
| Original replica owns failed primary slots | 2559.74 | 2347.92 |
| Full usable slot coverage | 2559.74 | 2347.92 |
| Surviving views converge | 2559.74 | 2347.92 |
| First all-shard SET/GET probe succeeds | 3366.94 | 3155.12 |
| Client probes sustain recovery | 4467.86 | 4256.04 |
| Topology and clients sustain recovery | 4467.86 | 4256.04 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 404 | 100.00% | 0 | 199.86 | 0.520 / 0.697 / 0.847 |
| fault | 429 | 90.91% | 39 | 96.02 | 0.472 / 1.030 / 396.141 |
| post_recovery | 200 | 100.00% | 0 | 199.90 | 0.472 / 0.664 / 1.081 |

Acknowledged writes: 506. Ambiguous SET outcomes: 21.

Global pacing: offered 1497; started 1033; completed 1033; dropped 464. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual workload SET/GET client calls; completed requests count their terminal metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and service probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 506; checked: 506; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | timeout | 39 |

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Measurement methodology

- Fault is SIGKILL of the deterministic primary owning the lowest initial slot; the process remains stopped. Initial topology is three primaries and one synchronized replica each.
- Time origin is the start of validated fault injection, including Docker ownership inspections before SIGKILL; successful command completion is also recorded. Redis event times are bounded by polling and request deadlines; views within a snapshot are concurrent and not an atomic Redis snapshot.
- A sample collects direct CLUSTER INFO/NODES and replication/server INFO from all six managed nodes, then parallel SET/GET probes to all three initial shards. Poll interval is a delay after sample completion; actual timestamps define resolution.
- Recovery requires five surviving nodes healthy with exactly one owner per 16384 slots, equal ownership maps, the original replica owning the failed primary slots, unaffected ownership preserved, and all-shard SET/GET probes succeeding for stable_for. Post-recovery probes continue to detect regressions.
- Probes are separate from workload metrics and do not guarantee every request succeeded between samples. The workload success rate, latency distributions and error counts report client-visible operations, including terminal errors after cluster redirections.
- Workload stops and waits for in-flight requests before verifying every acknowledged immutable write. Missing/corrupt reads during traffic also prevent a recovered outcome. Error-returning writes have unknown outcomes, not proven lost writes.
- At most 1200 topology samples are retained (first 1199 and latest); dropped_samples records omitted intermediates. Milestones are retained independently. No process restart or redundancy restoration is attempted.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `8e959199c0f2dff505399e403a8b7f25bfe40c39ec4dbc9ab905198f7d8d36c3`.
