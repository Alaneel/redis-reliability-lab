# Redis Reliability Lab — dataset-trial

Run: `9e35831d92bb07158449d38b`  
Outcome: **recovered**  
Started: 2026-10-11T19:46:23.364983+08:00  
Finished: 2026-10-11T19:47:21.8518+08:00

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

Frozen acknowledged SETs: 109. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T19:46:49.45619+08:00 |
| Immutable ledger frozen | 2026-10-11T19:46:49.456297+08:00 |
| All six stopped verified before restart | 2026-10-11T19:46:51.694282+08:00 |
| First restart requested | 2026-10-11T19:46:54.695601+08:00 |
| Restart sequence completed | 2026-10-11T19:47:04.294613+08:00 |
| Frozen-ledger validation completed | 2026-10-11T19:47:16.127382+08:00 |
| Post-validation diagnostics requested | 2026-10-11T19:47:16.127703+08:00 |

Drain to first kill request: 0.132 ms (runtime monotonic clock).

Drain to last kill completion: 1522.904 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 3001.322 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 18680.138 |
| Original topology sustained | 20154.891 |
| Replication redundancy sustained | 20154.891 |
| Post-validation diagnostic service sustained | 28134.809 |
| AOF frozen-value recovery observed | 26671.096 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 109; checked 109; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T19:46:49.456322+08:00 | 2026-10-11T19:46:49.701217+08:00 |  |
| 2 | node-5 | 2026-10-11T19:46:49.701235+08:00 | 2026-10-11T19:46:49.95796+08:00 |  |
| 3 | node-6 | 2026-10-11T19:46:49.957978+08:00 | 2026-10-11T19:46:50.235637+08:00 |  |
| 4 | node-1 | 2026-10-11T19:46:50.235651+08:00 | 2026-10-11T19:46:50.485202+08:00 |  |
| 5 | node-2 | 2026-10-11T19:46:50.485212+08:00 | 2026-10-11T19:46:50.734596+08:00 |  |
| 6 | node-3 | 2026-10-11T19:46:50.734616+08:00 | 2026-10-11T19:46:50.979092+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T19:46:54.695601+08:00 | 2026-10-11T19:46:56.317807+08:00 |  |
| 2 | node-2 | 2026-10-11T19:46:56.317813+08:00 | 2026-10-11T19:46:57.925503+08:00 |  |
| 3 | node-3 | 2026-10-11T19:46:57.925512+08:00 | 2026-10-11T19:46:59.556659+08:00 |  |
| 4 | node-4 | 2026-10-11T19:46:59.556672+08:00 | 2026-10-11T19:47:01.102853+08:00 |  |
| 5 | node-5 | 2026-10-11T19:47:01.102857+08:00 | 2026-10-11T19:47:02.692881+08:00 |  |
| 6 | node-6 | 2026-10-11T19:47:02.692894+08:00 | 2026-10-11T19:47:04.29461+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 318b1cbe6870eb13836d9e09e9f4c798901aa609 | 2026-10-11T19:46:51.108542+08:00 | false | no | managed Redis process is stopped |
| node-5 | 34809acef2f58828b4beaaeca2a3da5e932a5c5e | 2026-10-11T19:46:51.229525+08:00 | false | no | managed Redis process is stopped |
| node-6 | 6de6ba07a0a0868e93498535e328bc4459f996fe | 2026-10-11T19:46:51.346383+08:00 | false | no | managed Redis process is stopped |
| node-1 | cfa0514f36132aa07ebe4eebffcb7559e394d4e5 | 2026-10-11T19:46:51.46132+08:00 | false | no | managed Redis process is stopped |
| node-2 | 9e71b90d995a3c72b2eab894a7b714241fc893c4 | 2026-10-11T19:46:51.578396+08:00 | false | no | managed Redis process is stopped |
| node-3 | 3fa1b16365bfc29b0a09558ad5eed04bb2ceaf9e | 2026-10-11T19:46:51.694254+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `cfa0514f36132aa07ebe4eebffcb7559e394d4e5`. Process run ID: `5d5c8ed29df9ff7f3e5298864248dbf96c0e1fa6`. Observed: 2026-10-11T19:46:42.181191+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `cfa0514f36132aa07ebe4eebffcb7559e394d4e5`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
rdb-del-sync-files
yes
repl-diskless-sync
yes
aof-use-rdb-preamble
yes
maxmemory
134217728
appendfsync
always
maxmemory-policy
noeviction
save

appendonly
yes
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
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
rdb_last_save_time:1791719192
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:491520
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

Original Redis Cluster ID: `9e71b90d995a3c72b2eab894a7b714241fc893c4`. Process run ID: `58330443ccaf0f5afb931a06e3ad7452ce1bb370`. Observed: 2026-10-11T19:46:43.20286+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9e71b90d995a3c72b2eab894a7b714241fc893c4`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
aof-load-truncated
yes
maxmemory
134217728
maxmemory-policy
noeviction
save

appendfsync
always
appendonly
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
rdb-del-sync-files
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
rdb_last_save_time:1791719192
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:585728
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

Original Redis Cluster ID: `3fa1b16365bfc29b0a09558ad5eed04bb2ceaf9e`. Process run ID: `ee3fee68a37410ed713eab13159251bae1723371`. Observed: 2026-10-11T19:46:44.225379+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3fa1b16365bfc29b0a09558ad5eed04bb2ceaf9e`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendonly
yes
rdb-del-sync-files
yes
repl-diskless-sync
yes
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
repl-diskless-load
swapdb
maxmemory-policy
noeviction
maxmemory
134217728
no-appendfsync-on-rewrite
no
save

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
rdb_last_save_time:1791719193
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

### Prepared node-4

Original Redis Cluster ID: `318b1cbe6870eb13836d9e09e9f4c798901aa609`. Process run ID: `9e535f21c6b234619a007c9a8e6fc255d5092cfa`. Observed: 2026-10-11T19:46:45.238515+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `318b1cbe6870eb13836d9e09e9f4c798901aa609`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
appendonly
yes
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
appendfsync
always
maxmemory
134217728
repl-diskless-load
swapdb
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
rdb_last_save_time:1791719193
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

Original Redis Cluster ID: `34809acef2f58828b4beaaeca2a3da5e932a5c5e`. Process run ID: `683a9b65d616616867c850aced6b4b646f47bb94`. Observed: 2026-10-11T19:46:46.279235+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `34809acef2f58828b4beaaeca2a3da5e932a5c5e`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
repl-diskless-load
swapdb
appendonly
yes
aof-load-truncated
yes
appendfsync
always
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
save

auto-aof-rewrite-percentage
0
maxmemory
134217728
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
rdb_last_save_time:1791719193
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

Original Redis Cluster ID: `6de6ba07a0a0868e93498535e328bc4459f996fe`. Process run ID: `bea8084f0bc258abee187ce1eab8de66986f4bed`. Observed: 2026-10-11T19:46:47.294676+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6de6ba07a0a0868e93498535e328bc4459f996fe`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
save

maxmemory
134217728
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
aof-use-rdb-preamble
yes
appendonly
yes
appendfsync
always
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719194
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

Original Redis Cluster ID: `cfa0514f36132aa07ebe4eebffcb7559e394d4e5`. Observed after restart: 2026-10-11T19:46:56.196531+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `cfa0514f36132aa07ebe4eebffcb7559e394d4e5`. Restart policy: `no`. `/data` usage: 32 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
rdb-del-sync-files
yes
aof-load-truncated
yes
appendfsync
always
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
repl-diskless-load
swapdb
save

appendonly
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
rdb_changes_since_last_save:38
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719214
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
aof_current_size:8374
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=38,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 2: node-2

Original Redis Cluster ID: `9e71b90d995a3c72b2eab894a7b714241fc893c4`. Observed after restart: 2026-10-11T19:46:57.800683+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9e71b90d995a3c72b2eab894a7b714241fc893c4`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
appendfsync
always
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
appendonly
yes
aof-load-truncated
yes
repl-diskless-sync
yes
rdb-del-sync-files
yes
save

maxmemory-policy
noeviction
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
rdb_changes_since_last_save:37
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719216
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
aof_current_size:8155
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=37,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 3: node-3

Original Redis Cluster ID: `3fa1b16365bfc29b0a09558ad5eed04bb2ceaf9e`. Observed after restart: 2026-10-11T19:46:59.435724+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3fa1b16365bfc29b0a09558ad5eed04bb2ceaf9e`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
maxmemory
134217728
appendonly
yes
aof-load-truncated
yes
repl-diskless-sync
yes
appendfsync
always
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:37
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719218
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
aof_current_size:8155
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=37,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 4: node-4

Original Redis Cluster ID: `318b1cbe6870eb13836d9e09e9f4c798901aa609`. Observed after restart: 2026-10-11T19:47:00.985057+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `318b1cbe6870eb13836d9e09e9f4c798901aa609`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
aof-load-truncated
yes
save

no-appendfsync-on-rewrite
no
appendonly
yes
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
rdb-del-sync-files
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
rdb_changes_since_last_save:37
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719219
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
aof_current_size:8155
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=37,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791719219.1.rdb

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `34809acef2f58828b4beaaeca2a3da5e932a5c5e`. Observed after restart: 2026-10-11T19:47:02.571473+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `34809acef2f58828b4beaaeca2a3da5e932a5c5e`. Restart policy: `no`. `/data` usage: 32 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
aof-load-truncated
yes
appendonly
yes
maxmemory
134217728
repl-diskless-sync
yes
repl-diskless-load
swapdb
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
save

appendfsync
always
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
rdb_changes_since_last_save:38
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719221
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
aof_current_size:8374
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=38,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791719221.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `6de6ba07a0a0868e93498535e328bc4459f996fe`. Observed after restart: 2026-10-11T19:47:04.173512+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6de6ba07a0a0868e93498535e328bc4459f996fe`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
appendfsync
always
save

repl-diskless-sync
yes
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
repl-diskless-load
swapdb
appendonly
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
rdb_changes_since_last_save:37
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719222
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
aof_current_size:8155
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=37,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791719222.1.rdb

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `cfa0514f36132aa07ebe4eebffcb7559e394d4e5`. Process run ID: `8d55e492538a2919e6db8fb62db53c81972fbf49`. Observed: 2026-10-11T19:47:10.900274+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `cfa0514f36132aa07ebe4eebffcb7559e394d4e5`. Restart policy: `no`. `/data` usage: 32 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
rdb-del-sync-files
yes
aof-load-truncated
yes
appendfsync
always
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
repl-diskless-load
swapdb
save

appendonly
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
rdb_changes_since_last_save:38
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719214
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
aof_current_size:8374
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=38,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-2

Original Redis Cluster ID: `9e71b90d995a3c72b2eab894a7b714241fc893c4`. Process run ID: `8703f8cf408dcf26b5f75441bf5e0bf082dcf129`. Observed: 2026-10-11T19:47:11.907661+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9e71b90d995a3c72b2eab894a7b714241fc893c4`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
appendfsync
always
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
appendonly
yes
aof-load-truncated
yes
repl-diskless-sync
yes
rdb-del-sync-files
yes
save

maxmemory-policy
noeviction
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
rdb_changes_since_last_save:37
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719216
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:860160
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
aof_current_size:8155
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=37,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-3

Original Redis Cluster ID: `3fa1b16365bfc29b0a09558ad5eed04bb2ceaf9e`. Process run ID: `1910aba4fb10227c9ad804b04cae4358ea96b8a2`. Observed: 2026-10-11T19:47:12.925156+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3fa1b16365bfc29b0a09558ad5eed04bb2ceaf9e`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
maxmemory
134217728
appendonly
yes
aof-load-truncated
yes
repl-diskless-sync
yes
appendfsync
always
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:37
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719218
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:880640
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
aof_current_size:8155
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=37,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-4

Original Redis Cluster ID: `318b1cbe6870eb13836d9e09e9f4c798901aa609`. Process run ID: `7e0b4b0d9b1d3eef2c1b546cd7442f295b8a54b1`. Observed: 2026-10-11T19:47:13.938949+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `318b1cbe6870eb13836d9e09e9f4c798901aa609`. Restart policy: `no`. `/data` usage: 24 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
aof-load-truncated
yes
save

no-appendfsync-on-rewrite
no
appendonly
yes
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
rdb-del-sync-files
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
rdb_changes_since_last_save:37
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719219
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:37
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:684032
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:5443
aof_base_size:5443
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=37,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.3.incr.aof
/data/appendonlydir/appendonly.aof.3.base.rdb

```

#### Restored node-5

Original Redis Cluster ID: `34809acef2f58828b4beaaeca2a3da5e932a5c5e`. Process run ID: `e1b3cfffb2dd4af2bc038c50c85c276a48f93c67`. Observed: 2026-10-11T19:47:14.98748+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `34809acef2f58828b4beaaeca2a3da5e932a5c5e`. Restart policy: `no`. `/data` usage: 24 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
aof-load-truncated
yes
appendonly
yes
maxmemory
134217728
repl-diskless-sync
yes
repl-diskless-load
swapdb
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
save

appendfsync
always
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
rdb_changes_since_last_save:38
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719221
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:38
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:1134592
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:5583
aof_base_size:5583
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=38,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.3.incr.aof
/data/appendonlydir/appendonly.aof.3.base.rdb

```

#### Restored node-6

Original Redis Cluster ID: `6de6ba07a0a0868e93498535e328bc4459f996fe`. Process run ID: `3b5ff6fdeccd04262d50fa6a34a4fadaeb1ee065`. Observed: 2026-10-11T19:47:16.086091+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6de6ba07a0a0868e93498535e328bc4459f996fe`. Restart policy: `no`. `/data` usage: 24 KiB.

Observed persistence configuration:

```text
appendfsync
always
save

repl-diskless-sync
yes
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
repl-diskless-load
swapdb
appendonly
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
rdb_changes_since_last_save:37
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719222
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:37
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:741376
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:5443
aof_base_size:5443
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=37,expires=0,avg_ttl=0,subexpiry=0
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

Fault action completed 2237.967 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 214 | 100.00% | 0 | 99.74 | 0.934 / 2.478 / 2.869 |

Acknowledged writes: 109. Ambiguous SET outcomes: 0.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | cfa0514f36132aa07ebe4eebffcb7559e394d4e5 | 2026-10-11T19:46:42.181191+08:00 | cfa0514f36132aa07ebe4eebffcb7559e394d4e5 | true |  |
| node-2 | 9e71b90d995a3c72b2eab894a7b714241fc893c4 | 2026-10-11T19:46:43.20286+08:00 | 9e71b90d995a3c72b2eab894a7b714241fc893c4 | true |  |
| node-3 | 3fa1b16365bfc29b0a09558ad5eed04bb2ceaf9e | 2026-10-11T19:46:44.225379+08:00 | 3fa1b16365bfc29b0a09558ad5eed04bb2ceaf9e | true |  |
| node-4 | 318b1cbe6870eb13836d9e09e9f4c798901aa609 | 2026-10-11T19:46:45.238515+08:00 | 318b1cbe6870eb13836d9e09e9f4c798901aa609 | true |  |
| node-5 | 34809acef2f58828b4beaaeca2a3da5e932a5c5e | 2026-10-11T19:46:46.279235+08:00 | 34809acef2f58828b4beaaeca2a3da5e932a5c5e | true |  |
| node-6 | 6de6ba07a0a0868e93498535e328bc4459f996fe | 2026-10-11T19:46:47.294676+08:00 | 6de6ba07a0a0868e93498535e328bc4459f996fe | true |  |

Before-baseline `set` policy probe at 2026-10-11T19:46:47.310491+08:00; all shards successful: true.

| Original shard index | SET acknowledged |
|---|---|
| 0 | true |
| 1 | true |
| 2 | true |

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 109 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 109 | 1.449 / 2.869 / 3.012 |
| baseline | set | 109 | 1.449 / 2.732 / 3.012 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 109 | 50.801 |

Global pacing: offered 214; started 214; completed 214; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 109; checked: 109; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `f573aa82d77dec72471df5e359ab0dcaa4f4bb5f9bcacc2c2db65a8b625896bf`.
