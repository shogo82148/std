// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lex

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

// A Slice reads from a slice of Tokens.
type Slice struct {
	tokens []Token
	base   *src.PosBase
	line   int
	pos    int
}

func NewSlice(base *src.PosBase, line int, tokens []Token) *Slice

func (s *Slice) Next() ScanToken

func (s *Slice) Text() string

func (s *Slice) File() string

func (s *Slice) Base() *src.PosBase

func (s *Slice) SetBase(base *src.PosBase)

func (s *Slice) Line() int

func (s *Slice) Col() int

func (s *Slice) Close()
