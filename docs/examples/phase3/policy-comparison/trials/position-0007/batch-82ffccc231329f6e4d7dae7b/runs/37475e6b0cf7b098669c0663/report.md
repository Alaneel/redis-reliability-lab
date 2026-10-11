# Redis Reliability Lab — policy-aof-waitaof

Run: `37475e6b0cf7b098669c0663`  
Outcome: **recovered**  
Started: 2026-10-11T17:15:27.137138+08:00  
Finished: 2026-10-11T17:16:05.240204+08:00

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

Target: primary `node-1` (Redis ID `eba01c15acef5a958a2efff862b66b0919d499a8`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1499.688 ms.
Stable redundancy first observed 5949.375 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1813.686 |
| Replication interruption observed | 3450.078 |
| Target reachable after repair | 6253.872 |
| Membership rejoined | 6253.872 |
| Replica link up | 10913.946 |
| Synchronization in progress | unobserved |
| Synchronization complete | 10913.946 |
| Primary full-sync counter increased | 6253.872 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 12200.501 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `21a9a47a35222fc8669812e147ff9f4e864af72e`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 241.023 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1813.69 | 1572.66 |
| Original replica owns failed primary slots | 3450.08 | 3209.05 |
| Full usable slot coverage | 3450.08 | 3209.05 |
| Surviving views converge | 3450.08 | 3209.05 |
| First all-shard SET/GET probe succeeds | 3077.33 | 2836.31 |
| Client probes sustain recovery | 4254.83 | 4013.80 |
| Topology and clients sustain recovery | 4638.86 | 4397.84 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 803 | 100.00% | 0 | 199.81 | 1.314 / 2.478 / 3.163 |
| fault | 575 | 77.91% | 127 | 121.02 | 0.697 / 3.012 / 231.616 |
| repair | 300 | 82.33% | 53 | 200.04 | 0.495 / 2.248 / 3.662 |
| reconciliation | 1198 | 87.73% | 147 | 200.13 | 0.472 / 2.732 / 3.487 |
| post_recovery | 400 | 100.00% | 0 | 199.96 | 0.934 / 2.869 / 3.845 |

Acknowledged writes: 1629. Ambiguous SET outcomes: 33.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | eba01c15acef5a958a2efff862b66b0919d499a8 | 2026-10-11T17:15:39.054273+08:00 | eba01c15acef5a958a2efff862b66b0919d499a8 | true |  |
| node-2 | 368831204b55115e8cd536248d21e4315c00adf7 | 2026-10-11T17:15:40.063644+08:00 | 368831204b55115e8cd536248d21e4315c00adf7 | true |  |
| node-3 | b18cb8b0745aeb1cd68ee6acf028b4f1153b43fd | 2026-10-11T17:15:41.07829+08:00 | b18cb8b0745aeb1cd68ee6acf028b4f1153b43fd | true |  |
| node-4 | 204a3d38a19fcabe6d3f23c934fcc6756ffeaae1 | 2026-10-11T17:15:42.086197+08:00 | 204a3d38a19fcabe6d3f23c934fcc6756ffeaae1 | true |  |
| node-5 | 21a9a47a35222fc8669812e147ff9f4e864af72e | 2026-10-11T17:15:43.075221+08:00 | 21a9a47a35222fc8669812e147ff9f4e864af72e | true |  |
| node-6 | 03d79913768e7d1fb7168d60a1077664cd39f266 | 2026-10-11T17:15:44.051188+08:00 | 03d79913768e7d1fb7168d60a1077664cd39f266 | true |  |

Before-baseline `waitaof` policy probe at 2026-10-11T17:15:44.06177+08:00; all shards successful: true.
requested local/replica AOF counts: 1/1.

| Original shard index | SET acknowledged | WAITAOF attempted | Local / replica AOF reply |
|---|---|---|---|
| 0 | true | true | 1/1 |
| 1 | true | true | 1/1 |
| 2 | true | true | 1/1 |

Write policy: `waitaof`. Diagnostic stage requests are excluded from logical workload counts.
Logical write latency includes SET and WAITAOF on the same physical connection. WAITAOF policy success requires both requested local and replica AOF acknowledgement counts. These disk acknowledgements are separate from WAIT replica-memory acknowledgements; they do not establish strong consistency or a power-loss guarantee from this process-crash experiment.

| Phase | SET acknowledged | SET errors | WAITAOF attempted | Both satisfied | Short pair | Command errors | Client timeouts | Local satisfied | Replica satisfied |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline | 416 | 0 | 416 | 416 | 0 | 0 | 0 | 416 | 416 |
| fault | 240 | 33 | 240 | 185 | 55 | 0 | 0 | 240 | 185 |
| repair | 158 | 0 | 158 | 105 | 53 | 0 | 0 | 158 | 105 |
| reconciliation | 619 | 0 | 619 | 472 | 147 | 0 | 0 | 619 | 472 |
| post_recovery | 196 | 0 | 196 | 196 | 0 | 0 | 0 | 196 | 196 |

WAITAOF reply distribution preserves the local/replica pair. Only valid two-integer replies are counted; command errors and unattempted calls have no observed pair. `0/0` is an observed short reply. Local/replica satisfied counts can overlap and must not be added.

| Phase | Local / replica AOF acknowledgements | Replies |
|---|---|---:|
| baseline | 1/1 | 416 |
| fault | 1/0 | 55 |
| fault | 1/1 | 185 |
| repair | 1/0 | 53 |
| repair | 1/1 | 105 |
| reconciliation | 1/0 | 147 |
| reconciliation | 1/1 | 472 |
| post_recovery | 1/1 | 196 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 416 | 1.677 / 2.869 / 3.487 |
| baseline | set | 416 | 1.030 / 1.521 / 2.478 |
| baseline | waitaof | 416 | 0.632 / 1.677 / 2.141 |
| fault | Logical write completion | 273 | 1.849 / 111.411 / 128.972 |
| fault | set | 273 | 0.981 / 2.141 / 4.451 |
| fault | waitaof | 240 | 0.768 / 111.411 / 122.831 |
| repair | Logical write completion | 158 | 1.849 / 128.972 / 135.223 |
| repair | set | 158 | 0.934 / 1.380 / 2.039 |
| repair | waitaof | 158 | 0.697 / 128.972 / 132.845 |
| reconciliation | Logical write completion | 619 | 1.598 / 116.982 / 128.972 |
| reconciliation | set | 619 | 0.847 / 1.314 / 2.478 |
| reconciliation | waitaof | 619 | 0.632 / 116.982 / 128.972 |
| post_recovery | Logical write completion | 196 | 1.449 / 3.321 / 5.964 |
| post_recovery | set | 196 | 0.847 / 2.141 / 4.239 |
| post_recovery | waitaof | 196 | 0.573 / 1.849 / 2.186 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 416 | 103.513 |
| fault | 185 | 38.936 |
| repair | 105 | 70.013 |
| reconciliation | 472 | 78.850 |
| post_recovery | 196 | 97.983 |

SET acknowledgements without a subsequent WAITAOF attempt: 0; client timeouts among those: 0. WAITAOF command errors and WAITAOF client timeouts above count attempted WAITAOF client calls only. A short integer pair is separate from a client timeout.

Configured write policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 3651; started 3276; completed 3276; dropped 375. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1629; checked: 1629; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAITAOF-satisfied subset: expected 1374; checked 1374; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAITAOF.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 4 |
| fault | network | 29 |
| fault | timeout | 39 |
| fault | waitaof_short | 55 |
| reconciliation | waitaof_short | 147 |
| repair | waitaof_short | 53 |

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `975080021092fa9a75ee5da95f13abfb2ac4db2ce23d3e320a4867ae224e9637`.
