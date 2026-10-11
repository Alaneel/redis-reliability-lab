package workload

import (
	"context"
	"errors"
	"reflect"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/metrics"
	"github.com/redis/go-redis/v9"
)

func waitAOFConfig(writes int) Config {
	c := waitConfig(writes)
	c.WritePolicy = "waitaof"
	return c
}

func acknowledgedWaitAOF(local, replicas int64) WriteResult {
	return WriteResult{SetAcknowledged: true, WaitAOFAttempted: true,
		WaitAOFRequestedLocal: 1, WaitAOFRequestedReplicas: 1,
		WaitAOFLocalAcknowledgements: local, WaitAOFReplicaAcknowledgements: replicas,
		PolicySatisfied: local == 1 && replicas >= 1,
		SetLatency:      time.Millisecond, WaitAOFLatency: 3 * time.Millisecond, TotalLatency: 4 * time.Millisecond}
}

func TestWAITAOFReplyPairsStagesAndLedgerAreIndependent(t *testing.T) {
	c := newMemoryClient()
	results := []WriteResult{acknowledgedWaitAOF(1, 1), acknowledgedWaitAOF(1, 0), acknowledgedWaitAOF(0, 1), acknowledgedWaitAOF(0, 0), acknowledgedWaitAOF(0, 0), acknowledgedWaitAOF(0, 0), acknowledgedWaitAOF(2, 1), {SetError: errors.New("connection reset by peer"), SetLatency: time.Millisecond}}
	results[4].WaitAOFError = errors.New("connection reset by peer")
	results[5].WaitAOFError = context.DeadlineExceeded
	results[0].Redirects = 2
	var calls atomic.Uint64
	writer := writerFunc(func(ctx context.Context, key, value string) WriteResult {
		// Even the unacknowledged SET may execute; it stays outside the ledger.
		c.Set(ctx, key, value, 0)
		return results[calls.Add(1)-1]
	})
	r, err := NewWithWriter(c, writer, waitAOFConfig(len(results)), "waitaof-accounting")
	if err != nil {
		t.Fatal(err)
	}
	if err = r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return r.Stats().WritesAttempted == uint64(len(results)) })
	r.Stop()
	assertRequestAccounting(t, r)
	s, logical := r.Stats(), r.Snapshot().Totals.ByOperation["set_waitaof"]
	p := s.PolicyTotals
	if s.AcknowledgedWrites != 7 || s.AmbiguousWrites != 1 || len(r.Ledger()) != 7 || p.SetAcknowledged != 7 || p.SetErrors != 1 || p.PolicySatisfied != 1 || p.HandledRedirects != 2 || p.WaitAOFAttempted != 7 || p.WaitAOFSatisfied != 1 || p.WaitAOFShort != 3 || p.WaitAOFErrors != 3 || p.WaitAOFClientTimeouts != 1 || p.WaitAOFLocalSatisfied != 2 || p.WaitAOFReplicaSatisfied != 2 {
		t.Fatalf("WAITAOF counter reconciliation: %+v", s)
	}
	if !reflect.DeepEqual(p.WaitAOFReplies, map[string]uint64{"1/1": 1, "1/0": 1, "0/1": 1, "0/0": 1}) || p.WaitAttempted != 0 || p.WaitSatisfied != 0 || len(p.WaitAcknowledgements) != 0 {
		t.Fatalf("reply pairs or WAIT isolation wrong: %+v", p)
	}
	if logical.Attempts != 8 || logical.Successes != 1 || logical.Errors != 7 || logical.ErrorClasses["waitaof_short"] != 3 || logical.ErrorClasses["timeout"] != 1 || logical.ErrorClasses["network"] != 2 || logical.ErrorClasses["other"] != 1 || logical.Latency.Count != 8 {
		t.Fatalf("logical writes lost policy failures: %+v", logical)
	}
	stages := s.WriteStages.Totals
	if stages.Attempts != 15 || stages.ByOperation["set"].Attempts != 8 || stages.ByOperation["set"].Successes != 7 || stages.ByOperation["waitaof"].Attempts != 7 || stages.ByOperation["waitaof"].ErrorClasses["waitaof_short"] != 3 || stages.ByOperation["wait"].Attempts != 0 || r.Snapshot().Totals.Attempts != logical.Attempts+r.Snapshot().Totals.ByOperation["get"].Attempts {
		t.Fatalf("diagnostic stages leaked into logical traffic: %+v", stages)
	}
	v := r.Validate(context.Background(), time.Second)
	if v.Expected != 7 || v.Checked != 7 || v.Missing+v.Corrupt+v.Unreadable+v.Unverified != 0 || v.WaitAOFSatisfiedSubset.Expected != 1 || v.WaitAOFSatisfiedSubset.Checked != 1 || v.WaitSatisfiedSubset.Expected != 0 {
		t.Fatalf("successful SETs or policy subset lost: %+v", v)
	}
	s.PolicyTotals.WaitAOFReplies["1/1"] = 100
	phase := s.PolicyByPhase["baseline"]
	phase.WaitAOFReplies["0/0"] = 100
	if r.Stats().PolicyTotals.WaitAOFReplies["1/1"] != 1 || r.Stats().PolicyByPhase["baseline"].WaitAOFReplies["0/0"] != 1 {
		t.Fatal("snapshot aliases reply maps")
	}
}

func TestWAITAOFSubsetValidationUsesOneReadAndPreservesLoss(t *testing.T) {
	c := newMemoryClient()
	var calls atomic.Uint64
	writer := writerFunc(func(ctx context.Context, key, value string) WriteResult {
		c.Set(ctx, key, value, 0)
		if calls.Add(1) <= 3 {
			return acknowledgedWaitAOF(1, 1)
		}
		return acknowledgedWaitAOF(1, 0)
	})
	r, err := NewWithWriter(c, writer, waitAOFConfig(4), "waitaof-subset")
	if err != nil {
		t.Fatal(err)
	}
	if err = r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return r.Stats().WritesAttempted == 4 })
	r.Stop()
	ledger := r.Ledger()
	var gets atomic.Uint64
	c.getHook = func(_ context.Context, key string) (string, error) {
		gets.Add(1)
		switch key {
		case ledger[0].Key:
			return "", redis.Nil
		case ledger[1].Key:
			return "corrupt", nil
		case ledger[2].Key:
			return "", errors.New("connection refused")
		default:
			return ledger[3].Value, nil
		}
	}
	v := r.Validate(context.Background(), time.Second)
	want := ValidationSubset{Expected: 3, Checked: 3, Missing: 1, Corrupt: 1, Unreadable: 1, Complete: true}
	if gets.Load() != 4 || v.Checked != 4 || v.Missing != 1 || v.Corrupt != 1 || v.Unreadable != 1 || !v.Complete || v.WaitAOFSatisfiedSubset != want || v.WaitSatisfiedSubset.Expected != 0 {
		t.Fatalf("one GET must classify main ledger and AOF subset: gets=%d v=%+v", gets.Load(), v)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	v = r.Validate(ctx, time.Second)
	if v.Checked != 0 || v.Unverified != 4 || v.WaitAOFSatisfiedSubset.Checked != 0 || v.WaitAOFSatisfiedSubset.Unverified != 3 || v.WaitAOFSatisfiedSubset.Complete {
		t.Fatalf("canceled validation fabricated subset coverage: %+v", v)
	}
}

func TestWAITAOFDeadlineRetainsSETAndDistinguishesUnissuedCommand(t *testing.T) {
	for _, issued := range []bool{false, true} {
		t.Run(map[bool]string{false: "not issued", true: "issued"}[issued], func(t *testing.T) {
			c := newMemoryClient()
			started := make(chan time.Time, 1)
			writer := writerFunc(func(ctx context.Context, key, value string) WriteResult {
				c.Set(ctx, key, value, 0)
				at := time.Now()
				started <- at
				<-ctx.Done()
				result := acknowledgedWaitAOF(0, 0)
				result.SetAcknowledgedAt, result.WaitAOFAttempted, result.WaitAOFError = at, issued, ctx.Err()
				return result
			})
			cfg := waitAOFConfig(1)
			cfg.OperationTimeout = 20 * time.Millisecond
			r, err := NewWithWriter(c, writer, cfg, "waitaof-deadline")
			if err != nil {
				t.Fatal(err)
			}
			if err = r.Start(context.Background()); err != nil {
				t.Fatal(err)
			}
			at := <-started
			r.SetPhase("fault", time.Now())
			r.Stop()
			assertRequestAccounting(t, r)
			s, p := r.Stats(), r.Stats().PolicyTotals
			if s.AcknowledgedWrites != 1 || s.AmbiguousWrites != 0 || !r.Ledger()[0].AcknowledgedAt.Equal(at) || r.Snapshot().Phases["baseline"].ByOperation["set_waitaof"].ErrorClasses["timeout"] != 1 || s.PolicyByPhase["fault"].SetAcknowledged != 0 {
				t.Fatalf("SET acknowledgement or initial phase lost: %+v", s)
			}
			if issued {
				if p.WaitAOFAttempted != 1 || p.WaitAOFErrors != 1 || p.WaitAOFClientTimeouts != 1 || p.WaitAOFNotAttempted != 0 || s.WriteStages.Totals.Attempts != 2 {
					t.Fatalf("attempted timeout: %+v", p)
				}
			} else if p.WaitAOFAttempted != 0 || p.WaitAOFErrors != 0 || p.WaitAOFClientTimeouts != 0 || p.WaitAOFNotAttempted != 1 || p.WaitAOFNotAttemptedTimeouts != 1 || s.WriteStages.Totals.Attempts != 1 {
				t.Fatalf("unissued timeout: %+v", p)
			}
			if len(p.WaitAOFReplies) != 0 || p.WaitAttempted != 0 || p.PolicySatisfied != 0 {
				t.Fatal("timeout fabricated reply or policy success")
			}
			v := r.Validate(context.Background(), time.Second)
			if v.Expected != 1 || v.Checked != 1 || v.WaitAOFSatisfiedSubset.Expected != 0 {
				t.Fatalf("SET must survive incomplete policy: %+v", v)
			}
		})
	}
}

func TestWAITAOFRequiresBothActualCountsAndAffineWriter(t *testing.T) {
	if _, err := New(newMemoryClient(), waitAOFConfig(1), "unsafe-affinity"); err == nil {
		t.Fatal("WAITAOF allowed without affine writer")
	}
	for _, pair := range [][2]int64{{1, 0}, {0, 1}, {0, 0}} {
		result := acknowledgedWaitAOF(pair[0], pair[1])
		result.PolicySatisfied = true
		if !errors.Is(logicalWriteError("waitaof", result), metrics.ErrWaitAOFShort) {
			t.Fatalf("single count fabricated both-required success: %+v", result)
		}
	}
	for _, pair := range [][2]int64{{-1, 1}, {2, 1}, {1, -1}} {
		result := acknowledgedWaitAOF(pair[0], pair[1])
		result.PolicySatisfied = true
		if !errors.Is(logicalWriteError("waitaof", result), errWaitAOFInvalidReply) {
			t.Fatalf("invalid reply accepted: %+v", result)
		}
	}
	result := acknowledgedWaitAOF(1, 1)
	result.WaitAOFAttempted = false
	if !errors.Is(logicalWriteError("waitaof", result), errWaitAOFNotAttempted) {
		t.Fatal("unissued command invented success")
	}
}
