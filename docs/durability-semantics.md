# Persistence and durability semantics

Real all-process crash/restart runs verified complete loss in the memory control and preservation in both AOF modes. Counterbalanced persistence and SET/WAIT/WAITAOF comparisons each completed nine fresh trials with verified cleanup. The evidence concerns abrupt Redis process termination with retained storage and a running host. The earlier [replica lifecycle examples](replica-lifecycle.md) establish service and redundancy behavior; acknowledgement policy, readable values and persistence recovery remain separate facts.

## Observed persistence results

The three fresh runs executed in blocks, memory then everysec then always, using SET acknowledgement policy and the same pinned `redis:8.10.2-alpine` image, host, topology, 4-second baseline, 8 workers, aggregate rate of 200 requests/s, 50% reads, seed 42, 128-byte values, and 500 ms operation timeout. The host was macOS arm64 with Docker Desktop and Docker Engine 24.0.6. Actual configuration, resolved image, host metadata, per-node runtime settings, data files, process identities, and action order remain in each JSON report.

| Mode and saved evidence | Outcome | Frozen acknowledged writes checked | Missing | Persistence recovery observed | Cleanup verified |
|---|---|---:|---:|---|---|
| [memory](examples/phase3/persistence-memory/report.md) ([JSON](examples/phase3/persistence-memory/report.json)) | partial | 432 / 432 | 432 | Not applicable; AOF/RDB persistence disabled | true |
| [aof-everysec](examples/phase3/persistence-aof-everysec/report.md) ([JSON](examples/phase3/persistence-aof-everysec/report.json)) | recovered | 433 / 433 | 0 | true | true |
| [aof-always](examples/phase3/persistence-aof-always/report.md) ([JSON](examples/phase3/persistence-aof-always/report.json)) | recovered | 422 / 422 | 0 | true | true |

All three validations completed with zero corrupt, unreadable, or unverified values and no ambiguous SET outcomes. Memory restored the original six-node topology and working diagnostics while independently reporting data loss. AOF file inventories recorded the actual hybrid `.base.rdb`, incremental `.incr.aof`, and manifest files. After all six processes had been stopped, replication could copy a recovered persisted shard dataset to its replica; these results do not prove that each of the six AOFs independently contained or physically persisted every acknowledged write.

| Mode | Offered / started / completed / dropped requests | Baseline throughput, requests/s | Successful SET p50 / p95 / p99, ms | Drain to first / last kill, ms | All stopped to first restart, ms |
|---|---|---:|---|---|---:|
| memory | 828 / 828 / 828 / 0 | 199.932 | 0.428 / 0.520 / 0.664 | 0.167 / 1470.794 | 1001.563 |
| aof-everysec | 828 / 828 / 828 / 0 | 199.772 | 0.632 / 0.890 / 1.192 | 0.170 / 1578.968 | 1000.485 |
| aof-always | 829 / 816 / 816 / 13 | 196.720 | 1.030 / 2.141 / 62.038 | 0.180 / 1503.391 | 1001.214 |

These are individual trial measurements, with separate per-run latency quantiles and actual demand counters. There is one trial per mode, no pooled latency percentile, no statistical significance claim, and no reliable performance ranking. The roughly 1.5-second sequential killing interval and traffic drain permit background fsync to progress. Workload traffic was absent during fault and restoration, so those phases have no workload availability denominator.

The three-mode integration suite passed in 167.351 seconds. Applied restart-command failure (`3c360…`, failed), restoration timeout (`8fc…`, failed), and cancellation (`593fa…`, aborted) passed their failure-path assertions in 89.145 seconds; all verified cleanup. The [Phase 3 validation record](phase3-validation.md) contains their execution and evidence details.

### Counterbalanced persistence comparison

A separate [nine-trial comparison](examples/phase3/persistence-comparison/comparison.md) ([complete JSON](examples/phase3/persistence-comparison/comparison.json)) completed three fresh trials per mode with all individual evidence and cleanup verified. Its seed was `20261011`, workload seed `42`, and actual condition order was `[2,1,3,1,3,2,3,2,1]`: condition 1 memory, 2 everysec, 3 always. Each condition occupied every round position once. The overall partial outcome retains the expected memory controls rather than discarding them.

| Condition | Trials recovered / partial | Frozen writes checked | Missing | Evidence |
|---|---:|---:|---:|---|
| memory | 0 / 3 | 1,251 | 1,251 | [Condition summary](examples/phase3/persistence-comparison/conditions/condition-01/summary.md) |
| aof-everysec | 3 / 0 | 1,251 | 0 | [Condition summary](examples/phase3/persistence-comparison/conditions/condition-02/summary.md) |
| aof-always | 3 / 0 | 1,252 | 0 | [Condition summary](examples/phase3/persistence-comparison/conditions/condition-03/summary.md) |

All 3,754 frozen acknowledged writes were checked with zero corrupt, unreadable, unverified or ambiguous outcomes. Every trial verified cleanup. The raw comparison, condition summaries and individual reports were copied byte for byte. Three trials per condition do not establish reliability, significance or a recovery distribution; upper-tail recovery estimates remain suppressed at this sample size. Latency quantiles stay per trial, and the position-balanced order does not balance every predecessor pair or remove host/storage effects. The fault's running-host, sequential-kill and post-restart replication limits remain unchanged.

This persistence cohort preceded the common three-shard AOF policy warmup addition. Its saved method and reports remain the executed evidence; the later policy cohort below used the complete common quiet-preparation and selected-policy warmup path.

## Three controlled modes

| Mode | Redis configuration | What the frozen-ledger experiment examines |
|---|---|---|
| `memory` | Automatic RDB snapshots disabled; AOF disabled; replication RDB cleanup enabled | Negative control: retained membership files do not preserve volatile values when every Redis process is stopped |
| `aof-everysec` | Automatic RDB snapshots disabled; hybrid AOF enabled; `appendfsync everysec` | Value recovery after same-container process restart with background fsync |
| `aof-always` | Automatic RDB snapshots disabled; hybrid AOF enabled; `appendfsync always` | Value recovery with the configured AOF fsync-before-reply policy |

RDB records snapshots; AOF records changes for replay at startup. Redis distinguishes periodic AOF fsync from fsync when appended command batches complete. Concurrent clients can share an fsync batch under `always`. The experiment records the actual runtime settings instead of treating a requested mode as proof of the active configuration. [Redis persistence documentation](https://redis.io/docs/latest/operate/oss_and_stack/management/persistence/)

The official pinned image retains its bundled modules. Its RedisSearch startup rejected AOF with the RDB preamble disabled, so AOF modes use `aof-use-rdb-preamble yes`: an RDB-format AOF base plus incremental AOF files tracked by a manifest. `save ""` disables scheduled RDB snapshots; it does not prohibit the RDB-format AOF base. Memory mode keeps AOF disabled and its preamble setting `no`. [Pinned Redis AOF base formats](https://github.com/redis/redis/blob/8.10.2/redis.conf#L1524-L1527)

`repl-diskless-load swapdb` can fall back to disk when loaded modules lack the required error-handling or asynchronous-load support. Retained nodes therefore configure `rdb-del-sync-files yes`; Redis removes replication RDB files when both RDB and AOF persistence are disabled. This setting is ignored with AOF enabled. The memory control still requires observed absence of RDB/AOF value files before baseline and after restart; the lab does not delete files to manufacture that result. [Pinned diskless-load selection](https://github.com/redis/redis/blob/8.10.2/src/replication.c#L2090-L2108), [replication-file cleanup configuration](https://github.com/redis/redis/blob/8.10.2/redis.conf#L511-L522)

The checked-in definitions are [memory](../experiments/persistence-memory.yaml), [everysec](../experiments/persistence-aof-everysec.yaml), and [always](../experiments/persistence-aof-always.yaml). They use the same SET policy, Redis image, baseline duration, worker/rate limits, payload, seed, topology, node timeout, and observation budgets. Fresh environments avoid reusing data from another mode. Execution order, effective configuration, resolved image, host metadata, and per-run counts must accompany any reported comparison. One trial per mode cannot establish a recovery distribution, statistical significance, or a reliable performance ranking.

## Freeze before the fault

The experiment first provisions six managed containers with disposable local `/data` volumes and verifies healthy original owners and replica assignments. Administrative evidence binds each node to its original Redis Cluster ID, process identity, runtime persistence configuration, persistence/keyspace INFO, and data-file listing. Before baseline traffic, preparation waits for sampled loading, replication synchronization, AOF rewrite/scheduled rewrite, and pending background fsync to become quiet. This preparation does not request SAVE, BGSAVE, BGREWRITEAOF, or an explicit fsync barrier. Docker restart policy must be `no` so no stopped process silently restarts during the barrier.

Baseline traffic generates unique immutable values using the existing SET acknowledgement ledger. At the end of the baseline, a read-only pre-crash topology snapshot verifies original owners and assignments while traffic is still running. The runner then stops the dispatcher, drains every in-flight operation, and freezes acknowledged and ambiguous counts. Successful SETs remain the validation population; an errored SET may have executed but remains outside that population. No extra administrative or topology observation is inserted between draining and the first kill, avoiding an additional delay in which background fsync could progress. No fresh SET/GET service probes or workload writes may repopulate values before frozen-ledger validation.

The runner records each sequential owned SIGKILL request/completion and stopped-process observation. Before restarting any container, it requires fresh evidence that all six processes are stopped. This all-stopped barrier removes a live Redis replica as a source for restoring otherwise lost volatile values. Containers, bridge, and data volumes remain the same managed resources.

After restart, replication can copy one recovered persisted shard dataset to its replica. Preserving the frozen ledger therefore establishes recovery through the configured persistence-and-replication process; it does not prove that each of the six AOF file sets independently contained or physically persisted every acknowledged write. Per-node settings, files, and healthy links provide separate runtime evidence, not six independent disk-acknowledgement guarantees.

Restart requests and receipts preserve their actual order, including failed requests with no receipt. Each restarted process must retain its original Redis Cluster ID and endpoints while changing process start identity. The six-node readiness predicate requires original owners and replica assignments; it does not require promotion. Early physical restart receipts can contain transient full-sync or AOF status and remain unchanged.

After read-only topology stability, six fresh administrative observations record quiet loading/replication/rewrite status, matching persistence configuration, and actual data files under the same restoration budget. The report saves these as `persistence_fault.restored`, separately from physical restart receipts. They occur before frozen-ledger reads and establish the runtime persistence criterion; an early PONG or temporary replication RDB is insufficient. Missing or expired proof cannot establish persistence recovery.

## Validate values before diagnostic writes

After read-only topology readiness, validation reads every frozen acknowledged key and compares its exact value. Its expected count must match the frozen ledger. Missing, corrupt, unreadable, and unverified counts remain separate. `complete: true` means the attempted verification covered the ledger; it can coexist with measured missing values and a failed data verdict.

Only after that validation completes can separate diagnostics write fresh probe keys and measure sustained client service. Their request time and service milestone are recorded separately. A later working SET/GET probe cannot repair the data verdict, grow the frozen ledger, or hide missing immutable values. Main `validation` and `persistence_fault.frozen_validation` identify the same pre-diagnostic result.

Workload traffic is not offered during fault or restart observation. There is no workload availability denominator in those phases, and no zero-percent availability or zero-throughput interpretation. Baseline latency/throughput remains measured traffic; later probes are diagnostics rather than workload measurements. Offered/started/completed/dropped counts expose actual baseline demand and saturation.

## Independent results

Process restart, six-node topology, replica readiness, diagnostic client service, frozen data, and persistence evidence are separate conditions. A memory control can restore all processes, owners, replica links, and clients while reporting missing frozen values and a partial outcome. That is an observed control result, not a successful data-recovery claim.

For memory mode, `persistence_recovery_observed` is absent because there is no enabled persistence mechanism under test. The independent `persistence_recovered` dimension is false when the disabled runtime configuration was observed. Frozen data verification independently records preservation or loss.

An AOF persistence-recovery flag requires the all-stopped/restart evidence, actual runtime configuration/loading evidence, and complete preservation of the frozen acknowledged ledger. The file inventory records the manifest, RDB-format AOF base, incremental AOF, and any replication synchronization RDB separately. Restart or `cluster_state:ok` cannot establish preservation. Missing, contradictory, or incomplete evidence remains unknown or a failed observation rather than a fabricated zero-duration success. The recovery milestone records the first accepted verification, not an exact disk or Redis event time.

The report retains drain-to-first-kill, drain-to-last-kill, and all-stopped-to-first-restart durations measured by the runtime monotonic clock. Action timestamps show wall time and order. Raw process observations include Redis run/Cluster IDs, CONFIG GET output, INFO persistence/keyspace, data-file listings, and `/data` size. These indicate runtime behavior and files; they do not establish a file's physical durability.

## What this fault establishes

This is sequential abrupt Redis process termination and same-container restart while the Docker VM, host, and storage stay running. It is different from graceful shutdown, replacing a container with empty storage, deleting a volume, physical power loss, and host/storage failure. Data files and operating-system page cache survive this process-level fault.

Traffic drain, configuration/ownership probes, sequential kill calls, and stopped-process inspections consume time during which background fsync can progress. An everysec run that preserves all values therefore cannot establish that acknowledged writes survive sudden power loss, a strict one-second loss bound, or behavior equivalent to always. The report exposes those intervals instead of calling them an instantaneous simultaneous crash. [Redis AOF fsync policies](https://redis.io/docs/latest/operate/oss_and_stack/management/persistence/)

SET success, WAIT satisfaction, persistence settings, and observed final values establish different facts. WAIT counts replication acknowledgements for earlier writes on the same connection; it does not itself establish disk persistence or universal loss-free failover. AOF mode does not convert the lab into a distributed transaction or linearizability test. [WAIT semantics](https://redis.io/docs/latest/commands/wait/), [acknowledgement methodology](consistency-comparison.md)

## WAITAOF acknowledgement policy

WAITAOF requests AOF fsync acknowledgements for preceding writes on the same client connection. This lab uses `SET` followed by `WAITAOF 1 1 timeout`: one local and one replica AOF acknowledgement. The reply is exactly two integers, local then replica; both must meet their requested counts for policy success. A server timeout can return a valid short pair, including `0/0`, `0/1`, or `1/0`. A client or command error supplies no observed pair. WAITAOF does not make Redis strongly consistent or ensure universal loss-free failover. [Official WAITAOF semantics](https://redis.io/docs/latest/commands/waitaof/), [pinned implementation](https://github.com/redis/redis/blob/8.10.2/src/replication.c#L4545-L4580)

The write executor preserves one physical connection for SET and its selected acknowledgement command. It follows only validated managed-node redirects before SET success, then never replays SET on a transport failure or replays an errored WAITAOF. Each successful SET remains in the full acknowledged ledger even if WAITAOF is short, fails, or cannot start. A separate WAITAOF-satisfied subset records only writes meeting both requested counts; its final value check does not replace the full ledger.

All AOF policy trials observe the same six-node quiet runtime preparation and exercise their selected policy on all three original shards before baseline. This records actual command compatibility and warms the same managed routing path without attributing probes to workload traffic. WAITAOF request/reply masks are saved separately. It is never used as a post-baseline persistence preflight, since flushing earlier workload writes after drain could bias the frozen-ledger result.

Reports keep WAITAOF attempted, unattempted, satisfied, short, error, and client-timeout counts separate from WAIT. The paired `local/replicas` histogram excludes errors and malformed replies, and per-shard probe masks distinguish observed zero counts from unknown replies. Local-satisfied and replica-satisfied counts overlap; adding them would misstate success. Logical `set_waitaof` latency includes both stages, while the `waitaof` stage records only an attempted command. Ordinary SET/GET service and the configured acknowledgement policy have independent recovery flags.

This policy requires observed AOF-enabled runtime configuration and a primary connection in the pinned Redis 8.10.2 environment. Its acknowledgements concern Redis AOF fsync progress, including the hybrid RDB-format AOF base and incremental AOF. A process-crash run still keeps the host and storage running; it cannot establish physical power-loss durability.

### Observed WAITAOF primary repair

The standalone [WAITAOF repair report](examples/phase3/waitaof-repair/report.md) ([JSON](examples/phase3/waitaof-repair/report.json), run `34b6d0e4fd105ef1b3e0e9e0`) recovered with six observed AOF-prepared nodes and valid before-baseline `1/1` replies on all three shards. Its full ledger checked 1,635 SET acknowledgements with no missing, corrupt, unreadable or unverified values; 37 ambiguous SET outcomes remained separate. Of 1,635 valid WAITAOF replies, 1,372 were satisfied `1/1` pairs and 263 were local-only `1/0` short pairs. There were no WAITAOF command errors or unattempted calls. The 263 short writes remained in full-ledger validation.

| Observation | Measured timing |
|---|---:|
| Original replica owns failed primary slots | 3.539445 s after fault orchestration began |
| Ordinary service jointly stable | 4.752659 s after fault orchestration began |
| Former-primary restart action | 1.590417 s |
| All-six redundancy stable | 5.819080 s after restart completed |
| WAITAOF policy stable | 12.282862 s after fault orchestration began |

Ordinary service recovered before the restored replica could satisfy the disk-acknowledgement policy. Local-satisfied count was 1,635 and replica-satisfied count was 1,372; these overlap and are not additional writes. Cleanup was verified. This individual observation does not establish a loss-free guarantee or a cross-policy performance ranking.

Live Redis copies remained available through this primary fault. Its independent persistence-recovered dimension remains unobserved: lifecycle readiness and readable values do not establish recovery across an all-process storage-loading boundary. The frozen-ledger persistence experiments above require that separate barrier and fresh runtime evidence.

### Frozen WAITAOF persistence proof

A separate [all-process WAITAOF report](examples/phase3/waitaof-persistence/report.md) ([JSON](examples/phase3/waitaof-persistence/report.json), run `adcd47811ff02526680c06fa`) froze and checked 433 acknowledged values, all with satisfied `1/1` replies. It observed zero missing, corrupt, unreadable, unverified or ambiguous values. All six processes were verified stopped before any restart, and frozen GET validation completed before diagnostic writes. Fresh restored AOF runtime evidence, original topology, observed value preservation and cleanup all passed. Its persistence-recovery flag is therefore observed true under the defined process fault, with the running-host and per-file limitations described above.

### Counterbalanced SET / WAIT / WAITAOF comparison

The [complete policy comparison](examples/phase3/policy-comparison/comparison.md) ([JSON](examples/phase3/policy-comparison/comparison.json), `comparison-321638985ff49f9040049bf4`) recovered in all nine trials from 17:11:38 to 17:17:22 Singapore time on 11 October 2026. Three fresh trials per policy used matched AOF-always, primary SIGKILL with former-primary repair, the same image/host, workers, rate, payload, deadlines and workload seed `42`. All recorded six quiet prepared nodes and three successful selected-policy probes before baseline. Order seed `20261011` produced WAIT → SET → WAITAOF, then SET → WAITAOF → WAIT, then WAITAOF → WAIT → SET.

| Policy | SET acknowledgements checked | Command satisfied / short | Ambiguous SET outcomes | Evidence |
|---|---:|---|---:|---|
| SET | 4,646 | Not applicable | 124 | [Condition summary](examples/phase3/policy-comparison/conditions/condition-01/summary.md) |
| WAIT | 4,575 | 3,907 / 668 | 124 | [Condition summary](examples/phase3/policy-comparison/conditions/condition-02/summary.md) |
| WAITAOF | 4,759 | 4,029 / 730 | 111 | [Condition summary](examples/phase3/policy-comparison/conditions/condition-03/summary.md) |

All 13,980 acknowledged values were checked with zero missing, corrupt, unreadable or unverified values. The 359 ambiguous SET outcomes remain outside that population. WAIT and WAITAOF recorded zero command errors and zero unattempted calls. WAITAOF's valid reply pairs were 4,029 `1/1` and 730 `1/0`: all 4,759 satisfied the local count, while 4,029 satisfied the replica count. Both subsets and the full SET ledgers were verified separately. Short-command writes remained in the full ledgers.

| Policy | Median individual baseline successful logical-write p50 / p95 / p99, ms | Median successful baseline writes/s | Median ordinary service stable after fault start, ms | Median redundancy stable after repair completed, ms | Median policy stable after fault start, ms |
|---|---|---:|---:|---:|---:|
| SET | 1.029797 / 1.521483 / 2.360325 | 103.489 | 4523.550 | 5350.976 | Not applicable |
| WAIT | 1.677436 / 3.012442 / 3.661644 | 103.694 | 4717.017 | 4983.597 | 11080.822 |
| WAITAOF | 1.597558 / 2.868992 / 3.321219 | 103.523 | 4656.215 | 5459.118 | 11530.280 |

These are medians of three per-trial measurements, with baseline traffic separated from fault/repair traffic. Latency quantiles are histogram estimates; none are pooled across requests or trials. The sample and shared host do not support significance, a performance ranking, or reliable tail estimates. Ordinary service stabilized before the repaired shard could satisfy the acknowledgement policy. The independent persistence-recovered dimension remained unobserved for these individual-primary faults; there was no all-stopped boundary.

All nine evidence/cleanup checks passed, with zero incomplete or unstarted entries. The 44 saved files were copied byte for byte; comparison JSON SHA-256 is `f17f14c1584530ff5b4ad798c7ce9c3451d68e4575151f8725d3891aa0c56644`. Position balancing reduces ordering confounding without balancing all carryover pairs or proving independent samples. These results establish observed readability and value equality under this fault; future loss-free failover, power-loss survival and strong consistency remain unsupported.

## Managed storage and cleanup

Volumes have exact run/node labels, fixed names, local drivers, and validated `/data` mounts. Containers retain read-only root filesystems and no host bind mounts. Finite workload/probe deadlines, bounded payload/write counts, Redis `maxmemory 128mb`, and `auto-aof-rewrite-percentage 0` limit generated traffic and disable growth-triggered automatic AOF rewrite. That setting does not prohibit Redis from scheduling AOF initialization or rewrite during startup and full synchronization; actual rewrite status remains part of the observations. The volume filesystem has no hard quota; allocated storage and file listings are explicit observations.

Failures and cancellation still attempt cleanup with a fresh bounded context. Container and volume ownership are revalidated; altered or foreign resources are preserved and reported. Containers are removed before their tracked volumes, the empty bridge is removed, and absence is verified. A physically applied restart with a command error remains an error with its evidence retained. Cleanup failure prevents a recovered outcome. [Retained-resource decision](decisions/0008-retained-identity-and-independent-redundancy.md)

The saved persistence results above establish observed frozen-ledger recovery or loss under this defined process fault. No host/storage-failure comparison or general durability guarantee is claimed.
