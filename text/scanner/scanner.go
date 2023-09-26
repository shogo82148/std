// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package scanner provides a scanner and tokenizer for UTF-8-encoded text.
// It takes an io.Reader providing the source, which then can be tokenized
// through repeated calls to the Scan function.  For compatibility with
// existing tools, the NUL character is not allowed.
//
// By default, a Scanner skips white space and Go comments and recognizes all
// literals as defined by the Go language specification.  It may be
// customized to recognize only a subset of those literals and to recognize
// different white space characters.
//
// Basic usage pattern:
//
//	var s scanner.Scanner
//	s.Init(src)
//	tok := s.Scan()
//	for tok != scanner.EOF {
//		// do something with tok
//		tok = s.Scan()
//	}
package scanner

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/io"
)

// A source position is represented by a Position value.
// A position is valid if Line > 0.
type Position struct {
	Filename string
	Offset   int
	Line     int
	Column   int
}

// IsValid returns true if the position is valid.
func (pos *Position) IsValid() bool

func (pos Position) String() string

// Predefined mode bits to control recognition of tokens. For instance,
// to configure a Scanner such that it only recognizes (Go) identifiers,
// integers, and skips comments, set the Scanner's Mode field to:
//
//	ScanIdents | ScanInts | SkipComments
const (
	ScanIdents     = 1 << -Ident
	ScanInts       = 1 << -Int
	ScanFloats     = 1 << -Float
	ScanChars      = 1 << -Char
	ScanStrings    = 1 << -String
	ScanRawStrings = 1 << -RawString
	ScanComments   = 1 << -Comment
	SkipComments   = 1 << -skipComment
	GoTokens       = ScanIdents | ScanFloats | ScanChars | ScanStrings | ScanRawStrings | ScanComments | SkipComments
)

// The result of Scan is one of the following tokens or a Unicode character.
const (
	EOF = -(iota + 1)
	Ident
	Int
	Float
	Char
	String
	RawString
	Comment
)

// TokenString returns a printable string for a token or Unicode character.
func TokenString(tok rune) string

// GoWhitespace is the default value for the Scanner's Whitespace field.
// Its value selects Go's white space characters.
const GoWhitespace = 1<<'\t' | 1<<'\n' | 1<<'\r' | 1<<' '

// A Scanner implements reading of Unicode characters and tokens from an io.Reader.
type Scanner struct {
	src io.Reader

	srcBuf [bufLen + 1]byte
	srcPos int
	srcEnd int

	srcBufOffset int
	line         int
	column       int
	lastLineLen  int
	lastCharLen  int

	tokBuf bytes.Buffer
	tokPos int
	tokEnd int

	ch rune

	Error func(s *Scanner, msg string)

	ErrorCount int

	Mode uint

	Whitespace uint64

	Position
}

// Init initializes a Scanner with a new source and returns s.
// Error is set to nil, ErrorCount is set to 0, Mode is set to GoTokens,
// and Whitespace is set to GoWhitespace.
func (s *Scanner) Init(src io.Reader) *Scanner

// Next reads and returns the next Unicode character.
// It returns EOF at the end of the source. It reports
// a read error by calling s.Error, if not nil; otherwise
// it prints an error message to os.Stderr. Next does not
// update the Scanner's Position field; use Pos() to
// get the current position.
func (s *Scanner) Next() rune

// Peek returns the next Unicode character in the source without advancing
// the scanner. It returns EOF if the scanner's position is at the last
// character of the source.
func (s *Scanner) Peek() rune

// Scan reads the next token or Unicode character from source and returns it.
// It only recognizes tokens t for which the respective Mode bit (1<<-t) is set.
// It returns EOF at the end of the source. It reports scanner errors (read and
// token errors) by calling s.Error, if not nil; otherwise it prints an error
// message to os.Stderr.
func (s *Scanner) Scan() rune

// Pos returns the position of the character immediately after
// the character or token returned by the last call to Next or Scan.
func (s *Scanner) Pos() (pos Position)

// TokenText returns the string corresponding to the most recently scanned token.
// Valid after calling Scan().
func (s *Scanner) TokenText() string
