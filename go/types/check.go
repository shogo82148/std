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

// exprInfo stores information about an untyped expression.

// An environment represents the environment within which an object is
// type-checked.

// An importKey identifies an imported package by import path and source directory
// (directory containing the file containing the import). In practice, the directory
// may always be the same, or may not matter. Given an (import path, directory), an
// importer must always return the same package (but given two different import paths,
// an importer may still return the same package by mapping them to the same package
// paths).

// A dotImportKey describes a dot-imported object in the given scope.

// An action describes a (delayed) action.

// An actionDesc provides information on an action.
// For debugging only.

// A Checker maintains the state of the type checker.
// It must be created with NewChecker.
type Checker struct {
	conf *Config
	ctxt *Context
	fset *token.FileSet
	pkg  *Package
	*Info
	version version
	nextID  uint64
	objMap  map[Object]*declInfo
	impMap  map[importKey]*Package
	valids  instanceLookup

	pkgPathMap map[string]map[string]bool
	seenPkgMap map[*Package]bool

	files         []*ast.File
	imports       []*PkgName
	dotImportMap  map[dotImportKey]*PkgName
	recvTParamMap map[*ast.Ident]*TypeParam
	brokenAliases map[*TypeName]bool
	unionTypeSets map[*Union]*_TypeSet
	mono          monoGraph

	firstErr error
	methods  map[*TypeName][]*Func
	untyped  map[ast.Expr]exprInfo
	delayed  []action
	objPath  []Object
	cleaners []cleaner

	environment

	indent int
}

// NewChecker returns a new Checker instance for a given package.
// Package files may be added incrementally via checker.Files.
func NewChecker(conf *Config, fset *token.FileSet, pkg *Package, info *Info) *Checker

// A bailout panic is used for early termination.

// Files checks the provided files as part of the checker's package.
func (check *Checker) Files(files []*ast.File) error
