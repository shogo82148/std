// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package obj

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

type Plist struct {
	Firstpc *Prog
	Curfn   Func
}

// ProgAlloc is a function that allocates Progs.
// It is used to provide access to cached/bulk-allocated Progs to the assemblers.
type ProgAlloc func() *Prog

func Flushplist(ctxt *Link, plist *Plist, newprog ProgAlloc)

func (ctxt *Link) InitTextSym(s *LSym, flag int, start src.XPos)

func (ctxt *Link) Globl(s *LSym, size int64, flag int)

func (ctxt *Link) GloblPos(s *LSym, size int64, flag int, pos src.XPos)

// EmitEntryStackMap generates PCDATA Progs after p to switch to the
// liveness map active at the entry of function s.
func (ctxt *Link) EmitEntryStackMap(s *LSym, p *Prog, newprog ProgAlloc) *Prog

// StartUnsafePoint generates PCDATA Progs after p to mark the
// beginning of an unsafe point. The unsafe point starts immediately
// after p.
// It returns the last Prog generated.
func (ctxt *Link) StartUnsafePoint(p *Prog, newprog ProgAlloc) *Prog

// EndUnsafePoint generates PCDATA Progs after p to mark the end of an
// unsafe point, restoring the register map index to oldval.
// The unsafe point ends right after p.
// It returns the last Prog generated.
func (ctxt *Link) EndUnsafePoint(p *Prog, newprog ProgAlloc, oldval int64) *Prog

// MarkUnsafePoints inserts PCDATAs to mark nonpreemptible and restartable
// instruction sequences, based on isUnsafePoint and isRestartable predicate.
// p0 is the start of the instruction stream.
// isUnsafePoint(p) returns true if p is not safe for async preemption.
// isRestartable(p) returns true if we can restart at the start of p (this Prog)
// upon async preemption. (Currently multi-Prog restartable sequence is not
// supported.)
// isRestartable can be nil. In this case it is treated as always returning false.
// If isUnsafePoint(p) and isRestartable(p) are both true, it is treated as
// an unsafe point.
func MarkUnsafePoints(ctxt *Link, p0 *Prog, newprog ProgAlloc, isUnsafePoint, isRestartable func(*Prog) bool)
