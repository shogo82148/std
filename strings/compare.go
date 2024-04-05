// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

// Compareは、2つの文字列を辞書順で比較して整数を返します。
// a == bの場合は0、a < bの場合は-1、a > bの場合は+1になります。
//
<<<<<<< HEAD
// Compareは、パッケージbytesとの対称性のために含まれています。
// 通常、組み込みの文字列比較演算子==、<、>などを使用する方が明確で、常に高速です。
=======
// Use Compare when you need to perform a three-way comparison (with
// slices.SortFunc, for example). It is usually clearer and always faster
// to use the built-in string comparison operators ==, <, >, and so on.
>>>>>>> upstream/master
func Compare(a, b string) int
