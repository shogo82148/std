// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//simdgen:category Shifts

package spec

// ScaleSaturated multiplies x[i] by 2^scale[i], with saturation.
//
// Positive exponents scale up (shift left with saturation). Negative exponents
// scale down ({{if eq .xB "int"}}arithmetic{{else}}logical{{end}} shift right).
//
//	z[i] = sat(x[i] * 2^scale[i])
//
//specgen:require scale=Int{xN}x{xL}
func ScaleSaturated[E Ints | Uints, W Width, yE Ints](x Vec[E, W], scale Vec[yE, W]) (z Vec[E, W])

// {{if eq .xB "int" -}}
// ShiftAllRight arithmetically shifts each element of x right by shift bits.
// If shift is greater than the element width, the result is 0 or -1.
// {{- else -}}
// ShiftAllRight logically shifts each element of x right by shift bits.
// If shift is greater than the element width, the result is 0.
// {{- end}}
//
//	z[i] = x[i] >> shift
func ShiftAllRight[E Ints | Uints, W Width](x Vec[E, W], shift uint64) (z Vec[E, W])

// ShiftAllLeft shifts each element of x left by shift bits.
// If shift is greater than the element width, the result is 0.
//
//	z[i] = x[i] << shift
func ShiftAllLeft[E Ints | Uints, W Width](x Vec[E, W], shift uint64) (z Vec[E, W])

// {{if eq .xB "int" -}}
// ShiftRight arithmetically shifts x[i] right by shift[i] bits.
// If shift[i] is greater than the element width, the result is 0 or -1.
// {{- else -}}
// ShiftRight logically shifts x[i] right by shift[i] bits.
// If shift[i] is greater than the element width, the result is 0.
// {{- end}}
//
//	z[i] = x[i] >> shift[i]
//
//specgen:require shift=Uint{xN}x{xL}
func ShiftRight[E Ints | Uints, W Width, yE Uints](x Vec[E, W], shift Vec[yE, W]) (z Vec[E, W])

// ShiftLeft shifts x[i] left by shift[i] bits.
// If shift[i] is greater than the element width, the result is 0.
//
//	z[i] = x[i] << shift[i]
//
//specgen:require shift=Uint{xN}x{xL}
func ShiftLeft[E Ints | Uints, W Width, yE Uints](x Vec[E, W], shift Vec[yE, W]) (z Vec[E, W])

// RotateAllLeft rotates each element of x left by shift bits.
//
//	z[i] = rotateLeft(x[i], shift)
func RotateAllLeft[E Uints, W Width](x Vec[E, W], shift uint64) (z Vec[E, W])

// RotateLeft rotates x left by shift bits, elementwise.
//
//	z[i] = rotateLeft(x[i], shift[i])
func RotateLeft[E Uints, W Width](x, shift Vec[E, W]) (z Vec[E, W])

// RotateAllRight rotates each element of x right by shift bits.
//
//	z[i] = rotateRight(x[i], shift)
func RotateAllRight[E Uints, W Width](x Vec[E, W], shift uint64) (z Vec[E, W])

// RotateRight rotates x right by shift bits, elementwise.
//
//	z[i] = rotateRight(x[i], shift[i])
func RotateRight[E Uints, W Width](x, shift Vec[E, W]) (z Vec[E, W])
