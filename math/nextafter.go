// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Nextafter32 returns the next representable float32 value after x towards y.
// Special cases:
//
//		Nextafter32(x, x)   = x
//	     Nextafter32(NaN, y) = NaN
//	     Nextafter32(x, NaN) = NaN
func Nextafter32(x, y float32) (r float32)

// Nextafter returns the next representable float64 value after x towards y.
// Special cases:
//
//		Nextafter64(x, x)   = x
//	     Nextafter64(NaN, y) = NaN
//	     Nextafter64(x, NaN) = NaN
func Nextafter(x, y float64) (r float64)
