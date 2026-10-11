# Redis Reliability Lab — persistence-aof-always

Run: `bc5ed8626797ee3fa6013069`  
Outcome: **recovered**  
Started: 2026-10-11T16:52:01.100803+08:00  
Finished: 2026-10-11T16:52:53.040206+08:00

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

Frozen acknowledged SETs: 417. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:52:23.049007+08:00 |
| Immutable ledger frozen | 2026-10-11T16:52:23.04904+08:00 |
| All six stopped verified before restart | 2026-10-11T16:52:25.278928+08:00 |
| First restart requested | 2026-10-11T16:52:26.280371+08:00 |
| Restart sequence completed | 2026-10-11T16:52:35.886618+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:52:47.037369+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:52:47.037505+08:00 |

Drain to first kill request: 0.063 ms (runtime monotonic clock).

Drain to last kill completion: 1516.097 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.434 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16476.006 |
| Original topology sustained | 17744.665 |
| Replication redundancy sustained | 17744.665 |
| Post-validation diagnostic service sustained | 25201.248 |
| AOF frozen-value recovery observed | 23988.105 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 417; checked 417; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:52:23.04907+08:00 | 2026-10-11T16:52:23.295308+08:00 |  |
| 2 | node-5 | 2026-10-11T16:52:23.295314+08:00 | 2026-10-11T16:52:23.53596+08:00 |  |
| 3 | node-6 | 2026-10-11T16:52:23.535968+08:00 | 2026-10-11T16:52:23.798251+08:00 |  |
| 4 | node-1 | 2026-10-11T16:52:23.798255+08:00 | 2026-10-11T16:52:24.051735+08:00 |  |
| 5 | node-2 | 2026-10-11T16:52:24.05175+08:00 | 2026-10-11T16:52:24.31608+08:00 |  |
| 6 | node-3 | 2026-10-11T16:52:24.316099+08:00 | 2026-10-11T16:52:24.565116+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:52:26.280371+08:00 | 2026-10-11T16:52:27.871071+08:00 |  |
| 2 | node-2 | 2026-10-11T16:52:27.871073+08:00 | 2026-10-11T16:52:29.421308+08:00 |  |
| 3 | node-3 | 2026-10-11T16:52:29.421317+08:00 | 2026-10-11T16:52:31.01008+08:00 |  |
| 4 | node-4 | 2026-10-11T16:52:31.010083+08:00 | 2026-10-11T16:52:32.587769+08:00 |  |
| 5 | node-5 | 2026-10-11T16:52:32.58777+08:00 | 2026-10-11T16:52:34.220962+08:00 |  |
| 6 | node-6 | 2026-10-11T16:52:34.220976+08:00 | 2026-10-11T16:52:35.886617+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 15a2730c8e38e4b5e1f4944836428e54a16eb822 | 2026-10-11T16:52:24.692206+08:00 | false | no | managed Redis process is stopped |
| node-5 | 4f7ee0415eabdd7b08ae34a31a815cc8cd043156 | 2026-10-11T16:52:24.811607+08:00 | false | no | managed Redis process is stopped |
| node-6 | 43fb13b7fe59488d25750d112e56b91384890b27 | 2026-10-11T16:52:24.930505+08:00 | false | no | managed Redis process is stopped |
| node-1 | 51cf77678a5b26a97eed1c314fa1b83677eb8f6d | 2026-10-11T16:52:25.046391+08:00 | false | no | managed Redis process is stopped |
| node-2 | f5569fe3f2ce4f89bc33cafb2e49fbbfa4b110ad | 2026-10-11T16:52:25.161523+08:00 | false | no | managed Redis process is stopped |
| node-3 | 1f738a6112c6881caf322ab0d9387eb4ecbfbb4b | 2026-10-11T16:52:25.278918+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `51cf77678a5b26a97eed1c314fa1b83677eb8f6d`. Process run ID: `33bd37b46f5b26dd77d4c0b1db3c9999faab30d9`. Observed: 2026-10-11T16:52:14.048364+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `51cf77678a5b26a97eed1c314fa1b83677eb8f6d`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
appendonly
yes
appendfsync
always
repl-diskless-load
swapdb
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
repl-diskless-sync
yes
aof-load-truncated
yes
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
rdb_last_save_time:1791708724
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

### Prepared node-2

Original Redis Cluster ID: `f5569fe3f2ce4f89bc33cafb2e49fbbfa4b110ad`. Process run ID: `73f7e452cef7a65821b136f4b0c1ac91f96f7166`. Observed: 2026-10-11T16:52:15.059049+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f5569fe3f2ce4f89bc33cafb2e49fbbfa4b110ad`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
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
repl-diskless-load
swapdb
save

maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
maxmemory
134217728
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
rdb_last_save_time:1791708724
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:581632
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

Original Redis Cluster ID: `1f738a6112c6881caf322ab0d9387eb4ecbfbb4b`. Process run ID: `aac69805376e033a8ee7287ba4c41b00f561c9a1`. Observed: 2026-10-11T16:52:16.057291+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1f738a6112c6881caf322ab0d9387eb4ecbfbb4b`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
save

maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
aof-use-rdb-preamble
yes
appendfsync
always
auto-aof-rewrite-percentage
0
maxmemory
134217728
appendonly
yes
rdb-del-sync-files
yes
repl-diskless-sync
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
rdb_last_save_time:1791708724
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:475136
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

Original Redis Cluster ID: `15a2730c8e38e4b5e1f4944836428e54a16eb822`. Process run ID: `73d0d3af2e0709e8344a1baca8448ae6883600c9`. Observed: 2026-10-11T16:52:17.06644+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `15a2730c8e38e4b5e1f4944836428e54a16eb822`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendonly
yes
rdb-del-sync-files
yes
repl-diskless-sync
yes
appendfsync
always
maxmemory
134217728
repl-diskless-load
swapdb
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
aof-load-truncated
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
rdb_last_save_time:1791708725
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
aof_last_cow_size:385024
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

Original Redis Cluster ID: `4f7ee0415eabdd7b08ae34a31a815cc8cd043156`. Process run ID: `09b7a59981ab981a9051a9d61bf9bc533cdafcb3`. Observed: 2026-10-11T16:52:18.054281+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4f7ee0415eabdd7b08ae34a31a815cc8cd043156`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-load
swapdb
appendonly
yes
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
save

aof-use-rdb-preamble
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
rdb_last_save_time:1791708725
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
aof_last_cow_size:376832
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

Original Redis Cluster ID: `43fb13b7fe59488d25750d112e56b91384890b27`. Process run ID: `26d2ac56ddd783526690d8620ddf52cb7f1db73d`. Observed: 2026-10-11T16:52:19.022105+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `43fb13b7fe59488d25750d112e56b91384890b27`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendfsync
always
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
repl-diskless-sync
yes
save

maxmemory
134217728
repl-diskless-load
swapdb
aof-use-rdb-preamble
yes
aof-load-truncated
yes
appendonly
yes
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708725
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
aof_last_cow_size:823296
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

Original Redis Cluster ID: `51cf77678a5b26a97eed1c314fa1b83677eb8f6d`. Observed after restart: 2026-10-11T16:52:27.754784+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `51cf77678a5b26a97eed1c314fa1b83677eb8f6d`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendfsync
always
repl-diskless-sync
yes
maxmemory
134217728
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
aof-load-truncated
yes
save

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
rdb_last_save_time:1791708746
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

Original Redis Cluster ID: `f5569fe3f2ce4f89bc33cafb2e49fbbfa4b110ad`. Observed after restart: 2026-10-11T16:52:29.307664+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f5569fe3f2ce4f89bc33cafb2e49fbbfa4b110ad`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
save

appendonly
yes
no-appendfsync-on-rewrite
no
appendfsync
always
rdb-del-sync-files
yes
maxmemory-policy
noeviction
aof-load-truncated
yes
repl-diskless-sync
yes
maxmemory
134217728
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791708748
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

Original Redis Cluster ID: `1f738a6112c6881caf322ab0d9387eb4ecbfbb4b`. Observed after restart: 2026-10-11T16:52:30.894241+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1f738a6112c6881caf322ab0d9387eb4ecbfbb4b`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
rdb-del-sync-files
yes
maxmemory
134217728
aof-load-truncated
yes
save

appendonly
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
appendfsync
always
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
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
rdb_last_save_time:1791708749
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

Original Redis Cluster ID: `15a2730c8e38e4b5e1f4944836428e54a16eb822`. Observed after restart: 2026-10-11T16:52:32.467426+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `15a2730c8e38e4b5e1f4944836428e54a16eb822`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
save

no-appendfsync-on-rewrite
no
appendonly
yes
appendfsync
always
auto-aof-rewrite-percentage
0
maxmemory
134217728
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
repl-diskless-load
swapdb
rdb-del-sync-files
yes
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708751
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
/data/temp-1791708751.1.rdb

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `4f7ee0415eabdd7b08ae34a31a815cc8cd043156`. Observed after restart: 2026-10-11T16:52:34.032841+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4f7ee0415eabdd7b08ae34a31a815cc8cd043156`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
save

repl-diskless-load
swapdb
appendonly
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
maxmemory
134217728
appendfsync
always
aof-use-rdb-preamble
yes
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
rdb_last_save_time:1791708752
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
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791708752.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `43fb13b7fe59488d25750d112e56b91384890b27`. Observed after restart: 2026-10-11T16:52:35.766497+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `43fb13b7fe59488d25750d112e56b91384890b27`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
rdb-del-sync-files
yes
repl-diskless-load
swapdb
save

aof-load-truncated
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
appendonly
yes
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708754
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
/data/temp-1791708754.1.rdb

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `51cf77678a5b26a97eed1c314fa1b83677eb8f6d`. Process run ID: `0aadb72f8a5fbb84f4d02f8d01c642c1a9f9d4de`. Observed: 2026-10-11T16:52:41.868934+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `51cf77678a5b26a97eed1c314fa1b83677eb8f6d`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendfsync
always
repl-diskless-sync
yes
maxmemory
134217728
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
aof-load-truncated
yes
save

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
rdb_last_save_time:1791708746
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:819200
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

Original Redis Cluster ID: `f5569fe3f2ce4f89bc33cafb2e49fbbfa4b110ad`. Process run ID: `69ef73b8cdb0465e1997d041fbdda8f1a430c36a`. Observed: 2026-10-11T16:52:42.853639+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `f5569fe3f2ce4f89bc33cafb2e49fbbfa4b110ad`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
save

appendonly
yes
no-appendfsync-on-rewrite
no
appendfsync
always
rdb-del-sync-files
yes
maxmemory-policy
noeviction
aof-load-truncated
yes
repl-diskless-sync
yes
maxmemory
134217728
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791708748
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:827392
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

Original Redis Cluster ID: `1f738a6112c6881caf322ab0d9387eb4ecbfbb4b`. Process run ID: `0e0559f43e3e352f61c68cdd9b6e546a775b2163`. Observed: 2026-10-11T16:52:43.879001+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1f738a6112c6881caf322ab0d9387eb4ecbfbb4b`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
rdb-del-sync-files
yes
maxmemory
134217728
aof-load-truncated
yes
save

appendonly
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
appendfsync
always
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
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
rdb_last_save_time:1791708749
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:827392
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

Original Redis Cluster ID: `15a2730c8e38e4b5e1f4944836428e54a16eb822`. Process run ID: `acb2bf0dc5b488a89060b81206b961f722cc2d9d`. Observed: 2026-10-11T16:52:44.897374+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `15a2730c8e38e4b5e1f4944836428e54a16eb822`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
save

no-appendfsync-on-rewrite
no
appendonly
yes
appendfsync
always
auto-aof-rewrite-percentage
0
maxmemory
134217728
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
repl-diskless-load
swapdb
rdb-del-sync-files
yes
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708751
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
aof_last_cow_size:679936
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20052
aof_base_size:20052
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

Original Redis Cluster ID: `4f7ee0415eabdd7b08ae34a31a815cc8cd043156`. Process run ID: `2ac3598570b442998fd9981f58f64213b7fd96d9`. Observed: 2026-10-11T16:52:45.878632+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4f7ee0415eabdd7b08ae34a31a815cc8cd043156`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
save

repl-diskless-load
swapdb
appendonly
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
maxmemory
134217728
appendfsync
always
aof-use-rdb-preamble
yes
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
rdb_last_save_time:1791708752
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
aof_current_size:20046
aof_base_size:20046
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

Original Redis Cluster ID: `43fb13b7fe59488d25750d112e56b91384890b27`. Process run ID: `59d46228b48728e316b1bd85b315a6d519be4e54`. Observed: 2026-10-11T16:52:46.920955+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `43fb13b7fe59488d25750d112e56b91384890b27`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
rdb-del-sync-files
yes
repl-diskless-load
swapdb
save

aof-load-truncated
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
appendonly
yes
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708754
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
aof_current_size:20043
aof_base_size:20043
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

Fault action completed 2229.842 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 805 | 100.00% | 0 | 199.92 | 0.732 / 1.449 / 2.248 |

Acknowledged writes: 417. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 417 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 417 | 1.030 / 1.598 / 2.478 |
| baseline | set | 417 | 1.030 / 1.598 / 2.478 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 417 | 103.559 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 805; started 805; completed 805; dropped 0. These counts include all workload phases and exclude validation and probes.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `4d4fd037cd04b6f208194358a86091403dc4075d14d2149be6a06dc3067b954e`.
