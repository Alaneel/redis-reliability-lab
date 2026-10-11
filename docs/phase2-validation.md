# Phase 2 validation

Milestone A was executed on **11 October 2026** using **Go 1.26.9**, **macOS arm64**, and **Docker Desktop / Docker Engine 24.0.6**. Five fresh sequential primary-failure trials completed with exit code **0**, five recovered outcomes, and verified cleanup. These results validate the repeated-run pipeline and describe this host and configuration; they do not establish a production recovery tail or a Redis durability guarantee.

The complete original batch is preserved as [summary JSON](examples/phase2-crash-batch/summary.json), [summary Markdown](examples/phase2-crash-batch/summary.md), and the five individual JSON/Markdown reports linked below. All **12 files** were copied verbatim and their SHA-256 hashes compared with the generated originals. Recorded absolute paths describe the original execution location; relative report links remain usable in the copied example.

## Configuration and execution

Batch ID: `batch-e4b81963e956ed40e22b4a27`. The recorded batch interval was **15:15:35–15:17:15 Singapore time**, including provisioning and teardown. Reproduce the same experiment definition with:

```sh
./bin/redis-lab repeat experiments/primary-failure-repeat.yaml --runs 5 --output reports/phase2
```

The experiment used three primaries and three replicas, Redis **8.10.2**, a **1 s** Redis node timeout, **2 s** baseline, **250 ms** delay between completed observation samples, **1 s** required stability, **1 s** post-recovery observation, and **30 s** recovery deadline. Workload settings were **8 workers**, **200 offered operations/s**, **50% configured reads**, **128-byte values**, seed **42**, a **20,000 attempted-write cap**, and **500 ms** operation deadlines. Writes used normal `SET` acknowledgements and unique immutable keys distributed across the three initial shards.

Each container had a **256 MiB** memory limit, **1 CPU**, and temporary memory-backed data storage with Redis persistence disabled. The host had **12 logical CPUs**. The resolved image repository digest was:

```text
redis@sha256:3811787313eba226a2ef38658c6ccb91cd5e110edc89c37767de373120a0e5a0
```

Every individual report records the effective configuration, image identity, host/container metadata, topology source replies, and lifecycle timestamps. Their effective configuration SHA-256 was:

```text
8e959199c0f2dff505399e403a8b7f25bfe40c39ec4dbc9ab905198f7d8d36c3
```

This hash identifies the effective configuration recorded by the binary used for these trials. Later network and acknowledgement-policy support adds `experiment.isolation_duration`, `timeouts.restoration`, `timeouts.reconciliation`, `workload.write_policy`, `workload.wait_replicas`, and `workload.wait_timeout` defaults even when loading the same primary-crash YAML. A current reproduction therefore records those additional fields and a different configuration hash. The primary-crash fault and default SET acknowledgement policy remain the same, while newer write execution and routing are described in each new report's methodology; fresh timings are not measurements from this historical binary. The preserved reports and their original configuration hashes remain unchanged. Batch execution compares each newly produced report with its exact requested effective configuration; it does not import or rewrite historical reports as evidence from a new run.

## Individual outcomes

All five requested trials started and completed. Recovery success was **5/5**; partial, failed, aborted, and unstarted counts were all **0**. Cleanup was verified for every trial before the next trial began; independent discovery found no labelled containers or networks left over.

| Trial | Run ID | Outcome | Cleanup | Individual evidence |
|---:|---|---|---|---|
| 1 | `c71517bdf2d2fd24649843c9` | recovered | Verified | [JSON](examples/phase2-crash-batch/runs/c71517bdf2d2fd24649843c9/report.json), [Markdown](examples/phase2-crash-batch/runs/c71517bdf2d2fd24649843c9/report.md) |
| 2 | `447f02dcf25ca6d8a1931941` | recovered | Verified | [JSON](examples/phase2-crash-batch/runs/447f02dcf25ca6d8a1931941/report.json), [Markdown](examples/phase2-crash-batch/runs/447f02dcf25ca6d8a1931941/report.md) |
| 3 | `4052b45ef729eee33f8ba3f8` | recovered | Verified | [JSON](examples/phase2-crash-batch/runs/4052b45ef729eee33f8ba3f8/report.json), [Markdown](examples/phase2-crash-batch/runs/4052b45ef729eee33f8ba3f8/report.md) |
| 4 | `3a9679dd0e8a675a20d49be6` | recovered | Verified | [JSON](examples/phase2-crash-batch/runs/3a9679dd0e8a675a20d49be6/report.json), [Markdown](examples/phase2-crash-batch/runs/3a9679dd0e8a675a20d49be6/report.md) |
| 5 | `9c81a869a4e04108f3c73780` | recovered | Verified | [JSON](examples/phase2-crash-batch/runs/9c81a869a4e04108f3c73780/report.json), [Markdown](examples/phase2-crash-batch/runs/9c81a869a4e04108f3c73780/report.md) |

The killed primary remained stopped. Service and surviving topology recovered with reduced redundancy; process restart and restoration of the original replica count were not part of these trials.

## Observed recovery variability

The following values are seconds from **fault orchestration start**, including the ownership checks and Docker action. Every listed event was observed in all five trials. Timing statistics are conditional on those observations and include polling uncertainty.

| Observation | N | Min | Median | Mean | Sample SD | Max |
|---|---:|---:|---:|---:|---:|---:|
| Original replica promoted | 5 | 2.538939 | 2.558434 | 2.705453 | 0.339388 | 3.312375 |
| Full usable slot coverage | 5 | 2.538939 | 2.558434 | 2.705453 | 0.339388 | 3.312375 |
| First successful all-shard client probe | 5 | 3.058486 | 3.353813 | 3.243029 | 0.165055 | 3.369663 |
| Client probes sustain recovery | 5 | 4.172748 | 4.441869 | 4.347898 | 0.159720 | 4.482438 |
| Topology and clients jointly sustain recovery | 5 | 4.172748 | 4.454121 | 4.403807 | 0.130052 | 4.482438 |

Fault-action duration had mean **0.208304 s**, median **0.206691 s**, and range **0.199701–0.217500 s**. The action-completion anchor separates that measured orchestration interval from subsequent observation:

| Observation from action completion | N | Min | Median | Mean | Sample SD | Max |
|---|---:|---:|---:|---:|---:|---:|
| Original replica promoted | 5 | 2.333127 | 2.347920 | 2.497149 | 0.344163 | 3.112673 |
| Topology and clients jointly sustain recovery | 5 | 3.966936 | 4.254419 | 4.195503 | 0.129085 | 4.275747 |

The physical fault can begin before Docker reports action completion. Neither anchor is an exact Redis election timestamp. Observation work and request deadlines add to the configured sampling delay, and a successful probe does not prove availability between samples.

**N=5 is small.** P90/P95 across-trial timing estimates are intentionally omitted below 20 observed trials. The standard deviation uses the sample denominator N−1. No confidence bound, independent-host population claim, or production-tail estimate is inferred from this batch. Promotion and coverage shared sampled timestamps here; this does not establish that the underlying events were simultaneous.

## Client operations and correctness

| Phase | Completed requests | Successful | Terminal errors | Request-weighted success fraction |
|---|---:|---:|---:|---:|
| Baseline | 2,019 | 2,019 | 0 | 100% |
| Fault and recovery observation | 2,052 | 1,845 | 207 timeouts | 89.912281% |
| Post-recovery | 1,000 | 1,000 | 0 | 100% |

Fault-phase availability is **1,845/2,052**, weighting requests rather than averaging trial success percentages. The batch also retains the distribution of per-trial fractions. Successful and failed request latency populations remain separate. Across-trial distributions of each trial's histogram p50/p95/p99 are labelled accordingly; their means are **not pooled request percentiles**.

Whole-workload counters independently recomputed from the five individual `workload_stats` records were:

| Counter | Requests |
|---|---:|
| Offered schedule slots | 7,421 |
| Started client operations | 5,071 |
| Completed client operations | 5,071 |
| Dropped schedule slots | 2,350 |

They reconcile as **offered = started + dropped** and **started = completed** after quiescence. Busy workers and scheduling delays reduce starts without an unbounded queue or catch-up burst. Service probes and final validation are excluded. The preserved batch summary was produced before aggregate `load` fields were added; it therefore contains **no `load` field**. These totals come from individual raw reports, and the original summary has not been regenerated or edited to imply otherwise.

There were **2,564 attempted writes**, **2,459 acknowledged writes**, and **105 ambiguous SET outcomes**. All **2,459 acknowledged immutable records** were checked: **0 missing, 0 corrupt, 0 unreadable, and 0 unverified**. During traffic, no missing or mismatching acknowledged values were observed.

**No acknowledged-write loss was observed across the executed five trials.** This does not establish zero-loss Redis replication. A SET error may have executed and is not proof of loss; those 105 ambiguous outcomes remain outside the successful-acknowledgement ledger. Persistence was disabled, and this validation does not establish durable or linearizable writes.

## Failure cases and compatibility

The real `TestDockerRepeatedRecoveryTimeout` integration test passed with race detection in **37.902 s**. It executed two fresh trials with an intentionally insufficient **100 ms** recovery deadline: both returned failed outcomes, `timed_out: true`, and verified cleanup. Batch `batch-e3a8aef71917461d7de5322c` retained both trials, **0/2** recovery success, and two deadline-censored stable-recovery observations with **N=0**, without zero-duration samples.

The failed run IDs were `948d28b9474a42d4b514a332` and `184cb060cd5efa1e64511cac`. The original six-file timeout batch is preserved verbatim as [summary JSON](examples/repeated-timeout/summary.json), [summary Markdown](examples/repeated-timeout/summary.md), and both individual report pairs linked by that summary. These files are separate from the five-trial recovered example and retain the original recorded paths. A sample rejected after the deadline cannot establish accepted recovery. An absent accepted observation horizon is left absent; the configured deadline does not supply an invented sampled timestamp or prove eventual non-recovery.

The original single-run CLI was also executed again and returned exit code **0**, recovered outcome, and successful cleanup for run `6d68cf5c557f4594aae2b054`, recorded at **15:19:15–15:19:52 Singapore time**. This retains the Phase 1 entry point; its longer default configuration is not a direct comparison with the short repeated-trial definition.

## Automated validation and scope

The batch race tests and vetting passed:

```sh
go test -race ./internal/batch
go vet ./internal/batch
```

Coverage includes sequential execution, trial/outcome denominators, observed events in partial trials, deadline/abort censoring, unstarted trials, cleanup failures and contradictory cleanup flags, nil-report failures, duplicate IDs, missing/tampered report files, configuration identity, unsupported schemas, invalid validation accounting, zero-duration throughput, malformed recovered outcomes, statistical calculations, legacy timing fallback, and summary-write failure. CLI and experiment unit/compile checks also passed. Independent calculations from the individual reports agreed with outcome counts, descriptive statistics, request/data totals, unique IDs, effective configuration, lifecycle ordering, and pacing invariants. A subsequent audit rechecked all 12 copied files against their generated originals, verified the historical hash against its original configuration JSON, and resolved every relative report link in both this page and the copied batch summary.

After the network and policy accounting guards landed, the current source passed formatting checks, vet, race tests for **all nine packages**, and the CLI build:

```sh
make check build
```

A separate clean source copy, excluding `.git`, prebuilt `bin/`, and generated `reports/`, passed the same checks and build. `go mod verify` also returned **all modules verified** in that copy. These checks establish that the tested source builds without those generated artifacts; they do not claim a hosted CI execution. Later batch guard tests also verify summed phase policy counters and reply histograms, measured SET/WAIT stage attempts, and subset validation bounded by the full acknowledged-write ledger.

The five-trial evidence above covers **Milestone A**, primary process crashes, and normal SET acknowledgements. It supplies no WAIT-policy comparison evidence. Trials shared one host and may be correlated despite fresh containers. See [measurement methodology](measurement-methodology.md) and [Phase 1 validation](validation.md) for the measurement contract and earlier single-run evidence.

## Complete Docker integration suite

After adding network isolation and connection-affine SET/WAIT execution, the complete local Docker suite passed with race detection in **220.933 s**:

```sh
go test -race -tags=integration -count=1 -timeout=10m -v ./internal/experiment
```

The suite exercised real primary crashes, recovery deadlines, repeated deadline failures, cancellation, primary network isolation, restoration and reconciliation, interrupted network-fault teardown, and a WAIT-policy crash. It independently checked that managed containers and networks were absent after cleanup. The network and WAIT evidence below was produced by that suite. These reports are copied verbatim; content and SHA-256 hashes were compared with the originals. Their configurations and binary revision differ from the earlier Milestone A batch.

## Network isolation and failure paths

The fault disconnects the selected primary's sole Docker bridge endpoint in both directions while Redis continues running. It removes peer, replication, and published-client connectivity for that node. It does not simulate a selective partition or permit client writes to the isolated minority.

| Suite case | Run ID | Outcome | Connectivity restored | Reconciled | Final validation | Cleanup | Evidence |
|---|---|---|---|---|---|---|---|
| Isolation and reconciliation | `fc00ac9fc5a2c0b8b96bba88` | recovered | Yes | Yes | 1,547 checked | Verified | [JSON](examples/network-integration/fc00ac9fc5a2c0b8b96bba88/report.json), [Markdown](examples/network-integration/fc00ac9fc5a2c0b8b96bba88/report.md) |
| 100 ms reconciliation deadline | `011c149f7f247c9892b19156` | partial | Yes | No | 1,244 checked | Verified | [JSON](examples/network-failure-paths/011c149f7f247c9892b19156/report.json), [Markdown](examples/network-failure-paths/011c149f7f247c9892b19156/report.md) |
| Injected ambiguous disconnect response | `9a4541c05f49552d850e8a15` | failed | Yes | Unobserved | Skipped; 322 acknowledgements unchecked | Verified | [JSON](examples/network-failure-paths/9a4541c05f49552d850e8a15/report.json), [Markdown](examples/network-failure-paths/9a4541c05f49552d850e8a15/report.md) |
| Cancellation after disconnect | `e4544d8f12c536abaedb1905` | aborted | Yes | Unobserved | Skipped; 322 acknowledgements unchecked | Verified | [JSON](examples/network-failure-paths/e4544d8f12c536abaedb1905/report.json), [Markdown](examples/network-failure-paths/e4544d8f12c536abaedb1905/report.md) |

In the recovered case, an administrative probe observed the unchanged live Redis/container process while disconnected. Surviving topology and basic all-shard client probes sustained recovery before restoration. Reconnection preserved the original endpoints and process identity; the final six-node topology retained the promoted owner, demoted the original primary, and restored synchronized replication redundancy. All **1,547** acknowledged writes were checked, with **0 missing, corrupt, unreadable, or unverified** values. There were **57** ambiguous SET outcomes, outside that ledger. Promotion was sampled at **3.509503 s**, joint majority-side recovery at **5.983404 s** from orchestration start, and stable reconciliation at **2.145407 s** after restoration completed.

The deadline case also observed majority-side recovery during isolation and restored connectivity, but the **100 ms** reconciliation budget accepted no stable six-node recovery. `reconciliation_timed_out: true` remains separate from `recovery.timed_out: false`: the earlier majority-side recovery was observed. Complete checks of **1,244** acknowledged writes found no missing, corrupt, unreadable, or unverified values; data verification did not turn incomplete reconciliation into a recovered outcome.

The ambiguity test performs a real successful Docker disconnect, then injects an error response at the manager boundary. It demonstrates safe handling of an uncertain action result; it is not evidence of a naturally occurring Docker transport failure. The cancellation test cancels the experiment immediately after the real disconnect. Both use fresh bounded restoration contexts before independent cleanup, preserve the failure/abort outcome, and skip recovery and data-validation claims. Their zero default validation counters do not mean zero data loss was checked.

The earlier standalone CLI network run, `bbba656458a3cb89130c16de`, is preserved separately as [JSON](examples/primary-isolation/report.json) and [Markdown](examples/primary-isolation/report.md). It verified **1,389** acknowledged writes with no observed loss and successful cleanup. It is a different run from the suite's **1,547**-write case.

## WAIT-policy crash case

Run `470c45df26b8848ae11e929a` is preserved as [JSON](examples/wait-policy-crash/report.json) and [Markdown](examples/wait-policy-crash/report.md). It used **SET followed by WAIT for one replica on the same physical connection**, a **100 ms** WAIT timeout, **500 ms** overall operation timeout, **4 s** baseline, **3 s** post-recovery period, and **1 s** stability interval. Ordinary topology and all-shard SET/GET service recovered, while the promoted shard had no replica and the all-shard acknowledgement policy remained unavailable. The accurate overall outcome was **partial**, with verified cleanup.

| Write or validation population | Count |
|---|---:|
| SET attempts | 1,007 |
| Successful SET acknowledgements | 959 |
| Ambiguous SET errors | 48 |
| WAIT replies satisfying one replica | 810 |
| WAIT replies returning zero replicas | 149 |
| WAIT command errors / client timeouts | 0 / 0 |
| All acknowledged SETs checked | 959 |
| WAIT-satisfied subset checked | 810 |

All **959** acknowledged SETs remained in the ledger, including the **149** writes whose WAIT was short. Validation of the full ledger and the **810**-write subset found **0 missing, corrupt, unreadable, or unverified** values. During traffic, missing and mismatched acknowledged-value checks were also zero. This observation supplies neither a loss-free WAIT guarantee nor a disk-durability claim.

Logical write metrics reconcile as **1,007 SET+WAIT operations = 810 policy successes + 149 short replies + 48 SET errors**. Diagnostic stage counts were **1,007 SET attempts + 959 WAIT attempts = 1,966**, separate from **2,003** completed logical workload requests, including **996** GETs. Schedule accounting was **2,312 offered = 2,003 started + 309 dropped**. A zero-replica integer reply is a measured short WAIT result, distinct from a client timeout or an unattempted WAIT.

Before the fault, all **426** successful baseline SETs satisfied WAIT. After basic service recovery, **103/308** acknowledged post-recovery SETs returned zero replicas. The policy-stability milestone remains absent; it is not a zero-time recovery or a claim that WAIT was tested until the full configured recovery deadline after ordinary service had already recovered. This single integration case establishes behavior and accounting, not a latency comparison with the historical five-run batch.

## Matched SET and WAIT batches

Milestone C then executed **three fresh SET trials and three fresh WAIT trials** with matching effective configurations except for the experiment name and write policy. Both used the same connection-affine write executor, **8 s** baseline, **3 s** post-recovery observation, **1 s** stability interval, and **200 offered logical operations/s**. WAIT requested one replica with a **100 ms** server timeout inside the **500 ms** operation deadline. The complete comparison, measurement limits, and individual run links are in [SET/WAIT consistency comparison](consistency-comparison.md).

| Batch | Trial outcomes | Acknowledged SETs checked | Ambiguous SET outcomes | Evidence |
|---|---|---:|---:|---|
| SET `batch-bbc2fe8f6b545c67219b3727` | 3 recovered | 4,015 | 132 | [Summary JSON](examples/policy-set-batch/summary.json), [summary Markdown](examples/policy-set-batch/summary.md) |
| WAIT `batch-26adf7924f9e85799ff28453` | 3 partial | 4,070 | 124 | [Summary JSON](examples/policy-wait-batch/summary.json), [summary Markdown](examples/policy-wait-batch/summary.md) |

All six trials observed basic topology and SET/GET service recovery, completed validation, and verified cleanup. The WAIT trials accurately remained **partial**: basic service recovered, while the all-shard replica acknowledgement policy did not. Of their **4,070** successful SET acknowledgements, **3,588** satisfied WAIT and **482** returned fewer than the requested replica acknowledgements. All **482** short-WAIT writes remained in the full ledger. Validation checked both the **4,070**-write ledger and its **3,588**-write WAIT-satisfied subset. Both batches had **0 missing, corrupt, unreadable, or unverified** acknowledged values; ambiguous SET outcomes were excluded from those checks.

An independent audit checked the six unique run identities, **36 fresh container identities and six fresh networks**, matching configurations and configuration hashes, event statistics and phase aggregates against individual reports, cleanup and validation subsets, and pacing, stage, policy, and ledger accounting. Each batch contains three complete individual report pairs with no unstarted trial, cleanup failure, or incomplete evidence. All **16 copied batch files** were checked byte-for-byte and by SHA-256 against their generated originals. **N=3 per policy is small**: the comparison suppresses across-trial P90/P95, describes distributions of individual request percentiles rather than pooled request percentiles, and makes no production-tail or loss-free replication claim. SET ran before WAIT on the same host; the sequence was not randomized.

## Final compatibility and source checks

After the policy changes and matched batches, the original single-run command completed with exit code **0**:

```sh
./bin/redis-lab run --output reports/phase2-final experiments/primary-failure.yaml
```

Run `1f071109915a0bce646c07f5`, recorded at **15:48:01–15:48:39 Singapore time** on 11 October 2026, recovered and verified cleanup. It checked all **2,046** acknowledged SETs with **0 missing, corrupt, unreadable, or unverified** values; **121** ambiguous SET outcomes remained outside the ledger. Promotion was sampled at **5.339418 s** and joint stable service recovery at **8.537254 s** from fault orchestration start. Its [JSON](examples/final-primary-failure/report.json) and [Markdown](examples/final-primary-failure/report.md) are preserved verbatim, with both SHA-256 hashes checked against the originals.

The latest exact Go source in a separate clean copy passed **`make check build`**, including formatting, vet, all nine packages' race tests, and the CLI build; **`go mod verify`** again reported all modules verified. All six experiment YAML definitions passed validation. Final independent labelled Docker container and network inventories were empty. These results preserve the original CLI entry point alongside the completed repeated-run, network-isolation, and acknowledgement-policy evidence.

All evidence in this page was produced on one **macOS arm64 Docker Desktop host**. Hosted GitHub Actions and a separate Linux host were not executed for this validation. Fresh clusters isolate resource lifecycles, while shared host load, election timing, and asynchronous replication remain sources of variability.
