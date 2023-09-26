// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Page heap.
//
// See malloc.go for overview.

package runtime

// minPhysPageSize is a lower-bound on the physical page size. The
// true physical page size may be larger than this. In contrast,
// sys.PhysPageSize is an upper-bound on the physical page size.

// Main malloc heap.
// The heap itself is the "free[]" and "large" arrays,
// but all the other global data is here too.
//
// mheap must not be heap-allocated because it contains mSpanLists,
// which must not be heap-allocated.
//
//go:notinheap

// A heapArena stores metadata for a heap arena. heapArenas are stored
// outside of the Go heap and accessed via the mheap_.arenas index.
//
// This gets allocated directly from the OS, so ideally it should be a
// multiple of the system page size. For example, avoid adding small
// fields.
//
//go:notinheap

// arenaHint is a hint for where to grow the heap arenas. See
// mheap_.arenaHints.
//
//go:notinheap

// An MSpan representing actual memory has state _MSpanInUse,
// _MSpanManual, or _MSpanFree. Transitions between these states are
// constrained as follows:
//
// * A span may transition from free to in-use or manual during any GC
//   phase.
//
// * During sweeping (gcphase == _GCoff), a span may transition from
//   in-use to free (as a result of sweeping) or manual to free (as a
//   result of stacks being freed).
//
// * During GC (gcphase != _GCoff), a span *must not* transition from
//   manual or in-use to free. Because concurrent GC may read a pointer
//   and then look up its span, the span state must be monotonic.

// mSpanStateNames are the names of the span states, indexed by
// mSpanState.

// mSpanList heads a linked list of spans.
//
//go:notinheap

//go:notinheap

// A spanClass represents the size class and noscan-ness of a span.
//
// Each size class has a noscan spanClass and a scan spanClass. The
// noscan spanClass contains only noscan objects, which do not contain
// pointers and thus do not need to be scanned by the garbage
// collector.

//go:notinheap

// The described object has a finalizer set for it.
//
// specialfinalizer is allocated from non-GC'd memory, so any heap
// pointers must be specially handled.
//
//go:notinheap

// The described object is being heap profiled.
//
//go:notinheap

// gcBits is an alloc/mark bitmap. This is always used as *gcBits.
//
//go:notinheap

//go:notinheap
