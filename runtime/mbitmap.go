// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Garbage collector: type and heap bitmaps.
//
// Stack, data, and bss bitmaps
//
// Stack frames and global variables in the data and bss sections are described
// by 1-bit bitmaps in which 0 means uninteresting and 1 means live pointer
// to be visited during GC. The bits in each byte are consumed starting with
// the low bit: 1<<0, 1<<1, and so on.
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
// In each 2-bit entry, the lower bit holds the same information as in the 1-bit
// bitmaps: 0 means uninteresting and 1 means live pointer to be visited during GC.
// The meaning of the high bit depends on the position of the word being described
// in its allocated object. In all words *except* the second word, the
// high bit indicates that the object is still being described. In
// these words, if a bit pair with a high bit 0 is encountered, the
// low bit can also be assumed to be 0, and the object description is
// over. This 00 is called the ``dead'' encoding: it signals that the
// rest of the words in the object are uninteresting to the garbage
// collector.
//
// In the second word, the high bit is the GC ``checkmarked'' bit (see below).
//
// The 2-bit entries are split when written into the byte, so that the top half
// of the byte contains 4 high bits and the bottom half contains 4 low (pointer)
// bits.
// This form allows a copy from the 1-bit to the 4-bit form to keep the
// pointer bits contiguous, instead of having to space them out.
//
// The code makes use of the fact that the zero value for a heap bitmap
// has no live pointer bit set and is (depending on position), not used,
// not checkmarked, and is the dead encoding.
// These properties must be preserved when modifying the encoding.
//
// The bitmap for noscan spans is not maintained. Code must ensure
// that an object is scannable before consulting its bitmap by
// checking either the noscan bit in the span or by consulting its
// type's information.
//
// Checkmarks
//
// In a concurrent garbage collector, one worries about failing to mark
// a live object due to mutations without write barriers or bugs in the
// collector implementation. As a sanity check, the GC has a 'checkmark'
// mode that retraverses the object graph with the world stopped, to make
// sure that everything that should be marked is marked.
// In checkmark mode, in the heap bitmap, the high bit of the 2-bit entry
// for the second word of the object holds the checkmark bit.
// When not in checkmark mode, this bit is set to 1.
//
// The smallest possible allocation is 8 bytes. On a 32-bit machine, that
// means every allocated object has two words, so there is room for the
// checkmark bit. On a 64-bit machine, however, the 8-byte allocation is
// just one word, so the second bit pair is not available for encoding the
// checkmark. However, because non-pointer allocations are combined
// into larger 16-byte (maxTinySize) allocations, a plain 8-byte allocation
// must be a pointer, so the type bit in the first word is not actually needed.
// It is still used in general, except in checkmark the type bit is repurposed
// as the checkmark bit and then reinitialized (to 1) as the type bit when
// finished.
//

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

// oneBitCount is indexed by byte and produces the
// number of 1 bits in that byte. For example 128 has 1 bit set
// and oneBitCount[128] will holds 1.
