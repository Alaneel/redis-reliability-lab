// Package evaluation owns laboratory labels. The diagnosis package cannot
// import this package or receive a Case: only the observation report bytes cross
// the inference boundary. Dataset identifiers and partitions never reach rules.
package evaluation

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/diagnosis"
)

var Categories = []string{"primary_unavailability", "replica_unavailability", "redundancy_degradation", "acknowledgement_degradation", "client_recovery_lag", "data_validation_failure", "persistence_recovery", "validation_incomplete", "client_disruption", "node_unavailability", "acknowledgement_command_failure", "topology_disagreement", "client_latency_change"}

type Label struct {
	Incident   bool       `json:"incident"`
	Categories []string   `json:"categories"`
	Cause      string     `json:"injected_cause"`
	Target     string     `json:"target_node,omitempty"`
	StartedAt  *time.Time `json:"fault_start,omitempty"`
	EndedAt    *time.Time `json:"fault_end,omitempty"`
	Basis      []string   `json:"label_basis"`
}
type Case struct {
	ID        string `json:"case_id"`
	Partition string `json:"partition"`
	Family    string `json:"trial_family"`
	Report    string `json:"report"`
	SHA256    string `json:"sha256"`
	Label     Label  `json:"ground_truth"`
}
type Dataset struct {
	SchemaVersion int      `json:"schema_version"`
	RuleVersion   string   `json:"frozen_rule_version"`
	Cases         []Case   `json:"cases"`
	Methodology   []string `json:"methodology"`
}
type Matrix struct {
	TP int `json:"tp"`
	FP int `json:"fp"`
	TN int `json:"tn"`
	FN int `json:"fn"`
}

func (m Matrix) Precision() *float64 { return ratio(m.TP, m.TP+m.FP) }
func (m Matrix) Recall() *float64    { return ratio(m.TP, m.TP+m.FN) }
func ratio(n, d int) *float64 {
	if d == 0 {
		return nil
	}
	r := float64(n) / float64(d)
	return &r
}

type CategoryResult struct {
	Matrix    Matrix   `json:"confusion"`
	Precision *float64 `json:"precision"`
	Recall    *float64 `json:"recall"`
}
type CaseResult struct {
	ID        string   `json:"case_id"`
	Expected  []string `json:"expected_categories"`
	Predicted []string `json:"predicted_categories"`
	Detection string   `json:"detection"`
	RootCause string   `json:"root_cause"`
	LatencyMS float64  `json:"execution_latency_ms"`
	Error     string   `json:"error,omitempty"`
}
type Result struct {
	EngineVersion                 string                    `json:"engine_version"`
	DatasetRuleVersion            string                    `json:"dataset_rule_version"`
	SchemaVersion                 int                       `json:"schema_version"`
	Partition                     string                    `json:"partition"`
	DatasetSHA256                 string                    `json:"dataset_sha256"`
	Cases                         int                       `json:"cases"`
	Detection                     CategoryResult            `json:"incident_detection"`
	Categories                    map[string]CategoryResult `json:"categories"`
	Micro                         CategoryResult            `json:"multilabel_micro"`
	Abstentions                   int                       `json:"root_cause_abstentions"`
	AbstentionRate                *float64                  `json:"root_cause_abstention_rate"`
	IncorrectDefinitiveClaims     int                       `json:"incorrect_definitive_root_cause_claims"`
	DefinitiveClaims              int                       `json:"definitive_claims"`
	Citations                     int                       `json:"citations"`
	ValidCitations                int                       `json:"valid_citations"`
	CitationValidity              *float64                  `json:"citation_validity"`
	Recommendations               int                       `json:"recommendations"`
	UnsupportedRecommendations    int                       `json:"unsupported_recommendations"`
	UnsupportedRecommendationRate *float64                  `json:"unsupported_recommendation_rate"`
	Results                       []CaseResult              `json:"case_results"`
	Limitations                   []string                  `json:"limitations"`
}

func Load(path string) (Dataset, []byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return Dataset{}, nil, err
	}
	defer f.Close()
	b, err := io.ReadAll(io.LimitReader(f, 4<<20+1))
	if err != nil {
		return Dataset{}, nil, err
	}
	if len(b) > 4<<20 {
		return Dataset{}, nil, fmt.Errorf("dataset manifest too large")
	}
	var d Dataset
	dec := json.NewDecoder(strings.NewReader(string(b)))
	dec.DisallowUnknownFields()
	if err = dec.Decode(&d); err != nil {
		return d, nil, err
	}
	var trailing any
	if dec.Decode(&trailing) != io.EOF {
		return d, nil, fmt.Errorf("trailing manifest data")
	}
	return d, b, d.Validate()
}
func (d Dataset) Validate() error {
	if d.SchemaVersion != 1 || (d.RuleVersion != "rules-1" && d.RuleVersion != "rules-2") || len(d.Cases) == 0 || len(d.Cases) > 200 {
		return fmt.Errorf("invalid dataset schema, rule version or size")
	}
	ids := map[string]bool{}
	hashes := map[string]bool{}
	families := map[string]string{}
	for _, c := range d.Cases {
		if c.ID == "" || ids[c.ID] || c.Family == "" || len(c.Label.Basis) == 0 || c.Label.Cause == "" {
			return fmt.Errorf("invalid or duplicate case")
		}
		ids[c.ID] = true
		if c.Partition != "development" && c.Partition != "heldout" {
			return fmt.Errorf("invalid partition")
		}
		if old, ok := families[c.Family]; ok && old != c.Partition {
			return fmt.Errorf("trial family crosses partitions")
		}
		families[c.Family] = c.Partition
		hash, err := hex.DecodeString(c.SHA256)
		if err != nil || len(hash) != 32 || hashes[c.SHA256] {
			return fmt.Errorf("invalid or duplicate report checksum")
		}
		hashes[c.SHA256] = true
		if c.Report == "" || filepath.IsAbs(c.Report) || filepath.Clean(c.Report) != c.Report || c.Report == ".." || strings.HasPrefix(c.Report, ".."+string(filepath.Separator)) {
			return fmt.Errorf("report path escapes dataset")
		}
		seen := map[string]bool{}
		for _, k := range c.Label.Categories {
			if seen[k] || !has(Categories, k) {
				return fmt.Errorf("unknown or duplicate category %s", k)
			}
			seen[k] = true
		}
		if !c.Label.Incident && len(c.Label.Categories) > 0 {
			return fmt.Errorf("healthy label with incident categories")
		}
	}
	return nil
}
func has(s []string, v string) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}
func update(m Matrix, want, got bool) Matrix {
	if want && got {
		m.TP++
	} else if want {
		m.FN++
	} else if got {
		m.FP++
	} else {
		m.TN++
	}
	return m
}
func category(m Matrix) CategoryResult { return CategoryResult{m, m.Precision(), m.Recall()} }

func Run(ctx context.Context, manifest, partition string) (Result, error) {
	if partition != "development" && partition != "heldout" {
		return Result{}, fmt.Errorf("partition must be development or heldout")
	}
	d, b, err := Load(manifest)
	if err != nil {
		return Result{}, err
	}
	sum := sha256.Sum256(b)
	r := Result{EngineVersion: diagnosis.EngineVersion, DatasetRuleVersion: d.RuleVersion, SchemaVersion: 1, Partition: partition, DatasetSHA256: hex.EncodeToString(sum[:]), Categories: map[string]CategoryResult{}, Results: []CaseResult{}, Limitations: []string{
		"Small single-host laboratory dataset; no generalization or production claims.",
		"Category matrices treat every selected case as one multilabel decision per category; overlap is allowed. Micro metrics sum category TP/FP/FN. Undefined denominators are null.",
		"Input and analysis failures count as undetected incidents and negative category predictions, and remain visible in case results.",
		"Abstention denominator is every selected case, including controls; definitive-claim counts are separate.",
		"Citation validity checks JSON-pointer existence in original bytes, not the causal truth of an interpretation.",
		"Unsupported recommendations check cited evidence availability and the read-only safety contract, not operator effectiveness.",
		"Execution latency includes input normalization and deterministic analysis, excluding file reading and report writing; timings are not deterministic output.",
	}}
	matrices := map[string]Matrix{}
	detect := Matrix{}
	for _, c := range d.Cases {
		if c.Partition != partition {
			continue
		}
		if err := ctx.Err(); err != nil {
			return r, err
		}
		r.Cases++
		cr := CaseResult{ID: c.ID, Expected: c.Label.Categories, Predicted: []string{}, RootCause: "abstained", Detection: "insufficient_evidence"}
		root := filepath.Dir(manifest)
		path, err := filepath.EvalSymlinks(filepath.Join(root, c.Report))
		if err != nil {
			return r, err
		}
		absRoot, err := filepath.EvalSymlinks(root)
		if err != nil {
			return r, err
		}
		rel, err := filepath.Rel(absRoot, path)
		if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			return r, fmt.Errorf("report symlink escapes dataset")
		}
		f, err := os.Open(path)
		if err != nil {
			return r, err
		}
		raw, err := io.ReadAll(io.LimitReader(f, diagnosis.MaxInputBytes+1))
		f.Close()
		if err != nil {
			return r, err
		}
		actual := sha256.Sum256(raw)
		if hex.EncodeToString(actual[:]) != c.SHA256 {
			return r, fmt.Errorf("report checksum mismatch: %s", c.ID)
		}
		started := time.Now()
		in, analysisErr := diagnosis.Normalize(ctx, strings.NewReader(string(raw)))
		var a diagnosis.Assessment
		if analysisErr == nil {
			a, analysisErr = diagnosis.Analyze(ctx, in)
		}
		cr.LatencyMS = float64(time.Since(started)) / float64(time.Millisecond)
		if analysisErr != nil {
			cr.Error = analysisErr.Error()
		} else {
			cr.Detection = a.Detection
			cr.RootCause = a.RootCause
			for _, finding := range a.Findings {
				cr.Predicted = append(cr.Predicted, finding.Category)
			}
			sort.Strings(cr.Predicted)
			var doc any
			json.Unmarshal(raw, &doc)
			evidence := map[string]diagnosis.Evidence{}
			for _, e := range in.Evidence {
				evidence[e.ID] = e
			}
			refs := []string{}
			for _, finding := range a.Findings {
				refs = append(refs, finding.Support...)
				refs = append(refs, finding.Against...)
			}
			for _, event := range a.Timeline {
				refs = append(refs, event.Evidence...)
			}
			for _, plan := range a.Recommendations {
				r.Recommendations++
				supported := len(plan.Evidence) > 0 && plan.Execution == "read_only_plan" && plan.HumanRequired && len(plan.Preconditions) > 0 && len(plan.Risks) > 0 && len(plan.Verification) > 0
				for _, id := range plan.Evidence {
					e, ok := evidence[id]
					if !ok || e.Status != diagnosis.Valid || !pointerExists(doc, e.Reference) {
						supported = false
					}
				}
				if !supported {
					r.UnsupportedRecommendations++
				}
				refs = append(refs, plan.Evidence...)
			}
			for _, id := range refs {
				r.Citations++
				e, ok := evidence[id]
				if ok && pointerExists(doc, e.Reference) {
					r.ValidCitations++
				}
			}
		}
		if cr.RootCause == "abstained" {
			r.Abstentions++
		} else {
			r.DefinitiveClaims++
			if cr.RootCause != c.Label.Cause {
				r.IncorrectDefinitiveClaims++
			}
		}
		detect = update(detect, c.Label.Incident, cr.Detection == "incident")
		for _, k := range Categories {
			matrices[k] = update(matrices[k], has(c.Label.Categories, k), has(cr.Predicted, k))
		}
		r.Results = append(r.Results, cr)
	}
	if r.Cases == 0 {
		return r, fmt.Errorf("partition has no cases")
	}
	r.Detection = category(detect)
	micro := Matrix{}
	for _, k := range Categories {
		m := matrices[k]
		r.Categories[k] = category(m)
		micro.TP += m.TP
		micro.FP += m.FP
		micro.FN += m.FN
		micro.TN += m.TN
	}
	r.Micro = category(micro)
	r.AbstentionRate = ratio(r.Abstentions, r.Cases)
	r.CitationValidity = ratio(r.ValidCitations, r.Citations)
	r.UnsupportedRecommendationRate = ratio(r.UnsupportedRecommendations, r.Recommendations)
	return r, nil
}

func pointerExists(doc any, pointer string) bool {
	if pointer == "" {
		return true
	}
	if !strings.HasPrefix(pointer, "/") {
		return false
	}
	cur := doc
	for _, part := range strings.Split(pointer[1:], "/") {
		part = strings.ReplaceAll(strings.ReplaceAll(part, "~1", "/"), "~0", "~")
		switch v := cur.(type) {
		case map[string]any:
			var ok bool
			cur, ok = v[part]
			if !ok {
				return false
			}
		case []any:
			i, err := strconv.Atoi(part)
			if err != nil || i < 0 || i >= len(v) {
				return false
			}
			cur = v[i]
		default:
			return false
		}
	}
	return true
}

func (r Result) Markdown() string {
	var b strings.Builder
	fmt.Fprintf(&b, "# Diagnostic evaluation — %s\n\nCases: %d. Dataset SHA-256: `%s`.\n\nIncident detection: TP %d, FP %d, TN %d, FN %d. Precision %s, recall %s.\n\n| Category | TP | FP | TN | FN | Precision | Recall |\n|---|---:|---:|---:|---:|---:|---:|\n", r.Partition, r.Cases, r.DatasetSHA256, r.Detection.Matrix.TP, r.Detection.Matrix.FP, r.Detection.Matrix.TN, r.Detection.Matrix.FN, rate(r.Detection.Precision), rate(r.Detection.Recall))
	for _, k := range Categories {
		c := r.Categories[k]
		fmt.Fprintf(&b, "| %s | %d | %d | %d | %d | %s | %s |\n", k, c.Matrix.TP, c.Matrix.FP, c.Matrix.TN, c.Matrix.FN, rate(c.Precision), rate(c.Recall))
	}
	fmt.Fprintf(&b, "\nMicro precision %s, recall %s. Root-cause abstentions %d/%d; definitive claims %d, incorrect %d. Citations valid %d/%d. Unsupported recommendations %d/%d.\n\n## Individual outcomes\n\n| Case | Expected | Predicted | Detection | ms | Error |\n|---|---|---|---|---:|---|\n", rate(r.Micro.Precision), rate(r.Micro.Recall), r.Abstentions, r.Cases, r.DefinitiveClaims, r.IncorrectDefinitiveClaims, r.ValidCitations, r.Citations, r.UnsupportedRecommendations, r.Recommendations)
	for _, c := range r.Results {
		fmt.Fprintf(&b, "| %s | %s | %s | %s | %.3f | %s |\n", c.ID, strings.Join(c.Expected, ", "), strings.Join(c.Predicted, ", "), c.Detection, c.LatencyMS, c.Error)
	}
	b.WriteString("\n## Limitations\n\n")
	for _, s := range r.Limitations {
		fmt.Fprintf(&b, "- %s\n", s)
	}
	return b.String()
}
func rate(v *float64) string {
	if v == nil {
		return "undefined"
	}
	return fmt.Sprintf("%.6f", *v)
}
func Write(dir string, r Result) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	if err = os.WriteFile(filepath.Join(dir, "evaluation.json"), append(b, '\n'), 0644); err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(dir, "evaluation.md"), []byte(r.Markdown()), 0644)
}
