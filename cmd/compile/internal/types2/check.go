// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements the Check function, which drives type-checking.

package types2

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
)

// A Checker maintains the state of the type checker.
// It must be created with NewChecker.
type Checker struct {
	// package information
	// (initialized by NewChecker, valid for the life-time of checker)
	conf *Config
	ctxt *Context
	pkg  *Package
	*Info
	version goVersion
	nextID  uint64
	objMap  map[Object]*declInfo
	impMap  map[importKey]*Package

	// pkgPathMap maps package names to the set of distinct import paths we've
	// seen for that name, anywhere in the import graph. It is used for
	// disambiguating package names in error messages.
	//
	// pkgPathMap is allocated lazily, so that we don't pay the price of building
	// it on the happy path. seenPkgMap tracks the packages that we've already
	// walked.
	pkgPathMap map[string]map[string]bool
	seenPkgMap map[*Package]bool

	// information collected during type-checking of a set of package files
	// (initialized by Files, valid only for the duration of check.Files;
	// maps and lists are allocated on demand)
	files         []*syntax.File
	versions      map[*syntax.PosBase]string
	imports       []*PkgName
	dotImportMap  map[dotImportKey]*PkgName
	brokenAliases map[*TypeName]bool
	unionTypeSets map[*Union]*_TypeSet
	mono          monoGraph

	firstErr error
	methods  map[*TypeName][]*Func
	untyped  map[syntax.Expr]exprInfo
	delayed  []action
	objPath  []Object
	cleaners []cleaner

	// environment within which the current object is type-checked (valid only
	// for the duration of type-checking a specific object)
	environment

	// debugging
	indent int
}

// NewChecker returns a new Checker instance for a given package.
// Package files may be added incrementally via checker.Files.
func NewChecker(conf *Config, pkg *Package, info *Info) *Checker

// Files checks the provided files as part of the checker's package.
func (check *Checker) Files(files []*syntax.File) (err error)
