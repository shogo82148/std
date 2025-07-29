// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package windows provides the syscall primitives required for the runtime.

package windows

// MaxArgs should be divisible by 2, as Windows stack
// must be kept 16-byte aligned on syscall entry.
//
// Although it only permits maximum 42 parameters, it
// is arguably large enough.
const MaxArgs = 42

// StdCallInfo is a structure used to pass parameters to the system call.
type StdCallInfo struct {
	Fn   uintptr
	N    uintptr
	Args uintptr
	R1   uintptr
	R2   uintptr
	Err  uintptr
}

// StdCall calls a function using Windows' stdcall convention.
//
//go:noescape
func StdCall(fn *StdCallInfo)

// AsmStdCallAddr is the address of a function that accepts a pointer
// to [StdCallInfo] stored on the stack following the C calling convention,
// and calls the function using Windows' stdcall calling convention.
// Shouldn't be called directly from Go.
func AsmStdCallAddr() uintptr
