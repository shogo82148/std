// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

type Loop struct {
	Header *Block
	Outer  *Loop

	// Next three fields used by regalloc and/or
	// aid in computation of inner-ness and list of blocks.
	NBlocks int32
	Depth   int16
	IsInner bool

	// True if all paths through the loop have a call.
	// Computed and used by regalloc; stored here for convenience.
	ContainsUnavoidableCall bool
}

type LoopNest struct {
	F              *Func
	B2L            []*Loop
	Po             []*Block
	SDom           SparseTree
	Loops          []*Loop
	HasIrreducible bool
}

func Loopnestfor(f *Func) *LoopNest

func (l *Loop) String() string

func (l *Loop) LongString() string

func (l *Loop) IsWithinOrEq(ll *Loop) bool

// Depth returns the loop nesting level of block b.
func (ln *LoopNest) Depth(b ID) int16
