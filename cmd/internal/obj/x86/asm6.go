// Inferno utils/6l/span.c
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

package x86

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

type Optab struct {
	as     obj.As
	ytab   []ytab
	prefix uint8
	op     opBytes
}

const (
	Yxxx = iota
	Ynone
	Yi0
	Yi1
	Yu2
	Yi8
	Yu8
	Yu7
	Ys32
	Yi32
	Yi64
	Yiauto
	Yal
	Ycl
	Yax
	Ycx
	Yrb
	Yrl
	Yrl32
	Yrf
	Yf0
	Yrx
	Ymb
	Yml
	Ym
	Ybr
	Ycs
	Yss
	Yds
	Yes
	Yfs
	Ygs
	Ygdtr
	Yidtr
	Yldtr
	Ymsw
	Ytask
	Ycr0
	Ycr1
	Ycr2
	Ycr3
	Ycr4
	Ycr5
	Ycr6
	Ycr7
	Ycr8
	Ydr0
	Ydr1
	Ydr2
	Ydr3
	Ydr4
	Ydr5
	Ydr6
	Ydr7
	Ytr0
	Ytr1
	Ytr2
	Ytr3
	Ytr4
	Ytr5
	Ytr6
	Ytr7
	Ymr
	Ymm
	Yxr0
	YxrEvexMulti4
	Yxr
	YxrEvex
	Yxm
	YxmEvex
	Yxvm
	YxvmEvex
	YyrEvexMulti4
	Yyr
	YyrEvex
	Yym
	YymEvex
	Yyvm
	YyvmEvex
	YzrMulti4
	Yzr
	Yzm
	Yzvm
	Yk0
	Yknot0
	Yk
	Ykm
	Ytls
	Ytextsize
	Yindir
	Ymax
)

const (
	Zxxx = iota
	Zlit
	Zlitm_r
	Zlitr_m
	Zlit_m_r
	Z_rp
	Zbr
	Zcall
	Zcallcon
	Zcallduff
	Zcallind
	Zcallindreg
	Zib_
	Zib_rp
	Zibo_m
	Zibo_m_xm
	Zil_
	Zil_rp
	Ziq_rp
	Zilo_m
	Zjmp
	Zjmpcon
	Zloop
	Zo_iw
	Zm_o
	Zm_r
	Z_m_r
	Zm2_r
	Zm_r_xm
	Zm_r_i_xm
	Zm_r_xm_nr
	Zr_m_xm_nr
	Zibm_r
	Zibr_m
	Zmb_r
	Zaut_r
	Zo_m
	Zo_m64
	Zpseudo
	Zr_m
	Zr_m_xm
	Zrp_
	Z_ib
	Z_il
	Zm_ibo
	Zm_ilo
	Zib_rr
	Zil_rr
	Zbyte

	Zvex_rm_v_r
	Zvex_rm_v_ro
	Zvex_r_v_rm
	Zvex_i_rm_vo
	Zvex_v_rm_r
	Zvex_i_rm_r
	Zvex_i_r_v
	Zvex_i_rm_v_r
	Zvex
	Zvex_rm_r_vo
	Zvex_i_r_rm
	Zvex_hr_rm_v_r

	Zevex_first
	Zevex_i_r_k_rm
	Zevex_i_r_rm
	Zevex_i_rm_k_r
	Zevex_i_rm_k_vo
	Zevex_i_rm_r
	Zevex_i_rm_v_k_r
	Zevex_i_rm_v_r
	Zevex_i_rm_vo
	Zevex_k_rmo
	Zevex_r_k_rm
	Zevex_r_v_k_rm
	Zevex_r_v_rm
	Zevex_rm_k_r
	Zevex_rm_v_k_r
	Zevex_rm_v_r
	Zevex_last

	Zmax
)

const (
	Px   = 0
	Px1  = 1
	P32  = 0x32
	Pe   = 0x66
	Pm   = 0x0f
	Pq   = 0xff
	Pb   = 0xfe
	Pf2  = 0xf2
	Pf3  = 0xf3
	Pef3 = 0xf5
	Pq3  = 0x67
	Pq4  = 0x68
	Pq4w = 0x69
	Pq5  = 0x6a
	Pq5w = 0x6b
	Pfw  = 0xf4
	Pw   = 0x48
	Pw8  = 0x90
	Py   = 0x80
	Py1  = 0x81
	Py3  = 0x83
	Pavx = 0x84

	RxrEvex = 1 << 4
	Rxw     = 1 << 3
	Rxr     = 1 << 2
	Rxx     = 1 << 1
	Rxb     = 1 << 0
)

// AsmBuf is a simple buffer to assemble variable-length x86 instructions into
// and hold assembly state.
type AsmBuf struct {
	buf      [100]byte
	off      int
	rexflag  int
	vexflag  bool
	evexflag bool
	rep      bool
	repn     bool
	lock     bool

	evex evexBits
}

// Put1 appends one byte to the end of the buffer.
func (ab *AsmBuf) Put1(x byte)

// Put2 appends two bytes to the end of the buffer.
func (ab *AsmBuf) Put2(x, y byte)

// Put3 appends three bytes to the end of the buffer.
func (ab *AsmBuf) Put3(x, y, z byte)

// Put4 appends four bytes to the end of the buffer.
func (ab *AsmBuf) Put4(x, y, z, w byte)

// PutInt16 writes v into the buffer using little-endian encoding.
func (ab *AsmBuf) PutInt16(v int16)

// PutInt32 writes v into the buffer using little-endian encoding.
func (ab *AsmBuf) PutInt32(v int32)

// PutInt64 writes v into the buffer using little-endian encoding.
func (ab *AsmBuf) PutInt64(v int64)

// Put copies b into the buffer.
func (ab *AsmBuf) Put(b []byte)

// PutOpBytesLit writes zero terminated sequence of bytes from op,
// starting at specified offset (e.g. z counter value).
// Trailing 0 is not written.
//
// Intended to be used for literal Z cases.
// Literal Z cases usually have "Zlit" in their name (Zlit, Zlitr_m, Zlitm_r).
func (ab *AsmBuf) PutOpBytesLit(offset int, op *opBytes)

// Insert inserts b at offset i.
func (ab *AsmBuf) Insert(i int, b byte)

// Last returns the byte at the end of the buffer.
func (ab *AsmBuf) Last() byte

// Len returns the length of the buffer.
func (ab *AsmBuf) Len() int

// Bytes returns the contents of the buffer.
func (ab *AsmBuf) Bytes() []byte

// Reset empties the buffer.
func (ab *AsmBuf) Reset()

// At returns the byte at offset i.
func (ab *AsmBuf) At(i int) byte
