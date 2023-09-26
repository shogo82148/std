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

// For returns an Importer for the given compiler and lookup interface,
// or nil. Supported compilers are "gc", and "gccgo". If lookup is nil,
// the default package lookup mechanism for the given compiler is used.
// BUG(issue13847): For does not support non-nil lookup functions.
func For(compiler string, lookup Lookup) types.Importer

// Default returns an Importer for the compiler that built the running binary.
// If available, the result implements types.ImporterFrom.
func Default() types.Importer
