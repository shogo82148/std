// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package devirtualize implements two "devirtualization" optimization passes:
//
//   - "Static" devirtualization which replaces interface method calls with
//     direct concrete-type method calls where possible.
//   - "Profile-guided" devirtualization which replaces indirect calls with a
//     conditional direct call to the hottest concrete callee from a profile, as
//     well as a fallback using the original indirect call.
package devirtualize

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// StaticCall devirtualizes the given call if possible when the concrete callee
// is available statically.
func StaticCall(s *State, call *ir.CallExpr)

// State holds precomputed state for use in [StaticCall].
type State struct {
	// ifaceAssignments maps interface variables to all their assignments
	// defined inside functions stored in the analyzedFuncs set.
	// Note: it does not include direct assignments to nil.
	ifaceAssignments map[*ir.Name][]assignment

	// ifaceCallExprAssigns stores every [*ir.CallExpr], which has an interface
	// result, that is assigned to a variable.
	ifaceCallExprAssigns map[*ir.CallExpr][]ifaceAssignRef

	// analyzedFuncs is a set of Funcs that were analyzed for iface assignments.
	analyzedFuncs map[*ir.Func]struct{}
}

// InlinedCall updates the [State] to take into account a newly inlined call.
func (s *State) InlinedCall(fun *ir.Func, origCall *ir.CallExpr, inlinedCall *ir.InlinedCallExpr)
