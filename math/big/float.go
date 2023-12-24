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
<<<<<<< HEAD
// Unless specified otherwise, all operations (including setters) that
// specify a *Float variable for the result (usually via the receiver
// with the exception of [Float.MantExp]), round the numeric result according
// to the precision and rounding mode of the result variable.
=======
// 特に指定がない限り、結果として*Float変数を指定するすべての操作（セッターを含む）は、
// 通常レシーバを介して（MantExpの例外を除く）、結果変数の精度と丸めモードに従って数値結果を丸めます。
>>>>>>> release-branch.go1.21
//
// 提供された結果の精度が0（以下参照）の場合、それは丸めが行われる前に
// 最大の精度値を持つ引数の精度に設定され、丸めモードは変更されません。したがって、
// 結果の引数として提供される未初期化のFloatは、その精度がオペランドによって
// 決定される合理的な値に設定され、そのモードはRoundingModeのゼロ値（ToNearestEven）です。
//
<<<<<<< HEAD
// By setting the desired precision to 24 or 53 and using matching rounding
// mode (typically [ToNearestEven]), Float operations produce the same results
// as the corresponding float32 or float64 IEEE-754 arithmetic for operands
// that correspond to normal (i.e., not denormal) float32 or float64 numbers.
// Exponent underflow and overflow lead to a 0 or an Infinity for different
// values than IEEE-754 because Float exponents have a much larger range.
//
// The zero (uninitialized) value for a Float is ready to use and represents
// the number +0.0 exactly, with precision 0 and rounding mode [ToNearestEven].
//
// Operations always take pointer arguments (*Float) rather
// than Float values, and each unique Float value requires
// its own unique *Float pointer. To "copy" a Float value,
// an existing (or newly allocated) Float must be set to
// a new value using the [Float.Set] method; shallow copies
// of Floats are not supported and may lead to errors.
=======
// 望ましい精度を24または53に設定し、対応する丸めモード（通常はToNearestEven）を使用すると、
// Float操作は、正常（つまり、非正規化ではない）float32またはfloat64数に対応するオペランドに対して、
// 対応するfloat32またはfloat64 IEEE-754算術と同じ結果を生成します。
// 指数のアンダーフローとオーバーフローは、Floatの指数がはるかに大きな範囲を持つため、
// IEEE-754とは異なる値に対して0またはInfinityを導きます。
//
// Floatのゼロ（未初期化）値は使用準備が整っており、
// 精度0と丸めモードToNearestEvenで数値+0.0を正確に表します。
//
// 操作は常にポインタ引数（*Float）を取るのではなく、
// Float値を取り、各一意のFloat値は自身の一意の*Floatポインタを必要とします。
// Float値を「コピー」するには、既存の（または新しく割り当てられた）Floatを
// Float.Setメソッドを使用して新しい値に設定する必要があります。
// Floatの浅いコピーはサポートされておらず、エラーを引き起こす可能性があります。
>>>>>>> release-branch.go1.21
type Float struct {
	prec uint32
	mode RoundingMode
	acc  Accuracy
	form form
	neg  bool
	mant nat
	exp  int32
}

<<<<<<< HEAD
// An ErrNaN panic is raised by a [Float] operation that would lead to
// a NaN under IEEE-754 rules. An ErrNaN implements the error interface.
=======
// ErrNaNパニックは、IEEE-754のルールに従ってNaNになるFloat操作によって引き起こされます。
// ErrNaNはエラーインターフェースを実装します。
>>>>>>> release-branch.go1.21
type ErrNaN struct {
	msg string
}

func (err ErrNaN) Error() string

<<<<<<< HEAD
// NewFloat allocates and returns a new [Float] set to x,
// with precision 53 and rounding mode [ToNearestEven].
// NewFloat panics with [ErrNaN] if x is a NaN.
=======
// NewFloatは、精度53と丸めモードToNearestEvenでxに設定された新しいFloatを割り当てて返します。
// xがNaNの場合、NewFloatはErrNaNでパニックを起こします。
>>>>>>> release-branch.go1.21
func NewFloat(x float64) *Float

// 指数と精度の制限。
const (
	MaxExp  = math.MaxInt32
	MinExp  = math.MinInt32
	MaxPrec = math.MaxUint32
)

<<<<<<< HEAD
// RoundingMode determines how a [Float] value is rounded to the
// desired precision. Rounding may change the [Float] value; the
// rounding error is described by the [Float]'s [Accuracy].
=======
// RoundingModeは、Float値が望ましい精度に丸められる方法を決定します。
// 丸めはFloat値を変更する可能性があり、丸め誤差はFloatのAccuracyによって説明されます。
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// Accuracy describes the rounding error produced by the most recent
// operation that generated a [Float] value, relative to the exact value.
type Accuracy int8

// Constants describing the [Accuracy] of a [Float].
=======
// Accuracyは、Float値を生成した最新の操作によって生じた丸め誤差を、
// 正確な値に対して説明します。
type Accuracy int8

// Floatの精度を説明する定数。
>>>>>>> release-branch.go1.21
const (
	Below Accuracy = -1
	Exact Accuracy = 0
	Above Accuracy = +1
)

<<<<<<< HEAD
// SetPrec sets z's precision to prec and returns the (possibly) rounded
// value of z. Rounding occurs according to z's rounding mode if the mantissa
// cannot be represented in prec bits without loss of precision.
// SetPrec(0) maps all finite values to ±0; infinite values remain unchanged.
// If prec > [MaxPrec], it is set to [MaxPrec].
func (z *Float) SetPrec(prec uint) *Float

// SetMode sets z's rounding mode to mode and returns an exact z.
// z remains unchanged otherwise.
// z.SetMode(z.Mode()) is a cheap way to set z's accuracy to [Exact].
=======
// SetPrecはzの精度をprecに設定し、（可能な場合）zの丸められた
// 値を返します。仮数部が精度の損失なしにprecビットで表現できない場合、
// zの丸めモードに従って丸めが行われます。
// SetPrec(0)はすべての有限値を±0にマップします；無限値は変更されません。
// prec > MaxPrecの場合、precはMaxPrecに設定されます。
func (z *Float) SetPrec(prec uint) *Float

// SetModeはzの丸めモードをmodeに設定し、正確なzを返します。
// それ以外の場合、zは変更されません。
// z.SetMode(z.Mode())は、zの精度をExactに設定するための安価な方法です。
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// SetMantExp sets z to mant × 2**exp and returns z.
// The result z has the same precision and rounding mode
// as mant. SetMantExp is an inverse of [Float.MantExp] but does
// not require 0.5 <= |mant| < 1.0. Specifically, for a
// given x of type *[Float], SetMantExp relates to [Float.MantExp]
// as follows:
=======
// SetMantExpはzをmant × 2**expに設定し、zを返します。
// 結果のzは、mantと同じ精度と丸めモードを持ちます。
// SetMantExpはMantExpの逆ですが、0.5 <= |mant| < 1.0を必要としません。
// 特に、*Float型の指定されたxに対して、SetMantExpはMantExpと次のように関連しています:
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// SetFloat64 sets z to the (possibly rounded) value of x and returns z.
// If z's precision is 0, it is changed to 53 (and rounding will have
// no effect). SetFloat64 panics with [ErrNaN] if x is a NaN.
=======
// SetFloat64は、zをxの（可能性のある丸められた）値に設定し、zを返します。
// zの精度が0の場合、それは53に変更されます（そして丸めは影響を及ぼしません）。
// xがNaNの場合、SetFloat64はErrNaNでパニックを起こします。
>>>>>>> release-branch.go1.21
func (z *Float) SetFloat64(x float64) *Float

// SetIntは、zをxの（可能性のある丸められた）値に設定し、zを返します。
// zの精度が0の場合、それはx.BitLen()または64の大きい方に変更されます
// （そして丸めは影響を及ぼしません）。
func (z *Float) SetInt(x *Int) *Float

// SetRatは、zをxの（可能性のある丸められた）値に設定し、zを返します。
// zの精度が0の場合、それはa.BitLen()、b.BitLen()、または64の最大のものに変更されます；
// x = a/bとします。
func (z *Float) SetRat(x *Rat) *Float

<<<<<<< HEAD
// SetInf sets z to the infinite Float -Inf if signbit is
// set, or +Inf if signbit is not set, and returns z. The
// precision of z is unchanged and the result is always
// [Exact].
=======
// SetInfは、signbitが設定されている場合はzを無限のFloat -Infに、
// 設定されていない場合は+Infに設定し、zを返します。
// zの精度は変わらず、結果は常にExactです。
>>>>>>> release-branch.go1.21
func (z *Float) SetInf(signbit bool) *Float

// Setは、zをxの（可能性のある丸められた）値に設定し、zを返します。
// zの精度が0の場合、zを設定する前にxの精度に変更されます
// （そして丸めは影響を及ぼしません）。
// 丸めはzの精度と丸めモードに従って実行され、
// zの精度は正確な（丸められていない）結果に対する結果のエラーを報告します。
func (z *Float) Set(x *Float) *Float

// Copyは、zをxと同じ精度、丸めモード、および
// 精度で設定し、zを返します。zと
// xが同じであっても、xは変更されません。
func (z *Float) Copy(x *Float) *Float

<<<<<<< HEAD
// Uint64 returns the unsigned integer resulting from truncating x
// towards zero. If 0 <= x <= math.MaxUint64, the result is [Exact]
// if x is an integer and [Below] otherwise.
// The result is (0, [Above]) for x < 0, and ([math.MaxUint64], [Below])
// for x > [math.MaxUint64].
func (x *Float) Uint64() (uint64, Accuracy)

// Int64 returns the integer resulting from truncating x towards zero.
// If [math.MinInt64] <= x <= [math.MaxInt64], the result is [Exact] if x is
// an integer, and [Above] (x < 0) or [Below] (x > 0) otherwise.
// The result is ([math.MinInt64], [Above]) for x < [math.MinInt64],
// and ([math.MaxInt64], [Below]) for x > [math.MaxInt64].
func (x *Float) Int64() (int64, Accuracy)

// Float32 returns the float32 value nearest to x. If x is too small to be
// represented by a float32 (|x| < [math.SmallestNonzeroFloat32]), the result
// is (0, [Below]) or (-0, [Above]), respectively, depending on the sign of x.
// If x is too large to be represented by a float32 (|x| > [math.MaxFloat32]),
// the result is (+Inf, [Above]) or (-Inf, [Below]), depending on the sign of x.
func (x *Float) Float32() (float32, Accuracy)

// Float64 returns the float64 value nearest to x. If x is too small to be
// represented by a float64 (|x| < [math.SmallestNonzeroFloat64]), the result
// is (0, [Below]) or (-0, [Above]), respectively, depending on the sign of x.
// If x is too large to be represented by a float64 (|x| > [math.MaxFloat64]),
// the result is (+Inf, [Above]) or (-Inf, [Below]), depending on the sign of x.
func (x *Float) Float64() (float64, Accuracy)

// Int returns the result of truncating x towards zero;
// or nil if x is an infinity.
// The result is [Exact] if x.IsInt(); otherwise it is [Below]
// for x > 0, and [Above] for x < 0.
// If a non-nil *[Int] argument z is provided, [Int] stores
// the result in z instead of allocating a new [Int].
func (x *Float) Int(z *Int) (*Int, Accuracy)

// Rat returns the rational number corresponding to x;
// or nil if x is an infinity.
// The result is [Exact] if x is not an Inf.
// If a non-nil *[Rat] argument z is provided, [Rat] stores
// the result in z instead of allocating a new [Rat].
=======
// Uint64は、xをゼロに向かって切り捨てることによって得られる符号なし整数を返します。
// 0 <= x <= math.MaxUint64の場合、結果はxが整数の場合はExact、それ以外の場合はBelowです。
// x < 0の場合、結果は(0, Above)で、x > math.MaxUint64の場合は(math.MaxUint64, Below)です。
func (x *Float) Uint64() (uint64, Accuracy)

// Int64は、xをゼロに向かって切り捨てることによって得られる整数を返します。
// math.MinInt64 <= x <= math.MaxInt64の場合、結果はxが整数の場合はExact、それ以外の場合はAbove（x < 0）またはBelow（x > 0）です。
// 結果はx < math.MinInt64の場合は（math.MinInt64, Above）、x > math.MaxInt64の場合は（math.MaxInt64, Below）です。
func (x *Float) Int64() (int64, Accuracy)

// Float32は、xに最も近いfloat32の値を返します。xが小さすぎて
// float32で表現できない場合（|x| < math.SmallestNonzeroFloat32）、結果は
// （0, Below）または（-0, Above）となります。これはxの符号によります。
// xが大きすぎてfloat32で表現できない場合（|x| > math.MaxFloat32）、
// 結果は（+Inf, Above）または（-Inf, Below）となります。これもxの符号によります。
func (x *Float) Float32() (float32, Accuracy)

// Float64は、xに最も近いfloat64の値を返します。xが小さすぎて
// float64で表現できない場合（|x| < math.SmallestNonzeroFloat64）、結果は
// （0, Below）または（-0, Above）となります。これはxの符号によります。
// xが大きすぎてfloat64で表現できない場合（|x| > math.MaxFloat64）、
// 結果は（+Inf, Above）または（-Inf, Below）となります。これもxの符号によります。
func (x *Float) Float64() (float64, Accuracy)

// Intは、xをゼロに向かって切り捨てた結果を返します。
// または、xが無限大の場合はnilを返します。
// 結果はx.IsInt()の場合はExact、それ以外の場合はx > 0の場合はBelow、
// x < 0の場合はAboveです。
// 非nilの*Int引数zが提供された場合、Intは結果をzに格納します。
// 新しいIntを割り当てる代わりに。
func (x *Float) Int(z *Int) (*Int, Accuracy)

// Ratは、xに対応する有理数を返します。
// または、xが無限大の場合はnilを返します。
// 結果はxがInfでない場合はExactです。
// 非nilの*Rat引数zが提供された場合、Ratは結果をzに格納します。
// 新しいRatを割り当てる代わりに。
>>>>>>> release-branch.go1.21
func (x *Float) Rat(z *Rat) (*Rat, Accuracy)

// Absは、zを|x|（xの絶対値）の（可能性のある丸められた）値に設定し、zを返します。
func (z *Float) Abs(x *Float) *Float

// Negは、zを符号を反転したxの（可能性のある丸められた）値に設定し、zを返します。
func (z *Float) Neg(x *Float) *Float

<<<<<<< HEAD
// Add sets z to the rounded sum x+y and returns z. If z's precision is 0,
// it is changed to the larger of x's or y's precision before the operation.
// Rounding is performed according to z's precision and rounding mode; and
// z's accuracy reports the result error relative to the exact (not rounded)
// result. Add panics with [ErrNaN] if x and y are infinities with opposite
// signs. The value of z is undefined in that case.
func (z *Float) Add(x, y *Float) *Float

// Sub sets z to the rounded difference x-y and returns z.
// Precision, rounding, and accuracy reporting are as for [Float.Add].
// Sub panics with [ErrNaN] if x and y are infinities with equal
// signs. The value of z is undefined in that case.
func (z *Float) Sub(x, y *Float) *Float

// Mul sets z to the rounded product x*y and returns z.
// Precision, rounding, and accuracy reporting are as for [Float.Add].
// Mul panics with [ErrNaN] if one operand is zero and the other
// operand an infinity. The value of z is undefined in that case.
func (z *Float) Mul(x, y *Float) *Float

// Quo sets z to the rounded quotient x/y and returns z.
// Precision, rounding, and accuracy reporting are as for [Float.Add].
// Quo panics with [ErrNaN] if both operands are zero or infinities.
// The value of z is undefined in that case.
=======
// Addは、zを丸められた和x+yに設定し、zを返します。zの精度が0の場合、
// 操作前にxの精度またはyの精度の大きい方に変更されます。
// 丸めはzの精度と丸めモードに従って行われ、
// zの精度は正確な（丸められていない）結果に対する結果のエラーを報告します。
// xとyが逆の符号の無限大である場合、AddはErrNaNでパニックを起こします。
// その場合、zの値は未定義です。
func (z *Float) Add(x, y *Float) *Float

// Subは、zを丸められた差分x-yに設定し、zを返します。
// 精度、丸め、および精度報告はAddと同様です。
// xとyが同じ符号の無限大である場合、SubはErrNaNでパニックを起こします。
// その場合、zの値は未定義です。
func (z *Float) Sub(x, y *Float) *Float

// Mulは、zを丸められた積x*yに設定し、zを返します。
// 精度、丸め、および精度報告はAddと同様です。
// 一方のオペランドがゼロで、他方のオペランドが無限大である場合、MulはErrNaNでパニックを起こします。
// その場合、zの値は未定義です。
func (z *Float) Mul(x, y *Float) *Float

// Quoは、zを丸められた商x/yに設定し、zを返します。
// 精度、丸め、および精度報告はAddと同様です。
// 両方のオペランドがゼロまたは無限大である場合、QuoはErrNaNでパニックを起こします。
// その場合、zの値は未定義です。
>>>>>>> release-branch.go1.21
func (z *Float) Quo(x, y *Float) *Float

// Cmpはxとyを比較し、次の値を返します:
//
//	-1 は x <  y
//	 0 は x == y (これには -0 == 0, -Inf == -Inf, そして +Inf == +Inf も含まれます)
//	+1 は x >  y
func (x *Float) Cmp(y *Float) int
