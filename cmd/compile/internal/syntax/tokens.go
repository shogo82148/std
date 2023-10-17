// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

type Token uint

const (
	_ token = iota
)

const (
	// for BranchStmt
	Break       = _Break
	Continue    = _Continue
	Fallthrough = _Fallthrough
	Goto        = _Goto

	// for CallStmt
	Go    = _Go
	Defer = _Defer
)

// Make sure we have at most 64 tokens so we can use them in a set.
const _ uint64 = 1 << (tokenCount - 1)

type LitKind uint8

// TODO(gri) With the 'i' (imaginary) suffix now permitted on integer
// and floating-point numbers, having a single ImagLit does
// not represent the literal kind well anymore. Remove it?
const (
	IntLit LitKind = iota
	FloatLit
	ImagLit
	RuneLit
	StringLit
)

type Operator uint

const (
	_ Operator = iota

	// Def is the : in :=
	Def
	Not
	Recv
	Tilde

	// precOrOr
	OrOr

	// precAndAnd
	AndAnd

	// precCmp
	Eql
	Neq
	Lss
	Leq
	Gtr
	Geq

	// precAdd
	Add
	Sub
	Or
	Xor

	// precMul
	Mul
	Div
	Rem
	And
	AndNot
	Shl
	Shr
)

// Operator precedences
const (
	_ = iota
)
