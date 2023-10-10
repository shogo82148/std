// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package subtleは、暗号化コードでよく使用される関数を実装しますが、正しく使用するために注意深い考慮が必要です。
package subtle

// ConstantTimeCompareは、2つのスライスxとyが同じ内容を持つ場合は1を返し、そうでない場合は0を返します。実行時間はスライスの長さに依存し、内容には独立しています。xとyの長さが一致しない場合は、即座に0を返します。
func ConstantTimeCompare(x, y []byte) int

// ConstantTimeSelectは、vが1の場合はxを返し、vが0の場合はyを返します。
// vが他の値を取る場合、動作は未定義です。
func ConstantTimeSelect(v, x, y int) int

// ConstantTimeByteEq は、x が y と等しい場合は1を、そうでない場合は0を返します。
func ConstantTimeByteEq(x, y uint8) int

// ConstantTimeEqは、x == yの場合は1を返し、それ以外の場合は0を返します。
func ConstantTimeEq(x, y int32) int

// ConstantTimeCopyは、v == 1の場合、yの内容（長さが等しいスライス）をxにコピーします。
// v == 0の場合、xは変更されません。vが他の値を取る場合の動作は未定義です。
func ConstantTimeCopy(v int, x, y []byte)

// ConstantTimeLessOrEq は、x ≤ y の場合は1を返し、そうでない場合は0を返します。
// ただし、xまたはyが負数または2**31 - 1より大きい場合、動作は未定義です。
func ConstantTimeLessOrEq(x, y int) int
