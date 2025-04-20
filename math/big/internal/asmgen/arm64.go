// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asmgen

var ArchARM64 = &Arch{
	Name:          "arm64",
	WordBits:      64,
	WordBytes:     8,
	CarrySafeLoop: true,

	regs: []string{

		"R0", "R1", "R2", "R3", "R4", "R5", "R6", "R7", "R8", "R9",
		"R10", "R11", "R12", "R13", "R14", "R15", "R16", "R17", "R19",
		"R20", "R21", "R22", "R23", "R24", "R25", "R26",
	},
	reg0: "ZR",

	mov:   "MOVD",
	add:   "ADD",
	adds:  "ADDS",
	adc:   "ADC",
	adcs:  "ADCS",
	sub:   "SUB",
	subs:  "SUBS",
	sbc:   "SBC",
	sbcs:  "SBCS",
	mul:   "MUL",
	mulhi: "UMULH",
	lsh:   "LSL",
	rsh:   "LSR",
	and:   "AND",
	or:    "ORR",
	xor:   "EOR",

	addWords: "ADD %[1]s<<3, %[2]s, %[3]s",

	jmpZero:    "CBZ %s, %s",
	jmpNonZero: "CBNZ %s, %s",

	loadIncN:  arm64LoadIncN,
	loadDecN:  arm64LoadDecN,
	storeIncN: arm64StoreIncN,
	storeDecN: arm64StoreDecN,
}
