// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements multi-precision rational numbers.

package big

import (
	"github.com/shogo82148/std/fmt"
)

// A Rat represents a quotient a/b of arbitrary precision.
// The zero value for a Rat represents the value 0.
type Rat struct {
	a Int
	b nat
}

// NewRat creates a new Rat with numerator a and denominator b.
func NewRat(a, b int64) *Rat

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

// IsInt returns true if the denominator of x is 1.
func (x *Rat) IsInt() bool

// Num returns the numerator of x; it may be <= 0.
// The result is a reference to x's numerator; it
// may change if a new value is assigned to x.
func (x *Rat) Num() *Int

// Denom returns the denominator of x; it is always > 0.
// The result is a reference to x's denominator; it
// may change if a new value is assigned to x.
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

// Scan is a support routine for fmt.Scanner. It accepts the formats
// 'e', 'E', 'f', 'F', 'g', 'G', and 'v'. All formats are equivalent.
func (z *Rat) Scan(s fmt.ScanState, ch rune) error

// SetString sets z to the value of s and returns z and a boolean indicating
// success. s can be given as a fraction "a/b" or as a floating-point number
// optionally followed by an exponent. If the operation failed, the value of
// z is undefined but the returned value is nil.
func (z *Rat) SetString(s string) (*Rat, bool)

// String returns a string representation of z in the form "a/b" (even if b == 1).
func (x *Rat) String() string

// RatString returns a string representation of z in the form "a/b" if b != 1,
// and in the form "a" if b == 1.
func (x *Rat) RatString() string

// FloatString returns a string representation of z in decimal form with prec
// digits of precision after the decimal point and the last digit rounded.
func (x *Rat) FloatString(prec int) string

// Gob codec version. Permits backward-compatible changes to the encoding.

// GobEncode implements the gob.GobEncoder interface.
func (x *Rat) GobEncode() ([]byte, error)

// GobDecode implements the gob.GobDecoder interface.
func (z *Rat) GobDecode(buf []byte) error
