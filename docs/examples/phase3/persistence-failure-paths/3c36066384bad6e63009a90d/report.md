# Redis Reliability Lab — primary-failure

Run: `3c36066384bad6e63009a90d`  
Outcome: **failed**  
Started: 2026-10-11T16:47:03.512919+08:00  
Finished: 2026-10-11T16:47:32.425119+08:00

## Recovery

| Condition | Observed |
|---|---|
| Original six-node topology recovered | false |
| Post-validation diagnostic service sustained | false |
| Topology and diagnostic service jointly sustained | false |
| Frozen acknowledged writes verified | false |
| All managed processes restarted | false |
| Original replication redundancy restored | false |
| Restoration deadline expired | false |
| Diagnostic deadline expired | false |
| Cleanup verified | true |


## Frozen-ledger process restart

Configured persistence mode: `aof-always`. Failure model: `sequential SIGKILL of all six Redis processes, same-container restart; host and storage remain running`.

Workload traffic stops and drains before the sequential SIGKILL actions. No workload requests are offered during fault or restart observation; these phases do not measure zero availability or recovery throughput. Read-only frozen-ledger validation precedes separate diagnostic writes.

Frozen acknowledged SETs: 432. Frozen ambiguous SET outcomes: 0.

Restoration deadline expired: false.

| Evidence anchor | Observed time |
|---|---|
| Traffic fully drained | 2026-10-11T16:47:25.291832+08:00 |
| Immutable ledger frozen | 2026-10-11T16:47:25.291991+08:00 |
| All six stopped verified before restart | 2026-10-11T16:47:27.530685+08:00 |
| First restart requested | 2026-10-11T16:47:28.533036+08:00 |
| Restart sequence completed | unobserved |
| Frozen-ledger validation completed | unobserved |
| Post-validation diagnostics requested | unobserved |

Drain to first kill request: 0.200 ms (runtime monotonic clock).

Drain to last kill completion: 1510.899 ms (runtime monotonic clock).

All-stopped barrier to first restart request: 1002.343 ms (runtime monotonic clock).

| Restart observation | Since fault orchestration began, ms |
|---|---:|
| Original six-node replication ready | unobserved |
| Original topology sustained | unobserved |
| Replication redundancy sustained | unobserved |
| Post-validation diagnostic service sustained | unobserved |
| AOF frozen-value recovery observed | unobserved |

AOF frozen-value recovery observed: unobserved. This criterion requires actual loading/configuration evidence and complete frozen-value preservation; restart or topology readiness alone cannot establish it.

Frozen-ledger validation unobserved; 432 acknowledged values remain unchecked.

### Kill action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-4 | 2026-10-11T16:47:25.292031+08:00 | 2026-10-11T16:47:25.545175+08:00 |  |
| 2 | node-5 | 2026-10-11T16:47:25.545196+08:00 | 2026-10-11T16:47:25.800194+08:00 |  |
| 3 | node-6 | 2026-10-11T16:47:25.80022+08:00 | 2026-10-11T16:47:26.044611+08:00 |  |
| 4 | node-1 | 2026-10-11T16:47:26.04463+08:00 | 2026-10-11T16:47:26.294063+08:00 |  |
| 5 | node-2 | 2026-10-11T16:47:26.294072+08:00 | 2026-10-11T16:47:26.534415+08:00 |  |
| 6 | node-3 | 2026-10-11T16:47:26.534429+08:00 | 2026-10-11T16:47:26.802743+08:00 |  |

### Restart action order

| Order | Managed node | Requested | Completed | Error |
|---:|---|---|---|---|
| 1 | node-1 | 2026-10-11T16:47:28.533036+08:00 | unobserved | injected ambiguous all-process restart response |

### Stopped administrative observations

| Managed node | Original Redis Cluster ID | Observed | Container running | Restart policy | Status/error |
|---|---|---|---|---|---|
| node-4 | 38fad44ef90b0dc7688f37d5fe5458ff7afb287d | 2026-10-11T16:47:26.930777+08:00 | false | no | managed Redis process is stopped |
| node-5 | 32675455610bc02bc9e4ef01873162549dd10ad2 | 2026-10-11T16:47:27.049192+08:00 | false | no | managed Redis process is stopped |
| node-6 | 3aff3b88873034951def3585214465a7c7256e9f | 2026-10-11T16:47:27.169069+08:00 | false | no | managed Redis process is stopped |
| node-1 | 2bee0be9aea1afc66d0d756cecbab9d53f6fb557 | 2026-10-11T16:47:27.290696+08:00 | false | no | managed Redis process is stopped |
| node-2 | d67e5ad9f55156becef18bd91683fe942c6e9863 | 2026-10-11T16:47:27.411743+08:00 | false | no | managed Redis process is stopped |
| node-3 | c04f64531e8f03e7ebbc5b002e7c6bb5c682c412 | 2026-10-11T16:47:27.530655+08:00 | false | no | managed Redis process is stopped |

### Prepared node-1

Original Redis Cluster ID: `2bee0be9aea1afc66d0d756cecbab9d53f6fb557`. Process run ID: `ecc99fa458733428dca68124406cca8f694818e3`. Observed: 2026-10-11T16:47:16.083299+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2bee0be9aea1afc66d0d756cecbab9d53f6fb557`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
rdb-del-sync-files
yes
appendonly
yes
save

maxmemory-policy
noeviction
aof-load-truncated
yes
appendfsync
always
repl-diskless-sync
yes
maxmemory
134217728
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
repl-diskless-load
swapdb
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
rdb_last_save_time:1791708426
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

### Prepared node-2

Original Redis Cluster ID: `d67e5ad9f55156becef18bd91683fe942c6e9863`. Process run ID: `2df8fd0f20f1db2298b07a930550ee201b2b3e42`. Observed: 2026-10-11T16:47:17.094577+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `d67e5ad9f55156becef18bd91683fe942c6e9863`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-use-rdb-preamble
yes
auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
save

appendonly
yes
repl-diskless-load
swapdb
aof-load-truncated
yes
appendfsync
always
maxmemory-policy
noeviction
rdb-del-sync-files
yes
repl-diskless-sync
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
rdb_last_save_time:1791708426
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
rdb_saves:0
rdb_saves_consecutive_failures:0
rdb_last_cow_size:561152
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

Original Redis Cluster ID: `c04f64531e8f03e7ebbc5b002e7c6bb5c682c412`. Process run ID: `9d87e120fa1dc7ae57ffc10701682612871de2f4`. Observed: 2026-10-11T16:47:18.083847+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `c04f64531e8f03e7ebbc5b002e7c6bb5c682c412`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
maxmemory
134217728
no-appendfsync-on-rewrite
no
aof-use-rdb-preamble
yes
appendonly
yes
rdb-del-sync-files
yes
maxmemory-policy
noeviction
repl-diskless-sync
yes
aof-load-truncated
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
rdb_last_save_time:1791708427
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

Original Redis Cluster ID: `38fad44ef90b0dc7688f37d5fe5458ff7afb287d`. Process run ID: `396610fb753d447fa79790b591fc2995849966ab`. Observed: 2026-10-11T16:47:19.065575+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `38fad44ef90b0dc7688f37d5fe5458ff7afb287d`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
rdb-del-sync-files
yes
repl-diskless-sync
yes
save

aof-use-rdb-preamble
yes
no-appendfsync-on-rewrite
no
appendfsync
always
appendonly
yes
repl-diskless-load
swapdb
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
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
rdb_last_save_time:1791708427
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

Original Redis Cluster ID: `32675455610bc02bc9e4ef01873162549dd10ad2`. Process run ID: `8ff2786ab038de02247ce933800775baf8279209`. Observed: 2026-10-11T16:47:20.111588+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `32675455610bc02bc9e4ef01873162549dd10ad2`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
repl-diskless-sync
yes
rdb-del-sync-files
yes
appendfsync
always
auto-aof-rewrite-percentage
0
repl-diskless-load
swapdb
maxmemory
134217728
aof-use-rdb-preamble
yes
maxmemory-policy
noeviction
no-appendfsync-on-rewrite
no
appendonly
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
rdb_last_save_time:1791708427
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

Original Redis Cluster ID: `3aff3b88873034951def3585214465a7c7256e9f`. Process run ID: `abc2c7f300570eef73135603b2cff1ba22e443bf`. Observed: 2026-10-11T16:47:21.147694+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `3aff3b88873034951def3585214465a7c7256e9f`. Restart policy: `no`. `/data` usage: 20 KiB.

Observed persistence configuration:

```text
aof-load-truncated
yes
no-appendfsync-on-rewrite
no
appendonly
yes
repl-diskless-load
swapdb
appendfsync
always
auto-aof-rewrite-percentage
0
maxmemory-policy
noeviction
rdb-del-sync-files
yes
save

aof-use-rdb-preamble
yes
repl-diskless-sync
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
rdb_last_save_time:1791708428
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
aof_last_cow_size:827392
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

Original Redis Cluster ID: `2bee0be9aea1afc66d0d756cecbab9d53f6fb557`. Observed after restart: 2026-10-11T16:47:30.10522+08:00.

Container running: true. Redis PONG: true. Redis Cluster ID: `2bee0be9aea1afc66d0d756cecbab9d53f6fb557`. Restart policy: `no`. `/data` usage: 52 KiB.

Observed persistence configuration:

```text
maxmemory
134217728
appendfsync
always
save

auto-aof-rewrite-percentage
0
no-appendfsync-on-rewrite
no
repl-diskless-sync
yes
rdb-del-sync-files
yes
appendonly
yes
aof-use-rdb-preamble
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
rdb_changes_since_last_save:144
rdb_bgsave_in_progress:0
rdb_last_save_time:1791708448
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
aof_current_size:31629
aof_base_size:108
aof_pending_rewrite:0
aof_buffer_length:0
aof_pending_bio_fsync:0
aof_delayed_fsync:0
```

Observed INFO keyspace:

```text
# Keyspace
db0:keys=144,expires=0,avg_ttl=0,subexpiry=0
```

Observed managed data files:

```text
/data/nodes.conf
/data/appendonlydir/appendonly.aof.manifest
/data/appendonlydir/appendonly.aof.1.incr.aof
/data/appendonlydir/appendonly.aof.1.base.rdb

```

Post-synchronization administrative runtime observations: unobserved.

This is a sequential process crash and same-container restart with retained managed volumes. It does not remove Docker VM/host page cache or simulate power loss, storage loss, or a graceful shutdown. Draining and sequential kill/inspection time allow background fsync to progress, so everysec preservation does not establish a one-second-loss guarantee or equivalence to always.

Persistence experiment error: restart node-1: injected ambiguous all-process restart response

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

Fault action completed 2238.639 ms after orchestration began (measured with the monotonic runtime clock).

No primary promotion is required. Original owners and replica assignments must remain unchanged across the all-stopped barrier; frozen data and post-validation diagnostics are evaluated separately.

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|
| baseline | 828 | 100.00% | 0 | 199.83 | 0.807 / 1.449 / 2.248 |

Acknowledged writes: 432. Ambiguous SET outcomes: 0.

Write policy: `set`. Logical write latency includes SET and, for `wait`, WAIT; policy success requires the requested replica count. Diagnostic stage requests are excluded from logical workload counts.

| Phase | SET acknowledged | SET errors | WAIT attempted | WAIT satisfied | WAIT short | WAIT errors | WAIT client timeouts |
|---|---:|---:|---:|---:|---:|---:|---:|
| baseline | 432 | 0 | 0 | 0 | 0 | 0 | 0 |

| Phase | Measured path | Count | p50 / p95 / p99, ms |
|---|---|---:|---|
| baseline | Logical write completion | 432 | 1.081 / 1.598 / 2.360 |
| baseline | set | 432 | 1.081 / 1.598 / 2.360 |

| Phase | Successful logical writes | Successful writes/s over phase wall time |
|---|---:|---:|
| baseline | 432 | 104.258 |

SET acknowledgements without a subsequent WAIT attempt: 0; client timeouts among those: 0. WAIT command errors and WAIT client timeouts above count attempted WAIT client calls only.

Global pacing: offered 828; started 828; completed 828; dropped 0. These counts include all workload phases and exclude validation and probes.

Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases.

Validation attempted: false; expected: 0; checked: 0; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: false.
Validation was not attempted; 432 acknowledged writes remain unchecked.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|

Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here.

## Errors

- restart node-1: injected ambiguous all-process restart response

## Measurement methodology

- Quiet administrative process, CONFIG, persistence INFO, keyspace and bounded data-file preparation occurs before baseline traffic; it issues no SAVE, BGSAVE, BGREWRITEAOF or explicit durability barrier.
- The workload fully stops and joins in-flight acknowledgements before freezing its immutable ledger and counters. All later topology observations are read-only until every frozen acknowledged value has been checked. SET-error outcomes remain ambiguous, excluded from acknowledgement guarantees.
- Original replicas are SIGKILLed before original primaries. Every fixed owned container must produce ErrProcessStopped administrative evidence with restart policy no before any restart. Original primaries restart before replicas in the same containers and retained volumes; original Redis Cluster identities and slot ownership must survive.
- The model is sequential process SIGKILL while Docker VM, host and storage remain running, followed by same-container starts. It does not simulate power loss, host/page-cache loss or disk loss. Drain, kill and inspection intervals are reported because they permit background fsync progress.
- Topology restoration requires all six original identities, endpoints, primary slot owners and replica assignments synchronized for stable_for, without forced Redis topology commands. Physical receipts retain early raw evidence, which can be transient during full sync. Six fresh quiet administrative runtime/configuration/file observations after topology stability prove the final persistence mode on those same restarted processes, within the shared repair budget and before frozen-value validation. Samples crossing restoration or diagnostic deadlines are retained but excluded from successful milestones.
- Frozen-ledger GET validation happens exactly once before diagnostic SET/GET or configured replica-acknowledgement probes. Process/slot readiness and later diagnostic service are independent from frozen-value preservation; missing or corrupt values retain a partial data verdict even after service recovers.
- Memory mode has RDB/AOF disabled, rdb-del-sync-files=yes to remove replication-created RDB files, and observed absence of value persistence files. Its persistence-recovery event is not applicable and remains absent; its independent persistence dimension is false. AOF modes use the RDB base preamble required by the pinned image's bundled modules plus incremental AOF records. AOF recovery requires actual matching healthy runtime configuration, loading=0 and complete frozen acknowledged-value preservation; this observation establishes no loss-free or power-loss guarantee.
- No workload demand is offered after draining, so absent fault/recovery traffic does not measure zero availability or throughput. Post-validation probes are separate diagnostics and never enter workload metrics or the frozen ledger. Poll intervals follow sample completion; observations are sampled and non-atomic.

The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: `5f3362a0bdd93bf4a884c5958983a2e641eef0c9e963ba05b2dc59469d6b507f`.
