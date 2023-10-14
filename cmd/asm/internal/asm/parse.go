// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package asm implements the parser and instruction generator for the assembler.
// TODO: Split apart?
package asm

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/text/scanner"

	"github.com/shogo82148/std/cmd/asm/internal/arch"
	"github.com/shogo82148/std/cmd/asm/internal/lex"
	"github.com/shogo82148/std/cmd/internal/obj"
)

type Parser struct {
	lex              lex.TokenReader
	lineNum          int
	errorLine        int
	errorCount       int
	sawCode          bool
	pc               int64
	input            []lex.Token
	inputPos         int
	pendingLabels    []string
	labels           map[string]*obj.Prog
	toPatch          []Patch
	addr             []obj.Addr
	arch             *arch.Arch
	ctxt             *obj.Link
	firstProg        *obj.Prog
	lastProg         *obj.Prog
	dataAddr         map[string]int64
	isJump           bool
	compilingRuntime bool
	errorWriter      io.Writer
}

type Patch struct {
	addr  *obj.Addr
	label string
}

func NewParser(ctxt *obj.Link, ar *arch.Arch, lexer lex.TokenReader, compilingRuntime bool) *Parser

func (p *Parser) Parse() (*obj.Prog, bool)

// ParseSymABIs parses p's assembly code to find text symbol
// definitions and references and writes a symabis file to w.
func (p *Parser) ParseSymABIs(w io.Writer) bool

// EOF represents the end of input.
var EOF = lex.Make(scanner.EOF, "EOF")
