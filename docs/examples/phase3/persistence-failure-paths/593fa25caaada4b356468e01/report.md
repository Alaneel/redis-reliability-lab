# Redis Reliability Lab — primary-failure

Run: `593fa25caaada4b356468e01`  
Outcome: **aborted**  
Started: 2026-10-11T16:48:01.490477+08:00  
Finished: 2026-10-11T16:48:30.352121+08:00

## Recovery

| Condition | Observed |
|---|---|
| Original six-node topology recovered | false |
| Post-validation diagnostic service sustained | false |
| Topology and diagnostic service jointly sustained | false |
| Frozen acknowledged writes verified | false |
| All managed processes restarted | false |
| Original replication redundancy restored | false |
| Restoration deadline expired | false |
| Diagnostic deadline expired | false |
| Cleanup verified | true |


## Frozen-ledger process restart

Configured persistence mode: `aof-always`. Failure model: `sequential SIGKILL of all six Redis processes, same-container restart; host and storage remain running`.

Workload traffic stops and drains before the sequential SIGKILL actions. No workload requests are offered during fault or restart observation; these phases do not measure zero availability or recovery throughput. Read-only frozen-ledger validation precedes separate diagnostic writes.

Frozen acknowledged SETs: 432. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:48:23.413958+08:00 |
| Immutable ledger frozen | 2026-10-11T16:48:23.414141+08:00 |
| All six stopped verified before restart | 2026-10-11T16:48:25.654707+08:00 |
| First restart requested | 2026-10-11T16:48:26.656383+08:00 |
| Restart sequence completed | unobserved |
| Frozen-ledger validation completed | unobserved |
| Post-validation diagnostics requested | unobserved |

Drain to first kill request: 0.210 ms (runtime monotonic clock).

Drain to last kill completion: 1520.052 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.667 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | unobserved |
| Original topology sustained | unobserved |
| Replication redundancy sustained | unobserved |
| Post-validation diagnostic service sustained | unobserved |
| AOF frozen-value recovery observed | unobserved |

AOF frozen-value recovery observed: unobserved. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation unobserved; 432 acknowledged values remain unchecked.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:48:23.414168+08:00 | 2026-10-11T16:48:23.676605+08:00 |  |
| 2 | node-5 | 2026-10-11T16:48:23.676621+08:00 | 2026-10-11T16:48:23.940842+08:00 |  |
| 3 | node-6 | 2026-10-11T16:48:23.940868+08:00 | 2026-10-11T16:48:24.181095+08:00 |  |
| 4 | node-1 | 2026-10-11T16:48:24.181104+08:00 | 2026-10-11T16:48:24.434797+08:00 |  |
| 5 | node-2 | 2026-10-11T16:48:24.434806+08:00 | 2026-10-11T16:48:24.690528+08:00 |  |
| 6 | node-3 | 2026-10-11T16:48:24.690541+08:00 | 2026-10-11T16:48:24.934022+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:48:26.656383+08:00 | unobserved | context canceled |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | c7fb1d8a86cf8a06a6a353af3fdb60c37873efa0 | 2026-10-11T16:48:25.064309+08:00 | false | no | managed Redis process is stopped |
| node-5 | 6ef8d528e2316bc6b52f9ee6702aa24daa4e52bb | 2026-10-11T16:48:25.184419+08:00 | false | no | managed Redis process is stopped |
| node-6 | 84d23b850205504d43981af44cef1d156bde7df4 | 2026-10-11T16:48:25.301951+08:00 | false | no | managed Redis process is stopped |
| node-1 | 8a5b3895777511c0a03a54946a3a96d04332f0f4 | 2026-10-11T16:48:25.420831+08:00 | false | no | managed Redis process is stopped |
| node-2 | 6c5b34a2880817b78d2ff22979e4d52e861fdb5d | 2026-10-11T16:48:25.536667+08:00 | false | no | managed Redis process is stopped |
| node-3 | f2d82dc382633e6694c57bf11e75535a20fa49e9 | 2026-10-11T16:48:25.654675+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `8a5b3895777511c0a03a54946a3a96d04332f0f4`. Process run ID: `272d330f2dca7135420f991de49741d19150bc27`. Observed: 2026-10-11T16:48:14.043823+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `8a5b3895777511c0a03a54946a3a96d04332f0f4`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
maxmemory-policy
noeviction
appendonly
yes
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
rdb-del-sync-files
yes
maxmemory
134217728
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
appendfsync
always
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
rdb_last_save_time:1791708484
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:532480
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

Original Redis Cluster ID: `6c5b34a2880817b78d2ff22979e4d52e861fdb5d`. Process run ID: `761e474d6c4b4153a1c08dcbccbc5163c583e25d`. Observed: 2026-10-11T16:48:15.082097+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6c5b34a2880817b78d2ff22979e4d52e861fdb5d`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
save

aof-use-rdb-preamble
yes
appendonly
yes
appendfsync
always
repl-diskless-load
swapdb
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
repl-diskless-sync
yes
maxmemory
134217728
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708484
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:520192
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

Original Redis Cluster ID: `f2d82dc382633e6694c57bf11e75535a20fa49e9`. Process run ID: `fcd4107e6a5c7656504509ac9a6e1a875bb1539e`. Observed: 2026-10-11T16:48:16.098395+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f2d82dc382633e6694c57bf11e75535a20fa49e9`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
appendfsync
always
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
aof-load-truncated
yes
appendonly
yes
repl-diskless-load
swapdb
maxmemory-policy
noeviction
maxmemory
134217728
aof-use-rdb-preamble
yes
save

auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791708485
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

### Prepared node-4

Original Redis Cluster ID: `c7fb1d8a86cf8a06a6a353af3fdb60c37873efa0`. Process run ID: `f856d900cd99a43d2b53f0e59838f0c97dacac65`. Observed: 2026-10-11T16:48:17.131123+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c7fb1d8a86cf8a06a6a353af3fdb60c37873efa0`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
save

appendonly
yes
maxmemory
134217728
aof-use-rdb-preamble
yes
aof-load-truncated
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
repl-diskless-sync
yes
appendfsync
always
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
rdb_last_save_time:1791708485
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
aof_last_cow_size:761856
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

Original Redis Cluster ID: `6ef8d528e2316bc6b52f9ee6702aa24daa4e52bb`. Process run ID: `c6c5b27b7cc475b5aaa41e54d5d1f6a1b530c02c`. Observed: 2026-10-11T16:48:18.162849+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6ef8d528e2316bc6b52f9ee6702aa24daa4e52bb`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
appendfsync
always
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
appendonly
yes
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
maxmemory
134217728
repl-diskless-load
swapdb
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708485
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

### Prepared node-6

Original Redis Cluster ID: `84d23b850205504d43981af44cef1d156bde7df4`. Process run ID: `3c26a3d0f493daa056eeda9bc886c6b1e5d4fee5`. Observed: 2026-10-11T16:48:19.271138+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `84d23b850205504d43981af44cef1d156bde7df4`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
appendfsync
always
rdb-del-sync-files
yes
repl-diskless-load
swapdb
aof-load-truncated
yes
appendonly
yes
save

aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
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
rdb_last_save_time:1791708486
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

### Restart receipt 1: node-1

Original Redis Cluster ID: `8a5b3895777511c0a03a54946a3a96d04332f0f4`. Observed after restart: 2026-10-11T16:48:28.181535+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `8a5b3895777511c0a03a54946a3a96d04332f0f4`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
appendfsync
always
aof-load-truncated
yes
appendonly
yes
repl-diskless-load
swapdb
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
maxmemory
134217728
save

no-appendfsync-on-rewrite
no
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
rdb_changes_since_last_save:144
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708506
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
aof_current_size:31629
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=144,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

Post-synchronization administrative runtime observations: unobserved.

This is a sequential process crash and same-container restart with retained managed volumes. It does not remove Docker VM/host page cache or simulate power loss, storage loss, or a graceful shutdown. Draining and sequential kill/inspection time allow background fsync to progress, so everysec preservation does not establish a one-second-loss guarantee or equivalence to always.

Persistence experiment error: restart node-1: context canceled

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 2240.526 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 828 | 100.00% | 0 | 199.89 | 0.847 / 1.449 / 2.360 |

Acknowledged writes: 432. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 432 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 432 | 1.081 / 1.677 / 2.478 |
| baseline | set | 432 | 1.030 / 1.677 / 2.478 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 432 | 104.291 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 828; started 828; completed 828; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: false; expected: 0; checked: 0; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: false.
Validation was not attempted; 432 acknowledged writes remain unchecked.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Errors

- restart node-1: context canceled

## Measurement methodology

- Quiet administrative process, CONFIG, persistence INFO, keyspace and bounded data-file preparation occurs before baseline traffic; it issues no SAVE, BGSAVE, BGREWRITEAOF or explicit durability barrier.
- The workload fully stops and joins in-flight acknowledgements before freezing its immutable ledger and counters. All later topology observations are read-only until every frozen acknowledged value has been checked. SET-error outcomes remain ambiguous, excluded from acknowledgement guarantees.
- Original replicas are SIGKILLed before original primaries. Every fixed owned container must produce ErrProcessStopped administrative evidence with restart policy no before any restart. Original primaries restart before replicas in the same containers and retained volumes; original Redis Cluster identities and slot ownership must survive.
- The model is sequential process SIGKILL while Docker VM, host and storage remain running, followed by same-container starts. It does not simulate power loss, host/page-cache loss or disk loss. Drain, kill and inspection intervals are reported because they permit background fsync progress.
- Topology restoration requires all six original identities, endpoints, primary slot owners and replica assignments synchronized for stable_for, without forced Redis topology commands. Physical receipts retain early raw evidence, which can be transient during full sync. Six fresh quiet administrative runtime/configuration/file observations after topology stability prove the final persistence mode on those same restarted processes, within the shared repair budget and before frozen-value validation. Samples crossing restoration or diagnostic deadlines are retained but excluded from successful milestones.
- Frozen-ledger GET validation happens exactly once before diagnostic SET/GET or configured replica-acknowledgement probes. Process/slot readiness and later diagnostic service are independent from frozen-value preservation; missing or corrupt values retain a partial data verdict even after service recovers.
- Memory mode has RDB/AOF disabled, rdb-del-sync-files=yes to remove replication-created RDB files, and observed absence of value persistence files. Its persistence-recovery event is not applicable and remains absent; its independent persistence dimension is false. AOF modes use the RDB base preamble required by the pinned image's bundled modules plus incremental AOF records. AOF recovery requires actual matching healthy runtime configuration, loading=0 and complete frozen acknowledged-value preservation; this observation establishes no loss-free or power-loss guarantee.
- No workload demand is offered after draining, so absent fault/recovery traffic does not measure zero availability or throughput. Post-validation probes are separate diagnostics and never enter workload metrics or the frozen ledger. Poll intervals follow sample completion; observations are sampled and non-atomic.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `5f3362a0bdd93bf4a884c5958983a2e641eef0c9e963ba05b2dc59469d6b507f`.
