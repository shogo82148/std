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

func NewBool(pos src.XPos, b bool) Node

func NewInt(pos src.XPos, v int64) Node

func NewString(pos src.XPos, s string) Node

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
