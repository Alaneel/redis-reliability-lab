# Evidence-grounded incident diagnosis

Phase 4 adds an offline diagnostic pipeline to the existing laboratory. The Redis experiment runner and its controlled repair behavior are unchanged. No external service, AI API, Redis production endpoint, database or new runtime dependency is required.

## Pipeline and interfaces

`internal/diagnosis.Normalize(ctx, reader)` reads a bounded schema 1–3 experiment report and projects an explicit observation allowlist into a versioned `Incident`. `Analyze(ctx, Incident)` detects observable conditions and correlates supporting, opposing/recovery and missing evidence. A separate timeline reconstruction derives transitions from successive samples. Recommendations are structured read-only plans. `Write` validates the citation graph and writes JSON and Markdown with exclusive creation.

`internal/evaluation` owns labels, trial families, dataset partitions and metric computation. It passes only original report bytes to the normalizer. The diagnosis package imports neither the evaluation package nor experiment, environment, config, report, Redis or Docker packages. This dependency boundary avoids using the orchestrator's knowledge as a diagnostic answer.

```sh
make build
bin/redis-lab diagnose docs/examples/phase3/primary-repair-wait/report.json --output reports/diagnoses
bin/redis-lab evaluate --dataset DATASET/manifest.json --partition heldout --output reports/evaluation
```

Diagnosis outputs are under the normalized incident identifier. Choose a new output root to repeat the same diagnosis: existing output is protected against overwrite. Evaluation can replace its generated result files; it never modifies input reports.

## Evidence and compatibility

Every record retains timestamp, node/peer identity, signal source, typed observed value, status, quality, JSON-pointer reference and collection-error classification. The original bytes' SHA-256 is recorded for provenance. Inference never consumes that checksum or incident identifiers.

Valid, explicitly unavailable/null, missing from older schemas, collection failed, not applicable and invalid observations have distinct states. A missing replica-link field on a primary is not applicable; missing administrative process liveness remains unknown. A failed node collection does not turn missing replication metrics into zero replicas. Dedicated `process_live` observations can be represented if supplied in node views; existing historical collectors do not supply that channel. Fault-specific process probes and action receipts remain evaluation-only.

The adapter ignores experiment names, filenames, paths, run IDs, configuration and hashes, phase names, selected fault targets, lifecycle dimensions, injected actions and precomputed recovery/outcome labels. No raw collection-error text is used to infer causes. Independent node and peer views are retained; conflicting contemporaneous peer flags, ownership or assignments are reported without electing an authoritative witness.

Input is bounded to 64 MiB, 4,000 samples, 64 views/sample, 3,600 client windows and 100,000 retained evidence records. Managed node identities are capped at 256 bytes, peer identities at 128 bytes, peer states at 64 per view and each projected field at 64 KiB to prevent expansion from oversized identifiers or fields. Overflow is rejected. Out-of-frame node observation timestamps are invalid rather than silently promoted to current evidence. Deterministic analysis checks context cancellation and uses no goroutines. Reports use source pointers rather than copying entire Redis INFO replies or command text.

## Observable diagnoses and inference limits

Rules distinguish prior-primary unavailability, prior-replica unavailability/link disruption, zero connected replicas, SET acknowledgements with short WAIT/WAITAOF results, policy command unavailability, full recorded slot coverage with failed client probes, persistence loading/write errors, missing/corrupt final acknowledged values, incomplete validation and terminal client errors. Multiple findings can coexist. New request windows also retain measured all-attempt p95 latency in a bounded histogram; legacy windows explicitly lack this field. A documented heuristic flags windows with at least 20 attempts whose p95 is at least five times an earlier error-free measured window and at least 20 ms. This is workload-sensitive evidence of change, not a calibrated anomaly probability.

Replica promotion is inferred only from an observed replica-to-primary transition. Endpoint unavailability is not proof of process death. Healthy slot coverage does not prove client recovery or value preservation. Short acknowledgement replies do not undo successful SETs. Replication offsets describe progress, never a count of lost values. Final correctness requires explicit final validation counters and a positive fully checked population; default zero counters do not imply preservation.

Assessment levels describe logical rule support, not calibrated confidence probabilities. Physical crash and network isolation generally remain competing causes. The engine explicitly abstains from selecting a unique root cause with the available channels. Detection can be `incident`, `healthy_observed` or `insufficient_evidence`. Healthy observations describe only recorded points, not uninterrupted availability.

Timeline timestamps are first sampled observations, with previous-observation bounds where available. They do not assert exact fault time. Slot coverage convergence, client recovery, acknowledgement restoration and replica-count restoration are independent events; a connected replica count is not proof of synchronization completion. Regressions and never-recovered incidents remain visible.

## Safety

A recommendation contains objective, supporting evidence, prerequisites, risks, expected effect, human intervention and verification criteria. No recommendation is executed. Plans require current identity/topology checks, operator approval for state changes and preservation of evidence/backups. Validation rejects execution modes other than `read_only_plan`.

The diagnostic code cannot force failover, change configuration, rejoin nodes, execute arbitrary commands or delete data. Existing laboratory fault/repair actions remain in their original orchestrator and managed Docker boundary.

## LLM status

No LLM integration is implemented. Deterministic evidence processing, explicit uncertainty and measurable evaluation are the baseline. An explanation provider would be a separate optional layer after these capabilities are validated; it would not become the authority for facts or execute commands.

Coverage normalization derives a Boolean from `assigned == 16384` and empty invalid-claim, missing, duplicate and unhealthy-owner lists in the recorded coverage object. Numeric INFO strings are parsed only when valid nonnegative integers. These are documented projections of recorded input, not independent measurements or calibrated confidence.
