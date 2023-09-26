// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements the Bits type used for testing Float operations
// via an independent (albeit slower) representations for floating-point
// numbers.

package big

// A Bits value b represents a finite floating-point number x of the form
//
//	x = 2**b[0] + 2**b[1] + ... 2**b[len(b)-1]
//
// The order of slice elements is not significant. Negative elements may be
// used to form fractions. A Bits value is normalized if each b[i] occurs at
// most once. For instance Bits{0, 0, 1} is not normalized but represents the
// same floating-point number as Bits{2}, which is normalized. The zero (nil)
// value of Bits is a ready to use Bits value and represents the value 0.
type Bits []int
