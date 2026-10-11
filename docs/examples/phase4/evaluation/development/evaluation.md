# Diagnostic evaluation — development

Cases: 8. Dataset SHA-256: `39ddbf64fc33d99d7a630c12e3336eb7ec048ac04178552c8393a2c22e2e4aaf`.

Incident detection: TP 7, FP 0, TN 1, FN 0. Precision 1.000000, recall 1.000000.

| Category | TP | FP | TN | FN | Precision | Recall |
|---|---:|---:|---:|---:|---:|---:|
| primary_unavailability | 2 | 0 | 4 | 2 | 1.000000 | 0.500000 |
| replica_unavailability | 5 | 2 | 1 | 0 | 0.714286 | 1.000000 |
| redundancy_degradation | 3 | 0 | 1 | 4 | 1.000000 | 0.428571 |
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

Micro precision 0.789474, recall 0.714286. Root-cause abstentions 8/8; definitive claims 0, incorrect 0. Citations valid 842/842. Unsupported recommendations 0/19.

## Individual outcomes

| Case | Expected | Predicted | Detection | ms | Error |
|---|---|---|---|---:|---|
| case-001 | primary_unavailability, redundancy_degradation, client_disruption | client_disruption, primary_unavailability, replica_unavailability | incident | 56.910 |  |
| case-002 | replica_unavailability, redundancy_degradation | redundancy_degradation, replica_unavailability, topology_disagreement | incident | 108.023 |  |
| case-003 | primary_unavailability, redundancy_degradation, client_disruption | client_disruption, primary_unavailability, redundancy_degradation, replica_unavailability, topology_disagreement | incident | 173.306 |  |
| case-004 | replica_unavailability, redundancy_degradation | replica_unavailability | incident | 65.258 |  |
| case-005 | replica_unavailability, redundancy_degradation, client_disruption, acknowledgement_degradation | acknowledgement_degradation, client_disruption, redundancy_degradation, replica_unavailability | incident | 110.889 |  |
| case-006 | primary_unavailability, replica_unavailability, redundancy_degradation, data_validation_failure | data_validation_failure, replica_unavailability | incident | 117.948 |  |
| case-007 | primary_unavailability, replica_unavailability, redundancy_degradation | replica_unavailability | incident | 113.670 |  |
| case-016 |  |  | healthy_observed | 29.732 |  |

## Limitations

- Small single-host laboratory dataset; no generalization or production claims.
- Category matrices treat every selected case as one multilabel decision per category; overlap is allowed. Micro metrics sum category TP/FP/FN. Undefined denominators are null.
- Input and analysis failures count as undetected incidents and negative category predictions, and remain visible in case results.
- Abstention denominator is every selected case, including controls; definitive-claim counts are separate.
- Citation validity checks JSON-pointer existence in original bytes, not the causal truth of an interpretation.
- Unsupported recommendations check cited evidence availability and the read-only safety contract, not operator effectiveness.
- Execution latency includes input normalization and deterministic analysis, excluding file reading and report writing; timings are not deterministic output.
