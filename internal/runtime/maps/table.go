// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package maps implements Go's builtin map type.
package maps

import (
	"github.com/shogo82148/std/internal/abi"
	"github.com/shogo82148/std/unsafe"
)

func NewTable(mt *abi.SwissMapType, capacity uint64) *table

type Iter struct {
	key  unsafe.Pointer
	elem unsafe.Pointer
	typ  *abi.SwissMapType
	tab  *table

	// Snapshot of the groups at iteration initialization time. If the
	// table resizes during iteration, we continue to iterate over the old
	// groups.
	//
	// If the table grows we must consult the updated table to observe
	// changes, though we continue to use the snapshot to determine order
	// and avoid duplicating results.
	groups groupsReference

	// Copy of Table.clearSeq at iteration initialization time. Used to
	// detect clear during iteration.
	clearSeq uint64

	// Randomize iteration order by starting iteration at a random slot
	// offset.
	offset uint64

	// TODO: these could be merged into a single counter (and pre-offset
	// with offset).
	groupIdx uint64
	slotIdx  uint32
}

// Init initializes Iter for iteration.
func (it *Iter) Init(typ *abi.SwissMapType, t *table)

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
