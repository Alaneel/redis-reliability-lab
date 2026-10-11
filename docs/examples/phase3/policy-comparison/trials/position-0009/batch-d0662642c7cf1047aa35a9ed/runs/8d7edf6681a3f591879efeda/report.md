# Redis Reliability Lab — policy-aof-set

Run: `8d7edf6681a3f591879efeda`  
Outcome: **recovered**  
Started: 2026-10-11T17:16:42.829144+08:00  
Finished: 2026-10-11T17:17:22.210203+08:00

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

Target: primary `node-1` (Redis ID `353be4b2ffc7569fb1e31e9ff106e247b8fec4c8`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1510.389 ms.
Stable redundancy first observed 5976.766 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1804.405 |
| Replication interruption observed | 2588.466 |
| Target reachable after repair | 6039.703 |
| Membership rejoined | 6039.703 |
| Replica link up | 10733.578 |
| Synchronization in progress | unobserved |
| Synchronization complete | 10733.578 |
| Primary full-sync counter increased | 6039.703 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 12013.147 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `b1fbc4156d950a774b1df4d5689d3c3ae2068f73`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 236.912 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1804.40 | 1567.49 |
| Original replica owns failed primary slots | 2588.47 | 2351.55 |
| Full usable slot coverage | 2588.47 | 2351.55 |
| Surviving views converge | 2588.47 | 2351.55 |
| First all-shard SET/GET probe succeeds | 3398.08 | 3161.17 |
| Client probes sustain recovery | 4522.18 | 4285.27 |
| Topology and clients sustain recovery | 4522.18 | 4285.27 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 804 | 100.00% | 0 | 199.95 | 0.807 / 1.380 / 2.360 |
| fault | 543 | 85.64% | 78 | 119.97 | 0.697 / 2.360 / 325.906 |
| repair | 302 | 100.00% | 0 | 199.94 | 0.602 / 1.314 / 2.602 |
| reconciliation | 1202 | 100.00% | 0 | 199.96 | 0.632 / 1.192 / 1.677 |
| post_recovery | 400 | 100.00% | 0 | 199.94 | 0.697 / 2.039 / 2.478 |

Acknowledged writes: 1607. Ambiguous SET outcomes: 44.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | 353be4b2ffc7569fb1e31e9ff106e247b8fec4c8 | 2026-10-11T17:16:56.14886+08:00 | 353be4b2ffc7569fb1e31e9ff106e247b8fec4c8 | true |  |
| node-2 | 749146776e95fd6196d72cd7cc899013dae39693 | 2026-10-11T17:16:57.128048+08:00 | 749146776e95fd6196d72cd7cc899013dae39693 | true |  |
| node-3 | 577c44cb0f35bd911e3203a327fb3a54d8c208ca | 2026-10-11T17:16:58.190907+08:00 | 577c44cb0f35bd911e3203a327fb3a54d8c208ca | true |  |
| node-4 | b12b462d8bdb97fda0b7ad3d9b2230746afb75fc | 2026-10-11T17:16:59.204591+08:00 | b12b462d8bdb97fda0b7ad3d9b2230746afb75fc | true |  |
| node-5 | b1fbc4156d950a774b1df4d5689d3c3ae2068f73 | 2026-10-11T17:17:00.191377+08:00 | b1fbc4156d950a774b1df4d5689d3c3ae2068f73 | true |  |
| node-6 | 62e54ad9bea8615974929ffe5175cced3878ab45 | 2026-10-11T17:17:01.206357+08:00 | 62e54ad9bea8615974929ffe5175cced3878ab45 | true |  |

Before-baseline `set` policy probe at 2026-10-11T17:17:01.221043+08:00; all shards successful: true.

| Original shard index | SET acknowledged |
|---|---|
| 0 | true |
| 1 | true |
| 2 | true |

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 417 | 0 |
| fault | 217 | 44 |
| repair | 144 | 0 |
| reconciliation | 634 | 0 |
| post_recovery | 195 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 417 | 1.030 / 1.677 / 2.360 |
| baseline | set | 417 | 1.030 / 1.677 / 2.360 |
| fault | Logical write completion | 261 | 0.934 / 1.761 / 3.662 |
| fault | set | 261 | 0.934 / 1.761 / 3.662 |
| repair | Logical write completion | 144 | 0.934 / 1.761 / 2.602 |
| repair | set | 144 | 0.934 / 1.761 / 2.602 |
| reconciliation | Logical write completion | 634 | 0.890 / 1.380 / 2.039 |
| reconciliation | set | 634 | 0.890 / 1.380 / 2.039 |
| post_recovery | Logical write completion | 195 | 0.934 / 2.360 / 2.869 |
| post_recovery | set | 195 | 0.934 / 2.360 / 2.869 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 417 | 103.704 |
| fault | 217 | 47.945 |
| repair | 144 | 95.337 |
| reconciliation | 634 | 105.467 |
| post_recovery | 195 | 97.473 |

Global pacing: offered 3613; started 3251; completed 3251; dropped 362. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1607; checked: 1607; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 14 |
| fault | network | 30 |
| fault | timeout | 34 |

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
