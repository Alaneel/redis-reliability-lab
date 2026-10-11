# ADR 0004: Record and bound cluster client topology refresh

**Status:** Accepted

## Context

The client maintains a cached slot map. After a primary crashes, that primary cannot return a `MOVED` response directing requests to the promoted replica. A client's topology refresh policy therefore affects application-visible recovery independently of Redis election and gossip convergence.

An early real Docker failover run with go-redis v9.23.0's default client refresh setting observed surviving topology recovery at about 3.18 seconds from the start of validated fault injection but no client recovery before a 20-second deadline. Validation then encountered unreadable values through the stale route. Dependency code inspection confirmed a default 60-second lazy refresh interval. This was an actual unsuccessful experiment, rather than evidence of Redis recovery failure.

## Decision

Set the client lazy topology refresh interval explicitly to one second. Keep genuine cluster-aware routing/redirection behavior, direct observer connections, strict managed-address translation, and operation deadlines. Do not manually substitute the observer's topology into the workload client to force apparent recovery.

Set RESP2, redirect/retry-loop limit eight, disabled per-node command retries, one attempt per pool dial, and per-node pools of 16 connections. Record these settings as environment metadata so timing comparisons can account for client policy.

## Consequences

Client recovery timing includes this policy and cannot be attributed only to server-side election time. Refresh is lazy and asynchronous, triggered when the client uses an aged state, so one second is a configured interval rather than a hard recovery guarantee. The direct topology observer remains independent. Future experiments comparing clients should expose deliberate client variants and record their refresh/retry policies.
