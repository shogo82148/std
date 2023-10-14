// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the
// MIPS (MIPS64) instruction set, to minimize its interaction
// with the core of the assembler.

package arch

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

// IsMIPSCMP reports whether the op (as defined by an mips.A* constant) is
// one of the CMP instructions that require special handling.
func IsMIPSCMP(op obj.As) bool

// IsMIPSMUL reports whether the op (as defined by an mips.A* constant) is
// one of the MUL/DIV/REM/MADD/MSUB instructions that require special handling.
func IsMIPSMUL(op obj.As) bool
