# Redis Cluster primary failure

This document explains the concepts exercised by the lab's process-crash experiment. Redis behavior is defined by its version and configuration; the report describes one local run rather than a universal failover guarantee.

## Topology, slots, and replication

Redis Cluster partitions keys across 16,384 hash slots. Primaries own slots; replicas follow individual primaries and may replace them after failure. The lab starts three primaries with one replica each. Hash tags direct workload keys to one location in each initial shard. [Redis Cluster specification](https://redis.io/docs/latest/operate/oss_and_stack/reference/cluster-spec/)

The private cluster bus exchanges membership, failure information, and configuration updates. Client traffic uses the separate Redis command port. The host client uses published loopback ports while nodes communicate through internal addresses.

Each node has a local topology view. `CLUSTER NODES` exposes identities, roles, primary relationships, link status, failure flags, configuration epochs, and slot ownership. A single reply cannot prove every survivor has converged. The lab collects direct views and compares ownership and roles; local timestamps and observer-specific `myself` flags are excluded from equality. [CLUSTER NODES reference](https://redis.io/docs/latest/commands/cluster-nodes/)

Replication is asynchronous. A primary can acknowledge a write before its replica has received it. Losing the primary during that interval can lose an acknowledged value. An apparently healthy replication link does not prove a zero replication gap. The lab retains replication `INFO`, verifies acknowledged values, and reports missing data as an experimental outcome. [Redis replication](https://redis.io/docs/latest/operate/oss_and_stack/management/replication/)

## Failure detection and election

An abruptly killed node cannot communicate with peers. Redis initially records local suspicion (`PFAIL`, represented as `fail?` in `CLUSTER NODES`); failure reports and agreement among primaries can establish `FAIL`. An eligible replica then seeks election and promotion, with a majority of primaries participating. `cluster-node-timeout` influences detection, while election, networking, scheduling, and observation add further time. [Redis Cluster specification](https://redis.io/docs/latest/operate/oss_and_stack/reference/cluster-spec/)

The experiment injects `SIGKILL` into the primary owning the lowest initial slot. The designated replacement is its original replica. The framework never sends `CLUSTER FAILOVER`: that command initiates a manual failover with separate synchronization semantics, so it would exercise a different scenario. [Manual failover reference](https://redis.io/docs/latest/commands/cluster-failover/)

Full recovery requires every survivor to report a usable slot map and the expected replacement ownership. An elected primary can exist before every observer or client has caught up. `cluster_state:ok` and assigned/healthy slot counts are useful evidence, but the lab also parses ownership to reject missing slots, duplicates, unhealthy owners, and changed unaffected shards. [CLUSTER INFO reference](https://redis.io/docs/latest/commands/cluster-info/)

## Client routing and visible availability

`MOVED` directs a client to the current slot owner. `ASK` is a temporary redirect used during slot movement and requires the corresponding asking protocol. Handling these replies is ordinary Redis Cluster client behavior. The lab uses a cluster-aware client; errors in reports are terminal caller-visible errors after its configured redirect handling. [Redis Cluster specification](https://redis.io/docs/latest/operate/oss_and_stack/reference/cluster-spec/)

The dialer accepts only addresses mapped from this run's six containers. This handles Docker Desktop's separation between host and container addresses while preserving the managed-target boundary.

Client recovery probes run `SET` followed by matching `GET` on each of the three original shards. A successful probe of an unaffected shard alone cannot establish that the failed shard is writable. All three probes and topology predicates must remain positive across the configured stable interval. Background workload error rates and one-second windows quantify the other observed operations.

## Four meanings of recovery

| Meaning | Lab evidence |
|---|---|
| Process recovery | The crashed process starts again; deliberately absent here |
| Cluster recovery | Five survivors agree on expected roles and healthy complete slot coverage |
| Client-visible recovery | All three shard probes work for the stable interval; traffic metrics expose remaining errors |
| Data recovery | Every acknowledged immutable workload value remains readable and correct at validation |

Restoring replication redundancy is another independent condition. The promoted replica has no replacement in the MVP, so there are three surviving primaries and two replicas. A recovered service has less protection against a subsequent failure. The old primary can still appear as a failed member in gossip; recovery requires that it own no slots and remain unreachable.

## Availability and consistency limits

The crash leaves two original primaries and the failed primary's replica alive, allowing the intended automatic failover scenario. Different combinations of failed primaries/replicas or a partition that removes a primary majority can prevent service recovery. The lab models one process failure on one host; it does not simulate a distributed partition or prove availability under arbitrary failures.

Successful `SET` responses establish the ledger; response errors remain ambiguous because the server may have applied the command. Unique keys prevent overwrites from hiding missing acknowledged writes. Validation checks equality after traffic stops and reports incomplete checks, missing values, mismatches, and unreadable keys separately. Even a loss-free run cannot prove that acknowledged writes are always safe.

`WAIT` requests replication acknowledgements and can reduce some loss windows, but it does not establish strong consistency or guarantee that an acknowledged replica will be chosen after failure. The lab's [policy comparison](consistency-comparison.md) measures SET and connection-affine SET+WAIT separately, retains every SET acknowledgement for validation, and reports acknowledgement-policy degradation after promotion leaves the shard without a replica. [WAIT reference](https://redis.io/docs/latest/commands/wait/)

Persistence is disabled and data directories are temporary memory. The MVP therefore evaluates replication and failover rather than RDB/AOF durability, restart recovery, or storage resilience. Local Docker scheduling and network translation also affect timings; repeated runs with the same effective configuration provide a stronger comparison than a single result.
