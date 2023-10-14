// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

// PosMax is the largest line or column value that can be represented without loss.
// Incoming values (arguments) larger than PosMax will be set to PosMax.
//
// Keep this consistent with maxLineCol in go/scanner.
const PosMax = 1 << 30

// A Pos represents an absolute (line, col) source position
// with a reference to position base for computing relative
// (to a file, or line directive) position information.
// Pos values are intentionally light-weight so that they
// can be created without too much concern about space use.
type Pos struct {
	base      *PosBase
	line, col uint32
}

// MakePos returns a new Pos for the given PosBase, line and column.
func MakePos(base *PosBase, line, col uint) Pos

func (pos Pos) Pos() Pos
func (pos Pos) IsKnown() bool
func (pos Pos) Base() *PosBase
func (pos Pos) Line() uint
func (pos Pos) Col() uint

func (pos Pos) RelFilename() string

func (pos Pos) RelLine() uint

func (pos Pos) RelCol() uint

// Cmp compares the positions p and q and returns a result r as follows:
//
//	r <  0: p is before q
//	r == 0: p and q are the same position (but may not be identical)
//	r >  0: p is after q
//
// If p and q are in different files, p is before q if the filename
// of p sorts lexicographically before the filename of q.
func (p Pos) Cmp(q Pos) int

func (pos Pos) String() string

// A PosBase represents the base for relative position information:
// At position pos, the relative position is filename:line:col.
type PosBase struct {
	pos       Pos
	filename  string
	line, col uint32
	trimmed   bool
}

// NewFileBase returns a new PosBase for the given filename.
// A file PosBase's position is relative to itself, with the
// position being filename:1:1.
func NewFileBase(filename string) *PosBase

// NewTrimmedFileBase is like NewFileBase, but allows specifying Trimmed.
func NewTrimmedFileBase(filename string, trimmed bool) *PosBase

// NewLineBase returns a new PosBase for a line directive "line filename:line:col"
// relative to pos, which is the position of the character immediately following
// the comment containing the line directive. For a directive in a line comment,
// that position is the beginning of the next line (i.e., the newline character
// belongs to the line comment).
func NewLineBase(pos Pos, filename string, trimmed bool, line, col uint) *PosBase

func (base *PosBase) IsFileBase() bool

func (base *PosBase) Pos() (_ Pos)

func (base *PosBase) Filename() string

func (base *PosBase) Line() uint

func (base *PosBase) Col() uint

func (base *PosBase) Trimmed() bool
