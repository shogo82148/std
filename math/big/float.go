// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements multi-precision floating-point numbers.
// Like in the GNU MPFR library (https://www.mpfr.org/), operands
// can be of mixed precision. Unlike MPFR, the rounding mode is
// not specified with each operation, but with each operand. The
// rounding mode of the result operand determines the rounding
// mode of an operation. This is a from-scratch implementation.

package big

import (
	"github.com/shogo82148/std/math"
)

// 非ゼロ有限Floatは、多精度浮動小数点数を表しますr
//
//	符号 × 仮数部 × 2**指数
//
// 0.5 <= 仮数部 < 1.0、および MinExp <= 指数 <= MaxExpとなります。
// Floatはゼロ（+0、-0）または無限（+Inf、-Inf）でもあり得ます。
// すべてのFloatは順序付けられており、二つのFloat xとyの順序付けは
// x.Cmp(y)によって定義されます。
//
// 各Float値には、精度、丸めモード、および精度もあります。
// 精度は、値を表現するために利用可能な仮数部ビットの最大数です。
// 丸めモードは、結果が仮数部ビットに収まるようにどのように丸められるべきかを指定します、
// 精度は、正確な結果に対する丸め誤差を説明します。
//
// 特に指定がない限り、結果として*Float変数を指定するすべての操作（セッターを含む）は、
// 通常レシーバを介して（[Float.MantExp] の例外を除く）、結果変数の精度と丸めモードに従って数値結果を丸めます。
//
// 提供された結果の精度が0（以下参照）の場合、それは丸めが行われる前に
// 最大の精度値を持つ引数の精度に設定され、丸めモードは変更されません。したがって、
// 結果の引数として提供される未初期化のFloatは、その精度がオペランドによって
// 決定される合理的な値に設定され、そのモードはRoundingModeのゼロ値（ToNearestEven）です。
//
// 望ましい精度を24または53に設定し、対応する丸めモード（通常は [ToNearestEven]）を使用すると、
// Float操作は、正常（つまり、非正規化ではない）float32またはfloat64数に対応するオペランドに対して、
// 対応するfloat32またはfloat64 IEEE 754算術と同じ結果を生成します。
// 指数のアンダーフローとオーバーフローは、Floatの指数がはるかに大きな範囲を持つため、
// IEEE 754とは異なる値に対して0またはInfinityを導きます。
//
// Floatのゼロ（未初期化）値は使用準備が整っており、
// 精度0と丸めモード [ToNearestEven] で数値+0.0を正確に表します。
//
// 操作は常にポインタ引数（*Float）を取るのではなく、
// Float値を取り、各一意のFloat値は自身の一意の*Floatポインタを必要とします。
// Float値を「コピー」するには、既存の（または新しく割り当てられた）Floatを
// [Float.Set] メソッドを使用して新しい値に設定する必要があります。
// Floatの浅いコピーはサポートされておらず、エラーを引き起こす可能性があります。
type Float struct {
	prec uint32
	mode RoundingMode
	acc  Accuracy
	form form
	neg  bool
	mant nat
	exp  int32
}

// ErrNaNパニックは、IEEE 754のルールに従ってNaNになる [Float] 操作によって引き起こされます。
// ErrNaNはエラーインターフェースを実装します。
type ErrNaN struct {
	msg string
}

func (err ErrNaN) Error() string

// NewFloatは、精度53と丸めモード [ToNearestEven] でxに設定された新しい [Float] を割り当てて返します。
// xがNaNの場合、NewFloatは [ErrNaN] でパニックを起こします。
func NewFloat(x float64) *Float

// 指数と精度の制限。
const (
	MaxExp  = math.MaxInt32
	MinExp  = math.MinInt32
	MaxPrec = math.MaxUint32
)

// RoundingModeは、[Float] 値が望ましい精度に丸められる方法を決定します。
// 丸めは [Float] 値を変更する可能性があり、丸め誤差は [Float] の [Accuracy] によって説明されます。
type RoundingMode byte

// これらの定数は、サポートされている丸めモードを定義します。
const (
	ToNearestEven RoundingMode = iota
	ToNearestAway
	ToZero
	AwayFromZero
	ToNegativeInf
	ToPositiveInf
)

// Accuracyは、[Float] 値を生成した最新の操作によって生じた丸め誤差を、
// 正確な値に対して説明します。
type Accuracy int8

// [Float] の [Accuracy] を説明する定数。
const (
	Below Accuracy = -1
	Exact Accuracy = 0
	Above Accuracy = +1
)

// SetPrecはzの精度をprecに設定し、（可能な場合）zの丸められた
// 値を返します。仮数部が精度の損失なしにprecビットで表現できない場合、
// zの丸めモードに従って丸めが行われます。
// SetPrec(0)はすべての有限値を±0にマップします；無限値は変更されません。
// prec > [MaxPrec] の場合、precは [MaxPrec] に設定されます。
func (z *Float) SetPrec(prec uint) *Float

// SetModeはzの丸めモードをmodeに設定し、正確なzを返します。
// それ以外の場合、zは変更されません。
// z.SetMode(z.Mode())は、zの精度を [Exact] に設定するための安価な方法です。
func (z *Float) SetMode(mode RoundingMode) *Float

// Precは、xの仮数部の精度をビット単位で返します。
// 結果は、|x| == 0 および |x| == Inf の場合、0になる可能性があります。
func (x *Float) Prec() uint

// MinPrecは、xを正確に表現するために必要な最小精度を返します
// （つまり、x.SetPrec(prec)がxを丸め始める最小のprec）。
// 結果は、|x| == 0 および |x| == Inf の場合、0になります。
func (x *Float) MinPrec() uint

// Modeは、xの丸めモードを返します。
func (x *Float) Mode() RoundingMode

// Accは、最も最近の操作によって生成されたxの精度を返します。
// その操作が明示的に異なることを文書化していない限り。
func (x *Float) Acc() Accuracy

// Signは以下を返します:
//
//	-1 は x <   0 の場合
//	 0 は x が ±0 の場合
//	+1 は x >   0 の場合
func (x *Float) Sign() int

// MantExpはxをその仮数部と指数部に分解し、指数を返します。
// 非nilのmant引数が提供された場合、その値はxの仮数部に設定され、
// xと同じ精度と丸めモードを持ちます。コンポーネントは
// x == mant × 2**exp、0.5 <= |mant| < 1.0を満たします。
// nil引数でMantExpを呼び出すことは、レシーバの指数を効率的に取得する方法です。
//
// 特殊なケースは以下の通りです:
//
//	(  ±0).MantExp(mant) = 0、mantは   ±0に設定されます
//	(±Inf).MantExp(mant) = 0、mantは ±Infに設定されます
//
// xとmantは同じものである可能性があり、その場合、xはその
// 仮数部の値に設定されます。
func (x *Float) MantExp(mant *Float) (exp int)

// SetMantExpはzをmant × 2**expに設定し、zを返します。
// 結果のzは、mantと同じ精度と丸めモードを持ちます。
// SetMantExpは [Float.MantExp] の逆ですが、0.5 <= |mant| < 1.0を必要としません。
// 特に、*[Float] 型の指定されたxに対して、SetMantExpは [Float.MantExp] と次のように関連しています:
//
//	mant := new(Float)
//	new(Float).SetMantExp(mant, x.MantExp(mant)).Cmp(x) == 0
//
// 特殊なケースは以下の通りです:
//
//	z.SetMantExp(  ±0, exp) =   ±0
//	z.SetMantExp(±Inf, exp) = ±Inf
//
// zとmantは同じものである可能性があり、その場合、zの指数はexpに設定されます。
func (z *Float) SetMantExp(mant *Float, exp int) *Float

// Signbitは、xが負または負のゼロであるかどうかを報告します。
func (x *Float) Signbit() bool

// IsInfは、xが+Infまたは-Infであるかどうかを報告します。
func (x *Float) IsInf() bool

// IsIntは、xが整数であるかどうかを報告します。
// ±Infの値は整数ではありません。
func (x *Float) IsInt() bool

// SetUint64は、zをxの（可能性のある丸められた）値に設定し、zを返します。
// zの精度が0の場合、それは64に変更されます（そして丸めは影響を及ぼしません）。
func (z *Float) SetUint64(x uint64) *Float

// SetInt64は、zをxの（可能性のある丸められた）値に設定し、zを返します。
// zの精度が0の場合、それは64に変更されます（そして丸めは影響を及ぼしません）。
func (z *Float) SetInt64(x int64) *Float

// SetFloat64は、zをxの（可能性のある丸められた）値に設定し、zを返します。
// zの精度が0の場合、それは53に変更されます（そして丸めは影響を及ぼしません）。
// xがNaNの場合、SetFloat64は [ErrNaN] でパニックを起こします。
func (z *Float) SetFloat64(x float64) *Float

// SetIntは、zをxの（可能性のある丸められた）値に設定し、zを返します。
// zの精度が0の場合、それはx.BitLen()または64の大きい方に変更されます
// （そして丸めは影響を及ぼしません）。
func (z *Float) SetInt(x *Int) *Float

// SetRatは、zをxの（可能性のある丸められた）値に設定し、zを返します。
// zの精度が0の場合、それはa.BitLen()、b.BitLen()、または64の最大のものに変更されます；
// x = a/bとします。
func (z *Float) SetRat(x *Rat) *Float

// SetInfは、signbitが設定されている場合はzを無限のFloat -Infに、
// 設定されていない場合は+Infに設定し、zを返します。
// zの精度は変わらず、結果は常に [Exact] です。
func (z *Float) SetInf(signbit bool) *Float

// Setは、zをxの（可能性のある丸められた）値に設定し、zを返します。
// zの精度が0の場合、zを設定する前にxの精度に変更されます
// （そして丸めは影響を及ぼしません）。
// 丸めはzの精度と丸めモードに従って実行され、
// zの精度は正確な（丸められていない）結果に対する結果のエラーを報告します。
func (z *Float) Set(x *Float) *Float

// Copyはzをxと同じ精度、丸めモード、およびxと同じ精度で設定します。
// Copyはzを返します。xとzが同一である場合、Copyは何も操作しません。
func (z *Float) Copy(x *Float) *Float

// Uint64は、xをゼロに向かって切り捨てることによって得られる符号なし整数を返します。
// 0 <= x <= math.MaxUint64の場合、結果はxが整数の場合は [Exact] 、それ以外の場合は [Below] です。
// x < 0の場合、結果は(0, [Above])で、x > [math.MaxUint64] の場合は([math.MaxUint64], [Below])です。
func (x *Float) Uint64() (uint64, Accuracy)

// Int64は、xをゼロに向かって切り捨てることによって得られる整数を返します。
// [math.MinInt64] <= x <= [math.MaxInt64] の場合、結果はxが整数の場合は [Exact]、それ以外の場合は [Above]（x < 0）または [Below]（x > 0）です。
// 結果はx < [math.MinInt64] の場合は（[math.MinInt64], Above）、x > [math.MaxInt64] の場合は（[math.MaxInt64], [Below]）です。
func (x *Float) Int64() (int64, Accuracy)

// Float32は、xに最も近いfloat32の値を返します。xが小さすぎて
// float32で表現できない場合（|x| < [math.SmallestNonzeroFloat32] ）、結果は
// （0, [Below]）または（-0, [Above]）となります。これはxの符号によります。
// xが大きすぎてfloat32で表現できない場合（|x| > [math.MaxFloat32]）、
// 結果は（+Inf, [Above]）または（-Inf, [Below]）となります。これもxの符号によります。
func (x *Float) Float32() (float32, Accuracy)

// Float64は、xに最も近いfloat64の値を返します。xが小さすぎて
// float64で表現できない場合（|x| < [math.SmallestNonzeroFloat64]）、結果は
// （0, [Below]）または（-0, [Above]）となります。これはxの符号によります。
// xが大きすぎてfloat64で表現できない場合（|x| > [math.MaxFloat64]）、
// 結果は（+Inf, [Above]）または（-Inf, [Below]）となります。これもxの符号によります。
func (x *Float) Float64() (float64, Accuracy)

// Intは、xをゼロに向かって切り捨てた結果を返します。
// または、xが無限大の場合はnilを返します。
// 結果はx.IsInt()の場合は [Exact]、それ以外の場合はx > 0の場合は [Below]、
// x < 0の場合は [Above] です。
// 非nilの*[Int] 引数zが提供された場合、[Int] は結果をzに格納します。
// 新しい [Int] を割り当てる代わりに。
func (x *Float) Int(z *Int) (*Int, Accuracy)

// Ratは、xに対応する有理数を返します。
// または、xが無限大の場合はnilを返します。
// 結果はxがInfでない場合は [Exact] です。
// 非nilの*[Rat] 引数zが提供された場合、[Rat] は結果をzに格納します。
// 新しい[Rat] を割り当てる代わりに。
func (x *Float) Rat(z *Rat) (*Rat, Accuracy)

// Absは、zを|x|（xの絶対値）の（可能性のある丸められた）値に設定し、zを返します。
func (z *Float) Abs(x *Float) *Float

// Negは、zを符号を反転したxの（可能性のある丸められた）値に設定し、zを返します。
func (z *Float) Neg(x *Float) *Float

// Addは、zを丸められた和x+yに設定し、zを返します。zの精度が0の場合、
// 操作前にxの精度またはyの精度の大きい方に変更されます。
// 丸めはzの精度と丸めモードに従って行われ、
// zの精度は正確な（丸められていない）結果に対する結果のエラーを報告します。
// xとyが逆の符号の無限大である場合、Addは [ErrNaN] でパニックを起こします。
// その場合、zの値は未定義です。
func (z *Float) Add(x, y *Float) *Float

// Subは、zを丸められた差分x-yに設定し、zを返します。
// 精度、丸め、および精度報告は [Float.Add] と同様です。
// xとyが同じ符号の無限大である場合、Subは [ErrNaN] でパニックを起こします。
// その場合、zの値は未定義です。
func (z *Float) Sub(x, y *Float) *Float

// Mulは、zを丸められた積x*yに設定し、zを返します。
// 精度、丸め、および精度報告は [Float.Add] と同様です。
// 一方のオペランドがゼロで、他方のオペランドが無限大である場合、Mulは [ErrNaN] でパニックを起こします。
// その場合、zの値は未定義です。
func (z *Float) Mul(x, y *Float) *Float

// Quoは、zを丸められた商x/yに設定し、zを返します。
// 精度、丸め、および精度報告は [Float.Add] と同様です。
// 両方のオペランドがゼロまたは無限大である場合、Quoは [ErrNaN] でパニックを起こします。
// その場合、zの値は未定義です。
func (z *Float) Quo(x, y *Float) *Float

// Cmpはxとyを比較し、次の値を返します:
//
//	-1 は x <  y
//	 0 は x == y (これには -0 == 0, -Inf == -Inf, そして +Inf == +Inf も含まれます)
//	+1 は x >  y
func (x *Float) Cmp(y *Float) int
