// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package loopvar applies the proper variable capture, according
// to experiment, flags, language version, etc.
package loopvar

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/internal/src"
)

type VarAndLoop struct {
	Name    *ir.Name
	Loop    ir.Node
	LastPos src.XPos
}

// ForCapture transforms for and range loops that declare variables that might be
// captured by a closure or escaped to the heap, using a syntactic check that
// conservatively overestimates the loops where capture occurs, but still avoids
// transforming the (large) majority of loops. It returns the list of names
// subject to this change, that may (once transformed) be heap allocated in the
// process. (This allows checking after escape analysis to call out any such
// variables, in case it causes allocation/performance problems).
//
// The decision to transform loops is normally encoded in the For/Range loop node
// field DistinctVars but is also dependent on base.LoopVarHash, and some values
// of base.Debug.LoopVar (which is set per-package).  Decisions encoded in DistinctVars
// are preserved across inlining, so if package a calls b.F and loops in b.F are
// transformed, then they are always transformed, whether b.F is inlined or not.
//
// Per-package, the debug flag settings that affect this transformer:
//
// base.LoopVarHash != nil => use hash setting to govern transformation.
// note that LoopVarHash != nil sets base.Debug.LoopVar to 1 (unless it is >= 11, for testing/debugging).
//
// base.Debug.LoopVar == 11 => transform ALL loops ignoring syntactic/potential escape. Do not log, can be in addition to GOEXPERIMENT.
//
// The effect of GOEXPERIMENT=loopvar is to change the default value (0) of base.Debug.LoopVar to 1 for all packages.
func ForCapture(fn *ir.Func) []VarAndLoop

func LogTransformations(transformed []VarAndLoop)
