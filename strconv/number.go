// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/internal/strconv"
)

// IntSizeは、intまたはuint値のビット単位のサイズです。
const IntSize = strconv.IntSize

// ParseBoolは、文字列が表す真偽値を返します。
// 1、t、T、TRUE、true、True、0、f、F、FALSE、false、Falseを受け付けます。
// その他の値はエラーを返します。
func ParseBool(str string) (bool, error)

// FormatBoolは、bの値に応じて"true"または"false"を返します。
func FormatBool(b bool) string

// AppendBoolは、bの値に応じて"true"または"false"をdstに追加し、
// 拡張されたバッファーを返します。
func AppendBool(dst []byte, b bool) []byte

// ParseComplexは、文字列sをbitSizeで指定された精度（complex64の場合は64、
// complex128の場合は128）の複素数に変換します。
// bitSize=64の場合、結果の型はcomplex128のままですが、値を変更せずに
// complex64に変換できます。
//
// sが表す数値は、N、Ni、またはN±Niの形式でなければなりません。ここでNは
// [ParseFloat] が認識する浮動小数点数を表し、iは虚数成分です。
// 2番目のNが符号なしの場合、±で示されるように2つの成分の間に+記号が必要です。
// 2番目のNがNaNの場合、+記号のみが受け付けられます。
// 形式は括弧で囲むことができ、スペースを含めることはできません。
// 結果の複素数は、ParseFloatによって変換された2つの成分で構成されます。
//
// ParseComplexが返すエラーの具体的な型は [*NumError] であり、
// err.Num = sを含みます。
//
// sが構文的に正しくない場合、ParseComplexはerr.Err = ErrSyntaxを返します。
//
// sが構文的に正しいが、いずれかの成分が指定された成分のサイズの最大の浮動小数点数から
// 1/2 ULP以上離れている場合、ParseComplexはerr.Err = ErrRangeと、
// 該当する成分についてc = ±Infを返します。
func ParseComplex(s string, bitSize int) (complex128, error)

// ParseFloatは、文字列sをbitSizeで指定された精度（float32の場合は32、
// float64の場合は64）の浮動小数点数に変換します。
// bitSize=32の場合、結果の型はfloat64のままですが、値を変更せずに
// float32に変換できます。
//
// ParseFloatは、Goの[floating-point literals]の構文で定義された
// 10進数および16進数の浮動小数点数を受け付けます。
// sが正しい形式で有効な浮動小数点数に近い場合、ParseFloatは
// IEEE754の偏りのない丸めを使用して丸められた最も近い浮動小数点数を返します。
// （16進数の浮動小数点値のパースでは、16進数表現に仮数部に収まるより多くの
// ビットがある場合にのみ丸めが行われます。）
//
// ParseFloatが返すエラーの具体的な型は*NumErrorであり、
// err.Num = sを含みます。
//
// sが構文的に正しくない場合、ParseFloatはerr.Err = ErrSyntaxを返します。
//
// sが構文的に正しいが、指定されたサイズの最大の浮動小数点数から
// 1/2 ULP以上離れている場合、ParseFloatはf = ±Inf、err.Err = ErrRangeを返します。
//
// ParseFloatは、文字列"NaN"、および（符号付きの可能性がある）文字列"Inf"と"Infinity"を
// それぞれの特別な浮動小数点値として認識します。マッチングの際は大文字と小文字を区別しません。
//
// [floating-point literals]: https://go.dev/ref/spec#Floating-point_literals
func ParseFloat(s string, bitSize int) (float64, error)

// ParseUintは [ParseInt] と同様ですが、符号なしの数値用です。
//
// 符号の接頭辞は許可されません。
func ParseUint(s string, base int, bitSize int) (uint64, error)

// ParseIntは、指定された基数（0、2から36）とビットサイズ（0から64）で
// 文字列sを解釈し、対応する値iを返します。
//
// 文字列は先頭に符号"+"または"-"を付けることができます。
//
// base引数が0の場合、実際の基数は符号（存在する場合）に続く文字列の接頭辞によって
// 暗黙的に決まります。"0b"の場合は2、"0"または"0o"の場合は8、"0x"の場合は16、
// それ以外の場合は10です。また、base引数が0の場合に限り、Goの
// [integer literals]の構文で定義されているアンダースコア文字が許可されます。
//
// bitSize引数は、結果が収まらなければならない整数型を指定します。
// ビットサイズ0、8、16、32、64は、それぞれint、int8、int16、int32、int64に対応します。
// bitSizeが0未満または64を超える場合、エラーが返されます。
//
// ParseIntが返すエラーの具体的な型は [*NumError] であり、
// err.Num = sを含みます。sが空であるか無効な数字を含む場合、
// err.Err = [ErrSyntax] となり、返される値は0です。
// sに対応する値が指定されたサイズの符号付き整数で表現できない場合、
// err.Err = [ErrRange] となり、返される値は適切なbitSizeと符号の
// 絶対値が最大の整数です。
//
// [integer literals]: https://go.dev/ref/spec#Integer_literals
func ParseInt(s string, base int, bitSize int) (i int64, err error)

// Atoiは、ParseInt(s, 10, 0)をint型に変換したものと同等です。
func Atoi(s string) (int, error)

// FormatComplexは、複素数cを(a+bi)の形式の文字列に変換します。
// ここでaとbは実部と虚部であり、書式fmtと精度precに従ってフォーマットされます。
//
// 書式fmtと精度precは [FormatFloat] と同じ意味を持ちます。
// 元の値がbitSizeビットの複素数値から得られたものと仮定して結果を丸めます。
// bitSizeはcomplex64の場合は64、complex128の場合は128でなければなりません。
func FormatComplex(c complex128, fmt byte, prec, bitSize int) string

// FormatFloatは、浮動小数点数fを書式fmtと精度precに従って文字列に変換します。
// 元の値がbitSizeビット（float32の場合は32、float64の場合は64）の
// 浮動小数点値から得られたものと仮定して結果を丸めます。
//
// 書式fmtは次のいずれかです
//   - 'b' (-ddddp±ddd、2進指数),
//   - 'e' (-d.dddde±dd、10進指数),
//   - 'E' (-d.ddddE±dd、10進指数),
//   - 'f' (-ddd.dddd、指数なし),
//   - 'g' (大きな指数の場合は'e'、それ以外は'f'),
//   - 'G' (大きな指数の場合は'E'、それ以外は'f'),
//   - 'x' (-0xd.ddddp±ddd、16進数の仮数と2進指数), または
//   - 'X' (-0Xd.ddddP±ddd、16進数の仮数と2進指数).
//
// 精度precは、'e'、'E'、'f'、'g'、'G'、'x'、'X'形式で出力される
// 桁数（指数部は除く）を制御します。
// 'e'、'E'、'f'、'x'、'X'では、小数点以下の桁数を表します。
// 'g'と'G'では、有効数字の最大桁数を表します（末尾のゼロは削除されます）。
// 特別な精度-1は、ParseFloatがfを正確に返すために必要な最小の桁数を使用します。
// 指数は10進整数として書かれます。
// 'b'以外のすべての形式では、少なくとも2桁になります。
func FormatFloat(f float64, fmt byte, prec, bitSize int) string

// AppendFloatは、[FormatFloat] によって生成された浮動小数点数fの文字列形式を
// dstに追加し、拡張されたバッファーを返します。
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte

// FormatUintは、2 <= base <= 36の指定された基数でのiの文字列表現を返します。
// 結果は、10以上の数字の値に小文字の'a'から'z'を使用します。
func FormatUint(i uint64, base int) string

// FormatIntは、2 <= base <= 36の指定された基数でのiの文字列表現を返します。
// 結果は、10以上の数字の値に小文字の'a'から'z'を使用します。
func FormatInt(i int64, base int) string

// Itoaは [FormatInt](int64(i), 10)と同等です。
func Itoa(i int) string

// AppendIntは、[FormatInt] によって生成された整数iの文字列形式を
// dstに追加し、拡張されたバッファーを返します。
func AppendInt(dst []byte, i int64, base int) []byte

// AppendUintは、[FormatUint] によって生成された符号なし整数iの文字列形式を
// dstに追加し、拡張されたバッファーを返します。
func AppendUint(dst []byte, i uint64, base int) []byte

// ErrRangeは、値が対象の型の範囲外であることを示します。
var ErrRange = errors.New("value out of range")

// ErrSyntaxは、値が対象の型に対して正しい構文を持っていないことを示します。
var ErrSyntax = errors.New("invalid syntax")

// NumErrorは失敗した変換を記録します。
type NumError struct {
	Func string
	Num  string
	Err  error
}

func (e *NumError) Error() string

func (e *NumError) Unwrap() error
