// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

const ThatPeskyUnicodeDot = "\u00b7"

// Wt = Mt; for 0 <= t <= 15
//
// Line 50
func MSGSCHEDULE0(index int)

// Wt = SIGMA1(Wt-2) + Wt-7 + SIGMA0(Wt-15) + Wt-16; for 16 <= t <= 79
//
//	SIGMA0(x) = ROTR(1,x) XOR ROTR(8,x) XOR SHR(7,x)
//	SIGMA1(x) = ROTR(19,x) XOR ROTR(61,x) XOR SHR(6,x)
//
// Line 58
func MSGSCHEDULE1(index int)

// Calculate T1 in AX - uses AX, CX and DX registers.
// h is also used as an accumulator. Wt is passed in AX.
//
//	T1 = h + BIGSIGMA1(e) + Ch(e, f, g) + Kt + Wt
//	  BIGSIGMA1(x) = ROTR(14,x) XOR ROTR(18,x) XOR ROTR(41,x)
//	  Ch(x, y, z) = (x AND y) XOR (NOT x AND z)
//
// Line 85
func SHA512T1(konst uint64, e, f, g, h GPPhysical)

// Calculate T2 in BX - uses BX, CX, DX and DI registers.
//
//	T2 = BIGSIGMA0(a) + Maj(a, b, c)
//	  BIGSIGMA0(x) = ROTR(28,x) XOR ROTR(34,x) XOR ROTR(39,x)
//	  Maj(x, y, z) = (x AND y) XOR (x AND z) XOR (y AND z)
//
// Line 110
func SHA512T2(a, b, c GPPhysical)

// Calculate T1 and T2, then e = d + T1 and a = T1 + T2.
// The values for e and a are stored in d and h, ready for rotation.
//
// Line 131
func SHA512ROUND(index int, konst uint64, a, b, c, d, e, f, g, h GPPhysical)

// Line 169
func SHA512ROUND0(index int, konst uint64, a, b, c, d, e, f, g, h GPPhysical)

// Line 142
func SHA512ROUND1(index int, konst uint64, a, b, c, d, e, f, g, h GPPhysical)

// Line 289
var (
	YFER_SIZE int = (4 * 8)
	SRND_SIZE     = (1 * 8)
	INP_SIZE      = (1 * 8)
)

// Line 302
func COPY_YMM_AND_BSWAP(p1 VecPhysical, p2 Mem, p3 VecPhysical)

// Line 306
func MY_VPALIGNR(YDST, YSRC1, YSRC2 VecPhysical, RVAL int)

// Pointers for memoizing Data section symbols
var PSHUFFLE_BYTE_FLIP_MASK_DATA_ptr, MASK_YMM_LO_ptr *Mem

// Line 310
func PSHUFFLE_BYTE_FLIP_MASK_DATA() Mem

// Line 317
func MASK_YMM_LO_DATA() Mem
