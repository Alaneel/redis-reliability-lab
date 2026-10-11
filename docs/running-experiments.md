# Running experiments

Run commands from the repository root. The lab accepts local experiment files and managed run IDs; it has no arbitrary Redis address, remote Docker target, or arbitrary fault-command option. The [quick start](../README.md#quick-start) introduces the original primary crash; this guide covers the remaining CLI surface.

## Prerequisites and build

Use Go 1.26 or newer; `go.mod` selects toolchain `go1.26.9`, which Go can download automatically. Docker Engine or Docker Desktop must run Linux containers, and the Docker CLI must reach a local Unix socket. macOS and Linux fit this host model; native Windows named pipes and TCP/SSH endpoints are rejected. Prefer Docker Engine 28+ on native Linux because older engines have a localhost-publication caveat; see the [Security policy](../SECURITY.md). The retained Docker 24.0.6 observations used macOS Docker Desktop.

Allow six containers with limits of 256 MiB RAM and one CPU each, plus Docker and the Go process. Initial Go dependencies need internet access. Every provisioning run explicitly pulls the configured exact-version Redis image, even when cached, so registry access is required. Retained-data experiments use disposable owned volumes without a hard filesystem quota.

```sh
docker info
make build
bin/redis-lab validate experiments/primary-failure.yaml
bin/redis-lab version
```

`validate` checks the file without provisioning or contacting Docker. A source build prints the current application version (`dev` unless supplied at build time). Docker endpoint resolution honors `DOCKER_HOST` or the selected Docker context only when the endpoint is `unix:///absolute/socket`; an active run then keeps that socket fixed.

## Single runs and available definitions

```sh
bin/redis-lab run -output reports experiments/primary-failure.yaml
```

`run` flags precede its one experiment file. Each run prints its outcome and report directory, and saves `reports/<run-id>/report.json` and `report.md`. Start with Markdown; JSON retains effective configuration and its hash, Docker/image metadata, source topology replies, sampled observations, metrics, correctness results, and cleanup status. Initial image downloads can exceed the experiment duration; increase `timeouts.provision` in the YAML if needed.

| Experiment definition | Fault and intended observation |
|---|---|
| [primary-failure.yaml](../experiments/primary-failure.yaml) | Original primary SIGKILL; service recovery with the killed process still stopped |
| [primary-failure-repeat.yaml](../experiments/primary-failure-repeat.yaml) | Shorter primary-crash definition for repeated trials; settings differ from the original |
| [primary-failure-timeout.yaml](../experiments/primary-failure-timeout.yaml) | Intentionally insufficient 100 ms recovery deadline; expected nonzero exit with retained evidence |
| [primary-isolation.yaml](../experiments/primary-isolation.yaml) | Whole-node bidirectional primary isolation, restoration, and reconciliation |
| [replica-crash.yaml](../experiments/replica-crash.yaml) | Replica SIGKILL, same-container restart, and restored redundancy |
| [replica-isolation.yaml](../experiments/replica-isolation.yaml) | Replica isolation, exact reconnection, and restored redundancy |
| [primary-repair.yaml](../experiments/primary-repair.yaml) | SET primary crash followed by former-primary repair |
| [primary-repair-wait.yaml](../experiments/primary-repair-wait.yaml) | Former-primary repair plus independently observed WAIT availability |
| [policy-set.yaml](../experiments/policy-set.yaml), [policy-wait.yaml](../experiments/policy-wait.yaml) | Matched memory-only SET/WAIT primary crashes without repair |
| [persistence-memory.yaml](../experiments/persistence-memory.yaml) | All-process restart with a frozen volatile ledger; expected partial data-loss control |
| [persistence-aof-everysec.yaml](../experiments/persistence-aof-everysec.yaml), [persistence-aof-always.yaml](../experiments/persistence-aof-always.yaml) | All-process restart on retained AOF volumes |
| [policy-aof-set.yaml](../experiments/policy-aof-set.yaml), [policy-aof-wait.yaml](../experiments/policy-aof-wait.yaml), [policy-aof-waitaof.yaml](../experiments/policy-aof-waitaof.yaml) | Matched AOF-always primary-repair conditions differing only in name and write policy |

Replica faults always restore their target. Optional primary crash repair is enabled by `experiment.repair: true`; it waits for failover service, restarts the same former-primary container with its original Cluster ID, and requires it to rejoin as a synchronized replica of the promoted owner. [Replica lifecycle](replica-lifecycle.md).

```sh
bin/redis-lab run -output reports experiments/primary-isolation.yaml
bin/redis-lab run -output reports experiments/replica-crash.yaml
bin/redis-lab run -output reports experiments/replica-isolation.yaml
bin/redis-lab run -output reports experiments/primary-repair-wait.yaml
```

Whole-node isolation removes peer and host-client access in both directions. The lab observes process survival administratively, restores the original endpoints, and verifies identities, assignments, synchronized replicas, and stable client service. [Network semantics](network-partition.md).

## Repeated trials

```sh
bin/redis-lab repeat experiments/primary-failure-repeat.yaml --runs 5 --output reports
bin/redis-lab repeat experiments/policy-set.yaml --runs 3
bin/redis-lab repeat experiments/policy-wait.yaml --runs 3
```

`repeat` accepts flags before or after its one experiment file; `--runs` supports 1–100 trials and defaults to five. It writes `reports/batch-<id>/summary.json`, `summary.md`, and `runs/<run-id>/report.json`/`report.md`. Trials run sequentially in fresh clusters, with cleanup completed before the next provisioning. Failed and partial outcomes remain in the summary. Cancellation or uncertain cleanup stops further provisioning; unstarted trials are counted. Any unsuccessful trial makes the batch exit `1` while preserving reports.

The memory-only WAIT crash condition is expected to become partial: ordinary service can recover while the promoted shard has no replica to acknowledge writes. Every successful SET remains in data validation even when WAIT is short or fails. Compare healthy `baseline` measurements separately from `fault` and `post_recovery`. [Consistency comparison](consistency-comparison.md).

Recovery statistics are conditional on observed milestones, with unobserved/censored and not-at-risk counts retained. With fewer than 20 observations, batch p90/p95 recovery estimates are omitted. Aggregated request latency values describe per-trial quantiles rather than pooled request percentiles. [Measurement methodology](measurement-methodology.md).

## Persistence controls and counterbalanced comparisons

```sh
bin/redis-lab run experiments/persistence-memory.yaml
bin/redis-lab run experiments/persistence-aof-everysec.yaml
bin/redis-lab run experiments/persistence-aof-always.yaml
```

All-process restart drains traffic and freezes the acknowledged ledger, kills original replicas before primaries, verifies all six stopped, and restarts the same containers and managed volumes. Validation reads the frozen ledger before diagnostic writes. Require a nonempty ledger when reproducing the memory control; its expected data loss produces a partial outcome and exit `1`. No workload traffic is offered during fault/restoration, so later diagnostic service timing has no workload availability denominator. [Durability semantics](durability-semantics.md).

Use recorded order rotation for multiple conditions:

```sh
bin/redis-lab compare --rounds 3 --seed 20261011 --output reports \
  experiments/persistence-memory.yaml \
  experiments/persistence-aof-everysec.yaml \
  experiments/persistence-aof-always.yaml

bin/redis-lab compare --rounds 3 --seed 20261011 --output reports \
  experiments/policy-aof-set.yaml \
  experiments/policy-aof-wait.yaml \
  experiments/policy-aof-waitaof.yaml
```

`compare` flags precede its 2–8 files, with at most 100 total trials. Rounds default to three and the order seed defaults to `20261011`. It saves the full schedule, individual reports, and per-condition summaries under `reports/comparison-<id>/`. The order seed controls the schedule; workload seeds remain unchanged.

A seeded base permutation rotates each round. With N conditions, each condition occupies each position once in N rounds; incomplete blocks are only partly balanced. This does not balance every predecessor pair or establish statistical significance. Expected partial memory trials return `1` while retaining evidence. The retained [persistence comparison](examples/phase3/persistence-comparison/comparison.md) and [policy comparison](examples/phase3/policy-comparison/comparison.md) each contain nine actual fresh trials.

WAIT requests one replica acknowledgement for preceding SETs on the same physical connection. WAITAOF requests one local and one replica AOF acknowledgement on that connection, requires AOF on all managed nodes and Redis 7.2+, and retains the observed local/replica pair independently. Short replies and command errors remain distinct from SET acknowledgements. Neither policy establishes universal loss-free failover or power-loss durability. [Policy accounting](consistency-comparison.md) · [WAITAOF and persistence](durability-semantics.md).

## Configuration defaults

Copy a definition from `experiments/` and validate it before running. Unknown fields, multiple YAML documents, unsupported faults, unpinned images, and out-of-range values are rejected. Settings omitted from YAML receive these defaults:

| Setting | Default | Meaning |
|---|---|---|
| `environment.redis_image` | `redis:8.10.2-alpine` | Exact official version tag; optional `@sha256:...` digest |
| `environment.node_timeout` | `3s` | Redis Cluster failure-detection configuration |
| `environment.persistence` | `memory` | `memory`, `aof-everysec`, or `aof-always` |
| `workload.workers` | `8` | Maximum concurrent workload operations |
| `workload.rate_per_second` | `200` | Aggregate request-start ceiling; occupied workers cause dropped offers |
| `workload.read_percent` | `50` | Seeded read/write choice; reads check acknowledged keys |
| `workload.max_writes` | `20000` | Bound on distinct write attempts and the validation ledger |
| `workload.value_bytes` | `128` | Synthetic value size |
| `workload.seed` | `42` | Worker-local random-choice seed |
| `workload.operation_timeout` | `500ms` | Deadline for a whole logical operation |
| `workload.write_policy` | `set` | `set`, connection-affine `wait`, or AOF-enabled `waitaof` |
| `workload.wait_replicas` | `1` | Fixed one-replica acknowledgement request; WAITAOF also requests one local AOF acknowledgement |
| `workload.wait_timeout` | `100ms` | Server timeout in positive integral milliseconds; operation deadline also applies |
| `experiment.fault` | `primary-kill` | `primary-kill`, `primary-isolate`, `replica-kill`, `replica-isolate`, or `cluster-restart` |
| `experiment.repair` | `false` | Optional primary crash repair; replica faults always restore their target |
| `experiment.down_duration` | `3s` | Minimum replica crash hold or all-stopped restart hold |
| `experiment.isolation_duration` | `8s` | Minimum network hold after successful fault action, bounded by recovery timeout |
| `experiment.baseline` | `10s` | Healthy traffic interval before injection |
| `experiment.minimum_baseline_success` | `0.99` | Required baseline success fraction |
| `experiment.sample_interval` | `500ms` | Delay after completed observation and probes |
| `experiment.stable_for` | `3s` | Required span of consecutive positive recovery samples |
| `experiment.post_recovery` | `5s` | Continued observations and traffic after recovery |
| `timeouts.provision` | `2m` | Provisioning budget, including image pull |
| `timeouts.readiness` | `1m` | Initial topology and replica-readiness budget |
| `timeouts.recovery` | `45s` | Recovery-observation budget |
| `timeouts.repair` | `30s` | Combined restart/reconnect and synchronization budget, including all six restart nodes |
| `timeouts.restoration` | `30s` | Network-restoration budget, including a fresh attempt after interruption |
| `timeouts.reconciliation` | `30s` | Restored six-node topology and client-stability budget |
| `timeouts.validation` | `30s` | Acknowledged-write readback budget |
| `timeouts.cleanup` | `30s` | Automatic cleanup budget |

The topology is fixed at three primaries and one replica per primary. Traffic uses one hash tag per original shard rather than uniformly covering every slot. Values derive from recorded run ID, seed, and sequence; OS scheduling, replication, and elections remain nondeterministic.

Retained-data experiments bound `max_writes` to 20,000 and value size to 256 bytes, and restrict recovery, repair, and reconciliation budgets to two minutes. WAIT/WAITAOF require an operation deadline at least 10 ms beyond their server timeout. The complete bounds and compatibility rules live in [configuration validation](../internal/config/config.go).

## Interpreting reports and exit status

| Outcome | Interpretation |
|---|---|
| `recovered` | All applicable topology, client, requested repair/policy/persistence, data, and cleanup criteria passed |
| `partial` | Some recovery was observed, but complete recovery or data criteria were not met |
| `failed` | Setup, fault execution, recovery, or cleanup failed; inspect detailed errors |
| `aborted` | Cancellation interrupted the run; bounded cleanup was still attempted |

Exit code `0` means command success; an incomplete/unsuccessful run or report-writing error returns `1`; invalid command usage or configuration returns `2`. `validate` and `cleanup` return `0` on success. A deadline expiration can be a valid experimental result even though the command exits nonzero:

```sh
bin/redis-lab run -output reports experiments/primary-failure-timeout.yaml
```

Inspect independent topology, client service, data, process restart, redundancy, policy, and persistence fields. A default primary crash leaves its process stopped and redundancy reduced. WAIT/WAITAOF need sustained all-shard policy probes for full recovery; basic SET/GET recovery alone can be partial. Persistence additionally needs the all-stopped boundary, fresh runtime/file evidence, and preserved frozen values. Schema 3 retains these dimensions independently; absent evidence means unobserved. Older raw examples preserve their original schemas.

Only successful SET acknowledgements enter the validation ledger, including subsequent short or failed WAIT/WAITAOF calls. Error-returning SETs may have executed and remain ambiguous. Unique immutable keys prevent later writes from hiding loss. A skipped or incomplete validation with zero counters proves no preservation.

The cluster client uses RESP2, one-second lazy topology refresh, an eight-redirect limit, disabled per-node retries, one attempt per pool dial, and 16 connections per node pool. Write policies use separate owned pools and the same physical-connection routing path. Only explicit allowlisted MOVED/ASK SET rejections permit bounded replay; transport and acknowledgement-command errors are never replayed. Environment metadata records the settings.

Latencies include routing and handled redirects, with successful, failed, and all-attempt distributions separate. Logical SET+WAIT/WAITAOF completion is separate from command-stage diagnostics; stages do not increase offered or logical request counts. Percentiles are bounded histogram estimates with 5% bucket growth; inspect sample counts. Throughput divides completed attempts by phase wall time. Offered, started, completed, and dropped counters expose worker saturation. Handled write redirects and terminal redirect errors are recorded separately.

Milestones are first sampled observations relative to validated fault orchestration start, which includes ownership inspection. Fault completion and action duration are separate anchors. Actual sample spacing includes observation, probes, and configured delay; samples crossing a decision deadline cannot establish recovery. At most 1,200 topology samples are retained, with omitted samples counted. Probes establish sampled availability, while workload metrics describe other requests. [Measurement methodology](measurement-methodology.md) · [Architecture](architecture.md).

## Cleanup and interruption

Containers configure client-port publication on `127.0.0.1`, have no host bind mounts, use read-only root filesystems, and drop capabilities. Loopback publication is not an independent network-security guarantee; Docker version and host networking matter as described in the [Security policy](../SECURITY.md). Ordinary memory primary crash/isolation runs use `/data` tmpfs; retained-identity and AOF runs use disposable owned volumes. Generated traffic is bounded; volume file systems have no hard quota.

Every managed kill, restart, disconnect, reconnect, and removal verifies resource IDs, labels, expected names, and isolation settings. Crash repair starts the same stopped container and verifies retained Cluster ID and original endpoints. Reconnection refuses changed processes or occupied endpoints. Cleanup removes tracked containers, their empty network, and tracked volumes, verifies absence, and is idempotent. Unrelated resources and downloaded images remain.

`Ctrl+C` and `SIGTERM` cancel traffic and trigger cleanup with its own bounded context. A hard kill, machine shutdown, or unavailable daemon can leave resources. Replace `RUN_ID` with the exact run ID from logs or reports:

```sh
bin/redis-lab cleanup RUN_ID
```

The recovery command has a one-minute deadline and reconstructs resources from exact management labels and expected names. It can remove a validated disconnected node without first reconnecting it. If ownership or isolation checks fail, inspect the reported condition; altered resources are preserved. [Security policy](../SECURITY.md) · [Cleanup design](architecture.md#failure-handling-and-cleanup-guarantees).

For local development checks and the opt-in Docker suite, see [Contributing](../CONTRIBUTING.md). For executed results and original raw reports, see the [documentation index](README.md#validation-and-raw-evidence).
