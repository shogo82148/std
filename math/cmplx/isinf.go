// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmplx

// IsInfは、real(x)またはimag(x)のいずれかが無限大であるかどうかを報告します。
func IsInf(x complex128) bool

// Infは複素数の無限大、complex(+Inf, +Inf)を返します。
func Inf() complex128
