// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sve loads ARM64 SVE / SVE2 instruction definitions from the ARM A64
// ISA XML files and emits them as simdgen unify values.
// TODO: merge with the arm64 package, the approach taken here should take over
// the NEON loader.
// TODO: merge with x/arch/arm64/instgen?
//
// SVE registers are "scalable": their total bit width is the hardware
// implementation-defined vector length rather than a fixed 128/256/512 bits. So
// emitted vector operands carry only a base type and an element width, without a
// fixed bits/lanes count.
//
// Arrangement is per-operand. An SVE instruction template such as
//
//	ADD  <Zdn>.<T>, <Pg>/M, <Zdn>.<T>, <Zm>.<T>
//
// stands for a family of concrete instructions, one per value of the <T>
// arrangement symbol. simdgen enumerates them by resolving each operand's
// arrangement symbol from the section's explanations. Different symbols can be
// encoded in the same instruction field but interpreted differently, the
// loader also takes care of this.
//
// It emits register, mask, immediate, memory and special operands.
// Memory and special operands are opaque at this moment.
// Register-list operands are not modeled yet, except for single-register lists,
// so instructions carrying one are skipped (TODO); see classify.
//
// TODO: Peepholes might need the structure of memory operands, implement it?
// TODO: special operands are like registers with indexing, prefetch ops, etc.
// They seem too specialized that we might want to manually implment them instead
// of via simdgen, but we can revisit this.
package sve

import (
	"golang.org/x/arch/arm64/instgen/xmlspec"
)

// Instruction is a *logical* SVE instruction, one per iclass.
type Instruction struct {
	xmlspec.Instruction
	// iclass is the specific class this logical instruction represents.
	// A raw xmlspec.Instruction can hold several iclasses with distinct mnemonics.
	// If nil, the first iclass is used.
	iclass        *xmlspec.Iclass
	mnemonicCache string
}
