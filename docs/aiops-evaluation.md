# Diagnostic evaluation methodology

The dataset generator is an opt-in real Docker integration test in `internal/evaluation`. It provisions separate six-node clusters for development and heldout partitions, varying workload seeds, request rate, sampling interval and failure duration. Existing deterministic target selection still chooses the same shard position; fresh node identities alone do not establish broad target generalization.

```sh
REDIS_LAB_DATASET=reports/new-aiops-dataset make aiops-integration
bin/redis-lab evaluate --dataset reports/new-aiops-dataset/manifest.json --partition development --output reports/development
bin/redis-lab evaluate --dataset reports/new-aiops-dataset/manifest.json --partition heldout --output reports/heldout
```

An existing manifest cannot be overwritten by collection. Healthy controls use actual all-shard SET/GET probes. Fresh fault trials cover primary crash, replica crash, primary network isolation, replica network isolation, SET+WAIT replica loss, all-process memory restart and all-process AOF restart. Every trial verifies managed resource cleanup. One heldout primary trial also has an evidence-deletion derivative: its real incident label is retained, so an abstention causes a detection false negative. That derivative is not another independent physical trial.

The label manifest contains fault type, target, fault timing and label basis. Physical incident labels use controlled action receipts; client impact and final-value outcomes use independent runner metrics and the immutable acknowledged ledger. Labels are not generated from diagnostic predictions. Some observable symptoms can be absent despite a known physical fault: these cases remain false negatives rather than being relabelled healthy.

Trial families must stay within one partition, duplicate report hashes are rejected, unknown categories and inconsistent healthy labels fail validation. Input checksums and path/symlink containment prevent silently changing or substituting reports. Dataset manifest hashes identify the exact evaluated snapshot. Metadata-mutating tests verify that fault labels and experiment configuration do not alter the normalized evidence or inference.

## Metrics and denominators

Incident detection uses one binary decision per selected case. Category results use one binary decision per case per category, allowing overlapping conditions. Each category reports TP/FP/TN/FN, precision TP/(TP+FP) and recall TP/(TP+FN). Micro results sum these counts across categories. Undefined denominators are JSON null rather than invented zero/one percentages.

Normalization/analysis errors count as undetected incidents and negative category predictions; the actual errors remain in individual results. Root-cause abstention is counted across all cases, including controls, with definitive and incorrect definitive claims reported separately. This baseline abstains from unique physical causes because independent administrative evidence is insufficient; that behavior must not be presented as high root-cause accuracy.

Citation validity checks existence of every cited evidence record's JSON pointer in the original input, including findings, timelines and recommendations. It does not validate causal truth. Unsupported recommendation rate checks available cited evidence plus the read-only plan contract and prerequisites/risks/verification; it does not measure intervention effectiveness. Execution latency measures normalization plus analysis, excluding report input/output, and is reported per case.

The sample is small and single-host. Development and heldout trials share Redis version, fault mechanism and laboratory topology. Timing variation, seeds and fresh processes reduce exact duplication but do not establish production generalization. Negative tests cover malformed input, missing and stale observations, conflicting peers, cancellation, unknown schema and incomplete validation. Observability gaps and incorrect predictions must remain in the evaluation reports.

## Label scope and conservative category matching

The predeclared labels principally describe physical action targets and known outcomes. The engine also reports observational phenomena such as replica link disruption after a primary failure, peer-view disagreement and latency changes. These extra findings are counted as unmatched category false positives even when their cited observations are real; they were not independently adjudicated as additional gold categories. Consequently micro precision/recall measure exact agreement with this restricted label protocol, rather than complete expert-adjudicated diagnostic correctness. This limitation must accompany any quoted category metric. A future evaluation needs blinded human annotation of overlapping symptoms, not retroactively adding labels to match predictions.

The first rule version incorrectly treated an initial zero connected-replica observation during startup as a decline. Its original development and heldout results, exact frozen source archive and all original inputs are retained. The corrected rule requires a positive prior recorded replica count. It fixes the startup false positive but deliberately misses degradation when the historical baseline never recorded a positive count. A fresh eight-trial heldout cohort was collected after this correction. Its evidence-deletion derivative is a ninth evaluation record, not another independent trial.

## Executed results

Usable collection consists of 24 fresh physical trials: 16 in the initial development/evaluation dataset and eight in the replacement heldout cohort. Two evidence-deletion derivatives produce 26 evaluation records across these datasets; they are not independent trials. The initial collection also attempted two healthy controls that failed report writing after cleanup; they are excluded from the usable-trial count and their failure logs are retained.

The final scored partitions are eight original development trials and nine fresh heldout records (eight trials plus one derivative). Each physical cohort has one healthy control and one each of primary crash, replica crash, primary network isolation, replica network isolation, SET+WAIT replica loss, memory restart and AOF restart. Their shared single-host topology and deterministic shard selection limit generalization.

| Measure | Development | Replacement heldout |
|---|---:|---:|
| Cases | 8 | 9 |
| Root-cause abstentions | 8 | 9 |
| Definitive root-cause claims | 0 | 0 |
| Incorrect definitive claims | 0 | 0 |
| Citations checked | 842 | 906 |
| Valid citations | 842 | 906 |
| Read-only plans | 19 | 23 |
| Unsupported plans | 0 | 0 |
| Detection precision | 1.000000 | 1.000000 |
| Detection recall | 1.000000 | 0.875000 |
| Category micro precision | 0.789474 | 0.652174 |
| Category micro recall | 0.714286 | 0.625000 |

Detection counts: development TP=7, FP=0, TN=1, FN=0; replacement heldout TP=7, FP=0, TN=1, FN=1. The deliberately evidence-deleted physical primary incident is undetected and correctly returns insufficient evidence; it still counts as a detection false negative. Both healthy controls avoid incident classification under the corrected rule.

Heldout multilabel counts are TP=15, FP=8, TN=85, FN=9 across 13 categories and nine cases. Refer to the restricted label-scope warning above before interpreting category precision: additional observed latency/peer/replica-link phenomena count as unmatched findings, and absent positive replica baselines cause deliberate misses.

No unique physical root cause is selected: abstention is 9/9 heldout records, including the control. Zero incorrect definitive claims has a zero definitive-claim denominator; it is not 100% root-cause accuracy.

Observed development normalization/analysis latency: median 109.456 ms, maximum 173.306 ms on this invocation. Timings exclude input/output and are not a production SLA.
Observed heldout normalization/analysis latency: median 81.201 ms, maximum 173.335 ms on this invocation. Timings exclude input/output and are not a production SLA.

The new heldout collection passed all eight physical subcases with race detection in 280.008 s. The original full collection failed its two healthy-control report writes; all 14 fault subcases passed, and the corrected standalone healthy-control collection passed in 29.369 s. Failed invocations are not rewritten as suite passes. Label target and request/completion intervals were separately audited against actual receipts; category labels and input report bytes were not changed to match predictions.

[Final heldout JSON](examples/phase4/evaluation/heldout/evaluation.json), [complete confusion matrices and individual outcomes](examples/phase4/evaluation/heldout/evaluation.md), [development results](examples/phase4/evaluation/development/evaluation.md), [fresh original inputs and audited labels](examples/phase4/heldout-dataset/manifest.json), [collection log](examples/phase4/heldout-collection.log), [rules-2 source freeze](examples/phase4/rule-freeze-rules2.json), [exact source archive](examples/phase4/rules-2-source.tar.gz).

The rejected first rule version remains inspectable through [original heldout results](examples/phase4/initial-evaluation/heldout/evaluation.md), [original manifest](examples/phase4/initial-dataset/manifest.json), [separately audited label provenance](examples/phase4/initial-dataset/manifest-audited.json), [failed initial collection](examples/phase4/collection.log), [corrected control invocation](examples/phase4/controls-corrected.log), [first source freeze](examples/phase4/rule-freeze.json) and [exact rules-1 source archive](examples/phase4/rules-1-source.tar.gz).

A later parser hardening adds strict identifier/peer/field limits without changing rule decisions. [Complete assessment replay proof](examples/phase4/input-guard-replay-proof.json) shows identical hashes on all 29 retained corpus records against the exact frozen evaluation source; this replay is a regression check, not another blind evaluation.
