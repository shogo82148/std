// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gccgoimporter implements Import for gccgo-generated object files.
package gccgoimporter

import (
	"github.com/shogo82148/std/go/types"
	"github.com/shogo82148/std/io"
)

// A PackageInit describes an imported package that needs initialization.
type PackageInit struct {
	Name     string
	InitFunc string
	Priority int
}

// The gccgo-specific init data for a package.
type InitData struct {
	// Initialization priority of this package relative to other packages.
	// This is based on the maximum depth of the package's dependency graph;
	// it is guaranteed to be greater than that of its dependencies.
	Priority int

	// The list of packages which this package depends on to be initialized,
	// including itself if needed. This is the subset of the transitive closure of
	// the package's dependencies that need initialization.
	Inits []PackageInit
}

// An Importer resolves import paths to Packages. The imports map records
// packages already known, indexed by package path.
// An importer must determine the canonical package path and check imports
// to see if it is already present in the map. If so, the Importer can return
// the map entry. Otherwise, the importer must load the package data for the
// given path into a new *Package, record it in imports map, and return the
// package.
type Importer func(imports map[string]*types.Package, path, srcDir string, lookup func(string) (io.ReadCloser, error)) (*types.Package, error)

func GetImporter(searchpaths []string, initmap map[*types.Package]InitData) Importer
