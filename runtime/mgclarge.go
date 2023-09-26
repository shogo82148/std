// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Page heap.
//
// See malloc.go for the general overview.
//
// Allocation policy is the subject of this file. All free spans live in
// a treap for most of their time being free. See
// https://en.wikipedia.org/wiki/Treap or
// https://faculty.washington.edu/aragon/pubs/rst89.pdf for an overview.
// sema.go also holds an implementation of a treap.
//
// Each treapNode holds a single span. The treap is sorted by base address
// and each span necessarily has a unique base address.
// Spans are returned based on a first-fit algorithm, acquiring the span
// with the lowest base address which still satisfies the request.
//
// The first-fit algorithm is possible due to an augmentation of each
// treapNode to maintain the size of the largest span in the subtree rooted
// at that treapNode. Below we refer to this invariant as the maxPages
// invariant.
//
// The primary routines are
// insert: adds a span to the treap
// remove: removes the span from that treap that best fits the required size
// removeSpan: which removes a specific span from the treap
//
// Whenever a pointer to a span which is owned by the treap is acquired, that
// span must not be mutated. To mutate a span in the treap, remove it first.
//
// mheap_.lock must be held when manipulating this data structure.

package runtime

//go:notinheap

//go:notinheap

// treapIterType represents the type of iteration to perform
// over the treap. Each different flag is represented by a bit
// in the type, and types may be combined together by a bitwise
// or operation.
//
// Note that only 5 bits are available for treapIterType, do not
// use the 3 higher-order bits. This constraint is to allow for
// expansion into a treapIterFilter, which is a uint32.

// treapIterFilter is a bitwise filter of different spans by binary
// properties. Each bit of a treapIterFilter represents a unique
// combination of bits set in a treapIterType, in other words, it
// represents the power set of a treapIterType.
//
// The purpose of this representation is to allow the existence of
// a specific span type to bubble up in the treap (see the types
// field on treapNode).
//
// More specifically, any treapIterType may be transformed into a
// treapIterFilter for a specific combination of flags via the
// following operation: 1 << (0x1f&treapIterType).

// treapFilterAll represents the filter which allows all spans.

// treapIter is a bidirectional iterator type which may be used to iterate over a
// an mTreap in-order forwards (increasing order) or backwards (decreasing order).
// Its purpose is to hide details about the treap from users when trying to iterate
// over it.
//
// To create iterators over the treap, call start or end on an mTreap.
