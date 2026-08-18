// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/cmd/internal/src"
)

func FprintFunc(p funcPrinter, f *Func)

func PrintFunc(f *Func)

func StmtString(p src.XPos) string

type StringFuncPrinter struct {
	w         io.Writer
	printDead bool
}

func HashFunc(f *Func) []byte

func (f *Func) String() string

// RewriteHash returns a hash of f suitable for detecting rewrite cycles.
func (f *Func) RewriteHash() string

func (p StringFuncPrinter) Header(f *Func)

func (p StringFuncPrinter) StartBlock(b *Block, reachable bool)

func (p StringFuncPrinter) EndBlock(b *Block, reachable bool)

func (p StringFuncPrinter) Value(v *Value, live bool)

func (p StringFuncPrinter) StartDepCycle()

func (p StringFuncPrinter) EndDepCycle()

func (p StringFuncPrinter) Named(n LocalSlot, vals []*Value)
