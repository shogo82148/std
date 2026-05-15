// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package windows provides the syscall primitives required for the runtime.

package windows

// StdCallInfo is a structure used to pass parameters to the system call.
type StdCallInfo struct {
	Fn   uintptr
	N    uintptr
	Args uintptr
	R1   uintptr
	R2   uintptr
}

// StdCall calls a function using Windows' stdcall convention.
// The calling thread's last-error code value is cleared before calling the function,
// and stored in the return value.
//
//go:noescape
func StdCall(fn *StdCallInfo) uint32

// AsmStdCallAddr is the address of a function that accepts a pointer
// to [StdCallInfo] stored on the stack following the C calling convention,
// and calls the function using Windows' stdcall calling convention.
// Shouldn't be called directly from Go.
func AsmStdCallAddr() uintptr
