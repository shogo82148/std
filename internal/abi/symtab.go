// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

// A FuncFlag records bits about a function, passed to the runtime.
type FuncFlag uint8

const (
	// FuncFlagTopFrame indicates a function that appears at the top of its stack.
	// The traceback routine stop at such a function and consider that a
	// successful, complete traversal of the stack.
	// Examples of TopFrame functions include goexit, which appears
	// at the top of a user goroutine stack, and mstart, which appears
	// at the top of a system goroutine stack.
	FuncFlagTopFrame FuncFlag = 1 << iota

	// FuncFlagSPWrite indicates a function that writes an arbitrary value to SP
	// (any write other than adding or subtracting a constant amount).
	// The traceback routines cannot encode such changes into the
	// pcsp tables, so the function traceback cannot safely unwind past
	// SPWrite functions. Stopping at an SPWrite function is considered
	// to be an incomplete unwinding of the stack. In certain contexts
	// (in particular garbage collector stack scans) that is a fatal error.
	FuncFlagSPWrite

	// FuncFlagAsm indicates that a function was implemented in assembly.
	FuncFlagAsm
)

// A FuncID identifies particular functions that need to be treated
// specially by the runtime.
// Note that in some situations involving plugins, there may be multiple
// copies of a particular special runtime function.
type FuncID uint8

const (
	FuncIDNormal FuncID = iota
	FuncID_abort
	FuncID_asmcgocall
	FuncID_asyncPreempt
	FuncID_cgocallback
	FuncID_debugCallV2
	FuncID_gcBgMarkWorker
	FuncID_goexit
	FuncID_gogo
	FuncID_gopanic
	FuncID_handleAsyncEvent
	FuncID_mcall
	FuncID_morestack
	FuncID_mstart
	FuncID_panicwrap
	FuncID_rt0_go
	FuncID_runfinq
	FuncID_runtime_main
	FuncID_sigpanic
	FuncID_systemstack
	FuncID_systemstack_switch
	FuncIDWrapper
)

// ArgsSizeUnknown is set in Func.argsize to mark all functions
// whose argument size is unknown (C vararg functions, and
// assembly code without an explicit specification).
// This value is generated by the compiler, assembler, or linker.
const ArgsSizeUnknown = -0x80000000

// IDs for PCDATA and FUNCDATA tables in Go binaries.
//
// These must agree with ../../../runtime/funcdata.h.
const (
	PCDATA_UnsafePoint   = 0
	PCDATA_StackMapIndex = 1
	PCDATA_InlTreeIndex  = 2
	PCDATA_ArgLiveIndex  = 3

	FUNCDATA_ArgsPointerMaps    = 0
	FUNCDATA_LocalsPointerMaps  = 1
	FUNCDATA_StackObjects       = 2
	FUNCDATA_InlTree            = 3
	FUNCDATA_OpenCodedDeferInfo = 4
	FUNCDATA_ArgInfo            = 5
	FUNCDATA_ArgLiveInfo        = 6
	FUNCDATA_WrapInfo           = 7
)

// Special values for the PCDATA_UnsafePoint table.
const (
	UnsafePointSafe   = -1
	UnsafePointUnsafe = -2

	// UnsafePointRestart1(2) apply on a sequence of instructions, within
	// which if an async preemption happens, we should back off the PC
	// to the start of the sequence when resuming.
	// We need two so we can distinguish the start/end of the sequence
	// in case that two sequences are next to each other.
	UnsafePointRestart1 = -3
	UnsafePointRestart2 = -4

	// Like UnsafePointRestart1, but back to function entry if async preempted.
	UnsafePointRestartAtEntry = -5
)
