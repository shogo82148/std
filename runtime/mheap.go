// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Page heap.
//
// See malloc.go for overview.

package runtime

// Main malloc heap.
// The heap itself is the "free" and "scav" treaps,
// but all the other global data is here too.
//
// mheap must not be heap-allocated because it contains mSpanLists,
// which must not be heap-allocated.

// A heapArena stores metadata for a heap arena. heapArenas are stored
// outside of the Go heap and accessed via the mheap_.arenas index.

// arenaHint is a hint for where to grow the heap arenas. See
// mheap_.arenaHints.

// An mspan representing actual memory has state mSpanInUse,
// mSpanManual, or mSpanFree. Transitions between these states are
// constrained as follows:
//
//   - A span may transition from free to in-use or manual during any GC
//     phase.
//
//   - During sweeping (gcphase == _GCoff), a span may transition from
//     in-use to free (as a result of sweeping) or manual to free (as a
//     result of stacks being freed).
//
//   - During GC (gcphase != _GCoff), a span *must not* transition from
//     manual or in-use to free. Because concurrent GC may read a pointer
//     and then look up its span, the span state must be monotonic.
//
// Setting mspan.state to mSpanInUse or mSpanManual must be done
// atomically and only after all other span fields are valid.
// Likewise, if inspecting a span is contingent on it being
// mSpanInUse, the state should be loaded atomically and checked
// before depending on other fields. This allows the garbage collector
// to safely deal with potentially invalid pointers, since resolving
// such pointers may race with a span being allocated.

// mSpanStateNames are the names of the span states, indexed by
// mSpanState.

// mSpanStateBox holds an atomic.Uint8 to provide atomic operations on
// an mSpanState. This is a separate type to disallow accidental comparison
// or assignment with mSpanState.

// mSpanList heads a linked list of spans.

// A spanClass represents the size class and noscan-ness of a span.
//
// Each size class has a noscan spanClass and a scan spanClass. The
// noscan spanClass contains only noscan objects, which do not contain
// pointers and thus do not need to be scanned by the garbage
// collector.

// spanAllocType represents the type of allocation to make, or
// the type of allocation to be freed.

// The described object has a finalizer set for it.
//
// specialfinalizer is allocated from non-GC'd memory, so any heap
// pointers must be specially handled.

// The described object is being heap profiled.

// specialReachable tracks whether an object is reachable on the next
// GC cycle. This is used by testing.

// specialsIter helps iterate over specials lists.

// gcBits is an alloc/mark bitmap. This is always used as gcBits.x.
