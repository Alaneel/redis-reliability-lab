# Diagnostic evaluation — heldout

Cases: 9. Dataset SHA-256: `1e063311347352d6fcd22ecb1248a662e1a9b8c304133a8b2699468ba9fafc19`.

Incident detection: TP 7, FP 1, TN 0, FN 1. Precision 0.875000, recall 0.875000.

| Category | TP | FP | TN | FN | Precision | Recall |
|---|---:|---:|---:|---:|---:|---:|
| primary_unavailability | 2 | 0 | 4 | 3 | 1.000000 | 0.400000 |
| replica_unavailability | 5 | 2 | 2 | 0 | 0.714286 | 1.000000 |
| redundancy_degradation | 7 | 1 | 0 | 1 | 0.875000 | 0.875000 |
| acknowledgement_degradation | 1 | 0 | 8 | 0 | 1.000000 | 1.000000 |
| client_recovery_lag | 0 | 0 | 9 | 0 | undefined | undefined |
| data_validation_failure | 1 | 0 | 8 | 0 | 1.000000 | 1.000000 |
| persistence_recovery | 0 | 0 | 9 | 0 | undefined | undefined |
| validation_incomplete | 0 | 0 | 9 | 0 | undefined | undefined |
| client_disruption | 3 | 0 | 5 | 1 | 1.000000 | 0.750000 |
| node_unavailability | 0 | 0 | 9 | 0 | undefined | undefined |
| acknowledgement_command_failure | 0 | 0 | 9 | 0 | undefined | undefined |
| topology_disagreement | 0 | 4 | 5 | 0 | 0.000000 | undefined |
| client_latency_change | 0 | 0 | 9 | 0 | undefined | undefined |

Micro precision 0.730769, recall 0.791667. Root-cause abstentions 9/9; definitive claims 0, incorrect 0. Citations valid 1547/1547. Unsupported recommendations 0/26.

## Individual outcomes

| Case | Expected | Predicted | Detection | ms | Error |
|---|---|---|---|---:|---|
| case-008 | primary_unavailability, redundancy_degradation, client_disruption | client_disruption, primary_unavailability, redundancy_degradation, replica_unavailability | incident | 55.342 |  |
| case-008-incomplete | primary_unavailability, redundancy_degradation, client_disruption |  | insufficient_evidence | 0.036 |  |
| case-009 | replica_unavailability, redundancy_degradation | redundancy_degradation, replica_unavailability, topology_disagreement | incident | 105.126 |  |
| case-010 | primary_unavailability, redundancy_degradation, client_disruption | client_disruption, primary_unavailability, redundancy_degradation, replica_unavailability, topology_disagreement | incident | 160.530 |  |
| case-011 | replica_unavailability, redundancy_degradation | redundancy_degradation, replica_unavailability, topology_disagreement | incident | 63.263 |  |
| case-012 | replica_unavailability, redundancy_degradation, client_disruption, acknowledgement_degradation | acknowledgement_degradation, client_disruption, redundancy_degradation, replica_unavailability, topology_disagreement | incident | 92.237 |  |
| case-013 | primary_unavailability, replica_unavailability, redundancy_degradation, data_validation_failure | data_validation_failure, redundancy_degradation, replica_unavailability | incident | 83.806 |  |
| case-014 | primary_unavailability, replica_unavailability, redundancy_degradation | redundancy_degradation, replica_unavailability | incident | 83.982 |  |
| case-017 |  | redundancy_degradation | incident | 28.388 |  |

## Limitations

- Small single-host laboratory dataset; no generalization or production claims.
- Category matrices treat every selected case as one multilabel decision per category; overlap is allowed. Micro metrics sum category TP/FP/FN. Undefined denominators are null.
- Input and analysis failures count as undetected incidents and negative category predictions, and remain visible in case results.
- Abstention denominator is every selected case, including controls; definitive-claim counts are separate.
- Citation validity checks JSON-pointer existence in original bytes, not the causal truth of an interpretation.
- Unsupported recommendations check cited evidence availability and the read-only safety contract, not operator effectiveness.
- Execution latency includes input normalization and deterministic analysis, excluding file reading and report writing; timings are not deterministic output.
