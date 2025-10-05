// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Garbage collector liveness bitmap generation.

// The command line flag -live causes this code to print debug information.
// The levels are:
//
//	-live (aka -live=1): print liveness lists as code warnings at safe points
//	-live=2: print an assembly listing with liveness annotations
//
// Each level includes the earlier output as well.

package liveness

import (
	"github.com/shogo82148/std/cmd/compile/internal/abi"
	"github.com/shogo82148/std/cmd/compile/internal/bitvec"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/objw"
	"github.com/shogo82148/std/cmd/compile/internal/ssa"
)

// A collection of global state used by Liveness analysis.
type Liveness struct {
	fn         *ir.Func
	f          *ssa.Func
	vars       []*ir.Name
	idx        map[*ir.Name]int32
	stkptrsize int64

	be []blockEffects

	// allUnsafe indicates that all points in this function are
	// unsafe-points.
	allUnsafe bool
	// unsafePoints bit i is set if Value ID i is an unsafe-point
	// (preemption is not allowed). Only valid if !allUnsafe.
	unsafePoints bitvec.BitVec
	// unsafeBlocks bit i is set if Block ID i is an unsafe-point
	// (preemption is not allowed on any end-of-block
	// instructions). Only valid if !allUnsafe.
	unsafeBlocks bitvec.BitVec

	// An array with a bit vector for each safe point in the
	// current Block during liveness.epilogue. Indexed in Value
	// order for that block. Additionally, for the entry block
	// livevars[0] is the entry bitmap. liveness.compact moves
	// these to stackMaps.
	livevars []bitvec.BitVec

	// livenessMap maps from safe points (i.e., CALLs) to their
	// liveness map indexes.
	livenessMap Map
	stackMapSet bvecSet
	stackMaps   []bitvec.BitVec

	cache progeffectscache

	// partLiveArgs includes input arguments (PPARAM) that may
	// be partially live. That is, it is considered live because
	// a part of it is used, but we may not initialize all parts.
	partLiveArgs map[*ir.Name]bool

	doClobber     bool
	noClobberArgs bool

	// treat "dead" writes as equivalent to reads during the analysis;
	// used only during liveness analysis for stack slot merging (doesn't
	// make sense for stackmap analysis).
	conservativeWrites bool
}

// Map maps from *ssa.Value to StackMapIndex.
// Also keeps track of unsafe ssa.Values and ssa.Blocks.
// (unsafe = can't be interrupted during GC.)
type Map struct {
	Vals         map[ssa.ID]objw.StackMapIndex
	UnsafeVals   map[ssa.ID]bool
	UnsafeBlocks map[ssa.ID]bool
	// The set of live, pointer-containing variables at the DeferReturn
	// call (only set when open-coded defers are used).
	DeferReturn objw.StackMapIndex
}

func (m Map) Get(v *ssa.Value) objw.StackMapIndex

func (m Map) GetUnsafe(v *ssa.Value) bool

func (m Map) GetUnsafeBlock(b *ssa.Block) bool

// IsUnsafe indicates that all points in this function are
// unsafe-points.
func IsUnsafe(f *ssa.Func) bool

func (lv *Liveness) Format(v *ssa.Value) string

// Entry pointer for Compute analysis. Solves for the Compute of
// pointer variables in the function and emits a runtime data
// structure read by the garbage collector.
// Returns a map from GC safe points to their corresponding stack map index,
// and a map that contains all input parameters that may be partially live.
func Compute(curfn *ir.Func, f *ssa.Func, stkptrsize int64, pp *objw.Progs, retLiveness bool) (Map, map[*ir.Name]bool, *Liveness)

// WriteFuncMap writes the pointer bitmaps for bodyless function fn's
// inputs and outputs as the value of symbol <fn>.args_stackmap.
// If fn has outputs, two bitmaps are written, otherwise just one.
func WriteFuncMap(fn *ir.Func, abiInfo *abi.ABIParamResultInfo)
