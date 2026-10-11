# Redis Reliability Lab — dataset-trial

Run: `d4afc5196b21a7355b5de2ba`  
Outcome: **partial**  
Started: 2026-10-11T19:58:07.094564+08:00  
Finished: 2026-10-11T19:59:01.367074+08:00

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

Frozen acknowledged SETs: 191. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T19:58:27.72236+08:00 |
| Immutable ledger frozen | 2026-10-11T19:58:27.722508+08:00 |
| All six stopped verified before restart | 2026-10-11T19:58:30.095207+08:00 |
| First restart requested | 2026-10-11T19:58:34.096555+08:00 |
| Restart sequence completed | 2026-10-11T19:58:44.333496+08:00 |
| Frozen-ledger validation completed | 2026-10-11T19:58:55.983629+08:00 |
| Post-validation diagnostics requested | 2026-10-11T19:58:55.98401+08:00 |

Drain to first kill request: 0.206 ms (runtime monotonic clock).

Drain to last kill completion: 1567.854 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 4001.313 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 19656.829 |
| Original topology sustained | 21135.638 |
| Replication redundancy sustained | 21135.638 |
| Post-validation diagnostic service sustained | 29450.837 |
| AOF frozen-value recovery observed | unobserved |

Configured memory-only control disables RDB/AOF value persistence. Persistence recovery is not applicable; restored processes and topology can coexist with lost frozen values. The independent persistence dimension is false when the disabled runtime configuration was observed, while the persistence-recovery event remains absent.

Frozen-ledger validation: expected 191; checked 191; missing 191; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T19:58:27.722566+08:00 | 2026-10-11T19:58:27.980245+08:00 |  |
| 2 | node-5 | 2026-10-11T19:58:27.980263+08:00 | 2026-10-11T19:58:28.251144+08:00 |  |
| 3 | node-6 | 2026-10-11T19:58:28.251153+08:00 | 2026-10-11T19:58:28.510652+08:00 |  |
| 4 | node-1 | 2026-10-11T19:58:28.510665+08:00 | 2026-10-11T19:58:28.761983+08:00 |  |
| 5 | node-2 | 2026-10-11T19:58:28.761992+08:00 | 2026-10-11T19:58:29.04074+08:00 |  |
| 6 | node-3 | 2026-10-11T19:58:29.040758+08:00 | 2026-10-11T19:58:29.290227+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T19:58:34.096555+08:00 | 2026-10-11T19:58:35.749381+08:00 |  |
| 2 | node-2 | 2026-10-11T19:58:35.749387+08:00 | 2026-10-11T19:58:37.398951+08:00 |  |
| 3 | node-3 | 2026-10-11T19:58:37.39896+08:00 | 2026-10-11T19:58:39.025046+08:00 |  |
| 4 | node-4 | 2026-10-11T19:58:39.025057+08:00 | 2026-10-11T19:58:40.73416+08:00 |  |
| 5 | node-5 | 2026-10-11T19:58:40.734163+08:00 | 2026-10-11T19:58:42.57696+08:00 |  |
| 6 | node-6 | 2026-10-11T19:58:42.57697+08:00 | 2026-10-11T19:58:44.333492+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | e925eb9e29de6393ceec2a30056989de4c6f770a | 2026-10-11T19:58:29.424385+08:00 | false | no | managed Redis process is stopped |
| node-5 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | 2026-10-11T19:58:29.554962+08:00 | false | no | managed Redis process is stopped |
| node-6 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | 2026-10-11T19:58:29.684109+08:00 | false | no | managed Redis process is stopped |
| node-1 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | 2026-10-11T19:58:29.842341+08:00 | false | no | managed Redis process is stopped |
| node-2 | 7f385e1d484746697faa90765e66e2563b95f1b7 | 2026-10-11T19:58:29.971199+08:00 | false | no | managed Redis process is stopped |
| node-3 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | 2026-10-11T19:58:30.095163+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `4f9cdefa31d52b1ec07962fe48e90304775f32b4`. Process run ID: `f248eba6dc40b818bb533b1fb363d02c650196ea`. Observed: 2026-10-11T19:58:20.396721+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4f9cdefa31d52b1ec07962fe48e90304775f32b4`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
aof-load-truncated
yes
appendfsync
everysec
repl-diskless-load
swapdb
maxmemory
134217728
aof-use-rdb-preamble
no
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
maxmemory-policy
noeviction
appendonly
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
rdb_last_save_time:1791719890
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:450560
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

Original Redis Cluster ID: `7f385e1d484746697faa90765e66e2563b95f1b7`. Process run ID: `97f3f18a58cf2a66352f733bd69e5b844e6c4833`. Observed: 2026-10-11T19:58:21.428475+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7f385e1d484746697faa90765e66e2563b95f1b7`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
appendonly
no
aof-load-truncated
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
repl-diskless-sync
yes
save

aof-use-rdb-preamble
no
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791719890
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:524288
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

Original Redis Cluster ID: `3017030ae4cc288de4885b1c31d4086bf0a2037f`. Process run ID: `57597d77c843fa2915c215c342dbf1155b2a8b2c`. Observed: 2026-10-11T19:58:22.466255+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3017030ae4cc288de4885b1c31d4086bf0a2037f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
appendfsync
everysec
aof-use-rdb-preamble
no
maxmemory
134217728
rdb-del-sync-files
yes
repl-diskless-sync
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
appendonly
no
save

auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791719890
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:483328
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

Original Redis Cluster ID: `e925eb9e29de6393ceec2a30056989de4c6f770a`. Process run ID: `33d6ef1b7228021a8f2b12f4838ed6addf8b28a0`. Observed: 2026-10-11T19:58:23.526647+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e925eb9e29de6393ceec2a30056989de4c6f770a`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
save

rdb-del-sync-files
yes
appendfsync
everysec
aof-load-truncated
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
appendonly
no
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
rdb_last_save_time:1791719891
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

Original Redis Cluster ID: `361fdf9ad3ee6577fd53ca9afbf430d205fe6c30`. Process run ID: `719a05f5f1d613658f2725bc95ff9ba9c187811c`. Observed: 2026-10-11T19:58:24.55142+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `361fdf9ad3ee6577fd53ca9afbf430d205fe6c30`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
rdb-del-sync-files
yes
maxmemory
134217728
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
repl-diskless-sync
yes
aof-use-rdb-preamble
no
appendonly
no
repl-diskless-load
swapdb
save

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
rdb_last_save_time:1791719891
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

Original Redis Cluster ID: `1cfbe152fc4e94c0102c9e95c8a3e262e71872bc`. Process run ID: `98c8fab60b30181e3b4b32f39cc466bef7a722cd`. Observed: 2026-10-11T19:58:25.575179+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1cfbe152fc4e94c0102c9e95c8a3e262e71872bc`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
save

maxmemory
134217728
aof-load-truncated
yes
appendfsync
everysec
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
rdb_last_save_time:1791719891
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

Original Redis Cluster ID: `4f9cdefa31d52b1ec07962fe48e90304775f32b4`. Observed after restart: 2026-10-11T19:58:35.624347+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4f9cdefa31d52b1ec07962fe48e90304775f32b4`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
aof-use-rdb-preamble
no
maxmemory-policy
noeviction
rdb-del-sync-files
yes
aof-load-truncated
yes
save

repl-diskless-load
swapdb
appendfsync
everysec
maxmemory
134217728
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791719914
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

Original Redis Cluster ID: `7f385e1d484746697faa90765e66e2563b95f1b7`. Observed after restart: 2026-10-11T19:58:37.269352+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7f385e1d484746697faa90765e66e2563b95f1b7`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
rdb-del-sync-files
yes
aof-load-truncated
yes
appendfsync
everysec
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
save

no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
maxmemory
134217728
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791719916
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

Original Redis Cluster ID: `3017030ae4cc288de4885b1c31d4086bf0a2037f`. Observed after restart: 2026-10-11T19:58:38.901879+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3017030ae4cc288de4885b1c31d4086bf0a2037f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
maxmemory
134217728
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
repl-diskless-load
swapdb
save

aof-load-truncated
yes
appendonly
no
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791719917
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

Original Redis Cluster ID: `e925eb9e29de6393ceec2a30056989de4c6f770a`. Observed after restart: 2026-10-11T19:58:40.59164+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e925eb9e29de6393ceec2a30056989de4c6f770a`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
appendonly
no
save

aof-load-truncated
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791719919
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
/data/temp-1791719919.1.rdb

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `361fdf9ad3ee6577fd53ca9afbf430d205fe6c30`. Observed after restart: 2026-10-11T19:58:42.440883+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `361fdf9ad3ee6577fd53ca9afbf430d205fe6c30`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
save

no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendfsync
everysec
aof-load-truncated
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
maxmemory-policy
noeviction
aof-use-rdb-preamble
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
rdb_last_save_time:1791719921
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
/data/temp-1791719921.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `1cfbe152fc4e94c0102c9e95c8a3e262e71872bc`. Observed after restart: 2026-10-11T19:58:44.201263+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1cfbe152fc4e94c0102c9e95c8a3e262e71872bc`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
repl-diskless-load
swapdb
save

appendonly
no
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
appendfsync
everysec
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
maxmemory
134217728
rdb-del-sync-files
yes
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
rdb_last_save_time:1791719922
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
/data/temp-1791719922.1.rdb
/data/nodes.conf

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `4f9cdefa31d52b1ec07962fe48e90304775f32b4`. Process run ID: `bcc745b5f7cb4ca30ece39d4aca1e451e2b989be`. Observed: 2026-10-11T19:58:50.207607+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4f9cdefa31d52b1ec07962fe48e90304775f32b4`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
aof-use-rdb-preamble
no
maxmemory-policy
noeviction
rdb-del-sync-files
yes
aof-load-truncated
yes
save

repl-diskless-load
swapdb
appendfsync
everysec
maxmemory
134217728
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791719914
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

#### Restored node-2

Original Redis Cluster ID: `7f385e1d484746697faa90765e66e2563b95f1b7`. Process run ID: `bf46993f5f1a663fdae9f9014bc79ce3ddea9e71`. Observed: 2026-10-11T19:58:51.329405+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7f385e1d484746697faa90765e66e2563b95f1b7`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
rdb-del-sync-files
yes
aof-load-truncated
yes
appendfsync
everysec
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
save

no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
maxmemory
134217728
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791719916
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:483328
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

Original Redis Cluster ID: `3017030ae4cc288de4885b1c31d4086bf0a2037f`. Process run ID: `0c22da6c02485f5394d24016c6810d43a56861c8`. Observed: 2026-10-11T19:58:52.558769+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3017030ae4cc288de4885b1c31d4086bf0a2037f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
maxmemory
134217728
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
repl-diskless-load
swapdb
save

aof-load-truncated
yes
appendonly
no
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791719917
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

#### Restored node-4

Original Redis Cluster ID: `e925eb9e29de6393ceec2a30056989de4c6f770a`. Process run ID: `e040b61e1d02c6adc35c765cd323431d5aa6a656`. Observed: 2026-10-11T19:58:53.685829+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e925eb9e29de6393ceec2a30056989de4c6f770a`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
appendonly
no
save

aof-load-truncated
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791719919
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

Original Redis Cluster ID: `361fdf9ad3ee6577fd53ca9afbf430d205fe6c30`. Process run ID: `a02228b30ef9503fdae1758ea7b437658e981031`. Observed: 2026-10-11T19:58:54.824899+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `361fdf9ad3ee6577fd53ca9afbf430d205fe6c30`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
save

no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendfsync
everysec
aof-load-truncated
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
maxmemory-policy
noeviction
aof-use-rdb-preamble
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
rdb_last_save_time:1791719921
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

Original Redis Cluster ID: `1cfbe152fc4e94c0102c9e95c8a3e262e71872bc`. Process run ID: `2f7b915e9fc16c6d42653c0eff937fb6be3d957a`. Observed: 2026-10-11T19:58:55.89691+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `1cfbe152fc4e94c0102c9e95c8a3e262e71872bc`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
repl-diskless-load
swapdb
save

appendonly
no
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
appendfsync
everysec
maxmemory-policy
noeviction
auto-aof-rewrite-percentage
0
maxmemory
134217728
rdb-del-sync-files
yes
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
rdb_last_save_time:1791719922
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

Fault action completed 2372.646 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 364 | 100.00% | 0 | 169.57 | 0.472 / 0.632 / 1.314 |

Acknowledged writes: 191. Ambiguous SET outcomes: 0.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 191 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 191 | 0.472 / 0.664 / 1.314 |
| baseline | set | 191 | 0.472 / 0.632 / 1.314 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 191 | 88.979 |

Global pacing: offered 364; started 364; completed 364; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 191; checked: 191; missing: 191; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `bf422c1f305caa93e32e7535a6a108ca2b27f6c7c2b729f2418dc4bf88f1a5cb`.
