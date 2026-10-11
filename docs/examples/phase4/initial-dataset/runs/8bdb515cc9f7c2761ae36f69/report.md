# Redis Reliability Lab — dataset-trial

Run: `8bdb515cc9f7c2761ae36f69`  
Outcome: **partial**  
Started: 2026-10-11T19:45:28.790536+08:00  
Finished: 2026-10-11T19:46:22.961173+08:00

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

Frozen acknowledged SETs: 121. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T19:45:49.570843+08:00 |
| Immutable ledger frozen | 2026-10-11T19:45:49.570965+08:00 |
| All six stopped verified before restart | 2026-10-11T19:45:51.885588+08:00 |
| First restart requested | 2026-10-11T19:45:54.886935+08:00 |
| Restart sequence completed | 2026-10-11T19:46:04.703888+08:00 |
| Frozen-ledger validation completed | 2026-10-11T19:46:17.179854+08:00 |
| Post-validation diagnostics requested | 2026-10-11T19:46:17.180126+08:00 |

Drain to first kill request: 0.152 ms (runtime monotonic clock).

Drain to last kill completion: 1554.628 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 3001.350 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 19536.447 |
| Original topology sustained | 21001.342 |
| Replication redundancy sustained | 21001.342 |
| Post-validation diagnostic service sustained | 29068.960 |
| AOF frozen-value recovery observed | unobserved |

Configured memory-only control disables RDB/AOF value persistence. Persistence recovery is not applicable; restored processes and topology can coexist with lost frozen values. The independent persistence dimension is false when the disabled runtime configuration was observed, while the persistence-recovery event remains absent.

Frozen-ledger validation: expected 121; checked 121; missing 121; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T19:45:49.570996+08:00 | 2026-10-11T19:45:49.821303+08:00 |  |
| 2 | node-5 | 2026-10-11T19:45:49.82132+08:00 | 2026-10-11T19:45:50.083369+08:00 |  |
| 3 | node-6 | 2026-10-11T19:45:50.083384+08:00 | 2026-10-11T19:45:50.351688+08:00 |  |
| 4 | node-1 | 2026-10-11T19:45:50.351705+08:00 | 2026-10-11T19:45:50.620964+08:00 |  |
| 5 | node-2 | 2026-10-11T19:45:50.620973+08:00 | 2026-10-11T19:45:50.87328+08:00 |  |
| 6 | node-3 | 2026-10-11T19:45:50.873307+08:00 | 2026-10-11T19:45:51.125469+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T19:45:54.886935+08:00 | 2026-10-11T19:45:56.507678+08:00 |  |
| 2 | node-2 | 2026-10-11T19:45:56.507682+08:00 | 2026-10-11T19:45:58.132064+08:00 |  |
| 3 | node-3 | 2026-10-11T19:45:58.132079+08:00 | 2026-10-11T19:45:59.757306+08:00 |  |
| 4 | node-4 | 2026-10-11T19:45:59.757315+08:00 | 2026-10-11T19:46:01.391575+08:00 |  |
| 5 | node-5 | 2026-10-11T19:46:01.391578+08:00 | 2026-10-11T19:46:03.071061+08:00 |  |
| 6 | node-6 | 2026-10-11T19:46:03.071067+08:00 | 2026-10-11T19:46:04.703886+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | fca8c4a47930f008988e801fcb806fdba30e5b14 | 2026-10-11T19:45:51.262884+08:00 | false | no | managed Redis process is stopped |
| node-5 | 966f17b00394c5b5c7ac58428e19497a22b779e7 | 2026-10-11T19:45:51.389011+08:00 | false | no | managed Redis process is stopped |
| node-6 | 8905830ae19a4accc8d3ea0580d11f00aaabf924 | 2026-10-11T19:45:51.513124+08:00 | false | no | managed Redis process is stopped |
| node-1 | e3f5e7cd1ca22884fbbdfffc5280df0787b666f8 | 2026-10-11T19:45:51.638323+08:00 | false | no | managed Redis process is stopped |
| node-2 | 29255d56e132ceef12fea24f29f95ff28d6f8557 | 2026-10-11T19:45:51.761032+08:00 | false | no | managed Redis process is stopped |
| node-3 | b926f55d56cf8180e62c4bb92fb5110226f64356 | 2026-10-11T19:45:51.885518+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `e3f5e7cd1ca22884fbbdfffc5280df0787b666f8`. Process run ID: `874f61b31a3865e9f7437a782b51a0eb1b45a4b9`. Observed: 2026-10-11T19:45:42.197788+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e3f5e7cd1ca22884fbbdfffc5280df0787b666f8`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
aof-load-truncated
yes
save

repl-diskless-load
swapdb
repl-diskless-sync
yes
aof-use-rdb-preamble
no
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
appendfsync
everysec
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
rdb_last_save_time:1791719132
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

Original Redis Cluster ID: `29255d56e132ceef12fea24f29f95ff28d6f8557`. Process run ID: `76656286df04c06034b9f289c9eac4c1a8ae3273`. Observed: 2026-10-11T19:45:43.266999+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `29255d56e132ceef12fea24f29f95ff28d6f8557`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
repl-diskless-load
swapdb
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
save

maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
appendfsync
everysec
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791719132
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

### Prepared node-3

Original Redis Cluster ID: `b926f55d56cf8180e62c4bb92fb5110226f64356`. Process run ID: `634f933171e1cf5fa3aa21e0ac7cf0ca1964f27a`. Observed: 2026-10-11T19:45:44.274086+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b926f55d56cf8180e62c4bb92fb5110226f64356`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
appendonly
no
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
aof-use-rdb-preamble
no
appendfsync
everysec
save

maxmemory-policy
noeviction
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
rdb_last_save_time:1791719132
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:1
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

### Prepared node-4

Original Redis Cluster ID: `fca8c4a47930f008988e801fcb806fdba30e5b14`. Process run ID: `bdfaf1c71c5760a48d307afdd8be96b5d0fb967a`. Observed: 2026-10-11T19:45:45.288859+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `fca8c4a47930f008988e801fcb806fdba30e5b14`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
aof-load-truncated
yes
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
maxmemory
134217728
aof-use-rdb-preamble
no
repl-diskless-sync
yes
save

no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
repl-diskless-load
swapdb
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
rdb_last_save_time:1791719133
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

Original Redis Cluster ID: `966f17b00394c5b5c7ac58428e19497a22b779e7`. Process run ID: `09eb2e706f32af744fc09295be55850d1fa52679`. Observed: 2026-10-11T19:45:46.397143+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `966f17b00394c5b5c7ac58428e19497a22b779e7`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
aof-use-rdb-preamble
no
rdb-del-sync-files
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
save

aof-load-truncated
yes
repl-diskless-sync
yes
appendfsync
everysec
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
rdb_last_save_time:1791719133
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

Original Redis Cluster ID: `8905830ae19a4accc8d3ea0580d11f00aaabf924`. Process run ID: `71dec7c49f35ec5dccb29de64777b1fcf9de7b77`. Observed: 2026-10-11T19:45:47.427847+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `8905830ae19a4accc8d3ea0580d11f00aaabf924`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
save

repl-diskless-sync
yes
appendfsync
everysec
maxmemory
134217728
rdb-del-sync-files
yes
auto-aof-rewrite-percentage
0
appendonly
no
repl-diskless-load
swapdb
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
rdb_last_save_time:1791719133
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

Original Redis Cluster ID: `e3f5e7cd1ca22884fbbdfffc5280df0787b666f8`. Observed after restart: 2026-10-11T19:45:56.386743+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e3f5e7cd1ca22884fbbdfffc5280df0787b666f8`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
appendonly
no
rdb-del-sync-files
yes
repl-diskless-sync
yes
maxmemory-policy
noeviction
save

no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
aof-load-truncated
yes
appendfsync
everysec
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
rdb_last_save_time:1791719155
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

Original Redis Cluster ID: `29255d56e132ceef12fea24f29f95ff28d6f8557`. Observed after restart: 2026-10-11T19:45:58.003072+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `29255d56e132ceef12fea24f29f95ff28d6f8557`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
aof-use-rdb-preamble
no
save

aof-load-truncated
yes
repl-diskless-sync
yes
rdb-del-sync-files
yes
appendfsync
everysec
no-appendfsync-on-rewrite
no
maxmemory
134217728
repl-diskless-load
swapdb
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
rdb_last_save_time:1791719156
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

Original Redis Cluster ID: `b926f55d56cf8180e62c4bb92fb5110226f64356`. Observed after restart: 2026-10-11T19:45:59.638199+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b926f55d56cf8180e62c4bb92fb5110226f64356`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
maxmemory
134217728
appendfsync
everysec
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
appendonly
no
aof-use-rdb-preamble
no
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
save

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
rdb_last_save_time:1791719158
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

Original Redis Cluster ID: `fca8c4a47930f008988e801fcb806fdba30e5b14`. Observed after restart: 2026-10-11T19:46:01.268621+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `fca8c4a47930f008988e801fcb806fdba30e5b14`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
repl-diskless-load
swapdb
appendonly
no
maxmemory-policy
noeviction
save

aof-use-rdb-preamble
no
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791719160
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
/data/temp-1791719160.1.rdb

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `966f17b00394c5b5c7ac58428e19497a22b779e7`. Observed after restart: 2026-10-11T19:46:02.946606+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `966f17b00394c5b5c7ac58428e19497a22b779e7`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
save

aof-use-rdb-preamble
no
appendonly
no
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
maxmemory-policy
noeviction
appendfsync
everysec
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
rdb-del-sync-files
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
rdb_last_save_time:1791719161
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
/data/temp-1791719161.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `8905830ae19a4accc8d3ea0580d11f00aaabf924`. Observed after restart: 2026-10-11T19:46:04.586002+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `8905830ae19a4accc8d3ea0580d11f00aaabf924`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
repl-diskless-load
swapdb
appendfsync
everysec
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
no
maxmemory
134217728
rdb-del-sync-files
yes
maxmemory-policy
noeviction
appendonly
no
aof-load-truncated
yes
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
rdb_last_save_time:1791719163
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
/data/temp-1791719163.1.rdb
/data/nodes.conf

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `e3f5e7cd1ca22884fbbdfffc5280df0787b666f8`. Process run ID: `d9a52c6f295be938500fa6231f0d05a96641aacf`. Observed: 2026-10-11T19:46:12.01257+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `e3f5e7cd1ca22884fbbdfffc5280df0787b666f8`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
appendonly
no
rdb-del-sync-files
yes
repl-diskless-sync
yes
maxmemory-policy
noeviction
save

no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
aof-load-truncated
yes
appendfsync
everysec
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
rdb_last_save_time:1791719155
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

#### Restored node-2

Original Redis Cluster ID: `29255d56e132ceef12fea24f29f95ff28d6f8557`. Process run ID: `7dff58eda257a9fa5569817ca5a4db952fa495c1`. Observed: 2026-10-11T19:46:13.056073+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `29255d56e132ceef12fea24f29f95ff28d6f8557`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
aof-use-rdb-preamble
no
save

aof-load-truncated
yes
repl-diskless-sync
yes
rdb-del-sync-files
yes
appendfsync
everysec
no-appendfsync-on-rewrite
no
maxmemory
134217728
repl-diskless-load
swapdb
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
rdb_last_save_time:1791719156
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:1
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:548864
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

Original Redis Cluster ID: `b926f55d56cf8180e62c4bb92fb5110226f64356`. Process run ID: `242ca3b47566ffd1ba246ffddcb3711cf08f4a57`. Observed: 2026-10-11T19:46:14.088468+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b926f55d56cf8180e62c4bb92fb5110226f64356`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
maxmemory
134217728
appendfsync
everysec
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
appendonly
no
aof-use-rdb-preamble
no
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
repl-diskless-sync
yes
save

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
rdb_last_save_time:1791719158
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:532480
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

Original Redis Cluster ID: `fca8c4a47930f008988e801fcb806fdba30e5b14`. Process run ID: `a02cf019bed73e19e17633bde8d38c2e60f1b4e8`. Observed: 2026-10-11T19:46:15.133629+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `fca8c4a47930f008988e801fcb806fdba30e5b14`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
repl-diskless-load
swapdb
appendonly
no
maxmemory-policy
noeviction
save

aof-use-rdb-preamble
no
aof-load-truncated
yes
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791719160
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

Original Redis Cluster ID: `966f17b00394c5b5c7ac58428e19497a22b779e7`. Process run ID: `704efe4d3ad65495145652bb113bf10e6547ce78`. Observed: 2026-10-11T19:46:16.125168+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `966f17b00394c5b5c7ac58428e19497a22b779e7`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
save

aof-use-rdb-preamble
no
appendonly
no
auto-aof-rewrite-percentage
0
aof-load-truncated
yes
maxmemory-policy
noeviction
appendfsync
everysec
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
rdb-del-sync-files
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
rdb_last_save_time:1791719161
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

Original Redis Cluster ID: `8905830ae19a4accc8d3ea0580d11f00aaabf924`. Process run ID: `a5fe88f8f62782efe733636035b5954d23e57d48`. Observed: 2026-10-11T19:46:17.133657+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `8905830ae19a4accc8d3ea0580d11f00aaabf924`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
repl-diskless-load
swapdb
appendfsync
everysec
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
no
maxmemory
134217728
rdb-del-sync-files
yes
maxmemory-policy
noeviction
appendonly
no
aof-load-truncated
yes
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
rdb_last_save_time:1791719163
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

Fault action completed 2314.600 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 214 | 100.00% | 0 | 99.88 | 0.520 / 0.847 / 5.680 |

Acknowledged writes: 121. Ambiguous SET outcomes: 0.

Write policy: `set`. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors |
|---|---:|---:|
| baseline | 121 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 121 | 0.520 / 0.697 / 3.845 |
| baseline | set | 121 | 0.520 / 0.697 / 3.845 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 121 | 56.476 |

Global pacing: offered 214; started 214; completed 214; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 121; checked: 121; missing: 121; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `589ff7c7c45e95a58008f50f6ef0a9a01bf62aa88052e48d95927852327f35c2`.
