// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Extract example functions from file ASTs.

package doc

import (
	"github.com/shogo82148/std/go/ast"
)

// An Example represents an example function found in a source files.
type Example struct {
	Name        string
	Doc         string
	Code        ast.Node
	Play        *ast.File
	Comments    []*ast.CommentGroup
	Output      string
	EmptyOutput bool
	Order       int
}

// Examples returns the examples found in the files, sorted by Name field.
// The Order fields record the order in which the examples were encountered.
func Examples(files ...*ast.File) []*Example
