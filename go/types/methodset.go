// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements method sets.

package types

// A MethodSet is an ordered set of concrete or abstract (interface) methods;
// a method is a [MethodVal] selection, and they are ordered by ascending m.Obj().Id().
// The zero value for a MethodSet is a ready-to-use empty method set.
type MethodSet struct {
	list []*Selection
}

func (s *MethodSet) String() string

// Len returns the number of methods in s.
func (s *MethodSet) Len() int

// At returns the i'th method in s for 0 <= i < s.Len().
func (s *MethodSet) At(i int) *Selection

// Lookup returns the method with matching package and name, or nil if not found.
func (s *MethodSet) Lookup(pkg *Package, name string) *Selection

// NewMethodSet returns the method set for the given type T.
// It always returns a non-nil method set, even if it is empty.
func NewMethodSet(T Type) *MethodSet
