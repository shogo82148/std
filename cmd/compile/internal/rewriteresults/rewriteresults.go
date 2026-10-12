// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rewriteresults rewrites local variables returned directly by a
// function to use the corresponding result parameter's storage.
package rewriteresults

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// Funcs applies the rewriteresults pass to fns.
func Funcs(fns []*ir.Func)
