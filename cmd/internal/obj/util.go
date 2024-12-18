// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package obj

import (
	"github.com/shogo82148/std/io"
)

const REG_NONE = 0

// Line returns a string containing the filename and line number for p
func (p *Prog) Line() string

func (p *Prog) InnermostLine(w io.Writer)

// InnermostLineNumber returns a string containing the line number for the
// innermost inlined function (if any inlining) at p's position
func (p *Prog) InnermostLineNumber() string

// InnermostLineNumberHTML returns a string containing the line number for the
// innermost inlined function (if any inlining) at p's position
func (p *Prog) InnermostLineNumberHTML() string

// InnermostFilename returns a string containing the innermost
// (in inlining) filename at p's position
func (p *Prog) InnermostFilename() string

/* ARM scond byte */
const (
	C_SCOND     = (1 << 4) - 1
	C_SBIT      = 1 << 4
	C_PBIT      = 1 << 5
	C_WBIT      = 1 << 6
	C_FBIT      = 1 << 7
	C_UBIT      = 1 << 7
	C_SCOND_XOR = 14
)

// CConv formats opcode suffix bits (Prog.Scond).
func CConv(s uint8) string

// CConvARM formats ARM opcode suffix bits (mostly condition codes).
func CConvARM(s uint8) string

func (p *Prog) String() string

func (p *Prog) InnermostString(w io.Writer)

// InstructionString returns a string representation of the instruction without preceding
// program counter or file and line number.
func (p *Prog) InstructionString() string

// WriteInstructionString writes a string representation of the instruction without preceding
// program counter or file and line number.
func (p *Prog) WriteInstructionString(w io.Writer)

func (ctxt *Link) NewProg() *Prog

func (ctxt *Link) CanReuseProgs() bool

// Dconv accepts an argument 'a' within a prog 'p' and returns a string
// with a formatted version of the argument.
func Dconv(p *Prog, a *Addr) string

// DconvWithABIDetail accepts an argument 'a' within a prog 'p'
// and returns a string with a formatted version of the argument, in
// which text symbols are rendered with explicit ABI selectors.
func DconvWithABIDetail(p *Prog, a *Addr) string

// WriteDconv accepts an argument 'a' within a prog 'p'
// and writes a formatted version of the arg to the writer.
func WriteDconv(w io.Writer, p *Prog, a *Addr)

func (a *Addr) WriteNameTo(w io.Writer)

// RegisterOpSuffix assigns cconv function for formatting opcode suffixes
// when compiling for GOARCH=arch.
//
// cconv is never called with 0 argument.
func RegisterOpSuffix(arch string, cconv func(uint8) string)

const (
	// Because of masking operations in the encodings, each register
	// space should start at 0 modulo some power of 2.
	RBase386     = 1 * 1024
	RBaseAMD64   = 2 * 1024
	RBaseARM     = 3 * 1024
	RBasePPC64   = 4 * 1024
	RBaseARM64   = 8 * 1024
	RBaseMIPS    = 13 * 1024
	RBaseS390X   = 14 * 1024
	RBaseRISCV   = 15 * 1024
	RBaseWasm    = 16 * 1024
	RBaseLOONG64 = 19 * 1024
)

// RegisterRegister binds a pretty-printer (Rconv) for register
// numbers to a given register number range. Lo is inclusive,
// hi exclusive (valid registers are lo through hi-1).
func RegisterRegister(lo, hi int, Rconv func(int) string)

func Rconv(reg int) string

// Each architecture is allotted a distinct subspace: [Lo, Hi) for declaring its
// arch-specific register list numbers.
const (
	RegListARMLo = 0
	RegListARMHi = 1 << 16

	// arm64 uses the 60th bit to differentiate from other archs
	RegListARM64Lo = 1 << 60
	RegListARM64Hi = 1<<61 - 1

	// x86 uses the 61th bit to differentiate from other archs
	RegListX86Lo = 1 << 61
	RegListX86Hi = 1<<62 - 1
)

// RegisterRegisterList binds a pretty-printer (RLconv) for register list
// numbers to a given register list number range. Lo is inclusive,
// hi exclusive (valid register list are lo through hi-1).
func RegisterRegisterList(lo, hi int64, rlconv func(int64) string)

func RLconv(list int64) string

// RegisterSpecialOperands binds a pretty-printer (SPCconv) for special
// operand numbers to a given special operand number range. Lo is inclusive,
// hi is exclusive (valid special operands are lo through hi-1).
func RegisterSpecialOperands(lo, hi int64, rlconv func(int64) string)

// SPCconv returns the string representation of the special operand spc.
func SPCconv(spc int64) string

// RegisterOpcode binds a list of instruction names
// to a given instruction number range.
func RegisterOpcode(lo As, Anames []string)

func (a As) String() string

var Anames = []string{
	"XXX",
	"CALL",
	"DUFFCOPY",
	"DUFFZERO",
	"END",
	"FUNCDATA",
	"JMP",
	"NOP",
	"PCALIGN",
	"PCALIGNMAX",
	"PCDATA",
	"RET",
	"GETCALLERPC",
	"TEXT",
	"UNDEF",
}

func Bool2int(b bool) int

// AlignmentPadding bytes to add to align code as requested.
// Alignment is restricted to powers of 2 between 8 and 2048 inclusive.
//
// pc_: current offset in function, in bytes
// p:  a PCALIGN or PCALIGNMAX prog
// ctxt: the context, for current function
// cursym: current function being assembled
// returns number of bytes of padding needed,
// updates minimum alignment for the function.
func AlignmentPadding(pc int32, p *Prog, ctxt *Link, cursym *LSym) int

// AlignmentPaddingLength is the number of bytes to add to align code as requested.
// Alignment is restricted to powers of 2 between 8 and 2048 inclusive.
// This only computes the length and does not update the (missing parameter)
// current function's own required alignment.
//
// pc: current offset in function, in bytes
// p:  a PCALIGN or PCALIGNMAX prog
// ctxt: the context, for current function
// returns number of bytes of padding needed,
func AlignmentPaddingLength(pc int32, p *Prog, ctxt *Link) int
