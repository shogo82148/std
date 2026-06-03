// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && (amd64 || arm64 || wasm)

package simd

// VectorBitSize returns the bit length of the longest vector available
// on the current hardware.  It can be artificially reduced by setting
// GODEBUG=simd=<smaller size> environment variable before running a program.
func VectorBitSize() int

// Emulated returns whether simd operations are emulated or
// running on actual vector hardware.
func Emulated() bool

// HasHardwareCarrylessMultiply returns whether this platform
// as a hardware-implemented version of carryless multiply.
// With default GODEBUG=simd settings, if this is false,
// it is emulated and merely slow, but with non-default settings
// this can indicate the possibility of a missing instruction
// that will fail ("SIGILL") if it is executed.
func HasHardwareCarrylessMultiply() bool
