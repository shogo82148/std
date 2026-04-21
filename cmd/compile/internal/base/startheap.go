// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

// AdjustStartingHeap modifies GOGC so that GC should not occur until the heap
// grows to the requested size.  This is intended but not promised, though it
// is true-mostly, depending on when the adjustment occurs and on the
// compiler's input and behavior.  Once the live heap is approximately half
// this size, GOGC is reset to its value when AdjustStartingHeap was called;
// subsequent GCs may reduce the heap below the requested size, but this
// function does not affect that.
//
// logHeapTweaks (-d=gcadjust=1) enables logging of GOGC adjustment events.
//
// The temporarily requested GOGC is derated from what would be the "obvious"
// value necessary to hit the starting heap goal because the obvious
// (goal/live-1)*100 value seems to grow RSS a little more than it "should"
// (compared to GOMEMLIMIT, e.g.) and the assumption is that the GC's control
// algorithms are tuned for GOGC near 100, and not tuned for huge values of
// GOGC.  Different derating factors apply for "lo" and "hi" values of GOGC;
// lo is below derateBreak, hi is above derateBreak.  The derating factors,
// expressed as integer percentages, are derateLoPct and derateHiPct.
// 60-75 is an okay value for derateLoPct, 30-65 seems like a good value for
// derateHiPct, and 600 seems like a good value for derateBreak.  If these
// are zero, defaults are used instead.
//
// NOTE: If you think this code would help startup time in your own
// application and you decide to use it, please benchmark first to see if it
// actually works for you (it may not: the Go compiler is not typical), and
// whatever the outcome, please leave a comment on bug #56546.  This code
// uses supported interfaces, but depends more than we like on
// current+observed behavior of the garbage collector, so if many people need
// this feature, we should consider/propose a better way to accomplish it.
func AdjustStartingHeap(requestedHeapGoal, derateBreak, derateLoPct, derateHiPct uint64, logHeapTweaks bool)
