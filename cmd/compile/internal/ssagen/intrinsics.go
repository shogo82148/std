// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssagen

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

func IsIntrinsicCall(n *ir.CallExpr) bool

func IsIntrinsicSym(sym *types.Sym) bool

// GenIntrinsicBody generates the function body for a bodyless intrinsic.
// This is used when the intrinsic is used in a non-call context, e.g.
// as a function pointer, or (for a method) being referenced from the type
// descriptor.
//
// The compiler already recognizes a call to fn as an intrinsic and can
// directly generate code for it. So we just fill in the body with a call
// to fn.
func GenIntrinsicBody(fn *ir.Func)
