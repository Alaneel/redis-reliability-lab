# Incident diagnosis

Incident: `f22e11d33b9fb3a555bad521`  
Detection: **incident**  
Unique root cause: **abstained**  
Engine: `rules-2`

## Observed symptoms and diagnostic hypotheses

### data_validation_failure

Final acknowledged-value validation observed missing or corrupt values. Serving requests does not establish data preservation.

Assessment: strongly_supported. Rule: `R10-positive-validation-loss-counter`. Nodes: .

Supporting evidence: e002075.

Contrary or recovery observations: .

Missing evidence: Independent process liveness and network-path observations; Per-request routing/cache/redirect traces.

Alternatives: Process or endpoint failure; Network or observer connectivity problem; Planned administrative change.

### replica_unavailability

Replica INFO reports its primary link down; process and network causes remain alternatives.

Assessment: strongly_supported. Rule: `R4-replica-link-down`. Nodes: node-4, node-5, node-6.

Supporting evidence: e000366, e000404, e000442, e000518, e000632, e000784, e000860, e000936, e001088, e001354.

Contrary or recovery observations: e000017, e000131, e000207, e000363, e000401, e000439, e000515, e000629, e000667, e000781, e000819, e000857, e000933, e000971, e001085, e001161.

Missing evidence: Independent process liveness and network-path observations; Per-request routing/cache/redirect traces.

Alternatives: Process or endpoint failure; Network or observer connectivity problem; Planned administrative change.

## Timeline

| First observed time | Event | Evidence | Timing |
|---|---|---|---|
| 2026-10-11T19:58:27.576453+08:00 | baseline_full_slot_coverage_observed | e000025, e000063, e000101, e000139, e000177, e000215 | all recorded views in this non-atomic sample report full healthy coverage |
| 2026-10-11T19:58:44.333718+08:00 | first_observed_replica_unavailability | e000366 | first_sampled_observation; exact event time unknown |
| 2026-10-11T19:58:55.983629+08:00 | first_observed_data_validation_failure | e002075 | first_sampled_observation; exact event time unknown |

## Recovery and uncertainty

- Last client probe successful=true (e003452)
- Insufficient evidence to identify a unique root cause. Process failure, network disruption and observation-channel failure can be indistinguishable.

## Read-only remediation plans

- **data_validation_failure**: Preserve evidence and investigate acknowledged-value validation failures. Evidence: e002075. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability; Rejoining or overwriting a node may destroy recoverable values. Expected effect: Resolve alternative explanations without executing remediation. Verify: Inspect independently collected process and network status; Compare all node views and client outcomes at aligned timestamps; Read the frozen acknowledged ledger and compare expected values; retain unreadable and unchecked counts. Human approval required; no execution.
- **replica_unavailability**: Assess restoration of a synchronized replica. Evidence: e000366, e000404, e000442, e000518, e000632, e000784, e000860, e000936, e001088, e001354. Preconditions: Confirm affected node identity and current topology; Obtain operator approval before any state-changing action; Preserve logs and verify available backups; Confirm an authoritative healthy primary and compatible former-node identity and data. Risks: Stale observations may identify the wrong current owner; Changing topology can lose data or interrupt availability. Expected effect: Resolve alternative explanations without executing remediation. Verify: Inspect independently collected process and network status; Compare all node views and client outcomes at aligned timestamps; Observe replica assignment, link up, synchronization complete and repeated healthy service probes. Human approval required; no execution.

## Evidence provenance

Original bytes SHA-256: `2cc8bef726446e851d5f1109e0fad0d2057e4c8764c7f623c307c388d6be827b`. Input report schema: 3. References are JSON pointers in the original report.

| ID | Time | Node | Source | Signal | Status | Value | Reference |
|---|---|---|---|---|---|---|---|
| e000001 | 00:00:00 |  | client_probe | client_success | invalid | `false` | /samples/0/probe/successful |
| e000002 | 00:00:00 |  | client_probe | client_success | invalid | `false` | /samples/1/probe/successful |
| e000003 | 00:00:00 |  | client_probe | client_success | invalid | `false` | /samples/2/probe/successful |
| e000004 | 00:00:00 |  | client_probe | client_success | invalid | `false` | /samples/3/probe/successful |
| e000005 | 00:00:00 |  | client_probe | client_success | invalid | `false` | /samples/4/probe/successful |
| e000006 | 00:00:00 |  | client_probe | client_success | invalid | `false` | /samples/5/probe/successful |
| e000007 | 00:00:00 |  | client_probe | client_success | invalid | `false` | /samples/6/probe/successful |
| e000008 | 00:00:00 |  | client_probe | client_success | invalid | `false` | /samples/7/probe/successful |
| e000009 | 19:58:26.575746 |  | client_window | client.attempts | valid | `169` | /metrics/windows/0/attempts |
| e000010 | 19:58:26.575746 |  | client_window | client.errors | valid | `0` | /metrics/windows/0/errors |
| e000011 | 19:58:26.575746 |  | client_window | client.successes | valid | `169` | /metrics/windows/0/successes |
| e000012 | 19:58:26.575746 |  | client_window | client.p95_ms | valid | `0.632203` | /metrics/windows/0/p95_ms |
| e000013 | 19:58:27.575746 |  | client_window | client.attempts | valid | `170` | /metrics/windows/1/attempts |
| e000014 | 19:58:27.575746 |  | client_window | client.errors | valid | `0` | /metrics/windows/1/errors |
| e000015 | 19:58:27.575746 |  | client_window | client.successes | valid | `170` | /metrics/windows/1/successes |
| e000016 | 19:58:27.575746 |  | client_window | client.p95_ms | valid | `0.632203` | /metrics/windows/1/p95_ms |
| e000017 | 19:58:27.576404 | node-4 | node_client | reachable | valid | `true` | /initial/views/3/reachable |
| e000018 | 19:58:27.576404 | node-4 | info_replication | replication.role | valid | `"slave"` | /initial/views/3/replication/role |
| e000019 | 19:58:27.576404 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /initial/views/3/replication/connected_slaves |
| e000020 | 19:58:27.576404 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /initial/views/3/replication/master_link_status |
| e000021 | 19:58:27.576404 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /initial/views/3/replication/master_sync_in_progress |
| e000022 | 19:58:27.576404 | node-4 | info_replication | replication.master_repl_offset | valid | `13321` | /initial/views/3/replication/master_repl_offset |
| e000023 | 19:58:27.576404 | node-4 | info_replication | replication.slave_repl_offset | valid | `13321` | /initial/views/3/replication/slave_repl_offset |
| e000024 | 19:58:27.576404 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /initial/views/3/redis_id |
| e000025 | 19:58:27.576404 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/3/coverage |
| e000026 | 19:58:27.576404 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/3/nodes/0/flags |
| e000027 | 19:58:27.576404 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /initial/views/3/nodes/0/primary_id |
| e000028 | 19:58:27.576404 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/3/nodes/0/slots |
| e000029 | 19:58:27.576404 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/0/link_state |
| e000030 | 19:58:27.576404 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /initial/views/3/nodes/1/flags |
| e000031 | 19:58:27.576404 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /initial/views/3/nodes/1/primary_id |
| e000032 | 19:58:27.576404 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /initial/views/3/nodes/1/slots |
| e000033 | 19:58:27.576404 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/1/link_state |
| e000034 | 19:58:27.576404 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/3/nodes/2/flags |
| e000035 | 19:58:27.576404 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /initial/views/3/nodes/2/primary_id |
| e000036 | 19:58:27.576404 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /initial/views/3/nodes/2/slots |
| e000037 | 19:58:27.576404 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/2/link_state |
| e000038 | 19:58:27.576404 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/3/nodes/3/flags |
| e000039 | 19:58:27.576404 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /initial/views/3/nodes/3/primary_id |
| e000040 | 19:58:27.576404 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/3/nodes/3/slots |
| e000041 | 19:58:27.576404 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/3/link_state |
| e000042 | 19:58:27.576404 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/3/nodes/4/flags |
| e000043 | 19:58:27.576404 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /initial/views/3/nodes/4/primary_id |
| e000044 | 19:58:27.576404 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /initial/views/3/nodes/4/slots |
| e000045 | 19:58:27.576404 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/4/link_state |
| e000046 | 19:58:27.576404 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/3/nodes/5/flags |
| e000047 | 19:58:27.576404 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /initial/views/3/nodes/5/primary_id |
| e000048 | 19:58:27.576404 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/3/nodes/5/slots |
| e000049 | 19:58:27.576404 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /initial/views/3/nodes/5/link_state |
| e000050 | 19:58:27.576404 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/3/persistence/aof_enabled |
| e000051 | 19:58:27.576404 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/3/persistence/loading |
| e000052 | 19:58:27.576404 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/3/persistence/aof_last_write_status |
| e000053 | 19:58:27.576404 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/3/persistence/aof_rewrite_in_progress |
| e000054 | 19:58:27.576404 | node-4 | admin | process_live | missing_from_schema | `` | /initial/views/3/process_live |
| e000055 | 19:58:27.576414 | node-1 | node_client | reachable | valid | `true` | /initial/views/0/reachable |
| e000056 | 19:58:27.576414 | node-1 | info_replication | replication.role | valid | `"master"` | /initial/views/0/replication/role |
| e000057 | 19:58:27.576414 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /initial/views/0/replication/connected_slaves |
| e000058 | 19:58:27.576414 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /initial/views/0/replication/master_link_status |
| e000059 | 19:58:27.576414 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /initial/views/0/replication/master_sync_in_progress |
| e000060 | 19:58:27.576414 | node-1 | info_replication | replication.master_repl_offset | valid | `13541` | /initial/views/0/replication/master_repl_offset |
| e000061 | 19:58:27.576414 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /initial/views/0/replication/slave_repl_offset |
| e000062 | 19:58:27.576414 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /initial/views/0/redis_id |
| e000063 | 19:58:27.576414 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/0/coverage |
| e000064 | 19:58:27.576414 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /initial/views/0/nodes/0/flags |
| e000065 | 19:58:27.576414 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /initial/views/0/nodes/0/primary_id |
| e000066 | 19:58:27.576414 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/0/nodes/0/slots |
| e000067 | 19:58:27.576414 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/0/link_state |
| e000068 | 19:58:27.576414 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/0/nodes/1/flags |
| e000069 | 19:58:27.576414 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /initial/views/0/nodes/1/primary_id |
| e000070 | 19:58:27.576414 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /initial/views/0/nodes/1/slots |
| e000071 | 19:58:27.576414 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/1/link_state |
| e000072 | 19:58:27.576414 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/0/nodes/2/flags |
| e000073 | 19:58:27.576414 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /initial/views/0/nodes/2/primary_id |
| e000074 | 19:58:27.576414 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /initial/views/0/nodes/2/slots |
| e000075 | 19:58:27.576414 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/2/link_state |
| e000076 | 19:58:27.576414 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/0/nodes/3/flags |
| e000077 | 19:58:27.576414 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /initial/views/0/nodes/3/primary_id |
| e000078 | 19:58:27.576414 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/0/nodes/3/slots |
| e000079 | 19:58:27.576414 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/3/link_state |
| e000080 | 19:58:27.576414 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/0/nodes/4/flags |
| e000081 | 19:58:27.576414 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /initial/views/0/nodes/4/primary_id |
| e000082 | 19:58:27.576414 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /initial/views/0/nodes/4/slots |
| e000083 | 19:58:27.576414 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/4/link_state |
| e000084 | 19:58:27.576414 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/0/nodes/5/flags |
| e000085 | 19:58:27.576414 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /initial/views/0/nodes/5/primary_id |
| e000086 | 19:58:27.576414 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/0/nodes/5/slots |
| e000087 | 19:58:27.576414 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /initial/views/0/nodes/5/link_state |
| e000088 | 19:58:27.576414 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/0/persistence/aof_enabled |
| e000089 | 19:58:27.576414 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/0/persistence/loading |
| e000090 | 19:58:27.576414 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/0/persistence/aof_last_write_status |
| e000091 | 19:58:27.576414 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/0/persistence/aof_rewrite_in_progress |
| e000092 | 19:58:27.576414 | node-1 | admin | process_live | missing_from_schema | `` | /initial/views/0/process_live |
| e000093 | 19:58:27.576435 | node-2 | node_client | reachable | valid | `true` | /initial/views/1/reachable |
| e000094 | 19:58:27.576435 | node-2 | info_replication | replication.role | valid | `"master"` | /initial/views/1/replication/role |
| e000095 | 19:58:27.576435 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /initial/views/1/replication/connected_slaves |
| e000096 | 19:58:27.576435 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /initial/views/1/replication/master_link_status |
| e000097 | 19:58:27.576435 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /initial/views/1/replication/master_sync_in_progress |
| e000098 | 19:58:27.576435 | node-2 | info_replication | replication.master_repl_offset | valid | `13321` | /initial/views/1/replication/master_repl_offset |
| e000099 | 19:58:27.576435 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /initial/views/1/replication/slave_repl_offset |
| e000100 | 19:58:27.576435 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /initial/views/1/redis_id |
| e000101 | 19:58:27.576435 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/1/coverage |
| e000102 | 19:58:27.576435 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/1/nodes/0/flags |
| e000103 | 19:58:27.576435 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /initial/views/1/nodes/0/primary_id |
| e000104 | 19:58:27.576435 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/1/nodes/0/slots |
| e000105 | 19:58:27.576435 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/0/link_state |
| e000106 | 19:58:27.576435 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /initial/views/1/nodes/1/flags |
| e000107 | 19:58:27.576435 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /initial/views/1/nodes/1/primary_id |
| e000108 | 19:58:27.576435 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/1/nodes/1/slots |
| e000109 | 19:58:27.576435 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/1/link_state |
| e000110 | 19:58:27.576435 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/1/nodes/2/flags |
| e000111 | 19:58:27.576435 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /initial/views/1/nodes/2/primary_id |
| e000112 | 19:58:27.576435 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /initial/views/1/nodes/2/slots |
| e000113 | 19:58:27.576435 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/2/link_state |
| e000114 | 19:58:27.576435 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/1/nodes/3/flags |
| e000115 | 19:58:27.576435 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /initial/views/1/nodes/3/primary_id |
| e000116 | 19:58:27.576435 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/1/nodes/3/slots |
| e000117 | 19:58:27.576435 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/3/link_state |
| e000118 | 19:58:27.576435 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/1/nodes/4/flags |
| e000119 | 19:58:27.576435 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /initial/views/1/nodes/4/primary_id |
| e000120 | 19:58:27.576435 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /initial/views/1/nodes/4/slots |
| e000121 | 19:58:27.576435 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/4/link_state |
| e000122 | 19:58:27.576435 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/1/nodes/5/flags |
| e000123 | 19:58:27.576435 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /initial/views/1/nodes/5/primary_id |
| e000124 | 19:58:27.576435 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /initial/views/1/nodes/5/slots |
| e000125 | 19:58:27.576435 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /initial/views/1/nodes/5/link_state |
| e000126 | 19:58:27.576435 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/1/persistence/aof_enabled |
| e000127 | 19:58:27.576435 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/1/persistence/loading |
| e000128 | 19:58:27.576435 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/1/persistence/aof_last_write_status |
| e000129 | 19:58:27.576435 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/1/persistence/aof_rewrite_in_progress |
| e000130 | 19:58:27.576435 | node-2 | admin | process_live | missing_from_schema | `` | /initial/views/1/process_live |
| e000131 | 19:58:27.576438 | node-6 | node_client | reachable | valid | `true` | /initial/views/5/reachable |
| e000132 | 19:58:27.576438 | node-6 | info_replication | replication.role | valid | `"slave"` | /initial/views/5/replication/role |
| e000133 | 19:58:27.576438 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /initial/views/5/replication/connected_slaves |
| e000134 | 19:58:27.576438 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /initial/views/5/replication/master_link_status |
| e000135 | 19:58:27.576438 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /initial/views/5/replication/master_sync_in_progress |
| e000136 | 19:58:27.576438 | node-6 | info_replication | replication.master_repl_offset | valid | `13321` | /initial/views/5/replication/master_repl_offset |
| e000137 | 19:58:27.576438 | node-6 | info_replication | replication.slave_repl_offset | valid | `13321` | /initial/views/5/replication/slave_repl_offset |
| e000138 | 19:58:27.576438 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /initial/views/5/redis_id |
| e000139 | 19:58:27.576438 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/5/coverage |
| e000140 | 19:58:27.576438 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/5/nodes/0/flags |
| e000141 | 19:58:27.576438 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /initial/views/5/nodes/0/primary_id |
| e000142 | 19:58:27.576438 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/5/nodes/0/slots |
| e000143 | 19:58:27.576438 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/0/link_state |
| e000144 | 19:58:27.576438 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /initial/views/5/nodes/1/flags |
| e000145 | 19:58:27.576438 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /initial/views/5/nodes/1/primary_id |
| e000146 | 19:58:27.576438 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /initial/views/5/nodes/1/slots |
| e000147 | 19:58:27.576438 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/1/link_state |
| e000148 | 19:58:27.576438 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/5/nodes/2/flags |
| e000149 | 19:58:27.576438 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /initial/views/5/nodes/2/primary_id |
| e000150 | 19:58:27.576438 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/5/nodes/2/slots |
| e000151 | 19:58:27.576438 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/2/link_state |
| e000152 | 19:58:27.576438 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/5/nodes/3/flags |
| e000153 | 19:58:27.576438 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /initial/views/5/nodes/3/primary_id |
| e000154 | 19:58:27.576438 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/5/nodes/3/slots |
| e000155 | 19:58:27.576438 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/3/link_state |
| e000156 | 19:58:27.576438 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/5/nodes/4/flags |
| e000157 | 19:58:27.576438 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /initial/views/5/nodes/4/primary_id |
| e000158 | 19:58:27.576438 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /initial/views/5/nodes/4/slots |
| e000159 | 19:58:27.576438 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/4/link_state |
| e000160 | 19:58:27.576438 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/5/nodes/5/flags |
| e000161 | 19:58:27.576438 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /initial/views/5/nodes/5/primary_id |
| e000162 | 19:58:27.576438 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /initial/views/5/nodes/5/slots |
| e000163 | 19:58:27.576438 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /initial/views/5/nodes/5/link_state |
| e000164 | 19:58:27.576438 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/5/persistence/aof_enabled |
| e000165 | 19:58:27.576438 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/5/persistence/loading |
| e000166 | 19:58:27.576438 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/5/persistence/aof_last_write_status |
| e000167 | 19:58:27.576438 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/5/persistence/aof_rewrite_in_progress |
| e000168 | 19:58:27.576438 | node-6 | admin | process_live | missing_from_schema | `` | /initial/views/5/process_live |
| e000169 | 19:58:27.576452 | node-3 | node_client | reachable | valid | `true` | /initial/views/2/reachable |
| e000170 | 19:58:27.576452 | node-3 | info_replication | replication.role | valid | `"master"` | /initial/views/2/replication/role |
| e000171 | 19:58:27.576452 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /initial/views/2/replication/connected_slaves |
| e000172 | 19:58:27.576452 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /initial/views/2/replication/master_link_status |
| e000173 | 19:58:27.576452 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /initial/views/2/replication/master_sync_in_progress |
| e000174 | 19:58:27.576452 | node-3 | info_replication | replication.master_repl_offset | valid | `13321` | /initial/views/2/replication/master_repl_offset |
| e000175 | 19:58:27.576452 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /initial/views/2/replication/slave_repl_offset |
| e000176 | 19:58:27.576452 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /initial/views/2/redis_id |
| e000177 | 19:58:27.576452 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/2/coverage |
| e000178 | 19:58:27.576452 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/2/nodes/0/flags |
| e000179 | 19:58:27.576452 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /initial/views/2/nodes/0/primary_id |
| e000180 | 19:58:27.576452 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/2/nodes/0/slots |
| e000181 | 19:58:27.576452 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/0/link_state |
| e000182 | 19:58:27.576452 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /initial/views/2/nodes/1/flags |
| e000183 | 19:58:27.576452 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /initial/views/2/nodes/1/primary_id |
| e000184 | 19:58:27.576452 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/2/nodes/1/slots |
| e000185 | 19:58:27.576452 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/1/link_state |
| e000186 | 19:58:27.576452 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/2/nodes/2/flags |
| e000187 | 19:58:27.576452 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /initial/views/2/nodes/2/primary_id |
| e000188 | 19:58:27.576452 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/2/nodes/2/slots |
| e000189 | 19:58:27.576452 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/2/link_state |
| e000190 | 19:58:27.576452 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/2/nodes/3/flags |
| e000191 | 19:58:27.576452 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /initial/views/2/nodes/3/primary_id |
| e000192 | 19:58:27.576452 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /initial/views/2/nodes/3/slots |
| e000193 | 19:58:27.576452 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/3/link_state |
| e000194 | 19:58:27.576452 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/2/nodes/4/flags |
| e000195 | 19:58:27.576452 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /initial/views/2/nodes/4/primary_id |
| e000196 | 19:58:27.576452 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /initial/views/2/nodes/4/slots |
| e000197 | 19:58:27.576452 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/4/link_state |
| e000198 | 19:58:27.576452 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/2/nodes/5/flags |
| e000199 | 19:58:27.576452 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /initial/views/2/nodes/5/primary_id |
| e000200 | 19:58:27.576452 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /initial/views/2/nodes/5/slots |
| e000201 | 19:58:27.576452 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /initial/views/2/nodes/5/link_state |
| e000202 | 19:58:27.576452 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/2/persistence/aof_enabled |
| e000203 | 19:58:27.576452 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/2/persistence/loading |
| e000204 | 19:58:27.576452 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/2/persistence/aof_last_write_status |
| e000205 | 19:58:27.576452 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/2/persistence/aof_rewrite_in_progress |
| e000206 | 19:58:27.576452 | node-3 | admin | process_live | missing_from_schema | `` | /initial/views/2/process_live |
| e000207 | 19:58:27.576453 | node-5 | node_client | reachable | valid | `true` | /initial/views/4/reachable |
| e000208 | 19:58:27.576453 | node-5 | info_replication | replication.role | valid | `"slave"` | /initial/views/4/replication/role |
| e000209 | 19:58:27.576453 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /initial/views/4/replication/connected_slaves |
| e000210 | 19:58:27.576453 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /initial/views/4/replication/master_link_status |
| e000211 | 19:58:27.576453 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /initial/views/4/replication/master_sync_in_progress |
| e000212 | 19:58:27.576453 | node-5 | info_replication | replication.master_repl_offset | valid | `13541` | /initial/views/4/replication/master_repl_offset |
| e000213 | 19:58:27.576453 | node-5 | info_replication | replication.slave_repl_offset | valid | `13541` | /initial/views/4/replication/slave_repl_offset |
| e000214 | 19:58:27.576453 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /initial/views/4/redis_id |
| e000215 | 19:58:27.576453 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /initial/views/4/coverage |
| e000216 | 19:58:27.576453 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/4/nodes/0/flags |
| e000217 | 19:58:27.576453 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /initial/views/4/nodes/0/primary_id |
| e000218 | 19:58:27.576453 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /initial/views/4/nodes/0/slots |
| e000219 | 19:58:27.576453 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/0/link_state |
| e000220 | 19:58:27.576453 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/4/nodes/1/flags |
| e000221 | 19:58:27.576453 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /initial/views/4/nodes/1/primary_id |
| e000222 | 19:58:27.576453 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /initial/views/4/nodes/1/slots |
| e000223 | 19:58:27.576453 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/1/link_state |
| e000224 | 19:58:27.576453 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/4/nodes/2/flags |
| e000225 | 19:58:27.576453 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /initial/views/4/nodes/2/primary_id |
| e000226 | 19:58:27.576453 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /initial/views/4/nodes/2/slots |
| e000227 | 19:58:27.576453 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/2/link_state |
| e000228 | 19:58:27.576453 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /initial/views/4/nodes/3/flags |
| e000229 | 19:58:27.576453 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /initial/views/4/nodes/3/primary_id |
| e000230 | 19:58:27.576453 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /initial/views/4/nodes/3/slots |
| e000231 | 19:58:27.576453 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/3/link_state |
| e000232 | 19:58:27.576453 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /initial/views/4/nodes/4/flags |
| e000233 | 19:58:27.576453 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /initial/views/4/nodes/4/primary_id |
| e000234 | 19:58:27.576453 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /initial/views/4/nodes/4/slots |
| e000235 | 19:58:27.576453 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/4/link_state |
| e000236 | 19:58:27.576453 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /initial/views/4/nodes/5/flags |
| e000237 | 19:58:27.576453 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /initial/views/4/nodes/5/primary_id |
| e000238 | 19:58:27.576453 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /initial/views/4/nodes/5/slots |
| e000239 | 19:58:27.576453 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /initial/views/4/nodes/5/link_state |
| e000240 | 19:58:27.576453 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /initial/views/4/persistence/aof_enabled |
| e000241 | 19:58:27.576453 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /initial/views/4/persistence/loading |
| e000242 | 19:58:27.576453 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /initial/views/4/persistence/aof_last_write_status |
| e000243 | 19:58:27.576453 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /initial/views/4/persistence/aof_rewrite_in_progress |
| e000244 | 19:58:27.576453 | node-5 | admin | process_live | missing_from_schema | `` | /initial/views/4/process_live |
| e000245 | 19:58:27.722336 |  | client_window | client.attempts | valid | `25` | /metrics/windows/2/attempts |
| e000246 | 19:58:27.722336 |  | client_window | client.errors | valid | `0` | /metrics/windows/2/errors |
| e000247 | 19:58:27.722336 |  | client_window | client.successes | valid | `25` | /metrics/windows/2/successes |
| e000248 | 19:58:27.722336 |  | client_window | client.p95_ms | valid | `0.54612` | /metrics/windows/2/p95_ms |
| e000249 | 19:58:44.333694 | node-1 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/0/reachable |
| e000250 | 19:58:44.333694 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/0/snapshot/views/0/replication/role |
| e000251 | 19:58:44.333694 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/0/snapshot/views/0/replication/connected_slaves |
| e000252 | 19:58:44.333694 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/0/snapshot/views/0/replication/master_link_status |
| e000253 | 19:58:44.333694 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/0/snapshot/views/0/replication/master_sync_in_progress |
| e000254 | 19:58:44.333694 | node-1 | info_replication | replication.master_repl_offset | valid | `0` | /samples/0/snapshot/views/0/replication/master_repl_offset |
| e000255 | 19:58:44.333694 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/0/snapshot/views/0/replication/slave_repl_offset |
| e000256 | 19:58:44.333694 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/0/snapshot/views/0/redis_id |
| e000257 | 19:58:44.333694 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/0/coverage |
| e000258 | 19:58:44.333694 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/0/nodes/0/flags |
| e000259 | 19:58:44.333694 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/0/snapshot/views/0/nodes/0/primary_id |
| e000260 | 19:58:44.333694 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/0/slots |
| e000261 | 19:58:44.333694 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/0/nodes/0/link_state |
| e000262 | 19:58:44.333694 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/0/nodes/1/flags |
| e000263 | 19:58:44.333694 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/1/primary_id |
| e000264 | 19:58:44.333694 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/0/nodes/1/slots |
| e000265 | 19:58:44.333694 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/0/nodes/1/link_state |
| e000266 | 19:58:44.333694 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/0/nodes/2/flags |
| e000267 | 19:58:44.333694 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/2/primary_id |
| e000268 | 19:58:44.333694 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/0/nodes/2/slots |
| e000269 | 19:58:44.333694 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/0/nodes/2/link_state |
| e000270 | 19:58:44.333694 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/0/snapshot/views/0/nodes/3/flags |
| e000271 | 19:58:44.333694 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/3/primary_id |
| e000272 | 19:58:44.333694 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/0/nodes/3/slots |
| e000273 | 19:58:44.333694 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/0/nodes/3/link_state |
| e000274 | 19:58:44.333694 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/0/nodes/4/flags |
| e000275 | 19:58:44.333694 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/0/snapshot/views/0/nodes/4/primary_id |
| e000276 | 19:58:44.333694 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/4/slots |
| e000277 | 19:58:44.333694 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/0/nodes/4/link_state |
| e000278 | 19:58:44.333694 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/0/nodes/5/flags |
| e000279 | 19:58:44.333694 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/0/snapshot/views/0/nodes/5/primary_id |
| e000280 | 19:58:44.333694 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/0/nodes/5/slots |
| e000281 | 19:58:44.333694 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/0/nodes/5/link_state |
| e000282 | 19:58:44.333694 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/0/persistence/aof_enabled |
| e000283 | 19:58:44.333694 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/0/persistence/loading |
| e000284 | 19:58:44.333694 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/0/persistence/aof_last_write_status |
| e000285 | 19:58:44.333694 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e000286 | 19:58:44.333694 | node-1 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/0/process_live |
| e000287 | 19:58:44.333717 | node-3 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/2/reachable |
| e000288 | 19:58:44.333717 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/0/snapshot/views/2/replication/role |
| e000289 | 19:58:44.333717 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/0/snapshot/views/2/replication/connected_slaves |
| e000290 | 19:58:44.333717 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/0/snapshot/views/2/replication/master_link_status |
| e000291 | 19:58:44.333717 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/0/snapshot/views/2/replication/master_sync_in_progress |
| e000292 | 19:58:44.333717 | node-3 | info_replication | replication.master_repl_offset | valid | `0` | /samples/0/snapshot/views/2/replication/master_repl_offset |
| e000293 | 19:58:44.333717 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/0/snapshot/views/2/replication/slave_repl_offset |
| e000294 | 19:58:44.333717 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/0/snapshot/views/2/redis_id |
| e000295 | 19:58:44.333717 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/2/coverage |
| e000296 | 19:58:44.333717 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/2/nodes/0/flags |
| e000297 | 19:58:44.333717 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/0/primary_id |
| e000298 | 19:58:44.333717 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/2/nodes/0/slots |
| e000299 | 19:58:44.333717 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/0/link_state |
| e000300 | 19:58:44.333717 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/0/snapshot/views/2/nodes/1/flags |
| e000301 | 19:58:44.333717 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/1/primary_id |
| e000302 | 19:58:44.333717 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/2/nodes/1/slots |
| e000303 | 19:58:44.333717 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/1/link_state |
| e000304 | 19:58:44.333717 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/2/nodes/2/flags |
| e000305 | 19:58:44.333717 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/0/snapshot/views/2/nodes/2/primary_id |
| e000306 | 19:58:44.333717 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/2/slots |
| e000307 | 19:58:44.333717 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/2/link_state |
| e000308 | 19:58:44.333717 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/2/nodes/3/flags |
| e000309 | 19:58:44.333717 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/0/snapshot/views/2/nodes/3/primary_id |
| e000310 | 19:58:44.333717 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/3/slots |
| e000311 | 19:58:44.333717 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/3/link_state |
| e000312 | 19:58:44.333717 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/2/nodes/4/flags |
| e000313 | 19:58:44.333717 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/0/snapshot/views/2/nodes/4/primary_id |
| e000314 | 19:58:44.333717 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/4/slots |
| e000315 | 19:58:44.333717 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/4/link_state |
| e000316 | 19:58:44.333717 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/2/nodes/5/flags |
| e000317 | 19:58:44.333717 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/2/nodes/5/primary_id |
| e000318 | 19:58:44.333717 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/2/nodes/5/slots |
| e000319 | 19:58:44.333717 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/2/nodes/5/link_state |
| e000320 | 19:58:44.333717 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/2/persistence/aof_enabled |
| e000321 | 19:58:44.333717 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/2/persistence/loading |
| e000322 | 19:58:44.333717 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/2/persistence/aof_last_write_status |
| e000323 | 19:58:44.333717 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e000324 | 19:58:44.333717 | node-3 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/2/process_live |
| e000325 | 19:58:44.333718 | node-2 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/1/reachable |
| e000326 | 19:58:44.333718 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/0/snapshot/views/1/replication/role |
| e000327 | 19:58:44.333718 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/0/snapshot/views/1/replication/connected_slaves |
| e000328 | 19:58:44.333718 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/0/snapshot/views/1/replication/master_link_status |
| e000329 | 19:58:44.333718 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/0/snapshot/views/1/replication/master_sync_in_progress |
| e000330 | 19:58:44.333718 | node-2 | info_replication | replication.master_repl_offset | valid | `0` | /samples/0/snapshot/views/1/replication/master_repl_offset |
| e000331 | 19:58:44.333718 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/0/snapshot/views/1/replication/slave_repl_offset |
| e000332 | 19:58:44.333718 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/0/snapshot/views/1/redis_id |
| e000333 | 19:58:44.333718 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/1/coverage |
| e000334 | 19:58:44.333718 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/1/nodes/0/flags |
| e000335 | 19:58:44.333718 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/0/snapshot/views/1/nodes/0/primary_id |
| e000336 | 19:58:44.333718 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/0/slots |
| e000337 | 19:58:44.333718 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/0/link_state |
| e000338 | 19:58:44.333718 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/1/nodes/1/flags |
| e000339 | 19:58:44.333718 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/1/primary_id |
| e000340 | 19:58:44.333718 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/1/nodes/1/slots |
| e000341 | 19:58:44.333718 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/1/link_state |
| e000342 | 19:58:44.333718 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/1/nodes/2/flags |
| e000343 | 19:58:44.333718 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/0/snapshot/views/1/nodes/2/primary_id |
| e000344 | 19:58:44.333718 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/2/slots |
| e000345 | 19:58:44.333718 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/2/link_state |
| e000346 | 19:58:44.333718 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/1/nodes/3/flags |
| e000347 | 19:58:44.333718 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/3/primary_id |
| e000348 | 19:58:44.333718 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/1/nodes/3/slots |
| e000349 | 19:58:44.333718 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/3/link_state |
| e000350 | 19:58:44.333718 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/1/nodes/4/flags |
| e000351 | 19:58:44.333718 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/0/snapshot/views/1/nodes/4/primary_id |
| e000352 | 19:58:44.333718 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/4/slots |
| e000353 | 19:58:44.333718 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/4/link_state |
| e000354 | 19:58:44.333718 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/0/snapshot/views/1/nodes/5/flags |
| e000355 | 19:58:44.333718 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/1/nodes/5/primary_id |
| e000356 | 19:58:44.333718 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/1/nodes/5/slots |
| e000357 | 19:58:44.333718 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/1/nodes/5/link_state |
| e000358 | 19:58:44.333718 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/1/persistence/aof_enabled |
| e000359 | 19:58:44.333718 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/1/persistence/loading |
| e000360 | 19:58:44.333718 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/1/persistence/aof_last_write_status |
| e000361 | 19:58:44.333718 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e000362 | 19:58:44.333718 | node-2 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/1/process_live |
| e000363 | 19:58:44.333718 | node-4 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/3/reachable |
| e000364 | 19:58:44.333718 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/0/snapshot/views/3/replication/role |
| e000365 | 19:58:44.333718 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/0/snapshot/views/3/replication/connected_slaves |
| e000366 | 19:58:44.333718 | node-4 | info_replication | replication.master_link_status | valid | `"down"` | /samples/0/snapshot/views/3/replication/master_link_status |
| e000367 | 19:58:44.333718 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/0/snapshot/views/3/replication/master_sync_in_progress |
| e000368 | 19:58:44.333718 | node-4 | info_replication | replication.master_repl_offset | valid | `0` | /samples/0/snapshot/views/3/replication/master_repl_offset |
| e000369 | 19:58:44.333718 | node-4 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/0/snapshot/views/3/replication/slave_repl_offset |
| e000370 | 19:58:44.333718 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/0/snapshot/views/3/redis_id |
| e000371 | 19:58:44.333718 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/3/coverage |
| e000372 | 19:58:44.333718 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/3/nodes/0/flags |
| e000373 | 19:58:44.333718 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/0/primary_id |
| e000374 | 19:58:44.333718 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/3/nodes/0/slots |
| e000375 | 19:58:44.333718 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/0/link_state |
| e000376 | 19:58:44.333718 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/0/snapshot/views/3/nodes/1/flags |
| e000377 | 19:58:44.333718 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/0/snapshot/views/3/nodes/1/primary_id |
| e000378 | 19:58:44.333718 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/1/slots |
| e000379 | 19:58:44.333718 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/1/link_state |
| e000380 | 19:58:44.333718 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/3/nodes/2/flags |
| e000381 | 19:58:44.333718 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/0/snapshot/views/3/nodes/2/primary_id |
| e000382 | 19:58:44.333718 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/2/slots |
| e000383 | 19:58:44.333718 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/2/link_state |
| e000384 | 19:58:44.333718 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/3/nodes/3/flags |
| e000385 | 19:58:44.333718 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/3/primary_id |
| e000386 | 19:58:44.333718 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/3/nodes/3/slots |
| e000387 | 19:58:44.333718 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/3/link_state |
| e000388 | 19:58:44.333718 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/3/nodes/4/flags |
| e000389 | 19:58:44.333718 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/0/snapshot/views/3/nodes/4/primary_id |
| e000390 | 19:58:44.333718 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/4/slots |
| e000391 | 19:58:44.333718 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/4/link_state |
| e000392 | 19:58:44.333718 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/3/nodes/5/flags |
| e000393 | 19:58:44.333718 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/3/nodes/5/primary_id |
| e000394 | 19:58:44.333718 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/3/nodes/5/slots |
| e000395 | 19:58:44.333718 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/3/nodes/5/link_state |
| e000396 | 19:58:44.333718 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/3/persistence/aof_enabled |
| e000397 | 19:58:44.333718 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/3/persistence/loading |
| e000398 | 19:58:44.333718 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/3/persistence/aof_last_write_status |
| e000399 | 19:58:44.333718 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e000400 | 19:58:44.333718 | node-4 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/3/process_live |
| e000401 | 19:58:44.333725 | node-6 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/5/reachable |
| e000402 | 19:58:44.333725 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/0/snapshot/views/5/replication/role |
| e000403 | 19:58:44.333725 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/0/snapshot/views/5/replication/connected_slaves |
| e000404 | 19:58:44.333725 | node-6 | info_replication | replication.master_link_status | valid | `"down"` | /samples/0/snapshot/views/5/replication/master_link_status |
| e000405 | 19:58:44.333725 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/0/snapshot/views/5/replication/master_sync_in_progress |
| e000406 | 19:58:44.333725 | node-6 | info_replication | replication.master_repl_offset | valid | `0` | /samples/0/snapshot/views/5/replication/master_repl_offset |
| e000407 | 19:58:44.333725 | node-6 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/0/snapshot/views/5/replication/slave_repl_offset |
| e000408 | 19:58:44.333725 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/0/snapshot/views/5/redis_id |
| e000409 | 19:58:44.333725 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/5/coverage |
| e000410 | 19:58:44.333725 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/5/nodes/0/flags |
| e000411 | 19:58:44.333725 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/0/snapshot/views/5/nodes/0/primary_id |
| e000412 | 19:58:44.333725 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/0/slots |
| e000413 | 19:58:44.333725 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/0/link_state |
| e000414 | 19:58:44.333725 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/5/nodes/1/flags |
| e000415 | 19:58:44.333725 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/1/primary_id |
| e000416 | 19:58:44.333725 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/5/nodes/1/slots |
| e000417 | 19:58:44.333725 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/1/link_state |
| e000418 | 19:58:44.333725 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/5/nodes/2/flags |
| e000419 | 19:58:44.333725 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/2/primary_id |
| e000420 | 19:58:44.333725 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/5/nodes/2/slots |
| e000421 | 19:58:44.333725 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/2/link_state |
| e000422 | 19:58:44.333725 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/5/nodes/3/flags |
| e000423 | 19:58:44.333725 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/3/primary_id |
| e000424 | 19:58:44.333725 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/5/nodes/3/slots |
| e000425 | 19:58:44.333725 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/3/link_state |
| e000426 | 19:58:44.333725 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/5/nodes/4/flags |
| e000427 | 19:58:44.333725 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/0/snapshot/views/5/nodes/4/primary_id |
| e000428 | 19:58:44.333725 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/4/slots |
| e000429 | 19:58:44.333725 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/4/link_state |
| e000430 | 19:58:44.333725 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/0/snapshot/views/5/nodes/5/flags |
| e000431 | 19:58:44.333725 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/0/snapshot/views/5/nodes/5/primary_id |
| e000432 | 19:58:44.333725 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/5/nodes/5/slots |
| e000433 | 19:58:44.333725 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/5/nodes/5/link_state |
| e000434 | 19:58:44.333725 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/5/persistence/aof_enabled |
| e000435 | 19:58:44.333725 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/5/persistence/loading |
| e000436 | 19:58:44.333725 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/5/persistence/aof_last_write_status |
| e000437 | 19:58:44.333725 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e000438 | 19:58:44.333725 | node-6 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/5/process_live |
| e000439 | 19:58:44.333762 | node-5 | node_client | reachable | valid | `true` | /samples/0/snapshot/views/4/reachable |
| e000440 | 19:58:44.333762 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/0/snapshot/views/4/replication/role |
| e000441 | 19:58:44.333762 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/0/snapshot/views/4/replication/connected_slaves |
| e000442 | 19:58:44.333762 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/0/snapshot/views/4/replication/master_link_status |
| e000443 | 19:58:44.333762 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/0/snapshot/views/4/replication/master_sync_in_progress |
| e000444 | 19:58:44.333762 | node-5 | info_replication | replication.master_repl_offset | valid | `0` | /samples/0/snapshot/views/4/replication/master_repl_offset |
| e000445 | 19:58:44.333762 | node-5 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/0/snapshot/views/4/replication/slave_repl_offset |
| e000446 | 19:58:44.333762 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/0/snapshot/views/4/redis_id |
| e000447 | 19:58:44.333762 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/0/snapshot/views/4/coverage |
| e000448 | 19:58:44.333762 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/4/nodes/0/flags |
| e000449 | 19:58:44.333762 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/0/primary_id |
| e000450 | 19:58:44.333762 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/0/snapshot/views/4/nodes/0/slots |
| e000451 | 19:58:44.333762 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/0/link_state |
| e000452 | 19:58:44.333762 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/4/nodes/1/flags |
| e000453 | 19:58:44.333762 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/0/snapshot/views/4/nodes/1/primary_id |
| e000454 | 19:58:44.333762 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/1/slots |
| e000455 | 19:58:44.333762 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/1/link_state |
| e000456 | 19:58:44.333762 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/4/nodes/2/flags |
| e000457 | 19:58:44.333762 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/2/primary_id |
| e000458 | 19:58:44.333762 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/0/snapshot/views/4/nodes/2/slots |
| e000459 | 19:58:44.333762 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/2/link_state |
| e000460 | 19:58:44.333762 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/0/snapshot/views/4/nodes/3/flags |
| e000461 | 19:58:44.333762 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/3/primary_id |
| e000462 | 19:58:44.333762 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/0/snapshot/views/4/nodes/3/slots |
| e000463 | 19:58:44.333762 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/3/link_state |
| e000464 | 19:58:44.333762 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/0/snapshot/views/4/nodes/4/flags |
| e000465 | 19:58:44.333762 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/0/snapshot/views/4/nodes/4/primary_id |
| e000466 | 19:58:44.333762 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/4/slots |
| e000467 | 19:58:44.333762 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/4/link_state |
| e000468 | 19:58:44.333762 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/0/snapshot/views/4/nodes/5/flags |
| e000469 | 19:58:44.333762 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/0/snapshot/views/4/nodes/5/primary_id |
| e000470 | 19:58:44.333762 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/0/snapshot/views/4/nodes/5/slots |
| e000471 | 19:58:44.333762 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/0/snapshot/views/4/nodes/5/link_state |
| e000472 | 19:58:44.333762 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/0/snapshot/views/4/persistence/aof_enabled |
| e000473 | 19:58:44.333762 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/0/snapshot/views/4/persistence/loading |
| e000474 | 19:58:44.333762 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/0/snapshot/views/4/persistence/aof_last_write_status |
| e000475 | 19:58:44.333762 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/0/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e000476 | 19:58:44.333762 | node-5 | admin | process_live | missing_from_schema | `` | /samples/0/snapshot/views/4/process_live |
| e000477 | 19:58:44.930151 | node-1 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/0/reachable |
| e000478 | 19:58:44.930151 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/1/snapshot/views/0/replication/role |
| e000479 | 19:58:44.930151 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/1/snapshot/views/0/replication/connected_slaves |
| e000480 | 19:58:44.930151 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/1/snapshot/views/0/replication/master_link_status |
| e000481 | 19:58:44.930151 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/1/snapshot/views/0/replication/master_sync_in_progress |
| e000482 | 19:58:44.930151 | node-1 | info_replication | replication.master_repl_offset | valid | `14` | /samples/1/snapshot/views/0/replication/master_repl_offset |
| e000483 | 19:58:44.930151 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/1/snapshot/views/0/replication/slave_repl_offset |
| e000484 | 19:58:44.930151 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/1/snapshot/views/0/redis_id |
| e000485 | 19:58:44.930151 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/0/coverage |
| e000486 | 19:58:44.930151 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/0/nodes/0/flags |
| e000487 | 19:58:44.930151 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/1/snapshot/views/0/nodes/0/primary_id |
| e000488 | 19:58:44.930151 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/0/slots |
| e000489 | 19:58:44.930151 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/0/link_state |
| e000490 | 19:58:44.930151 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/0/nodes/1/flags |
| e000491 | 19:58:44.930151 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/1/primary_id |
| e000492 | 19:58:44.930151 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/0/nodes/1/slots |
| e000493 | 19:58:44.930151 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/1/link_state |
| e000494 | 19:58:44.930151 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/0/nodes/2/flags |
| e000495 | 19:58:44.930151 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/2/primary_id |
| e000496 | 19:58:44.930151 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/0/nodes/2/slots |
| e000497 | 19:58:44.930151 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/2/link_state |
| e000498 | 19:58:44.930151 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/1/snapshot/views/0/nodes/3/flags |
| e000499 | 19:58:44.930151 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/3/primary_id |
| e000500 | 19:58:44.930151 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/0/nodes/3/slots |
| e000501 | 19:58:44.930151 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/3/link_state |
| e000502 | 19:58:44.930151 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/0/nodes/4/flags |
| e000503 | 19:58:44.930151 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/1/snapshot/views/0/nodes/4/primary_id |
| e000504 | 19:58:44.930151 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/4/slots |
| e000505 | 19:58:44.930151 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/4/link_state |
| e000506 | 19:58:44.930151 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/0/nodes/5/flags |
| e000507 | 19:58:44.930151 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/1/snapshot/views/0/nodes/5/primary_id |
| e000508 | 19:58:44.930151 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/0/nodes/5/slots |
| e000509 | 19:58:44.930151 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/0/nodes/5/link_state |
| e000510 | 19:58:44.930151 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/0/persistence/aof_enabled |
| e000511 | 19:58:44.930151 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/0/persistence/loading |
| e000512 | 19:58:44.930151 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/0/persistence/aof_last_write_status |
| e000513 | 19:58:44.930151 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e000514 | 19:58:44.930151 | node-1 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/0/process_live |
| e000515 | 19:58:44.930174 | node-5 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/4/reachable |
| e000516 | 19:58:44.930174 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/1/snapshot/views/4/replication/role |
| e000517 | 19:58:44.930174 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/1/snapshot/views/4/replication/connected_slaves |
| e000518 | 19:58:44.930174 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/1/snapshot/views/4/replication/master_link_status |
| e000519 | 19:58:44.930174 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/1/snapshot/views/4/replication/master_sync_in_progress |
| e000520 | 19:58:44.930174 | node-5 | info_replication | replication.master_repl_offset | valid | `0` | /samples/1/snapshot/views/4/replication/master_repl_offset |
| e000521 | 19:58:44.930174 | node-5 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/1/snapshot/views/4/replication/slave_repl_offset |
| e000522 | 19:58:44.930174 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/1/snapshot/views/4/redis_id |
| e000523 | 19:58:44.930174 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/4/coverage |
| e000524 | 19:58:44.930174 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/4/nodes/0/flags |
| e000525 | 19:58:44.930174 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/0/primary_id |
| e000526 | 19:58:44.930174 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/4/nodes/0/slots |
| e000527 | 19:58:44.930174 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/0/link_state |
| e000528 | 19:58:44.930174 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/4/nodes/1/flags |
| e000529 | 19:58:44.930174 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/1/snapshot/views/4/nodes/1/primary_id |
| e000530 | 19:58:44.930174 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/1/slots |
| e000531 | 19:58:44.930174 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/1/link_state |
| e000532 | 19:58:44.930174 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/4/nodes/2/flags |
| e000533 | 19:58:44.930174 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/2/primary_id |
| e000534 | 19:58:44.930174 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/4/nodes/2/slots |
| e000535 | 19:58:44.930174 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/2/link_state |
| e000536 | 19:58:44.930174 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/4/nodes/3/flags |
| e000537 | 19:58:44.930174 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/3/primary_id |
| e000538 | 19:58:44.930174 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/4/nodes/3/slots |
| e000539 | 19:58:44.930174 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/3/link_state |
| e000540 | 19:58:44.930174 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/1/snapshot/views/4/nodes/4/flags |
| e000541 | 19:58:44.930174 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/1/snapshot/views/4/nodes/4/primary_id |
| e000542 | 19:58:44.930174 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/4/slots |
| e000543 | 19:58:44.930174 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/4/link_state |
| e000544 | 19:58:44.930174 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/4/nodes/5/flags |
| e000545 | 19:58:44.930174 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/1/snapshot/views/4/nodes/5/primary_id |
| e000546 | 19:58:44.930174 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/4/nodes/5/slots |
| e000547 | 19:58:44.930174 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/4/nodes/5/link_state |
| e000548 | 19:58:44.930174 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/4/persistence/aof_enabled |
| e000549 | 19:58:44.930174 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/4/persistence/loading |
| e000550 | 19:58:44.930174 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/4/persistence/aof_last_write_status |
| e000551 | 19:58:44.930174 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e000552 | 19:58:44.930174 | node-5 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/4/process_live |
| e000553 | 19:58:44.930178 | node-3 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/2/reachable |
| e000554 | 19:58:44.930178 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/1/snapshot/views/2/replication/role |
| e000555 | 19:58:44.930178 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/1/snapshot/views/2/replication/connected_slaves |
| e000556 | 19:58:44.930178 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/1/snapshot/views/2/replication/master_link_status |
| e000557 | 19:58:44.930178 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/1/snapshot/views/2/replication/master_sync_in_progress |
| e000558 | 19:58:44.930178 | node-3 | info_replication | replication.master_repl_offset | valid | `0` | /samples/1/snapshot/views/2/replication/master_repl_offset |
| e000559 | 19:58:44.930178 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/1/snapshot/views/2/replication/slave_repl_offset |
| e000560 | 19:58:44.930178 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/1/snapshot/views/2/redis_id |
| e000561 | 19:58:44.930178 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/2/coverage |
| e000562 | 19:58:44.930178 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/2/nodes/0/flags |
| e000563 | 19:58:44.930178 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/0/primary_id |
| e000564 | 19:58:44.930178 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/2/nodes/0/slots |
| e000565 | 19:58:44.930178 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/0/link_state |
| e000566 | 19:58:44.930178 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/1/snapshot/views/2/nodes/1/flags |
| e000567 | 19:58:44.930178 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/1/primary_id |
| e000568 | 19:58:44.930178 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/2/nodes/1/slots |
| e000569 | 19:58:44.930178 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/1/link_state |
| e000570 | 19:58:44.930178 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/2/nodes/2/flags |
| e000571 | 19:58:44.930178 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/1/snapshot/views/2/nodes/2/primary_id |
| e000572 | 19:58:44.930178 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/2/slots |
| e000573 | 19:58:44.930178 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/2/link_state |
| e000574 | 19:58:44.930178 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/2/nodes/3/flags |
| e000575 | 19:58:44.930178 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/1/snapshot/views/2/nodes/3/primary_id |
| e000576 | 19:58:44.930178 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/3/slots |
| e000577 | 19:58:44.930178 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/3/link_state |
| e000578 | 19:58:44.930178 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/2/nodes/4/flags |
| e000579 | 19:58:44.930178 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/1/snapshot/views/2/nodes/4/primary_id |
| e000580 | 19:58:44.930178 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/4/slots |
| e000581 | 19:58:44.930178 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/4/link_state |
| e000582 | 19:58:44.930178 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/2/nodes/5/flags |
| e000583 | 19:58:44.930178 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/2/nodes/5/primary_id |
| e000584 | 19:58:44.930178 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/2/nodes/5/slots |
| e000585 | 19:58:44.930178 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/2/nodes/5/link_state |
| e000586 | 19:58:44.930178 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/2/persistence/aof_enabled |
| e000587 | 19:58:44.930178 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/2/persistence/loading |
| e000588 | 19:58:44.930178 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/2/persistence/aof_last_write_status |
| e000589 | 19:58:44.930178 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e000590 | 19:58:44.930178 | node-3 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/2/process_live |
| e000591 | 19:58:44.930181 | node-2 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/1/reachable |
| e000592 | 19:58:44.930181 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/1/snapshot/views/1/replication/role |
| e000593 | 19:58:44.930181 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/1/snapshot/views/1/replication/connected_slaves |
| e000594 | 19:58:44.930181 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/1/snapshot/views/1/replication/master_link_status |
| e000595 | 19:58:44.930181 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/1/snapshot/views/1/replication/master_sync_in_progress |
| e000596 | 19:58:44.930181 | node-2 | info_replication | replication.master_repl_offset | valid | `0` | /samples/1/snapshot/views/1/replication/master_repl_offset |
| e000597 | 19:58:44.930181 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/1/snapshot/views/1/replication/slave_repl_offset |
| e000598 | 19:58:44.930181 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/1/snapshot/views/1/redis_id |
| e000599 | 19:58:44.930181 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/1/coverage |
| e000600 | 19:58:44.930181 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/1/nodes/0/flags |
| e000601 | 19:58:44.930181 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/1/snapshot/views/1/nodes/0/primary_id |
| e000602 | 19:58:44.930181 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/0/slots |
| e000603 | 19:58:44.930181 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/0/link_state |
| e000604 | 19:58:44.930181 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/1/nodes/1/flags |
| e000605 | 19:58:44.930181 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/1/primary_id |
| e000606 | 19:58:44.930181 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/1/nodes/1/slots |
| e000607 | 19:58:44.930181 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/1/link_state |
| e000608 | 19:58:44.930181 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/1/nodes/2/flags |
| e000609 | 19:58:44.930181 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/1/snapshot/views/1/nodes/2/primary_id |
| e000610 | 19:58:44.930181 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/2/slots |
| e000611 | 19:58:44.930181 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/2/link_state |
| e000612 | 19:58:44.930181 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/1/nodes/3/flags |
| e000613 | 19:58:44.930181 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/3/primary_id |
| e000614 | 19:58:44.930181 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/1/nodes/3/slots |
| e000615 | 19:58:44.930181 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/3/link_state |
| e000616 | 19:58:44.930181 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/1/nodes/4/flags |
| e000617 | 19:58:44.930181 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/1/snapshot/views/1/nodes/4/primary_id |
| e000618 | 19:58:44.930181 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/4/slots |
| e000619 | 19:58:44.930181 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/4/link_state |
| e000620 | 19:58:44.930181 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/1/snapshot/views/1/nodes/5/flags |
| e000621 | 19:58:44.930181 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/1/nodes/5/primary_id |
| e000622 | 19:58:44.930181 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/1/nodes/5/slots |
| e000623 | 19:58:44.930181 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/1/nodes/5/link_state |
| e000624 | 19:58:44.930181 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/1/persistence/aof_enabled |
| e000625 | 19:58:44.930181 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/1/persistence/loading |
| e000626 | 19:58:44.930181 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/1/persistence/aof_last_write_status |
| e000627 | 19:58:44.930181 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e000628 | 19:58:44.930181 | node-2 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/1/process_live |
| e000629 | 19:58:44.930187 | node-6 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/5/reachable |
| e000630 | 19:58:44.930187 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/1/snapshot/views/5/replication/role |
| e000631 | 19:58:44.930187 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/1/snapshot/views/5/replication/connected_slaves |
| e000632 | 19:58:44.930187 | node-6 | info_replication | replication.master_link_status | valid | `"down"` | /samples/1/snapshot/views/5/replication/master_link_status |
| e000633 | 19:58:44.930187 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/1/snapshot/views/5/replication/master_sync_in_progress |
| e000634 | 19:58:44.930187 | node-6 | info_replication | replication.master_repl_offset | valid | `0` | /samples/1/snapshot/views/5/replication/master_repl_offset |
| e000635 | 19:58:44.930187 | node-6 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/1/snapshot/views/5/replication/slave_repl_offset |
| e000636 | 19:58:44.930187 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/1/snapshot/views/5/redis_id |
| e000637 | 19:58:44.930187 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/5/coverage |
| e000638 | 19:58:44.930187 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/5/nodes/0/flags |
| e000639 | 19:58:44.930187 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/1/snapshot/views/5/nodes/0/primary_id |
| e000640 | 19:58:44.930187 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/0/slots |
| e000641 | 19:58:44.930187 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/0/link_state |
| e000642 | 19:58:44.930187 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/5/nodes/1/flags |
| e000643 | 19:58:44.930187 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/1/primary_id |
| e000644 | 19:58:44.930187 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/5/nodes/1/slots |
| e000645 | 19:58:44.930187 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/1/link_state |
| e000646 | 19:58:44.930187 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/5/nodes/2/flags |
| e000647 | 19:58:44.930187 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/2/primary_id |
| e000648 | 19:58:44.930187 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/5/nodes/2/slots |
| e000649 | 19:58:44.930187 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/2/link_state |
| e000650 | 19:58:44.930187 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/5/nodes/3/flags |
| e000651 | 19:58:44.930187 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/3/primary_id |
| e000652 | 19:58:44.930187 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/5/nodes/3/slots |
| e000653 | 19:58:44.930187 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/3/link_state |
| e000654 | 19:58:44.930187 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/5/nodes/4/flags |
| e000655 | 19:58:44.930187 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/1/snapshot/views/5/nodes/4/primary_id |
| e000656 | 19:58:44.930187 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/4/slots |
| e000657 | 19:58:44.930187 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/4/link_state |
| e000658 | 19:58:44.930187 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/1/snapshot/views/5/nodes/5/flags |
| e000659 | 19:58:44.930187 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/1/snapshot/views/5/nodes/5/primary_id |
| e000660 | 19:58:44.930187 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/5/nodes/5/slots |
| e000661 | 19:58:44.930187 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/5/nodes/5/link_state |
| e000662 | 19:58:44.930187 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/5/persistence/aof_enabled |
| e000663 | 19:58:44.930187 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/5/persistence/loading |
| e000664 | 19:58:44.930187 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/5/persistence/aof_last_write_status |
| e000665 | 19:58:44.930187 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e000666 | 19:58:44.930187 | node-6 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/5/process_live |
| e000667 | 19:58:44.930195 | node-4 | node_client | reachable | valid | `true` | /samples/1/snapshot/views/3/reachable |
| e000668 | 19:58:44.930195 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/1/snapshot/views/3/replication/role |
| e000669 | 19:58:44.930195 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/1/snapshot/views/3/replication/connected_slaves |
| e000670 | 19:58:44.930195 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/1/snapshot/views/3/replication/master_link_status |
| e000671 | 19:58:44.930195 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/1/snapshot/views/3/replication/master_sync_in_progress |
| e000672 | 19:58:44.930195 | node-4 | info_replication | replication.master_repl_offset | valid | `0` | /samples/1/snapshot/views/3/replication/master_repl_offset |
| e000673 | 19:58:44.930195 | node-4 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/1/snapshot/views/3/replication/slave_repl_offset |
| e000674 | 19:58:44.930195 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/1/snapshot/views/3/redis_id |
| e000675 | 19:58:44.930195 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/1/snapshot/views/3/coverage |
| e000676 | 19:58:44.930195 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/3/nodes/0/flags |
| e000677 | 19:58:44.930195 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/0/primary_id |
| e000678 | 19:58:44.930195 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/1/snapshot/views/3/nodes/0/slots |
| e000679 | 19:58:44.930195 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/0/link_state |
| e000680 | 19:58:44.930195 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/1/snapshot/views/3/nodes/1/flags |
| e000681 | 19:58:44.930195 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/1/snapshot/views/3/nodes/1/primary_id |
| e000682 | 19:58:44.930195 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/1/slots |
| e000683 | 19:58:44.930195 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/1/link_state |
| e000684 | 19:58:44.930195 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/3/nodes/2/flags |
| e000685 | 19:58:44.930195 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/1/snapshot/views/3/nodes/2/primary_id |
| e000686 | 19:58:44.930195 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/2/slots |
| e000687 | 19:58:44.930195 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/2/link_state |
| e000688 | 19:58:44.930195 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/3/nodes/3/flags |
| e000689 | 19:58:44.930195 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/3/primary_id |
| e000690 | 19:58:44.930195 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/1/snapshot/views/3/nodes/3/slots |
| e000691 | 19:58:44.930195 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/3/link_state |
| e000692 | 19:58:44.930195 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/1/snapshot/views/3/nodes/4/flags |
| e000693 | 19:58:44.930195 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/1/snapshot/views/3/nodes/4/primary_id |
| e000694 | 19:58:44.930195 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/4/slots |
| e000695 | 19:58:44.930195 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/4/link_state |
| e000696 | 19:58:44.930195 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/1/snapshot/views/3/nodes/5/flags |
| e000697 | 19:58:44.930195 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/1/snapshot/views/3/nodes/5/primary_id |
| e000698 | 19:58:44.930195 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/1/snapshot/views/3/nodes/5/slots |
| e000699 | 19:58:44.930195 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/1/snapshot/views/3/nodes/5/link_state |
| e000700 | 19:58:44.930195 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/1/snapshot/views/3/persistence/aof_enabled |
| e000701 | 19:58:44.930195 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/1/snapshot/views/3/persistence/loading |
| e000702 | 19:58:44.930195 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/1/snapshot/views/3/persistence/aof_last_write_status |
| e000703 | 19:58:44.930195 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/1/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e000704 | 19:58:44.930195 | node-4 | admin | process_live | missing_from_schema | `` | /samples/1/snapshot/views/3/process_live |
| e000705 | 19:58:45.521713 | node-3 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/2/reachable |
| e000706 | 19:58:45.521713 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/2/snapshot/views/2/replication/role |
| e000707 | 19:58:45.521713 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/2/snapshot/views/2/replication/connected_slaves |
| e000708 | 19:58:45.521713 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/2/snapshot/views/2/replication/master_link_status |
| e000709 | 19:58:45.521713 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/2/snapshot/views/2/replication/master_sync_in_progress |
| e000710 | 19:58:45.521713 | node-3 | info_replication | replication.master_repl_offset | valid | `0` | /samples/2/snapshot/views/2/replication/master_repl_offset |
| e000711 | 19:58:45.521713 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/2/snapshot/views/2/replication/slave_repl_offset |
| e000712 | 19:58:45.521713 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/2/snapshot/views/2/redis_id |
| e000713 | 19:58:45.521713 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/2/snapshot/views/2/coverage |
| e000714 | 19:58:45.521713 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/2/nodes/0/flags |
| e000715 | 19:58:45.521713 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/0/primary_id |
| e000716 | 19:58:45.521713 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/2/nodes/0/slots |
| e000717 | 19:58:45.521713 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/0/link_state |
| e000718 | 19:58:45.521713 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/2/snapshot/views/2/nodes/1/flags |
| e000719 | 19:58:45.521713 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/1/primary_id |
| e000720 | 19:58:45.521713 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/2/nodes/1/slots |
| e000721 | 19:58:45.521713 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/1/link_state |
| e000722 | 19:58:45.521713 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/2/nodes/2/flags |
| e000723 | 19:58:45.521713 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/2/snapshot/views/2/nodes/2/primary_id |
| e000724 | 19:58:45.521713 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/2/slots |
| e000725 | 19:58:45.521713 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/2/link_state |
| e000726 | 19:58:45.521713 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/2/nodes/3/flags |
| e000727 | 19:58:45.521713 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/2/snapshot/views/2/nodes/3/primary_id |
| e000728 | 19:58:45.521713 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/3/slots |
| e000729 | 19:58:45.521713 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/3/link_state |
| e000730 | 19:58:45.521713 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/2/nodes/4/flags |
| e000731 | 19:58:45.521713 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/2/snapshot/views/2/nodes/4/primary_id |
| e000732 | 19:58:45.521713 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/4/slots |
| e000733 | 19:58:45.521713 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/4/link_state |
| e000734 | 19:58:45.521713 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/2/nodes/5/flags |
| e000735 | 19:58:45.521713 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/2/nodes/5/primary_id |
| e000736 | 19:58:45.521713 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/2/nodes/5/slots |
| e000737 | 19:58:45.521713 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/2/nodes/5/link_state |
| e000738 | 19:58:45.521713 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/2/persistence/aof_enabled |
| e000739 | 19:58:45.521713 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/2/persistence/loading |
| e000740 | 19:58:45.521713 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/2/persistence/aof_last_write_status |
| e000741 | 19:58:45.521713 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e000742 | 19:58:45.521713 | node-3 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/2/process_live |
| e000743 | 19:58:45.521715 | node-2 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/1/reachable |
| e000744 | 19:58:45.521715 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/2/snapshot/views/1/replication/role |
| e000745 | 19:58:45.521715 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/2/snapshot/views/1/replication/connected_slaves |
| e000746 | 19:58:45.521715 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/2/snapshot/views/1/replication/master_link_status |
| e000747 | 19:58:45.521715 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/2/snapshot/views/1/replication/master_sync_in_progress |
| e000748 | 19:58:45.521715 | node-2 | info_replication | replication.master_repl_offset | valid | `0` | /samples/2/snapshot/views/1/replication/master_repl_offset |
| e000749 | 19:58:45.521715 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/2/snapshot/views/1/replication/slave_repl_offset |
| e000750 | 19:58:45.521715 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/2/snapshot/views/1/redis_id |
| e000751 | 19:58:45.521715 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/2/snapshot/views/1/coverage |
| e000752 | 19:58:45.521715 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/1/nodes/0/flags |
| e000753 | 19:58:45.521715 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/2/snapshot/views/1/nodes/0/primary_id |
| e000754 | 19:58:45.521715 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/0/slots |
| e000755 | 19:58:45.521715 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/0/link_state |
| e000756 | 19:58:45.521715 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/1/nodes/1/flags |
| e000757 | 19:58:45.521715 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/1/primary_id |
| e000758 | 19:58:45.521715 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/1/nodes/1/slots |
| e000759 | 19:58:45.521715 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/1/link_state |
| e000760 | 19:58:45.521715 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/1/nodes/2/flags |
| e000761 | 19:58:45.521715 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/2/snapshot/views/1/nodes/2/primary_id |
| e000762 | 19:58:45.521715 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/2/slots |
| e000763 | 19:58:45.521715 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/2/link_state |
| e000764 | 19:58:45.521715 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/1/nodes/3/flags |
| e000765 | 19:58:45.521715 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/3/primary_id |
| e000766 | 19:58:45.521715 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/1/nodes/3/slots |
| e000767 | 19:58:45.521715 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/3/link_state |
| e000768 | 19:58:45.521715 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/1/nodes/4/flags |
| e000769 | 19:58:45.521715 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/2/snapshot/views/1/nodes/4/primary_id |
| e000770 | 19:58:45.521715 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/4/slots |
| e000771 | 19:58:45.521715 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/4/link_state |
| e000772 | 19:58:45.521715 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/2/snapshot/views/1/nodes/5/flags |
| e000773 | 19:58:45.521715 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/1/nodes/5/primary_id |
| e000774 | 19:58:45.521715 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/1/nodes/5/slots |
| e000775 | 19:58:45.521715 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/1/nodes/5/link_state |
| e000776 | 19:58:45.521715 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/1/persistence/aof_enabled |
| e000777 | 19:58:45.521715 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/1/persistence/loading |
| e000778 | 19:58:45.521715 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/1/persistence/aof_last_write_status |
| e000779 | 19:58:45.521715 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e000780 | 19:58:45.521715 | node-2 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/1/process_live |
| e000781 | 19:58:45.521729 | node-6 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/5/reachable |
| e000782 | 19:58:45.521729 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/2/snapshot/views/5/replication/role |
| e000783 | 19:58:45.521729 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/2/snapshot/views/5/replication/connected_slaves |
| e000784 | 19:58:45.521729 | node-6 | info_replication | replication.master_link_status | valid | `"down"` | /samples/2/snapshot/views/5/replication/master_link_status |
| e000785 | 19:58:45.521729 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/2/snapshot/views/5/replication/master_sync_in_progress |
| e000786 | 19:58:45.521729 | node-6 | info_replication | replication.master_repl_offset | valid | `0` | /samples/2/snapshot/views/5/replication/master_repl_offset |
| e000787 | 19:58:45.521729 | node-6 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/2/snapshot/views/5/replication/slave_repl_offset |
| e000788 | 19:58:45.521729 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/2/snapshot/views/5/redis_id |
| e000789 | 19:58:45.521729 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/2/snapshot/views/5/coverage |
| e000790 | 19:58:45.521729 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/5/nodes/0/flags |
| e000791 | 19:58:45.521729 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/2/snapshot/views/5/nodes/0/primary_id |
| e000792 | 19:58:45.521729 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/0/slots |
| e000793 | 19:58:45.521729 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/0/link_state |
| e000794 | 19:58:45.521729 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/5/nodes/1/flags |
| e000795 | 19:58:45.521729 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/1/primary_id |
| e000796 | 19:58:45.521729 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/5/nodes/1/slots |
| e000797 | 19:58:45.521729 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/1/link_state |
| e000798 | 19:58:45.521729 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/5/nodes/2/flags |
| e000799 | 19:58:45.521729 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/2/primary_id |
| e000800 | 19:58:45.521729 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/5/nodes/2/slots |
| e000801 | 19:58:45.521729 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/2/link_state |
| e000802 | 19:58:45.521729 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/5/nodes/3/flags |
| e000803 | 19:58:45.521729 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/3/primary_id |
| e000804 | 19:58:45.521729 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/5/nodes/3/slots |
| e000805 | 19:58:45.521729 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/3/link_state |
| e000806 | 19:58:45.521729 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/5/nodes/4/flags |
| e000807 | 19:58:45.521729 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/2/snapshot/views/5/nodes/4/primary_id |
| e000808 | 19:58:45.521729 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/4/slots |
| e000809 | 19:58:45.521729 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/4/link_state |
| e000810 | 19:58:45.521729 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/2/snapshot/views/5/nodes/5/flags |
| e000811 | 19:58:45.521729 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/2/snapshot/views/5/nodes/5/primary_id |
| e000812 | 19:58:45.521729 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/5/nodes/5/slots |
| e000813 | 19:58:45.521729 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/5/nodes/5/link_state |
| e000814 | 19:58:45.521729 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/5/persistence/aof_enabled |
| e000815 | 19:58:45.521729 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/5/persistence/loading |
| e000816 | 19:58:45.521729 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/5/persistence/aof_last_write_status |
| e000817 | 19:58:45.521729 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e000818 | 19:58:45.521729 | node-6 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/5/process_live |
| e000819 | 19:58:45.521737 | node-4 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/3/reachable |
| e000820 | 19:58:45.521737 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/2/snapshot/views/3/replication/role |
| e000821 | 19:58:45.521737 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/2/snapshot/views/3/replication/connected_slaves |
| e000822 | 19:58:45.521737 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/2/snapshot/views/3/replication/master_link_status |
| e000823 | 19:58:45.521737 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/2/snapshot/views/3/replication/master_sync_in_progress |
| e000824 | 19:58:45.521737 | node-4 | info_replication | replication.master_repl_offset | valid | `0` | /samples/2/snapshot/views/3/replication/master_repl_offset |
| e000825 | 19:58:45.521737 | node-4 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/2/snapshot/views/3/replication/slave_repl_offset |
| e000826 | 19:58:45.521737 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/2/snapshot/views/3/redis_id |
| e000827 | 19:58:45.521737 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/2/snapshot/views/3/coverage |
| e000828 | 19:58:45.521737 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/3/nodes/0/flags |
| e000829 | 19:58:45.521737 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/0/primary_id |
| e000830 | 19:58:45.521737 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/3/nodes/0/slots |
| e000831 | 19:58:45.521737 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/0/link_state |
| e000832 | 19:58:45.521737 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/2/snapshot/views/3/nodes/1/flags |
| e000833 | 19:58:45.521737 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/2/snapshot/views/3/nodes/1/primary_id |
| e000834 | 19:58:45.521737 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/1/slots |
| e000835 | 19:58:45.521737 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/1/link_state |
| e000836 | 19:58:45.521737 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/3/nodes/2/flags |
| e000837 | 19:58:45.521737 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/2/snapshot/views/3/nodes/2/primary_id |
| e000838 | 19:58:45.521737 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/2/slots |
| e000839 | 19:58:45.521737 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/2/link_state |
| e000840 | 19:58:45.521737 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/3/nodes/3/flags |
| e000841 | 19:58:45.521737 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/3/primary_id |
| e000842 | 19:58:45.521737 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/3/nodes/3/slots |
| e000843 | 19:58:45.521737 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/3/link_state |
| e000844 | 19:58:45.521737 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/3/nodes/4/flags |
| e000845 | 19:58:45.521737 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/2/snapshot/views/3/nodes/4/primary_id |
| e000846 | 19:58:45.521737 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/4/slots |
| e000847 | 19:58:45.521737 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/4/link_state |
| e000848 | 19:58:45.521737 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/3/nodes/5/flags |
| e000849 | 19:58:45.521737 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/3/nodes/5/primary_id |
| e000850 | 19:58:45.521737 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/3/nodes/5/slots |
| e000851 | 19:58:45.521737 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/3/nodes/5/link_state |
| e000852 | 19:58:45.521737 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/3/persistence/aof_enabled |
| e000853 | 19:58:45.521737 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/3/persistence/loading |
| e000854 | 19:58:45.521737 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/3/persistence/aof_last_write_status |
| e000855 | 19:58:45.521737 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e000856 | 19:58:45.521737 | node-4 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/3/process_live |
| e000857 | 19:58:45.521753 | node-5 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/4/reachable |
| e000858 | 19:58:45.521753 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/2/snapshot/views/4/replication/role |
| e000859 | 19:58:45.521753 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/2/snapshot/views/4/replication/connected_slaves |
| e000860 | 19:58:45.521753 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/2/snapshot/views/4/replication/master_link_status |
| e000861 | 19:58:45.521753 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/2/snapshot/views/4/replication/master_sync_in_progress |
| e000862 | 19:58:45.521753 | node-5 | info_replication | replication.master_repl_offset | valid | `0` | /samples/2/snapshot/views/4/replication/master_repl_offset |
| e000863 | 19:58:45.521753 | node-5 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/2/snapshot/views/4/replication/slave_repl_offset |
| e000864 | 19:58:45.521753 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/2/snapshot/views/4/redis_id |
| e000865 | 19:58:45.521753 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/2/snapshot/views/4/coverage |
| e000866 | 19:58:45.521753 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/4/nodes/0/flags |
| e000867 | 19:58:45.521753 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/0/primary_id |
| e000868 | 19:58:45.521753 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/4/nodes/0/slots |
| e000869 | 19:58:45.521753 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/0/link_state |
| e000870 | 19:58:45.521753 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/4/nodes/1/flags |
| e000871 | 19:58:45.521753 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/2/snapshot/views/4/nodes/1/primary_id |
| e000872 | 19:58:45.521753 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/1/slots |
| e000873 | 19:58:45.521753 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/1/link_state |
| e000874 | 19:58:45.521753 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/4/nodes/2/flags |
| e000875 | 19:58:45.521753 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/2/primary_id |
| e000876 | 19:58:45.521753 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/4/nodes/2/slots |
| e000877 | 19:58:45.521753 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/2/link_state |
| e000878 | 19:58:45.521753 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/4/nodes/3/flags |
| e000879 | 19:58:45.521753 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/3/primary_id |
| e000880 | 19:58:45.521753 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/4/nodes/3/slots |
| e000881 | 19:58:45.521753 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/3/link_state |
| e000882 | 19:58:45.521753 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/2/snapshot/views/4/nodes/4/flags |
| e000883 | 19:58:45.521753 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/2/snapshot/views/4/nodes/4/primary_id |
| e000884 | 19:58:45.521753 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/4/slots |
| e000885 | 19:58:45.521753 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/4/link_state |
| e000886 | 19:58:45.521753 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/4/nodes/5/flags |
| e000887 | 19:58:45.521753 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/2/snapshot/views/4/nodes/5/primary_id |
| e000888 | 19:58:45.521753 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/4/nodes/5/slots |
| e000889 | 19:58:45.521753 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/4/nodes/5/link_state |
| e000890 | 19:58:45.521753 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/4/persistence/aof_enabled |
| e000891 | 19:58:45.521753 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/4/persistence/loading |
| e000892 | 19:58:45.521753 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/4/persistence/aof_last_write_status |
| e000893 | 19:58:45.521753 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e000894 | 19:58:45.521753 | node-5 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/4/process_live |
| e000895 | 19:58:45.52176 | node-1 | node_client | reachable | valid | `true` | /samples/2/snapshot/views/0/reachable |
| e000896 | 19:58:45.52176 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/2/snapshot/views/0/replication/role |
| e000897 | 19:58:45.52176 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/2/snapshot/views/0/replication/connected_slaves |
| e000898 | 19:58:45.52176 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/2/snapshot/views/0/replication/master_link_status |
| e000899 | 19:58:45.52176 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/2/snapshot/views/0/replication/master_sync_in_progress |
| e000900 | 19:58:45.52176 | node-1 | info_replication | replication.master_repl_offset | valid | `14` | /samples/2/snapshot/views/0/replication/master_repl_offset |
| e000901 | 19:58:45.52176 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/2/snapshot/views/0/replication/slave_repl_offset |
| e000902 | 19:58:45.52176 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/2/snapshot/views/0/redis_id |
| e000903 | 19:58:45.52176 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/2/snapshot/views/0/coverage |
| e000904 | 19:58:45.52176 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/0/nodes/0/flags |
| e000905 | 19:58:45.52176 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/2/snapshot/views/0/nodes/0/primary_id |
| e000906 | 19:58:45.52176 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/0/slots |
| e000907 | 19:58:45.52176 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/0/link_state |
| e000908 | 19:58:45.52176 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/0/nodes/1/flags |
| e000909 | 19:58:45.52176 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/1/primary_id |
| e000910 | 19:58:45.52176 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/2/snapshot/views/0/nodes/1/slots |
| e000911 | 19:58:45.52176 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/1/link_state |
| e000912 | 19:58:45.52176 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/2/snapshot/views/0/nodes/2/flags |
| e000913 | 19:58:45.52176 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/2/primary_id |
| e000914 | 19:58:45.52176 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/2/snapshot/views/0/nodes/2/slots |
| e000915 | 19:58:45.52176 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/2/link_state |
| e000916 | 19:58:45.52176 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/2/snapshot/views/0/nodes/3/flags |
| e000917 | 19:58:45.52176 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/3/primary_id |
| e000918 | 19:58:45.52176 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/2/snapshot/views/0/nodes/3/slots |
| e000919 | 19:58:45.52176 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/3/link_state |
| e000920 | 19:58:45.52176 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/0/nodes/4/flags |
| e000921 | 19:58:45.52176 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/2/snapshot/views/0/nodes/4/primary_id |
| e000922 | 19:58:45.52176 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/4/slots |
| e000923 | 19:58:45.52176 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/4/link_state |
| e000924 | 19:58:45.52176 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/2/snapshot/views/0/nodes/5/flags |
| e000925 | 19:58:45.52176 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/2/snapshot/views/0/nodes/5/primary_id |
| e000926 | 19:58:45.52176 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/2/snapshot/views/0/nodes/5/slots |
| e000927 | 19:58:45.52176 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/2/snapshot/views/0/nodes/5/link_state |
| e000928 | 19:58:45.52176 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/2/snapshot/views/0/persistence/aof_enabled |
| e000929 | 19:58:45.52176 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/2/snapshot/views/0/persistence/loading |
| e000930 | 19:58:45.52176 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/2/snapshot/views/0/persistence/aof_last_write_status |
| e000931 | 19:58:45.52176 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/2/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e000932 | 19:58:45.52176 | node-1 | admin | process_live | missing_from_schema | `` | /samples/2/snapshot/views/0/process_live |
| e000933 | 19:58:46.125869 | node-6 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/5/reachable |
| e000934 | 19:58:46.125869 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/3/snapshot/views/5/replication/role |
| e000935 | 19:58:46.125869 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/3/snapshot/views/5/replication/connected_slaves |
| e000936 | 19:58:46.125869 | node-6 | info_replication | replication.master_link_status | valid | `"down"` | /samples/3/snapshot/views/5/replication/master_link_status |
| e000937 | 19:58:46.125869 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/3/snapshot/views/5/replication/master_sync_in_progress |
| e000938 | 19:58:46.125869 | node-6 | info_replication | replication.master_repl_offset | valid | `0` | /samples/3/snapshot/views/5/replication/master_repl_offset |
| e000939 | 19:58:46.125869 | node-6 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/3/snapshot/views/5/replication/slave_repl_offset |
| e000940 | 19:58:46.125869 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/3/snapshot/views/5/redis_id |
| e000941 | 19:58:46.125869 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/5/coverage |
| e000942 | 19:58:46.125869 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/5/nodes/0/flags |
| e000943 | 19:58:46.125869 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/3/snapshot/views/5/nodes/0/primary_id |
| e000944 | 19:58:46.125869 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/0/slots |
| e000945 | 19:58:46.125869 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/0/link_state |
| e000946 | 19:58:46.125869 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/5/nodes/1/flags |
| e000947 | 19:58:46.125869 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/1/primary_id |
| e000948 | 19:58:46.125869 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/5/nodes/1/slots |
| e000949 | 19:58:46.125869 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/1/link_state |
| e000950 | 19:58:46.125869 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/5/nodes/2/flags |
| e000951 | 19:58:46.125869 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/2/primary_id |
| e000952 | 19:58:46.125869 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/5/nodes/2/slots |
| e000953 | 19:58:46.125869 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/2/link_state |
| e000954 | 19:58:46.125869 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/5/nodes/3/flags |
| e000955 | 19:58:46.125869 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/3/primary_id |
| e000956 | 19:58:46.125869 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/5/nodes/3/slots |
| e000957 | 19:58:46.125869 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/3/link_state |
| e000958 | 19:58:46.125869 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/5/nodes/4/flags |
| e000959 | 19:58:46.125869 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/3/snapshot/views/5/nodes/4/primary_id |
| e000960 | 19:58:46.125869 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/4/slots |
| e000961 | 19:58:46.125869 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/4/link_state |
| e000962 | 19:58:46.125869 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/3/snapshot/views/5/nodes/5/flags |
| e000963 | 19:58:46.125869 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/3/snapshot/views/5/nodes/5/primary_id |
| e000964 | 19:58:46.125869 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/5/nodes/5/slots |
| e000965 | 19:58:46.125869 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/5/nodes/5/link_state |
| e000966 | 19:58:46.125869 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/5/persistence/aof_enabled |
| e000967 | 19:58:46.125869 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/5/persistence/loading |
| e000968 | 19:58:46.125869 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/5/persistence/aof_last_write_status |
| e000969 | 19:58:46.125869 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e000970 | 19:58:46.125869 | node-6 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/5/process_live |
| e000971 | 19:58:46.125871 | node-4 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/3/reachable |
| e000972 | 19:58:46.125871 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/3/snapshot/views/3/replication/role |
| e000973 | 19:58:46.125871 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/3/snapshot/views/3/replication/connected_slaves |
| e000974 | 19:58:46.125871 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/3/snapshot/views/3/replication/master_link_status |
| e000975 | 19:58:46.125871 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/3/snapshot/views/3/replication/master_sync_in_progress |
| e000976 | 19:58:46.125871 | node-4 | info_replication | replication.master_repl_offset | valid | `0` | /samples/3/snapshot/views/3/replication/master_repl_offset |
| e000977 | 19:58:46.125871 | node-4 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/3/snapshot/views/3/replication/slave_repl_offset |
| e000978 | 19:58:46.125871 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/3/snapshot/views/3/redis_id |
| e000979 | 19:58:46.125871 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/3/coverage |
| e000980 | 19:58:46.125871 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/3/nodes/0/flags |
| e000981 | 19:58:46.125871 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/0/primary_id |
| e000982 | 19:58:46.125871 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/3/nodes/0/slots |
| e000983 | 19:58:46.125871 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/0/link_state |
| e000984 | 19:58:46.125871 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/3/snapshot/views/3/nodes/1/flags |
| e000985 | 19:58:46.125871 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/3/snapshot/views/3/nodes/1/primary_id |
| e000986 | 19:58:46.125871 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/1/slots |
| e000987 | 19:58:46.125871 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/1/link_state |
| e000988 | 19:58:46.125871 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/3/nodes/2/flags |
| e000989 | 19:58:46.125871 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/3/snapshot/views/3/nodes/2/primary_id |
| e000990 | 19:58:46.125871 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/2/slots |
| e000991 | 19:58:46.125871 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/2/link_state |
| e000992 | 19:58:46.125871 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/3/nodes/3/flags |
| e000993 | 19:58:46.125871 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/3/primary_id |
| e000994 | 19:58:46.125871 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/3/nodes/3/slots |
| e000995 | 19:58:46.125871 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/3/link_state |
| e000996 | 19:58:46.125871 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/3/nodes/4/flags |
| e000997 | 19:58:46.125871 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/3/snapshot/views/3/nodes/4/primary_id |
| e000998 | 19:58:46.125871 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/4/slots |
| e000999 | 19:58:46.125871 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/4/link_state |
| e001000 | 19:58:46.125871 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/3/nodes/5/flags |
| e001001 | 19:58:46.125871 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/3/nodes/5/primary_id |
| e001002 | 19:58:46.125871 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/3/nodes/5/slots |
| e001003 | 19:58:46.125871 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/3/nodes/5/link_state |
| e001004 | 19:58:46.125871 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/3/persistence/aof_enabled |
| e001005 | 19:58:46.125871 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/3/persistence/loading |
| e001006 | 19:58:46.125871 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/3/persistence/aof_last_write_status |
| e001007 | 19:58:46.125871 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001008 | 19:58:46.125871 | node-4 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/3/process_live |
| e001009 | 19:58:46.125872 | node-1 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/0/reachable |
| e001010 | 19:58:46.125872 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/3/snapshot/views/0/replication/role |
| e001011 | 19:58:46.125872 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/3/snapshot/views/0/replication/connected_slaves |
| e001012 | 19:58:46.125872 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/3/snapshot/views/0/replication/master_link_status |
| e001013 | 19:58:46.125872 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/3/snapshot/views/0/replication/master_sync_in_progress |
| e001014 | 19:58:46.125872 | node-1 | info_replication | replication.master_repl_offset | valid | `14` | /samples/3/snapshot/views/0/replication/master_repl_offset |
| e001015 | 19:58:46.125872 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/3/snapshot/views/0/replication/slave_repl_offset |
| e001016 | 19:58:46.125872 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/3/snapshot/views/0/redis_id |
| e001017 | 19:58:46.125872 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/0/coverage |
| e001018 | 19:58:46.125872 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/0/nodes/0/flags |
| e001019 | 19:58:46.125872 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/3/snapshot/views/0/nodes/0/primary_id |
| e001020 | 19:58:46.125872 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/0/slots |
| e001021 | 19:58:46.125872 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/0/link_state |
| e001022 | 19:58:46.125872 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/0/nodes/1/flags |
| e001023 | 19:58:46.125872 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/1/primary_id |
| e001024 | 19:58:46.125872 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/0/nodes/1/slots |
| e001025 | 19:58:46.125872 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/1/link_state |
| e001026 | 19:58:46.125872 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/0/nodes/2/flags |
| e001027 | 19:58:46.125872 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/2/primary_id |
| e001028 | 19:58:46.125872 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/0/nodes/2/slots |
| e001029 | 19:58:46.125872 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/2/link_state |
| e001030 | 19:58:46.125872 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/3/snapshot/views/0/nodes/3/flags |
| e001031 | 19:58:46.125872 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/3/primary_id |
| e001032 | 19:58:46.125872 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/0/nodes/3/slots |
| e001033 | 19:58:46.125872 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/3/link_state |
| e001034 | 19:58:46.125872 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/0/nodes/4/flags |
| e001035 | 19:58:46.125872 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/3/snapshot/views/0/nodes/4/primary_id |
| e001036 | 19:58:46.125872 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/4/slots |
| e001037 | 19:58:46.125872 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/4/link_state |
| e001038 | 19:58:46.125872 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/0/nodes/5/flags |
| e001039 | 19:58:46.125872 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/3/snapshot/views/0/nodes/5/primary_id |
| e001040 | 19:58:46.125872 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/0/nodes/5/slots |
| e001041 | 19:58:46.125872 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/0/nodes/5/link_state |
| e001042 | 19:58:46.125872 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/0/persistence/aof_enabled |
| e001043 | 19:58:46.125872 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/0/persistence/loading |
| e001044 | 19:58:46.125872 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/0/persistence/aof_last_write_status |
| e001045 | 19:58:46.125872 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001046 | 19:58:46.125872 | node-1 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/0/process_live |
| e001047 | 19:58:46.125886 | node-3 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/2/reachable |
| e001048 | 19:58:46.125886 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/3/snapshot/views/2/replication/role |
| e001049 | 19:58:46.125886 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/3/snapshot/views/2/replication/connected_slaves |
| e001050 | 19:58:46.125886 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/3/snapshot/views/2/replication/master_link_status |
| e001051 | 19:58:46.125886 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/3/snapshot/views/2/replication/master_sync_in_progress |
| e001052 | 19:58:46.125886 | node-3 | info_replication | replication.master_repl_offset | valid | `0` | /samples/3/snapshot/views/2/replication/master_repl_offset |
| e001053 | 19:58:46.125886 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/3/snapshot/views/2/replication/slave_repl_offset |
| e001054 | 19:58:46.125886 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/3/snapshot/views/2/redis_id |
| e001055 | 19:58:46.125886 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/2/coverage |
| e001056 | 19:58:46.125886 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/2/nodes/0/flags |
| e001057 | 19:58:46.125886 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/0/primary_id |
| e001058 | 19:58:46.125886 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/2/nodes/0/slots |
| e001059 | 19:58:46.125886 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/0/link_state |
| e001060 | 19:58:46.125886 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/3/snapshot/views/2/nodes/1/flags |
| e001061 | 19:58:46.125886 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/1/primary_id |
| e001062 | 19:58:46.125886 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/2/nodes/1/slots |
| e001063 | 19:58:46.125886 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/1/link_state |
| e001064 | 19:58:46.125886 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/2/nodes/2/flags |
| e001065 | 19:58:46.125886 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/3/snapshot/views/2/nodes/2/primary_id |
| e001066 | 19:58:46.125886 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/2/slots |
| e001067 | 19:58:46.125886 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/2/link_state |
| e001068 | 19:58:46.125886 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/2/nodes/3/flags |
| e001069 | 19:58:46.125886 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/3/snapshot/views/2/nodes/3/primary_id |
| e001070 | 19:58:46.125886 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/3/slots |
| e001071 | 19:58:46.125886 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/3/link_state |
| e001072 | 19:58:46.125886 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/2/nodes/4/flags |
| e001073 | 19:58:46.125886 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/3/snapshot/views/2/nodes/4/primary_id |
| e001074 | 19:58:46.125886 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/4/slots |
| e001075 | 19:58:46.125886 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/4/link_state |
| e001076 | 19:58:46.125886 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/2/nodes/5/flags |
| e001077 | 19:58:46.125886 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/2/nodes/5/primary_id |
| e001078 | 19:58:46.125886 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/2/nodes/5/slots |
| e001079 | 19:58:46.125886 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/2/nodes/5/link_state |
| e001080 | 19:58:46.125886 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/2/persistence/aof_enabled |
| e001081 | 19:58:46.125886 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/2/persistence/loading |
| e001082 | 19:58:46.125886 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/2/persistence/aof_last_write_status |
| e001083 | 19:58:46.125886 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001084 | 19:58:46.125886 | node-3 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/2/process_live |
| e001085 | 19:58:46.125922 | node-5 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/4/reachable |
| e001086 | 19:58:46.125922 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/3/snapshot/views/4/replication/role |
| e001087 | 19:58:46.125922 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/3/snapshot/views/4/replication/connected_slaves |
| e001088 | 19:58:46.125922 | node-5 | info_replication | replication.master_link_status | valid | `"down"` | /samples/3/snapshot/views/4/replication/master_link_status |
| e001089 | 19:58:46.125922 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/3/snapshot/views/4/replication/master_sync_in_progress |
| e001090 | 19:58:46.125922 | node-5 | info_replication | replication.master_repl_offset | valid | `0` | /samples/3/snapshot/views/4/replication/master_repl_offset |
| e001091 | 19:58:46.125922 | node-5 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/3/snapshot/views/4/replication/slave_repl_offset |
| e001092 | 19:58:46.125922 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/3/snapshot/views/4/redis_id |
| e001093 | 19:58:46.125922 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/4/coverage |
| e001094 | 19:58:46.125922 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/4/nodes/0/flags |
| e001095 | 19:58:46.125922 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/0/primary_id |
| e001096 | 19:58:46.125922 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/4/nodes/0/slots |
| e001097 | 19:58:46.125922 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/0/link_state |
| e001098 | 19:58:46.125922 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/4/nodes/1/flags |
| e001099 | 19:58:46.125922 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/3/snapshot/views/4/nodes/1/primary_id |
| e001100 | 19:58:46.125922 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/1/slots |
| e001101 | 19:58:46.125922 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/1/link_state |
| e001102 | 19:58:46.125922 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/4/nodes/2/flags |
| e001103 | 19:58:46.125922 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/2/primary_id |
| e001104 | 19:58:46.125922 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/4/nodes/2/slots |
| e001105 | 19:58:46.125922 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/2/link_state |
| e001106 | 19:58:46.125922 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/4/nodes/3/flags |
| e001107 | 19:58:46.125922 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/3/primary_id |
| e001108 | 19:58:46.125922 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/4/nodes/3/slots |
| e001109 | 19:58:46.125922 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/3/link_state |
| e001110 | 19:58:46.125922 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/3/snapshot/views/4/nodes/4/flags |
| e001111 | 19:58:46.125922 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/3/snapshot/views/4/nodes/4/primary_id |
| e001112 | 19:58:46.125922 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/4/slots |
| e001113 | 19:58:46.125922 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/4/link_state |
| e001114 | 19:58:46.125922 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/4/nodes/5/flags |
| e001115 | 19:58:46.125922 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/3/snapshot/views/4/nodes/5/primary_id |
| e001116 | 19:58:46.125922 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/4/nodes/5/slots |
| e001117 | 19:58:46.125922 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/4/nodes/5/link_state |
| e001118 | 19:58:46.125922 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/4/persistence/aof_enabled |
| e001119 | 19:58:46.125922 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/4/persistence/loading |
| e001120 | 19:58:46.125922 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/4/persistence/aof_last_write_status |
| e001121 | 19:58:46.125922 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001122 | 19:58:46.125922 | node-5 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/4/process_live |
| e001123 | 19:58:46.125936 | node-2 | node_client | reachable | valid | `true` | /samples/3/snapshot/views/1/reachable |
| e001124 | 19:58:46.125936 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/3/snapshot/views/1/replication/role |
| e001125 | 19:58:46.125936 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/3/snapshot/views/1/replication/connected_slaves |
| e001126 | 19:58:46.125936 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/3/snapshot/views/1/replication/master_link_status |
| e001127 | 19:58:46.125936 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/3/snapshot/views/1/replication/master_sync_in_progress |
| e001128 | 19:58:46.125936 | node-2 | info_replication | replication.master_repl_offset | valid | `0` | /samples/3/snapshot/views/1/replication/master_repl_offset |
| e001129 | 19:58:46.125936 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/3/snapshot/views/1/replication/slave_repl_offset |
| e001130 | 19:58:46.125936 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/3/snapshot/views/1/redis_id |
| e001131 | 19:58:46.125936 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/3/snapshot/views/1/coverage |
| e001132 | 19:58:46.125936 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/1/nodes/0/flags |
| e001133 | 19:58:46.125936 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/3/snapshot/views/1/nodes/0/primary_id |
| e001134 | 19:58:46.125936 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/0/slots |
| e001135 | 19:58:46.125936 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/0/link_state |
| e001136 | 19:58:46.125936 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/1/nodes/1/flags |
| e001137 | 19:58:46.125936 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/1/primary_id |
| e001138 | 19:58:46.125936 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/3/snapshot/views/1/nodes/1/slots |
| e001139 | 19:58:46.125936 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/1/link_state |
| e001140 | 19:58:46.125936 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/1/nodes/2/flags |
| e001141 | 19:58:46.125936 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/3/snapshot/views/1/nodes/2/primary_id |
| e001142 | 19:58:46.125936 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/2/slots |
| e001143 | 19:58:46.125936 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/2/link_state |
| e001144 | 19:58:46.125936 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/3/snapshot/views/1/nodes/3/flags |
| e001145 | 19:58:46.125936 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/3/primary_id |
| e001146 | 19:58:46.125936 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/3/snapshot/views/1/nodes/3/slots |
| e001147 | 19:58:46.125936 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/3/link_state |
| e001148 | 19:58:46.125936 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/3/snapshot/views/1/nodes/4/flags |
| e001149 | 19:58:46.125936 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/3/snapshot/views/1/nodes/4/primary_id |
| e001150 | 19:58:46.125936 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/4/slots |
| e001151 | 19:58:46.125936 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/4/link_state |
| e001152 | 19:58:46.125936 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/3/snapshot/views/1/nodes/5/flags |
| e001153 | 19:58:46.125936 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/3/snapshot/views/1/nodes/5/primary_id |
| e001154 | 19:58:46.125936 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/3/snapshot/views/1/nodes/5/slots |
| e001155 | 19:58:46.125936 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/3/snapshot/views/1/nodes/5/link_state |
| e001156 | 19:58:46.125936 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/3/snapshot/views/1/persistence/aof_enabled |
| e001157 | 19:58:46.125936 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/3/snapshot/views/1/persistence/loading |
| e001158 | 19:58:46.125936 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/3/snapshot/views/1/persistence/aof_last_write_status |
| e001159 | 19:58:46.125936 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/3/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001160 | 19:58:46.125936 | node-2 | admin | process_live | missing_from_schema | `` | /samples/3/snapshot/views/1/process_live |
| e001161 | 19:58:46.731454 | node-4 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/3/reachable |
| e001162 | 19:58:46.731454 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/4/snapshot/views/3/replication/role |
| e001163 | 19:58:46.731454 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/4/snapshot/views/3/replication/connected_slaves |
| e001164 | 19:58:46.731454 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/4/snapshot/views/3/replication/master_link_status |
| e001165 | 19:58:46.731454 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/4/snapshot/views/3/replication/master_sync_in_progress |
| e001166 | 19:58:46.731454 | node-4 | info_replication | replication.master_repl_offset | valid | `0` | /samples/4/snapshot/views/3/replication/master_repl_offset |
| e001167 | 19:58:46.731454 | node-4 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/4/snapshot/views/3/replication/slave_repl_offset |
| e001168 | 19:58:46.731454 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/4/snapshot/views/3/redis_id |
| e001169 | 19:58:46.731454 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/3/coverage |
| e001170 | 19:58:46.731454 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/3/nodes/0/flags |
| e001171 | 19:58:46.731454 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/0/primary_id |
| e001172 | 19:58:46.731454 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/3/nodes/0/slots |
| e001173 | 19:58:46.731454 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/0/link_state |
| e001174 | 19:58:46.731454 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/4/snapshot/views/3/nodes/1/flags |
| e001175 | 19:58:46.731454 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/4/snapshot/views/3/nodes/1/primary_id |
| e001176 | 19:58:46.731454 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/1/slots |
| e001177 | 19:58:46.731454 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/1/link_state |
| e001178 | 19:58:46.731454 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/3/nodes/2/flags |
| e001179 | 19:58:46.731454 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/4/snapshot/views/3/nodes/2/primary_id |
| e001180 | 19:58:46.731454 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/2/slots |
| e001181 | 19:58:46.731454 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/2/link_state |
| e001182 | 19:58:46.731454 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/3/nodes/3/flags |
| e001183 | 19:58:46.731454 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/3/primary_id |
| e001184 | 19:58:46.731454 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/3/nodes/3/slots |
| e001185 | 19:58:46.731454 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/3/link_state |
| e001186 | 19:58:46.731454 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/3/nodes/4/flags |
| e001187 | 19:58:46.731454 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/4/snapshot/views/3/nodes/4/primary_id |
| e001188 | 19:58:46.731454 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/4/slots |
| e001189 | 19:58:46.731454 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/4/link_state |
| e001190 | 19:58:46.731454 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/3/nodes/5/flags |
| e001191 | 19:58:46.731454 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/3/nodes/5/primary_id |
| e001192 | 19:58:46.731454 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/3/nodes/5/slots |
| e001193 | 19:58:46.731454 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/3/nodes/5/link_state |
| e001194 | 19:58:46.731454 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/3/persistence/aof_enabled |
| e001195 | 19:58:46.731454 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/3/persistence/loading |
| e001196 | 19:58:46.731454 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/3/persistence/aof_last_write_status |
| e001197 | 19:58:46.731454 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001198 | 19:58:46.731454 | node-4 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/3/process_live |
| e001199 | 19:58:46.732709 | node-1 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/0/reachable |
| e001200 | 19:58:46.732709 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/4/snapshot/views/0/replication/role |
| e001201 | 19:58:46.732709 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/4/snapshot/views/0/replication/connected_slaves |
| e001202 | 19:58:46.732709 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/4/snapshot/views/0/replication/master_link_status |
| e001203 | 19:58:46.732709 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/4/snapshot/views/0/replication/master_sync_in_progress |
| e001204 | 19:58:46.732709 | node-1 | info_replication | replication.master_repl_offset | valid | `14` | /samples/4/snapshot/views/0/replication/master_repl_offset |
| e001205 | 19:58:46.732709 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/4/snapshot/views/0/replication/slave_repl_offset |
| e001206 | 19:58:46.732709 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/4/snapshot/views/0/redis_id |
| e001207 | 19:58:46.732709 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/0/coverage |
| e001208 | 19:58:46.732709 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/0/nodes/0/flags |
| e001209 | 19:58:46.732709 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/4/snapshot/views/0/nodes/0/primary_id |
| e001210 | 19:58:46.732709 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/0/slots |
| e001211 | 19:58:46.732709 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/0/link_state |
| e001212 | 19:58:46.732709 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/0/nodes/1/flags |
| e001213 | 19:58:46.732709 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/1/primary_id |
| e001214 | 19:58:46.732709 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/0/nodes/1/slots |
| e001215 | 19:58:46.732709 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/1/link_state |
| e001216 | 19:58:46.732709 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/0/nodes/2/flags |
| e001217 | 19:58:46.732709 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/2/primary_id |
| e001218 | 19:58:46.732709 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/0/nodes/2/slots |
| e001219 | 19:58:46.732709 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/2/link_state |
| e001220 | 19:58:46.732709 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/4/snapshot/views/0/nodes/3/flags |
| e001221 | 19:58:46.732709 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/3/primary_id |
| e001222 | 19:58:46.732709 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/0/nodes/3/slots |
| e001223 | 19:58:46.732709 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/3/link_state |
| e001224 | 19:58:46.732709 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/0/nodes/4/flags |
| e001225 | 19:58:46.732709 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/4/snapshot/views/0/nodes/4/primary_id |
| e001226 | 19:58:46.732709 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/4/slots |
| e001227 | 19:58:46.732709 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/4/link_state |
| e001228 | 19:58:46.732709 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/0/nodes/5/flags |
| e001229 | 19:58:46.732709 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/4/snapshot/views/0/nodes/5/primary_id |
| e001230 | 19:58:46.732709 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/0/nodes/5/slots |
| e001231 | 19:58:46.732709 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/0/nodes/5/link_state |
| e001232 | 19:58:46.732709 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/0/persistence/aof_enabled |
| e001233 | 19:58:46.732709 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/0/persistence/loading |
| e001234 | 19:58:46.732709 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/0/persistence/aof_last_write_status |
| e001235 | 19:58:46.732709 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001236 | 19:58:46.732709 | node-1 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/0/process_live |
| e001237 | 19:58:46.733 | node-2 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/1/reachable |
| e001238 | 19:58:46.733 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/4/snapshot/views/1/replication/role |
| e001239 | 19:58:46.733 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/4/snapshot/views/1/replication/connected_slaves |
| e001240 | 19:58:46.733 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/4/snapshot/views/1/replication/master_link_status |
| e001241 | 19:58:46.733 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/4/snapshot/views/1/replication/master_sync_in_progress |
| e001242 | 19:58:46.733 | node-2 | info_replication | replication.master_repl_offset | valid | `14` | /samples/4/snapshot/views/1/replication/master_repl_offset |
| e001243 | 19:58:46.733 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/4/snapshot/views/1/replication/slave_repl_offset |
| e001244 | 19:58:46.733 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/4/snapshot/views/1/redis_id |
| e001245 | 19:58:46.733 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/1/coverage |
| e001246 | 19:58:46.733 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/1/nodes/0/flags |
| e001247 | 19:58:46.733 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/4/snapshot/views/1/nodes/0/primary_id |
| e001248 | 19:58:46.733 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/0/slots |
| e001249 | 19:58:46.733 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/0/link_state |
| e001250 | 19:58:46.733 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/1/nodes/1/flags |
| e001251 | 19:58:46.733 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/1/primary_id |
| e001252 | 19:58:46.733 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/1/nodes/1/slots |
| e001253 | 19:58:46.733 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/1/link_state |
| e001254 | 19:58:46.733 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/1/nodes/2/flags |
| e001255 | 19:58:46.733 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/4/snapshot/views/1/nodes/2/primary_id |
| e001256 | 19:58:46.733 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/2/slots |
| e001257 | 19:58:46.733 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/2/link_state |
| e001258 | 19:58:46.733 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/1/nodes/3/flags |
| e001259 | 19:58:46.733 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/3/primary_id |
| e001260 | 19:58:46.733 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/1/nodes/3/slots |
| e001261 | 19:58:46.733 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/3/link_state |
| e001262 | 19:58:46.733 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/1/nodes/4/flags |
| e001263 | 19:58:46.733 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/4/snapshot/views/1/nodes/4/primary_id |
| e001264 | 19:58:46.733 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/4/slots |
| e001265 | 19:58:46.733 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/4/link_state |
| e001266 | 19:58:46.733 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/4/snapshot/views/1/nodes/5/flags |
| e001267 | 19:58:46.733 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/1/nodes/5/primary_id |
| e001268 | 19:58:46.733 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/1/nodes/5/slots |
| e001269 | 19:58:46.733 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/1/nodes/5/link_state |
| e001270 | 19:58:46.733 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/1/persistence/aof_enabled |
| e001271 | 19:58:46.733 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/1/persistence/loading |
| e001272 | 19:58:46.733 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/1/persistence/aof_last_write_status |
| e001273 | 19:58:46.733 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001274 | 19:58:46.733 | node-2 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/1/process_live |
| e001275 | 19:58:46.733075 | node-3 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/2/reachable |
| e001276 | 19:58:46.733075 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/4/snapshot/views/2/replication/role |
| e001277 | 19:58:46.733075 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/4/snapshot/views/2/replication/connected_slaves |
| e001278 | 19:58:46.733075 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/4/snapshot/views/2/replication/master_link_status |
| e001279 | 19:58:46.733075 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/4/snapshot/views/2/replication/master_sync_in_progress |
| e001280 | 19:58:46.733075 | node-3 | info_replication | replication.master_repl_offset | valid | `0` | /samples/4/snapshot/views/2/replication/master_repl_offset |
| e001281 | 19:58:46.733075 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/4/snapshot/views/2/replication/slave_repl_offset |
| e001282 | 19:58:46.733075 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/4/snapshot/views/2/redis_id |
| e001283 | 19:58:46.733075 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/2/coverage |
| e001284 | 19:58:46.733075 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/2/nodes/0/flags |
| e001285 | 19:58:46.733075 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/0/primary_id |
| e001286 | 19:58:46.733075 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/2/nodes/0/slots |
| e001287 | 19:58:46.733075 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/0/link_state |
| e001288 | 19:58:46.733075 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/4/snapshot/views/2/nodes/1/flags |
| e001289 | 19:58:46.733075 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/1/primary_id |
| e001290 | 19:58:46.733075 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/2/nodes/1/slots |
| e001291 | 19:58:46.733075 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/1/link_state |
| e001292 | 19:58:46.733075 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/2/nodes/2/flags |
| e001293 | 19:58:46.733075 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/4/snapshot/views/2/nodes/2/primary_id |
| e001294 | 19:58:46.733075 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/2/slots |
| e001295 | 19:58:46.733075 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/2/link_state |
| e001296 | 19:58:46.733075 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/2/nodes/3/flags |
| e001297 | 19:58:46.733075 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/4/snapshot/views/2/nodes/3/primary_id |
| e001298 | 19:58:46.733075 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/3/slots |
| e001299 | 19:58:46.733075 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/3/link_state |
| e001300 | 19:58:46.733075 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/2/nodes/4/flags |
| e001301 | 19:58:46.733075 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/4/snapshot/views/2/nodes/4/primary_id |
| e001302 | 19:58:46.733075 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/4/slots |
| e001303 | 19:58:46.733075 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/4/link_state |
| e001304 | 19:58:46.733075 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/2/nodes/5/flags |
| e001305 | 19:58:46.733075 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/2/nodes/5/primary_id |
| e001306 | 19:58:46.733075 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/2/nodes/5/slots |
| e001307 | 19:58:46.733075 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/2/nodes/5/link_state |
| e001308 | 19:58:46.733075 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/2/persistence/aof_enabled |
| e001309 | 19:58:46.733075 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/2/persistence/loading |
| e001310 | 19:58:46.733075 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/2/persistence/aof_last_write_status |
| e001311 | 19:58:46.733075 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001312 | 19:58:46.733075 | node-3 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/2/process_live |
| e001313 | 19:58:46.733133 | node-5 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/4/reachable |
| e001314 | 19:58:46.733133 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/4/snapshot/views/4/replication/role |
| e001315 | 19:58:46.733133 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/4/snapshot/views/4/replication/connected_slaves |
| e001316 | 19:58:46.733133 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/4/snapshot/views/4/replication/master_link_status |
| e001317 | 19:58:46.733133 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/4/snapshot/views/4/replication/master_sync_in_progress |
| e001318 | 19:58:46.733133 | node-5 | info_replication | replication.master_repl_offset | valid | `14` | /samples/4/snapshot/views/4/replication/master_repl_offset |
| e001319 | 19:58:46.733133 | node-5 | info_replication | replication.slave_repl_offset | valid | `14` | /samples/4/snapshot/views/4/replication/slave_repl_offset |
| e001320 | 19:58:46.733133 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/4/snapshot/views/4/redis_id |
| e001321 | 19:58:46.733133 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/4/coverage |
| e001322 | 19:58:46.733133 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/4/nodes/0/flags |
| e001323 | 19:58:46.733133 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/0/primary_id |
| e001324 | 19:58:46.733133 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/4/nodes/0/slots |
| e001325 | 19:58:46.733133 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/4/nodes/0/link_state |
| e001326 | 19:58:46.733133 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/4/nodes/1/flags |
| e001327 | 19:58:46.733133 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/4/snapshot/views/4/nodes/1/primary_id |
| e001328 | 19:58:46.733133 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/1/slots |
| e001329 | 19:58:46.733133 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/4/nodes/1/link_state |
| e001330 | 19:58:46.733133 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/4/nodes/2/flags |
| e001331 | 19:58:46.733133 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/2/primary_id |
| e001332 | 19:58:46.733133 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/4/nodes/2/slots |
| e001333 | 19:58:46.733133 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/4/nodes/2/link_state |
| e001334 | 19:58:46.733133 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/4/nodes/3/flags |
| e001335 | 19:58:46.733133 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/3/primary_id |
| e001336 | 19:58:46.733133 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/4/nodes/3/slots |
| e001337 | 19:58:46.733133 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/4/nodes/3/link_state |
| e001338 | 19:58:46.733133 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/4/snapshot/views/4/nodes/4/flags |
| e001339 | 19:58:46.733133 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/4/snapshot/views/4/nodes/4/primary_id |
| e001340 | 19:58:46.733133 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/4/slots |
| e001341 | 19:58:46.733133 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/4/nodes/4/link_state |
| e001342 | 19:58:46.733133 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/4/nodes/5/flags |
| e001343 | 19:58:46.733133 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/4/snapshot/views/4/nodes/5/primary_id |
| e001344 | 19:58:46.733133 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/4/nodes/5/slots |
| e001345 | 19:58:46.733133 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/4/nodes/5/link_state |
| e001346 | 19:58:46.733133 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/4/persistence/aof_enabled |
| e001347 | 19:58:46.733133 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/4/persistence/loading |
| e001348 | 19:58:46.733133 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/4/persistence/aof_last_write_status |
| e001349 | 19:58:46.733133 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001350 | 19:58:46.733133 | node-5 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/4/process_live |
| e001351 | 19:58:46.733184 | node-6 | node_client | reachable | valid | `true` | /samples/4/snapshot/views/5/reachable |
| e001352 | 19:58:46.733184 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/4/snapshot/views/5/replication/role |
| e001353 | 19:58:46.733184 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/4/snapshot/views/5/replication/connected_slaves |
| e001354 | 19:58:46.733184 | node-6 | info_replication | replication.master_link_status | valid | `"down"` | /samples/4/snapshot/views/5/replication/master_link_status |
| e001355 | 19:58:46.733184 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/4/snapshot/views/5/replication/master_sync_in_progress |
| e001356 | 19:58:46.733184 | node-6 | info_replication | replication.master_repl_offset | valid | `0` | /samples/4/snapshot/views/5/replication/master_repl_offset |
| e001357 | 19:58:46.733184 | node-6 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/4/snapshot/views/5/replication/slave_repl_offset |
| e001358 | 19:58:46.733184 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/4/snapshot/views/5/redis_id |
| e001359 | 19:58:46.733184 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/4/snapshot/views/5/coverage |
| e001360 | 19:58:46.733184 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/5/nodes/0/flags |
| e001361 | 19:58:46.733184 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/4/snapshot/views/5/nodes/0/primary_id |
| e001362 | 19:58:46.733184 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/0/slots |
| e001363 | 19:58:46.733184 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/0/link_state |
| e001364 | 19:58:46.733184 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/5/nodes/1/flags |
| e001365 | 19:58:46.733184 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/1/primary_id |
| e001366 | 19:58:46.733184 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/4/snapshot/views/5/nodes/1/slots |
| e001367 | 19:58:46.733184 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/1/link_state |
| e001368 | 19:58:46.733184 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/5/nodes/2/flags |
| e001369 | 19:58:46.733184 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/2/primary_id |
| e001370 | 19:58:46.733184 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/4/snapshot/views/5/nodes/2/slots |
| e001371 | 19:58:46.733184 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/2/link_state |
| e001372 | 19:58:46.733184 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/4/snapshot/views/5/nodes/3/flags |
| e001373 | 19:58:46.733184 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/3/primary_id |
| e001374 | 19:58:46.733184 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/4/snapshot/views/5/nodes/3/slots |
| e001375 | 19:58:46.733184 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/3/link_state |
| e001376 | 19:58:46.733184 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/4/snapshot/views/5/nodes/4/flags |
| e001377 | 19:58:46.733184 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/4/snapshot/views/5/nodes/4/primary_id |
| e001378 | 19:58:46.733184 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/4/slots |
| e001379 | 19:58:46.733184 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/4/link_state |
| e001380 | 19:58:46.733184 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/4/snapshot/views/5/nodes/5/flags |
| e001381 | 19:58:46.733184 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/4/snapshot/views/5/nodes/5/primary_id |
| e001382 | 19:58:46.733184 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/4/snapshot/views/5/nodes/5/slots |
| e001383 | 19:58:46.733184 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/4/snapshot/views/5/nodes/5/link_state |
| e001384 | 19:58:46.733184 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/4/snapshot/views/5/persistence/aof_enabled |
| e001385 | 19:58:46.733184 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/4/snapshot/views/5/persistence/loading |
| e001386 | 19:58:46.733184 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/4/snapshot/views/5/persistence/aof_last_write_status |
| e001387 | 19:58:46.733184 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/4/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e001388 | 19:58:46.733184 | node-6 | admin | process_live | missing_from_schema | `` | /samples/4/snapshot/views/5/process_live |
| e001389 | 19:58:47.3548 | node-6 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/5/reachable |
| e001390 | 19:58:47.3548 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/5/snapshot/views/5/replication/role |
| e001391 | 19:58:47.3548 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/5/snapshot/views/5/replication/connected_slaves |
| e001392 | 19:58:47.3548 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/5/snapshot/views/5/replication/master_link_status |
| e001393 | 19:58:47.3548 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/5/snapshot/views/5/replication/master_sync_in_progress |
| e001394 | 19:58:47.3548 | node-6 | info_replication | replication.master_repl_offset | valid | `14` | /samples/5/snapshot/views/5/replication/master_repl_offset |
| e001395 | 19:58:47.3548 | node-6 | info_replication | replication.slave_repl_offset | valid | `14` | /samples/5/snapshot/views/5/replication/slave_repl_offset |
| e001396 | 19:58:47.3548 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/5/snapshot/views/5/redis_id |
| e001397 | 19:58:47.3548 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/5/coverage |
| e001398 | 19:58:47.3548 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/5/nodes/0/flags |
| e001399 | 19:58:47.3548 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/5/snapshot/views/5/nodes/0/primary_id |
| e001400 | 19:58:47.3548 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/0/slots |
| e001401 | 19:58:47.3548 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/0/link_state |
| e001402 | 19:58:47.3548 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/5/nodes/1/flags |
| e001403 | 19:58:47.3548 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/1/primary_id |
| e001404 | 19:58:47.3548 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/5/nodes/1/slots |
| e001405 | 19:58:47.3548 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/1/link_state |
| e001406 | 19:58:47.3548 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/5/nodes/2/flags |
| e001407 | 19:58:47.3548 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/2/primary_id |
| e001408 | 19:58:47.3548 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/5/nodes/2/slots |
| e001409 | 19:58:47.3548 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/2/link_state |
| e001410 | 19:58:47.3548 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/5/nodes/3/flags |
| e001411 | 19:58:47.3548 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/3/primary_id |
| e001412 | 19:58:47.3548 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/5/nodes/3/slots |
| e001413 | 19:58:47.3548 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/3/link_state |
| e001414 | 19:58:47.3548 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/5/nodes/4/flags |
| e001415 | 19:58:47.3548 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/5/snapshot/views/5/nodes/4/primary_id |
| e001416 | 19:58:47.3548 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/4/slots |
| e001417 | 19:58:47.3548 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/4/link_state |
| e001418 | 19:58:47.3548 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/5/snapshot/views/5/nodes/5/flags |
| e001419 | 19:58:47.3548 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/5/snapshot/views/5/nodes/5/primary_id |
| e001420 | 19:58:47.3548 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/5/nodes/5/slots |
| e001421 | 19:58:47.3548 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/5/nodes/5/link_state |
| e001422 | 19:58:47.3548 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/5/persistence/aof_enabled |
| e001423 | 19:58:47.3548 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/5/persistence/loading |
| e001424 | 19:58:47.3548 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/5/persistence/aof_last_write_status |
| e001425 | 19:58:47.3548 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e001426 | 19:58:47.3548 | node-6 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/5/process_live |
| e001427 | 19:58:47.354866 | node-5 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/4/reachable |
| e001428 | 19:58:47.354866 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/5/snapshot/views/4/replication/role |
| e001429 | 19:58:47.354866 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/5/snapshot/views/4/replication/connected_slaves |
| e001430 | 19:58:47.354866 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/5/snapshot/views/4/replication/master_link_status |
| e001431 | 19:58:47.354866 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/5/snapshot/views/4/replication/master_sync_in_progress |
| e001432 | 19:58:47.354866 | node-5 | info_replication | replication.master_repl_offset | valid | `14` | /samples/5/snapshot/views/4/replication/master_repl_offset |
| e001433 | 19:58:47.354866 | node-5 | info_replication | replication.slave_repl_offset | valid | `14` | /samples/5/snapshot/views/4/replication/slave_repl_offset |
| e001434 | 19:58:47.354866 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/5/snapshot/views/4/redis_id |
| e001435 | 19:58:47.354866 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/4/coverage |
| e001436 | 19:58:47.354866 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/4/nodes/0/flags |
| e001437 | 19:58:47.354866 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/0/primary_id |
| e001438 | 19:58:47.354866 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/4/nodes/0/slots |
| e001439 | 19:58:47.354866 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/0/link_state |
| e001440 | 19:58:47.354866 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/4/nodes/1/flags |
| e001441 | 19:58:47.354866 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/5/snapshot/views/4/nodes/1/primary_id |
| e001442 | 19:58:47.354866 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/1/slots |
| e001443 | 19:58:47.354866 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/1/link_state |
| e001444 | 19:58:47.354866 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/4/nodes/2/flags |
| e001445 | 19:58:47.354866 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/2/primary_id |
| e001446 | 19:58:47.354866 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/4/nodes/2/slots |
| e001447 | 19:58:47.354866 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/2/link_state |
| e001448 | 19:58:47.354866 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/4/nodes/3/flags |
| e001449 | 19:58:47.354866 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/3/primary_id |
| e001450 | 19:58:47.354866 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/4/nodes/3/slots |
| e001451 | 19:58:47.354866 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/3/link_state |
| e001452 | 19:58:47.354866 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/5/snapshot/views/4/nodes/4/flags |
| e001453 | 19:58:47.354866 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/5/snapshot/views/4/nodes/4/primary_id |
| e001454 | 19:58:47.354866 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/4/slots |
| e001455 | 19:58:47.354866 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/4/link_state |
| e001456 | 19:58:47.354866 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/4/nodes/5/flags |
| e001457 | 19:58:47.354866 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/5/snapshot/views/4/nodes/5/primary_id |
| e001458 | 19:58:47.354866 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/4/nodes/5/slots |
| e001459 | 19:58:47.354866 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/4/nodes/5/link_state |
| e001460 | 19:58:47.354866 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/4/persistence/aof_enabled |
| e001461 | 19:58:47.354866 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/4/persistence/loading |
| e001462 | 19:58:47.354866 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/4/persistence/aof_last_write_status |
| e001463 | 19:58:47.354866 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001464 | 19:58:47.354866 | node-5 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/4/process_live |
| e001465 | 19:58:47.355513 | node-4 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/3/reachable |
| e001466 | 19:58:47.355513 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/5/snapshot/views/3/replication/role |
| e001467 | 19:58:47.355513 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/5/snapshot/views/3/replication/connected_slaves |
| e001468 | 19:58:47.355513 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/5/snapshot/views/3/replication/master_link_status |
| e001469 | 19:58:47.355513 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/5/snapshot/views/3/replication/master_sync_in_progress |
| e001470 | 19:58:47.355513 | node-4 | info_replication | replication.master_repl_offset | valid | `0` | /samples/5/snapshot/views/3/replication/master_repl_offset |
| e001471 | 19:58:47.355513 | node-4 | info_replication | replication.slave_repl_offset | valid | `0` | /samples/5/snapshot/views/3/replication/slave_repl_offset |
| e001472 | 19:58:47.355513 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/5/snapshot/views/3/redis_id |
| e001473 | 19:58:47.355513 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/3/coverage |
| e001474 | 19:58:47.355513 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/3/nodes/0/flags |
| e001475 | 19:58:47.355513 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/0/primary_id |
| e001476 | 19:58:47.355513 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/3/nodes/0/slots |
| e001477 | 19:58:47.355513 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/0/link_state |
| e001478 | 19:58:47.355513 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/5/snapshot/views/3/nodes/1/flags |
| e001479 | 19:58:47.355513 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/5/snapshot/views/3/nodes/1/primary_id |
| e001480 | 19:58:47.355513 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/1/slots |
| e001481 | 19:58:47.355513 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/1/link_state |
| e001482 | 19:58:47.355513 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/3/nodes/2/flags |
| e001483 | 19:58:47.355513 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/5/snapshot/views/3/nodes/2/primary_id |
| e001484 | 19:58:47.355513 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/2/slots |
| e001485 | 19:58:47.355513 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/2/link_state |
| e001486 | 19:58:47.355513 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/3/nodes/3/flags |
| e001487 | 19:58:47.355513 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/3/primary_id |
| e001488 | 19:58:47.355513 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/3/nodes/3/slots |
| e001489 | 19:58:47.355513 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/3/link_state |
| e001490 | 19:58:47.355513 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/3/nodes/4/flags |
| e001491 | 19:58:47.355513 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/5/snapshot/views/3/nodes/4/primary_id |
| e001492 | 19:58:47.355513 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/4/slots |
| e001493 | 19:58:47.355513 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/4/link_state |
| e001494 | 19:58:47.355513 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/3/nodes/5/flags |
| e001495 | 19:58:47.355513 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/3/nodes/5/primary_id |
| e001496 | 19:58:47.355513 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/3/nodes/5/slots |
| e001497 | 19:58:47.355513 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/3/nodes/5/link_state |
| e001498 | 19:58:47.355513 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/3/persistence/aof_enabled |
| e001499 | 19:58:47.355513 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/3/persistence/loading |
| e001500 | 19:58:47.355513 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/3/persistence/aof_last_write_status |
| e001501 | 19:58:47.355513 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001502 | 19:58:47.355513 | node-4 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/3/process_live |
| e001503 | 19:58:47.35578 | node-1 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/0/reachable |
| e001504 | 19:58:47.35578 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/5/snapshot/views/0/replication/role |
| e001505 | 19:58:47.35578 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/5/snapshot/views/0/replication/connected_slaves |
| e001506 | 19:58:47.35578 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/5/snapshot/views/0/replication/master_link_status |
| e001507 | 19:58:47.35578 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/5/snapshot/views/0/replication/master_sync_in_progress |
| e001508 | 19:58:47.35578 | node-1 | info_replication | replication.master_repl_offset | valid | `14` | /samples/5/snapshot/views/0/replication/master_repl_offset |
| e001509 | 19:58:47.35578 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/5/snapshot/views/0/replication/slave_repl_offset |
| e001510 | 19:58:47.35578 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/5/snapshot/views/0/redis_id |
| e001511 | 19:58:47.35578 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/0/coverage |
| e001512 | 19:58:47.35578 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/0/nodes/0/flags |
| e001513 | 19:58:47.35578 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/5/snapshot/views/0/nodes/0/primary_id |
| e001514 | 19:58:47.35578 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/0/slots |
| e001515 | 19:58:47.35578 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/0/nodes/0/link_state |
| e001516 | 19:58:47.35578 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/0/nodes/1/flags |
| e001517 | 19:58:47.35578 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/1/primary_id |
| e001518 | 19:58:47.35578 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/0/nodes/1/slots |
| e001519 | 19:58:47.35578 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/0/nodes/1/link_state |
| e001520 | 19:58:47.35578 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/0/nodes/2/flags |
| e001521 | 19:58:47.35578 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/2/primary_id |
| e001522 | 19:58:47.35578 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/0/nodes/2/slots |
| e001523 | 19:58:47.35578 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/0/nodes/2/link_state |
| e001524 | 19:58:47.35578 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/5/snapshot/views/0/nodes/3/flags |
| e001525 | 19:58:47.35578 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/3/primary_id |
| e001526 | 19:58:47.35578 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/0/nodes/3/slots |
| e001527 | 19:58:47.35578 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/0/nodes/3/link_state |
| e001528 | 19:58:47.35578 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/0/nodes/4/flags |
| e001529 | 19:58:47.35578 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/5/snapshot/views/0/nodes/4/primary_id |
| e001530 | 19:58:47.35578 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/4/slots |
| e001531 | 19:58:47.35578 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/0/nodes/4/link_state |
| e001532 | 19:58:47.35578 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/0/nodes/5/flags |
| e001533 | 19:58:47.35578 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/5/snapshot/views/0/nodes/5/primary_id |
| e001534 | 19:58:47.35578 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/0/nodes/5/slots |
| e001535 | 19:58:47.35578 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/0/nodes/5/link_state |
| e001536 | 19:58:47.35578 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/0/persistence/aof_enabled |
| e001537 | 19:58:47.35578 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/0/persistence/loading |
| e001538 | 19:58:47.35578 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/0/persistence/aof_last_write_status |
| e001539 | 19:58:47.35578 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001540 | 19:58:47.35578 | node-1 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/0/process_live |
| e001541 | 19:58:47.355964 | node-2 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/1/reachable |
| e001542 | 19:58:47.355964 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/5/snapshot/views/1/replication/role |
| e001543 | 19:58:47.355964 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/5/snapshot/views/1/replication/connected_slaves |
| e001544 | 19:58:47.355964 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/5/snapshot/views/1/replication/master_link_status |
| e001545 | 19:58:47.355964 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/5/snapshot/views/1/replication/master_sync_in_progress |
| e001546 | 19:58:47.355964 | node-2 | info_replication | replication.master_repl_offset | valid | `14` | /samples/5/snapshot/views/1/replication/master_repl_offset |
| e001547 | 19:58:47.355964 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/5/snapshot/views/1/replication/slave_repl_offset |
| e001548 | 19:58:47.355964 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/5/snapshot/views/1/redis_id |
| e001549 | 19:58:47.355964 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/1/coverage |
| e001550 | 19:58:47.355964 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/1/nodes/0/flags |
| e001551 | 19:58:47.355964 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/5/snapshot/views/1/nodes/0/primary_id |
| e001552 | 19:58:47.355964 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/0/slots |
| e001553 | 19:58:47.355964 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/0/link_state |
| e001554 | 19:58:47.355964 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/1/nodes/1/flags |
| e001555 | 19:58:47.355964 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/1/primary_id |
| e001556 | 19:58:47.355964 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/1/nodes/1/slots |
| e001557 | 19:58:47.355964 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/1/link_state |
| e001558 | 19:58:47.355964 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/1/nodes/2/flags |
| e001559 | 19:58:47.355964 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/5/snapshot/views/1/nodes/2/primary_id |
| e001560 | 19:58:47.355964 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/2/slots |
| e001561 | 19:58:47.355964 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/2/link_state |
| e001562 | 19:58:47.355964 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/1/nodes/3/flags |
| e001563 | 19:58:47.355964 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/3/primary_id |
| e001564 | 19:58:47.355964 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/1/nodes/3/slots |
| e001565 | 19:58:47.355964 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/3/link_state |
| e001566 | 19:58:47.355964 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/1/nodes/4/flags |
| e001567 | 19:58:47.355964 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/5/snapshot/views/1/nodes/4/primary_id |
| e001568 | 19:58:47.355964 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/4/slots |
| e001569 | 19:58:47.355964 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/4/link_state |
| e001570 | 19:58:47.355964 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/5/snapshot/views/1/nodes/5/flags |
| e001571 | 19:58:47.355964 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/1/nodes/5/primary_id |
| e001572 | 19:58:47.355964 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/1/nodes/5/slots |
| e001573 | 19:58:47.355964 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/1/nodes/5/link_state |
| e001574 | 19:58:47.355964 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/1/persistence/aof_enabled |
| e001575 | 19:58:47.355964 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/1/persistence/loading |
| e001576 | 19:58:47.355964 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/1/persistence/aof_last_write_status |
| e001577 | 19:58:47.355964 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001578 | 19:58:47.355964 | node-2 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/1/process_live |
| e001579 | 19:58:47.356109 | node-3 | node_client | reachable | valid | `true` | /samples/5/snapshot/views/2/reachable |
| e001580 | 19:58:47.356109 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/5/snapshot/views/2/replication/role |
| e001581 | 19:58:47.356109 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/5/snapshot/views/2/replication/connected_slaves |
| e001582 | 19:58:47.356109 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/5/snapshot/views/2/replication/master_link_status |
| e001583 | 19:58:47.356109 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/5/snapshot/views/2/replication/master_sync_in_progress |
| e001584 | 19:58:47.356109 | node-3 | info_replication | replication.master_repl_offset | valid | `0` | /samples/5/snapshot/views/2/replication/master_repl_offset |
| e001585 | 19:58:47.356109 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/5/snapshot/views/2/replication/slave_repl_offset |
| e001586 | 19:58:47.356109 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/5/snapshot/views/2/redis_id |
| e001587 | 19:58:47.356109 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/5/snapshot/views/2/coverage |
| e001588 | 19:58:47.356109 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/2/nodes/0/flags |
| e001589 | 19:58:47.356109 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/0/primary_id |
| e001590 | 19:58:47.356109 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/5/snapshot/views/2/nodes/0/slots |
| e001591 | 19:58:47.356109 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/0/link_state |
| e001592 | 19:58:47.356109 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/5/snapshot/views/2/nodes/1/flags |
| e001593 | 19:58:47.356109 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/1/primary_id |
| e001594 | 19:58:47.356109 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/5/snapshot/views/2/nodes/1/slots |
| e001595 | 19:58:47.356109 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/1/link_state |
| e001596 | 19:58:47.356109 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/2/nodes/2/flags |
| e001597 | 19:58:47.356109 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/5/snapshot/views/2/nodes/2/primary_id |
| e001598 | 19:58:47.356109 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/2/slots |
| e001599 | 19:58:47.356109 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/2/link_state |
| e001600 | 19:58:47.356109 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/2/nodes/3/flags |
| e001601 | 19:58:47.356109 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/5/snapshot/views/2/nodes/3/primary_id |
| e001602 | 19:58:47.356109 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/3/slots |
| e001603 | 19:58:47.356109 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/3/link_state |
| e001604 | 19:58:47.356109 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/5/snapshot/views/2/nodes/4/flags |
| e001605 | 19:58:47.356109 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/5/snapshot/views/2/nodes/4/primary_id |
| e001606 | 19:58:47.356109 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/4/slots |
| e001607 | 19:58:47.356109 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/4/link_state |
| e001608 | 19:58:47.356109 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/5/snapshot/views/2/nodes/5/flags |
| e001609 | 19:58:47.356109 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/5/snapshot/views/2/nodes/5/primary_id |
| e001610 | 19:58:47.356109 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/5/snapshot/views/2/nodes/5/slots |
| e001611 | 19:58:47.356109 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/5/snapshot/views/2/nodes/5/link_state |
| e001612 | 19:58:47.356109 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/5/snapshot/views/2/persistence/aof_enabled |
| e001613 | 19:58:47.356109 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/5/snapshot/views/2/persistence/loading |
| e001614 | 19:58:47.356109 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/5/snapshot/views/2/persistence/aof_last_write_status |
| e001615 | 19:58:47.356109 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/5/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001616 | 19:58:47.356109 | node-3 | admin | process_live | missing_from_schema | `` | /samples/5/snapshot/views/2/process_live |
| e001617 | 19:58:48.216611 | node-6 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/5/reachable |
| e001618 | 19:58:48.216611 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/6/snapshot/views/5/replication/role |
| e001619 | 19:58:48.216611 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/6/snapshot/views/5/replication/connected_slaves |
| e001620 | 19:58:48.216611 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/6/snapshot/views/5/replication/master_link_status |
| e001621 | 19:58:48.216611 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/6/snapshot/views/5/replication/master_sync_in_progress |
| e001622 | 19:58:48.216611 | node-6 | info_replication | replication.master_repl_offset | valid | `14` | /samples/6/snapshot/views/5/replication/master_repl_offset |
| e001623 | 19:58:48.216611 | node-6 | info_replication | replication.slave_repl_offset | valid | `14` | /samples/6/snapshot/views/5/replication/slave_repl_offset |
| e001624 | 19:58:48.216611 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/6/snapshot/views/5/redis_id |
| e001625 | 19:58:48.216611 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/5/coverage |
| e001626 | 19:58:48.216611 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/5/nodes/0/flags |
| e001627 | 19:58:48.216611 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/6/snapshot/views/5/nodes/0/primary_id |
| e001628 | 19:58:48.216611 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/0/slots |
| e001629 | 19:58:48.216611 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/0/link_state |
| e001630 | 19:58:48.216611 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/5/nodes/1/flags |
| e001631 | 19:58:48.216611 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/1/primary_id |
| e001632 | 19:58:48.216611 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/5/nodes/1/slots |
| e001633 | 19:58:48.216611 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/1/link_state |
| e001634 | 19:58:48.216611 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/5/nodes/2/flags |
| e001635 | 19:58:48.216611 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/2/primary_id |
| e001636 | 19:58:48.216611 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/5/nodes/2/slots |
| e001637 | 19:58:48.216611 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/2/link_state |
| e001638 | 19:58:48.216611 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/5/nodes/3/flags |
| e001639 | 19:58:48.216611 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/3/primary_id |
| e001640 | 19:58:48.216611 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/5/nodes/3/slots |
| e001641 | 19:58:48.216611 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/3/link_state |
| e001642 | 19:58:48.216611 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/5/nodes/4/flags |
| e001643 | 19:58:48.216611 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/6/snapshot/views/5/nodes/4/primary_id |
| e001644 | 19:58:48.216611 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/4/slots |
| e001645 | 19:58:48.216611 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/4/link_state |
| e001646 | 19:58:48.216611 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/6/snapshot/views/5/nodes/5/flags |
| e001647 | 19:58:48.216611 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/6/snapshot/views/5/nodes/5/primary_id |
| e001648 | 19:58:48.216611 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/5/nodes/5/slots |
| e001649 | 19:58:48.216611 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/5/nodes/5/link_state |
| e001650 | 19:58:48.216611 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/5/persistence/aof_enabled |
| e001651 | 19:58:48.216611 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/5/persistence/loading |
| e001652 | 19:58:48.216611 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/5/persistence/aof_last_write_status |
| e001653 | 19:58:48.216611 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e001654 | 19:58:48.216611 | node-6 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/5/process_live |
| e001655 | 19:58:48.216613 | node-3 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/2/reachable |
| e001656 | 19:58:48.216613 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/6/snapshot/views/2/replication/role |
| e001657 | 19:58:48.216613 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/6/snapshot/views/2/replication/connected_slaves |
| e001658 | 19:58:48.216613 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/6/snapshot/views/2/replication/master_link_status |
| e001659 | 19:58:48.216613 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/6/snapshot/views/2/replication/master_sync_in_progress |
| e001660 | 19:58:48.216613 | node-3 | info_replication | replication.master_repl_offset | valid | `14` | /samples/6/snapshot/views/2/replication/master_repl_offset |
| e001661 | 19:58:48.216613 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/6/snapshot/views/2/replication/slave_repl_offset |
| e001662 | 19:58:48.216613 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/6/snapshot/views/2/redis_id |
| e001663 | 19:58:48.216613 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/2/coverage |
| e001664 | 19:58:48.216613 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/2/nodes/0/flags |
| e001665 | 19:58:48.216613 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/0/primary_id |
| e001666 | 19:58:48.216613 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/2/nodes/0/slots |
| e001667 | 19:58:48.216613 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/0/link_state |
| e001668 | 19:58:48.216613 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/6/snapshot/views/2/nodes/1/flags |
| e001669 | 19:58:48.216613 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/1/primary_id |
| e001670 | 19:58:48.216613 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/2/nodes/1/slots |
| e001671 | 19:58:48.216613 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/1/link_state |
| e001672 | 19:58:48.216613 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/2/nodes/2/flags |
| e001673 | 19:58:48.216613 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/6/snapshot/views/2/nodes/2/primary_id |
| e001674 | 19:58:48.216613 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/2/slots |
| e001675 | 19:58:48.216613 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/2/link_state |
| e001676 | 19:58:48.216613 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/2/nodes/3/flags |
| e001677 | 19:58:48.216613 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/6/snapshot/views/2/nodes/3/primary_id |
| e001678 | 19:58:48.216613 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/3/slots |
| e001679 | 19:58:48.216613 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/3/link_state |
| e001680 | 19:58:48.216613 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/2/nodes/4/flags |
| e001681 | 19:58:48.216613 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/6/snapshot/views/2/nodes/4/primary_id |
| e001682 | 19:58:48.216613 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/4/slots |
| e001683 | 19:58:48.216613 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/4/link_state |
| e001684 | 19:58:48.216613 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/2/nodes/5/flags |
| e001685 | 19:58:48.216613 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/2/nodes/5/primary_id |
| e001686 | 19:58:48.216613 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/2/nodes/5/slots |
| e001687 | 19:58:48.216613 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/2/nodes/5/link_state |
| e001688 | 19:58:48.216613 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/2/persistence/aof_enabled |
| e001689 | 19:58:48.216613 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/2/persistence/loading |
| e001690 | 19:58:48.216613 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/2/persistence/aof_last_write_status |
| e001691 | 19:58:48.216613 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001692 | 19:58:48.216613 | node-3 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/2/process_live |
| e001693 | 19:58:48.216618 | node-4 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/3/reachable |
| e001694 | 19:58:48.216618 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/6/snapshot/views/3/replication/role |
| e001695 | 19:58:48.216618 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/6/snapshot/views/3/replication/connected_slaves |
| e001696 | 19:58:48.216618 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/6/snapshot/views/3/replication/master_link_status |
| e001697 | 19:58:48.216618 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/6/snapshot/views/3/replication/master_sync_in_progress |
| e001698 | 19:58:48.216618 | node-4 | info_replication | replication.master_repl_offset | valid | `14` | /samples/6/snapshot/views/3/replication/master_repl_offset |
| e001699 | 19:58:48.216618 | node-4 | info_replication | replication.slave_repl_offset | valid | `14` | /samples/6/snapshot/views/3/replication/slave_repl_offset |
| e001700 | 19:58:48.216618 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/6/snapshot/views/3/redis_id |
| e001701 | 19:58:48.216618 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/3/coverage |
| e001702 | 19:58:48.216618 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/3/nodes/0/flags |
| e001703 | 19:58:48.216618 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/0/primary_id |
| e001704 | 19:58:48.216618 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/3/nodes/0/slots |
| e001705 | 19:58:48.216618 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/0/link_state |
| e001706 | 19:58:48.216618 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/6/snapshot/views/3/nodes/1/flags |
| e001707 | 19:58:48.216618 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/6/snapshot/views/3/nodes/1/primary_id |
| e001708 | 19:58:48.216618 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/1/slots |
| e001709 | 19:58:48.216618 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/1/link_state |
| e001710 | 19:58:48.216618 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/3/nodes/2/flags |
| e001711 | 19:58:48.216618 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/6/snapshot/views/3/nodes/2/primary_id |
| e001712 | 19:58:48.216618 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/2/slots |
| e001713 | 19:58:48.216618 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/2/link_state |
| e001714 | 19:58:48.216618 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/3/nodes/3/flags |
| e001715 | 19:58:48.216618 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/3/primary_id |
| e001716 | 19:58:48.216618 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/3/nodes/3/slots |
| e001717 | 19:58:48.216618 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/3/link_state |
| e001718 | 19:58:48.216618 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/3/nodes/4/flags |
| e001719 | 19:58:48.216618 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/6/snapshot/views/3/nodes/4/primary_id |
| e001720 | 19:58:48.216618 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/4/slots |
| e001721 | 19:58:48.216618 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/4/link_state |
| e001722 | 19:58:48.216618 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/3/nodes/5/flags |
| e001723 | 19:58:48.216618 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/3/nodes/5/primary_id |
| e001724 | 19:58:48.216618 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/3/nodes/5/slots |
| e001725 | 19:58:48.216618 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/3/nodes/5/link_state |
| e001726 | 19:58:48.216618 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/3/persistence/aof_enabled |
| e001727 | 19:58:48.216618 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/3/persistence/loading |
| e001728 | 19:58:48.216618 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/3/persistence/aof_last_write_status |
| e001729 | 19:58:48.216618 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001730 | 19:58:48.216618 | node-4 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/3/process_live |
| e001731 | 19:58:48.216649 | node-2 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/1/reachable |
| e001732 | 19:58:48.216649 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/6/snapshot/views/1/replication/role |
| e001733 | 19:58:48.216649 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/6/snapshot/views/1/replication/connected_slaves |
| e001734 | 19:58:48.216649 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/6/snapshot/views/1/replication/master_link_status |
| e001735 | 19:58:48.216649 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/6/snapshot/views/1/replication/master_sync_in_progress |
| e001736 | 19:58:48.216649 | node-2 | info_replication | replication.master_repl_offset | valid | `14` | /samples/6/snapshot/views/1/replication/master_repl_offset |
| e001737 | 19:58:48.216649 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/6/snapshot/views/1/replication/slave_repl_offset |
| e001738 | 19:58:48.216649 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/6/snapshot/views/1/redis_id |
| e001739 | 19:58:48.216649 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/1/coverage |
| e001740 | 19:58:48.216649 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/1/nodes/0/flags |
| e001741 | 19:58:48.216649 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/6/snapshot/views/1/nodes/0/primary_id |
| e001742 | 19:58:48.216649 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/0/slots |
| e001743 | 19:58:48.216649 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/0/link_state |
| e001744 | 19:58:48.216649 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/1/nodes/1/flags |
| e001745 | 19:58:48.216649 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/1/primary_id |
| e001746 | 19:58:48.216649 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/1/nodes/1/slots |
| e001747 | 19:58:48.216649 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/1/link_state |
| e001748 | 19:58:48.216649 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/1/nodes/2/flags |
| e001749 | 19:58:48.216649 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/6/snapshot/views/1/nodes/2/primary_id |
| e001750 | 19:58:48.216649 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/2/slots |
| e001751 | 19:58:48.216649 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/2/link_state |
| e001752 | 19:58:48.216649 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/1/nodes/3/flags |
| e001753 | 19:58:48.216649 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/3/primary_id |
| e001754 | 19:58:48.216649 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/1/nodes/3/slots |
| e001755 | 19:58:48.216649 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/3/link_state |
| e001756 | 19:58:48.216649 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/1/nodes/4/flags |
| e001757 | 19:58:48.216649 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/6/snapshot/views/1/nodes/4/primary_id |
| e001758 | 19:58:48.216649 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/4/slots |
| e001759 | 19:58:48.216649 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/4/link_state |
| e001760 | 19:58:48.216649 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/6/snapshot/views/1/nodes/5/flags |
| e001761 | 19:58:48.216649 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/1/nodes/5/primary_id |
| e001762 | 19:58:48.216649 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/1/nodes/5/slots |
| e001763 | 19:58:48.216649 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/1/nodes/5/link_state |
| e001764 | 19:58:48.216649 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/1/persistence/aof_enabled |
| e001765 | 19:58:48.216649 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/1/persistence/loading |
| e001766 | 19:58:48.216649 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/1/persistence/aof_last_write_status |
| e001767 | 19:58:48.216649 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e001768 | 19:58:48.216649 | node-2 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/1/process_live |
| e001769 | 19:58:48.216653 | node-5 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/4/reachable |
| e001770 | 19:58:48.216653 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/6/snapshot/views/4/replication/role |
| e001771 | 19:58:48.216653 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/6/snapshot/views/4/replication/connected_slaves |
| e001772 | 19:58:48.216653 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/6/snapshot/views/4/replication/master_link_status |
| e001773 | 19:58:48.216653 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/6/snapshot/views/4/replication/master_sync_in_progress |
| e001774 | 19:58:48.216653 | node-5 | info_replication | replication.master_repl_offset | valid | `14` | /samples/6/snapshot/views/4/replication/master_repl_offset |
| e001775 | 19:58:48.216653 | node-5 | info_replication | replication.slave_repl_offset | valid | `14` | /samples/6/snapshot/views/4/replication/slave_repl_offset |
| e001776 | 19:58:48.216653 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/6/snapshot/views/4/redis_id |
| e001777 | 19:58:48.216653 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/4/coverage |
| e001778 | 19:58:48.216653 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/4/nodes/0/flags |
| e001779 | 19:58:48.216653 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/0/primary_id |
| e001780 | 19:58:48.216653 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/4/nodes/0/slots |
| e001781 | 19:58:48.216653 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/0/link_state |
| e001782 | 19:58:48.216653 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/4/nodes/1/flags |
| e001783 | 19:58:48.216653 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/6/snapshot/views/4/nodes/1/primary_id |
| e001784 | 19:58:48.216653 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/1/slots |
| e001785 | 19:58:48.216653 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/1/link_state |
| e001786 | 19:58:48.216653 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/4/nodes/2/flags |
| e001787 | 19:58:48.216653 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/2/primary_id |
| e001788 | 19:58:48.216653 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/4/nodes/2/slots |
| e001789 | 19:58:48.216653 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/2/link_state |
| e001790 | 19:58:48.216653 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/4/nodes/3/flags |
| e001791 | 19:58:48.216653 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/3/primary_id |
| e001792 | 19:58:48.216653 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/4/nodes/3/slots |
| e001793 | 19:58:48.216653 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/3/link_state |
| e001794 | 19:58:48.216653 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/6/snapshot/views/4/nodes/4/flags |
| e001795 | 19:58:48.216653 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/6/snapshot/views/4/nodes/4/primary_id |
| e001796 | 19:58:48.216653 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/4/slots |
| e001797 | 19:58:48.216653 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/4/link_state |
| e001798 | 19:58:48.216653 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/4/nodes/5/flags |
| e001799 | 19:58:48.216653 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/6/snapshot/views/4/nodes/5/primary_id |
| e001800 | 19:58:48.216653 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/4/nodes/5/slots |
| e001801 | 19:58:48.216653 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/4/nodes/5/link_state |
| e001802 | 19:58:48.216653 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/4/persistence/aof_enabled |
| e001803 | 19:58:48.216653 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/4/persistence/loading |
| e001804 | 19:58:48.216653 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/4/persistence/aof_last_write_status |
| e001805 | 19:58:48.216653 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001806 | 19:58:48.216653 | node-5 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/4/process_live |
| e001807 | 19:58:48.216662 | node-1 | node_client | reachable | valid | `true` | /samples/6/snapshot/views/0/reachable |
| e001808 | 19:58:48.216662 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/6/snapshot/views/0/replication/role |
| e001809 | 19:58:48.216662 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/6/snapshot/views/0/replication/connected_slaves |
| e001810 | 19:58:48.216662 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/6/snapshot/views/0/replication/master_link_status |
| e001811 | 19:58:48.216662 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/6/snapshot/views/0/replication/master_sync_in_progress |
| e001812 | 19:58:48.216662 | node-1 | info_replication | replication.master_repl_offset | valid | `14` | /samples/6/snapshot/views/0/replication/master_repl_offset |
| e001813 | 19:58:48.216662 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/6/snapshot/views/0/replication/slave_repl_offset |
| e001814 | 19:58:48.216662 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/6/snapshot/views/0/redis_id |
| e001815 | 19:58:48.216662 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/6/snapshot/views/0/coverage |
| e001816 | 19:58:48.216662 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/0/nodes/0/flags |
| e001817 | 19:58:48.216662 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/6/snapshot/views/0/nodes/0/primary_id |
| e001818 | 19:58:48.216662 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/0/slots |
| e001819 | 19:58:48.216662 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/0/link_state |
| e001820 | 19:58:48.216662 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/0/nodes/1/flags |
| e001821 | 19:58:48.216662 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/1/primary_id |
| e001822 | 19:58:48.216662 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/6/snapshot/views/0/nodes/1/slots |
| e001823 | 19:58:48.216662 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/1/link_state |
| e001824 | 19:58:48.216662 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/6/snapshot/views/0/nodes/2/flags |
| e001825 | 19:58:48.216662 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/2/primary_id |
| e001826 | 19:58:48.216662 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/6/snapshot/views/0/nodes/2/slots |
| e001827 | 19:58:48.216662 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/2/link_state |
| e001828 | 19:58:48.216662 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/6/snapshot/views/0/nodes/3/flags |
| e001829 | 19:58:48.216662 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/3/primary_id |
| e001830 | 19:58:48.216662 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/6/snapshot/views/0/nodes/3/slots |
| e001831 | 19:58:48.216662 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/3/link_state |
| e001832 | 19:58:48.216662 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/0/nodes/4/flags |
| e001833 | 19:58:48.216662 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/6/snapshot/views/0/nodes/4/primary_id |
| e001834 | 19:58:48.216662 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/4/slots |
| e001835 | 19:58:48.216662 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/4/link_state |
| e001836 | 19:58:48.216662 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/6/snapshot/views/0/nodes/5/flags |
| e001837 | 19:58:48.216662 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/6/snapshot/views/0/nodes/5/primary_id |
| e001838 | 19:58:48.216662 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/6/snapshot/views/0/nodes/5/slots |
| e001839 | 19:58:48.216662 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/6/snapshot/views/0/nodes/5/link_state |
| e001840 | 19:58:48.216662 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/6/snapshot/views/0/persistence/aof_enabled |
| e001841 | 19:58:48.216662 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/6/snapshot/views/0/persistence/loading |
| e001842 | 19:58:48.216662 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/6/snapshot/views/0/persistence/aof_last_write_status |
| e001843 | 19:58:48.216662 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/6/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001844 | 19:58:48.216662 | node-1 | admin | process_live | missing_from_schema | `` | /samples/6/snapshot/views/0/process_live |
| e001845 | 19:58:48.846929 | node-1 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/0/reachable |
| e001846 | 19:58:48.846929 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/7/snapshot/views/0/replication/role |
| e001847 | 19:58:48.846929 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/7/snapshot/views/0/replication/connected_slaves |
| e001848 | 19:58:48.846929 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/7/snapshot/views/0/replication/master_link_status |
| e001849 | 19:58:48.846929 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/7/snapshot/views/0/replication/master_sync_in_progress |
| e001850 | 19:58:48.846929 | node-1 | info_replication | replication.master_repl_offset | valid | `14` | /samples/7/snapshot/views/0/replication/master_repl_offset |
| e001851 | 19:58:48.846929 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/7/snapshot/views/0/replication/slave_repl_offset |
| e001852 | 19:58:48.846929 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/7/snapshot/views/0/redis_id |
| e001853 | 19:58:48.846929 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/0/coverage |
| e001854 | 19:58:48.846929 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/0/nodes/0/flags |
| e001855 | 19:58:48.846929 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/7/snapshot/views/0/nodes/0/primary_id |
| e001856 | 19:58:48.846929 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/0/slots |
| e001857 | 19:58:48.846929 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/0/link_state |
| e001858 | 19:58:48.846929 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/0/nodes/1/flags |
| e001859 | 19:58:48.846929 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/1/primary_id |
| e001860 | 19:58:48.846929 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/0/nodes/1/slots |
| e001861 | 19:58:48.846929 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/1/link_state |
| e001862 | 19:58:48.846929 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/0/nodes/2/flags |
| e001863 | 19:58:48.846929 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/2/primary_id |
| e001864 | 19:58:48.846929 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/0/nodes/2/slots |
| e001865 | 19:58:48.846929 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/2/link_state |
| e001866 | 19:58:48.846929 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/7/snapshot/views/0/nodes/3/flags |
| e001867 | 19:58:48.846929 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/3/primary_id |
| e001868 | 19:58:48.846929 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/0/nodes/3/slots |
| e001869 | 19:58:48.846929 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/3/link_state |
| e001870 | 19:58:48.846929 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/0/nodes/4/flags |
| e001871 | 19:58:48.846929 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/7/snapshot/views/0/nodes/4/primary_id |
| e001872 | 19:58:48.846929 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/4/slots |
| e001873 | 19:58:48.846929 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/4/link_state |
| e001874 | 19:58:48.846929 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/0/nodes/5/flags |
| e001875 | 19:58:48.846929 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/7/snapshot/views/0/nodes/5/primary_id |
| e001876 | 19:58:48.846929 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/0/nodes/5/slots |
| e001877 | 19:58:48.846929 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/0/nodes/5/link_state |
| e001878 | 19:58:48.846929 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/0/persistence/aof_enabled |
| e001879 | 19:58:48.846929 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/0/persistence/loading |
| e001880 | 19:58:48.846929 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/0/persistence/aof_last_write_status |
| e001881 | 19:58:48.846929 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e001882 | 19:58:48.846929 | node-1 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/0/process_live |
| e001883 | 19:58:48.846934 | node-5 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/4/reachable |
| e001884 | 19:58:48.846934 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/7/snapshot/views/4/replication/role |
| e001885 | 19:58:48.846934 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/7/snapshot/views/4/replication/connected_slaves |
| e001886 | 19:58:48.846934 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/7/snapshot/views/4/replication/master_link_status |
| e001887 | 19:58:48.846934 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/7/snapshot/views/4/replication/master_sync_in_progress |
| e001888 | 19:58:48.846934 | node-5 | info_replication | replication.master_repl_offset | valid | `14` | /samples/7/snapshot/views/4/replication/master_repl_offset |
| e001889 | 19:58:48.846934 | node-5 | info_replication | replication.slave_repl_offset | valid | `14` | /samples/7/snapshot/views/4/replication/slave_repl_offset |
| e001890 | 19:58:48.846934 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/7/snapshot/views/4/redis_id |
| e001891 | 19:58:48.846934 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/4/coverage |
| e001892 | 19:58:48.846934 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/4/nodes/0/flags |
| e001893 | 19:58:48.846934 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/0/primary_id |
| e001894 | 19:58:48.846934 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/4/nodes/0/slots |
| e001895 | 19:58:48.846934 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/0/link_state |
| e001896 | 19:58:48.846934 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/4/nodes/1/flags |
| e001897 | 19:58:48.846934 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/7/snapshot/views/4/nodes/1/primary_id |
| e001898 | 19:58:48.846934 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/1/slots |
| e001899 | 19:58:48.846934 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/1/link_state |
| e001900 | 19:58:48.846934 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/4/nodes/2/flags |
| e001901 | 19:58:48.846934 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/2/primary_id |
| e001902 | 19:58:48.846934 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/4/nodes/2/slots |
| e001903 | 19:58:48.846934 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/2/link_state |
| e001904 | 19:58:48.846934 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/4/nodes/3/flags |
| e001905 | 19:58:48.846934 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/3/primary_id |
| e001906 | 19:58:48.846934 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/4/nodes/3/slots |
| e001907 | 19:58:48.846934 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/3/link_state |
| e001908 | 19:58:48.846934 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/7/snapshot/views/4/nodes/4/flags |
| e001909 | 19:58:48.846934 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/7/snapshot/views/4/nodes/4/primary_id |
| e001910 | 19:58:48.846934 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/4/slots |
| e001911 | 19:58:48.846934 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/4/link_state |
| e001912 | 19:58:48.846934 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/4/nodes/5/flags |
| e001913 | 19:58:48.846934 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/7/snapshot/views/4/nodes/5/primary_id |
| e001914 | 19:58:48.846934 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/4/nodes/5/slots |
| e001915 | 19:58:48.846934 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/4/nodes/5/link_state |
| e001916 | 19:58:48.846934 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/4/persistence/aof_enabled |
| e001917 | 19:58:48.846934 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/4/persistence/loading |
| e001918 | 19:58:48.846934 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/4/persistence/aof_last_write_status |
| e001919 | 19:58:48.846934 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e001920 | 19:58:48.846934 | node-5 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/4/process_live |
| e001921 | 19:58:48.846958 | node-3 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/2/reachable |
| e001922 | 19:58:48.846958 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/7/snapshot/views/2/replication/role |
| e001923 | 19:58:48.846958 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/7/snapshot/views/2/replication/connected_slaves |
| e001924 | 19:58:48.846958 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/7/snapshot/views/2/replication/master_link_status |
| e001925 | 19:58:48.846958 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/7/snapshot/views/2/replication/master_sync_in_progress |
| e001926 | 19:58:48.846958 | node-3 | info_replication | replication.master_repl_offset | valid | `14` | /samples/7/snapshot/views/2/replication/master_repl_offset |
| e001927 | 19:58:48.846958 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/7/snapshot/views/2/replication/slave_repl_offset |
| e001928 | 19:58:48.846958 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/7/snapshot/views/2/redis_id |
| e001929 | 19:58:48.846958 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/2/coverage |
| e001930 | 19:58:48.846958 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/2/nodes/0/flags |
| e001931 | 19:58:48.846958 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/0/primary_id |
| e001932 | 19:58:48.846958 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/2/nodes/0/slots |
| e001933 | 19:58:48.846958 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/0/link_state |
| e001934 | 19:58:48.846958 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/7/snapshot/views/2/nodes/1/flags |
| e001935 | 19:58:48.846958 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/1/primary_id |
| e001936 | 19:58:48.846958 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/2/nodes/1/slots |
| e001937 | 19:58:48.846958 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/1/link_state |
| e001938 | 19:58:48.846958 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/2/nodes/2/flags |
| e001939 | 19:58:48.846958 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/7/snapshot/views/2/nodes/2/primary_id |
| e001940 | 19:58:48.846958 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/2/slots |
| e001941 | 19:58:48.846958 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/2/link_state |
| e001942 | 19:58:48.846958 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/2/nodes/3/flags |
| e001943 | 19:58:48.846958 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/7/snapshot/views/2/nodes/3/primary_id |
| e001944 | 19:58:48.846958 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/3/slots |
| e001945 | 19:58:48.846958 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/3/link_state |
| e001946 | 19:58:48.846958 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/2/nodes/4/flags |
| e001947 | 19:58:48.846958 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/7/snapshot/views/2/nodes/4/primary_id |
| e001948 | 19:58:48.846958 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/4/slots |
| e001949 | 19:58:48.846958 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/4/link_state |
| e001950 | 19:58:48.846958 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/2/nodes/5/flags |
| e001951 | 19:58:48.846958 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/2/nodes/5/primary_id |
| e001952 | 19:58:48.846958 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/2/nodes/5/slots |
| e001953 | 19:58:48.846958 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/2/nodes/5/link_state |
| e001954 | 19:58:48.846958 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/2/persistence/aof_enabled |
| e001955 | 19:58:48.846958 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/2/persistence/loading |
| e001956 | 19:58:48.846958 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/2/persistence/aof_last_write_status |
| e001957 | 19:58:48.846958 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e001958 | 19:58:48.846958 | node-3 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/2/process_live |
| e001959 | 19:58:48.84697 | node-4 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/3/reachable |
| e001960 | 19:58:48.84697 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/7/snapshot/views/3/replication/role |
| e001961 | 19:58:48.84697 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/7/snapshot/views/3/replication/connected_slaves |
| e001962 | 19:58:48.84697 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/7/snapshot/views/3/replication/master_link_status |
| e001963 | 19:58:48.84697 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/7/snapshot/views/3/replication/master_sync_in_progress |
| e001964 | 19:58:48.84697 | node-4 | info_replication | replication.master_repl_offset | valid | `14` | /samples/7/snapshot/views/3/replication/master_repl_offset |
| e001965 | 19:58:48.84697 | node-4 | info_replication | replication.slave_repl_offset | valid | `14` | /samples/7/snapshot/views/3/replication/slave_repl_offset |
| e001966 | 19:58:48.84697 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/7/snapshot/views/3/redis_id |
| e001967 | 19:58:48.84697 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/3/coverage |
| e001968 | 19:58:48.84697 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/3/nodes/0/flags |
| e001969 | 19:58:48.84697 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/0/primary_id |
| e001970 | 19:58:48.84697 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/3/nodes/0/slots |
| e001971 | 19:58:48.84697 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/0/link_state |
| e001972 | 19:58:48.84697 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/7/snapshot/views/3/nodes/1/flags |
| e001973 | 19:58:48.84697 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/7/snapshot/views/3/nodes/1/primary_id |
| e001974 | 19:58:48.84697 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/1/slots |
| e001975 | 19:58:48.84697 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/1/link_state |
| e001976 | 19:58:48.84697 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/3/nodes/2/flags |
| e001977 | 19:58:48.84697 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/7/snapshot/views/3/nodes/2/primary_id |
| e001978 | 19:58:48.84697 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/2/slots |
| e001979 | 19:58:48.84697 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/2/link_state |
| e001980 | 19:58:48.84697 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/3/nodes/3/flags |
| e001981 | 19:58:48.84697 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/3/primary_id |
| e001982 | 19:58:48.84697 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/3/nodes/3/slots |
| e001983 | 19:58:48.84697 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/3/link_state |
| e001984 | 19:58:48.84697 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/3/nodes/4/flags |
| e001985 | 19:58:48.84697 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/7/snapshot/views/3/nodes/4/primary_id |
| e001986 | 19:58:48.84697 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/4/slots |
| e001987 | 19:58:48.84697 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/4/link_state |
| e001988 | 19:58:48.84697 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/3/nodes/5/flags |
| e001989 | 19:58:48.84697 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/3/nodes/5/primary_id |
| e001990 | 19:58:48.84697 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/3/nodes/5/slots |
| e001991 | 19:58:48.84697 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/3/nodes/5/link_state |
| e001992 | 19:58:48.84697 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/3/persistence/aof_enabled |
| e001993 | 19:58:48.84697 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/3/persistence/loading |
| e001994 | 19:58:48.84697 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/3/persistence/aof_last_write_status |
| e001995 | 19:58:48.84697 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e001996 | 19:58:48.84697 | node-4 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/3/process_live |
| e001997 | 19:58:48.84698 | node-6 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/5/reachable |
| e001998 | 19:58:48.84698 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/7/snapshot/views/5/replication/role |
| e001999 | 19:58:48.84698 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/7/snapshot/views/5/replication/connected_slaves |
| e002000 | 19:58:48.84698 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/7/snapshot/views/5/replication/master_link_status |
| e002001 | 19:58:48.84698 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/7/snapshot/views/5/replication/master_sync_in_progress |
| e002002 | 19:58:48.84698 | node-6 | info_replication | replication.master_repl_offset | valid | `14` | /samples/7/snapshot/views/5/replication/master_repl_offset |
| e002003 | 19:58:48.84698 | node-6 | info_replication | replication.slave_repl_offset | valid | `14` | /samples/7/snapshot/views/5/replication/slave_repl_offset |
| e002004 | 19:58:48.84698 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/7/snapshot/views/5/redis_id |
| e002005 | 19:58:48.84698 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/5/coverage |
| e002006 | 19:58:48.84698 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/5/nodes/0/flags |
| e002007 | 19:58:48.84698 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/7/snapshot/views/5/nodes/0/primary_id |
| e002008 | 19:58:48.84698 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/0/slots |
| e002009 | 19:58:48.84698 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/0/link_state |
| e002010 | 19:58:48.84698 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/5/nodes/1/flags |
| e002011 | 19:58:48.84698 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/1/primary_id |
| e002012 | 19:58:48.84698 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/5/nodes/1/slots |
| e002013 | 19:58:48.84698 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/1/link_state |
| e002014 | 19:58:48.84698 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/5/nodes/2/flags |
| e002015 | 19:58:48.84698 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/2/primary_id |
| e002016 | 19:58:48.84698 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/5/nodes/2/slots |
| e002017 | 19:58:48.84698 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/2/link_state |
| e002018 | 19:58:48.84698 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/5/nodes/3/flags |
| e002019 | 19:58:48.84698 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/3/primary_id |
| e002020 | 19:58:48.84698 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/5/nodes/3/slots |
| e002021 | 19:58:48.84698 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/3/link_state |
| e002022 | 19:58:48.84698 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/5/nodes/4/flags |
| e002023 | 19:58:48.84698 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/7/snapshot/views/5/nodes/4/primary_id |
| e002024 | 19:58:48.84698 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/4/slots |
| e002025 | 19:58:48.84698 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/4/link_state |
| e002026 | 19:58:48.84698 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/7/snapshot/views/5/nodes/5/flags |
| e002027 | 19:58:48.84698 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/7/snapshot/views/5/nodes/5/primary_id |
| e002028 | 19:58:48.84698 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/5/nodes/5/slots |
| e002029 | 19:58:48.84698 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/5/nodes/5/link_state |
| e002030 | 19:58:48.84698 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/5/persistence/aof_enabled |
| e002031 | 19:58:48.84698 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/5/persistence/loading |
| e002032 | 19:58:48.84698 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/5/persistence/aof_last_write_status |
| e002033 | 19:58:48.84698 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002034 | 19:58:48.84698 | node-6 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/5/process_live |
| e002035 | 19:58:48.847002 | node-2 | node_client | reachable | valid | `true` | /samples/7/snapshot/views/1/reachable |
| e002036 | 19:58:48.847002 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/7/snapshot/views/1/replication/role |
| e002037 | 19:58:48.847002 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/7/snapshot/views/1/replication/connected_slaves |
| e002038 | 19:58:48.847002 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/7/snapshot/views/1/replication/master_link_status |
| e002039 | 19:58:48.847002 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/7/snapshot/views/1/replication/master_sync_in_progress |
| e002040 | 19:58:48.847002 | node-2 | info_replication | replication.master_repl_offset | valid | `14` | /samples/7/snapshot/views/1/replication/master_repl_offset |
| e002041 | 19:58:48.847002 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/7/snapshot/views/1/replication/slave_repl_offset |
| e002042 | 19:58:48.847002 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/7/snapshot/views/1/redis_id |
| e002043 | 19:58:48.847002 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/7/snapshot/views/1/coverage |
| e002044 | 19:58:48.847002 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/1/nodes/0/flags |
| e002045 | 19:58:48.847002 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/7/snapshot/views/1/nodes/0/primary_id |
| e002046 | 19:58:48.847002 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/0/slots |
| e002047 | 19:58:48.847002 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/0/link_state |
| e002048 | 19:58:48.847002 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/1/nodes/1/flags |
| e002049 | 19:58:48.847002 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/1/primary_id |
| e002050 | 19:58:48.847002 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/7/snapshot/views/1/nodes/1/slots |
| e002051 | 19:58:48.847002 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/1/link_state |
| e002052 | 19:58:48.847002 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/1/nodes/2/flags |
| e002053 | 19:58:48.847002 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/7/snapshot/views/1/nodes/2/primary_id |
| e002054 | 19:58:48.847002 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/2/slots |
| e002055 | 19:58:48.847002 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/2/link_state |
| e002056 | 19:58:48.847002 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/7/snapshot/views/1/nodes/3/flags |
| e002057 | 19:58:48.847002 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/3/primary_id |
| e002058 | 19:58:48.847002 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/7/snapshot/views/1/nodes/3/slots |
| e002059 | 19:58:48.847002 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/3/link_state |
| e002060 | 19:58:48.847002 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/7/snapshot/views/1/nodes/4/flags |
| e002061 | 19:58:48.847002 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/7/snapshot/views/1/nodes/4/primary_id |
| e002062 | 19:58:48.847002 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/4/slots |
| e002063 | 19:58:48.847002 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/4/link_state |
| e002064 | 19:58:48.847002 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/7/snapshot/views/1/nodes/5/flags |
| e002065 | 19:58:48.847002 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/7/snapshot/views/1/nodes/5/primary_id |
| e002066 | 19:58:48.847002 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/7/snapshot/views/1/nodes/5/slots |
| e002067 | 19:58:48.847002 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/7/snapshot/views/1/nodes/5/link_state |
| e002068 | 19:58:48.847002 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/7/snapshot/views/1/persistence/aof_enabled |
| e002069 | 19:58:48.847002 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/7/snapshot/views/1/persistence/loading |
| e002070 | 19:58:48.847002 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/7/snapshot/views/1/persistence/aof_last_write_status |
| e002071 | 19:58:48.847002 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/7/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e002072 | 19:58:48.847002 | node-2 | admin | process_live | missing_from_schema | `` | /samples/7/snapshot/views/1/process_live |
| e002073 | 19:58:55.983629 |  | value_validation | validation.expected | valid | `191` | /validation/expected |
| e002074 | 19:58:55.983629 |  | value_validation | validation.checked | valid | `191` | /validation/checked |
| e002075 | 19:58:55.983629 |  | value_validation | validation.missing | valid | `191` | /validation/missing |
| e002076 | 19:58:55.983629 |  | value_validation | validation.corrupt | valid | `0` | /validation/corrupt |
| e002077 | 19:58:55.983629 |  | value_validation | validation.unreadable | valid | `0` | /validation/unreadable |
| e002078 | 19:58:55.983629 |  | value_validation | validation.unverified | valid | `0` | /validation/unverified |
| e002079 | 19:58:55.983629 |  | value_validation | validation.complete | valid | `true` | /validation/complete |
| e002080 | 19:58:55.984097 | node-4 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/3/reachable |
| e002081 | 19:58:55.984097 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/8/snapshot/views/3/replication/role |
| e002082 | 19:58:55.984097 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/8/snapshot/views/3/replication/connected_slaves |
| e002083 | 19:58:55.984097 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/8/snapshot/views/3/replication/master_link_status |
| e002084 | 19:58:55.984097 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/8/snapshot/views/3/replication/master_sync_in_progress |
| e002085 | 19:58:55.984097 | node-4 | info_replication | replication.master_repl_offset | valid | `14` | /samples/8/snapshot/views/3/replication/master_repl_offset |
| e002086 | 19:58:55.984097 | node-4 | info_replication | replication.slave_repl_offset | valid | `14` | /samples/8/snapshot/views/3/replication/slave_repl_offset |
| e002087 | 19:58:55.984097 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/8/snapshot/views/3/redis_id |
| e002088 | 19:58:55.984097 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/3/coverage |
| e002089 | 19:58:55.984097 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/3/nodes/0/flags |
| e002090 | 19:58:55.984097 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/0/primary_id |
| e002091 | 19:58:55.984097 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/3/nodes/0/slots |
| e002092 | 19:58:55.984097 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/0/link_state |
| e002093 | 19:58:55.984097 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/8/snapshot/views/3/nodes/1/flags |
| e002094 | 19:58:55.984097 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/8/snapshot/views/3/nodes/1/primary_id |
| e002095 | 19:58:55.984097 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/1/slots |
| e002096 | 19:58:55.984097 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/1/link_state |
| e002097 | 19:58:55.984097 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/3/nodes/2/flags |
| e002098 | 19:58:55.984097 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/8/snapshot/views/3/nodes/2/primary_id |
| e002099 | 19:58:55.984097 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/2/slots |
| e002100 | 19:58:55.984097 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/2/link_state |
| e002101 | 19:58:55.984097 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/3/nodes/3/flags |
| e002102 | 19:58:55.984097 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/3/primary_id |
| e002103 | 19:58:55.984097 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/3/nodes/3/slots |
| e002104 | 19:58:55.984097 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/3/link_state |
| e002105 | 19:58:55.984097 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/3/nodes/4/flags |
| e002106 | 19:58:55.984097 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/8/snapshot/views/3/nodes/4/primary_id |
| e002107 | 19:58:55.984097 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/4/slots |
| e002108 | 19:58:55.984097 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/4/link_state |
| e002109 | 19:58:55.984097 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/3/nodes/5/flags |
| e002110 | 19:58:55.984097 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/3/nodes/5/primary_id |
| e002111 | 19:58:55.984097 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/3/nodes/5/slots |
| e002112 | 19:58:55.984097 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/3/nodes/5/link_state |
| e002113 | 19:58:55.984097 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/3/persistence/aof_enabled |
| e002114 | 19:58:55.984097 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/3/persistence/loading |
| e002115 | 19:58:55.984097 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/3/persistence/aof_last_write_status |
| e002116 | 19:58:55.984097 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e002117 | 19:58:55.984097 | node-4 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/3/process_live |
| e002118 | 19:58:55.984112 | node-3 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/2/reachable |
| e002119 | 19:58:55.984112 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/8/snapshot/views/2/replication/role |
| e002120 | 19:58:55.984112 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/8/snapshot/views/2/replication/connected_slaves |
| e002121 | 19:58:55.984112 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/8/snapshot/views/2/replication/master_link_status |
| e002122 | 19:58:55.984112 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/8/snapshot/views/2/replication/master_sync_in_progress |
| e002123 | 19:58:55.984112 | node-3 | info_replication | replication.master_repl_offset | valid | `14` | /samples/8/snapshot/views/2/replication/master_repl_offset |
| e002124 | 19:58:55.984112 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/8/snapshot/views/2/replication/slave_repl_offset |
| e002125 | 19:58:55.984112 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/8/snapshot/views/2/redis_id |
| e002126 | 19:58:55.984112 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/2/coverage |
| e002127 | 19:58:55.984112 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/2/nodes/0/flags |
| e002128 | 19:58:55.984112 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/0/primary_id |
| e002129 | 19:58:55.984112 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/2/nodes/0/slots |
| e002130 | 19:58:55.984112 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/0/link_state |
| e002131 | 19:58:55.984112 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/8/snapshot/views/2/nodes/1/flags |
| e002132 | 19:58:55.984112 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/1/primary_id |
| e002133 | 19:58:55.984112 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/2/nodes/1/slots |
| e002134 | 19:58:55.984112 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/1/link_state |
| e002135 | 19:58:55.984112 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/2/nodes/2/flags |
| e002136 | 19:58:55.984112 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/8/snapshot/views/2/nodes/2/primary_id |
| e002137 | 19:58:55.984112 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/2/slots |
| e002138 | 19:58:55.984112 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/2/link_state |
| e002139 | 19:58:55.984112 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/2/nodes/3/flags |
| e002140 | 19:58:55.984112 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/8/snapshot/views/2/nodes/3/primary_id |
| e002141 | 19:58:55.984112 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/3/slots |
| e002142 | 19:58:55.984112 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/3/link_state |
| e002143 | 19:58:55.984112 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/2/nodes/4/flags |
| e002144 | 19:58:55.984112 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/8/snapshot/views/2/nodes/4/primary_id |
| e002145 | 19:58:55.984112 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/4/slots |
| e002146 | 19:58:55.984112 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/4/link_state |
| e002147 | 19:58:55.984112 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/2/nodes/5/flags |
| e002148 | 19:58:55.984112 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/2/nodes/5/primary_id |
| e002149 | 19:58:55.984112 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/2/nodes/5/slots |
| e002150 | 19:58:55.984112 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/2/nodes/5/link_state |
| e002151 | 19:58:55.984112 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/2/persistence/aof_enabled |
| e002152 | 19:58:55.984112 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/2/persistence/loading |
| e002153 | 19:58:55.984112 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/2/persistence/aof_last_write_status |
| e002154 | 19:58:55.984112 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e002155 | 19:58:55.984112 | node-3 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/2/process_live |
| e002156 | 19:58:55.984121 | node-5 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/4/reachable |
| e002157 | 19:58:55.984121 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/8/snapshot/views/4/replication/role |
| e002158 | 19:58:55.984121 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/8/snapshot/views/4/replication/connected_slaves |
| e002159 | 19:58:55.984121 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/8/snapshot/views/4/replication/master_link_status |
| e002160 | 19:58:55.984121 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/8/snapshot/views/4/replication/master_sync_in_progress |
| e002161 | 19:58:55.984121 | node-5 | info_replication | replication.master_repl_offset | valid | `28` | /samples/8/snapshot/views/4/replication/master_repl_offset |
| e002162 | 19:58:55.984121 | node-5 | info_replication | replication.slave_repl_offset | valid | `28` | /samples/8/snapshot/views/4/replication/slave_repl_offset |
| e002163 | 19:58:55.984121 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/8/snapshot/views/4/redis_id |
| e002164 | 19:58:55.984121 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/4/coverage |
| e002165 | 19:58:55.984121 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/4/nodes/0/flags |
| e002166 | 19:58:55.984121 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/0/primary_id |
| e002167 | 19:58:55.984121 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/4/nodes/0/slots |
| e002168 | 19:58:55.984121 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/0/link_state |
| e002169 | 19:58:55.984121 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/4/nodes/1/flags |
| e002170 | 19:58:55.984121 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/8/snapshot/views/4/nodes/1/primary_id |
| e002171 | 19:58:55.984121 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/1/slots |
| e002172 | 19:58:55.984121 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/1/link_state |
| e002173 | 19:58:55.984121 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/4/nodes/2/flags |
| e002174 | 19:58:55.984121 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/2/primary_id |
| e002175 | 19:58:55.984121 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/4/nodes/2/slots |
| e002176 | 19:58:55.984121 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/2/link_state |
| e002177 | 19:58:55.984121 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/4/nodes/3/flags |
| e002178 | 19:58:55.984121 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/3/primary_id |
| e002179 | 19:58:55.984121 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/4/nodes/3/slots |
| e002180 | 19:58:55.984121 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/3/link_state |
| e002181 | 19:58:55.984121 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/8/snapshot/views/4/nodes/4/flags |
| e002182 | 19:58:55.984121 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/8/snapshot/views/4/nodes/4/primary_id |
| e002183 | 19:58:55.984121 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/4/slots |
| e002184 | 19:58:55.984121 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/4/link_state |
| e002185 | 19:58:55.984121 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/4/nodes/5/flags |
| e002186 | 19:58:55.984121 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/8/snapshot/views/4/nodes/5/primary_id |
| e002187 | 19:58:55.984121 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/4/nodes/5/slots |
| e002188 | 19:58:55.984121 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/4/nodes/5/link_state |
| e002189 | 19:58:55.984121 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/4/persistence/aof_enabled |
| e002190 | 19:58:55.984121 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/4/persistence/loading |
| e002191 | 19:58:55.984121 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/4/persistence/aof_last_write_status |
| e002192 | 19:58:55.984121 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e002193 | 19:58:55.984121 | node-5 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/4/process_live |
| e002194 | 19:58:55.984142 | node-6 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/5/reachable |
| e002195 | 19:58:55.984142 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/8/snapshot/views/5/replication/role |
| e002196 | 19:58:55.984142 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/8/snapshot/views/5/replication/connected_slaves |
| e002197 | 19:58:55.984142 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/8/snapshot/views/5/replication/master_link_status |
| e002198 | 19:58:55.984142 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/8/snapshot/views/5/replication/master_sync_in_progress |
| e002199 | 19:58:55.984142 | node-6 | info_replication | replication.master_repl_offset | valid | `14` | /samples/8/snapshot/views/5/replication/master_repl_offset |
| e002200 | 19:58:55.984142 | node-6 | info_replication | replication.slave_repl_offset | valid | `14` | /samples/8/snapshot/views/5/replication/slave_repl_offset |
| e002201 | 19:58:55.984142 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/8/snapshot/views/5/redis_id |
| e002202 | 19:58:55.984142 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/5/coverage |
| e002203 | 19:58:55.984142 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/5/nodes/0/flags |
| e002204 | 19:58:55.984142 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/8/snapshot/views/5/nodes/0/primary_id |
| e002205 | 19:58:55.984142 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/0/slots |
| e002206 | 19:58:55.984142 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/0/link_state |
| e002207 | 19:58:55.984142 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/5/nodes/1/flags |
| e002208 | 19:58:55.984142 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/1/primary_id |
| e002209 | 19:58:55.984142 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/5/nodes/1/slots |
| e002210 | 19:58:55.984142 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/1/link_state |
| e002211 | 19:58:55.984142 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/5/nodes/2/flags |
| e002212 | 19:58:55.984142 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/2/primary_id |
| e002213 | 19:58:55.984142 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/5/nodes/2/slots |
| e002214 | 19:58:55.984142 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/2/link_state |
| e002215 | 19:58:55.984142 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/5/nodes/3/flags |
| e002216 | 19:58:55.984142 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/3/primary_id |
| e002217 | 19:58:55.984142 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/5/nodes/3/slots |
| e002218 | 19:58:55.984142 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/3/link_state |
| e002219 | 19:58:55.984142 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/5/nodes/4/flags |
| e002220 | 19:58:55.984142 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/8/snapshot/views/5/nodes/4/primary_id |
| e002221 | 19:58:55.984142 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/4/slots |
| e002222 | 19:58:55.984142 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/4/link_state |
| e002223 | 19:58:55.984142 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/8/snapshot/views/5/nodes/5/flags |
| e002224 | 19:58:55.984142 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/8/snapshot/views/5/nodes/5/primary_id |
| e002225 | 19:58:55.984142 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/5/nodes/5/slots |
| e002226 | 19:58:55.984142 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/5/nodes/5/link_state |
| e002227 | 19:58:55.984142 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/5/persistence/aof_enabled |
| e002228 | 19:58:55.984142 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/5/persistence/loading |
| e002229 | 19:58:55.984142 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/5/persistence/aof_last_write_status |
| e002230 | 19:58:55.984142 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002231 | 19:58:55.984142 | node-6 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/5/process_live |
| e002232 | 19:58:55.984456 | node-1 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/0/reachable |
| e002233 | 19:58:55.984456 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/8/snapshot/views/0/replication/role |
| e002234 | 19:58:55.984456 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/8/snapshot/views/0/replication/connected_slaves |
| e002235 | 19:58:55.984456 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/8/snapshot/views/0/replication/master_link_status |
| e002236 | 19:58:55.984456 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/8/snapshot/views/0/replication/master_sync_in_progress |
| e002237 | 19:58:55.984456 | node-1 | info_replication | replication.master_repl_offset | valid | `28` | /samples/8/snapshot/views/0/replication/master_repl_offset |
| e002238 | 19:58:55.984456 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/8/snapshot/views/0/replication/slave_repl_offset |
| e002239 | 19:58:55.984456 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/8/snapshot/views/0/redis_id |
| e002240 | 19:58:55.984456 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/0/coverage |
| e002241 | 19:58:55.984456 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/0/nodes/0/flags |
| e002242 | 19:58:55.984456 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/8/snapshot/views/0/nodes/0/primary_id |
| e002243 | 19:58:55.984456 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/0/slots |
| e002244 | 19:58:55.984456 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/0/link_state |
| e002245 | 19:58:55.984456 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/0/nodes/1/flags |
| e002246 | 19:58:55.984456 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/1/primary_id |
| e002247 | 19:58:55.984456 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/0/nodes/1/slots |
| e002248 | 19:58:55.984456 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/1/link_state |
| e002249 | 19:58:55.984456 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/0/nodes/2/flags |
| e002250 | 19:58:55.984456 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/2/primary_id |
| e002251 | 19:58:55.984456 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/0/nodes/2/slots |
| e002252 | 19:58:55.984456 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/2/link_state |
| e002253 | 19:58:55.984456 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/8/snapshot/views/0/nodes/3/flags |
| e002254 | 19:58:55.984456 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/3/primary_id |
| e002255 | 19:58:55.984456 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/0/nodes/3/slots |
| e002256 | 19:58:55.984456 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/3/link_state |
| e002257 | 19:58:55.984456 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/0/nodes/4/flags |
| e002258 | 19:58:55.984456 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/8/snapshot/views/0/nodes/4/primary_id |
| e002259 | 19:58:55.984456 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/4/slots |
| e002260 | 19:58:55.984456 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/4/link_state |
| e002261 | 19:58:55.984456 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/0/nodes/5/flags |
| e002262 | 19:58:55.984456 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/8/snapshot/views/0/nodes/5/primary_id |
| e002263 | 19:58:55.984456 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/0/nodes/5/slots |
| e002264 | 19:58:55.984456 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/0/nodes/5/link_state |
| e002265 | 19:58:55.984456 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/0/persistence/aof_enabled |
| e002266 | 19:58:55.984456 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/0/persistence/loading |
| e002267 | 19:58:55.984456 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/0/persistence/aof_last_write_status |
| e002268 | 19:58:55.984456 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e002269 | 19:58:55.984456 | node-1 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/0/process_live |
| e002270 | 19:58:55.984576 | node-2 | node_client | reachable | valid | `true` | /samples/8/snapshot/views/1/reachable |
| e002271 | 19:58:55.984576 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/8/snapshot/views/1/replication/role |
| e002272 | 19:58:55.984576 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/8/snapshot/views/1/replication/connected_slaves |
| e002273 | 19:58:55.984576 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/8/snapshot/views/1/replication/master_link_status |
| e002274 | 19:58:55.984576 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/8/snapshot/views/1/replication/master_sync_in_progress |
| e002275 | 19:58:55.984576 | node-2 | info_replication | replication.master_repl_offset | valid | `14` | /samples/8/snapshot/views/1/replication/master_repl_offset |
| e002276 | 19:58:55.984576 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/8/snapshot/views/1/replication/slave_repl_offset |
| e002277 | 19:58:55.984576 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/8/snapshot/views/1/redis_id |
| e002278 | 19:58:55.984576 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/8/snapshot/views/1/coverage |
| e002279 | 19:58:55.984576 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/1/nodes/0/flags |
| e002280 | 19:58:55.984576 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/8/snapshot/views/1/nodes/0/primary_id |
| e002281 | 19:58:55.984576 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/0/slots |
| e002282 | 19:58:55.984576 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/0/link_state |
| e002283 | 19:58:55.984576 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/1/nodes/1/flags |
| e002284 | 19:58:55.984576 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/1/primary_id |
| e002285 | 19:58:55.984576 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/8/snapshot/views/1/nodes/1/slots |
| e002286 | 19:58:55.984576 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/1/link_state |
| e002287 | 19:58:55.984576 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/1/nodes/2/flags |
| e002288 | 19:58:55.984576 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/8/snapshot/views/1/nodes/2/primary_id |
| e002289 | 19:58:55.984576 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/2/slots |
| e002290 | 19:58:55.984576 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/2/link_state |
| e002291 | 19:58:55.984576 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/8/snapshot/views/1/nodes/3/flags |
| e002292 | 19:58:55.984576 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/3/primary_id |
| e002293 | 19:58:55.984576 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/8/snapshot/views/1/nodes/3/slots |
| e002294 | 19:58:55.984576 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/3/link_state |
| e002295 | 19:58:55.984576 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/8/snapshot/views/1/nodes/4/flags |
| e002296 | 19:58:55.984576 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/8/snapshot/views/1/nodes/4/primary_id |
| e002297 | 19:58:55.984576 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/4/slots |
| e002298 | 19:58:55.984576 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/4/link_state |
| e002299 | 19:58:55.984576 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/8/snapshot/views/1/nodes/5/flags |
| e002300 | 19:58:55.984576 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/8/snapshot/views/1/nodes/5/primary_id |
| e002301 | 19:58:55.984576 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/8/snapshot/views/1/nodes/5/slots |
| e002302 | 19:58:55.984576 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/8/snapshot/views/1/nodes/5/link_state |
| e002303 | 19:58:55.984576 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/8/snapshot/views/1/persistence/aof_enabled |
| e002304 | 19:58:55.984576 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/8/snapshot/views/1/persistence/loading |
| e002305 | 19:58:55.984576 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/8/snapshot/views/1/persistence/aof_last_write_status |
| e002306 | 19:58:55.984576 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/8/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e002307 | 19:58:55.984576 | node-2 | admin | process_live | missing_from_schema | `` | /samples/8/snapshot/views/1/process_live |
| e002308 | 19:58:55.991023 |  | client_probe | client_success | valid | `true` | /samples/8/probe/successful |
| e002309 | 19:58:56.576526 | node-2 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/1/reachable |
| e002310 | 19:58:56.576526 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/9/snapshot/views/1/replication/role |
| e002311 | 19:58:56.576526 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/9/snapshot/views/1/replication/connected_slaves |
| e002312 | 19:58:56.576526 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/9/snapshot/views/1/replication/master_link_status |
| e002313 | 19:58:56.576526 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/9/snapshot/views/1/replication/master_sync_in_progress |
| e002314 | 19:58:56.576526 | node-2 | info_replication | replication.master_repl_offset | valid | `138` | /samples/9/snapshot/views/1/replication/master_repl_offset |
| e002315 | 19:58:56.576526 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/9/snapshot/views/1/replication/slave_repl_offset |
| e002316 | 19:58:56.576526 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/9/snapshot/views/1/redis_id |
| e002317 | 19:58:56.576526 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/1/coverage |
| e002318 | 19:58:56.576526 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/1/nodes/0/flags |
| e002319 | 19:58:56.576526 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/9/snapshot/views/1/nodes/0/primary_id |
| e002320 | 19:58:56.576526 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/0/slots |
| e002321 | 19:58:56.576526 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/0/link_state |
| e002322 | 19:58:56.576526 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/1/nodes/1/flags |
| e002323 | 19:58:56.576526 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/1/primary_id |
| e002324 | 19:58:56.576526 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/1/nodes/1/slots |
| e002325 | 19:58:56.576526 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/1/link_state |
| e002326 | 19:58:56.576526 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/1/nodes/2/flags |
| e002327 | 19:58:56.576526 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/9/snapshot/views/1/nodes/2/primary_id |
| e002328 | 19:58:56.576526 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/2/slots |
| e002329 | 19:58:56.576526 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/2/link_state |
| e002330 | 19:58:56.576526 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/1/nodes/3/flags |
| e002331 | 19:58:56.576526 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/3/primary_id |
| e002332 | 19:58:56.576526 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/1/nodes/3/slots |
| e002333 | 19:58:56.576526 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/3/link_state |
| e002334 | 19:58:56.576526 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/1/nodes/4/flags |
| e002335 | 19:58:56.576526 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/9/snapshot/views/1/nodes/4/primary_id |
| e002336 | 19:58:56.576526 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/4/slots |
| e002337 | 19:58:56.576526 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/4/link_state |
| e002338 | 19:58:56.576526 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/9/snapshot/views/1/nodes/5/flags |
| e002339 | 19:58:56.576526 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/1/nodes/5/primary_id |
| e002340 | 19:58:56.576526 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/1/nodes/5/slots |
| e002341 | 19:58:56.576526 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/1/nodes/5/link_state |
| e002342 | 19:58:56.576526 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/1/persistence/aof_enabled |
| e002343 | 19:58:56.576526 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/1/persistence/loading |
| e002344 | 19:58:56.576526 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/1/persistence/aof_last_write_status |
| e002345 | 19:58:56.576526 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e002346 | 19:58:56.576526 | node-2 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/1/process_live |
| e002347 | 19:58:56.576531 | node-1 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/0/reachable |
| e002348 | 19:58:56.576531 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/9/snapshot/views/0/replication/role |
| e002349 | 19:58:56.576531 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/9/snapshot/views/0/replication/connected_slaves |
| e002350 | 19:58:56.576531 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/9/snapshot/views/0/replication/master_link_status |
| e002351 | 19:58:56.576531 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/9/snapshot/views/0/replication/master_sync_in_progress |
| e002352 | 19:58:56.576531 | node-1 | info_replication | replication.master_repl_offset | valid | `152` | /samples/9/snapshot/views/0/replication/master_repl_offset |
| e002353 | 19:58:56.576531 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/9/snapshot/views/0/replication/slave_repl_offset |
| e002354 | 19:58:56.576531 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/9/snapshot/views/0/redis_id |
| e002355 | 19:58:56.576531 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/0/coverage |
| e002356 | 19:58:56.576531 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/0/nodes/0/flags |
| e002357 | 19:58:56.576531 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/9/snapshot/views/0/nodes/0/primary_id |
| e002358 | 19:58:56.576531 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/0/slots |
| e002359 | 19:58:56.576531 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/0/link_state |
| e002360 | 19:58:56.576531 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/0/nodes/1/flags |
| e002361 | 19:58:56.576531 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/1/primary_id |
| e002362 | 19:58:56.576531 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/0/nodes/1/slots |
| e002363 | 19:58:56.576531 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/1/link_state |
| e002364 | 19:58:56.576531 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/0/nodes/2/flags |
| e002365 | 19:58:56.576531 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/2/primary_id |
| e002366 | 19:58:56.576531 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/0/nodes/2/slots |
| e002367 | 19:58:56.576531 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/2/link_state |
| e002368 | 19:58:56.576531 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/9/snapshot/views/0/nodes/3/flags |
| e002369 | 19:58:56.576531 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/3/primary_id |
| e002370 | 19:58:56.576531 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/0/nodes/3/slots |
| e002371 | 19:58:56.576531 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/3/link_state |
| e002372 | 19:58:56.576531 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/0/nodes/4/flags |
| e002373 | 19:58:56.576531 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/9/snapshot/views/0/nodes/4/primary_id |
| e002374 | 19:58:56.576531 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/4/slots |
| e002375 | 19:58:56.576531 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/4/link_state |
| e002376 | 19:58:56.576531 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/0/nodes/5/flags |
| e002377 | 19:58:56.576531 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/9/snapshot/views/0/nodes/5/primary_id |
| e002378 | 19:58:56.576531 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/0/nodes/5/slots |
| e002379 | 19:58:56.576531 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/0/nodes/5/link_state |
| e002380 | 19:58:56.576531 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/0/persistence/aof_enabled |
| e002381 | 19:58:56.576531 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/0/persistence/loading |
| e002382 | 19:58:56.576531 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/0/persistence/aof_last_write_status |
| e002383 | 19:58:56.576531 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e002384 | 19:58:56.576531 | node-1 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/0/process_live |
| e002385 | 19:58:56.57654 | node-4 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/3/reachable |
| e002386 | 19:58:56.57654 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/9/snapshot/views/3/replication/role |
| e002387 | 19:58:56.57654 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/9/snapshot/views/3/replication/connected_slaves |
| e002388 | 19:58:56.57654 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/9/snapshot/views/3/replication/master_link_status |
| e002389 | 19:58:56.57654 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/9/snapshot/views/3/replication/master_sync_in_progress |
| e002390 | 19:58:56.57654 | node-4 | info_replication | replication.master_repl_offset | valid | `138` | /samples/9/snapshot/views/3/replication/master_repl_offset |
| e002391 | 19:58:56.57654 | node-4 | info_replication | replication.slave_repl_offset | valid | `138` | /samples/9/snapshot/views/3/replication/slave_repl_offset |
| e002392 | 19:58:56.57654 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/9/snapshot/views/3/redis_id |
| e002393 | 19:58:56.57654 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/3/coverage |
| e002394 | 19:58:56.57654 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/3/nodes/0/flags |
| e002395 | 19:58:56.57654 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/0/primary_id |
| e002396 | 19:58:56.57654 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/3/nodes/0/slots |
| e002397 | 19:58:56.57654 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/0/link_state |
| e002398 | 19:58:56.57654 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/9/snapshot/views/3/nodes/1/flags |
| e002399 | 19:58:56.57654 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/9/snapshot/views/3/nodes/1/primary_id |
| e002400 | 19:58:56.57654 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/1/slots |
| e002401 | 19:58:56.57654 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/1/link_state |
| e002402 | 19:58:56.57654 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/3/nodes/2/flags |
| e002403 | 19:58:56.57654 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/9/snapshot/views/3/nodes/2/primary_id |
| e002404 | 19:58:56.57654 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/2/slots |
| e002405 | 19:58:56.57654 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/2/link_state |
| e002406 | 19:58:56.57654 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/3/nodes/3/flags |
| e002407 | 19:58:56.57654 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/3/primary_id |
| e002408 | 19:58:56.57654 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/3/nodes/3/slots |
| e002409 | 19:58:56.57654 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/3/link_state |
| e002410 | 19:58:56.57654 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/3/nodes/4/flags |
| e002411 | 19:58:56.57654 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/9/snapshot/views/3/nodes/4/primary_id |
| e002412 | 19:58:56.57654 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/4/slots |
| e002413 | 19:58:56.57654 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/4/link_state |
| e002414 | 19:58:56.57654 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/3/nodes/5/flags |
| e002415 | 19:58:56.57654 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/3/nodes/5/primary_id |
| e002416 | 19:58:56.57654 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/3/nodes/5/slots |
| e002417 | 19:58:56.57654 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/3/nodes/5/link_state |
| e002418 | 19:58:56.57654 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/3/persistence/aof_enabled |
| e002419 | 19:58:56.57654 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/3/persistence/loading |
| e002420 | 19:58:56.57654 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/3/persistence/aof_last_write_status |
| e002421 | 19:58:56.57654 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e002422 | 19:58:56.57654 | node-4 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/3/process_live |
| e002423 | 19:58:56.576573 | node-5 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/4/reachable |
| e002424 | 19:58:56.576573 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/9/snapshot/views/4/replication/role |
| e002425 | 19:58:56.576573 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/9/snapshot/views/4/replication/connected_slaves |
| e002426 | 19:58:56.576573 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/9/snapshot/views/4/replication/master_link_status |
| e002427 | 19:58:56.576573 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/9/snapshot/views/4/replication/master_sync_in_progress |
| e002428 | 19:58:56.576573 | node-5 | info_replication | replication.master_repl_offset | valid | `152` | /samples/9/snapshot/views/4/replication/master_repl_offset |
| e002429 | 19:58:56.576573 | node-5 | info_replication | replication.slave_repl_offset | valid | `152` | /samples/9/snapshot/views/4/replication/slave_repl_offset |
| e002430 | 19:58:56.576573 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/9/snapshot/views/4/redis_id |
| e002431 | 19:58:56.576573 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/4/coverage |
| e002432 | 19:58:56.576573 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/4/nodes/0/flags |
| e002433 | 19:58:56.576573 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/0/primary_id |
| e002434 | 19:58:56.576573 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/4/nodes/0/slots |
| e002435 | 19:58:56.576573 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/0/link_state |
| e002436 | 19:58:56.576573 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/4/nodes/1/flags |
| e002437 | 19:58:56.576573 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/9/snapshot/views/4/nodes/1/primary_id |
| e002438 | 19:58:56.576573 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/1/slots |
| e002439 | 19:58:56.576573 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/1/link_state |
| e002440 | 19:58:56.576573 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/4/nodes/2/flags |
| e002441 | 19:58:56.576573 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/2/primary_id |
| e002442 | 19:58:56.576573 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/4/nodes/2/slots |
| e002443 | 19:58:56.576573 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/2/link_state |
| e002444 | 19:58:56.576573 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/4/nodes/3/flags |
| e002445 | 19:58:56.576573 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/3/primary_id |
| e002446 | 19:58:56.576573 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/4/nodes/3/slots |
| e002447 | 19:58:56.576573 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/3/link_state |
| e002448 | 19:58:56.576573 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/9/snapshot/views/4/nodes/4/flags |
| e002449 | 19:58:56.576573 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/9/snapshot/views/4/nodes/4/primary_id |
| e002450 | 19:58:56.576573 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/4/slots |
| e002451 | 19:58:56.576573 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/4/link_state |
| e002452 | 19:58:56.576573 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/4/nodes/5/flags |
| e002453 | 19:58:56.576573 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/9/snapshot/views/4/nodes/5/primary_id |
| e002454 | 19:58:56.576573 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/4/nodes/5/slots |
| e002455 | 19:58:56.576573 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/4/nodes/5/link_state |
| e002456 | 19:58:56.576573 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/4/persistence/aof_enabled |
| e002457 | 19:58:56.576573 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/4/persistence/loading |
| e002458 | 19:58:56.576573 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/4/persistence/aof_last_write_status |
| e002459 | 19:58:56.576573 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e002460 | 19:58:56.576573 | node-5 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/4/process_live |
| e002461 | 19:58:56.576582 | node-3 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/2/reachable |
| e002462 | 19:58:56.576582 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/9/snapshot/views/2/replication/role |
| e002463 | 19:58:56.576582 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/9/snapshot/views/2/replication/connected_slaves |
| e002464 | 19:58:56.576582 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/9/snapshot/views/2/replication/master_link_status |
| e002465 | 19:58:56.576582 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/9/snapshot/views/2/replication/master_sync_in_progress |
| e002466 | 19:58:56.576582 | node-3 | info_replication | replication.master_repl_offset | valid | `138` | /samples/9/snapshot/views/2/replication/master_repl_offset |
| e002467 | 19:58:56.576582 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/9/snapshot/views/2/replication/slave_repl_offset |
| e002468 | 19:58:56.576582 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/9/snapshot/views/2/redis_id |
| e002469 | 19:58:56.576582 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/2/coverage |
| e002470 | 19:58:56.576582 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/2/nodes/0/flags |
| e002471 | 19:58:56.576582 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/0/primary_id |
| e002472 | 19:58:56.576582 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/2/nodes/0/slots |
| e002473 | 19:58:56.576582 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/0/link_state |
| e002474 | 19:58:56.576582 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/9/snapshot/views/2/nodes/1/flags |
| e002475 | 19:58:56.576582 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/1/primary_id |
| e002476 | 19:58:56.576582 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/2/nodes/1/slots |
| e002477 | 19:58:56.576582 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/1/link_state |
| e002478 | 19:58:56.576582 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/2/nodes/2/flags |
| e002479 | 19:58:56.576582 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/9/snapshot/views/2/nodes/2/primary_id |
| e002480 | 19:58:56.576582 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/2/slots |
| e002481 | 19:58:56.576582 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/2/link_state |
| e002482 | 19:58:56.576582 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/2/nodes/3/flags |
| e002483 | 19:58:56.576582 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/9/snapshot/views/2/nodes/3/primary_id |
| e002484 | 19:58:56.576582 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/3/slots |
| e002485 | 19:58:56.576582 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/3/link_state |
| e002486 | 19:58:56.576582 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/2/nodes/4/flags |
| e002487 | 19:58:56.576582 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/9/snapshot/views/2/nodes/4/primary_id |
| e002488 | 19:58:56.576582 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/4/slots |
| e002489 | 19:58:56.576582 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/4/link_state |
| e002490 | 19:58:56.576582 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/2/nodes/5/flags |
| e002491 | 19:58:56.576582 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/2/nodes/5/primary_id |
| e002492 | 19:58:56.576582 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/2/nodes/5/slots |
| e002493 | 19:58:56.576582 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/2/nodes/5/link_state |
| e002494 | 19:58:56.576582 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/2/persistence/aof_enabled |
| e002495 | 19:58:56.576582 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/2/persistence/loading |
| e002496 | 19:58:56.576582 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/2/persistence/aof_last_write_status |
| e002497 | 19:58:56.576582 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e002498 | 19:58:56.576582 | node-3 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/2/process_live |
| e002499 | 19:58:56.576588 | node-6 | node_client | reachable | valid | `true` | /samples/9/snapshot/views/5/reachable |
| e002500 | 19:58:56.576588 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/9/snapshot/views/5/replication/role |
| e002501 | 19:58:56.576588 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/9/snapshot/views/5/replication/connected_slaves |
| e002502 | 19:58:56.576588 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/9/snapshot/views/5/replication/master_link_status |
| e002503 | 19:58:56.576588 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/9/snapshot/views/5/replication/master_sync_in_progress |
| e002504 | 19:58:56.576588 | node-6 | info_replication | replication.master_repl_offset | valid | `138` | /samples/9/snapshot/views/5/replication/master_repl_offset |
| e002505 | 19:58:56.576588 | node-6 | info_replication | replication.slave_repl_offset | valid | `138` | /samples/9/snapshot/views/5/replication/slave_repl_offset |
| e002506 | 19:58:56.576588 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/9/snapshot/views/5/redis_id |
| e002507 | 19:58:56.576588 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/9/snapshot/views/5/coverage |
| e002508 | 19:58:56.576588 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/5/nodes/0/flags |
| e002509 | 19:58:56.576588 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/9/snapshot/views/5/nodes/0/primary_id |
| e002510 | 19:58:56.576588 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/0/slots |
| e002511 | 19:58:56.576588 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/0/link_state |
| e002512 | 19:58:56.576588 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/5/nodes/1/flags |
| e002513 | 19:58:56.576588 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/1/primary_id |
| e002514 | 19:58:56.576588 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/9/snapshot/views/5/nodes/1/slots |
| e002515 | 19:58:56.576588 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/1/link_state |
| e002516 | 19:58:56.576588 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/5/nodes/2/flags |
| e002517 | 19:58:56.576588 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/2/primary_id |
| e002518 | 19:58:56.576588 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/9/snapshot/views/5/nodes/2/slots |
| e002519 | 19:58:56.576588 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/2/link_state |
| e002520 | 19:58:56.576588 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/9/snapshot/views/5/nodes/3/flags |
| e002521 | 19:58:56.576588 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/3/primary_id |
| e002522 | 19:58:56.576588 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/9/snapshot/views/5/nodes/3/slots |
| e002523 | 19:58:56.576588 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/3/link_state |
| e002524 | 19:58:56.576588 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/9/snapshot/views/5/nodes/4/flags |
| e002525 | 19:58:56.576588 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/9/snapshot/views/5/nodes/4/primary_id |
| e002526 | 19:58:56.576588 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/4/slots |
| e002527 | 19:58:56.576588 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/4/link_state |
| e002528 | 19:58:56.576588 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/9/snapshot/views/5/nodes/5/flags |
| e002529 | 19:58:56.576588 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/9/snapshot/views/5/nodes/5/primary_id |
| e002530 | 19:58:56.576588 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/9/snapshot/views/5/nodes/5/slots |
| e002531 | 19:58:56.576588 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/9/snapshot/views/5/nodes/5/link_state |
| e002532 | 19:58:56.576588 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/9/snapshot/views/5/persistence/aof_enabled |
| e002533 | 19:58:56.576588 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/9/snapshot/views/5/persistence/loading |
| e002534 | 19:58:56.576588 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/9/snapshot/views/5/persistence/aof_last_write_status |
| e002535 | 19:58:56.576588 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/9/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002536 | 19:58:56.576588 | node-6 | admin | process_live | missing_from_schema | `` | /samples/9/snapshot/views/5/process_live |
| e002537 | 19:58:56.583224 |  | client_probe | client_success | valid | `true` | /samples/9/probe/successful |
| e002538 | 19:58:57.166774 | node-1 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/0/reachable |
| e002539 | 19:58:57.166774 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/10/snapshot/views/0/replication/role |
| e002540 | 19:58:57.166774 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/10/snapshot/views/0/replication/connected_slaves |
| e002541 | 19:58:57.166774 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/10/snapshot/views/0/replication/master_link_status |
| e002542 | 19:58:57.166774 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/10/snapshot/views/0/replication/master_sync_in_progress |
| e002543 | 19:58:57.166774 | node-1 | info_replication | replication.master_repl_offset | valid | `253` | /samples/10/snapshot/views/0/replication/master_repl_offset |
| e002544 | 19:58:57.166774 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/10/snapshot/views/0/replication/slave_repl_offset |
| e002545 | 19:58:57.166774 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/10/snapshot/views/0/redis_id |
| e002546 | 19:58:57.166774 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/0/coverage |
| e002547 | 19:58:57.166774 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/0/nodes/0/flags |
| e002548 | 19:58:57.166774 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/10/snapshot/views/0/nodes/0/primary_id |
| e002549 | 19:58:57.166774 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/0/slots |
| e002550 | 19:58:57.166774 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/0/link_state |
| e002551 | 19:58:57.166774 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/0/nodes/1/flags |
| e002552 | 19:58:57.166774 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/1/primary_id |
| e002553 | 19:58:57.166774 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/0/nodes/1/slots |
| e002554 | 19:58:57.166774 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/1/link_state |
| e002555 | 19:58:57.166774 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/0/nodes/2/flags |
| e002556 | 19:58:57.166774 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/2/primary_id |
| e002557 | 19:58:57.166774 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/0/nodes/2/slots |
| e002558 | 19:58:57.166774 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/2/link_state |
| e002559 | 19:58:57.166774 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/10/snapshot/views/0/nodes/3/flags |
| e002560 | 19:58:57.166774 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/3/primary_id |
| e002561 | 19:58:57.166774 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/0/nodes/3/slots |
| e002562 | 19:58:57.166774 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/3/link_state |
| e002563 | 19:58:57.166774 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/0/nodes/4/flags |
| e002564 | 19:58:57.166774 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/10/snapshot/views/0/nodes/4/primary_id |
| e002565 | 19:58:57.166774 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/4/slots |
| e002566 | 19:58:57.166774 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/4/link_state |
| e002567 | 19:58:57.166774 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/0/nodes/5/flags |
| e002568 | 19:58:57.166774 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/10/snapshot/views/0/nodes/5/primary_id |
| e002569 | 19:58:57.166774 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/0/nodes/5/slots |
| e002570 | 19:58:57.166774 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/0/nodes/5/link_state |
| e002571 | 19:58:57.166774 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/0/persistence/aof_enabled |
| e002572 | 19:58:57.166774 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/0/persistence/loading |
| e002573 | 19:58:57.166774 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/0/persistence/aof_last_write_status |
| e002574 | 19:58:57.166774 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e002575 | 19:58:57.166774 | node-1 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/0/process_live |
| e002576 | 19:58:57.166791 | node-6 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/5/reachable |
| e002577 | 19:58:57.166791 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/10/snapshot/views/5/replication/role |
| e002578 | 19:58:57.166791 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/10/snapshot/views/5/replication/connected_slaves |
| e002579 | 19:58:57.166791 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/10/snapshot/views/5/replication/master_link_status |
| e002580 | 19:58:57.166791 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/10/snapshot/views/5/replication/master_sync_in_progress |
| e002581 | 19:58:57.166791 | node-6 | info_replication | replication.master_repl_offset | valid | `239` | /samples/10/snapshot/views/5/replication/master_repl_offset |
| e002582 | 19:58:57.166791 | node-6 | info_replication | replication.slave_repl_offset | valid | `239` | /samples/10/snapshot/views/5/replication/slave_repl_offset |
| e002583 | 19:58:57.166791 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/10/snapshot/views/5/redis_id |
| e002584 | 19:58:57.166791 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/5/coverage |
| e002585 | 19:58:57.166791 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/5/nodes/0/flags |
| e002586 | 19:58:57.166791 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/10/snapshot/views/5/nodes/0/primary_id |
| e002587 | 19:58:57.166791 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/0/slots |
| e002588 | 19:58:57.166791 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/0/link_state |
| e002589 | 19:58:57.166791 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/5/nodes/1/flags |
| e002590 | 19:58:57.166791 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/1/primary_id |
| e002591 | 19:58:57.166791 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/5/nodes/1/slots |
| e002592 | 19:58:57.166791 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/1/link_state |
| e002593 | 19:58:57.166791 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/5/nodes/2/flags |
| e002594 | 19:58:57.166791 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/2/primary_id |
| e002595 | 19:58:57.166791 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/5/nodes/2/slots |
| e002596 | 19:58:57.166791 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/2/link_state |
| e002597 | 19:58:57.166791 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/5/nodes/3/flags |
| e002598 | 19:58:57.166791 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/3/primary_id |
| e002599 | 19:58:57.166791 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/5/nodes/3/slots |
| e002600 | 19:58:57.166791 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/3/link_state |
| e002601 | 19:58:57.166791 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/5/nodes/4/flags |
| e002602 | 19:58:57.166791 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/10/snapshot/views/5/nodes/4/primary_id |
| e002603 | 19:58:57.166791 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/4/slots |
| e002604 | 19:58:57.166791 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/4/link_state |
| e002605 | 19:58:57.166791 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/10/snapshot/views/5/nodes/5/flags |
| e002606 | 19:58:57.166791 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/10/snapshot/views/5/nodes/5/primary_id |
| e002607 | 19:58:57.166791 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/5/nodes/5/slots |
| e002608 | 19:58:57.166791 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/5/nodes/5/link_state |
| e002609 | 19:58:57.166791 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/5/persistence/aof_enabled |
| e002610 | 19:58:57.166791 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/5/persistence/loading |
| e002611 | 19:58:57.166791 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/5/persistence/aof_last_write_status |
| e002612 | 19:58:57.166791 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002613 | 19:58:57.166791 | node-6 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/5/process_live |
| e002614 | 19:58:57.166798 | node-5 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/4/reachable |
| e002615 | 19:58:57.166798 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/10/snapshot/views/4/replication/role |
| e002616 | 19:58:57.166798 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/10/snapshot/views/4/replication/connected_slaves |
| e002617 | 19:58:57.166798 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/10/snapshot/views/4/replication/master_link_status |
| e002618 | 19:58:57.166798 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/10/snapshot/views/4/replication/master_sync_in_progress |
| e002619 | 19:58:57.166798 | node-5 | info_replication | replication.master_repl_offset | valid | `253` | /samples/10/snapshot/views/4/replication/master_repl_offset |
| e002620 | 19:58:57.166798 | node-5 | info_replication | replication.slave_repl_offset | valid | `253` | /samples/10/snapshot/views/4/replication/slave_repl_offset |
| e002621 | 19:58:57.166798 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/10/snapshot/views/4/redis_id |
| e002622 | 19:58:57.166798 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/4/coverage |
| e002623 | 19:58:57.166798 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/4/nodes/0/flags |
| e002624 | 19:58:57.166798 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/0/primary_id |
| e002625 | 19:58:57.166798 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/4/nodes/0/slots |
| e002626 | 19:58:57.166798 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/0/link_state |
| e002627 | 19:58:57.166798 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/4/nodes/1/flags |
| e002628 | 19:58:57.166798 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/10/snapshot/views/4/nodes/1/primary_id |
| e002629 | 19:58:57.166798 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/1/slots |
| e002630 | 19:58:57.166798 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/1/link_state |
| e002631 | 19:58:57.166798 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/4/nodes/2/flags |
| e002632 | 19:58:57.166798 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/2/primary_id |
| e002633 | 19:58:57.166798 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/4/nodes/2/slots |
| e002634 | 19:58:57.166798 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/2/link_state |
| e002635 | 19:58:57.166798 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/4/nodes/3/flags |
| e002636 | 19:58:57.166798 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/3/primary_id |
| e002637 | 19:58:57.166798 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/4/nodes/3/slots |
| e002638 | 19:58:57.166798 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/3/link_state |
| e002639 | 19:58:57.166798 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/10/snapshot/views/4/nodes/4/flags |
| e002640 | 19:58:57.166798 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/10/snapshot/views/4/nodes/4/primary_id |
| e002641 | 19:58:57.166798 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/4/slots |
| e002642 | 19:58:57.166798 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/4/link_state |
| e002643 | 19:58:57.166798 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/4/nodes/5/flags |
| e002644 | 19:58:57.166798 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/10/snapshot/views/4/nodes/5/primary_id |
| e002645 | 19:58:57.166798 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/4/nodes/5/slots |
| e002646 | 19:58:57.166798 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/4/nodes/5/link_state |
| e002647 | 19:58:57.166798 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/4/persistence/aof_enabled |
| e002648 | 19:58:57.166798 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/4/persistence/loading |
| e002649 | 19:58:57.166798 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/4/persistence/aof_last_write_status |
| e002650 | 19:58:57.166798 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e002651 | 19:58:57.166798 | node-5 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/4/process_live |
| e002652 | 19:58:57.166802 | node-4 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/3/reachable |
| e002653 | 19:58:57.166802 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/10/snapshot/views/3/replication/role |
| e002654 | 19:58:57.166802 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/10/snapshot/views/3/replication/connected_slaves |
| e002655 | 19:58:57.166802 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/10/snapshot/views/3/replication/master_link_status |
| e002656 | 19:58:57.166802 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/10/snapshot/views/3/replication/master_sync_in_progress |
| e002657 | 19:58:57.166802 | node-4 | info_replication | replication.master_repl_offset | valid | `239` | /samples/10/snapshot/views/3/replication/master_repl_offset |
| e002658 | 19:58:57.166802 | node-4 | info_replication | replication.slave_repl_offset | valid | `239` | /samples/10/snapshot/views/3/replication/slave_repl_offset |
| e002659 | 19:58:57.166802 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/10/snapshot/views/3/redis_id |
| e002660 | 19:58:57.166802 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/3/coverage |
| e002661 | 19:58:57.166802 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/3/nodes/0/flags |
| e002662 | 19:58:57.166802 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/0/primary_id |
| e002663 | 19:58:57.166802 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/3/nodes/0/slots |
| e002664 | 19:58:57.166802 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/0/link_state |
| e002665 | 19:58:57.166802 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/10/snapshot/views/3/nodes/1/flags |
| e002666 | 19:58:57.166802 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/10/snapshot/views/3/nodes/1/primary_id |
| e002667 | 19:58:57.166802 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/1/slots |
| e002668 | 19:58:57.166802 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/1/link_state |
| e002669 | 19:58:57.166802 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/3/nodes/2/flags |
| e002670 | 19:58:57.166802 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/10/snapshot/views/3/nodes/2/primary_id |
| e002671 | 19:58:57.166802 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/2/slots |
| e002672 | 19:58:57.166802 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/2/link_state |
| e002673 | 19:58:57.166802 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/3/nodes/3/flags |
| e002674 | 19:58:57.166802 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/3/primary_id |
| e002675 | 19:58:57.166802 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/3/nodes/3/slots |
| e002676 | 19:58:57.166802 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/3/link_state |
| e002677 | 19:58:57.166802 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/3/nodes/4/flags |
| e002678 | 19:58:57.166802 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/10/snapshot/views/3/nodes/4/primary_id |
| e002679 | 19:58:57.166802 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/4/slots |
| e002680 | 19:58:57.166802 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/4/link_state |
| e002681 | 19:58:57.166802 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/3/nodes/5/flags |
| e002682 | 19:58:57.166802 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/3/nodes/5/primary_id |
| e002683 | 19:58:57.166802 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/3/nodes/5/slots |
| e002684 | 19:58:57.166802 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/3/nodes/5/link_state |
| e002685 | 19:58:57.166802 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/3/persistence/aof_enabled |
| e002686 | 19:58:57.166802 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/3/persistence/loading |
| e002687 | 19:58:57.166802 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/3/persistence/aof_last_write_status |
| e002688 | 19:58:57.166802 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e002689 | 19:58:57.166802 | node-4 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/3/process_live |
| e002690 | 19:58:57.166804 | node-3 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/2/reachable |
| e002691 | 19:58:57.166804 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/10/snapshot/views/2/replication/role |
| e002692 | 19:58:57.166804 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/10/snapshot/views/2/replication/connected_slaves |
| e002693 | 19:58:57.166804 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/10/snapshot/views/2/replication/master_link_status |
| e002694 | 19:58:57.166804 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/10/snapshot/views/2/replication/master_sync_in_progress |
| e002695 | 19:58:57.166804 | node-3 | info_replication | replication.master_repl_offset | valid | `239` | /samples/10/snapshot/views/2/replication/master_repl_offset |
| e002696 | 19:58:57.166804 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/10/snapshot/views/2/replication/slave_repl_offset |
| e002697 | 19:58:57.166804 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/10/snapshot/views/2/redis_id |
| e002698 | 19:58:57.166804 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/2/coverage |
| e002699 | 19:58:57.166804 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/2/nodes/0/flags |
| e002700 | 19:58:57.166804 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/0/primary_id |
| e002701 | 19:58:57.166804 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/2/nodes/0/slots |
| e002702 | 19:58:57.166804 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/0/link_state |
| e002703 | 19:58:57.166804 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/10/snapshot/views/2/nodes/1/flags |
| e002704 | 19:58:57.166804 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/1/primary_id |
| e002705 | 19:58:57.166804 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/2/nodes/1/slots |
| e002706 | 19:58:57.166804 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/1/link_state |
| e002707 | 19:58:57.166804 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/2/nodes/2/flags |
| e002708 | 19:58:57.166804 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/10/snapshot/views/2/nodes/2/primary_id |
| e002709 | 19:58:57.166804 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/2/slots |
| e002710 | 19:58:57.166804 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/2/link_state |
| e002711 | 19:58:57.166804 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/2/nodes/3/flags |
| e002712 | 19:58:57.166804 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/10/snapshot/views/2/nodes/3/primary_id |
| e002713 | 19:58:57.166804 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/3/slots |
| e002714 | 19:58:57.166804 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/3/link_state |
| e002715 | 19:58:57.166804 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/2/nodes/4/flags |
| e002716 | 19:58:57.166804 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/10/snapshot/views/2/nodes/4/primary_id |
| e002717 | 19:58:57.166804 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/4/slots |
| e002718 | 19:58:57.166804 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/4/link_state |
| e002719 | 19:58:57.166804 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/2/nodes/5/flags |
| e002720 | 19:58:57.166804 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/2/nodes/5/primary_id |
| e002721 | 19:58:57.166804 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/2/nodes/5/slots |
| e002722 | 19:58:57.166804 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/2/nodes/5/link_state |
| e002723 | 19:58:57.166804 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/2/persistence/aof_enabled |
| e002724 | 19:58:57.166804 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/2/persistence/loading |
| e002725 | 19:58:57.166804 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/2/persistence/aof_last_write_status |
| e002726 | 19:58:57.166804 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e002727 | 19:58:57.166804 | node-3 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/2/process_live |
| e002728 | 19:58:57.166852 | node-2 | node_client | reachable | valid | `true` | /samples/10/snapshot/views/1/reachable |
| e002729 | 19:58:57.166852 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/10/snapshot/views/1/replication/role |
| e002730 | 19:58:57.166852 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/10/snapshot/views/1/replication/connected_slaves |
| e002731 | 19:58:57.166852 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/10/snapshot/views/1/replication/master_link_status |
| e002732 | 19:58:57.166852 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/10/snapshot/views/1/replication/master_sync_in_progress |
| e002733 | 19:58:57.166852 | node-2 | info_replication | replication.master_repl_offset | valid | `239` | /samples/10/snapshot/views/1/replication/master_repl_offset |
| e002734 | 19:58:57.166852 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/10/snapshot/views/1/replication/slave_repl_offset |
| e002735 | 19:58:57.166852 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/10/snapshot/views/1/redis_id |
| e002736 | 19:58:57.166852 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/10/snapshot/views/1/coverage |
| e002737 | 19:58:57.166852 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/1/nodes/0/flags |
| e002738 | 19:58:57.166852 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/10/snapshot/views/1/nodes/0/primary_id |
| e002739 | 19:58:57.166852 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/0/slots |
| e002740 | 19:58:57.166852 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/0/link_state |
| e002741 | 19:58:57.166852 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/1/nodes/1/flags |
| e002742 | 19:58:57.166852 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/1/primary_id |
| e002743 | 19:58:57.166852 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/10/snapshot/views/1/nodes/1/slots |
| e002744 | 19:58:57.166852 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/1/link_state |
| e002745 | 19:58:57.166852 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/1/nodes/2/flags |
| e002746 | 19:58:57.166852 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/10/snapshot/views/1/nodes/2/primary_id |
| e002747 | 19:58:57.166852 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/2/slots |
| e002748 | 19:58:57.166852 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/2/link_state |
| e002749 | 19:58:57.166852 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/10/snapshot/views/1/nodes/3/flags |
| e002750 | 19:58:57.166852 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/3/primary_id |
| e002751 | 19:58:57.166852 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/10/snapshot/views/1/nodes/3/slots |
| e002752 | 19:58:57.166852 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/3/link_state |
| e002753 | 19:58:57.166852 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/10/snapshot/views/1/nodes/4/flags |
| e002754 | 19:58:57.166852 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/10/snapshot/views/1/nodes/4/primary_id |
| e002755 | 19:58:57.166852 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/4/slots |
| e002756 | 19:58:57.166852 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/4/link_state |
| e002757 | 19:58:57.166852 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/10/snapshot/views/1/nodes/5/flags |
| e002758 | 19:58:57.166852 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/10/snapshot/views/1/nodes/5/primary_id |
| e002759 | 19:58:57.166852 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/10/snapshot/views/1/nodes/5/slots |
| e002760 | 19:58:57.166852 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/10/snapshot/views/1/nodes/5/link_state |
| e002761 | 19:58:57.166852 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/10/snapshot/views/1/persistence/aof_enabled |
| e002762 | 19:58:57.166852 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/10/snapshot/views/1/persistence/loading |
| e002763 | 19:58:57.166852 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/10/snapshot/views/1/persistence/aof_last_write_status |
| e002764 | 19:58:57.166852 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/10/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e002765 | 19:58:57.166852 | node-2 | admin | process_live | missing_from_schema | `` | /samples/10/snapshot/views/1/process_live |
| e002766 | 19:58:57.173638 |  | client_probe | client_success | valid | `true` | /samples/10/probe/successful |
| e002767 | 19:58:57.760237 | node-3 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/2/reachable |
| e002768 | 19:58:57.760237 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/11/snapshot/views/2/replication/role |
| e002769 | 19:58:57.760237 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/11/snapshot/views/2/replication/connected_slaves |
| e002770 | 19:58:57.760237 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/11/snapshot/views/2/replication/master_link_status |
| e002771 | 19:58:57.760237 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/11/snapshot/views/2/replication/master_sync_in_progress |
| e002772 | 19:58:57.760237 | node-3 | info_replication | replication.master_repl_offset | valid | `340` | /samples/11/snapshot/views/2/replication/master_repl_offset |
| e002773 | 19:58:57.760237 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/11/snapshot/views/2/replication/slave_repl_offset |
| e002774 | 19:58:57.760237 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/11/snapshot/views/2/redis_id |
| e002775 | 19:58:57.760237 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/2/coverage |
| e002776 | 19:58:57.760237 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/2/nodes/0/flags |
| e002777 | 19:58:57.760237 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/0/primary_id |
| e002778 | 19:58:57.760237 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/2/nodes/0/slots |
| e002779 | 19:58:57.760237 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/0/link_state |
| e002780 | 19:58:57.760237 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/11/snapshot/views/2/nodes/1/flags |
| e002781 | 19:58:57.760237 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/1/primary_id |
| e002782 | 19:58:57.760237 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/2/nodes/1/slots |
| e002783 | 19:58:57.760237 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/1/link_state |
| e002784 | 19:58:57.760237 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/2/nodes/2/flags |
| e002785 | 19:58:57.760237 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/11/snapshot/views/2/nodes/2/primary_id |
| e002786 | 19:58:57.760237 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/2/slots |
| e002787 | 19:58:57.760237 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/2/link_state |
| e002788 | 19:58:57.760237 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/2/nodes/3/flags |
| e002789 | 19:58:57.760237 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/11/snapshot/views/2/nodes/3/primary_id |
| e002790 | 19:58:57.760237 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/3/slots |
| e002791 | 19:58:57.760237 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/3/link_state |
| e002792 | 19:58:57.760237 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/2/nodes/4/flags |
| e002793 | 19:58:57.760237 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/11/snapshot/views/2/nodes/4/primary_id |
| e002794 | 19:58:57.760237 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/4/slots |
| e002795 | 19:58:57.760237 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/4/link_state |
| e002796 | 19:58:57.760237 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/2/nodes/5/flags |
| e002797 | 19:58:57.760237 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/2/nodes/5/primary_id |
| e002798 | 19:58:57.760237 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/2/nodes/5/slots |
| e002799 | 19:58:57.760237 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/2/nodes/5/link_state |
| e002800 | 19:58:57.760237 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/2/persistence/aof_enabled |
| e002801 | 19:58:57.760237 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/2/persistence/loading |
| e002802 | 19:58:57.760237 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/2/persistence/aof_last_write_status |
| e002803 | 19:58:57.760237 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e002804 | 19:58:57.760237 | node-3 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/2/process_live |
| e002805 | 19:58:57.76024 | node-1 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/0/reachable |
| e002806 | 19:58:57.76024 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/11/snapshot/views/0/replication/role |
| e002807 | 19:58:57.76024 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/11/snapshot/views/0/replication/connected_slaves |
| e002808 | 19:58:57.76024 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/11/snapshot/views/0/replication/master_link_status |
| e002809 | 19:58:57.76024 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/11/snapshot/views/0/replication/master_sync_in_progress |
| e002810 | 19:58:57.76024 | node-1 | info_replication | replication.master_repl_offset | valid | `354` | /samples/11/snapshot/views/0/replication/master_repl_offset |
| e002811 | 19:58:57.76024 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/11/snapshot/views/0/replication/slave_repl_offset |
| e002812 | 19:58:57.76024 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/11/snapshot/views/0/redis_id |
| e002813 | 19:58:57.76024 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/0/coverage |
| e002814 | 19:58:57.76024 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/0/nodes/0/flags |
| e002815 | 19:58:57.76024 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/11/snapshot/views/0/nodes/0/primary_id |
| e002816 | 19:58:57.76024 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/0/slots |
| e002817 | 19:58:57.76024 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/0/link_state |
| e002818 | 19:58:57.76024 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/0/nodes/1/flags |
| e002819 | 19:58:57.76024 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/1/primary_id |
| e002820 | 19:58:57.76024 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/0/nodes/1/slots |
| e002821 | 19:58:57.76024 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/1/link_state |
| e002822 | 19:58:57.76024 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/0/nodes/2/flags |
| e002823 | 19:58:57.76024 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/2/primary_id |
| e002824 | 19:58:57.76024 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/0/nodes/2/slots |
| e002825 | 19:58:57.76024 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/2/link_state |
| e002826 | 19:58:57.76024 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/11/snapshot/views/0/nodes/3/flags |
| e002827 | 19:58:57.76024 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/3/primary_id |
| e002828 | 19:58:57.76024 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/0/nodes/3/slots |
| e002829 | 19:58:57.76024 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/3/link_state |
| e002830 | 19:58:57.76024 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/0/nodes/4/flags |
| e002831 | 19:58:57.76024 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/11/snapshot/views/0/nodes/4/primary_id |
| e002832 | 19:58:57.76024 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/4/slots |
| e002833 | 19:58:57.76024 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/4/link_state |
| e002834 | 19:58:57.76024 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/0/nodes/5/flags |
| e002835 | 19:58:57.76024 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/11/snapshot/views/0/nodes/5/primary_id |
| e002836 | 19:58:57.76024 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/0/nodes/5/slots |
| e002837 | 19:58:57.76024 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/0/nodes/5/link_state |
| e002838 | 19:58:57.76024 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/0/persistence/aof_enabled |
| e002839 | 19:58:57.76024 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/0/persistence/loading |
| e002840 | 19:58:57.76024 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/0/persistence/aof_last_write_status |
| e002841 | 19:58:57.76024 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e002842 | 19:58:57.76024 | node-1 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/0/process_live |
| e002843 | 19:58:57.760241 | node-2 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/1/reachable |
| e002844 | 19:58:57.760241 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/11/snapshot/views/1/replication/role |
| e002845 | 19:58:57.760241 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/11/snapshot/views/1/replication/connected_slaves |
| e002846 | 19:58:57.760241 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/11/snapshot/views/1/replication/master_link_status |
| e002847 | 19:58:57.760241 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/11/snapshot/views/1/replication/master_sync_in_progress |
| e002848 | 19:58:57.760241 | node-2 | info_replication | replication.master_repl_offset | valid | `340` | /samples/11/snapshot/views/1/replication/master_repl_offset |
| e002849 | 19:58:57.760241 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/11/snapshot/views/1/replication/slave_repl_offset |
| e002850 | 19:58:57.760241 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/11/snapshot/views/1/redis_id |
| e002851 | 19:58:57.760241 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/1/coverage |
| e002852 | 19:58:57.760241 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/1/nodes/0/flags |
| e002853 | 19:58:57.760241 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/11/snapshot/views/1/nodes/0/primary_id |
| e002854 | 19:58:57.760241 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/0/slots |
| e002855 | 19:58:57.760241 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/0/link_state |
| e002856 | 19:58:57.760241 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/1/nodes/1/flags |
| e002857 | 19:58:57.760241 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/1/primary_id |
| e002858 | 19:58:57.760241 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/1/nodes/1/slots |
| e002859 | 19:58:57.760241 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/1/link_state |
| e002860 | 19:58:57.760241 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/1/nodes/2/flags |
| e002861 | 19:58:57.760241 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/11/snapshot/views/1/nodes/2/primary_id |
| e002862 | 19:58:57.760241 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/2/slots |
| e002863 | 19:58:57.760241 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/2/link_state |
| e002864 | 19:58:57.760241 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/1/nodes/3/flags |
| e002865 | 19:58:57.760241 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/3/primary_id |
| e002866 | 19:58:57.760241 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/1/nodes/3/slots |
| e002867 | 19:58:57.760241 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/3/link_state |
| e002868 | 19:58:57.760241 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/1/nodes/4/flags |
| e002869 | 19:58:57.760241 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/11/snapshot/views/1/nodes/4/primary_id |
| e002870 | 19:58:57.760241 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/4/slots |
| e002871 | 19:58:57.760241 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/4/link_state |
| e002872 | 19:58:57.760241 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/11/snapshot/views/1/nodes/5/flags |
| e002873 | 19:58:57.760241 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/1/nodes/5/primary_id |
| e002874 | 19:58:57.760241 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/1/nodes/5/slots |
| e002875 | 19:58:57.760241 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/1/nodes/5/link_state |
| e002876 | 19:58:57.760241 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/1/persistence/aof_enabled |
| e002877 | 19:58:57.760241 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/1/persistence/loading |
| e002878 | 19:58:57.760241 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/1/persistence/aof_last_write_status |
| e002879 | 19:58:57.760241 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e002880 | 19:58:57.760241 | node-2 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/1/process_live |
| e002881 | 19:58:57.760247 | node-5 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/4/reachable |
| e002882 | 19:58:57.760247 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/11/snapshot/views/4/replication/role |
| e002883 | 19:58:57.760247 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/11/snapshot/views/4/replication/connected_slaves |
| e002884 | 19:58:57.760247 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/11/snapshot/views/4/replication/master_link_status |
| e002885 | 19:58:57.760247 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/11/snapshot/views/4/replication/master_sync_in_progress |
| e002886 | 19:58:57.760247 | node-5 | info_replication | replication.master_repl_offset | valid | `354` | /samples/11/snapshot/views/4/replication/master_repl_offset |
| e002887 | 19:58:57.760247 | node-5 | info_replication | replication.slave_repl_offset | valid | `354` | /samples/11/snapshot/views/4/replication/slave_repl_offset |
| e002888 | 19:58:57.760247 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/11/snapshot/views/4/redis_id |
| e002889 | 19:58:57.760247 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/4/coverage |
| e002890 | 19:58:57.760247 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/4/nodes/0/flags |
| e002891 | 19:58:57.760247 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/0/primary_id |
| e002892 | 19:58:57.760247 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/4/nodes/0/slots |
| e002893 | 19:58:57.760247 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/0/link_state |
| e002894 | 19:58:57.760247 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/4/nodes/1/flags |
| e002895 | 19:58:57.760247 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/11/snapshot/views/4/nodes/1/primary_id |
| e002896 | 19:58:57.760247 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/1/slots |
| e002897 | 19:58:57.760247 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/1/link_state |
| e002898 | 19:58:57.760247 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/4/nodes/2/flags |
| e002899 | 19:58:57.760247 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/2/primary_id |
| e002900 | 19:58:57.760247 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/4/nodes/2/slots |
| e002901 | 19:58:57.760247 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/2/link_state |
| e002902 | 19:58:57.760247 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/4/nodes/3/flags |
| e002903 | 19:58:57.760247 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/3/primary_id |
| e002904 | 19:58:57.760247 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/4/nodes/3/slots |
| e002905 | 19:58:57.760247 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/3/link_state |
| e002906 | 19:58:57.760247 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/11/snapshot/views/4/nodes/4/flags |
| e002907 | 19:58:57.760247 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/11/snapshot/views/4/nodes/4/primary_id |
| e002908 | 19:58:57.760247 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/4/slots |
| e002909 | 19:58:57.760247 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/4/link_state |
| e002910 | 19:58:57.760247 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/4/nodes/5/flags |
| e002911 | 19:58:57.760247 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/11/snapshot/views/4/nodes/5/primary_id |
| e002912 | 19:58:57.760247 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/4/nodes/5/slots |
| e002913 | 19:58:57.760247 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/4/nodes/5/link_state |
| e002914 | 19:58:57.760247 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/4/persistence/aof_enabled |
| e002915 | 19:58:57.760247 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/4/persistence/loading |
| e002916 | 19:58:57.760247 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/4/persistence/aof_last_write_status |
| e002917 | 19:58:57.760247 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e002918 | 19:58:57.760247 | node-5 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/4/process_live |
| e002919 | 19:58:57.760339 | node-6 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/5/reachable |
| e002920 | 19:58:57.760339 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/11/snapshot/views/5/replication/role |
| e002921 | 19:58:57.760339 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/11/snapshot/views/5/replication/connected_slaves |
| e002922 | 19:58:57.760339 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/11/snapshot/views/5/replication/master_link_status |
| e002923 | 19:58:57.760339 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/11/snapshot/views/5/replication/master_sync_in_progress |
| e002924 | 19:58:57.760339 | node-6 | info_replication | replication.master_repl_offset | valid | `340` | /samples/11/snapshot/views/5/replication/master_repl_offset |
| e002925 | 19:58:57.760339 | node-6 | info_replication | replication.slave_repl_offset | valid | `340` | /samples/11/snapshot/views/5/replication/slave_repl_offset |
| e002926 | 19:58:57.760339 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/11/snapshot/views/5/redis_id |
| e002927 | 19:58:57.760339 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/5/coverage |
| e002928 | 19:58:57.760339 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/5/nodes/0/flags |
| e002929 | 19:58:57.760339 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/11/snapshot/views/5/nodes/0/primary_id |
| e002930 | 19:58:57.760339 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/0/slots |
| e002931 | 19:58:57.760339 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/0/link_state |
| e002932 | 19:58:57.760339 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/5/nodes/1/flags |
| e002933 | 19:58:57.760339 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/1/primary_id |
| e002934 | 19:58:57.760339 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/5/nodes/1/slots |
| e002935 | 19:58:57.760339 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/1/link_state |
| e002936 | 19:58:57.760339 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/5/nodes/2/flags |
| e002937 | 19:58:57.760339 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/2/primary_id |
| e002938 | 19:58:57.760339 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/5/nodes/2/slots |
| e002939 | 19:58:57.760339 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/2/link_state |
| e002940 | 19:58:57.760339 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/5/nodes/3/flags |
| e002941 | 19:58:57.760339 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/3/primary_id |
| e002942 | 19:58:57.760339 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/5/nodes/3/slots |
| e002943 | 19:58:57.760339 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/3/link_state |
| e002944 | 19:58:57.760339 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/5/nodes/4/flags |
| e002945 | 19:58:57.760339 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/11/snapshot/views/5/nodes/4/primary_id |
| e002946 | 19:58:57.760339 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/4/slots |
| e002947 | 19:58:57.760339 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/4/link_state |
| e002948 | 19:58:57.760339 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/11/snapshot/views/5/nodes/5/flags |
| e002949 | 19:58:57.760339 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/11/snapshot/views/5/nodes/5/primary_id |
| e002950 | 19:58:57.760339 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/5/nodes/5/slots |
| e002951 | 19:58:57.760339 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/5/nodes/5/link_state |
| e002952 | 19:58:57.760339 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/5/persistence/aof_enabled |
| e002953 | 19:58:57.760339 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/5/persistence/loading |
| e002954 | 19:58:57.760339 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/5/persistence/aof_last_write_status |
| e002955 | 19:58:57.760339 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e002956 | 19:58:57.760339 | node-6 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/5/process_live |
| e002957 | 19:58:57.760358 | node-4 | node_client | reachable | valid | `true` | /samples/11/snapshot/views/3/reachable |
| e002958 | 19:58:57.760358 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/11/snapshot/views/3/replication/role |
| e002959 | 19:58:57.760358 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/11/snapshot/views/3/replication/connected_slaves |
| e002960 | 19:58:57.760358 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/11/snapshot/views/3/replication/master_link_status |
| e002961 | 19:58:57.760358 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/11/snapshot/views/3/replication/master_sync_in_progress |
| e002962 | 19:58:57.760358 | node-4 | info_replication | replication.master_repl_offset | valid | `340` | /samples/11/snapshot/views/3/replication/master_repl_offset |
| e002963 | 19:58:57.760358 | node-4 | info_replication | replication.slave_repl_offset | valid | `340` | /samples/11/snapshot/views/3/replication/slave_repl_offset |
| e002964 | 19:58:57.760358 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/11/snapshot/views/3/redis_id |
| e002965 | 19:58:57.760358 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/11/snapshot/views/3/coverage |
| e002966 | 19:58:57.760358 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/3/nodes/0/flags |
| e002967 | 19:58:57.760358 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/0/primary_id |
| e002968 | 19:58:57.760358 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/11/snapshot/views/3/nodes/0/slots |
| e002969 | 19:58:57.760358 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/0/link_state |
| e002970 | 19:58:57.760358 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/11/snapshot/views/3/nodes/1/flags |
| e002971 | 19:58:57.760358 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/11/snapshot/views/3/nodes/1/primary_id |
| e002972 | 19:58:57.760358 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/1/slots |
| e002973 | 19:58:57.760358 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/1/link_state |
| e002974 | 19:58:57.760358 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/3/nodes/2/flags |
| e002975 | 19:58:57.760358 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/11/snapshot/views/3/nodes/2/primary_id |
| e002976 | 19:58:57.760358 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/2/slots |
| e002977 | 19:58:57.760358 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/2/link_state |
| e002978 | 19:58:57.760358 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/3/nodes/3/flags |
| e002979 | 19:58:57.760358 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/3/primary_id |
| e002980 | 19:58:57.760358 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/11/snapshot/views/3/nodes/3/slots |
| e002981 | 19:58:57.760358 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/3/link_state |
| e002982 | 19:58:57.760358 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/11/snapshot/views/3/nodes/4/flags |
| e002983 | 19:58:57.760358 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/11/snapshot/views/3/nodes/4/primary_id |
| e002984 | 19:58:57.760358 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/4/slots |
| e002985 | 19:58:57.760358 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/4/link_state |
| e002986 | 19:58:57.760358 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/11/snapshot/views/3/nodes/5/flags |
| e002987 | 19:58:57.760358 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/11/snapshot/views/3/nodes/5/primary_id |
| e002988 | 19:58:57.760358 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/11/snapshot/views/3/nodes/5/slots |
| e002989 | 19:58:57.760358 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/11/snapshot/views/3/nodes/5/link_state |
| e002990 | 19:58:57.760358 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/11/snapshot/views/3/persistence/aof_enabled |
| e002991 | 19:58:57.760358 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/11/snapshot/views/3/persistence/loading |
| e002992 | 19:58:57.760358 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/11/snapshot/views/3/persistence/aof_last_write_status |
| e002993 | 19:58:57.760358 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/11/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e002994 | 19:58:57.760358 | node-4 | admin | process_live | missing_from_schema | `` | /samples/11/snapshot/views/3/process_live |
| e002995 | 19:58:57.766931 |  | client_probe | client_success | valid | `true` | /samples/11/probe/successful |
| e002996 | 19:58:58.350984 | node-2 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/1/reachable |
| e002997 | 19:58:58.350984 | node-2 | info_replication | replication.role | valid | `"master"` | /samples/12/snapshot/views/1/replication/role |
| e002998 | 19:58:58.350984 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /samples/12/snapshot/views/1/replication/connected_slaves |
| e002999 | 19:58:58.350984 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /samples/12/snapshot/views/1/replication/master_link_status |
| e003000 | 19:58:58.350984 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/12/snapshot/views/1/replication/master_sync_in_progress |
| e003001 | 19:58:58.350984 | node-2 | info_replication | replication.master_repl_offset | valid | `441` | /samples/12/snapshot/views/1/replication/master_repl_offset |
| e003002 | 19:58:58.350984 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/12/snapshot/views/1/replication/slave_repl_offset |
| e003003 | 19:58:58.350984 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/12/snapshot/views/1/redis_id |
| e003004 | 19:58:58.350984 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/1/coverage |
| e003005 | 19:58:58.350984 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/1/nodes/0/flags |
| e003006 | 19:58:58.350984 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/12/snapshot/views/1/nodes/0/primary_id |
| e003007 | 19:58:58.350984 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/0/slots |
| e003008 | 19:58:58.350984 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/0/link_state |
| e003009 | 19:58:58.350984 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/1/nodes/1/flags |
| e003010 | 19:58:58.350984 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/1/primary_id |
| e003011 | 19:58:58.350984 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/1/nodes/1/slots |
| e003012 | 19:58:58.350984 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/1/link_state |
| e003013 | 19:58:58.350984 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/1/nodes/2/flags |
| e003014 | 19:58:58.350984 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/12/snapshot/views/1/nodes/2/primary_id |
| e003015 | 19:58:58.350984 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/2/slots |
| e003016 | 19:58:58.350984 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/2/link_state |
| e003017 | 19:58:58.350984 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/1/nodes/3/flags |
| e003018 | 19:58:58.350984 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/3/primary_id |
| e003019 | 19:58:58.350984 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/1/nodes/3/slots |
| e003020 | 19:58:58.350984 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/3/link_state |
| e003021 | 19:58:58.350984 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/1/nodes/4/flags |
| e003022 | 19:58:58.350984 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/12/snapshot/views/1/nodes/4/primary_id |
| e003023 | 19:58:58.350984 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/4/slots |
| e003024 | 19:58:58.350984 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/4/link_state |
| e003025 | 19:58:58.350984 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/12/snapshot/views/1/nodes/5/flags |
| e003026 | 19:58:58.350984 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/1/nodes/5/primary_id |
| e003027 | 19:58:58.350984 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/1/nodes/5/slots |
| e003028 | 19:58:58.350984 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/1/nodes/5/link_state |
| e003029 | 19:58:58.350984 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/1/persistence/aof_enabled |
| e003030 | 19:58:58.350984 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/1/persistence/loading |
| e003031 | 19:58:58.350984 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/1/persistence/aof_last_write_status |
| e003032 | 19:58:58.350984 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/1/persistence/aof_rewrite_in_progress |
| e003033 | 19:58:58.350984 | node-2 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/1/process_live |
| e003034 | 19:58:58.350984 | node-2 | node_client | reachable | valid | `true` | /final/views/1/reachable |
| e003035 | 19:58:58.350984 | node-2 | info_replication | replication.role | valid | `"master"` | /final/views/1/replication/role |
| e003036 | 19:58:58.350984 | node-2 | info_replication | replication.connected_slaves | valid | `1` | /final/views/1/replication/connected_slaves |
| e003037 | 19:58:58.350984 | node-2 | info_replication | replication.master_link_status | not_applicable | `` | /final/views/1/replication/master_link_status |
| e003038 | 19:58:58.350984 | node-2 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /final/views/1/replication/master_sync_in_progress |
| e003039 | 19:58:58.350984 | node-2 | info_replication | replication.master_repl_offset | valid | `441` | /final/views/1/replication/master_repl_offset |
| e003040 | 19:58:58.350984 | node-2 | info_replication | replication.slave_repl_offset | not_applicable | `` | /final/views/1/replication/slave_repl_offset |
| e003041 | 19:58:58.350984 | node-2 | cluster_nodes | redis_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /final/views/1/redis_id |
| e003042 | 19:58:58.350984 | node-2 | cluster_nodes | slot_coverage | valid | `true` | /final/views/1/coverage |
| e003043 | 19:58:58.350984 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/1/nodes/0/flags |
| e003044 | 19:58:58.350984 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /final/views/1/nodes/0/primary_id |
| e003045 | 19:58:58.350984 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.slots | missing_from_schema | `` | /final/views/1/nodes/0/slots |
| e003046 | 19:58:58.350984 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/0/link_state |
| e003047 | 19:58:58.350984 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /final/views/1/nodes/1/flags |
| e003048 | 19:58:58.350984 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /final/views/1/nodes/1/primary_id |
| e003049 | 19:58:58.350984 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/1/nodes/1/slots |
| e003050 | 19:58:58.350984 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/1/link_state |
| e003051 | 19:58:58.350984 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/1/nodes/2/flags |
| e003052 | 19:58:58.350984 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /final/views/1/nodes/2/primary_id |
| e003053 | 19:58:58.350984 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.slots | missing_from_schema | `` | /final/views/1/nodes/2/slots |
| e003054 | 19:58:58.350984 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/2/link_state |
| e003055 | 19:58:58.350984 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.flags | valid | `[
              "master"
            ]` | /final/views/1/nodes/3/flags |
| e003056 | 19:58:58.350984 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /final/views/1/nodes/3/primary_id |
| e003057 | 19:58:58.350984 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/1/nodes/3/slots |
| e003058 | 19:58:58.350984 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/3/link_state |
| e003059 | 19:58:58.350984 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/1/nodes/4/flags |
| e003060 | 19:58:58.350984 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /final/views/1/nodes/4/primary_id |
| e003061 | 19:58:58.350984 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.slots | missing_from_schema | `` | /final/views/1/nodes/4/slots |
| e003062 | 19:58:58.350984 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/4/link_state |
| e003063 | 19:58:58.350984 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /final/views/1/nodes/5/flags |
| e003064 | 19:58:58.350984 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.primary_id | missing_from_schema | `` | /final/views/1/nodes/5/primary_id |
| e003065 | 19:58:58.350984 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/1/nodes/5/slots |
| e003066 | 19:58:58.350984 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-2 | peer.link_state | valid | `"connected"` | /final/views/1/nodes/5/link_state |
| e003067 | 19:58:58.350984 | node-2 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/1/persistence/aof_enabled |
| e003068 | 19:58:58.350984 | node-2 | info_persistence | persistence.loading | valid | `"0"` | /final/views/1/persistence/loading |
| e003069 | 19:58:58.350984 | node-2 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/1/persistence/aof_last_write_status |
| e003070 | 19:58:58.350984 | node-2 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/1/persistence/aof_rewrite_in_progress |
| e003071 | 19:58:58.350984 | node-2 | admin | process_live | missing_from_schema | `` | /final/views/1/process_live |
| e003072 | 19:58:58.350991 | node-1 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/0/reachable |
| e003073 | 19:58:58.350991 | node-1 | info_replication | replication.role | valid | `"master"` | /samples/12/snapshot/views/0/replication/role |
| e003074 | 19:58:58.350991 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /samples/12/snapshot/views/0/replication/connected_slaves |
| e003075 | 19:58:58.350991 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /samples/12/snapshot/views/0/replication/master_link_status |
| e003076 | 19:58:58.350991 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/12/snapshot/views/0/replication/master_sync_in_progress |
| e003077 | 19:58:58.350991 | node-1 | info_replication | replication.master_repl_offset | valid | `455` | /samples/12/snapshot/views/0/replication/master_repl_offset |
| e003078 | 19:58:58.350991 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/12/snapshot/views/0/replication/slave_repl_offset |
| e003079 | 19:58:58.350991 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/12/snapshot/views/0/redis_id |
| e003080 | 19:58:58.350991 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/0/coverage |
| e003081 | 19:58:58.350991 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/0/nodes/0/flags |
| e003082 | 19:58:58.350991 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/12/snapshot/views/0/nodes/0/primary_id |
| e003083 | 19:58:58.350991 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/0/slots |
| e003084 | 19:58:58.350991 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/0/link_state |
| e003085 | 19:58:58.350991 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/0/nodes/1/flags |
| e003086 | 19:58:58.350991 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/1/primary_id |
| e003087 | 19:58:58.350991 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/0/nodes/1/slots |
| e003088 | 19:58:58.350991 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/1/link_state |
| e003089 | 19:58:58.350991 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/0/nodes/2/flags |
| e003090 | 19:58:58.350991 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/2/primary_id |
| e003091 | 19:58:58.350991 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/0/nodes/2/slots |
| e003092 | 19:58:58.350991 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/2/link_state |
| e003093 | 19:58:58.350991 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/12/snapshot/views/0/nodes/3/flags |
| e003094 | 19:58:58.350991 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/3/primary_id |
| e003095 | 19:58:58.350991 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/0/nodes/3/slots |
| e003096 | 19:58:58.350991 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/3/link_state |
| e003097 | 19:58:58.350991 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/0/nodes/4/flags |
| e003098 | 19:58:58.350991 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/12/snapshot/views/0/nodes/4/primary_id |
| e003099 | 19:58:58.350991 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/4/slots |
| e003100 | 19:58:58.350991 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/4/link_state |
| e003101 | 19:58:58.350991 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/0/nodes/5/flags |
| e003102 | 19:58:58.350991 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/12/snapshot/views/0/nodes/5/primary_id |
| e003103 | 19:58:58.350991 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/0/nodes/5/slots |
| e003104 | 19:58:58.350991 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/0/nodes/5/link_state |
| e003105 | 19:58:58.350991 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/0/persistence/aof_enabled |
| e003106 | 19:58:58.350991 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/0/persistence/loading |
| e003107 | 19:58:58.350991 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/0/persistence/aof_last_write_status |
| e003108 | 19:58:58.350991 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/0/persistence/aof_rewrite_in_progress |
| e003109 | 19:58:58.350991 | node-1 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/0/process_live |
| e003110 | 19:58:58.350991 | node-1 | node_client | reachable | valid | `true` | /final/views/0/reachable |
| e003111 | 19:58:58.350991 | node-1 | info_replication | replication.role | valid | `"master"` | /final/views/0/replication/role |
| e003112 | 19:58:58.350991 | node-1 | info_replication | replication.connected_slaves | valid | `1` | /final/views/0/replication/connected_slaves |
| e003113 | 19:58:58.350991 | node-1 | info_replication | replication.master_link_status | not_applicable | `` | /final/views/0/replication/master_link_status |
| e003114 | 19:58:58.350991 | node-1 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /final/views/0/replication/master_sync_in_progress |
| e003115 | 19:58:58.350991 | node-1 | info_replication | replication.master_repl_offset | valid | `455` | /final/views/0/replication/master_repl_offset |
| e003116 | 19:58:58.350991 | node-1 | info_replication | replication.slave_repl_offset | not_applicable | `` | /final/views/0/replication/slave_repl_offset |
| e003117 | 19:58:58.350991 | node-1 | cluster_nodes | redis_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /final/views/0/redis_id |
| e003118 | 19:58:58.350991 | node-1 | cluster_nodes | slot_coverage | valid | `true` | /final/views/0/coverage |
| e003119 | 19:58:58.350991 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/0/nodes/0/flags |
| e003120 | 19:58:58.350991 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /final/views/0/nodes/0/primary_id |
| e003121 | 19:58:58.350991 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.slots | missing_from_schema | `` | /final/views/0/nodes/0/slots |
| e003122 | 19:58:58.350991 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/0/link_state |
| e003123 | 19:58:58.350991 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.flags | valid | `[
              "master"
            ]` | /final/views/0/nodes/1/flags |
| e003124 | 19:58:58.350991 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /final/views/0/nodes/1/primary_id |
| e003125 | 19:58:58.350991 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/0/nodes/1/slots |
| e003126 | 19:58:58.350991 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/1/link_state |
| e003127 | 19:58:58.350991 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.flags | valid | `[
              "master"
            ]` | /final/views/0/nodes/2/flags |
| e003128 | 19:58:58.350991 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /final/views/0/nodes/2/primary_id |
| e003129 | 19:58:58.350991 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/0/nodes/2/slots |
| e003130 | 19:58:58.350991 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/2/link_state |
| e003131 | 19:58:58.350991 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /final/views/0/nodes/3/flags |
| e003132 | 19:58:58.350991 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.primary_id | missing_from_schema | `` | /final/views/0/nodes/3/primary_id |
| e003133 | 19:58:58.350991 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/0/nodes/3/slots |
| e003134 | 19:58:58.350991 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/3/link_state |
| e003135 | 19:58:58.350991 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/0/nodes/4/flags |
| e003136 | 19:58:58.350991 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /final/views/0/nodes/4/primary_id |
| e003137 | 19:58:58.350991 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.slots | missing_from_schema | `` | /final/views/0/nodes/4/slots |
| e003138 | 19:58:58.350991 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/4/link_state |
| e003139 | 19:58:58.350991 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/0/nodes/5/flags |
| e003140 | 19:58:58.350991 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /final/views/0/nodes/5/primary_id |
| e003141 | 19:58:58.350991 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.slots | missing_from_schema | `` | /final/views/0/nodes/5/slots |
| e003142 | 19:58:58.350991 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-1 | peer.link_state | valid | `"connected"` | /final/views/0/nodes/5/link_state |
| e003143 | 19:58:58.350991 | node-1 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/0/persistence/aof_enabled |
| e003144 | 19:58:58.350991 | node-1 | info_persistence | persistence.loading | valid | `"0"` | /final/views/0/persistence/loading |
| e003145 | 19:58:58.350991 | node-1 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/0/persistence/aof_last_write_status |
| e003146 | 19:58:58.350991 | node-1 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/0/persistence/aof_rewrite_in_progress |
| e003147 | 19:58:58.350991 | node-1 | admin | process_live | missing_from_schema | `` | /final/views/0/process_live |
| e003148 | 19:58:58.350992 | node-4 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/3/reachable |
| e003149 | 19:58:58.350992 | node-4 | info_replication | replication.role | valid | `"slave"` | /samples/12/snapshot/views/3/replication/role |
| e003150 | 19:58:58.350992 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /samples/12/snapshot/views/3/replication/connected_slaves |
| e003151 | 19:58:58.350992 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /samples/12/snapshot/views/3/replication/master_link_status |
| e003152 | 19:58:58.350992 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/12/snapshot/views/3/replication/master_sync_in_progress |
| e003153 | 19:58:58.350992 | node-4 | info_replication | replication.master_repl_offset | valid | `441` | /samples/12/snapshot/views/3/replication/master_repl_offset |
| e003154 | 19:58:58.350992 | node-4 | info_replication | replication.slave_repl_offset | valid | `441` | /samples/12/snapshot/views/3/replication/slave_repl_offset |
| e003155 | 19:58:58.350992 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /samples/12/snapshot/views/3/redis_id |
| e003156 | 19:58:58.350992 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/3/coverage |
| e003157 | 19:58:58.350992 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/3/nodes/0/flags |
| e003158 | 19:58:58.350992 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/0/primary_id |
| e003159 | 19:58:58.350992 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/3/nodes/0/slots |
| e003160 | 19:58:58.350992 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/0/link_state |
| e003161 | 19:58:58.350992 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/12/snapshot/views/3/nodes/1/flags |
| e003162 | 19:58:58.350992 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/12/snapshot/views/3/nodes/1/primary_id |
| e003163 | 19:58:58.350992 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/1/slots |
| e003164 | 19:58:58.350992 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/1/link_state |
| e003165 | 19:58:58.350992 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/3/nodes/2/flags |
| e003166 | 19:58:58.350992 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/12/snapshot/views/3/nodes/2/primary_id |
| e003167 | 19:58:58.350992 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/2/slots |
| e003168 | 19:58:58.350992 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/2/link_state |
| e003169 | 19:58:58.350992 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/3/nodes/3/flags |
| e003170 | 19:58:58.350992 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/3/primary_id |
| e003171 | 19:58:58.350992 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/3/nodes/3/slots |
| e003172 | 19:58:58.350992 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/3/link_state |
| e003173 | 19:58:58.350992 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/3/nodes/4/flags |
| e003174 | 19:58:58.350992 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/12/snapshot/views/3/nodes/4/primary_id |
| e003175 | 19:58:58.350992 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/4/slots |
| e003176 | 19:58:58.350992 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/4/link_state |
| e003177 | 19:58:58.350992 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/3/nodes/5/flags |
| e003178 | 19:58:58.350992 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/3/nodes/5/primary_id |
| e003179 | 19:58:58.350992 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/3/nodes/5/slots |
| e003180 | 19:58:58.350992 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/3/nodes/5/link_state |
| e003181 | 19:58:58.350992 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/3/persistence/aof_enabled |
| e003182 | 19:58:58.350992 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/3/persistence/loading |
| e003183 | 19:58:58.350992 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/3/persistence/aof_last_write_status |
| e003184 | 19:58:58.350992 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/3/persistence/aof_rewrite_in_progress |
| e003185 | 19:58:58.350992 | node-4 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/3/process_live |
| e003186 | 19:58:58.350992 | node-4 | node_client | reachable | valid | `true` | /final/views/3/reachable |
| e003187 | 19:58:58.350992 | node-4 | info_replication | replication.role | valid | `"slave"` | /final/views/3/replication/role |
| e003188 | 19:58:58.350992 | node-4 | info_replication | replication.connected_slaves | valid | `0` | /final/views/3/replication/connected_slaves |
| e003189 | 19:58:58.350992 | node-4 | info_replication | replication.master_link_status | valid | `"up"` | /final/views/3/replication/master_link_status |
| e003190 | 19:58:58.350992 | node-4 | info_replication | replication.master_sync_in_progress | valid | `0` | /final/views/3/replication/master_sync_in_progress |
| e003191 | 19:58:58.350992 | node-4 | info_replication | replication.master_repl_offset | valid | `441` | /final/views/3/replication/master_repl_offset |
| e003192 | 19:58:58.350992 | node-4 | info_replication | replication.slave_repl_offset | valid | `441` | /final/views/3/replication/slave_repl_offset |
| e003193 | 19:58:58.350992 | node-4 | cluster_nodes | redis_id | valid | `"e925eb9e29de6393ceec2a30056989de4c6f770a"` | /final/views/3/redis_id |
| e003194 | 19:58:58.350992 | node-4 | cluster_nodes | slot_coverage | valid | `true` | /final/views/3/coverage |
| e003195 | 19:58:58.350992 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /final/views/3/nodes/0/flags |
| e003196 | 19:58:58.350992 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /final/views/3/nodes/0/primary_id |
| e003197 | 19:58:58.350992 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/3/nodes/0/slots |
| e003198 | 19:58:58.350992 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/0/link_state |
| e003199 | 19:58:58.350992 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /final/views/3/nodes/1/flags |
| e003200 | 19:58:58.350992 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /final/views/3/nodes/1/primary_id |
| e003201 | 19:58:58.350992 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.slots | missing_from_schema | `` | /final/views/3/nodes/1/slots |
| e003202 | 19:58:58.350992 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/1/link_state |
| e003203 | 19:58:58.350992 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/3/nodes/2/flags |
| e003204 | 19:58:58.350992 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /final/views/3/nodes/2/primary_id |
| e003205 | 19:58:58.350992 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.slots | missing_from_schema | `` | /final/views/3/nodes/2/slots |
| e003206 | 19:58:58.350992 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/2/link_state |
| e003207 | 19:58:58.350992 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /final/views/3/nodes/3/flags |
| e003208 | 19:58:58.350992 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /final/views/3/nodes/3/primary_id |
| e003209 | 19:58:58.350992 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/3/nodes/3/slots |
| e003210 | 19:58:58.350992 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/3/link_state |
| e003211 | 19:58:58.350992 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/3/nodes/4/flags |
| e003212 | 19:58:58.350992 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /final/views/3/nodes/4/primary_id |
| e003213 | 19:58:58.350992 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.slots | missing_from_schema | `` | /final/views/3/nodes/4/slots |
| e003214 | 19:58:58.350992 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/4/link_state |
| e003215 | 19:58:58.350992 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.flags | valid | `[
              "master"
            ]` | /final/views/3/nodes/5/flags |
| e003216 | 19:58:58.350992 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.primary_id | missing_from_schema | `` | /final/views/3/nodes/5/primary_id |
| e003217 | 19:58:58.350992 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/3/nodes/5/slots |
| e003218 | 19:58:58.350992 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-4 | peer.link_state | valid | `"connected"` | /final/views/3/nodes/5/link_state |
| e003219 | 19:58:58.350992 | node-4 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/3/persistence/aof_enabled |
| e003220 | 19:58:58.350992 | node-4 | info_persistence | persistence.loading | valid | `"0"` | /final/views/3/persistence/loading |
| e003221 | 19:58:58.350992 | node-4 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/3/persistence/aof_last_write_status |
| e003222 | 19:58:58.350992 | node-4 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/3/persistence/aof_rewrite_in_progress |
| e003223 | 19:58:58.350992 | node-4 | admin | process_live | missing_from_schema | `` | /final/views/3/process_live |
| e003224 | 19:58:58.350996 | node-3 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/2/reachable |
| e003225 | 19:58:58.350996 | node-3 | info_replication | replication.role | valid | `"master"` | /samples/12/snapshot/views/2/replication/role |
| e003226 | 19:58:58.350996 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /samples/12/snapshot/views/2/replication/connected_slaves |
| e003227 | 19:58:58.350996 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /samples/12/snapshot/views/2/replication/master_link_status |
| e003228 | 19:58:58.350996 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /samples/12/snapshot/views/2/replication/master_sync_in_progress |
| e003229 | 19:58:58.350996 | node-3 | info_replication | replication.master_repl_offset | valid | `441` | /samples/12/snapshot/views/2/replication/master_repl_offset |
| e003230 | 19:58:58.350996 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /samples/12/snapshot/views/2/replication/slave_repl_offset |
| e003231 | 19:58:58.350996 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/12/snapshot/views/2/redis_id |
| e003232 | 19:58:58.350996 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/2/coverage |
| e003233 | 19:58:58.350996 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/2/nodes/0/flags |
| e003234 | 19:58:58.350996 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/0/primary_id |
| e003235 | 19:58:58.350996 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/2/nodes/0/slots |
| e003236 | 19:58:58.350996 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/0/link_state |
| e003237 | 19:58:58.350996 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
                  "myself",
                  "master"
                ]` | /samples/12/snapshot/views/2/nodes/1/flags |
| e003238 | 19:58:58.350996 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/1/primary_id |
| e003239 | 19:58:58.350996 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/2/nodes/1/slots |
| e003240 | 19:58:58.350996 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/1/link_state |
| e003241 | 19:58:58.350996 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/2/nodes/2/flags |
| e003242 | 19:58:58.350996 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/12/snapshot/views/2/nodes/2/primary_id |
| e003243 | 19:58:58.350996 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/2/slots |
| e003244 | 19:58:58.350996 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/2/link_state |
| e003245 | 19:58:58.350996 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/2/nodes/3/flags |
| e003246 | 19:58:58.350996 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/12/snapshot/views/2/nodes/3/primary_id |
| e003247 | 19:58:58.350996 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/3/slots |
| e003248 | 19:58:58.350996 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/3/link_state |
| e003249 | 19:58:58.350996 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/2/nodes/4/flags |
| e003250 | 19:58:58.350996 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/12/snapshot/views/2/nodes/4/primary_id |
| e003251 | 19:58:58.350996 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/4/slots |
| e003252 | 19:58:58.350996 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/4/link_state |
| e003253 | 19:58:58.350996 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/2/nodes/5/flags |
| e003254 | 19:58:58.350996 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/2/nodes/5/primary_id |
| e003255 | 19:58:58.350996 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/2/nodes/5/slots |
| e003256 | 19:58:58.350996 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/2/nodes/5/link_state |
| e003257 | 19:58:58.350996 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/2/persistence/aof_enabled |
| e003258 | 19:58:58.350996 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/2/persistence/loading |
| e003259 | 19:58:58.350996 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/2/persistence/aof_last_write_status |
| e003260 | 19:58:58.350996 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/2/persistence/aof_rewrite_in_progress |
| e003261 | 19:58:58.350996 | node-3 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/2/process_live |
| e003262 | 19:58:58.350996 | node-3 | node_client | reachable | valid | `true` | /final/views/2/reachable |
| e003263 | 19:58:58.350996 | node-3 | info_replication | replication.role | valid | `"master"` | /final/views/2/replication/role |
| e003264 | 19:58:58.350996 | node-3 | info_replication | replication.connected_slaves | valid | `1` | /final/views/2/replication/connected_slaves |
| e003265 | 19:58:58.350996 | node-3 | info_replication | replication.master_link_status | not_applicable | `` | /final/views/2/replication/master_link_status |
| e003266 | 19:58:58.350996 | node-3 | info_replication | replication.master_sync_in_progress | not_applicable | `` | /final/views/2/replication/master_sync_in_progress |
| e003267 | 19:58:58.350996 | node-3 | info_replication | replication.master_repl_offset | valid | `441` | /final/views/2/replication/master_repl_offset |
| e003268 | 19:58:58.350996 | node-3 | info_replication | replication.slave_repl_offset | not_applicable | `` | /final/views/2/replication/slave_repl_offset |
| e003269 | 19:58:58.350996 | node-3 | cluster_nodes | redis_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /final/views/2/redis_id |
| e003270 | 19:58:58.350996 | node-3 | cluster_nodes | slot_coverage | valid | `true` | /final/views/2/coverage |
| e003271 | 19:58:58.350996 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /final/views/2/nodes/0/flags |
| e003272 | 19:58:58.350996 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /final/views/2/nodes/0/primary_id |
| e003273 | 19:58:58.350996 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/2/nodes/0/slots |
| e003274 | 19:58:58.350996 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/0/link_state |
| e003275 | 19:58:58.350996 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.flags | valid | `[
              "myself",
              "master"
            ]` | /final/views/2/nodes/1/flags |
| e003276 | 19:58:58.350996 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /final/views/2/nodes/1/primary_id |
| e003277 | 19:58:58.350996 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/2/nodes/1/slots |
| e003278 | 19:58:58.350996 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/1/link_state |
| e003279 | 19:58:58.350996 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/2/nodes/2/flags |
| e003280 | 19:58:58.350996 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /final/views/2/nodes/2/primary_id |
| e003281 | 19:58:58.350996 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.slots | missing_from_schema | `` | /final/views/2/nodes/2/slots |
| e003282 | 19:58:58.350996 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/2/link_state |
| e003283 | 19:58:58.350996 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/2/nodes/3/flags |
| e003284 | 19:58:58.350996 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /final/views/2/nodes/3/primary_id |
| e003285 | 19:58:58.350996 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.slots | missing_from_schema | `` | /final/views/2/nodes/3/slots |
| e003286 | 19:58:58.350996 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/3/link_state |
| e003287 | 19:58:58.350996 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/2/nodes/4/flags |
| e003288 | 19:58:58.350996 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /final/views/2/nodes/4/primary_id |
| e003289 | 19:58:58.350996 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.slots | missing_from_schema | `` | /final/views/2/nodes/4/slots |
| e003290 | 19:58:58.350996 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/4/link_state |
| e003291 | 19:58:58.350996 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.flags | valid | `[
              "master"
            ]` | /final/views/2/nodes/5/flags |
| e003292 | 19:58:58.350996 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.primary_id | missing_from_schema | `` | /final/views/2/nodes/5/primary_id |
| e003293 | 19:58:58.350996 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/2/nodes/5/slots |
| e003294 | 19:58:58.350996 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-3 | peer.link_state | valid | `"connected"` | /final/views/2/nodes/5/link_state |
| e003295 | 19:58:58.350996 | node-3 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/2/persistence/aof_enabled |
| e003296 | 19:58:58.350996 | node-3 | info_persistence | persistence.loading | valid | `"0"` | /final/views/2/persistence/loading |
| e003297 | 19:58:58.350996 | node-3 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/2/persistence/aof_last_write_status |
| e003298 | 19:58:58.350996 | node-3 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/2/persistence/aof_rewrite_in_progress |
| e003299 | 19:58:58.350996 | node-3 | admin | process_live | missing_from_schema | `` | /final/views/2/process_live |
| e003300 | 19:58:58.351006 | node-6 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/5/reachable |
| e003301 | 19:58:58.351006 | node-6 | info_replication | replication.role | valid | `"slave"` | /samples/12/snapshot/views/5/replication/role |
| e003302 | 19:58:58.351006 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /samples/12/snapshot/views/5/replication/connected_slaves |
| e003303 | 19:58:58.351006 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /samples/12/snapshot/views/5/replication/master_link_status |
| e003304 | 19:58:58.351006 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/12/snapshot/views/5/replication/master_sync_in_progress |
| e003305 | 19:58:58.351006 | node-6 | info_replication | replication.master_repl_offset | valid | `441` | /samples/12/snapshot/views/5/replication/master_repl_offset |
| e003306 | 19:58:58.351006 | node-6 | info_replication | replication.slave_repl_offset | valid | `441` | /samples/12/snapshot/views/5/replication/slave_repl_offset |
| e003307 | 19:58:58.351006 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /samples/12/snapshot/views/5/redis_id |
| e003308 | 19:58:58.351006 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/5/coverage |
| e003309 | 19:58:58.351006 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/5/nodes/0/flags |
| e003310 | 19:58:58.351006 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/12/snapshot/views/5/nodes/0/primary_id |
| e003311 | 19:58:58.351006 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/0/slots |
| e003312 | 19:58:58.351006 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/0/link_state |
| e003313 | 19:58:58.351006 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/5/nodes/1/flags |
| e003314 | 19:58:58.351006 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/1/primary_id |
| e003315 | 19:58:58.351006 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/5/nodes/1/slots |
| e003316 | 19:58:58.351006 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/1/link_state |
| e003317 | 19:58:58.351006 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/5/nodes/2/flags |
| e003318 | 19:58:58.351006 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/2/primary_id |
| e003319 | 19:58:58.351006 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/5/nodes/2/slots |
| e003320 | 19:58:58.351006 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/2/link_state |
| e003321 | 19:58:58.351006 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/5/nodes/3/flags |
| e003322 | 19:58:58.351006 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/3/primary_id |
| e003323 | 19:58:58.351006 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/5/nodes/3/slots |
| e003324 | 19:58:58.351006 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/3/link_state |
| e003325 | 19:58:58.351006 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/5/nodes/4/flags |
| e003326 | 19:58:58.351006 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/12/snapshot/views/5/nodes/4/primary_id |
| e003327 | 19:58:58.351006 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/4/slots |
| e003328 | 19:58:58.351006 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/4/link_state |
| e003329 | 19:58:58.351006 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/12/snapshot/views/5/nodes/5/flags |
| e003330 | 19:58:58.351006 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/12/snapshot/views/5/nodes/5/primary_id |
| e003331 | 19:58:58.351006 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/5/nodes/5/slots |
| e003332 | 19:58:58.351006 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/5/nodes/5/link_state |
| e003333 | 19:58:58.351006 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/5/persistence/aof_enabled |
| e003334 | 19:58:58.351006 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/5/persistence/loading |
| e003335 | 19:58:58.351006 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/5/persistence/aof_last_write_status |
| e003336 | 19:58:58.351006 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/5/persistence/aof_rewrite_in_progress |
| e003337 | 19:58:58.351006 | node-6 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/5/process_live |
| e003338 | 19:58:58.351006 | node-6 | node_client | reachable | valid | `true` | /final/views/5/reachable |
| e003339 | 19:58:58.351006 | node-6 | info_replication | replication.role | valid | `"slave"` | /final/views/5/replication/role |
| e003340 | 19:58:58.351006 | node-6 | info_replication | replication.connected_slaves | valid | `0` | /final/views/5/replication/connected_slaves |
| e003341 | 19:58:58.351006 | node-6 | info_replication | replication.master_link_status | valid | `"up"` | /final/views/5/replication/master_link_status |
| e003342 | 19:58:58.351006 | node-6 | info_replication | replication.master_sync_in_progress | valid | `0` | /final/views/5/replication/master_sync_in_progress |
| e003343 | 19:58:58.351006 | node-6 | info_replication | replication.master_repl_offset | valid | `441` | /final/views/5/replication/master_repl_offset |
| e003344 | 19:58:58.351006 | node-6 | info_replication | replication.slave_repl_offset | valid | `441` | /final/views/5/replication/slave_repl_offset |
| e003345 | 19:58:58.351006 | node-6 | cluster_nodes | redis_id | valid | `"1cfbe152fc4e94c0102c9e95c8a3e262e71872bc"` | /final/views/5/redis_id |
| e003346 | 19:58:58.351006 | node-6 | cluster_nodes | slot_coverage | valid | `true` | /final/views/5/coverage |
| e003347 | 19:58:58.351006 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/5/nodes/0/flags |
| e003348 | 19:58:58.351006 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /final/views/5/nodes/0/primary_id |
| e003349 | 19:58:58.351006 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.slots | missing_from_schema | `` | /final/views/5/nodes/0/slots |
| e003350 | 19:58:58.351006 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/0/link_state |
| e003351 | 19:58:58.351006 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /final/views/5/nodes/1/flags |
| e003352 | 19:58:58.351006 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /final/views/5/nodes/1/primary_id |
| e003353 | 19:58:58.351006 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/5/nodes/1/slots |
| e003354 | 19:58:58.351006 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/1/link_state |
| e003355 | 19:58:58.351006 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /final/views/5/nodes/2/flags |
| e003356 | 19:58:58.351006 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /final/views/5/nodes/2/primary_id |
| e003357 | 19:58:58.351006 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/5/nodes/2/slots |
| e003358 | 19:58:58.351006 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/2/link_state |
| e003359 | 19:58:58.351006 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.flags | valid | `[
              "master"
            ]` | /final/views/5/nodes/3/flags |
| e003360 | 19:58:58.351006 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.primary_id | missing_from_schema | `` | /final/views/5/nodes/3/primary_id |
| e003361 | 19:58:58.351006 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/5/nodes/3/slots |
| e003362 | 19:58:58.351006 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/3/link_state |
| e003363 | 19:58:58.351006 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/5/nodes/4/flags |
| e003364 | 19:58:58.351006 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /final/views/5/nodes/4/primary_id |
| e003365 | 19:58:58.351006 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.slots | missing_from_schema | `` | /final/views/5/nodes/4/slots |
| e003366 | 19:58:58.351006 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/4/link_state |
| e003367 | 19:58:58.351006 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /final/views/5/nodes/5/flags |
| e003368 | 19:58:58.351006 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /final/views/5/nodes/5/primary_id |
| e003369 | 19:58:58.351006 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.slots | missing_from_schema | `` | /final/views/5/nodes/5/slots |
| e003370 | 19:58:58.351006 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-6 | peer.link_state | valid | `"connected"` | /final/views/5/nodes/5/link_state |
| e003371 | 19:58:58.351006 | node-6 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/5/persistence/aof_enabled |
| e003372 | 19:58:58.351006 | node-6 | info_persistence | persistence.loading | valid | `"0"` | /final/views/5/persistence/loading |
| e003373 | 19:58:58.351006 | node-6 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/5/persistence/aof_last_write_status |
| e003374 | 19:58:58.351006 | node-6 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/5/persistence/aof_rewrite_in_progress |
| e003375 | 19:58:58.351006 | node-6 | admin | process_live | missing_from_schema | `` | /final/views/5/process_live |
| e003376 | 19:58:58.351025 | node-5 | node_client | reachable | valid | `true` | /samples/12/snapshot/views/4/reachable |
| e003377 | 19:58:58.351025 | node-5 | info_replication | replication.role | valid | `"slave"` | /samples/12/snapshot/views/4/replication/role |
| e003378 | 19:58:58.351025 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /samples/12/snapshot/views/4/replication/connected_slaves |
| e003379 | 19:58:58.351025 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /samples/12/snapshot/views/4/replication/master_link_status |
| e003380 | 19:58:58.351025 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /samples/12/snapshot/views/4/replication/master_sync_in_progress |
| e003381 | 19:58:58.351025 | node-5 | info_replication | replication.master_repl_offset | valid | `455` | /samples/12/snapshot/views/4/replication/master_repl_offset |
| e003382 | 19:58:58.351025 | node-5 | info_replication | replication.slave_repl_offset | valid | `455` | /samples/12/snapshot/views/4/replication/slave_repl_offset |
| e003383 | 19:58:58.351025 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /samples/12/snapshot/views/4/redis_id |
| e003384 | 19:58:58.351025 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /samples/12/snapshot/views/4/coverage |
| e003385 | 19:58:58.351025 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/4/nodes/0/flags |
| e003386 | 19:58:58.351025 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/0/primary_id |
| e003387 | 19:58:58.351025 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 5461,
                    "end": 10922
                  }
                ]` | /samples/12/snapshot/views/4/nodes/0/slots |
| e003388 | 19:58:58.351025 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/0/link_state |
| e003389 | 19:58:58.351025 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/4/nodes/1/flags |
| e003390 | 19:58:58.351025 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /samples/12/snapshot/views/4/nodes/1/primary_id |
| e003391 | 19:58:58.351025 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/1/slots |
| e003392 | 19:58:58.351025 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/1/link_state |
| e003393 | 19:58:58.351025 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/4/nodes/2/flags |
| e003394 | 19:58:58.351025 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/2/primary_id |
| e003395 | 19:58:58.351025 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 10923,
                    "end": 16383
                  }
                ]` | /samples/12/snapshot/views/4/nodes/2/slots |
| e003396 | 19:58:58.351025 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/2/link_state |
| e003397 | 19:58:58.351025 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
                  "master"
                ]` | /samples/12/snapshot/views/4/nodes/3/flags |
| e003398 | 19:58:58.351025 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/3/primary_id |
| e003399 | 19:58:58.351025 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
                  {
                    "start": 0,
                    "end": 5460
                  }
                ]` | /samples/12/snapshot/views/4/nodes/3/slots |
| e003400 | 19:58:58.351025 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/3/link_state |
| e003401 | 19:58:58.351025 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
                  "myself",
                  "slave"
                ]` | /samples/12/snapshot/views/4/nodes/4/flags |
| e003402 | 19:58:58.351025 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /samples/12/snapshot/views/4/nodes/4/primary_id |
| e003403 | 19:58:58.351025 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/4/slots |
| e003404 | 19:58:58.351025 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/4/link_state |
| e003405 | 19:58:58.351025 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
                  "slave"
                ]` | /samples/12/snapshot/views/4/nodes/5/flags |
| e003406 | 19:58:58.351025 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /samples/12/snapshot/views/4/nodes/5/primary_id |
| e003407 | 19:58:58.351025 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /samples/12/snapshot/views/4/nodes/5/slots |
| e003408 | 19:58:58.351025 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /samples/12/snapshot/views/4/nodes/5/link_state |
| e003409 | 19:58:58.351025 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /samples/12/snapshot/views/4/persistence/aof_enabled |
| e003410 | 19:58:58.351025 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /samples/12/snapshot/views/4/persistence/loading |
| e003411 | 19:58:58.351025 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /samples/12/snapshot/views/4/persistence/aof_last_write_status |
| e003412 | 19:58:58.351025 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /samples/12/snapshot/views/4/persistence/aof_rewrite_in_progress |
| e003413 | 19:58:58.351025 | node-5 | admin | process_live | missing_from_schema | `` | /samples/12/snapshot/views/4/process_live |
| e003414 | 19:58:58.351025 | node-5 | node_client | reachable | valid | `true` | /final/views/4/reachable |
| e003415 | 19:58:58.351025 | node-5 | info_replication | replication.role | valid | `"slave"` | /final/views/4/replication/role |
| e003416 | 19:58:58.351025 | node-5 | info_replication | replication.connected_slaves | valid | `0` | /final/views/4/replication/connected_slaves |
| e003417 | 19:58:58.351025 | node-5 | info_replication | replication.master_link_status | valid | `"up"` | /final/views/4/replication/master_link_status |
| e003418 | 19:58:58.351025 | node-5 | info_replication | replication.master_sync_in_progress | valid | `0` | /final/views/4/replication/master_sync_in_progress |
| e003419 | 19:58:58.351025 | node-5 | info_replication | replication.master_repl_offset | valid | `455` | /final/views/4/replication/master_repl_offset |
| e003420 | 19:58:58.351025 | node-5 | info_replication | replication.slave_repl_offset | valid | `455` | /final/views/4/replication/slave_repl_offset |
| e003421 | 19:58:58.351025 | node-5 | cluster_nodes | redis_id | valid | `"361fdf9ad3ee6577fd53ca9afbf430d205fe6c30"` | /final/views/4/redis_id |
| e003422 | 19:58:58.351025 | node-5 | cluster_nodes | slot_coverage | valid | `true` | /final/views/4/coverage |
| e003423 | 19:58:58.351025 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /final/views/4/nodes/0/flags |
| e003424 | 19:58:58.351025 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /final/views/4/nodes/0/primary_id |
| e003425 | 19:58:58.351025 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 5461,
                "end": 10922
              }
            ]` | /final/views/4/nodes/0/slots |
| e003426 | 19:58:58.351025 | 7f385e1d484746697faa90765e66e2563b95f1b7 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/0/link_state |
| e003427 | 19:58:58.351025 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/4/nodes/1/flags |
| e003428 | 19:58:58.351025 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.primary_id | valid | `"3017030ae4cc288de4885b1c31d4086bf0a2037f"` | /final/views/4/nodes/1/primary_id |
| e003429 | 19:58:58.351025 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.slots | missing_from_schema | `` | /final/views/4/nodes/1/slots |
| e003430 | 19:58:58.351025 | e925eb9e29de6393ceec2a30056989de4c6f770a | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/1/link_state |
| e003431 | 19:58:58.351025 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /final/views/4/nodes/2/flags |
| e003432 | 19:58:58.351025 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /final/views/4/nodes/2/primary_id |
| e003433 | 19:58:58.351025 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 10923,
                "end": 16383
              }
            ]` | /final/views/4/nodes/2/slots |
| e003434 | 19:58:58.351025 | 3017030ae4cc288de4885b1c31d4086bf0a2037f | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/2/link_state |
| e003435 | 19:58:58.351025 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.flags | valid | `[
              "master"
            ]` | /final/views/4/nodes/3/flags |
| e003436 | 19:58:58.351025 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.primary_id | missing_from_schema | `` | /final/views/4/nodes/3/primary_id |
| e003437 | 19:58:58.351025 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.slots | valid | `[
              {
                "start": 0,
                "end": 5460
              }
            ]` | /final/views/4/nodes/3/slots |
| e003438 | 19:58:58.351025 | 4f9cdefa31d52b1ec07962fe48e90304775f32b4 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/3/link_state |
| e003439 | 19:58:58.351025 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.flags | valid | `[
              "myself",
              "slave"
            ]` | /final/views/4/nodes/4/flags |
| e003440 | 19:58:58.351025 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.primary_id | valid | `"4f9cdefa31d52b1ec07962fe48e90304775f32b4"` | /final/views/4/nodes/4/primary_id |
| e003441 | 19:58:58.351025 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.slots | missing_from_schema | `` | /final/views/4/nodes/4/slots |
| e003442 | 19:58:58.351025 | 361fdf9ad3ee6577fd53ca9afbf430d205fe6c30 | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/4/link_state |
| e003443 | 19:58:58.351025 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.flags | valid | `[
              "slave"
            ]` | /final/views/4/nodes/5/flags |
| e003444 | 19:58:58.351025 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.primary_id | valid | `"7f385e1d484746697faa90765e66e2563b95f1b7"` | /final/views/4/nodes/5/primary_id |
| e003445 | 19:58:58.351025 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.slots | missing_from_schema | `` | /final/views/4/nodes/5/slots |
| e003446 | 19:58:58.351025 | 1cfbe152fc4e94c0102c9e95c8a3e262e71872bc | peer_view:node-5 | peer.link_state | valid | `"connected"` | /final/views/4/nodes/5/link_state |
| e003447 | 19:58:58.351025 | node-5 | info_persistence | persistence.aof_enabled | valid | `"0"` | /final/views/4/persistence/aof_enabled |
| e003448 | 19:58:58.351025 | node-5 | info_persistence | persistence.loading | valid | `"0"` | /final/views/4/persistence/loading |
| e003449 | 19:58:58.351025 | node-5 | info_persistence | persistence.aof_last_write_status | valid | `"ok"` | /final/views/4/persistence/aof_last_write_status |
| e003450 | 19:58:58.351025 | node-5 | info_persistence | persistence.aof_rewrite_in_progress | valid | `"0"` | /final/views/4/persistence/aof_rewrite_in_progress |
| e003451 | 19:58:58.351025 | node-5 | admin | process_live | missing_from_schema | `` | /final/views/4/process_live |
| e003452 | 19:58:58.36262 |  | client_probe | client_success | valid | `true` | /samples/12/probe/successful |

## Analysis limitations

- Offline laboratory observation analysis; no production integration.
- Node views are non-atomic and retained independently; sample timestamps are not exact server event times.
- No fault labels, injected action results, run identifiers, configuration hashes or phase names are inference inputs.
- Client reachability does not establish process death; no administrative process telemetry is accepted from fault-specific fields.
- Validation proves only the recorded checked population, never general durability or power-loss safety.
