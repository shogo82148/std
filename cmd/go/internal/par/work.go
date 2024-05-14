// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package par implements parallel execution helpers.
package par

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/sync"
)

// Work manages a set of work items to be executed in parallel, at most once each.
// The items in the set must all be valid map keys.
type Work[T comparable] struct {
	f       func(T)
	running int

	mu      sync.Mutex
	added   map[T]bool
	todo    []T
	wait    sync.Cond
	waiting int
}

// Add adds item to the work set, if it hasn't already been added.
func (w *Work[T]) Add(item T)

// Do runs f in parallel on items from the work set,
// with at most n invocations of f running at a time.
// It returns when everything added to the work set has been processed.
// At least one item should have been added to the work set
// before calling Do (or else Do returns immediately),
// but it is allowed for f(item) to add new items to the set.
// Do should only be used once on a given Work.
func (w *Work[T]) Do(n int, f func(item T))

// ErrCache is like Cache except that it also stores
// an error value alongside the cached value V.
type ErrCache[K comparable, V any] struct {
	Cache[K, errValue[V]]
}

func (c *ErrCache[K, V]) Do(key K, f func() (V, error)) (V, error)

var ErrCacheEntryNotFound = errors.New("cache entry not found")

// Get returns the cached result associated with key.
// It returns ErrCacheEntryNotFound if there is no such result.
func (c *ErrCache[K, V]) Get(key K) (V, error)

// Cache runs an action once per key and caches the result.
type Cache[K comparable, V any] struct {
	m sync.Map
}

// Do calls the function f if and only if Do is being called for the first time with this key.
// No call to Do with a given key returns until the one call to f returns.
// Do returns the value returned by the one call to f.
func (c *Cache[K, V]) Do(key K, f func() V) V

// Get returns the cached result associated with key
// and reports whether there is such a result.
//
// If the result for key is being computed, Get does not wait for the computation to finish.
func (c *Cache[K, V]) Get(key K) (V, bool)
