// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run make_tables.go

// Package bits implements bit counting and manipulation
// functions for the predeclared unsigned integer types.
package bits

// UintSize is the size of a uint in bits.
const UintSize = uintSize

// LeadingZeros returns the number of leading zero bits in x; the result is UintSize for x == 0.
func LeadingZeros(x uint) int

// LeadingZeros8 returns the number of leading zero bits in x; the result is 8 for x == 0.
func LeadingZeros8(x uint8) int

// LeadingZeros16 returns the number of leading zero bits in x; the result is 16 for x == 0.
func LeadingZeros16(x uint16) int

// LeadingZeros32 returns the number of leading zero bits in x; the result is 32 for x == 0.
func LeadingZeros32(x uint32) int

// LeadingZeros64 returns the number of leading zero bits in x; the result is 64 for x == 0.
func LeadingZeros64(x uint64) int

// See http://supertech.csail.mit.edu/papers/debruijn.pdf

// TrailingZeros returns the number of trailing zero bits in x; the result is UintSize for x == 0.
func TrailingZeros(x uint) int

// TrailingZeros8 returns the number of trailing zero bits in x; the result is 8 for x == 0.
func TrailingZeros8(x uint8) int

// TrailingZeros16 returns the number of trailing zero bits in x; the result is 16 for x == 0.
func TrailingZeros16(x uint16) (n int)

// TrailingZeros32 returns the number of trailing zero bits in x; the result is 32 for x == 0.
func TrailingZeros32(x uint32) int

// TrailingZeros64 returns the number of trailing zero bits in x; the result is 64 for x == 0.
func TrailingZeros64(x uint64) int

// OnesCount returns the number of one bits ("population count") in x.
func OnesCount(x uint) int

// OnesCount8 returns the number of one bits ("population count") in x.
func OnesCount8(x uint8) int

// OnesCount16 returns the number of one bits ("population count") in x.
func OnesCount16(x uint16) int

// OnesCount32 returns the number of one bits ("population count") in x.
func OnesCount32(x uint32) int

// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int

// RotateLeft returns the value of x rotated left by (k mod UintSize) bits.
// To rotate x right by k bits, call RotateLeft(x, -k).
func RotateLeft(x uint, k int) uint

// RotateLeft8 returns the value of x rotated left by (k mod 8) bits.
// To rotate x right by k bits, call RotateLeft8(x, -k).
func RotateLeft8(x uint8, k int) uint8

// RotateLeft16 returns the value of x rotated left by (k mod 16) bits.
// To rotate x right by k bits, call RotateLeft16(x, -k).
func RotateLeft16(x uint16, k int) uint16

// RotateLeft32 returns the value of x rotated left by (k mod 32) bits.
// To rotate x right by k bits, call RotateLeft32(x, -k).
func RotateLeft32(x uint32, k int) uint32

// RotateLeft64 returns the value of x rotated left by (k mod 64) bits.
// To rotate x right by k bits, call RotateLeft64(x, -k).
func RotateLeft64(x uint64, k int) uint64

// Reverse returns the value of x with its bits in reversed order.
func Reverse(x uint) uint

// Reverse8 returns the value of x with its bits in reversed order.
func Reverse8(x uint8) uint8

// Reverse16 returns the value of x with its bits in reversed order.
func Reverse16(x uint16) uint16

// Reverse32 returns the value of x with its bits in reversed order.
func Reverse32(x uint32) uint32

// Reverse64 returns the value of x with its bits in reversed order.
func Reverse64(x uint64) uint64

// ReverseBytes returns the value of x with its bytes in reversed order.
func ReverseBytes(x uint) uint

// ReverseBytes16 returns the value of x with its bytes in reversed order.
func ReverseBytes16(x uint16) uint16

// ReverseBytes32 returns the value of x with its bytes in reversed order.
func ReverseBytes32(x uint32) uint32

// ReverseBytes64 returns the value of x with its bytes in reversed order.
func ReverseBytes64(x uint64) uint64

// Len returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len(x uint) int

// Len8 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len8(x uint8) int

// Len16 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len16(x uint16) (n int)

// Len32 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len32(x uint32) (n int)

// Len64 returns the minimum number of bits required to represent x; the result is 0 for x == 0.
func Len64(x uint64) (n int)
