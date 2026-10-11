# Redis Reliability Lab — policy-aof-waitaof

Run: `34b6d0e4fd105ef1b3e0e9e0`  
Outcome: **recovered**  
Started: 2026-10-11T16:59:57.88768+08:00  
Finished: 2026-10-11T17:00:37.215231+08:00

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

Target: primary `node-1` (Redis ID `ae5824d0cbaece9f93ed25014521931e9f5fed78`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1590.417 ms.
Stable redundancy first observed 5819.080 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1859.951 |
| Replication interruption observed | 3539.445 |
| Target reachable after repair | 6465.189 |
| Membership rejoined | 6465.189 |
| Replica link up | 11034.121 |
| Synchronization in progress | unobserved |
| Synchronization complete | 11034.121 |
| Primary full-sync counter increased | 6465.189 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 12281.065 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `5b3767babd8dc8ee1c3e4aceb80173ad3c951fda`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 282.051 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1859.95 | 1577.90 |
| Original replica owns failed primary slots | 3539.44 | 3257.39 |
| Full usable slot coverage | 3539.44 | 3257.39 |
| Surviving views converge | 3539.44 | 3257.39 |
| First all-shard SET/GET probe succeeds | 3565.13 | 3283.08 |
| Client probes sustain recovery | 4752.66 | 4470.61 |
| Topology and clients sustain recovery | 4752.66 | 4470.61 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 803 | 100.00% | 0 | 199.82 | 1.081 / 2.602 / 3.012 |
| fault | 614 | 77.85% | 136 | 126.04 | 0.520 / 3.487 / 62.038 |
| repair | 319 | 83.07% | 54 | 200.57 | 0.520 / 2.869 / 10.711 |
| reconciliation | 1170 | 87.18% | 150 | 199.99 | 0.449 / 2.602 / 5.964 |
| post_recovery | 400 | 100.00% | 0 | 199.87 | 0.573 / 2.602 / 3.662 |

Acknowledged writes: 1635. Ambiguous SET outcomes: 37.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | ae5824d0cbaece9f93ed25014521931e9f5fed78 | 2026-10-11T17:00:10.968449+08:00 | ae5824d0cbaece9f93ed25014521931e9f5fed78 | true |  |
| node-2 | 17980a3c13e0a85b23ec0e6812f2d78fddd77ecd | 2026-10-11T17:00:11.940869+08:00 | 17980a3c13e0a85b23ec0e6812f2d78fddd77ecd | true |  |
| node-3 | 759988c840fa4fd07815143b8bad16a0ae28d2bb | 2026-10-11T17:00:12.977676+08:00 | 759988c840fa4fd07815143b8bad16a0ae28d2bb | true |  |
| node-4 | 9a97744494e5e91d9d4a1552a765943efd45508c | 2026-10-11T17:00:13.976328+08:00 | 9a97744494e5e91d9d4a1552a765943efd45508c | true |  |
| node-5 | 5b3767babd8dc8ee1c3e4aceb80173ad3c951fda | 2026-10-11T17:00:14.950847+08:00 | 5b3767babd8dc8ee1c3e4aceb80173ad3c951fda | true |  |
| node-6 | 1982265c5414f6dce7db879002869bd3da987c5b | 2026-10-11T17:00:15.916055+08:00 | 1982265c5414f6dce7db879002869bd3da987c5b | true |  |

Before-baseline policy probe at 2026-10-11T17:00:15.926612+08:00; all shards successful: true; requested local/replica AOF counts: 1/1.

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
| fault | 259 | 37 | 259 | 200 | 59 | 0 | 0 | 259 | 200 |
| repair | 162 | 0 | 162 | 108 | 54 | 0 | 0 | 162 | 108 |
| reconciliation | 602 | 0 | 602 | 452 | 150 | 0 | 0 | 602 | 452 |
| post_recovery | 196 | 0 | 196 | 196 | 0 | 0 | 0 | 196 | 196 |

WAITAOF reply distribution preserves the local/replica pair. Only valid two-integer replies are counted; command errors and unattempted calls have no observed pair. `0/0` is an observed short reply. Local/replica satisfied counts can overlap and must not be added.

| Phase | Local / replica AOF acknowledgements | Replies |
|---|---|---:|
| baseline | 1/1 | 416 |
| fault | 1/0 | 59 |
| fault | 1/1 | 200 |
| repair | 1/0 | 54 |
| repair | 1/1 | 108 |
| reconciliation | 1/0 | 150 |
| reconciliation | 1/1 | 452 |
| post_recovery | 1/1 | 196 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 416 | 1.449 / 2.869 / 3.012 |
| baseline | set | 416 | 0.847 / 1.521 / 2.248 |
| baseline | waitaof | 416 | 0.573 / 1.849 / 2.039 |
| fault | Logical write completion | 296 | 1.521 / 116.982 / 128.972 |
| fault | set | 296 | 0.847 / 2.141 / 6.905 |
| fault | waitaof | 259 | 0.632 / 116.982 / 128.972 |
| repair | Logical write completion | 162 | 2.039 / 135.421 / 181.477 |
| repair | set | 162 | 0.981 / 2.478 / 9.253 |
| repair | waitaof | 162 | 0.768 / 128.972 / 181.477 |
| reconciliation | Logical write completion | 602 | 1.598 / 122.831 / 128.972 |
| reconciliation | set | 602 | 0.847 / 2.248 / 4.451 |
| reconciliation | waitaof | 602 | 0.602 / 116.982 / 128.972 |
| post_recovery | Logical write completion | 196 | 1.380 / 3.012 / 4.037 |
| post_recovery | set | 196 | 0.847 / 1.942 / 2.360 |
| post_recovery | waitaof | 196 | 0.520 / 1.761 / 2.360 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 416 | 103.521 |
| fault | 200 | 41.055 |
| repair | 108 | 67.905 |
| reconciliation | 452 | 77.262 |
| post_recovery | 196 | 97.936 |

SET acknowledgements without a subsequent WAITAOF attempt: 0; client timeouts among those: 0. WAITAOF command errors and WAITAOF client timeouts above count attempted WAITAOF client calls only. A short integer pair is separate from a client timeout.

Configured write policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 3666; started 3306; completed 3306; dropped 360. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1635; checked: 1635; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAITAOF-satisfied subset: expected 1372; checked 1372; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAITAOF.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 4 |
| fault | network | 33 |
| fault | timeout | 40 |
| fault | waitaof_short | 59 |
| reconciliation | waitaof_short | 150 |
| repair | waitaof_short | 54 |

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
- Traffic drains before validating every acknowledged immutable SET, including short/error WAIT outcomes. Ambiguous SET errors remain separate. Cancellation skips unsupported recovery/data claims and still uses fresh bounded restoration and cleanup contexts, including managed volumes.
- Configured acknowledgement policy is a separate current condition, measured through all-three-shard SET plus WAIT/WAITAOF on one physical connection. A stable policy milestone is historical; a later unsuccessful probe resets the current condition.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `975080021092fa9a75ee5da95f13abfb2ac4db2ce23d3e320a4867ae224e9637`.
