# Redis Reliability Lab — policy-aof-waitaof

Run: `e519c821277d8791d7e6796e`  
Outcome: **recovered**  
Started: 2026-10-11T17:12:55.6879+08:00  
Finished: 2026-10-11T17:13:33.537038+08:00

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

Target: primary `node-1` (Redis ID `669d97916061711ff170ce2badb6d1418c925d3f`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1559.327 ms.
Stable redundancy first observed 5098.398 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1801.791 |
| Replication interruption observed | 2586.054 |
| Target reachable after repair | 6329.500 |
| Membership rejoined | 6329.500 |
| Replica link up | 10142.302 |
| Synchronization in progress | unobserved |
| Synchronization complete | 10142.302 |
| Primary full-sync counter increased | 6329.500 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 11424.893 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `dc2f255a8e6c94fcdd9b914bd700de2253da65e5`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 236.692 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1801.79 | 1565.10 |
| Original replica owns failed primary slots | 2586.05 | 2349.36 |
| Full usable slot coverage | 2586.05 | 2349.36 |
| Surviving views converge | 2586.05 | 2349.36 |
| First all-shard SET/GET probe succeeds | 3497.77 | 3261.08 |
| Client probes sustain recovery | 4656.21 | 4419.52 |
| Topology and clients sustain recovery | 4656.21 | 4419.52 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 803 | 100.00% | 0 | 199.83 | 1.252 / 2.360 / 2.869 |
| fault | 584 | 76.88% | 135 | 122.50 | 0.546 / 3.163 / 255.356 |
| repair | 311 | 82.96% | 53 | 199.44 | 0.495 / 2.869 / 6.576 |
| reconciliation | 1027 | 87.54% | 128 | 200.04 | 0.495 / 2.478 / 4.239 |
| post_recovery | 400 | 100.00% | 0 | 199.99 | 0.573 / 1.942 / 2.869 |

Acknowledged writes: 1544. Ambiguous SET outcomes: 42.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | 669d97916061711ff170ce2badb6d1418c925d3f | 2026-10-11T17:13:08.21381+08:00 | 669d97916061711ff170ce2badb6d1418c925d3f | true |  |
| node-2 | a4df1ee5c962798f4a06c83805ef510aa688ec30 | 2026-10-11T17:13:09.209249+08:00 | a4df1ee5c962798f4a06c83805ef510aa688ec30 | true |  |
| node-3 | 7c7d564dedd3c3fe6391dd69dc6ce46b935ba735 | 2026-10-11T17:13:10.174669+08:00 | 7c7d564dedd3c3fe6391dd69dc6ce46b935ba735 | true |  |
| node-4 | f807a791aa679679772ba401b779651fba235b65 | 2026-10-11T17:13:11.161355+08:00 | f807a791aa679679772ba401b779651fba235b65 | true |  |
| node-5 | dc2f255a8e6c94fcdd9b914bd700de2253da65e5 | 2026-10-11T17:13:12.165195+08:00 | dc2f255a8e6c94fcdd9b914bd700de2253da65e5 | true |  |
| node-6 | 31d025121b5f735836a8ae69c8c94b03b8f8d69e | 2026-10-11T17:13:13.157233+08:00 | 31d025121b5f735836a8ae69c8c94b03b8f8d69e | true |  |

Before-baseline `waitaof` policy probe at 2026-10-11T17:13:13.167513+08:00; all shards successful: true.
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
| fault | 234 | 42 | 234 | 179 | 55 | 0 | 0 | 234 | 179 |
| repair | 158 | 0 | 158 | 105 | 53 | 0 | 0 | 158 | 105 |
| reconciliation | 545 | 0 | 545 | 417 | 128 | 0 | 0 | 545 | 417 |
| post_recovery | 191 | 0 | 191 | 191 | 0 | 0 | 0 | 191 | 191 |

WAITAOF reply distribution preserves the local/replica pair. Only valid two-integer replies are counted; command errors and unattempted calls have no observed pair. `0/0` is an observed short reply. Local/replica satisfied counts can overlap and must not be added.

| Phase | Local / replica AOF acknowledgements | Replies |
|---|---|---:|
| baseline | 1/1 | 416 |
| fault | 1/0 | 55 |
| fault | 1/1 | 179 |
| repair | 1/0 | 53 |
| repair | 1/1 | 105 |
| reconciliation | 1/0 | 128 |
| reconciliation | 1/1 | 417 |
| post_recovery | 1/1 | 191 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 416 | 1.598 / 2.732 / 3.012 |
| baseline | set | 416 | 0.981 / 1.521 / 2.602 |
| baseline | waitaof | 416 | 0.602 / 1.598 / 1.849 |
| fault | Logical write completion | 276 | 1.761 / 116.982 / 135.421 |
| fault | set | 276 | 0.934 / 1.942 / 2.602 |
| fault | waitaof | 234 | 0.697 / 116.982 / 134.218 |
| repair | Logical write completion | 158 | 2.141 / 116.982 / 128.972 |
| repair | set | 158 | 0.934 / 2.478 / 4.037 |
| repair | waitaof | 158 | 0.807 / 116.982 / 128.972 |
| reconciliation | Logical write completion | 545 | 1.598 / 116.982 / 122.831 |
| reconciliation | set | 545 | 0.847 / 1.314 / 2.360 |
| reconciliation | waitaof | 545 | 0.602 / 116.982 / 122.831 |
| post_recovery | Logical write completion | 191 | 1.380 / 2.478 / 3.163 |
| post_recovery | set | 191 | 0.847 / 1.380 / 2.360 |
| post_recovery | waitaof | 191 | 0.520 / 1.598 / 1.942 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 416 | 103.523 |
| fault | 179 | 37.549 |
| repair | 105 | 67.335 |
| reconciliation | 417 | 81.225 |
| post_recovery | 191 | 95.493 |

SET acknowledgements without a subsequent WAITAOF attempt: 0; client timeouts among those: 0. WAITAOF command errors and WAITAOF client timeouts above count attempted WAITAOF client calls only. A short integer pair is separate from a client timeout.

Configured write policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 3495; started 3125; completed 3125; dropped 370. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1544; checked: 1544; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAITAOF-satisfied subset: expected 1308; checked 1308; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAITAOF.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 11 |
| fault | network | 31 |
| fault | timeout | 38 |
| fault | waitaof_short | 55 |
| reconciliation | waitaof_short | 128 |
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
