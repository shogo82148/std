// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

// This type and constants are for encoding different
// kinds of bounds check failures.
type BoundsErrorCode uint8

const (
	BoundsIndex      BoundsErrorCode = iota
	BoundsSliceAlen
	BoundsSliceAcap
	BoundsSliceB
	BoundsSlice3Alen
	BoundsSlice3Acap
	BoundsSlice3B
	BoundsSlice3C
	BoundsConvert
)

const (
	BoundsMaxReg   = 15
	BoundsMaxConst = 31
)

// Encode bounds failure information into an integer for PCDATA_PanicBounds.
// Register numbers must be in 0-15. Constants must be in 0-31.
func BoundsEncode(code BoundsErrorCode, signed, xIsReg, yIsReg bool, xVal, yVal int) int

func BoundsDecode(v int) (code BoundsErrorCode, signed, xIsReg, yIsReg bool, xVal, yVal int)
