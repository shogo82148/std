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

// Grow increases the length of s.P to lsiz.
func (s *LSym) Grow(lsiz int64)

// GrowCap increases the capacity of s.P to c.
func (s *LSym) GrowCap(c int64)

// WriteFloat32 writes f into s at offset off.
func (s *LSym) WriteFloat32(ctxt *Link, off int64, f float32)

// WriteFloat64 writes f into s at offset off.
func (s *LSym) WriteFloat64(ctxt *Link, off int64, f float64)

// WriteInt writes an integer i of size siz into s at offset off.
func (s *LSym) WriteInt(ctxt *Link, off int64, siz int, i int64)

// WriteAddr writes an address of size siz into s at offset off.
// rsym and roff specify the relocation for the address.
func (s *LSym) WriteAddr(ctxt *Link, off int64, siz int, rsym *LSym, roff int64)

// WriteWeakAddr writes an address of size siz into s at offset off.
// rsym and roff specify the relocation for the address.
// This is a weak reference.
func (s *LSym) WriteWeakAddr(ctxt *Link, off int64, siz int, rsym *LSym, roff int64)

// WriteCURelativeAddr writes a pointer-sized address into s at offset off.
// rsym and roff specify the relocation for the address which will be
// resolved by the linker to an offset from the DW_AT_low_pc attribute of
// the DWARF Compile Unit of rsym.
func (s *LSym) WriteCURelativeAddr(ctxt *Link, off int64, rsym *LSym, roff int64)

// WriteOff writes a 4 byte offset to rsym+roff into s at offset off.
// After linking the 4 bytes stored at s+off will be
// rsym+roff-(start of section that s is in).
func (s *LSym) WriteOff(ctxt *Link, off int64, rsym *LSym, roff int64)

// WriteWeakOff writes a weak 4 byte offset to rsym+roff into s at offset off.
// After linking the 4 bytes stored at s+off will be
// rsym+roff-(start of section that s is in).
func (s *LSym) WriteWeakOff(ctxt *Link, off int64, rsym *LSym, roff int64)

// WriteString writes a string of size siz into s at offset off.
func (s *LSym) WriteString(ctxt *Link, off int64, siz int, str string)

// WriteBytes writes a slice of bytes into s at offset off.
func (s *LSym) WriteBytes(ctxt *Link, off int64, b []byte) int64

func Addrel(s *LSym) *Reloc
