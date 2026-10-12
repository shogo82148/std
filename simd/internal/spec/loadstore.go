// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//simdgen:category Load/Store

package spec

// BroadcastZ returns a vector with the input x assigned to all elements of the
// result.
//
//specgen:name Broadcast{{.z}}
func BroadcastZ[E Elt, W Width](x E) (z Vec[E, W])

// LoadZ loads a slice into a vector. If len(s) < {{lanes .z "z"}}, it panics.
//
//specgen:name Load{{.z}}
func LoadZ[E Elt, W Width](s []E) (z Vec[E, W])

// LoadZArray loads an array into a vector.
//
//specgen:name Load{{.z}}Array
func LoadZArray[E Elt, W FixedWidth](x *Array[E, W]) (z Vec[E, W])

// LoadZPart loads n=min(len(s), {{lanes .z "z"}}) elements from slice s as a
// vector and returns the vector and n. If len(s) < {{lanes .z "z"}}, the
// remaining vector elements will be zero.
//
//specgen:name Load{{.z}}Part
func LoadZPart[E Elt, W Width](s []E) (z Vec[E, W], n int)

// Store stores the elements of x into a slice. If len(s) < {{lanes .x "x"}}, it panics.
func Store[E Elt, W Width](x Vec[E, W], s []E)

// StoreArray stores the elements of x to an array.
func StoreArray[E Elt, W FixedWidth](x Vec[E, W], y *Array[E, W])

// StoreArrayMasked stores the masked elements of x to an array. It does not
// modify elements of y that are false in the mask.
//
//specgen:require maskN=xN
func StoreArrayMasked[E Elt, W FixedWidth, mE MaskElt](x Vec[E, W], y *Array[E, W], mask Vec[mE, W])

// StorePart stores n=min(len(s), {{lanes .x "x"}}) elements of x into s and returns n.
func StorePart[E Elt, W Width](x Vec[E, W], s []E) (n int)
