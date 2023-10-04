// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// A Union represents a union of terms embedded in an interface.
type Union struct {
	terms []*Term
}

// NewUnion returns a new Union type with the given terms.
// It is an error to create an empty union; they are syntactically not possible.
func NewUnion(terms []*Term) *Union

func (u *Union) Len() int
func (u *Union) Term(i int) *Term

func (u *Union) Underlying() Type
func (u *Union) String() string

// A Term represents a term in a Union.
type Term term

// NewTerm returns a new union term.
func NewTerm(tilde bool, typ Type) *Term

func (t *Term) Tilde() bool
func (t *Term) Type() Type
func (t *Term) String() string
