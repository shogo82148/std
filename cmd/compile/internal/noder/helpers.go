// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package noder

import (
	"github.com/shogo82148/std/go/constant"

	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/src"
)

type ImplicitNode interface {
	ir.Node
	SetImplicit(x bool)
}

// Implicit returns n after marking it as Implicit.
func Implicit(n ImplicitNode) ImplicitNode

func OrigConst(pos src.XPos, typ *types.Type, val constant.Value, op ir.Op, raw string) ir.Node

// FixValue returns val after converting and truncating it as
// appropriate for typ.
func FixValue(typ *types.Type, val constant.Value) constant.Value

func Nil(pos src.XPos, typ *types.Type) ir.Node

func Addr(pos src.XPos, x ir.Node) *ir.AddrExpr

func Assert(pos src.XPos, x ir.Node, typ *types.Type) ir.Node

func Binary(pos src.XPos, op ir.Op, typ *types.Type, x, y ir.Node) *ir.BinaryExpr

func Compare(pos src.XPos, typ *types.Type, op ir.Op, x, y ir.Node) *ir.BinaryExpr

func Deref(pos src.XPos, typ *types.Type, x ir.Node) *ir.StarExpr

func DotField(pos src.XPos, x ir.Node, index int) *ir.SelectorExpr

func DotMethod(pos src.XPos, x ir.Node, index int) *ir.SelectorExpr

// MethodExpr returns a OMETHEXPR node with the indicated index into the methods
// of typ. The receiver type is set from recv, which is different from typ if the
// method was accessed via embedded fields. Similarly, the X value of the
// ir.SelectorExpr is recv, the original OTYPE node before passing through the
// embedded fields.
func MethodExpr(pos src.XPos, recv ir.Node, embed *types.Type, index int) *ir.SelectorExpr

func Index(pos src.XPos, typ *types.Type, x, index ir.Node) *ir.IndexExpr

func Slice(pos src.XPos, typ *types.Type, x, low, high, max ir.Node) *ir.SliceExpr

func Unary(pos src.XPos, typ *types.Type, op ir.Op, x ir.Node) ir.Node

func IncDec(pos src.XPos, op ir.Op, x ir.Node) *ir.AssignOpStmt
