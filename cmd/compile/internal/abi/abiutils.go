// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
)

// ABIParamResultInfo stores the results of processing a given
// function type to compute stack layout and register assignments. For
// each input and output parameter we capture whether the param was
// register-assigned (and to which register(s)) or the stack offset
// for the param if is not going to be passed in registers according
// to the rules in the Go internal ABI specification (1.17).
type ABIParamResultInfo struct {
	inparams          []ABIParamAssignment
	outparams         []ABIParamAssignment
	offsetToSpillArea int64
	spillAreaSize     int64
	inRegistersUsed   int
	outRegistersUsed  int
	config            *ABIConfig
}

func (a *ABIParamResultInfo) Config() *ABIConfig

func (a *ABIParamResultInfo) InParams() []ABIParamAssignment

func (a *ABIParamResultInfo) OutParams() []ABIParamAssignment

func (a *ABIParamResultInfo) InRegistersUsed() int

func (a *ABIParamResultInfo) OutRegistersUsed() int

func (a *ABIParamResultInfo) InParam(i int) *ABIParamAssignment

func (a *ABIParamResultInfo) OutParam(i int) *ABIParamAssignment

func (a *ABIParamResultInfo) SpillAreaOffset() int64

func (a *ABIParamResultInfo) SpillAreaSize() int64

// ArgWidth returns the amount of stack needed for all the inputs
// and outputs of a function or method, including ABI-defined parameter
// slots and ABI-defined spill slots for register-resident parameters.
// The name is inherited from (*Type).ArgWidth(), which it replaces.
func (a *ABIParamResultInfo) ArgWidth() int64

// RegIndex stores the index into the set of machine registers used by
// the ABI on a specific architecture for parameter passing.  RegIndex
// values 0 through N-1 (where N is the number of integer registers
// used for param passing according to the ABI rules) describe integer
// registers; values N through M (where M is the number of floating
// point registers used).  Thus if the ABI says there are 5 integer
// registers and 7 floating point registers, then RegIndex value of 4
// indicates the 5th integer register, and a RegIndex value of 11
// indicates the 7th floating point register.
type RegIndex uint8

// ABIParamAssignment holds information about how a specific param or
// result will be passed: in registers (in which case 'Registers' is
// populated) or on the stack (in which case 'Offset' is set to a
// non-negative stack offset). The values in 'Registers' are indices
// (as described above), not architected registers.
type ABIParamAssignment struct {
	Type      *types.Type
	Name      *ir.Name
	Registers []RegIndex
	offset    int32
}

// Offset returns the stack offset for addressing the parameter that "a" describes.
// This will panic if "a" describes a register-allocated parameter.
func (a *ABIParamAssignment) Offset() int32

// RegisterTypes returns a slice of the types of the registers
// corresponding to a slice of parameters.  The returned slice
// has capacity for one more, likely a memory type.
func RegisterTypes(apa []ABIParamAssignment) []*types.Type

func (pa *ABIParamAssignment) RegisterTypesAndOffsets() ([]*types.Type, []int64)

// FrameOffset returns the frame-pointer-relative location that a function
// would spill its input or output parameter to, if such a spill slot exists.
// If there is none defined (e.g., register-allocated outputs) it panics.
// For register-allocated inputs that is their spill offset reserved for morestack;
// for stack-allocated inputs and outputs, that is their location on the stack.
// (In a future version of the ABI, register-resident inputs may lose their defined
// spill area to help reduce stack sizes.)
func (a *ABIParamAssignment) FrameOffset(i *ABIParamResultInfo) int64

// RegAmounts holds a specified number of integer/float registers.
type RegAmounts struct {
	intRegs   int
	floatRegs int
}

// ABIConfig captures the number of registers made available
// by the ABI rules for parameter passing and result returning.
type ABIConfig struct {
	// Do we need anything more than this?
	offsetForLocals int64
	regAmounts      RegAmounts
	which           obj.ABI
}

// NewABIConfig returns a new ABI configuration for an architecture with
// iRegsCount integer/pointer registers and fRegsCount floating point registers.
func NewABIConfig(iRegsCount, fRegsCount int, offsetForLocals int64, which uint8) *ABIConfig

// Copy returns config.
//
// TODO(mdempsky): Remove.
func (config *ABIConfig) Copy() *ABIConfig

// Which returns the ABI number
func (config *ABIConfig) Which() obj.ABI

// LocalsOffset returns the architecture-dependent offset from SP for args and results.
// In theory this is only used for debugging; it ought to already be incorporated into
// results from the ABI-related methods
func (config *ABIConfig) LocalsOffset() int64

// FloatIndexFor translates r into an index in the floating point parameter
// registers.  If the result is negative, the input index was actually for the
// integer parameter registers.
func (config *ABIConfig) FloatIndexFor(r RegIndex) int64

// NumParamRegs returns the total number of registers used to
// represent a parameter of the given type, which must be register
// assignable.
func (config *ABIConfig) NumParamRegs(typ *types.Type) int

// ABIAnalyzeTypes takes slices of parameter and result types, and returns an ABIParamResultInfo,
// based on the given configuration.  This is the same result computed by config.ABIAnalyze applied to the
// corresponding method/function type, except that all the embedded parameter names are nil.
// This is intended for use by ssagen/ssa.go:(*state).rtcall, for runtime functions that lack a parsed function type.
func (config *ABIConfig) ABIAnalyzeTypes(params, results []*types.Type) *ABIParamResultInfo

// ABIAnalyzeFuncType takes a function type 'ft' and an ABI rules description
// 'config' and analyzes the function to determine how its parameters
// and results will be passed (in registers or on the stack), returning
// an ABIParamResultInfo object that holds the results of the analysis.
func (config *ABIConfig) ABIAnalyzeFuncType(ft *types.Type) *ABIParamResultInfo

// ABIAnalyze returns the same result as ABIAnalyzeFuncType, but also
// updates the offsets of all the receiver, input, and output fields.
// If setNname is true, it also sets the FrameOffset of the Nname for
// the field(s); this is for use when compiling a function and figuring out
// spill locations.  Doing this for callers can cause races for register
// outputs because their frame location transitions from BOGUS_FUNARG_OFFSET
// to zero to an as-if-AUTO offset that has no use for callers.
func (config *ABIConfig) ABIAnalyze(t *types.Type, setNname bool) *ABIParamResultInfo

// ToString method renders an ABIParamAssignment in human-readable
// form, suitable for debugging or unit testing.
func (ri *ABIParamAssignment) ToString(config *ABIConfig, extra bool) string

// String method renders an ABIParamResultInfo in human-readable
// form, suitable for debugging or unit testing.
func (ri *ABIParamResultInfo) String() string

// ComputePadding returns a list of "post element" padding values in
// the case where we have a structure being passed in registers. Given
// a param assignment corresponding to a struct, it returns a list
// containing padding values for each field, e.g. the Kth element in
// the list is the amount of padding between field K and the following
// field. For things that are not structs (or structs without padding)
// it returns a list of zeros. Example:
//
//	type small struct {
//		x uint16
//		y uint8
//		z int32
//		w int32
//	}
//
// For this struct we would return a list [0, 1, 0, 0], meaning that
// we have one byte of padding after the second field, and no bytes of
// padding after any of the other fields. Input parameter "storage" is
// a slice with enough capacity to accommodate padding elements for
// the architected register set in question.
func (pa *ABIParamAssignment) ComputePadding(storage []uint64) []uint64
