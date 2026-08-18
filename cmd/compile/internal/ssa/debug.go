// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/compile/internal/abt"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/ssa/ssabase"
	"github.com/shogo82148/std/cmd/compile/internal/ssa/ssaop"
	"github.com/shogo82148/std/cmd/internal/obj"
)

type BlockDebug struct {
	// State at the start and end of the block. These are initialized,
	// and updated from new information that flows on back edges.
	startState, endState abt.T
	// Use these to avoid excess work in the merge. If none of the
	// predecessors has changed since the last check, the old answer is
	// still good.
	lastCheckedTime, lastChangedTime int32
	// Whether the block had any changes to user variables at all.
	relevant bool
	// false until the block has been processed at least once. This
	// affects how the merge is done; the goal is to maximize sharing
	// and avoid allocation.
	everProcessed bool
}

var BlockEnd = &Value{
	ID:  -20000,
	Op:  ssaop.OpInvalid,
	Aux: StringToAux("BlockEnd"),
}

var BlockStart = &Value{
	ID:  -10000,
	Op:  ssaop.OpInvalid,
	Aux: StringToAux("BlockStart"),
}

type DebugState struct {
	// See FuncDebug.
	Slots    []LocalSlot
	Vars     []*ir.Name
	VarSlots [][]SlotID
	Lists    [][]LocListEntry

	// The user variable that each slot rolls up to, indexed by SlotID.
	SlotVars []VarID

	F             *Func
	LoggingLevel  int
	ConvergeCount int
	Registers     []ssabase.Register
	StackOffset   func(LocalSlot) int32
	Ctxt          *obj.Link

	// The names (slots) associated with each value, indexed by Value ID.
	ValueNames [][]SlotID

	// The current state of whatever analysis is running.
	currentState StateAtPC
	changedVars  *SparseSet
	changedSlots *SparseSet

	// The pending location list entry for each user variable, indexed by VarID.
	pendingEntries []pendingEntry

	VarParts        map[*ir.Name][]SlotID
	blockDebug      []BlockDebug
	pendingSlotLocs []VarLoc
}

var FuncEnd = &Value{
	ID:  -30000,
	Op:  ssaop.OpInvalid,
	Aux: StringToAux("FuncEnd"),
}

// IsVarWantedForDebug returns true if the debug info for the node should
// be generated.
// For example, internal variables for range-over-func loops have little
// value to users, so we don't generate debug info for them.
func IsVarWantedForDebug(n ir.Node) bool

// LocListEntry represents a single entry in a location list.
// StartBlock/StartValue and EndBlock/EndValue are SSA coordinates
// that get resolved to PCs during final encoding.
type LocListEntry struct {
	StartBlock, StartValue ID
	EndBlock, EndValue     ID
	Expr                   []byte
}

// RegisterSet is a bitmap of registers, indexed by Register.num.
type RegisterSet uint64

type SlotID int32

// StackOffset encodes whether a value is on the stack and if so, where.
// It is a 31-bit integer followed by a presence flag at the low-order
// bit.
type StackOffset int32

// StateAtPC is the current state of all variables at some point.
type StateAtPC struct {
	// The location of each known slot, indexed by SlotID.
	slots []VarLoc
	// The slots present in each register, indexed by register number.
	registers [][]SlotID
}

type VarID int32

// A VarLoc describes the storage for part of a user variable.
type VarLoc struct {
	// The registers this variable is available in. There can be more than
	// one in various situations, e.g. it's being moved between registers.
	Registers RegisterSet

	StackOffset
}

func (s *DebugState) LocString(loc VarLoc) string

// Logf prints debug-specific logging to stdout (always stdout) if the
// current function is tagged by GOSSAFUNC (for ssa output directed
// either to stdout or html).
func (s *DebugState) Logf(msg string, args ...any)

func (state *DebugState) InitializeCache(f *Func, numVars, numSlots int)

// Liveness walks the function in control flow order, calculating the start
// and end state of each block.
func (state *DebugState) Liveness() []*BlockDebug

// BuildLocationLists builds location lists for all the user variables
// in state.f, using the information about block state in blockLocs.
// The returned location lists are not fully complete. They are in
// terms of SSA values rather than PCs, and have no base address/end
// entries. They will be finished by PutLocationList.
func (state *DebugState) BuildLocationLists(blockLocs []*BlockDebug)
