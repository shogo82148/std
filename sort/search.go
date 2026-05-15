// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements binary search.

package sort

// Searchは二分探索を使って、[0, n) の範囲で f(i) が true となる
// 最小のインデックス i を見つけて返します。ここで [0, n) の範囲において
// f(i) == true なら f(i+1) == true であると仮定します。つまり Search は、
// 入力範囲 [0, n) の先頭側の（空の場合もある）接頭部分で f が false、
// 残りの（空の場合もある）部分で true となることを要求し、
// 最初に true となるインデックスを返します。そのようなインデックスが
// ない場合、Search は n を返します。
// （たとえば strings.Index のように、"見つからない" 場合の戻り値が
// -1 ではない点に注意してください。）
// Search は [0, n) の範囲の i に対してのみ f(i) を呼び出します。
//
// Search の一般的な用途は、配列やスライスのようなソート済みで
// インデックスアクセス可能なデータ構造において、値 x のインデックス i を
// 見つけることです。この場合、引数 f（通常はクロージャ）は、
// 探索対象の値と、データ構造のインデックス付けおよび順序付けの方法を
// 捕捉します。
//
// たとえば、昇順にソートされたスライス data があるとき、
// Search(len(data), func(i int) bool { return data[i] >= 23 })
// の呼び出しは data[i] >= 23 となる最小のインデックス i を返します。
// 呼び出し側が 23 がスライス内に存在するかを調べたい場合は、
// data[i] == 23 を別途確認する必要があります。
//
// 降順にソートされたデータを探索する場合は、>= 演算子の代わりに
// <= 演算子を使います。
//
// 上の例を完成させると、次のコードは昇順にソートされた整数スライス data から
// 値 x を探します。
//
//	x := 23
//	i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
//	if i < len(data) && data[i] == x {
//		// x は data[i] に存在する
//	} else {
//		// x は data には存在しないが、
//		// i は挿入される位置のインデックスである。
//	}
//
// もう少し遊び心のある例として、次のプログラムはあなたの数を当てます。
//
//	func GuessingGame() {
//		var s string
//		fmt.Printf("0 から 100 までの整数を 1 つ選んでください。\n")
//		answer := sort.Search(100, func(i int) bool {
//			fmt.Printf("あなたの数は %d 以下ですか? ", i)
//			fmt.Scanf("%s", &s)
//			return s != "" && s[0] == 'y'
//		})
//		fmt.Printf("あなたの数は %d です。\n", answer)
//	}
func Search(n int, f func(int) bool) int

// Findは二分探索を使って、[0, n) の範囲で cmp(i) <= 0 となる
// 最小のインデックス i を見つけて返します。そのような i がない場合、
// Find は i = n を返します。
// i < n かつ cmp(i) == 0 のとき、found の結果は true です。
// Find は [0, n) の範囲の i に対してのみ cmp(i) を呼び出します。
//
// 二分探索を可能にするため、Find は範囲の先頭側の接頭部分で cmp(i) > 0、
// 中央で cmp(i) == 0、末尾側の接尾部分で cmp(i) < 0 であることを
// 要求します。（各部分範囲は空でも構いません。）
// この条件を満たす一般的な方法は、cmp(i) を、
// 目的のターゲット値 t と基礎となるインデックス付きデータ構造 x の
// 要素 i との比較として解釈し、
// それぞれ t < x[i]、t == x[i]、t > x[i] のときに
// <0、0、>0 を返すようにすることです。
//
// たとえば、ソート済みでランダムアクセス可能な文字列リストから
// 特定の文字列を探すには次のようにします。
//
//	i, found := sort.Find(x.Len(), func(i int) int {
//	    return strings.Compare(target, x.At(i))
//	})
//	if found {
//	    fmt.Printf("%s がエントリ %d で見つかりました\n", target, i)
//	} else {
//	    fmt.Printf("%s は見つかりませんでした。挿入位置は %d です", target, i)
//	}
func Find(n int, cmp func(int) int) (i int, found bool)

// SearchIntsはソート済みのintスライスから x を探索し、[Search] で
// 指定されるインデックスを返します。戻り値は x が存在しない場合に x を
// 挿入するインデックスです（len(a) になる場合があります）。
// スライスは昇順にソートされている必要があります。
func SearchInts(a []int, x int) int

// SearchFloat64sはソート済みのfloat64スライスから x を探索し、[Search] で
// 指定されるインデックスを返します。戻り値は x が存在しない場合に x を
// 挿入するインデックスです（len(a) になる場合があります）。
// スライスは昇順にソートされている必要があります。
func SearchFloat64s(a []float64, x float64) int

// SearchStringsはソート済みのstringスライスから x を探索し、Search で
// 指定されるインデックスを返します。戻り値は x が存在しない場合に x を
// 挿入するインデックスです（len(a) になる場合があります）。
// スライスは昇順にソートされている必要があります。
func SearchStrings(a []string, x string) int

// Searchはレシーバと x に [SearchInts] を適用した結果を返します。
func (p IntSlice) Search(x int) int

// Searchはレシーバと x に [SearchFloat64s] を適用した結果を返します。
func (p Float64Slice) Search(x float64) int

// Searchはレシーバと x に [SearchStrings] を適用した結果を返します。
func (p StringSlice) Search(x string) int
