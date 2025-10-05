// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

const ThatPeskyUnicodeDot = "\u00b7"

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
