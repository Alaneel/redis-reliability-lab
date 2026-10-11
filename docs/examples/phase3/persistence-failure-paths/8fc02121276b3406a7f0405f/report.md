# Redis Reliability Lab — primary-failure

Run: `8fc02121276b3406a7f0405f`  
Outcome: **failed**  
Started: 2026-10-11T16:47:33.272241+08:00  
Finished: 2026-10-11T16:48:00.671262+08:00

## Recovery

| Condition | Observed |
|---|---|
| Original six-node topology recovered | false |
| Post-validation diagnostic service sustained | false |
| Topology and diagnostic service jointly sustained | false |
| Frozen acknowledged writes verified | false |
| All managed processes restarted | false |
| Original replication redundancy restored | false |
| Restoration deadline expired | true |
| Diagnostic deadline expired | false |
| Cleanup verified | true |


## Frozen-ledger process restart

Configured persistence mode: `aof-everysec`. Failure model: `sequential SIGKILL of all six Redis processes, same-container restart; host and storage remain running`.

Workload traffic stops and drains before the sequential SIGKILL actions. No workload requests are offered during fault or restart observation; these phases do not measure zero availability or recovery throughput. Read-only frozen-ledger validation precedes separate diagnostic writes.

Frozen acknowledged SETs: 433. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: true.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:47:55.335045+08:00 |
| Immutable ledger frozen | 2026-10-11T16:47:55.335214+08:00 |
| All six stopped verified before restart | 2026-10-11T16:47:57.610753+08:00 |
| First restart requested | 2026-10-11T16:47:58.612375+08:00 |
| Restart sequence completed | unobserved |
| Frozen-ledger validation completed | unobserved |
| Post-validation diagnostics requested | unobserved |

Drain to first kill request: 0.202 ms (runtime monotonic clock).

Drain to last kill completion: 1546.182 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.615 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | unobserved |
| Original topology sustained | unobserved |
| Replication redundancy sustained | unobserved |
| Post-validation diagnostic service sustained | unobserved |
| AOF frozen-value recovery observed | unobserved |

AOF frozen-value recovery observed: unobserved. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation unobserved; 433 acknowledged values remain unchecked.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:47:55.335247+08:00 | 2026-10-11T16:47:55.574173+08:00 |  |
| 2 | node-5 | 2026-10-11T16:47:55.57419+08:00 | 2026-10-11T16:47:55.842617+08:00 |  |
| 3 | node-6 | 2026-10-11T16:47:55.842634+08:00 | 2026-10-11T16:47:56.096914+08:00 |  |
| 4 | node-1 | 2026-10-11T16:47:56.096926+08:00 | 2026-10-11T16:47:56.356544+08:00 |  |
| 5 | node-2 | 2026-10-11T16:47:56.356551+08:00 | 2026-10-11T16:47:56.618831+08:00 |  |
| 6 | node-3 | 2026-10-11T16:47:56.618846+08:00 | 2026-10-11T16:47:56.881239+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:47:58.612375+08:00 | unobserved | context deadline exceeded |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 8dcdb210959acac5f748d695473376036cda738d | 2026-10-11T16:47:57.004817+08:00 | false | no | managed Redis process is stopped |
| node-5 | 1f06dc3acb2587c1f455487eb52ea9cd250be1c1 | 2026-10-11T16:47:57.123618+08:00 | false | no | managed Redis process is stopped |
| node-6 | e48f40e1a4e91c57a02dd39b6d4e991c57c7c2d0 | 2026-10-11T16:47:57.240245+08:00 | false | no | managed Redis process is stopped |
| node-1 | 5e9f1dcef0a4faa0d164df27573ddaf2862cc723 | 2026-10-11T16:47:57.355686+08:00 | false | no | managed Redis process is stopped |
| node-2 | bba057a517663973cffbad6a736bd7a9dcdc3724 | 2026-10-11T16:47:57.477586+08:00 | false | no | managed Redis process is stopped |
| node-3 | f1b1cd60090be3114a4cb9d7d40ee05c5584e60a | 2026-10-11T16:47:57.610716+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `5e9f1dcef0a4faa0d164df27573ddaf2862cc723`. Process run ID: `9a6317de6ae38c7c4746add5606d0eefbbf1a921`. Observed: 2026-10-11T16:47:46.16337+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `5e9f1dcef0a4faa0d164df27573ddaf2862cc723`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
repl-diskless-load
swapdb
maxmemory
134217728
repl-diskless-sync
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
save

appendonly
yes
aof-load-truncated
yes
```

Observed INFO persistence:

```text
# Persistence
loading:0
async_loading:0
current_cow_peak:0
current_cow_size:0
current_cow_size_age:0
current_fork_perc:0.00
current_save_keys_processed:0
current_save_keys_total:0
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708456
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:471040
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:-1
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:0
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:0
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:0
aof_base_size:0
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Prepared node-2

Original Redis Cluster ID: `bba057a517663973cffbad6a736bd7a9dcdc3724`. Process run ID: `5a68c525ffdb791df79dc39b26128054330f9891`. Observed: 2026-10-11T16:47:47.189847+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `bba057a517663973cffbad6a736bd7a9dcdc3724`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
repl-diskless-load
swapdb
maxmemory-policy
noeviction
save

appendonly
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
aof-load-truncated
yes
maxmemory
134217728
```

Observed INFO persistence:

```text
# Persistence
loading:0
async_loading:0
current_cow_peak:0
current_cow_size:0
current_cow_size_age:0
current_fork_perc:0.00
current_save_keys_processed:0
current_save_keys_total:0
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708456
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:446464
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:-1
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:0
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:0
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:0
aof_base_size:0
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Prepared node-3

Original Redis Cluster ID: `f1b1cd60090be3114a4cb9d7d40ee05c5584e60a`. Process run ID: `e044d9404687b75767915b222a9827a9aa0db96c`. Observed: 2026-10-11T16:47:48.190108+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f1b1cd60090be3114a4cb9d7d40ee05c5584e60a`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
appendonly
yes
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
appendfsync
everysec
repl-diskless-load
swapdb
repl-diskless-sync
yes
rdb-del-sync-files
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
yes
save

```

Observed INFO persistence:

```text
# Persistence
loading:0
async_loading:0
current_cow_peak:0
current_cow_size:0
current_cow_size_age:0
current_fork_perc:0.00
current_save_keys_processed:0
current_save_keys_total:0
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708457
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:540672
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:-1
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:0
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:0
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:0
aof_base_size:0
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Prepared node-4

Original Redis Cluster ID: `8dcdb210959acac5f748d695473376036cda738d`. Process run ID: `3e76743313e85d4b08b72763bed306ff52bf5977`. Observed: 2026-10-11T16:47:49.197987+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `8dcdb210959acac5f748d695473376036cda738d`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
maxmemory
134217728
appendfsync
everysec
repl-diskless-sync
yes
aof-load-truncated
yes
appendonly
yes
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
maxmemory-policy
noeviction
save

```

Observed INFO persistence:

```text
# Persistence
loading:0
async_loading:0
current_cow_peak:0
current_cow_size:0
current_cow_size_age:0
current_fork_perc:0.00
current_save_keys_processed:0
current_save_keys_total:0
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708457
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:389120
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:108
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Prepared node-5

Original Redis Cluster ID: `1f06dc3acb2587c1f455487eb52ea9cd250be1c1`. Process run ID: `7b16aac2c50dd04bcc45287b97a4fe6e9d52d43d`. Observed: 2026-10-11T16:47:50.19316+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1f06dc3acb2587c1f455487eb52ea9cd250be1c1`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
maxmemory
134217728
appendfsync
everysec
aof-use-rdb-preamble
yes
save

no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
repl-diskless-load
swapdb
appendonly
yes
```

Observed INFO persistence:

```text
# Persistence
loading:0
async_loading:0
current_cow_peak:0
current_cow_size:0
current_cow_size_age:0
current_fork_perc:0.00
current_save_keys_processed:0
current_save_keys_total:0
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708457
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:770048
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:108
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Prepared node-6

Original Redis Cluster ID: `e48f40e1a4e91c57a02dd39b6d4e991c57c7c2d0`. Process run ID: `35b8bf77f85d4cc70cd5dd550198d49a3d23f0a9`. Observed: 2026-10-11T16:47:51.183264+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e48f40e1a4e91c57a02dd39b6d4e991c57c7c2d0`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendonly
yes
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
repl-diskless-load
swapdb
appendfsync
everysec
aof-load-truncated
yes
repl-diskless-sync
yes
maxmemory
134217728
save

```

Observed INFO persistence:

```text
# Persistence
loading:0
async_loading:0
current_cow_peak:0
current_cow_size:0
current_cow_size_age:0
current_fork_perc:0.00
current_save_keys_processed:0
current_save_keys_total:0
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708458
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:393216
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:108
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

Post-synchronization administrative runtime observations: unobserved.

This is a sequential process crash and same-container restart with retained managed volumes. It does not remove Docker VM/host page cache or simulate power loss, storage loss, or a graceful shutdown. Draining and sequential kill/inspection time allow background fsync to progress, so everysec preservation does not establish a one-second-loss guarantee or equivalence to always.

Persistence experiment error: restart node-1: context deadline exceeded

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 2275.492 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 830 | 100.00% | 0 | 199.94 | 0.632 / 0.847 / 1.252 |

Acknowledged writes: 433. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 433 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 433 | 0.632 / 0.890 / 1.252 |
| baseline | set | 433 | 0.602 / 0.847 / 1.192 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 433 | 104.306 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 830; started 830; completed 830; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: false; expected: 0; checked: 0; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: false.
Validation was not attempted; 433 acknowledged writes remain unchecked.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Errors

- restart node-1: context deadline exceeded

## Measurement methodology

- Quiet administrative process, CONFIG, persistence INFO, keyspace and bounded data-file preparation occurs before baseline traffic; it issues no SAVE, BGSAVE, BGREWRITEAOF or explicit durability barrier.
- The workload fully stops and joins in-flight acknowledgements before freezing its immutable ledger and counters. All later topology observations are read-only until every frozen acknowledged value has been checked. SET-error outcomes remain ambiguous, excluded from acknowledgement guarantees.
- Original replicas are SIGKILLed before original primaries. Every fixed owned container must produce ErrProcessStopped administrative evidence with restart policy no before any restart. Original primaries restart before replicas in the same containers and retained volumes; original Redis Cluster identities and slot ownership must survive.
- The model is sequential process SIGKILL while Docker VM, host and storage remain running, followed by same-container starts. It does not simulate power loss, host/page-cache loss or disk loss. Drain, kill and inspection intervals are reported because they permit background fsync progress.
- Topology restoration requires all six original identities, endpoints, primary slot owners and replica assignments synchronized for stable_for, without forced Redis topology commands. Physical receipts retain early raw evidence, which can be transient during full sync. Six fresh quiet administrative runtime/configuration/file observations after topology stability prove the final persistence mode on those same restarted processes, within the shared repair budget and before frozen-value validation. Samples crossing restoration or diagnostic deadlines are retained but excluded from successful milestones.
- Frozen-ledger GET validation happens exactly once before diagnostic SET/GET or configured replica-acknowledgement probes. Process/slot readiness and later diagnostic service are independent from frozen-value preservation; missing or corrupt values retain a partial data verdict even after service recovers.
- Memory mode has RDB/AOF disabled, rdb-del-sync-files=yes to remove replication-created RDB files, and observed absence of value persistence files. Its persistence-recovery event is not applicable and remains absent; its independent persistence dimension is false. AOF modes use the RDB base preamble required by the pinned image's bundled modules plus incremental AOF records. AOF recovery requires actual matching healthy runtime configuration, loading=0 and complete frozen acknowledged-value preservation; this observation establishes no loss-free or power-loss guarantee.
- No workload demand is offered after draining, so absent fault/recovery traffic does not measure zero availability or throughput. Post-validation probes are separate diagnostics and never enter workload metrics or the frozen ledger. Poll intervals follow sample completion; observations are sampled and non-atomic.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `79160ee95fc0aad724fab7a2c6b4e9f4d1af92d59da547583ee75842d0f8ca81`.
