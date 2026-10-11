package metrics

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestQuantilesAndPhaseWallTime(t *testing.T) {
	start := time.Unix(100, 0)
	c := NewCollector(start)
	for i := 1; i <= 100; i++ {
		c.RecordAt("baseline", "set", time.Duration(i)*time.Millisecond, nil, start.Add(time.Second))
	}
	c.SetPhase("fault", start.Add(2*time.Second))
	c.RecordAt("fault", "get", 500*time.Millisecond, errors.New("CLUSTERDOWN no slot"), start.Add(3*time.Second))
	c.Close(start.Add(4 * time.Second))
	s := c.Snapshot(start.Add(10 * time.Second))
	b := s.Phases["baseline"]
	if b.Attempts != 100 || b.Successes != 100 || b.DurationSeconds != 2 || b.Throughput != 50 {
		t.Fatalf("baseline summary: %+v", b)
	}
	for _, test := range []struct {
		name              string
		got, lower, upper float64
	}{
		{"p50", b.Latency.P50MS, 50, 52.51}, {"p95", b.Latency.P95MS, 95, 99.76}, {"p99", b.Latency.P99MS, 99, 100},
	} {
		if test.got < test.lower || test.got > test.upper {
			t.Errorf("%s=%v outside [%v,%v]", test.name, test.got, test.lower, test.upper)
		}
	}
	f := s.Phases["fault"]
	if f.Errors != 1 || f.ErrorClasses["CLUSTERDOWN"] != 1 || f.FailedLatency.Count != 1 || f.SuccessfulLatency.Count != 0 {
		t.Fatalf("failure summary: %+v", f)
	}
	if f.Latency.P50MS != 500 || s.Totals.DurationSeconds != 4 || s.Totals.SuccessfulLatency.Count != 100 || s.Totals.FailedLatency.Count != 1 {
		t.Fatalf("overall summary: %+v", s.Totals)
	}
}

func TestWindowsIncludeIdleIntervalsAndDrainOnce(t *testing.T) {
	start := time.Unix(100, 0)
	c := NewCollector(start)
	c.RecordAt("baseline", "get", time.Millisecond, nil, start.Add(100*time.Millisecond))
	windows := c.DrainWindows(start.Add(3 * time.Second))
	if len(windows) != 3 || windows[0].Attempts != 1 || windows[0].SuccessRate != 1 || windows[1].Attempts != 0 || windows[1].SuccessRate != 0 {
		t.Fatalf("windows: %+v", windows)
	}
	if again := c.DrainWindows(start.Add(3 * time.Second)); len(again) != 0 {
		t.Fatalf("drained twice: %+v", again)
	}
	windows[0].Phases["baseline"] = 99
	if c.Snapshot(start.Add(3 * time.Second)).Windows[0].Phases["baseline"] != 1 {
		t.Fatal("snapshot aliases drained maps")
	}
	c.Close(start.Add(3250 * time.Millisecond))
	last := c.DrainWindows(start.Add(20 * time.Second))
	if len(last) != 1 || last[0].EndedAt.Sub(last[0].StartedAt) != 250*time.Millisecond {
		t.Fatalf("partial closing window: %+v", last)
	}
}

func TestConcurrentRecordAndSnapshot(t *testing.T) {
	c := NewCollector(time.Now())
	var wg sync.WaitGroup
	for worker := 0; worker < 10; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 200; i++ {
				c.Record(c.Phase(), "get", time.Millisecond, nil)
				_ = c.Snapshot(time.Now())
			}
		}()
	}
	wg.Wait()
	c.Close(time.Now())
	s := c.Snapshot(time.Now())
	if s.Totals.Attempts != 2000 || s.Totals.Successes != 2000 {
		t.Fatalf("lost operations: %+v", s.Totals)
	}
	s.Phases["baseline"].ErrorClasses["other"] = 99
	if c.Snapshot(time.Now()).Phases["baseline"].ErrorClasses["other"] != 0 {
		t.Fatal("snapshot aliases collector maps")
	}
}

func TestClassifyTerminalErrors(t *testing.T) {
	for _, test := range []struct {
		err  error
		want string
	}{
		{redis.Nil, "missing"}, {ErrMissing, "missing"}, {fmt.Errorf("wrapped: %w", ErrMismatch), "mismatch"},
		{ErrWaitAOFShort, "waitaof_short"}, {fmt.Errorf("wrapped: %w", ErrWaitAOFShort), "waitaof_short"},
		{ErrWaitShort, "wait_short"}, {fmt.Errorf("wrapped: %w", ErrWaitShort), "wait_short"},
		{context.DeadlineExceeded, "timeout"}, {context.Canceled, "canceled"},
		{&net.OpError{Op: "dial", Net: "tcp", Err: errors.New("refused")}, "network"},
		{errors.New("EOF"), "network"}, {errors.New("MOVED 3999 127.0.0.1:7001"), "MOVED"},
		{errors.New("ASK 3999 127.0.0.1:7001"), "ASK"}, {errors.New("CLUSTERDOWN state fail"), "CLUSTERDOWN"},
		{errors.New("READONLY replica"), "READONLY"}, {errors.New("unknown"), "other"},
		{errors.New("redis: connection pool timeout"), "timeout"},
	} {
		if got := ClassifyError(test.err); got != test.want {
			t.Errorf("%v classified %q; want %q", test.err, got, test.want)
		}
	}
}

func TestBoundedWindowHistory(t *testing.T) {
	start := time.Unix(100, 0)
	c := NewCollector(start)
	windows := c.DrainWindows(start.Add((maxWindows + 3) * time.Second))
	s := c.Snapshot(start.Add((maxWindows + 3) * time.Second))
	if len(windows) != maxWindows || len(s.Windows) != maxWindows || s.DroppedWindows != 3 {
		t.Fatalf("unbounded windows: drained=%d snapshot=%d dropped=%d", len(windows), len(s.Windows), s.DroppedWindows)
	}
}
