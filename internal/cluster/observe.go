package cluster

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type Observer struct {
	nodes   []Node
	clients []*redis.Client
	timeout time.Duration
}

func New(nodes []Node, timeout time.Duration) *Observer {
	if timeout <= 0 {
		timeout = time.Second
	}
	o := &Observer{nodes: append([]Node(nil), nodes...), timeout: timeout}
	for _, node := range nodes {
		o.clients = append(o.clients, redis.NewClient(&redis.Options{
			Addr: node.Addr, DialTimeout: timeout, ReadTimeout: timeout, WriteTimeout: timeout,
			Protocol: 2, DialerRetries: 1, ContextTimeoutEnabled: true, MaxRetries: -1, PoolSize: 1,
		}))
	}
	return o
}

// Observe returns failures as evidence, including for the deliberately killed
// node. Observation work is bounded by the caller and a per-node timeout.
func (o *Observer) Observe(ctx context.Context) Snapshot {
	s := Snapshot{At: time.Now(), Views: make([]View, len(o.nodes))}
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 6)
	for i := range o.nodes {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			select {
			case semaphore <- struct{}{}:
				defer func() { <-semaphore }()
			case <-ctx.Done():
				s.Views[i] = View{Node: o.nodes[i], At: time.Now(), Errors: []string{ctx.Err().Error()}}
				return
			}
			s.Views[i] = o.observeNode(ctx, i)
		}(i)
	}
	wg.Wait()
	s.Completed = time.Now()
	return s
}

func (o *Observer) observeNode(parent context.Context, i int) View {
	ctx, cancel := context.WithTimeout(parent, o.timeout)
	defer cancel()
	v := View{Node: o.nodes[i], At: time.Now()}
	var info, nodes, replication, server, persistence, stats *redis.StringCmd
	_, pipelineErr := o.clients[i].Pipelined(ctx, func(p redis.Pipeliner) error {
		info = p.ClusterInfo(ctx)
		nodes = p.ClusterNodes(ctx)
		replication = p.Info(ctx, "replication")
		server = p.Info(ctx, "server")
		persistence = p.Info(ctx, "persistence")
		stats = p.Info(ctx, "stats")
		return nil
	})
	for _, response := range []struct {
		name string
		cmd  *redis.StringCmd
		dst  *string
	}{
		{"CLUSTER INFO", info, &v.RawClusterInfo}, {"CLUSTER NODES", nodes, &v.RawNodes},
		{"INFO replication", replication, &v.RawReplication}, {"INFO server", server, &v.RawServer},
		{"INFO persistence", persistence, &v.RawPersistence}, {"INFO stats", stats, &v.RawStats},
	} {
		if response.cmd == nil {
			v.Errors = append(v.Errors, response.name+": command not issued")
			continue
		}
		value, err := response.cmd.Result()
		if err != nil {
			v.Errors = append(v.Errors, response.name+": "+err.Error())
			continue
		}
		v.Reachable = true
		*response.dst = value
	}
	if pipelineErr != nil && len(v.Errors) == 0 {
		v.Errors = append(v.Errors, "observation pipeline: "+pipelineErr.Error())
	}
	v.ClusterInfo = ParseInfo(v.RawClusterInfo)
	v.Replication = ParseInfo(v.RawReplication)
	v.Persistence = ParseInfo(v.RawPersistence)
	v.Stats = ParseInfo(v.RawStats)
	v.Server = ParseInfo(v.RawServer)
	v.RedisVersion = v.Server["redis_version"]
	derived := DeriveReplication(v.Replication)
	v.ReplicationMetrics = &derived
	if v.RawNodes != "" {
		parsed, err := ParseNodes(v.RawNodes)
		if err != nil {
			v.Errors = append(v.Errors, err.Error())
		} else {
			v.Nodes = parsed
			for _, node := range parsed {
				if node.HasFlag("myself") {
					if v.RedisID != "" {
						v.Errors = append(v.Errors, "CLUSTER NODES contains multiple myself nodes")
					}
					v.RedisID = node.ID
				}
			}
			if v.RedisID == "" {
				v.Errors = append(v.Errors, "CLUSTER NODES contains no myself node")
			}
			v.Coverage = AnalyzeCoverage(parsed)
		}
	}
	return v
}

func (o *Observer) Close() error {
	var errs []error
	for i, client := range o.clients {
		if err := client.Close(); err != nil {
			errs = append(errs, fmt.Errorf("close observer %s: %w", o.nodes[i].ID, err))
		}
	}
	return errors.Join(errs...)
}
