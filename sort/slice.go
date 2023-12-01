// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort

// Sliceは与えられたless関数に基づいてスライスxをソートします。
// xがスライスでない場合はパニックを起こします。
//
// このソートは安定していることは保証されません：等しい要素は
// 元の順序から逆になる場合があります。
// 安定したソートをするには、[SliceStable] を使用してください。
//
// less関数は、Interface型のLessメソッドと同じ要件を満たす必要があります。
//
// 注意：多くの場合、より新しい [slices.SortFunc] 関数の方が操作性が高く、実行速度も速くなります。
func Slice(x any, less func(i, j int) bool)

// SliceStableは、与えられた比較関数を使用してスライスxをソートし、等しい要素を元の順序で保持します。
// xがスライスでない場合、パニックを起こします。
//
// less関数は、Interface型のLessメソッドと同じ要件を満たす必要があります。
//
// 注意：多くの場合、より新しい [slices.SortStableFunc] 関数の方が操作性が高く、実行速度も速くなります。
func SliceStable(x any, less func(i, j int) bool)

// SliceIsSortedは、提供されたless関数に従ってスライスxがソートされているかどうかを報告します。
// xがスライスでない場合、panicします。
//
// 注意：多くの場合、より新しい [slices.IsSortedFunc] 関数の方が操作性が高く、実行速度も速くなります。
func SliceIsSorted(x any, less func(i, j int) bool) bool
