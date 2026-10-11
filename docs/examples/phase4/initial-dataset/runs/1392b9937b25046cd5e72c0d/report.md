# Redis Reliability Lab — dataset-trial

Run: `1392b9937b25046cd5e72c0d`  
Outcome: **recovered**  
Started: 2026-10-11T19:47:59.856173+08:00  
Finished: 2026-10-11T19:48:39.342642+08:00

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

Target: replica `node-5` (Redis ID `b1ce5ac1d200ef403e8478480fb408d408312f0e`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1591.753 ms.
Stable redundancy first observed 5415.539 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1400.655 |
| Replication interruption observed | 265.917 |
| Target reachable after repair | 6359.325 |
| Membership rejoined | 6359.325 |
| Replica link up | 10320.107 |
| Synchronization in progress | unobserved |
| Synchronization complete | 10320.107 |
| Primary full-sync counter increased | 6359.325 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 11764.855 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `8ecd39fb257cf10e257b2391bc34fe6ab5ee70f5`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 259.031 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1400.65 | 1141.62 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 265.92 | 6.89 |
| Surviving views converge | 265.92 | 6.89 |
| First all-shard SET/GET probe succeeds | 392.61 | 133.58 |
| Client probes sustain recovery | 1526.66 | 1267.63 |
| Topology and clients sustain recovery | 1526.66 | 1267.63 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 357 | 100.00% | 0 | 169.89 | 0.495 / 0.732 / 1.521 |
| fault | 808 | 100.00% | 0 | 169.84 | 0.520 / 0.732 / 0.981 |
| repair | 271 | 100.00% | 0 | 170.23 | 0.602 / 1.030 / 1.380 |
| reconciliation | 955 | 100.00% | 0 | 170.05 | 0.495 / 0.732 / 1.521 |
| post_recovery | 170 | 100.00% | 0 | 169.96 | 0.495 / 0.664 / 0.764 |

Acknowledged writes: 1309. Ambiguous SET outcomes: 0.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 185 | 0 |
| fault | 399 | 0 |
| repair | 129 | 0 |
| reconciliation | 510 | 0 |
| post_recovery | 86 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 185 | 0.520 / 0.697 / 1.677 |
| baseline | set | 185 | 0.495 / 0.697 / 1.677 |
| fault | Logical write completion | 399 | 0.546 / 0.768 / 0.981 |
| fault | set | 399 | 0.520 / 0.732 / 0.981 |
| repair | Logical write completion | 129 | 0.602 / 1.081 / 1.449 |
| repair | set | 129 | 0.573 / 0.934 / 1.314 |
| reconciliation | Logical write completion | 510 | 0.495 / 0.732 / 1.449 |
| reconciliation | set | 510 | 0.472 / 0.732 / 1.380 |
| post_recovery | Logical write completion | 86 | 0.520 / 0.664 / 0.743 |
| post_recovery | set | 86 | 0.495 / 0.632 / 0.720 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 185 | 88.037 |
| fault | 399 | 83.869 |
| repair | 129 | 81.031 |
| reconciliation | 510 | 90.815 |
| post_recovery | 86 | 85.982 |

Global pacing: offered 2561; started 2561; completed 2561; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1309; checked: 1309; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `f2bb2e5a3d24f1e2bbafbebb0ab2384bfd332a1340507a6e35a93860c8a81b13`.
