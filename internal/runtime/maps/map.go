// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package maps implements Go's builtin map type.
package maps

import (
	"github.com/shogo82148/std/internal/abi"
	"github.com/shogo82148/std/internal/goarch"
	"github.com/shogo82148/std/unsafe"
)

// Note: changes here must be reflected in cmd/compile/internal/reflectdata/map.go:MapType.
type Map struct {
	// The number of filled slots (i.e. the number of elements in all
	// tables). Excludes deleted slots.
	// Must be first (known by the compiler, for len() builtin).
	used uint64

	// seed is the hash seed, computed as a unique random number per map.
	seed uintptr

	// The directory of tables.
	//
	// Normally dirPtr points to an array of table pointers
	//
	// dirPtr *[dirLen]*table
	//
	// The length (dirLen) of this array is `1 << globalDepth`. Multiple
	// entries may point to the same table. See top-level comment for more
	// details.
	//
	// Small map optimization: if the map always contained
	// abi.MapGroupSlots or fewer entries, it fits entirely in a
	// single group. In that case dirPtr points directly to a single group.
	//
	// dirPtr *group
	//
	// In this case, dirLen is 0. used counts the number of used slots in
	// the group. Note that small maps never have deleted slots (as there
	// is no probe sequence to maintain).
	dirPtr unsafe.Pointer
	dirLen int

	// The number of bits to use in table directory lookups.
	globalDepth uint8

	// The number of bits to shift out of the hash for directory lookups.
	// On 64-bit systems, this is 64 - globalDepth.
	globalShift uint8

	// writing is a flag that is toggled (XOR 1) while the map is being
	// written. Normally it is set to 1 when writing, but if there are
	// multiple concurrent writers, then toggling increases the probability
	// that both sides will detect the race.
	writing uint8

	// tombstonePossible is false if we know that no table in this map
	// contains a tombstone.
	tombstonePossible bool

	// clearSeq is a sequence counter of calls to Clear. It is used to
	// detect map clears during iteration.
	clearSeq uint64
}

// Use 64-bit hash on 64-bit systems, except on Wasm, where we use
// 32-bit hash (see runtime/hash32.go).
const Use64BitHash = goarch.PtrSize == 8 && goarch.IsWasm == 0

// If m is non-nil, it should be used rather than allocating.
//
// maxAlloc should be runtime.maxAlloc.
//
// TODO(prattmic): Put maxAlloc somewhere accessible.
func NewMap(mt *abi.MapType, hint uintptr, m *Map, maxAlloc uintptr) *Map

func NewEmptyMap() *Map

func (m *Map) Used() uint64

// Get performs a lookup of the key that key points to. It returns a pointer to
// the element, or false if the key doesn't exist.
func (m *Map) Get(typ *abi.MapType, key unsafe.Pointer) (unsafe.Pointer, bool)

func (m *Map) Put(typ *abi.MapType, key, elem unsafe.Pointer)

// PutSlot returns a pointer to the element slot where an inserted element
// should be written.
//
// PutSlot never returns nil.
func (m *Map) PutSlot(typ *abi.MapType, key unsafe.Pointer) unsafe.Pointer

func (m *Map) Delete(typ *abi.MapType, key unsafe.Pointer)

// Clear deletes all entries from the map resulting in an empty map.
func (m *Map) Clear(typ *abi.MapType)

func (m *Map) Clone(typ *abi.MapType) *Map
