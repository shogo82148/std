// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// An Array represents an array type.
type Array struct {
	len  int64
	elem Type
}

// NewArray returns a new array type for the given element type and length.
// A negative length indicates an unknown length.
func NewArray(elem Type, len int64) *Array

// Len returns the length of array a.
// A negative result indicates an unknown length.
func (a *Array) Len() int64

// Elem returns element type of array a.
func (a *Array) Elem() Type

func (t *Array) Underlying() Type
func (t *Array) String() string
