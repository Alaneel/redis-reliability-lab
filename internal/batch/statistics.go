package batch

import (
	"math"
	"sort"
)

const minimumPercentileSamples = 20

// Statistics describes trial-level observations, never pooled request latencies.
// Sample standard deviation is undefined for N<2. High percentiles are omitted
// for N<20; even larger samples describe only the executed trials.
type Statistics struct {
	N            int      `json:"n"`
	Min          *float64 `json:"min,omitempty"`
	Median       *float64 `json:"median,omitempty"`
	Mean         *float64 `json:"mean,omitempty"`
	SampleStdDev *float64 `json:"sample_stddev,omitempty"`
	P90          *float64 `json:"p90,omitempty"`
	P95          *float64 `json:"p95,omitempty"`
	Max          *float64 `json:"max,omitempty"`
	SmallSample  bool     `json:"small_sample"`
	Methodology  string   `json:"methodology"`
}

func number(value float64) *float64 { return &value }

func describe(values []float64) Statistics {
	s := Statistics{N: len(values), SmallSample: len(values) < minimumPercentileSamples,
		Methodology: "Descriptive statistics over available trial observations. Standard deviation uses N-1. P90/P95 use nearest rank and require at least 20 observations. Small samples do not support reliable tail estimates; these statistics are not confidence intervals or universal reliability guarantees."}
	if len(values) == 0 {
		return s
	}
	ordered := append([]float64(nil), values...)
	sort.Float64s(ordered)
	var mean, m2 float64
	for i, value := range ordered {
		delta := value - mean
		mean += delta / float64(i+1)
		m2 += delta * (value - mean)
	}
	median := ordered[len(ordered)/2]
	if len(ordered)%2 == 0 {
		median = ordered[len(ordered)/2-1]/2 + median/2
	}
	s.Min, s.Median, s.Mean, s.Max = number(ordered[0]), number(median), number(mean), number(ordered[len(ordered)-1])
	if len(ordered) >= 2 {
		s.SampleStdDev = number(math.Sqrt(math.Max(0, m2/float64(len(ordered)-1))))
	}
	if len(ordered) >= minimumPercentileSamples {
		s.P90 = number(ordered[int(math.Ceil(.90*float64(len(ordered))))-1])
		s.P95 = number(ordered[int(math.Ceil(.95*float64(len(ordered))))-1])
	}
	return s
}

func finite(value float64) bool { return !math.IsNaN(value) && !math.IsInf(value, 0) }
