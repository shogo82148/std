// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modindex

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/go/build"
)

// Module represents and encoded module index file. It is used to
// do the equivalent of build.Import of packages in the module and answer other
// questions based on the index file's data.
type Module struct {
	modroot string
	d       *decoder
	n       int
}

var ErrNotIndexed = errors.New("not in module index")

// GetPackage returns the IndexPackage for the package at the given path.
// It will return ErrNotIndexed if the directory should be read without
// using the index, for instance because the index is disabled, or the package
// is not in a module.
func GetPackage(modroot, pkgdir string) (*IndexPackage, error)

// GetModule returns the Module for the given modroot.
// It will return ErrNotIndexed if the directory should be read without
// using the index, for instance because the index is disabled, or the package
// is not in a module.
func GetModule(modroot string) (*Module, error)

// Walk calls f for each package in the index, passing the path to that package relative to the module root.
func (m *Module) Walk(f func(path string))

// Import is the equivalent of build.Import given the information in Module.
func (rp *IndexPackage) Import(bctxt build.Context, mode build.ImportMode) (p *build.Package, err error)

// IsStandardPackage reports whether path is a standard package
// for the goroot and compiler using the module index if possible,
// and otherwise falling back to internal/goroot.IsStandardPackage
func IsStandardPackage(goroot_, compiler, path string) bool

// IsDirWithGoFiles is the equivalent of fsys.IsDirWithGoFiles using the information in the index.
func (rp *IndexPackage) IsDirWithGoFiles() (_ bool, err error)

// ScanDir implements imports.ScanDir using the information in the index.
func (rp *IndexPackage) ScanDir(tags map[string]bool) (sortedImports []string, sortedTestImports []string, err error)

// IndexPackage holds the information needed to access information in the
// index needed to load a package in a specific directory.
type IndexPackage struct {
	error error
	dir   string

	modroot string

	// Source files
	sourceFiles []*sourceFile
}

// Package and returns finds the package with the given path (relative to the module root).
// If the package does not exist, Package returns an IndexPackage that will return an
// appropriate error from its methods.
func (m *Module) Package(path string) *IndexPackage
