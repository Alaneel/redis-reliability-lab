# Redis Reliability Lab — primary-failure

Run: `243139221390b39795c9a0bd`  
Outcome: **partial**  
Started: 2026-10-11T16:43:47.560873+08:00  
Finished: 2026-10-11T16:44:42.082161+08:00

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

Frozen acknowledged SETs: 432. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:44:10.116637+08:00 |
| Immutable ledger frozen | 2026-10-11T16:44:10.116772+08:00 |
| All six stopped verified before restart | 2026-10-11T16:44:12.30094+08:00 |
| First restart requested | 2026-10-11T16:44:13.302512+08:00 |
| Restart sequence completed | 2026-10-11T16:44:22.666335+08:00 |
| Frozen-ledger validation completed | 2026-10-11T16:44:35.331349+08:00 |
| Post-validation diagnostics requested | 2026-10-11T16:44:35.331621+08:00 |

Drain to first kill request: 0.167 ms (runtime monotonic clock).

Drain to last kill completion: 1470.794 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1001.563 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | 16450.470 |
| Original topology sustained | 17904.132 |
| Replication redundancy sustained | 17904.132 |
| Post-validation diagnostic service sustained | 26644.225 |
| AOF frozen-value recovery observed | unobserved |

Configured memory-only control disables RDB/AOF value persistence. Persistence recovery is not applicable; restored processes and topology can coexist with lost frozen values. The independent persistence dimension is false when the disabled runtime configuration was observed, while the persistence-recovery event remains absent.

Frozen-ledger validation: expected 432; checked 432; missing 432; corrupt 0; unreadable 0; unverified 0; complete true. Completion describes verification coverage; observed loss remains a data failure.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:44:10.116803+08:00 | 2026-10-11T16:44:10.352808+08:00 |  |
| 2 | node-5 | 2026-10-11T16:44:10.352819+08:00 | 2026-10-11T16:44:10.6016+08:00 |  |
| 3 | node-6 | 2026-10-11T16:44:10.601617+08:00 | 2026-10-11T16:44:10.858997+08:00 |  |
| 4 | node-1 | 2026-10-11T16:44:10.859013+08:00 | 2026-10-11T16:44:11.093143+08:00 |  |
| 5 | node-2 | 2026-10-11T16:44:11.093154+08:00 | 2026-10-11T16:44:11.332268+08:00 |  |
| 6 | node-3 | 2026-10-11T16:44:11.332285+08:00 | 2026-10-11T16:44:11.587443+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:44:13.302512+08:00 | 2026-10-11T16:44:14.893014+08:00 |  |
| 2 | node-2 | 2026-10-11T16:44:14.893018+08:00 | 2026-10-11T16:44:16.433741+08:00 |  |
| 3 | node-3 | 2026-10-11T16:44:16.433753+08:00 | 2026-10-11T16:44:17.981295+08:00 |  |
| 4 | node-4 | 2026-10-11T16:44:17.981302+08:00 | 2026-10-11T16:44:19.592107+08:00 |  |
| 5 | node-5 | 2026-10-11T16:44:19.59211+08:00 | 2026-10-11T16:44:21.138777+08:00 |  |
| 6 | node-6 | 2026-10-11T16:44:21.138783+08:00 | 2026-10-11T16:44:22.666332+08:00 |  |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | a513f784a0b819b33ec364d20c0e4e916402651b | 2026-10-11T16:44:11.714412+08:00 | false | no | managed Redis process is stopped |
| node-5 | ede840c696cb1ba76d33577498dfcac661ce5af8 | 2026-10-11T16:44:11.831192+08:00 | false | no | managed Redis process is stopped |
| node-6 | af41380c5204875e4ac9b78860887ee8d4ef4595 | 2026-10-11T16:44:11.947684+08:00 | false | no | managed Redis process is stopped |
| node-1 | 56dbc41ed0d4bc74d9dadf10882c27d1b0c27083 | 2026-10-11T16:44:12.064091+08:00 | false | no | managed Redis process is stopped |
| node-2 | 3aa9e76b105dc2066f76488146f323dcda22de18 | 2026-10-11T16:44:12.180978+08:00 | false | no | managed Redis process is stopped |
| node-3 | b2148048350c10e564ed4318f9a5b4e3bdaa22a0 | 2026-10-11T16:44:12.300914+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `56dbc41ed0d4bc74d9dadf10882c27d1b0c27083`. Process run ID: `d05972d17ff61e502e97d02c4b70b67924aa7400`. Observed: 2026-10-11T16:44:00.900843+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `56dbc41ed0d4bc74d9dadf10882c27d1b0c27083`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
maxmemory-policy
noeviction
save

repl-diskless-sync
yes
maxmemory
134217728
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
appendonly
no
aof-use-rdb-preamble
no
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
rdb_last_save_time:1791708230
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

### Prepared node-2

Original Redis Cluster ID: `3aa9e76b105dc2066f76488146f323dcda22de18`. Process run ID: `c2cb8720bb940e08f186e0d6026ab9042c177992`. Observed: 2026-10-11T16:44:01.92132+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3aa9e76b105dc2066f76488146f323dcda22de18`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendonly
no
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
aof-load-truncated
yes
maxmemory-policy
noeviction
repl-diskless-load
swapdb
maxmemory
134217728
appendfsync
everysec
repl-diskless-sync
yes
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708231
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

### Prepared node-3

Original Redis Cluster ID: `b2148048350c10e564ed4318f9a5b4e3bdaa22a0`. Process run ID: `58cdd4fbb180641b12630ea1ecd0ecc744bb8387`. Observed: 2026-10-11T16:44:02.923222+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b2148048350c10e564ed4318f9a5b4e3bdaa22a0`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
save

aof-use-rdb-preamble
no
appendfsync
everysec
maxmemory
134217728
repl-diskless-load
swapdb
maxmemory-policy
noeviction
repl-diskless-sync
yes
appendonly
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
rdb_last_save_time:1791708231
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:499712
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

Original Redis Cluster ID: `a513f784a0b819b33ec364d20c0e4e916402651b`. Process run ID: `eeada5e986f3ef92c4487b68deb9d13f0b68b0af`. Observed: 2026-10-11T16:44:03.960773+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `a513f784a0b819b33ec364d20c0e4e916402651b`. Restart policy: `no`. `/data` usage: 8 KiB.

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
repl-diskless-sync
yes
repl-diskless-load
swapdb
appendonly
no
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
aof-load-truncated
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708232
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

Original Redis Cluster ID: `ede840c696cb1ba76d33577498dfcac661ce5af8`. Process run ID: `9dc3c0512fcad4ca1329ab9c7f5cf42728ca96e6`. Observed: 2026-10-11T16:44:04.963683+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ede840c696cb1ba76d33577498dfcac661ce5af8`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
aof-use-rdb-preamble
no
maxmemory-policy
noeviction
maxmemory
134217728
repl-diskless-load
swapdb
save

aof-load-truncated
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
rdb-del-sync-files
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
rdb_last_save_time:1791708232
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

Original Redis Cluster ID: `af41380c5204875e4ac9b78860887ee8d4ef4595`. Process run ID: `eb9850e4b60d8e3f0ca6bcc2171d007ce6c65d53`. Observed: 2026-10-11T16:44:05.974754+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `af41380c5204875e4ac9b78860887ee8d4ef4595`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
no-appendfsync-on-rewrite
no
auto-aof-rewrite-percentage
0
maxmemory
134217728
aof-use-rdb-preamble
no
appendfsync
everysec
save

rdb-del-sync-files
yes
repl-diskless-load
swapdb
aof-load-truncated
yes
maxmemory-policy
noeviction
repl-diskless-sync
yes
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
rdb_last_save_time:1791708232
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

Original Redis Cluster ID: `56dbc41ed0d4bc74d9dadf10882c27d1b0c27083`. Observed after restart: 2026-10-11T16:44:14.777011+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `56dbc41ed0d4bc74d9dadf10882c27d1b0c27083`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
appendfsync
everysec
appendonly
no
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
repl-diskless-sync
yes
aof-load-truncated
yes
save

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
rdb_last_save_time:1791708253
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

Original Redis Cluster ID: `3aa9e76b105dc2066f76488146f323dcda22de18`. Observed after restart: 2026-10-11T16:44:16.319564+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3aa9e76b105dc2066f76488146f323dcda22de18`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
appendonly
no
aof-use-rdb-preamble
no
aof-load-truncated
yes
save

auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
rdb-del-sync-files
yes
repl-diskless-sync
yes
maxmemory
134217728
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
rdb_last_save_time:1791708255
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

Original Redis Cluster ID: `b2148048350c10e564ed4318f9a5b4e3bdaa22a0`. Observed after restart: 2026-10-11T16:44:17.868992+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b2148048350c10e564ed4318f9a5b4e3bdaa22a0`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
appendonly
no
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
maxmemory
134217728
aof-use-rdb-preamble
no
appendfsync
everysec
rdb-del-sync-files
yes
repl-diskless-sync
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708256
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

Original Redis Cluster ID: `a513f784a0b819b33ec364d20c0e4e916402651b`. Observed after restart: 2026-10-11T16:44:19.477011+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `a513f784a0b819b33ec364d20c0e4e916402651b`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
save

aof-load-truncated
yes
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
aof-use-rdb-preamble
no
appendonly
no
repl-diskless-load
swapdb
maxmemory
134217728
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708258
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
/data/temp-1791708258.1.rdb

```

### Restart receipt 5: node-5

Original Redis Cluster ID: `ede840c696cb1ba76d33577498dfcac661ce5af8`. Observed after restart: 2026-10-11T16:44:21.024934+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ede840c696cb1ba76d33577498dfcac661ce5af8`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
no
repl-diskless-sync
yes
maxmemory-policy
noeviction
appendfsync
everysec
appendonly
no
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
save

maxmemory
134217728
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708259
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
/data/temp-1791708259.1.rdb

```

### Restart receipt 6: node-6

Original Redis Cluster ID: `af41380c5204875e4ac9b78860887ee8d4ef4595`. Observed after restart: 2026-10-11T16:44:22.554606+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `af41380c5204875e4ac9b78860887ee8d4ef4595`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
save

aof-use-rdb-preamble
no
repl-diskless-sync
yes
maxmemory
134217728
aof-load-truncated
yes
appendonly
no
repl-diskless-load
swapdb
no-appendfsync-on-rewrite
no
appendfsync
everysec
auto-aof-rewrite-percentage
0
rdb-del-sync-files
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
rdb_last_save_time:1791708261
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
/data/temp-1791708261.1.rdb

```

### Post-synchronization runtime evidence

These fresh read-only administrative observations follow topology stability and precede frozen-ledger validation. They establish the sampled persistence configuration, loading status and data files; the earlier physical restart receipts remain unchanged.

#### Restored node-1

Original Redis Cluster ID: `56dbc41ed0d4bc74d9dadf10882c27d1b0c27083`. Process run ID: `039f7d67b0522d2192f4f433223d82c978865ee3`. Observed: 2026-10-11T16:44:29.948609+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `56dbc41ed0d4bc74d9dadf10882c27d1b0c27083`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
appendfsync
everysec
appendonly
no
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
aof-use-rdb-preamble
no
repl-diskless-sync
yes
aof-load-truncated
yes
save

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
rdb_last_save_time:1791708253
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

#### Restored node-2

Original Redis Cluster ID: `3aa9e76b105dc2066f76488146f323dcda22de18`. Process run ID: `0b040847201a2137fc8cf56e4be7930e9246b51e`. Observed: 2026-10-11T16:44:31.048324+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3aa9e76b105dc2066f76488146f323dcda22de18`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
appendfsync
everysec
appendonly
no
aof-use-rdb-preamble
no
aof-load-truncated
yes
save

auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
rdb-del-sync-files
yes
repl-diskless-sync
yes
maxmemory
134217728
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
rdb_last_save_time:1791708255
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

#### Restored node-3

Original Redis Cluster ID: `b2148048350c10e564ed4318f9a5b4e3bdaa22a0`. Process run ID: `005d3cfa5f8ce634a8b3005278e6d312b7e58881`. Observed: 2026-10-11T16:44:32.081264+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `b2148048350c10e564ed4318f9a5b4e3bdaa22a0`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
appendonly
no
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
maxmemory
134217728
aof-use-rdb-preamble
no
appendfsync
everysec
rdb-del-sync-files
yes
repl-diskless-sync
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708256
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

Original Redis Cluster ID: `a513f784a0b819b33ec364d20c0e4e916402651b`. Process run ID: `ee480b22cedd2b96d3cf6e6cd67ac0beec128388`. Observed: 2026-10-11T16:44:33.123717+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `a513f784a0b819b33ec364d20c0e4e916402651b`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
maxmemory-policy
noeviction
save

aof-load-truncated
yes
rdb-del-sync-files
yes
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
aof-use-rdb-preamble
no
appendonly
no
repl-diskless-load
swapdb
maxmemory
134217728
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
rdb_changes_since_last_save:0
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708258
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

Original Redis Cluster ID: `ede840c696cb1ba76d33577498dfcac661ce5af8`. Process run ID: `23170db52f4e12ec3f12cc3a21fa2b55facc7137`. Observed: 2026-10-11T16:44:34.133077+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `ede840c696cb1ba76d33577498dfcac661ce5af8`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
no
repl-diskless-sync
yes
maxmemory-policy
noeviction
appendfsync
everysec
appendonly
no
auto-aof-rewrite-percentage
0
rdb-del-sync-files
yes
save

maxmemory
134217728
no-appendfsync-on-rewrite
no
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
rdb_last_save_time:1791708259
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

Original Redis Cluster ID: `af41380c5204875e4ac9b78860887ee8d4ef4595`. Process run ID: `abfb7f9b6a70bb5ce166746af80836853e68a71b`. Observed: 2026-10-11T16:44:35.175579+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `af41380c5204875e4ac9b78860887ee8d4ef4595`. Restart policy: `no`. `/data` usage: 8 KiB.

Observed persistence configuration:

```text
save

aof-use-rdb-preamble
no
repl-diskless-sync
yes
maxmemory
134217728
aof-load-truncated
yes
appendonly
no
repl-diskless-load
swapdb
no-appendfsync-on-rewrite
no
appendfsync
everysec
auto-aof-rewrite-percentage
0
rdb-del-sync-files
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
rdb_last_save_time:1791708261
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

Fault action completed 2184.123 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 828 | 100.00% | 0 | 199.93 | 0.408 / 0.520 / 0.664 |

Acknowledged writes: 432. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 432 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 432 | 0.428 / 0.520 / 0.664 |
| baseline | set | 432 | 0.408 / 0.495 / 0.632 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 432 | 104.312 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 828; started 828; completed 828; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: true; expected: 432; checked: 432; missing: 432; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: true.

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

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `4d51a2946a4630a2745a809e542f14a752a84eac8b35429962bf08a41f6174d1`.
