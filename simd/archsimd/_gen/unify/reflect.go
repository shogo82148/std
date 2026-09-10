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
// Any type may implement [DecoderEncoder], in which case its DecodeUnified
// method will be called instead of using the default decoding scheme.
func (v *Value) Decode(into any) error

// Encode constructs a Value from a Go value. It is the inverse of
// [Value.Decode].
//
// If a struct has an "Encode<Field>" field, it will be used for encoding the
// field named by <Field>, overriding the default field. This behavior is useful
// when a single type is used as an input and an output from unification, but
// the input and output have different requirements (e.g., optionality or
// strings vs regexps).
func Encode(gv any) *Value

// DecoderEncoder can be implemented by types as a custom implementation of
// [Decode] and [Encode] for that type. These must be inverses.
type DecoderEncoder interface {
	DecodeUnified(v *Value) error
	EncodeUnified() *Value
}
