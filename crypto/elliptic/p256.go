// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !amd64
// +build !amd64

package elliptic

// p256Precomputed contains precomputed values to aid the calculation of scalar
// multiples of the base point, G. It's actually two, equal length, tables
// concatenated.
//
// The first table contains (x,y) field element pairs for 16 multiples of the
// base point, G.
//
//   Index  |  Index (binary) | Value
//       0  |           0000  | 0G (all zeros, omitted)
//       1  |           0001  | G
//       2  |           0010  | 2**64G
//       3  |           0011  | 2**64G + G
//       4  |           0100  | 2**128G
//       5  |           0101  | 2**128G + G
//       6  |           0110  | 2**128G + 2**64G
//       7  |           0111  | 2**128G + 2**64G + G
//       8  |           1000  | 2**192G
//       9  |           1001  | 2**192G + G
//      10  |           1010  | 2**192G + 2**64G
//      11  |           1011  | 2**192G + 2**64G + G
//      12  |           1100  | 2**192G + 2**128G
//      13  |           1101  | 2**192G + 2**128G + G
//      14  |           1110  | 2**192G + 2**128G + 2**64G
//      15  |           1111  | 2**192G + 2**128G + 2**64G + G
//
// The second table follows the same style, but the terms are 2**32G,
// 2**96G, 2**160G, 2**224G.
//
// This is ~2KB of data.

// p256Zero31 is 0 mod p.
