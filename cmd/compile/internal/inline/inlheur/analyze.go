// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inlheur

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

func AnalyzeFunc(fn *ir.Func, canInline func(*ir.Func), inlineMaxBudget int32) *FuncProps

// RevisitInlinability revisits the question of whether to continue to
// treat function 'fn' as an inline candidate based on the set of
// properties we've computed for it. If (for example) it has an
// initial size score of 150 and no interesting properties to speak
// of, then there isn't really any point to moving ahead with it as an
// inline candidate.
func RevisitInlinability(fn *ir.Func, budgetForFunc func(*ir.Func) int32)

func UnitTesting() bool

// DumpFuncProps computes and caches function properties for the func
// 'fn' and any closures it contains, or if fn is nil, it writes out the
// cached set of properties to the file given in 'dumpfile'. Used for
// the "-d=dumpinlfuncprops=..." command line flag, intended for use
// primarily in unit testing.
func DumpFuncProps(fn *ir.Func, dumpfile string, canInline func(*ir.Func), inlineMaxBudget int32)
