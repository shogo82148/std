// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/time"
)

// A Value can represent any Go value, but unlike type any,
// it can represent most small values without an allocation.
// The zero Value corresponds to nil.
type Value struct {
	_ [0]func()

	num uint64

	any any
}

// Kind is the kind of a Value.
type Kind int

const (
	KindAny Kind = iota
	KindBool
	KindDuration
	KindFloat64
	KindInt64
	KindString
	KindTime
	KindUint64
	KindGroup
	KindLogValuer
)

func (k Kind) String() string

// Unexported version of Kind, just so we can store Kinds in Values.
// (No user-provided value has this type.)

// Kind returns v's Kind.
func (v Value) Kind() Kind

// StringValue returns a new Value for a string.
func StringValue(value string) Value

// IntValue returns a Value for an int.
func IntValue(v int) Value

// Int64Value returns a Value for an int64.
func Int64Value(v int64) Value

// Uint64Value returns a Value for a uint64.
func Uint64Value(v uint64) Value

// Float64Value returns a Value for a floating-point number.
func Float64Value(v float64) Value

// BoolValue returns a Value for a bool.
func BoolValue(v bool) Value

// Unexported version of *time.Location, just so we can store *time.Locations in
// Values. (No user-provided value has this type.)

// TimeValue returns a Value for a time.Time.
// It discards the monotonic portion.
func TimeValue(v time.Time) Value

// DurationValue returns a Value for a time.Duration.
func DurationValue(v time.Duration) Value

// GroupValue returns a new Value for a list of Attrs.
// The caller must not subsequently mutate the argument slice.
func GroupValue(as ...Attr) Value

// AnyValue returns a Value for the supplied value.
//
// If the supplied value is of type Value, it is returned
// unmodified.
//
// Given a value of one of Go's predeclared string, bool, or
// (non-complex) numeric types, AnyValue returns a Value of kind
// String, Bool, Uint64, Int64, or Float64. The width of the
// original numeric type is not preserved.
//
// Given a time.Time or time.Duration value, AnyValue returns a Value of kind
// KindTime or KindDuration. The monotonic time is not preserved.
//
// For nil, or values of all other types, including named types whose
// underlying type is numeric, AnyValue returns a value of kind KindAny.
func AnyValue(v any) Value

// Any returns v's value as an any.
func (v Value) Any() any

// String returns Value's value as a string, formatted like fmt.Sprint. Unlike
// the methods Int64, Float64, and so on, which panic if v is of the
// wrong kind, String never panics.
func (v Value) String() string

// Int64 returns v's value as an int64. It panics
// if v is not a signed integer.
func (v Value) Int64() int64

// Uint64 returns v's value as a uint64. It panics
// if v is not an unsigned integer.
func (v Value) Uint64() uint64

// Bool returns v's value as a bool. It panics
// if v is not a bool.
func (v Value) Bool() bool

// Duration returns v's value as a time.Duration. It panics
// if v is not a time.Duration.
func (a Value) Duration() time.Duration

// Float64 returns v's value as a float64. It panics
// if v is not a float64.
func (v Value) Float64() float64

// Time returns v's value as a time.Time. It panics
// if v is not a time.Time.
func (v Value) Time() time.Time

// LogValuer returns v's value as a LogValuer. It panics
// if v is not a LogValuer.
func (v Value) LogValuer() LogValuer

// Group returns v's value as a []Attr.
// It panics if v's Kind is not KindGroup.
func (v Value) Group() []Attr

// Equal reports whether v and w represent the same Go value.
func (v Value) Equal(w Value) bool

// A LogValuer is any Go value that can convert itself into a Value for logging.
//
// This mechanism may be used to defer expensive operations until they are
// needed, or to expand a single value into a sequence of components.
type LogValuer interface {
	LogValue() Value
}

// Resolve repeatedly calls LogValue on v while it implements LogValuer,
// and returns the result.
// If v resolves to a group, the group's attributes' values are not recursively
// resolved.
// If the number of LogValue calls exceeds a threshold, a Value containing an
// error is returned.
// Resolve's return value is guaranteed not to be of Kind KindLogValuer.
func (v Value) Resolve() (rv Value)
