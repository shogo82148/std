// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/internal/coverage"
	"github.com/shogo82148/std/internal/coverage/encodemeta"

	"github.com/shogo82148/std/cmd/internal/edit"
)

// Block represents the information about a basic block to be recorded in the analysis.
// Note: Our definition of basic block is based on control structures; we don't break
// apart && and ||. We could but it doesn't seem important enough to bother.
type Block struct {
	startByte token.Pos
	endByte   token.Pos
	numStmt   int
}

// Package holds package-specific state.
type Package struct {
	mdb            *encodemeta.CoverageMetaDataBuilder
	counterLengths []int
}

// Function holds func-specific state.
type Func struct {
	units      []coverage.CoverableUnit
	counterVar string
}

// File is a wrapper for the state of a file used in the parser.
// The basic parse tree walker is a method of this type.
type File struct {
	fset    *token.FileSet
	name    string
	astFile *ast.File
	blocks  []Block
	content []byte
	edit    *edit.Buffer
	mdb     *encodemeta.CoverageMetaDataBuilder
	fn      Func
	pkg     *Package
}

// Visit implements the ast.Visitor interface.
func (f *File) Visit(node ast.Node) ast.Visitor
