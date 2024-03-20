// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typecheck

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/src"
)

// MakeDotArgs package all the arguments that match a ... T parameter into a []T.
func MakeDotArgs(pos src.XPos, typ *types.Type, args []ir.Node) ir.Node

// FixVariadicCall rewrites calls to variadic functions to use an
// explicit ... argument if one is not already present.
func FixVariadicCall(call *ir.CallExpr)

// FixMethodCall rewrites a method call t.M(...) into a function call T.M(t, ...).
func FixMethodCall(call *ir.CallExpr)

func AssertFixedCall(call *ir.CallExpr)

// ClosureType returns the struct type used to hold all the information
// needed in the closure for clo (clo must be a OCLOSURE node).
// The address of a variable of the returned type can be cast to a func.
func ClosureType(clo *ir.ClosureExpr) *types.Type

// MethodValueType returns the struct type used to hold all the information
// needed in the closure for a OMETHVALUE node. The address of a variable of
// the returned type can be cast to a func.
func MethodValueType(n *ir.SelectorExpr) *types.Type

// ClosureStructIter iterates through a slice of closure variables returning
// their type and offset in the closure struct.
type ClosureStructIter struct {
	closureVars []*ir.Name
	offset      int64
	next        int
}

// NewClosureStructIter creates a new ClosureStructIter for closureVars.
func NewClosureStructIter(closureVars []*ir.Name) *ClosureStructIter

// Next returns the next name, type and offset of the next closure variable.
// A nil name is returned after the last closure variable.
func (iter *ClosureStructIter) Next() (n *ir.Name, typ *types.Type, offset int64)
