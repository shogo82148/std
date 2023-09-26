// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Memory statistics

package runtime

// A MemStats records statistics about the memory allocator.
type MemStats struct {
	Alloc uint64

	TotalAlloc uint64

	Sys uint64

	Lookups uint64

	Mallocs uint64

	Frees uint64

	HeapAlloc uint64

	HeapSys uint64

	HeapIdle uint64

	HeapInuse uint64

	HeapReleased uint64

	HeapObjects uint64

	StackInuse uint64

	StackSys uint64

	MSpanInuse uint64

	MSpanSys uint64

	MCacheInuse uint64

	MCacheSys uint64

	BuckHashSys uint64

	GCSys uint64

	OtherSys uint64

	NextGC uint64

	LastGC uint64

	PauseTotalNs uint64

	PauseNs [256]uint64

	PauseEnd [256]uint64

	NumGC uint32

	NumForcedGC uint32

	GCCPUFraction float64

	EnableGC bool

	DebugGC bool

	BySize [61]struct {
		Size uint32

		Mallocs uint64

		Frees uint64
	}
}

// ReadMemStats populates m with memory allocator statistics.
//
// The returned memory allocator statistics are up to date as of the
// call to ReadMemStats. This is in contrast with a heap profile,
// which is a snapshot as of the most recently completed garbage
// collection cycle.
func ReadMemStats(m *MemStats)

// sysMemStat represents a global system statistic that is managed atomically.
//
// This type must structurally be a uint64 so that mstats aligns with MemStats.

// heapStatsDelta contains deltas of various runtime memory statistics
// that need to be updated together in order for them to be kept
// consistent with one another.

// consistentHeapStats represents a set of various memory statistics
// whose updates must be viewed completely to get a consistent
// state of the world.
//
// To write updates to memory stats use the acquire and release
// methods. To obtain a consistent global snapshot of these statistics,
// use read.
