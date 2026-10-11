# ADR 0003: Use bounded pacing, histograms, and an acknowledged-write ledger

**Status:** Accepted

## Context

During outage, unbounded request queues distort latency and consume memory. Retaining every operation for exact percentiles is unnecessary for a local reliability experiment. Overwritten keys can hide acknowledged-write loss.

## Decision

Use a shared ticker consumed by a bounded number of workers, dropping ticks when workers are occupied. Record caller-visible completions in separate successful, failed, and all-attempt geometric latency histograms. Attribute operations to the phase when they start. Retain bounded one-second windows and bounded topology samples with explicit dropped counts.

Write unique immutable keys and retain only successfully acknowledged values in a bounded ledger. Treat error-returning writes as ambiguous. Check ledger values during reads and after draining the workload. Use one tag per initial shard to guarantee traffic on the shard selected for failure.

## Consequences

The configured rate is an aggregate start ceiling, not a guaranteed open-loop arrival rate. Percentiles are estimates, with 5% geometric bucket growth. Scheduling and failover remain nondeterministic despite a seeded workload. Three tagged shard locations provide controlled coverage rather than a uniform keyspace benchmark. Reports state these limits explicitly.
