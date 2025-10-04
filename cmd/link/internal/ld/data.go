// Derived from Inferno utils/6l/obj.c and utils/6l/span.c
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/6l/obj.c
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/6l/span.c
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
	"github.com/shogo82148/std/cmd/internal/gcprog"
	"github.com/shogo82148/std/cmd/link/internal/loader"
	"github.com/shogo82148/std/cmd/link/internal/sym"
)

// FoldSubSymbolOffset computes the offset of symbol s to its top-level outer
// symbol. Returns the top-level symbol and the offset.
// This is used in generating external relocations.
func FoldSubSymbolOffset(ldr *loader.Loader, s loader.Sym) (loader.Sym, int64)

// ExtrelocSimple creates a simple external relocation from r, with the same
// symbol and addend.
func ExtrelocSimple(ldr *loader.Loader, r loader.Reloc) loader.ExtReloc

// ExtrelocViaOuterSym creates an external relocation from r targeting the
// outer symbol and folding the subsymbol's offset into the addend.
func ExtrelocViaOuterSym(ldr *loader.Loader, r loader.Reloc, s loader.Sym) loader.ExtReloc

func CodeblkPad(ctxt *Link, out *OutBuf, addr int64, size int64, pad []byte)

// Used only on Wasm for now.
func DatblkBytes(ctxt *Link, addr int64, size int64) []byte

type GCProg struct {
	ctxt *Link
	sym  *loader.SymbolBuilder
	w    gcprog.Writer
}

func (p *GCProg) Init(ctxt *Link, name string)

func (p *GCProg) End(size int64)

func (p *GCProg) AddSym(s loader.Sym)

// Add to the gc program the ptr bits for the type typ at
// byte offset off in the region being described.
// The type must have a pointer in it.
func (p *GCProg) AddType(off int64, typ loader.Sym)

// add a trampoline with symbol s (to be laid down after the current function)
func (ctxt *Link) AddTramp(s *loader.SymbolBuilder, typ sym.SymKind)
