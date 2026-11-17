// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd

package simd

// VectorBits returns the bit length of the vectors used by this package.
//
// If [Emulated] is true, this is 128. Otherwise, it is the longest vector
// length supported by the hardware and the simd package. It can be artificially
// reduced by setting the GODEBUG=simd=<smaller size> environment variable
// before running a program.
func VectorBits() int

// Emulated reports whether this package's operations are emulated in software
// rather than executed on vector hardware.
func Emulated() bool

// HasHardwareCarrylessMultiply returns whether this platform
// has a hardware-implemented version of carryless multiply.
// With default GODEBUG=simd settings, if this is false,
// it is emulated and merely slow, but with non-default settings
// this can indicate the possibility of a missing instruction
// that will fail ("SIGILL") if it is executed.
func HasHardwareCarrylessMultiply() bool
