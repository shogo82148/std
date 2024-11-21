// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This code was translated into a form compatible with 6a from the public
// domain sources at https://github.com/gvanas/KeccakCodePackage

package main

// Round Constants for use in the ι step.
var RoundConstants = [24]uint64{
	0x0000000000000001,
	0x0000000000008082,
	0x800000000000808A,
	0x8000000080008000,
	0x000000000000808B,
	0x0000000080000001,
	0x8000000080008081,
	0x8000000000008009,
	0x000000000000008A,
	0x0000000000000088,
	0x0000000080008009,
	0x000000008000000A,
	0x000000008000808B,
	0x800000000000008B,
	0x8000000000008089,
	0x8000000000008003,
	0x8000000000008002,
	0x8000000000000080,
	0x000000000000800A,
	0x800000008000000A,
	0x8000000080008081,
	0x8000000000008080,
	0x0000000080000001,
	0x8000000080008008,
}

func MOVQ_RBI_RCE()
func XORQ_RT1_RCA()
func XORQ_RT1_RCE()
func XORQ_RBA_RCU()
func XORQ_RBE_RCU()
func XORQ_RDU_RCU()
func XORQ_RDA_RCA()
func XORQ_RDE_RCE()

type ArgMacro func()
