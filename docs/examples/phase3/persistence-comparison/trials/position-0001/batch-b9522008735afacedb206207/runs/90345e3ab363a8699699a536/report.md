# Redis Reliability Lab — persistence-aof-everysec

Run: `90345e3ab363a8699699a536`  
Outcome: **recovered**  
Started: 2026-10-11T16:50:18.123412+08:00  
Finished: 2026-10-11T16:51:09.87141+08:00

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

Frozen acknowledged SETs: 417. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:50:39.855872+08:00 |
| Immutable ledger frozen | 2026-10-11T16:50:39.855915+08:00 |
| All six stopped verified before restart | 2026-10-11T16:50:42.122674+08:00 |
| First restart requested | 2026-10-11T16:50:43.125019+08:00 |
| Restart sequence completed | 2026-10-11T16:50:52.748345+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:51:03.737054+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:51:03.737214+08:00 |

Drain to first kill request: 0.066 ms (runtime monotonic clock).

Drain to last kill completion: 1502.983 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1002.336 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16609.126 |
| Original topology sustained | 17798.110 |
| Replication redundancy sustained | 17798.110 |
| Post-validation diagnostic service sustained | 25144.089 |
| AOF frozen-value recovery observed | 23880.922 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 417; checked 417; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:50:39.855939+08:00 | 2026-10-11T16:50:40.100737+08:00 |  |
| 2 | node-5 | 2026-10-11T16:50:40.100743+08:00 | 2026-10-11T16:50:40.36728+08:00 |  |
| 3 | node-6 | 2026-10-11T16:50:40.367299+08:00 | 2026-10-11T16:50:40.624471+08:00 |  |
| 4 | node-1 | 2026-10-11T16:50:40.624474+08:00 | 2026-10-11T16:50:40.871642+08:00 |  |
| 5 | node-2 | 2026-10-11T16:50:40.871645+08:00 | 2026-10-11T16:50:41.105404+08:00 |  |
| 6 | node-3 | 2026-10-11T16:50:41.105412+08:00 | 2026-10-11T16:50:41.358868+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:50:43.125019+08:00 | 2026-10-11T16:50:44.725408+08:00 |  |
| 2 | node-2 | 2026-10-11T16:50:44.72541+08:00 | 2026-10-11T16:50:46.340779+08:00 |  |
| 3 | node-3 | 2026-10-11T16:50:46.340783+08:00 | 2026-10-11T16:50:47.927026+08:00 |  |
| 4 | node-4 | 2026-10-11T16:50:47.927028+08:00 | 2026-10-11T16:50:49.578109+08:00 |  |
| 5 | node-5 | 2026-10-11T16:50:49.57811+08:00 | 2026-10-11T16:50:51.181701+08:00 |  |
| 6 | node-6 | 2026-10-11T16:50:51.181702+08:00 | 2026-10-11T16:50:52.748344+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | c6c4c7cf096296fe711cf8351dbd4f9115566e07 | 2026-10-11T16:50:41.487691+08:00 | false | no | managed Redis process is stopped |
| node-5 | 920ed0feec4971fe528066efb6825813d7ebc8b8 | 2026-10-11T16:50:41.603481+08:00 | false | no | managed Redis process is stopped |
| node-6 | 239adae38f6e0ef275dd92e8029f0f30b51f9744 | 2026-10-11T16:50:41.720076+08:00 | false | no | managed Redis process is stopped |
| node-1 | ac81918c0a466b21315873b2ccf077dc56c3e729 | 2026-10-11T16:50:41.837067+08:00 | false | no | managed Redis process is stopped |
| node-2 | 147c896212a7b740d9edade8d0a5d261f658ba94 | 2026-10-11T16:50:41.955141+08:00 | false | no | managed Redis process is stopped |
| node-3 | 5c0331076522c7c3f13ecd50003432d7d39948f8 | 2026-10-11T16:50:42.12267+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `ac81918c0a466b21315873b2ccf077dc56c3e729`. Process run ID: `16431b70c3942a7e5f4cf1c78dbd69ddb5171f75`. Observed: 2026-10-11T16:50:30.901409+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ac81918c0a466b21315873b2ccf077dc56c3e729`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
save

maxmemory
134217728
appendonly
yes
repl-diskless-load
swapdb
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
rdb_last_save_time:1791708621
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:487424
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

Original Redis Cluster ID: `147c896212a7b740d9edade8d0a5d261f658ba94`. Process run ID: `4bf827f02049cbfc50bad25034ca23069de5ea51`. Observed: 2026-10-11T16:50:31.895339+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `147c896212a7b740d9edade8d0a5d261f658ba94`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
appendonly
yes
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
aof-load-truncated
yes
maxmemory
134217728
repl-diskless-load
swapdb
appendfsync
everysec
auto-aof-rewrite-percentage
0
save

repl-diskless-sync
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
rdb_last_save_time:1791708621
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:512000
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

Original Redis Cluster ID: `5c0331076522c7c3f13ecd50003432d7d39948f8`. Process run ID: `40804a336d40c436caef4b28007dbc5e3313388f`. Observed: 2026-10-11T16:50:32.878992+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `5c0331076522c7c3f13ecd50003432d7d39948f8`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
save

appendonly
yes
appendfsync
everysec
aof-load-truncated
yes
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
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
rdb_last_save_time:1791708621
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

Original Redis Cluster ID: `c6c4c7cf096296fe711cf8351dbd4f9115566e07`. Process run ID: `615d0e965c4e8ced36784d0953286b7498ecb8f8`. Observed: 2026-10-11T16:50:33.851321+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c6c4c7cf096296fe711cf8351dbd4f9115566e07`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
maxmemory-policy
noeviction
appendonly
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
maxmemory
134217728
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
repl-diskless-sync
yes
save

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
rdb_last_save_time:1791708622
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

Original Redis Cluster ID: `920ed0feec4971fe528066efb6825813d7ebc8b8`. Process run ID: `2a3cdbc57661126725f93ad3f466050363953c56`. Observed: 2026-10-11T16:50:34.830303+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `920ed0feec4971fe528066efb6825813d7ebc8b8`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
maxmemory-policy
noeviction
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendonly
yes
maxmemory
134217728
appendfsync
everysec
aof-use-rdb-preamble
yes
save

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
rdb_last_save_time:1791708622
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

Original Redis Cluster ID: `239adae38f6e0ef275dd92e8029f0f30b51f9744`. Process run ID: `6e49933bfa8d8cf3dfadc73a5898192be15c663c`. Observed: 2026-10-11T16:50:35.825173+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `239adae38f6e0ef275dd92e8029f0f30b51f9744`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
repl-diskless-sync
yes
repl-diskless-load
swapdb
rdb-del-sync-files
yes
appendonly
yes
auto-aof-rewrite-percentage
0
save

maxmemory
134217728
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
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
rdb_last_save_time:1791708622
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

Original Redis Cluster ID: `ac81918c0a466b21315873b2ccf077dc56c3e729`. Observed after restart: 2026-10-11T16:50:44.602511+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ac81918c0a466b21315873b2ccf077dc56c3e729`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
aof-load-truncated
yes
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
save

rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-load
swapdb
appendfsync
everysec
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791708643
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
aof_current_size:30534
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

### Restart receipt 2: node-2

Original Redis Cluster ID: `147c896212a7b740d9edade8d0a5d261f658ba94`. Observed after restart: 2026-10-11T16:50:46.222139+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `147c896212a7b740d9edade8d0a5d261f658ba94`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
aof-load-truncated
yes
appendonly
yes
repl-diskless-load
swapdb
save

auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
maxmemory
134217728
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708644
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

Original Redis Cluster ID: `5c0331076522c7c3f13ecd50003432d7d39948f8`. Observed after restart: 2026-10-11T16:50:47.806641+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `5c0331076522c7c3f13ecd50003432d7d39948f8`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
rdb-del-sync-files
yes
aof-load-truncated
yes
repl-diskless-load
swapdb
appendonly
yes
auto-aof-rewrite-percentage
0
save

appendfsync
everysec
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
repl-diskless-sync
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
rdb_last_save_time:1791708646
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

Original Redis Cluster ID: `c6c4c7cf096296fe711cf8351dbd4f9115566e07`. Observed after restart: 2026-10-11T16:50:49.457674+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c6c4c7cf096296fe711cf8351dbd4f9115566e07`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
rdb-del-sync-files
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
appendonly
yes
appendfsync
everysec
aof-use-rdb-preamble
yes
maxmemory
134217728
aof-load-truncated
yes
save

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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708648
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
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791708648.1.rdb

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `920ed0feec4971fe528066efb6825813d7ebc8b8`. Observed after restart: 2026-10-11T16:50:51.067542+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `920ed0feec4971fe528066efb6825813d7ebc8b8`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
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
aof-load-truncated
yes
maxmemory-policy
noeviction
appendonly
yes
save

appendfsync
everysec
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708649
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
aof_current_size:30534
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
/data/temp-1791708649.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `239adae38f6e0ef275dd92e8029f0f30b51f9744`. Observed after restart: 2026-10-11T16:50:52.635877+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `239adae38f6e0ef275dd92e8029f0f30b51f9744`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
save

maxmemory
134217728
aof-load-truncated
yes
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
appendfsync
everysec
appendonly
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
rdb_last_save_time:1791708651
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
/data/temp-1791708651.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `ac81918c0a466b21315873b2ccf077dc56c3e729`. Process run ID: `e2d88c77207bd5b75b88a4588d5ef3e80dafc957`. Observed: 2026-10-11T16:50:58.707449+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ac81918c0a466b21315873b2ccf077dc56c3e729`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
aof-load-truncated
yes
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
save

rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-load
swapdb
appendfsync
everysec
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791708643
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
aof_current_size:30534
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

#### Restored node-2

Original Redis Cluster ID: `147c896212a7b740d9edade8d0a5d261f658ba94`. Process run ID: `ee832dd0e29c46bcfb2a42713c9643fbc54dd4ce`. Observed: 2026-10-11T16:50:59.687757+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `147c896212a7b740d9edade8d0a5d261f658ba94`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
aof-load-truncated
yes
appendonly
yes
repl-diskless-load
swapdb
save

auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
maxmemory
134217728
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708644
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:839680
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

Original Redis Cluster ID: `5c0331076522c7c3f13ecd50003432d7d39948f8`. Process run ID: `5dcc68b95d631c3fb54e587e2a299bf8eef56049`. Observed: 2026-10-11T16:51:00.654194+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `5c0331076522c7c3f13ecd50003432d7d39948f8`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
rdb-del-sync-files
yes
aof-load-truncated
yes
repl-diskless-load
swapdb
appendonly
yes
auto-aof-rewrite-percentage
0
save

appendfsync
everysec
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
repl-diskless-sync
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
rdb_last_save_time:1791708646
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:815104
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

Original Redis Cluster ID: `c6c4c7cf096296fe711cf8351dbd4f9115566e07`. Process run ID: `25b1569e2260e3d77a05eff3b195043097277eac`. Observed: 2026-10-11T16:51:01.655932+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c6c4c7cf096296fe711cf8351dbd4f9115566e07`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
rdb-del-sync-files
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
appendonly
yes
appendfsync
everysec
aof-use-rdb-preamble
yes
maxmemory
134217728
aof-load-truncated
yes
save

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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708648
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
aof_current_size:20039
aof_base_size:20039
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

Original Redis Cluster ID: `920ed0feec4971fe528066efb6825813d7ebc8b8`. Process run ID: `413f2e479401c17b0369998005e3d351cc7c2ab5`. Observed: 2026-10-11T16:51:02.647412+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `920ed0feec4971fe528066efb6825813d7ebc8b8`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
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
aof-load-truncated
yes
maxmemory-policy
noeviction
appendonly
yes
save

appendfsync
everysec
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708649
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
aof_last_cow_size:663552
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20041
aof_base_size:20041
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

#### Restored node-6

Original Redis Cluster ID: `239adae38f6e0ef275dd92e8029f0f30b51f9744`. Process run ID: `ee8b3b6442bef0aec972fb7e72ffa87092778fe8`. Observed: 2026-10-11T16:51:03.619834+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `239adae38f6e0ef275dd92e8029f0f30b51f9744`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
save

maxmemory
134217728
aof-load-truncated
yes
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
appendfsync
everysec
appendonly
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
rdb_last_save_time:1791708651
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
aof_last_cow_size:667648
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20041
aof_base_size:20041
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

Fault action completed 2266.718 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 805 | 100.00% | 0 | 199.73 | 0.546 / 0.768 / 0.934 |

Acknowledged writes: 417. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 417 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 417 | 0.573 / 0.807 / 0.934 |
| baseline | set | 417 | 0.573 / 0.807 / 0.934 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 417 | 103.463 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 806; started 805; completed 805; dropped 1. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 417; checked: 417; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `216e98ee4d66a4cd497a054a7ad6776564916f52380dbc7586e7f3e402dca046`.
