// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lex

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/text/scanner"

	"github.com/shogo82148/std/cmd/internal/src"
)

// A Tokenizer is a simple wrapping of text/scanner.Scanner, configured
// for our purposes and made a TokenReader. It forms the lowest level,
// turning text from readers into tokens.
type Tokenizer struct {
	tok  ScanToken
	s    *scanner.Scanner
	base *src.PosBase
	line int
	file *os.File
}

func NewTokenizer(name string, r io.Reader, file *os.File) *Tokenizer

func (t *Tokenizer) Text() string

func (t *Tokenizer) File() string

func (t *Tokenizer) Base() *src.PosBase

func (t *Tokenizer) SetBase(base *src.PosBase)

func (t *Tokenizer) Line() int

func (t *Tokenizer) Col() int

func (t *Tokenizer) Next() ScanToken

func (t *Tokenizer) Close()
