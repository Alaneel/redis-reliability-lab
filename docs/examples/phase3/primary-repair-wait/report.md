# Redis Reliability Lab — primary-repair-wait

Run: `7ce5fd7f49d14f7c38460823`  
Outcome: **recovered**  
Started: 2026-10-11T16:10:19.723237+08:00  
Finished: 2026-10-11T16:10:50.340974+08:00

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

Target: primary `node-1` (Redis ID `9c20c23f87bd68faa0616d849cb05a67b692a929`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1553.768 ms.
Stable redundancy first observed 5913.530 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1817.312 |
| Replication interruption observed | 3472.588 |
| Target reachable after repair | 6330.557 |
| Membership rejoined | 6330.557 |
| Replica link up | 10970.081 |
| Synchronization in progress | unobserved |
| Synchronization complete | 10970.081 |
| Primary full-sync counter increased | 6330.557 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 12241.066 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `3443ad37b00fb35be10ba504aaf25e369e4b5d9f`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 260.599 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1817.31 | 1556.71 |
| Original replica owns failed primary slots | 3472.59 | 3211.99 |
| Full usable slot coverage | 3472.59 | 3211.99 |
| Surviving views converge | 3472.59 | 3211.99 |
| First all-shard SET/GET probe succeeds | 3493.34 | 3232.75 |
| Client probes sustain recovery | 4664.92 | 4404.32 |
| Topology and clients sustain recovery | 4664.92 | 4404.32 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 394 | 100.00% | 0 | 195.39 | 0.732 / 2.039 / 5.964 |
| fault | 626 | 76.36% | 148 | 131.13 | 0.602 / 1.252 / 377.277 |
| repair | 310 | 85.16% | 46 | 199.51 | 0.632 / 2.602 / 5.152 |
| reconciliation | 1191 | 87.41% | 150 | 200.11 | 0.546 / 1.192 / 3.662 |
| post_recovery | 200 | 100.00% | 0 | 199.96 | 0.732 / 1.677 / 3.012 |

Acknowledged writes: 1325. Ambiguous SET outcomes: 55.

Write policy: `wait`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 209 | 0 | 209 | 209 | 0 | 0 | 0 |
| fault | 263 | 55 | 263 | 204 | 58 | 1 | 0 |
| repair | 139 | 0 | 139 | 93 | 46 | 0 | 0 |
| reconciliation | 608 | 0 | 608 | 458 | 150 | 0 | 0 |
| post_recovery | 106 | 0 | 106 | 106 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 209 | 0.847 / 2.141 / 10.201 |
| baseline | set | 209 | 0.428 / 1.380 / 7.250 |
| baseline | wait | 209 | 0.408 / 0.847 / 3.487 |
| fault | Logical write completion | 318 | 0.847 / 116.982 / 128.972 |
| fault | set | 318 | 0.428 / 0.664 / 1.252 |
| fault | wait | 263 | 0.428 / 116.982 / 135.421 |
| repair | Logical write completion | 139 | 1.314 / 128.972 / 135.421 |
| repair | set | 139 | 0.495 / 2.248 / 5.152 |
| repair | wait | 139 | 0.632 / 128.972 / 135.421 |
| reconciliation | Logical write completion | 608 | 0.981 / 116.982 / 128.972 |
| reconciliation | set | 608 | 0.449 / 0.732 / 1.942 |
| reconciliation | wait | 608 | 0.472 / 116.982 / 128.972 |
| post_recovery | Logical write completion | 106 | 0.847 / 2.248 / 3.012 |
| post_recovery | set | 106 | 0.428 / 1.252 / 1.380 |
| post_recovery | wait | 106 | 0.408 / 1.030 / 1.521 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 209 | 103.648 |
| fault | 204 | 42.734 |
| repair | 93 | 59.852 |
| reconciliation | 458 | 76.951 |
| post_recovery | 106 | 105.977 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Replica-acknowledgement policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 3059; started 2721; completed 2721; dropped 338. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1325; checked: 1325; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAIT-satisfied subset: expected 1070; checked 1070; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAIT.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 13 |
| fault | network | 43 |
| fault | timeout | 34 |
| fault | wait_short | 58 |
| reconciliation | wait_short | 150 |
| repair | wait_short | 46 |

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Measurement methodology

- Fault selection preserves the initial primary and its replica. The destructive target is the selected primary after stable five-survivor service recovery; only strictly owned local Docker resources are affected.
- Replica faults preserve all original slot owners and do not require promotion. down_duration is a minimum hold from successful fault action completion, bounded by the recovery observation budget. Client service, failure gossip and replication interruption are separately sampled.
- Replica network isolation removes its complete bidirectional Docker bridge connectivity while Redis remains running. It is not a selective replication-only partition. Restoration uses the same owned endpoint and unchanged live process identity.
- Crash repair starts the same stopped owned container with its bounded disposable data volume. Redis Cluster identity must match the pre-fault identity. Restart changes the process; it does not simulate power loss, host/storage failure or persistence durability.
- The repair deadline covers the restart/reconnect action and subsequent synchronization. Redis's current membership is inspected without issuing CLUSTER REPLICATE or forcing slot ownership. Replica repair preserves original assignments; former-primary repair requires the original replica to retain promoted ownership and the old primary to rejoin as its replica.
- Fresh stability after repair requires all six original identities and endpoints, healthy equal slot maps, exactly one synchronized replica per primary, and all-shard client probes. Historical service recovery remains separate from final redundancy and configured acknowledgement-policy availability.
- Replication link-up and master_sync_in_progress=0 are sampled indicators, not exact zero lag, bytes preserved, or identification of partial versus full resynchronization. Raw replication INFO is retained for inspection; unobserved milestones remain absent.
- Traffic drains before validating every acknowledged immutable SET, including short/error WAIT outcomes. Ambiguous SET errors remain separate. Cancellation skips unsupported recovery/data claims and still uses fresh bounded restoration and cleanup contexts, including managed volumes.
- SET+WAIT uses one physical connection, records every successful SET even when WAIT is incomplete, and separately tracks all-shard replica-acknowledgement policy probes. Basic service recovery can succeed while that policy remains unavailable. WAIT is not a durability or strong-consistency guarantee.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `90ae5cf077684763514bcff25c9e0f86d917c8254c97316276d7725a396d31ae`.
