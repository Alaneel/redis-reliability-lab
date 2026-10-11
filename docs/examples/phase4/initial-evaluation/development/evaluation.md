# Diagnostic evaluation — development

Cases: 8. Dataset SHA-256: `1e063311347352d6fcd22ecb1248a662e1a9b8c304133a8b2699468ba9fafc19`.

Incident detection: TP 7, FP 1, TN 0, FN 0. Precision 0.875000, recall 1.000000.

| Category | TP | FP | TN | FN | Precision | Recall |
|---|---:|---:|---:|---:|---:|---:|
| primary_unavailability | 2 | 0 | 4 | 2 | 1.000000 | 0.500000 |
| replica_unavailability | 5 | 2 | 1 | 0 | 0.714286 | 1.000000 |
| redundancy_degradation | 7 | 1 | 0 | 0 | 0.875000 | 1.000000 |
| acknowledgement_degradation | 1 | 0 | 7 | 0 | 1.000000 | 1.000000 |
| client_recovery_lag | 0 | 0 | 8 | 0 | undefined | undefined |
| data_validation_failure | 1 | 0 | 7 | 0 | 1.000000 | 1.000000 |
| persistence_recovery | 0 | 0 | 8 | 0 | undefined | undefined |
| validation_incomplete | 0 | 0 | 8 | 0 | undefined | undefined |
| client_disruption | 3 | 0 | 5 | 0 | 1.000000 | 1.000000 |
| node_unavailability | 0 | 0 | 8 | 0 | undefined | undefined |
| acknowledgement_command_failure | 0 | 0 | 8 | 0 | undefined | undefined |
| topology_disagreement | 0 | 2 | 6 | 0 | 0.000000 | undefined |
| client_latency_change | 0 | 0 | 8 | 0 | undefined | undefined |

Micro precision 0.791667, recall 0.904762. Root-cause abstentions 8/8; definitive claims 0, incorrect 0. Citations valid 1665/1665. Unsupported recommendations 0/24.

## Individual outcomes

| Case | Expected | Predicted | Detection | ms | Error |
|---|---|---|---|---:|---|
| case-001 | primary_unavailability, redundancy_degradation, client_disruption | client_disruption, primary_unavailability, redundancy_degradation, replica_unavailability | incident | 66.517 |  |
| case-002 | replica_unavailability, redundancy_degradation | redundancy_degradation, replica_unavailability, topology_disagreement | incident | 106.663 |  |
| case-003 | primary_unavailability, redundancy_degradation, client_disruption | client_disruption, primary_unavailability, redundancy_degradation, replica_unavailability, topology_disagreement | incident | 171.086 |  |
| case-004 | replica_unavailability, redundancy_degradation | redundancy_degradation, replica_unavailability | incident | 64.780 |  |
| case-005 | replica_unavailability, redundancy_degradation, client_disruption, acknowledgement_degradation | acknowledgement_degradation, client_disruption, redundancy_degradation, replica_unavailability | incident | 104.730 |  |
| case-006 | primary_unavailability, replica_unavailability, redundancy_degradation, data_validation_failure | data_validation_failure, redundancy_degradation, replica_unavailability | incident | 125.778 |  |
| case-007 | primary_unavailability, replica_unavailability, redundancy_degradation | redundancy_degradation, replica_unavailability | incident | 113.233 |  |
| case-016 |  | redundancy_degradation | incident | 29.322 |  |

## Limitations

- Small single-host laboratory dataset; no generalization or production claims.
- Category matrices treat every selected case as one multilabel decision per category; overlap is allowed. Micro metrics sum category TP/FP/FN. Undefined denominators are null.
- Input and analysis failures count as undetected incidents and negative category predictions, and remain visible in case results.
- Abstention denominator is every selected case, including controls; definitive-claim counts are separate.
- Citation validity checks JSON-pointer existence in original bytes, not the causal truth of an interpretation.
- Unsupported recommendations check cited evidence availability and the read-only safety contract, not operator effectiveness.
- Execution latency includes input normalization and deterministic analysis, excluding file reading and report writing; timings are not deterministic output.
