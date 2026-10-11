# Redis Reliability Lab — dataset-trial

Run: `6fcd217008a29878e09b60d0`  
Outcome: **recovered**  
Started: 2026-10-11T19:56:27.699771+08:00  
Finished: 2026-10-11T19:57:07.74891+08:00

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

Restoration action duration: 1349.931 ms.
Stable reconciliation first observed 6880.241 ms after restoration completed.

| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |
|---|---|---|---|---:|---|---|
| 2026-10-11T19:56:44.380793+08:00 | true | true | false | 74593 | 01d99bb6c16c450ea40665b7d1a0a15effa3d673 |  |
| 2026-10-11T19:56:47.284224+08:00 | true | true | false | 74593 | 01d99bb6c16c450ea40665b7d1a0a15effa3d673 |  |
| 2026-10-11T19:56:50.089239+08:00 | true | true | false | 74593 | 01d99bb6c16c450ea40665b7d1a0a15effa3d673 |  |
| 2026-10-11T19:56:52.857984+08:00 | true | true | false | 74593 | 01d99bb6c16c450ea40665b7d1a0a15effa3d673 |  |
| 2026-10-11T19:57:04.021307+08:00 | true | true | true | 74593 | 01d99bb6c16c450ea40665b7d1a0a15effa3d673 |  |

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 1477.642 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 2964.62 | 1486.97 |
| Original replica owns failed primary slots | 3905.94 | 2428.30 |
| Full usable slot coverage | 3905.94 | 2428.30 |
| Surviving views converge | 3905.94 | 2428.30 |
| First all-shard SET/GET probe succeeds | 4038.30 | 2560.66 |
| Client probes sustain recovery | 5055.57 | 3577.93 |
| Topology and clients sustain recovery | 5055.57 | 3577.93 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 357 | 100.00% | 0 | 169.91 | 0.602 / 0.890 / 2.602 |
| fault | 2066 | 95.93% | 84 | 146.87 | 0.602 / 0.981 / 2.602 |
| reconciliation | 1290 | 100.00% | 0 | 170.03 | 0.546 / 0.890 / 2.732 |
| post_recovery | 170 | 100.00% | 0 | 169.90 | 0.602 / 0.807 / 1.677 |

Acknowledged writes: 1932. Ambiguous SET outcomes: 41.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 206 | 0 |
| fault | 1005 | 41 |
| reconciliation | 631 | 0 |
| post_recovery | 90 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 206 | 0.602 / 0.934 / 2.602 |
| baseline | set | 206 | 0.573 / 0.890 / 2.478 |
| fault | Logical write completion | 1046 | 0.602 / 1.081 / 2.039 |
| fault | set | 1046 | 0.573 / 0.981 / 2.039 |
| reconciliation | Logical write completion | 631 | 0.573 / 0.890 / 3.662 |
| reconciliation | set | 631 | 0.546 / 0.890 / 3.662 |
| post_recovery | Logical write completion | 90 | 0.632 / 0.807 / 1.941 |
| post_recovery | set | 90 | 0.602 / 0.807 / 1.879 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 206 | 98.042 |
| fault | 1005 | 71.445 |
| reconciliation | 631 | 83.171 |
| post_recovery | 90 | 89.947 |

Global pacing: offered 4208; started 3883; completed 3883; dropped 325. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1932; checked: 1932; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 10 |
| fault | network | 31 |
| fault | timeout | 43 |

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `a8e714fdcf8dcb787d0f2aaf1975f5cd0266f961aa59fb6ba7bde445c47b53a5`.
