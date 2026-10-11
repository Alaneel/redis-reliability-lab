# Redis Reliability Lab — dataset-trial

Run: `61f327f21a9b8adb4f662cc5`  
Outcome: **recovered**  
Started: 2026-10-11T19:57:08.245221+08:00  
Finished: 2026-10-11T19:57:35.974242+08:00

## Recovery

| Condition | Observed |
|---|---|
| Surviving topology converged | true |
| All-shard client recovery sustained | true |
| Topology and client recovery jointly sustained | true |
| All acknowledged writes verified | true |
| Target process restarted | false |
| Original replication redundancy restored | true |
| Recovery deadline expired | false |
| Cleanup verified | true |

The replica remains running during whole-node bidirectional network isolation. Connectivity restoration, topology reconciliation and replication redundancy are evaluated separately. Redis replication is asynchronous; acknowledged SETs are not durability guarantees.

## Network restoration

| Condition | Observed |
|---|---|
| Original endpoints and process restored | true |
| Topology and clients jointly reconciled | true |
| Reconciliation deadline expired | false |

Restoration action duration: 1717.017 ms.
Stable reconciliation first observed 1493.010 ms after restoration completed.

| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |
|---|---|---|---|---:|---|---|
| 2026-10-11T19:57:25.615578+08:00 | true | true | false | 81332 | a1b6996ff8522930459464f062290d00341590b3 |  |
| 2026-10-11T19:57:28.39123+08:00 | true | true | false | 81332 | a1b6996ff8522930459464f062290d00341590b3 |  |

## Managed node lifecycle

Target: replica `node-5` (Redis ID `aca134303e542438704f1ef8a7b4b89df4cd0f8b`). Repair: `reconnect`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1717.019 ms.
Stable redundancy first observed 1493.009 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 2725.198 |
| Replication interruption observed | unobserved |
| Target reachable after repair | 7776.624 |
| Membership rejoined | 7776.624 |
| Replica link up | 7776.624 |
| Synchronization in progress | unobserved |
| Synchronization complete | 7776.624 |
| Primary full-sync counter increased | unobserved |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 9253.095 |

Resynchronization evidence: `no_counter_increase_observed`; primary Redis ID `a5a2b706e17893938afc13da87362b6fe45cb3ca`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 1592.127 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 2725.20 | 1133.07 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 2160.70 | 568.57 |
| Surviving views converge | 2160.70 | 568.57 |
| First all-shard SET/GET probe succeeds | 2285.81 | 693.68 |
| Client probes sustain recovery | 3408.75 | 1816.62 |
| Topology and clients sustain recovery | 3408.75 | 1816.62 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 357 | 100.00% | 0 | 169.54 | 0.520 / 1.030 / 3.012 |
| fault | 1027 | 100.00% | 0 | 169.95 | 0.632 / 1.521 / 4.451 |
| repair | 283 | 100.00% | 0 | 164.81 | 0.847 / 38.086 / 83.137 |
| reconciliation | 287 | 100.00% | 0 | 169.67 | 0.573 / 1.192 / 13.020 |
| post_recovery | 183 | 100.00% | 0 | 169.77 | 0.495 / 0.768 / 1.449 |

Acknowledged writes: 1076. Ambiguous SET outcomes: 0.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 165 | 0 |
| fault | 541 | 0 |
| repair | 139 | 0 |
| reconciliation | 149 | 0 |
| post_recovery | 82 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 165 | 0.546 / 1.135 / 2.869 |
| baseline | set | 165 | 0.520 / 1.081 / 2.869 |
| fault | Logical write completion | 541 | 0.632 / 1.598 / 5.410 |
| fault | set | 541 | 0.602 / 1.521 / 5.152 |
| repair | Logical write completion | 139 | 0.934 / 39.990 / 79.178 |
| repair | set | 139 | 0.890 / 39.990 / 79.178 |
| reconciliation | Logical write completion | 149 | 0.573 / 1.314 / 13.020 |
| reconciliation | set | 149 | 0.546 / 1.252 / 13.020 |
| post_recovery | Logical write completion | 82 | 0.520 / 0.768 / 1.759 |
| post_recovery | set | 82 | 0.495 / 0.768 / 1.728 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 165 | 78.358 |
| fault | 541 | 89.525 |
| repair | 139 | 80.948 |
| reconciliation | 149 | 88.086 |
| post_recovery | 82 | 76.072 |

Global pacing: offered 2148; started 2137; completed 2137; dropped 11. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1076; checked: 1076; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Measurement methodology

- Fault selection preserves the initial primary and its replica. The destructive target is the selected replica; only strictly owned local Docker resources are affected.
- Replica faults preserve all original slot owners and do not require promotion. down_duration is a minimum hold from successful fault action completion, bounded by the recovery observation budget. Client service, failure gossip and replication interruption are separately sampled.
- Replica network isolation removes its complete bidirectional Docker bridge connectivity while Redis remains running. It is not a selective replication-only partition. Restoration uses the same owned endpoint and unchanged live process identity.
- Administrative endpoint-disconnection evidence is separate from Redis INFO link or connected-replica indicators. A short disconnection can finish before TCP/link failure detection; absent interruption or synchronization-counter milestones remain unobserved and are not invented from the Docker interface state.
- Crash repair starts the same stopped owned container with its bounded disposable data volume. Redis Cluster identity must match the pre-fault identity. Restart changes the process; it does not simulate power loss, host/storage failure or persistence durability.
- The repair deadline covers the restart/reconnect action and subsequent synchronization. Redis's current membership is inspected without issuing CLUSTER REPLICATE or forcing slot ownership. Replica repair preserves original assignments; former-primary repair requires the original replica to retain promoted ownership and the old primary to rejoin as its replica.
- Fresh stability after repair requires all six original identities and endpoints, healthy equal slot maps, exactly one synchronized replica per primary, and all-shard client probes. Historical service recovery remains separate from final redundancy and configured acknowledgement-policy availability.
- Replication link-up and master_sync_in_progress=0 are sampled indicators, not exact zero lag, bytes preserved, or identification of partial versus full resynchronization. Raw replication INFO is retained for inspection; unobserved milestones remain absent.
- Traffic drains before validating every acknowledged immutable SET, including short/error WAIT/WAITAOF outcomes. Ambiguous SET errors remain separate. Cancellation skips unsupported recovery/data claims and still uses fresh bounded restoration and cleanup contexts, including managed volumes.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `b21b838ba4a0a801142948a2de1aa459d7e4c8bedb000ccd821ee852165c156d`.
