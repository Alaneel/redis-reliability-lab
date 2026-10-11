package batch

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/config"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
	"github.com/alanwang/redis-reliability-lab/internal/metrics"
	"github.com/alanwang/redis-reliability-lab/internal/report"
	"github.com/alanwang/redis-reliability-lab/internal/workload"
)

func persistenceFixture(t *testing.T) *report.Report {
	t.Helper()
	c := config.Default()
	c.Experiment.Fault, c.Environment.Persistence = "cluster-restart", "aof-always"
	r := fixture(c, 1)
	r.SchemaVersion = report.LifecycleSchemaVersion
	base := r.StartedAt
	at := func(seconds int) *time.Time { value := base.Add(time.Duration(seconds) * time.Second); return &value }
	snapshot := cluster.Snapshot{At: base, Completed: base}
	for self := 0; self < 6; self++ {
		var raw strings.Builder
		for i := 0; i < 6; i++ {
			flags, parent, slots := "master", "-", ""
			if i >= 3 {
				flags, parent = "slave", strings.Repeat(string(rune('a'+i-3)), 40)
			} else {
				slots = []string{"0-5460", "5461-10922", "10923-16383"}[i]
			}
			if self == i {
				flags = "myself," + flags
			}
			fmt.Fprintf(&raw, "%s 172.22.0.%d:6379@16379 %s %s 0 1 %d connected %s\n", strings.Repeat(string(rune('a'+i)), 40), i+2, flags, parent, i+1, slots)
		}
		nodes, err := cluster.ParseNodes(raw.String())
		if err != nil {
			t.Fatal(err)
		}
		replication := map[string]string{"role": "master", "connected_slaves": "1"}
		if self >= 3 {
			replication = map[string]string{"role": "slave", "master_link_status": "up", "master_sync_in_progress": "0", "master_host": fmt.Sprintf("172.22.0.%d", self-1), "master_port": "6379"}
		}
		node := environment.Node{ID: fmt.Sprintf("node-%d", self+1), ContainerID: strings.Repeat(string(rune('a'+self)), 64), Name: fmt.Sprintf("owned-%d", self+1), Addr: fmt.Sprintf("127.0.0.1:%d", 7000+self), InternalAddr: fmt.Sprintf("172.22.0.%d:6379", self+2), RedisClusterID: strings.Repeat(string(rune('a'+self)), 40)}
		r.Environment.Nodes = append(r.Environment.Nodes, node)
		r.Environment.Volumes = append(r.Environment.Volumes, environment.Volume{Name: fmt.Sprintf("volume-%d", self+1), NodeID: node.ID})
		snapshot.Views = append(snapshot.Views, cluster.View{Node: cluster.Node{ID: node.ID, Addr: node.Addr, InternalAddr: node.InternalAddr}, At: base, RedisID: node.RedisClusterID, Reachable: true, RedisVersion: "8.10.2", ClusterInfo: map[string]string{"cluster_state": "ok", "cluster_slots_assigned": "16384", "cluster_slots_ok": "16384", "cluster_slots_fail": "0", "cluster_slots_pfail": "0"}, Replication: replication, Nodes: nodes, Coverage: cluster.AnalyzeCoverage(nodes)})
	}
	selection, err := cluster.SelectPrimary(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	final := snapshot
	final.At, final.Completed = *at(35), *at(36)
	final.Views = append([]cluster.View(nil), snapshot.Views...)
	for i := range final.Views {
		final.Views[i].At = *at(35)
	}
	r.Initial, r.Final, r.Selection = &snapshot, &final, &selection
	r.FinishedAt = *at(40)
	r.Environment.RetainData, r.Environment.Persistence = true, "aof-always"
	yes := true
	r.Dimensions = &report.RecoveryDimensions{DataVerified: &yes, PersistenceRecovered: &yes, RedundancyRestored: &yes}
	r.Recovery.RedundancyRestored = true
	p := &report.PersistenceFault{Mode: "aof-always", FailureModel: report.PersistenceFailureModel, TrafficDrainedAt: at(1), LedgerFrozenAt: at(1), AllStoppedAt: at(10), RestartRequestedAt: at(11), RestartCompletedAt: at(30), DataCheckedAt: at(32), DiagnosticsRequestedAt: at(33), AcknowledgedFrozen: 4, AmbiguousFrozen: 1, PreCrashSnapshot: &snapshot, DrainToFirstKillMS: number(1000), DrainToLastKillMS: number(7000), AllStoppedToFirstRestartMS: number(1000), PersistenceRecoveryObserved: &yes}
	stable := &report.Milestone{At: *at(31), SinceFaultMS: 29000}
	p.TopologyStable, p.StableRedundancy = stable, stable
	p.DiagnosticServiceStable = &report.Milestone{At: *at(35), SinceFaultMS: 33000}
	p.PersistenceRecoveryAt = &report.Milestone{At: *at(32), SinceFaultMS: 30000}
	r.PersistenceFault = p
	r.Recovery.CombinedStable, r.Recovery.ClientStable, r.Recovery.ClientFirstSuccess = p.DiagnosticServiceStable, p.DiagnosticServiceStable, p.DiagnosticServiceStable
	r.Recovery.TopologyConverged, r.Recovery.FullCoverage = p.TopologyStable, p.TopologyStable
	r.Recovery.ReplicaPromoted, r.Recovery.FailureDetected = nil, nil
	r.Metrics = metrics.Snapshot{Totals: metrics.Summary{Attempts: 10, Successes: 9, Errors: 1}, Phases: map[string]metrics.PhaseSummary{"baseline": {Attempts: 10, Successes: 9, Errors: 1, DurationSeconds: 1, SuccessRate: .9}}}
	r.WorkloadStats.PacingMethodology = workload.PacingMethodology
	r.WorkloadStats.OfferedRequests, r.WorkloadStats.StartedRequests, r.WorkloadStats.CompletedRequests = 10, 10, 10
	r.FaultRequestedAt, r.FaultCompletedAt, r.FaultActionDurationMS = at(2), at(10), number(8000)
	r.Validation.StartedAt, r.Validation.CompletedAt = *at(32), *at(32)
	validation := r.Validation
	p.FrozenValidation = &validation
	configRaw := "repl-diskless-load\nswapdb\nrepl-diskless-sync\nyes\nauto-aof-rewrite-percentage\n0\nmaxmemory\n134217728\nmaxmemory-policy\nnoeviction\nappendonly\nyes\nappendfsync\nalways\naof-use-rdb-preamble\nyes\nno-appendfsync-on-rewrite\nno\naof-load-truncated\nyes\nrdb-del-sync-files\nyes\nsave\n"
	for i, node := range r.Environment.Nodes {
		prepared := environment.ProcessEvidence{ObservedAt: base, ContainerRunning: true, ContainerPID: 100 + i, ContainerStartedAt: "original-start", ContainerRestartPolicy: "no", RedisPONG: true, RedisPID: 1, RedisRunID: fmt.Sprintf("old-run-%d", i), RedisClusterID: node.RedisClusterID, NetworkConnected: true, NetworkCount: 1, RawPersistenceConfig: configRaw, RawPersistence: "loading:0\naof_enabled:1\naof_last_write_status:ok\naof_last_bgrewrite_status:ok\n", RawDataFiles: "/data/nodes.conf\n/data/appendonlydir/appendonly.aof.1.incr.aof\n/data/appendonlydir/appendonly.aof.manifest\n"}
		stopped := environment.ProcessEvidence{ObservedAt: *at(9), ContainerStartedAt: prepared.ContainerStartedAt, ContainerRestartPolicy: "no"}
		p.Prepared = append(p.Prepared, report.NodeProcessSample{Node: node, OriginalRedisID: node.RedisClusterID, Evidence: prepared})
		p.Stopped = append(p.Stopped, report.NodeProcessSample{Node: node, OriginalRedisID: node.RedisClusterID, Evidence: stopped, Error: environment.ErrProcessStopped.Error()})
		after := prepared
		after.ObservedAt = *at(12 + i*2)
		after.ContainerStartedAt = "new-start"
		after.RedisRunID = fmt.Sprintf("new-run-%d", i)
		restartBefore := stopped
		restartBefore.ObservedAt = *at(11 + i*2)
		p.Restarts = append(p.Restarts, &environment.RestartReceipt{RunID: r.RunID, NodeID: node.ID, Node: node, OriginalRedisID: node.RedisClusterID, Before: restartBefore, After: after, RequestedAt: *at(11 + i*2), CompletedAt: *at(12 + i*2)})
		p.RestartActions = append(p.RestartActions, report.NodeAction{Node: node, RequestedAt: *at(11 + i*2), CompletedAt: at(12 + i*2)})
		restored := after
		restored.ObservedAt = *at(31)
		p.Restored = append(p.Restored, report.NodeProcessSample{Node: node, OriginalRedisID: node.RedisClusterID, Evidence: restored})
	}
	for i, index := range []int{3, 4, 5, 0, 1, 2} {
		p.KillActions = append(p.KillActions, report.NodeAction{Node: r.Environment.Nodes[index], RequestedAt: *at(2 + i), CompletedAt: at(3 + i)})
	}
	return r
}

func TestRecoveredPersistenceRequiresFrozenBarrierEvidence(t *testing.T) {
	if err := recoveredPersistenceEvidence(persistenceFixture(t)); err != nil {
		t.Fatal(err)
	}
	cases := map[string]func(*report.Report){
		"pre-diagnostic client recovery": func(r *report.Report) { r.Recovery.ClientStable = &report.Milestone{At: r.StartedAt, SinceFaultMS: 0} },
		"contradictory inventory":        func(r *report.Report) { r.Environment.Nodes[0].RedisClusterID = "foreign" },
		"workload afterfreeze":           func(r *report.Report) { r.Metrics.Phases["fault"] = metrics.PhaseSummary{Attempts: 1} },
		"stale final":                    func(r *report.Report) { r.Final = r.Initial },
		"impossible finished time":       func(r *report.Report) { r.FinishedAt = r.StartedAt },
		"receipt outside action":         func(r *report.Report) { r.PersistenceFault.Restarts[0].RequestedAt = r.StartedAt.Add(25 * time.Second) },
		"missing barrier":                func(r *report.Report) { r.PersistenceFault.AllStoppedAt = nil },
		"five stopped":                   func(r *report.Report) { r.PersistenceFault.Stopped = r.PersistenceFault.Stopped[:5] },
		"generic inspection error":       func(r *report.Report) { r.PersistenceFault.Stopped[0].Error = "permission denied" },
		"autonomous restart":             func(r *report.Report) { r.PersistenceFault.Stopped[0].Evidence.ContainerRestartPolicy = "always" },
		"stopped still alive":            func(r *report.Report) { r.PersistenceFault.Stopped[0].Evidence.ContainerRunning = true },
		"early restart":                  func(r *report.Report) { r.PersistenceFault.Restarts[0].RequestedAt = r.StartedAt },
		"duplicate restart":              func(r *report.Report) { r.PersistenceFault.Restarts[1] = r.PersistenceFault.Restarts[0] },
		"foreign volume":                 func(r *report.Report) { r.Environment.Volumes[0].NodeID = "foreign" },
		"foreign container":              func(r *report.Report) { r.PersistenceFault.Restarts[0].Node.ContainerID = "foreign" },
		"same process": func(r *report.Report) {
			r.PersistenceFault.Restarts[0].After.RedisRunID = r.PersistenceFault.Prepared[0].Evidence.RedisRunID
		},
		"AOF disabled": func(r *report.Report) {
			r.PersistenceFault.Restored[0].Evidence.RawPersistence = "loading:0\naof_enabled:0\naof_last_write_status:ok"
		},
		"loading": func(r *report.Report) {
			r.PersistenceFault.Restored[0].Evidence.RawPersistence = "loading:1\naof_enabled:1\naof_last_write_status:ok"
		},
		"wrong fsync": func(r *report.Report) {
			r.PersistenceFault.Restored[0].Evidence.RawPersistenceConfig = strings.ReplaceAll(r.PersistenceFault.Restored[0].Evidence.RawPersistenceConfig, "always", "everysec")
		},
		"missing files":             func(r *report.Report) { r.PersistenceFault.Restored[0].Evidence.RawDataFiles = "" },
		"unfrozen acknowledgements": func(r *report.Report) { r.PersistenceFault.AcknowledgedFrozen++ },
		"different validation":      func(r *report.Report) { r.PersistenceFault.FrozenValidation.Missing++ },
		"validation after diagnostics": func(r *report.Report) {
			early := r.Validation.StartedAt.Add(-time.Second)
			r.PersistenceFault.DiagnosticsRequestedAt = &early
		},
		"missing persistence dimension": func(r *report.Report) { r.Dimensions.PersistenceRecovered = nil },
		"changed slot owner":            func(r *report.Report) { r.Final.Views[0].Nodes[0].Slots = nil },
		"unobserved stability":          func(r *report.Report) { r.PersistenceFault.StableRedundancy = nil },
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			r := persistenceFixture(t)
			mutate(r)
			if err := recoveredPersistenceEvidence(r); err == nil {
				t.Fatal("contradictory persistence evidence accepted")
			}
		})
	}
}

func TestPersistenceBatchPreservesPartialControlAndCensoring(t *testing.T) {
	r := persistenceFixture(t)
	r.Config.Environment.Persistence = "memory"
	r.Outcome = "partial"
	r.PersistenceFault.Mode = "memory"
	for i := range r.PersistenceFault.Prepared {
		e := &r.PersistenceFault.Prepared[i].Evidence
		e.RawPersistenceConfig = strings.ReplaceAll(strings.ReplaceAll(e.RawPersistenceConfig, "appendonly\nyes", "appendonly\nno"), "aof-use-rdb-preamble\nyes", "aof-use-rdb-preamble\nno")
		e.RawPersistence = "loading:0\naof_enabled:0\n"
		e.RawDataFiles = "/data/nodes.conf\n"
		after := *e
		after.ObservedAt = r.PersistenceFault.Restarts[i].After.ObservedAt
		after.ContainerStartedAt = r.PersistenceFault.Restarts[i].After.ContainerStartedAt
		after.RedisRunID = r.PersistenceFault.Restarts[i].After.RedisRunID
		r.PersistenceFault.Restarts[i].After = after
	}
	r.Recovery.DataVerified = false
	r.Recovery.ReplicaPromoted = nil
	r.PersistenceFault.PersistenceRecoveryAt = nil
	r.PersistenceFault.PersistenceRecoveryObserved = nil
	no := false
	r.Dimensions.DataVerified = &no
	r.Dimensions.PersistenceRecovered = &no
	r.Validation.Missing = 4
	validation := r.Validation
	r.PersistenceFault.FrozenValidation = &validation
	s, err := Run(context.Background(), r.Config, 1, t.TempDir(), func(_ context.Context, c config.Config, root string) (*report.Report, error) {
		fresh := fixture(c, 1)
		r.ConfigSHA256 = fresh.ConfigSHA256
		save(t, root, r)
		return r, nil
	})
	if err == nil || s.Outcomes["partial"] != 1 || s.IncompleteEvidenceTrials != 0 || s.Data.Missing != 4 || s.Events["persistence_recovered"].NotAtRisk != 1 || s.Events["replica_promoted"].NotAtRisk != 1 || s.Events["stable_redundancy"].Observed != 1 {
		t.Fatalf("control evidence lost: %+v %v", s, err)
	}
	r.Config.Environment.Persistence = "aof-always"
	r.PersistenceFault.Mode = "aof-always"
	r.PersistenceFault.TopologyStable = nil
	r.PersistenceFault.StableRedundancy = nil
	r.PersistenceFault.RestoreTimedOut = true
	o := observations(r)
	if o["persistence_recovered"].Status != "censored_deadline" || o["original_topology_stable"].Status != "censored_deadline" || o["persistence_recovered"].SinceFaultRequestedMS != nil {
		t.Fatalf("timeout became successful observation: %+v", o)
	}
}
