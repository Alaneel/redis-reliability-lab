# Redis Reliability Lab — primary-network-isolation

Run: `e4544d8f12c536abaedb1905`  
Outcome: **aborted**  
Started: 2026-10-11T15:38:15.873259+08:00  
Finished: 2026-10-11T15:38:33.415668+08:00

## Recovery

| Condition | Observed |
|---|---|
| Surviving topology converged | false |
| All-shard client recovery sustained | false |
| Topology and client recovery jointly sustained | false |
| All acknowledged writes verified | false |
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
| Reconciliation deadline expired | false |

Restoration action duration: 1198.808 ms.

| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |
|---|---|---|---|---:|---|---|

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 1319.161 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | unobserved | unobserved |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | unobserved | unobserved |
| Surviving views converge | unobserved | unobserved |
| First all-shard SET/GET probe succeeds | unobserved | unobserved |
| Client probes sustain recovery | unobserved | unobserved |
| Topology and clients sustain recovery | unobserved | unobserved |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 420 | 100.00% | 0 | 199.47 | 0.546 / 0.847 / 1.314 |
| fault | 213 | 87.79% | 26 | 161.45 | 0.520 / 0.890 / 1.449 |

Acknowledged writes: 322. Ambiguous SET outcomes: 14.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 221 | 0 | 0 | 0 | 0 | 0 | 0 |
| fault | 101 | 14 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 221 | 0.546 / 0.890 / 1.252 |
| baseline | set | 221 | 0.546 / 0.847 / 1.192 |
| fault | Logical write completion | 115 | 0.546 / 0.890 / 1.192 |
| fault | set | 115 | 0.520 / 0.807 / 0.981 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 221 | 104.959 |
| fault | 101 | 76.558 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count issued WAIT commands only.

Global pacing: offered 684; started 633; completed 633; dropped 51. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual workload SET/GET client calls; completed requests count their terminal metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and service probes are excluded. These global counters are not attributed to phases.

Validation attempted: false; expected: 0; checked: 0; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: false.
Validation was not attempted; 322 acknowledged writes remain unchecked.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | canceled | 7 |
| fault | network | 14 |
| fault | timeout | 5 |

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Errors

- context canceled

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
