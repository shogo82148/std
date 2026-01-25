// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflectdata

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

// EqFor returns ONAME node represents type t's equal function, and a boolean
// to indicates whether a length needs to be passed when calling the function.
func EqFor(t *types.Type) (ir.Node, bool)
