# Redis Reliability Lab — dataset-trial

Run: `42ff17db5678dfbe4b21a5d7`  
Outcome: **recovered**  
Started: 2026-10-11T19:57:36.290068+08:00  
Finished: 2026-10-11T19:58:06.72437+08:00

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

Target: replica `node-5` (Redis ID `35a19f4babbca2395b87729c60bd927ff4a61972`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1682.537 ms.
Stable redundancy first observed 6040.731 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1614.742 |
| Replication interruption observed | 257.732 |
| Target reachable after repair | 5966.817 |
| Membership rejoined | 5966.817 |
| Replica link up | 10542.055 |
| Synchronization in progress | unobserved |
| Synchronization complete | 10542.055 |
| Primary full-sync counter increased | 5966.817 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 11998.527 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `51d3089b3fef8198d88399f390d14bfa31ba0b26`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 251.413 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1614.74 | 1363.33 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 257.73 | 6.32 |
| Surviving views converge | 257.73 | 6.32 |
| First all-shard SET/GET probe succeeds | 379.16 | 127.75 |
| Client probes sustain recovery | 1736.37 | 1484.96 |
| Topology and clients sustain recovery | 1736.37 | 1484.96 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 357 | 100.00% | 0 | 169.94 | 0.847 / 1.449 / 4.907 |
| fault | 726 | 84.71% | 111 | 169.82 | 0.546 / 1.135 / 1.942 |
| repair | 286 | 83.92% | 46 | 169.97 | 0.732 / 1.449 / 2.478 |
| reconciliation | 1062 | 89.27% | 114 | 170.08 | 0.768 / 1.761 / 6.576 |
| post_recovery | 170 | 100.00% | 0 | 169.52 | 0.807 / 1.192 / 1.521 |

Acknowledged writes: 1303. Ambiguous SET outcomes: 0.

Write policy: `wait`. Diagnostic stage requests are excluded from logical workload counts.
Logical write latency includes SET and WAIT on the same physical connection; policy success requires the requested replica-memory acknowledgement count.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 183 | 0 | 183 | 183 | 0 | 0 | 0 |
| fault | 351 | 0 | 351 | 240 | 111 | 0 | 0 |
| repair | 136 | 0 | 136 | 90 | 46 | 0 | 0 |
| reconciliation | 547 | 0 | 547 | 433 | 114 | 0 | 0 |
| post_recovery | 86 | 0 | 86 | 86 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 183 | 0.981 / 1.677 / 4.907 |
| baseline | set | 183 | 0.495 / 1.030 / 4.451 |
| baseline | wait | 183 | 0.449 / 0.664 / 0.981 |
| fault | Logical write completion | 351 | 1.030 / 128.972 / 135.421 |
| fault | set | 351 | 0.495 / 0.732 / 1.598 |
| fault | wait | 351 | 0.495 / 122.831 / 135.421 |
| repair | Logical write completion | 136 | 1.252 / 128.972 / 142.072 |
| repair | set | 136 | 0.602 / 0.847 / 1.677 |
| repair | wait | 136 | 0.546 / 128.972 / 135.421 |
| reconciliation | Logical write completion | 547 | 1.030 / 122.831 / 135.421 |
| reconciliation | set | 547 | 0.520 / 1.081 / 4.037 |
| reconciliation | wait | 547 | 0.495 / 122.831 / 135.421 |
| post_recovery | Logical write completion | 86 | 0.981 / 1.192 / 1.452 |
| post_recovery | set | 86 | 0.495 / 0.664 / 0.924 |
| post_recovery | wait | 86 | 0.449 / 0.546 / 0.588 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 183 | 87.111 |
| fault | 240 | 56.137 |
| repair | 90 | 53.488 |
| reconciliation | 433 | 69.345 |
| post_recovery | 86 | 85.756 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Configured write policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 2601; started 2601; completed 2601; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1303; checked: 1303; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAIT-satisfied subset: expected 1032; checked 1032; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAIT.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | wait_short | 111 |
| reconciliation | wait_short | 114 |
| repair | wait_short | 46 |

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
- Configured acknowledgement policy is a separate current condition, measured through all-three-shard SET plus WAIT/WAITAOF on one physical connection. A stable policy milestone is historical; a later unsuccessful probe resets the current condition.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `b30dd65ed452d45cebd6f40c6a78a936c7b7cb31833b3ff3185b2e1e654a1c2c`.
