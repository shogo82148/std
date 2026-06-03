// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package midway

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/cmd/compile/internal/types2"
)

type Rewriter struct {
	pkg      *types2.Package
	analyzer *Analyzer
	info     *types2.Info
	sizes    []int
}

func NewRewriter(pkg *types2.Package, info *types2.Info, analyzer *Analyzer, sizes []int) *Rewriter

func (r *Rewriter) Rewrite(files []*syntax.File)

// Generate an API matching the standalone compilation call
func RewriteWrapper(pkg *types2.Package, info *types2.Info, files []*syntax.File) bool
