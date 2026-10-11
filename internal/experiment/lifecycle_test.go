package experiment

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/cluster"
	"github.com/Alaneel/redis-reliability-lab/internal/config"
	"github.com/Alaneel/redis-reliability-lab/internal/environment"
	"github.com/Alaneel/redis-reliability-lab/internal/report"
)

func lifecycleSnapshot(t *testing.T, at time.Time) cluster.Snapshot {
	t.Helper()
	s := cluster.Snapshot{At: at, Completed: at}
	ids := []string{"primary-a", "primary-b", "primary-c", "replica-a", "replica-b", "replica-c"}
	for self := 0; self < 6; self++ {
		v := cluster.View{Node: cluster.Node{ID: fmt.Sprintf("node-%d", self+1), Addr: fmt.Sprintf("127.0.0.1:%d", 7000+self), InternalAddr: fmt.Sprintf("172.22.0.%d:6379", self+2)}, At: at, RedisID: ids[self], RedisVersion: "8.10.2", Reachable: true, ClusterInfo: map[string]string{"cluster_state": "ok", "cluster_slots_assigned": "16384", "cluster_slots_ok": "16384", "cluster_slots_fail": "0", "cluster_slots_pfail": "0"}, Replication: map[string]string{"role": "master", "connected_slaves": "1", "master_replid": "repl-id"}, Server: map[string]string{"run_id": fmt.Sprint("process-", self)}, Stats: map[string]string{"sync_full": "0", "sync_partial_ok": "0", "sync_partial_err": "0"}}
		if self >= 3 {
			v.Replication = map[string]string{"role": "slave", "master_link_status": "up", "master_sync_in_progress": "0", "master_host": fmt.Sprintf("172.22.0.%d", self-1), "master_port": "6379"}
		}
		for i, id := range ids {
			n := cluster.NodeState{ID: id, Addr: fmt.Sprintf("172.22.0.%d:6379", i+2), Flags: []string{"master"}, LinkState: "connected"}
			if i < 3 {
				n.Slots = []cluster.SlotRange{{Start: []int{0, 5461, 10923}[i], End: []int{5460, 10922, 16383}[i]}}
			} else {
				n.Flags, n.PrimaryID = []string{"slave"}, ids[i-3]
			}
			if self == i {
				n.Flags = append(n.Flags, "myself")
			}
			v.Nodes = append(v.Nodes, n)
		}
		v.Coverage = cluster.AnalyzeCoverage(v.Nodes)
		s.Views = append(s.Views, v)
	}
	if ready, why := cluster.Ready(s); !ready {
		t.Fatalf("invalid lifecycle fixture: %v", why)
	}
	return s
}

func lifecycleReport(t *testing.T) (*report.Report, cluster.Selection, cluster.Snapshot) {
	t.Helper()
	at := time.Now()
	initial := lifecycleSnapshot(t, at.Add(-5*time.Second))
	selection, err := cluster.SelectPrimary(initial)
	if err != nil {
		t.Fatal(err)
	}
	fault, completed := at.Add(-4*time.Second), at.Add(-3*time.Second)
	r := &report.Report{RunID: "test-run", Config: config.Default(), FaultRequestedAt: &fault, FaultCompletedAt: &completed, Initial: &initial, Selection: &selection, LifecycleFault: &report.LifecycleFault{TargetKind: "replica", Target: selection.Replica, TargetRedisID: selection.ReplicaID, RepairCompleted: true, RepairCompletedAt: &at}}
	r.Config.Experiment.Fault = "replica-kill"
	r.Config.Experiment.StableFor = config.Duration(10 * time.Millisecond)
	r.Config.Experiment.SampleInterval = config.Duration(time.Millisecond)
	return r, selection, initial
}

type helperRestartManager struct {
	context context.Context
	target  environment.Node
	receipt *environment.RestartReceipt
	err     error
}

func validRestartFixture(r *report.Report, target environment.Node) *environment.RestartReceipt {
	at := time.Now()
	return &environment.RestartReceipt{RunID: r.RunID, NodeID: target.ID, Node: target, OriginalRedisID: r.LifecycleFault.TargetRedisID, RequestedAt: at, CompletedAt: at.Add(time.Millisecond), Before: environment.ProcessEvidence{ObservedAt: at, ContainerStartedAt: "old-container-start"}, After: environment.ProcessEvidence{ObservedAt: at.Add(time.Millisecond), RedisClusterID: r.LifecycleFault.TargetRedisID, ContainerRunning: true, ContainerPID: 12345, ContainerStartedAt: "new-container-start", RedisPONG: true, RedisPID: 1, RedisRunID: "new-redis-process", NetworkConnected: true, NetworkCount: 1}}
}

func (m *helperRestartManager) Restart(ctx context.Context, _ *environment.Environment, n environment.Node) (*environment.RestartReceipt, error) {
	m.context, m.target = ctx, n
	return m.receipt, m.err
}

func TestLifecycleRepairRetainsActionEvidenceAndFailure(t *testing.T) {
	for _, name := range []string{"success", "ambiguous failure", "wrong identity", "deadline"} {
		t.Run(name, func(t *testing.T) {
			r, _, _ := lifecycleReport(t)
			r.LifecycleFault.RepairCompleted, r.LifecycleFault.RepairCompletedAt = false, nil
			n := environment.Node{ID: r.LifecycleFault.Target.ID, ContainerID: "owned-container", Addr: r.LifecycleFault.Target.Addr, InternalAddr: r.LifecycleFault.Target.InternalAddr, RedisClusterID: r.LifecycleFault.TargetRedisID}
			m := &helperRestartManager{receipt: validRestartFixture(r, n)}
			switch name {
			case "ambiguous failure":
				m.err = errors.New("injected start response error")
			case "wrong identity":
				m.receipt.After.RedisClusterID = "replacement-node"
			case "deadline":
				m.err = context.DeadlineExceeded
			}
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			err := performLifecycleRepair(ctx, r, m, nil, n)
			l := r.LifecycleFault
			if m.context != ctx || m.target != n || l.RestartReceipt != m.receipt || l.RepairRequestedAt == nil {
				t.Fatal("bounded repair context, target or partial receipt discarded")
			}
			if name == "success" {
				if err != nil || !l.RepairCompleted || !r.Recovery.ProcessRestarted || l.RepairCompletedAt == nil || l.RepairActionDurationMS == nil || *l.RepairActionDurationMS < 0 || l.StableRedundancy != nil || r.Recovery.RedundancyRestored {
					t.Fatalf("physical restart fabricated or lost restoration evidence: %+v %v", l, err)
				}
			} else if err == nil || l.RepairCompleted || l.RepairCompletedAt != nil || l.RepairActionDurationMS != nil || r.Recovery.CombinedRecovered || len(l.Errors) == 0 || l.RepairTimedOut != (name == "deadline") {
				t.Fatalf("failed/contradictory restart became completed repair: %+v %v", l, err)
			}
		})
	}
}

func TestLifecycleRepairRejectsIncompleteOrContradictoryRestartReceipts(t *testing.T) {
	cases := map[string]func(*environment.RestartReceipt){
		"missing completion":         func(x *environment.RestartReceipt) { x.CompletedAt = time.Time{} },
		"completion before request":  func(x *environment.RestartReceipt) { x.CompletedAt = x.RequestedAt.Add(-time.Millisecond) },
		"wrong container":            func(x *environment.RestartReceipt) { x.Node.ContainerID = "another-container" },
		"changed endpoint":           func(x *environment.RestartReceipt) { x.Node.Addr = "127.0.0.1:9000" },
		"wrong run":                  func(x *environment.RestartReceipt) { x.RunID = "another-run" },
		"wrong node":                 func(x *environment.RestartReceipt) { x.NodeID = "node-6" },
		"before still running":       func(x *environment.RestartReceipt) { x.Before.ContainerRunning = true },
		"before start missing":       func(x *environment.RestartReceipt) { x.Before.ContainerStartedAt = "" },
		"after start missing":        func(x *environment.RestartReceipt) { x.After.ContainerStartedAt = "" },
		"after start unchanged":      func(x *environment.RestartReceipt) { x.After.ContainerStartedAt = x.Before.ContainerStartedAt },
		"network detached":           func(x *environment.RestartReceipt) { x.After.NetworkConnected = false },
		"multiple networks":          func(x *environment.RestartReceipt) { x.After.NetworkCount = 2 },
		"after Redis process absent": func(x *environment.RestartReceipt) { x.After.RedisRunID = "" },
		"process unchanged":          func(x *environment.RestartReceipt) { x.Before.RedisRunID = x.After.RedisRunID },
		"observation absent":         func(x *environment.RestartReceipt) { x.After.ObservedAt = time.Time{} },
		"before observation absent":  func(x *environment.RestartReceipt) { x.Before.ObservedAt = time.Time{} },
		"observation reversed":       func(x *environment.RestartReceipt) { x.After.ObservedAt = x.Before.ObservedAt.Add(-time.Millisecond) },
	}
	for name, mutate := range cases {
		t.Run(name, func(t *testing.T) {
			r, _, _ := lifecycleReport(t)
			r.LifecycleFault.RepairCompleted, r.LifecycleFault.RepairCompletedAt = false, nil
			target := environment.Node{ID: r.LifecycleFault.Target.ID, ContainerID: "owned-container", Addr: r.LifecycleFault.Target.Addr, InternalAddr: r.LifecycleFault.Target.InternalAddr, RedisClusterID: r.LifecycleFault.TargetRedisID}
			m := &helperRestartManager{receipt: validRestartFixture(r, target)}
			mutate(m.receipt)
			if err := performLifecycleRepair(context.Background(), r, m, nil, target); err == nil || r.LifecycleFault.RepairCompleted || r.Recovery.ProcessRestarted || r.Recovery.CombinedRecovered || r.LifecycleFault.RepairCompletedAt != nil || r.LifecycleFault.RestartReceipt == nil {
				t.Fatal("incomplete or contradictory restart adapter evidence became an accepted repair")
			}
		})
	}
}

func TestLifecycleSynchronizationRequiresFreshStableInterval(t *testing.T) {
	r, selected, initial := lifecycleReport(t)
	start := *r.LifecycleFault.RepairCompletedAt
	r.Recovery.CombinedStable = milestone(start.Add(-time.Second), *r.FaultRequestedAt)
	r.Recovery.CombinedRecovered, r.Recovery.ClientRecovered, r.Recovery.TopologyRecovered = true, true, true
	var combined, service, policy Stability
	count := 0
	collect := func(context.Context) report.Sample {
		at := start.Add(time.Duration(count) * 5 * time.Millisecond)
		s := lifecycleSnapshot(t, at)
		ok := count != 1
		count++
		return report.Sample{Snapshot: s, Probe: report.Probe{At: at, Successful: ok}}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	observeLifecycleSynchronization(ctx, r.Config, r, selected, initial, &combined, &service, &policy, collect)
	if count != 5 || !r.Recovery.CombinedRecovered || !r.Recovery.RedundancyRestored || r.LifecycleFault.StableRedundancy == nil || !r.LifecycleFault.StableRedundancy.At.Equal(start.Add(20*time.Millisecond)) || r.LifecycleFault.StableRedundancySinceRepairMS == nil || *r.LifecycleFault.StableRedundancySinceRepairMS != 20 {
		t.Fatalf("earlier service success or a negative sample bypassed fresh repair stability: count=%d lifecycle=%+v", count, r.LifecycleFault)
	}
	if r.Recovery.CombinedStable.At.Equal(r.LifecycleFault.StableRedundancy.At) || r.Dimensions == nil || r.LifecycleFault.Dimensions != r.Dimensions || r.Dimensions.DataVerified != nil || r.Dimensions.PersistenceRecovered != nil {
		t.Fatal("historical service conflated with redundancy, or unmeasured correctness/durability invented")
	}
}

func TestLifecycleSynchronizationDeadlineAndCancellationExcludeLateSample(t *testing.T) {
	for _, canceled := range []bool{false, true} {
		t.Run(fmt.Sprint("canceled=", canceled), func(t *testing.T) {
			r, selected, initial := lifecycleReport(t)
			r.Recovery.CombinedStable = milestone(time.Now().Add(-time.Second), *r.FaultRequestedAt)
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
			defer cancel()
			collect := func(ctx context.Context) report.Sample {
				if canceled {
					cancel()
				} else {
					<-ctx.Done()
				}
				at := time.Now()
				return report.Sample{Snapshot: lifecycleSnapshot(t, at), Probe: report.Probe{At: at, Successful: true}}
			}
			var combined, service, policy Stability
			observeLifecycleSynchronization(ctx, r.Config, r, selected, initial, &combined, &service, &policy, collect)
			l := r.LifecycleFault
			if len(r.Samples) != 1 || !r.Samples[0].Excluded || l.StableRedundancy != nil || l.SyncComplete != nil || l.RepairSnapshot != nil || r.Recovery.CombinedRecovered || r.Recovery.RedundancyRestored || l.RepairTimedOut == canceled || r.Recovery.CombinedStable == nil {
				t.Fatalf("late sample, cancellation, deadline or historical service misclassified: %+v", l)
			}
		})
	}
}

func TestLifecycleSynchronizationRejectsUnsynchronizedReplica(t *testing.T) {
	r, selected, initial := lifecycleReport(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	collect := func(context.Context) report.Sample {
		at := time.Now()
		s := lifecycleSnapshot(t, at)
		s.Views[3].Replication["master_sync_in_progress"] = "1"
		s.Views[3].Replication["master_link_status"] = "down"
		return report.Sample{Snapshot: s, Probe: report.Probe{At: at, Successful: true}}
	}
	var combined, service, policy Stability
	observeLifecycleSynchronization(ctx, r.Config, r, selected, initial, &combined, &service, &policy, collect)
	if r.LifecycleFault.SyncStarted == nil || r.LifecycleFault.LinkUp != nil || r.LifecycleFault.SyncComplete != nil || r.LifecycleFault.StableRedundancy != nil || !r.LifecycleFault.RepairTimedOut || r.Recovery.CombinedRecovered {
		t.Fatal("available clients or syncing replica were mistaken for restored redundancy")
	}
}

func TestReplicaMilestonesDoNotRequireOrInventPromotion(t *testing.T) {
	r, selected, initial := lifecycleReport(t)
	s := lifecycleSnapshot(t, time.Now())
	s.Views[3].Reachable = false
	s.Views[3].Errors = []string{"connection refused"}
	s.Views[0].Replication["connected_slaves"] = "0"
	for i := range s.Views {
		s.Views[i].Nodes[3].Flags = append(s.Views[i].Nodes[3].Flags, "fail")
	}
	sample := report.Sample{Phase: "fault", Snapshot: s, Probe: report.Probe{At: s.Completed, Successful: true}, TopologyOK: true}
	r.LifecycleFault.RepairCompleted = false
	updateReplicaMilestones(r, sample, selected, *r.FaultRequestedAt)
	updateLifecycleMilestones(r, sample, selected, initial)
	if r.Recovery.ReplicaPromoted != nil || r.Recovery.FailureDetected == nil || r.Recovery.FullCoverage == nil || r.Recovery.ClientFirstSuccess == nil || r.LifecycleFault.ReplicationInterrupted == nil || *r.Dimensions.MembershipHealthy || *r.Dimensions.RedundancyRestored || !*r.Dimensions.LeadersAvailable || !*r.Dimensions.SlotCoverage || *r.Dimensions.TargetProcessLive {
		t.Fatal("replica failure conflated promotion, slot health, process liveness or replication health")
	}
}

func TestLifecycleResyncCountersAndUnknownFailedReads(t *testing.T) {
	r, selected, initial := lifecycleReport(t)
	r.LifecycleFault.FaultSnapshot = &initial
	s := lifecycleSnapshot(t, time.Now())
	s.Views[0].Stats["sync_full"] = "1"
	writeOK, readOK := true, false
	sample := report.Sample{Phase: "reconciliation", Snapshot: s, Probe: report.Probe{At: s.Completed, Successful: false, WriteSuccessful: &writeOK, ReadSuccessful: &readOK}}
	updateLifecycleMilestones(r, sample, selected, initial)
	if r.LifecycleFault.Resynchronization == nil || r.LifecycleFault.Resynchronization.Inference != "full_sync_counter_increased" || r.LifecycleFault.FullSyncObserved == nil || r.LifecycleFault.PartialSyncObserved != nil || r.Dimensions.WriteAvailable == nil || !*r.Dimensions.WriteAvailable || r.Dimensions.ReadAvailable == nil || *r.Dimensions.ReadAvailable {
		t.Fatal("measured counter delta or independent read/write results were lost")
	}
	sample.Probe = report.Probe{At: s.Completed, Successful: false}
	updateLifecycleMilestones(r, sample, selected, initial)
	if r.Dimensions.ReadAvailable != nil || r.Dimensions.WriteAvailable != nil {
		t.Fatal("legacy failed combined probe invented separate availability")
	}
}

func TestRunLifecycleRequiresRestartBeforeProvisioning(t *testing.T) {
	for _, fault := range []string{"replica-kill", "primary-kill"} {
		c := config.Default()
		c.Experiment.Fault, c.Experiment.Repair = fault, fault == "primary-kill"
		m := &fakeManager{env: testManagedEnvironment()}
		provisioned := false
		m.afterProvision = func() { provisioned = true }
		r, err := quietRunner(m).Run(context.Background(), c, t.TempDir())
		if r != nil || err == nil || !strings.Contains(err.Error(), "restart manager") || provisioned || m.cleaned != nil || m.kills != 0 {
			t.Fatalf("unsupported restart capability reached resources: fault=%s report=%+v err=%v", fault, r, err)
		}
	}
}

func TestLifecycleDimensionsRejectUnknownMembersAndUnreachableOwners(t *testing.T) {
	for _, name := range []string{"unknown member", "duplicate member", "changed endpoint", "unreachable owner", "owner actually replica", "wrong sync primary"} {
		t.Run(name, func(t *testing.T) {
			r, selected, initial := lifecycleReport(t)
			s := lifecycleSnapshot(t, time.Now())
			switch name {
			case "unknown member":
				s.Views[1].Nodes[5].ID = "unknown-redis-id"
			case "duplicate member":
				s.Views[1].Nodes[5].ID = s.Views[1].Nodes[4].ID
			case "changed endpoint":
				s.Views[5].Node.Addr = "127.0.0.1:9001"
			case "unreachable owner":
				s.Views[0].Reachable = false
				s.Views[0].Errors = []string{"connection refused"}
			case "owner actually replica":
				s.Views[0].Replication["role"] = "slave"
			case "wrong sync primary":
				s.Views[3].Nodes[3].PrimaryID = selected.ReplicaID
			}
			updateLifecycleMilestones(r, report.Sample{Snapshot: s, Probe: report.Probe{At: s.Completed}}, selected, initial)
			if name == "wrong sync primary" {
				if r.LifecycleFault.LinkUp != nil || r.LifecycleFault.SyncComplete != nil {
					t.Fatal("link to incorrect primary established intended replica recovery")
				}
			} else if name != "owner actually replica" && (r.Dimensions.MembershipHealthy == nil || *r.Dimensions.MembershipHealthy || r.LifecycleFault.MembershipRejoined != nil) {
				t.Fatal("unknown or inconsistent identities established managed membership recovery")
			}
			if (name == "unreachable owner" || name == "owner actually replica") && (r.Dimensions.SlotCoverage == nil || *r.Dimensions.SlotCoverage || *r.Dimensions.LeadersAvailable) {
				t.Fatal("assigned gossip slots with unreachable owner established usable coverage")
			}
		})
	}
}

func TestLifecycleRepairDoesNotHideUnavailableWAITPolicy(t *testing.T) {
	r, selected, initial := lifecycleReport(t)
	r.Config.Workload.WritePolicy = "wait"
	previouslyAvailable := true
	r.Recovery.WritePolicyRecovered = &previouslyAvailable
	start, count := *r.LifecycleFault.RepairCompletedAt, 0
	collect := func(context.Context) report.Sample {
		at := start.Add(time.Duration(count) * 10 * time.Millisecond)
		count++
		return report.Sample{Snapshot: lifecycleSnapshot(t, at), Probe: report.Probe{At: at, Successful: true}, WritePolicyProbe: &report.PolicyProbe{At: at, Successful: false}}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var combined, service, policy Stability
	observeLifecycleSynchronization(ctx, r.Config, r, selected, initial, &combined, &service, &policy, collect)
	r.Recovery.DataVerified = true
	if !r.Recovery.CombinedRecovered || !r.Recovery.RedundancyRestored || r.LifecycleFault.StableRedundancy == nil || r.Recovery.WritePolicyRecovered == nil || *r.Recovery.WritePolicyRecovered || r.Dimensions.WritePolicyAvailable == nil || *r.Dimensions.WritePolicyAvailable || Outcome(r.Recovery) != "partial" {
		t.Fatal("physical/replication repair converted unavailable configured WAIT policy to full recovery")
	}
}
