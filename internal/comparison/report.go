package comparison

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/batch"
)

func (s Summary) Markdown() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "# Redis Reliability Lab — counterbalanced comparison\n\nComparison: `%s`  \nOutcome: **%s**  \nOrder seed: `%d`  \nStarted: %s  \nFinished: %s\n\n", s.ComparisonID, s.Outcome, s.OrderSeed, s.StartedAt.Format(time.RFC3339Nano), s.FinishedAt.Format(time.RFC3339Nano))
	fmt.Fprintf(&b, "Conditions: %d; rounds: %d; requested: %d; started: %d; callbacks completed: %d; unstarted: %d. Base condition order: %v.\n\n", len(s.Conditions), s.Rounds, s.RequestedTrials, s.StartedTrials, s.CompletedTrials, s.UnstartedTrials, s.BaseOrder)
	if s.StopReason != "" {
		fmt.Fprintf(&b, "Execution stopped: `%s`. Remaining entries retain their planned order and unstarted status.\n\n", s.StopReason)
	}
	fmt.Fprint(&b, "## Conditions\n\nStatistics stay within each condition. Failed and partial trials remain represented; censored or missing observations are not zero-duration recoveries. No pooled request latency quantiles or causal inference is presented.\n\n| Condition | Experiment | Workload seed | Recovered | Partial | Failed | Aborted | Unstarted | Summary |\n|---|---|---:|---:|---:|---:|---:|---:|---|\n")
	for _, c := range s.Conditions {
		o := c.Summary.Outcomes
		md := filepath.ToSlash(filepath.Join(filepath.Dir(c.SummaryPath), "summary.md"))
		fmt.Fprintf(&b, "| %s | %s | %d | %d | %d | %d | %d | %d | [JSON](%s), [Markdown](%s) |\n", c.ID, c.Config.Name, c.Config.Workload.Seed, o["recovered"], o["partial"], o["failed"], o["aborted"], o["unstarted"], c.SummaryPath, md)
	}
	fmt.Fprint(&b, "\n## Actual schedule\n\n| Position | Round | Round position | Condition | Run ID | Outcome | Evidence verified | Cleanup verified | Report |\n|---:|---:|---:|---|---|---|---|---|---|\n")
	for _, entry := range s.Schedule {
		link := "unstarted"
		if entry.ReportPath != "" {
			link = fmt.Sprintf("[JSON](%s)", entry.ReportPath)
			if entry.Trial.ReportMarkdownAvailable {
				link += fmt.Sprintf(", [Markdown](%s)", filepath.ToSlash(filepath.Join(filepath.Dir(entry.ReportPath), "report.md")))
			}
		} else if entry.Trial.Started {
			link = "unavailable"
		}
		fmt.Fprintf(&b, "| %d | %d | %d | %s | %s | %s | %t | %t | %s |\n", entry.Position, entry.Round, entry.RoundPosition, entry.ConditionID, entry.Trial.RunID, entry.Trial.Outcome, entry.Trial.EvidenceComplete, entry.Trial.Cleanup.Successful, link)
	}
	if len(s.Errors) > 0 {
		fmt.Fprint(&b, "\n## Execution errors\n\n")
		for _, err := range s.Errors {
			fmt.Fprintf(&b, "- %s\n", strings.ReplaceAll(err, "\n", "; "))
		}
	}
	fmt.Fprint(&b, "\n## Measurement methodology\n\n")
	for _, method := range s.Methodology {
		fmt.Fprintf(&b, "- %s\n", method)
	}
	return b.String()
}

// Write atomically replaces each completed file. The collection of comparison
// and condition summaries is not a filesystem transaction; errors are returned.
func Write(dir string, s Summary) error {
	for _, condition := range s.Conditions {
		if err := batch.Write(filepath.Join(dir, "conditions", condition.ID), condition.Summary); err != nil {
			return fmt.Errorf("write %s summary: %w", condition.ID, err)
		}
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("encode comparison: %w", err)
	}
	if err := atomicWrite(filepath.Join(dir, "comparison.json"), append(data, '\n')); err != nil {
		return err
	}
	return atomicWrite(filepath.Join(dir, "comparison.md"), []byte(s.Markdown()))
}

func atomicWrite(path string, data []byte) error {
	f, err := os.CreateTemp(filepath.Dir(path), ".comparison-*")
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	if _, err := f.Write(data); err != nil {
		f.Close()
		return err
	}
	if err := f.Sync(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(f.Name(), path)
}
