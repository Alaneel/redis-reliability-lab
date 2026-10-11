# Redis Reliability Lab — replica-crash

Run: `abd3356a3b9c2474c06cb6e6`  
Outcome: **recovered**  
Started: 2026-10-11T16:07:36.072976+08:00  
Finished: 2026-10-11T16:08:04.198464+08:00

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

Target: replica `node-5` (Redis ID `3183b588fce8a80c8c1bd5992355615de3231bd0`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1519.627 ms.
Stable redundancy first observed 1238.892 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1718.323 |
| Replication interruption observed | 245.255 |
| Target reachable after repair | 5008.592 |
| Membership rejoined | 5008.592 |
| Replica link up | 5008.592 |
| Synchronization in progress | unobserved |
| Synchronization complete | 5008.592 |
| Primary full-sync counter increased | unobserved |
| Primary partial-sync counter increased | 5008.592 |
| Replication redundancy sustained | 6244.397 |

Resynchronization evidence: `partial_sync_counter_increased`; primary Redis ID `6ff7bf4084e773b5ddf2fb43a84f414c9f349a6d`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 242.281 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1718.32 | 1476.04 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 245.25 | 2.97 |
| Surviving views converge | 245.25 | 2.97 |
| First all-shard SET/GET probe succeeds | 264.55 | 22.27 |
| Client probes sustain recovery | 1447.32 | 1205.04 |
| Topology and clients sustain recovery | 1447.32 | 1205.04 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 403 | 100.00% | 0 | 199.52 | 0.495 / 0.732 / 1.449 |
| fault | 698 | 100.00% | 0 | 200.24 | 0.449 / 0.697 / 1.192 |
| repair | 304 | 100.00% | 0 | 200.04 | 0.472 / 1.192 / 17.448 |
| reconciliation | 253 | 100.00% | 0 | 199.45 | 0.428 / 1.314 / 2.039 |
| post_recovery | 200 | 100.00% | 0 | 199.98 | 0.408 / 0.573 / 0.627 |

Acknowledged writes: 936. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 213 | 0 | 0 | 0 | 0 | 0 | 0 |
| fault | 353 | 0 | 0 | 0 | 0 | 0 | 0 |
| repair | 131 | 0 | 0 | 0 | 0 | 0 | 0 |
| reconciliation | 143 | 0 | 0 | 0 | 0 | 0 | 0 |
| post_recovery | 96 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 213 | 0.520 / 0.732 / 1.449 |
| baseline | set | 213 | 0.495 / 0.732 / 1.449 |
| fault | Logical write completion | 353 | 0.449 / 0.697 / 1.192 |
| fault | set | 353 | 0.449 / 0.697 / 1.192 |
| repair | Logical write completion | 131 | 0.472 / 1.252 / 11.247 |
| repair | set | 131 | 0.472 / 1.252 / 11.247 |
| reconciliation | Logical write completion | 143 | 0.449 / 1.030 / 2.039 |
| reconciliation | set | 143 | 0.449 / 1.030 / 2.039 |
| post_recovery | Logical write completion | 96 | 0.428 / 0.573 / 0.627 |
| post_recovery | set | 96 | 0.428 / 0.573 / 0.623 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 213 | 105.453 |
| fault | 353 | 101.268 |
| repair | 131 | 86.200 |
| reconciliation | 143 | 112.731 |
| post_recovery | 96 | 95.992 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 1858; started 1858; completed 1858; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 936; checked: 936; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `5cd7588193d4582e7bb93d778105aefd6829d73098973dc91726b4b54064e39b`.
