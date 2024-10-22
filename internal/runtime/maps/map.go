// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package maps implements Go's builtin map type.
package maps

import (
	"github.com/shogo82148/std/internal/abi"
	"github.com/shogo82148/std/unsafe"
)

type Map struct {
	// The number of filled slots (i.e. the number of elements in all
	// tables).
	used uint64

	// Type of this map.
	//
	// TODO(prattmic): Old maps pass this into every call instead of
	// keeping a reference in the map header. This is probably more
	// efficient and arguably more robust (crafty users can't reach into to
	// the map to change its type), but I leave it here for now for
	// simplicity.
	typ *abi.SwissMapType

	// seed is the hash seed, computed as a unique random number per map.
	// TODO(prattmic): Populate this on table initialization.
	seed uintptr

	// The directory of tables. The length of this slice is
	// `1 << globalDepth`. Multiple entries may point to the same table.
	// See top-level comment for more details.
	directory []*table

	// The number of bits to use in table directory lookups.
	globalDepth uint8

	// clearSeq is a sequence counter of calls to Clear. It is used to
	// detect map clears during iteration.
	clearSeq uint64
}

func NewMap(mt *abi.SwissMapType, capacity uint64) *Map

func (m *Map) Type() *abi.SwissMapType

func (m *Map) Used() uint64

// Get performs a lookup of the key that key points to. It returns a pointer to
// the element, or false if the key doesn't exist.
func (m *Map) Get(key unsafe.Pointer) (unsafe.Pointer, bool)

func (m *Map) Put(key, elem unsafe.Pointer)

// PutSlot returns a pointer to the element slot where an inserted element
// should be written.
//
// PutSlot never returns nil.
func (m *Map) PutSlot(key unsafe.Pointer) unsafe.Pointer

func (m *Map) Delete(key unsafe.Pointer)

// Clear deletes all entries from the map resulting in an empty map.
func (m *Map) Clear()
