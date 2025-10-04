// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Writes dwarf information to object files.

package obj

import (
	"github.com/shogo82148/std/cmd/internal/dwarf"
	"github.com/shogo82148/std/sync"
)

// Generate a sequence of opcodes that is as short as possible.
// See section 6.2.5
const (
	LINE_BASE   = -4
	LINE_RANGE  = 10
	PC_RANGE    = (255 - OPCODE_BASE) / LINE_RANGE
	OPCODE_BASE = 11
)

// DwarfIntConst creates a link symbol for an integer constant with the
// given name, type and value.
func (ctxt *Link) DwarfIntConst(name, typename string, val int64)

// DwarfGlobal creates a link symbol containing a DWARF entry for
// a global variable.
func (ctxt *Link) DwarfGlobal(typename string, varSym *LSym)

func (ctxt *Link) DwarfAbstractFunc(curfn Func, s *LSym)

// This table is designed to aid in the creation of references between
// DWARF subprogram DIEs.
//
// In most cases when one DWARF DIE has to refer to another DWARF DIE,
// the target of the reference has an LSym, which makes it easy to use
// the existing relocation mechanism. For DWARF inlined routine DIEs,
// however, the subprogram DIE has to refer to a child
// parameter/variable DIE of the abstract subprogram. This child DIE
// doesn't have an LSym, and also of interest is the fact that when
// DWARF generation is happening for inlined function F within caller
// G, it's possible that DWARF generation hasn't happened yet for F,
// so there is no way to know the offset of a child DIE within F's
// abstract function. Making matters more complex, each inlined
// instance of F may refer to a subset of the original F's variables
// (depending on what happens with optimization, some vars may be
// eliminated).
//
// The fixup table below helps overcome this hurdle. At the point
// where a parameter/variable reference is made (via a call to
// "ReferenceChildDIE"), a fixup record is generate that records
// the relocation that is targeting that child variable. At a later
// point when the abstract function DIE is emitted, there will be
// a call to "RegisterChildDIEOffsets", at which point the offsets
// needed to apply fixups are captured. Finally, once the parallel
// portion of the compilation is done, fixups can actually be applied
// during the "Finalize" method (this can't be done during the
// parallel portion of the compile due to the possibility of data
// races).
//
// This table is also used to record the "precursor" function node for
// each function that is the target of an inline -- child DIE references
// have to be made with respect to the original pre-optimization
// version of the function (to allow for the fact that each inlined
// body may be optimized differently).
type DwarfFixupTable struct {
	ctxt      *Link
	mu        sync.Mutex
	symtab    map[*LSym]int
	svec      []symFixups
	precursor map[*LSym]fnState
}

func NewDwarfFixupTable(ctxt *Link) *DwarfFixupTable

func (ft *DwarfFixupTable) GetPrecursorFunc(s *LSym) Func

func (ft *DwarfFixupTable) SetPrecursorFunc(s *LSym, fn Func)

// Make a note of a child DIE reference: relocation 'ridx' within symbol 's'
// is targeting child 'c' of DIE with symbol 'tgt'.
func (ft *DwarfFixupTable) ReferenceChildDIE(s *LSym, ridx int, tgt *LSym, dclidx int, inlIndex int)

// Called once DWARF generation is complete for a given abstract function,
// whose children might have been referenced via a call above. Stores
// the offsets for any child DIEs (vars, params) so that they can be
// consumed later in on DwarfFixupTable.Finalize, which applies any
// outstanding fixups.
func (ft *DwarfFixupTable) RegisterChildDIEOffsets(s *LSym, vars []*dwarf.Var, coffsets []int32)

// return the LSym corresponding to the 'abstract subprogram' DWARF
// info entry for a function.
func (ft *DwarfFixupTable) AbsFuncDwarfSym(fnsym *LSym) *LSym

// Called after all functions have been compiled; the main job of this
// function is to identify cases where there are outstanding fixups.
// This scenario crops up when we have references to variables of an
// inlined routine, but that routine is defined in some other package.
// This helper walks through and locate these fixups, then invokes a
// helper to create an abstract subprogram DIE for each one.
func (ft *DwarfFixupTable) Finalize(myimportpath string, trace bool)
