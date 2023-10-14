// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sym

import "github.com/shogo82148/std/cmd/internal/dwarf"

// LoaderSym holds a loader.Sym value. We can't refer to this
// type from the sym package since loader imports sym.
type LoaderSym uint32

// A CompilationUnit represents a set of source files that are compiled
// together. Since all Go sources in a Go package are compiled together,
// there's one CompilationUnit per package that represents all Go sources in
// that package, plus one for each assembly file.
//
// Equivalently, there's one CompilationUnit per object file in each Library
// loaded by the linker.
//
// These are used for both DWARF and pclntab generation.
type CompilationUnit struct {
	Lib       *Library
	PclnIndex int
	PCs       []dwarf.Range
	DWInfo    *dwarf.DWDie
	FileTable []string

	Consts    LoaderSym
	FuncDIEs  []LoaderSym
	VarDIEs   []LoaderSym
	AbsFnDIEs []LoaderSym
	RangeSyms []LoaderSym
	Textp     []LoaderSym
}
