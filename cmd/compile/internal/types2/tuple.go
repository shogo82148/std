// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

// A Tuple represents an ordered list of variables; a nil *Tuple is a valid (empty) tuple.
// Tuples are used as components of signatures and to represent the type of multiple
// assignments; they are not first class types of Go.
type Tuple struct {
	vars []*Var
}

// NewTuple returns a new tuple for the given variables.
func NewTuple(x ...*Var) *Tuple

// Len returns the number variables of tuple t.
func (t *Tuple) Len() int

// At returns the i'th variable of tuple t.
func (t *Tuple) At(i int) *Var

func (t *Tuple) Underlying() Type
func (t *Tuple) String() string
