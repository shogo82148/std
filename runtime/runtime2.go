// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

/*
 * defined constants
 */

// A guintptr holds a goroutine pointer, but typed as a uintptr
// to bypass write barriers. It is used in the Gobuf goroutine state
// and in scheduling lists that are manipulated without a P.
//
// The Gobuf.g goroutine pointer is almost always updated by assembly code.
// In one of the few places it is updated by Go code - func save - it must be
// treated as a uintptr to avoid a write barrier being emitted at a bad time.
// Instead of figuring out how to emit the write barriers missing in the
// assembly manipulation, we change the type of the field to uintptr,
// so that it does not require write barriers at all.
//
// Goroutine structs are published in the allg list and never freed.
// That will keep the goroutine structs from being collected.
// There is never a time that Gobuf.g's contain the only references
// to a goroutine: the publishing of the goroutine in allg comes first.
// Goroutine pointers are also kept in non-GC-visible places like TLS,
// so I can't see them ever moving. If we did want to start moving data
// in the GC, we'd need to allocate the goroutine structs from an
// alternate arena. Using guintptr doesn't make that problem any worse.

// Known to compiler.
// Changes here must also be made in src/cmd/internal/gc/select.go's selecttype.

// describes how to handle callback

// Stack describes a Go execution stack.
// The bounds of the stack are exactly [lo, hi),
// with no implicit data structures on either side.

// stkbar records the state of a G's stack barrier.

// The m->locked word holds two pieces of state counting active calls to LockOSThread/lockOSThread.
// The low bit (LockExternal) is a boolean reporting whether any LockOSThread call is active.
// External locks are not recursive; a second lock is silently ignored.
// The upper bits of m->locked record the nesting depth of calls to lockOSThread
// (counting up by LockInternal), popped by unlockOSThread (counting down by LockInternal).
// Internal locks can be recursive. For instance, a lock for cgo can occur while the main
// goroutine is holding the lock during the initialization phase.

// Layout of in-memory per-function information prepared by linker
// See https://golang.org/s/go12symtab.
// Keep in sync with linker
// and with package debug/gosym and with symtab.go in package runtime.

// layout of Itab known to compilers
// allocated in non-garbage-collected memory

// Lock-free stack node.
// // Also known to export_test.go.

/*
 * known to compiler
 */

// startup_random_data holds random bytes initialized at startup.  These come from
// the ELF AT_RANDOM auxiliary vector (vdso_linux_amd64.go or os_linux_386.go).

/*
 * deferred subroutine calls
 */

/*
 * panics
 */

// Set by the linker so the runtime can determine the buildmode.
