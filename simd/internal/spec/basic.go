// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//specgen:category Basics

package spec

// Len returns the number of elements in x.
func Len[E EltOrMask, W Width](x Vec[E, W]) int

// GetElem returns the i'th element x. If i >= {{lanes .x "x"}}, it panics.
//
//	z = x[i]
func GetElem[E Elt, W Width](x Vec[E, W], i uint8) (z E)

// SetElem returns x with its i'th element set to y. If i >= {{lanes .x "x"}}, it panics.
//
//	z = x
//	z[i] = y
func SetElem[E Elt, W Width](x Vec[E, W], i uint8, y E) (z Vec[E, W])

// GetLo returns the lower half of x.
//
// {{if eq .xL 4}}
//
//	z = {x[0], x[1]}
//
// {{else}}
//
//	z = {x[0], x[1], ..., x[n/2-1]}, where n = {{lanes .x "x"}}
//
// {{end}}
//
//specgen:require z={xB}{xN}x{xL/2}
func GetLo[E Elt, W SplittableWidth, zW Width](x Vec[E, W]) (z Vec[E, zW])

// GetHi returns the upper half of x.
//
// {{if eq .xL 4}}
//
//	z = {x[2], x[3]}
//
// {{else}}
//
//	z = {x[n/2], x[n/2+1], ..., x[n-1]}, where n = {{lanes .x "x"}}
//
// {{end}}
//
//specgen:require z={xB}{xN}x{xL/2}
func GetHi[E Elt, W SplittableWidth, zW Width](x Vec[E, W]) (z Vec[E, zW])

// SetLo returns x with its lower half set to y.
//
// {{if eq .xL 4}}
//
//	z = {y[0], y[1], x[2], x[3]}
//
// {{else}}
//
//	z = {y[0], y[1], ..., y[n/2-1], x[n/2], ..., x[n-1]}, where n = {{lanes .x "x"}}
//
// {{end}}
//
//specgen:require y={xB}{xN}x{xL/2}
func SetLo[E Elt, W SplittableWidth, yW Width](x Vec[E, W], y Vec[E, yW]) (z Vec[E, W])

// SetHi returns x with its upper half set to y.
//
// {{if eq .xL 4}}
//
//	z = {x[0], x[1], y[0], y[1]}
//
// {{else}}
//
//	z = {x[0], ..., x[n/2-1], y[0], y[1], ..., y[n/2-1]}, where n = {{lanes .x "x"}}
//
// {{end}}
//
//specgen:require y={xB}{xN}x{xL/2}
func SetHi[E Elt, W SplittableWidth, yW Width](x Vec[E, W], y Vec[E, yW]) (z Vec[E, W])

// String returns x in the form "{x[0],x[1],...}".
func String[E Nums, W Width](x Vec[E, W]) string

// StringMask returns x in the form "{x[0],x[1],...}", where 0 means false and 1
// means true.
//
//specgen:name String
func StringMask[E MaskElt, W Width](x Vec[E, W]) string
