# Redis Reliability Lab — persistence-memory

Run: `aa2a96b0899c2885b2aa3a32`  
Outcome: **partial**  
Started: 2026-10-11T16:51:10.03148+08:00  
Finished: 2026-10-11T16:52:00.950618+08:00

## Recovery

| Condition | Observed |
|---|---|
| Original six-node topology recovered | true |
| Post-validation diagnostic service sustained | true |
| Topology and diagnostic service jointly sustained | true |
| Frozen acknowledged writes verified | false |
| All managed processes restarted | true |
| Original replication redundancy restored | true |
| Restoration deadline expired | false |
| Diagnostic deadline expired | false |
| Cleanup verified | true |


## Frozen-ledger process restart

Configured persistence mode: `memory`. Failure model: `sequential SIGKILL of all six Redis processes, same-container restart; host and storage remain running`.

Workload traffic stops and drains before the sequential SIGKILL actions. No workload requests are offered during fault or restart observation; these phases do not measure zero availability or recovery throughput. Read-only frozen-ledger validation precedes separate diagnostic writes.

Frozen acknowledged SETs: 416. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:51:31.062182+08:00 |
| Immutable ledger frozen | 2026-10-11T16:51:31.062211+08:00 |
| All six stopped verified before restart | 2026-10-11T16:51:33.282098+08:00 |
| First restart requested | 2026-10-11T16:51:34.283658+08:00 |
| Restart sequence completed | 2026-10-11T16:51:43.968463+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:51:54.862836+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:51:54.862964+08:00 |

Drain to first kill request: 0.045 ms (runtime monotonic clock).

Drain to last kill completion: 1504.372 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.552 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16375.412 |
| Original topology sustained | 17602.125 |
| Replication redundancy sustained | 17602.125 |
| Post-validation diagnostic service sustained | 25047.236 |
| AOF frozen-value recovery observed | unobserved |

Configured memory-only control disables RDB/AOF value persistence. Persistence recovery is not applicable; restored processes and topology can coexist with lost frozen values. The independent persistence dimension is false when the disabled runtime configuration was observed, while the persistence-recovery event remains absent.

Frozen-ledger validation: expected 416; checked 416; missing 416; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:51:31.062227+08:00 | 2026-10-11T16:51:31.306114+08:00 |  |
| 2 | node-5 | 2026-10-11T16:51:31.306129+08:00 | 2026-10-11T16:51:31.563889+08:00 |  |
| 3 | node-6 | 2026-10-11T16:51:31.563901+08:00 | 2026-10-11T16:51:31.815692+08:00 |  |
| 4 | node-1 | 2026-10-11T16:51:31.815702+08:00 | 2026-10-11T16:51:32.067092+08:00 |  |
| 5 | node-2 | 2026-10-11T16:51:32.067095+08:00 | 2026-10-11T16:51:32.30758+08:00 |  |
| 6 | node-3 | 2026-10-11T16:51:32.307586+08:00 | 2026-10-11T16:51:32.566566+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:51:34.283658+08:00 | 2026-10-11T16:51:35.940022+08:00 |  |
| 2 | node-2 | 2026-10-11T16:51:35.940024+08:00 | 2026-10-11T16:51:37.560562+08:00 |  |
| 3 | node-3 | 2026-10-11T16:51:37.560565+08:00 | 2026-10-11T16:51:39.225153+08:00 |  |
| 4 | node-4 | 2026-10-11T16:51:39.225157+08:00 | 2026-10-11T16:51:40.815089+08:00 |  |
| 5 | node-5 | 2026-10-11T16:51:40.81509+08:00 | 2026-10-11T16:51:42.396101+08:00 |  |
| 6 | node-6 | 2026-10-11T16:51:42.396102+08:00 | 2026-10-11T16:51:43.968462+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 1ff08c5910353237fccd0e38113efce553830763 | 2026-10-11T16:51:32.690349+08:00 | false | no | managed Redis process is stopped |
| node-5 | 98ef7673d7aeace6a9797164e44269a6f8641a99 | 2026-10-11T16:51:32.809807+08:00 | false | no | managed Redis process is stopped |
| node-6 | ff8e163493df323fcbe87d9a3763a1b114cd76cc | 2026-10-11T16:51:32.929871+08:00 | false | no | managed Redis process is stopped |
| node-1 | 60da8323e592f98015673eddf39dfae143b2a91c | 2026-10-11T16:51:33.048922+08:00 | false | no | managed Redis process is stopped |
| node-2 | 39826df6f139e86b2404ff37e026e69153ce0c7f | 2026-10-11T16:51:33.166657+08:00 | false | no | managed Redis process is stopped |
| node-3 | 24ef51786d5fb18310454148d5f42476a7d0632d | 2026-10-11T16:51:33.282093+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `60da8323e592f98015673eddf39dfae143b2a91c`. Process run ID: `bcbf7a46db891f5b4331593540600e22aaf1c1f0`. Observed: 2026-10-11T16:51:22.00262+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `60da8323e592f98015673eddf39dfae143b2a91c`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
appendonly
no
repl-diskless-sync
yes
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
rdb-del-sync-files
yes
appendfsync
everysec
aof-load-truncated
yes
save

aof-use-rdb-preamble
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
rdb_last_save_time:1791708672
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:507904
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

### Prepared node-2

Original Redis Cluster ID: `39826df6f139e86b2404ff37e026e69153ce0c7f`. Process run ID: `a1ff15f3f39f8652876471b8a2b714c5a1793b7f`. Observed: 2026-10-11T16:51:23.003067+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `39826df6f139e86b2404ff37e026e69153ce0c7f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
aof-use-rdb-preamble
no
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
save

appendfsync
everysec
appendonly
no
repl-diskless-sync
yes
maxmemory
134217728
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
rdb_last_save_time:1791708673
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:663552
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

### Prepared node-3

Original Redis Cluster ID: `24ef51786d5fb18310454148d5f42476a7d0632d`. Process run ID: `96646429820b44a684276b03fb014f2841c81e7e`. Observed: 2026-10-11T16:51:24.0373+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `24ef51786d5fb18310454148d5f42476a7d0632d`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
appendfsync
everysec
maxmemory-policy
noeviction
repl-diskless-load
swapdb
rdb-del-sync-files
yes
maxmemory
134217728
aof-use-rdb-preamble
no
appendonly
no
save

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
rdb_last_save_time:1791708673
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:552960
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

### Prepared node-4

Original Redis Cluster ID: `1ff08c5910353237fccd0e38113efce553830763`. Process run ID: `6b0b9ab5e3d3953853a59a4cb7e8270105b89721`. Observed: 2026-10-11T16:51:25.000856+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1ff08c5910353237fccd0e38113efce553830763`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
appendonly
no
rdb-del-sync-files
yes
aof-load-truncated
yes
appendfsync
everysec
aof-use-rdb-preamble
no
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708673
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

### Prepared node-5

Original Redis Cluster ID: `98ef7673d7aeace6a9797164e44269a6f8641a99`. Process run ID: `9ee9acc74f70f8261c9381026a39d8500ecc808b`. Observed: 2026-10-11T16:51:26.000337+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `98ef7673d7aeace6a9797164e44269a6f8641a99`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
aof-load-truncated
yes
appendfsync
everysec
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
appendonly
no
maxmemory-policy
noeviction
repl-diskless-load
swapdb
maxmemory
134217728
aof-use-rdb-preamble
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
rdb_last_save_time:1791708674
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

### Prepared node-6

Original Redis Cluster ID: `ff8e163493df323fcbe87d9a3763a1b114cd76cc`. Process run ID: `6e0bdd5b9913553467c6efb957b4ea72f11f6256`. Observed: 2026-10-11T16:51:27.033393+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ff8e163493df323fcbe87d9a3763a1b114cd76cc`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
no
aof-load-truncated
yes
maxmemory
134217728
repl-diskless-load
swapdb
maxmemory-policy
noeviction
appendonly
no
rdb-del-sync-files
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
appendfsync
everysec
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
rdb_last_save_time:1791708674
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

### Restart receipt 1: node-1

Original Redis Cluster ID: `60da8323e592f98015673eddf39dfae143b2a91c`. Observed after restart: 2026-10-11T16:51:35.820709+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `60da8323e592f98015673eddf39dfae143b2a91c`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
repl-diskless-sync
yes
maxmemory-policy
noeviction
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
appendfsync
everysec
aof-use-rdb-preamble
no
repl-diskless-load
swapdb
maxmemory
134217728
appendonly
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
rdb_last_save_time:1791708694
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

### Restart receipt 2: node-2

Original Redis Cluster ID: `39826df6f139e86b2404ff37e026e69153ce0c7f`. Observed after restart: 2026-10-11T16:51:37.432289+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `39826df6f139e86b2404ff37e026e69153ce0c7f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
appendonly
no
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
repl-diskless-sync
yes
repl-diskless-load
swapdb
save

rdb-del-sync-files
yes
aof-load-truncated
yes
appendfsync
everysec
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
rdb_last_save_time:1791708696
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

### Restart receipt 3: node-3

Original Redis Cluster ID: `24ef51786d5fb18310454148d5f42476a7d0632d`. Observed after restart: 2026-10-11T16:51:39.105222+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `24ef51786d5fb18310454148d5f42476a7d0632d`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
appendfsync
everysec
aof-load-truncated
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
repl-diskless-sync
yes
save

auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
appendonly
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
rdb_last_save_time:1791708697
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

### Restart receipt 4: node-4

Original Redis Cluster ID: `1ff08c5910353237fccd0e38113efce553830763`. Observed after restart: 2026-10-11T16:51:40.695779+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1ff08c5910353237fccd0e38113efce553830763`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
appendonly
no
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
maxmemory
134217728
rdb-del-sync-files
yes
appendfsync
everysec
repl-diskless-load
swapdb
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791708699
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf
/data/temp-1791708699.1.rdb

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `98ef7673d7aeace6a9797164e44269a6f8641a99`. Observed after restart: 2026-10-11T16:51:42.281133+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `98ef7673d7aeace6a9797164e44269a6f8641a99`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
appendfsync
everysec
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
save

auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
appendonly
no
maxmemory-policy
noeviction
aof-load-truncated
yes
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791708701
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf
/data/temp-1791708701.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `ff8e163493df323fcbe87d9a3763a1b114cd76cc`. Observed after restart: 2026-10-11T16:51:43.854871+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ff8e163493df323fcbe87d9a3763a1b114cd76cc`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
no-appendfsync-on-rewrite
no
save

appendonly
no
maxmemory
134217728
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
maxmemory-policy
noeviction
rdb-del-sync-files
yes
repl-diskless-load
swapdb
aof-use-rdb-preamble
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
rdb_last_save_time:1791708702
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/temp-1791708702.1.rdb
/data/nodes.conf

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `60da8323e592f98015673eddf39dfae143b2a91c`. Process run ID: `0cbe324940eaa795909b368dacc8ef1af5545120`. Observed: 2026-10-11T16:51:49.732498+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `60da8323e592f98015673eddf39dfae143b2a91c`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
repl-diskless-sync
yes
maxmemory-policy
noeviction
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
appendfsync
everysec
aof-use-rdb-preamble
no
repl-diskless-load
swapdb
maxmemory
134217728
appendonly
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
rdb_last_save_time:1791708694
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:487424
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

#### Restored node-2

Original Redis Cluster ID: `39826df6f139e86b2404ff37e026e69153ce0c7f`. Process run ID: `7f37f2a8884f2edbd1c420cbfe5b0752c30d2709`. Observed: 2026-10-11T16:51:50.77338+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `39826df6f139e86b2404ff37e026e69153ce0c7f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
appendonly
no
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
repl-diskless-sync
yes
repl-diskless-load
swapdb
save

rdb-del-sync-files
yes
aof-load-truncated
yes
appendfsync
everysec
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
rdb_last_save_time:1791708696
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:458752
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

#### Restored node-3

Original Redis Cluster ID: `24ef51786d5fb18310454148d5f42476a7d0632d`. Process run ID: `57e05eb48de996ac65dcec8612b1fac4ab668512`. Observed: 2026-10-11T16:51:51.796632+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `24ef51786d5fb18310454148d5f42476a7d0632d`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
appendfsync
everysec
aof-load-truncated
yes
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
repl-diskless-sync
yes
save

auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
appendonly
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
rdb_last_save_time:1791708697
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:458752
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

#### Restored node-4

Original Redis Cluster ID: `1ff08c5910353237fccd0e38113efce553830763`. Process run ID: `bc6672b1bcbcf13fceae3aaeed53ea16268a1a6a`. Observed: 2026-10-11T16:51:52.783774+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1ff08c5910353237fccd0e38113efce553830763`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
appendonly
no
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
maxmemory
134217728
rdb-del-sync-files
yes
appendfsync
everysec
repl-diskless-load
swapdb
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791708699
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

#### Restored node-5

Original Redis Cluster ID: `98ef7673d7aeace6a9797164e44269a6f8641a99`. Process run ID: `5a75df237fe15aa886ef89b79248246819b45725`. Observed: 2026-10-11T16:51:53.733272+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `98ef7673d7aeace6a9797164e44269a6f8641a99`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
appendfsync
everysec
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
save

auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
appendonly
no
maxmemory-policy
noeviction
aof-load-truncated
yes
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791708701
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

#### Restored node-6

Original Redis Cluster ID: `ff8e163493df323fcbe87d9a3763a1b114cd76cc`. Process run ID: `1d89f016b81b150a1a7626505ae0c76b7bc48b85`. Observed: 2026-10-11T16:51:54.748141+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ff8e163493df323fcbe87d9a3763a1b114cd76cc`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
no-appendfsync-on-rewrite
no
save

appendonly
no
maxmemory
134217728
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
maxmemory-policy
noeviction
rdb-del-sync-files
yes
repl-diskless-load
swapdb
aof-use-rdb-preamble
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
rdb_last_save_time:1791708702
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:-1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:0
rdb_last_load_keys_expired:0
rdb_last_load_keys_loaded:0
aof_enabled:0
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
```

Observed INFO keyspace:

```text
# Keyspace
```

Observed managed data files:

```text
/data/nodes.conf

```

This is a sequential process crash and same-container restart with retained managed volumes. It does not remove Docker VM/host page cache or simulate power loss, storage loss, or a graceful shutdown. Draining and sequential kill/inspection time allow background fsync to progress, so everysec preservation does not establish a one-second-loss guarantee or equivalence to always.

Persistence experiment error: persistence experiment ended partial: frozen data verified=false, original redundancy stable=true, restoration deadline=false

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
| Acknowledged data verified | false |
| Configured write policy available | true |
| Replication redundancy restored | true |
| Persistence recovery verified | false |

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 2219.856 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 804 | 100.00% | 0 | 199.57 | 0.428 / 0.697 / 1.449 |

Acknowledged writes: 416. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 416 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 416 | 0.449 / 0.732 / 1.521 |
| baseline | set | 416 | 0.428 / 0.732 / 1.521 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 416 | 103.263 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 805; started 804; completed 804; dropped 1. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 416; checked: 416; missing: 416; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Errors

- persistence experiment ended partial: frozen data verified=false, original redundancy stable=true, restoration deadline=false

## Measurement methodology

- Quiet administrative process, CONFIG, persistence INFO, keyspace and bounded data-file preparation occurs before baseline traffic; it issues no SAVE, BGSAVE, BGREWRITEAOF or explicit durability barrier.
- The workload fully stops and joins in-flight acknowledgements before freezing its immutable ledger and counters. All later topology observations are read-only until every frozen acknowledged value has been checked. SET-error outcomes remain ambiguous, excluded from acknowledgement guarantees.
- Original replicas are SIGKILLed before original primaries. Every fixed owned container must produce ErrProcessStopped administrative evidence with restart policy no before any restart. Original primaries restart before replicas in the same containers and retained volumes; original Redis Cluster identities and slot ownership must survive.
- The model is sequential process SIGKILL while Docker VM, host and storage remain running, followed by same-container starts. It does not simulate power loss, host/page-cache loss or disk loss. Drain, kill and inspection intervals are reported because they permit background fsync progress.
- Topology restoration requires all six original identities, endpoints, primary slot owners and replica assignments synchronized for stable_for, without forced Redis topology commands. Physical receipts retain early raw evidence, which can be transient during full sync. Six fresh quiet administrative runtime/configuration/file observations after topology stability prove the final persistence mode on those same restarted processes, within the shared repair budget and before frozen-value validation. Samples crossing restoration or diagnostic deadlines are retained but excluded from successful milestones.
- Frozen-ledger GET validation happens exactly once before diagnostic SET/GET or configured replica-acknowledgement probes. Process/slot readiness and later diagnostic service are independent from frozen-value preservation; missing or corrupt values retain a partial data verdict even after service recovers.
- Memory mode has RDB/AOF disabled, rdb-del-sync-files=yes to remove replication-created RDB files, and observed absence of value persistence files. Its persistence-recovery event is not applicable and remains absent; its independent persistence dimension is false. AOF modes use the RDB base preamble required by the pinned image's bundled modules plus incremental AOF records. AOF recovery requires actual matching healthy runtime configuration, loading=0 and complete frozen acknowledged-value preservation; this observation establishes no loss-free or power-loss guarantee.
- No workload demand is offered after draining, so absent fault/recovery traffic does not measure zero availability or throughput. Post-validation probes are separate diagnostics and never enter workload metrics or the frozen ledger. Poll intervals follow sample completion; observations are sampled and non-atomic.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `359df9c4ebd051da3e82059877c38c43334b012d0ec20f156604a3c2447ffe2a`.
