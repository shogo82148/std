// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asmgen

var ArchARM = &Arch{
	Name:          "arm",
	WordBits:      32,
	WordBytes:     4,
	CarrySafeLoop: true,

	regs: []string{

		"R0", "R1", "R2", "R3", "R4", "R5", "R6", "R7", "R8", "R9", "R11", "R12",
	},
	regShift: true,

	mov:  "MOVW",
	add:  "ADD",
	adds: "ADD.S",
	adc:  "ADC",
	adcs: "ADC.S",
	sub:  "SUB",
	subs: "SUB.S",
	sbc:  "SBC",
	sbcs: "SBC.S",
	rsb:  "RSB",
	and:  "AND",
	or:   "ORR",
	xor:  "EOR",

	mulWideF: armMulWide,

	addWords: "ADD %s<<2, %s, %s",

	jmpZero:    "TEQ $0, %s; BEQ %s",
	jmpNonZero: "TEQ $0, %s; BNE %s",

	loadIncN:  armLoadIncN,
	loadDecN:  armLoadDecN,
	storeIncN: armStoreIncN,
	storeDecN: armStoreDecN,
}
