# Redis Reliability Lab — dataset-trial

Run: `9d637ef7c7c4dce5e3adf2fd`  
Outcome: **partial**  
Started: 2026-10-11T19:50:14.924141+08:00  
Finished: 2026-10-11T19:51:09.755916+08:00

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

Frozen acknowledged SETs: 197. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T19:50:36.878855+08:00 |
| Immutable ledger frozen | 2026-10-11T19:50:36.879+08:00 |
| All six stopped verified before restart | 2026-10-11T19:50:39.155635+08:00 |
| First restart requested | 2026-10-11T19:50:43.158099+08:00 |
| Restart sequence completed | 2026-10-11T19:50:53.216858+08:00 |
| Frozen-ledger validation completed | 2026-10-11T19:51:04.497299+08:00 |
| Post-validation diagnostics requested | 2026-10-11T19:51:04.497575+08:00 |

Drain to first kill request: 0.261 ms (runtime monotonic clock).

Drain to last kill completion: 1558.723 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 4002.468 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 19916.513 |
| Original topology sustained | 21120.685 |
| Replication redundancy sustained | 21120.685 |
| Post-validation diagnostic service sustained | 28845.201 |
| AOF frozen-value recovery observed | unobserved |

Configured memory-only control disables RDB/AOF value persistence. Persistence recovery is not applicable; restored processes and topology can coexist with lost frozen values. The independent persistence dimension is false when the disabled runtime configuration was observed, while the persistence-recovery event remains absent.

Frozen-ledger validation: expected 197; checked 197; missing 197; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T19:50:36.879117+08:00 | 2026-10-11T19:50:37.137182+08:00 |  |
| 2 | node-5 | 2026-10-11T19:50:37.137196+08:00 | 2026-10-11T19:50:37.404819+08:00 |  |
| 3 | node-6 | 2026-10-11T19:50:37.404844+08:00 | 2026-10-11T19:50:37.671352+08:00 |  |
| 4 | node-1 | 2026-10-11T19:50:37.671372+08:00 | 2026-10-11T19:50:37.932806+08:00 |  |
| 5 | node-2 | 2026-10-11T19:50:37.932815+08:00 | 2026-10-11T19:50:38.18672+08:00 |  |
| 6 | node-3 | 2026-10-11T19:50:38.186735+08:00 | 2026-10-11T19:50:38.437577+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T19:50:43.158099+08:00 | 2026-10-11T19:50:44.746085+08:00 |  |
| 2 | node-2 | 2026-10-11T19:50:44.746089+08:00 | 2026-10-11T19:50:46.552167+08:00 |  |
| 3 | node-3 | 2026-10-11T19:50:46.552179+08:00 | 2026-10-11T19:50:48.22322+08:00 |  |
| 4 | node-4 | 2026-10-11T19:50:48.223229+08:00 | 2026-10-11T19:50:49.891122+08:00 |  |
| 5 | node-5 | 2026-10-11T19:50:49.891129+08:00 | 2026-10-11T19:50:51.558641+08:00 |  |
| 6 | node-6 | 2026-10-11T19:50:51.558649+08:00 | 2026-10-11T19:50:53.216854+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | cb5a9b1c09fc1456842750ca847adad7ab56de7b | 2026-10-11T19:50:38.564357+08:00 | false | no | managed Redis process is stopped |
| node-5 | e2ccdf9333d5fb076fc9151e111384bf85ead9d8 | 2026-10-11T19:50:38.684837+08:00 | false | no | managed Redis process is stopped |
| node-6 | 5833d47aba403296d1f6e130da8e14496fa7f253 | 2026-10-11T19:50:38.799803+08:00 | false | no | managed Redis process is stopped |
| node-1 | 6f5d43c99995e4332b23a24fa5b75ed7d8117f69 | 2026-10-11T19:50:38.919126+08:00 | false | no | managed Redis process is stopped |
| node-2 | 02639372d5007df31125d2928a4ef4bf57dd9083 | 2026-10-11T19:50:39.038082+08:00 | false | no | managed Redis process is stopped |
| node-3 | 0f87624dd96e210c6ab980ce59aa93d880d22a56 | 2026-10-11T19:50:39.155595+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `6f5d43c99995e4332b23a24fa5b75ed7d8117f69`. Process run ID: `0e8e3e27e895b92708f7283bd8e39029d28a368c`. Observed: 2026-10-11T19:50:29.538705+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6f5d43c99995e4332b23a24fa5b75ed7d8117f69`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
appendonly
no
aof-load-truncated
yes
maxmemory
134217728
save

appendfsync
everysec
repl-diskless-sync
yes
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
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
rdb_last_save_time:1791719419
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

### Prepared node-2

Original Redis Cluster ID: `02639372d5007df31125d2928a4ef4bf57dd9083`. Process run ID: `182fbc2e41a47826385daa3cbe4499610dd8e58a`. Observed: 2026-10-11T19:50:30.577618+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `02639372d5007df31125d2928a4ef4bf57dd9083`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
rdb-del-sync-files
yes
appendfsync
everysec
appendonly
no
save

maxmemory-policy
noeviction
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
rdb_last_save_time:1791719419
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:561152
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

Original Redis Cluster ID: `0f87624dd96e210c6ab980ce59aa93d880d22a56`. Process run ID: `d32b56fbcd3f7079680789af0c182354207ddba8`. Observed: 2026-10-11T19:50:31.593546+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `0f87624dd96e210c6ab980ce59aa93d880d22a56`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendfsync
everysec
maxmemory
134217728
maxmemory-policy
noeviction
repl-diskless-sync
yes
aof-use-rdb-preamble
no
aof-load-truncated
yes
repl-diskless-load
swapdb
save

auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791719419
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

### Prepared node-4

Original Redis Cluster ID: `cb5a9b1c09fc1456842750ca847adad7ab56de7b`. Process run ID: `9cd4414722bd51cc3ce98cba43a6cc2eab96e188`. Observed: 2026-10-11T19:50:32.650462+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `cb5a9b1c09fc1456842750ca847adad7ab56de7b`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
no-appendfsync-on-rewrite
no
save

maxmemory-policy
noeviction
repl-diskless-sync
yes
maxmemory
134217728
repl-diskless-load
swapdb
rdb-del-sync-files
yes
aof-use-rdb-preamble
no
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791719420
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

Original Redis Cluster ID: `e2ccdf9333d5fb076fc9151e111384bf85ead9d8`. Process run ID: `d71bfdde6cb9542e260117b7403057538df1d306`. Observed: 2026-10-11T19:50:33.696084+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e2ccdf9333d5fb076fc9151e111384bf85ead9d8`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
aof-use-rdb-preamble
no
maxmemory
134217728
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
aof-load-truncated
yes
repl-diskless-sync
yes
appendfsync
everysec
appendonly
no
save

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
rdb_last_save_time:1791719420
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

Original Redis Cluster ID: `5833d47aba403296d1f6e130da8e14496fa7f253`. Process run ID: `95117168aa9f008673353307d69aab47156a8ac1`. Observed: 2026-10-11T19:50:34.731278+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `5833d47aba403296d1f6e130da8e14496fa7f253`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
repl-diskless-sync
yes
save

rdb-del-sync-files
yes
maxmemory
134217728
repl-diskless-load
swapdb
aof-use-rdb-preamble
no
appendonly
no
appendfsync
everysec
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
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
rdb_last_save_time:1791719420
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

Original Redis Cluster ID: `6f5d43c99995e4332b23a24fa5b75ed7d8117f69`. Observed after restart: 2026-10-11T19:50:44.627393+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6f5d43c99995e4332b23a24fa5b75ed7d8117f69`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
appendfsync
everysec
appendonly
no
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
repl-diskless-load
swapdb
maxmemory
134217728
maxmemory-policy
noeviction
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
rdb_last_save_time:1791719443
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

Original Redis Cluster ID: `02639372d5007df31125d2928a4ef4bf57dd9083`. Observed after restart: 2026-10-11T19:50:46.421504+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `02639372d5007df31125d2928a4ef4bf57dd9083`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
appendonly
no
maxmemory
134217728
aof-use-rdb-preamble
no
repl-diskless-sync
yes
repl-diskless-load
swapdb
save

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
rdb_last_save_time:1791719445
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

Original Redis Cluster ID: `0f87624dd96e210c6ab980ce59aa93d880d22a56`. Observed after restart: 2026-10-11T19:50:48.102334+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `0f87624dd96e210c6ab980ce59aa93d880d22a56`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
repl-diskless-sync
yes
maxmemory
134217728
aof-load-truncated
yes
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
no
maxmemory-policy
noeviction
save

appendonly
no
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791719446
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

Original Redis Cluster ID: `cb5a9b1c09fc1456842750ca847adad7ab56de7b`. Observed after restart: 2026-10-11T19:50:49.762203+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `cb5a9b1c09fc1456842750ca847adad7ab56de7b`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
save

appendfsync
everysec
appendonly
no
rdb-del-sync-files
yes
repl-diskless-load
swapdb
maxmemory
134217728
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-sync
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
rdb_last_save_time:1791719448
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
/data/temp-1791719448.1.rdb
/data/nodes.conf

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `e2ccdf9333d5fb076fc9151e111384bf85ead9d8`. Observed after restart: 2026-10-11T19:50:51.433851+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e2ccdf9333d5fb076fc9151e111384bf85ead9d8`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
no
repl-diskless-load
swapdb
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendonly
no
maxmemory-policy
noeviction
appendfsync
everysec
aof-load-truncated
yes
save

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
rdb_last_save_time:1791719450
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
/data/temp-1791719450.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `5833d47aba403296d1f6e130da8e14496fa7f253`. Observed after restart: 2026-10-11T19:50:53.086578+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `5833d47aba403296d1f6e130da8e14496fa7f253`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
aof-use-rdb-preamble
no
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
aof-load-truncated
yes
save

repl-diskless-sync
yes
maxmemory
134217728
appendonly
no
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791719451
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
/data/temp-1791719451.1.rdb

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `6f5d43c99995e4332b23a24fa5b75ed7d8117f69`. Process run ID: `7c0f7e2c3c3c4ccdb8458e198122e04c2f701504`. Observed: 2026-10-11T19:50:59.265722+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `6f5d43c99995e4332b23a24fa5b75ed7d8117f69`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
appendfsync
everysec
appendonly
no
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
repl-diskless-load
swapdb
maxmemory
134217728
maxmemory-policy
noeviction
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
rdb_last_save_time:1791719443
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

#### Restored node-2

Original Redis Cluster ID: `02639372d5007df31125d2928a4ef4bf57dd9083`. Process run ID: `335c9bdfd578d2e1cc79de197913f0d40215c577`. Observed: 2026-10-11T19:51:00.265641+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `02639372d5007df31125d2928a4ef4bf57dd9083`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
appendonly
no
maxmemory
134217728
aof-use-rdb-preamble
no
repl-diskless-sync
yes
repl-diskless-load
swapdb
save

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
rdb_last_save_time:1791719445
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

Original Redis Cluster ID: `0f87624dd96e210c6ab980ce59aa93d880d22a56`. Process run ID: `e70b8c1acc7c4b2d2c98a5da4d1faa3da1719e3e`. Observed: 2026-10-11T19:51:01.277327+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `0f87624dd96e210c6ab980ce59aa93d880d22a56`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
repl-diskless-sync
yes
maxmemory
134217728
aof-load-truncated
yes
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
no
maxmemory-policy
noeviction
save

appendonly
no
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791719446
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:479232
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

Original Redis Cluster ID: `cb5a9b1c09fc1456842750ca847adad7ab56de7b`. Process run ID: `01cfddae5438e73cd34903ebf95cd66f887c7b09`. Observed: 2026-10-11T19:51:02.300783+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `cb5a9b1c09fc1456842750ca847adad7ab56de7b`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
save

appendfsync
everysec
appendonly
no
rdb-del-sync-files
yes
repl-diskless-load
swapdb
maxmemory
134217728
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
repl-diskless-sync
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
rdb_last_save_time:1791719448
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

Original Redis Cluster ID: `e2ccdf9333d5fb076fc9151e111384bf85ead9d8`. Process run ID: `df5e0706a1eaad3bd773b32cd70dc76b5ebde1ce`. Observed: 2026-10-11T19:51:03.352319+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e2ccdf9333d5fb076fc9151e111384bf85ead9d8`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
no
repl-diskless-load
swapdb
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
rdb-del-sync-files
yes
appendonly
no
maxmemory-policy
noeviction
appendfsync
everysec
aof-load-truncated
yes
save

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
rdb_last_save_time:1791719450
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

Original Redis Cluster ID: `5833d47aba403296d1f6e130da8e14496fa7f253`. Process run ID: `7a69a40a8b1ebffd20760ab1222dec874ddcabcb`. Observed: 2026-10-11T19:51:04.425325+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `5833d47aba403296d1f6e130da8e14496fa7f253`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
aof-use-rdb-preamble
no
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
aof-load-truncated
yes
save

repl-diskless-sync
yes
maxmemory
134217728
appendonly
no
auto-aof-rewrite-percentage
0
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
rdb_last_save_time:1791719451
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

Fault action completed 2276.590 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 364 | 100.00% | 0 | 169.54 | 0.602 / 0.807 / 2.039 |

Acknowledged writes: 197. Ambiguous SET outcomes: 0.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 197 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 197 | 0.602 / 0.847 / 1.252 |
| baseline | set | 197 | 0.573 / 0.807 / 1.192 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 197 | 91.755 |

Global pacing: offered 364; started 364; completed 364; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 197; checked: 197; missing: 197; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `85ec7fe735261ce30ad6447fe4f26cae8a8e98160a0a45921f0ccb1eb4c7f5c0`.
