// Package workload generates a bounded workload over an injected Redis client.
// Writes use distinct immutable keys. Only SETs returning success enter the
// acknowledged ledger; a failed SET may still have executed on Redis.
package workload

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Alaneel/redis-reliability-lab/internal/metrics"
	"github.com/redis/go-redis/v9"
)

// Client is satisfied by redis.ClusterClient. ContextTimeoutEnabled must be
// enabled on that client so Redis network I/O honors each operation deadline.
type Client interface {
	Set(context.Context, string, interface{}, time.Duration) *redis.StatusCmd
	Get(context.Context, string) *redis.StringCmd
}

type Config struct {
	WritePolicy      string        `json:"write_policy"`
	Workers          int           `json:"workers"`
	RatePerSecond    float64       `json:"rate_per_second"`
	MaxWrites        int           `json:"max_writes"`
	ReadPercent      int           `json:"read_percent"`
	ValueBytes       int           `json:"value_bytes"`
	Seed             int64         `json:"seed"`
	OperationTimeout time.Duration `json:"operation_timeout"`
}

func (c Config) validate() error {
	if c.WritePolicy != "" && c.WritePolicy != "set" && c.WritePolicy != "wait" && c.WritePolicy != "waitaof" {
		return errors.New("workload write policy must be set, wait or waitaof")
	}
	if c.Workers < 1 || c.Workers > 1024 {
		return errors.New("workload workers must be between 1 and 1024")
	}
	if math.IsNaN(c.RatePerSecond) || math.IsInf(c.RatePerSecond, 0) || c.RatePerSecond <= 0 || c.RatePerSecond > 1e6 {
		return errors.New("workload rate must be finite and in (0, 1000000]")
	}
	if float64(time.Second)/c.RatePerSecond >= float64(math.MaxInt64) {
		return errors.New("workload rate interval exceeds the supported timer duration")
	}
	if c.MaxWrites <= 0 {
		return errors.New("workload maximum writes must be positive")
	}
	if c.ReadPercent < 0 || c.ReadPercent > 100 {
		return errors.New("workload read percentage must be between 0 and 100")
	}
	if c.ValueBytes < 1 || c.ValueBytes > 1<<20 {
		return errors.New("workload value bytes must be between 1 and 1048576")
	}
	if c.OperationTimeout <= 0 {
		return errors.New("workload operation timeout must be positive")
	}
	return nil
}

type Write struct {
	Sequence                       uint64    `json:"sequence"`
	Key                            string    `json:"key"`
	Value                          string    `json:"value"`
	Slot                           uint16    `json:"slot"`
	AcknowledgedAt                 time.Time `json:"acknowledged_at"`
	WaitAttempted                  bool      `json:"wait_attempted"`
	WaitRequested                  int       `json:"wait_requested"`
	WaitAcknowledgements           int64     `json:"wait_acknowledgements"`
	WaitAOFAttempted               bool      `json:"waitaof_attempted"`
	WaitAOFRequestedLocal          int       `json:"waitaof_requested_local"`
	WaitAOFRequestedReplicas       int       `json:"waitaof_requested_replicas"`
	WaitAOFLocalAcknowledgements   int64     `json:"waitaof_local_acknowledgements"`
	WaitAOFReplicaAcknowledgements int64     `json:"waitaof_replica_acknowledgements"`
	PolicySatisfied                bool      `json:"policy_satisfied"`
}

type Stats struct {
	WritePolicy        string                  `json:"write_policy"`
	PolicyTotals       PolicyCounts            `json:"policy_totals"`
	PolicyByPhase      map[string]PolicyCounts `json:"policy_by_phase"`
	WriteStages        metrics.Snapshot        `json:"write_stages"`
	OfferedRequests    uint64                  `json:"offered_requests"`
	StartedRequests    uint64                  `json:"started_requests"`
	CompletedRequests  uint64                  `json:"completed_requests"`
	DroppedRequests    uint64                  `json:"dropped_requests"`
	PacingMethodology  string                  `json:"pacing_methodology"`
	WritesAttempted    uint64                  `json:"writes_attempted"`
	AcknowledgedWrites uint64                  `json:"acknowledged_writes"`
	AmbiguousWrites    uint64                  `json:"ambiguous_writes"`
	ReadChecks         uint64                  `json:"read_checks"`
	MissingReads       uint64                  `json:"missing_reads"`
	MismatchedReads    uint64                  `json:"mismatched_reads"`
	WriteLimitReached  bool                    `json:"write_limit_reached"`
	Guarantee          string                  `json:"guarantee"`
}

const Guarantee = "Every successful client SET acknowledgement is validated, including writes whose subsequent WAIT or WAITAOF returned too few acknowledgements or an error. Every write uses a unique immutable key. SET errors are ambiguous outcomes and may have executed; they are excluded from the acknowledged-write guarantee. WAIT-satisfied and WAITAOF-satisfied subsets are reported separately. WAIT concerns replica memory, while WAITAOF requires observed local and replica AOF counts; neither subset establishes loss-free failover or a power-loss guarantee from this process-crash experiment. Redis asynchronous replication can lose acknowledged writes during failover. This check establishes observed post-experiment readability and value equality, not a general durability guarantee."

const PacingMethodology = "Offered requests are ideal period slots from workload Start until the dispatcher observes Stop or parent cancellation; the first slot occurs after one period. The period is one second divided by the configured aggregate rate, truncated to whole nanoseconds. A central dispatcher offers only the latest due slot over an unbuffered channel; missed slots, busy workers, stopped tokens, and tokens with no eligible operation are dropped, without catch-up requests. Started requests count actual logical workload GET or write executions; a SET+WAIT/WAITAOF write counts once. Completed requests count their terminal logical metrics attempts. After Stop, offered = started + dropped and started = completed. Validation and ordinary/acknowledgement-policy probes are excluded. These global counters are not attributed to phases."

type Runner struct {
	client           Client
	writer           WriteExecutor
	config           Config
	runID            string
	tags             [3]string
	mu               sync.Mutex
	started, stopped bool
	collector        *metrics.Collector
	stages           *metrics.Collector
	policyMu         sync.Mutex
	policyTotals     PolicyCounts
	policyByPhase    map[string]PolicyCounts
	stop             chan struct{}
	done             chan struct{}
	stopOnce         sync.Once
	wg               sync.WaitGroup
	dispatchDone     chan struct{}
	ledgerMu         sync.Mutex
	ledger           []Write
	allocated        atomic.Uint64
	ambiguous        atomic.Uint64
	reads            atomic.Uint64
	missing          atomic.Uint64
	mismatched       atomic.Uint64
	offered          atomic.Uint64
	requestsStarted  atomic.Uint64
	completed        atomic.Uint64
	dropped          atomic.Uint64
}

func New(client Client, config Config, runID string) (*Runner, error) {
	if config.WritePolicy == "wait" || config.WritePolicy == "waitaof" {
		return nil, errors.New("wait or waitaof policy requires NewWithWriter and a connection-affine write executor")
	}
	return NewWithWriter(client, normalWriter{client: client}, config, runID)
}

func NewWithWriter(client Client, writer WriteExecutor, config Config, runID string) (*Runner, error) {
	if client == nil {
		return nil, errors.New("workload client is required")
	}
	if writer == nil {
		return nil, errors.New("workload write executor is required")
	}
	if err := config.validate(); err != nil {
		return nil, err
	}
	if runID == "" {
		return nil, errors.New("workload run ID is required")
	}
	if config.WritePolicy == "" {
		config.WritePolicy = "set"
	}
	started := time.Now()
	r := &Runner{client: client, writer: writer, config: config, runID: runID, stop: make(chan struct{}), done: make(chan struct{}), dispatchDone: make(chan struct{}), collector: metrics.NewCollector(started), stages: metrics.NewCollector(started), policyTotals: newPolicyCounts(), policyByPhase: make(map[string]PolicyCounts)}
	// The managed three-primary cluster starts with contiguous slot thirds.
	// Tags guarantee equal attempted write coverage of those three slot ranges.
	for i := 0; i < 3; i++ {
		for candidate := 0; ; candidate++ {
			tag := fmt.Sprintf("redis-lab-shard-%d-%d", i, candidate)
			slot := HashSlot(tag)
			minimum := [3]uint16{0, 5461, 10923}
			maximum := [3]uint16{5460, 10922, 16383}
			if slot >= minimum[i] && slot <= maximum[i] {
				r.tags[i] = tag
				break
			}
		}
	}
	return r, nil
}

// Start starts workers once. One dispatcher accounts for the ideal aggregate
// schedule, dropping offers when a worker cannot start them immediately.
func (r *Runner) Start(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.started || r.stopped {
		return errors.New("workload runner can only be started once")
	}
	r.started = true
	started := time.Now()
	r.collector = metrics.NewCollector(started)
	r.stages = metrics.NewCollector(started)
	period := time.Duration(float64(time.Second) / r.config.RatePerSecond)
	if period < time.Nanosecond {
		period = time.Nanosecond
	}
	tokens := make(chan struct{})
	for worker := 0; worker < r.config.Workers; worker++ {
		r.wg.Add(1)
		go r.worker(ctx, tokens, worker)
	}
	r.wg.Add(1)
	go r.dispatch(ctx, tokens, started, period)
	go func() {
		r.wg.Wait()
		ended := time.Now()
		r.collector.Close(ended)
		r.stages.Close(ended)
		r.mu.Lock()
		r.stopped = true
		r.mu.Unlock()
		close(r.done)
	}()
	return nil
}

// offerSlots accounts for every ideal slot elapsed since the preceding wakeup.
// It returns at most one dispatchable offer, regardless of scheduler delay.
func (r *Runner) offerSlots(started, at time.Time, period time.Duration, last *uint64, dispatch bool) bool {
	elapsed := at.Sub(started)
	if elapsed <= 0 {
		return false
	}
	due := uint64(elapsed / period)
	if due <= *last {
		return false
	}
	n := due - *last
	*last = due
	r.offered.Add(n)
	if !dispatch {
		r.dropped.Add(n)
		return false
	}
	r.dropped.Add(n - 1)
	return true
}

func (r *Runner) dispatch(ctx context.Context, tokens chan<- struct{}, started time.Time, period time.Duration) {
	defer r.wg.Done()
	defer close(r.dispatchDone)
	defer close(tokens)
	ticker := time.NewTicker(period)
	defer ticker.Stop()
	var last uint64
	for {
		select {
		case <-r.stop:
			r.offerSlots(started, time.Now(), period, &last, false)
			return
		case <-ctx.Done():
			r.offerSlots(started, time.Now(), period, &last, false)
			return
		case <-ticker.C:
		}
		// Use the current monotonic elapsed time, not a possibly stale tick.
		// Go timers may skip ticks or the scheduler may wake up late.
		stopping := ctx.Err() != nil
		select {
		case <-r.stop:
			stopping = true
		default:
		}
		if !r.offerSlots(started, time.Now(), period, &last, !stopping) {
			if stopping {
				return
			}
			continue
		}
		select {
		case tokens <- struct{}{}:
		default:
			r.dropped.Add(1)
		}
	}
}

func (r *Runner) worker(ctx context.Context, tokens <-chan struct{}, worker int) {
	defer r.wg.Done()
	rng := rand.New(rand.NewSource(r.config.Seed + int64(worker)*1_000_003))
	var acknowledged []Write
	for {
		// Prioritize stop before accepting another rate token.
		select {
		case <-r.stop:
			return
		case <-ctx.Done():
			return
		default:
		}
		select {
		case <-r.stop:
			return
		case <-ctx.Done():
			return
		case _, ok := <-tokens:
			if !ok {
				return
			}
		}
		select {
		case <-r.stop:
			r.dropped.Add(1)
			return
		case <-ctx.Done():
			r.dropped.Add(1)
			return
		default:
		}
		if len(acknowledged) > 0 && (rng.Intn(100) < r.config.ReadPercent || r.allocated.Load() >= uint64(r.config.MaxWrites)) {
			r.read(ctx, acknowledged[rng.Intn(len(acknowledged))])
			continue
		}
		seq, ok := r.reserveWrite()
		if !ok {
			r.dropped.Add(1)
			continue
		}
		write := r.makeWrite(seq)
		phase := r.collector.Phase()
		started := time.Now()
		opCtx, cancel := context.WithTimeout(ctx, r.config.OperationTimeout)
		r.requestsStarted.Add(1)
		result := r.writer.Write(opCtx, write.Key, write.Value)
		cancel()
		elapsed := time.Since(started)
		if result.TotalLatency > 0 {
			elapsed = result.TotalLatency
		}
		err := logicalWriteError(r.config.WritePolicy, result)
		operation := "set"
		if r.config.WritePolicy == "wait" {
			operation = "set_wait"
		}
		if r.config.WritePolicy == "waitaof" {
			operation = "set_waitaof"
		}
		r.collector.Record(phase, operation, elapsed, err)
		r.recordWriteResult(phase, result)
		r.completed.Add(1)
		if !result.SetAcknowledged {
			r.ambiguous.Add(1)
			continue
		}
		// Do this even if Stop was requested while SET was in flight: a client
		// success must always be represented in the final validation ledger.
		write.AcknowledgedAt = result.SetAcknowledgedAt
		if write.AcknowledgedAt.IsZero() {
			// Compatibility for injected test executors. Production executors
			// record acknowledgement immediately after SET and before the policy command.
			write.AcknowledgedAt = time.Now()
		}
		write.WaitAttempted, write.WaitRequested, write.WaitAcknowledgements, write.PolicySatisfied = result.WaitAttempted, result.WaitRequested, result.WaitAcknowledgements, err == nil
		write.WaitAOFAttempted, write.WaitAOFRequestedLocal, write.WaitAOFRequestedReplicas = result.WaitAOFAttempted, result.WaitAOFRequestedLocal, result.WaitAOFRequestedReplicas
		write.WaitAOFLocalAcknowledgements, write.WaitAOFReplicaAcknowledgements = result.WaitAOFLocalAcknowledgements, result.WaitAOFReplicaAcknowledgements
		r.ledgerMu.Lock()
		r.ledger = append(r.ledger, write)
		r.ledgerMu.Unlock()
		acknowledged = append(acknowledged, write)
	}
}

func (r *Runner) reserveWrite() (uint64, bool) {
	for {
		n := r.allocated.Load()
		if n >= uint64(r.config.MaxWrites) {
			return 0, false
		}
		if r.allocated.CompareAndSwap(n, n+1) {
			return n, true
		}
	}
}

func (r *Runner) makeWrite(seq uint64) Write {
	key := fmt.Sprintf("redis-lab:{%s}:%s:%d:%d", r.tags[seq%3], r.runID, r.config.Seed, seq)
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s/%d/%d", r.runID, r.config.Seed, seq)))
	encoded := hex.EncodeToString(sum[:])
	value := strings.Repeat(encoded, (r.config.ValueBytes+len(encoded)-1)/len(encoded))[:r.config.ValueBytes]
	return Write{Sequence: seq, Key: key, Value: value, Slot: HashSlot(key)}
}

func (r *Runner) read(ctx context.Context, write Write) {
	phase := r.collector.Phase()
	started := time.Now()
	opCtx, cancel := context.WithTimeout(ctx, r.config.OperationTimeout)
	r.requestsStarted.Add(1)
	actual, err := r.client.Get(opCtx, write.Key).Result()
	cancel()
	r.reads.Add(1)
	if errors.Is(err, redis.Nil) {
		r.missing.Add(1)
	}
	if err == nil && actual != write.Value {
		err = metrics.ErrMismatch
		r.mismatched.Add(1)
	}
	r.collector.Record(phase, "get", time.Since(started), err)
	r.completed.Add(1)
}

// Stop ends scheduling and waits for current requests to finish within their
// operation deadlines. It intentionally does not cancel already-issued requests.
func (r *Runner) Stop() {
	r.stopOnce.Do(func() { close(r.stop) })
	r.mu.Lock()
	if !r.started {
		r.stopped = true
		r.collector.Close(time.Now())
		r.stages.Close(time.Now())
		r.mu.Unlock()
		return
	}
	r.mu.Unlock()
	<-r.done
}

func (r *Runner) SetPhase(name string, at time.Time) {
	r.mu.Lock()
	collector, stages := r.collector, r.stages
	r.mu.Unlock()
	collector.SetPhase(name, at)
	stages.SetPhase(name, at)
}

func (r *Runner) Metrics() *metrics.Collector { r.mu.Lock(); defer r.mu.Unlock(); return r.collector }

func (r *Runner) Snapshot() metrics.Snapshot { return r.Metrics().Snapshot(time.Now()) }

func (r *Runner) Ledger() []Write {
	r.ledgerMu.Lock()
	out := append([]Write(nil), r.ledger...)
	r.ledgerMu.Unlock()
	sort.Slice(out, func(i, j int) bool { return out[i].Sequence < out[j].Sequence })
	return out
}

func (r *Runner) Stats() Stats {
	r.ledgerMu.Lock()
	ack := uint64(len(r.ledger))
	r.ledgerMu.Unlock()
	r.policyMu.Lock()
	totals := clonePolicyCounts(r.policyTotals)
	byPhase := make(map[string]PolicyCounts, len(r.policyByPhase))
	for phase, counts := range r.policyByPhase {
		byPhase[phase] = clonePolicyCounts(counts)
	}
	r.policyMu.Unlock()
	r.mu.Lock()
	stages := r.stages
	r.mu.Unlock()
	stageSnapshot := stages.Snapshot(time.Now())
	stageSnapshot.Methodology = WriteStagesMethodology + " " + stageSnapshot.Methodology
	return Stats{WritePolicy: r.config.WritePolicy, PolicyTotals: totals, PolicyByPhase: byPhase, WriteStages: stageSnapshot, WritesAttempted: r.allocated.Load(), AcknowledgedWrites: ack, AmbiguousWrites: r.ambiguous.Load(),
		OfferedRequests: r.offered.Load(), StartedRequests: r.requestsStarted.Load(), CompletedRequests: r.completed.Load(), DroppedRequests: r.dropped.Load(), PacingMethodology: PacingMethodology,
		ReadChecks: r.reads.Load(), MissingReads: r.missing.Load(), MismatchedReads: r.mismatched.Load(),
		WriteLimitReached: r.allocated.Load() >= uint64(r.config.MaxWrites), Guarantee: Guarantee}
}

type ValidationSample struct {
	Key     string `json:"key"`
	Outcome string `json:"outcome"`
	Error   string `json:"error,omitempty"`
}

type Validation struct {
	WaitSatisfiedSubset    ValidationSubset   `json:"wait_satisfied_subset"`
	WaitAOFSatisfiedSubset ValidationSubset   `json:"waitaof_satisfied_subset"`
	StartedAt              time.Time          `json:"started_at"`
	CompletedAt            time.Time          `json:"completed_at"`
	Complete               bool               `json:"complete"`
	Expected               int                `json:"expected"`
	Checked                int                `json:"checked"`
	Missing                int                `json:"missing"`
	Corrupt                int                `json:"corrupt"`
	Unreadable             int                `json:"unreadable"`
	Unverified             int                `json:"unverified"`
	Samples                []ValidationSample `json:"samples"`
	Error                  string             `json:"error,omitempty"`
	Guarantee              string             `json:"guarantee"`
}

// Validate checks every acknowledged immutable write after the workload stops.
// The caller supplies a bounded overall context; each GET has its own timeout.
// Timeout/cancellation leaves remaining entries explicitly unverified.
func (r *Runner) Validate(ctx context.Context, timeoutPerOperation time.Duration) Validation {
	out := Validation{StartedAt: time.Now(), Guarantee: Guarantee, Samples: make([]ValidationSample, 0)}
	r.mu.Lock()
	stopped := r.stopped
	r.mu.Unlock()
	ledger := r.Ledger()
	out.Expected = len(ledger)
	for _, write := range ledger {
		if waitSatisfiedWrite(write) {
			out.WaitSatisfiedSubset.Expected++
		}
		if waitAOFSatisfiedWrite(write) {
			out.WaitAOFSatisfiedSubset.Expected++
		}
	}
	if !stopped || timeoutPerOperation <= 0 {
		out.Error = "validation requires a stopped workload and a positive operation timeout"
		out.Unverified = out.Expected
		out.WaitSatisfiedSubset.Unverified = out.WaitSatisfiedSubset.Expected
		out.WaitAOFSatisfiedSubset.Unverified = out.WaitAOFSatisfiedSubset.Expected
		out.CompletedAt = time.Now()
		return out
	}
	for _, write := range ledger {
		if ctx.Err() != nil {
			out.Error = ctx.Err().Error()
			break
		}
		opCtx, cancel := context.WithTimeout(ctx, timeoutPerOperation)
		actual, err := r.client.Get(opCtx, write.Key).Result()
		cancel()
		out.Checked++
		waitSatisfied := waitSatisfiedWrite(write)
		waitAOFSatisfied := waitAOFSatisfiedWrite(write)
		if waitSatisfied {
			out.WaitSatisfiedSubset.Checked++
		}
		if waitAOFSatisfied {
			out.WaitAOFSatisfiedSubset.Checked++
		}
		outcome := ""
		switch {
		case errors.Is(err, redis.Nil):
			out.Missing++
			if waitSatisfied {
				out.WaitSatisfiedSubset.Missing++
			}
			if waitAOFSatisfied {
				out.WaitAOFSatisfiedSubset.Missing++
			}
			outcome = "missing"
		case err != nil:
			out.Unreadable++
			if waitSatisfied {
				out.WaitSatisfiedSubset.Unreadable++
			}
			if waitAOFSatisfied {
				out.WaitAOFSatisfiedSubset.Unreadable++
			}
			outcome = "unreadable"
		case actual != write.Value:
			out.Corrupt++
			if waitSatisfied {
				out.WaitSatisfiedSubset.Corrupt++
			}
			if waitAOFSatisfied {
				out.WaitAOFSatisfiedSubset.Corrupt++
			}
			outcome = "corrupt"
		}
		if outcome != "" && len(out.Samples) < 10 {
			sample := ValidationSample{Key: write.Key, Outcome: outcome}
			if err != nil && !errors.Is(err, redis.Nil) {
				sample.Error = err.Error()
			}
			out.Samples = append(out.Samples, sample)
		}
	}
	out.Unverified = out.Expected - out.Checked
	out.Complete = out.Unverified == 0
	out.WaitSatisfiedSubset.Unverified = out.WaitSatisfiedSubset.Expected - out.WaitSatisfiedSubset.Checked
	out.WaitSatisfiedSubset.Complete = out.WaitSatisfiedSubset.Unverified == 0
	out.WaitAOFSatisfiedSubset.Unverified = out.WaitAOFSatisfiedSubset.Expected - out.WaitAOFSatisfiedSubset.Checked
	out.WaitAOFSatisfiedSubset.Complete = out.WaitAOFSatisfiedSubset.Unverified == 0
	out.CompletedAt = time.Now()
	return out
}

// HashSlot implements Redis Cluster's CRC16/XMODEM hash tags and 16,384 slots.
func HashSlot(key string) uint16 {
	if start := strings.IndexByte(key, '{'); start >= 0 {
		if end := strings.IndexByte(key[start+1:], '}'); end > 0 {
			key = key[start+1 : start+1+end]
		}
	}
	var crc uint16
	for _, b := range []byte(key) {
		crc ^= uint16(b) << 8
		for bit := 0; bit < 8; bit++ {
			if crc&0x8000 != 0 {
				crc = (crc << 1) ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}
	return crc % 16384
}
