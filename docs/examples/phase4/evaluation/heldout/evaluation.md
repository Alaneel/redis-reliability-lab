# Diagnostic evaluation — heldout

Cases: 9. Dataset SHA-256: `e9e50a12ce4b9dd219a443df8b2e14f5a3ea760c54d9a816f0b50759ecac4cd8`.

Incident detection: TP 7, FP 0, TN 1, FN 1. Precision 1.000000, recall 0.875000.

| Category | TP | FP | TN | FN | Precision | Recall |
|---|---:|---:|---:|---:|---:|---:|
| primary_unavailability | 2 | 0 | 4 | 3 | 1.000000 | 0.400000 |
| replica_unavailability | 5 | 2 | 2 | 0 | 0.714286 | 1.000000 |
| redundancy_degradation | 3 | 0 | 1 | 5 | 1.000000 | 0.375000 |
| acknowledgement_degradation | 1 | 0 | 8 | 0 | 1.000000 | 1.000000 |
| client_recovery_lag | 0 | 0 | 9 | 0 | undefined | undefined |
| data_validation_failure | 1 | 0 | 8 | 0 | 1.000000 | 1.000000 |
| persistence_recovery | 0 | 0 | 9 | 0 | undefined | undefined |
| validation_incomplete | 0 | 0 | 9 | 0 | undefined | undefined |
| client_disruption | 3 | 0 | 5 | 1 | 1.000000 | 0.750000 |
| node_unavailability | 0 | 0 | 9 | 0 | undefined | undefined |
| acknowledgement_command_failure | 0 | 0 | 9 | 0 | undefined | undefined |
| topology_disagreement | 0 | 2 | 7 | 0 | 0.000000 | undefined |
| client_latency_change | 0 | 4 | 5 | 0 | 0.000000 | undefined |

Micro precision 0.652174, recall 0.625000. Root-cause abstentions 9/9; definitive claims 0, incorrect 0. Citations valid 906/906. Unsupported recommendations 0/23.

## Individual outcomes

| Case | Expected | Predicted | Detection | ms | Error |
|---|---|---|---|---:|---|
| case-001 |  |  | healthy_observed | 34.699 |  |
| case-002 | primary_unavailability, redundancy_degradation, client_disruption | client_disruption, client_latency_change, primary_unavailability, replica_unavailability | incident | 56.122 |  |
| case-002-incomplete | primary_unavailability, redundancy_degradation, client_disruption |  | insufficient_evidence | 0.160 |  |
| case-003 | replica_unavailability, redundancy_degradation | redundancy_degradation, replica_unavailability | incident | 96.463 |  |
| case-004 | primary_unavailability, redundancy_degradation, client_disruption | client_disruption, client_latency_change, primary_unavailability, redundancy_degradation, replica_unavailability, topology_disagreement | incident | 173.335 |  |
| case-005 | replica_unavailability, redundancy_degradation | client_latency_change, replica_unavailability | incident | 66.633 |  |
| case-006 | replica_unavailability, redundancy_degradation, client_disruption, acknowledgement_degradation | acknowledgement_degradation, client_disruption, client_latency_change, redundancy_degradation, replica_unavailability, topology_disagreement | incident | 104.618 |  |
| case-007 | primary_unavailability, replica_unavailability, redundancy_degradation, data_validation_failure | data_validation_failure, replica_unavailability | incident | 81.201 |  |
| case-008 | primary_unavailability, replica_unavailability, redundancy_degradation | replica_unavailability | incident | 92.297 |  |

## Limitations

- Small single-host laboratory dataset; no generalization or production claims.
- Category matrices treat every selected case as one multilabel decision per category; overlap is allowed. Micro metrics sum category TP/FP/FN. Undefined denominators are null.
- Input and analysis failures count as undetected incidents and negative category predictions, and remain visible in case results.
- Abstention denominator is every selected case, including controls; definitive-claim counts are separate.
- Citation validity checks JSON-pointer existence in original bytes, not the causal truth of an interpretation.
- Unsupported recommendations check cited evidence availability and the read-only safety contract, not operator effectiveness.
- Execution latency includes input normalization and deterministic analysis, excluding file reading and report writing; timings are not deterministic output.
