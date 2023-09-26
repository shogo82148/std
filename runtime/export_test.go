// Copyright 2010 The Go Authors.  All rights reserved.
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

var Entersyscall = entersyscall
var Exitsyscall = exitsyscall
var LockedOSThread = golockedOSThread
var Stackguard = stackguard

type LFNode struct {
	Next    *LFNode
	Pushcnt uintptr
}

var LFStackPush = lfstackpush_go
var LFStackPop = lfstackpop_go

type ParFor struct {
	body    *byte
	done    uint32
	Nthr    uint32
	nthrmax uint32
	thrseq  uint32
	Cnt     uint32
	Ctx     *byte
	wait    bool
}

var NewParFor = newParFor
var ParForSetup = parForSetup
var ParForDo = parForDo

var TestSchedLocalQueue1 = testSchedLocalQueue
var TestSchedLocalQueueSteal1 = testSchedLocalQueueSteal

var HaveGoodHash = haveGoodHash
var StringHash = stringHash
var BytesHash = bytesHash
var Int32Hash = int32Hash
var Int64Hash = int64Hash

var HashLoad = &hashLoad

var MemclrBytes = memclrBytes

var GogoBytes = gogoBytes
