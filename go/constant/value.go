// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// constant パッケージは、Go の未型指定の定数とその対応する操作を表現する値を実装します。
//
// エラーにより値が不明な場合、特別な Unknown 値を使用できます。
// 不明な値に対する操作は、明示的に指定されない限り、不明な値を生成します。
package constant

import (
	"github.com/shogo82148/std/go/token"
)

<<<<<<< HEAD
// KindはValueが表す値の種類を指定します。
=======
// Kind specifies the kind of value represented by a [Value].
>>>>>>> upstream/master
type Kind int

const (
	// 未知の値
	Unknown Kind = iota

	// 数値ではない値
	Bool
	String

	// 数値の値
	Int
	Float
	Complex
)

// ValueはGo言語の定数の値を表します。
type Value interface {
	Kind() Kind

	String() string

	ExactString() string

	implementsValue()
}

<<<<<<< HEAD
// MakeUnknownはUnknownの値を返します。
func MakeUnknown() Value

// MakeBoolはbのBool値を返します。
func MakeBool(b bool) Value

// MakeStringはsのString値を返します。
func MakeString(s string) Value

// MakeInt64はxのInt値を返します。
func MakeInt64(x int64) Value

// MakeUint64はxのInt値を返します。
func MakeUint64(x uint64) Value

// MakeFloat64はxのFloat値を返します。
// xが-0.0の場合、結果は0.0です。
// xが有限ではない場合、結果はUnknownです。
func MakeFloat64(x float64) Value

// MakeFromLiteralは、Goリテラル文字列に対応する整数、浮動小数点、虚数、文字、または文字列の値を返します。tokの値は、token.INT、token.FLOAT、token.IMAG、token.CHAR、またはtoken.STRINGのいずれかでなければなりません。最後の引数はゼロでなければなりません。リテラルの文字列構文が無効な場合、結果はUnknownです。
func MakeFromLiteral(lit string, tok token.Token, zero uint) Value

// BoolValは、BoolまたはUnknownである必要があるxのGoのブール値を返します。
// xがUnknownの場合、結果はfalseです。
func BoolVal(x Value) bool

// StringValはxのGo文字列の値を返します。xはStringまたはUnknownである必要があります。
// xがUnknownの場合、結果は""です。
func StringVal(x Value) string

// Int64Valは、xのGo int64値と結果が正確であるかどうかを返します。
// xはIntまたはUnknownでなければなりません。結果が正確でない場合は、値は未定義です。
// xがUnknownの場合、結果は(0、偽)です。
func Int64Val(x Value) (int64, bool)

// Uint64ValはxのGo uint64の値と結果が正確かどうかを返します。
// xはIntまたはUnknownでなければなりません。結果が正確でない場合、その値は未定義です。
// xがUnknownの場合、結果は(0, false)です。
func Uint64Val(x Value) (uint64, bool)

// Float32Valは、float64ではなくfloat32のためのFloat64Valと同様です。
func Float32Val(x Value) (float32, bool)

// Float64Valは、xの最も近いGoのfloat64値とその結果が正確かどうかを返します。
// xは数値またはUnknownである必要がありますが、Complexではありません。float64として表現するのに
// 小さすぎる値（0に近すぎる）の場合、Float64Valは静かに0にアンダーフローします。結果の符号は常に
// xの符号と一致しますが、0の場合でもです。
// xがUnknownの場合、結果は（0、false）です。
=======
// MakeUnknown returns the [Unknown] value.
func MakeUnknown() Value

// MakeBool returns the [Bool] value for b.
func MakeBool(b bool) Value

// MakeString returns the [String] value for s.
func MakeString(s string) Value

// MakeInt64 returns the [Int] value for x.
func MakeInt64(x int64) Value

// MakeUint64 returns the [Int] value for x.
func MakeUint64(x uint64) Value

// MakeFloat64 returns the [Float] value for x.
// If x is -0.0, the result is 0.0.
// If x is not finite, the result is an [Unknown].
func MakeFloat64(x float64) Value

// MakeFromLiteral returns the corresponding integer, floating-point,
// imaginary, character, or string value for a Go literal string. The
// tok value must be one of [token.INT], [token.FLOAT], [token.IMAG],
// [token.CHAR], or [token.STRING]. The final argument must be zero.
// If the literal string syntax is invalid, the result is an [Unknown].
func MakeFromLiteral(lit string, tok token.Token, zero uint) Value

// BoolVal returns the Go boolean value of x, which must be a [Bool] or an [Unknown].
// If x is [Unknown], the result is false.
func BoolVal(x Value) bool

// StringVal returns the Go string value of x, which must be a [String] or an [Unknown].
// If x is [Unknown], the result is "".
func StringVal(x Value) string

// Int64Val returns the Go int64 value of x and whether the result is exact;
// x must be an [Int] or an [Unknown]. If the result is not exact, its value is undefined.
// If x is [Unknown], the result is (0, false).
func Int64Val(x Value) (int64, bool)

// Uint64Val returns the Go uint64 value of x and whether the result is exact;
// x must be an [Int] or an [Unknown]. If the result is not exact, its value is undefined.
// If x is [Unknown], the result is (0, false).
func Uint64Val(x Value) (uint64, bool)

// Float32Val is like [Float64Val] but for float32 instead of float64.
func Float32Val(x Value) (float32, bool)

// Float64Val returns the nearest Go float64 value of x and whether the result is exact;
// x must be numeric or an [Unknown], but not [Complex]. For values too small (too close to 0)
// to represent as float64, [Float64Val] silently underflows to 0. The result sign always
// matches the sign of x, even for 0.
// If x is [Unknown], the result is (0, false).
>>>>>>> upstream/master
func Float64Val(x Value) (float64, bool)

// Valは指定された定数の基になる値を返します。インターフェースを返すため、呼び出し元が結果を期待する型にキャストすることが求められます。可能な動的な戻り値の型は次の通りです：
//
//	xの種類            結果の型
//	-------------------------------------------
//	Bool               bool
//	String             string
//	Int                int64または*big.Int
//	Float              *big.Floatまたは*big.Rat
//	その他のすべて      nil
func Val(x Value) any

<<<<<<< HEAD
// Makeはxの値に対するValueを返します。
=======
// Make returns the [Value] for x.
>>>>>>> upstream/master
//
//	xの型            結果のKind
//	----------------------------
//	bool             Bool
//	string           String
//	int64            Int
//	*big.Int         Int
//	*big.Float       Float
//	*big.Rat         Float
//	それ以外の型     Unknown
func Make(x any) Value

<<<<<<< HEAD
// BitLenは、絶対値xを2進表現するために必要なビット数を返します。xはIntまたはUnknownである必要があります。
// xがUnknownの場合、結果は0です。
func BitLen(x Value) int

// Signは、xが0より小さい場合-1、xが0と等しい場合0、xが0より大きい場合1を返します。
// xは数値またはUnknownである必要があります。複素数の場合、xが0と等しい場合は0であり、
// それ以外の場合は0ではありません。xがUnknownの場合、結果は1です。
func Sign(x Value) int

// Bytes関数はxの絶対値のバイトをリトルエンディアンの
// 2進表現で返します。xはInt型である必要があります。
func Bytes(x Value) []byte

// MakeFromBytesは、リトルエンディアンのバイナリ表現のバイトを与えられた場合に、Int値を返します。空のバイトスライス引数は0を表します。
func MakeFromBytes(bytes []byte) Value

// Numはxの分子を返します。xはInt、Float、またはUnknownでなければなりません。
// xがUnknownであるか、分数として表現するには大きすぎるまたは小さすぎる場合は、結果はUnknownです。
// それ以外の場合、結果はxと同じ符号のIntです。
func Num(x Value) Value

// Denomはxの分母を返します。xはInt、Float、またはUnknownでなければなりません。
// もしxがUnknownであるか、それを分数として表現するのに大きすぎるか小さすぎる場合、結果はUnknownです。
// それ以外の場合、結果は1以上のIntです。
func Denom(x Value) Value

// MakeImagはComplex値 x*iを返します；
// xはInt、Float、またはUnknownである必要があります。
// もしxがUnknownの場合、結果もUnknownです。
func MakeImag(x Value) Value

// Realは、数値または未知の値でなければならないxの実数部を返します。
// xがUnknownの場合、結果はUnknownです。
func Real(x Value) Value

// Imagはxの虚数部分を返します。xは数値または不明な値である必要があります。
// xが不明な場合、結果は不明です。
func Imag(x Value) Value

// ToIntは、xがIntとして表現可能な場合、xをInt値に変換します。
// それ以外の場合は、Unknownを返します。
func ToInt(x Value) Value

// ToFloatはxがFloatとして表現可能な場合、xをFloat値に変換します。
// それ以外の場合は、Unknownを返します。
func ToFloat(x Value) Value

// ToComplexは、xがComplexとして表現可能な場合はComplexの値に変換します。
// それ以外の場合はUnknownを返します。
func ToComplex(x Value) Value

// UnaryOpは単項演算子op yの結果を返します。
// 演算はオペランドに対して定義されている必要があります。
// prec > 0の場合、^（XOR）の結果のビット数を指定します。
// yがUnknownの場合、結果はUnknownです。
func UnaryOp(op token.Token, y Value, prec uint) Value

// BinaryOpは、バイナリ式x op yの結果を返します。
// 演算は、オペランドに対して定義されている必要があります。オペランドの1つがUnknownの場合、結果はUnknownです。
// BinaryOpは比較やシフトを処理しません。代わりにCompareまたはShiftを使用してください。
//
// Intオペランドの整数除算を強制するには、token.QUOのかわりにop == token.QUO_ASSIGNを使用します。この場合、結果はIntが保証されます。
// ゼロでの除算はランタイムパニックを引き起こします。
func BinaryOp(x_ Value, op token.Token, y_ Value) Value

// Shiftはshift式x op sの結果を返します
// opがtoken.SHLまたはtoken.SHR（<<または>>）である場合の結果です。xは
// IntまたはUnknownでなければなりません。xがUnknownの場合、結果はxです。
func Shift(x Value, op token.Token, s uint) Value

// Compareはx op yの比較結果を返します。
// オペランドの比較は定義されている必要があります。
// オペランドの一つがUnknownの場合、結果はfalseです。
=======
// BitLen returns the number of bits required to represent
// the absolute value x in binary representation; x must be an [Int] or an [Unknown].
// If x is [Unknown], the result is 0.
func BitLen(x Value) int

// Sign returns -1, 0, or 1 depending on whether x < 0, x == 0, or x > 0;
// x must be numeric or [Unknown]. For complex values x, the sign is 0 if x == 0,
// otherwise it is != 0. If x is [Unknown], the result is 1.
func Sign(x Value) int

// Bytes returns the bytes for the absolute value of x in little-
// endian binary representation; x must be an [Int].
func Bytes(x Value) []byte

// MakeFromBytes returns the [Int] value given the bytes of its little-endian
// binary representation. An empty byte slice argument represents 0.
func MakeFromBytes(bytes []byte) Value

// Num returns the numerator of x; x must be [Int], [Float], or [Unknown].
// If x is [Unknown], or if it is too large or small to represent as a
// fraction, the result is [Unknown]. Otherwise the result is an [Int]
// with the same sign as x.
func Num(x Value) Value

// Denom returns the denominator of x; x must be [Int], [Float], or [Unknown].
// If x is [Unknown], or if it is too large or small to represent as a
// fraction, the result is [Unknown]. Otherwise the result is an [Int] >= 1.
func Denom(x Value) Value

// MakeImag returns the [Complex] value x*i;
// x must be [Int], [Float], or [Unknown].
// If x is [Unknown], the result is [Unknown].
func MakeImag(x Value) Value

// Real returns the real part of x, which must be a numeric or unknown value.
// If x is [Unknown], the result is [Unknown].
func Real(x Value) Value

// Imag returns the imaginary part of x, which must be a numeric or unknown value.
// If x is [Unknown], the result is [Unknown].
func Imag(x Value) Value

// ToInt converts x to an [Int] value if x is representable as an [Int].
// Otherwise it returns an [Unknown].
func ToInt(x Value) Value

// ToFloat converts x to a [Float] value if x is representable as a [Float].
// Otherwise it returns an [Unknown].
func ToFloat(x Value) Value

// ToComplex converts x to a [Complex] value if x is representable as a [Complex].
// Otherwise it returns an [Unknown].
func ToComplex(x Value) Value

// UnaryOp returns the result of the unary expression op y.
// The operation must be defined for the operand.
// If prec > 0 it specifies the ^ (xor) result size in bits.
// If y is [Unknown], the result is [Unknown].
func UnaryOp(op token.Token, y Value, prec uint) Value

// BinaryOp returns the result of the binary expression x op y.
// The operation must be defined for the operands. If one of the
// operands is [Unknown], the result is [Unknown].
// BinaryOp doesn't handle comparisons or shifts; use [Compare]
// or [Shift] instead.
//
// To force integer division of [Int] operands, use op == [token.QUO_ASSIGN]
// instead of [token.QUO]; the result is guaranteed to be [Int] in this case.
// Division by zero leads to a run-time panic.
func BinaryOp(x_ Value, op token.Token, y_ Value) Value

// Shift returns the result of the shift expression x op s
// with op == [token.SHL] or [token.SHR] (<< or >>). x must be
// an [Int] or an [Unknown]. If x is [Unknown], the result is x.
func Shift(x Value, op token.Token, s uint) Value

// Compare returns the result of the comparison x op y.
// The comparison must be defined for the operands.
// If one of the operands is [Unknown], the result is
// false.
>>>>>>> upstream/master
func Compare(x_ Value, op token.Token, y_ Value) bool
