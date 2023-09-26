// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements multi-precision rational numbers.

package big

// A Rat represents a quotient a/b of arbitrary precision.
// The zero value for a Rat represents the value 0.
type Rat struct {
	a, b Int
}

// NewRat creates a new Rat with numerator a and denominator b.
func NewRat(a, b int64) *Rat

// SetFloat64 sets z to exactly f and returns z.
// If f is not finite, SetFloat returns nil.
func (z *Rat) SetFloat64(f float64) *Rat

// Float32 returns the nearest float32 value for x and a bool indicating
// whether f represents x exactly. If the magnitude of x is too large to
// be represented by a float32, f is an infinity and exact is false.
// The sign of f always matches the sign of x, even if f == 0.
func (x *Rat) Float32() (f float32, exact bool)

// Float64 returns the nearest float64 value for x and a bool indicating
// whether f represents x exactly. If the magnitude of x is too large to
// be represented by a float64, f is an infinity and exact is false.
// The sign of f always matches the sign of x, even if f == 0.
func (x *Rat) Float64() (f float64, exact bool)

// SetFrac sets z to a/b and returns z.
func (z *Rat) SetFrac(a, b *Int) *Rat

// SetFrac64 sets z to a/b and returns z.
func (z *Rat) SetFrac64(a, b int64) *Rat

// SetInt sets z to x (by making a copy of x) and returns z.
func (z *Rat) SetInt(x *Int) *Rat

// SetInt64 sets z to x and returns z.
func (z *Rat) SetInt64(x int64) *Rat

// Set sets z to x (by making a copy of x) and returns z.
func (z *Rat) Set(x *Rat) *Rat

// Abs sets z to |x| (the absolute value of x) and returns z.
func (z *Rat) Abs(x *Rat) *Rat

// Neg sets z to -x and returns z.
func (z *Rat) Neg(x *Rat) *Rat

// Inv sets z to 1/x and returns z.
func (z *Rat) Inv(x *Rat) *Rat

// Sign returns:
//
//	-1 if x <  0
//	 0 if x == 0
//	+1 if x >  0
func (x *Rat) Sign() int

// IsInt reports whether the denominator of x is 1.
func (x *Rat) IsInt() bool

// Num returns the numerator of x; it may be <= 0.
// The result is a reference to x's numerator; it
// may change if a new value is assigned to x, and vice versa.
// The sign of the numerator corresponds to the sign of x.
func (x *Rat) Num() *Int

// Denom returns the denominator of x; it is always > 0.
// The result is a reference to x's denominator; it
// may change if a new value is assigned to x, and vice versa.
func (x *Rat) Denom() *Int

// Cmp compares x and y and returns:
//
//	-1 if x <  y
//	 0 if x == y
//	+1 if x >  y
func (x *Rat) Cmp(y *Rat) int

// Add sets z to the sum x+y and returns z.
func (z *Rat) Add(x, y *Rat) *Rat

// Sub sets z to the difference x-y and returns z.
func (z *Rat) Sub(x, y *Rat) *Rat

// Mul sets z to the product x*y and returns z.
func (z *Rat) Mul(x, y *Rat) *Rat

// Quo sets z to the quotient x/y and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
func (z *Rat) Quo(x, y *Rat) *Rat
