package environment

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	retainedLabel    = "org.redis-reliability-lab.retained-data"
	persistenceLabel = "org.redis-reliability-lab.persistence"
	volumeNodeLabel  = "org.redis-reliability-lab.node-id"
)

// Volume is an explicitly created disposable local Docker volume. Docker volume
// identity is its immutable name; labels and the local driver are revalidated.
type Volume struct {
	Name   string `json:"name"`
	NodeID string `json:"node_id"`
}

type mountInspect struct {
	Type        string
	Name        string
	Destination string
	RW          bool
}

type volumeInspect struct {
	Name    string
	Driver  string
	Scope   string
	Labels  map[string]string
	Options map[string]string
}

// RestartReceipt is returned before starting a stopped managed container. A
// command error remains an error even when subsequent inspection observes that
// Docker applied it. Redis Cluster identity must survive; process identity must
// change. Memory mode retains nodes.conf but cannot retain volatile key values.
type RestartReceipt struct {
	RunID           string          `json:"run_id"`
	NodeID          string          `json:"node_id"`
	Node            Node            `json:"node"`
	RequestedAt     time.Time       `json:"requested_at"`
	CompletedAt     time.Time       `json:"completed_at,omitempty"`
	Before          ProcessEvidence `json:"before"`
	After           ProcessEvidence `json:"after"`
	OriginalRedisID string          `json:"original_redis_id"`
}

func validPersistence(mode string) bool {
	return mode == "memory" || mode == "aof-everysec" || mode == "aof-always"
}

// redis-cli --raw emits one newline per array element. Remove only its final
// output terminator: when save is the last key, its empty value must survive as
// the final empty field. Trimming whitespace would corrupt this valid reply.
func persistenceConfigReply(reply []byte) string {
	return strings.TrimSuffix(strings.ReplaceAll(string(reply), "\r\n", "\n"), "\n")
}

func validatePersistenceConfig(raw, mode string) error {
	if !validPersistence(mode) {
		return errors.New("invalid managed persistence mode")
	}
	lines := strings.Split(raw, "\n")
	if len(lines)%2 != 0 {
		return errors.New("invalid persistence CONFIG GET observation")
	}
	values := map[string]string{}
	for index := 0; index < len(lines); index += 2 {
		key := lines[index]
		if key == "" || strings.TrimSpace(key) != key {
			return errors.New("invalid persistence CONFIG GET key")
		}
		if _, duplicate := values[key]; duplicate {
			return errors.New("duplicate persistence CONFIG GET key")
		}
		values[key] = lines[index+1]
	}
	appendonly := "no"
	if mode != "memory" {
		appendonly = "yes"
	}
	preamble := "no"
	if mode != "memory" {
		preamble = "yes"
	}
	wantedValues := map[string]string{"appendonly": appendonly, "save": "", "repl-diskless-load": "swapdb", "repl-diskless-sync": "yes", "auto-aof-rewrite-percentage": "0", "maxmemory": "134217728", "maxmemory-policy": "noeviction", "aof-use-rdb-preamble": preamble, "no-appendfsync-on-rewrite": "no", "aof-load-truncated": "yes", "rdb-del-sync-files": "yes"}
	if len(values) != len(wantedValues)+1 {
		return errors.New("unexpected persistence CONFIG GET fields")
	}
	for key, wanted := range wantedValues {
		if value, ok := values[key]; !ok || value != wanted {
			return fmt.Errorf("managed Redis persistence configuration changed: %s", key)
		}
	}
	if policy := values["appendfsync"]; policy != "always" && policy != "everysec" && policy != "no" {
		return errors.New("invalid managed Redis AOF fsync policy")
	}
	if mode != "memory" {
		wanted := strings.TrimPrefix(mode, "aof-")
		if values["appendfsync"] != wanted {
			return errors.New("managed Redis AOF fsync policy changed")
		}
	}
	return nil
}

// validateDataFiles accepts only canonical file paths within the fixed find
// scope. The raw response is retained separately; no listed path is executed.
func validateDataFiles(raw string) error {
	if len(raw) > 1<<20 || strings.ContainsAny(raw, "\r\x00") {
		return errors.New("invalid managed data file listing")
	}
	if raw == "" {
		return nil
	}
	if !strings.HasSuffix(raw, "\n") {
		return errors.New("incomplete managed data file listing")
	}
	lines := strings.Split(strings.TrimSuffix(raw, "\n"), "\n")
	if len(lines) > 4096 {
		return errors.New("managed data file listing exceeds observation bound")
	}
	seen := map[string]bool{}
	for _, name := range lines {
		if !strings.HasPrefix(name, "/data/") || path.Clean(name) != name || strings.TrimSpace(name) != name || seen[name] {
			return errors.New("foreign or malformed managed data file path")
		}
		parts := strings.Split(strings.TrimPrefix(name, "/data/"), "/")
		if len(parts) > 2 {
			return errors.New("managed data file path exceeds observation depth")
		}
		for _, part := range parts {
			if part == "" || part == "." || part == ".." {
				return errors.New("malformed managed data file path")
			}
			for _, c := range part {
				if !(c >= 'a' && c <= 'z') && !(c >= 'A' && c <= 'Z') && !(c >= '0' && c <= '9') && c != '.' && c != '-' && c != '_' {
					return errors.New("invalid managed data file name")
				}
			}
		}
		seen[name] = true
	}
	return nil
}

func validateKeyspace(raw string) error {
	lines := strings.Split(strings.ReplaceAll(raw, "\r\n", "\n"), "\n")
	if len(raw) > 1<<20 || len(lines) == 0 || lines[0] != "# Keyspace" {
		return errors.New("invalid managed keyspace observation")
	}
	seen := map[string]bool{}
	for _, line := range lines[1:] {
		db, values, ok := strings.Cut(line, ":")
		index, parseErr := strconv.ParseUint(strings.TrimPrefix(db, "db"), 10, 32)
		if !ok || parseErr != nil || db != "db"+strconv.FormatUint(index, 10) || seen[db] {
			return errors.New("invalid managed keyspace database")
		}
		fields := map[string]bool{}
		for _, entry := range strings.Split(values, ",") {
			key, value, ok := strings.Cut(entry, "=")
			if !ok || key == "" || fields[key] {
				return errors.New("invalid managed keyspace field")
			}
			for _, c := range key {
				if !(c >= 'a' && c <= 'z') && c != '_' {
					return errors.New("invalid managed keyspace field name")
				}
			}
			n, parseErr := strconv.ParseUint(value, 10, 64)
			if parseErr != nil || value != strconv.FormatUint(n, 10) {
				return errors.New("invalid managed keyspace count")
			}
			fields[key] = true
		}
		if !fields["keys"] || !fields["expires"] || !fields["avg_ttl"] {
			return errors.New("incomplete managed keyspace observation")
		}
		seen[db] = true
	}
	return nil
}

func validNodeKey(key string) bool {
	for index := 1; index <= nodeCount; index++ {
		if key == fmt.Sprintf("node-%d", index) {
			return true
		}
	}
	return false
}

func volumeName(runID, nodeID string) string { return resourceName(runID) + "-" + nodeID + "-data" }

func retainedRunArgs(args []string, runID, key, mode string) []string {
	out := []string{"run", "--label", retainedLabel + "=true", "--label", persistenceLabel + "=" + mode,
		"--mount", "type=volume,src=" + volumeName(runID, key) + ",dst=/data"}
	for index := 1; index < len(args); index++ {
		if args[index] == "--tmpfs" {
			index++
			continue
		}
		out = append(out, args[index])
	}
	preamble := "no"
	if mode != "memory" {
		preamble = "yes"
	}
	out = append(out, "--aof-use-rdb-preamble", preamble, "--no-appendfsync-on-rewrite", "no", "--rdb-del-sync-files", "yes")
	if mode != "memory" {
		for index := range out {
			if out[index] == "--appendonly" {
				out[index+1] = "yes"
			}
		}
		policy := "everysec"
		if mode == "aof-always" {
			policy = "always"
		}
		out = append(out, "--appendfsync", policy)
	}
	return out
}

func validateVolume(volume volumeInspect, env *Environment, tracked Volume) error {
	if env == nil || !env.RetainData || !validPersistence(env.Persistence) || !validNodeKey(tracked.NodeID) || tracked.Name != volumeName(env.RunID, tracked.NodeID) || volume.Name != tracked.Name || volume.Driver != "local" || volume.Scope != "local" || len(volume.Options) != 0 || volume.Labels[managedLabel] != "true" || volume.Labels[runLabel] != env.RunID || volume.Labels[retainedLabel] != "true" || volume.Labels[volumeNodeLabel] != tracked.NodeID || volume.Labels[persistenceLabel] != env.Persistence {
		return errors.New("refusing operation: volume is not this node's managed disposable local data volume")
	}
	return nil
}

func (m *Manager) inspectVolume(ctx context.Context, name string) (volumeInspect, error) {
	var volume volumeInspect
	// Names alone are never authorization: callers must validate exact run/node
	// names and immutable ownership labels before using or removing the volume.
	if len(name) > 100 || !strings.HasPrefix(name, "redis-lab-") || strings.ContainsAny(name, "/\\ ,\t\n") {
		return volume, errors.New("invalid tracked Docker volume name")
	}
	out, err := m.command(ctx, 15*time.Second, "volume", "inspect", name, "--format", "{{json .}}")
	if err != nil {
		return volume, err
	}
	if err = json.Unmarshal(out, &volume); err != nil {
		return volume, fmt.Errorf("decode volume inspection: %w", err)
	}
	if volume.Name != name {
		return volume, errors.New("volume inspection returned a different resource name")
	}
	return volume, nil
}

func (m *Manager) createVolume(ctx context.Context, env *Environment, volume Volume) error {
	out, err := m.command(ctx, 15*time.Second, "volume", "create", "--driver", "local", "--label", managedLabel+"=true", "--label", runLabel+"="+env.RunID, "--label", retainedLabel+"=true", "--label", persistenceLabel+"="+env.Persistence, "--label", volumeNodeLabel+"="+volume.NodeID, volume.Name)
	if err != nil {
		return fmt.Errorf("create %s data volume: %w", volume.NodeID, err)
	}
	if strings.TrimSpace(string(out)) != volume.Name {
		return errors.New("Docker returned an unexpected data volume name")
	}
	inspected, err := m.inspectVolume(ctx, volume.Name)
	if err != nil {
		return err
	}
	return validateVolume(inspected, env, volume)
}

func (m *Manager) discoverVolumes(ctx context.Context, env *Environment) error {
	out, err := m.command(ctx, 15*time.Second, "volume", "ls", "--filter", "label="+managedLabel+"=true", "--filter", "label="+runLabel+"="+env.RunID, "--format", "{{.Name}}")
	if err != nil {
		return err
	}
	seen := map[string]bool{}
	for _, name := range strings.Fields(string(out)) {
		volume, inspectErr := m.inspectVolume(ctx, name)
		if inspectErr != nil {
			return inspectErr
		}
		key := volume.Labels[volumeNodeLabel]
		mode := volume.Labels[persistenceLabel]
		if !validPersistence(mode) || env.Persistence != "" && env.Persistence != mode || seen[key] {
			return errors.New("ambiguous retained data volume inventory")
		}
		env.RetainData, env.Persistence = true, mode
		tracked := Volume{Name: name, NodeID: key}
		if inspectErr = validateVolume(volume, env, tracked); inspectErr != nil {
			return inspectErr
		}
		seen[key] = true
		env.Volumes = append(env.Volumes, tracked)
	}
	sort.Slice(env.Volumes, func(i, j int) bool { return env.Volumes[i].NodeID < env.Volumes[j].NodeID })
	return nil
}

func validateDataMounts(container containerInspect, env *Environment, node Node) error {
	retained := container.Config.Labels[retainedLabel] == "true"
	if retained != env.RetainData {
		return errors.New("refusing operation: container retained-data capability does not match this run")
	}
	count := 0
	for _, mount := range container.Mounts {
		if mount.Type == "bind" {
			return errors.New("refusing operation: managed container has a host bind mount")
		}
		if mount.Type == "volume" {
			count++
			if !retained || mount.Name != volumeName(env.RunID, node.ID) || mount.Destination != "/data" || !mount.RW {
				return errors.New("refusing operation: container has a foreign or altered data volume mount")
			}
		}
		if mount.Type != "volume" && mount.Type != "tmpfs" {
			return errors.New("refusing operation: unexpected container mount type")
		}
	}
	if retained {
		if count != 1 || container.Config.Labels[persistenceLabel] != env.Persistence {
			return errors.New("refusing operation: retained data mount or persistence mode changed")
		}
		tracked := false
		for _, volume := range env.Volumes {
			if volume.NodeID == node.ID && volume.Name == volumeName(env.RunID, node.ID) {
				tracked = true
			}
		}
		if !tracked {
			return errors.New("refusing operation: container data volume is not tracked")
		}
	}
	return nil
}

func (m *Manager) validateNodeVolume(ctx context.Context, env *Environment, node Node) error {
	if !env.RetainData {
		return nil
	}
	tracked := Volume{Name: volumeName(env.RunID, node.ID), NodeID: node.ID}
	volume, err := m.inspectVolume(ctx, tracked.Name)
	if err != nil {
		return err
	}
	return validateVolume(volume, env, tracked)
}

func (m *Manager) processNode(ctx context.Context, env *Environment, node Node) (containerInspect, networkInspect, error) {
	if env != nil && env.NetworkIsolation {
		return m.isolationNode(ctx, env, node)
	}
	var container containerInspect
	var network networkInspect
	if env == nil {
		return container, network, errors.New("nil managed environment")
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
	if !tracked {
		return container, network, errors.New("refusing process observation of an untracked node")
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
	if err = validateContainer(container, env, node, false); err != nil {
		return container, network, err
	}
	if err = m.validateNodeVolume(ctx, env, node); err != nil {
		return container, network, err
	}
	if env.RetainData {
		if err = validateRetainedEndpoint(container, network); err != nil {
			return container, network, err
		}
	}
	return container, network, nil
}

// Restart starts the same retained managed container after a crash. It uses
// container start rather than stopping a live process or creating a replacement.
func (m *Manager) Restart(ctx context.Context, env *Environment, node Node) (*RestartReceipt, error) {
	requestedAt := time.Now()
	if env == nil || !env.RetainData || len(node.RedisClusterID) != 40 || !isHex(node.RedisClusterID) {
		return nil, errors.New("restart requires retained managed data and the original Redis Cluster identity")
	}
	restartCtx, cancel := context.WithTimeout(ctx, 45*time.Second)
	defer cancel()
	container, _, err := m.processNode(restartCtx, env, node)
	if err != nil {
		return nil, err
	}
	receipt := &RestartReceipt{RunID: env.RunID, NodeID: node.ID, Node: node, RequestedAt: requestedAt, OriginalRedisID: node.RedisClusterID}
	receipt.Before = ProcessEvidence{StartedAt: receipt.RequestedAt, ObservedAt: time.Now(), ContainerRunning: container.State.Running, ContainerPID: container.State.Pid, ContainerStartedAt: container.State.StartedAt, ContainerRestartPolicy: container.HostConfig.RestartPolicy.Name, NetworkCount: len(container.NetworkSettings.Networks), NetworkConnected: container.State.Running && len(container.NetworkSettings.Networks) == 1}
	if container.State.Running {
		receipt.Before, err = m.ProbeProcess(restartCtx, env, node)
		return receipt, errors.Join(errors.New("refusing to restart a running managed process"), err)
	}
	// Docker releases runtime network endpoints on stop. Immutable creation
	// labels, configured fixed port and exact network mode bound container start;
	// the independently inspected postcondition requires the original IP/port.
	_, commandErr := m.command(restartCtx, 15*time.Second, "container", "start", node.ContainerID)
	checkCtx := restartCtx
	checkCancel := func() {}
	independent := commandErr != nil || restartCtx.Err() != nil
	if independent {
		checkCtx, checkCancel = context.WithTimeout(context.Background(), 15*time.Second)
	}
	checkErr := m.inspectRestart(checkCtx, env, node, receipt)
	checkCancel()
	if !independent && restartCtx.Err() != nil {
		// Cancellation while observing must not hide an applied start. Observe
		// once on an independent budget, while preserving the caller's error.
		freshCtx, freshCancel := context.WithTimeout(context.Background(), 15*time.Second)
		checkErr = m.inspectRestart(freshCtx, env, node, receipt)
		freshCancel()
	}
	return receipt, errors.Join(commandErr, restartCtx.Err(), checkErr)
}

func (m *Manager) inspectRestart(checkCtx context.Context, env *Environment, node Node, receipt *RestartReceipt) error {
	current, _, checkErr := m.processNode(checkCtx, env, node)
	if checkErr == nil && !current.State.Running {
		checkErr = errors.New("managed node remains stopped after container start")
	}
	if checkErr == nil {
		checkErr = m.waitReady(checkCtx, env, node)
	}
	if checkErr == nil {
		receipt.After, checkErr = m.ProbeProcess(checkCtx, env, node)
	}
	if checkErr == nil {
		after, currentNetwork, inspectErr := m.processNode(checkCtx, env, node)
		checkErr = inspectErr
		if checkErr == nil {
			checkErr = validateContainer(after, env, node, true)
		}
		if checkErr == nil {
			if _, ok := currentNetwork.Containers[node.ContainerID]; !ok {
				checkErr = errors.New("managed bridge does not report restarted target membership")
			}
		}
		if checkErr == nil && (receipt.After.RedisClusterID != receipt.OriginalRedisID || receipt.After.ContainerStartedAt == receipt.Before.ContainerStartedAt) {
			checkErr = errors.New("restart did not preserve Redis Cluster identity and change process start time")
		}
		if checkErr == nil {
			receipt.CompletedAt = time.Now()
		}
	}
	return checkErr
}

func (m *Manager) cleanupVolumes(ctx context.Context, env *Environment) error {
	var failures []error
	for _, tracked := range env.Volumes {
		volume, err := m.inspectVolume(ctx, tracked.Name)
		if isMissing(err) {
			continue
		}
		if err == nil {
			err = validateVolume(volume, env, tracked)
		}
		if err == nil {
			_, err = m.command(ctx, 15*time.Second, "volume", "rm", tracked.Name)
		}
		if err == nil {
			_, err = m.inspectVolume(ctx, tracked.Name)
			if isMissing(err) {
				err = nil
			} else if err == nil {
				err = errors.New("managed data volume remains after removal")
			}
		}
		if err != nil {
			failures = append(failures, err)
		}
	}
	return errors.Join(failures...)
}

// Verify the original retained IP is also contained in this run's fixed pool.
func validateRetainedEndpoint(container containerInspect, network networkInspect) error {
	_, subnet, _ := net.ParseCIDR(network.Labels[subnetLabel])
	if subnet == nil || !subnet.Contains(net.ParseIP(container.Config.Labels[originalIPLabel])) {
		return errors.New("retained node endpoint is outside the managed private pool")
	}
	return nil
}
