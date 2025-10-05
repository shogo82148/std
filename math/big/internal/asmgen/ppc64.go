// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asmgen

var ArchPPC64x = &Arch{
	Name:          "ppc64x",
	Build:         "ppc64 || ppc64le",
	WordBits:      64,
	WordBytes:     8,
	CarrySafeLoop: true,

	regs: []string{

		"R3", "R4", "R5", "R6", "R7", "R8", "R9",
		"R10", "R11", "R12", "R14", "R15", "R16", "R17", "R18", "R19",
		"R20", "R21", "R22", "R23", "R24", "R25", "R26", "R27", "R28", "R29",
	},
	reg0:   "R0",
	regTmp: "R31",

	mov:   "MOVD",
	add:   "ADD",
	adds:  "ADDC",
	adcs:  "ADDE",
	sub:   "SUB",
	subs:  "SUBC",
	sbcs:  "SUBE",
	mul:   "MULLD",
	mulhi: "MULHDU",
	lsh:   "SLD",
	rsh:   "SRD",
	and:   "ANDCC",
	or:    "OR",
	xor:   "XOR",

	jmpZero:    "CMP %[1]s, $0; BEQ %[2]s",
	jmpNonZero: "CMP %s, $0; BNE %s",

	loopTop:    "CMP %[1]s, $0; BEQ %[2]s; MOVD %[1]s, CTR",
	loopBottom: "BDNZ %[2]s",
}
