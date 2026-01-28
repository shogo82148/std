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

package sym

import "github.com/shogo82148/std/cmd/internal/objabi"

// A SymKind describes the kind of memory represented by a symbol.
type SymKind uint8

// Defined SymKind values.
//
// TODO(rsc): Give idiomatic Go names.
//
//go:generate stringer -type=SymKind
const (
	// An otherwise invalid zero value for the type.
	Sxxx SymKind = iota
	// The text segment, containing executable instructions.
	STEXT
	STEXTFIPSSTART
	STEXTFIPS
	STEXTFIPSEND
	STEXTEND
	SELFRXSECT
	SMACHOPLT

	// Read-only, non-executable, unrelocated segment.
	SSTRING
	SGOSTRING
	SGCBITS
	SRODATA
	SRODATAFIPSSTART
	SRODATAFIPS
	SRODATAFIPSEND
	SRODATAEND
	SPCLNTAB
	STYPELINK
	SELFROSECT

	// Read-only, non-executable, dynamically relocatable segment.
	//
	// This segment holds read-only data that contains pointers to
	// other parts of the program. When generating a position
	// independent executable or a shared library, these sections
	// are "relro", meaning that they start as writable, and are
	// changed to be read-only after dynamic relocations are applied.
	//
	// When no dynamic relocations are required, as when generating
	// an executable that is not position independent, this is just
	// part of the normal read-only segment.
	SRODATARELRO
	STYPE
	SGOFUNC
	SELFRELROSECT
	SMACHORELROSECT

	SITABLINK

	// Allocated writable segment.
	SFirstWritable
	SBUILDINFO
	SFIPSINFO
	SELFSECT
	SMACHO
	SWINDOWS
	SMODULEDATA
	SELFGOT
	SMACHOGOT
	SNOPTRDATA
	SNOPTRDATAFIPSSTART
	SNOPTRDATAFIPS
	SNOPTRDATAFIPSEND
	SNOPTRDATAEND
	SINITARR
	SDATA
	SDATAFIPSSTART
	SDATAFIPS
	SDATAFIPSEND
	SDATAEND
	SXCOFFTOC

	// Allocated zero-initialized segment.
	SBSS
	SNOPTRBSS
	SLIBFUZZER_8BIT_COUNTER
	SCOVERAGE_COUNTER
	SCOVERAGE_AUXVAR
	STLSBSS

	// Unallocated segment.
	SFirstUnallocated
	SXREF
	SMACHOSYMSTR
	SMACHOSYMTAB
	SMACHOINDIRECTPLT
	SMACHOINDIRECTGOT
	SDYNIMPORT
	SHOSTOBJ
	SUNDEFEXT

	// Unallocated DWARF debugging segment.
	SDWARFSECT
	// DWARF symbol types created by compiler or linker.
	SDWARFCUINFO
	SDWARFCONST
	SDWARFFCN
	SDWARFABSFCN
	SDWARFTYPE
	SDWARFVAR
	SDWARFRANGE
	SDWARFLOC
	SDWARFLINES
	SDWARFADDR

	// SEH symbol types. These are probably allocated at run time.
	SSEHUNWINDINFO
	SSEHSECT
)

// AbiSymKindToSymKind maps values read from object files (which are
// of type cmd/internal/objabi.SymKind) to values of type SymKind.
var AbiSymKindToSymKind = [...]SymKind{
	objabi.Sxxx:                    Sxxx,
	objabi.STEXT:                   STEXT,
	objabi.STEXTFIPS:               STEXTFIPS,
	objabi.SRODATA:                 SRODATA,
	objabi.SRODATAFIPS:             SRODATAFIPS,
	objabi.SNOPTRDATA:              SNOPTRDATA,
	objabi.SNOPTRDATAFIPS:          SNOPTRDATAFIPS,
	objabi.SDATA:                   SDATA,
	objabi.SDATAFIPS:               SDATAFIPS,
	objabi.SBSS:                    SBSS,
	objabi.SNOPTRBSS:               SNOPTRBSS,
	objabi.STLSBSS:                 STLSBSS,
	objabi.SDWARFCUINFO:            SDWARFCUINFO,
	objabi.SDWARFCONST:             SDWARFCONST,
	objabi.SDWARFFCN:               SDWARFFCN,
	objabi.SDWARFABSFCN:            SDWARFABSFCN,
	objabi.SDWARFTYPE:              SDWARFTYPE,
	objabi.SDWARFVAR:               SDWARFVAR,
	objabi.SDWARFRANGE:             SDWARFRANGE,
	objabi.SDWARFLOC:               SDWARFLOC,
	objabi.SDWARFLINES:             SDWARFLINES,
	objabi.SDWARFADDR:              SDWARFADDR,
	objabi.SLIBFUZZER_8BIT_COUNTER: SLIBFUZZER_8BIT_COUNTER,
	objabi.SCOVERAGE_COUNTER:       SCOVERAGE_COUNTER,
	objabi.SCOVERAGE_AUXVAR:        SCOVERAGE_AUXVAR,
	objabi.SSEHUNWINDINFO:          SSEHUNWINDINFO,
}

// ReadOnly are the symbol kinds that form read-only sections
// that never require runtime relocations.
var ReadOnly = []SymKind{
	SSTRING,
	SGOSTRING,
	SGCBITS,
	SRODATA,
	SRODATAFIPSSTART,
	SRODATAFIPS,
	SRODATAFIPSEND,
	SRODATAEND,
}

// IsText returns true if t is a text type.
func (t SymKind) IsText() bool

// IsData returns true if t is any kind of data type.
func (t SymKind) IsData() bool

// IsDATA returns true if t is one of the SDATA types.
func (t SymKind) IsDATA() bool

// IsRODATA returns true if t is one of the SRODATA types.
func (t SymKind) IsRODATA() bool

// IsNOPTRDATA returns true if t is one of the SNOPTRDATA types.
func (t SymKind) IsNOPTRDATA() bool

func (t SymKind) IsDWARF() bool
