# Redis Reliability Lab — dataset-trial

Run: `6bd90861768a51c6096c09a4`  
Outcome: **recovered**  
Started: 2026-10-11T19:44:27.406872+08:00  
Finished: 2026-10-11T19:44:55.065157+08:00

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

Restoration action duration: 1304.602 ms.
Stable reconciliation first observed 1240.416 ms after restoration completed.

| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |
|---|---|---|---|---:|---|---|
| 2026-10-11T19:44:46.057286+08:00 | true | true | false | 10333 | 9afbe1f637f60ed073dea91474d36af8cf46e893 |  |
| 2026-10-11T19:44:48.835886+08:00 | true | true | false | 10333 | 9afbe1f637f60ed073dea91474d36af8cf46e893 |  |

## Managed node lifecycle

Target: replica `node-5` (Redis ID `f8bc258d1ed314cf06fa040b2ad7a15451279168`). Repair: `reconnect`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1304.611 ms.
Stable redundancy first observed 1240.408 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 2465.663 |
| Replication interruption observed | unobserved |
| Target reachable after repair | 6548.694 |
| Membership rejoined | 6548.694 |
| Replica link up | 6548.694 |
| Synchronization in progress | unobserved |
| Synchronization complete | 6548.694 |
| Primary full-sync counter increased | unobserved |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 7779.144 |

Resynchronization evidence: `no_counter_increase_observed`; primary Redis ID `9713a715e675cdfdb24232a4101a018743bfcd17`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 1466.246 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 2465.66 | 999.42 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 2011.27 | 545.03 |
| Surviving views converge | 2011.27 | 545.03 |
| First all-shard SET/GET probe succeeds | 2135.41 | 669.16 |
| Client probes sustain recovery | 3492.72 | 2026.47 |
| Topology and clients sustain recovery | 3492.72 | 2026.47 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 210 | 100.00% | 0 | 99.87 | 0.732 / 1.192 / 1.849 |
| fault | 523 | 100.00% | 0 | 99.92 | 0.664 / 1.030 / 1.598 |
| repair | 131 | 100.00% | 0 | 100.40 | 0.732 / 1.030 / 1.192 |
| reconciliation | 144 | 100.00% | 0 | 99.90 | 0.495 / 0.732 / 1.030 |
| post_recovery | 100 | 100.00% | 0 | 99.90 | 0.546 / 0.732 / 0.934 |

Acknowledged writes: 534. Ambiguous SET outcomes: 0.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 108 | 0 |
| fault | 242 | 0 |
| repair | 60 | 0 |
| reconciliation | 75 | 0 |
| post_recovery | 49 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 108 | 0.768 / 1.192 / 1.521 |
| baseline | set | 108 | 0.732 / 1.135 / 1.521 |
| fault | Logical write completion | 242 | 0.697 / 1.030 / 1.942 |
| fault | set | 242 | 0.664 / 0.934 / 1.521 |
| repair | Logical write completion | 60 | 0.732 / 1.030 / 1.574 |
| repair | set | 60 | 0.697 / 0.981 / 1.494 |
| reconciliation | Logical write completion | 75 | 0.495 / 0.768 / 1.018 |
| reconciliation | set | 75 | 0.495 / 0.732 / 0.987 |
| post_recovery | Logical write completion | 49 | 0.546 / 0.768 / 2.530 |
| post_recovery | set | 49 | 0.520 / 0.732 / 2.471 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 108 | 51.359 |
| fault | 242 | 46.235 |
| repair | 60 | 45.986 |
| reconciliation | 75 | 52.032 |
| post_recovery | 49 | 48.953 |

Global pacing: offered 1108; started 1108; completed 1108; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 534; checked: 534; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `1d4e377103ce227e5b962d00c98e92b3bb149c60180bac3380b2d478a452c2f8`.
