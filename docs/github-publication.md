# GitHub publication validation

The public repository is [Alaneel/redis-reliability-lab](https://github.com/Alaneel/redis-reliability-lab), published on 11 October 2026. Its Go module and internal import paths match that address. Existing experiment reports and initial local audit evidence are retained unchanged; they describe the source and environment that produced them.

## First hosted run and test correction

The [first Ubuntu workflow](https://github.com/Alaneel/redis-reliability-lab/actions/runs/38134070090), at commit `d0e7c08acf61d939a76148c970031b2a084d41c0`, installed Go 1.26.9 and verified dependencies, but failed three subcases of `TestPersistenceDiagnosticsKeepPostStabilityRegressionAndPolicyIndependent`. Docker integration was consequently skipped. The [original failure log](examples/github-publication/first-unit-failure.log) is preserved.

That test combined synthetic sample timestamps with a one-second wall-clock context. Race-instrumented full-slot validation on a slower host could consume that budget, excluding the third regression sample and retaining earlier stable flags. The test now uses explicit cancellation and a fourth-sample guard, preserving the assertion that the third sample changes current service/policy state. Cancellation exclusion remains tested. Production recovery code and its real timeout budgets are unchanged.

A controlled diagnostic used temporary copies with the same 1.1-second delay before the third sample. The [old fixture failed](examples/github-publication/old-controlled-delay.log) with `context deadline exceeded` and an excluded final sample; the [corrected fixture passed](examples/github-publication/fixed-controlled-delay.log). Those induced failures are a diagnostic reproduction, not failed real Redis experiments. The unchanged repository fixture passed [five repeated race-test runs](examples/github-publication/focused-fixed.log), covering all four subcases each time. The [complete corrected local check](examples/github-publication/local-fixed-check.log) also passed formatting, vet, fresh race tests in all ten packages and the release build.

## Current hosted status

Consult the [Actions workflow](https://github.com/Alaneel/redis-reliability-lab/actions/workflows/ci.yml) for the latest exact commit and outcome. A source push or local pass alone is not hosted verification. The v1.0.0 tag and GitHub release are created only after the corrected commit passes both the Ubuntu unit/build job and the real Docker integration job. Final hosted results and their raw integration evidence are linked in the [GitHub Releases](https://github.com/Alaneel/redis-reliability-lab/releases).
