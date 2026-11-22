// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/go/ast"
)

// ExportFunc describes an //export Go function callable from C in a cgo
// binding package.
type ExportFunc struct {
	Name   string
	Params []string
	Ret    string
	Func   *ast.FuncDecl
	Imp    fileImports
}
