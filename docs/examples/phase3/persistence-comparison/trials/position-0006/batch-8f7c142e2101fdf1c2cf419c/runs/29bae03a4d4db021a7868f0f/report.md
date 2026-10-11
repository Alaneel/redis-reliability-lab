# Redis Reliability Lab — persistence-aof-everysec

Run: `29bae03a4d4db021a7868f0f`  
Outcome: **recovered**  
Started: 2026-10-11T16:54:36.841165+08:00  
Finished: 2026-10-11T16:55:27.75971+08:00

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
| Traffic fully drained | 2026-10-11T16:54:58.159132+08:00 |
| Immutable ledger frozen | 2026-10-11T16:54:58.159159+08:00 |
| All six stopped verified before restart | 2026-10-11T16:55:00.345433+08:00 |
| First restart requested | 2026-10-11T16:55:01.346871+08:00 |
| Restart sequence completed | 2026-10-11T16:55:10.75801+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:55:21.759744+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:55:21.759884+08:00 |

Drain to first kill request: 0.042 ms (runtime monotonic clock).

Drain to last kill completion: 1482.832 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.429 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16340.912 |
| Original topology sustained | 17585.182 |
| Replication redundancy sustained | 17585.182 |
| Post-validation diagnostic service sustained | 24855.190 |
| AOF frozen-value recovery observed | 23600.379 |

AOF frozen-value recovery observed: true. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation: expected 417; checked 417; missing 0; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:54:58.159174+08:00 | 2026-10-11T16:54:58.398819+08:00 |  |
| 2 | node-5 | 2026-10-11T16:54:58.398827+08:00 | 2026-10-11T16:54:58.651059+08:00 |  |
| 3 | node-6 | 2026-10-11T16:54:58.651077+08:00 | 2026-10-11T16:54:58.898412+08:00 |  |
| 4 | node-1 | 2026-10-11T16:54:58.898418+08:00 | 2026-10-11T16:54:59.150566+08:00 |  |
| 5 | node-2 | 2026-10-11T16:54:59.15057+08:00 | 2026-10-11T16:54:59.394555+08:00 |  |
| 6 | node-3 | 2026-10-11T16:54:59.394559+08:00 | 2026-10-11T16:54:59.641976+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:55:01.346871+08:00 | 2026-10-11T16:55:02.942919+08:00 |  |
| 2 | node-2 | 2026-10-11T16:55:02.94292+08:00 | 2026-10-11T16:55:04.467913+08:00 |  |
| 3 | node-3 | 2026-10-11T16:55:04.46792+08:00 | 2026-10-11T16:55:06.05301+08:00 |  |
| 4 | node-4 | 2026-10-11T16:55:06.053012+08:00 | 2026-10-11T16:55:07.614458+08:00 |  |
| 5 | node-5 | 2026-10-11T16:55:07.614459+08:00 | 2026-10-11T16:55:09.194882+08:00 |  |
| 6 | node-6 | 2026-10-11T16:55:09.194887+08:00 | 2026-10-11T16:55:10.758008+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | e23102cc8d62389dd829be895fbeb6317867563c | 2026-10-11T16:54:59.768247+08:00 | false | no | managed Redis process is stopped |
| node-5 | 607266574aca65963c44c856b4fd4b5d54fe3717 | 2026-10-11T16:54:59.884405+08:00 | false | no | managed Redis process is stopped |
| node-6 | 8506326a3b7e749e87f2f4c93c337ee48bcd74b1 | 2026-10-11T16:54:59.997727+08:00 | false | no | managed Redis process is stopped |
| node-1 | ae9a021f3d7bb615e1f647d55fcd19712535a12e | 2026-10-11T16:55:00.112214+08:00 | false | no | managed Redis process is stopped |
| node-2 | 885fad1ab78ea2d232380c0dc5610d5b6cffb791 | 2026-10-11T16:55:00.233617+08:00 | false | no | managed Redis process is stopped |
| node-3 | 92f27955e9993e78992a7049208f42a2e00c144c | 2026-10-11T16:55:00.345425+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `ae9a021f3d7bb615e1f647d55fcd19712535a12e`. Process run ID: `9b76a7f4c21d236da9eee6c9ffe04a279756bbc2`. Observed: 2026-10-11T16:54:49.104593+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ae9a021f3d7bb615e1f647d55fcd19712535a12e`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
aof-load-truncated
yes
maxmemory
134217728
save

appendonly
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
appendfsync
everysec
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
rdb_last_save_time:1791708879
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

### Prepared node-2

Original Redis Cluster ID: `885fad1ab78ea2d232380c0dc5610d5b6cffb791`. Process run ID: `7bad8e32d7b3a624903bab92654bfb988e81a597`. Observed: 2026-10-11T16:54:50.137604+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `885fad1ab78ea2d232380c0dc5610d5b6cffb791`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
repl-diskless-sync
yes
appendonly
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
save

aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
appendfsync
everysec
maxmemory
134217728
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
rdb_last_save_time:1791708879
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:569344
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

Original Redis Cluster ID: `92f27955e9993e78992a7049208f42a2e00c144c`. Process run ID: `be231ea56da0cad973638dc783431616c86205bf`. Observed: 2026-10-11T16:54:51.135802+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `92f27955e9993e78992a7049208f42a2e00c144c`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
appendonly
yes
rdb-del-sync-files
yes
repl-diskless-load
swapdb
maxmemory-policy
noeviction
aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
maxmemory
134217728
save

auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
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
rdb_last_save_time:1791708880
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:548864
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

Original Redis Cluster ID: `e23102cc8d62389dd829be895fbeb6317867563c`. Process run ID: `4a98fa0770f8b6c2c63d5bf766a6a736eba69ebd`. Observed: 2026-10-11T16:54:52.10928+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e23102cc8d62389dd829be895fbeb6317867563c`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
appendonly
yes
auto-aof-rewrite-percentage
0
appendfsync
everysec
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
yes
rdb-del-sync-files
yes
repl-diskless-load
swapdb
maxmemory
134217728
repl-diskless-sync
yes
maxmemory-policy
noeviction
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
rdb_last_save_time:1791708880
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

### Prepared node-5

Original Redis Cluster ID: `607266574aca65963c44c856b4fd4b5d54fe3717`. Process run ID: `55783eb9b739e9b8d0ef6a2d43ad657cf8368c81`. Observed: 2026-10-11T16:54:53.124397+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `607266574aca65963c44c856b4fd4b5d54fe3717`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
rdb-del-sync-files
yes
repl-diskless-load
swapdb
aof-load-truncated
yes
save

auto-aof-rewrite-percentage
0
appendonly
yes
maxmemory
134217728
appendfsync
everysec
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
rdb_last_save_time:1791708880
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

Original Redis Cluster ID: `8506326a3b7e749e87f2f4c93c337ee48bcd74b1`. Process run ID: `89f17c058d1839598f97125fd784065d4504274e`. Observed: 2026-10-11T16:54:54.13497+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `8506326a3b7e749e87f2f4c93c337ee48bcd74b1`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
aof-use-rdb-preamble
yes
repl-diskless-sync
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
appendfsync
everysec
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
save

repl-diskless-load
swapdb
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
rdb_last_save_time:1791708881
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

Original Redis Cluster ID: `ae9a021f3d7bb615e1f647d55fcd19712535a12e`. Observed after restart: 2026-10-11T16:55:02.830053+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ae9a021f3d7bb615e1f647d55fcd19712535a12e`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
save

aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
appendfsync
everysec
appendonly
yes
maxmemory-policy
noeviction
aof-load-truncated
yes
repl-diskless-sync
yes
maxmemory
134217728
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
rdb_last_save_time:1791708901
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

Original Redis Cluster ID: `885fad1ab78ea2d232380c0dc5610d5b6cffb791`. Observed after restart: 2026-10-11T16:55:04.356958+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `885fad1ab78ea2d232380c0dc5610d5b6cffb791`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
appendonly
yes
auto-aof-rewrite-percentage
0
save

no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
maxmemory-policy
noeviction
rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-load
swapdb
aof-load-truncated
yes
appendfsync
everysec
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
rdb_last_save_time:1791708903
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

Original Redis Cluster ID: `92f27955e9993e78992a7049208f42a2e00c144c`. Observed after restart: 2026-10-11T16:55:05.936484+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `92f27955e9993e78992a7049208f42a2e00c144c`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
appendfsync
everysec
repl-diskless-sync
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
yes
aof-load-truncated
yes
repl-diskless-load
swapdb
appendonly
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708904
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

Original Redis Cluster ID: `e23102cc8d62389dd829be895fbeb6317867563c`. Observed after restart: 2026-10-11T16:55:07.499409+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e23102cc8d62389dd829be895fbeb6317867563c`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
save

repl-diskless-sync
yes
appendfsync
everysec
appendonly
yes
maxmemory
134217728
maxmemory-policy
noeviction
rdb-del-sync-files
yes
repl-diskless-load
swapdb
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708906
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
/data/temp-1791708906.1.rdb
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.2.base.rdb
/data/appendonlydir/appendonly.aof.2.incr.aof

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `607266574aca65963c44c856b4fd4b5d54fe3717`. Observed after restart: 2026-10-11T16:55:09.07677+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `607266574aca65963c44c856b4fd4b5d54fe3717`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-load
swapdb
repl-diskless-sync
yes
appendonly
yes
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708907
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
/data/temp-1791708907.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `8506326a3b7e749e87f2f4c93c337ee48bcd74b1`. Observed after restart: 2026-10-11T16:55:10.642238+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `8506326a3b7e749e87f2f4c93c337ee48bcd74b1`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
aof-load-truncated
yes
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
appendonly
yes
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
maxmemory
134217728
save

aof-use-rdb-preamble
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708909
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
/data/temp-1791708909.1.rdb

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `ae9a021f3d7bb615e1f647d55fcd19712535a12e`. Process run ID: `8e0ed4a71bf86fd0f6cbc6e4a49693194c6b3c01`. Observed: 2026-10-11T16:55:16.758054+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ae9a021f3d7bb615e1f647d55fcd19712535a12e`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
save

aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
appendfsync
everysec
appendonly
yes
maxmemory-policy
noeviction
aof-load-truncated
yes
repl-diskless-sync
yes
maxmemory
134217728
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
rdb_last_save_time:1791708901
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:811008
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

Original Redis Cluster ID: `885fad1ab78ea2d232380c0dc5610d5b6cffb791`. Process run ID: `0208de039f5cc1300e1ca5c4a2e45279c101f975`. Observed: 2026-10-11T16:55:17.750052+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `885fad1ab78ea2d232380c0dc5610d5b6cffb791`. Restart policy: `no`. `/data` usage: 56 KiB.

Observed persistence configuration:

```text
appendonly
yes
auto-aof-rewrite-percentage
0
save

no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
maxmemory-policy
noeviction
rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-load
swapdb
aof-load-truncated
yes
appendfsync
everysec
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
rdb_last_save_time:1791708903
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

#### Restored node-3

Original Redis Cluster ID: `92f27955e9993e78992a7049208f42a2e00c144c`. Process run ID: `7a433796515c354352e3f03032db33e9ba79141b`. Observed: 2026-10-11T16:55:18.739148+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `92f27955e9993e78992a7049208f42a2e00c144c`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
appendfsync
everysec
repl-diskless-sync
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
yes
aof-load-truncated
yes
repl-diskless-load
swapdb
appendonly
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708904
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

Original Redis Cluster ID: `e23102cc8d62389dd829be895fbeb6317867563c`. Process run ID: `93b2d441cdf4de16d3eb8589c19a514dd9786f07`. Observed: 2026-10-11T16:55:19.687394+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e23102cc8d62389dd829be895fbeb6317867563c`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
save

repl-diskless-sync
yes
appendfsync
everysec
appendonly
yes
maxmemory
134217728
maxmemory-policy
noeviction
rdb-del-sync-files
yes
repl-diskless-load
swapdb
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708906
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
aof_last_cow_size:647168
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20050
aof_base_size:20050
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

Original Redis Cluster ID: `607266574aca65963c44c856b4fd4b5d54fe3717`. Process run ID: `4c89c2efd289d05d35d380614c7115d7dcf843f8`. Observed: 2026-10-11T16:55:20.654422+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `607266574aca65963c44c856b4fd4b5d54fe3717`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-load
swapdb
repl-diskless-sync
yes
appendonly
yes
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
yes
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708907
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
aof_last_cow_size:712704
backup_in_progress:0
module_fork_in_progress:0
module_fork_last_cow_size:0
aof_current_size:20042
aof_base_size:20042
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

Original Redis Cluster ID: `8506326a3b7e749e87f2f4c93c337ee48bcd74b1`. Process run ID: `67df0b346d3d62c5f289cb42fd9f4c26e3560d91`. Observed: 2026-10-11T16:55:21.627602+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `8506326a3b7e749e87f2f4c93c337ee48bcd74b1`. Restart policy: `no`. `/data` usage: 36 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
aof-load-truncated
yes
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
appendonly
yes
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
maxmemory
134217728
save

aof-use-rdb-preamble
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
rdb_changes_since_last_save:139
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708909
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
aof_current_size:20044
aof_base_size:20044
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

Fault action completed 2186.243 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 804 | 100.00% | 0 | 199.81 | 0.520 / 0.732 / 0.981 |

Acknowledged writes: 417. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 417 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 417 | 0.546 / 0.768 / 0.981 |
| baseline | set | 417 | 0.546 / 0.768 / 0.981 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 417 | 103.631 |

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
