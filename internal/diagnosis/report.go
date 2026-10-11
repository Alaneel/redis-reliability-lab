package diagnosis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (a Assessment) Markdown() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "# Incident diagnosis\n\nIncident: `%s`  \nDetection: **%s**  \nUnique root cause: **%s**  \nEngine: `%s`\n\n", a.Incident.ID, a.Detection, a.RootCause, a.EngineVersion)
	fmt.Fprint(&b, "## Observed symptoms and diagnostic hypotheses\n\n")
	for _, f := range a.Findings {
		fmt.Fprintf(&b, "### %s\n\n%s\n\nAssessment: %s. Rule: `%s`. Nodes: %s.\n\nSupporting evidence: %s.\n\nContrary or recovery observations: %s.\n\nMissing evidence: %s.\n\nAlternatives: %s.\n\n", f.Category, f.Description, f.Status, f.Rule, strings.Join(f.Nodes, ", "), strings.Join(f.Support, ", "), strings.Join(f.Against, ", "), strings.Join(f.Missing, "; "), strings.Join(f.Alternatives, "; "))
	}
	fmt.Fprintln(&b, "## Timeline\n\n| First observed time | Event | Evidence | Timing |\n|---|---|---|---|")
	for _, e := range a.Timeline {
		fmt.Fprintf(&b, "| %s | %s | %s | %s |\n", e.At.Format("2006-01-02T15:04:05.999999999Z07:00"), e.Kind, strings.Join(e.Evidence, ", "), e.Timing)
	}
	fmt.Fprint(&b, "\n## Recovery and uncertainty\n\n")
	for _, s := range append(append([]string{}, a.Recovery...), a.Uncertainty...) {
		fmt.Fprintf(&b, "- %s\n", s)
	}
	fmt.Fprint(&b, "\n## Read-only remediation plans\n\n")
	for _, r := range a.Recommendations {
		fmt.Fprintf(&b, "- **%s**: %s. Evidence: %s. Preconditions: %s. Risks: %s. Expected effect: %s. Verify: %s. Human approval required; no execution.\n", r.Category, r.Objective, strings.Join(r.Evidence, ", "), strings.Join(r.Preconditions, "; "), strings.Join(r.Risks, "; "), r.ExpectedEffects, strings.Join(r.Verification, "; "))
	}
	fmt.Fprintf(&b, "\n## Evidence provenance\n\nOriginal bytes SHA-256: `%s`. Input report schema: %d. References are JSON pointers in the original report.\n\n| ID | Time | Node | Source | Signal | Status | Value | Reference |\n|---|---|---|---|---|---|---|---|\n", a.Incident.SourceSHA256, a.Incident.InputSchema)
	for _, e := range a.Incident.Evidence {
		v := strings.ReplaceAll(string(e.Value), "|", "\\|")
		fmt.Fprintf(&b, "| %s | %s | %s | %s | %s | %s | `%s` | %s |\n", e.ID, e.At.Format("15:04:05.999999999"), e.Node, e.Source, e.Signal, e.Status, v, e.Reference)
	}
	fmt.Fprint(&b, "\n## Analysis limitations\n\n")
	for _, s := range a.Incident.Limitations {
		fmt.Fprintf(&b, "- %s\n", s)
	}
	return b.String()
}

// Write uses exclusive creation to avoid overwriting source or prior diagnoses.
func Write(dir string, a Assessment) error {
	if err := a.Validate(); err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	for _, file := range []struct {
		name string
		data []byte
	}{{"diagnosis.json", append(data, '\n')}, {"diagnosis.md", []byte(a.Markdown())}} {
		f, err := os.OpenFile(filepath.Join(dir, file.name), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
		if err != nil {
			return err
		}
		_, err = f.Write(file.data)
		closeErr := f.Close()
		if err != nil {
			return err
		}
		if closeErr != nil {
			return closeErr
		}
	}
	return nil
}
