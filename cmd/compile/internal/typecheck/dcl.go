// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typecheck

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/src"
)

var DeclContext ir.Class = ir.PEXTERN

func DeclFunc(sym *types.Sym, recv *ir.Field, params, results []*ir.Field) *ir.Func

// Declare records that Node n declares symbol n.Sym in the specified
// declaration context.
func Declare(n *ir.Name, ctxt ir.Class)

// Export marks n for export (or reexport).
func Export(n *ir.Name)

// declare the function proper
// and declare the arguments.
// called in extern-declaration context
// returns in auto-declaration context.
func StartFuncBody(fn *ir.Func)

// finish the body.
// called in auto-declaration context.
// returns in extern-declaration context.
func FinishFuncBody()

func CheckFuncStack()

func Temp(t *types.Type) *ir.Name

// make a new Node off the books.
func TempAt(pos src.XPos, curfn *ir.Func, t *types.Type) *ir.Name

// f is method type, with receiver.
// return function type, receiver as first argument (or not).
func NewMethodType(sig *types.Type, recv *types.Type) *types.Type
