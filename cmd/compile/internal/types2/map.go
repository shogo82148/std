// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

// A Map represents a map type.
type Map struct {
	key, elem Type
}

// NewMap returns a new map for the given key and element types.
func NewMap(key, elem Type) *Map

// Key returns the key type of map m.
func (m *Map) Key() Type

// Elem returns the element type of map m.
func (m *Map) Elem() Type

func (t *Map) Underlying() Type
func (t *Map) String() string
