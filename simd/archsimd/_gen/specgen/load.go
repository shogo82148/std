// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package specgen

import (
	"github.com/shogo82148/std/go/ast"
	"github.com/shogo82148/std/io"
)

type LoadOptions struct {
	// Filter, if non-nil, causes Load to process only spec functions satisfying
	// Filter.
	Filter func(*ast.FuncDecl) bool

	// Trace, if non-nil, causes Load to log a debug trace of solver steps to
	// Trace.
	Trace io.Writer
}

// Load loads a Go SIMD spec from the package in directory dir. This is the main
// entrypoint to this package.
func Load(dir string, opts *LoadOptions) ([]*Func, error)
