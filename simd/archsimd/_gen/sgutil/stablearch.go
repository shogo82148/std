// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sgutil

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/go/build/constraint"
	"github.com/shogo82148/std/go/token"
)

// BuildConstraints are the build constraints of a SIMD source file: its
// //go:build line, and the GOARCH that its file name implies, if any.
type BuildConstraints struct {
	expr     constraint.Expr
	nameArch string

	phase          Phase
	stableArchOnly bool

	free []string
}

// ParseBuildConstraints returns the build constraints of file, which must have
// been parsed with comments using fset. It's an error if the file never
// builds, or if its //go:build line is too complicated to analyze.
func ParseBuildConstraints(fset *token.FileSet, file *ast.File) (*BuildConstraints, error)

// Phase returns the phase whose GOEXPERIMENT the file needs: [PhaseStable] if
// it builds with GOEXPERIMENT=simdstable, [PhaseExperiment] if it builds only
// with GOEXPERIMENT=simd (which implies simdstable), or 0 if it builds with
// neither.
func (c *BuildConstraints) Phase() Phase

// StableArchOnly reports whether, with the GOEXPERIMENT of its phase, the file
// only builds for architectures whose archsimd API can be generally
// available (amd64, arm64 and wasm). Only such a file may declare generally
// available archsimd API.
func (c *BuildConstraints) StableArchOnly() bool
