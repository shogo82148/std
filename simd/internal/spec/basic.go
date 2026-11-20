// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//specgen:category Basics

package spec

// Vector is a {{if scalable .T}}scalable SIMD vector of {{.TE}}s{{else}}{{.TW}}-bit SIMD vector of {{.TL}} {{.TE}}s{{end}}.
//
// Values of type {{.T}} are comparable. a == b if a[i] == b[i] for all elements.
// {{if eq .TB "float"}}Like {{.TE}} scalars, NaN != NaN, and +0 == -0.{{end}}
//
//specgen:name {{.T}}
type Vector[E Elt, W Width] Vec[E, W]

// Mask is a mask for a {{if scalable .T}}scalable SIMD vector of {{.TN}}-bit elements{{else}}SIMD vector of {{.TL}} {{.TN}}-bit elements{{end}}.
//
// Values of type {{.T}} are comparable.
//
//specgen:name {{.T}}
type Mask[E MaskElt, W Width] Vec[E, W]

// Len returns the number of elements in x.
func Len[E EltOrMask, W Width](x Vec[E, W]) int

// Elem returns the i'th element x. If i >= {{lanes .x "x"}}, it panics.
//
//	z = x[i]
func Elem[E Elt, W Width](x Vec[E, W], i uint8) (z E)

// SetElem returns x with its i'th element set to y. If i >= {{lanes .x "x"}}, it panics.
//
//	z = x
//	z[i] = y
func SetElem[E Elt, W Width](x Vec[E, W], i uint8, y E) (z Vec[E, W])

// Lo returns the lower half of x.
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
// See also [Enlarge].
//
//specgen:require z={xB}{xN}x{xL/2}
func Lo[E Elt, W SplittableWidth, zW Width](x Vec[E, W]) (z Vec[E, zW])

// Hi returns the upper half of x.
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
func Hi[E Elt, W SplittableWidth, zW Width](x Vec[E, W]) (z Vec[E, zW])

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
// See also [Enlarge].
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

// Enlarge returns x with twice as many elements, with the upper half zeroed.
// [Lo] reverses this: x.Enlarge().Lo() returns x.
//
// {{if eq .xL 2}}
//
//	z = concat(x, {0, 0})
//
// {{else}}
//
//	z = concat(x, {0, ..., 0})
//
// {{end}}
//
//specgen:require z={xB}{xN}x{xL*2}
func Enlarge[E Elt, W DoublableWidth, zW Width](x Vec[E, W]) (z Vec[E, zW])

// String returns x in the form "{x[0],x[1],...}". Each element is formatted
// like the %v format specifier.
func String[E Nums, W Width](x Vec[E, W]) string

// StringMask returns x in the form "{x[0],x[1],...}", where 0 means false and 1
// means true.
//
//specgen:name String
func StringMask[E MaskElt, W Width](x Vec[E, W]) string
