// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
)

// outputfiles is the list of *.cover.go instrumented outputs to write,
// one per input (set when -pkgcfg is in use)

// covervarsoutfile is an additional Go source file into which we'll
// write definitions of coverage counter variables + meta data variables
// (set when -pkgcfg is in use).

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

// funcLitFinder implements the ast.Visitor pattern to find the location of any
// function literal in a subtree.

// pos2 is a pair of token.Position values, used as a map key type.

// seenPos2 tracks whether we have seen a token.Position pair.
