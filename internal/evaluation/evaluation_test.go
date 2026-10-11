package evaluation

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMatricesAndUndefinedDenominators(t *testing.T) {
	m := Matrix{}
	m = update(m, true, true)
	m = update(m, true, false)
	m = update(m, false, true)
	m = update(m, false, false)
	if m.TP != 1 || m.FP != 1 || m.TN != 1 || m.FN != 1 || *m.Precision() != .5 || *m.Recall() != .5 {
		t.Fatal(m)
	}
	if (Matrix{}).Precision() != nil || (Matrix{}).Recall() != nil {
		t.Fatal("undefined ratio fabricated")
	}
}
func caseFor(raw, id, part, family string) Case {
	hash := sha256.Sum256([]byte(raw))
	return Case{ID: id, Partition: part, Family: family, Report: id + ".json", SHA256: hex.EncodeToString(hash[:]), Label: Label{Cause: "unknown", Categories: []string{}, Basis: []string{"test receipt"}}}
}
func TestManifestPartitionAndSchemaValidation(t *testing.T) {
	c := caseFor(`{"schema_version":3}`, "a", "development", "trial-a")
	d := Dataset{SchemaVersion: 1, RuleVersion: "rules-1", Cases: []Case{c}}
	if err := d.Validate(); err != nil {
		t.Fatal(err)
	}
	for _, mutate := range []func(*Dataset){
		func(d *Dataset) { d.Cases[0].Partition = "test" }, func(d *Dataset) { d.Cases[0].Report = "../secret" }, func(d *Dataset) { d.Cases[0].Label.Categories = []string{"invented"} }, func(d *Dataset) { d.Cases[0].Label.Categories = []string{"primary_unavailability"} }, func(d *Dataset) { d.Cases[0].SHA256 = "invalid" }, func(d *Dataset) { d.Cases[0].Label.Basis = nil },
	} {
		copyCase := c
		copyCase.Label = Label{Cause: "unknown", Categories: []string{}, Basis: []string{"receipt"}}
		bad := Dataset{SchemaVersion: 1, RuleVersion: "rules-1", Cases: []Case{copyCase}}
		mutate(&bad)
		if bad.Validate() == nil {
			t.Fatal("accepted invalid manifest")
		}
	}
	other := caseFor(`{"schema_version":2}`, "b", "heldout", "trial-a")
	d.Cases = append(d.Cases, other)
	if d.Validate() == nil {
		t.Fatal("trial family leaked across partition")
	}
	other.Family = "trial-b"
	other.SHA256 = c.SHA256
	d.Cases[1] = other
	if d.Validate() == nil {
		t.Fatal("duplicate input crossed partition")
	}
}
func TestActualEvaluationPreservesMissesAndAbstention(t *testing.T) {
	root := t.TempDir()
	healthy := `{"schema_version":3,"initial":{"views":[{"at":"2026-10-11T00:00:00Z","managed_node":{"id":"a"},"reachable":true,"coverage":{"assigned":16384}},{"at":"2026-10-11T00:00:00Z","managed_node":{"id":"b"},"reachable":true,"coverage":{"assigned":16384}}]}}`
	missing := `{"schema_version":3,"samples":[]}`
	corrupt := `{"schema_version":3,`
	cases := []Case{caseFor(healthy, "control", "heldout", "f1"), caseFor(missing, "insufficient", "heldout", "f2"), caseFor(corrupt, "malformed", "heldout", "f3")}
	cases[1].Label.Incident = true
	cases[1].Label.Categories = []string{"primary_unavailability"}
	cases[2].Label.Incident = true
	cases[2].Label.Categories = []string{"replica_unavailability"}
	for i, raw := range []string{healthy, missing, corrupt} {
		if err := os.WriteFile(filepath.Join(root, cases[i].Report), []byte(raw), 0644); err != nil {
			t.Fatal(err)
		}
	}
	d := Dataset{SchemaVersion: 1, RuleVersion: "rules-1", Cases: cases}
	b, _ := json.Marshal(d)
	manifest := filepath.Join(root, "manifest.json")
	os.WriteFile(manifest, b, 0644)
	r, err := Run(context.Background(), manifest, "heldout")
	if err != nil {
		t.Fatal(err)
	}
	if r.Cases != 3 || r.Detection.Matrix.TN != 1 || r.Detection.Matrix.FN != 2 || r.Abstentions != 3 || r.Results[2].Error == "" || r.DefinitiveClaims != 0 {
		t.Fatalf("%+v", r)
	}
	if r.Citations == 0 || r.CitationValidity == nil || *r.CitationValidity != 1 {
		t.Fatal("healthy timeline citations not validated")
	}
	if !strings.Contains(r.Markdown(), "malformed") {
		t.Fatal("failure hidden")
	}
	os.WriteFile(filepath.Join(root, cases[0].Report), []byte("changed"), 0644)
	if _, err := Run(context.Background(), manifest, "heldout"); err == nil {
		t.Fatal("ignored checksum failure")
	}
}
func TestPointerValidation(t *testing.T) {
	doc := map[string]any{"a": []any{map[string]any{"x/y": true}}}
	if !pointerExists(doc, "/a/0/x~1y") {
		t.Fatal("valid pointer rejected")
	}
	for _, p := range []string{"/a/1", "/absent", "not-pointer", "/a/0/x"} {
		if pointerExists(doc, p) {
			t.Fatalf("invalid %s", p)
		}
	}
}
