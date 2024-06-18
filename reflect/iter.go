// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

import "github.com/shogo82148/std/iter"

// Seq returns an iter.Seq[Value] that loops over the elements of v.
// If v's kind is Func, it must be a function that has no results and
// that takes a single argument of type func(T) bool for some type T.
// If v's kind is Pointer, the pointer element type must have kind Array.
// Otherwise v's kind must be Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr,
// Array, Chan, Map, Slice, or String.
func (v Value) Seq() iter.Seq[Value]

// Seq2 returns an iter.Seq2[Value, Value] that loops over the elements of v.
// If v's kind is Func, it must be a function that has no results and
// that takes a single argument of type func(K, V) bool for some type K, V.
// If v's kind is Pointer, the pointer element type must have kind Array.
// Otherwise v's kind must be Array, Map, Slice, or String.
func (v Value) Seq2() iter.Seq2[Value, Value]
