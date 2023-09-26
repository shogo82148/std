// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// pow10tab stores the pre-computed values 10**i for i < 32.

// pow10postab32 stores the pre-computed value for 10**(i*32) at index i.

// pow10negtab32 stores the pre-computed value for 10**(-i*32) at index i.

// Pow10 returns 10**n, the base-10 exponential of n.
//
// Special cases are:
//
//	Pow10(n) =    0 for n < -323
//	Pow10(n) = +Inf for n > 308
func Pow10(n int) float64
