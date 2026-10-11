# Third-party notices

Redis Reliability Lab source is covered by the [MIT license](LICENSE). This does not replace the licenses of its dependencies. The table below covers the external Go modules linked by `cmd/redis-lab` at the pinned versions, plus the Go runtime and standard library used by compiled binaries. Complete upstream license texts and YAML's NOTICE are retained verbatim in [licenses/](licenses/).

| Component | Version | License and notice |
|---|---|---|
| go-redis | v9.23.0 | [BSD-2-Clause](licenses/go-redis-BSD-2-Clause.txt) |
| go-yaml | v3.0.5 | [Mixed Apache-2.0/MIT terms](licenses/go-yaml-LICENSE.txt), [Apache-2.0 full text](licenses/Apache-2.0.txt), [NOTICE](licenses/go-yaml-NOTICE.txt) |
| xxhash | v2.3.0 | [MIT](licenses/xxhash-MIT.txt) |
| atomic | v1.12.0 | [MIT](licenses/atomic-MIT.txt) |
| golang.org/x/sys | v0.48.0 | [BSD-3-Clause](licenses/x-sys-BSD-3-Clause.txt) |
| Go runtime and standard library | go1.26.9 | [BSD-3-Clause](licenses/Go-BSD-3-Clause.txt), [patent grant](licenses/Go-PATENTS.txt) |

If distributing a compiled Redis Lab binary, include this file, the project LICENSE, and the `licenses/` directory. Test-only module dependencies, the Docker CLI/daemon, compiler tools and operating-system libraries remain separately licensed and are not bundled by this repository. `go.mod` and `go.sum` retain the dependency graph and checksums; the release audit records the actual linked-module inventory.

## External Redis and Docker runtime

The lab downloads and invokes a separate, unmodified official Redis container image. It does not redistribute Redis or claim that Redis is covered by this project's MIT license. Redis 8 and its bundled components have separate licensing choices described by [Redis's official licensing page](https://redis.io/legal/licenses/). Users must review those terms for their own use or redistribution. Docker and the optional release-validation tools also retain their respective upstream licenses.
