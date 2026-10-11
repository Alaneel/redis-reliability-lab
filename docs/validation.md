# Phase 1 validation evidence

This page preserves the original Phase 1 evidence. Later expanded-suite results are recorded in [Phase 2 validation](phase2-validation.md) and [Phase 3 validation](phase3-validation.md).

Checks were executed on **11 October 2026** using **Go 1.26.9**, a **macOS arm64** host, and **Docker Desktop with Docker Engine 24.0.6 Linux containers**. These are observed results from this environment, not expected values or benchmark guarantees. Hosted GitHub Actions and a separate Linux host were not executed locally.

## Build and deterministic checks

The following checks passed:

```sh
go test -race -count=1 ./...
go vet ./...
go mod verify
gofmt -l cmd internal
go build -trimpath -buildvcs=false -o bin/redis-lab ./cmd/redis-lab
```

Formatting returned no filenames. Configuration, cluster parsing/readiness/recovery, workload and metrics, state transitions, report generation, client target restrictions, Docker safety, cancellation, partial provisioning, and cleanup paths have automated coverage. Environment tests run against a deterministic Docker fake; real Redis/container checks are separate.

The final combined formatting/vetting/uncached race-test/build check passed. A separate temporary copy of the source also built successfully with `-mod=readonly -trimpath -buildvcs=false`, confirming the build does not depend on generated experiment reports or local output files.

## Real Docker integration

The final full suite passed with race detection:

```sh
go test -race -tags=integration -count=1 -timeout=10m -v ./internal/experiment
```

The final passing suite reported **56.975 seconds** and covered automatic primary failover, an intentionally insufficient **100 ms** recovery deadline, and cancellation immediately after the managed primary kill. Every run checked cleanup and independently verified that no labelled run containers or network remained. The failover and timeout cases also decoded their saved JSON reports and checked their Markdown files.

| Case | Run ID | Observed outcome | Cleanup verified |
|---|---|---|---|
| Automatic primary failover | `0e8142da5c18f79f0c8225c4` | `recovered` | Yes |
| 100 ms recovery deadline | `8120ec03a2cc7d3edf6e87e4` | `failed`, `timed_out: true` | Yes |
| Cancellation after kill | `172e17d70d20bdefd7dd1a1c` | `aborted` | Yes |

The timeout case retained one sample marked `excluded_from_recovery: true`. Deadline expiry means recovery was not accepted within the observation budget; it does not prove Redis could never recover given more time.

The automatic-failover case validates the full topology/service contract and acknowledged-write checks. Because replication is asynchronous, real missing acknowledged writes are allowed as an accurately reported **partial** experiment outcome; they must never be called recovered. The suite passing is evidence that the framework evaluated the observed result correctly, not a promise that every failover preserves every write.

The final failover case observed original-replica promotion at **2.740397 s**, full usable coverage and topology recovery at **3.260984 s**, the first successful all-shard client probe at **3.047390 s**, and combined stable recovery at **4.410092 s**. All **626** acknowledged writes were verified with no missing, corrupt, unreadable, or unverified values. Client probe success preceding complete topology agreement illustrates why these conditions are recorded separately. The full CLI example uses a longer baseline/stability configuration and different node timeout, so this shorter integration run is not a direct performance comparison.

All milestone times begin at the start of validated fault injection, including Docker ownership checks before `SIGKILL`. They are first sampled observations; actual election and state transitions can happen between samples.

## Canonical default CLI run

The final default experiment completed with exit code **0** and outcome **recovered**:

```sh
./bin/redis-lab run -output reports experiments/primary-failure.yaml
```

Run ID: `56dfe45e38604c041b1c6087`. It ran from **14:52:12 to 14:52:50 Singapore time** on 11 October 2026. Exact generated evidence is checked in as the [Markdown report](examples/primary-failure.md) and [JSON report](examples/primary-failure.json).

| Observed milestone | Seconds from validated fault start |
|---|---:|
| Failure flagged in gossip | 4.333341 |
| Original replica promoted | 6.399665 |
| Full usable slot coverage | 6.399665 |
| Surviving topology recovered | 6.399665 |
| First successful all-shard client probe | 6.425577 |
| Topology and clients jointly stable | 9.577199 |

Fault ownership inspections and command completion took about **0.217 s**, included in that time origin. The default Redis node timeout was **3 s** and required stability span **3 s**.

| Workload phase | Completed attempts | Success rate | Terminal errors | Successful p95 / p99, ms |
|---|---:|---:|---:|---:|
| Baseline | 2,004 | 100% | 0 | 0.731856 / 0.847216 |
| Fault/recovery observation | 947 | 90.3907% | 91 timeouts | 1.597558 / 172.835069 |
| Post-recovery | 999 | 100% | 0 | 0.663814 / 1.677436 |

Successful latency excludes failed requests. During the fault phase, all-attempt p95 and p99 were both about **501.103 ms**, reflecting the 500 ms operation deadline and scheduling overhead. Values are histogram estimates; raw distributions and phase durations are in JSON.

There were **2,004 attempted writes**, **1,960 acknowledged writes**, and **44 ambiguous write outcomes**. All **1,960 acknowledged values** were checked: **0 missing, 0 corrupt, 0 unreadable, 0 unverified**. Traffic performed **1,946 read checks**, with no missing or mismatching acknowledged values observed. These results describe this run; they do not prove Redis replication is lossless.

Cleanup verified removal of all **six containers and the run network**. The killed process was never restarted, and replication redundancy was not restored; both report fields correctly remain false.

The observed runtime was Redis **8.10.2**, Go **1.26.9**, Darwin **arm64**, Docker Engine **24.0.6**, and Docker Linux kernel **6.4.16-linuxkit**. The image repository digest was:

```text
redis@sha256:3811787313eba226a2ef38658c6ccb91cd5e110edc89c37767de373120a0e5a0
```

The report also records the effective configuration hash, resolved image ID, client refresh/retry policies, topology source replies, and lifecycle timestamps.

## Runtime findings incorporated into the implementation

- The tested Docker daemon did not expose active published ports for containers on an `--internal` bridge. The final provisioner uses a per-run bridge with egress IP masquerading disabled and loopback-only publishing. Cleanup also recovers legacy managed internal bridges and partially created containers.
- An early run recovered surviving topology in about **3.18 s** while client operations remained unsuccessful before a **20 s** deadline. go-redis v9.23.0's default **60 s** lazy refresh interval left cached dead-primary ownership too long for the intended observation window. The final client explicitly records and uses a **1 s** refresh interval; see [ADR 0004](decisions/0004-explicit-client-refresh.md).
- A stale gossip view could temporarily assign all slots to the known dead primary before peers marked it failed. The final full-coverage milestone requires each reported owner to be directly reachable and excludes the known killed primary. A regression test covers this case.
- A recovery sample that finishes after its deadline is retained as raw evidence with `excluded_from_recovery: true`, but cannot establish a milestone or accepted recovery. This avoids treating deadline-crossing observations as timely success.
- A successful Docker subprocess can emit a diagnostic warning on stderr while returning its resource ID or JSON on stdout. The final executor keeps these streams separate, logs the warning, and parses only stdout. Real stub-subprocess tests exercise warning-bearing IDs/JSON and retained error diagnostics.

These unsuccessful intermediate runs led to concrete behavior and measurement fixes. They are not presented as passing evidence.

## Validation limits

The real runtime evidence covers one macOS arm64 Docker Desktop environment. Linux CI is configured, but no hosted workflow result is claimed. Windows and remote Docker daemons are unsupported. Runs share one host and use temporary memory with persistence disabled; they do not validate independent host faults, network partitions, disk durability, or production capacity.

A run with no observed missing writes cannot establish a zero-loss guarantee for asynchronous Redis replication. Recovery probe success does not imply every unsampled client operation succeeded. Reproduce results with the full effective configuration, image digest, client policy, and environment metadata, and use repeated runs to evaluate variability.
