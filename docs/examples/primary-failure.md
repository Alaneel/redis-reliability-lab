# Redis Reliability Lab — primary-failure

Run: `56dfe45e38604c041b1c6087`  
Outcome: **recovered**  
Started: 2026-10-11T06:52:12.08323Z  
Finished: 2026-10-11T06:52:50.832092Z

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

Times are milliseconds from starting validated fault injection, including Docker ownership checks before SIGKILL. Milestones are first sampled observations, not Redis event timestamps. Unobserved milestones remain absent.

| Observation | Milliseconds |
|---|---:|
| Failure flagged in gossip | 4333.34 |
| Original replica owns failed primary slots | 6399.66 |
| Full usable slot coverage | 6399.66 |
| Surviving views converge | 6399.66 |
| First all-shard SET/GET probe succeeds | 6425.58 |
| Client probes sustain recovery | 9577.20 |
| Topology and clients sustain recovery | 9577.20 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 2004 | 100.00% | 0 | 199.99 | 0.495 / 0.732 / 0.847 |
| fault | 947 | 90.39% | 91 | 98.88 | 0.573 / 1.598 / 172.835 |
| post_recovery | 999 | 100.00% | 0 | 199.79 | 0.428 / 0.664 / 1.677 |

Acknowledged writes: 1960. Ambiguous SET outcomes: 44.

Validation attempted: true; expected: 1960; checked: 1960; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | timeout | 91 |

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Measurement methodology

- Fault is SIGKILL of the deterministic primary owning the lowest initial slot; the process remains stopped. Initial topology is three primaries and one synchronized replica each.
- Time origin is the start of validated fault injection, including Docker ownership inspections before SIGKILL; successful command completion is also recorded. Redis event times are bounded by polling and request deadlines; views within a snapshot are concurrent and not an atomic Redis snapshot.
- A sample collects direct CLUSTER INFO/NODES and replication/server INFO from all six managed nodes, then parallel SET/GET probes to all three initial shards. Poll interval is a delay after sample completion; actual timestamps define resolution.
- Recovery requires five surviving nodes healthy with exactly one owner per 16384 slots, equal ownership maps, the original replica owning the failed primary slots, unaffected ownership preserved, and all-shard SET/GET probes succeeding for stable_for. Post-recovery probes continue to detect regressions.
- Probes are separate from workload metrics and do not guarantee every request succeeded between samples. The workload success rate, latency distributions and error counts report client-visible operations, including terminal errors after cluster redirections.
- Workload stops and waits for in-flight requests before verifying every acknowledged immutable write. Missing/corrupt reads during traffic also prevent a recovered outcome. Error-returning writes have unknown outcomes, not proven lost writes.
- At most 1200 topology samples are retained (first 1199 and latest); dropped_samples records omitted intermediates. Milestones are retained independently. No process restart or redundancy restoration is attempted.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `a4b32016c81316b3a003e553a8a7082b64a17c77059d113d447dd4ebe0409670`.
