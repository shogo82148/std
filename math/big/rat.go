// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements multi-precision rational numbers.

package big

// Ratは、任意の精度の商a/bを表します。
// Ratのゼロ値は値0を表します。
//
// 操作は常にポインタ引数（*Rat）を取る
// 代わりにRat値、そして各ユニークなRat値は
// 自身のユニークな*Ratポインタが必要です。Rat値を「コピー」するには、
// 既存の（または新しく割り当てられた）Ratを
// [Rat.Set] メソッドを使用して新しい値に設定する必要があります。Ratsの浅いコピーは
// サポートされておらず、エラーを引き起こす可能性があります。
type Rat struct {
	// To make zero values for Rat work w/o initialization,
	// a zero value of b (len(b) == 0) acts like b == 1. At
	// the earliest opportunity (when an assignment to the Rat
	// is made), such uninitialized denominators are set to 1.
	// a.neg determines the sign of the Rat, b.neg is ignored.
	a, b Int
}

// NewRatは、分子aと分母bを持つ新しい [Rat] を作成します。
func NewRat(a, b int64) *Rat

// SetFloat64は、zを正確にfに設定し、zを返します。
// もしfが有限でない場合、SetFloatはnilを返します。
func (z *Rat) SetFloat64(f float64) *Rat

// Float32は、xに最も近いfloat32値と、
// fがxを正確に表現しているかどうかを示すbool値を返します。
// もしxの絶対値がfloat32で表現できる範囲を超えている場合、
// fは無限大となり、exactはfalseとなります。
// fの符号は、fが0であっても、常にxの符号と一致します。
func (x *Rat) Float32() (f float32, exact bool)

// Float64は、xに最も近いfloat64値と、
// fがxを正確に表現しているかどうかを示すbool値を返します。
// もしxの絶対値がfloat64で表現できる範囲を超えている場合、
// fは無限大となり、exactはfalseとなります。
// fの符号は、fが0であっても、常にxの符号と一致します。
func (x *Rat) Float64() (f float64, exact bool)

// SetFracは、zをa/bに設定し、zを返します。
// もしb == 0の場合、SetFracはパニックを引き起こします。
func (z *Rat) SetFrac(a, b *Int) *Rat

// SetFrac64は、zをa/bに設定し、zを返します。
// もしb == 0の場合、SetFrac64はパニックを引き起こします。
func (z *Rat) SetFrac64(a, b int64) *Rat

// SetIntは、zをxに設定します（xのコピーを作成します）そしてzを返します。
func (z *Rat) SetInt(x *Int) *Rat

// SetInt64は、zをxに設定し、zを返します。
func (z *Rat) SetInt64(x int64) *Rat

// SetUint64は、zをxに設定し、zを返します。
func (z *Rat) SetUint64(x uint64) *Rat

// Setは、zをxに設定します（xのコピーを作成します）そしてzを返します。
func (z *Rat) Set(x *Rat) *Rat

// Absは、zを|x|に設定します（xの絶対値）そしてzを返します。
func (z *Rat) Abs(x *Rat) *Rat

// Negは、zを-xに設定し、zを返します。
func (z *Rat) Neg(x *Rat) *Rat

// Invは、zを1/xに設定し、zを返します。
// もしx == 0の場合、Invはパニックを引き起こします。
func (z *Rat) Inv(x *Rat) *Rat

// Signは以下を返します:
//
//   - -1 x <  0 の場合
//   - 0 x == 0 の場合
//   - +1 x >  0 の場合
func (x *Rat) Sign() int

// IsIntは、xの分母が1であるかどうかを報告します。
func (x *Rat) IsInt() bool

// Numはxの分子を返します。これは0以下になる可能性があります。
// 結果はxの分子への参照であり、xに新しい値が割り当てられると変更される可能性があります。逆も同様です。
// 分子の符号はxの符号に対応します。
func (x *Rat) Num() *Int

// Denomはxの分母を返します。これは常に> 0です。
// 結果はxの分母への参照であり、
// xが初期化されていない（ゼロ値の）[Rat] の場合、
// 結果は値1の新しい [Int] になります。（xを初期化するには、
// xを設定する任意の操作が適用できます、x.Set(x)を含む。）
// 結果がxの分母への参照である場合、
// 新しい値がxに割り当てられると変更される可能性があります。逆も同様です。
func (x *Rat) Denom() *Int

// Cmpはxとyを比較し、以下を返します:
//
//   - -1 x <  y の場合
//   - 0 x == y の場合
//   - +1 x >  y の場合
func (x *Rat) Cmp(y *Rat) int

// Addはzをx+yの和に設定し、zを返します。
func (z *Rat) Add(x, y *Rat) *Rat

// Subはzをx-yの差に設定し、zを返します。
func (z *Rat) Sub(x, y *Rat) *Rat

// Mulはzをx*yの積に設定し、zを返します。
func (z *Rat) Mul(x, y *Rat) *Rat

// Quoはzをx/yの商に設定し、zを返します。
// もしy == 0の場合、Quoはパニックを引き起こします。
func (z *Rat) Quo(x, y *Rat) *Rat
