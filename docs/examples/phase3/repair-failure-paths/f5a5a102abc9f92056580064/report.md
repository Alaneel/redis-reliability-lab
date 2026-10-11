# Redis Reliability Lab — primary-failure

Run: `f5a5a102abc9f92056580064`  
Outcome: **failed**  
Started: 2026-10-11T16:19:34.963947+08:00  
Finished: 2026-10-11T16:19:56.006896+08:00

## Recovery

| Condition | Observed |
|---|---|
| Surviving topology converged | true |
| All-shard client recovery sustained | true |
| Topology and client recovery jointly sustained | false |
| All acknowledged writes verified | false |
| Target process restarted | false |
| Original replication redundancy restored | false |
| Recovery deadline expired | false |
| Cleanup verified | true |


## Managed node lifecycle

Target: replica `node-5` (Redis ID `093ac1cb6538847e01f545f5ad4bdd03ab1f0db7`). Repair: `restart`.

Repair action completed: false. Repair or synchronization deadline expired: true.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1600.065 |
| Replication interruption observed | 256.860 |
| Target reachable after repair | unobserved |
| Membership rejoined | unobserved |
| Replica link up | unobserved |
| Synchronization in progress | unobserved |
| Synchronization complete | unobserved |
| Primary full-sync counter increased | unobserved |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | unobserved |

Lifecycle error: context deadline exceeded

## Independent recovery dimensions

Target observed role: `unobserved`.

| Independent condition | Observed |
|---|---|
| Target process live | false |
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

Fault action completed 250.537 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1600.06 | 1349.53 |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | 256.86 | 6.32 |
| Surviving views converge | 256.86 | 6.32 |
| First all-shard SET/GET probe succeeds | 377.98 | 127.45 |
| Client probes sustain recovery | 2612.67 | 2362.13 |
| Topology and clients sustain recovery | 2612.67 | 2362.13 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 419 | 100.00% | 0 | 199.43 | 0.573 / 0.807 / 1.314 |
| fault | 855 | 100.00% | 0 | 200.25 | 0.472 / 0.664 / 0.934 |
| repair | 20 | 100.00% | 0 | 198.99 | 0.602 / 0.768 / 1.043 |

Acknowledged writes: 653. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 221 | 0 | 0 | 0 | 0 | 0 | 0 |
| fault | 423 | 0 | 0 | 0 | 0 | 0 | 0 |
| repair | 9 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 221 | 0.573 / 0.807 / 1.252 |
| baseline | set | 221 | 0.546 / 0.768 / 1.192 |
| fault | Logical write completion | 423 | 0.472 / 0.664 / 0.934 |
| fault | set | 423 | 0.449 / 0.632 / 0.934 |
| repair | Logical write completion | 9 | 0.602 / 1.043 / 1.043 |
| repair | set | 9 | 0.573 / 0.972 / 0.972 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 221 | 105.187 |
| fault | 423 | 99.071 |
| repair | 9 | 89.545 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 1294; started 1294; completed 1294; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: false; expected: 0; checked: 0; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: false.
Validation was not attempted; 653 acknowledged writes remain unchecked.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Errors

- repair node-5: context deadline exceeded

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `703d804056ef9e020784ff013daae6c6ed9333117dc1d70f84ce45fe63f0dd2d`.
