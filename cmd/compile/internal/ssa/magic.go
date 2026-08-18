// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

func Sdivisible(n uint, c int64) SdivisibleData

type SdivisibleData struct {
	K   int64
	M   uint64
	A   uint64
	Max uint64
}

// SdivisibleOK reports whether we should strength reduce a signed n-bit divisibility check by c.
func SdivisibleOK(n uint, c int64) bool

func SdivisibleOK16(c int16) bool

func SdivisibleOK8(c int8) bool

// Smagic computes the constants needed to strength reduce signed n-bit divides by the constant c.
// Must have c>0.
// The return values satisfy for all -2^(n-1) <= x < 2^(n-1)
//
//	trunc(x / c) = x * m >> (n+s) + (x < 0 ? 1 : 0)
func Smagic(n uint, c int64) SmagicData

type SmagicData struct {
	S int64
	M uint64
}

func SmagicOK(n uint, c int64) bool

func Udivisible(n uint, c int64) UdivisibleData

type UdivisibleData struct {
	K   int64
	M   uint64
	Max uint64
}

// UdivisibleOK reports whether we should strength reduce an unsigned n-bit divisibility check by c.
func UdivisibleOK(n uint, c int64) bool

func UdivisibleOK16(c int16) bool

func UdivisibleOK8(c int8) bool

// Umagic computes the constants needed to strength reduce unsigned n-bit divides by the constant uint64(c).
// The return values satisfy for all 0 <= x < 2^n
//
//	floor(x / uint64(c)) = x * (m + 2^n) >> (n+s)
func Umagic(n uint, c int64) UmagicData

type UmagicData struct {
	S int64
	M uint64
}
