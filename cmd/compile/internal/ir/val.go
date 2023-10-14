// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

import (
	"github.com/shogo82148/std/go/constant"

	"github.com/shogo82148/std/cmd/compile/internal/types"
)

func ConstType(n Node) constant.Kind

// IntVal returns v converted to int64.
// Note: if t is uint64, very large values will be converted to negative int64.
func IntVal(t *types.Type, v constant.Value) int64

func AssertValidTypeForConst(t *types.Type, v constant.Value)

func ValidTypeForConst(t *types.Type, v constant.Value) bool

var OKForConst [types.NTYPE]bool

// Int64Val returns n as an int64.
// n must be an integer or rune constant.
func Int64Val(n Node) int64

// Uint64Val returns n as a uint64.
// n must be an integer or rune constant.
func Uint64Val(n Node) uint64

// BoolVal returns n as a bool.
// n must be a boolean constant.
func BoolVal(n Node) bool

// StringVal returns the value of a literal string Node as a string.
// n must be a string constant.
func StringVal(n Node) string
