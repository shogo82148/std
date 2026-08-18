// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

// NewSparseTree creates a SparseTree from a block-to-parent map (array indexed by Block.ID).
// The children of a given node are in reverse postorder.
// This has the nice property that for a given tree walk, the source block of all
// non-retreating edges are visited before their destination block.
func NewSparseTree(f *Func, parentOf []*Block) SparseTree

// A SparseTree is a tree of Blocks.
// It allows rapid ancestor queries,
// such as whether one block dominates another.
type SparseTree []SparseTreeNode

type SparseTreeNode struct {
	Child   *Block
	Sibling *Block
	parent  *Block

	// Every block has 6 numbers associated with it:
	// Entry-1, Entry, Entry+1, Exit-1, and Exit, Exit+1.
	// Entry and Exit are conceptually the top of the block (phi functions)
	// Entry+1 and Exit-1 are conceptually the bottom of the block (ordinary defs)
	// Entry-1 and Exit+1 are conceptually "just before" the block (conditions flowing in)
	//
	// This simplifies life if we wish to query information about x
	// when x is both an input to and output of a block.
	Entry, Exit int32
}

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

func (s *SparseTreeNode) String() string

// Treestructure provides a string description of the dominator
// tree and flow structure of block b and all blocks that it
// dominates.
func (t SparseTree) Treestructure(b *Block) string

func (t SparseTree) NumberBlock(b *Block, n int32) int32

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

// IsAncestor reports whether x is a strict ancestor of y.
func (t SparseTree) IsAncestor(x, y *Block) bool

// DomOrder returns a value for dominator-oriented sorting.
// Block domination does not provide a total ordering,
// but DomOrder two has useful properties.
//  1. If DomOrder(x) > DomOrder(y) then x does not dominate y.
//  2. If DomOrder(x) < DomOrder(y) and DomOrder(y) < DomOrder(z) and x does not dominate y,
//     then x does not dominate z.
//
// Property (1) means that blocks sorted by DomOrder always have a maximal dominant block first.
// Property (2) allows searches for dominated blocks to exit early.
func (t SparseTree) DomOrder(x *Block) int32
