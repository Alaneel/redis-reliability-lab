# Real incident diagnosis case studies

These are fresh Redis Docker trials from the replacement heldout cohort. Ground truth is shown here for evaluation; it is excluded from diagnostic inference. Each diagnosis was produced by the actual offline CLI against the unmodified report bytes. No unique physical root cause is asserted by the engine.

## Primary endpoint failure and promotion

Original trial `a272b5d9e6773d9c5fa6c493`, case `case-002`, outcome `recovered`. [Original JSON](examples/phase4/heldout-dataset/runs/a272b5d9e6773d9c5fa6c493/report.json), [diagnosis JSON](examples/phase4/diagnoses/b1ae66e62637f3e24d310204/diagnosis.json), [diagnosis Markdown](examples/phase4/diagnoses/b1ae66e62637f3e24d310204/diagnosis.md).

Ground truth: controlled `primary` action, target `node-1`. Action request/completion anchors are in the separate [label manifest](examples/phase4/heldout-dataset/manifest.json). The action label is not a diagnostic input.

The diagnostic input is the allowlisted node/client/replication/persistence/value observations. Normalization records the original byte checksum and JSON-pointer provenance, omitting the action, target selection, configuration, filenames and lifecycle conclusions.

- `client_disruption`: The all-shard client probe failed. Evidence `e000441, e000443, e000650`; rule `R5-observed-client-probe-failure`.
- `client_latency_change`: Terminal request-window p95 is at least five times an earlier error-free measured window and at least 20 ms. Workload mix and acknowledgement waits can explain this; it is not a calibrated anomaly probability. Evidence `e000445, e000004, e000442`; rule `R15-measured-window-p95-5x-and-20ms-with-20-attempts`.
- `primary_unavailability`: Node observation endpoint unavailable; the cause remains unresolved. Evidence `e000237, e000124, e000560`; rule `R2-endpoint-unavailable-with-prior-role`.
- `replica_unavailability`: Replica INFO reports its primary link down; process and network causes remain alternatives. Evidence `e000254, e000449, e000786`; rule `R4-replica-link-down`.

Final validation checked 504 of 504 acknowledged values: missing 0, corrupt 0, unreadable 0, unverified 0. This population excludes ambiguous SET outcomes; it does not prove general durability.

The engine identifies primary endpoint unavailability and role transition while abstaining between crash and network causes. It misses the independently labelled redundancy degradation because the saved pre-fault baseline lacks a positive primary replica count. The additional replica-link/latency findings are unmatched by the predeclared physical-target labels and count as category false positives in this evaluation.

Recommendations require current identity/topology verification, preservation of evidence/backups and human approval before state changes. No Redis command is executed by diagnosis. Remaining uncertainties include independent process/network evidence, per-request routing traces and complete temporal coverage.

Reproduce with `bin/redis-lab diagnose docs/examples/phase4/heldout-dataset/runs/a272b5d9e6773d9c5fa6c493/report.json --output reports/reproduce-primary`. Existing output cannot be overwritten.

## Replica failure with short WAIT acknowledgements

Original trial `42ff17db5678dfbe4b21a5d7`, case `case-006`, outcome `recovered`. [Original JSON](examples/phase4/heldout-dataset/runs/42ff17db5678dfbe4b21a5d7/report.json), [diagnosis JSON](examples/phase4/diagnoses/d386be492b555b00da8cbd11/diagnosis.json), [diagnosis Markdown](examples/phase4/diagnoses/d386be492b555b00da8cbd11/diagnosis.md).

Ground truth: controlled `wait` action, target `node-5`. Action request/completion anchors are in the separate [label manifest](examples/phase4/heldout-dataset/manifest.json). The action label is not a diagnostic input.

The diagnostic input is the allowlisted node/client/replication/persistence/value observations. Normalization records the original byte checksum and JSON-pointer provenance, omitting the action, target selection, configuration, filenames and lifecycle conclusions.

- `acknowledgement_degradation`: At least one acknowledged SET received a valid short WAIT/WAITAOF reply. The SET succeeded; the additional acknowledgement policy was not satisfied. Evidence `e000443, e000445, e000442`; rule `R8-acknowledged-set-valid-short-policy-reply`.
- `client_disruption`: Client request window contains terminal errors; policy and SET outcomes must be inspected separately. Evidence `e000454, e000879, e001110`; rule `R7-terminal-client-errors`.
- `client_latency_change`: Terminal request-window p95 is at least five times an earlier error-free measured window and at least 20 ms. Workload mix and acknowledgement waits can explain this; it is not a calibrated anomaly probability. Evidence `e000456, e000004, e000453`; rule `R15-measured-window-p95-5x-and-20ms-with-20-attempts`.
- `redundancy_degradation`: A directly observed primary reports no connected replicas; service availability and synchronization remain separate questions. Evidence `e000253, e000125, e000625`; rule `R3-primary-zero-connected-replicas-after-positive-baseline`.
- `replica_unavailability`: Node observation endpoint unavailable; the cause remains unresolved. Evidence `e000237, e000048, e000457`; rule `R2-endpoint-unavailable-with-prior-role`.
- `topology_disagreement`: Contemporaneous peer views disagree about the same node; no single view is chosen as authoritative. Evidence `e000694, e000738`; rule `R14-conflicting-contemporaneous-peer-views`.

Final validation checked 1303 of 1303 acknowledged values: missing 0, corrupt 0, unreadable 0, unverified 0. This population excludes ambiguous SET outcomes; it does not prove general durability.

The workload ledger separately recorded 271 valid short WAIT replies. Diagnostic probes cite successful SET acknowledgements together with insufficient replica counts; a short WAIT is not relabelled a failed SET. The engine reports replica disruption, redundancy decline and acknowledgement degradation, with additional latency and peer-disagreement findings retained as unmatched categories.

Recommendations require current identity/topology verification, preservation of evidence/backups and human approval before state changes. No Redis command is executed by diagnosis. Remaining uncertainties include independent process/network evidence, per-request routing traces and complete temporal coverage.

Reproduce with `bin/redis-lab diagnose docs/examples/phase4/heldout-dataset/runs/42ff17db5678dfbe4b21a5d7/report.json --output reports/reproduce-wait`. Existing output cannot be overwritten.

## Serving requests after memory restart with acknowledged-value loss

Original trial `d4afc5196b21a7355b5de2ba`, case `case-007`, outcome `partial`. [Original JSON](examples/phase4/heldout-dataset/runs/d4afc5196b21a7355b5de2ba/report.json), [diagnosis JSON](examples/phase4/diagnoses/f22e11d33b9fb3a555bad521/diagnosis.json), [diagnosis Markdown](examples/phase4/diagnoses/f22e11d33b9fb3a555bad521/diagnosis.md).

Ground truth: controlled `memory` action, target `all-six-managed-nodes`. Action request/completion anchors are in the separate [label manifest](examples/phase4/heldout-dataset/manifest.json). The action label is not a diagnostic input.

The diagnostic input is the allowlisted node/client/replication/persistence/value observations. Normalization records the original byte checksum and JSON-pointer provenance, omitting the action, target selection, configuration, filenames and lifecycle conclusions.

- `data_validation_failure`: Final acknowledged-value validation observed missing or corrupt values. Serving requests does not establish data preservation. Evidence `e002075`; rule `R10-positive-validation-loss-counter`.
- `replica_unavailability`: Replica INFO reports its primary link down; process and network causes remain alternatives. Evidence `e000366, e000404, e000442`; rule `R4-replica-link-down`.

Final validation checked 191 of 191 acknowledged values: missing 191, corrupt 0, unreadable 0, unverified 0. This population excludes ambiguous SET outcomes; it does not prove general durability.

All checked acknowledged values were missing even though later diagnostic service probes succeeded. The engine reports observed data-validation failure rather than equating service recovery with preservation. Physical primary unavailability and redundancy degradation remain missed because intermediate administrative stop/action fields are excluded and a usable positive replica baseline is absent. These are visible false negatives, not reconstructed from the injected label.

Recommendations require current identity/topology verification, preservation of evidence/backups and human approval before state changes. No Redis command is executed by diagnosis. Remaining uncertainties include independent process/network evidence, per-request routing traces and complete temporal coverage.

Reproduce with `bin/redis-lab diagnose docs/examples/phase4/heldout-dataset/runs/d4afc5196b21a7355b5de2ba/report.json --output reports/reproduce-memory`. Existing output cannot be overwritten.
