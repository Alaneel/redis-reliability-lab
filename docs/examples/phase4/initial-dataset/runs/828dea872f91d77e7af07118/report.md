# Redis Reliability Lab — dataset-trial

Run: `828dea872f91d77e7af07118`  
Outcome: **recovered**  
Started: 2026-10-11T19:44:55.391667+08:00  
Finished: 2026-10-11T19:45:28.383406+08:00

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

Target: replica `node-5` (Redis ID `dbd3d150660190e74211b958dd269e40041346f6`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1674.518 ms.
Stable redundancy first observed 5858.393 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1983.420 |
| Replication interruption observed | 278.034 |
| Target reachable after repair | 5369.509 |
| Membership rejoined | 5369.509 |
| Replica link up | 9996.956 |
| Synchronization in progress | unobserved |
| Synchronization complete | 9996.956 |
| Primary full-sync counter increased | 5369.509 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 11216.661 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `89ec9ae0b6fb81d8b1037578bc2aca5d62bb0bfc`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 272.462 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1983.42 | 1710.96 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 278.03 | 5.57 |
| Surviving views converge | 278.03 | 5.57 |
| First all-shard SET/GET probe succeeds | 406.74 | 134.28 |
| Client probes sustain recovery | 1535.97 | 1263.50 |
| Topology and clients sustain recovery | 1535.97 | 1263.50 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 210 | 100.00% | 0 | 99.84 | 0.934 / 2.141 / 13.020 |
| fault | 368 | 81.79% | 67 | 99.90 | 0.807 / 1.380 / 1.849 |
| repair | 168 | 82.74% | 29 | 100.32 | 0.847 / 1.677 / 1.849 |
| reconciliation | 606 | 88.28% | 71 | 100.02 | 0.664 / 1.252 / 1.380 |
| post_recovery | 100 | 100.00% | 0 | 99.91 | 0.573 / 1.135 / 1.142 |

Acknowledged writes: 730. Ambiguous SET outcomes: 0.

Write policy: `wait`. Diagnostic stage requests are excluded from logical workload counts.
Logical write latency includes SET and WAIT on the same physical connection; policy success requires the requested replica-memory acknowledgement count.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 101 | 0 | 101 | 101 | 0 | 0 | 0 |
| fault | 207 | 0 | 207 | 140 | 67 | 0 | 0 |
| repair | 87 | 0 | 87 | 58 | 29 | 0 | 0 |
| reconciliation | 289 | 0 | 289 | 218 | 71 | 0 | 0 |
| post_recovery | 46 | 0 | 46 | 46 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 101 | 1.135 / 1.849 / 18.320 |
| baseline | set | 101 | 0.602 / 1.252 / 4.037 |
| baseline | wait | 101 | 0.495 / 0.697 / 1.252 |
| fault | Logical write completion | 207 | 1.192 / 135.421 / 141.341 |
| fault | set | 207 | 0.573 / 0.807 / 1.449 |
| fault | wait | 207 | 0.520 / 135.421 / 140.797 |
| repair | Logical write completion | 87 | 1.449 / 135.421 / 140.682 |
| repair | set | 87 | 0.697 / 0.934 / 1.450 |
| repair | wait | 87 | 0.602 / 135.421 / 139.789 |
| reconciliation | Logical write completion | 289 | 1.135 / 135.421 / 142.192 |
| reconciliation | set | 289 | 0.573 / 0.732 / 0.847 |
| reconciliation | wait | 289 | 0.495 / 135.421 / 142.192 |
| post_recovery | Logical write completion | 46 | 0.981 / 1.135 / 1.142 |
| post_recovery | set | 46 | 0.520 / 0.632 / 0.657 |
| post_recovery | wait | 46 | 0.428 / 0.520 / 0.546 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 101 | 48.020 |
| fault | 140 | 38.005 |
| repair | 58 | 34.635 |
| reconciliation | 218 | 35.979 |
| post_recovery | 46 | 45.958 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Configured write policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 1452; started 1452; completed 1452; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 730; checked: 730; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAIT-satisfied subset: expected 563; checked 563; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAIT.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | wait_short | 67 |
| reconciliation | wait_short | 71 |
| repair | wait_short | 29 |

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `17ada32017dbe51dbd45e81da45e7fa375444c4333ce4d1a6451c598176d6566`.
