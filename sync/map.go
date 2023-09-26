// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"github.com/shogo82148/std/sync/atomic"
)

// Map is a concurrent map with amortized-constant-time loads, stores, and deletes.
// It is safe for multiple goroutines to call a Map's methods concurrently.
//
// It is optimized for use in concurrent loops with keys that are
// stable over time, and either few steady-state stores, or stores
// localized to one goroutine per key.
//
// For use cases that do not share these attributes, it will likely have
// comparable or worse performance and worse type safety than an ordinary
// map paired with a read-write mutex.
//
// The zero Map is valid and empty.
//
// A Map must not be copied after first use.
type Map struct {
	mu Mutex

	read atomic.Value

	dirty map[interface{}]*entry

	misses int
}

// readOnly is an immutable struct stored atomically in the Map.read field.

// expunged is an arbitrary pointer that marks entries which have been deleted
// from the dirty map.

// An entry is a slot in the map corresponding to a particular key.

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *Map) Load(key interface{}) (value interface{}, ok bool)

// Store sets the value for a key.
func (m *Map) Store(key, value interface{})

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)

// Delete deletes the value for a key.
func (m *Map) Delete(key interface{})

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's
// contents: no key will be visited more than once, but if the value for any key
// is stored or deleted concurrently, Range may reflect any mapping for that key
// from any point during the Range call.
//
// Range may be O(N) with the number of elements in the map even if f returns
// false after a constant number of calls.
func (m *Map) Range(f func(key, value interface{}) bool)
