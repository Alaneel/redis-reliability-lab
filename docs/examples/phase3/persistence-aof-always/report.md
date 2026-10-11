# Redis Reliability Lab — primary-failure

Run: `7d3b278ba09b73b1cab9adf0`  
Outcome: **recovered**  
Started: 2026-10-11T16:45:37.525195+08:00  
Finished: 2026-10-11T16:46:31.917869+08:00

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

Frozen acknowledged SETs: 422. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:46:00.454223+08:00 |
| Immutable ledger frozen | 2026-10-11T16:46:00.454373+08:00 |
| All six stopped verified before restart | 2026-10-11T16:46:02.682173+08:00 |
| First restart requested | 2026-10-11T16:46:03.683395+08:00 |
| Restart sequence completed | 2026-10-11T16:46:13.297084+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:46:25.188674+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:46:25.189001+08:00 |

Drain to first kill request: 0.180 ms (runtime monotonic clock).

Drain to last kill completion: 1503.391 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.214 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16715.918 |
| Original topology sustained | 18178.602 |
| Replication redundancy sustained | 18178.602 |
| Post-validation diagnostic service sustained | 26197.466 |
| AOF frozen-value recovery observed | 24734.075 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 422; checked 422; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:46:00.454402+08:00 | 2026-10-11T16:46:00.698938+08:00 |  |
| 2 | node-5 | 2026-10-11T16:46:00.698952+08:00 | 2026-10-11T16:46:00.953883+08:00 |  |
| 3 | node-6 | 2026-10-11T16:46:00.953905+08:00 | 2026-10-11T16:46:01.203153+08:00 |  |
| 4 | node-1 | 2026-10-11T16:46:01.203176+08:00 | 2026-10-11T16:46:01.456081+08:00 |  |
| 5 | node-2 | 2026-10-11T16:46:01.456091+08:00 | 2026-10-11T16:46:01.694444+08:00 |  |
| 6 | node-3 | 2026-10-11T16:46:01.694464+08:00 | 2026-10-11T16:46:01.957626+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:46:03.683395+08:00 | 2026-10-11T16:46:05.331969+08:00 |  |
| 2 | node-2 | 2026-10-11T16:46:05.331974+08:00 | 2026-10-11T16:46:06.904742+08:00 |  |
| 3 | node-3 | 2026-10-11T16:46:06.904754+08:00 | 2026-10-11T16:46:08.511861+08:00 |  |
| 4 | node-4 | 2026-10-11T16:46:08.511874+08:00 | 2026-10-11T16:46:10.120723+08:00 |  |
| 5 | node-5 | 2026-10-11T16:46:10.120727+08:00 | 2026-10-11T16:46:11.708629+08:00 |  |
| 6 | node-6 | 2026-10-11T16:46:11.708637+08:00 | 2026-10-11T16:46:13.29708+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | f93dad28b3c6a102c50505cb59d931189392d575 | 2026-10-11T16:46:02.089747+08:00 | false | no | managed Redis process is stopped |
| node-5 | 9dfae9dc4c37dc540e6f388370007e3f6d2695f0 | 2026-10-11T16:46:02.21052+08:00 | false | no | managed Redis process is stopped |
| node-6 | f242672e37e92b66fe8dbb28fbfefcd3a331b087 | 2026-10-11T16:46:02.325554+08:00 | false | no | managed Redis process is stopped |
| node-1 | bbdbffa3b380ff18bf3c84b15a7cbcfe5c382ca6 | 2026-10-11T16:46:02.439785+08:00 | false | no | managed Redis process is stopped |
| node-2 | 2f3281a4ab41fff08e9d10104a625cda590ddb02 | 2026-10-11T16:46:02.561033+08:00 | false | no | managed Redis process is stopped |
| node-3 | 20210f51d8e82776c96335912817a98b2f8a6ca0 | 2026-10-11T16:46:02.682142+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `bbdbffa3b380ff18bf3c84b15a7cbcfe5c382ca6`. Process run ID: `cf18ee4b5b2268ff540275087a337a36364a8d25`. Observed: 2026-10-11T16:45:51.170227+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `bbdbffa3b380ff18bf3c84b15a7cbcfe5c382ca6`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendfsync
always
no-appendfsync-on-rewrite
no
save

aof-load-truncated
yes
auto-aof-rewrite-percentage
0
appendonly
yes
repl-diskless-load
swapdb
maxmemory
134217728
maxmemory-policy
noeviction
repl-diskless-sync
yes
aof-use-rdb-preamble
yes
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
rdb_last_save_time:1791708340
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:1
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

Original Redis Cluster ID: `2f3281a4ab41fff08e9d10104a625cda590ddb02`. Process run ID: `0f5c4358ba59b95232d94aa962870a5ac7ff167e`. Observed: 2026-10-11T16:45:52.205167+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2f3281a4ab41fff08e9d10104a625cda590ddb02`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
appendfsync
always
appendonly
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
repl-diskless-sync
yes
repl-diskless-load
swapdb
save

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
rdb_last_save_time:1791708341
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:462848
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

Original Redis Cluster ID: `20210f51d8e82776c96335912817a98b2f8a6ca0`. Process run ID: `d37750189389a9e7fb14b43084033ec19e65f3ea`. Observed: 2026-10-11T16:45:53.189293+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `20210f51d8e82776c96335912817a98b2f8a6ca0`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendonly
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
maxmemory
134217728
save

appendfsync
always
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
aof-load-truncated
yes
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
rdb_last_save_time:1791708341
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

Original Redis Cluster ID: `f93dad28b3c6a102c50505cb59d931189392d575`. Process run ID: `1dd8be9f314fa0a3ff58782c7390e77852673ff6`. Observed: 2026-10-11T16:45:54.236159+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f93dad28b3c6a102c50505cb59d931189392d575`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
aof-use-rdb-preamble
yes
maxmemory
134217728
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
appendfsync
always
rdb-del-sync-files
yes
maxmemory-policy
noeviction
save

appendonly
yes
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
rdb_last_save_time:1791708342
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

Original Redis Cluster ID: `9dfae9dc4c37dc540e6f388370007e3f6d2695f0`. Process run ID: `63ba8df83190c98a50156050507b690811d606ef`. Observed: 2026-10-11T16:45:55.28722+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9dfae9dc4c37dc540e6f388370007e3f6d2695f0`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
maxmemory
134217728
aof-use-rdb-preamble
yes
aof-load-truncated
yes
save

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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708342
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
aof_last_cow_size:593920
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

Original Redis Cluster ID: `f242672e37e92b66fe8dbb28fbfefcd3a331b087`. Process run ID: `99973f330020aac8846ff378e7334db196671a96`. Observed: 2026-10-11T16:45:56.305687+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f242672e37e92b66fe8dbb28fbfefcd3a331b087`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
maxmemory-policy
noeviction
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
appendfsync
always
save

maxmemory
134217728
repl-diskless-sync
yes
aof-use-rdb-preamble
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
rdb_last_save_time:1791708342
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

Original Redis Cluster ID: `bbdbffa3b380ff18bf3c84b15a7cbcfe5c382ca6`. Observed after restart: 2026-10-11T16:46:05.210876+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `bbdbffa3b380ff18bf3c84b15a7cbcfe5c382ca6`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
rdb-del-sync-files
yes
appendfsync
always
repl-diskless-sync
yes
repl-diskless-load
swapdb
save

auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
aof-load-truncated
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
rdb_changes_since_last_save:141
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708363
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
aof_current_size:30972
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=141,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 2: node-2

Original Redis Cluster ID: `2f3281a4ab41fff08e9d10104a625cda590ddb02`. Observed after restart: 2026-10-11T16:46:06.78854+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2f3281a4ab41fff08e9d10104a625cda590ddb02`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
save

appendonly
yes
maxmemory
134217728
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
repl-diskless-sync
yes
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
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
rdb_changes_since_last_save:141
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708365
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
aof_current_size:30974
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=141,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

### Restart receipt 3: node-3

Original Redis Cluster ID: `20210f51d8e82776c96335912817a98b2f8a6ca0`. Observed after restart: 2026-10-11T16:46:08.387349+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `20210f51d8e82776c96335912817a98b2f8a6ca0`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
maxmemory
134217728
rdb-del-sync-files
yes
repl-diskless-sync
yes
save

appendfsync
always
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
rdb_changes_since_last_save:140
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708367
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
aof_current_size:30755
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

### Restart receipt 4: node-4

Original Redis Cluster ID: `f93dad28b3c6a102c50505cb59d931189392d575`. Observed after restart: 2026-10-11T16:46:10.003094+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f93dad28b3c6a102c50505cb59d931189392d575`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
repl-diskless-sync
yes
rdb-del-sync-files
yes
repl-diskless-load
swapdb
save

appendfsync
always
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
rdb_changes_since_last_save:140
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708368
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
aof_current_size:30755
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
/data/temp-1791708368.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `9dfae9dc4c37dc540e6f388370007e3f6d2695f0`. Observed after restart: 2026-10-11T16:46:11.590706+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9dfae9dc4c37dc540e6f388370007e3f6d2695f0`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
save

auto-aof-rewrite-percentage
0
maxmemory
134217728
rdb-del-sync-files
yes
repl-diskless-sync
yes
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
appendonly
yes
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:141
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708370
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
aof_current_size:30972
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=141,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791708370.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `f242672e37e92b66fe8dbb28fbfefcd3a331b087`. Observed after restart: 2026-10-11T16:46:13.176708+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f242672e37e92b66fe8dbb28fbfefcd3a331b087`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
repl-diskless-sync
yes
save

auto-aof-rewrite-percentage
0
maxmemory
134217728
aof-load-truncated
yes
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
appendfsync
always
appendonly
yes
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
rdb_changes_since_last_save:141
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708371
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
aof_current_size:30974
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=141,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791708371.1.rdb

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `bbdbffa3b380ff18bf3c84b15a7cbcfe5c382ca6`. Process run ID: `279bad9d0b31ed4e8375ee37ed31be4f80be5ada`. Observed: 2026-10-11T16:46:19.989575+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `bbdbffa3b380ff18bf3c84b15a7cbcfe5c382ca6`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
rdb-del-sync-files
yes
appendfsync
always
repl-diskless-sync
yes
repl-diskless-load
swapdb
save

auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
aof-load-truncated
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
rdb_changes_since_last_save:141
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708363
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
aof_current_size:30972
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=141,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-2

Original Redis Cluster ID: `2f3281a4ab41fff08e9d10104a625cda590ddb02`. Process run ID: `1fdf6037c0d51432393622cd5424d84ec9613bad`. Observed: 2026-10-11T16:46:21.008491+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2f3281a4ab41fff08e9d10104a625cda590ddb02`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
save

appendonly
yes
maxmemory
134217728
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
repl-diskless-sync
yes
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
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
rdb_changes_since_last_save:141
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708365
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:1036288
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
aof_current_size:30974
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=141,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

#### Restored node-3

Original Redis Cluster ID: `20210f51d8e82776c96335912817a98b2f8a6ca0`. Process run ID: `8d4d630fea08030c1a91d24f506371cbc2183aef`. Observed: 2026-10-11T16:46:22.034766+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `20210f51d8e82776c96335912817a98b2f8a6ca0`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
maxmemory
134217728
rdb-del-sync-files
yes
repl-diskless-sync
yes
save

appendfsync
always
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
rdb_changes_since_last_save:140
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708367
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:901120
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
aof_current_size:30755
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

#### Restored node-4

Original Redis Cluster ID: `f93dad28b3c6a102c50505cb59d931189392d575`. Process run ID: `f94b05083b3affc3450843dc1554b7e6f7aeaaf5`. Observed: 2026-10-11T16:46:23.041965+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f93dad28b3c6a102c50505cb59d931189392d575`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
repl-diskless-sync
yes
rdb-del-sync-files
yes
repl-diskless-load
swapdb
save

appendfsync
always
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
rdb_changes_since_last_save:140
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708368
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
aof_last_cow_size:667648
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20186
aof_base_size:20186
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

#### Restored node-5

Original Redis Cluster ID: `9dfae9dc4c37dc540e6f388370007e3f6d2695f0`. Process run ID: `f267a8472981270e4f7f0293ba7c493d18b23d1a`. Observed: 2026-10-11T16:46:24.053438+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9dfae9dc4c37dc540e6f388370007e3f6d2695f0`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
save

auto-aof-rewrite-percentage
0
maxmemory
134217728
rdb-del-sync-files
yes
repl-diskless-sync
yes
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
appendonly
yes
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:141
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708370
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:141
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
aof_current_size:20318
aof_base_size:20318
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=141,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.3.incr.aof
/data/appendonlydir/appendonly.aof.3.base.rdb

```

#### Restored node-6

Original Redis Cluster ID: `f242672e37e92b66fe8dbb28fbfefcd3a331b087`. Process run ID: `243350747dec130a26a40801ffcb749cc1c0df64`. Observed: 2026-10-11T16:46:25.050265+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f242672e37e92b66fe8dbb28fbfefcd3a331b087`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
repl-diskless-sync
yes
save

auto-aof-rewrite-percentage
0
maxmemory
134217728
aof-load-truncated
yes
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
appendfsync
always
appendonly
yes
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
rdb_changes_since_last_save:141
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708371
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:141
aof_enabled:1
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:0
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_rewrites:1
aof_rewrites_consecutive_failures:0
aof_last_write_status:ok
aof_last_cow_size:675840
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20338
aof_base_size:20338
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=141,expires=0,avg_ttl=0,subexpiry=0
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

Fault action completed 2227.757 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 816 | 100.00% | 0 | 196.72 | 0.732 / 1.761 / 13.020 |

Acknowledged writes: 422. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 422 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 422 | 1.030 / 2.141 / 62.038 |
| baseline | set | 422 | 0.981 / 2.141 / 62.038 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 422 | 101.735 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 829; started 816; completed 816; dropped 13. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 422; checked: 422; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `5f3362a0bdd93bf4a884c5958983a2e641eef0c9e963ba05b2dc59469d6b507f`.
