// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージrandは、シミュレーションなどのタスクに適した擬似乱数生成器を実装しますが、
// セキュリティに敏感な作業には使用しないでください。
//
// 乱数は[Source]によって生成され、通常は [Rand] でラップされます。
// これらの型は一度に1つのゴルーチンで使用する必要があります：複数のゴルーチン間で共有するには何らかの同期が必要です。
//
// トップレベルの関数、たとえば [Float64] や [Int] などは、
// 複数のゴルーチンによる並行使用に対して安全です。
//
// このパッケージの出力は、どのようにシードされていても容易に予測可能かもしれません。
// セキュリティに敏感な作業に適したランダムな数値については、crypto/randパッケージを参照してください。
package rand

// Sourceは、範囲[0, 1<<63)内の一様に分布した
// 擬似乱数int64値のソースを表します。
//
// Sourceは、複数のゴルーチンによる並行使用には安全ではありません。
type Source interface {
	Int63() int64
	Seed(seed int64)
}

<<<<<<< HEAD
// A Source64 is a [Source] that can also generate
// uniformly-distributed pseudo-random uint64 values in
// the range [0, 1<<64) directly.
// If a [Rand] r's underlying [Source] s implements Source64,
// then r.Uint64 returns the result of one call to s.Uint64
// instead of making two calls to s.Int63.
=======
// Source64は、範囲[0, 1<<64)内の一様に分布した
// 擬似乱数uint64値を直接生成することもできるSourceです。
// Rand rの基礎となるSource sがSource64を実装している場合、
// r.Uint64はs.Int63を2回呼び出す代わりに、s.Uint64を1回呼び出した結果を返します。
>>>>>>> release-branch.go1.21
type Source64 interface {
	Source
	Uint64() uint64
}

<<<<<<< HEAD
// NewSource returns a new pseudo-random [Source] seeded with the given value.
// Unlike the default [Source] used by top-level functions, this source is not
// safe for concurrent use by multiple goroutines.
// The returned [Source] implements [Source64].
=======
// NewSourceは、指定された値でシードされた新しい擬似乱数Sourceを返します。
// トップレベルの関数で使用されるデフォルトのSourceとは異なり、このSourceは
// 複数のゴルーチンによる並行使用には安全ではありません。
// 返されるSourceはSource64を実装します。
>>>>>>> release-branch.go1.21
func NewSource(seed int64) Source

// Randは、乱数のソースです。
type Rand struct {
	src Source
	s64 Source64

	// readVal contains remainder of 63-bit integer used for bytes
	// generation during most recent Read call.
	// It is saved so next Read call can start where the previous
	// one finished.
	readVal int64
	// readPos indicates the number of low-order bytes of readVal
	// that are still valid.
	readPos int8
}

<<<<<<< HEAD
// New returns a new [Rand] that uses random values from src
// to generate other random values.
func New(src Source) *Rand

// Seed uses the provided seed value to initialize the generator to a deterministic state.
// Seed should not be called concurrently with any other [Rand] method.
=======
// Newは、他の乱数を生成するためにsrcから乱数を使用する新しいRandを返します。
func New(src Source) *Rand

// Seedは、提供されたシード値を使用してジェネレータを決定的な状態に初期化します。
// Seedは、他のRandメソッドと同時に呼び出すべきではありません。
>>>>>>> release-branch.go1.21
func (r *Rand) Seed(seed int64)

// Int63は、非負の擬似乱数63ビット整数をint64として返します。
func (r *Rand) Int63() int64

// Uint32は、擬似乱数32ビット値をuint32として返します。
func (r *Rand) Uint32() uint32

// Uint64は、擬似乱数64ビット値をuint64として返します。
func (r *Rand) Uint64() uint64

// Int31は、非負の擬似乱数31ビット整数をint32として返します。
func (r *Rand) Int31() int32

// Intは、非負の擬似乱数intを返します。
func (r *Rand) Int() int

// Int63nは、半開区間[0,n)内の非負の擬似乱数をint64として返します。
// nが0以下の場合、パニックを引き起こします。
func (r *Rand) Int63n(n int64) int64

// Int31nは、半開区間[0,n)内の非負の擬似乱数をint32として返します。
// nが0以下の場合、パニックを引き起こします。
func (r *Rand) Int31n(n int32) int32

// Intnは、半開区間[0,n)内の非負の擬似乱数をintとして返します。
// nが0以下の場合、パニックを引き起こします。
func (r *Rand) Intn(n int) int

// Float64は、半開区間[0.0,1.0)内の擬似乱数をfloat64として返します。
func (r *Rand) Float64() float64

// Float32は、半開区間[0.0,1.0)内の擬似乱数をfloat32として返します。
func (r *Rand) Float32() float32

// Permは、半開区間[0,n)内の整数の擬似乱数順列を、n個のintのスライスとして返します。
func (r *Rand) Perm(n int) []int

// Shuffleは要素の順序を擬似ランダムにします。
// nは要素の数です。n < 0の場合、Shuffleはパニックを引き起こします。
// swapは、インデックスiとjの要素を交換します。
func (r *Rand) Shuffle(n int, swap func(i, j int))

// Readは、len(p)個のランダムなバイトを生成し、それらをpに書き込みます。
// 常にlen(p)とnilエラーを返します。
// Readは、他のRandメソッドと同時に呼び出すべきではありません。
func (r *Rand) Read(p []byte) (n int, err error)

<<<<<<< HEAD
// Seed uses the provided seed value to initialize the default Source to a
// deterministic state. Seed values that have the same remainder when
// divided by 2³¹-1 generate the same pseudo-random sequence.
// Seed, unlike the [Rand.Seed] method, is safe for concurrent use.
=======
// Seedは、提供されたシード値を使用してデフォルトのSourceを
// 決定的な状態に初期化します。2³¹-1で割った余りが同じであるシード値は、
// 同じ擬似乱数系列を生成します。
// Seedは、Rand.Seedメソッドとは異なり、並行使用に安全です。
>>>>>>> release-branch.go1.21
//
// Seedが呼び出されない場合、ジェネレータはプログラムの起動時にランダムにシードされます。
//
// Go 1.20より前では、ジェネレータはプログラムの起動時にSeed(1)のようにシードされました。
// 古い振る舞いを強制するには、プログラムの起動時にSeed(1)を呼び出します。
// あるいは、このパッケージの関数を呼び出す前に環境変数でGODEBUG=randautoseed=0を設定します。
//
// Deprecated: Go 1.20以降、ランダムな値でSeedを呼び出す理由はありません。
// 特定の結果のシーケンスを得るために既知の値でSeedを呼び出すプログラムは、
// New(NewSource(seed))を使用してローカルのランダムジェネレータを取得するべきです。
func Seed(seed int64)

<<<<<<< HEAD
// Int63 returns a non-negative pseudo-random 63-bit integer as an int64
// from the default [Source].
func Int63() int64

// Uint32 returns a pseudo-random 32-bit value as a uint32
// from the default [Source].
func Uint32() uint32

// Uint64 returns a pseudo-random 64-bit value as a uint64
// from the default [Source].
func Uint64() uint64

// Int31 returns a non-negative pseudo-random 31-bit integer as an int32
// from the default [Source].
func Int31() int32

// Int returns a non-negative pseudo-random int from the default [Source].
func Int() int

// Int63n returns, as an int64, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default [Source].
// It panics if n <= 0.
func Int63n(n int64) int64

// Int31n returns, as an int32, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default [Source].
// It panics if n <= 0.
func Int31n(n int32) int32

// Intn returns, as an int, a non-negative pseudo-random number in the half-open interval [0,n)
// from the default [Source].
// It panics if n <= 0.
func Intn(n int) int

// Float64 returns, as a float64, a pseudo-random number in the half-open interval [0.0,1.0)
// from the default [Source].
func Float64() float64

// Float32 returns, as a float32, a pseudo-random number in the half-open interval [0.0,1.0)
// from the default [Source].
func Float32() float32

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers
// in the half-open interval [0,n) from the default [Source].
func Perm(n int) []int

// Shuffle pseudo-randomizes the order of elements using the default [Source].
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func Shuffle(n int, swap func(i, j int))

// Read generates len(p) random bytes from the default [Source] and
// writes them into p. It always returns len(p) and a nil error.
// Read, unlike the [Rand.Read] method, is safe for concurrent use.
//
// Deprecated: For almost all use cases, [crypto/rand.Read] is more appropriate.
func Read(p []byte) (n int, err error)

// NormFloat64 returns a normally distributed float64 in the range
// [-[math.MaxFloat64], +[math.MaxFloat64]] with
// standard normal distribution (mean = 0, stddev = 1)
// from the default [Source].
// To produce a different normal distribution, callers can
// adjust the output using:
=======
// Int63は、デフォルトのSourceから非負の擬似乱数63ビット整数をint64として返します。
func Int63() int64

// Uint32は、デフォルトのSourceから擬似乱数32ビット値をuint32として返します。
func Uint32() uint32

// Uint64は、デフォルトのSourceから擬似乱数64ビット値をuint64として返します。
func Uint64() uint64

// Int31は、デフォルトのSourceから非負の擬似乱数31ビット整数をint32として返します。
func Int31() int32

// Intは、デフォルトのSourceから非負の擬似乱数intを返します。
func Int() int

// Int63nは、デフォルトのSourceから半開区間[0,n)内の非負の擬似乱数をint64として返します。
// nが0以下の場合、パニックを引き起こします。
func Int63n(n int64) int64

// Int31nは、デフォルトのSourceから半開区間[0,n)内の非負の擬似乱数をint32として返します。
// nが0以下の場合、パニックを引き起こします。
func Int31n(n int32) int32

// Intnは、デフォルトのSourceから半開区間[0,n)内の非負の擬似乱数をintとして返します。
// nが0以下の場合、パニックを引き起こします。
func Intn(n int) int

// Float64は、デフォルトのSourceから半開区間[0.0,1.0)内の擬似乱数をfloat64として返します。
func Float64() float64

// Float32は、デフォルトのSourceから半開区間[0.0,1.0)内の擬似乱数をfloat32として返します。
func Float32() float32

// Permは、デフォルトのSourceから半開区間[0,n)内の整数の擬似乱数順列を、n個のintのスライスとして返します。
func Perm(n int) []int

// ShuffleはデフォルトのSourceを使用して要素の順序を擬似ランダムにします。
// nは要素の数です。n < 0の場合、Shuffleはパニックを引き起こします。
// swapは、インデックスiとjの要素を交換します。
func Shuffle(n int, swap func(i, j int))

// Readは、デフォルトのSourceからlen(p)個のランダムなバイトを生成し、それらをpに書き込みます。
// 常にlen(p)とnilエラーを返します。
// Readは、Rand.Readメソッドとは異なり、並行使用に安全です。
//
// Deprecated: ほとんどの使用ケースでは、crypto/rand.Readの方が適切です。
func Read(p []byte) (n int, err error)

// NormFloat64は、デフォルトのSourceから、範囲
// [-math.MaxFloat64, +math.MaxFloat64]内の正規分布に従うfloat64を返します。
// 標準正規分布（平均 = 0、標準偏差 = 1）です。
// 異なる正規分布を生成するために、呼び出し元は
// 出力を調整することができます：
>>>>>>> release-branch.go1.21
//
//	sample = NormFloat64() * desiredStdDev + desiredMean
func NormFloat64() float64

<<<<<<< HEAD
// ExpFloat64 returns an exponentially distributed float64 in the range
// (0, +[math.MaxFloat64]] with an exponential distribution whose rate parameter
// (lambda) is 1 and whose mean is 1/lambda (1) from the default [Source].
// To produce a distribution with a different rate parameter,
// callers can adjust the output using:
=======
// ExpFloat64は、デフォルトのSourceから、範囲
// (0, +math.MaxFloat64]内の指数分布に従うfloat64を返します。
// レートパラメータ（ラムダ）が1で、平均が1/ラムダ（1）の指数分布です。
// 異なるレートパラメータの分布を生成するために、
// 呼び出し元は出力を調整することができます：
>>>>>>> release-branch.go1.21
//
//	sample = ExpFloat64() / desiredRateParameter
func ExpFloat64() float64
