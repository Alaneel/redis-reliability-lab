# Redis Reliability Lab — dataset-trial

Run: `f5eca49fc4fc2128a29cda93`  
Outcome: **recovered**  
Started: 2026-10-11T19:48:39.766482+08:00  
Finished: 2026-10-11T19:49:17.615149+08:00

## Recovery

| Condition | Observed |
|---|---|
| Surviving topology converged | true |
| All-shard client recovery sustained | true |
| Topology and client recovery jointly sustained | true |
| All acknowledged writes verified | true |
| Target process restarted | false |
| Original replication redundancy restored | true |
| Recovery deadline expired | false |
| Cleanup verified | true |

The primary remains running during whole-node bidirectional network isolation. Connectivity restoration, topology reconciliation and replication redundancy are evaluated separately. Redis replication is asynchronous; acknowledged SETs are not durability guarantees.

## Network restoration

| Condition | Observed |
|---|---|
| Original endpoints and process restored | true |
| Topology and clients jointly reconciled | true |
| Reconciliation deadline expired | false |

Restoration action duration: 1293.868 ms.
Stable reconciliation first observed 5651.855 ms after restoration completed.

| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |
|---|---|---|---|---:|---|---|
| 2026-10-11T19:48:56.339915+08:00 | true | true | false | 32297 | 3cc3a26eafb1e60e08a680a7f5069e6c847f2b23 |  |
| 2026-10-11T19:48:59.140703+08:00 | true | true | false | 32297 | 3cc3a26eafb1e60e08a680a7f5069e6c847f2b23 |  |
| 2026-10-11T19:49:01.920789+08:00 | true | true | false | 32297 | 3cc3a26eafb1e60e08a680a7f5069e6c847f2b23 |  |
| 2026-10-11T19:49:04.670556+08:00 | true | true | false | 32297 | 3cc3a26eafb1e60e08a680a7f5069e6c847f2b23 |  |
| 2026-10-11T19:49:14.510756+08:00 | true | true | true | 32297 | 3cc3a26eafb1e60e08a680a7f5069e6c847f2b23 |  |

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 1436.895 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1970.74 | 533.85 |
| Original replica owns failed primary slots | 2911.90 | 1475.00 |
| Full usable slot coverage | 2911.90 | 1475.00 |
| Surviving views converge | 2911.90 | 1475.00 |
| First all-shard SET/GET probe succeeds | 3039.58 | 1602.69 |
| Client probes sustain recovery | 4897.00 | 3460.11 |
| Topology and clients sustain recovery | 4897.00 | 3460.11 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 357 | 100.00% | 0 | 169.83 | 0.573 / 0.768 / 1.677 |
| fault | 2174 | 97.19% | 61 | 157.27 | 0.546 / 0.847 / 1.380 |
| reconciliation | 1072 | 100.00% | 0 | 169.92 | 0.495 / 0.664 / 0.847 |
| post_recovery | 170 | 100.00% | 0 | 169.88 | 0.520 / 1.252 / 2.039 |

Acknowledged writes: 1849. Ambiguous SET outcomes: 32.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 176 | 0 |
| fault | 1053 | 32 |
| reconciliation | 546 | 0 |
| post_recovery | 74 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 176 | 0.573 / 0.768 / 1.252 |
| baseline | set | 176 | 0.546 / 0.732 / 1.252 |
| fault | Logical write completion | 1085 | 0.573 / 0.890 / 1.314 |
| fault | set | 1085 | 0.546 / 0.807 / 1.252 |
| reconciliation | Logical write completion | 546 | 0.495 / 0.664 / 0.847 |
| reconciliation | set | 546 | 0.495 / 0.632 / 0.807 |
| post_recovery | Logical write completion | 74 | 0.546 / 0.890 / 1.973 |
| post_recovery | set | 74 | 0.546 / 0.847 / 1.905 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 176 | 83.726 |
| fault | 1053 | 76.173 |
| reconciliation | 546 | 86.546 |
| post_recovery | 74 | 73.947 |

Global pacing: offered 3949; started 3773; completed 3773; dropped 176. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1849; checked: 1849; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 2 |
| fault | network | 30 |
| fault | timeout | 29 |

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Measurement methodology

- Fault is a complete bidirectional Docker bridge disconnection of the deterministic primary owning the lowest initial slot. Redis remains running. All peer and published client connectivity is removed; this is not a selective or one-way partition.
- The orchestration and action-completion anchors are recorded separately using monotonic runtime durations. isolation_duration is a minimum hold from successful action completion, bounded by the recovery deadline; samples and bounded process probes can extend the hold. Actual restoration timestamps record the observed duration.
- During isolation, recovery requires the five surviving nodes to converge on healthy usable slot ownership, the original replica to own the selected primary slots, unchanged unaffected ownership, and successful all-shard SET/GET probes over stable_for. A stable milestone is retained while observation continues until restoration.
- Docker exec runs fixed Redis liveness and topology commands independent of the disconnected data plane. PONG and unchanged container PID, start time and Redis run ID prove sampled process survival, not cluster health. Isolation success requires at least one such disconnected observation.
- Restoration reconnects only the same owned node to the same owned network with its original private IP and loopback port. Process identity and endpoint identity must remain unchanged. Cancellation uses a fresh bounded restoration attempt before independent cleanup; safe cleanup can directly remove an owned detached node.
- After restoration, fresh stability tracking requires all six original identities and endpoints, three healthy primaries with synchronized replicas, the promoted replica retaining shard ownership, former-primary demotion, unchanged unaffected shards, and all-shard client probes. Service recovery, connectivity restoration and replication redundancy have separate evidence.
- Poll intervals are delays after sample completion; concurrent views are not atomic snapshots. Probes do not prove continuous availability. Workload metrics include terminal client errors separately from probe results. At most 1200 topology samples and 64 process samples are retained.
- Traffic drains before validating every successfully acknowledged immutable SET. Missing, corrupt, unreadable or unverified records prevent full recovery. Error-returning SET outcomes remain ambiguous. No client can write to the disconnected minority, so this experiment does not exercise minority-side write conflicts or establish loss-free asynchronous replication.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `eb2a0312e2381d0d006ad9eb67f8d01f66d522f628c825b2af991755bf22e5cb`.
