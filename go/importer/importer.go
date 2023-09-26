// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package importer provides access to export data importers.
package importer

import (
	"github.com/shogo82148/std/go/types"
	"github.com/shogo82148/std/io"
)

// A Lookup function returns a reader to access package data for
// a given import path, or an error if no matching package is found.
type Lookup func(path string) (io.ReadCloser, error)

// For returns an Importer for importing from installed packages
// for the compilers "gc" and "gccgo", or for importing directly
// from the source if the compiler argument is "source". In this
// latter case, importing may fail under circumstances where the
// exported API is not entirely defined in pure Go source code
// (if the package API depends on cgo-defined entities, the type
// checker won't have access to those).
//
// If lookup is nil, the default package lookup mechanism for the
// given compiler is used, and the resulting importer attempts
// to resolve relative and absolute import paths to canonical
// import path IDs before finding the imported file.
//
// If lookup is non-nil, then the returned importer calls lookup
// each time it needs to resolve an import path. In this mode
// the importer can only be invoked with canonical import paths
// (not relative or absolute ones); it is assumed that the translation
// to canonical import paths is being done by the client of the
// importer.
func For(compiler string, lookup Lookup) types.Importer

// Default returns an Importer for the compiler that built the running binary.
// If available, the result implements types.ImporterFrom.
func Default() types.Importer
