// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package midway

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/cmd/compile/internal/types2"
)

type Rewriter struct {
	pkg          *types2.Package
	analyzer     *Analyzer
	info         *types2.Info
	sizes        []int
	blankCounter int
}

func NewRewriter(pkg *types2.Package, info *types2.Info, analyzer *Analyzer, sizes []int) *Rewriter

func (r *Rewriter) Rewrite(files []*syntax.File)

// RewriteWrapper applies the midway rewrite to files, which are the type
// checked files of pkg.
//
// If any rewriting was performed, it returns a non-nil InitOrderPlan capturing
// the pre-rewrite Go-spec initialization order, which the caller must Apply to
// the re-checked types2.Info after typechecking finishes.
//
// It also returns the names of the package-level declarations that are
// originals (see the comment at the top of this file). Every method of an
// original type is also an original.
func RewriteWrapper(pkg *types2.Package, info *types2.Info, files []*syntax.File) (originals []string, initOrderPlan *InitOrderPlan)
