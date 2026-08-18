// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssadebug

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/ssa"
	"github.com/shogo82148/std/cmd/internal/obj"
)

// A FuncDebug contains all the debug information for the variables in a
// function. Variables are identified by their LocalSlot, which may be
// the result of decomposing a larger variable.
type FuncDebug struct {
	// Slots is all the slots used in the debug info, indexed by their SlotID.
	Slots []ssa.LocalSlot
	// The user variables, indexed by VarID.
	Vars []*ir.Name
	// The slots that make up each variable, indexed by VarID.
	VarSlots [][]ssa.SlotID
	// The location list data, indexed by VarID. Must be processed by PutLocationList.
	LocationLists [][]ssa.LocListEntry
	// Register-resident output parameters for the function. This is filled in at
	// SSA generation time.
	RegOutputParams []*ir.Name
	// Variable declarations that were removed during optimization
	OptDcl []*ir.Name
	// The ssa.Func.EntryID value, used to build location lists for
	// return values promoted to heap in later DWARF generation.
	EntryID ssa.ID

	// Filled in by the user. Translates Block and Value ID to PC.
	//
	// NOTE: block is only used if value is BlockStart.ID or BlockEnd.ID.
	// Otherwise, it is ignored.
	GetPC func(block, value ssa.ID) int64
}

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
func PopulateABIInRegArgOps(f *ssa.Func)

// BuildFuncDebug builds debug information for f, placing the results
// in "rval". f must be fully processed, so that each Value is where it
// will be when machine code is emitted.
func BuildFuncDebug(ctxt *obj.Link, f *ssa.Func, loggingLevel int, stackOffset func(ssa.LocalSlot) int32, rval *FuncDebug)

// PutLocationList adds entries (a location list in structured form)
// to listSym, encoding it in the appropriate DWARF format.
func (debugInfo *FuncDebug) PutLocationList(entries []ssa.LocListEntry, ctxt *obj.Link, listSym, startPC *obj.LSym)

// PutLocationListDwarf5 adds entries (a location list in structured form)
// to listSym in DWARF 5 format.
func (debugInfo *FuncDebug) PutLocationListDwarf5(entries []ssa.LocListEntry, ctxt *obj.Link, listSym, startPC *obj.LSym)

// PutLocationListDwarf4 adds entries (a location list in structured form)
// to listSym in DWARF 4 format.
func (debugInfo *FuncDebug) PutLocationListDwarf4(entries []ssa.LocListEntry, ctxt *obj.Link, listSym, startPC *obj.LSym)

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
func BuildFuncDebugNoOptimized(ctxt *obj.Link, f *ssa.Func, loggingEnabled bool, stackOffset func(ssa.LocalSlot) int32, rval *FuncDebug)
