// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

import (
	"github.com/shogo82148/std/unsafe"
)

// ValueKind is a tag for a metric [Value] which indicates its type.
type ValueKind int

const (
	// KindBad indicates that the Value has no type and should not be used.
	KindBad ValueKind = iota

	// KindUint64 indicates that the type of the Value is a uint64.
	KindUint64

	// KindFloat64 indicates that the type of the Value is a float64.
	KindFloat64

	// KindFloat64Histogram indicates that the type of the Value is a *Float64Histogram.
	KindFloat64Histogram
)

// Value represents a metric value returned by the runtime.
type Value struct {
	kind    ValueKind
	scalar  uint64
	pointer unsafe.Pointer
}

// Kind returns the tag representing the kind of value this is.
func (v Value) Kind() ValueKind

// Uint64 returns the internal uint64 value for the metric.
//
// If v.Kind() != KindUint64, this method panics.
func (v Value) Uint64() uint64

// Float64 returns the internal float64 value for the metric.
//
// If v.Kind() != KindFloat64, this method panics.
func (v Value) Float64() float64

// Float64Histogram returns the internal *Float64Histogram value for the metric.
//
// If v.Kind() != KindFloat64Histogram, this method panics.
func (v Value) Float64Histogram() *Float64Histogram
