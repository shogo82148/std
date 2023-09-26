// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big

// Sqrt sets z to the rounded square root of x, and returns it.
//
// If z's precision is 0, it is changed to x's precision before the
// operation. Rounding is performed according to z's precision and
// rounding mode.
//
// The function panics if z < 0. The value of z is undefined in that
// case.
func (z *Float) Sqrt(x *Float) *Float
