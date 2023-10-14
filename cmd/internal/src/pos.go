// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements the encoding of source positions.

package src

import (
	"github.com/shogo82148/std/io"
)

// A Pos encodes a source position consisting of a (line, column) number pair
// and a position base. A zero Pos is a ready to use "unknown" position (nil
// position base and zero line number).
//
// The (line, column) values refer to a position in a file independent of any
// position base ("absolute" file position).
//
// The position base is used to determine the "relative" position, that is the
// filename and line number relative to the position base. If the base refers
// to the current file, there is no difference between absolute and relative
// positions. If it refers to a //line directive, a relative position is relative
// to that directive. A position base in turn contains the position at which it
// was introduced in the current file.
type Pos struct {
	base *PosBase
	lico
}

// NoPos is a valid unknown position.
var NoPos Pos

// MakePos creates a new Pos value with the given base, and (file-absolute)
// line and column.
func MakePos(base *PosBase, line, col uint) Pos

// IsKnown reports whether the position p is known.
// A position is known if it either has a non-nil
// position base, or a non-zero line number.
func (p Pos) IsKnown() bool

// Before reports whether the position p comes before q in the source.
// For positions in different files, ordering is by filename.
func (p Pos) Before(q Pos) bool

// After reports whether the position p comes after q in the source.
// For positions in different files, ordering is by filename.
func (p Pos) After(q Pos) bool

func (p Pos) LineNumber() string

func (p Pos) LineNumberHTML() string

// Filename returns the name of the actual file containing this position.
func (p Pos) Filename() string

// Base returns the position base.
func (p Pos) Base() *PosBase

// SetBase sets the position base.
func (p *Pos) SetBase(base *PosBase)

// RelFilename returns the filename recorded with the position's base.
func (p Pos) RelFilename() string

// RelLine returns the line number relative to the position's base.
func (p Pos) RelLine() uint

// RelCol returns the column number relative to the position's base.
func (p Pos) RelCol() uint

// AbsFilename() returns the absolute filename recorded with the position's base.
func (p Pos) AbsFilename() string

// SymFilename() returns the absolute filename recorded with the position's base,
// prefixed by FileSymPrefix to make it appropriate for use as a linker symbol.
func (p Pos) SymFilename() string

func (p Pos) String() string

// Format formats a position as "filename:line" or "filename:line:column",
// controlled by the showCol flag and if the column is known (!= 0).
// For positions relative to line directives, the original position is
// shown as well, as in "filename:line[origfile:origline:origcolumn] if
// showOrig is set.
func (p Pos) Format(showCol, showOrig bool) string

// WriteTo a position to w, formatted as Format does.
func (p Pos) WriteTo(w io.Writer, showCol, showOrig bool)

// A PosBase encodes a filename and base position.
// Typically, each file and line directive introduce a PosBase.
type PosBase struct {
	pos         Pos
	filename    string
	absFilename string
	symFilename string
	line, col   uint
	inl         int
}

// NewFileBase returns a new *PosBase for a file with the given (relative and
// absolute) filenames.
func NewFileBase(filename, absFilename string) *PosBase

// NewLinePragmaBase returns a new *PosBase for a line directive of the form
//
//	//line filename:line:col
//	/*line filename:line:col*/
//
// at position pos.
func NewLinePragmaBase(pos Pos, filename, absFilename string, line, col uint) *PosBase

// NewInliningBase returns a copy of the old PosBase with the given inlining
// index. If old == nil, the resulting PosBase has no filename.
func NewInliningBase(old *PosBase, inlTreeIndex int) *PosBase

// Pos returns the position at which base is located.
// If b == nil, the result is the zero position.
func (b *PosBase) Pos() *Pos

// Filename returns the filename recorded with the base.
// If b == nil, the result is the empty string.
func (b *PosBase) Filename() string

// AbsFilename returns the absolute filename recorded with the base.
// If b == nil, the result is the empty string.
func (b *PosBase) AbsFilename() string

const FileSymPrefix = "gofile.."

// SymFilename returns the absolute filename recorded with the base,
// prefixed by FileSymPrefix to make it appropriate for use as a linker symbol.
// If b is nil, SymFilename returns FileSymPrefix + "??".
func (b *PosBase) SymFilename() string

// Line returns the line number recorded with the base.
// If b == nil, the result is 0.
func (b *PosBase) Line() uint

// Col returns the column number recorded with the base.
// If b == nil, the result is 0.
func (b *PosBase) Col() uint

// InliningIndex returns the index into the global inlining
// tree recorded with the base. If b == nil or the base has
// not been inlined, the result is < 0.
func (b *PosBase) InliningIndex() int

const (
	// It is expected that the front end or a phase in SSA will usually generate positions tagged with
	// PosDefaultStmt, but note statement boundaries with PosIsStmt.  Simple statements will have a single
	// boundary; for loops with initialization may have one for their entry and one for their back edge
	// (this depends on exactly how the loop is compiled; the intent is to provide a good experience to a
	// user debugging a program; the goal is that a breakpoint set on the loop line fires both on entry
	// and on iteration).  Proper treatment of non-gofmt input with multiple simple statements on a single
	// line is TBD.
	//
	// Optimizing compilation will move instructions around, and some of these will become known-bad as
	// step targets for debugging purposes (examples: register spills and reloads; code generated into
	// the entry block; invariant code hoisted out of loops) but those instructions will still have interesting
	// positions for profiling purposes. To reflect this these positions will be changed to PosNotStmt.
	//
	// When the optimizer removes an instruction marked PosIsStmt; it should attempt to find a nearby
	// instruction with the same line marked PosDefaultStmt to be the new statement boundary.  I.e., the
	// optimizer should make a best-effort to conserve statement boundary positions, and might be enhanced
	// to note when a statement boundary is not conserved.
	//
	// Code cloning, e.g. loop unrolling or loop unswitching, is an exception to the conservation rule
	// because a user running a debugger would expect to see breakpoints active in the copies of the code.
	//
	// In non-optimizing compilation there is still a role for PosNotStmt because of code generation
	// into the entry block.  PosIsStmt statement positions should be conserved.
	//
	// When code generation occurs any remaining default-marked positions are replaced with not-statement
	// positions.
	//
	PosDefaultStmt uint = iota
	PosIsStmt
	PosNotStmt
)

type PosXlogue uint

const (
	PosDefaultLogue PosXlogue = iota
	PosPrologueEnd
	PosEpilogueBegin
)
