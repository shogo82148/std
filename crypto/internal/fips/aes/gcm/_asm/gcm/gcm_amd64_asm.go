// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is an optimized implementation of AES-GCM using AES-NI and CLMUL-NI
// The implementation uses some optimization as described in:
// [1] Gueron, S., Kounavis, M.E.: Intel® Carry-Less Multiplication
//     Instruction and its Usage for Computing the GCM Mode rev. 2.02
// [2] Gueron, S., Krasnov, V.: Speeding up Counter Mode in Software and
//     Hardware

package main

var (
	B0 VecPhysical = X0
	B1             = X1
	B2             = X2
	B3             = X3
	B4             = X4
	B5             = X5
	B6             = X6
	B7             = X7

	ACC0 VecPhysical = X8
	ACC1             = X9
	ACCM             = X10

	T0    VecPhysical = X11
	T1                = X12
	T2                = X13
	POLY              = X14
	BSWAP             = X15
)
