package environment

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestRetainedProcessCapturesActualKeyspaceFilesAndRewriteConfiguration(t *testing.T) {
	for _, mode := range []string{"memory", "aof-everysec", "aof-always"} {
		t.Run(mode, func(t *testing.T) {
			manager, fake := newFake()
			env := provisionRetained(t, manager, mode)
			original := manager.execute
			manager.execute = func(ctx context.Context, args ...string) ([]byte, error) {
				out, err := original(ctx, args...)
				if strings.HasSuffix(strings.Join(args, " "), "info keyspace") {
					return []byte("# Keyspace\r\ndb0:keys=17,expires=0,avg_ttl=0,subexpiry=0\r\n"), err
				}
				return out, err
			}
			fake.calls = nil
			evidence, err := manager.ProbeProcess(context.Background(), env, env.Nodes[0])
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(evidence.RawKeyspace, "keys=17") || evidence.DataSizeKiB != 8 || !strings.Contains(evidence.RawDataFiles, "/data/nodes.conf\n") || evidence.ContainerRestartPolicy != "no" {
				t.Fatalf("actual storage/keyspace observations lost: %#v", evidence)
			}
			if strings.Contains(evidence.RawDataFiles, "dump.rdb") {
				t.Fatalf("unexpected RDB snapshot in configured diskless fixture: %q", evidence.RawDataFiles)
			}
			if mode != "memory" {
				for _, suffix := range []string{"appendonly.aof.manifest\n", "appendonly.aof.1.base.rdb\n", "appendonly.aof.1.incr.aof\n"} {
					if !strings.Contains(evidence.RawDataFiles, suffix) {
						t.Fatalf("AOF file evidence absent: %q", evidence.RawDataFiles)
					}
				}
			} else if strings.Contains(evidence.RawDataFiles, "appendonly") {
				t.Fatal("memory fixture reported AOF files")
			}
			preamble := "no"
			if mode != "memory" {
				preamble = "yes"
			}
			for _, field := range []string{"aof-use-rdb-preamble\n" + preamble, "rdb-del-sync-files\nyes", "no-appendfsync-on-rewrite\nno", "aof-load-truncated\nyes"} {
				if !strings.Contains(evidence.RawPersistenceConfig, field) {
					t.Fatalf("runtime policy absent: %s", field)
				}
			}
			if !strings.Contains(env.Metadata["storage_bound"], "full synchronization") || !strings.Contains(env.Metadata["aof_load_truncated"], "default yes") {
				t.Fatal("rewrite/truncated-tail limitations absent")
			}
			if countOperation(fake, "info keyspace") != 1 || countOperation(fake, "find /data -maxdepth 2 -type f") != 1 || countOperation(fake, "du -sk /data") != 1 {
				t.Fatal("storage observations were missing or duplicated")
			}
			for _, call := range fake.calls {
				text := strings.Join(call, " ")
				for _, forbidden := range []string{"WAITAOF", "waitaof", "bgsave", "bgrewriteaof", " save", "sh -c"} {
					if strings.Contains(text, forbidden) && !strings.Contains(text, "config get") {
						t.Fatalf("observation changed persistence timing: %q", call)
					}
				}
			}
		})
	}
}

func TestStoppedProcessEvidenceRecordsInspectedNoRestartPolicy(t *testing.T) {
	manager, fake := newFake()
	env := provisionRetained(t, manager, "memory")
	node := env.Nodes[0]
	if err := manager.Kill(context.Background(), env, node); err != nil {
		t.Fatal(err)
	}
	fake.calls = nil
	evidence, err := manager.ProbeProcess(context.Background(), env, node)
	if !errors.Is(err, ErrProcessStopped) || evidence.ContainerRunning || evidence.ContainerRestartPolicy != "no" {
		t.Fatalf("stopped barrier lacks actual restart-policy evidence: %#v %v", evidence, err)
	}
	if hasOperation(fake, "exec") {
		t.Fatal("stopped process probe executed administrative commands")
	}
}

func TestRetainedArgumentsConfigureCompatiblePreambleAndSyncFileRemoval(t *testing.T) {
	for _, mode := range []string{"memory", "aof-everysec", "aof-always"} {
		t.Run(mode, func(t *testing.T) {
			manager, fake := newFake()
			provisionRetained(t, manager, mode)
			preamble := "no"
			if mode != "memory" {
				preamble = "yes"
			}
			for _, field := range []string{"--aof-use-rdb-preamble " + preamble, "--rdb-del-sync-files yes", "--no-appendfsync-on-rewrite no"} {
				if countOperation(fake, field) != nodeCount {
					t.Fatalf("policy was not explicit on every retained node: %s", field)
				}
			}
			if hasOperation(fake, "--aof-load-truncated") {
				t.Fatal("default truncated-load policy unexpectedly overridden")
			}
		})
	}
}

func TestDataFileListingRejectsForeignMalformedAndTruncatedOutput(t *testing.T) {
	for _, raw := range []string{"", "/data/nodes.conf\n", "/data/nodes.conf\n/data/appendonlydir/appendonly.aof.manifest\n/data/appendonlydir/appendonly.aof.1.incr.aof\n"} {
		if err := validateDataFiles(raw); err != nil {
			t.Fatalf("valid fixed-scope listing rejected: %q: %v", raw, err)
		}
	}
	for _, raw := range []string{
		"/etc/passwd\n", "/data-foreign/nodes.conf\n", "/data/../foreign\n", "/data//nodes.conf\n", "/data/./nodes.conf\n",
		"/data/nodes.conf", "/data/nodes.conf\n\n", "/data/nodes.conf\n/data/nodes.conf\n", "/data/a/b/c\n",
		"/data/nodes.conf\r\n", "/data/nodes\x00.conf\n", "/data/white space\n", "/data\n", "/data/\n",
		strings.Repeat("x", (1<<20)+1), strings.Repeat("/data/file\n", 4097),
	} {
		if err := validateDataFiles(raw); err == nil {
			t.Fatalf("foreign or malformed listing accepted: %.100q", raw)
		}
	}
}

func TestKeyspaceObservationAcceptsEmptyDatabaseAndRejectsMalformedCounts(t *testing.T) {
	for _, raw := range []string{"# Keyspace", "# Keyspace\ndb0:keys=17,expires=0,avg_ttl=0", "# Keyspace\r\ndb0:keys=17,expires=1,avg_ttl=99,subexpiry=0"} {
		if err := validateKeyspace(raw); err != nil {
			t.Fatalf("valid keyspace rejected: %q: %v", raw, err)
		}
	}
	for _, raw := range []string{
		"", "keys=17", "# Keyspace\nforeign:keys=1,expires=0,avg_ttl=0", "# Keyspace\ndb01:keys=1,expires=0,avg_ttl=0",
		"# Keyspace\ndb0:keys=-1,expires=0,avg_ttl=0", "# Keyspace\ndb0:keys=x,expires=0,avg_ttl=0", "# Keyspace\ndb0:keys=1,expires=0",
		"# Keyspace\ndb0:keys=1,keys=2,expires=0,avg_ttl=0", "# Keyspace\ndb0:keys=1,expires=0,avg_ttl=0\ndb0:keys=2,expires=0,avg_ttl=0",
		"# Keyspace\ndb0:keys=1,expires=0,avg_ttl=0,malformed", "# Keyspace\ndb0:keys=1,expires=0,avg_ttl=0,foreign/path=1",
	} {
		if err := validateKeyspace(raw); err == nil {
			t.Fatalf("malformed keyspace accepted: %q", raw)
		}
	}
}

func TestPersistenceEvidenceRejectsChangedPoliciesAndPreservesRawReply(t *testing.T) {
	for _, change := range []struct{ before, after string }{
		{"aof-use-rdb-preamble\nyes", "aof-use-rdb-preamble\nno"},
		{"no-appendfsync-on-rewrite\nno", "no-appendfsync-on-rewrite\nyes"},
		{"aof-load-truncated\nyes", "aof-load-truncated\nno"},
		{"appendfsync\neverysec", "appendfsync\ninvalid"},
		{"maxmemory-policy\nnoeviction", "maxmemory-policy\nnoeviction\nappendonly\nyes"},
	} {
		t.Run(change.after, func(t *testing.T) {
			manager, _ := newFake()
			env := provisionRetained(t, manager, "aof-everysec")
			original := manager.execute
			manager.execute = func(ctx context.Context, args ...string) ([]byte, error) {
				out, err := original(ctx, args...)
				if strings.Contains(strings.Join(args, " "), "config get") {
					out = []byte(strings.Replace(string(out), change.before, change.after, 1))
				}
				return out, err
			}
			evidence, err := manager.ProbeProcess(context.Background(), env, env.Nodes[0])
			if err == nil || !strings.Contains(evidence.RawPersistenceConfig, change.after) {
				t.Fatalf("changed policy accepted or evidence erased: %#v %v", evidence, err)
			}
		})
	}
}

func TestPersistenceStorageProbePreservesErrorsAndOwnedScope(t *testing.T) {
	for _, fault := range []string{"foreign-node", "foreign-volume", "foreign-listing", "malformed-usage", "canceled-before", "canceled-during-listing", "restart-policy-changed-during-listing"} {
		t.Run(fault, func(t *testing.T) {
			manager, fake := newFake()
			env := provisionRetained(t, manager, "memory")
			node := env.Nodes[0]
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			original := manager.execute
			manager.execute = func(ctx context.Context, args ...string) ([]byte, error) {
				out, err := original(ctx, args...)
				if len(args) >= 5 && args[2] == "exec" && args[4] == "find" {
					if !reflect.DeepEqual(args[4:], []string{"find", "/data", "-maxdepth", "2", "-type", "f"}) {
						t.Fatalf("file observation escaped fixed arguments: %q", args)
					}
					switch fault {
					case "foreign-listing":
						out = []byte("/host/foreign\n")
					case "canceled-during-listing":
						cancel()
					case "restart-policy-changed-during-listing":
						container := fake.containers[node.ContainerID]
						container.HostConfig.RestartPolicy.Name = "always"
						fake.containers[node.ContainerID] = container
					}
				}
				if fault == "malformed-usage" && strings.Contains(strings.Join(args, " "), "du -sk") {
					out = []byte("8\t/foreign\n")
				}
				return out, err
			}
			switch fault {
			case "foreign-node":
				node.ContainerID = strings.Repeat("f", 64)
			case "foreign-volume":
				volume := fake.volumes[env.Volumes[0].Name]
				volume.Labels[runLabel] = "other-run"
				fake.volumes[volume.Name] = volume
			case "canceled-before":
				cancel()
			}
			fake.calls = nil
			_, err := manager.ProbeProcess(ctx, env, node)
			if err == nil || errors.Is(err, ErrProcessStopped) {
				t.Fatalf("unsafe/incomplete observation accepted: %v", err)
			}
			if strings.HasPrefix(fault, "canceled") && !errors.Is(err, context.Canceled) {
				t.Fatalf("caller cancellation hidden: %v", err)
			}
			if fault == "foreign-node" || fault == "foreign-volume" || fault == "canceled-before" {
				if hasOperation(fake, "exec") {
					t.Fatal("administrative command ran outside validated ownership")
				}
			}
		})
	}
}
