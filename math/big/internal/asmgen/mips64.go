// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asmgen

var ArchMIPS64x = &Arch{
	Name:          "mips64x",
	Build:         "mips64 || mips64le",
	WordBits:      64,
	WordBytes:     8,
	CarrySafeLoop: true,

	regs: []string{

		"R1", "R2", "R3", "R4", "R5", "R6", "R7", "R8", "R9",
		"R10", "R11", "R12", "R13", "R14", "R15", "R16", "R17", "R18", "R19",
		"R20", "R21", "R22", "R24", "R25",
	},
	reg0:        "R0",
	regTmp:      "R23",
	regCarry:    "R24",
	regAltCarry: "R25",

	mov:      "MOVV",
	add:      "ADDVU",
	sltu:     "SGTU",
	sub:      "SUBVU",
	mulWideF: mips64MulWide,
	lsh:      "SLLV",
	rsh:      "SRLV",
	and:      "AND",
	or:       "OR",
	xor:      "XOR",

	jmpZero:    "BEQ %s, %s",
	jmpNonZero: "BNE %s, %s",
}
