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

// KindはValueが表す値の種類を指定します。
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
	// Kind returns the value kind.
	Kind() Kind

	// String returns a short, quoted (human-readable) form of the value.
	// For numeric values, the result may be an approximation;
	// for String values the result may be a shortened string.
	// Use ExactString for a string representing a value exactly.
	String() string

	// ExactString returns an exact, quoted (human-readable) form of the value.
	// If the Value is of Kind String, use StringVal to obtain the unquoted string.
	ExactString() string

	// Prevent external implementations.
	implementsValue()
}

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

// Makeはxの値に対するValueを返します。
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
func Compare(x_ Value, op token.Token, y_ Value) bool
