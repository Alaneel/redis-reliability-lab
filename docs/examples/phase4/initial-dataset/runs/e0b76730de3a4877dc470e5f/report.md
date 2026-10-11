# Redis Reliability Lab — dataset-trial

Run: `e0b76730de3a4877dc470e5f`  
Outcome: **recovered**  
Started: 2026-10-11T19:43:52.333153+08:00  
Finished: 2026-10-11T19:44:26.929186+08:00

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

Restoration action duration: 1371.340 ms.
Stable reconciliation first observed 6188.076 ms after restoration completed.

| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |
|---|---|---|---|---:|---|---|
| 2026-10-11T19:44:07.108335+08:00 | true | true | false | 7195 | 456193e39fd53f1485075530dce359d82ce8518d |  |
| 2026-10-11T19:44:09.995014+08:00 | true | true | false | 7195 | 456193e39fd53f1485075530dce359d82ce8518d |  |
| 2026-10-11T19:44:12.556199+08:00 | true | true | false | 7195 | 456193e39fd53f1485075530dce359d82ce8518d |  |
| 2026-10-11T19:44:15.083499+08:00 | true | true | false | 7195 | 456193e39fd53f1485075530dce359d82ce8518d |  |
| 2026-10-11T19:44:23.997245+08:00 | true | true | true | 7195 | 456193e39fd53f1485075530dce359d82ce8518d |  |

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 1396.712 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1940.81 | 544.10 |
| Original replica owns failed primary slots | 3601.56 | 2204.84 |
| Full usable slot coverage | 3601.56 | 2204.84 |
| Surviving views converge | 3601.56 | 2204.84 |
| First all-shard SET/GET probe succeeds | 4100.39 | 2703.68 |
| Client probes sustain recovery | 5290.60 | 3893.89 |
| Topology and clients sustain recovery | 5290.60 | 3893.89 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 209 | 100.00% | 0 | 99.46 | 0.697 / 0.981 / 3.321 |
| fault | 1130 | 91.86% | 92 | 94.56 | 0.632 / 1.081 / 2.360 |
| reconciliation | 687 | 100.00% | 0 | 99.94 | 0.546 / 0.807 / 1.135 |
| post_recovery | 100 | 100.00% | 0 | 99.88 | 0.602 / 0.768 / 0.934 |

Acknowledged writes: 1020. Ambiguous SET outcomes: 56.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 107 | 0 |
| fault | 498 | 56 |
| reconciliation | 359 | 0 |
| post_recovery | 56 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 107 | 0.732 / 0.981 / 3.321 |
| baseline | set | 107 | 0.697 / 0.934 / 3.321 |
| fault | Logical write completion | 554 | 0.664 / 1.030 / 1.449 |
| fault | set | 554 | 0.632 / 0.981 / 1.449 |
| reconciliation | Logical write completion | 359 | 0.573 / 0.847 / 1.252 |
| reconciliation | set | 359 | 0.546 / 0.807 / 1.135 |
| post_recovery | Logical write completion | 56 | 0.602 / 0.768 / 0.843 |
| post_recovery | set | 56 | 0.573 / 0.732 / 0.813 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 107 | 50.918 |
| fault | 498 | 41.674 |
| reconciliation | 359 | 52.225 |
| post_recovery | 56 | 55.933 |

Global pacing: offered 2192; started 2126; completed 2126; dropped 66. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1020; checked: 1020; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 14 |
| fault | network | 42 |
| fault | timeout | 36 |

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `562959e5dcf87e2113d36d7c5951566a68df3f72cda490a01ac85d79b9b6e3b7`.
