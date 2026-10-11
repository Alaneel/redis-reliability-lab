# Replica lifecycle and redundancy repair

Milestone A extends the existing service-failover experiment with replica crash/restart, replica isolation/reconnect, and optional repair of a failed primary. Three real CLI examples below reached verified six-node redundancy and completed acknowledged-write validation and cleanup. These examples establish the implemented lifecycle behavior; they do not establish persistence recovery or complete all of Phase 3.

## Run the experiments

Use the same build and local Docker prerequisites as the [quick start](../README.md#quick-start):

```sh
./bin/redis-lab run -output reports experiments/replica-crash.yaml
./bin/redis-lab run -output reports experiments/replica-isolation.yaml
./bin/redis-lab run -output reports experiments/primary-repair-wait.yaml
```

The [crash definition](../experiments/replica-crash.yaml) uses `replica-kill`; the [isolation definition](../experiments/replica-isolation.yaml) uses `replica-isolate`. Replica faults always repair their target. `down_duration` sets the minimum replica hold after successful fault action completion; the recovery observation budget still bounds that hold. The [primary repair definition](../experiments/primary-repair.yaml) adds `experiment.repair: true` to `primary-kill`. The WAIT variant exercises acknowledgement-policy recovery as a separate criterion.

Selection remains deterministic: select the original primary owning slot zero and retain its replica and slot set. A replica experiment faults that replica; a primary experiment faults the original primary. Managed node names, Redis Cluster IDs, and network endpoints are separate identities in the report.

## Service and redundancy have separate criteria

During a replica fault, service topology requires the other five managed nodes to remain directly reachable, all three original primaries to retain their original slot ownership, and unaffected replica links to remain ready. The failed replica can be absent or unhealthy in gossip. No promotion is required, and continued service does not require waiting for a `FAIL` flag. Fault receipts and process evidence establish that the action occurred independently of this service predicate.

After replica restart or reconnect, acceptance requires all six original managed and Redis identities and endpoints, unchanged owners and replica assignments, equal full slot maps, three primaries with one replica each, and replication link-up with initial synchronization complete. All-shard SET/GET probes must overlap this condition for a fresh `stable_for` interval. A completed repair action or a running process alone cannot establish restored redundancy.

Primary repair begins after five-survivor failover service has stabilized. The original replica must keep every failed-primary slot. The restarted former primary must be observed as its replica, with the expected primary endpoint and synchronization readiness. The runner inspects Redis's resulting topology; it does not issue `CLUSTER REPLICATE`, transfer slots, or force an old primary to recover ownership. Unaffected shards must remain unchanged.

The `repair` deadline covers the action and subsequent synchronization. Stability tracking resets when repair starts; a pre-repair success interval cannot satisfy final restoration. Historical service milestones remain visible if repair later fails. Current final conditions, complete data validation, applicable write-policy recovery, and verified cleanup determine the outcome.

Schema 3 adds `lifecycle_fault`, per-sample dimensions, and final `recovery_dimensions`. Independent flags describe process liveness, membership, usable slots, leaders, reads, writes, data, write-policy availability, redundancy, and persistence. An absent pointer means unobserved. For example, ordinary memory-only replication readiness leaves persistence recovery unobserved.

## Real observations

These were sequential individual CLI runs on 2026-10-11, using Redis 8.10.2, Go 1.26.9, macOS arm64, and Docker Desktop with Docker Engine 24.0.6. Each used eight workers, a 200 requests/s offered limit, 50% reads, seed 42, 128-byte values, a 500 ms operation deadline, 1 s Redis node timeout, 2 s baseline, 250 ms inter-sample delay, 1 s required stability, and 1 s post-recovery traffic. Replica holds were 3 s; recovery and repair budgets were 30 s. Effective configurations, resolved image digest, and platform metadata are retained in each JSON report.

| Observation | A1: replica crash/restart | A2: replica isolation/reconnect | A3: primary crash/repair with WAIT |
|---|---:|---:|---:|
| Outcome | recovered | recovered | recovered |
| First failure flag, since fault orchestration | 1.718 s | 1.896 s | 1.817 s |
| Stable ordinary service, since fault orchestration | 1.447 s | 3.090 s | 4.665 s |
| Repair action duration | 1.520 s | 1.240 s | 1.554 s |
| Stable six-node redundancy, after repair completion | 1.239 s | 1.570 s | 5.914 s |
| Completed workload requests | 1,858 | 2,125 | 2,721 |
| Workload requests with terminal errors | 0 | 0 | 344 |
| Acknowledged immutable SETs verified | 936 | 1,069 | 1,325 |
| Ambiguous SET outcomes, outside the ledger | 0 | 0 | 55 |
| Missing / corrupt / unreadable / unverified acknowledged values | 0 / 0 / 0 / 0 | 0 / 0 / 0 / 0 | 0 / 0 / 0 / 0 |
| Cleanup verified | yes | yes | yes |

Read the verbatim [A1 Markdown](examples/phase3/replica-crash/report.md) and [JSON](examples/phase3/replica-crash/report.json), [A2 Markdown](examples/phase3/replica-isolation/report.md) and [JSON](examples/phase3/replica-isolation/report.json), and [A3 Markdown](examples/phase3/primary-repair-wait/report.md) and [JSON](examples/phase3/primary-repair-wait/report.json). The [Milestone A validation record](phase3-validation.md) records executed checks and detailed observations. These are one execution per scenario, with different fault and acknowledgement policies; the table is descriptive and makes no comparative performance or recovery-distribution claim. Quantiles in individual reports are bounded histogram estimates, not exact request quantiles.

A1 kept all workload requests successful while the original primary remained available. Its primary's accepted-partial-sync counter increased by one; the full-sync and rejected-partial counters did not increase. Gossip failure detection occurred after sampled service stability, which is consistent with independently measuring service and failure detection.

A2 proved whole-node bidirectional interface disconnection: administrative probes recorded Redis PONG, a running container, zero attached networks, and unchanged container PID and Redis process run ID while direct clients could not reach the target. The replica offset remained at 21,009 across the two detached probes, while last-I/O age increased from 1 to 4 seconds. Those probes still reported `master_link_status:up`; the primary continued reporting one connected replica, and none of its full/partial counters increased over the comparison interval. The saved report therefore leaves replication-interruption and resync-event milestones unobserved. Physical disconnection is established by the owned Docker receipt and process probes; these samples cannot establish a Redis link-down transition, classify a resync event, or show that every relevant transition was captured. The mechanism affects peer and host-client access together. [Network fault semantics](network-partition.md)

A3 observed promotion at 3.473 s and stable ordinary service at 4.665 s from fault orchestration. At that stage the failed process and membership were unhealthy, while slots, leaders, reads, and ordinary writes were available; WAIT policy and redundancy remained unavailable. After restart, the final former primary was a synchronized replica of the promoted owner. Stable redundancy appeared at 12.241 s and stable WAIT policy at 12.245 s from fault orchestration. The promoted primary's full-sync and rejected-partial counters each increased by one. Validation included all 1,325 successful SETs: 1,070 WAIT-satisfied writes, 254 short-WAIT writes, and one write with a WAIT error. All were observed intact; 55 ambiguous SET errors remain outside that guarantee. The pacing ledger recorded 3,059 offered, 2,721 started/completed, and 338 dropped requests, so the configured rate was not the achieved rate during failure.

## What synchronization evidence establishes

The observer retains raw and parsed `INFO replication`, `INFO stats`, and `INFO persistence` alongside topology views. Optional derived fields include offsets, connected replicas, link/sync flags, last-I/O age, backlog length, and per-replica lag. Unknown or invalid values remain absent; offset gaps do not measure lost values.

Redis documents `sync_full`, `sync_partial_ok`, and `sync_partial_err` as full synchronization, accepted partial synchronization, and rejected partial synchronization counters. The lab compares them only on the same directly observed primary, endpoint, process run ID, and replication lineage. Missing, invalid, or decreasing counters make the interval unknown. These are primary-wide counters; a positive delta does not identify an exact target transfer or prove its completion. Link-up and `master_sync_in_progress=0` are separate sampled readiness checks. [Redis INFO fields](https://redis.io/docs/latest/commands/info/)

Milestones mark the first accepted sampled observation, not Redis event timestamps. Fault orchestration includes ownership checks before the Docker action; action duration and completion are separate. Repair-relative durations use the runtime monotonic clock. The configured delay follows completed observations and probes, so actual sample spacing defines temporal resolution. In these examples synchronization-in-progress was not sampled; no exact full/partial synchronization duration is claimed. [Measurement methodology](measurement-methodology.md)

## Retained resources and limits

Crash repair starts the same stopped managed container and preserves its Redis Cluster identity through explicitly owned disposable data volumes. It verifies the original network endpoints and a changed process start time. Memory-only mode retains membership files, while RDB and AOF remain disabled and volatile key values must be supplied by replication after restart. Isolation reconnects the same live process without a restart.

Volumes have exact run/node ownership labels, fixed names, local drivers, and validated `/data` mounts. There are no host bind mounts. Finite traffic and probe budgets, a 20,000-write/256-byte retained-data configuration cap, and Redis `maxmemory 128mb` limit experimental data generation; Docker volumes have no hard filesystem quota. Cleanup validates and removes owned containers, the empty bridge, and tracked volumes, then verifies absence. Cancellation and ambiguous fault/action responses preserve evidence and use fresh bounded restoration/cleanup contexts. Altered or foreign resources are refused. [ADR 0008](decisions/0008-retained-identity-and-independent-redundancy.md)

Observed post-experiment value equality and replica readiness do not establish zero lag, disk durability, uninterrupted availability, or survival of another failure. A process crash/container restart retains managed storage and does not simulate physical power loss, storage failure, or host failure. Redis replication and WAIT retain their existing limitations. [Redis replication](https://redis.io/docs/latest/operate/oss_and_stack/management/replication/), [SET/WAIT methodology](consistency-comparison.md)
