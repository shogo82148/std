// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements the Check function, which drives type-checking.

package types

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/token"
)

// debugging/development support

// If Strict is set, the type-checker enforces additional
// rules not specified by the Go 1 spec, but which will
// catch guaranteed run-time errors if the respective
// code is executed. In other words, programs passing in
// Strict mode are Go 1 compliant, but not all Go 1 programs
// will pass in Strict mode. The additional rules are:
//
// - A type assertion x.(T) where T is an interface type
//   is invalid if any (statically known) method that exists
//   for both x and T have different signatures.
//

// exprInfo stores information about an untyped expression.

// A context represents the context within which an object is type-checked.

// An importKey identifies an imported package by import path and source directory
// (directory containing the file containing the import). In practice, the directory
// may always be the same, or may not matter. Given an (import path, directory), an
// importer must always return the same package (but given two different import paths,
// an importer may still return the same package by mapping them to the same package
// paths).

// A Checker maintains the state of the type checker.
// It must be created with NewChecker.
type Checker struct {
	conf *Config
	fset *token.FileSet
	pkg  *Package
	*Info
	objMap map[Object]*declInfo
	impMap map[importKey]*Package
	posMap map[*Interface][]token.Pos
	pkgCnt map[string]int

	files            []*ast.File
	unusedDotImports map[*Scope]map[*Package]token.Pos

	firstErr error
	methods  map[*TypeName][]*Func
	untyped  map[ast.Expr]exprInfo
	delayed  []func()
	finals   []func()
	objPath  []Object

	context

	indent int
}

// NewChecker returns a new Checker instance for a given package.
// Package files may be added incrementally via checker.Files.
func NewChecker(conf *Config, fset *token.FileSet, pkg *Package, info *Info) *Checker

// A bailout panic is used for early termination.

// Files checks the provided files as part of the checker's package.
func (check *Checker) Files(files []*ast.File) error
