// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

import (
	"github.com/shogo82148/std/time"
)

// Global lock to ensure only one benchmark runs at a time.

// Used for every benchmark for measuring memory.

// An internal type but exported because it is cross-package; part of the implementation
// of the "go test" command.
type InternalBenchmark struct {
	Name string
	F    func(b *B)
}

// B is a type passed to Benchmark functions to manage benchmark
// timing and to specify the number of iterations to run.
type B struct {
	common
	N                int
	previousN        int
	previousDuration time.Duration
	benchmark        InternalBenchmark
	bytes            int64
	timerOn          bool
	showAllocResult  bool
	result           BenchmarkResult
	parallelism      int

	startAllocs uint64
	startBytes  uint64

	netAllocs uint64
	netBytes  uint64
}

// StartTimer starts timing a test.  This function is called automatically
// before a benchmark starts, but it can also used to resume timing after
// a call to StopTimer.
func (b *B) StartTimer()

// StopTimer stops timing a test.  This can be used to pause the timer
// while performing complex initialization that you don't
// want to measure.
func (b *B) StopTimer()

// ResetTimer zeros the elapsed benchmark time and memory allocation counters.
// It does not affect whether the timer is running.
func (b *B) ResetTimer()

// SetBytes records the number of bytes processed in a single operation.
// If this is called, the benchmark will report ns/op and MB/s.
func (b *B) SetBytes(n int64)

// ReportAllocs enables malloc statistics for this benchmark.
// It is equivalent to setting -test.benchmem, but it only affects the
// benchmark function that calls ReportAllocs.
func (b *B) ReportAllocs()

// The results of a benchmark run.
type BenchmarkResult struct {
	N         int
	T         time.Duration
	Bytes     int64
	MemAllocs uint64
	MemBytes  uint64
}

func (r BenchmarkResult) NsPerOp() int64

func (r BenchmarkResult) AllocsPerOp() int64

func (r BenchmarkResult) AllocedBytesPerOp() int64

func (r BenchmarkResult) String() string

func (r BenchmarkResult) MemString() string

// An internal function but exported because it is cross-package; part of the implementation
// of the "go test" command.
func RunBenchmarks(matchString func(pat, str string) (bool, error), benchmarks []InternalBenchmark)

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
// because they have global effect.
func (b *B) RunParallel(body func(*PB))

// SetParallelism sets the number of goroutines used by RunParallel to p*GOMAXPROCS.
// There is usually no need to call SetParallelism for CPU-bound benchmarks.
// If p is less than 1, this call will have no effect.
func (b *B) SetParallelism(p int)

// Benchmark benchmarks a single function. Useful for creating
// custom benchmarks that do not use the "go test" command.
func Benchmark(f func(b *B)) BenchmarkResult
