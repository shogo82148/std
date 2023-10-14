// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lex

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

// A Stack is a stack of TokenReaders. As the top TokenReader hits EOF,
// it resumes reading the next one down.
type Stack struct {
	tr []TokenReader
}

// Push adds tr to the top (end) of the input stack. (Popping happens automatically.)
func (s *Stack) Push(tr TokenReader)

func (s *Stack) Next() ScanToken

func (s *Stack) Text() string

func (s *Stack) File() string

func (s *Stack) Base() *src.PosBase

func (s *Stack) SetBase(base *src.PosBase)

func (s *Stack) Line() int

func (s *Stack) Col() int

func (s *Stack) Close()
