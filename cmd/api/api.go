// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package api computes the exported API of a set of Go packages.
// It is only a test, not a command, nor a usefully importable package.
package api

import (
	"github.com/shogo82148/std/go/build"
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/go/types"
	"github.com/shogo82148/std/testing"
)

// contexts are the default contexts which are scanned.

func Check(t *testing.T)

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
	current     *apiPackage
	deprecated  map[token.Pos]bool
	features    map[string]bool
	imported    map[string]*apiPackage
	stdPackages []string
	importMap   map[string]map[string]string
	importDir   map[string]string
}

func NewWalker(context *build.Context, root string) *Walker

func (w *Walker) Features() (fs []string)

// Disable before debugging non-obvious errors from the type-checker.

// listSem is a semaphore restricting concurrent invocations of 'go list'. 'go
// list' has its own internal concurrency, so we use a hard-coded constant (to
// allow the I/O-intensive phases of 'go list' to overlap) instead of scaling
// all the way up to GOMAXPROCS.

// Importing is a sentinel taking the place in Walker.imported
// for a package that is in the process of being imported.

// Import implements types.Importer.
func (w *Walker) Import(name string) (*types.Package, error)

// ImportFrom implements types.ImporterFrom.
func (w *Walker) ImportFrom(fromPath, fromDir string, mode types.ImportMode) (*types.Package, error)
