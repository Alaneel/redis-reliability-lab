# Redis Reliability Lab — policy-aof-set

Run: `60141689fee92e0e4cd9bc77`  
Outcome: **recovered**  
Started: 2026-10-11T17:13:33.704203+08:00  
Finished: 2026-10-11T17:14:10.688361+08:00

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

Target: primary `node-1` (Redis ID `00608c51212cc3bd03e9accaeed493cf7adf5533`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1570.270 ms.
Stable redundancy first observed 4725.556 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1807.338 |
| Replication interruption observed | 2590.549 |
| Target reachable after repair | 6103.880 |
| Membership rejoined | 6103.880 |
| Replica link up | 9549.942 |
| Synchronization in progress | unobserved |
| Synchronization complete | 9549.942 |
| Primary full-sync counter increased | 6103.880 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 10826.177 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `2ad9648cabe877cd6ae4ed38eb640e0e34700d6b`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 233.343 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1807.34 | 1573.99 |
| Original replica owns failed primary slots | 2590.55 | 2357.21 |
| Full usable slot coverage | 2590.55 | 2357.21 |
| Surviving views converge | 2590.55 | 2357.21 |
| First all-shard SET/GET probe succeeds | 3406.26 | 3172.91 |
| Client probes sustain recovery | 4526.97 | 4293.63 |
| Topology and clients sustain recovery | 4526.97 | 4293.63 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 803 | 100.00% | 0 | 199.73 | 0.807 / 1.380 / 2.248 |
| fault | 507 | 86.39% | 69 | 111.91 | 0.732 / 2.478 / 135.421 |
| repair | 314 | 100.00% | 0 | 199.96 | 0.732 / 1.598 / 3.487 |
| reconciliation | 952 | 100.00% | 0 | 199.95 | 0.632 / 1.252 / 2.248 |
| post_recovery | 400 | 100.00% | 0 | 199.98 | 0.697 / 1.252 / 2.360 |

Acknowledged writes: 1494. Ambiguous SET outcomes: 28.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | 00608c51212cc3bd03e9accaeed493cf7adf5533 | 2026-10-11T17:13:45.943858+08:00 | 00608c51212cc3bd03e9accaeed493cf7adf5533 | true |  |
| node-2 | ea443a21b8ac0abb54dc72c3b9b1d33de01afb57 | 2026-10-11T17:13:46.916998+08:00 | ea443a21b8ac0abb54dc72c3b9b1d33de01afb57 | true |  |
| node-3 | cdd4c2a069c2e11488c8ec019595bbaecd77b8a4 | 2026-10-11T17:13:47.883787+08:00 | cdd4c2a069c2e11488c8ec019595bbaecd77b8a4 | true |  |
| node-4 | db2375b1b80fa5a6351abc1d788dfdef19bfcadc | 2026-10-11T17:13:48.875378+08:00 | db2375b1b80fa5a6351abc1d788dfdef19bfcadc | true |  |
| node-5 | 2ad9648cabe877cd6ae4ed38eb640e0e34700d6b | 2026-10-11T17:13:49.871894+08:00 | 2ad9648cabe877cd6ae4ed38eb640e0e34700d6b | true |  |
| node-6 | f2300b47eaa9f74130fc10589c78d5163fec3f1c | 2026-10-11T17:13:50.841403+08:00 | f2300b47eaa9f74130fc10589c78d5163fec3f1c | true |  |

Before-baseline `set` policy probe at 2026-10-11T17:13:50.852062+08:00; all shards successful: true.

| Original shard index | SET acknowledged |
|---|---|
| 0 | true |
| 1 | true |
| 2 | true |

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 416 | 0 |
| fault | 211 | 28 |
| repair | 161 | 0 |
| reconciliation | 495 | 0 |
| post_recovery | 211 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 416 | 1.030 / 1.521 / 2.602 |
| baseline | set | 416 | 1.030 / 1.521 / 2.602 |
| fault | Logical write completion | 239 | 1.030 / 2.732 / 3.845 |
| fault | set | 239 | 1.030 / 2.732 / 3.845 |
| repair | Logical write completion | 161 | 0.981 / 2.732 / 4.907 |
| repair | set | 161 | 0.981 / 2.602 / 4.907 |
| reconciliation | Logical write completion | 495 | 0.847 / 1.521 / 2.478 |
| reconciliation | set | 495 | 0.847 / 1.521 / 2.478 |
| post_recovery | Logical write completion | 211 | 0.847 / 1.521 / 2.478 |
| post_recovery | set | 211 | 0.847 / 1.521 / 2.478 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 416 | 103.473 |
| fault | 211 | 46.575 |
| repair | 161 | 102.526 |
| reconciliation | 495 | 103.968 |
| post_recovery | 211 | 105.487 |

Global pacing: offered 3376; started 2976; completed 2976; dropped 400. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1494; checked: 1494; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 3 |
| fault | network | 25 |
| fault | timeout | 41 |

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `de05e21df2ed5b0a82dc0af8fa135f9173da7c1fdf43019ed7cd93af6d12f3cf`.
