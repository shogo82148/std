// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The deadlocals pass removes assignments to unused local variables.
package deadlocals

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// Funcs applies the deadlocals pass to fns.
func Funcs(fns []*ir.Func)
