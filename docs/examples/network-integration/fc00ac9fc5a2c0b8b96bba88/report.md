# Redis Reliability Lab — primary-network-isolation

Run: `fc00ac9fc5a2c0b8b96bba88`  
Outcome: **recovered**  
Started: 2026-10-11T15:36:55.932241+08:00  
Finished: 2026-10-11T15:37:26.796038+08:00

## Recovery

| Condition | Observed |
|---|---|
| Surviving topology converged | true |
| All-shard client recovery sustained | true |
| Topology and client recovery jointly sustained | true |
| All acknowledged writes verified | true |
| Killed process restarted | false |
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

Restoration action duration: 1247.796 ms.
Stable reconciliation first observed 2145.407 ms after restoration completed.

| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |
|---|---|---|---|---:|---|---|
| 2026-10-11T15:37:11.223484+08:00 | true | true | false | 9443 | 73cebbcbe4bf2fb3ea581c5a34a21e6b2834e035 |  |
| 2026-10-11T15:37:14.227247+08:00 | true | true | false | 9443 | 73cebbcbe4bf2fb3ea581c5a34a21e6b2834e035 |  |
| 2026-10-11T15:37:16.75867+08:00 | true | true | false | 9443 | 73cebbcbe4bf2fb3ea581c5a34a21e6b2834e035 |  |
| 2026-10-11T15:37:22.831334+08:00 | true | true | true | 9443 | 73cebbcbe4bf2fb3ea581c5a34a21e6b2834e035 |  |

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 1344.145 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1844.94 | 500.80 |
| Original replica owns failed primary slots | 3509.50 | 2165.36 |
| Full usable slot coverage | 3509.50 | 2165.36 |
| Surviving views converge | 3509.50 | 2165.36 |
| First all-shard SET/GET probe succeeds | 4976.80 | 3632.66 |
| Client probes sustain recovery | 5983.40 | 4639.26 |
| Topology and clients sustain recovery | 5983.40 | 4639.26 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 419 | 100.00% | 0 | 199.64 | 0.449 / 0.573 / 1.521 |
| fault | 1791 | 94.53% | 98 | 168.24 | 0.472 / 0.768 / 1.314 |
| reconciliation | 560 | 100.00% | 0 | 200.06 | 0.449 / 0.632 / 0.768 |
| post_recovery | 400 | 100.00% | 0 | 199.94 | 0.428 / 0.573 / 0.732 |

Acknowledged writes: 1547. Ambiguous SET outcomes: 57.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 220 | 0 | 0 | 0 | 0 | 0 | 0 |
| fault | 838 | 57 | 0 | 0 | 0 | 0 | 0 |
| reconciliation | 295 | 0 | 0 | 0 | 0 | 0 | 0 |
| post_recovery | 194 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 220 | 0.449 / 0.573 / 1.380 |
| baseline | set | 220 | 0.428 / 0.573 / 1.314 |
| fault | Logical write completion | 895 | 0.495 / 0.807 / 1.252 |
| fault | set | 895 | 0.472 / 0.768 / 1.135 |
| reconciliation | Logical write completion | 295 | 0.449 / 0.664 / 1.135 |
| reconciliation | set | 295 | 0.449 / 0.632 / 0.981 |
| post_recovery | Logical write completion | 194 | 0.449 / 0.602 / 0.732 |
| post_recovery | set | 194 | 0.428 / 0.573 / 0.697 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 220 | 104.824 |
| fault | 838 | 78.721 |
| reconciliation | 295 | 105.390 |
| post_recovery | 194 | 96.971 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count issued WAIT commands only.

Global pacing: offered 3508; started 3170; completed 3170; dropped 338. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual workload SET/GET client calls; completed requests count their terminal metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and service probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1547; checked: 1547; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 6 |
| fault | network | 51 |
| fault | timeout | 41 |

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `1ee187349b5fed1612b885597e30e3715a6edbd937cf35e1b8be33975cffb09c`.
