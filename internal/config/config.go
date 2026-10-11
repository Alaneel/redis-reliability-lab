// Package config strictly decodes and bounds the supported experiment definition.
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"go.yaml.in/yaml/v3"
)

const maxConfigBytes = 1 << 20

type Duration time.Duration

func (d Duration) Value() time.Duration { return time.Duration(d) }
func (d Duration) String() string       { return d.Value().String() }
func (d *Duration) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.ScalarNode || n.Tag != "!!str" {
		return errors.New("duration must be a string such as 5s")
	}
	v, err := time.ParseDuration(n.Value)
	if err != nil {
		return err
	}
	*d = Duration(v)
	return nil
}
func (d Duration) MarshalJSON() ([]byte, error) { return []byte(fmt.Sprintf("%q", d.String())), nil }
func (d *Duration) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	parsed, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	*d = Duration(parsed)
	return nil
}
func (d Duration) MarshalYAML() (any, error) { return d.String(), nil }

type Config struct {
	Version     int         `yaml:"version" json:"version"`
	Name        string      `yaml:"name" json:"name"`
	Environment Environment `yaml:"environment" json:"environment"`
	Workload    Workload    `yaml:"workload" json:"workload"`
	Experiment  Experiment  `yaml:"experiment" json:"experiment"`
	Timeouts    Timeouts    `yaml:"timeouts" json:"timeouts"`
}
type Environment struct {
	Persistence string   `yaml:"persistence" json:"persistence"`
	RedisImage  string   `yaml:"redis_image" json:"redis_image"`
	NodeTimeout Duration `yaml:"node_timeout" json:"node_timeout"`
}
type Workload struct {
	WritePolicy      string   `yaml:"write_policy" json:"write_policy"`
	WaitReplicas     int      `yaml:"wait_replicas" json:"wait_replicas"`
	WaitTimeout      Duration `yaml:"wait_timeout" json:"wait_timeout"`
	Workers          int      `yaml:"workers" json:"workers"`
	RatePerSecond    float64  `yaml:"rate_per_second" json:"rate_per_second"`
	MaxWrites        int      `yaml:"max_writes" json:"max_writes"`
	ReadPercent      int      `yaml:"read_percent" json:"read_percent"`
	ValueBytes       int      `yaml:"value_bytes" json:"value_bytes"`
	Seed             int64    `yaml:"seed" json:"seed"`
	OperationTimeout Duration `yaml:"operation_timeout" json:"operation_timeout"`
}
type Experiment struct {
	Fault                  string   `yaml:"fault" json:"fault"`
	Repair                 bool     `yaml:"repair" json:"repair"`
	DownDuration           Duration `yaml:"down_duration" json:"down_duration"`
	IsolationDuration      Duration `yaml:"isolation_duration" json:"isolation_duration"`
	Baseline               Duration `yaml:"baseline" json:"baseline"`
	PostRecovery           Duration `yaml:"post_recovery" json:"post_recovery"`
	SampleInterval         Duration `yaml:"sample_interval" json:"sample_interval"`
	StableFor              Duration `yaml:"stable_for" json:"stable_for"`
	MinimumBaselineSuccess float64  `yaml:"minimum_baseline_success" json:"minimum_baseline_success"`
}
type Timeouts struct {
	Provision      Duration `yaml:"provision" json:"provision"`
	Readiness      Duration `yaml:"readiness" json:"readiness"`
	Recovery       Duration `yaml:"recovery" json:"recovery"`
	Repair         Duration `yaml:"repair" json:"repair"`
	Restoration    Duration `yaml:"restoration" json:"restoration"`
	Reconciliation Duration `yaml:"reconciliation" json:"reconciliation"`
	Validation     Duration `yaml:"validation" json:"validation"`
	Cleanup        Duration `yaml:"cleanup" json:"cleanup"`
}

func Default() Config {
	return Config{Version: 1, Name: "primary-failure", Environment: Environment{Persistence: "memory", RedisImage: "redis:8.10.2-alpine", NodeTimeout: Duration(3 * time.Second)},
		Workload:   Workload{WritePolicy: "set", WaitReplicas: 1, WaitTimeout: Duration(100 * time.Millisecond), Workers: 8, RatePerSecond: 200, MaxWrites: 20000, ReadPercent: 50, ValueBytes: 128, Seed: 42, OperationTimeout: Duration(500 * time.Millisecond)},
		Experiment: Experiment{Fault: "primary-kill", DownDuration: Duration(3 * time.Second), IsolationDuration: Duration(8 * time.Second), Baseline: Duration(10 * time.Second), PostRecovery: Duration(5 * time.Second), SampleInterval: Duration(500 * time.Millisecond), StableFor: Duration(3 * time.Second), MinimumBaselineSuccess: 0.99},
		Timeouts:   Timeouts{Repair: Duration(30 * time.Second), Provision: Duration(2 * time.Minute), Readiness: Duration(time.Minute), Recovery: Duration(45 * time.Second), Restoration: Duration(30 * time.Second), Reconciliation: Duration(30 * time.Second), Validation: Duration(30 * time.Second), Cleanup: Duration(30 * time.Second)}}
}

func Load(path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()
	return Decode(f)
}
func Decode(r io.Reader) (Config, error) {
	c := Default()
	data, err := io.ReadAll(io.LimitReader(r, maxConfigBytes+1))
	if err != nil {
		return c, fmt.Errorf("read experiment: %w", err)
	}
	if len(data) > maxConfigBytes {
		return c, fmt.Errorf("experiment file must not exceed %d bytes", maxConfigBytes)
	}
	d := yaml.NewDecoder(bytes.NewReader(data))
	d.KnownFields(true)
	if err := d.Decode(&c); err != nil {
		return c, fmt.Errorf("decode experiment: %w", err)
	}
	var extra any
	if err := d.Decode(&extra); err != io.EOF {
		return c, errors.New("exactly one YAML document is required")
	}
	return c, c.Validate()
}

var imagePattern = regexp.MustCompile(`^redis:[0-9]+\.[0-9]+\.[0-9]+-alpine(@sha256:[a-f0-9]{64})?$`)
var namePattern = regexp.MustCompile(`^[a-z][a-z0-9-]{0,63}$`)

func (c Config) Validate() error {
	var errs []error
	check := func(ok bool, msg string) {
		if !ok {
			errs = append(errs, errors.New(msg))
		}
	}
	check(c.Version == 1, "version must be 1")
	check(namePattern.MatchString(c.Name), "name must be a lowercase slug of 1–64 characters")
	check(imagePattern.MatchString(c.Environment.RedisImage), "redis_image must pin an official redis:X.Y.Z-alpine tag (optional sha256 digest)")
	check(c.Experiment.Fault == "primary-kill" || c.Experiment.Fault == "primary-isolate" || c.Experiment.Fault == "replica-kill" || c.Experiment.Fault == "replica-isolate" || c.Experiment.Fault == "cluster-restart", "fault must be primary-kill, primary-isolate, replica-kill, replica-isolate or cluster-restart")
	check(c.Environment.Persistence == "memory" || c.Environment.Persistence == "aof-everysec" || c.Environment.Persistence == "aof-always", "persistence must be memory, aof-everysec or aof-always")
	check(!c.Experiment.Repair || c.Experiment.Fault == "primary-kill", "repair is supported only for primary-kill; replica faults always restore their target")
	if c.RetainData() {
		check(c.Workload.MaxWrites <= 20000 && c.Workload.ValueBytes <= 256, "retained-data experiments require max_writes <=20000 and value_bytes <=256 to bound disposable storage traffic")
		check(c.Timeouts.Recovery.Value() <= 2*time.Minute && c.Timeouts.Repair.Value() <= 2*time.Minute && c.Timeouts.Reconciliation.Value() <= 2*time.Minute, "retained-data observation and repair budgets must not exceed 2m")
	}
	check(c.Workload.WritePolicy == "set" || c.Workload.WritePolicy == "wait" || c.Workload.WritePolicy == "waitaof", "write_policy must be set, wait or waitaof")
	check(c.Workload.WaitReplicas == 1, "wait_replicas must be 1 for the fixed one-replica-per-primary topology")
	check(c.Workload.WaitTimeout.Value() >= time.Millisecond && c.Workload.WaitTimeout.Value() <= time.Second && c.Workload.WaitTimeout.Value()%time.Millisecond == 0, "wait_timeout must be an integral number of milliseconds between 1ms and 1s")
	if c.Workload.WritePolicy == "wait" || c.Workload.WritePolicy == "waitaof" {
		check(c.Workload.OperationTimeout.Value() >= c.Workload.WaitTimeout.Value()+10*time.Millisecond, "operation_timeout must exceed wait_timeout by at least 10ms")
	}
	if c.Workload.WritePolicy == "waitaof" {
		check(c.Environment.Persistence == "aof-everysec" || c.Environment.Persistence == "aof-always", "waitaof requires AOF persistence on every managed node")
		version := strings.Split(strings.TrimPrefix(c.Environment.RedisImage, "redis:"), ".")
		compatible := false
		if len(version) >= 2 {
			major, majorErr := strconv.Atoi(version[0])
			minor, minorErr := strconv.Atoi(version[1])
			compatible = majorErr == nil && minorErr == nil && (major > 7 || major == 7 && minor >= 2)
		}
		check(compatible, "waitaof requires Redis 7.2 or newer")
	}
	check(c.Workload.Workers >= 1 && c.Workload.Workers <= 128, "workers must be 1–128")
	check(c.Workload.RatePerSecond >= 1 && c.Workload.RatePerSecond <= 10000, "rate_per_second must be 1–10000")
	check(c.Workload.MaxWrites >= c.Workload.Workers && c.Workload.MaxWrites <= 100000, "max_writes must be at least workers and at most 100000")
	check(c.Workload.ReadPercent >= 0 && c.Workload.ReadPercent <= 99, "read_percent must be 0–99 (some writes are required)")
	check(c.Workload.ValueBytes >= 16 && c.Workload.ValueBytes <= 1024, "value_bytes must be 16–1024")
	check(c.Experiment.MinimumBaselineSuccess >= 0.9 && c.Experiment.MinimumBaselineSuccess <= 1, "minimum_baseline_success must be 0.9–1")
	for _, v := range []struct {
		name     string
		value    Duration
		min, max time.Duration
	}{
		{"node_timeout", c.Environment.NodeTimeout, time.Second, 30 * time.Second},
		{"operation_timeout", c.Workload.OperationTimeout, 10 * time.Millisecond, 5 * time.Second},
		{"baseline", c.Experiment.Baseline, time.Second, 5 * time.Minute},
		{"post_recovery", c.Experiment.PostRecovery, time.Second, time.Minute},
		{"sample_interval", c.Experiment.SampleInterval, 100 * time.Millisecond, 5 * time.Second},
		{"stable_for", c.Experiment.StableFor, time.Second, 30 * time.Second},
		{"provision", c.Timeouts.Provision, time.Second, 10 * time.Minute},
		{"readiness", c.Timeouts.Readiness, time.Second, 5 * time.Minute},
		{"recovery", c.Timeouts.Recovery, 100 * time.Millisecond, 5 * time.Minute},
		{"repair", c.Timeouts.Repair, 100 * time.Millisecond, 2 * time.Minute},
		{"down_duration", c.Experiment.DownDuration, time.Second, 30 * time.Second},
		{"isolation_duration", c.Experiment.IsolationDuration, time.Second, 2 * time.Minute},
		{"restoration", c.Timeouts.Restoration, time.Second, 2 * time.Minute},
		{"reconciliation", c.Timeouts.Reconciliation, 100 * time.Millisecond, 2 * time.Minute},
		{"validation", c.Timeouts.Validation, time.Second, 5 * time.Minute},
		{"cleanup", c.Timeouts.Cleanup, time.Second, 2 * time.Minute},
	} {
		check(v.value.Value() >= v.min && v.value.Value() <= v.max, fmt.Sprintf("%s must be between %s and %s", v.name, v.min, v.max))
	}
	check(c.Experiment.StableFor >= c.Experiment.SampleInterval, "stable_for must be at least sample_interval")
	return errors.Join(errs...)
}

// RetainData preserves only explicitly managed, disposable Redis data and membership files.
func (c Config) RetainData() bool {
	return c.Experiment.Fault == "cluster-restart" || c.Experiment.Fault == "replica-kill" || (c.Experiment.Fault == "primary-kill" && c.Experiment.Repair) || c.Environment.Persistence != "memory"
}
