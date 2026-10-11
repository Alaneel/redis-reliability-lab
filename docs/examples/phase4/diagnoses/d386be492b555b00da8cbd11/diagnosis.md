# Incident diagnosis

Incident: `d386be492b555b00da8cbd11`  
Detection: **incident**  
Unique root cause: **abstained**  
Engine: `rules-2`

## Observed symptoms and diagnostic hypotheses

### acknowledgement_degradation

At least one acknowledged SET received a valid short WAIT/WAITAOF reply. The SET succeeded; the additional acknowledgement policy was not satisfied.

Assessment: strongly_supported. Rule: `R8-acknowledged-set-valid-short-policy-reply`. Nodes: .

Supporting evidence: e000443, e000445, e000442, e000446, e000447, e000444, e000663, e000665, e000662, e000666, e000667, e000664, e000883, e000885, e000882, e000886, e000887, e000884, e001099, e001101, e001098, e001102, e001103, e001100, e001319, e001321, e001318, e001322, e001323, e001320, e001539, e001541, e001538, e001542, e001543, e001540, e001787, e001789, e001786, e001790, e001791, e001788, e002031, e002033, e002030, e002034, e002035, e002032, e002271, e002273, e002270, e002274, e002275, e002272, e002515, e002517, e002514, e002518, e002519, e002516, e002759, e002761, e002758, e002762, e002763, e002760.

Contrary or recovery observations: .

Missing evidence: Independent process liveness and network-path observations; Per-request routing/cache/redirect traces.

Alternatives: Process or endpoint failure; Network or observer connectivity problem; Planned administrative change.

### client_disruption

Client request window contains terminal errors; policy and SET outcomes must be inspected separately.

Assessment: strongly_supported. Rule: `R7-terminal-client-errors`. Nodes: .

Supporting evidence: e000454, e000879, e001110, e001535, e001550, e001554, e002027, e002282, e002526, e002998, e003254.

Contrary or recovery observations: e000441, e000661, e000877, e001097, e001317, e001533, e001785, e002025, e002269, e002513, e002757, e003001, e003241, e003485, e003729, e003969.

Missing evidence: Independent process liveness and network-path observations; Per-request routing/cache/redirect traces.

Alternatives: Process or endpoint failure; Network or observer connectivity problem; Planned administrative change.

### client_latency_change

Terminal request-window p95 is at least five times an earlier error-free measured window and at least 20 ms. Workload mix and acknowledgement waits can explain this; it is not a calibrated anomaly probability.

Assessment: strongly_supported. Rule: `R15-measured-window-p95-5x-and-20ms-with-20-attempts`. Nodes: .

Supporting evidence: e000456, e000004, e000453, e000881, e000878, e001112, e001109, e001537, e001534, e001552, e001549, e001556, e001553, e002029, e002026, e002284, e002281, e002528, e002525, e003000, e002997.

Contrary or recovery observations: .

Missing evidence: Independent process liveness and network-path observations; Per-request routing/cache/redirect traces.

Alternatives: Process or endpoint failure; Network or observer connectivity problem; Planned administrative change.

### redundancy_degradation

A directly observed primary reports no connected replicas; service availability and synchronization remain separate questions.

Assessment: strongly_supported. Rule: `R3-primary-zero-connected-replicas-after-positive-baseline`. Nodes: node-1.

Supporting evidence: e000253, e000125, e000625, e000841, e001061, e001115, e001497.

Contrary or recovery observations: .

Missing evidence: Independent process liveness and network-path observations; Per-request routing/cache/redirect traces.

Alternatives: Process or endpoint failure; Network or observer connectivity problem; Planned administrative change.

### replica_unavailability

Node observation endpoint unavailable; the cause remains unresolved.

Assessment: strongly_supported. Rule: `R2-endpoint-unavailable-with-prior-role`. Nodes: node-5.

Supporting evidence: e000237, e000048, e000457, e000711, e000969, e001265, e001405, e001674, e001914, e002196, e002402, e002646, e002962.

Contrary or recovery observations: e000047, e001671, e001911, e002193, e002399, e002643, e002959, e003203, e003447, e003573, e003893, e004361, e004399.

Missing evidence: Independent process liveness and network-path observations; Per-request routing/cache/redirect traces.

Alternatives: Process or endpoint failure; Network or observer connectivity problem; Planned administrative change.

### topology_disagreement

Contemporaneous peer views disagree about the same node; no single view is chosen as authoritative.

Assessment: strongly_supported. Rule: `R14-conflicting-contemporaneous-peer-views`. Nodes: 35a19f4babbca2395b87729c60bd927ff4a61972.

Supporting evidence: e000694, e000738.

Contrary or recovery observations: .

Missing evidence: Atomic topology snapshot is unavailable.

Alternatives: Normal gossip convergence; Stale peer view; Partitioned membership.

## Timeline

| First observed time | Event | Evidence | Timing |
|---|---|---|---|
| 2026-10-11T19:57:50.381842+08:00 | baseline_full_slot_coverage_observed | e000017, e000055, e000093, e000131, e000169, e000207 | all recorded views in this non-atomic sample report full healthy coverage |
| 2026-10-11T19:57:50.733451+08:00 | component_unavailability_observed | e000237, e000047 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:57:50.733451+08:00 | first_observed_replica_unavailability | e000237 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:57:50.733461+08:00 | first_observed_redundancy_degradation | e000253 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:57:51.049017+08:00 | first_observed_acknowledgement_degradation | e000443 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:57:51.381014+08:00 | first_observed_client_disruption | e000454 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:57:51.381014+08:00 | first_observed_client_latency_change | e000456 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:57:52.087733+08:00 | topology_disagreement_observed | e000694 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:57:52.087752+08:00 | cluster_failure_flag_observed | e000738, e000560 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:57:52.087774+08:00 | cluster_failure_flag_observed | e000818, e000602 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:57:52.08782+08:00 | cluster_failure_flag_observed | e000848, e000632 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:57:52.751649+08:00 | cluster_failure_flag_observed | e000922, e000792 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:57:52.751653+08:00 | cluster_failure_flag_observed | e000952, e000694 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:57:56.439839+08:00 | primary_connected_replica_count_restored | e001635, e001497 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:57:56.439851+08:00 | node_reachability_restored | e001671, e001405 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:57:56.439857+08:00 | slot_coverage_convergence_observed | e001565, e001603, e001641, e001679, e001717, e001755 | all recorded views in this non-atomic sample report full healthy coverage |
| 2026-10-11T19:58:00.4729+08:00 | acknowledgement_policy_restoration_observed | e003003, e002759 | transition sampled within observation interval; exact time unknown |
| 2026-10-11T19:58:04.108348+08:00 | final_correctness_validation_observed | e004463 | first_sampled_observation; exact event time unknown |

## Recovery and uncertainty

- Last client probe successful=true (e004441)
- Final checked acknowledged-value population preserved; see /validation evidence. This is not a durability guarantee.
- Insufficient evidence to identify a unique root cause. Process failure, network disruption and observation-channel failure can be indistinguishable.

## Read-only remediation plans

- **acknowledgement_degradation**: Inspect configured acknowledgement thresholds and replica/AOF health. Evidence: e000443, e000445, e000442, e000446, e000447, e000444, e000663, e000665, e000662, e000666, e000667, e000664, e000883, e000885, e000882, e000886, e000887, e000884, e001099, e001101, e001098, e001102, e001103, e001100, e001319, e001321, e001318, e001322, e001323, e001320, e001539, e001541, e001538, e001542, e001543, e001540, e001787, e001789, e001786, e001790, e001791, e001788, e002031, e002033, e002030, e002034, e002035, e002032, e002271, e002273, e002270, e002274, e002275, e002272, e002515, e002517, e002514, e002518, e002519, e002516, e002759, e002761, e002758, e002762, e002763, e002760. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability; Reducing acknowledgement thresholds changes durability expectations. Expected effect: Resolve alternative explanations without executing remediation. Verify: Inspect independently collected process and network status; Compare all node views and client outcomes at aligned timestamps; Observe successful SET and explicit sufficient WAIT/WAITAOF counts on the same connection. Human approval required; no execution.
- **client_disruption**: Collect independent evidence before changing system state. Evidence: e000454, e000879, e001110, e001535, e001550, e001554, e002027, e002282, e002526, e002998, e003254. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability. Expected effect: Resolve alternative explanations without executing remediation. Verify: Inspect independently collected process and network status; Compare all node views and client outcomes at aligned timestamps. Human approval required; no execution.
- **client_latency_change**: Collect independent evidence before changing system state. Evidence: e000456, e000004, e000453, e000881, e000878, e001112, e001109, e001537, e001534, e001552, e001549, e001556, e001553, e002029, e002026, e002284, e002281, e002528, e002525, e003000, e002997. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability. Expected effect: Resolve alternative explanations without executing remediation. Verify: Inspect independently collected process and network status; Compare all node views and client outcomes at aligned timestamps. Human approval required; no execution.
- **redundancy_degradation**: Assess restoration of a synchronized replica. Evidence: e000253, e000125, e000625, e000841, e001061, e001115, e001497. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups; Confirm an authoritative healthy primary and compatible former-node identity and data. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability. Expected effect: Resolve alternative explanations without executing remediation. Verify: Inspect independently collected process and network status; Compare all node views and client outcomes at aligned timestamps; Observe replica assignment, link up, synchronization complete and repeated healthy service probes. Human approval required; no execution.
- **replica_unavailability**: Assess restoration of a synchronized replica. Evidence: e000237, e000048, e000457, e000711, e000969, e001265, e001405, e001674, e001914, e002196, e002402, e002646, e002962. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups; Confirm an authoritative healthy primary and compatible former-node identity and data. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability. Expected effect: Resolve alternative explanations without executing remediation. Verify: Inspect independently collected process and network status; Compare all node views and client outcomes at aligned timestamps; Observe replica assignment, link up, synchronization complete and repeated healthy service probes. Human approval required; no execution.
- **topology_disagreement**: Collect independent evidence before changing system state. Evidence: e000694, e000738. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability. Expected effect: Resolve alternative explanations without executing remediation. Verify: Collect repeated independent views and compare per-node epochs and membership. Human approval required; no execution.

## Evidence provenance

Original bytes SHA-256: `96608a41004ade3fec33399f0e3d1ef57a42c54a40673eb8a24297d8dffc74d2`. Input report schema: 3. References are JSON pointers in the original report.

| ID | Time | Node | Source | Signal | Status | Value | Reference |
|---|---|---|---|---|---|---|---|
| e000001 | 19:57:49.381014 |  | client_window | client.attempts | valid | `169` | /metrics/windows/0/attempts |
| e000002 | 19:57:49.381014 |  | client_window | client.errors | valid | `0` | /metrics/windows/0/errors |
| e000003 | 19:57:49.381014 |  | client_window | client.successes | valid | `169` | /metrics/windows/0/successes |
| e000004 | 19:57:49.381014 |  | client_window | client.p95_ms | valid | `1.849374` | /metrics/windows/0/p95_ms |
| e000005 | 19:57:50.381014 |  | client_window | client.attempts | valid | `170` | /metrics/windows/1/attempts |
| e000006 | 19:57:50.381014 |  | client_window | client.errors | valid | `0` | /metrics/windows/1/errors |
| e000007 | 19:57:50.381014 |  | client_window | client.successes | valid | `170` | /metrics/windows/1/successes |
| e000008 | 19:57:50.381014 |  | client_window | client.p95_ms | valid | `1.19212` | /metrics/windows/1/p95_ms |
| e000009 | 19:57:50.381788 | node-2 | node_client | reachable | valid | `true` | /initial/views/1/reachable |
| e000010 | 19:57:50.381788 | node-2 | info_replication | replication.role | valid | `"master"` | /initial/views/1/replication/role |
| e000011 | 19:57:50.381788 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /initial/views/1/replication/connected_slaves |
| e000012 | 19:57:50.381788 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /initial/views/1/replication/master_link_status |
| e000013 | 19:57:50.381788 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /initial/views/1/replication/master_sync_in_progress |
| e000014 | 19:57:50.381788 | node-2 | info_replication | replication.master_repl_offset | valid | `15009` | /initial/views/1/replication/master_repl_offset |
| e000015 | 19:57:50.381788 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /initial/views/1/replication/slave_repl_offset |
| e000016 | 19:57:50.381788 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /initial/views/1/redis_id |
| e000017 | 19:57:50.381788 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/1/coverage |
| e000018 | 19:57:50.381788 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/1/nodes/0/flags |
| e000019 | 19:57:50.381788 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /initial/views/1/nodes/0/primary_id |
| e000020 | 19:57:50.381788 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/1/nodes/0/slots |
| e000021 | 19:57:50.381788 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/0/link_state |
| e000022 | 19:57:50.381788 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/1/nodes/1/flags |
| e000023 | 19:57:50.381788 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /initial/views/1/nodes/1/primary_id |
| e000024 | 19:57:50.381788 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /initial/views/1/nodes/1/slots |
| e000025 | 19:57:50.381788 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/1/link_state |
| e000026 | 19:57:50.381788 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/1/nodes/2/flags |
| e000027 | 19:57:50.381788 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /initial/views/1/nodes/2/primary_id |
| e000028 | 19:57:50.381788 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /initial/views/1/nodes/2/slots |
| e000029 | 19:57:50.381788 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/2/link_state |
| e000030 | 19:57:50.381788 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/1/nodes/3/flags |
| e000031 | 19:57:50.381788 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /initial/views/1/nodes/3/primary_id |
| e000032 | 19:57:50.381788 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /initial/views/1/nodes/3/slots |
| e000033 | 19:57:50.381788 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/3/link_state |
| e000034 | 19:57:50.381788 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /initial/views/1/nodes/4/flags |
| e000035 | 19:57:50.381788 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /initial/views/1/nodes/4/primary_id |
| e000036 | 19:57:50.381788 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/1/nodes/4/slots |
| e000037 | 19:57:50.381788 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/4/link_state |
| e000038 | 19:57:50.381788 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/1/nodes/5/flags |
| e000039 | 19:57:50.381788 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /initial/views/1/nodes/5/primary_id |
| e000040 | 19:57:50.381788 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/1/nodes/5/slots |
| e000041 | 19:57:50.381788 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/5/link_state |
| e000042 | 19:57:50.381788 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/1/persistence/aof_enabled |
| e000043 | 19:57:50.381788 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/1/persistence/loading |
| e000044 | 19:57:50.381788 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/1/persistence/aof_last_write_status |
| e000045 | 19:57:50.381788 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/1/persistence/aof_rewrite_in_progress |
| e000046 | 19:57:50.381788 | node-2 | admin | process_live | missing_from_schema | `` | /initial/views/1/process_live |
| e000047 | 19:57:50.381788 | node-5 | node_client | reachable | valid | `true` | /initial/views/4/reachable |
| e000048 | 19:57:50.381788 | node-5 | info_replication | replication.role | valid | `"slave"` | /initial/views/4/replication/role |
| e000049 | 19:57:50.381788 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /initial/views/4/replication/connected_slaves |
| e000050 | 19:57:50.381788 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /initial/views/4/replication/master_link_status |
| e000051 | 19:57:50.381788 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /initial/views/4/replication/master_sync_in_progress |
| e000052 | 19:57:50.381788 | node-5 | info_replication | replication.master_repl_offset | valid | `15007` | /initial/views/4/replication/master_repl_offset |
| e000053 | 19:57:50.381788 | node-5 | info_replication | replication.slave_repl_offset | valid | `15007` | /initial/views/4/replication/slave_repl_offset |
| e000054 | 19:57:50.381788 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /initial/views/4/redis_id |
| e000055 | 19:57:50.381788 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/4/coverage |
| e000056 | 19:57:50.381788 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/4/nodes/0/flags |
| e000057 | 19:57:50.381788 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /initial/views/4/nodes/0/primary_id |
| e000058 | 19:57:50.381788 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/4/nodes/0/slots |
| e000059 | 19:57:50.381788 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/0/link_state |
| e000060 | 19:57:50.381788 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/4/nodes/1/flags |
| e000061 | 19:57:50.381788 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /initial/views/4/nodes/1/primary_id |
| e000062 | 19:57:50.381788 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /initial/views/4/nodes/1/slots |
| e000063 | 19:57:50.381788 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/1/link_state |
| e000064 | 19:57:50.381788 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/4/nodes/2/flags |
| e000065 | 19:57:50.381788 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /initial/views/4/nodes/2/primary_id |
| e000066 | 19:57:50.381788 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /initial/views/4/nodes/2/slots |
| e000067 | 19:57:50.381788 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/2/link_state |
| e000068 | 19:57:50.381788 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/4/nodes/3/flags |
| e000069 | 19:57:50.381788 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /initial/views/4/nodes/3/primary_id |
| e000070 | 19:57:50.381788 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/4/nodes/3/slots |
| e000071 | 19:57:50.381788 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/3/link_state |
| e000072 | 19:57:50.381788 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /initial/views/4/nodes/4/flags |
| e000073 | 19:57:50.381788 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /initial/views/4/nodes/4/primary_id |
| e000074 | 19:57:50.381788 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /initial/views/4/nodes/4/slots |
| e000075 | 19:57:50.381788 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/4/link_state |
| e000076 | 19:57:50.381788 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/4/nodes/5/flags |
| e000077 | 19:57:50.381788 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /initial/views/4/nodes/5/primary_id |
| e000078 | 19:57:50.381788 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/4/nodes/5/slots |
| e000079 | 19:57:50.381788 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/5/link_state |
| e000080 | 19:57:50.381788 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/4/persistence/aof_enabled |
| e000081 | 19:57:50.381788 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/4/persistence/loading |
| e000082 | 19:57:50.381788 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/4/persistence/aof_last_write_status |
| e000083 | 19:57:50.381788 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/4/persistence/aof_rewrite_in_progress |
| e000084 | 19:57:50.381788 | node-5 | admin | process_live | missing_from_schema | `` | /initial/views/4/process_live |
| e000085 | 19:57:50.381793 | node-3 | node_client | reachable | valid | `true` | /initial/views/2/reachable |
| e000086 | 19:57:50.381793 | node-3 | info_replication | replication.role | valid | `"master"` | /initial/views/2/replication/role |
| e000087 | 19:57:50.381793 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /initial/views/2/replication/connected_slaves |
| e000088 | 19:57:50.381793 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /initial/views/2/replication/master_link_status |
| e000089 | 19:57:50.381793 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /initial/views/2/replication/master_sync_in_progress |
| e000090 | 19:57:50.381793 | node-3 | info_replication | replication.master_repl_offset | valid | `14750` | /initial/views/2/replication/master_repl_offset |
| e000091 | 19:57:50.381793 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /initial/views/2/replication/slave_repl_offset |
| e000092 | 19:57:50.381793 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /initial/views/2/redis_id |
| e000093 | 19:57:50.381793 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/2/coverage |
| e000094 | 19:57:50.381793 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/2/nodes/0/flags |
| e000095 | 19:57:50.381793 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /initial/views/2/nodes/0/primary_id |
| e000096 | 19:57:50.381793 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /initial/views/2/nodes/0/slots |
| e000097 | 19:57:50.381793 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/0/link_state |
| e000098 | 19:57:50.381793 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /initial/views/2/nodes/1/flags |
| e000099 | 19:57:50.381793 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /initial/views/2/nodes/1/primary_id |
| e000100 | 19:57:50.381793 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/2/nodes/1/slots |
| e000101 | 19:57:50.381793 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/1/link_state |
| e000102 | 19:57:50.381793 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/2/nodes/2/flags |
| e000103 | 19:57:50.381793 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /initial/views/2/nodes/2/primary_id |
| e000104 | 19:57:50.381793 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /initial/views/2/nodes/2/slots |
| e000105 | 19:57:50.381793 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/2/link_state |
| e000106 | 19:57:50.381793 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/2/nodes/3/flags |
| e000107 | 19:57:50.381793 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /initial/views/2/nodes/3/primary_id |
| e000108 | 19:57:50.381793 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/2/nodes/3/slots |
| e000109 | 19:57:50.381793 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/3/link_state |
| e000110 | 19:57:50.381793 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/2/nodes/4/flags |
| e000111 | 19:57:50.381793 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /initial/views/2/nodes/4/primary_id |
| e000112 | 19:57:50.381793 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/2/nodes/4/slots |
| e000113 | 19:57:50.381793 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/4/link_state |
| e000114 | 19:57:50.381793 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/2/nodes/5/flags |
| e000115 | 19:57:50.381793 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /initial/views/2/nodes/5/primary_id |
| e000116 | 19:57:50.381793 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /initial/views/2/nodes/5/slots |
| e000117 | 19:57:50.381793 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/5/link_state |
| e000118 | 19:57:50.381793 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/2/persistence/aof_enabled |
| e000119 | 19:57:50.381793 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/2/persistence/loading |
| e000120 | 19:57:50.381793 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/2/persistence/aof_last_write_status |
| e000121 | 19:57:50.381793 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/2/persistence/aof_rewrite_in_progress |
| e000122 | 19:57:50.381793 | node-3 | admin | process_live | missing_from_schema | `` | /initial/views/2/process_live |
| e000123 | 19:57:50.3818 | node-1 | node_client | reachable | valid | `true` | /initial/views/0/reachable |
| e000124 | 19:57:50.3818 | node-1 | info_replication | replication.role | valid | `"master"` | /initial/views/0/replication/role |
| e000125 | 19:57:50.3818 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /initial/views/0/replication/connected_slaves |
| e000126 | 19:57:50.3818 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /initial/views/0/replication/master_link_status |
| e000127 | 19:57:50.3818 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /initial/views/0/replication/master_sync_in_progress |
| e000128 | 19:57:50.3818 | node-1 | info_replication | replication.master_repl_offset | valid | `15007` | /initial/views/0/replication/master_repl_offset |
| e000129 | 19:57:50.3818 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /initial/views/0/replication/slave_repl_offset |
| e000130 | 19:57:50.3818 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /initial/views/0/redis_id |
| e000131 | 19:57:50.3818 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/0/coverage |
| e000132 | 19:57:50.3818 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/0/nodes/0/flags |
| e000133 | 19:57:50.3818 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /initial/views/0/nodes/0/primary_id |
| e000134 | 19:57:50.3818 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /initial/views/0/nodes/0/slots |
| e000135 | 19:57:50.3818 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/0/link_state |
| e000136 | 19:57:50.3818 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/0/nodes/1/flags |
| e000137 | 19:57:50.3818 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /initial/views/0/nodes/1/primary_id |
| e000138 | 19:57:50.3818 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /initial/views/0/nodes/1/slots |
| e000139 | 19:57:50.3818 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/1/link_state |
| e000140 | 19:57:50.3818 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/0/nodes/2/flags |
| e000141 | 19:57:50.3818 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /initial/views/0/nodes/2/primary_id |
| e000142 | 19:57:50.3818 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/0/nodes/2/slots |
| e000143 | 19:57:50.3818 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/2/link_state |
| e000144 | 19:57:50.3818 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /initial/views/0/nodes/3/flags |
| e000145 | 19:57:50.3818 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /initial/views/0/nodes/3/primary_id |
| e000146 | 19:57:50.3818 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/0/nodes/3/slots |
| e000147 | 19:57:50.3818 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/3/link_state |
| e000148 | 19:57:50.3818 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/0/nodes/4/flags |
| e000149 | 19:57:50.3818 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /initial/views/0/nodes/4/primary_id |
| e000150 | 19:57:50.3818 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /initial/views/0/nodes/4/slots |
| e000151 | 19:57:50.3818 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/4/link_state |
| e000152 | 19:57:50.3818 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/0/nodes/5/flags |
| e000153 | 19:57:50.3818 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /initial/views/0/nodes/5/primary_id |
| e000154 | 19:57:50.3818 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/0/nodes/5/slots |
| e000155 | 19:57:50.3818 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/5/link_state |
| e000156 | 19:57:50.3818 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/0/persistence/aof_enabled |
| e000157 | 19:57:50.3818 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/0/persistence/loading |
| e000158 | 19:57:50.3818 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/0/persistence/aof_last_write_status |
| e000159 | 19:57:50.3818 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/0/persistence/aof_rewrite_in_progress |
| e000160 | 19:57:50.3818 | node-1 | admin | process_live | missing_from_schema | `` | /initial/views/0/process_live |
| e000161 | 19:57:50.381841 | node-6 | node_client | reachable | valid | `true` | /initial/views/5/reachable |
| e000162 | 19:57:50.381841 | node-6 | info_replication | replication.role | valid | `"slave"` | /initial/views/5/replication/role |
| e000163 | 19:57:50.381841 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /initial/views/5/replication/connected_slaves |
| e000164 | 19:57:50.381841 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /initial/views/5/replication/master_link_status |
| e000165 | 19:57:50.381841 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /initial/views/5/replication/master_sync_in_progress |
| e000166 | 19:57:50.381841 | node-6 | info_replication | replication.master_repl_offset | valid | `15009` | /initial/views/5/replication/master_repl_offset |
| e000167 | 19:57:50.381841 | node-6 | info_replication | replication.slave_repl_offset | valid | `15009` | /initial/views/5/replication/slave_repl_offset |
| e000168 | 19:57:50.381841 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /initial/views/5/redis_id |
| e000169 | 19:57:50.381841 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/5/coverage |
| e000170 | 19:57:50.381841 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /initial/views/5/nodes/0/flags |
| e000171 | 19:57:50.381841 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /initial/views/5/nodes/0/primary_id |
| e000172 | 19:57:50.381841 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /initial/views/5/nodes/0/slots |
| e000173 | 19:57:50.381841 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/0/link_state |
| e000174 | 19:57:50.381841 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/5/nodes/1/flags |
| e000175 | 19:57:50.381841 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /initial/views/5/nodes/1/primary_id |
| e000176 | 19:57:50.381841 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/5/nodes/1/slots |
| e000177 | 19:57:50.381841 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/1/link_state |
| e000178 | 19:57:50.381841 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/5/nodes/2/flags |
| e000179 | 19:57:50.381841 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /initial/views/5/nodes/2/primary_id |
| e000180 | 19:57:50.381841 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /initial/views/5/nodes/2/slots |
| e000181 | 19:57:50.381841 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/2/link_state |
| e000182 | 19:57:50.381841 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/5/nodes/3/flags |
| e000183 | 19:57:50.381841 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /initial/views/5/nodes/3/primary_id |
| e000184 | 19:57:50.381841 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/5/nodes/3/slots |
| e000185 | 19:57:50.381841 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/3/link_state |
| e000186 | 19:57:50.381841 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/5/nodes/4/flags |
| e000187 | 19:57:50.381841 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /initial/views/5/nodes/4/primary_id |
| e000188 | 19:57:50.381841 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /initial/views/5/nodes/4/slots |
| e000189 | 19:57:50.381841 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/4/link_state |
| e000190 | 19:57:50.381841 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/5/nodes/5/flags |
| e000191 | 19:57:50.381841 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /initial/views/5/nodes/5/primary_id |
| e000192 | 19:57:50.381841 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/5/nodes/5/slots |
| e000193 | 19:57:50.381841 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/5/link_state |
| e000194 | 19:57:50.381841 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/5/persistence/aof_enabled |
| e000195 | 19:57:50.381841 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/5/persistence/loading |
| e000196 | 19:57:50.381841 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/5/persistence/aof_last_write_status |
| e000197 | 19:57:50.381841 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/5/persistence/aof_rewrite_in_progress |
| e000198 | 19:57:50.381841 | node-6 | admin | process_live | missing_from_schema | `` | /initial/views/5/process_live |
| e000199 | 19:57:50.381842 | node-4 | node_client | reachable | valid | `true` | /initial/views/3/reachable |
| e000200 | 19:57:50.381842 | node-4 | info_replication | replication.role | valid | `"slave"` | /initial/views/3/replication/role |
| e000201 | 19:57:50.381842 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /initial/views/3/replication/connected_slaves |
| e000202 | 19:57:50.381842 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /initial/views/3/replication/master_link_status |
| e000203 | 19:57:50.381842 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /initial/views/3/replication/master_sync_in_progress |
| e000204 | 19:57:50.381842 | node-4 | info_replication | replication.master_repl_offset | valid | `14750` | /initial/views/3/replication/master_repl_offset |
| e000205 | 19:57:50.381842 | node-4 | info_replication | replication.slave_repl_offset | valid | `14750` | /initial/views/3/replication/slave_repl_offset |
| e000206 | 19:57:50.381842 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /initial/views/3/redis_id |
| e000207 | 19:57:50.381842 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/3/coverage |
| e000208 | 19:57:50.381842 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/3/nodes/0/flags |
| e000209 | 19:57:50.381842 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /initial/views/3/nodes/0/primary_id |
| e000210 | 19:57:50.381842 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/3/nodes/0/slots |
| e000211 | 19:57:50.381842 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/0/link_state |
| e000212 | 19:57:50.381842 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/3/nodes/1/flags |
| e000213 | 19:57:50.381842 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /initial/views/3/nodes/1/primary_id |
| e000214 | 19:57:50.381842 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /initial/views/3/nodes/1/slots |
| e000215 | 19:57:50.381842 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/1/link_state |
| e000216 | 19:57:50.381842 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/3/nodes/2/flags |
| e000217 | 19:57:50.381842 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /initial/views/3/nodes/2/primary_id |
| e000218 | 19:57:50.381842 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /initial/views/3/nodes/2/slots |
| e000219 | 19:57:50.381842 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/2/link_state |
| e000220 | 19:57:50.381842 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /initial/views/3/nodes/3/flags |
| e000221 | 19:57:50.381842 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /initial/views/3/nodes/3/primary_id |
| e000222 | 19:57:50.381842 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /initial/views/3/nodes/3/slots |
| e000223 | 19:57:50.381842 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/3/link_state |
| e000224 | 19:57:50.381842 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/3/nodes/4/flags |
| e000225 | 19:57:50.381842 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /initial/views/3/nodes/4/primary_id |
| e000226 | 19:57:50.381842 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/3/nodes/4/slots |
| e000227 | 19:57:50.381842 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/4/link_state |
| e000228 | 19:57:50.381842 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/3/nodes/5/flags |
| e000229 | 19:57:50.381842 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /initial/views/3/nodes/5/primary_id |
| e000230 | 19:57:50.381842 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/3/nodes/5/slots |
| e000231 | 19:57:50.381842 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/5/link_state |
| e000232 | 19:57:50.381842 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/3/persistence/aof_enabled |
| e000233 | 19:57:50.381842 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/3/persistence/loading |
| e000234 | 19:57:50.381842 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/3/persistence/aof_last_write_status |
| e000235 | 19:57:50.381842 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/3/persistence/aof_rewrite_in_progress |
| e000236 | 19:57:50.381842 | node-4 | admin | process_live | missing_from_schema | `` | /initial/views/3/process_live |
| e000237 | 19:57:50.733451 | node-5 | node_client | reachable | valid | `false` | /samples/0/snapshot/views/4/reachable |
| e000238 | 19:57:50.733451 | node-5 | info_replication | replication.role | collection_failed | `` | /samples/0/snapshot/views/4/replication/role |
| e000239 | 19:57:50.733451 | node-5 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/0/snapshot/views/4/replication/connected_slaves |
| e000240 | 19:57:50.733451 | node-5 | info_replication | replication.master_link_status | collection_failed | `` | /samples/0/snapshot/views/4/replication/master_link_status |
| e000241 | 19:57:50.733451 | node-5 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/0/snapshot/views/4/replication/master_sync_in_progress |
| e000242 | 19:57:50.733451 | node-5 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/0/snapshot/views/4/replication/master_repl_offset |
| e000243 | 19:57:50.733451 | node-5 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/0/snapshot/views/4/replication/slave_repl_offset |
| e000244 | 19:57:50.733451 | node-5 | cluster_nodes | redis_id | collection_failed | `` | /samples/0/snapshot/views/4/redis_id |
| e000245 | 19:57:50.733451 | node-5 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/0/snapshot/views/4/coverage |
| e000246 | 19:57:50.733451 | node-5 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/0/snapshot/views/4/persistence/aof_enabled |
| e000247 | 19:57:50.733451 | node-5 | info_persistence | persistence.loading | collection_failed | `` | /samples/0/snapshot/views/4/persistence/loading |
| e000248 | 19:57:50.733451 | node-5 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/0/snapshot/views/4/persistence/aof_last_write_status |
| e000249 | 19:57:50.733451 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/0/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e000250 | 19:57:50.733451 | node-5 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/4/process_live |
| e000251 | 19:57:50.733461 | node-1 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/0/reachable |
| e000252 | 19:57:50.733461 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/0/snapshot/views/0/replication/role |
| e000253 | 19:57:50.733461 | node-1 | info_replication | replication.connected_slaves | valid | `0` | /samples/0/snapshot/views/0/replication/connected_slaves |
| e000254 | 19:57:50.733461 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/0/snapshot/views/0/replication/master_link_status |
| e000255 | 19:57:50.733461 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/0/snapshot/views/0/replication/master_sync_in_progress |
| e000256 | 19:57:50.733461 | node-1 | info_replication | replication.master_repl_offset | valid | `17856` | /samples/0/snapshot/views/0/replication/master_repl_offset |
| e000257 | 19:57:50.733461 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/0/snapshot/views/0/replication/slave_repl_offset |
| e000258 | 19:57:50.733461 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/0/snapshot/views/0/redis_id |
| e000259 | 19:57:50.733461 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/0/coverage |
| e000260 | 19:57:50.733461 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/0/nodes/0/flags |
| e000261 | 19:57:50.733461 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/0/snapshot/views/0/nodes/0/primary_id |
| e000262 | 19:57:50.733461 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/0/slots |
| e000263 | 19:57:50.733461 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"disconnected"` | /samples/0/snapshot/views/0/nodes/0/link_state |
| e000264 | 19:57:50.733461 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/0/nodes/1/flags |
| e000265 | 19:57:50.733461 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/0/snapshot/views/0/nodes/1/primary_id |
| e000266 | 19:57:50.733461 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/1/slots |
| e000267 | 19:57:50.733461 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/0/nodes/1/link_state |
| e000268 | 19:57:50.733461 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/0/nodes/2/flags |
| e000269 | 19:57:50.733461 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/2/primary_id |
| e000270 | 19:57:50.733461 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/0/nodes/2/slots |
| e000271 | 19:57:50.733461 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/0/nodes/2/link_state |
| e000272 | 19:57:50.733461 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/0/snapshot/views/0/nodes/3/flags |
| e000273 | 19:57:50.733461 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/3/primary_id |
| e000274 | 19:57:50.733461 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/0/nodes/3/slots |
| e000275 | 19:57:50.733461 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/0/nodes/3/link_state |
| e000276 | 19:57:50.733461 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/0/nodes/4/flags |
| e000277 | 19:57:50.733461 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/0/snapshot/views/0/nodes/4/primary_id |
| e000278 | 19:57:50.733461 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/4/slots |
| e000279 | 19:57:50.733461 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/0/nodes/4/link_state |
| e000280 | 19:57:50.733461 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/0/nodes/5/flags |
| e000281 | 19:57:50.733461 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/5/primary_id |
| e000282 | 19:57:50.733461 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/0/nodes/5/slots |
| e000283 | 19:57:50.733461 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/0/nodes/5/link_state |
| e000284 | 19:57:50.733461 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/0/persistence/aof_enabled |
| e000285 | 19:57:50.733461 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/0/persistence/loading |
| e000286 | 19:57:50.733461 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/0/persistence/aof_last_write_status |
| e000287 | 19:57:50.733461 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e000288 | 19:57:50.733461 | node-1 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/0/process_live |
| e000289 | 19:57:50.73354 | node-6 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/5/reachable |
| e000290 | 19:57:50.73354 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/0/snapshot/views/5/replication/role |
| e000291 | 19:57:50.73354 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/0/snapshot/views/5/replication/connected_slaves |
| e000292 | 19:57:50.73354 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/0/snapshot/views/5/replication/master_link_status |
| e000293 | 19:57:50.73354 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/0/snapshot/views/5/replication/master_sync_in_progress |
| e000294 | 19:57:50.73354 | node-6 | info_replication | replication.master_repl_offset | valid | `17858` | /samples/0/snapshot/views/5/replication/master_repl_offset |
| e000295 | 19:57:50.73354 | node-6 | info_replication | replication.slave_repl_offset | valid | `17858` | /samples/0/snapshot/views/5/replication/slave_repl_offset |
| e000296 | 19:57:50.73354 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/0/snapshot/views/5/redis_id |
| e000297 | 19:57:50.73354 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/5/coverage |
| e000298 | 19:57:50.73354 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/0/snapshot/views/5/nodes/0/flags |
| e000299 | 19:57:50.73354 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/0/snapshot/views/5/nodes/0/primary_id |
| e000300 | 19:57:50.73354 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/0/slots |
| e000301 | 19:57:50.73354 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/0/link_state |
| e000302 | 19:57:50.73354 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/5/nodes/1/flags |
| e000303 | 19:57:50.73354 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/1/primary_id |
| e000304 | 19:57:50.73354 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/5/nodes/1/slots |
| e000305 | 19:57:50.73354 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/1/link_state |
| e000306 | 19:57:50.73354 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/5/nodes/2/flags |
| e000307 | 19:57:50.73354 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/0/snapshot/views/5/nodes/2/primary_id |
| e000308 | 19:57:50.73354 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/2/slots |
| e000309 | 19:57:50.73354 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/2/link_state |
| e000310 | 19:57:50.73354 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/5/nodes/3/flags |
| e000311 | 19:57:50.73354 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/3/primary_id |
| e000312 | 19:57:50.73354 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/5/nodes/3/slots |
| e000313 | 19:57:50.73354 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/3/link_state |
| e000314 | 19:57:50.73354 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/5/nodes/4/flags |
| e000315 | 19:57:50.73354 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/0/snapshot/views/5/nodes/4/primary_id |
| e000316 | 19:57:50.73354 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/4/slots |
| e000317 | 19:57:50.73354 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/4/link_state |
| e000318 | 19:57:50.73354 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/5/nodes/5/flags |
| e000319 | 19:57:50.73354 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/5/primary_id |
| e000320 | 19:57:50.73354 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/5/nodes/5/slots |
| e000321 | 19:57:50.73354 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/5/link_state |
| e000322 | 19:57:50.73354 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/5/persistence/aof_enabled |
| e000323 | 19:57:50.73354 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/5/persistence/loading |
| e000324 | 19:57:50.73354 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/5/persistence/aof_last_write_status |
| e000325 | 19:57:50.73354 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e000326 | 19:57:50.73354 | node-6 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/5/process_live |
| e000327 | 19:57:50.733582 | node-2 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/1/reachable |
| e000328 | 19:57:50.733582 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/0/snapshot/views/1/replication/role |
| e000329 | 19:57:50.733582 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/0/snapshot/views/1/replication/connected_slaves |
| e000330 | 19:57:50.733582 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/0/snapshot/views/1/replication/master_link_status |
| e000331 | 19:57:50.733582 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/0/snapshot/views/1/replication/master_sync_in_progress |
| e000332 | 19:57:50.733582 | node-2 | info_replication | replication.master_repl_offset | valid | `17858` | /samples/0/snapshot/views/1/replication/master_repl_offset |
| e000333 | 19:57:50.733582 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/0/snapshot/views/1/replication/slave_repl_offset |
| e000334 | 19:57:50.733582 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/0/snapshot/views/1/redis_id |
| e000335 | 19:57:50.733582 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/1/coverage |
| e000336 | 19:57:50.733582 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/1/nodes/0/flags |
| e000337 | 19:57:50.733582 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/0/primary_id |
| e000338 | 19:57:50.733582 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/1/nodes/0/slots |
| e000339 | 19:57:50.733582 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/0/link_state |
| e000340 | 19:57:50.733582 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/1/nodes/1/flags |
| e000341 | 19:57:50.733582 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/0/snapshot/views/1/nodes/1/primary_id |
| e000342 | 19:57:50.733582 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/1/slots |
| e000343 | 19:57:50.733582 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/1/link_state |
| e000344 | 19:57:50.733582 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/1/nodes/2/flags |
| e000345 | 19:57:50.733582 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/0/snapshot/views/1/nodes/2/primary_id |
| e000346 | 19:57:50.733582 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/2/slots |
| e000347 | 19:57:50.733582 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/2/link_state |
| e000348 | 19:57:50.733582 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/1/nodes/3/flags |
| e000349 | 19:57:50.733582 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/0/snapshot/views/1/nodes/3/primary_id |
| e000350 | 19:57:50.733582 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/3/slots |
| e000351 | 19:57:50.733582 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/3/link_state |
| e000352 | 19:57:50.733582 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/0/snapshot/views/1/nodes/4/flags |
| e000353 | 19:57:50.733582 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/4/primary_id |
| e000354 | 19:57:50.733582 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/1/nodes/4/slots |
| e000355 | 19:57:50.733582 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/4/link_state |
| e000356 | 19:57:50.733582 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/1/nodes/5/flags |
| e000357 | 19:57:50.733582 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/5/primary_id |
| e000358 | 19:57:50.733582 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/1/nodes/5/slots |
| e000359 | 19:57:50.733582 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/5/link_state |
| e000360 | 19:57:50.733582 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/1/persistence/aof_enabled |
| e000361 | 19:57:50.733582 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/1/persistence/loading |
| e000362 | 19:57:50.733582 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/1/persistence/aof_last_write_status |
| e000363 | 19:57:50.733582 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e000364 | 19:57:50.733582 | node-2 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/1/process_live |
| e000365 | 19:57:50.733823 | node-3 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/2/reachable |
| e000366 | 19:57:50.733823 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/0/snapshot/views/2/replication/role |
| e000367 | 19:57:50.733823 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/0/snapshot/views/2/replication/connected_slaves |
| e000368 | 19:57:50.733823 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/0/snapshot/views/2/replication/master_link_status |
| e000369 | 19:57:50.733823 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/0/snapshot/views/2/replication/master_sync_in_progress |
| e000370 | 19:57:50.733823 | node-3 | info_replication | replication.master_repl_offset | valid | `17599` | /samples/0/snapshot/views/2/replication/master_repl_offset |
| e000371 | 19:57:50.733823 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/0/snapshot/views/2/replication/slave_repl_offset |
| e000372 | 19:57:50.733823 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/0/snapshot/views/2/redis_id |
| e000373 | 19:57:50.733823 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/2/coverage |
| e000374 | 19:57:50.733823 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/2/nodes/0/flags |
| e000375 | 19:57:50.733823 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/0/snapshot/views/2/nodes/0/primary_id |
| e000376 | 19:57:50.733823 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/0/slots |
| e000377 | 19:57:50.733823 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/0/link_state |
| e000378 | 19:57:50.733823 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/0/snapshot/views/2/nodes/1/flags |
| e000379 | 19:57:50.733823 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/1/primary_id |
| e000380 | 19:57:50.733823 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/2/nodes/1/slots |
| e000381 | 19:57:50.733823 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/1/link_state |
| e000382 | 19:57:50.733823 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/2/nodes/2/flags |
| e000383 | 19:57:50.733823 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/0/snapshot/views/2/nodes/2/primary_id |
| e000384 | 19:57:50.733823 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/2/slots |
| e000385 | 19:57:50.733823 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/2/link_state |
| e000386 | 19:57:50.733823 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/2/nodes/3/flags |
| e000387 | 19:57:50.733823 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/3/primary_id |
| e000388 | 19:57:50.733823 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/2/nodes/3/slots |
| e000389 | 19:57:50.733823 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/3/link_state |
| e000390 | 19:57:50.733823 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/2/nodes/4/flags |
| e000391 | 19:57:50.733823 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/4/primary_id |
| e000392 | 19:57:50.733823 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/2/nodes/4/slots |
| e000393 | 19:57:50.733823 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/4/link_state |
| e000394 | 19:57:50.733823 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/2/nodes/5/flags |
| e000395 | 19:57:50.733823 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/0/snapshot/views/2/nodes/5/primary_id |
| e000396 | 19:57:50.733823 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/5/slots |
| e000397 | 19:57:50.733823 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/5/link_state |
| e000398 | 19:57:50.733823 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/2/persistence/aof_enabled |
| e000399 | 19:57:50.733823 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/2/persistence/loading |
| e000400 | 19:57:50.733823 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/2/persistence/aof_last_write_status |
| e000401 | 19:57:50.733823 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e000402 | 19:57:50.733823 | node-3 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/2/process_live |
| e000403 | 19:57:50.733985 | node-4 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/3/reachable |
| e000404 | 19:57:50.733985 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/0/snapshot/views/3/replication/role |
| e000405 | 19:57:50.733985 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/0/snapshot/views/3/replication/connected_slaves |
| e000406 | 19:57:50.733985 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/0/snapshot/views/3/replication/master_link_status |
| e000407 | 19:57:50.733985 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/0/snapshot/views/3/replication/master_sync_in_progress |
| e000408 | 19:57:50.733985 | node-4 | info_replication | replication.master_repl_offset | valid | `17599` | /samples/0/snapshot/views/3/replication/master_repl_offset |
| e000409 | 19:57:50.733985 | node-4 | info_replication | replication.slave_repl_offset | valid | `17599` | /samples/0/snapshot/views/3/replication/slave_repl_offset |
| e000410 | 19:57:50.733985 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/0/snapshot/views/3/redis_id |
| e000411 | 19:57:50.733985 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/3/coverage |
| e000412 | 19:57:50.733985 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/3/nodes/0/flags |
| e000413 | 19:57:50.733985 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/0/primary_id |
| e000414 | 19:57:50.733985 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/3/nodes/0/slots |
| e000415 | 19:57:50.733985 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/0/link_state |
| e000416 | 19:57:50.733985 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/3/nodes/1/flags |
| e000417 | 19:57:50.733985 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/0/snapshot/views/3/nodes/1/primary_id |
| e000418 | 19:57:50.733985 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/1/slots |
| e000419 | 19:57:50.733985 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/1/link_state |
| e000420 | 19:57:50.733985 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/3/nodes/2/flags |
| e000421 | 19:57:50.733985 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/0/snapshot/views/3/nodes/2/primary_id |
| e000422 | 19:57:50.733985 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/2/slots |
| e000423 | 19:57:50.733985 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/2/link_state |
| e000424 | 19:57:50.733985 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/0/snapshot/views/3/nodes/3/flags |
| e000425 | 19:57:50.733985 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/0/snapshot/views/3/nodes/3/primary_id |
| e000426 | 19:57:50.733985 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/3/slots |
| e000427 | 19:57:50.733985 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/3/link_state |
| e000428 | 19:57:50.733985 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/3/nodes/4/flags |
| e000429 | 19:57:50.733985 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/4/primary_id |
| e000430 | 19:57:50.733985 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/3/nodes/4/slots |
| e000431 | 19:57:50.733985 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/4/link_state |
| e000432 | 19:57:50.733985 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/3/nodes/5/flags |
| e000433 | 19:57:50.733985 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/5/primary_id |
| e000434 | 19:57:50.733985 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/3/nodes/5/slots |
| e000435 | 19:57:50.733985 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/5/link_state |
| e000436 | 19:57:50.733985 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/3/persistence/aof_enabled |
| e000437 | 19:57:50.733985 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/3/persistence/loading |
| e000438 | 19:57:50.733985 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/3/persistence/aof_last_write_status |
| e000439 | 19:57:50.733985 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e000440 | 19:57:50.733985 | node-4 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/3/process_live |
| e000441 | 19:57:50.86096 |  | client_probe | client_success | valid | `true` | /samples/0/probe/successful |
| e000442 | 19:57:51.049017 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/0/write_policy_probe/policy |
| e000443 | 19:57:51.049017 |  | ack_probe | ack.successful | valid | `false` | /samples/0/write_policy_probe/successful |
| e000444 | 19:57:51.049017 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/0/write_policy_probe/requested_replicas |
| e000445 | 19:57:51.049017 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/0/write_policy_probe/set_acknowledged |
| e000446 | 19:57:51.049017 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          0
        ]` | /samples/0/write_policy_probe/acknowledgements |
| e000447 | 19:57:51.049017 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/0/write_policy_probe/acknowledgement_observed |
| e000448 | 19:57:51.049017 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/0/write_policy_probe/waitaof_requested_local |
| e000449 | 19:57:51.049017 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/0/write_policy_probe/waitaof_requested_replicas |
| e000450 | 19:57:51.049017 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/0/write_policy_probe/waitaof_local_acknowledgements |
| e000451 | 19:57:51.049017 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/0/write_policy_probe/waitaof_replica_acknowledgements |
| e000452 | 19:57:51.049017 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/0/write_policy_probe/waitaof_acknowledgement_observed |
| e000453 | 19:57:51.381014 |  | client_window | client.attempts | valid | `167` | /metrics/windows/2/attempts |
| e000454 | 19:57:51.381014 |  | client_window | client.errors | valid | `18` | /metrics/windows/2/errors |
| e000455 | 19:57:51.381014 |  | client_window | client.successes | valid | `149` | /metrics/windows/2/successes |
| e000456 | 19:57:51.381014 |  | client_window | client.p95_ms | valid | `111.411023` | /metrics/windows/2/p95_ms |
| e000457 | 19:57:51.393924 | node-5 | node_client | reachable | valid | `false` | /samples/1/snapshot/views/4/reachable |
| e000458 | 19:57:51.393924 | node-5 | info_replication | replication.role | collection_failed | `` | /samples/1/snapshot/views/4/replication/role |
| e000459 | 19:57:51.393924 | node-5 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/1/snapshot/views/4/replication/connected_slaves |
| e000460 | 19:57:51.393924 | node-5 | info_replication | replication.master_link_status | collection_failed | `` | /samples/1/snapshot/views/4/replication/master_link_status |
| e000461 | 19:57:51.393924 | node-5 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/1/snapshot/views/4/replication/master_sync_in_progress |
| e000462 | 19:57:51.393924 | node-5 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/1/snapshot/views/4/replication/master_repl_offset |
| e000463 | 19:57:51.393924 | node-5 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/1/snapshot/views/4/replication/slave_repl_offset |
| e000464 | 19:57:51.393924 | node-5 | cluster_nodes | redis_id | collection_failed | `` | /samples/1/snapshot/views/4/redis_id |
| e000465 | 19:57:51.393924 | node-5 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/1/snapshot/views/4/coverage |
| e000466 | 19:57:51.393924 | node-5 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/1/snapshot/views/4/persistence/aof_enabled |
| e000467 | 19:57:51.393924 | node-5 | info_persistence | persistence.loading | collection_failed | `` | /samples/1/snapshot/views/4/persistence/loading |
| e000468 | 19:57:51.393924 | node-5 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/1/snapshot/views/4/persistence/aof_last_write_status |
| e000469 | 19:57:51.393924 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/1/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e000470 | 19:57:51.393924 | node-5 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/4/process_live |
| e000471 | 19:57:51.393926 | node-3 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/2/reachable |
| e000472 | 19:57:51.393926 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/1/snapshot/views/2/replication/role |
| e000473 | 19:57:51.393926 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/1/snapshot/views/2/replication/connected_slaves |
| e000474 | 19:57:51.393926 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/1/snapshot/views/2/replication/master_link_status |
| e000475 | 19:57:51.393926 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/1/snapshot/views/2/replication/master_sync_in_progress |
| e000476 | 19:57:51.393926 | node-3 | info_replication | replication.master_repl_offset | valid | `23010` | /samples/1/snapshot/views/2/replication/master_repl_offset |
| e000477 | 19:57:51.393926 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/1/snapshot/views/2/replication/slave_repl_offset |
| e000478 | 19:57:51.393926 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/1/snapshot/views/2/redis_id |
| e000479 | 19:57:51.393926 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/2/coverage |
| e000480 | 19:57:51.393926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/2/nodes/0/flags |
| e000481 | 19:57:51.393926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/1/snapshot/views/2/nodes/0/primary_id |
| e000482 | 19:57:51.393926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/0/slots |
| e000483 | 19:57:51.393926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/0/link_state |
| e000484 | 19:57:51.393926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/1/snapshot/views/2/nodes/1/flags |
| e000485 | 19:57:51.393926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/1/primary_id |
| e000486 | 19:57:51.393926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/2/nodes/1/slots |
| e000487 | 19:57:51.393926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/1/link_state |
| e000488 | 19:57:51.393926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/2/nodes/2/flags |
| e000489 | 19:57:51.393926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/1/snapshot/views/2/nodes/2/primary_id |
| e000490 | 19:57:51.393926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/2/slots |
| e000491 | 19:57:51.393926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/2/link_state |
| e000492 | 19:57:51.393926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/2/nodes/3/flags |
| e000493 | 19:57:51.393926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/3/primary_id |
| e000494 | 19:57:51.393926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/2/nodes/3/slots |
| e000495 | 19:57:51.393926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/3/link_state |
| e000496 | 19:57:51.393926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/2/nodes/4/flags |
| e000497 | 19:57:51.393926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/4/primary_id |
| e000498 | 19:57:51.393926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/2/nodes/4/slots |
| e000499 | 19:57:51.393926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/4/link_state |
| e000500 | 19:57:51.393926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/2/nodes/5/flags |
| e000501 | 19:57:51.393926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/1/snapshot/views/2/nodes/5/primary_id |
| e000502 | 19:57:51.393926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/5/slots |
| e000503 | 19:57:51.393926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/5/link_state |
| e000504 | 19:57:51.393926 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/2/persistence/aof_enabled |
| e000505 | 19:57:51.393926 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/2/persistence/loading |
| e000506 | 19:57:51.393926 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/2/persistence/aof_last_write_status |
| e000507 | 19:57:51.393926 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e000508 | 19:57:51.393926 | node-3 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/2/process_live |
| e000509 | 19:57:51.393964 | node-2 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/1/reachable |
| e000510 | 19:57:51.393964 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/1/snapshot/views/1/replication/role |
| e000511 | 19:57:51.393964 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/1/snapshot/views/1/replication/connected_slaves |
| e000512 | 19:57:51.393964 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/1/snapshot/views/1/replication/master_link_status |
| e000513 | 19:57:51.393964 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/1/snapshot/views/1/replication/master_sync_in_progress |
| e000514 | 19:57:51.393964 | node-2 | info_replication | replication.master_repl_offset | valid | `23010` | /samples/1/snapshot/views/1/replication/master_repl_offset |
| e000515 | 19:57:51.393964 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/1/snapshot/views/1/replication/slave_repl_offset |
| e000516 | 19:57:51.393964 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/1/snapshot/views/1/redis_id |
| e000517 | 19:57:51.393964 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/1/coverage |
| e000518 | 19:57:51.393964 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/1/nodes/0/flags |
| e000519 | 19:57:51.393964 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/0/primary_id |
| e000520 | 19:57:51.393964 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/1/nodes/0/slots |
| e000521 | 19:57:51.393964 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/0/link_state |
| e000522 | 19:57:51.393964 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/1/nodes/1/flags |
| e000523 | 19:57:51.393964 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/1/snapshot/views/1/nodes/1/primary_id |
| e000524 | 19:57:51.393964 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/1/slots |
| e000525 | 19:57:51.393964 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/1/link_state |
| e000526 | 19:57:51.393964 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/1/nodes/2/flags |
| e000527 | 19:57:51.393964 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/1/snapshot/views/1/nodes/2/primary_id |
| e000528 | 19:57:51.393964 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/2/slots |
| e000529 | 19:57:51.393964 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/2/link_state |
| e000530 | 19:57:51.393964 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/1/nodes/3/flags |
| e000531 | 19:57:51.393964 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/1/snapshot/views/1/nodes/3/primary_id |
| e000532 | 19:57:51.393964 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/3/slots |
| e000533 | 19:57:51.393964 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/3/link_state |
| e000534 | 19:57:51.393964 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/1/snapshot/views/1/nodes/4/flags |
| e000535 | 19:57:51.393964 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/4/primary_id |
| e000536 | 19:57:51.393964 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/1/nodes/4/slots |
| e000537 | 19:57:51.393964 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/4/link_state |
| e000538 | 19:57:51.393964 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/1/nodes/5/flags |
| e000539 | 19:57:51.393964 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/5/primary_id |
| e000540 | 19:57:51.393964 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/1/nodes/5/slots |
| e000541 | 19:57:51.393964 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/5/link_state |
| e000542 | 19:57:51.393964 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/1/persistence/aof_enabled |
| e000543 | 19:57:51.393964 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/1/persistence/loading |
| e000544 | 19:57:51.393964 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/1/persistence/aof_last_write_status |
| e000545 | 19:57:51.393964 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e000546 | 19:57:51.393964 | node-2 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/1/process_live |
| e000547 | 19:57:51.393974 | node-4 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/3/reachable |
| e000548 | 19:57:51.393974 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/1/snapshot/views/3/replication/role |
| e000549 | 19:57:51.393974 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/1/snapshot/views/3/replication/connected_slaves |
| e000550 | 19:57:51.393974 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/1/snapshot/views/3/replication/master_link_status |
| e000551 | 19:57:51.393974 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/1/snapshot/views/3/replication/master_sync_in_progress |
| e000552 | 19:57:51.393974 | node-4 | info_replication | replication.master_repl_offset | valid | `23010` | /samples/1/snapshot/views/3/replication/master_repl_offset |
| e000553 | 19:57:51.393974 | node-4 | info_replication | replication.slave_repl_offset | valid | `23010` | /samples/1/snapshot/views/3/replication/slave_repl_offset |
| e000554 | 19:57:51.393974 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/1/snapshot/views/3/redis_id |
| e000555 | 19:57:51.393974 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/3/coverage |
| e000556 | 19:57:51.393974 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/3/nodes/0/flags |
| e000557 | 19:57:51.393974 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/0/primary_id |
| e000558 | 19:57:51.393974 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/3/nodes/0/slots |
| e000559 | 19:57:51.393974 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/0/link_state |
| e000560 | 19:57:51.393974 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/3/nodes/1/flags |
| e000561 | 19:57:51.393974 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/1/snapshot/views/3/nodes/1/primary_id |
| e000562 | 19:57:51.393974 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/1/slots |
| e000563 | 19:57:51.393974 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/1/link_state |
| e000564 | 19:57:51.393974 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/3/nodes/2/flags |
| e000565 | 19:57:51.393974 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/1/snapshot/views/3/nodes/2/primary_id |
| e000566 | 19:57:51.393974 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/2/slots |
| e000567 | 19:57:51.393974 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/2/link_state |
| e000568 | 19:57:51.393974 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/1/snapshot/views/3/nodes/3/flags |
| e000569 | 19:57:51.393974 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/1/snapshot/views/3/nodes/3/primary_id |
| e000570 | 19:57:51.393974 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/3/slots |
| e000571 | 19:57:51.393974 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/3/link_state |
| e000572 | 19:57:51.393974 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/3/nodes/4/flags |
| e000573 | 19:57:51.393974 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/4/primary_id |
| e000574 | 19:57:51.393974 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/3/nodes/4/slots |
| e000575 | 19:57:51.393974 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/4/link_state |
| e000576 | 19:57:51.393974 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/3/nodes/5/flags |
| e000577 | 19:57:51.393974 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/5/primary_id |
| e000578 | 19:57:51.393974 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/3/nodes/5/slots |
| e000579 | 19:57:51.393974 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/5/link_state |
| e000580 | 19:57:51.393974 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/3/persistence/aof_enabled |
| e000581 | 19:57:51.393974 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/3/persistence/loading |
| e000582 | 19:57:51.393974 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/3/persistence/aof_last_write_status |
| e000583 | 19:57:51.393974 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e000584 | 19:57:51.393974 | node-4 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/3/process_live |
| e000585 | 19:57:51.393999 | node-6 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/5/reachable |
| e000586 | 19:57:51.393999 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/1/snapshot/views/5/replication/role |
| e000587 | 19:57:51.393999 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/1/snapshot/views/5/replication/connected_slaves |
| e000588 | 19:57:51.393999 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/1/snapshot/views/5/replication/master_link_status |
| e000589 | 19:57:51.393999 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/1/snapshot/views/5/replication/master_sync_in_progress |
| e000590 | 19:57:51.393999 | node-6 | info_replication | replication.master_repl_offset | valid | `23010` | /samples/1/snapshot/views/5/replication/master_repl_offset |
| e000591 | 19:57:51.393999 | node-6 | info_replication | replication.slave_repl_offset | valid | `23010` | /samples/1/snapshot/views/5/replication/slave_repl_offset |
| e000592 | 19:57:51.393999 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/1/snapshot/views/5/redis_id |
| e000593 | 19:57:51.393999 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/5/coverage |
| e000594 | 19:57:51.393999 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/1/snapshot/views/5/nodes/0/flags |
| e000595 | 19:57:51.393999 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/1/snapshot/views/5/nodes/0/primary_id |
| e000596 | 19:57:51.393999 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/0/slots |
| e000597 | 19:57:51.393999 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/0/link_state |
| e000598 | 19:57:51.393999 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/5/nodes/1/flags |
| e000599 | 19:57:51.393999 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/1/primary_id |
| e000600 | 19:57:51.393999 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/5/nodes/1/slots |
| e000601 | 19:57:51.393999 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/1/link_state |
| e000602 | 19:57:51.393999 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/5/nodes/2/flags |
| e000603 | 19:57:51.393999 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/1/snapshot/views/5/nodes/2/primary_id |
| e000604 | 19:57:51.393999 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/2/slots |
| e000605 | 19:57:51.393999 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/2/link_state |
| e000606 | 19:57:51.393999 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/5/nodes/3/flags |
| e000607 | 19:57:51.393999 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/3/primary_id |
| e000608 | 19:57:51.393999 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/5/nodes/3/slots |
| e000609 | 19:57:51.393999 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/3/link_state |
| e000610 | 19:57:51.393999 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/5/nodes/4/flags |
| e000611 | 19:57:51.393999 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/1/snapshot/views/5/nodes/4/primary_id |
| e000612 | 19:57:51.393999 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/4/slots |
| e000613 | 19:57:51.393999 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/4/link_state |
| e000614 | 19:57:51.393999 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/5/nodes/5/flags |
| e000615 | 19:57:51.393999 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/5/primary_id |
| e000616 | 19:57:51.393999 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/5/nodes/5/slots |
| e000617 | 19:57:51.393999 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/5/link_state |
| e000618 | 19:57:51.393999 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/5/persistence/aof_enabled |
| e000619 | 19:57:51.393999 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/5/persistence/loading |
| e000620 | 19:57:51.393999 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/5/persistence/aof_last_write_status |
| e000621 | 19:57:51.393999 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e000622 | 19:57:51.393999 | node-6 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/5/process_live |
| e000623 | 19:57:51.394098 | node-1 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/0/reachable |
| e000624 | 19:57:51.394098 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/1/snapshot/views/0/replication/role |
| e000625 | 19:57:51.394098 | node-1 | info_replication | replication.connected_slaves | valid | `0` | /samples/1/snapshot/views/0/replication/connected_slaves |
| e000626 | 19:57:51.394098 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/1/snapshot/views/0/replication/master_link_status |
| e000627 | 19:57:51.394098 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/1/snapshot/views/0/replication/master_sync_in_progress |
| e000628 | 19:57:51.394098 | node-1 | info_replication | replication.master_repl_offset | valid | `23230` | /samples/1/snapshot/views/0/replication/master_repl_offset |
| e000629 | 19:57:51.394098 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/1/snapshot/views/0/replication/slave_repl_offset |
| e000630 | 19:57:51.394098 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/1/snapshot/views/0/redis_id |
| e000631 | 19:57:51.394098 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/0/coverage |
| e000632 | 19:57:51.394098 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/0/nodes/0/flags |
| e000633 | 19:57:51.394098 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/1/snapshot/views/0/nodes/0/primary_id |
| e000634 | 19:57:51.394098 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/0/slots |
| e000635 | 19:57:51.394098 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/0/link_state |
| e000636 | 19:57:51.394098 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/0/nodes/1/flags |
| e000637 | 19:57:51.394098 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/1/snapshot/views/0/nodes/1/primary_id |
| e000638 | 19:57:51.394098 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/1/slots |
| e000639 | 19:57:51.394098 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/1/link_state |
| e000640 | 19:57:51.394098 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/0/nodes/2/flags |
| e000641 | 19:57:51.394098 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/2/primary_id |
| e000642 | 19:57:51.394098 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/0/nodes/2/slots |
| e000643 | 19:57:51.394098 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/2/link_state |
| e000644 | 19:57:51.394098 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/1/snapshot/views/0/nodes/3/flags |
| e000645 | 19:57:51.394098 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/3/primary_id |
| e000646 | 19:57:51.394098 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/0/nodes/3/slots |
| e000647 | 19:57:51.394098 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/3/link_state |
| e000648 | 19:57:51.394098 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/0/nodes/4/flags |
| e000649 | 19:57:51.394098 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/1/snapshot/views/0/nodes/4/primary_id |
| e000650 | 19:57:51.394098 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/4/slots |
| e000651 | 19:57:51.394098 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/4/link_state |
| e000652 | 19:57:51.394098 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/0/nodes/5/flags |
| e000653 | 19:57:51.394098 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/5/primary_id |
| e000654 | 19:57:51.394098 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/0/nodes/5/slots |
| e000655 | 19:57:51.394098 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/5/link_state |
| e000656 | 19:57:51.394098 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/0/persistence/aof_enabled |
| e000657 | 19:57:51.394098 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/0/persistence/loading |
| e000658 | 19:57:51.394098 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/0/persistence/aof_last_write_status |
| e000659 | 19:57:51.394098 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e000660 | 19:57:51.394098 | node-1 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/0/process_live |
| e000661 | 19:57:51.520037 |  | client_probe | client_success | valid | `true` | /samples/1/probe/successful |
| e000662 | 19:57:51.741115 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/1/write_policy_probe/policy |
| e000663 | 19:57:51.741115 |  | ack_probe | ack.successful | valid | `false` | /samples/1/write_policy_probe/successful |
| e000664 | 19:57:51.741115 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/1/write_policy_probe/requested_replicas |
| e000665 | 19:57:51.741115 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/1/write_policy_probe/set_acknowledged |
| e000666 | 19:57:51.741115 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          0
        ]` | /samples/1/write_policy_probe/acknowledgements |
| e000667 | 19:57:51.741115 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/1/write_policy_probe/acknowledgement_observed |
| e000668 | 19:57:51.741115 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/1/write_policy_probe/waitaof_requested_local |
| e000669 | 19:57:51.741115 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/1/write_policy_probe/waitaof_requested_replicas |
| e000670 | 19:57:51.741115 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/1/write_policy_probe/waitaof_local_acknowledgements |
| e000671 | 19:57:51.741115 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/1/write_policy_probe/waitaof_replica_acknowledgements |
| e000672 | 19:57:51.741115 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/1/write_policy_probe/waitaof_acknowledgement_observed |
| e000673 | 19:57:52.087733 | node-2 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/1/reachable |
| e000674 | 19:57:52.087733 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/2/snapshot/views/1/replication/role |
| e000675 | 19:57:52.087733 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/2/snapshot/views/1/replication/connected_slaves |
| e000676 | 19:57:52.087733 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/2/snapshot/views/1/replication/master_link_status |
| e000677 | 19:57:52.087733 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/2/snapshot/views/1/replication/master_sync_in_progress |
| e000678 | 19:57:52.087733 | node-2 | info_replication | replication.master_repl_offset | valid | `28421` | /samples/2/snapshot/views/1/replication/master_repl_offset |
| e000679 | 19:57:52.087733 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/2/snapshot/views/1/replication/slave_repl_offset |
| e000680 | 19:57:52.087733 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/2/snapshot/views/1/redis_id |
| e000681 | 19:57:52.087733 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/2/snapshot/views/1/coverage |
| e000682 | 19:57:52.087733 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/1/nodes/0/flags |
| e000683 | 19:57:52.087733 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/0/primary_id |
| e000684 | 19:57:52.087733 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/1/nodes/0/slots |
| e000685 | 19:57:52.087733 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/0/link_state |
| e000686 | 19:57:52.087733 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/1/nodes/1/flags |
| e000687 | 19:57:52.087733 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/2/snapshot/views/1/nodes/1/primary_id |
| e000688 | 19:57:52.087733 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/1/slots |
| e000689 | 19:57:52.087733 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/1/link_state |
| e000690 | 19:57:52.087733 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/1/nodes/2/flags |
| e000691 | 19:57:52.087733 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/2/snapshot/views/1/nodes/2/primary_id |
| e000692 | 19:57:52.087733 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/2/slots |
| e000693 | 19:57:52.087733 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/2/link_state |
| e000694 | 19:57:52.087733 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/1/nodes/3/flags |
| e000695 | 19:57:52.087733 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/2/snapshot/views/1/nodes/3/primary_id |
| e000696 | 19:57:52.087733 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/3/slots |
| e000697 | 19:57:52.087733 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/3/link_state |
| e000698 | 19:57:52.087733 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/2/snapshot/views/1/nodes/4/flags |
| e000699 | 19:57:52.087733 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/4/primary_id |
| e000700 | 19:57:52.087733 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/1/nodes/4/slots |
| e000701 | 19:57:52.087733 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/4/link_state |
| e000702 | 19:57:52.087733 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/1/nodes/5/flags |
| e000703 | 19:57:52.087733 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/5/primary_id |
| e000704 | 19:57:52.087733 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/1/nodes/5/slots |
| e000705 | 19:57:52.087733 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/5/link_state |
| e000706 | 19:57:52.087733 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/1/persistence/aof_enabled |
| e000707 | 19:57:52.087733 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/1/persistence/loading |
| e000708 | 19:57:52.087733 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/1/persistence/aof_last_write_status |
| e000709 | 19:57:52.087733 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e000710 | 19:57:52.087733 | node-2 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/1/process_live |
| e000711 | 19:57:52.087742 | node-5 | node_client | reachable | valid | `false` | /samples/2/snapshot/views/4/reachable |
| e000712 | 19:57:52.087742 | node-5 | info_replication | replication.role | collection_failed | `` | /samples/2/snapshot/views/4/replication/role |
| e000713 | 19:57:52.087742 | node-5 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/2/snapshot/views/4/replication/connected_slaves |
| e000714 | 19:57:52.087742 | node-5 | info_replication | replication.master_link_status | collection_failed | `` | /samples/2/snapshot/views/4/replication/master_link_status |
| e000715 | 19:57:52.087742 | node-5 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/2/snapshot/views/4/replication/master_sync_in_progress |
| e000716 | 19:57:52.087742 | node-5 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/2/snapshot/views/4/replication/master_repl_offset |
| e000717 | 19:57:52.087742 | node-5 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/2/snapshot/views/4/replication/slave_repl_offset |
| e000718 | 19:57:52.087742 | node-5 | cluster_nodes | redis_id | collection_failed | `` | /samples/2/snapshot/views/4/redis_id |
| e000719 | 19:57:52.087742 | node-5 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/2/snapshot/views/4/coverage |
| e000720 | 19:57:52.087742 | node-5 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/2/snapshot/views/4/persistence/aof_enabled |
| e000721 | 19:57:52.087742 | node-5 | info_persistence | persistence.loading | collection_failed | `` | /samples/2/snapshot/views/4/persistence/loading |
| e000722 | 19:57:52.087742 | node-5 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/2/snapshot/views/4/persistence/aof_last_write_status |
| e000723 | 19:57:52.087742 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/2/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e000724 | 19:57:52.087742 | node-5 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/4/process_live |
| e000725 | 19:57:52.087752 | node-4 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/3/reachable |
| e000726 | 19:57:52.087752 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/2/snapshot/views/3/replication/role |
| e000727 | 19:57:52.087752 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/2/snapshot/views/3/replication/connected_slaves |
| e000728 | 19:57:52.087752 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/2/snapshot/views/3/replication/master_link_status |
| e000729 | 19:57:52.087752 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/2/snapshot/views/3/replication/master_sync_in_progress |
| e000730 | 19:57:52.087752 | node-4 | info_replication | replication.master_repl_offset | valid | `28162` | /samples/2/snapshot/views/3/replication/master_repl_offset |
| e000731 | 19:57:52.087752 | node-4 | info_replication | replication.slave_repl_offset | valid | `28162` | /samples/2/snapshot/views/3/replication/slave_repl_offset |
| e000732 | 19:57:52.087752 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/2/snapshot/views/3/redis_id |
| e000733 | 19:57:52.087752 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/2/snapshot/views/3/coverage |
| e000734 | 19:57:52.087752 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/3/nodes/0/flags |
| e000735 | 19:57:52.087752 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/0/primary_id |
| e000736 | 19:57:52.087752 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/3/nodes/0/slots |
| e000737 | 19:57:52.087752 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/0/link_state |
| e000738 | 19:57:52.087752 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave",
                  "fail?"
                ]` | /samples/2/snapshot/views/3/nodes/1/flags |
| e000739 | 19:57:52.087752 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/2/snapshot/views/3/nodes/1/primary_id |
| e000740 | 19:57:52.087752 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/1/slots |
| e000741 | 19:57:52.087752 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/1/link_state |
| e000742 | 19:57:52.087752 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/3/nodes/2/flags |
| e000743 | 19:57:52.087752 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/2/snapshot/views/3/nodes/2/primary_id |
| e000744 | 19:57:52.087752 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/2/slots |
| e000745 | 19:57:52.087752 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/2/link_state |
| e000746 | 19:57:52.087752 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/2/snapshot/views/3/nodes/3/flags |
| e000747 | 19:57:52.087752 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/2/snapshot/views/3/nodes/3/primary_id |
| e000748 | 19:57:52.087752 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/3/slots |
| e000749 | 19:57:52.087752 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/3/link_state |
| e000750 | 19:57:52.087752 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/3/nodes/4/flags |
| e000751 | 19:57:52.087752 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/4/primary_id |
| e000752 | 19:57:52.087752 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/3/nodes/4/slots |
| e000753 | 19:57:52.087752 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/4/link_state |
| e000754 | 19:57:52.087752 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/3/nodes/5/flags |
| e000755 | 19:57:52.087752 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/5/primary_id |
| e000756 | 19:57:52.087752 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/3/nodes/5/slots |
| e000757 | 19:57:52.087752 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/5/link_state |
| e000758 | 19:57:52.087752 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/3/persistence/aof_enabled |
| e000759 | 19:57:52.087752 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/3/persistence/loading |
| e000760 | 19:57:52.087752 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/3/persistence/aof_last_write_status |
| e000761 | 19:57:52.087752 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e000762 | 19:57:52.087752 | node-4 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/3/process_live |
| e000763 | 19:57:52.087756 | node-3 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/2/reachable |
| e000764 | 19:57:52.087756 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/2/snapshot/views/2/replication/role |
| e000765 | 19:57:52.087756 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/2/snapshot/views/2/replication/connected_slaves |
| e000766 | 19:57:52.087756 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/2/snapshot/views/2/replication/master_link_status |
| e000767 | 19:57:52.087756 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/2/snapshot/views/2/replication/master_sync_in_progress |
| e000768 | 19:57:52.087756 | node-3 | info_replication | replication.master_repl_offset | valid | `28162` | /samples/2/snapshot/views/2/replication/master_repl_offset |
| e000769 | 19:57:52.087756 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/2/snapshot/views/2/replication/slave_repl_offset |
| e000770 | 19:57:52.087756 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/2/snapshot/views/2/redis_id |
| e000771 | 19:57:52.087756 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/2/snapshot/views/2/coverage |
| e000772 | 19:57:52.087756 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/2/nodes/0/flags |
| e000773 | 19:57:52.087756 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/2/snapshot/views/2/nodes/0/primary_id |
| e000774 | 19:57:52.087756 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/0/slots |
| e000775 | 19:57:52.087756 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/0/link_state |
| e000776 | 19:57:52.087756 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/2/snapshot/views/2/nodes/1/flags |
| e000777 | 19:57:52.087756 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/1/primary_id |
| e000778 | 19:57:52.087756 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/2/nodes/1/slots |
| e000779 | 19:57:52.087756 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/1/link_state |
| e000780 | 19:57:52.087756 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/2/nodes/2/flags |
| e000781 | 19:57:52.087756 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/2/snapshot/views/2/nodes/2/primary_id |
| e000782 | 19:57:52.087756 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/2/slots |
| e000783 | 19:57:52.087756 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/2/link_state |
| e000784 | 19:57:52.087756 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/2/nodes/3/flags |
| e000785 | 19:57:52.087756 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/3/primary_id |
| e000786 | 19:57:52.087756 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/2/nodes/3/slots |
| e000787 | 19:57:52.087756 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/3/link_state |
| e000788 | 19:57:52.087756 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/2/nodes/4/flags |
| e000789 | 19:57:52.087756 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/4/primary_id |
| e000790 | 19:57:52.087756 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/2/nodes/4/slots |
| e000791 | 19:57:52.087756 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/4/link_state |
| e000792 | 19:57:52.087756 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/2/nodes/5/flags |
| e000793 | 19:57:52.087756 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/2/snapshot/views/2/nodes/5/primary_id |
| e000794 | 19:57:52.087756 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/5/slots |
| e000795 | 19:57:52.087756 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/5/link_state |
| e000796 | 19:57:52.087756 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/2/persistence/aof_enabled |
| e000797 | 19:57:52.087756 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/2/persistence/loading |
| e000798 | 19:57:52.087756 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/2/persistence/aof_last_write_status |
| e000799 | 19:57:52.087756 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e000800 | 19:57:52.087756 | node-3 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/2/process_live |
| e000801 | 19:57:52.087774 | node-6 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/5/reachable |
| e000802 | 19:57:52.087774 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/2/snapshot/views/5/replication/role |
| e000803 | 19:57:52.087774 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/2/snapshot/views/5/replication/connected_slaves |
| e000804 | 19:57:52.087774 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/2/snapshot/views/5/replication/master_link_status |
| e000805 | 19:57:52.087774 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/2/snapshot/views/5/replication/master_sync_in_progress |
| e000806 | 19:57:52.087774 | node-6 | info_replication | replication.master_repl_offset | valid | `28421` | /samples/2/snapshot/views/5/replication/master_repl_offset |
| e000807 | 19:57:52.087774 | node-6 | info_replication | replication.slave_repl_offset | valid | `28421` | /samples/2/snapshot/views/5/replication/slave_repl_offset |
| e000808 | 19:57:52.087774 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/2/snapshot/views/5/redis_id |
| e000809 | 19:57:52.087774 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/2/snapshot/views/5/coverage |
| e000810 | 19:57:52.087774 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/2/snapshot/views/5/nodes/0/flags |
| e000811 | 19:57:52.087774 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/2/snapshot/views/5/nodes/0/primary_id |
| e000812 | 19:57:52.087774 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/0/slots |
| e000813 | 19:57:52.087774 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/0/link_state |
| e000814 | 19:57:52.087774 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/5/nodes/1/flags |
| e000815 | 19:57:52.087774 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/1/primary_id |
| e000816 | 19:57:52.087774 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/5/nodes/1/slots |
| e000817 | 19:57:52.087774 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/1/link_state |
| e000818 | 19:57:52.087774 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave",
                  "fail?"
                ]` | /samples/2/snapshot/views/5/nodes/2/flags |
| e000819 | 19:57:52.087774 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/2/snapshot/views/5/nodes/2/primary_id |
| e000820 | 19:57:52.087774 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/2/slots |
| e000821 | 19:57:52.087774 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/2/link_state |
| e000822 | 19:57:52.087774 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/5/nodes/3/flags |
| e000823 | 19:57:52.087774 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/3/primary_id |
| e000824 | 19:57:52.087774 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/5/nodes/3/slots |
| e000825 | 19:57:52.087774 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/3/link_state |
| e000826 | 19:57:52.087774 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/5/nodes/4/flags |
| e000827 | 19:57:52.087774 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/2/snapshot/views/5/nodes/4/primary_id |
| e000828 | 19:57:52.087774 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/4/slots |
| e000829 | 19:57:52.087774 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/4/link_state |
| e000830 | 19:57:52.087774 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/5/nodes/5/flags |
| e000831 | 19:57:52.087774 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/5/primary_id |
| e000832 | 19:57:52.087774 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/5/nodes/5/slots |
| e000833 | 19:57:52.087774 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/5/link_state |
| e000834 | 19:57:52.087774 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/5/persistence/aof_enabled |
| e000835 | 19:57:52.087774 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/5/persistence/loading |
| e000836 | 19:57:52.087774 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/5/persistence/aof_last_write_status |
| e000837 | 19:57:52.087774 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e000838 | 19:57:52.087774 | node-6 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/5/process_live |
| e000839 | 19:57:52.08782 | node-1 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/0/reachable |
| e000840 | 19:57:52.08782 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/2/snapshot/views/0/replication/role |
| e000841 | 19:57:52.08782 | node-1 | info_replication | replication.connected_slaves | valid | `0` | /samples/2/snapshot/views/0/replication/connected_slaves |
| e000842 | 19:57:52.08782 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/2/snapshot/views/0/replication/master_link_status |
| e000843 | 19:57:52.08782 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/2/snapshot/views/0/replication/master_sync_in_progress |
| e000844 | 19:57:52.08782 | node-1 | info_replication | replication.master_repl_offset | valid | `28419` | /samples/2/snapshot/views/0/replication/master_repl_offset |
| e000845 | 19:57:52.08782 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/2/snapshot/views/0/replication/slave_repl_offset |
| e000846 | 19:57:52.08782 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/2/snapshot/views/0/redis_id |
| e000847 | 19:57:52.08782 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/2/snapshot/views/0/coverage |
| e000848 | 19:57:52.08782 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave",
                  "fail?"
                ]` | /samples/2/snapshot/views/0/nodes/0/flags |
| e000849 | 19:57:52.08782 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/2/snapshot/views/0/nodes/0/primary_id |
| e000850 | 19:57:52.08782 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/0/slots |
| e000851 | 19:57:52.08782 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/0/link_state |
| e000852 | 19:57:52.08782 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/0/nodes/1/flags |
| e000853 | 19:57:52.08782 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/2/snapshot/views/0/nodes/1/primary_id |
| e000854 | 19:57:52.08782 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/1/slots |
| e000855 | 19:57:52.08782 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/1/link_state |
| e000856 | 19:57:52.08782 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/0/nodes/2/flags |
| e000857 | 19:57:52.08782 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/2/primary_id |
| e000858 | 19:57:52.08782 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/0/nodes/2/slots |
| e000859 | 19:57:52.08782 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/2/link_state |
| e000860 | 19:57:52.08782 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/2/snapshot/views/0/nodes/3/flags |
| e000861 | 19:57:52.08782 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/3/primary_id |
| e000862 | 19:57:52.08782 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/0/nodes/3/slots |
| e000863 | 19:57:52.08782 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/3/link_state |
| e000864 | 19:57:52.08782 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/0/nodes/4/flags |
| e000865 | 19:57:52.08782 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/2/snapshot/views/0/nodes/4/primary_id |
| e000866 | 19:57:52.08782 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/4/slots |
| e000867 | 19:57:52.08782 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/4/link_state |
| e000868 | 19:57:52.08782 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/0/nodes/5/flags |
| e000869 | 19:57:52.08782 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/5/primary_id |
| e000870 | 19:57:52.08782 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/0/nodes/5/slots |
| e000871 | 19:57:52.08782 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/5/link_state |
| e000872 | 19:57:52.08782 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/0/persistence/aof_enabled |
| e000873 | 19:57:52.08782 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/0/persistence/loading |
| e000874 | 19:57:52.08782 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/0/persistence/aof_last_write_status |
| e000875 | 19:57:52.08782 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e000876 | 19:57:52.08782 | node-1 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/0/process_live |
| e000877 | 19:57:52.218183 |  | client_probe | client_success | valid | `true` | /samples/2/probe/successful |
| e000878 | 19:57:52.381014 |  | client_window | client.attempts | valid | `170` | /metrics/windows/3/attempts |
| e000879 | 19:57:52.381014 |  | client_window | client.errors | valid | `28` | /metrics/windows/3/errors |
| e000880 | 19:57:52.381014 |  | client_window | client.successes | valid | `142` | /metrics/windows/3/successes |
| e000881 | 19:57:52.381014 |  | client_window | client.p95_ms | valid | `122.830654` | /metrics/windows/3/p95_ms |
| e000882 | 19:57:52.405679 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/2/write_policy_probe/policy |
| e000883 | 19:57:52.405679 |  | ack_probe | ack.successful | valid | `false` | /samples/2/write_policy_probe/successful |
| e000884 | 19:57:52.405679 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/2/write_policy_probe/requested_replicas |
| e000885 | 19:57:52.405679 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/2/write_policy_probe/set_acknowledged |
| e000886 | 19:57:52.405679 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          0
        ]` | /samples/2/write_policy_probe/acknowledgements |
| e000887 | 19:57:52.405679 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/2/write_policy_probe/acknowledgement_observed |
| e000888 | 19:57:52.405679 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/2/write_policy_probe/waitaof_requested_local |
| e000889 | 19:57:52.405679 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/2/write_policy_probe/waitaof_requested_replicas |
| e000890 | 19:57:52.405679 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/2/write_policy_probe/waitaof_local_acknowledgements |
| e000891 | 19:57:52.405679 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/2/write_policy_probe/waitaof_replica_acknowledgements |
| e000892 | 19:57:52.405679 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/2/write_policy_probe/waitaof_acknowledgement_observed |
| e000893 | 19:57:52.751649 | node-3 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/2/reachable |
| e000894 | 19:57:52.751649 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/3/snapshot/views/2/replication/role |
| e000895 | 19:57:52.751649 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/3/snapshot/views/2/replication/connected_slaves |
| e000896 | 19:57:52.751649 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/3/snapshot/views/2/replication/master_link_status |
| e000897 | 19:57:52.751649 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/3/snapshot/views/2/replication/master_sync_in_progress |
| e000898 | 19:57:52.751649 | node-3 | info_replication | replication.master_repl_offset | valid | `32796` | /samples/3/snapshot/views/2/replication/master_repl_offset |
| e000899 | 19:57:52.751649 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/3/snapshot/views/2/replication/slave_repl_offset |
| e000900 | 19:57:52.751649 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/3/snapshot/views/2/redis_id |
| e000901 | 19:57:52.751649 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/2/coverage |
| e000902 | 19:57:52.751649 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/2/nodes/0/flags |
| e000903 | 19:57:52.751649 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/3/snapshot/views/2/nodes/0/primary_id |
| e000904 | 19:57:52.751649 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/0/slots |
| e000905 | 19:57:52.751649 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/0/link_state |
| e000906 | 19:57:52.751649 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/3/snapshot/views/2/nodes/1/flags |
| e000907 | 19:57:52.751649 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/1/primary_id |
| e000908 | 19:57:52.751649 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/2/nodes/1/slots |
| e000909 | 19:57:52.751649 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/1/link_state |
| e000910 | 19:57:52.751649 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/2/nodes/2/flags |
| e000911 | 19:57:52.751649 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/3/snapshot/views/2/nodes/2/primary_id |
| e000912 | 19:57:52.751649 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/2/slots |
| e000913 | 19:57:52.751649 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/2/link_state |
| e000914 | 19:57:52.751649 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/2/nodes/3/flags |
| e000915 | 19:57:52.751649 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/3/primary_id |
| e000916 | 19:57:52.751649 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/2/nodes/3/slots |
| e000917 | 19:57:52.751649 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/3/link_state |
| e000918 | 19:57:52.751649 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/2/nodes/4/flags |
| e000919 | 19:57:52.751649 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/4/primary_id |
| e000920 | 19:57:52.751649 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/2/nodes/4/slots |
| e000921 | 19:57:52.751649 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/4/link_state |
| e000922 | 19:57:52.751649 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/3/snapshot/views/2/nodes/5/flags |
| e000923 | 19:57:52.751649 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/3/snapshot/views/2/nodes/5/primary_id |
| e000924 | 19:57:52.751649 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/5/slots |
| e000925 | 19:57:52.751649 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/5/link_state |
| e000926 | 19:57:52.751649 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/2/persistence/aof_enabled |
| e000927 | 19:57:52.751649 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/2/persistence/loading |
| e000928 | 19:57:52.751649 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/2/persistence/aof_last_write_status |
| e000929 | 19:57:52.751649 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e000930 | 19:57:52.751649 | node-3 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/2/process_live |
| e000931 | 19:57:52.751653 | node-2 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/1/reachable |
| e000932 | 19:57:52.751653 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/3/snapshot/views/1/replication/role |
| e000933 | 19:57:52.751653 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/3/snapshot/views/1/replication/connected_slaves |
| e000934 | 19:57:52.751653 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/3/snapshot/views/1/replication/master_link_status |
| e000935 | 19:57:52.751653 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/3/snapshot/views/1/replication/master_sync_in_progress |
| e000936 | 19:57:52.751653 | node-2 | info_replication | replication.master_repl_offset | valid | `32796` | /samples/3/snapshot/views/1/replication/master_repl_offset |
| e000937 | 19:57:52.751653 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/3/snapshot/views/1/replication/slave_repl_offset |
| e000938 | 19:57:52.751653 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/3/snapshot/views/1/redis_id |
| e000939 | 19:57:52.751653 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/1/coverage |
| e000940 | 19:57:52.751653 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/1/nodes/0/flags |
| e000941 | 19:57:52.751653 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/0/primary_id |
| e000942 | 19:57:52.751653 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/1/nodes/0/slots |
| e000943 | 19:57:52.751653 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/0/link_state |
| e000944 | 19:57:52.751653 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/1/nodes/1/flags |
| e000945 | 19:57:52.751653 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/3/snapshot/views/1/nodes/1/primary_id |
| e000946 | 19:57:52.751653 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/1/slots |
| e000947 | 19:57:52.751653 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/1/link_state |
| e000948 | 19:57:52.751653 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/1/nodes/2/flags |
| e000949 | 19:57:52.751653 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/3/snapshot/views/1/nodes/2/primary_id |
| e000950 | 19:57:52.751653 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/2/slots |
| e000951 | 19:57:52.751653 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/2/link_state |
| e000952 | 19:57:52.751653 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/3/snapshot/views/1/nodes/3/flags |
| e000953 | 19:57:52.751653 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/3/snapshot/views/1/nodes/3/primary_id |
| e000954 | 19:57:52.751653 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/3/slots |
| e000955 | 19:57:52.751653 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/3/link_state |
| e000956 | 19:57:52.751653 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/3/snapshot/views/1/nodes/4/flags |
| e000957 | 19:57:52.751653 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/4/primary_id |
| e000958 | 19:57:52.751653 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/1/nodes/4/slots |
| e000959 | 19:57:52.751653 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/4/link_state |
| e000960 | 19:57:52.751653 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/1/nodes/5/flags |
| e000961 | 19:57:52.751653 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/5/primary_id |
| e000962 | 19:57:52.751653 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/1/nodes/5/slots |
| e000963 | 19:57:52.751653 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/5/link_state |
| e000964 | 19:57:52.751653 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/1/persistence/aof_enabled |
| e000965 | 19:57:52.751653 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/1/persistence/loading |
| e000966 | 19:57:52.751653 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/1/persistence/aof_last_write_status |
| e000967 | 19:57:52.751653 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e000968 | 19:57:52.751653 | node-2 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/1/process_live |
| e000969 | 19:57:52.751672 | node-5 | node_client | reachable | valid | `false` | /samples/3/snapshot/views/4/reachable |
| e000970 | 19:57:52.751672 | node-5 | info_replication | replication.role | collection_failed | `` | /samples/3/snapshot/views/4/replication/role |
| e000971 | 19:57:52.751672 | node-5 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/3/snapshot/views/4/replication/connected_slaves |
| e000972 | 19:57:52.751672 | node-5 | info_replication | replication.master_link_status | collection_failed | `` | /samples/3/snapshot/views/4/replication/master_link_status |
| e000973 | 19:57:52.751672 | node-5 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/3/snapshot/views/4/replication/master_sync_in_progress |
| e000974 | 19:57:52.751672 | node-5 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/3/snapshot/views/4/replication/master_repl_offset |
| e000975 | 19:57:52.751672 | node-5 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/3/snapshot/views/4/replication/slave_repl_offset |
| e000976 | 19:57:52.751672 | node-5 | cluster_nodes | redis_id | collection_failed | `` | /samples/3/snapshot/views/4/redis_id |
| e000977 | 19:57:52.751672 | node-5 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/3/snapshot/views/4/coverage |
| e000978 | 19:57:52.751672 | node-5 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/3/snapshot/views/4/persistence/aof_enabled |
| e000979 | 19:57:52.751672 | node-5 | info_persistence | persistence.loading | collection_failed | `` | /samples/3/snapshot/views/4/persistence/loading |
| e000980 | 19:57:52.751672 | node-5 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/3/snapshot/views/4/persistence/aof_last_write_status |
| e000981 | 19:57:52.751672 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/3/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e000982 | 19:57:52.751672 | node-5 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/4/process_live |
| e000983 | 19:57:52.751709 | node-6 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/5/reachable |
| e000984 | 19:57:52.751709 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/3/snapshot/views/5/replication/role |
| e000985 | 19:57:52.751709 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/3/snapshot/views/5/replication/connected_slaves |
| e000986 | 19:57:52.751709 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/3/snapshot/views/5/replication/master_link_status |
| e000987 | 19:57:52.751709 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/3/snapshot/views/5/replication/master_sync_in_progress |
| e000988 | 19:57:52.751709 | node-6 | info_replication | replication.master_repl_offset | valid | `32796` | /samples/3/snapshot/views/5/replication/master_repl_offset |
| e000989 | 19:57:52.751709 | node-6 | info_replication | replication.slave_repl_offset | valid | `32796` | /samples/3/snapshot/views/5/replication/slave_repl_offset |
| e000990 | 19:57:52.751709 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/3/snapshot/views/5/redis_id |
| e000991 | 19:57:52.751709 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/5/coverage |
| e000992 | 19:57:52.751709 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/3/snapshot/views/5/nodes/0/flags |
| e000993 | 19:57:52.751709 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/3/snapshot/views/5/nodes/0/primary_id |
| e000994 | 19:57:52.751709 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/0/slots |
| e000995 | 19:57:52.751709 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/0/link_state |
| e000996 | 19:57:52.751709 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/5/nodes/1/flags |
| e000997 | 19:57:52.751709 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/1/primary_id |
| e000998 | 19:57:52.751709 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/5/nodes/1/slots |
| e000999 | 19:57:52.751709 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/1/link_state |
| e001000 | 19:57:52.751709 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/3/snapshot/views/5/nodes/2/flags |
| e001001 | 19:57:52.751709 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/3/snapshot/views/5/nodes/2/primary_id |
| e001002 | 19:57:52.751709 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/2/slots |
| e001003 | 19:57:52.751709 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/2/link_state |
| e001004 | 19:57:52.751709 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/5/nodes/3/flags |
| e001005 | 19:57:52.751709 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/3/primary_id |
| e001006 | 19:57:52.751709 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/5/nodes/3/slots |
| e001007 | 19:57:52.751709 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/3/link_state |
| e001008 | 19:57:52.751709 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/5/nodes/4/flags |
| e001009 | 19:57:52.751709 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/3/snapshot/views/5/nodes/4/primary_id |
| e001010 | 19:57:52.751709 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/4/slots |
| e001011 | 19:57:52.751709 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/4/link_state |
| e001012 | 19:57:52.751709 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/5/nodes/5/flags |
| e001013 | 19:57:52.751709 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/5/primary_id |
| e001014 | 19:57:52.751709 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/5/nodes/5/slots |
| e001015 | 19:57:52.751709 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/5/link_state |
| e001016 | 19:57:52.751709 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/5/persistence/aof_enabled |
| e001017 | 19:57:52.751709 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/5/persistence/loading |
| e001018 | 19:57:52.751709 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/5/persistence/aof_last_write_status |
| e001019 | 19:57:52.751709 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e001020 | 19:57:52.751709 | node-6 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/5/process_live |
| e001021 | 19:57:52.751713 | node-4 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/3/reachable |
| e001022 | 19:57:52.751713 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/3/snapshot/views/3/replication/role |
| e001023 | 19:57:52.751713 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/3/snapshot/views/3/replication/connected_slaves |
| e001024 | 19:57:52.751713 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/3/snapshot/views/3/replication/master_link_status |
| e001025 | 19:57:52.751713 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/3/snapshot/views/3/replication/master_sync_in_progress |
| e001026 | 19:57:52.751713 | node-4 | info_replication | replication.master_repl_offset | valid | `32796` | /samples/3/snapshot/views/3/replication/master_repl_offset |
| e001027 | 19:57:52.751713 | node-4 | info_replication | replication.slave_repl_offset | valid | `32796` | /samples/3/snapshot/views/3/replication/slave_repl_offset |
| e001028 | 19:57:52.751713 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/3/snapshot/views/3/redis_id |
| e001029 | 19:57:52.751713 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/3/coverage |
| e001030 | 19:57:52.751713 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/3/nodes/0/flags |
| e001031 | 19:57:52.751713 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/0/primary_id |
| e001032 | 19:57:52.751713 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/3/nodes/0/slots |
| e001033 | 19:57:52.751713 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/0/link_state |
| e001034 | 19:57:52.751713 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/3/snapshot/views/3/nodes/1/flags |
| e001035 | 19:57:52.751713 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/3/snapshot/views/3/nodes/1/primary_id |
| e001036 | 19:57:52.751713 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/1/slots |
| e001037 | 19:57:52.751713 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/1/link_state |
| e001038 | 19:57:52.751713 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/3/nodes/2/flags |
| e001039 | 19:57:52.751713 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/3/snapshot/views/3/nodes/2/primary_id |
| e001040 | 19:57:52.751713 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/2/slots |
| e001041 | 19:57:52.751713 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/2/link_state |
| e001042 | 19:57:52.751713 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/3/snapshot/views/3/nodes/3/flags |
| e001043 | 19:57:52.751713 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/3/snapshot/views/3/nodes/3/primary_id |
| e001044 | 19:57:52.751713 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/3/slots |
| e001045 | 19:57:52.751713 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/3/link_state |
| e001046 | 19:57:52.751713 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/3/nodes/4/flags |
| e001047 | 19:57:52.751713 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/4/primary_id |
| e001048 | 19:57:52.751713 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/3/nodes/4/slots |
| e001049 | 19:57:52.751713 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/4/link_state |
| e001050 | 19:57:52.751713 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/3/nodes/5/flags |
| e001051 | 19:57:52.751713 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/5/primary_id |
| e001052 | 19:57:52.751713 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/3/nodes/5/slots |
| e001053 | 19:57:52.751713 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/5/link_state |
| e001054 | 19:57:52.751713 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/3/persistence/aof_enabled |
| e001055 | 19:57:52.751713 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/3/persistence/loading |
| e001056 | 19:57:52.751713 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/3/persistence/aof_last_write_status |
| e001057 | 19:57:52.751713 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001058 | 19:57:52.751713 | node-4 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/3/process_live |
| e001059 | 19:57:52.751725 | node-1 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/0/reachable |
| e001060 | 19:57:52.751725 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/3/snapshot/views/0/replication/role |
| e001061 | 19:57:52.751725 | node-1 | info_replication | replication.connected_slaves | valid | `0` | /samples/3/snapshot/views/0/replication/connected_slaves |
| e001062 | 19:57:52.751725 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/3/snapshot/views/0/replication/master_link_status |
| e001063 | 19:57:52.751725 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/3/snapshot/views/0/replication/master_sync_in_progress |
| e001064 | 19:57:52.751725 | node-1 | info_replication | replication.master_repl_offset | valid | `33053` | /samples/3/snapshot/views/0/replication/master_repl_offset |
| e001065 | 19:57:52.751725 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/3/snapshot/views/0/replication/slave_repl_offset |
| e001066 | 19:57:52.751725 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/3/snapshot/views/0/redis_id |
| e001067 | 19:57:52.751725 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/0/coverage |
| e001068 | 19:57:52.751725 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/3/snapshot/views/0/nodes/0/flags |
| e001069 | 19:57:52.751725 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/3/snapshot/views/0/nodes/0/primary_id |
| e001070 | 19:57:52.751725 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/0/slots |
| e001071 | 19:57:52.751725 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/0/link_state |
| e001072 | 19:57:52.751725 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/0/nodes/1/flags |
| e001073 | 19:57:52.751725 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/3/snapshot/views/0/nodes/1/primary_id |
| e001074 | 19:57:52.751725 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/1/slots |
| e001075 | 19:57:52.751725 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/1/link_state |
| e001076 | 19:57:52.751725 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/0/nodes/2/flags |
| e001077 | 19:57:52.751725 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/2/primary_id |
| e001078 | 19:57:52.751725 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/0/nodes/2/slots |
| e001079 | 19:57:52.751725 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/2/link_state |
| e001080 | 19:57:52.751725 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/3/snapshot/views/0/nodes/3/flags |
| e001081 | 19:57:52.751725 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/3/primary_id |
| e001082 | 19:57:52.751725 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/0/nodes/3/slots |
| e001083 | 19:57:52.751725 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/3/link_state |
| e001084 | 19:57:52.751725 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/0/nodes/4/flags |
| e001085 | 19:57:52.751725 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/3/snapshot/views/0/nodes/4/primary_id |
| e001086 | 19:57:52.751725 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/4/slots |
| e001087 | 19:57:52.751725 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/4/link_state |
| e001088 | 19:57:52.751725 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/0/nodes/5/flags |
| e001089 | 19:57:52.751725 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/5/primary_id |
| e001090 | 19:57:52.751725 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/0/nodes/5/slots |
| e001091 | 19:57:52.751725 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/5/link_state |
| e001092 | 19:57:52.751725 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/0/persistence/aof_enabled |
| e001093 | 19:57:52.751725 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/0/persistence/loading |
| e001094 | 19:57:52.751725 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/0/persistence/aof_last_write_status |
| e001095 | 19:57:52.751725 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001096 | 19:57:52.751725 | node-1 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/0/process_live |
| e001097 | 19:57:52.876782 |  | client_probe | client_success | valid | `true` | /samples/3/probe/successful |
| e001098 | 19:57:53.063964 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/3/write_policy_probe/policy |
| e001099 | 19:57:53.063964 |  | ack_probe | ack.successful | valid | `false` | /samples/3/write_policy_probe/successful |
| e001100 | 19:57:53.063964 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/3/write_policy_probe/requested_replicas |
| e001101 | 19:57:53.063964 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/3/write_policy_probe/set_acknowledged |
| e001102 | 19:57:53.063964 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          0
        ]` | /samples/3/write_policy_probe/acknowledgements |
| e001103 | 19:57:53.063964 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/3/write_policy_probe/acknowledgement_observed |
| e001104 | 19:57:53.063964 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/3/write_policy_probe/waitaof_requested_local |
| e001105 | 19:57:53.063964 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/3/write_policy_probe/waitaof_requested_replicas |
| e001106 | 19:57:53.063964 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/3/write_policy_probe/waitaof_local_acknowledgements |
| e001107 | 19:57:53.063964 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/3/write_policy_probe/waitaof_replica_acknowledgements |
| e001108 | 19:57:53.063964 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/3/write_policy_probe/waitaof_acknowledgement_observed |
| e001109 | 19:57:53.381014 |  | client_window | client.attempts | valid | `171` | /metrics/windows/4/attempts |
| e001110 | 19:57:53.381014 |  | client_window | client.errors | valid | `26` | /metrics/windows/4/errors |
| e001111 | 19:57:53.381014 |  | client_window | client.successes | valid | `145` | /metrics/windows/4/successes |
| e001112 | 19:57:53.381014 |  | client_window | client.p95_ms | valid | `116.981575` | /metrics/windows/4/p95_ms |
| e001113 | 19:57:53.409479 | node-1 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/0/reachable |
| e001114 | 19:57:53.409479 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/4/snapshot/views/0/replication/role |
| e001115 | 19:57:53.409479 | node-1 | info_replication | replication.connected_slaves | valid | `0` | /samples/4/snapshot/views/0/replication/connected_slaves |
| e001116 | 19:57:53.409479 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/4/snapshot/views/0/replication/master_link_status |
| e001117 | 19:57:53.409479 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/4/snapshot/views/0/replication/master_sync_in_progress |
| e001118 | 19:57:53.409479 | node-1 | info_replication | replication.master_repl_offset | valid | `37946` | /samples/4/snapshot/views/0/replication/master_repl_offset |
| e001119 | 19:57:53.409479 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/4/snapshot/views/0/replication/slave_repl_offset |
| e001120 | 19:57:53.409479 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/4/snapshot/views/0/redis_id |
| e001121 | 19:57:53.409479 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/0/coverage |
| e001122 | 19:57:53.409479 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/4/snapshot/views/0/nodes/0/flags |
| e001123 | 19:57:53.409479 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/4/snapshot/views/0/nodes/0/primary_id |
| e001124 | 19:57:53.409479 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/0/slots |
| e001125 | 19:57:53.409479 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/0/link_state |
| e001126 | 19:57:53.409479 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/0/nodes/1/flags |
| e001127 | 19:57:53.409479 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/4/snapshot/views/0/nodes/1/primary_id |
| e001128 | 19:57:53.409479 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/1/slots |
| e001129 | 19:57:53.409479 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/1/link_state |
| e001130 | 19:57:53.409479 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/0/nodes/2/flags |
| e001131 | 19:57:53.409479 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/2/primary_id |
| e001132 | 19:57:53.409479 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/0/nodes/2/slots |
| e001133 | 19:57:53.409479 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/2/link_state |
| e001134 | 19:57:53.409479 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/4/snapshot/views/0/nodes/3/flags |
| e001135 | 19:57:53.409479 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/3/primary_id |
| e001136 | 19:57:53.409479 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/0/nodes/3/slots |
| e001137 | 19:57:53.409479 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/3/link_state |
| e001138 | 19:57:53.409479 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/0/nodes/4/flags |
| e001139 | 19:57:53.409479 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/4/snapshot/views/0/nodes/4/primary_id |
| e001140 | 19:57:53.409479 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/4/slots |
| e001141 | 19:57:53.409479 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/4/link_state |
| e001142 | 19:57:53.409479 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/0/nodes/5/flags |
| e001143 | 19:57:53.409479 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/5/primary_id |
| e001144 | 19:57:53.409479 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/0/nodes/5/slots |
| e001145 | 19:57:53.409479 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/5/link_state |
| e001146 | 19:57:53.409479 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/0/persistence/aof_enabled |
| e001147 | 19:57:53.409479 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/0/persistence/loading |
| e001148 | 19:57:53.409479 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/0/persistence/aof_last_write_status |
| e001149 | 19:57:53.409479 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001150 | 19:57:53.409479 | node-1 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/0/process_live |
| e001151 | 19:57:53.409483 | node-4 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/3/reachable |
| e001152 | 19:57:53.409483 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/4/snapshot/views/3/replication/role |
| e001153 | 19:57:53.409483 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/4/snapshot/views/3/replication/connected_slaves |
| e001154 | 19:57:53.409483 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/4/snapshot/views/3/replication/master_link_status |
| e001155 | 19:57:53.409483 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/4/snapshot/views/3/replication/master_sync_in_progress |
| e001156 | 19:57:53.409483 | node-4 | info_replication | replication.master_repl_offset | valid | `37689` | /samples/4/snapshot/views/3/replication/master_repl_offset |
| e001157 | 19:57:53.409483 | node-4 | info_replication | replication.slave_repl_offset | valid | `37689` | /samples/4/snapshot/views/3/replication/slave_repl_offset |
| e001158 | 19:57:53.409483 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/4/snapshot/views/3/redis_id |
| e001159 | 19:57:53.409483 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/3/coverage |
| e001160 | 19:57:53.409483 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/3/nodes/0/flags |
| e001161 | 19:57:53.409483 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/0/primary_id |
| e001162 | 19:57:53.409483 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/3/nodes/0/slots |
| e001163 | 19:57:53.409483 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/0/link_state |
| e001164 | 19:57:53.409483 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/4/snapshot/views/3/nodes/1/flags |
| e001165 | 19:57:53.409483 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/4/snapshot/views/3/nodes/1/primary_id |
| e001166 | 19:57:53.409483 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/1/slots |
| e001167 | 19:57:53.409483 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/1/link_state |
| e001168 | 19:57:53.409483 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/3/nodes/2/flags |
| e001169 | 19:57:53.409483 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/4/snapshot/views/3/nodes/2/primary_id |
| e001170 | 19:57:53.409483 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/2/slots |
| e001171 | 19:57:53.409483 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/2/link_state |
| e001172 | 19:57:53.409483 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/4/snapshot/views/3/nodes/3/flags |
| e001173 | 19:57:53.409483 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/4/snapshot/views/3/nodes/3/primary_id |
| e001174 | 19:57:53.409483 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/3/slots |
| e001175 | 19:57:53.409483 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/3/link_state |
| e001176 | 19:57:53.409483 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/3/nodes/4/flags |
| e001177 | 19:57:53.409483 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/4/primary_id |
| e001178 | 19:57:53.409483 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/3/nodes/4/slots |
| e001179 | 19:57:53.409483 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/4/link_state |
| e001180 | 19:57:53.409483 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/3/nodes/5/flags |
| e001181 | 19:57:53.409483 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/5/primary_id |
| e001182 | 19:57:53.409483 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/3/nodes/5/slots |
| e001183 | 19:57:53.409483 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/5/link_state |
| e001184 | 19:57:53.409483 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/3/persistence/aof_enabled |
| e001185 | 19:57:53.409483 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/3/persistence/loading |
| e001186 | 19:57:53.409483 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/3/persistence/aof_last_write_status |
| e001187 | 19:57:53.409483 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001188 | 19:57:53.409483 | node-4 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/3/process_live |
| e001189 | 19:57:53.409488 | node-3 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/2/reachable |
| e001190 | 19:57:53.409488 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/4/snapshot/views/2/replication/role |
| e001191 | 19:57:53.409488 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/4/snapshot/views/2/replication/connected_slaves |
| e001192 | 19:57:53.409488 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/4/snapshot/views/2/replication/master_link_status |
| e001193 | 19:57:53.409488 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/4/snapshot/views/2/replication/master_sync_in_progress |
| e001194 | 19:57:53.409488 | node-3 | info_replication | replication.master_repl_offset | valid | `37689` | /samples/4/snapshot/views/2/replication/master_repl_offset |
| e001195 | 19:57:53.409488 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/4/snapshot/views/2/replication/slave_repl_offset |
| e001196 | 19:57:53.409488 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/4/snapshot/views/2/redis_id |
| e001197 | 19:57:53.409488 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/2/coverage |
| e001198 | 19:57:53.409488 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/2/nodes/0/flags |
| e001199 | 19:57:53.409488 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/4/snapshot/views/2/nodes/0/primary_id |
| e001200 | 19:57:53.409488 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/0/slots |
| e001201 | 19:57:53.409488 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/0/link_state |
| e001202 | 19:57:53.409488 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/4/snapshot/views/2/nodes/1/flags |
| e001203 | 19:57:53.409488 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/1/primary_id |
| e001204 | 19:57:53.409488 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/2/nodes/1/slots |
| e001205 | 19:57:53.409488 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/1/link_state |
| e001206 | 19:57:53.409488 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/2/nodes/2/flags |
| e001207 | 19:57:53.409488 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/4/snapshot/views/2/nodes/2/primary_id |
| e001208 | 19:57:53.409488 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/2/slots |
| e001209 | 19:57:53.409488 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/2/link_state |
| e001210 | 19:57:53.409488 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/2/nodes/3/flags |
| e001211 | 19:57:53.409488 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/3/primary_id |
| e001212 | 19:57:53.409488 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/2/nodes/3/slots |
| e001213 | 19:57:53.409488 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/3/link_state |
| e001214 | 19:57:53.409488 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/2/nodes/4/flags |
| e001215 | 19:57:53.409488 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/4/primary_id |
| e001216 | 19:57:53.409488 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/2/nodes/4/slots |
| e001217 | 19:57:53.409488 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/4/link_state |
| e001218 | 19:57:53.409488 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/4/snapshot/views/2/nodes/5/flags |
| e001219 | 19:57:53.409488 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/4/snapshot/views/2/nodes/5/primary_id |
| e001220 | 19:57:53.409488 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/5/slots |
| e001221 | 19:57:53.409488 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/5/link_state |
| e001222 | 19:57:53.409488 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/2/persistence/aof_enabled |
| e001223 | 19:57:53.409488 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/2/persistence/loading |
| e001224 | 19:57:53.409488 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/2/persistence/aof_last_write_status |
| e001225 | 19:57:53.409488 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001226 | 19:57:53.409488 | node-3 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/2/process_live |
| e001227 | 19:57:53.409509 | node-2 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/1/reachable |
| e001228 | 19:57:53.409509 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/4/snapshot/views/1/replication/role |
| e001229 | 19:57:53.409509 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/4/snapshot/views/1/replication/connected_slaves |
| e001230 | 19:57:53.409509 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/4/snapshot/views/1/replication/master_link_status |
| e001231 | 19:57:53.409509 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/4/snapshot/views/1/replication/master_sync_in_progress |
| e001232 | 19:57:53.409509 | node-2 | info_replication | replication.master_repl_offset | valid | `37689` | /samples/4/snapshot/views/1/replication/master_repl_offset |
| e001233 | 19:57:53.409509 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/4/snapshot/views/1/replication/slave_repl_offset |
| e001234 | 19:57:53.409509 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/4/snapshot/views/1/redis_id |
| e001235 | 19:57:53.409509 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/1/coverage |
| e001236 | 19:57:53.409509 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/1/nodes/0/flags |
| e001237 | 19:57:53.409509 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/0/primary_id |
| e001238 | 19:57:53.409509 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/1/nodes/0/slots |
| e001239 | 19:57:53.409509 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/0/link_state |
| e001240 | 19:57:53.409509 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/1/nodes/1/flags |
| e001241 | 19:57:53.409509 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/4/snapshot/views/1/nodes/1/primary_id |
| e001242 | 19:57:53.409509 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/1/slots |
| e001243 | 19:57:53.409509 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/1/link_state |
| e001244 | 19:57:53.409509 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/1/nodes/2/flags |
| e001245 | 19:57:53.409509 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/4/snapshot/views/1/nodes/2/primary_id |
| e001246 | 19:57:53.409509 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/2/slots |
| e001247 | 19:57:53.409509 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/2/link_state |
| e001248 | 19:57:53.409509 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/4/snapshot/views/1/nodes/3/flags |
| e001249 | 19:57:53.409509 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/4/snapshot/views/1/nodes/3/primary_id |
| e001250 | 19:57:53.409509 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/3/slots |
| e001251 | 19:57:53.409509 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/3/link_state |
| e001252 | 19:57:53.409509 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/4/snapshot/views/1/nodes/4/flags |
| e001253 | 19:57:53.409509 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/4/primary_id |
| e001254 | 19:57:53.409509 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/1/nodes/4/slots |
| e001255 | 19:57:53.409509 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/4/link_state |
| e001256 | 19:57:53.409509 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/1/nodes/5/flags |
| e001257 | 19:57:53.409509 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/5/primary_id |
| e001258 | 19:57:53.409509 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/1/nodes/5/slots |
| e001259 | 19:57:53.409509 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/5/link_state |
| e001260 | 19:57:53.409509 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/1/persistence/aof_enabled |
| e001261 | 19:57:53.409509 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/1/persistence/loading |
| e001262 | 19:57:53.409509 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/1/persistence/aof_last_write_status |
| e001263 | 19:57:53.409509 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001264 | 19:57:53.409509 | node-2 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/1/process_live |
| e001265 | 19:57:53.409518 | node-5 | node_client | reachable | valid | `false` | /samples/4/snapshot/views/4/reachable |
| e001266 | 19:57:53.409518 | node-5 | info_replication | replication.role | collection_failed | `` | /samples/4/snapshot/views/4/replication/role |
| e001267 | 19:57:53.409518 | node-5 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/4/snapshot/views/4/replication/connected_slaves |
| e001268 | 19:57:53.409518 | node-5 | info_replication | replication.master_link_status | collection_failed | `` | /samples/4/snapshot/views/4/replication/master_link_status |
| e001269 | 19:57:53.409518 | node-5 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/4/snapshot/views/4/replication/master_sync_in_progress |
| e001270 | 19:57:53.409518 | node-5 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/4/snapshot/views/4/replication/master_repl_offset |
| e001271 | 19:57:53.409518 | node-5 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/4/snapshot/views/4/replication/slave_repl_offset |
| e001272 | 19:57:53.409518 | node-5 | cluster_nodes | redis_id | collection_failed | `` | /samples/4/snapshot/views/4/redis_id |
| e001273 | 19:57:53.409518 | node-5 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/4/snapshot/views/4/coverage |
| e001274 | 19:57:53.409518 | node-5 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/4/snapshot/views/4/persistence/aof_enabled |
| e001275 | 19:57:53.409518 | node-5 | info_persistence | persistence.loading | collection_failed | `` | /samples/4/snapshot/views/4/persistence/loading |
| e001276 | 19:57:53.409518 | node-5 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/4/snapshot/views/4/persistence/aof_last_write_status |
| e001277 | 19:57:53.409518 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/4/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001278 | 19:57:53.409518 | node-5 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/4/process_live |
| e001279 | 19:57:53.409525 | node-6 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/5/reachable |
| e001280 | 19:57:53.409525 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/4/snapshot/views/5/replication/role |
| e001281 | 19:57:53.409525 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/4/snapshot/views/5/replication/connected_slaves |
| e001282 | 19:57:53.409525 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/4/snapshot/views/5/replication/master_link_status |
| e001283 | 19:57:53.409525 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/4/snapshot/views/5/replication/master_sync_in_progress |
| e001284 | 19:57:53.409525 | node-6 | info_replication | replication.master_repl_offset | valid | `37689` | /samples/4/snapshot/views/5/replication/master_repl_offset |
| e001285 | 19:57:53.409525 | node-6 | info_replication | replication.slave_repl_offset | valid | `37689` | /samples/4/snapshot/views/5/replication/slave_repl_offset |
| e001286 | 19:57:53.409525 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/4/snapshot/views/5/redis_id |
| e001287 | 19:57:53.409525 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/5/coverage |
| e001288 | 19:57:53.409525 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/4/snapshot/views/5/nodes/0/flags |
| e001289 | 19:57:53.409525 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/4/snapshot/views/5/nodes/0/primary_id |
| e001290 | 19:57:53.409525 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/0/slots |
| e001291 | 19:57:53.409525 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/0/link_state |
| e001292 | 19:57:53.409525 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/5/nodes/1/flags |
| e001293 | 19:57:53.409525 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/1/primary_id |
| e001294 | 19:57:53.409525 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/5/nodes/1/slots |
| e001295 | 19:57:53.409525 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/1/link_state |
| e001296 | 19:57:53.409525 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/4/snapshot/views/5/nodes/2/flags |
| e001297 | 19:57:53.409525 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/4/snapshot/views/5/nodes/2/primary_id |
| e001298 | 19:57:53.409525 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/2/slots |
| e001299 | 19:57:53.409525 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/2/link_state |
| e001300 | 19:57:53.409525 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/5/nodes/3/flags |
| e001301 | 19:57:53.409525 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/3/primary_id |
| e001302 | 19:57:53.409525 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/5/nodes/3/slots |
| e001303 | 19:57:53.409525 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/3/link_state |
| e001304 | 19:57:53.409525 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/5/nodes/4/flags |
| e001305 | 19:57:53.409525 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/4/snapshot/views/5/nodes/4/primary_id |
| e001306 | 19:57:53.409525 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/4/slots |
| e001307 | 19:57:53.409525 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/4/link_state |
| e001308 | 19:57:53.409525 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/5/nodes/5/flags |
| e001309 | 19:57:53.409525 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/5/primary_id |
| e001310 | 19:57:53.409525 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/5/nodes/5/slots |
| e001311 | 19:57:53.409525 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/5/link_state |
| e001312 | 19:57:53.409525 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/5/persistence/aof_enabled |
| e001313 | 19:57:53.409525 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/5/persistence/loading |
| e001314 | 19:57:53.409525 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/5/persistence/aof_last_write_status |
| e001315 | 19:57:53.409525 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e001316 | 19:57:53.409525 | node-6 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/5/process_live |
| e001317 | 19:57:53.537723 |  | client_probe | client_success | valid | `true` | /samples/4/probe/successful |
| e001318 | 19:57:53.735253 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/4/write_policy_probe/policy |
| e001319 | 19:57:53.735253 |  | ack_probe | ack.successful | valid | `false` | /samples/4/write_policy_probe/successful |
| e001320 | 19:57:53.735253 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/4/write_policy_probe/requested_replicas |
| e001321 | 19:57:53.735253 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/4/write_policy_probe/set_acknowledged |
| e001322 | 19:57:53.735253 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          0
        ]` | /samples/4/write_policy_probe/acknowledgements |
| e001323 | 19:57:53.735253 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/4/write_policy_probe/acknowledgement_observed |
| e001324 | 19:57:53.735253 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/4/write_policy_probe/waitaof_requested_local |
| e001325 | 19:57:53.735253 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/4/write_policy_probe/waitaof_requested_replicas |
| e001326 | 19:57:53.735253 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/4/write_policy_probe/waitaof_local_acknowledgements |
| e001327 | 19:57:53.735253 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/4/write_policy_probe/waitaof_replica_acknowledgements |
| e001328 | 19:57:53.735253 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/4/write_policy_probe/waitaof_acknowledgement_observed |
| e001329 | 19:57:54.080919 | node-3 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/2/reachable |
| e001330 | 19:57:54.080919 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/5/snapshot/views/2/replication/role |
| e001331 | 19:57:54.080919 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/5/snapshot/views/2/replication/connected_slaves |
| e001332 | 19:57:54.080919 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/5/snapshot/views/2/replication/master_link_status |
| e001333 | 19:57:54.080919 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/5/snapshot/views/2/replication/master_sync_in_progress |
| e001334 | 19:57:54.080919 | node-3 | info_replication | replication.master_repl_offset | valid | `42582` | /samples/5/snapshot/views/2/replication/master_repl_offset |
| e001335 | 19:57:54.080919 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/5/snapshot/views/2/replication/slave_repl_offset |
| e001336 | 19:57:54.080919 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/5/snapshot/views/2/redis_id |
| e001337 | 19:57:54.080919 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/2/coverage |
| e001338 | 19:57:54.080919 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/2/nodes/0/flags |
| e001339 | 19:57:54.080919 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/5/snapshot/views/2/nodes/0/primary_id |
| e001340 | 19:57:54.080919 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/0/slots |
| e001341 | 19:57:54.080919 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/0/link_state |
| e001342 | 19:57:54.080919 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/5/snapshot/views/2/nodes/1/flags |
| e001343 | 19:57:54.080919 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/1/primary_id |
| e001344 | 19:57:54.080919 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/2/nodes/1/slots |
| e001345 | 19:57:54.080919 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/1/link_state |
| e001346 | 19:57:54.080919 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/2/nodes/2/flags |
| e001347 | 19:57:54.080919 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/5/snapshot/views/2/nodes/2/primary_id |
| e001348 | 19:57:54.080919 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/2/slots |
| e001349 | 19:57:54.080919 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/2/link_state |
| e001350 | 19:57:54.080919 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/2/nodes/3/flags |
| e001351 | 19:57:54.080919 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/3/primary_id |
| e001352 | 19:57:54.080919 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/2/nodes/3/slots |
| e001353 | 19:57:54.080919 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/3/link_state |
| e001354 | 19:57:54.080919 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/2/nodes/4/flags |
| e001355 | 19:57:54.080919 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/4/primary_id |
| e001356 | 19:57:54.080919 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/2/nodes/4/slots |
| e001357 | 19:57:54.080919 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/4/link_state |
| e001358 | 19:57:54.080919 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/5/snapshot/views/2/nodes/5/flags |
| e001359 | 19:57:54.080919 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/5/snapshot/views/2/nodes/5/primary_id |
| e001360 | 19:57:54.080919 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/5/slots |
| e001361 | 19:57:54.080919 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/5/link_state |
| e001362 | 19:57:54.080919 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/2/persistence/aof_enabled |
| e001363 | 19:57:54.080919 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/2/persistence/loading |
| e001364 | 19:57:54.080919 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/2/persistence/aof_last_write_status |
| e001365 | 19:57:54.080919 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001366 | 19:57:54.080919 | node-3 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/2/process_live |
| e001367 | 19:57:54.08092 | node-6 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/5/reachable |
| e001368 | 19:57:54.08092 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/5/snapshot/views/5/replication/role |
| e001369 | 19:57:54.08092 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/5/snapshot/views/5/replication/connected_slaves |
| e001370 | 19:57:54.08092 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/5/snapshot/views/5/replication/master_link_status |
| e001371 | 19:57:54.08092 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/5/snapshot/views/5/replication/master_sync_in_progress |
| e001372 | 19:57:54.08092 | node-6 | info_replication | replication.master_repl_offset | valid | `42582` | /samples/5/snapshot/views/5/replication/master_repl_offset |
| e001373 | 19:57:54.08092 | node-6 | info_replication | replication.slave_repl_offset | valid | `42582` | /samples/5/snapshot/views/5/replication/slave_repl_offset |
| e001374 | 19:57:54.08092 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/5/snapshot/views/5/redis_id |
| e001375 | 19:57:54.08092 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/5/coverage |
| e001376 | 19:57:54.08092 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/5/snapshot/views/5/nodes/0/flags |
| e001377 | 19:57:54.08092 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/5/snapshot/views/5/nodes/0/primary_id |
| e001378 | 19:57:54.08092 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/0/slots |
| e001379 | 19:57:54.08092 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/0/link_state |
| e001380 | 19:57:54.08092 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/5/nodes/1/flags |
| e001381 | 19:57:54.08092 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/1/primary_id |
| e001382 | 19:57:54.08092 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/5/nodes/1/slots |
| e001383 | 19:57:54.08092 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/1/link_state |
| e001384 | 19:57:54.08092 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/5/snapshot/views/5/nodes/2/flags |
| e001385 | 19:57:54.08092 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/5/snapshot/views/5/nodes/2/primary_id |
| e001386 | 19:57:54.08092 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/2/slots |
| e001387 | 19:57:54.08092 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"disconnected"` | /samples/5/snapshot/views/5/nodes/2/link_state |
| e001388 | 19:57:54.08092 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/5/nodes/3/flags |
| e001389 | 19:57:54.08092 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/3/primary_id |
| e001390 | 19:57:54.08092 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/5/nodes/3/slots |
| e001391 | 19:57:54.08092 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/3/link_state |
| e001392 | 19:57:54.08092 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/5/nodes/4/flags |
| e001393 | 19:57:54.08092 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/5/snapshot/views/5/nodes/4/primary_id |
| e001394 | 19:57:54.08092 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/4/slots |
| e001395 | 19:57:54.08092 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/4/link_state |
| e001396 | 19:57:54.08092 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/5/nodes/5/flags |
| e001397 | 19:57:54.08092 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/5/primary_id |
| e001398 | 19:57:54.08092 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/5/nodes/5/slots |
| e001399 | 19:57:54.08092 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/5/link_state |
| e001400 | 19:57:54.08092 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/5/persistence/aof_enabled |
| e001401 | 19:57:54.08092 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/5/persistence/loading |
| e001402 | 19:57:54.08092 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/5/persistence/aof_last_write_status |
| e001403 | 19:57:54.08092 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e001404 | 19:57:54.08092 | node-6 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/5/process_live |
| e001405 | 19:57:54.080938 | node-5 | node_client | reachable | valid | `false` | /samples/5/snapshot/views/4/reachable |
| e001406 | 19:57:54.080938 | node-5 | info_replication | replication.role | collection_failed | `` | /samples/5/snapshot/views/4/replication/role |
| e001407 | 19:57:54.080938 | node-5 | info_replication | replication.connected_slaves | collection_failed | `` | /samples/5/snapshot/views/4/replication/connected_slaves |
| e001408 | 19:57:54.080938 | node-5 | info_replication | replication.master_link_status | collection_failed | `` | /samples/5/snapshot/views/4/replication/master_link_status |
| e001409 | 19:57:54.080938 | node-5 | info_replication | replication.master_sync_in_progress | collection_failed | `` | /samples/5/snapshot/views/4/replication/master_sync_in_progress |
| e001410 | 19:57:54.080938 | node-5 | info_replication | replication.master_repl_offset | collection_failed | `` | /samples/5/snapshot/views/4/replication/master_repl_offset |
| e001411 | 19:57:54.080938 | node-5 | info_replication | replication.slave_repl_offset | collection_failed | `` | /samples/5/snapshot/views/4/replication/slave_repl_offset |
| e001412 | 19:57:54.080938 | node-5 | cluster_nodes | redis_id | collection_failed | `` | /samples/5/snapshot/views/4/redis_id |
| e001413 | 19:57:54.080938 | node-5 | cluster_nodes | slot_coverage | collection_failed | `` | /samples/5/snapshot/views/4/coverage |
| e001414 | 19:57:54.080938 | node-5 | info_persistence | persistence.aof_enabled | collection_failed | `` | /samples/5/snapshot/views/4/persistence/aof_enabled |
| e001415 | 19:57:54.080938 | node-5 | info_persistence | persistence.loading | collection_failed | `` | /samples/5/snapshot/views/4/persistence/loading |
| e001416 | 19:57:54.080938 | node-5 | info_persistence | persistence.aof_last_write_status | collection_failed | `` | /samples/5/snapshot/views/4/persistence/aof_last_write_status |
| e001417 | 19:57:54.080938 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | collection_failed | `` | /samples/5/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001418 | 19:57:54.080938 | node-5 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/4/process_live |
| e001419 | 19:57:54.080944 | node-2 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/1/reachable |
| e001420 | 19:57:54.080944 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/5/snapshot/views/1/replication/role |
| e001421 | 19:57:54.080944 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/5/snapshot/views/1/replication/connected_slaves |
| e001422 | 19:57:54.080944 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/5/snapshot/views/1/replication/master_link_status |
| e001423 | 19:57:54.080944 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/5/snapshot/views/1/replication/master_sync_in_progress |
| e001424 | 19:57:54.080944 | node-2 | info_replication | replication.master_repl_offset | valid | `42582` | /samples/5/snapshot/views/1/replication/master_repl_offset |
| e001425 | 19:57:54.080944 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/5/snapshot/views/1/replication/slave_repl_offset |
| e001426 | 19:57:54.080944 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/5/snapshot/views/1/redis_id |
| e001427 | 19:57:54.080944 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/1/coverage |
| e001428 | 19:57:54.080944 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/1/nodes/0/flags |
| e001429 | 19:57:54.080944 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/0/primary_id |
| e001430 | 19:57:54.080944 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/1/nodes/0/slots |
| e001431 | 19:57:54.080944 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/0/link_state |
| e001432 | 19:57:54.080944 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/1/nodes/1/flags |
| e001433 | 19:57:54.080944 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/5/snapshot/views/1/nodes/1/primary_id |
| e001434 | 19:57:54.080944 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/1/slots |
| e001435 | 19:57:54.080944 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/1/link_state |
| e001436 | 19:57:54.080944 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/1/nodes/2/flags |
| e001437 | 19:57:54.080944 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/5/snapshot/views/1/nodes/2/primary_id |
| e001438 | 19:57:54.080944 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/2/slots |
| e001439 | 19:57:54.080944 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/2/link_state |
| e001440 | 19:57:54.080944 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/5/snapshot/views/1/nodes/3/flags |
| e001441 | 19:57:54.080944 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/5/snapshot/views/1/nodes/3/primary_id |
| e001442 | 19:57:54.080944 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/3/slots |
| e001443 | 19:57:54.080944 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/3/link_state |
| e001444 | 19:57:54.080944 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/5/snapshot/views/1/nodes/4/flags |
| e001445 | 19:57:54.080944 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/4/primary_id |
| e001446 | 19:57:54.080944 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/1/nodes/4/slots |
| e001447 | 19:57:54.080944 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/4/link_state |
| e001448 | 19:57:54.080944 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/1/nodes/5/flags |
| e001449 | 19:57:54.080944 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/5/primary_id |
| e001450 | 19:57:54.080944 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/1/nodes/5/slots |
| e001451 | 19:57:54.080944 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/5/link_state |
| e001452 | 19:57:54.080944 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/1/persistence/aof_enabled |
| e001453 | 19:57:54.080944 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/1/persistence/loading |
| e001454 | 19:57:54.080944 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/1/persistence/aof_last_write_status |
| e001455 | 19:57:54.080944 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001456 | 19:57:54.080944 | node-2 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/1/process_live |
| e001457 | 19:57:54.080957 | node-4 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/3/reachable |
| e001458 | 19:57:54.080957 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/5/snapshot/views/3/replication/role |
| e001459 | 19:57:54.080957 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/5/snapshot/views/3/replication/connected_slaves |
| e001460 | 19:57:54.080957 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/5/snapshot/views/3/replication/master_link_status |
| e001461 | 19:57:54.080957 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/5/snapshot/views/3/replication/master_sync_in_progress |
| e001462 | 19:57:54.080957 | node-4 | info_replication | replication.master_repl_offset | valid | `42582` | /samples/5/snapshot/views/3/replication/master_repl_offset |
| e001463 | 19:57:54.080957 | node-4 | info_replication | replication.slave_repl_offset | valid | `42582` | /samples/5/snapshot/views/3/replication/slave_repl_offset |
| e001464 | 19:57:54.080957 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/5/snapshot/views/3/redis_id |
| e001465 | 19:57:54.080957 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/3/coverage |
| e001466 | 19:57:54.080957 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/3/nodes/0/flags |
| e001467 | 19:57:54.080957 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/0/primary_id |
| e001468 | 19:57:54.080957 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/3/nodes/0/slots |
| e001469 | 19:57:54.080957 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/0/link_state |
| e001470 | 19:57:54.080957 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/5/snapshot/views/3/nodes/1/flags |
| e001471 | 19:57:54.080957 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/5/snapshot/views/3/nodes/1/primary_id |
| e001472 | 19:57:54.080957 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/1/slots |
| e001473 | 19:57:54.080957 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/1/link_state |
| e001474 | 19:57:54.080957 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/3/nodes/2/flags |
| e001475 | 19:57:54.080957 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/5/snapshot/views/3/nodes/2/primary_id |
| e001476 | 19:57:54.080957 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/2/slots |
| e001477 | 19:57:54.080957 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/2/link_state |
| e001478 | 19:57:54.080957 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/5/snapshot/views/3/nodes/3/flags |
| e001479 | 19:57:54.080957 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/5/snapshot/views/3/nodes/3/primary_id |
| e001480 | 19:57:54.080957 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/3/slots |
| e001481 | 19:57:54.080957 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/3/link_state |
| e001482 | 19:57:54.080957 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/3/nodes/4/flags |
| e001483 | 19:57:54.080957 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/4/primary_id |
| e001484 | 19:57:54.080957 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/3/nodes/4/slots |
| e001485 | 19:57:54.080957 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/4/link_state |
| e001486 | 19:57:54.080957 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/3/nodes/5/flags |
| e001487 | 19:57:54.080957 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/5/primary_id |
| e001488 | 19:57:54.080957 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/3/nodes/5/slots |
| e001489 | 19:57:54.080957 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/5/link_state |
| e001490 | 19:57:54.080957 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/3/persistence/aof_enabled |
| e001491 | 19:57:54.080957 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/3/persistence/loading |
| e001492 | 19:57:54.080957 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/3/persistence/aof_last_write_status |
| e001493 | 19:57:54.080957 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001494 | 19:57:54.080957 | node-4 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/3/process_live |
| e001495 | 19:57:54.080967 | node-1 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/0/reachable |
| e001496 | 19:57:54.080967 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/5/snapshot/views/0/replication/role |
| e001497 | 19:57:54.080967 | node-1 | info_replication | replication.connected_slaves | valid | `0` | /samples/5/snapshot/views/0/replication/connected_slaves |
| e001498 | 19:57:54.080967 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/5/snapshot/views/0/replication/master_link_status |
| e001499 | 19:57:54.080967 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/5/snapshot/views/0/replication/master_sync_in_progress |
| e001500 | 19:57:54.080967 | node-1 | info_replication | replication.master_repl_offset | valid | `42839` | /samples/5/snapshot/views/0/replication/master_repl_offset |
| e001501 | 19:57:54.080967 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/5/snapshot/views/0/replication/slave_repl_offset |
| e001502 | 19:57:54.080967 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/5/snapshot/views/0/redis_id |
| e001503 | 19:57:54.080967 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/0/coverage |
| e001504 | 19:57:54.080967 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave",
                  "fail"
                ]` | /samples/5/snapshot/views/0/nodes/0/flags |
| e001505 | 19:57:54.080967 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/5/snapshot/views/0/nodes/0/primary_id |
| e001506 | 19:57:54.080967 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/0/slots |
| e001507 | 19:57:54.080967 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"disconnected"` | /samples/5/snapshot/views/0/nodes/0/link_state |
| e001508 | 19:57:54.080967 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/0/nodes/1/flags |
| e001509 | 19:57:54.080967 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/5/snapshot/views/0/nodes/1/primary_id |
| e001510 | 19:57:54.080967 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/1/slots |
| e001511 | 19:57:54.080967 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/0/nodes/1/link_state |
| e001512 | 19:57:54.080967 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/0/nodes/2/flags |
| e001513 | 19:57:54.080967 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/2/primary_id |
| e001514 | 19:57:54.080967 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/0/nodes/2/slots |
| e001515 | 19:57:54.080967 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/0/nodes/2/link_state |
| e001516 | 19:57:54.080967 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/5/snapshot/views/0/nodes/3/flags |
| e001517 | 19:57:54.080967 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/3/primary_id |
| e001518 | 19:57:54.080967 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/0/nodes/3/slots |
| e001519 | 19:57:54.080967 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/0/nodes/3/link_state |
| e001520 | 19:57:54.080967 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/0/nodes/4/flags |
| e001521 | 19:57:54.080967 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/5/snapshot/views/0/nodes/4/primary_id |
| e001522 | 19:57:54.080967 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/4/slots |
| e001523 | 19:57:54.080967 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/0/nodes/4/link_state |
| e001524 | 19:57:54.080967 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/0/nodes/5/flags |
| e001525 | 19:57:54.080967 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/5/primary_id |
| e001526 | 19:57:54.080967 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/0/nodes/5/slots |
| e001527 | 19:57:54.080967 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/0/nodes/5/link_state |
| e001528 | 19:57:54.080967 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/0/persistence/aof_enabled |
| e001529 | 19:57:54.080967 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/0/persistence/loading |
| e001530 | 19:57:54.080967 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/0/persistence/aof_last_write_status |
| e001531 | 19:57:54.080967 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001532 | 19:57:54.080967 | node-1 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/0/process_live |
| e001533 | 19:57:54.208373 |  | client_probe | client_success | valid | `true` | /samples/5/probe/successful |
| e001534 | 19:57:54.381014 |  | client_window | client.attempts | valid | `170` | /metrics/windows/5/attempts |
| e001535 | 19:57:54.381014 |  | client_window | client.errors | valid | `27` | /metrics/windows/5/errors |
| e001536 | 19:57:54.381014 |  | client_window | client.successes | valid | `143` | /metrics/windows/5/successes |
| e001537 | 19:57:54.381014 |  | client_window | client.p95_ms | valid | `116.981575` | /metrics/windows/5/p95_ms |
| e001538 | 19:57:54.411721 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/5/write_policy_probe/policy |
| e001539 | 19:57:54.411721 |  | ack_probe | ack.successful | valid | `false` | /samples/5/write_policy_probe/successful |
| e001540 | 19:57:54.411721 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/5/write_policy_probe/requested_replicas |
| e001541 | 19:57:54.411721 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/5/write_policy_probe/set_acknowledged |
| e001542 | 19:57:54.411721 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          0
        ]` | /samples/5/write_policy_probe/acknowledgements |
| e001543 | 19:57:54.411721 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/5/write_policy_probe/acknowledgement_observed |
| e001544 | 19:57:54.411721 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/5/write_policy_probe/waitaof_requested_local |
| e001545 | 19:57:54.411721 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/5/write_policy_probe/waitaof_requested_replicas |
| e001546 | 19:57:54.411721 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/5/write_policy_probe/waitaof_local_acknowledgements |
| e001547 | 19:57:54.411721 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/5/write_policy_probe/waitaof_replica_acknowledgements |
| e001548 | 19:57:54.411721 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/5/write_policy_probe/waitaof_acknowledgement_observed |
| e001549 | 19:57:55.381014 |  | client_window | client.attempts | valid | `169` | /metrics/windows/6/attempts |
| e001550 | 19:57:55.381014 |  | client_window | client.errors | valid | `26` | /metrics/windows/6/errors |
| e001551 | 19:57:55.381014 |  | client_window | client.successes | valid | `143` | /metrics/windows/6/successes |
| e001552 | 19:57:55.381014 |  | client_window | client.p95_ms | valid | `116.981575` | /metrics/windows/6/p95_ms |
| e001553 | 19:57:56.381014 |  | client_window | client.attempts | valid | `170` | /metrics/windows/7/attempts |
| e001554 | 19:57:56.381014 |  | client_window | client.errors | valid | `27` | /metrics/windows/7/errors |
| e001555 | 19:57:56.381014 |  | client_window | client.successes | valid | `143` | /metrics/windows/7/successes |
| e001556 | 19:57:56.381014 |  | client_window | client.p95_ms | valid | `122.830654` | /metrics/windows/7/p95_ms |
| e001557 | 19:57:56.439802 | node-3 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/2/reachable |
| e001558 | 19:57:56.439802 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/6/snapshot/views/2/replication/role |
| e001559 | 19:57:56.439802 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/6/snapshot/views/2/replication/connected_slaves |
| e001560 | 19:57:56.439802 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/6/snapshot/views/2/replication/master_link_status |
| e001561 | 19:57:56.439802 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/6/snapshot/views/2/replication/master_sync_in_progress |
| e001562 | 19:57:56.439802 | node-3 | info_replication | replication.master_repl_offset | valid | `59130` | /samples/6/snapshot/views/2/replication/master_repl_offset |
| e001563 | 19:57:56.439802 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/6/snapshot/views/2/replication/slave_repl_offset |
| e001564 | 19:57:56.439802 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/6/snapshot/views/2/redis_id |
| e001565 | 19:57:56.439802 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/2/coverage |
| e001566 | 19:57:56.439802 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/2/nodes/0/flags |
| e001567 | 19:57:56.439802 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/6/snapshot/views/2/nodes/0/primary_id |
| e001568 | 19:57:56.439802 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/0/slots |
| e001569 | 19:57:56.439802 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/0/link_state |
| e001570 | 19:57:56.439802 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/6/snapshot/views/2/nodes/1/flags |
| e001571 | 19:57:56.439802 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/1/primary_id |
| e001572 | 19:57:56.439802 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/2/nodes/1/slots |
| e001573 | 19:57:56.439802 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/1/link_state |
| e001574 | 19:57:56.439802 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/2/nodes/2/flags |
| e001575 | 19:57:56.439802 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/6/snapshot/views/2/nodes/2/primary_id |
| e001576 | 19:57:56.439802 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/2/slots |
| e001577 | 19:57:56.439802 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/2/link_state |
| e001578 | 19:57:56.439802 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/2/nodes/3/flags |
| e001579 | 19:57:56.439802 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/3/primary_id |
| e001580 | 19:57:56.439802 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/2/nodes/3/slots |
| e001581 | 19:57:56.439802 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/3/link_state |
| e001582 | 19:57:56.439802 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/2/nodes/4/flags |
| e001583 | 19:57:56.439802 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/4/primary_id |
| e001584 | 19:57:56.439802 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/2/nodes/4/slots |
| e001585 | 19:57:56.439802 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/4/link_state |
| e001586 | 19:57:56.439802 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/2/nodes/5/flags |
| e001587 | 19:57:56.439802 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/6/snapshot/views/2/nodes/5/primary_id |
| e001588 | 19:57:56.439802 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/5/slots |
| e001589 | 19:57:56.439802 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/5/link_state |
| e001590 | 19:57:56.439802 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/2/persistence/aof_enabled |
| e001591 | 19:57:56.439802 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/2/persistence/loading |
| e001592 | 19:57:56.439802 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/2/persistence/aof_last_write_status |
| e001593 | 19:57:56.439802 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001594 | 19:57:56.439802 | node-3 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/2/process_live |
| e001595 | 19:57:56.439819 | node-2 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/1/reachable |
| e001596 | 19:57:56.439819 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/6/snapshot/views/1/replication/role |
| e001597 | 19:57:56.439819 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/6/snapshot/views/1/replication/connected_slaves |
| e001598 | 19:57:56.439819 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/6/snapshot/views/1/replication/master_link_status |
| e001599 | 19:57:56.439819 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/6/snapshot/views/1/replication/master_sync_in_progress |
| e001600 | 19:57:56.439819 | node-2 | info_replication | replication.master_repl_offset | valid | `59130` | /samples/6/snapshot/views/1/replication/master_repl_offset |
| e001601 | 19:57:56.439819 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/6/snapshot/views/1/replication/slave_repl_offset |
| e001602 | 19:57:56.439819 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/6/snapshot/views/1/redis_id |
| e001603 | 19:57:56.439819 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/1/coverage |
| e001604 | 19:57:56.439819 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/1/nodes/0/flags |
| e001605 | 19:57:56.439819 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/0/primary_id |
| e001606 | 19:57:56.439819 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/1/nodes/0/slots |
| e001607 | 19:57:56.439819 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/0/link_state |
| e001608 | 19:57:56.439819 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/1/nodes/1/flags |
| e001609 | 19:57:56.439819 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/6/snapshot/views/1/nodes/1/primary_id |
| e001610 | 19:57:56.439819 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/1/slots |
| e001611 | 19:57:56.439819 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/1/link_state |
| e001612 | 19:57:56.439819 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/1/nodes/2/flags |
| e001613 | 19:57:56.439819 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/6/snapshot/views/1/nodes/2/primary_id |
| e001614 | 19:57:56.439819 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/2/slots |
| e001615 | 19:57:56.439819 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/2/link_state |
| e001616 | 19:57:56.439819 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/1/nodes/3/flags |
| e001617 | 19:57:56.439819 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/6/snapshot/views/1/nodes/3/primary_id |
| e001618 | 19:57:56.439819 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/3/slots |
| e001619 | 19:57:56.439819 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/3/link_state |
| e001620 | 19:57:56.439819 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/6/snapshot/views/1/nodes/4/flags |
| e001621 | 19:57:56.439819 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/4/primary_id |
| e001622 | 19:57:56.439819 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/1/nodes/4/slots |
| e001623 | 19:57:56.439819 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/4/link_state |
| e001624 | 19:57:56.439819 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/1/nodes/5/flags |
| e001625 | 19:57:56.439819 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/5/primary_id |
| e001626 | 19:57:56.439819 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/1/nodes/5/slots |
| e001627 | 19:57:56.439819 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/5/link_state |
| e001628 | 19:57:56.439819 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/1/persistence/aof_enabled |
| e001629 | 19:57:56.439819 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/1/persistence/loading |
| e001630 | 19:57:56.439819 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/1/persistence/aof_last_write_status |
| e001631 | 19:57:56.439819 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001632 | 19:57:56.439819 | node-2 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/1/process_live |
| e001633 | 19:57:56.439839 | node-1 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/0/reachable |
| e001634 | 19:57:56.439839 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/6/snapshot/views/0/replication/role |
| e001635 | 19:57:56.439839 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/6/snapshot/views/0/replication/connected_slaves |
| e001636 | 19:57:56.439839 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/6/snapshot/views/0/replication/master_link_status |
| e001637 | 19:57:56.439839 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/6/snapshot/views/0/replication/master_sync_in_progress |
| e001638 | 19:57:56.439839 | node-1 | info_replication | replication.master_repl_offset | valid | `59387` | /samples/6/snapshot/views/0/replication/master_repl_offset |
| e001639 | 19:57:56.439839 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/6/snapshot/views/0/replication/slave_repl_offset |
| e001640 | 19:57:56.439839 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/6/snapshot/views/0/redis_id |
| e001641 | 19:57:56.439839 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/0/coverage |
| e001642 | 19:57:56.439839 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/0/nodes/0/flags |
| e001643 | 19:57:56.439839 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/6/snapshot/views/0/nodes/0/primary_id |
| e001644 | 19:57:56.439839 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/0/slots |
| e001645 | 19:57:56.439839 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/0/link_state |
| e001646 | 19:57:56.439839 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/0/nodes/1/flags |
| e001647 | 19:57:56.439839 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/6/snapshot/views/0/nodes/1/primary_id |
| e001648 | 19:57:56.439839 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/1/slots |
| e001649 | 19:57:56.439839 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/1/link_state |
| e001650 | 19:57:56.439839 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/0/nodes/2/flags |
| e001651 | 19:57:56.439839 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/2/primary_id |
| e001652 | 19:57:56.439839 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/0/nodes/2/slots |
| e001653 | 19:57:56.439839 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/2/link_state |
| e001654 | 19:57:56.439839 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/6/snapshot/views/0/nodes/3/flags |
| e001655 | 19:57:56.439839 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/3/primary_id |
| e001656 | 19:57:56.439839 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/0/nodes/3/slots |
| e001657 | 19:57:56.439839 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/3/link_state |
| e001658 | 19:57:56.439839 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/0/nodes/4/flags |
| e001659 | 19:57:56.439839 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/6/snapshot/views/0/nodes/4/primary_id |
| e001660 | 19:57:56.439839 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/4/slots |
| e001661 | 19:57:56.439839 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/4/link_state |
| e001662 | 19:57:56.439839 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/0/nodes/5/flags |
| e001663 | 19:57:56.439839 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/5/primary_id |
| e001664 | 19:57:56.439839 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/0/nodes/5/slots |
| e001665 | 19:57:56.439839 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/5/link_state |
| e001666 | 19:57:56.439839 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/0/persistence/aof_enabled |
| e001667 | 19:57:56.439839 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/0/persistence/loading |
| e001668 | 19:57:56.439839 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/0/persistence/aof_last_write_status |
| e001669 | 19:57:56.439839 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001670 | 19:57:56.439839 | node-1 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/0/process_live |
| e001671 | 19:57:56.439851 | node-5 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/4/reachable |
| e001672 | 19:57:56.439851 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/6/snapshot/views/4/replication/role |
| e001673 | 19:57:56.439851 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/6/snapshot/views/4/replication/connected_slaves |
| e001674 | 19:57:56.439851 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/6/snapshot/views/4/replication/master_link_status |
| e001675 | 19:57:56.439851 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/6/snapshot/views/4/replication/master_sync_in_progress |
| e001676 | 19:57:56.439851 | node-5 | info_replication | replication.master_repl_offset | valid | `0` | /samples/6/snapshot/views/4/replication/master_repl_offset |
| e001677 | 19:57:56.439851 | node-5 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/6/snapshot/views/4/replication/slave_repl_offset |
| e001678 | 19:57:56.439851 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /samples/6/snapshot/views/4/redis_id |
| e001679 | 19:57:56.439851 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/4/coverage |
| e001680 | 19:57:56.439851 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/4/nodes/0/flags |
| e001681 | 19:57:56.439851 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/0/primary_id |
| e001682 | 19:57:56.439851 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/4/nodes/0/slots |
| e001683 | 19:57:56.439851 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/0/link_state |
| e001684 | 19:57:56.439851 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/4/nodes/1/flags |
| e001685 | 19:57:56.439851 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/6/snapshot/views/4/nodes/1/primary_id |
| e001686 | 19:57:56.439851 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/1/slots |
| e001687 | 19:57:56.439851 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/1/link_state |
| e001688 | 19:57:56.439851 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/6/snapshot/views/4/nodes/2/flags |
| e001689 | 19:57:56.439851 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/6/snapshot/views/4/nodes/2/primary_id |
| e001690 | 19:57:56.439851 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/2/slots |
| e001691 | 19:57:56.439851 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/2/link_state |
| e001692 | 19:57:56.439851 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/4/nodes/3/flags |
| e001693 | 19:57:56.439851 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/3/primary_id |
| e001694 | 19:57:56.439851 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/4/nodes/3/slots |
| e001695 | 19:57:56.439851 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/3/link_state |
| e001696 | 19:57:56.439851 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/4/nodes/4/flags |
| e001697 | 19:57:56.439851 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/6/snapshot/views/4/nodes/4/primary_id |
| e001698 | 19:57:56.439851 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/4/slots |
| e001699 | 19:57:56.439851 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/4/link_state |
| e001700 | 19:57:56.439851 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/4/nodes/5/flags |
| e001701 | 19:57:56.439851 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/5/primary_id |
| e001702 | 19:57:56.439851 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/4/nodes/5/slots |
| e001703 | 19:57:56.439851 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/5/link_state |
| e001704 | 19:57:56.439851 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/4/persistence/aof_enabled |
| e001705 | 19:57:56.439851 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/4/persistence/loading |
| e001706 | 19:57:56.439851 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/4/persistence/aof_last_write_status |
| e001707 | 19:57:56.439851 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001708 | 19:57:56.439851 | node-5 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/4/process_live |
| e001709 | 19:57:56.439853 | node-6 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/5/reachable |
| e001710 | 19:57:56.439853 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/6/snapshot/views/5/replication/role |
| e001711 | 19:57:56.439853 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/6/snapshot/views/5/replication/connected_slaves |
| e001712 | 19:57:56.439853 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/6/snapshot/views/5/replication/master_link_status |
| e001713 | 19:57:56.439853 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/6/snapshot/views/5/replication/master_sync_in_progress |
| e001714 | 19:57:56.439853 | node-6 | info_replication | replication.master_repl_offset | valid | `59130` | /samples/6/snapshot/views/5/replication/master_repl_offset |
| e001715 | 19:57:56.439853 | node-6 | info_replication | replication.slave_repl_offset | valid | `59130` | /samples/6/snapshot/views/5/replication/slave_repl_offset |
| e001716 | 19:57:56.439853 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/6/snapshot/views/5/redis_id |
| e001717 | 19:57:56.439853 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/5/coverage |
| e001718 | 19:57:56.439853 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/6/snapshot/views/5/nodes/0/flags |
| e001719 | 19:57:56.439853 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/6/snapshot/views/5/nodes/0/primary_id |
| e001720 | 19:57:56.439853 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/0/slots |
| e001721 | 19:57:56.439853 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/0/link_state |
| e001722 | 19:57:56.439853 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/5/nodes/1/flags |
| e001723 | 19:57:56.439853 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/1/primary_id |
| e001724 | 19:57:56.439853 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/5/nodes/1/slots |
| e001725 | 19:57:56.439853 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/1/link_state |
| e001726 | 19:57:56.439853 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/5/nodes/2/flags |
| e001727 | 19:57:56.439853 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/6/snapshot/views/5/nodes/2/primary_id |
| e001728 | 19:57:56.439853 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/2/slots |
| e001729 | 19:57:56.439853 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/2/link_state |
| e001730 | 19:57:56.439853 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/5/nodes/3/flags |
| e001731 | 19:57:56.439853 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/3/primary_id |
| e001732 | 19:57:56.439853 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/5/nodes/3/slots |
| e001733 | 19:57:56.439853 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/3/link_state |
| e001734 | 19:57:56.439853 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/5/nodes/4/flags |
| e001735 | 19:57:56.439853 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/6/snapshot/views/5/nodes/4/primary_id |
| e001736 | 19:57:56.439853 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/4/slots |
| e001737 | 19:57:56.439853 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/4/link_state |
| e001738 | 19:57:56.439853 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/5/nodes/5/flags |
| e001739 | 19:57:56.439853 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/5/primary_id |
| e001740 | 19:57:56.439853 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/5/nodes/5/slots |
| e001741 | 19:57:56.439853 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/5/link_state |
| e001742 | 19:57:56.439853 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/5/persistence/aof_enabled |
| e001743 | 19:57:56.439853 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/5/persistence/loading |
| e001744 | 19:57:56.439853 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/5/persistence/aof_last_write_status |
| e001745 | 19:57:56.439853 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e001746 | 19:57:56.439853 | node-6 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/5/process_live |
| e001747 | 19:57:56.439857 | node-4 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/3/reachable |
| e001748 | 19:57:56.439857 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/6/snapshot/views/3/replication/role |
| e001749 | 19:57:56.439857 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/6/snapshot/views/3/replication/connected_slaves |
| e001750 | 19:57:56.439857 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/6/snapshot/views/3/replication/master_link_status |
| e001751 | 19:57:56.439857 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/6/snapshot/views/3/replication/master_sync_in_progress |
| e001752 | 19:57:56.439857 | node-4 | info_replication | replication.master_repl_offset | valid | `59130` | /samples/6/snapshot/views/3/replication/master_repl_offset |
| e001753 | 19:57:56.439857 | node-4 | info_replication | replication.slave_repl_offset | valid | `59130` | /samples/6/snapshot/views/3/replication/slave_repl_offset |
| e001754 | 19:57:56.439857 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/6/snapshot/views/3/redis_id |
| e001755 | 19:57:56.439857 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/3/coverage |
| e001756 | 19:57:56.439857 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/3/nodes/0/flags |
| e001757 | 19:57:56.439857 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/0/primary_id |
| e001758 | 19:57:56.439857 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/3/nodes/0/slots |
| e001759 | 19:57:56.439857 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/0/link_state |
| e001760 | 19:57:56.439857 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/3/nodes/1/flags |
| e001761 | 19:57:56.439857 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/6/snapshot/views/3/nodes/1/primary_id |
| e001762 | 19:57:56.439857 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/1/slots |
| e001763 | 19:57:56.439857 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/1/link_state |
| e001764 | 19:57:56.439857 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/3/nodes/2/flags |
| e001765 | 19:57:56.439857 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/6/snapshot/views/3/nodes/2/primary_id |
| e001766 | 19:57:56.439857 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/2/slots |
| e001767 | 19:57:56.439857 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/2/link_state |
| e001768 | 19:57:56.439857 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/6/snapshot/views/3/nodes/3/flags |
| e001769 | 19:57:56.439857 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/6/snapshot/views/3/nodes/3/primary_id |
| e001770 | 19:57:56.439857 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/3/slots |
| e001771 | 19:57:56.439857 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/3/link_state |
| e001772 | 19:57:56.439857 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/3/nodes/4/flags |
| e001773 | 19:57:56.439857 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/4/primary_id |
| e001774 | 19:57:56.439857 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/3/nodes/4/slots |
| e001775 | 19:57:56.439857 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/4/link_state |
| e001776 | 19:57:56.439857 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/3/nodes/5/flags |
| e001777 | 19:57:56.439857 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/5/primary_id |
| e001778 | 19:57:56.439857 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/3/nodes/5/slots |
| e001779 | 19:57:56.439857 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/5/link_state |
| e001780 | 19:57:56.439857 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/3/persistence/aof_enabled |
| e001781 | 19:57:56.439857 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/3/persistence/loading |
| e001782 | 19:57:56.439857 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/3/persistence/aof_last_write_status |
| e001783 | 19:57:56.439857 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001784 | 19:57:56.439857 | node-4 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/3/process_live |
| e001785 | 19:57:56.58532 |  | client_probe | client_success | valid | `true` | /samples/6/probe/successful |
| e001786 | 19:57:56.700009 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/6/write_policy_probe/policy |
| e001787 | 19:57:56.700009 |  | ack_probe | ack.successful | valid | `false` | /samples/6/write_policy_probe/successful |
| e001788 | 19:57:56.700009 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/6/write_policy_probe/requested_replicas |
| e001789 | 19:57:56.700009 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/6/write_policy_probe/set_acknowledged |
| e001790 | 19:57:56.700009 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          0
        ]` | /samples/6/write_policy_probe/acknowledgements |
| e001791 | 19:57:56.700009 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/6/write_policy_probe/acknowledgement_observed |
| e001792 | 19:57:56.700009 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/6/write_policy_probe/waitaof_requested_local |
| e001793 | 19:57:56.700009 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/6/write_policy_probe/waitaof_requested_replicas |
| e001794 | 19:57:56.700009 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/6/write_policy_probe/waitaof_local_acknowledgements |
| e001795 | 19:57:56.700009 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/6/write_policy_probe/waitaof_replica_acknowledgements |
| e001796 | 19:57:56.700009 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/6/write_policy_probe/waitaof_acknowledgement_observed |
| e001797 | 19:57:57.211209 | node-3 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/2/reachable |
| e001798 | 19:57:57.211209 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/7/snapshot/views/2/replication/role |
| e001799 | 19:57:57.211209 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/7/snapshot/views/2/replication/connected_slaves |
| e001800 | 19:57:57.211209 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/7/snapshot/views/2/replication/master_link_status |
| e001801 | 19:57:57.211209 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/7/snapshot/views/2/replication/master_sync_in_progress |
| e001802 | 19:57:57.211209 | node-3 | info_replication | replication.master_repl_offset | valid | `65318` | /samples/7/snapshot/views/2/replication/master_repl_offset |
| e001803 | 19:57:57.211209 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/7/snapshot/views/2/replication/slave_repl_offset |
| e001804 | 19:57:57.211209 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/7/snapshot/views/2/redis_id |
| e001805 | 19:57:57.211209 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/2/coverage |
| e001806 | 19:57:57.211209 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/2/nodes/0/flags |
| e001807 | 19:57:57.211209 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/7/snapshot/views/2/nodes/0/primary_id |
| e001808 | 19:57:57.211209 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/0/slots |
| e001809 | 19:57:57.211209 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/0/link_state |
| e001810 | 19:57:57.211209 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/7/snapshot/views/2/nodes/1/flags |
| e001811 | 19:57:57.211209 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/1/primary_id |
| e001812 | 19:57:57.211209 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/2/nodes/1/slots |
| e001813 | 19:57:57.211209 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/1/link_state |
| e001814 | 19:57:57.211209 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/2/nodes/2/flags |
| e001815 | 19:57:57.211209 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/7/snapshot/views/2/nodes/2/primary_id |
| e001816 | 19:57:57.211209 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/2/slots |
| e001817 | 19:57:57.211209 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/2/link_state |
| e001818 | 19:57:57.211209 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/2/nodes/3/flags |
| e001819 | 19:57:57.211209 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/3/primary_id |
| e001820 | 19:57:57.211209 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/2/nodes/3/slots |
| e001821 | 19:57:57.211209 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/3/link_state |
| e001822 | 19:57:57.211209 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/2/nodes/4/flags |
| e001823 | 19:57:57.211209 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/4/primary_id |
| e001824 | 19:57:57.211209 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/2/nodes/4/slots |
| e001825 | 19:57:57.211209 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/4/link_state |
| e001826 | 19:57:57.211209 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/2/nodes/5/flags |
| e001827 | 19:57:57.211209 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/7/snapshot/views/2/nodes/5/primary_id |
| e001828 | 19:57:57.211209 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/5/slots |
| e001829 | 19:57:57.211209 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/5/link_state |
| e001830 | 19:57:57.211209 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/2/persistence/aof_enabled |
| e001831 | 19:57:57.211209 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/2/persistence/loading |
| e001832 | 19:57:57.211209 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/2/persistence/aof_last_write_status |
| e001833 | 19:57:57.211209 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001834 | 19:57:57.211209 | node-3 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/2/process_live |
| e001835 | 19:57:57.211213 | node-4 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/3/reachable |
| e001836 | 19:57:57.211213 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/7/snapshot/views/3/replication/role |
| e001837 | 19:57:57.211213 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/7/snapshot/views/3/replication/connected_slaves |
| e001838 | 19:57:57.211213 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/7/snapshot/views/3/replication/master_link_status |
| e001839 | 19:57:57.211213 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/7/snapshot/views/3/replication/master_sync_in_progress |
| e001840 | 19:57:57.211213 | node-4 | info_replication | replication.master_repl_offset | valid | `65318` | /samples/7/snapshot/views/3/replication/master_repl_offset |
| e001841 | 19:57:57.211213 | node-4 | info_replication | replication.slave_repl_offset | valid | `65318` | /samples/7/snapshot/views/3/replication/slave_repl_offset |
| e001842 | 19:57:57.211213 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/7/snapshot/views/3/redis_id |
| e001843 | 19:57:57.211213 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/3/coverage |
| e001844 | 19:57:57.211213 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/3/nodes/0/flags |
| e001845 | 19:57:57.211213 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/0/primary_id |
| e001846 | 19:57:57.211213 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/3/nodes/0/slots |
| e001847 | 19:57:57.211213 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/0/link_state |
| e001848 | 19:57:57.211213 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/3/nodes/1/flags |
| e001849 | 19:57:57.211213 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/7/snapshot/views/3/nodes/1/primary_id |
| e001850 | 19:57:57.211213 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/1/slots |
| e001851 | 19:57:57.211213 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/1/link_state |
| e001852 | 19:57:57.211213 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/3/nodes/2/flags |
| e001853 | 19:57:57.211213 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/7/snapshot/views/3/nodes/2/primary_id |
| e001854 | 19:57:57.211213 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/2/slots |
| e001855 | 19:57:57.211213 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/2/link_state |
| e001856 | 19:57:57.211213 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/7/snapshot/views/3/nodes/3/flags |
| e001857 | 19:57:57.211213 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/7/snapshot/views/3/nodes/3/primary_id |
| e001858 | 19:57:57.211213 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/3/slots |
| e001859 | 19:57:57.211213 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/3/link_state |
| e001860 | 19:57:57.211213 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/3/nodes/4/flags |
| e001861 | 19:57:57.211213 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/4/primary_id |
| e001862 | 19:57:57.211213 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/3/nodes/4/slots |
| e001863 | 19:57:57.211213 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/4/link_state |
| e001864 | 19:57:57.211213 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/3/nodes/5/flags |
| e001865 | 19:57:57.211213 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/5/primary_id |
| e001866 | 19:57:57.211213 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/3/nodes/5/slots |
| e001867 | 19:57:57.211213 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/5/link_state |
| e001868 | 19:57:57.211213 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/3/persistence/aof_enabled |
| e001869 | 19:57:57.211213 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/3/persistence/loading |
| e001870 | 19:57:57.211213 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/3/persistence/aof_last_write_status |
| e001871 | 19:57:57.211213 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001872 | 19:57:57.211213 | node-4 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/3/process_live |
| e001873 | 19:57:57.211257 | node-1 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/0/reachable |
| e001874 | 19:57:57.211257 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/7/snapshot/views/0/replication/role |
| e001875 | 19:57:57.211257 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/7/snapshot/views/0/replication/connected_slaves |
| e001876 | 19:57:57.211257 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/7/snapshot/views/0/replication/master_link_status |
| e001877 | 19:57:57.211257 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/7/snapshot/views/0/replication/master_sync_in_progress |
| e001878 | 19:57:57.211257 | node-1 | info_replication | replication.master_repl_offset | valid | `65575` | /samples/7/snapshot/views/0/replication/master_repl_offset |
| e001879 | 19:57:57.211257 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/7/snapshot/views/0/replication/slave_repl_offset |
| e001880 | 19:57:57.211257 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/7/snapshot/views/0/redis_id |
| e001881 | 19:57:57.211257 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/0/coverage |
| e001882 | 19:57:57.211257 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/0/nodes/0/flags |
| e001883 | 19:57:57.211257 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/7/snapshot/views/0/nodes/0/primary_id |
| e001884 | 19:57:57.211257 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/0/slots |
| e001885 | 19:57:57.211257 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/0/link_state |
| e001886 | 19:57:57.211257 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/0/nodes/1/flags |
| e001887 | 19:57:57.211257 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/7/snapshot/views/0/nodes/1/primary_id |
| e001888 | 19:57:57.211257 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/1/slots |
| e001889 | 19:57:57.211257 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/1/link_state |
| e001890 | 19:57:57.211257 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/0/nodes/2/flags |
| e001891 | 19:57:57.211257 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/2/primary_id |
| e001892 | 19:57:57.211257 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/0/nodes/2/slots |
| e001893 | 19:57:57.211257 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/2/link_state |
| e001894 | 19:57:57.211257 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/7/snapshot/views/0/nodes/3/flags |
| e001895 | 19:57:57.211257 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/3/primary_id |
| e001896 | 19:57:57.211257 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/0/nodes/3/slots |
| e001897 | 19:57:57.211257 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/3/link_state |
| e001898 | 19:57:57.211257 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/0/nodes/4/flags |
| e001899 | 19:57:57.211257 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/7/snapshot/views/0/nodes/4/primary_id |
| e001900 | 19:57:57.211257 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/4/slots |
| e001901 | 19:57:57.211257 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/4/link_state |
| e001902 | 19:57:57.211257 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/0/nodes/5/flags |
| e001903 | 19:57:57.211257 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/5/primary_id |
| e001904 | 19:57:57.211257 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/0/nodes/5/slots |
| e001905 | 19:57:57.211257 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/5/link_state |
| e001906 | 19:57:57.211257 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/0/persistence/aof_enabled |
| e001907 | 19:57:57.211257 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/0/persistence/loading |
| e001908 | 19:57:57.211257 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/0/persistence/aof_last_write_status |
| e001909 | 19:57:57.211257 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001910 | 19:57:57.211257 | node-1 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/0/process_live |
| e001911 | 19:57:57.211269 | node-5 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/4/reachable |
| e001912 | 19:57:57.211269 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/7/snapshot/views/4/replication/role |
| e001913 | 19:57:57.211269 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/7/snapshot/views/4/replication/connected_slaves |
| e001914 | 19:57:57.211269 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/7/snapshot/views/4/replication/master_link_status |
| e001915 | 19:57:57.211269 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/7/snapshot/views/4/replication/master_sync_in_progress |
| e001916 | 19:57:57.211269 | node-5 | info_replication | replication.master_repl_offset | valid | `0` | /samples/7/snapshot/views/4/replication/master_repl_offset |
| e001917 | 19:57:57.211269 | node-5 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/7/snapshot/views/4/replication/slave_repl_offset |
| e001918 | 19:57:57.211269 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /samples/7/snapshot/views/4/redis_id |
| e001919 | 19:57:57.211269 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/4/coverage |
| e001920 | 19:57:57.211269 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/4/nodes/0/flags |
| e001921 | 19:57:57.211269 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/0/primary_id |
| e001922 | 19:57:57.211269 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/4/nodes/0/slots |
| e001923 | 19:57:57.211269 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/0/link_state |
| e001924 | 19:57:57.211269 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/4/nodes/1/flags |
| e001925 | 19:57:57.211269 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/7/snapshot/views/4/nodes/1/primary_id |
| e001926 | 19:57:57.211269 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/1/slots |
| e001927 | 19:57:57.211269 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/1/link_state |
| e001928 | 19:57:57.211269 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/7/snapshot/views/4/nodes/2/flags |
| e001929 | 19:57:57.211269 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/7/snapshot/views/4/nodes/2/primary_id |
| e001930 | 19:57:57.211269 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/2/slots |
| e001931 | 19:57:57.211269 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/2/link_state |
| e001932 | 19:57:57.211269 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/4/nodes/3/flags |
| e001933 | 19:57:57.211269 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/3/primary_id |
| e001934 | 19:57:57.211269 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/4/nodes/3/slots |
| e001935 | 19:57:57.211269 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/3/link_state |
| e001936 | 19:57:57.211269 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/4/nodes/4/flags |
| e001937 | 19:57:57.211269 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/7/snapshot/views/4/nodes/4/primary_id |
| e001938 | 19:57:57.211269 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/4/slots |
| e001939 | 19:57:57.211269 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/4/link_state |
| e001940 | 19:57:57.211269 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/4/nodes/5/flags |
| e001941 | 19:57:57.211269 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/5/primary_id |
| e001942 | 19:57:57.211269 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/4/nodes/5/slots |
| e001943 | 19:57:57.211269 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/5/link_state |
| e001944 | 19:57:57.211269 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/4/persistence/aof_enabled |
| e001945 | 19:57:57.211269 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/4/persistence/loading |
| e001946 | 19:57:57.211269 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/4/persistence/aof_last_write_status |
| e001947 | 19:57:57.211269 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001948 | 19:57:57.211269 | node-5 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/4/process_live |
| e001949 | 19:57:57.211277 | node-2 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/1/reachable |
| e001950 | 19:57:57.211277 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/7/snapshot/views/1/replication/role |
| e001951 | 19:57:57.211277 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/7/snapshot/views/1/replication/connected_slaves |
| e001952 | 19:57:57.211277 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/7/snapshot/views/1/replication/master_link_status |
| e001953 | 19:57:57.211277 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/7/snapshot/views/1/replication/master_sync_in_progress |
| e001954 | 19:57:57.211277 | node-2 | info_replication | replication.master_repl_offset | valid | `65540` | /samples/7/snapshot/views/1/replication/master_repl_offset |
| e001955 | 19:57:57.211277 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/7/snapshot/views/1/replication/slave_repl_offset |
| e001956 | 19:57:57.211277 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/7/snapshot/views/1/redis_id |
| e001957 | 19:57:57.211277 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/1/coverage |
| e001958 | 19:57:57.211277 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/1/nodes/0/flags |
| e001959 | 19:57:57.211277 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/0/primary_id |
| e001960 | 19:57:57.211277 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/1/nodes/0/slots |
| e001961 | 19:57:57.211277 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/0/link_state |
| e001962 | 19:57:57.211277 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/1/nodes/1/flags |
| e001963 | 19:57:57.211277 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/7/snapshot/views/1/nodes/1/primary_id |
| e001964 | 19:57:57.211277 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/1/slots |
| e001965 | 19:57:57.211277 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/1/link_state |
| e001966 | 19:57:57.211277 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/1/nodes/2/flags |
| e001967 | 19:57:57.211277 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/7/snapshot/views/1/nodes/2/primary_id |
| e001968 | 19:57:57.211277 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/2/slots |
| e001969 | 19:57:57.211277 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/2/link_state |
| e001970 | 19:57:57.211277 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/1/nodes/3/flags |
| e001971 | 19:57:57.211277 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/7/snapshot/views/1/nodes/3/primary_id |
| e001972 | 19:57:57.211277 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/3/slots |
| e001973 | 19:57:57.211277 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/3/link_state |
| e001974 | 19:57:57.211277 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/7/snapshot/views/1/nodes/4/flags |
| e001975 | 19:57:57.211277 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/4/primary_id |
| e001976 | 19:57:57.211277 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/1/nodes/4/slots |
| e001977 | 19:57:57.211277 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/4/link_state |
| e001978 | 19:57:57.211277 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/1/nodes/5/flags |
| e001979 | 19:57:57.211277 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/5/primary_id |
| e001980 | 19:57:57.211277 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/1/nodes/5/slots |
| e001981 | 19:57:57.211277 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/5/link_state |
| e001982 | 19:57:57.211277 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/1/persistence/aof_enabled |
| e001983 | 19:57:57.211277 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/1/persistence/loading |
| e001984 | 19:57:57.211277 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/1/persistence/aof_last_write_status |
| e001985 | 19:57:57.211277 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001986 | 19:57:57.211277 | node-2 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/1/process_live |
| e001987 | 19:57:57.211284 | node-6 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/5/reachable |
| e001988 | 19:57:57.211284 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/7/snapshot/views/5/replication/role |
| e001989 | 19:57:57.211284 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/7/snapshot/views/5/replication/connected_slaves |
| e001990 | 19:57:57.211284 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/7/snapshot/views/5/replication/master_link_status |
| e001991 | 19:57:57.211284 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/7/snapshot/views/5/replication/master_sync_in_progress |
| e001992 | 19:57:57.211284 | node-6 | info_replication | replication.master_repl_offset | valid | `65318` | /samples/7/snapshot/views/5/replication/master_repl_offset |
| e001993 | 19:57:57.211284 | node-6 | info_replication | replication.slave_repl_offset | valid | `65318` | /samples/7/snapshot/views/5/replication/slave_repl_offset |
| e001994 | 19:57:57.211284 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/7/snapshot/views/5/redis_id |
| e001995 | 19:57:57.211284 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/5/coverage |
| e001996 | 19:57:57.211284 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/7/snapshot/views/5/nodes/0/flags |
| e001997 | 19:57:57.211284 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/7/snapshot/views/5/nodes/0/primary_id |
| e001998 | 19:57:57.211284 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/0/slots |
| e001999 | 19:57:57.211284 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/0/link_state |
| e002000 | 19:57:57.211284 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/5/nodes/1/flags |
| e002001 | 19:57:57.211284 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/1/primary_id |
| e002002 | 19:57:57.211284 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/5/nodes/1/slots |
| e002003 | 19:57:57.211284 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/1/link_state |
| e002004 | 19:57:57.211284 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/5/nodes/2/flags |
| e002005 | 19:57:57.211284 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/7/snapshot/views/5/nodes/2/primary_id |
| e002006 | 19:57:57.211284 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/2/slots |
| e002007 | 19:57:57.211284 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/2/link_state |
| e002008 | 19:57:57.211284 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/5/nodes/3/flags |
| e002009 | 19:57:57.211284 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/3/primary_id |
| e002010 | 19:57:57.211284 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/5/nodes/3/slots |
| e002011 | 19:57:57.211284 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/3/link_state |
| e002012 | 19:57:57.211284 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/5/nodes/4/flags |
| e002013 | 19:57:57.211284 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/7/snapshot/views/5/nodes/4/primary_id |
| e002014 | 19:57:57.211284 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/4/slots |
| e002015 | 19:57:57.211284 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/4/link_state |
| e002016 | 19:57:57.211284 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/5/nodes/5/flags |
| e002017 | 19:57:57.211284 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/5/primary_id |
| e002018 | 19:57:57.211284 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/5/nodes/5/slots |
| e002019 | 19:57:57.211284 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/5/link_state |
| e002020 | 19:57:57.211284 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/5/persistence/aof_enabled |
| e002021 | 19:57:57.211284 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/5/persistence/loading |
| e002022 | 19:57:57.211284 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/5/persistence/aof_last_write_status |
| e002023 | 19:57:57.211284 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002024 | 19:57:57.211284 | node-6 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/5/process_live |
| e002025 | 19:57:57.353828 |  | client_probe | client_success | valid | `true` | /samples/7/probe/successful |
| e002026 | 19:57:57.381014 |  | client_window | client.attempts | valid | `170` | /metrics/windows/8/attempts |
| e002027 | 19:57:57.381014 |  | client_window | client.errors | valid | `30` | /metrics/windows/8/errors |
| e002028 | 19:57:57.381014 |  | client_window | client.successes | valid | `140` | /metrics/windows/8/successes |
| e002029 | 19:57:57.381014 |  | client_window | client.p95_ms | valid | `116.981575` | /metrics/windows/8/p95_ms |
| e002030 | 19:57:57.470533 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/7/write_policy_probe/policy |
| e002031 | 19:57:57.470533 |  | ack_probe | ack.successful | valid | `false` | /samples/7/write_policy_probe/successful |
| e002032 | 19:57:57.470533 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/7/write_policy_probe/requested_replicas |
| e002033 | 19:57:57.470533 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/7/write_policy_probe/set_acknowledged |
| e002034 | 19:57:57.470533 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          0
        ]` | /samples/7/write_policy_probe/acknowledgements |
| e002035 | 19:57:57.470533 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/7/write_policy_probe/acknowledgement_observed |
| e002036 | 19:57:57.470533 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/7/write_policy_probe/waitaof_requested_local |
| e002037 | 19:57:57.470533 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/7/write_policy_probe/waitaof_requested_replicas |
| e002038 | 19:57:57.470533 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/7/write_policy_probe/waitaof_local_acknowledgements |
| e002039 | 19:57:57.470533 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/7/write_policy_probe/waitaof_replica_acknowledgements |
| e002040 | 19:57:57.470533 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/7/write_policy_probe/waitaof_acknowledgement_observed |
| e002041 | 19:57:57.982912 | node-2 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/1/reachable |
| e002042 | 19:57:57.982912 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/8/snapshot/views/1/replication/role |
| e002043 | 19:57:57.982912 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/8/snapshot/views/1/replication/connected_slaves |
| e002044 | 19:57:57.982912 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/8/snapshot/views/1/replication/master_link_status |
| e002045 | 19:57:57.982912 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/8/snapshot/views/1/replication/master_sync_in_progress |
| e002046 | 19:57:57.982912 | node-2 | info_replication | replication.master_repl_offset | valid | `71506` | /samples/8/snapshot/views/1/replication/master_repl_offset |
| e002047 | 19:57:57.982912 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/8/snapshot/views/1/replication/slave_repl_offset |
| e002048 | 19:57:57.982912 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/8/snapshot/views/1/redis_id |
| e002049 | 19:57:57.982912 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/1/coverage |
| e002050 | 19:57:57.982912 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/1/nodes/0/flags |
| e002051 | 19:57:57.982912 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/0/primary_id |
| e002052 | 19:57:57.982912 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/1/nodes/0/slots |
| e002053 | 19:57:57.982912 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/0/link_state |
| e002054 | 19:57:57.982912 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/1/nodes/1/flags |
| e002055 | 19:57:57.982912 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/8/snapshot/views/1/nodes/1/primary_id |
| e002056 | 19:57:57.982912 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/1/slots |
| e002057 | 19:57:57.982912 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/1/link_state |
| e002058 | 19:57:57.982912 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/1/nodes/2/flags |
| e002059 | 19:57:57.982912 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/8/snapshot/views/1/nodes/2/primary_id |
| e002060 | 19:57:57.982912 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/2/slots |
| e002061 | 19:57:57.982912 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/2/link_state |
| e002062 | 19:57:57.982912 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/1/nodes/3/flags |
| e002063 | 19:57:57.982912 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/8/snapshot/views/1/nodes/3/primary_id |
| e002064 | 19:57:57.982912 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/3/slots |
| e002065 | 19:57:57.982912 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/3/link_state |
| e002066 | 19:57:57.982912 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/8/snapshot/views/1/nodes/4/flags |
| e002067 | 19:57:57.982912 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/4/primary_id |
| e002068 | 19:57:57.982912 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/1/nodes/4/slots |
| e002069 | 19:57:57.982912 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/4/link_state |
| e002070 | 19:57:57.982912 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/1/nodes/5/flags |
| e002071 | 19:57:57.982912 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/5/primary_id |
| e002072 | 19:57:57.982912 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/1/nodes/5/slots |
| e002073 | 19:57:57.982912 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/5/link_state |
| e002074 | 19:57:57.982912 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/1/persistence/aof_enabled |
| e002075 | 19:57:57.982912 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/1/persistence/loading |
| e002076 | 19:57:57.982912 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/1/persistence/aof_last_write_status |
| e002077 | 19:57:57.982912 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e002078 | 19:57:57.982912 | node-2 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/1/process_live |
| e002079 | 19:57:57.982919 | node-1 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/0/reachable |
| e002080 | 19:57:57.982919 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/8/snapshot/views/0/replication/role |
| e002081 | 19:57:57.982919 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/8/snapshot/views/0/replication/connected_slaves |
| e002082 | 19:57:57.982919 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/8/snapshot/views/0/replication/master_link_status |
| e002083 | 19:57:57.982919 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/8/snapshot/views/0/replication/master_sync_in_progress |
| e002084 | 19:57:57.982919 | node-1 | info_replication | replication.master_repl_offset | valid | `71504` | /samples/8/snapshot/views/0/replication/master_repl_offset |
| e002085 | 19:57:57.982919 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/8/snapshot/views/0/replication/slave_repl_offset |
| e002086 | 19:57:57.982919 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/8/snapshot/views/0/redis_id |
| e002087 | 19:57:57.982919 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/0/coverage |
| e002088 | 19:57:57.982919 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/0/nodes/0/flags |
| e002089 | 19:57:57.982919 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/8/snapshot/views/0/nodes/0/primary_id |
| e002090 | 19:57:57.982919 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/0/slots |
| e002091 | 19:57:57.982919 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/0/link_state |
| e002092 | 19:57:57.982919 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/0/nodes/1/flags |
| e002093 | 19:57:57.982919 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/8/snapshot/views/0/nodes/1/primary_id |
| e002094 | 19:57:57.982919 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/1/slots |
| e002095 | 19:57:57.982919 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/1/link_state |
| e002096 | 19:57:57.982919 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/0/nodes/2/flags |
| e002097 | 19:57:57.982919 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/2/primary_id |
| e002098 | 19:57:57.982919 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/0/nodes/2/slots |
| e002099 | 19:57:57.982919 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/2/link_state |
| e002100 | 19:57:57.982919 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/8/snapshot/views/0/nodes/3/flags |
| e002101 | 19:57:57.982919 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/3/primary_id |
| e002102 | 19:57:57.982919 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/0/nodes/3/slots |
| e002103 | 19:57:57.982919 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/3/link_state |
| e002104 | 19:57:57.982919 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/0/nodes/4/flags |
| e002105 | 19:57:57.982919 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/8/snapshot/views/0/nodes/4/primary_id |
| e002106 | 19:57:57.982919 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/4/slots |
| e002107 | 19:57:57.982919 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/4/link_state |
| e002108 | 19:57:57.982919 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/0/nodes/5/flags |
| e002109 | 19:57:57.982919 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/5/primary_id |
| e002110 | 19:57:57.982919 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/0/nodes/5/slots |
| e002111 | 19:57:57.982919 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/5/link_state |
| e002112 | 19:57:57.982919 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/0/persistence/aof_enabled |
| e002113 | 19:57:57.982919 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/0/persistence/loading |
| e002114 | 19:57:57.982919 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/0/persistence/aof_last_write_status |
| e002115 | 19:57:57.982919 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e002116 | 19:57:57.982919 | node-1 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/0/process_live |
| e002117 | 19:57:57.982931 | node-4 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/3/reachable |
| e002118 | 19:57:57.982931 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/8/snapshot/views/3/replication/role |
| e002119 | 19:57:57.982931 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/8/snapshot/views/3/replication/connected_slaves |
| e002120 | 19:57:57.982931 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/8/snapshot/views/3/replication/master_link_status |
| e002121 | 19:57:57.982931 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/8/snapshot/views/3/replication/master_sync_in_progress |
| e002122 | 19:57:57.982931 | node-4 | info_replication | replication.master_repl_offset | valid | `71506` | /samples/8/snapshot/views/3/replication/master_repl_offset |
| e002123 | 19:57:57.982931 | node-4 | info_replication | replication.slave_repl_offset | valid | `71506` | /samples/8/snapshot/views/3/replication/slave_repl_offset |
| e002124 | 19:57:57.982931 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/8/snapshot/views/3/redis_id |
| e002125 | 19:57:57.982931 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/3/coverage |
| e002126 | 19:57:57.982931 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/3/nodes/0/flags |
| e002127 | 19:57:57.982931 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/0/primary_id |
| e002128 | 19:57:57.982931 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/3/nodes/0/slots |
| e002129 | 19:57:57.982931 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/0/link_state |
| e002130 | 19:57:57.982931 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/3/nodes/1/flags |
| e002131 | 19:57:57.982931 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/8/snapshot/views/3/nodes/1/primary_id |
| e002132 | 19:57:57.982931 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/1/slots |
| e002133 | 19:57:57.982931 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/1/link_state |
| e002134 | 19:57:57.982931 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/3/nodes/2/flags |
| e002135 | 19:57:57.982931 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/8/snapshot/views/3/nodes/2/primary_id |
| e002136 | 19:57:57.982931 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/2/slots |
| e002137 | 19:57:57.982931 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/2/link_state |
| e002138 | 19:57:57.982931 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/8/snapshot/views/3/nodes/3/flags |
| e002139 | 19:57:57.982931 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/8/snapshot/views/3/nodes/3/primary_id |
| e002140 | 19:57:57.982931 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/3/slots |
| e002141 | 19:57:57.982931 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/3/link_state |
| e002142 | 19:57:57.982931 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/3/nodes/4/flags |
| e002143 | 19:57:57.982931 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/4/primary_id |
| e002144 | 19:57:57.982931 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/3/nodes/4/slots |
| e002145 | 19:57:57.982931 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/4/link_state |
| e002146 | 19:57:57.982931 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/3/nodes/5/flags |
| e002147 | 19:57:57.982931 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/5/primary_id |
| e002148 | 19:57:57.982931 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/3/nodes/5/slots |
| e002149 | 19:57:57.982931 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/5/link_state |
| e002150 | 19:57:57.982931 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/3/persistence/aof_enabled |
| e002151 | 19:57:57.982931 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/3/persistence/loading |
| e002152 | 19:57:57.982931 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/3/persistence/aof_last_write_status |
| e002153 | 19:57:57.982931 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e002154 | 19:57:57.982931 | node-4 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/3/process_live |
| e002155 | 19:57:57.982968 | node-3 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/2/reachable |
| e002156 | 19:57:57.982968 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/8/snapshot/views/2/replication/role |
| e002157 | 19:57:57.982968 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/8/snapshot/views/2/replication/connected_slaves |
| e002158 | 19:57:57.982968 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/8/snapshot/views/2/replication/master_link_status |
| e002159 | 19:57:57.982968 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/8/snapshot/views/2/replication/master_sync_in_progress |
| e002160 | 19:57:57.982968 | node-3 | info_replication | replication.master_repl_offset | valid | `71506` | /samples/8/snapshot/views/2/replication/master_repl_offset |
| e002161 | 19:57:57.982968 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/8/snapshot/views/2/replication/slave_repl_offset |
| e002162 | 19:57:57.982968 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/8/snapshot/views/2/redis_id |
| e002163 | 19:57:57.982968 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/2/coverage |
| e002164 | 19:57:57.982968 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/2/nodes/0/flags |
| e002165 | 19:57:57.982968 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/8/snapshot/views/2/nodes/0/primary_id |
| e002166 | 19:57:57.982968 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/0/slots |
| e002167 | 19:57:57.982968 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/0/link_state |
| e002168 | 19:57:57.982968 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/8/snapshot/views/2/nodes/1/flags |
| e002169 | 19:57:57.982968 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/1/primary_id |
| e002170 | 19:57:57.982968 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/2/nodes/1/slots |
| e002171 | 19:57:57.982968 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/1/link_state |
| e002172 | 19:57:57.982968 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/2/nodes/2/flags |
| e002173 | 19:57:57.982968 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/8/snapshot/views/2/nodes/2/primary_id |
| e002174 | 19:57:57.982968 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/2/slots |
| e002175 | 19:57:57.982968 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/2/link_state |
| e002176 | 19:57:57.982968 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/2/nodes/3/flags |
| e002177 | 19:57:57.982968 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/3/primary_id |
| e002178 | 19:57:57.982968 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/2/nodes/3/slots |
| e002179 | 19:57:57.982968 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/3/link_state |
| e002180 | 19:57:57.982968 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/2/nodes/4/flags |
| e002181 | 19:57:57.982968 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/4/primary_id |
| e002182 | 19:57:57.982968 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/2/nodes/4/slots |
| e002183 | 19:57:57.982968 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/4/link_state |
| e002184 | 19:57:57.982968 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/2/nodes/5/flags |
| e002185 | 19:57:57.982968 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/8/snapshot/views/2/nodes/5/primary_id |
| e002186 | 19:57:57.982968 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/5/slots |
| e002187 | 19:57:57.982968 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/5/link_state |
| e002188 | 19:57:57.982968 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/2/persistence/aof_enabled |
| e002189 | 19:57:57.982968 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/2/persistence/loading |
| e002190 | 19:57:57.982968 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/2/persistence/aof_last_write_status |
| e002191 | 19:57:57.982968 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e002192 | 19:57:57.982968 | node-3 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/2/process_live |
| e002193 | 19:57:57.982976 | node-5 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/4/reachable |
| e002194 | 19:57:57.982976 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/8/snapshot/views/4/replication/role |
| e002195 | 19:57:57.982976 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/8/snapshot/views/4/replication/connected_slaves |
| e002196 | 19:57:57.982976 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/8/snapshot/views/4/replication/master_link_status |
| e002197 | 19:57:57.982976 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/8/snapshot/views/4/replication/master_sync_in_progress |
| e002198 | 19:57:57.982976 | node-5 | info_replication | replication.master_repl_offset | valid | `0` | /samples/8/snapshot/views/4/replication/master_repl_offset |
| e002199 | 19:57:57.982976 | node-5 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/8/snapshot/views/4/replication/slave_repl_offset |
| e002200 | 19:57:57.982976 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /samples/8/snapshot/views/4/redis_id |
| e002201 | 19:57:57.982976 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/4/coverage |
| e002202 | 19:57:57.982976 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/4/nodes/0/flags |
| e002203 | 19:57:57.982976 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/0/primary_id |
| e002204 | 19:57:57.982976 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/4/nodes/0/slots |
| e002205 | 19:57:57.982976 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/0/link_state |
| e002206 | 19:57:57.982976 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/4/nodes/1/flags |
| e002207 | 19:57:57.982976 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/8/snapshot/views/4/nodes/1/primary_id |
| e002208 | 19:57:57.982976 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/1/slots |
| e002209 | 19:57:57.982976 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/1/link_state |
| e002210 | 19:57:57.982976 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/8/snapshot/views/4/nodes/2/flags |
| e002211 | 19:57:57.982976 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/8/snapshot/views/4/nodes/2/primary_id |
| e002212 | 19:57:57.982976 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/2/slots |
| e002213 | 19:57:57.982976 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/2/link_state |
| e002214 | 19:57:57.982976 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/4/nodes/3/flags |
| e002215 | 19:57:57.982976 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/3/primary_id |
| e002216 | 19:57:57.982976 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/4/nodes/3/slots |
| e002217 | 19:57:57.982976 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/3/link_state |
| e002218 | 19:57:57.982976 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/4/nodes/4/flags |
| e002219 | 19:57:57.982976 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/8/snapshot/views/4/nodes/4/primary_id |
| e002220 | 19:57:57.982976 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/4/slots |
| e002221 | 19:57:57.982976 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/4/link_state |
| e002222 | 19:57:57.982976 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/4/nodes/5/flags |
| e002223 | 19:57:57.982976 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/5/primary_id |
| e002224 | 19:57:57.982976 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/4/nodes/5/slots |
| e002225 | 19:57:57.982976 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/5/link_state |
| e002226 | 19:57:57.982976 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/4/persistence/aof_enabled |
| e002227 | 19:57:57.982976 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/4/persistence/loading |
| e002228 | 19:57:57.982976 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/4/persistence/aof_last_write_status |
| e002229 | 19:57:57.982976 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e002230 | 19:57:57.982976 | node-5 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/4/process_live |
| e002231 | 19:57:57.982999 | node-6 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/5/reachable |
| e002232 | 19:57:57.982999 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/8/snapshot/views/5/replication/role |
| e002233 | 19:57:57.982999 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/8/snapshot/views/5/replication/connected_slaves |
| e002234 | 19:57:57.982999 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/8/snapshot/views/5/replication/master_link_status |
| e002235 | 19:57:57.982999 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/8/snapshot/views/5/replication/master_sync_in_progress |
| e002236 | 19:57:57.982999 | node-6 | info_replication | replication.master_repl_offset | valid | `71506` | /samples/8/snapshot/views/5/replication/master_repl_offset |
| e002237 | 19:57:57.982999 | node-6 | info_replication | replication.slave_repl_offset | valid | `71506` | /samples/8/snapshot/views/5/replication/slave_repl_offset |
| e002238 | 19:57:57.982999 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/8/snapshot/views/5/redis_id |
| e002239 | 19:57:57.982999 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/5/coverage |
| e002240 | 19:57:57.982999 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/8/snapshot/views/5/nodes/0/flags |
| e002241 | 19:57:57.982999 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/8/snapshot/views/5/nodes/0/primary_id |
| e002242 | 19:57:57.982999 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/0/slots |
| e002243 | 19:57:57.982999 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/0/link_state |
| e002244 | 19:57:57.982999 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/5/nodes/1/flags |
| e002245 | 19:57:57.982999 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/1/primary_id |
| e002246 | 19:57:57.982999 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/5/nodes/1/slots |
| e002247 | 19:57:57.982999 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/1/link_state |
| e002248 | 19:57:57.982999 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/5/nodes/2/flags |
| e002249 | 19:57:57.982999 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/8/snapshot/views/5/nodes/2/primary_id |
| e002250 | 19:57:57.982999 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/2/slots |
| e002251 | 19:57:57.982999 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/2/link_state |
| e002252 | 19:57:57.982999 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/5/nodes/3/flags |
| e002253 | 19:57:57.982999 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/3/primary_id |
| e002254 | 19:57:57.982999 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/5/nodes/3/slots |
| e002255 | 19:57:57.982999 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/3/link_state |
| e002256 | 19:57:57.982999 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/5/nodes/4/flags |
| e002257 | 19:57:57.982999 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/8/snapshot/views/5/nodes/4/primary_id |
| e002258 | 19:57:57.982999 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/4/slots |
| e002259 | 19:57:57.982999 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/4/link_state |
| e002260 | 19:57:57.982999 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/5/nodes/5/flags |
| e002261 | 19:57:57.982999 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/5/primary_id |
| e002262 | 19:57:57.982999 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/5/nodes/5/slots |
| e002263 | 19:57:57.982999 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/5/link_state |
| e002264 | 19:57:57.982999 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/5/persistence/aof_enabled |
| e002265 | 19:57:57.982999 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/5/persistence/loading |
| e002266 | 19:57:57.982999 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/5/persistence/aof_last_write_status |
| e002267 | 19:57:57.982999 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002268 | 19:57:57.982999 | node-6 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/5/process_live |
| e002269 | 19:57:58.128257 |  | client_probe | client_success | valid | `true` | /samples/8/probe/successful |
| e002270 | 19:57:58.246889 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/8/write_policy_probe/policy |
| e002271 | 19:57:58.246889 |  | ack_probe | ack.successful | valid | `false` | /samples/8/write_policy_probe/successful |
| e002272 | 19:57:58.246889 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/8/write_policy_probe/requested_replicas |
| e002273 | 19:57:58.246889 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/8/write_policy_probe/set_acknowledged |
| e002274 | 19:57:58.246889 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          0
        ]` | /samples/8/write_policy_probe/acknowledgements |
| e002275 | 19:57:58.246889 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/8/write_policy_probe/acknowledgement_observed |
| e002276 | 19:57:58.246889 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/8/write_policy_probe/waitaof_requested_local |
| e002277 | 19:57:58.246889 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/8/write_policy_probe/waitaof_requested_replicas |
| e002278 | 19:57:58.246889 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/8/write_policy_probe/waitaof_local_acknowledgements |
| e002279 | 19:57:58.246889 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/8/write_policy_probe/waitaof_replica_acknowledgements |
| e002280 | 19:57:58.246889 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/8/write_policy_probe/waitaof_acknowledgement_observed |
| e002281 | 19:57:58.381014 |  | client_window | client.attempts | valid | `170` | /metrics/windows/9/attempts |
| e002282 | 19:57:58.381014 |  | client_window | client.errors | valid | `29` | /metrics/windows/9/errors |
| e002283 | 19:57:58.381014 |  | client_window | client.successes | valid | `141` | /metrics/windows/9/successes |
| e002284 | 19:57:58.381014 |  | client_window | client.p95_ms | valid | `122.830654` | /metrics/windows/9/p95_ms |
| e002285 | 19:57:58.756693 | node-1 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/0/reachable |
| e002286 | 19:57:58.756693 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/9/snapshot/views/0/replication/role |
| e002287 | 19:57:58.756693 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/9/snapshot/views/0/replication/connected_slaves |
| e002288 | 19:57:58.756693 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/9/snapshot/views/0/replication/master_link_status |
| e002289 | 19:57:58.756693 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/9/snapshot/views/0/replication/master_sync_in_progress |
| e002290 | 19:57:58.756693 | node-1 | info_replication | replication.master_repl_offset | valid | `77951` | /samples/9/snapshot/views/0/replication/master_repl_offset |
| e002291 | 19:57:58.756693 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/9/snapshot/views/0/replication/slave_repl_offset |
| e002292 | 19:57:58.756693 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/9/snapshot/views/0/redis_id |
| e002293 | 19:57:58.756693 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/0/coverage |
| e002294 | 19:57:58.756693 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/0/nodes/0/flags |
| e002295 | 19:57:58.756693 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/9/snapshot/views/0/nodes/0/primary_id |
| e002296 | 19:57:58.756693 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/0/slots |
| e002297 | 19:57:58.756693 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/0/link_state |
| e002298 | 19:57:58.756693 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/0/nodes/1/flags |
| e002299 | 19:57:58.756693 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/9/snapshot/views/0/nodes/1/primary_id |
| e002300 | 19:57:58.756693 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/1/slots |
| e002301 | 19:57:58.756693 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/1/link_state |
| e002302 | 19:57:58.756693 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/0/nodes/2/flags |
| e002303 | 19:57:58.756693 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/2/primary_id |
| e002304 | 19:57:58.756693 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/0/nodes/2/slots |
| e002305 | 19:57:58.756693 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/2/link_state |
| e002306 | 19:57:58.756693 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/9/snapshot/views/0/nodes/3/flags |
| e002307 | 19:57:58.756693 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/3/primary_id |
| e002308 | 19:57:58.756693 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/0/nodes/3/slots |
| e002309 | 19:57:58.756693 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/3/link_state |
| e002310 | 19:57:58.756693 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/0/nodes/4/flags |
| e002311 | 19:57:58.756693 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/9/snapshot/views/0/nodes/4/primary_id |
| e002312 | 19:57:58.756693 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/4/slots |
| e002313 | 19:57:58.756693 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/4/link_state |
| e002314 | 19:57:58.756693 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/0/nodes/5/flags |
| e002315 | 19:57:58.756693 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/5/primary_id |
| e002316 | 19:57:58.756693 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/0/nodes/5/slots |
| e002317 | 19:57:58.756693 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/5/link_state |
| e002318 | 19:57:58.756693 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/0/persistence/aof_enabled |
| e002319 | 19:57:58.756693 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/0/persistence/loading |
| e002320 | 19:57:58.756693 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/0/persistence/aof_last_write_status |
| e002321 | 19:57:58.756693 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e002322 | 19:57:58.756693 | node-1 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/0/process_live |
| e002323 | 19:57:58.756704 | node-2 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/1/reachable |
| e002324 | 19:57:58.756704 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/9/snapshot/views/1/replication/role |
| e002325 | 19:57:58.756704 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/9/snapshot/views/1/replication/connected_slaves |
| e002326 | 19:57:58.756704 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/9/snapshot/views/1/replication/master_link_status |
| e002327 | 19:57:58.756704 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/9/snapshot/views/1/replication/master_sync_in_progress |
| e002328 | 19:57:58.756704 | node-2 | info_replication | replication.master_repl_offset | valid | `77694` | /samples/9/snapshot/views/1/replication/master_repl_offset |
| e002329 | 19:57:58.756704 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/9/snapshot/views/1/replication/slave_repl_offset |
| e002330 | 19:57:58.756704 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/9/snapshot/views/1/redis_id |
| e002331 | 19:57:58.756704 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/1/coverage |
| e002332 | 19:57:58.756704 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/1/nodes/0/flags |
| e002333 | 19:57:58.756704 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/0/primary_id |
| e002334 | 19:57:58.756704 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/1/nodes/0/slots |
| e002335 | 19:57:58.756704 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/0/link_state |
| e002336 | 19:57:58.756704 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/1/nodes/1/flags |
| e002337 | 19:57:58.756704 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/9/snapshot/views/1/nodes/1/primary_id |
| e002338 | 19:57:58.756704 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/1/slots |
| e002339 | 19:57:58.756704 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/1/link_state |
| e002340 | 19:57:58.756704 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/1/nodes/2/flags |
| e002341 | 19:57:58.756704 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/9/snapshot/views/1/nodes/2/primary_id |
| e002342 | 19:57:58.756704 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/2/slots |
| e002343 | 19:57:58.756704 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/2/link_state |
| e002344 | 19:57:58.756704 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/1/nodes/3/flags |
| e002345 | 19:57:58.756704 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/9/snapshot/views/1/nodes/3/primary_id |
| e002346 | 19:57:58.756704 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/3/slots |
| e002347 | 19:57:58.756704 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/3/link_state |
| e002348 | 19:57:58.756704 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/9/snapshot/views/1/nodes/4/flags |
| e002349 | 19:57:58.756704 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/4/primary_id |
| e002350 | 19:57:58.756704 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/1/nodes/4/slots |
| e002351 | 19:57:58.756704 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/4/link_state |
| e002352 | 19:57:58.756704 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/1/nodes/5/flags |
| e002353 | 19:57:58.756704 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/5/primary_id |
| e002354 | 19:57:58.756704 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/1/nodes/5/slots |
| e002355 | 19:57:58.756704 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/5/link_state |
| e002356 | 19:57:58.756704 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/1/persistence/aof_enabled |
| e002357 | 19:57:58.756704 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/1/persistence/loading |
| e002358 | 19:57:58.756704 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/1/persistence/aof_last_write_status |
| e002359 | 19:57:58.756704 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e002360 | 19:57:58.756704 | node-2 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/1/process_live |
| e002361 | 19:57:58.756712 | node-4 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/3/reachable |
| e002362 | 19:57:58.756712 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/9/snapshot/views/3/replication/role |
| e002363 | 19:57:58.756712 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/9/snapshot/views/3/replication/connected_slaves |
| e002364 | 19:57:58.756712 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/9/snapshot/views/3/replication/master_link_status |
| e002365 | 19:57:58.756712 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/9/snapshot/views/3/replication/master_sync_in_progress |
| e002366 | 19:57:58.756712 | node-4 | info_replication | replication.master_repl_offset | valid | `77694` | /samples/9/snapshot/views/3/replication/master_repl_offset |
| e002367 | 19:57:58.756712 | node-4 | info_replication | replication.slave_repl_offset | valid | `77694` | /samples/9/snapshot/views/3/replication/slave_repl_offset |
| e002368 | 19:57:58.756712 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/9/snapshot/views/3/redis_id |
| e002369 | 19:57:58.756712 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/3/coverage |
| e002370 | 19:57:58.756712 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/3/nodes/0/flags |
| e002371 | 19:57:58.756712 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/0/primary_id |
| e002372 | 19:57:58.756712 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/3/nodes/0/slots |
| e002373 | 19:57:58.756712 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/0/link_state |
| e002374 | 19:57:58.756712 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/3/nodes/1/flags |
| e002375 | 19:57:58.756712 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/9/snapshot/views/3/nodes/1/primary_id |
| e002376 | 19:57:58.756712 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/1/slots |
| e002377 | 19:57:58.756712 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/1/link_state |
| e002378 | 19:57:58.756712 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/3/nodes/2/flags |
| e002379 | 19:57:58.756712 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/9/snapshot/views/3/nodes/2/primary_id |
| e002380 | 19:57:58.756712 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/2/slots |
| e002381 | 19:57:58.756712 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/2/link_state |
| e002382 | 19:57:58.756712 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/9/snapshot/views/3/nodes/3/flags |
| e002383 | 19:57:58.756712 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/9/snapshot/views/3/nodes/3/primary_id |
| e002384 | 19:57:58.756712 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/3/slots |
| e002385 | 19:57:58.756712 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/3/link_state |
| e002386 | 19:57:58.756712 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/3/nodes/4/flags |
| e002387 | 19:57:58.756712 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/4/primary_id |
| e002388 | 19:57:58.756712 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/3/nodes/4/slots |
| e002389 | 19:57:58.756712 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/4/link_state |
| e002390 | 19:57:58.756712 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/3/nodes/5/flags |
| e002391 | 19:57:58.756712 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/5/primary_id |
| e002392 | 19:57:58.756712 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/3/nodes/5/slots |
| e002393 | 19:57:58.756712 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/5/link_state |
| e002394 | 19:57:58.756712 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/3/persistence/aof_enabled |
| e002395 | 19:57:58.756712 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/3/persistence/loading |
| e002396 | 19:57:58.756712 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/3/persistence/aof_last_write_status |
| e002397 | 19:57:58.756712 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e002398 | 19:57:58.756712 | node-4 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/3/process_live |
| e002399 | 19:57:58.756731 | node-5 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/4/reachable |
| e002400 | 19:57:58.756731 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/9/snapshot/views/4/replication/role |
| e002401 | 19:57:58.756731 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/9/snapshot/views/4/replication/connected_slaves |
| e002402 | 19:57:58.756731 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/9/snapshot/views/4/replication/master_link_status |
| e002403 | 19:57:58.756731 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/9/snapshot/views/4/replication/master_sync_in_progress |
| e002404 | 19:57:58.756731 | node-5 | info_replication | replication.master_repl_offset | valid | `0` | /samples/9/snapshot/views/4/replication/master_repl_offset |
| e002405 | 19:57:58.756731 | node-5 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/9/snapshot/views/4/replication/slave_repl_offset |
| e002406 | 19:57:58.756731 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /samples/9/snapshot/views/4/redis_id |
| e002407 | 19:57:58.756731 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/4/coverage |
| e002408 | 19:57:58.756731 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/4/nodes/0/flags |
| e002409 | 19:57:58.756731 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/0/primary_id |
| e002410 | 19:57:58.756731 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/4/nodes/0/slots |
| e002411 | 19:57:58.756731 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/0/link_state |
| e002412 | 19:57:58.756731 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/4/nodes/1/flags |
| e002413 | 19:57:58.756731 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/9/snapshot/views/4/nodes/1/primary_id |
| e002414 | 19:57:58.756731 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/1/slots |
| e002415 | 19:57:58.756731 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/1/link_state |
| e002416 | 19:57:58.756731 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/9/snapshot/views/4/nodes/2/flags |
| e002417 | 19:57:58.756731 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/9/snapshot/views/4/nodes/2/primary_id |
| e002418 | 19:57:58.756731 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/2/slots |
| e002419 | 19:57:58.756731 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/2/link_state |
| e002420 | 19:57:58.756731 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/4/nodes/3/flags |
| e002421 | 19:57:58.756731 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/3/primary_id |
| e002422 | 19:57:58.756731 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/4/nodes/3/slots |
| e002423 | 19:57:58.756731 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/3/link_state |
| e002424 | 19:57:58.756731 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/4/nodes/4/flags |
| e002425 | 19:57:58.756731 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/9/snapshot/views/4/nodes/4/primary_id |
| e002426 | 19:57:58.756731 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/4/slots |
| e002427 | 19:57:58.756731 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/4/link_state |
| e002428 | 19:57:58.756731 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/4/nodes/5/flags |
| e002429 | 19:57:58.756731 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/5/primary_id |
| e002430 | 19:57:58.756731 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/4/nodes/5/slots |
| e002431 | 19:57:58.756731 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/5/link_state |
| e002432 | 19:57:58.756731 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/4/persistence/aof_enabled |
| e002433 | 19:57:58.756731 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/4/persistence/loading |
| e002434 | 19:57:58.756731 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/4/persistence/aof_last_write_status |
| e002435 | 19:57:58.756731 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e002436 | 19:57:58.756731 | node-5 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/4/process_live |
| e002437 | 19:57:58.756736 | node-3 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/2/reachable |
| e002438 | 19:57:58.756736 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/9/snapshot/views/2/replication/role |
| e002439 | 19:57:58.756736 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/9/snapshot/views/2/replication/connected_slaves |
| e002440 | 19:57:58.756736 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/9/snapshot/views/2/replication/master_link_status |
| e002441 | 19:57:58.756736 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/9/snapshot/views/2/replication/master_sync_in_progress |
| e002442 | 19:57:58.756736 | node-3 | info_replication | replication.master_repl_offset | valid | `77694` | /samples/9/snapshot/views/2/replication/master_repl_offset |
| e002443 | 19:57:58.756736 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/9/snapshot/views/2/replication/slave_repl_offset |
| e002444 | 19:57:58.756736 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/9/snapshot/views/2/redis_id |
| e002445 | 19:57:58.756736 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/2/coverage |
| e002446 | 19:57:58.756736 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/2/nodes/0/flags |
| e002447 | 19:57:58.756736 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/9/snapshot/views/2/nodes/0/primary_id |
| e002448 | 19:57:58.756736 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/0/slots |
| e002449 | 19:57:58.756736 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/0/link_state |
| e002450 | 19:57:58.756736 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/9/snapshot/views/2/nodes/1/flags |
| e002451 | 19:57:58.756736 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/1/primary_id |
| e002452 | 19:57:58.756736 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/2/nodes/1/slots |
| e002453 | 19:57:58.756736 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/1/link_state |
| e002454 | 19:57:58.756736 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/2/nodes/2/flags |
| e002455 | 19:57:58.756736 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/9/snapshot/views/2/nodes/2/primary_id |
| e002456 | 19:57:58.756736 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/2/slots |
| e002457 | 19:57:58.756736 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/2/link_state |
| e002458 | 19:57:58.756736 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/2/nodes/3/flags |
| e002459 | 19:57:58.756736 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/3/primary_id |
| e002460 | 19:57:58.756736 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/2/nodes/3/slots |
| e002461 | 19:57:58.756736 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/3/link_state |
| e002462 | 19:57:58.756736 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/2/nodes/4/flags |
| e002463 | 19:57:58.756736 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/4/primary_id |
| e002464 | 19:57:58.756736 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/2/nodes/4/slots |
| e002465 | 19:57:58.756736 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/4/link_state |
| e002466 | 19:57:58.756736 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/2/nodes/5/flags |
| e002467 | 19:57:58.756736 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/9/snapshot/views/2/nodes/5/primary_id |
| e002468 | 19:57:58.756736 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/5/slots |
| e002469 | 19:57:58.756736 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/5/link_state |
| e002470 | 19:57:58.756736 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/2/persistence/aof_enabled |
| e002471 | 19:57:58.756736 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/2/persistence/loading |
| e002472 | 19:57:58.756736 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/2/persistence/aof_last_write_status |
| e002473 | 19:57:58.756736 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e002474 | 19:57:58.756736 | node-3 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/2/process_live |
| e002475 | 19:57:58.756737 | node-6 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/5/reachable |
| e002476 | 19:57:58.756737 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/9/snapshot/views/5/replication/role |
| e002477 | 19:57:58.756737 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/9/snapshot/views/5/replication/connected_slaves |
| e002478 | 19:57:58.756737 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/9/snapshot/views/5/replication/master_link_status |
| e002479 | 19:57:58.756737 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/9/snapshot/views/5/replication/master_sync_in_progress |
| e002480 | 19:57:58.756737 | node-6 | info_replication | replication.master_repl_offset | valid | `77694` | /samples/9/snapshot/views/5/replication/master_repl_offset |
| e002481 | 19:57:58.756737 | node-6 | info_replication | replication.slave_repl_offset | valid | `77694` | /samples/9/snapshot/views/5/replication/slave_repl_offset |
| e002482 | 19:57:58.756737 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/9/snapshot/views/5/redis_id |
| e002483 | 19:57:58.756737 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/5/coverage |
| e002484 | 19:57:58.756737 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/9/snapshot/views/5/nodes/0/flags |
| e002485 | 19:57:58.756737 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/9/snapshot/views/5/nodes/0/primary_id |
| e002486 | 19:57:58.756737 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/0/slots |
| e002487 | 19:57:58.756737 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/0/link_state |
| e002488 | 19:57:58.756737 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/5/nodes/1/flags |
| e002489 | 19:57:58.756737 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/1/primary_id |
| e002490 | 19:57:58.756737 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/5/nodes/1/slots |
| e002491 | 19:57:58.756737 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/1/link_state |
| e002492 | 19:57:58.756737 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/5/nodes/2/flags |
| e002493 | 19:57:58.756737 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/9/snapshot/views/5/nodes/2/primary_id |
| e002494 | 19:57:58.756737 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/2/slots |
| e002495 | 19:57:58.756737 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/2/link_state |
| e002496 | 19:57:58.756737 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/5/nodes/3/flags |
| e002497 | 19:57:58.756737 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/3/primary_id |
| e002498 | 19:57:58.756737 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/5/nodes/3/slots |
| e002499 | 19:57:58.756737 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/3/link_state |
| e002500 | 19:57:58.756737 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/5/nodes/4/flags |
| e002501 | 19:57:58.756737 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/9/snapshot/views/5/nodes/4/primary_id |
| e002502 | 19:57:58.756737 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/4/slots |
| e002503 | 19:57:58.756737 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/4/link_state |
| e002504 | 19:57:58.756737 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/5/nodes/5/flags |
| e002505 | 19:57:58.756737 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/5/primary_id |
| e002506 | 19:57:58.756737 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/5/nodes/5/slots |
| e002507 | 19:57:58.756737 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/5/link_state |
| e002508 | 19:57:58.756737 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/5/persistence/aof_enabled |
| e002509 | 19:57:58.756737 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/5/persistence/loading |
| e002510 | 19:57:58.756737 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/5/persistence/aof_last_write_status |
| e002511 | 19:57:58.756737 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002512 | 19:57:58.756737 | node-6 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/5/process_live |
| e002513 | 19:57:58.901382 |  | client_probe | client_success | valid | `true` | /samples/9/probe/successful |
| e002514 | 19:57:59.005167 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/9/write_policy_probe/policy |
| e002515 | 19:57:59.005167 |  | ack_probe | ack.successful | valid | `false` | /samples/9/write_policy_probe/successful |
| e002516 | 19:57:59.005167 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/9/write_policy_probe/requested_replicas |
| e002517 | 19:57:59.005167 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/9/write_policy_probe/set_acknowledged |
| e002518 | 19:57:59.005167 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          0
        ]` | /samples/9/write_policy_probe/acknowledgements |
| e002519 | 19:57:59.005167 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/9/write_policy_probe/acknowledgement_observed |
| e002520 | 19:57:59.005167 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/9/write_policy_probe/waitaof_requested_local |
| e002521 | 19:57:59.005167 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/9/write_policy_probe/waitaof_requested_replicas |
| e002522 | 19:57:59.005167 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/9/write_policy_probe/waitaof_local_acknowledgements |
| e002523 | 19:57:59.005167 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/9/write_policy_probe/waitaof_replica_acknowledgements |
| e002524 | 19:57:59.005167 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/9/write_policy_probe/waitaof_acknowledgement_observed |
| e002525 | 19:57:59.381014 |  | client_window | client.attempts | valid | `169` | /metrics/windows/10/attempts |
| e002526 | 19:57:59.381014 |  | client_window | client.errors | valid | `29` | /metrics/windows/10/errors |
| e002527 | 19:57:59.381014 |  | client_window | client.successes | valid | `140` | /metrics/windows/10/successes |
| e002528 | 19:57:59.381014 |  | client_window | client.p95_ms | valid | `122.830654` | /metrics/windows/10/p95_ms |
| e002529 | 19:57:59.518058 | node-2 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/1/reachable |
| e002530 | 19:57:59.518058 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/10/snapshot/views/1/replication/role |
| e002531 | 19:57:59.518058 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/10/snapshot/views/1/replication/connected_slaves |
| e002532 | 19:57:59.518058 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/10/snapshot/views/1/replication/master_link_status |
| e002533 | 19:57:59.518058 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/10/snapshot/views/1/replication/master_sync_in_progress |
| e002534 | 19:57:59.518058 | node-2 | info_replication | replication.master_repl_offset | valid | `83623` | /samples/10/snapshot/views/1/replication/master_repl_offset |
| e002535 | 19:57:59.518058 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/10/snapshot/views/1/replication/slave_repl_offset |
| e002536 | 19:57:59.518058 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/10/snapshot/views/1/redis_id |
| e002537 | 19:57:59.518058 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/1/coverage |
| e002538 | 19:57:59.518058 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/1/nodes/0/flags |
| e002539 | 19:57:59.518058 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/0/primary_id |
| e002540 | 19:57:59.518058 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/1/nodes/0/slots |
| e002541 | 19:57:59.518058 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/0/link_state |
| e002542 | 19:57:59.518058 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/1/nodes/1/flags |
| e002543 | 19:57:59.518058 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/10/snapshot/views/1/nodes/1/primary_id |
| e002544 | 19:57:59.518058 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/1/slots |
| e002545 | 19:57:59.518058 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/1/link_state |
| e002546 | 19:57:59.518058 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/1/nodes/2/flags |
| e002547 | 19:57:59.518058 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/10/snapshot/views/1/nodes/2/primary_id |
| e002548 | 19:57:59.518058 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/2/slots |
| e002549 | 19:57:59.518058 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/2/link_state |
| e002550 | 19:57:59.518058 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/1/nodes/3/flags |
| e002551 | 19:57:59.518058 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/10/snapshot/views/1/nodes/3/primary_id |
| e002552 | 19:57:59.518058 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/3/slots |
| e002553 | 19:57:59.518058 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/3/link_state |
| e002554 | 19:57:59.518058 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/10/snapshot/views/1/nodes/4/flags |
| e002555 | 19:57:59.518058 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/4/primary_id |
| e002556 | 19:57:59.518058 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/1/nodes/4/slots |
| e002557 | 19:57:59.518058 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/4/link_state |
| e002558 | 19:57:59.518058 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/1/nodes/5/flags |
| e002559 | 19:57:59.518058 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/5/primary_id |
| e002560 | 19:57:59.518058 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/1/nodes/5/slots |
| e002561 | 19:57:59.518058 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/5/link_state |
| e002562 | 19:57:59.518058 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/1/persistence/aof_enabled |
| e002563 | 19:57:59.518058 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/1/persistence/loading |
| e002564 | 19:57:59.518058 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/1/persistence/aof_last_write_status |
| e002565 | 19:57:59.518058 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e002566 | 19:57:59.518058 | node-2 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/1/process_live |
| e002567 | 19:57:59.518071 | node-4 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/3/reachable |
| e002568 | 19:57:59.518071 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/10/snapshot/views/3/replication/role |
| e002569 | 19:57:59.518071 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/10/snapshot/views/3/replication/connected_slaves |
| e002570 | 19:57:59.518071 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/10/snapshot/views/3/replication/master_link_status |
| e002571 | 19:57:59.518071 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/10/snapshot/views/3/replication/master_sync_in_progress |
| e002572 | 19:57:59.518071 | node-4 | info_replication | replication.master_repl_offset | valid | `83364` | /samples/10/snapshot/views/3/replication/master_repl_offset |
| e002573 | 19:57:59.518071 | node-4 | info_replication | replication.slave_repl_offset | valid | `83364` | /samples/10/snapshot/views/3/replication/slave_repl_offset |
| e002574 | 19:57:59.518071 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/10/snapshot/views/3/redis_id |
| e002575 | 19:57:59.518071 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/3/coverage |
| e002576 | 19:57:59.518071 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/3/nodes/0/flags |
| e002577 | 19:57:59.518071 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/0/primary_id |
| e002578 | 19:57:59.518071 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/3/nodes/0/slots |
| e002579 | 19:57:59.518071 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/0/link_state |
| e002580 | 19:57:59.518071 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/3/nodes/1/flags |
| e002581 | 19:57:59.518071 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/10/snapshot/views/3/nodes/1/primary_id |
| e002582 | 19:57:59.518071 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/1/slots |
| e002583 | 19:57:59.518071 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/1/link_state |
| e002584 | 19:57:59.518071 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/3/nodes/2/flags |
| e002585 | 19:57:59.518071 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/10/snapshot/views/3/nodes/2/primary_id |
| e002586 | 19:57:59.518071 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/2/slots |
| e002587 | 19:57:59.518071 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/2/link_state |
| e002588 | 19:57:59.518071 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/10/snapshot/views/3/nodes/3/flags |
| e002589 | 19:57:59.518071 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/10/snapshot/views/3/nodes/3/primary_id |
| e002590 | 19:57:59.518071 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/3/slots |
| e002591 | 19:57:59.518071 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/3/link_state |
| e002592 | 19:57:59.518071 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/3/nodes/4/flags |
| e002593 | 19:57:59.518071 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/4/primary_id |
| e002594 | 19:57:59.518071 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/3/nodes/4/slots |
| e002595 | 19:57:59.518071 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/4/link_state |
| e002596 | 19:57:59.518071 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/3/nodes/5/flags |
| e002597 | 19:57:59.518071 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/5/primary_id |
| e002598 | 19:57:59.518071 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/3/nodes/5/slots |
| e002599 | 19:57:59.518071 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/5/link_state |
| e002600 | 19:57:59.518071 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/3/persistence/aof_enabled |
| e002601 | 19:57:59.518071 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/3/persistence/loading |
| e002602 | 19:57:59.518071 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/3/persistence/aof_last_write_status |
| e002603 | 19:57:59.518071 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e002604 | 19:57:59.518071 | node-4 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/3/process_live |
| e002605 | 19:57:59.518082 | node-3 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/2/reachable |
| e002606 | 19:57:59.518082 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/10/snapshot/views/2/replication/role |
| e002607 | 19:57:59.518082 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/10/snapshot/views/2/replication/connected_slaves |
| e002608 | 19:57:59.518082 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/10/snapshot/views/2/replication/master_link_status |
| e002609 | 19:57:59.518082 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/10/snapshot/views/2/replication/master_sync_in_progress |
| e002610 | 19:57:59.518082 | node-3 | info_replication | replication.master_repl_offset | valid | `83364` | /samples/10/snapshot/views/2/replication/master_repl_offset |
| e002611 | 19:57:59.518082 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/10/snapshot/views/2/replication/slave_repl_offset |
| e002612 | 19:57:59.518082 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/10/snapshot/views/2/redis_id |
| e002613 | 19:57:59.518082 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/2/coverage |
| e002614 | 19:57:59.518082 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/2/nodes/0/flags |
| e002615 | 19:57:59.518082 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/10/snapshot/views/2/nodes/0/primary_id |
| e002616 | 19:57:59.518082 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/0/slots |
| e002617 | 19:57:59.518082 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/0/link_state |
| e002618 | 19:57:59.518082 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/10/snapshot/views/2/nodes/1/flags |
| e002619 | 19:57:59.518082 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/1/primary_id |
| e002620 | 19:57:59.518082 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/2/nodes/1/slots |
| e002621 | 19:57:59.518082 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/1/link_state |
| e002622 | 19:57:59.518082 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/2/nodes/2/flags |
| e002623 | 19:57:59.518082 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/10/snapshot/views/2/nodes/2/primary_id |
| e002624 | 19:57:59.518082 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/2/slots |
| e002625 | 19:57:59.518082 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/2/link_state |
| e002626 | 19:57:59.518082 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/2/nodes/3/flags |
| e002627 | 19:57:59.518082 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/3/primary_id |
| e002628 | 19:57:59.518082 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/2/nodes/3/slots |
| e002629 | 19:57:59.518082 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/3/link_state |
| e002630 | 19:57:59.518082 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/2/nodes/4/flags |
| e002631 | 19:57:59.518082 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/4/primary_id |
| e002632 | 19:57:59.518082 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/2/nodes/4/slots |
| e002633 | 19:57:59.518082 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/4/link_state |
| e002634 | 19:57:59.518082 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/2/nodes/5/flags |
| e002635 | 19:57:59.518082 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/10/snapshot/views/2/nodes/5/primary_id |
| e002636 | 19:57:59.518082 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/5/slots |
| e002637 | 19:57:59.518082 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/5/link_state |
| e002638 | 19:57:59.518082 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/2/persistence/aof_enabled |
| e002639 | 19:57:59.518082 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/2/persistence/loading |
| e002640 | 19:57:59.518082 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/2/persistence/aof_last_write_status |
| e002641 | 19:57:59.518082 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e002642 | 19:57:59.518082 | node-3 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/2/process_live |
| e002643 | 19:57:59.518092 | node-5 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/4/reachable |
| e002644 | 19:57:59.518092 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/10/snapshot/views/4/replication/role |
| e002645 | 19:57:59.518092 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/10/snapshot/views/4/replication/connected_slaves |
| e002646 | 19:57:59.518092 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/10/snapshot/views/4/replication/master_link_status |
| e002647 | 19:57:59.518092 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/10/snapshot/views/4/replication/master_sync_in_progress |
| e002648 | 19:57:59.518092 | node-5 | info_replication | replication.master_repl_offset | valid | `0` | /samples/10/snapshot/views/4/replication/master_repl_offset |
| e002649 | 19:57:59.518092 | node-5 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/10/snapshot/views/4/replication/slave_repl_offset |
| e002650 | 19:57:59.518092 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /samples/10/snapshot/views/4/redis_id |
| e002651 | 19:57:59.518092 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/4/coverage |
| e002652 | 19:57:59.518092 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/4/nodes/0/flags |
| e002653 | 19:57:59.518092 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/0/primary_id |
| e002654 | 19:57:59.518092 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/4/nodes/0/slots |
| e002655 | 19:57:59.518092 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/0/link_state |
| e002656 | 19:57:59.518092 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/4/nodes/1/flags |
| e002657 | 19:57:59.518092 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/10/snapshot/views/4/nodes/1/primary_id |
| e002658 | 19:57:59.518092 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/1/slots |
| e002659 | 19:57:59.518092 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/1/link_state |
| e002660 | 19:57:59.518092 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/10/snapshot/views/4/nodes/2/flags |
| e002661 | 19:57:59.518092 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/10/snapshot/views/4/nodes/2/primary_id |
| e002662 | 19:57:59.518092 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/2/slots |
| e002663 | 19:57:59.518092 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/2/link_state |
| e002664 | 19:57:59.518092 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/4/nodes/3/flags |
| e002665 | 19:57:59.518092 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/3/primary_id |
| e002666 | 19:57:59.518092 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/4/nodes/3/slots |
| e002667 | 19:57:59.518092 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/3/link_state |
| e002668 | 19:57:59.518092 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/4/nodes/4/flags |
| e002669 | 19:57:59.518092 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/10/snapshot/views/4/nodes/4/primary_id |
| e002670 | 19:57:59.518092 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/4/slots |
| e002671 | 19:57:59.518092 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/4/link_state |
| e002672 | 19:57:59.518092 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/4/nodes/5/flags |
| e002673 | 19:57:59.518092 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/5/primary_id |
| e002674 | 19:57:59.518092 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/4/nodes/5/slots |
| e002675 | 19:57:59.518092 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/5/link_state |
| e002676 | 19:57:59.518092 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/4/persistence/aof_enabled |
| e002677 | 19:57:59.518092 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/4/persistence/loading |
| e002678 | 19:57:59.518092 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/4/persistence/aof_last_write_status |
| e002679 | 19:57:59.518092 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e002680 | 19:57:59.518092 | node-5 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/4/process_live |
| e002681 | 19:57:59.518117 | node-1 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/0/reachable |
| e002682 | 19:57:59.518117 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/10/snapshot/views/0/replication/role |
| e002683 | 19:57:59.518117 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/10/snapshot/views/0/replication/connected_slaves |
| e002684 | 19:57:59.518117 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/10/snapshot/views/0/replication/master_link_status |
| e002685 | 19:57:59.518117 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/10/snapshot/views/0/replication/master_sync_in_progress |
| e002686 | 19:57:59.518117 | node-1 | info_replication | replication.master_repl_offset | valid | `83621` | /samples/10/snapshot/views/0/replication/master_repl_offset |
| e002687 | 19:57:59.518117 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/10/snapshot/views/0/replication/slave_repl_offset |
| e002688 | 19:57:59.518117 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/10/snapshot/views/0/redis_id |
| e002689 | 19:57:59.518117 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/0/coverage |
| e002690 | 19:57:59.518117 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/0/nodes/0/flags |
| e002691 | 19:57:59.518117 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/10/snapshot/views/0/nodes/0/primary_id |
| e002692 | 19:57:59.518117 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/0/slots |
| e002693 | 19:57:59.518117 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/0/link_state |
| e002694 | 19:57:59.518117 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/0/nodes/1/flags |
| e002695 | 19:57:59.518117 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/10/snapshot/views/0/nodes/1/primary_id |
| e002696 | 19:57:59.518117 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/1/slots |
| e002697 | 19:57:59.518117 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/1/link_state |
| e002698 | 19:57:59.518117 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/0/nodes/2/flags |
| e002699 | 19:57:59.518117 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/2/primary_id |
| e002700 | 19:57:59.518117 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/0/nodes/2/slots |
| e002701 | 19:57:59.518117 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/2/link_state |
| e002702 | 19:57:59.518117 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/10/snapshot/views/0/nodes/3/flags |
| e002703 | 19:57:59.518117 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/3/primary_id |
| e002704 | 19:57:59.518117 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/0/nodes/3/slots |
| e002705 | 19:57:59.518117 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/3/link_state |
| e002706 | 19:57:59.518117 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/0/nodes/4/flags |
| e002707 | 19:57:59.518117 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/10/snapshot/views/0/nodes/4/primary_id |
| e002708 | 19:57:59.518117 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/4/slots |
| e002709 | 19:57:59.518117 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/4/link_state |
| e002710 | 19:57:59.518117 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/0/nodes/5/flags |
| e002711 | 19:57:59.518117 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/5/primary_id |
| e002712 | 19:57:59.518117 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/0/nodes/5/slots |
| e002713 | 19:57:59.518117 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/5/link_state |
| e002714 | 19:57:59.518117 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/0/persistence/aof_enabled |
| e002715 | 19:57:59.518117 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/0/persistence/loading |
| e002716 | 19:57:59.518117 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/0/persistence/aof_last_write_status |
| e002717 | 19:57:59.518117 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e002718 | 19:57:59.518117 | node-1 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/0/process_live |
| e002719 | 19:57:59.51815 | node-6 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/5/reachable |
| e002720 | 19:57:59.51815 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/10/snapshot/views/5/replication/role |
| e002721 | 19:57:59.51815 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/10/snapshot/views/5/replication/connected_slaves |
| e002722 | 19:57:59.51815 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/10/snapshot/views/5/replication/master_link_status |
| e002723 | 19:57:59.51815 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/10/snapshot/views/5/replication/master_sync_in_progress |
| e002724 | 19:57:59.51815 | node-6 | info_replication | replication.master_repl_offset | valid | `83623` | /samples/10/snapshot/views/5/replication/master_repl_offset |
| e002725 | 19:57:59.51815 | node-6 | info_replication | replication.slave_repl_offset | valid | `83623` | /samples/10/snapshot/views/5/replication/slave_repl_offset |
| e002726 | 19:57:59.51815 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/10/snapshot/views/5/redis_id |
| e002727 | 19:57:59.51815 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/5/coverage |
| e002728 | 19:57:59.51815 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/10/snapshot/views/5/nodes/0/flags |
| e002729 | 19:57:59.51815 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/10/snapshot/views/5/nodes/0/primary_id |
| e002730 | 19:57:59.51815 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/0/slots |
| e002731 | 19:57:59.51815 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/0/link_state |
| e002732 | 19:57:59.51815 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/5/nodes/1/flags |
| e002733 | 19:57:59.51815 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/1/primary_id |
| e002734 | 19:57:59.51815 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/5/nodes/1/slots |
| e002735 | 19:57:59.51815 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/1/link_state |
| e002736 | 19:57:59.51815 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/5/nodes/2/flags |
| e002737 | 19:57:59.51815 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/10/snapshot/views/5/nodes/2/primary_id |
| e002738 | 19:57:59.51815 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/2/slots |
| e002739 | 19:57:59.51815 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/2/link_state |
| e002740 | 19:57:59.51815 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/5/nodes/3/flags |
| e002741 | 19:57:59.51815 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/3/primary_id |
| e002742 | 19:57:59.51815 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/5/nodes/3/slots |
| e002743 | 19:57:59.51815 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/3/link_state |
| e002744 | 19:57:59.51815 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/5/nodes/4/flags |
| e002745 | 19:57:59.51815 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/10/snapshot/views/5/nodes/4/primary_id |
| e002746 | 19:57:59.51815 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/4/slots |
| e002747 | 19:57:59.51815 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/4/link_state |
| e002748 | 19:57:59.51815 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/5/nodes/5/flags |
| e002749 | 19:57:59.51815 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/5/primary_id |
| e002750 | 19:57:59.51815 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/5/nodes/5/slots |
| e002751 | 19:57:59.51815 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/5/link_state |
| e002752 | 19:57:59.51815 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/5/persistence/aof_enabled |
| e002753 | 19:57:59.51815 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/5/persistence/loading |
| e002754 | 19:57:59.51815 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/5/persistence/aof_last_write_status |
| e002755 | 19:57:59.51815 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002756 | 19:57:59.51815 | node-6 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/5/process_live |
| e002757 | 19:57:59.659166 |  | client_probe | client_success | valid | `true` | /samples/10/probe/successful |
| e002758 | 19:57:59.776476 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/10/write_policy_probe/policy |
| e002759 | 19:57:59.776476 |  | ack_probe | ack.successful | valid | `false` | /samples/10/write_policy_probe/successful |
| e002760 | 19:57:59.776476 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/10/write_policy_probe/requested_replicas |
| e002761 | 19:57:59.776476 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/10/write_policy_probe/set_acknowledged |
| e002762 | 19:57:59.776476 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          0
        ]` | /samples/10/write_policy_probe/acknowledgements |
| e002763 | 19:57:59.776476 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/10/write_policy_probe/acknowledgement_observed |
| e002764 | 19:57:59.776476 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/10/write_policy_probe/waitaof_requested_local |
| e002765 | 19:57:59.776476 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/10/write_policy_probe/waitaof_requested_replicas |
| e002766 | 19:57:59.776476 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/10/write_policy_probe/waitaof_local_acknowledgements |
| e002767 | 19:57:59.776476 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/10/write_policy_probe/waitaof_replica_acknowledgements |
| e002768 | 19:57:59.776476 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/10/write_policy_probe/waitaof_acknowledgement_observed |
| e002769 | 19:58:00.287965 | node-6 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/5/reachable |
| e002770 | 19:58:00.287965 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/11/snapshot/views/5/replication/role |
| e002771 | 19:58:00.287965 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/11/snapshot/views/5/replication/connected_slaves |
| e002772 | 19:58:00.287965 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/11/snapshot/views/5/replication/master_link_status |
| e002773 | 19:58:00.287965 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/11/snapshot/views/5/replication/master_sync_in_progress |
| e002774 | 19:58:00.287965 | node-6 | info_replication | replication.master_repl_offset | valid | `89555` | /samples/11/snapshot/views/5/replication/master_repl_offset |
| e002775 | 19:58:00.287965 | node-6 | info_replication | replication.slave_repl_offset | valid | `89555` | /samples/11/snapshot/views/5/replication/slave_repl_offset |
| e002776 | 19:58:00.287965 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/11/snapshot/views/5/redis_id |
| e002777 | 19:58:00.287965 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/5/coverage |
| e002778 | 19:58:00.287965 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/11/snapshot/views/5/nodes/0/flags |
| e002779 | 19:58:00.287965 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/11/snapshot/views/5/nodes/0/primary_id |
| e002780 | 19:58:00.287965 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/0/slots |
| e002781 | 19:58:00.287965 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/0/link_state |
| e002782 | 19:58:00.287965 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/5/nodes/1/flags |
| e002783 | 19:58:00.287965 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/1/primary_id |
| e002784 | 19:58:00.287965 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/5/nodes/1/slots |
| e002785 | 19:58:00.287965 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/1/link_state |
| e002786 | 19:58:00.287965 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/5/nodes/2/flags |
| e002787 | 19:58:00.287965 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/11/snapshot/views/5/nodes/2/primary_id |
| e002788 | 19:58:00.287965 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/2/slots |
| e002789 | 19:58:00.287965 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/2/link_state |
| e002790 | 19:58:00.287965 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/5/nodes/3/flags |
| e002791 | 19:58:00.287965 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/3/primary_id |
| e002792 | 19:58:00.287965 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/5/nodes/3/slots |
| e002793 | 19:58:00.287965 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/3/link_state |
| e002794 | 19:58:00.287965 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/5/nodes/4/flags |
| e002795 | 19:58:00.287965 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/11/snapshot/views/5/nodes/4/primary_id |
| e002796 | 19:58:00.287965 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/4/slots |
| e002797 | 19:58:00.287965 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/4/link_state |
| e002798 | 19:58:00.287965 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/5/nodes/5/flags |
| e002799 | 19:58:00.287965 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/5/primary_id |
| e002800 | 19:58:00.287965 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/5/nodes/5/slots |
| e002801 | 19:58:00.287965 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/5/link_state |
| e002802 | 19:58:00.287965 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/5/persistence/aof_enabled |
| e002803 | 19:58:00.287965 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/5/persistence/loading |
| e002804 | 19:58:00.287965 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/5/persistence/aof_last_write_status |
| e002805 | 19:58:00.287965 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002806 | 19:58:00.287965 | node-6 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/5/process_live |
| e002807 | 19:58:00.287967 | node-3 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/2/reachable |
| e002808 | 19:58:00.287967 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/11/snapshot/views/2/replication/role |
| e002809 | 19:58:00.287967 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/11/snapshot/views/2/replication/connected_slaves |
| e002810 | 19:58:00.287967 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/11/snapshot/views/2/replication/master_link_status |
| e002811 | 19:58:00.287967 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/11/snapshot/views/2/replication/master_sync_in_progress |
| e002812 | 19:58:00.287967 | node-3 | info_replication | replication.master_repl_offset | valid | `89518` | /samples/11/snapshot/views/2/replication/master_repl_offset |
| e002813 | 19:58:00.287967 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/11/snapshot/views/2/replication/slave_repl_offset |
| e002814 | 19:58:00.287967 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/11/snapshot/views/2/redis_id |
| e002815 | 19:58:00.287967 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/2/coverage |
| e002816 | 19:58:00.287967 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/2/nodes/0/flags |
| e002817 | 19:58:00.287967 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/11/snapshot/views/2/nodes/0/primary_id |
| e002818 | 19:58:00.287967 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/0/slots |
| e002819 | 19:58:00.287967 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/0/link_state |
| e002820 | 19:58:00.287967 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/11/snapshot/views/2/nodes/1/flags |
| e002821 | 19:58:00.287967 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/1/primary_id |
| e002822 | 19:58:00.287967 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/2/nodes/1/slots |
| e002823 | 19:58:00.287967 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/1/link_state |
| e002824 | 19:58:00.287967 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/2/nodes/2/flags |
| e002825 | 19:58:00.287967 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/11/snapshot/views/2/nodes/2/primary_id |
| e002826 | 19:58:00.287967 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/2/slots |
| e002827 | 19:58:00.287967 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/2/link_state |
| e002828 | 19:58:00.287967 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/2/nodes/3/flags |
| e002829 | 19:58:00.287967 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/3/primary_id |
| e002830 | 19:58:00.287967 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/2/nodes/3/slots |
| e002831 | 19:58:00.287967 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/3/link_state |
| e002832 | 19:58:00.287967 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/2/nodes/4/flags |
| e002833 | 19:58:00.287967 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/4/primary_id |
| e002834 | 19:58:00.287967 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/2/nodes/4/slots |
| e002835 | 19:58:00.287967 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/4/link_state |
| e002836 | 19:58:00.287967 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/2/nodes/5/flags |
| e002837 | 19:58:00.287967 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/11/snapshot/views/2/nodes/5/primary_id |
| e002838 | 19:58:00.287967 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/5/slots |
| e002839 | 19:58:00.287967 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/5/link_state |
| e002840 | 19:58:00.287967 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/2/persistence/aof_enabled |
| e002841 | 19:58:00.287967 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/2/persistence/loading |
| e002842 | 19:58:00.287967 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/2/persistence/aof_last_write_status |
| e002843 | 19:58:00.287967 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e002844 | 19:58:00.287967 | node-3 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/2/process_live |
| e002845 | 19:58:00.288003 | node-1 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/0/reachable |
| e002846 | 19:58:00.288003 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/11/snapshot/views/0/replication/role |
| e002847 | 19:58:00.288003 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/11/snapshot/views/0/replication/connected_slaves |
| e002848 | 19:58:00.288003 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/11/snapshot/views/0/replication/master_link_status |
| e002849 | 19:58:00.288003 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/11/snapshot/views/0/replication/master_sync_in_progress |
| e002850 | 19:58:00.288003 | node-1 | info_replication | replication.master_repl_offset | valid | `89552` | /samples/11/snapshot/views/0/replication/master_repl_offset |
| e002851 | 19:58:00.288003 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/11/snapshot/views/0/replication/slave_repl_offset |
| e002852 | 19:58:00.288003 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/11/snapshot/views/0/redis_id |
| e002853 | 19:58:00.288003 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/0/coverage |
| e002854 | 19:58:00.288003 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/0/nodes/0/flags |
| e002855 | 19:58:00.288003 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/11/snapshot/views/0/nodes/0/primary_id |
| e002856 | 19:58:00.288003 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/0/slots |
| e002857 | 19:58:00.288003 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/0/link_state |
| e002858 | 19:58:00.288003 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/0/nodes/1/flags |
| e002859 | 19:58:00.288003 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/11/snapshot/views/0/nodes/1/primary_id |
| e002860 | 19:58:00.288003 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/1/slots |
| e002861 | 19:58:00.288003 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/1/link_state |
| e002862 | 19:58:00.288003 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/0/nodes/2/flags |
| e002863 | 19:58:00.288003 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/2/primary_id |
| e002864 | 19:58:00.288003 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/0/nodes/2/slots |
| e002865 | 19:58:00.288003 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/2/link_state |
| e002866 | 19:58:00.288003 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/11/snapshot/views/0/nodes/3/flags |
| e002867 | 19:58:00.288003 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/3/primary_id |
| e002868 | 19:58:00.288003 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/0/nodes/3/slots |
| e002869 | 19:58:00.288003 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/3/link_state |
| e002870 | 19:58:00.288003 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/0/nodes/4/flags |
| e002871 | 19:58:00.288003 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/11/snapshot/views/0/nodes/4/primary_id |
| e002872 | 19:58:00.288003 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/4/slots |
| e002873 | 19:58:00.288003 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/4/link_state |
| e002874 | 19:58:00.288003 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/0/nodes/5/flags |
| e002875 | 19:58:00.288003 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/5/primary_id |
| e002876 | 19:58:00.288003 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/0/nodes/5/slots |
| e002877 | 19:58:00.288003 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/5/link_state |
| e002878 | 19:58:00.288003 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/0/persistence/aof_enabled |
| e002879 | 19:58:00.288003 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/0/persistence/loading |
| e002880 | 19:58:00.288003 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/0/persistence/aof_last_write_status |
| e002881 | 19:58:00.288003 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e002882 | 19:58:00.288003 | node-1 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/0/process_live |
| e002883 | 19:58:00.288034 | node-2 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/1/reachable |
| e002884 | 19:58:00.288034 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/11/snapshot/views/1/replication/role |
| e002885 | 19:58:00.288034 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/11/snapshot/views/1/replication/connected_slaves |
| e002886 | 19:58:00.288034 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/11/snapshot/views/1/replication/master_link_status |
| e002887 | 19:58:00.288034 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/11/snapshot/views/1/replication/master_sync_in_progress |
| e002888 | 19:58:00.288034 | node-2 | info_replication | replication.master_repl_offset | valid | `89555` | /samples/11/snapshot/views/1/replication/master_repl_offset |
| e002889 | 19:58:00.288034 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/11/snapshot/views/1/replication/slave_repl_offset |
| e002890 | 19:58:00.288034 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/11/snapshot/views/1/redis_id |
| e002891 | 19:58:00.288034 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/1/coverage |
| e002892 | 19:58:00.288034 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/1/nodes/0/flags |
| e002893 | 19:58:00.288034 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/0/primary_id |
| e002894 | 19:58:00.288034 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/1/nodes/0/slots |
| e002895 | 19:58:00.288034 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/0/link_state |
| e002896 | 19:58:00.288034 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/1/nodes/1/flags |
| e002897 | 19:58:00.288034 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/11/snapshot/views/1/nodes/1/primary_id |
| e002898 | 19:58:00.288034 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/1/slots |
| e002899 | 19:58:00.288034 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/1/link_state |
| e002900 | 19:58:00.288034 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/1/nodes/2/flags |
| e002901 | 19:58:00.288034 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/11/snapshot/views/1/nodes/2/primary_id |
| e002902 | 19:58:00.288034 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/2/slots |
| e002903 | 19:58:00.288034 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/2/link_state |
| e002904 | 19:58:00.288034 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/1/nodes/3/flags |
| e002905 | 19:58:00.288034 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/11/snapshot/views/1/nodes/3/primary_id |
| e002906 | 19:58:00.288034 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/3/slots |
| e002907 | 19:58:00.288034 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/3/link_state |
| e002908 | 19:58:00.288034 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/11/snapshot/views/1/nodes/4/flags |
| e002909 | 19:58:00.288034 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/4/primary_id |
| e002910 | 19:58:00.288034 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/1/nodes/4/slots |
| e002911 | 19:58:00.288034 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/4/link_state |
| e002912 | 19:58:00.288034 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/1/nodes/5/flags |
| e002913 | 19:58:00.288034 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/5/primary_id |
| e002914 | 19:58:00.288034 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/1/nodes/5/slots |
| e002915 | 19:58:00.288034 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/5/link_state |
| e002916 | 19:58:00.288034 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/1/persistence/aof_enabled |
| e002917 | 19:58:00.288034 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/1/persistence/loading |
| e002918 | 19:58:00.288034 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/1/persistence/aof_last_write_status |
| e002919 | 19:58:00.288034 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e002920 | 19:58:00.288034 | node-2 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/1/process_live |
| e002921 | 19:58:00.28805 | node-4 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/3/reachable |
| e002922 | 19:58:00.28805 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/11/snapshot/views/3/replication/role |
| e002923 | 19:58:00.28805 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/11/snapshot/views/3/replication/connected_slaves |
| e002924 | 19:58:00.28805 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/11/snapshot/views/3/replication/master_link_status |
| e002925 | 19:58:00.28805 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/11/snapshot/views/3/replication/master_sync_in_progress |
| e002926 | 19:58:00.28805 | node-4 | info_replication | replication.master_repl_offset | valid | `89518` | /samples/11/snapshot/views/3/replication/master_repl_offset |
| e002927 | 19:58:00.28805 | node-4 | info_replication | replication.slave_repl_offset | valid | `89518` | /samples/11/snapshot/views/3/replication/slave_repl_offset |
| e002928 | 19:58:00.28805 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/11/snapshot/views/3/redis_id |
| e002929 | 19:58:00.28805 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/3/coverage |
| e002930 | 19:58:00.28805 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/3/nodes/0/flags |
| e002931 | 19:58:00.28805 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/0/primary_id |
| e002932 | 19:58:00.28805 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/3/nodes/0/slots |
| e002933 | 19:58:00.28805 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/0/link_state |
| e002934 | 19:58:00.28805 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/3/nodes/1/flags |
| e002935 | 19:58:00.28805 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/11/snapshot/views/3/nodes/1/primary_id |
| e002936 | 19:58:00.28805 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/1/slots |
| e002937 | 19:58:00.28805 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/1/link_state |
| e002938 | 19:58:00.28805 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/3/nodes/2/flags |
| e002939 | 19:58:00.28805 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/11/snapshot/views/3/nodes/2/primary_id |
| e002940 | 19:58:00.28805 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/2/slots |
| e002941 | 19:58:00.28805 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/2/link_state |
| e002942 | 19:58:00.28805 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/11/snapshot/views/3/nodes/3/flags |
| e002943 | 19:58:00.28805 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/11/snapshot/views/3/nodes/3/primary_id |
| e002944 | 19:58:00.28805 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/3/slots |
| e002945 | 19:58:00.28805 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/3/link_state |
| e002946 | 19:58:00.28805 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/3/nodes/4/flags |
| e002947 | 19:58:00.28805 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/4/primary_id |
| e002948 | 19:58:00.28805 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/3/nodes/4/slots |
| e002949 | 19:58:00.28805 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/4/link_state |
| e002950 | 19:58:00.28805 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/3/nodes/5/flags |
| e002951 | 19:58:00.28805 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/5/primary_id |
| e002952 | 19:58:00.28805 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/3/nodes/5/slots |
| e002953 | 19:58:00.28805 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/5/link_state |
| e002954 | 19:58:00.28805 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/3/persistence/aof_enabled |
| e002955 | 19:58:00.28805 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/3/persistence/loading |
| e002956 | 19:58:00.28805 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/3/persistence/aof_last_write_status |
| e002957 | 19:58:00.28805 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e002958 | 19:58:00.28805 | node-4 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/3/process_live |
| e002959 | 19:58:00.288057 | node-5 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/4/reachable |
| e002960 | 19:58:00.288057 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/11/snapshot/views/4/replication/role |
| e002961 | 19:58:00.288057 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/11/snapshot/views/4/replication/connected_slaves |
| e002962 | 19:58:00.288057 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/11/snapshot/views/4/replication/master_link_status |
| e002963 | 19:58:00.288057 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/11/snapshot/views/4/replication/master_sync_in_progress |
| e002964 | 19:58:00.288057 | node-5 | info_replication | replication.master_repl_offset | valid | `0` | /samples/11/snapshot/views/4/replication/master_repl_offset |
| e002965 | 19:58:00.288057 | node-5 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/11/snapshot/views/4/replication/slave_repl_offset |
| e002966 | 19:58:00.288057 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /samples/11/snapshot/views/4/redis_id |
| e002967 | 19:58:00.288057 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/4/coverage |
| e002968 | 19:58:00.288057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/4/nodes/0/flags |
| e002969 | 19:58:00.288057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/0/primary_id |
| e002970 | 19:58:00.288057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/4/nodes/0/slots |
| e002971 | 19:58:00.288057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/0/link_state |
| e002972 | 19:58:00.288057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/4/nodes/1/flags |
| e002973 | 19:58:00.288057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/11/snapshot/views/4/nodes/1/primary_id |
| e002974 | 19:58:00.288057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/1/slots |
| e002975 | 19:58:00.288057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/1/link_state |
| e002976 | 19:58:00.288057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/11/snapshot/views/4/nodes/2/flags |
| e002977 | 19:58:00.288057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/11/snapshot/views/4/nodes/2/primary_id |
| e002978 | 19:58:00.288057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/2/slots |
| e002979 | 19:58:00.288057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/2/link_state |
| e002980 | 19:58:00.288057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/4/nodes/3/flags |
| e002981 | 19:58:00.288057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/3/primary_id |
| e002982 | 19:58:00.288057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/4/nodes/3/slots |
| e002983 | 19:58:00.288057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/3/link_state |
| e002984 | 19:58:00.288057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/4/nodes/4/flags |
| e002985 | 19:58:00.288057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/11/snapshot/views/4/nodes/4/primary_id |
| e002986 | 19:58:00.288057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/4/slots |
| e002987 | 19:58:00.288057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/4/link_state |
| e002988 | 19:58:00.288057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/4/nodes/5/flags |
| e002989 | 19:58:00.288057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/5/primary_id |
| e002990 | 19:58:00.288057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/4/nodes/5/slots |
| e002991 | 19:58:00.288057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/5/link_state |
| e002992 | 19:58:00.288057 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/4/persistence/aof_enabled |
| e002993 | 19:58:00.288057 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/4/persistence/loading |
| e002994 | 19:58:00.288057 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/4/persistence/aof_last_write_status |
| e002995 | 19:58:00.288057 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e002996 | 19:58:00.288057 | node-5 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/4/process_live |
| e002997 | 19:58:00.381014 |  | client_window | client.attempts | valid | `172` | /metrics/windows/11/attempts |
| e002998 | 19:58:00.381014 |  | client_window | client.errors | valid | `29` | /metrics/windows/11/errors |
| e002999 | 19:58:00.381014 |  | client_window | client.successes | valid | `143` | /metrics/windows/11/successes |
| e003000 | 19:58:00.381014 |  | client_window | client.p95_ms | valid | `116.981575` | /metrics/windows/11/p95_ms |
| e003001 | 19:58:00.456292 |  | client_probe | client_success | valid | `true` | /samples/11/probe/successful |
| e003002 | 19:58:00.4729 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/11/write_policy_probe/policy |
| e003003 | 19:58:00.4729 |  | ack_probe | ack.successful | valid | `true` | /samples/11/write_policy_probe/successful |
| e003004 | 19:58:00.4729 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/11/write_policy_probe/requested_replicas |
| e003005 | 19:58:00.4729 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/11/write_policy_probe/set_acknowledged |
| e003006 | 19:58:00.4729 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          1
        ]` | /samples/11/write_policy_probe/acknowledgements |
| e003007 | 19:58:00.4729 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/11/write_policy_probe/acknowledgement_observed |
| e003008 | 19:58:00.4729 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/11/write_policy_probe/waitaof_requested_local |
| e003009 | 19:58:00.4729 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/11/write_policy_probe/waitaof_requested_replicas |
| e003010 | 19:58:00.4729 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/11/write_policy_probe/waitaof_local_acknowledgements |
| e003011 | 19:58:00.4729 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/11/write_policy_probe/waitaof_replica_acknowledgements |
| e003012 | 19:58:00.4729 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/11/write_policy_probe/waitaof_acknowledgement_observed |
| e003013 | 19:58:01.0172 | node-1 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/0/reachable |
| e003014 | 19:58:01.0172 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/12/snapshot/views/0/replication/role |
| e003015 | 19:58:01.0172 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/12/snapshot/views/0/replication/connected_slaves |
| e003016 | 19:58:01.0172 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/12/snapshot/views/0/replication/master_link_status |
| e003017 | 19:58:01.0172 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/12/snapshot/views/0/replication/master_sync_in_progress |
| e003018 | 19:58:01.0172 | node-1 | info_replication | replication.master_repl_offset | valid | `95006` | /samples/12/snapshot/views/0/replication/master_repl_offset |
| e003019 | 19:58:01.0172 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/12/snapshot/views/0/replication/slave_repl_offset |
| e003020 | 19:58:01.0172 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/12/snapshot/views/0/redis_id |
| e003021 | 19:58:01.0172 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/0/coverage |
| e003022 | 19:58:01.0172 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/0/nodes/0/flags |
| e003023 | 19:58:01.0172 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/12/snapshot/views/0/nodes/0/primary_id |
| e003024 | 19:58:01.0172 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/0/slots |
| e003025 | 19:58:01.0172 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/0/link_state |
| e003026 | 19:58:01.0172 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/0/nodes/1/flags |
| e003027 | 19:58:01.0172 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/12/snapshot/views/0/nodes/1/primary_id |
| e003028 | 19:58:01.0172 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/1/slots |
| e003029 | 19:58:01.0172 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/1/link_state |
| e003030 | 19:58:01.0172 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/0/nodes/2/flags |
| e003031 | 19:58:01.0172 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/2/primary_id |
| e003032 | 19:58:01.0172 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/0/nodes/2/slots |
| e003033 | 19:58:01.0172 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/2/link_state |
| e003034 | 19:58:01.0172 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/12/snapshot/views/0/nodes/3/flags |
| e003035 | 19:58:01.0172 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/3/primary_id |
| e003036 | 19:58:01.0172 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/0/nodes/3/slots |
| e003037 | 19:58:01.0172 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/3/link_state |
| e003038 | 19:58:01.0172 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/0/nodes/4/flags |
| e003039 | 19:58:01.0172 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/12/snapshot/views/0/nodes/4/primary_id |
| e003040 | 19:58:01.0172 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/4/slots |
| e003041 | 19:58:01.0172 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/4/link_state |
| e003042 | 19:58:01.0172 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/0/nodes/5/flags |
| e003043 | 19:58:01.0172 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/5/primary_id |
| e003044 | 19:58:01.0172 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/0/nodes/5/slots |
| e003045 | 19:58:01.0172 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/5/link_state |
| e003046 | 19:58:01.0172 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/0/persistence/aof_enabled |
| e003047 | 19:58:01.0172 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/0/persistence/loading |
| e003048 | 19:58:01.0172 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/0/persistence/aof_last_write_status |
| e003049 | 19:58:01.0172 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e003050 | 19:58:01.0172 | node-1 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/0/process_live |
| e003051 | 19:58:01.017243 | node-3 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/2/reachable |
| e003052 | 19:58:01.017243 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/12/snapshot/views/2/replication/role |
| e003053 | 19:58:01.017243 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/12/snapshot/views/2/replication/connected_slaves |
| e003054 | 19:58:01.017243 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/12/snapshot/views/2/replication/master_link_status |
| e003055 | 19:58:01.017243 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/12/snapshot/views/2/replication/master_sync_in_progress |
| e003056 | 19:58:01.017243 | node-3 | info_replication | replication.master_repl_offset | valid | `94912` | /samples/12/snapshot/views/2/replication/master_repl_offset |
| e003057 | 19:58:01.017243 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/12/snapshot/views/2/replication/slave_repl_offset |
| e003058 | 19:58:01.017243 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/12/snapshot/views/2/redis_id |
| e003059 | 19:58:01.017243 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/2/coverage |
| e003060 | 19:58:01.017243 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/2/nodes/0/flags |
| e003061 | 19:58:01.017243 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/12/snapshot/views/2/nodes/0/primary_id |
| e003062 | 19:58:01.017243 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/0/slots |
| e003063 | 19:58:01.017243 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/0/link_state |
| e003064 | 19:58:01.017243 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/12/snapshot/views/2/nodes/1/flags |
| e003065 | 19:58:01.017243 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/1/primary_id |
| e003066 | 19:58:01.017243 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/2/nodes/1/slots |
| e003067 | 19:58:01.017243 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/1/link_state |
| e003068 | 19:58:01.017243 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/2/nodes/2/flags |
| e003069 | 19:58:01.017243 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/12/snapshot/views/2/nodes/2/primary_id |
| e003070 | 19:58:01.017243 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/2/slots |
| e003071 | 19:58:01.017243 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/2/link_state |
| e003072 | 19:58:01.017243 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/2/nodes/3/flags |
| e003073 | 19:58:01.017243 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/3/primary_id |
| e003074 | 19:58:01.017243 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/2/nodes/3/slots |
| e003075 | 19:58:01.017243 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/3/link_state |
| e003076 | 19:58:01.017243 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/2/nodes/4/flags |
| e003077 | 19:58:01.017243 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/4/primary_id |
| e003078 | 19:58:01.017243 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/2/nodes/4/slots |
| e003079 | 19:58:01.017243 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/4/link_state |
| e003080 | 19:58:01.017243 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/2/nodes/5/flags |
| e003081 | 19:58:01.017243 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/12/snapshot/views/2/nodes/5/primary_id |
| e003082 | 19:58:01.017243 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/5/slots |
| e003083 | 19:58:01.017243 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/5/link_state |
| e003084 | 19:58:01.017243 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/2/persistence/aof_enabled |
| e003085 | 19:58:01.017243 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/2/persistence/loading |
| e003086 | 19:58:01.017243 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/2/persistence/aof_last_write_status |
| e003087 | 19:58:01.017243 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e003088 | 19:58:01.017243 | node-3 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/2/process_live |
| e003089 | 19:58:01.01729 | node-2 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/1/reachable |
| e003090 | 19:58:01.01729 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/12/snapshot/views/1/replication/role |
| e003091 | 19:58:01.01729 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/12/snapshot/views/1/replication/connected_slaves |
| e003092 | 19:58:01.01729 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/12/snapshot/views/1/replication/master_link_status |
| e003093 | 19:58:01.01729 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/12/snapshot/views/1/replication/master_sync_in_progress |
| e003094 | 19:58:01.01729 | node-2 | info_replication | replication.master_repl_offset | valid | `94986` | /samples/12/snapshot/views/1/replication/master_repl_offset |
| e003095 | 19:58:01.01729 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/12/snapshot/views/1/replication/slave_repl_offset |
| e003096 | 19:58:01.01729 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/12/snapshot/views/1/redis_id |
| e003097 | 19:58:01.01729 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/1/coverage |
| e003098 | 19:58:01.01729 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/1/nodes/0/flags |
| e003099 | 19:58:01.01729 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/0/primary_id |
| e003100 | 19:58:01.01729 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/1/nodes/0/slots |
| e003101 | 19:58:01.01729 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/0/link_state |
| e003102 | 19:58:01.01729 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/1/nodes/1/flags |
| e003103 | 19:58:01.01729 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/12/snapshot/views/1/nodes/1/primary_id |
| e003104 | 19:58:01.01729 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/1/slots |
| e003105 | 19:58:01.01729 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/1/link_state |
| e003106 | 19:58:01.01729 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/1/nodes/2/flags |
| e003107 | 19:58:01.01729 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/12/snapshot/views/1/nodes/2/primary_id |
| e003108 | 19:58:01.01729 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/2/slots |
| e003109 | 19:58:01.01729 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/2/link_state |
| e003110 | 19:58:01.01729 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/1/nodes/3/flags |
| e003111 | 19:58:01.01729 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/12/snapshot/views/1/nodes/3/primary_id |
| e003112 | 19:58:01.01729 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/3/slots |
| e003113 | 19:58:01.01729 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/3/link_state |
| e003114 | 19:58:01.01729 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/12/snapshot/views/1/nodes/4/flags |
| e003115 | 19:58:01.01729 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/4/primary_id |
| e003116 | 19:58:01.01729 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/1/nodes/4/slots |
| e003117 | 19:58:01.01729 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/4/link_state |
| e003118 | 19:58:01.01729 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/1/nodes/5/flags |
| e003119 | 19:58:01.01729 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/5/primary_id |
| e003120 | 19:58:01.01729 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/1/nodes/5/slots |
| e003121 | 19:58:01.01729 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/5/link_state |
| e003122 | 19:58:01.01729 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/1/persistence/aof_enabled |
| e003123 | 19:58:01.01729 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/1/persistence/loading |
| e003124 | 19:58:01.01729 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/1/persistence/aof_last_write_status |
| e003125 | 19:58:01.01729 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e003126 | 19:58:01.01729 | node-2 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/1/process_live |
| e003127 | 19:58:01.017315 | node-4 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/3/reachable |
| e003128 | 19:58:01.017315 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/12/snapshot/views/3/replication/role |
| e003129 | 19:58:01.017315 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/12/snapshot/views/3/replication/connected_slaves |
| e003130 | 19:58:01.017315 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/12/snapshot/views/3/replication/master_link_status |
| e003131 | 19:58:01.017315 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/12/snapshot/views/3/replication/master_sync_in_progress |
| e003132 | 19:58:01.017315 | node-4 | info_replication | replication.master_repl_offset | valid | `94912` | /samples/12/snapshot/views/3/replication/master_repl_offset |
| e003133 | 19:58:01.017315 | node-4 | info_replication | replication.slave_repl_offset | valid | `94912` | /samples/12/snapshot/views/3/replication/slave_repl_offset |
| e003134 | 19:58:01.017315 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/12/snapshot/views/3/redis_id |
| e003135 | 19:58:01.017315 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/3/coverage |
| e003136 | 19:58:01.017315 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/3/nodes/0/flags |
| e003137 | 19:58:01.017315 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/0/primary_id |
| e003138 | 19:58:01.017315 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/3/nodes/0/slots |
| e003139 | 19:58:01.017315 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/0/link_state |
| e003140 | 19:58:01.017315 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/3/nodes/1/flags |
| e003141 | 19:58:01.017315 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/12/snapshot/views/3/nodes/1/primary_id |
| e003142 | 19:58:01.017315 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/1/slots |
| e003143 | 19:58:01.017315 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/1/link_state |
| e003144 | 19:58:01.017315 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/3/nodes/2/flags |
| e003145 | 19:58:01.017315 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/12/snapshot/views/3/nodes/2/primary_id |
| e003146 | 19:58:01.017315 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/2/slots |
| e003147 | 19:58:01.017315 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/2/link_state |
| e003148 | 19:58:01.017315 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/12/snapshot/views/3/nodes/3/flags |
| e003149 | 19:58:01.017315 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/12/snapshot/views/3/nodes/3/primary_id |
| e003150 | 19:58:01.017315 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/3/slots |
| e003151 | 19:58:01.017315 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/3/link_state |
| e003152 | 19:58:01.017315 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/3/nodes/4/flags |
| e003153 | 19:58:01.017315 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/4/primary_id |
| e003154 | 19:58:01.017315 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/3/nodes/4/slots |
| e003155 | 19:58:01.017315 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/4/link_state |
| e003156 | 19:58:01.017315 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/3/nodes/5/flags |
| e003157 | 19:58:01.017315 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/5/primary_id |
| e003158 | 19:58:01.017315 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/3/nodes/5/slots |
| e003159 | 19:58:01.017315 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/5/link_state |
| e003160 | 19:58:01.017315 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/3/persistence/aof_enabled |
| e003161 | 19:58:01.017315 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/3/persistence/loading |
| e003162 | 19:58:01.017315 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/3/persistence/aof_last_write_status |
| e003163 | 19:58:01.017315 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e003164 | 19:58:01.017315 | node-4 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/3/process_live |
| e003165 | 19:58:01.017347 | node-6 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/5/reachable |
| e003166 | 19:58:01.017347 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/12/snapshot/views/5/replication/role |
| e003167 | 19:58:01.017347 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/12/snapshot/views/5/replication/connected_slaves |
| e003168 | 19:58:01.017347 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/12/snapshot/views/5/replication/master_link_status |
| e003169 | 19:58:01.017347 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/12/snapshot/views/5/replication/master_sync_in_progress |
| e003170 | 19:58:01.017347 | node-6 | info_replication | replication.master_repl_offset | valid | `94986` | /samples/12/snapshot/views/5/replication/master_repl_offset |
| e003171 | 19:58:01.017347 | node-6 | info_replication | replication.slave_repl_offset | valid | `94986` | /samples/12/snapshot/views/5/replication/slave_repl_offset |
| e003172 | 19:58:01.017347 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/12/snapshot/views/5/redis_id |
| e003173 | 19:58:01.017347 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/5/coverage |
| e003174 | 19:58:01.017347 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/12/snapshot/views/5/nodes/0/flags |
| e003175 | 19:58:01.017347 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/12/snapshot/views/5/nodes/0/primary_id |
| e003176 | 19:58:01.017347 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/0/slots |
| e003177 | 19:58:01.017347 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/0/link_state |
| e003178 | 19:58:01.017347 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/5/nodes/1/flags |
| e003179 | 19:58:01.017347 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/1/primary_id |
| e003180 | 19:58:01.017347 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/5/nodes/1/slots |
| e003181 | 19:58:01.017347 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/1/link_state |
| e003182 | 19:58:01.017347 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/5/nodes/2/flags |
| e003183 | 19:58:01.017347 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/12/snapshot/views/5/nodes/2/primary_id |
| e003184 | 19:58:01.017347 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/2/slots |
| e003185 | 19:58:01.017347 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/2/link_state |
| e003186 | 19:58:01.017347 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/5/nodes/3/flags |
| e003187 | 19:58:01.017347 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/3/primary_id |
| e003188 | 19:58:01.017347 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/5/nodes/3/slots |
| e003189 | 19:58:01.017347 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/3/link_state |
| e003190 | 19:58:01.017347 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/5/nodes/4/flags |
| e003191 | 19:58:01.017347 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/12/snapshot/views/5/nodes/4/primary_id |
| e003192 | 19:58:01.017347 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/4/slots |
| e003193 | 19:58:01.017347 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/4/link_state |
| e003194 | 19:58:01.017347 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/5/nodes/5/flags |
| e003195 | 19:58:01.017347 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/5/primary_id |
| e003196 | 19:58:01.017347 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/5/nodes/5/slots |
| e003197 | 19:58:01.017347 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/5/link_state |
| e003198 | 19:58:01.017347 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/5/persistence/aof_enabled |
| e003199 | 19:58:01.017347 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/5/persistence/loading |
| e003200 | 19:58:01.017347 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/5/persistence/aof_last_write_status |
| e003201 | 19:58:01.017347 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e003202 | 19:58:01.017347 | node-6 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/5/process_live |
| e003203 | 19:58:01.017387 | node-5 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/4/reachable |
| e003204 | 19:58:01.017387 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/12/snapshot/views/4/replication/role |
| e003205 | 19:58:01.017387 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/12/snapshot/views/4/replication/connected_slaves |
| e003206 | 19:58:01.017387 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/12/snapshot/views/4/replication/master_link_status |
| e003207 | 19:58:01.017387 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/12/snapshot/views/4/replication/master_sync_in_progress |
| e003208 | 19:58:01.017387 | node-5 | info_replication | replication.master_repl_offset | valid | `95006` | /samples/12/snapshot/views/4/replication/master_repl_offset |
| e003209 | 19:58:01.017387 | node-5 | info_replication | replication.slave_repl_offset | valid | `95006` | /samples/12/snapshot/views/4/replication/slave_repl_offset |
| e003210 | 19:58:01.017387 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /samples/12/snapshot/views/4/redis_id |
| e003211 | 19:58:01.017387 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/4/coverage |
| e003212 | 19:58:01.017387 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/4/nodes/0/flags |
| e003213 | 19:58:01.017387 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/0/primary_id |
| e003214 | 19:58:01.017387 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/4/nodes/0/slots |
| e003215 | 19:58:01.017387 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/0/link_state |
| e003216 | 19:58:01.017387 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/4/nodes/1/flags |
| e003217 | 19:58:01.017387 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/12/snapshot/views/4/nodes/1/primary_id |
| e003218 | 19:58:01.017387 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/1/slots |
| e003219 | 19:58:01.017387 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/1/link_state |
| e003220 | 19:58:01.017387 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/12/snapshot/views/4/nodes/2/flags |
| e003221 | 19:58:01.017387 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/12/snapshot/views/4/nodes/2/primary_id |
| e003222 | 19:58:01.017387 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/2/slots |
| e003223 | 19:58:01.017387 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/2/link_state |
| e003224 | 19:58:01.017387 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/4/nodes/3/flags |
| e003225 | 19:58:01.017387 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/3/primary_id |
| e003226 | 19:58:01.017387 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/4/nodes/3/slots |
| e003227 | 19:58:01.017387 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/3/link_state |
| e003228 | 19:58:01.017387 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/4/nodes/4/flags |
| e003229 | 19:58:01.017387 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/12/snapshot/views/4/nodes/4/primary_id |
| e003230 | 19:58:01.017387 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/4/slots |
| e003231 | 19:58:01.017387 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/4/link_state |
| e003232 | 19:58:01.017387 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/4/nodes/5/flags |
| e003233 | 19:58:01.017387 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/5/primary_id |
| e003234 | 19:58:01.017387 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/4/nodes/5/slots |
| e003235 | 19:58:01.017387 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/5/link_state |
| e003236 | 19:58:01.017387 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/4/persistence/aof_enabled |
| e003237 | 19:58:01.017387 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/4/persistence/loading |
| e003238 | 19:58:01.017387 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/4/persistence/aof_last_write_status |
| e003239 | 19:58:01.017387 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e003240 | 19:58:01.017387 | node-5 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/4/process_live |
| e003241 | 19:58:01.164377 |  | client_probe | client_success | valid | `true` | /samples/12/probe/successful |
| e003242 | 19:58:01.1659 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/12/write_policy_probe/policy |
| e003243 | 19:58:01.1659 |  | ack_probe | ack.successful | valid | `true` | /samples/12/write_policy_probe/successful |
| e003244 | 19:58:01.1659 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/12/write_policy_probe/requested_replicas |
| e003245 | 19:58:01.1659 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/12/write_policy_probe/set_acknowledged |
| e003246 | 19:58:01.1659 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          1
        ]` | /samples/12/write_policy_probe/acknowledgements |
| e003247 | 19:58:01.1659 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/12/write_policy_probe/acknowledgement_observed |
| e003248 | 19:58:01.1659 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/12/write_policy_probe/waitaof_requested_local |
| e003249 | 19:58:01.1659 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/12/write_policy_probe/waitaof_requested_replicas |
| e003250 | 19:58:01.1659 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/12/write_policy_probe/waitaof_local_acknowledgements |
| e003251 | 19:58:01.1659 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/12/write_policy_probe/waitaof_replica_acknowledgements |
| e003252 | 19:58:01.1659 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/12/write_policy_probe/waitaof_acknowledgement_observed |
| e003253 | 19:58:01.381014 |  | client_window | client.attempts | valid | `172` | /metrics/windows/12/attempts |
| e003254 | 19:58:01.381014 |  | client_window | client.errors | valid | `2` | /metrics/windows/12/errors |
| e003255 | 19:58:01.381014 |  | client_window | client.successes | valid | `170` | /metrics/windows/12/successes |
| e003256 | 19:58:01.381014 |  | client_window | client.p95_ms | valid | `10.201222` | /metrics/windows/12/p95_ms |
| e003257 | 19:58:01.681942 | node-1 | node_client | reachable | valid | `true` | /samples/13/snapshot/views/0/reachable |
| e003258 | 19:58:01.681942 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/13/snapshot/views/0/replication/role |
| e003259 | 19:58:01.681942 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/13/snapshot/views/0/replication/connected_slaves |
| e003260 | 19:58:01.681942 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/13/snapshot/views/0/replication/master_link_status |
| e003261 | 19:58:01.681942 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/13/snapshot/views/0/replication/master_sync_in_progress |
| e003262 | 19:58:01.681942 | node-1 | info_replication | replication.master_repl_offset | valid | `100400` | /samples/13/snapshot/views/0/replication/master_repl_offset |
| e003263 | 19:58:01.681942 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/13/snapshot/views/0/replication/slave_repl_offset |
| e003264 | 19:58:01.681942 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/13/snapshot/views/0/redis_id |
| e003265 | 19:58:01.681942 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/13/snapshot/views/0/coverage |
| e003266 | 19:58:01.681942 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/0/nodes/0/flags |
| e003267 | 19:58:01.681942 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/13/snapshot/views/0/nodes/0/primary_id |
| e003268 | 19:58:01.681942 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/0/nodes/0/slots |
| e003269 | 19:58:01.681942 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/0/nodes/0/link_state |
| e003270 | 19:58:01.681942 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/0/nodes/1/flags |
| e003271 | 19:58:01.681942 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/13/snapshot/views/0/nodes/1/primary_id |
| e003272 | 19:58:01.681942 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/0/nodes/1/slots |
| e003273 | 19:58:01.681942 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/0/nodes/1/link_state |
| e003274 | 19:58:01.681942 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/0/nodes/2/flags |
| e003275 | 19:58:01.681942 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/0/nodes/2/primary_id |
| e003276 | 19:58:01.681942 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/13/snapshot/views/0/nodes/2/slots |
| e003277 | 19:58:01.681942 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/0/nodes/2/link_state |
| e003278 | 19:58:01.681942 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/13/snapshot/views/0/nodes/3/flags |
| e003279 | 19:58:01.681942 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/0/nodes/3/primary_id |
| e003280 | 19:58:01.681942 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/13/snapshot/views/0/nodes/3/slots |
| e003281 | 19:58:01.681942 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/0/nodes/3/link_state |
| e003282 | 19:58:01.681942 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/0/nodes/4/flags |
| e003283 | 19:58:01.681942 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/13/snapshot/views/0/nodes/4/primary_id |
| e003284 | 19:58:01.681942 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/0/nodes/4/slots |
| e003285 | 19:58:01.681942 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/0/nodes/4/link_state |
| e003286 | 19:58:01.681942 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/0/nodes/5/flags |
| e003287 | 19:58:01.681942 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/0/nodes/5/primary_id |
| e003288 | 19:58:01.681942 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/13/snapshot/views/0/nodes/5/slots |
| e003289 | 19:58:01.681942 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/0/nodes/5/link_state |
| e003290 | 19:58:01.681942 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/13/snapshot/views/0/persistence/aof_enabled |
| e003291 | 19:58:01.681942 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/13/snapshot/views/0/persistence/loading |
| e003292 | 19:58:01.681942 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/13/snapshot/views/0/persistence/aof_last_write_status |
| e003293 | 19:58:01.681942 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/13/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e003294 | 19:58:01.681942 | node-1 | admin | process_live | missing_from_schema | `` | /samples/13/snapshot/views/0/process_live |
| e003295 | 19:58:01.681948 | node-2 | node_client | reachable | valid | `true` | /samples/13/snapshot/views/1/reachable |
| e003296 | 19:58:01.681948 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/13/snapshot/views/1/replication/role |
| e003297 | 19:58:01.681948 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/13/snapshot/views/1/replication/connected_slaves |
| e003298 | 19:58:01.681948 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/13/snapshot/views/1/replication/master_link_status |
| e003299 | 19:58:01.681948 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/13/snapshot/views/1/replication/master_sync_in_progress |
| e003300 | 19:58:01.681948 | node-2 | info_replication | replication.master_repl_offset | valid | `100417` | /samples/13/snapshot/views/1/replication/master_repl_offset |
| e003301 | 19:58:01.681948 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/13/snapshot/views/1/replication/slave_repl_offset |
| e003302 | 19:58:01.681948 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/13/snapshot/views/1/redis_id |
| e003303 | 19:58:01.681948 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/13/snapshot/views/1/coverage |
| e003304 | 19:58:01.681948 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/1/nodes/0/flags |
| e003305 | 19:58:01.681948 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/1/nodes/0/primary_id |
| e003306 | 19:58:01.681948 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/13/snapshot/views/1/nodes/0/slots |
| e003307 | 19:58:01.681948 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/1/nodes/0/link_state |
| e003308 | 19:58:01.681948 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/1/nodes/1/flags |
| e003309 | 19:58:01.681948 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/13/snapshot/views/1/nodes/1/primary_id |
| e003310 | 19:58:01.681948 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/1/nodes/1/slots |
| e003311 | 19:58:01.681948 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/1/nodes/1/link_state |
| e003312 | 19:58:01.681948 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/1/nodes/2/flags |
| e003313 | 19:58:01.681948 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/13/snapshot/views/1/nodes/2/primary_id |
| e003314 | 19:58:01.681948 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/1/nodes/2/slots |
| e003315 | 19:58:01.681948 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/1/nodes/2/link_state |
| e003316 | 19:58:01.681948 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/1/nodes/3/flags |
| e003317 | 19:58:01.681948 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/13/snapshot/views/1/nodes/3/primary_id |
| e003318 | 19:58:01.681948 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/1/nodes/3/slots |
| e003319 | 19:58:01.681948 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/1/nodes/3/link_state |
| e003320 | 19:58:01.681948 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/13/snapshot/views/1/nodes/4/flags |
| e003321 | 19:58:01.681948 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/1/nodes/4/primary_id |
| e003322 | 19:58:01.681948 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/13/snapshot/views/1/nodes/4/slots |
| e003323 | 19:58:01.681948 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/1/nodes/4/link_state |
| e003324 | 19:58:01.681948 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/1/nodes/5/flags |
| e003325 | 19:58:01.681948 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/1/nodes/5/primary_id |
| e003326 | 19:58:01.681948 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/13/snapshot/views/1/nodes/5/slots |
| e003327 | 19:58:01.681948 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/1/nodes/5/link_state |
| e003328 | 19:58:01.681948 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/13/snapshot/views/1/persistence/aof_enabled |
| e003329 | 19:58:01.681948 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/13/snapshot/views/1/persistence/loading |
| e003330 | 19:58:01.681948 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/13/snapshot/views/1/persistence/aof_last_write_status |
| e003331 | 19:58:01.681948 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/13/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e003332 | 19:58:01.681948 | node-2 | admin | process_live | missing_from_schema | `` | /samples/13/snapshot/views/1/process_live |
| e003333 | 19:58:01.681949 | node-6 | node_client | reachable | valid | `true` | /samples/13/snapshot/views/5/reachable |
| e003334 | 19:58:01.681949 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/13/snapshot/views/5/replication/role |
| e003335 | 19:58:01.681949 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/13/snapshot/views/5/replication/connected_slaves |
| e003336 | 19:58:01.681949 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/13/snapshot/views/5/replication/master_link_status |
| e003337 | 19:58:01.681949 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/13/snapshot/views/5/replication/master_sync_in_progress |
| e003338 | 19:58:01.681949 | node-6 | info_replication | replication.master_repl_offset | valid | `100417` | /samples/13/snapshot/views/5/replication/master_repl_offset |
| e003339 | 19:58:01.681949 | node-6 | info_replication | replication.slave_repl_offset | valid | `100417` | /samples/13/snapshot/views/5/replication/slave_repl_offset |
| e003340 | 19:58:01.681949 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/13/snapshot/views/5/redis_id |
| e003341 | 19:58:01.681949 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/13/snapshot/views/5/coverage |
| e003342 | 19:58:01.681949 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/13/snapshot/views/5/nodes/0/flags |
| e003343 | 19:58:01.681949 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/13/snapshot/views/5/nodes/0/primary_id |
| e003344 | 19:58:01.681949 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/5/nodes/0/slots |
| e003345 | 19:58:01.681949 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/5/nodes/0/link_state |
| e003346 | 19:58:01.681949 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/5/nodes/1/flags |
| e003347 | 19:58:01.681949 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/5/nodes/1/primary_id |
| e003348 | 19:58:01.681949 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/13/snapshot/views/5/nodes/1/slots |
| e003349 | 19:58:01.681949 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/5/nodes/1/link_state |
| e003350 | 19:58:01.681949 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/5/nodes/2/flags |
| e003351 | 19:58:01.681949 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/13/snapshot/views/5/nodes/2/primary_id |
| e003352 | 19:58:01.681949 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/5/nodes/2/slots |
| e003353 | 19:58:01.681949 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/5/nodes/2/link_state |
| e003354 | 19:58:01.681949 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/5/nodes/3/flags |
| e003355 | 19:58:01.681949 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/5/nodes/3/primary_id |
| e003356 | 19:58:01.681949 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/13/snapshot/views/5/nodes/3/slots |
| e003357 | 19:58:01.681949 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/5/nodes/3/link_state |
| e003358 | 19:58:01.681949 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/5/nodes/4/flags |
| e003359 | 19:58:01.681949 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/13/snapshot/views/5/nodes/4/primary_id |
| e003360 | 19:58:01.681949 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/5/nodes/4/slots |
| e003361 | 19:58:01.681949 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/5/nodes/4/link_state |
| e003362 | 19:58:01.681949 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/5/nodes/5/flags |
| e003363 | 19:58:01.681949 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/5/nodes/5/primary_id |
| e003364 | 19:58:01.681949 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/13/snapshot/views/5/nodes/5/slots |
| e003365 | 19:58:01.681949 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/5/nodes/5/link_state |
| e003366 | 19:58:01.681949 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/13/snapshot/views/5/persistence/aof_enabled |
| e003367 | 19:58:01.681949 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/13/snapshot/views/5/persistence/loading |
| e003368 | 19:58:01.681949 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/13/snapshot/views/5/persistence/aof_last_write_status |
| e003369 | 19:58:01.681949 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/13/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e003370 | 19:58:01.681949 | node-6 | admin | process_live | missing_from_schema | `` | /samples/13/snapshot/views/5/process_live |
| e003371 | 19:58:01.681957 | node-3 | node_client | reachable | valid | `true` | /samples/13/snapshot/views/2/reachable |
| e003372 | 19:58:01.681957 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/13/snapshot/views/2/replication/role |
| e003373 | 19:58:01.681957 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/13/snapshot/views/2/replication/connected_slaves |
| e003374 | 19:58:01.681957 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/13/snapshot/views/2/replication/master_link_status |
| e003375 | 19:58:01.681957 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/13/snapshot/views/2/replication/master_sync_in_progress |
| e003376 | 19:58:01.681957 | node-3 | info_replication | replication.master_repl_offset | valid | `100120` | /samples/13/snapshot/views/2/replication/master_repl_offset |
| e003377 | 19:58:01.681957 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/13/snapshot/views/2/replication/slave_repl_offset |
| e003378 | 19:58:01.681957 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/13/snapshot/views/2/redis_id |
| e003379 | 19:58:01.681957 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/13/snapshot/views/2/coverage |
| e003380 | 19:58:01.681957 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/2/nodes/0/flags |
| e003381 | 19:58:01.681957 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/13/snapshot/views/2/nodes/0/primary_id |
| e003382 | 19:58:01.681957 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/2/nodes/0/slots |
| e003383 | 19:58:01.681957 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/2/nodes/0/link_state |
| e003384 | 19:58:01.681957 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/13/snapshot/views/2/nodes/1/flags |
| e003385 | 19:58:01.681957 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/2/nodes/1/primary_id |
| e003386 | 19:58:01.681957 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/13/snapshot/views/2/nodes/1/slots |
| e003387 | 19:58:01.681957 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/2/nodes/1/link_state |
| e003388 | 19:58:01.681957 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/2/nodes/2/flags |
| e003389 | 19:58:01.681957 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/13/snapshot/views/2/nodes/2/primary_id |
| e003390 | 19:58:01.681957 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/2/nodes/2/slots |
| e003391 | 19:58:01.681957 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/2/nodes/2/link_state |
| e003392 | 19:58:01.681957 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/2/nodes/3/flags |
| e003393 | 19:58:01.681957 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/2/nodes/3/primary_id |
| e003394 | 19:58:01.681957 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/13/snapshot/views/2/nodes/3/slots |
| e003395 | 19:58:01.681957 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/2/nodes/3/link_state |
| e003396 | 19:58:01.681957 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/2/nodes/4/flags |
| e003397 | 19:58:01.681957 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/2/nodes/4/primary_id |
| e003398 | 19:58:01.681957 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/13/snapshot/views/2/nodes/4/slots |
| e003399 | 19:58:01.681957 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/2/nodes/4/link_state |
| e003400 | 19:58:01.681957 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/2/nodes/5/flags |
| e003401 | 19:58:01.681957 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/13/snapshot/views/2/nodes/5/primary_id |
| e003402 | 19:58:01.681957 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/2/nodes/5/slots |
| e003403 | 19:58:01.681957 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/2/nodes/5/link_state |
| e003404 | 19:58:01.681957 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/13/snapshot/views/2/persistence/aof_enabled |
| e003405 | 19:58:01.681957 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/13/snapshot/views/2/persistence/loading |
| e003406 | 19:58:01.681957 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/13/snapshot/views/2/persistence/aof_last_write_status |
| e003407 | 19:58:01.681957 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/13/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e003408 | 19:58:01.681957 | node-3 | admin | process_live | missing_from_schema | `` | /samples/13/snapshot/views/2/process_live |
| e003409 | 19:58:01.681966 | node-4 | node_client | reachable | valid | `true` | /samples/13/snapshot/views/3/reachable |
| e003410 | 19:58:01.681966 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/13/snapshot/views/3/replication/role |
| e003411 | 19:58:01.681966 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/13/snapshot/views/3/replication/connected_slaves |
| e003412 | 19:58:01.681966 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/13/snapshot/views/3/replication/master_link_status |
| e003413 | 19:58:01.681966 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/13/snapshot/views/3/replication/master_sync_in_progress |
| e003414 | 19:58:01.681966 | node-4 | info_replication | replication.master_repl_offset | valid | `100120` | /samples/13/snapshot/views/3/replication/master_repl_offset |
| e003415 | 19:58:01.681966 | node-4 | info_replication | replication.slave_repl_offset | valid | `100120` | /samples/13/snapshot/views/3/replication/slave_repl_offset |
| e003416 | 19:58:01.681966 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/13/snapshot/views/3/redis_id |
| e003417 | 19:58:01.681966 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/13/snapshot/views/3/coverage |
| e003418 | 19:58:01.681966 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/3/nodes/0/flags |
| e003419 | 19:58:01.681966 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/3/nodes/0/primary_id |
| e003420 | 19:58:01.681966 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/13/snapshot/views/3/nodes/0/slots |
| e003421 | 19:58:01.681966 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/3/nodes/0/link_state |
| e003422 | 19:58:01.681966 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/3/nodes/1/flags |
| e003423 | 19:58:01.681966 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/13/snapshot/views/3/nodes/1/primary_id |
| e003424 | 19:58:01.681966 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/3/nodes/1/slots |
| e003425 | 19:58:01.681966 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/3/nodes/1/link_state |
| e003426 | 19:58:01.681966 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/3/nodes/2/flags |
| e003427 | 19:58:01.681966 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/13/snapshot/views/3/nodes/2/primary_id |
| e003428 | 19:58:01.681966 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/3/nodes/2/slots |
| e003429 | 19:58:01.681966 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/3/nodes/2/link_state |
| e003430 | 19:58:01.681966 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/13/snapshot/views/3/nodes/3/flags |
| e003431 | 19:58:01.681966 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/13/snapshot/views/3/nodes/3/primary_id |
| e003432 | 19:58:01.681966 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/3/nodes/3/slots |
| e003433 | 19:58:01.681966 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/3/nodes/3/link_state |
| e003434 | 19:58:01.681966 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/3/nodes/4/flags |
| e003435 | 19:58:01.681966 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/3/nodes/4/primary_id |
| e003436 | 19:58:01.681966 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/13/snapshot/views/3/nodes/4/slots |
| e003437 | 19:58:01.681966 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/3/nodes/4/link_state |
| e003438 | 19:58:01.681966 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/3/nodes/5/flags |
| e003439 | 19:58:01.681966 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/3/nodes/5/primary_id |
| e003440 | 19:58:01.681966 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/13/snapshot/views/3/nodes/5/slots |
| e003441 | 19:58:01.681966 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/3/nodes/5/link_state |
| e003442 | 19:58:01.681966 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/13/snapshot/views/3/persistence/aof_enabled |
| e003443 | 19:58:01.681966 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/13/snapshot/views/3/persistence/loading |
| e003444 | 19:58:01.681966 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/13/snapshot/views/3/persistence/aof_last_write_status |
| e003445 | 19:58:01.681966 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/13/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e003446 | 19:58:01.681966 | node-4 | admin | process_live | missing_from_schema | `` | /samples/13/snapshot/views/3/process_live |
| e003447 | 19:58:01.682056 | node-5 | node_client | reachable | valid | `true` | /samples/13/snapshot/views/4/reachable |
| e003448 | 19:58:01.682056 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/13/snapshot/views/4/replication/role |
| e003449 | 19:58:01.682056 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/13/snapshot/views/4/replication/connected_slaves |
| e003450 | 19:58:01.682056 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/13/snapshot/views/4/replication/master_link_status |
| e003451 | 19:58:01.682056 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/13/snapshot/views/4/replication/master_sync_in_progress |
| e003452 | 19:58:01.682056 | node-5 | info_replication | replication.master_repl_offset | valid | `100400` | /samples/13/snapshot/views/4/replication/master_repl_offset |
| e003453 | 19:58:01.682056 | node-5 | info_replication | replication.slave_repl_offset | valid | `100400` | /samples/13/snapshot/views/4/replication/slave_repl_offset |
| e003454 | 19:58:01.682056 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /samples/13/snapshot/views/4/redis_id |
| e003455 | 19:58:01.682056 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/13/snapshot/views/4/coverage |
| e003456 | 19:58:01.682056 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/4/nodes/0/flags |
| e003457 | 19:58:01.682056 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/4/nodes/0/primary_id |
| e003458 | 19:58:01.682056 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/13/snapshot/views/4/nodes/0/slots |
| e003459 | 19:58:01.682056 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/4/nodes/0/link_state |
| e003460 | 19:58:01.682056 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/4/nodes/1/flags |
| e003461 | 19:58:01.682056 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/13/snapshot/views/4/nodes/1/primary_id |
| e003462 | 19:58:01.682056 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/4/nodes/1/slots |
| e003463 | 19:58:01.682056 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/4/nodes/1/link_state |
| e003464 | 19:58:01.682056 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/13/snapshot/views/4/nodes/2/flags |
| e003465 | 19:58:01.682056 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/13/snapshot/views/4/nodes/2/primary_id |
| e003466 | 19:58:01.682056 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/4/nodes/2/slots |
| e003467 | 19:58:01.682056 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/4/nodes/2/link_state |
| e003468 | 19:58:01.682056 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/4/nodes/3/flags |
| e003469 | 19:58:01.682056 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/4/nodes/3/primary_id |
| e003470 | 19:58:01.682056 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/13/snapshot/views/4/nodes/3/slots |
| e003471 | 19:58:01.682056 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/4/nodes/3/link_state |
| e003472 | 19:58:01.682056 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/13/snapshot/views/4/nodes/4/flags |
| e003473 | 19:58:01.682056 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/13/snapshot/views/4/nodes/4/primary_id |
| e003474 | 19:58:01.682056 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/13/snapshot/views/4/nodes/4/slots |
| e003475 | 19:58:01.682056 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/4/nodes/4/link_state |
| e003476 | 19:58:01.682056 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/13/snapshot/views/4/nodes/5/flags |
| e003477 | 19:58:01.682056 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/13/snapshot/views/4/nodes/5/primary_id |
| e003478 | 19:58:01.682056 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/13/snapshot/views/4/nodes/5/slots |
| e003479 | 19:58:01.682056 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/13/snapshot/views/4/nodes/5/link_state |
| e003480 | 19:58:01.682056 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/13/snapshot/views/4/persistence/aof_enabled |
| e003481 | 19:58:01.682056 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/13/snapshot/views/4/persistence/loading |
| e003482 | 19:58:01.682056 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/13/snapshot/views/4/persistence/aof_last_write_status |
| e003483 | 19:58:01.682056 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/13/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e003484 | 19:58:01.682056 | node-5 | admin | process_live | missing_from_schema | `` | /samples/13/snapshot/views/4/process_live |
| e003485 | 19:58:01.823703 |  | client_probe | client_success | valid | `true` | /samples/13/probe/successful |
| e003486 | 19:58:01.825159 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/13/write_policy_probe/policy |
| e003487 | 19:58:01.825159 |  | ack_probe | ack.successful | valid | `true` | /samples/13/write_policy_probe/successful |
| e003488 | 19:58:01.825159 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/13/write_policy_probe/requested_replicas |
| e003489 | 19:58:01.825159 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/13/write_policy_probe/set_acknowledged |
| e003490 | 19:58:01.825159 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          1
        ]` | /samples/13/write_policy_probe/acknowledgements |
| e003491 | 19:58:01.825159 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/13/write_policy_probe/acknowledgement_observed |
| e003492 | 19:58:01.825159 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/13/write_policy_probe/waitaof_requested_local |
| e003493 | 19:58:01.825159 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/13/write_policy_probe/waitaof_requested_replicas |
| e003494 | 19:58:01.825159 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/13/write_policy_probe/waitaof_local_acknowledgements |
| e003495 | 19:58:01.825159 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/13/write_policy_probe/waitaof_replica_acknowledgements |
| e003496 | 19:58:01.825159 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/13/write_policy_probe/waitaof_acknowledgement_observed |
| e003497 | 19:58:02.338054 | node-4 | node_client | reachable | valid | `true` | /samples/14/snapshot/views/3/reachable |
| e003498 | 19:58:02.338054 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/14/snapshot/views/3/replication/role |
| e003499 | 19:58:02.338054 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/14/snapshot/views/3/replication/connected_slaves |
| e003500 | 19:58:02.338054 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/14/snapshot/views/3/replication/master_link_status |
| e003501 | 19:58:02.338054 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/14/snapshot/views/3/replication/master_sync_in_progress |
| e003502 | 19:58:02.338054 | node-4 | info_replication | replication.master_repl_offset | valid | `105291` | /samples/14/snapshot/views/3/replication/master_repl_offset |
| e003503 | 19:58:02.338054 | node-4 | info_replication | replication.slave_repl_offset | valid | `105291` | /samples/14/snapshot/views/3/replication/slave_repl_offset |
| e003504 | 19:58:02.338054 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/14/snapshot/views/3/redis_id |
| e003505 | 19:58:02.338054 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/14/snapshot/views/3/coverage |
| e003506 | 19:58:02.338054 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/3/nodes/0/flags |
| e003507 | 19:58:02.338054 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/3/nodes/0/primary_id |
| e003508 | 19:58:02.338054 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/14/snapshot/views/3/nodes/0/slots |
| e003509 | 19:58:02.338054 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/3/nodes/0/link_state |
| e003510 | 19:58:02.338054 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/3/nodes/1/flags |
| e003511 | 19:58:02.338054 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/14/snapshot/views/3/nodes/1/primary_id |
| e003512 | 19:58:02.338054 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/3/nodes/1/slots |
| e003513 | 19:58:02.338054 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/3/nodes/1/link_state |
| e003514 | 19:58:02.338054 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/3/nodes/2/flags |
| e003515 | 19:58:02.338054 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/14/snapshot/views/3/nodes/2/primary_id |
| e003516 | 19:58:02.338054 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/3/nodes/2/slots |
| e003517 | 19:58:02.338054 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/3/nodes/2/link_state |
| e003518 | 19:58:02.338054 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/14/snapshot/views/3/nodes/3/flags |
| e003519 | 19:58:02.338054 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/14/snapshot/views/3/nodes/3/primary_id |
| e003520 | 19:58:02.338054 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/3/nodes/3/slots |
| e003521 | 19:58:02.338054 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/3/nodes/3/link_state |
| e003522 | 19:58:02.338054 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/3/nodes/4/flags |
| e003523 | 19:58:02.338054 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/3/nodes/4/primary_id |
| e003524 | 19:58:02.338054 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/14/snapshot/views/3/nodes/4/slots |
| e003525 | 19:58:02.338054 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/3/nodes/4/link_state |
| e003526 | 19:58:02.338054 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/3/nodes/5/flags |
| e003527 | 19:58:02.338054 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/3/nodes/5/primary_id |
| e003528 | 19:58:02.338054 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/14/snapshot/views/3/nodes/5/slots |
| e003529 | 19:58:02.338054 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/3/nodes/5/link_state |
| e003530 | 19:58:02.338054 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/14/snapshot/views/3/persistence/aof_enabled |
| e003531 | 19:58:02.338054 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/14/snapshot/views/3/persistence/loading |
| e003532 | 19:58:02.338054 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/14/snapshot/views/3/persistence/aof_last_write_status |
| e003533 | 19:58:02.338054 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/14/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e003534 | 19:58:02.338054 | node-4 | admin | process_live | missing_from_schema | `` | /samples/14/snapshot/views/3/process_live |
| e003535 | 19:58:02.338057 | node-3 | node_client | reachable | valid | `true` | /samples/14/snapshot/views/2/reachable |
| e003536 | 19:58:02.338057 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/14/snapshot/views/2/replication/role |
| e003537 | 19:58:02.338057 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/14/snapshot/views/2/replication/connected_slaves |
| e003538 | 19:58:02.338057 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/14/snapshot/views/2/replication/master_link_status |
| e003539 | 19:58:02.338057 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/14/snapshot/views/2/replication/master_sync_in_progress |
| e003540 | 19:58:02.338057 | node-3 | info_replication | replication.master_repl_offset | valid | `105291` | /samples/14/snapshot/views/2/replication/master_repl_offset |
| e003541 | 19:58:02.338057 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/14/snapshot/views/2/replication/slave_repl_offset |
| e003542 | 19:58:02.338057 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/14/snapshot/views/2/redis_id |
| e003543 | 19:58:02.338057 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/14/snapshot/views/2/coverage |
| e003544 | 19:58:02.338057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/2/nodes/0/flags |
| e003545 | 19:58:02.338057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/14/snapshot/views/2/nodes/0/primary_id |
| e003546 | 19:58:02.338057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/2/nodes/0/slots |
| e003547 | 19:58:02.338057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/2/nodes/0/link_state |
| e003548 | 19:58:02.338057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/14/snapshot/views/2/nodes/1/flags |
| e003549 | 19:58:02.338057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/2/nodes/1/primary_id |
| e003550 | 19:58:02.338057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/14/snapshot/views/2/nodes/1/slots |
| e003551 | 19:58:02.338057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/2/nodes/1/link_state |
| e003552 | 19:58:02.338057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/2/nodes/2/flags |
| e003553 | 19:58:02.338057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/14/snapshot/views/2/nodes/2/primary_id |
| e003554 | 19:58:02.338057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/2/nodes/2/slots |
| e003555 | 19:58:02.338057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/2/nodes/2/link_state |
| e003556 | 19:58:02.338057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/2/nodes/3/flags |
| e003557 | 19:58:02.338057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/2/nodes/3/primary_id |
| e003558 | 19:58:02.338057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/14/snapshot/views/2/nodes/3/slots |
| e003559 | 19:58:02.338057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/2/nodes/3/link_state |
| e003560 | 19:58:02.338057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/2/nodes/4/flags |
| e003561 | 19:58:02.338057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/2/nodes/4/primary_id |
| e003562 | 19:58:02.338057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/14/snapshot/views/2/nodes/4/slots |
| e003563 | 19:58:02.338057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/2/nodes/4/link_state |
| e003564 | 19:58:02.338057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/2/nodes/5/flags |
| e003565 | 19:58:02.338057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/14/snapshot/views/2/nodes/5/primary_id |
| e003566 | 19:58:02.338057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/2/nodes/5/slots |
| e003567 | 19:58:02.338057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/2/nodes/5/link_state |
| e003568 | 19:58:02.338057 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/14/snapshot/views/2/persistence/aof_enabled |
| e003569 | 19:58:02.338057 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/14/snapshot/views/2/persistence/loading |
| e003570 | 19:58:02.338057 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/14/snapshot/views/2/persistence/aof_last_write_status |
| e003571 | 19:58:02.338057 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/14/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e003572 | 19:58:02.338057 | node-3 | admin | process_live | missing_from_schema | `` | /samples/14/snapshot/views/2/process_live |
| e003573 | 19:58:02.338057 | node-5 | node_client | reachable | valid | `true` | /samples/14/snapshot/views/4/reachable |
| e003574 | 19:58:02.338057 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/14/snapshot/views/4/replication/role |
| e003575 | 19:58:02.338057 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/14/snapshot/views/4/replication/connected_slaves |
| e003576 | 19:58:02.338057 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/14/snapshot/views/4/replication/master_link_status |
| e003577 | 19:58:02.338057 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/14/snapshot/views/4/replication/master_sync_in_progress |
| e003578 | 19:58:02.338057 | node-5 | info_replication | replication.master_repl_offset | valid | `105571` | /samples/14/snapshot/views/4/replication/master_repl_offset |
| e003579 | 19:58:02.338057 | node-5 | info_replication | replication.slave_repl_offset | valid | `105571` | /samples/14/snapshot/views/4/replication/slave_repl_offset |
| e003580 | 19:58:02.338057 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /samples/14/snapshot/views/4/redis_id |
| e003581 | 19:58:02.338057 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/14/snapshot/views/4/coverage |
| e003582 | 19:58:02.338057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/4/nodes/0/flags |
| e003583 | 19:58:02.338057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/4/nodes/0/primary_id |
| e003584 | 19:58:02.338057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/14/snapshot/views/4/nodes/0/slots |
| e003585 | 19:58:02.338057 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/4/nodes/0/link_state |
| e003586 | 19:58:02.338057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/4/nodes/1/flags |
| e003587 | 19:58:02.338057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/14/snapshot/views/4/nodes/1/primary_id |
| e003588 | 19:58:02.338057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/4/nodes/1/slots |
| e003589 | 19:58:02.338057 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/4/nodes/1/link_state |
| e003590 | 19:58:02.338057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/14/snapshot/views/4/nodes/2/flags |
| e003591 | 19:58:02.338057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/14/snapshot/views/4/nodes/2/primary_id |
| e003592 | 19:58:02.338057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/4/nodes/2/slots |
| e003593 | 19:58:02.338057 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/4/nodes/2/link_state |
| e003594 | 19:58:02.338057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/4/nodes/3/flags |
| e003595 | 19:58:02.338057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/4/nodes/3/primary_id |
| e003596 | 19:58:02.338057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/14/snapshot/views/4/nodes/3/slots |
| e003597 | 19:58:02.338057 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/4/nodes/3/link_state |
| e003598 | 19:58:02.338057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/4/nodes/4/flags |
| e003599 | 19:58:02.338057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/14/snapshot/views/4/nodes/4/primary_id |
| e003600 | 19:58:02.338057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/4/nodes/4/slots |
| e003601 | 19:58:02.338057 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/4/nodes/4/link_state |
| e003602 | 19:58:02.338057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/4/nodes/5/flags |
| e003603 | 19:58:02.338057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/4/nodes/5/primary_id |
| e003604 | 19:58:02.338057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/14/snapshot/views/4/nodes/5/slots |
| e003605 | 19:58:02.338057 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/4/nodes/5/link_state |
| e003606 | 19:58:02.338057 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/14/snapshot/views/4/persistence/aof_enabled |
| e003607 | 19:58:02.338057 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/14/snapshot/views/4/persistence/loading |
| e003608 | 19:58:02.338057 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/14/snapshot/views/4/persistence/aof_last_write_status |
| e003609 | 19:58:02.338057 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/14/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e003610 | 19:58:02.338057 | node-5 | admin | process_live | missing_from_schema | `` | /samples/14/snapshot/views/4/process_live |
| e003611 | 19:58:02.338078 | node-2 | node_client | reachable | valid | `true` | /samples/14/snapshot/views/1/reachable |
| e003612 | 19:58:02.338078 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/14/snapshot/views/1/replication/role |
| e003613 | 19:58:02.338078 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/14/snapshot/views/1/replication/connected_slaves |
| e003614 | 19:58:02.338078 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/14/snapshot/views/1/replication/master_link_status |
| e003615 | 19:58:02.338078 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/14/snapshot/views/1/replication/master_sync_in_progress |
| e003616 | 19:58:02.338078 | node-2 | info_replication | replication.master_repl_offset | valid | `105328` | /samples/14/snapshot/views/1/replication/master_repl_offset |
| e003617 | 19:58:02.338078 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/14/snapshot/views/1/replication/slave_repl_offset |
| e003618 | 19:58:02.338078 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/14/snapshot/views/1/redis_id |
| e003619 | 19:58:02.338078 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/14/snapshot/views/1/coverage |
| e003620 | 19:58:02.338078 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/1/nodes/0/flags |
| e003621 | 19:58:02.338078 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/1/nodes/0/primary_id |
| e003622 | 19:58:02.338078 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/14/snapshot/views/1/nodes/0/slots |
| e003623 | 19:58:02.338078 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/1/nodes/0/link_state |
| e003624 | 19:58:02.338078 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/1/nodes/1/flags |
| e003625 | 19:58:02.338078 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/14/snapshot/views/1/nodes/1/primary_id |
| e003626 | 19:58:02.338078 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/1/nodes/1/slots |
| e003627 | 19:58:02.338078 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/1/nodes/1/link_state |
| e003628 | 19:58:02.338078 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/1/nodes/2/flags |
| e003629 | 19:58:02.338078 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/14/snapshot/views/1/nodes/2/primary_id |
| e003630 | 19:58:02.338078 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/1/nodes/2/slots |
| e003631 | 19:58:02.338078 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/1/nodes/2/link_state |
| e003632 | 19:58:02.338078 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/1/nodes/3/flags |
| e003633 | 19:58:02.338078 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/14/snapshot/views/1/nodes/3/primary_id |
| e003634 | 19:58:02.338078 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/1/nodes/3/slots |
| e003635 | 19:58:02.338078 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/1/nodes/3/link_state |
| e003636 | 19:58:02.338078 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/14/snapshot/views/1/nodes/4/flags |
| e003637 | 19:58:02.338078 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/1/nodes/4/primary_id |
| e003638 | 19:58:02.338078 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/14/snapshot/views/1/nodes/4/slots |
| e003639 | 19:58:02.338078 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/1/nodes/4/link_state |
| e003640 | 19:58:02.338078 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/1/nodes/5/flags |
| e003641 | 19:58:02.338078 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/1/nodes/5/primary_id |
| e003642 | 19:58:02.338078 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/14/snapshot/views/1/nodes/5/slots |
| e003643 | 19:58:02.338078 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/1/nodes/5/link_state |
| e003644 | 19:58:02.338078 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/14/snapshot/views/1/persistence/aof_enabled |
| e003645 | 19:58:02.338078 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/14/snapshot/views/1/persistence/loading |
| e003646 | 19:58:02.338078 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/14/snapshot/views/1/persistence/aof_last_write_status |
| e003647 | 19:58:02.338078 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/14/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e003648 | 19:58:02.338078 | node-2 | admin | process_live | missing_from_schema | `` | /samples/14/snapshot/views/1/process_live |
| e003649 | 19:58:02.338079 | node-1 | node_client | reachable | valid | `true` | /samples/14/snapshot/views/0/reachable |
| e003650 | 19:58:02.338079 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/14/snapshot/views/0/replication/role |
| e003651 | 19:58:02.338079 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/14/snapshot/views/0/replication/connected_slaves |
| e003652 | 19:58:02.338079 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/14/snapshot/views/0/replication/master_link_status |
| e003653 | 19:58:02.338079 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/14/snapshot/views/0/replication/master_sync_in_progress |
| e003654 | 19:58:02.338079 | node-1 | info_replication | replication.master_repl_offset | valid | `105571` | /samples/14/snapshot/views/0/replication/master_repl_offset |
| e003655 | 19:58:02.338079 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/14/snapshot/views/0/replication/slave_repl_offset |
| e003656 | 19:58:02.338079 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/14/snapshot/views/0/redis_id |
| e003657 | 19:58:02.338079 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/14/snapshot/views/0/coverage |
| e003658 | 19:58:02.338079 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/0/nodes/0/flags |
| e003659 | 19:58:02.338079 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/14/snapshot/views/0/nodes/0/primary_id |
| e003660 | 19:58:02.338079 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/0/nodes/0/slots |
| e003661 | 19:58:02.338079 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/0/nodes/0/link_state |
| e003662 | 19:58:02.338079 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/0/nodes/1/flags |
| e003663 | 19:58:02.338079 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/14/snapshot/views/0/nodes/1/primary_id |
| e003664 | 19:58:02.338079 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/0/nodes/1/slots |
| e003665 | 19:58:02.338079 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/0/nodes/1/link_state |
| e003666 | 19:58:02.338079 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/0/nodes/2/flags |
| e003667 | 19:58:02.338079 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/0/nodes/2/primary_id |
| e003668 | 19:58:02.338079 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/14/snapshot/views/0/nodes/2/slots |
| e003669 | 19:58:02.338079 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/0/nodes/2/link_state |
| e003670 | 19:58:02.338079 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/14/snapshot/views/0/nodes/3/flags |
| e003671 | 19:58:02.338079 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/0/nodes/3/primary_id |
| e003672 | 19:58:02.338079 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/14/snapshot/views/0/nodes/3/slots |
| e003673 | 19:58:02.338079 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/0/nodes/3/link_state |
| e003674 | 19:58:02.338079 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/0/nodes/4/flags |
| e003675 | 19:58:02.338079 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/14/snapshot/views/0/nodes/4/primary_id |
| e003676 | 19:58:02.338079 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/0/nodes/4/slots |
| e003677 | 19:58:02.338079 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/0/nodes/4/link_state |
| e003678 | 19:58:02.338079 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/0/nodes/5/flags |
| e003679 | 19:58:02.338079 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/0/nodes/5/primary_id |
| e003680 | 19:58:02.338079 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/14/snapshot/views/0/nodes/5/slots |
| e003681 | 19:58:02.338079 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/0/nodes/5/link_state |
| e003682 | 19:58:02.338079 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/14/snapshot/views/0/persistence/aof_enabled |
| e003683 | 19:58:02.338079 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/14/snapshot/views/0/persistence/loading |
| e003684 | 19:58:02.338079 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/14/snapshot/views/0/persistence/aof_last_write_status |
| e003685 | 19:58:02.338079 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/14/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e003686 | 19:58:02.338079 | node-1 | admin | process_live | missing_from_schema | `` | /samples/14/snapshot/views/0/process_live |
| e003687 | 19:58:02.338086 | node-6 | node_client | reachable | valid | `true` | /samples/14/snapshot/views/5/reachable |
| e003688 | 19:58:02.338086 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/14/snapshot/views/5/replication/role |
| e003689 | 19:58:02.338086 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/14/snapshot/views/5/replication/connected_slaves |
| e003690 | 19:58:02.338086 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/14/snapshot/views/5/replication/master_link_status |
| e003691 | 19:58:02.338086 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/14/snapshot/views/5/replication/master_sync_in_progress |
| e003692 | 19:58:02.338086 | node-6 | info_replication | replication.master_repl_offset | valid | `105328` | /samples/14/snapshot/views/5/replication/master_repl_offset |
| e003693 | 19:58:02.338086 | node-6 | info_replication | replication.slave_repl_offset | valid | `105328` | /samples/14/snapshot/views/5/replication/slave_repl_offset |
| e003694 | 19:58:02.338086 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/14/snapshot/views/5/redis_id |
| e003695 | 19:58:02.338086 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/14/snapshot/views/5/coverage |
| e003696 | 19:58:02.338086 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/14/snapshot/views/5/nodes/0/flags |
| e003697 | 19:58:02.338086 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/14/snapshot/views/5/nodes/0/primary_id |
| e003698 | 19:58:02.338086 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/5/nodes/0/slots |
| e003699 | 19:58:02.338086 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/5/nodes/0/link_state |
| e003700 | 19:58:02.338086 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/5/nodes/1/flags |
| e003701 | 19:58:02.338086 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/5/nodes/1/primary_id |
| e003702 | 19:58:02.338086 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/14/snapshot/views/5/nodes/1/slots |
| e003703 | 19:58:02.338086 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/5/nodes/1/link_state |
| e003704 | 19:58:02.338086 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/5/nodes/2/flags |
| e003705 | 19:58:02.338086 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/14/snapshot/views/5/nodes/2/primary_id |
| e003706 | 19:58:02.338086 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/5/nodes/2/slots |
| e003707 | 19:58:02.338086 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/5/nodes/2/link_state |
| e003708 | 19:58:02.338086 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/5/nodes/3/flags |
| e003709 | 19:58:02.338086 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/5/nodes/3/primary_id |
| e003710 | 19:58:02.338086 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/14/snapshot/views/5/nodes/3/slots |
| e003711 | 19:58:02.338086 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/5/nodes/3/link_state |
| e003712 | 19:58:02.338086 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/14/snapshot/views/5/nodes/4/flags |
| e003713 | 19:58:02.338086 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/14/snapshot/views/5/nodes/4/primary_id |
| e003714 | 19:58:02.338086 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/14/snapshot/views/5/nodes/4/slots |
| e003715 | 19:58:02.338086 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/5/nodes/4/link_state |
| e003716 | 19:58:02.338086 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/14/snapshot/views/5/nodes/5/flags |
| e003717 | 19:58:02.338086 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/14/snapshot/views/5/nodes/5/primary_id |
| e003718 | 19:58:02.338086 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/14/snapshot/views/5/nodes/5/slots |
| e003719 | 19:58:02.338086 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/14/snapshot/views/5/nodes/5/link_state |
| e003720 | 19:58:02.338086 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/14/snapshot/views/5/persistence/aof_enabled |
| e003721 | 19:58:02.338086 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/14/snapshot/views/5/persistence/loading |
| e003722 | 19:58:02.338086 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/14/snapshot/views/5/persistence/aof_last_write_status |
| e003723 | 19:58:02.338086 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/14/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e003724 | 19:58:02.338086 | node-6 | admin | process_live | missing_from_schema | `` | /samples/14/snapshot/views/5/process_live |
| e003725 | 19:58:02.381014 |  | client_window | client.attempts | valid | `170` | /metrics/windows/13/attempts |
| e003726 | 19:58:02.381014 |  | client_window | client.errors | valid | `0` | /metrics/windows/13/errors |
| e003727 | 19:58:02.381014 |  | client_window | client.successes | valid | `170` | /metrics/windows/13/successes |
| e003728 | 19:58:02.381014 |  | client_window | client.p95_ms | valid | `1.449031` | /metrics/windows/13/p95_ms |
| e003729 | 19:58:02.480432 |  | client_probe | client_success | valid | `true` | /samples/14/probe/successful |
| e003730 | 19:58:02.481852 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/14/write_policy_probe/policy |
| e003731 | 19:58:02.481852 |  | ack_probe | ack.successful | valid | `true` | /samples/14/write_policy_probe/successful |
| e003732 | 19:58:02.481852 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/14/write_policy_probe/requested_replicas |
| e003733 | 19:58:02.481852 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/14/write_policy_probe/set_acknowledged |
| e003734 | 19:58:02.481852 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          1
        ]` | /samples/14/write_policy_probe/acknowledgements |
| e003735 | 19:58:02.481852 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/14/write_policy_probe/acknowledgement_observed |
| e003736 | 19:58:02.481852 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/14/write_policy_probe/waitaof_requested_local |
| e003737 | 19:58:02.481852 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/14/write_policy_probe/waitaof_requested_replicas |
| e003738 | 19:58:02.481852 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/14/write_policy_probe/waitaof_local_acknowledgements |
| e003739 | 19:58:02.481852 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/14/write_policy_probe/waitaof_replica_acknowledgements |
| e003740 | 19:58:02.481852 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/14/write_policy_probe/waitaof_acknowledgement_observed |
| e003741 | 19:58:02.684 | node-2 | node_client | reachable | valid | `true` | /samples/15/snapshot/views/1/reachable |
| e003742 | 19:58:02.684 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/15/snapshot/views/1/replication/role |
| e003743 | 19:58:02.684 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/15/snapshot/views/1/replication/connected_slaves |
| e003744 | 19:58:02.684 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/15/snapshot/views/1/replication/master_link_status |
| e003745 | 19:58:02.684 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/15/snapshot/views/1/replication/master_sync_in_progress |
| e003746 | 19:58:02.684 | node-2 | info_replication | replication.master_repl_offset | valid | `108679` | /samples/15/snapshot/views/1/replication/master_repl_offset |
| e003747 | 19:58:02.684 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/15/snapshot/views/1/replication/slave_repl_offset |
| e003748 | 19:58:02.684 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/15/snapshot/views/1/redis_id |
| e003749 | 19:58:02.684 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/15/snapshot/views/1/coverage |
| e003750 | 19:58:02.684 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/1/nodes/0/flags |
| e003751 | 19:58:02.684 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/1/nodes/0/primary_id |
| e003752 | 19:58:02.684 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/15/snapshot/views/1/nodes/0/slots |
| e003753 | 19:58:02.684 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/1/nodes/0/link_state |
| e003754 | 19:58:02.684 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/1/nodes/1/flags |
| e003755 | 19:58:02.684 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/15/snapshot/views/1/nodes/1/primary_id |
| e003756 | 19:58:02.684 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/1/nodes/1/slots |
| e003757 | 19:58:02.684 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/1/nodes/1/link_state |
| e003758 | 19:58:02.684 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/1/nodes/2/flags |
| e003759 | 19:58:02.684 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/15/snapshot/views/1/nodes/2/primary_id |
| e003760 | 19:58:02.684 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/1/nodes/2/slots |
| e003761 | 19:58:02.684 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/1/nodes/2/link_state |
| e003762 | 19:58:02.684 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/1/nodes/3/flags |
| e003763 | 19:58:02.684 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/15/snapshot/views/1/nodes/3/primary_id |
| e003764 | 19:58:02.684 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/1/nodes/3/slots |
| e003765 | 19:58:02.684 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/1/nodes/3/link_state |
| e003766 | 19:58:02.684 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/15/snapshot/views/1/nodes/4/flags |
| e003767 | 19:58:02.684 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/1/nodes/4/primary_id |
| e003768 | 19:58:02.684 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/15/snapshot/views/1/nodes/4/slots |
| e003769 | 19:58:02.684 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/1/nodes/4/link_state |
| e003770 | 19:58:02.684 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/1/nodes/5/flags |
| e003771 | 19:58:02.684 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/1/nodes/5/primary_id |
| e003772 | 19:58:02.684 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/15/snapshot/views/1/nodes/5/slots |
| e003773 | 19:58:02.684 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/1/nodes/5/link_state |
| e003774 | 19:58:02.684 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/15/snapshot/views/1/persistence/aof_enabled |
| e003775 | 19:58:02.684 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/15/snapshot/views/1/persistence/loading |
| e003776 | 19:58:02.684 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/15/snapshot/views/1/persistence/aof_last_write_status |
| e003777 | 19:58:02.684 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/15/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e003778 | 19:58:02.684 | node-2 | admin | process_live | missing_from_schema | `` | /samples/15/snapshot/views/1/process_live |
| e003779 | 19:58:02.684007 | node-4 | node_client | reachable | valid | `true` | /samples/15/snapshot/views/3/reachable |
| e003780 | 19:58:02.684007 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/15/snapshot/views/3/replication/role |
| e003781 | 19:58:02.684007 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/15/snapshot/views/3/replication/connected_slaves |
| e003782 | 19:58:02.684007 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/15/snapshot/views/3/replication/master_link_status |
| e003783 | 19:58:02.684007 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/15/snapshot/views/3/replication/master_sync_in_progress |
| e003784 | 19:58:02.684007 | node-4 | info_replication | replication.master_repl_offset | valid | `108382` | /samples/15/snapshot/views/3/replication/master_repl_offset |
| e003785 | 19:58:02.684007 | node-4 | info_replication | replication.slave_repl_offset | valid | `108382` | /samples/15/snapshot/views/3/replication/slave_repl_offset |
| e003786 | 19:58:02.684007 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/15/snapshot/views/3/redis_id |
| e003787 | 19:58:02.684007 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/15/snapshot/views/3/coverage |
| e003788 | 19:58:02.684007 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/3/nodes/0/flags |
| e003789 | 19:58:02.684007 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/3/nodes/0/primary_id |
| e003790 | 19:58:02.684007 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/15/snapshot/views/3/nodes/0/slots |
| e003791 | 19:58:02.684007 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/3/nodes/0/link_state |
| e003792 | 19:58:02.684007 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/3/nodes/1/flags |
| e003793 | 19:58:02.684007 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/15/snapshot/views/3/nodes/1/primary_id |
| e003794 | 19:58:02.684007 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/3/nodes/1/slots |
| e003795 | 19:58:02.684007 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/3/nodes/1/link_state |
| e003796 | 19:58:02.684007 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/3/nodes/2/flags |
| e003797 | 19:58:02.684007 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/15/snapshot/views/3/nodes/2/primary_id |
| e003798 | 19:58:02.684007 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/3/nodes/2/slots |
| e003799 | 19:58:02.684007 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/3/nodes/2/link_state |
| e003800 | 19:58:02.684007 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/15/snapshot/views/3/nodes/3/flags |
| e003801 | 19:58:02.684007 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/15/snapshot/views/3/nodes/3/primary_id |
| e003802 | 19:58:02.684007 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/3/nodes/3/slots |
| e003803 | 19:58:02.684007 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/3/nodes/3/link_state |
| e003804 | 19:58:02.684007 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/3/nodes/4/flags |
| e003805 | 19:58:02.684007 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/3/nodes/4/primary_id |
| e003806 | 19:58:02.684007 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/15/snapshot/views/3/nodes/4/slots |
| e003807 | 19:58:02.684007 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/3/nodes/4/link_state |
| e003808 | 19:58:02.684007 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/3/nodes/5/flags |
| e003809 | 19:58:02.684007 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/3/nodes/5/primary_id |
| e003810 | 19:58:02.684007 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/15/snapshot/views/3/nodes/5/slots |
| e003811 | 19:58:02.684007 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/3/nodes/5/link_state |
| e003812 | 19:58:02.684007 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/15/snapshot/views/3/persistence/aof_enabled |
| e003813 | 19:58:02.684007 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/15/snapshot/views/3/persistence/loading |
| e003814 | 19:58:02.684007 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/15/snapshot/views/3/persistence/aof_last_write_status |
| e003815 | 19:58:02.684007 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/15/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e003816 | 19:58:02.684007 | node-4 | admin | process_live | missing_from_schema | `` | /samples/15/snapshot/views/3/process_live |
| e003817 | 19:58:02.684012 | node-3 | node_client | reachable | valid | `true` | /samples/15/snapshot/views/2/reachable |
| e003818 | 19:58:02.684012 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/15/snapshot/views/2/replication/role |
| e003819 | 19:58:02.684012 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/15/snapshot/views/2/replication/connected_slaves |
| e003820 | 19:58:02.684012 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/15/snapshot/views/2/replication/master_link_status |
| e003821 | 19:58:02.684012 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/15/snapshot/views/2/replication/master_sync_in_progress |
| e003822 | 19:58:02.684012 | node-3 | info_replication | replication.master_repl_offset | valid | `108382` | /samples/15/snapshot/views/2/replication/master_repl_offset |
| e003823 | 19:58:02.684012 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/15/snapshot/views/2/replication/slave_repl_offset |
| e003824 | 19:58:02.684012 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/15/snapshot/views/2/redis_id |
| e003825 | 19:58:02.684012 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/15/snapshot/views/2/coverage |
| e003826 | 19:58:02.684012 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/2/nodes/0/flags |
| e003827 | 19:58:02.684012 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/15/snapshot/views/2/nodes/0/primary_id |
| e003828 | 19:58:02.684012 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/2/nodes/0/slots |
| e003829 | 19:58:02.684012 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/2/nodes/0/link_state |
| e003830 | 19:58:02.684012 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/15/snapshot/views/2/nodes/1/flags |
| e003831 | 19:58:02.684012 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/2/nodes/1/primary_id |
| e003832 | 19:58:02.684012 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/15/snapshot/views/2/nodes/1/slots |
| e003833 | 19:58:02.684012 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/2/nodes/1/link_state |
| e003834 | 19:58:02.684012 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/2/nodes/2/flags |
| e003835 | 19:58:02.684012 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/15/snapshot/views/2/nodes/2/primary_id |
| e003836 | 19:58:02.684012 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/2/nodes/2/slots |
| e003837 | 19:58:02.684012 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/2/nodes/2/link_state |
| e003838 | 19:58:02.684012 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/2/nodes/3/flags |
| e003839 | 19:58:02.684012 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/2/nodes/3/primary_id |
| e003840 | 19:58:02.684012 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/15/snapshot/views/2/nodes/3/slots |
| e003841 | 19:58:02.684012 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/2/nodes/3/link_state |
| e003842 | 19:58:02.684012 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/2/nodes/4/flags |
| e003843 | 19:58:02.684012 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/2/nodes/4/primary_id |
| e003844 | 19:58:02.684012 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/15/snapshot/views/2/nodes/4/slots |
| e003845 | 19:58:02.684012 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/2/nodes/4/link_state |
| e003846 | 19:58:02.684012 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/2/nodes/5/flags |
| e003847 | 19:58:02.684012 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/15/snapshot/views/2/nodes/5/primary_id |
| e003848 | 19:58:02.684012 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/2/nodes/5/slots |
| e003849 | 19:58:02.684012 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/2/nodes/5/link_state |
| e003850 | 19:58:02.684012 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/15/snapshot/views/2/persistence/aof_enabled |
| e003851 | 19:58:02.684012 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/15/snapshot/views/2/persistence/loading |
| e003852 | 19:58:02.684012 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/15/snapshot/views/2/persistence/aof_last_write_status |
| e003853 | 19:58:02.684012 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/15/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e003854 | 19:58:02.684012 | node-3 | admin | process_live | missing_from_schema | `` | /samples/15/snapshot/views/2/process_live |
| e003855 | 19:58:02.68402 | node-1 | node_client | reachable | valid | `true` | /samples/15/snapshot/views/0/reachable |
| e003856 | 19:58:02.68402 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/15/snapshot/views/0/replication/role |
| e003857 | 19:58:02.68402 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/15/snapshot/views/0/replication/connected_slaves |
| e003858 | 19:58:02.68402 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/15/snapshot/views/0/replication/master_link_status |
| e003859 | 19:58:02.68402 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/15/snapshot/views/0/replication/master_sync_in_progress |
| e003860 | 19:58:02.68402 | node-1 | info_replication | replication.master_repl_offset | valid | `108662` | /samples/15/snapshot/views/0/replication/master_repl_offset |
| e003861 | 19:58:02.68402 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/15/snapshot/views/0/replication/slave_repl_offset |
| e003862 | 19:58:02.68402 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/15/snapshot/views/0/redis_id |
| e003863 | 19:58:02.68402 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/15/snapshot/views/0/coverage |
| e003864 | 19:58:02.68402 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/0/nodes/0/flags |
| e003865 | 19:58:02.68402 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/15/snapshot/views/0/nodes/0/primary_id |
| e003866 | 19:58:02.68402 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/0/nodes/0/slots |
| e003867 | 19:58:02.68402 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/0/nodes/0/link_state |
| e003868 | 19:58:02.68402 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/0/nodes/1/flags |
| e003869 | 19:58:02.68402 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/15/snapshot/views/0/nodes/1/primary_id |
| e003870 | 19:58:02.68402 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/0/nodes/1/slots |
| e003871 | 19:58:02.68402 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/0/nodes/1/link_state |
| e003872 | 19:58:02.68402 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/0/nodes/2/flags |
| e003873 | 19:58:02.68402 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/0/nodes/2/primary_id |
| e003874 | 19:58:02.68402 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/15/snapshot/views/0/nodes/2/slots |
| e003875 | 19:58:02.68402 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/0/nodes/2/link_state |
| e003876 | 19:58:02.68402 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/15/snapshot/views/0/nodes/3/flags |
| e003877 | 19:58:02.68402 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/0/nodes/3/primary_id |
| e003878 | 19:58:02.68402 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/15/snapshot/views/0/nodes/3/slots |
| e003879 | 19:58:02.68402 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/0/nodes/3/link_state |
| e003880 | 19:58:02.68402 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/0/nodes/4/flags |
| e003881 | 19:58:02.68402 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/15/snapshot/views/0/nodes/4/primary_id |
| e003882 | 19:58:02.68402 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/0/nodes/4/slots |
| e003883 | 19:58:02.68402 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/0/nodes/4/link_state |
| e003884 | 19:58:02.68402 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/0/nodes/5/flags |
| e003885 | 19:58:02.68402 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/0/nodes/5/primary_id |
| e003886 | 19:58:02.68402 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/15/snapshot/views/0/nodes/5/slots |
| e003887 | 19:58:02.68402 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/0/nodes/5/link_state |
| e003888 | 19:58:02.68402 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/15/snapshot/views/0/persistence/aof_enabled |
| e003889 | 19:58:02.68402 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/15/snapshot/views/0/persistence/loading |
| e003890 | 19:58:02.68402 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/15/snapshot/views/0/persistence/aof_last_write_status |
| e003891 | 19:58:02.68402 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/15/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e003892 | 19:58:02.68402 | node-1 | admin | process_live | missing_from_schema | `` | /samples/15/snapshot/views/0/process_live |
| e003893 | 19:58:02.684029 | node-5 | node_client | reachable | valid | `true` | /samples/15/snapshot/views/4/reachable |
| e003894 | 19:58:02.684029 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/15/snapshot/views/4/replication/role |
| e003895 | 19:58:02.684029 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/15/snapshot/views/4/replication/connected_slaves |
| e003896 | 19:58:02.684029 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/15/snapshot/views/4/replication/master_link_status |
| e003897 | 19:58:02.684029 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/15/snapshot/views/4/replication/master_sync_in_progress |
| e003898 | 19:58:02.684029 | node-5 | info_replication | replication.master_repl_offset | valid | `108662` | /samples/15/snapshot/views/4/replication/master_repl_offset |
| e003899 | 19:58:02.684029 | node-5 | info_replication | replication.slave_repl_offset | valid | `108662` | /samples/15/snapshot/views/4/replication/slave_repl_offset |
| e003900 | 19:58:02.684029 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /samples/15/snapshot/views/4/redis_id |
| e003901 | 19:58:02.684029 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/15/snapshot/views/4/coverage |
| e003902 | 19:58:02.684029 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/4/nodes/0/flags |
| e003903 | 19:58:02.684029 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/4/nodes/0/primary_id |
| e003904 | 19:58:02.684029 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/15/snapshot/views/4/nodes/0/slots |
| e003905 | 19:58:02.684029 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/4/nodes/0/link_state |
| e003906 | 19:58:02.684029 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/4/nodes/1/flags |
| e003907 | 19:58:02.684029 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/15/snapshot/views/4/nodes/1/primary_id |
| e003908 | 19:58:02.684029 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/4/nodes/1/slots |
| e003909 | 19:58:02.684029 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/4/nodes/1/link_state |
| e003910 | 19:58:02.684029 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/15/snapshot/views/4/nodes/2/flags |
| e003911 | 19:58:02.684029 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/15/snapshot/views/4/nodes/2/primary_id |
| e003912 | 19:58:02.684029 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/4/nodes/2/slots |
| e003913 | 19:58:02.684029 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/4/nodes/2/link_state |
| e003914 | 19:58:02.684029 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/4/nodes/3/flags |
| e003915 | 19:58:02.684029 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/4/nodes/3/primary_id |
| e003916 | 19:58:02.684029 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/15/snapshot/views/4/nodes/3/slots |
| e003917 | 19:58:02.684029 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/4/nodes/3/link_state |
| e003918 | 19:58:02.684029 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/4/nodes/4/flags |
| e003919 | 19:58:02.684029 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/15/snapshot/views/4/nodes/4/primary_id |
| e003920 | 19:58:02.684029 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/4/nodes/4/slots |
| e003921 | 19:58:02.684029 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/4/nodes/4/link_state |
| e003922 | 19:58:02.684029 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/4/nodes/5/flags |
| e003923 | 19:58:02.684029 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/4/nodes/5/primary_id |
| e003924 | 19:58:02.684029 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/15/snapshot/views/4/nodes/5/slots |
| e003925 | 19:58:02.684029 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/4/nodes/5/link_state |
| e003926 | 19:58:02.684029 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/15/snapshot/views/4/persistence/aof_enabled |
| e003927 | 19:58:02.684029 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/15/snapshot/views/4/persistence/loading |
| e003928 | 19:58:02.684029 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/15/snapshot/views/4/persistence/aof_last_write_status |
| e003929 | 19:58:02.684029 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/15/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e003930 | 19:58:02.684029 | node-5 | admin | process_live | missing_from_schema | `` | /samples/15/snapshot/views/4/process_live |
| e003931 | 19:58:02.684037 | node-6 | node_client | reachable | valid | `true` | /samples/15/snapshot/views/5/reachable |
| e003932 | 19:58:02.684037 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/15/snapshot/views/5/replication/role |
| e003933 | 19:58:02.684037 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/15/snapshot/views/5/replication/connected_slaves |
| e003934 | 19:58:02.684037 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/15/snapshot/views/5/replication/master_link_status |
| e003935 | 19:58:02.684037 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/15/snapshot/views/5/replication/master_sync_in_progress |
| e003936 | 19:58:02.684037 | node-6 | info_replication | replication.master_repl_offset | valid | `108679` | /samples/15/snapshot/views/5/replication/master_repl_offset |
| e003937 | 19:58:02.684037 | node-6 | info_replication | replication.slave_repl_offset | valid | `108679` | /samples/15/snapshot/views/5/replication/slave_repl_offset |
| e003938 | 19:58:02.684037 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/15/snapshot/views/5/redis_id |
| e003939 | 19:58:02.684037 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/15/snapshot/views/5/coverage |
| e003940 | 19:58:02.684037 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/15/snapshot/views/5/nodes/0/flags |
| e003941 | 19:58:02.684037 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/15/snapshot/views/5/nodes/0/primary_id |
| e003942 | 19:58:02.684037 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/5/nodes/0/slots |
| e003943 | 19:58:02.684037 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/5/nodes/0/link_state |
| e003944 | 19:58:02.684037 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/5/nodes/1/flags |
| e003945 | 19:58:02.684037 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/5/nodes/1/primary_id |
| e003946 | 19:58:02.684037 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/15/snapshot/views/5/nodes/1/slots |
| e003947 | 19:58:02.684037 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/5/nodes/1/link_state |
| e003948 | 19:58:02.684037 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/5/nodes/2/flags |
| e003949 | 19:58:02.684037 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/15/snapshot/views/5/nodes/2/primary_id |
| e003950 | 19:58:02.684037 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/5/nodes/2/slots |
| e003951 | 19:58:02.684037 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/5/nodes/2/link_state |
| e003952 | 19:58:02.684037 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/5/nodes/3/flags |
| e003953 | 19:58:02.684037 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/5/nodes/3/primary_id |
| e003954 | 19:58:02.684037 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/15/snapshot/views/5/nodes/3/slots |
| e003955 | 19:58:02.684037 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/5/nodes/3/link_state |
| e003956 | 19:58:02.684037 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/15/snapshot/views/5/nodes/4/flags |
| e003957 | 19:58:02.684037 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/15/snapshot/views/5/nodes/4/primary_id |
| e003958 | 19:58:02.684037 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/15/snapshot/views/5/nodes/4/slots |
| e003959 | 19:58:02.684037 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/5/nodes/4/link_state |
| e003960 | 19:58:02.684037 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/15/snapshot/views/5/nodes/5/flags |
| e003961 | 19:58:02.684037 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/15/snapshot/views/5/nodes/5/primary_id |
| e003962 | 19:58:02.684037 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/15/snapshot/views/5/nodes/5/slots |
| e003963 | 19:58:02.684037 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/15/snapshot/views/5/nodes/5/link_state |
| e003964 | 19:58:02.684037 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/15/snapshot/views/5/persistence/aof_enabled |
| e003965 | 19:58:02.684037 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/15/snapshot/views/5/persistence/loading |
| e003966 | 19:58:02.684037 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/15/snapshot/views/5/persistence/aof_last_write_status |
| e003967 | 19:58:02.684037 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/15/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e003968 | 19:58:02.684037 | node-6 | admin | process_live | missing_from_schema | `` | /samples/15/snapshot/views/5/process_live |
| e003969 | 19:58:02.823684 |  | client_probe | client_success | valid | `true` | /samples/15/probe/successful |
| e003970 | 19:58:02.939459 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/15/write_policy_probe/policy |
| e003971 | 19:58:02.939459 |  | ack_probe | ack.successful | valid | `true` | /samples/15/write_policy_probe/successful |
| e003972 | 19:58:02.939459 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/15/write_policy_probe/requested_replicas |
| e003973 | 19:58:02.939459 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/15/write_policy_probe/set_acknowledged |
| e003974 | 19:58:02.939459 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          1
        ]` | /samples/15/write_policy_probe/acknowledgements |
| e003975 | 19:58:02.939459 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/15/write_policy_probe/acknowledgement_observed |
| e003976 | 19:58:02.939459 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/15/write_policy_probe/waitaof_requested_local |
| e003977 | 19:58:02.939459 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/15/write_policy_probe/waitaof_requested_replicas |
| e003978 | 19:58:02.939459 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/15/write_policy_probe/waitaof_local_acknowledgements |
| e003979 | 19:58:02.939459 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/15/write_policy_probe/waitaof_replica_acknowledgements |
| e003980 | 19:58:02.939459 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/15/write_policy_probe/waitaof_acknowledgement_observed |
| e003981 | 19:58:03.337924 | node-6 | node_client | reachable | valid | `true` | /samples/16/snapshot/views/5/reachable |
| e003982 | 19:58:03.337924 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/16/snapshot/views/5/replication/role |
| e003983 | 19:58:03.337924 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/16/snapshot/views/5/replication/connected_slaves |
| e003984 | 19:58:03.337924 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/16/snapshot/views/5/replication/master_link_status |
| e003985 | 19:58:03.337924 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/16/snapshot/views/5/replication/master_sync_in_progress |
| e003986 | 19:58:03.337924 | node-6 | info_replication | replication.master_repl_offset | valid | `113590` | /samples/16/snapshot/views/5/replication/master_repl_offset |
| e003987 | 19:58:03.337924 | node-6 | info_replication | replication.slave_repl_offset | valid | `113590` | /samples/16/snapshot/views/5/replication/slave_repl_offset |
| e003988 | 19:58:03.337924 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /samples/16/snapshot/views/5/redis_id |
| e003989 | 19:58:03.337924 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/16/snapshot/views/5/coverage |
| e003990 | 19:58:03.337924 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/16/snapshot/views/5/nodes/0/flags |
| e003991 | 19:58:03.337924 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/16/snapshot/views/5/nodes/0/primary_id |
| e003992 | 19:58:03.337924 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/5/nodes/0/slots |
| e003993 | 19:58:03.337924 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/5/nodes/0/link_state |
| e003994 | 19:58:03.337924 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/5/nodes/1/flags |
| e003995 | 19:58:03.337924 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/5/nodes/1/primary_id |
| e003996 | 19:58:03.337924 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/16/snapshot/views/5/nodes/1/slots |
| e003997 | 19:58:03.337924 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/5/nodes/1/link_state |
| e003998 | 19:58:03.337924 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/5/nodes/2/flags |
| e003999 | 19:58:03.337924 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/16/snapshot/views/5/nodes/2/primary_id |
| e004000 | 19:58:03.337924 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/5/nodes/2/slots |
| e004001 | 19:58:03.337924 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/5/nodes/2/link_state |
| e004002 | 19:58:03.337924 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/5/nodes/3/flags |
| e004003 | 19:58:03.337924 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/5/nodes/3/primary_id |
| e004004 | 19:58:03.337924 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/16/snapshot/views/5/nodes/3/slots |
| e004005 | 19:58:03.337924 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/5/nodes/3/link_state |
| e004006 | 19:58:03.337924 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/5/nodes/4/flags |
| e004007 | 19:58:03.337924 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/16/snapshot/views/5/nodes/4/primary_id |
| e004008 | 19:58:03.337924 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/5/nodes/4/slots |
| e004009 | 19:58:03.337924 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/5/nodes/4/link_state |
| e004010 | 19:58:03.337924 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/5/nodes/5/flags |
| e004011 | 19:58:03.337924 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/5/nodes/5/primary_id |
| e004012 | 19:58:03.337924 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/16/snapshot/views/5/nodes/5/slots |
| e004013 | 19:58:03.337924 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/5/nodes/5/link_state |
| e004014 | 19:58:03.337924 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/16/snapshot/views/5/persistence/aof_enabled |
| e004015 | 19:58:03.337924 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/16/snapshot/views/5/persistence/loading |
| e004016 | 19:58:03.337924 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/16/snapshot/views/5/persistence/aof_last_write_status |
| e004017 | 19:58:03.337924 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/16/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e004018 | 19:58:03.337924 | node-6 | admin | process_live | missing_from_schema | `` | /samples/16/snapshot/views/5/process_live |
| e004019 | 19:58:03.337924 | node-6 | node_client | reachable | valid | `true` | /final/views/5/reachable |
| e004020 | 19:58:03.337924 | node-6 | info_replication | replication.role | valid | `"slave"` | /final/views/5/replication/role |
| e004021 | 19:58:03.337924 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /final/views/5/replication/connected_slaves |
| e004022 | 19:58:03.337924 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /final/views/5/replication/master_link_status |
| e004023 | 19:58:03.337924 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /final/views/5/replication/master_sync_in_progress |
| e004024 | 19:58:03.337924 | node-6 | info_replication | replication.master_repl_offset | valid | `113590` | /final/views/5/replication/master_repl_offset |
| e004025 | 19:58:03.337924 | node-6 | info_replication | replication.slave_repl_offset | valid | `113590` | /final/views/5/replication/slave_repl_offset |
| e004026 | 19:58:03.337924 | node-6 | cluster_nodes | redis_id | valid | `"bc00717ad0825a914118d066bbb0e87129689ab3"` | /final/views/5/redis_id |
| e004027 | 19:58:03.337924 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /final/views/5/coverage |
| e004028 | 19:58:03.337924 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /final/views/5/nodes/0/flags |
| e004029 | 19:58:03.337924 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /final/views/5/nodes/0/primary_id |
| e004030 | 19:58:03.337924 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /final/views/5/nodes/0/slots |
| e004031 | 19:58:03.337924 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/0/link_state |
| e004032 | 19:58:03.337924 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /final/views/5/nodes/1/flags |
| e004033 | 19:58:03.337924 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /final/views/5/nodes/1/primary_id |
| e004034 | 19:58:03.337924 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/5/nodes/1/slots |
| e004035 | 19:58:03.337924 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/1/link_state |
| e004036 | 19:58:03.337924 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/5/nodes/2/flags |
| e004037 | 19:58:03.337924 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /final/views/5/nodes/2/primary_id |
| e004038 | 19:58:03.337924 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /final/views/5/nodes/2/slots |
| e004039 | 19:58:03.337924 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/2/link_state |
| e004040 | 19:58:03.337924 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /final/views/5/nodes/3/flags |
| e004041 | 19:58:03.337924 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /final/views/5/nodes/3/primary_id |
| e004042 | 19:58:03.337924 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/5/nodes/3/slots |
| e004043 | 19:58:03.337924 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/3/link_state |
| e004044 | 19:58:03.337924 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/5/nodes/4/flags |
| e004045 | 19:58:03.337924 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /final/views/5/nodes/4/primary_id |
| e004046 | 19:58:03.337924 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.slots | missing_from_schema | `` | /final/views/5/nodes/4/slots |
| e004047 | 19:58:03.337924 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/4/link_state |
| e004048 | 19:58:03.337924 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /final/views/5/nodes/5/flags |
| e004049 | 19:58:03.337924 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /final/views/5/nodes/5/primary_id |
| e004050 | 19:58:03.337924 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/5/nodes/5/slots |
| e004051 | 19:58:03.337924 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/5/link_state |
| e004052 | 19:58:03.337924 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/5/persistence/aof_enabled |
| e004053 | 19:58:03.337924 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /final/views/5/persistence/loading |
| e004054 | 19:58:03.337924 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/5/persistence/aof_last_write_status |
| e004055 | 19:58:03.337924 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/5/persistence/aof_rewrite_in_progress |
| e004056 | 19:58:03.337924 | node-6 | admin | process_live | missing_from_schema | `` | /final/views/5/process_live |
| e004057 | 19:58:03.337926 | node-1 | node_client | reachable | valid | `true` | /samples/16/snapshot/views/0/reachable |
| e004058 | 19:58:03.337926 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/16/snapshot/views/0/replication/role |
| e004059 | 19:58:03.337926 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/16/snapshot/views/0/replication/connected_slaves |
| e004060 | 19:58:03.337926 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/16/snapshot/views/0/replication/master_link_status |
| e004061 | 19:58:03.337926 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/16/snapshot/views/0/replication/master_sync_in_progress |
| e004062 | 19:58:03.337926 | node-1 | info_replication | replication.master_repl_offset | valid | `113833` | /samples/16/snapshot/views/0/replication/master_repl_offset |
| e004063 | 19:58:03.337926 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/16/snapshot/views/0/replication/slave_repl_offset |
| e004064 | 19:58:03.337926 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/16/snapshot/views/0/redis_id |
| e004065 | 19:58:03.337926 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/16/snapshot/views/0/coverage |
| e004066 | 19:58:03.337926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/0/nodes/0/flags |
| e004067 | 19:58:03.337926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/16/snapshot/views/0/nodes/0/primary_id |
| e004068 | 19:58:03.337926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/0/nodes/0/slots |
| e004069 | 19:58:03.337926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/0/nodes/0/link_state |
| e004070 | 19:58:03.337926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/0/nodes/1/flags |
| e004071 | 19:58:03.337926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/16/snapshot/views/0/nodes/1/primary_id |
| e004072 | 19:58:03.337926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/0/nodes/1/slots |
| e004073 | 19:58:03.337926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/0/nodes/1/link_state |
| e004074 | 19:58:03.337926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/0/nodes/2/flags |
| e004075 | 19:58:03.337926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/0/nodes/2/primary_id |
| e004076 | 19:58:03.337926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/16/snapshot/views/0/nodes/2/slots |
| e004077 | 19:58:03.337926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/0/nodes/2/link_state |
| e004078 | 19:58:03.337926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/16/snapshot/views/0/nodes/3/flags |
| e004079 | 19:58:03.337926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/0/nodes/3/primary_id |
| e004080 | 19:58:03.337926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/16/snapshot/views/0/nodes/3/slots |
| e004081 | 19:58:03.337926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/0/nodes/3/link_state |
| e004082 | 19:58:03.337926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/0/nodes/4/flags |
| e004083 | 19:58:03.337926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/16/snapshot/views/0/nodes/4/primary_id |
| e004084 | 19:58:03.337926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/0/nodes/4/slots |
| e004085 | 19:58:03.337926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/0/nodes/4/link_state |
| e004086 | 19:58:03.337926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/0/nodes/5/flags |
| e004087 | 19:58:03.337926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/0/nodes/5/primary_id |
| e004088 | 19:58:03.337926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/16/snapshot/views/0/nodes/5/slots |
| e004089 | 19:58:03.337926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/0/nodes/5/link_state |
| e004090 | 19:58:03.337926 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/16/snapshot/views/0/persistence/aof_enabled |
| e004091 | 19:58:03.337926 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/16/snapshot/views/0/persistence/loading |
| e004092 | 19:58:03.337926 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/16/snapshot/views/0/persistence/aof_last_write_status |
| e004093 | 19:58:03.337926 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/16/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e004094 | 19:58:03.337926 | node-1 | admin | process_live | missing_from_schema | `` | /samples/16/snapshot/views/0/process_live |
| e004095 | 19:58:03.337926 | node-1 | node_client | reachable | valid | `true` | /final/views/0/reachable |
| e004096 | 19:58:03.337926 | node-1 | info_replication | replication.role | valid | `"master"` | /final/views/0/replication/role |
| e004097 | 19:58:03.337926 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /final/views/0/replication/connected_slaves |
| e004098 | 19:58:03.337926 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /final/views/0/replication/master_link_status |
| e004099 | 19:58:03.337926 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /final/views/0/replication/master_sync_in_progress |
| e004100 | 19:58:03.337926 | node-1 | info_replication | replication.master_repl_offset | valid | `113833` | /final/views/0/replication/master_repl_offset |
| e004101 | 19:58:03.337926 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /final/views/0/replication/slave_repl_offset |
| e004102 | 19:58:03.337926 | node-1 | cluster_nodes | redis_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /final/views/0/redis_id |
| e004103 | 19:58:03.337926 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /final/views/0/coverage |
| e004104 | 19:58:03.337926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/0/nodes/0/flags |
| e004105 | 19:58:03.337926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /final/views/0/nodes/0/primary_id |
| e004106 | 19:58:03.337926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /final/views/0/nodes/0/slots |
| e004107 | 19:58:03.337926 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/0/link_state |
| e004108 | 19:58:03.337926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/0/nodes/1/flags |
| e004109 | 19:58:03.337926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /final/views/0/nodes/1/primary_id |
| e004110 | 19:58:03.337926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /final/views/0/nodes/1/slots |
| e004111 | 19:58:03.337926 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/1/link_state |
| e004112 | 19:58:03.337926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.flags | valid | `[
              "master"
            ]` | /final/views/0/nodes/2/flags |
| e004113 | 19:58:03.337926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /final/views/0/nodes/2/primary_id |
| e004114 | 19:58:03.337926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/0/nodes/2/slots |
| e004115 | 19:58:03.337926 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/2/link_state |
| e004116 | 19:58:03.337926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /final/views/0/nodes/3/flags |
| e004117 | 19:58:03.337926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /final/views/0/nodes/3/primary_id |
| e004118 | 19:58:03.337926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/0/nodes/3/slots |
| e004119 | 19:58:03.337926 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/3/link_state |
| e004120 | 19:58:03.337926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/0/nodes/4/flags |
| e004121 | 19:58:03.337926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /final/views/0/nodes/4/primary_id |
| e004122 | 19:58:03.337926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.slots | missing_from_schema | `` | /final/views/0/nodes/4/slots |
| e004123 | 19:58:03.337926 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/4/link_state |
| e004124 | 19:58:03.337926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.flags | valid | `[
              "master"
            ]` | /final/views/0/nodes/5/flags |
| e004125 | 19:58:03.337926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /final/views/0/nodes/5/primary_id |
| e004126 | 19:58:03.337926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/0/nodes/5/slots |
| e004127 | 19:58:03.337926 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/5/link_state |
| e004128 | 19:58:03.337926 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/0/persistence/aof_enabled |
| e004129 | 19:58:03.337926 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /final/views/0/persistence/loading |
| e004130 | 19:58:03.337926 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/0/persistence/aof_last_write_status |
| e004131 | 19:58:03.337926 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/0/persistence/aof_rewrite_in_progress |
| e004132 | 19:58:03.337926 | node-1 | admin | process_live | missing_from_schema | `` | /final/views/0/process_live |
| e004133 | 19:58:03.337929 | node-4 | node_client | reachable | valid | `true` | /samples/16/snapshot/views/3/reachable |
| e004134 | 19:58:03.337929 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/16/snapshot/views/3/replication/role |
| e004135 | 19:58:03.337929 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/16/snapshot/views/3/replication/connected_slaves |
| e004136 | 19:58:03.337929 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/16/snapshot/views/3/replication/master_link_status |
| e004137 | 19:58:03.337929 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/16/snapshot/views/3/replication/master_sync_in_progress |
| e004138 | 19:58:03.337929 | node-4 | info_replication | replication.master_repl_offset | valid | `113553` | /samples/16/snapshot/views/3/replication/master_repl_offset |
| e004139 | 19:58:03.337929 | node-4 | info_replication | replication.slave_repl_offset | valid | `113553` | /samples/16/snapshot/views/3/replication/slave_repl_offset |
| e004140 | 19:58:03.337929 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /samples/16/snapshot/views/3/redis_id |
| e004141 | 19:58:03.337929 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/16/snapshot/views/3/coverage |
| e004142 | 19:58:03.337929 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/3/nodes/0/flags |
| e004143 | 19:58:03.337929 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/3/nodes/0/primary_id |
| e004144 | 19:58:03.337929 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/16/snapshot/views/3/nodes/0/slots |
| e004145 | 19:58:03.337929 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/3/nodes/0/link_state |
| e004146 | 19:58:03.337929 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/3/nodes/1/flags |
| e004147 | 19:58:03.337929 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/16/snapshot/views/3/nodes/1/primary_id |
| e004148 | 19:58:03.337929 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/3/nodes/1/slots |
| e004149 | 19:58:03.337929 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/3/nodes/1/link_state |
| e004150 | 19:58:03.337929 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/3/nodes/2/flags |
| e004151 | 19:58:03.337929 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/16/snapshot/views/3/nodes/2/primary_id |
| e004152 | 19:58:03.337929 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/3/nodes/2/slots |
| e004153 | 19:58:03.337929 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/3/nodes/2/link_state |
| e004154 | 19:58:03.337929 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/16/snapshot/views/3/nodes/3/flags |
| e004155 | 19:58:03.337929 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/16/snapshot/views/3/nodes/3/primary_id |
| e004156 | 19:58:03.337929 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/3/nodes/3/slots |
| e004157 | 19:58:03.337929 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/3/nodes/3/link_state |
| e004158 | 19:58:03.337929 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/3/nodes/4/flags |
| e004159 | 19:58:03.337929 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/3/nodes/4/primary_id |
| e004160 | 19:58:03.337929 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/16/snapshot/views/3/nodes/4/slots |
| e004161 | 19:58:03.337929 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/3/nodes/4/link_state |
| e004162 | 19:58:03.337929 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/3/nodes/5/flags |
| e004163 | 19:58:03.337929 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/3/nodes/5/primary_id |
| e004164 | 19:58:03.337929 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/16/snapshot/views/3/nodes/5/slots |
| e004165 | 19:58:03.337929 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/3/nodes/5/link_state |
| e004166 | 19:58:03.337929 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/16/snapshot/views/3/persistence/aof_enabled |
| e004167 | 19:58:03.337929 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/16/snapshot/views/3/persistence/loading |
| e004168 | 19:58:03.337929 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/16/snapshot/views/3/persistence/aof_last_write_status |
| e004169 | 19:58:03.337929 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/16/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e004170 | 19:58:03.337929 | node-4 | admin | process_live | missing_from_schema | `` | /samples/16/snapshot/views/3/process_live |
| e004171 | 19:58:03.337929 | node-4 | node_client | reachable | valid | `true` | /final/views/3/reachable |
| e004172 | 19:58:03.337929 | node-4 | info_replication | replication.role | valid | `"slave"` | /final/views/3/replication/role |
| e004173 | 19:58:03.337929 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /final/views/3/replication/connected_slaves |
| e004174 | 19:58:03.337929 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /final/views/3/replication/master_link_status |
| e004175 | 19:58:03.337929 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /final/views/3/replication/master_sync_in_progress |
| e004176 | 19:58:03.337929 | node-4 | info_replication | replication.master_repl_offset | valid | `113553` | /final/views/3/replication/master_repl_offset |
| e004177 | 19:58:03.337929 | node-4 | info_replication | replication.slave_repl_offset | valid | `113553` | /final/views/3/replication/slave_repl_offset |
| e004178 | 19:58:03.337929 | node-4 | cluster_nodes | redis_id | valid | `"5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf"` | /final/views/3/redis_id |
| e004179 | 19:58:03.337929 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /final/views/3/coverage |
| e004180 | 19:58:03.337929 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /final/views/3/nodes/0/flags |
| e004181 | 19:58:03.337929 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /final/views/3/nodes/0/primary_id |
| e004182 | 19:58:03.337929 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/3/nodes/0/slots |
| e004183 | 19:58:03.337929 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/0/link_state |
| e004184 | 19:58:03.337929 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/3/nodes/1/flags |
| e004185 | 19:58:03.337929 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /final/views/3/nodes/1/primary_id |
| e004186 | 19:58:03.337929 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /final/views/3/nodes/1/slots |
| e004187 | 19:58:03.337929 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/1/link_state |
| e004188 | 19:58:03.337929 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/3/nodes/2/flags |
| e004189 | 19:58:03.337929 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /final/views/3/nodes/2/primary_id |
| e004190 | 19:58:03.337929 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /final/views/3/nodes/2/slots |
| e004191 | 19:58:03.337929 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/2/link_state |
| e004192 | 19:58:03.337929 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /final/views/3/nodes/3/flags |
| e004193 | 19:58:03.337929 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /final/views/3/nodes/3/primary_id |
| e004194 | 19:58:03.337929 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.slots | missing_from_schema | `` | /final/views/3/nodes/3/slots |
| e004195 | 19:58:03.337929 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/3/link_state |
| e004196 | 19:58:03.337929 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /final/views/3/nodes/4/flags |
| e004197 | 19:58:03.337929 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /final/views/3/nodes/4/primary_id |
| e004198 | 19:58:03.337929 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/3/nodes/4/slots |
| e004199 | 19:58:03.337929 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/4/link_state |
| e004200 | 19:58:03.337929 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /final/views/3/nodes/5/flags |
| e004201 | 19:58:03.337929 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /final/views/3/nodes/5/primary_id |
| e004202 | 19:58:03.337929 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/3/nodes/5/slots |
| e004203 | 19:58:03.337929 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/5/link_state |
| e004204 | 19:58:03.337929 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/3/persistence/aof_enabled |
| e004205 | 19:58:03.337929 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /final/views/3/persistence/loading |
| e004206 | 19:58:03.337929 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/3/persistence/aof_last_write_status |
| e004207 | 19:58:03.337929 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/3/persistence/aof_rewrite_in_progress |
| e004208 | 19:58:03.337929 | node-4 | admin | process_live | missing_from_schema | `` | /final/views/3/process_live |
| e004209 | 19:58:03.338005 | node-2 | node_client | reachable | valid | `true` | /samples/16/snapshot/views/1/reachable |
| e004210 | 19:58:03.338005 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/16/snapshot/views/1/replication/role |
| e004211 | 19:58:03.338005 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/16/snapshot/views/1/replication/connected_slaves |
| e004212 | 19:58:03.338005 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/16/snapshot/views/1/replication/master_link_status |
| e004213 | 19:58:03.338005 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/16/snapshot/views/1/replication/master_sync_in_progress |
| e004214 | 19:58:03.338005 | node-2 | info_replication | replication.master_repl_offset | valid | `113590` | /samples/16/snapshot/views/1/replication/master_repl_offset |
| e004215 | 19:58:03.338005 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/16/snapshot/views/1/replication/slave_repl_offset |
| e004216 | 19:58:03.338005 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/16/snapshot/views/1/redis_id |
| e004217 | 19:58:03.338005 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/16/snapshot/views/1/coverage |
| e004218 | 19:58:03.338005 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/1/nodes/0/flags |
| e004219 | 19:58:03.338005 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/1/nodes/0/primary_id |
| e004220 | 19:58:03.338005 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/16/snapshot/views/1/nodes/0/slots |
| e004221 | 19:58:03.338005 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/1/nodes/0/link_state |
| e004222 | 19:58:03.338005 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/1/nodes/1/flags |
| e004223 | 19:58:03.338005 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/16/snapshot/views/1/nodes/1/primary_id |
| e004224 | 19:58:03.338005 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/1/nodes/1/slots |
| e004225 | 19:58:03.338005 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/1/nodes/1/link_state |
| e004226 | 19:58:03.338005 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/1/nodes/2/flags |
| e004227 | 19:58:03.338005 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/16/snapshot/views/1/nodes/2/primary_id |
| e004228 | 19:58:03.338005 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/1/nodes/2/slots |
| e004229 | 19:58:03.338005 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/1/nodes/2/link_state |
| e004230 | 19:58:03.338005 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/1/nodes/3/flags |
| e004231 | 19:58:03.338005 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/16/snapshot/views/1/nodes/3/primary_id |
| e004232 | 19:58:03.338005 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/1/nodes/3/slots |
| e004233 | 19:58:03.338005 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/1/nodes/3/link_state |
| e004234 | 19:58:03.338005 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/16/snapshot/views/1/nodes/4/flags |
| e004235 | 19:58:03.338005 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/1/nodes/4/primary_id |
| e004236 | 19:58:03.338005 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/16/snapshot/views/1/nodes/4/slots |
| e004237 | 19:58:03.338005 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/1/nodes/4/link_state |
| e004238 | 19:58:03.338005 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/1/nodes/5/flags |
| e004239 | 19:58:03.338005 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/1/nodes/5/primary_id |
| e004240 | 19:58:03.338005 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/16/snapshot/views/1/nodes/5/slots |
| e004241 | 19:58:03.338005 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/1/nodes/5/link_state |
| e004242 | 19:58:03.338005 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/16/snapshot/views/1/persistence/aof_enabled |
| e004243 | 19:58:03.338005 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/16/snapshot/views/1/persistence/loading |
| e004244 | 19:58:03.338005 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/16/snapshot/views/1/persistence/aof_last_write_status |
| e004245 | 19:58:03.338005 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/16/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e004246 | 19:58:03.338005 | node-2 | admin | process_live | missing_from_schema | `` | /samples/16/snapshot/views/1/process_live |
| e004247 | 19:58:03.338005 | node-2 | node_client | reachable | valid | `true` | /final/views/1/reachable |
| e004248 | 19:58:03.338005 | node-2 | info_replication | replication.role | valid | `"master"` | /final/views/1/replication/role |
| e004249 | 19:58:03.338005 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /final/views/1/replication/connected_slaves |
| e004250 | 19:58:03.338005 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /final/views/1/replication/master_link_status |
| e004251 | 19:58:03.338005 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /final/views/1/replication/master_sync_in_progress |
| e004252 | 19:58:03.338005 | node-2 | info_replication | replication.master_repl_offset | valid | `113590` | /final/views/1/replication/master_repl_offset |
| e004253 | 19:58:03.338005 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /final/views/1/replication/slave_repl_offset |
| e004254 | 19:58:03.338005 | node-2 | cluster_nodes | redis_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /final/views/1/redis_id |
| e004255 | 19:58:03.338005 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /final/views/1/coverage |
| e004256 | 19:58:03.338005 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /final/views/1/nodes/0/flags |
| e004257 | 19:58:03.338005 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /final/views/1/nodes/0/primary_id |
| e004258 | 19:58:03.338005 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/1/nodes/0/slots |
| e004259 | 19:58:03.338005 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/0/link_state |
| e004260 | 19:58:03.338005 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/1/nodes/1/flags |
| e004261 | 19:58:03.338005 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /final/views/1/nodes/1/primary_id |
| e004262 | 19:58:03.338005 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /final/views/1/nodes/1/slots |
| e004263 | 19:58:03.338005 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/1/link_state |
| e004264 | 19:58:03.338005 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/1/nodes/2/flags |
| e004265 | 19:58:03.338005 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /final/views/1/nodes/2/primary_id |
| e004266 | 19:58:03.338005 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.slots | missing_from_schema | `` | /final/views/1/nodes/2/slots |
| e004267 | 19:58:03.338005 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/2/link_state |
| e004268 | 19:58:03.338005 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/1/nodes/3/flags |
| e004269 | 19:58:03.338005 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /final/views/1/nodes/3/primary_id |
| e004270 | 19:58:03.338005 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /final/views/1/nodes/3/slots |
| e004271 | 19:58:03.338005 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/3/link_state |
| e004272 | 19:58:03.338005 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /final/views/1/nodes/4/flags |
| e004273 | 19:58:03.338005 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /final/views/1/nodes/4/primary_id |
| e004274 | 19:58:03.338005 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/1/nodes/4/slots |
| e004275 | 19:58:03.338005 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/4/link_state |
| e004276 | 19:58:03.338005 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /final/views/1/nodes/5/flags |
| e004277 | 19:58:03.338005 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /final/views/1/nodes/5/primary_id |
| e004278 | 19:58:03.338005 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/1/nodes/5/slots |
| e004279 | 19:58:03.338005 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/5/link_state |
| e004280 | 19:58:03.338005 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/1/persistence/aof_enabled |
| e004281 | 19:58:03.338005 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /final/views/1/persistence/loading |
| e004282 | 19:58:03.338005 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/1/persistence/aof_last_write_status |
| e004283 | 19:58:03.338005 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/1/persistence/aof_rewrite_in_progress |
| e004284 | 19:58:03.338005 | node-2 | admin | process_live | missing_from_schema | `` | /final/views/1/process_live |
| e004285 | 19:58:03.338015 | node-3 | node_client | reachable | valid | `true` | /samples/16/snapshot/views/2/reachable |
| e004286 | 19:58:03.338015 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/16/snapshot/views/2/replication/role |
| e004287 | 19:58:03.338015 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/16/snapshot/views/2/replication/connected_slaves |
| e004288 | 19:58:03.338015 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/16/snapshot/views/2/replication/master_link_status |
| e004289 | 19:58:03.338015 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/16/snapshot/views/2/replication/master_sync_in_progress |
| e004290 | 19:58:03.338015 | node-3 | info_replication | replication.master_repl_offset | valid | `113553` | /samples/16/snapshot/views/2/replication/master_repl_offset |
| e004291 | 19:58:03.338015 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/16/snapshot/views/2/replication/slave_repl_offset |
| e004292 | 19:58:03.338015 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/16/snapshot/views/2/redis_id |
| e004293 | 19:58:03.338015 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/16/snapshot/views/2/coverage |
| e004294 | 19:58:03.338015 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/2/nodes/0/flags |
| e004295 | 19:58:03.338015 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/16/snapshot/views/2/nodes/0/primary_id |
| e004296 | 19:58:03.338015 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/2/nodes/0/slots |
| e004297 | 19:58:03.338015 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/2/nodes/0/link_state |
| e004298 | 19:58:03.338015 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/16/snapshot/views/2/nodes/1/flags |
| e004299 | 19:58:03.338015 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/2/nodes/1/primary_id |
| e004300 | 19:58:03.338015 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/16/snapshot/views/2/nodes/1/slots |
| e004301 | 19:58:03.338015 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/2/nodes/1/link_state |
| e004302 | 19:58:03.338015 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/2/nodes/2/flags |
| e004303 | 19:58:03.338015 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/16/snapshot/views/2/nodes/2/primary_id |
| e004304 | 19:58:03.338015 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/2/nodes/2/slots |
| e004305 | 19:58:03.338015 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/2/nodes/2/link_state |
| e004306 | 19:58:03.338015 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/2/nodes/3/flags |
| e004307 | 19:58:03.338015 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/2/nodes/3/primary_id |
| e004308 | 19:58:03.338015 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/16/snapshot/views/2/nodes/3/slots |
| e004309 | 19:58:03.338015 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/2/nodes/3/link_state |
| e004310 | 19:58:03.338015 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/2/nodes/4/flags |
| e004311 | 19:58:03.338015 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/2/nodes/4/primary_id |
| e004312 | 19:58:03.338015 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/16/snapshot/views/2/nodes/4/slots |
| e004313 | 19:58:03.338015 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/2/nodes/4/link_state |
| e004314 | 19:58:03.338015 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/2/nodes/5/flags |
| e004315 | 19:58:03.338015 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/16/snapshot/views/2/nodes/5/primary_id |
| e004316 | 19:58:03.338015 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/2/nodes/5/slots |
| e004317 | 19:58:03.338015 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/2/nodes/5/link_state |
| e004318 | 19:58:03.338015 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/16/snapshot/views/2/persistence/aof_enabled |
| e004319 | 19:58:03.338015 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/16/snapshot/views/2/persistence/loading |
| e004320 | 19:58:03.338015 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/16/snapshot/views/2/persistence/aof_last_write_status |
| e004321 | 19:58:03.338015 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/16/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e004322 | 19:58:03.338015 | node-3 | admin | process_live | missing_from_schema | `` | /samples/16/snapshot/views/2/process_live |
| e004323 | 19:58:03.338015 | node-3 | node_client | reachable | valid | `true` | /final/views/2/reachable |
| e004324 | 19:58:03.338015 | node-3 | info_replication | replication.role | valid | `"master"` | /final/views/2/replication/role |
| e004325 | 19:58:03.338015 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /final/views/2/replication/connected_slaves |
| e004326 | 19:58:03.338015 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /final/views/2/replication/master_link_status |
| e004327 | 19:58:03.338015 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /final/views/2/replication/master_sync_in_progress |
| e004328 | 19:58:03.338015 | node-3 | info_replication | replication.master_repl_offset | valid | `113553` | /final/views/2/replication/master_repl_offset |
| e004329 | 19:58:03.338015 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /final/views/2/replication/slave_repl_offset |
| e004330 | 19:58:03.338015 | node-3 | cluster_nodes | redis_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /final/views/2/redis_id |
| e004331 | 19:58:03.338015 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /final/views/2/coverage |
| e004332 | 19:58:03.338015 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/2/nodes/0/flags |
| e004333 | 19:58:03.338015 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /final/views/2/nodes/0/primary_id |
| e004334 | 19:58:03.338015 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /final/views/2/nodes/0/slots |
| e004335 | 19:58:03.338015 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/0/link_state |
| e004336 | 19:58:03.338015 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /final/views/2/nodes/1/flags |
| e004337 | 19:58:03.338015 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /final/views/2/nodes/1/primary_id |
| e004338 | 19:58:03.338015 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/2/nodes/1/slots |
| e004339 | 19:58:03.338015 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/1/link_state |
| e004340 | 19:58:03.338015 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/2/nodes/2/flags |
| e004341 | 19:58:03.338015 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /final/views/2/nodes/2/primary_id |
| e004342 | 19:58:03.338015 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.slots | missing_from_schema | `` | /final/views/2/nodes/2/slots |
| e004343 | 19:58:03.338015 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/2/link_state |
| e004344 | 19:58:03.338015 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /final/views/2/nodes/3/flags |
| e004345 | 19:58:03.338015 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /final/views/2/nodes/3/primary_id |
| e004346 | 19:58:03.338015 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/2/nodes/3/slots |
| e004347 | 19:58:03.338015 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/3/link_state |
| e004348 | 19:58:03.338015 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /final/views/2/nodes/4/flags |
| e004349 | 19:58:03.338015 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /final/views/2/nodes/4/primary_id |
| e004350 | 19:58:03.338015 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/2/nodes/4/slots |
| e004351 | 19:58:03.338015 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/4/link_state |
| e004352 | 19:58:03.338015 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/2/nodes/5/flags |
| e004353 | 19:58:03.338015 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /final/views/2/nodes/5/primary_id |
| e004354 | 19:58:03.338015 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /final/views/2/nodes/5/slots |
| e004355 | 19:58:03.338015 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/5/link_state |
| e004356 | 19:58:03.338015 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/2/persistence/aof_enabled |
| e004357 | 19:58:03.338015 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /final/views/2/persistence/loading |
| e004358 | 19:58:03.338015 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/2/persistence/aof_last_write_status |
| e004359 | 19:58:03.338015 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/2/persistence/aof_rewrite_in_progress |
| e004360 | 19:58:03.338015 | node-3 | admin | process_live | missing_from_schema | `` | /final/views/2/process_live |
| e004361 | 19:58:03.338024 | node-5 | node_client | reachable | valid | `true` | /samples/16/snapshot/views/4/reachable |
| e004362 | 19:58:03.338024 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/16/snapshot/views/4/replication/role |
| e004363 | 19:58:03.338024 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/16/snapshot/views/4/replication/connected_slaves |
| e004364 | 19:58:03.338024 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/16/snapshot/views/4/replication/master_link_status |
| e004365 | 19:58:03.338024 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/16/snapshot/views/4/replication/master_sync_in_progress |
| e004366 | 19:58:03.338024 | node-5 | info_replication | replication.master_repl_offset | valid | `113833` | /samples/16/snapshot/views/4/replication/master_repl_offset |
| e004367 | 19:58:03.338024 | node-5 | info_replication | replication.slave_repl_offset | valid | `113833` | /samples/16/snapshot/views/4/replication/slave_repl_offset |
| e004368 | 19:58:03.338024 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /samples/16/snapshot/views/4/redis_id |
| e004369 | 19:58:03.338024 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/16/snapshot/views/4/coverage |
| e004370 | 19:58:03.338024 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/4/nodes/0/flags |
| e004371 | 19:58:03.338024 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/4/nodes/0/primary_id |
| e004372 | 19:58:03.338024 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/16/snapshot/views/4/nodes/0/slots |
| e004373 | 19:58:03.338024 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/4/nodes/0/link_state |
| e004374 | 19:58:03.338024 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/4/nodes/1/flags |
| e004375 | 19:58:03.338024 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /samples/16/snapshot/views/4/nodes/1/primary_id |
| e004376 | 19:58:03.338024 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/4/nodes/1/slots |
| e004377 | 19:58:03.338024 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/4/nodes/1/link_state |
| e004378 | 19:58:03.338024 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/16/snapshot/views/4/nodes/2/flags |
| e004379 | 19:58:03.338024 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /samples/16/snapshot/views/4/nodes/2/primary_id |
| e004380 | 19:58:03.338024 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/4/nodes/2/slots |
| e004381 | 19:58:03.338024 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/4/nodes/2/link_state |
| e004382 | 19:58:03.338024 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/4/nodes/3/flags |
| e004383 | 19:58:03.338024 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/4/nodes/3/primary_id |
| e004384 | 19:58:03.338024 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/16/snapshot/views/4/nodes/3/slots |
| e004385 | 19:58:03.338024 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/4/nodes/3/link_state |
| e004386 | 19:58:03.338024 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/16/snapshot/views/4/nodes/4/flags |
| e004387 | 19:58:03.338024 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /samples/16/snapshot/views/4/nodes/4/primary_id |
| e004388 | 19:58:03.338024 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/16/snapshot/views/4/nodes/4/slots |
| e004389 | 19:58:03.338024 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/4/nodes/4/link_state |
| e004390 | 19:58:03.338024 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/16/snapshot/views/4/nodes/5/flags |
| e004391 | 19:58:03.338024 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/16/snapshot/views/4/nodes/5/primary_id |
| e004392 | 19:58:03.338024 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/16/snapshot/views/4/nodes/5/slots |
| e004393 | 19:58:03.338024 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/16/snapshot/views/4/nodes/5/link_state |
| e004394 | 19:58:03.338024 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/16/snapshot/views/4/persistence/aof_enabled |
| e004395 | 19:58:03.338024 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/16/snapshot/views/4/persistence/loading |
| e004396 | 19:58:03.338024 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/16/snapshot/views/4/persistence/aof_last_write_status |
| e004397 | 19:58:03.338024 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/16/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e004398 | 19:58:03.338024 | node-5 | admin | process_live | missing_from_schema | `` | /samples/16/snapshot/views/4/process_live |
| e004399 | 19:58:03.338024 | node-5 | node_client | reachable | valid | `true` | /final/views/4/reachable |
| e004400 | 19:58:03.338024 | node-5 | info_replication | replication.role | valid | `"slave"` | /final/views/4/replication/role |
| e004401 | 19:58:03.338024 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /final/views/4/replication/connected_slaves |
| e004402 | 19:58:03.338024 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /final/views/4/replication/master_link_status |
| e004403 | 19:58:03.338024 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /final/views/4/replication/master_sync_in_progress |
| e004404 | 19:58:03.338024 | node-5 | info_replication | replication.master_repl_offset | valid | `113833` | /final/views/4/replication/master_repl_offset |
| e004405 | 19:58:03.338024 | node-5 | info_replication | replication.slave_repl_offset | valid | `113833` | /final/views/4/replication/slave_repl_offset |
| e004406 | 19:58:03.338024 | node-5 | cluster_nodes | redis_id | valid | `"35a19f4babbca2395b87729c60bd927ff4a61972"` | /final/views/4/redis_id |
| e004407 | 19:58:03.338024 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /final/views/4/coverage |
| e004408 | 19:58:03.338024 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /final/views/4/nodes/0/flags |
| e004409 | 19:58:03.338024 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /final/views/4/nodes/0/primary_id |
| e004410 | 19:58:03.338024 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/4/nodes/0/slots |
| e004411 | 19:58:03.338024 | 32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/0/link_state |
| e004412 | 19:58:03.338024 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/4/nodes/1/flags |
| e004413 | 19:58:03.338024 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.primary_id | valid | `"32613a70bb1af812d9bd5fd8f48d0b6ad1c5184b"` | /final/views/4/nodes/1/primary_id |
| e004414 | 19:58:03.338024 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.slots | missing_from_schema | `` | /final/views/4/nodes/1/slots |
| e004415 | 19:58:03.338024 | 5ea13b969c72bacfbd0bf0baa0a6bf3016711fbf | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/1/link_state |
| e004416 | 19:58:03.338024 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /final/views/4/nodes/2/flags |
| e004417 | 19:58:03.338024 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.primary_id | valid | `"51d3089b3fef8198d88399f390d14bfa31ba0b26"` | /final/views/4/nodes/2/primary_id |
| e004418 | 19:58:03.338024 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /final/views/4/nodes/2/slots |
| e004419 | 19:58:03.338024 | 35a19f4babbca2395b87729c60bd927ff4a61972 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/2/link_state |
| e004420 | 19:58:03.338024 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /final/views/4/nodes/3/flags |
| e004421 | 19:58:03.338024 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /final/views/4/nodes/3/primary_id |
| e004422 | 19:58:03.338024 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/4/nodes/3/slots |
| e004423 | 19:58:03.338024 | 51d3089b3fef8198d88399f390d14bfa31ba0b26 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/3/link_state |
| e004424 | 19:58:03.338024 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/4/nodes/4/flags |
| e004425 | 19:58:03.338024 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.primary_id | valid | `"3dba8e31798f51b7da23cdf632bf27f37d96fa1c"` | /final/views/4/nodes/4/primary_id |
| e004426 | 19:58:03.338024 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /final/views/4/nodes/4/slots |
| e004427 | 19:58:03.338024 | bc00717ad0825a914118d066bbb0e87129689ab3 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/4/link_state |
| e004428 | 19:58:03.338024 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /final/views/4/nodes/5/flags |
| e004429 | 19:58:03.338024 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /final/views/4/nodes/5/primary_id |
| e004430 | 19:58:03.338024 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/4/nodes/5/slots |
| e004431 | 19:58:03.338024 | 3dba8e31798f51b7da23cdf632bf27f37d96fa1c | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/5/link_state |
| e004432 | 19:58:03.338024 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/4/persistence/aof_enabled |
| e004433 | 19:58:03.338024 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /final/views/4/persistence/loading |
| e004434 | 19:58:03.338024 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/4/persistence/aof_last_write_status |
| e004435 | 19:58:03.338024 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/4/persistence/aof_rewrite_in_progress |
| e004436 | 19:58:03.338024 | node-5 | admin | process_live | missing_from_schema | `` | /final/views/4/process_live |
| e004437 | 19:58:03.381014 |  | client_window | client.attempts | valid | `170` | /metrics/windows/14/attempts |
| e004438 | 19:58:03.381014 |  | client_window | client.errors | valid | `0` | /metrics/windows/14/errors |
| e004439 | 19:58:03.381014 |  | client_window | client.successes | valid | `170` | /metrics/windows/14/successes |
| e004440 | 19:58:03.381014 |  | client_window | client.p95_ms | valid | `1.19212` | /metrics/windows/14/p95_ms |
| e004441 | 19:58:03.485649 |  | client_probe | client_success | valid | `true` | /samples/16/probe/successful |
| e004442 | 19:58:03.600583 |  | ack_probe | ack.policy | valid | `"wait"` | /samples/16/write_policy_probe/policy |
| e004443 | 19:58:03.600583 |  | ack_probe | ack.successful | valid | `true` | /samples/16/write_policy_probe/successful |
| e004444 | 19:58:03.600583 |  | ack_probe | ack.requested_replicas | valid | `1` | /samples/16/write_policy_probe/requested_replicas |
| e004445 | 19:58:03.600583 |  | ack_probe | ack.set_acknowledged | valid | `[
          true,
          true,
          true
        ]` | /samples/16/write_policy_probe/set_acknowledged |
| e004446 | 19:58:03.600583 |  | ack_probe | ack.acknowledgements | valid | `[
          1,
          1,
          1
        ]` | /samples/16/write_policy_probe/acknowledgements |
| e004447 | 19:58:03.600583 |  | ack_probe | ack.acknowledgement_observed | valid | `[
          true,
          true,
          true
        ]` | /samples/16/write_policy_probe/acknowledgement_observed |
| e004448 | 19:58:03.600583 |  | ack_probe | ack.waitaof_requested_local | missing_from_schema | `` | /samples/16/write_policy_probe/waitaof_requested_local |
| e004449 | 19:58:03.600583 |  | ack_probe | ack.waitaof_requested_replicas | missing_from_schema | `` | /samples/16/write_policy_probe/waitaof_requested_replicas |
| e004450 | 19:58:03.600583 |  | ack_probe | ack.waitaof_local_acknowledgements | missing_from_schema | `` | /samples/16/write_policy_probe/waitaof_local_acknowledgements |
| e004451 | 19:58:03.600583 |  | ack_probe | ack.waitaof_replica_acknowledgements | missing_from_schema | `` | /samples/16/write_policy_probe/waitaof_replica_acknowledgements |
| e004452 | 19:58:03.600583 |  | ack_probe | ack.waitaof_acknowledgement_observed | missing_from_schema | `` | /samples/16/write_policy_probe/waitaof_acknowledgement_observed |
| e004453 | 19:58:03.686788 |  | client_window | client.attempts | valid | `52` | /metrics/windows/15/attempts |
| e004454 | 19:58:03.686788 |  | client_window | client.errors | valid | `0` | /metrics/windows/15/errors |
| e004455 | 19:58:03.686788 |  | client_window | client.successes | valid | `52` | /metrics/windows/15/successes |
| e004456 | 19:58:03.686788 |  | client_window | client.p95_ms | valid | `1.19212` | /metrics/windows/15/p95_ms |
| e004457 | 19:58:04.108348 |  | value_validation | validation.expected | valid | `1303` | /validation/expected |
| e004458 | 19:58:04.108348 |  | value_validation | validation.checked | valid | `1303` | /validation/checked |
| e004459 | 19:58:04.108348 |  | value_validation | validation.missing | valid | `0` | /validation/missing |
| e004460 | 19:58:04.108348 |  | value_validation | validation.corrupt | valid | `0` | /validation/corrupt |
| e004461 | 19:58:04.108348 |  | value_validation | validation.unreadable | valid | `0` | /validation/unreadable |
| e004462 | 19:58:04.108348 |  | value_validation | validation.unverified | valid | `0` | /validation/unverified |
| e004463 | 19:58:04.108348 |  | value_validation | validation.complete | valid | `true` | /validation/complete |

## Analysis limitations

- Offline laboratory observation analysis; no production integration.
- Node views are non-atomic and retained independently; sample timestamps are not exact server event times.
- No fault labels, injected action results, run identifiers, configuration hashes or phase names are inference inputs.
- Client reachability does not establish process death; no administrative process telemetry is accepted from fault-specific fields.
- Validation proves only the recorded checked population, never general durability or power-loss safety.
