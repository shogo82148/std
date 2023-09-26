// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Nextafter returns the next representable value after x towards y.
// If x == y, then x is returned.
//
// Special cases are:
//
//	Nextafter(NaN, y) = NaN
//	Nextafter(x, NaN) = NaN
func Nextafter(x, y float64) (r float64)
