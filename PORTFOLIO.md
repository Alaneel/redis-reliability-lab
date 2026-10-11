# Portfolio walkthrough

## The project in 60 seconds

Redis Reliability Lab is a Go CLI for studying what Redis Cluster recovery actually means. It creates a fresh local cluster with three primaries and three replicas, drives bounded synthetic traffic, injects a managed process or network fault, observes recovery, verifies acknowledged immutable writes, and checks cleanup before saving JSON and Markdown evidence.

The central design separates process liveness, topology, client service, replication redundancy, acknowledgement availability and data preservation. A promoted replica can serve requests while still lacking the replica needed by WAIT. Persistence tests stop all six processes and verify a frozen ledger before diagnostic writes can conceal missing data. SET, WAIT and WAITAOF use the same connection-aware routing path, with different success criteria. Repeated and counterbalanced runs preserve unsuccessful outcomes and measurement limits. The result is an executable distributed-systems experiment with inspectable evidence, rather than a claim that Redis failover is universally loss-free.

Start with the [first experiment](README.md#quick-start), [real CLI report](docs/examples/phase3/primary-failure-regression/report.md), [architecture](docs/architecture.md), and [release validation](docs/release-audit.md).

## Three engineering challenges and solutions

### 1. Service recovery can conceal degraded redundancy

A Redis process responding to PING or a promoted primary accepting SET does not establish full recovery. The observer reads each node directly, independently of the workload client's routing cache. Sustained topology and all-shard service checks have separate milestones from six-node membership, replica assignment, synchronization and acknowledgement-policy availability. Restarting a former primary preserves its Cluster ID, then requires observed demotion and synchronization with the promoted owner rather than forcing topology changes. Deadline-crossing samples cannot establish recovery.

The [replica lifecycle evidence](docs/replica-lifecycle.md) and [retained-identity decision](docs/decisions/0008-retained-identity-and-independent-redundancy.md) show the implementation and measured limits.

### 2. Acknowledgement policy failures must not erase successful SETs

WAIT and WAITAOF refer to preceding writes on the same physical connection. The managed writer holds one connection through SET and the selected acknowledgement command. Only validated MOVED/ASK rejections permit bounded SET redirection before acknowledgement; transport failures and post-SET acknowledgement failures are not replayed. Every successful SET enters the immutable ledger even when WAIT/WAITAOF returns a short reply or errors. SET errors remain ambiguous outcomes, while satisfied policy subsets are verified separately.

The real counterbalanced [AOF policy cohort](docs/examples/phase3/policy-comparison/comparison.md) checked 13,980 acknowledged values, including 668 short-WAIT and 730 short-WAITAOF records, with zero observed missing, corrupt, unreadable or unverified values. Its 359 ambiguous SET outcomes were excluded from that verified population. These are experimental observations, not universal guarantees. [Connection and accounting details](docs/consistency-comparison.md).

### 3. Persistence tests can accidentally preserve or recreate their own evidence

Leaving live replicas available after a primary crash cannot establish recovery from files. Writing diagnostic values before validation can hide missing records. The persistence experiment drains traffic, freezes a nonempty acknowledged ledger, kills replicas before primaries, verifies all six stopped without automatic restart, and restarts the same managed containers and disposable volumes. It verifies actual runtime configuration and files, then reads the frozen ledger before diagnostic writes.

The [persistence cohort](docs/examples/phase3/persistence-comparison/comparison.md) recorded all 1,251 memory-control values missing, while AOF everysec and always preserved 1,251 and 1,252 values. The host and storage remained running, and serial process kills permitted background fsync; these results do not establish power-loss durability or equivalence of fsync policies. [Persistence methodology](docs/durability-semantics.md).

## Five technical interview questions

1. Why can ordinary Redis Cluster service recover while WAIT remains unsatisfiable, and which observations establish restored redundancy?
2. When a SET times out, what can be concluded about execution, and how do immutable keys and separate ambiguous outcomes affect correctness validation?
3. Why must SET and WAIT/WAITAOF share a connection, and how should MOVED, ASK and post-SET transport failures change retry behavior?
4. What prevents a persistence experiment from being rescued by a live replica or repopulated by its diagnostic probes, and why does SIGKILL differ from physical power loss?
5. How do counterbalanced order, censored recovery outcomes and per-trial latency quantiles improve interpretation, and what remains uncertain with three trials per condition?

## Resume bullet candidates

- Built a Go/Docker Redis Cluster reliability lab with controlled fault injection, cluster-aware workloads, independent service and redundancy recovery, immutable-write validation, and ownership-checked cleanup.
- Validated failover, replica repair and AOF recovery with 22 real Docker integration test groups; retained counterbalanced SET/WAIT/WAITAOF evidence verifying 13,980 acknowledged writes while separating 359 ambiguous outcomes.

## Current limitations

- All Redis nodes share one physical host and Docker environment; independent machine, availability-zone, host and storage failures are outside the implemented model.
- Isolation removes an entire node's bridge interface in both directions. Selective links, one-way partitions and minority-side writes are unsupported.
- Persistence faults retain storage and page cache. Neither AOF preservation nor WAITAOF acknowledgements establish linearizability, physical power-loss durability or universally loss-free failover.
- Traffic and probes exercise one hash-tag location per initial shard, with bounded offered load and sampled observations. This is not a maximum-throughput benchmark or exhaustive slot test.
- Three trials per comparison condition do not support reliable population tails or statistical significance. Timing estimates include orchestration and sampling effects.
- The local Docker daemon and host user are trusted. Ownership checks cannot remove inspection-to-action races; synthetic Redis is unauthenticated, and disposable volumes lack hard filesystem quotas. [Security policy](SECURITY.md).
- Platform checks and hosted CI status are explicitly scoped in the [release audit](docs/release-audit.md). Existing reports preserve the configuration and runtime that actually produced them.
