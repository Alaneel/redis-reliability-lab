# Redis Reliability Lab — dataset-trial

Run: `0f7c4cf14829038fbb88fd6d`  
Outcome: **recovered**  
Started: 2026-10-11T19:49:44.411409+08:00  
Finished: 2026-10-11T19:50:14.570612+08:00

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

Target: replica `node-5` (Redis ID `f8ed325c445e75ecaeb2718880af0a34e6dba328`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1595.719 ms.
Stable redundancy first observed 5352.409 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1611.292 |
| Replication interruption observed | 268.545 |
| Target reachable after repair | 5895.210 |
| Membership rejoined | 5895.210 |
| Replica link up | 9768.735 |
| Synchronization in progress | unobserved |
| Synchronization complete | 9768.735 |
| Primary full-sync counter increased | 5895.210 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 11237.406 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `7d91ba3956ae19f4ee2145add59446980c729dce`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 258.739 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1611.29 | 1352.55 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 268.55 | 9.81 |
| Surviving views converge | 268.55 | 9.81 |
| First all-shard SET/GET probe succeeds | 390.73 | 131.99 |
| Client probes sustain recovery | 1732.61 | 1473.87 |
| Topology and clients sustain recovery | 1732.61 | 1473.87 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 357 | 100.00% | 0 | 169.51 | 0.890 / 1.314 / 2.732 |
| fault | 730 | 83.84% | 118 | 170.19 | 0.573 / 1.135 / 1.314 |
| repair | 271 | 82.66% | 47 | 169.82 | 0.697 / 1.449 / 1.761 |
| reconciliation | 944 | 90.47% | 90 | 169.96 | 0.573 / 1.135 / 2.360 |
| post_recovery | 170 | 100.00% | 0 | 169.85 | 0.602 / 1.135 / 2.478 |

Acknowledged writes: 1223. Ambiguous SET outcomes: 0.

Write policy: `wait`. Diagnostic stage requests are excluded from logical workload counts.
Logical write latency includes SET and WAIT on the same physical connection; policy success requires the requested replica-memory acknowledgement count.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 184 | 0 | 184 | 184 | 0 | 0 | 0 |
| fault | 367 | 0 | 367 | 249 | 118 | 0 | 0 |
| repair | 141 | 0 | 141 | 94 | 47 | 0 | 0 |
| reconciliation | 452 | 0 | 452 | 362 | 90 | 0 | 0 |
| post_recovery | 79 | 0 | 79 | 79 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 184 | 1.030 / 1.314 / 4.239 |
| baseline | set | 184 | 0.546 / 0.697 / 1.449 |
| baseline | wait | 184 | 0.472 / 0.602 / 1.192 |
| fault | Logical write completion | 367 | 1.030 / 122.831 / 135.421 |
| fault | set | 367 | 0.495 / 0.664 / 1.849 |
| fault | wait | 367 | 0.472 / 122.831 / 135.421 |
| repair | Logical write completion | 141 | 1.192 / 122.831 / 127.288 |
| repair | set | 141 | 0.546 / 0.847 / 1.135 |
| repair | wait | 141 | 0.520 / 122.831 / 126.597 |
| reconciliation | Logical write completion | 452 | 0.981 / 122.831 / 135.421 |
| reconciliation | set | 452 | 0.495 / 0.664 / 1.380 |
| reconciliation | wait | 452 | 0.449 / 122.831 / 135.421 |
| post_recovery | Logical write completion | 79 | 0.934 / 1.252 / 2.828 |
| post_recovery | set | 79 | 0.495 / 0.697 / 2.245 |
| post_recovery | wait | 79 | 0.428 / 0.573 / 0.680 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 184 | 87.366 |
| fault | 249 | 58.053 |
| repair | 94 | 58.903 |
| reconciliation | 362 | 65.174 |
| post_recovery | 79 | 78.930 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Configured write policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 2472; started 2472; completed 2472; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1223; checked: 1223; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAIT-satisfied subset: expected 968; checked 968; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAIT.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | wait_short | 118 |
| reconciliation | wait_short | 90 |
| repair | wait_short | 47 |

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `39ea4e364a486a2ca98748999dc04aa492aee68e812a30784da39e615ac3746d`.
