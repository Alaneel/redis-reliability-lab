# Redis Reliability Lab — replica-isolation

Run: `a3f6bb0e004815a4c522c39d`  
Outcome: **recovered**  
Started: 2026-10-11T16:08:53.441508+08:00  
Finished: 2026-10-11T16:09:17.405423+08:00

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

Restoration action duration: 1239.843 ms.
Stable reconciliation first observed 1569.601 ms after restoration completed.

| Process probe time | Redis PONG | Container running | Network attached | Container PID | Redis run ID | Error |
|---|---|---|---|---:|---|---|
| 2026-10-11T16:09:08.897469+08:00 | true | true | false | 4745 | afc51953fb3b973fa9a1482891dd843f3ebffe8e |  |
| 2026-10-11T16:09:11.480659+08:00 | true | true | false | 4745 | afc51953fb3b973fa9a1482891dd843f3ebffe8e |  |

## Managed node lifecycle

Target: replica `node-5` (Redis ID `2ce608dcdb1498744e354cb957d69a8b93ba365e`). Repair: `reconnect`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1239.845 ms.
Stable redundancy first observed 1569.599 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1896.341 |
| Replication interruption observed | unobserved |
| Target reachable after repair | 6010.197 |
| Membership rejoined | 6315.034 |
| Replica link up | 6010.197 |
| Synchronization in progress | unobserved |
| Synchronization complete | 6010.197 |
| Primary full-sync counter increased | unobserved |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 7577.280 |

Resynchronization evidence: `no_counter_increase_observed`; primary Redis ID `ee9d228309807a4cd804b2d420c80b9de3111243`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 1345.077 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1896.34 | 551.26 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 1896.34 | 551.26 |
| Surviving views converge | 1896.34 | 551.26 |
| First all-shard SET/GET probe succeeds | 1916.18 | 571.11 |
| Client probes sustain recovery | 3089.97 | 1744.89 |
| Topology and clients sustain recovery | 3089.97 | 1744.89 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 403 | 100.00% | 0 | 199.78 | 0.520 / 0.768 / 1.521 |
| fault | 953 | 100.00% | 0 | 199.88 | 0.472 / 0.847 / 1.849 |
| repair | 248 | 100.00% | 0 | 200.01 | 0.449 / 0.697 / 0.981 |
| reconciliation | 320 | 100.00% | 0 | 200.00 | 0.428 / 0.632 / 1.380 |
| post_recovery | 201 | 100.00% | 0 | 200.59 | 0.449 / 0.664 / 0.768 |

Acknowledged writes: 1069. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 214 | 0 | 0 | 0 | 0 | 0 | 0 |
| fault | 468 | 0 | 0 | 0 | 0 | 0 | 0 |
| repair | 128 | 0 | 0 | 0 | 0 | 0 | 0 |
| reconciliation | 157 | 0 | 0 | 0 | 0 | 0 | 0 |
| post_recovery | 102 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 214 | 0.546 / 0.807 / 1.521 |
| baseline | set | 214 | 0.520 / 0.807 / 1.521 |
| fault | Logical write completion | 468 | 0.472 / 1.030 / 1.942 |
| fault | set | 468 | 0.472 / 1.030 / 1.942 |
| repair | Logical write completion | 128 | 0.449 / 0.732 / 1.030 |
| repair | set | 128 | 0.449 / 0.732 / 1.030 |
| reconciliation | Logical write completion | 157 | 0.428 / 0.664 / 1.380 |
| reconciliation | set | 157 | 0.428 / 0.664 / 1.380 |
| post_recovery | Logical write completion | 102 | 0.472 / 0.664 / 0.697 |
| post_recovery | set | 102 | 0.472 / 0.632 / 0.697 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 214 | 106.084 |
| fault | 468 | 98.159 |
| repair | 128 | 103.231 |
| reconciliation | 157 | 98.125 |
| post_recovery | 102 | 101.790 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 2125; started 2125; completed 2125; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1069; checked: 1069; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Measurement methodology

- Fault selection preserves the initial primary and its replica. The destructive target is the selected replica; only strictly owned local Docker resources are affected.
- Replica faults preserve all original slot owners and do not require promotion. down_duration is a minimum hold from successful fault action completion, bounded by the recovery observation budget. Client service, failure gossip and replication interruption are separately sampled.
- Replica network isolation removes its complete bidirectional Docker bridge connectivity while Redis remains running. It is not a selective replication-only partition. Restoration uses the same owned endpoint and unchanged live process identity.
- Crash repair starts the same stopped owned container with its bounded disposable data volume. Redis Cluster identity must match the pre-fault identity. Restart changes the process; it does not simulate power loss, host/storage failure or persistence durability.
- The repair deadline covers the restart/reconnect action and subsequent synchronization. Redis's current membership is inspected without issuing CLUSTER REPLICATE or forcing slot ownership. Replica repair preserves original assignments; former-primary repair requires the original replica to retain promoted ownership and the old primary to rejoin as its replica.
- Fresh stability after repair requires all six original identities and endpoints, healthy equal slot maps, exactly one synchronized replica per primary, and all-shard client probes. Historical service recovery remains separate from final redundancy and configured acknowledgement-policy availability.
- Replication link-up and master_sync_in_progress=0 are sampled indicators, not exact zero lag, bytes preserved, or identification of partial versus full resynchronization. Raw replication INFO is retained for inspection; unobserved milestones remain absent.
- Traffic drains before validating every acknowledged immutable SET, including short/error WAIT outcomes. Ambiguous SET errors remain separate. Cancellation skips unsupported recovery/data claims and still uses fresh bounded restoration and cleanup contexts, including managed volumes.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `3636c096ac26e760510d748b8bbcad4b5a3352765591c52fa8d2715106b8816d`.
