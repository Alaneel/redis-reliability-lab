package experiment

import (
	"fmt"
	"time"
)

type Stage string

const (
	Created      Stage = "created"
	Provisioning Stage = "provisioning"
	Readiness    Stage = "readiness"
	Baseline     Stage = "baseline"
	Quiescing    Stage = "quiescing"
	Injecting    Stage = "injecting"
	Observing    Stage = "observing"
	Restoring    Stage = "restoring"
	Repairing    Stage = "repairing"
	Reconciling  Stage = "reconciling"
	PostRecovery Stage = "post_recovery"
	Validating   Stage = "validating"
	Diagnosing   Stage = "diagnosing"
	Cleaning     Stage = "cleanup"
	Complete     Stage = "complete"
)

type Transition struct {
	Stage Stage     `json:"stage"`
	At    time.Time `json:"at"`
}
type Lifecycle struct {
	Current     Stage        `json:"current"`
	Transitions []Transition `json:"transitions"`
}

func NewLifecycle(at time.Time) Lifecycle { return Lifecycle{Created, []Transition{{Created, at}}} }
func (s *Lifecycle) Move(next Stage, at time.Time) error {
	allowed := map[Stage]Stage{Created: Provisioning, Provisioning: Readiness, Readiness: Baseline, Baseline: Injecting, Injecting: Observing, Observing: PostRecovery, PostRecovery: Validating, Validating: Cleaning, Cleaning: Complete}
	branch := (s.Current == Observing && (next == Validating || next == Restoring || next == Repairing)) ||
		(s.Current == Baseline && next == Quiescing) ||
		(s.Current == Quiescing && next == Injecting) ||
		(s.Current == Restoring && next == Reconciling) ||
		(s.Current == Repairing && next == Reconciling) ||
		(s.Current == Reconciling && (next == PostRecovery || next == Validating)) ||
		(s.Current == Validating && next == Diagnosing)
	if next != Cleaning && !branch && allowed[s.Current] != next {
		return fmt.Errorf("illegal stage transition %s -> %s", s.Current, next)
	}
	if next == Cleaning && (s.Current == Complete || s.Current == Cleaning) {
		return fmt.Errorf("cannot clean from %s", s.Current)
	}
	if len(s.Transitions) > 0 && at.Before(s.Transitions[len(s.Transitions)-1].At) {
		return fmt.Errorf("stage timestamp regressed")
	}
	s.Current = next
	s.Transitions = append(s.Transitions, Transition{next, at})
	return nil
}

// Stability requires a sequence of positive samples spanning the configured
// duration. A negative sample resets it. It is evidence at sampled times only.
type Stability struct {
	Since time.Time
	Last  time.Time
}

func (s *Stability) Update(at time.Time, ok bool, duration time.Duration) bool {
	if !s.Last.IsZero() && at.Before(s.Last) {
		s.Since = time.Time{}
		return false
	}
	s.Last = at
	if !ok {
		s.Since = time.Time{}
		return false
	}
	if s.Since.IsZero() {
		s.Since = at
	}
	return at.Sub(s.Since) >= duration
}
