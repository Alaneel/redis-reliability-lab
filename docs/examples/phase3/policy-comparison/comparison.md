# Redis Reliability Lab — counterbalanced comparison

Comparison: `comparison-321638985ff49f9040049bf4`  
Outcome: **recovered**  
Order seed: `20261011`  
Started: 2026-10-11T09:11:38.844832Z  
Finished: 2026-10-11T09:17:22.385436Z

Conditions: 3; rounds: 3; requested: 9; started: 9; callbacks completed: 9; unstarted: 0. Base condition order: [2 1 3].

## Conditions

Statistics stay within each condition. Failed and partial trials remain represented; censored or missing observations are not zero-duration recoveries. No pooled request latency quantiles or causal inference is presented.

| Condition | Experiment | Workload seed | Recovered | Partial | Failed | Aborted | Unstarted | Summary |
|---|---|---:|---:|---:|---:|---:|---:|---|
| condition-01 | policy-aof-set | 42 | 3 | 0 | 0 | 0 | 0 | [JSON](conditions/condition-01/summary.json), [Markdown](conditions/condition-01/summary.md) |
| condition-02 | policy-aof-wait | 42 | 3 | 0 | 0 | 0 | 0 | [JSON](conditions/condition-02/summary.json), [Markdown](conditions/condition-02/summary.md) |
| condition-03 | policy-aof-waitaof | 42 | 3 | 0 | 0 | 0 | 0 | [JSON](conditions/condition-03/summary.json), [Markdown](conditions/condition-03/summary.md) |

## Actual schedule

| Position | Round | Round position | Condition | Run ID | Outcome | Evidence verified | Cleanup verified | Report |
|---:|---:|---:|---|---|---|---|---|---|
| 1 | 1 | 1 | condition-02 | dc0464195769627e90802e83 | recovered | true | true | [JSON](trials/position-0001/batch-28685cb04eeec1fdcdf25b49/runs/dc0464195769627e90802e83/report.json), [Markdown](trials/position-0001/batch-28685cb04eeec1fdcdf25b49/runs/dc0464195769627e90802e83/report.md) |
| 2 | 1 | 2 | condition-01 | dab1232736a2b59ad498e657 | recovered | true | true | [JSON](trials/position-0002/batch-bb39bfe07e9868b38d98ea5d/runs/dab1232736a2b59ad498e657/report.json), [Markdown](trials/position-0002/batch-bb39bfe07e9868b38d98ea5d/runs/dab1232736a2b59ad498e657/report.md) |
| 3 | 1 | 3 | condition-03 | e519c821277d8791d7e6796e | recovered | true | true | [JSON](trials/position-0003/batch-bb6f098573c5605a2fe65248/runs/e519c821277d8791d7e6796e/report.json), [Markdown](trials/position-0003/batch-bb6f098573c5605a2fe65248/runs/e519c821277d8791d7e6796e/report.md) |
| 4 | 2 | 1 | condition-01 | 60141689fee92e0e4cd9bc77 | recovered | true | true | [JSON](trials/position-0004/batch-c7b388dcb147ac10049fac7b/runs/60141689fee92e0e4cd9bc77/report.json), [Markdown](trials/position-0004/batch-c7b388dcb147ac10049fac7b/runs/60141689fee92e0e4cd9bc77/report.md) |
| 5 | 2 | 2 | condition-03 | 688a7dc24b138ea24d1b7e0d | recovered | true | true | [JSON](trials/position-0005/batch-05a52c53dbeff938cd618ed7/runs/688a7dc24b138ea24d1b7e0d/report.json), [Markdown](trials/position-0005/batch-05a52c53dbeff938cd618ed7/runs/688a7dc24b138ea24d1b7e0d/report.md) |
| 6 | 2 | 3 | condition-02 | a41f7d40e0b182eec96b39bb | recovered | true | true | [JSON](trials/position-0006/batch-a4333dacfdfe4f1ff20e1e7a/runs/a41f7d40e0b182eec96b39bb/report.json), [Markdown](trials/position-0006/batch-a4333dacfdfe4f1ff20e1e7a/runs/a41f7d40e0b182eec96b39bb/report.md) |
| 7 | 3 | 1 | condition-03 | 37475e6b0cf7b098669c0663 | recovered | true | true | [JSON](trials/position-0007/batch-82ffccc231329f6e4d7dae7b/runs/37475e6b0cf7b098669c0663/report.json), [Markdown](trials/position-0007/batch-82ffccc231329f6e4d7dae7b/runs/37475e6b0cf7b098669c0663/report.md) |
| 8 | 3 | 2 | condition-02 | 73a8d11d94ba9a7f56fb11b1 | recovered | true | true | [JSON](trials/position-0008/batch-4c4de00323adeee2bc70b1a9/runs/73a8d11d94ba9a7f56fb11b1/report.json), [Markdown](trials/position-0008/batch-4c4de00323adeee2bc70b1a9/runs/73a8d11d94ba9a7f56fb11b1/report.md) |
| 9 | 3 | 3 | condition-01 | 8d7edf6681a3f591879efeda | recovered | true | true | [JSON](trials/position-0009/batch-d0662642c7cf1047aa35a9ed/runs/8d7edf6681a3f591879efeda/report.json), [Markdown](trials/position-0009/batch-d0662642c7cf1047aa35a9ed/runs/8d7edf6681a3f591879efeda/report.md) |

## Measurement methodology

- A seeded permutation of condition indices is rotated by one position each round. Within each complete block of N rounds, every condition occupies every round position once. An incomplete block is only partially balanced; this rotating Latin order does not balance all predecessor/successor pairs.
- The order seed controls only the condition schedule. Effective workload seeds and all other configuration values remain unchanged. Each entry executes one fresh sequential batch trial with persisted-report integrity and cleanup checks.
- Further trials require verified individual evidence and successful cleanup, or verified evidence that no environment was provisioned. Unsafe cleanup, missing evidence, cancellation and summary-write failures stop execution; all remaining planned entries stay explicitly unstarted.
- Failed and partial trials with valid evidence and safe cleanup remain in their condition statistics. Timing statistics are conditional on observed milestones; missing and censored observations are not zero-duration events. Request latency quantiles are not pooled across trials or conditions.
- Condition summaries reuse batch aggregation. Condition-summary report paths are relative to their own summary directory; schedule report paths and batch references are relative to the comparison directory. Host and environment references point to individual JSON report sections.
- Ordering reduces systematic position confounding but does not prove causality, independent samples, equal host conditions or population reliability. Compare individual environmental metadata and actual order before interpreting differences.
