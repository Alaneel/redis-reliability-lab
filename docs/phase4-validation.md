# Phase 4 implementation and completion evidence

The extension implements offline incident intelligence in Go, preserving the original distributed-systems experiment runner. It normalizes allowlisted schema 1–3 observations, records missingness and provenance, detects correlated symptoms, reconstructs sampled timelines, preserves competing explanations and generates read-only plans. It evaluates separate laboratory labels and emits inspectable JSON/Markdown. There is no LLM provider or API key requirement.

## Architecture and implemented interfaces

`diagnosis.Normalize(ctx, reader) → Incident → diagnosis.Analyze(ctx, Incident) → Assessment → diagnosis.Write(directory, assessment)`. The engine has no Redis, Docker, experiment or evaluation imports. The evaluator receives labels separately and supplies only original report bytes to inference. The CLI exposes `diagnose` and `evaluate`. The existing metrics collector additionally records bounded request-window p95 latency; older windows retain missing status rather than fabricated latency.

See [architecture](aiops-architecture.md), [evaluation methodology and actual results](aiops-evaluation.md) and [three actual case studies](aiops-case-studies.md).

## Executed engineering validation

Executed locally with Go 1.26.9 on macOS arm64, Docker Engine 24.0.6 and Redis 8.10.2:

```sh
make check build
# gofmt verification, go vet ./..., fresh go test -race -count=1 ./...
# all 12 packages passed; binary built with trimpath and buildvcs=false

go test -race -count=1 ./internal/diagnosis
# extra replay, aligned-correlation and real timeout/cancellation tests passed

REDIS_LAB_DATASET=reports/phase4-heldout-rules2 \
REDIS_LAB_DATASET_PARTITION=heldout make aiops-integration
# eight fresh physical heldout subcases passed; 280.008 seconds

make integration
# go test -race -tags=integration -count=1 -timeout=20m -v ./internal/experiment
# all 22 top-level Docker groups, including three persistence mode subtests,
# passed in 792.501 seconds
```

The first dataset invocation failed its two healthy-control report writes because the output directories were missing. Its 14 physical fault subcases passed. The corrected control invocation passed both controls in 29.369 seconds; the complete replacement heldout invocation passed all eight subcases. The original failed invocation remains saved. The evaluator also exposed and retained the first rule version's startup replica-count false positive, corrected it using development evidence and collected a new heldout cohort. Original and corrected rules have exact source archives and hashes.

The complete regression produced 24 JSON/Markdown experiment reports: nine recovered, three partial, eight failed and four aborted experimental outcomes. Those negative outcomes are expected failure-path assertions; all Go test groups passed. One additional WAITAOF compatibility run is represented by the actual command log rather than an experiment report. Every saved run report verified cleanup; independent post-suite Docker inventories found zero managed containers, networks and volumes.

[Full test log](examples/phase4/full-integration.log), [raw regression reports archive](examples/phase4/full-integration-reports.tar.gz), [derived regression summary](examples/phase4/integration-summary.json), [final 12-package check/build log](examples/phase4/final-check-build.log), [final additional diagnosis tests](examples/phase4/diagnosis-final-tests.log), [fresh heldout collection log](examples/phase4/heldout-collection.log).

The GitHub Actions definition now includes unit/build, original Docker regression and fresh AIOps dataset/evaluation jobs with preserved evidence artifacts. The three-job YAML parsed successfully. At completion of the local validation recorded here, Phase 4 had not been pushed and the added hosted job had not run. The public v1.0 hosted result validates the previous release. Subsequent hosted validation and publication evidence are recorded with the [v1.1.0 release](https://github.com/Alaneel/redis-reliability-lab/releases/tag/v1.1.0); the measurements below remain the original local results.


A final input-boundary review added identifier, peer-count and projected-field caps. These guards change rejection of oversized malformed input, not diagnostic rules. Isolated replay binaries built from the exact frozen rules-2 archive and current guarded source produced identical complete-assessment hashes for all 29 original/fresh/real-failure corpus records. [Replay proof](examples/phase4/input-guard-replay-proof.json) and [adversarial guard tests](examples/phase4/input-guard-tests.log) preserve this distinction; the frozen source archive is not silently replaced by later code.

## Requirement evidence audit

| Requirement | Implemented evidence / executed verification |
|---|---|
| Preserve existing Redis experiments | Complete 22-group real Docker regression plus fresh 12-package checks; existing runner source unchanged |
| Accept real historical reports | Offline CLI processed primary repair, replica crash and memory reports; tests cover schema 1–3 and real failure-path reports |
| Structured evidence / timeline | Versioned typed model, allowlist normalizer, source checksums/pointers, sampled transitions and explicit unavailable/missing/invalid states |
| Multiple categories | Primary/replica endpoint observations, replication decline, short acknowledgement policies, client lag/latency, persistence and value-validation conditions; meaningful positive/negative tests |
| Evidence citations and competing explanations | Citation graph validation, same-sample correlation tests, peer disagreement retention and deterministic root-cause abstention |
| Healthy controls / insufficient evidence | Both corrected healthy controls avoided incident classification; evidence-deleted real incident returned insufficient evidence and remains a false negative |
| Read-only remediation | Structured prerequisite/risk/effect/verification plans; unsafe execution and missing-prerequisite tests; no Redis/Docker dependencies in diagnostic package |
| Ground-truth isolation | Label/config/identifier mutation tests, one-way projection and package boundary; action/target receipts used only for label audits |
| Fresh heldout evaluation | Replacement eight-trial cohort after rule correction, with separate physical identities/seeds/timing and same-family derivative |
| Actual metrics / visible errors | Category matrices, precision/recall, abstention, definitive errors, citation/plan rates and per-case latency; original false positive and current misses retained |
| JSON and Markdown output | Actual CLI outputs for three fresh incidents and both evaluated partitions |
| Negative/adversarial coverage | Missing/stale/conflicting input, unknown signal/schema, malformed reports, cancellation, real failed/aborted/partial trajectories, false-positive prevention and deterministic replay |
| Resource bounds / idiomatic implementation | Bounded input/sample/view/window/evidence counts, existing bounded histogram, context checks, no new runtime dependency or diagnostic goroutines |
| Documentation / reproducibility | Architecture, methodology, three real case studies, exact frozen rule source, raw input reports and checksum manifest |
| CI integration | Three explicit jobs configured; local completion predates hosted execution. See the v1.1.0 release for subsequent hosted evidence |
| Optional LLM | Not implemented; deterministic evidence processing and evaluation were prioritized |

## Final evaluation and limitations

The replacement heldout set contains eight independent physical trials and one evidence-deletion derivative. Detection is TP=7, FP=0, TN=1, FN=1: precision 7/7, recall 7/8. The missing-evidence derivative explains the detection false negative. Category micro agreement is TP=15, FP=8, TN=85, FN=9: precision 15/23, recall 15/24 across 13 categories and nine records. The restricted physical-target label protocol does not adjudicate every observed latency, peer disagreement or replica-link symptom; quote category results with this limitation.

All nine heldout records abstain from selecting a unique physical root cause, including the healthy control. There are zero definitive claims; zero incorrect definitive claims is not a root-cause accuracy score. All 906 cited references resolve into original report bytes. All 23 recommendations satisfy the read-only/cited-prerequisite safety contract; this does not measure operational effectiveness. Median observed normalization/analysis latency is 81.201 ms, maximum 173.335 ms in the recorded invocation.

This is a small, single-host lab validation. It neither distinguishes every crash from a network fault nor establishes production integration, causally correct explanations from pointer existence, or strong generalization. Missing positive pre-fault replica counts cause deliberate missed redundancy classifications. All-process stopped-barrier action fields are excluded from inference, leaving primary interruption incompletely observable. Failed/partial/aborted experimental outcomes, classifier errors and uncertainty remain visible.

## Defensible resume bullet candidates

- Built an offline Go incident-diagnosis pipeline for Redis Cluster, normalizing versioned telemetry into source-cited findings, sampled timelines, explicit uncertainty and safety-constrained read-only remediation plans.
- Implemented isolated diagnostic evaluation using 24 fresh six-node Redis trials and preserved failures; a replacement nine-record heldout set achieved 7/7 incident-detection precision and 7/8 recall while abstaining from unsupported unique physical root causes.
- Added bounded request-window latency telemetry and deterministic provenance/correlation tests; validated 12 packages with race detection and passed 22 real Docker regression groups spanning failover, replication repair, WAIT/WAITAOF and persistence.

## Three highest-value next improvements

1. Collect continuous baseline and restart-phase observations independently of fault orchestration, including anonymous administrative process/network evidence, so missing positive replica baselines and process interruption no longer cause predictable misses.
2. Add read-only per-request routing, redirect, connection and timeout traces to distinguish observed client recovery lag from specific cached-route or connection causes.
3. Expand fresh blinded cohorts across physical target shards, load, failure duration and hosts, with independent expert adjudication of overlapping symptom categories. Recalibrate thresholds only against development data and report confidence intervals without treating heuristic scores as probabilities.
