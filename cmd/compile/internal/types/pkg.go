// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

type Pkg struct {
	Path    string
	Name    string
	Prefix  string
	Syms    map[string]*Sym
	Pathsym *obj.LSym

	Direct bool
}

// NewPkg returns a new Pkg for the given package path and name.
// Unless name is the empty string, if the package exists already,
// the existing package name and the provided name must match.
func NewPkg(path, name string) *Pkg

func PkgMap() map[string]*Pkg

func (pkg *Pkg) Lookup(name string) *Sym

// LookupOK looks up name in pkg and reports whether it previously existed.
func (pkg *Pkg) LookupOK(name string) (s *Sym, existed bool)

func (pkg *Pkg) LookupBytes(name []byte) *Sym

// LookupNum looks up the symbol starting with prefix and ending with
// the decimal n. If prefix is too long, LookupNum panics.
func (pkg *Pkg) LookupNum(prefix string, n int) *Sym

// Selector looks up a selector identifier.
func (pkg *Pkg) Selector(name string) *Sym

func InternString(b []byte) string

// CleanroomDo invokes f in an environment with no preexisting packages.
// For testing of import/export only.
func CleanroomDo(f func())
