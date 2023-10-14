// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lex implements lexical analysis for the assembler.
package lex

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

// A ScanToken represents an input item. It is a simple wrapping of rune, as
// returned by text/scanner.Scanner, plus a couple of extra values.
type ScanToken rune

const (
	// Asm defines some two-character lexemes. We make up
	// a rune/ScanToken value for them - ugly but simple.
	LSH ScanToken = -1000 - iota
	RSH
	ARR
	ROT
	Include
	BuildComment
)

// IsRegisterShift reports whether the token is one of the ARM register shift operators.
func IsRegisterShift(r ScanToken) bool

func (t ScanToken) String() string

// NewLexer returns a lexer for the named file and the given link context.
func NewLexer(name string) TokenReader

// A TokenReader is like a reader, but returns lex tokens of type Token. It also can tell you what
// the text of the most recently returned token is, and where it was found.
// The underlying scanner elides all spaces except newline, so the input looks like a stream of
// Tokens; original spacing is lost but we don't need it.
type TokenReader interface {
	Next() ScanToken

	Text() string

	File() string

	Base() *src.PosBase

	SetBase(*src.PosBase)

	Line() int

	Col() int

	Close()
}

// A Token is a scan token plus its string value.
// A macro is stored as a sequence of Tokens with spaces stripped.
type Token struct {
	ScanToken
	text string
}

// Make returns a Token with the given rune (ScanToken) and text representation.
func Make(token ScanToken, text string) Token

func (l Token) String() string

// A Macro represents the definition of a #defined macro.
type Macro struct {
	name   string
	args   []string
	tokens []Token
}

// Tokenize turns a string into a list of Tokens; used to parse the -D flag and in tests.
func Tokenize(str string) []Token
