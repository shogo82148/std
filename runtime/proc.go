// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Gosched yields the processor, allowing other goroutines to run. It does not
// suspend the current goroutine, so execution resumes automatically.
//
//go:nosplit
func Gosched()

// Breakpoint executes a breakpoint trap.
func Breakpoint()

// LockOSThread wires the calling goroutine to its current operating system thread.
// The calling goroutine will always execute in that thread,
// and no other goroutine will execute in it,
// until the calling goroutine has made as many calls to
// [UnlockOSThread] as to LockOSThread.
// If the calling goroutine exits without unlocking the thread,
// the thread will be terminated.
//
// All init functions are run on the startup thread. Calling LockOSThread
// from an init function will cause the main function to be invoked on
// that thread.
//
// A goroutine should call LockOSThread before calling OS services or
// non-Go library functions that depend on per-thread state.
//
//go:nosplit
func LockOSThread()

// UnlockOSThread undoes an earlier call to LockOSThread.
// If this drops the number of active LockOSThread calls on the
// calling goroutine to zero, it unwires the calling goroutine from
// its fixed operating system thread.
// If there are no active LockOSThread calls, this is a no-op.
//
// Before calling UnlockOSThread, the caller must ensure that the OS
// thread is suitable for running other goroutines. If the caller made
// any permanent changes to the state of the thread that would affect
// other goroutines, it should not call this function and thus leave
// the goroutine locked to the OS thread until the goroutine (and
// hence the thread) exits.
//
//go:nosplit
func UnlockOSThread()
