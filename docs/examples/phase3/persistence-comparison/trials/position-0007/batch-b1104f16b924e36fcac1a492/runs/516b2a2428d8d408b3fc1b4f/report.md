# Redis Reliability Lab — persistence-aof-always

Run: `516b2a2428d8d408b3fc1b4f`  
Outcome: **recovered**  
Started: 2026-10-11T16:55:27.919392+08:00  
Finished: 2026-10-11T16:56:19.757093+08:00

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

Frozen acknowledged SETs: 418. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:55:50.065754+08:00 |
| Immutable ledger frozen | 2026-10-11T16:55:50.065875+08:00 |
| All six stopped verified before restart | 2026-10-11T16:55:52.224796+08:00 |
| First restart requested | 2026-10-11T16:55:53.226124+08:00 |
| Restart sequence completed | 2026-10-11T16:56:02.653164+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:56:13.743838+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:56:13.744113+08:00 |

Drain to first kill request: 0.176 ms (runtime monotonic clock).

Drain to last kill completion: 1476.576 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.320 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16323.188 |
| Original topology sustained | 17565.176 |
| Replication redundancy sustained | 17565.176 |
| Post-validation diagnostic service sustained | 24954.929 |
| AOF frozen-value recovery observed | 23677.718 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 418; checked 418; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:55:50.06593+08:00 | 2026-10-11T16:55:50.304813+08:00 |  |
| 2 | node-5 | 2026-10-11T16:55:50.304827+08:00 | 2026-10-11T16:55:50.548737+08:00 |  |
| 3 | node-6 | 2026-10-11T16:55:50.548748+08:00 | 2026-10-11T16:55:50.803881+08:00 |  |
| 4 | node-1 | 2026-10-11T16:55:50.803887+08:00 | 2026-10-11T16:55:51.057769+08:00 |  |
| 5 | node-2 | 2026-10-11T16:55:51.057772+08:00 | 2026-10-11T16:55:51.296803+08:00 |  |
| 6 | node-3 | 2026-10-11T16:55:51.296807+08:00 | 2026-10-11T16:55:51.542342+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:55:53.226124+08:00 | 2026-10-11T16:55:54.789436+08:00 |  |
| 2 | node-2 | 2026-10-11T16:55:54.789442+08:00 | 2026-10-11T16:55:56.406082+08:00 |  |
| 3 | node-3 | 2026-10-11T16:55:56.406087+08:00 | 2026-10-11T16:55:58.033126+08:00 |  |
| 4 | node-4 | 2026-10-11T16:55:58.03313+08:00 | 2026-10-11T16:55:59.558829+08:00 |  |
| 5 | node-5 | 2026-10-11T16:55:59.55883+08:00 | 2026-10-11T16:56:01.098798+08:00 |  |
| 6 | node-6 | 2026-10-11T16:56:01.098809+08:00 | 2026-10-11T16:56:02.653162+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 7cd6c90674fa08cebafb66b260bea00407e20849 | 2026-10-11T16:55:51.662387+08:00 | false | no | managed Redis process is stopped |
| node-5 | 9f24cb9e1d9e052106aed9eb57897c1f280af210 | 2026-10-11T16:55:51.775564+08:00 | false | no | managed Redis process is stopped |
| node-6 | 688446465e7611ca9643225b343f519d9c7e25fb | 2026-10-11T16:55:51.887549+08:00 | false | no | managed Redis process is stopped |
| node-1 | c37e839e04d256003bb2b71f6a467460f5f7a283 | 2026-10-11T16:55:51.999794+08:00 | false | no | managed Redis process is stopped |
| node-2 | aa3b1633f5fd3d83db63665d7b93fd5a22d085fa | 2026-10-11T16:55:52.111222+08:00 | false | no | managed Redis process is stopped |
| node-3 | f29b3efe695d49ec16351da3ac3d5e45a2494488 | 2026-10-11T16:55:52.224782+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `c37e839e04d256003bb2b71f6a467460f5f7a283`. Process run ID: `4151761533579507864ae8cac1dd6a35b44fd9e3`. Observed: 2026-10-11T16:55:41.147499+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c37e839e04d256003bb2b71f6a467460f5f7a283`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
appendfsync
always
repl-diskless-load
swapdb
maxmemory
134217728
save

rdb-del-sync-files
yes
repl-diskless-sync
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708931
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:483328
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

Original Redis Cluster ID: `aa3b1633f5fd3d83db63665d7b93fd5a22d085fa`. Process run ID: `be0ffa99544f3baca2d7dd6cf061deedbc28f538`. Observed: 2026-10-11T16:55:42.127791+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `aa3b1633f5fd3d83db63665d7b93fd5a22d085fa`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendfsync
always
maxmemory-policy
noeviction
repl-diskless-sync
yes
appendonly
yes
maxmemory
134217728
repl-diskless-load
swapdb
aof-use-rdb-preamble
yes
aof-load-truncated
yes
rdb-del-sync-files
yes
save

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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708931
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:573440
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

Original Redis Cluster ID: `f29b3efe695d49ec16351da3ac3d5e45a2494488`. Process run ID: `3b62c52af5adfa064906dbcbd99eab97cdab0891`. Observed: 2026-10-11T16:55:43.098399+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f29b3efe695d49ec16351da3ac3d5e45a2494488`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
appendonly
yes
appendfsync
always
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
maxmemory
134217728
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791708931
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:1
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

Original Redis Cluster ID: `7cd6c90674fa08cebafb66b260bea00407e20849`. Process run ID: `7e35b1ba87da78332644aef8141d3eeb5244db19`. Observed: 2026-10-11T16:55:44.078309+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7cd6c90674fa08cebafb66b260bea00407e20849`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
appendonly
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
save

maxmemory
134217728
appendfsync
always
maxmemory-policy
noeviction
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
rdb_last_save_time:1791708932
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
aof_last_rewrite_time_sec:1
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

Original Redis Cluster ID: `9f24cb9e1d9e052106aed9eb57897c1f280af210`. Process run ID: `46798fbaf534f981b94ce2b7d8ed7b1b563bd51d`. Observed: 2026-10-11T16:55:45.03425+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9f24cb9e1d9e052106aed9eb57897c1f280af210`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
save

appendfsync
always
aof-load-truncated
yes
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
rdb-del-sync-files
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
appendonly
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
rdb_last_save_time:1791708932
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

### Prepared node-6

Original Redis Cluster ID: `688446465e7611ca9643225b343f519d9c7e25fb`. Process run ID: `5899b20e0b180b52ea17144c268963743122c721`. Observed: 2026-10-11T16:55:46.038967+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `688446465e7611ca9643225b343f519d9c7e25fb`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendonly
yes
repl-diskless-sync
yes
repl-diskless-load
swapdb
appendfsync
always
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
save

rdb-del-sync-files
yes
maxmemory
134217728
maxmemory-policy
noeviction
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
rdb_last_save_time:1791708932
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

Original Redis Cluster ID: `c37e839e04d256003bb2b71f6a467460f5f7a283`. Observed after restart: 2026-10-11T16:55:54.675862+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c37e839e04d256003bb2b71f6a467460f5f7a283`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
maxmemory
134217728
no-appendfsync-on-rewrite
no
appendonly
yes
rdb-del-sync-files
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
save

repl-diskless-sync
yes
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
rdb_changes_since_last_save:140
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708953
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
aof_current_size:30753
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=140,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 2: node-2

Original Redis Cluster ID: `aa3b1633f5fd3d83db63665d7b93fd5a22d085fa`. Observed after restart: 2026-10-11T16:55:56.287607+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `aa3b1633f5fd3d83db63665d7b93fd5a22d085fa`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
repl-diskless-load
swapdb
appendfsync
always
repl-diskless-sync
yes
maxmemory
134217728
maxmemory-policy
noeviction
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708955
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
aof_current_size:30536
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=139,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 3: node-3

Original Redis Cluster ID: `f29b3efe695d49ec16351da3ac3d5e45a2494488`. Observed after restart: 2026-10-11T16:55:57.915541+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f29b3efe695d49ec16351da3ac3d5e45a2494488`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
aof-load-truncated
yes
repl-diskless-sync
yes
save

maxmemory-policy
noeviction
repl-diskless-load
swapdb
appendfsync
always
auto-aof-rewrite-percentage
0
maxmemory
134217728
no-appendfsync-on-rewrite
no
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708956
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
aof_current_size:30536
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=139,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 4: node-4

Original Redis Cluster ID: `7cd6c90674fa08cebafb66b260bea00407e20849`. Observed after restart: 2026-10-11T16:55:59.445612+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7cd6c90674fa08cebafb66b260bea00407e20849`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
save

appendonly
yes
repl-diskless-load
swapdb
maxmemory-policy
noeviction
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
appendfsync
always
maxmemory
134217728
rdb-del-sync-files
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708958
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
aof_current_size:30536
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=139,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/temp-1791708958.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `9f24cb9e1d9e052106aed9eb57897c1f280af210`. Observed after restart: 2026-10-11T16:56:00.983257+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9f24cb9e1d9e052106aed9eb57897c1f280af210`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
save

auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
appendonly
yes
rdb-del-sync-files
yes
aof-load-truncated
yes
maxmemory
134217728
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
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
rdb_changes_since_last_save:140
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708959
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
aof_current_size:30753
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=140,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791708959.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `688446465e7611ca9643225b343f519d9c7e25fb`. Observed after restart: 2026-10-11T16:56:02.535863+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `688446465e7611ca9643225b343f519d9c7e25fb`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
appendfsync
always
appendonly
yes
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
repl-diskless-load
swapdb
save

maxmemory
134217728
aof-load-truncated
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708961
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
aof_current_size:30536
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=139,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/temp-1791708961.1.rdb
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `c37e839e04d256003bb2b71f6a467460f5f7a283`. Process run ID: `dee000b98eb3dfd649365215aa0591fdb3f8c4a7`. Observed: 2026-10-11T16:56:08.683639+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c37e839e04d256003bb2b71f6a467460f5f7a283`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
maxmemory
134217728
no-appendfsync-on-rewrite
no
appendonly
yes
rdb-del-sync-files
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
save

repl-diskless-sync
yes
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
rdb_changes_since_last_save:140
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708953
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
aof_current_size:30753
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=140,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-2

Original Redis Cluster ID: `aa3b1633f5fd3d83db63665d7b93fd5a22d085fa`. Process run ID: `6d8822a4f8ac26a99deb1bac91f8c26eb2a35e5b`. Observed: 2026-10-11T16:56:09.651238+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `aa3b1633f5fd3d83db63665d7b93fd5a22d085fa`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
repl-diskless-load
swapdb
appendfsync
always
repl-diskless-sync
yes
maxmemory
134217728
maxmemory-policy
noeviction
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708955
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:909312
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
aof_current_size:30536
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=139,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-3

Original Redis Cluster ID: `f29b3efe695d49ec16351da3ac3d5e45a2494488`. Process run ID: `ff6599f96529f8bfc48495020bcca876347ba8b8`. Observed: 2026-10-11T16:56:10.625222+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f29b3efe695d49ec16351da3ac3d5e45a2494488`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
aof-load-truncated
yes
repl-diskless-sync
yes
save

maxmemory-policy
noeviction
repl-diskless-load
swapdb
appendfsync
always
auto-aof-rewrite-percentage
0
maxmemory
134217728
no-appendfsync-on-rewrite
no
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708956
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:905216
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
aof_current_size:30536
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=139,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-4

Original Redis Cluster ID: `7cd6c90674fa08cebafb66b260bea00407e20849`. Process run ID: `58200f90bd100d35aa08a3d13f355ea205f17c33`. Observed: 2026-10-11T16:56:11.598398+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7cd6c90674fa08cebafb66b260bea00407e20849`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
save

appendonly
yes
repl-diskless-load
swapdb
maxmemory-policy
noeviction
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
appendfsync
always
maxmemory
134217728
rdb-del-sync-files
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708958
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:139
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:1052672
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20034
aof_base_size:20034
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=139,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.3.incr.aof
/data/appendonlydir/appendonly.aof.3.base.rdb

```

#### Restored node-5

Original Redis Cluster ID: `9f24cb9e1d9e052106aed9eb57897c1f280af210`. Process run ID: `a442f0ba5dc79db04bd523686e30bc8acd7a67fb`. Observed: 2026-10-11T16:56:12.608047+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9f24cb9e1d9e052106aed9eb57897c1f280af210`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
save

auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
appendonly
yes
rdb-del-sync-files
yes
aof-load-truncated
yes
maxmemory
134217728
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
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
rdb_changes_since_last_save:140
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708959
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:140
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:671744
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20196
aof_base_size:20196
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=140,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.3.incr.aof
/data/appendonlydir/appendonly.aof.3.base.rdb

```

#### Restored node-6

Original Redis Cluster ID: `688446465e7611ca9643225b343f519d9c7e25fb`. Process run ID: `6c2bd8f8c8a5f9fcefd42f082c7b1bea92390a0b`. Observed: 2026-10-11T16:56:13.609488+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `688446465e7611ca9643225b343f519d9c7e25fb`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
appendfsync
always
appendonly
yes
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
repl-diskless-load
swapdb
save

maxmemory
134217728
aof-load-truncated
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708961
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:139
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:671744
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20051
aof_base_size:20051
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=139,expires=0,avg_ttl=0,subexpiry=0
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

Fault action completed 2158.851 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 805 | 100.00% | 0 | 199.92 | 0.732 / 1.380 / 2.478 |

Acknowledged writes: 418. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 418 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 418 | 0.981 / 1.521 / 2.602 |
| baseline | set | 418 | 0.981 / 1.521 / 2.602 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 418 | 103.811 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 805; started 805; completed 805; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 418; checked: 418; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `4d4fd037cd04b6f208194358a86091403dc4075d14d2149be6a06dc3067b954e`.
