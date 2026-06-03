// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arm64

// OperandType defines the type of an operand for ARM64 instruction generation.
type OperandType int

const (
	OperandVReg  OperandType = iota
	OperandGReg
	OperandImm
	OperandVElem
	OperandList
)

func (t OperandType) String() string

// Operand represents an arm64 instruction operand instantiated for concrete arrangement.
type Operand struct {
	Type     OperandType
	Class    string
	BaseType string
	ElemBits int
	Bits     int
	Lanes    int
	ImmMax   int
	// The operand's role. Possible values:
	//   - "destination":      the output register
	//   - "original":         the original SSA value of "destination" (for resultInArg0 instructions)
	//   - ends with "_i":     vector element index: should get ImmMax = lanes-1
	//   - "op0", "op1", ...:  input registers
	//   - other strings:      immediate names (e.g. "immshift", "amount", "immzero")
	Role       string
	ListNumber int
	AsmPos     int
}
