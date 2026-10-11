# ADR 0007: Preserve connection affinity and the full SET acknowledgement ledger

Status: accepted; implemented, unit tested, and exercised in three real trials per policy on 2026-10-11

## Context

Adding replica acknowledgement changes both write completion cost and the definition of an application success. WAIT applies to preceding writes on its current connection, so independently routing SET and WAIT through a cluster client cannot establish the intended connection relationship. The server returns a replica count that the caller must compare with its requirement. See [Redis WAIT](https://redis.io/docs/latest/commands/wait/).

A successful SET followed by incomplete replica acknowledgement is still a client-acknowledged write. Dropping that record from final validation would conceal loss in exactly the cases this experiment should investigate. Replica acknowledgement also does not establish durable storage or strong consistency; see [Redis replication](https://redis.io/docs/latest/operate/oss_and_stack/management/replication/).

## Decision

Use the same managed writer for normal SET and SET plus WAIT. Borrow the cluster client's cached slot mapping, then use owned node pools with the same bounded configuration and private-to-loopback endpoint allowlist. Hold one sticky physical connection across SET and WAIT. Keep GET routing through the cluster client.

Follow only explicit MOVED/ASK rejections for the expected slot and an allowlisted destination, bounded to eight redirects. ASKING and SET use one ephemeral connection to prevent session-state reuse. Never replay a SET after a transport failure, and never retry, relocate, or replace the connection for WAIT after SET success. The overall client deadline bounds the logical operation.

Store every successful SET in the full immutable acknowledgement ledger. Preserve separate WAIT-attempt, count, error, and policy-satisfaction fields. Validate the full ledger and a distinct WAIT-satisfied subset after traffic stops. A short WAIT reply or WAIT error makes the logical policy unsatisfied while leaving SET acknowledgement evidence intact. Failed SET outcomes remain a separately counted ambiguous set outside the read-back guarantee.

Count each GET or logical write as one offered/started/completed operation. Record SET and WAIT diagnostic stage latency separately; these stages do not double the logical request counters. Report successful write throughput separately from throughput that includes reads and failed operations.

Keep ordinary SET/GET service recovery distinct from acknowledgement-policy recovery. Promotion can restore client service before replication redundancy is restored; a promoted primary with no replica cannot meet the selected replica-count policy. Cluster failover and acknowledgement receipt are observations with separate requirements. See the [Redis Cluster specification](https://redis.io/docs/latest/operate/oss_and_stack/reference/cluster-spec/).

## Consequences

Both policies share routing and retry behavior, which reduces a client-path confound in their comparison. The writer introduces owned per-node pools and an explicit lifecycle; it closes those pools without closing the borrowed cluster client. Policy failures remain visible in latency, throughput, count, and correctness accounting.

The fixed topology supports waiting for one replica. Increasing this requirement needs a different provisioned topology and readiness criteria. The experiment records final readability and value equality, not durability, linearizability, or an assurance that every future failover preserves acknowledged data. Detailed controls, comparison caveats, and executed evidence belong in [the comparison methodology](../consistency-comparison.md).
