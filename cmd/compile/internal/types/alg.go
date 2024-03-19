// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// AlgKind describes the kind of algorithms used for comparing and
// hashing a Type.
type AlgKind int8

const (
	AUNK   AlgKind = iota
	ANOEQ
	ANOALG
	AMEM
	AMEM0
	AMEM8
	AMEM16
	AMEM32
	AMEM64
	AMEM128
	ASTRING
	AINTER
	ANILINTER
	AFLOAT32
	AFLOAT64
	ACPLX64
	ACPLX128
	ASPECIAL
)

// AlgType returns the AlgKind used for comparing and hashing Type t.
func AlgType(t *Type) AlgKind

// TypeHasNoAlg reports whether t does not have any associated hash/eq
// algorithms because t, or some component of t, is marked Noalg.
func TypeHasNoAlg(t *Type) bool

// IsComparable reports whether t is a comparable type.
func IsComparable(t *Type) bool

// IncomparableField returns an incomparable Field of struct Type t, if any.
func IncomparableField(t *Type) *Field

// IsPaddedField reports whether the i'th field of struct type t is followed
// by padding.
func IsPaddedField(t *Type, i int) bool
