package workload

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/metrics"
	"github.com/redis/go-redis/v9"
)

type writerFunc func(context.Context, string, string) WriteResult

func (f writerFunc) Write(ctx context.Context, key, value string) WriteResult {
	return f(ctx, key, value)
}

func waitConfig(writes int) Config {
	c := testConfig()
	c.Workers, c.MaxWrites, c.ReadPercent = 1, writes, 0
	c.WritePolicy = "wait"
	return c
}

func acknowledgedWait(acknowledgements int64) WriteResult {
	return WriteResult{SetAcknowledged: true, WaitAttempted: true, WaitRequested: 1,
		WaitAcknowledgements: acknowledgements, PolicySatisfied: acknowledgements >= 1,
		SetLatency: time.Millisecond, WaitLatency: 3 * time.Millisecond, TotalLatency: 4 * time.Millisecond}
}

func TestWAITLogicalStagesAndAcknowledgedLedgerRemainDistinct(t *testing.T) {
	c := newMemoryClient()
	var calls atomic.Uint64
	writer := writerFunc(func(ctx context.Context, key, value string) WriteResult {
		// Even the final error-returning SET may have executed. Only its client
		// acknowledgement status decides whether the key enters the ledger.
		if err := c.Set(ctx, key, value, 0).Err(); err != nil {
			t.Fatal(err)
		}
		switch calls.Add(1) {
		case 1:
			result := acknowledgedWait(1)
			result.Redirects = 2
			return result
		case 2:
			return acknowledgedWait(0)
		case 3:
			result := acknowledgedWait(0)
			result.WaitError = errors.New("connection reset by peer")
			return result
		case 4:
			result := acknowledgedWait(0)
			result.WaitError = context.DeadlineExceeded
			return result
		default:
			return WriteResult{SetError: errors.New("connection reset by peer"), SetLatency: time.Millisecond, TotalLatency: time.Millisecond}
		}
	})
	r, err := NewWithWriter(c, writer, waitConfig(5), "wait-accounting")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return r.Stats().WritesAttempted == 5 })
	r.Stop()
	assertRequestAccounting(t, r)
	s := r.Stats()
	p := s.PolicyTotals
	if s.AcknowledgedWrites != 4 || s.AmbiguousWrites != 1 || p.SetAcknowledged != 4 || p.SetErrors != 1 || p.WaitAttempted != 4 || p.WaitSatisfied != 1 || p.WaitShort != 1 || p.WaitErrors != 2 || p.WaitClientTimeouts != 1 || p.PolicySatisfied != 1 || p.HandledRedirects != 2 {
		t.Fatalf("policy accounting: %+v", s)
	}
	if len(p.WaitAcknowledgements) != 2 || p.WaitAcknowledgements[1] != 1 || p.WaitAcknowledgements[0] != 1 {
		t.Fatalf("failed WAIT outcomes invented acknowledgement replies: %+v", p.WaitAcknowledgements)
	}
	logical := r.Snapshot().Totals.ByOperation["set_wait"]
	if logical.Attempts != 5 || logical.Successes != 1 || logical.Errors != 4 || logical.ErrorClasses["wait_short"] != 1 || logical.ErrorClasses["timeout"] != 1 {
		t.Fatalf("logical write outcomes: %+v", logical)
	}
	stages := s.WriteStages.Totals
	set, wait := stages.ByOperation["set"], stages.ByOperation["wait"]
	if stages.Attempts != 9 || set.Attempts != 5 || set.Successes != 4 || wait.Attempts != 4 || wait.Successes != 1 || wait.Errors != 3 {
		t.Fatalf("stage counts contaminate or lose operations: %+v", stages)
	}
	if set.Latency.MinMS != 1 || set.Latency.MaxMS != 1 || wait.Latency.MinMS != 3 || wait.Latency.MaxMS != 3 || logical.SuccessfulLatency.MinMS != 4 {
		t.Fatal("SET, WAIT, and total logical latency were not recorded separately")
	}
	for i, write := range r.Ledger() {
		if !write.WaitAttempted || write.WaitRequested != 1 || write.PolicySatisfied != (i == 0) {
			t.Fatalf("SET acknowledgement lost WAIT policy result: %+v", write)
		}
	}
	v := r.Validate(context.Background(), time.Second)
	if v.Expected != 4 || v.Checked != 4 || !v.Complete || v.WaitSatisfiedSubset.Expected != 1 || v.WaitSatisfiedSubset.Checked != 1 || !v.WaitSatisfiedSubset.Complete {
		t.Fatalf("validation lost unsuccessful WAIT SET acknowledgements: %+v", v)
	}
	// Snapshot maps must be independent of callers and concurrent future reads.
	s.PolicyTotals.WaitAcknowledgements[0] = 1000
	counts := s.PolicyByPhase["baseline"]
	counts.WaitAcknowledgements[0] = 1000
	if r.Stats().PolicyTotals.WaitAcknowledgements[0] != 1 || r.Stats().PolicyByPhase["baseline"].WaitAcknowledgements[0] != 1 {
		t.Fatal("policy snapshot exposes mutable collector state")
	}
}

func TestWAITSubsetUsesSameGETAndPreservesLossAccounting(t *testing.T) {
	c := newMemoryClient()
	var calls atomic.Uint64
	writer := writerFunc(func(ctx context.Context, key, value string) WriteResult {
		c.Set(ctx, key, value, 0)
		n := calls.Add(1)
		if n == 3 {
			return acknowledgedWait(0)
		}
		result := acknowledgedWait(1)
		if n == 4 {
			result.WaitError, result.PolicySatisfied = errors.New("connection reset by peer"), false
		}
		return result
	})
	r, err := NewWithWriter(c, writer, waitConfig(5), "wait-subset-loss")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return r.Stats().WritesAttempted == 5 })
	r.Stop()
	ledger := r.Ledger()
	var gets atomic.Uint64
	c.getHook = func(_ context.Context, key string) (string, error) {
		gets.Add(1)
		switch key {
		case ledger[0].Key, ledger[2].Key:
			return "", redis.Nil
		case ledger[1].Key:
			return "corrupt", nil
		case ledger[4].Key:
			return "", errors.New("connection refused")
		default:
			return ledger[3].Value, nil
		}
	}
	v := r.Validate(context.Background(), time.Second)
	subset := v.WaitSatisfiedSubset
	if gets.Load() != 5 || v.Checked != 5 || v.Missing != 2 || v.Corrupt != 1 || v.Unreadable != 1 || !v.Complete {
		t.Fatalf("all-SET validation: gets=%d validation=%+v", gets.Load(), v)
	}
	if subset.Expected != 3 || subset.Checked != 3 || subset.Missing != 1 || subset.Corrupt != 1 || subset.Unreadable != 1 || subset.Unverified != 0 || !subset.Complete {
		t.Fatalf("WAIT subset validation: %+v", subset)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	unverified := r.Validate(ctx, time.Second)
	if unverified.Checked != 0 || unverified.Unverified != 5 || unverified.Complete || unverified.WaitSatisfiedSubset.Checked != 0 || unverified.WaitSatisfiedSubset.Unverified != 3 || unverified.WaitSatisfiedSubset.Complete {
		t.Fatalf("cancellation invented checked subset values: %+v", unverified)
	}
}

func TestStopWaitsForWAITAndRetainsSuccessfulSETAfterClientTimeout(t *testing.T) {
	c := newMemoryClient()
	started := make(chan time.Time, 1)
	writer := writerFunc(func(ctx context.Context, key, value string) WriteResult {
		c.Set(ctx, key, value, 0)
		at := time.Now()
		started <- at
		<-ctx.Done()
		result := acknowledgedWait(0)
		result.WaitError = ctx.Err()
		result.SetAcknowledgedAt = at
		result.WaitLatency = time.Since(at)
		result.TotalLatency = result.SetLatency + result.WaitLatency
		return result
	})
	cfg := waitConfig(1)
	cfg.OperationTimeout = 30 * time.Millisecond
	r, err := NewWithWriter(c, writer, cfg, "wait-stop-timeout")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	acknowledgedAt := <-started
	r.SetPhase("fault", time.Now())
	r.Stop()
	assertRequestAccounting(t, r)
	s := r.Stats()
	if s.AcknowledgedWrites != 1 || s.AmbiguousWrites != 0 || s.PolicyTotals.WaitClientTimeouts != 1 || s.PolicyTotals.PolicySatisfied != 0 || len(r.Ledger()) != 1 {
		t.Fatalf("in-flight SET acknowledgement lost: %+v", s)
	}
	if !r.Ledger()[0].AcknowledgedAt.Equal(acknowledgedAt) {
		t.Fatal("ledger acknowledgement timestamp moved from SET completion to WAIT completion")
	}
	if s.PolicyByPhase["baseline"].SetAcknowledged != 1 || s.PolicyByPhase["fault"].SetAcknowledged != 0 || r.Snapshot().Phases["baseline"].ByOperation["set_wait"].Errors != 1 || s.WriteStages.Phases["baseline"].ByOperation["wait"].Errors != 1 {
		t.Fatal("operation crossing phase boundary changed attribution")
	}
}

func TestClientDeadlineBeforeWAITPreservesSETAndCountsUnattemptedWAIT(t *testing.T) {
	c := newMemoryClient()
	writer := writerFunc(func(ctx context.Context, key, value string) WriteResult {
		c.Set(ctx, key, value, 0)
		acknowledgedAt := time.Now()
		<-ctx.Done()
		return WriteResult{SetAcknowledged: true, SetAcknowledgedAt: acknowledgedAt, WaitRequested: 1,
			WaitError: ctx.Err(), SetLatency: time.Millisecond, TotalLatency: time.Since(acknowledgedAt)}
	})
	cfg := waitConfig(1)
	cfg.OperationTimeout = 10 * time.Millisecond
	r, err := NewWithWriter(c, writer, cfg, "wait-not-issued")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return r.Stats().PolicyTotals.WaitNotAttempted == 1 })
	r.Stop()
	assertRequestAccounting(t, r)
	s := r.Stats()
	p := s.PolicyTotals
	if s.AcknowledgedWrites != 1 || s.AmbiguousWrites != 0 || p.WaitAttempted != 0 || p.WaitErrors != 0 || p.WaitClientTimeouts != 0 || p.WaitNotAttempted != 1 || p.WaitNotAttemptedTimeouts != 1 || p.PolicySatisfied != 0 || len(p.WaitAcknowledgements) != 0 {
		t.Fatalf("unattempted WAIT accounting: %+v", s)
	}
	if s.WriteStages.Totals.Attempts != 1 || s.WriteStages.Totals.ByOperation["wait"].Attempts != 0 || r.Snapshot().Totals.ByOperation["set_wait"].ErrorClasses["timeout"] != 1 {
		t.Fatal("deadline was misclassified or an unissued WAIT was recorded as a stage command")
	}
	v := r.Validate(context.Background(), time.Second)
	if v.Expected != 1 || v.Checked != 1 || v.WaitSatisfiedSubset.Expected != 0 {
		t.Fatalf("acknowledged SET disappeared when WAIT was never issued: %+v", v)
	}
}

func TestPolicyValidationAndNormalSETCompatibility(t *testing.T) {
	c := newMemoryClient()
	if _, err := New(c, waitConfig(1), "requires-affinity"); err == nil {
		t.Fatal("WAIT allowed without an injected connection-affine executor")
	}
	if _, err := NewWithWriter(c, nil, testConfig(), "missing-writer"); err == nil {
		t.Fatal("missing writer accepted")
	}
	invalid := testConfig()
	invalid.WritePolicy = "arbitrary"
	if _, err := New(c, invalid, "invalid-policy"); err == nil {
		t.Fatal("unsupported policy accepted")
	}
	r, err := New(c, testConfig(), "normal-policy")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return r.Stats().WritesAttempted == 30 })
	r.Stop()
	assertRequestAccounting(t, r)
	s := r.Stats()
	if s.WritePolicy != "set" || s.PolicyTotals.SetAcknowledged != 30 || s.PolicyTotals.PolicySatisfied != 30 || s.PolicyTotals.WaitAttempted != 0 || s.WriteStages.Totals.Attempts != 30 {
		t.Fatalf("normal SET behavior changed: %+v", s)
	}
	if r.Snapshot().Totals.ByOperation["set"].Attempts != 30 || r.Snapshot().Totals.ByOperation["set_wait"].Attempts != 0 {
		t.Fatal("normal SET used WAIT operation population")
	}
	v := r.Validate(context.Background(), time.Second)
	if v.WaitSatisfiedSubset.Expected != 0 || v.WaitSatisfiedSubset.Checked != 0 || !v.WaitSatisfiedSubset.Complete {
		t.Fatalf("normal SET invented WAIT-satisfied records: %+v", v)
	}
}

func TestUnsatisfiedWAITIsATerminalLogicalError(t *testing.T) {
	for _, result := range []WriteResult{
		{SetAcknowledged: true, PolicySatisfied: true},
		{SetAcknowledged: true, WaitAttempted: true, WaitRequested: 1, WaitAcknowledgements: 1},
		{PolicySatisfied: true},
	} {
		if err := logicalWriteError("wait", result); err == nil {
			t.Fatalf("incomplete writer result falsely succeeded: %s", fmt.Sprint(result))
		}
	}
	if metrics.ClassifyError(logicalWriteError("wait", acknowledgedWait(0))) != "wait_short" {
		t.Fatal("short WAIT was not classified separately from client timeout")
	}
}
