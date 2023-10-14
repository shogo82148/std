// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

// Sym represents an object name in a segmented (pkg, name) namespace.
// Most commonly, this is a Go identifier naming an object declared within a package,
// but Syms are also used to name internal synthesized objects.
//
// As an exception, field and method names that are exported use the Sym
// associated with localpkg instead of the package that declared them. This
// allows using Sym pointer equality to test for Go identifier uniqueness when
// handling selector expressions.
//
// Ideally, Sym should be used for representing Go language constructs,
// while cmd/internal/obj.LSym is used for representing emitted artifacts.
//
// NOTE: In practice, things can be messier than the description above
// for various reasons (historical, convenience).
type Sym struct {
	Linkname string

	Pkg  *Pkg
	Name string

	// The unique ONAME, OTYPE, OPACK, or OLITERAL node that this symbol is
	// bound to within the current scope. (Most parts of the compiler should
	// prefer passing the Node directly, rather than relying on this field.)
	//
	// Deprecated: New code should avoid depending on Sym.Def. Add
	// mdempsky@ as a reviewer for any CLs involving Sym.Def.
	Def Object

	flags bitset8
}

func (sym *Sym) OnExportList() bool
func (sym *Sym) Uniq() bool
func (sym *Sym) Siggen() bool
func (sym *Sym) Asm() bool
func (sym *Sym) Func() bool

func (sym *Sym) SetOnExportList(b bool)
func (sym *Sym) SetUniq(b bool)
func (sym *Sym) SetSiggen(b bool)
func (sym *Sym) SetAsm(b bool)
func (sym *Sym) SetFunc(b bool)

func (sym *Sym) IsBlank() bool

// Deprecated: This method should not be used directly. Instead, use a
// higher-level abstraction that directly returns the linker symbol
// for a named object. For example, reflectdata.TypeLinksym(t) instead
// of reflectdata.TypeSym(t).Linksym().
func (sym *Sym) Linksym() *obj.LSym

// Deprecated: This method should not be used directly. Instead, use a
// higher-level abstraction that directly returns the linker symbol
// for a named object. For example, (*ir.Name).LinksymABI(abi) instead
// of (*ir.Name).Sym().LinksymABI(abi).
func (sym *Sym) LinksymABI(abi obj.ABI) *obj.LSym

// Less reports whether symbol a is ordered before symbol b.
//
// Symbols are ordered exported before non-exported, then by name, and
// finally (for non-exported symbols) by package path.
func (a *Sym) Less(b *Sym) bool

// IsExported reports whether name is an exported Go symbol (that is,
// whether it begins with an upper-case letter).
func IsExported(name string) bool
