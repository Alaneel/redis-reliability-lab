# Documentation

Start with the [repository quick start](../README.md#quick-start) or [portfolio walkthrough](../PORTFOLIO.md). All examples use fresh local six-node clusters and synthetic data; the saved effective configuration is authoritative for each report.

## Run and understand experiments

| Guide | What it covers |
|---|---|
| [Running experiments](running-experiments.md) | Prerequisites, CLI commands, configuration defaults, outcomes, and cleanup |
| [Architecture](architecture.md) | Local Docker boundary, orchestration, observations, workload, report contract, and safety |
| [Measurement methodology](measurement-methodology.md) | Sequential trials, counterbalanced order, timing anchors, denominators, censoring, and small-sample limits |
| [Redis failover](redis-failover.md) | Slots, gossip, elections, asynchronous replication, routing, and service recovery |
| [Network partition](network-partition.md) | Whole-node bidirectional isolation, unchanged process survival, exact endpoint restoration, and reconciliation |
| [Replica lifecycle](replica-lifecycle.md) | Replica faults, retained identity, former-primary repair, synchronization evidence, and restored redundancy |
| [Consistency comparison](consistency-comparison.md) | SET/WAIT/WAITAOF accounting, connection affinity, policy recovery, and matched comparisons |
| [Durability semantics](durability-semantics.md) | Memory and AOF controls, all-process stopped barrier, frozen ledger, runtime/file evidence, and process-crash limits |

## Validation and raw evidence

Validation pages record what was executed, including earlier failed checks and final regression results. They are historical local evidence, not proof of a hosted CI run. Raw examples under `examples/` remain unchanged.

| Record | Evidence to inspect first |
|---|---|
| [Release audit](release-audit.md) | Release checks, packaging, and platform scope |
| [Phase 3 validation](phase3-validation.md) | Replica lifecycle, persistence, WAITAOF, two counterbalanced comparisons, and final regression suite |
| [Phase 2 validation](phase2-validation.md) | Repeated crash trials, network failure paths, and SET/WAIT comparison |
| [Original validation](validation.md) | Initial primary-failover implementation and executed checks |

- **Original CLI regression:** [Markdown](examples/phase3/primary-failure-regression/report.md) · [JSON](examples/phase3/primary-failure-regression/report.json). Service recovered with reduced redundancy; all 2,030 acknowledged values were checked, and 87 ambiguous SETs were separate.
- **Final integration suite:** [complete raw log](examples/phase3/full-integration.log) · [execution record](phase3-validation.md#final-complete-docker-regression-suite). All 22 top-level Docker test groups passed in the recorded run.
- **Persistence comparison:** [Markdown](examples/phase3/persistence-comparison/comparison.md) · [JSON](examples/phase3/persistence-comparison/comparison.json). Nine fresh trials include observed memory loss and AOF preservation within the retained-storage process-crash model.
- **Acknowledgement-policy comparison:** [Markdown](examples/phase3/policy-comparison/comparison.md) · [JSON](examples/phase3/policy-comparison/comparison.json). Nine fresh AOF-always primary-repair trials retain policy subsets, short replies, ambiguous SETs, and individual reports.
- **Frozen WAITAOF persistence:** [Markdown](examples/phase3/waitaof-persistence/report.md) · [JSON](examples/phase3/waitaof-persistence/report.json). Read verification occurs before diagnostic writes.

Each comparison's links lead to per-condition summaries and individual reports. Read the relevant methodology before comparing values across definitions, phases, or report schemas.

## Architecture decisions

| Decision | Reason to read it |
|---|---|
| [0001 — Local Docker boundary](decisions/0001-local-docker-boundary.md) | Reproducibility and owned-resource limits |
| [0002 — Recovery is observed service and data](decisions/0002-recovery-is-observed-service-and-data.md) | Why process liveness cannot establish recovery |
| [0003 — Bounded workload and measurements](decisions/0003-bounded-workload-and-measurements.md) | Load pacing, immutable values, and bounded evidence |
| [0004 — Explicit client refresh](decisions/0004-explicit-client-refresh.md) | Cached routing after abrupt primary failure |
| [0005 — Sequential trials and explicit censoring](decisions/0005-sequential-trials-and-explicit-censoring.md) | Trial isolation and conditional statistics |
| [0006 — Full-node network isolation](decisions/0006-full-node-network-isolation.md) | Exact fault model and original endpoint restoration |
| [0007 — Connection-affine replica acknowledgements](decisions/0007-connection-affine-replica-acknowledgements.md) | SET and WAIT must refer to the same connection history |
| [0008 — Retained identity and independent redundancy](decisions/0008-retained-identity-and-independent-redundancy.md) | Repair, process replacement, and redundancy criteria |
| [0009 — Frozen-ledger all-process persistence](decisions/0009-frozen-ledger-all-process-persistence.md) | Preventing live replicas or diagnostic writes from masking loss |

For development and resource safety, see [Contributing](../CONTRIBUTING.md) and the [Security policy](../SECURITY.md).

## Incident intelligence

- [AIOps architecture, evidence model, diagnostic rules and safety](aiops-architecture.md)
- [Evaluation construction, partitions, metrics and limitations](aiops-evaluation.md)
- [Three real incident diagnosis case studies](aiops-case-studies.md)
- [Phase 4 validation, acceptance audit, portfolio bullets and remaining work](phase4-validation.md)
