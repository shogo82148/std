// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package spec

// ConvertToZ converts element values to {zE}. The result has the same number of lanes.
//
//specgen:name ConvertTo{zE}
//specgen:require zL=xL zE!=xE
func ConvertToZ[xE Nums, xW Width, zE Nums, zW Width](x Vec[xE, xW]) (z Vec[zE, zW])

// ExtendLoLToZ extends the lowest {zL} vector elements to {zE}.
//
//specgen:name ExtendLo{zL}To{zE}
//specgen:require zB=xB zN>xN
func ExtendLoLToZ[E Ints | Uints, W FixedWidth, zE Ints | Uints](x Vec[E, W]) (z Vec[zE, W])

// ConvertLoLToZ converts the low-indexed {zL} elements of x to {zE}.
//
//specgen:name ConvertLo{zL}To{zE}
//specgen:require zL<xL
func ConvertLoLToZ[E Nums, W FixedWidth, zE Floats](x Vec[E, W]) (z Vec[zE, W])
