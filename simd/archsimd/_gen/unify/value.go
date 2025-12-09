// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unify

import (
	"github.com/shogo82148/std/iter"
)

// A Value represents a structured, non-deterministic value consisting of
// strings, tuples of Values, and string-keyed maps of Values. A
// non-deterministic Value will also contain variables, which are resolved via
// an environment as part of a [Closure].
//
// For debugging, a Value can also track the source position it was read from in
// an input file, and its provenance from other Values.
type Value struct {
	Domain Domain

	// A Value has either a pos or parents (or neither).
	pos     *Pos
	parents *[2]*Value
}

// NewValue returns a new [Value] with the given domain and no position
// information.
func NewValue(d Domain) *Value

// NewValuePos returns a new [Value] with the given domain at position p.
func NewValuePos(d Domain, p Pos) *Value

func (v *Value) Pos() Pos

func (v *Value) PosString() string

func (v *Value) WhyNotExact() string

func (v *Value) Exact() bool

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

// Provenance iterates over all of the source Values that have contributed to
// this Value.
func (v *Value) Provenance() iter.Seq[*Value]
