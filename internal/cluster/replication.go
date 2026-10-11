package cluster

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

// ReplicationMetrics preserves optional typed indicators alongside the source
// INFO map. Offset differences describe replication progress, not lost values.
type ReplicationMetrics struct {
	Role              string            `json:"role,omitempty"`
	ReplicationID     string            `json:"replication_id,omitempty"`
	MasterOffset      *int64            `json:"master_offset,omitempty"`
	ReplicaOffset     *int64            `json:"replica_offset,omitempty"`
	ReplicaReadOffset *int64            `json:"replica_read_offset,omitempty"`
	ConnectedReplicas *int64            `json:"connected_replicas,omitempty"`
	LinkUp            *bool             `json:"link_up,omitempty"`
	SyncInProgress    *bool             `json:"sync_in_progress,omitempty"`
	LastIOSeconds     *int64            `json:"last_io_seconds,omitempty"`
	LinkDownSeconds   *int64            `json:"link_down_seconds,omitempty"`
	BacklogBytes      *int64            `json:"backlog_bytes,omitempty"`
	Replicas          []ReplicaProgress `json:"replicas,omitempty"`
}

type ReplicaProgress struct {
	Name           string `json:"name"`
	IP             string `json:"ip,omitempty"`
	Port           string `json:"port,omitempty"`
	State          string `json:"state,omitempty"`
	Offset         *int64 `json:"offset,omitempty"`
	LagSeconds     *int64 `json:"lag_seconds,omitempty"`
	OffsetGapBytes *int64 `json:"offset_gap_bytes,omitempty"`
}

func infoNumber(value string) *int64 {
	n, err := strconv.ParseInt(value, 10, 64)
	if err != nil || n < 0 {
		return nil
	}
	return &n
}

func infoFlag(value, yes, no string) *bool {
	if value != yes && value != no {
		return nil
	}
	flag := value == yes
	return &flag
}

func DeriveReplication(info map[string]string) ReplicationMetrics {
	m := ReplicationMetrics{
		Role: info["role"], ReplicationID: info["master_replid"],
		MasterOffset: infoNumber(info["master_repl_offset"]), ReplicaOffset: infoNumber(info["slave_repl_offset"]),
		ReplicaReadOffset: infoNumber(info["slave_read_repl_offset"]), ConnectedReplicas: infoNumber(info["connected_slaves"]),
		LinkUp: infoFlag(info["master_link_status"], "up", "down"), SyncInProgress: infoFlag(info["master_sync_in_progress"], "1", "0"),
		LastIOSeconds: infoNumber(info["master_last_io_seconds_ago"]), LinkDownSeconds: infoNumber(info["master_link_down_since_seconds"]),
		BacklogBytes: infoNumber(info["repl_backlog_histlen"]),
	}
	for name, value := range info {
		if !strings.HasPrefix(name, "slave") {
			continue
		}
		index, err := strconv.Atoi(strings.TrimPrefix(name, "slave"))
		if err != nil || index < 0 {
			continue
		}
		fields := make(map[string]string)
		for _, field := range strings.Split(value, ",") {
			key, value, found := strings.Cut(field, "=")
			if found {
				fields[key] = value
			}
		}
		r := ReplicaProgress{Name: name, IP: fields["ip"], Port: fields["port"], State: fields["state"], Offset: infoNumber(fields["offset"]), LagSeconds: infoNumber(fields["lag"])}
		if m.MasterOffset != nil && r.Offset != nil && *m.MasterOffset >= *r.Offset {
			gap := *m.MasterOffset - *r.Offset
			r.OffsetGapBytes = &gap
		}
		m.Replicas = append(m.Replicas, r)
	}
	sort.Slice(m.Replicas, func(i, j int) bool {
		a, _ := strconv.Atoi(strings.TrimPrefix(m.Replicas[i].Name, "slave"))
		b, _ := strconv.Atoi(strings.TrimPrefix(m.Replicas[j].Name, "slave"))
		return a < b
	})
	return m
}

type ResynchronizationEvidence struct {
	PrimaryRedisID       string    `json:"primary_redis_id,omitempty"`
	BeforeAt             time.Time `json:"before_at"`
	AfterAt              time.Time `json:"after_at"`
	FullSyncDelta        *uint64   `json:"full_sync_delta,omitempty"`
	PartialAcceptedDelta *uint64   `json:"partial_accepted_delta,omitempty"`
	PartialRejectedDelta *uint64   `json:"partial_rejected_delta,omitempty"`
	Inference            string    `json:"inference"`
	Reasons              []string  `json:"reasons,omitempty"`
	Methodology          string    `json:"methodology"`
}

const ResynchronizationMethodology = "INFO stats sync_full, sync_partial_ok and sync_partial_err deltas are compared only on the same directly observed primary identity, endpoints, Redis process run ID and replication ID. Decreasing, missing or invalid counters remain unknown. Positive counters are server-wide evidence of synchronization requests and do not prove a particular replica completed synchronization; link-up and master_sync_in_progress=0 are measured separately. Replication offsets and lag are progress indicators, never exact byte-loss or durability measurements."

// ResynchronizationDelta uses primary-side counters. Replica restarts reset
// local counters and cannot independently establish full or partial resync.
// Redis documents these fields at https://redis.io/docs/latest/commands/info/.
func ResynchronizationDelta(before, after View) ResynchronizationEvidence {
	e := ResynchronizationEvidence{PrimaryRedisID: after.RedisID, BeforeAt: before.At, AfterAt: after.At, Inference: "unknown", Methodology: ResynchronizationMethodology}
	for _, v := range []View{before, after} {
		self, found := findNode(v.Nodes, v.RedisID)
		if !v.Reachable || len(v.Errors) != 0 || !found || !self.Primary() || self.Replica() || !self.Healthy() || v.Replication["role"] != "master" {
			e.Reasons = append(e.Reasons, "both observations must directly identify a healthy primary")
			return e
		}
	}
	if before.RedisID == "" || before.RedisID != after.RedisID || before.Node != after.Node {
		e.Reasons = append(e.Reasons, "primary identity or managed endpoints changed")
		return e
	}
	server := func(v View) map[string]string {
		if v.Server != nil {
			return v.Server
		}
		return ParseInfo(v.RawServer)
	}
	if server(before)["run_id"] == "" || server(before)["run_id"] != server(after)["run_id"] {
		e.Reasons = append(e.Reasons, "primary Redis process run ID is absent or changed")
		return e
	}
	if before.Replication["master_replid"] == "" || before.Replication["master_replid"] != after.Replication["master_replid"] {
		e.Reasons = append(e.Reasons, "primary replication ID is absent or changed")
		return e
	}
	if before.At.IsZero() || after.At.IsZero() || after.At.Before(before.At) {
		e.Reasons = append(e.Reasons, "counter observation interval is absent or reversed")
		return e
	}
	delta := func(key string) *uint64 {
		a, aerr := strconv.ParseUint(before.Stats[key], 10, 64)
		b, berr := strconv.ParseUint(after.Stats[key], 10, 64)
		if aerr != nil || berr != nil || b < a {
			e.Reasons = append(e.Reasons, key+" is missing, invalid, or decreased")
			return nil
		}
		n := b - a
		return &n
	}
	e.FullSyncDelta, e.PartialAcceptedDelta, e.PartialRejectedDelta = delta("sync_full"), delta("sync_partial_ok"), delta("sync_partial_err")
	if len(e.Reasons) != 0 {
		// A decrease may represent a server-wide CONFIG RESETSTAT. Without
		// complete comparable counters, do not infer deltas for the interval.
		e.FullSyncDelta, e.PartialAcceptedDelta, e.PartialRejectedDelta = nil, nil, nil
		return e
	}
	positive := func(n *uint64) bool { return n != nil && *n > 0 }
	switch {
	case positive(e.FullSyncDelta) && positive(e.PartialAcceptedDelta):
		e.Inference = "full_and_partial_sync_counters_increased"
	case positive(e.FullSyncDelta):
		e.Inference = "full_sync_counter_increased"
	case positive(e.PartialAcceptedDelta):
		e.Inference = "partial_sync_counter_increased"
	case positive(e.PartialRejectedDelta):
		e.Inference = "partial_sync_request_rejected"
	case len(e.Reasons) == 0:
		e.Inference = "no_counter_increase_observed"
	}
	return e
}
