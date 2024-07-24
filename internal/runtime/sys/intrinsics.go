// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

// TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.
func TrailingZeros32(x uint32) int

// TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.
func TrailingZeros64(x uint64) int

// TrailingZeros8 returns the number of trailing zero bits in x; the result is 8 for x == 0.
func TrailingZeros8(x uint8) int

// Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
//
// nosplit because this is used in src/runtime/histogram.go, which make run in sensitive contexts.
//
//go:nosplit
func Len64(x uint64) (n int)

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int

// Len8 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len8(x uint8) int

// Bswap64 returns its input with byte order reversed
// 0x0102030405060708 -> 0x0807060504030201
func Bswap64(x uint64) uint64

// Bswap32 returns its input with byte order reversed
// 0x01020304 -> 0x04030201
func Bswap32(x uint32) uint32

// Prefetch prefetches data from memory addr to cache
//
// AMD64: Produce PREFETCHT0 instruction
//
// ARM64: Produce PRFM instruction with PLDL1KEEP option
func Prefetch(addr uintptr)

// PrefetchStreamed prefetches data from memory addr, with a hint that this data is being streamed.
// That is, it is likely to be accessed very soon, but only once. If possible, this will avoid polluting the cache.
//
// AMD64: Produce PREFETCHNTA instruction
//
// ARM64: Produce PRFM instruction with PLDL1STRM option
func PrefetchStreamed(addr uintptr)
