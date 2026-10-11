# Comparing SET, WAIT and WAITAOF

## Status

The matched AOF-always comparison completed nine real trials on 2026-10-11: three each for SET, WAIT and WAITAOF, all recovered with verified evidence and cleanup. All 13,980 acknowledged immutable SETs and both policy-satisfied subsets were checked with zero missing, corrupt, unreadable or unverified records. The 359 error-returning SET outcomes remain ambiguous and outside that verified population. The complete [comparison and individual evidence](examples/phase3/policy-comparison/comparison.md) are preserved.

The historical Phase 2 cohort below contains six primary-crash trials from earlier that day: three SET trials recovered, while three WAIT trials restored ordinary service but remained partial because the promoted primary had no replica. All six completed validation and cleanup. The complete [SET batch](examples/policy-set-batch/summary.md) and [WAIT batch](examples/policy-wait-batch/summary.md), including all individual JSON/Markdown reports, are preserved verbatim; SHA-256 comparisons verified all 16 copied files against the original reports. These persistence-disabled, unrepaired trials do not measure WAITAOF or the new AOF repair comparison.

## Policies and the meaning of success

`set` completes when Redis returns a successful SET response. `wait` performs SET followed by `WAIT 1 100` and completes successfully only when SET succeeds and WAIT returns at least one replica acknowledgement.

`waitaof` performs SET followed by `WAITAOF 1 1 100` on the same physical connection. Success requires both local and replica AOF acknowledgement thresholds. Its reply is a two-integer pair; a short pair is distinct from a client error. WAITAOF requires Redis 7.2 or later, and this lab restricts it to retained AOF configurations on every managed node. Its fsync acknowledgement differs from WAIT's replica-memory acknowledgement and does not establish strong consistency. See the [Redis WAITAOF command](https://redis.io/docs/latest/commands/waitaof/).

WAIT covers preceding writes on the **same client connection**, using the client's replication offset. Its integer reply is a replica count, including when the requested count is not reached. The timeout is in milliseconds; zero allows unbounded waiting and is rejected here. In transactions or scripts WAIT cannot provide the same blocking behavior. A low count does not undo a successful SET. See the [Redis WAIT command](https://redis.io/docs/latest/commands/wait/).

Redis normally replicates asynchronously. Replica receipt does not establish persistence, linearizability, or a universally loss-free failover. The historical Phase 2 trials disable persistence and test final readability and value equality of immutable records. Current retained AOF experiments record actual persistence settings and process-restart preservation separately. See [Redis replication](https://redis.io/docs/latest/operate/oss_and_stack/management/replication/).

## Connection and routing controls

All three policies use the same managed writer, slot routing, pool limits, endpoint allowlist, operation deadline, and redirect rules. GET reads continue through the cluster client. The writer selects the cached primary for the key and borrows a sticky connection from its own pool. SET and optional WAIT/WAITAOF run sequentially on that connection before release.

Only a Redis MOVED or ASK response for the key's actual slot may redirect an unacknowledged SET, with a limit of eight redirects and an exact managed-address match. ASKING is issued immediately before SET on an ephemeral connection that is discarded afterward. Transport failures are never replayed, and WAIT/WAITAOF are never rerouted or retried after a SET acknowledgement. A new connection or another cluster node would not be evidence for the original SET.

The implementation targets pinned go-redis v9.23.0. Inspect [the managed writer](../internal/experiment/writer.go) and [the workload policy accounting](../internal/workload/policy.go). The client operation context bounds the entire routing/SET/acknowledgement sequence, independently of the server command's timeout.

## Ledgers and accounting

Every SET acknowledged by Redis enters the full immutable-write ledger, including SETs followed by a short WAIT/WAITAOF reply, a command error, or cancellation before the acknowledgement command can start. The final validator checks the full ledger. It separately checks the subset whose configured acknowledgement thresholds were satisfied. A successful GET of a later overwrite cannot conceal loss of an earlier record because each SET has a unique key.

| Observation | Full SET ledger | Configured policy-satisfied subset | Policy result |
|---|---|---|---|
| SET success under `set` | Included | Excluded | Successful |
| SET success; WAIT reaches one replica | Included | Included | Successful |
| SET success; WAIT returns a lower count | Included | Excluded | Failed policy acknowledgement |
| SET success; WAIT error or cannot start | Included | Excluded | Failed policy acknowledgement |
| SET success; WAITAOF reaches both local and replica thresholds | Included | Included in WAITAOF subset | Successful |
| SET success; WAITAOF returns a short pair | Included | Excluded | Failed policy acknowledgement |
| SET success; WAITAOF error or cannot start | Included | Excluded | Failed policy acknowledgement |
| SET error, including timeout or disconnection | Excluded; counted separately as an ambiguous outcome | Excluded | Failed |

The counters preserve SET acknowledgements and errors, each policy's attempted and unattempted commands, valid reply distribution, threshold satisfaction, short replies, errors and client timeouts. WAITAOF histograms retain canonical `local/replicas` pairs. Independent local and replica satisfaction counts overlap; both thresholds must hold in the same reply for policy success. Errors have unknown counts, while `0/0` is an observed short pair. Policy probes retain separate attempt and reply-observed masks. Records excluded from the ledger may still have executed, so exclusion narrows the verified set instead of proving absence.

One schedule slot represents one logical operation: GET, SET, SET plus WAIT, or SET plus WAITAOF. Diagnostic stages do not add another offered or completed logical request. Reports retain separate stage latency summaries. Use the logical `set`, `set_wait` or `set_waitaof` operation's success count and latency population for write comparisons; phase totals that also include GET describe the mixed workload. Batch integrity checks reconcile logical write attempts/successes/errors with SET and policy counters before statistics include them.

## Service availability and policy availability

Ordinary all-shard SET/GET probes measure service recovery. Separate probes measure whether the requested write policy can be satisfied. After a primary crash, its promoted replica can serve requests while having no replica of its own. With this lab's original primary kept stopped, WAIT's one-replica requirement can remain unsatisfied on that shard even after service recovers. That is a redundancy and acknowledgement-policy limitation, and should not be described as failure of ordinary client service. The [Redis Cluster specification](https://redis.io/docs/latest/operate/oss_and_stack/reference/cluster-spec/) describes asynchronous replication, failover, and the possibility of acknowledged-write loss.

## Current controlled comparison commands

The current matched AOF policy definitions all use `aof-always` and primary SIGKILL followed by repair. They differ in experiment name and write policy. Before baseline traffic, all six processes must have matching quiet AOF runtime evidence, and an all-shard compatibility probe must succeed. For WAITAOF it observes local 1 and at least one replica AOF acknowledgement per shard. Preparation and this probe precede workload writes; they are separate from final immutable-ledger validation. Recovered WAIT/WAITAOF outcomes additionally require sustained policy success and complete preservation of both the full SET ledger and the policy-satisfied subset. SET has no separate acknowledgement-command subset.

```sh
redis-lab compare --rounds 3 --seed 20261011 --output reports \
  experiments/policy-aof-set.yaml \
  experiments/policy-aof-wait.yaml \
  experiments/policy-aof-waitaof.yaml
```

Flags precede configuration files. Condition order uses a seeded base permutation rotated each round; every condition occupies every round position once across three rounds. Workload seed 42 remains unchanged. JSON preserves the actual schedule, full effective configurations and SHA-256 hashes, outcomes and raw host/environment references. Verified partial/failed trials remain included, with censoring and missing observations separate from zero-duration events. Unsafe cleanup, missing evidence and cancellation leave remaining positions explicitly unstarted. See [measurement methodology](measurement-methodology.md) for ordering limits and report paths.

Fresh containers share the same physical host and Docker VM. Position balance does not remove scheduling, cache, synchronization or host-load effects. Three trials per condition support pipeline validation and descriptive observations, without causal rankings, reliable tail estimates or pooled request quantiles.

## Current AOF policy comparison results

Comparison `comparison-321638985ff49f9040049bf4` ran from 09:11:38.844832 to 09:17:22.385436 UTC on 2026-10-11. File-order conditions were SET, WAIT and WAITAOF; actual order was `2,1,3 / 1,3,2 / 3,2,1`, with order seed 20261011. Workload seed 42 remained unchanged. All nine requested trials started, completed and recovered, with no unstarted, incomplete-evidence or cleanup-failure entries and no early stop. Each trial retained six quiet before-baseline AOF preparations and a successful three-shard compatibility probe.

The common settings were Redis `8.10.2-alpine`, AOF always on all six nodes, primary SIGKILL followed by former-primary repair, a 4-second baseline, 2-second post-recovery observation, 8 workers, 200 offered logical requests/s, 50% configured reads, 128-byte values and at most 20,000 writes. WAIT and WAITAOF requested one replica with a 100 ms server timeout, inside the common 500 ms operation deadline. Host/runtime settings and resolved image identity are recorded per trial; source configuration hashes alone do not describe host or implementation changes.

### Healthy baseline write observations

These values use successful logical `set`, `set_wait` and `set_waitaof` operations from individual baseline reports, excluding GET and diagnostic stages. All baseline writes succeeded. Writes/s is the median of three per-trial successful-write counts divided by that trial's recorded phase duration.

| Policy / condition summary | Successful baseline writes across three trials | Median trial write p50, ms | Median trial write p95, ms | Median trial write p99, ms | Median successful writes/s |
|---|---:|---:|---:|---:|---:|
| [SET](examples/phase3/policy-comparison/conditions/condition-01/summary.md) | 1,249 | 1.029797 | 1.521483 | 2.360325 | 103.489 |
| [WAIT](examples/phase3/policy-comparison/conditions/condition-02/summary.md) | 1,250 | 1.677436 | 3.012442 | 3.661644 | 103.694 |
| [WAITAOF](examples/phase3/policy-comparison/conditions/condition-03/summary.md) | 1,248 | 1.597558 | 2.868992 | 3.321219 | 103.523 |

Quantiles are bounded-histogram estimates. Each cell is a median of three per-trial quantile estimates, **not a pooled request percentile**. WAITAOF's observed median quantiles were below WAIT's in this execution; this small sample does not establish that WAITAOF is universally faster. Both acknowledgement policies had higher observed median baseline write latency than SET. The rate-limited workload and shared host do not establish capacity, statistical significance, causal effects or a general policy ranking. Condition-summary phase quantiles describe mixed reads and writes and must not replace the logical-write values above.

### Acknowledgements, recovery and verification

| Policy | SET acknowledged / checked | Ambiguous SET outcomes | Policy-satisfied logical writes | Short acknowledgement replies | Satisfied subset checked |
|---|---:|---:|---:|---:|---:|
| SET | 4,646 / 4,646 | 124 | 4,646 | Not applicable | Not applicable |
| WAIT | 4,575 / 4,575 | 124 | 3,907 | 668 | 3,907 |
| WAITAOF | 4,759 / 4,759 | 111 | 4,029 | 730 | 4,029 |

All 4,575 WAIT attempts returned a replica count: 3,907 returned one and 668 returned zero. All 4,759 WAITAOF attempts returned a valid pair: 4,029 returned `1/1` and 730 returned `1/0`. Thus every WAITAOF reply satisfied the local AOF threshold, while 4,029 satisfied the replica threshold and both thresholds together. Independent local and replica counts overlap. Neither cohort recorded acknowledgement-command errors, client timeouts or successful SETs without a subsequent command attempt. The 668 short WAIT and 730 short WAITAOF SETs remain in the verified full ledger.

| Policy | Median ordinary service/topology stable after fault start, s | Median six-node redundancy stable after fault start, s | Median configured acknowledgement policy stable after fault start, s |
|---|---:|---:|---:|
| SET | 4.524 | 11.387 | Not applicable |
| WAIT | 4.717 | 11.388 | 11.081 |
| WAITAOF | 4.656 | 11.838 | 11.530 |

Ordinary service returned before repaired redundancy and acknowledgement availability were sustained. Policy and full redundancy are sampled separately, so a policy milestone can precede the accepted six-node redundancy interval. All six WAIT/WAITAOF policy milestones were observed; none were fabricated from missing or censored events. Their observed timing medians describe these three trials per policy. Main-ledger and policy-subset validation had zero missing, corrupt, unreadable or unverified records; that establishes observed final values for the acknowledged population, without proving universal loss prevention or physical power-loss durability. The 359 ambiguous SET outcomes remain outside that claim.

| Policy | Offered | Started | Completed | Dropped |
|---|---:|---:|---:|---:|
| SET | 10,477 | 9,385 | 9,385 | 1,092 |
| WAIT | 10,410 | 9,276 | 9,276 | 1,134 |
| WAITAOF | 10,724 | 9,600 | 9,600 | 1,124 |

These whole-workload counters include all traffic phases, satisfy offered = started + dropped and started = completed, and exclude administrative preparation, policy/service probes and final validation. Differences in total request counts reflect actual durations and worker availability, despite matched configured load.

All 44 [comparison, condition, single-trial batch and raw report files](examples/phase3/policy-comparison/comparison.md) were preserved byte for byte. The [comparison JSON](examples/phase3/policy-comparison/comparison.json) records the complete actual schedule and per-trial configuration/host/environment references; SHA-256: `f17f14c1584530ff5b4ad798c7ce9c3451d68e4575151f8725d3891aa0c56644`. The following historical cohort has different persistence, repair and timing controls and should not be pooled with this comparison.

## Historical persistence-disabled matched trials

Three fresh clusters per policy ran with [policy-set.yaml](../experiments/policy-set.yaml) and [policy-wait.yaml](../experiments/policy-wait.yaml). The effective configurations differ only in policy and experiment name. The six reports have unique run/network IDs and 36 unique container IDs.

To repeat the same comparison:

```sh
redis-lab repeat experiments/policy-set.yaml --runs 3 --output reports/phase2-policy
redis-lab repeat experiments/policy-wait.yaml --runs 3 --output reports/phase2-policy
```

The WAIT-policy batch returns a nonzero CLI status for these partial outcomes while preserving all reports and continuing the batch after verified cleanup.

| Control | Value |
|---|---|
| Redis image | `redis:8.10.2-alpine`; actual image ID recorded per trial |
| Topology | Three primaries, each with one replica before the fault |
| Fault | Primary SIGKILL; original primary remains stopped |
| Node timeout | 1 second |
| Healthy baseline / post-recovery observation | 8 seconds / 3 seconds |
| Workers / aggregate offered rate | 8 / 200 logical requests per second |
| Configured read percentage / seed | 50% / 42 |
| Payload / maximum writes | 128 bytes / 20,000 |
| Operation deadline | 500 milliseconds |
| WAIT policy | One replica / 100 milliseconds |
| Topology sample delay / stable interval | 250 milliseconds / 1 second |

The recorded host was macOS/arm64 with 12 logical CPUs and Go 1.26.9; Docker Desktop with Docker Engine 24.0.6 ran Linux kernel `6.4.16-linuxkit`. Each Redis container had one CPU, 256 MiB memory, and a 128 MiB tmpfs with persistence disabled. All six trials used image ID `sha256:675f7644831d85995868250485bf7a71970257f74da31c62969a868d3f3eff2e` and recorded digest `redis@sha256:3811787313eba226a2ef38658c6ccb91cd5e110edc89c37767de373120a0e5a0`.

The requested read percentage is a selection rule, not a promise of an exact completed mix. Slow or failed writes can change which workers are available and which acknowledged records can be read. Compare actual GET/write counts and global offered/started/completed/dropped counters. Workload saturation must remain visible.

Analyze **healthy baseline** metrics separately from fault and post-recovery metrics. Baseline logical write latency and successful write throughput estimate the observed acknowledgement cost before topology degradation. Fault-phase success counts, WAIT counts, service recovery, policy recovery, and both validation sets answer different questions.

The configured sample interval is a delay after a completed observation. Topology queries, ordinary SET/GET probes, and then policy probes also consume time; WAIT can lengthen that cadence. Milestones are first sampled observations, and successful probes do not establish uninterrupted availability between them. Preserve both fault-orchestration and action-completion anchors when comparing recovery times.

Execution used a block order: three SET trials from 07:40:37 to 07:42:04 UTC, then three WAIT trials from 07:42:19 to 07:43:46 UTC on the same host. This order was not randomized or counterbalanced: changing host load, cache state, Docker startup behavior, and time can confound the comparison. These results describe the executed trials, without causal or population guarantees. Later experiments should alternate or randomize policies and increase the trial count.

## Historical healthy baseline results

The healthy baseline completed 4,799 mixed requests for SET and 4,809 for WAIT, all successful. The table uses only logical write latency (`set` or `set_wait`), excluding GET and diagnostic stage counts. Because every baseline write succeeded, the all-write and successful-write distributions coincide here. Successful writes/s divides the trial's successful logical write count by its recorded baseline phase duration.

| Trial / individual JSON | Successful writes | Write p50, ms | Write p95, ms | Write p99, ms | Successful writes/s |
|---|---:|---:|---:|---:|---:|
| [SET 1](examples/policy-set-batch/runs/0207534b36545671d4c6bcb3/report.json) | 811 | 0.573 | 0.807 | 1.192 | 101.112 |
| [SET 2](examples/policy-set-batch/runs/fc9ec7fdcee8870aaa67e2ce/report.json) | 809 | 0.520 | 0.807 | 2.869 | 100.912 |
| [SET 3](examples/policy-set-batch/runs/461f750112eb15258eeb6892/report.json) | 806 | 0.697 | 2.869 | 7.250 | 100.424 |
| **SET median across three trials** | — | **0.573** | **0.807** | **2.869** | **100.912** |
| [WAIT 1](examples/policy-wait-batch/runs/92d758d529f45b4bb1836474/report.json) | 811 | 1.081 | 1.380 | 2.360 | 101.149 |
| [WAIT 2](examples/policy-wait-batch/runs/2f20da5016af35f23348d0de/report.json) | 810 | 1.081 | 1.521 | 2.869 | 101.009 |
| [WAIT 3](examples/policy-wait-batch/runs/07aadc068ea2995602d6959d/report.json) | 809 | 0.981 | 1.314 | 1.849 | 100.918 |
| **WAIT median across three trials** | — | **1.081** | **1.380** | **2.360** | **101.009** |

WAIT's median trial p50 was 0.508 ms higher in this execution. Both policies sustained approximately 101 successful writes/s under the configured mixed offered load; this rate-limited experiment does not measure maximum throughput or capacity. The median p99 was lower for WAIT, with substantial variation in the three SET trials, so these results do not establish a consistent tail-latency improvement or penalty.

The bold values are medians of **per-trial quantile estimates**, not pooled request percentiles. Report histograms quantize latency estimates. Three trials per policy do not support recovery-tail estimates, significance tests, or general policy rankings; no pooled latency quantiles are computed.

## Historical fault and post-service-recovery observations

The following counts include GET and logical writes, excluding probes and validation. Successful ops/s and successful writes/s divide summed counts by summed recorded phase durations. Actual durations include phase boundaries and in-flight completion; configured durations and offered rate do not replace those denominators.

| Policy / phase | Completed | Successful | Failed | Request success fraction | Successful ops/s | Successful writes/s |
|---|---:|---:|---:|---:|---:|---:|
| SET / baseline | 4,799 | 4,799 | 0 | 100.00% | 199.429 | 100.816 |
| WAIT / baseline | 4,809 | 4,809 | 0 | 100.00% | 199.931 | 101.026 |
| SET / fault | 1,527 | 1,279 | 248 | 83.76% | 96.945 | 47.298 |
| WAIT / fault | 1,611 | 1,210 | 401 | 75.11% | 86.830 | 36.813 |
| SET / post-recovery | 1,800 | 1,800 | 0 | 100.00% | 199.970 | 107.206 |
| WAIT / post-recovery | 1,800 | 1,478 | 322 | 82.11% | 157.761 | 68.847 |

The `post_recovery` phase begins after ordinary topology/service recovery; its name does not mean the WAIT policy recovered. Every one of the 322 failed post-recovery WAIT logical writes had a successful SET followed by an observed zero-replica reply. The promoted primary's final replication INFO reported zero connected replicas in each WAIT trial. The other two primaries retained their replicas; all-shard SET/GET probes succeeded while policy probes observed one shard returning zero and the other two returning one.

| Policy / phase | SET acknowledged | SET errors | WAIT attempted | WAIT reached one | WAIT returned zero | WAIT errors |
|---|---:|---:|---:|---:|---:|---:|
| SET / baseline | 2,426 | 0 | 0 | 0 | 0 | 0 |
| WAIT / baseline | 2,430 | 0 | 2,430 | 2,430 | 0 | 0 |
| SET / fault | 624 | 132 | 0 | 0 | 0 | 0 |
| WAIT / fault | 673 | 124 | 673 | 513 | 160 | 0 |
| SET / post-recovery | 965 | 0 | 0 | 0 | 0 | 0 |
| WAIT / post-recovery | 967 | 0 | 967 | 645 | 322 | 0 |

All 4,070 recorded workload WAIT attempts returned a count: 3,588 returned one and 482 returned zero. Errors/timeouts among attempted WAIT calls and SET acknowledgements without a subsequent WAIT attempt were both zero. A returned zero represents incomplete replica acknowledgement after the server's bounded wait; it does not establish that SET failed. Separate terminal request errors during the fault included CLUSTERDOWN, network errors, and client timeouts.

Whole-workload pacing remained observable:

| Policy | Offered | Started | Completed | Dropped |
|---|---:|---:|---:|---:|
| SET | 9,251 | 8,126 | 8,126 | 1,125 |
| WAIT | 9,397 | 8,220 | 8,220 | 1,177 |

These global counters satisfy offered = started + dropped and started = completed. They include all workload phases, with elapsed ideal schedule slots defining offered load. Busy workers and skipped slots remain visible instead of being interpreted as successful offered requests.

## Historical recovery, validation, and interpretation

| Observation | SET cohort | WAIT cohort |
|---|---:|---:|
| Requested / started / completed trials | 3 / 3 / 3 | 3 / 3 / 3 |
| Recovered / partial / failed / aborted | 3 / 0 / 0 / 0 | 0 / 3 / 0 / 0 |
| Final ordinary topology/client recovery | 3 / 3 | 3 / 3 |
| Replica-acknowledgement policy sustained after fault | Not applicable | 0 / 3 observed |
| Median ordinary combined recovery after fault action, ms | 4,257.376 | 4,294.363 |
| Full SET ledger checked | 4,015 | 4,070 |
| WAIT-satisfied subset checked | Not applicable | 3,588 |
| Missing / corrupt / unreadable / unverified ledger records | 0 / 0 / 0 / 0 | 0 / 0 / 0 / 0 |
| Ambiguous SET outcomes outside the ledger | 132 | 124 |
| Cleanup failures / incomplete report evidence | 0 / 0 | 0 / 0 |

The WAIT trials are accurately partial: ordinary service and data verification succeeded, but a one-replica write policy was unavailable on the promoted shard. No process restart or redundancy restoration is attempted in this primary-crash experiment. The SET outcome does not require replacement redundancy; the WAIT outcome additionally requires sustained all-shard policy success.

WAIT policy recovery was **unobserved by the accepted probe horizon**, approximately 7.240, 7.221, and 7.414 seconds after completed SIGKILL in trial order. The runner stops fault observation when ordinary recovery stabilizes and then observes the configured post-recovery phase; it does not necessarily observe WAIT policy for the full 30-second recovery ceiling. Batch JSON classifies these absent policy milestones as `censored_error` because each partial trial returns an explicit outcome error. They are neither zero-duration recoveries nor 30-second deadline timeouts. This censoring describes the observed horizon, not a later physical recovery event.

**No acknowledged-write loss observed across the executed trials.** All 8,085 successful SET acknowledgements were checked, including the 482 SETs whose WAIT returned zero. The WAIT-satisfied subset also had zero missing, corrupt, unreadable, or unverified records, and traffic read checks had no missing or mismatched values. The 256 error-returning SET outcomes remain ambiguous and outside that verified set. These observations establish final readability and value equality for the checked immutable records; they do not establish durability, linearizability, or universal loss prevention by WAIT.

The historical six-trial results above remain SET/WAIT evidence only. Use the AOF policy comparison definitions for WAITAOF and inspect its paired acknowledgement evidence separately; these historical numbers do not transfer to that experiment.
