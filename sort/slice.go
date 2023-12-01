// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort

// Sliceは与えられたless関数に基づいてスライスxをソートします。
// xがスライスでない場合はパニックを起こします。
//
<<<<<<< HEAD
// このソートは安定していることは保証されません：等しい要素は
// 元の順序から逆になる場合があります。
// 安定したソートをするには、SliceStableを使用してください。
=======
// The sort is not guaranteed to be stable: equal elements
// may be reversed from their original order.
// For a stable sort, use [SliceStable].
>>>>>>> upstream/master
//
// less関数は、Interface型のLessメソッドと同じ要件を満たす必要があります。
//
<<<<<<< HEAD
// 注意：多くの場合、より新しいslices.SortFunc関数の方が操作性が高く、実行速度も速くなります。
=======
// Note: in many situations, the newer [slices.SortFunc] function is more
// ergonomic and runs faster.
>>>>>>> upstream/master
func Slice(x any, less func(i, j int) bool)

// SliceStableは、与えられた比較関数を使用してスライスxをソートし、等しい要素を元の順序で保持します。
// xがスライスでない場合、パニックを起こします。
//
// less関数は、Interface型のLessメソッドと同じ要件を満たす必要があります。
//
<<<<<<< HEAD
// 注意：多くの場合、より新しいslices.SortStableFunc関数の方が操作性が高く、実行速度も速くなります。
=======
// Note: in many situations, the newer [slices.SortStableFunc] function is more
// ergonomic and runs faster.
>>>>>>> upstream/master
func SliceStable(x any, less func(i, j int) bool)

// SliceIsSortedは、提供されたless関数に従ってスライスxがソートされているかどうかを報告します。
// xがスライスでない場合、panicします。
//
<<<<<<< HEAD
// 注意：多くの場合、より新しいslices.IsSortedFunc関数の方が操作性が高く、実行速度も速くなります。
=======
// Note: in many situations, the newer [slices.IsSortedFunc] function is more
// ergonomic and runs faster.
>>>>>>> upstream/master
func SliceIsSorted(x any, less func(i, j int) bool) bool
