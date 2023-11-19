// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inlheur

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// AnalyzeFunc computes function properties for fn and its contained
// closures, updating the global 'fpmap' table. It is assumed that
// "CanInline" has been run on fn and on the closures that feed
// directly into calls; other closures not directly called will also
// be checked inlinability for inlinability here in case they are
// returned as a result.
func AnalyzeFunc(fn *ir.Func, canInline func(*ir.Func), budgetForFunc func(*ir.Func) int32, inlineMaxBudget int)

// TearDown is invoked at the end of the main inlining pass; doing
// function analysis and call site scoring is unlikely to help a lot
// after this point, so nil out fpmap and other globals to reclaim
// storage.
func TearDown()

func Enabled() bool

func UnitTesting() bool

// DumpFuncProps computes and caches function properties for the func
// 'fn', writing out a description of the previously computed set of
// properties to the file given in 'dumpfile'. Used for the
// "-d=dumpinlfuncprops=..." command line flag, intended for use
// primarily in unit testing.
func DumpFuncProps(fn *ir.Func, dumpfile string)
