// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package inlheur

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

func AnalyzeFunc(fn *ir.Func, canInline func(*ir.Func)) *FuncProps

func UnitTesting() bool

// DumpFuncProps computes and caches function properties for the func
// 'fn' and any closures it contains, or if fn is nil, it writes out the
// cached set of properties to the file given in 'dumpfile'. Used for
// the "-d=dumpinlfuncprops=..." command line flag, intended for use
// primarily in unit testing.
func DumpFuncProps(fn *ir.Func, dumpfile string, canInline func(*ir.Func))
