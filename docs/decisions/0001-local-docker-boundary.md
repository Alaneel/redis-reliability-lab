# ADR 0001: Restrict destructive operations to managed local Docker resources

**Status:** Accepted

## Context

The framework must be reproducible without infrastructure accounts and must not turn a malformed experiment into a production Redis fault. Container creation, killing, and removal require a clear ownership boundary.

## Decision

Use the installed Docker CLI through fixed argument lists. Accept only local Unix-socket daemons. Create one labelled bridge with egress IP masquerading disabled and six labelled, read-only containers per run, with temporary data and loopback-only client ports. The tested Docker 24.0.6 daemon did not create active published ports on `--internal` bridges, so that option could not support the host-side client; the selected bridge mode keeps host access while avoiding outbound address translation. This is not a firewall against the Docker host or explicit routes. Revalidate IDs, labels, expected names, and isolation before destructive operations. Expose no arbitrary Redis address or arbitrary command hook.

Return a partial inventory after failed provisioning; reconstruct a interrupted run through exact labels and expected names. Remove only exact IDs, then verify absence. Preserve images and unrelated resources.

## Consequences

Linux and macOS with Linux Docker containers are supported. Remote Docker and native Windows named pipes are excluded. Docker CLI output is parsed through JSON rather than display tables, and the Docker executable remains a prerequisite. Hard process termination can require explicit `cleanup <run-id>`. Altered resources are preserved for inspection instead of forcibly removed.
