package environment

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"strings"
	"testing"
	"time"
)

func provisionIsolation(t *testing.T, manager *Manager) *Environment {
	t.Helper()
	env, err := manager.Provision(context.Background(), testRun, Config{NetworkIsolation: true})
	if err != nil {
		t.Fatal(err)
	}
	return env
}

func isolate(t *testing.T, manager *Manager, env *Environment) *IsolationReceipt {
	t.Helper()
	receipt, err := manager.Isolate(context.Background(), env, env.Nodes[0])
	if err != nil {
		t.Fatal(err)
	}
	return receipt
}

func countOperation(fake *fakeDocker, operation string) int {
	count := 0
	for _, args := range fake.calls {
		if strings.Contains(strings.Join(args, " "), operation) {
			count++
		}
	}
	return count
}

func TestIsolationProvisionUsesExplicitSafeEndpoints(t *testing.T) {
	manager, fake := newFake()
	env := provisionIsolation(t, manager)
	if !env.NetworkIsolation || !validIsolationSubnet(env.Metadata["network_subnet"]) || fake.network.Labels[isolationLabel] != "true" {
		t.Fatalf("missing managed isolation IPAM: %#v", env)
	}
	if err := validateNetwork(fake.network, env); err != nil {
		t.Fatal(err)
	}
	for _, node := range env.Nodes {
		container := fake.containers[node.ContainerID]
		if container.Config.Labels[isolationLabel] != "true" || container.Config.Labels[originalPortLabel] == "" || container.Config.Labels[originalIPLabel] == "" {
			t.Fatal("missing persistent endpoint capability labels")
		}
		ip, _, _ := net.SplitHostPort(node.InternalAddr)
		_, pool, _ := net.ParseCIDR(env.Metadata["network_subnet"])
		if !pool.Contains(net.ParseIP(ip)) || node.Addr != "127.0.0.1:"+container.Config.Labels[originalPortLabel] {
			t.Fatalf("wrong isolation endpoint: %#v", node)
		}
	}
	if hasOperation(fake, "127.0.0.1::6379") || !hasOperation(fake, "--subnet "+env.Metadata["network_subnet"]) {
		t.Fatal("isolation did not use explicit IPAM and host ports")
	}
}

func TestIsolationLeavesProcessAliveAndRestoresExactEndpointsIdempotently(t *testing.T) {
	manager, fake := newFake()
	env := provisionIsolation(t, manager)
	receipt := isolate(t, manager, env)
	process, err := manager.ProbeProcess(context.Background(), env, receipt.Node)
	if err != nil || !sameProcess(receipt.BeforeProcess, process) || process.NetworkConnected || process.NetworkCount != 0 || process.RawClusterInfo == "" || process.RawClusterNodes == "" || process.RawReplication == "" {
		t.Fatalf("bad isolated process evidence: %#v, %v", process, err)
	}
	if receipt.CommandStartedAt.IsZero() || receipt.CommandCompletedAt.Before(receipt.CommandStartedAt) || receipt.DisconnectedAt.Before(receipt.CommandCompletedAt) {
		t.Fatalf("missing fault timings: %#v", receipt)
	}
	if !hasOperation(fake, "network disconnect "+env.NetworkID+" "+receipt.Node.ContainerID) || hasOperation(fake, "network disconnect --force") {
		t.Fatal("disconnect was not the fixed exact-ID operation")
	}
	for i := 0; i < 2; i++ {
		if err := manager.Restore(context.Background(), env, receipt); err != nil {
			t.Fatal(err)
		}
	}
	if countOperation(fake, "network connect") != 1 || receipt.RestoredAt.IsZero() {
		t.Fatal("restore was not idempotent")
	}
	container := fake.containers[receipt.Node.ContainerID]
	current, err := containerNode(container, env, receipt.Node)
	if err != nil || current != receipt.Node {
		t.Fatalf("restore changed frozen endpoints: %#v, %v", current, err)
	}
}

func TestIsolationRefusesNonCapableAndAlteredResources(t *testing.T) {
	for _, fault := range []string{"ordinary-run", "untracked", "forged-address", "wrong-label", "wrong-network", "extra-network", "wrong-IP", "public-port", "read-write-root", "bind-mount", "changed-subnet", "unmanaged-network"} {
		t.Run(fault, func(t *testing.T) {
			manager, fake := newFake()
			var env *Environment
			if fault == "ordinary-run" {
				env = provision(t, manager)
			} else {
				env = provisionIsolation(t, manager)
			}
			node := env.Nodes[0]
			container := fake.containers[node.ContainerID]
			switch fault {
			case "untracked":
				node.ContainerID = strings.Repeat("f", 64)
			case "forged-address":
				node.Addr = "127.0.0.1:1234"
			case "wrong-label":
				container.Config.Labels[runLabel] = "other123"
			case "wrong-network":
				container.HostConfig.NetworkMode = strings.Repeat("b", 64)
			case "extra-network":
				container.NetworkSettings.Networks["foreign"] = struct{ NetworkID, IPAddress string }{strings.Repeat("b", 64), "192.168.1.2"}
			case "wrong-IP":
				member := container.NetworkSettings.Networks[fake.network.Name]
				member.IPAddress = "192.168.10.2"
				container.NetworkSettings.Networks[fake.network.Name] = member
			case "public-port":
				container.HostConfig.PortBindings["6379/tcp"][0].HostIP = "0.0.0.0"
			case "read-write-root":
				container.HostConfig.ReadonlyRootfs = false
			case "bind-mount":
				container.Mounts = []mountInspect{{Type: "bind"}}
			case "changed-subnet":
				fake.network.IPAM.Config[0].Subnet = "192.168.1.0/24"
			case "unmanaged-network":
				fake.network.Labels[managedLabel] = "false"
			}
			fake.containers[env.Nodes[0].ContainerID] = container
			if _, err := manager.Isolate(context.Background(), env, node); err == nil {
				t.Fatalf("accepted unsafe isolation: %s", fault)
			}
			if hasOperation(fake, "network disconnect") {
				t.Fatal("unsafe disconnect reached Docker")
			}
		})
	}
}

func TestDisconnectErrorRetainsReceiptAndReconcilesAppliedMutation(t *testing.T) {
	for _, cause := range []error{context.Canceled, context.DeadlineExceeded, errors.New("response lost after detach")} {
		manager, fake := newFake()
		env := provisionIsolation(t, manager)
		fake.disconnectError = cause
		receipt, err := manager.Isolate(context.Background(), env, env.Nodes[0])
		if !errors.Is(err, cause) || receipt == nil || receipt.DisconnectedAt.IsZero() || !receipt.BeforeProcess.RedisPONG {
			t.Fatalf("lost ambiguous mutation receipt/cause: %#v %v", receipt, err)
		}
		if err = manager.Restore(context.Background(), env, receipt); err != nil {
			t.Fatal(err)
		}
		if err = manager.Cleanup(context.Background(), env); err != nil {
			t.Fatal(err)
		}
	}
}

func TestDisconnectSuccessfulExitRequiresDetachedPostcondition(t *testing.T) {
	manager, fake := newFake()
	env := provisionIsolation(t, manager)
	fake.disconnectNoop = true
	receipt, err := manager.Isolate(context.Background(), env, env.Nodes[0])
	if err == nil || receipt == nil || !receipt.DisconnectedAt.IsZero() {
		t.Fatalf("accepted no-op disconnect: %#v %v", receipt, err)
	}
	if err = manager.Restore(context.Background(), env, receipt); err != nil {
		t.Fatal(err)
	}
	if hasOperation(fake, "network connect") {
		t.Fatal("reattached an already attached original endpoint")
	}
}

func TestRestoreErrorsReconcilePostconditionWithoutHidingCause(t *testing.T) {
	for _, applied := range []bool{false, true} {
		manager, fake := newFake()
		env := provisionIsolation(t, manager)
		receipt := isolate(t, manager, env)
		cause := errors.New("connect response lost")
		fake.connectError = cause
		fake.connectNoop = !applied
		err := manager.Restore(context.Background(), env, receipt)
		if !errors.Is(err, cause) || receipt.RestoredAt.IsZero() == applied {
			t.Fatalf("restore hid cause or misclassified attachment: applied=%v, %#v %v", applied, receipt, err)
		}
		if err = manager.Cleanup(context.Background(), env); err != nil {
			t.Fatal(err)
		}
	}
}

func TestRestoreRejectsChangedProcessEndpointsAndOccupiedIP(t *testing.T) {
	for _, fault := range []string{"container-pid", "container-started", "redis-run-id", "redis-pid", "stopped", "changed-port", "other-network", "occupied-IP", "receipt-run", "receipt-IP", "receipt-node"} {
		t.Run(fault, func(t *testing.T) {
			manager, fake := newFake()
			env := provisionIsolation(t, manager)
			receipt := isolate(t, manager, env)
			container := fake.containers[receipt.Node.ContainerID]
			switch fault {
			case "container-pid":
				container.State.Pid++
			case "container-started":
				container.State.StartedAt = "2026-10-11T01:00:00Z"
			case "redis-run-id":
				fake.redisRunID = strings.Repeat("d", 40)
			case "redis-pid":
				fake.redisPID = 2
			case "stopped":
				container.State.Running = false
			case "changed-port":
				container.HostConfig.PortBindings["6379/tcp"][0].HostPort = "44000"
			case "other-network":
				container.NetworkSettings.Networks = map[string]struct{ NetworkID, IPAddress string }{"foreign": {strings.Repeat("b", 64), receipt.OriginalIP}}
			case "occupied-IP":
				fake.foreignMembers = map[string]json.RawMessage{strings.Repeat("f", 64): json.RawMessage(`{"IPv4Address":"` + receipt.OriginalIP + `/24"}`)}
			case "receipt-run":
				receipt.RunID = "other123"
			case "receipt-IP":
				receipt.OriginalIP = "192.168.1.2"
			case "receipt-node":
				receipt.Node.ContainerID = strings.Repeat("f", 64)
			}
			fake.containers[env.Nodes[0].ContainerID] = container
			if err := manager.Restore(context.Background(), env, receipt); err == nil {
				t.Fatalf("accepted unsafe restore: %s", fault)
			}
			if hasOperation(fake, "network connect") {
				t.Fatal("unsafe restore reached Docker")
			}
		})
	}
}

func TestCancellationAfterDisconnectSideEffectStillTracksAndRestores(t *testing.T) {
	manager, fake := newFake()
	env := provisionIsolation(t, manager)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	manager.execute = func(callCtx context.Context, args ...string) ([]byte, error) {
		out, err := fake.execute(callCtx, args...)
		if len(args) > 3 && args[2] == "network" && args[3] == "disconnect" {
			cancel()
			return out, context.Canceled
		}
		return out, err
	}
	receipt, err := manager.Isolate(ctx, env, env.Nodes[0])
	if !errors.Is(err, context.Canceled) || receipt == nil || receipt.DisconnectedAt.IsZero() {
		t.Fatalf("canceled side effect lost receipt: %#v %v", receipt, err)
	}
	if err := manager.Restore(context.Background(), env, receipt); err != nil {
		t.Fatal(err)
	}
	if err := manager.Cleanup(context.Background(), env); err != nil {
		t.Fatal(err)
	}
}

func TestIsolationPartialProvisionAndLostNetworkOutputRemainCleanable(t *testing.T) {
	for _, fault := range []string{"lost-network-response", "lost-container-response", "port-selection"} {
		t.Run(fault, func(t *testing.T) {
			manager, fake := newFake()
			switch fault {
			case "lost-network-response":
				manager.execute = func(ctx context.Context, args ...string) ([]byte, error) {
					out, err := fake.execute(ctx, args...)
					if len(args) > 3 && args[2] == "network" && args[3] == "create" {
						return nil, errors.New("network creation response lost")
					}
					return out, err
				}
			case "lost-container-response":
				fake.failRun, fake.loseRunOutput = 3, true
			case "port-selection":
				manager.choosePort = func(map[string]bool) (string, error) { return "", errors.New("ephemeral port unavailable") }
			}
			env, err := manager.Provision(context.Background(), testRun, Config{NetworkIsolation: true})
			if err == nil || env.NetworkID == "" || !env.NetworkIsolation {
				t.Fatalf("partial isolation inventory lost: %#v %v", env, err)
			}
			if err = manager.Cleanup(context.Background(), env); err != nil || fake.networkExists || len(fake.containers) != 0 {
				t.Fatalf("partial isolation resources leaked: %v", err)
			}
		})
	}
}

func TestHardStopDiscoverAndCleanupDetachedNodeWithoutReconnect(t *testing.T) {
	manager, fake := newFake()
	env := provisionIsolation(t, manager)
	receipt := isolate(t, manager, env)
	newManager := &Manager{execute: fake.execute, getenv: manager.getenv}
	discovered, err := newManager.Discover(context.Background(), env.RunID)
	if err != nil || !discovered.NetworkIsolation || len(discovered.Nodes) != 6 || discovered.Nodes[0] != receipt.Node {
		t.Fatalf("hard-stop discovery lost detached endpoints: %#v %v", discovered, err)
	}
	for i := 0; i < 2; i++ {
		if err = newManager.Cleanup(context.Background(), discovered); err != nil {
			t.Fatal(err)
		}
	}
	if len(fake.containers) != 0 || fake.networkExists || hasOperation(fake, "network connect") {
		t.Fatal("detached cleanup leaked resources or reconnected unnecessarily")
	}
}

func TestOrdinaryRunningDetachedNodeCannotOptIntoCleanup(t *testing.T) {
	manager, fake := newFake()
	env := provision(t, manager)
	container := fake.containers[env.Nodes[0].ContainerID]
	container.NetworkSettings.Networks = nil
	container.NetworkSettings.Ports = nil
	fake.containers[container.ID] = container
	if _, err := manager.Discover(context.Background(), env.RunID); err == nil {
		t.Fatal("ordinary detached node accepted without capability")
	}
	if err := manager.Cleanup(context.Background(), env); err == nil {
		t.Fatal("ordinary detached node removed without capability")
	}
	if _, exists := fake.containers[container.ID]; !exists {
		t.Fatal("out-of-scope detached resource was deleted")
	}
}

func TestIsolationSubnetCollisionRetryBoundedAndPreservesUnknownMember(t *testing.T) {
	manager, fake := newFake()
	fake.overlapAttempts = 2
	env := provisionIsolation(t, manager)
	if env.Metadata["network_subnet"] != isolationSubnet(testRun, 2) || countOperation(fake, "network create") != 3 {
		t.Fatal("did not retry exact rejected IPAM candidates")
	}
	fake.foreignMembers = map[string]json.RawMessage{strings.Repeat("f", 64): json.RawMessage(`{"IPv4Address":"192.168.1.2/24"}`)}
	if err := manager.Cleanup(context.Background(), env); err == nil || hasOperation(fake, "network rm") || !fake.networkExists {
		t.Fatal("cleanup removed a network containing an unknown endpoint")
	}
	manager, fake = newFake()
	fake.overlapAttempts = 50
	if _, err := manager.Provision(context.Background(), testRun, Config{NetworkIsolation: true}); err == nil || countOperation(fake, "network create") != 32 || fake.networkExists {
		t.Fatal("subnet collision retry was not bounded")
	}
}

func TestIsolatedCancellationAllowsFreshBoundedCleanup(t *testing.T) {
	manager, fake := newFake()
	env := provisionIsolation(t, manager)
	receipt := isolate(t, manager, env)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := manager.Restore(ctx, env, receipt); !errors.Is(err, context.Canceled) {
		t.Fatalf("restore ignored cancellation: %v", err)
	}
	cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), time.Second)
	defer cleanupCancel()
	if err := manager.Cleanup(cleanupCtx, env); err != nil || len(fake.containers) != 0 || fake.networkExists {
		t.Fatalf("fresh detached cleanup failed: %v", err)
	}
}
