# Redis Reliability Lab — dataset-trial

Run: `dd88217ca44abba63ff3ccb6`  
Outcome: **recovered**  
Started: 2026-10-11T19:51:10.115831+08:00  
Finished: 2026-10-11T19:52:06.395289+08:00

## Recovery

| Condition | Observed |
|---|---|
| Original six-node topology recovered | true |
| Post-validation diagnostic service sustained | true |
| Topology and diagnostic service jointly sustained | true |
| Frozen acknowledged writes verified | true |
| All managed processes restarted | true |
| Original replication redundancy restored | true |
| Restoration deadline expired | false |
| Diagnostic deadline expired | false |
| Cleanup verified | true |


## Frozen-ledger process restart

Configured persistence mode: `aof-always`. Failure model: `sequential SIGKILL of all six Redis processes, same-container restart; host and storage remain running`.

Workload traffic stops and drains before the sequential SIGKILL actions. No workload requests are offered during fault or restart observation; these phases do not measure zero availability or recovery throughput. Read-only frozen-ledger validation precedes separate diagnostic writes.

Frozen acknowledged SETs: 179. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T19:51:33.601617+08:00 |
| Immutable ledger frozen | 2026-10-11T19:51:33.601743+08:00 |
| All six stopped verified before restart | 2026-10-11T19:51:35.95823+08:00 |
| First restart requested | 2026-10-11T19:51:39.960536+08:00 |
| Restart sequence completed | 2026-10-11T19:51:49.89592+08:00 |
| Frozen-ledger validation completed | 2026-10-11T19:52:01.159763+08:00 |
| Post-validation diagnostics requested | 2026-10-11T19:52:01.160093+08:00 |

Drain to first kill request: 0.162 ms (runtime monotonic clock).

Drain to last kill completion: 1571.667 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 4002.311 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 19909.041 |
| Original topology sustained | 21097.224 |
| Replication redundancy sustained | 21097.224 |
| Post-validation diagnostic service sustained | 28781.659 |
| AOF frozen-value recovery observed | 27558.023 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 179; checked 179; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T19:51:33.601778+08:00 | 2026-10-11T19:51:33.859829+08:00 |  |
| 2 | node-5 | 2026-10-11T19:51:33.859847+08:00 | 2026-10-11T19:51:34.142209+08:00 |  |
| 3 | node-6 | 2026-10-11T19:51:34.142227+08:00 | 2026-10-11T19:51:34.396485+08:00 |  |
| 4 | node-1 | 2026-10-11T19:51:34.396499+08:00 | 2026-10-11T19:51:34.658263+08:00 |  |
| 5 | node-2 | 2026-10-11T19:51:34.658271+08:00 | 2026-10-11T19:51:34.923374+08:00 |  |
| 6 | node-3 | 2026-10-11T19:51:34.923387+08:00 | 2026-10-11T19:51:35.173282+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T19:51:39.960536+08:00 | 2026-10-11T19:51:41.640329+08:00 |  |
| 2 | node-2 | 2026-10-11T19:51:41.640339+08:00 | 2026-10-11T19:51:43.304582+08:00 |  |
| 3 | node-3 | 2026-10-11T19:51:43.304595+08:00 | 2026-10-11T19:51:44.978367+08:00 |  |
| 4 | node-4 | 2026-10-11T19:51:44.978378+08:00 | 2026-10-11T19:51:46.627979+08:00 |  |
| 5 | node-5 | 2026-10-11T19:51:46.627982+08:00 | 2026-10-11T19:51:48.306134+08:00 |  |
| 6 | node-6 | 2026-10-11T19:51:48.306144+08:00 | 2026-10-11T19:51:49.895918+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | f82206901ddc3e53f8d1f3102583b6b09c5bd5a5 | 2026-10-11T19:51:35.311964+08:00 | false | no | managed Redis process is stopped |
| node-5 | 2c93c2f0924ede03bcdb930a51ad423d766fd96b | 2026-10-11T19:51:35.44156+08:00 | false | no | managed Redis process is stopped |
| node-6 | 70f581b3e04404f9ac902671fb28dba270482430 | 2026-10-11T19:51:35.571785+08:00 | false | no | managed Redis process is stopped |
| node-1 | a2940fba25ad5ccece469610654aaadaac1459cb | 2026-10-11T19:51:35.7026+08:00 | false | no | managed Redis process is stopped |
| node-2 | 75b3e52641107447681347766872db10a3f12e5a | 2026-10-11T19:51:35.830228+08:00 | false | no | managed Redis process is stopped |
| node-3 | b6bd377d90a667f0a9f283d1d8826567473daccc | 2026-10-11T19:51:35.958194+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `a2940fba25ad5ccece469610654aaadaac1459cb`. Process run ID: `afe43e37e4b8fde209cc06b73dfea84ea9b75590`. Observed: 2026-10-11T19:51:26.065456+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `a2940fba25ad5ccece469610654aaadaac1459cb`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
appendfsync
always
maxmemory
134217728
repl-diskless-sync
yes
rdb-del-sync-files
yes
save

appendonly
yes
aof-use-rdb-preamble
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
rdb_last_save_time:1791719473
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:561152
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

Original Redis Cluster ID: `75b3e52641107447681347766872db10a3f12e5a`. Process run ID: `a8f73201cad89b6c4bb69b9ff32bd497fb24a49a`. Observed: 2026-10-11T19:51:27.218538+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `75b3e52641107447681347766872db10a3f12e5a`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
aof-load-truncated
yes
appendfsync
always
maxmemory
134217728
appendonly
yes
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
rdb-del-sync-files
yes
repl-diskless-sync
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
rdb_last_save_time:1791719474
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:569344
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

Original Redis Cluster ID: `b6bd377d90a667f0a9f283d1d8826567473daccc`. Process run ID: `58e4fb9d93d79d04e206613c1a7a8705531a9f7b`. Observed: 2026-10-11T19:51:28.301216+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b6bd377d90a667f0a9f283d1d8826567473daccc`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
repl-diskless-load
swapdb
maxmemory
134217728
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
save

aof-load-truncated
yes
appendonly
yes
rdb-del-sync-files
yes
maxmemory-policy
noeviction
appendfsync
always
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
rdb_last_save_time:1791719474
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:536576
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

Original Redis Cluster ID: `f82206901ddc3e53f8d1f3102583b6b09c5bd5a5`. Process run ID: `d23e8d62f593c5c30606480a68c2619872a07f27`. Observed: 2026-10-11T19:51:29.33719+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f82206901ddc3e53f8d1f3102583b6b09c5bd5a5`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
appendfsync
always
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
repl-diskless-sync
yes
save

appendonly
yes
maxmemory
134217728
repl-diskless-load
swapdb
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
rdb_last_save_time:1791719475
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

### Prepared node-5

Original Redis Cluster ID: `2c93c2f0924ede03bcdb930a51ad423d766fd96b`. Process run ID: `3467dc8db14c64e9c8011cf115d3f7074604f80f`. Observed: 2026-10-11T19:51:30.375797+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2c93c2f0924ede03bcdb930a51ad423d766fd96b`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
maxmemory
134217728
maxmemory-policy
noeviction
appendonly
yes
repl-diskless-load
swapdb
rdb-del-sync-files
yes
repl-diskless-sync
yes
save

aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
appendfsync
always
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
rdb_last_save_time:1791719475
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
aof_last_cow_size:438272
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

Original Redis Cluster ID: `70f581b3e04404f9ac902671fb28dba270482430`. Process run ID: `56397b0e0e5fba317d6c28cb3985a28f7584fba8`. Observed: 2026-10-11T19:51:31.44423+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `70f581b3e04404f9ac902671fb28dba270482430`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
rdb-del-sync-files
yes
aof-load-truncated
yes
repl-diskless-sync
yes
save

appendfsync
always
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791719476
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
aof_last_cow_size:454656
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

Original Redis Cluster ID: `a2940fba25ad5ccece469610654aaadaac1459cb`. Observed after restart: 2026-10-11T19:51:41.517161+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `a2940fba25ad5ccece469610654aaadaac1459cb`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
maxmemory-policy
noeviction
aof-load-truncated
yes
repl-diskless-sync
yes
appendfsync
always
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
save

appendonly
yes
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719500
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
aof_current_size:13517
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=61,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 2: node-2

Original Redis Cluster ID: `75b3e52641107447681347766872db10a3f12e5a`. Observed after restart: 2026-10-11T19:51:43.174943+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `75b3e52641107447681347766872db10a3f12e5a`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
save

no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
appendfsync
always
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
maxmemory
134217728
appendonly
yes
aof-load-truncated
yes
rdb-del-sync-files
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719501
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
aof_current_size:13519
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=61,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 3: node-3

Original Redis Cluster ID: `b6bd377d90a667f0a9f283d1d8826567473daccc`. Observed after restart: 2026-10-11T19:51:44.85396+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b6bd377d90a667f0a9f283d1d8826567473daccc`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
appendonly
yes
save

aof-load-truncated
yes
maxmemory
134217728
repl-diskless-sync
yes
appendfsync
always
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
repl-diskless-load
swapdb
aof-use-rdb-preamble
yes
rdb-del-sync-files
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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719503
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
aof_current_size:13297
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=60,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 4: node-4

Original Redis Cluster ID: `f82206901ddc3e53f8d1f3102583b6b09c5bd5a5`. Observed after restart: 2026-10-11T19:51:46.510006+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f82206901ddc3e53f8d1f3102583b6b09c5bd5a5`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
appendonly
yes
repl-diskless-load
swapdb
repl-diskless-sync
yes
maxmemory-policy
noeviction
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
save

maxmemory
134217728
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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719505
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
aof_current_size:13297
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=60,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/temp-1791719505.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `2c93c2f0924ede03bcdb930a51ad423d766fd96b`. Observed after restart: 2026-10-11T19:51:48.180725+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2c93c2f0924ede03bcdb930a51ad423d766fd96b`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
save

repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
maxmemory
134217728
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
appendfsync
always
repl-diskless-sync
yes
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719506
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
aof_current_size:13517
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=61,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/temp-1791719506.1.rdb
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `70f581b3e04404f9ac902671fb28dba270482430`. Observed after restart: 2026-10-11T19:51:49.773922+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `70f581b3e04404f9ac902671fb28dba270482430`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
repl-diskless-sync
yes
rdb-del-sync-files
yes
save

appendonly
yes
aof-load-truncated
yes
appendfsync
always
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719508
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
aof_current_size:13519
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=61,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/temp-1791719508.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `a2940fba25ad5ccece469610654aaadaac1459cb`. Process run ID: `140eac3434f36829bc5ff042c3bfd9b4b03b5080`. Observed: 2026-10-11T19:51:56.012726+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `a2940fba25ad5ccece469610654aaadaac1459cb`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
maxmemory-policy
noeviction
aof-load-truncated
yes
repl-diskless-sync
yes
appendfsync
always
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
save

appendonly
yes
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719500
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:831488
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
aof_current_size:13517
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=61,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-2

Original Redis Cluster ID: `75b3e52641107447681347766872db10a3f12e5a`. Process run ID: `ad84833d77270178556406c75fcb731ce325e431`. Observed: 2026-10-11T19:51:57.025096+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `75b3e52641107447681347766872db10a3f12e5a`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
save

no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
appendfsync
always
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
maxmemory
134217728
appendonly
yes
aof-load-truncated
yes
rdb-del-sync-files
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719501
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:884736
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
aof_current_size:13519
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=61,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-3

Original Redis Cluster ID: `b6bd377d90a667f0a9f283d1d8826567473daccc`. Process run ID: `722a8cc2952d87ea5a897a15ca8cc20d83c739dd`. Observed: 2026-10-11T19:51:58.022049+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b6bd377d90a667f0a9f283d1d8826567473daccc`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
appendonly
yes
save

aof-load-truncated
yes
maxmemory
134217728
repl-diskless-sync
yes
appendfsync
always
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
repl-diskless-load
swapdb
aof-use-rdb-preamble
yes
rdb-del-sync-files
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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719503
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:876544
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
aof_current_size:13297
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=60,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-4

Original Redis Cluster ID: `f82206901ddc3e53f8d1f3102583b6b09c5bd5a5`. Process run ID: `40df2dc6a1da779e058985cba3b2e3c6c2abf963`. Observed: 2026-10-11T19:51:59.015596+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f82206901ddc3e53f8d1f3102583b6b09c5bd5a5`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
appendonly
yes
repl-diskless-load
swapdb
repl-diskless-sync
yes
maxmemory-policy
noeviction
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
save

maxmemory
134217728
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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719505
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:60
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:700416
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:8875
aof_base_size:8875
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=60,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.3.incr.aof
/data/appendonlydir/appendonly.aof.3.base.rdb

```

#### Restored node-5

Original Redis Cluster ID: `2c93c2f0924ede03bcdb930a51ad423d766fd96b`. Process run ID: `3ce1b6deb6968587c1016c35b1b9d425d2961e2f`. Observed: 2026-10-11T19:52:00.096454+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2c93c2f0924ede03bcdb930a51ad423d766fd96b`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
save

repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
maxmemory
134217728
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
appendfsync
always
repl-diskless-sync
yes
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719506
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:61
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:688128
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:9025
aof_base_size:9025
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=61,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.3.incr.aof
/data/appendonlydir/appendonly.aof.3.base.rdb

```

#### Restored node-6

Original Redis Cluster ID: `70f581b3e04404f9ac902671fb28dba270482430`. Process run ID: `7fee120b108a8e3e1ac5c21203721768caf8b158`. Observed: 2026-10-11T19:52:01.098322+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `70f581b3e04404f9ac902671fb28dba270482430`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
repl-diskless-sync
yes
rdb-del-sync-files
yes
save

appendonly
yes
aof-load-truncated
yes
appendfsync
always
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719508
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:61
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:688128
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:9020
aof_base_size:9020
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=61,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.3.incr.aof
/data/appendonlydir/appendonly.aof.3.base.rdb

```

This is a sequential process crash and same-container restart with retained managed volumes. It does not remove Docker VM/host page cache or simulate power loss, storage loss, or a graceful shutdown. Draining and sequential kill/inspection time allow background fsync to progress, so everysec preservation does not establish a one-second-loss guarantee or equivalence to always.

## Independent recovery dimensions

Target observed role: `unobserved`.

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
| Persistence recovery verified | true |

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 2356.460 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 363 | 100.00% | 0 | 169.54 | 0.602 / 1.314 / 2.602 |

Acknowledged writes: 179. Ambiguous SET outcomes: 0.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | a2940fba25ad5ccece469610654aaadaac1459cb | 2026-10-11T19:51:26.065456+08:00 | a2940fba25ad5ccece469610654aaadaac1459cb | true |  |
| node-2 | 75b3e52641107447681347766872db10a3f12e5a | 2026-10-11T19:51:27.218538+08:00 | 75b3e52641107447681347766872db10a3f12e5a | true |  |
| node-3 | b6bd377d90a667f0a9f283d1d8826567473daccc | 2026-10-11T19:51:28.301216+08:00 | b6bd377d90a667f0a9f283d1d8826567473daccc | true |  |
| node-4 | f82206901ddc3e53f8d1f3102583b6b09c5bd5a5 | 2026-10-11T19:51:29.33719+08:00 | f82206901ddc3e53f8d1f3102583b6b09c5bd5a5 | true |  |
| node-5 | 2c93c2f0924ede03bcdb930a51ad423d766fd96b | 2026-10-11T19:51:30.375797+08:00 | 2c93c2f0924ede03bcdb930a51ad423d766fd96b | true |  |
| node-6 | 70f581b3e04404f9ac902671fb28dba270482430 | 2026-10-11T19:51:31.44423+08:00 | 70f581b3e04404f9ac902671fb28dba270482430 | true |  |

Before-baseline `set` policy probe at 2026-10-11T19:51:31.460415+08:00; all shards successful: true.

| Original shard index | SET acknowledged |
|---|---|
| 0 | true |
| 1 | true |
| 2 | true |

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 179 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 179 | 1.030 / 1.521 / 2.732 |
| baseline | set | 179 | 1.030 / 1.521 / 2.725 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 179 | 83.602 |

Global pacing: offered 363; started 363; completed 363; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 179; checked: 179; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Measurement methodology

- Quiet administrative process, CONFIG, persistence INFO, keyspace and bounded data-file preparation occurs before baseline traffic; it issues no SAVE, BGSAVE, BGREWRITEAOF or explicit durability barrier.
- The workload fully stops and joins in-flight acknowledgements before freezing its immutable ledger and counters. All later topology observations are read-only until every frozen acknowledged value has been checked. SET-error outcomes remain ambiguous, excluded from acknowledgement guarantees.
- Original replicas are SIGKILLed before original primaries. Every fixed owned container must produce ErrProcessStopped administrative evidence with restart policy no before any restart. Original primaries restart before replicas in the same containers and retained volumes; original Redis Cluster identities and slot ownership must survive.
- The model is sequential process SIGKILL while Docker VM, host and storage remain running, followed by same-container starts. It does not simulate power loss, host/page-cache loss or disk loss. Drain, kill and inspection intervals are reported because they permit background fsync progress.
- Topology restoration requires all six original identities, endpoints, primary slot owners and replica assignments synchronized for stable_for, without forced Redis topology commands. Physical receipts retain early raw evidence, which can be transient during full sync. Six fresh quiet administrative runtime/configuration/file observations after topology stability prove the final persistence mode on those same restarted processes, within the shared repair budget and before frozen-value validation. Samples crossing restoration or diagnostic deadlines are retained but excluded from successful milestones.
- Frozen-ledger GET validation happens exactly once before diagnostic SET/GET or configured replica-acknowledgement probes. Process/slot readiness and later diagnostic service are independent from frozen-value preservation; missing or corrupt values retain a partial data verdict even after service recovers.
- Memory mode has RDB/AOF disabled, rdb-del-sync-files=yes to remove replication-created RDB files, and observed absence of value persistence files. Its persistence-recovery event is not applicable and remains absent; its independent persistence dimension is false. AOF modes use the RDB base preamble required by the pinned image's bundled modules plus incremental AOF records. AOF recovery requires actual matching healthy runtime configuration, loading=0 and complete frozen acknowledged-value preservation; this observation establishes no loss-free or power-loss guarantee.
- No workload demand is offered after draining, so absent fault/recovery traffic does not measure zero availability or throughput. Post-validation probes are separate diagnostics and never enter workload metrics or the frozen ledger. Poll intervals follow sample completion; observations are sampled and non-atomic.
- All AOF modes and write policies use six-node quiet runtime preparation followed by exactly three all-shard policy SET/GET warmup writes before baseline starts. Warmup writes are excluded from the workload ledger and metrics; no WAITAOF compatibility command is issued after baseline to alter crash-trial durability.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `f2bac3b9d90a5211a7af54d371f6dc0eeffc094f678d5c4b227214e716aa0761`.
