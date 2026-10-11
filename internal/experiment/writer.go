package experiment

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
)

// ManagedWriter owns its write pools, while borrowing cluster topology only.
// SET and WAIT/WAITAOF use one sticky physical connection. Transport failures are never
// replayed; only an explicit allowlisted MOVED/ASK rejection permits another SET.
type ManagedWriter struct {
	cluster     *redis.ClusterClient
	nodes       map[string]*redis.Client
	aliases     map[string]string
	timeout     time.Duration
	policy      string
	replicas    int
	waitTimeout time.Duration
	mu          sync.RWMutex
	closed      bool
	closeErr    error
}

var _ workload.WriteExecutor = (*ManagedWriter)(nil)

// NewWriter validates a bounded acknowledgement policy and the managed address
// inventory. Closing it cannot close the supplied cluster client or its pools.
func NewWriter(clusterClient *redis.ClusterClient, env *environment.Environment, timeout time.Duration, policy string, replicas int, waitTimeout time.Duration) (*ManagedWriter, error) {
	if clusterClient == nil || env == nil || len(env.Nodes) != 6 || timeout <= 0 {
		return nil, errors.New("writer requires a cluster client, six managed nodes, and a positive operation timeout")
	}
	if policy == "set" {
		// Shared configuration defaults may include WAIT parameters. They have
		// no effect under SET and never cause a WAIT command to be issued.
	} else if policy == "wait" || policy == "waitaof" {
		if replicas != 1 || waitTimeout < time.Millisecond || waitTimeout > time.Second || waitTimeout%time.Millisecond != 0 || timeout <= waitTimeout {
			return nil, errors.New("WAIT/WAITAOF requires one replica, integral positive milliseconds, and a larger operation timeout")
		}
		if policy == "waitaof" && (!env.RetainData || env.Persistence != "aof-everysec" && env.Persistence != "aof-always") {
			return nil, errors.New("WAITAOF requires retained AOF persistence on all managed nodes")
		}
	} else {
		return nil, errors.New("write policy must be set, wait or waitaof")
	}
	w := &ManagedWriter{cluster: clusterClient, nodes: map[string]*redis.Client{}, aliases: map[string]string{}, timeout: timeout, policy: policy, replicas: replicas, waitTimeout: waitTimeout}
	ids, endpoints := map[string]bool{}, map[string]bool{}
	for _, node := range env.Nodes {
		host, port, err := net.SplitHostPort(node.Addr)
		portNumber, parseErr := strconv.Atoi(port)
		if err != nil || host != "127.0.0.1" || parseErr != nil || portNumber < 1 || portNumber > 65535 || endpoints[node.Addr] || node.ID == "" || ids[node.ID] {
			return nil, errors.New("writer requires unique managed identities and numeric loopback ports")
		}
		ip, redisPort, internalErr := net.SplitHostPort(node.InternalAddr)
		parsedIP := net.ParseIP(ip)
		if internalErr != nil || parsedIP == nil || !parsedIP.IsPrivate() || redisPort != "6379" || w.aliases[node.InternalAddr] != "" {
			return nil, errors.New("writer requires unique private managed Redis endpoints")
		}
		ids[node.ID], endpoints[node.Addr] = true, true
		w.aliases[node.InternalAddr], w.aliases[node.Addr] = node.InternalAddr, node.InternalAddr
	}
	for _, seed := range clusterClient.Options().Addrs {
		if w.aliases[seed] == "" {
			return nil, fmt.Errorf("writer refuses unmanaged cluster seed %q", seed)
		}
	}
	for _, node := range env.Nodes {
		canonical, published := node.InternalAddr, node.Addr
		dialer := net.Dialer{Timeout: timeout}
		w.nodes[canonical] = redis.NewClient(&redis.Options{
			Addr: canonical, Protocol: 2, MaxRetries: -1, DialerRetries: 1,
			DialTimeout: timeout, ReadTimeout: timeout, WriteTimeout: timeout, PoolTimeout: timeout,
			ContextTimeoutEnabled: true, PoolSize: 16, DisableIdentity: true,
			MaintNotificationsConfig: &maintnotifications.Config{Mode: maintnotifications.ModeDisabled},
			Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
				if network != "tcp" || addr != canonical {
					return nil, fmt.Errorf("writer refuses unmanaged dial target %q", addr)
				}
				return dialer.DialContext(ctx, "tcp", published)
			},
		})
	}
	return w, nil
}

// Write keeps every successful SET in the acknowledgement result, including
// when the following WAIT/WAITAOF is short, fails, or cannot start before cancellation.
func (w *ManagedWriter) Write(parent context.Context, key, value string) (result workload.WriteResult) {
	started := time.Now()
	defer func() { result.TotalLatency = time.Since(started) }()
	w.mu.RLock()
	defer w.mu.RUnlock()
	if w.closed {
		result.SetError = errors.New("managed writer is closed")
		return result
	}
	if w.policy == "wait" {
		result.WaitRequested = w.replicas
	} else if w.policy == "waitaof" {
		result.WaitAOFRequestedLocal, result.WaitAOFRequestedReplicas = 1, w.replicas
	}
	ctx, cancel := context.WithTimeout(parent, w.timeout)
	defer cancel()
	setStarted := time.Now()
	finishSet := func() { result.SetLatency = time.Since(setStarted) }
	master, err := w.cluster.MasterForKey(ctx, key)
	if err != nil {
		result.SetError = err
		finishSet()
		return result
	}
	canonical := w.aliases[master.Options().Addr]
	if canonical == "" {
		result.SetError = fmt.Errorf("writer refuses unmanaged slot owner %q", master.Options().Addr)
		finishSet()
		return result
	}
	asking := false
	for {
		if err := ctx.Err(); err != nil {
			result.SetError = err
			finishSet()
			return result
		}
		client := w.nodes[canonical]
		var conn *redis.Conn
		if asking {
			// ASKING changes connection session state. Discard this socket on
			// release, including cancellation before SET consumes that state.
			conn = client.EphemeralConn()
			status, err := conn.Do(ctx, "ASKING").Text()
			if err == nil && status != "OK" {
				err = fmt.Errorf("unexpected ASKING acknowledgement %q", status)
			}
			if err != nil {
				result.SetError = err
				_ = conn.Close()
				finishSet()
				return result
			}
		} else {
			conn = client.Conn()
		}
		status, setErr := conn.Set(ctx, key, value, 0).Result()
		if setErr == nil && status != "OK" {
			setErr = fmt.Errorf("unexpected SET acknowledgement %q", status)
		}
		if setErr != nil {
			_ = conn.Close()
			kind, target, redirected := managedRedirect(setErr, key, w.aliases)
			if !redirected || result.Redirects >= 8 {
				result.SetError = setErr
				finishSet()
				return result
			}
			result.Redirects++
			canonical, asking = target, kind == "ASK"
			if kind == "MOVED" {
				w.cluster.ReloadState(ctx)
			}
			continue
		}
		result.SetAcknowledged = true
		result.SetAcknowledgedAt = time.Now()
		finishSet()
		if w.policy == "set" {
			result.PolicySatisfied = true
			_ = conn.Close()
			return result
		}
		if err := ctx.Err(); err != nil {
			if w.policy == "waitaof" {
				result.WaitAOFError = err
			} else {
				result.WaitError = err
			}
			_ = conn.Close()
			return result
		}
		waitStarted := time.Now()
		if w.policy == "waitaof" {
			result.WaitAOFAttempted = true
			var reply []int64
			reply, result.WaitAOFError = conn.WaitAOF(ctx, 1, w.replicas, w.waitTimeout).Result()
			result.WaitAOFLatency = time.Since(waitStarted)
			if result.WaitAOFError == nil {
				if len(reply) != 2 || reply[0] < 0 || reply[0] > 1 || reply[1] < 0 {
					result.WaitAOFError = errors.New("invalid WAITAOF local/replica acknowledgement reply")
				} else {
					result.WaitAOFLocalAcknowledgements, result.WaitAOFReplicaAcknowledgements = reply[0], reply[1]
					result.PolicySatisfied = reply[0] >= 1 && reply[1] >= int64(w.replicas)
				}
			}
			_ = conn.Close()
			return result
		}
		result.WaitAttempted = true
		result.WaitAcknowledgements, result.WaitError = conn.Wait(ctx, w.replicas, w.waitTimeout).Result()
		result.WaitLatency = time.Since(waitStarted)
		result.PolicySatisfied = result.WaitError == nil && result.WaitAcknowledgements >= int64(w.replicas)
		_ = conn.Close()
		return result
	}
}

func managedRedirect(err error, key string, aliases map[string]string) (kind, target string, ok bool) {
	var serverErr redis.Error
	if !errors.As(err, &serverErr) {
		return "", "", false
	}
	fields := strings.Fields(serverErr.Error())
	if len(fields) != 3 || fields[0] != "MOVED" && fields[0] != "ASK" {
		return "", "", false
	}
	slot, parseErr := strconv.Atoi(fields[1])
	if parseErr != nil || slot != int(workload.HashSlot(key)) || aliases[fields[2]] == "" {
		return "", "", false
	}
	return fields[0], aliases[fields[2]], true
}

func (w *ManagedWriter) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.closed {
		return w.closeErr
	}
	w.closed = true
	var failures []error
	for _, client := range w.nodes {
		if err := client.Close(); err != nil {
			failures = append(failures, err)
		}
	}
	w.closeErr = errors.Join(failures...)
	return w.closeErr
}
