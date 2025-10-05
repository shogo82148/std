// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Green Tea mark algorithm
//
// The core idea behind Green Tea is simple: achieve better locality during
// mark/scan by delaying scanning so that we can accumulate objects to scan
// within the same span, then scan the objects that have accumulated on the
// span all together.
//
// By batching objects this way, we increase the chance that adjacent objects
// will be accessed, amortize the cost of accessing object metadata, and create
// better opportunities for prefetching. We can take this even further and
// optimize the scan loop by size class (not yet completed) all the way to the
// point of applying SIMD techniques to really tear through the heap.
//
// Naturally, this depends on being able to create opportunties to batch objects
// together. The basic idea here is to have two sets of mark bits. One set is the
// regular set of mark bits ("marks"), while the other essentially says that the
// objects have been scanned already ("scans"). When we see a pointer for the first
// time we set its mark and enqueue its span. We track these spans in work queues
// with a FIFO policy, unlike workbufs which have a LIFO policy. Empirically, a
// FIFO policy appears to work best for accumulating objects to scan on a span.
// Later, when we dequeue the span, we find both the union and intersection of the
// mark and scan bitsets. The union is then written back into the scan bits, while
// the intersection is used to decide which objects need scanning, such that the GC
// is still precise.
//
// Below is the bulk of the implementation, focusing on the worst case
// for locality, small objects. Specifically, those that are smaller than
// a few cache lines in size and whose metadata is stored the same way (at the
// end of the span).

//go:build goexperiment.greenteagc

package runtime
