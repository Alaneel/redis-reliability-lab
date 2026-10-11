package diagnosis

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"
)

type object map[string]json.RawMessage

func obj(b json.RawMessage) object { var o object; json.Unmarshal(b, &o); return o }
func arr(b json.RawMessage) []json.RawMessage {
	var a []json.RawMessage
	json.Unmarshal(b, &a)
	return a
}
func txt(b json.RawMessage) string      { var s string; json.Unmarshal(b, &s); return s }
func stamp(b json.RawMessage) time.Time { var t time.Time; json.Unmarshal(b, &t); return t }

// Normalize is a one-way allowlist adapter. Fault actions, run identifiers,
// filenames, configuration, lifecycle dimensions and expected outcomes are
// never decoded into inference signals. References are JSON pointers only.
func Normalize(ctx context.Context, reader io.Reader) (Incident, error) {
	data, err := io.ReadAll(io.LimitReader(reader, MaxInputBytes+1))
	if err != nil {
		return Incident{}, err
	}
	if len(data) > MaxInputBytes {
		return Incident{}, fmt.Errorf("report exceeds %d bytes", MaxInputBytes)
	}
	var root object
	dec := json.NewDecoder(bytes.NewReader(data))
	if err = dec.Decode(&root); err != nil {
		return Incident{}, fmt.Errorf("invalid report: %w", err)
	}
	var extra any
	if dec.Decode(&extra) != io.EOF {
		return Incident{}, fmt.Errorf("trailing report data")
	}
	for _, k := range []string{"initial", "final", "validation", "metrics"} {
		if b, ok := root[k]; ok && !bytes.Equal(b, []byte("null")) {
			var v object
			if json.Unmarshal(b, &v) != nil {
				return Incident{}, fmt.Errorf("%s must be an object", k)
			}
		}
	}
	if b, ok := root["samples"]; ok && !bytes.Equal(b, []byte("null")) {
		var v []json.RawMessage
		if json.Unmarshal(b, &v) != nil {
			return Incident{}, fmt.Errorf("samples must be an array")
		}
	}
	var schema int
	if json.Unmarshal(root["schema_version"], &schema) != nil || schema < 1 || schema > 3 {
		return Incident{}, fmt.Errorf("unsupported report schema")
	}
	sum := sha256.Sum256(data)
	out := Incident{SchemaVersion: SchemaVersion, InputSchema: schema, SourceSHA256: hex.EncodeToString(sum[:]), Evidence: []Evidence{}, Limitations: []string{
		"Offline laboratory observation analysis; no production integration.",
		"Node views are non-atomic and retained independently; sample timestamps are not exact server event times.",
		"No fault labels, injected action results, run identifiers, configuration hashes or phase names are inference inputs.",
		"Client reachability does not establish process death; no administrative process telemetry is accepted from fault-specific fields.",
		"Validation proves only the recorded checked population, never general durability or power-loss safety.",
	}}
	stalePrefixes := map[string]bool{}
	overflow := false
	var projectionError error
	add := func(at time.Time, node, source, signal, ref string, raw json.RawMessage, status Status) {
		if len(raw) > 65536 {
			projectionError = fmt.Errorf("projected field exceeds 64 KiB: %s", ref)
			return
		}
		if status == Valid {
			if len(raw) == 0 {
				status = Missing
			} else if bytes.Equal(raw, []byte("null")) {
				status = Unavailable
			} else if at.IsZero() {
				status = Invalid
			}
		}
		if status == Valid {
			if before, after, ok := strings.Cut(ref, "/views/"); ok {
				index, _, _ := strings.Cut(after, "/")
				if stalePrefixes[before+"/views/"+index] {
					status = Invalid
				}
			}
		}
		if status == Valid && !validSignalValue(signal, raw) {
			status = Invalid
		}
		quality := "direct_recorded_observation"
		if status != Valid {
			quality = "unknown"
		}
		if len(out.Evidence) >= MaxEvidence {
			overflow = true
			return
		}
		out.Evidence = append(out.Evidence, Evidence{At: at, Node: node, Source: source, Signal: signal, Reference: ref, Value: raw, Status: status, Quality: quality})
	}
	encoded := func(v any) json.RawMessage { b, _ := json.Marshal(v); return b }
	snapshot := func(raw json.RawMessage, ref string) error {
		s := obj(raw)
		if len(s) == 0 {
			return nil
		}
		var views []json.RawMessage
		if b, ok := s["views"]; ok && json.Unmarshal(b, &views) != nil {
			return fmt.Errorf("views must be an array")
		}
		begin, end := stamp(s["at"]), stamp(s["completed"])
		if len(views) > 64 {
			return fmt.Errorf("too many node views")
		}
		for i, rawView := range views {
			if err := ctx.Err(); err != nil {
				return err
			}
			v := obj(rawView)
			node := txt(obj(v["managed_node"])["id"])
			if len(node) > 256 {
				return fmt.Errorf("node identity exceeds 256 bytes")
			}
			if node == "" {
				return fmt.Errorf("view lacks node identity")
			}
			at := stamp(v["at"])
			p := fmt.Sprintf("%s/views/%d", ref, i)
			if (!begin.IsZero() && at.Before(begin.Add(-time.Second))) || (!end.IsZero() && at.After(end.Add(time.Second))) {
				stalePrefixes[p] = true
			}
			add(at, node, "node_client", "reachable", p+"/reachable", v["reachable"], Valid)
			var reachable bool
			json.Unmarshal(v["reachable"], &reachable)
			status := Valid
			if len(v["reachable"]) == 0 {
				status = Missing
			} else if json.Unmarshal(v["reachable"], &reachable) != nil {
				status = Invalid
			} else if !reachable {
				status = Failed
			}
			repl := obj(v["replication"])
			for _, key := range []string{"role", "connected_slaves", "master_link_status", "master_sync_in_progress", "master_repl_offset", "slave_repl_offset"} {
				r := repl[key]
				signal := "replication." + key
				if key == "connected_slaves" || key == "master_sync_in_progress" || key == "master_repl_offset" || key == "slave_repl_offset" {
					if len(r) > 0 {
						n, err := strconv.ParseInt(txt(r), 10, 64)
						if err == nil && n >= 0 {
							r = encoded(n)
						} else {
							add(at, node, "info_replication", signal, p+"/replication/"+key, r, Invalid)
							continue
						}
					}
				}
				st := status
				if reachable && len(r) == 0 {
					st = Missing
				}
				if reachable && ((key == "master_link_status" || key == "master_sync_in_progress" || key == "slave_repl_offset") && txt(repl["role"]) == "master") {
					st = NotApplicable
				}
				add(at, node, "info_replication", signal, p+"/replication/"+key, r, st)
			}
			add(at, node, "cluster_nodes", "redis_id", p+"/redis_id", v["redis_id"], status)
			coverage := obj(v["coverage"])
			var assigned int
			json.Unmarshal(coverage["assigned"], &assigned)
			if status == Valid && len(coverage["assigned"]) > 0 {
				full := assigned == 16384
				for _, k := range []string{"missing", "duplicate", "invalid_claims", "unhealthy_owners"} {
					if len(arr(coverage[k])) > 0 {
						full = false
					}
				}
				add(at, node, "cluster_nodes", "slot_coverage", p+"/coverage", encoded(full), Valid)
			} else {
				add(at, node, "cluster_nodes", "slot_coverage", p+"/coverage", nil, status)
			}
			peers := arr(v["nodes"])
			if len(peers) > 64 {
				return fmt.Errorf("too many peer states in one view")
			}
			for j, peerRaw := range peers {
				peer := obj(peerRaw)
				peerID := txt(peer["redis_id"])
				if len(peerID) > 128 {
					return fmt.Errorf("peer identity exceeds 128 bytes")
				}
				if peerID == "" {
					continue
				}
				for _, k := range []string{"flags", "primary_id", "slots", "link_state"} {
					add(at, peerID, "peer_view:"+node, "peer."+k, fmt.Sprintf("%s/nodes/%d/%s", p, j, k), peer[k], status)
				}
			}
			persistence := obj(v["persistence"])
			for _, key := range []string{"aof_enabled", "loading", "aof_last_write_status", "aof_rewrite_in_progress"} {
				st := status
				if reachable && len(persistence[key]) == 0 {
					st = Missing
				}
				add(at, node, "info_persistence", "persistence."+key, p+"/persistence/"+key, persistence[key], st)
			}
			add(at, node, "admin", "process_live", p+"/process_live", v["process_live"], Valid)
		}
		return nil
	}
	if err := snapshot(root["initial"], "/initial"); err != nil {
		return Incident{}, err
	}
	samples := arr(root["samples"])
	if len(samples) > 4000 {
		return Incident{}, fmt.Errorf("too many snapshots")
	}
	for i, r := range samples {
		s := obj(r)
		if s == nil {
			return Incident{}, fmt.Errorf("sample must be an object")
		}
		prefix := fmt.Sprintf("/samples/%d", i)
		if err := snapshot(s["snapshot"], prefix+"/snapshot"); err != nil {
			return Incident{}, err
		}
		probe := obj(s["probe"])
		at := stamp(probe["at"])
		add(at, "", "client_probe", "client_success", prefix+"/probe/successful", probe["successful"], Valid)
		policy := obj(s["write_policy_probe"])
		if len(policy) > 0 {
			at = stamp(policy["at"])
			for _, k := range []string{"policy", "successful", "requested_replicas", "set_acknowledged", "acknowledgements", "acknowledgement_observed", "waitaof_requested_local", "waitaof_requested_replicas", "waitaof_local_acknowledgements", "waitaof_replica_acknowledgements", "waitaof_acknowledgement_observed"} {
				add(at, "", "ack_probe", "ack."+k, prefix+"/write_policy_probe/"+k, policy[k], Valid)
			}
		}
	}
	if err := snapshot(root["final"], "/final"); err != nil {
		return Incident{}, err
	}
	validation := obj(root["validation"])
	at := stamp(validation["completed_at"])
	if at.IsZero() {
		at = stamp(validation["started_at"])
	}
	for _, k := range []string{"expected", "checked", "missing", "corrupt", "unreadable", "unverified", "complete"} {
		st := Valid
		if at.IsZero() {
			st = Unavailable
		}
		add(at, "", "value_validation", "validation."+k, "/validation/"+k, validation[k], st)
	}
	windows := arr(obj(root["metrics"])["windows"])
	if len(windows) > 3600 {
		return Incident{}, fmt.Errorf("too many client windows")
	}
	for i, w := range windows {
		o := obj(w)
		at := stamp(o["ended_at"])
		for _, k := range []string{"attempts", "errors", "successes", "p95_ms"} {
			add(at, "", "client_window", "client."+k, fmt.Sprintf("/metrics/windows/%d/%s", i, k), o[k], Valid)
		}
	}
	if projectionError != nil {
		return Incident{}, projectionError
	}
	if overflow {
		return Incident{}, fmt.Errorf("too much evidence")
	}
	sort.SliceStable(out.Evidence, func(i, j int) bool { return out.Evidence[i].At.Before(out.Evidence[j].At) })
	for i := range out.Evidence {
		out.Evidence[i].ID = fmt.Sprintf("e%06d", i+1)
		if out.Evidence[i].Status == Failed {
			out.Evidence[i].CollectionError = "node collection unavailable"
		}
	}
	// The incident identity depends exclusively on the normalized observation
	// projection. Source checksum is provenance, never a diagnostic signal.
	b, _ := json.Marshal(out.Evidence)
	hash := sha256.Sum256(b)
	out.ID = hex.EncodeToString(hash[:12])
	if err := ctx.Err(); err != nil {
		return Incident{}, err
	}
	return out, nil
}

// Signals carry typed bounded values. Invalid values are retained explicitly,
// never silently interpreted as false or zero.
func validSignalValue(signal string, raw json.RawMessage) bool {
	switch signal {
	case "client.p95_ms":
		var n float64
		return json.Unmarshal(raw, &n) == nil && n >= 0
	case "reachable", "process_live", "slot_coverage", "client_success", "ack.successful", "validation.complete":
		var b bool
		return json.Unmarshal(raw, &b) == nil
	case "replication.role":
		s := txt(raw)
		return s == "master" || s == "slave"
	case "replication.master_link_status":
		s := txt(raw)
		return s == "up" || s == "down"
	case "ack.policy":
		s := txt(raw)
		return s == "set" || s == "wait" || s == "waitaof"
	case "ack.set_acknowledged", "ack.acknowledgement_observed", "ack.waitaof_acknowledgement_observed":
		var v []bool
		return json.Unmarshal(raw, &v) == nil && len(v) <= 64
	case "ack.acknowledgements", "ack.waitaof_local_acknowledgements", "ack.waitaof_replica_acknowledgements":
		var v []int64
		if json.Unmarshal(raw, &v) != nil || len(v) > 64 {
			return false
		}
		for _, n := range v {
			if n < 0 {
				return false
			}
		}
		return true
	}
	if strings.HasPrefix(signal, "validation.") || strings.HasPrefix(signal, "client.") || strings.HasPrefix(signal, "ack.requested_") || strings.HasPrefix(signal, "ack.waitaof_requested_") || signal == "replication.connected_slaves" || signal == "replication.master_sync_in_progress" || strings.HasSuffix(signal, "_offset") {
		var n int64
		return json.Unmarshal(raw, &n) == nil && n >= 0
	}
	return json.Valid(raw)
}
