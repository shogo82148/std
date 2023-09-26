// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Memory statistics

package runtime

// Statistics.
// If you edit this structure, also edit type MemStats below.

// A MemStats records statistics about the memory allocator.
type MemStats struct {
	Alloc      uint64
	TotalAlloc uint64
	Sys        uint64
	Lookups    uint64
	Mallocs    uint64
	Frees      uint64

	HeapAlloc    uint64
	HeapSys      uint64
	HeapIdle     uint64
	HeapInuse    uint64
	HeapReleased uint64
	HeapObjects  uint64

	StackInuse  uint64
	StackSys    uint64
	MSpanInuse  uint64
	MSpanSys    uint64
	MCacheInuse uint64
	MCacheSys   uint64
	BuckHashSys uint64
	GCSys       uint64
	OtherSys    uint64

	NextGC        uint64
	LastGC        uint64
	PauseTotalNs  uint64
	PauseNs       [256]uint64
	PauseEnd      [256]uint64
	NumGC         uint32
	GCCPUFraction float64
	EnableGC      bool
	DebugGC       bool

	BySize [61]struct {
		Size    uint32
		Mallocs uint64
		Frees   uint64
	}
}

// Size of the trailing by_size array differs between Go and C,
// and all data after by_size is local to runtime, not exported.
// NumSizeClasses was changed, but we can not change Go struct because of backward compatibility.
// sizeof_C_MStats is what C thinks about size of Go struct.

// ReadMemStats populates m with memory allocator statistics.
func ReadMemStats(m *MemStats)
