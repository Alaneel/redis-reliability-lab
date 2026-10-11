# ADR 0008: Retained managed identity and independent redundancy restoration

Status: accepted; exercised by the real [replica crash](../examples/phase3/replica-crash/report.md), [replica isolation](../examples/phase3/replica-isolation/report.md), and [primary repair](../examples/phase3/primary-repair-wait/report.md) CLI examples.

## Context

Service can recover after failover while the promoted shard has no replica. Restarting an empty container can also lose its Redis Cluster identity. A running process or completed Docker action does not establish the intended membership, owner map, replica assignment, or data preservation.

## Decision

Crash-repair experiments retain six explicitly owned disposable local Docker volumes, including Redis membership files. Restart only the same stopped managed container, verify its original Redis Cluster ID and endpoints, and independently verify a new process start. Preserve an action receipt even when command output is ambiguous. Keep memory-only mode distinct from persisted key data.

Use pure baseline-aware topology predicates. Replica repair preserves all original leaders, slots, identities, endpoints, and replica assignments. Former-primary repair requires the original replica to retain promoted ownership and the restarted old primary to become its synchronized replica. Inspect current Redis state without forcing replication assignments or slot changes.

Reset stability at repair and require all-six topology and all-shard clients to sustain a fresh interval. Report process, membership, leaders, slots, reads, writes, data, acknowledgement policy, redundancy, and persistence independently. A historical service milestone cannot replace current final restoration. Primary-wide resync counters remain evidence with explicit identity and sampling limitations.

Validate volume labels, exact names, local drivers, and mounts before use/removal. Bound generated traffic and memory, expose the lack of a filesystem quota, and verify removal of tracked volumes during fresh-context cleanup. Preserve unrelated or altered resources.

## Consequences

The original primary-crash experiment still ends with reduced redundancy unless repair is requested. New repair runs can prove six-node replication readiness while leaving persistence recovery unobserved. Storage survives only the stated container/process fault; this does not model host/storage loss or prove durable acknowledgements. See [replica lifecycle evidence and limits](../replica-lifecycle.md).
