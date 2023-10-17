// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

import (
	"github.com/shogo82148/std/go/constant"
	"github.com/shogo82148/std/math/big"

	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/src"
)

// NewBool returns an OLITERAL representing b as an untyped boolean.
func NewBool(pos src.XPos, b bool) Node

// NewInt returns an OLITERAL representing v as an untyped integer.
func NewInt(pos src.XPos, v int64) Node

// NewString returns an OLITERAL representing s as an untyped string.
func NewString(pos src.XPos, s string) Node

// NewUintptr returns an OLITERAL representing v as a uintptr.
func NewUintptr(pos src.XPos, v int64) Node

// NewZero returns a zero value of the given type.
func NewZero(pos src.XPos, typ *types.Type) Node

// NewOne returns an OLITERAL representing 1 with the given type.
func NewOne(pos src.XPos, typ *types.Type) Node

const (
	// Maximum size in bits for big.Ints before signaling
	// overflow and also mantissa precision for big.Floats.
	ConstPrec = 512
)

func BigFloat(v constant.Value) *big.Float

// ConstOverflow reports whether constant value v is too large
// to represent with type t.
func ConstOverflow(v constant.Value, t *types.Type) bool

// IsConstNode reports whether n is a Go language constant (as opposed to a
// compile-time constant).
//
// Expressions derived from nil, like string([]byte(nil)), while they
// may be known at compile time, are not Go language constants.
func IsConstNode(n Node) bool

func IsSmallIntConst(n Node) bool
