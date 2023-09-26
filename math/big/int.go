// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements signed multi-precision integers.

package big

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/math/rand"
)

// An Int represents a signed multi-precision integer.
// The zero value for an Int represents the value 0.
type Int struct {
	neg bool
	abs nat
}

// Sign returns:
//
//	-1 if x <  0
//	 0 if x == 0
//	+1 if x >  0
func (x *Int) Sign() int

// SetInt64 sets z to x and returns z.
func (z *Int) SetInt64(x int64) *Int

// SetUint64 sets z to x and returns z.
func (z *Int) SetUint64(x uint64) *Int

// NewInt allocates and returns a new Int set to x.
func NewInt(x int64) *Int

// Set sets z to x and returns z.
func (z *Int) Set(x *Int) *Int

// Bits provides raw (unchecked but fast) access to x by returning its
// absolute value as a little-endian Word slice. The result and x share
// the same underlying array.
// Bits is intended to support implementation of missing low-level Int
// functionality outside this package; it should be avoided otherwise.
func (x *Int) Bits() []Word

// SetBits provides raw (unchecked but fast) access to z by setting its
// value to abs, interpreted as a little-endian Word slice, and returning
// z. The result and abs share the same underlying array.
// SetBits is intended to support implementation of missing low-level Int
// functionality outside this package; it should be avoided otherwise.
func (z *Int) SetBits(abs []Word) *Int

// Abs sets z to |x| (the absolute value of x) and returns z.
func (z *Int) Abs(x *Int) *Int

// Neg sets z to -x and returns z.
func (z *Int) Neg(x *Int) *Int

// Add sets z to the sum x+y and returns z.
func (z *Int) Add(x, y *Int) *Int

// Sub sets z to the difference x-y and returns z.
func (z *Int) Sub(x, y *Int) *Int

// Mul sets z to the product x*y and returns z.
func (z *Int) Mul(x, y *Int) *Int

// MulRange sets z to the product of all integers
// in the range [a, b] inclusively and returns z.
// If a > b (empty range), the result is 1.
func (z *Int) MulRange(a, b int64) *Int

// Binomial sets z to the binomial coefficient of (n, k) and returns z.
func (z *Int) Binomial(n, k int64) *Int

// Quo sets z to the quotient x/y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Quo implements truncated division (like Go); see QuoRem for more details.
func (z *Int) Quo(x, y *Int) *Int

// Rem sets z to the remainder x%y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Rem implements truncated modulus (like Go); see QuoRem for more details.
func (z *Int) Rem(x, y *Int) *Int

// QuoRem sets z to the quotient x/y and r to the remainder x%y
// and returns the pair (z, r) for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
//
// QuoRem implements T-division and modulus (like Go):
//
//	q = x/y      with the result truncated to zero
//	r = x - y*q
//
// (See Daan Leijen, “Division and Modulus for Computer Scientists”.)
// See DivMod for Euclidean division and modulus (unlike Go).
func (z *Int) QuoRem(x, y, r *Int) (*Int, *Int)

// Div sets z to the quotient x/y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see DivMod for more details.
func (z *Int) Div(x, y *Int) *Int

// Mod sets z to the modulus x%y for y != 0 and returns z.
// If y == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean modulus (unlike Go); see DivMod for more details.
func (z *Int) Mod(x, y *Int) *Int

// DivMod sets z to the quotient x div y and m to the modulus x mod y
// and returns the pair (z, m) for y != 0.
// If y == 0, a division-by-zero run-time panic occurs.
//
// DivMod implements Euclidean division and modulus (unlike Go):
//
//	q = x div y  such that
//	m = x - y*q  with 0 <= m < |q|
//
// (See Raymond T. Boute, “The Euclidean definition of the functions
// div and mod”. ACM Transactions on Programming Languages and
// Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992.
// ACM press.)
// See QuoRem for T-division and modulus (like Go).
func (z *Int) DivMod(x, y, m *Int) (*Int, *Int)

// Cmp compares x and y and returns:
//
//	-1 if x <  y
//	 0 if x == y
//	+1 if x >  y
func (x *Int) Cmp(y *Int) (r int)

func (x *Int) String() string

// Format is a support routine for fmt.Formatter. It accepts
// the formats 'b' (binary), 'o' (octal), 'd' (decimal), 'x'
// (lowercase hexadecimal), and 'X' (uppercase hexadecimal).
// Also supported are the full suite of package fmt's format
// verbs for integral types, including '+', '-', and ' '
// for sign control, '#' for leading zero in octal and for
// hexadecimal, a leading "0x" or "0X" for "%#x" and "%#X"
// respectively, specification of minimum digits precision,
// output field width, space or zero padding, and left or
// right justification.
func (x *Int) Format(s fmt.State, ch rune)

// Scan is a support routine for fmt.Scanner; it sets z to the value of
// the scanned number. It accepts the formats 'b' (binary), 'o' (octal),
// 'd' (decimal), 'x' (lowercase hexadecimal), and 'X' (uppercase hexadecimal).
func (z *Int) Scan(s fmt.ScanState, ch rune) error

// Int64 returns the int64 representation of x.
// If x cannot be represented in an int64, the result is undefined.
func (x *Int) Int64() int64

// Uint64 returns the uint64 representation of x.
// If x cannot be represented in a uint64, the result is undefined.
func (x *Int) Uint64() uint64

// SetString sets z to the value of s, interpreted in the given base,
// and returns z and a boolean indicating success. If SetString fails,
// the value of z is undefined but the returned value is nil.
//
// The base argument must be 0 or a value from 2 through MaxBase. If the base
// is 0, the string prefix determines the actual conversion base. A prefix of
// “0x” or “0X” selects base 16; the “0” prefix selects base 8, and a
// “0b” or “0B” prefix selects base 2. Otherwise the selected base is 10.
func (z *Int) SetString(s string, base int) (*Int, bool)

// SetBytes interprets buf as the bytes of a big-endian unsigned
// integer, sets z to that value, and returns z.
func (z *Int) SetBytes(buf []byte) *Int

// Bytes returns the absolute value of x as a big-endian byte slice.
func (x *Int) Bytes() []byte

// BitLen returns the length of the absolute value of x in bits.
// The bit length of 0 is 0.
func (x *Int) BitLen() int

// Exp sets z = x**y mod |m| (i.e. the sign of m is ignored), and returns z.
// If y <= 0, the result is 1 mod |m|; if m == nil or m == 0, z = x**y.
// See Knuth, volume 2, section 4.6.3.
func (z *Int) Exp(x, y, m *Int) *Int

// GCD sets z to the greatest common divisor of a and b, which both must
// be > 0, and returns z.
// If x and y are not nil, GCD sets x and y such that z = a*x + b*y.
// If either a or b is <= 0, GCD sets z = x = y = 0.
func (z *Int) GCD(x, y, a, b *Int) *Int

// ProbablyPrime performs n Miller-Rabin tests to check whether x is prime.
// If it returns true, x is prime with probability 1 - 1/4^n.
// If it returns false, x is not prime.
func (x *Int) ProbablyPrime(n int) bool

// Rand sets z to a pseudo-random number in [0, n) and returns z.
func (z *Int) Rand(rnd *rand.Rand, n *Int) *Int

// ModInverse sets z to the multiplicative inverse of g in the group ℤ/pℤ (where
// p is a prime) and returns z.
func (z *Int) ModInverse(g, p *Int) *Int

// Lsh sets z = x << n and returns z.
func (z *Int) Lsh(x *Int, n uint) *Int

// Rsh sets z = x >> n and returns z.
func (z *Int) Rsh(x *Int, n uint) *Int

// Bit returns the value of the i'th bit of x. That is, it
// returns (x>>i)&1. The bit index i must be >= 0.
func (x *Int) Bit(i int) uint

// SetBit sets z to x, with x's i'th bit set to b (0 or 1).
// That is, if b is 1 SetBit sets z = x | (1 << i);
// if b is 0 SetBit sets z = x &^ (1 << i). If b is not 0 or 1,
// SetBit will panic.
func (z *Int) SetBit(x *Int, i int, b uint) *Int

// And sets z = x & y and returns z.
func (z *Int) And(x, y *Int) *Int

// AndNot sets z = x &^ y and returns z.
func (z *Int) AndNot(x, y *Int) *Int

// Or sets z = x | y and returns z.
func (z *Int) Or(x, y *Int) *Int

// Xor sets z = x ^ y and returns z.
func (z *Int) Xor(x, y *Int) *Int

// Not sets z = ^x and returns z.
func (z *Int) Not(x *Int) *Int

// Gob codec version. Permits backward-compatible changes to the encoding.

// GobEncode implements the gob.GobEncoder interface.
func (x *Int) GobEncode() ([]byte, error)

// GobDecode implements the gob.GobDecoder interface.
func (z *Int) GobDecode(buf []byte) error

// MarshalJSON implements the json.Marshaler interface.
func (z *Int) MarshalJSON() ([]byte, error)

// UnmarshalJSON implements the json.Unmarshaler interface.
func (z *Int) UnmarshalJSON(text []byte) error

// MarshalText implements the encoding.TextMarshaler interface
func (z *Int) MarshalText() (text []byte, err error)

// UnmarshalText implements the encoding.TextUnmarshaler interface
func (z *Int) UnmarshalText(text []byte) error
