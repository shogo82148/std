// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package interleaved implements the interleaved devirtualization and
// inlining pass.
package interleaved

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/pgoir"
)

// DevirtualizeAndInlinePackage interleaves devirtualization and inlining on
// all functions within pkg.
func DevirtualizeAndInlinePackage(pkg *ir.Package, profile *pgoir.Profile)

// DevirtualizeAndInlineFunc interleaves devirtualization and inlining
// on a single function.
func DevirtualizeAndInlineFunc(fn *ir.Func, profile *pgoir.Profile)
