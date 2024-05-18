// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

// Value is a dynamically-typed value obtained from a trace.
type Value struct {
	kind   ValueKind
	scalar uint64
}

// ValueKind is the type of a dynamically-typed value from a trace.
type ValueKind uint8

const (
	ValueBad ValueKind = iota
	ValueUint64
)

// Kind returns the ValueKind of the value.
//
// It represents the underlying structure of the value.
//
// New ValueKinds may be added in the future. Users of this type must be robust
// to that possibility.
func (v Value) Kind() ValueKind

// Uint64 returns the uint64 value for a MetricSampleUint64.
//
// Panics if this metric sample's Kind is not MetricSampleUint64.
func (v Value) Uint64() uint64
