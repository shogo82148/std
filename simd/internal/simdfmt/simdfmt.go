// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package simdfmt contains formatting helpers shared by the simd packages.
package simdfmt

// Number is the set of element types that SliceToString can format.
type Number interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// SliceToString returns a string representation of x of the form "{a,b,c}".
func SliceToString[T Number](x []T) string
