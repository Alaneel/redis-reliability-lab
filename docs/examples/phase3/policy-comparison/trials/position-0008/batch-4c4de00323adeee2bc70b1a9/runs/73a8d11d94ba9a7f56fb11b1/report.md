# Redis Reliability Lab — policy-aof-wait

Run: `73a8d11d94ba9a7f56fb11b1`  
Outcome: **recovered**  
Started: 2026-10-11T17:16:05.416608+08:00  
Finished: 2026-10-11T17:16:42.639174+08:00

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

Target: primary `node-1` (Redis ID `db4a927f3e910db533a248bddaf2ad81a537fa37`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1537.263 ms.
Stable redundancy first observed 4644.760 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1797.670 |
| Replication interruption observed | 2572.051 |
| Target reachable after repair | 5882.436 |
| Membership rejoined | 5882.436 |
| Replica link up | 9247.574 |
| Synchronization in progress | unobserved |
| Synchronization complete | 9247.574 |
| Primary full-sync counter increased | 5882.436 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 10523.834 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `a032ad225e3c64165b36a5842e2d72548fce4c5d`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 235.555 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1797.67 | 1562.12 |
| Original replica owns failed primary slots | 2572.05 | 2336.50 |
| Full usable slot coverage | 2572.05 | 2336.50 |
| Surviving views converge | 2572.05 | 2336.50 |
| First all-shard SET/GET probe succeeds | 3035.83 | 2800.27 |
| Client probes sustain recovery | 4230.16 | 3994.61 |
| Topology and clients sustain recovery | 4230.16 | 3994.61 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 803 | 100.00% | 0 | 199.77 | 1.252 / 2.732 / 3.321 |
| fault | 494 | 76.52% | 116 | 113.78 | 0.768 / 3.163 / 436.746 |
| repair | 307 | 84.36% | 48 | 199.70 | 0.449 / 2.360 / 3.321 |
| reconciliation | 936 | 89.10% | 102 | 199.90 | 0.520 / 2.478 / 4.239 |
| post_recovery | 401 | 100.00% | 0 | 200.46 | 1.192 / 3.012 / 6.263 |

Acknowledged writes: 1462. Ambiguous SET outcomes: 35.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | db4a927f3e910db533a248bddaf2ad81a537fa37 | 2026-10-11T17:16:18.095894+08:00 | db4a927f3e910db533a248bddaf2ad81a537fa37 | true |  |
| node-2 | 5fff76d8a527bcf725f6cd93072d4b79421ff848 | 2026-10-11T17:16:19.109642+08:00 | 5fff76d8a527bcf725f6cd93072d4b79421ff848 | true |  |
| node-3 | 9ca12628df178b78d6e0a950985ac4eeff940b16 | 2026-10-11T17:16:20.147143+08:00 | 9ca12628df178b78d6e0a950985ac4eeff940b16 | true |  |
| node-4 | 639eb01ed32b8feae9e4a08eab376aa37fd4f82e | 2026-10-11T17:16:21.154061+08:00 | 639eb01ed32b8feae9e4a08eab376aa37fd4f82e | true |  |
| node-5 | a032ad225e3c64165b36a5842e2d72548fce4c5d | 2026-10-11T17:16:22.146533+08:00 | a032ad225e3c64165b36a5842e2d72548fce4c5d | true |  |
| node-6 | 0c094e0c5a57e6a1f50e2059feeb59939bb1c409 | 2026-10-11T17:16:23.133608+08:00 | 0c094e0c5a57e6a1f50e2059feeb59939bb1c409 | true |  |

Before-baseline `wait` policy probe at 2026-10-11T17:16:23.143778+08:00; all shards successful: true.
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
| baseline | 417 | 0 | 417 | 417 | 0 | 0 | 0 |
| fault | 209 | 35 | 209 | 165 | 44 | 0 | 0 |
| repair | 146 | 0 | 146 | 98 | 48 | 0 | 0 |
| reconciliation | 479 | 0 | 479 | 377 | 102 | 0 | 0 |
| post_recovery | 211 | 0 | 211 | 211 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 417 | 1.677 / 3.012 / 3.662 |
| baseline | set | 417 | 0.981 / 1.449 / 2.602 |
| baseline | wait | 417 | 0.664 / 1.942 / 2.141 |
| fault | Logical write completion | 244 | 1.521 / 116.982 / 125.721 |
| fault | set | 244 | 0.847 / 1.761 / 3.662 |
| fault | wait | 209 | 0.602 / 116.982 / 122.831 |
| repair | Logical write completion | 146 | 1.942 / 122.831 / 135.421 |
| repair | set | 146 | 0.934 / 1.677 / 2.732 |
| repair | wait | 146 | 0.732 / 122.831 / 134.659 |
| reconciliation | Logical write completion | 479 | 1.521 / 116.982 / 128.972 |
| reconciliation | set | 479 | 0.847 / 1.380 / 3.321 |
| reconciliation | wait | 479 | 0.602 / 116.982 / 128.972 |
| post_recovery | Logical write completion | 211 | 1.449 / 3.845 / 7.250 |
| post_recovery | set | 211 | 0.890 / 2.602 / 5.410 |
| post_recovery | wait | 211 | 0.546 / 1.942 / 3.487 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 417 | 103.744 |
| fault | 165 | 38.003 |
| repair | 98 | 63.748 |
| reconciliation | 377 | 80.516 |
| post_recovery | 211 | 105.480 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Configured write policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 3316; started 2941; completed 2941; dropped 375. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1462; checked: 1462; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAIT-satisfied subset: expected 1268; checked 1268; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAIT.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 4 |
| fault | network | 31 |
| fault | timeout | 37 |
| fault | wait_short | 44 |
| reconciliation | wait_short | 102 |
| repair | wait_short | 48 |

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
