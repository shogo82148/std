// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// randパッケージは、シミュレーションなどのタスクに適した擬似乱数生成器を実装しますが、セキュリティに敏感な作業には使用しないでください。
//
// 乱数は [Source] によって生成され、通常は [Rand] でラップされます。
// 両方のタイプは一度に1つのゴルーチンから使用されるべきです：複数のゴルーチン間で共有するには何らかの同期が必要です。
//
// トップレベルの関数、例えば [Float64] や [Int] は、
// 複数のゴルーチンによる並行使用が安全です。
//
// このパッケージの出力は、どのようにシードされていても容易に予測可能かもしれません。セキュリティに敏感な作業に適した乱数については、
// [crypto/rand] パッケージを参照してください。
package rand

// Sourceは、範囲[0, 1<<64)内の一様に分布した
// 擬似乱数uint64値のソースです。
//
// Sourceは、複数のゴルーチンによる並行使用には安全ではありません。
type Source interface {
	Uint64() uint64
}

// Randは、乱数のソースです。
type Rand struct {
	src Source
}

// Newは、他の乱数を生成するためにsrcから乱数を使用する新しいRandを返します。
func New(src Source) *Rand

// Int64は、非負の擬似乱数63ビット整数をint64として返します。
func (r *Rand) Int64() int64

// Uint32は、擬似乱数32ビット値をuint32として返します。
func (r *Rand) Uint32() uint32

// Uint64は、擬似乱数64ビット値をuint64として返します。
func (r *Rand) Uint64() uint64

// Int32は、非負の擬似乱数31ビット整数をint32として返します。
func (r *Rand) Int32() int32

// Intは、非負の擬似乱数intを返します。
func (r *Rand) Int() int

// Uint は擬似乱数の uint を返します。
func (r *Rand) Uint() uint

// Int64Nは、半開放区間[0,n)内の非負の擬似乱数をint64として返します。
// nが0以下の場合、パニックを引き起こします。
func (r *Rand) Int64N(n int64) int64

// Uint64Nは、半開放区間[0,n)内の非負の擬似乱数をuint64として返します。
// nが0の場合、パニックを引き起こします。
func (r *Rand) Uint64N(n uint64) uint64

// Int32Nは、半開放区間[0,n)内の非負の擬似乱数をint32として返します。
// nが0以下の場合、パニックを引き起こします。
func (r *Rand) Int32N(n int32) int32

// Uint32Nは、半開放区間[0,n)内の非負の擬似乱数をuint32として返します。
// nが0の場合、パニックを引き起こします。
func (r *Rand) Uint32N(n uint32) uint32

// IntNは、半開放区間[0,n)内の非負の擬似乱数をintとして返します。
// nが0以下の場合、パニックを引き起こします。
func (r *Rand) IntN(n int) int

// UintNは、半開放区間[0,n)内の非負の擬似乱数をuintとして返します。
// nが0の場合、パニックを引き起こします。
func (r *Rand) UintN(n uint) uint

// Float64は、半開放区間[0.0,1.0)内の擬似乱数をfloat64として返します。
func (r *Rand) Float64() float64

// Float32は、半開放区間[0.0,1.0)内の擬似乱数をfloat32として返します。
func (r *Rand) Float32() float32

// Permは、半開放区間[0,n)内の整数の擬似乱数順列をn個のintのスライスとして返します。
func (r *Rand) Perm(n int) []int

// Shuffleは要素の順序を擬似ランダムにします。
// nは要素の数です。n < 0の場合、Shuffleはパニックを引き起こします。
// swapは、インデックスiとjの要素を交換します。
func (r *Rand) Shuffle(n int, swap func(i, j int))

// Int64は、デフォルトのSourceから非負の擬似乱数63ビット整数をint64として返します。
func Int64() int64

// Uint32は、デフォルトのSourceから擬似乱数32ビット値をuint32として返します。
func Uint32() uint32

// Uint64Nは、デフォルトのSourceから半開区間[0,n)内の擬似乱数をuint64として返します。
// nが0の場合、パニックを引き起こします。
func Uint64N(n uint64) uint64

// Uint32Nは、デフォルトのSourceから半開区間[0,n)内の擬似乱数をuint32として返します。
// nが0の場合、パニックを引き起こします。
func Uint32N(n uint32) uint32

// Uint64は、デフォルトのSourceから擬似乱数64ビット値をuint64として返します。
func Uint64() uint64

// Int32は、デフォルトのSourceから非負の擬似乱数31ビット整数をint32として返します。
func Int32() int32

// Intは、デフォルトのSourceから非負の擬似乱数intを返します。
func Int() int

// Uint はデフォルトのソースから擬似乱数の uint を返します。
func Uint() uint

// Int64Nは、デフォルトのSourceから半開放区間[0,n)内の非負の擬似乱数をint64として返します。
// nが0以下の場合、パニックを引き起こします。
func Int64N(n int64) int64

// Int32Nは、デフォルトのSourceから半開放区間[0,n)内の非負の擬似乱数をint32として返します。
// nが0以下の場合、パニックを引き起こします。
func Int32N(n int32) int32

// IntNは、デフォルトのSourceから半開放区間[0,n)内の非負の擬似乱数をintとして返します。
// nが0以下の場合、パニックを引き起こします。
func IntN(n int) int

// UintNは、デフォルトのSourceから半開区間[0,n)内の擬似乱数をuintとして返します。
// nが0の場合、パニックを引き起こします。
func UintN(n uint) uint

// Nは、デフォルトのSourceから半開放区間[0,n)内の擬似乱数を返します。
// 型パラメータIntは任意の整数型にすることができます。
// nが0以下の場合、パニックを引き起こします。
func N[Int intType](n Int) Int

// Float64は、デフォルトのSourceから半開放区間[0.0,1.0)内の擬似乱数をfloat64として返します。
func Float64() float64

// Float32は、デフォルトのSourceから半開放区間[0.0,1.0)内の擬似乱数をfloat32として返します。
func Float32() float32

// Permは、デフォルトのSourceから半開放区間[0,n)内の整数の擬似乱数順列をn個のintのスライスとして返します。
func Perm(n int) []int

// ShuffleはデフォルトのSourceを使用して要素の順序を擬似ランダムにします。
// nは要素の数です。n < 0の場合、Shuffleはパニックを引き起こします。
// swapは、インデックスiとjの要素を交換します。
func Shuffle(n int, swap func(i, j int))

// NormFloat64は、デフォルトのSourceから標準正規分布（平均 = 0、標準偏差 = 1）に従う
// 範囲[-math.MaxFloat64, +math.MaxFloat64]の正規分布のfloat64を返します。
// 異なる正規分布を生成するために、呼び出し元は出力を調整できます：
//
//	sample = NormFloat64() * desiredStdDev + desiredMean
func NormFloat64() float64

// ExpFloat64は、デフォルトのSourceからレートパラメータ（lambda）が1で平均が1/lambda（1）の指数分布に従う
// 範囲(0, +math.MaxFloat64]の指数分布のfloat64を返します。
// 異なるレートパラメータの分布を生成するために、呼び出し元は出力を調整できます：
//
//	sample = ExpFloat64() / desiredRateParameter
func ExpFloat64() float64
