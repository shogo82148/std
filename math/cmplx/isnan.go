// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

// IsNaNは、real(x)またはimag(x)のいずれかがNaN（非数）であり、
// どちらも無限大でないかどうかを報告します。
func IsNaN(x complex128) bool

// NaNは複素数の「非数」値を返します。
func NaN() complex128
