// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Scavenging free pages.
//
// This file implements scavenging (the release of physical pages backing mapped
// memory) of free and unused pages in the heap as a way to deal with page-level
// fragmentation and reduce the RSS of Go applications.
//
// Scavenging in Go happens on two fronts: there's the background
// (asynchronous) scavenger and the allocation-time (synchronous) scavenger.
//
// The former happens on a goroutine much like the background sweeper which is
// soft-capped at using scavengePercent of the mutator's time, based on
// order-of-magnitude estimates of the costs of scavenging. The latter happens
// when allocating pages from the heap.
//
// The scavenger's primary goal is to bring the estimated heap RSS of the
// application down to a goal.
//
// Before we consider what this looks like, we need to split the world into two
// halves. One in which a memory limit is not set, and one in which it is.
//
// For the former, the goal is defined as:
//   (retainExtraPercent+100) / 100 * (heapGoal / lastHeapGoal) * lastHeapInUse
//
// Essentially, we wish to have the application's RSS track the heap goal, but
// the heap goal is defined in terms of bytes of objects, rather than pages like
// RSS. As a result, we need to take into account for fragmentation internal to
// spans. heapGoal / lastHeapGoal defines the ratio between the current heap goal
// and the last heap goal, which tells us by how much the heap is growing and
// shrinking. We estimate what the heap will grow to in terms of pages by taking
// this ratio and multiplying it by heapInUse at the end of the last GC, which
// allows us to account for this additional fragmentation. Note that this
// procedure makes the assumption that the degree of fragmentation won't change
// dramatically over the next GC cycle. Overestimating the amount of
// fragmentation simply results in higher memory use, which will be accounted
// for by the next pacing up date. Underestimating the fragmentation however
// could lead to performance degradation. Handling this case is not within the
// scope of the scavenger. Situations where the amount of fragmentation balloons
// over the course of a single GC cycle should be considered pathologies,
// flagged as bugs, and fixed appropriately.
//
// An additional factor of retainExtraPercent is added as a buffer to help ensure
// that there's more unscavenged memory to allocate out of, since each allocation
// out of scavenged memory incurs a potentially expensive page fault.
//
// If a memory limit is set, then we wish to pick a scavenge goal that maintains
// that memory limit. For that, we look at total memory that has been committed
// (memstats.mappedReady) and try to bring that down below the limit. In this case,
// we want to give buffer space in the *opposite* direction. When the application
// is close to the limit, we want to make sure we push harder to keep it under, so
// if we target below the memory limit, we ensure that the background scavenger is
// giving the situation the urgency it deserves.
//
// In this case, the goal is defined as:
//    (100-reduceExtraPercent) / 100 * memoryLimit
//
// We compute both of these goals, and check whether either of them have been met.
// The background scavenger continues operating as long as either one of the goals
// has not been met.
//
// The goals are updated after each GC.
//
// Synchronous scavenging happens for one of two reasons: if an allocation would
// exceed the memory limit or whenever the heap grows in size, for some
// definition of heap-growth. The intuition behind this second reason is that the
// application had to grow the heap because existing fragments were not sufficiently
// large to satisfy a page-level memory allocation, so we scavenge those fragments
// eagerly to offset the growth in RSS that results.
//
// Lastly, not all pages are available for scavenging at all times and in all cases.
// The background scavenger and heap-growth scavenger only release memory in chunks
// that have not been densely-allocated for at least 1 full GC cycle. The reason
// behind this is likelihood of reuse: the Go heap is allocated in a first-fit order
// and by the end of the GC mark phase, the heap tends to be densely packed. Releasing
// memory in these densely packed chunks while they're being packed is counter-productive,
// and worse, it breaks up huge pages on systems that support them. The scavenger (invoked
// during memory allocation) further ensures that chunks it identifies as "dense" are
// immediately eligible for being backed by huge pages. Note that for the most part these
// density heuristics are best-effort heuristics. It's totally possible (but unlikely)
// that a chunk that just became dense is scavenged in the case of a race between memory
// allocation and scavenging.
//
// When synchronously scavenging for the memory limit or for debug.FreeOSMemory, these
// "dense" packing heuristics are ignored (in other words, scavenging is "forced") because
// in these scenarios returning memory to the OS is more important than keeping CPU
// overheads low.

package runtime
