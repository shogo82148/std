// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file encapsulates some of the odd characteristics of the ARM64
// instruction set, to minimize its interaction with the core of the
// assembler.

package arch

import (
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/obj/arm64"
)

// GetARM64SpecialOperand returns the internal representation of a special operand.
func GetARM64SpecialOperand(name string) arm64.SpecialOperand

// IsARM64ADR reports whether the op (as defined by an arm64.A* constant) is
// one of the comparison instructions that require special handling.
func IsARM64ADR(op obj.As) bool

// IsARM64CMP reports whether the op (as defined by an arm64.A* constant) is
// one of the comparison instructions that require special handling.
func IsARM64CMP(op obj.As) bool

// IsARM64STLXR reports whether the op (as defined by an arm64.A*
// constant) is one of the STLXR-like instructions that require special
// handling.
func IsARM64STLXR(op obj.As) bool

// IsARM64TBL reports whether the op (as defined by an arm64.A*
// constant) is one of the TBL-like instructions and one of its
// inputs does not fit into prog.Reg, so require special handling.
func IsARM64TBL(op obj.As) bool

// IsARM64CASP reports whether the op (as defined by an arm64.A*
// constant) is one of the CASP-like instructions, and its 2nd
// destination is a register pair that require special handling.
func IsARM64CASP(op obj.As) bool

// ARM64Suffix handles the special suffix for the ARM64.
// It returns a boolean to indicate success; failure means
// cond was unrecognized.
func ARM64Suffix(prog *obj.Prog, cond string) bool

// ARM64RegisterShift constructs an ARM64 register with shift operation.
func ARM64RegisterShift(reg, op, count int16) (int64, error)

// ARM64RegisterExtension constructs an ARM64 register with extension or arrangement.
func ARM64RegisterExtension(a *obj.Addr, ext string, reg, num int16, isAmount, isIndex bool) error

// ARM64RegisterArrangement constructs an ARM64 vector register arrangement.
func ARM64RegisterArrangement(reg int16, name, arng string) (int64, error)

// ARM64RegisterListOffset generates offset encoding according to AArch64 specification.
func ARM64RegisterListOffset(firstReg, regCnt int, arrangement int64) (int64, error)
