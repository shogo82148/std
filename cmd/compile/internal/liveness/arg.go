// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package liveness

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/objw"
	"github.com/shogo82148/std/cmd/compile/internal/ssa"
)

// ArgLiveness computes the liveness information of register argument spill slots.
// An argument's spill slot is "live" if we know it contains a meaningful value,
// that is, we have stored the register value to it.
// Returns the liveness map indices at each Block entry and at each Value (where
// it changes).
func ArgLiveness(fn *ir.Func, f *ssa.Func, pp *objw.Progs) (blockIdx, valueIdx map[ssa.ID]int)
