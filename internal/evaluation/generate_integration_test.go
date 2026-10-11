//go:build integration

package evaluation

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/cluster"
	"github.com/Alaneel/redis-reliability-lab/internal/config"
	"github.com/Alaneel/redis-reliability-lab/internal/environment"
	"github.com/Alaneel/redis-reliability-lab/internal/experiment"
	"github.com/Alaneel/redis-reliability-lab/internal/report"
)

// This opt-in collection test performs actual managed faults and records labels
// separately. No inference is called while constructing or freezing the dataset.
func TestDockerGenerateIncidentDataset(t *testing.T) {
	root := os.Getenv("REDIS_LAB_DATASET")
	if root == "" {
		t.Skip("set REDIS_LAB_DATASET to opt into fresh dataset generation")
	}
	root, err := filepath.Abs(root)
	if err != nil {
		t.Fatal(err)
	}
	appendControls := os.Getenv("REDIS_LAB_DATASET_APPEND_CONTROLS") == "1"
	if _, err := os.Stat(filepath.Join(root, "manifest.json")); err == nil && !appendControls {
		t.Fatal("refusing to overwrite an existing dataset")
	}
	if err := os.MkdirAll(root, 0755); err != nil {
		t.Fatal(err)
	}
	d := Dataset{SchemaVersion: 1, RuleVersion: "rules-2", Cases: []Case{}, Methodology: []string{
		"All base cases are freshly provisioned six-node Redis clusters; no diagnostic inference is run during collection or label construction.",
		"Development and heldout trials use separate physical clusters, workload seeds, rates, sample intervals and timing. Derivatives share their parent's trial family and partition.",
		"Targets are selected by existing deterministic cluster inventory, with fresh Redis identities and ports; target shard position is not randomized. This limits generalization.",
		"Physical primary/replica unavailability and redundancy labels come from applied controlled actions; transient client and value outcomes are labelled from independent runner counters, not diagnostic predictions.",
		"Longer network holds allow server observations to reflect disruptions; absence of detectable evidence can still cause false negatives, retained in evaluation.",
		"Persistence faults are labelled as primary and replica unavailability even when intermediate node observations are absent. The engine must not read stopped-process action fields to recover those labels.",
		"Incomplete derivatives retain true physical incident labels, rather than changing labels to match missing observations. They are not additional independent trials.",
	}}
	index := 0
	if appendControls {
		var err error
		d, _, err = Load(filepath.Join(root, "manifest.json"))
		if err != nil {
			t.Fatal(err)
		}
		index = len(d.Cases)
		d.Methodology = append(d.Methodology, "Healthy controls were collected in a corrected separate invocation after the initial directory-writing defect; original failed collection log retained.")
	}
	scenarios := []string{"healthy", "primary", "replica", "primary-network", "replica-network", "wait", "memory", "aof"}
	if appendControls {
		scenarios = []string{"healthy"}
	}
	for pi, partition := range []string{"development", "heldout"} {
		if selected := os.Getenv("REDIS_LAB_DATASET_PARTITION"); selected != "" && selected != partition {
			continue
		}
		for _, scenario := range scenarios {
			t.Run(partition+"-"+scenario, func(t *testing.T) {
				c := config.Default()
				c.Name = "dataset-trial"
				c.Environment.NodeTimeout = config.Duration(time.Second)
				c.Experiment.Baseline = config.Duration(2 * time.Second)
				c.Experiment.PostRecovery = config.Duration(time.Second)
				c.Experiment.StableFor = config.Duration(time.Second)
				c.Experiment.SampleInterval = config.Duration(time.Duration(200+pi*110) * time.Millisecond)
				c.Workload.Seed = int64(27101 + index*101 + pi*10000)
				c.Workload.RatePerSecond = float64(100 + pi*70)
				c.Experiment.DownDuration = config.Duration(time.Duration(3+pi) * time.Second)
				c.Experiment.IsolationDuration = config.Duration(time.Duration(9+pi*2) * time.Second)
				c.Timeouts.Recovery = config.Duration(30 * time.Second)
				switch scenario {
				case "replica":
					c.Experiment.Fault = "replica-kill"
				case "primary-network":
					c.Experiment.Fault = "primary-isolate"
				case "replica-network":
					c.Experiment.Fault = "replica-isolate"
				case "wait":
					c.Experiment.Fault = "replica-kill"
					c.Workload.WritePolicy = "wait"
				case "memory":
					c.Experiment.Fault = "cluster-restart"
				case "aof":
					c.Experiment.Fault = "cluster-restart"
					c.Environment.Persistence = "aof-always"
				}
				if err := c.Validate(); err != nil {
					t.Fatal(err)
				}
				ctx, cancel := context.WithTimeout(context.Background(), 4*time.Minute)
				defer cancel()
				var r *report.Report
				if scenario == "healthy" {
					r = healthyTrial(t, ctx, c, root)
				} else {
					var runErr error
					r, runErr = (experiment.Runner{Manager: environment.New()}).Run(ctx, c, filepath.Join(root, "runs"))
					if r == nil {
						t.Fatal(runErr)
					}
					if runErr != nil {
						t.Logf("retained experimental outcome=%s err=%v", r.Outcome, runErr)
					}
				}
				if !r.Cleanup.Successful {
					t.Fatalf("unverified cleanup for %s", r.RunID)
				}
				inventory, err := environment.New().Discover(context.Background(), r.RunID)
				if err != nil || inventory.NetworkID != "" || len(inventory.Nodes) > 0 || len(inventory.Volumes) > 0 {
					t.Fatalf("resources remain: %+v %v", inventory, err)
				}
				// Independently retain fault actions in the original report as evaluation
				// provenance; the allowlist inference adapter cannot ingest those fields.
				label := Label{Incident: scenario != "healthy", Cause: scenario, Categories: []string{}, Basis: []string{"Fresh managed trial; physical action and cleanup receipts retained in source report"}, StartedAt: r.FaultRequestedAt, EndedAt: r.FaultCompletedAt}
				if r.Selection != nil {
					label.Target = r.Selection.Target.ID
					if r.LifecycleFault != nil && r.LifecycleFault.TargetKind == "replica" {
						label.Target = r.LifecycleFault.Target.ID
					}
				}
				if scenario == "memory" || scenario == "aof" {
					label.Target = "all-six-managed-nodes"
				}
				switch scenario {
				case "primary", "primary-network":
					label.Categories = append(label.Categories, "primary_unavailability", "redundancy_degradation")
				case "replica", "replica-network", "wait":
					label.Categories = append(label.Categories, "replica_unavailability", "redundancy_degradation")
				case "memory", "aof":
					label.Categories = append(label.Categories, "primary_unavailability", "replica_unavailability", "redundancy_degradation")
				}
				if r.Metrics.Totals.Errors > 0 {
					label.Categories = append(label.Categories, "client_disruption")
					label.Basis = append(label.Basis, "Terminal client metric errors > 0")
				}
				if r.Validation.Missing > 0 || r.Validation.Corrupt > 0 {
					label.Categories = append(label.Categories, "data_validation_failure")
					label.Basis = append(label.Basis, "Final immutable acknowledged-value ledger has missing or unequal values")
				}
				if r.Validation.Unreadable > 0 || r.Validation.Unverified > 0 {
					label.Categories = append(label.Categories, "validation_incomplete")
				}
				// Independent laboratory outcome counters distinguish valid short replies
				// from command failure and never treat a short WAIT as a failed SET.
				counts := r.WorkloadStats.PolicyTotals
				if counts.WaitShort+counts.WaitAOFShort > 0 {
					label.Categories = append(label.Categories, "acknowledgement_degradation")
					label.Basis = append(label.Basis, "Workload policy ledger recorded valid short acknowledgements")
				}
				index++
				caseID := fmt.Sprintf("case-%03d", index)
				family := r.RunID
				path := filepath.Join(root, "runs", r.RunID, "report.json")
				b, err := os.ReadFile(path)
				if err != nil {
					t.Fatal(err)
				}
				sum := sha256.Sum256(b)
				rel, _ := filepath.Rel(root, path)
				d.Cases = append(d.Cases, Case{ID: caseID, Partition: partition, Family: family, Report: rel, SHA256: hex.EncodeToString(sum[:]), Label: label})
				if scenario == "primary" && partition == "heldout" {
					// Evidence deletion is a documented adversarial derivative of a real incident,
					// not a synthetic successful experiment. It stays in the same trial family.
					incomplete := map[string]any{"schema_version": 3, "samples": []any{}, "validation": map[string]any{}}
					b, _ := json.MarshalIndent(incomplete, "", "  ")
					derivative := filepath.Join(root, "inputs", caseID+"-incomplete.json")
					os.MkdirAll(filepath.Dir(derivative), 0755)
					if err := os.WriteFile(derivative, b, 0644); err != nil {
						t.Fatal(err)
					}
					sum := sha256.Sum256(b)
					rel, _ := filepath.Rel(root, derivative)
					d.Cases = append(d.Cases, Case{ID: caseID + "-incomplete", Partition: partition, Family: family, Report: rel, SHA256: hex.EncodeToString(sum[:]), Label: label})
				}
				b, err = json.MarshalIndent(d, "", "  ")
				if err != nil {
					t.Fatal(err)
				}
				if err = os.WriteFile(filepath.Join(root, "manifest.json"), append(b, '\n'), 0644); err != nil {
					t.Fatal(err)
				}
				t.Logf("case=%s partition=%s scenario=%s run=%s outcome=%s", caseID, partition, scenario, r.RunID, r.Outcome)
			})
		}
	}
	if err := d.Validate(); err != nil {
		t.Fatal(err)
	}
}

func healthyTrial(t *testing.T, ctx context.Context, c config.Config, root string) *report.Report {
	t.Helper()
	random := make([]byte, 12)
	if _, err := rand.Read(random); err != nil {
		t.Fatal(err)
	}
	id := hex.EncodeToString(random)
	m := environment.New()
	env, err := m.Provision(ctx, id, environment.Config{Image: c.Environment.RedisImage, NodeTimeout: c.Environment.NodeTimeout.Value()})
	if err != nil {
		t.Fatal(err)
	}
	r := &report.Report{SchemaVersion: 3, RunID: id, StartedAt: time.Now(), Config: c, Environment: env, Samples: []report.Sample{}, Outcome: "healthy-control"}
	defer func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		r.Cleanup.Attempted = true
		err := m.Cleanup(cleanupCtx, env)
		r.Cleanup.Successful = err == nil
		if err != nil {
			t.Error(err)
		}
		r.FinishedAt = time.Now()
		if err := os.MkdirAll(filepath.Join(root, "runs", id), 0755); err != nil {
			t.Fatal(err)
		}
		if err := report.Write(filepath.Join(root, "runs", id), *r); err != nil {
			t.Fatal(err)
		}
	}()
	nodes := []cluster.Node{}
	for _, n := range env.Nodes {
		nodes = append(nodes, cluster.Node{ID: n.ID, Addr: n.Addr, InternalAddr: n.InternalAddr})
	}
	o := cluster.New(nodes, time.Second)
	defer o.Close()
	readyCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	for {
		snap := o.Observe(readyCtx)
		if ok, _ := cluster.Ready(snap); ok {
			r.Initial = &snap
			break
		}
		select {
		case <-readyCtx.Done():
			t.Fatal(readyCtx.Err())
		case <-time.After(200 * time.Millisecond):
		}
	}
	client, err := experiment.NewClient(env, time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	keys, err := experiment.ProbeKeys(*r.Initial)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 4; i++ {
		snap := o.Observe(ctx)
		probe := experiment.ProbeAll(ctx, client, keys, "healthy-control", time.Second)
		if !probe.Successful {
			t.Fatalf("healthy control probe failed: %+v", probe)
		}
		r.Samples = append(r.Samples, report.Sample{Snapshot: snap, Probe: probe})
		time.Sleep(c.Experiment.SampleInterval.Value())
	}
	final := o.Observe(ctx)
	r.Final = &final
	return r
}
