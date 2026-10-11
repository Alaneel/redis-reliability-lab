# Redis Reliability Lab — waitaof-frozen-persistence

Run: `adcd47811ff02526680c06fa`  
Outcome: **recovered**  
Started: 2026-10-11T17:02:40.466821+08:00  
Finished: 2026-10-11T17:03:35.590267+08:00

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

Frozen acknowledged SETs: 433. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T17:03:02.449667+08:00 |
| Immutable ledger frozen | 2026-10-11T17:03:02.449839+08:00 |
| All six stopped verified before restart | 2026-10-11T17:03:04.739824+08:00 |
| First restart requested | 2026-10-11T17:03:05.742529+08:00 |
| Restart sequence completed | 2026-10-11T17:03:15.298465+08:00 |
| Frozen-ledger validation completed | 2026-10-11T17:03:28.12986+08:00 |
| Post-validation diagnostics requested | 2026-10-11T17:03:28.130165+08:00 |

Drain to first kill request: 0.213 ms (runtime monotonic clock).

Drain to last kill completion: 1572.766 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1002.712 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 17699.263 |
| Original topology sustained | 19165.824 |
| Replication redundancy sustained | 19165.824 |
| Post-validation diagnostic service sustained | 27161.461 |
| AOF frozen-value recovery observed | 25680.174 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 433; checked 433; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T17:03:02.449879+08:00 | 2026-10-11T17:03:02.721975+08:00 |  |
| 2 | node-5 | 2026-10-11T17:03:02.721986+08:00 | 2026-10-11T17:03:02.976158+08:00 |  |
| 3 | node-6 | 2026-10-11T17:03:02.976174+08:00 | 2026-10-11T17:03:03.243006+08:00 |  |
| 4 | node-1 | 2026-10-11T17:03:03.243024+08:00 | 2026-10-11T17:03:03.504806+08:00 |  |
| 5 | node-2 | 2026-10-11T17:03:03.504819+08:00 | 2026-10-11T17:03:03.754562+08:00 |  |
| 6 | node-3 | 2026-10-11T17:03:03.754584+08:00 | 2026-10-11T17:03:04.02242+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T17:03:05.742529+08:00 | 2026-10-11T17:03:07.363703+08:00 |  |
| 2 | node-2 | 2026-10-11T17:03:07.363707+08:00 | 2026-10-11T17:03:08.935292+08:00 |  |
| 3 | node-3 | 2026-10-11T17:03:08.935301+08:00 | 2026-10-11T17:03:10.553437+08:00 |  |
| 4 | node-4 | 2026-10-11T17:03:10.553445+08:00 | 2026-10-11T17:03:12.114144+08:00 |  |
| 5 | node-5 | 2026-10-11T17:03:12.114149+08:00 | 2026-10-11T17:03:13.724723+08:00 |  |
| 6 | node-6 | 2026-10-11T17:03:13.724738+08:00 | 2026-10-11T17:03:15.298462+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 2817ac1b299aa7fc5a0fc39536d74880afde65fe | 2026-10-11T17:03:04.148245+08:00 | false | no | managed Redis process is stopped |
| node-5 | f410dadcdae84a9facaba744c266912a88881b28 | 2026-10-11T17:03:04.264732+08:00 | false | no | managed Redis process is stopped |
| node-6 | 6aa13ec6d54eeb4df4fbe41ddd7e5c6871ae39f3 | 2026-10-11T17:03:04.381336+08:00 | false | no | managed Redis process is stopped |
| node-1 | b710487161e5d37c4e0af1f95494ab7ff0c4497a | 2026-10-11T17:03:04.498336+08:00 | false | no | managed Redis process is stopped |
| node-2 | 18dfaf0dbd2559ebbed5cca11be79228dc136bd8 | 2026-10-11T17:03:04.620424+08:00 | false | no | managed Redis process is stopped |
| node-3 | 6a28f58a66a696e356dd6bbf9480d1d2abd183e9 | 2026-10-11T17:03:04.739799+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `b710487161e5d37c4e0af1f95494ab7ff0c4497a`. Process run ID: `514a894de711663e4b2c53b60f52ae0febf58eb9`. Observed: 2026-10-11T17:02:53.218494+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b710487161e5d37c4e0af1f95494ab7ff0c4497a`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
save

aof-use-rdb-preamble
yes
repl-diskless-sync
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
aof-load-truncated
yes
maxmemory-policy
noeviction
appendfsync
always
rdb-del-sync-files
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
rdb_last_save_time:1791709363
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

Original Redis Cluster ID: `18dfaf0dbd2559ebbed5cca11be79228dc136bd8`. Process run ID: `d60f126ed301953bae8a46d907268204943cf7c2`. Observed: 2026-10-11T17:02:54.277806+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `18dfaf0dbd2559ebbed5cca11be79228dc136bd8`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
repl-diskless-load
swapdb
aof-load-truncated
yes
appendfsync
always
maxmemory-policy
noeviction
rdb-del-sync-files
yes
save

appendonly
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791709363
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

### Prepared node-3

Original Redis Cluster ID: `6a28f58a66a696e356dd6bbf9480d1d2abd183e9`. Process run ID: `be5ab1c936db2c174ad6c459b29246caf1138b44`. Observed: 2026-10-11T17:02:55.291754+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6a28f58a66a696e356dd6bbf9480d1d2abd183e9`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
aof-load-truncated
yes
repl-diskless-sync
yes
save

appendfsync
always
appendonly
yes
maxmemory
134217728
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
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
rdb_last_save_time:1791709363
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

### Prepared node-4

Original Redis Cluster ID: `2817ac1b299aa7fc5a0fc39536d74880afde65fe`. Process run ID: `655fcea2874a275b8a7302f0438e2b7724d24855`. Observed: 2026-10-11T17:02:56.300853+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2817ac1b299aa7fc5a0fc39536d74880afde65fe`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
save

aof-load-truncated
yes
repl-diskless-load
swapdb
maxmemory-policy
noeviction
appendfsync
always
appendonly
yes
maxmemory
134217728
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
repl-diskless-sync
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
rdb_last_save_time:1791709364
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

Original Redis Cluster ID: `f410dadcdae84a9facaba744c266912a88881b28`. Process run ID: `2819c209812052e2917db49191d00af359d04ec9`. Observed: 2026-10-11T17:02:57.286622+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f410dadcdae84a9facaba744c266912a88881b28`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
appendonly
yes
maxmemory
134217728
repl-diskless-load
swapdb
aof-load-truncated
yes
appendfsync
always
maxmemory-policy
noeviction
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
rdb_last_save_time:1791709364
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

Original Redis Cluster ID: `6aa13ec6d54eeb4df4fbe41ddd7e5c6871ae39f3`. Process run ID: `4c3d5669098b8e51fac11ca34c0ea6c9f8464788`. Observed: 2026-10-11T17:02:58.284155+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6aa13ec6d54eeb4df4fbe41ddd7e5c6871ae39f3`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
appendfsync
always
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
save

aof-use-rdb-preamble
yes
appendonly
yes
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709364
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

Original Redis Cluster ID: `b710487161e5d37c4e0af1f95494ab7ff0c4497a`. Observed after restart: 2026-10-11T17:03:07.246709+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b710487161e5d37c4e0af1f95494ab7ff0c4497a`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
appendfsync
always
no-appendfsync-on-rewrite
no
maxmemory
134217728
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
save

repl-diskless-sync
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
appendonly
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
rdb_changes_since_last_save:146
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709386
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
aof_current_size:31952
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=146,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 2: node-2

Original Redis Cluster ID: `18dfaf0dbd2559ebbed5cca11be79228dc136bd8`. Observed after restart: 2026-10-11T17:03:08.814861+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `18dfaf0dbd2559ebbed5cca11be79228dc136bd8`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
appendonly
yes
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
repl-diskless-sync
yes
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
appendfsync
always
maxmemory-policy
noeviction
save

auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
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
rdb_last_save_time:1791709387
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
aof_current_size:31735
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

### Restart receipt 3: node-3

Original Redis Cluster ID: `6a28f58a66a696e356dd6bbf9480d1d2abd183e9`. Observed after restart: 2026-10-11T17:03:10.436329+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6a28f58a66a696e356dd6bbf9480d1d2abd183e9`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
appendonly
yes
repl-diskless-load
swapdb
aof-load-truncated
yes
maxmemory
134217728
rdb-del-sync-files
yes
save

appendfsync
always
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
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
rdb_changes_since_last_save:145
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709389
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
aof_current_size:31735
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

### Restart receipt 4: node-4

Original Redis Cluster ID: `2817ac1b299aa7fc5a0fc39536d74880afde65fe`. Observed after restart: 2026-10-11T17:03:11.998363+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2817ac1b299aa7fc5a0fc39536d74880afde65fe`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
appendfsync
always
repl-diskless-load
swapdb
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
save

aof-load-truncated
yes
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
maxmemory
134217728
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
rdb_changes_since_last_save:145
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709390
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
aof_current_size:31735
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
/data/temp-1791709390.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `f410dadcdae84a9facaba744c266912a88881b28`. Observed after restart: 2026-10-11T17:03:13.603828+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f410dadcdae84a9facaba744c266912a88881b28`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
save

aof-load-truncated
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendonly
yes
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
maxmemory
134217728
aof-use-rdb-preamble
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
rdb_changes_since_last_save:146
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709392
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
aof_current_size:31952
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=146,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791709392.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `6aa13ec6d54eeb4df4fbe41ddd7e5c6871ae39f3`. Observed after restart: 2026-10-11T17:03:15.182045+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6aa13ec6d54eeb4df4fbe41ddd7e5c6871ae39f3`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
repl-diskless-sync
yes
appendonly
yes
maxmemory
134217728
aof-load-truncated
yes
appendfsync
always
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:145
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709394
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
aof_current_size:31735
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
/data/temp-1791709394.1.rdb
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `b710487161e5d37c4e0af1f95494ab7ff0c4497a`. Process run ID: `daa9d68b66a3f3e0a88002e6811890a708e0584d`. Observed: 2026-10-11T17:03:22.865629+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b710487161e5d37c4e0af1f95494ab7ff0c4497a`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
appendfsync
always
no-appendfsync-on-rewrite
no
maxmemory
134217728
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
save

repl-diskless-sync
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
appendonly
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
rdb_changes_since_last_save:146
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709386
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:913408
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
aof_current_size:31952
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=146,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-2

Original Redis Cluster ID: `18dfaf0dbd2559ebbed5cca11be79228dc136bd8`. Process run ID: `9a9cb91ca5a20e9fe534ff0cf8f5d5870c620d70`. Observed: 2026-10-11T17:03:23.879758+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `18dfaf0dbd2559ebbed5cca11be79228dc136bd8`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
appendonly
yes
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
repl-diskless-sync
yes
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
appendfsync
always
maxmemory-policy
noeviction
save

auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
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
rdb_last_save_time:1791709387
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
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
aof_current_size:31735
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

#### Restored node-3

Original Redis Cluster ID: `6a28f58a66a696e356dd6bbf9480d1d2abd183e9`. Process run ID: `b968fe6e1ed4fe242b11da9d1abee542575cee84`. Observed: 2026-10-11T17:03:24.880742+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6a28f58a66a696e356dd6bbf9480d1d2abd183e9`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
appendonly
yes
repl-diskless-load
swapdb
aof-load-truncated
yes
maxmemory
134217728
rdb-del-sync-files
yes
save

appendfsync
always
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
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
rdb_changes_since_last_save:145
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709389
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:917504
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
aof_current_size:31735
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

#### Restored node-4

Original Redis Cluster ID: `2817ac1b299aa7fc5a0fc39536d74880afde65fe`. Process run ID: `41590b1481cf75495711164308e8923c69d32cea`. Observed: 2026-10-11T17:03:25.901726+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2817ac1b299aa7fc5a0fc39536d74880afde65fe`. Restart policy: `no`. `/data` usage: 40 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
appendfsync
always
repl-diskless-load
swapdb
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
save

aof-load-truncated
yes
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
maxmemory
134217728
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
rdb_changes_since_last_save:145
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709390
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
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:1073152
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20851
aof_base_size:20851
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

#### Restored node-5

Original Redis Cluster ID: `f410dadcdae84a9facaba744c266912a88881b28`. Process run ID: `c07e8c81030df57881c74d73c84e01392ea0cbc1`. Observed: 2026-10-11T17:03:26.912082+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f410dadcdae84a9facaba744c266912a88881b28`. Restart policy: `no`. `/data` usage: 40 KiB.

Observed persistence configuration:

```text
save

aof-load-truncated
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendonly
yes
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
maxmemory
134217728
aof-use-rdb-preamble
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
rdb_changes_since_last_save:146
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709392
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:146
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:679936
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20986
aof_base_size:20986
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=146,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.3.incr.aof
/data/appendonlydir/appendonly.aof.3.base.rdb

```

#### Restored node-6

Original Redis Cluster ID: `6aa13ec6d54eeb4df4fbe41ddd7e5c6871ae39f3`. Process run ID: `d13b642a00760cf1e2dab127cd63eacdebd995f4`. Observed: 2026-10-11T17:03:27.986583+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6aa13ec6d54eeb4df4fbe41ddd7e5c6871ae39f3`. Restart policy: `no`. `/data` usage: 40 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
repl-diskless-sync
yes
appendonly
yes
maxmemory
134217728
aof-load-truncated
yes
appendfsync
always
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:145
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709394
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
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:729088
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20844
aof_base_size:20844
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

Fault action completed 2289.966 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 829 | 100.00% | 0 | 199.87 | 1.314 / 2.478 / 3.321 |

Acknowledged writes: 433. Ambiguous SET outcomes: 0.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | b710487161e5d37c4e0af1f95494ab7ff0c4497a | 2026-10-11T17:02:53.218494+08:00 | b710487161e5d37c4e0af1f95494ab7ff0c4497a | true |  |
| node-2 | 18dfaf0dbd2559ebbed5cca11be79228dc136bd8 | 2026-10-11T17:02:54.277806+08:00 | 18dfaf0dbd2559ebbed5cca11be79228dc136bd8 | true |  |
| node-3 | 6a28f58a66a696e356dd6bbf9480d1d2abd183e9 | 2026-10-11T17:02:55.291754+08:00 | 6a28f58a66a696e356dd6bbf9480d1d2abd183e9 | true |  |
| node-4 | 2817ac1b299aa7fc5a0fc39536d74880afde65fe | 2026-10-11T17:02:56.300853+08:00 | 2817ac1b299aa7fc5a0fc39536d74880afde65fe | true |  |
| node-5 | f410dadcdae84a9facaba744c266912a88881b28 | 2026-10-11T17:02:57.286622+08:00 | f410dadcdae84a9facaba744c266912a88881b28 | true |  |
| node-6 | 6aa13ec6d54eeb4df4fbe41ddd7e5c6871ae39f3 | 2026-10-11T17:02:58.284155+08:00 | 6aa13ec6d54eeb4df4fbe41ddd7e5c6871ae39f3 | true |  |

Before-baseline `waitaof` policy probe at 2026-10-11T17:02:58.302028+08:00; all shards successful: true.
requested local/replica AOF counts: 1/1.

| Original shard index | SET acknowledged | WAITAOF attempted | Local / replica AOF reply |
|---|---|---|---|
| 0 | true | true | 1/1 |
| 1 | true | true | 1/1 |
| 2 | true | true | 1/1 |

Write policy: `waitaof`. Diagnostic stage requests are excluded from logical workload counts.
Logical write latency includes SET and WAITAOF on the same physical connection. WAITAOF policy success requires both requested local and replica AOF acknowledgement counts. These disk acknowledgements are separate from WAIT replica-memory acknowledgements; they do not establish strong consistency or a power-loss guarantee from this process-crash experiment.

| Phase | SET acknowledged | SET errors | WAITAOF attempted | Both satisfied | Short pair | Command errors | Client timeouts | Local satisfied | Replica satisfied |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| baseline | 433 | 0 | 433 | 433 | 0 | 0 | 0 | 433 | 433 |

WAITAOF reply distribution preserves the local/replica pair. Only valid two-integer replies are counted; command errors and unattempted calls have no observed pair. `0/0` is an observed short reply. Local/replica satisfied counts can overlap and must not be added.

| Phase | Local / replica AOF acknowledgements | Replies |
|---|---|---:|
| baseline | 1/1 | 433 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 433 | 1.677 / 3.012 / 3.845 |
| baseline | set | 433 | 1.030 / 1.449 / 2.602 |
| baseline | waitaof | 433 | 0.632 / 1.761 / 2.141 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 433 | 104.398 |

SET acknowledgements without a subsequent WAITAOF attempt: 0; client timeouts among those: 0. WAITAOF command errors and WAITAOF client timeouts above count attempted WAITAOF client calls only. A short integer pair is separate from a client timeout.

Configured write policy recovered for the stable interval: true. This criterion is independent of ordinary SET/GET service recovery.

Global pacing: offered 829; started 829; completed 829; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 433; checked: 433; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

WAITAOF-satisfied subset: expected 433; checked 433; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. The main ledger also includes every successful SET followed by incomplete WAITAOF.

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
- Configured acknowledgement policy is a separate current condition, measured through all-three-shard SET plus WAIT/WAITAOF on one physical connection. A stable policy milestone is historical; a later unsuccessful probe resets the current condition.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `88ab3f40ee8f36409a06c243810114981f4c932dee7444c9720c4846e2de28bc`.
