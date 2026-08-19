// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/encoding/binary"

	"github.com/shogo82148/std/cmd/compile/internal/ssa/ssaop"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/obj/s390x"
	"github.com/shogo82148/std/cmd/internal/src"
)

// Aux is an interface to hold miscellaneous data in Blocks and Values.
type Aux interface {
	CanBeAnSSAAux()
}

func BoolToAuxInt(b bool) int64

// FlagConstant represents the result of a compile-time comparison.
// The sense of these flags does not necessarily represent the hardware's notion
// of a flags register - these are just a compile-time construct.
// We happen to match the semantics to those of arm/arm64.
// Note that these semantics differ from x86: the carry flag has the opposite
// sense on a subtraction!
//
//	On amd64, C=1 represents a borrow, e.g. SBB on amd64 does x - y - C.
//	On arm64, C=0 represents a borrow, e.g. SBC on arm64 does x - y - ^C.
//	 (because it does x + ^y + C).
//
// See https://en.wikipedia.org/wiki/Carry_flag#Vs._borrow_flag
type FlagConstant uint8

func IsNewObjectCall(aux Aux) bool

func IsSpecializedMalloc(aux Aux) bool

// StringAux wraps string values for use in Aux.
type StringAux string

func StringToAux(s string) Aux

func AuxIntToArm64ConditionalParams(i int64) Arm64ConditionalParams

func LogLargeCopy(funcName string, pos src.XPos, s int64)

var AuxMark auxMark

// PanicBoundsC contains a constant for a bounds failure.
type PanicBoundsC struct {
	C int64
}

// PanicBoundsCC contains 2 constants for a bounds failure.
type PanicBoundsCC struct {
	Cx int64
	Cy int64
}

func GetPPC64Shiftsh(auxint int64) int64

func GetPPC64Shiftmb(auxint int64) int64

// DecodePPC64RotateMask is the inverse operation of encodePPC64RotateMask.  The values returned as
// mb and me satisfy the POWER ISA definition of MASK(x,y) where MASK(mb,me) = mask.
func DecodePPC64RotateMask(sauxint int64) (rotate, mb, me int64, mask uint64)

// DivisionNeedsFixUp reports whether the division needs fix-up code.
func DivisionNeedsFixUp(v *Value) bool

// ZeroUpper32Bits checks if value zeroes out upper 32-bit of 64-bit register.
// depth limits recursion depth. In AMD64.rules 3 is used as limit,
// because it catches same amount of cases as 4.
func ZeroUpper32Bits(x *Value) bool

// ZeroUpper48Bits is similar to ZeroUpper32Bits, but for upper 48 bits.
func ZeroUpper48Bits(x *Value) bool

// ZeroUpper56Bits is similar to ZeroUpper32Bits, but for upper 56 bits.
func ZeroUpper56Bits(x *Value) bool

// IsSamePtr reports whether p1 and p2 point to the same address.
func IsSamePtr(p1, p2 *Value) bool

// Disjoint reports whether the memory region specified by [p1:p1+t1.Size())
// does not overlap with [p2:p2+t2.Size()).
// A return value of false does not imply the regions overlap.
func Disjoint(p1 *Value, t1 *types.Type, p2 *Value, t2 *types.Type) bool

// Disjoint1 reports whether the memory region specified by [p1:p1+n1)
// does not overlap with [p2:p2+n2).
// A return value of false does not imply the regions overlap.
func Disjoint1(p1 *Value, n1 int64, p2 *Value, n2 int64) bool

// DisjointTypes reports whether a memory region pointed to by a pointer of type
// t1 does not overlap with a memory region pointed to by a pointer of type t2 --
// based on type aliasing rules.
func DisjointTypes(t1 *types.Type, t2 *types.Type) bool

// Overlap reports whether the ranges given by the given offset and
// size pairs Overlap.
func Overlap(offset1, size1, offset2, size2 int64) bool

func IsInlinableMemmove(dst, src *Value, sz int64, c *Config) bool

func (StringAux) CanBeAnSSAAux()

// returns the Lsb part of the auxInt field of arm64 bitfield ops.
func (bfc Arm64BitField) Lsb() int64

// returns the Width part of the auxInt field of arm64 bitfield ops.
func (bfc Arm64BitField) Width() int64

// extracts NZCV flags from auxint.
func (condParams Arm64ConditionalParams) Nzcv() int64

// extracts constant value from auxint if present.
func (condParams Arm64ConditionalParams) ConstValue() (int64, bool)

// N reports whether the result of an operation is negative (high bit set).
func (fc FlagConstant) N() bool

// Z reports whether the result of an operation is 0.
func (fc FlagConstant) Z() bool

// C reports whether an unsigned add overflowed (carry), or an
// unsigned subtract did not underflow (borrow).
func (fc FlagConstant) C() bool

// V reports whether a signed operation overflowed or underflowed.
func (fc FlagConstant) V() bool

func (fc FlagConstant) Eq() bool

func (fc FlagConstant) Ne() bool

func (fc FlagConstant) Lt() bool

func (fc FlagConstant) Le() bool

func (fc FlagConstant) Gt() bool

func (fc FlagConstant) Ge() bool

func (fc FlagConstant) Ult() bool

func (fc FlagConstant) Ule() bool

func (fc FlagConstant) Ugt() bool

func (fc FlagConstant) Uge() bool

func (fc FlagConstant) LtNoov() bool

func (fc FlagConstant) LeNoov() bool

func (fc FlagConstant) GtNoov() bool

func (fc FlagConstant) GeNoov() bool

func (fc FlagConstant) String() string

func (p PanicBoundsC) CanBeAnSSAAux()

func (p PanicBoundsCC) CanBeAnSSAAux()

// AddFlags32 returns the flags that would be set from computing x+y.
func AddFlags32(x, y int32) FlagConstant

// AddFlags64 returns the flags that would be set from computing x+y.
func AddFlags64(x, y int64) FlagConstant

func Arm64BitFieldToAuxInt(v Arm64BitField) int64

func Arm64ConditionalParamsToAuxInt(v Arm64ConditionalParams) int64

// encodes the lsb and width for arm(64) bitfield ops into the expected auxInt format.
func ArmBFAuxInt(lsb, width int64) Arm64BitField

func AuxIntToArm64BitField(i int64) Arm64BitField

func AuxIntToBool(i int64) bool

func AuxIntToFlagConstant(x int64) FlagConstant

func AuxIntToFloat32(i int64) float32

func AuxIntToFloat64(i int64) float64

func AuxIntToInt16(i int64) int16

func AuxIntToInt32(i int64) int32

func AuxIntToInt64(i int64) int64

func AuxIntToInt8(i int64) int8

func AuxIntToOp(cc int64) ssaop.Op

func AuxIntToUint64(i int64) uint64

func AuxIntToUint8(i int64) uint8

func AuxIntToValAndOff(i int64) ValAndOff

func AuxToCall(i Aux) *AuxCall

func AuxToPanicBoundsC(i Aux) PanicBoundsC

func AuxToPanicBoundsCC(i Aux) PanicBoundsCC

func AuxToS390xCCMask(i Aux) s390x.CCMask

func AuxToS390xRotateParams(i Aux) s390x.RotateParams

func AuxToString(i Aux) string

func AuxToSym(i Aux) Sym

func AuxToType(i Aux) *types.Type

// B2i translates a boolean value to 0 or 1 for assigning to auxInt.
func B2i(b bool) int64

// B2i32 translates a boolean value to 0 or 1.
func B2i32(b bool) int32

func CallToAux(s *AuxCall) Aux

// CanMergeLoad reports whether the load can be merged into target without
// invalidating the schedule.
func CanMergeLoad(target, load *Value) bool

// CanMergeLoadClobber reports whether the load can be merged into target without
// invalidating the schedule.
// It also checks that the other non-load argument x is something we
// are ok with clobbering.
func CanMergeLoadClobber(target, load, x *Value) bool

func CanMergeSym(x, y Sym) bool

func CanMulStrengthReduce(config *Config, x int64) bool

func CanMulStrengthReduce32(config *Config, x int32) bool

// CanonLessThan returns whether x is "ordered" less than y, for purposes of normalizing
// generated code as much as possible.
func CanonLessThan(x, y *Value) bool

// Clobber invalidates values. Returns true.
// Clobber is used by rewrite rules to:
//
//	A) make sure the values are really dead and never used again.
//	B) decrement use counts of the values' args.
func Clobber(vv ...*Value) bool

// ClobberIfDead resets v when use count is 1. Returns true.
// ClobberIfDead is used by rewrite rules to decrement
// use counts of v's args when v is dead and never used.
func ClobberIfDead(v *Value) bool

// CountRule increments Func.ruleMatches[key].
// If Func.ruleMatches is non-nil at the end
// of compilation, it will be printed to stdout.
// This is intended to make it easier to find which functions
// which contain lots of rules matches when developing new rules.
func CountRule(v *Value, key string) bool

type DeadValueChoice bool

// Compress mask and shift into single value of the form
// me | mb<<8 | rotate<<16 | nbits<<24 where me and mb can
// be used to regenerate the input mask.
func EncodePPC64RotateMask(rotate, mask, nbits int64) int64

// for a pseudo-op like (LessThan x), extract x.
func FlagArg(v *Value) *Value

type FlagConstantBuilder struct {
	N bool
	Z bool
	C bool
	V bool
}

func FlagConstantToAuxInt(x FlagConstant) int64

func Float32ToAuxInt(f float32) int64

func Float64ToAuxInt(f float64) int64

// When v is (IMake typ (StructMake ...)), convert to
// (IMake typ arg) where arg is the pointer-y argument to
// the StructMake (there must be exactly one).
func ImakeOfStructMake(v *Value) *Value

func Int16ToAuxInt(i int16) int64

func Int32ToAuxInt(i int32) int64

func Int64ToAuxInt(i int64) int64

func Int8ToAuxInt(i int8) int64

// Is12Bit reports whether n can be represented as a signed 12 bit integer.
func Is12Bit(n int64) bool

// Is16Bit reports whether n can be represented as a signed 16 bit integer.
func Is16Bit(n int64) bool

func Is16BitInt(t *types.Type) bool

// Is20Bit reports whether n can be represented as a signed 20 bit integer.
func Is20Bit(n int64) bool

// Is32Bit reports whether n can be represented as a signed 32 bit integer.
func Is32Bit(n int64) bool

func Is32BitFloat(t *types.Type) bool

func Is32BitInt(t *types.Type) bool

func Is64BitFloat(t *types.Type) bool

func Is64BitInt(t *types.Type) bool

func Is8BitInt(t *types.Type) bool

// isPowerOfTwoX functions report whether n is a power of 2.
func IsPowerOfTwo[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64](n T) bool

// This verifies that the mask is a set of
// consecutive bits including the least
// significant bit.
func IsPPC64ValidShiftMask(v int64) bool

// Test if this value can encoded as a mask for a rlwinm like
// operation.  Masks can also extend from the msb and wrap to
// the lsb too.  That is, the valid masks are 32 bit strings
// of the form: 0..01..10..0 or 1..10..01..1 or 1...1
//
// Note: This ignores the upper 32 bits of the input. When a
// zero extended result is desired (e.g a 64 bit result), the
// user must verify the upper 32 bits are 0 and the mask is
// contiguous (that is, non-wrapping).
func IsPPC64WordRotateMask(v64 int64) bool

func IsPtr(t *types.Type) bool

// IsSameCall reports whether aux is the same as the given named symbol.
func IsSameCall(aux Aux, name string) bool

// IsU32Bit reports whether n can be represented as an unsigned 32 bit integer.
func IsU32Bit(n int64) bool

// IsVolatile reports whether v is a pointer to argument region on stack which
// will be clobbered by a function call.
func IsVolatile(v *Value) bool

const (
	LeaveDeadValues  DeadValueChoice = false
	RemoveDeadValues                 = true

	RepZeroThreshold = 1408
	RepMoveThreshold = 1408
)

func Log16(n int16) int64

func Log16u(n uint16) int64

func Log32(n int32) int64

func Log32u(n uint32) int64

func Log64(n int64) int64

func Log64u(n uint64) int64

// logXu returns the logarithm of n base 2.
// n must be a power of 2 (isPowerOfTwo returns true)
func Log8u(n uint8) int64

// LogicFlags32 returns flags set to the sign/zeroness of x.
// C and V are set to false.
func LogicFlags32(x int32) FlagConstant

// LogicFlags64 returns flags set to the sign/zeroness of x.
// C and V are set to false.
func LogicFlags64(x int64) FlagConstant

// LogLargeCopyValue logs the occurrence of a large copy.
// The best place to do this is in the rewrite rules where the size of the move is easy to find.
// "Large" is arbitrarily chosen to be 128 bytes; this may change.
func LogLargeCopyValue(v *Value, s int64) bool

// LogRule logs the use of the rule s. This will only be enabled if
// rewrite rules were generated with the -log option, see _gen/rulegen.go.
func LogRule(s string)

func MakeJumpTableSym(b *Block) *obj.LSym

// Combine (ANDconst [m] (SRWconst [s])) into (RLWINM [y]) or return 0
func MergePPC64AndSrwi(m, s int64) int64

// Test if a RLWINM feeding into a CLRLSLDI can be merged into RLWINM.  Return
// the encoded RLWINM constant, or 0 if they cannot be merged.
func MergePPC64ClrlsldiRlwinm(sld int32, rlw int64) int64

// Test if a word shift right feeding into a CLRLSLDI can be merged into RLWINM.
// Return the encoded RLWINM constant, or 0 if they cannot be merged.
func MergePPC64ClrlsldiSrw(sld, srw int64) int64

// Decompose a shift right into an equivalent rotate/mask,
// and return mask & m.
func MergePPC64RShiftMask(m, s, nbits int64) int64

// Compute the encoded RLWINM constant from combining (SLDconst [sld] (SRWconst [srw] x)),
// or return 0 if they cannot be combined.
func MergePPC64SldiSrw(sld, srw int64) int64

// MergeSym merges two symbolic offsets. There is no real merging of
// offsets, we just pick the non-nil one.
func MergeSym(x, y Sym) Sym

// MoveSize returns the number of bytes an aligned MOV instruction moves.
func MoveSize(align int64, c *Config) int64

// MulStrengthReduce returns v*x evaluated at the location
// (block and source position) of m.
// canMulStrengthReduce must have returned true.
func MulStrengthReduce(m *Value, v *Value, x int64) *Value

// MulStrengthReduce32 returns v*x evaluated at the location
// (block and source position) of m.
// canMulStrengthReduce32 must have returned true.
// The upper 32 bits of m might be set to junk.
func MulStrengthReduce32(m *Value, v *Value, x int32) *Value

func NewPPC64ShiftAuxInt(sh, mb, me, sz int64) int32

// NoteRule is an easy way to track if a rule is matched when writing
// new ones.  Make the rule of interest also conditional on
//
//	NoteRule("note to self: rule of interest matched")
//
// and that message will print when the rule matches.
func NoteRule(s string) bool

// ntzX returns the number of trailing zeros.
func Ntz64(x int64) int

// OneBit reports whether x contains exactly one set bit.
func OneBit[T int8 | int16 | int32 | int64](x T) bool

func OpToAuxInt(o ssaop.Op) int64

func PanicBoundsCCToAux(p PanicBoundsCC) Aux

func PanicBoundsCToAux(p PanicBoundsC) Aux

// Read16 reads two bytes from the read-only global sym at offset off.
func Read16(sym Sym, off int64, byteorder binary.ByteOrder) uint16

// Read32 reads four bytes from the read-only global sym at offset off.
func Read32(sym Sym, off int64, byteorder binary.ByteOrder) uint32

// Read64 reads eight bytes from the read-only global sym at offset off.
func Read64(sym Sym, off int64, byteorder binary.ByteOrder) uint64

// Read8 reads one byte from the read-only global sym at offset off.
func Read8(sym Sym, off int64) uint8

func RewriteStructStore(v *Value) *Value

func S390xCCMaskToAux(c s390x.CCMask) Aux

func S390xRotateParamsToAux(r s390x.RotateParams) Aux

// SetPos sets the position of v to pos, then returns true.
// Useful for setting the result of a rewrite's position to
// something other than the default.
func SetPos(v *Value, pos src.XPos) bool

// ShiftIsBounded reports whether (left/right) shift Value v is known to be bounded.
// A shift is bounded if it is shifting by less than the width of the shifted value.
func ShiftIsBounded(v *Value) bool

// SubFlags32 returns the flags that would be set from computing x-y.
func SubFlags32(x, y int32) FlagConstant

// SubFlags64 returns the flags that would be set from computing x-y.
func SubFlags64(x, y int64) FlagConstant

func SupportsPPC64PCRel() bool

// SymIsRO reports whether sym is a read-only global.
func SymIsRO(sym Sym) bool

func SymToAux(s Sym) Aux

func TypeToAux(t *types.Type) Aux

func Uint64ToAuxInt(i uint64) int64

func Uint8ToAuxInt(i uint8) int64

func ValAndOffToAuxInt(v ValAndOff) int64

func (fcs FlagConstantBuilder) Encode() FlagConstant

func IsConstZero(v *Value) bool
