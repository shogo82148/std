// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Implementation of (safe) user arenas.
//
// This file contains the implementation of user arenas wherein Go values can
// be manually allocated and freed in bulk. The act of manually freeing memory,
// potentially before a GC cycle, means that a garbage collection cycle can be
// delayed, improving efficiency by reducing GC cycle frequency. There are other
// potential efficiency benefits, such as improved locality and access to a more
// efficient allocation strategy.
//
// What makes the arenas here safe is that once they are freed, accessing the
// arena's memory will cause an explicit program fault, and the arena's address
// space will not be reused until no more pointers into it are found. There's one
// exception to this: if an arena allocated memory that isn't exhausted, it's placed
// back into a pool for reuse. This means that a crash is not always guaranteed.
//
// While this may seem unsafe, it still prevents memory corruption, and is in fact
// necessary in order to make new(T) a valid implementation of arenas. Such a property
// is desirable to allow for a trivial implementation. (It also avoids complexities
// that arise from synchronization with the GC when trying to set the arena chunks to
// fault while the GC is active.)
//
// The implementation works in layers. At the bottom, arenas are managed in chunks.
// Each chunk must be a multiple of the heap arena size, or the heap arena size must
// be divisible by the arena chunks. The address space for each chunk, and each
// corresponding heapArena for that addres space, are eternelly reserved for use as
// arena chunks. That is, they can never be used for the general heap. Each chunk
// is also represented by a single mspan, and is modeled as a single large heap
// allocation. It must be, because each chunk contains ordinary Go values that may
// point into the heap, so it must be scanned just like any other object. Any
// pointer into a chunk will therefore always cause the whole chunk to be scanned
// while its corresponding arena is still live.
//
// Chunks may be allocated either from new memory mapped by the OS on our behalf,
// or by reusing old freed chunks. When chunks are freed, their underlying memory
// is returned to the OS, set to fault on access, and may not be reused until the
// program doesn't point into the chunk anymore (the code refers to this state as
// "quarantined"), a property checked by the GC.
//
// The sweeper handles moving chunks out of this quarantine state to be ready for
// reuse. When the chunk is placed into the quarantine state, its corresponding
// span is marked as noscan so that the GC doesn't try to scan memory that would
// cause a fault.
//
// At the next layer are the user arenas themselves. They consist of a single
// active chunk which new Go values are bump-allocated into and a list of chunks
// that were exhausted when allocating into the arena. Once the arena is freed,
// it frees all full chunks it references, and places the active one onto a reuse
// list for a future arena to use. Each arena keeps its list of referenced chunks
// explicitly live until it is freed. Each user arena also maps to an object which
// has a finalizer attached that ensures the arena's chunks are all freed even if
// the arena itself is never explicitly freed.
//
// Pointer-ful memory is bump-allocated from low addresses to high addresses in each
// chunk, while pointer-free memory is bump-allocated from high address to low
// addresses. The reason for this is to take advantage of a GC optimization wherein
// the GC will stop scanning an object when there are no more pointers in it, which
// also allows us to elide clearing the heap bitmap for pointer-free Go values
// allocated into arenas.
//
// Note that arenas are not safe to use concurrently.
//
// In summary, there are 2 resources: arenas, and arena chunks. They exist in the
// following lifecycle:
//
// (1) A new arena is created via newArena.
// (2) Chunks are allocated to hold memory allocated into the arena with new or slice.
//    (a) Chunks are first allocated from the reuse list of partially-used chunks.
//    (b) If there are no such chunks, then chunks on the ready list are taken.
//    (c) Failing all the above, memory for a new chunk is mapped.
// (3) The arena is freed, or all references to it are dropped, triggering its finalizer.
//    (a) If the GC is not active, exhausted chunks are set to fault and placed on a
//        quarantine list.
//    (b) If the GC is active, exhausted chunks are placed on a fault list and will
//        go through step (a) at a later point in time.
//    (c) Any remaining partially-used chunk is placed on a reuse list.
// (4) Once no more pointers are found into quarantined arena chunks, the sweeper
//     takes these chunks out of quarantine and places them on the ready list.

package runtime
