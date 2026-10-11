# ADR 0009: Freeze the ledger before an all-process persistence fault

Status: accepted.

An individual Redis crash leaves live replicas that can restore volatile values. Fresh diagnostic SETs can also repopulate keys, while acknowledgement or explicit fsync commands after traffic drain can improve persistence before the intended crash. Those paths would make a persistence result difficult to interpret.

The all-process experiment observes quiet runtime settings and files before baseline, then takes an end-baseline read-only topology snapshot. It stops and drains traffic and freezes the complete successful-SET ledger. It kills the original replicas before the original primaries and requires fresh ownership-checked proof that all six processes are stopped before restarting any. Automatic Docker restart is disabled. Containers, endpoints, Redis Cluster identities and disposable managed volumes are retained.

Restoration requires unchanged owners and replica assignments, followed by fresh quiet runtime/file observations. The runner reads every frozen acknowledged value before any diagnostic SET/GET or acknowledgement-policy probe. Missing, corrupt, unreadable and unverified results remain distinct. `complete` describes verification coverage and can coexist with missing values. Physical restart, topology, redundancy, data and persistence are independent verdicts; cleanup must also be verified.

SET success followed by short, failed or unattempted WAIT/WAITAOF remains in the full ledger. A command-compatibility probe, including WAITAOF's fsync request, must precede baseline so it cannot flush frozen workload writes. Subset validation supplements the full ledger.

This chooses a reproducible process-crash boundary rather than a power-loss model. The host, Docker VM, storage and page cache keep running. Drain and sequential kill/inspection intervals are recorded because background fsync can progress. Post-restart replication may distribute one persisted shard copy to its replica; preservation does not prove each of six AOF file sets independently persisted every write. Memory is the negative control and must have a nonempty frozen ledger and observed absence of value-persistence files.

Reports retain the raw evidence and separate pre-diagnostic validation from later diagnostic service timing. No workload is offered during fault/restoration, so these timings have no workload availability denominator. See [durability semantics](../durability-semantics.md) and [validation evidence](../phase3-validation.md).
