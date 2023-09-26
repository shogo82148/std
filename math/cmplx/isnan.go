// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

// IsNaN returns true if either real(x) or imag(x) is NaN
// and neither is an infinity.
func IsNaN(x complex128) bool

// NaN returns a complex “not-a-number” value.
func NaN() complex128
