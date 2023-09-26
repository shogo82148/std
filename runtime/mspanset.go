// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// A spanSet is a set of *mspans.
//
// spanSet is safe for concurrent push and pop operations.

// atomicSpanSetSpinePointer is an atomically-accessed spanSetSpinePointer.
//
// It has the same semantics as atomic.UnsafePointer.

// spanSetSpinePointer represents a pointer to a contiguous block of atomic.Pointer[spanSetBlock].

// spanSetBlockPool is a global pool of spanSetBlocks.

// spanSetBlockAlloc represents a concurrent pool of spanSetBlocks.

// haidTailIndex represents a combined 32-bit head and 32-bit tail
// of a queue into a single 64-bit value.

// atomicHeadTailIndex is an atomically-accessed headTailIndex.

// atomicMSpanPointer is an atomic.Pointer[mspan]. Can't use generics because it's NotInHeap.
