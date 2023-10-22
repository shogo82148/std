// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/compile/internal/abi"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/src"
)

// A Config holds readonly compilation information.
// It is created once, early during compilation,
// and shared across all compilations.
type Config struct {
	arch           string
	PtrSize        int64
	RegSize        int64
	Types          Types
	lowerBlock     blockRewriter
	lowerValue     valueRewriter
	lateLowerBlock blockRewriter
	lateLowerValue valueRewriter
	splitLoad      valueRewriter
	registers      []Register
	gpRegMask      regMask
	fpRegMask      regMask
	fp32RegMask    regMask
	fp64RegMask    regMask
	specialRegMask regMask
	intParamRegs   []int8
	floatParamRegs []int8
	ABI1           *abi.ABIConfig
	ABI0           *abi.ABIConfig
	GCRegMap       []*Register
	FPReg          int8
	LinkReg        int8
	hasGReg        bool
	ctxt           *obj.Link
	optimize       bool
	noDuffDevice   bool
	useSSE         bool
	useAvg         bool
	useHmul        bool
	SoftFloat      bool
	Race           bool
	BigEndian      bool
	UseFMA         bool
	unalignedOK    bool
	haveBswap64    bool
	haveBswap32    bool
	haveBswap16    bool
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
}

// NewTypes creates and populates a Types.
func NewTypes() *Types

// SetTypPtrs populates t.
func (t *Types) SetTypPtrs()

type Logger interface {
	Logf(string, ...interface{})

	Log() bool

	Fatalf(pos src.XPos, msg string, args ...interface{})

	Warnl(pos src.XPos, fmt_ string, args ...interface{})

	Debug_checknil() bool
}

type Frontend interface {
	Logger

	StringData(string) *obj.LSym

	SplitSlot(parent *LocalSlot, suffix string, offset int64, t *types.Type) LocalSlot

	Syslook(string) *obj.LSym

	UseWriteBarrier() bool

	Func() *ir.Func
}

// NewConfig returns a new configuration object for the given architecture.
func NewConfig(arch string, types Types, ctxt *obj.Link, optimize, softfloat bool) *Config

func (c *Config) Ctxt() *obj.Link
