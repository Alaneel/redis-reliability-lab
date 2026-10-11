// Package metrics collects bounded, concurrency-safe client operation metrics.
package metrics

import (
	"context"
	"errors"
	"math"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ErrMissing      = errors.New("acknowledged key missing")
	ErrMismatch     = errors.New("acknowledged value mismatch")
	ErrWaitShort    = errors.New("WAIT returned fewer replica acknowledgements than requested")
	ErrWaitAOFShort = errors.New("WAITAOF returned fewer local or replica AOF acknowledgements than requested")
)

const Methodology = "Latencies use 512 bounded geometric histogram buckets, starting at 1 microsecond and increasing by 5%; quantiles are nearest-rank bucket upper bounds. The last bucket covers the remaining duration range. All terminal client attempts, successful attempts, and failed attempts have separate histograms. Throughput is completed attempts divided by phase wall-clock seconds. Operations are attributed to the phase in which they start. Redirect counts include only terminal MOVED/ASK errors; redirects handled by the client are not observable here."

const histogramBuckets = 512
const maxWindows = 3600

var bounds = makeBounds()

func makeBounds() [histogramBuckets]time.Duration {
	var b [histogramBuckets]time.Duration
	b[0] = time.Microsecond
	for i := 1; i < len(b)-1; i++ {
		b[i] = time.Duration(math.Ceil(float64(b[i-1]) * 1.05))
	}
	b[len(b)-1] = time.Duration(math.MaxInt64)
	return b
}

type histogram struct {
	counts   [histogramBuckets]uint64
	count    uint64
	min, max time.Duration
}

func (h *histogram) add(d time.Duration) {
	if d < 0 {
		d = 0
	}
	if h.count == 0 || d < h.min {
		h.min = d
	}
	if d > h.max {
		h.max = d
	}
	h.count++
	lo, hi := 0, len(bounds)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if bounds[mid] < d {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	h.counts[lo]++
}

func (h *histogram) quantile(p float64) time.Duration {
	if h.count == 0 {
		return 0
	}
	rank := uint64(math.Ceil(float64(h.count) * p))
	var n uint64
	for i, count := range h.counts {
		n += count
		if n >= rank {
			// The actual maximum gives a useful finite result for the overflow bucket.
			if bounds[i] > h.max {
				return h.max
			}
			return bounds[i]
		}
	}
	return h.max
}

type LatencySummary struct {
	Count uint64  `json:"count"`
	MinMS float64 `json:"min_ms"`
	P50MS float64 `json:"p50_ms"`
	P95MS float64 `json:"p95_ms"`
	P99MS float64 `json:"p99_ms"`
	MaxMS float64 `json:"max_ms"`
}

func (h *histogram) snapshot() LatencySummary {
	ms := func(d time.Duration) float64 { return float64(d) / float64(time.Millisecond) }
	return LatencySummary{h.count, ms(h.min), ms(h.quantile(.50)), ms(h.quantile(.95)), ms(h.quantile(.99)), ms(h.max)}
}

type OperationSummary struct {
	Attempts          uint64            `json:"attempts"`
	Successes         uint64            `json:"successes"`
	Errors            uint64            `json:"errors"`
	ErrorClasses      map[string]uint64 `json:"error_classes"`
	Latency           LatencySummary    `json:"latency"`
	SuccessfulLatency LatencySummary    `json:"successful_latency"`
	FailedLatency     LatencySummary    `json:"failed_latency"`
}

type Summary struct {
	Attempts          uint64                      `json:"attempts"`
	Successes         uint64                      `json:"successes"`
	Errors            uint64                      `json:"errors"`
	ErrorClasses      map[string]uint64           `json:"error_classes"`
	ByOperation       map[string]OperationSummary `json:"by_operation"`
	Latency           LatencySummary              `json:"latency"`
	SuccessfulLatency LatencySummary              `json:"successful_latency"`
	FailedLatency     LatencySummary              `json:"failed_latency"`
	DurationSeconds   float64                     `json:"duration_seconds"`
	Throughput        float64                     `json:"throughput_ops_per_second"`
	SuccessRate       float64                     `json:"success_rate"`
}

type PhaseSummary = Summary

type PhaseTransition struct {
	Phase string    `json:"phase"`
	At    time.Time `json:"at"`
}

type Window struct {
	StartedAt   time.Time         `json:"started_at"`
	EndedAt     time.Time         `json:"ended_at"`
	Attempts    uint64            `json:"attempts"`
	Successes   uint64            `json:"successes"`
	Errors      uint64            `json:"errors"`
	SuccessRate float64           `json:"success_rate"`
	Phases      map[string]uint64 `json:"phases"`
}

type Snapshot struct {
	StartedAt        time.Time               `json:"started_at"`
	EndedAt          time.Time               `json:"ended_at"`
	Phases           map[string]PhaseSummary `json:"phases"`
	Totals           Summary                 `json:"totals"`
	Windows          []Window                `json:"windows"`
	DroppedWindows   uint64                  `json:"dropped_windows"`
	PhaseTransitions []PhaseTransition       `json:"phase_transitions"`
	Methodology      string                  `json:"methodology"`
}

type aggregate struct {
	attempts, successes   uint64
	errors                map[string]uint64
	all, success, failure histogram
	ops                   map[string]*aggregate
}

func newAggregate() *aggregate {
	return &aggregate{errors: make(map[string]uint64), ops: make(map[string]*aggregate)}
}

func (a *aggregate) record(operation string, d time.Duration, err error) {
	a.attempts++
	a.all.add(d)
	if err == nil {
		a.successes++
		a.success.add(d)
	} else {
		a.errors[ClassifyError(err)]++
		a.failure.add(d)
	}
	if operation != "" {
		op := a.ops[operation]
		if op == nil {
			op = newAggregate()
			a.ops[operation] = op
		}
		op.record("", d, err)
	}
}

func copyCounts(source map[string]uint64) map[string]uint64 {
	dest := make(map[string]uint64, len(source))
	for key, count := range source {
		dest[key] = count
	}
	return dest
}

func (a *aggregate) snapshot(duration time.Duration) Summary {
	s := Summary{Attempts: a.attempts, Successes: a.successes, Errors: a.attempts - a.successes,
		ErrorClasses: copyCounts(a.errors), ByOperation: make(map[string]OperationSummary),
		Latency: a.all.snapshot(), SuccessfulLatency: a.success.snapshot(), FailedLatency: a.failure.snapshot(),
		DurationSeconds: duration.Seconds()}
	if duration > 0 {
		s.Throughput = float64(a.attempts) / duration.Seconds()
	}
	if a.attempts > 0 {
		s.SuccessRate = float64(a.successes) / float64(a.attempts)
	}
	for name, op := range a.ops {
		s.ByOperation[name] = OperationSummary{op.attempts, op.successes, op.attempts - op.successes, copyCounts(op.errors), op.all.snapshot(), op.success.snapshot(), op.failure.snapshot()}
	}
	return s
}

// Collector retains no individual operation samples. Window history and its
// undrained queue are each capped at 3,600 windows (one hour at one-second windows).
type Collector struct {
	mu               sync.Mutex
	started, ended   time.Time
	phase            string
	phaseStarted     time.Time
	durations        map[string]time.Duration
	phases           map[string]*aggregate
	totals           *aggregate
	transitions      []PhaseTransition
	window           Window
	windows, pending []Window
	dropped          uint64
}

func NewCollector(start time.Time) *Collector {
	return &Collector{started: start, phase: "baseline", phaseStarted: start,
		durations: make(map[string]time.Duration), phases: map[string]*aggregate{"baseline": newAggregate()},
		totals: newAggregate(), transitions: []PhaseTransition{{"baseline", start}},
		window: Window{StartedAt: start, Phases: make(map[string]uint64)}}
}

func (c *Collector) Phase() string { c.mu.Lock(); defer c.mu.Unlock(); return c.phase }

func (c *Collector) SetPhase(name string, at time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.ended.IsZero() || name == c.phase || name == "" {
		return
	}
	if at.Before(c.phaseStarted) {
		at = c.phaseStarted
	}
	c.durations[c.phase] += at.Sub(c.phaseStarted)
	c.phase, c.phaseStarted = name, at
	if c.phases[name] == nil {
		c.phases[name] = newAggregate()
	}
	c.transitions = append(c.transitions, PhaseTransition{name, at})
}

// Record records a completed operation. The caller captures Phase before starting
// its request, so requests crossing a phase boundary remain attributed correctly.
func (c *Collector) Record(phase, operation string, elapsed time.Duration, err error) {
	c.RecordAt(phase, operation, elapsed, err, time.Now())
}

// RecordAt permits deterministic clocks in tests. Windows use completion time.
func (c *Collector) RecordAt(phase, operation string, elapsed time.Duration, err error, at time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.ended.IsZero() {
		return
	}
	c.advanceWindows(at)
	a := c.phases[phase]
	if a == nil {
		a = newAggregate()
		c.phases[phase] = a
	}
	a.record(operation, elapsed, err)
	c.totals.record(operation, elapsed, err)
	c.window.Attempts++
	c.window.Phases[phase]++
	if err == nil {
		c.window.Successes++
	} else {
		c.window.Errors++
	}
}

func (c *Collector) appendWindow(w Window) {
	if w.Attempts > 0 {
		w.SuccessRate = float64(w.Successes) / float64(w.Attempts)
	}
	if len(c.windows) == maxWindows {
		c.windows = c.windows[1:]
		c.dropped++
	}
	if len(c.pending) == maxWindows {
		c.pending = c.pending[1:]
	}
	c.windows = append(c.windows, w)
	c.pending = append(c.pending, w)
}

func (c *Collector) advanceWindows(at time.Time) {
	for !at.Before(c.window.StartedAt.Add(time.Second)) {
		c.window.EndedAt = c.window.StartedAt.Add(time.Second)
		c.appendWindow(c.window)
		c.window = Window{StartedAt: c.window.EndedAt, Phases: make(map[string]uint64)}
	}
}

func cloneWindows(windows []Window) []Window {
	out := make([]Window, len(windows))
	for i, window := range windows {
		out[i] = window
		out[i].Phases = copyCounts(window.Phases)
	}
	return out
}

func (c *Collector) DrainWindows(at time.Time) []Window {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.ended.IsZero() {
		c.advanceWindows(at)
	}
	out := cloneWindows(c.pending)
	c.pending = nil
	return out
}

// Close freezes the clock after all in-flight operations have been recorded.
func (c *Collector) Close(at time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.ended.IsZero() {
		return
	}
	if at.Before(c.phaseStarted) {
		at = c.phaseStarted
	}
	c.advanceWindows(at)
	if at.After(c.window.StartedAt) {
		c.window.EndedAt = at
		c.appendWindow(c.window)
	}
	c.durations[c.phase] += at.Sub(c.phaseStarted)
	c.ended = at
}

func (c *Collector) Snapshot(at time.Time) Snapshot {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.ended.IsZero() {
		at = c.ended
	}
	if at.Before(c.started) {
		at = c.started
	}
	s := Snapshot{StartedAt: c.started, EndedAt: at, Phases: make(map[string]PhaseSummary),
		Totals: c.totals.snapshot(at.Sub(c.started)), Windows: cloneWindows(c.windows),
		DroppedWindows: c.dropped, PhaseTransitions: append([]PhaseTransition(nil), c.transitions...), Methodology: Methodology}
	for name, agg := range c.phases {
		d := c.durations[name]
		if c.ended.IsZero() && name == c.phase && at.After(c.phaseStarted) {
			d += at.Sub(c.phaseStarted)
		}
		s.Phases[name] = agg.snapshot(d)
	}
	return s
}

// ClassifyError describes terminal application-visible outcomes, after the
// cluster client has exhausted its configured retries/redirection handling.
func ClassifyError(err error) string {
	if err == nil {
		return ""
	}
	if errors.Is(err, redis.Nil) || errors.Is(err, ErrMissing) {
		return "missing"
	}
	if errors.Is(err, ErrMismatch) {
		return "mismatch"
	}
	if errors.Is(err, ErrWaitShort) {
		return "wait_short"
	}
	if errors.Is(err, ErrWaitAOFShort) {
		return "waitaof_short"
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return "timeout"
	}
	if errors.Is(err, context.Canceled) {
		return "canceled"
	}
	var ne net.Error
	if errors.As(err, &ne) {
		if ne.Timeout() {
			return "timeout"
		}
		return "network"
	}
	message := strings.ToUpper(err.Error())
	for _, class := range []string{"CLUSTERDOWN", "MOVED", "ASK", "READONLY", "LOADING", "TRYAGAIN", "MASTERDOWN", "NOAUTH", "WRONGPASS"} {
		if strings.HasPrefix(message, class+" ") || message == class {
			return class
		}
	}
	// go-redis's connection pool deadline uses a plain error rather than a
	// net.Error, so classify that bounded wait explicitly.
	if strings.Contains(message, "CONNECTION POOL TIMEOUT") {
		return "timeout"
	}
	for _, fragment := range []string{"EOF", "CONNECTION REFUSED", "CONNECTION RESET", "BROKEN PIPE", "NETWORK IS UNREACHABLE", "CLOSED NETWORK CONNECTION"} {
		if strings.Contains(message, fragment) {
			return "network"
		}
	}
	return "other"
}
