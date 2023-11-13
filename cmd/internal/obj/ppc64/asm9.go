// cmd/9l/optab.c, cmd/9l/asmout.c from Vita Nuova.
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2008 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2008 Lucent Technologies Inc. and others
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

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

const (
	// R bit option in prefixed load/store/add D-form operations
	PFX_R_ABS   = 0
	PFX_R_PCREL = 1
)

const (
	// The preferred hardware nop instruction.
	NOP = 0x60000000
)

type Optab struct {
	as    obj.As
	a1    uint8
	a2    uint8
	a3    uint8
	a4    uint8
	a5    uint8
	a6    uint8
	type_ int8
	size  int8

	// A prefixed instruction is generated by this opcode. This cannot be placed
	// across a 64B PC address. Opcodes should not translate to more than one
	// prefixed instruction. The prefixed instruction should be written first
	// (e.g when Optab.size > 8).
	ispfx bool

	asmout func(*ctxt9, *obj.Prog, *Optab, *[5]uint32)
}

// These are opcodes above which may generate different sequences depending on whether prefix opcode support
// is available
type PrefixableOptab struct {
	Optab
	minGOPPC64 int
	pfxsize    int8
}

// Determine if the build configuration requires a TOC pointer.
// It is assumed this always called after buildop.
func NeedTOCpointer(ctxt *obj.Link) bool

func OPVXX1(o uint32, xo uint32, oe uint32) uint32

func OPVXX2(o uint32, xo uint32, oe uint32) uint32

func OPVXX2VA(o uint32, xo uint32, oe uint32) uint32

func OPVXX3(o uint32, xo uint32, oe uint32) uint32

func OPVXX4(o uint32, xo uint32, oe uint32) uint32

func OPDQ(o uint32, xo uint32, oe uint32) uint32

func OPVX(o uint32, xo uint32, oe uint32, rc uint32) uint32

func OPVC(o uint32, xo uint32, oe uint32, rc uint32) uint32

func OPVCC(o uint32, xo uint32, oe uint32, rc uint32) uint32

func OPCC(o uint32, xo uint32, rc uint32) uint32

/* Generate MD-form opcode */
func OPMD(o, xo, rc uint32) uint32

/* the order is dest, a/s, b/imm for both arithmetic and logical operations. */
func AOP_RRR(op uint32, d uint32, a uint32, b uint32) uint32

/* VX-form 2-register operands, r/none/r */
func AOP_RR(op uint32, d uint32, a uint32) uint32

/* VA-form 4-register operands */
func AOP_RRRR(op uint32, d uint32, a uint32, b uint32, c uint32) uint32

func AOP_IRR(op uint32, d uint32, a uint32, simm uint32) uint32

/* VX-form 2-register + UIM operands */
func AOP_VIRR(op uint32, d uint32, a uint32, simm uint32) uint32

/* VX-form 2-register + ST + SIX operands */
func AOP_IIRR(op uint32, d uint32, a uint32, sbit uint32, simm uint32) uint32

/* VA-form 3-register + SHB operands */
func AOP_IRRR(op uint32, d uint32, a uint32, b uint32, simm uint32) uint32

/* VX-form 1-register + SIM operands */
func AOP_IR(op uint32, d uint32, simm uint32) uint32

/* XX1-form 3-register operands, 1 VSR operand */
func AOP_XX1(op uint32, r uint32, a uint32, b uint32) uint32

/* XX2-form 3-register operands, 2 VSR operands */
func AOP_XX2(op uint32, xt uint32, a uint32, xb uint32) uint32

/* XX3-form 3 VSR operands */
func AOP_XX3(op uint32, xt uint32, xa uint32, xb uint32) uint32

/* XX3-form 3 VSR operands + immediate */
func AOP_XX3I(op uint32, xt uint32, xa uint32, xb uint32, c uint32) uint32

/* XX4-form, 4 VSR operands */
func AOP_XX4(op uint32, xt uint32, xa uint32, xb uint32, xc uint32) uint32

/* DQ-form, VSR register, register + offset operands */
func AOP_DQ(op uint32, xt uint32, a uint32, b uint32) uint32

/* Z23-form, 3-register operands + CY field */
func AOP_Z23I(op uint32, d uint32, a uint32, b uint32, c uint32) uint32

/* X-form, 3-register operands + EH field */
func AOP_RRRI(op uint32, d uint32, a uint32, b uint32, c uint32) uint32

func LOP_RRR(op uint32, a uint32, s uint32, b uint32) uint32

func LOP_IRR(op uint32, a uint32, s uint32, uimm uint32) uint32

func OP_BR(op uint32, li uint32, aa uint32) uint32

func OP_BC(op uint32, bo uint32, bi uint32, bd uint32, aa uint32) uint32

func OP_BCR(op uint32, bo uint32, bi uint32) uint32

func OP_RLW(op uint32, a uint32, s uint32, sh uint32, mb uint32, me uint32) uint32

func AOP_EXTSWSLI(op uint32, a uint32, s uint32, sh uint32) uint32

func AOP_ISEL(op uint32, t uint32, a uint32, b uint32, bc uint32) uint32

/* MD-form 2-register, 2 6-bit immediate operands */
func AOP_MD(op uint32, a uint32, s uint32, sh uint32, m uint32) uint32

/* MDS-form 3-register, 1 6-bit immediate operands. rsh argument is a register. */
func AOP_MDS(op, to, from, rsh, m uint32) uint32

func AOP_PFX_00_8LS(r, ie uint32) uint32

func AOP_PFX_10_MLS(r, ie uint32) uint32

const (
	/* each rhs is OPVCC(_, _, _, _) */
	OP_ADD      = 31<<26 | 266<<1 | 0<<10 | 0
	OP_ADDI     = 14<<26 | 0<<1 | 0<<10 | 0
	OP_ADDIS    = 15<<26 | 0<<1 | 0<<10 | 0
	OP_ANDI     = 28<<26 | 0<<1 | 0<<10 | 0
	OP_EXTSB    = 31<<26 | 954<<1 | 0<<10 | 0
	OP_EXTSH    = 31<<26 | 922<<1 | 0<<10 | 0
	OP_EXTSW    = 31<<26 | 986<<1 | 0<<10 | 0
	OP_ISEL     = 31<<26 | 15<<1 | 0<<10 | 0
	OP_MCRF     = 19<<26 | 0<<1 | 0<<10 | 0
	OP_MCRFS    = 63<<26 | 64<<1 | 0<<10 | 0
	OP_MCRXR    = 31<<26 | 512<<1 | 0<<10 | 0
	OP_MFCR     = 31<<26 | 19<<1 | 0<<10 | 0
	OP_MFFS     = 63<<26 | 583<<1 | 0<<10 | 0
	OP_MFSPR    = 31<<26 | 339<<1 | 0<<10 | 0
	OP_MFSR     = 31<<26 | 595<<1 | 0<<10 | 0
	OP_MFSRIN   = 31<<26 | 659<<1 | 0<<10 | 0
	OP_MTCRF    = 31<<26 | 144<<1 | 0<<10 | 0
	OP_MTFSF    = 63<<26 | 711<<1 | 0<<10 | 0
	OP_MTFSFI   = 63<<26 | 134<<1 | 0<<10 | 0
	OP_MTSPR    = 31<<26 | 467<<1 | 0<<10 | 0
	OP_MTSR     = 31<<26 | 210<<1 | 0<<10 | 0
	OP_MTSRIN   = 31<<26 | 242<<1 | 0<<10 | 0
	OP_MULLW    = 31<<26 | 235<<1 | 0<<10 | 0
	OP_MULLD    = 31<<26 | 233<<1 | 0<<10 | 0
	OP_OR       = 31<<26 | 444<<1 | 0<<10 | 0
	OP_ORI      = 24<<26 | 0<<1 | 0<<10 | 0
	OP_ORIS     = 25<<26 | 0<<1 | 0<<10 | 0
	OP_RLWINM   = 21<<26 | 0<<1 | 0<<10 | 0
	OP_RLWNM    = 23<<26 | 0<<1 | 0<<10 | 0
	OP_SUBF     = 31<<26 | 40<<1 | 0<<10 | 0
	OP_RLDIC    = 30<<26 | 4<<1 | 0<<10 | 0
	OP_RLDICR   = 30<<26 | 2<<1 | 0<<10 | 0
	OP_RLDICL   = 30<<26 | 0<<1 | 0<<10 | 0
	OP_RLDCL    = 30<<26 | 8<<1 | 0<<10 | 0
	OP_EXTSWSLI = 31<<26 | 445<<2
	OP_SETB     = 31<<26 | 128<<1
)

const (
	D_FORM = iota
	DS_FORM
)
