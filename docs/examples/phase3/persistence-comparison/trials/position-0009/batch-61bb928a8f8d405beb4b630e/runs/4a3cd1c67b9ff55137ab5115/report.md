# Redis Reliability Lab — persistence-memory

Run: `4a3cd1c67b9ff55137ab5115`  
Outcome: **partial**  
Started: 2026-10-11T16:57:10.97486+08:00  
Finished: 2026-10-11T16:58:02.420328+08:00

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

Frozen acknowledged SETs: 417. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:57:32.770167+08:00 |
| Immutable ledger frozen | 2026-10-11T16:57:32.770196+08:00 |
| All six stopped verified before restart | 2026-10-11T16:57:34.968138+08:00 |
| First restart requested | 2026-10-11T16:57:35.969273+08:00 |
| Restart sequence completed | 2026-10-11T16:57:45.443637+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:57:56.452578+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:57:56.452734+08:00 |

Drain to first kill request: 0.045 ms (runtime monotonic clock).

Drain to last kill completion: 1510.013 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.127 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16394.367 |
| Original topology sustained | 17654.187 |
| Replication redundancy sustained | 17654.187 |
| Post-validation diagnostic service sustained | 24887.334 |
| AOF frozen-value recovery observed | unobserved |

Configured memory-only control disables RDB/AOF value persistence. Persistence recovery is not applicable; restored processes and topology can coexist with lost frozen values. The independent persistence dimension is false when the disabled runtime configuration was observed, while the persistence-recovery event remains absent.

Frozen-ledger validation: expected 417; checked 417; missing 417; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:57:32.770212+08:00 | 2026-10-11T16:57:33.022978+08:00 |  |
| 2 | node-5 | 2026-10-11T16:57:33.022983+08:00 | 2026-10-11T16:57:33.286023+08:00 |  |
| 3 | node-6 | 2026-10-11T16:57:33.286033+08:00 | 2026-10-11T16:57:33.547267+08:00 |  |
| 4 | node-1 | 2026-10-11T16:57:33.547274+08:00 | 2026-10-11T16:57:33.790721+08:00 |  |
| 5 | node-2 | 2026-10-11T16:57:33.790725+08:00 | 2026-10-11T16:57:34.03563+08:00 |  |
| 6 | node-3 | 2026-10-11T16:57:34.035637+08:00 | 2026-10-11T16:57:34.280193+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:57:35.969273+08:00 | 2026-10-11T16:57:37.558501+08:00 |  |
| 2 | node-2 | 2026-10-11T16:57:37.558502+08:00 | 2026-10-11T16:57:39.121742+08:00 |  |
| 3 | node-3 | 2026-10-11T16:57:39.121748+08:00 | 2026-10-11T16:57:40.794691+08:00 |  |
| 4 | node-4 | 2026-10-11T16:57:40.794692+08:00 | 2026-10-11T16:57:42.357446+08:00 |  |
| 5 | node-5 | 2026-10-11T16:57:42.357448+08:00 | 2026-10-11T16:57:43.898348+08:00 |  |
| 6 | node-6 | 2026-10-11T16:57:43.898351+08:00 | 2026-10-11T16:57:45.443636+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 9bec4937572a4bd3682e276791fa9ea0235d484f | 2026-10-11T16:57:34.401882+08:00 | false | no | managed Redis process is stopped |
| node-5 | 113c72026f58143e233a5cfb0689a7755f15037e | 2026-10-11T16:57:34.513114+08:00 | false | no | managed Redis process is stopped |
| node-6 | 7203320f44c58d0c04919751149f7b89fb4c3e5f | 2026-10-11T16:57:34.628601+08:00 | false | no | managed Redis process is stopped |
| node-1 | 98fc0662f08091c356307aa3d8e8b2ec3d43eff7 | 2026-10-11T16:57:34.740859+08:00 | false | no | managed Redis process is stopped |
| node-2 | fb811d18a22f7d3866a180924e4e901c9926f431 | 2026-10-11T16:57:34.855269+08:00 | false | no | managed Redis process is stopped |
| node-3 | 33d3749f9f31c4d97b5bd47f3b30ab6651dd525b | 2026-10-11T16:57:34.968119+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `98fc0662f08091c356307aa3d8e8b2ec3d43eff7`. Process run ID: `50db7d60fe3e96dff972f73665ef78b137b5a684`. Observed: 2026-10-11T16:57:23.832613+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `98fc0662f08091c356307aa3d8e8b2ec3d43eff7`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
appendfsync
everysec
rdb-del-sync-files
yes
aof-use-rdb-preamble
no
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
maxmemory-policy
noeviction
appendonly
no
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
maxmemory
134217728
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
rdb_last_save_time:1791709033
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:516096
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

Original Redis Cluster ID: `fb811d18a22f7d3866a180924e4e901c9926f431`. Process run ID: `d956001bf2017a758e26d2e68bf1cd9a75b8bda1`. Observed: 2026-10-11T16:57:24.835393+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `fb811d18a22f7d3866a180924e4e901c9926f431`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
save

repl-diskless-sync
yes
appendonly
no
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
appendfsync
everysec
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
aof-load-truncated
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709034
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:540672
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

Original Redis Cluster ID: `33d3749f9f31c4d97b5bd47f3b30ab6651dd525b`. Process run ID: `4f02c4986ddd1f1960728705f90815fc62c1aaff`. Observed: 2026-10-11T16:57:25.816465+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `33d3749f9f31c4d97b5bd47f3b30ab6651dd525b`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
maxmemory-policy
noeviction
save

maxmemory
134217728
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
aof-use-rdb-preamble
no
appendonly
no
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
rdb_last_save_time:1791709034
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:536576
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

Original Redis Cluster ID: `9bec4937572a4bd3682e276791fa9ea0235d484f`. Process run ID: `8372c0e2b6faf10429aedcb58fa4d793637cbed3`. Observed: 2026-10-11T16:57:26.789749+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9bec4937572a4bd3682e276791fa9ea0235d484f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
save

maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
appendonly
no
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
no
maxmemory
134217728
aof-load-truncated
yes
repl-diskless-sync
yes
appendfsync
everysec
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
rdb_last_save_time:1791709034
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

Original Redis Cluster ID: `113c72026f58143e233a5cfb0689a7755f15037e`. Process run ID: `6a4167311c38e20d942246a26ed62bb46205fec4`. Observed: 2026-10-11T16:57:27.770007+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `113c72026f58143e233a5cfb0689a7755f15037e`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
save

aof-use-rdb-preamble
no
appendfsync
everysec
maxmemory-policy
noeviction
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
repl-diskless-load
swapdb
appendonly
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
rdb_last_save_time:1791709035
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

Original Redis Cluster ID: `7203320f44c58d0c04919751149f7b89fb4c3e5f`. Process run ID: `93b54ea56547f148c2f1706bd7604d8bac657837`. Observed: 2026-10-11T16:57:28.746938+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7203320f44c58d0c04919751149f7b89fb4c3e5f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
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
appendonly
no
save

aof-use-rdb-preamble
no
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709035
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

Original Redis Cluster ID: `98fc0662f08091c356307aa3d8e8b2ec3d43eff7`. Observed after restart: 2026-10-11T16:57:37.441267+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `98fc0662f08091c356307aa3d8e8b2ec3d43eff7`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
rdb-del-sync-files
yes
appendonly
no
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
save

maxmemory
134217728
maxmemory-policy
noeviction
repl-diskless-load
swapdb
appendfsync
everysec
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
rdb_last_save_time:1791709056
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

Original Redis Cluster ID: `fb811d18a22f7d3866a180924e4e901c9926f431`. Observed after restart: 2026-10-11T16:57:39.008039+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `fb811d18a22f7d3866a180924e4e901c9926f431`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
aof-load-truncated
yes
save

appendonly
no
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
aof-use-rdb-preamble
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
rdb_last_save_time:1791709057
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

Original Redis Cluster ID: `33d3749f9f31c4d97b5bd47f3b30ab6651dd525b`. Observed after restart: 2026-10-11T16:57:40.646933+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `33d3749f9f31c4d97b5bd47f3b30ab6651dd525b`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
no
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
appendfsync
everysec
save

maxmemory-policy
noeviction
aof-load-truncated
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
appendonly
no
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
rdb_last_save_time:1791709059
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

Original Redis Cluster ID: `9bec4937572a4bd3682e276791fa9ea0235d484f`. Observed after restart: 2026-10-11T16:57:42.246061+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9bec4937572a4bd3682e276791fa9ea0235d484f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
save

maxmemory
134217728
aof-use-rdb-preamble
no
no-appendfsync-on-rewrite
no
appendfsync
everysec
rdb-del-sync-files
yes
maxmemory-policy
noeviction
aof-load-truncated
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
rdb_last_save_time:1791709061
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
/data/temp-1791709061.1.rdb

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `113c72026f58143e233a5cfb0689a7755f15037e`. Observed after restart: 2026-10-11T16:57:43.785728+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `113c72026f58143e233a5cfb0689a7755f15037e`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
rdb-del-sync-files
yes
aof-load-truncated
yes
aof-use-rdb-preamble
no
no-appendfsync-on-rewrite
no
save

maxmemory
134217728
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
repl-diskless-sync
yes
repl-diskless-load
swapdb
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
rdb_last_save_time:1791709062
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
/data/temp-1791709062.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `7203320f44c58d0c04919751149f7b89fb4c3e5f`. Observed after restart: 2026-10-11T16:57:45.327656+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7203320f44c58d0c04919751149f7b89fb4c3e5f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
aof-load-truncated
yes
save

appendfsync
everysec
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
no
rdb-del-sync-files
yes
repl-diskless-sync
yes
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709064
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
/data/temp-1791709064.1.rdb

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `98fc0662f08091c356307aa3d8e8b2ec3d43eff7`. Process run ID: `547b4a007072b6484a75956d6e1f385ffd05315a`. Observed: 2026-10-11T16:57:51.466804+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `98fc0662f08091c356307aa3d8e8b2ec3d43eff7`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
rdb-del-sync-files
yes
appendonly
no
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
save

maxmemory
134217728
maxmemory-policy
noeviction
repl-diskless-load
swapdb
appendfsync
everysec
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
rdb_last_save_time:1791709056
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

#### Restored node-2

Original Redis Cluster ID: `fb811d18a22f7d3866a180924e4e901c9926f431`. Process run ID: `dbdb54d39717f36f4beaef91a23f29e60a81c1bb`. Observed: 2026-10-11T16:57:52.43186+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `fb811d18a22f7d3866a180924e4e901c9926f431`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
aof-load-truncated
yes
save

appendonly
no
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
aof-use-rdb-preamble
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
rdb_last_save_time:1791709057
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:471040
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

Original Redis Cluster ID: `33d3749f9f31c4d97b5bd47f3b30ab6651dd525b`. Process run ID: `4de767de4278b1849ad86f04316909adf265decf`. Observed: 2026-10-11T16:57:53.409451+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `33d3749f9f31c4d97b5bd47f3b30ab6651dd525b`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
no
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
appendfsync
everysec
save

maxmemory-policy
noeviction
aof-load-truncated
yes
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
appendonly
no
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
rdb_last_save_time:1791709059
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:573440
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

Original Redis Cluster ID: `9bec4937572a4bd3682e276791fa9ea0235d484f`. Process run ID: `ba1f7819132b09aec2dce4b87bc1b979bac27f9a`. Observed: 2026-10-11T16:57:54.39014+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `9bec4937572a4bd3682e276791fa9ea0235d484f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
save

maxmemory
134217728
aof-use-rdb-preamble
no
no-appendfsync-on-rewrite
no
appendfsync
everysec
rdb-del-sync-files
yes
maxmemory-policy
noeviction
aof-load-truncated
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
rdb_last_save_time:1791709061
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

Original Redis Cluster ID: `113c72026f58143e233a5cfb0689a7755f15037e`. Process run ID: `d5ed4b3f4555c7603b58bd740eaa8edbbe3fd769`. Observed: 2026-10-11T16:57:55.367174+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `113c72026f58143e233a5cfb0689a7755f15037e`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
rdb-del-sync-files
yes
aof-load-truncated
yes
aof-use-rdb-preamble
no
no-appendfsync-on-rewrite
no
save

maxmemory
134217728
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
repl-diskless-sync
yes
repl-diskless-load
swapdb
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
rdb_last_save_time:1791709062
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

Original Redis Cluster ID: `7203320f44c58d0c04919751149f7b89fb4c3e5f`. Process run ID: `cfb9590bb84fbd988efcad0a531288fca75f660a`. Observed: 2026-10-11T16:57:56.336287+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `7203320f44c58d0c04919751149f7b89fb4c3e5f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
aof-load-truncated
yes
save

appendfsync
everysec
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
no
rdb-del-sync-files
yes
repl-diskless-sync
yes
maxmemory-policy
noeviction
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791709064
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

Fault action completed 2197.910 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 804 | 100.00% | 0 | 199.85 | 0.472 / 0.664 / 1.598 |

Acknowledged writes: 417. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 417 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 417 | 0.472 / 0.664 / 1.849 |
| baseline | set | 417 | 0.472 / 0.664 / 1.849 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 417 | 103.654 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 804; started 804; completed 804; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 417; checked: 417; missing: 417; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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
