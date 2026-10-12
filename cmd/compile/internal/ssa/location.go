// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

type LocPair [2]Location

// A LocalSlot is a location in the stack frame, which identifies and stores
// part or all of a PPARAM, PPARAMOUT, or PAUTO ONAME node.
// It can represent a whole variable, part of a larger stack slot, or part of a
// variable that has been decomposed into multiple stack slots.
// As an example, a string could have the following configurations:
//
//	          stack layout              LocalSlots
//
//	Optimizations are disabled. s is on the stack and represented in its entirety.
//	[ ------- s string ---- ] { N: s, Type: string, Off: 0 }
//
//	s was not decomposed, but the SSA operates on its parts individually, so
//	there is a LocalSlot for each of its fields that points into the single stack slot.
//	[ ------- s string ---- ] { N: s, Type: *uint8, Off: 0 }, {N: s, Type: int, Off: 8}
//
//	s was decomposed. Each of its fields is in its own stack slot and has its own LocalSLot.
//	[ ptr *uint8 ] [ len int] { N: ptr, Type: *uint8, Off: 0, SplitOf: parent, SplitOffset: 0},
//	                          { N: len, Type: int, Off: 0, SplitOf: parent, SplitOffset: 8}
//	                          parent = &{N: s, Type: string}
type LocalSlot struct {
	N    *ir.Name
	Type *types.Type
	Off  int64

	SplitOf     *LocalSlot
	SplitOffset int64
}

// A place that an ssa variable can reside.
type Location interface {
	String() string
}

type Spill struct {
	Type   *types.Type
	Offset int64
	Reg    int16
}

type LocResults []Location

func (s LocalSlot) String() string

func (t LocPair) String() string

func (t LocResults) String() string
