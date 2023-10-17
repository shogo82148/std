// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inlheur

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// ShouldFoldIfNameConstant analyzes expression tree 'e' to see
// whether it contains only combinations of simple references to all
// of the names in 'names' with selected constants + operators. The
// intent is to identify expression that could be folded away to a
// constant if the value of 'n' were available. Return value is TRUE
// if 'e' does look foldable given the value of 'n', and given that
// 'e' actually makes reference to 'n'. Some examples where the type
// of "n" is int64, type of "s" is string, and type of "p" is *byte:
//
//	Simple?		Expr
//	yes			n<10
//	yes			n*n-100
//	yes			(n < 10 || n > 100) && (n >= 12 || n <= 99 || n != 101)
//	yes			s == "foo"
//	yes			p == nil
//	no			n<foo()
//	no			n<1 || n>m
//	no			float32(n)<1.0
//	no			*p == 1
//	no			1 + 100
//	no			1 / n
//	no			1 + unsafe.Sizeof(n)
//
// To avoid complexities (e.g. nan, inf) we stay way from folding and
// floating point or complex operations (integers, bools, and strings
// only). We also try to be conservative about avoiding any operation
// that might result in a panic at runtime, e.g. for "n" with type
// int64:
//
//	1<<(n-9) < 100/(n<<9999)
//
// we would return FALSE due to the negative shift count and/or
// potential divide by zero.
func ShouldFoldIfNameConstant(n ir.Node, names []*ir.Name) bool
