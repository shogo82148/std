// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// PostorderWithNumbering provides a DFS postordering.
// This seems to make loop-finding more robust.
func PostorderWithNumbering(f *Func, ponums []int32) []*Block

func Dominators(f *Func) []*Block

// DominatorsSimple computes the dominator tree for f. It returns a slice
// which maps block ID to the immediate dominator of that block.
// Unreachable blocks map to nil. The entry block maps to nil.
func DominatorsSimple(f *Func) []*Block
