package workload

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

type memoryClient struct {
	mu              sync.Mutex
	values          map[string]string
	setHook         func(context.Context, string, string) error
	getHook         func(context.Context, string) (string, error)
	duplicateWrites int
	unknownReads    int
}

func newMemoryClient() *memoryClient { return &memoryClient{values: make(map[string]string)} }

func (c *memoryClient) Set(ctx context.Context, key string, value interface{}, _ time.Duration) *redis.StatusCmd {
	text := value.(string)
	if c.setHook != nil {
		if err := c.setHook(ctx, key, text); err != nil {
			return redis.NewStatusResult("", err)
		}
	}
	c.mu.Lock()
	if _, found := c.values[key]; found {
		c.duplicateWrites++
	}
	c.values[key] = text
	c.mu.Unlock()
	return redis.NewStatusResult("OK", nil)
}

func (c *memoryClient) Get(ctx context.Context, key string) *redis.StringCmd {
	if c.getHook != nil {
		value, err := c.getHook(ctx, key)
		return redis.NewStringResult(value, err)
	}
	c.mu.Lock()
	value, found := c.values[key]
	if !found {
		c.unknownReads++
	}
	c.mu.Unlock()
	if !found {
		return redis.NewStringResult("", redis.Nil)
	}
	return redis.NewStringResult(value, nil)
}

func testConfig() Config {
	return Config{Workers: 4, RatePerSecond: 5000, MaxWrites: 30, ReadPercent: 35, ValueBytes: 80, Seed: 7, OperationTimeout: 100 * time.Millisecond}
}

func waitFor(t *testing.T, condition func() bool) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for !condition() {
		if time.Now().After(deadline) {
			t.Fatal("timed out waiting for workload condition")
		}
		time.Sleep(time.Millisecond)
	}
}

func assertRequestAccounting(t *testing.T, r *Runner) {
	t.Helper()
	s := r.Stats()
	if s.OfferedRequests != s.StartedRequests+s.DroppedRequests {
		t.Fatalf("offered requests do not reconcile after Stop: %+v", s)
	}
	if s.StartedRequests != s.CompletedRequests || s.CompletedRequests != r.Snapshot().Totals.Attempts {
		t.Fatalf("started/completed requests do not reconcile with terminal metrics: stats=%+v metrics=%+v", s, r.Snapshot().Totals)
	}
	if s.StartedRequests != s.WritesAttempted+s.ReadChecks {
		t.Fatalf("started requests do not equal actual SET/GET calls: %+v", s)
	}
	if s.PacingMethodology == "" {
		t.Fatal("pacing methodology missing")
	}
}

func TestImmutableBoundedAcknowledgedLedger(t *testing.T) {
	c := newMemoryClient()
	r, err := New(c, testConfig(), "test-run")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return r.Stats().WritesAttempted == 30 })
	r.SetPhase("fault", time.Now())
	waitFor(t, func() bool { return r.Stats().ReadChecks >= 15 })
	r.Stop()
	assertRequestAccounting(t, r)
	if err := r.Start(context.Background()); err == nil {
		t.Fatal("runner restarted")
	}
	stats := r.Stats()
	if stats.AcknowledgedWrites != 30 || stats.WritesAttempted != 30 || stats.AmbiguousWrites != 0 || !stats.WriteLimitReached {
		t.Fatalf("stats: %+v", stats)
	}
	ledger := r.Ledger()
	var shards [3]int
	for i, write := range ledger {
		if write.Sequence != uint64(i) || len(write.Value) != 80 {
			t.Fatalf("write: %+v", write)
		}
		switch {
		case write.Slot <= 5460:
			shards[0]++
		case write.Slot <= 10922:
			shards[1]++
		default:
			shards[2]++
		}
	}
	if shards != [3]int{10, 10, 10} {
		t.Fatalf("shard coverage: %v", shards)
	}
	if c.duplicateWrites != 0 || c.unknownReads != 0 {
		t.Fatalf("unsafe accesses: duplicates=%d unknown=%d", c.duplicateWrites, c.unknownReads)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	validation := r.Validate(ctx, 100*time.Millisecond)
	if !validation.Complete || validation.Checked != 30 || validation.Missing+validation.Corrupt+validation.Unreadable != 0 {
		t.Fatalf("validation: %+v", validation)
	}
	metrics := r.Snapshot()
	if metrics.Totals.Attempts != stats.WritesAttempted+stats.ReadChecks {
		t.Fatalf("metrics missing attempts: %+v", metrics.Totals)
	}
}

func TestAmbiguousSETCanExecuteButIsExcluded(t *testing.T) {
	c := newMemoryClient()
	var count atomic.Uint64
	c.setHook = func(_ context.Context, key, value string) error {
		if count.Add(1)%3 != 0 {
			return nil
		}
		c.mu.Lock()
		c.values[key] = value
		c.mu.Unlock()
		return errors.New("connection reset by peer")
	}
	config := testConfig()
	config.ReadPercent = 0
	r, err := New(c, config, "unknown-outcomes")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return r.Stats().WritesAttempted == 30 })
	r.Stop()
	assertRequestAccounting(t, r)
	stats := r.Stats()
	if stats.AcknowledgedWrites != 20 || stats.AmbiguousWrites != 10 || len(c.values) != 30 {
		t.Fatalf("unknown outcomes: stats=%+v stored=%d", stats, len(c.values))
	}
	result := r.Validate(context.Background(), time.Second)
	if result.Expected != 20 || result.Checked != 20 || !result.Complete {
		t.Fatalf("validated unacknowledged data: %+v", result)
	}
	if r.Snapshot().Totals.ErrorClasses["network"] != 10 {
		t.Fatal("unknown SET outcomes missing from metrics")
	}
}

func TestStopWaitsForSuccessfulInflightAcknowledgement(t *testing.T) {
	c := newMemoryClient()
	entered, release := make(chan struct{}), make(chan struct{})
	c.setHook = func(context.Context, string, string) error { close(entered); <-release; return nil }
	config := testConfig()
	config.Workers = 1
	config.MaxWrites = 1
	r, err := New(c, config, "stop")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	select {
	case <-entered:
	case <-time.After(time.Second):
		t.Fatal("SET not issued")
	}
	stopped := make(chan struct{})
	go func() { r.Stop(); close(stopped) }()
	select {
	case <-stopped:
		t.Fatal("Stop returned before SET finished")
	case <-time.After(10 * time.Millisecond):
	}
	close(release)
	select {
	case <-stopped:
	case <-time.After(time.Second):
		t.Fatal("Stop hung")
	}
	if r.Stats().AcknowledgedWrites != 1 || len(r.Ledger()) != 1 {
		t.Fatal("successful in-flight SET was lost")
	}
	assertRequestAccounting(t, r)
}

func TestOperationTimeoutAndCanceledParent(t *testing.T) {
	for _, cancelParent := range []bool{false, true} {
		t.Run(fmt.Sprint(cancelParent), func(t *testing.T) {
			c := newMemoryClient()
			entered := make(chan struct{})
			c.setHook = func(ctx context.Context, _, _ string) error { close(entered); <-ctx.Done(); return ctx.Err() }
			config := testConfig()
			config.Workers = 1
			config.MaxWrites = 1
			config.OperationTimeout = 15 * time.Millisecond
			r, err := New(c, config, "deadline")
			if err != nil {
				t.Fatal(err)
			}
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if err := r.Start(ctx); err != nil {
				t.Fatal(err)
			}
			select {
			case <-entered:
			case <-time.After(time.Second):
				t.Fatal("SET not issued")
			}
			if cancelParent {
				cancel()
			}
			r.Stop()
			assertRequestAccounting(t, r)
			if r.Stats().AmbiguousWrites != 1 || r.Stats().AcknowledgedWrites != 0 {
				t.Fatalf("stats: %+v", r.Stats())
			}
			class := "timeout"
			if cancelParent {
				class = "canceled"
			}
			if r.Snapshot().Totals.ErrorClasses[class] != 1 {
				t.Fatalf("wrong failure class: %+v", r.Snapshot())
			}
		})
	}
}

func TestValidationSeparatesMissingCorruptUnreadableAndCapsSamples(t *testing.T) {
	c := newMemoryClient()
	r, err := New(c, testConfig(), "validate")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 15; i++ {
		write := r.makeWrite(uint64(i))
		r.ledger = append(r.ledger, write)
		c.values[write.Key] = write.Value
	}
	r.Stop()
	index := make(map[string]int)
	for i, write := range r.Ledger() {
		index[write.Key] = i
	}
	c.getHook = func(_ context.Context, key string) (string, error) {
		switch index[key] % 3 {
		case 0:
			return "", redis.Nil
		case 1:
			return "wrong", nil
		default:
			return "", errors.New("CLUSTERDOWN failed")
		}
	}
	result := r.Validate(context.Background(), time.Second)
	if !result.Complete || result.Checked != 15 || result.Missing != 5 || result.Corrupt != 5 || result.Unreadable != 5 || len(result.Samples) != 10 {
		t.Fatalf("validation: %+v", result)
	}
}

func TestValidationDeadlineLeavesUnverifiedEntries(t *testing.T) {
	c := newMemoryClient()
	r, err := New(c, testConfig(), "validation-timeout")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		r.ledger = append(r.ledger, r.makeWrite(uint64(i)))
	}
	r.Stop()
	c.getHook = func(ctx context.Context, _ string) (string, error) { <-ctx.Done(); return "", ctx.Err() }
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	result := r.Validate(ctx, time.Second)
	if result.Complete || result.Checked != 1 || result.Unreadable != 1 || result.Unverified != 2 || result.Error == "" {
		t.Fatalf("validation: %+v", result)
	}
}

func TestDeterministicKeysValuesAndCRC16(t *testing.T) {
	r, _ := New(newMemoryClient(), testConfig(), "same-run")
	other, _ := New(newMemoryClient(), testConfig(), "same-run")
	for i := uint64(0); i < 12; i++ {
		left, right := r.makeWrite(i), other.makeWrite(i)
		if left.Key != right.Key || left.Value != right.Value || left.Slot != right.Slot {
			t.Fatal("workload data is nondeterministic")
		}
	}
	if HashSlot("123456789") != 0x31c3%16384 || HashSlot("foo{bar}zap") != HashSlot("bar") || HashSlot("foo{}zap") == HashSlot("") {
		t.Fatal("CRC16/hash tag semantics incorrect")
	}
}

func TestAggregateRequestRate(t *testing.T) {
	config := testConfig()
	config.RatePerSecond = 40
	config.MaxWrites = 100
	config.ReadPercent = 0
	r, err := New(newMemoryClient(), config, "rate")
	if err != nil {
		t.Fatal(err)
	}
	started := time.Now()
	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	time.Sleep(200 * time.Millisecond)
	r.Stop()
	elapsed := time.Since(started).Seconds()
	if got := r.Snapshot().Totals.Attempts; got < 3 || float64(got) > elapsed*config.RatePerSecond+2 {
		t.Fatalf("global rate bound broken: %d in %v seconds", got, elapsed)
	}
	assertRequestAccounting(t, r)
}

func TestDispatcherCountsBusyOffersAndStopDrainsWithoutFurtherOffers(t *testing.T) {
	c := newMemoryClient()
	entered := make(chan struct{}, 2)
	release := make(chan struct{})
	c.setHook = func(context.Context, string, string) error {
		entered <- struct{}{}
		<-release
		return nil
	}
	config := testConfig()
	config.Workers = 2
	config.MaxWrites = 1000
	config.ReadPercent = 0
	config.RatePerSecond = 1000
	r, err := New(c, config, "saturated")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < config.Workers; i++ {
		select {
		case <-entered:
		case <-time.After(time.Second):
			close(release)
			r.Stop()
			t.Fatal("workers did not enter SET")
		}
	}
	waitFor(t, func() bool { return r.Stats().OfferedRequests >= 12 })
	before := r.Stats()
	if before.StartedRequests != 2 || before.CompletedRequests != 0 || before.DroppedRequests == 0 {
		close(release)
		r.Stop()
		t.Fatalf("saturation missing from counters: %+v", before)
	}
	stopped := make(chan struct{})
	go func() { r.Stop(); close(stopped) }()
	select {
	case <-r.dispatchDone:
	case <-time.After(time.Second):
		close(release)
		t.Fatal("dispatcher did not stop while requests were in flight")
	}
	offered := r.Stats().OfferedRequests
	time.Sleep(5 * time.Millisecond)
	if got := r.Stats().OfferedRequests; got != offered {
		close(release)
		<-stopped
		t.Fatalf("offered requests increased during Stop drain: %d -> %d", offered, got)
	}
	select {
	case <-stopped:
		t.Fatal("Stop did not drain in-flight requests")
	default:
	}
	close(release)
	select {
	case <-stopped:
	case <-time.After(time.Second):
		t.Fatal("Stop did not finish after in-flight requests completed")
	}
	if got := r.Stats(); got.StartedRequests != 2 || got.AcknowledgedWrites != 2 {
		t.Fatalf("queued or lost requests at Stop: %+v", got)
	}
	assertRequestAccounting(t, r)
}

func TestMissedScheduleSlotsProduceOnlyOneCurrentOffer(t *testing.T) {
	r, err := New(newMemoryClient(), testConfig(), "grid")
	if err != nil {
		t.Fatal(err)
	}
	start := time.Now()
	period := 10 * time.Millisecond
	var last uint64
	if r.offerSlots(start, start.Add(period-time.Nanosecond), period, &last, true) {
		t.Fatal("offer occurred before the first ideal period")
	}
	// A delayed dispatcher accounts for ten offers but permits only one start.
	if !r.offerSlots(start, start.Add(105*time.Millisecond), period, &last, true) {
		t.Fatal("latest slot was not dispatchable")
	}
	if s := r.Stats(); s.OfferedRequests != 10 || s.DroppedRequests != 9 {
		t.Fatalf("missed slots not counted: %+v", s)
	}
	if r.offerSlots(start, start.Add(109*time.Millisecond), period, &last, true) {
		t.Fatal("duplicate offer within the same ideal slot")
	}
	// Termination accounts for remaining ideal slots without issuing a request.
	if r.offerSlots(start, start.Add(139*time.Millisecond), period, &last, false) {
		t.Fatal("termination offered a new request")
	}
	if s := r.Stats(); s.OfferedRequests != 13 || s.DroppedRequests != 12 {
		t.Fatalf("terminal skipped slots not counted: %+v", s)
	}
}

func TestNoEligibleOperationsDropOffersWithoutExceedingWriteLimit(t *testing.T) {
	c := newMemoryClient()
	c.setHook = func(context.Context, string, string) error { return errors.New("connection reset by peer") }
	config := testConfig()
	config.Workers = 8
	config.MaxWrites = 3
	config.ReadPercent = 100
	config.RatePerSecond = 10000
	r, err := New(c, config, "no-eligible")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
	waitFor(t, func() bool { return r.Stats().OfferedRequests >= 30 && r.Stats().AmbiguousWrites == 3 })
	r.Stop()
	s := r.Stats()
	if s.StartedRequests != 3 || s.WritesAttempted != 3 || s.AmbiguousWrites != 3 || s.ReadChecks != 0 || s.DroppedRequests == 0 {
		t.Fatalf("ineligible requests or write-limit overflow: %+v", s)
	}
	assertRequestAccounting(t, r)
}

func TestStopWithoutStartHasZeroOfferedRequests(t *testing.T) {
	r, err := New(newMemoryClient(), testConfig(), "not-started")
	if err != nil {
		t.Fatal(err)
	}
	r.Stop()
	s := r.Stats()
	if s.OfferedRequests+s.StartedRequests+s.CompletedRequests+s.DroppedRequests != 0 {
		t.Fatalf("unstarted workload had requests: %+v", s)
	}
	assertRequestAccounting(t, r)
}

func TestRateCannotOverflowTimerInterval(t *testing.T) {
	config := testConfig()
	config.RatePerSecond = 1e-20
	if _, err := New(newMemoryClient(), config, "tiny-rate"); err == nil {
		t.Fatal("accepted a rate whose interval overflows time.Duration")
	}
}
