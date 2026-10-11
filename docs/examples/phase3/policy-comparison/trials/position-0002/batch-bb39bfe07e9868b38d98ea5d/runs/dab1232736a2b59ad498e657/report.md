# Redis Reliability Lab — policy-aof-set

Run: `dab1232736a2b59ad498e657`  
Outcome: **recovered**  
Started: 2026-10-11T17:12:17.766032+08:00  
Finished: 2026-10-11T17:12:55.505079+08:00

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

Target: primary `node-1` (Redis ID `b21aebb5aacdad45fec76670db0fe4dc66a8ced2`). Repair: `restart`.

Repair action completed: true. Repair or synchronization deadline expired: false.
Repair action duration: 1509.084 ms.
Stable redundancy first observed 5350.976 ms after repair action completed.

A completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.

| Lifecycle observation | Since fault orchestration began, ms |
|---|---:|
| Target failure detected | 1813.866 |
| Replication interruption observed | 3383.025 |
| Target reachable after repair | 6038.904 |
| Membership rejoined | 6038.904 |
| Replica link up | 10105.573 |
| Synchronization in progress | unobserved |
| Synchronization complete | 10105.573 |
| Primary full-sync counter increased | 6038.904 |
| Primary partial-sync counter increased | unobserved |
| Replication redundancy sustained | 11386.981 |

Resynchronization evidence: `full_sync_counter_increased`; primary Redis ID `c3aff2f7c9977e891f864b4bb6c248dcb33ed09f`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.

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

Fault action completed 242.783 ms after orchestration began (measured with the monotonic runtime clock).

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | 1813.87 | 1571.08 |
| Original replica owns failed primary slots | 3383.02 | 3140.24 |
| Full usable slot coverage | 3383.02 | 3140.24 |
| Surviving views converge | 3383.02 | 3140.24 |
| First all-shard SET/GET probe succeeds | 3405.68 | 3162.90 |
| Client probes sustain recovery | 4523.55 | 4280.77 |
| Topology and clients sustain recovery | 4523.55 | 4280.77 |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 803 | 100.00% | 0 | 199.76 | 0.664 / 1.252 / 2.248 |
| fault | 576 | 84.20% | 91 | 127.24 | 0.732 / 1.942 / 149.301 |
| repair | 302 | 100.00% | 0 | 200.12 | 0.664 / 1.314 / 1.849 |
| reconciliation | 1077 | 100.00% | 0 | 199.95 | 0.632 / 1.192 / 2.141 |
| post_recovery | 400 | 100.00% | 0 | 199.98 | 0.573 / 1.252 / 2.141 |

Acknowledged writes: 1545. Ambiguous SET outcomes: 52.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | b21aebb5aacdad45fec76670db0fe4dc66a8ced2 | 2026-10-11T17:12:30.002668+08:00 | b21aebb5aacdad45fec76670db0fe4dc66a8ced2 | true |  |
| node-2 | 6a6c4217abe33ad747a033cd2ffaefebce4da276 | 2026-10-11T17:12:31.084499+08:00 | 6a6c4217abe33ad747a033cd2ffaefebce4da276 | true |  |
| node-3 | 5099554547f1d5eb9e955a97f0cab5269659d4e7 | 2026-10-11T17:12:32.182883+08:00 | 5099554547f1d5eb9e955a97f0cab5269659d4e7 | true |  |
| node-4 | 8424390f082ee10f3e0950630648bc851cf63383 | 2026-10-11T17:12:33.190983+08:00 | 8424390f082ee10f3e0950630648bc851cf63383 | true |  |
| node-5 | c3aff2f7c9977e891f864b4bb6c248dcb33ed09f | 2026-10-11T17:12:34.174327+08:00 | c3aff2f7c9977e891f864b4bb6c248dcb33ed09f | true |  |
| node-6 | 7cb8fc8b2061d17adb294d460a8ae61138cdd6de | 2026-10-11T17:12:35.189913+08:00 | 7cb8fc8b2061d17adb294d460a8ae61138cdd6de | true |  |

Before-baseline `set` policy probe at 2026-10-11T17:12:35.199358+08:00; all shards successful: true.

| Original shard index | SET acknowledged |
|---|---|
| 0 | true |
| 1 | true |
| 2 | true |

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 416 | 0 |
| fault | 230 | 52 |
| repair | 153 | 0 |
| reconciliation | 556 | 0 |
| post_recovery | 190 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 416 | 0.890 / 1.314 / 2.141 |
| baseline | set | 416 | 0.890 / 1.314 / 2.141 |
| fault | Logical write completion | 282 | 0.934 / 1.942 / 2.602 |
| fault | set | 282 | 0.890 / 1.942 / 2.602 |
| repair | Logical write completion | 153 | 0.981 / 1.677 / 2.602 |
| repair | set | 153 | 0.981 / 1.677 / 2.602 |
| reconciliation | Logical write completion | 556 | 0.890 / 1.380 / 2.360 |
| reconciliation | set | 556 | 0.890 / 1.380 / 2.360 |
| post_recovery | Logical write completion | 190 | 0.847 / 1.521 / 2.360 |
| post_recovery | set | 190 | 0.847 / 1.521 / 2.360 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 416 | 103.489 |
| fault | 230 | 50.807 |
| repair | 153 | 101.384 |
| reconciliation | 556 | 103.226 |
| post_recovery | 190 | 94.992 |

Global pacing: offered 3488; started 3158; completed 3158; dropped 330. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 1545; checked: 1545; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|
| fault | CLUSTERDOWN | 15 |
| fault | network | 37 |
| fault | timeout | 39 |

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Measurement methodology

- Fault selection preserves the initial primary and its replica. The destructive target is the selected primary after stable five-survivor service recovery; only strictly owned local Docker resources are affected.
- Replica faults preserve all original slot owners and do not require promotion. down_duration is a minimum hold from successful fault action completion, bounded by the recovery observation budget. Client service, failure gossip and replication interruption are separately sampled.
- Replica network isolation removes its complete bidirectional Docker bridge connectivity while Redis remains running. It is not a selective replication-only partition. Restoration uses the same owned endpoint and unchanged live process identity.
- Administrative endpoint-disconnection evidence is separate from Redis INFO link or connected-replica indicators. A short disconnection can finish before TCP/link failure detection; absent interruption or synchronization-counter milestones remain unobserved and are not invented from the Docker interface state.
- Crash repair starts the same stopped owned container with its bounded disposable data volume. Redis Cluster identity must match the pre-fault identity. Restart changes the process; it does not simulate power loss, host/storage failure or persistence durability.
- The repair deadline covers the restart/reconnect action and subsequent synchronization. Redis's current membership is inspected without issuing CLUSTER REPLICATE or forcing slot ownership. Replica repair preserves original assignments; former-primary repair requires the original replica to retain promoted ownership and the old primary to rejoin as its replica.
- Fresh stability after repair requires all six original identities and endpoints, healthy equal slot maps, exactly one synchronized replica per primary, and all-shard client probes. Historical service recovery remains separate from final redundancy and configured acknowledgement-policy availability.
- Replication link-up and master_sync_in_progress=0 are sampled indicators, not exact zero lag, bytes preserved, or identification of partial versus full resynchronization. Raw replication INFO is retained for inspection; unobserved milestones remain absent.
- Traffic drains before validating every acknowledged immutable SET, including short/error WAIT/WAITAOF outcomes. Ambiguous SET errors remain separate. Cancellation skips unsupported recovery/data claims and still uses fresh bounded restoration and cleanup contexts, including managed volumes.
- All AOF modes and write policies use six-node quiet runtime preparation followed by exactly three all-shard policy SET/GET warmup writes before baseline starts. Warmup writes are excluded from the workload ledger and metrics; no WAITAOF compatibility command is issued after baseline to alter crash-trial durability.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `de05e21df2ed5b0a82dc0af8fa135f9173da7c1fdf43019ed7cd93af6d12f3cf`.
