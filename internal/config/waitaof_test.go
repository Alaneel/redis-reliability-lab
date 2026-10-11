package config

import (
	"testing"
	"time"
)

func TestWAITAOFRequiresPersistenceVersionAndBoundedTimeout(t *testing.T) {
	base := Default()
	base.Workload.WritePolicy, base.Environment.Persistence = "waitaof", "aof-always"
	if err := base.Validate(); err != nil {
		t.Fatal(err)
	}
	for _, mutate := range []func(*Config){
		func(c *Config) { c.Environment.Persistence = "memory" },
		func(c *Config) { c.Environment.RedisImage = "redis:7.0.15-alpine" },
		func(c *Config) { c.Environment.RedisImage = "redis:6.2.19-alpine" },
		func(c *Config) { c.Workload.WaitTimeout = 0 },
		func(c *Config) { c.Workload.WaitTimeout = Duration(1500 * time.Microsecond) },
		func(c *Config) { c.Workload.OperationTimeout = c.Workload.WaitTimeout },
		func(c *Config) { c.Workload.WaitReplicas = 0 },
	} {
		c := base
		mutate(&c)
		if c.Validate() == nil {
			t.Fatalf("invalid WAITAOF accepted: %#v", c)
		}
	}
	for _, image := range []string{"redis:7.2.0-alpine", "redis:8.10.2-alpine"} {
		c := base
		c.Environment.RedisImage = image
		if err := c.Validate(); err != nil {
			t.Fatal(err)
		}
	}
}
