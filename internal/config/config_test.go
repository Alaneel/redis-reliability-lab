package config

import (
	"strings"
	"testing"
	"time"
)

func TestDefaultsAndOverrides(t *testing.T) {
	c, err := Decode(strings.NewReader("version: 1\nworkload:\n  workers: 4\n  operation_timeout: 200ms\n"))
	if err != nil {
		t.Fatal(err)
	}
	if c.Workload.Workers != 4 || c.Workload.OperationTimeout.Value() != 200*time.Millisecond || c.Environment.RedisImage != "redis:8.10.2-alpine" {
		t.Fatalf("bad merge: %+v", c)
	}
}
func TestRejectUnsafeOrInvalidConfig(t *testing.T) {
	for _, s := range []string{"", "version: 2", "unknown: true", "environment:\n  address: redis.example.com:6379", "environment:\n  redis_image: redis:latest", "experiment:\n  fault: exec", "workload:\n  workers: -1", "workload:\n  read_percent: 100", "workload:\n  rate_per_second: .nan", "workload:\n  operation_timeout: 500", "timeouts:\n  recovery: 0s", "version: 1\n---\nversion: 1", "version: 1\nversion: 1"} {
		t.Run(s, func(t *testing.T) {
			if _, err := Decode(strings.NewReader(s)); err == nil {
				t.Fatalf("accepted %q", s)
			}
		})
	}
}

func TestConfigByteLimitRejectsTruncation(t *testing.T) {
	header := "version: 1\nname: byte-limit\n"
	exactLimit := header + "#" + strings.Repeat("x", maxConfigBytes-len(header)-2) + "\n"
	c, err := Decode(strings.NewReader(exactLimit))
	if err != nil || c.Name != "byte-limit" {
		t.Fatalf("valid configuration at byte limit was rejected: name=%q err=%v", c.Name, err)
	}
	for name, input := range map[string]string{
		"otherwise valid": exactLimit + "\n",
		"second document": exactLimit + "---\nversion: 1\nname: second-document\n",
		"unknown field":   exactLimit + "unknown: true\n",
	} {
		t.Run(name, func(t *testing.T) {
			if _, err := Decode(strings.NewReader(input)); err == nil || !strings.Contains(err.Error(), "must not exceed") {
				t.Fatalf("oversized configuration was accepted or failed to report byte limit: %v", err)
			}
		})
	}
}

func TestFaultAndWritePolicyBounds(t *testing.T) {
	for _, s := range []string{
		"experiment:\n  isolation_duration: 0s", "timeouts:\n  restoration: 0s", "timeouts:\n  reconciliation: 0s",
		"workload:\n  write_policy: wait-aof", "workload:\n  wait_replicas: 2", "workload:\n  wait_timeout: 0ms",
		"workload:\n  wait_timeout: 1500us", "workload:\n  write_policy: wait\n  operation_timeout: 100ms\n  wait_timeout: 100ms",
	} {
		if _, err := Decode(strings.NewReader(s)); err == nil {
			t.Fatalf("accepted invalid fault or policy definition: %s", s)
		}
	}
	if _, err := Decode(strings.NewReader("experiment:\n  fault: primary-isolate\nworkload:\n  write_policy: wait\n  wait_timeout: 100ms\n  operation_timeout: 500ms")); err != nil {
		t.Fatal(err)
	}
}
