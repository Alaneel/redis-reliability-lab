package environment

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

func provisionRetained(t *testing.T, manager *Manager, mode string) *Environment {
	t.Helper()
	env, err := manager.Provision(context.Background(), testRun, Config{RetainData: true, Persistence: mode})
	if err != nil {
		t.Fatal(err)
	}
	return env
}

func TestRetainedProvisionUsesOwnedVolumesAndFixedEndpointsWithoutIsolationPermission(t *testing.T) {
	manager, fake := newFake()
	env := provisionRetained(t, manager, "memory")
	if !env.RetainData || env.NetworkIsolation || len(env.Volumes) != 6 || fake.network.Labels[isolationLabel] != "" {
		t.Fatalf("wrong retained capabilities: %#v", env)
	}
	if env.Metadata["storage_bound"] == "" || !strings.Contains(env.Metadata["container_data"], "no filesystem quota") {
		t.Fatal("storage limitations missing")
	}
	for _, node := range env.Nodes {
		container := fake.containers[node.ContainerID]
		if len(node.RedisClusterID) != 40 || node.RedisClusterID != node.ContainerID[24:] {
			t.Fatalf("initial Redis Cluster identity absent: %#v", node)
		}
		if container.Config.Labels[originalIPLabel] == "" || container.HostConfig.PortBindings["6379/tcp"][0].HostPort == "" {
			t.Fatal("retained endpoint not fixed")
		}
		if len(container.Mounts) != 1 || container.Mounts[0].Name != volumeName(env.RunID, node.ID) {
			t.Fatal("data mount not explicit owned volume")
		}
	}
	for _, call := range fake.calls {
		if len(call) > 2 && call[2] == "run" {
			text := strings.Join(call, " ")
			for _, flag := range []string{"--mount type=volume,src=redis-lab-abc12345-node-", "--appendonly no", "--save ", "--maxmemory 128mb", "--maxmemory-policy noeviction", "--auto-aof-rewrite-percentage 0", "--repl-diskless-sync yes", "--repl-diskless-load swapdb", "--read-only", "--user 999:999"} {
				if !strings.Contains(text, flag) {
					t.Fatalf("missing %s in %q", flag, call)
				}
			}
			if strings.Contains(text, "--tmpfs") || strings.Contains(text, "--privileged") {
				t.Fatalf("unexpected retained storage or privilege: %q", call)
			}
		}
	}
	if _, err := manager.Isolate(context.Background(), env, env.Nodes[0]); err == nil {
		t.Fatal("retained storage incorrectly grants network isolation capability")
	}
}

func TestRetainedRestartPreservesClusterIdentityAndEndpointsWithNewProcess(t *testing.T) {
	manager, fake := newFake()
	env := provisionRetained(t, manager, "memory")
	node := env.Nodes[0]
	live, err := manager.ProbeProcess(context.Background(), env, node)
	if err != nil {
		t.Fatal(err)
	}
	if err = manager.Kill(context.Background(), env, node); err != nil {
		t.Fatal(err)
	}
	// Real Docker clears a stopped endpoint's runtime IP/network ID and port
	// publication. The configured fixed creation endpoints remain authoritative.
	stopped := fake.containers[node.ContainerID]
	member := stopped.NetworkSettings.Networks[fake.network.Name]
	member.NetworkID, member.IPAddress = "", ""
	stopped.NetworkSettings.Networks[fake.network.Name] = member
	stopped.NetworkSettings.Ports = nil
	fake.containers[node.ContainerID] = stopped
	receipt, err := manager.Restart(context.Background(), env, node)
	if err != nil {
		t.Fatal(err)
	}
	if receipt.Before.ContainerRunning || !receipt.After.ContainerRunning || !receipt.After.RedisPONG || receipt.After.RedisClusterID != node.RedisClusterID || receipt.After.RedisRunID == live.RedisRunID || receipt.After.ContainerStartedAt == live.ContainerStartedAt || receipt.CompletedAt.IsZero() {
		t.Fatalf("invalid restart evidence: %#v", receipt)
	}
	if receipt.After.RawPersistence == "" || receipt.After.RawPersistenceConfig == "" || receipt.After.DataSizeKiB != 8 {
		t.Fatal("runtime persistence/storage observation absent")
	}
	if countOperation(fake, "container start") != 1 || countOperation(fake, "network connect") != 0 || countOperation(fake, "run --") != 6 {
		t.Fatal("restart recreated or reconnected resources")
	}
	if err = manager.Cleanup(context.Background(), env); err != nil {
		t.Fatal(err)
	}
	if len(fake.containers) != 0 || len(fake.volumes) != 0 || fake.networkExists {
		t.Fatal("restart resources leaked")
	}
}

func TestStoppedProcessEvidenceRequiresSuccessfulOwnershipInspection(t *testing.T) {
	manager, fake := newFake()
	env := provisionRetained(t, manager, "memory")
	node := env.Nodes[0]
	if err := manager.Kill(context.Background(), env, node); err != nil {
		t.Fatal(err)
	}
	evidence, err := manager.ProbeProcess(context.Background(), env, node)
	if !errors.Is(err, ErrProcessStopped) || evidence.ContainerRunning || evidence.NetworkConnected || evidence.ObservedAt.IsZero() {
		t.Fatalf("stopped state not established: %#v %v", evidence, err)
	}
	container := fake.containers[node.ContainerID]
	container.Config.Labels[runLabel] = "foreign-run"
	fake.containers[node.ContainerID] = container
	_, err = manager.ProbeProcess(context.Background(), env, node)
	if err == nil || errors.Is(err, ErrProcessStopped) {
		t.Fatalf("ownership failure misclassified as stopped: %v", err)
	}
}

func TestRestartRefusesUnsafeResourcesBeforeStart(t *testing.T) {
	for _, fault := range []string{"ordinary", "untracked", "running", "cluster-id-absent", "wrong-label", "privileged", "wrong-volume-label", "volume-driver", "volume-options", "volume-name", "volume-destination", "extra-volume", "host-bind", "wrong-port", "wrong-IP", "wrong-network", "detached-running"} {
		t.Run(fault, func(t *testing.T) {
			manager, fake := newFake()
			env := provisionRetained(t, manager, "memory")
			node := env.Nodes[0]
			if fault != "running" {
				if err := manager.Kill(context.Background(), env, node); err != nil {
					t.Fatal(err)
				}
			}
			container := fake.containers[node.ContainerID]
			volume := fake.volumes[volumeName(env.RunID, node.ID)]
			switch fault {
			case "ordinary":
				env.RetainData = false
			case "untracked":
				node.ContainerID = strings.Repeat("f", 64)
			case "cluster-id-absent":
				node.RedisClusterID = ""
			case "wrong-label":
				container.Config.Labels[runLabel] = "some-other-run"
			case "privileged":
				container.HostConfig.Privileged = true
			case "wrong-volume-label":
				volume.Labels[runLabel] = "some-other-run"
			case "volume-driver":
				volume.Driver = "remote-plugin"
			case "volume-options":
				volume.Options = map[string]string{"device": "/Users/foreign-data", "type": "none", "o": "bind"}
			case "volume-name":
				container.Mounts[0].Name = "foreign-volume"
			case "volume-destination":
				container.Mounts[0].Destination = "/etc"
			case "extra-volume":
				container.Mounts = append(container.Mounts, mountInspect{Type: "volume", Name: "foreign", Destination: "/foreign", RW: true})
			case "host-bind":
				container.Mounts = append(container.Mounts, mountInspect{Type: "bind", Destination: "/host"})
			case "wrong-port":
				container.HostConfig.PortBindings["6379/tcp"][0].HostPort = "44444"
			case "wrong-IP":
				member := container.NetworkSettings.Networks[fake.network.Name]
				member.IPAddress = "10.224.0.99"
				container.NetworkSettings.Networks[fake.network.Name] = member
			case "wrong-network":
				container.HostConfig.NetworkMode = strings.Repeat("f", 64)
			case "detached-running":
				container.State.Running = true
				container.NetworkSettings.Networks = nil
			}
			fake.containers[env.Nodes[0].ContainerID] = container
			fake.volumes[volume.Name] = volume
			if _, err := manager.Restart(context.Background(), env, node); err == nil {
				t.Fatal("unsafe restart accepted")
			}
			if countOperation(fake, "container start") != 0 {
				t.Fatal("unsafe start executed")
			}
		})
	}
}

func TestRestartDetectsChangedClusterIdentityAndUnexpectedEndpointAfterStart(t *testing.T) {
	for _, fault := range []string{"cluster-identity", "port", "IP", "start-noop"} {
		t.Run(fault, func(t *testing.T) {
			manager, fake := newFake()
			env := provisionRetained(t, manager, "memory")
			node := env.Nodes[0]
			if err := manager.Kill(context.Background(), env, node); err != nil {
				t.Fatal(err)
			}
			original := manager.execute
			manager.execute = func(ctx context.Context, args ...string) ([]byte, error) {
				out, err := original(ctx, args...)
				if strings.Contains(strings.Join(args, " "), "container start") {
					container := fake.containers[node.ContainerID]
					switch fault {
					case "cluster-identity":
						fake.clusterID = strings.Repeat("e", 40)
					case "port":
						container.NetworkSettings.Ports["6379/tcp"][0].HostPort = "44444"
					case "IP":
						member := container.NetworkSettings.Networks[fake.network.Name]
						member.IPAddress = "10.224.0.99"
						container.NetworkSettings.Networks[fake.network.Name] = member
					}
					fake.containers[node.ContainerID] = container
				}
				return out, err
			}
			fake.startNoop = fault == "start-noop"
			receipt, err := manager.Restart(context.Background(), env, node)
			if err == nil || receipt == nil || !receipt.CompletedAt.IsZero() {
				t.Fatalf("unexpected restart accepted: %#v %v", receipt, err)
			}
		})
	}
}

func TestAmbiguousRestartAndCancellationPreserveAppliedStartEvidenceAndError(t *testing.T) {
	for _, fault := range []string{"lost-response", "cancel-after-start", "cancel-during-observation", "deadline-after-start"} {
		t.Run(fault, func(t *testing.T) {
			manager, fake := newFake()
			env := provisionRetained(t, manager, "memory")
			node := env.Nodes[0]
			if err := manager.Kill(context.Background(), env, node); err != nil {
				t.Fatal(err)
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if fault == "deadline-after-start" {
				ctx, cancel = context.WithTimeout(context.Background(), 20*time.Millisecond)
				defer cancel()
			}
			lost := errors.New("Docker start response lost")
			original := manager.execute
			started, canceled := false, false
			manager.execute = func(callCtx context.Context, args ...string) ([]byte, error) {
				out, err := original(callCtx, args...)
				text := strings.Join(args, " ")
				if strings.Contains(text, "container start") {
					started = true
					switch fault {
					case "lost-response":
						return out, lost
					case "cancel-after-start":
						cancel()
						return out, context.Canceled
					case "deadline-after-start":
						<-callCtx.Done()
						return out, callCtx.Err()
					}
				}
				if fault == "cancel-during-observation" && started && !canceled && strings.HasSuffix(text, "ping") {
					canceled = true
					cancel()
				}
				return out, err
			}
			receipt, err := manager.Restart(ctx, env, node)
			if err == nil || receipt == nil || !receipt.After.RedisPONG || receipt.After.RedisClusterID != node.RedisClusterID || receipt.CompletedAt.IsZero() {
				t.Fatalf("applied start evidence lost: %#v %v", receipt, err)
			}
			want := lost
			if fault == "deadline-after-start" {
				want = context.DeadlineExceeded
			} else if fault != "lost-response" {
				want = context.Canceled
			}
			if !errors.Is(err, want) {
				t.Fatalf("start error hidden: %v", err)
			}
			if countOperation(fake, "container start") != 1 {
				t.Fatal("ambiguous start replayed")
			}
			if err = manager.Cleanup(context.Background(), env); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestRetainedPartialProvisionAndLostVolumeOutputAreDiscoverableAndCleanable(t *testing.T) {
	for _, fault := range []string{"volume-failure", "lost-volume-response", "container-failure", "lost-container-response"} {
		t.Run(fault, func(t *testing.T) {
			manager, fake := newFake()
			if strings.Contains(fault, "volume") {
				fake.failVolume = 3
				fake.loseVolumeOutput = strings.HasPrefix(fault, "lost-")
			} else {
				fake.failRun = 3
				fake.loseRunOutput = strings.HasPrefix(fault, "lost-")
			}
			env, err := manager.Provision(context.Background(), testRun, Config{RetainData: true})
			if err == nil || env.NetworkID == "" || len(env.Volumes) < 2 {
				t.Fatalf("partial retained inventory lost: %#v %v", env, err)
			}
			discovered, err := manager.Discover(context.Background(), testRun)
			if err != nil || len(discovered.Volumes) != len(fake.volumes) {
				t.Fatalf("volume discovery failed: %#v %v", discovered, err)
			}
			if err = manager.Cleanup(context.Background(), env); err != nil {
				t.Fatal(err)
			}
			if len(fake.volumes) != 0 || len(fake.containers) != 0 || fake.networkExists {
				t.Fatal("persistent partial resources leaked")
			}
		})
	}
}

func TestVolumeOnlyDiscoveryCleanupAndExistingRunRefusal(t *testing.T) {
	manager, fake := newFake()
	env := &Environment{RunID: testRun, RetainData: true, Persistence: "memory", Volumes: []Volume{{Name: volumeName(testRun, "node-1"), NodeID: "node-1"}}}
	if err := manager.createVolume(context.Background(), env, env.Volumes[0]); err != nil {
		t.Fatal(err)
	}
	discovered, err := manager.Discover(context.Background(), testRun)
	if err != nil || discovered.NetworkID != "" || len(discovered.Volumes) != 1 {
		t.Fatalf("volume-only inventory lost: %#v %v", discovered, err)
	}
	partial, err := manager.Provision(context.Background(), testRun, Config{RetainData: true})
	if err == nil || len(partial.Volumes) != 0 || countOperation(fake, "network create") != 0 {
		t.Fatal("existing volume ownership transferred or ignored")
	}
	if err = manager.Cleanup(context.Background(), discovered); err != nil {
		t.Fatal(err)
	}
	if err = manager.Cleanup(context.Background(), discovered); err != nil {
		t.Fatal(err)
	}
	if len(fake.volumes) != 0 {
		t.Fatal("volume-only cleanup leaked")
	}
}

func TestVolumeCleanupRefusesForeignOwnershipAndVerifiesRemovalWithoutForce(t *testing.T) {
	for _, fault := range []string{"foreign-label", "driver-options", "still-present", "foreign-user"} {
		t.Run(fault, func(t *testing.T) {
			manager, fake := newFake()
			env := provisionRetained(t, manager, "memory")
			first := env.Volumes[0]
			volume := fake.volumes[first.Name]
			switch fault {
			case "foreign-label":
				volume.Labels[runLabel] = "some-other-run"
			case "driver-options":
				volume.Options = map[string]string{"device": "/foreign"}
			case "still-present":
				fake.leaveVolumeAfterRemove = true
			case "foreign-user":
				fake.volumeInUse = true
			}
			fake.volumes[first.Name] = volume
			if err := manager.Cleanup(context.Background(), env); err == nil {
				t.Fatal("unsafe/incomplete volume cleanup reported success")
			}
			if !fake.networkExists {
				t.Fatal("network cleanup hid volume failure")
			}
			for _, call := range fake.calls {
				if strings.Contains(strings.Join(call, " "), "volume rm --force") {
					t.Fatal("volume forced out of foreign use")
				}
			}
			if fault == "foreign-label" || fault == "driver-options" {
				if _, ok := fake.containers[env.Nodes[0].ContainerID]; !ok {
					t.Fatal("container using unowned volume removed")
				}
			}
		})
	}
}

func TestPersistenceAllowlistAndActualRuntimeConfig(t *testing.T) {
	for _, mode := range []string{"memory", "aof-everysec", "aof-always"} {
		t.Run(mode, func(t *testing.T) {
			manager, fake := newFake()
			env := provisionRetained(t, manager, mode)
			evidence, err := manager.ProbeProcess(context.Background(), env, env.Nodes[0])
			if err != nil {
				t.Fatal(err)
			}
			if err = validatePersistenceConfig(evidence.RawPersistenceConfig, mode); err != nil {
				t.Fatal(err)
			}
			if mode != "memory" && (!hasOperation(fake, "--appendonly yes") || !hasOperation(fake, "--appendfsync "+strings.TrimPrefix(mode, "aof-")) || !hasOperation(fake, "--aof-use-rdb-preamble yes")) {
				t.Fatal("AOF mode not explicitly configured")
			}
			wrong := strings.Replace(evidence.RawPersistenceConfig, "noeviction", "allkeys-lru", 1)
			if validatePersistenceConfig(wrong, mode) == nil {
				t.Fatal("altered runtime persistence configuration accepted")
			}
		})
	}
	for _, cfg := range []Config{{RetainData: true, Persistence: "rdb"}, {Persistence: "aof-everysec"}, {RetainData: true, Persistence: "--appendonly yes"}} {
		manager, fake := newFake()
		if _, err := manager.Provision(context.Background(), testRun, cfg); err == nil {
			t.Fatal("invalid persistence accepted")
		}
		if len(fake.calls) != 0 {
			t.Fatal("invalid persistence touched Docker")
		}
	}
}

func TestPersistenceConfigPreservesEmptySaveInAnyReplyPosition(t *testing.T) {
	fields := []string{"appendonly", "no", "appendfsync", "everysec", "repl-diskless-load", "swapdb", "repl-diskless-sync", "yes", "auto-aof-rewrite-percentage", "0", "maxmemory", "134217728", "maxmemory-policy", "noeviction", "aof-use-rdb-preamble", "no", "no-appendfsync-on-rewrite", "no", "aof-load-truncated", "yes", "rdb-del-sync-files", "yes"}
	for _, test := range []struct {
		name, reply string
		valid       bool
	}{
		{"save-last-unix", strings.Join(fields, "\n") + "\nsave\n\n", true},
		{"save-last-crlf", strings.Join(fields, "\r\n") + "\r\nsave\r\n\r\n", true},
		{"empty-in-middle", "save\n\n" + strings.Join(fields, "\n") + "\n", true},
		{"missing-final-value", strings.Join(fields, "\n") + "\nsave\n", false},
		{"malformed-odd", "appendonly\nno\nsave", false},
	} {
		t.Run(test.name, func(t *testing.T) {
			parsed := persistenceConfigReply([]byte(test.reply))
			err := validatePersistenceConfig(parsed, "memory")
			if (err == nil) != test.valid {
				t.Fatalf("reply=%q parsed=%q error=%v", test.reply, parsed, err)
			}
			if strings.HasPrefix(test.name, "save-last") && !strings.HasSuffix(parsed, "save\n") {
				t.Fatalf("final empty save value lost: %q", parsed)
			}
		})
	}
}

func TestProcessProbePreservesFinalEmptyConfigValueAndOtherProbeNormalization(t *testing.T) {
	manager, fake := newFake()
	env := provisionRetained(t, manager, "memory")
	original := manager.execute
	manager.execute = func(ctx context.Context, args ...string) ([]byte, error) {
		out, err := original(ctx, args...)
		text := strings.Join(args, " ")
		if strings.Contains(text, "config get") {
			// Move the empty value to the end, as real Redis CONFIG GET may do.
			normalized := strings.Replace(string(out), "save\n\n", "", 1)
			out = []byte(strings.ReplaceAll(normalized+"save\n\n", "\n", "\r\n"))
		}
		if strings.HasSuffix(text, "ping") {
			out = []byte(" PONG\r\n ")
		}
		return out, err
	}
	evidence, err := manager.ProbeProcess(context.Background(), env, env.Nodes[0])
	if err != nil || !evidence.RedisPONG || !strings.HasSuffix(evidence.RawPersistenceConfig, "save\n") {
		t.Fatalf("raw config or other probe changed: %#v %v", evidence, err)
	}
	if countOperation(fake, "config get") != 1 {
		t.Fatal("fixed CONFIG GET probe missing or duplicated")
	}
}
