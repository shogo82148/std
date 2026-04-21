// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

const (
	InvalidIn     inShape = iota
	PureVregIn
	OneKmaskIn
	OneImmIn
	OneKmaskImmIn
	PureKmaskIn
)

const (
	InvalidOut     outShape = iota
	NoOut
	OneVregOut
	OneGregOut
	OneKmaskOut
	OneVregOutAtIn
)

const (
	InvalidMask maskShape = iota
	NoMask
	OneMask
	AllMasks
)

const (
	InvalidImm  immShape = iota
	NoImm
	ConstImm
	VarImm
	ConstVarImm
)

const (
	InvalidMem memShape = iota
	NoMem
	VregMemIn
)

// SSAType returns the string for the type reference in SSA generation,
// for example in the intrinsics generating template.
func (op Operation) SSAType() string

// GoType returns the Go type returned by this operation (relative to the simd package),
// for example "int32" or "Int8x16".  This is used in a template.
func (op Operation) GoType() string

// ImmName returns the name to use for an operation's immediate operand.
// This can be overriden in the yaml with "name" on an operand,
// otherwise, for now, "constant"
func (op Operation) ImmName() string

func (o Operand) OpName(s string) string

func (o Operand) OpNameAndType(s string) string

// GoExported returns [Go] with first character capitalized.
func (op Operation) GoExported() string

// DocumentationExported returns [Documentation] with method name capitalized.
func (op Operation) DocumentationExported() string

// Op0Name returns the name to use for the 0 operand,
// if any is present, otherwise the parameter is used.
func (op Operation) Op0Name(s string) string

// Op1Name returns the name to use for the 1 operand,
// if any is present, otherwise the parameter is used.
func (op Operation) Op1Name(s string) string

// Op2Name returns the name to use for the 2 operand,
// if any is present, otherwise the parameter is used.
func (op Operation) Op2Name(s string) string

// Op3Name returns the name to use for the 3 operand,
// if any is present, otherwise the parameter is used.
func (op Operation) Op3Name(s string) string

// Op0NameAndType returns the name and type to use for
// the 0 operand, if a name is provided, otherwise
// the parameter value is used as the default.
func (op Operation) Op0NameAndType(s string) string

// Op1NameAndType returns the name and type to use for
// the 1 operand, if a name is provided, otherwise
// the parameter value is used as the default.
func (op Operation) Op1NameAndType(s string) string

// Op2NameAndType returns the name and type to use for
// the 2 operand, if a name is provided, otherwise
// the parameter value is used as the default.
func (op Operation) Op2NameAndType(s string) string

// Op3NameAndType returns the name and type to use for
// the 3 operand, if a name is provided, otherwise
// the parameter value is used as the default.
func (op Operation) Op3NameAndType(s string) string

// Op4NameAndType returns the name and type to use for
// the 4 operand, if a name is provided, otherwise
// the parameter value is used as the default.
func (op Operation) Op4NameAndType(s string) string

func (op Operation) GenericName() string

func (o Operation) String() string

func (op Operand) String() string
