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

package mips

import (
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/sys"
)

type Optab struct {
	as     obj.As
	a1     uint8
	a2     uint8
	a3     uint8
	type_  int8
	size   int8
	param  int16
	family sys.ArchFamily
	flag   uint8
}

const (
	// Optab.flag
	NOTUSETMP = 1 << iota
)

func OP(x uint32, y uint32) uint32

func SP(x uint32, y uint32) uint32

func BCOND(x uint32, y uint32) uint32

func MMU(x uint32, y uint32) uint32

func FPF(x uint32, y uint32) uint32

func FPD(x uint32, y uint32) uint32

func FPW(x uint32, y uint32) uint32

func FPV(x uint32, y uint32) uint32

func OP_RRR(op uint32, r1 uint32, r2 uint32, r3 uint32) uint32

func OP_IRR(op uint32, i uint32, r2 uint32, r3 uint32) uint32

func OP_SRR(op uint32, s uint32, r2 uint32, r3 uint32) uint32

func OP_FRRR(op uint32, r1 uint32, r2 uint32, r3 uint32) uint32

func OP_JMP(op uint32, i uint32) uint32

func OP_VI10(op uint32, df uint32, s10 int32, wd uint32, minor uint32) uint32

func OP_VMI10(s10 int32, rs uint32, wd uint32, minor uint32, df uint32) uint32
