// Package environment provisions isolated, explicitly labelled local Redis containers.
package environment

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	DefaultImage       = "redis:8.10.2-alpine"
	managedLabel       = "org.redis-reliability-lab.managed"
	runLabel           = "org.redis-reliability-lab.run-id"
	nodeCount          = 6
	noMasqueradeOption = "com.docker.network.bridge.enable_ip_masquerade"
)

var (
	runPattern   = regexp.MustCompile(`^[a-z0-9][a-z0-9-]{7,63}$`)
	idPattern    = regexp.MustCompile(`^[a-f0-9]{64}$`)
	imagePattern = regexp.MustCompile(`^redis:[0-9]+\.[0-9]+\.[0-9]+-alpine(?:@sha256:[a-f0-9]{64})?$`)
)

// Config pins the official Redis image and Redis failure detection timeout.
type Config struct {
	Image            string        `json:"image"`
	NodeTimeout      time.Duration `json:"node_timeout"`
	NetworkIsolation bool          `json:"network_isolation"`
	Persistence      string        `json:"persistence"`
	RetainData       bool          `json:"retain_data"`
}

// Node holds both the Redis network address and the host's loopback address.
// ID is a local node key; Redis cluster IDs are observed separately.
type Node struct {
	ID             string `json:"id"`
	ContainerID    string `json:"container_id"`
	Name           string `json:"name"`
	InternalAddr   string `json:"internal_addr"`
	Addr           string `json:"addr"`
	RedisClusterID string `json:"redis_cluster_id,omitempty"`
}

type Environment struct {
	RunID            string            `json:"run_id"`
	NetworkID        string            `json:"network_id"`
	Nodes            []Node            `json:"nodes"`
	Metadata         map[string]string `json:"metadata"`
	NetworkIsolation bool              `json:"network_isolation,omitempty"`
	RetainData       bool              `json:"retain_data,omitempty"`
	Persistence      string            `json:"persistence,omitempty"`
	Volumes          []Volume          `json:"volumes,omitempty"`
}

// Manager uses Docker's local Unix socket. Every command after endpoint discovery
// specifies that socket explicitly, so context changes cannot retarget a run.
type Manager struct {
	execute    func(context.Context, ...string) ([]byte, error)
	getenv     func(string) string
	mu         sync.Mutex
	endpoint   string
	choosePort func(map[string]bool) (string, error)
}

func New() *Manager {
	return &Manager{execute: executeDocker, getenv: os.Getenv}
}

func executeDocker(ctx context.Context, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "docker", args...)
	// Explicit --host must remain authoritative even if the caller's environment
	// selects another context or TLS configuration. Context inspection below is
	// given its requested context as an argument before target commands begin.
	cmd.Env = make([]string, 0, len(os.Environ()))
	for _, entry := range os.Environ() {
		key, _, _ := strings.Cut(entry, "=")
		if key != "DOCKER_HOST" && key != "DOCKER_CONTEXT" && key != "DOCKER_TLS_VERIFY" && key != "DOCKER_CERT_PATH" {
			cmd.Env = append(cmd.Env, entry)
		}
	}
	// Resource IDs and JSON arrive on stdout. A successful Docker invocation can
	// emit warnings on stderr; mixing these streams corrupts protocol parsing.
	var output, diagnostic bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &diagnostic
	err := cmd.Run()
	if err != nil {
		if ctx.Err() != nil {
			return output.Bytes(), ctx.Err()
		}
		return output.Bytes(), fmt.Errorf("docker %s: %w: %s", args[0], err, strings.TrimSpace(diagnostic.String()+"\n"+output.String()))
	}
	if warning := strings.TrimSpace(diagnostic.String()); warning != "" {
		slog.Warn("Docker diagnostic", "message", warning)
	}
	return output.Bytes(), nil
}

func validateRun(runID string) error {
	if !runPattern.MatchString(runID) {
		return errors.New("run ID must be 8–64 lowercase letters, numbers or hyphens")
	}
	return nil
}

func (m *Manager) ensureEndpoint(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.endpoint != "" {
		return nil
	}
	host := m.getenv("DOCKER_HOST")
	if host == "" {
		args := []string{"context", "inspect"}
		if name := m.getenv("DOCKER_CONTEXT"); name != "" {
			args = append(args, name)
		}
		args = append(args, "--format", "{{.Endpoints.docker.Host}}")
		callCtx, cancel := context.WithTimeout(ctx, 15*time.Second)
		out, err := m.execute(callCtx, args...)
		cancel()
		if err != nil {
			return fmt.Errorf("resolve local Docker endpoint: %w", err)
		}
		host = strings.TrimSpace(string(out))
	}
	u, err := url.Parse(host)
	if err != nil || u.Scheme != "unix" || u.Host != "" || !strings.HasPrefix(u.Path, "/") || u.RawQuery != "" || u.Fragment != "" || u.User != nil {
		return fmt.Errorf("refusing Docker endpoint %q: only local unix:///absolute/socket endpoints are allowed", host)
	}
	m.endpoint = host
	return nil
}

func (m *Manager) command(ctx context.Context, timeout time.Duration, args ...string) ([]byte, error) {
	if err := m.ensureEndpoint(ctx); err != nil {
		return nil, err
	}
	callCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	m.mu.Lock()
	host := m.endpoint
	m.mu.Unlock()
	return m.execute(callCtx, append([]string{"--host", host}, args...)...)
}

func resourceName(runID string) string { return "redis-lab-" + runID }

// Provision returns a non-nil partial environment on failure. The caller must
// clean it up with a fresh bounded context even when the run context was cancelled.
func (m *Manager) Provision(ctx context.Context, runID string, cfg Config) (env *Environment, err error) {
	env = &Environment{RunID: runID, Metadata: map[string]string{}, NetworkIsolation: cfg.NetworkIsolation, RetainData: cfg.RetainData}
	if err = validateRun(runID); err != nil {
		return env, err
	}
	if cfg.Image == "" {
		cfg.Image = DefaultImage
	}
	if !imagePattern.MatchString(cfg.Image) {
		return env, errors.New("image must pin the official redis:X.Y.Z-alpine image, optionally with @sha256:digest")
	}
	if cfg.NodeTimeout == 0 {
		cfg.NodeTimeout = time.Second
	}
	if cfg.NodeTimeout < 100*time.Millisecond || cfg.NodeTimeout > time.Minute {
		return env, errors.New("Redis node timeout must be between 100ms and 1m")
	}
	if cfg.Persistence == "" {
		cfg.Persistence = "memory"
	}
	if !validPersistence(cfg.Persistence) || cfg.Persistence != "memory" && !cfg.RetainData {
		return env, errors.New("persistence must be memory, aof-everysec, or aof-always; AOF requires retained managed data")
	}
	env.Persistence = cfg.Persistence
	if err = m.ensureEndpoint(ctx); err != nil {
		return env, err
	}
	existing, discoverErr := m.Discover(ctx, runID)
	if discoverErr != nil {
		return env, fmt.Errorf("check run resources: %w", discoverErr)
	}
	if existing.NetworkID != "" || len(existing.Nodes) != 0 || len(existing.Volumes) != 0 {
		return env, errors.New("resources already exist for this run ID; clean them up or choose a new run ID")
	}
	creationAttempted := false
	defer func() {
		if err == nil || !creationAttempted {
			return
		}
		// Docker may finish creating a resource after command cancellation, before
		// its ID is read. Resolve only this run's exact labels and expected names.
		discoverCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		found, findErr := m.Discover(discoverCtx, runID)
		if found.NetworkID != "" {
			env.NetworkID = found.NetworkID
		}
		if len(found.Nodes) > 0 {
			env.Nodes = found.Nodes
		}
		if len(found.Volumes) > 0 {
			env.Volumes = found.Volumes
		}
		if findErr != nil {
			err = errors.Join(err, fmt.Errorf("recover partial resource inventory: %w", findErr))
		}
	}()
	m.mu.Lock()
	env.Metadata["docker_endpoint"] = m.endpoint
	m.mu.Unlock()
	env.Metadata["redis_image"] = cfg.Image
	env.Metadata["redis_node_timeout_ms"] = strconv.FormatInt(cfg.NodeTimeout.Milliseconds(), 10)
	env.Metadata["container_memory"] = "256m"
	env.Metadata["container_cpus"] = "1"
	env.Metadata["container_data"] = "128m tmpfs; persistence disabled"
	env.Metadata["redis_persistence"] = cfg.Persistence
	env.Metadata["redis_maxmemory"] = "128mb; noeviction"
	env.Metadata["container_restart_policy"] = "no; inspected before managed operations"
	if cfg.RetainData {
		env.Metadata["container_data"] = "six disposable managed local volumes; no filesystem quota; nodes.conf retained across container stops"
		env.Metadata["storage_bound"] = "finite workload and probes; Redis maxmemory 128mb; RDB snapshots disabled; growth-triggered AOF rewrite disabled; successful full synchronization can still trigger AOF rewrite; volume filesystem has no hard quota"
		env.Metadata["aof_rewrite_policy"] = "auto-aof-rewrite-percentage 0 disables growth-triggered rewrites only; full synchronization can trigger rewrites; AOF RDB preamble enabled for pinned image module compatibility; rewrite-time fsync suppression disabled"
		env.Metadata["memory_sync_files"] = "rdb-del-sync-files yes removes replication-created RDB files when AOF and scheduled snapshots are disabled; actual absence is verified before the memory control baseline"
		env.Metadata["aof_load_truncated"] = "Redis default yes is retained and observed in CONFIG GET; truncated AOF tail loading can discard data"
	}
	env.Metadata["container_network"] = "per-run bridge; egress IP masquerading disabled; client ports loopback only"
	info, infoErr := m.command(ctx, 20*time.Second, "info", "--format", "{{json .}}")
	if infoErr != nil {
		return env, fmt.Errorf("Docker daemon unavailable: %w", infoErr)
	}
	var daemon struct {
		ServerVersion, OperatingSystem, OSType, Architecture, KernelVersion, Driver string
		NCPU                                                                        int
		MemTotal                                                                    int64
	}
	if err = json.Unmarshal(info, &daemon); err != nil {
		return env, fmt.Errorf("decode Docker daemon information: %w", err)
	}
	if daemon.OSType != "linux" {
		return env, errors.New("Docker must use Linux containers")
	}
	env.Metadata["docker_version"] = daemon.ServerVersion
	env.Metadata["docker_os"] = daemon.OperatingSystem
	env.Metadata["docker_architecture"] = daemon.Architecture
	env.Metadata["docker_kernel"] = daemon.KernelVersion
	if daemon.NCPU > 0 {
		env.Metadata["docker_logical_cpus"] = strconv.Itoa(daemon.NCPU)
	}
	if daemon.MemTotal > 0 {
		env.Metadata["docker_memory_bytes"] = strconv.FormatInt(daemon.MemTotal, 10)
	}
	if daemon.Driver != "" {
		env.Metadata["docker_storage_driver"] = daemon.Driver
	}
	if _, err = m.command(ctx, 3*time.Minute, "image", "pull", cfg.Image); err != nil {
		return env, fmt.Errorf("pull pinned Redis image: %w", err)
	}
	imageInfo, imageErr := m.command(ctx, 15*time.Second, "image", "inspect", cfg.Image, "--format", "{{json .}}")
	if imageErr != nil {
		return env, imageErr
	}
	var imageData struct {
		ID          string
		RepoDigests []string
	}
	if err = json.Unmarshal(imageInfo, &imageData); err != nil {
		return env, fmt.Errorf("decode Redis image metadata: %w", err)
	}
	env.Metadata["redis_image_id"] = imageData.ID
	env.Metadata["redis_image_digests"] = strings.Join(imageData.RepoDigests, ",")
	creationAttempted = true
	// Docker's --internal bridge suppresses published ports on older daemons.
	// Disable bridge egress masquerading while retaining host loopback publishing.
	networkOut, subnet, networkErr := m.createRunNetwork(ctx, runID, cfg.NetworkIsolation, cfg.RetainData, cfg.Persistence)
	if networkErr != nil {
		return env, fmt.Errorf("create run network: %w", networkErr)
	}
	env.NetworkID = strings.TrimSpace(string(networkOut))
	if !idPattern.MatchString(env.NetworkID) {
		return env, errors.New("Docker returned an invalid network ID")
	}
	if cfg.NetworkIsolation || cfg.RetainData {
		env.Metadata["network_subnet"] = subnet
		env.Metadata["container_ports"] = "fixed ephemeral loopback ports retained across network reconnect and container start"
	}
	if cfg.NetworkIsolation {
		env.Metadata["network_isolation"] = "full-node bidirectional disconnect; peer and host client traffic unavailable"
		env.Metadata["network_subnet"] = subnet
		env.Metadata["container_ports"] = "fixed ephemeral loopback ports retained across network reconnect"
	}
	usedPorts := map[string]bool{}
	for index := 1; index <= nodeCount; index++ {
		key := fmt.Sprintf("node-%d", index)
		name := resourceName(runID) + "-" + key
		if cfg.RetainData {
			volume := Volume{Name: volumeName(runID, key), NodeID: key}
			env.Volumes = append(env.Volumes, volume)
			if err = m.createVolume(ctx, env, volume); err != nil {
				return env, err
			}
		}
		args := []string{"run", "--detach", "--restart", "no", "--name", name, "--label", managedLabel + "=true", "--label", runLabel + "=" + runID,
			"--network", env.NetworkID, "--publish", "127.0.0.1::6379", "--read-only", "--tmpfs", "/data:rw,noexec,nosuid,size=128m,uid=999,gid=999,mode=700",
			"--user", "999:999", "--cap-drop", "ALL", "--security-opt", "no-new-privileges", "--pids-limit", "100", "--memory", "256m", "--cpus", "1",
			cfg.Image, "redis-server", "--bind", "0.0.0.0", "--port", "6379", "--protected-mode", "no", "--cluster-enabled", "yes",
			"--cluster-config-file", "/data/nodes.conf", "--cluster-node-timeout", strconv.FormatInt(cfg.NodeTimeout.Milliseconds(), 10), "--appendonly", "no", "--save", "", "--dir", "/data"}
		args = append(args, "--maxmemory", "128mb", "--maxmemory-policy", "noeviction", "--auto-aof-rewrite-percentage", "0", "--repl-diskless-sync", "yes", "--repl-diskless-load", "swapdb")
		if cfg.RetainData {
			args = retainedRunArgs(args, runID, key, cfg.Persistence)
		}
		if cfg.NetworkIsolation || cfg.RetainData {
			choose := m.choosePort
			if choose == nil {
				choose = chooseLoopbackPort
			}
			port, portErr := choose(usedPorts)
			if portErr != nil {
				return env, portErr
			}
			usedPorts[port] = true
			ip := strings.TrimSuffix(subnet, "0/24") + strconv.Itoa(index+1)
			// Creation labels retain the allowed original endpoints even after a
			// hard stop leaves this running container without a network endpoint.
			endpointArgs := []string{"run", "--label", originalIPLabel + "=" + ip, "--label", originalPortLabel + "=" + port, "--ip", ip}
			if cfg.NetworkIsolation {
				endpointArgs = append(endpointArgs, "--label", isolationLabel+"=true")
			}
			args = append(endpointArgs, args[1:]...)
			for i := range args {
				if args[i] == "127.0.0.1::6379" {
					args[i] = "127.0.0.1:" + port + ":6379"
				}
			}
		}
		out, runErr := m.command(ctx, 30*time.Second, args...)
		if runErr != nil {
			return env, fmt.Errorf("start %s: %w", key, runErr)
		}
		node := Node{ID: key, Name: name, ContainerID: strings.TrimSpace(string(out))}
		if !idPattern.MatchString(node.ContainerID) {
			return env, errors.New("Docker returned an invalid container ID")
		}
		env.Nodes = append(env.Nodes, node)
		container, inspectErr := m.inspectContainer(ctx, node.ContainerID)
		if inspectErr != nil {
			return env, inspectErr
		}
		if err = validateContainer(container, env, node, true); err != nil {
			return env, err
		}
		node, err = containerNode(container, env, node)
		if err != nil {
			return env, err
		}
		if node.InternalAddr == "" || node.Addr == "" {
			return env, errors.New("managed container is missing its address or loopback port binding")
		}
		env.Nodes[len(env.Nodes)-1] = node
	}
	for _, node := range env.Nodes {
		if err = m.waitReady(ctx, env, node); err != nil {
			return env, err
		}
	}
	args := []string{"exec", env.Nodes[0].ContainerID, "redis-cli", "--cluster", "create"}
	for _, node := range env.Nodes {
		args = append(args, node.InternalAddr)
	}
	args = append(args, "--cluster-replicas", "1", "--cluster-yes")
	if _, err = m.command(ctx, 60*time.Second, args...); err != nil {
		return env, fmt.Errorf("bootstrap Redis cluster: %w", err)
	}
	if cfg.RetainData {
		for index := range env.Nodes {
			out, identityErr := m.command(ctx, 3*time.Second, "exec", env.Nodes[index].ContainerID, "redis-cli", "--raw", "cluster", "myid")
			identity := strings.TrimSpace(string(out))
			if identityErr != nil || len(identity) != 40 || !isHex(identity) {
				return env, errors.Join(errors.New("retained node lacks a valid initial Redis Cluster identity"), identityErr)
			}
			env.Nodes[index].RedisClusterID = identity
		}
	}
	return env, nil
}

type containerInspect struct {
	ID     string `json:"Id"`
	Name   string
	Config struct{ Labels map[string]string }
	State  struct {
		Running   bool
		Pid       int
		StartedAt string
	}
	HostConfig struct {
		ReadonlyRootfs bool
		Privileged     bool
		RestartPolicy  struct{ Name string }
		NetworkMode    string
		PortBindings   map[string][]struct{ HostIP, HostPort string }
	}
	Mounts          []mountInspect
	NetworkSettings struct {
		Networks map[string]struct{ NetworkID, IPAddress string }
		Ports    map[string][]struct{ HostIP, HostPort string }
	}
}

type networkInspect struct {
	ID         string `json:"Id"`
	Name       string
	Driver     string
	Internal   bool
	Labels     map[string]string
	Options    map[string]string
	Containers map[string]json.RawMessage
	IPAM       struct {
		Config []struct{ Subnet string }
	}
}

func (m *Manager) inspectContainer(ctx context.Context, id string) (containerInspect, error) {
	var value containerInspect
	if !idPattern.MatchString(id) {
		return value, errors.New("invalid tracked Docker container ID")
	}
	out, err := m.command(ctx, 15*time.Second, "container", "inspect", id, "--format", "{{json .}}")
	if err != nil {
		return value, err
	}
	if err = json.Unmarshal(out, &value); err != nil {
		return value, fmt.Errorf("decode container inspection: %w", err)
	}
	if value.ID != id {
		return value, errors.New("container inspection returned a different resource ID")
	}
	return value, nil
}

func (m *Manager) inspectNetwork(ctx context.Context, id string) (networkInspect, error) {
	var value networkInspect
	if !idPattern.MatchString(id) {
		return value, errors.New("invalid tracked Docker network ID")
	}
	out, err := m.command(ctx, 15*time.Second, "network", "inspect", id, "--format", "{{json .}}")
	if err != nil {
		return value, err
	}
	if err = json.Unmarshal(out, &value); err != nil {
		return value, fmt.Errorf("decode network inspection: %w", err)
	}
	if value.ID != id {
		return value, errors.New("network inspection returned a different resource ID")
	}
	return value, nil
}

func validateNetwork(network networkInspect, env *Environment) error {
	if network.Labels[managedLabel] != "true" || network.Labels[runLabel] != env.RunID || network.Name != resourceName(env.RunID) || network.Driver != "bridge" || (!network.Internal && network.Options[noMasqueradeOption] != "false") {
		return errors.New("refusing operation: network is not this run's managed bridge with egress masquerading disabled")
	}
	if env.NetworkIsolation && network.Labels[isolationLabel] != "true" {
		return errors.New("refusing operation: network lacks managed isolation capability")
	}
	if (network.Labels[retainedLabel] == "true") != env.RetainData || env.RetainData && network.Labels[persistenceLabel] != env.Persistence {
		return errors.New("refusing operation: network retained-data capability or persistence mode changed")
	}
	if network.Labels[isolationLabel] == "true" || env.RetainData {
		if err := validateIsolationNetwork(network); err != nil {
			return err
		}
	}
	return nil
}

func validateContainer(container containerInspect, env *Environment, node Node, requireRunning bool) error {
	validKey := false
	for index := 1; index <= nodeCount; index++ {
		if node.ID == fmt.Sprintf("node-%d", index) {
			validKey = true
		}
	}
	if !validKey {
		return errors.New("refusing operation: unexpected managed node key")
	}
	if container.Config.Labels[managedLabel] != "true" || container.Config.Labels[runLabel] != env.RunID || strings.TrimPrefix(container.Name, "/") != node.Name || node.Name != resourceName(env.RunID)+"-"+node.ID {
		return errors.New("refusing operation: container identity or management labels do not match this run")
	}
	if !container.HostConfig.ReadonlyRootfs || container.HostConfig.Privileged {
		return errors.New("refusing operation: managed container does not have a read-only root filesystem")
	}
	if container.HostConfig.RestartPolicy.Name != "no" {
		return errors.New("refusing operation: managed container restart policy must be explicitly no")
	}
	if err := validateDataMounts(container, env, node); err != nil {
		return err
	}
	if container.HostConfig.NetworkMode != env.NetworkID {
		return errors.New("refusing operation: managed container has unexpected network mode")
	}
	// A cancelled docker run can leave a created/stopped container before its
	// endpoint is attached. Its configured exact network ID still bounds cleanup.
	capable := container.Config.Labels[isolationLabel] == "true"
	if capable && !env.NetworkIsolation {
		return errors.New("refusing operation: isolation container does not match network capability")
	}
	if capable || env.RetainData {
		if err := validateOriginalEndpoints(container, node); err != nil {
			return err
		}
	}
	partialWithoutEndpoint := !requireRunning && (!container.State.Running || capable) && len(container.NetworkSettings.Networks) == 0
	if len(container.NetworkSettings.Networks) != 1 && !partialWithoutEndpoint {
		return errors.New("refusing operation: managed container has unexpected network membership")
	}
	member, ok := container.NetworkSettings.Networks[resourceName(env.RunID)]
	partialEndpoint := !requireRunning && !container.State.Running && ok && member.NetworkID == ""
	if (!ok || member.NetworkID != env.NetworkID) && !partialWithoutEndpoint && !partialEndpoint {
		return errors.New("refusing operation: container is outside this run's network")
	}
	bindings := container.NetworkSettings.Ports["6379/tcp"]
	if len(bindings) > 1 || (len(bindings) == 1 && bindings[0].HostIP != "127.0.0.1") {
		return errors.New("refusing operation: Redis port binding is not restricted to loopback")
	}
	if requireRunning && len(bindings) != 1 {
		return errors.New("managed container is missing its loopback port binding")
	}
	configuredBindings := container.HostConfig.PortBindings
	if len(configuredBindings) != 1 || len(configuredBindings["6379/tcp"]) != 1 || configuredBindings["6379/tcp"][0].HostIP != "127.0.0.1" {
		return errors.New("refusing operation: configured container port publication is not restricted to Redis on loopback")
	}
	if requireRunning && !container.State.Running {
		return errors.New("managed container is not running")
	}
	return nil
}

func containerNode(container containerInspect, env *Environment, node Node) (Node, error) {
	if container.Config.Labels[isolationLabel] == "true" || container.Config.Labels[retainedLabel] == "true" {
		node.InternalAddr = net.JoinHostPort(container.Config.Labels[originalIPLabel], "6379")
		node.Addr = net.JoinHostPort("127.0.0.1", container.Config.Labels[originalPortLabel])
	}
	member := container.NetworkSettings.Networks[resourceName(env.RunID)]
	if member.IPAddress != "" {
		ip := net.ParseIP(member.IPAddress)
		if ip == nil || !ip.IsPrivate() {
			return node, errors.New("managed container has an invalid private network address")
		}
		node.InternalAddr = net.JoinHostPort(member.IPAddress, "6379")
	}
	bindings := container.NetworkSettings.Ports["6379/tcp"]
	if len(bindings) > 0 {
		if len(bindings) != 1 || bindings[0].HostIP != "127.0.0.1" {
			return node, errors.New("refusing non-loopback Redis port binding")
		}
		port, err := strconv.Atoi(bindings[0].HostPort)
		if err != nil || port < 1 || port > 65535 {
			return node, errors.New("managed container has an invalid host port")
		}
		node.Addr = net.JoinHostPort("127.0.0.1", bindings[0].HostPort)
	}
	return node, nil
}

func (m *Manager) waitReady(ctx context.Context, env *Environment, node Node) error {
	readyCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		out, err := m.command(readyCtx, 3*time.Second, "exec", node.ContainerID, "redis-cli", "ping")
		if err == nil && strings.TrimSpace(string(out)) == "PONG" {
			return nil
		}
		select {
		case <-readyCtx.Done():
			return fmt.Errorf("wait for %s Redis readiness: %w", node.ID, readyCtx.Err())
		case <-ticker.C:
		}
	}
}

// Kill abruptly terminates exactly one tracked node after revalidating ownership.
func (m *Manager) Kill(ctx context.Context, env *Environment, node Node) error {
	if env == nil {
		return errors.New("nil managed environment")
	}
	if err := validateRun(env.RunID); err != nil {
		return err
	}
	tracked := false
	for _, candidate := range env.Nodes {
		if candidate.ContainerID == node.ContainerID && candidate.ID == node.ID && candidate.Name == node.Name {
			tracked = true
		}
	}
	if !tracked {
		return errors.New("refusing to kill an untracked node")
	}
	network, err := m.inspectNetwork(ctx, env.NetworkID)
	if err != nil {
		return err
	}
	if err = validateNetwork(network, env); err != nil {
		return err
	}
	container, err := m.inspectContainer(ctx, node.ContainerID)
	if err != nil {
		return err
	}
	if err = validateContainer(container, env, node, true); err != nil {
		return err
	}
	if err = m.validateNodeVolume(ctx, env, node); err != nil {
		return err
	}
	_, err = m.command(ctx, 15*time.Second, "container", "kill", "--signal", "KILL", node.ContainerID)
	return err
}

func isMissing(err error) bool {
	if err == nil {
		return false
	}
	message := strings.ToLower(err.Error())
	return strings.Contains(message, "no such container") || strings.Contains(message, "no such network") || strings.Contains(message, "no such object") || strings.Contains(message, "no such volume") || regexp.MustCompile(`network [a-f0-9]{64} not found`).MatchString(message)
}

// Cleanup is idempotent and removes only exact tracked resources. Ownership,
// isolation, and network membership are inspected before every destructive call.
func (m *Manager) Cleanup(ctx context.Context, env *Environment) error {
	if env == nil {
		return nil
	}
	if err := validateRun(env.RunID); err != nil {
		return err
	}
	if env.NetworkID == "" && len(env.Nodes) == 0 && len(env.Volumes) == 0 {
		return nil
	}
	var network networkInspect
	var err error
	if env.NetworkID != "" {
		network, err = m.inspectNetwork(ctx, env.NetworkID)
	} else if len(env.Nodes) != 0 {
		return errors.New("refusing cleanup of containers without their tracked managed network")
	}
	if err != nil && !isMissing(err) {
		return fmt.Errorf("validate cleanup network: %w", err)
	}
	if err == nil && env.NetworkID != "" {
		if err = validateNetwork(network, env); err != nil {
			return err
		}
	}
	var failures []error
	for _, node := range env.Nodes {
		container, inspectErr := m.inspectContainer(ctx, node.ContainerID)
		if isMissing(inspectErr) {
			continue
		}
		if inspectErr != nil {
			failures = append(failures, inspectErr)
			continue
		}
		if inspectErr = validateContainer(container, env, node, false); inspectErr != nil {
			failures = append(failures, inspectErr)
			continue
		}
		if inspectErr = m.validateNodeVolume(ctx, env, node); inspectErr != nil {
			failures = append(failures, inspectErr)
			continue
		}
		if _, removeErr := m.command(ctx, 30*time.Second, "container", "rm", "--force", node.ContainerID); removeErr != nil {
			failures = append(failures, removeErr)
			continue
		}
		if _, checkErr := m.inspectContainer(ctx, node.ContainerID); !isMissing(checkErr) {
			if checkErr == nil {
				checkErr = errors.New("container remains after removal")
			}
			failures = append(failures, checkErr)
		}
	}
	if len(failures) > 0 {
		return errors.Join(failures...)
	}
	if err = m.cleanupVolumes(ctx, env); err != nil {
		return err
	}
	if env.NetworkID == "" {
		return nil
	}
	network, err = m.inspectNetwork(ctx, env.NetworkID)
	if isMissing(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if err = validateNetwork(network, env); err != nil {
		return err
	}
	if len(network.Containers) != 0 {
		return errors.New("refusing network removal: network still contains containers")
	}
	if _, err = m.command(ctx, 30*time.Second, "network", "rm", env.NetworkID); err != nil {
		return err
	}
	if _, err = m.inspectNetwork(ctx, env.NetworkID); !isMissing(err) {
		if err == nil {
			err = errors.New("network remains after removal")
		}
		return err
	}
	return nil
}

// Discover recovers a run's inventory after interruption. It cannot select an
// arbitrary Docker target: run labels, expected names and isolation are checked.
func (m *Manager) Discover(ctx context.Context, runID string) (*Environment, error) {
	env := &Environment{RunID: runID, Metadata: map[string]string{}}
	if err := validateRun(runID); err != nil {
		return env, err
	}
	if err := m.discoverVolumes(ctx, env); err != nil {
		return env, err
	}
	filters := []string{"--filter", "label=" + managedLabel + "=true", "--filter", "label=" + runLabel + "=" + runID, "--no-trunc", "--format", "{{.ID}}"}
	networkIDs, err := m.command(ctx, 15*time.Second, append([]string{"network", "ls"}, filters...)...)
	if err != nil {
		return env, err
	}
	ids := strings.Fields(string(networkIDs))
	if len(ids) > 1 {
		return env, errors.New("multiple managed networks have this run ID; refusing ambiguous cleanup")
	}
	if len(ids) == 1 {
		env.NetworkID = ids[0]
		network, inspectErr := m.inspectNetwork(ctx, env.NetworkID)
		if inspectErr != nil {
			return env, inspectErr
		}
		env.NetworkIsolation = network.Labels[isolationLabel] == "true"
		if network.Labels[retainedLabel] == "true" {
			mode := network.Labels[persistenceLabel]
			if !validPersistence(mode) || env.Persistence != "" && env.Persistence != mode {
				return env, errors.New("managed network and volumes disagree on persistence mode")
			}
			env.RetainData, env.Persistence = true, mode
		}
		if env.NetworkIsolation || env.RetainData {
			env.Metadata["network_subnet"] = network.Labels[subnetLabel]
		}
		if inspectErr = validateNetwork(network, env); inspectErr != nil {
			return env, inspectErr
		}
	}
	containerIDs, err := m.command(ctx, 15*time.Second, append([]string{"container", "ls", "--all"}, filters...)...)
	if err != nil {
		return env, err
	}
	for _, id := range strings.Fields(string(containerIDs)) {
		container, inspectErr := m.inspectContainer(ctx, id)
		if inspectErr != nil {
			return env, inspectErr
		}
		name := strings.TrimPrefix(container.Name, "/")
		key := strings.TrimPrefix(name, resourceName(runID)+"-")
		validKey := false
		for index := 1; index <= nodeCount; index++ {
			if key == fmt.Sprintf("node-%d", index) {
				validKey = true
			}
		}
		if !validKey {
			return env, errors.New("managed labels found on an unexpected container name")
		}
		node := Node{ID: key, Name: name, ContainerID: id}
		if inspectErr = validateContainer(container, env, node, false); inspectErr != nil {
			return env, inspectErr
		}
		node, inspectErr = containerNode(container, env, node)
		if inspectErr != nil {
			return env, inspectErr
		}
		env.Nodes = append(env.Nodes, node)
	}
	sort.Slice(env.Nodes, func(i, j int) bool { return env.Nodes[i].ID < env.Nodes[j].ID })
	return env, nil
}
