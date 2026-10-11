# Redis Reliability Lab — primary-failure

Run: `488a31734f8e918c27af7ca0`  
Outcome: **aborted**  
Started: 2026-10-11T16:20:19.444794+08:00  
Finished: 2026-10-11T16:20:40.669214+08:00

## Recovery

| Condition | Observed |
|---|---|
| Surviving topology converged | false |
| All-shard client recovery sustained | false |
| Topology and client recovery jointly sustained | false |
| All acknowledged writes verified | false |
| Target process restarted | true |
| Original replication redundancy restored | false |
| Recovery deadline expired | false |
| Cleanup verified | true |


## Managed node lifecycle

Target: replica `node-5` (Redis ID `d7f6ccec8e4e24319718f05d5fa589fa554659cd`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1511.278 ms.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1607.801 |
| Replication interruption observed | 258.613 |
| Target reachable after repair | unobserved |
| Membership rejoined | unobserved |
| Replica link up | unobserved |
| Synchronization in progress | unobserved |
| Synchronization complete | unobserved |
| Primary full-sync counter increased | unobserved |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | unobserved |

Lifecycle error: original six-node redundancy and clients did not stabilize within the 30s repair budget

## Independent recovery dimensions

Target observed role: `slave`.

| Independent condition | Observed |
|---|---|
| Target process live | true |
| Cluster membership healthy | false |
| Full usable slot coverage | true |
| All shard leaders available | true |
| Client reads available | true |
| Client writes available | true |
| Acknowledged data verified | unobserved |
| Configured write policy available | true |
| Replication redundancy restored | false |
| Persistence recovery verified | unobserved |

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 252.570 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1607.80 | 1355.23 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 258.61 | 6.04 |
| Surviving views converge | 258.61 | 6.04 |
| First all-shard SET/GET probe succeeds | 379.67 | 127.10 |
| Client probes sustain recovery | 1731.64 | 1479.07 |
| Topology and clients sustain recovery | 1731.64 | 1479.07 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 420 | 100.00% | 0 | 199.71 | 0.573 / 0.807 / 1.598 |
| fault | 677 | 100.00% | 0 | 199.89 | 0.449 / 0.664 / 1.135 |
| repair | 303 | 100.00% | 0 | 200.48 | 0.573 / 1.380 / 2.602 |
| reconciliation | 0 | 0.00% | 0 | 0.00 | unobserved |

Acknowledged writes: 696. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 221 | 0 | 0 | 0 | 0 | 0 | 0 |
| fault | 344 | 0 | 0 | 0 | 0 | 0 | 0 |
| repair | 131 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 221 | 0.602 / 0.807 / 1.252 |
| baseline | set | 221 | 0.573 / 0.768 / 1.192 |
| fault | Logical write completion | 344 | 0.472 / 0.664 / 1.314 |
| fault | set | 344 | 0.449 / 0.602 / 1.314 |
| repair | Logical write completion | 131 | 0.602 / 1.521 / 2.732 |
| repair | set | 131 | 0.546 / 1.449 / 2.602 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 221 | 105.086 |
| fault | 344 | 101.571 |
| repair | 131 | 86.676 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 1400; started 1400; completed 1400; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: false; expected: 0; checked: 0; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: false.
Validation was not attempted; 696 acknowledged writes remain unchecked.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Errors

- original six-node redundancy and clients did not stabilize within the 30s repair budget
- context canceled

## Measurement methodology

- Fault selection preserves the initial primary and its replica. The destructive target is the selected replica; only strictly owned local Docker resources are affected.
- Replica faults preserve all original slot owners and do not require promotion. down_duration is a minimum hold from successful fault action completion, bounded by the recovery observation budget. Client service, failure gossip and replication interruption are separately sampled.
- Replica network isolation removes its complete bidirectional Docker bridge connectivity while Redis remains running. It is not a selective replication-only partition. Restoration uses the same owned endpoint and unchanged live process identity.
- Administrative endpoint-disconnection evidence is separate from Redis INFO link or connected-replica indicators. A short disconnection can finish before TCP/link failure detection; absent interruption or synchronization-counter milestones remain unobserved and are not invented from the Docker interface state.
- Crash repair starts the same stopped owned container with its bounded disposable data volume. Redis Cluster identity must match the pre-fault identity. Restart changes the process; it does not simulate power loss, host/storage failure or persistence durability.
- The repair deadline covers the restart/reconnect action and subsequent synchronization. Redis's current membership is inspected without issuing CLUSTER REPLICATE or forcing slot ownership. Replica repair preserves original assignments; former-primary repair requires the original replica to retain promoted ownership and the old primary to rejoin as its replica.
- Fresh stability after repair requires all six original identities and endpoints, healthy equal slot maps, exactly one synchronized replica per primary, and all-shard client probes. Historical service recovery remains separate from final redundancy and configured acknowledgement-policy availability.
- Replication link-up and master_sync_in_progress=0 are sampled indicators, not exact zero lag, bytes preserved, or identification of partial versus full resynchronization. Raw replication INFO is retained for inspection; unobserved milestones remain absent.
- Traffic drains before validating every acknowledged immutable SET, including short/error WAIT outcomes. Ambiguous SET errors remain separate. Cancellation skips unsupported recovery/data claims and still uses fresh bounded restoration and cleanup contexts, including managed volumes.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `3b13d98713125809f4730f1041b2d6ad83c5c30c7457042a17d4041038a39eaa`.
