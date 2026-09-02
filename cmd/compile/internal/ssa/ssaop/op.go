// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssaop

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

type AuxType int8

const (
	AuxTypeNone AuxType = iota
	AuxTypeBool
	AuxTypeInt8
	AuxTypeInt16
	AuxTypeInt32
	AuxTypeInt64
	AuxTypeInt128
	AuxTypeUInt8
	AuxTypeFloat32
	AuxTypeFloat64
	AuxTypeFlagConstant
	AuxTypeCCop
	AuxTypeNameOffsetInt8
	AuxTypeString
	AuxTypeSym
	AuxTypeSymOff
	AuxTypeSymValAndOff
	AuxTypeTyp
	AuxTypeTypSize
	AuxTypeCall
	AuxTypeCallOff

	AuxTypePanicBoundsC
	AuxTypePanicBoundsCC

	// architecture specific aux types
	AuxTypeARM64BitField
	AuxTypeARM64ConditionalParams
	AuxTypeS390XRotateParams
	AuxTypeS390XCCMask
	AuxTypeS390XCCMaskInt8
	AuxTypeS390XCCMaskUint8
	AuxTypeSizeAndAlign
)

// An Op encodes the specific operation that a Value performs.
// Opcodes' semantics can be modified by the type and aux fields of the Value.
// For instance, OpAdd can be 32 or 64 bit, signed or unsigned, float or complex, depending on Value.Type.
// Semantics of each op are described in the opcode files in _gen/*Ops.go.
// There is one file for generic (architecture-independent) ops and one file
// for each architecture.
type Op int32

type OpInfo struct {
	Name              string
	Reg               RegInfo
	AuxType           AuxType
	ArgLen            int32
	asm               obj.As
	Generic           bool
	Rematerializeable bool
	Commutative       bool
	ResultInArg0      bool
	ResultNotInArgs   bool
	ClobberFlags      bool
	NeedIntTemp       bool
	Call              bool
	tailCall          bool
	NilCheck          bool
	FaultOnNilArg0    bool
	FaultOnNilArg1    bool
	usesScratch       bool
	HasSideEffects    bool
	ZeroWidth         bool
	unsafePoint       bool
	FixedReg          bool
	EarlyOk           bool
	AddrSinkArg0      bool
	AddrSinkArg1      bool
	symEffect         SymEffect
	scale             uint8
	ZeroUpperBits     uint8
}

type OutputInfo struct {
	Idx  int
	Regs RegMask
}

type RegInfo struct {
	// Inputs encodes the register restrictions for an instruction's Inputs.
	// Each entry specifies an allowed register set for a particular input.
	// They are listed in the order in which regalloc should pick a register
	// from the register set (most constrained first).
	// Inputs which do not need registers are not listed.
	Inputs []InputInfo
	// Clobbers encodes the set of registers that are overwritten by
	// the instruction (other than the output registers).
	Clobbers RegMask
	// Instruction clobbers the register containing input 0.
	ClobbersArg0 bool
	// Instruction clobbers the register containing input 1.
	ClobbersArg1 bool
	// Outputs is the same as inputs, but for the Outputs of the instruction.
	Outputs []OutputInfo
}

// A SymEffect describes the effect that an SSA Value has on the variable
// identified by the symbol in its Aux field.
type SymEffect int8

const (
	SymRead SymEffect = 1 << iota
	SymWrite
	SymAddr

	SymRdWr = SymRead | SymWrite

	SymNone SymEffect = 0
)

type InputInfo struct {
	Idx  int
	Regs RegMask
}

func (r *RegInfo) String() string

func (op Op) IsLoweredGetClosurePtr() bool
