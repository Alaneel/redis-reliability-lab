# ADR 0006: Full-node isolation through managed Docker endpoint removal

Status: accepted; verified by real CLI run `bbba656458a3cb89130c16de`

## Context

A process crash cannot show a healthy Redis process becoming unreachable to peers. Network injection must preserve the local ownership boundary without privileged containers, host firewall manipulation, or arbitrary commands. The existing observer and cluster client use fixed private-to-loopback address maps.

## Decision

Add `primary-isolate` using fixed Docker network disconnect/reconnect commands against the exact tracked primary and run bridge. Describe the actual fault as bidirectional whole-node isolation affecting peers and host clients. Preserve the Redis process and observe it through bounded local administrative PING/INFO commands.

Provision isolation experiments with explicit private IPAM and fixed ephemeral loopback ports, retaining original endpoints and isolation capability in creation labels. Capture a receipt before disconnect. Reconnect only the same container and process at the same IP and port; reject contradictory inspection results and retain command errors even when the side effect completed.

Allow cleanup of a detached node only within its verified isolation-capable environment. Cleanup does not require successful reconnect. Restore and cleanup use independent bounded contexts after cancellation. Preserve all unrelated resources and unknown network members.

## Consequences

The mechanism is narrow and reproducible without a firewall layer. It cannot model selective one-way links or writes from a minority-side client. Explicit subnet candidates can collide, and ephemeral port selection has a bind race; both fail visibly within bounded provisioning and cleanup. Connectivity restoration, cluster reconciliation, replication redundancy, service recovery, and data verification remain separate observations.

The original primary crash provisioning path remains unchanged. Unit tests verify the state and safety boundaries. A real CLI trial observed the unchanged Redis process while detached, same-endpoint restoration, all-six-node reconciliation and redundancy, 1389 acknowledged writes verified without errors, and successful cleanup. See the verbatim [report](../examples/primary-isolation/report.md) and [Phase 2 validation](../phase2-validation.md); one execution does not establish universal recovery or data-preservation behavior.
