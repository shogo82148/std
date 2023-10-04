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

const PreemptMSupported = preemptMSupported

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

const PtrSize = goarch.PtrSize

var ForceGCPeriod = &forcegcperiod

var ReadUnaligned32 = readUnaligned32
var ReadUnaligned64 = readUnaligned64

const (
	ProfBufBlocking    = profBufBlocking
	ProfBufNonBlocking = profBufNonBlocking
)

const RuntimeHmapSize = unsafe.Sizeof(hmap{})

var CasGStatusAlwaysTrack = &casgstatusAlwaysTrack

const (
	PageSize         = pageSize
	PallocChunkPages = pallocChunkPages
	PageAlloc64Bit   = pageAlloc64Bit
	PallocSumBytes   = pallocSumBytes
)

const PageCachePages = pageCachePages

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

var Semacquire = semacquire
var Semrelease1 = semrelease1

const SemTableSize = semTabSize

const (
	TimeHistSubBucketBits = timeHistSubBucketBits
	TimeHistNumSubBuckets = timeHistNumSubBuckets
	TimeHistNumBuckets    = timeHistNumBuckets
	TimeHistMinBucketBits = timeHistMinBucketBits
	TimeHistMaxBucketBits = timeHistMaxBucketBits
)

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

var Timediv = timediv

const (
	CapacityPerProc          = capacityPerProc
	GCCPULimiterUpdatePeriod = gcCPULimiterUpdatePeriod
)

const ScavengePercent = scavengePercent

const GTrackingPeriod = gTrackingPeriod

var ZeroBase = unsafe.Pointer(&zerobase)

const UserArenaChunkBytes = userArenaChunkBytes

var AlignUp = alignUp

var (
	IsPinned      = isPinned
	GetPinCounter = pinnerGetPinCounter
)
