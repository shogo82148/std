// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simdstable

package archsimd

// CarrylessMulEven computes the elementwise carryless multiplication of
// even-indexed elements of x and y. The result of each carryless multiply is
// twice the width of the input elements. The high bits are stored in elements
// z[2*i+1] and the low bits are stored in elements z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i], y[2*i])
//
// A carryless multiplication uses bitwise XOR instead of add-with-carry.
// For example, to compute the carryless multiply of 0b1110 and 0b1011:
//
//	     1110
//	   ⊗ 1011
//	  ───────
//	     1110
//	    1110
//	   0000
//	⊕ 1110
//	  ───────
//	  1100010
//
// Carryless multiply can also be viewed as multiplying polynomials with
// coefficients from GF(2). For example, the above example can be represented as
//
//	  (x^3 + x^2 + x^1) * (x^3 + x^1 + x^0)
//	= (x^6 + x^5 + (1^1)x^4 + (1^1)x^3 + (1^1)x^2 + x^1)
//	= (x^6 + x^5 + x^1)
//
// Asm: VPCLMULQDQ if CPU feature AVXCLMUL, Emulated otherwise, CPU Feature: AVX
func (x Uint64x2) CarrylessMulEven(y Uint64x2) (z Uint64x2)

// CarrylessMulOdd computes the elementwise carryless multiplication of
// odd-indexed elements of x and y. The result of each carryless multiply is
// twice the width of the input elements. The high bits are stored in elements
// z[2*i+1] and the low bits are stored in elements z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i+1], y[2*i+1])
//
// See [CarrylessMulEven] for details about carryless multiply.
//
// Asm: VPCLMULQDQ if CPU feature AVXCLMUL, Emulated otherwise, CPU Feature: AVX
func (x Uint64x2) CarrylessMulOdd(y Uint64x2) (z Uint64x2)

// CarrylessMulOddEven computes the elementwise carryless multiplication of
// odd-indexed elements of x and even-indexed elements of y. The result of each
// carryless multiply is twice the width of the input elements. The high bits
// are stored in elements z[2*i+1] and the low bits are stored in elements
// z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i+1], y[2*i])
//
// See [CarrylessMulEven] for details about carryless multiply.
//
// Asm: VPCLMULQDQ if CPU feature AVXCLMUL, Emulated otherwise, CPU Feature: AVX
func (x Uint64x2) CarrylessMulOddEven(y Uint64x2) (z Uint64x2)

// CarrylessMulEvenOdd computes the elementwise carryless multiplication of
// even-indexed elements of x and odd-indexed elements of y. The result of each
// carryless multiply is twice the width of the input elements. The high bits
// are stored in elements z[2*i+1] and the low bits are stored in elements
// z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i], y[2*i+1])
//
// See [CarrylessMulEven] for details about carryless multiply.
//
// Asm: VPCLMULQDQ if CPU feature AVXCLMUL, Emulated otherwise, CPU Feature: AVX
func (x Uint64x2) CarrylessMulEvenOdd(y Uint64x2) (z Uint64x2)

// CarrylessMulEven computes the elementwise carryless multiplication of
// even-indexed elements of x and y. The result of each carryless multiply is
// twice the width of the input elements. The high bits are stored in elements
// z[2*i+1] and the low bits are stored in elements z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i], y[2*i])
//
// A carryless multiplication uses bitwise XOR instead of add-with-carry.
// For example, to compute the carryless multiply of 0b1110 and 0b1011:
//
//	     1110
//	   ⊗ 1011
//	  ───────
//	     1110
//	    1110
//	   0000
//	⊕ 1110
//	  ───────
//	  1100010
//
// Carryless multiply can also be viewed as multiplying polynomials with
// coefficients from GF(2). For example, the above example can be represented as
//
//	  (x^3 + x^2 + x^1) * (x^3 + x^1 + x^0)
//	= (x^6 + x^5 + (1^1)x^4 + (1^1)x^3 + (1^1)x^2 + x^1)
//	= (x^6 + x^5 + x^1)
//
// Asm: VPCLMULQDQ if CPU feature AVX2CLMUL, Emulated otherwise, CPU Feature: AVX2
func (x Uint64x4) CarrylessMulEven(y Uint64x4) (z Uint64x4)

// CarrylessMulOdd computes the elementwise carryless multiplication of
// odd-indexed elements of x and y. The result of each carryless multiply is
// twice the width of the input elements. The high bits are stored in elements
// z[2*i+1] and the low bits are stored in elements z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i+1], y[2*i+1])
//
// See [CarrylessMulEven] for details about carryless multiply.
//
// Asm: VPCLMULQDQ if CPU feature AVX2CLMUL, Emulated otherwise, CPU Feature: AVX2
func (x Uint64x4) CarrylessMulOdd(y Uint64x4) (z Uint64x4)

// CarrylessMulOddEven computes the elementwise carryless multiplication of
// odd-indexed elements of x and even-indexed elements of y. The result of each
// carryless multiply is twice the width of the input elements. The high bits
// are stored in elements z[2*i+1] and the low bits are stored in elements
// z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i+1], y[2*i])
//
// See [CarrylessMulEven] for details about carryless multiply.
//
// Asm: VPCLMULQDQ if CPU feature AVX2CLMUL, Emulated otherwise, CPU Feature: AVX2
func (x Uint64x4) CarrylessMulOddEven(y Uint64x4) (z Uint64x4)

// CarrylessMulEvenOdd computes the elementwise carryless multiplication of
// even-indexed elements of x and odd-indexed elements of y. The result of each
// carryless multiply is twice the width of the input elements. The high bits
// are stored in elements z[2*i+1] and the low bits are stored in elements
// z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i], y[2*i+1])
//
// See [CarrylessMulEven] for details about carryless multiply.
//
// Asm: VPCLMULQDQ if CPU feature AVX2CLMUL, Emulated otherwise, CPU Feature: AVX2
func (x Uint64x4) CarrylessMulEvenOdd(y Uint64x4) (z Uint64x4)

// CarrylessMulEven computes the elementwise carryless multiplication of
// even-indexed elements of x and y. The result of each carryless multiply is
// twice the width of the input elements. The high bits are stored in elements
// z[2*i+1] and the low bits are stored in elements z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i], y[2*i])
//
// A carryless multiplication uses bitwise XOR instead of add-with-carry.
// For example, to compute the carryless multiply of 0b1110 and 0b1011:
//
//	     1110
//	   ⊗ 1011
//	  ───────
//	     1110
//	    1110
//	   0000
//	⊕ 1110
//	  ───────
//	  1100010
//
// Carryless multiply can also be viewed as multiplying polynomials with
// coefficients from GF(2). For example, the above example can be represented as
//
//	  (x^3 + x^2 + x^1) * (x^3 + x^1 + x^0)
//	= (x^6 + x^5 + (1^1)x^4 + (1^1)x^3 + (1^1)x^2 + x^1)
//	= (x^6 + x^5 + x^1)
//
// Asm: VPCLMULQDQ if CPU feature AVX512CLMUL, Emulated otherwise, CPU Feature: AVX512
func (x Uint64x8) CarrylessMulEven(y Uint64x8) (z Uint64x8)

// CarrylessMulOdd computes the elementwise carryless multiplication of
// odd-indexed elements of x and y. The result of each carryless multiply is
// twice the width of the input elements. The high bits are stored in elements
// z[2*i+1] and the low bits are stored in elements z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i+1], y[2*i+1])
//
// See [CarrylessMulEven] for details about carryless multiply.
//
// Asm: VPCLMULQDQ if CPU feature AVX512CLMUL, Emulated otherwise, CPU Feature: AVX512
func (x Uint64x8) CarrylessMulOdd(y Uint64x8) (z Uint64x8)

// CarrylessMulOddEven computes the elementwise carryless multiplication of
// odd-indexed elements of x and even-indexed elements of y. The result of each
// carryless multiply is twice the width of the input elements. The high bits
// are stored in elements z[2*i+1] and the low bits are stored in elements
// z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i+1], y[2*i])
//
// See [CarrylessMulEven] for details about carryless multiply.
//
// Asm: VPCLMULQDQ if CPU feature AVX512CLMUL, Emulated otherwise, CPU Feature: AVX512
func (x Uint64x8) CarrylessMulOddEven(y Uint64x8) (z Uint64x8)

// CarrylessMulEvenOdd computes the elementwise carryless multiplication of
// even-indexed elements of x and odd-indexed elements of y. The result of each
// carryless multiply is twice the width of the input elements. The high bits
// are stored in elements z[2*i+1] and the low bits are stored in elements
// z[2*i].
//
//	concat(z[2*i+1], z[2*i]) = clmul(x[2*i], y[2*i+1])
//
// See [CarrylessMulEven] for details about carryless multiply.
//
// Asm: VPCLMULQDQ if CPU feature AVX512CLMUL, Emulated otherwise, CPU Feature: AVX512
func (x Uint64x8) CarrylessMulEvenOdd(y Uint64x8) (z Uint64x8)
