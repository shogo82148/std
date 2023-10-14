// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the ARM
// instruction set, to minimize its interaction with the core of the
// assembler.

package arch

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

// IsARMCMP reports whether the op (as defined by an arm.A* constant) is
// one of the comparison instructions that require special handling.
func IsARMCMP(op obj.As) bool

// IsARMSTREX reports whether the op (as defined by an arm.A* constant) is
// one of the STREX-like instructions that require special handling.
func IsARMSTREX(op obj.As) bool

// IsARMMRC reports whether the op (as defined by an arm.A* constant) is
// MRC or MCR.
func IsARMMRC(op obj.As) bool

// IsARMBFX reports whether the op (as defined by an arm.A* constant) is one the
// BFX-like instructions which are in the form of "op $width, $LSB, (Reg,) Reg".
func IsARMBFX(op obj.As) bool

// IsARMFloatCmp reports whether the op is a floating comparison instruction.
func IsARMFloatCmp(op obj.As) bool

// ARMMRCOffset implements the peculiar encoding of the MRC and MCR instructions.
// The difference between MRC and MCR is represented by a bit high in the word, not
// in the usual way by the opcode itself. Asm must use AMRC for both instructions, so
// we return the opcode for MRC so that asm doesn't need to import obj/arm.
func ARMMRCOffset(op obj.As, cond string, x0, x1, x2, x3, x4, x5 int64) (offset int64, op0 obj.As, ok bool)

// IsARMMULA reports whether the op (as defined by an arm.A* constant) is
// MULA, MULS, MMULA, MMULS, MULABB, MULAWB or MULAWT, the 4-operand instructions.
func IsARMMULA(op obj.As) bool

// ARMConditionCodes handles the special condition code situation for the ARM.
// It returns a boolean to indicate success; failure means cond was unrecognized.
func ARMConditionCodes(prog *obj.Prog, cond string) bool

// ParseARMCondition parses the conditions attached to an ARM instruction.
// The input is a single string consisting of period-separated condition
// codes, such as ".P.W". An initial period is ignored.
func ParseARMCondition(cond string) (uint8, bool)
