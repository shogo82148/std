// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package benchmark provides a Metrics object that enables memory and CPU
// profiling for the linker. The Metrics objects can be used to mark stages
// of the code, and name the measurements during that stage. There is also
// optional GCs that can be performed at the end of each stage, so you
// can get an accurate measurement of how each stage changes live memory.
package benchmark

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
)

type Flags int

const (
	GC         = 1 << iota
	NoGC Flags = 0
)

type Metrics struct {
	gc        Flags
	marks     []*mark
	curMark   *mark
	filebase  string
	pprofFile *os.File
}

// New creates a new Metrics object.
//
// Typical usage should look like:
//
//	func main() {
//	  filename := "" // Set to enable per-phase pprof file output.
//	  bench := benchmark.New(benchmark.GC, filename)
//	  defer bench.Report(os.Stdout)
//	  // etc
//	  bench.Start("foo")
//	  foo()
//	  bench.Start("bar")
//	  bar()
//	}
//
// Note that a nil Metrics object won't cause any errors, so one could write
// code like:
//
//	func main() {
//	  enableBenchmarking := flag.Bool("enable", true, "enables benchmarking")
//	  flag.Parse()
//	  var bench *benchmark.Metrics
//	  if *enableBenchmarking {
//	    bench = benchmark.New(benchmark.GC)
//	  }
//	  bench.Start("foo")
//	  // etc.
//	}
func New(gc Flags, filebase string) *Metrics

// Report reports the metrics.
// Closes the currently Start(ed) range, and writes the report to the given io.Writer.
func (m *Metrics) Report(w io.Writer)

// Start marks the beginning of a new measurement phase.
// Once a metric is started, it continues until either a Report is issued, or another Start is called.
func (m *Metrics) Start(name string)
