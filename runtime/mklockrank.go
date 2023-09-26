// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// mklockrank records the static rank graph of the locks in the
// runtime and generates the rank checking structures in lockrank.go.
package main

// ranks describes the lock rank graph. See "go doc internal/dag" for
// the syntax.
//
// "a < b" means a must be acquired before b if both are held
// (or, if b is held, a cannot be acquired).
//
// "NONE < a" means no locks may be held when a is acquired.
//
// If a lock is not given a rank, then it is assumed to be a leaf
// lock, which means no other lock can be acquired while it is held.
// Therefore, leaf locks do not need to be given an explicit rank.
//
// Ranks in all caps are pseudo-nodes that help define order, but do
// not actually define a rank.
//
// TODO: It's often hard to correlate rank names to locks. Change
// these to be more consistent with the locks they label.

// cyclicRanks lists lock ranks that allow multiple locks of the same
// rank to be acquired simultaneously. The runtime enforces ordering
// within these ranks using a separate mechanism.
