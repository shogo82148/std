// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Vet is a simple checker for static errors in Go source code.
// See doc.go for more information.
package main

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/go/types"
)

// "-all" flag enables all non-experimental checks

// Flags to control which individual checks to perform.

// experimental records the flags enabling experimental features. These must be
// requested explicitly; they are not enabled by -all.

// setTrueCount record how many flags are explicitly set to true.

// dirsRun and filesRun indicate whether the vet is applied to directory or
// file targets. The distinction affects which checks are run.

// includesNonTest indicates whether the vet is applied to non-test targets.
// Certain checks are relevant only if they touch both test and non-test files.

// A triState is a boolean that knows whether it has been set to either true or false.
// It is used to identify if a flag appears; the standard boolean flag cannot
// distinguish missing from unset. It also satisfies flag.Value.

// Usage is a replacement usage function for the flags package.
func Usage()

// File is a wrapper for the state of a file used in the parser.
// The parse tree walkers are all methods of this type.
type File struct {
	pkg     *Package
	fset    *token.FileSet
	name    string
	content []byte
	file    *ast.File
	b       bytes.Buffer

	basePkg *Package

	stringers map[*ast.Object]bool

	checkers map[ast.Node][]func(*File, ast.Node)

	dead map[ast.Node]bool
}

type Package struct {
	path      string
	defs      map[*ast.Ident]types.Object
	uses      map[*ast.Ident]types.Object
	selectors map[*ast.SelectorExpr]*types.Selection
	types     map[ast.Expr]types.TypeAndValue
	spans     map[types.Object]Span
	files     []*File
	typesPkg  *types.Package
}

// Println is fmt.Println guarded by -v.
func Println(args ...interface{})

// Printf is fmt.Printf guarded by -v.
func Printf(format string, args ...interface{})

// Bad reports an error and sets the exit code..
func (f *File) Bad(pos token.Pos, args ...interface{})

// Badf reports a formatted error and sets the exit code.
func (f *File) Badf(pos token.Pos, format string, args ...interface{})

// Warn reports an error but does not set the exit code.
func (f *File) Warn(pos token.Pos, args ...interface{})

// Warnf reports a formatted error but does not set the exit code.
func (f *File) Warnf(pos token.Pos, format string, args ...interface{})

// Visit implements the ast.Visitor interface.
func (f *File) Visit(node ast.Node) ast.Visitor
