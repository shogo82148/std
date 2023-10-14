// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

func AtExit(f func())

func Exit(code int)

// To enable tracing support (-t flag), set EnableTrace to true.
const EnableTrace = false

// AdjustStartingHeap modifies GOGC so that GC should not occur until the heap
// grows to the requested size.  This is intended but not promised, though it
// is true-mostly, depending on when the adjustment occurs and on the
// compiler's input and behavior.  Once this size is approximately reached
// GOGC is reset to 100; subsequent GCs may reduce the heap below the requested
// size, but this function does not affect that.
//
// -d=gcadjust=1 enables logging of GOGC adjustment events.
//
// NOTE: If you think this code would help startup time in your own
// application and you decide to use it, please benchmark first to see if it
// actually works for you (it may not: the Go compiler is not typical), and
// whatever the outcome, please leave a comment on bug #56546.  This code
// uses supported interfaces, but depends more than we like on
// current+observed behavior of the garbage collector, so if many people need
// this feature, we should consider/propose a better way to accomplish it.
func AdjustStartingHeap(requestedHeapGoal uint64)

func Compiling(pkgs []string) bool

// Do not instrument the following packages at all,
// at best instrumentation would cause infinite recursion.
var NoInstrumentPkgs = []string{
	"runtime/internal/atomic",
	"runtime/internal/math",
	"runtime/internal/sys",
	"runtime/internal/syscall",
	"runtime",
	"runtime/race",
	"runtime/msan",
	"runtime/asan",
	"internal/cpu",
	"internal/abi",
}

// Don't insert racefuncenter/racefuncexit into the following packages.
// Memory accesses in the packages are either uninteresting or will cause false positives.
var NoRacePkgs = []string{"sync", "sync/atomic"}
