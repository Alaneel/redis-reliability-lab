# Redis Reliability Lab — dataset-trial

Run: `abf1c05abe1dbfbc9ccad73d`  
Outcome: **healthy-control**  
Started: 2026-10-11T19:52:59.525776+08:00  
Finished: 2026-10-11T19:53:07.602202+08:00

## Recovery

| Condition | Observed |
|---|---|
| Surviving topology converged | false |
| All-shard client recovery sustained | false |
| Topology and client recovery jointly sustained | false |
| All acknowledged writes verified | false |
| Target process restarted | false |
| Original replication redundancy restored | false |
| Recovery deadline expired | false |
| Cleanup verified | true |

The primary stays stopped. Service recovery can succeed with reduced redundancy. Redis replication is asynchronous; an acknowledged SET is not a durability guarantee. Missing acknowledged writes cause a partial outcome.

## Timing

Milestones are first sampled observations, not Redis event timestamps. Orchestration includes ownership checks and fault execution. The action-completion column separates that overhead; neither anchor identifies the exact Redis event time. Unobserved milestones remain absent.

| Observation | Since orchestration began, ms | Since action completed, ms |
|---|---:|---:|
| Failure flagged in gossip | unobserved | unobserved |
| Original replica owns failed primary slots | unobserved | unobserved |
| Full usable slot coverage | unobserved | unobserved |
| Surviving views converge | unobserved | unobserved |
| First all-shard SET/GET probe succeeds | unobserved | unobserved |
| Client probes sustain recovery | unobserved | unobserved |
| Topology and clients sustain recovery | unobserved | unobserved |

## Workload

| Phase | Attempts | Success rate | Errors | Throughput (ops/s) | Successful p50/p95/p99 (ms) |
|---|---:|---:|---:|---:|---|

Acknowledged writes: 0. Ambiguous SET outcomes: 0.

Validation attempted: false; expected: 0; checked: 0; missing: 0; corrupt: 0; unreadable: 0; unverified: 0. Validation complete: false.
Validation was not attempted; 0 acknowledged writes remain unchecked.

### Terminal workload errors

| Phase | Error class | Count |
|---|---|---:|



## Measurement methodology


The JSON report contains effective configuration, environment metadata, raw topology replies, per-operation metrics, validation samples, and lifecycle timestamps. Effective configuration SHA-256: ``.
