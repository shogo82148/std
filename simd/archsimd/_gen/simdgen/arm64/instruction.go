// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arm64

import (
	"golang.org/x/arch/arm64/instgen/xmlspec"
)

// Instruction represents a parsed ARM64 instruction with domain logic
type Instruction struct {
	xmlspec.Instruction
	arrangementsCache []Arrangement
	mnemonicCache     string
	arngShape         ArngShape
}

// BaseTypeSet allows to specify the type set of values independent of arrangement's size, e.g.:
// - Float (instruction used for floating point values in lanes),
// - Uint (instruction used only for unsigned integer values in lanes with any arrangement),
// - Float|Int|Uint (e.g. VMOV V1.S[i], V0.S[j]: copy i-th lane from src vreg to j-th lane of dst vreg: basically don't care about base type).
type BaseTypeSet int

const (
	BaseTypeInt = 1 << iota
	BaseTypeUint
	BaseTypeFloat
)

func (t BaseTypeSet) String() string

// Arrangement defines the properties of a vector arrangement.
type Arrangement struct {
	arrangement string
	baseType    string
	elemBits    int
	bits        int
	lanes       int
}

// ArngShape makes certain vreg operands half or double bits wide.
type ArngShape int

const (
	DefaultArngs = ArngShape(iota)
	NarrowArngs
	LongArngs
	WideArngs
	UnsupportedArngs
)

// Mnemonic extracts the mnemonic from docvars
func (instruction *Instruction) Mnemonic() string

// Bitwise returns true if the instruction is a bitwise operation
// by detecting "Bitwise " prefix in the brief description
func (instruction *Instruction) Bitwise() bool

// InstrClass returns the instruction Class from docvars
func (instruction *Instruction) InstrClass() string

// ResultInArg0 determines if result shares register with first argument.
// This occurs when the destination register is also read as an input operand.
func (instruction *Instruction) ResultInArg0() bool

// IsAlias returns true if this instruction is an alias of another instruction
func (instruction *Instruction) IsAlias() bool

// BaseTypeSet determines if an instruction operates on integers or floats
func (instruction *Instruction) BaseTypeSet() BaseTypeSet

// Arrangements collects valid arrangement/type specifiers for the instruction
func (instruction *Instruction) Arrangements() ([]Arrangement, ArngShape)

// ArngShape returns the arrangement shape.
// Returns UnsupportedArngs for instructions we don't support yet - those will not be emitted.
// If we were not able to classify the instruction, panic to prevent wrong yaml generation.
func (instruction *Instruction) ArngShape() ArngShape

// Documentation extracts detailed instruction documentation from the XML
func (instruction *Instruction) Documentation() string

// Brief returns the brief description from XML
func (instruction *Instruction) Brief() string
