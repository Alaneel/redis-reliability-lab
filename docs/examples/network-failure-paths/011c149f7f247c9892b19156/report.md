# Redis Reliability Lab — primary-network-isolation

Run: `011c149f7f247c9892b19156`  
Outcome: **partial**  
Started: 2026-10-11T15:37:29.161396+08:00  
Finished: 2026-10-11T15:37:55.568407+08:00

## Recovery

| Condition | Observed |
|---|---|
| Surviving topology converged | false |
| All-shard client recovery sustained | false |
| Topology and client recovery jointly sustained | false |
| All acknowledged writes verified | true |
| Killed process restarted | false |
| Original replication redundancy restored | false |
| Recovery deadline expired | false |
| Cleanup verified | true |

The primary remains running during whole-node bidirectional network isolation. Connectivity restoration, topology reconciliation and replication redundancy are evaluated separately. Redis replication is asynchronous; acknowledged SETs are not durability guarantees.

## Network restoration

| Condition | Observed |
|---|---|
| Original endpoints and process restored | true |
| Topology and clients jointly reconciled | false |
| Reconciliation deadline expired | true |

Restoration action duration: 1218.374 ms.

| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |
|---|---|---|---|---:|---|---|
| 2026-10-11T15:37:43.824891+08:00 | true | true | false | 11744 | 9ca077b7fcb5357928d6645fc67bf8ed99d96c82 |  |
| 2026-10-11T15:37:46.365441+08:00 | true | true | false | 11744 | 9ca077b7fcb5357928d6645fc67bf8ed99d96c82 |  |
| 2026-10-11T15:37:48.890292+08:00 | true | true | false | 11744 | 9ca077b7fcb5357928d6645fc67bf8ed99d96c82 |  |
| 2026-10-11T15:37:51.430844+08:00 | true | true | false | 11744 | 9ca077b7fcb5357928d6645fc67bf8ed99d96c82 |  |
| 2026-10-11T15:37:53.72937+08:00 | true | true | true | 11744 | 9ca077b7fcb5357928d6645fc67bf8ed99d96c82 |  |

Network fault error: topology and client reconciliation did not stabilize within 100ms

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 1363.587 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 2711.87 | 1348.28 |
| Original replica owns failed primary slots | 3251.40 | 1887.81 |
| Full usable slot coverage | 3251.40 | 1887.81 |
| Surviving views converge | 3251.40 | 1887.81 |
| First all-shard SET/GET probe succeeds | 3045.14 | 1681.56 |
| Client probes sustain recovery | 4543.70 | 3180.11 |
| Topology and clients sustain recovery | 4543.70 | 3180.11 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 421 | 100.00% | 0 | 199.53 | 0.495 / 0.697 / 1.598 |
| fault | 1941 | 96.70% | 64 | 176.00 | 0.472 / 0.732 / 1.380 |
| reconciliation | 150 | 100.00% | 0 | 200.52 | 0.495 / 0.664 / 0.890 |

Acknowledged writes: 1244. Ambiguous SET outcomes: 34.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 221 | 0 | 0 | 0 | 0 | 0 | 0 |
| fault | 943 | 34 | 0 | 0 | 0 | 0 | 0 |
| reconciliation | 80 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 221 | 0.495 / 0.664 / 1.252 |
| baseline | set | 221 | 0.495 / 0.664 / 1.192 |
| fault | Logical write completion | 977 | 0.495 / 0.768 / 1.314 |
| fault | set | 977 | 0.472 / 0.768 / 1.135 |
| reconciliation | Logical write completion | 80 | 0.495 / 0.664 / 0.886 |
| reconciliation | set | 80 | 0.472 / 0.602 / 0.793 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 221 | 104.742 |
| fault | 943 | 85.506 |
| reconciliation | 80 | 106.943 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count issued WAIT commands only.

Global pacing: offered 2777; started 2512; completed 2512; dropped 265. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual workload SET/GET client calls; completed requests count their terminal metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and service probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1244; checked: 1244; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 4 |
| fault | network | 30 |
| fault | timeout | 30 |

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Errors

- experiment outcome: partial

## Measurement methodology

- Fault is a complete bidirectional Docker bridge disconnection of the deterministic primary owning the lowest initial slot. Redis remains running. All peer and published client connectivity is removed; this is not a selective or one-way partition.
- The orchestration and action-completion anchors are recorded separately using monotonic runtime durations. isolation_duration is a minimum hold from successful action completion, bounded by the recovery deadline; samples and bounded process probes can extend the hold. Actual restoration timestamps record the observed duration.
- During isolation, recovery requires the five surviving nodes to converge on healthy usable slot ownership, the original replica to own the selected primary slots, unchanged unaffected ownership, and successful all-shard SET/GET probes over stable_for. A stable milestone is retained while observation continues until restoration.
- Docker exec runs fixed Redis liveness and topology commands independent of the disconnected data plane. PONG and unchanged container PID, start time and Redis run ID prove sampled process survival, not cluster health. Isolation success requires at least one such disconnected observation.
- Restoration reconnects only the same owned node to the same owned network with its original private IP and loopback port. Process identity and endpoint identity must remain unchanged. Cancellation uses a fresh bounded restoration attempt before independent cleanup; safe cleanup can directly remove an owned detached node.
- After restoration, fresh stability tracking requires all six original identities and endpoints, three healthy primaries with synchronized replicas, the promoted replica retaining shard ownership, former-primary demotion, unchanged unaffected shards, and all-shard client probes. Service recovery, connectivity restoration and replication redundancy have separate evidence.
- Poll intervals are delays after sample completion; concurrent views are not atomic snapshots. Probes do not prove continuous availability. Workload metrics include terminal client errors separately from probe results. At most 1200 topology samples and 64 process samples are retained.
- Traffic drains before validating every successfully acknowledged immutable SET. Missing, corrupt, unreadable or unverified records prevent full recovery. Error-returning SET outcomes remain ambiguous. No client can write to the disconnected minority, so this experiment does not exercise minority-side write conflicts or establish loss-free asynchronous replication.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `9f492d82a1064758f96588cf2388be02b2fbf81d4594cac65bbc487e45bb7436`.
