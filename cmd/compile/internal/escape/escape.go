// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package escape

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

func Funcs(all []ir.Node)

// Batch performs escape analysis on a minimal batch of
// functions.
func Batch(fns []*ir.Func, recursive bool)
