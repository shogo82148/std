// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

// A Prog is a compiled regular expression program.
type Prog struct {
	Inst   []Inst
	Start  int
	NumCap int
}

// An InstOp is an instruction opcode.
type InstOp uint8

const (
	InstAlt InstOp = iota
	InstAltMatch
	InstCapture
	InstEmptyWidth
	InstMatch
	InstFail
	InstNop
	InstRune
	InstRune1
	InstRuneAny
	InstRuneAnyNotNL
)

func (i InstOp) String() string

// An EmptyOp specifies a kind or mixture of zero-width assertions.
type EmptyOp uint8

const (
	EmptyBeginLine EmptyOp = 1 << iota
	EmptyEndLine
	EmptyBeginText
	EmptyEndText
	EmptyWordBoundary
	EmptyNoWordBoundary
)

// EmptyOpContext returns the zero-width assertions
// satisfied at the position between the runes r1 and r2.
// Passing r1 == -1 indicates that the position is
// at the beginning of the text.
// Passing r2 == -1 indicates that the position is
// at the end of the text.
func EmptyOpContext(r1, r2 rune) EmptyOp

// IsWordChar reports whether r is consider a “word character”
// during the evaluation of the \b and \B zero-width assertions.
// These assertions are ASCII-only: the word characters are [A-Za-z0-9_].
func IsWordChar(r rune) bool

// An Inst is a single instruction in a regular expression program.
type Inst struct {
	Op   InstOp
	Out  uint32
	Arg  uint32
	Rune []rune
}

func (p *Prog) String() string

// Prefix returns a literal string that all matches for the
// regexp must start with.  Complete is true if the prefix
// is the entire match.
func (p *Prog) Prefix() (prefix string, complete bool)

// StartCond returns the leading empty-width conditions that must
// be true in any match.  It returns ^EmptyOp(0) if no matches are possible.
func (p *Prog) StartCond() EmptyOp

// MatchRune reports whether the instruction matches (and consumes) r.
// It should only be called when i.Op == InstRune.
func (i *Inst) MatchRune(r rune) bool

// MatchRunePos checks whether the instruction matches (and consumes) r.
// If so, MatchRunePos returns the index of the matching rune pair
// (or, when len(i.Rune) == 1, rune singleton).
// If not, MatchRunePos returns -1.
// MatchRunePos should only be called when i.Op == InstRune.
func (i *Inst) MatchRunePos(r rune) int

// MatchEmptyWidth reports whether the instruction matches
// an empty string between the runes before and after.
// It should only be called when i.Op == InstEmptyWidth.
func (i *Inst) MatchEmptyWidth(before rune, after rune) bool

func (i *Inst) String() string
