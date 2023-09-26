// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Export guts for testing.

package runtime

import (
	"unsafe"
)

var Fadd64 = fadd64
var Fsub64 = fsub64
var Fmul64 = fmul64
var Fdiv64 = fdiv64
var F64to32 = f64to32
var F32to64 = f32to64
var Fcmp64 = fcmp64
var Fintto64 = fintto64
var F64toint = f64toint

var Entersyscall = entersyscall
var Exitsyscall = exitsyscall
var LockedOSThread = lockedOSThread
var Xadduintptr = atomic.Xadduintptr

var FuncPC = funcPC

var Fastlog2 = fastlog2

var Atoi = atoi
var Atoi32 = atoi32

var Nanotime = nanotime
var NetpollBreak = netpollBreak
var Usleep = usleep

var PhysPageSize = physPageSize
var PhysHugePageSize = physHugePageSize

var NetpollGenericInit = netpollGenericInit

var Memmove = memmove
var MemclrNoHeapPointers = memclrNoHeapPointers

var LockPartialOrder = lockPartialOrder

type LockRank lockRank

const PreemptMSupported = preemptMSupported

type LFNode struct {
	Next    uint64
	Pushcnt uintptr
}

// Temporary to enable register ABI bringup.
// TODO(register args): convert back to local variables in RunSchedLocalQueueEmptyTest that
// get passed to the "go" stmts there.
var RunSchedLocalQueueEmptyState struct {
	done  chan bool
	ready *uint32
	p     *p
}

var (
	StringHash = stringHash
	BytesHash  = bytesHash
	Int32Hash  = int32Hash
	Int64Hash  = int64Hash
	MemHash    = memhash
	MemHash32  = memhash32
	MemHash64  = memhash64
	EfaceHash  = efaceHash
	IfaceHash  = ifaceHash
)

var UseAeshash = &useAeshash

var HashLoad = &hashLoad

var Open = open
var Close = closefd
var Read = read
var Write = write

var BigEndian = sys.BigEndian

const PtrSize = sys.PtrSize

var ForceGCPeriod = &forcegcperiod

var ReadUnaligned32 = readUnaligned32
var ReadUnaligned64 = readUnaligned64

type ProfBuf profBuf

const (
	ProfBufBlocking    = profBufBlocking
	ProfBufNonBlocking = profBufNonBlocking
)

type RWMutex struct {
	rw rwmutex
}

const RuntimeHmapSize = unsafe.Sizeof(hmap{})

type G = g

type Sudog = sudog

const (
	PageSize         = pageSize
	PallocChunkPages = pallocChunkPages
	PageAlloc64Bit   = pageAlloc64Bit
	PallocSumBytes   = pallocSumBytes
)

// Expose pallocSum for testing.
type PallocSum pallocSum

// Expose pallocBits for testing.
type PallocBits pallocBits

// Expose pallocData for testing.
type PallocData pallocData

// Expose pageCache for testing.
type PageCache pageCache

const PageCachePages = pageCachePages

// Expose chunk index type.
type ChunkIdx chunkIdx

// Expose pageAlloc for testing. Note that because pageAlloc is
// not in the heap, so is PageAlloc.
type PageAlloc pageAlloc

// AddrRange is a wrapper around addrRange for testing.
type AddrRange struct {
	addrRange
}

// AddrRanges is a wrapper around addrRanges for testing.
type AddrRanges struct {
	addrRanges
	mutable bool
}

// BitRange represents a range over a bitmap.
type BitRange struct {
	I, N uint
}

// BaseChunkIdx is a convenient chunkIdx value which works on both
// 64 bit and 32 bit platforms, allowing the tests to share code
// between the two.
//
// This should not be higher than 0x100*pallocChunkBytes to support
// mips and mipsle, which only have 31-bit address spaces.
var BaseChunkIdx = func() ChunkIdx {
	var prefix uintptr
	if pageAlloc64Bit != 0 {
		prefix = 0xc000
	} else {
		prefix = 0x100
	}
	baseAddr := prefix * pallocChunkBytes
	if sys.GoosAix != 0 {
		baseAddr += arenaBaseOffset
	}
	return ChunkIdx(chunkIndex(baseAddr))
}()

type BitsMismatch struct {
	Base      uintptr
	Got, Want uint64
}

var Semacquire = semacquire
var Semrelease1 = semrelease1

// mspan wrapper for testing.
//
//go:notinheap
type MSpan mspan

const (
	TimeHistSubBucketBits   = timeHistSubBucketBits
	TimeHistNumSubBuckets   = timeHistNumSubBuckets
	TimeHistNumSuperBuckets = timeHistNumSuperBuckets
)

type TimeHistogram timeHistogram

var TimeHistogramMetricsBuckets = timeHistogramMetricsBuckets

// For GCTestMoveStackOnNextCall, it's important not to introduce an
// extra layer of call, since then there's a return before the "real"
// next call.
var GCTestMoveStackOnNextCall = gcTestMoveStackOnNextCall

const Raceenabled = raceenabled
