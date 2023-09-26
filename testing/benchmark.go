// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

import (
	"github.com/shogo82148/std/time"
)

// Global lock to ensure only one benchmark runs at a time.

// Used for every benchmark for measuring memory.

// InternalBenchmark is an internal type but exported because it is cross-package;
// it is part of the implementation of the "go test" command.
type InternalBenchmark struct {
	Name string
	F    func(b *B)
}

// B is a type passed to Benchmark functions to manage benchmark
// timing and to specify the number of iterations to run.
//
// A benchmark ends when its Benchmark function returns or calls any of the methods
// FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. Those methods must be called
// only from the goroutine running the Benchmark function.
// The other reporting methods, such as the variations of Log and Error,
// may be called simultaneously from multiple goroutines.
//
// Like in tests, benchmark logs are accumulated during execution
// and dumped to standard output when done. Unlike in tests, benchmark logs
// are always printed, so as not to hide output whose existence may be
// affecting benchmark results.
type B struct {
	common
	importPath       string
	context          *benchContext
	N                int
	previousN        int
	previousDuration time.Duration
	benchFunc        func(b *B)
	benchTime        durationOrCountFlag
	bytes            int64
	missingBytes     bool
	timerOn          bool
	showAllocResult  bool
	result           BenchmarkResult
	parallelism      int

	startAllocs uint64
	startBytes  uint64

	netAllocs uint64
	netBytes  uint64

	extra map[string]float64
}

// StartTimer starts timing a test. This function is called automatically
// before a benchmark starts, but it can also be used to resume timing after
// a call to StopTimer.
func (b *B) StartTimer()

// StopTimer stops timing a test. This can be used to pause the timer
// while performing complex initialization that you don't
// want to measure.
func (b *B) StopTimer()

// ResetTimer zeroes the elapsed benchmark time and memory allocation counters
// and deletes user-reported metrics.
// It does not affect whether the timer is running.
func (b *B) ResetTimer()

// SetBytes records the number of bytes processed in a single operation.
// If this is called, the benchmark will report ns/op and MB/s.
func (b *B) SetBytes(n int64)

// ReportAllocs enables malloc statistics for this benchmark.
// It is equivalent to setting -test.benchmem, but it only affects the
// benchmark function that calls ReportAllocs.
func (b *B) ReportAllocs()

// ReportMetric adds "n unit" to the reported benchmark results.
// If the metric is per-iteration, the caller should divide by b.N,
// and by convention units should end in "/op".
// ReportMetric overrides any previously reported value for the same unit.
// ReportMetric panics if unit is the empty string or if unit contains
// any whitespace.
// If unit is a unit normally reported by the benchmark framework itself
// (such as "allocs/op"), ReportMetric will override that metric.
// Setting "ns/op" to 0 will suppress that built-in metric.
func (b *B) ReportMetric(n float64, unit string)

// BenchmarkResult contains the results of a benchmark run.
type BenchmarkResult struct {
	N         int
	T         time.Duration
	Bytes     int64
	MemAllocs uint64
	MemBytes  uint64

	Extra map[string]float64
}

// NsPerOp returns the "ns/op" metric.
func (r BenchmarkResult) NsPerOp() int64

// AllocsPerOp returns the "allocs/op" metric,
// which is calculated as r.MemAllocs / r.N.
func (r BenchmarkResult) AllocsPerOp() int64

// AllocedBytesPerOp returns the "B/op" metric,
// which is calculated as r.MemBytes / r.N.
func (r BenchmarkResult) AllocedBytesPerOp() int64

// String returns a summary of the benchmark results.
// It follows the benchmark result line format from
// https://golang.org/design/14313-benchmark-format, not including the
// benchmark name.
// Extra metrics override built-in metrics of the same name.
// String does not include allocs/op or B/op, since those are reported
// by MemString.
func (r BenchmarkResult) String() string

// MemString returns r.AllocedBytesPerOp and r.AllocsPerOp in the same format as 'go test'.
func (r BenchmarkResult) MemString() string

// RunBenchmarks is an internal function but exported because it is cross-package;
// it is part of the implementation of the "go test" command.
func RunBenchmarks(matchString func(pat, str string) (bool, error), benchmarks []InternalBenchmark)

// Run benchmarks f as a subbenchmark with the given name. It reports
// whether there were any failures.
//
// A subbenchmark is like any other benchmark. A benchmark that calls Run at
// least once will not be measured itself and will be called once with N=1.
func (b *B) Run(name string, f func(b *B)) bool

// A PB is used by RunParallel for running parallel benchmarks.
type PB struct {
	globalN *uint64
	grain   uint64
	cache   uint64
	bN      uint64
}

// Next reports whether there are more iterations to execute.
func (pb *PB) Next() bool

// RunParallel runs a benchmark in parallel.
// It creates multiple goroutines and distributes b.N iterations among them.
// The number of goroutines defaults to GOMAXPROCS. To increase parallelism for
// non-CPU-bound benchmarks, call SetParallelism before RunParallel.
// RunParallel is usually used with the go test -cpu flag.
//
// The body function will be run in each goroutine. It should set up any
// goroutine-local state and then iterate until pb.Next returns false.
// It should not use the StartTimer, StopTimer, or ResetTimer functions,
// because they have global effect. It should also not call Run.
func (b *B) RunParallel(body func(*PB))

// SetParallelism sets the number of goroutines used by RunParallel to p*GOMAXPROCS.
// There is usually no need to call SetParallelism for CPU-bound benchmarks.
// If p is less than 1, this call will have no effect.
func (b *B) SetParallelism(p int)

// Benchmark benchmarks a single function. It is useful for creating
// custom benchmarks that do not use the "go test" command.
//
// If f depends on testing flags, then Init must be used to register
// those flags before calling Benchmark and before calling flag.Parse.
//
// If f calls Run, the result will be an estimate of running all its
// subbenchmarks that don't call Run in sequence in a single benchmark.
func Benchmark(f func(b *B)) BenchmarkResult
