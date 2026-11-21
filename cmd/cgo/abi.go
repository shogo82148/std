// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/shogo82148/std/io"
)

// BindingType describes a C or Go type used by cgo binding declarations
// for C ABI classification and Go wrapper generation.
type BindingType struct {
	Name  string
	Def   string
	Size  int64
	Align int64
}

// BindingFunc describes a bound C function declared with //cgo:func.
type BindingFunc struct {
	// Identifier and symbol names:
	Name   string
	GoName string
	CName  string
	BindC  bool

	// Signature (Go type expressions, e.g. "int32", "*byte", "_Ctype_S"):
	Params []string
	Ret    string

	// Go wrapper and import metadata (used by binding.go):
	Library string
	LibC    bool
	Static  bool
}

// ParamLoc describes where a function parameter is passed in the C ABI.
type ParamLoc struct {
	GoOffset  int64
	Size      int64
	IsFloat   bool
	Regs      []string
	StackOff  int64
	PassByPtr bool
}

// RetLoc describes how a function return value is returned in the C ABI.
type RetLoc struct {
	GoOffset int64
	Size     int64
	IsFloat  bool
	Indirect bool
	Regs     []string
}

// FuncLayout describes the complete Go argument struct and C ABI layout for a call.
type FuncLayout struct {
	Params        []ParamLoc
	Ret           RetLoc
	StackSize     int64
	FloatRegCount int
}

// ABIConfig defines the C calling convention rules and assembly primitives
// for a target OS/architecture. OS, PtrSize, CLongSize, StackAlign, and AssignParam must always be set.
type ABIConfig struct {
	Arch        string
	OS          string
	PtrSize     int64
	CLongSize   int64
	StackAlign  int64
	ShadowSpace int64

	IntRegs      []string
	FloatRegs    []string
	IntRetRegs   []string
	FloatRetRegs []string

	MaxRegStructSize int64
	IndirectRetReg   string
	IndirectRetInArg bool
	PackStackArgs    bool

	// AssignParam allocates a single parameter in registers or on the C stack.
	AssignParam func(cfg *ABIConfig, fn BindingFunc, idx int, info typeInfo, st *abiState) ParamLoc

	// AssignRet determines how the return value is returned in the C ABI.
	// ABIs without special return rules use defaultAssignRet.
	AssignRet func(cfg *ABIConfig, fn BindingFunc, info typeInfo, st *abiState) RetLoc

	// Optional platform-specific ABI hook. Returning handled=false falls back
	// to default GPR chunking.
	CompositeRegSlots func(cfg *ABIConfig, info typeInfo, isFloat bool, regs []string) ([]regSlot, bool)

	// Trampoline assembly generation parameters:
	SPReg          string
	CalleeSavedReg string
	TmpReg         string
	TmpReg2        string
	LRReserve      int64

	asm                asmOps
	TrampolinePrologue func(w io.Writer, lay FuncLayout) (epilogue func(w io.Writer))
	PreCall            func(cfg *ABIConfig, w io.Writer, fn BindingFunc, lay FuncLayout)
	CallSymbol         func(w io.Writer, sym string)
}
