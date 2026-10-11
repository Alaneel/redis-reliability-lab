# Redis Reliability Lab — persistence-memory

Run: `fa1b57f987b8c265178b1829`  
Outcome: **partial**  
Started: 2026-10-11T16:52:53.206239+08:00  
Finished: 2026-10-11T16:53:45.059144+08:00

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

Frozen acknowledged SETs: 418. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:53:15.17304+08:00 |
| Immutable ledger frozen | 2026-10-11T16:53:15.173092+08:00 |
| All six stopped verified before restart | 2026-10-11T16:53:17.569558+08:00 |
| First restart requested | 2026-10-11T16:53:18.571042+08:00 |
| Restart sequence completed | 2026-10-11T16:53:28.09094+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:53:39.131541+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:53:39.131685+08:00 |

Drain to first kill request: 0.147 ms (runtime monotonic clock).

Drain to last kill completion: 1623.689 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.476 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16633.900 |
| Original topology sustained | 17817.071 |
| Replication redundancy sustained | 17817.071 |
| Post-validation diagnostic service sustained | 25204.776 |
| AOF frozen-value recovery observed | unobserved |

Configured memory-only control disables RDB/AOF value persistence. Persistence recovery is not applicable; restored processes and topology can coexist with lost frozen values. The independent persistence dimension is false when the disabled runtime configuration was observed, while the persistence-recovery event remains absent.

Frozen-ledger validation: expected 418; checked 418; missing 418; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:53:15.173187+08:00 | 2026-10-11T16:53:15.417184+08:00 |  |
| 2 | node-5 | 2026-10-11T16:53:15.417189+08:00 | 2026-10-11T16:53:15.666556+08:00 |  |
| 3 | node-6 | 2026-10-11T16:53:15.666567+08:00 | 2026-10-11T16:53:16.003719+08:00 |  |
| 4 | node-1 | 2026-10-11T16:53:16.003726+08:00 | 2026-10-11T16:53:16.276364+08:00 |  |
| 5 | node-2 | 2026-10-11T16:53:16.276368+08:00 | 2026-10-11T16:53:16.536679+08:00 |  |
| 6 | node-3 | 2026-10-11T16:53:16.53671+08:00 | 2026-10-11T16:53:16.796742+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:53:18.571042+08:00 | 2026-10-11T16:53:20.181995+08:00 |  |
| 2 | node-2 | 2026-10-11T16:53:20.181997+08:00 | 2026-10-11T16:53:21.800394+08:00 |  |
| 3 | node-3 | 2026-10-11T16:53:21.800402+08:00 | 2026-10-11T16:53:23.368338+08:00 |  |
| 4 | node-4 | 2026-10-11T16:53:23.368341+08:00 | 2026-10-11T16:53:24.976812+08:00 |  |
| 5 | node-5 | 2026-10-11T16:53:24.976813+08:00 | 2026-10-11T16:53:26.52815+08:00 |  |
| 6 | node-6 | 2026-10-11T16:53:26.528152+08:00 | 2026-10-11T16:53:28.090938+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 82074e3798227e809cf614145a633504ff65d5b4 | 2026-10-11T16:53:16.933117+08:00 | false | no | managed Redis process is stopped |
| node-5 | 08422d688e2f5cdcb8cea639fd2b48628bde9d2f | 2026-10-11T16:53:17.093236+08:00 | false | no | managed Redis process is stopped |
| node-6 | 0fdcc6893bab895b36c58509cb0db60cd2827e6f | 2026-10-11T16:53:17.21725+08:00 | false | no | managed Redis process is stopped |
| node-1 | 78d73678f8ebe8b7b4c0768a08bf51c2bf2fa254 | 2026-10-11T16:53:17.336715+08:00 | false | no | managed Redis process is stopped |
| node-2 | 571e49629aa19f09758a4a13fbfb01c91f999c98 | 2026-10-11T16:53:17.452571+08:00 | false | no | managed Redis process is stopped |
| node-3 | 4abc5361436730516621bc5ca240247f5c8e21df | 2026-10-11T16:53:17.569549+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `78d73678f8ebe8b7b4c0768a08bf51c2bf2fa254`. Process run ID: `ea6427a41341007f7ca5d628ab1bb2c6649370fb`. Observed: 2026-10-11T16:53:06.16047+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `78d73678f8ebe8b7b4c0768a08bf51c2bf2fa254`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
repl-diskless-load
swapdb
save

appendfsync
everysec
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
aof-load-truncated
yes
appendonly
no
rdb-del-sync-files
yes
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791708776
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:475136
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

Original Redis Cluster ID: `571e49629aa19f09758a4a13fbfb01c91f999c98`. Process run ID: `6910c75b09285a7a077daa7ba2d267eb1345f4f9`. Observed: 2026-10-11T16:53:07.182025+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `571e49629aa19f09758a4a13fbfb01c91f999c98`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
aof-load-truncated
yes
repl-diskless-sync
yes
repl-diskless-load
swapdb
save

maxmemory
134217728
appendfsync
everysec
rdb-del-sync-files
yes
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
rdb_last_save_time:1791708776
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:565248
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

Original Redis Cluster ID: `4abc5361436730516621bc5ca240247f5c8e21df`. Process run ID: `9d268c6001ea489f9ee5ddd8c012c1125b1262e9`. Observed: 2026-10-11T16:53:08.146458+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4abc5361436730516621bc5ca240247f5c8e21df`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
maxmemory
134217728
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
repl-diskless-load
swapdb
rdb-del-sync-files
yes
appendonly
no
appendfsync
everysec
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791708776
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:491520
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

Original Redis Cluster ID: `82074e3798227e809cf614145a633504ff65d5b4`. Process run ID: `379b87c3d1877325001e9e38de244adab0b07c94`. Observed: 2026-10-11T16:53:09.134741+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `82074e3798227e809cf614145a633504ff65d5b4`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
rdb-del-sync-files
yes
appendonly
no
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
save

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
rdb_last_save_time:1791708776
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

Original Redis Cluster ID: `08422d688e2f5cdcb8cea639fd2b48628bde9d2f`. Process run ID: `a7c32f8353176504abb57a472c0318b6b1364277`. Observed: 2026-10-11T16:53:10.135148+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `08422d688e2f5cdcb8cea639fd2b48628bde9d2f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
aof-load-truncated
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
maxmemory
134217728
repl-diskless-load
swapdb
appendonly
no
rdb-del-sync-files
yes
save

maxmemory-policy
noeviction
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
rdb_last_save_time:1791708777
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

Original Redis Cluster ID: `0fdcc6893bab895b36c58509cb0db60cd2827e6f`. Process run ID: `30b04a549fe4c8a9566e1406c159c7038074b372`. Observed: 2026-10-11T16:53:11.142042+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `0fdcc6893bab895b36c58509cb0db60cd2827e6f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
appendonly
no
save

maxmemory-policy
noeviction
repl-diskless-sync
yes
aof-use-rdb-preamble
no
repl-diskless-load
swapdb
no-appendfsync-on-rewrite
no
maxmemory
134217728
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
rdb_last_save_time:1791708777
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

Original Redis Cluster ID: `78d73678f8ebe8b7b4c0768a08bf51c2bf2fa254`. Observed after restart: 2026-10-11T16:53:20.061987+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `78d73678f8ebe8b7b4c0768a08bf51c2bf2fa254`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
save

maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
maxmemory
134217728
aof-use-rdb-preamble
no
aof-load-truncated
yes
appendonly
no
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
rdb_last_save_time:1791708798
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

Original Redis Cluster ID: `571e49629aa19f09758a4a13fbfb01c91f999c98`. Observed after restart: 2026-10-11T16:53:21.680077+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `571e49629aa19f09758a4a13fbfb01c91f999c98`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
no
save

aof-load-truncated
yes
appendonly
no
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
maxmemory
134217728
maxmemory-policy
noeviction
rdb-del-sync-files
yes
repl-diskless-sync
yes
appendfsync
everysec
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
rdb_last_save_time:1791708800
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

Original Redis Cluster ID: `4abc5361436730516621bc5ca240247f5c8e21df`. Observed after restart: 2026-10-11T16:53:23.254704+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4abc5361436730516621bc5ca240247f5c8e21df`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
save

maxmemory
134217728
repl-diskless-sync
yes
appendfsync
everysec
repl-diskless-load
swapdb
appendonly
no
maxmemory-policy
noeviction
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
no
aof-load-truncated
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
rdb_last_save_time:1791708802
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

Original Redis Cluster ID: `82074e3798227e809cf614145a633504ff65d5b4`. Observed after restart: 2026-10-11T16:53:24.855765+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `82074e3798227e809cf614145a633504ff65d5b4`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
repl-diskless-load
swapdb
appendonly
no
repl-diskless-sync
yes
maxmemory-policy
noeviction
rdb-del-sync-files
yes
appendfsync
everysec
maxmemory
134217728
aof-use-rdb-preamble
no
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708803
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
/data/temp-1791708803.1.rdb

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `08422d688e2f5cdcb8cea639fd2b48628bde9d2f`. Observed after restart: 2026-10-11T16:53:26.415416+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `08422d688e2f5cdcb8cea639fd2b48628bde9d2f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
aof-load-truncated
yes
aof-use-rdb-preamble
no
save

appendfsync
everysec
appendonly
no
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708805
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
/data/temp-1791708805.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `0fdcc6893bab895b36c58509cb0db60cd2827e6f`. Observed after restart: 2026-10-11T16:53:27.974201+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `0fdcc6893bab895b36c58509cb0db60cd2827e6f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
no
rdb-del-sync-files
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
save

maxmemory
134217728
repl-diskless-sync
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
appendonly
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
rdb_last_save_time:1791708806
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
/data/temp-1791708806.1.rdb

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `78d73678f8ebe8b7b4c0768a08bf51c2bf2fa254`. Process run ID: `4b074fa091eee322aae549c40218ccb8844703f6`. Observed: 2026-10-11T16:53:34.033134+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `78d73678f8ebe8b7b4c0768a08bf51c2bf2fa254`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
save

maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
maxmemory
134217728
aof-use-rdb-preamble
no
aof-load-truncated
yes
appendonly
no
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
rdb_last_save_time:1791708798
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:462848
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

Original Redis Cluster ID: `571e49629aa19f09758a4a13fbfb01c91f999c98`. Process run ID: `b5fef630bb4b31ccc778935516456b32589cbbc6`. Observed: 2026-10-11T16:53:35.024531+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `571e49629aa19f09758a4a13fbfb01c91f999c98`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
no
save

aof-load-truncated
yes
appendonly
no
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
maxmemory
134217728
maxmemory-policy
noeviction
rdb-del-sync-files
yes
repl-diskless-sync
yes
appendfsync
everysec
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
rdb_last_save_time:1791708800
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:466944
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

Original Redis Cluster ID: `4abc5361436730516621bc5ca240247f5c8e21df`. Process run ID: `debd1bb4e2a1adcae0439ed90ccdccc72a0ed2ff`. Observed: 2026-10-11T16:53:36.02251+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `4abc5361436730516621bc5ca240247f5c8e21df`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
save

maxmemory
134217728
repl-diskless-sync
yes
appendfsync
everysec
repl-diskless-load
swapdb
appendonly
no
maxmemory-policy
noeviction
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
no
aof-load-truncated
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
rdb_last_save_time:1791708802
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

Original Redis Cluster ID: `82074e3798227e809cf614145a633504ff65d5b4`. Process run ID: `b863c1cbd62ce042dd83698d69b83418b5f8fb76`. Observed: 2026-10-11T16:53:37.060994+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `82074e3798227e809cf614145a633504ff65d5b4`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
repl-diskless-load
swapdb
appendonly
no
repl-diskless-sync
yes
maxmemory-policy
noeviction
rdb-del-sync-files
yes
appendfsync
everysec
maxmemory
134217728
aof-use-rdb-preamble
no
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708803
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

Original Redis Cluster ID: `08422d688e2f5cdcb8cea639fd2b48628bde9d2f`. Process run ID: `e18e2a48769eaedf708eeb25ee75093b21af2cf5`. Observed: 2026-10-11T16:53:38.040782+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `08422d688e2f5cdcb8cea639fd2b48628bde9d2f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
aof-load-truncated
yes
aof-use-rdb-preamble
no
save

appendfsync
everysec
appendonly
no
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
maxmemory
134217728
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708805
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

Original Redis Cluster ID: `0fdcc6893bab895b36c58509cb0db60cd2827e6f`. Process run ID: `df89386f80e47282921f8e036bdeafd3a6e7bc73`. Observed: 2026-10-11T16:53:39.015276+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `0fdcc6893bab895b36c58509cb0db60cd2827e6f`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
no
rdb-del-sync-files
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
save

maxmemory
134217728
repl-diskless-sync
yes
appendfsync
everysec
auto-aof-rewrite-percentage
0
appendonly
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
rdb_last_save_time:1791708806
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

Fault action completed 2396.354 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 806 | 100.00% | 0 | 199.96 | 0.546 / 0.732 / 0.934 |

Acknowledged writes: 418. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 418 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 418 | 0.573 / 0.768 / 0.934 |
| baseline | set | 418 | 0.573 / 0.768 / 0.934 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 418 | 103.702 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 806; started 806; completed 806; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 418; checked: 418; missing: 418; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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
