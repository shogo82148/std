// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strconv

import "github.com/shogo82148/std/errors"

// ErrRangeは、値が対象の型の範囲外であることを示します。
var ErrRange = errors.New("value out of range")

// ErrSyntaxは、値がターゲットの型の正しい構文ではないことを示します。
var ErrSyntax = errors.New("invalid syntax")

// NumErrorは変換に失敗したことを記録します。
type NumError struct {
	Func string
	Num  string
	Err  error
}

func (e *NumError) Error() string

func (e *NumError) Unwrap() error

// IntSizeはintまたはuint値のビットサイズです。
const IntSize = intSize

// ParseUintはParseIntと同じですが、符号の接頭辞は許可されていません。
func ParseUint(s string, base int, bitSize int) (uint64, error)

// ParseIntは与えられた基数（0、2から36）とビットサイズ（0から64）で文字列sを解釈し、対応する値iを返します。
//
// 文字列は先頭に符号 "+" または "-" を持つことができます。
//
// 基数引数が0の場合、真の基数は符号の後に続く文字列の接頭辞で推測されます（存在する場合）："0b"の場合は2、"0"または"0o"の場合は8、"0x"の場合は16、それ以外の場合は10です。また、基数0の場合だけアンダースコア文字が許可されます。これはGoの構文で定義されている[整数リテラル]です。
//
// bitSize引数は結果が適合する必要のある整数型を指定します。ビットサイズ0、8、16、32、64はint、int8、int16、int32、int64に対応します。bitSizeが0未満または64を超える場合、エラーが返されます。
//
// ParseIntが返すエラーは具体的な型*NumErrorを持ち、err.Num = sとなります。sが空であるか無効な数字を含んでいる場合、err.Err = ErrSyntaxとなり返される値は0です。sに対応する値を指定のサイズの符号付き整数で表現することができない場合、err.Err = ErrRangeとなり、返される値は適切なbitSizeと符号の最大の大きさの整数です。
//
// [整数リテラル]: https://go.dev/ref/spec#Integer_literals
func ParseInt(s string, base int, bitSize int) (i int64, err error)

// AtoiはParseInt(s, 10, 0)と同じであり、int型に変換されます。
func Atoi(s string) (int, error)
