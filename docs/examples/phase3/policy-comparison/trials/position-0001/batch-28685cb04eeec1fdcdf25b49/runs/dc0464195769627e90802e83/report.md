# Redis Reliability Lab — policy-aof-wait

Run: `dc0464195769627e90802e83`  
Outcome: **recovered**  
Started: 2026-10-11T17:11:38.886901+08:00  
Finished: 2026-10-11T17:12:17.592128+08:00

## Recovery

| Condition | Observed |
|---|---|
| Surviving topology converged | true |
| All-shard client recovery sustained | true |
| Topology and client recovery jointly sustained | true |
| All acknowledged writes verified | true |
| Target process restarted | true |
| Original replication redundancy restored | true |
| Recovery deadline expired | false |
| Cleanup verified | true |


## Managed node lifecycle

Target: primary `node-1` (Redis ID `f3b9eace3d420a8e58af6efd3e321637c1e9eb27`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1541.732 ms.
Stable redundancy first observed 4983.597 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1825.250 |
| Replication interruption observed | 2615.063 |
| Target reachable after repair | 6407.196 |
| Membership rejoined | 6407.196 |
| Replica link up | 10118.887 |
| Synchronization in progress | unobserved |
| Synchronization complete | 10118.887 |
| Primary full-sync counter increased | 6407.196 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 11388.004 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `daccf40e5fd1979e36ac63c8b99faaed53cbd8d0`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

## Independent recovery dimensions

Target observed role: `slave`.

| Independent condition | Observed |
|---|---|
| Target process live | true |
| Cluster membership healthy | true |
| Full usable slot coverage | true |
| All shard leaders available | true |
| Client reads available | true |
| Client writes available | true |
| Acknowledged data verified | true |
| Configured write policy available | true |
| Replication redundancy restored | true |
| Persistence recovery verified | unobserved |

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 236.903 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1825.25 | 1588.35 |
| Original replica owns failed primary slots | 2615.06 | 2378.16 |
| Full usable slot coverage | 3510.96 | 3274.06 |
| Surviving views converge | 3510.96 | 3274.06 |
| First all-shard SET/GET probe succeeds | 3128.11 | 2891.21 |
| Client probes sustain recovery | 4333.92 | 4097.02 |
| Topology and clients sustain recovery | 4728.71 | 4491.81 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 803 | 100.00% | 0 | 199.84 | 1.192 / 2.732 / 4.673 |
| fault | 608 | 77.14% | 139 | 125.03 | 0.546 / 3.321 / 436.746 |
| repair | 308 | 82.47% | 54 | 199.77 | 0.520 / 2.732 / 4.037 |
| reconciliation | 1004 | 88.65% | 114 | 200.04 | 0.546 / 2.732 / 4.037 |
| post_recovery | 400 | 100.00% | 0 | 199.98 | 0.546 / 2.732 / 3.321 |

Acknowledged writes: 1530. Ambiguous SET outcomes: 47.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | f3b9eace3d420a8e58af6efd3e321637c1e9eb27 | 2026-10-11T17:11:52.104358+08:00 | f3b9eace3d420a8e58af6efd3e321637c1e9eb27 | true |  |
| node-2 | 34a4e713deeaf072b200615fde2fff3ae55ef769 | 2026-10-11T17:11:53.080723+08:00 | 34a4e713deeaf072b200615fde2fff3ae55ef769 | true |  |
| node-3 | 416514404167e8ad69d052c0075552aeceb3f80d | 2026-10-11T17:11:54.073425+08:00 | 416514404167e8ad69d052c0075552aeceb3f80d | true |  |
| node-4 | 48a7692490d492489e6b26d6708ca56f7707c09a | 2026-10-11T17:11:55.176618+08:00 | 48a7692490d492489e6b26d6708ca56f7707c09a | true |  |
| node-5 | daccf40e5fd1979e36ac63c8b99faaed53cbd8d0 | 2026-10-11T17:11:56.201815+08:00 | daccf40e5fd1979e36ac63c8b99faaed53cbd8d0 | true |  |
| node-6 | 9070e779cb261545f5728fc5e4c00699a4104769 | 2026-10-11T17:11:57.16589+08:00 | 9070e779cb261545f5728fc5e4c00699a4104769 | true |  |

Before-baseline `wait` policy probe at 2026-10-11T17:11:57.17842+08:00; all shards successful: true.
Requested replica-memory count: 1.

| Original shard index | SET acknowledged | WAIT attempted | Replica-memory reply |
|---|---|---|---|
| 0 | true | true | 1 |
| 1 | true | true | 1 |
| 2 | true | true | 1 |

Write policy: `wait`. Diagnostic stage requests are excluded from logical workload counts.
Logical write latency includes SET and WAIT on the same physical connection; policy success requires the requested replica-memory acknowledgement count.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 416 | 0 | 416 | 416 | 0 | 0 | 0 |
| fault | 240 | 47 | 240 | 183 | 57 | 0 | 0 |
| repair | 162 | 0 | 162 | 108 | 54 | 0 | 0 |
| reconciliation | 523 | 0 | 523 | 409 | 114 | 0 | 0 |
| post_recovery | 189 | 0 | 189 | 189 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 416 | 1.598 / 3.012 / 7.250 |
| baseline | set | 416 | 0.981 / 1.761 / 5.410 |
| baseline | wait | 416 | 0.602 / 1.942 / 2.360 |
| fault | Logical write completion | 287 | 1.761 / 116.982 / 126.227 |
| fault | set | 287 | 0.981 / 2.360 / 4.239 |
| fault | wait | 240 | 0.697 / 116.982 / 122.831 |
| repair | Logical write completion | 162 | 1.849 / 116.982 / 135.421 |
| repair | set | 162 | 1.030 / 1.598 / 2.360 |
| repair | wait | 162 | 0.732 / 116.982 / 134.971 |
| reconciliation | Logical write completion | 523 | 1.598 / 116.982 / 122.831 |
| reconciliation | set | 523 | 0.847 / 2.039 / 2.602 |
| reconciliation | wait | 523 | 0.602 / 116.982 / 122.831 |
| post_recovery | Logical write completion | 189 | 1.449 / 3.163 / 4.451 |
| post_recovery | set | 189 | 0.890 / 1.380 / 3.662 |
| post_recovery | wait | 189 | 0.546 / 1.942 / 2.248 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 416 | 103.529 |
| fault | 183 | 37.634 |
| repair | 108 | 70.049 |
| reconciliation | 409 | 81.492 |
| post_recovery | 189 | 94.490 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Configured write policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 3488; started 3123; completed 3123; dropped 365. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1530; checked: 1530; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAIT-satisfied subset: expected 1305; checked 1305; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAIT.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 14 |
| fault | network | 33 |
| fault | timeout | 35 |
| fault | wait_short | 57 |
| reconciliation | wait_short | 114 |
| repair | wait_short | 54 |

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Measurement methodology

- Fault selection preserves the initial primary and its replica. The destructive target is the selected primary after stable five-survivor service recovery; only strictly owned local Docker resources are affected.
- Replica faults preserve all original slot owners and do not require promotion. down_duration is a minimum hold from successful fault action completion, bounded by the recovery observation budget. Client service, failure gossip and replication interruption are separately sampled.
- Replica network isolation removes its complete bidirectional Docker bridge connectivity while Redis remains running. It is not a selective replication-only partition. Restoration uses the same owned endpoint and unchanged live process identity.
- Administrative endpoint-disconnection evidence is separate from Redis INFO link or connected-replica indicators. A short disconnection can finish before TCP/link failure detection; absent interruption or synchronization-counter milestones remain unobserved and are not invented from the Docker interface state.
- Crash repair starts the same stopped owned container with its bounded disposable data volume. Redis Cluster identity must match the pre-fault identity. Restart changes the process; it does not simulate power loss, host/storage failure or persistence durability.
- The repair deadline covers the restart/reconnect action and subsequent synchronization. Redis's current membership is inspected without issuing CLUSTER REPLICATE or forcing slot ownership. Replica repair preserves original assignments; former-primary repair requires the original replica to retain promoted ownership and the old primary to rejoin as its replica.
- Fresh stability after repair requires all six original identities and endpoints, healthy equal slot maps, exactly one synchronized replica per primary, and all-shard client probes. Historical service recovery remains separate from final redundancy and configured acknowledgement-policy availability.
- Replication link-up and master_sync_in_progress=0 are sampled indicators, not exact zero lag, bytes preserved, or identification of partial versus full resynchronization. Raw replication INFO is retained for inspection; unobserved milestones remain absent.
- Traffic drains before validating every acknowledged immutable SET, including short/error WAIT/WAITAOF outcomes. Ambiguous SET errors remain separate. Cancellation skips unsupported recovery/data claims and still uses fresh bounded restoration and cleanup contexts, including managed volumes.
- All AOF modes and write policies use six-node quiet runtime preparation followed by exactly three all-shard policy SET/GET warmup writes before baseline starts. Warmup writes are excluded from the workload ledger and metrics; no WAITAOF compatibility command is issued after baseline to alter crash-trial durability.
- Configured acknowledgement policy is a separate current condition, measured through all-three-shard SET plus WAIT/WAITAOF on one physical connection. A stable policy milestone is historical; a later unsuccessful probe resets the current condition.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `d7e106ea5cb6011eb4736e886b40e76b334dd77ec754a2cae619035e090e1ef9`.
