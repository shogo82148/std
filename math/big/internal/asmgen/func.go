// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asmgen

// A Func represents a single assembly function.
type Func struct {
	Name    string
	Asm     *Asm
	inputs  []string
	outputs []string
	args    map[string]int
}

// Func starts a new function in the assembly output.
func (a *Asm) Func(decl string) *Func

// Arg allocates a new register, copies the named argument (or result) into it,
// and returns that register.
func (f *Func) Arg(name string) Reg

// ArgHint is like Arg but uses a register allocation hint.
func (f *Func) ArgHint(name string, hint Hint) Reg

// ArgPtr is like Arg but returns a RegPtr.
func (f *Func) ArgPtr(name string) RegPtr

// StoreArg stores src into the named argument (or result).
func (f *Func) StoreArg(src Reg, name string)
