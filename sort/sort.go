// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen_sort_variants.go

// Package sort はスライスやユーザー定義のコレクションをソートするための基本機能を提供します。
package sort

// このパッケージのルーチンによって、インタフェースの実装はソート可能です。
// メソッドは、整数インデックスによって基礎コレクションの要素を参照します。
type Interface interface {
	Len() int

	Less(i, j int) bool

	Swap(i, j int)
}

// SortはLessメソッドによって決定される昇順でデータをソートします。
// data.Lenを1回呼び出してnを決定し、O(n*log(n))回のdata.Lessとdata.Swapを呼び出します。ソートは安定するとは限りません。
//
// 注意：多くの場合、よりエルゴノミックで高速なslices.SortFunc関数を使用する方が好ましいです。
func Sort(data Interface)

// Reverseはdataの逆順を返します。
func Reverse(data Interface) Interface

// IsSortedはデータがソートされているかどうかを報告します。
//
// 注意：多くの場合、新しいslices.IsSortedFunc関数の方が使いやすく、高速です。
func IsSorted(data Interface) bool

// IntSlice は、Interface のメソッドを []int にアタッチし、昇順でソートします。
type IntSlice []int

func (x IntSlice) Len() int
func (x IntSlice) Less(i, j int) bool
func (x IntSlice) Swap(i, j int)

// Sort は便利なメソッドです: x.Sort() は Sort(x) を呼び出します。
func (x IntSlice) Sort()

// Float64Sliceは、[]float64のデータを増加順に並べ替えるためのインターフェースを実装します。
// NaN（非数値）の値は他の値よりも前に並べます。
type Float64Slice []float64

func (x Float64Slice) Len() int

// Less関数は、ソートインターフェースの要件に従って、x[i]がx[j]の前に並べられるべきかどうかを報告します。
// フロート比較自体は推移的な関係ではありませんことに注意してください：NaN（非数）の値については一貫した順序を報告しません。
// このLess関数の実装では、NaN値を他の値よりも前に配置します：
//
//	x[i] < x[j] || (math.IsNaN(x[i]) && !math.IsNaN(x[j]))
func (x Float64Slice) Less(i, j int) bool
func (x Float64Slice) Swap(i, j int)

// Sortは便利なメソッドです：x.Sort()はSort(x)を呼び出します。
func (x Float64Slice) Sort()

// StringSliceはInterfaceのメソッドを[]stringに追加し、昇順でソートします。
type StringSlice []string

func (x StringSlice) Len() int
func (x StringSlice) Less(i, j int) bool
func (x StringSlice) Swap(i, j int)

// Sort は利便性のためのメソッドです: x.Sort() は Sort(x) を呼び出します。
func (x StringSlice) Sort()

// Intsはintのスライスを昇順にソートします。
//
// 注意：Go 1.22以降、この関数は単にslices.Sortを呼び出すだけです。
func Ints(x []int)

// Float64sはfloat64のスライスを昇順でソートします。
// NaN(非数)の値は他の値よりも優先的に並べられます。
//
// 注意：Go 1.22以降、この関数は単にslices.Sortを呼び出すだけです。
func Float64s(x []float64)

// Stringsは文字列のスライスを昇順でソートします。
//
// 注意：Go 1.22以降、この関数は単にslices.Sortを呼び出すだけです。
func Strings(x []string)

// IntsAreSortedは、スライスxが昇順にソートされているかどうかを報告します。
//
// 注意：Go 1.22以降、この関数は単にslices.IsSortedを呼び出すだけです。
func IntsAreSorted(x []int) bool

// Float64sAreSortedは、スライスxが昇順に並んでいるか、NaN（非数値）の値が他の値の前にあるかどうかを報告します。
//
// 注意：Go 1.22以降、この関数は単にslices.IsSortedを呼び出すだけです。
func Float64sAreSorted(x []float64) bool

// StringsAreSortedは、スライスxが昇順に並んでいるかどうかを報告します。
//
// 注意：Go 1.22以降、この関数は単にslices.IsSortedを呼び出すだけです。
func StringsAreSorted(x []string) bool

// Lessメソッドによって決定される昇順でデータを安定的にソートします。
// 同じ要素の元の順序を保持します。
//
// data.Lenを1回呼び出してnを決定し、data.LessをO(n*log(n))回呼び出し、
// data.SwapをO(n*log(n)*log(n))回呼び出します。
//
// 注意: 多くの場合、新しいslices.SortStableFunc関数の方が使いやすく、
// より高速に実行されます。
func Stable(data Interface)
