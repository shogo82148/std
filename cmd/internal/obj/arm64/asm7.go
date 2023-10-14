// cmd/7l/asm.c, cmd/7l/asmout.c, cmd/7l/optab.c, cmd/7l/span.c, cmd/ld/sub.c, cmd/ld/mod.c, from Vita Nuova.
// https://code.google.com/p/ken-cc/source/browse/
//
// 	Copyright © 1994-1999 Lucent Technologies Inc. All rights reserved.
// 	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
// 	Portions Copyright © 1997-1999 Vita Nuova Limited
// 	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
// 	Portions Copyright © 2004,2006 Bruce Ellis
// 	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
// 	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
// 	Portions Copyright © 2009 The Go Authors. All rights reserved.
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

package arm64

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

const (
	REGFROM = 1
)

type Optab struct {
	as    obj.As
	a1    uint8
	a2    uint8
	a3    uint8
	a4    uint8
	type_ int8
	size_ int8
	param int16
	flag  int8
	scond uint16
}

func IsAtomicInstruction(as obj.As) bool

const (
	S32     = 0 << 31
	S64     = 1 << 31
	Sbit    = 1 << 29
	LSL0_32 = 2 << 13
	LSL0_64 = 3 << 13
)

func OPDP2(x uint32) uint32

func OPDP3(sf uint32, op54 uint32, op31 uint32, o0 uint32) uint32

func OPBcc(x uint32) uint32

func OPBLR(x uint32) uint32

func SYSOP(l uint32, op0 uint32, op1 uint32, crn uint32, crm uint32, op2 uint32, rt uint32) uint32

func SYSHINT(x uint32) uint32

func LDSTR(sz uint32, v uint32, opc uint32) uint32

func LD2STR(o uint32) uint32

func LDSTX(sz uint32, o2 uint32, l uint32, o1 uint32, o0 uint32) uint32

func FPCMP(m uint32, s uint32, type_ uint32, op uint32, op2 uint32) uint32

func FPCCMP(m uint32, s uint32, type_ uint32, op uint32) uint32

func FPOP1S(m uint32, s uint32, type_ uint32, op uint32) uint32

func FPOP2S(m uint32, s uint32, type_ uint32, op uint32) uint32

func FPOP3S(m uint32, s uint32, type_ uint32, op uint32, op2 uint32) uint32

func FPCVTI(sf uint32, s uint32, type_ uint32, rmode uint32, op uint32) uint32

func ADR(p uint32, o uint32, rt uint32) uint32

func OPBIT(x uint32) uint32

func MOVCONST(d int64, s int, rt int) uint32

const (
	// Optab.flag
	LFROM = 1 << iota
	LFROM128
	LTO
	NOTUSETMP
	BRANCH14BITS
	BRANCH19BITS
)

// Used for padding NOOP instruction
const OP_NOOP = 0xd503201f

/* form offset parameter to SYS; special register number */
func SYSARG5(op0 int, op1 int, Cn int, Cm int, op2 int) int

func SYSARG4(op1 int, Cn int, Cm int, op2 int) int
