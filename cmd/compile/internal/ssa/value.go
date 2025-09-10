// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/src"
)

// A Value represents a value in the SSA representation of the program.
// The ID and Type fields must not be modified. The remainder may be modified
// if they preserve the value of the Value (e.g. changing a (mul 2 x) to an (add x x)).
type Value struct {
	// A unique identifier for the value. For performance we allocate these IDs
	// densely starting at 1.  There is no guarantee that there won't be occasional holes, though.
	ID ID

	// The operation that computes this value. See op.go.
	Op Op

	// The type of this value. Normally this will be a Go type, but there
	// are a few other pseudo-types, see ../types/type.go.
	Type *types.Type

	// Auxiliary info for this value. The type of this information depends on the opcode and type.
	// AuxInt is used for integer values, Aux is used for other values.
	// Floats are stored in AuxInt using math.Float64bits(f).
	// Unused portions of AuxInt are filled by sign-extending the used portion,
	// even if the represented value is unsigned.
	// Users of AuxInt which interpret AuxInt as unsigned (e.g. shifts) must be careful.
	// Use Value.AuxUnsigned to get the zero-extended value of AuxInt.
	AuxInt int64
	Aux    Aux

	// Arguments of this value
	Args []*Value

	// Containing basic block
	Block *Block

	// Source position
	Pos src.XPos

	// Use count. Each appearance in Value.Args and Block.Controls counts once.
	Uses int32

	// wasm: Value stays on the WebAssembly stack. This value will not get a "register" (WebAssembly variable)
	// nor a slot on Go stack, and the generation of this value is delayed to its use time.
	OnWasmStack bool

	// Is this value in the per-function constant cache? If so, remove from cache before changing it or recycling it.
	InCache bool

	// Storage for the first three args
	argstorage [3]*Value
}

// short form print. Just v#.
func (v *Value) String() string

func (v *Value) AuxInt8() int8

func (v *Value) AuxUInt8() uint8

func (v *Value) AuxInt16() int16

func (v *Value) AuxInt32() int32

// AuxUnsigned returns v.AuxInt as an unsigned value for OpConst*.
// v.AuxInt is always sign-extended to 64 bits, even if the
// represented value is unsigned. This undoes that sign extension.
func (v *Value) AuxUnsigned() uint64

func (v *Value) AuxFloat() float64

func (v *Value) AuxValAndOff() ValAndOff

func (v *Value) AuxArm64BitField() arm64BitField

func (v *Value) AuxArm64ConditionalParams() arm64ConditionalParams

// long form print.  v# = opcode <type> [aux] args [: reg] (names)
func (v *Value) LongString() string

// If/when midstack inlining is enabled (-l=4), the compiler gets both larger and slower.
// Not-inlining this method is a help (*Value.reset and *Block.NewValue0 are similar).
//
//go:noinline
func (v *Value) AddArg(w *Value)

//go:noinline
func (v *Value) AddArg2(w1, w2 *Value)

//go:noinline
func (v *Value) AddArg3(w1, w2, w3 *Value)

//go:noinline
func (v *Value) AddArg4(w1, w2, w3, w4 *Value)

//go:noinline
func (v *Value) AddArg5(w1, w2, w3, w4, w5 *Value)

//go:noinline
func (v *Value) AddArg6(w1, w2, w3, w4, w5, w6 *Value)

func (v *Value) AddArgs(a ...*Value)

func (v *Value) SetArg(i int, w *Value)

func (v *Value) SetArgs1(a *Value)

func (v *Value) SetArgs2(a, b *Value)

func (v *Value) SetArgs3(a, b, c *Value)

func (v *Value) SetArgs4(a, b, c, d *Value)

func (v *Value) Logf(msg string, args ...interface{})
func (v *Value) Log() bool
func (v *Value) Fatalf(msg string, args ...interface{})

// ResultReg returns the result register assigned to v, in cmd/internal/obj/$ARCH numbering.
// It is similar to Reg and Reg0, except that it is usable interchangeably for all Value Ops.
// If you know v.Op, using Reg or Reg0 (as appropriate) will be more efficient.
func (v *Value) ResultReg() int16

// Reg returns the register assigned to v, in cmd/internal/obj/$ARCH numbering.
func (v *Value) Reg() int16

// Reg0 returns the register assigned to the first output of v, in cmd/internal/obj/$ARCH numbering.
func (v *Value) Reg0() int16

// Reg1 returns the register assigned to the second output of v, in cmd/internal/obj/$ARCH numbering.
func (v *Value) Reg1() int16

// RegTmp returns the temporary register assigned to v, in cmd/internal/obj/$ARCH numbering.
func (v *Value) RegTmp() int16

func (v *Value) RegName() string

// MemoryArg returns the memory argument for the Value.
// The returned value, if non-nil, will be memory-typed (or a tuple with a memory-typed second part).
// Otherwise, nil is returned.
func (v *Value) MemoryArg() *Value

// LackingPos indicates whether v is a value that is unlikely to have a correct
// position assigned to it.  Ignoring such values leads to more user-friendly positions
// assigned to nearby values and the blocks containing them.
func (v *Value) LackingPos() bool

// AutoVar returns a *Name and int64 representing the auto variable and offset within it
// where v should be spilled.
func AutoVar(v *Value) (*ir.Name, int64)

// CanSSA reports whether values of type t can be represented as a Value.
func CanSSA(t *types.Type) bool
