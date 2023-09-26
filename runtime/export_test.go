// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Export guts for testing.

package runtime

var Fadd64 = fadd64
var Fsub64 = fsub64
var Fmul64 = fmul64
var Fdiv64 = fdiv64
var F64to32 = f64to32
var F32to64 = f32to64
var Fcmp64 = fcmp64
var Fintto64 = fintto64
var F64toint = f64toint
var Sqrt = sqrt

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

var StringHash = stringHash
var BytesHash = bytesHash
var Int32Hash = int32Hash
var Int64Hash = int64Hash
var EfaceHash = efaceHash
var IfaceHash = ifaceHash

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
