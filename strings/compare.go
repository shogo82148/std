// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

// Compareは、2つの文字列を辞書順で比較して整数を返します。
// a == bの場合は0、a < bの場合は-1、a > bの場合は+1になります。
//
// 3つの比較を行う必要があるとき（例えば、slices.SortFuncと一緒に）にCompareを使用します。
// 通常、組み込みの文字列比較演算子 ==、<、>などを使用する方が明確で、常に高速です。
func Compare(a, b string) int
