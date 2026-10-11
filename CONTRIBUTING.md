# Contributing

Keep changes focused on reproducible reliability experiments and the evidence they produce. A useful contribution states its initial conditions, fault semantics, deadlines, recovery criteria, correctness contract, and cleanup behavior.

## Development

Use Go 1.26 or newer with the pinned toolchain. Format Go changes and run:

```sh
go test -race ./...
go vet ./...
go build -trimpath -buildvcs=false ./cmd/redis-lab
```

For Docker or lifecycle changes, also run the real integration suite against a local Linux-container daemon:

```sh
make integration
```

Record exactly which checks ran. If Docker was unavailable, state that integration remains unverified. Avoid committing generated binaries, transient reports, profiling artifacts, or container state.

## Review expectations

- Preserve the managed-local-target boundary. Do not add arbitrary Redis endpoints, shell commands, remote daemons, or daemon-wide cleanup.
- Propagate cancellation and explicitly bound waits, network operations, and subprocesses.
- Keep recovery predicates and metric calculations testable independently of Docker.
- Test failure and timeout behavior when changes affect it; do not replace real integration evidence with fakes.
- Document measurement limitations and avoid treating one observed result as a universal guarantee.
- Keep configuration and report schema changes deliberate, versioned when incompatible, and reflected in the example and documentation.
- Use only public technologies and synthetic data; include no credentials or employer-specific material.

Pull requests should explain the concrete behavior change, its motivation, and validation evidence. Additional faults should be introduced one complete, safely cleaned-up experiment at a time.
