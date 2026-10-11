package report

import (
	"fmt"
	"io"
	"time"

	"github.com/alanwang/redis-reliability-lab/internal/cluster"
	"github.com/alanwang/redis-reliability-lab/internal/environment"
)

// LifecycleSchemaVersion adds independent replica lifecycle and persistence
// observations. Existing schema 1 and 2 reports retain their original fields.
const LifecycleSchemaVersion = 3

// RecoveryDimensions contains independently observed conditions. A nil flag
// means unobserved, not healthy or failed. Persistence recovery must be based on
// a persistence experiment; ordinary replication readiness cannot establish it.
type RecoveryDimensions struct {
	TargetRole           string `json:"target_role,omitempty"`
	TargetProcessLive    *bool  `json:"target_process_live,omitempty"`
	MembershipHealthy    *bool  `json:"membership_healthy,omitempty"`
	SlotCoverage         *bool  `json:"slot_coverage,omitempty"`
	LeadersAvailable     *bool  `json:"leaders_available,omitempty"`
	ReadAvailable        *bool  `json:"read_available,omitempty"`
	WriteAvailable       *bool  `json:"write_available,omitempty"`
	DataVerified         *bool  `json:"data_verified,omitempty"`
	WritePolicyAvailable *bool  `json:"write_policy_available,omitempty"`
	RedundancyRestored   *bool  `json:"redundancy_restored,omitempty"`
	PersistenceRecovered *bool  `json:"persistence_recovered,omitempty"`
}

// LifecycleFault separates completion of the physical repair action from
// sampled membership, replication readiness and sustained redundancy.
// FaultSnapshot is the accepted observation before repair, while RepairSnapshot
// is an observation after repair; neither is an atomic server event snapshot.
type LifecycleFault struct {
	TargetKind                    string                             `json:"target_kind"`
	Target                        cluster.Node                       `json:"target"`
	TargetRedisID                 string                             `json:"target_redis_id"`
	RepairKind                    string                             `json:"repair_kind,omitempty"`
	RestartReceipt                *environment.RestartReceipt        `json:"restart_receipt,omitempty"`
	RepairRequestedAt             *time.Time                         `json:"repair_requested_at,omitempty"`
	RepairCompletedAt             *time.Time                         `json:"repair_completed_at,omitempty"`
	RepairActionDurationMS        *float64                           `json:"repair_action_duration_ms,omitempty"`
	FaultSnapshot                 *cluster.Snapshot                  `json:"fault_snapshot,omitempty"`
	RepairSnapshot                *cluster.Snapshot                  `json:"repair_snapshot,omitempty"`
	ProcessSamples                []ProcessSample                    `json:"process_samples,omitempty"`
	FailureDetected               *Milestone                         `json:"failure_detected,omitempty"`
	ReplicationInterrupted        *Milestone                         `json:"replication_interrupted,omitempty"`
	TargetReachable               *Milestone                         `json:"target_reachable,omitempty"`
	MembershipRejoined            *Milestone                         `json:"membership_rejoined,omitempty"`
	LinkUp                        *Milestone                         `json:"link_up,omitempty"`
	SyncStarted                   *Milestone                         `json:"sync_started,omitempty"`
	SyncComplete                  *Milestone                         `json:"sync_complete,omitempty"`
	FullSyncObserved              *Milestone                         `json:"full_sync_observed,omitempty"`
	PartialSyncObserved           *Milestone                         `json:"partial_sync_observed,omitempty"`
	StableRedundancy              *Milestone                         `json:"stable_redundancy,omitempty"`
	StableRedundancySinceRepairMS *float64                           `json:"stable_redundancy_since_repair_completed_ms,omitempty"`
	Resynchronization             *cluster.ResynchronizationEvidence `json:"resynchronization,omitempty"`
	RepairCompleted               bool                               `json:"repair_completed"`
	RepairTimedOut                bool                               `json:"repair_timed_out"`
	Dimensions                    *RecoveryDimensions                `json:"dimensions,omitempty"`
	Errors                        []string                           `json:"errors,omitempty"`
}

func observedFlag(v *bool) string {
	if v == nil {
		return "unobserved"
	}
	return fmt.Sprint(*v)
}

func (d RecoveryDimensions) markdown(b io.Writer) {
	role := d.TargetRole
	if role == "" {
		role = "unobserved"
	}
	fmt.Fprintf(b, "\n## Independent recovery dimensions\n\nTarget observed role: `%s`.\n\n| Independent condition | Observed |\n|---|---|\n", role)
	for _, condition := range []struct {
		name string
		flag *bool
	}{
		{"Target process live", d.TargetProcessLive}, {"Cluster membership healthy", d.MembershipHealthy},
		{"Full usable slot coverage", d.SlotCoverage}, {"All shard leaders available", d.LeadersAvailable},
		{"Client reads available", d.ReadAvailable}, {"Client writes available", d.WriteAvailable},
		{"Acknowledged data verified", d.DataVerified}, {"Configured write policy available", d.WritePolicyAvailable},
		{"Replication redundancy restored", d.RedundancyRestored}, {"Persistence recovery verified", d.PersistenceRecovered},
	} {
		fmt.Fprintf(b, "| %s | %s |\n", condition.name, observedFlag(condition.flag))
	}
}

func (l LifecycleFault) markdown(b io.Writer) {
	fmt.Fprintf(b, "\n## Managed node lifecycle\n\nTarget: %s `%s` (Redis ID `%s`). Repair: `%s`.\n", l.TargetKind, l.Target.ID, l.TargetRedisID, l.RepairKind)
	fmt.Fprintf(b, "\nRepair action completed: %t. Repair or synchronization deadline expired: %t.\n", l.RepairCompleted, l.RepairTimedOut)
	if l.RepairActionDurationMS != nil {
		fmt.Fprintf(b, "Repair action duration: %.3f ms.\n", *l.RepairActionDurationMS)
	}
	if l.StableRedundancySinceRepairMS != nil {
		fmt.Fprintf(b, "Stable redundancy first observed %.3f ms after repair action completed.\n", *l.StableRedundancySinceRepairMS)
	}
	fmt.Fprintln(b, "\nA completed restart or reconnect does not establish topology or replication recovery. Link-up and completed initial synchronization are sampled readiness conditions, not zero lag or durability guarantees.")
	fmt.Fprint(b, "\n| Lifecycle observation | Since fault orchestration began, ms |\n|---|---:|\n")
	for _, entry := range []struct {
		name string
		at   *Milestone
	}{
		{"Target failure detected", l.FailureDetected}, {"Replication interruption observed", l.ReplicationInterrupted},
		{"Target reachable after repair", l.TargetReachable}, {"Membership rejoined", l.MembershipRejoined},
		{"Replica link up", l.LinkUp}, {"Synchronization in progress", l.SyncStarted}, {"Synchronization complete", l.SyncComplete},
		{"Primary full-sync counter increased", l.FullSyncObserved}, {"Primary partial-sync counter increased", l.PartialSyncObserved},
		{"Replication redundancy sustained", l.StableRedundancy},
	} {
		if entry.at == nil {
			fmt.Fprintf(b, "| %s | unobserved |\n", entry.name)
		} else {
			fmt.Fprintf(b, "| %s | %.3f |\n", entry.name, entry.at.SinceFaultMS)
		}
	}
	if l.Resynchronization != nil {
		r := l.Resynchronization
		fmt.Fprintf(b, "\nResynchronization evidence: `%s`; primary Redis ID `%s`. Counters are server-wide primary deltas, not an exact target transfer or byte-loss measurement.\n", r.Inference, r.PrimaryRedisID)
		for _, why := range r.Reasons {
			fmt.Fprintf(b, "\nResynchronization limitation: %s\n", why)
		}
	}
	for _, err := range l.Errors {
		fmt.Fprintf(b, "\nLifecycle error: %s\n", err)
	}
}
