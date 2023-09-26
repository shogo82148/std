// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// This table might overflow 127-bit exponent representations.
// In that case, truncate it after 1.0e38.

// Pow10 returns 10**e, the base-10 exponential of e.
//
// Special cases are:
//
//	Pow10(e) = +Inf for e > 309
//	Pow10(e) = 0 for e < -324
func Pow10(e int) float64
