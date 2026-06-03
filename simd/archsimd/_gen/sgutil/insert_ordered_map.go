// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sgutil

import "github.com/shogo82148/std/iter"

type InsertMap[K comparable, V any] struct {
	m map[K]uint32
	v []V
	k []K
}

// Put inserts or updates a key-value pair in the map.
// If the key already exists, its value is updated while its insertion order
// remains unchanged. Returns the old value if there is one,
// and a boolean indicating whether an update occurred.
func (im *InsertMap[K, V]) Put(key K, val V) (old V, updated bool)

// Contains reports whether the map contains the specified key.
func (im *InsertMap[K, V]) Contains(key K) bool

// Get returns the value associated with the specified key.
// If the key does not exist, the zero value of type V is returned.
func (im *InsertMap[K, V]) Get(key K) V

// GetOk returns the value associated with the specified key and
// a boolean indicating whether the key exists in the map.
// If the key is not in the map, the zero value is returned along with false.
func (im *InsertMap[K, V]) GetOk(key K) (V, bool)

// Compare compares two keys based on their insertion order.
// It returns:
//   - -1 if 'a' was inserted before 'b'
//   - 1 if 'a' was inserted after 'b'
//   - 0 if 'a' and 'b' are equal, or if neither key exists in the map.
//
// If only one of the keys exists in the map, the existing key is considered
// to be "before" (less than) the non-existing key, returning -1 if 'a' exists,
// and 1 if 'b' exists.
func (im *InsertMap[K, V]) Compare(a, b K) int

// All returns an iterator over the key-value pairs of the map in insertion order.
func (im *InsertMap[K, V]) All() iter.Seq2[K, V]

// Keys returns an iterator over the keys of the map in insertion order.
func (im *InsertMap[K, V]) Keys() iter.Seq[K]

// Values returns an iterator over the values of the map in insertion order.
func (im *InsertMap[K, V]) Values() iter.Seq[V]
