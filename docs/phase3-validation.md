# Phase 3 validation

The replica lifecycle milestone was exercised on 11 October 2026 using Go 1.26.9, macOS arm64 with 12 logical CPUs, Docker Engine 24.0.6, and Redis 8.10.2. Each fresh cluster used three primaries and three replicas, 256 MiB and one CPU per container, and private loopback-only client endpoints. The retained-data trials used six disposable, strictly labelled Docker volumes. Their filesystem has no hard quota; workload, payload, Redis memory, probe count and observation deadlines bound the generated traffic.

## Replica lifecycle observations

The standalone trials below all returned `recovered` and verified cleanup. Their original JSON and Markdown files were copied verbatim and compared byte for byte with generated reports. These are single-host observations from small samples, not production recovery guarantees.

| Scenario | Run ID | Ordinary service stable after fault start | Repair action | Stable redundancy after repair completed | Acknowledged values checked | Evidence |
|---|---|---:|---:|---:|---:|---|
| Replica SIGKILL and restart | `abd3356a3b9c2474c06cb6e6` | 1.447321 s | 1.519627 s | 1.238892 s | 936 | [JSON](examples/phase3/replica-crash/report.json), [Markdown](examples/phase3/replica-crash/report.md) |
| Replica whole-node isolation and reconnect | `a3f6bb0e004815a4c522c39d` | 3.089971 s | 1.239845 s | 1.569599 s | 1,069 | [JSON](examples/phase3/replica-isolation/report.json), [Markdown](examples/phase3/replica-isolation/report.md) |
| Primary SIGKILL, failover, former-primary repair with SET+WAIT | `7ce5fd7f49d14f7c38460823` | 4.664917 s | 1.553769 s | 5.913530 s | 1,325 | [JSON](examples/phase3/primary-repair-wait/report.json), [Markdown](examples/phase3/primary-repair-wait/report.md) |

All 3,330 acknowledged immutable values were checked: none were missing, corrupt, unreadable or unverified. The replica trials observed no ambiguous SET outcomes. The primary repair trial recorded 55 ambiguous SET outcomes, excluded from the acknowledged ledger; its verified ledger included 1,070 WAIT-satisfied writes, 254 short-WAIT writes and one SET acknowledgement followed by a WAIT error. No loss was observed in these trials; asynchronous replication does not guarantee this result in future failures.

In the primary trial the original replica was first sampled owning the failed slots at 3.472588 s. Ordinary service stabilized with a stopped former primary and degraded redundancy. Starting that primary preserved its Redis Cluster identity while replacing its process. Redis demoted it to a replica of the promoted owner; the unaffected owners remained unchanged. Six-node redundancy stabilized at 12.241066 s and the configured WAIT policy at 12.244734 s from fault orchestration start. The report therefore separates serving requests from restored replication and acknowledgement availability.

The replica crash trial observed one increase in the primary's partial-sync acceptance counter. The primary repair trial observed one full-sync and one partial-sync rejection counter increase on the promoted primary. These are server-wide counter observations from the same primary process and replication history. They do not identify an exact transfer duration or prove that a particular replica completed a transfer. Replica link-up and `master_sync_in_progress=0`, correct assignment, membership convergence and sustained all-shard probes are evaluated separately.

The three-second replica network hold produced administrative evidence of the same running Redis process on a detached sole bridge endpoint. It produced no sampled INFO-level replication interruption and no synchronization counter increase. Reconnection may occur before a surviving TCP connection times out; absent events remain unobserved. This fault removes the whole interface in both directions and does not represent a selective partition.

## Reproduction

```sh
make check build
bin/redis-lab run experiments/replica-crash.yaml
bin/redis-lab run experiments/replica-isolation.yaml
bin/redis-lab run experiments/primary-repair-wait.yaml
```

## Replica lifecycle integration validation

After a real integration run exposed a trailing-empty `CONFIG GET save` parsing defect, the parser was corrected and regression-tested. All six lifecycle Docker tests then passed with race detection in **149.325 s**:

```sh
go test -race -tags=integration -count=1 -timeout=5m \
  -run '^TestDocker(ReplicaCrashRestart|ReplicaNetworkIsolation|PrimaryRepairRestoresRedundancy|ReplicaRepairTimeout|ReplicaRepairAmbiguousFailureCleansVolumes|CancellationDuringReplicaSynchronization)$' \
  -v ./internal/experiment
```

The three happy-path suite runs were separate fresh clusters from the standalone examples above: replica crash `9e093ff339a2c491569f1c6d` verified 942 values, replica isolation `3086489a904ed89d7b913f46` verified 1,094, and primary repair `acb2ae25a9349bd0f7f5f0e7` verified 1,382. No missing or corrupt acknowledged values were observed in these suite cases.

| Failure-path suite case | Run ID | Outcome | Restart observed | Stable redundancy | Evidence |
|---|---|---|---|---|---|
| 100 ms repair budget | `f5a5a102abc9f92056580064` | failed, repair deadline expired | No accepted completed action | Unobserved | [JSON](examples/phase3/repair-failure-paths/f5a5a102abc9f92056580064/report.json), [Markdown](examples/phase3/repair-failure-paths/f5a5a102abc9f92056580064/report.md) |
| Error injected after successful restart | `ab7af306ada2be3182c8226c` | failed | Yes, error retained | Unobserved | [JSON](examples/phase3/repair-failure-paths/ab7af306ada2be3182c8226c/report.json), [Markdown](examples/phase3/repair-failure-paths/ab7af306ada2be3182c8226c/report.md) |
| Cancellation after restart, during synchronization | `488a31734f8e918c27af7ca0` | aborted | Yes | Unobserved | [JSON](examples/phase3/repair-failure-paths/488a31734f8e918c27af7ca0/report.json), [Markdown](examples/phase3/repair-failure-paths/488a31734f8e918c27af7ca0/report.md) |

Every failure-path case verified container, network and retained-volume cleanup using a fresh bounded cleanup context. Validation was skipped in these early exits: zero default validation counters do not establish zero loss, and acknowledged records remain unchecked. The ambiguous-response case deliberately injects a manager-boundary error after a real applied restart; it does not claim a naturally occurring Docker transport failure. Cancellation preserves the applied physical restart without claiming synchronization completion. Independent label-filtered Docker inventories were empty after the suite.

The earlier full integration run passed all nine existing Docker cases and five of the six new lifecycle cases; its single parser-related failure is recorded as a failure rather than a full-suite pass. The final expanded full-suite run passed after persistence and WAITAOF implementation; its result is recorded below.


## Persistence process-crash observations

All three real Docker mode tests passed with race detection in **167.351 s** after correcting the initial setup defects described below. Every test drained the running workload, froze its final acknowledged-write ledger, killed the three original replicas before the three original primaries, verified all six processes stopped with automatic restart disabled, held that barrier for one second, and restarted the same six containers on their retained managed volumes. Validation read the frozen ledger before any diagnostic SET or acknowledgement-policy probe.

| Mode | Run ID | Outcome | Acknowledged / checked | Missing | Original topology stable after fault start | Diagnostic service stable after fault start | Evidence |
|---|---|---|---:|---:|---:|---:|---|
| Memory | `243139221390b39795c9a0bd` | partial | 432 / 432 | 432 | 17.904132 s | 26.644225 s | [JSON](examples/phase3/persistence-memory/report.json), [Markdown](examples/phase3/persistence-memory/report.md) |
| AOF everysec | `8010385e1d747640095e5e8f` | recovered | 433 / 433 | 0 | 17.639774 s | 25.624264 s | [JSON](examples/phase3/persistence-aof-everysec/report.json), [Markdown](examples/phase3/persistence-aof-everysec/report.md) |
| AOF always | `7d3b278ba09b73b1cab9adf0` | recovered | 422 / 422 | 0 | 18.178602 s | 26.197466 s | [JSON](examples/phase3/persistence-aof-always/report.json), [Markdown](examples/phase3/persistence-aof-always/report.md) |

None of these trials recorded corrupt, unreadable or unverified values, or ambiguous SET outcomes. All verified cleanup. Memory restored processes, topology, redundancy and diagnostic service while losing the entire frozen ledger, so it correctly remained partial. AOF preservation was established only after final fresh runtime configuration, quiet loading/rewrite status, actual retained file inventory, original topology and complete read validation were observed. Early restart receipts remain raw observations; final evidence is collected separately after recovery.

The host, Docker VM, filesystem and storage remained running. Draining and serial SIGKILL took about 1.5 seconds from the first to last kill, permitting background fsync. AOF everysec and always preserving these small ledgers does not establish equivalent durability, a strict one-second loss bound or survival of physical power loss. Replication after restart can reconcile node files; the result does not prove every node's AOF independently held every write. The Docker VM reported 13 CPUs, 8,232,312,832 bytes of memory and the overlay2 storage driver; host runtime information is recorded separately.

The first mode suite failed rather than passed. Redis 8.10.2's bundled RedisSearch module requires `aof-use-rdb-preamble yes`; the attempted pure-command AOF base prevented AOF containers from starting. The corrected modes retain a valid hybrid RDB-format base, incremental AOF and manifest. The memory control also exposed a diskless-sync fallback that briefly left a replication RDB file. Redis-owned `rdb-del-sync-files yes` and a quiet, file-inspected preparation barrier now establish that the memory control has no persistence files. No data files are manually removed to manufacture the result. See [durability semantics](durability-semantics.md) for the runtime checks and Redis source references.

```sh
go test -race -tags=integration -count=1 -timeout=5m \
  -run '^TestDockerPersistenceModes$' -v ./internal/experiment
bin/redis-lab run experiments/persistence-memory.yaml
bin/redis-lab run experiments/persistence-aof-everysec.yaml
bin/redis-lab run experiments/persistence-aof-always.yaml
```

These example reports came from the integration cases, whose 200 ms sampling interval differs from the standalone YAML's 250 ms; the saved effective configuration is authoritative. The memory control exits nonzero for its expected partial outcome. An actual everysec report also passed the complete batch evidence verifier, including saved artifacts, ledger accounting, all-process barrier, identity-preserving restart, restored runtime evidence and recovery milestones.

## Persistence failure-path validation

All three additional real Docker tests passed with race detection in **89.145 s**:

```sh
go test -race -tags=integration -count=1 -timeout=5m \
  -run '^TestDockerPersistence(RestartFailureCleansVolumes|RestartTimeout|CancellationAfterRestart)$' \
  -v ./internal/experiment
```

| Case | Run ID | Outcome | Evidence |
|---|---|---|---|
| Error injected after the first real restart | `3c36066384bad6e63009a90d` | failed | [JSON](examples/phase3/persistence-failure-paths/3c36066384bad6e63009a90d/report.json), [Markdown](examples/phase3/persistence-failure-paths/3c36066384bad6e63009a90d/report.md) |
| 100 ms restart budget | `8fc02121276b3406a7f0405f` | failed | [JSON](examples/phase3/persistence-failure-paths/8fc02121276b3406a7f0405f/report.json), [Markdown](examples/phase3/persistence-failure-paths/8fc02121276b3406a7f0405f/report.md) |
| Cancellation after the first real restart | `593fa25caaada4b356468e01` | aborted | [JSON](examples/phase3/persistence-failure-paths/593fa25caaada4b356468e01/report.json), [Markdown](examples/phase3/persistence-failure-paths/593fa25caaada4b356468e01/report.md) |

Each case reached the six-process stopped barrier, preserved its failed or cancelled action and attempted fresh-context cleanup. Every cleanup succeeded, and independent managed-label Docker container, network and volume inventories were empty. Validation was skipped; default zero validation counters establish no preservation claim. The injected-response case is a deliberate manager-boundary error after an applied action, rather than evidence of a natural Docker transport failure. Saved failure reports were copied byte for byte.


## Counterbalanced persistence comparison

The real CLI comparison `comparison-dde55f9ada692afef5838308` completed all nine fresh trials from 16:50:18 to 16:58:02 Singapore time on 11 October 2026. The order seed `20261011` selected everysec → memory → always, then memory → always → everysec, then always → everysec → memory. Each mode occupied each within-round position once. Workload seed `42` and effective settings stayed unchanged; the actual order, configuration hashes, resolved Redis image, host and Docker VM metadata are saved with every trial.

```sh
bin/redis-lab compare --rounds 3 --seed 20261011 --output reports \
  experiments/persistence-memory.yaml \
  experiments/persistence-aof-everysec.yaml \
  experiments/persistence-aof-always.yaml
```

| Mode | Recovered / partial / failed | Acknowledged values checked | Missing | Median original-topology stability from fault start | Median diagnostic service stability from fault start | Median per-trial baseline write p50 / p95 / p99 |
|---|---|---:|---:|---:|---:|---|
| Memory | 0 / 3 / 0 | 1,251 | 1,251 | 17.654187 s | 25.047236 s | 0.471758 / 0.731856 / 1.521483 ms |
| AOF everysec | 3 / 0 / 0 | 1,251 | 0 | 17.705670 s | 25.112851 s | 0.546120 / 0.806872 / 0.980759 ms |
| AOF always | 3 / 0 / 0 | 1,252 | 0 | 17.633924 s | 25.028941 s | 0.980759 / 1.597558 / 2.478342 ms |

All 3,754 acknowledged immutable values were checked. No trial recorded ambiguity, corruption, unreadability or incomplete validation. Every report passed batch integrity/measurement verification, every cleanup succeeded, and independent owned-label inventories were empty after the cohort. The comparison completed six recovered AOF trials and three partial memory controls, with no unstarted or incomplete-evidence entries. Its exit code was correctly nonzero because the memory control's observed data loss prevents full recovery.

Latency values are medians of three individual successful logical SET latency quantiles measured during baseline. They are histogram estimates, not pooled request quantiles or reliable population tail estimates. The three baseline write-latency observations per mode do not establish significance, a universal ranking or a durability bound. This was a shared development host; counterbalancing reduces systematic position confounding but does not eliminate other runtime variability. Cyclic order also does not balance all carryover pairs. Fault and restart phases had no offered workload, so their timings are topology/diagnostic milestones rather than workload outage or throughput measurements.

The [comparison JSON](examples/phase3/persistence-comparison/comparison.json) and [Markdown](examples/phase3/persistence-comparison/comparison.md) link to all condition summaries, individual one-trial batches and raw reports. All 44 saved files were copied byte for byte. The comparison JSON SHA-256 is `083bd2d707e9c04f84bee5a8a61e2fa5c4baf36949dc51d6e4d34803c8f1df2e`. The runtime recorded Redis 8.10.2 and the same resolved image identity across conditions. These persistence comparisons preceded the new common AOF policy-preflight addition; their saved effective configuration and method are the executed evidence.


## SET+WAITAOF failover and repair

The standalone CLI run `34b6d0e4fd105ef1b3e0e9e0` used [the matched AOF-always WAITAOF definition](../experiments/policy-aof-waitaof.yaml) and returned `recovered`. Before baseline, it recorded six quiet, actual mode-matched AOF runtime observations and three successful shard probes returning local/replica counts `1/1`. No preflight or explicit fsync barrier followed baseline traffic.

```sh
bin/redis-lab run experiments/policy-aof-waitaof.yaml
```

The trial verified all 1,635 acknowledged immutable SETs with no missing, corrupt, unreadable or unverified values. Its 1,372 WAITAOF-satisfied writes formed a separately verified subset; 263 valid short replies returned `1/0`, so every observed reply satisfied the requested local AOF count but those writes lacked the requested replica AOF count. No WAITAOF errors or unissued calls occurred. The 37 ambiguous SET outcomes remain excluded from the acknowledged population.

Promotion was first sampled 3.539445 s after fault start and ordinary service stabilized at 4.752659 s. Same-container former-primary repair took 1.590417 s; restored redundancy stabilized 5.819080 s after repair completion. The configured WAITAOF policy stabilized at 12.282862 s after fault start, separately from earlier ordinary service. Six-node topology, replica readiness, policy availability and data verification were all observed. Persistence recovery is not claimed by this single-primary fault merely because AOF acknowledgements were received.

The [JSON](examples/phase3/waitaof-repair/report.json) and [Markdown](examples/phase3/waitaof-repair/report.md) were copied byte for byte. JSON SHA-256: `24afb519ec3f7956b83ed89be80ac43487f2402c906bca6fc80ce6030dcbe1b9`. Cleanup succeeded and independent owned-label inventories were empty. The observed `1/1` fsync acknowledgements and post-experiment values are separate facts; they do not establish distributed transactions, linearizability, loss-free failover or survival of physical power loss. [Redis WAITAOF semantics](https://redis.io/docs/latest/commands/waitaof/)


## WAITAOF real command and frozen-ledger validation

The primary-repair and all-six-process persistence tests passed in the first WAITAOF integration invocation (39.98 s and 56.77 s respectively). The invocation was not a full pass: its negative replica compatibility case selected a primary through a client routing helper. A second invocation exposed the same targeting assumption. The corrected test binds each replica to the managed readiness inventory and verifies live `INFO replication` role `slave` on the same physical connection. That corrected test passed in **16.681 s**. This was a test targeting defect; the two earlier invocations remain failed checks.

The compatibility case observed three memory-only primaries accepting immutable SETs, returning the expected appendonly-disabled WAITAOF error, and retaining the SET value on GET. Each of the three verified replicas rejected WAITAOF with the replica-instance error without receiving a SET. The [raw successful command log](examples/phase3/waitaof-command-compatibility.log) preserves the actual errors and owned endpoints.

The all-process AOF-always WAITAOF trial `adcd47811ff02526680c06fa` froze and checked **433 acknowledged values**, all with satisfied local/replica counts `1/1`, with zero missing, corrupt, unreadable, unverified or ambiguous values. It verified all six processes stopped before any restart and checked the frozen ledger before diagnostic writes. Persistence recovery and cleanup were verified. Its [JSON](examples/phase3/waitaof-persistence/report.json) and [Markdown](examples/phase3/waitaof-persistence/report.md) were copied byte for byte. The retained-volume process crash establishes observed value preservation in this failure model, not physical power-loss durability.

```sh
go test -race -tags=integration -count=1 -timeout=5m \
  -run '^TestDockerWAITAOF(PrimaryRepairRestoresPolicy|FrozenPersistenceValidation|DisabledRejectsCommandRetainsSET)$' \
  -v ./internal/experiment
```

The final complete integration result is recorded below after all regression cases passed. Independent inventories using the actual `org.redis-reliability-lab.managed=true` ownership label were empty after the corrected compatibility case.


## Fresh-source build verification

A fresh source copy excluding Git metadata, generated binaries and report output passed `make check build`: gofmt verification, `go vet ./...`, and fresh `go test -race -count=1 ./...` across all ten packages. `go mod verify` returned `all modules verified`. All 86 Go/source/build files were compared byte for byte with the working source. Both trimmed, VCS-independent builds produced identical binary SHA-256 `46759133fc5e41b026a93384525e1ef0634522f3d50b74bbb071bec91ced16cd`. The [raw check/build log](examples/phase3/clean-source-check.log) retains the actual results.

The final CLI validated all 16 YAML examples. The three AOF policy comparison definitions were checked to differ only in experiment name and write policy. This verifies the current source independently of existing binaries; this repository has no committed checkout or remote yet. Hosted GitHub Actions is configured but was not executed in this local validation.


## Counterbalanced AOF acknowledgement-policy comparison

The final real CLI comparison `comparison-321638985ff49f9040049bf4` completed all nine fresh trials from **17:11:38 to 17:17:22 Singapore time on 11 October 2026**, exiting successfully. All conditions used AOF `appendfsync always`, primary SIGKILL followed by same-container former-primary repair, six actual quiet persistence observations and three selected-policy compatibility probes before baseline. The three definitions differ only in name and write policy. Order seed `20261011` selected WAIT → SET → WAITAOF, then SET → WAITAOF → WAIT, then WAITAOF → WAIT → SET; workload seed remained `42`. Every policy occupied each round position once.

```sh
bin/redis-lab compare --rounds 3 --seed 20261011 --output reports \
  experiments/policy-aof-set.yaml \
  experiments/policy-aof-wait.yaml \
  experiments/policy-aof-waitaof.yaml
```

| Policy | Recovered / requested | Acknowledged / checked | Ambiguous SETs | Satisfied policy / valid short replies | Median individual baseline write p50 / p95 / p99 |
|---|---:|---:|---:|---:|---|
| SET | 3 / 3 | 4,646 / 4,646 | 124 | 4,646 / not applicable | 1.029797 / 1.521483 / 2.360325 ms |
| SET+WAIT | 3 / 3 | 4,575 / 4,575 | 124 | 3,907 / 668 | 1.677436 / 3.012442 / 3.661644 ms |
| SET+WAITAOF | 3 / 3 | 4,759 / 4,759 | 111 | 4,029 / 730 | 1.597558 / 2.868992 / 3.321219 ms |

All **13,980 acknowledged immutable values** were checked with zero missing, corrupt, unreadable or unverified records. The **359 ambiguous SET outcomes** remain outside that population. Every short WAIT/WAITAOF result stayed in the full SET ledger. WAITAOF returned 4,029 observed `1/1` pairs and 730 observed `1/0` pairs: all 4,759 replies met the local AOF threshold, while 4,029 also met the replica AOF threshold. Neither acknowledgement policy recorded a command error or unissued post-SET call. Their satisfied subsets were completely checked with zero observed loss.

| Policy | Median ordinary service stable after fault start | Median repair action | Median redundancy stable after repair completion | Median configured policy stable after fault start |
|---|---:|---:|---:|---:|
| SET | 4.523550 s | 1.510390 s | 5.350976 s | Not a separate acknowledgement event |
| SET+WAIT | 4.717017 s | 1.541732 s | 4.983597 s | 11.080822 s |
| SET+WAITAOF | 4.656215 s | 1.550036 s | 5.459118 s | 11.530280 s |

All nine trials verified restored six-node redundancy, with the former primary following the promoted owner and unchanged ownership on unaffected shards. Each recorded a same-primary-process full-sync counter increase of one; this is server-wide synchronization-request evidence rather than an exact replica transfer duration. The policy and redundancy stable intervals are independent observations. These primary faults do not test all-process persistence recovery; that event is correctly counted as not at risk, rather than a zero-duration success.

Healthy baselines recorded 1,249 successful logical writes for SET, 1,250 for WAIT and 1,248 for WAITAOF. The latency table uses only successful logical `set`, `set_wait` and `set_waitaof` measurements, excluding GET, probes, fault phases and diagnostic command-stage populations. Values are medians of three histogram-estimated per-trial quantiles, not pooled request quantiles or population tail estimates. WAITAOF's observed median was below WAIT's in this cohort; this does not establish a universal performance ordering or equivalent semantics. All trials used the same host, Docker VM and resolved Redis image; position balance does not remove environmental variability or balance all carryover pairs.

The [comparison JSON](examples/phase3/policy-comparison/comparison.json) and [Markdown](examples/phase3/policy-comparison/comparison.md) retain the actual schedule, effective configurations and hashes, runtime metadata, conditional timing denominators, all per-condition summaries and individual raw reports. All **44 files** were copied byte for byte; comparison JSON SHA-256 is `f17f14c1584530ff5b4ad798c7ce9c3451d68e4575151f8725d3891aa0c56644`. No trial was failed, partial, unstarted or excluded for incomplete evidence. Every cleanup was verified, and independent inventories using `org.redis-reliability-lab.managed=true` were empty before the final regression suite.


## Final complete Docker regression suite

The frozen final source passed the complete real Docker integration suite with race detection in **781.960 s** on 11 October 2026:

```sh
make integration
# go test -race -tags=integration -count=1 -timeout=20m -v ./internal/experiment
```

All **22 top-level Docker test groups** passed, including all three memory/everysec/always persistence subtests. The command also executed the deterministic experiment-package tests included under the integration build tag. There were no failed test cases. The [complete unmodified test log](examples/phase3/full-integration.log) records every test name, fault-stage result and generated run ID. Earlier failed parser, persistence setup and replica-targeting checks described above remain historical failed checks; this fresh expanded suite establishes the final regression result.

Coverage includes original primary failover, recovery deadline, cancellation and repeated timeout experiments; replica crash, isolation, former-primary repair, repair timeout, ambiguous applied restart and cancellation during synchronization; primary network isolation, reconciliation deadline and interrupted restoration; frozen-ledger persistence in all three modes, applied restart failure, restart timeout and cancellation; legacy WAIT degradation, WAITAOF repair and frozen persistence, and actual role-verified WAITAOF command rejection. All integration checks used the managed local Docker environment. Failed/partial/aborted experimental outcomes in expected failure-path tests are retained and checked rather than relabelled recovered.

Independent inventories using the exact `org.redis-reliability-lab.managed=true` ownership label were empty for containers, networks and volumes after the suite. The earlier fresh-source formatting, vet, ten-package race checks, build and dependency verification passed on the same unchanged Go source. Linux GitHub Actions remains configured rather than locally verified on a separate Linux host.


## Final original CLI regression

The unchanged README command ran on a fresh cluster after the complete integration suite and exited **0** with `recovered`:

```sh
bin/redis-lab run -output reports experiments/primary-failure.yaml
```

Run `5a38321b7d2a351e9600b705` checked all **2,030 acknowledged immutable values** with zero missing, corrupt, unreadable or unverified records. Its **87 ambiguous SET outcomes** remain separate. Promotion was first sampled 5.370301 s and combined ordinary service stability 8.567055 s after fault orchestration start. The default intentionally leaves the killed primary stopped and has no repair/persistence-recovery claim. Its [JSON](examples/phase3/primary-failure-regression/report.json) and [Markdown](examples/phase3/primary-failure-regression/report.md) were copied byte for byte. Cleanup succeeded; final independent managed-label inventories were empty for containers, networks and volumes.
