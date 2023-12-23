// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements signed multi-precision integers.

package big

import (
	"github.com/shogo82148/std/math/rand"
)

// Intは、符号付きの多倍長整数を表します。
// Intのゼロ値は値0を表します。
//
// 操作は常にポインタ引数（*Int）を取り、
// 各ユニークなInt値は自身のユニークな*Intポインタを必要とします。
// Int値を「コピー」するには、既存の（または新しく割り当てられた）Intを
// Int.Setメソッドを使用して新しい値に設定する必要があります。
// Intの浅いコピーはサポートされておらず、エラーを引き起こす可能性があります。
//
// メソッドは、タイミングのサイドチャネルを通じてIntの値を漏らす可能性があることに注意してください。
// このため、そして実装の範囲と複雑さのため、Intは暗号化操作を実装するのに適していません。
// 標準ライブラリは、攻撃者が制御する入力に対して非自明なIntメソッドを公開することを避け、
// math/bigのバグがセキュリティ脆弱性と見なされるかどうかは、標準ライブラリへの影響によって決まる可能性があります。
type Int struct {
	neg bool
	abs nat
}

// Signは次の値を返します:
//
//	-1 は x <  0 の場合
//	 0 は x == 0 の場合
//	+1 は x >  0 の場合
func (x *Int) Sign() int

// SetInt64はzをxに設定し、zを返します。
func (z *Int) SetInt64(x int64) *Int

// SetUint64はzをxに設定し、zを返します。
func (z *Int) SetUint64(x uint64) *Int

// NewIntは新しいIntを割り当て、xに設定して返します。
func NewInt(x int64) *Int

// Setはzをxに設定し、zを返します。
func (z *Int) Set(x *Int) *Int

// Bitsは、xの絶対値をリトルエンディアンのWordスライスとして返すことで、
// xへの生の（チェックされていないが高速な）アクセスを提供します。結果とxは
// 同じ基本配列を共有します。
// Bitsは、このパッケージ外部で欠けている低レベルのInt機能の実装をサポートすることを目的としています。
// それ以外の場合は避けるべきです。
func (x *Int) Bits() []Word

// SetBitsは、zの値をリトルエンディアンのWordスライスとして解釈されるabsに設定し、
// zを返すことで、zへの生の（チェックされていないが高速な）アクセスを提供します。
// 結果とabsは同じ基本配列を共有します。
// SetBitsは、このパッケージ外部で欠けている低レベルのInt機能の実装をサポートすることを目的としています。
// それ以外の場合は避けるべきです。
func (z *Int) SetBits(abs []Word) *Int

// Absはzを|x|（xの絶対値）に設定し、zを返します。
func (z *Int) Abs(x *Int) *Int

// Negはzを-xに設定し、zを返します。
func (z *Int) Neg(x *Int) *Int

// Addはzをx+yの和に設定し、zを返します。
func (z *Int) Add(x, y *Int) *Int

// Subはzをx-yの差に設定し、zを返します。
func (z *Int) Sub(x, y *Int) *Int

// Mulはzをx*yの積に設定し、zを返します。
func (z *Int) Mul(x, y *Int) *Int

// MulRangeは、zを範囲[a, b]（両端を含む）内のすべての整数の積に設定し、zを返します。
// a > b（範囲が空）の場合、結果は1です。
func (z *Int) MulRange(a, b int64) *Int

// Binomialは、zを二項係数C(n, k)に設定し、zを返します。
func (z *Int) Binomial(n, k int64) *Int

// Quoは、y != 0の場合、zを商x/yに設定し、zを返します。
// y == 0の場合、ゼロ除算のランタイムパニックが発生します。
// Quoは切り捨て除算（Goと同様）を実装します。詳細はQuoRemを参照してください。
func (z *Int) Quo(x, y *Int) *Int

// Remは、y != 0の場合、zを余りx%yに設定し、zを返します。
// y == 0の場合、ゼロ除算のランタイムパニックが発生します。
// Remは切り捨てモジュラス（Goと同様）を実装します。詳細はQuoRemを参照してください。
func (z *Int) Rem(x, y *Int) *Int

// QuoRemは、y != 0の場合、zを商x/yに、rを余りx%yに設定し、
// ペア(z, r)を返します。
// y == 0の場合、ゼロ除算のランタイムパニックが発生します。
//
// QuoRemはT-除算とモジュラス（Goと同様）を実装します：
//
//	q = x/y      結果はゼロに切り捨てられます
//	r = x - y*q
//
// （Daan Leijenの「コンピュータサイエンティストのための除算とモジュラス」を参照）
// ユークリッド除算とモジュラス（Goとは異なる）についてはDivModを参照してください。
func (z *Int) QuoRem(x, y, r *Int) (*Int, *Int)

// Divは、y != 0の場合、zを商x/yに設定し、zを返します。
// y == 0の場合、ゼロ除算のランタイムパニックが発生します。
// Divはユークリッド除算を実装します（Goとは異なります）；詳細はDivModを参照してください。
func (z *Int) Div(x, y *Int) *Int

// Modは、y != 0の場合、zを余りx%yに設定し、zを返します。
// y == 0の場合、ゼロ除算のランタイムパニックが発生します。
// Modはユークリッドのモジュラスを実装します（Goとは異なります）；詳細はDivModを参照してください。
func (z *Int) Mod(x, y *Int) *Int

// DivModは、y != 0の場合、zを商x div yに、mを余りx mod yに設定し、
// ペア(z, m)を返します。
// y == 0の場合、ゼロ除算のランタイムパニックが発生します。
//
// DivModはユークリッドの除算とモジュラスを実装します（Goとは異なります）：
//
//	q = x div y  となるような
//	m = x - y*q  で 0 <= m < |y|
//
// （Raymond T. Boute, "The Euclidean definition of the functions
// div and mod". ACM Transactions on Programming Languages and
// Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992.
// ACM press.を参照）
// T-除算とモジュラス（Goと同様）についてはQuoRemを参照してください。
func (z *Int) DivMod(x, y, m *Int) (*Int, *Int)

// Cmpはxとyを比較し、次の値を返します:
//
//	-1 は x <  y の場合
//	 0 は x == y の場合
//	+1 は x >  y の場合
func (x *Int) Cmp(y *Int) (r int)

// CmpAbsはxとyの絶対値を比較し、次の値を返します:
//
//	-1 は |x| <  |y| の場合
//	 0 は |x| == |y| の場合
//	+1 は |x| >  |y| の場合
func (x *Int) CmpAbs(y *Int) int

// Int64はxのint64表現を返します。
// もしxがint64で表現できない場合、結果は未定義です。
func (x *Int) Int64() int64

// Uint64はxのuint64表現を返します。
// もしxがuint64で表現できない場合、結果は未定義です。
func (x *Int) Uint64() uint64

// IsInt64は、xがint64として表現できるかどうかを報告します。
func (x *Int) IsInt64() bool

// IsUint64は、xがuint64として表現できるかどうかを報告します。
func (x *Int) IsUint64() bool

// Float64は、xに最も近いfloat64の値と、
// 発生した丸め処理の有無を示す指標を返します。
func (x *Int) Float64() (float64, Accuracy)

// SetStringは、zを指定された基数で解釈されたsの値に設定し、
// zと成功を示すブール値を返します。成功するためには、文字列全体（プレフィックスだけでなく）
// が有効である必要があります。SetStringが失敗した場合、zの値は未定義ですが、
// 返される値はnilです。
//
// 基数引数は0または2からMaxBaseの間の値でなければなりません。
// 基数が0の場合、数値のプレフィックスが実際の基数を決定します：プレフィックスが
// "0b"または"0B"は基数2を選択し、"0"、"0o"または"0O"は基数8を選択し、
// "0x"または"0X"は基数16を選択します。それ以外の場合、選択された基数は10であり、
// プレフィックスは受け付けられません。
//
// 基数が36以下の場合、小文字と大文字は同じとみなされます：
// 文字 'a' から 'z' と 'A' から 'Z' は、数字の値 10 から 35 を表します。
// 基数が36より大きい場合、大文字の 'A' から 'Z' は、数字の値 36 から 61 を表します。
//
// 基数が0の場合、アンダースコア文字 "_" は基数のプレフィックスと隣接する数字の間、
// または連続する数字の間に現れることがあります。このようなアンダースコアは数値の値に影響しません。
// アンダースコアの配置が不適切な場合、他にエラーがない場合にエラーとして報告されます。
// 基数が0でない場合、アンダースコアは認識されず、有効な数字でない他の任意の文字と同様に動作します。
func (z *Int) SetString(s string, base int) (*Int, bool)

// SetBytesは、bufをビッグエンディアンの符号なし整数のバイトとして解釈し、
// zをその値に設定し、zを返します。
func (z *Int) SetBytes(buf []byte) *Int

// Bytesは、xの絶対値をビッグエンディアンのバイトスライスとして返します。
//
// 固定長のスライスや、事前に割り当てられたものを使用するには、FillBytesを使用します。
func (x *Int) Bytes() []byte

// FillBytesは、bufをxの絶対値に設定し、それをゼロ拡張のビッグエンディアンのバイトスライスとして格納し、
// bufを返します。
//
// もしxの絶対値がbufに収まらない場合、FillBytesはパニックを起こします。
func (x *Int) FillBytes(buf []byte) []byte

// BitLenは、xの絶対値の長さをビット単位で返します。
// 0のビット長は0です。
func (x *Int) BitLen() int

// TrailingZeroBitsは、|x|の連続する最下位ゼロビットの数を返します。
func (x *Int) TrailingZeroBits() uint

// Expは、z = x**y mod |m|（つまり、mの符号は無視されます）を設定し、zを返します。
// もしm == nilまたはm == 0なら、y <= 0ならz = 1、それ以外の場合はz = x**yです。
// もしm != 0、y < 0、そしてxとmが相互に素ではない場合、zは変更されず、nilが返されます。
//
// 特定のサイズの入力のモジュラ指数は、暗号学的に一定時間の操作ではありません。
func (z *Int) Exp(x, y, m *Int) *Int

// GCDは、zをaとbの最大公約数に設定し、zを返します。
// もしxまたはyがnilでなければ、GCDはz = a*x + b*yとなるようにそれらの値を設定します。
//
// aとbは正、ゼロ、または負のいずれかである可能性があります。（Go 1.14以前は両方とも
// > 0である必要がありました。）aとbの符号に関係なく、zは常に>= 0です。
//
// もしa == b == 0なら、GCDはz = x = y = 0に設定します。
//
// もしa == 0でb != 0なら、GCDはz = |b|、x = 0、y = sign(b) * 1に設定します。
//
// もしa != 0でb == 0なら、GCDはz = |a|、x = sign(a) * 1、y = 0に設定します。
func (z *Int) GCD(x, y, a, b *Int) *Int

// Randは、zを[0, n)の範囲の擬似乱数に設定し、zを返します。
//
// これはmath/randパッケージを使用しているため、
// セキュリティに敏感な作業には使用してはなりません。代わりにcrypto/rand.Intを使用してください。
func (z *Int) Rand(rnd *rand.Rand, n *Int) *Int

// ModInverseは、zを環ℤ/nℤにおけるgの乗法的逆数に設定し、zを返します。
// もしgとnが互いに素でない場合、gは環ℤ/nℤに乗法的逆数を持ちません。
// この場合、zは変更されず、戻り値はnilです。もしn == 0なら、ゼロ除算のランタイムパニックが発生します。
func (z *Int) ModInverse(g, n *Int) *Int

// Jacobiは、ヤコビ記号 (x/y) を返します。これは+1、-1、または0のいずれかです。
// y引数は奇数でなければなりません。
func Jacobi(x, y *Int) int

// ModSqrtは、存在する場合、zをx mod pの平方根に設定し、zを返します。
// 剰余pは奇数の素数でなければなりません。もしxがp modの平方でない場合、
// ModSqrtはzを変更せず、nilを返します。この関数は、pが奇数でない場合にパニックを起こします。
// pが奇数だが素数でない場合の動作は未定義です。
func (z *Int) ModSqrt(x, p *Int) *Int

// Lshは、z = x << nを設定し、zを返します。
func (z *Int) Lsh(x *Int, n uint) *Int

// Rshは、z = x >> nを設定し、zを返します。
func (z *Int) Rsh(x *Int, n uint) *Int

// Bitは、xのi番目のビットの値を返します。つまり、
// (x>>i)&1を返します。ビットインデックスiは0以上でなければなりません。
func (x *Int) Bit(i int) uint

// SetBitは、xのi番目のビットをb（0または1）に設定したxをzに設定します。
// つまり、もしbが1なら、SetBitはz = x | (1 << i)を設定します。
// もしbが0なら、SetBitはz = x &^ (1 << i)を設定します。もしbが0または1でない場合、
// SetBitはパニックを起こします。
func (z *Int) SetBit(x *Int, i int, b uint) *Int

// Andは、z = x & yを設定し、zを返します。
func (z *Int) And(x, y *Int) *Int

// AndNotは、z = x &^ yを設定し、zを返します。
func (z *Int) AndNot(x, y *Int) *Int

// Orは、z = x | yを設定し、zを返します。
func (z *Int) Or(x, y *Int) *Int

// Xorは、z = x ^ yを設定し、zを返します。
func (z *Int) Xor(x, y *Int) *Int

// Notは、z = ^xを設定し、zを返します。
func (z *Int) Not(x *Int) *Int

// Sqrtは、zを⌊√x⌋（つまり、z² ≤ xとなる最大の整数）に設定し、zを返します。
// xが負の場合、パニックを起こします。
func (z *Int) Sqrt(x *Int) *Int
