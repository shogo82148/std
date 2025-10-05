// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the RISCV64
// instruction set, to minimize its interaction with the core of the
// assembler.

package arch

import (
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/obj/riscv"
)

// IsRISCV64AMO reports whether op is an AMO instruction that requires
// special handling.
func IsRISCV64AMO(op obj.As) bool

// IsRISCV64VTypeI reports whether op is a vtype immediate instruction that
// requires special handling.
func IsRISCV64VTypeI(op obj.As) bool

// RISCV64SpecialOperand returns the internal representation of a special operand.
func RISCV64SpecialOperand(name string) riscv.SpecialOperand

// RISCV64ValidateVectorType reports whether the given configuration is a
// valid vector type.
func RISCV64ValidateVectorType(vsew, vlmul, vtail, vmask int64) error
