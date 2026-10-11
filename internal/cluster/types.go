// Package cluster observes the managed Redis nodes without relying on the
// cluster client's topology cache. Every observation retains the source replies.
package cluster

import "time"

const SlotCount = 16384

// Node identifies a managed container and its host and cluster-network endpoints.
// ID is a managed environment ID, not a Redis Cluster node ID.
type Node struct {
	ID           string `json:"id"`
	Addr         string `json:"addr"`
	InternalAddr string `json:"internal_addr"`
}

type SlotRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type OwnerRange struct {
	Start   int    `json:"start"`
	End     int    `json:"end"`
	RedisID string `json:"redis_id"`
}

type NodeState struct {
	ID              string      `json:"redis_id"`
	Addr            string      `json:"addr"`
	Flags           []string    `json:"flags"`
	PrimaryID       string      `json:"primary_id,omitempty"`
	PingSent        int64       `json:"ping_sent"`
	PongReceived    int64       `json:"pong_received"`
	ConfigEpoch     uint64      `json:"config_epoch"`
	LinkState       string      `json:"link_state"`
	Slots           []SlotRange `json:"slots,omitempty"`
	SlotTransitions []string    `json:"slot_transitions,omitempty"`
}

func (n NodeState) HasFlag(flag string) bool {
	for _, f := range n.Flags {
		if f == flag {
			return true
		}
	}
	return false
}

func (n NodeState) Primary() bool { return n.HasFlag("master") }
func (n NodeState) Replica() bool { return n.HasFlag("slave") }
func (n NodeState) Healthy() bool {
	return n.LinkState == "connected" && !n.HasFlag("fail") && !n.HasFlag("fail?") &&
		!n.HasFlag("handshake") && !n.HasFlag("noaddr")
}

type Coverage struct {
	Assigned        int          `json:"assigned"`
	InvalidClaims   []string     `json:"invalid_claims,omitempty"`
	Missing         []SlotRange  `json:"missing,omitempty"`
	Duplicate       []SlotRange  `json:"duplicate,omitempty"`
	UnhealthyOwners []SlotRange  `json:"unhealthy_owners,omitempty"`
	Owners          []OwnerRange `json:"owners"`
}

func (c Coverage) Full() bool {
	return c.Assigned == SlotCount && len(c.InvalidClaims) == 0 && len(c.Missing) == 0 && len(c.Duplicate) == 0 && len(c.UnhealthyOwners) == 0
}

type View struct {
	Node               Node                `json:"managed_node"`
	At                 time.Time           `json:"at"`
	RedisID            string              `json:"redis_id,omitempty"`
	Reachable          bool                `json:"reachable"`
	Errors             []string            `json:"errors,omitempty"`
	ClusterInfo        map[string]string   `json:"cluster_info,omitempty"`
	Replication        map[string]string   `json:"replication,omitempty"`
	Persistence        map[string]string   `json:"persistence,omitempty"`
	Stats              map[string]string   `json:"stats,omitempty"`
	Server             map[string]string   `json:"server,omitempty"`
	ReplicationMetrics *ReplicationMetrics `json:"replication_metrics,omitempty"`
	RedisVersion       string              `json:"redis_version,omitempty"`
	Nodes              []NodeState         `json:"nodes,omitempty"`
	Coverage           Coverage            `json:"coverage"`
	RawClusterInfo     string              `json:"raw_cluster_info,omitempty"`
	RawNodes           string              `json:"raw_cluster_nodes,omitempty"`
	RawReplication     string              `json:"raw_replication,omitempty"`
	RawServer          string              `json:"raw_server,omitempty"`
	RawPersistence     string              `json:"raw_persistence,omitempty"`
	RawStats           string              `json:"raw_stats,omitempty"`
}

type Snapshot struct {
	At        time.Time `json:"at"`
	Completed time.Time `json:"completed"`
	Views     []View    `json:"views"`
}

type Selection struct {
	Target    Node   `json:"target"`
	TargetID  string `json:"target_redis_id"`
	Replica   Node   `json:"replica"`
	ReplicaID string `json:"replica_redis_id"`
	Slots     []int  `json:"slots"`
}
