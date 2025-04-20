// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asmgen

var Arch386 = &Arch{
	Name:      "386",
	WordBits:  32,
	WordBytes: 4,

	regs: []string{
		"BX", "SI", "DI", "BP",
		"CX", "DX", "AX",
	},
	op3:              x86Op3,
	hint:             x86Hint,
	memOK:            true,
	subCarryIsBorrow: true,
	maxColumns:       1,

	memIndex: _386MemIndex,

	mov:      "MOVL",
	adds:     "ADDL",
	adcs:     "ADCL",
	subs:     "SUBL",
	sbcs:     "SBBL",
	lsh:      "SHLL",
	lshd:     "SHLL",
	rsh:      "SHRL",
	rshd:     "SHRL",
	and:      "ANDL",
	or:       "ORL",
	xor:      "XORL",
	neg:      "NEGL",
	lea:      "LEAL",
	mulWideF: x86MulWide,

	addWords: "LEAL (%[2]s)(%[1]s*4), %[3]s",

	jmpZero:       "TESTL %[1]s, %[1]s; JZ %[2]s",
	jmpNonZero:    "TESTL %[1]s, %[1]s; JNZ %[2]s",
	loopBottom:    "SUBL $1, %[1]s; JNZ %[2]s",
	loopBottomNeg: "ADDL $1, %[1]s; JNZ %[2]s",
}
