// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

// Block represents a basic block in the control flow graph of a function.
type Block struct {
	// A unique identifier for the block. The system will attempt to allocate
	// these IDs densely, but no guarantees.
	ID ID

	// Source position for block's control operation
	Pos src.XPos

	// The kind of block this is.
	Kind BlockKind

	// Likely direction for branches.
	// If BranchLikely, Succs[0] is the most likely branch taken.
	// If BranchUnlikely, Succs[1] is the most likely branch taken.
	// Ignored if len(Succs) < 2.
	// Fatal if not BranchUnknown and len(Succs) > 2.
	Likely BranchPrediction

	// After flagalloc, records whether flags are live at the end of the block.
	FlagsLiveAtEnd bool

	// A block that would be good to align (according to the optimizer's guesses)
	Hotness Hotness

	// Subsequent blocks, if any. The number and order depend on the block kind.
	Succs []Edge

	// Inverse of successors.
	// The order is significant to Phi nodes in the block.
	// TODO: predecessors is a pain to maintain. Can we somehow order phi
	// arguments by block id and have this field computed explicitly when needed?
	Preds []Edge

	// A list of values that determine how the block is exited. The number
	// and type of control values depends on the Kind of the block. For
	// instance, a BlockIf has a single boolean control value and BlockExit
	// has a single memory control value.
	//
	// The ControlValues() method may be used to get a slice with the non-nil
	// control values that can be ranged over.
	//
	// Controls[1] must be nil if Controls[0] is nil.
	Controls [2]*Value

	// Auxiliary info for the block. Its value depends on the Kind.
	Aux    Aux
	AuxInt int64

	// The unordered set of Values that define the operation of this block.
	// After the scheduling pass, this list is ordered.
	Values []*Value

	// The containing function
	Func *Func

	// Storage for Succs, Preds and Values.
	succstorage [2]Edge
	predstorage [4]Edge
	valstorage  [9]*Value
}

// Edge represents a CFG edge.
// Example edges for b branching to either c or d.
// (c and d have other predecessors.)
//
//	b.Succs = [{c,3}, {d,1}]
//	c.Preds = [?, ?, ?, {b,0}]
//	d.Preds = [?, {b,1}, ?]
//
// These indexes allow us to edit the CFG in constant time.
// In addition, it informs phi ops in degenerate cases like:
//
//	b:
//	   if k then c else c
//	c:
//	   v = Phi(x, y)
//
// Then the indexes tell you whether x is chosen from
// the if or else branch from b.
//
//	b.Succs = [{c,0},{c,1}]
//	c.Preds = [{b,0},{b,1}]
//
// means x is chosen if k is true.
type Edge struct {
	// block edge goes to (in a Succs list) or from (in a Preds list)
	b *Block
	// index of reverse edge.  Invariant:
	//   e := x.Succs[idx]
	//   e.b.Preds[e.i] = Edge{x,idx}
	// and similarly for predecessors.
	i int
}

func (e Edge) Block() *Block

func (e Edge) Index() int

func (e Edge) String() string

// BlockKind is the kind of SSA block.
type BlockKind uint8

// short form print
func (b *Block) String() string

// long form print
func (b *Block) LongString() string

// NumControls returns the number of non-nil control values the
// block has.
func (b *Block) NumControls() int

// ControlValues returns a slice containing the non-nil control
// values of the block. The index of each control value will be
// the same as it is in the Controls property and can be used
// in ReplaceControl calls.
func (b *Block) ControlValues() []*Value

// SetControl removes all existing control values and then adds
// the control value provided. The number of control values after
// a call to SetControl will always be 1.
func (b *Block) SetControl(v *Value)

// ResetControls sets the number of controls for the block to 0.
func (b *Block) ResetControls()

// AddControl appends a control value to the existing list of control values.
func (b *Block) AddControl(v *Value)

// ReplaceControl exchanges the existing control value at the index provided
// for the new value. The index must refer to a valid control value.
func (b *Block) ReplaceControl(i int, v *Value)

// CopyControls replaces the controls for this block with those from the
// provided block. The provided block is not modified.
func (b *Block) CopyControls(from *Block)

// Reset sets the block to the provided kind and clears all the blocks control
// and auxiliary values. Other properties of the block, such as its successors,
// predecessors and values are left unmodified.
func (b *Block) Reset(kind BlockKind)

// AddEdgeTo adds an edge from block b to block c.
func (b *Block) AddEdgeTo(c *Block)

// LackingPos indicates whether b is a block whose position should be inherited
// from its successors.  This is true if all the values within it have unreliable positions
// and if it is "plain", meaning that there is no control flow that is also very likely
// to correspond to a well-understood source position.
func (b *Block) LackingPos() bool

func (b *Block) AuxIntString() string

func (b *Block) Logf(msg string, args ...any)
func (b *Block) Log() bool
func (b *Block) Fatalf(msg string, args ...any)

type BranchPrediction int8

const (
	BranchUnlikely = BranchPrediction(-1)
	BranchUnknown  = BranchPrediction(0)
	BranchLikely   = BranchPrediction(+1)
)

type Hotness int8

const (
	// These values are arranged in what seems to be order of increasing alignment importance.
	// Currently only a few are relevant.  Implicitly, they are all in a loop.
	HotNotFlowIn Hotness = 1 << iota
	HotInitial
	HotPgo

	HotNot                 = 0
	HotInitialNotFlowIn    = HotInitial | HotNotFlowIn
	HotPgoInitial          = HotPgo | HotInitial
	HotPgoInitialNotFLowIn = HotPgo | HotInitial | HotNotFlowIn
)
