// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// slices パッケージは、任意の型のスライスで使用できるさまざまな関数を定義します。
package slices

import (
	"github.com/shogo82148/std/cmp"
)

// Equal は、2つのスライスが等しいかどうかを報告します。
// 長さが異なる場合、Equal は false を返します。
// そうでない場合、要素はインデックスの昇順で比較され、比較は最初の異なるペアで停止します。
// 浮動小数点 NaN は等しくないと見なされます。
func Equal[S ~[]E, E comparable](s1, s2 S) bool { return false }

// EqualFunc は、各要素のペアに対して等値性関数を使用して2つのスライスが等しいかどうかを報告します。
// 長さが異なる場合、EqualFunc は false を返します。
// そうでない場合、要素はインデックスの昇順で比較され、比較は最初の異なるペアで停止します。
func EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool { return false }

// Compare は、s1 と s2 の要素を、各ペアで [cmp.Compare] を使用して比較します。
// 要素は、インデックス 0 から順番に比較され、一方の要素が他方と等しくなくなるまで比較が続けられます。
// 最初の不一致要素の比較結果が返されます。
// どちらかのスライスが終端に達するまで、両方のスライスが等しい場合、短いスライスが長いスライスよりも小さいと見なされます。
// 結果は、s1 == s2 の場合は 0、s1 < s2 の場合は -1、s1 > s2 の場合は +1 です。
func Compare[S ~[]E, E cmp.Ordered](s1, s2 S) int { return 0 }

// CompareFunc は、各要素のペアに対してカスタム比較関数を使用して s1 と s2 の要素を比較します。
// 結果は、cmp の最初の非ゼロ結果です。cmp が常に 0 を返す場合、結果は len(s1) == len(s2) の場合は 0、len(s1) < len(s2) の場合は -1、len(s1) > len(s2) の場合は +1 です。
func CompareFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, cmp func(E1, E2) int) int {
	return 0
}

// Index は、s 内で最初に v が出現するインデックスを返します。
// v が存在しない場合は -1 を返します。
func Index[S ~[]E, E comparable](s S, v E) int { return 0 }

// IndexFunc は、f(s[i]) を満たす最初のインデックス i を返します。
// そのようなインデックスが存在しない場合は -1 を返します。
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int { return 0 }

// Contains は、s 内に v が存在するかどうかを報告します。
func Contains[S ~[]E, E comparable](s S, v E) bool { return false }

// ContainsFunc は、s の少なくとも1つの要素 e が f(e) を満たすかどうかを報告します。
func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool { return false }

// Insert は、値 v... を s のインデックス i に挿入し、変更されたスライスを返します。
// s[i:] の要素は上にシフトされ、スペースが作成されます。
// 返されるスライス r では、r[i] == v[0] であり、r[i+len(v)] == r[i] に元々あった値です。
// i が範囲外の場合、Insert は panic を発生させます。
// この関数の計算量は O(len(s) + len(v)) です。
func Insert[S ~[]E, E any](s S, i int, v ...E) S { return nil }

<<<<<<< HEAD
// Delete は、s[i:j] の要素を削除し、変更されたスライスを返します。
// s[i:j] が s の有効なスライスでない場合、Delete は panic を発生させます。
// Delete は O(len(s)-j) であり、多数のアイテムを削除する必要がある場合は、
// 1つずつ削除するよりも、まとめて削除する単一の呼び出しを作成する方が良いです。
// Delete は、s[len(s)-(j-i):len(s)] の要素を変更しない場合があります。
// これらの要素にポインタが含まれている場合は、参照するオブジェクトをガベージコレクションできるように、
// これらの要素をゼロにすることを検討する必要があります。
func Delete[S ~[]E, E any](s S, i, j int) S { return nil }
=======
// Delete removes the elements s[i:j] from s, returning the modified slice.
// Delete panics if j > len(s) or s[i:j] is not a valid slice of s.
// Delete is O(len(s)-i), so if many items must be deleted, it is better to
// make a single call deleting them all together than to delete one at a time.
// Delete zeroes the elements s[len(s)-(j-i):len(s)].
func Delete[S ~[]E, E any](s S, i, j int) S
>>>>>>> upstream/master

<<<<<<< HEAD
// DeleteFunc は、del が true を返す要素を s から削除し、変更されたスライスを返します。
// DeleteFunc が m 個の要素を削除すると、s[len(s)-m:len(s)] の要素を変更しない場合があります。
// これらの要素にポインタが含まれている場合は、参照するオブジェクトをガベージコレクションできるように、
// これらの要素をゼロにすることを検討する必要があります。
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S { return nil }
=======
// DeleteFunc removes any elements from s for which del returns true,
// returning the modified slice.
// DeleteFunc zeroes the elements between the new length and the original length.
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S
>>>>>>> upstream/master

<<<<<<< HEAD
// Replace は、s[i:j] の要素を与えられた v で置き換え、変更されたスライスを返します。
// s[i:j] が s の有効なスライスでない場合、Replace は panic を発生させます。
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S { return nil }
=======
// Replace replaces the elements s[i:j] by the given v, and returns the
// modified slice.
// Replace panics if j > len(s) or s[i:j] is not a valid slice of s.
// When len(v) < (j-i), Replace zeroes the elements between the new length and the original length.
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S
>>>>>>> upstream/master

// Clone は、スライスのコピーを返します。
// 要素は代入を使用してコピーされるため、これは浅いクローンです。
func Clone[S ~[]E, E any](s S) S { return nil }

<<<<<<< HEAD
// Compact は、等しい要素の連続したランを単一のコピーに置き換えます。
// これは、Unix で見つかる uniq コマンドのようです。
// Compact はスライス s の内容を変更し、変更されたスライスを返します。
// 変更されたスライスの長さが短くなる場合があります。
// Compact が合計で m 個の要素を破棄すると、s[len(s)-m:len(s)] の要素を変更しない場合があります。
// これらの要素にポインタが含まれている場合は、参照するオブジェクトをガベージコレクションできるように、
// これらの要素をゼロにすることを検討する必要があります。
func Compact[S ~[]E, E comparable](s S) S { return nil }

// CompactFunc は、等価性関数を使用して要素を比較する Compact です。
// 等しい要素のランに対して、CompactFunc は最初の要素を保持します。
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S { return nil }
=======
// Compact replaces consecutive runs of equal elements with a single copy.
// This is like the uniq command found on Unix.
// Compact modifies the contents of the slice s and returns the modified slice,
// which may have a smaller length.
// Compact zeroes the elements between the new length and the original length.
func Compact[S ~[]E, E comparable](s S) S

// CompactFunc is like [Compact] but uses an equality function to compare elements.
// For runs of elements that compare equal, CompactFunc keeps the first one.
// CompactFunc zeroes the elements between the new length and the original length.
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S
>>>>>>> upstream/master

// Grow は、必要に応じてスライスの容量を増やし、別の n 要素のスペースを保証します。
// Grow(n) の後、スライスには、別の割り当てなしで n 要素が追加できます。
// n が負の場合、またはメモリを割り当てるには大きすぎる場合、Grow は panic を発生させます。
func Grow[S ~[]E, E any](s S, n int) S { return nil }

// Clip は、スライスから未使用の容量を削除し、s[:len(s):len(s)] を返します。
func Clip[S ~[]E, E any](s S) S { return nil }

// Reverse は、スライスの要素を逆順にします。
func Reverse[S ~[]E, E any](s S) { return }

// Concat は、渡されたスライスを連結した新しいスライスを返します。
func Concat[S ~[]E, E any](slices ...S) S { return nil }
