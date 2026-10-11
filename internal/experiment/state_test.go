package experiment

import (
	"testing"
	"time"
)

func TestLifecycle(t *testing.T) {
	at := time.Now()
	s := NewLifecycle(at)
	if s.Move(Observing, at) == nil {
		t.Fatal("skipped provisioning")
	}
	for _, next := range []Stage{Provisioning, Readiness, Baseline, Injecting, Observing, Validating, Cleaning, Complete} {
		if err := s.Move(next, at); err != nil {
			t.Fatal(err)
		}
	}
	if s.Move(Cleaning, at) == nil {
		t.Fatal("reentered cleanup")
	}
}
func TestFailureCleanupAndClock(t *testing.T) {
	s := NewLifecycle(time.Now())
	if s.Move(Cleaning, time.Now().Add(-time.Hour)) == nil {
		t.Fatal("regression accepted")
	}
	if err := s.Move(Cleaning, time.Now()); err != nil {
		t.Fatal(err)
	}
}
func TestSustainedRecovery(t *testing.T) {
	at := time.Now()
	s := Stability{}
	if s.Update(at, true, time.Second) || s.Update(at.Add(time.Second/2), true, time.Second) {
		t.Fatal("early recovery")
	}
	if s.Update(at.Add(time.Second), false, time.Second) {
		t.Fatal("false recovered")
	}
	if s.Update(at.Add(2*time.Second), true, time.Second) {
		t.Fatal("did not reset")
	}
	if !s.Update(at.Add(3*time.Second), true, time.Second) {
		t.Fatal("not stable")
	}
}

func TestNetworkRestorationLifecycleBranches(t *testing.T) {
	for _, afterReconciliation := range []Stage{PostRecovery, Validating} {
		t.Run(string(afterReconciliation), func(t *testing.T) {
			at := time.Now()
			s := NewLifecycle(at)
			stages := []Stage{Provisioning, Readiness, Baseline, Injecting, Observing, Restoring, Reconciling, afterReconciliation}
			if afterReconciliation == PostRecovery {
				stages = append(stages, Validating)
			}
			stages = append(stages, Cleaning, Complete)
			for _, next := range stages {
				at = at.Add(time.Millisecond)
				if err := s.Move(next, at); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
	for _, failedAt := range []Stage{Restoring, Reconciling} {
		t.Run("cleanup from "+string(failedAt), func(t *testing.T) {
			at := time.Now()
			s := NewLifecycle(at)
			for _, next := range []Stage{Provisioning, Readiness, Baseline, Injecting, Observing, Restoring, Reconciling} {
				if err := s.Move(next, at); err != nil {
					t.Fatal(err)
				}
				if next == failedAt {
					break
				}
			}
			if err := s.Move(Cleaning, at); err != nil {
				t.Fatal("failed network phase cannot reach cleanup:", err)
			}
		})
	}
}

func TestReplicaRepairLifecycleBranchesAndFailureCleanup(t *testing.T) {
	for _, next := range []Stage{PostRecovery, Validating, Cleaning} {
		s := NewLifecycle(time.Now())
		for _, stage := range []Stage{Provisioning, Readiness, Baseline, Injecting, Observing, Repairing} {
			if err := s.Move(stage, time.Now()); err != nil {
				t.Fatal(err)
			}
		}
		if next != Cleaning {
			if err := s.Move(Reconciling, time.Now()); err != nil {
				t.Fatal(err)
			}
		}
		if err := s.Move(next, time.Now()); err != nil {
			t.Fatal(err)
		}
	}
}
