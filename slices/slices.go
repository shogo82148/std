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
// 空のスライスと nil スライスは等しいと見なされます。
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

// Insertは、値v...をsのインデックスiに挿入し、変更されたスライスを返します。
// s[i:]の要素は空きスペースを作るために後ろにずらされます。
// 返されたスライスrでは、r[i] == v[0]となり、
// また、i < len(s)の場合、r[i+len(v)]は元々r[i]にあった値になります。
// i > len(s)の場合、Insertはパニックを起こします。
// この関数の計算量はO(len(s) + len(v))です。
// 結果が空の場合、sのnil性を保持します。
func Insert[S ~[]E, E any](s S, i int, v ...E) S

// Deleteは、s[i:j]の要素をsから削除し、変更されたスライスを返します。
// j > len(s)の場合やs[i:j]がsの有効なスライスでない場合、Deleteはパニックを起こします。
// Deleteの計算量はO(len(s)-i)なので、多くの要素を削除する場合は、まとめて一度に削除する方が、一つずつ削除するより効率的です。
// Deleteはs[len(s)-(j-i):len(s)]の要素をゼロにします。
// 結果が空の場合、sのnil性を保持します。
func Delete[S ~[]E, E any](s S, i, j int) S

// DeleteFuncは、delがtrueを返すsの要素をすべて削除し、変更されたスライスを返します。
// DeleteFuncは新しい長さから元の長さまでの要素をゼロにします。
// 結果が空の場合、sのnil性を保持します。
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S

// Replaceは、s[i:j]の要素を指定されたvで置き換え、変更されたスライスを返します。
// j > len(s)の場合やs[i:j]がsの有効なスライスでない場合、Replaceはパニックを起こします。
// len(v) < (j-i)の場合、新しい長さから元の長さまでの要素をゼロにします。
// 結果が空の場合、sのnil性を保持します。
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S

// Cloneはスライスのコピーを返します。
// 要素は代入によってコピーされるため、これは浅いコピーです。
// 結果は追加の未使用容量を持つ場合があります。
// 結果はsのnil性を保持します。
func Clone[S ~[]E, E any](s S) S

// Compactは、連続する等しい要素の並びを1つのコピーに置き換えます。
// これはUnixのuniqコマンドのような動作です。
// Compactはスライスsの内容を変更し、変更されたスライスを返します。
// 長さが短くなる場合があります。
// Compactは新しい長さから元の長さまでの要素をゼロにします。
// 結果はsのnil性を保持します。
func Compact[S ~[]E, E comparable](s S) S

// CompactFuncは [Compact] と同様ですが、要素の比較に等値関数を使用します。
// 等しいと判定された並びのうち、CompactFuncは最初の要素を残します。
// CompactFuncは新しい長さから元の長さまでの要素をゼロにします。
// 結果はsのnil性を保持します。
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S

// Growは、必要に応じてスライスの容量を増やし、
// 追加でn個の要素分の空き容量を保証します。Grow(n)の後は、追加の割り当てなしで少なくともn個の要素をスライスに追加できます。
// nが負またはメモリ割り当てが大きすぎる場合、Growはパニックを起こします。
// 結果はsのnil性を保持します。
func Grow[S ~[]E, E any](s S, n int) S

// Clipはスライスの未使用容量を削除し、s[:len(s):len(s)]を返します。
// 結果はsのnil性を保持します。
func Clip[S ~[]E, E any](s S) S

// Reverse は、スライスの要素を逆順にします。
func Reverse[S ~[]E, E any](s S) { return }

// Concatは、渡されたスライスを連結した新しいスライスを返します。
// 連結結果が空の場合、結果はnilになります。
func Concat[S ~[]E, E any](slices ...S) S

// Repeatは、指定されたスライスを指定された回数だけ繰り返す新しいスライスを返します。
// 結果の長さと容量は (len(x) * count) です。
// 結果は決してnilになりません。
// countが負の場合、または(len(x) * count)の結果がオーバーフローする場合、Repeatはパニックを起こします。
func Repeat[S ~[]E, E any](x S, count int) S { return nil }
