package experiment

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/config"
	"github.com/Alaneel/redis-reliability-lab/internal/environment"
	"github.com/Alaneel/redis-reliability-lab/internal/report"
)

type helperNetworkManager struct {
	restoreErr        error
	restoreCalled     bool
	restoreContext    context.Context
	restoredReceipt   *environment.IsolationReceipt
	restoredInventory *environment.Environment
	process           environment.ProcessEvidence
	processErr        error
}

func (m *helperNetworkManager) Isolate(context.Context, *environment.Environment, environment.Node) (*environment.IsolationReceipt, error) {
	return nil, errors.New("unexpected isolation in helper test")
}
func (m *helperNetworkManager) Restore(ctx context.Context, env *environment.Environment, receipt *environment.IsolationReceipt) error {
	m.restoreCalled, m.restoreContext, m.restoredReceipt, m.restoredInventory = true, ctx, receipt, env
	return m.restoreErr
}
func (m *helperNetworkManager) ProbeProcess(context.Context, *environment.Environment, environment.Node) (environment.ProcessEvidence, error) {
	return m.process, m.processErr
}

func TestRestoreNetworkRecordsCompletedActionSeparatelyFromErrors(t *testing.T) {
	for _, failed := range []bool{false, true} {
		t.Run(map[bool]string{false: "restored", true: "restore error"}[failed], func(t *testing.T) {
			manager := &helperNetworkManager{}
			if failed {
				manager.restoreErr = errors.New("network reconnect rejected")
			}
			receipt := &environment.IsolationReceipt{RunID: "test-run"}
			r := &report.Report{Environment: testManagedEnvironment(), NetworkFault: &report.NetworkFault{Receipt: receipt}}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			err := restoreNetwork(ctx, manager, r)
			if !manager.restoreCalled || manager.restoreContext != ctx || manager.restoredReceipt != receipt || manager.restoredInventory != r.Environment || r.NetworkFault.RestoreRequestedAt == nil {
				t.Fatal("restoration context, inventory, receipt or requested timestamp lost")
			}
			if failed {
				if !errors.Is(err, manager.restoreErr) || r.NetworkFault.ConnectivityRestored || r.NetworkFault.RestoreCompletedAt != nil || r.NetworkFault.RestoreDurationMS != nil || len(r.NetworkFault.Errors) != 1 {
					t.Fatalf("restore failure became successful evidence: %+v err=%v", r.NetworkFault, err)
				}
			} else if err != nil || !r.NetworkFault.ConnectivityRestored || r.NetworkFault.RestoreCompletedAt == nil || r.NetworkFault.RestoreDurationMS == nil || *r.NetworkFault.RestoreDurationMS < 0 || len(r.NetworkFault.Errors) != 0 {
				t.Fatalf("completed restoration missing: %+v err=%v", r.NetworkFault, err)
			}
		})
	}
}

func TestRunIsolationRejectsNoncapableManagerBeforeProvision(t *testing.T) {
	c := config.Default()
	c.Experiment.Fault = "primary-isolate"
	m := &fakeManager{env: testManagedEnvironment()}
	provisioned := false
	m.afterProvision = func() { provisioned = true }
	r, err := quietRunner(m).Run(context.Background(), c, t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "network fault manager") || r != nil || provisioned || m.kills != 0 || m.cleaned != nil {
		t.Fatalf("unsupported isolation manager reached resource operations: report=%+v error=%v", r, err)
	}
}

func isolatedProcessFixture() *report.NetworkFault {
	before := environment.ProcessEvidence{ContainerRunning: true, ContainerPID: 1234, ContainerStartedAt: "2026-10-11T01:00:00Z", RedisPONG: true, RedisPID: 1, RedisRunID: strings.Repeat("a", 40), NetworkConnected: true, NetworkCount: 1}
	after := before
	after.NetworkConnected, after.NetworkCount = false, 0
	return &report.NetworkFault{Receipt: &environment.IsolationReceipt{BeforeProcess: before}, ProcessSamples: []report.ProcessSample{{Evidence: after}}}
}

func TestProcessSurvivalRequiresDisconnectedRedisAndStableIdentity(t *testing.T) {
	if !processSurvivedIsolation(isolatedProcessFixture()) {
		t.Fatal("valid isolated process evidence rejected")
	}
	if processSurvivedIsolation(nil) || processSurvivedIsolation(&report.NetworkFault{}) {
		t.Fatal("missing receipt treated as process survival")
	}
	cases := map[string]func(*report.NetworkFault){
		"no observations":           func(n *report.NetworkFault) { n.ProcessSamples = nil },
		"probe failed":              func(n *report.NetworkFault) { n.ProcessSamples[0].Error = "docker exec failed" },
		"process stopped":           func(n *report.NetworkFault) { n.ProcessSamples[0].Evidence.ContainerRunning = false },
		"no Redis PONG":             func(n *report.NetworkFault) { n.ProcessSamples[0].Evidence.RedisPONG = false },
		"network connected":         func(n *report.NetworkFault) { n.ProcessSamples[0].Evidence.NetworkConnected = true },
		"unexpected other network":  func(n *report.NetworkFault) { n.ProcessSamples[0].Evidence.NetworkCount = 1 },
		"different container PID":   func(n *report.NetworkFault) { n.ProcessSamples[0].Evidence.ContainerPID++ },
		"different container start": func(n *report.NetworkFault) { n.ProcessSamples[0].Evidence.ContainerStartedAt = "2026-10-11T01:00:01Z" },
		"different Redis PID":       func(n *report.NetworkFault) { n.ProcessSamples[0].Evidence.RedisPID++ },
		"different Redis run ID":    func(n *report.NetworkFault) { n.ProcessSamples[0].Evidence.RedisRunID = strings.Repeat("b", 40) },
		"missing Redis identity":    func(n *report.NetworkFault) { n.ProcessSamples[0].Evidence.RedisRunID = "" },
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			n := isolatedProcessFixture()
			mutate(n)
			if processSurvivedIsolation(n) {
				t.Fatal("incomplete or changed process evidence accepted")
			}
		})
	}
	// A later successful isolated observation is useful even after a failed probe.
	n := isolatedProcessFixture()
	n.ProcessSamples = append([]report.ProcessSample{{Error: "transient exec timeout"}}, n.ProcessSamples...)
	if !processSurvivedIsolation(n) {
		t.Fatal("a failed earlier probe hid valid disconnected process evidence")
	}
}

func TestCaptureProcessPreservesErrorsAndBoundedLatestEvidence(t *testing.T) {
	m := &helperNetworkManager{process: isolatedProcessFixture().ProcessSamples[0].Evidence, processErr: errors.New("INFO failed")}
	r := &report.Report{Environment: testManagedEnvironment(), NetworkFault: &report.NetworkFault{}}
	captureProcess(context.Background(), m, r, r.Environment.Nodes[0])
	if len(r.NetworkFault.ProcessSamples) != 1 || r.NetworkFault.ProcessSamples[0].Error != "INFO failed" || r.NetworkFault.ProcessSamples[0].Evidence.RedisRunID == "" {
		t.Fatal("failed process probe lost partial evidence")
	}
	m.processErr = nil
	for i := 0; i < 70; i++ {
		m.process.ContainerPID = i + 100
		captureProcess(context.Background(), m, r, r.Environment.Nodes[0])
	}
	if len(r.NetworkFault.ProcessSamples) != 64 || r.NetworkFault.ProcessSamples[0].Error != "INFO failed" || r.NetworkFault.ProcessSamples[63].Evidence.ContainerPID != 169 {
		t.Fatal("process evidence was unbounded or first/latest samples were lost")
	}
}

func TestIsolationReceiptBindsSelectedNodeAndOriginalEndpoints(t *testing.T) {
	cases := map[string]func(*environment.IsolationReceipt){
		"another owned node":         func(x *environment.IsolationReceipt) { x.Node = testManagedEnvironment().Nodes[5] },
		"another container":          func(x *environment.IsolationReceipt) { x.Node.ContainerID = "another-container" },
		"another run":                func(x *environment.IsolationReceipt) { x.RunID = "another-run" },
		"another network":            func(x *environment.IsolationReceipt) { x.NetworkID = "another-network" },
		"another private endpoint":   func(x *environment.IsolationReceipt) { x.OriginalIP = "172.22.0.99" },
		"another published endpoint": func(x *environment.IsolationReceipt) { x.HostAddr = "127.0.0.1:9999" },
	}
	env := testManagedEnvironment()
	target := env.Nodes[3]
	valid := &environment.IsolationReceipt{RunID: env.RunID, NetworkID: env.NetworkID, Node: target, OriginalIP: "172.22.0.5", HostAddr: target.Addr}
	if err := validateIsolationReceipt(valid, env, target); err != nil {
		t.Fatalf("valid selected-node receipt rejected: %v", err)
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			x := *valid
			mutate(&x)
			if err := validateIsolationReceipt(&x, env, target); err == nil {
				t.Fatal("another node/resource/endpoint was accepted as the selected fault action")
			}
		})
	}
	if validateIsolationReceipt(nil, env, target) == nil {
		t.Fatal("absent isolation receipt accepted")
	}
}
