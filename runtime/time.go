// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Time-related runtime and pieces of package time.

package runtime

// Package time knows the layout of this structure.
// If this struct changes, adjust ../time/sleep.go:/runtimeTimer.
// For GOOS=nacl, package syscall knows the layout of this structure.
// If this struct changes, adjust ../syscall/net_nacl.go:/runtimeTimer.

// timersLen is the length of timers array.
//
// Ideally, this would be set to GOMAXPROCS, but that would require
// dynamic reallocation
//
// The current value is a compromise between memory usage and performance
// that should cover the majority of GOMAXPROCS values used in the wild.

// timers contains "per-P" timer heaps.
//
// Timers are queued into timersBucket associated with the current P,
// so each P may work with its own timers independently of other P instances.
//
// Each timersBucket may be associated with multiple P
// if GOMAXPROCS > timersLen.

//go:notinheap

// nacl fake time support - time in nanoseconds since 1970

// Monotonic times are reported as offsets from startNano.
// We initialize startNano to nanotime() - 1 so that on systems where
// monotonic time resolution is fairly low (e.g. Windows 2008
// which appears to have a default resolution of 15ms),
// we avoid ever reporting a nanotime of 0.
// (Callers may want to use 0 as "time not set".)
