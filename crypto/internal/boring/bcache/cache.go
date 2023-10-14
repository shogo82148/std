// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bcache implements a GC-friendly cache (see [Cache]) for BoringCrypto.
package bcache

import (
	"github.com/shogo82148/std/sync/atomic"
)

// A Cache is a GC-friendly concurrent map from unsafe.Pointer to
// unsafe.Pointer. It is meant to be used for maintaining shadow
// BoringCrypto state associated with certain allocated structs, in
// particular public and private RSA and ECDSA keys.
//
// The cache is GC-friendly in the sense that the keys do not
// indefinitely prevent the garbage collector from collecting them.
// Instead, at the start of each GC, the cache is cleared entirely. That
// is, the cache is lossy, and the loss happens at the start of each GC.
// This means that clients need to be able to cope with cache entries
// disappearing, but it also means that clients don't need to worry about
// cache entries keeping the keys from being collected.
type Cache[K, V any] struct {
	// The runtime atomically stores nil to ptable at the start of each GC.
	ptable atomic.Pointer[cacheTable[K, V]]
}

// Register registers the cache with the runtime,
// so that c.ptable can be cleared at the start of each GC.
// Register must be called during package initialization.
func (c *Cache[K, V]) Register()

// Clear clears the cache.
// The runtime does this automatically at each garbage collection;
// this method is exposed only for testing.
func (c *Cache[K, V]) Clear()

// Get returns the cached value associated with v,
// which is either the value v corresponding to the most recent call to Put(k, v)
// or nil if that cache entry has been dropped.
func (c *Cache[K, V]) Get(k *K) *V

// Put sets the cached value associated with k to v.
func (c *Cache[K, V]) Put(k *K, v *V)
