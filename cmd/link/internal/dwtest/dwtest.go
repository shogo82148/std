// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dwtest

import (
	"github.com/shogo82148/std/debug/dwarf"
)

type Examiner struct {
	dies        []*dwarf.Entry
	idxByOffset map[dwarf.Offset]int
	kids        map[int][]int
	parent      map[int]int
	byname      map[string][]int
}

// Populate the Examiner using the DIEs read from rdr.
func (ex *Examiner) Populate(rdr *dwarf.Reader) error

func (ex *Examiner) DIEs() []*dwarf.Entry

// For debugging new tests
func (ex *Examiner) DumpEntry(idx int, dumpKids bool, ilevel int)

// Given a DIE offset, return the previously read dwarf.Entry, or nil
func (ex *Examiner) EntryFromOffset(off dwarf.Offset) *dwarf.Entry

// Return the ID that Examiner uses to refer to the DIE at offset off
func (ex *Examiner) IdxFromOffset(off dwarf.Offset) int

// Returns a list of child entries for a die with ID 'idx'
func (ex *Examiner) Children(idx int) []*dwarf.Entry

// Returns parent DIE for DIE 'idx', or nil if the DIE is top level
func (ex *Examiner) Parent(idx int) *dwarf.Entry

// ParentCU returns the enclosing compilation unit DIE for the DIE
// with a given index, or nil if for some reason we can't establish a
// parent.
func (ex *Examiner) ParentCU(idx int) *dwarf.Entry

// FileRef takes a given DIE by index and a numeric file reference
// (presumably from a decl_file or call_file attribute), looks up the
// reference in the .debug_line file table, and returns the proper
// string for it. We need to know which DIE is making the reference
// so as to find the right compilation unit.
func (ex *Examiner) FileRef(dw *dwarf.Data, dieIdx int, fileRef int64) (string, error)

// Return a list of all DIEs with name 'name'. When searching for DIEs
// by name, keep in mind that the returned results will include child
// DIEs such as params/variables. For example, asking for all DIEs named
// "p" for even a small program will give you 400-500 entries.
func (ex *Examiner) Named(name string) []*dwarf.Entry

// SubprogLoAndHighPc returns the values of the lo_pc and high_pc
// attrs of the DWARF DIE subprogdie.  For DWARF versions 2-3, both of
// these attributes had to be of class address; with DWARF 4 the rules
// were changed, allowing compilers to emit a high PC attr of class
// constant, where the high PC could be computed by starting with the
// low PC address and then adding in the high_pc attr offset.  This
// function accepts both styles of specifying a hi/lo pair, returning
// the values or an error if the attributes are malformed in some way.
func SubprogLoAndHighPc(subprogdie *dwarf.Entry) (lo uint64, hi uint64, err error)
