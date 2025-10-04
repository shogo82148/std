// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asmgen

import (
	"github.com/shogo82148/std/bytes"
)

// An Asm is an assembly file being written.
type Asm struct {
	Arch     *Arch
	out      bytes.Buffer
	regavail uint64
	enabled  map[Option]bool
}

// NewAsm returns a new Asm preparing assembly
// for the given architecture to be written to file.
func NewAsm(arch *Arch) *Asm

// Fatalf reports a fatal error by panicking.
// Panicking is appropriate because there is a bug in the generator,
// and panicking will show the exact source lines leading to that bug.
func (a *Asm) Fatalf(format string, args ...any)

// ZR returns the zero register (the specific register guaranteed to hold the integer 0),
// or else the zero Reg (Reg{}, which has r.Valid() == false).
func (a *Asm) ZR() Reg

// Carry returns the carry register, or else the zero Reg.
func (a *Asm) Carry() Reg

// AltCarry returns the secondary carry register, or else the zero Reg.
func (a *Asm) AltCarry() Reg

// Imm returns a Reg representing an immediate (constant) value.
func (a *Asm) Imm(x int) Reg

// IsZero reports whether r is a zero immediate or the zero register.
func (a *Asm) IsZero(r Reg) bool

// Reg allocates a new register.
func (a *Asm) Reg() Reg

// RegHint allocates a new register, with a hint as to its purpose.
func (a *Asm) RegHint(hint Hint) Reg

// Free frees a previously allocated register.
// If r is not a register (if it's an immediate or a memory reference), Free is a no-op.
func (a *Asm) Free(r Reg)

// Unfree reallocates a previously freed register r.
// If r is not a register (if it's an immediate or a memory reference), Unfree is a no-op.
// If r is not free for allocation, Unfree panics.
// A Free paired with Unfree can release a register for use temporarily
// but then reclaim it, such as at the end of a loop body when it must be restored.
func (a *Asm) Unfree(r Reg)

// A RegsUsed is a snapshot of which registers are allocated.
type RegsUsed struct {
	avail uint64
}

// RegsUsed returns a snapshot of which registers are currently allocated,
// which can be passed to a future call to [Asm.SetRegsUsed].
func (a *Asm) RegsUsed() RegsUsed

// SetRegsUsed sets which registers are currently allocated.
// The argument should have been returned from a previous
// call to [Asm.RegsUsed].
func (a *Asm) SetRegsUsed(used RegsUsed)

// FreeAll frees all known registers.
func (a *Asm) FreeAll()

// Printf emits to the assembly output.
func (a *Asm) Printf(format string, args ...any)

// Comment emits a line comment to the assembly output.
func (a *Asm) Comment(format string, args ...any)

// EOL appends an end-of-line comment to the previous line.
func (a *Asm) EOL(format string, args ...any)

// JmpEnable emits a test for the optional CPU feature that jumps to label if the feature is present.
// If JmpEnable returns false, the feature is not available on this architecture and no code was emitted.
func (a *Asm) JmpEnable(option Option, label string) bool

// Enabled reports whether the optional CPU feature is considered
// to be enabled at this point in the assembly output.
func (a *Asm) Enabled(option Option) bool

// SetOption changes whether the optional CPU feature should be
// considered to be enabled.
func (a *Asm) SetOption(option Option, on bool)

// Mov emits dst = src.
func (a *Asm) Mov(src, dst Reg)

// AddWords emits dst = src1*WordBytes + src2.
// It does not set or use the carry flag.
func (a *Asm) AddWords(src1 Reg, src2, dst RegPtr)

// And emits dst = src1 & src2
// It may modify the carry flag.
func (a *Asm) And(src1, src2, dst Reg)

// Or emits dst = src1 | src2
// It may modify the carry flag.
func (a *Asm) Or(src1, src2, dst Reg)

// Xor emits dst = src1 ^ src2
// It may modify the carry flag.
func (a *Asm) Xor(src1, src2, dst Reg)

// Neg emits dst = -src.
// It may modify the carry flag.
func (a *Asm) Neg(src, dst Reg)

// HasRegShift reports whether the architecture can use shift expressions as operands.
func (a *Asm) HasRegShift() bool

// LshReg returns a shift-expression operand src<<shift.
// If a.HasRegShift() == false, LshReg panics.
func (a *Asm) LshReg(shift, src Reg) Reg

// Lsh emits dst = src << shift.
// It may modify the carry flag.
func (a *Asm) Lsh(shift, src, dst Reg)

// LshWide emits dst = src << shift with low bits shifted from adj.
// It may modify the carry flag.
func (a *Asm) LshWide(shift, adj, src, dst Reg)

// RshReg returns a shift-expression operand src>>shift.
// If a.HasRegShift() == false, RshReg panics.
func (a *Asm) RshReg(shift, src Reg) Reg

// Rsh emits dst = src >> shift.
// It may modify the carry flag.
func (a *Asm) Rsh(shift, src, dst Reg)

// RshWide emits dst = src >> shift with high bits shifted from adj.
// It may modify the carry flag.
func (a *Asm) RshWide(shift, adj, src, dst Reg)

// SLTU emits dst = src2 < src1 (0 or 1), using an unsigned comparison.
func (a *Asm) SLTU(src1, src2, dst Reg)

// Add emits dst = src1+src2, with the specified carry behavior.
func (a *Asm) Add(src1, src2, dst Reg, carry Carry)

// Sub emits dst = src2-src1, with the specified carry behavior.
func (a *Asm) Sub(src1, src2, dst Reg, carry Carry)

// ClearCarry clears the carry flag.
// The ‘which’ parameter must be AddCarry or SubCarry to specify how the flag will be used.
// (On some systems, the sub carry's actual processor bit is inverted from its usual value.)
func (a *Asm) ClearCarry(which Carry)

// SaveCarry saves the carry flag into dst.
// The meaning of the bits in dst is architecture-dependent.
// The carry flag is left in an undefined state.
func (a *Asm) SaveCarry(dst Reg)

// RestoreCarry restores the carry flag from src.
// src is left in an undefined state.
func (a *Asm) RestoreCarry(src Reg)

// ConvertCarry converts the carry flag in dst from the internal format to a 0 or 1.
// The carry flag is left in an undefined state.
func (a *Asm) ConvertCarry(which Carry, dst Reg)

// SaveConvertCarry saves and converts the carry flag into dst: 0 unset, 1 set.
// The carry flag is left in an undefined state.
func (a *Asm) SaveConvertCarry(which Carry, dst Reg)

// MulWide emits dstlo = src1 * src2 and dsthi = (src1 * src2) >> WordBits.
// The carry flag is left in an undefined state.
// If dstlo or dsthi is the zero Reg, then those outputs are discarded.
func (a *Asm) MulWide(src1, src2, dstlo, dsthi Reg)

// Jmp jumps to the label.
func (a *Asm) Jmp(label string)

// JmpZero jumps to the label if src is zero.
// It may modify the carry flag unless a.Arch.CarrySafeLoop is true.
func (a *Asm) JmpZero(src Reg, label string)

// JmpNonZero jumps to the label if src is non-zero.
// It may modify the carry flag unless a.Arch,CarrySafeLoop is true.
func (a *Asm) JmpNonZero(src Reg, label string)

// Label emits a label with the given name.
func (a *Asm) Label(name string)

// Ret returns.
func (a *Asm) Ret()
