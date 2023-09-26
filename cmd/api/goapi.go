// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Binary api computes the exported API of a set of Go packages.
package main

import (
	"github.com/shogo82148/std/go/build"
	"github.com/shogo82148/std/go/types"
)

// Flags

// contexts are the default contexts which are scanned, unless
// overridden by the -contexts flag.

type Walker struct {
	context  *build.Context
	root     string
	scope    []string
	current  *types.Package
	features map[string]bool
	imported map[string]*types.Package
}

func NewWalker(context *build.Context, root string) *Walker

func (w *Walker) Features() (fs []string)

// The package cache doesn't operate correctly in rare (so far artificial)
// circumstances (issue 8425). Disable before debugging non-obvious errors
// from the type-checker.

// Importing is a sentinel taking the place in Walker.imported
// for a package that is in the process of being imported.

func (w *Walker) Import(name string) (*types.Package, error)
