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

// Deleteは、sから要素s[i:j]を削除し、変更されたスライスを返します。
// j > len(s)またはs[i:j]がsの有効なスライスでない場合、Deleteはパニックを引き起こします。
// DeleteはO(len(s)-i)なので、多くのアイテムを削除する必要がある場合、
// 一度に一つずつ削除するよりも、一度に全てを削除するための単一の呼び出しを行う方が良いです。
// Deleteは要素s[len(s)-(j-i):len(s)]をゼロにします。
func Delete[S ~[]E, E any](s S, i, j int) S { return nil }

// DeleteFuncは、delがtrueを返すsの任意の要素を削除し、
// 変更されたスライスを返します。
// DeleteFuncは新しい長さと元の長さの間の要素をゼロにします。
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S { return nil }

// Replaceは、要素s[i:j]を与えられたvで置き換え、
// 変更されたスライスを返します。
// Replaceは、j > len(s)またはs[i:j]がsの有効なスライスでない場合、パニックを引き起こします。
// len(v) < (j-i)の場合、Replaceは新しい長さと元の長さの間の要素をゼロにします。
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S { return nil }

// Clone は、スライスのコピーを返します。
// 要素は代入を使用してコピーされるため、これは浅いクローンです。
// 結果には追加の未使用の容量が含まれる可能性があります。
func Clone[S ~[]E, E any](s S) S { return nil }

// Compactは、等しい要素の連続した実行を単一のコピーで置き換えます。
// これはUnixにあるuniqコマンドのようなものです。
// Compactはスライスsの内容を変更し、変更されたスライスを返します。
// これはより小さい長さを持つ可能性があります。
// Compactは新しい長さと元の長さの間の要素をゼロにします。
func Compact[S ~[]E, E comparable](s S) S { return nil }

// CompactFuncは [Compact] と似ていますが、要素を比較するための等価性関数を使用します。
// 等しいと比較される要素の連続について、CompactFuncは最初のものを保持します。
// CompactFuncは新しい長さと元の長さの間の要素をゼロにします。
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S { return nil }

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

// Repeat returns a new slice that repeats the provided slice the given number of times.
// The result has length and capacity len(x) * count.
// The result is never nil.
// Repeat panics if count is negative or if the result of (len(x) * count)
// overflows.
func Repeat[S ~[]E, E any](x S, count int) S { return nil }
