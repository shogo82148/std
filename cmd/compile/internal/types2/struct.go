// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

// A Struct represents a struct type.
type Struct struct {
	fields []*Var
	tags   []string
}

// NewStruct returns a new struct with the given fields and corresponding field tags.
// If a field with index i has a tag, tags[i] must be that tag, but len(tags) may be
// only as long as required to hold the tag with the largest index i. Consequently,
// if no field has a tag, tags may be nil.
func NewStruct(fields []*Var, tags []string) *Struct

// NumFields returns the number of fields in the struct (including blank and embedded fields).
func (s *Struct) NumFields() int

// Field returns the i'th field for 0 <= i < NumFields().
func (s *Struct) Field(i int) *Var

// Tag returns the i'th field tag for 0 <= i < NumFields().
func (s *Struct) Tag(i int) string

func (s *Struct) Underlying() Type
func (s *Struct) String() string
