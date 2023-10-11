// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package math

// Float32bitsは、fのIEEE 754バイナリ表現を返します。
// fの符号ビットと結果は同じビット位置にあります。
// Float32bits(Float32frombits(x)) == x.
func Float32bits(f float32) uint32

// Float32frombitsは、符号ビットの位置が同じであるように
// IEEE 754バイナリ表現bに対応する浮動小数点数を返します。
// Float32frombits（Float32bits（x））== x。
func Float32frombits(b uint32) float32

// Float64bitsは、fのIEEE 754バイナリ表現を返します。
// fの符号ビットと結果が同じビット位置になります。
// また、Float64bits(Float64frombits(x)) == x となります。
func Float64bits(f float64) uint64

// Float64frombitsは、IEEE 754のバイナリ表現bに対応する浮動小数点数を返します。bの符号ビットと結果は同じビット位置にあります。
// Float64frombits(Float64bits(x)) == x.
func Float64frombits(b uint64) float64
