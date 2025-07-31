// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package maps implements Go's builtin map type.
package maps

import (
	"github.com/shogo82148/std/internal/abi"
	"github.com/shogo82148/std/unsafe"
)

// Ensure the max capacity fits in uint16, used for capacity and growthLeft
// below.
var _ = uint16(maxTableCapacity)

type Iter struct {
	key  unsafe.Pointer
	elem unsafe.Pointer
	typ  *abi.MapType
	m    *Map

	// Randomize iteration order by starting iteration at a random slot
	// offset. The offset into the directory uses a separate offset, as it
	// must adjust when the directory grows.
	entryOffset uint64
	dirOffset   uint64

	// Snapshot of Map.clearSeq at iteration initialization time. Used to
	// detect clear during iteration.
	clearSeq uint64

	// Value of Map.globalDepth during the last call to Next. Used to
	// detect directory grow during iteration.
	globalDepth uint8

	// dirIdx is the current directory index, prior to adjustment by
	// dirOffset.
	dirIdx int

	// tab is the table at dirIdx during the previous call to Next.
	tab *table

	// group is the group at entryIdx during the previous call to Next.
	group groupReference

	// entryIdx is the current entry index, prior to adjustment by entryOffset.
	// The lower 3 bits of the index are the slot index, and the upper bits
	// are the group index.
	entryIdx uint64
}

// Init initializes Iter for iteration.
func (it *Iter) Init(typ *abi.MapType, m *Map)

func (it *Iter) Initialized() bool

// Map returns the map this iterator is iterating over.
func (it *Iter) Map() *Map

// Key returns a pointer to the current key. nil indicates end of iteration.
//
// Must not be called prior to Next.
func (it *Iter) Key() unsafe.Pointer

// Key returns a pointer to the current element. nil indicates end of
// iteration.
//
// Must not be called prior to Next.
func (it *Iter) Elem() unsafe.Pointer

// Next proceeds to the next element in iteration, which can be accessed via
// the Key and Elem methods.
//
// The table can be mutated during iteration, though there is no guarantee that
// the mutations will be visible to the iteration.
//
// Init must be called prior to Next.
func (it *Iter) Next()
