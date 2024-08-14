// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loong64

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

const (
	NSNAME = 8
	NSYM   = 50
	NREG   = 32
	NFREG  = 32
)

const (
	REG_R0 = obj.RBaseLOONG64 + iota
	REG_R1
	REG_R2
	REG_R3
	REG_R4
	REG_R5
	REG_R6
	REG_R7
	REG_R8
	REG_R9
	REG_R10
	REG_R11
	REG_R12
	REG_R13
	REG_R14
	REG_R15
	REG_R16
	REG_R17
	REG_R18
	REG_R19
	REG_R20
	REG_R21
	REG_R22
	REG_R23
	REG_R24
	REG_R25
	REG_R26
	REG_R27
	REG_R28
	REG_R29
	REG_R30
	REG_R31

	REG_F0
	REG_F1
	REG_F2
	REG_F3
	REG_F4
	REG_F5
	REG_F6
	REG_F7
	REG_F8
	REG_F9
	REG_F10
	REG_F11
	REG_F12
	REG_F13
	REG_F14
	REG_F15
	REG_F16
	REG_F17
	REG_F18
	REG_F19
	REG_F20
	REG_F21
	REG_F22
	REG_F23
	REG_F24
	REG_F25
	REG_F26
	REG_F27
	REG_F28
	REG_F29
	REG_F30
	REG_F31

	REG_FCSR0
	REG_FCSR1
	REG_FCSR2
	REG_FCSR3
	REG_FCSR4
	REG_FCSR5
	REG_FCSR6
	REG_FCSR7
	REG_FCSR8
	REG_FCSR9
	REG_FCSR10
	REG_FCSR11
	REG_FCSR12
	REG_FCSR13
	REG_FCSR14
	REG_FCSR15
	REG_FCSR16
	REG_FCSR17
	REG_FCSR18
	REG_FCSR19
	REG_FCSR20
	REG_FCSR21
	REG_FCSR22
	REG_FCSR23
	REG_FCSR24
	REG_FCSR25
	REG_FCSR26
	REG_FCSR27
	REG_FCSR28
	REG_FCSR29
	REG_FCSR30
	REG_FCSR31

	REG_FCC0
	REG_FCC1
	REG_FCC2
	REG_FCC3
	REG_FCC4
	REG_FCC5
	REG_FCC6
	REG_FCC7
	REG_FCC8
	REG_FCC9
	REG_FCC10
	REG_FCC11
	REG_FCC12
	REG_FCC13
	REG_FCC14
	REG_FCC15
	REG_FCC16
	REG_FCC17
	REG_FCC18
	REG_FCC19
	REG_FCC20
	REG_FCC21
	REG_FCC22
	REG_FCC23
	REG_FCC24
	REG_FCC25
	REG_FCC26
	REG_FCC27
	REG_FCC28
	REG_FCC29
	REG_FCC30
	REG_FCC31

	REG_LAST = REG_FCC31

	REG_SPECIAL = REG_FCSR0

	REGZERO = REG_R0
	REGLINK = REG_R1
	REGSP   = REG_R3
	REGRET  = REG_R20
	REGARG  = -1
	REGRT1  = REG_R20
	REGRT2  = REG_R21
	REGCTXT = REG_R29
	REGG    = REG_R22
	REGTMP  = REG_R30
	FREGRET = REG_F0
)

var LOONG64DWARFRegisters = map[int16]int16{}

const (
	BIG = 2046
)

const (
	// mark flags
	LABEL  = 1 << 0
	LEAF   = 1 << 1
	SYNC   = 1 << 2
	BRANCH = 1 << 3
)

const (
	C_NONE = iota
	C_REG
	C_FREG
	C_FCSRREG
	C_FCCREG
	C_ZCON
	C_SCON
	C_UCON
	C_ADD0CON
	C_AND0CON
	C_ADDCON
	C_ANDCON
	C_LCON
	C_DCON
	C_SACON
	C_SECON
	C_LACON
	C_LECON
	C_DACON
	C_STCON
	C_SBRA
	C_LBRA
	C_SAUTO
	C_LAUTO
	C_SEXT
	C_LEXT
	C_ZOREG
	C_SOREG
	C_LOREG
	C_GOK
	C_ADDR
	C_TLS_LE
	C_TLS_IE
	C_GOTADDR
	C_TEXTSIZE

	C_NCLASS
)

const (
	AABSD = obj.ABaseLoong64 + obj.A_ARCHSPECIFIC + iota
	AABSF
	AADD
	AADDD
	AADDF
	AADDU

	AADDW
	AAND
	ABEQ
	ABGEZ
	ABLEZ
	ABGTZ
	ABLTZ
	ABFPF
	ABFPT

	ABNE
	ABREAK
	ACLO
	ACLZ

	ACMPEQD
	ACMPEQF

	ACMPGED
	ACMPGEF
	ACMPGTD
	ACMPGTF

	ALU12IW
	ALU32ID
	ALU52ID
	APCALAU12I
	APCADDU12I
	AJIRL
	ABGE
	ABLT
	ABLTU
	ABGEU

	ADIV
	ADIVD
	ADIVF
	ADIVU
	ADIVW

	ALL
	ALLV

	ALUI

	AMOVB
	AMOVBU

	AMOVD
	AMOVDF
	AMOVDW
	AMOVF
	AMOVFD
	AMOVFW

	AMOVH
	AMOVHU
	AMOVW

	AMOVWD
	AMOVWF

	AMOVWL
	AMOVWR

	AMUL
	AMULD
	AMULF
	AMULU
	AMULH
	AMULHU
	AMULW
	ANEGD
	ANEGF

	ANEGW
	ANEGV

	ANOOP
	ANOR
	AOR
	AREM
	AREMU

	ARFE

	ASC
	ASCV

	ASGT
	ASGTU

	ASLL
	ASQRTD
	ASQRTF
	ASRA
	ASRL
	AROTR
	ASUB
	ASUBD
	ASUBF

	ASUBU
	ASUBW
	ADBAR
	ASYSCALL

	ATEQ
	ATNE

	AWORD

	AXOR

	AMASKEQZ
	AMASKNEZ

	// 64-bit
	AMOVV
	AMOVVL
	AMOVVR

	ASLLV
	ASRAV
	ASRLV
	AROTRV
	ADIVV
	ADIVVU

	AREMV
	AREMVU

	AMULV
	AMULVU
	AMULHV
	AMULHVU
	AADDV
	AADDVU
	ASUBV
	ASUBVU

	// 64-bit FP
	ATRUNCFV
	ATRUNCDV
	ATRUNCFW
	ATRUNCDW

	AMOVWU
	AMOVFV
	AMOVDV
	AMOVVF
	AMOVVD

	// 2.2.7. Atomic Memory Access Instructions
	AAMSWAPB
	AAMSWAPH
	AAMSWAPW
	AAMSWAPV
	AAMCASB
	AAMCASH
	AAMCASW
	AAMCASV
	AAMADDW
	AAMADDV
	AAMANDW
	AAMANDV
	AAMORW
	AAMORV
	AAMXORW
	AAMXORV
	AAMMAXW
	AAMMAXV
	AAMMINW
	AAMMINV
	AAMMAXWU
	AAMMAXVU
	AAMMINWU
	AAMMINVU
	AAMSWAPDBB
	AAMSWAPDBH
	AAMSWAPDBW
	AAMSWAPDBV
	AAMCASDBB
	AAMCASDBH
	AAMCASDBW
	AAMCASDBV
	AAMADDDBW
	AAMADDDBV
	AAMANDDBW
	AAMANDDBV
	AAMORDBW
	AAMORDBV
	AAMXORDBW
	AAMXORDBV
	AAMMAXDBW
	AAMMAXDBV
	AAMMINDBW
	AAMMINDBV
	AAMMAXDBWU
	AAMMAXDBVU
	AAMMINDBWU
	AAMMINDBVU

	// 2.2.10. Other Miscellaneous Instructions
	ARDTIMELW
	ARDTIMEHW
	ARDTIMED
	ACPUCFG

	ALAST

	// aliases
	AJMP = obj.AJMP
	AJAL = obj.ACALL
	ARET = obj.ARET
)
