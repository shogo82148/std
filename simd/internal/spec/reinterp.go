// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//simdgen:category Reinterpretation

package spec

// ToBits reinterprets the bits of each element of x as type {{.zE}}.
//
//specgen:name ToBits
//specgen:require xN=zN
func ToBits[xE Ints, xW Width, zE Uints](x Vec[xE, xW]) (z Vec[zE, xW])

// ToBitsFloat returns the IEEE 754 binary representation of each element of x.
//
//specgen:name ToBits
//specgen:require xN=zN
func ToBitsFloat[xE float32 | float64, xW Width, zE Uints](x Vec[xE, xW]) (z Vec[zE, xW])

// BitsTo reinterprets the bits of each element of x as type {{.zE}}.
//
//specgen:name BitsTo{{.zE | title}}
//specgen:require xN=zN
func BitsTo[xE Uints, xW Width, zE Ints](x Vec[xE, xW]) (z Vec[zE, xW])

// BitsToFloat reinterprets the bits of each element of x as type {{.zE}}.
//
//specgen:name BitsTo{{.zE | title}}
//specgen:require xN=zN
func BitsToFloat[xE Uints, xW Width, zE Floats](x Vec[xE, xW]) (z Vec[zE, xW])

// ReshapeToUints reinterprets the bits of x as a {{.z}} vector.
//
// Both the vector elements and the bits of each element are interpreted in
// little endian order.
//
// {{reshapeDiagram .x .z}}
//
//specgen:name ReshapeToUint{{.zN}}s
//specgen:require xN!=zN
func ReshapeToUints[xE Uints, xW Width, zE Uints](x Vec[xE, xW]) (z Vec[zE, xW])
