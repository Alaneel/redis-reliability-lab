# Redis Reliability Lab — dataset-trial

Run: `d7245c097de839daf38fe9c6`  
Outcome: **recovered**  
Started: 2026-10-11T19:49:18.056977+08:00  
Finished: 2026-10-11T19:49:44.10939+08:00

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

Restoration action duration: 1309.058 ms.
Stable reconciliation first observed 1452.552 ms after restoration completed.

| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |
|---|---|---|---|---:|---|---|
| 2026-10-11T19:49:34.327597+08:00 | true | true | false | 35433 | a44adc63ab8941b959f2d3b975bf384a92514897 |  |
| 2026-10-11T19:49:37.085729+08:00 | true | true | false | 35433 | a44adc63ab8941b959f2d3b975bf384a92514897 |  |

## Managed node lifecycle

Target: replica `node-5` (Redis ID `40b3dcb5606712ef69ea42a1ffa8ce5b97a57401`). Repair: `reconnect`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1309.061 ms.
Stable redundancy first observed 1452.549 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1873.297 |
| Replication interruption observed | unobserved |
| Target reachable after repair | 7062.954 |
| Membership rejoined | 7062.954 |
| Replica link up | 7062.954 |
| Synchronization in progress | unobserved |
| Synchronization complete | 7062.954 |
| Primary full-sync counter increased | unobserved |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 8505.364 |

Resynchronization evidence: `no_counter_increase_observed`; primary Redis ID `35e676a91f84ce9f165f0db4cab6192b25c18eba`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 1348.029 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1873.30 | 525.27 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 1873.30 | 525.27 |
| Surviving views converge | 1873.30 | 525.27 |
| First all-shard SET/GET probe succeeds | 1993.16 | 645.13 |
| Client probes sustain recovery | 3115.99 | 1767.96 |
| Topology and clients sustain recovery | 3115.99 | 1767.96 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 356 | 100.00% | 0 | 169.69 | 0.573 / 0.981 / 2.602 |
| fault | 977 | 100.00% | 0 | 170.10 | 0.520 / 0.732 / 1.135 |
| repair | 222 | 100.00% | 0 | 169.57 | 0.632 / 0.890 / 1.192 |
| reconciliation | 281 | 100.00% | 0 | 170.31 | 0.495 / 0.632 / 0.981 |
| post_recovery | 183 | 100.00% | 0 | 169.71 | 0.449 / 0.573 / 0.697 |

Acknowledged writes: 969. Ambiguous SET outcomes: 0.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 163 | 0 |
| fault | 474 | 0 |
| repair | 101 | 0 |
| reconciliation | 132 | 0 |
| post_recovery | 99 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 163 | 0.573 / 0.934 / 1.942 |
| baseline | set | 163 | 0.546 / 0.934 / 1.849 |
| fault | Logical write completion | 474 | 0.546 / 0.807 / 1.252 |
| fault | set | 474 | 0.520 / 0.732 / 1.192 |
| repair | Logical write completion | 101 | 0.632 / 0.890 / 1.192 |
| repair | set | 101 | 0.573 / 0.807 / 0.917 |
| reconciliation | Logical write completion | 132 | 0.520 / 0.632 / 0.981 |
| reconciliation | set | 132 | 0.495 / 0.602 / 0.981 |
| post_recovery | Logical write completion | 99 | 0.472 / 0.602 / 1.304 |
| post_recovery | set | 99 | 0.449 / 0.573 / 1.280 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 163 | 77.695 |
| fault | 474 | 82.525 |
| repair | 101 | 77.147 |
| reconciliation | 132 | 80.003 |
| post_recovery | 99 | 91.810 |

Global pacing: offered 2019; started 2019; completed 2019; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 969; checked: 969; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `1c7699bca874cae6ac1747c75fafcd4fbeddd578f8a8b5a9f06b7b85ff634cde`.
