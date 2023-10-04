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

func Check(t *testing.T)

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

// Import implements types.Importer.
func (w *Walker) Import(name string) (*types.Package, error)

// ImportFrom implements types.ImporterFrom.
func (w *Walker) ImportFrom(fromPath, fromDir string, mode types.ImportMode) (*types.Package, error)
