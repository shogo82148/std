// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modload

import (
	"github.com/shogo82148/std/context"

	"github.com/shogo82148/std/cmd/go/internal/gover"
	"github.com/shogo82148/std/cmd/go/internal/search"

	"golang.org/x/mod/module"
)

// PackageOpts control the behavior of the LoadPackages function.
type PackageOpts struct {
	// TidyGoVersion is the Go version to which the go.mod file should be updated
	// after packages have been loaded.
	//
	// An empty TidyGoVersion means to use the Go version already specified in the
	// main module's go.mod file, or the latest Go version if there is no main
	// module.
	TidyGoVersion string

	// Tags are the build tags in effect (as interpreted by the
	// cmd/go/internal/imports package).
	// If nil, treated as equivalent to imports.Tags().
	Tags map[string]bool

	// Tidy, if true, requests that the build list and go.sum file be reduced to
	// the minimal dependencies needed to reproducibly reload the requested
	// packages.
	Tidy bool

	// TidyDiff, if true, analyzes the necessary changes to go.mod and go.sum
	// to make them tidy. It does not modify these files, but exits with
	// a non-zero code if updates are needed.
	TidyDiff bool

	// TidyCompatibleVersion is the oldest Go version that must be able to
	// reproducibly reload the requested packages.
	//
	// If empty, the compatible version is the Go version immediately prior to the
	// 'go' version listed in the go.mod file.
	TidyCompatibleVersion string

	// VendorModulesInGOROOTSrc indicates that if we are within a module in
	// GOROOT/src, packages in the module's vendor directory should be resolved as
	// actual module dependencies (instead of standard-library packages).
	VendorModulesInGOROOTSrc bool

	// ResolveMissingImports indicates that we should attempt to add module
	// dependencies as needed to resolve imports of packages that are not found.
	//
	// For commands that support the -mod flag, resolving imports may still fail
	// if the flag is set to "readonly" (the default) or "vendor".
	ResolveMissingImports bool

	// AssumeRootsImported indicates that the transitive dependencies of the root
	// packages should be treated as if those roots will be imported by the main
	// module.
	AssumeRootsImported bool

	// AllowPackage, if non-nil, is called after identifying the module providing
	// each package. If AllowPackage returns a non-nil error, that error is set
	// for the package, and the imports and test of that package will not be
	// loaded.
	//
	// AllowPackage may be invoked concurrently by multiple goroutines,
	// and may be invoked multiple times for a given package path.
	AllowPackage func(ctx context.Context, path string, mod module.Version) error

	// LoadTests loads the test dependencies of each package matching a requested
	// pattern. If ResolveMissingImports is also true, test dependencies will be
	// resolved if missing.
	LoadTests bool

	// UseVendorAll causes the "all" package pattern to be interpreted as if
	// running "go mod vendor" (or building with "-mod=vendor").
	//
	// This is a no-op for modules that declare 'go 1.16' or higher, for which this
	// is the default (and only) interpretation of the "all" pattern in module mode.
	UseVendorAll bool

	// AllowErrors indicates that LoadPackages should not terminate the process if
	// an error occurs.
	AllowErrors bool

	// SilencePackageErrors indicates that LoadPackages should not print errors
	// that occur while matching or loading packages, and should not terminate the
	// process if such an error occurs.
	//
	// Errors encountered in the module graph will still be reported.
	//
	// The caller may retrieve the silenced package errors using the Lookup
	// function, and matching errors are still populated in the Errs field of the
	// associated search.Match.)
	SilencePackageErrors bool

	// SilenceMissingStdImports indicates that LoadPackages should not print
	// errors or terminate the process if an imported package is missing, and the
	// import path looks like it might be in the standard library (perhaps in a
	// future version).
	SilenceMissingStdImports bool

	// SilenceNoGoErrors indicates that LoadPackages should not print
	// imports.ErrNoGo errors.
	// This allows the caller to invoke LoadPackages (and report other errors)
	// without knowing whether the requested packages exist for the given tags.
	//
	// Note that if a requested package does not exist *at all*, it will fail
	// during module resolution and the error will not be suppressed.
	SilenceNoGoErrors bool

	// SilenceUnmatchedWarnings suppresses the warnings normally emitted for
	// patterns that did not match any packages.
	SilenceUnmatchedWarnings bool

	// Resolve the query against this module.
	MainModule module.Version

	// If Switcher is non-nil, then LoadPackages passes all encountered errors
	// to Switcher.Error and tries Switcher.Switch before base.ExitIfErrors.
	Switcher gover.Switcher
}

// LoadPackages identifies the set of packages matching the given patterns and
// loads the packages in the import graph rooted at that set.
func LoadPackages(ctx context.Context, opts PackageOpts, patterns ...string) (matches []*search.Match, loadedPackages []string)

// ImportFromFiles adds modules to the build list as needed
// to satisfy the imports in the named Go source files.
//
// Errors in missing dependencies are silenced.
//
// TODO(bcmills): Silencing errors seems off. Take a closer look at this and
// figure out what the error-reporting actually ought to be.
func ImportFromFiles(ctx context.Context, gofiles []string)

// DirImportPath returns the effective import path for dir,
// provided it is within a main module, or else returns ".".
func (mms *MainModuleSet) DirImportPath(ctx context.Context, dir string) (path string, m module.Version)

// PackageModule returns the module providing the package named by the import path.
func PackageModule(path string) module.Version

// Lookup returns the source directory, import path, and any loading error for
// the package at path as imported from the package in parentDir.
// Lookup requires that one of the Load functions in this package has already
// been called.
func Lookup(parentPath string, parentIsStd bool, path string) (dir, realPath string, err error)

// Why returns the "go mod why" output stanza for the given package,
// without the leading # comment.
// The package graph must have been loaded already, usually by LoadPackages.
// If there is no reason for the package to be in the current build,
// Why returns an empty string.
func Why(path string) string

// WhyDepth returns the number of steps in the Why listing.
// If there is no reason for the package to be in the current build,
// WhyDepth returns 0.
func WhyDepth(path string) int
