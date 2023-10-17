// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typecheck

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/src"
)

// DotField returns a field selector expression that selects the
// index'th field of the given expression, which must be of struct or
// pointer-to-struct type.
func DotField(pos src.XPos, x ir.Node, index int) *ir.SelectorExpr

// XDotMethod returns an expression representing the field selection
// x.sym. If any implicit field selection are necessary, those are
// inserted too.
func XDotField(pos src.XPos, x ir.Node, sym *types.Sym) *ir.SelectorExpr

// XDotMethod returns an expression representing the method value
// x.sym (i.e., x is a value, not a type). If any implicit field
// selection are necessary, those are inserted too.
//
// If callee is true, the result is an ODOTMETH/ODOTINTER, otherwise
// an OMETHVALUE.
func XDotMethod(pos src.XPos, x ir.Node, sym *types.Sym, callee bool) *ir.SelectorExpr
