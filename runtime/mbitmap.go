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
// The heap bitmap comprises 1 bit for each pointer-sized word in the heap,
// recording whether a pointer is stored in that word or not. This bitmap
// is stored in the heapArena metadata backing each heap arena.
// That is, if ha is the heapArena for the arena starting at "start",
// then ha.bitmap[0] holds the 64 bits for the 64 words "start"
// through start+63*ptrSize, ha.bitmap[1] holds the entries for
// start+64*ptrSize through start+127*ptrSize, and so on.
// Bits correspond to words in little-endian order. ha.bitmap[0]&1 represents
// the word at "start", ha.bitmap[0]>>1&1 represents the word at start+8, etc.
// (For 32-bit platforms, s/64/32/.)
//
// We also keep a noMorePtrs bitmap which allows us to stop scanning
// the heap bitmap early in certain situations. If ha.noMorePtrs[i]>>j&1
// is 1, then the object containing the last word described by ha.bitmap[8*i+j]
// has no more pointers beyond those described by ha.bitmap[8*i+j].
// If ha.noMorePtrs[i]>>j&1 is set, the entries in ha.bitmap[8*i+j+1] and
// beyond must all be zero until the start of the next object.
//
// The bitmap for noscan spans is set to all zero at span allocation time.
//
// The bitmap for unallocated objects in scannable spans is not maintained
// (can be junk).

package runtime
