package config

import (
	"strings"
	"testing"
)

func TestLifecycleDefinitions(t *testing.T) {
	for _, input := range []string{"experiment:\n  fault: replica-kill", "experiment:\n  fault: replica-isolate", "experiment:\n  fault: primary-kill\n  repair: true", "environment:\n  persistence: aof-everysec\nexperiment:\n  fault: replica-kill", "environment:\n  persistence: aof-always"} {
		c, err := Decode(strings.NewReader(input))
		if err != nil {
			t.Fatalf("%s: %v", input, err)
		}
		if c.Experiment.Fault == "replica-kill" || c.Experiment.Repair || c.Environment.Persistence != "memory" {
			if !c.RetainData() {
				t.Fatal("retention missing")
			}
		}
	}
}
func TestRejectUnsafeLifecycleDefinitions(t *testing.T) {
	for _, input := range []string{"environment:\n  persistence: external", "experiment:\n  fault: replica-isolate\n  repair: true", "experiment:\n  down_duration: 0s", "timeouts:\n  repair: 0s", "experiment:\n  fault: replica-kill\nworkload:\n  max_writes: 100000", "experiment:\n  repair: true\nworkload:\n  value_bytes: 1024", "experiment:\n  fault: replica-kill\ntimeouts:\n  recovery: 5m"} {
		if _, err := Decode(strings.NewReader(input)); err == nil {
			t.Fatalf("accepted %s", input)
		}
	}
	c := Default()
	if c.RetainData() {
		t.Fatal("old default unexpectedly retains files")
	}
}

func TestClusterRestartRequiresRetainedBoundedDataInEveryMode(t *testing.T) {
	for _, mode := range []string{"memory", "aof-everysec", "aof-always"} {
		c, err := Decode(strings.NewReader("environment:\n  persistence: " + mode + "\nexperiment:\n  fault: cluster-restart\n  down_duration: 1s"))
		if err != nil || !c.RetainData() {
			t.Fatalf("cluster restart in %s lost its disposable membership/data resources: %+v %v", mode, c, err)
		}
		c.Workload.MaxWrites = 100000
		if err := c.Validate(); err == nil {
			t.Fatal("accepted unbounded retained traffic")
		}
		c = Default()
		c.Experiment.Fault = "cluster-restart"
		c.Experiment.Repair = true
		if err := c.Validate(); err == nil {
			t.Fatal("accepted primary-only repair flag for whole-cluster restart")
		}
	}
}
