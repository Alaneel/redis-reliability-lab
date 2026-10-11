# Redis Reliability Lab — primary-isolation

Run: `bbba656458a3cb89130c16de`  
Outcome: **recovered**  
Started: 2026-10-11T15:27:31.249078+08:00  
Finished: 2026-10-11T15:28:01.622473+08:00

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

Restoration action duration: 1286.372 ms.
Stable reconciliation first observed 1238.500 ms after restoration completed.

| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |
|---|---|---|---|---:|---|---|
| 2026-10-11T15:27:47.139609+08:00 | true | true | false | 15692 | faf98cf12d7ca8bd9c4fc4b1ac59438037807767 |  |
| 2026-10-11T15:27:49.962877+08:00 | true | true | false | 15692 | faf98cf12d7ca8bd9c4fc4b1ac59438037807767 |  |
| 2026-10-11T15:27:52.697026+08:00 | true | true | false | 15692 | faf98cf12d7ca8bd9c4fc4b1ac59438037807767 |  |
| 2026-10-11T15:27:57.684662+08:00 | true | true | true | 15692 | faf98cf12d7ca8bd9c4fc4b1ac59438037807767 |  |

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 1363.464 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1899.79 | 536.32 |
| Original replica owns failed primary slots | 3442.66 | 2079.20 |
| Full usable slot coverage | 3442.66 | 2079.20 |
| Surviving views converge | 3442.66 | 2079.20 |
| First all-shard SET/GET probe succeeds | 4745.36 | 3381.90 |
| Client probes sustain recovery | 5852.77 | 4489.31 |
| Topology and clients sustain recovery | 5852.77 | 4489.31 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 404 | 100.00% | 0 | 199.76 | 0.520 / 0.768 / 1.449 |
| fault | 1610 | 97.02% | 48 | 150.90 | 0.495 / 0.697 / 1.252 |
| reconciliation | 355 | 100.00% | 0 | 200.12 | 0.472 / 0.697 / 1.252 |
| post_recovery | 400 | 100.00% | 0 | 199.90 | 0.472 / 0.697 / 0.807 |

Acknowledged writes: 1389. Ambiguous SET outcomes: 26.

Global pacing: offered 3293; started 2769; completed 2769; dropped 524. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual workload SET/GET client calls; completed requests count their terminal metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and service probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1389; checked: 1389; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | timeout | 48 |

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `c2dc3ff051851a79194fab4e00956840ba1994cb0b21cbffce91c8011b5b0880`.
