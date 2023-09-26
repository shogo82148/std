// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package constant implements Values representing untyped
// Go constants and their corresponding operations.
//
// A special Unknown value may be used when a value
// is unknown due to an error. Operations on unknown
// values produce unknown values unless specified
// otherwise.
package constant

import (
	"github.com/shogo82148/std/go/token"
)

// Kind specifies the kind of value represented by a Value.
type Kind int

const (
	// unknown values
	Unknown Kind = iota

	// non-numeric values
	Bool
	String

	// numeric values
	Int
	Float
	Complex
)

// A Value represents the value of a Go constant.
type Value interface {
	Kind() Kind

	String() string

	ExactString() string

	implementsValue()
}

// Maximum supported mantissa precision.
// The spec requires at least 256 bits; typical implementations use 512 bits.

// Permit fractions with component sizes up to maxExp
// before switching to using floating-point numbers.

// MakeUnknown returns the Unknown value.
func MakeUnknown() Value

// MakeBool returns the Bool value for b.
func MakeBool(b bool) Value

// MakeString returns the String value for s.
func MakeString(s string) Value

// MakeInt64 returns the Int value for x.
func MakeInt64(x int64) Value

// MakeUint64 returns the Int value for x.
func MakeUint64(x uint64) Value

// MakeFloat64 returns the Float value for x.
// If x is -0.0, the result is 0.0.
// If x is not finite, the result is an Unknown.
func MakeFloat64(x float64) Value

// MakeFromLiteral returns the corresponding integer, floating-point,
// imaginary, character, or string value for a Go literal string. The
// tok value must be one of token.INT, token.FLOAT, token.IMAG,
// token.CHAR, or token.STRING. The final argument must be zero.
// If the literal string syntax is invalid, the result is an Unknown.
func MakeFromLiteral(lit string, tok token.Token, zero uint) Value

// BoolVal returns the Go boolean value of x, which must be a Bool or an Unknown.
// If x is Unknown, the result is false.
func BoolVal(x Value) bool

// StringVal returns the Go string value of x, which must be a String or an Unknown.
// If x is Unknown, the result is "".
func StringVal(x Value) string

// Int64Val returns the Go int64 value of x and whether the result is exact;
// x must be an Int or an Unknown. If the result is not exact, its value is undefined.
// If x is Unknown, the result is (0, false).
func Int64Val(x Value) (int64, bool)

// Uint64Val returns the Go uint64 value of x and whether the result is exact;
// x must be an Int or an Unknown. If the result is not exact, its value is undefined.
// If x is Unknown, the result is (0, false).
func Uint64Val(x Value) (uint64, bool)

// Float32Val is like Float64Val but for float32 instead of float64.
func Float32Val(x Value) (float32, bool)

// Float64Val returns the nearest Go float64 value of x and whether the result is exact;
// x must be numeric or an Unknown, but not Complex. For values too small (too close to 0)
// to represent as float64, Float64Val silently underflows to 0. The result sign always
// matches the sign of x, even for 0.
// If x is Unknown, the result is (0, false).
func Float64Val(x Value) (float64, bool)

// Val returns the underlying value for a given constant. Since it returns an
// interface, it is up to the caller to type assert the result to the expected
// type. The possible dynamic return types are:
//
//	x Kind             type of result
//	-----------------------------------------
//	Bool               bool
//	String             string
//	Int                int64 or *big.Int
//	Float              *big.Float or *big.Rat
//	everything else    nil
func Val(x Value) any

// Make returns the Value for x.
//
//	type of x        result Kind
//	----------------------------
//	bool             Bool
//	string           String
//	int64            Int
//	*big.Int         Int
//	*big.Float       Float
//	*big.Rat         Float
//	anything else    Unknown
func Make(x any) Value

// BitLen returns the number of bits required to represent
// the absolute value x in binary representation; x must be an Int or an Unknown.
// If x is Unknown, the result is 0.
func BitLen(x Value) int

// Sign returns -1, 0, or 1 depending on whether x < 0, x == 0, or x > 0;
// x must be numeric or Unknown. For complex values x, the sign is 0 if x == 0,
// otherwise it is != 0. If x is Unknown, the result is 1.
func Sign(x Value) int

// Bytes returns the bytes for the absolute value of x in little-
// endian binary representation; x must be an Int.
func Bytes(x Value) []byte

// MakeFromBytes returns the Int value given the bytes of its little-endian
// binary representation. An empty byte slice argument represents 0.
func MakeFromBytes(bytes []byte) Value

// Num returns the numerator of x; x must be Int, Float, or Unknown.
// If x is Unknown, or if it is too large or small to represent as a
// fraction, the result is Unknown. Otherwise the result is an Int
// with the same sign as x.
func Num(x Value) Value

// Denom returns the denominator of x; x must be Int, Float, or Unknown.
// If x is Unknown, or if it is too large or small to represent as a
// fraction, the result is Unknown. Otherwise the result is an Int >= 1.
func Denom(x Value) Value

// MakeImag returns the Complex value x*i;
// x must be Int, Float, or Unknown.
// If x is Unknown, the result is Unknown.
func MakeImag(x Value) Value

// Real returns the real part of x, which must be a numeric or unknown value.
// If x is Unknown, the result is Unknown.
func Real(x Value) Value

// Imag returns the imaginary part of x, which must be a numeric or unknown value.
// If x is Unknown, the result is Unknown.
func Imag(x Value) Value

// ToInt converts x to an Int value if x is representable as an Int.
// Otherwise it returns an Unknown.
func ToInt(x Value) Value

// ToFloat converts x to a Float value if x is representable as a Float.
// Otherwise it returns an Unknown.
func ToFloat(x Value) Value

// ToComplex converts x to a Complex value if x is representable as a Complex.
// Otherwise it returns an Unknown.
func ToComplex(x Value) Value

// UnaryOp returns the result of the unary expression op y.
// The operation must be defined for the operand.
// If prec > 0 it specifies the ^ (xor) result size in bits.
// If y is Unknown, the result is Unknown.
func UnaryOp(op token.Token, y Value, prec uint) Value

// BinaryOp returns the result of the binary expression x op y.
// The operation must be defined for the operands. If one of the
// operands is Unknown, the result is Unknown.
// BinaryOp doesn't handle comparisons or shifts; use Compare
// or Shift instead.
//
// To force integer division of Int operands, use op == token.QUO_ASSIGN
// instead of token.QUO; the result is guaranteed to be Int in this case.
// Division by zero leads to a run-time panic.
func BinaryOp(x_ Value, op token.Token, y_ Value) Value

// Shift returns the result of the shift expression x op s
// with op == token.SHL or token.SHR (<< or >>). x must be
// an Int or an Unknown. If x is Unknown, the result is x.
func Shift(x Value, op token.Token, s uint) Value

// Compare returns the result of the comparison x op y.
// The comparison must be defined for the operands.
// If one of the operands is Unknown, the result is
// false.
func Compare(x_ Value, op token.Token, y_ Value) bool
