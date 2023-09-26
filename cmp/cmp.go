// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージcmpは、順序付けられた値を比較するための型と関数に関連するものを提供します。
package cmp

// Ordered < <= >= >演算子をサポートする任意の順序付けられた型を許可する制約です。
// もし将来のGoリリースで新しい順序付けられた型が追加された場合、
// この制約はそれらを含めるように変更されます。
//
// 浮動小数点型にはNaN（「非数」）値が含まれる場合があります。
// NaN値と他の値、NaNであろうとなかろうと、比較演算子（==、<など）は常にfalseを報告します。
// NaN値を比較する一貫した方法については、 [Compare] 関数を参照してください。
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Less xがyより小さい場合にtrueを報告します。
// 浮動小数点型にはNaN（「非数」）値が含まれる場合があります。
// NaN値は、任意の非NaN値よりも小さいと見なされ、-0.0は0.0より小さくありません（等しいです）。
func Less[T Ordered](x, y T) bool { return false }

// Compareは、
//
//	xがyより小さい場合は-1、
//	xがyと等しい場合は0、
//	xがyより大きい場合は+1を返します。
//
// 浮動小数点型にはNaN（「非数」）値が含まれる場合があります。
// NaN値は、任意の非NaN値よりも小さいと見なされ、NaN値はNaN値と等しく、-0.0は0.0と等しいです。
func Compare[T Ordered](x, y T) int { return 0 }
