// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

import (
	"github.com/shogo82148/std/unsafe"
)

<<<<<<< HEAD
// ValueKindはメトリックの値のタグで、そのタイプを示します。
=======
// ValueKind is a tag for a metric [Value] which indicates its type.
>>>>>>> upstream/master
type ValueKind int

const (
	// KindBadは、Valueに型がなく、使用すべきではないことを示します。
	KindBad ValueKind = iota

	// KindUint64 は Value の型が uint64 であることを示します。
	KindUint64

	// KindFloat64はValueの型がfloat64であることを示します。
	KindFloat64

	// KindFloat64HistogramはValueの型が*Float64Histogramであることを示します。
	KindFloat64Histogram
)

// Valueはランタイムが返すメトリック値を表します。
type Value struct {
	kind    ValueKind
	scalar  uint64
	pointer unsafe.Pointer
}

// Kindは、この値の種類を表すタグを返します。
func (v Value) Kind() ValueKind

// Uint64はメトリックの内部uint64値を返します。
//
// もしv.Kind() != KindUint64の場合、このメソッドはパニックします。
func (v Value) Uint64() uint64

// Float64はメトリックの内部float64値を返します。
//
// もしv.Kind() != KindFloat64なら、このメソッドはパニックを引き起こします。
func (v Value) Float64() float64

func (v Value) Float64Histogram() *Float64Histogram
