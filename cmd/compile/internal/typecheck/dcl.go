// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typecheck

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/src"
)

// DeclFunc declares the parameters for fn and adds it to
// Target.Funcs.
//
// Before returning, it sets CurFunc to fn. When the caller is done
// constructing fn, it must call FinishFuncBody to restore CurFunc.
func DeclFunc(fn *ir.Func)

// FinishFuncBody restores ir.CurFunc to its state before the last
// call to DeclFunc.
func FinishFuncBody()

func CheckFuncStack()

// make a new Node off the books.
func TempAt(pos src.XPos, curfn *ir.Func, typ *types.Type) *ir.Name

// f is method type, with receiver.
// return function type, receiver as first argument (or not).
func NewMethodType(sig *types.Type, recv *types.Type) *types.Type
