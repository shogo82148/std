// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maps

import "github.com/shogo82148/std/iter"

// All returns an iterator over key-value pairs from m.
// The iteration order is not specified and is not guaranteed
// to be the same from one call to the next.
func All[Map ~map[K]V, K comparable, V any](m Map) iter.Seq2[K, V]

// Keys returns an iterator over keys in m.
// The iteration order is not specified and is not guaranteed
// to be the same from one call to the next.
func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K]

// Values returns an iterator over values in m.
// The iteration order is not specified and is not guaranteed
// to be the same from one call to the next.
func Values[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[V]

// Insert adds the key-value pairs from seq to m.
// If a key in seq already exists in m, its value will be overwritten.
func Insert[Map ~map[K]V, K comparable, V any](m Map, seq iter.Seq2[K, V])

// Collect collects key-value pairs from seq into a new map
// and returns it.
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V
