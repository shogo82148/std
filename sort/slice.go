// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort

// Sliceは与えられたless関数に基づいてスライスxをソートします。
// xがスライスでない場合はパニックを起こします。
//
// このソートは安定していることは保証されません：等しい要素は
// 元の順序から逆になる場合があります。
// 安定したソートをするには、SliceStableを使用してください。
//
// less関数は、Interface型のLessメソッドと同様の要件を満たす必要があります。
func Slice(x any, less func(i, j int) bool)

// SliceStableは、与えられた比較関数を使用してスライスxをソートし、等しい要素を元の順序で保持します。
// xがスライスでない場合、パニックを起こします。
//
// 比較関数は、Interface型のLessメソッドと同じ要件を満たす必要があります。
func SliceStable(x any, less func(i, j int) bool)

// SliceIsSortedは、与えられたless関数に従ってスライスxがソートされているかどうかを報告します。
// スライスxがスライスでない場合、パニックを起こします。
func SliceIsSorted(x any, less func(i, j int) bool) bool
