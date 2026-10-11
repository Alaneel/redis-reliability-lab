# Redis Reliability Lab — counterbalanced comparison

Comparison: `comparison-dde55f9ada692afef5838308`  
Outcome: **partial**  
Order seed: `20261011`  
Started: 2026-10-11T08:50:18.083444Z  
Finished: 2026-10-11T08:58:02.56308Z

Conditions: 3; rounds: 3; requested: 9; started: 9; callbacks completed: 9; unstarted: 0. Base condition order: [2 1 3].

## Conditions

Statistics stay within each condition. Failed and partial trials remain represented; censored or missing observations are not zero-duration recoveries. No pooled request latency quantiles or causal inference is presented.

| Condition | Experiment | Workload seed | Recovered | Partial | Failed | Aborted | Unstarted | Summary |
|---|---|---:|---:|---:|---:|---:|---:|---|
| condition-01 | persistence-memory | 42 | 0 | 3 | 0 | 0 | 0 | [JSON](conditions/condition-01/summary.json), [Markdown](conditions/condition-01/summary.md) |
| condition-02 | persistence-aof-everysec | 42 | 3 | 0 | 0 | 0 | 0 | [JSON](conditions/condition-02/summary.json), [Markdown](conditions/condition-02/summary.md) |
| condition-03 | persistence-aof-always | 42 | 3 | 0 | 0 | 0 | 0 | [JSON](conditions/condition-03/summary.json), [Markdown](conditions/condition-03/summary.md) |

## Actual schedule

| Position | Round | Round position | Condition | Run ID | Outcome | Evidence verified | Cleanup verified | Report |
|---:|---:|---:|---|---|---|---|---|---|
| 1 | 1 | 1 | condition-02 | 90345e3ab363a8699699a536 | recovered | true | true | [JSON](trials/position-0001/batch-b9522008735afacedb206207/runs/90345e3ab363a8699699a536/report.json), [Markdown](trials/position-0001/batch-b9522008735afacedb206207/runs/90345e3ab363a8699699a536/report.md) |
| 2 | 1 | 2 | condition-01 | aa2a96b0899c2885b2aa3a32 | partial | true | true | [JSON](trials/position-0002/batch-aedea818cdd8ecb1335a5885/runs/aa2a96b0899c2885b2aa3a32/report.json), [Markdown](trials/position-0002/batch-aedea818cdd8ecb1335a5885/runs/aa2a96b0899c2885b2aa3a32/report.md) |
| 3 | 1 | 3 | condition-03 | bc5ed8626797ee3fa6013069 | recovered | true | true | [JSON](trials/position-0003/batch-b36014fbb9afa7b331dd7462/runs/bc5ed8626797ee3fa6013069/report.json), [Markdown](trials/position-0003/batch-b36014fbb9afa7b331dd7462/runs/bc5ed8626797ee3fa6013069/report.md) |
| 4 | 2 | 1 | condition-01 | fa1b57f987b8c265178b1829 | partial | true | true | [JSON](trials/position-0004/batch-9b1b272f35ff29ce65565565/runs/fa1b57f987b8c265178b1829/report.json), [Markdown](trials/position-0004/batch-9b1b272f35ff29ce65565565/runs/fa1b57f987b8c265178b1829/report.md) |
| 5 | 2 | 2 | condition-03 | a50433e55ee7e3e72db5414c | recovered | true | true | [JSON](trials/position-0005/batch-ea3add7d0827fb08c748b425/runs/a50433e55ee7e3e72db5414c/report.json), [Markdown](trials/position-0005/batch-ea3add7d0827fb08c748b425/runs/a50433e55ee7e3e72db5414c/report.md) |
| 6 | 2 | 3 | condition-02 | 29bae03a4d4db021a7868f0f | recovered | true | true | [JSON](trials/position-0006/batch-8f7c142e2101fdf1c2cf419c/runs/29bae03a4d4db021a7868f0f/report.json), [Markdown](trials/position-0006/batch-8f7c142e2101fdf1c2cf419c/runs/29bae03a4d4db021a7868f0f/report.md) |
| 7 | 3 | 1 | condition-03 | 516b2a2428d8d408b3fc1b4f | recovered | true | true | [JSON](trials/position-0007/batch-b1104f16b924e36fcac1a492/runs/516b2a2428d8d408b3fc1b4f/report.json), [Markdown](trials/position-0007/batch-b1104f16b924e36fcac1a492/runs/516b2a2428d8d408b3fc1b4f/report.md) |
| 8 | 3 | 2 | condition-02 | f8e160bbc8659f1e524fb427 | recovered | true | true | [JSON](trials/position-0008/batch-ec728a99fe31dc8937ae3d2d/runs/f8e160bbc8659f1e524fb427/report.json), [Markdown](trials/position-0008/batch-ec728a99fe31dc8937ae3d2d/runs/f8e160bbc8659f1e524fb427/report.md) |
| 9 | 3 | 3 | condition-01 | 4a3cd1c67b9ff55137ab5115 | partial | true | true | [JSON](trials/position-0009/batch-61bb928a8f8d405beb4b630e/runs/4a3cd1c67b9ff55137ab5115/report.json), [Markdown](trials/position-0009/batch-61bb928a8f8d405beb4b630e/runs/4a3cd1c67b9ff55137ab5115/report.md) |

## Execution errors

- schedule position 2 (condition-01): trial 1: persistence experiment ended partial: frozen data verified=false, original redundancy stable=true, restoration deadline=false; schedule position 4 (condition-01): trial 1: persistence experiment ended partial: frozen data verified=false, original redundancy stable=true, restoration deadline=false; schedule position 9 (condition-01): trial 1: persistence experiment ended partial: frozen data verified=false, original redundancy stable=true, restoration deadline=false

## Measurement methodology

- A seeded permutation of condition indices is rotated by one position each round. Within each complete block of N rounds, every condition occupies every round position once. An incomplete block is only partially balanced; this rotating Latin order does not balance all predecessor/successor pairs.
- The order seed controls only the condition schedule. Effective workload seeds and all other configuration values remain unchanged. Each entry executes one fresh sequential batch trial with persisted-report integrity and cleanup checks.
- Further trials require verified individual evidence and successful cleanup, or verified evidence that no environment was provisioned. Unsafe cleanup, missing evidence, cancellation and summary-write failures stop execution; all remaining planned entries stay explicitly unstarted.
- Failed and partial trials with valid evidence and safe cleanup remain in their condition statistics. Timing statistics are conditional on observed milestones; missing and censored observations are not zero-duration events. Request latency quantiles are not pooled across trials or conditions.
- Condition summaries reuse batch aggregation. Condition-summary report paths are relative to their own summary directory; schedule report paths and batch references are relative to the comparison directory. Host and environment references point to individual JSON report sections.
- Ordering reduces systematic position confounding but does not prove causality, independent samples, equal host conditions or population reliability. Compare individual environmental metadata and actual order before interpreting differences.
