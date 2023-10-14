// Derived from Inferno utils/6c/txt.c
// https://bitbucket.org/inferno-os/inferno-os/src/master/utils/6c/txt.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package objw

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/src"
)

// NewProgs returns a new Progs for fn.
// worker indicates which of the backend workers will use the Progs.
func NewProgs(fn *ir.Func, worker int) *Progs

// Progs accumulates Progs for a function and converts them into machine code.
type Progs struct {
	Text       *obj.Prog
	Next       *obj.Prog
	PC         int64
	Pos        src.XPos
	CurFunc    *ir.Func
	Cache      []obj.Prog
	CacheIndex int

	NextLive LivenessIndex
	PrevLive LivenessIndex
}

// LivenessIndex stores the liveness map information for a Value.
type LivenessIndex struct {
	StackMapIndex int

	// IsUnsafePoint indicates that this is an unsafe-point.
	//
	// Note that it's possible for a call Value to have a stack
	// map while also being an unsafe-point. This means it cannot
	// be preempted at this instruction, but that a preemption or
	// stack growth may happen in the called function.
	IsUnsafePoint bool
}

// StackMapDontCare indicates that the stack map index at a Value
// doesn't matter.
//
// This is a sentinel value that should never be emitted to the PCDATA
// stream. We use -1000 because that's obviously never a valid stack
// index (but -1 is).
const StackMapDontCare = -1000

// LivenessDontCare indicates that the liveness information doesn't
// matter. Currently it is used in deferreturn liveness when we don't
// actually need it. It should never be emitted to the PCDATA stream.
var LivenessDontCare = LivenessIndex{StackMapDontCare, true}

func (idx LivenessIndex) StackMapValid() bool

func (pp *Progs) NewProg() *obj.Prog

// Flush converts from pp to machine code.
func (pp *Progs) Flush()

// Free clears pp and any associated resources.
func (pp *Progs) Free()

// Prog adds a Prog with instruction As to pp.
func (pp *Progs) Prog(as obj.As) *obj.Prog

func (pp *Progs) Clear(p *obj.Prog)

func (pp *Progs) Append(p *obj.Prog, as obj.As, ftype obj.AddrType, freg int16, foffset int64, ttype obj.AddrType, treg int16, toffset int64) *obj.Prog

func (pp *Progs) SetText(fn *ir.Func)

// LosesStmtMark reports whether a prog with op as loses its statement mark on the way to DWARF.
// The attributes from some opcodes are lost in translation.
// TODO: this is an artifact of how funcpctab combines information for instructions at a single PC.
// Should try to fix it there.
func LosesStmtMark(as obj.As) bool
