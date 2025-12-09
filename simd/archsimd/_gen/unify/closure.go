// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unify

import (
	"github.com/shogo82148/std/iter"
)

type Closure struct {
	val *Value
	env envSet
}

func NewSum(vs ...*Value) Closure

// IsBottom returns whether c consists of no values.
func (c Closure) IsBottom() bool

// Summands returns the top-level Values of c. This assumes the top-level of c
// was constructed as a sum, and is mostly useful for debugging.
func (c Closure) Summands() iter.Seq[*Value]

// All enumerates all possible concrete values of c by substituting variables
// from the environment.
//
// E.g., enumerating this Value
//
//	a: !sum [1, 2]
//	b: !sum [3, 4]
//
// results in
//
//   - {a: 1, b: 3}
//   - {a: 1, b: 4}
//   - {a: 2, b: 3}
//   - {a: 2, b: 4}
func (c Closure) All() iter.Seq[*Value]
