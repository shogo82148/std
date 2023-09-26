// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Memory statistics

package runtime

// Statistics.
// If you edit this structure, also edit type MemStats below.
// Their layouts must match exactly.
//
// For detailed descriptions see the documentation for MemStats.
// Fields that differ from MemStats are further documented here.
//
// Many of these fields are updated on the fly, while others are only
// updated when updatememstats is called.

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

// Size of the trailing by_size array differs between mstats and MemStats,
// and all data after by_size is local to runtime, not exported.
// NumSizeClasses was changed, but we cannot change MemStats because of backward compatibility.
// sizeof_C_MStats is the size of the prefix of mstats that
// corresponds to MemStats. It should match Sizeof(MemStats{}).

// ReadMemStats populates m with memory allocator statistics.
//
// The returned memory allocator statistics are up to date as of the
// call to ReadMemStats. This is in contrast with a heap profile,
// which is a snapshot as of the most recently completed garbage
// collection cycle.
func ReadMemStats(m *MemStats)
