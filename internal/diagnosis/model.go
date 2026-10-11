// Package diagnosis analyzes allowlisted observations offline. It has no Redis,
// Docker, orchestrator, configuration or ground-truth dependency.
package diagnosis

import (
	"encoding/json"
	"fmt"
	"time"
)

const SchemaVersion = 1
const EngineVersion = "rules-2"
const MaxEvidence = 100000
const MaxInputBytes = 64 << 20

type Status string

const (
	Valid         Status = "valid"
	Missing       Status = "missing_from_schema"
	Unavailable   Status = "explicitly_unavailable"
	Failed        Status = "collection_failed"
	NotApplicable Status = "not_applicable"
	Invalid       Status = "invalid"
)

type Evidence struct {
	ID              string          `json:"id"`
	At              time.Time       `json:"at"`
	Node            string          `json:"node,omitempty"`
	Source          string          `json:"source"`
	Signal          string          `json:"signal"`
	Value           json.RawMessage `json:"value,omitempty"`
	Status          Status          `json:"status"`
	Reference       string          `json:"reference"`
	Quality         string          `json:"quality"`
	CollectionError string          `json:"collection_error,omitempty"`
}

type Incident struct {
	SchemaVersion int        `json:"schema_version"`
	ID            string     `json:"incident_id"`
	SourceSHA256  string     `json:"source_sha256"`
	InputSchema   int        `json:"input_schema_version"`
	Evidence      []Evidence `json:"evidence"`
	Limitations   []string   `json:"limitations"`
}

type Finding struct {
	Category     string   `json:"category"`
	Description  string   `json:"description"`
	Nodes        []string `json:"affected_nodes"`
	Support      []string `json:"supporting_evidence"`
	Against      []string `json:"contradicting_evidence"`
	Missing      []string `json:"missing_evidence"`
	Status       string   `json:"assessment"`
	Rule         string   `json:"rule"`
	Alternatives []string `json:"alternative_explanations"`
	Verification []string `json:"verification_steps"`
}

type Event struct {
	Kind     string     `json:"kind"`
	At       time.Time  `json:"first_observed_at"`
	After    *time.Time `json:"last_previous_observation_at,omitempty"`
	Timing   string     `json:"timing"`
	Evidence []string   `json:"evidence"`
}

type Recommendation struct {
	Objective       string   `json:"objective"`
	Category        string   `json:"category"`
	Evidence        []string `json:"required_evidence"`
	Preconditions   []string `json:"preconditions"`
	Risks           []string `json:"risks"`
	ExpectedEffects string   `json:"expected_effects"`
	Verification    []string `json:"verification_criteria"`
	HumanRequired   bool     `json:"human_intervention_required"`
	Execution       string   `json:"execution"`
}

type Assessment struct {
	SchemaVersion   int              `json:"schema_version"`
	EngineVersion   string           `json:"engine_version"`
	Incident        Incident         `json:"incident"`
	Detection       string           `json:"detection"` // incident, healthy_observed, insufficient_evidence
	RootCause       string           `json:"root_cause"`
	Findings        []Finding        `json:"hypotheses"`
	Timeline        []Event          `json:"timeline"`
	Recommendations []Recommendation `json:"recommendations"`
	Uncertainty     []string         `json:"unresolved_uncertainty"`
	Recovery        []string         `json:"recovery_state"`
}

func value(e Evidence, out any) bool    { return e.Status == Valid && json.Unmarshal(e.Value, out) == nil }
func boolValue(e Evidence) (bool, bool) { var b bool; ok := value(e, &b); return b, ok }
func number(e Evidence) (float64, bool) { var n float64; ok := value(e, &n); return n, ok }
func stringValue(e Evidence) string     { var s string; value(e, &s); return s }

// Validate checks the citation graph, safety contract and provenance before output.
func (a Assessment) Validate() error {
	if a.SchemaVersion != SchemaVersion || a.Incident.SchemaVersion != SchemaVersion {
		return fmt.Errorf("unsupported assessment schema")
	}
	ids := map[string]Evidence{}
	for _, e := range a.Incident.Evidence {
		if e.ID == "" || e.Reference == "" || e.Source == "" {
			return fmt.Errorf("incomplete provenance")
		}
		if _, ok := ids[e.ID]; ok {
			return fmt.Errorf("duplicate evidence %s", e.ID)
		}
		ids[e.ID] = e
	}
	check := func(refs []string) error {
		for _, id := range refs {
			if _, ok := ids[id]; !ok {
				return fmt.Errorf("unknown evidence %s", id)
			}
		}
		return nil
	}
	for _, f := range a.Findings {
		if f.Rule == "" || len(f.Support) == 0 {
			return fmt.Errorf("unsupported finding")
		}
		if err := check(append(append([]string{}, f.Support...), f.Against...)); err != nil {
			return err
		}
	}
	for _, e := range a.Timeline {
		if len(e.Evidence) == 0 {
			return fmt.Errorf("uncited timeline event")
		}
		if err := check(e.Evidence); err != nil {
			return err
		}
	}
	for _, r := range a.Recommendations {
		if !r.HumanRequired || r.Execution != "read_only_plan" || len(r.Preconditions) == 0 || len(r.Risks) == 0 || len(r.Verification) == 0 {
			return fmt.Errorf("unsafe recommendation")
		}
		if err := check(r.Evidence); err != nil {
			return err
		}
	}
	return nil
}
