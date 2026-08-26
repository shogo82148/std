// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unify

// Decode decodes v into a Go value.
//
// v must be exact, except that it can include Top. into must be a pointer.
// [Def]s are decoded into structs. [Tuple]s are decoded into slices. [String]s
// are decoded into strings or ints. Any field can itself be a pointer to one of
// these types. Top can be decoded into a pointer-typed field and will set the
// field to nil. Anything else will allocate a value if necessary.
//
// Any type may implement [Decoder], in which case its DecodeUnified method will
// be called instead of using the default decoding scheme.
func (v *Value) Decode(into any) error

// Decoder can be implemented by types as a custom implementation of [Decode]
// for that type.
type Decoder interface {
	DecodeUnified(v *Value) error
}
