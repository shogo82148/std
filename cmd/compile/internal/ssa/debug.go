// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/compile/internal/abt"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/internal/obj"
)

type SlotID int32
type VarID int32

// A FuncDebug contains all the debug information for the variables in a
// function. Variables are identified by their LocalSlot, which may be
// the result of decomposing a larger variable.
type FuncDebug struct {
	// Slots is all the slots used in the debug info, indexed by their SlotID.
	Slots []LocalSlot
	// The user variables, indexed by VarID.
	Vars []*ir.Name
	// The slots that make up each variable, indexed by VarID.
	VarSlots [][]SlotID
	// The location list data, indexed by VarID. Must be processed by PutLocationList.
	LocationLists [][]byte
	// Register-resident output parameters for the function. This is filled in at
	// SSA generation time.
	RegOutputParams []*ir.Name
	// Variable declarations that were removed during optimization
	OptDcl []*ir.Name

	// Filled in by the user. Translates Block and Value ID to PC.
	//
	// NOTE: block is only used if value is BlockStart.ID or BlockEnd.ID.
	// Otherwise, it is ignored.
	GetPC func(block, value ID) int64
}

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

// StackOffset encodes whether a value is on the stack and if so, where.
// It is a 31-bit integer followed by a presence flag at the low-order
// bit.
type StackOffset int32

// A VarLoc describes the storage for part of a user variable.
type VarLoc struct {
	// The registers this variable is available in. There can be more than
	// one in various situations, e.g. it's being moved between registers.
	Registers RegisterSet

	StackOffset
}

var BlockStart = &Value{
	ID:  -10000,
	Op:  OpInvalid,
	Aux: StringToAux("BlockStart"),
}

var BlockEnd = &Value{
	ID:  -20000,
	Op:  OpInvalid,
	Aux: StringToAux("BlockEnd"),
}

var FuncEnd = &Value{
	ID:  -30000,
	Op:  OpInvalid,
	Aux: StringToAux("FuncEnd"),
}

// RegisterSet is a bitmap of registers, indexed by Register.num.
type RegisterSet uint64

type SlKeyIdx uint32

// PopulateABIInRegArgOps examines the entry block of the function
// and looks for incoming parameters that have missing or partial
// OpArg{Int,Float}Reg values, inserting additional values in
// cases where they are missing. Example:
//
//	func foo(s string, used int, notused int) int {
//	  return len(s) + used
//	}
//
// In the function above, the incoming parameter "used" is fully live,
// "notused" is not live, and "s" is partially live (only the length
// field of the string is used). At the point where debug value
// analysis runs, we might expect to see an entry block with:
//
//	b1:
//	  v4 = ArgIntReg <uintptr> {s+8} [0] : BX
//	  v5 = ArgIntReg <int> {used} [0] : CX
//
// While this is an accurate picture of the live incoming params,
// we also want to have debug locations for non-live params (or
// their non-live pieces), e.g. something like
//
//	b1:
//	  v9 = ArgIntReg <*uint8> {s+0} [0] : AX
//	  v4 = ArgIntReg <uintptr> {s+8} [0] : BX
//	  v5 = ArgIntReg <int> {used} [0] : CX
//	  v10 = ArgIntReg <int> {unused} [0] : DI
//
// This function examines the live OpArg{Int,Float}Reg values and
// synthesizes new (dead) values for the non-live params or the
// non-live pieces of partially live params.
func PopulateABIInRegArgOps(f *Func)

// BuildFuncDebug debug information for f, placing the results
// in "rval". f must be fully processed, so that each Value is where it
// will be when machine code is emitted.
func BuildFuncDebug(ctxt *obj.Link, f *Func, loggingLevel int, stackOffset func(LocalSlot) int32, rval *FuncDebug)

// PutLocationList adds list (a location list in its intermediate
// representation) to listSym.
func (debugInfo *FuncDebug) PutLocationList(list []byte, ctxt *obj.Link, listSym, startPC *obj.LSym)

// PutLocationListDwarf5 adds list (a location list in its intermediate
// representation) to listSym in DWARF 5 format. NB: this is a somewhat
// hacky implementation in that it actually reads a DWARF4 encoded
// info from list (with all its DWARF4-specific quirks) then re-encodes
// it in DWARF5. It would probably be better at some point to have
// ssa/debug encode the list in a version-independent form and then
// have this func (and PutLocationListDwarf4) intoduce the quirks.
func (debugInfo *FuncDebug) PutLocationListDwarf5(list []byte, ctxt *obj.Link, listSym, startPC *obj.LSym)

// PutLocationListDwarf4 adds list (a location list in its intermediate
// representation) to listSym in DWARF 4 format.
func (debugInfo *FuncDebug) PutLocationListDwarf4(list []byte, ctxt *obj.Link, listSym, startPC *obj.LSym)

// BuildFuncDebugNoOptimized populates a FuncDebug object "rval" with
// entries corresponding to the register-resident input parameters for
// the function "f"; it is used when we are compiling without
// optimization but the register ABI is enabled. For each reg param,
// it constructs a 2-element location list: the first element holds
// the input register, and the second element holds the stack location
// of the param (the assumption being that when optimization is off,
// each input param reg will be spilled in the prolog). In addition
// to the register params, here we also build location lists (where
// appropriate for the ".closureptr" compiler-synthesized variable
// needed by the debugger for range func bodies.
func BuildFuncDebugNoOptimized(ctxt *obj.Link, f *Func, loggingEnabled bool, stackOffset func(LocalSlot) int32, rval *FuncDebug)

// IsVarWantedForDebug returns true if the debug info for the node should
// be generated.
// For example, internal variables for range-over-func loops have little
// value to users, so we don't generate debug info for them.
func IsVarWantedForDebug(n ir.Node) bool
