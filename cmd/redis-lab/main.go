package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/batch"
	"github.com/Alaneel/redis-reliability-lab/internal/comparison"
	"github.com/Alaneel/redis-reliability-lab/internal/config"
	"github.com/Alaneel/redis-reliability-lab/internal/environment"
	"github.com/Alaneel/redis-reliability-lab/internal/experiment"
)

var version = "dev"

func main() { os.Exit(execute(os.Args[1:])) }
func execute(args []string) int {
	if len(args) == 0 {
		usage()
		return 2
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	switch args[0] {
	case "compare":
		flags := flag.NewFlagSet("compare", flag.ContinueOnError)
		output := flags.String("output", "reports", "root directory for comparison and individual reports")
		rounds := flags.Int("rounds", 3, "counterbalanced rounds; one trial per condition per round")
		seed := flags.Int64("seed", 20261011, "execution order seed; workload seeds remain in each experiment")
		if err := flags.Parse(args[1:]); err != nil {
			return 2
		}
		if *rounds < 1 || flags.NArg() < 2 || flags.NArg() > 8 || *rounds > batch.MaxRuns/flags.NArg() {
			fmt.Fprintln(os.Stderr, "compare requires 2–8 experiment files and 1–100 total trials; flags precede files")
			return 2
		}
		configs := make([]config.Config, 0, flags.NArg())
		for _, path := range flags.Args() {
			c, err := config.Load(path)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return 2
			}
			configs = append(configs, c)
		}
		logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
		runner := experiment.Runner{Manager: environment.New(), Logger: logger}
		summary, err := comparison.Run(ctx, configs, *rounds, *seed, *output, runner.Run)
		if summary != nil {
			fmt.Printf("Comparison reports: %s\n", summary.Directory)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	case "repeat":
		flags := flag.NewFlagSet("repeat", flag.ContinueOnError)
		output := flags.String("output", "reports", "root directory for batch and individual reports")
		runs := flags.Int("runs", 5, "number of fresh sequential trials (1–100)")
		path, err := experimentArgument(flags, args[1:])
		if err != nil || *runs < 1 || *runs > 100 {
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Fprintln(os.Stderr, "runs must be 1–100")
			}
			return 2
		}
		c, err := config.Load(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
		runner := experiment.Runner{Manager: environment.New(), Logger: logger}
		summary, err := batch.Run(ctx, c, *runs, *output, runner.Run)
		if summary != nil {
			fmt.Printf("Batch: %s\nReports: %s\n", summary.BatchID, summary.Directory)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	case "run":
		flags := flag.NewFlagSet("run", flag.ContinueOnError)
		output := flags.String("output", "reports", "root directory for per-run JSON/Markdown reports")
		if err := flags.Parse(args[1:]); err != nil {
			return 2
		}
		if flags.NArg() != 1 {
			usage()
			return 2
		}
		c, err := config.Load(flags.Arg(0))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
		r, err := (experiment.Runner{Manager: environment.New(), Logger: logger}).Run(ctx, c, *output)
		if r != nil {
			abs, _ := filepath.Abs(filepath.Join(*output, r.RunID))
			fmt.Printf("Outcome: %s\nReports: %s\n", r.Outcome, abs)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		return 0
	case "validate":
		if len(args) != 2 {
			usage()
			return 2
		}
		c, err := config.Load(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 2
		}
		fmt.Printf("Valid experiment: %s (%s)\n", c.Name, c.Experiment.Fault)
		return 0
	case "cleanup":
		if len(args) != 2 {
			usage()
			return 2
		}
		cleanupCtx, cleanupCancel := context.WithTimeout(ctx, time.Minute)
		defer cleanupCancel()
		manager := environment.New()
		env, err := manager.Discover(cleanupCtx, args[1])
		if err == nil {
			err = manager.Cleanup(cleanupCtx, env)
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		fmt.Printf("Cleanup verified: %s\n", args[1])
		return 0
	case "version":
		if len(args) != 1 {
			usage()
			return 2
		}
		fmt.Println("redis-lab", version)
		return 0
	case "help", "-h", "--help":
		usage()
		return 0
	default:
		fmt.Fprintln(os.Stderr, errors.New("unknown command: "+args[0]))
		usage()
		return 2
	}
}

// Repeat accepts its one experiment path either before or after flags. This
// keeps the familiar single-run flag convention while supporting "repeat file
// --runs 5" without reordering flag values or accepting extra positional input.
func experimentArgument(flags *flag.FlagSet, args []string) (string, error) {
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		if err := flags.Parse(args[1:]); err != nil {
			return "", err
		}
		if flags.NArg() != 0 {
			return "", errors.New("exactly one experiment file is required")
		}
		return args[0], nil
	}
	if err := flags.Parse(args); err != nil {
		return "", err
	}
	if flags.NArg() != 1 {
		return "", errors.New("exactly one experiment file is required")
	}
	return flags.Arg(0), nil
}
func usage() {
	fmt.Fprintln(os.Stderr, "Usage:\n  redis-lab run [-output reports] experiments/primary-failure.yaml\n  redis-lab repeat experiments/primary-failure.yaml [--runs 5] [--output reports]\n  redis-lab compare [--rounds 3] [--seed 20261011] [--output reports] CONDITION.yaml CONDITION.yaml ...\n  redis-lab validate experiments/primary-failure.yaml\n  redis-lab cleanup RUN_ID\n  redis-lab version")
}
