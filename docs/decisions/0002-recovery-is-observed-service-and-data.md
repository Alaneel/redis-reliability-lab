# ADR 0002: Separate topology, service, data, process, and redundancy recovery

**Status:** Accepted

## Context

Redis asynchronous replication and eventual gossip mean a live process, an elected primary, a working client, and retained data can recover at different times. One aggregate boolean would hide these differences.

## Decision

Observe every managed node directly and retain its source replies. Require healthy complete slot coverage, convergence among five survivors, promotion of the original replica, and preservation of unaffected ownership. Combine that with probes to all three original shards spanning a configured stability interval. Stop traffic before checking every acknowledged immutable write.

Report topology recovery, client recovery, data verification, process restart, and redundancy restoration separately. The killed process remains stopped; neither restarting it nor restoring redundancy is required for this service-failover experiment. An acknowledged data violation prevents a recovered outcome. A recovery timeout is an experimental result, not fabricated success.

## Consequences

Recovery predicates are stricter than a single `PING` or `cluster_state:ok` check. Sampling has bounded temporal resolution and cannot prove uninterrupted availability between observations. Reports explicitly expose reduced redundancy after a successful failover. Future faults must define their own expected topology and recovery predicates.
