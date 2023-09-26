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

var HashLoad = &hashLoad

type Uintreg sys.Uintreg

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

// Span is a safe wrapper around an mspan, whose memory
// is managed manually.
type Span struct {
	*mspan
}

type TreapIter struct {
	treapIter
}

// Treap is a safe wrapper around mTreap for testing.
//
// It must never be heap-allocated because mTreap is
// notinheap.
//
//go:notinheap
type Treap struct {
	mTreap
}
