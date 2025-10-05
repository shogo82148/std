// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asmgen

var ArchAMD64 = &Arch{
	Name:      "amd64",
	WordBits:  64,
	WordBytes: 8,

	regs: []string{
		"BX", "SI", "DI",
		"R8", "R9", "R10", "R11", "R12", "R13", "R14", "R15",
		"AX", "DX", "CX",
	},
	op3:              x86Op3,
	hint:             x86Hint,
	memOK:            true,
	subCarryIsBorrow: true,

	options: map[Option]func(*Asm, string){
		OptionAltCarry: amd64JmpADX,
	},

	mov:      "MOVQ",
	adds:     "ADDQ",
	adcs:     "ADCQ",
	subs:     "SUBQ",
	sbcs:     "SBBQ",
	lsh:      "SHLQ",
	lshd:     "SHLQ",
	rsh:      "SHRQ",
	rshd:     "SHRQ",
	and:      "ANDQ",
	or:       "ORQ",
	xor:      "XORQ",
	neg:      "NEGQ",
	lea:      "LEAQ",
	addF:     amd64Add,
	mulWideF: x86MulWide,

	addWords: "LEAQ (%[2]s)(%[1]s*8), %[3]s",

	jmpZero:       "TESTQ %[1]s, %[1]s; JZ %[2]s",
	jmpNonZero:    "TESTQ %[1]s, %[1]s; JNZ %[2]s",
	loopBottom:    "SUBQ $1, %[1]s; JNZ %[2]s",
	loopBottomNeg: "ADDQ $1, %[1]s; JNZ %[2]s",
}
