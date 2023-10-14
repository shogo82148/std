// Derived from Inferno utils/6l/l.h and related files.
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/6l/l.h
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package ld

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/cmd/link/internal/loader"
	"github.com/shogo82148/std/cmd/link/internal/sym"
	"github.com/shogo82148/std/debug/elf"
)

type Shlib struct {
	Path string
	Hash []byte
	Deps []string
	File *elf.File
}

// Link holds the context for writing object code from a compiler
// or for reading that input into the linker.
type Link struct {
	Target
	ErrorReporter
	ArchSyms

	outSem chan int
	Out    *OutBuf

	version int

	Debugvlog int
	Bso       *bufio.Writer

	Loaded bool

	compressDWARF bool

	Libdir       []string
	Library      []*sym.Library
	LibraryByPkg map[string]*sym.Library
	Shlibs       []Shlib
	Textp        []loader.Sym
	Moduledata   loader.Sym

	PackageFile  map[string]string
	PackageShlib map[string]string

	tramps []loader.Sym

	compUnits []*sym.CompilationUnit
	runtimeCU *sym.CompilationUnit

	loader  *loader.Loader
	cgodata []cgodata

	datap  []loader.Sym
	dynexp []loader.Sym

	// Elf symtab variables.
	numelfsym int

	// These are symbols that created and written by the linker.
	// Rather than creating a symbol, and writing all its data into the heap,
	// you can create a symbol, and just a generation function will be called
	// after the symbol's been created in the output mmap.
	generatorSyms map[loader.Sym]generatorFunc
}

func (ctxt *Link) Logf(format string, args ...interface{})

// Allocate a new version (i.e. symbol namespace).
func (ctxt *Link) IncVersion() int

// returns the maximum version number
func (ctxt *Link) MaxVersion() int
