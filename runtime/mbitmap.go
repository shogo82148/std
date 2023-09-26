// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Garbage collector: type and heap bitmaps.
//
// Stack, data, and bss bitmaps
//
// Stack frames and global variables in the data and bss sections are
// described by bitmaps with 1 bit per pointer-sized word. A "1" bit
// means the word is a live pointer to be visited by the GC (referred to
// as "pointer"). A "0" bit means the word should be ignored by GC
// (referred to as "scalar", though it could be a dead pointer value).
//
// Heap bitmap
//
// The heap bitmap comprises 2 bits for each pointer-sized word in the heap,
// stored in the heapArena metadata backing each heap arena.
// That is, if ha is the heapArena for the arena starting a start,
// then ha.bitmap[0] holds the 2-bit entries for the four words start
// through start+3*ptrSize, ha.bitmap[1] holds the entries for
// start+4*ptrSize through start+7*ptrSize, and so on.
//
// In each 2-bit entry, the lower bit is a pointer/scalar bit, just
// like in the stack/data bitmaps described above. The upper bit
// indicates scan/dead: a "1" value ("scan") indicates that there may
// be pointers in later words of the allocation, and a "0" value
// ("dead") indicates there are no more pointers in the allocation. If
// the upper bit is 0, the lower bit must also be 0, and this
// indicates scanning can ignore the rest of the allocation.
//
// The 2-bit entries are split when written into the byte, so that the top half
// of the byte contains 4 high (scan) bits and the bottom half contains 4 low
// (pointer) bits. This form allows a copy from the 1-bit to the 4-bit form to
// keep the pointer bits contiguous, instead of having to space them out.
//
// The code makes use of the fact that the zero value for a heap
// bitmap means scalar/dead. This property must be preserved when
// modifying the encoding.
//
// The bitmap for noscan spans is not maintained. Code must ensure
// that an object is scannable before consulting its bitmap by
// checking either the noscan bit in the span or by consulting its
// type's information.

package runtime

// heapBits provides access to the bitmap bits for a single heap word.
// The methods on heapBits take value receivers so that the compiler
// can more easily inline calls to those methods and registerize the
// struct fields independently.

// Make the compiler check that heapBits.arena is large enough to hold
// the maximum arena frame number.
var _ = heapBits{arena: (1<<heapAddrBits)/heapArenaBytes - 1}

// markBits provides access to the mark bit for an object in the heap.
// bytep points to the byte holding the mark bit.
// mask is a byte with a single bit set that can be &ed with *bytep
// to see if the bit has been set.
// *m.byte&m.mask != 0 indicates the mark bit is set.
// index can be used along with span information to generate
// the address of the object in the heap.
// We maintain one set of mark bits for allocation and one for
// marking purposes.

// clobberdeadPtr is a special value that is used by the compiler to
// clobber dead stack slots, when -clobberdead flag is set.
