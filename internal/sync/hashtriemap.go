// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"github.com/shogo82148/std/sync/atomic"
)

// HashTrieMap is an implementation of a concurrent hash-trie. The implementation
// is designed around frequent loads, but offers decent performance for stores
// and deletes as well, especially if the map is larger. Its primary use-case is
// the unique package, but can be used elsewhere as well.
//
// The zero HashTrieMap is empty and ready to use.
// It must not be copied after first use.
type HashTrieMap[K comparable, V any] struct {
	inited   atomic.Uint32
	initMu   Mutex
	root     atomic.Pointer[indirect[K, V]]
	keyHash  hashFunc
	valEqual equalFunc
	seed     uintptr
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (ht *HashTrieMap[K, V]) Load(key K) (value V, ok bool)

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (ht *HashTrieMap[K, V]) LoadOrStore(key K, value V) (result V, loaded bool)

// Store sets the value for a key.
func (ht *HashTrieMap[K, V]) Store(key K, new V)

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (ht *HashTrieMap[K, V]) Swap(key K, new V) (previous V, loaded bool)

// CompareAndSwap swaps the old and new values for key
// if the value stored in the map is equal to old.
// The value type must be of a comparable type, otherwise CompareAndSwap will panic.
func (ht *HashTrieMap[K, V]) CompareAndSwap(key K, old, new V) (swapped bool)

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (ht *HashTrieMap[K, V]) LoadAndDelete(key K) (value V, loaded bool)

// Delete deletes the value for a key.
func (ht *HashTrieMap[K, V]) Delete(key K)

// CompareAndDelete deletes the entry for key if its value is equal to old.
// The value type must be comparable, otherwise this CompareAndDelete will panic.
//
// If there is no current value for key in the map, CompareAndDelete returns false
// (even if the old value is the nil interface value).
func (ht *HashTrieMap[K, V]) CompareAndDelete(key K, old V) (deleted bool)

// All returns an iterator over each key and value present in the map.
//
// The iterator does not necessarily correspond to any consistent snapshot of the
// HashTrieMap's contents: no key will be visited more than once, but if the value
// for any key is stored or deleted concurrently (including by yield), the iterator
// may reflect any mapping for that key from any point during iteration. The iterator
// does not block other methods on the receiver; even yield itself may call any
// method on the HashTrieMap.
func (ht *HashTrieMap[K, V]) All() func(yield func(K, V) bool)

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.
//
// This exists for compatibility with sync.Map; All should be preferred.
// It provides the same guarantees as sync.Map, and All.
func (ht *HashTrieMap[K, V]) Range(yield func(K, V) bool)

// Clear deletes all the entries, resulting in an empty HashTrieMap.
func (ht *HashTrieMap[K, V]) Clear()
