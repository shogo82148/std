// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

// FileExports trims the AST for a Go source file in place such that
// only exported nodes remain: all top-level identifiers which are not exported
// and their associated information (such as type, initial value, or function
// body) are removed. Non-exported fields and methods of exported types are
// stripped. The [File.Comments] list is not changed.
//
// FileExports reports whether there are exported declarations.
func FileExports(src *File) bool

// PackageExports trims the AST for a Go package in place such that
// only exported nodes remain. The pkg.Files list is not changed, so that
// file names and top-level package comments don't get lost.
//
// PackageExports reports whether there are exported declarations;
// it returns false otherwise.
//
// Deprecated: use the type checker [go/types] instead of [Package];
// see [Object]. Alternatively, use [FileExports].
func PackageExports(pkg *Package) bool

type Filter func(string) bool

// FilterDecl trims the AST for a Go declaration in place by removing
// all names (including struct field and interface method names, but
// not from parameter lists) that don't pass through the filter f.
//
// FilterDecl reports whether there are any declared names left after
// filtering.
func FilterDecl(decl Decl, f Filter) bool

// FilterFile trims the AST for a Go file in place by removing all
// names from top-level declarations (including struct field and
// interface method names, but not from parameter lists) that don't
// pass through the filter f. If the declaration is empty afterwards,
// the declaration is removed from the AST. Import declarations are
// always removed. The [File.Comments] list is not changed.
//
// FilterFile reports whether there are any top-level declarations
// left after filtering.
func FilterFile(src *File, f Filter) bool

// FilterPackage trims the AST for a Go package in place by removing
// all names from top-level declarations (including struct field and
// interface method names, but not from parameter lists) that don't
// pass through the filter f. If the declaration is empty afterwards,
// the declaration is removed from the AST. The pkg.Files list is not
// changed, so that file names and top-level package comments don't get
// lost.
//
// FilterPackage reports whether there are any top-level declarations
// left after filtering.
//
// Deprecated: use the type checker [go/types] instead of [Package];
// see [Object]. Alternatively, use [FilterFile].
func FilterPackage(pkg *Package, f Filter) bool

// The MergeMode flags control the behavior of [MergePackageFiles].
//
// Deprecated: use the type checker [go/types] instead of [Package];
// see [Object].
type MergeMode uint

// Deprecated: use the type checker [go/types] instead of [Package];
// see [Object].
const (
	// If set, duplicate function declarations are excluded.
	FilterFuncDuplicates MergeMode = 1 << iota
	// If set, comments that are not associated with a specific
	// AST node (as Doc or Comment) are excluded.
	FilterUnassociatedComments
	// If set, duplicate import declarations are excluded.
	FilterImportDuplicates
)

// MergePackageFiles creates a file AST by merging the ASTs of the
// files belonging to a package. The mode flags control merging behavior.
//
// Deprecated: this function is poorly specified and has unfixable
// bugs; also [Package] is deprecated.
func MergePackageFiles(pkg *Package, mode MergeMode) *File
