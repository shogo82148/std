// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lex

// Input is the main input: a stack of readers and some macro definitions.
// It also handles #include processing (by pushing onto the input stack)
// and parses and instantiates macro definitions.
type Input struct {
	Stack
	includes        []string
	beginningOfLine bool
	ifdefStack      []bool
	macros          map[string]*Macro
	text            string
	peek            bool
	peekToken       ScanToken
	peekText        string
}

// NewInput returns an Input from the given path.
func NewInput(name string) *Input

func (in *Input) Error(args ...any)

func (in *Input) Next() ScanToken

func (in *Input) Text() string

func (in *Input) Push(r TokenReader)

func (in *Input) Close()
