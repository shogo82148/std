// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Binary api computes the exported API of a set of Go packages.
package main

import (
	"github.com/shogo82148/std/go/build"
	"github.com/shogo82148/std/go/types"
	exec "internal/execabs"
)

// Flags

// contexts are the default contexts which are scanned, unless
// overridden by the -contexts flag.

type Walker struct {
	context   *build.Context
	root      string
	scope     []string
	current   *types.Package
	features  map[string]bool
	imported  map[string]*types.Package
	importMap map[string]map[string]string
	importDir map[string]string
}

func NewWalker(context *build.Context, root string) *Walker

func (w *Walker) Features() (fs []string)

// Disable before debugging non-obvious errors from the type-checker.

// Importing is a sentinel taking the place in Walker.imported
// for a package that is in the process of being imported.

func (w *Walker) Import(name string) (*types.Package, error)

func (w *Walker) ImportFrom(fromPath, fromDir string, mode types.ImportMode) (*types.Package, error)
