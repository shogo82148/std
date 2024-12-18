// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func LOAD(index int)

func SHUFFLE(index int)

func FUNC1(a, b, c, d, e GPPhysical)

func FUNC2(a, b, c, d, e GPPhysical)

func FUNC3(a, b, c, d, e GPPhysical)

func FUNC4(a, b, c, d, e GPPhysical)

func MIX(a, b, c, d, e GPPhysical, konst int)

func ROUND1(a, b, c, d, e GPPhysical, index int)

func ROUND1x(a, b, c, d, e GPPhysical, index int)

func ROUND2(a, b, c, d, e GPPhysical, index int)

func ROUND3(a, b, c, d, e GPPhysical, index int)

func ROUND4(a, b, c, d, e GPPhysical, index int)

func UPDATE_HASH(A, TB, C, D, E GPPhysical)

func PRECALC_0(OFFSET int)

func PRECALC_1(OFFSET int)

func PRECALC_2(YREG VecPhysical)

func PRECALC_4(YREG VecPhysical, K_OFFSET int)

func PRECALC_7(OFFSET int)

// Message scheduling pre-compute for rounds 0-15
//
//   - R13 is a pointer to even 64-byte block
//   - R10 is a pointer to odd 64-byte block
//   - R14 is a pointer to temp buffer
//   - X0 is used as temp register
//   - YREG is clobbered as part of computation
//   - OFFSET chooses 16 byte chunk within a block
//   - R8 is a pointer to constants block
//   - K_OFFSET chooses K constants relevant to this round
//   - X10 holds swap mask
func PRECALC_00_15(OFFSET int, YREG VecPhysical)

func PRECALC_16(REG_SUB_16, REG_SUB_12, REG_SUB_4, REG VecPhysical)

func PRECALC_17(REG_SUB_16, REG_SUB_8, REG VecPhysical)

func PRECALC_18(REG VecPhysical)

func PRECALC_19(REG VecPhysical)

func PRECALC_20(REG VecPhysical)

func PRECALC_21(REG VecPhysical)

func PRECALC_23(REG VecPhysical, K_OFFSET, OFFSET int)

// Message scheduling pre-compute for rounds 16-31
//   - calculating last 32 w[i] values in 8 XMM registers
//   - pre-calculate K+w[i] values and store to mem
//   - for later load by ALU add instruction.
//   - "brute force" vectorization for rounds 16-31 only
//   - due to w[i]->w[i-3] dependency.
//   - clobbers 5 input ymm registers REG_SUB*
//   - uses X0 and X9 as temp registers
//   - As always, R8 is a pointer to constants block
//   - and R14 is a pointer to temp buffer
func PRECALC_16_31(REG, REG_SUB_4, REG_SUB_8, REG_SUB_12, REG_SUB_16 VecPhysical, K_OFFSET, OFFSET int)

func PRECALC_32(REG_SUB_8, REG_SUB_4 VecPhysical)

func PRECALC_33(REG_SUB_28, REG VecPhysical)

func PRECALC_34(REG_SUB_16 VecPhysical)

func PRECALC_35(REG VecPhysical)

func PRECALC_36(REG VecPhysical)

func PRECALC_37(REG VecPhysical)

func PRECALC_39(REG VecPhysical, K_OFFSET, OFFSET int)

func PRECALC_32_79(REG, REG_SUB_4, REG_SUB_8, REG_SUB_16, REG_SUB_28 VecPhysical, K_OFFSET, OFFSET int)

func PRECALC()

func CALC_F1_PRE(OFFSET int, REG_A, REG_B, REG_C, REG_E GPPhysical)

func CALC_F1_POST(REG_A, REG_B, REG_E GPPhysical)

func CALC_0()

func CALC_1()

func CALC_2()

func CALC_3()

func CALC_4()

func CALC_5()

func CALC_6()

func CALC_7()

func CALC_8()

func CALC_9()

func CALC_10()

func CALC_11()

func CALC_12()

func CALC_13()

func CALC_14()

func CALC_15()

func CALC_16()

func CALC_17()

func CALC_18()

func CALC_F2_PRE(OFFSET int, REG_A, REG_B, REG_E GPPhysical)

func CALC_F2_POST(REG_A, REG_B, REG_C, REG_E GPPhysical)

func CALC_19()

func CALC_20()

func CALC_21()

func CALC_22()

func CALC_23()

func CALC_24()

func CALC_25()

func CALC_26()

func CALC_27()

func CALC_28()

func CALC_29()

func CALC_30()

func CALC_31()

func CALC_32()

func CALC_33()

func CALC_34()

func CALC_35()

func CALC_36()

func CALC_37()

func CALC_38()

func CALC_F3_PRE(OFFSET int, REG_E GPPhysical)

func CALC_F3_POST(REG_A, REG_B, REG_C, REG_E, REG_TB GPPhysical)

func CALC_39()

func CALC_40()

func CALC_41()

func CALC_42()

func CALC_43()

func CALC_44()

func CALC_45()

func CALC_46()

func CALC_47()

func CALC_48()

func CALC_49()

func CALC_50()

func CALC_51()

func CALC_52()

func CALC_53()

func CALC_54()

func CALC_55()

func CALC_56()

func CALC_57()

func CALC_58()

func CALC_59()

func CALC_60()

func CALC_61()

func CALC_62()

func CALC_63()

func CALC_64()

func CALC_65()

func CALC_66()

func CALC_67()

func CALC_68()

func CALC_69()

func CALC_70()

func CALC_71()

func CALC_72()

func CALC_73()

func CALC_74()

func CALC_75()

func CALC_76()

func CALC_77()

func CALC_78()

func CALC_79()

// Similar to CALC_0
func CALC_80()

func CALC_81()

func CALC_82()

func CALC_83()

func CALC_84()

func CALC_85()

func CALC_86()

func CALC_87()

func CALC_88()

func CALC_89()

func CALC_90()

func CALC_91()

func CALC_92()

func CALC_93()

func CALC_94()

func CALC_95()

func CALC_96()

func CALC_97()

func CALC_98()

func CALC_99()

func CALC_100()

func CALC_101()

func CALC_102()

func CALC_103()

func CALC_104()

func CALC_105()

func CALC_106()

func CALC_107()

func CALC_108()

func CALC_109()

func CALC_110()

func CALC_111()

func CALC_112()

func CALC_113()

func CALC_114()

func CALC_115()

func CALC_116()

func CALC_117()

func CALC_118()

func CALC_119()

func CALC_120()

func CALC_121()

func CALC_122()

func CALC_123()

func CALC_124()

func CALC_125()

func CALC_126()

func CALC_127()

func CALC_128()

func CALC_129()

func CALC_130()

func CALC_131()

func CALC_132()

func CALC_133()

func CALC_134()

func CALC_135()

func CALC_136()

func CALC_137()

func CALC_138()

func CALC_139()

func CALC_140()

func CALC_141()

func CALC_142()

func CALC_143()

func CALC_144()

func CALC_145()

func CALC_146()

func CALC_147()

func CALC_148()

func CALC_149()

func CALC_150()

func CALC_151()

func CALC_152()

func CALC_153()

func CALC_154()

func CALC_155()

func CALC_156()

func CALC_157()

func CALC_158()

func CALC_159()

func CALC()

// Pointers for memoizing Data section symbols
var (
	K_XMM_AR_ptr, BSWAP_SHUFB_CTL_ptr *Mem
)

func K_XMM_AR_DATA() Mem

var BSWAP_SHUFB_CTL_CONSTANTS = [8]uint32{
	0x00010203,
	0x04050607,
	0x08090a0b,
	0x0c0d0e0f,
	0x00010203,
	0x04050607,
	0x08090a0b,
	0x0c0d0e0f,
}

func BSWAP_SHUFB_CTL_DATA() Mem
