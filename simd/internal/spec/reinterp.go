// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spec

// ReshapeToUints reinterprets the bits of x as a {z} vector. The least
// significant bit of element 0 is bit 0
//
//specgen:name ReshapeToUint{zN}s
//specgen:require xN!=zN
func ReshapeToUints[xE Uints, xW Width, zE Uints](x Vec[xE, xW]) (z Vec[zE, xW])
