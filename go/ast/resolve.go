// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements NewPackage.

package ast

import (
	"github.com/shogo82148/std/go/token"
)

// An Importer resolves import paths to package Objects.
// The imports map records the packages already imported,
// indexed by package id (canonical import path).
// An Importer must determine the canonical import path and
// check the map to see if it is already present in the imports map.
// If so, the Importer can return the map entry.  Otherwise, the
// Importer should load the package data for the given path into
// a new *Object (pkg), record pkg in the imports map, and then
// return pkg.
type Importer func(imports map[string]*Object, path string) (pkg *Object, err error)

// NewPackage creates a new Package node from a set of File nodes. It resolves
// unresolved identifiers across files and updates each file's Unresolved list
// accordingly. If a non-nil importer and universe scope are provided, they are
// used to resolve identifiers not declared in any of the package files. Any
// remaining unresolved identifiers are reported as undeclared. If the files
// belong to different packages, one package name is selected and files with
// different package names are reported and then ignored.
// The result is a package node and a scanner.ErrorList if there were errors.
func NewPackage(fset *token.FileSet, files map[string]*File, importer Importer, universe *Scope) (*Package, error)
