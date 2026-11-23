// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd && wasm

package archsimd

// TruncToInt8 converts each element of x to int8 by discarding its high
// bits.
//
//	z[i] = int8(x[i])
//
// The results are stored in the low 8 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Int16x8) TruncToInt8() (z Int8x16)

// TruncToUint8 converts each element of x to uint8 by discarding its high
// bits.
//
//	z[i] = uint8(x[i])
//
// The results are stored in the low 8 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Uint16x8) TruncToUint8() (z Uint8x16)

// TruncToInt16 converts each element of x to int16 by discarding its high
// bits.
//
//	z[i] = int16(x[i])
//
// The results are stored in the low 4 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Int32x4) TruncToInt16() (z Int16x8)

// TruncToUint16 converts each element of x to uint16 by discarding its high
// bits.
//
//	z[i] = uint16(x[i])
//
// The results are stored in the low 4 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Uint32x4) TruncToUint16() (z Uint16x8)

// TruncToInt32 converts each element of x to int32 by discarding its high
// bits.
//
//	z[i] = int32(x[i])
//
// The results are stored in the low 2 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Int64x2) TruncToInt32() (z Int32x4)

// TruncToUint32 converts each element of x to uint32 by discarding its high
// bits.
//
//	z[i] = uint32(x[i])
//
// The results are stored in the low 2 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Uint64x2) TruncToUint32() (z Uint32x4)

// SaturateToInt8 converts each element of x to int8, saturating values
// that are outside the range of int8.
//
//	z[i] = sat(x[i])
//
// The results are stored in the low 8 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Int16x8) SaturateToInt8() (z Int8x16)

// SaturateToUint8 converts each element of x to uint8, saturating values
// that are outside the range of uint8. Negative values become 0.
//
//	z[i] = sat(x[i])
//
// The results are stored in the low 8 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Int16x8) SaturateToUint8() (z Uint8x16)

// SaturateToUint8 converts each element of x to uint8, saturating values
// that are outside the range of uint8.
//
//	z[i] = sat(x[i])
//
// The results are stored in the low 8 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Uint16x8) SaturateToUint8() (z Uint8x16)

// SaturateToInt16 converts each element of x to int16, saturating values
// that are outside the range of int16.
//
//	z[i] = sat(x[i])
//
// The results are stored in the low 4 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Int32x4) SaturateToInt16() (z Int16x8)

// SaturateToUint16 converts each element of x to uint16, saturating values
// that are outside the range of uint16. Negative values become 0.
//
//	z[i] = sat(x[i])
//
// The results are stored in the low 4 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Int32x4) SaturateToUint16() (z Uint16x8)

// SaturateToUint16 converts each element of x to uint16, saturating values
// that are outside the range of uint16.
//
//	z[i] = sat(x[i])
//
// The results are stored in the low 4 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Uint32x4) SaturateToUint16() (z Uint16x8)

// SaturateToInt32 converts each element of x to int32, saturating values
// that are outside the range of int32.
//
//	z[i] = sat(x[i])
//
// The results are stored in the low 2 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Int64x2) SaturateToInt32() (z Int32x4)

// SaturateToUint32 converts each element of x to uint32, saturating values
// that are outside the range of uint32. Negative values become 0.
//
//	z[i] = sat(x[i])
//
// The results are stored in the low 2 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Int64x2) SaturateToUint32() (z Uint32x4)

// SaturateToUint32 converts each element of x to uint32, saturating values
// that are outside the range of uint32.
//
//	z[i] = sat(x[i])
//
// The results are stored in the low 2 elements of z.
// The remaining elements of z are zero.
//
// Emulated
func (x Uint64x2) SaturateToUint32() (z Uint32x4)
