// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

// IsInf reports whether either real(x) or imag(x) is an infinity.
func IsInf(x complex128) bool

// Inf returns a complex infinity, complex(+Inf, +Inf).
func Inf() complex128
