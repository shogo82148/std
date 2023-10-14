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
<<<<<<< HEAD
// The less function must satisfy the same requirements as
// the Interface type's Less method.
//
// Note: in many situations, the newer slices.SortFunc function is more
// ergonomic and runs faster.
=======
// less関数は、Interface型のLessメソッドと同様の要件を満たす必要があります。
>>>>>>> release-branch.go1.21
func Slice(x any, less func(i, j int) bool)

// SliceStableは、与えられた比較関数を使用してスライスxをソートし、等しい要素を元の順序で保持します。
// xがスライスでない場合、パニックを起こします。
//
<<<<<<< HEAD
// The less function must satisfy the same requirements as
// the Interface type's Less method.
//
// Note: in many situations, the newer slices.SortStableFunc function is more
// ergonomic and runs faster.
func SliceStable(x any, less func(i, j int) bool)

// SliceIsSorted reports whether the slice x is sorted according to the provided less function.
// It panics if x is not a slice.
//
// Note: in many situations, the newer slices.IsSortedFunc function is more
// ergonomic and runs faster.
=======
// 比較関数は、Interface型のLessメソッドと同じ要件を満たす必要があります。
func SliceStable(x any, less func(i, j int) bool)

// SliceIsSortedは、与えられたless関数に従ってスライスxがソートされているかどうかを報告します。
// スライスxがスライスでない場合、パニックを起こします。
>>>>>>> release-branch.go1.21
func SliceIsSorted(x any, less func(i, j int) bool) bool
