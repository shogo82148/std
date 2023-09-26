// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// main_init_done is a signal used by cgocallbackg that initialization
// has been completed. It is made before _cgo_notify_runtime_init_done,
// so all cgo calls can rely on it existing. When main_init is complete,
// it is closed, meaning cgocallbackg can reliably receive from it.

// runtimeInitTime is the nanotime() at which the runtime started.

// Value to use for signal mask for newly created M's.

// Gosched yields the processor, allowing other goroutines to run.  It does not
// suspend the current goroutine, so execution resumes automatically.
func Gosched()

// freezeStopWait is a large value that freezetheworld sets
// sched.stopwait to in order to request that all Gs permanently stop.

// Holding worldsema grants an M the right to try to stop the world
// and prevents gomaxprocs from changing concurrently.

// When running with cgo, we call _cgo_thread_start
// to start threads for us so that we can play nicely with
// foreign code.

// Breakpoint executes a breakpoint trap.
func Breakpoint()

// LockOSThread wires the calling goroutine to its current operating system thread.
// Until the calling goroutine exits or calls UnlockOSThread, it will always
// execute in that thread, and no other goroutine can.
func LockOSThread()

// UnlockOSThread unwires the calling goroutine from its fixed operating system thread.
// If the calling goroutine has not called LockOSThread, UnlockOSThread is a no-op.
func UnlockOSThread()

// forcegcperiod is the maximum time in nanoseconds between garbage
// collections. If we go this long without a garbage collection, one
// is forced to run.
//
// This is a variable for testing purposes. It normally doesn't change.

// forcePreemptNS is the time slice given to a G before it is
// preempted.

// To shake out latent assumptions about scheduling order,
// we introduce some randomness into scheduling decisions
// when running with the race detector.
// The need for this was made obvious by changing the
// (deterministic) scheduling order in Go 1.5 and breaking
// many poorly-written tests.
// With the randomness here, as long as the tests pass
// consistently with -race, they shouldn't have latent scheduling
// assumptions.
