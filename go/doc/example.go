// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Extract example functions from file ASTs.

package doc

import (
	"github.com/shogo82148/std/go/ast"
)

// An Example represents an example function found in a test source file.
type Example struct {
	Name        string
	Suffix      string
	Doc         string
	Code        ast.Node
	Play        *ast.File
	Comments    []*ast.CommentGroup
	Output      string
	Unordered   bool
	EmptyOutput bool
	Order       int
}

// Examples returns the examples found in testFiles, sorted by Name field.
// The Order fields record the order in which the examples were encountered.
// The Suffix field is not populated when Examples is called directly, it is
// only populated by [NewFromFiles] for examples it finds in _test.go files.
//
// Playable Examples must be in a package whose name ends in "_test".
// An Example is "playable" (the Play field is non-nil) in either of these
// circumstances:
//   - The example function is self-contained: the function references only
//     identifiers from other packages (or predeclared identifiers, such as
//     "int") and the test file does not include a dot import.
//   - The entire test file is the example: the file contains exactly one
//     example function, zero test, fuzz test, or benchmark function, and at
//     least one top-level function, type, variable, or constant declaration
//     other than the example function.
func Examples(testFiles ...*ast.File) []*Example
