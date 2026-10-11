# Measurement methodology

Repeated experiments measure variability that one successful failover cannot establish. The CLI runs each trial sequentially in a fresh six-node cluster, records the effective configuration and resolved image, and verifies cleanup before provisioning the next cluster. A batch is a collection of independent resource lifecycles on the same host, rather than independent physical machines.

```sh
redis-lab repeat experiments/primary-failure-repeat.yaml --runs 5
```

The full default primary-failure definition is also supported. The shorter repeat definition changes baseline duration, node timeout, observation spacing, and stability duration; compare trials within a definition, not those different configurations as if they were identical.

## Comparing configurations in a recorded order

Use `compare` for multiple conditions. Put flags before the configuration files:

```sh
redis-lab compare --rounds 3 --seed 20261011 --output reports \
  experiments/persistence-memory.yaml \
  experiments/persistence-aof-everysec.yaml \
  experiments/persistence-aof-always.yaml
```

Conditions are numbered in file order. A seeded base permutation rotates by one position each round. With N conditions, each condition occupies every position once in a complete block of N rounds. An incomplete block has partial balance. This rotating Latin order does not balance all predecessor/successor pairs or remove host-load and carryover effects. The order seed changes only the schedule; it does not replace each effective workload seed. Supported bounds are 2–8 conditions, 1–100 rounds, and at most 100 total trials.

`comparison.json` and `comparison.md` preserve every planned position, round, condition, configuration hash, outcome and actual trial reference, including unstarted positions. Each position contains one fresh single-trial batch. Per-condition `conditions/condition-XX/summary.json` and Markdown use the same failure-aware batch statistics. Schedule paths are relative to the comparison directory; condition-summary paths are relative to that summary directory. Individual JSON references retain host and Docker metadata, resolved image, effective configuration, raw observations and runtime controls. Each summary file is replaced atomically, while the JSON/Markdown pair is not a filesystem transaction.

The comparison proceeds after a failed or partial trial only when its persisted evidence is verified and cleanup is successful, or the report establishes that no environment was provisioned. Missing or mismatched evidence, unsafe cleanup, cancellation and summary-write failures stop further provisioning. A nonzero CLI status preserves valid partial/failed evidence and does not by itself mean execution stopped early. Compare `stop_reason`, the full schedule and per-condition denominators before interpreting results.

Ordering improves position balance within this host execution, without establishing independent samples or causal effects. Configuration hashes identify exact effective configurations; a matching hash does not establish identical host conditions. Inspect CPU architecture, host CPU count, Docker VM CPU/memory settings, storage driver and image digest in each trial. Host and Docker VM CPU counts describe different environments and may differ.

## Populations and missing observations

`summary.json` records requested, invoked, completed, and unstarted trials. Every invoked trial remains represented, including setup failures, data-loss partial outcomes, deadline failures, and cancellation. If cleanup cannot be verified, the batch stops before another environment is provisioned. Cancellation likewise stops further trials. Remaining requested trials are explicitly unstarted.

Outcome fractions and timing statistics answer different questions. Recovery success uses all requested trials as its denominator, with started-trial fractions also reported. Timing distributions describe observed milestones conditional on those milestones being observed. A partial trial can have a valid promotion measurement even if acknowledged-write verification later fails. A setup failure never entered the fault-observation population.

An unobserved recovery event after an injected fault is retained with its observation horizon and censor reason. It is neither zero duration nor evidence that Redis can never recover. A deadline means the observer did not accept recovery before that deadline. Polling can miss an event, and cancellation can end observation early. Conditional timing statistics do not estimate the unconditional recovery-time distribution or solve censoring with a survival model.

Descriptive statistics use exact trial values: minimum, median, arithmetic mean, sample standard deviation when at least two observations exist, and maximum. Trial p90/p95 values are withheld below 20 observations. Even above that threshold, they are descriptive estimates with sampling uncertainty, not confidence bounds. Five trials are enough to test the pipeline and show a small observed range, not characterize a production tail.

## Fault clocks and observation uncertainty

Reports distinguish fault orchestration beginning, fault action completion, first failure flag, first observed promotion, first successful client probe, and stable recovery. Runtime interval calculations retain Go's monotonic clock. Wall timestamps remain useful for correlating evidence; they do not replace runtime duration measurements.

Milestone `since_fault_ms` preserves the Phase 1 orchestration anchor. `fault_action_duration_ms` records the duration of ownership checks plus the action. Batch summaries and Markdown also expose milestone durations relative to action completion. The exact physical fault can occur before Docker returns, so these anchors bound an orchestration interval rather than identify an exact Redis event time.

The configured sample interval is a delay after a completed sample. Direct observation, operation deadlines, client probes, and scheduling add to the actual spacing. Each snapshot retains start/end and per-node timestamps. Concurrent node views are not an atomic distributed snapshot. A sample crossing the recovery deadline remains raw evidence but cannot establish accepted recovery.

A stable interval means consecutive positive topology and client observations span the required duration. It does not prove continuous availability between samples. Workload errors, one-second windows, and post-recovery checks provide additional evidence, without converting sampled probes into a universal availability guarantee.

## Workload and latency populations

Record configured offered rate separately from actual request counts. Worker saturation can reduce started traffic; it must not be hidden by labelling completed throughput as offered load. There is no unbounded queue or catch-up burst. A logical write operation includes all policy work associated with that immutable record; policy stages are not extra offered requests.

Global counters account for elapsed ideal periods, including skipped timer periods. Busy workers and unavailable eligible operations drop offers. After stopping and draining, `offered = started + dropped` and `started = completed = logical metrics attempts`. These counters exclude validation and both service and acknowledgement-policy probes. Logical `set`, `set_wait` and `set_waitaof` writes span their complete routing and command sequence. Diagnostic SET/WAIT/WAITAOF histograms do not add requests. The SET stage includes routing and pre-send errors. An attempted acknowledgement command means its client method was invoked, not proof the command reached Redis. Policy probes follow ordinary service probes, so they can increase actual sample spacing.

Metrics distinguish completed operations, successful and failed operations, terminal error classes, and separate all/success/failure latency distributions. Each operation belongs to the phase in which it starts. Throughput divides completed requests by phase elapsed time, and successful throughput uses successes. Fault-phase pooled availability divides total successes by total attempts across trials; the distribution of per-trial success fractions is a separate statistic.

Per-trial request quantiles use bounded geometric histograms with 5% bucket growth. Batch and comparison summaries preserve the distribution of per-trial quantile estimates. They do **not** average those estimates and call the result a pooled request percentile. Phase-level quantiles describe the mixed logical request population when reads and writes are both present. Write-policy latency comparisons must use the individual phase's successful `set`, `set_wait` or `set_waitaof` operation summaries. A median across those per-trial p95 estimates remains a median of trial estimates, not a pooled request p95. Individual JSON reports retain operation-level summaries, phase durations, availability windows, configuration, and topology source replies.

## Correctness and external validity

Every attempted write has a unique immutable key and deterministic value. Read-back tests cannot hide an earlier lost write behind a later overwrite. A successful `SET` enters the acknowledgement ledger. A client error is an ambiguous outcome: the command might have executed, and lack of an acknowledgement is not proof of loss.

Validation runs after workload scheduling stops and in-flight requests settle. Missing, corrupt, unreadable, and unchecked values are separate outcomes. Deadline-expired validation is incomplete, not successful. Policy-specific replica acknowledgements must remain separate from `SET` acknowledgement and final verification.

WAIT accounting distinguishes requested count, integer acknowledgement replies, short replies, attempted-command errors/client timeouts, and SET acknowledgements for which WAIT could not be attempted. WAITAOF additionally preserves each canonical `local/replicas` reply pair and independently counts local, replica and both-threshold satisfaction. Local and replica counts overlap. Errors and unattempted calls have unknown acknowledgement counts; they are not observed `0` or `0/0` replies. Every successful SET enters the main ledger, and each policy-satisfied subset is validated separately through the same final reads. A short reply or client error does not undo SET. Batch integrity checks reconcile phase and total counters, reply histograms, diagnostic attempts and logical write success/error counts before aggregation. Skipped subset validation remains unverified. Recovered policy trials require sustained policy availability and complete subset preservation. [Policy semantics and comparison](consistency-comparison.md)

## Retained process-crash measurements

The `cluster-restart` experiment stops and drains traffic, freezes the immutable SET ledger, kills original replicas before primaries, and verifies all six stopped processes with inspected restart policy `no` before restarting any. It restarts original primaries before replicas in the same containers and retained managed volumes. Frozen-ledger read-back completes before diagnostic SET/GET or policy probes issue new writes. No workload traffic runs during the crash or restoration phases, so they do not supply workload availability or recovery-throughput measurements.

Read-only administrative preparation runs before baseline traffic. It records actual CONFIG values, persistence and keyspace INFO, bounded `/data` file listings and data footprint for all six owned processes, and requires quiet persistence and synchronized replication. WAITAOF also requires a successful before-baseline all-shard compatibility probe after preparation. Capturing this evidence before traffic avoids inserting an additional fsync-producing probe between draining and killing. The runner issues no SAVE, BGSAVE or BGREWRITEAOF, and adds no durability barrier between traffic drain and the crash.

Actual modes distinguish disabled memory persistence, AOF everysec, and AOF always. Scheduled RDB snapshots are disabled. Growth-triggered AOF rewriting is disabled, while successful full synchronization can still trigger rewriting; AOF RDB preambles remain enabled for the pinned image's module compatibility. Rewrite-time fsync suppression is disabled, and `aof-load-truncated yes` is recorded. The memory control verifies the absence of value-persistence files before baseline. Managed volumes are disposable and have no filesystem quota; finite workload/probes, container limits and Redis no-eviction memory bounds limit this lab's resources.

This failure model keeps the host and storage running. Draining, sequential kills and inspection allow background fsync to progress; Docker VM and host page caches remain. Preserved everysec values therefore do not establish a one-second-loss bound, equality with always, or resilience to power/storage failure. Runtime readiness, restored topology, frozen-value preservation and later diagnostic service are separate observations. See the [executed Phase 3 evidence](phase3-validation.md).

### Recorded persistence comparison

The [preserved comparison](examples/phase3/persistence-comparison/comparison.md) ran nine fresh clusters on 2026-10-11, from 08:50:18 to 08:58:02 UTC. File-order conditions were memory, everysec and always; actual order was `2,1,3 / 1,3,2 / 3,2,1`, with order seed 20261011 and workload seed 42. All nine trials completed with verified evidence and cleanup, with no early stop. The expected memory-control losses produced three partial outcomes and a nonzero CLI status.

| Condition / summary | Recovered / partial | SET acknowledged / checked | Missing | Median trial successful write p50 / p95 / p99, ms |
|---|---:|---:|---:|---|
| [Memory](examples/phase3/persistence-comparison/conditions/condition-01/summary.md) | 0 / 3 | 1,251 / 1,251 | 1,251 | 0.472 / 0.732 / 1.521 |
| [AOF everysec](examples/phase3/persistence-comparison/conditions/condition-02/summary.md) | 3 / 0 | 1,251 / 1,251 | 0 | 0.546 / 0.807 / 0.981 |
| [AOF always](examples/phase3/persistence-comparison/conditions/condition-03/summary.md) | 3 / 0 | 1,252 / 1,252 | 0 | 0.981 / 1.598 / 2.478 |

All three cohorts had zero ambiguous SET outcomes, corrupt, unreadable and unverified records. Latency cells are medians of three **per-trial successful logical SET** baseline quantile estimates, excluding GET and diagnostic stages. They are bounded-histogram estimates and are not pooled request percentiles. This rate-limited, small-sample execution does not measure capacity or establish a performance ranking or durability guarantee.

The 44 preserved JSON/Markdown files match the original bytes. The [comparison JSON](examples/phase3/persistence-comparison/comparison.json) retains every schedule position, raw trial link, configuration hash and environment reference; its SHA-256 is `083bd2d707e9c04f84bee5a8a61e2fa5c4baf36949dc51d6e4d34803c8f1df2e`. Trials record 12 host logical CPUs and 13 Docker VM CPUs, with 1 CPU/256 MiB per container and Redis maxmemory 128 MiB. These separate environments and limits belong to this execution. This cohort preceded the common AOF policy-probe warmup addition; each preserved trial's methodology describes its executed controls.

### Recorded AOF policy comparison

The [matched AOF-always policy comparison](examples/phase3/policy-comparison/comparison.md) completed nine fresh primary-crash-and-repair trials from 09:11:38 to 09:17:22 UTC on 2026-10-11. SET, WAIT and WAITAOF each occupied every round position once, in actual order `2,1,3 / 1,3,2 / 3,2,1`; schedule seed 20261011 and workload seed 42 remained separate. All nine recovered with verified evidence and cleanup. Common preparation included six quiet AOF runtime samples and a successful three-shard policy compatibility probe before each baseline.

| Policy | Acknowledged SETs checked | Missing / corrupt / unreadable / unverified | Ambiguous SET outcomes | Median trial successful write p50 / p95 / p99, ms |
|---|---:|---|---:|---|
| SET | 4,646 | 0 / 0 / 0 / 0 | 124 | 1.029797 / 1.521483 / 2.360325 |
| WAIT | 4,575 | 0 / 0 / 0 / 0 | 124 | 1.677436 / 3.012442 / 3.661644 |
| WAITAOF | 4,759 | 0 / 0 / 0 / 0 | 111 | 1.597558 / 2.868992 / 3.321219 |

These latency values are medians of three per-trial successful logical-write baseline estimates, excluding mixed GET traffic and diagnostic stages. They do not pool request quantiles. WAITAOF's lower observed median values than WAIT's do not establish universal speed or a causal ranking. The configured 200-request/s mixed workload measures acknowledgement cost under these conditions, without establishing maximum throughput.

All 13,980 acknowledged SETs were checked, including 668 short WAIT and 730 short WAITAOF writes. The separate satisfied subsets contained 3,907 WAIT records and 4,029 WAITAOF records, both completely checked with no observed loss. WAITAOF's paired histogram was `1/1: 4,029; 1/0: 730`, preserving local and replica AOF acknowledgement semantics. The 359 ambiguous SET outcomes are outside the verified population. See the [policy comparison](consistency-comparison.md) for counts, recovery anchors, pacing and interpretation.

All 44 cohort files were preserved byte for byte. The [comparison JSON](examples/phase3/policy-comparison/comparison.json) SHA-256 is `f17f14c1584530ff5b4ad798c7ce9c3451d68e4575151f8725d3891aa0c56644`; its schedule links to every raw report and runtime/configuration reference. These nine trials are separate from the historical persistence-disabled, unrepaired Phase 2 SET/WAIT batches and the earlier all-process persistence cohort.

Record image digest, Redis version, host and Docker environment, effective configuration, client routing/retry policy, workload rate/concurrency/mix/payload, failure timing, and initial health. Sources of variability include election jitter, gossip, replication scheduling, host load, connection-pool contention, image/provisioning activity, and observation spacing. Trials share a host and may remain correlated despite fresh containers. Results establish observations within these settings, not durability, linearizability, or cross-host production behavior.
