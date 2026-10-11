# Redis Reliability Lab — persistence-aof-always

Run: `a50433e55ee7e3e72db5414c`  
Outcome: **recovered**  
Started: 2026-10-11T16:53:45.198299+08:00  
Finished: 2026-10-11T16:54:36.667419+08:00

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
| Traffic fully drained | 2026-10-11T16:54:06.91534+08:00 |
| Immutable ledger frozen | 2026-10-11T16:54:06.91546+08:00 |
| All six stopped verified before restart | 2026-10-11T16:54:09.179641+08:00 |
| First restart requested | 2026-10-11T16:54:10.180953+08:00 |
| Restart sequence completed | 2026-10-11T16:54:19.678701+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:54:30.697339+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:54:30.69747+08:00 |

Drain to first kill request: 0.150 ms (runtime monotonic clock).

Drain to last kill completion: 1498.397 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.304 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16454.906 |
| Original topology sustained | 17633.924 |
| Replication redundancy sustained | 17633.924 |
| Post-validation diagnostic service sustained | 25028.941 |
| AOF frozen-value recovery observed | 23781.657 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 417; checked 417; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:54:06.91549+08:00 | 2026-10-11T16:54:07.16633+08:00 |  |
| 2 | node-5 | 2026-10-11T16:54:07.166339+08:00 | 2026-10-11T16:54:07.400002+08:00 |  |
| 3 | node-6 | 2026-10-11T16:54:07.400012+08:00 | 2026-10-11T16:54:07.666819+08:00 |  |
| 4 | node-1 | 2026-10-11T16:54:07.666826+08:00 | 2026-10-11T16:54:07.92417+08:00 |  |
| 5 | node-2 | 2026-10-11T16:54:07.924174+08:00 | 2026-10-11T16:54:08.165355+08:00 |  |
| 6 | node-3 | 2026-10-11T16:54:08.165362+08:00 | 2026-10-11T16:54:08.41375+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:54:10.180953+08:00 | 2026-10-11T16:54:11.74777+08:00 |  |
| 2 | node-2 | 2026-10-11T16:54:11.747773+08:00 | 2026-10-11T16:54:13.349095+08:00 |  |
| 3 | node-3 | 2026-10-11T16:54:13.349104+08:00 | 2026-10-11T16:54:14.99566+08:00 |  |
| 4 | node-4 | 2026-10-11T16:54:14.995662+08:00 | 2026-10-11T16:54:16.573613+08:00 |  |
| 5 | node-5 | 2026-10-11T16:54:16.573615+08:00 | 2026-10-11T16:54:18.105988+08:00 |  |
| 6 | node-6 | 2026-10-11T16:54:18.10599+08:00 | 2026-10-11T16:54:19.6787+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 31ae5a20099e918d6c448a631c8405b758ef9bf2 | 2026-10-11T16:54:08.537773+08:00 | false | no | managed Redis process is stopped |
| node-5 | 6ac3b761d2d030ca810736c04b67fdb74aa6c784 | 2026-10-11T16:54:08.657298+08:00 | false | no | managed Redis process is stopped |
| node-6 | ab677c48120b7e560592710dbe802d2ca7162d6b | 2026-10-11T16:54:08.780396+08:00 | false | no | managed Redis process is stopped |
| node-1 | d5170c4ddf1973cd50fae05154d96b9b0e2ddd3e | 2026-10-11T16:54:08.947229+08:00 | false | no | managed Redis process is stopped |
| node-2 | 7b7be838251cab8b2d549eb25e31413d8a8711c7 | 2026-10-11T16:54:09.066441+08:00 | false | no | managed Redis process is stopped |
| node-3 | 1352a46eb9b24bf3395a65921e58ea4100cec296 | 2026-10-11T16:54:09.179633+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `d5170c4ddf1973cd50fae05154d96b9b0e2ddd3e`. Process run ID: `32d9270e7add9aa25b787ca8b8766967f44519fd`. Observed: 2026-10-11T16:53:57.836171+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `d5170c4ddf1973cd50fae05154d96b9b0e2ddd3e`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
repl-diskless-load
swapdb
maxmemory-policy
noeviction
aof-load-truncated
yes
save

aof-use-rdb-preamble
yes
appendfsync
always
appendonly
yes
no-appendfsync-on-rewrite
no
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708828
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

### Prepared node-2

Original Redis Cluster ID: `7b7be838251cab8b2d549eb25e31413d8a8711c7`. Process run ID: `884744a22d733dbb08817c5f9d669d5280553c6a`. Observed: 2026-10-11T16:53:58.807098+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7b7be838251cab8b2d549eb25e31413d8a8711c7`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendonly
yes
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
maxmemory
134217728
maxmemory-policy
noeviction
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
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
rdb_last_save_time:1791708828
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:544768
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

Original Redis Cluster ID: `1352a46eb9b24bf3395a65921e58ea4100cec296`. Process run ID: `0dba84d0c127fe7d137393195f13d4f910a9b6d0`. Observed: 2026-10-11T16:53:59.882793+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1352a46eb9b24bf3395a65921e58ea4100cec296`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendonly
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
save

maxmemory
134217728
appendfsync
always
rdb-del-sync-files
yes
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708828
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

### Prepared node-4

Original Redis Cluster ID: `31ae5a20099e918d6c448a631c8405b758ef9bf2`. Process run ID: `8cab98fe695a4adce0fe25dd57a884c557e7ddb5`. Observed: 2026-10-11T16:54:00.898897+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `31ae5a20099e918d6c448a631c8405b758ef9bf2`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
appendonly
yes
repl-diskless-load
swapdb
save

no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
rdb-del-sync-files
yes
maxmemory
134217728
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
maxmemory-policy
noeviction
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
rdb_last_save_time:1791708829
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

Original Redis Cluster ID: `6ac3b761d2d030ca810736c04b67fdb74aa6c784`. Process run ID: `7b668628146934b8c60978c031b736df79d44aa8`. Observed: 2026-10-11T16:54:01.898098+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6ac3b761d2d030ca810736c04b67fdb74aa6c784`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendfsync
always
rdb-del-sync-files
yes
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
save

maxmemory
134217728
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
rdb_last_save_time:1791708829
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

### Prepared node-6

Original Redis Cluster ID: `ab677c48120b7e560592710dbe802d2ca7162d6b`. Process run ID: `24a980e010a894b0858bac9cc38b45c84e11f58d`. Observed: 2026-10-11T16:54:02.887071+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ab677c48120b7e560592710dbe802d2ca7162d6b`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
maxmemory
134217728
appendonly
yes
repl-diskless-sync
yes
repl-diskless-load
swapdb
aof-use-rdb-preamble
yes
appendfsync
always
rdb-del-sync-files
yes
aof-load-truncated
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708829
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

Original Redis Cluster ID: `d5170c4ddf1973cd50fae05154d96b9b0e2ddd3e`. Observed after restart: 2026-10-11T16:54:11.633808+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `d5170c4ddf1973cd50fae05154d96b9b0e2ddd3e`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
appendonly
yes
save

maxmemory
134217728
appendfsync
always
rdb-del-sync-files
yes
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708850
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

Original Redis Cluster ID: `7b7be838251cab8b2d549eb25e31413d8a8711c7`. Observed after restart: 2026-10-11T16:54:13.2285+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7b7be838251cab8b2d549eb25e31413d8a8711c7`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
save

maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
aof-load-truncated
yes
repl-diskless-sync
yes
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
appendfsync
always
appendonly
yes
repl-diskless-load
swapdb
maxmemory
134217728
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
rdb_last_save_time:1791708852
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

Original Redis Cluster ID: `1352a46eb9b24bf3395a65921e58ea4100cec296`. Observed after restart: 2026-10-11T16:54:14.877963+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1352a46eb9b24bf3395a65921e58ea4100cec296`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
save

maxmemory
134217728
repl-diskless-load
swapdb
aof-load-truncated
yes
appendfsync
always
appendonly
yes
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
rdb_last_save_time:1791708853
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

Original Redis Cluster ID: `31ae5a20099e918d6c448a631c8405b758ef9bf2`. Observed after restart: 2026-10-11T16:54:16.461801+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `31ae5a20099e918d6c448a631c8405b758ef9bf2`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
repl-diskless-load
swapdb
repl-diskless-sync
yes
appendfsync
always
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
maxmemory
134217728
save

no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708855
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
/data/temp-1791708855.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `6ac3b761d2d030ca810736c04b67fdb74aa6c784`. Observed after restart: 2026-10-11T16:54:17.990281+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6ac3b761d2d030ca810736c04b67fdb74aa6c784`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
appendfsync
always
rdb-del-sync-files
yes
maxmemory-policy
noeviction
aof-load-truncated
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
repl-diskless-load
swapdb
appendonly
yes
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708856
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
/data/temp-1791708856.1.rdb
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `ab677c48120b7e560592710dbe802d2ca7162d6b`. Observed after restart: 2026-10-11T16:54:19.564525+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ab677c48120b7e560592710dbe802d2ca7162d6b`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
appendfsync
always
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
repl-diskless-sync
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
appendonly
yes
maxmemory
134217728
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708858
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
/data/temp-1791708858.1.rdb

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `d5170c4ddf1973cd50fae05154d96b9b0e2ddd3e`. Process run ID: `793fd61abda39042a535bf129b06cf197a526b32`. Observed: 2026-10-11T16:54:25.596724+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `d5170c4ddf1973cd50fae05154d96b9b0e2ddd3e`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
appendonly
yes
save

maxmemory
134217728
appendfsync
always
rdb-del-sync-files
yes
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708850
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

Original Redis Cluster ID: `7b7be838251cab8b2d549eb25e31413d8a8711c7`. Process run ID: `7103e0fceb95a2172e566343ce038cce5ef9ed4e`. Observed: 2026-10-11T16:54:26.575783+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7b7be838251cab8b2d549eb25e31413d8a8711c7`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
save

maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
aof-load-truncated
yes
repl-diskless-sync
yes
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
appendfsync
always
appendonly
yes
repl-diskless-load
swapdb
maxmemory
134217728
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
rdb_last_save_time:1791708852
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

Original Redis Cluster ID: `1352a46eb9b24bf3395a65921e58ea4100cec296`. Process run ID: `12704e25c57b35568e5ab407fa410ea4302e22a6`. Observed: 2026-10-11T16:54:27.614576+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1352a46eb9b24bf3395a65921e58ea4100cec296`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
save

maxmemory
134217728
repl-diskless-load
swapdb
aof-load-truncated
yes
appendfsync
always
appendonly
yes
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
rdb_last_save_time:1791708853
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:897024
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

Original Redis Cluster ID: `31ae5a20099e918d6c448a631c8405b758ef9bf2`. Process run ID: `6c077872dfeb3c55a4211db93b3093aff9340165`. Observed: 2026-10-11T16:54:28.598209+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `31ae5a20099e918d6c448a631c8405b758ef9bf2`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
repl-diskless-load
swapdb
repl-diskless-sync
yes
appendfsync
always
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
maxmemory
134217728
save

no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708855
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
aof_last_cow_size:675840
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

#### Restored node-5

Original Redis Cluster ID: `6ac3b761d2d030ca810736c04b67fdb74aa6c784`. Process run ID: `2179ec104c18f86aefa5fb8302b068f6ae24cb1d`. Observed: 2026-10-11T16:54:29.582577+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6ac3b761d2d030ca810736c04b67fdb74aa6c784`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
appendfsync
always
rdb-del-sync-files
yes
maxmemory-policy
noeviction
aof-load-truncated
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
repl-diskless-load
swapdb
appendonly
yes
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708856
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
aof_last_cow_size:675840
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

#### Restored node-6

Original Redis Cluster ID: `ab677c48120b7e560592710dbe802d2ca7162d6b`. Process run ID: `9e31876775c46960ccc2202d2ed4d4953d8af17e`. Observed: 2026-10-11T16:54:30.57951+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ab677c48120b7e560592710dbe802d2ca7162d6b`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
appendfsync
always
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
repl-diskless-sync
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
appendonly
yes
maxmemory
134217728
repl-diskless-load
swapdb
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708858
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
aof_last_cow_size:733184
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

Fault action completed 2264.135 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 805 | 100.00% | 0 | 199.86 | 0.664 / 1.380 / 2.360 |

Acknowledged writes: 417. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 417 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 417 | 0.934 / 1.598 / 2.478 |
| baseline | set | 417 | 0.934 / 1.598 / 2.478 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 417 | 103.532 |

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
