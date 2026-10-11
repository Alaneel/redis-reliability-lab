# Redis Reliability Lab — dataset-trial

Run: `497e76933ed98eabedcf5f9a`  
Outcome: **recovered**  
Started: 2026-10-11T19:43:22.839446+08:00  
Finished: 2026-10-11T19:43:51.950753+08:00

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

Target: replica `node-5` (Redis ID `d6e93eb87ac49ee8ea1411e61f947d3d0b2183bc`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1600.701 ms.
Stable redundancy first observed 5044.457 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1603.513 |
| Replication interruption observed | 257.283 |
| Target reachable after repair | 5003.153 |
| Membership rejoined | 5003.153 |
| Replica link up | 8807.322 |
| Synchronization in progress | unobserved |
| Synchronization complete | 8807.322 |
| Primary full-sync counter increased | 5003.153 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 10037.948 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `830fc750bc41e8874d5ee5f098aeafd420fb11b5`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 251.291 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1603.51 | 1352.22 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 257.28 | 5.99 |
| Surviving views converge | 257.28 | 5.99 |
| First all-shard SET/GET probe succeeds | 379.17 | 127.88 |
| Client probes sustain recovery | 1724.05 | 1472.76 |
| Topology and clients sustain recovery | 1724.05 | 1472.76 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 210 | 100.00% | 0 | 99.56 | 0.697 / 1.030 / 1.942 |
| fault | 340 | 100.00% | 0 | 100.22 | 0.546 / 0.768 / 0.981 |
| repair | 160 | 100.00% | 0 | 99.95 | 0.632 / 1.081 / 1.380 |
| reconciliation | 524 | 100.00% | 0 | 99.94 | 0.546 / 0.732 / 0.890 |
| post_recovery | 100 | 100.00% | 0 | 99.97 | 0.573 / 0.732 / 0.890 |

Acknowledged writes: 683. Ambiguous SET outcomes: 0.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 109 | 0 |
| fault | 180 | 0 |
| repair | 76 | 0 |
| reconciliation | 270 | 0 |
| post_recovery | 48 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 109 | 0.697 / 1.030 / 1.942 |
| baseline | set | 109 | 0.664 / 0.981 / 1.849 |
| fault | Logical write completion | 180 | 0.546 / 0.768 / 0.981 |
| fault | set | 180 | 0.520 / 0.732 / 0.934 |
| repair | Logical write completion | 76 | 0.632 / 0.890 / 1.306 |
| repair | set | 76 | 0.602 / 0.847 / 1.232 |
| reconciliation | Logical write completion | 270 | 0.546 / 0.768 / 0.934 |
| reconciliation | set | 270 | 0.546 / 0.732 / 0.890 |
| post_recovery | Logical write completion | 48 | 0.602 / 0.732 / 0.823 |
| post_recovery | set | 48 | 0.573 / 0.732 / 0.803 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 109 | 51.678 |
| fault | 180 | 53.055 |
| repair | 76 | 47.474 |
| reconciliation | 270 | 51.495 |
| post_recovery | 48 | 47.988 |

Global pacing: offered 1334; started 1334; completed 1334; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 683; checked: 683; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `c97545ea94937baecf80f417ea603d0bacabb93a7a0f6874e1d77cdecad7df4e`.
