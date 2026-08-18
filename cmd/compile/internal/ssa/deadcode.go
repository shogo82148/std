// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// LiveValues returns the live values in f and a list of values that are eligible
// to be statements in reversed data flow order.
// The second result is used to help conserve statement boundaries for debugging.
// reachable is a map from block ID to whether the block is reachable.
// The caller should call f.Cache.freeBoolSlice(live) and f.Cache.freeValueSlice(liveOrderStmts).
// when they are done with the return values.
func LiveValues(f *Func, reachable []bool) (live []bool, liveOrderStmts []*Value)

// ReachableBlocks returns the reachable blocks in f.
func ReachableBlocks(f *Func) []bool

// RemoveEdge removes the i'th outgoing edge from b (and
// the corresponding incoming edge from b.Succs[i].b).
// Note that this potentially reorders successors of b, so it
// must be used very carefully.
func (b *Block) RemoveEdge(i int)
