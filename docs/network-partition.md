# Primary network isolation

`primary-isolate` disconnects one selected primary container from its sole run-owned Docker bridge, then reconnects that same running process. It creates a **bidirectional full-node isolation**: Redis client traffic, replication traffic, and cluster-bus traffic cannot use the removed interface. The other five nodes retain their original network. Host clients also lose access to the isolated node's published port. This scenario does not create a selective or one-way partition, or a client-reachable minority write island. [Docker disconnect reference](https://docs.docker.com/reference/cli/docker/network/disconnect/)

## Running the experiment

```sh
go build -trimpath -buildvcs=false -o bin/redis-lab ./cmd/redis-lab
bin/redis-lab validate experiments/primary-isolation.yaml
bin/redis-lab run experiments/primary-isolation.yaml
```

The example uses a one-second Redis node timeout, an eight-second isolation hold, and separate bounded restoration and reconciliation phases. Failure detection and promotion remain observed outcomes; choosing these durations does not guarantee promotion. A nonzero exit and its JSON/Markdown reports are valid evidence for failed, partial, or aborted runs.

**Executed evidence:** Real CLI run `bbba656458a3cb89130c16de` exited successfully with recovery and verified cleanup. Three detached observations returned local PONG with unchanged container PID and Redis `run_id`; a subsequent connected observation retained those identities. Promotion was first observed 3442.663 ms after fault orchestration began, and jointly sustained service recovery at 5852.771 ms. Fault orchestration took 1363.464 ms, restoration 1286.373 ms, and stable reconciliation 1238.500 ms after restoration completed. All six nodes restored replication redundancy, and all 1389 acknowledged writes were checked with no missing, corrupt, unreadable, or unverified records. These are observations from one trial on Docker 24.0.6, not general guarantees. The original [JSON report](examples/primary-isolation/report.json) and [Markdown report](examples/primary-isolation/report.md) are copied verbatim. Further tests and executed results belong in [Phase 2 validation](phase2-validation.md).

## Preserving identities during restoration

Isolation experiments use an explicit private `/24` subnet selected from a bounded internal candidate pool. Docker rejects overlaps with existing allocations; the lab retries only that definite IPAM rejection. Each node receives a fixed private IP and a selected ephemeral `127.0.0.1` host port. These settings are confined to isolation experiments; the existing primary crash path keeps its default provisioning behavior.

The explicit subnet permits `docker network connect --ip` to restore the original IP. Fixed port bindings are needed because Docker 24 releases bridge port mappings when disconnecting and allocates them again when reconnecting. The client and observer retain their original address maps, so any IP or port change fails restoration. [Docker connect reference](https://docs.docker.com/reference/cli/docker/network/connect/), [Docker 24.0.6 network implementation](https://raw.githubusercontent.com/moby/moby/v24.0.6/daemon/container_operations.go), [Docker 24.0.6 bridge implementation](https://raw.githubusercontent.com/moby/moby/v24.0.6/libnetwork/drivers/bridge/bridge.go)

Selecting a free host port and having Docker bind it cannot be atomic across processes. A competing bind produces a visible provisioning failure followed by exact owned cleanup. The lab does not take ports or remove networks owned by another application.

## Liveness, reachability, and recovery

The lab observes the isolated Redis process through fixed, bounded `docker exec` commands against the container's own loopback interface: `PING`, `INFO server`, `CLUSTER INFO`, `CLUSTER NODES`, and `INFO replication`. This administrative observation path does not carry workload writes and does not restore peer communication. Each observation records Docker running state, PID and start time, Redis PID and `run_id`, attachment state, raw Redis topology/replication output, and observation start/end timestamps.

Container running status alone is insufficient. Unchanged Docker process identity, unchanged Redis process identity, and a local PONG support process survival at the observed times. They do not establish client availability or continuous cluster health between samples. Majority-side promotion, usable slot coverage, direct client operation success, and sustained service recovery are measured independently.

Reconnecting the interface is also insufficient for full recovery. Topology reconciliation must establish agreement across all six nodes, the former primary's replica relationship to the promoted owner, healthy replication links, and restored redundancy. Final immutable acknowledged-write validation remains a separate condition. Asynchronous Redis replication can lose acknowledged writes; record the observed result rather than infer a universal loss-free property. See [Redis Cluster specification](https://redis.io/docs/latest/operate/oss_and_stack/reference/cluster-spec/) and [failover interpretation](redis-failover.md).

The receipt preserves the Docker disconnect command's start and completion, the first verified detached state, and restoration completion. Report fault orchestration also includes ownership checks and process observations. Use these separate timestamps when interpreting recovery duration; topology and liveness polling limit event-time precision.

## Failure handling and ownership

Before disconnect or reconnect, the manager verifies the full tracked container/network IDs, expected names, run and management labels, bridge settings, original endpoints, read-only root filesystem, absence of host bind mounts, and loopback-only publication. Restoration also requires the original Docker and Redis process identities. The fixed disconnect operation does not use `--force`; restoration never replaces a process or disconnects an unexpected network.

A receipt is captured before mutation. If a command times out or its response is lost, bounded inspection determines whether its side effect occurred while retaining the original error. The runner uses a fresh bounded context for deferred restoration after cancellation, then cleans resources independently.

Creation-time isolation capability and endpoint labels allow `cleanup <run-id>` to discover a running detached node after a hard stop. Cleanup removes only exact owned IDs and can remove that detached container directly without reconnecting it. A normal crash experiment's unexpectedly detached running container does not acquire this permission. Altered ownership, another attached network or changed port causes refusal. Restoration additionally refuses an occupied original IP or changed process identity. Cleanup can remove an exact owned detached container after a process change; it refuses network removal while unknown members remain. Images and unrelated Docker resources are preserved.

This mechanism uses Docker's existing local networking API. It introduces no privileged container, host firewall command, remote Docker endpoint, arbitrary Redis target, or arbitrary shell command. The managed bridge disables egress IP masquerading; this setting is not a firewall against explicit host routes.
