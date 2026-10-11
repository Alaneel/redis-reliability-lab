# Redis Reliability Lab — dataset-trial

Run: `440307057705312273b742fe`  
Outcome: **recovered**  
Started: 2026-10-11T19:59:01.719231+08:00  
Finished: 2026-10-11T19:59:57.263034+08:00

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

Frozen acknowledged SETs: 178. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T19:59:22.729787+08:00 |
| Immutable ledger frozen | 2026-10-11T19:59:22.729906+08:00 |
| All six stopped verified before restart | 2026-10-11T19:59:25.185545+08:00 |
| First restart requested | 2026-10-11T19:59:29.187072+08:00 |
| Restart sequence completed | 2026-10-11T19:59:39.573711+08:00 |
| Frozen-ledger validation completed | 2026-10-11T19:59:51.924325+08:00 |
| Post-validation diagnostics requested | 2026-10-11T19:59:51.924628+08:00 |

Drain to first kill request: 0.159 ms (runtime monotonic clock).

Drain to last kill completion: 1632.310 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 4001.492 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 21013.593 |
| Original topology sustained | 22203.309 |
| Replication redundancy sustained | 22203.309 |
| Post-validation diagnostic service sustained | 30392.892 |
| AOF frozen-value recovery observed | 29194.127 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 178; checked 178; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T19:59:22.729945+08:00 | 2026-10-11T19:59:22.994204+08:00 |  |
| 2 | node-5 | 2026-10-11T19:59:22.994217+08:00 | 2026-10-11T19:59:23.255209+08:00 |  |
| 3 | node-6 | 2026-10-11T19:59:23.255233+08:00 | 2026-10-11T19:59:23.552185+08:00 |  |
| 4 | node-1 | 2026-10-11T19:59:23.552209+08:00 | 2026-10-11T19:59:23.820542+08:00 |  |
| 5 | node-2 | 2026-10-11T19:59:23.820553+08:00 | 2026-10-11T19:59:24.086014+08:00 |  |
| 6 | node-3 | 2026-10-11T19:59:24.086032+08:00 | 2026-10-11T19:59:24.362112+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T19:59:29.187072+08:00 | 2026-10-11T19:59:30.995204+08:00 |  |
| 2 | node-2 | 2026-10-11T19:59:30.995214+08:00 | 2026-10-11T19:59:32.745409+08:00 |  |
| 3 | node-3 | 2026-10-11T19:59:32.745419+08:00 | 2026-10-11T19:59:34.464709+08:00 |  |
| 4 | node-4 | 2026-10-11T19:59:34.464719+08:00 | 2026-10-11T19:59:36.189989+08:00 |  |
| 5 | node-5 | 2026-10-11T19:59:36.190007+08:00 | 2026-10-11T19:59:37.883631+08:00 |  |
| 6 | node-6 | 2026-10-11T19:59:37.883638+08:00 | 2026-10-11T19:59:39.573707+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 4a862010dd6d8d7577f856d93622db91fc28a591 | 2026-10-11T19:59:24.496324+08:00 | false | no | managed Redis process is stopped |
| node-5 | 2030e1684c35fb229e379f5d2713ca598814a111 | 2026-10-11T19:59:24.63158+08:00 | false | no | managed Redis process is stopped |
| node-6 | 1346c223f4fccd93b36810b84266bd1d3dc83c34 | 2026-10-11T19:59:24.766045+08:00 | false | no | managed Redis process is stopped |
| node-1 | 3590b7b79dbfa3dee36180a810362c6db3da40aa | 2026-10-11T19:59:24.912086+08:00 | false | no | managed Redis process is stopped |
| node-2 | eff38617be643479bf150d926099e769782144e1 | 2026-10-11T19:59:25.047961+08:00 | false | no | managed Redis process is stopped |
| node-3 | c9f5c02396be4206dedd79076d0aa92517c46ef2 | 2026-10-11T19:59:25.18551+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `3590b7b79dbfa3dee36180a810362c6db3da40aa`. Process run ID: `c8c6d02abe0e49725f25520c2490405971d2bbca`. Observed: 2026-10-11T19:59:15.121876+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3590b7b79dbfa3dee36180a810362c6db3da40aa`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
appendfsync
always
repl-diskless-load
swapdb
maxmemory
134217728
no-appendfsync-on-rewrite
no
appendonly
yes
save

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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719945
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

### Prepared node-2

Original Redis Cluster ID: `eff38617be643479bf150d926099e769782144e1`. Process run ID: `8fcc03bd3208f5924a3da654392fbdf79c981905`. Observed: 2026-10-11T19:59:16.222182+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `eff38617be643479bf150d926099e769782144e1`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
appendonly
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
maxmemory
134217728
auto-aof-rewrite-percentage
0
save

repl-diskless-sync
yes
repl-diskless-load
swapdb
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
rdb_last_save_time:1791719945
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
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

### Prepared node-3

Original Redis Cluster ID: `c9f5c02396be4206dedd79076d0aa92517c46ef2`. Process run ID: `1fe2de1535b1afd1c79b6fab0161457c4cf905d9`. Observed: 2026-10-11T19:59:17.301695+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c9f5c02396be4206dedd79076d0aa92517c46ef2`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
repl-diskless-load
swapdb
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
save

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
rdb_last_save_time:1791719945
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

Original Redis Cluster ID: `4a862010dd6d8d7577f856d93622db91fc28a591`. Process run ID: `6d2f2e68b6632ca81e7ec728f9e2c88e386e5c98`. Observed: 2026-10-11T19:59:18.393823+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4a862010dd6d8d7577f856d93622db91fc28a591`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
no-appendfsync-on-rewrite
no
maxmemory
134217728
appendfsync
always
aof-load-truncated
yes
rdb-del-sync-files
yes
appendonly
yes
auto-aof-rewrite-percentage
0
save

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
rdb_last_save_time:1791719946
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

Original Redis Cluster ID: `2030e1684c35fb229e379f5d2713ca598814a111`. Process run ID: `6780e6eef110443747a17ca9bd28dd5c0b7ffdee`. Observed: 2026-10-11T19:59:19.468865+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2030e1684c35fb229e379f5d2713ca598814a111`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
maxmemory-policy
noeviction
repl-diskless-sync
yes
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
appendonly
yes
maxmemory
134217728
save

appendfsync
always
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791719946
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
aof_last_cow_size:819200
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

Original Redis Cluster ID: `1346c223f4fccd93b36810b84266bd1d3dc83c34`. Process run ID: `69f22c788c21ca30627ab53a3ed358027b911064`. Observed: 2026-10-11T19:59:20.568185+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1346c223f4fccd93b36810b84266bd1d3dc83c34`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendonly
yes
maxmemory
134217728
repl-diskless-load
swapdb
aof-load-truncated
yes
maxmemory-policy
noeviction
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
save

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
rdb_last_save_time:1791719946
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

Original Redis Cluster ID: `3590b7b79dbfa3dee36180a810362c6db3da40aa`. Observed after restart: 2026-10-11T19:59:30.852676+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3590b7b79dbfa3dee36180a810362c6db3da40aa`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
save

aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
appendonly
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
rdb-del-sync-files
yes
aof-load-truncated
yes
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719969
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

Original Redis Cluster ID: `eff38617be643479bf150d926099e769782144e1`. Observed after restart: 2026-10-11T19:59:32.617162+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `eff38617be643479bf150d926099e769782144e1`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
appendfsync
always
rdb-del-sync-files
yes
appendonly
yes
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
save

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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719971
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

### Restart receipt 3: node-3

Original Redis Cluster ID: `c9f5c02396be4206dedd79076d0aa92517c46ef2`. Observed after restart: 2026-10-11T19:59:34.335968+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c9f5c02396be4206dedd79076d0aa92517c46ef2`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
appendfsync
always
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
save

maxmemory
134217728
repl-diskless-load
swapdb
appendonly
yes
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719973
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

Original Redis Cluster ID: `4a862010dd6d8d7577f856d93622db91fc28a591`. Observed after restart: 2026-10-11T19:59:36.05588+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4a862010dd6d8d7577f856d93622db91fc28a591`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
save

maxmemory-policy
noeviction
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
appendfsync
always
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719974
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
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791719974.1.rdb

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `2030e1684c35fb229e379f5d2713ca598814a111`. Observed after restart: 2026-10-11T19:59:37.748441+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2030e1684c35fb229e379f5d2713ca598814a111`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
appendfsync
always
aof-use-rdb-preamble
yes
aof-load-truncated
yes
save

maxmemory-policy
noeviction
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719976
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
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof
/data/temp-1791719976.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `1346c223f4fccd93b36810b84266bd1d3dc83c34`. Observed after restart: 2026-10-11T19:59:39.439497+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1346c223f4fccd93b36810b84266bd1d3dc83c34`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
save

appendfsync
always
appendonly
yes
maxmemory
134217728
rdb-del-sync-files
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
aof-load-truncated
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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719978
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
/data/temp-1791719978.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `3590b7b79dbfa3dee36180a810362c6db3da40aa`. Process run ID: `b7ae5c3d810cfa94359ec731949d4696e2a49953`. Observed: 2026-10-11T19:59:46.283774+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3590b7b79dbfa3dee36180a810362c6db3da40aa`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
save

aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
appendonly
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
rdb-del-sync-files
yes
aof-load-truncated
yes
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719969
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

Original Redis Cluster ID: `eff38617be643479bf150d926099e769782144e1`. Process run ID: `ecf4148fa670f195d0c9fa4d9f0ace7e0402563d`. Observed: 2026-10-11T19:59:47.338548+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `eff38617be643479bf150d926099e769782144e1`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
appendfsync
always
rdb-del-sync-files
yes
appendonly
yes
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
save

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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719971
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:868352
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

#### Restored node-3

Original Redis Cluster ID: `c9f5c02396be4206dedd79076d0aa92517c46ef2`. Process run ID: `8b90fb1fcfda0deb2d10c6f8d9f286c650fc8cf4`. Observed: 2026-10-11T19:59:48.488946+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c9f5c02396be4206dedd79076d0aa92517c46ef2`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
appendfsync
always
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
save

maxmemory
134217728
repl-diskless-load
swapdb
appendonly
yes
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719973
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

Original Redis Cluster ID: `4a862010dd6d8d7577f856d93622db91fc28a591`. Process run ID: `d451ccda68d38a6e619d6cd032e647b528748c8a`. Observed: 2026-10-11T19:59:49.664041+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4a862010dd6d8d7577f856d93622db91fc28a591`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
save

maxmemory-policy
noeviction
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
appendfsync
always
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719974
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
aof_last_cow_size:688128
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:8874
aof_base_size:8874
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

Original Redis Cluster ID: `2030e1684c35fb229e379f5d2713ca598814a111`. Process run ID: `796c4f7045cef513f9eadf56d8108484922f2d55`. Observed: 2026-10-11T19:59:50.760585+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2030e1684c35fb229e379f5d2713ca598814a111`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
appendfsync
always
aof-use-rdb-preamble
yes
aof-load-truncated
yes
save

maxmemory-policy
noeviction
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
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
rdb_changes_since_last_save:61
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719976
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
aof_last_cow_size:696320
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:9028
aof_base_size:9028
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

Original Redis Cluster ID: `1346c223f4fccd93b36810b84266bd1d3dc83c34`. Process run ID: `efaf3869bbb82f908d40a445dba587d965f293ed`. Observed: 2026-10-11T19:59:51.857007+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1346c223f4fccd93b36810b84266bd1d3dc83c34`. Restart policy: `no`. `/data` usage: 28 KiB.

Observed persistence configuration:

```text
save

appendfsync
always
appendonly
yes
maxmemory
134217728
rdb-del-sync-files
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
aof-load-truncated
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
rdb_changes_since_last_save:60
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719978
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
aof_last_cow_size:688128
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

Fault action completed 2455.585 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 364 | 100.00% | 0 | 169.62 | 0.732 / 1.521 / 2.478 |

Acknowledged writes: 178. Ambiguous SET outcomes: 0.

### Before-baseline persistence and command evidence

Persistence preparation contains 6 node observations. Raw CONFIG, INFO and data-file replies are retained in JSON; these observations are separate from frozen-ledger verification.

| Node | Original Redis Cluster ID | Observed at | Observed Redis Cluster ID | Redis PONG | Error |
|---|---|---|---|---|---|
| node-1 | 3590b7b79dbfa3dee36180a810362c6db3da40aa | 2026-10-11T19:59:15.121876+08:00 | 3590b7b79dbfa3dee36180a810362c6db3da40aa | true |  |
| node-2 | eff38617be643479bf150d926099e769782144e1 | 2026-10-11T19:59:16.222182+08:00 | eff38617be643479bf150d926099e769782144e1 | true |  |
| node-3 | c9f5c02396be4206dedd79076d0aa92517c46ef2 | 2026-10-11T19:59:17.301695+08:00 | c9f5c02396be4206dedd79076d0aa92517c46ef2 | true |  |
| node-4 | 4a862010dd6d8d7577f856d93622db91fc28a591 | 2026-10-11T19:59:18.393823+08:00 | 4a862010dd6d8d7577f856d93622db91fc28a591 | true |  |
| node-5 | 2030e1684c35fb229e379f5d2713ca598814a111 | 2026-10-11T19:59:19.468865+08:00 | 2030e1684c35fb229e379f5d2713ca598814a111 | true |  |
| node-6 | 1346c223f4fccd93b36810b84266bd1d3dc83c34 | 2026-10-11T19:59:20.568185+08:00 | 1346c223f4fccd93b36810b84266bd1d3dc83c34 | true |  |

Before-baseline `set` policy probe at 2026-10-11T19:59:20.583626+08:00; all shards successful: true.

| Original shard index | SET acknowledged |
|---|---|
| 0 | true |
| 1 | true |
| 2 | true |

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 178 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 178 | 1.135 / 2.039 / 2.732 |
| baseline | set | 178 | 1.081 / 2.039 / 2.732 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 178 | 82.944 |

Global pacing: offered 364; started 364; completed 364; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 178; checked: 178; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `451309c55308664870ca3178164b2bd1ea718b89d5d4405735a92c5ef9a57197`.
