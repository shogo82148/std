// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// A Pointer represents a pointer type.
type Pointer struct {
	base Type
}

// NewPointer returns a new pointer type for the given element (base) type.
func NewPointer(elem Type) *Pointer

// Elem returns the element type for the given pointer p.
func (p *Pointer) Elem() Type

func (t *Pointer) Underlying() Type
func (t *Pointer) String() string
