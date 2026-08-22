// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sve

// OperandType classifies an SVE instruction operand.
type OperandType int

const (
	// OperandZReg is a scalable vector register (Z), e.g. <Zd>.<T>, <Zn>.<T>.
	// It has no fixed total bit width: the width is the implementation-defined
	// vector length. Only its element type and element width are known.
	OperandZReg OperandType = iota
	// OperandPReg is a scalable predicate register (P), e.g. <Pg>/M, <Pd>.<T>.
	// A predicate is modeled as a Go mask value.
	OperandPReg
	// OperandGReg is a general-purpose scalar register (W/X/R).
	OperandGReg
	// OperandVFP is a SIMD&FP scalar register (<Dd>, <V><d>, ...): a single
	// fixed-width value, such as a horizontal reduction's result (SADDV <Dd>) or
	// a DUP scalar source. Unlike a Z register it is not scalable.
	OperandVFP
	// OperandImm is an immediate.
	OperandImm
	// OperandMem is a memory operand, e.g. [<Xn|SP>{, #<imm>, MUL VL}] or a
	// gather/scatter address like [<Xn|SP>, <Zm>.D, SXTW]. simdgen does not yet
	// distinguish the memory addressing modes; they are all one "mem" class.
	OperandMem
	// OperandList is a register list, e.g. { <Zt>.B } or { <Zt1>.D-<Zt2>.D }.
	// TODO: register lists are not modeled yet; instructions carrying one are
	// skipped (see classify).
	OperandList
	// OperandSpecial is a recognized but not-yet-detailed operand: an indexed
	// register (<Zm>.<T>[<index>]), a register with an optional modifier
	// ({, <pattern>}), or a special token (<prfop>, <vl>, <pattern>, <const>,
	// <mod>, and NEON-style <Vd>/<Dd> reduction results).
	OperandSpecial
	// OperandUnknown is a token the classifier could not place at all; an anomaly.
	OperandUnknown
)

func (t OperandType) String() string

// Operand is an SVE instruction operand instantiated for a concrete element size.
type Operand struct {
	Type     OperandType
	Class    string
	BaseType string
	ElemBits int
	// Bits and Lanes are set for a fixed-width scalar register — a general-purpose
	// greg (<Xd>) or a SIMD&FP vreg (<Dd>): the total register width and lane
	// count (always 1). A scalable Z-vector leaves them 0 and is marked
	// "scalable" in the emitted def instead.
	Bits  int
	Lanes int

	// Predication is "M" (merging) or "Z" (zeroing) for governing predicates,
	// otherwise "".
	Predication string
	// AsmPos is the position in the assembly syntax (0 for the destination
	// register, 1+ for inputs). It mirrors the source template order and is the
	// field simdgen uses to order operands.
	AsmPos int
	// Raw is the source operand token, retained for deferred (mem/list/special)
	// and unknown operands so diagnostics can name what was skipped.
	Raw string

	// role is the operand's internal role: "destination", "op0"/"op1"/..., or
	// "mask" (a governing predicate). It drives out/in/inVariant partitioning at
	// emit time but is NOT emitted (simdgen orders operands by AsmPos, so a role
	// field in the YAML would be redundant).
	role string
	// arngLink is the <a> link of this operand's arrangement symbol (<T>/<Ta>/
	// <Tb>), used to resolve its per-operand element widths. Empty if the
	// operand has a fixed or no arrangement.
	arngLink string
	// fixedElem is a hardcoded element width (from e.g. ".D"), or 0.
	fixedElem int
	// fixedBits is the fixed total width of a SIMD&FP scalar named by a size
	// letter (<Dd> -> 64, <Sd> -> 32, ...), or 0 for an element-sized <V><d>.
	fixedBits int
	// isList reports that this register came from a single-register list
	// ("{ <Zt>.<T> }"). It is a distinct assembler encoding from a bare register,
	// so it is preserved (emitted as listNumber) even though the register is
	// otherwise handled like any vreg.
	isList bool
	// regName is the inner register symbol, e.g. "Zdn", "Zm", "Pg".
	regName string
}
