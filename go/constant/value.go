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

// Kindは [Value] が表す値の種類を指定します。
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

// MakeUnknownは [Unknown] の値を返します。
func MakeUnknown() Value

// MakeBoolはbの [Bool] 値を返します。
func MakeBool(b bool) Value

// MakeStringはsの [String] 値を返します。
func MakeString(s string) Value

// MakeInt64はxの [Int] 値を返します。
func MakeInt64(x int64) Value

// MakeUint64はxの [Int] 値を返します。
func MakeUint64(x uint64) Value

// MakeFloat64はxの [Float] 値を返します。
// xが-0.0の場合、結果は0.0です。
// xが有限ではない場合、結果はUnknownです。
func MakeFloat64(x float64) Value

// MakeFromLiteralは、Goリテラル文字列に対応する整数、浮動小数点、虚数、文字、または文字列の値を返します。tokの値は、[token.INT]、[token.FLOAT]、[token.IMAG]、[token.CHAR]、または [token.STRING] のいずれかでなければなりません。最後の引数はゼロでなければなりません。リテラルの文字列構文が無効な場合、結果は [Unknown] です。
func MakeFromLiteral(lit string, tok token.Token, zero uint) Value

// BoolValは、[Bool] または [Unknown] である必要があるxのGoのブール値を返します。
// xが [Unknown] の場合、結果はfalseです。
func BoolVal(x Value) bool

// StringValはxのGo文字列の値を返します。xは [String] または [Unknown] である必要があります。
// xが [Unknown] の場合、結果は""です。
func StringVal(x Value) string

// Int64Valは、xのGo int64値と結果が正確であるかどうかを返します。
// xは [Int] または [Unknown] でなければなりません。結果が正確でない場合は、値は未定義です。
// xがUnknownの場合、結果は(0、false)です。
func Int64Val(x Value) (int64, bool)

// Uint64ValはxのGo uint64の値と結果が正確かどうかを返します。
// xは [Int] または [Unknown] でなければなりません。結果が正確でない場合、その値は未定義です。
// xがUnknownの場合、結果は(0, false)です。
func Uint64Val(x Value) (uint64, bool)

// Float32Valは、float64ではなくfloat32のための [Float64Val] と同様です。
func Float32Val(x Value) (float32, bool)

// Float64Valは、xの最も近いGoのfloat64値とその結果が正確かどうかを返します。
// xは数値または [Unknown] である必要がありますが、[Complex] ではありません。float64として表現するのに
// 小さすぎる値（0に近すぎる）の場合、[Float64Val] は静かに0にアンダーフローします。結果の符号は常に
// xの符号と一致しますが、0の場合でもです。
// xが [Unknown] の場合、結果は（0、false）です。
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

// Makeはxの値に対する [Value] を返します。
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

// BitLenは、絶対値xを2進表現するために必要なビット数を返します。xは [Int] または [Unknown] である必要があります。
// xが [Unknown] の場合、結果は0です。
func BitLen(x Value) int

// Signは、xが0より小さい場合-1、xが0と等しい場合0、xが0より大きい場合1を返します。
// xは数値または [Unknown] である必要があります。複素数の場合、xが0と等しい場合は0であり、
// それ以外の場合は0ではありません。xが [Unknown] の場合、結果は1です。
func Sign(x Value) int

// Bytes関数はxの絶対値のバイトをリトルエンディアンの
// 2進表現で返します。xは [Int] 型である必要があります。
func Bytes(x Value) []byte

// MakeFromBytesは、リトルエンディアンのバイナリ表現のバイトを与えられた場合に、 [Int] 値を返します。空のバイトスライス引数は0を表します。
func MakeFromBytes(bytes []byte) Value

// Numはxの分子を返します。xは [Int]、[Float]、または [Unknown] でなければなりません。
// xが [Unknown] であるか、分数として表現するには大きすぎるまたは小さすぎる場合は、結果は [Unknown] です。
// それ以外の場合、結果はxと同じ符号の [Int] です。
func Num(x Value) Value

// Denomはxの分母を返します。xは [Int]、[Float]、または [Unknown] でなければなりません。
// もしxが [Unknown] であるか、それを分数として表現するのに大きすぎるか小さすぎる場合、結果は [Unknown] です。
// それ以外の場合、結果は1以上の [Int] です。
func Denom(x Value) Value

// MakeImagはComplex値 x*iを返します；
// xはInt、Float、またはUnknownである必要があります。
// もしxがUnknownの場合、結果もUnknownです。
func MakeImag(x Value) Value

// Realは、数値または未知の値でなければならないxの実数部を返します。
// xが [Unknown] の場合、結果は [Unknown] です。
func Real(x Value) Value

// Imagはxの虚数部分を返します。xは数値または不明な値である必要があります。
// xが不明な場合、結果は不明です。
func Imag(x Value) Value

// ToIntは、xが [Int] として表現可能な場合、xを [Int] 値に変換します。
// それ以外の場合は、[Unknown] を返します。
func ToInt(x Value) Value

// ToFloatはxが [Float] として表現可能な場合、xを [Float] 値に変換します。
// それ以外の場合は、[Unknown] を返します。
func ToFloat(x Value) Value

// ToComplexは、xが [Complex] として表現可能な場合は [Complex] の値に変換します。
// それ以外の場合は [Unknown] を返します。
func ToComplex(x Value) Value

// UnaryOpは単項演算子op yの結果を返します。
// 演算はオペランドに対して定義されている必要があります。
// prec > 0の場合、^（XOR）の結果のビット数を指定します。
// yが [Unknown] の場合、結果は [Unknown] です。
func UnaryOp(op token.Token, y Value, prec uint) Value

// BinaryOpは、バイナリ式x op yの結果を返します。
// 演算は、オペランドに対して定義されている必要があります。オペランドの1つが [Unknown] の場合、結果は [Unknown] です。
// BinaryOpは比較やシフトを処理しません。代わりに [Compare] または [Shift] を使用してください。
//
// Intオペランドの整数除算を強制するには、[token.QUO] のかわりにop == [token.QUO_ASSIGN] を使用します。この場合、結果は [Int] が保証されます。
// ゼロでの除算はランタイムパニックを引き起こします。
func BinaryOp(x_ Value, op token.Token, y_ Value) Value

// Shiftはshift式x op sの結果を返します
// opが [token.SHL] または [token.SHR]（<<または>>）である場合の結果です。xは
// [Int] または [Unknown] でなければなりません。xが [Unknown] の場合、結果はxです。
func Shift(x Value, op token.Token, s uint) Value

// Compareはx op yの比較結果を返します。
// オペランドの比較は定義されている必要があります。
// オペランドの一つが [Unknown] の場合、結果はfalseです。
func Compare(x_ Value, op token.Token, y_ Value) bool
