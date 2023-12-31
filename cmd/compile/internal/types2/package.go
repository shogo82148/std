// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

// A Package describes a Go package.
type Package struct {
	path      string
	name      string
	scope     *Scope
	imports   []*Package
	complete  bool
	fake      bool
	cgo       bool
	goVersion string
}

// NewPackage returns a new Package for the given package path and name.
// The package is not complete and contains no explicit imports.
func NewPackage(path, name string) *Package

// Path returns the package path.
func (pkg *Package) Path() string

// Name returns the package name.
func (pkg *Package) Name() string

// SetName sets the package name.
func (pkg *Package) SetName(name string)

// GoVersion returns the minimum Go version required by this package.
// If the minimum version is unknown, GoVersion returns the empty string.
// Individual source files may specify a different minimum Go version,
// as reported in the [go/ast.File.GoVersion] field.
func (pkg *Package) GoVersion() string

// Scope returns the (complete or incomplete) package scope
// holding the objects declared at package level (TypeNames,
// Consts, Vars, and Funcs).
// For a nil pkg receiver, Scope returns the Universe scope.
func (pkg *Package) Scope() *Scope

// A package is complete if its scope contains (at least) all
// exported objects; otherwise it is incomplete.
func (pkg *Package) Complete() bool

// MarkComplete marks a package as complete.
func (pkg *Package) MarkComplete()

// Imports returns the list of packages directly imported by
// pkg; the list is in source order.
//
// If pkg was loaded from export data, Imports includes packages that
// provide package-level objects referenced by pkg. This may be more or
// less than the set of packages directly imported by pkg's source code.
//
// If pkg uses cgo and the FakeImportC configuration option
// was enabled, the imports list may contain a fake "C" package.
func (pkg *Package) Imports() []*Package

// SetImports sets the list of explicitly imported packages to list.
// It is the caller's responsibility to make sure list elements are unique.
func (pkg *Package) SetImports(list []*Package)

func (pkg *Package) String() string
