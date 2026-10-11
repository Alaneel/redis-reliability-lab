package main

import (
	"flag"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestRepeatArgumentOrdersAndValidation(t *testing.T) {
	for _, test := range []struct {
		name      string
		args      []string
		wantError bool
	}{
		{"path first", []string{"experiment.yaml", "--runs", "7", "--output", "some directory"}, false},
		{"path last", []string{"--runs", "7", "--output", "some directory", "experiment.yaml"}, false},
		{"extra file", []string{"experiment.yaml", "other.yaml"}, true},
		{"missing file", []string{"--runs", "7"}, true},
		{"unknown flag", []string{"experiment.yaml", "--unknown"}, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			flags := flag.NewFlagSet("repeat", flag.ContinueOnError)
			flags.SetOutput(io.Discard)
			runs := flags.Int("runs", 5, "")
			output := flags.String("output", "reports", "")
			path, err := experimentArgument(flags, test.args)
			if test.wantError {
				if err == nil {
					t.Fatal("accepted invalid arguments")
				}
				return
			}
			if err != nil || path != "experiment.yaml" || *runs != 7 || *output != "some directory" {
				t.Fatalf("path=%q runs=%d output=%q error=%v", path, *runs, *output, err)
			}
		})
	}
}

func TestComparisonRejectsAllInvalidInputsBeforeCreatingEvidence(t *testing.T) {
	dir := t.TempDir()
	valid := filepath.Join(dir, "valid.yaml")
	invalid := filepath.Join(dir, "invalid.yaml")
	if err := os.WriteFile(valid, []byte("version: 1\nname: valid\n"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(invalid, []byte("version: 1\nunknown_option: true\n"), 0600); err != nil {
		t.Fatal(err)
	}
	output := filepath.Join(dir, "reports")
	for _, args := range [][]string{
		{"compare", "--output", output},
		{"compare", "--rounds", "0", "--output", output, valid, valid},
		{"compare", "--rounds", "51", "--output", output, valid, valid},
		{"compare", "--output", output, valid, invalid},
		{"compare", "--output", output, valid, filepath.Join(dir, "missing.yaml")},
		{"compare", "--unknown", "--output", output, valid, valid},
	} {
		if got := execute(args); got != 2 {
			t.Fatalf("invalid input exit=%d for %v", got, args)
		}
		if _, err := os.Stat(output); !os.IsNotExist(err) {
			t.Fatalf("invalid comparison created evidence or reached provisioning: %v", err)
		}
	}
}
