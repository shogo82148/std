// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unify

import (
	"github.com/shogo82148/std/iter"
	"github.com/shogo82148/std/reflect"
	"github.com/shogo82148/std/regexp"
)

// A Domain is a non-empty set of values, all of the same kind.
//
// Domain may be a scalar:
//
//   - [String] - Represents string-typed values.
//
// Or a composite:
//
//   - [Def] - A mapping from fixed keys to [Domain]s.
//
//   - [Tuple] - A fixed-length sequence of [Domain]s or
//     all possible lengths repeating a [Domain].
//
// Or top or bottom:
//
//   - [Top] - Represents all possible values of all kinds.
//
//   - nil - Represents no values.
//
// Or a variable:
//
//   - [Var] - A value captured in the environment.
type Domain interface {
	Exact() bool
	WhyNotExact() string

	decode(reflect.Value) error
}

// Top represents all possible values of all possible types.
type Top struct{}

func (t Top) Exact() bool
func (t Top) WhyNotExact() string

// A Def is a mapping from field names to [Value]s. Any fields not explicitly
// listed have [Value] [Top].
type Def struct {
	fields map[string]*Value
}

// A DefBuilder builds a [Def] one field at a time. The zero value is an empty
// [Def].
type DefBuilder struct {
	fields map[string]*Value
}

func (b *DefBuilder) Add(name string, v *Value)

// Build constructs a [Def] from the fields added to this builder.
func (b *DefBuilder) Build() Def

// Exact returns true if all field Values are exact.
func (d Def) Exact() bool

// WhyNotExact returns why the value is not exact
func (d Def) WhyNotExact() string

func (d Def) All() iter.Seq2[string, *Value]

// A Tuple is a sequence of Values in one of two forms: 1. a fixed-length tuple,
// where each Value can be different or 2. a "repeated tuple", which is a Value
// repeated 0 or more times.
type Tuple struct {
	vs []*Value

	// repeat, if non-nil, means this Tuple consists of an element repeated 0 or
	// more times. If repeat is non-nil, vs must be nil. This is a generator
	// function because we don't necessarily want *exactly* the same Value
	// repeated. For example, in YAML encoding, a !sum in a repeated tuple needs
	// a fresh variable in each instance.
	repeat []func(envSet) (*Value, envSet)
}

func NewTuple(vs ...*Value) Tuple

func NewRepeat(gens ...func(envSet) (*Value, envSet)) Tuple

func (d Tuple) Exact() bool

func (d Tuple) WhyNotExact() string

// A String represents a set of strings. It can represent the intersection of a
// set of regexps, or a single exact string. In general, the domain of a String
// is non-empty, but we do not attempt to prove emptiness of a regexp value.
type String struct {
	kind  stringKind
	re    []*regexp.Regexp
	exact string
}

func NewStringRegex(exprs ...string) (String, error)

func NewStringExact(s string) String

// Exact returns whether this Value is known to consist of a single string.
func (d String) Exact() bool

func (d String) WhyNotExact() string
