// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

// A Regexp is a node in a regular expression syntax tree.
type Regexp struct {
	Op       Op
	Flags    Flags
	Sub      []*Regexp
	Sub0     [1]*Regexp
	Rune     []rune
	Rune0    [2]rune
	Min, Max int
	Cap      int
	Name     string
}

// An Op is a single regular expression operator.
type Op uint8

const (
	OpNoMatch        Op = 1 + iota
	OpEmptyMatch
	OpLiteral
	OpCharClass
	OpAnyCharNotNL
	OpAnyChar
	OpBeginLine
	OpEndLine
	OpBeginText
	OpEndText
	OpWordBoundary
	OpNoWordBoundary
	OpCapture
	OpStar
	OpPlus
	OpQuest
	OpRepeat
	OpConcat
	OpAlternate
)

// Equal reports whether x and y have identical structure.
func (x *Regexp) Equal(y *Regexp) bool

// printFlags is a bit set indicating which flags (including non-capturing parens) to print around a regexp.

func (re *Regexp) String() string

// MaxCap walks the regexp to find the maximum capture index.
func (re *Regexp) MaxCap() int

// CapNames walks the regexp to find the names of capturing groups.
func (re *Regexp) CapNames() []string
