// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements the compressed encoding of source
// positions using a lookup table.

package src

// XPos is a more compact representation of Pos.
type XPos struct {
	index int32
	lico
}

// NoXPos is a valid unknown position.
var NoXPos XPos

// IsKnown reports whether the position p is known.
// XPos.IsKnown() matches Pos.IsKnown() for corresponding
// positions.
func (p XPos) IsKnown() bool

// Before reports whether the position p comes before q in the source.
// For positions with different bases, ordering is by base index.
func (p XPos) Before(q XPos) bool

// SameFile reports whether p and q are positions in the same file.
func (p XPos) SameFile(q XPos) bool

// SameFileAndLine reports whether p and q are positions on the same line in the same file.
func (p XPos) SameFileAndLine(q XPos) bool

// After reports whether the position p comes after q in the source.
// For positions with different bases, ordering is by base index.
func (p XPos) After(q XPos) bool

// WithNotStmt returns the same location to be marked with DWARF is_stmt=0
func (p XPos) WithNotStmt() XPos

// WithDefaultStmt returns the same location with undetermined is_stmt
func (p XPos) WithDefaultStmt() XPos

// WithIsStmt returns the same location to be marked with DWARF is_stmt=1
func (p XPos) WithIsStmt() XPos

// WithBogusLine returns a bogus line that won't match any recorded for the source code.
// Its use is to disrupt the statements within an infinite loop so that the debugger
// will not itself loop infinitely waiting for the line number to change.
// gdb chooses not to display the bogus line; delve shows it with a complaint, but the
// alternative behavior is to hang.
func (p XPos) WithBogusLine() XPos

// WithXlogue returns the same location but marked with DWARF function prologue/epilogue
func (p XPos) WithXlogue(x PosXlogue) XPos

// LineNumber returns a string for the line number, "?" if it is not known.
func (p XPos) LineNumber() string

// FileIndex returns a smallish non-negative integer corresponding to the
// file for this source position.  Smallish is relative; it can be thousands
// large, but not millions.
func (p XPos) FileIndex() int32

func (p XPos) LineNumberHTML() string

// AtColumn1 returns the same location but shifted to column 1.
func (p XPos) AtColumn1() XPos

// A PosTable tracks Pos -> XPos conversions and vice versa.
// Its zero value is a ready-to-use PosTable.
type PosTable struct {
	baseList []*PosBase
	indexMap map[*PosBase]int
	nameMap  map[string]int
}

// XPos returns the corresponding XPos for the given pos,
// adding pos to t if necessary.
func (t *PosTable) XPos(pos Pos) XPos

// Pos returns the corresponding Pos for the given p.
// If p cannot be translated via t, the function panics.
func (t *PosTable) Pos(p XPos) Pos

// FileTable returns a slice of all files used to build this package.
func (t *PosTable) FileTable() []string
