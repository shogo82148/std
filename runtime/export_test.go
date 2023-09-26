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

var Fastlog2 = fastlog2

var Atoi = atoi
var Atoi32 = atoi32
var ParseByteCount = parseByteCount

var Nanotime = nanotime
var NetpollBreak = netpollBreak
var Usleep = usleep

var PhysPageSize = physPageSize
var PhysHugePageSize = physHugePageSize

var NetpollGenericInit = netpollGenericInit

var Memmove = memmove
var MemclrNoHeapPointers = memclrNoHeapPointers

var CgoCheckPointer = cgoCheckPointer

const TracebackInnerFrames = tracebackInnerFrames
const TracebackOuterFrames = tracebackOuterFrames

var LockPartialOrder = lockPartialOrder

type LockRank lockRank

const PreemptMSupported = preemptMSupported

type LFNode struct {
	Next    uint64
	Pushcnt uintptr
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

const HashLoad = hashLoad

var Open = open
var Close = closefd
var Read = read
var Write = write

// blockWrapper is a wrapper type that ensures a T is placed within a
// large object. This is necessary for safely benchmarking things
// that manipulate the heap bitmap, like heapBitsSetType.
//
// More specifically, allocating threads assume they're the sole writers
// to their span's heap bits, which allows those writes to be non-atomic.
// The heap bitmap is written byte-wise, so if one tried to call heapBitsSetType
// on an existing object in a small object span, we might corrupt that
// span's bitmap with a concurrent byte write to the heap bitmap. Large
// object spans contain exactly one object, so we can be sure no other P
// is going to be allocating from it concurrently, hence this wrapper type
// which ensures we have a T in a large object span.

// arrayBlockWrapper is like blockWrapper, but the interior value is intended
// to be used as a backing store for a slice.

// arrayLargeBlockWrapper is like arrayBlockWrapper, but the interior array
// accommodates many more elements.

const PtrSize = goarch.PtrSize

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

var CasGStatusAlwaysTrack = &casgstatusAlwaysTrack

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

// testSysStat is the sysStat passed to test versions of various
// runtime structures. We do actually have to keep track of this
// because otherwise memstats.mappedReady won't actually line up
// with other stats in the runtime during tests.

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
	if goos.IsAix != 0 {
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

const SemTableSize = semTabSize

// SemTable is a wrapper around semTable exported for testing.
type SemTable struct {
	semTable
}

// mspan wrapper for testing.
type MSpan mspan

const (
	TimeHistSubBucketBits = timeHistSubBucketBits
	TimeHistNumSubBuckets = timeHistNumSubBuckets
	TimeHistNumBuckets    = timeHistNumBuckets
	TimeHistMinBucketBits = timeHistMinBucketBits
	TimeHistMaxBucketBits = timeHistMaxBucketBits
)

type TimeHistogram timeHistogram

var TimeHistogramMetricsBuckets = timeHistogramMetricsBuckets

// For GCTestMoveStackOnNextCall, it's important not to introduce an
// extra layer of call, since then there's a return before the "real"
// next call.
var GCTestMoveStackOnNextCall = gcTestMoveStackOnNextCall

const Raceenabled = raceenabled

const (
	GCBackgroundUtilization            = gcBackgroundUtilization
	GCGoalUtilization                  = gcGoalUtilization
	DefaultHeapMinimum                 = defaultHeapMinimum
	MemoryLimitHeapGoalHeadroomPercent = memoryLimitHeapGoalHeadroomPercent
	MemoryLimitMinHeapGoalHeadroom     = memoryLimitMinHeapGoalHeadroom
)

type GCController struct {
	gcControllerState
}

type GCControllerReviseDelta struct {
	HeapLive        int64
	HeapScan        int64
	HeapScanWork    int64
	StackScanWork   int64
	GlobalsScanWork int64
}

var Timediv = timediv

type PIController struct {
	piController
}

const (
	CapacityPerProc          = capacityPerProc
	GCCPULimiterUpdatePeriod = gcCPULimiterUpdatePeriod
)

type GCCPULimiter struct {
	limiter gcCPULimiterState
}

const ScavengePercent = scavengePercent

type Scavenger struct {
	Sleep      func(int64) int64
	Scavenge   func(uintptr) (uintptr, int64)
	ShouldStop func() bool
	GoMaxProcs func() int32

	released  atomic.Uintptr
	scavenger scavengerState
	stop      chan<- struct{}
	done      <-chan struct{}
}

type ScavengeIndex struct {
	i scavengeIndex
}

const GTrackingPeriod = gTrackingPeriod

var ZeroBase = unsafe.Pointer(&zerobase)

const UserArenaChunkBytes = userArenaChunkBytes

type UserArena struct {
	arena *userArena
}

var AlignUp = alignUp

var (
	IsPinned      = isPinned
	GetPinCounter = pinnerGetPinCounter
)
