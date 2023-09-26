// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package des implements the Data Encryption Standard (DES) and the
// Triple Data Encryption Algorithm (TDEA) as defined
// in U.S. Federal Information Processing Standards Publication 46-3.
package des

// Used to perform an initial permutation of a 64-bit input block.

// Used to perform a final permutation of a 4-bit preoutput block. This is the
// inverse of initialPermutation

// Used to expand an input block of 32 bits, producing an output block of 48
// bits.

// Yields a 32-bit output from a 32-bit input

// Used in the key schedule to select 56 bits
// from a 64-bit input.

// Used in the key schedule to produce each subkey by selecting 48 bits from
// the 56-bit input

// 8 S-boxes composed of 4 rows and 16 columns
// Used in the DES cipher function

// Size of left rotation per round in each half of the key schedule
