// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package midway

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/cmd/compile/internal/types2"
)

// Analyzer holds the state for SIMD dependency analysis
type Analyzer struct {
	pkg          *types2.Package
	info         *types2.Info
	dependentObj map[types2.Object]bool
	visited      map[types2.Type]bool
	inSimd       bool
}

func NewAnalyzer(pkg *types2.Package, info *types2.Info) *Analyzer

// Analyze builds the set of SIMD-dependent objects
func (a *Analyzer) Analyze(files []*syntax.File) bool

func (a *Analyzer) HasDependentSignature(sig *types2.Signature) bool
