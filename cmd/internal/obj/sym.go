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

package obj

func Linknew(arch *LinkArch) *Link

// LookupDerived looks up or creates the symbol with name derived from symbol s.
// The resulting symbol will be static iff s is.
func (ctxt *Link) LookupDerived(s *LSym, name string) *LSym

// LookupStatic looks up the static symbol with name name.
// If it does not exist, it creates it.
func (ctxt *Link) LookupStatic(name string) *LSym

// LookupABI looks up a symbol with the given ABI.
// If it does not exist, it creates it.
func (ctxt *Link) LookupABI(name string, abi ABI) *LSym

// LookupABIInit looks up a symbol with the given ABI.
// If it does not exist, it creates it and
// passes it to init for one-time initialization.
func (ctxt *Link) LookupABIInit(name string, abi ABI, init func(s *LSym)) *LSym

// Lookup looks up the symbol with name name.
// If it does not exist, it creates it.
func (ctxt *Link) Lookup(name string) *LSym

// LookupInit looks up the symbol with name name.
// If it does not exist, it creates it and
// passes it to init for one-time initialization.
func (ctxt *Link) LookupInit(name string, init func(s *LSym)) *LSym

func (ctxt *Link) Float32Sym(f float32) *LSym

func (ctxt *Link) Float64Sym(f float64) *LSym

func (ctxt *Link) Int64Sym(i int64) *LSym

// GCLocalsSym generates a content-addressable sym containing data.
func (ctxt *Link) GCLocalsSym(data []byte) *LSym

// Assign index to symbols.
// asm is set to true if this is called by the assembler (i.e. not the compiler),
// in which case all the symbols are non-package (for now).
func (ctxt *Link) NumberSyms()

// StaticNamePref is the prefix the front end applies to static temporary
// variables. When turned into LSyms, these can be tagged as static so
// as to avoid inserting them into the linker's name lookup tables.
const StaticNamePref = ".stmp_"
