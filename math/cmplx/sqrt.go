// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

// Sqrt returns the square root of x.
// The result r is chosen so that real(r) ≥ 0 and imag(r) has the same sign as imag(x).
func Sqrt(x complex128) complex128
