// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/compile/internal/abi"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/ssa/ssaop"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
)

// For ABI register index r, returns the (dense) register number used in
// SSA backend.
func ArchRegForAbiReg(r abi.RegIndex, c *Config) uint8

// Arm64BitField is the GO type of ARM64BitField auxInt.
// if x is an ARM64BitField, then width=x&0xff, lsb=(x>>8)&0xff, and
// width+lsb<64 for 64-bit variant, width+lsb<32 for 32-bit variant.
// the meaning of width and lsb are instruction-dependent.
type Arm64BitField int16

// Arm64ConditionalParams is the GO type of ARM64ConditionalParams auxInt.
type Arm64ConditionalParams struct {
	Cond     ssaop.Op
	NzcvVal  uint8
	ConstVal uint8
	Ind      bool
}

type AuxCall struct {
	Fn       *obj.LSym
	RegCache *ssaop.RegInfo
	AbiInfo  *abi.ABIParamResultInfo
}

type AuxNameOffset struct {
	Name   *ir.Name
	Offset int64
}

func MakeValAndOff(val, off int32) ValAndOff

// A ValAndOff is used by the several opcodes. It holds
// both a value and a pointer offset.
// A ValAndOff is intended to be encoded into an AuxInt field.
// The zero ValAndOff encodes a value of 0 and an offset of 0.
// The high 32 bits hold a value.
// The low 32 bits hold a pointer offset.
type ValAndOff int64

type BoundsKind uint8

const (
	BoundsIndex BoundsKind = iota
	BoundsIndexU
	BoundsSliceAlen
	BoundsSliceAlenU
	BoundsSliceAcap
	BoundsSliceAcapU
	BoundsSliceB
	BoundsSliceBU
	BoundsSlice3Alen
	BoundsSlice3AlenU
	BoundsSlice3Acap
	BoundsSlice3AcapU
	BoundsSlice3B
	BoundsSlice3BU
	BoundsSlice3C
	BoundsSlice3CU
	BoundsConvert
	BoundsKindCount
)

// For ABI register index r, returns the register number used in the obj
// package (assembler).
func ObjRegForAbiReg(r abi.RegIndex, c *Config) int16

// StaticAuxCall returns an AuxCall for a static call.
func StaticAuxCall(sym *obj.LSym, paramResultInfo *abi.ABIParamResultInfo) *AuxCall

// InterfaceAuxCall returns an AuxCall for an interface call.
func InterfaceAuxCall(paramResultInfo *abi.ABIParamResultInfo) *AuxCall

// ClosureAuxCall returns an AuxCall for a closure call.
func ClosureAuxCall(paramResultInfo *abi.ABIParamResultInfo) *AuxCall

// OwnAuxCall returns a function's own AuxCall.
func OwnAuxCall(fn *obj.LSym, paramResultInfo *abi.ABIParamResultInfo) *AuxCall

// A Sym represents a symbolic offset from a base register.
// Currently a Sym can be one of 3 things:
//   - a *ir.Name, for an offset from SP (the stack pointer)
//   - a *obj.LSym, for an offset from SB (the global pointer)
//   - nil, for no offset
type Sym interface {
	Aux
	CanBeAnSSASym()
}

func (a *AuxNameOffset) CanBeAnSSAAux()

func (a *AuxNameOffset) String() string

func (a *AuxNameOffset) FrameOffset() int64

// Reg returns the regInfo for a given call, combining the derived in/out register masks
// with the machine-specific register information in the input i.  (The machine-specific
// regInfo is much handier at the call site than it is when the AuxCall is being constructed,
// therefore do this lazily).
//
// TODO: there is a Clever Hack that allows pre-generation of a small-ish number of the slices
// of inputInfo and outputInfo used here, provided that we are willing to reorder the inputs
// and outputs from calls, so that all integer registers come first, then all floating registers.
// At this point (active development of register ABI) that is very premature,
// but if this turns out to be a cost, we could do it.
func (a *AuxCall) Reg(i *ssaop.RegInfo, c *Config) *ssaop.RegInfo

func (a *AuxCall) ABI() *abi.ABIConfig

func (a *AuxCall) ABIInfo() *abi.ABIParamResultInfo

func (a *AuxCall) ResultReg(c *Config) *ssaop.RegInfo

// ArgWidth returns the amount of stack needed for all the inputs
// and outputs of a function or method, including ABI-defined parameter
// slots and ABI-defined spill slots for register-resident parameters.
//
// The name is taken from the types package's ArgWidth(<function type>),
// which predated changes to the ABI; this version handles those changes.
func (a *AuxCall) ArgWidth() int64

// ParamAssignmentForResult returns the ABI Parameter assignment for result which (indexed 0, 1, etc).
func (a *AuxCall) ParamAssignmentForResult(which int64) *abi.ABIParamAssignment

// OffsetOfResult returns the SP offset of result which (indexed 0, 1, etc).
func (a *AuxCall) OffsetOfResult(which int64) int64

// OffsetOfArg returns the SP offset of argument which (indexed 0, 1, etc).
// If the call is to a method, the receiver is the first argument (i.e., index 0)
func (a *AuxCall) OffsetOfArg(which int64) int64

// RegsOfResult returns the register(s) used for result which (indexed 0, 1, etc).
func (a *AuxCall) RegsOfResult(which int64) []abi.RegIndex

// RegsOfArg returns the register(s) used for argument which (indexed 0, 1, etc).
// If the call is to a method, the receiver is the first argument (i.e., index 0)
func (a *AuxCall) RegsOfArg(which int64) []abi.RegIndex

// NameOfResult returns the ir.Name of result which (indexed 0, 1, etc).
func (a *AuxCall) NameOfResult(which int64) *ir.Name

// TypeOfResult returns the type of result which (indexed 0, 1, etc).
func (a *AuxCall) TypeOfResult(which int64) *types.Type

// TypeOfArg returns the type of argument which (indexed 0, 1, etc).
// If the call is to a method, the receiver is the first argument (i.e., index 0)
func (a *AuxCall) TypeOfArg(which int64) *types.Type

// SizeOfResult returns the size of result which (indexed 0, 1, etc).
func (a *AuxCall) SizeOfResult(which int64) int64

// SizeOfArg returns the size of argument which (indexed 0, 1, etc).
// If the call is to a method, the receiver is the first argument (i.e., index 0)
func (a *AuxCall) SizeOfArg(which int64) int64

// NResults returns the number of results.
func (a *AuxCall) NResults() int64

// LateExpansionResultType returns the result type (including trailing mem)
// for a call that will be expanded later in the SSA phase.
func (a *AuxCall) LateExpansionResultType() *types.Type

// NArgs returns the number of arguments (including receiver, if there is one).
func (a *AuxCall) NArgs() int64

// String returns "AuxCall{<fn>}"
func (a *AuxCall) String() string

func (*AuxCall) CanBeAnSSAAux()

func (x ValAndOff) Val() int32

func (x ValAndOff) Val64() int64

func (x ValAndOff) Val16() int16

func (x ValAndOff) Val8() int8

func (x ValAndOff) Off64() int64

func (x ValAndOff) Off() int32

func (x ValAndOff) String() string

func (x ValAndOff) CanAdd32(off int32) bool

func (x ValAndOff) CanAdd64(off int64) bool

func (x ValAndOff) AddOffset32(off int32) ValAndOff

func (x ValAndOff) AddOffset64(off int64) ValAndOff

// Returns the bounds error code needed by the runtime, and
// whether the x field is signed.
func (b BoundsKind) Code() (rtabi.BoundsErrorCode, bool)
