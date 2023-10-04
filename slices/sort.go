// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices

import (
	"github.com/shogo82148/std/cmp"
)

// Sort は、任意の順序付け可能な型のスライスを昇順にソートします。
// 浮動小数点数をソートする場合、NaN は他の値の前に並べられます。
func Sort[S ~[]E, E cmp.Ordered](x S) {}

// SortFunc は、cmp 関数によって決定される昇順でスライス x をソートします。
// このソートは安定であることは保証されません。
// cmp(a, b) は、a < b の場合は負の数、a > b の場合は正の数、a == b の場合はゼロを返す必要があります。
//
// SortFunc は、cmp が厳密な弱順序であることを要求します。
// https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings を参照してください。
func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {}

// SortStableFunc は、cmp 関数によって決定される順序でスライス x をソートします。
// このソートは安定であり、等しい要素の元の順序を維持します。
// cmp(a, b) は、a < b の場合は負の数、a > b の場合は正の数、a == b の場合はゼロを返す必要があります。
//
// SortStableFunc は、cmp が厳密な弱順序であることを要求します。
// https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings を参照してください。
func SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int) {}

// IsSorted は、x が昇順にソートされているかどうかを報告します。
func IsSorted[S ~[]E, E cmp.Ordered](x S) bool { return false }

// IsSortedFunc は、x が昇順にソートされているかどうかを報告します。
// このソートは、cmp 関数によって定義された比較関数によって行われます。
func IsSortedFunc[S ~[]E, E any](x S, cmp func(a, b E) int) bool { return false }

// Min は、x 内の最小値を返します。x が空の場合は panic を発生させます。
// 浮動小数点数の場合、Min は NaN を伝播します（x 内の任意の NaN 値は、出力を NaN にします）。
func Min[S ~[]E, E cmp.Ordered](x S) E { var zero E; return zero }

// MinFunc は、cmp 関数によって比較された x 内の最小値を返します。
// x が空の場合、MinFunc は panic を発生させます。
// cmp 関数によって複数の最小要素がある場合、MinFunc は最初の要素を返します。
func MinFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E { var zero E; return zero }

// Max は、x 内の最大値を返します。x が空の場合は panic を発生させます。
// 浮動小数点数の場合、Max は NaN を伝播します（x 内の任意の NaN 値は、出力を NaN にします）。
func Max[S ~[]E, E cmp.Ordered](x S) E { var zero E; return zero }

// MaxFunc は、cmp 関数によって比較された x 内の最大値を返します。
// x が空の場合、MaxFunc は panic を発生させます。
// cmp 関数によって複数の最大要素がある場合、MaxFunc は最初の要素を返します。
func MaxFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E { var zero E; return zero }

// BinarySearch は、ソートされたスライス内で target を検索し、target が見つかる位置、
// またはソート順序で target が表示される位置を返します。
// また、スライス内に target が本当に見つかったかどうかを示す bool も返します。
// スライスは昇順にソートする必要があります。
func BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool) { return 0, false }

<<<<<<< HEAD
// BinarySearchFunc は、カスタム比較関数を使用して [BinarySearch] と同様に動作します。
// スライスは、cmp によって定義される増加順でソートする必要があります。
// cmp は、スライス要素がターゲットに一致する場合は 0、スライス要素がターゲットよりも前の場合は負の数、
// スライス要素がターゲットよりも後ろの場合は正の数を返す必要があります。
// cmp は、スライスと同じ順序付けを実装する必要があります。つまり、
// cmp(a, t) < 0 かつ cmp(b, t) >= 0 の場合、a はスライス内で b よりも前になければなりません。
func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool) {
	return 0, false
}

// xorshift paper: https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
=======
// BinarySearchFunc works like [BinarySearch], but uses a custom comparison
// function. The slice must be sorted in increasing order, where "increasing"
// is defined by cmp. cmp should return 0 if the slice element matches
// the target, a negative number if the slice element precedes the target,
// or a positive number if the slice element follows the target.
// cmp must implement the same ordering as the slice, such that if
// cmp(a, t) < 0 and cmp(b, t) >= 0, then a must precede b in the slice.
func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool)
>>>>>>> upstream/release-branch.go1.21
