package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOfflineDiagnosisCLI(t *testing.T) {
	root := t.TempDir()
	input := filepath.Join(root, "input.json")
	if err := os.WriteFile(input, []byte(`{"schema_version":3,"samples":[]}`), 0644); err != nil {
		t.Fatal(err)
	}
	if got := execute([]string{"diagnose", input, "--output", filepath.Join(root, "output")}); got != 0 {
		t.Fatalf("diagnose exit=%d", got)
	}
	entries, err := os.ReadDir(filepath.Join(root, "output"))
	if err != nil || len(entries) != 1 {
		t.Fatal("missing diagnosis output")
	}
	for _, name := range []string{"diagnosis.json", "diagnosis.md"} {
		if _, err := os.Stat(filepath.Join(root, "output", entries[0].Name(), name)); err != nil {
			t.Fatal(err)
		}
	}
	if got := execute([]string{"diagnose", input, "--output", filepath.Join(root, "output")}); got != 1 {
		t.Fatal("existing diagnosis overwritten")
	}
	if got := execute([]string{"diagnose"}); got != 2 {
		t.Fatal("accepted missing input")
	}
	if got := execute([]string{"evaluate"}); got != 2 {
		t.Fatal("accepted missing dataset")
	}
}
