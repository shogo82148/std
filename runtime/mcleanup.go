// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// AddCleanup attaches a cleanup function to ptr. Some time after ptr is no longer
// reachable, the runtime will call cleanup(arg) in a separate goroutine.
//
// If ptr is reachable from cleanup or arg, ptr will never be collected
// and the cleanup will never run. AddCleanup panics if arg is equal to ptr.
//
// The cleanup(arg) call is not always guaranteed to run; in particular it is not
// guaranteed to run before program exit.
//
// Cleanups are not guaranteed to run if the size of T is zero bytes, because
// it may share same address with other zero-size objects in memory. See
// https://go.dev/ref/spec#Size_and_alignment_guarantees.
//
// There is no specified order in which cleanups will run.
//
// A single goroutine runs all cleanup calls for a program, sequentially. If a
// cleanup function must run for a long time, it should create a new goroutine.
//
// If ptr has both a cleanup and a finalizer, the cleanup will only run once
// it has been finalized and becomes unreachable without an associated finalizer.
//
// It is not guaranteed that a cleanup will run for objects allocated
// in initializers for package-level variables. Such objects may be
// linker-allocated, not heap-allocated.
//
// Note that because cleanups may execute arbitrarily far into the future
// after an object is no longer referenced, the runtime is allowed to perform
// a space-saving optimization that batches objects together in a single
// allocation slot. The cleanup for an unreferenced object in such an
// allocation may never run if it always exists in the same batch as a
// referenced object. Typically, this batching only happens for tiny
// (on the order of 16 bytes or less) and pointer-free objects.
func AddCleanup[T, S any](ptr *T, cleanup func(S), arg S) Cleanup

// Cleanup is a handle to a cleanup call for a specific object.
type Cleanup struct {
	// id is the unique identifier for the cleanup within the arena.
	id uint64
	// ptr contains the pointer to the object.
	ptr uintptr
}

// Stop cancels the cleanup call. Stop will have no effect if the cleanup call
// has already been queued for execution (because ptr became unreachable).
// To guarantee that Stop removes the cleanup function, the caller must ensure
// that the pointer that was passed to AddCleanup is reachable across the call to Stop.
func (c Cleanup) Stop()
