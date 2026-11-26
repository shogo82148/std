// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.simd

package archsimd

// AVXPCLMULQDQ returns whether the CPU supports carryless multiplication on
// 128-bit vectors.
//
// Deprecated: use X86.AVXCLMUL() instead.
//
//go:fix inline
func (X86Features) AVXPCLMULQDQ() bool

// VPCLMULQDQ returns whether the CPU supports carryless multiplication on
// 256-bit vectors.
//
// Deprecated: use X86.AVX2CLMUL() instead.
//
//go:fix inline
func (X86Features) VPCLMULQDQ() bool

// AVX512VPCLMULQDQ returns whether the CPU supports carryless multiplication on
// 512-bit vectors.
//
// Deprecated: use X86.AVX512CLMUL() instead.
//
//go:fix inline
func (X86Features) AVX512VPCLMULQDQ() bool

// VAES returns whether the CPU supports AES operations on 256-bit vectors.
//
// Deprecated: use X86.AVX2AES() instead.
//
//go:fix inline
func (X86Features) VAES() bool

// AVX512VAES returns whether the CPU supports AES operations on 512-bit vectors.
//
// Deprecated: use X86.AVX512AES() instead.
//
//go:fix inline
func (X86Features) AVX512VAES() bool
