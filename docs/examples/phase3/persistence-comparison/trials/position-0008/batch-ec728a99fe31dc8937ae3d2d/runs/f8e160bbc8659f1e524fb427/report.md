# Redis Reliability Lab — persistence-aof-everysec

Run: `f8e160bbc8659f1e524fb427`  
Outcome: **recovered**  
Started: 2026-10-11T16:56:19.923701+08:00  
Finished: 2026-10-11T16:57:10.815303+08:00

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
| Traffic fully drained | 2026-10-11T16:56:40.901557+08:00 |
| Immutable ledger frozen | 2026-10-11T16:56:40.901607+08:00 |
| All six stopped verified before restart | 2026-10-11T16:56:43.126737+08:00 |
| First restart requested | 2026-10-11T16:56:44.127917+08:00 |
| Restart sequence completed | 2026-10-11T16:56:53.620146+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:57:04.753019+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:57:04.753169+08:00 |

Drain to first kill request: 0.084 ms (runtime monotonic clock).

Drain to last kill completion: 1503.462 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.172 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16445.386 |
| Original topology sustained | 17705.670 |
| Replication redundancy sustained | 17705.670 |
| Post-validation diagnostic service sustained | 25112.851 |
| AOF frozen-value recovery observed | 23851.185 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 417; checked 417; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:56:40.90164+08:00 | 2026-10-11T16:56:41.150564+08:00 |  |
| 2 | node-5 | 2026-10-11T16:56:41.150565+08:00 | 2026-10-11T16:56:41.40825+08:00 |  |
| 3 | node-6 | 2026-10-11T16:56:41.408269+08:00 | 2026-10-11T16:56:41.654119+08:00 |  |
| 4 | node-1 | 2026-10-11T16:56:41.654124+08:00 | 2026-10-11T16:56:41.90196+08:00 |  |
| 5 | node-2 | 2026-10-11T16:56:41.901965+08:00 | 2026-10-11T16:56:42.151804+08:00 |  |
| 6 | node-3 | 2026-10-11T16:56:42.151825+08:00 | 2026-10-11T16:56:42.405031+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:56:44.127917+08:00 | 2026-10-11T16:56:45.761063+08:00 |  |
| 2 | node-2 | 2026-10-11T16:56:45.761065+08:00 | 2026-10-11T16:56:47.406446+08:00 |  |
| 3 | node-3 | 2026-10-11T16:56:47.406454+08:00 | 2026-10-11T16:56:49.005638+08:00 |  |
| 4 | node-4 | 2026-10-11T16:56:49.005643+08:00 | 2026-10-11T16:56:50.528489+08:00 |  |
| 5 | node-5 | 2026-10-11T16:56:50.528491+08:00 | 2026-10-11T16:56:52.07831+08:00 |  |
| 6 | node-6 | 2026-10-11T16:56:52.078312+08:00 | 2026-10-11T16:56:53.620145+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 1b170a2bad63b2b10f397624c3bc22a30c1829fb | 2026-10-11T16:56:42.53606+08:00 | false | no | managed Redis process is stopped |
| node-5 | be8f126c6d3cced280700cd8b214453079e8b7cb | 2026-10-11T16:56:42.652904+08:00 | false | no | managed Redis process is stopped |
| node-6 | 3b76fbb62930f13784873374e5b469b4e4b16b93 | 2026-10-11T16:56:42.773612+08:00 | false | no | managed Redis process is stopped |
| node-1 | 5c6b7ac19b8db054b4f6c10519b5c42f4c6e850a | 2026-10-11T16:56:42.89152+08:00 | false | no | managed Redis process is stopped |
| node-2 | 0476edb5edf3058d5b7ba4f7fc77c1565cf352b4 | 2026-10-11T16:56:43.009325+08:00 | false | no | managed Redis process is stopped |
| node-3 | 3d1a44be841a53169e772a5bf137ef1e00ce715e | 2026-10-11T16:56:43.126727+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `5c6b7ac19b8db054b4f6c10519b5c42f4c6e850a`. Process run ID: `04a6c04edfad29c5b04758677d85d81f150f95f8`. Observed: 2026-10-11T16:56:31.977287+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `5c6b7ac19b8db054b4f6c10519b5c42f4c6e850a`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
save

appendonly
yes
aof-load-truncated
yes
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
appendfsync
everysec
maxmemory
134217728
maxmemory-policy
noeviction
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
rdb_last_save_time:1791708982
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

Original Redis Cluster ID: `0476edb5edf3058d5b7ba4f7fc77c1565cf352b4`. Process run ID: `6a60051c1b1067b1dab1ff39d1ec089a13d35b25`. Observed: 2026-10-11T16:56:32.953214+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `0476edb5edf3058d5b7ba4f7fc77c1565cf352b4`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
aof-load-truncated
yes
aof-use-rdb-preamble
yes
maxmemory
134217728
appendonly
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
save

maxmemory-policy
noeviction
rdb-del-sync-files
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
rdb_last_save_time:1791708982
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

Original Redis Cluster ID: `3d1a44be841a53169e772a5bf137ef1e00ce715e`. Process run ID: `ff6a3ed87a3800ad281e4473f7a0846b999da034`. Observed: 2026-10-11T16:56:33.963054+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3d1a44be841a53169e772a5bf137ef1e00ce715e`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
save

appendonly
yes
maxmemory
134217728
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
appendfsync
everysec
repl-diskless-sync
yes
repl-diskless-load
swapdb
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708983
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

### Prepared node-4

Original Redis Cluster ID: `1b170a2bad63b2b10f397624c3bc22a30c1829fb`. Process run ID: `4769badb1f0feb4f354a08d2486a025ad2da58db`. Observed: 2026-10-11T16:56:34.944064+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1b170a2bad63b2b10f397624c3bc22a30c1829fb`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
appendfsync
everysec
save

appendonly
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
aof-load-truncated
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708983
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

Original Redis Cluster ID: `be8f126c6d3cced280700cd8b214453079e8b7cb`. Process run ID: `cd34c8dce75c03113ed8aba4ada66e68a1047a06`. Observed: 2026-10-11T16:56:35.906421+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `be8f126c6d3cced280700cd8b214453079e8b7cb`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
save

no-appendfsync-on-rewrite
no
appendonly
yes
maxmemory
134217728
aof-load-truncated
yes
repl-diskless-load
swapdb
appendfsync
everysec
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791708983
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

Original Redis Cluster ID: `3b76fbb62930f13784873374e5b469b4e4b16b93`. Process run ID: `fa7f81888f2b2c0cba405aaa08fdef3b067928c2`. Observed: 2026-10-11T16:56:36.877053+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3b76fbb62930f13784873374e5b469b4e4b16b93`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
save

appendonly
yes
maxmemory
134217728
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
aof-load-truncated
yes
repl-diskless-sync
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
rdb_last_save_time:1791708984
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

### Restart receipt 1: node-1

Original Redis Cluster ID: `5c6b7ac19b8db054b4f6c10519b5c42f4c6e850a`. Observed after restart: 2026-10-11T16:56:45.638062+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `5c6b7ac19b8db054b4f6c10519b5c42f4c6e850a`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
appendfsync
everysec
maxmemory
134217728
appendonly
yes
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
save

repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709004
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

Original Redis Cluster ID: `0476edb5edf3058d5b7ba4f7fc77c1565cf352b4`. Observed after restart: 2026-10-11T16:56:47.286228+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `0476edb5edf3058d5b7ba4f7fc77c1565cf352b4`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
appendfsync
everysec
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
appendonly
yes
maxmemory
134217728
aof-load-truncated
yes
repl-diskless-sync
yes
rdb-del-sync-files
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709006
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

Original Redis Cluster ID: `3d1a44be841a53169e772a5bf137ef1e00ce715e`. Observed after restart: 2026-10-11T16:56:48.884666+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3d1a44be841a53169e772a5bf137ef1e00ce715e`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
save

aof-load-truncated
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
appendfsync
everysec
aof-use-rdb-preamble
yes
maxmemory
134217728
appendonly
yes
rdb-del-sync-files
yes
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709007
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

Original Redis Cluster ID: `1b170a2bad63b2b10f397624c3bc22a30c1829fb`. Observed after restart: 2026-10-11T16:56:50.415508+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1b170a2bad63b2b10f397624c3bc22a30c1829fb`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
rdb-del-sync-files
yes
appendonly
yes
appendfsync
everysec
maxmemory
134217728
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
save

maxmemory-policy
noeviction
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
rdb_last_save_time:1791709009
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
/data/temp-1791709009.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `be8f126c6d3cced280700cd8b214453079e8b7cb`. Observed after restart: 2026-10-11T16:56:51.965611+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `be8f126c6d3cced280700cd8b214453079e8b7cb`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
save

aof-load-truncated
yes
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
maxmemory
134217728
rdb-del-sync-files
yes
appendonly
yes
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709010
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
/data/temp-1791709010.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `3b76fbb62930f13784873374e5b469b4e4b16b93`. Observed after restart: 2026-10-11T16:56:53.503448+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3b76fbb62930f13784873374e5b469b4e4b16b93`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
appendfsync
everysec
maxmemory
134217728
repl-diskless-load
swapdb
appendonly
yes
aof-use-rdb-preamble
yes
aof-load-truncated
yes
maxmemory-policy
noeviction
repl-diskless-sync
yes
save

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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709012
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
/data/temp-1791709012.1.rdb

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `5c6b7ac19b8db054b4f6c10519b5c42f4c6e850a`. Process run ID: `22a081c5dd039fc5c2778e6f9630af3ee5494fe2`. Observed: 2026-10-11T16:56:59.62686+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `5c6b7ac19b8db054b4f6c10519b5c42f4c6e850a`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
appendfsync
everysec
maxmemory
134217728
appendonly
yes
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
save

repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709004
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

Original Redis Cluster ID: `0476edb5edf3058d5b7ba4f7fc77c1565cf352b4`. Process run ID: `c5943c8260e44a7e297f804794d62433e7d2b2be`. Observed: 2026-10-11T16:57:00.633219+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `0476edb5edf3058d5b7ba4f7fc77c1565cf352b4`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
appendfsync
everysec
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
appendonly
yes
maxmemory
134217728
aof-load-truncated
yes
repl-diskless-sync
yes
rdb-del-sync-files
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709006
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:921600
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

Original Redis Cluster ID: `3d1a44be841a53169e772a5bf137ef1e00ce715e`. Process run ID: `a875257ac41651baeec09cf9f9934a98b1483240`. Observed: 2026-10-11T16:57:01.664448+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3d1a44be841a53169e772a5bf137ef1e00ce715e`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
save

aof-load-truncated
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
appendfsync
everysec
aof-use-rdb-preamble
yes
maxmemory
134217728
appendonly
yes
rdb-del-sync-files
yes
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709007
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

Original Redis Cluster ID: `1b170a2bad63b2b10f397624c3bc22a30c1829fb`. Process run ID: `dbd7698069f9a29b62a8cc2deea104d46ca51e43`. Observed: 2026-10-11T16:57:02.656549+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1b170a2bad63b2b10f397624c3bc22a30c1829fb`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
rdb-del-sync-files
yes
appendonly
yes
appendfsync
everysec
maxmemory
134217728
aof-use-rdb-preamble
yes
repl-diskless-load
swapdb
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
save

maxmemory-policy
noeviction
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
rdb_last_save_time:1791709009
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
aof_current_size:20040
aof_base_size:20040
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

Original Redis Cluster ID: `be8f126c6d3cced280700cd8b214453079e8b7cb`. Process run ID: `912b076c2b5f224a61572c09a281f5a48524e55c`. Observed: 2026-10-11T16:57:03.663884+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `be8f126c6d3cced280700cd8b214453079e8b7cb`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
save

aof-load-truncated
yes
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
maxmemory
134217728
rdb-del-sync-files
yes
appendonly
yes
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709010
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
aof_last_cow_size:888832
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20037
aof_base_size:20037
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

Original Redis Cluster ID: `3b76fbb62930f13784873374e5b469b4e4b16b93`. Process run ID: `5ad6576a0bb6350222f10f84900bae51de5deb2e`. Observed: 2026-10-11T16:57:04.634095+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3b76fbb62930f13784873374e5b469b4e4b16b93`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
appendfsync
everysec
maxmemory
134217728
repl-diskless-load
swapdb
appendonly
yes
aof-use-rdb-preamble
yes
aof-load-truncated
yes
maxmemory-policy
noeviction
repl-diskless-sync
yes
save

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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709012
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

Fault action completed 2225.079 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 804 | 100.00% | 0 | 199.79 | 0.520 / 1.081 / 3.012 |

Acknowledged writes: 417. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 417 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 417 | 0.546 / 1.380 / 3.163 |
| baseline | set | 417 | 0.546 / 1.314 / 3.163 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 417 | 103.622 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 804; started 804; completed 804; dropped 0. These counts include all workload phases and exclude validation and probes.

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
