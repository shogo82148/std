// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package noder

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/cmd/compile/internal/types2"
)

// An HTMLWriter dumps syntax nodes to multicolumn HTML, similar to what the
// ssa backend does for GOSSAFUNC.
type HTMLWriter struct {
	ir.HTMLWriterBase

	Decl *syntax.FuncDecl
	pkg  *types2.Package
	file *syntax.File
	info *types2.Info
}

func NewHTMLWriter(pkg *types2.Package, file *syntax.File, info *types2.Info, path string, decl *syntax.FuncDecl, cfgMask string) *HTMLWriter

func (w *HTMLWriter) DeclHTML(phase string) func()
