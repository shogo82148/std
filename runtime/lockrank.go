// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file records the static ranks of the locks in the runtime. If a lock
// is not given a rank, then it is assumed to be a leaf lock, which means no other
// lock can be acquired while it is held. Therefore, leaf locks do not need to be
// given an explicit rank. We list all of the architecture-independent leaf locks
// for documentation purposes, but don't list any of the architecture-dependent
// locks (which are all leaf locks). debugLock is ignored for ranking, since it is used
// when printing out lock ranking errors.
//
// lockInit(l *mutex, rank int) is used to set the rank of lock before it is used.
// If there is no clear place to initialize a lock, then the rank of a lock can be
// specified during the lock call itself via lockWithrank(l *mutex, rank int).
//
// Besides the static lock ranking (which is a total ordering of the locks), we
// also represent and enforce the actual partial order among the locks in the
// arcs[] array below. That is, if it is possible that lock B can be acquired when
// lock A is the previous acquired lock that is still held, then there should be
// an entry for A in arcs[B][]. We will currently fail not only if the total order
// (the lock ranking) is violated, but also if there is a missing entry in the
// partial order.

package runtime

// Constants representing the lock rank of the architecture-independent locks in
// the runtime. Locks with lower rank must be taken before locks with higher
// rank.

// lockRankLeafRank is the rank of lock that does not have a declared rank, and hence is
// a leaf lock.

// lockNames gives the names associated with each of the above ranks

// lockPartialOrder is a partial order among the various lock types, listing the immediate
// ordering that has actually been observed in the runtime. Each entry (which
// corresponds to a particular lock rank) specifies the list of locks that can be
// already be held immediately "above" it.
//
// So, for example, the lockRankSched entry shows that all the locks preceding it in
// rank can actually be held. The fin lock shows that only the sched, timers, or
// hchan lock can be held immediately above it when it is acquired.
