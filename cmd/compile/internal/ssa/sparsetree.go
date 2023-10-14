// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

type SparseTreeNode struct {
	child   *Block
	sibling *Block
	parent  *Block

	// Every block has 6 numbers associated with it:
	// entry-1, entry, entry+1, exit-1, and exit, exit+1.
	// entry and exit are conceptually the top of the block (phi functions)
	// entry+1 and exit-1 are conceptually the bottom of the block (ordinary defs)
	// entry-1 and exit+1 are conceptually "just before" the block (conditions flowing in)
	//
	// This simplifies life if we wish to query information about x
	// when x is both an input to and output of a block.
	entry, exit int32
}

func (s *SparseTreeNode) String() string

func (s *SparseTreeNode) Entry() int32

func (s *SparseTreeNode) Exit() int32

const (
	// When used to lookup up definitions in a sparse tree,
	// these adjustments to a block's entry (+adjust) and
	// exit (-adjust) numbers allow a distinction to be made
	// between assignments (typically branch-dependent
	// conditionals) occurring "before" the block (e.g., as inputs
	// to the block and its phi functions), "within" the block,
	// and "after" the block.
	AdjustBefore = -1
	AdjustWithin = 0
	AdjustAfter  = 1
)

// A SparseTree is a tree of Blocks.
// It allows rapid ancestor queries,
// such as whether one block dominates another.
type SparseTree []SparseTreeNode

// Sibling returns a sibling of x in the dominator tree (i.e.,
// a node with the same immediate dominator) or nil if there
// are no remaining siblings in the arbitrary but repeatable
// order chosen. Because the Child-Sibling order is used
// to assign entry and exit numbers in the treewalk, those
// numbers are also consistent with this order (i.e.,
// Sibling(x) has entry number larger than x's exit number).
func (t SparseTree) Sibling(x *Block) *Block

// Child returns a child of x in the dominator tree, or
// nil if there are none. The choice of first child is
// arbitrary but repeatable.
func (t SparseTree) Child(x *Block) *Block

// Parent returns the parent of x in the dominator tree, or
// nil if x is the function's entry.
func (t SparseTree) Parent(x *Block) *Block

// IsAncestorEq reports whether x is an ancestor of or equal to y.
func (t SparseTree) IsAncestorEq(x, y *Block) bool
