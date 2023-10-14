// Inferno utils/5l/asm.c
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/5l/asm.c
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

package ppc64

const (
	// For genstub, the type of stub required by the caller.
	STUB_TOC = iota
	STUB_PCREL
)

const (
	OP_TOCRESTORE    = 0xe8410018
	OP_TOCSAVE       = 0xf8410018
	OP_NOP           = 0x60000000
	OP_BL            = 0x48000001
	OP_BCTR          = 0x4e800420
	OP_BCTRL         = 0x4e800421
	OP_BCL           = 0x40000001
	OP_ADDI          = 0x38000000
	OP_ADDIS         = 0x3c000000
	OP_LD            = 0xe8000000
	OP_PLA_PFX       = 0x06100000
	OP_PLA_SFX       = 0x38000000
	OP_PLD_PFX_PCREL = 0x04100000
	OP_PLD_SFX       = 0xe4000000
	OP_MFLR          = 0x7c0802a6
	OP_MTLR          = 0x7c0803a6
	OP_MFCTR         = 0x7c0902a6
	OP_MTCTR         = 0x7c0903a6

	OP_ADDIS_R12_R2  = OP_ADDIS | 12<<21 | 2<<16
	OP_ADDIS_R12_R12 = OP_ADDIS | 12<<21 | 12<<16
	OP_ADDI_R12_R12  = OP_ADDI | 12<<21 | 12<<16
	OP_PLD_SFX_R12   = OP_PLD_SFX | 12<<21
	OP_PLA_SFX_R12   = OP_PLA_SFX | 12<<21
	OP_LIS_R12       = OP_ADDIS | 12<<21
	OP_LD_R12_R12    = OP_LD | 12<<21 | 12<<16
	OP_MTCTR_R12     = OP_MTCTR | 12<<21
	OP_MFLR_R12      = OP_MFLR | 12<<21
	OP_MFLR_R0       = OP_MFLR | 0<<21
	OP_MTLR_R0       = OP_MTLR | 0<<21

	// This is a special, preferred form of bcl to obtain the next
	// instruction address (NIA, aka PC+4) in LR.
	OP_BCL_NIA = OP_BCL | 20<<21 | 31<<16 | 1<<2

	// Masks to match opcodes
	MASK_PLD_PFX  = 0xfff70000
	MASK_PLD_SFX  = 0xfc1f0000
	MASK_PLD_RT   = 0x03e00000
	MASK_OP_LD    = 0xfc000003
	MASK_OP_ADDIS = 0xfc000000
)
