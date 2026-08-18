// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

// An XPosMap is a map from fileindex and line of src.XPos to int32,
// implemented sparsely to save space (column and statement status are ignored).
// The sparse skeleton is constructed once, and then reused by ssa phases
// that (re)move values with statements attached.
type XPosMap struct {
	// A map from file index to maps from line range to integers (block numbers)
	maps map[int32]*BiasedSparseMap
	// The next two fields provide a single-item cache for common case of repeated lines from same file.
	lastIndex int32
	lastMap   *BiasedSparseMap
}

// NewXPosMap constructs an xposmap valid for inputs which have a file index in the keys of x,
// and line numbers in the range x[file index].
// The resulting xposmap will panic if a caller attempts to set or add an XPos not in that range.
func NewXPosMap(x map[int]LineRange) *XPosMap

type LineRange struct {
	First, Last uint32
}

// Clear removes data from the map but leaves the sparse skeleton.
func (m *XPosMap) Clear()

// MapFor returns the line range map for a given file index.
func (m *XPosMap) MapFor(index int32) *BiasedSparseMap

// Set inserts p->v into the map.
// If p does not fall within the Set of fileindex->lineRange used to construct m, this will panic.
func (m *XPosMap) Set(p src.XPos, v int32)

// Get returns the int32 associated with the file index and line of p.
func (m *XPosMap) Get(p src.XPos) (int32, bool)

// Add adds p to m, treating m as a set instead of as a map.
// If p does not fall within the set of fileindex->lineRange used to construct m, this will panic.
// Use clear() in between set/map interpretations of m.
func (m *XPosMap) Add(p src.XPos)

// Contains returns whether the file index and line of p are in m,
// treating m as a set instead of as a map.
func (m *XPosMap) Contains(p src.XPos) bool

// Remove removes the file index and line for p from m,
// whether m is currently treated as a map or set.
func (m *XPosMap) Remove(p src.XPos)

// ForeachEntry applies f to each (fileindex, line, value) triple in m.
func (m *XPosMap) ForeachEntry(f func(j int32, l uint, v int32))
