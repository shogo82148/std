// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package compare contains code for generating comparison
// routines for structs, strings and interfaces.
package compare

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

// IsRegularMemory reports whether t can be compared/hashed as regular memory.
func IsRegularMemory(t *types.Type) bool

// Memrun finds runs of struct fields for which memory-only algs are appropriate.
// t is the parent struct type, and start is the field index at which to start the run.
// size is the length in bytes of the memory included in the run.
// next is the index just after the end of the memory run.
func Memrun(t *types.Type, start int) (size int64, next int)

// EqCanPanic reports whether == on type t could panic (has an interface somewhere).
// t must be comparable.
func EqCanPanic(t *types.Type) bool

// EqStructCost returns the cost of an equality comparison of two structs.
//
// The cost is determined using an algorithm which takes into consideration
// the size of the registers in the current architecture and the size of the
// memory-only fields in the struct.
func EqStructCost(t *types.Type) int64

// EqStruct compares two structs np and nq for equality.
// It works by building a list of boolean conditions to satisfy.
// Conditions must be evaluated in the returned order and
// properly short-circuited by the caller.
// The first return value is the flattened list of conditions,
// the second value is a boolean indicating whether any of the
// comparisons could panic.
func EqStruct(t *types.Type, np, nq ir.Node) ([]ir.Node, bool)

// EqString returns the nodes
//
//	len(s) == len(t)
//
// and
//
//	memequal(s.ptr, t.ptr, len(s))
//
// which can be used to construct string equality comparison.
// eqlen must be evaluated before eqmem, and shortcircuiting is required.
func EqString(s, t ir.Node) (eqlen *ir.BinaryExpr, eqmem *ir.CallExpr)

// EqInterface returns the nodes
//
//	s.tab == t.tab (or s.typ == t.typ, as appropriate)
//
// and
//
//	ifaceeq(s.tab, s.data, t.data) (or efaceeq(s.typ, s.data, t.data), as appropriate)
//
// which can be used to construct interface equality comparison.
// eqtab must be evaluated before eqdata, and shortcircuiting is required.
func EqInterface(s, t ir.Node) (eqtab *ir.BinaryExpr, eqdata *ir.CallExpr)
