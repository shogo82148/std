// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssagen

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/pgoir"
	"github.com/shogo82148/std/cmd/compile/internal/ssa"
	"github.com/shogo82148/std/cmd/internal/obj"
)

// Compile builds an SSA backend function,
// uses it to generate a plist,
// and flushes that plist to machine code.
// worker indicates which of the backend workers is doing the processing.
func Compile(fn *ir.Func, worker int, profile *pgoir.Profile)

// RegisterMapInitLsym records "s" in the set of outlined map initializer
// functions.
func RegisterMapInitLsym(s *obj.LSym)

// StackOffset returns the stack location of a LocalSlot relative to the
// stack pointer, suitable for use in a DWARF location entry. This has nothing
// to do with its offset in the user variable.
func StackOffset(slot ssa.LocalSlot) int32

func CheckLargeStacks()
