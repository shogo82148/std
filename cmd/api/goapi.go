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

// aliasReplacer applies type aliases to earlier API files,
// to avoid misleading negative results.
// This makes all the references to os.FileInfo in go1.txt
// be read as if they said fs.FileInfo, since os.FileInfo is now an alias.
// If there are many of these, we could do a more general solution,
// but for now the replacer is fine.

type Walker struct {
	context     *build.Context
	root        string
	scope       []string
	current     *types.Package
	features    map[string]bool
	imported    map[string]*types.Package
	stdPackages []string
	importMap   map[string]map[string]string
	importDir   map[string]string
}

func NewWalker(context *build.Context, root string) *Walker

func (w *Walker) Features() (fs []string)

// Disable before debugging non-obvious errors from the type-checker.

// listSem is a semaphore restricting concurrent invocations of 'go list'.

// Importing is a sentinel taking the place in Walker.imported
// for a package that is in the process of being imported.

func (w *Walker) Import(name string) (*types.Package, error)

func (w *Walker) ImportFrom(fromPath, fromDir string, mode types.ImportMode) (*types.Package, error)
