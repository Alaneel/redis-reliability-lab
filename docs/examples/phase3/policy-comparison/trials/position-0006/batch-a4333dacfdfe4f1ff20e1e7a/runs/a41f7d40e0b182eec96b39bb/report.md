# Redis Reliability Lab — policy-aof-wait

Run: `a41f7d40e0b182eec96b39bb`  
Outcome: **recovered**  
Started: 2026-10-11T17:14:49.130648+08:00  
Finished: 2026-10-11T17:15:26.954205+08:00

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

Target: primary `node-1` (Redis ID `764f5c311198bb95c7db761c338d81b32e9ccf31`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1589.307 ms.
Stable redundancy first observed 5529.567 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1835.464 |
| Replication interruption observed | 2625.178 |
| Target reachable after repair | 6445.564 |
| Membership rejoined | 6445.564 |
| Replica link up | 10700.484 |
| Synchronization in progress | unobserved |
| Synchronization complete | 10700.484 |
| Primary full-sync counter increased | 6445.564 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 11972.450 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `39b3e7d11f74662b5f1541a52a854bfb80dac99b`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 254.839 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1835.46 | 1580.62 |
| Original replica owns failed primary slots | 2625.18 | 2370.34 |
| Full usable slot coverage | 2625.18 | 2370.34 |
| Surviving views converge | 2625.18 | 2370.34 |
| First all-shard SET/GET probe succeeds | 3531.59 | 3276.75 |
| Client probes sustain recovery | 4717.02 | 4462.18 |
| Topology and clients sustain recovery | 4717.02 | 4462.18 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 804 | 100.00% | 0 | 199.93 | 1.252 / 2.732 / 3.163 |
| fault | 577 | 76.78% | 134 | 118.88 | 0.664 / 3.845 / 200.078 |
| repair | 317 | 82.65% | 55 | 199.45 | 0.546 / 2.602 / 3.321 |
| reconciliation | 1114 | 87.52% | 139 | 200.12 | 0.520 / 2.732 / 4.451 |
| post_recovery | 400 | 100.00% | 0 | 199.99 | 0.520 / 2.360 / 2.869 |

Acknowledged writes: 1583. Ambiguous SET outcomes: 42.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | 764f5c311198bb95c7db761c338d81b32e9ccf31 | 2026-10-11T17:15:01.093897+08:00 | 764f5c311198bb95c7db761c338d81b32e9ccf31 | true |  |
| node-2 | c4c96aedd1e3f1c22bfcc92d49f753bca6b193de | 2026-10-11T17:15:02.072178+08:00 | c4c96aedd1e3f1c22bfcc92d49f753bca6b193de | true |  |
| node-3 | 88b38bb54e622e814c71f3ae48a8a56b4715c863 | 2026-10-11T17:15:03.081147+08:00 | 88b38bb54e622e814c71f3ae48a8a56b4715c863 | true |  |
| node-4 | d5c3421f9d62c73af7c1aebf2b31ae3fe5ec70f5 | 2026-10-11T17:15:04.060847+08:00 | d5c3421f9d62c73af7c1aebf2b31ae3fe5ec70f5 | true |  |
| node-5 | 39b3e7d11f74662b5f1541a52a854bfb80dac99b | 2026-10-11T17:15:05.070612+08:00 | 39b3e7d11f74662b5f1541a52a854bfb80dac99b | true |  |
| node-6 | 6d8c8fc6d99949ee88851c0a72b1151201cf486b | 2026-10-11T17:15:06.04933+08:00 | 6d8c8fc6d99949ee88851c0a72b1151201cf486b | true |  |

Before-baseline `wait` policy probe at 2026-10-11T17:15:06.059618+08:00; all shards successful: true.
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
| fault | 229 | 42 | 229 | 174 | 55 | 0 | 0 |
| repair | 165 | 0 | 165 | 110 | 55 | 0 | 0 |
| reconciliation | 585 | 0 | 585 | 446 | 139 | 0 | 0 |
| post_recovery | 187 | 0 | 187 | 187 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 417 | 1.677 / 3.012 / 3.163 |
| baseline | set | 417 | 0.981 / 1.598 / 2.478 |
| baseline | wait | 417 | 0.632 / 1.761 / 2.039 |
| fault | Logical write completion | 271 | 1.677 / 116.982 / 142.192 |
| fault | set | 271 | 0.934 / 2.141 / 5.152 |
| fault | wait | 229 | 0.768 / 116.982 / 142.192 |
| repair | Logical write completion | 165 | 2.039 / 122.831 / 135.104 |
| repair | set | 165 | 0.981 / 1.598 / 2.687 |
| repair | wait | 165 | 0.768 / 122.831 / 128.972 |
| reconciliation | Logical write completion | 585 | 1.677 / 122.831 / 128.972 |
| reconciliation | set | 585 | 0.847 / 1.677 / 2.732 |
| reconciliation | wait | 585 | 0.632 / 116.982 / 128.972 |
| post_recovery | Logical write completion | 187 | 1.449 / 2.732 / 3.012 |
| post_recovery | set | 187 | 0.890 / 1.598 / 2.478 |
| post_recovery | wait | 187 | 0.546 / 1.761 / 2.039 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 417 | 103.694 |
| fault | 174 | 35.850 |
| repair | 110 | 69.210 |
| reconciliation | 446 | 80.120 |
| post_recovery | 187 | 93.495 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Configured write policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 3606; started 3212; completed 3212; dropped 394. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1583; checked: 1583; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAIT-satisfied subset: expected 1334; checked 1334; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAIT.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 12 |
| fault | network | 30 |
| fault | timeout | 37 |
| fault | wait_short | 55 |
| reconciliation | wait_short | 139 |
| repair | wait_short | 55 |

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
