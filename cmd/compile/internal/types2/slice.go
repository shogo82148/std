// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

// A Slice represents a slice type.
type Slice struct {
	elem Type
}

// NewSlice returns a new slice type for the given element type.
func NewSlice(elem Type) *Slice

// Elem returns the element type of slice s.
func (s *Slice) Elem() Type

func (s *Slice) Underlying() Type
func (s *Slice) String() string
