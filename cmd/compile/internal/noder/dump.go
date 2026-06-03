// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package noder

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/cmd/compile/internal/types2"
)

// MatchASTDump returns true if the fn matches the value
// of the astdump debug flag.
func MatchASTDump(fn *syntax.FuncDecl) bool

// DumpNodeHTML dumps the node n to the HTML writer for fn.
func DumpNodeHTML(pkg *types2.Package, file *syntax.File, info *types2.Info, fn *syntax.FuncDecl, why string, n syntax.Node)

// CloseHTMLWriters closes the HTML writer for fn, if one exists.
func CloseHTMLWriters()
