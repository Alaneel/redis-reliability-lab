package environment

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

func TestManagedOperationsRejectAlteredRestartPolicy(t *testing.T) {
	for _, policy := range []string{"always", "unless-stopped", "on-failure", "unknown", ""} {
		name := policy
		if name == "" {
			name = "missing-name"
		}
		for _, operation := range []string{"kill", "running-probe", "stopped-probe", "restart", "isolate", "restore", "discover", "cleanup"} {
			t.Run(name+"/"+operation, func(t *testing.T) {
				manager, fake := newFake()
				env, err := manager.Provision(context.Background(), testRun, Config{RetainData: true, Persistence: "memory", NetworkIsolation: true})
				if err != nil {
					t.Fatal(err)
				}
				node := env.Nodes[0]
				var receipt *IsolationReceipt
				if operation == "stopped-probe" || operation == "restart" {
					if err := manager.Kill(context.Background(), env, node); err != nil {
						t.Fatal(err)
					}
				}
				if operation == "restore" {
					receipt, err = manager.Isolate(context.Background(), env, node)
					if err != nil {
						t.Fatal(err)
					}
				}
				container := fake.containers[node.ContainerID]
				container.HostConfig.RestartPolicy.Name = policy
				fake.containers[node.ContainerID] = container
				fake.calls = nil
				switch operation {
				case "kill":
					err = manager.Kill(context.Background(), env, node)
				case "running-probe", "stopped-probe":
					_, err = manager.ProbeProcess(context.Background(), env, node)
					if errors.Is(err, ErrProcessStopped) {
						t.Fatal("unsafe restart policy was misclassified as a proven stopped process")
					}
				case "restart":
					_, err = manager.Restart(context.Background(), env, node)
				case "isolate":
					_, err = manager.Isolate(context.Background(), env, node)
				case "restore":
					err = manager.Restore(context.Background(), env, receipt)
				case "discover":
					_, err = manager.Discover(context.Background(), env.RunID)
				case "cleanup":
					err = manager.Cleanup(context.Background(), env)
				}
				if err == nil || !strings.Contains(err.Error(), "restart policy") {
					t.Fatalf("operation accepted or failed to inspect restart policy %q: %v", policy, err)
				}
				for _, call := range fake.calls {
					joined := strings.Join(call, " ")
					for _, mutation := range []string{"container kill", "container start", "container rm", "network disconnect", "network connect", "exec"} {
						if strings.Contains(joined, mutation) && strings.Contains(joined, node.ContainerID) {
							t.Fatalf("unsafe target operation reached Docker: %q", call)
						}
					}
				}
				if _, exists := fake.containers[node.ContainerID]; !exists {
					t.Fatal("unsafe target was removed")
				}
			})
		}
	}
}

func TestMissingRestartPolicyInspectionFailsClosed(t *testing.T) {
	manager, fake := newFake()
	env := provisionRetained(t, manager, "memory")
	node := env.Nodes[0]
	if err := manager.Kill(context.Background(), env, node); err != nil {
		t.Fatal(err)
	}
	manager.execute = func(ctx context.Context, args ...string) ([]byte, error) {
		out, err := fake.execute(ctx, args...)
		if err == nil && len(args) > 4 && args[2] == "container" && args[3] == "inspect" {
			var reply map[string]any
			if err := json.Unmarshal(out, &reply); err != nil {
				return nil, err
			}
			delete(reply["HostConfig"].(map[string]any), "RestartPolicy")
			return json.Marshal(reply)
		}
		return out, err
	}
	fake.calls = nil
	_, err := manager.ProbeProcess(context.Background(), env, node)
	if err == nil || errors.Is(err, ErrProcessStopped) || !strings.Contains(err.Error(), "restart policy") {
		t.Fatalf("missing restart policy was accepted as stopped evidence: %v", err)
	}
	if hasOperation(fake, "exec") {
		t.Fatal("probe executed Redis commands without a safe restart policy")
	}
}
