// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/cmp"
	"github.com/shogo82148/std/slices"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/testing"
	"github.com/shogo82148/std/text/template"
)

func ExampleB_Loop() {
	simpleFunc := func(i int) int {
		return i + 1
	}
	n := 0
	testing.Benchmark(func(b *testing.B) {
		// Unlike "for i := range N {...}" style loops, this
		// setup logic will only be executed once, so simpleFunc
		// will always get argument 1.
		n++
		// It behaves just like "for i := range N {...}", except with keeping
		// function call parameters and results alive.
		for b.Loop() {
			// This function call, if was in a normal loop, will be optimized away
			// completely, first by inlining, then by dead code elimination.
			// In a b.Loop loop, the compiler ensures that this function is not optimized away.
			simpleFunc(n)
		}
		// This clean-up will only be executed once, so after the benchmark, the user
		// will see n == 2.
		n++
		// Use b.ReportMetric as usual just like what a user may do after
		// b.N loop.
	})
	// We can expect n == 2 here.

	// The return value of the above Benchmark could be used just like
	// a b.N loop benchmark as well.
}

func ExampleB_RunParallel() {
	// Parallel benchmark for text/template.Template.Execute on a single object.
	testing.Benchmark(func(b *testing.B) {
		templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
		// RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		b.RunParallel(func(pb *testing.PB) {
			// Each goroutine has its own bytes.Buffer.
			var buf bytes.Buffer
			for pb.Next() {
				// The loop body is executed b.N times total across all goroutines.
				buf.Reset()
				templ.Execute(&buf, "World")
			}
		})
	})
}

func ExampleB_ReportMetric() {
	// This reports a custom benchmark metric relevant to a
	// specific algorithm (in this case, sorting).
	testing.Benchmark(func(b *testing.B) {
		var compares int64
		for i := 0; i < b.N; i++ {
			s := []int{5, 4, 3, 2, 1}
			slices.SortFunc(s, func(a, b int) int {
				compares++
				return cmp.Compare(a, b)
			})
		}
		// This metric is per-operation, so divide by b.N and
		// report it as a "/op" unit.
		b.ReportMetric(float64(compares)/float64(b.N), "compares/op")
		// This metric is per-time, so divide by b.Elapsed and
		// report it as a "/ns" unit.
		b.ReportMetric(float64(compares)/float64(b.Elapsed().Nanoseconds()), "compares/ns")
	})
}

func ExampleB_ReportMetric_parallel() {
	// This reports a custom benchmark metric relevant to a
	// specific algorithm (in this case, sorting) in parallel.
	testing.Benchmark(func(b *testing.B) {
		var compares atomic.Int64
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := []int{5, 4, 3, 2, 1}
				slices.SortFunc(s, func(a, b int) int {
					// Because RunParallel runs the function many
					// times in parallel, we must increment the
					// counter atomically to avoid racing writes.
					compares.Add(1)
					return cmp.Compare(a, b)
				})
			}
		})

		// NOTE: Report each metric once, after all of the parallel
		// calls have completed.

		// This metric is per-operation, so divide by b.N and
		// report it as a "/op" unit.
		b.ReportMetric(float64(compares.Load())/float64(b.N), "compares/op")
		// This metric is per-time, so divide by b.Elapsed and
		// report it as a "/ns" unit.
		b.ReportMetric(float64(compares.Load())/float64(b.Elapsed().Nanoseconds()), "compares/ns")
	})
}
