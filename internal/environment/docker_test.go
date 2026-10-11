package environment

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"
)

const testRun = "abc12345"

type fakeDocker struct {
	calls                  [][]string
	network                networkInspect
	networkExists          bool
	containers             map[string]containerInspect
	runCount               int
	failRun                int
	loseRunOutput          bool
	leaveAfterRemove       bool
	pingFailure            bool
	disconnectError        error
	connectError           error
	disconnectNoop         bool
	connectNoop            bool
	foreignMembers         map[string]json.RawMessage
	overlapAttempts        int
	redisRunID             string
	redisPID               int
	volumes                map[string]volumeInspect
	volumeCount            int
	failVolume             int
	loseVolumeOutput       bool
	leaveVolumeAfterRemove bool
	volumeInUse            bool
	startError             error
	startNoop              bool
	starts                 map[string]int
	clusterID              string
}

func newFake() (*Manager, *fakeDocker) {
	fake := &fakeDocker{containers: map[string]containerInspect{}, volumes: map[string]volumeInspect{}, starts: map[string]int{}}
	manager := &Manager{execute: fake.execute, getenv: func(key string) string {
		if key == "DOCKER_HOST" {
			return "unix:///tmp/docker.sock"
		}
		return ""
	}}
	manager.choosePort = func(used map[string]bool) (string, error) { return strconvPort(len(used) + 1), nil }
	return manager, fake
}

func encode(value any) ([]byte, error) { return json.Marshal(value) }

func (f *fakeDocker) execute(ctx context.Context, supplied ...string) ([]byte, error) {
	f.calls = append(f.calls, append([]string(nil), supplied...))
	if _, ok := ctx.Deadline(); !ok {
		return nil, errors.New("unbounded Docker invocation")
	}
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	args := supplied
	if len(args) >= 2 && args[0] == "--host" {
		args = args[2:]
	}
	switch {
	case reflect.DeepEqual(args, []string{"context", "inspect", "--format", "{{.Endpoints.docker.Host}}"}):
		return []byte("unix:///tmp/docker.sock"), nil
	case args[0] == "info":
		return []byte(`{"ServerVersion":"28.0.0","OSType":"linux","OperatingSystem":"Fake Linux","Architecture":"aarch64","KernelVersion":"6.1"}`), nil
	case args[0] == "image" && args[1] == "pull":
		return nil, nil
	case args[0] == "image" && args[1] == "inspect":
		return []byte(`{"ID":"sha256:pinned","RepoDigests":["redis@sha256:verified"]}`), nil
	case args[0] == "volume" && args[1] == "ls":
		var names []string
		for name := range f.volumes {
			names = append(names, name)
		}
		return []byte(strings.Join(names, "\n")), nil
	case args[0] == "volume" && args[1] == "create":
		f.volumeCount++
		if f.volumeCount == f.failVolume && !f.loseVolumeOutput {
			return nil, errors.New("injected volume creation failure")
		}
		name := args[len(args)-1]
		volume := volumeInspect{Name: name, Driver: "local", Scope: "local", Labels: map[string]string{}}
		for i, arg := range args {
			if arg == "--label" {
				key, value, _ := strings.Cut(args[i+1], "=")
				volume.Labels[key] = value
			}
		}
		f.volumes[name] = volume
		if f.volumeCount == f.failVolume {
			return nil, errors.New("volume created but response lost")
		}
		return []byte(name), nil
	case args[0] == "volume" && args[1] == "inspect":
		volume, ok := f.volumes[args[2]]
		if !ok {
			return nil, errors.New("No such volume: " + args[2])
		}
		return encode(volume)
	case args[0] == "volume" && args[1] == "rm":
		if f.volumeInUse {
			return nil, errors.New("volume is in use by another container")
		}
		if !f.leaveVolumeAfterRemove {
			delete(f.volumes, args[2])
		}
		return nil, nil
	case args[0] == "network" && args[1] == "ls":
		if f.networkExists {
			return []byte(f.network.ID), nil
		}
		return nil, nil
	case args[0] == "container" && args[1] == "ls":
		var ids []string
		for id := range f.containers {
			ids = append(ids, id)
		}
		return []byte(strings.Join(ids, "\n")), nil
	case args[0] == "network" && args[1] == "create":
		if f.overlapAttempts > 0 {
			f.overlapAttempts--
			return nil, errors.New("Pool overlaps with other one on this address space")
		}
		f.network.ID = strings.Repeat("a", 64)
		f.network.Name = resourceName(testRun)
		f.network.Driver = "bridge"
		f.network.Internal = false
		f.network.Options = map[string]string{noMasqueradeOption: "false"}
		f.network.Labels = map[string]string{managedLabel: "true", runLabel: testRun}
		for i, arg := range args {
			if arg == "--label" {
				key, value, _ := strings.Cut(args[i+1], "=")
				f.network.Labels[key] = value
			}
			if arg == "--subnet" {
				f.network.IPAM.Config = []struct{ Subnet string }{{args[i+1]}}
			}
		}
		f.networkExists = true
		return []byte(f.network.ID), nil
	case args[0] == "network" && args[1] == "inspect":
		if !f.networkExists {
			return nil, fmt.Errorf("Error response from daemon: network %s not found", args[2])
		}
		f.network.Containers = map[string]json.RawMessage{}
		for id, container := range f.containers {
			if member, ok := container.NetworkSettings.Networks[f.network.Name]; ok {
				f.network.Containers[id] = json.RawMessage(`{"IPv4Address":"` + member.IPAddress + `/24"}`)
			}
		}
		for id, member := range f.foreignMembers {
			f.network.Containers[id] = member
		}
		return encode(f.network)
	case args[0] == "network" && args[1] == "rm":
		f.networkExists = false
		return nil, nil
	case args[0] == "network" && args[1] == "disconnect":
		id := args[3]
		if !f.disconnectNoop {
			container := f.containers[id]
			container.NetworkSettings.Networks = nil
			container.NetworkSettings.Ports = nil
			f.containers[id] = container
		}
		return nil, f.disconnectError
	case args[0] == "network" && args[1] == "connect":
		id := args[len(args)-1]
		if !f.connectNoop {
			container := f.containers[id]
			container.NetworkSettings.Networks = map[string]struct{ NetworkID, IPAddress string }{f.network.Name: {NetworkID: f.network.ID, IPAddress: args[3]}}
			container.NetworkSettings.Ports = map[string][]struct{ HostIP, HostPort string }{"6379/tcp": {{HostIP: "127.0.0.1", HostPort: container.Config.Labels[originalPortLabel]}}}
			f.containers[id] = container
		}
		return nil, f.connectError
	case args[0] == "run":
		f.runCount++
		if f.runCount == f.failRun && !f.loseRunOutput {
			return nil, errors.New("injected container start failure")
		}
		id := fmt.Sprintf("%064x", f.runCount)
		container := containerInspect{ID: id, Name: "/" + resourceName(testRun) + fmt.Sprintf("-node-%d", f.runCount)}
		container.Config.Labels = map[string]string{managedLabel: "true", runLabel: testRun}
		container.State.Running = true
		container.State.Pid = 1000 + f.runCount
		container.State.StartedAt = "2026-10-11T00:00:00Z"
		container.HostConfig.ReadonlyRootfs = true
		container.HostConfig.NetworkMode = f.network.ID
		container.HostConfig.PortBindings = map[string][]struct{ HostIP, HostPort string }{"6379/tcp": {{HostIP: "127.0.0.1", HostPort: ""}}}
		container.NetworkSettings.Networks = map[string]struct{ NetworkID, IPAddress string }{f.network.Name: {NetworkID: f.network.ID, IPAddress: fmt.Sprintf("172.25.0.%d", f.runCount+1)}}
		container.NetworkSettings.Ports = map[string][]struct{ HostIP, HostPort string }{"6379/tcp": {{HostIP: "127.0.0.1", HostPort: strconvPort(f.runCount)}}}
		for i, arg := range args {
			if arg == "--restart" {
				container.HostConfig.RestartPolicy.Name = args[i+1]
			}
			if arg == "--label" {
				key, value, _ := strings.Cut(args[i+1], "=")
				container.Config.Labels[key] = value
			}
			if arg == "--ip" {
				member := container.NetworkSettings.Networks[f.network.Name]
				member.IPAddress = args[i+1]
				container.NetworkSettings.Networks[f.network.Name] = member
			}
			if arg == "--publish" && args[i+1] != "127.0.0.1::6379" {
				port := strings.Split(args[i+1], ":")[1]
				container.HostConfig.PortBindings["6379/tcp"][0].HostPort = port
				container.NetworkSettings.Ports["6379/tcp"][0].HostPort = port
			}
			if arg == "--mount" {
				parts := strings.Split(args[i+1], ",")
				name := strings.TrimPrefix(parts[1], "src=")
				container.Mounts = []mountInspect{{Type: "volume", Name: name, Destination: "/data", RW: true}}
			}
		}
		f.containers[id] = container
		if f.runCount == f.failRun {
			return nil, errors.New("output lost after resource creation")
		}
		return []byte(id), nil
	case args[0] == "container" && args[1] == "inspect":
		container, exists := f.containers[args[2]]
		if !exists {
			return nil, errors.New("Error: No such object: " + args[2])
		}
		return encode(container)
	case args[0] == "container" && args[1] == "kill":
		id := args[len(args)-1]
		container := f.containers[id]
		container.State.Running = false
		f.containers[id] = container
		return nil, nil
	case args[0] == "container" && args[1] == "start":
		id := args[len(args)-1]
		if !f.startNoop {
			container := f.containers[id]
			f.starts[id]++
			container.State.Running = true
			container.State.Pid += 100
			container.State.StartedAt = fmt.Sprintf("2026-10-11T00:01:%02dZ", f.starts[id])
			container.NetworkSettings.Networks = map[string]struct{ NetworkID, IPAddress string }{f.network.Name: {NetworkID: f.network.ID, IPAddress: container.Config.Labels[originalIPLabel]}}
			container.NetworkSettings.Ports = map[string][]struct{ HostIP, HostPort string }{"6379/tcp": {{HostIP: "127.0.0.1", HostPort: container.Config.Labels[originalPortLabel]}}}
			f.containers[id] = container
		}
		return []byte(id), f.startError
	case args[0] == "container" && args[1] == "rm":
		if !f.leaveAfterRemove {
			delete(f.containers, args[len(args)-1])
		}
		return nil, nil
	case args[0] == "exec":
		if reflect.DeepEqual(args[2:], []string{"find", "/data", "-maxdepth", "2", "-type", "f"}) {
			files := "/data/nodes.conf\n"
			if mode := f.containers[args[1]].Config.Labels[persistenceLabel]; mode == "aof-everysec" || mode == "aof-always" {
				files += "/data/appendonlydir/appendonly.aof.manifest\n/data/appendonlydir/appendonly.aof.1.base.rdb\n/data/appendonlydir/appendonly.aof.1.incr.aof\n"
			}
			return []byte(files), nil
		}
		if len(args) >= 4 && args[2] == "du" {
			return []byte("8\t/data\n"), nil
		}
		if len(args) > 6 && args[6] == "config" {
			mode := f.containers[args[1]].Config.Labels[persistenceLabel]
			aof, policy, preamble := "no", "everysec", "no"
			if mode != "memory" && mode != "" {
				aof, policy, preamble = "yes", strings.TrimPrefix(mode, "aof-"), "yes"
			}
			return []byte("appendonly\n" + aof + "\nappendfsync\n" + policy + "\nsave\n\nrepl-diskless-load\nswapdb\nrepl-diskless-sync\nyes\nauto-aof-rewrite-percentage\n0\nmaxmemory\n134217728\nmaxmemory-policy\nnoeviction\naof-use-rdb-preamble\n" + preamble + "\nno-appendfsync-on-rewrite\nno\naof-load-truncated\nyes\nrdb-del-sync-files\nyes\n"), nil
		}
		if args[len(args)-1] == "ping" {
			if f.pingFailure {
				return nil, errors.New("Redis is not ready")
			}
			return []byte("PONG\n"), nil
		}
		if args[len(args)-1] == "server" {
			runID := f.redisRunID
			if runID == "" {
				runID = strings.Repeat("c", 40)
				if f.starts[args[1]] > 0 {
					runID = fmt.Sprintf("%040x", 100+f.starts[args[1]])
				}
			}
			pid := f.redisPID
			if pid == 0 {
				pid = 1
			}
			return []byte(fmt.Sprintf("run_id:%s\r\nprocess_id:%d\r\n", runID, pid)), nil
		}
		if args[len(args)-1] == "replication" {
			return []byte("role:master\r\nconnected_slaves:1\r\n"), nil
		}
		if args[len(args)-1] == "persistence" {
			return []byte("loading:0\r\naof_enabled:0\r\n"), nil
		}
		if args[len(args)-1] == "keyspace" {
			return []byte("# Keyspace\r\n"), nil
		}
		if args[len(args)-1] == "myid" {
			if f.clusterID != "" {
				return []byte(f.clusterID), nil
			}
			return []byte(args[1][24:] + "\n"), nil
		}
		if args[len(args)-1] == "info" {
			return []byte("cluster_state:ok\r\n"), nil
		}
		if args[len(args)-1] == "nodes" {
			return []byte("observed cluster nodes\n"), nil
		}
		return []byte("[OK] All 16384 slots covered."), nil
	default:
		return nil, fmt.Errorf("unexpected Docker operation: %q", args)
	}
}

func strconvPort(index int) string { return fmt.Sprintf("%d", 33000+index) }

func hasOperation(f *fakeDocker, operation string) bool {
	for _, args := range f.calls {
		if strings.Contains(strings.Join(args, " "), operation) {
			return true
		}
	}
	return false
}

func provision(t *testing.T, manager *Manager) *Environment {
	t.Helper()
	env, err := manager.Provision(context.Background(), testRun, Config{NodeTimeout: time.Second})
	if err != nil {
		t.Fatal(err)
	}
	return env
}

func TestProvisionIsolationAndPinnedMetadata(t *testing.T) {
	manager, fake := newFake()
	env := provision(t, manager)
	if len(env.Nodes) != 6 {
		t.Fatalf("got %d nodes", len(env.Nodes))
	}
	if env.Metadata["redis_image"] != DefaultImage || env.Metadata["redis_image_digests"] == "" || env.Metadata["docker_architecture"] == "" {
		t.Fatalf("missing metadata: %#v", env.Metadata)
	}
	for _, args := range fake.calls {
		if len(args) < 2 || args[0] != "--host" || args[1] != "unix:///tmp/docker.sock" {
			t.Fatalf("operation omitted explicit local socket: %q", args)
		}
		if args[2] == "run" {
			joined := strings.Join(args, " ")
			for _, flag := range []string{"--restart no", "--read-only", "--tmpfs /data:rw,noexec,nosuid", "--publish 127.0.0.1::6379", "--cap-drop ALL", "--security-opt no-new-privileges", "--memory 256m", "--user 999:999", "--cluster-enabled yes", "--cluster-node-timeout 1000"} {
				if !strings.Contains(joined, flag) {
					t.Fatalf("missing %s in %q", flag, args)
				}
			}
			for _, arg := range args {
				if arg == "--volume" || arg == "-v" || arg == "--privileged" || arg == "--network=host" {
					t.Fatalf("unsafe flag: %q", args)
				}
			}
		}
	}
	if !hasOperation(fake, "network create --driver bridge --opt "+noMasqueradeOption+"=false") || !hasOperation(fake, "--cluster-replicas 1 --cluster-yes") {
		t.Fatal("missing scoped bridge or cluster bootstrap")
	}
	for index, node := range env.Nodes {
		if node.ID != fmt.Sprintf("node-%d", index+1) || node.Addr != "127.0.0.1:"+strconvPort(index+1) || !strings.HasSuffix(node.InternalAddr, ":6379") {
			t.Fatalf("bad node addresses: %#v", node)
		}
	}
}

func TestRemoteDockerEndpointsRejectedBeforeAnyOperation(t *testing.T) {
	for _, host := range []string{"tcp://127.0.0.1:2375", "ssh://host", "unix://host/socket", "unix:relative", "unix:///tmp/socket?forward=tcp", "unix:///tmp/socket#target", "https://host", ""} {
		t.Run(host, func(t *testing.T) {
			manager, fake := newFake()
			manager.getenv = func(string) string { return host }
			if host == "" {
				manager.execute = func(context.Context, ...string) ([]byte, error) { return []byte("ssh://remote"), nil }
			}
			_, err := manager.Provision(context.Background(), testRun, Config{})
			if err == nil {
				t.Fatal("accepted remote Docker endpoint")
			}
			if len(fake.calls) != 0 {
				t.Fatalf("Docker target touched: %q", fake.calls)
			}
		})
	}
}

func TestEndpointResolvedOnceAndPassedExplicitly(t *testing.T) {
	manager, fake := newFake()
	manager.getenv = func(string) string { return "" }
	env := provision(t, manager)
	manager.getenv = func(string) string { return "ssh://changed-context" }
	if err := manager.Kill(context.Background(), env, env.Nodes[0]); err != nil {
		t.Fatal(err)
	}
	contextCalls := 0
	for _, args := range fake.calls {
		if args[0] == "context" {
			contextCalls++
			continue
		}
		if args[0] != "--host" || args[1] != "unix:///tmp/docker.sock" {
			t.Fatalf("lost endpoint binding: %q", args)
		}
	}
	if contextCalls != 1 {
		t.Fatalf("resolved endpoint %d times", contextCalls)
	}
}

func TestInvalidConfigurationDoesNotTouchDocker(t *testing.T) {
	for _, test := range []struct {
		run string
		cfg Config
	}{
		{"../outside", Config{}}, {"short", Config{}}, {"ValidButUppercase", Config{}},
		{testRun, Config{Image: "redis:latest"}}, {testRun, Config{Image: "host/redis:8.10.2-alpine"}},
		{testRun, Config{NodeTimeout: time.Nanosecond}}, {testRun, Config{NodeTimeout: time.Hour}},
	} {
		manager, fake := newFake()
		if _, err := manager.Provision(context.Background(), test.run, test.cfg); err == nil {
			t.Fatalf("accepted invalid config: %#v", test)
		}
		if len(fake.calls) != 0 {
			t.Fatal("invalid configuration executed Docker")
		}
	}
}

func TestProvisionFailureTracksResourcesIncludingLostOutput(t *testing.T) {
	for _, lost := range []bool{false, true} {
		t.Run(fmt.Sprint(lost), func(t *testing.T) {
			manager, fake := newFake()
			fake.failRun = 3
			fake.loseRunOutput = lost
			env, err := manager.Provision(context.Background(), testRun, Config{})
			if err == nil {
				t.Fatal("expected injected start failure")
			}
			expected := 2
			if lost {
				expected = 3
			}
			if env.NetworkID == "" || len(env.Nodes) != expected {
				t.Fatalf("partial inventory lost resources: %#v", env)
			}
			if err = manager.Cleanup(context.Background(), env); err != nil {
				t.Fatal(err)
			}
			if len(fake.containers) != 0 || fake.networkExists {
				t.Fatal("partial resources leaked")
			}
		})
	}
}

func TestKillRejectsUntrackedAndChangedOwnership(t *testing.T) {
	for _, fault := range []string{"untracked", "wrong-run-label", "wrong-network", "bind-mount", "public-port", "read-write-root", "different-name"} {
		t.Run(fault, func(t *testing.T) {
			manager, fake := newFake()
			env := provision(t, manager)
			node := env.Nodes[0]
			container := fake.containers[node.ContainerID]
			switch fault {
			case "untracked":
				node.ContainerID = strings.Repeat("f", 64)
			case "wrong-run-label":
				container.Config.Labels[runLabel] = "another-run"
			case "wrong-network":
				container.HostConfig.NetworkMode = strings.Repeat("b", 64)
			case "bind-mount":
				container.Mounts = []mountInspect{{Type: "bind"}}
			case "public-port":
				container.NetworkSettings.Ports["6379/tcp"][0].HostIP = "0.0.0.0"
			case "read-write-root":
				container.HostConfig.ReadonlyRootfs = false
			case "different-name":
				container.Name = "/some-other-container"
			}
			fake.containers[env.Nodes[0].ContainerID] = container
			if err := manager.Kill(context.Background(), env, node); err == nil {
				t.Fatalf("kill accepted %s", fault)
			}
			if hasOperation(fake, "container kill") {
				t.Fatal("unsafe kill executed")
			}
		})
	}
}

func TestCleanupIdempotentAfterFault(t *testing.T) {
	manager, fake := newFake()
	env := provision(t, manager)
	if err := manager.Kill(context.Background(), env, env.Nodes[0]); err != nil {
		t.Fatal(err)
	}
	if err := manager.Cleanup(context.Background(), env); err != nil {
		t.Fatal(err)
	}
	if err := manager.Cleanup(context.Background(), env); err != nil {
		t.Fatal(err)
	}
	if len(fake.containers) != 0 || fake.networkExists {
		t.Fatal("managed resources remain")
	}
	removals := 0
	for _, args := range fake.calls {
		if len(args) > 3 && args[2] == "container" && args[3] == "rm" {
			removals++
			if !idPattern.MatchString(args[len(args)-1]) {
				t.Fatalf("removal was not by full ID: %q", args)
			}
		}
	}
	if removals != 6 {
		t.Fatalf("got %d container removals", removals)
	}
}

func TestCleanupRefusesUnmanagedNetworkAndVerifiesDeletion(t *testing.T) {
	manager, fake := newFake()
	env := provision(t, manager)
	fake.network.Labels[managedLabel] = "false"
	if err := manager.Cleanup(context.Background(), env); err == nil {
		t.Fatal("accepted unmanaged network")
	}
	if hasOperation(fake, "container rm") || hasOperation(fake, "network rm") {
		t.Fatal("unsafe cleanup executed")
	}
	fake.network.Labels[managedLabel] = "true"
	fake.leaveAfterRemove = true
	if err := manager.Cleanup(context.Background(), env); err == nil {
		t.Fatal("cleanup accepted containers remaining after removal")
	}
	if hasOperation(fake, "network rm") {
		t.Fatal("network removed despite failed container cleanup")
	}
}

func TestDiscoverValidatesNamesAndIsolation(t *testing.T) {
	manager, fake := newFake()
	expected := provision(t, manager)
	env, err := manager.Discover(context.Background(), testRun)
	if err != nil || env.NetworkID != expected.NetworkID || len(env.Nodes) != 6 {
		t.Fatalf("discovery: %#v %v", env, err)
	}
	container := fake.containers[expected.Nodes[0].ContainerID]
	container.Name = "/another-service"
	fake.containers[container.ID] = container
	if _, err = manager.Discover(context.Background(), testRun); err == nil {
		t.Fatal("discovered a container with unexpected name")
	}
	if hasOperation(fake, "container rm") || hasOperation(fake, "container kill") {
		t.Fatal("discovery mutated resources")
	}
}

func TestExistingRunRefusedWithoutTransferringOwnership(t *testing.T) {
	manager, fake := newFake()
	provision(t, manager)
	env, err := manager.Provision(context.Background(), testRun, Config{})
	if err == nil {
		t.Fatal("accepted existing run")
	}
	if env.NetworkID != "" || len(env.Nodes) != 0 {
		t.Fatal("preexisting resources transferred to failed provision caller")
	}
	if fake.runCount != 6 {
		t.Fatal("existing run started additional nodes")
	}
}

func TestMissingDetectionDoesNotHideDaemonFailures(t *testing.T) {
	for _, err := range []error{errors.New("docker executable file not found"), errors.New("connection refused"), context.Canceled, errors.New("permission denied")} {
		if isMissing(err) {
			t.Fatalf("failure mistaken for resource absence: %v", err)
		}
	}
}

func TestReadinessCancellationReturnsInventoryForFreshCleanup(t *testing.T) {
	manager, fake := newFake()
	fake.pingFailure = true
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()
	env, err := manager.Provision(ctx, testRun, Config{})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("readiness did not honor deadline: %v", err)
	}
	if len(env.Nodes) != 6 {
		t.Fatal("cancelled readiness lost the resource inventory")
	}
	if err := manager.Cleanup(context.Background(), env); err != nil {
		t.Fatal(err)
	}
	if fake.networkExists || len(fake.containers) != 0 {
		t.Fatal("cancelled readiness leaked resources")
	}
}

func TestCleanupRefusesContainerMovedOutOfRunNetwork(t *testing.T) {
	manager, fake := newFake()
	env := provision(t, manager)
	container := fake.containers[env.Nodes[0].ContainerID]
	container.HostConfig.NetworkMode = strings.Repeat("b", 64)
	fake.containers[container.ID] = container
	if err := manager.Cleanup(context.Background(), env); err == nil {
		t.Fatal("cleanup accepted a moved container")
	}
	if _, stillExists := fake.containers[container.ID]; !stillExists {
		t.Fatal("cleanup deleted an out-of-scope container")
	}
	if !fake.networkExists {
		t.Fatal("cleanup removed a network despite a scope violation")
	}
}

func TestCleanupLegacyInternalBridgeWithMissingDynamicPorts(t *testing.T) {
	manager, fake := newFake()
	env := provision(t, manager)
	fake.network.Internal = true
	fake.network.Options = nil
	for id, container := range fake.containers {
		container.NetworkSettings.Ports = nil
		fake.containers[id] = container
	}
	discovered, err := manager.Discover(context.Background(), testRun)
	if err != nil {
		t.Fatal(err)
	}
	if len(discovered.Nodes) != 6 {
		t.Fatal("discovery lost partially provisioned containers")
	}
	if err := manager.Kill(context.Background(), discovered, discovered.Nodes[0]); err == nil {
		t.Fatal("kill accepted missing live Redis binding")
	}
	if err := manager.Cleanup(context.Background(), env); err != nil {
		t.Fatal(err)
	}
	if fake.networkExists || len(fake.containers) != 0 {
		t.Fatal("legacy partial environment leaked")
	}
}

func TestCleanupRejectsConfiguredPublicPortEvenIfStopped(t *testing.T) {
	manager, fake := newFake()
	env := provision(t, manager)
	node := env.Nodes[0]
	container := fake.containers[node.ContainerID]
	container.State.Running = false
	container.NetworkSettings.Ports = nil
	container.HostConfig.PortBindings["6379/tcp"][0].HostIP = "0.0.0.0"
	fake.containers[node.ContainerID] = container
	if err := manager.Cleanup(context.Background(), env); err == nil {
		t.Fatal("cleanup accepted public configured port publication")
	}
	if _, exists := fake.containers[node.ContainerID]; !exists {
		t.Fatal("cleanup removed altered container")
	}
}

func TestDiscoverAndCleanupCreatedContainerWithoutNetworkEndpoint(t *testing.T) {
	manager, fake := newFake()
	env := provision(t, manager)
	container := fake.containers[env.Nodes[0].ContainerID]
	container.State.Running = false
	container.NetworkSettings.Networks = nil
	container.NetworkSettings.Ports = nil
	fake.containers[container.ID] = container
	discovered, err := manager.Discover(context.Background(), testRun)
	if err != nil {
		t.Fatal(err)
	}
	if len(discovered.Nodes) != 6 {
		t.Fatal("created partial resource was not inventoried")
	}
	if err := manager.Cleanup(context.Background(), discovered); err != nil {
		t.Fatal(err)
	}
	if len(fake.containers) != 0 || fake.networkExists {
		t.Fatal("created resource leaked")
	}
}

func TestDockerSubprocessKeepsWarningsOutOfProtocolOutput(t *testing.T) {
	for _, protocol := range []struct{ name, output string }{
		{"resource-id", strings.Repeat("a", 64)},
		{"json", `{"Id":"managed-resource","State":{"Running":true}}`},
	} {
		t.Run(protocol.name, func(t *testing.T) {
			binaryDir := t.TempDir()
			stub := "#!/bin/sh\nprintf '%s\\n' '" + protocol.output + "'\nprintf '%s\\n' 'WARNING: harmless Docker diagnostic' >&2\n"
			if err := os.WriteFile(filepath.Join(binaryDir, "docker"), []byte(stub), 0755); err != nil {
				t.Fatal(err)
			}
			t.Setenv("PATH", binaryDir)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			output, err := executeDocker(ctx, "container", "inspect", "managed-id")
			if err != nil {
				t.Fatal(err)
			}
			if strings.TrimSpace(string(output)) != protocol.output {
				t.Fatalf("diagnostic corrupted protocol: %q", output)
			}
		})
	}
}

func TestDockerSubprocessRetainsFailureDiagnostics(t *testing.T) {
	binaryDir := t.TempDir()
	stub := "#!/bin/sh\nprintf '%s\\n' 'partial stdout'\nprintf '%s\\n' 'Error: No such object: managed-id' >&2\nexit 1\n"
	if err := os.WriteFile(filepath.Join(binaryDir, "docker"), []byte(stub), 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", binaryDir)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	output, err := executeDocker(ctx, "container", "inspect", "managed-id")
	if err == nil || !strings.Contains(err.Error(), "No such object") || !strings.Contains(err.Error(), "partial stdout") {
		t.Fatalf("lost failure diagnostics: %q, %v", output, err)
	}
	if string(output) != "partial stdout\n" {
		t.Fatalf("stdout mixed with stderr on failure: %q", output)
	}
}
