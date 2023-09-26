// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package aes implements AES encryption (formerly Rijndael), as defined in
// U.S. Federal Information Processing Standards Publication 197.
package aes

// AES is based on the mathematical behavior of binary polynomials
// (polynomials over GF(2)) modulo the irreducible polynomial x⁸ + x⁴ + x³ + x + 1.
// Addition of these binary polynomials corresponds to binary xor.
// Reducing mod poly corresponds to binary xor with poly every
// time a 0x100 bit appears.

// Powers of x mod poly in GF(2).

// FIPS-197 Figure 7. S-box substitution values in hexadecimal format.

// FIPS-197 Figure 14.  Inverse S-box substitution values in hexadecimal format.
