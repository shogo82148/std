// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var (
	XDWORD0 VecPhysical = Y4
	XDWORD1             = Y5
	XDWORD2             = Y6
	XDWORD3             = Y7

	XWORD0 = X4
	XWORD1 = X5
	XWORD2 = X6
	XWORD3 = X7

	XTMP0 = Y0
	XTMP1 = Y1
	XTMP2 = Y2
	XTMP3 = Y3
	XTMP4 = Y8
	XTMP5 = Y11

	XFER = Y9

	BYTE_FLIP_MASK   = Y13
	X_BYTE_FLIP_MASK = X13

	NUM_BYTES GPPhysical = RDX
	INP                  = RDI

	CTX = RSI

	TBL = RBP

	SRND = RSI

	T1 = R12L

	// Offsets
	XFER_SIZE    = 2 * 64 * 4
	INP_END_SIZE = 8
	INP_SIZE     = 8

	STACK_SIZE = _INP + INP_SIZE
)

// Pointers for memoizing Data section symbols
var K256Ptr *Mem

// Round specific constants
func K256_DATA() Mem
