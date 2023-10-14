// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// The inlining facility makes 2 passes: first CanInline determines which
// functions are suitable for inlining, and for those that are it
// saves a copy of the body. Then InlineCalls walks each function body to
// expand calls to inlinable functions.
//
// The Debug.l flag controls the aggressiveness. Note that main() swaps level 0 and 1,
// making 1 the default and -l disable. Additional levels (beyond -l) may be buggy and
// are not supported.
//      0: disabled
//      1: 80-nodes leaf functions, oneliners, panic, lazy typechecking (default)
//      2: (unassigned)
//      3: (unassigned)
//      4: allow non-leaf functions
//
// At some point this may get another default and become switch-offable with -N.
//
// The -d typcheckinl flag enables early typechecking of all imported bodies,
// which is useful to flush out bugs.
//
// The Debug.m flag enables diagnostic output.  a single -m is useful for verifying
// which calls get inlined or not, more is for debugging, and may go away at any point.

package inline

import (
	"github.com/shogo82148/std/cmd/compile/internal/base"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/pgo"
)

// InlinePackage finds functions that can be inlined and clones them before walk expands them.
func InlinePackage(p *pgo.Profile)

// InlineDecls applies inlining to the given batch of declarations.
func InlineDecls(p *pgo.Profile, decls []ir.Node, doInline bool)

// CanInline determines whether fn is inlineable.
// If so, CanInline saves copies of fn.Body and fn.Dcl in fn.Inl.
// fn and fn.Body will already have been typechecked.
func CanInline(fn *ir.Func, profile *pgo.Profile)

// InlineImpossible returns a non-empty reason string if fn is impossible to
// inline regardless of cost or contents.
func InlineImpossible(fn *ir.Func) string

// InlineCalls/inlnode walks fn's statements and expressions and substitutes any
// calls made to inlineable functions. This is the external entry point.
func InlineCalls(fn *ir.Func, profile *pgo.Profile)

// SSADumpInline gives the SSA back end a chance to dump the function
// when producing output for debugging the compiler itself.
var SSADumpInline = func(*ir.Func) {}

// InlineCall allows the inliner implementation to be overridden.
// If it returns nil, the function will not be inlined.
var InlineCall = func(call *ir.CallExpr, fn *ir.Func, inlIndex int) *ir.InlinedCallExpr {
	base.Fatalf("inline.InlineCall not overridden")
	panic("unreachable")
}

// CalleeEffects appends any side effects from evaluating callee to init.
func CalleeEffects(init *ir.Nodes, callee ir.Node)