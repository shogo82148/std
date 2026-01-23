// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package staticinit

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

type Entry struct {
	Xoffset int64
	Expr    ir.Node
}

type Plan struct {
	E []Entry
}

// An Schedule is used to decompose assignment statements into
// static and dynamic initialization parts. Static initializations are
// handled by populating variables' linker symbol data, while dynamic
// initializations are accumulated to be executed in order.
type Schedule struct {
	// Out is the ordered list of dynamic initialization
	// statements.
	Out []ir.Node

	Plans map[ir.Node]*Plan
	Temps map[ir.Node]*ir.Name

	// seenMutation tracks whether we've seen an initialization
	// expression that may have modified other package-scope variables
	// within this package.
	seenMutation bool
}

// StaticInit adds an initialization statement n to the schedule.
func (s *Schedule) StaticInit(n ir.Node)

// MapInitToVar is the inverse of VarToMapInit; it maintains a mapping
// from a compiler-generated init function to the map the function is
// initializing.
var MapInitToVar map[*ir.Func]*ir.Name

func (s *Schedule) StaticAssign(l *ir.Name, loff int64, r ir.Node, typ *types.Type) bool

// StaticName returns a name backed by a (writable) static data symbol.
func StaticName(t *types.Type) *ir.Name

// StaticLoc returns the static address of n, if n has one, or else nil.
func StaticLoc(n ir.Node) (name *ir.Name, offset int64, ok bool)

// AnySideEffects reports whether n contains any operations that could have observable side effects.
func AnySideEffects(n ir.Node) bool

// AddKeepRelocations adds a dummy "R_KEEP" relocation from each
// global map variable V to its associated outlined init function.
// These relocation ensure that if the map var itself is determined to
// be reachable at link time, we also mark the init function as
// reachable.
func AddKeepRelocations()

// OutlineMapInits replaces global map initializers with outlined
// calls to separate "map init" functions (where possible and
// profitable), to facilitate better dead-code elimination by the
// linker.
func OutlineMapInits(fn *ir.Func)

// SplitLargeInit breaks up a large "init" function into smaller chunks to avoid slow compilation.
func SplitLargeInit(fn *ir.Func)

// CanOptimize reports whether the given fn can be optimized for static assignments.
func CanOptimize(fn *ir.Func) bool
