# Redis Reliability Lab — dataset-trial

Run: `5e2f1051ec26c8e7bb07b752`  
Outcome: **recovered**  
Started: 2026-10-11T19:55:57.89242+08:00  
Finished: 2026-10-11T19:56:27.329602+08:00

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

Target: replica `node-5` (Redis ID `d20137f353081ceae2db3da69edad89e1dd2869a`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1804.297 ms.
Stable redundancy first observed 4747.076 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1963.172 |
| Replication interruption observed | 278.947 |
| Target reachable after repair | 6116.795 |
| Membership rejoined | 6116.795 |
| Replica link up | 9401.844 |
| Synchronization in progress | unobserved |
| Synchronization complete | 9401.844 |
| Primary full-sync counter increased | 6116.795 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 10853.646 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `2698c61383131ae1c1f41a01764d188c71f4c4e4`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 270.988 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1963.17 | 1692.18 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 278.95 | 7.96 |
| Surviving views converge | 278.95 | 7.96 |
| First all-shard SET/GET probe succeeds | 403.57 | 132.59 |
| Client probes sustain recovery | 1525.29 | 1254.31 |
| Topology and clients sustain recovery | 1525.29 | 1254.31 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 356 | 100.00% | 0 | 169.73 | 0.602 / 0.847 / 1.849 |
| fault | 729 | 100.00% | 0 | 169.45 | 0.697 / 2.478 / 6.905 |
| repair | 307 | 100.00% | 0 | 170.14 | 0.768 / 1.314 / 2.039 |
| reconciliation | 841 | 100.00% | 0 | 169.94 | 0.520 / 0.768 / 1.677 |
| post_recovery | 171 | 100.00% | 0 | 170.78 | 0.472 / 0.664 / 0.981 |

Acknowledged writes: 1248. Ambiguous SET outcomes: 0.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 185 | 0 |
| fault | 393 | 0 |
| repair | 154 | 0 |
| reconciliation | 423 | 0 |
| post_recovery | 93 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 185 | 0.632 / 0.934 / 1.677 |
| baseline | set | 185 | 0.602 / 0.890 / 1.598 |
| fault | Logical write completion | 393 | 0.697 / 3.012 / 6.905 |
| fault | set | 393 | 0.664 / 2.869 / 6.905 |
| repair | Logical write completion | 154 | 0.768 / 1.314 / 1.677 |
| repair | set | 154 | 0.732 / 1.192 / 1.521 |
| reconciliation | Logical write completion | 423 | 0.546 / 0.768 / 1.380 |
| reconciliation | set | 423 | 0.520 / 0.732 / 1.380 |
| post_recovery | Logical write completion | 93 | 0.495 / 0.732 / 4.066 |
| post_recovery | set | 93 | 0.495 / 0.697 / 4.048 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 185 | 88.203 |
| fault | 393 | 91.348 |
| repair | 154 | 85.347 |
| reconciliation | 423 | 85.476 |
| post_recovery | 93 | 92.883 |

Global pacing: offered 2406; started 2404; completed 2404; dropped 2. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1248; checked: 1248; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `98c05c302b4226ae2941439d237cb243952ffaf69327345b6c46771044a9efd1`.
