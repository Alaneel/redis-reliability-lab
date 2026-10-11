package experiment

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/metrics"
	"github.com/alanwang/redis-reliability-lab/internal/report"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
	"github.com/redis/go-redis/v9"
)

type persistenceLog struct {
	mu                    sync.Mutex
	calls                 []string
	validated             bool
	values                map[string]string
	writeBeforeValidation bool
}

func (l *persistenceLog) add(s string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.calls = append(l.calls, s)
}

type persistenceFakeTraffic struct {
	log               *persistenceLog
	stopped           bool
	missing           int
	validationCalls   int
	invalidValidation bool
}

func (f *persistenceFakeTraffic) Stop() { f.log.add("stop"); f.stopped = true }
func (f *persistenceFakeTraffic) Stats() workload.Stats {
	f.log.add("stats")
	if !f.stopped {
		panic("stats frozen before drain")
	}
	return workload.Stats{AcknowledgedWrites: 2, AmbiguousWrites: 1, WritesAttempted: 3, StartedRequests: 4, CompletedRequests: 4, OfferedRequests: 5, DroppedRequests: 1}
}
func (f *persistenceFakeTraffic) Snapshot() metrics.Snapshot {
	f.log.add("metrics")
	return metrics.Snapshot{Totals: metrics.Summary{Attempts: 4}, Phases: map[string]metrics.PhaseSummary{"baseline": {Attempts: 4, Successes: 4, SuccessRate: 1}}}
}
func (f *persistenceFakeTraffic) Ledger() []workload.Write {
	f.log.add("ledger")
	return []workload.Write{{Key: "one", Value: "first"}, {Key: "two", Value: "second"}}
}
func (f *persistenceFakeTraffic) Validate(context.Context, time.Duration) workload.Validation {
	if !f.stopped {
		panic("validation before drain")
	}
	f.validationCalls++
	if f.invalidValidation {
		return workload.Validation{}
	}
	f.log.mu.Lock()
	f.log.calls = append(f.log.calls, "validate")
	f.log.validated = true
	f.log.mu.Unlock()
	at := time.Now()
	return workload.Validation{StartedAt: at, CompletedAt: time.Now(), Complete: true, Expected: 2, Checked: 2, Missing: f.missing}
}

type persistenceFakeClient struct{ log *persistenceLog }

func (f persistenceFakeClient) Set(_ context.Context, key string, value interface{}, _ time.Duration) *redis.StatusCmd {
	f.log.mu.Lock()
	defer f.log.mu.Unlock()
	f.log.calls = append(f.log.calls, "diagnostic-set")
	f.log.writeBeforeValidation = f.log.writeBeforeValidation || !f.log.validated
	f.log.values[key] = fmt.Sprint(value)
	return redis.NewStatusResult("OK", nil)
}
func (f persistenceFakeClient) Get(_ context.Context, key string) *redis.StringCmd {
	f.log.mu.Lock()
	defer f.log.mu.Unlock()
	return redis.NewStringResult(f.log.values[key], nil)
}

type persistenceFakeManager struct {
	log                                       *persistenceLog
	mode                                      string
	killed                                    map[string]bool
	killErrorAt, restartErrorAt, probeErrorAt int
	kills, restarts, stoppedProbes            int
	wrongRestartProcess                       bool
	busyPreparation                           int
	probes                                    int
	cancel                                    context.CancelFunc
	started                                   map[string]bool
	stoppedError                              error
	stoppedPolicy                             string
	oldRestartRequest                         bool
	earlyReceiptTransient                     bool
}

func (m *persistenceFakeManager) Provision(context.Context, string, environment.Config) (*environment.Environment, error) {
	panic("not used")
}
func (m *persistenceFakeManager) Cleanup(context.Context, *environment.Environment) error {
	m.log.add("cleanup")
	return nil
}
func (m *persistenceFakeManager) Kill(_ context.Context, _ *environment.Environment, n environment.Node) error {
	m.log.add("kill:" + n.ID)
	m.kills++
	if m.killErrorAt == m.kills {
		return errors.New("injected kill error")
	}
	m.killed[n.ID] = true
	return nil
}
func (m *persistenceFakeManager) ProbeProcess(_ context.Context, _ *environment.Environment, n environment.Node) (environment.ProcessEvidence, error) {
	m.probes++
	if m.killed[n.ID] && !m.started[n.ID] {
		m.log.add("stopped:" + n.ID)
		m.stoppedProbes++
		if m.probeErrorAt == m.stoppedProbes {
			return environment.ProcessEvidence{}, errors.New("ownership inspection error")
		}
		at := time.Now()
		e := environment.ProcessEvidence{StartedAt: at, ObservedAt: time.Now(), ContainerStartedAt: "old-" + n.ID, ContainerRestartPolicy: "no"}
		if m.stoppedPolicy != "" {
			e.ContainerRestartPolicy = m.stoppedPolicy
		}
		if m.stoppedError != nil {
			return e, m.stoppedError
		}
		return e, environment.ErrProcessStopped
	}
	m.log.add("prepare:" + n.ID)
	e := persistenceLiveEvidence(m.mode, n)
	if m.started[n.ID] {
		e.ContainerStartedAt = "new-" + n.ID
		e.RedisRunID = "new-" + n.ID
	}
	if m.busyPreparation > 0 {
		m.busyPreparation--
		e.RawPersistence += "aof_rewrite_in_progress:1\n"
	}
	return e, nil
}
func (m *persistenceFakeManager) Restart(_ context.Context, env *environment.Environment, n environment.Node) (*environment.RestartReceipt, error) {
	m.log.add("restart:" + n.ID)
	m.restarts++
	if len(m.killed) != 6 || m.stoppedProbes != 6 {
		panic("restart before all six stopped proofs")
	}
	at := time.Now()
	a := persistenceLiveEvidence(m.mode, n)
	a.ContainerStartedAt = "new-" + n.ID
	a.RedisRunID = "new-" + n.ID
	if m.wrongRestartProcess {
		a.RedisRunID = "old-" + n.ID
	}
	x := &environment.RestartReceipt{RunID: env.RunID, NodeID: n.ID, Node: n, OriginalRedisID: n.RedisClusterID, RequestedAt: at, Before: environment.ProcessEvidence{ObservedAt: at, ContainerStartedAt: "old-" + n.ID, ContainerRestartPolicy: "no"}, After: a, CompletedAt: time.Now()}
	m.started[n.ID] = true
	if m.oldRestartRequest {
		x.RequestedAt = at.Add(-time.Second)
	}
	if m.earlyReceiptTransient {
		if m.mode == "memory" {
			x.After.RawDataFiles += "/data/dump.rdb\n"
		} else {
			x.After.RawPersistence = strings.Replace(x.After.RawPersistence, "aof_enabled:1", "aof_enabled:0", 1)
		}
	}
	if m.cancel != nil {
		m.cancel()
	}
	if m.restarts == m.restartErrorAt {
		return x, errors.New("injected applied restart error")
	}
	return x, nil
}

func persistenceLiveEvidence(mode string, n environment.Node) environment.ProcessEvidence {
	values := map[string]string{"appendonly": "no", "appendfsync": "everysec", "save": "", "repl-diskless-load": "swapdb", "repl-diskless-sync": "yes", "auto-aof-rewrite-percentage": "0", "maxmemory": "134217728", "maxmemory-policy": "noeviction", "aof-use-rdb-preamble": "no", "no-appendfsync-on-rewrite": "no", "aof-load-truncated": "yes", "rdb-del-sync-files": "yes"}
	files, enabled := "/data/nodes.conf\n", "0"
	if mode != "memory" {
		values["appendonly"] = "yes"
		values["appendfsync"] = strings.TrimPrefix(mode, "aof-")
		values["aof-use-rdb-preamble"] = "yes"
		files += "/data/appendonlydir/appendonly.aof.manifest\n/data/appendonlydir/appendonly.aof.1.incr.aof\n"
		enabled = "1"
	}
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	lines := make([]string, 0, len(keys)*2)
	for _, key := range keys {
		lines = append(lines, key, values[key])
	}
	at := time.Now()
	return environment.ProcessEvidence{StartedAt: at, ObservedAt: time.Now(), ContainerRunning: true, ContainerPID: 1234, ContainerStartedAt: "old-" + n.ID, ContainerRestartPolicy: "no", RedisPONG: true, RedisPID: 1, RedisRunID: "old-" + n.ID, RedisClusterID: n.RedisClusterID, NetworkConnected: true, NetworkCount: 1, RawPersistenceConfig: strings.Join(lines, "\n"), RawPersistence: "loading:0\naof_enabled:" + enabled + "\naof_last_write_status:ok\naof_last_bgrewrite_status:ok\naof_rewrite_in_progress:0\naof_rewrite_scheduled:0\naof_pending_bio_fsync:0\naof_pending_rewrite:0\nrdb_bgsave_in_progress:0\n", RawDataFiles: files, RawReplication: "role:master\nconnected_slaves:1\n"}
}

type persistenceObserverFunc func(context.Context) cluster.Snapshot

func (f persistenceObserverFunc) Observe(ctx context.Context) cluster.Snapshot { return f(ctx) }

func persistenceFixture(t *testing.T, mode string) (config.Config, *report.Report, *persistenceFakeManager, *persistenceFakeTraffic, persistenceObserverFunc) {
	t.Helper()
	initial := lifecycleSnapshot(t, time.Now())
	selected, err := cluster.SelectPrimary(initial)
	if err != nil {
		t.Fatal(err)
	}
	c := config.Default()
	c.Experiment.Fault = "cluster-restart"
	c.Environment.Persistence = mode
	c.Experiment.DownDuration = config.Duration(time.Millisecond)
	c.Experiment.SampleInterval = config.Duration(time.Millisecond)
	c.Experiment.StableFor = config.Duration(time.Millisecond)
	c.Experiment.PostRecovery = config.Duration(time.Millisecond)
	c.Timeouts.Repair = config.Duration(5 * time.Second)
	c.Timeouts.Recovery = config.Duration(5 * time.Second)
	env := &environment.Environment{RunID: "persistence-test", RetainData: true, Persistence: mode}
	for _, v := range initial.Views {
		env.Nodes = append(env.Nodes, environment.Node{ID: v.Node.ID, Addr: v.Node.Addr, InternalAddr: v.Node.InternalAddr, ContainerID: "owned-" + v.Node.ID, RedisClusterID: v.RedisID})
		env.Volumes = append(env.Volumes, environment.Volume{Name: "volume-" + v.Node.ID, NodeID: v.Node.ID})
	}
	r := &report.Report{RunID: env.RunID, Config: c, SchemaVersion: report.LifecycleSchemaVersion, Environment: env, Initial: &initial, Selection: &selected, Outcome: "failed", PersistenceFault: &report.PersistenceFault{Mode: mode, FailureModel: PersistenceFailureModel}}
	l := &persistenceLog{values: map[string]string{}}
	m := &persistenceFakeManager{log: l, mode: mode, killed: map[string]bool{}, started: map[string]bool{}}
	traffic := &persistenceFakeTraffic{log: l}
	observer := persistenceObserverFunc(func(context.Context) cluster.Snapshot {
		l.add("readonly-observe")
		s := initial
		s.At, s.Completed = time.Now(), time.Now()
		return s
	})
	if err := preparePersistence(context.Background(), c, r, m); err != nil {
		t.Fatal(err)
	}
	l.calls = nil
	return c, r, m, traffic, observer
}

func runPersistenceFixture(ctx context.Context, c config.Config, r *report.Report, m *persistenceFakeManager, traffic *persistenceFakeTraffic, observer persistenceObserverFunc) error {
	return runPersistenceExperiment(ctx, c, r, m, m, m, observer, traffic, persistenceFakeClient{m.log}, nil, []string{"probe-a", "probe-b", "probe-c"}, func(stage Stage) { m.log.add("stage:" + string(stage)) })
}

func TestPersistenceFreezeBarrierAndValidationOrder(t *testing.T) {
	for _, mode := range []string{"memory", "aof-everysec", "aof-always"} {
		t.Run(mode, func(t *testing.T) {
			c, r, m, traffic, o := persistenceFixture(t, mode)
			if mode == "memory" {
				traffic.missing = 2
			}
			err := runPersistenceFixture(context.Background(), c, r, m, traffic, o)
			p := r.PersistenceFault
			if p.AcknowledgedFrozen != 2 || p.AmbiguousFrozen != 1 || traffic.validationCalls != 1 || m.log.writeBeforeValidation || len(p.KillActions) != 6 || len(p.Stopped) != 6 || len(p.Restarts) != 6 || len(p.Restored) != 6 || len(p.RestartActions) != 6 || p.StableRedundancy == nil || p.DiagnosticServiceStable == nil || p.RestoreTimedOut {
				t.Fatalf("frozen/barrier evidence incomplete: %+v err=%v", p, err)
			}
			if !reflect.DeepEqual(*p.FrozenValidation, r.Validation) || p.TrafficDrainedAt.After(*p.LedgerFrozenAt) || p.LedgerFrozenAt.After(p.KillActions[0].RequestedAt) || p.AllStoppedAt.After(*p.RestartRequestedAt) || p.RestartCompletedAt.After(r.Validation.StartedAt) || p.DataCheckedAt.After(*p.DiagnosticsRequestedAt) || p.DrainToFirstKillMS == nil || p.DrainToLastKillMS == nil || p.AllStoppedToFirstRestartMS == nil {
				t.Fatal("stage ordering or frozen validation changed")
			}
			kills, restarts, _ := persistenceNodeOrder(r.Environment, r.Initial)
			for i := range kills {
				if p.KillActions[i].Node != kills[i] || p.RestartActions[i].Node != restarts[i] || p.Stopped[i].Error != environment.ErrProcessStopped.Error() {
					t.Fatal("original-role order or authoritative stopped status lost")
				}
			}
			for _, s := range r.Samples {
				if s.Phase == "persistence_recovery" && (!s.Probe.At.IsZero() || s.WritePolicyProbe != nil || s.Dimensions.ReadAvailable != nil || s.Dimensions.WriteAvailable != nil || s.Dimensions.DataVerified != nil) {
					t.Fatal("pre-validation observation fabricated service/data probes")
				}
			}
			if r.Dimensions.MembershipHealthy == nil || !*r.Dimensions.MembershipHealthy || r.Dimensions.PersistenceRecovered == nil || r.Dimensions.ReadAvailable == nil || !*r.Dimensions.ReadAvailable || r.Dimensions.WritePolicyAvailable == nil || !*r.Dimensions.WritePolicyAvailable {
				t.Fatalf("independent final dimensions missing: %+v", r.Dimensions)
			}
			if mode == "memory" {
				if err == nil || r.Outcome != "partial" || r.Recovery.DataVerified || p.PersistenceRecoveryObserved != nil || p.PersistenceRecoveryAt != nil || *r.Dimensions.PersistenceRecovered {
					t.Fatal("memory data loss/persistence absence became recovery")
				}
			} else if err != nil || r.Outcome != "recovered" || !r.Recovery.DataVerified || p.PersistenceRecoveryObserved == nil || !*p.PersistenceRecoveryObserved || p.PersistenceRecoveryAt == nil || !*r.Dimensions.PersistenceRecovered {
				t.Fatalf("measured AOF recovery missing: %v", err)
			}
		})
	}
}

func TestPersistenceFaultBarrierRefusesPartialKillOrUnownedProbe(t *testing.T) {
	for _, name := range []string{"partial kill", "ownership probe", "joined ownership error", "restart policy", "same process", "old action receipt", "applied restart failure", "cancel restart"} {
		t.Run(name, func(t *testing.T) {
			c, r, m, traffic, o := persistenceFixture(t, "aof-always")
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			switch name {
			case "partial kill":
				m.killErrorAt = 2
			case "ownership probe":
				m.probeErrorAt = 3
			case "joined ownership error":
				m.stoppedError = errors.Join(environment.ErrProcessStopped, errors.New("ownership inspection failed"))
			case "restart policy":
				m.stoppedPolicy = "always"
			case "same process":
				m.wrongRestartProcess = true
			case "old action receipt":
				m.oldRestartRequest = true
			case "applied restart failure":
				m.restartErrorAt = 1
			case "cancel restart":
				m.cancel = cancel
			}
			err := runPersistenceFixture(ctx, c, r, m, traffic, o)
			if err == nil || r.Outcome == "recovered" || r.Recovery.ProcessRestarted || r.PersistenceFault.StableRedundancy != nil || traffic.validationCalls != 0 || m.log.writeBeforeValidation {
				t.Fatalf("failure/cancellation became restoration: err=%v report=%+v", err, r.PersistenceFault)
			}
			if (name == "partial kill" || name == "ownership probe" || name == "joined ownership error" || name == "restart policy") && (m.restarts != 0 || r.PersistenceFault.AllStoppedAt != nil) {
				t.Fatal("restart crossed incomplete stopped barrier")
			}
			if name == "applied restart failure" && (len(r.PersistenceFault.Restarts) != 1 || !r.PersistenceFault.Restarts[0].After.ContainerRunning || len(r.PersistenceFault.RestartActions) != 1 || r.PersistenceFault.RestartActions[0].Error == "" || r.PersistenceFault.RestartActions[0].CompletedAt != nil) {
				t.Fatal("ambiguous applied action proof or error lost")
			}
			if name == "cancel restart" && !errors.Is(err, context.Canceled) {
				t.Fatal("parent cancellation lost")
			}
		})
	}
}

func TestPersistenceReadOnlyRestorationRejectsChangedOwnershipAndLateSamples(t *testing.T) {
	for _, name := range []string{"changed ownership", "crossed deadline", "cancellation"} {
		t.Run(name, func(t *testing.T) {
			c, r, _, _, _ := persistenceFixture(t, "aof-always")
			fault := time.Now()
			r.FaultRequestedAt = &fault
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer cancel()
			o := persistenceObserverFunc(func(ctx context.Context) cluster.Snapshot {
				s := *r.Initial
				s.At, s.Completed = time.Now(), time.Now()
				switch name {
				case "changed ownership":
					s.Views = append([]cluster.View(nil), s.Views...)
					for i := range s.Views {
						s.Views[i].Nodes = append([]cluster.NodeState(nil), s.Views[i].Nodes...)
						s.Views[i].Nodes[0].Slots, s.Views[i].Nodes[1].Slots = s.Views[i].Nodes[1].Slots, s.Views[i].Nodes[0].Slots
						s.Views[i].Coverage = cluster.AnalyzeCoverage(s.Views[i].Nodes)
					}
				case "crossed deadline":
					<-ctx.Done()
					s.Completed = time.Now()
				case "cancellation":
					cancel()
				}
				return s
			})
			observePersistenceTopology(ctx, c, r, o)
			if r.PersistenceFault.StableRedundancy != nil || r.Recovery.TopologyRecovered || r.Recovery.RedundancyRestored || r.Recovery.ClientRecovered || r.Recovery.CombinedStable != nil {
				t.Fatal("unaccepted original topology became recovered")
			}
			if name == "crossed deadline" && (len(r.Samples) != 1 || !r.Samples[0].Excluded || r.PersistenceFault.ReplicaReady != nil || !r.PersistenceFault.RestoreTimedOut) {
				t.Fatal("deadline-excluded sample raised milestones")
			}
			if name == "cancellation" && r.PersistenceFault.RestoreTimedOut {
				t.Fatal("cancellation mislabeled restoration deadline")
			}
		})
	}
}

func TestPersistencePreparationRetriesOnlyObservedTransientBeforeBaseline(t *testing.T) {
	c, r, m, _, _ := persistenceFixture(t, "aof-everysec")
	r.PersistenceFault.Prepared = nil
	m.busyPreparation = 2
	m.probes = 0
	if err := preparePersistence(context.Background(), c, r, m); err != nil || len(r.PersistenceFault.Prepared) != 6 || m.probes != 8 || len(r.PersistenceFault.KillActions) != 0 {
		t.Fatalf("preparation did not await observed quiet state before any mutation: probes=%d err=%v", m.probes, err)
	}
	m.busyPreparation = 10000
	r.PersistenceFault.Prepared = nil
	c.Timeouts.Readiness = config.Duration(5 * time.Millisecond)
	if err := preparePersistence(context.Background(), c, r, m); !errors.Is(err, context.DeadlineExceeded) || len(r.PersistenceFault.Prepared) != 1 || r.PersistenceFault.Prepared[0].Error == "" {
		t.Fatal("busy preparation failed to honor its deadline or preserve final observation")
	}
}

func TestPersistenceRuntimeEvidenceRefusesInventedDurability(t *testing.T) {
	n := environment.Node{ID: "n", RedisClusterID: "redis-n"}
	for _, name := range []string{"missing configuration", "wrong fsync", "loading", "AOF error", "missing files", "memory RDB", "unknown restart policy"} {
		t.Run(name, func(t *testing.T) {
			mode := "aof-always"
			e := persistenceLiveEvidence(mode, n)
			switch name {
			case "missing configuration":
				e.RawPersistenceConfig = ""
			case "wrong fsync":
				e.RawPersistenceConfig = strings.Replace(e.RawPersistenceConfig, "appendfsync\nalways", "appendfsync\neverysec", 1)
			case "loading":
				e.RawPersistence = strings.Replace(e.RawPersistence, "loading:0", "loading:1", 1)
			case "AOF error":
				e.RawPersistence = strings.Replace(e.RawPersistence, "aof_last_write_status:ok", "aof_last_write_status:err", 1)
			case "missing files":
				e.RawDataFiles = ""
			case "memory RDB":
				mode = "memory"
				e = persistenceLiveEvidence(mode, n)
				e.RawDataFiles += "/data/dump.rdb\n"
			case "unknown restart policy":
				e.ContainerRestartPolicy = ""
			}
			if persistenceRuntimeEvidence(e, mode) == nil {
				t.Fatal("incomplete/contradictory raw evidence accepted")
			}
		})
	}
}

func TestPersistenceAOFVerdictKeepsValidationUncertaintyAndObservedLoss(t *testing.T) {
	for _, name := range []string{"missing", "unreadable", "unverified", "memory preserved"} {
		t.Run(name, func(t *testing.T) {
			mode := "aof-always"
			if name == "memory preserved" {
				mode = "memory"
			}
			c, r, m, traffic, o := persistenceFixture(t, mode)
			if err := runPersistenceFixture(context.Background(), c, r, m, traffic, o); err != nil {
				t.Fatal(err)
			}
			r.PersistenceFault.PersistenceRecoveryObserved = nil
			r.PersistenceFault.PersistenceRecoveryAt = nil
			r.Dimensions.PersistenceRecovered = nil
			r.Recovery.DataVerified = false
			switch name {
			case "missing":
				r.Validation.Missing = 1
			case "unreadable":
				r.Validation.Unreadable = 1
			case "unverified":
				r.Validation.Complete = false
				r.Validation.Unverified = 1
			case "memory preserved":
				r.Recovery.DataVerified = true
			}
			updatePersistenceVerdict(r)
			if name == "missing" {
				if r.PersistenceFault.PersistenceRecoveryObserved == nil || *r.PersistenceFault.PersistenceRecoveryObserved || r.Dimensions.PersistenceRecovered == nil || *r.Dimensions.PersistenceRecovered {
					t.Fatal("observed loss became persistence recovery")
				}
			} else if name == "memory preserved" {
				if r.PersistenceFault.PersistenceRecoveryObserved != nil || r.Dimensions.PersistenceRecovered == nil || *r.Dimensions.PersistenceRecovered {
					t.Fatal("unexpected memory preservation became AOF recovery")
				}
			} else if r.PersistenceFault.PersistenceRecoveryObserved != nil || r.Dimensions.PersistenceRecovered != nil {
				t.Fatal("unreadable/unverified became measured preservation failure")
			}
		})
	}
}

func TestPersistenceDiagnosticsKeepPostStabilityRegressionAndPolicyIndependent(t *testing.T) {
	for _, name := range []string{"ordinary service regression", "WAIT regression", "WAITAOF regression", "deadline sample"} {
		t.Run(name, func(t *testing.T) {
			c, r, _, _, _ := persistenceFixture(t, "aof-always")
			at := time.Now()
			r.FaultRequestedAt = &at
			c.Experiment.StableFor = config.Duration(time.Millisecond)
			c.Experiment.PostRecovery = config.Duration(time.Millisecond)
			if name == "WAIT regression" || name == "WAITAOF regression" {
				c.Workload.WritePolicy = "wait"
				if name == "WAITAOF regression" {
					c.Workload.WritePolicy = "waitaof"
				}
				r.Config.Workload.WritePolicy = c.Workload.WritePolicy
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			count := 0
			collect := func(context.Context) report.Sample {
				count++
				stamp := at.Add(time.Duration(count) * time.Millisecond)
				snapshot := *r.Initial
				snapshot.At, snapshot.Completed = stamp, stamp
				good := true
				s := report.Sample{Snapshot: snapshot, Probe: report.Probe{At: stamp, Successful: true, ReadSuccessful: &good, WriteSuccessful: &good}}
				if name == "ordinary service regression" && count >= 3 {
					s.Probe.Successful = false
					bad := false
					s.Probe.ReadSuccessful = &bad
				}
				if name == "WAIT regression" || name == "WAITAOF regression" {
					s.WritePolicyProbe = &report.PolicyProbe{At: stamp, Successful: count < 3}
				}
				if name == "deadline sample" {
					cancel()
				}
				return s
			}
			observePersistenceDiagnostics(ctx, c, r, collect)
			if name == "deadline sample" {
				if len(r.Samples) != 1 || !r.Samples[0].Excluded || r.Recovery.CombinedStable != nil || r.PersistenceFault.DiagnosticServiceStable != nil || r.Recovery.ClientFirstSuccess != nil || r.Recovery.CombinedRecovered {
					t.Fatal("deadline/canceled diagnostics raised recovery milestones")
				}
				return
			}
			if r.Recovery.CombinedStable == nil || r.PersistenceFault.DiagnosticServiceStable == nil || count != 3 {
				t.Fatal("historical sampled stability was lost")
			}
			if name == "ordinary service regression" && (r.Recovery.ClientRecovered || r.Recovery.CombinedRecovered || r.Dimensions.ReadAvailable == nil || *r.Dimensions.ReadAvailable) {
				t.Fatal("post-stability client failure was masked")
			}
			if (name == "WAIT regression" || name == "WAITAOF regression") && (!r.Recovery.ClientRecovered || !r.Recovery.CombinedRecovered || r.Recovery.WritePolicyRecovered == nil || *r.Recovery.WritePolicyRecovered || r.Dimensions.WritePolicyAvailable == nil || *r.Dimensions.WritePolicyAvailable || Outcome(r.Recovery) == "recovered") {
				t.Fatal("basic service conflated with acknowledgement policy recovery")
			}
		})
	}
}

func TestPersistenceMissingValidationWithholdsDiagnosticWrites(t *testing.T) {
	c, r, m, traffic, o := persistenceFixture(t, "aof-always")
	traffic.invalidValidation = true
	if err := runPersistenceFixture(context.Background(), c, r, m, traffic, o); err == nil || r.PersistenceFault.DiagnosticsRequestedAt != nil || r.Recovery.DataVerified {
		t.Fatal("unattempted validation crossed diagnostic-write boundary")
	}
	for _, call := range m.log.calls {
		if call == "diagnostic-set" {
			t.Fatal("diagnostic SET happened without a completed validation attempt")
		}
	}
}

func TestPersistenceQuietPreparationUnknownsAndAfterRestartRewrite(t *testing.T) {
	n := environment.Node{ID: "n", RedisClusterID: "redis-n"}
	e := persistenceLiveEvidence("aof-always", n)
	if persistencePreparationQuiet(e) != nil || persistencePreparationBusy(e) {
		t.Fatal("known quiet preparation rejected")
	}
	for _, field := range []string{"rdb_bgsave_in_progress:0\n", "aof_pending_bio_fsync:0\n", "aof_rewrite_scheduled:0\n"} {
		unknown := e
		unknown.RawPersistence = strings.Replace(unknown.RawPersistence, field, "", 1)
		if persistencePreparationQuiet(unknown) == nil || persistencePreparationBusy(unknown) {
			t.Fatalf("unknown indicator became quiet or an observed transient: %s", field)
		}
	}
	rewriting := e
	rewriting.RawPersistence = strings.Replace(rewriting.RawPersistence, "aof_rewrite_in_progress:0", "aof_rewrite_in_progress:1", 1)
	if !persistencePreparationBusy(rewriting) || persistencePreparationQuiet(rewriting) == nil || persistenceRuntimeEvidence(rewriting, "aof-always") != nil {
		t.Fatal("initial quiet preparation was conflated with healthy post-restart full-sync rewrite")
	}
}

func TestNonRestartAOFPoliciesShareQuietPreparationEvidence(t *testing.T) {
	for _, policy := range []string{"set", "wait", "waitaof"} {
		t.Run(policy, func(t *testing.T) {
			c, r, m, _, _ := persistenceFixture(t, "aof-always")
			c.Experiment.Fault, c.Workload.WritePolicy = "primary-kill", policy
			r.Config, r.PersistenceFault = c, nil
			if err := preparePersistence(context.Background(), c, r, m); err != nil {
				t.Fatal(err)
			}
			if len(r.PersistencePrepared) != 6 {
				t.Fatal("AOF policy preparation skipped an original node")
			}
			for i, sample := range r.PersistencePrepared {
				if sample.Node != r.Environment.Nodes[i] || sample.OriginalRedisID != sample.Node.RedisClusterID || sample.Error != "" || persistenceRuntimeEvidence(sample.Evidence, "aof-always") != nil || persistencePreparationQuiet(sample.Evidence) != nil {
					t.Fatal("common preparation is not actual quiet original-process evidence")
				}
			}
			for _, call := range m.log.calls {
				if strings.HasPrefix(call, "kill") || strings.HasPrefix(call, "restart") || call == "diagnostic-set" || call == "validate" {
					t.Fatalf("preparation performed a fault/write/validation: %s", call)
				}
			}
		})
	}
}

func TestPersistenceFinalRuntimeIsFreshWithoutRewritingEarlyReceipts(t *testing.T) {
	for _, mode := range []string{"memory", "aof-everysec"} {
		t.Run(mode, func(t *testing.T) {
			c, r, m, traffic, o := persistenceFixture(t, mode)
			m.earlyReceiptTransient = true
			if mode == "memory" {
				traffic.missing = 2
			}
			err := runPersistenceFixture(context.Background(), c, r, m, traffic, o)
			if len(r.PersistenceFault.Restored) != 6 || r.PersistenceFault.DiagnosticsRequestedAt == nil || r.PersistenceFault.RestoreTimedOut || (mode != "memory" && err != nil) {
				t.Fatalf("healthy final evidence did not resolve early transient observations: %v", err)
			}
			for i, x := range r.PersistenceFault.Restarts {
				if persistenceRuntimeEvidence(x.After, mode) == nil || persistenceRuntimeEvidence(r.PersistenceFault.Restored[i].Evidence, mode) != nil {
					t.Fatal("past physical receipt was changed instead of retaining a distinct healthy final observation")
				}
			}
		})
	}
}
