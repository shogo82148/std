// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run make_tables.go

// bitsパッケージは、事前に宣言された符号なし整数型のためのビットカウントと操作
// 関数を実装します。
//
// このパッケージの関数は、パフォーマンス向上のためにコンパイラによって直接実装される可能性があります。
// そのような関数の場合、このパッケージのコードは使用されません。
// どの関数がコンパイラによって実装されるかは、アーキテクチャとGoのリリースによります。
package bits

// UintSizeは、uintのビット単位のサイズです。
const UintSize = uintSize

<<<<<<< HEAD
// LeadingZeros returns the number of leading zero bits in x; the result is [UintSize] for x == 0.
=======
// LeadingZerosは、xの先頭のゼロビットの数を返します。x == 0の場合、結果はUintSizeです。
>>>>>>> release-branch.go1.21
func LeadingZeros(x uint) int

// LeadingZeros8は、xの先頭のゼロビットの数を返します。x == 0の場合、結果は8です。
func LeadingZeros8(x uint8) int

// LeadingZeros16は、xの先頭のゼロビットの数を返します。x == 0の場合、結果は16です。
func LeadingZeros16(x uint16) int

// LeadingZeros32は、xの先頭のゼロビットの数を返します。x == 0の場合、結果は32です。
func LeadingZeros32(x uint32) int

// LeadingZeros64は、xの先頭のゼロビットの数を返します。x == 0の場合、結果は64です。
func LeadingZeros64(x uint64) int

<<<<<<< HEAD
// TrailingZeros returns the number of trailing zero bits in x; the result is [UintSize] for x == 0.
=======
// TrailingZerosは、xの末尾のゼロビットの数を返します。x == 0の場合、結果はUintSizeです。
>>>>>>> release-branch.go1.21
func TrailingZeros(x uint) int

// TrailingZeros8は、xの末尾のゼロビットの数を返します。x == 0の場合、結果は8です。
func TrailingZeros8(x uint8) int

// TrailingZeros16は、xの末尾のゼロビットの数を返します。x == 0の場合、結果は16です。
func TrailingZeros16(x uint16) int

// TrailingZeros32は、xの末尾のゼロビットの数を返します。x == 0の場合、結果は32です。
func TrailingZeros32(x uint32) int

// TrailingZeros64は、xの末尾のゼロビットの数を返します。x == 0の場合、結果は64です。
func TrailingZeros64(x uint64) int

// OnesCountは、xの1ビットの数（"ポピュレーションカウント"）を返します。
func OnesCount(x uint) int

// OnesCount8は、xの1ビットの数（"ポピュレーションカウント"）を返します。
func OnesCount8(x uint8) int

// OnesCount16は、xの1ビットの数（"ポピュレーションカウント"）を返します。
func OnesCount16(x uint16) int

// OnesCount32は、xの1ビットの数（"ポピュレーションカウント"）を返します。
func OnesCount32(x uint32) int

// OnesCount64は、xの1ビットの数（"ポピュレーションカウント"）を返します。
func OnesCount64(x uint64) int

<<<<<<< HEAD
// RotateLeft returns the value of x rotated left by (k mod [UintSize]) bits.
// To rotate x right by k bits, call RotateLeft(x, -k).
=======
// RotateLeftは、xを左に（k mod UintSize）ビット回転させた値を返します。
// xをkビット右に回転させるには、RotateLeft(x, -k)を呼び出します。
>>>>>>> release-branch.go1.21
//
// この関数の実行時間は入力に依存しません。
func RotateLeft(x uint, k int) uint

// RotateLeft8は、xを左に（k mod 8）ビット回転させた値を返します。
// xをkビット右に回転させるには、RotateLeft8(x, -k)を呼び出します。
//
// この関数の実行時間は入力に依存しません。
func RotateLeft8(x uint8, k int) uint8

// RotateLeft16は、xを左に（k mod 16）ビット回転させた値を返します。
// xをkビット右に回転させるには、RotateLeft16(x, -k)を呼び出します。
//
// この関数の実行時間は入力に依存しません。
func RotateLeft16(x uint16, k int) uint16

// RotateLeft32は、xを左に（k mod 32）ビット回転させた値を返します。
// xをkビット右に回転させるには、RotateLeft32(x, -k)を呼び出します。
//
// この関数の実行時間は入力に依存しません。
func RotateLeft32(x uint32, k int) uint32

// RotateLeft64は、xを左に（k mod 64）ビット回転させた値を返します。
// xをkビット右に回転させるには、RotateLeft64(x, -k)を呼び出します。
//
// この関数の実行時間は入力に依存しません。
func RotateLeft64(x uint64, k int) uint64

// Reverseは、ビットが逆順になったxの値を返します。
func Reverse(x uint) uint

// Reverse8は、ビットが逆順になったxの値を返します。
func Reverse8(x uint8) uint8

// Reverse16は、ビットが逆順になったxの値を返します。
func Reverse16(x uint16) uint16

// Reverse32は、ビットが逆順になったxの値を返します。
func Reverse32(x uint32) uint32

// Reverse64は、ビットが逆順になったxの値を返します。
func Reverse64(x uint64) uint64

// ReverseBytesは、バイトが逆順になったxの値を返します。
//
// この関数の実行時間は入力に依存しません。
func ReverseBytes(x uint) uint

// ReverseBytes16は、バイトが逆順になったxの値を返します。
//
// この関数の実行時間は入力に依存しません。
func ReverseBytes16(x uint16) uint16

// ReverseBytes32は、バイトが逆順になったxの値を返します。
//
// この関数の実行時間は入力に依存しません。
func ReverseBytes32(x uint32) uint32

// ReverseBytes64は、バイトが逆順になったxの値を返します。
//
// この関数の実行時間は入力に依存しません。
func ReverseBytes64(x uint64) uint64

// Lenは、xを表現するために必要なビットの最小数を返します。x == 0の場合、結果は0です。
func Len(x uint) int

// Len8は、xを表現するために必要なビットの最小数を返します。x == 0の場合、結果は0です。
func Len8(x uint8) int

// Len16は、xを表現するために必要なビットの最小数を返します。x == 0の場合、結果は0です。
func Len16(x uint16) (n int)

// Len32は、xを表現するために必要なビットの最小数を返します。x == 0の場合、結果は0です。
func Len32(x uint32) (n int)

// Len64は、xを表現するために必要なビットの最小数を返します。x == 0の場合、結果は0です。
func Len64(x uint64) (n int)

// Addは、x、y、およびcarryの和を返します：sum = x + y + carry。
// carry入力は0または1でなければなりません。それ以外の場合、動作は未定義です。
// carryOut出力は0または1であることが保証されています。
//
// この関数の実行時間は入力に依存しません。
func Add(x, y, carry uint) (sum, carryOut uint)

// Add32は、x、y、およびcarryの和を返します：sum = x + y + carry。
// carry入力は0または1でなければなりません。それ以外の場合、動作は未定義です。
// carryOut出力は0または1であることが保証されています。
//
// この関数の実行時間は入力に依存しません。
func Add32(x, y, carry uint32) (sum, carryOut uint32)

// Add64は、x、y、およびcarryの和を返します：sum = x + y + carry。
// carry入力は0または1でなければなりません。それ以外の場合、動作は未定義です。
// carryOut出力は0または1であることが保証されています。
//
// この関数の実行時間は入力に依存しません。
func Add64(x, y, carry uint64) (sum, carryOut uint64)

// Subは、x、y、およびborrowの差を返します：diff = x - y - borrow。
// borrow入力は0または1でなければなりません。それ以外の場合、動作は未定義です。
// borrowOut出力は0または1であることが保証されています。
//
// この関数の実行時間は入力に依存しません。
func Sub(x, y, borrow uint) (diff, borrowOut uint)

// Sub32は、x、y、およびborrowの差を返します：diff = x - y - borrow。
// borrow入力は0または1でなければなりません。それ以外の場合、動作は未定義です。
// borrowOut出力は0または1であることが保証されています。
//
// この関数の実行時間は入力に依存しません。
func Sub32(x, y, borrow uint32) (diff, borrowOut uint32)

// Sub64は、x、y、およびborrowの差を返します：diff = x - y - borrow。
// borrow入力は0または1でなければなりません。それ以外の場合、動作は未定義です。
// borrowOut出力は0または1であることが保証されています。
//
// この関数の実行時間は入力に依存しません。
func Sub64(x, y, borrow uint64) (diff, borrowOut uint64)

// Mulは、xとyの全幅の積を返します：(hi, lo) = x * y
// 積のビットの上半分はhiに、下半分はloに返されます。
//
// この関数の実行時間は入力に依存しません。
func Mul(x, y uint) (hi, lo uint)

// Mul32は、xとyの64ビットの積を返します：(hi, lo) = x * y
// 積のビットの上半分はhiに、下半分はloに返されます。
//
// この関数の実行時間は入力に依存しません。
func Mul32(x, y uint32) (hi, lo uint32)

// Mul64は、xとyの128ビットの積を返します：(hi, lo) = x * y
// 積のビットの上半分はhiに、下半分はloに返されます。
//
// この関数の実行時間は入力に依存しません。
func Mul64(x, y uint64) (hi, lo uint64)

// Divは、(hi, lo)をyで割った商と余りを返します：
// quo = (hi, lo)/y, rem = (hi, lo)%y
// 被除数のビットの上半分はパラメータhiに、下半分はパラメータloにあります。
// y == 0の場合（ゼロ除算）またはy <= hiの場合（商のオーバーフロー）、Divはパニックします。
func Div(hi, lo, y uint) (quo, rem uint)

// Div32は、(hi, lo)をyで割った商と余りを返します：
// quo = (hi, lo)/y, rem = (hi, lo)%y
// 被除数のビットの上半分はパラメータhiに、下半分はパラメータloにあります。
// y == 0の場合（ゼロ除算）またはy <= hiの場合（商のオーバーフロー）、Div32はパニックします。
func Div32(hi, lo, y uint32) (quo, rem uint32)

// Div64は、(hi, lo)をyで割った商と余りを返します：
// quo = (hi, lo)/y, rem = (hi, lo)%y
// 被除数のビットの上半分はパラメータhiに、下半分はパラメータloにあります。
// y == 0の場合（ゼロ除算）またはy <= hiの場合（商のオーバーフロー）、Div64はパニックします。
func Div64(hi, lo, y uint64) (quo, rem uint64)

// Remは、(hi, lo)をyで割った余りを返します。Remは
// y == 0の場合（ゼロ除算）にパニックしますが、Divとは異なり、
// 商のオーバーフローではパニックしません。
func Rem(hi, lo, y uint) uint

<<<<<<< HEAD
// Rem32 returns the remainder of (hi, lo) divided by y. Rem32 panics
// for y == 0 (division by zero) but, unlike [Div32], it doesn't panic
// on a quotient overflow.
func Rem32(hi, lo, y uint32) uint32

// Rem64 returns the remainder of (hi, lo) divided by y. Rem64 panics
// for y == 0 (division by zero) but, unlike [Div64], it doesn't panic
// on a quotient overflow.
=======
// Rem32は、(hi, lo)をyで割った余りを返します。Rem32は
// y == 0の場合（ゼロ除算）にパニックしますが、Div32とは異なり、
// 商のオーバーフローではパニックしません。
func Rem32(hi, lo, y uint32) uint32

// Rem64は、(hi, lo)をyで割った余りを返します。Rem64は
// y == 0の場合（ゼロ除算）にパニックしますが、Div64とは異なり、
// 商のオーバーフローではパニックしません。
>>>>>>> release-branch.go1.21
func Rem64(hi, lo, y uint64) uint64
