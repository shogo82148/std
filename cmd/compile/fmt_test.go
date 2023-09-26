// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements TestFormats; a test that verifies
// format strings in the compiler (this directory and all
// subdirectories, recursively).
//
// TestFormats finds potential (Printf, etc.) format strings.
// If they are used in a call, the format verbs are verified
// based on the matching argument type against a precomputed
// table of valid formats. The knownFormats table can be used
// to automatically rewrite format strings with the -u flag.
//
// A new knownFormats table based on the found formats is printed
// when the test is run in verbose mode (-v flag). The table
// needs to be updated whenever a new (type, format) combination
// is found and the format verb is not 'v' or 'T' (as in "%v" or
// "%T").
//
// Run as: go test -run Formats [-u][-v]
//
// Known bugs:
//   - indexed format strings ("%[2]s", etc.) are not supported
//     (the test will fail)
//   - format strings that are not simple string literals cannot
//     be updated automatically
//     (the test will fail with respective warnings)
//   - format strings in _test packages outside the current
//     package are not processed
//     (the test will report those files)
package main_test

import (
	"go/ast"
)

// The following variables collect information across all processed files.

// A File is a corresponding (filename, ast) pair.
type File struct {
	name string
	ast  *ast.File
}

// A callSite describes a function call that appears to contain
// a format string.

// blacklistedPackages is the set of packages which can
// be ignored.

// blacklistedFunctions is the set of functions which may have
// format-like arguments but which don't do any formatting and
// thus may be ignored.

// knownFormats entries are of the form "typename format" -> "newformat".
// An absent entry means that the format is not recognized as valid.
// An empty new format means that the format should remain unchanged.
// To print out a new table, run: go test -run Formats -v.
