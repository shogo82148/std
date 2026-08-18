// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/compile/internal/abi"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/ssa/ssabase"
	"github.com/shogo82148/std/cmd/compile/internal/ssa/ssaop"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/src"
)

type (
	BlockRewriter func(*Block) bool
	ValueRewriter func(*Value) bool
)

// A Config holds readonly compilation information.
// It is created once, early during compilation,
// and shared across all compilations.
type Config struct {
	Arch           string
	PtrSize        int64
	RegSize        int64
	Types          Types
	LowerBlock     BlockRewriter
	LowerValue     ValueRewriter
	LateLowerBlock BlockRewriter
	LateLowerValue ValueRewriter
	SplitLoad      ValueRewriter
	Registers      []ssabase.Register
	GpRegMask      ssaop.RegMask
	FpRegMask      ssaop.RegMask
	Fp32RegMask    ssaop.RegMask
	Fp64RegMask    ssaop.RegMask
	SimdRegMask    ssaop.RegMask
	SpecialRegMask ssaop.RegMask
	IntParamRegs   []int8
	FloatParamRegs []int8
	ABI1           *abi.ABIConfig
	ABI0           *abi.ABIConfig
	FPReg          int8
	LinkReg        int8
	HasGReg        bool
	Ctxt           *obj.Link
	Optimize       bool
	SoftFloat      bool
	Race           bool
	BigEndian      bool
	UnalignedOK    bool
	HaveBswap64    bool
	HaveBswap32    bool
	HaveBswap16    bool
	HaveCondSelect bool

	// MulRecipes[x] = function to build v * x from v.
	MulRecipes map[int64]mulRecipe
}

type Frontend interface {
	Logger

	StringData(string) *obj.LSym

	SplitSlot(parent *LocalSlot, suffix string, offset int64, t *types.Type) LocalSlot

	Syslook(string) *obj.LSym

	UseWriteBarrier() bool

	Func() *ir.Func
}

type Logger interface {
	Logf(string, ...any)

	Log() bool

	Fatalf(pos src.XPos, msg string, args ...any)

	Warnl(pos src.XPos, fmt_ string, args ...any)

	Debug_checknil() bool
}

type Types struct {
	Bool       *types.Type
	Int8       *types.Type
	Int16      *types.Type
	Int32      *types.Type
	Int64      *types.Type
	UInt8      *types.Type
	UInt16     *types.Type
	UInt32     *types.Type
	UInt64     *types.Type
	Int        *types.Type
	Float32    *types.Type
	Float64    *types.Type
	UInt       *types.Type
	Uintptr    *types.Type
	String     *types.Type
	BytePtr    *types.Type
	Int32Ptr   *types.Type
	UInt32Ptr  *types.Type
	IntPtr     *types.Type
	UintptrPtr *types.Type
	Float32Ptr *types.Type
	Float64Ptr *types.Type
	BytePtrPtr *types.Type
	Vec128     *types.Type
	Vec256     *types.Type
	Vec512     *types.Type
	Mask       *types.Type
}

// NewTypes creates and populates a Types.
func NewTypes() *Types

// SetTypPtrs populates t.
func (t *Types) SetTypPtrs()

func (c *Config) HaveByteSwap(size int64) bool

func (c *Config) BuildRecipes(arch string)
