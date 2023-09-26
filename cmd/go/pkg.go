// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/go/build"
	"github.com/shogo82148/std/os"
	pathpkg "path"
)

// A Package describes a single package found in a directory.
type Package struct {
	Dir           string `json:",omitempty"`
	ImportPath    string `json:",omitempty"`
	ImportComment string `json:",omitempty"`
	Name          string `json:",omitempty"`
	Doc           string `json:",omitempty"`
	Target        string `json:",omitempty"`
	Goroot        bool   `json:",omitempty"`
	Standard      bool   `json:",omitempty"`
	Stale         bool   `json:",omitempty"`
	Root          string `json:",omitempty"`
	ConflictDir   string `json:",omitempty"`

	GoFiles        []string `json:",omitempty"`
	CgoFiles       []string `json:",omitempty"`
	IgnoredGoFiles []string `json:",omitempty"`
	CFiles         []string `json:",omitempty"`
	CXXFiles       []string `json:",omitempty"`
	MFiles         []string `json:",omitempty"`
	HFiles         []string `json:",omitempty"`
	SFiles         []string `json:",omitempty"`
	SwigFiles      []string `json:",omitempty"`
	SwigCXXFiles   []string `json:",omitempty"`
	SysoFiles      []string `json:",omitempty"`

	CgoCFLAGS    []string `json:",omitempty"`
	CgoCPPFLAGS  []string `json:",omitempty"`
	CgoCXXFLAGS  []string `json:",omitempty"`
	CgoLDFLAGS   []string `json:",omitempty"`
	CgoPkgConfig []string `json:",omitempty"`

	Imports []string `json:",omitempty"`
	Deps    []string `json:",omitempty"`

	Incomplete bool            `json:",omitempty"`
	Error      *PackageError   `json:",omitempty"`
	DepsErrors []*PackageError `json:",omitempty"`

	TestGoFiles  []string `json:",omitempty"`
	TestImports  []string `json:",omitempty"`
	XTestGoFiles []string `json:",omitempty"`
	XTestImports []string `json:",omitempty"`

	build        *build.Package
	pkgdir       string
	imports      []*Package
	deps         []*Package
	gofiles      []string
	sfiles       []string
	allgofiles   []string
	target       string
	fake         bool
	forceBuild   bool
	forceLibrary bool
	cmdline      bool
	local        bool
	localPrefix  string
	exeName      string
	coverMode    string
	coverVars    map[string]*CoverVar
	omitDWARF    bool
}

// CoverVar holds the name of the generated coverage variables targeting the named file.
type CoverVar struct {
	File string
	Var  string
}

// A PackageError describes an error loading information about a package.
type PackageError struct {
	ImportStack   []string
	Pos           string
	Err           string
	isImportCycle bool
	hard          bool
}

func (p *PackageError) Error() string

// An importStack is a stack of import paths.

// packageCache is a lookup cache for loadPackage,
// so that if we look up a package multiple times
// we return the same pointer each time.

// goTools is a map of Go program import path to install target directory.

// The runtime version string takes one of two forms:
// "go1.X[.Y]" for Go releases, and "devel +hash" at tip.
// Determine whether we are in a released copy by
// inspecting the version.

var _ = os.Getwd()
