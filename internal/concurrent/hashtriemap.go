// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package concurrent

// HashTrieMap is an implementation of a concurrent hash-trie. The implementation
// is designed around frequent loads, but offers decent performance for stores
// and deletes as well, especially if the map is larger. It's primary use-case is
// the unique package, but can be used elsewhere as well.
type HashTrieMap[K, V comparable] struct {
	root     *indirect[K, V]
	keyHash  hashFunc
	keyEqual equalFunc
	valEqual equalFunc
	seed     uintptr
}

// NewHashTrieMap creates a new HashTrieMap for the provided key and value.
func NewHashTrieMap[K, V comparable]() *HashTrieMap[K, V]

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (ht *HashTrieMap[K, V]) Load(key K) (value V, ok bool)

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (ht *HashTrieMap[K, V]) LoadOrStore(key K, value V) (result V, loaded bool)

// CompareAndDelete deletes the entry for key if its value is equal to old.
//
// If there is no current value for key in the map, CompareAndDelete returns false
// (even if the old value is the nil interface value).
func (ht *HashTrieMap[K, V]) CompareAndDelete(key K, old V) (deleted bool)

// Enumerate produces all key-value pairs in the map. The enumeration does
// not represent any consistent snapshot of the map, but is guaranteed
// to visit each unique key-value pair only once. It is safe to operate
// on the tree during iteration. No particular enumeration order is
// guaranteed.
func (ht *HashTrieMap[K, V]) Enumerate(yield func(key K, value V) bool)
