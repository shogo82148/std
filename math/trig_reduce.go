// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// reduceThreshold is the maximum value of x where the reduction using Pi/4
// in 3 float64 parts still gives accurate results. This threshold
// is set by y*C being representable as a float64 without error
// where y is given by y = floor(x * (4 / Pi)) and C is the leading partial
// terms of 4/Pi. Since the leading terms (PI4A and PI4B in sin.go) have 30
// and 32 trailing zero bits, y should have less than 30 significant bits.
//	y < 1<<30  -> floor(x*4/Pi) < 1<<30 -> x < (1<<30 - 1) * Pi/4
// So, conservatively we can take x < 1<<29.
// Above this threshold Payne-Hanek range reduction must be used.

// mPi4 is the binary digits of 4/pi as a uint64 array,
// that is, 4/pi = Sum mPi4[i]*2^(-64*i)
// 19 64-bit digits and the leading one bit give 1217 bits
// of precision to handle the largest possible float64 exponent.
