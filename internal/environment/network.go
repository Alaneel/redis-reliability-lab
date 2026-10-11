package environment

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	isolationLabel    = "org.redis-reliability-lab.network-isolation"
	subnetLabel       = "org.redis-reliability-lab.isolation-subnet"
	originalIPLabel   = "org.redis-reliability-lab.original-ip"
	originalPortLabel = "org.redis-reliability-lab.original-host-port"
)

// ErrProcessStopped means ownership and resource inspection succeeded and the
// tracked Docker process was observed stopped. Other errors prove no such state.
var ErrProcessStopped = errors.New("managed Redis process is stopped")

// ProcessEvidence is an administrative observation through Docker exec. It is
// not evidence that the process is reachable by clients or cluster peers.
type ProcessEvidence struct {
	StartedAt              time.Time `json:"started_at"`
	ObservedAt             time.Time `json:"observed_at"`
	ContainerRunning       bool      `json:"container_running"`
	ContainerPID           int       `json:"container_pid"`
	ContainerStartedAt     string    `json:"container_started_at"`
	ContainerRestartPolicy string    `json:"container_restart_policy"`
	RedisPONG              bool      `json:"redis_pong"`
	RedisPID               int       `json:"redis_pid"`
	RedisRunID             string    `json:"redis_run_id"`
	RedisClusterID         string    `json:"redis_cluster_id,omitempty"`
	NetworkConnected       bool      `json:"network_connected"`
	NetworkCount           int       `json:"network_count"`
	RawClusterInfo         string    `json:"raw_cluster_info"`
	RawClusterNodes        string    `json:"raw_cluster_nodes"`
	RawReplication         string    `json:"raw_replication"`
	RawPersistence         string    `json:"raw_persistence,omitempty"`
	RawPersistenceConfig   string    `json:"raw_persistence_config,omitempty"`
	RawKeyspace            string    `json:"raw_keyspace,omitempty"`
	RawDataFiles           string    `json:"raw_data_files,omitempty"`
	DataSizeKiB            int64     `json:"data_size_kib,omitempty"`
}

// IsolationReceipt bounds restoration to the original owned resources and
// process. It is returned before disconnect so an ambiguous command error can
// still be followed by safe restoration using a fresh bounded context.
type IsolationReceipt struct {
	RunID              string          `json:"run_id"`
	NetworkID          string          `json:"network_id"`
	Node               Node            `json:"node"`
	OriginalIP         string          `json:"original_ip"`
	HostAddr           string          `json:"host_addr"`
	BeforeProcess      ProcessEvidence `json:"before_process"`
	CommandStartedAt   time.Time       `json:"command_started_at"`
	CommandCompletedAt time.Time       `json:"command_completed_at"`
	DisconnectedAt     time.Time       `json:"disconnected_at,omitempty"`
	RestoredAt         time.Time       `json:"restored_at,omitempty"`
}

func isolationSubnet(runID string, attempt int) string {
	hash := sha256.Sum256([]byte(runID))
	index := (int(hash[0])<<8 | int(hash[1])) + attempt
	return fmt.Sprintf("10.%d.%d.0/24", 224+(index/256)%16, index%256)
}

func validIsolationSubnet(subnet string) bool {
	ip, pool, err := net.ParseCIDR(subnet)
	if err != nil || ip.To4() == nil || subnet != pool.String() {
		return false
	}
	bits, width := pool.Mask.Size()
	v4 := ip.To4()
	return bits == 24 && width == 32 && v4[0] == 10 && v4[1] >= 224 && v4[1] <= 239
}

func (m *Manager) createRunNetwork(ctx context.Context, runID string, isolation, retained bool, persistence string) ([]byte, string, error) {
	for attempt := 0; attempt < 32; attempt++ {
		args := []string{"network", "create", "--driver", "bridge", "--opt", noMasqueradeOption + "=false", "--label", managedLabel + "=true", "--label", runLabel + "=" + runID}
		subnet := ""
		if isolation || retained {
			subnet = isolationSubnet(runID, attempt)
			args = append(args, "--subnet", subnet, "--label", subnetLabel+"="+subnet)
		}
		if isolation {
			args = append(args, "--label", isolationLabel+"=true")
		}
		if retained {
			args = append(args, "--label", retainedLabel+"=true", "--label", persistenceLabel+"="+persistence)
		}
		out, err := m.command(ctx, 30*time.Second, append(args, resourceName(runID))...)
		if err == nil || !(isolation || retained) || !strings.Contains(strings.ToLower(err.Error()), "pool overlaps") || ctx.Err() != nil {
			return out, subnet, err
		}
		// Only Docker's definite IPAM overlap rejection is retried. Other errors
		// may have created a resource and are reconciled by Provision discovery.
	}
	return nil, "", errors.New("no free managed private isolation subnet after 32 candidates")
}

func chooseLoopbackPort(used map[string]bool) (string, error) {
	for attempt := 0; attempt < 32; attempt++ {
		listener, err := net.Listen("tcp4", "127.0.0.1:0")
		if err != nil {
			return "", fmt.Errorf("select loopback port: %w", err)
		}
		port := strconv.Itoa(listener.Addr().(*net.TCPAddr).Port)
		closeErr := listener.Close()
		if closeErr != nil {
			return "", closeErr
		}
		if !used[port] {
			// Docker binds this exact port. A competing bind in this unavoidable
			// interval causes an explicit provision failure and owned cleanup.
			return port, nil
		}
	}
	return "", errors.New("could not select a unique ephemeral loopback port")
}

func validateIsolationNetwork(network networkInspect) error {
	subnet := network.Labels[subnetLabel]
	if network.Internal || network.Options[noMasqueradeOption] != "false" || !validIsolationSubnet(subnet) || len(network.IPAM.Config) != 1 || network.IPAM.Config[0].Subnet != subnet {
		return errors.New("refusing operation: isolation bridge has unexpected IPAM or egress configuration")
	}
	return nil
}

func validateOriginalEndpoints(container containerInspect, node Node) error {
	ip := container.Config.Labels[originalIPLabel]
	port := container.Config.Labels[originalPortLabel]
	parsedIP := net.ParseIP(ip)
	n, err := strconv.Atoi(port)
	if parsedIP == nil || parsedIP.To4() == nil || !parsedIP.IsPrivate() || err != nil || n < 1024 || n > 65535 || strconv.Itoa(n) != port {
		return errors.New("refusing operation: invalid original isolation endpoints")
	}
	bindings := container.HostConfig.PortBindings["6379/tcp"]
	if len(bindings) != 1 || bindings[0].HostIP != "127.0.0.1" || bindings[0].HostPort != port {
		return errors.New("refusing operation: configured isolation port changed")
	}
	if node.InternalAddr != "" && node.InternalAddr != net.JoinHostPort(ip, "6379") || node.Addr != "" && node.Addr != net.JoinHostPort("127.0.0.1", port) {
		return errors.New("refusing operation: tracked isolation endpoints changed")
	}
	for _, member := range container.NetworkSettings.Networks {
		if member.IPAddress != "" && member.IPAddress != ip {
			return errors.New("refusing operation: isolation node changed its original IP")
		}
	}
	for _, binding := range container.NetworkSettings.Ports["6379/tcp"] {
		if binding.HostIP != "127.0.0.1" || binding.HostPort != port {
			return errors.New("refusing operation: isolation node changed its original published port")
		}
	}
	return nil
}

func (m *Manager) isolationNode(ctx context.Context, env *Environment, node Node) (containerInspect, networkInspect, error) {
	var container containerInspect
	var network networkInspect
	if env == nil || !env.NetworkIsolation {
		return container, network, errors.New("environment was not provisioned for network isolation")
	}
	if err := validateRun(env.RunID); err != nil {
		return container, network, err
	}
	tracked := false
	for _, candidate := range env.Nodes {
		if candidate == node {
			tracked = true
		}
	}
	if !tracked || node.InternalAddr == "" || node.Addr == "" {
		return container, network, errors.New("refusing operation on an untracked isolation node")
	}
	var err error
	network, err = m.inspectNetwork(ctx, env.NetworkID)
	if err != nil {
		return container, network, err
	}
	if err = validateNetwork(network, env); err != nil {
		return container, network, err
	}
	container, err = m.inspectContainer(ctx, node.ContainerID)
	if err != nil {
		return container, network, err
	}
	if container.Config.Labels[isolationLabel] != "true" {
		return container, network, errors.New("node lacks managed network isolation capability")
	}
	if err = validateContainer(container, env, node, false); err != nil {
		return container, network, err
	}
	if err = m.validateNodeVolume(ctx, env, node); err != nil {
		return container, network, err
	}
	_, subnet, _ := net.ParseCIDR(network.Labels[subnetLabel])
	if subnet == nil || !subnet.Contains(net.ParseIP(container.Config.Labels[originalIPLabel])) {
		return container, network, errors.New("isolation node IP is outside its managed subnet")
	}
	return container, network, nil
}

// ProbeProcess observes a tracked process even when its only network interface
// has been disconnected. All administrative Redis commands are fixed and local.
func (m *Manager) ProbeProcess(ctx context.Context, env *Environment, node Node) (evidence ProcessEvidence, err error) {
	evidence.StartedAt = time.Now()
	defer func() { evidence.ObservedAt = time.Now() }()
	probeCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	container, _, err := m.processNode(probeCtx, env, node)
	if err != nil {
		return evidence, err
	}
	evidence.ContainerRunning = container.State.Running
	evidence.ContainerPID = container.State.Pid
	evidence.ContainerStartedAt = container.State.StartedAt
	evidence.ContainerRestartPolicy = container.HostConfig.RestartPolicy.Name
	evidence.NetworkCount = len(container.NetworkSettings.Networks)
	evidence.NetworkConnected = container.State.Running && evidence.NetworkCount == 1
	if !container.State.Running {
		return evidence, ErrProcessStopped
	}
	cli := func(command ...string) (string, error) {
		args := append([]string{"exec", node.ContainerID, "redis-cli", "--raw", "-h", "127.0.0.1"}, command...)
		out, callErr := m.command(probeCtx, 3*time.Second, args...)
		return strings.TrimSpace(string(out)), callErr
	}
	ping, pingErr := cli("ping")
	evidence.RedisPONG = pingErr == nil && ping == "PONG"
	if !evidence.RedisPONG {
		return evidence, errors.Join(errors.New("managed Redis process did not return local PONG"), pingErr)
	}
	server, serverErr := cli("info", "server")
	if serverErr != nil {
		return evidence, serverErr
	}
	for _, line := range strings.Split(server, "\n") {
		key, value, _ := strings.Cut(strings.TrimSpace(line), ":")
		switch key {
		case "run_id":
			evidence.RedisRunID = value
		case "process_id":
			evidence.RedisPID, _ = strconv.Atoi(value)
		}
	}
	if len(evidence.RedisRunID) != 40 || !isHex(evidence.RedisRunID) || evidence.RedisPID < 1 || evidence.ContainerPID < 1 || evidence.ContainerStartedAt == "" {
		return evidence, errors.New("managed process lacks a valid Redis or Docker process identity")
	}
	evidence.RawClusterInfo, err = cli("cluster", "info")
	if err != nil {
		return evidence, err
	}
	evidence.RawClusterNodes, err = cli("cluster", "nodes")
	if err != nil {
		return evidence, err
	}
	evidence.RawReplication, err = cli("info", "replication")
	if err != nil {
		return evidence, err
	}
	if env.RetainData {
		evidence.RedisClusterID, err = cli("cluster", "myid")
		if err != nil || len(evidence.RedisClusterID) != 40 || !isHex(evidence.RedisClusterID) {
			return evidence, errors.Join(errors.New("managed Redis process lacks a valid Cluster identity"), err)
		}
		evidence.RawPersistence, err = cli("info", "persistence")
		if err != nil {
			return evidence, err
		}
		configArgs := []string{"exec", node.ContainerID, "redis-cli", "--raw", "-h", "127.0.0.1", "config", "get", "appendonly", "appendfsync", "save", "repl-diskless-load", "repl-diskless-sync", "auto-aof-rewrite-percentage", "maxmemory", "maxmemory-policy", "aof-use-rdb-preamble", "no-appendfsync-on-rewrite", "aof-load-truncated", "rdb-del-sync-files"}
		configReply, configErr := m.command(probeCtx, 3*time.Second, configArgs...)
		err = configErr
		evidence.RawPersistenceConfig = persistenceConfigReply(configReply)
		if err != nil {
			return evidence, err
		}
		if err = validatePersistenceConfig(evidence.RawPersistenceConfig, env.Persistence); err != nil {
			return evidence, err
		}
		evidence.RawKeyspace, err = cli("info", "keyspace")
		if err != nil {
			return evidence, err
		}
		if err = validateKeyspace(evidence.RawKeyspace); err != nil {
			return evidence, err
		}
		files, filesErr := m.command(probeCtx, 3*time.Second, "exec", node.ContainerID, "find", "/data", "-maxdepth", "2", "-type", "f")
		evidence.RawDataFiles = string(files)
		if filesErr != nil {
			return evidence, filesErr
		}
		if err = validateDataFiles(evidence.RawDataFiles); err != nil {
			return evidence, err
		}
		size, sizeErr := m.command(probeCtx, 3*time.Second, "exec", node.ContainerID, "du", "-sk", "/data")
		fields := strings.Fields(string(size))
		if sizeErr != nil || len(fields) != 2 || fields[1] != "/data" {
			return evidence, errors.Join(errors.New("observe managed data usage"), sizeErr)
		}
		evidence.DataSizeKiB, err = strconv.ParseInt(fields[0], 10, 64)
		if err != nil || evidence.DataSizeKiB < 0 {
			return evidence, errors.New("invalid managed data usage")
		}
	}
	current, _, checkErr := m.processNode(probeCtx, env, node)
	if checkErr != nil {
		return evidence, checkErr
	}
	if !current.State.Running || current.State.Pid != container.State.Pid || current.State.StartedAt != container.State.StartedAt {
		return evidence, errors.New("managed process changed during administrative observation")
	}
	if len(current.NetworkSettings.Networks) != evidence.NetworkCount {
		return evidence, errors.New("managed network attachment changed during administrative observation")
	}
	return evidence, nil
}

func isHex(value string) bool {
	for _, c := range value {
		if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'f') {
			return false
		}
	}
	return true
}

func sameProcess(before, after ProcessEvidence) bool {
	return before.ContainerRunning && after.ContainerRunning && before.RedisPONG && after.RedisPONG && before.ContainerPID == after.ContainerPID && before.ContainerStartedAt == after.ContainerStartedAt && before.RedisPID == after.RedisPID && before.RedisRunID == after.RedisRunID && before.RedisRunID != ""
}

// Isolate disconnects the sole managed bridge endpoint. This is a whole-node,
// bidirectional fault affecting peer, replication, and host-published traffic.
func (m *Manager) Isolate(ctx context.Context, env *Environment, node Node) (*IsolationReceipt, error) {
	container, _, err := m.isolationNode(ctx, env, node)
	if err != nil {
		return nil, err
	}
	receipt := &IsolationReceipt{RunID: env.RunID, NetworkID: env.NetworkID, Node: node, OriginalIP: container.Config.Labels[originalIPLabel], HostAddr: node.Addr}
	if err = validateContainer(container, env, node, true); err != nil {
		return receipt, err
	}
	receipt.BeforeProcess, err = m.ProbeProcess(ctx, env, node)
	if err != nil {
		return receipt, err
	}
	if !receipt.BeforeProcess.NetworkConnected || receipt.BeforeProcess.NetworkCount != 1 {
		return receipt, errors.New("node disconnected before fault injection")
	}
	current, network, err := m.isolationNode(ctx, env, node)
	if err != nil {
		return receipt, err
	}
	if err = validateContainer(current, env, node, true); err != nil {
		return receipt, err
	}
	if current.State.Pid != receipt.BeforeProcess.ContainerPID || current.State.StartedAt != receipt.BeforeProcess.ContainerStartedAt {
		return receipt, errors.New("managed process changed before fault injection")
	}
	if _, member := network.Containers[node.ContainerID]; !member {
		return receipt, errors.New("managed network does not report target membership before isolation")
	}
	receipt.CommandStartedAt = time.Now()
	_, commandErr := m.command(ctx, 15*time.Second, "network", "disconnect", env.NetworkID, node.ContainerID)
	receipt.CommandCompletedAt = time.Now()
	// Inspect on an independent bounded context even if the mutation's context
	// was canceled, because an error does not imply that Docker did not detach.
	checkCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	after, network, checkErr := m.isolationNode(checkCtx, env, node)
	if checkErr == nil && len(after.NetworkSettings.Networks) == 0 {
		if _, member := network.Containers[node.ContainerID]; member {
			checkErr = errors.New("network still reports isolated node membership")
		} else {
			receipt.DisconnectedAt = time.Now()
		}
	} else if checkErr == nil {
		checkErr = errors.New("node remains attached after network disconnect")
	}
	if commandErr != nil || checkErr != nil {
		return receipt, errors.Join(commandErr, checkErr)
	}
	process, processErr := m.ProbeProcess(checkCtx, env, node)
	if processErr != nil || !sameProcess(receipt.BeforeProcess, process) {
		return receipt, errors.Join(errors.New("Redis process did not survive network isolation unchanged"), processErr)
	}
	return receipt, nil
}

// Restore reattaches the same process to its original private IP and published
// port. It never recreates resources or disconnects unexpected network members.
func (m *Manager) Restore(ctx context.Context, env *Environment, receipt *IsolationReceipt) error {
	if receipt == nil || env == nil || receipt.RunID != env.RunID || receipt.NetworkID != env.NetworkID || receipt.HostAddr != receipt.Node.Addr || net.JoinHostPort(receipt.OriginalIP, "6379") != receipt.Node.InternalAddr {
		return errors.New("invalid or mismatched isolation receipt")
	}
	container, network, err := m.isolationNode(ctx, env, receipt.Node)
	if err != nil {
		return err
	}
	if container.Config.Labels[originalIPLabel] != receipt.OriginalIP {
		return errors.New("isolation receipt IP does not match the provisioned endpoint")
	}
	process, err := m.ProbeProcess(ctx, env, receipt.Node)
	if err != nil || !sameProcess(receipt.BeforeProcess, process) {
		return errors.Join(errors.New("refusing restore: original Redis process changed or is unavailable"), err)
	}
	if len(container.NetworkSettings.Networks) == 1 {
		if err = validateContainer(container, env, receipt.Node, true); err != nil {
			return err
		}
		if _, member := network.Containers[receipt.Node.ContainerID]; !member {
			return errors.New("network membership disagrees with the attached original endpoint")
		}
		receipt.RestoredAt = time.Now()
		return nil
	}
	if _, member := network.Containers[receipt.Node.ContainerID]; member {
		return errors.New("refusing restore: network still reports a detached target endpoint")
	}
	for id, raw := range network.Containers {
		var member struct{ IPv4Address string }
		if decodeErr := json.Unmarshal(raw, &member); decodeErr != nil {
			return fmt.Errorf("decode managed network member: %w", decodeErr)
		}
		ip, _, _ := strings.Cut(member.IPv4Address, "/")
		if id != receipt.Node.ContainerID && ip == receipt.OriginalIP {
			return errors.New("refusing restore: original IP is occupied by another endpoint")
		}
	}
	_, commandErr := m.command(ctx, 15*time.Second, "network", "connect", "--ip", receipt.OriginalIP, env.NetworkID, receipt.Node.ContainerID)
	checkCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	after, currentNetwork, checkErr := m.isolationNode(checkCtx, env, receipt.Node)
	if checkErr == nil {
		checkErr = validateContainer(after, env, receipt.Node, true)
	}
	if checkErr == nil {
		if _, member := currentNetwork.Containers[receipt.Node.ContainerID]; !member {
			checkErr = errors.New("managed network does not report restored node membership")
		}
	}
	if checkErr == nil {
		current, processErr := m.ProbeProcess(checkCtx, env, receipt.Node)
		if processErr != nil || !sameProcess(receipt.BeforeProcess, current) {
			checkErr = errors.Join(errors.New("Redis process changed during network restoration"), processErr)
		}
	}
	if checkErr == nil {
		receipt.RestoredAt = time.Now()
	}
	return errors.Join(commandErr, checkErr)
}
