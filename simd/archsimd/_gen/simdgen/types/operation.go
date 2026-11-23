// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/regexp"
	"github.com/shogo82148/std/simd/archsimd/_gen/specgen/specexpr"
	"github.com/shogo82148/std/simd/archsimd/_gen/unify"
)

// RawOperation is the unifier representation of an [Operation]. It is
// translated into a more parsed form after unifier decoding.
type RawOperation struct {
	Go string

	GOARCH       string
	Asm          string
	Arrangement  *string
	OperandOrder *string

	// Sftimm, if true, adds rules that fold a shift by a constant scalar into
	// the immediate form of the machine op (Asm + "const").
	Sftimm *bool
	// AsmRules are rules on the machine op, each of the form
	// "<cond> => <result>", where <cond> is required. simdgen supplies the
	// left-hand side: the machine op applied to a pattern variable for each
	// of its operands. These are named by position, as in the default
	// lowering rule:
	//
	//	[a]      the immediate, if the operation has a variable immediate
	//	         (for a constVarImm operation, a includes the constant offset)
	//	x, y, z  the vector operands, in order (there are at most three)
	//	mask     the mask, if the machine op is masked
	//
	// <cond> and <result> refer to these variables. For example, "a==0 => x"
	// on VSSHR4S produces
	//
	//	(VSSHR4S [a] x) && a==0 => x
	//
	// Since the left-hand side comes from the operation's shape, one asm rule
	// covers every machine op the operation expands to. Write the rule for
	// the unmasked machine op. On amd64, simdgen derives the rule for the
	// AVX-512 masked form, which computes the unmasked op and then zeroes the
	// inactive lanes, by adding mask to the left-hand side and applying a
	// zeroing masked move to <result>. For example, "a==0 => x" on
	// VPSLLW128const also produces
	//
	//	(VPSLLWMasked128const [a] x mask) && a==0 => (VMOVDQU16Masked128 x mask)
	//
	// Similarly, for 512-bit vectors, it derives the rule for the merging
	// form, which blends the unmasked result into dst under the mask:
	//
	//	(VPSLLWMasked512constMerging dst [a] x mask) && a==0 => (VPBLENDMWMasked512 dst x mask)
	//
	// Format specifiers are not expanded in asm rules.
	AsmRules *[]string
	// LowerRules replace the default lowering rule of the generic op. They are
	// complete rules, usually on %g, and are emitted in order, so a rule for a
	// special case can precede a general rule.
	LowerRules *[]string
	// PeepholeRules are complete rules in addition to the default lowering
	// rule, usually on the machine op (%h).
	PeepholeRules *[]string

	// HiHalfAsm is the assembly mnemonic for the hi-half "2" variant of this operation,
	// specified in go_arm64.yaml (e.g., "VSHRN2", "VUMULL2").
	// When non-nil, simdgen generates the "2" variant machine op and folding rules.
	HiHalfAsm *string

	In              []Operand
	InVariant       []Operand
	Out             []Operand
	MemFeatures     *string
	MemFeaturesData *string
	Commutative     bool
	CPUFeature      string
	Zeroing         *bool
	Documentation   *string
	AddDoc          *string
	// ConstMask is a hack to reduce the size of defs the user writes for const-immediate
	// If present, it will be copied to [In[0].Const].
	ConstImm *string
	// NameAndSizeCheck is used to check [BWDQ] maps to (8|16|32|64) elemBits.
	NameAndSizeCheck *bool
	// If non-nil, all generation in gen_simdTypes.go and gen_intrinsics will be skipped.
	NoTypes *string
	// If non-nil, all generation in gen_simdGenericOps and gen_simdrules will be skipped.
	NoGenericOps *string
	// If non-nil, this string will be attached to the machine ssa op name.  E.g. "const"
	SSAVariant *string
	// If true, do not emit method declarations, generic ops, or intrinsics for masked variants
	// DO emit the architecture-specific opcodes and optimizations.
	HideMaskMethods *bool
	// UnpredCPUFeature is set when an SVE operation's unpredicated encoding
	// needs a higher feature level than the operation itself: the CPUFeature
	// field is the floor its predicated sibling provides, and the lowering may
	// use the unpredicated encoding only in blocks where the ssa cpufeatures
	// analysis proves this level (e.g. integer MUL: CPUFeature "SVE",
	// UnpredCPUFeature "SVE2").
	UnpredCPUFeature *string
	// WidthAgnostic marks an SVE bitwise operation whose unpredicated encoding
	// is written .D but computes the same bits under any element view. The
	// loader emits one def per element width so every Go type gets the API;
	// the unpredicated machine op collapses back to the single .D instruction
	// (see machineOpName), while the per-<T> predicated forms stay per width.
	WidthAgnostic *bool
}

// MaxVectorBits is the maximum vector length in bits Go currently supports (256
// bits / 32 bytes). It is used where a concrete upper bound is required for
// scalable SVE vectors (e.g., SSA vector types and buffer allocations).
const MaxVectorBits = 256

type Operand struct {
	Class string

	Go *string

	AsmPos int

	Base       *string
	EncodeBase *regexp.Regexp
	ElemBits   *int
	Bits       VectorSize
	EncodeBits *VectorSize

	Const *string
	// Optional immediate arg offsets. If this field is non-nil,
	// This operand will be an immediate operand:
	// The compiler will right-shift the user-passed value by ImmOffset and set it as the AuxInt
	// field of the operation.
	ImmOffset *string
	ImmMin    *int
	ImmMax    *int
	// ImmDesc is the ISA's description of an immediate, from which simdgen
	// resolves its range on SVE (see sveImmRanges).
	ImmDesc *string
	Name    *string
	Lanes   *int
	// TreatLikeAScalarOfSize means only the lower $TreatLikeAScalarOfSize bits of the vector
	// is used, so at the API level we can make it just a scalar value of this size; Then we
	// can overwrite it to a vector of the right size during intrinsics stage.
	TreatLikeAScalarOfSize *int
	// If non-nil, it means the [Class] field is overwritten here, right now this is used to
	// overwrite the results of AVX2 compares to masks.
	OverwriteClass *string
	// If non-nil, it means the [Base] field is overwritten here. This field exist solely
	// because Intel's XED data is inconsistent. e.g. VANDNP[SD] marks its operand int.
	OverwriteBase *string
	// If non-nil, it means the [ElementBits] field is overwritten. This field exist solely
	// because Intel's XED data is inconsistent. e.g. AVX512 VPMADDUBSW marks its operand
	// elemBits 16, which should be 8.
	OverwriteElementBits *int
	// For greg, specifically VPEXTR[BW], their results are specified by Intel as 32 bits,
	// but they really are 8/16 bits.
	OverwriteBits *int
	// GoBits sets the bit width of a vreg operand's Go type, which may be wider
	// than the register the instruction uses for it. For example, an AVX-512 narrowing
	// conversion from 256 bits writes a 128-bit register, but its Go result type is 256
	// bits, whose upper half the instruction zeroes. Likewise, a widening conversion to
	// 256 bits reads only a 128-bit register, the low half of its 256-bit Go input.
	//
	// [Bits] is always the register's width, while [Go] and [Lanes] describe the Go
	// type. Use [Operand.GoWidth] for the Go type's width and [Operand.NarrowReg] to
	// tell whether the two differ.
	//
	// If a vreg output has a narrower register, simdgen treats the operation's
	// opmask as "opmask-partial" (see [MaskRole]), since the instruction zeroes the
	// rest of the Go result even when merging.
	GoBits *int
	// FixedReg is the name of the fixed registers
	FixedReg *string
	// If non-nil, marks this vreg as a register list operand (for TBL/TBX).
	// Currently only list number 0 is supported (we might need to teach regalloc handle register lists
	// to support more than one register in the list).
	ListNumber *int
	// RegName is the assembly template's register symbol for this operand, e.g.
	// "Zdn", "Zn", "Pg" (SVE only). Comparing it across operands is how the
	// shape of an instruction is recognised: an input naming the same register
	// as the destination is written in place.
	RegName *string
	// PredRegName is the symbol this operand has in each of the operation's
	// predicated encodings, indexed to match InVariant (SVE only). It is nil
	// for an operation with no predicated encoding, and for every other target.
	PredRegName *[]string
	// Predication is the SVE governing-predicate qualifier, "M" (merging) or
	// "Z" (zeroing). It is set on mask operands of predicated encodings and
	// decides whether the generated machine op is the merging or the zeroing
	// form (see sveMaskSuffix).
	Predication *string
	// MaskRole describes the role of a mask operand in an amd64 instruction. It
	// is derived from XED, is set on every amd64 mask operand, and is one of:
	//
	//	"opmask"         - The operand is the AVX-512 opmask and masks the
	//	                   entire result: each result element is computed
	//	                   as if unmasked, and elements where the mask is
	//	                   false are zeroed or merged.
	//	"opmask-partial" - Like "opmask", except the instruction computes
	//	                   only the low elements of its result and zeroes
	//	                   the rest regardless of masking (so the upper
	//	                   elements are zeroed even when merging). This
	//	                   includes instructions that write a register
	//	                   narrower than the result's Go type (see
	//	                   [GoBits]).
	//	"select"         - The operand is in the opmask position, but the
	//	                   instruction uses it to select elements in some
	//	                   other way, as in blend, compress, and expand.
	//	"data"           - The operand isn't in the opmask position. For
	//	                   example, the mask result of a comparison, or
	//	                   the input of VPMOVM2D.
	//
	// It is nil on other architectures.
	MaskRole *string
	// Governing marks the SVE governing predicate among an instruction's
	// operands — the one that selects which lanes the instruction acts on, as
	// opposed to a predicate it merely reads as data (SEL's <Pv>, the <Pn>/<Pm>
	// of a predicate-logical op).
	//
	// It is set only where the instruction has no unpredicated encoding, since
	// otherwise that encoding carries the operation and its predicated sibling's
	// predicate becomes an InVariant instead. So a governing predicate here is
	// always one the Go API hides: the generated method, generic op and
	// intrinsic omit it, and the lowering synthesizes an all-true predicate in
	// its place, which is how predicated-only instructions (ZCMPGT) expose an
	// unpredicated API. See #79781.
	//
	// This is independent of [Operand.Predication]: a governing predicate need
	// not carry a qualifier (SADDV <Dd>, <Pg>, <Zn>.<T> has no lanes to merge
	// into), and a qualified predicate need not be governing in this sense (the
	// InVariant of a paired operation is a real operand a peephole supplies).
	Governing *bool
}

// VectorSize is a unifier value that is either a number or the string "scalable".
type VectorSize struct {
	Scalable bool
	NRaw     int
}

// N returns vs.NRaw, or panics if vs.Scalable.
func (vs VectorSize) N() int

func (vs VectorSize) Num() specexpr.Num

func (vs VectorSize) String() string

// IsGoverning reports whether this operand is the SVE governing predicate, and
// so is dropped from the API and filled with an all-true predicate at lowering.
func (o *Operand) IsGoverning() bool

// GoWidth returns the bit width of o's Go type. This is [Operand.Bits] unless
// [Operand.GoBits] makes the Go type wider than o's register.
func (o *Operand) GoWidth() VectorSize

// NarrowReg reports whether o's register is narrower than its Go type (see
// [Operand.GoBits]). If so, the instruction reads or writes only the low
// [Operand.Bits] bits of o's Go vector.
func (o *Operand) NarrowReg() bool

func (o Operand) OpName(s string) string

func (o Operand) OpNameAndType(s string) string

// Compare sorts operands into basic Go API order: immediates and memory,
// vectors and scalars, then masks. Ties are broken by their assembly order.
// Note that operand order may be further overridden on a case-by-case basis.
func (o Operand) Compare(p Operand) int

func (vs *VectorSize) DecodeUnified(v *unify.Value) error

func (vs *VectorSize) EncodeUnified() *unify.Value
