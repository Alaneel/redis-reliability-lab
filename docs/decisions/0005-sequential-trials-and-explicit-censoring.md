# ADR 0005: Sequential trials with explicit missing-event populations

Status: accepted

## Context

One failover gives neither a recovery-time distribution nor evidence that the same system behaves consistently. Parallel trials would compete for host resources and could obscure cleanup failures. Excluding failed trials or substituting zero for unobserved recovery events biases results.

## Decision

Run bounded batches sequentially, invoking the existing complete lifecycle for a fresh random run ID each time. Preserve each original report and write atomic per-file JSON/Markdown summaries. Verify report evidence and cleanup before continuing. Stop after cancellation or uncertain cleanup, explicitly retaining unstarted trial counts.

Summarize outcome counts over the entire requested population. Summarize event durations conditional on observation, alongside at-risk, not-at-risk, observed and unobserved populations and horizons. Partial outcomes can contribute genuine observed milestones. Never convert an unobserved event to zero or silently omit the unsuccessful trial.

At small trial counts, report descriptive min/median/mean/sample standard deviation/max and suppress trial tail percentiles. Preserve per-trial request quantiles and label statistics over them correctly; do not infer pooled request quantiles from their means.

## Consequences

Batches take longer than concurrent execution but offer clearer attribution and resource safety. Conditional event summaries are honest descriptive evidence, not a survival-analysis estimate. Shared-host correlations, limited sample counts, and sampling resolution remain explicit limitations.
