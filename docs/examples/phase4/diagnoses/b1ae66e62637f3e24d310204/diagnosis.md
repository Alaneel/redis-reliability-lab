# Incident diagnosis

Incident: `b1ae66e62637f3e24d310204`  
Detection: **incident**  
Unique root cause: **abstained**  
Engine: `rules-2`

## Observed symptoms and diagnostic hypotheses

### client_disruption

The all-shard client probe failed.

Assessment: strongly_supported. Rule: `R5-observed-client-probe-failure`. Nodes: .

Supporting evidence: e000441, e000443, e000650, e000652, e000859, e000861, e001275, e002510.

Contrary or recovery observations: e001068, e001273, e001482, e001687, e001892, e002305.

Missing evidence: Independent process liveness and network-path observations; Per-request routing/cache/redirect traces.

Alternatives: Process or endpoint failure; Network or observer connectivity problem; Planned administrative change.

### client_latency_change

Terminal request-window p95 is at least five times an earlier error-free measured window and at least 20 ms. Workload mix and acknowledgement waits can explain this; it is not a calibrated anomaly probability.

Assessment: strongly_supported. Rule: `R15-measured-window-p95-5x-and-20ms-with-20-attempts`. Nodes: .

Supporting evidence: e000445, e000004, e000442, e000654, e000651, e000863, e000860.

Contrary or recovery observations: .

Missing evidence: Independent process liveness and network-path observations; Per-request routing/cache/redirect traces.

Alternatives: Process or endpoint failure; Network or observer connectivity problem; Planned administrative change.

### primary_unavailability

Node observation endpoint unavailable; the cause remains unresolved.

Assessment: strongly_supported. Rule: `R2-endpoint-unavailable-with-prior-role`. Nodes: node-1, node-5.

Supporting evidence: e000237, e000124, e000560, e000693, e001017, e000784, e001054, e001259, e001392, e001559, e001688, e001897, e001911, e002382.

Contrary or recovery observations: e000009, e000123, e000251, e000446, e000783, e001016, e001069, e001444, e001611, e001816, e002001, e002039, e002434.

Missing evidence: Independent process liveness and network-path observations; Per-request routing/cache/redirect traces.

Alternatives: Process or endpoint failure; Network or observer connectivity problem; Planned administrative change.

### replica_unavailability

Replica INFO reports its primary link down; process and network causes remain alternatives.

Assessment: strongly_supported. Rule: `R4-replica-link-down`. Nodes: node-5.

Supporting evidence: e000254, e000449, e000786.

Contrary or recovery observations: e000009, e000251, e000446, e000783, e001016, e001069, e001444, e001611, e001816, e002001, e002039, e002434.

Missing evidence: Independent process liveness and network-path observations; Per-request routing/cache/redirect traces.

Alternatives: Process or endpoint failure; Network or observer connectivity problem; Planned administrative change.

## Timeline

| First observed time | Event | Evidence | Timing |
|---|---|---|---|
| 2026-10-11T19:55:50.225559+08:00 | baseline_full_slot_coverage_observed | e000017, e000055, e000093, e000131, e000169, e000207 | all recorded views in this non-atomic sample report full healthy coverage |
| 2026-10-11T19:55:50.546014+08:00 | component_unavailability_observed | e000237, e000123 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:55:50.546014+08:00 | first_observed_primary_unavailability | e000237 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:55:50.546065+08:00 | first_observed_replica_unavailability | e000254 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:55:51.176954+08:00 | first_observed_client_disruption | e000441 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:55:51.224589+08:00 | first_observed_client_latency_change | e000445 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:55:52.425739+08:00 | cluster_failure_flag_observed | e000672, e000629 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:55:52.425778+08:00 | cluster_failure_flag_observed | e000716, e000583 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:55:52.425804+08:00 | cluster_failure_flag_observed | e000758, e000497 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:55:52.425809+08:00 | cluster_failure_flag_observed | e000800, e000463 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:55:52.425814+08:00 | cluster_failure_flag_observed | e000850, e000551 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:55:53.365922+08:00 | replica_promotion_observed | e001017, e000784 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:55:53.498116+08:00 | client_visible_recovery_observed | e001068, e000859 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:55:55.880843+08:00 | client_recovery_regression_observed | e002510, e002305 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:55:56.080706+08:00 | final_correctness_validation_observed | e002521 | first_sampled_observation; exact event time unknown |

## Recovery and uncertainty

- Last client probe successful=false (e002510)
- Final checked acknowledged-value population preserved; see /validation evidence. This is not a durability guarantee.
- Insufficient evidence to identify a unique root cause. Process failure, network disruption and observation-channel failure can be indistinguishable.

## Read-only remediation plans

- **client_disruption**: Collect independent evidence before changing system state. Evidence: e000441, e000443, e000650, e000652, e000859, e000861, e001275, e002510. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability. Expected effect: Resolve alternative explanations without executing remediation. Verify: Inspect independently collected process and network status; Compare all node views and client outcomes at aligned timestamps. Human approval required; no execution.
- **client_latency_change**: Collect independent evidence before changing system state. Evidence: e000445, e000004, e000442, e000654, e000651, e000863, e000860. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability. Expected effect: Resolve alternative explanations without executing remediation. Verify: Inspect independently collected process and network status; Compare all node views and client outcomes at aligned timestamps. Human approval required; no execution.
- **primary_unavailability**: Collect independent evidence before changing system state. Evidence: e000237, e000124, e000560, e000693, e001017, e000784, e001054, e001259, e001392, e001559, e001688, e001897, e001911, e002382. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability. Expected effect: Resolve alternative explanations without executing remediation. Verify: Inspect independently collected process and network status; Compare all node views and client outcomes at aligned timestamps. Human approval required; no execution.
- **replica_unavailability**: Assess restoration of a synchronized replica. Evidence: e000254, e000449, e000786. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups; Confirm an authoritative healthy primary and compatible former-node identity and data. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability. Expected effect: Resolve alternative explanations without executing remediation. Verify: Inspect independently collected process and network status; Compare all node views and client outcomes at aligned timestamps; Observe replica assignment, link up, synchronization complete and repeated healthy service probes. Human approval required; no execution.

## Evidence provenance

Original bytes SHA-256: `24feabe4e114282db315d0b5df5e5c7c6d4415203c84a6dfecd2f01c97b1a72d`. Input report schema: 2. References are JSON pointers in the original report.

| ID | Time | Node | Source | Signal | Status | Value | Reference |
|---|---|---|---|---|---|---|---|
| e000001 | 19:55:49.224589 |  | client_window | client.attempts | valid | `169` | /metrics/windows/0/attempts |
| e000002 | 19:55:49.224589 |  | client_window | client.errors | valid | `0` | /metrics/windows/0/errors |
| e000003 | 19:55:49.224589 |  | client_window | client.successes | valid | `169` | /metrics/windows/0/successes |
| e000004 | 19:55:49.224589 |  | client_window | client.p95_ms | valid | `1.380029` | /metrics/windows/0/p95_ms |
| e000005 | 19:55:50.224589 |  | client_window | client.attempts | valid | `170` | /metrics/windows/1/attempts |
| e000006 | 19:55:50.224589 |  | client_window | client.errors | valid | `0` | /metrics/windows/1/errors |
| e000007 | 19:55:50.224589 |  | client_window | client.successes | valid | `170` | /metrics/windows/1/successes |
| e000008 | 19:55:50.224589 |  | client_window | client.p95_ms | valid | `0.731856` | /metrics/windows/1/p95_ms |
| e000009 | 19:55:50.225517 | node-5 | node_client | reachable | valid | `true` | /initial/views/4/reachable |
| e000010 | 19:55:50.225517 | node-5 | info_replication | replication.role | valid | `"slave"` | /initial/views/4/replication/role |
| e000011 | 19:55:50.225517 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /initial/views/4/replication/connected_slaves |
| e000012 | 19:55:50.225517 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /initial/views/4/replication/master_link_status |
| e000013 | 19:55:50.225517 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /initial/views/4/replication/master_sync_in_progress |
| e000014 | 19:55:50.225517 | node-5 | info_replication | replication.master_repl_offset | valid | `12417` | /initial/views/4/replication/master_repl_offset |
| e000015 | 19:55:50.225517 | node-5 | info_replication | replication.slave_repl_offset | valid | `12417` | /initial/views/4/replication/slave_repl_offset |
| e000016 | 19:55:50.225517 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /initial/views/4/redis_id |
| e000017 | 19:55:50.225517 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/4/coverage |
| e000018 | 19:55:50.225517 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/4/nodes/0/flags |
| e000019 | 19:55:50.225517 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /initial/views/4/nodes/0/primary_id |
| e000020 | 19:55:50.225517 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /initial/views/4/nodes/0/slots |
| e000021 | 19:55:50.225517 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/0/link_state |
| e000022 | 19:55:50.225517 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/4/nodes/1/flags |
| e000023 | 19:55:50.225517 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /initial/views/4/nodes/1/primary_id |
| e000024 | 19:55:50.225517 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/4/nodes/1/slots |
| e000025 | 19:55:50.225517 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/1/link_state |
| e000026 | 19:55:50.225517 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/4/nodes/2/flags |
| e000027 | 19:55:50.225517 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /initial/views/4/nodes/2/primary_id |
| e000028 | 19:55:50.225517 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/4/nodes/2/slots |
| e000029 | 19:55:50.225517 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/2/link_state |
| e000030 | 19:55:50.225517 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /initial/views/4/nodes/3/flags |
| e000031 | 19:55:50.225517 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /initial/views/4/nodes/3/primary_id |
| e000032 | 19:55:50.225517 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /initial/views/4/nodes/3/slots |
| e000033 | 19:55:50.225517 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/3/link_state |
| e000034 | 19:55:50.225517 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/4/nodes/4/flags |
| e000035 | 19:55:50.225517 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /initial/views/4/nodes/4/primary_id |
| e000036 | 19:55:50.225517 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /initial/views/4/nodes/4/slots |
| e000037 | 19:55:50.225517 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/4/link_state |
| e000038 | 19:55:50.225517 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/4/nodes/5/flags |
| e000039 | 19:55:50.225517 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /initial/views/4/nodes/5/primary_id |
| e000040 | 19:55:50.225517 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/4/nodes/5/slots |
| e000041 | 19:55:50.225517 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/5/link_state |
| e000042 | 19:55:50.225517 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/4/persistence/aof_enabled |
| e000043 | 19:55:50.225517 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/4/persistence/loading |
| e000044 | 19:55:50.225517 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/4/persistence/aof_last_write_status |
| e000045 | 19:55:50.225517 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/4/persistence/aof_rewrite_in_progress |
| e000046 | 19:55:50.225517 | node-5 | admin | process_live | missing_from_schema | `` | /initial/views/4/process_live |
| e000047 | 19:55:50.225522 | node-4 | node_client | reachable | valid | `true` | /initial/views/3/reachable |
| e000048 | 19:55:50.225522 | node-4 | info_replication | replication.role | valid | `"slave"` | /initial/views/3/replication/role |
| e000049 | 19:55:50.225522 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /initial/views/3/replication/connected_slaves |
| e000050 | 19:55:50.225522 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /initial/views/3/replication/master_link_status |
| e000051 | 19:55:50.225522 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /initial/views/3/replication/master_sync_in_progress |
| e000052 | 19:55:50.225522 | node-4 | info_replication | replication.master_repl_offset | valid | `12197` | /initial/views/3/replication/master_repl_offset |
| e000053 | 19:55:50.225522 | node-4 | info_replication | replication.slave_repl_offset | valid | `12197` | /initial/views/3/replication/slave_repl_offset |
| e000054 | 19:55:50.225522 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /initial/views/3/redis_id |
| e000055 | 19:55:50.225522 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/3/coverage |
| e000056 | 19:55:50.225522 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/3/nodes/0/flags |
| e000057 | 19:55:50.225522 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /initial/views/3/nodes/0/primary_id |
| e000058 | 19:55:50.225522 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/3/nodes/0/slots |
| e000059 | 19:55:50.225522 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/0/link_state |
| e000060 | 19:55:50.225522 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /initial/views/3/nodes/1/flags |
| e000061 | 19:55:50.225522 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /initial/views/3/nodes/1/primary_id |
| e000062 | 19:55:50.225522 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /initial/views/3/nodes/1/slots |
| e000063 | 19:55:50.225522 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/1/link_state |
| e000064 | 19:55:50.225522 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/3/nodes/2/flags |
| e000065 | 19:55:50.225522 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /initial/views/3/nodes/2/primary_id |
| e000066 | 19:55:50.225522 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /initial/views/3/nodes/2/slots |
| e000067 | 19:55:50.225522 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/2/link_state |
| e000068 | 19:55:50.225522 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/3/nodes/3/flags |
| e000069 | 19:55:50.225522 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /initial/views/3/nodes/3/primary_id |
| e000070 | 19:55:50.225522 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /initial/views/3/nodes/3/slots |
| e000071 | 19:55:50.225522 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/3/link_state |
| e000072 | 19:55:50.225522 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/3/nodes/4/flags |
| e000073 | 19:55:50.225522 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /initial/views/3/nodes/4/primary_id |
| e000074 | 19:55:50.225522 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/3/nodes/4/slots |
| e000075 | 19:55:50.225522 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/4/link_state |
| e000076 | 19:55:50.225522 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/3/nodes/5/flags |
| e000077 | 19:55:50.225522 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /initial/views/3/nodes/5/primary_id |
| e000078 | 19:55:50.225522 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/3/nodes/5/slots |
| e000079 | 19:55:50.225522 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/5/link_state |
| e000080 | 19:55:50.225522 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/3/persistence/aof_enabled |
| e000081 | 19:55:50.225522 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/3/persistence/loading |
| e000082 | 19:55:50.225522 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/3/persistence/aof_last_write_status |
| e000083 | 19:55:50.225522 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/3/persistence/aof_rewrite_in_progress |
| e000084 | 19:55:50.225522 | node-4 | admin | process_live | missing_from_schema | `` | /initial/views/3/process_live |
| e000085 | 19:55:50.225527 | node-6 | node_client | reachable | valid | `true` | /initial/views/5/reachable |
| e000086 | 19:55:50.225527 | node-6 | info_replication | replication.role | valid | `"slave"` | /initial/views/5/replication/role |
| e000087 | 19:55:50.225527 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /initial/views/5/replication/connected_slaves |
| e000088 | 19:55:50.225527 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /initial/views/5/replication/master_link_status |
| e000089 | 19:55:50.225527 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /initial/views/5/replication/master_sync_in_progress |
| e000090 | 19:55:50.225527 | node-6 | info_replication | replication.master_repl_offset | valid | `12419` | /initial/views/5/replication/master_repl_offset |
| e000091 | 19:55:50.225527 | node-6 | info_replication | replication.slave_repl_offset | valid | `12419` | /initial/views/5/replication/slave_repl_offset |
| e000092 | 19:55:50.225527 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /initial/views/5/redis_id |
| e000093 | 19:55:50.225527 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/5/coverage |
| e000094 | 19:55:50.225527 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/5/nodes/0/flags |
| e000095 | 19:55:50.225527 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /initial/views/5/nodes/0/primary_id |
| e000096 | 19:55:50.225527 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /initial/views/5/nodes/0/slots |
| e000097 | 19:55:50.225527 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/0/link_state |
| e000098 | 19:55:50.225527 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/5/nodes/1/flags |
| e000099 | 19:55:50.225527 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /initial/views/5/nodes/1/primary_id |
| e000100 | 19:55:50.225527 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/5/nodes/1/slots |
| e000101 | 19:55:50.225527 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/1/link_state |
| e000102 | 19:55:50.225527 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/5/nodes/2/flags |
| e000103 | 19:55:50.225527 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /initial/views/5/nodes/2/primary_id |
| e000104 | 19:55:50.225527 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/5/nodes/2/slots |
| e000105 | 19:55:50.225527 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/2/link_state |
| e000106 | 19:55:50.225527 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/5/nodes/3/flags |
| e000107 | 19:55:50.225527 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /initial/views/5/nodes/3/primary_id |
| e000108 | 19:55:50.225527 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/5/nodes/3/slots |
| e000109 | 19:55:50.225527 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/3/link_state |
| e000110 | 19:55:50.225527 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /initial/views/5/nodes/4/flags |
| e000111 | 19:55:50.225527 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /initial/views/5/nodes/4/primary_id |
| e000112 | 19:55:50.225527 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /initial/views/5/nodes/4/slots |
| e000113 | 19:55:50.225527 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/4/link_state |
| e000114 | 19:55:50.225527 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/5/nodes/5/flags |
| e000115 | 19:55:50.225527 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /initial/views/5/nodes/5/primary_id |
| e000116 | 19:55:50.225527 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /initial/views/5/nodes/5/slots |
| e000117 | 19:55:50.225527 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/5/link_state |
| e000118 | 19:55:50.225527 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/5/persistence/aof_enabled |
| e000119 | 19:55:50.225527 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/5/persistence/loading |
| e000120 | 19:55:50.225527 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/5/persistence/aof_last_write_status |
| e000121 | 19:55:50.225527 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/5/persistence/aof_rewrite_in_progress |
| e000122 | 19:55:50.225527 | node-6 | admin | process_live | missing_from_schema | `` | /initial/views/5/process_live |
| e000123 | 19:55:50.225544 | node-1 | node_client | reachable | valid | `true` | /initial/views/0/reachable |
| e000124 | 19:55:50.225544 | node-1 | info_replication | replication.role | valid | `"master"` | /initial/views/0/replication/role |
| e000125 | 19:55:50.225544 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /initial/views/0/replication/connected_slaves |
| e000126 | 19:55:50.225544 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /initial/views/0/replication/master_link_status |
| e000127 | 19:55:50.225544 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /initial/views/0/replication/master_sync_in_progress |
| e000128 | 19:55:50.225544 | node-1 | info_replication | replication.master_repl_offset | valid | `12417` | /initial/views/0/replication/master_repl_offset |
| e000129 | 19:55:50.225544 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /initial/views/0/replication/slave_repl_offset |
| e000130 | 19:55:50.225544 | node-1 | cluster_nodes | redis_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /initial/views/0/redis_id |
| e000131 | 19:55:50.225544 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/0/coverage |
| e000132 | 19:55:50.225544 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-1 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/0/nodes/0/flags |
| e000133 | 19:55:50.225544 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /initial/views/0/nodes/0/primary_id |
| e000134 | 19:55:50.225544 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/0/nodes/0/slots |
| e000135 | 19:55:50.225544 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/0/link_state |
| e000136 | 19:55:50.225544 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/0/nodes/1/flags |
| e000137 | 19:55:50.225544 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-1 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /initial/views/0/nodes/1/primary_id |
| e000138 | 19:55:50.225544 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /initial/views/0/nodes/1/slots |
| e000139 | 19:55:50.225544 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/1/link_state |
| e000140 | 19:55:50.225544 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/0/nodes/2/flags |
| e000141 | 19:55:50.225544 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-1 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /initial/views/0/nodes/2/primary_id |
| e000142 | 19:55:50.225544 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /initial/views/0/nodes/2/slots |
| e000143 | 19:55:50.225544 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/2/link_state |
| e000144 | 19:55:50.225544 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/0/nodes/3/flags |
| e000145 | 19:55:50.225544 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-1 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /initial/views/0/nodes/3/primary_id |
| e000146 | 19:55:50.225544 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /initial/views/0/nodes/3/slots |
| e000147 | 19:55:50.225544 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/3/link_state |
| e000148 | 19:55:50.225544 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-1 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/0/nodes/4/flags |
| e000149 | 19:55:50.225544 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /initial/views/0/nodes/4/primary_id |
| e000150 | 19:55:50.225544 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/0/nodes/4/slots |
| e000151 | 19:55:50.225544 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/4/link_state |
| e000152 | 19:55:50.225544 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-1 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /initial/views/0/nodes/5/flags |
| e000153 | 19:55:50.225544 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /initial/views/0/nodes/5/primary_id |
| e000154 | 19:55:50.225544 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/0/nodes/5/slots |
| e000155 | 19:55:50.225544 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/5/link_state |
| e000156 | 19:55:50.225544 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/0/persistence/aof_enabled |
| e000157 | 19:55:50.225544 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/0/persistence/loading |
| e000158 | 19:55:50.225544 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/0/persistence/aof_last_write_status |
| e000159 | 19:55:50.225544 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/0/persistence/aof_rewrite_in_progress |
| e000160 | 19:55:50.225544 | node-1 | admin | process_live | missing_from_schema | `` | /initial/views/0/process_live |
| e000161 | 19:55:50.225555 | node-2 | node_client | reachable | valid | `true` | /initial/views/1/reachable |
| e000162 | 19:55:50.225555 | node-2 | info_replication | replication.role | valid | `"master"` | /initial/views/1/replication/role |
| e000163 | 19:55:50.225555 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /initial/views/1/replication/connected_slaves |
| e000164 | 19:55:50.225555 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /initial/views/1/replication/master_link_status |
| e000165 | 19:55:50.225555 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /initial/views/1/replication/master_sync_in_progress |
| e000166 | 19:55:50.225555 | node-2 | info_replication | replication.master_repl_offset | valid | `12419` | /initial/views/1/replication/master_repl_offset |
| e000167 | 19:55:50.225555 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /initial/views/1/replication/slave_repl_offset |
| e000168 | 19:55:50.225555 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /initial/views/1/redis_id |
| e000169 | 19:55:50.225555 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/1/coverage |
| e000170 | 19:55:50.225555 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/1/nodes/0/flags |
| e000171 | 19:55:50.225555 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /initial/views/1/nodes/0/primary_id |
| e000172 | 19:55:50.225555 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/1/nodes/0/slots |
| e000173 | 19:55:50.225555 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/0/link_state |
| e000174 | 19:55:50.225555 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/1/nodes/1/flags |
| e000175 | 19:55:50.225555 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /initial/views/1/nodes/1/primary_id |
| e000176 | 19:55:50.225555 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /initial/views/1/nodes/1/slots |
| e000177 | 19:55:50.225555 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/1/link_state |
| e000178 | 19:55:50.225555 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/1/nodes/2/flags |
| e000179 | 19:55:50.225555 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /initial/views/1/nodes/2/primary_id |
| e000180 | 19:55:50.225555 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/1/nodes/2/slots |
| e000181 | 19:55:50.225555 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/2/link_state |
| e000182 | 19:55:50.225555 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /initial/views/1/nodes/3/flags |
| e000183 | 19:55:50.225555 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /initial/views/1/nodes/3/primary_id |
| e000184 | 19:55:50.225555 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/1/nodes/3/slots |
| e000185 | 19:55:50.225555 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/3/link_state |
| e000186 | 19:55:50.225555 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/1/nodes/4/flags |
| e000187 | 19:55:50.225555 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /initial/views/1/nodes/4/primary_id |
| e000188 | 19:55:50.225555 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /initial/views/1/nodes/4/slots |
| e000189 | 19:55:50.225555 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/4/link_state |
| e000190 | 19:55:50.225555 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/1/nodes/5/flags |
| e000191 | 19:55:50.225555 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /initial/views/1/nodes/5/primary_id |
| e000192 | 19:55:50.225555 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /initial/views/1/nodes/5/slots |
| e000193 | 19:55:50.225555 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/5/link_state |
| e000194 | 19:55:50.225555 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/1/persistence/aof_enabled |
| e000195 | 19:55:50.225555 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/1/persistence/loading |
| e000196 | 19:55:50.225555 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/1/persistence/aof_last_write_status |
| e000197 | 19:55:50.225555 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/1/persistence/aof_rewrite_in_progress |
| e000198 | 19:55:50.225555 | node-2 | admin | process_live | missing_from_schema | `` | /initial/views/1/process_live |
| e000199 | 19:55:50.225559 | node-3 | node_client | reachable | valid | `true` | /initial/views/2/reachable |
| e000200 | 19:55:50.225559 | node-3 | info_replication | replication.role | valid | `"master"` | /initial/views/2/replication/role |
| e000201 | 19:55:50.225559 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /initial/views/2/replication/connected_slaves |
| e000202 | 19:55:50.225559 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /initial/views/2/replication/master_link_status |
| e000203 | 19:55:50.225559 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /initial/views/2/replication/master_sync_in_progress |
| e000204 | 19:55:50.225559 | node-3 | info_replication | replication.master_repl_offset | valid | `12419` | /initial/views/2/replication/master_repl_offset |
| e000205 | 19:55:50.225559 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /initial/views/2/replication/slave_repl_offset |
| e000206 | 19:55:50.225559 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /initial/views/2/redis_id |
| e000207 | 19:55:50.225559 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/2/coverage |
| e000208 | 19:55:50.225559 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /initial/views/2/nodes/0/flags |
| e000209 | 19:55:50.225559 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /initial/views/2/nodes/0/primary_id |
| e000210 | 19:55:50.225559 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/2/nodes/0/slots |
| e000211 | 19:55:50.225559 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/0/link_state |
| e000212 | 19:55:50.225559 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/2/nodes/1/flags |
| e000213 | 19:55:50.225559 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /initial/views/2/nodes/1/primary_id |
| e000214 | 19:55:50.225559 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/2/nodes/1/slots |
| e000215 | 19:55:50.225559 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/1/link_state |
| e000216 | 19:55:50.225559 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/2/nodes/2/flags |
| e000217 | 19:55:50.225559 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /initial/views/2/nodes/2/primary_id |
| e000218 | 19:55:50.225559 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /initial/views/2/nodes/2/slots |
| e000219 | 19:55:50.225559 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/2/link_state |
| e000220 | 19:55:50.225559 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/2/nodes/3/flags |
| e000221 | 19:55:50.225559 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /initial/views/2/nodes/3/primary_id |
| e000222 | 19:55:50.225559 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/2/nodes/3/slots |
| e000223 | 19:55:50.225559 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/3/link_state |
| e000224 | 19:55:50.225559 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/2/nodes/4/flags |
| e000225 | 19:55:50.225559 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /initial/views/2/nodes/4/primary_id |
| e000226 | 19:55:50.225559 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /initial/views/2/nodes/4/slots |
| e000227 | 19:55:50.225559 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/4/link_state |
| e000228 | 19:55:50.225559 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/2/nodes/5/flags |
| e000229 | 19:55:50.225559 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /initial/views/2/nodes/5/primary_id |
| e000230 | 19:55:50.225559 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /initial/views/2/nodes/5/slots |
| e000231 | 19:55:50.225559 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/5/link_state |
| e000232 | 19:55:50.225559 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/2/persistence/aof_enabled |
| e000233 | 19:55:50.225559 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/2/persistence/loading |
| e000234 | 19:55:50.225559 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/2/persistence/aof_last_write_status |
| e000235 | 19:55:50.225559 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/2/persistence/aof_rewrite_in_progress |
| e000236 | 19:55:50.225559 | node-3 | admin | process_live | missing_from_schema | `` | /initial/views/2/process_live |
| e000237 | 19:55:50.546014 | node-1 | node_client | reachable | valid | `false` | /samples/0/snapshot/views/0/reachable |
| e000238 | 19:55:50.546014 | node-1 | info_replication | replication.role | collection_failed | `` | /samples/0/snapshot/views/0/replication/role |
| e000239 | 19:55:50.546014 | node-1 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/0/snapshot/views/0/replication/connected_slaves |
| e000240 | 19:55:50.546014 | node-1 | info_replication | replication.master_link_status | collection_failed | `` | /samples/0/snapshot/views/0/replication/master_link_status |
| e000241 | 19:55:50.546014 | node-1 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/0/snapshot/views/0/replication/master_sync_in_progress |
| e000242 | 19:55:50.546014 | node-1 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/0/snapshot/views/0/replication/master_repl_offset |
| e000243 | 19:55:50.546014 | node-1 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/0/snapshot/views/0/replication/slave_repl_offset |
| e000244 | 19:55:50.546014 | node-1 | cluster_nodes | redis_id | collection_failed | `` | /samples/0/snapshot/views/0/redis_id |
| e000245 | 19:55:50.546014 | node-1 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/0/snapshot/views/0/coverage |
| e000246 | 19:55:50.546014 | node-1 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/0/snapshot/views/0/persistence/aof_enabled |
| e000247 | 19:55:50.546014 | node-1 | info_persistence | persistence.loading | collection_failed | `` | /samples/0/snapshot/views/0/persistence/loading |
| e000248 | 19:55:50.546014 | node-1 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/0/snapshot/views/0/persistence/aof_last_write_status |
| e000249 | 19:55:50.546014 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/0/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e000250 | 19:55:50.546014 | node-1 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/0/process_live |
| e000251 | 19:55:50.546065 | node-5 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/4/reachable |
| e000252 | 19:55:50.546065 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/0/snapshot/views/4/replication/role |
| e000253 | 19:55:50.546065 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/0/snapshot/views/4/replication/connected_slaves |
| e000254 | 19:55:50.546065 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/0/snapshot/views/4/replication/master_link_status |
| e000255 | 19:55:50.546065 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/0/snapshot/views/4/replication/master_sync_in_progress |
| e000256 | 19:55:50.546065 | node-5 | info_replication | replication.master_repl_offset | valid | `14193` | /samples/0/snapshot/views/4/replication/master_repl_offset |
| e000257 | 19:55:50.546065 | node-5 | info_replication | replication.slave_repl_offset | valid | `14193` | /samples/0/snapshot/views/4/replication/slave_repl_offset |
| e000258 | 19:55:50.546065 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /samples/0/snapshot/views/4/redis_id |
| e000259 | 19:55:50.546065 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/4/coverage |
| e000260 | 19:55:50.546065 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/4/nodes/0/flags |
| e000261 | 19:55:50.546065 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/0/snapshot/views/4/nodes/0/primary_id |
| e000262 | 19:55:50.546065 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/0/slots |
| e000263 | 19:55:50.546065 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/0/link_state |
| e000264 | 19:55:50.546065 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/4/nodes/1/flags |
| e000265 | 19:55:50.546065 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/1/primary_id |
| e000266 | 19:55:50.546065 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/4/nodes/1/slots |
| e000267 | 19:55:50.546065 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/1/link_state |
| e000268 | 19:55:50.546065 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/4/nodes/2/flags |
| e000269 | 19:55:50.546065 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/2/primary_id |
| e000270 | 19:55:50.546065 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/4/nodes/2/slots |
| e000271 | 19:55:50.546065 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/2/link_state |
| e000272 | 19:55:50.546065 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/0/snapshot/views/4/nodes/3/flags |
| e000273 | 19:55:50.546065 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/0/snapshot/views/4/nodes/3/primary_id |
| e000274 | 19:55:50.546065 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/3/slots |
| e000275 | 19:55:50.546065 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/3/link_state |
| e000276 | 19:55:50.546065 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/4/nodes/4/flags |
| e000277 | 19:55:50.546065 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/0/snapshot/views/4/nodes/4/primary_id |
| e000278 | 19:55:50.546065 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/4/slots |
| e000279 | 19:55:50.546065 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/4/link_state |
| e000280 | 19:55:50.546065 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/4/nodes/5/flags |
| e000281 | 19:55:50.546065 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/5/primary_id |
| e000282 | 19:55:50.546065 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/4/nodes/5/slots |
| e000283 | 19:55:50.546065 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/5/link_state |
| e000284 | 19:55:50.546065 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/4/persistence/aof_enabled |
| e000285 | 19:55:50.546065 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/4/persistence/loading |
| e000286 | 19:55:50.546065 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/4/persistence/aof_last_write_status |
| e000287 | 19:55:50.546065 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e000288 | 19:55:50.546065 | node-5 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/4/process_live |
| e000289 | 19:55:50.546084 | node-4 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/3/reachable |
| e000290 | 19:55:50.546084 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/0/snapshot/views/3/replication/role |
| e000291 | 19:55:50.546084 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/0/snapshot/views/3/replication/connected_slaves |
| e000292 | 19:55:50.546084 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/0/snapshot/views/3/replication/master_link_status |
| e000293 | 19:55:50.546084 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/0/snapshot/views/3/replication/master_sync_in_progress |
| e000294 | 19:55:50.546084 | node-4 | info_replication | replication.master_repl_offset | valid | `14639` | /samples/0/snapshot/views/3/replication/master_repl_offset |
| e000295 | 19:55:50.546084 | node-4 | info_replication | replication.slave_repl_offset | valid | `14639` | /samples/0/snapshot/views/3/replication/slave_repl_offset |
| e000296 | 19:55:50.546084 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /samples/0/snapshot/views/3/redis_id |
| e000297 | 19:55:50.546084 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/3/coverage |
| e000298 | 19:55:50.546084 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/3/nodes/0/flags |
| e000299 | 19:55:50.546084 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/0/primary_id |
| e000300 | 19:55:50.546084 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/3/nodes/0/slots |
| e000301 | 19:55:50.546084 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/0/link_state |
| e000302 | 19:55:50.546084 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/0/snapshot/views/3/nodes/1/flags |
| e000303 | 19:55:50.546084 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/0/snapshot/views/3/nodes/1/primary_id |
| e000304 | 19:55:50.546084 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/1/slots |
| e000305 | 19:55:50.546084 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/1/link_state |
| e000306 | 19:55:50.546084 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/3/nodes/2/flags |
| e000307 | 19:55:50.546084 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/0/snapshot/views/3/nodes/2/primary_id |
| e000308 | 19:55:50.546084 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/2/slots |
| e000309 | 19:55:50.546084 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/2/link_state |
| e000310 | 19:55:50.546084 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/3/nodes/3/flags |
| e000311 | 19:55:50.546084 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/0/snapshot/views/3/nodes/3/primary_id |
| e000312 | 19:55:50.546084 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/3/slots |
| e000313 | 19:55:50.546084 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/3/link_state |
| e000314 | 19:55:50.546084 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/3/nodes/4/flags |
| e000315 | 19:55:50.546084 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/4/primary_id |
| e000316 | 19:55:50.546084 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/3/nodes/4/slots |
| e000317 | 19:55:50.546084 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/4/link_state |
| e000318 | 19:55:50.546084 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/3/nodes/5/flags |
| e000319 | 19:55:50.546084 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/5/primary_id |
| e000320 | 19:55:50.546084 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/3/nodes/5/slots |
| e000321 | 19:55:50.546084 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/5/link_state |
| e000322 | 19:55:50.546084 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/3/persistence/aof_enabled |
| e000323 | 19:55:50.546084 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/3/persistence/loading |
| e000324 | 19:55:50.546084 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/3/persistence/aof_last_write_status |
| e000325 | 19:55:50.546084 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e000326 | 19:55:50.546084 | node-4 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/3/process_live |
| e000327 | 19:55:50.546095 | node-2 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/1/reachable |
| e000328 | 19:55:50.546095 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/0/snapshot/views/1/replication/role |
| e000329 | 19:55:50.546095 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/0/snapshot/views/1/replication/connected_slaves |
| e000330 | 19:55:50.546095 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/0/snapshot/views/1/replication/master_link_status |
| e000331 | 19:55:50.546095 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/0/snapshot/views/1/replication/master_sync_in_progress |
| e000332 | 19:55:50.546095 | node-2 | info_replication | replication.master_repl_offset | valid | `14861` | /samples/0/snapshot/views/1/replication/master_repl_offset |
| e000333 | 19:55:50.546095 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/0/snapshot/views/1/replication/slave_repl_offset |
| e000334 | 19:55:50.546095 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/0/snapshot/views/1/redis_id |
| e000335 | 19:55:50.546095 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/1/coverage |
| e000336 | 19:55:50.546095 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/1/nodes/0/flags |
| e000337 | 19:55:50.546095 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/0/primary_id |
| e000338 | 19:55:50.546095 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/1/nodes/0/slots |
| e000339 | 19:55:50.546095 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/0/link_state |
| e000340 | 19:55:50.546095 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/1/nodes/1/flags |
| e000341 | 19:55:50.546095 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/0/snapshot/views/1/nodes/1/primary_id |
| e000342 | 19:55:50.546095 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/1/slots |
| e000343 | 19:55:50.546095 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/1/link_state |
| e000344 | 19:55:50.546095 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/1/nodes/2/flags |
| e000345 | 19:55:50.546095 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/2/primary_id |
| e000346 | 19:55:50.546095 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/1/nodes/2/slots |
| e000347 | 19:55:50.546095 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/2/link_state |
| e000348 | 19:55:50.546095 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/0/snapshot/views/1/nodes/3/flags |
| e000349 | 19:55:50.546095 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/3/primary_id |
| e000350 | 19:55:50.546095 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/1/nodes/3/slots |
| e000351 | 19:55:50.546095 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/3/link_state |
| e000352 | 19:55:50.546095 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/1/nodes/4/flags |
| e000353 | 19:55:50.546095 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/0/snapshot/views/1/nodes/4/primary_id |
| e000354 | 19:55:50.546095 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/4/slots |
| e000355 | 19:55:50.546095 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/4/link_state |
| e000356 | 19:55:50.546095 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/1/nodes/5/flags |
| e000357 | 19:55:50.546095 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/0/snapshot/views/1/nodes/5/primary_id |
| e000358 | 19:55:50.546095 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/5/slots |
| e000359 | 19:55:50.546095 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/5/link_state |
| e000360 | 19:55:50.546095 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/1/persistence/aof_enabled |
| e000361 | 19:55:50.546095 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/1/persistence/loading |
| e000362 | 19:55:50.546095 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/1/persistence/aof_last_write_status |
| e000363 | 19:55:50.546095 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e000364 | 19:55:50.546095 | node-2 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/1/process_live |
| e000365 | 19:55:50.546156 | node-3 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/2/reachable |
| e000366 | 19:55:50.546156 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/0/snapshot/views/2/replication/role |
| e000367 | 19:55:50.546156 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/0/snapshot/views/2/replication/connected_slaves |
| e000368 | 19:55:50.546156 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/0/snapshot/views/2/replication/master_link_status |
| e000369 | 19:55:50.546156 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/0/snapshot/views/2/replication/master_sync_in_progress |
| e000370 | 19:55:50.546156 | node-3 | info_replication | replication.master_repl_offset | valid | `14639` | /samples/0/snapshot/views/2/replication/master_repl_offset |
| e000371 | 19:55:50.546156 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/0/snapshot/views/2/replication/slave_repl_offset |
| e000372 | 19:55:50.546156 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/0/snapshot/views/2/redis_id |
| e000373 | 19:55:50.546156 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/2/coverage |
| e000374 | 19:55:50.546156 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/0/snapshot/views/2/nodes/0/flags |
| e000375 | 19:55:50.546156 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/0/primary_id |
| e000376 | 19:55:50.546156 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/2/nodes/0/slots |
| e000377 | 19:55:50.546156 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/0/link_state |
| e000378 | 19:55:50.546156 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/2/nodes/1/flags |
| e000379 | 19:55:50.546156 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/1/primary_id |
| e000380 | 19:55:50.546156 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/2/nodes/1/slots |
| e000381 | 19:55:50.546156 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/1/link_state |
| e000382 | 19:55:50.546156 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/2/nodes/2/flags |
| e000383 | 19:55:50.546156 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/0/snapshot/views/2/nodes/2/primary_id |
| e000384 | 19:55:50.546156 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/2/slots |
| e000385 | 19:55:50.546156 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/2/link_state |
| e000386 | 19:55:50.546156 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/2/nodes/3/flags |
| e000387 | 19:55:50.546156 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/3/primary_id |
| e000388 | 19:55:50.546156 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/2/nodes/3/slots |
| e000389 | 19:55:50.546156 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/3/link_state |
| e000390 | 19:55:50.546156 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/2/nodes/4/flags |
| e000391 | 19:55:50.546156 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/0/snapshot/views/2/nodes/4/primary_id |
| e000392 | 19:55:50.546156 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/4/slots |
| e000393 | 19:55:50.546156 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/4/link_state |
| e000394 | 19:55:50.546156 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/2/nodes/5/flags |
| e000395 | 19:55:50.546156 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/0/snapshot/views/2/nodes/5/primary_id |
| e000396 | 19:55:50.546156 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/5/slots |
| e000397 | 19:55:50.546156 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/5/link_state |
| e000398 | 19:55:50.546156 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/2/persistence/aof_enabled |
| e000399 | 19:55:50.546156 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/2/persistence/loading |
| e000400 | 19:55:50.546156 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/2/persistence/aof_last_write_status |
| e000401 | 19:55:50.546156 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e000402 | 19:55:50.546156 | node-3 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/2/process_live |
| e000403 | 19:55:50.546232 | node-6 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/5/reachable |
| e000404 | 19:55:50.546232 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/0/snapshot/views/5/replication/role |
| e000405 | 19:55:50.546232 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/0/snapshot/views/5/replication/connected_slaves |
| e000406 | 19:55:50.546232 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/0/snapshot/views/5/replication/master_link_status |
| e000407 | 19:55:50.546232 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/0/snapshot/views/5/replication/master_sync_in_progress |
| e000408 | 19:55:50.546232 | node-6 | info_replication | replication.master_repl_offset | valid | `14861` | /samples/0/snapshot/views/5/replication/master_repl_offset |
| e000409 | 19:55:50.546232 | node-6 | info_replication | replication.slave_repl_offset | valid | `14861` | /samples/0/snapshot/views/5/replication/slave_repl_offset |
| e000410 | 19:55:50.546232 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /samples/0/snapshot/views/5/redis_id |
| e000411 | 19:55:50.546232 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/5/coverage |
| e000412 | 19:55:50.546232 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/5/nodes/0/flags |
| e000413 | 19:55:50.546232 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/0/snapshot/views/5/nodes/0/primary_id |
| e000414 | 19:55:50.546232 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/0/slots |
| e000415 | 19:55:50.546232 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/0/link_state |
| e000416 | 19:55:50.546232 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/5/nodes/1/flags |
| e000417 | 19:55:50.546232 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/1/primary_id |
| e000418 | 19:55:50.546232 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/5/nodes/1/slots |
| e000419 | 19:55:50.546232 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/1/link_state |
| e000420 | 19:55:50.546232 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/5/nodes/2/flags |
| e000421 | 19:55:50.546232 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/2/primary_id |
| e000422 | 19:55:50.546232 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/5/nodes/2/slots |
| e000423 | 19:55:50.546232 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/2/link_state |
| e000424 | 19:55:50.546232 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/5/nodes/3/flags |
| e000425 | 19:55:50.546232 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/3/primary_id |
| e000426 | 19:55:50.546232 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/5/nodes/3/slots |
| e000427 | 19:55:50.546232 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/3/link_state |
| e000428 | 19:55:50.546232 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/0/snapshot/views/5/nodes/4/flags |
| e000429 | 19:55:50.546232 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/0/snapshot/views/5/nodes/4/primary_id |
| e000430 | 19:55:50.546232 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/4/slots |
| e000431 | 19:55:50.546232 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/4/link_state |
| e000432 | 19:55:50.546232 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/5/nodes/5/flags |
| e000433 | 19:55:50.546232 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/0/snapshot/views/5/nodes/5/primary_id |
| e000434 | 19:55:50.546232 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/5/slots |
| e000435 | 19:55:50.546232 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/5/link_state |
| e000436 | 19:55:50.546232 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/5/persistence/aof_enabled |
| e000437 | 19:55:50.546232 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/5/persistence/loading |
| e000438 | 19:55:50.546232 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/5/persistence/aof_last_write_status |
| e000439 | 19:55:50.546232 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e000440 | 19:55:50.546232 | node-6 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/5/process_live |
| e000441 | 19:55:51.176954 |  | client_probe | client_success | valid | `false` | /samples/0/probe/successful |
| e000442 | 19:55:51.224589 |  | client_window | client.attempts | valid | `98` | /metrics/windows/2/attempts |
| e000443 | 19:55:51.224589 |  | client_window | client.errors | valid | `18` | /metrics/windows/2/errors |
| e000444 | 19:55:51.224589 |  | client_window | client.successes | valid | `80` | /metrics/windows/2/successes |
| e000445 | 19:55:51.224589 |  | client_window | client.p95_ms | valid | `500.524291` | /metrics/windows/2/p95_ms |
| e000446 | 19:55:51.487135 | node-5 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/4/reachable |
| e000447 | 19:55:51.487135 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/1/snapshot/views/4/replication/role |
| e000448 | 19:55:51.487135 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/1/snapshot/views/4/replication/connected_slaves |
| e000449 | 19:55:51.487135 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/1/snapshot/views/4/replication/master_link_status |
| e000450 | 19:55:51.487135 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/1/snapshot/views/4/replication/master_sync_in_progress |
| e000451 | 19:55:51.487135 | node-5 | info_replication | replication.master_repl_offset | valid | `14193` | /samples/1/snapshot/views/4/replication/master_repl_offset |
| e000452 | 19:55:51.487135 | node-5 | info_replication | replication.slave_repl_offset | valid | `14193` | /samples/1/snapshot/views/4/replication/slave_repl_offset |
| e000453 | 19:55:51.487135 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /samples/1/snapshot/views/4/redis_id |
| e000454 | 19:55:51.487135 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/4/coverage |
| e000455 | 19:55:51.487135 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/4/nodes/0/flags |
| e000456 | 19:55:51.487135 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/1/snapshot/views/4/nodes/0/primary_id |
| e000457 | 19:55:51.487135 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/0/slots |
| e000458 | 19:55:51.487135 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/0/link_state |
| e000459 | 19:55:51.487135 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/4/nodes/1/flags |
| e000460 | 19:55:51.487135 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/1/primary_id |
| e000461 | 19:55:51.487135 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/4/nodes/1/slots |
| e000462 | 19:55:51.487135 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/1/link_state |
| e000463 | 19:55:51.487135 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/4/nodes/2/flags |
| e000464 | 19:55:51.487135 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/2/primary_id |
| e000465 | 19:55:51.487135 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/4/nodes/2/slots |
| e000466 | 19:55:51.487135 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/2/link_state |
| e000467 | 19:55:51.487135 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/1/snapshot/views/4/nodes/3/flags |
| e000468 | 19:55:51.487135 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/1/snapshot/views/4/nodes/3/primary_id |
| e000469 | 19:55:51.487135 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/3/slots |
| e000470 | 19:55:51.487135 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/3/link_state |
| e000471 | 19:55:51.487135 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/4/nodes/4/flags |
| e000472 | 19:55:51.487135 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/1/snapshot/views/4/nodes/4/primary_id |
| e000473 | 19:55:51.487135 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/4/slots |
| e000474 | 19:55:51.487135 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/4/link_state |
| e000475 | 19:55:51.487135 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/4/nodes/5/flags |
| e000476 | 19:55:51.487135 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/5/primary_id |
| e000477 | 19:55:51.487135 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/4/nodes/5/slots |
| e000478 | 19:55:51.487135 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/5/link_state |
| e000479 | 19:55:51.487135 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/4/persistence/aof_enabled |
| e000480 | 19:55:51.487135 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/4/persistence/loading |
| e000481 | 19:55:51.487135 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/4/persistence/aof_last_write_status |
| e000482 | 19:55:51.487135 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e000483 | 19:55:51.487135 | node-5 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/4/process_live |
| e000484 | 19:55:51.487146 | node-3 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/2/reachable |
| e000485 | 19:55:51.487146 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/1/snapshot/views/2/replication/role |
| e000486 | 19:55:51.487146 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/1/snapshot/views/2/replication/connected_slaves |
| e000487 | 19:55:51.487146 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/1/snapshot/views/2/replication/master_link_status |
| e000488 | 19:55:51.487146 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/1/snapshot/views/2/replication/master_sync_in_progress |
| e000489 | 19:55:51.487146 | node-3 | info_replication | replication.master_repl_offset | valid | `18059` | /samples/1/snapshot/views/2/replication/master_repl_offset |
| e000490 | 19:55:51.487146 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/1/snapshot/views/2/replication/slave_repl_offset |
| e000491 | 19:55:51.487146 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/1/snapshot/views/2/redis_id |
| e000492 | 19:55:51.487146 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/2/coverage |
| e000493 | 19:55:51.487146 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/1/snapshot/views/2/nodes/0/flags |
| e000494 | 19:55:51.487146 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/0/primary_id |
| e000495 | 19:55:51.487146 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/2/nodes/0/slots |
| e000496 | 19:55:51.487146 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/0/link_state |
| e000497 | 19:55:51.487146 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/2/nodes/1/flags |
| e000498 | 19:55:51.487146 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/1/primary_id |
| e000499 | 19:55:51.487146 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/2/nodes/1/slots |
| e000500 | 19:55:51.487146 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/1/link_state |
| e000501 | 19:55:51.487146 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/2/nodes/2/flags |
| e000502 | 19:55:51.487146 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/1/snapshot/views/2/nodes/2/primary_id |
| e000503 | 19:55:51.487146 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/2/slots |
| e000504 | 19:55:51.487146 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/2/link_state |
| e000505 | 19:55:51.487146 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/2/nodes/3/flags |
| e000506 | 19:55:51.487146 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/3/primary_id |
| e000507 | 19:55:51.487146 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/2/nodes/3/slots |
| e000508 | 19:55:51.487146 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/3/link_state |
| e000509 | 19:55:51.487146 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/2/nodes/4/flags |
| e000510 | 19:55:51.487146 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/1/snapshot/views/2/nodes/4/primary_id |
| e000511 | 19:55:51.487146 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/4/slots |
| e000512 | 19:55:51.487146 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/4/link_state |
| e000513 | 19:55:51.487146 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/2/nodes/5/flags |
| e000514 | 19:55:51.487146 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/1/snapshot/views/2/nodes/5/primary_id |
| e000515 | 19:55:51.487146 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/5/slots |
| e000516 | 19:55:51.487146 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/5/link_state |
| e000517 | 19:55:51.487146 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/2/persistence/aof_enabled |
| e000518 | 19:55:51.487146 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/2/persistence/loading |
| e000519 | 19:55:51.487146 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/2/persistence/aof_last_write_status |
| e000520 | 19:55:51.487146 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e000521 | 19:55:51.487146 | node-3 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/2/process_live |
| e000522 | 19:55:51.487162 | node-4 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/3/reachable |
| e000523 | 19:55:51.487162 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/1/snapshot/views/3/replication/role |
| e000524 | 19:55:51.487162 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/1/snapshot/views/3/replication/connected_slaves |
| e000525 | 19:55:51.487162 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/1/snapshot/views/3/replication/master_link_status |
| e000526 | 19:55:51.487162 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/1/snapshot/views/3/replication/master_sync_in_progress |
| e000527 | 19:55:51.487162 | node-4 | info_replication | replication.master_repl_offset | valid | `18059` | /samples/1/snapshot/views/3/replication/master_repl_offset |
| e000528 | 19:55:51.487162 | node-4 | info_replication | replication.slave_repl_offset | valid | `18059` | /samples/1/snapshot/views/3/replication/slave_repl_offset |
| e000529 | 19:55:51.487162 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /samples/1/snapshot/views/3/redis_id |
| e000530 | 19:55:51.487162 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/3/coverage |
| e000531 | 19:55:51.487162 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/3/nodes/0/flags |
| e000532 | 19:55:51.487162 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/0/primary_id |
| e000533 | 19:55:51.487162 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/3/nodes/0/slots |
| e000534 | 19:55:51.487162 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/0/link_state |
| e000535 | 19:55:51.487162 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/1/snapshot/views/3/nodes/1/flags |
| e000536 | 19:55:51.487162 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/1/snapshot/views/3/nodes/1/primary_id |
| e000537 | 19:55:51.487162 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/1/slots |
| e000538 | 19:55:51.487162 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/1/link_state |
| e000539 | 19:55:51.487162 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/3/nodes/2/flags |
| e000540 | 19:55:51.487162 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/1/snapshot/views/3/nodes/2/primary_id |
| e000541 | 19:55:51.487162 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/2/slots |
| e000542 | 19:55:51.487162 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/2/link_state |
| e000543 | 19:55:51.487162 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/3/nodes/3/flags |
| e000544 | 19:55:51.487162 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/1/snapshot/views/3/nodes/3/primary_id |
| e000545 | 19:55:51.487162 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/3/slots |
| e000546 | 19:55:51.487162 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/3/link_state |
| e000547 | 19:55:51.487162 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/3/nodes/4/flags |
| e000548 | 19:55:51.487162 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/4/primary_id |
| e000549 | 19:55:51.487162 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/3/nodes/4/slots |
| e000550 | 19:55:51.487162 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/4/link_state |
| e000551 | 19:55:51.487162 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/3/nodes/5/flags |
| e000552 | 19:55:51.487162 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/5/primary_id |
| e000553 | 19:55:51.487162 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/3/nodes/5/slots |
| e000554 | 19:55:51.487162 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/5/link_state |
| e000555 | 19:55:51.487162 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/3/persistence/aof_enabled |
| e000556 | 19:55:51.487162 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/3/persistence/loading |
| e000557 | 19:55:51.487162 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/3/persistence/aof_last_write_status |
| e000558 | 19:55:51.487162 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e000559 | 19:55:51.487162 | node-4 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/3/process_live |
| e000560 | 19:55:51.487165 | node-1 | node_client | reachable | valid | `false` | /samples/1/snapshot/views/0/reachable |
| e000561 | 19:55:51.487165 | node-1 | info_replication | replication.role | collection_failed | `` | /samples/1/snapshot/views/0/replication/role |
| e000562 | 19:55:51.487165 | node-1 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/1/snapshot/views/0/replication/connected_slaves |
| e000563 | 19:55:51.487165 | node-1 | info_replication | replication.master_link_status | collection_failed | `` | /samples/1/snapshot/views/0/replication/master_link_status |
| e000564 | 19:55:51.487165 | node-1 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/1/snapshot/views/0/replication/master_sync_in_progress |
| e000565 | 19:55:51.487165 | node-1 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/1/snapshot/views/0/replication/master_repl_offset |
| e000566 | 19:55:51.487165 | node-1 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/1/snapshot/views/0/replication/slave_repl_offset |
| e000567 | 19:55:51.487165 | node-1 | cluster_nodes | redis_id | collection_failed | `` | /samples/1/snapshot/views/0/redis_id |
| e000568 | 19:55:51.487165 | node-1 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/1/snapshot/views/0/coverage |
| e000569 | 19:55:51.487165 | node-1 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/1/snapshot/views/0/persistence/aof_enabled |
| e000570 | 19:55:51.487165 | node-1 | info_persistence | persistence.loading | collection_failed | `` | /samples/1/snapshot/views/0/persistence/loading |
| e000571 | 19:55:51.487165 | node-1 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/1/snapshot/views/0/persistence/aof_last_write_status |
| e000572 | 19:55:51.487165 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/1/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e000573 | 19:55:51.487165 | node-1 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/0/process_live |
| e000574 | 19:55:51.487165 | node-2 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/1/reachable |
| e000575 | 19:55:51.487165 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/1/snapshot/views/1/replication/role |
| e000576 | 19:55:51.487165 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/1/snapshot/views/1/replication/connected_slaves |
| e000577 | 19:55:51.487165 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/1/snapshot/views/1/replication/master_link_status |
| e000578 | 19:55:51.487165 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/1/snapshot/views/1/replication/master_sync_in_progress |
| e000579 | 19:55:51.487165 | node-2 | info_replication | replication.master_repl_offset | valid | `18059` | /samples/1/snapshot/views/1/replication/master_repl_offset |
| e000580 | 19:55:51.487165 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/1/snapshot/views/1/replication/slave_repl_offset |
| e000581 | 19:55:51.487165 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/1/snapshot/views/1/redis_id |
| e000582 | 19:55:51.487165 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/1/coverage |
| e000583 | 19:55:51.487165 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/1/nodes/0/flags |
| e000584 | 19:55:51.487165 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/0/primary_id |
| e000585 | 19:55:51.487165 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/1/nodes/0/slots |
| e000586 | 19:55:51.487165 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/0/link_state |
| e000587 | 19:55:51.487165 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/1/nodes/1/flags |
| e000588 | 19:55:51.487165 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/1/snapshot/views/1/nodes/1/primary_id |
| e000589 | 19:55:51.487165 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/1/slots |
| e000590 | 19:55:51.487165 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/1/link_state |
| e000591 | 19:55:51.487165 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/1/nodes/2/flags |
| e000592 | 19:55:51.487165 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/2/primary_id |
| e000593 | 19:55:51.487165 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/1/nodes/2/slots |
| e000594 | 19:55:51.487165 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/2/link_state |
| e000595 | 19:55:51.487165 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/1/snapshot/views/1/nodes/3/flags |
| e000596 | 19:55:51.487165 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/3/primary_id |
| e000597 | 19:55:51.487165 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/1/nodes/3/slots |
| e000598 | 19:55:51.487165 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/3/link_state |
| e000599 | 19:55:51.487165 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/1/nodes/4/flags |
| e000600 | 19:55:51.487165 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/1/snapshot/views/1/nodes/4/primary_id |
| e000601 | 19:55:51.487165 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/4/slots |
| e000602 | 19:55:51.487165 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/4/link_state |
| e000603 | 19:55:51.487165 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/1/nodes/5/flags |
| e000604 | 19:55:51.487165 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/1/snapshot/views/1/nodes/5/primary_id |
| e000605 | 19:55:51.487165 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/5/slots |
| e000606 | 19:55:51.487165 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/5/link_state |
| e000607 | 19:55:51.487165 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/1/persistence/aof_enabled |
| e000608 | 19:55:51.487165 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/1/persistence/loading |
| e000609 | 19:55:51.487165 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/1/persistence/aof_last_write_status |
| e000610 | 19:55:51.487165 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e000611 | 19:55:51.487165 | node-2 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/1/process_live |
| e000612 | 19:55:51.487185 | node-6 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/5/reachable |
| e000613 | 19:55:51.487185 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/1/snapshot/views/5/replication/role |
| e000614 | 19:55:51.487185 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/1/snapshot/views/5/replication/connected_slaves |
| e000615 | 19:55:51.487185 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/1/snapshot/views/5/replication/master_link_status |
| e000616 | 19:55:51.487185 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/1/snapshot/views/5/replication/master_sync_in_progress |
| e000617 | 19:55:51.487185 | node-6 | info_replication | replication.master_repl_offset | valid | `18059` | /samples/1/snapshot/views/5/replication/master_repl_offset |
| e000618 | 19:55:51.487185 | node-6 | info_replication | replication.slave_repl_offset | valid | `18059` | /samples/1/snapshot/views/5/replication/slave_repl_offset |
| e000619 | 19:55:51.487185 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /samples/1/snapshot/views/5/redis_id |
| e000620 | 19:55:51.487185 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/5/coverage |
| e000621 | 19:55:51.487185 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/5/nodes/0/flags |
| e000622 | 19:55:51.487185 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/1/snapshot/views/5/nodes/0/primary_id |
| e000623 | 19:55:51.487185 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/0/slots |
| e000624 | 19:55:51.487185 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/0/link_state |
| e000625 | 19:55:51.487185 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/5/nodes/1/flags |
| e000626 | 19:55:51.487185 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/1/primary_id |
| e000627 | 19:55:51.487185 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/5/nodes/1/slots |
| e000628 | 19:55:51.487185 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/1/link_state |
| e000629 | 19:55:51.487185 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/5/nodes/2/flags |
| e000630 | 19:55:51.487185 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/2/primary_id |
| e000631 | 19:55:51.487185 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/5/nodes/2/slots |
| e000632 | 19:55:51.487185 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/2/link_state |
| e000633 | 19:55:51.487185 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/5/nodes/3/flags |
| e000634 | 19:55:51.487185 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/3/primary_id |
| e000635 | 19:55:51.487185 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/5/nodes/3/slots |
| e000636 | 19:55:51.487185 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/3/link_state |
| e000637 | 19:55:51.487185 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/1/snapshot/views/5/nodes/4/flags |
| e000638 | 19:55:51.487185 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/1/snapshot/views/5/nodes/4/primary_id |
| e000639 | 19:55:51.487185 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/4/slots |
| e000640 | 19:55:51.487185 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/4/link_state |
| e000641 | 19:55:51.487185 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/5/nodes/5/flags |
| e000642 | 19:55:51.487185 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/1/snapshot/views/5/nodes/5/primary_id |
| e000643 | 19:55:51.487185 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/5/slots |
| e000644 | 19:55:51.487185 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/5/link_state |
| e000645 | 19:55:51.487185 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/5/persistence/aof_enabled |
| e000646 | 19:55:51.487185 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/5/persistence/loading |
| e000647 | 19:55:51.487185 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/5/persistence/aof_last_write_status |
| e000648 | 19:55:51.487185 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e000649 | 19:55:51.487185 | node-6 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/5/process_live |
| e000650 | 19:55:52.115614 |  | client_probe | client_success | valid | `false` | /samples/1/probe/successful |
| e000651 | 19:55:52.224589 |  | client_window | client.attempts | valid | `102` | /metrics/windows/3/attempts |
| e000652 | 19:55:52.224589 |  | client_window | client.errors | valid | `33` | /metrics/windows/3/errors |
| e000653 | 19:55:52.224589 |  | client_window | client.successes | valid | `69` | /metrics/windows/3/successes |
| e000654 | 19:55:52.224589 |  | client_window | client.p95_ms | valid | `500.591125` | /metrics/windows/3/p95_ms |
| e000655 | 19:55:52.425739 | node-6 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/5/reachable |
| e000656 | 19:55:52.425739 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/2/snapshot/views/5/replication/role |
| e000657 | 19:55:52.425739 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/2/snapshot/views/5/replication/connected_slaves |
| e000658 | 19:55:52.425739 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/2/snapshot/views/5/replication/master_link_status |
| e000659 | 19:55:52.425739 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/2/snapshot/views/5/replication/master_sync_in_progress |
| e000660 | 19:55:52.425739 | node-6 | info_replication | replication.master_repl_offset | valid | `19703` | /samples/2/snapshot/views/5/replication/master_repl_offset |
| e000661 | 19:55:52.425739 | node-6 | info_replication | replication.slave_repl_offset | valid | `19703` | /samples/2/snapshot/views/5/replication/slave_repl_offset |
| e000662 | 19:55:52.425739 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /samples/2/snapshot/views/5/redis_id |
| e000663 | 19:55:52.425739 | node-6 | cluster_nodes | slot_coverage | valid | `false` | /samples/2/snapshot/views/5/coverage |
| e000664 | 19:55:52.425739 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/5/nodes/0/flags |
| e000665 | 19:55:52.425739 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/2/snapshot/views/5/nodes/0/primary_id |
| e000666 | 19:55:52.425739 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/0/slots |
| e000667 | 19:55:52.425739 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/0/link_state |
| e000668 | 19:55:52.425739 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/5/nodes/1/flags |
| e000669 | 19:55:52.425739 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/1/primary_id |
| e000670 | 19:55:52.425739 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/5/nodes/1/slots |
| e000671 | 19:55:52.425739 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/1/link_state |
| e000672 | 19:55:52.425739 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/2/snapshot/views/5/nodes/2/flags |
| e000673 | 19:55:52.425739 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/2/primary_id |
| e000674 | 19:55:52.425739 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/5/nodes/2/slots |
| e000675 | 19:55:52.425739 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/2/link_state |
| e000676 | 19:55:52.425739 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/5/nodes/3/flags |
| e000677 | 19:55:52.425739 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/3/primary_id |
| e000678 | 19:55:52.425739 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/5/nodes/3/slots |
| e000679 | 19:55:52.425739 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/3/link_state |
| e000680 | 19:55:52.425739 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/2/snapshot/views/5/nodes/4/flags |
| e000681 | 19:55:52.425739 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/2/snapshot/views/5/nodes/4/primary_id |
| e000682 | 19:55:52.425739 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/4/slots |
| e000683 | 19:55:52.425739 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/4/link_state |
| e000684 | 19:55:52.425739 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/5/nodes/5/flags |
| e000685 | 19:55:52.425739 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/2/snapshot/views/5/nodes/5/primary_id |
| e000686 | 19:55:52.425739 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/5/slots |
| e000687 | 19:55:52.425739 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/5/link_state |
| e000688 | 19:55:52.425739 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/5/persistence/aof_enabled |
| e000689 | 19:55:52.425739 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/5/persistence/loading |
| e000690 | 19:55:52.425739 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/5/persistence/aof_last_write_status |
| e000691 | 19:55:52.425739 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e000692 | 19:55:52.425739 | node-6 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/5/process_live |
| e000693 | 19:55:52.425742 | node-1 | node_client | reachable | valid | `false` | /samples/2/snapshot/views/0/reachable |
| e000694 | 19:55:52.425742 | node-1 | info_replication | replication.role | collection_failed | `` | /samples/2/snapshot/views/0/replication/role |
| e000695 | 19:55:52.425742 | node-1 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/2/snapshot/views/0/replication/connected_slaves |
| e000696 | 19:55:52.425742 | node-1 | info_replication | replication.master_link_status | collection_failed | `` | /samples/2/snapshot/views/0/replication/master_link_status |
| e000697 | 19:55:52.425742 | node-1 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/2/snapshot/views/0/replication/master_sync_in_progress |
| e000698 | 19:55:52.425742 | node-1 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/2/snapshot/views/0/replication/master_repl_offset |
| e000699 | 19:55:52.425742 | node-1 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/2/snapshot/views/0/replication/slave_repl_offset |
| e000700 | 19:55:52.425742 | node-1 | cluster_nodes | redis_id | collection_failed | `` | /samples/2/snapshot/views/0/redis_id |
| e000701 | 19:55:52.425742 | node-1 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/2/snapshot/views/0/coverage |
| e000702 | 19:55:52.425742 | node-1 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/2/snapshot/views/0/persistence/aof_enabled |
| e000703 | 19:55:52.425742 | node-1 | info_persistence | persistence.loading | collection_failed | `` | /samples/2/snapshot/views/0/persistence/loading |
| e000704 | 19:55:52.425742 | node-1 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/2/snapshot/views/0/persistence/aof_last_write_status |
| e000705 | 19:55:52.425742 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/2/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e000706 | 19:55:52.425742 | node-1 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/0/process_live |
| e000707 | 19:55:52.425778 | node-2 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/1/reachable |
| e000708 | 19:55:52.425778 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/2/snapshot/views/1/replication/role |
| e000709 | 19:55:52.425778 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/2/snapshot/views/1/replication/connected_slaves |
| e000710 | 19:55:52.425778 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/2/snapshot/views/1/replication/master_link_status |
| e000711 | 19:55:52.425778 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/2/snapshot/views/1/replication/master_sync_in_progress |
| e000712 | 19:55:52.425778 | node-2 | info_replication | replication.master_repl_offset | valid | `19703` | /samples/2/snapshot/views/1/replication/master_repl_offset |
| e000713 | 19:55:52.425778 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/2/snapshot/views/1/replication/slave_repl_offset |
| e000714 | 19:55:52.425778 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/2/snapshot/views/1/redis_id |
| e000715 | 19:55:52.425778 | node-2 | cluster_nodes | slot_coverage | valid | `false` | /samples/2/snapshot/views/1/coverage |
| e000716 | 19:55:52.425778 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/2/snapshot/views/1/nodes/0/flags |
| e000717 | 19:55:52.425778 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/0/primary_id |
| e000718 | 19:55:52.425778 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/1/nodes/0/slots |
| e000719 | 19:55:52.425778 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/0/link_state |
| e000720 | 19:55:52.425778 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/1/nodes/1/flags |
| e000721 | 19:55:52.425778 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/2/snapshot/views/1/nodes/1/primary_id |
| e000722 | 19:55:52.425778 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/1/slots |
| e000723 | 19:55:52.425778 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/1/link_state |
| e000724 | 19:55:52.425778 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/1/nodes/2/flags |
| e000725 | 19:55:52.425778 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/2/primary_id |
| e000726 | 19:55:52.425778 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/1/nodes/2/slots |
| e000727 | 19:55:52.425778 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/2/link_state |
| e000728 | 19:55:52.425778 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/2/snapshot/views/1/nodes/3/flags |
| e000729 | 19:55:52.425778 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/3/primary_id |
| e000730 | 19:55:52.425778 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/1/nodes/3/slots |
| e000731 | 19:55:52.425778 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/3/link_state |
| e000732 | 19:55:52.425778 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/1/nodes/4/flags |
| e000733 | 19:55:52.425778 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/2/snapshot/views/1/nodes/4/primary_id |
| e000734 | 19:55:52.425778 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/4/slots |
| e000735 | 19:55:52.425778 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/4/link_state |
| e000736 | 19:55:52.425778 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/1/nodes/5/flags |
| e000737 | 19:55:52.425778 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/2/snapshot/views/1/nodes/5/primary_id |
| e000738 | 19:55:52.425778 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/5/slots |
| e000739 | 19:55:52.425778 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/5/link_state |
| e000740 | 19:55:52.425778 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/1/persistence/aof_enabled |
| e000741 | 19:55:52.425778 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/1/persistence/loading |
| e000742 | 19:55:52.425778 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/1/persistence/aof_last_write_status |
| e000743 | 19:55:52.425778 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e000744 | 19:55:52.425778 | node-2 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/1/process_live |
| e000745 | 19:55:52.425804 | node-3 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/2/reachable |
| e000746 | 19:55:52.425804 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/2/snapshot/views/2/replication/role |
| e000747 | 19:55:52.425804 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/2/snapshot/views/2/replication/connected_slaves |
| e000748 | 19:55:52.425804 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/2/snapshot/views/2/replication/master_link_status |
| e000749 | 19:55:52.425804 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/2/snapshot/views/2/replication/master_sync_in_progress |
| e000750 | 19:55:52.425804 | node-3 | info_replication | replication.master_repl_offset | valid | `19703` | /samples/2/snapshot/views/2/replication/master_repl_offset |
| e000751 | 19:55:52.425804 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/2/snapshot/views/2/replication/slave_repl_offset |
| e000752 | 19:55:52.425804 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/2/snapshot/views/2/redis_id |
| e000753 | 19:55:52.425804 | node-3 | cluster_nodes | slot_coverage | valid | `false` | /samples/2/snapshot/views/2/coverage |
| e000754 | 19:55:52.425804 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/2/snapshot/views/2/nodes/0/flags |
| e000755 | 19:55:52.425804 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/0/primary_id |
| e000756 | 19:55:52.425804 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/2/nodes/0/slots |
| e000757 | 19:55:52.425804 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/0/link_state |
| e000758 | 19:55:52.425804 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/2/snapshot/views/2/nodes/1/flags |
| e000759 | 19:55:52.425804 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/1/primary_id |
| e000760 | 19:55:52.425804 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/2/nodes/1/slots |
| e000761 | 19:55:52.425804 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/1/link_state |
| e000762 | 19:55:52.425804 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/2/nodes/2/flags |
| e000763 | 19:55:52.425804 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/2/snapshot/views/2/nodes/2/primary_id |
| e000764 | 19:55:52.425804 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/2/slots |
| e000765 | 19:55:52.425804 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/2/link_state |
| e000766 | 19:55:52.425804 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/2/nodes/3/flags |
| e000767 | 19:55:52.425804 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/3/primary_id |
| e000768 | 19:55:52.425804 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/2/nodes/3/slots |
| e000769 | 19:55:52.425804 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/3/link_state |
| e000770 | 19:55:52.425804 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/2/nodes/4/flags |
| e000771 | 19:55:52.425804 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/2/snapshot/views/2/nodes/4/primary_id |
| e000772 | 19:55:52.425804 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/4/slots |
| e000773 | 19:55:52.425804 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/4/link_state |
| e000774 | 19:55:52.425804 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/2/nodes/5/flags |
| e000775 | 19:55:52.425804 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/2/snapshot/views/2/nodes/5/primary_id |
| e000776 | 19:55:52.425804 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/5/slots |
| e000777 | 19:55:52.425804 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/5/link_state |
| e000778 | 19:55:52.425804 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/2/persistence/aof_enabled |
| e000779 | 19:55:52.425804 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/2/persistence/loading |
| e000780 | 19:55:52.425804 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/2/persistence/aof_last_write_status |
| e000781 | 19:55:52.425804 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e000782 | 19:55:52.425804 | node-3 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/2/process_live |
| e000783 | 19:55:52.425809 | node-5 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/4/reachable |
| e000784 | 19:55:52.425809 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/2/snapshot/views/4/replication/role |
| e000785 | 19:55:52.425809 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/2/snapshot/views/4/replication/connected_slaves |
| e000786 | 19:55:52.425809 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/2/snapshot/views/4/replication/master_link_status |
| e000787 | 19:55:52.425809 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/2/snapshot/views/4/replication/master_sync_in_progress |
| e000788 | 19:55:52.425809 | node-5 | info_replication | replication.master_repl_offset | valid | `14193` | /samples/2/snapshot/views/4/replication/master_repl_offset |
| e000789 | 19:55:52.425809 | node-5 | info_replication | replication.slave_repl_offset | valid | `14193` | /samples/2/snapshot/views/4/replication/slave_repl_offset |
| e000790 | 19:55:52.425809 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /samples/2/snapshot/views/4/redis_id |
| e000791 | 19:55:52.425809 | node-5 | cluster_nodes | slot_coverage | valid | `false` | /samples/2/snapshot/views/4/coverage |
| e000792 | 19:55:52.425809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/4/nodes/0/flags |
| e000793 | 19:55:52.425809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/2/snapshot/views/4/nodes/0/primary_id |
| e000794 | 19:55:52.425809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/0/slots |
| e000795 | 19:55:52.425809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/0/link_state |
| e000796 | 19:55:52.425809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/4/nodes/1/flags |
| e000797 | 19:55:52.425809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/1/primary_id |
| e000798 | 19:55:52.425809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/4/nodes/1/slots |
| e000799 | 19:55:52.425809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/1/link_state |
| e000800 | 19:55:52.425809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/2/snapshot/views/4/nodes/2/flags |
| e000801 | 19:55:52.425809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/2/primary_id |
| e000802 | 19:55:52.425809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/4/nodes/2/slots |
| e000803 | 19:55:52.425809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/2/link_state |
| e000804 | 19:55:52.425809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/2/snapshot/views/4/nodes/3/flags |
| e000805 | 19:55:52.425809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/2/snapshot/views/4/nodes/3/primary_id |
| e000806 | 19:55:52.425809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/3/slots |
| e000807 | 19:55:52.425809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/3/link_state |
| e000808 | 19:55:52.425809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/4/nodes/4/flags |
| e000809 | 19:55:52.425809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/2/snapshot/views/4/nodes/4/primary_id |
| e000810 | 19:55:52.425809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/4/slots |
| e000811 | 19:55:52.425809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/4/link_state |
| e000812 | 19:55:52.425809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/4/nodes/5/flags |
| e000813 | 19:55:52.425809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/5/primary_id |
| e000814 | 19:55:52.425809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/4/nodes/5/slots |
| e000815 | 19:55:52.425809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/5/link_state |
| e000816 | 19:55:52.425809 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/4/persistence/aof_enabled |
| e000817 | 19:55:52.425809 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/4/persistence/loading |
| e000818 | 19:55:52.425809 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/4/persistence/aof_last_write_status |
| e000819 | 19:55:52.425809 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e000820 | 19:55:52.425809 | node-5 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/4/process_live |
| e000821 | 19:55:52.425814 | node-4 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/3/reachable |
| e000822 | 19:55:52.425814 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/2/snapshot/views/3/replication/role |
| e000823 | 19:55:52.425814 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/2/snapshot/views/3/replication/connected_slaves |
| e000824 | 19:55:52.425814 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/2/snapshot/views/3/replication/master_link_status |
| e000825 | 19:55:52.425814 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/2/snapshot/views/3/replication/master_sync_in_progress |
| e000826 | 19:55:52.425814 | node-4 | info_replication | replication.master_repl_offset | valid | `19703` | /samples/2/snapshot/views/3/replication/master_repl_offset |
| e000827 | 19:55:52.425814 | node-4 | info_replication | replication.slave_repl_offset | valid | `19703` | /samples/2/snapshot/views/3/replication/slave_repl_offset |
| e000828 | 19:55:52.425814 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /samples/2/snapshot/views/3/redis_id |
| e000829 | 19:55:52.425814 | node-4 | cluster_nodes | slot_coverage | valid | `false` | /samples/2/snapshot/views/3/coverage |
| e000830 | 19:55:52.425814 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/3/nodes/0/flags |
| e000831 | 19:55:52.425814 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/0/primary_id |
| e000832 | 19:55:52.425814 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/3/nodes/0/slots |
| e000833 | 19:55:52.425814 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/0/link_state |
| e000834 | 19:55:52.425814 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/2/snapshot/views/3/nodes/1/flags |
| e000835 | 19:55:52.425814 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/2/snapshot/views/3/nodes/1/primary_id |
| e000836 | 19:55:52.425814 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/1/slots |
| e000837 | 19:55:52.425814 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/1/link_state |
| e000838 | 19:55:52.425814 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/3/nodes/2/flags |
| e000839 | 19:55:52.425814 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | valid | `"146512ed263b70c95686db0f88d848a4b2c631cf"` | /samples/2/snapshot/views/3/nodes/2/primary_id |
| e000840 | 19:55:52.425814 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/2/slots |
| e000841 | 19:55:52.425814 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/2/link_state |
| e000842 | 19:55:52.425814 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/3/nodes/3/flags |
| e000843 | 19:55:52.425814 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/2/snapshot/views/3/nodes/3/primary_id |
| e000844 | 19:55:52.425814 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/3/slots |
| e000845 | 19:55:52.425814 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/3/link_state |
| e000846 | 19:55:52.425814 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/3/nodes/4/flags |
| e000847 | 19:55:52.425814 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/4/primary_id |
| e000848 | 19:55:52.425814 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/3/nodes/4/slots |
| e000849 | 19:55:52.425814 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/4/link_state |
| e000850 | 19:55:52.425814 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/2/snapshot/views/3/nodes/5/flags |
| e000851 | 19:55:52.425814 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/5/primary_id |
| e000852 | 19:55:52.425814 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/3/nodes/5/slots |
| e000853 | 19:55:52.425814 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/5/link_state |
| e000854 | 19:55:52.425814 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/3/persistence/aof_enabled |
| e000855 | 19:55:52.425814 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/3/persistence/loading |
| e000856 | 19:55:52.425814 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/3/persistence/aof_last_write_status |
| e000857 | 19:55:52.425814 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e000858 | 19:55:52.425814 | node-4 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/3/process_live |
| e000859 | 19:55:53.055222 |  | client_probe | client_success | valid | `false` | /samples/2/probe/successful |
| e000860 | 19:55:53.224589 |  | client_window | client.attempts | valid | `66` | /metrics/windows/4/attempts |
| e000861 | 19:55:53.224589 |  | client_window | client.errors | valid | `25` | /metrics/windows/4/errors |
| e000862 | 19:55:53.224589 |  | client_window | client.successes | valid | `41` | /metrics/windows/4/successes |
| e000863 | 19:55:53.224589 |  | client_window | client.p95_ms | valid | `500.522084` | /metrics/windows/4/p95_ms |
| e000864 | 19:55:53.365888 | node-6 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/5/reachable |
| e000865 | 19:55:53.365888 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/3/snapshot/views/5/replication/role |
| e000866 | 19:55:53.365888 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/3/snapshot/views/5/replication/connected_slaves |
| e000867 | 19:55:53.365888 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/3/snapshot/views/5/replication/master_link_status |
| e000868 | 19:55:53.365888 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/3/snapshot/views/5/replication/master_sync_in_progress |
| e000869 | 19:55:53.365888 | node-6 | info_replication | replication.master_repl_offset | valid | `22013` | /samples/3/snapshot/views/5/replication/master_repl_offset |
| e000870 | 19:55:53.365888 | node-6 | info_replication | replication.slave_repl_offset | valid | `22013` | /samples/3/snapshot/views/5/replication/slave_repl_offset |
| e000871 | 19:55:53.365888 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /samples/3/snapshot/views/5/redis_id |
| e000872 | 19:55:53.365888 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/5/coverage |
| e000873 | 19:55:53.365888 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/5/nodes/0/flags |
| e000874 | 19:55:53.365888 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/3/snapshot/views/5/nodes/0/primary_id |
| e000875 | 19:55:53.365888 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/0/slots |
| e000876 | 19:55:53.365888 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/0/link_state |
| e000877 | 19:55:53.365888 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/5/nodes/1/flags |
| e000878 | 19:55:53.365888 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/1/primary_id |
| e000879 | 19:55:53.365888 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/5/nodes/1/slots |
| e000880 | 19:55:53.365888 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/1/link_state |
| e000881 | 19:55:53.365888 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/3/snapshot/views/5/nodes/2/flags |
| e000882 | 19:55:53.365888 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/2/primary_id |
| e000883 | 19:55:53.365888 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/2/slots |
| e000884 | 19:55:53.365888 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/2/link_state |
| e000885 | 19:55:53.365888 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/5/nodes/3/flags |
| e000886 | 19:55:53.365888 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/3/primary_id |
| e000887 | 19:55:53.365888 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/5/nodes/3/slots |
| e000888 | 19:55:53.365888 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/3/link_state |
| e000889 | 19:55:53.365888 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/3/snapshot/views/5/nodes/4/flags |
| e000890 | 19:55:53.365888 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/3/snapshot/views/5/nodes/4/primary_id |
| e000891 | 19:55:53.365888 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/4/slots |
| e000892 | 19:55:53.365888 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/4/link_state |
| e000893 | 19:55:53.365888 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/5/nodes/5/flags |
| e000894 | 19:55:53.365888 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/5/primary_id |
| e000895 | 19:55:53.365888 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/5/nodes/5/slots |
| e000896 | 19:55:53.365888 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/5/link_state |
| e000897 | 19:55:53.365888 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/5/persistence/aof_enabled |
| e000898 | 19:55:53.365888 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/5/persistence/loading |
| e000899 | 19:55:53.365888 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/5/persistence/aof_last_write_status |
| e000900 | 19:55:53.365888 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e000901 | 19:55:53.365888 | node-6 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/5/process_live |
| e000902 | 19:55:53.36589 | node-3 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/2/reachable |
| e000903 | 19:55:53.36589 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/3/snapshot/views/2/replication/role |
| e000904 | 19:55:53.36589 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/3/snapshot/views/2/replication/connected_slaves |
| e000905 | 19:55:53.36589 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/3/snapshot/views/2/replication/master_link_status |
| e000906 | 19:55:53.36589 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/3/snapshot/views/2/replication/master_sync_in_progress |
| e000907 | 19:55:53.36589 | node-3 | info_replication | replication.master_repl_offset | valid | `22013` | /samples/3/snapshot/views/2/replication/master_repl_offset |
| e000908 | 19:55:53.36589 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/3/snapshot/views/2/replication/slave_repl_offset |
| e000909 | 19:55:53.36589 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/3/snapshot/views/2/redis_id |
| e000910 | 19:55:53.36589 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/2/coverage |
| e000911 | 19:55:53.36589 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/3/snapshot/views/2/nodes/0/flags |
| e000912 | 19:55:53.36589 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/0/primary_id |
| e000913 | 19:55:53.36589 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/2/nodes/0/slots |
| e000914 | 19:55:53.36589 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/0/link_state |
| e000915 | 19:55:53.36589 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/3/snapshot/views/2/nodes/1/flags |
| e000916 | 19:55:53.36589 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/1/primary_id |
| e000917 | 19:55:53.36589 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/1/slots |
| e000918 | 19:55:53.36589 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/1/link_state |
| e000919 | 19:55:53.36589 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/2/nodes/2/flags |
| e000920 | 19:55:53.36589 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/3/snapshot/views/2/nodes/2/primary_id |
| e000921 | 19:55:53.36589 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/2/slots |
| e000922 | 19:55:53.36589 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/2/link_state |
| e000923 | 19:55:53.36589 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/2/nodes/3/flags |
| e000924 | 19:55:53.36589 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/3/primary_id |
| e000925 | 19:55:53.36589 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/2/nodes/3/slots |
| e000926 | 19:55:53.36589 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/3/link_state |
| e000927 | 19:55:53.36589 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/2/nodes/4/flags |
| e000928 | 19:55:53.36589 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/4/primary_id |
| e000929 | 19:55:53.36589 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/2/nodes/4/slots |
| e000930 | 19:55:53.36589 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/4/link_state |
| e000931 | 19:55:53.36589 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/2/nodes/5/flags |
| e000932 | 19:55:53.36589 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/3/snapshot/views/2/nodes/5/primary_id |
| e000933 | 19:55:53.36589 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/5/slots |
| e000934 | 19:55:53.36589 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/5/link_state |
| e000935 | 19:55:53.36589 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/2/persistence/aof_enabled |
| e000936 | 19:55:53.36589 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/2/persistence/loading |
| e000937 | 19:55:53.36589 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/2/persistence/aof_last_write_status |
| e000938 | 19:55:53.36589 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e000939 | 19:55:53.36589 | node-3 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/2/process_live |
| e000940 | 19:55:53.365891 | node-2 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/1/reachable |
| e000941 | 19:55:53.365891 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/3/snapshot/views/1/replication/role |
| e000942 | 19:55:53.365891 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/3/snapshot/views/1/replication/connected_slaves |
| e000943 | 19:55:53.365891 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/3/snapshot/views/1/replication/master_link_status |
| e000944 | 19:55:53.365891 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/3/snapshot/views/1/replication/master_sync_in_progress |
| e000945 | 19:55:53.365891 | node-2 | info_replication | replication.master_repl_offset | valid | `22235` | /samples/3/snapshot/views/1/replication/master_repl_offset |
| e000946 | 19:55:53.365891 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/3/snapshot/views/1/replication/slave_repl_offset |
| e000947 | 19:55:53.365891 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/3/snapshot/views/1/redis_id |
| e000948 | 19:55:53.365891 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/1/coverage |
| e000949 | 19:55:53.365891 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/3/snapshot/views/1/nodes/0/flags |
| e000950 | 19:55:53.365891 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/0/primary_id |
| e000951 | 19:55:53.365891 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/0/slots |
| e000952 | 19:55:53.365891 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/0/link_state |
| e000953 | 19:55:53.365891 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/1/nodes/1/flags |
| e000954 | 19:55:53.365891 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/3/snapshot/views/1/nodes/1/primary_id |
| e000955 | 19:55:53.365891 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/1/slots |
| e000956 | 19:55:53.365891 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/1/link_state |
| e000957 | 19:55:53.365891 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/1/nodes/2/flags |
| e000958 | 19:55:53.365891 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/2/primary_id |
| e000959 | 19:55:53.365891 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/1/nodes/2/slots |
| e000960 | 19:55:53.365891 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/2/link_state |
| e000961 | 19:55:53.365891 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/3/snapshot/views/1/nodes/3/flags |
| e000962 | 19:55:53.365891 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/3/primary_id |
| e000963 | 19:55:53.365891 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/1/nodes/3/slots |
| e000964 | 19:55:53.365891 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/3/link_state |
| e000965 | 19:55:53.365891 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/1/nodes/4/flags |
| e000966 | 19:55:53.365891 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/4/primary_id |
| e000967 | 19:55:53.365891 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/1/nodes/4/slots |
| e000968 | 19:55:53.365891 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/4/link_state |
| e000969 | 19:55:53.365891 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/1/nodes/5/flags |
| e000970 | 19:55:53.365891 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/3/snapshot/views/1/nodes/5/primary_id |
| e000971 | 19:55:53.365891 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/5/slots |
| e000972 | 19:55:53.365891 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/5/link_state |
| e000973 | 19:55:53.365891 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/1/persistence/aof_enabled |
| e000974 | 19:55:53.365891 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/1/persistence/loading |
| e000975 | 19:55:53.365891 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/1/persistence/aof_last_write_status |
| e000976 | 19:55:53.365891 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e000977 | 19:55:53.365891 | node-2 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/1/process_live |
| e000978 | 19:55:53.365904 | node-4 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/3/reachable |
| e000979 | 19:55:53.365904 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/3/snapshot/views/3/replication/role |
| e000980 | 19:55:53.365904 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/3/snapshot/views/3/replication/connected_slaves |
| e000981 | 19:55:53.365904 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/3/snapshot/views/3/replication/master_link_status |
| e000982 | 19:55:53.365904 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/3/snapshot/views/3/replication/master_sync_in_progress |
| e000983 | 19:55:53.365904 | node-4 | info_replication | replication.master_repl_offset | valid | `22013` | /samples/3/snapshot/views/3/replication/master_repl_offset |
| e000984 | 19:55:53.365904 | node-4 | info_replication | replication.slave_repl_offset | valid | `22013` | /samples/3/snapshot/views/3/replication/slave_repl_offset |
| e000985 | 19:55:53.365904 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /samples/3/snapshot/views/3/redis_id |
| e000986 | 19:55:53.365904 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/3/coverage |
| e000987 | 19:55:53.365904 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/3/nodes/0/flags |
| e000988 | 19:55:53.365904 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/0/primary_id |
| e000989 | 19:55:53.365904 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/3/nodes/0/slots |
| e000990 | 19:55:53.365904 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/0/link_state |
| e000991 | 19:55:53.365904 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/3/snapshot/views/3/nodes/1/flags |
| e000992 | 19:55:53.365904 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/3/snapshot/views/3/nodes/1/primary_id |
| e000993 | 19:55:53.365904 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/1/slots |
| e000994 | 19:55:53.365904 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/1/link_state |
| e000995 | 19:55:53.365904 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/3/nodes/2/flags |
| e000996 | 19:55:53.365904 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/2/primary_id |
| e000997 | 19:55:53.365904 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/3/nodes/2/slots |
| e000998 | 19:55:53.365904 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/2/link_state |
| e000999 | 19:55:53.365904 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/3/nodes/3/flags |
| e001000 | 19:55:53.365904 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/3/snapshot/views/3/nodes/3/primary_id |
| e001001 | 19:55:53.365904 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/3/slots |
| e001002 | 19:55:53.365904 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/3/link_state |
| e001003 | 19:55:53.365904 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/3/nodes/4/flags |
| e001004 | 19:55:53.365904 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/4/primary_id |
| e001005 | 19:55:53.365904 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/3/nodes/4/slots |
| e001006 | 19:55:53.365904 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/4/link_state |
| e001007 | 19:55:53.365904 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/3/snapshot/views/3/nodes/5/flags |
| e001008 | 19:55:53.365904 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/5/primary_id |
| e001009 | 19:55:53.365904 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/5/slots |
| e001010 | 19:55:53.365904 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/5/link_state |
| e001011 | 19:55:53.365904 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/3/persistence/aof_enabled |
| e001012 | 19:55:53.365904 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/3/persistence/loading |
| e001013 | 19:55:53.365904 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/3/persistence/aof_last_write_status |
| e001014 | 19:55:53.365904 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001015 | 19:55:53.365904 | node-4 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/3/process_live |
| e001016 | 19:55:53.365922 | node-5 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/4/reachable |
| e001017 | 19:55:53.365922 | node-5 | info_replication | replication.role | valid | `"master"` | /samples/3/snapshot/views/4/replication/role |
| e001018 | 19:55:53.365922 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/3/snapshot/views/4/replication/connected_slaves |
| e001019 | 19:55:53.365922 | node-5 | info_replication | replication.master_link_status | not_applicable | `` | /samples/3/snapshot/views/4/replication/master_link_status |
| e001020 | 19:55:53.365922 | node-5 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/3/snapshot/views/4/replication/master_sync_in_progress |
| e001021 | 19:55:53.365922 | node-5 | info_replication | replication.master_repl_offset | valid | `14882` | /samples/3/snapshot/views/4/replication/master_repl_offset |
| e001022 | 19:55:53.365922 | node-5 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/3/snapshot/views/4/replication/slave_repl_offset |
| e001023 | 19:55:53.365922 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /samples/3/snapshot/views/4/redis_id |
| e001024 | 19:55:53.365922 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/4/coverage |
| e001025 | 19:55:53.365922 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/4/nodes/0/flags |
| e001026 | 19:55:53.365922 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/3/snapshot/views/4/nodes/0/primary_id |
| e001027 | 19:55:53.365922 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/0/slots |
| e001028 | 19:55:53.365922 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/0/link_state |
| e001029 | 19:55:53.365922 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/4/nodes/1/flags |
| e001030 | 19:55:53.365922 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/1/primary_id |
| e001031 | 19:55:53.365922 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/4/nodes/1/slots |
| e001032 | 19:55:53.365922 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/1/link_state |
| e001033 | 19:55:53.365922 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/3/snapshot/views/4/nodes/2/flags |
| e001034 | 19:55:53.365922 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/2/primary_id |
| e001035 | 19:55:53.365922 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/2/slots |
| e001036 | 19:55:53.365922 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/2/link_state |
| e001037 | 19:55:53.365922 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/3/snapshot/views/4/nodes/3/flags |
| e001038 | 19:55:53.365922 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/3/primary_id |
| e001039 | 19:55:53.365922 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/4/nodes/3/slots |
| e001040 | 19:55:53.365922 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/3/link_state |
| e001041 | 19:55:53.365922 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/4/nodes/4/flags |
| e001042 | 19:55:53.365922 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/3/snapshot/views/4/nodes/4/primary_id |
| e001043 | 19:55:53.365922 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/4/slots |
| e001044 | 19:55:53.365922 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/4/link_state |
| e001045 | 19:55:53.365922 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/4/nodes/5/flags |
| e001046 | 19:55:53.365922 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/5/primary_id |
| e001047 | 19:55:53.365922 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/4/nodes/5/slots |
| e001048 | 19:55:53.365922 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/5/link_state |
| e001049 | 19:55:53.365922 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/4/persistence/aof_enabled |
| e001050 | 19:55:53.365922 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/4/persistence/loading |
| e001051 | 19:55:53.365922 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/4/persistence/aof_last_write_status |
| e001052 | 19:55:53.365922 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001053 | 19:55:53.365922 | node-5 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/4/process_live |
| e001054 | 19:55:53.365931 | node-1 | node_client | reachable | valid | `false` | /samples/3/snapshot/views/0/reachable |
| e001055 | 19:55:53.365931 | node-1 | info_replication | replication.role | collection_failed | `` | /samples/3/snapshot/views/0/replication/role |
| e001056 | 19:55:53.365931 | node-1 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/3/snapshot/views/0/replication/connected_slaves |
| e001057 | 19:55:53.365931 | node-1 | info_replication | replication.master_link_status | collection_failed | `` | /samples/3/snapshot/views/0/replication/master_link_status |
| e001058 | 19:55:53.365931 | node-1 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/3/snapshot/views/0/replication/master_sync_in_progress |
| e001059 | 19:55:53.365931 | node-1 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/3/snapshot/views/0/replication/master_repl_offset |
| e001060 | 19:55:53.365931 | node-1 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/3/snapshot/views/0/replication/slave_repl_offset |
| e001061 | 19:55:53.365931 | node-1 | cluster_nodes | redis_id | collection_failed | `` | /samples/3/snapshot/views/0/redis_id |
| e001062 | 19:55:53.365931 | node-1 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/3/snapshot/views/0/coverage |
| e001063 | 19:55:53.365931 | node-1 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/3/snapshot/views/0/persistence/aof_enabled |
| e001064 | 19:55:53.365931 | node-1 | info_persistence | persistence.loading | collection_failed | `` | /samples/3/snapshot/views/0/persistence/loading |
| e001065 | 19:55:53.365931 | node-1 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/3/snapshot/views/0/persistence/aof_last_write_status |
| e001066 | 19:55:53.365931 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/3/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001067 | 19:55:53.365931 | node-1 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/0/process_live |
| e001068 | 19:55:53.498116 |  | client_probe | client_success | valid | `true` | /samples/3/probe/successful |
| e001069 | 19:55:53.808961 | node-5 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/4/reachable |
| e001070 | 19:55:53.808961 | node-5 | info_replication | replication.role | valid | `"master"` | /samples/4/snapshot/views/4/replication/role |
| e001071 | 19:55:53.808961 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/4/snapshot/views/4/replication/connected_slaves |
| e001072 | 19:55:53.808961 | node-5 | info_replication | replication.master_link_status | not_applicable | `` | /samples/4/snapshot/views/4/replication/master_link_status |
| e001073 | 19:55:53.808961 | node-5 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/4/snapshot/views/4/replication/master_sync_in_progress |
| e001074 | 19:55:53.808961 | node-5 | info_replication | replication.master_repl_offset | valid | `18302` | /samples/4/snapshot/views/4/replication/master_repl_offset |
| e001075 | 19:55:53.808961 | node-5 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/4/snapshot/views/4/replication/slave_repl_offset |
| e001076 | 19:55:53.808961 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /samples/4/snapshot/views/4/redis_id |
| e001077 | 19:55:53.808961 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/4/coverage |
| e001078 | 19:55:53.808961 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/4/nodes/0/flags |
| e001079 | 19:55:53.808961 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/4/snapshot/views/4/nodes/0/primary_id |
| e001080 | 19:55:53.808961 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/0/slots |
| e001081 | 19:55:53.808961 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/4/nodes/0/link_state |
| e001082 | 19:55:53.808961 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/4/nodes/1/flags |
| e001083 | 19:55:53.808961 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/1/primary_id |
| e001084 | 19:55:53.808961 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/4/nodes/1/slots |
| e001085 | 19:55:53.808961 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/4/nodes/1/link_state |
| e001086 | 19:55:53.808961 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/4/snapshot/views/4/nodes/2/flags |
| e001087 | 19:55:53.808961 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/2/primary_id |
| e001088 | 19:55:53.808961 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/2/slots |
| e001089 | 19:55:53.808961 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"disconnected"` | /samples/4/snapshot/views/4/nodes/2/link_state |
| e001090 | 19:55:53.808961 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/4/snapshot/views/4/nodes/3/flags |
| e001091 | 19:55:53.808961 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/3/primary_id |
| e001092 | 19:55:53.808961 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/4/nodes/3/slots |
| e001093 | 19:55:53.808961 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/4/nodes/3/link_state |
| e001094 | 19:55:53.808961 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/4/nodes/4/flags |
| e001095 | 19:55:53.808961 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/4/snapshot/views/4/nodes/4/primary_id |
| e001096 | 19:55:53.808961 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/4/slots |
| e001097 | 19:55:53.808961 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/4/nodes/4/link_state |
| e001098 | 19:55:53.808961 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/4/nodes/5/flags |
| e001099 | 19:55:53.808961 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/5/primary_id |
| e001100 | 19:55:53.808961 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/4/nodes/5/slots |
| e001101 | 19:55:53.808961 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/4/nodes/5/link_state |
| e001102 | 19:55:53.808961 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/4/persistence/aof_enabled |
| e001103 | 19:55:53.808961 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/4/persistence/loading |
| e001104 | 19:55:53.808961 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/4/persistence/aof_last_write_status |
| e001105 | 19:55:53.808961 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001106 | 19:55:53.808961 | node-5 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/4/process_live |
| e001107 | 19:55:53.808968 | node-4 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/3/reachable |
| e001108 | 19:55:53.808968 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/4/snapshot/views/3/replication/role |
| e001109 | 19:55:53.808968 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/4/snapshot/views/3/replication/connected_slaves |
| e001110 | 19:55:53.808968 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/4/snapshot/views/3/replication/master_link_status |
| e001111 | 19:55:53.808968 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/4/snapshot/views/3/replication/master_sync_in_progress |
| e001112 | 19:55:53.808968 | node-4 | info_replication | replication.master_repl_offset | valid | `25433` | /samples/4/snapshot/views/3/replication/master_repl_offset |
| e001113 | 19:55:53.808968 | node-4 | info_replication | replication.slave_repl_offset | valid | `25433` | /samples/4/snapshot/views/3/replication/slave_repl_offset |
| e001114 | 19:55:53.808968 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /samples/4/snapshot/views/3/redis_id |
| e001115 | 19:55:53.808968 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/3/coverage |
| e001116 | 19:55:53.808968 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/3/nodes/0/flags |
| e001117 | 19:55:53.808968 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/0/primary_id |
| e001118 | 19:55:53.808968 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/3/nodes/0/slots |
| e001119 | 19:55:53.808968 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/0/link_state |
| e001120 | 19:55:53.808968 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/4/snapshot/views/3/nodes/1/flags |
| e001121 | 19:55:53.808968 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/4/snapshot/views/3/nodes/1/primary_id |
| e001122 | 19:55:53.808968 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/1/slots |
| e001123 | 19:55:53.808968 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/1/link_state |
| e001124 | 19:55:53.808968 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/3/nodes/2/flags |
| e001125 | 19:55:53.808968 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/2/primary_id |
| e001126 | 19:55:53.808968 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/3/nodes/2/slots |
| e001127 | 19:55:53.808968 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/2/link_state |
| e001128 | 19:55:53.808968 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/3/nodes/3/flags |
| e001129 | 19:55:53.808968 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/4/snapshot/views/3/nodes/3/primary_id |
| e001130 | 19:55:53.808968 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/3/slots |
| e001131 | 19:55:53.808968 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/3/link_state |
| e001132 | 19:55:53.808968 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/3/nodes/4/flags |
| e001133 | 19:55:53.808968 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/4/primary_id |
| e001134 | 19:55:53.808968 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/3/nodes/4/slots |
| e001135 | 19:55:53.808968 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/4/link_state |
| e001136 | 19:55:53.808968 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/4/snapshot/views/3/nodes/5/flags |
| e001137 | 19:55:53.808968 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/5/primary_id |
| e001138 | 19:55:53.808968 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/5/slots |
| e001139 | 19:55:53.808968 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"disconnected"` | /samples/4/snapshot/views/3/nodes/5/link_state |
| e001140 | 19:55:53.808968 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/3/persistence/aof_enabled |
| e001141 | 19:55:53.808968 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/3/persistence/loading |
| e001142 | 19:55:53.808968 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/3/persistence/aof_last_write_status |
| e001143 | 19:55:53.808968 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001144 | 19:55:53.808968 | node-4 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/3/process_live |
| e001145 | 19:55:53.808991 | node-2 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/1/reachable |
| e001146 | 19:55:53.808991 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/4/snapshot/views/1/replication/role |
| e001147 | 19:55:53.808991 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/4/snapshot/views/1/replication/connected_slaves |
| e001148 | 19:55:53.808991 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/4/snapshot/views/1/replication/master_link_status |
| e001149 | 19:55:53.808991 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/4/snapshot/views/1/replication/master_sync_in_progress |
| e001150 | 19:55:53.808991 | node-2 | info_replication | replication.master_repl_offset | valid | `25655` | /samples/4/snapshot/views/1/replication/master_repl_offset |
| e001151 | 19:55:53.808991 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/4/snapshot/views/1/replication/slave_repl_offset |
| e001152 | 19:55:53.808991 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/4/snapshot/views/1/redis_id |
| e001153 | 19:55:53.808991 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/1/coverage |
| e001154 | 19:55:53.808991 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/4/snapshot/views/1/nodes/0/flags |
| e001155 | 19:55:53.808991 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/0/primary_id |
| e001156 | 19:55:53.808991 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/0/slots |
| e001157 | 19:55:53.808991 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"disconnected"` | /samples/4/snapshot/views/1/nodes/0/link_state |
| e001158 | 19:55:53.808991 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/1/nodes/1/flags |
| e001159 | 19:55:53.808991 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/4/snapshot/views/1/nodes/1/primary_id |
| e001160 | 19:55:53.808991 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/1/slots |
| e001161 | 19:55:53.808991 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/1/link_state |
| e001162 | 19:55:53.808991 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/1/nodes/2/flags |
| e001163 | 19:55:53.808991 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/2/primary_id |
| e001164 | 19:55:53.808991 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/1/nodes/2/slots |
| e001165 | 19:55:53.808991 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/2/link_state |
| e001166 | 19:55:53.808991 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/4/snapshot/views/1/nodes/3/flags |
| e001167 | 19:55:53.808991 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/3/primary_id |
| e001168 | 19:55:53.808991 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/1/nodes/3/slots |
| e001169 | 19:55:53.808991 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/3/link_state |
| e001170 | 19:55:53.808991 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/1/nodes/4/flags |
| e001171 | 19:55:53.808991 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/4/primary_id |
| e001172 | 19:55:53.808991 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/1/nodes/4/slots |
| e001173 | 19:55:53.808991 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/4/link_state |
| e001174 | 19:55:53.808991 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/1/nodes/5/flags |
| e001175 | 19:55:53.808991 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/4/snapshot/views/1/nodes/5/primary_id |
| e001176 | 19:55:53.808991 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/5/slots |
| e001177 | 19:55:53.808991 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/5/link_state |
| e001178 | 19:55:53.808991 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/1/persistence/aof_enabled |
| e001179 | 19:55:53.808991 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/1/persistence/loading |
| e001180 | 19:55:53.808991 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/1/persistence/aof_last_write_status |
| e001181 | 19:55:53.808991 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001182 | 19:55:53.808991 | node-2 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/1/process_live |
| e001183 | 19:55:53.809 | node-3 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/2/reachable |
| e001184 | 19:55:53.809 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/4/snapshot/views/2/replication/role |
| e001185 | 19:55:53.809 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/4/snapshot/views/2/replication/connected_slaves |
| e001186 | 19:55:53.809 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/4/snapshot/views/2/replication/master_link_status |
| e001187 | 19:55:53.809 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/4/snapshot/views/2/replication/master_sync_in_progress |
| e001188 | 19:55:53.809 | node-3 | info_replication | replication.master_repl_offset | valid | `25433` | /samples/4/snapshot/views/2/replication/master_repl_offset |
| e001189 | 19:55:53.809 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/4/snapshot/views/2/replication/slave_repl_offset |
| e001190 | 19:55:53.809 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/4/snapshot/views/2/redis_id |
| e001191 | 19:55:53.809 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/2/coverage |
| e001192 | 19:55:53.809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/4/snapshot/views/2/nodes/0/flags |
| e001193 | 19:55:53.809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/0/primary_id |
| e001194 | 19:55:53.809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/2/nodes/0/slots |
| e001195 | 19:55:53.809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/0/link_state |
| e001196 | 19:55:53.809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/4/snapshot/views/2/nodes/1/flags |
| e001197 | 19:55:53.809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/1/primary_id |
| e001198 | 19:55:53.809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/1/slots |
| e001199 | 19:55:53.809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"disconnected"` | /samples/4/snapshot/views/2/nodes/1/link_state |
| e001200 | 19:55:53.809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/2/nodes/2/flags |
| e001201 | 19:55:53.809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/4/snapshot/views/2/nodes/2/primary_id |
| e001202 | 19:55:53.809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/2/slots |
| e001203 | 19:55:53.809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/2/link_state |
| e001204 | 19:55:53.809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/2/nodes/3/flags |
| e001205 | 19:55:53.809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/3/primary_id |
| e001206 | 19:55:53.809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/2/nodes/3/slots |
| e001207 | 19:55:53.809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/3/link_state |
| e001208 | 19:55:53.809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/2/nodes/4/flags |
| e001209 | 19:55:53.809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/4/primary_id |
| e001210 | 19:55:53.809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/2/nodes/4/slots |
| e001211 | 19:55:53.809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/4/link_state |
| e001212 | 19:55:53.809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/2/nodes/5/flags |
| e001213 | 19:55:53.809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/4/snapshot/views/2/nodes/5/primary_id |
| e001214 | 19:55:53.809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/5/slots |
| e001215 | 19:55:53.809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/5/link_state |
| e001216 | 19:55:53.809 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/2/persistence/aof_enabled |
| e001217 | 19:55:53.809 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/2/persistence/loading |
| e001218 | 19:55:53.809 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/2/persistence/aof_last_write_status |
| e001219 | 19:55:53.809 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001220 | 19:55:53.809 | node-3 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/2/process_live |
| e001221 | 19:55:53.809 | node-6 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/5/reachable |
| e001222 | 19:55:53.809 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/4/snapshot/views/5/replication/role |
| e001223 | 19:55:53.809 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/4/snapshot/views/5/replication/connected_slaves |
| e001224 | 19:55:53.809 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/4/snapshot/views/5/replication/master_link_status |
| e001225 | 19:55:53.809 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/4/snapshot/views/5/replication/master_sync_in_progress |
| e001226 | 19:55:53.809 | node-6 | info_replication | replication.master_repl_offset | valid | `25655` | /samples/4/snapshot/views/5/replication/master_repl_offset |
| e001227 | 19:55:53.809 | node-6 | info_replication | replication.slave_repl_offset | valid | `25655` | /samples/4/snapshot/views/5/replication/slave_repl_offset |
| e001228 | 19:55:53.809 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /samples/4/snapshot/views/5/redis_id |
| e001229 | 19:55:53.809 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/5/coverage |
| e001230 | 19:55:53.809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/5/nodes/0/flags |
| e001231 | 19:55:53.809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/4/snapshot/views/5/nodes/0/primary_id |
| e001232 | 19:55:53.809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/0/slots |
| e001233 | 19:55:53.809 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/0/link_state |
| e001234 | 19:55:53.809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/5/nodes/1/flags |
| e001235 | 19:55:53.809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/1/primary_id |
| e001236 | 19:55:53.809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/5/nodes/1/slots |
| e001237 | 19:55:53.809 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/1/link_state |
| e001238 | 19:55:53.809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/4/snapshot/views/5/nodes/2/flags |
| e001239 | 19:55:53.809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/2/primary_id |
| e001240 | 19:55:53.809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/2/slots |
| e001241 | 19:55:53.809 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"disconnected"` | /samples/4/snapshot/views/5/nodes/2/link_state |
| e001242 | 19:55:53.809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/5/nodes/3/flags |
| e001243 | 19:55:53.809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/3/primary_id |
| e001244 | 19:55:53.809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/5/nodes/3/slots |
| e001245 | 19:55:53.809 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/3/link_state |
| e001246 | 19:55:53.809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/4/snapshot/views/5/nodes/4/flags |
| e001247 | 19:55:53.809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/4/snapshot/views/5/nodes/4/primary_id |
| e001248 | 19:55:53.809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/4/slots |
| e001249 | 19:55:53.809 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/4/link_state |
| e001250 | 19:55:53.809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/5/nodes/5/flags |
| e001251 | 19:55:53.809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/5/primary_id |
| e001252 | 19:55:53.809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/5/nodes/5/slots |
| e001253 | 19:55:53.809 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/5/link_state |
| e001254 | 19:55:53.809 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/5/persistence/aof_enabled |
| e001255 | 19:55:53.809 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/5/persistence/loading |
| e001256 | 19:55:53.809 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/5/persistence/aof_last_write_status |
| e001257 | 19:55:53.809 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e001258 | 19:55:53.809 | node-6 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/5/process_live |
| e001259 | 19:55:53.809041 | node-1 | node_client | reachable | valid | `false` | /samples/4/snapshot/views/0/reachable |
| e001260 | 19:55:53.809041 | node-1 | info_replication | replication.role | collection_failed | `` | /samples/4/snapshot/views/0/replication/role |
| e001261 | 19:55:53.809041 | node-1 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/4/snapshot/views/0/replication/connected_slaves |
| e001262 | 19:55:53.809041 | node-1 | info_replication | replication.master_link_status | collection_failed | `` | /samples/4/snapshot/views/0/replication/master_link_status |
| e001263 | 19:55:53.809041 | node-1 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/4/snapshot/views/0/replication/master_sync_in_progress |
| e001264 | 19:55:53.809041 | node-1 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/4/snapshot/views/0/replication/master_repl_offset |
| e001265 | 19:55:53.809041 | node-1 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/4/snapshot/views/0/replication/slave_repl_offset |
| e001266 | 19:55:53.809041 | node-1 | cluster_nodes | redis_id | collection_failed | `` | /samples/4/snapshot/views/0/redis_id |
| e001267 | 19:55:53.809041 | node-1 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/4/snapshot/views/0/coverage |
| e001268 | 19:55:53.809041 | node-1 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/4/snapshot/views/0/persistence/aof_enabled |
| e001269 | 19:55:53.809041 | node-1 | info_persistence | persistence.loading | collection_failed | `` | /samples/4/snapshot/views/0/persistence/loading |
| e001270 | 19:55:53.809041 | node-1 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/4/snapshot/views/0/persistence/aof_last_write_status |
| e001271 | 19:55:53.809041 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/4/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001272 | 19:55:53.809041 | node-1 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/0/process_live |
| e001273 | 19:55:53.942102 |  | client_probe | client_success | valid | `true` | /samples/4/probe/successful |
| e001274 | 19:55:54.224589 |  | client_window | client.attempts | valid | `167` | /metrics/windows/5/attempts |
| e001275 | 19:55:54.224589 |  | client_window | client.errors | valid | `4` | /metrics/windows/5/errors |
| e001276 | 19:55:54.224589 |  | client_window | client.successes | valid | `163` | /metrics/windows/5/successes |
| e001277 | 19:55:54.224589 |  | client_window | client.p95_ms | valid | `1.314313` | /metrics/windows/5/p95_ms |
| e001278 | 19:55:54.252285 | node-2 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/1/reachable |
| e001279 | 19:55:54.252285 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/5/snapshot/views/1/replication/role |
| e001280 | 19:55:54.252285 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/5/snapshot/views/1/replication/connected_slaves |
| e001281 | 19:55:54.252285 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/5/snapshot/views/1/replication/master_link_status |
| e001282 | 19:55:54.252285 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/5/snapshot/views/1/replication/master_sync_in_progress |
| e001283 | 19:55:54.252285 | node-2 | info_replication | replication.master_repl_offset | valid | `28631` | /samples/5/snapshot/views/1/replication/master_repl_offset |
| e001284 | 19:55:54.252285 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/5/snapshot/views/1/replication/slave_repl_offset |
| e001285 | 19:55:54.252285 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/5/snapshot/views/1/redis_id |
| e001286 | 19:55:54.252285 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/1/coverage |
| e001287 | 19:55:54.252285 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/5/snapshot/views/1/nodes/0/flags |
| e001288 | 19:55:54.252285 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/0/primary_id |
| e001289 | 19:55:54.252285 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/0/slots |
| e001290 | 19:55:54.252285 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/0/link_state |
| e001291 | 19:55:54.252285 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/1/nodes/1/flags |
| e001292 | 19:55:54.252285 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/5/snapshot/views/1/nodes/1/primary_id |
| e001293 | 19:55:54.252285 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/1/slots |
| e001294 | 19:55:54.252285 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/1/link_state |
| e001295 | 19:55:54.252285 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/1/nodes/2/flags |
| e001296 | 19:55:54.252285 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/2/primary_id |
| e001297 | 19:55:54.252285 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/1/nodes/2/slots |
| e001298 | 19:55:54.252285 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/2/link_state |
| e001299 | 19:55:54.252285 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/5/snapshot/views/1/nodes/3/flags |
| e001300 | 19:55:54.252285 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/3/primary_id |
| e001301 | 19:55:54.252285 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/1/nodes/3/slots |
| e001302 | 19:55:54.252285 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/3/link_state |
| e001303 | 19:55:54.252285 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/1/nodes/4/flags |
| e001304 | 19:55:54.252285 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/4/primary_id |
| e001305 | 19:55:54.252285 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/1/nodes/4/slots |
| e001306 | 19:55:54.252285 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/4/link_state |
| e001307 | 19:55:54.252285 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/1/nodes/5/flags |
| e001308 | 19:55:54.252285 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/5/snapshot/views/1/nodes/5/primary_id |
| e001309 | 19:55:54.252285 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/5/slots |
| e001310 | 19:55:54.252285 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/5/link_state |
| e001311 | 19:55:54.252285 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/1/persistence/aof_enabled |
| e001312 | 19:55:54.252285 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/1/persistence/loading |
| e001313 | 19:55:54.252285 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/1/persistence/aof_last_write_status |
| e001314 | 19:55:54.252285 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001315 | 19:55:54.252285 | node-2 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/1/process_live |
| e001316 | 19:55:54.252287 | node-6 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/5/reachable |
| e001317 | 19:55:54.252287 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/5/snapshot/views/5/replication/role |
| e001318 | 19:55:54.252287 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/5/snapshot/views/5/replication/connected_slaves |
| e001319 | 19:55:54.252287 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/5/snapshot/views/5/replication/master_link_status |
| e001320 | 19:55:54.252287 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/5/snapshot/views/5/replication/master_sync_in_progress |
| e001321 | 19:55:54.252287 | node-6 | info_replication | replication.master_repl_offset | valid | `28631` | /samples/5/snapshot/views/5/replication/master_repl_offset |
| e001322 | 19:55:54.252287 | node-6 | info_replication | replication.slave_repl_offset | valid | `28631` | /samples/5/snapshot/views/5/replication/slave_repl_offset |
| e001323 | 19:55:54.252287 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /samples/5/snapshot/views/5/redis_id |
| e001324 | 19:55:54.252287 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/5/coverage |
| e001325 | 19:55:54.252287 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/5/nodes/0/flags |
| e001326 | 19:55:54.252287 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/5/snapshot/views/5/nodes/0/primary_id |
| e001327 | 19:55:54.252287 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/0/slots |
| e001328 | 19:55:54.252287 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/0/link_state |
| e001329 | 19:55:54.252287 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/5/nodes/1/flags |
| e001330 | 19:55:54.252287 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/1/primary_id |
| e001331 | 19:55:54.252287 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/5/nodes/1/slots |
| e001332 | 19:55:54.252287 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/1/link_state |
| e001333 | 19:55:54.252287 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/5/snapshot/views/5/nodes/2/flags |
| e001334 | 19:55:54.252287 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/2/primary_id |
| e001335 | 19:55:54.252287 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/2/slots |
| e001336 | 19:55:54.252287 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/2/link_state |
| e001337 | 19:55:54.252287 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/5/nodes/3/flags |
| e001338 | 19:55:54.252287 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/3/primary_id |
| e001339 | 19:55:54.252287 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/5/nodes/3/slots |
| e001340 | 19:55:54.252287 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/3/link_state |
| e001341 | 19:55:54.252287 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/5/snapshot/views/5/nodes/4/flags |
| e001342 | 19:55:54.252287 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/5/snapshot/views/5/nodes/4/primary_id |
| e001343 | 19:55:54.252287 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/4/slots |
| e001344 | 19:55:54.252287 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/4/link_state |
| e001345 | 19:55:54.252287 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/5/nodes/5/flags |
| e001346 | 19:55:54.252287 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/5/primary_id |
| e001347 | 19:55:54.252287 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/5/nodes/5/slots |
| e001348 | 19:55:54.252287 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/5/link_state |
| e001349 | 19:55:54.252287 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/5/persistence/aof_enabled |
| e001350 | 19:55:54.252287 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/5/persistence/loading |
| e001351 | 19:55:54.252287 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/5/persistence/aof_last_write_status |
| e001352 | 19:55:54.252287 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e001353 | 19:55:54.252287 | node-6 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/5/process_live |
| e001354 | 19:55:54.252289 | node-3 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/2/reachable |
| e001355 | 19:55:54.252289 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/5/snapshot/views/2/replication/role |
| e001356 | 19:55:54.252289 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/5/snapshot/views/2/replication/connected_slaves |
| e001357 | 19:55:54.252289 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/5/snapshot/views/2/replication/master_link_status |
| e001358 | 19:55:54.252289 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/5/snapshot/views/2/replication/master_sync_in_progress |
| e001359 | 19:55:54.252289 | node-3 | info_replication | replication.master_repl_offset | valid | `28409` | /samples/5/snapshot/views/2/replication/master_repl_offset |
| e001360 | 19:55:54.252289 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/5/snapshot/views/2/replication/slave_repl_offset |
| e001361 | 19:55:54.252289 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/5/snapshot/views/2/redis_id |
| e001362 | 19:55:54.252289 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/2/coverage |
| e001363 | 19:55:54.252289 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/5/snapshot/views/2/nodes/0/flags |
| e001364 | 19:55:54.252289 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/0/primary_id |
| e001365 | 19:55:54.252289 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/2/nodes/0/slots |
| e001366 | 19:55:54.252289 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/0/link_state |
| e001367 | 19:55:54.252289 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/5/snapshot/views/2/nodes/1/flags |
| e001368 | 19:55:54.252289 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/1/primary_id |
| e001369 | 19:55:54.252289 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/1/slots |
| e001370 | 19:55:54.252289 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/1/link_state |
| e001371 | 19:55:54.252289 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/2/nodes/2/flags |
| e001372 | 19:55:54.252289 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/5/snapshot/views/2/nodes/2/primary_id |
| e001373 | 19:55:54.252289 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/2/slots |
| e001374 | 19:55:54.252289 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/2/link_state |
| e001375 | 19:55:54.252289 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/2/nodes/3/flags |
| e001376 | 19:55:54.252289 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/3/primary_id |
| e001377 | 19:55:54.252289 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/2/nodes/3/slots |
| e001378 | 19:55:54.252289 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/3/link_state |
| e001379 | 19:55:54.252289 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/2/nodes/4/flags |
| e001380 | 19:55:54.252289 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/4/primary_id |
| e001381 | 19:55:54.252289 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/2/nodes/4/slots |
| e001382 | 19:55:54.252289 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/4/link_state |
| e001383 | 19:55:54.252289 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/2/nodes/5/flags |
| e001384 | 19:55:54.252289 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/5/snapshot/views/2/nodes/5/primary_id |
| e001385 | 19:55:54.252289 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/5/slots |
| e001386 | 19:55:54.252289 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/5/link_state |
| e001387 | 19:55:54.252289 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/2/persistence/aof_enabled |
| e001388 | 19:55:54.252289 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/2/persistence/loading |
| e001389 | 19:55:54.252289 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/2/persistence/aof_last_write_status |
| e001390 | 19:55:54.252289 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001391 | 19:55:54.252289 | node-3 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/2/process_live |
| e001392 | 19:55:54.252304 | node-1 | node_client | reachable | valid | `false` | /samples/5/snapshot/views/0/reachable |
| e001393 | 19:55:54.252304 | node-1 | info_replication | replication.role | collection_failed | `` | /samples/5/snapshot/views/0/replication/role |
| e001394 | 19:55:54.252304 | node-1 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/5/snapshot/views/0/replication/connected_slaves |
| e001395 | 19:55:54.252304 | node-1 | info_replication | replication.master_link_status | collection_failed | `` | /samples/5/snapshot/views/0/replication/master_link_status |
| e001396 | 19:55:54.252304 | node-1 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/5/snapshot/views/0/replication/master_sync_in_progress |
| e001397 | 19:55:54.252304 | node-1 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/5/snapshot/views/0/replication/master_repl_offset |
| e001398 | 19:55:54.252304 | node-1 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/5/snapshot/views/0/replication/slave_repl_offset |
| e001399 | 19:55:54.252304 | node-1 | cluster_nodes | redis_id | collection_failed | `` | /samples/5/snapshot/views/0/redis_id |
| e001400 | 19:55:54.252304 | node-1 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/5/snapshot/views/0/coverage |
| e001401 | 19:55:54.252304 | node-1 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/5/snapshot/views/0/persistence/aof_enabled |
| e001402 | 19:55:54.252304 | node-1 | info_persistence | persistence.loading | collection_failed | `` | /samples/5/snapshot/views/0/persistence/loading |
| e001403 | 19:55:54.252304 | node-1 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/5/snapshot/views/0/persistence/aof_last_write_status |
| e001404 | 19:55:54.252304 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/5/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001405 | 19:55:54.252304 | node-1 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/0/process_live |
| e001406 | 19:55:54.252333 | node-4 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/3/reachable |
| e001407 | 19:55:54.252333 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/5/snapshot/views/3/replication/role |
| e001408 | 19:55:54.252333 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/5/snapshot/views/3/replication/connected_slaves |
| e001409 | 19:55:54.252333 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/5/snapshot/views/3/replication/master_link_status |
| e001410 | 19:55:54.252333 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/5/snapshot/views/3/replication/master_sync_in_progress |
| e001411 | 19:55:54.252333 | node-4 | info_replication | replication.master_repl_offset | valid | `28409` | /samples/5/snapshot/views/3/replication/master_repl_offset |
| e001412 | 19:55:54.252333 | node-4 | info_replication | replication.slave_repl_offset | valid | `28409` | /samples/5/snapshot/views/3/replication/slave_repl_offset |
| e001413 | 19:55:54.252333 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /samples/5/snapshot/views/3/redis_id |
| e001414 | 19:55:54.252333 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/3/coverage |
| e001415 | 19:55:54.252333 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/3/nodes/0/flags |
| e001416 | 19:55:54.252333 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/0/primary_id |
| e001417 | 19:55:54.252333 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/3/nodes/0/slots |
| e001418 | 19:55:54.252333 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/0/link_state |
| e001419 | 19:55:54.252333 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/5/snapshot/views/3/nodes/1/flags |
| e001420 | 19:55:54.252333 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/5/snapshot/views/3/nodes/1/primary_id |
| e001421 | 19:55:54.252333 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/1/slots |
| e001422 | 19:55:54.252333 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/1/link_state |
| e001423 | 19:55:54.252333 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/3/nodes/2/flags |
| e001424 | 19:55:54.252333 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/2/primary_id |
| e001425 | 19:55:54.252333 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/3/nodes/2/slots |
| e001426 | 19:55:54.252333 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/2/link_state |
| e001427 | 19:55:54.252333 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/3/nodes/3/flags |
| e001428 | 19:55:54.252333 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/5/snapshot/views/3/nodes/3/primary_id |
| e001429 | 19:55:54.252333 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/3/slots |
| e001430 | 19:55:54.252333 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/3/link_state |
| e001431 | 19:55:54.252333 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/3/nodes/4/flags |
| e001432 | 19:55:54.252333 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/4/primary_id |
| e001433 | 19:55:54.252333 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/3/nodes/4/slots |
| e001434 | 19:55:54.252333 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/4/link_state |
| e001435 | 19:55:54.252333 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/5/snapshot/views/3/nodes/5/flags |
| e001436 | 19:55:54.252333 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/5/primary_id |
| e001437 | 19:55:54.252333 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/5/slots |
| e001438 | 19:55:54.252333 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/5/link_state |
| e001439 | 19:55:54.252333 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/3/persistence/aof_enabled |
| e001440 | 19:55:54.252333 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/3/persistence/loading |
| e001441 | 19:55:54.252333 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/3/persistence/aof_last_write_status |
| e001442 | 19:55:54.252333 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001443 | 19:55:54.252333 | node-4 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/3/process_live |
| e001444 | 19:55:54.252344 | node-5 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/4/reachable |
| e001445 | 19:55:54.252344 | node-5 | info_replication | replication.role | valid | `"master"` | /samples/5/snapshot/views/4/replication/role |
| e001446 | 19:55:54.252344 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/5/snapshot/views/4/replication/connected_slaves |
| e001447 | 19:55:54.252344 | node-5 | info_replication | replication.master_link_status | not_applicable | `` | /samples/5/snapshot/views/4/replication/master_link_status |
| e001448 | 19:55:54.252344 | node-5 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/5/snapshot/views/4/replication/master_sync_in_progress |
| e001449 | 19:55:54.252344 | node-5 | info_replication | replication.master_repl_offset | valid | `21278` | /samples/5/snapshot/views/4/replication/master_repl_offset |
| e001450 | 19:55:54.252344 | node-5 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/5/snapshot/views/4/replication/slave_repl_offset |
| e001451 | 19:55:54.252344 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /samples/5/snapshot/views/4/redis_id |
| e001452 | 19:55:54.252344 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/4/coverage |
| e001453 | 19:55:54.252344 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/4/nodes/0/flags |
| e001454 | 19:55:54.252344 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/5/snapshot/views/4/nodes/0/primary_id |
| e001455 | 19:55:54.252344 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/0/slots |
| e001456 | 19:55:54.252344 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/0/link_state |
| e001457 | 19:55:54.252344 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/4/nodes/1/flags |
| e001458 | 19:55:54.252344 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/1/primary_id |
| e001459 | 19:55:54.252344 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/4/nodes/1/slots |
| e001460 | 19:55:54.252344 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/1/link_state |
| e001461 | 19:55:54.252344 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/5/snapshot/views/4/nodes/2/flags |
| e001462 | 19:55:54.252344 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/2/primary_id |
| e001463 | 19:55:54.252344 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/2/slots |
| e001464 | 19:55:54.252344 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/2/link_state |
| e001465 | 19:55:54.252344 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/5/snapshot/views/4/nodes/3/flags |
| e001466 | 19:55:54.252344 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/3/primary_id |
| e001467 | 19:55:54.252344 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/4/nodes/3/slots |
| e001468 | 19:55:54.252344 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/3/link_state |
| e001469 | 19:55:54.252344 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/4/nodes/4/flags |
| e001470 | 19:55:54.252344 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/5/snapshot/views/4/nodes/4/primary_id |
| e001471 | 19:55:54.252344 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/4/slots |
| e001472 | 19:55:54.252344 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/4/link_state |
| e001473 | 19:55:54.252344 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/4/nodes/5/flags |
| e001474 | 19:55:54.252344 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/5/primary_id |
| e001475 | 19:55:54.252344 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/4/nodes/5/slots |
| e001476 | 19:55:54.252344 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/5/link_state |
| e001477 | 19:55:54.252344 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/4/persistence/aof_enabled |
| e001478 | 19:55:54.252344 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/4/persistence/loading |
| e001479 | 19:55:54.252344 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/4/persistence/aof_last_write_status |
| e001480 | 19:55:54.252344 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001481 | 19:55:54.252344 | node-5 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/4/process_live |
| e001482 | 19:55:54.388668 |  | client_probe | client_success | valid | `true` | /samples/5/probe/successful |
| e001483 | 19:55:54.699046 | node-3 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/2/reachable |
| e001484 | 19:55:54.699046 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/6/snapshot/views/2/replication/role |
| e001485 | 19:55:54.699046 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/6/snapshot/views/2/replication/connected_slaves |
| e001486 | 19:55:54.699046 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/6/snapshot/views/2/replication/master_link_status |
| e001487 | 19:55:54.699046 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/6/snapshot/views/2/replication/master_sync_in_progress |
| e001488 | 19:55:54.699046 | node-3 | info_replication | replication.master_repl_offset | valid | `32051` | /samples/6/snapshot/views/2/replication/master_repl_offset |
| e001489 | 19:55:54.699046 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/6/snapshot/views/2/replication/slave_repl_offset |
| e001490 | 19:55:54.699046 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/6/snapshot/views/2/redis_id |
| e001491 | 19:55:54.699046 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/2/coverage |
| e001492 | 19:55:54.699046 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/6/snapshot/views/2/nodes/0/flags |
| e001493 | 19:55:54.699046 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/0/primary_id |
| e001494 | 19:55:54.699046 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/2/nodes/0/slots |
| e001495 | 19:55:54.699046 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/0/link_state |
| e001496 | 19:55:54.699046 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/6/snapshot/views/2/nodes/1/flags |
| e001497 | 19:55:54.699046 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/1/primary_id |
| e001498 | 19:55:54.699046 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/1/slots |
| e001499 | 19:55:54.699046 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/1/link_state |
| e001500 | 19:55:54.699046 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/2/nodes/2/flags |
| e001501 | 19:55:54.699046 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/6/snapshot/views/2/nodes/2/primary_id |
| e001502 | 19:55:54.699046 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/2/slots |
| e001503 | 19:55:54.699046 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/2/link_state |
| e001504 | 19:55:54.699046 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/2/nodes/3/flags |
| e001505 | 19:55:54.699046 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/3/primary_id |
| e001506 | 19:55:54.699046 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/2/nodes/3/slots |
| e001507 | 19:55:54.699046 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/3/link_state |
| e001508 | 19:55:54.699046 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/2/nodes/4/flags |
| e001509 | 19:55:54.699046 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/4/primary_id |
| e001510 | 19:55:54.699046 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/2/nodes/4/slots |
| e001511 | 19:55:54.699046 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/4/link_state |
| e001512 | 19:55:54.699046 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/2/nodes/5/flags |
| e001513 | 19:55:54.699046 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/6/snapshot/views/2/nodes/5/primary_id |
| e001514 | 19:55:54.699046 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/5/slots |
| e001515 | 19:55:54.699046 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/5/link_state |
| e001516 | 19:55:54.699046 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/2/persistence/aof_enabled |
| e001517 | 19:55:54.699046 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/2/persistence/loading |
| e001518 | 19:55:54.699046 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/2/persistence/aof_last_write_status |
| e001519 | 19:55:54.699046 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001520 | 19:55:54.699046 | node-3 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/2/process_live |
| e001521 | 19:55:54.699048 | node-6 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/5/reachable |
| e001522 | 19:55:54.699048 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/6/snapshot/views/5/replication/role |
| e001523 | 19:55:54.699048 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/6/snapshot/views/5/replication/connected_slaves |
| e001524 | 19:55:54.699048 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/6/snapshot/views/5/replication/master_link_status |
| e001525 | 19:55:54.699048 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/6/snapshot/views/5/replication/master_sync_in_progress |
| e001526 | 19:55:54.699048 | node-6 | info_replication | replication.master_repl_offset | valid | `32051` | /samples/6/snapshot/views/5/replication/master_repl_offset |
| e001527 | 19:55:54.699048 | node-6 | info_replication | replication.slave_repl_offset | valid | `32051` | /samples/6/snapshot/views/5/replication/slave_repl_offset |
| e001528 | 19:55:54.699048 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /samples/6/snapshot/views/5/redis_id |
| e001529 | 19:55:54.699048 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/5/coverage |
| e001530 | 19:55:54.699048 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/5/nodes/0/flags |
| e001531 | 19:55:54.699048 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/6/snapshot/views/5/nodes/0/primary_id |
| e001532 | 19:55:54.699048 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/0/slots |
| e001533 | 19:55:54.699048 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/0/link_state |
| e001534 | 19:55:54.699048 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/5/nodes/1/flags |
| e001535 | 19:55:54.699048 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/1/primary_id |
| e001536 | 19:55:54.699048 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/5/nodes/1/slots |
| e001537 | 19:55:54.699048 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/1/link_state |
| e001538 | 19:55:54.699048 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/6/snapshot/views/5/nodes/2/flags |
| e001539 | 19:55:54.699048 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/2/primary_id |
| e001540 | 19:55:54.699048 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/2/slots |
| e001541 | 19:55:54.699048 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/2/link_state |
| e001542 | 19:55:54.699048 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/5/nodes/3/flags |
| e001543 | 19:55:54.699048 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/3/primary_id |
| e001544 | 19:55:54.699048 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/5/nodes/3/slots |
| e001545 | 19:55:54.699048 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/3/link_state |
| e001546 | 19:55:54.699048 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/6/snapshot/views/5/nodes/4/flags |
| e001547 | 19:55:54.699048 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/6/snapshot/views/5/nodes/4/primary_id |
| e001548 | 19:55:54.699048 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/4/slots |
| e001549 | 19:55:54.699048 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/4/link_state |
| e001550 | 19:55:54.699048 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/5/nodes/5/flags |
| e001551 | 19:55:54.699048 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/5/primary_id |
| e001552 | 19:55:54.699048 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/5/nodes/5/slots |
| e001553 | 19:55:54.699048 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/5/link_state |
| e001554 | 19:55:54.699048 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/5/persistence/aof_enabled |
| e001555 | 19:55:54.699048 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/5/persistence/loading |
| e001556 | 19:55:54.699048 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/5/persistence/aof_last_write_status |
| e001557 | 19:55:54.699048 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e001558 | 19:55:54.699048 | node-6 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/5/process_live |
| e001559 | 19:55:54.699051 | node-1 | node_client | reachable | valid | `false` | /samples/6/snapshot/views/0/reachable |
| e001560 | 19:55:54.699051 | node-1 | info_replication | replication.role | collection_failed | `` | /samples/6/snapshot/views/0/replication/role |
| e001561 | 19:55:54.699051 | node-1 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/6/snapshot/views/0/replication/connected_slaves |
| e001562 | 19:55:54.699051 | node-1 | info_replication | replication.master_link_status | collection_failed | `` | /samples/6/snapshot/views/0/replication/master_link_status |
| e001563 | 19:55:54.699051 | node-1 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/6/snapshot/views/0/replication/master_sync_in_progress |
| e001564 | 19:55:54.699051 | node-1 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/6/snapshot/views/0/replication/master_repl_offset |
| e001565 | 19:55:54.699051 | node-1 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/6/snapshot/views/0/replication/slave_repl_offset |
| e001566 | 19:55:54.699051 | node-1 | cluster_nodes | redis_id | collection_failed | `` | /samples/6/snapshot/views/0/redis_id |
| e001567 | 19:55:54.699051 | node-1 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/6/snapshot/views/0/coverage |
| e001568 | 19:55:54.699051 | node-1 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/6/snapshot/views/0/persistence/aof_enabled |
| e001569 | 19:55:54.699051 | node-1 | info_persistence | persistence.loading | collection_failed | `` | /samples/6/snapshot/views/0/persistence/loading |
| e001570 | 19:55:54.699051 | node-1 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/6/snapshot/views/0/persistence/aof_last_write_status |
| e001571 | 19:55:54.699051 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/6/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001572 | 19:55:54.699051 | node-1 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/0/process_live |
| e001573 | 19:55:54.699061 | node-4 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/3/reachable |
| e001574 | 19:55:54.699061 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/6/snapshot/views/3/replication/role |
| e001575 | 19:55:54.699061 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/6/snapshot/views/3/replication/connected_slaves |
| e001576 | 19:55:54.699061 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/6/snapshot/views/3/replication/master_link_status |
| e001577 | 19:55:54.699061 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/6/snapshot/views/3/replication/master_sync_in_progress |
| e001578 | 19:55:54.699061 | node-4 | info_replication | replication.master_repl_offset | valid | `32051` | /samples/6/snapshot/views/3/replication/master_repl_offset |
| e001579 | 19:55:54.699061 | node-4 | info_replication | replication.slave_repl_offset | valid | `32051` | /samples/6/snapshot/views/3/replication/slave_repl_offset |
| e001580 | 19:55:54.699061 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /samples/6/snapshot/views/3/redis_id |
| e001581 | 19:55:54.699061 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/3/coverage |
| e001582 | 19:55:54.699061 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/3/nodes/0/flags |
| e001583 | 19:55:54.699061 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/0/primary_id |
| e001584 | 19:55:54.699061 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/3/nodes/0/slots |
| e001585 | 19:55:54.699061 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/0/link_state |
| e001586 | 19:55:54.699061 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/6/snapshot/views/3/nodes/1/flags |
| e001587 | 19:55:54.699061 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/6/snapshot/views/3/nodes/1/primary_id |
| e001588 | 19:55:54.699061 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/1/slots |
| e001589 | 19:55:54.699061 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/1/link_state |
| e001590 | 19:55:54.699061 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/3/nodes/2/flags |
| e001591 | 19:55:54.699061 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/2/primary_id |
| e001592 | 19:55:54.699061 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/3/nodes/2/slots |
| e001593 | 19:55:54.699061 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/2/link_state |
| e001594 | 19:55:54.699061 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/3/nodes/3/flags |
| e001595 | 19:55:54.699061 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/6/snapshot/views/3/nodes/3/primary_id |
| e001596 | 19:55:54.699061 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/3/slots |
| e001597 | 19:55:54.699061 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/3/link_state |
| e001598 | 19:55:54.699061 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/3/nodes/4/flags |
| e001599 | 19:55:54.699061 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/4/primary_id |
| e001600 | 19:55:54.699061 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/3/nodes/4/slots |
| e001601 | 19:55:54.699061 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/4/link_state |
| e001602 | 19:55:54.699061 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/6/snapshot/views/3/nodes/5/flags |
| e001603 | 19:55:54.699061 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/5/primary_id |
| e001604 | 19:55:54.699061 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/5/slots |
| e001605 | 19:55:54.699061 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/5/link_state |
| e001606 | 19:55:54.699061 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/3/persistence/aof_enabled |
| e001607 | 19:55:54.699061 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/3/persistence/loading |
| e001608 | 19:55:54.699061 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/3/persistence/aof_last_write_status |
| e001609 | 19:55:54.699061 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001610 | 19:55:54.699061 | node-4 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/3/process_live |
| e001611 | 19:55:54.699088 | node-5 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/4/reachable |
| e001612 | 19:55:54.699088 | node-5 | info_replication | replication.role | valid | `"master"` | /samples/6/snapshot/views/4/replication/role |
| e001613 | 19:55:54.699088 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/6/snapshot/views/4/replication/connected_slaves |
| e001614 | 19:55:54.699088 | node-5 | info_replication | replication.master_link_status | not_applicable | `` | /samples/6/snapshot/views/4/replication/master_link_status |
| e001615 | 19:55:54.699088 | node-5 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/6/snapshot/views/4/replication/master_sync_in_progress |
| e001616 | 19:55:54.699088 | node-5 | info_replication | replication.master_repl_offset | valid | `24698` | /samples/6/snapshot/views/4/replication/master_repl_offset |
| e001617 | 19:55:54.699088 | node-5 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/6/snapshot/views/4/replication/slave_repl_offset |
| e001618 | 19:55:54.699088 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /samples/6/snapshot/views/4/redis_id |
| e001619 | 19:55:54.699088 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/4/coverage |
| e001620 | 19:55:54.699088 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/4/nodes/0/flags |
| e001621 | 19:55:54.699088 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/6/snapshot/views/4/nodes/0/primary_id |
| e001622 | 19:55:54.699088 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/0/slots |
| e001623 | 19:55:54.699088 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/0/link_state |
| e001624 | 19:55:54.699088 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/4/nodes/1/flags |
| e001625 | 19:55:54.699088 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/1/primary_id |
| e001626 | 19:55:54.699088 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/4/nodes/1/slots |
| e001627 | 19:55:54.699088 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/1/link_state |
| e001628 | 19:55:54.699088 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/6/snapshot/views/4/nodes/2/flags |
| e001629 | 19:55:54.699088 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/2/primary_id |
| e001630 | 19:55:54.699088 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/2/slots |
| e001631 | 19:55:54.699088 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/2/link_state |
| e001632 | 19:55:54.699088 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/6/snapshot/views/4/nodes/3/flags |
| e001633 | 19:55:54.699088 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/3/primary_id |
| e001634 | 19:55:54.699088 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/4/nodes/3/slots |
| e001635 | 19:55:54.699088 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/3/link_state |
| e001636 | 19:55:54.699088 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/4/nodes/4/flags |
| e001637 | 19:55:54.699088 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/6/snapshot/views/4/nodes/4/primary_id |
| e001638 | 19:55:54.699088 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/4/slots |
| e001639 | 19:55:54.699088 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/4/link_state |
| e001640 | 19:55:54.699088 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/4/nodes/5/flags |
| e001641 | 19:55:54.699088 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/5/primary_id |
| e001642 | 19:55:54.699088 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/4/nodes/5/slots |
| e001643 | 19:55:54.699088 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/5/link_state |
| e001644 | 19:55:54.699088 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/4/persistence/aof_enabled |
| e001645 | 19:55:54.699088 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/4/persistence/loading |
| e001646 | 19:55:54.699088 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/4/persistence/aof_last_write_status |
| e001647 | 19:55:54.699088 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001648 | 19:55:54.699088 | node-5 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/4/process_live |
| e001649 | 19:55:54.699095 | node-2 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/1/reachable |
| e001650 | 19:55:54.699095 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/6/snapshot/views/1/replication/role |
| e001651 | 19:55:54.699095 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/6/snapshot/views/1/replication/connected_slaves |
| e001652 | 19:55:54.699095 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/6/snapshot/views/1/replication/master_link_status |
| e001653 | 19:55:54.699095 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/6/snapshot/views/1/replication/master_sync_in_progress |
| e001654 | 19:55:54.699095 | node-2 | info_replication | replication.master_repl_offset | valid | `32051` | /samples/6/snapshot/views/1/replication/master_repl_offset |
| e001655 | 19:55:54.699095 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/6/snapshot/views/1/replication/slave_repl_offset |
| e001656 | 19:55:54.699095 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/6/snapshot/views/1/redis_id |
| e001657 | 19:55:54.699095 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/1/coverage |
| e001658 | 19:55:54.699095 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/6/snapshot/views/1/nodes/0/flags |
| e001659 | 19:55:54.699095 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/0/primary_id |
| e001660 | 19:55:54.699095 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/0/slots |
| e001661 | 19:55:54.699095 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/0/link_state |
| e001662 | 19:55:54.699095 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/1/nodes/1/flags |
| e001663 | 19:55:54.699095 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/6/snapshot/views/1/nodes/1/primary_id |
| e001664 | 19:55:54.699095 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/1/slots |
| e001665 | 19:55:54.699095 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/1/link_state |
| e001666 | 19:55:54.699095 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/1/nodes/2/flags |
| e001667 | 19:55:54.699095 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/2/primary_id |
| e001668 | 19:55:54.699095 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/1/nodes/2/slots |
| e001669 | 19:55:54.699095 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/2/link_state |
| e001670 | 19:55:54.699095 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/6/snapshot/views/1/nodes/3/flags |
| e001671 | 19:55:54.699095 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/3/primary_id |
| e001672 | 19:55:54.699095 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/1/nodes/3/slots |
| e001673 | 19:55:54.699095 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/3/link_state |
| e001674 | 19:55:54.699095 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/1/nodes/4/flags |
| e001675 | 19:55:54.699095 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/4/primary_id |
| e001676 | 19:55:54.699095 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/1/nodes/4/slots |
| e001677 | 19:55:54.699095 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/4/link_state |
| e001678 | 19:55:54.699095 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/1/nodes/5/flags |
| e001679 | 19:55:54.699095 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/6/snapshot/views/1/nodes/5/primary_id |
| e001680 | 19:55:54.699095 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/5/slots |
| e001681 | 19:55:54.699095 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/5/link_state |
| e001682 | 19:55:54.699095 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/1/persistence/aof_enabled |
| e001683 | 19:55:54.699095 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/1/persistence/loading |
| e001684 | 19:55:54.699095 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/1/persistence/aof_last_write_status |
| e001685 | 19:55:54.699095 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001686 | 19:55:54.699095 | node-2 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/1/process_live |
| e001687 | 19:55:54.833112 |  | client_probe | client_success | valid | `true` | /samples/6/probe/successful |
| e001688 | 19:55:54.833247 | node-1 | node_client | reachable | valid | `false` | /samples/7/snapshot/views/0/reachable |
| e001689 | 19:55:54.833247 | node-1 | info_replication | replication.role | collection_failed | `` | /samples/7/snapshot/views/0/replication/role |
| e001690 | 19:55:54.833247 | node-1 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/7/snapshot/views/0/replication/connected_slaves |
| e001691 | 19:55:54.833247 | node-1 | info_replication | replication.master_link_status | collection_failed | `` | /samples/7/snapshot/views/0/replication/master_link_status |
| e001692 | 19:55:54.833247 | node-1 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/7/snapshot/views/0/replication/master_sync_in_progress |
| e001693 | 19:55:54.833247 | node-1 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/7/snapshot/views/0/replication/master_repl_offset |
| e001694 | 19:55:54.833247 | node-1 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/7/snapshot/views/0/replication/slave_repl_offset |
| e001695 | 19:55:54.833247 | node-1 | cluster_nodes | redis_id | collection_failed | `` | /samples/7/snapshot/views/0/redis_id |
| e001696 | 19:55:54.833247 | node-1 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/7/snapshot/views/0/coverage |
| e001697 | 19:55:54.833247 | node-1 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/7/snapshot/views/0/persistence/aof_enabled |
| e001698 | 19:55:54.833247 | node-1 | info_persistence | persistence.loading | collection_failed | `` | /samples/7/snapshot/views/0/persistence/loading |
| e001699 | 19:55:54.833247 | node-1 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/7/snapshot/views/0/persistence/aof_last_write_status |
| e001700 | 19:55:54.833247 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/7/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001701 | 19:55:54.833247 | node-1 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/0/process_live |
| e001702 | 19:55:54.83325 | node-2 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/1/reachable |
| e001703 | 19:55:54.83325 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/7/snapshot/views/1/replication/role |
| e001704 | 19:55:54.83325 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/7/snapshot/views/1/replication/connected_slaves |
| e001705 | 19:55:54.83325 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/7/snapshot/views/1/replication/master_link_status |
| e001706 | 19:55:54.83325 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/7/snapshot/views/1/replication/master_sync_in_progress |
| e001707 | 19:55:54.83325 | node-2 | info_replication | replication.master_repl_offset | valid | `33029` | /samples/7/snapshot/views/1/replication/master_repl_offset |
| e001708 | 19:55:54.83325 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/7/snapshot/views/1/replication/slave_repl_offset |
| e001709 | 19:55:54.83325 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/7/snapshot/views/1/redis_id |
| e001710 | 19:55:54.83325 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/1/coverage |
| e001711 | 19:55:54.83325 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/7/snapshot/views/1/nodes/0/flags |
| e001712 | 19:55:54.83325 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/0/primary_id |
| e001713 | 19:55:54.83325 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/0/slots |
| e001714 | 19:55:54.83325 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/0/link_state |
| e001715 | 19:55:54.83325 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/1/nodes/1/flags |
| e001716 | 19:55:54.83325 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/7/snapshot/views/1/nodes/1/primary_id |
| e001717 | 19:55:54.83325 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/1/slots |
| e001718 | 19:55:54.83325 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/1/link_state |
| e001719 | 19:55:54.83325 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/1/nodes/2/flags |
| e001720 | 19:55:54.83325 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/2/primary_id |
| e001721 | 19:55:54.83325 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/1/nodes/2/slots |
| e001722 | 19:55:54.83325 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/2/link_state |
| e001723 | 19:55:54.83325 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/7/snapshot/views/1/nodes/3/flags |
| e001724 | 19:55:54.83325 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/3/primary_id |
| e001725 | 19:55:54.83325 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/1/nodes/3/slots |
| e001726 | 19:55:54.83325 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/3/link_state |
| e001727 | 19:55:54.83325 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/1/nodes/4/flags |
| e001728 | 19:55:54.83325 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/4/primary_id |
| e001729 | 19:55:54.83325 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/1/nodes/4/slots |
| e001730 | 19:55:54.83325 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/4/link_state |
| e001731 | 19:55:54.83325 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/1/nodes/5/flags |
| e001732 | 19:55:54.83325 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/7/snapshot/views/1/nodes/5/primary_id |
| e001733 | 19:55:54.83325 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/5/slots |
| e001734 | 19:55:54.83325 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/5/link_state |
| e001735 | 19:55:54.83325 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/1/persistence/aof_enabled |
| e001736 | 19:55:54.83325 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/1/persistence/loading |
| e001737 | 19:55:54.83325 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/1/persistence/aof_last_write_status |
| e001738 | 19:55:54.83325 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001739 | 19:55:54.83325 | node-2 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/1/process_live |
| e001740 | 19:55:54.833255 | node-6 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/5/reachable |
| e001741 | 19:55:54.833255 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/7/snapshot/views/5/replication/role |
| e001742 | 19:55:54.833255 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/7/snapshot/views/5/replication/connected_slaves |
| e001743 | 19:55:54.833255 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/7/snapshot/views/5/replication/master_link_status |
| e001744 | 19:55:54.833255 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/7/snapshot/views/5/replication/master_sync_in_progress |
| e001745 | 19:55:54.833255 | node-6 | info_replication | replication.master_repl_offset | valid | `33029` | /samples/7/snapshot/views/5/replication/master_repl_offset |
| e001746 | 19:55:54.833255 | node-6 | info_replication | replication.slave_repl_offset | valid | `33029` | /samples/7/snapshot/views/5/replication/slave_repl_offset |
| e001747 | 19:55:54.833255 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /samples/7/snapshot/views/5/redis_id |
| e001748 | 19:55:54.833255 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/5/coverage |
| e001749 | 19:55:54.833255 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/5/nodes/0/flags |
| e001750 | 19:55:54.833255 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/7/snapshot/views/5/nodes/0/primary_id |
| e001751 | 19:55:54.833255 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/0/slots |
| e001752 | 19:55:54.833255 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/0/link_state |
| e001753 | 19:55:54.833255 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/5/nodes/1/flags |
| e001754 | 19:55:54.833255 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/1/primary_id |
| e001755 | 19:55:54.833255 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/5/nodes/1/slots |
| e001756 | 19:55:54.833255 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/1/link_state |
| e001757 | 19:55:54.833255 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/7/snapshot/views/5/nodes/2/flags |
| e001758 | 19:55:54.833255 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/2/primary_id |
| e001759 | 19:55:54.833255 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/2/slots |
| e001760 | 19:55:54.833255 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/2/link_state |
| e001761 | 19:55:54.833255 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/5/nodes/3/flags |
| e001762 | 19:55:54.833255 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/3/primary_id |
| e001763 | 19:55:54.833255 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/5/nodes/3/slots |
| e001764 | 19:55:54.833255 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/3/link_state |
| e001765 | 19:55:54.833255 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/7/snapshot/views/5/nodes/4/flags |
| e001766 | 19:55:54.833255 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/7/snapshot/views/5/nodes/4/primary_id |
| e001767 | 19:55:54.833255 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/4/slots |
| e001768 | 19:55:54.833255 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/4/link_state |
| e001769 | 19:55:54.833255 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/5/nodes/5/flags |
| e001770 | 19:55:54.833255 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/5/primary_id |
| e001771 | 19:55:54.833255 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/5/nodes/5/slots |
| e001772 | 19:55:54.833255 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/5/link_state |
| e001773 | 19:55:54.833255 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/5/persistence/aof_enabled |
| e001774 | 19:55:54.833255 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/5/persistence/loading |
| e001775 | 19:55:54.833255 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/5/persistence/aof_last_write_status |
| e001776 | 19:55:54.833255 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e001777 | 19:55:54.833255 | node-6 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/5/process_live |
| e001778 | 19:55:54.833268 | node-3 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/2/reachable |
| e001779 | 19:55:54.833268 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/7/snapshot/views/2/replication/role |
| e001780 | 19:55:54.833268 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/7/snapshot/views/2/replication/connected_slaves |
| e001781 | 19:55:54.833268 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/7/snapshot/views/2/replication/master_link_status |
| e001782 | 19:55:54.833268 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/7/snapshot/views/2/replication/master_sync_in_progress |
| e001783 | 19:55:54.833268 | node-3 | info_replication | replication.master_repl_offset | valid | `33029` | /samples/7/snapshot/views/2/replication/master_repl_offset |
| e001784 | 19:55:54.833268 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/7/snapshot/views/2/replication/slave_repl_offset |
| e001785 | 19:55:54.833268 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/7/snapshot/views/2/redis_id |
| e001786 | 19:55:54.833268 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/2/coverage |
| e001787 | 19:55:54.833268 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/7/snapshot/views/2/nodes/0/flags |
| e001788 | 19:55:54.833268 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/0/primary_id |
| e001789 | 19:55:54.833268 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/2/nodes/0/slots |
| e001790 | 19:55:54.833268 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/0/link_state |
| e001791 | 19:55:54.833268 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/7/snapshot/views/2/nodes/1/flags |
| e001792 | 19:55:54.833268 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/1/primary_id |
| e001793 | 19:55:54.833268 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/1/slots |
| e001794 | 19:55:54.833268 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/1/link_state |
| e001795 | 19:55:54.833268 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/2/nodes/2/flags |
| e001796 | 19:55:54.833268 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/7/snapshot/views/2/nodes/2/primary_id |
| e001797 | 19:55:54.833268 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/2/slots |
| e001798 | 19:55:54.833268 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/2/link_state |
| e001799 | 19:55:54.833268 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/2/nodes/3/flags |
| e001800 | 19:55:54.833268 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/3/primary_id |
| e001801 | 19:55:54.833268 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/2/nodes/3/slots |
| e001802 | 19:55:54.833268 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/3/link_state |
| e001803 | 19:55:54.833268 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/2/nodes/4/flags |
| e001804 | 19:55:54.833268 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/4/primary_id |
| e001805 | 19:55:54.833268 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/2/nodes/4/slots |
| e001806 | 19:55:54.833268 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/4/link_state |
| e001807 | 19:55:54.833268 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/2/nodes/5/flags |
| e001808 | 19:55:54.833268 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/7/snapshot/views/2/nodes/5/primary_id |
| e001809 | 19:55:54.833268 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/5/slots |
| e001810 | 19:55:54.833268 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/5/link_state |
| e001811 | 19:55:54.833268 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/2/persistence/aof_enabled |
| e001812 | 19:55:54.833268 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/2/persistence/loading |
| e001813 | 19:55:54.833268 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/2/persistence/aof_last_write_status |
| e001814 | 19:55:54.833268 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001815 | 19:55:54.833268 | node-3 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/2/process_live |
| e001816 | 19:55:54.833282 | node-5 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/4/reachable |
| e001817 | 19:55:54.833282 | node-5 | info_replication | replication.role | valid | `"master"` | /samples/7/snapshot/views/4/replication/role |
| e001818 | 19:55:54.833282 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/7/snapshot/views/4/replication/connected_slaves |
| e001819 | 19:55:54.833282 | node-5 | info_replication | replication.master_link_status | not_applicable | `` | /samples/7/snapshot/views/4/replication/master_link_status |
| e001820 | 19:55:54.833282 | node-5 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/7/snapshot/views/4/replication/master_sync_in_progress |
| e001821 | 19:55:54.833282 | node-5 | info_replication | replication.master_repl_offset | valid | `25676` | /samples/7/snapshot/views/4/replication/master_repl_offset |
| e001822 | 19:55:54.833282 | node-5 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/7/snapshot/views/4/replication/slave_repl_offset |
| e001823 | 19:55:54.833282 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /samples/7/snapshot/views/4/redis_id |
| e001824 | 19:55:54.833282 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/4/coverage |
| e001825 | 19:55:54.833282 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/4/nodes/0/flags |
| e001826 | 19:55:54.833282 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/7/snapshot/views/4/nodes/0/primary_id |
| e001827 | 19:55:54.833282 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/0/slots |
| e001828 | 19:55:54.833282 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/0/link_state |
| e001829 | 19:55:54.833282 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/4/nodes/1/flags |
| e001830 | 19:55:54.833282 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/1/primary_id |
| e001831 | 19:55:54.833282 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/4/nodes/1/slots |
| e001832 | 19:55:54.833282 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/1/link_state |
| e001833 | 19:55:54.833282 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/7/snapshot/views/4/nodes/2/flags |
| e001834 | 19:55:54.833282 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/2/primary_id |
| e001835 | 19:55:54.833282 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/2/slots |
| e001836 | 19:55:54.833282 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/2/link_state |
| e001837 | 19:55:54.833282 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/7/snapshot/views/4/nodes/3/flags |
| e001838 | 19:55:54.833282 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/3/primary_id |
| e001839 | 19:55:54.833282 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/4/nodes/3/slots |
| e001840 | 19:55:54.833282 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/3/link_state |
| e001841 | 19:55:54.833282 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/4/nodes/4/flags |
| e001842 | 19:55:54.833282 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/7/snapshot/views/4/nodes/4/primary_id |
| e001843 | 19:55:54.833282 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/4/slots |
| e001844 | 19:55:54.833282 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/4/link_state |
| e001845 | 19:55:54.833282 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/4/nodes/5/flags |
| e001846 | 19:55:54.833282 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/5/primary_id |
| e001847 | 19:55:54.833282 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/4/nodes/5/slots |
| e001848 | 19:55:54.833282 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/5/link_state |
| e001849 | 19:55:54.833282 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/4/persistence/aof_enabled |
| e001850 | 19:55:54.833282 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/4/persistence/loading |
| e001851 | 19:55:54.833282 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/4/persistence/aof_last_write_status |
| e001852 | 19:55:54.833282 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001853 | 19:55:54.833282 | node-5 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/4/process_live |
| e001854 | 19:55:54.833284 | node-4 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/3/reachable |
| e001855 | 19:55:54.833284 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/7/snapshot/views/3/replication/role |
| e001856 | 19:55:54.833284 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/7/snapshot/views/3/replication/connected_slaves |
| e001857 | 19:55:54.833284 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/7/snapshot/views/3/replication/master_link_status |
| e001858 | 19:55:54.833284 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/7/snapshot/views/3/replication/master_sync_in_progress |
| e001859 | 19:55:54.833284 | node-4 | info_replication | replication.master_repl_offset | valid | `33029` | /samples/7/snapshot/views/3/replication/master_repl_offset |
| e001860 | 19:55:54.833284 | node-4 | info_replication | replication.slave_repl_offset | valid | `33029` | /samples/7/snapshot/views/3/replication/slave_repl_offset |
| e001861 | 19:55:54.833284 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /samples/7/snapshot/views/3/redis_id |
| e001862 | 19:55:54.833284 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/3/coverage |
| e001863 | 19:55:54.833284 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/3/nodes/0/flags |
| e001864 | 19:55:54.833284 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/0/primary_id |
| e001865 | 19:55:54.833284 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/3/nodes/0/slots |
| e001866 | 19:55:54.833284 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/0/link_state |
| e001867 | 19:55:54.833284 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/7/snapshot/views/3/nodes/1/flags |
| e001868 | 19:55:54.833284 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/7/snapshot/views/3/nodes/1/primary_id |
| e001869 | 19:55:54.833284 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/1/slots |
| e001870 | 19:55:54.833284 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/1/link_state |
| e001871 | 19:55:54.833284 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/3/nodes/2/flags |
| e001872 | 19:55:54.833284 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/2/primary_id |
| e001873 | 19:55:54.833284 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/3/nodes/2/slots |
| e001874 | 19:55:54.833284 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/2/link_state |
| e001875 | 19:55:54.833284 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/3/nodes/3/flags |
| e001876 | 19:55:54.833284 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/7/snapshot/views/3/nodes/3/primary_id |
| e001877 | 19:55:54.833284 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/3/slots |
| e001878 | 19:55:54.833284 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/3/link_state |
| e001879 | 19:55:54.833284 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/3/nodes/4/flags |
| e001880 | 19:55:54.833284 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/4/primary_id |
| e001881 | 19:55:54.833284 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/3/nodes/4/slots |
| e001882 | 19:55:54.833284 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/4/link_state |
| e001883 | 19:55:54.833284 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/7/snapshot/views/3/nodes/5/flags |
| e001884 | 19:55:54.833284 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/5/primary_id |
| e001885 | 19:55:54.833284 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/5/slots |
| e001886 | 19:55:54.833284 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/5/link_state |
| e001887 | 19:55:54.833284 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/3/persistence/aof_enabled |
| e001888 | 19:55:54.833284 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/3/persistence/loading |
| e001889 | 19:55:54.833284 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/3/persistence/aof_last_write_status |
| e001890 | 19:55:54.833284 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001891 | 19:55:54.833284 | node-4 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/3/process_live |
| e001892 | 19:55:54.967498 |  | client_probe | client_success | valid | `true` | /samples/7/probe/successful |
| e001893 | 19:55:55.224589 |  | client_window | client.attempts | valid | `170` | /metrics/windows/6/attempts |
| e001894 | 19:55:55.224589 |  | client_window | client.errors | valid | `0` | /metrics/windows/6/errors |
| e001895 | 19:55:55.224589 |  | client_window | client.successes | valid | `170` | /metrics/windows/6/successes |
| e001896 | 19:55:55.224589 |  | client_window | client.p95_ms | valid | `0.663814` | /metrics/windows/6/p95_ms |
| e001897 | 19:55:55.278164 | node-1 | node_client | reachable | valid | `false` | /samples/8/snapshot/views/0/reachable |
| e001898 | 19:55:55.278164 | node-1 | info_replication | replication.role | collection_failed | `` | /samples/8/snapshot/views/0/replication/role |
| e001899 | 19:55:55.278164 | node-1 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/8/snapshot/views/0/replication/connected_slaves |
| e001900 | 19:55:55.278164 | node-1 | info_replication | replication.master_link_status | collection_failed | `` | /samples/8/snapshot/views/0/replication/master_link_status |
| e001901 | 19:55:55.278164 | node-1 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/8/snapshot/views/0/replication/master_sync_in_progress |
| e001902 | 19:55:55.278164 | node-1 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/8/snapshot/views/0/replication/master_repl_offset |
| e001903 | 19:55:55.278164 | node-1 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/8/snapshot/views/0/replication/slave_repl_offset |
| e001904 | 19:55:55.278164 | node-1 | cluster_nodes | redis_id | collection_failed | `` | /samples/8/snapshot/views/0/redis_id |
| e001905 | 19:55:55.278164 | node-1 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/8/snapshot/views/0/coverage |
| e001906 | 19:55:55.278164 | node-1 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/8/snapshot/views/0/persistence/aof_enabled |
| e001907 | 19:55:55.278164 | node-1 | info_persistence | persistence.loading | collection_failed | `` | /samples/8/snapshot/views/0/persistence/loading |
| e001908 | 19:55:55.278164 | node-1 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/8/snapshot/views/0/persistence/aof_last_write_status |
| e001909 | 19:55:55.278164 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/8/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001910 | 19:55:55.278164 | node-1 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/0/process_live |
| e001911 | 19:55:55.278164 | node-1 | node_client | reachable | valid | `false` | /final/views/0/reachable |
| e001912 | 19:55:55.278164 | node-1 | info_replication | replication.role | collection_failed | `` | /final/views/0/replication/role |
| e001913 | 19:55:55.278164 | node-1 | info_replication | replication.connected_slaves | collection_failed | `` | /final/views/0/replication/connected_slaves |
| e001914 | 19:55:55.278164 | node-1 | info_replication | replication.master_link_status | collection_failed | `` | /final/views/0/replication/master_link_status |
| e001915 | 19:55:55.278164 | node-1 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /final/views/0/replication/master_sync_in_progress |
| e001916 | 19:55:55.278164 | node-1 | info_replication | replication.master_repl_offset | collection_failed | `` | /final/views/0/replication/master_repl_offset |
| e001917 | 19:55:55.278164 | node-1 | info_replication | replication.slave_repl_offset | collection_failed | `` | /final/views/0/replication/slave_repl_offset |
| e001918 | 19:55:55.278164 | node-1 | cluster_nodes | redis_id | collection_failed | `` | /final/views/0/redis_id |
| e001919 | 19:55:55.278164 | node-1 | cluster_nodes | slot_coverage | collection_failed | `` | /final/views/0/coverage |
| e001920 | 19:55:55.278164 | node-1 | info_persistence | persistence.aof_enabled | collection_failed | `` | /final/views/0/persistence/aof_enabled |
| e001921 | 19:55:55.278164 | node-1 | info_persistence | persistence.loading | collection_failed | `` | /final/views/0/persistence/loading |
| e001922 | 19:55:55.278164 | node-1 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /final/views/0/persistence/aof_last_write_status |
| e001923 | 19:55:55.278164 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /final/views/0/persistence/aof_rewrite_in_progress |
| e001924 | 19:55:55.278164 | node-1 | admin | process_live | missing_from_schema | `` | /final/views/0/process_live |
| e001925 | 19:55:55.278173 | node-3 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/2/reachable |
| e001926 | 19:55:55.278173 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/8/snapshot/views/2/replication/role |
| e001927 | 19:55:55.278173 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/8/snapshot/views/2/replication/connected_slaves |
| e001928 | 19:55:55.278173 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/8/snapshot/views/2/replication/master_link_status |
| e001929 | 19:55:55.278173 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/8/snapshot/views/2/replication/master_sync_in_progress |
| e001930 | 19:55:55.278173 | node-3 | info_replication | replication.master_repl_offset | valid | `36227` | /samples/8/snapshot/views/2/replication/master_repl_offset |
| e001931 | 19:55:55.278173 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/8/snapshot/views/2/replication/slave_repl_offset |
| e001932 | 19:55:55.278173 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/8/snapshot/views/2/redis_id |
| e001933 | 19:55:55.278173 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/2/coverage |
| e001934 | 19:55:55.278173 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/8/snapshot/views/2/nodes/0/flags |
| e001935 | 19:55:55.278173 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/0/primary_id |
| e001936 | 19:55:55.278173 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/2/nodes/0/slots |
| e001937 | 19:55:55.278173 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/0/link_state |
| e001938 | 19:55:55.278173 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/8/snapshot/views/2/nodes/1/flags |
| e001939 | 19:55:55.278173 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/1/primary_id |
| e001940 | 19:55:55.278173 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/1/slots |
| e001941 | 19:55:55.278173 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/1/link_state |
| e001942 | 19:55:55.278173 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/2/nodes/2/flags |
| e001943 | 19:55:55.278173 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/8/snapshot/views/2/nodes/2/primary_id |
| e001944 | 19:55:55.278173 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/2/slots |
| e001945 | 19:55:55.278173 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/2/link_state |
| e001946 | 19:55:55.278173 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/2/nodes/3/flags |
| e001947 | 19:55:55.278173 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/3/primary_id |
| e001948 | 19:55:55.278173 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/2/nodes/3/slots |
| e001949 | 19:55:55.278173 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/3/link_state |
| e001950 | 19:55:55.278173 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/2/nodes/4/flags |
| e001951 | 19:55:55.278173 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/4/primary_id |
| e001952 | 19:55:55.278173 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/2/nodes/4/slots |
| e001953 | 19:55:55.278173 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/4/link_state |
| e001954 | 19:55:55.278173 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/2/nodes/5/flags |
| e001955 | 19:55:55.278173 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/8/snapshot/views/2/nodes/5/primary_id |
| e001956 | 19:55:55.278173 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/5/slots |
| e001957 | 19:55:55.278173 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/5/link_state |
| e001958 | 19:55:55.278173 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/2/persistence/aof_enabled |
| e001959 | 19:55:55.278173 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/2/persistence/loading |
| e001960 | 19:55:55.278173 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/2/persistence/aof_last_write_status |
| e001961 | 19:55:55.278173 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001962 | 19:55:55.278173 | node-3 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/2/process_live |
| e001963 | 19:55:55.278173 | node-3 | node_client | reachable | valid | `true` | /final/views/2/reachable |
| e001964 | 19:55:55.278173 | node-3 | info_replication | replication.role | valid | `"master"` | /final/views/2/replication/role |
| e001965 | 19:55:55.278173 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /final/views/2/replication/connected_slaves |
| e001966 | 19:55:55.278173 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /final/views/2/replication/master_link_status |
| e001967 | 19:55:55.278173 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /final/views/2/replication/master_sync_in_progress |
| e001968 | 19:55:55.278173 | node-3 | info_replication | replication.master_repl_offset | valid | `36227` | /final/views/2/replication/master_repl_offset |
| e001969 | 19:55:55.278173 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /final/views/2/replication/slave_repl_offset |
| e001970 | 19:55:55.278173 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /final/views/2/redis_id |
| e001971 | 19:55:55.278173 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /final/views/2/coverage |
| e001972 | 19:55:55.278173 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /final/views/2/nodes/0/flags |
| e001973 | 19:55:55.278173 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /final/views/2/nodes/0/primary_id |
| e001974 | 19:55:55.278173 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/2/nodes/0/slots |
| e001975 | 19:55:55.278173 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/0/link_state |
| e001976 | 19:55:55.278173 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
              "master",
              "fail"
            ]` | /final/views/2/nodes/1/flags |
| e001977 | 19:55:55.278173 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /final/views/2/nodes/1/primary_id |
| e001978 | 19:55:55.278173 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /final/views/2/nodes/1/slots |
| e001979 | 19:55:55.278173 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/1/link_state |
| e001980 | 19:55:55.278173 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/2/nodes/2/flags |
| e001981 | 19:55:55.278173 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /final/views/2/nodes/2/primary_id |
| e001982 | 19:55:55.278173 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /final/views/2/nodes/2/slots |
| e001983 | 19:55:55.278173 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/2/link_state |
| e001984 | 19:55:55.278173 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /final/views/2/nodes/3/flags |
| e001985 | 19:55:55.278173 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /final/views/2/nodes/3/primary_id |
| e001986 | 19:55:55.278173 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/2/nodes/3/slots |
| e001987 | 19:55:55.278173 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/3/link_state |
| e001988 | 19:55:55.278173 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /final/views/2/nodes/4/flags |
| e001989 | 19:55:55.278173 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /final/views/2/nodes/4/primary_id |
| e001990 | 19:55:55.278173 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/2/nodes/4/slots |
| e001991 | 19:55:55.278173 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/4/link_state |
| e001992 | 19:55:55.278173 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/2/nodes/5/flags |
| e001993 | 19:55:55.278173 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /final/views/2/nodes/5/primary_id |
| e001994 | 19:55:55.278173 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /final/views/2/nodes/5/slots |
| e001995 | 19:55:55.278173 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/5/link_state |
| e001996 | 19:55:55.278173 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/2/persistence/aof_enabled |
| e001997 | 19:55:55.278173 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /final/views/2/persistence/loading |
| e001998 | 19:55:55.278173 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/2/persistence/aof_last_write_status |
| e001999 | 19:55:55.278173 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/2/persistence/aof_rewrite_in_progress |
| e002000 | 19:55:55.278173 | node-3 | admin | process_live | missing_from_schema | `` | /final/views/2/process_live |
| e002001 | 19:55:55.278176 | node-5 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/4/reachable |
| e002002 | 19:55:55.278176 | node-5 | info_replication | replication.role | valid | `"master"` | /samples/8/snapshot/views/4/replication/role |
| e002003 | 19:55:55.278176 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/8/snapshot/views/4/replication/connected_slaves |
| e002004 | 19:55:55.278176 | node-5 | info_replication | replication.master_link_status | not_applicable | `` | /samples/8/snapshot/views/4/replication/master_link_status |
| e002005 | 19:55:55.278176 | node-5 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/8/snapshot/views/4/replication/master_sync_in_progress |
| e002006 | 19:55:55.278176 | node-5 | info_replication | replication.master_repl_offset | valid | `29096` | /samples/8/snapshot/views/4/replication/master_repl_offset |
| e002007 | 19:55:55.278176 | node-5 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/8/snapshot/views/4/replication/slave_repl_offset |
| e002008 | 19:55:55.278176 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /samples/8/snapshot/views/4/redis_id |
| e002009 | 19:55:55.278176 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/4/coverage |
| e002010 | 19:55:55.278176 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/4/nodes/0/flags |
| e002011 | 19:55:55.278176 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/8/snapshot/views/4/nodes/0/primary_id |
| e002012 | 19:55:55.278176 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/0/slots |
| e002013 | 19:55:55.278176 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/0/link_state |
| e002014 | 19:55:55.278176 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/4/nodes/1/flags |
| e002015 | 19:55:55.278176 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/1/primary_id |
| e002016 | 19:55:55.278176 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/4/nodes/1/slots |
| e002017 | 19:55:55.278176 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/1/link_state |
| e002018 | 19:55:55.278176 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/8/snapshot/views/4/nodes/2/flags |
| e002019 | 19:55:55.278176 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/2/primary_id |
| e002020 | 19:55:55.278176 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/2/slots |
| e002021 | 19:55:55.278176 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/2/link_state |
| e002022 | 19:55:55.278176 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/8/snapshot/views/4/nodes/3/flags |
| e002023 | 19:55:55.278176 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/3/primary_id |
| e002024 | 19:55:55.278176 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/4/nodes/3/slots |
| e002025 | 19:55:55.278176 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/3/link_state |
| e002026 | 19:55:55.278176 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/4/nodes/4/flags |
| e002027 | 19:55:55.278176 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/8/snapshot/views/4/nodes/4/primary_id |
| e002028 | 19:55:55.278176 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/4/slots |
| e002029 | 19:55:55.278176 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/4/link_state |
| e002030 | 19:55:55.278176 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/4/nodes/5/flags |
| e002031 | 19:55:55.278176 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/5/primary_id |
| e002032 | 19:55:55.278176 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/4/nodes/5/slots |
| e002033 | 19:55:55.278176 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/5/link_state |
| e002034 | 19:55:55.278176 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/4/persistence/aof_enabled |
| e002035 | 19:55:55.278176 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/4/persistence/loading |
| e002036 | 19:55:55.278176 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/4/persistence/aof_last_write_status |
| e002037 | 19:55:55.278176 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e002038 | 19:55:55.278176 | node-5 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/4/process_live |
| e002039 | 19:55:55.278176 | node-5 | node_client | reachable | valid | `true` | /final/views/4/reachable |
| e002040 | 19:55:55.278176 | node-5 | info_replication | replication.role | valid | `"master"` | /final/views/4/replication/role |
| e002041 | 19:55:55.278176 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /final/views/4/replication/connected_slaves |
| e002042 | 19:55:55.278176 | node-5 | info_replication | replication.master_link_status | not_applicable | `` | /final/views/4/replication/master_link_status |
| e002043 | 19:55:55.278176 | node-5 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /final/views/4/replication/master_sync_in_progress |
| e002044 | 19:55:55.278176 | node-5 | info_replication | replication.master_repl_offset | valid | `29096` | /final/views/4/replication/master_repl_offset |
| e002045 | 19:55:55.278176 | node-5 | info_replication | replication.slave_repl_offset | not_applicable | `` | /final/views/4/replication/slave_repl_offset |
| e002046 | 19:55:55.278176 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /final/views/4/redis_id |
| e002047 | 19:55:55.278176 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /final/views/4/coverage |
| e002048 | 19:55:55.278176 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/4/nodes/0/flags |
| e002049 | 19:55:55.278176 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /final/views/4/nodes/0/primary_id |
| e002050 | 19:55:55.278176 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /final/views/4/nodes/0/slots |
| e002051 | 19:55:55.278176 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/0/link_state |
| e002052 | 19:55:55.278176 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /final/views/4/nodes/1/flags |
| e002053 | 19:55:55.278176 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /final/views/4/nodes/1/primary_id |
| e002054 | 19:55:55.278176 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/4/nodes/1/slots |
| e002055 | 19:55:55.278176 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/1/link_state |
| e002056 | 19:55:55.278176 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
              "master",
              "fail"
            ]` | /final/views/4/nodes/2/flags |
| e002057 | 19:55:55.278176 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /final/views/4/nodes/2/primary_id |
| e002058 | 19:55:55.278176 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /final/views/4/nodes/2/slots |
| e002059 | 19:55:55.278176 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/2/link_state |
| e002060 | 19:55:55.278176 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /final/views/4/nodes/3/flags |
| e002061 | 19:55:55.278176 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /final/views/4/nodes/3/primary_id |
| e002062 | 19:55:55.278176 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/4/nodes/3/slots |
| e002063 | 19:55:55.278176 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/3/link_state |
| e002064 | 19:55:55.278176 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/4/nodes/4/flags |
| e002065 | 19:55:55.278176 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /final/views/4/nodes/4/primary_id |
| e002066 | 19:55:55.278176 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /final/views/4/nodes/4/slots |
| e002067 | 19:55:55.278176 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/4/link_state |
| e002068 | 19:55:55.278176 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /final/views/4/nodes/5/flags |
| e002069 | 19:55:55.278176 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /final/views/4/nodes/5/primary_id |
| e002070 | 19:55:55.278176 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/4/nodes/5/slots |
| e002071 | 19:55:55.278176 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/5/link_state |
| e002072 | 19:55:55.278176 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/4/persistence/aof_enabled |
| e002073 | 19:55:55.278176 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /final/views/4/persistence/loading |
| e002074 | 19:55:55.278176 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/4/persistence/aof_last_write_status |
| e002075 | 19:55:55.278176 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/4/persistence/aof_rewrite_in_progress |
| e002076 | 19:55:55.278176 | node-5 | admin | process_live | missing_from_schema | `` | /final/views/4/process_live |
| e002077 | 19:55:55.27818 | node-4 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/3/reachable |
| e002078 | 19:55:55.27818 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/8/snapshot/views/3/replication/role |
| e002079 | 19:55:55.27818 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/8/snapshot/views/3/replication/connected_slaves |
| e002080 | 19:55:55.27818 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/8/snapshot/views/3/replication/master_link_status |
| e002081 | 19:55:55.27818 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/8/snapshot/views/3/replication/master_sync_in_progress |
| e002082 | 19:55:55.27818 | node-4 | info_replication | replication.master_repl_offset | valid | `36227` | /samples/8/snapshot/views/3/replication/master_repl_offset |
| e002083 | 19:55:55.27818 | node-4 | info_replication | replication.slave_repl_offset | valid | `36227` | /samples/8/snapshot/views/3/replication/slave_repl_offset |
| e002084 | 19:55:55.27818 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /samples/8/snapshot/views/3/redis_id |
| e002085 | 19:55:55.27818 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/3/coverage |
| e002086 | 19:55:55.27818 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/3/nodes/0/flags |
| e002087 | 19:55:55.27818 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/0/primary_id |
| e002088 | 19:55:55.27818 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/3/nodes/0/slots |
| e002089 | 19:55:55.27818 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/0/link_state |
| e002090 | 19:55:55.27818 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/8/snapshot/views/3/nodes/1/flags |
| e002091 | 19:55:55.27818 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/8/snapshot/views/3/nodes/1/primary_id |
| e002092 | 19:55:55.27818 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/1/slots |
| e002093 | 19:55:55.27818 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/1/link_state |
| e002094 | 19:55:55.27818 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/3/nodes/2/flags |
| e002095 | 19:55:55.27818 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/2/primary_id |
| e002096 | 19:55:55.27818 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/3/nodes/2/slots |
| e002097 | 19:55:55.27818 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/2/link_state |
| e002098 | 19:55:55.27818 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/3/nodes/3/flags |
| e002099 | 19:55:55.27818 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/8/snapshot/views/3/nodes/3/primary_id |
| e002100 | 19:55:55.27818 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/3/slots |
| e002101 | 19:55:55.27818 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/3/link_state |
| e002102 | 19:55:55.27818 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/3/nodes/4/flags |
| e002103 | 19:55:55.27818 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/4/primary_id |
| e002104 | 19:55:55.27818 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/3/nodes/4/slots |
| e002105 | 19:55:55.27818 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/4/link_state |
| e002106 | 19:55:55.27818 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/8/snapshot/views/3/nodes/5/flags |
| e002107 | 19:55:55.27818 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/5/primary_id |
| e002108 | 19:55:55.27818 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/5/slots |
| e002109 | 19:55:55.27818 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/5/link_state |
| e002110 | 19:55:55.27818 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/3/persistence/aof_enabled |
| e002111 | 19:55:55.27818 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/3/persistence/loading |
| e002112 | 19:55:55.27818 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/3/persistence/aof_last_write_status |
| e002113 | 19:55:55.27818 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e002114 | 19:55:55.27818 | node-4 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/3/process_live |
| e002115 | 19:55:55.27818 | node-4 | node_client | reachable | valid | `true` | /final/views/3/reachable |
| e002116 | 19:55:55.27818 | node-4 | info_replication | replication.role | valid | `"slave"` | /final/views/3/replication/role |
| e002117 | 19:55:55.27818 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /final/views/3/replication/connected_slaves |
| e002118 | 19:55:55.27818 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /final/views/3/replication/master_link_status |
| e002119 | 19:55:55.27818 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /final/views/3/replication/master_sync_in_progress |
| e002120 | 19:55:55.27818 | node-4 | info_replication | replication.master_repl_offset | valid | `36227` | /final/views/3/replication/master_repl_offset |
| e002121 | 19:55:55.27818 | node-4 | info_replication | replication.slave_repl_offset | valid | `36227` | /final/views/3/replication/slave_repl_offset |
| e002122 | 19:55:55.27818 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /final/views/3/redis_id |
| e002123 | 19:55:55.27818 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /final/views/3/coverage |
| e002124 | 19:55:55.27818 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /final/views/3/nodes/0/flags |
| e002125 | 19:55:55.27818 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /final/views/3/nodes/0/primary_id |
| e002126 | 19:55:55.27818 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/3/nodes/0/slots |
| e002127 | 19:55:55.27818 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/0/link_state |
| e002128 | 19:55:55.27818 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /final/views/3/nodes/1/flags |
| e002129 | 19:55:55.27818 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /final/views/3/nodes/1/primary_id |
| e002130 | 19:55:55.27818 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /final/views/3/nodes/1/slots |
| e002131 | 19:55:55.27818 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/1/link_state |
| e002132 | 19:55:55.27818 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /final/views/3/nodes/2/flags |
| e002133 | 19:55:55.27818 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /final/views/3/nodes/2/primary_id |
| e002134 | 19:55:55.27818 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/3/nodes/2/slots |
| e002135 | 19:55:55.27818 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/2/link_state |
| e002136 | 19:55:55.27818 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/3/nodes/3/flags |
| e002137 | 19:55:55.27818 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /final/views/3/nodes/3/primary_id |
| e002138 | 19:55:55.27818 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /final/views/3/nodes/3/slots |
| e002139 | 19:55:55.27818 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/3/link_state |
| e002140 | 19:55:55.27818 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /final/views/3/nodes/4/flags |
| e002141 | 19:55:55.27818 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /final/views/3/nodes/4/primary_id |
| e002142 | 19:55:55.27818 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/3/nodes/4/slots |
| e002143 | 19:55:55.27818 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/4/link_state |
| e002144 | 19:55:55.27818 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
              "master",
              "fail"
            ]` | /final/views/3/nodes/5/flags |
| e002145 | 19:55:55.27818 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /final/views/3/nodes/5/primary_id |
| e002146 | 19:55:55.27818 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /final/views/3/nodes/5/slots |
| e002147 | 19:55:55.27818 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/5/link_state |
| e002148 | 19:55:55.27818 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/3/persistence/aof_enabled |
| e002149 | 19:55:55.27818 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /final/views/3/persistence/loading |
| e002150 | 19:55:55.27818 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/3/persistence/aof_last_write_status |
| e002151 | 19:55:55.27818 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/3/persistence/aof_rewrite_in_progress |
| e002152 | 19:55:55.27818 | node-4 | admin | process_live | missing_from_schema | `` | /final/views/3/process_live |
| e002153 | 19:55:55.278236 | node-6 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/5/reachable |
| e002154 | 19:55:55.278236 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/8/snapshot/views/5/replication/role |
| e002155 | 19:55:55.278236 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/8/snapshot/views/5/replication/connected_slaves |
| e002156 | 19:55:55.278236 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/8/snapshot/views/5/replication/master_link_status |
| e002157 | 19:55:55.278236 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/8/snapshot/views/5/replication/master_sync_in_progress |
| e002158 | 19:55:55.278236 | node-6 | info_replication | replication.master_repl_offset | valid | `36227` | /samples/8/snapshot/views/5/replication/master_repl_offset |
| e002159 | 19:55:55.278236 | node-6 | info_replication | replication.slave_repl_offset | valid | `36227` | /samples/8/snapshot/views/5/replication/slave_repl_offset |
| e002160 | 19:55:55.278236 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /samples/8/snapshot/views/5/redis_id |
| e002161 | 19:55:55.278236 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/5/coverage |
| e002162 | 19:55:55.278236 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/5/nodes/0/flags |
| e002163 | 19:55:55.278236 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/8/snapshot/views/5/nodes/0/primary_id |
| e002164 | 19:55:55.278236 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/0/slots |
| e002165 | 19:55:55.278236 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/0/link_state |
| e002166 | 19:55:55.278236 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/5/nodes/1/flags |
| e002167 | 19:55:55.278236 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/1/primary_id |
| e002168 | 19:55:55.278236 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/5/nodes/1/slots |
| e002169 | 19:55:55.278236 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/1/link_state |
| e002170 | 19:55:55.278236 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/8/snapshot/views/5/nodes/2/flags |
| e002171 | 19:55:55.278236 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/2/primary_id |
| e002172 | 19:55:55.278236 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/2/slots |
| e002173 | 19:55:55.278236 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/2/link_state |
| e002174 | 19:55:55.278236 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/5/nodes/3/flags |
| e002175 | 19:55:55.278236 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/3/primary_id |
| e002176 | 19:55:55.278236 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/5/nodes/3/slots |
| e002177 | 19:55:55.278236 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/3/link_state |
| e002178 | 19:55:55.278236 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/8/snapshot/views/5/nodes/4/flags |
| e002179 | 19:55:55.278236 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/8/snapshot/views/5/nodes/4/primary_id |
| e002180 | 19:55:55.278236 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/4/slots |
| e002181 | 19:55:55.278236 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/4/link_state |
| e002182 | 19:55:55.278236 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/5/nodes/5/flags |
| e002183 | 19:55:55.278236 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/5/primary_id |
| e002184 | 19:55:55.278236 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/5/nodes/5/slots |
| e002185 | 19:55:55.278236 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/5/link_state |
| e002186 | 19:55:55.278236 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/5/persistence/aof_enabled |
| e002187 | 19:55:55.278236 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/5/persistence/loading |
| e002188 | 19:55:55.278236 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/5/persistence/aof_last_write_status |
| e002189 | 19:55:55.278236 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002190 | 19:55:55.278236 | node-6 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/5/process_live |
| e002191 | 19:55:55.278236 | node-6 | node_client | reachable | valid | `true` | /final/views/5/reachable |
| e002192 | 19:55:55.278236 | node-6 | info_replication | replication.role | valid | `"slave"` | /final/views/5/replication/role |
| e002193 | 19:55:55.278236 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /final/views/5/replication/connected_slaves |
| e002194 | 19:55:55.278236 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /final/views/5/replication/master_link_status |
| e002195 | 19:55:55.278236 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /final/views/5/replication/master_sync_in_progress |
| e002196 | 19:55:55.278236 | node-6 | info_replication | replication.master_repl_offset | valid | `36227` | /final/views/5/replication/master_repl_offset |
| e002197 | 19:55:55.278236 | node-6 | info_replication | replication.slave_repl_offset | valid | `36227` | /final/views/5/replication/slave_repl_offset |
| e002198 | 19:55:55.278236 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /final/views/5/redis_id |
| e002199 | 19:55:55.278236 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /final/views/5/coverage |
| e002200 | 19:55:55.278236 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/5/nodes/0/flags |
| e002201 | 19:55:55.278236 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /final/views/5/nodes/0/primary_id |
| e002202 | 19:55:55.278236 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /final/views/5/nodes/0/slots |
| e002203 | 19:55:55.278236 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/0/link_state |
| e002204 | 19:55:55.278236 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /final/views/5/nodes/1/flags |
| e002205 | 19:55:55.278236 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /final/views/5/nodes/1/primary_id |
| e002206 | 19:55:55.278236 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/5/nodes/1/slots |
| e002207 | 19:55:55.278236 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/1/link_state |
| e002208 | 19:55:55.278236 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
              "master",
              "fail"
            ]` | /final/views/5/nodes/2/flags |
| e002209 | 19:55:55.278236 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /final/views/5/nodes/2/primary_id |
| e002210 | 19:55:55.278236 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /final/views/5/nodes/2/slots |
| e002211 | 19:55:55.278236 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/2/link_state |
| e002212 | 19:55:55.278236 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /final/views/5/nodes/3/flags |
| e002213 | 19:55:55.278236 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /final/views/5/nodes/3/primary_id |
| e002214 | 19:55:55.278236 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/5/nodes/3/slots |
| e002215 | 19:55:55.278236 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/3/link_state |
| e002216 | 19:55:55.278236 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /final/views/5/nodes/4/flags |
| e002217 | 19:55:55.278236 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /final/views/5/nodes/4/primary_id |
| e002218 | 19:55:55.278236 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /final/views/5/nodes/4/slots |
| e002219 | 19:55:55.278236 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/4/link_state |
| e002220 | 19:55:55.278236 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /final/views/5/nodes/5/flags |
| e002221 | 19:55:55.278236 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /final/views/5/nodes/5/primary_id |
| e002222 | 19:55:55.278236 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/5/nodes/5/slots |
| e002223 | 19:55:55.278236 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/5/link_state |
| e002224 | 19:55:55.278236 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/5/persistence/aof_enabled |
| e002225 | 19:55:55.278236 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /final/views/5/persistence/loading |
| e002226 | 19:55:55.278236 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/5/persistence/aof_last_write_status |
| e002227 | 19:55:55.278236 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/5/persistence/aof_rewrite_in_progress |
| e002228 | 19:55:55.278236 | node-6 | admin | process_live | missing_from_schema | `` | /final/views/5/process_live |
| e002229 | 19:55:55.278262 | node-2 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/1/reachable |
| e002230 | 19:55:55.278262 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/8/snapshot/views/1/replication/role |
| e002231 | 19:55:55.278262 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/8/snapshot/views/1/replication/connected_slaves |
| e002232 | 19:55:55.278262 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/8/snapshot/views/1/replication/master_link_status |
| e002233 | 19:55:55.278262 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/8/snapshot/views/1/replication/master_sync_in_progress |
| e002234 | 19:55:55.278262 | node-2 | info_replication | replication.master_repl_offset | valid | `36227` | /samples/8/snapshot/views/1/replication/master_repl_offset |
| e002235 | 19:55:55.278262 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/8/snapshot/views/1/replication/slave_repl_offset |
| e002236 | 19:55:55.278262 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/8/snapshot/views/1/redis_id |
| e002237 | 19:55:55.278262 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/1/coverage |
| e002238 | 19:55:55.278262 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/8/snapshot/views/1/nodes/0/flags |
| e002239 | 19:55:55.278262 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/0/primary_id |
| e002240 | 19:55:55.278262 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/0/slots |
| e002241 | 19:55:55.278262 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/0/link_state |
| e002242 | 19:55:55.278262 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/1/nodes/1/flags |
| e002243 | 19:55:55.278262 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/8/snapshot/views/1/nodes/1/primary_id |
| e002244 | 19:55:55.278262 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/1/slots |
| e002245 | 19:55:55.278262 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/1/link_state |
| e002246 | 19:55:55.278262 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/1/nodes/2/flags |
| e002247 | 19:55:55.278262 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/2/primary_id |
| e002248 | 19:55:55.278262 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/1/nodes/2/slots |
| e002249 | 19:55:55.278262 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/2/link_state |
| e002250 | 19:55:55.278262 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/8/snapshot/views/1/nodes/3/flags |
| e002251 | 19:55:55.278262 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/3/primary_id |
| e002252 | 19:55:55.278262 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/1/nodes/3/slots |
| e002253 | 19:55:55.278262 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/3/link_state |
| e002254 | 19:55:55.278262 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/1/nodes/4/flags |
| e002255 | 19:55:55.278262 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/4/primary_id |
| e002256 | 19:55:55.278262 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/1/nodes/4/slots |
| e002257 | 19:55:55.278262 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/4/link_state |
| e002258 | 19:55:55.278262 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/1/nodes/5/flags |
| e002259 | 19:55:55.278262 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/8/snapshot/views/1/nodes/5/primary_id |
| e002260 | 19:55:55.278262 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/5/slots |
| e002261 | 19:55:55.278262 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/5/link_state |
| e002262 | 19:55:55.278262 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/1/persistence/aof_enabled |
| e002263 | 19:55:55.278262 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/1/persistence/loading |
| e002264 | 19:55:55.278262 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/1/persistence/aof_last_write_status |
| e002265 | 19:55:55.278262 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e002266 | 19:55:55.278262 | node-2 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/1/process_live |
| e002267 | 19:55:55.278262 | node-2 | node_client | reachable | valid | `true` | /final/views/1/reachable |
| e002268 | 19:55:55.278262 | node-2 | info_replication | replication.role | valid | `"master"` | /final/views/1/replication/role |
| e002269 | 19:55:55.278262 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /final/views/1/replication/connected_slaves |
| e002270 | 19:55:55.278262 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /final/views/1/replication/master_link_status |
| e002271 | 19:55:55.278262 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /final/views/1/replication/master_sync_in_progress |
| e002272 | 19:55:55.278262 | node-2 | info_replication | replication.master_repl_offset | valid | `36227` | /final/views/1/replication/master_repl_offset |
| e002273 | 19:55:55.278262 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /final/views/1/replication/slave_repl_offset |
| e002274 | 19:55:55.278262 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /final/views/1/redis_id |
| e002275 | 19:55:55.278262 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /final/views/1/coverage |
| e002276 | 19:55:55.278262 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
              "master",
              "fail"
            ]` | /final/views/1/nodes/0/flags |
| e002277 | 19:55:55.278262 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /final/views/1/nodes/0/primary_id |
| e002278 | 19:55:55.278262 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /final/views/1/nodes/0/slots |
| e002279 | 19:55:55.278262 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/0/link_state |
| e002280 | 19:55:55.278262 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/1/nodes/1/flags |
| e002281 | 19:55:55.278262 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /final/views/1/nodes/1/primary_id |
| e002282 | 19:55:55.278262 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /final/views/1/nodes/1/slots |
| e002283 | 19:55:55.278262 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/1/link_state |
| e002284 | 19:55:55.278262 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /final/views/1/nodes/2/flags |
| e002285 | 19:55:55.278262 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /final/views/1/nodes/2/primary_id |
| e002286 | 19:55:55.278262 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/1/nodes/2/slots |
| e002287 | 19:55:55.278262 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/2/link_state |
| e002288 | 19:55:55.278262 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /final/views/1/nodes/3/flags |
| e002289 | 19:55:55.278262 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /final/views/1/nodes/3/primary_id |
| e002290 | 19:55:55.278262 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/1/nodes/3/slots |
| e002291 | 19:55:55.278262 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/3/link_state |
| e002292 | 19:55:55.278262 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /final/views/1/nodes/4/flags |
| e002293 | 19:55:55.278262 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /final/views/1/nodes/4/primary_id |
| e002294 | 19:55:55.278262 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/1/nodes/4/slots |
| e002295 | 19:55:55.278262 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/4/link_state |
| e002296 | 19:55:55.278262 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/1/nodes/5/flags |
| e002297 | 19:55:55.278262 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /final/views/1/nodes/5/primary_id |
| e002298 | 19:55:55.278262 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /final/views/1/nodes/5/slots |
| e002299 | 19:55:55.278262 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/5/link_state |
| e002300 | 19:55:55.278262 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/1/persistence/aof_enabled |
| e002301 | 19:55:55.278262 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /final/views/1/persistence/loading |
| e002302 | 19:55:55.278262 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/1/persistence/aof_last_write_status |
| e002303 | 19:55:55.278262 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/1/persistence/aof_rewrite_in_progress |
| e002304 | 19:55:55.278262 | node-2 | admin | process_live | missing_from_schema | `` | /final/views/1/process_live |
| e002305 | 19:55:55.431086 |  | client_probe | client_success | valid | `true` | /samples/8/probe/successful |
| e002306 | 19:55:55.741263 | node-2 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/1/reachable |
| e002307 | 19:55:55.741263 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/9/snapshot/views/1/replication/role |
| e002308 | 19:55:55.741263 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/9/snapshot/views/1/replication/connected_slaves |
| e002309 | 19:55:55.741263 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/9/snapshot/views/1/replication/master_link_status |
| e002310 | 19:55:55.741263 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/9/snapshot/views/1/replication/master_sync_in_progress |
| e002311 | 19:55:55.741263 | node-2 | info_replication | replication.master_repl_offset | valid | `39425` | /samples/9/snapshot/views/1/replication/master_repl_offset |
| e002312 | 19:55:55.741263 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/9/snapshot/views/1/replication/slave_repl_offset |
| e002313 | 19:55:55.741263 | node-2 | cluster_nodes | redis_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/9/snapshot/views/1/redis_id |
| e002314 | 19:55:55.741263 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/1/coverage |
| e002315 | 19:55:55.741263 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/9/snapshot/views/1/nodes/0/flags |
| e002316 | 19:55:55.741263 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/0/primary_id |
| e002317 | 19:55:55.741263 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/0/slots |
| e002318 | 19:55:55.741263 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/0/link_state |
| e002319 | 19:55:55.741263 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/1/nodes/1/flags |
| e002320 | 19:55:55.741263 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/9/snapshot/views/1/nodes/1/primary_id |
| e002321 | 19:55:55.741263 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/1/slots |
| e002322 | 19:55:55.741263 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/1/link_state |
| e002323 | 19:55:55.741263 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/1/nodes/2/flags |
| e002324 | 19:55:55.741263 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/2/primary_id |
| e002325 | 19:55:55.741263 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/1/nodes/2/slots |
| e002326 | 19:55:55.741263 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/2/link_state |
| e002327 | 19:55:55.741263 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/9/snapshot/views/1/nodes/3/flags |
| e002328 | 19:55:55.741263 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/3/primary_id |
| e002329 | 19:55:55.741263 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/1/nodes/3/slots |
| e002330 | 19:55:55.741263 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/3/link_state |
| e002331 | 19:55:55.741263 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/1/nodes/4/flags |
| e002332 | 19:55:55.741263 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/4/primary_id |
| e002333 | 19:55:55.741263 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/1/nodes/4/slots |
| e002334 | 19:55:55.741263 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/4/link_state |
| e002335 | 19:55:55.741263 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/1/nodes/5/flags |
| e002336 | 19:55:55.741263 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/9/snapshot/views/1/nodes/5/primary_id |
| e002337 | 19:55:55.741263 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/5/slots |
| e002338 | 19:55:55.741263 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/5/link_state |
| e002339 | 19:55:55.741263 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/1/persistence/aof_enabled |
| e002340 | 19:55:55.741263 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/1/persistence/loading |
| e002341 | 19:55:55.741263 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/1/persistence/aof_last_write_status |
| e002342 | 19:55:55.741263 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e002343 | 19:55:55.741263 | node-2 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/1/process_live |
| e002344 | 19:55:55.741265 | node-4 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/3/reachable |
| e002345 | 19:55:55.741265 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/9/snapshot/views/3/replication/role |
| e002346 | 19:55:55.741265 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/9/snapshot/views/3/replication/connected_slaves |
| e002347 | 19:55:55.741265 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/9/snapshot/views/3/replication/master_link_status |
| e002348 | 19:55:55.741265 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/9/snapshot/views/3/replication/master_sync_in_progress |
| e002349 | 19:55:55.741265 | node-4 | info_replication | replication.master_repl_offset | valid | `39203` | /samples/9/snapshot/views/3/replication/master_repl_offset |
| e002350 | 19:55:55.741265 | node-4 | info_replication | replication.slave_repl_offset | valid | `39203` | /samples/9/snapshot/views/3/replication/slave_repl_offset |
| e002351 | 19:55:55.741265 | node-4 | cluster_nodes | redis_id | valid | `"6c0fa4cacebf158b03ae69957d7caf1eea7b78f8"` | /samples/9/snapshot/views/3/redis_id |
| e002352 | 19:55:55.741265 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/3/coverage |
| e002353 | 19:55:55.741265 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/3/nodes/0/flags |
| e002354 | 19:55:55.741265 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/0/primary_id |
| e002355 | 19:55:55.741265 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/3/nodes/0/slots |
| e002356 | 19:55:55.741265 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/0/link_state |
| e002357 | 19:55:55.741265 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/9/snapshot/views/3/nodes/1/flags |
| e002358 | 19:55:55.741265 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/9/snapshot/views/3/nodes/1/primary_id |
| e002359 | 19:55:55.741265 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/1/slots |
| e002360 | 19:55:55.741265 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/1/link_state |
| e002361 | 19:55:55.741265 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/3/nodes/2/flags |
| e002362 | 19:55:55.741265 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/2/primary_id |
| e002363 | 19:55:55.741265 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/3/nodes/2/slots |
| e002364 | 19:55:55.741265 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/2/link_state |
| e002365 | 19:55:55.741265 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/3/nodes/3/flags |
| e002366 | 19:55:55.741265 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/9/snapshot/views/3/nodes/3/primary_id |
| e002367 | 19:55:55.741265 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/3/slots |
| e002368 | 19:55:55.741265 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/3/link_state |
| e002369 | 19:55:55.741265 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/3/nodes/4/flags |
| e002370 | 19:55:55.741265 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/4/primary_id |
| e002371 | 19:55:55.741265 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/3/nodes/4/slots |
| e002372 | 19:55:55.741265 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/4/link_state |
| e002373 | 19:55:55.741265 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/9/snapshot/views/3/nodes/5/flags |
| e002374 | 19:55:55.741265 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/5/primary_id |
| e002375 | 19:55:55.741265 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/5/slots |
| e002376 | 19:55:55.741265 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/5/link_state |
| e002377 | 19:55:55.741265 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/3/persistence/aof_enabled |
| e002378 | 19:55:55.741265 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/3/persistence/loading |
| e002379 | 19:55:55.741265 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/3/persistence/aof_last_write_status |
| e002380 | 19:55:55.741265 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e002381 | 19:55:55.741265 | node-4 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/3/process_live |
| e002382 | 19:55:55.741267 | node-1 | node_client | reachable | valid | `false` | /samples/9/snapshot/views/0/reachable |
| e002383 | 19:55:55.741267 | node-1 | info_replication | replication.role | collection_failed | `` | /samples/9/snapshot/views/0/replication/role |
| e002384 | 19:55:55.741267 | node-1 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/9/snapshot/views/0/replication/connected_slaves |
| e002385 | 19:55:55.741267 | node-1 | info_replication | replication.master_link_status | collection_failed | `` | /samples/9/snapshot/views/0/replication/master_link_status |
| e002386 | 19:55:55.741267 | node-1 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/9/snapshot/views/0/replication/master_sync_in_progress |
| e002387 | 19:55:55.741267 | node-1 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/9/snapshot/views/0/replication/master_repl_offset |
| e002388 | 19:55:55.741267 | node-1 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/9/snapshot/views/0/replication/slave_repl_offset |
| e002389 | 19:55:55.741267 | node-1 | cluster_nodes | redis_id | collection_failed | `` | /samples/9/snapshot/views/0/redis_id |
| e002390 | 19:55:55.741267 | node-1 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/9/snapshot/views/0/coverage |
| e002391 | 19:55:55.741267 | node-1 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/9/snapshot/views/0/persistence/aof_enabled |
| e002392 | 19:55:55.741267 | node-1 | info_persistence | persistence.loading | collection_failed | `` | /samples/9/snapshot/views/0/persistence/loading |
| e002393 | 19:55:55.741267 | node-1 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/9/snapshot/views/0/persistence/aof_last_write_status |
| e002394 | 19:55:55.741267 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/9/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e002395 | 19:55:55.741267 | node-1 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/0/process_live |
| e002396 | 19:55:55.741278 | node-3 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/2/reachable |
| e002397 | 19:55:55.741278 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/9/snapshot/views/2/replication/role |
| e002398 | 19:55:55.741278 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/9/snapshot/views/2/replication/connected_slaves |
| e002399 | 19:55:55.741278 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/9/snapshot/views/2/replication/master_link_status |
| e002400 | 19:55:55.741278 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/9/snapshot/views/2/replication/master_sync_in_progress |
| e002401 | 19:55:55.741278 | node-3 | info_replication | replication.master_repl_offset | valid | `39203` | /samples/9/snapshot/views/2/replication/master_repl_offset |
| e002402 | 19:55:55.741278 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/9/snapshot/views/2/replication/slave_repl_offset |
| e002403 | 19:55:55.741278 | node-3 | cluster_nodes | redis_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/9/snapshot/views/2/redis_id |
| e002404 | 19:55:55.741278 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/2/coverage |
| e002405 | 19:55:55.741278 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/9/snapshot/views/2/nodes/0/flags |
| e002406 | 19:55:55.741278 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/0/primary_id |
| e002407 | 19:55:55.741278 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/2/nodes/0/slots |
| e002408 | 19:55:55.741278 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/0/link_state |
| e002409 | 19:55:55.741278 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/9/snapshot/views/2/nodes/1/flags |
| e002410 | 19:55:55.741278 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/1/primary_id |
| e002411 | 19:55:55.741278 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/1/slots |
| e002412 | 19:55:55.741278 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/1/link_state |
| e002413 | 19:55:55.741278 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/2/nodes/2/flags |
| e002414 | 19:55:55.741278 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/9/snapshot/views/2/nodes/2/primary_id |
| e002415 | 19:55:55.741278 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/2/slots |
| e002416 | 19:55:55.741278 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/2/link_state |
| e002417 | 19:55:55.741278 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/2/nodes/3/flags |
| e002418 | 19:55:55.741278 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/3/primary_id |
| e002419 | 19:55:55.741278 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/2/nodes/3/slots |
| e002420 | 19:55:55.741278 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/3/link_state |
| e002421 | 19:55:55.741278 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/2/nodes/4/flags |
| e002422 | 19:55:55.741278 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/4/primary_id |
| e002423 | 19:55:55.741278 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/2/nodes/4/slots |
| e002424 | 19:55:55.741278 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/4/link_state |
| e002425 | 19:55:55.741278 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/2/nodes/5/flags |
| e002426 | 19:55:55.741278 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/9/snapshot/views/2/nodes/5/primary_id |
| e002427 | 19:55:55.741278 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/5/slots |
| e002428 | 19:55:55.741278 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/5/link_state |
| e002429 | 19:55:55.741278 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/2/persistence/aof_enabled |
| e002430 | 19:55:55.741278 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/2/persistence/loading |
| e002431 | 19:55:55.741278 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/2/persistence/aof_last_write_status |
| e002432 | 19:55:55.741278 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e002433 | 19:55:55.741278 | node-3 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/2/process_live |
| e002434 | 19:55:55.741284 | node-5 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/4/reachable |
| e002435 | 19:55:55.741284 | node-5 | info_replication | replication.role | valid | `"master"` | /samples/9/snapshot/views/4/replication/role |
| e002436 | 19:55:55.741284 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/9/snapshot/views/4/replication/connected_slaves |
| e002437 | 19:55:55.741284 | node-5 | info_replication | replication.master_link_status | not_applicable | `` | /samples/9/snapshot/views/4/replication/master_link_status |
| e002438 | 19:55:55.741284 | node-5 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/9/snapshot/views/4/replication/master_sync_in_progress |
| e002439 | 19:55:55.741284 | node-5 | info_replication | replication.master_repl_offset | valid | `32072` | /samples/9/snapshot/views/4/replication/master_repl_offset |
| e002440 | 19:55:55.741284 | node-5 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/9/snapshot/views/4/replication/slave_repl_offset |
| e002441 | 19:55:55.741284 | node-5 | cluster_nodes | redis_id | valid | `"f66e6677e1becbbc01c8361d997d3a68409cb5e8"` | /samples/9/snapshot/views/4/redis_id |
| e002442 | 19:55:55.741284 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/4/coverage |
| e002443 | 19:55:55.741284 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/4/nodes/0/flags |
| e002444 | 19:55:55.741284 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/9/snapshot/views/4/nodes/0/primary_id |
| e002445 | 19:55:55.741284 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/0/slots |
| e002446 | 19:55:55.741284 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/0/link_state |
| e002447 | 19:55:55.741284 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/4/nodes/1/flags |
| e002448 | 19:55:55.741284 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/1/primary_id |
| e002449 | 19:55:55.741284 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/4/nodes/1/slots |
| e002450 | 19:55:55.741284 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/1/link_state |
| e002451 | 19:55:55.741284 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/9/snapshot/views/4/nodes/2/flags |
| e002452 | 19:55:55.741284 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/2/primary_id |
| e002453 | 19:55:55.741284 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/2/slots |
| e002454 | 19:55:55.741284 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/2/link_state |
| e002455 | 19:55:55.741284 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/9/snapshot/views/4/nodes/3/flags |
| e002456 | 19:55:55.741284 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/3/primary_id |
| e002457 | 19:55:55.741284 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/4/nodes/3/slots |
| e002458 | 19:55:55.741284 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/3/link_state |
| e002459 | 19:55:55.741284 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/4/nodes/4/flags |
| e002460 | 19:55:55.741284 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/9/snapshot/views/4/nodes/4/primary_id |
| e002461 | 19:55:55.741284 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/4/slots |
| e002462 | 19:55:55.741284 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/4/link_state |
| e002463 | 19:55:55.741284 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/4/nodes/5/flags |
| e002464 | 19:55:55.741284 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/5/primary_id |
| e002465 | 19:55:55.741284 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/4/nodes/5/slots |
| e002466 | 19:55:55.741284 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/5/link_state |
| e002467 | 19:55:55.741284 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/4/persistence/aof_enabled |
| e002468 | 19:55:55.741284 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/4/persistence/loading |
| e002469 | 19:55:55.741284 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/4/persistence/aof_last_write_status |
| e002470 | 19:55:55.741284 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e002471 | 19:55:55.741284 | node-5 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/4/process_live |
| e002472 | 19:55:55.741305 | node-6 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/5/reachable |
| e002473 | 19:55:55.741305 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/9/snapshot/views/5/replication/role |
| e002474 | 19:55:55.741305 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/9/snapshot/views/5/replication/connected_slaves |
| e002475 | 19:55:55.741305 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/9/snapshot/views/5/replication/master_link_status |
| e002476 | 19:55:55.741305 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/9/snapshot/views/5/replication/master_sync_in_progress |
| e002477 | 19:55:55.741305 | node-6 | info_replication | replication.master_repl_offset | valid | `39425` | /samples/9/snapshot/views/5/replication/master_repl_offset |
| e002478 | 19:55:55.741305 | node-6 | info_replication | replication.slave_repl_offset | valid | `39425` | /samples/9/snapshot/views/5/replication/slave_repl_offset |
| e002479 | 19:55:55.741305 | node-6 | cluster_nodes | redis_id | valid | `"aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8"` | /samples/9/snapshot/views/5/redis_id |
| e002480 | 19:55:55.741305 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/5/coverage |
| e002481 | 19:55:55.741305 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/5/nodes/0/flags |
| e002482 | 19:55:55.741305 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.primary_id | valid | `"e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be"` | /samples/9/snapshot/views/5/nodes/0/primary_id |
| e002483 | 19:55:55.741305 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/0/slots |
| e002484 | 19:55:55.741305 | 6c0fa4cacebf158b03ae69957d7caf1eea7b78f8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/0/link_state |
| e002485 | 19:55:55.741305 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/5/nodes/1/flags |
| e002486 | 19:55:55.741305 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/1/primary_id |
| e002487 | 19:55:55.741305 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/5/nodes/1/slots |
| e002488 | 19:55:55.741305 | 70b92c82b60b348aa5fa8b9ab8682308cd798517 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/1/link_state |
| e002489 | 19:55:55.741305 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.flags | valid | `[
                  "master",
                  "fail"
                ]` | /samples/9/snapshot/views/5/nodes/2/flags |
| e002490 | 19:55:55.741305 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/2/primary_id |
| e002491 | 19:55:55.741305 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/2/slots |
| e002492 | 19:55:55.741305 | 146512ed263b70c95686db0f88d848a4b2c631cf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/2/link_state |
| e002493 | 19:55:55.741305 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/5/nodes/3/flags |
| e002494 | 19:55:55.741305 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/3/primary_id |
| e002495 | 19:55:55.741305 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/5/nodes/3/slots |
| e002496 | 19:55:55.741305 | e90ead82ed15fd65666b6a3aa6fb4c8a297dd7be | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/3/link_state |
| e002497 | 19:55:55.741305 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/9/snapshot/views/5/nodes/4/flags |
| e002498 | 19:55:55.741305 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.primary_id | valid | `"70b92c82b60b348aa5fa8b9ab8682308cd798517"` | /samples/9/snapshot/views/5/nodes/4/primary_id |
| e002499 | 19:55:55.741305 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/4/slots |
| e002500 | 19:55:55.741305 | aec27de7d78f4bc2e55b33fcae8ad8f2048d1df8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/4/link_state |
| e002501 | 19:55:55.741305 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/5/nodes/5/flags |
| e002502 | 19:55:55.741305 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/5/primary_id |
| e002503 | 19:55:55.741305 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/5/nodes/5/slots |
| e002504 | 19:55:55.741305 | f66e6677e1becbbc01c8361d997d3a68409cb5e8 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/5/link_state |
| e002505 | 19:55:55.741305 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/5/persistence/aof_enabled |
| e002506 | 19:55:55.741305 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/5/persistence/loading |
| e002507 | 19:55:55.741305 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/5/persistence/aof_last_write_status |
| e002508 | 19:55:55.741305 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002509 | 19:55:55.741305 | node-6 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/5/process_live |
| e002510 | 19:55:55.880843 |  | client_probe | client_success | valid | `false` | /samples/9/probe/successful |
| e002511 | 19:55:55.880926 |  | client_window | client.attempts | valid | `112` | /metrics/windows/7/attempts |
| e002512 | 19:55:55.880926 |  | client_window | client.errors | valid | `0` | /metrics/windows/7/errors |
| e002513 | 19:55:55.880926 |  | client_window | client.successes | valid | `112` | /metrics/windows/7/successes |
| e002514 | 19:55:55.880926 |  | client_window | client.p95_ms | valid | `1.19212` | /metrics/windows/7/p95_ms |
| e002515 | 19:55:56.080706 |  | value_validation | validation.expected | valid | `504` | /validation/expected |
| e002516 | 19:55:56.080706 |  | value_validation | validation.checked | valid | `504` | /validation/checked |
| e002517 | 19:55:56.080706 |  | value_validation | validation.missing | valid | `0` | /validation/missing |
| e002518 | 19:55:56.080706 |  | value_validation | validation.corrupt | valid | `0` | /validation/corrupt |
| e002519 | 19:55:56.080706 |  | value_validation | validation.unreadable | valid | `0` | /validation/unreadable |
| e002520 | 19:55:56.080706 |  | value_validation | validation.unverified | valid | `0` | /validation/unverified |
| e002521 | 19:55:56.080706 |  | value_validation | validation.complete | valid | `true` | /validation/complete |

## Analysis limitations

- Offline laboratory observation analysis; no production integration.
- Node views are non-atomic and retained independently; sample timestamps are not exact server event times.
- No fault labels, injected action results, run identifiers, configuration hashes or phase names are inference inputs.
- Client reachability does not establish process death; no administrative process telemetry is accepted from fault-specific fields.
- Validation proves only the recorded checked population, never general durability or power-loss safety.
