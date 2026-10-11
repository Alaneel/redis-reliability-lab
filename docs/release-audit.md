# Portfolio release v1.0 audit

This audit records initial local release preparation on 11 October 2026. It separates local source/build validation, actual Docker experiments, and hosted GitHub Actions. At that point, no repository had been published or remote created. The owner subsequently authorized publication to [Alaneel/redis-reliability-lab](https://github.com/Alaneel/redis-reliability-lab). The Go module and internal import paths were changed to that actual repository address; the saved local checks and source manifests below describe the prepublication snapshot. Hosted validation must be assessed separately from those historical local results.

## Scope and changes

The existing architecture and fault models are retained. The README now introduces the problem, first experiment, distributed-systems concepts and real evidence before deeper guides. [The documentation index](README.md) and [running guide](running-experiments.md) organize the detailed material. [PORTFOLIO.md](../PORTFOLIO.md) supplies a technical explanation, challenges, interview questions, factual resume candidates and limitations.

The release audit found one strict-config defect: the former 1 MiB limited reader silently truncated larger files, allowing a second document or unknown field beyond the truncation boundary to disappear. The decoder now reads at most limit plus one byte and rejects oversized input before YAML decoding. Tests cover exact-limit acceptance and oversized valid, multi-document and unknown-field inputs. No experiment feature or failure model was added. All 460 local file, directory and heading links across 112 Markdown files resolved during the documentation audit.

GitHub Actions retains verified exact action commit pins, least-privilege permissions and disabled persisted checkout credentials. It now uses the installed pinned toolchain and readonly module mode, verifies dependencies before checks, sequences integration after unit validation, captures complete integration output, and bounds jobs/steps. `make build VERSION=1.0.0` stamps the existing version variable; default source builds remain `dev`.

## Executed local checks

| Check | Evidence and scope |
|---|---|
| macOS arm64 formatting, vet, fresh race tests and build | [Complete log](examples/release-v1.0/macos-check-build.log); all ten packages passed with Go 1.26.9, release-stamped build completed |
| macOS arm64 full Docker integration | [Complete fresh log](examples/release-v1.0/macos-full-integration.log); all 22 top-level Docker groups and three persistence subcases passed with race detection in **785.061 s**, including expected failure, timeout and cancellation outcomes |
| Linux arm64 formatting, vet, fresh race tests, dependency integrity and build | [Complete log](examples/release-v1.0/linux-check-build.log); all ten packages passed natively in a Debian 12 container, Go 1.26.9, Linux 6.4.16-linuxkit |
| Linux amd64 build | Cross-built successfully with `GOOS=linux GOARCH=amd64 CGO_ENABLED=0`; not executed on an amd64 host |
| Example configuration validation | All 16 YAML definitions accepted by the Linux release CLI; saved check log preserves individual results |
| Clean Git index export | [Build and validation log](examples/release-v1.0/clean-index-build.log); dependencies verified, all 16 configurations valid, `1.0.0` binary byte-identical to the workspace build with SHA256 `b54f87e94fe88f306ffad71cefb52b4468ed251a234c00131061bd4b686a1419` |
| Clean-export original CLI experiment | [CLI log](examples/release-v1.0/clean-index-cli.log), [report](examples/release-v1.0/primary-failure/report.md) and [raw JSON](examples/release-v1.0/primary-failure/report.json); `recovered`, all 2,020 acknowledged writes checked with zero missing, corrupt, unreadable or unverified values; 91 ambiguous SET outcomes remain separate; cleanup successful |
| Workflow syntax | [`actionlint 1.7.12`](examples/release-v1.0/actionlint.log) passed; downloaded upstream checksum verified; optional unavailable ShellCheck/Pyflakes analyzers disabled |
| Go vulnerability scan | `govulncheck v1.9.0` scanned `./...` against the official database and reported [no vulnerabilities found](examples/release-v1.0/govulncheck.log) at scan time |
| Prospective public source secret scan | `Gitleaks v8.30.1`, verified archive checksum, redacted directory scan of 311 candidate files reported [no leaks](examples/release-v1.0/gitleaks-source-snapshot.log); a separate [321-file Git index snapshot](examples/release-v1.0/gitleaks-index-snapshot.log) also passed |

The clean-export CLI first sampled promotion at **5.333 s** and joint topology/client stability at **8.535 s** after fault orchestration began. The failed primary remained stopped, so redundancy was not restored. These are observations from one local run, not general recovery or durability guarantees. Independent Docker inventories after the full suite and CLI run found no managed containers, networks or volumes remaining.

Scanner success is scoped to its analyzed packages, patterns and database at that time. It does not establish absence of all secrets or vulnerabilities, or audit the Docker/Redis images and host software.

## Linux runtime boundary and CI gate

The Linux runner used official Go 1.26.9 and Docker CLI 24.0.6, a local Unix daemon socket, six CPUs and 3 GiB memory, dropped capabilities and no-new-privileges. It was not a privileged container. The [source manifest](examples/release-v1.0/linux-source-manifest.sha256) identifies the source and definitions that were tested.

The attempted full Docker suite from inside that runner failed at readiness because the published `127.0.0.1` endpoints were unreachable there. The [unmodified failed attempt](examples/release-v1.0/linux-integration-attempt.log) records connection refusals and its termination. This is **not a passing Linux Docker integration result**. The attempt used Docker Desktop's Linux VM from a macOS host, rather than a native Linux host. Docker documents its [Desktop VM and host forwarding boundary](https://docs.docker.com/desktop/features/networking/) and requires Desktop 4.34+ with an enabled setting for [Desktop host networking](https://docs.docker.com/engine/network/drivers/host/). No Desktop setting was changed to make the test pass.

The incompatible harness was stopped. Its exact managed run IDs were cleaned through `redis-lab cleanup`, and managed container/network/volume inventories were empty before the next host suite. The temporary owned validation image was subsequently removed. Native Linux Docker end-to-end behavior and hosted Ubuntu GitHub Actions remain **unverified**. A successful hosted workflow or native Linux `make integration` is a release gate; an amd64 cross-build does not satisfy it.

## Security, licensing and artifact review

Local Docker target discovery, fixed-argument subprocesses, resource ownership and isolation checks, allowlisted redirects, bounded traffic, independent deadlines and cleanup verification were reviewed against their adverse-case tests. The daemon and local host user remain trusted; inspection-to-action races, routed network exposure and filesystem quotas remain explicit limits. [SECURITY.md](../SECURITY.md) now distinguishes configured loopback binding from guaranteed isolation and documents Docker's [pre-28 localhost publication issue](https://docs.docker.com/engine/network/port-publishing/).

Project source remains MIT. [Third-party notices](../THIRD_PARTY_NOTICES.md) and complete verbatim upstream license texts cover all five linked Go modules, the Go runtime, YAML's mixed MIT/Apache terms and its NOTICE. Redis and Docker are external runtime dependencies and retain their own licenses. Binary distributions must include the project LICENSE, notices and `licenses/` directory.

Generated binaries, transient reports, profiler output, agent metadata and common credential files are ignored. Original raw examples are deliberate release evidence, not transient artifacts: all 171 existing files are byte-identical to the [preserved manifest](examples/release-v1.0/original-evidence-manifest.sha256). Historical configurations, failed checks, small-sample results and platform limitations were preserved. No existing experiment evidence was rewritten. Repository attributes retain raw evidence and upstream legal texts byte for byte across Git checkouts. All 171 original evidence hashes and the license bytes were checked in the staged Git index, in addition to the working tree.

## Publishing manually

The local preparation does not authorize publishing. When you choose to publish, create the empty repository and push the source to trigger hosted validation. Wait for passing native Linux Docker or hosted CI evidence and resolve any failures before creating the v1.0 tag or GitHub release. Perform the commands below yourself, replacing `OWNER` with the intended account, replacing `RUN_ID` with the workflow run ID, and confirming the repository visibility.

```sh
# Create an empty GitHub repository named redis-reliability-lab in the GitHub UI.
# Leave automatic README/license/.gitignore initialization disabled.
git remote add origin https://github.com/OWNER/redis-reliability-lab.git
git push -u origin main
# Inspect the hosted workflow triggered by that push; do not claim it passed prematurely.
gh run list --workflow ci.yml --limit 1
gh run watch RUN_ID --exit-status
# Only after green checks and review:
git tag -a v1.0.0 -m "Redis Reliability Lab v1.0.0"
git push origin v1.0.0
```

Create the GitHub release from tag `v1.0.0`, link this audit and the real example reports, and state the platform scope. For optional binary assets, build each intended OS/architecture, confirm `redis-lab version` reports `1.0.0`, and include `LICENSE`, `THIRD_PARTY_NOTICES.md` and the complete `licenses/` directory in each archive. Source-only publishing does not require uploading binaries.

Suggested description: **A Go/Docker lab for reproducible Redis Cluster fault injection, failover, replication repair and durability experiments, with verifiable reports.**

Suggested topics: `go`, `redis`, `redis-cluster`, `distributed-systems`, `fault-injection`, `reliability`, `docker`, `replication`, `durability`, `backend`.
