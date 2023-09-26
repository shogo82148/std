// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package constant implements Values representing untyped
// Go constants and the corresponding operations. Values
// and operations may have arbitrary or unlimited precision.
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

// A Value represents a mathematically exact value of a given Kind.
type Value interface {
	Kind() Kind

	String() string

	implementsValue()
}

// int64 bounds

// MakeUnknown returns the Unknown value.
func MakeUnknown() Value

// MakeBool returns the Bool value for x.
func MakeBool(b bool) Value

// MakeString returns the String value for x.
func MakeString(s string) Value

// MakeInt64 returns the Int value for x.
func MakeInt64(x int64) Value

// MakeUint64 returns the Int value for x.
func MakeUint64(x uint64) Value

// MakeFloat64 returns the numeric value for x.
// If x is not finite, the result is unknown.
func MakeFloat64(x float64) Value

// MakeFromLiteral returns the corresponding integer, floating-point,
// imaginary, character, or string value for a Go literal string.
// If prec > 0, prec specifies an upper limit for the precision of
// a numeric value. If the literal string is invalid, the result is
// nil.
// BUG(gri) Only prec == 0 is supported at the moment.
func MakeFromLiteral(lit string, tok token.Token, prec uint) Value

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
// x must be numeric but not Complex, or Unknown. For values too small (too close to 0)
// to represent as float64, Float64Val silently underflows to 0. The result sign always
// matches the sign of x, even for 0.
// If x is Unknown, the result is (0, false).
func Float64Val(x Value) (float64, bool)

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
// If x is Unknown, the result is Unknown, otherwise it is an Int
// with the same sign as x.
func Num(x Value) Value

// Denom returns the denominator of x; x must be Int, Float, or Unknown.
// If x is Unknown, the result is Unknown, otherwise it is an Int >= 1.
func Denom(x Value) Value

// MakeImag returns the numeric value x*i (possibly 0);
// x must be Int, Float, or Unknown.
// If x is Unknown, the result is Unknown.
func MakeImag(x Value) Value

// Real returns the real part of x, which must be a numeric or unknown value.
// If x is Unknown, the result is Unknown.
func Real(x Value) Value

// Imag returns the imaginary part of x, which must be a numeric or unknown value.
// If x is Unknown, the result is Unknown.
func Imag(x Value) Value

// UnaryOp returns the result of the unary expression op y.
// The operation must be defined for the operand.
// If prec > 0 it specifies the ^ (xor) result size in bits.
// If y is Unknown, the result is Unknown.
func UnaryOp(op token.Token, y Value, prec uint) Value

// BinaryOp returns the result of the binary expression x op y.
// The operation must be defined for the operands. If one of the
// operands is Unknown, the result is Unknown.
// To force integer division of Int operands, use op == token.QUO_ASSIGN
// instead of token.QUO; the result is guaranteed to be Int in this case.
// Division by zero leads to a run-time panic.
func BinaryOp(x Value, op token.Token, y Value) Value

// Shift returns the result of the shift expression x op s
// with op == token.SHL or token.SHR (<< or >>). x must be
// an Int or an Unknown. If x is Unknown, the result is x.
func Shift(x Value, op token.Token, s uint) Value

// Compare returns the result of the comparison x op y.
// The comparison must be defined for the operands.
// If one of the operands is Unknown, the result is
// false.
func Compare(x Value, op token.Token, y Value) bool
