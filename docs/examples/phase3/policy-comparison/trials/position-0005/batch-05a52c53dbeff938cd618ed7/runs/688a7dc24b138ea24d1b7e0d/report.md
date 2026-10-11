# Redis Reliability Lab — policy-aof-waitaof

Run: `688a7dc24b138ea24d1b7e0d`  
Outcome: **recovered**  
Started: 2026-10-11T17:14:10.864572+08:00  
Finished: 2026-10-11T17:14:48.945774+08:00

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

Target: primary `node-1` (Redis ID `5253792ba4414681b143ee5ec61e8a17c6dcc78a`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1550.036 ms.
Stable redundancy first observed 5459.118 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1823.702 |
| Replication interruption observed | 2611.642 |
| Target reachable after repair | 6381.267 |
| Membership rejoined | 6381.267 |
| Replica link up | 10572.642 |
| Synchronization in progress | unobserved |
| Synchronization complete | 10572.642 |
| Primary full-sync counter increased | 6381.267 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 11837.719 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `4874c114517cd1a1dbcd8990ee790ab3e11192ff`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 249.882 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1823.70 | 1573.82 |
| Original replica owns failed primary slots | 2611.64 | 2361.76 |
| Full usable slot coverage | 2611.64 | 2361.76 |
| Surviving views converge | 2611.64 | 2361.76 |
| First all-shard SET/GET probe succeeds | 3531.25 | 3281.37 |
| Client probes sustain recovery | 4712.73 | 4462.84 |
| Topology and clients sustain recovery | 4712.73 | 4462.84 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 803 | 100.00% | 0 | 199.93 | 1.135 / 2.602 / 3.163 |
| fault | 586 | 77.65% | 131 | 121.36 | 0.495 / 2.869 / 41.990 |
| repair | 310 | 81.94% | 56 | 199.99 | 0.546 / 2.602 / 3.163 |
| reconciliation | 1099 | 88.35% | 128 | 200.00 | 0.472 / 2.478 / 3.163 |
| post_recovery | 401 | 100.00% | 0 | 200.42 | 0.981 / 2.602 / 3.163 |

Acknowledged writes: 1586. Ambiguous SET outcomes: 36.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | 5253792ba4414681b143ee5ec61e8a17c6dcc78a | 2026-10-11T17:14:23.186769+08:00 | 5253792ba4414681b143ee5ec61e8a17c6dcc78a | true |  |
| node-2 | 50757b28e88007a7ab9aa6f1382168e02f932243 | 2026-10-11T17:14:24.184435+08:00 | 50757b28e88007a7ab9aa6f1382168e02f932243 | true |  |
| node-3 | 590451160c5462dd0df2c540918d728af059f045 | 2026-10-11T17:14:25.154442+08:00 | 590451160c5462dd0df2c540918d728af059f045 | true |  |
| node-4 | fbaaaad4951f9760ccffd81ca20a2592211018ea | 2026-10-11T17:14:26.154455+08:00 | fbaaaad4951f9760ccffd81ca20a2592211018ea | true |  |
| node-5 | 4874c114517cd1a1dbcd8990ee790ab3e11192ff | 2026-10-11T17:14:27.138805+08:00 | 4874c114517cd1a1dbcd8990ee790ab3e11192ff | true |  |
| node-6 | f6bb5858a2dfd010202363e0fb9045c66598f520 | 2026-10-11T17:14:28.14628+08:00 | f6bb5858a2dfd010202363e0fb9045c66598f520 | true |  |

Before-baseline `waitaof` policy probe at 2026-10-11T17:14:28.157242+08:00; all shards successful: true.
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
| fault | 236 | 36 | 236 | 181 | 55 | 0 | 0 | 236 | 181 |
| repair | 170 | 0 | 170 | 114 | 56 | 0 | 0 | 170 | 114 |
| reconciliation | 565 | 0 | 565 | 437 | 128 | 0 | 0 | 565 | 437 |
| post_recovery | 199 | 0 | 199 | 199 | 0 | 0 | 0 | 199 | 199 |

WAITAOF reply distribution preserves the local/replica pair. Only valid two-integer replies are counted; command errors and unattempted calls have no observed pair. `0/0` is an observed short reply. Local/replica satisfied counts can overlap and must not be added.

| Phase | Local / replica AOF acknowledgements | Replies |
|---|---|---:|
| baseline | 1/1 | 416 |
| fault | 1/0 | 55 |
| fault | 1/1 | 181 |
| repair | 1/0 | 56 |
| repair | 1/1 | 114 |
| reconciliation | 1/0 | 128 |
| reconciliation | 1/1 | 437 |
| post_recovery | 1/1 | 199 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 416 | 1.598 / 2.869 / 3.321 |
| baseline | set | 416 | 0.890 / 1.314 / 2.248 |
| baseline | waitaof | 416 | 0.632 / 1.849 / 2.236 |
| fault | Logical write completion | 272 | 1.761 / 116.982 / 132.115 |
| fault | set | 272 | 0.981 / 1.942 / 2.869 |
| fault | waitaof | 236 | 0.697 / 122.831 / 130.761 |
| repair | Logical write completion | 170 | 1.942 / 116.982 / 128.972 |
| repair | set | 170 | 0.934 / 1.677 / 2.602 |
| repair | waitaof | 170 | 0.768 / 116.982 / 128.972 |
| reconciliation | Logical write completion | 565 | 1.521 / 116.982 / 128.972 |
| reconciliation | set | 565 | 0.847 / 1.314 / 2.478 |
| reconciliation | waitaof | 565 | 0.602 / 116.982 / 128.972 |
| post_recovery | Logical write completion | 199 | 1.380 / 3.012 / 3.487 |
| post_recovery | set | 199 | 0.807 / 1.380 / 2.602 |
| post_recovery | waitaof | 199 | 0.546 / 1.849 / 2.360 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 416 | 103.576 |
| fault | 181 | 37.485 |
| repair | 114 | 73.544 |
| reconciliation | 437 | 79.527 |
| post_recovery | 199 | 99.460 |

SET acknowledgements without a subsequent WAITAOF attempt: 0; client timeouts among those: 0. WAITAOF command errors and WAITAOF client timeouts above count attempted WAITAOF client calls only. A short integer pair is separate from a client timeout.

Configured write policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 3578; started 3199; completed 3199; dropped 379. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1586; checked: 1586; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAITAOF-satisfied subset: expected 1347; checked 1347; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAITAOF.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 6 |
| fault | network | 30 |
| fault | timeout | 40 |
| fault | waitaof_short | 55 |
| reconciliation | waitaof_short | 128 |
| repair | waitaof_short | 56 |

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
