package experiment

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/cluster"
	"github.com/Alaneel/redis-reliability-lab/internal/config"
	"github.com/Alaneel/redis-reliability-lab/internal/environment"
	"github.com/Alaneel/redis-reliability-lab/internal/report"
	"github.com/Alaneel/redis-reliability-lab/internal/workload"
	"github.com/redis/go-redis/v9"
)

func TestOutcomeDistinguishesFullPartialAndFailedRecovery(t *testing.T) {
	m := &report.Milestone{At: time.Now()}
	policyUnavailable := false
	tests := []struct {
		name     string
		recovery report.Recovery
		want     string
	}{
		{"full observable recovery", report.Recovery{TopologyRecovered: true, ClientRecovered: true, DataVerified: true, CombinedRecovered: true, CombinedStable: m}, "recovered"},
		{"basic service recovered but WAIT policy unavailable", report.Recovery{TopologyRecovered: true, ClientRecovered: true, DataVerified: true, CombinedRecovered: true, CombinedStable: m, WritePolicyRecovered: &policyUnavailable}, "partial"},
		{"regression before stable interval reestablished", report.Recovery{TopologyRecovered: true, ClientRecovered: true, DataVerified: true, CombinedStable: m}, "partial"},
		{"asynchronous loss despite service recovery", report.Recovery{TopologyRecovered: true, ClientRecovered: true, CombinedStable: m}, "partial"},
		{"topology without clients", report.Recovery{TopologyRecovered: true, DataVerified: true}, "partial"},
		{"clients without converged topology", report.Recovery{ClientRecovered: true, DataVerified: true}, "partial"},
		{"healthy sample without stable interval", report.Recovery{TopologyRecovered: true, ClientRecovered: true, DataVerified: true}, "partial"},
		{"deadline with partial observations", report.Recovery{TopologyRecovered: true, ClientRecovered: true, DataVerified: true, CombinedStable: m, TimedOut: true}, "partial"},
		{"regression after stable interval", report.Recovery{DataVerified: true, CombinedStable: m}, "partial"},
		{"unrecovered deadline", report.Recovery{TimedOut: true}, "failed"},
		{"unobserved recovery", report.Recovery{}, "failed"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Outcome(tt.recovery); got != tt.want {
				t.Fatalf("got %q; want %q", got, tt.want)
			}
		})
	}
}

func TestNewClientRejectsUnmanagedAndNonLoopbackTargets(t *testing.T) {
	tests := map[string]func(*environment.Environment){
		"remote published endpoint":   func(e *environment.Environment) { e.Nodes[0].Addr = "203.0.113.1:6379" },
		"private published endpoint":  func(e *environment.Environment) { e.Nodes[0].Addr = "172.22.0.2:6379" },
		"hostname published endpoint": func(e *environment.Environment) { e.Nodes[0].Addr = "localhost:6379" },
		"malformed host endpoint":     func(e *environment.Environment) { e.Nodes[0].Addr = "127.0.0.1" },
		"empty port":                  func(e *environment.Environment) { e.Nodes[0].Addr = "127.0.0.1:" },
		"malformed internal endpoint": func(e *environment.Environment) { e.Nodes[0].InternalAddr = "172.22.0.2" },
		"duplicate internal endpoint": func(e *environment.Environment) { e.Nodes[1].InternalAddr = e.Nodes[0].InternalAddr },
		"incomplete inventory":        func(e *environment.Environment) { e.Nodes = e.Nodes[:5] },
		"nonnumeric host port":        func(e *environment.Environment) { e.Nodes[0].Addr = "127.0.0.1:redis" },
		"zero host port":              func(e *environment.Environment) { e.Nodes[0].Addr = "127.0.0.1:0" },
		"out of range host port":      func(e *environment.Environment) { e.Nodes[0].Addr = "127.0.0.1:65536" },
		"duplicate host endpoint":     func(e *environment.Environment) { e.Nodes[1].Addr = e.Nodes[0].Addr },
		"duplicate managed identity":  func(e *environment.Environment) { e.Nodes[1].ID = e.Nodes[0].ID },
		"empty managed identity":      func(e *environment.Environment) { e.Nodes[0].ID = "" },
	}
	for name, mutate := range tests {
		t.Run(name, func(t *testing.T) {
			e := testManagedEnvironment()
			mutate(e)
			client, err := NewClient(e, time.Second)
			if client != nil {
				defer client.Close()
			}
			if err == nil {
				t.Fatal("unsafe inventory accepted")
			}
		})
	}
	if client, err := NewClient(nil, time.Second); err == nil || client != nil {
		t.Fatal("nil managed environment accepted")
	}
	client, err := NewClient(testManagedEnvironment(), time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			t.Error(err)
		}
	}()
	for _, target := range []string{"203.0.113.1:6379", "127.0.0.1:9999", "172.22.0.99:6379", "example.com:6379"} {
		conn, err := client.Options().Dialer(context.Background(), "tcp", target)
		if conn != nil {
			conn.Close()
			t.Fatalf("unmanaged target %s returned connection", target)
		}
		if err == nil || !strings.Contains(err.Error(), "unmanaged") {
			t.Fatalf("unmanaged target %s was not explicitly rejected: %v", target, err)
		}
	}
}

func TestProbeKeysCoversEveryInitialShard(t *testing.T) {
	initial := cluster.Snapshot{Views: []cluster.View{{Coverage: cluster.Coverage{Owners: []cluster.OwnerRange{
		{Start: 10923, End: 16383, RedisID: "primary-c"},
		{Start: 0, End: 100, RedisID: "primary-a"},
		{Start: 101, End: 10922, RedisID: "primary-b"},
	}}}}}
	keys, err := ProbeKeys(initial)
	if err != nil {
		t.Fatal(err)
	}
	if len(keys) != 3 {
		t.Fatalf("got %d probes", len(keys))
	}
	seen := map[string]bool{}
	for _, key := range keys {
		slot := int(workload.HashSlot(key))
		owner := ""
		for _, r := range initial.Views[0].Coverage.Owners {
			if slot >= r.Start && slot <= r.End {
				owner = r.RedisID
			}
		}
		if owner == "" || seen[owner] {
			t.Fatalf("probe %q slot %d misses or repeats a shard", key, slot)
		}
		seen[owner] = true
	}
	other, err := ProbeKeys(initial)
	if err != nil || strings.Join(keys, ",") != strings.Join(other, ",") {
		t.Fatal("probe key selection is not deterministic")
	}
	if _, err := ProbeKeys(cluster.Snapshot{}); err == nil {
		t.Fatal("missing topology accepted")
	}
	if _, err := ProbeKeys(cluster.Snapshot{Views: []cluster.View{{Coverage: cluster.Coverage{Owners: []cluster.OwnerRange{{Start: 0, End: 16383, RedisID: "only-primary"}}}}}}); err == nil {
		t.Fatal("single-shard probes accepted")
	}
}

func TestProbeAllRequiresSuccessfulReadWriteOnEveryShard(t *testing.T) {
	keys := []string{"a", "b", "c"}
	tests := []struct {
		name       string
		configure  func(*probeClient)
		successful bool
		errors     int
	}{
		{"all shards read and write", func(*probeClient) {}, true, 0},
		{"write unavailable", func(c *probeClient) { c.setErrors["a"] = errors.New("CLUSTERDOWN") }, false, 1},
		{"read unavailable", func(c *probeClient) { c.getErrors["b"] = errors.New("connection refused") }, false, 1},
		{"missing acknowledged probe", func(c *probeClient) { c.getErrors["c"] = redis.Nil }, false, 1},
		{"corrupt probe value", func(c *probeClient) { c.getValues["a"] = "wrong-value" }, false, 1},
		{"multiple failed shards", func(c *probeClient) { c.setErrors["a"] = errors.New("unreachable"); c.getErrors["c"] = redis.Nil }, false, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newProbeClient()
			tt.configure(c)
			p := ProbeAll(context.Background(), c, keys, "expected", 100*time.Millisecond)
			if p.Successful != tt.successful || len(p.Errors) != tt.errors || p.At.IsZero() {
				t.Fatalf("wrong probe result: %+v", p)
			}
			if c.setCalls != 3 {
				t.Fatal("probe did not attempt every initial shard")
			}
		})
	}
	c := newProbeClient()
	if p := ProbeAll(context.Background(), c, keys[:2], "expected", time.Second); p.Successful {
		t.Fatal("two-shard success mistaken for all-shard recovery")
	}
}

func TestProbeAllPropagatesBoundedOperationContexts(t *testing.T) {
	p := ProbeAll(context.Background(), deadlineProbeClient{}, []string{"a", "b", "c"}, "expected", 10*time.Millisecond)
	if p.Successful || len(p.Errors) != 3 {
		t.Fatalf("timed-out probes accepted: %+v", p)
	}
	for _, err := range p.Errors {
		if !strings.Contains(err, "context deadline exceeded") {
			t.Fatalf("deadline omitted: %s", err)
		}
	}
}

func TestProbeAllSeparatesWriteFailureFromUnattemptedRead(t *testing.T) {
	keys := []string{"a", "b", "c"}
	t.Run("all writes acknowledged but one read fails", func(t *testing.T) {
		client := newProbeClient()
		client.getErrors["b"] = errors.New("connection reset")
		p := ProbeAll(context.Background(), client, keys, "expected", time.Second)
		if p.Successful || p.WriteSuccessful == nil || !*p.WriteSuccessful || p.ReadSuccessful == nil || *p.ReadSuccessful {
			t.Fatalf("failed read erased successful write observation: %+v", p)
		}
	})
	t.Run("failed SET means not all reads attempted", func(t *testing.T) {
		client := newProbeClient()
		client.setErrors["b"] = errors.New("CLUSTERDOWN")
		p := ProbeAll(context.Background(), client, keys, "expected", time.Second)
		if p.Successful || p.WriteSuccessful == nil || *p.WriteSuccessful || p.ReadSuccessful != nil {
			t.Fatalf("failed SET invented failed GET evidence: %+v", p)
		}
	})
	t.Run("incomplete shard population remains unobserved", func(t *testing.T) {
		p := ProbeAll(context.Background(), newProbeClient(), keys[:2], "expected", time.Second)
		if p.Successful || p.ReadSuccessful != nil || p.WriteSuccessful != nil {
			t.Fatal("fewer than three shards established all-shard availability")
		}
	})
}

func TestRunProvisionFailureCleansPartialInventoryAndWritesEvidence(t *testing.T) {
	partial := testManagedEnvironment()
	partial.Nodes = partial.Nodes[:2]
	m := &fakeManager{env: partial, provisionErr: errors.New("node 3 creation failed")}
	output := t.TempDir()
	r, err := quietRunner(m).Run(context.Background(), config.Default(), output)
	if err == nil || !strings.Contains(err.Error(), "node 3 creation failed") {
		t.Fatalf("provision error lost: %v", err)
	}
	if r == nil || r.Outcome != "failed" || !r.Cleanup.Attempted || !r.Cleanup.Successful {
		t.Fatalf("wrong failure report: %+v", r)
	}
	assertFreshCleanup(t, m)
	if m.cleaned != partial || m.kills != 0 {
		t.Fatal("partial inventory not cleaned, or fault ran after provision failure")
	}
	assertReportFiles(t, output, r.RunID, "failed")
	stages := reportStages(r)
	if stages != "created,provisioning,cleanup,complete" {
		t.Fatalf("unexpected lifecycle: %s", stages)
	}
}

func TestRunCleanupFailureIsVisibleInReportAndReturnedError(t *testing.T) {
	m := &fakeManager{env: testManagedEnvironment(), provisionErr: errors.New("provision failed"), cleanupErr: errors.New("container removal failed")}
	output := t.TempDir()
	r, err := quietRunner(m).Run(context.Background(), config.Default(), output)
	if err == nil || !strings.Contains(err.Error(), "provision failed") || !strings.Contains(err.Error(), "container removal failed") {
		t.Fatalf("combined errors missing: %v", err)
	}
	if r.Cleanup.Successful || r.Cleanup.Error == "" || r.Outcome != "failed" {
		t.Fatal("cleanup failure hidden")
	}
	assertFreshCleanup(t, m)
	assertReportFiles(t, output, r.RunID, "failed")
}

func TestRunCanceledReadinessStillCleansAndWritesAbortedReport(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	m := &fakeManager{env: testManagedEnvironment(), afterProvision: cancel}
	output := t.TempDir()
	r, err := quietRunner(m).Run(ctx, config.Default(), output)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("cancellation cause lost: %v", err)
	}
	if r == nil || r.Outcome != "aborted" || !r.Cleanup.Successful || r.Initial == nil {
		t.Fatalf("canceled readiness evidence missing: %+v", r)
	}
	assertFreshCleanup(t, m)
	assertReportFiles(t, output, r.RunID, "aborted")
	if stages := reportStages(r); stages != "created,provisioning,readiness,cleanup,complete" {
		t.Fatalf("unexpected canceled lifecycle: %s", stages)
	}
	if m.kills != 0 {
		t.Fatal("fault injection ran after cancellation")
	}
}

func TestRunCancellationDuringCleanupPreservesCauseAndFreshDeadline(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	m := &fakeManager{env: testManagedEnvironment(), provisionErr: errors.New("provision failed"), afterCleanup: cancel}
	output := t.TempDir()
	r, err := quietRunner(m).Run(ctx, config.Default(), output)
	if !errors.Is(err, context.Canceled) || !strings.Contains(err.Error(), "provision failed") {
		t.Fatalf("late cancellation or earlier failure cause lost: %v", err)
	}
	if r.Outcome != "aborted" || !r.Cleanup.Successful {
		t.Fatalf("late cancellation prevented cleanup or was hidden: %+v", r.Cleanup)
	}
	assertFreshCleanup(t, m)
	assertReportFiles(t, output, r.RunID, "aborted")
}

func TestRunRejectsNilSuccessfulProvisionerInventory(t *testing.T) {
	m := &fakeManager{}
	output := t.TempDir()
	r, err := quietRunner(m).Run(context.Background(), config.Default(), output)
	if err == nil || r == nil || r.Outcome != "failed" || !strings.Contains(err.Error(), "six managed nodes") {
		t.Fatalf("invalid successful provisioning not rejected: report=%+v err=%v", r, err)
	}
	if r.Cleanup.Attempted || m.cleaned != nil {
		t.Fatal("nil inventory caused a cleanup attempt")
	}
	if _, err := os.Stat(filepath.Join(output, r.RunID, "report.json")); err != nil {
		t.Fatal("invalid inventory error report missing:", err)
	}
}

func TestMilestonesDistinguishPromotionFromTopologyConvergence(t *testing.T) {
	fault := time.Unix(1728763200, 0).UTC()
	at := fault.Add(2 * time.Second)
	selection := cluster.Selection{Target: cluster.Node{ID: "node-1"}, TargetID: "old-primary", ReplicaID: "replacement", Slots: []int{10, 11, 12}}
	view := cluster.View{Node: cluster.Node{ID: "node-4"}, Reachable: true, Nodes: []cluster.NodeState{
		{ID: "old-primary", Flags: []string{"master", "fail"}, LinkState: "disconnected"},
		{ID: "replacement", Flags: []string{"master", "myself"}, LinkState: "connected", Slots: []cluster.SlotRange{{Start: 10, End: 12}}},
	}, Coverage: cluster.Coverage{Assigned: 3, Owners: []cluster.OwnerRange{{Start: 10, End: 12, RedisID: "replacement"}}}}
	s := report.Sample{Snapshot: cluster.Snapshot{At: at, Completed: at, Views: []cluster.View{view}}, Probe: report.Probe{At: at.Add(10 * time.Millisecond), Successful: true}, TopologyOK: false}
	r := &report.Report{}
	updateMilestones(r, s, selection, fault)
	if r.Recovery.FailureDetected == nil || r.Recovery.ReplicaPromoted == nil || r.Recovery.ClientFirstSuccess == nil {
		t.Fatal("independent milestone evidence discarded")
	}
	if r.Recovery.TopologyConverged != nil || r.Recovery.FullCoverage != nil {
		t.Fatal("partial single-observer topology treated as converged full coverage")
	}
	if r.Recovery.ReplicaPromoted.SinceFaultMS != 2000 {
		t.Fatalf("wrong sampled promotion duration: %+v", r.Recovery.ReplicaPromoted)
	}
	s.Snapshot.Completed = at.Add(time.Second)
	s.TopologyOK = true
	updateMilestones(r, s, selection, fault)
	if r.Recovery.ReplicaPromoted.SinceFaultMS != 2000 || r.Recovery.TopologyConverged == nil || r.Recovery.TopologyConverged.SinceFaultMS != 3000 {
		t.Fatal("first promotion timestamp overwritten or convergence conflated")
	}
}

func TestFullCoverageMilestoneRejectsStaleOwnershipOfKilledPrimary(t *testing.T) {
	fault := time.Unix(1728763200, 0).UTC()
	at := fault.Add(time.Second)
	ids := []string{"old-primary", "primary-b", "primary-c", "replacement", "replica-b", "replica-c"}
	selection := cluster.Selection{Target: cluster.Node{ID: "node-1"}, TargetID: ids[0], ReplicaID: ids[3], Slots: []int{0, 1, 2}}
	snapshot := cluster.Snapshot{At: at, Completed: at}
	for self := 0; self < 6; self++ {
		view := cluster.View{Node: cluster.Node{ID: fmt.Sprintf("node-%d", self+1)}, RedisID: ids[self], Reachable: self != 0, ClusterInfo: map[string]string{"cluster_state": "ok"}}
		if self == 0 {
			view.Errors = []string{"connection refused"}
			snapshot.Views = append(snapshot.Views, view)
			continue
		}
		for i, id := range ids {
			n := cluster.NodeState{ID: id, Flags: []string{"master"}, LinkState: "connected"}
			if i < 3 {
				n.Slots = []cluster.SlotRange{{Start: []int{0, 5461, 10923}[i], End: []int{5460, 10922, 16383}[i]}}
			} else {
				n.Flags = []string{"slave"}
				n.PrimaryID = ids[i-3]
			}
			view.Nodes = append(view.Nodes, n)
		}
		view.Coverage = cluster.AnalyzeCoverage(view.Nodes)
		if !view.Coverage.Full() {
			t.Fatal("fixture must represent a stale complete healthy owner map")
		}
		snapshot.Views = append(snapshot.Views, view)
	}
	r := &report.Report{}
	sample := report.Sample{Snapshot: snapshot, Probe: report.Probe{At: at}}
	updateMilestones(r, sample, selection, fault)
	if r.Recovery.FullCoverage != nil {
		t.Fatalf("stale complete maps with unreachable owner reported usable coverage: %+v", r.Recovery.FullCoverage)
	}
	// After ownership changes to the reachable promoted replica, full coverage
	// has independent evidence even before cross-view convergence is confirmed.
	for i := 1; i < 6; i++ {
		v := &snapshot.Views[i]
		v.Nodes[0].Flags = []string{"master", "fail"}
		v.Nodes[0].LinkState = "disconnected"
		v.Nodes[0].Slots = nil
		v.Nodes[3].Flags = []string{"master"}
		v.Nodes[3].PrimaryID = ""
		v.Nodes[3].Slots = []cluster.SlotRange{{Start: 0, End: 5460}}
		v.Coverage = cluster.AnalyzeCoverage(v.Nodes)
	}
	snapshot.Completed = at.Add(time.Second)
	sample.Snapshot = snapshot
	updateMilestones(r, sample, selection, fault)
	if r.Recovery.FullCoverage == nil || r.Recovery.FullCoverage.SinceFaultMS != 2000 {
		t.Fatal("reachable replacement coverage milestone missing or wrongly timed")
	}
}

type fakeManager struct {
	env                      *environment.Environment
	provisionErr, cleanupErr error
	afterProvision           func()
	afterCleanup             func()
	cleaned                  *environment.Environment
	cleanupContextErr        error
	cleanupDeadline          bool
	kills                    int
}

func (m *fakeManager) Provision(_ context.Context, id string, _ environment.Config) (*environment.Environment, error) {
	if m.env != nil {
		m.env.RunID = id
	}
	if m.afterProvision != nil {
		m.afterProvision()
	}
	return m.env, m.provisionErr
}
func (m *fakeManager) Kill(context.Context, *environment.Environment, environment.Node) error {
	m.kills++
	return nil
}
func (m *fakeManager) Cleanup(ctx context.Context, e *environment.Environment) error {
	m.cleaned = e
	m.cleanupContextErr = ctx.Err()
	_, m.cleanupDeadline = ctx.Deadline()
	if m.afterCleanup != nil {
		m.afterCleanup()
	}
	return m.cleanupErr
}

func assertFreshCleanup(t *testing.T, m *fakeManager) {
	t.Helper()
	if m.cleaned == nil || m.cleanupContextErr != nil || !m.cleanupDeadline {
		t.Fatalf("cleanup did not get fresh bounded context: inventory=%v err=%v bounded=%t", m.cleaned, m.cleanupContextErr, m.cleanupDeadline)
	}
}
func quietRunner(m EnvironmentManager) Runner {
	return Runner{Manager: m, Logger: slog.New(slog.NewTextHandler(io.Discard, nil))}
}
func reportStages(r *report.Report) string {
	var stages []string
	for _, e := range r.Events {
		stages = append(stages, e.Stage)
	}
	return strings.Join(stages, ",")
}
func assertReportFiles(t *testing.T, output, id, outcome string) {
	t.Helper()
	dir := filepath.Join(output, id)
	raw, err := os.ReadFile(filepath.Join(dir, "report.json"))
	if err != nil {
		t.Fatal(err)
	}
	var fields struct {
		Outcome string         `json:"outcome"`
		Errors  []string       `json:"errors"`
		Cleanup report.Cleanup `json:"cleanup"`
	}
	if err := json.Unmarshal(raw, &fields); err != nil {
		t.Fatal(err)
	}
	if fields.Outcome != outcome || !fields.Cleanup.Attempted || len(fields.Errors) == 0 {
		t.Fatalf("failure evidence absent from JSON: %+v", fields)
	}
	md, err := os.ReadFile(filepath.Join(dir, "report.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(md), "Outcome: **"+outcome+"**") {
		t.Fatal("Markdown outcome differs from JSON")
	}
}
func testManagedEnvironment() *environment.Environment {
	e := &environment.Environment{}
	for i := 0; i < 6; i++ {
		e.Nodes = append(e.Nodes, environment.Node{ID: fmt.Sprintf("node-%d", i+1), Addr: fmt.Sprintf("127.0.0.1:%d", 7000+i), InternalAddr: fmt.Sprintf("172.22.0.%d:6379", i+2)})
	}
	return e
}

type probeClient struct {
	mu                   sync.Mutex
	stored, getValues    map[string]string
	setErrors, getErrors map[string]error
	setCalls             int
}

func newProbeClient() *probeClient {
	return &probeClient{stored: map[string]string{}, getValues: map[string]string{}, setErrors: map[string]error{}, getErrors: map[string]error{}}
}
func (c *probeClient) Set(ctx context.Context, key string, value interface{}, _ time.Duration) *redis.StatusCmd {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.setCalls++
	cmd := redis.NewStatusCmd(ctx)
	if err := c.setErrors[key]; err != nil {
		cmd.SetErr(err)
	} else {
		c.stored[key] = fmt.Sprint(value)
		cmd.SetVal("OK")
	}
	return cmd
}
func (c *probeClient) Get(ctx context.Context, key string) *redis.StringCmd {
	c.mu.Lock()
	defer c.mu.Unlock()
	cmd := redis.NewStringCmd(ctx)
	if err := c.getErrors[key]; err != nil {
		cmd.SetErr(err)
	} else if value, ok := c.getValues[key]; ok {
		cmd.SetVal(value)
	} else {
		cmd.SetVal(c.stored[key])
	}
	return cmd
}

type deadlineProbeClient struct{}

func (deadlineProbeClient) Set(ctx context.Context, _ string, _ interface{}, _ time.Duration) *redis.StatusCmd {
	<-ctx.Done()
	cmd := redis.NewStatusCmd(ctx)
	cmd.SetErr(ctx.Err())
	return cmd
}
func (deadlineProbeClient) Get(ctx context.Context, _ string) *redis.StringCmd {
	cmd := redis.NewStringCmd(ctx)
	cmd.SetErr(errors.New("Get must not run after failed SET"))
	return cmd
}
