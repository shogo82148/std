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
// map of valid formats (knownFormats). This map can be used to
// automatically rewrite format strings across all compiler
// files with the -r flag.
//
// The format map needs to be updated whenever a new (type,
// format) combination is found and the format verb is not
// 'v' or 'T' (as in "%v" or "%T"). To update the map auto-
// matically from the compiler source's use of format strings,
// use the -u flag. (Whether formats are valid for the values
// to be formatted must be verified manually, of course.)
//
// The -v flag prints out the names of all functions called
// with a format string, the names of files that were not
// processed, and any format rewrites made (with -r).
//
// Run as: go test -run Formats [-r][-u][-v]
//
// Known shortcomings:
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

// ignoredPackages is the set of packages which can
// be ignored.

// ignoredFunctions is the set of functions which may have
// format-like arguments but which don't do any formatting and
// thus may be ignored.
