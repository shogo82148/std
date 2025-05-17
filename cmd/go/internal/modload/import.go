// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modload

import (
	"golang.org/x/mod/module"
)

type ImportMissingError struct {
	Path     string
	Module   module.Version
	QueryErr error

	ImportingMainModule module.Version

	// isStd indicates whether we would expect to find the package in the standard
	// library. This is normally true for all dotless import paths, but replace
	// directives can cause us to treat the replaced paths as also being in
	// modules.
	isStd bool

	// importerGoVersion is the version the module containing the import error
	// specified. It is only set when isStd is true.
	importerGoVersion string

	// replaced the highest replaced version of the module where the replacement
	// contains the package. replaced is only set if the replacement is unused.
	replaced module.Version

	// newMissingVersion is set to a newer version of Module if one is present
	// in the build list. When set, we can't automatically upgrade.
	newMissingVersion string
}

func (e *ImportMissingError) Error() string

func (e *ImportMissingError) Unwrap() error

func (e *ImportMissingError) ImportPath() string

// An AmbiguousImportError indicates an import of a package found in multiple
// modules in the build list, or found in both the main module and its vendor
// directory.
type AmbiguousImportError struct {
	importPath string
	Dirs       []string
	Modules    []module.Version
}

func (e *AmbiguousImportError) ImportPath() string

func (e *AmbiguousImportError) Error() string

// A DirectImportFromImplicitDependencyError indicates a package directly
// imported by a package or test in the main module that is satisfied by a
// dependency that is not explicit in the main module's go.mod file.
type DirectImportFromImplicitDependencyError struct {
	ImporterPath string
	ImportedPath string
	Module       module.Version
}

func (e *DirectImportFromImplicitDependencyError) Error() string

func (e *DirectImportFromImplicitDependencyError) ImportPath() string

// ImportMissingSumError is reported in readonly mode when we need to check
// if a module contains a package, but we don't have a sum for its .zip file.
// We might need sums for multiple modules to verify the package is unique.
//
// TODO(#43653): consolidate multiple errors of this type into a single error
// that suggests a 'go get' command for root packages that transitively import
// packages from modules with missing sums. load.CheckPackageErrors would be
// a good place to consolidate errors, but we'll need to attach the import
// stack here.
type ImportMissingSumError struct {
	importPath                string
	found                     bool
	mods                      []module.Version
	importer, importerVersion string
	importerIsTest            bool
}

func (e *ImportMissingSumError) Error() string

func (e *ImportMissingSumError) ImportPath() string

// PkgIsInLocalModule reports whether the directory of the package with
// the given pkgpath, exists in the module with the given modpath
// at the given modroot, and contains go source files.
func PkgIsInLocalModule(pkgpath, modpath, modroot string) bool
