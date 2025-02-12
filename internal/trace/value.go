// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/unsafe"
)

// Value is a dynamically-typed value obtained from a trace.
type Value struct {
	kind    ValueKind
	pointer unsafe.Pointer
	scalar  uint64
}

// ValueKind is the type of a dynamically-typed value from a trace.
type ValueKind uint8

const (
	ValueBad ValueKind = iota
	ValueUint64
	ValueString
)

// Kind returns the ValueKind of the value.
//
// It represents the underlying structure of the value.
//
// New ValueKinds may be added in the future. Users of this type must be robust
// to that possibility.
func (v Value) Kind() ValueKind

// ToUint64 returns the uint64 value for a ValueUint64.
//
// Panics if this Value's Kind is not ValueUint64.
func (v Value) ToUint64() uint64

// ToString returns the uint64 value for a ValueString.
//
// Panics if this Value's Kind is not ValueString.
func (v Value) ToString() string

// String returns the string representation of the value.
func (v Value) String() string
