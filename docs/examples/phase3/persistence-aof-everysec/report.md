# Redis Reliability Lab — primary-failure

Run: `8010385e1d747640095e5e8f`  
Outcome: **recovered**  
Started: 2026-10-11T16:44:43.426058+08:00  
Finished: 2026-10-11T16:45:36.135407+08:00

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

Configured persistence mode: `aof-everysec`. Failure model: `sequential SIGKILL of all six Redis processes, same-container restart; host and storage remain running`.

Workload traffic stops and drains before the sequential SIGKILL actions. No workload requests are offered during fault or restart observation; these phases do not measure zero availability or recovery throughput. Read-only frozen-ledger validation precedes separate diagnostic writes.

Frozen acknowledged SETs: 433. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:45:05.262263+08:00 |
| Immutable ledger frozen | 2026-10-11T16:45:05.262405+08:00 |
| All six stopped verified before restart | 2026-10-11T16:45:07.546247+08:00 |
| First restart requested | 2026-10-11T16:45:08.546739+08:00 |
| Restart sequence completed | 2026-10-11T16:45:18.023671+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:45:29.428244+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:45:29.428522+08:00 |

Drain to first kill request: 0.170 ms (runtime monotonic clock).

Drain to last kill completion: 1578.968 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1000.485 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16631.091 |
| Original topology sustained | 17639.774 |
| Replication redundancy sustained | 17639.774 |
| Post-validation diagnostic service sustained | 25624.264 |
| AOF frozen-value recovery observed | 24165.619 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 433; checked 433; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:45:05.262433+08:00 | 2026-10-11T16:45:05.532701+08:00 |  |
| 2 | node-5 | 2026-10-11T16:45:05.532715+08:00 | 2026-10-11T16:45:05.797245+08:00 |  |
| 3 | node-6 | 2026-10-11T16:45:05.797263+08:00 | 2026-10-11T16:45:06.037247+08:00 |  |
| 4 | node-1 | 2026-10-11T16:45:06.037263+08:00 | 2026-10-11T16:45:06.308988+08:00 |  |
| 5 | node-2 | 2026-10-11T16:45:06.309002+08:00 | 2026-10-11T16:45:06.565733+08:00 |  |
| 6 | node-3 | 2026-10-11T16:45:06.56575+08:00 | 2026-10-11T16:45:06.841244+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:45:08.546739+08:00 | 2026-10-11T16:45:10.155748+08:00 |  |
| 2 | node-2 | 2026-10-11T16:45:10.155753+08:00 | 2026-10-11T16:45:11.716082+08:00 |  |
| 3 | node-3 | 2026-10-11T16:45:11.716097+08:00 | 2026-10-11T16:45:13.304803+08:00 |  |
| 4 | node-4 | 2026-10-11T16:45:13.304816+08:00 | 2026-10-11T16:45:14.87054+08:00 |  |
| 5 | node-5 | 2026-10-11T16:45:14.870548+08:00 | 2026-10-11T16:45:16.4245+08:00 |  |
| 6 | node-6 | 2026-10-11T16:45:16.424508+08:00 | 2026-10-11T16:45:18.023667+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | ad8ae597bfb05df43406b5c8c23d7abc4711eecd | 2026-10-11T16:45:06.968098+08:00 | false | no | managed Redis process is stopped |
| node-5 | 79d1519d78125dde931ba5e0b6018eba5d16aa6e | 2026-10-11T16:45:07.084251+08:00 | false | no | managed Redis process is stopped |
| node-6 | ad38f995c42cff19be4f930ff607fbeb9407801b | 2026-10-11T16:45:07.199021+08:00 | false | no | managed Redis process is stopped |
| node-1 | 598e1d4f860a06065d51575c30f30d3fdee9ee24 | 2026-10-11T16:45:07.314603+08:00 | false | no | managed Redis process is stopped |
| node-2 | 17ec5fb9af9141cc36bb60e2be93b1284510a4cb | 2026-10-11T16:45:07.429654+08:00 | false | no | managed Redis process is stopped |
| node-3 | c8322008da4aa9f52cc4fe9a89285a612a172420 | 2026-10-11T16:45:07.546212+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `598e1d4f860a06065d51575c30f30d3fdee9ee24`. Process run ID: `fa52bc81e425a41b96ee56e6d41d865916a0b122`. Observed: 2026-10-11T16:44:56.161915+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `598e1d4f860a06065d51575c30f30d3fdee9ee24`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
save

no-appendfsync-on-rewrite
no
appendonly
yes
aof-use-rdb-preamble
yes
aof-load-truncated
yes
repl-diskless-sync
yes
maxmemory-policy
noeviction
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
maxmemory
134217728
appendfsync
everysec
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
rdb_last_save_time:1791708286
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:557056
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

Original Redis Cluster ID: `17ec5fb9af9141cc36bb60e2be93b1284510a4cb`. Process run ID: `5ff3b4fcd2bf6991bfbbf81002d7315b65d1f756`. Observed: 2026-10-11T16:44:57.15799+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `17ec5fb9af9141cc36bb60e2be93b1284510a4cb`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
appendfsync
everysec
save

aof-load-truncated
yes
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
appendonly
yes
repl-diskless-load
swapdb
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708286
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:466944
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

Original Redis Cluster ID: `c8322008da4aa9f52cc4fe9a89285a612a172420`. Process run ID: `5add454854fe723507ccc03620a7540ea37d3673`. Observed: 2026-10-11T16:44:58.146885+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c8322008da4aa9f52cc4fe9a89285a612a172420`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
maxmemory
134217728
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
appendfsync
everysec
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
save

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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708287
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:479232
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

Original Redis Cluster ID: `ad8ae597bfb05df43406b5c8c23d7abc4711eecd`. Process run ID: `dabc89043bb83c44c9f7eb1412f332aa1cc57d6c`. Observed: 2026-10-11T16:44:59.137828+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ad8ae597bfb05df43406b5c8c23d7abc4711eecd`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
save

maxmemory-policy
noeviction
repl-diskless-load
swapdb
rdb-del-sync-files
yes
appendonly
yes
appendfsync
everysec
aof-load-truncated
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791708287
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

Original Redis Cluster ID: `79d1519d78125dde931ba5e0b6018eba5d16aa6e`. Process run ID: `e33b9019993c6774c2b952569a74fec6fd941b52`. Observed: 2026-10-11T16:45:00.140674+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `79d1519d78125dde931ba5e0b6018eba5d16aa6e`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
save

appendonly
yes
aof-load-truncated
yes
appendfsync
everysec
rdb-del-sync-files
yes
repl-diskless-sync
yes
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
maxmemory
134217728
aof-use-rdb-preamble
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
rdb_last_save_time:1791708287
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

Original Redis Cluster ID: `ad38f995c42cff19be4f930ff607fbeb9407801b`. Process run ID: `d45abfcb628950a2c32eb69a5ca5e432e8b07555`. Observed: 2026-10-11T16:45:01.116982+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ad38f995c42cff19be4f930ff607fbeb9407801b`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
maxmemory-policy
noeviction
appendonly
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
save

appendfsync
everysec
aof-use-rdb-preamble
yes
aof-load-truncated
yes
repl-diskless-load
swapdb
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
rdb_last_save_time:1791708287
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

Original Redis Cluster ID: `598e1d4f860a06065d51575c30f30d3fdee9ee24`. Observed after restart: 2026-10-11T16:45:10.039419+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `598e1d4f860a06065d51575c30f30d3fdee9ee24`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
appendonly
yes
no-appendfsync-on-rewrite
no
save

maxmemory-policy
noeviction
aof-load-truncated
yes
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
repl-diskless-sync
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
rdb_changes_since_last_save:145
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708308
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
aof_current_size:31848
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=145,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 2: node-2

Original Redis Cluster ID: `17ec5fb9af9141cc36bb60e2be93b1284510a4cb`. Observed after restart: 2026-10-11T16:45:11.599014+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `17ec5fb9af9141cc36bb60e2be93b1284510a4cb`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
rdb-del-sync-files
yes
save

repl-diskless-load
swapdb
appendfsync
everysec
aof-load-truncated
yes
maxmemory-policy
noeviction
appendonly
yes
maxmemory
134217728
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791708310
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
aof_current_size:31631
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

### Restart receipt 3: node-3

Original Redis Cluster ID: `c8322008da4aa9f52cc4fe9a89285a612a172420`. Observed after restart: 2026-10-11T16:45:13.182309+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c8322008da4aa9f52cc4fe9a89285a612a172420`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
repl-diskless-sync
yes
aof-load-truncated
yes
appendonly
yes
rdb-del-sync-files
yes
save

no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
appendfsync
everysec
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:144
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708311
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
aof_current_size:31631
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

### Restart receipt 4: node-4

Original Redis Cluster ID: `ad8ae597bfb05df43406b5c8c23d7abc4711eecd`. Observed after restart: 2026-10-11T16:45:14.75467+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ad8ae597bfb05df43406b5c8c23d7abc4711eecd`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
maxmemory
134217728
maxmemory-policy
noeviction
rdb-del-sync-files
yes
aof-load-truncated
yes
save

aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
appendfsync
everysec
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
rdb_changes_since_last_save:144
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708313
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
aof_current_size:31631
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
/data/temp-1791708313.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `79d1519d78125dde931ba5e0b6018eba5d16aa6e`. Observed after restart: 2026-10-11T16:45:16.30408+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `79d1519d78125dde931ba5e0b6018eba5d16aa6e`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
aof-load-truncated
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
maxmemory-policy
noeviction
save

appendonly
yes
rdb-del-sync-files
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
rdb_changes_since_last_save:145
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708315
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
aof_current_size:31848
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=145,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791708315.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `ad38f995c42cff19be4f930ff607fbeb9407801b`. Observed after restart: 2026-10-11T16:45:17.90781+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ad38f995c42cff19be4f930ff607fbeb9407801b`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
save

appendonly
yes
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
aof-load-truncated
yes
repl-diskless-sync
yes
maxmemory
134217728
rdb-del-sync-files
yes
appendfsync
everysec
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
rdb_changes_since_last_save:144
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708316
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
aof_current_size:31631
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
/data/temp-1791708316.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `598e1d4f860a06065d51575c30f30d3fdee9ee24`. Process run ID: `23bd1d5c5865c5736a15cc8acdc59f1488228832`. Observed: 2026-10-11T16:45:24.238404+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `598e1d4f860a06065d51575c30f30d3fdee9ee24`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
appendonly
yes
no-appendfsync-on-rewrite
no
save

maxmemory-policy
noeviction
aof-load-truncated
yes
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
repl-diskless-sync
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
rdb_changes_since_last_save:145
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708308
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:872448
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
aof_current_size:31848
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=145,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-2

Original Redis Cluster ID: `17ec5fb9af9141cc36bb60e2be93b1284510a4cb`. Process run ID: `0a342f88f4c5f675e63152e7cdd415ce712b5882`. Observed: 2026-10-11T16:45:25.22263+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `17ec5fb9af9141cc36bb60e2be93b1284510a4cb`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
rdb-del-sync-files
yes
save

repl-diskless-load
swapdb
appendfsync
everysec
aof-load-truncated
yes
maxmemory-policy
noeviction
appendonly
yes
maxmemory
134217728
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791708310
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
aof_current_size:31631
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

#### Restored node-3

Original Redis Cluster ID: `c8322008da4aa9f52cc4fe9a89285a612a172420`. Process run ID: `b0d431a896c9c19dcd9ac065306810380446556b`. Observed: 2026-10-11T16:45:26.248273+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c8322008da4aa9f52cc4fe9a89285a612a172420`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
repl-diskless-sync
yes
aof-load-truncated
yes
appendonly
yes
rdb-del-sync-files
yes
save

no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
appendfsync
everysec
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:144
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708311
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:888832
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
aof_current_size:31631
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

#### Restored node-4

Original Redis Cluster ID: `ad8ae597bfb05df43406b5c8c23d7abc4711eecd`. Process run ID: `d13588d0aefb197be83a04c3e603b97bd249b0fa`. Observed: 2026-10-11T16:45:27.270512+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ad8ae597bfb05df43406b5c8c23d7abc4711eecd`. Restart policy: `no`. `/data` usage: 40 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
maxmemory
134217728
maxmemory-policy
noeviction
rdb-del-sync-files
yes
aof-load-truncated
yes
save

aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
appendfsync
everysec
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
rdb_changes_since_last_save:144
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708313
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:144
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:667648
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20758
aof_base_size:20758
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
/data/appendonlydir/appendonly.aof.3.incr.aof
/data/appendonlydir/appendonly.aof.3.base.rdb

```

#### Restored node-5

Original Redis Cluster ID: `79d1519d78125dde931ba5e0b6018eba5d16aa6e`. Process run ID: `19a0c332150d2cfe6292c7a1af9de37737305947`. Observed: 2026-10-11T16:45:28.297116+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `79d1519d78125dde931ba5e0b6018eba5d16aa6e`. Restart policy: `no`. `/data` usage: 40 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
aof-load-truncated
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
maxmemory-policy
noeviction
save

appendonly
yes
rdb-del-sync-files
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
rdb_changes_since_last_save:145
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708315
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:145
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:1
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:671744
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20918
aof_base_size:20918
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=145,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.3.incr.aof
/data/appendonlydir/appendonly.aof.3.base.rdb

```

#### Restored node-6

Original Redis Cluster ID: `ad38f995c42cff19be4f930ff607fbeb9407801b`. Process run ID: `8119beea49d6a8d876001e832be5bd7829a38f4b`. Observed: 2026-10-11T16:45:29.285992+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ad38f995c42cff19be4f930ff607fbeb9407801b`. Restart policy: `no`. `/data` usage: 40 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
save

appendonly
yes
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
aof-load-truncated
yes
repl-diskless-sync
yes
maxmemory
134217728
rdb-del-sync-files
yes
appendfsync
everysec
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
rdb_changes_since_last_save:144
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708316
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:144
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:712704
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20765
aof_base_size:20765
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

Fault action completed 2283.800 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 828 | 100.00% | 0 | 199.77 | 0.632 / 0.847 / 1.081 |

Acknowledged writes: 433. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 433 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 433 | 0.632 / 0.890 / 1.192 |
| baseline | set | 433 | 0.602 / 0.847 / 1.135 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 433 | 104.470 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 828; started 828; completed 828; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 433; checked: 433; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `c2359755c7610c20f08fd8edb3a86d616be0dad17ff5b91cf389978cf97e4009`.
