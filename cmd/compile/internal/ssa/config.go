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
	// Logf logs a message from the compiler.
	Logf(string, ...interface{})

	// Log reports whether logging is not a no-op
	// some logging calls account for more than a few heap allocations.
	Log() bool

	// Fatal reports a compiler error and exits.
	Fatalf(pos src.XPos, msg string, args ...interface{})

	// Warnl writes compiler messages in the form expected by "errorcheck" tests
	Warnl(pos src.XPos, fmt_ string, args ...interface{})

	// Forwards the Debug flags from gc
	Debug_checknil() bool
}

type Frontend interface {
	Logger

	// CanSSA reports whether variables of type t are SSA-able.
	CanSSA(t *types.Type) bool

	// StringData returns a symbol pointing to the given string's contents.
	StringData(string) *obj.LSym

	// Auto returns a Node for an auto variable of the given type.
	// The SSA compiler uses this function to allocate space for spills.
	Auto(src.XPos, *types.Type) *ir.Name

	// Given the name for a compound type, returns the name we should use
	// for the parts of that compound type.
	SplitSlot(parent *LocalSlot, suffix string, offset int64, t *types.Type) LocalSlot

	// AllocFrame assigns frame offsets to all live auto variables.
	AllocFrame(f *Func)

	// Syslook returns a symbol of the runtime function/variable with the
	// given name.
	Syslook(string) *obj.LSym

	// UseWriteBarrier reports whether write barrier is enabled
	UseWriteBarrier() bool

	// MyImportPath provides the import name (roughly, the package) for the function being compiled.
	MyImportPath() string

	// Func returns the ir.Func of the function being compiled.
	Func() *ir.Func
}

// NewConfig returns a new configuration object for the given architecture.
func NewConfig(arch string, types Types, ctxt *obj.Link, optimize, softfloat bool) *Config

func (c *Config) Ctxt() *obj.Link
