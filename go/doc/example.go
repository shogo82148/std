// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Extract example functions from file ASTs.

package doc

import (
	"github.com/shogo82148/std/go/ast"
)

type Example struct {
	Name     string
	Doc      string
	Code     ast.Node
	Comments []*ast.CommentGroup
	Output   string
}

func Examples(files ...*ast.File) []*Example
