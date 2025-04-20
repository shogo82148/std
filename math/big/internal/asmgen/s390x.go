// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asmgen

var ArchS390X = &Arch{
	Name:          "s390x",
	WordBits:      64,
	WordBytes:     8,
	CarrySafeLoop: true,

	regs: []string{

		"R1", "R2", "R3", "R4", "R5", "R6", "R7", "R8", "R9",
		"R10", "R11", "R12",
	},
	reg0:       "R0",
	regTmp:     "R10",
	setup:      s390xSetup,
	maxColumns: 2,
	op3:        s390xOp3,
	hint:       s390xHint,

	mov:      "MOVD",
	adds:     "ADDC",
	adcs:     "ADDE",
	subs:     "SUBC",
	sbcs:     "SUBE",
	mulWideF: s390MulWide,
	lsh:      "SLD",
	rsh:      "SRD",
	and:      "AND",
	or:       "OR",
	xor:      "XOR",
	neg:      "NEG",
	lea:      "LAY",

	jmpZero:    "CMPBEQ %s, $0, %s",
	jmpNonZero: "CMPBNE %s, $0, %s",
}
