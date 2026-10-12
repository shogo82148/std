// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package specexpr manipulates symbolic constraints on vector shapes.
//
// # Shapes
//
// A vector shape consists of a base type B (e.g., int, float), element width N
// (8, 16, 32, or 64), and a vector width W (128, 256, 512, or scalable). It
// also has a lane count L, which is the vector width / element width. These are
// written like "Int32x4" or "Int32w128" or, for a scalable vector, "Int32s"
// Masks are represented as vectors with a base type of "Mask", e.g.,
// "Mask32x4".
//
// Scalar shapes consist only of a base type and an element width, e.g.,
// "uint32".
//
// # Expressions
//
// Constraints are written as boolean expressions over shapes and a few basic
// types.
//
// The primary expressions are:
//
//   - Variable ([a-zA-Z]+), such as x or xL
//
//   - Integers ([0-9]+)
//
//   - Shapes, written in the form given above, but where each component can be
//     written as a bracketed expression, such as "Int32x{z/2}".
//
// These can be combined with operators * or /, or comparison operators =, >, <,
// >=, <=. Arithmetic operators bind more tightly than comparison operators.
//
// # Example
//
// Consider a DotProductPairs function that takes two vectors x and y that have
// the same shape and produces a vector z that has the same base type as x and
// y, but has half as many elements, each of double the width. This can be
// expressed as:
//
//	y=x
//	z={xB}{xN*2}x{xL/2}
//
// # Width rounding
//
// The minimum vector width is 128 bits. Sometimes, operations would naturally
// produce a width smaller than this, so hardware simply pads the vector out to
// 128 bits. Shapes implement this behavior. For example, consider a "convert to
// float32 operation" with constraints
//
//	z=Float32x{xL}
//
// If x is Float64x2, then z would naturally be Float32x2, but since this is
// only 64 bits, the shape is "rounded" up to Float32x4.
//
// # Limitations
//
// The solver is intentionally simple. See [Solver] for a description of its
// limitations. If you run up against its limitations, you're probably being too
// clever.
package specexpr

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/iter"
)

// A Solver solves a set of constraints.
//
// This is a simple monotonic solver. It looks for a single order in which it
// can resolve all constraints, assertions of the form "var=expr" are treated as
// candidates for resolving the value of "var", and anything else is treated as
// a boolean check. Variables that appear only on the right hand side of
// assignments are "independent" and it will enumerate all possible values of
// these variables. It never tries to invert any formulas and refuses to solve a
// system with cycles. This is intentional to keep this solver fast: if your
// formulas are cyclic, you're doing something too complicated.
type Solver struct {
	vars    map[Variable][]any
	asserts []Expr
	tracer  *tracer
}

// SetTrace enables emitting a solver trace to w.
func (s *Solver) SetTrace(w io.Writer)

// Declare declares a variable and its domain.
//
// Any "int" values in domain will be converted to [Int].
func (s *Solver) Declare(v Variable, domain []any)

// Assign is a convenience for asserting that v=val.
func (s *Solver) Assign(v Variable, val Expr) Variable

// Assert asserts a boolean condition must be true.
func (s *Solver) Assert(cond Expr)

func (s *Solver) Fprint(w io.Writer)

// Bindings is a set of variable values.
type Bindings struct {
	varNames map[Variable]int
	vals     []any
}

// Get returns the value of v if resolved, or nil.
func (b *Bindings) Get(v Variable) any

// All yields all variable bindings.
func (b *Bindings) All() iter.Seq2[Variable, any]

func (b *Bindings) String() string

// Solve yields all satisfying assignments of the variables in s.
func (s *Solver) Solve() iter.Seq2[*Bindings, error]
