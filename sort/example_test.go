// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/math"
	"github.com/shogo82148/std/sort"
)

func ExampleInts() {
	s := []int{5, 2, 6, 3, 1, 4} // ソートされていない
	sort.Ints(s)
	fmt.Println(s)
	// Output: [1 2 3 4 5 6]
}

func ExampleIntsAreSorted() {
	s := []int{1, 2, 3, 4, 5, 6} // 昇順でソートされています
	fmt.Println(sort.IntsAreSorted(s))

	s = []int{6, 5, 4, 3, 2, 1} // 降順で並べ替え済み
	fmt.Println(sort.IntsAreSorted(s))

	s = []int{3, 2, 4, 1, 5} // 未ソート
	fmt.Println(sort.IntsAreSorted(s))

	// Output: true
	// false
	// false
}

func ExampleFloat64s() {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // ソートされていない
	sort.Float64s(s)
	fmt.Println(s)

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // 未整列
	sort.Float64s(s)
	fmt.Println(s)

	// Output: [-3.8 -1.3 0.7 2.6 5.2]
	// [NaN -Inf 0 +Inf]
}

func ExampleFloat64sAreSorted() {
	s := []float64{0.7, 1.3, 2.6, 3.8, 5.2} // 昇順でソートされています
	fmt.Println(sort.Float64sAreSorted(s))

	s = []float64{5.2, 3.8, 2.6, 1.3, 0.7} // 降順でソート済み
	fmt.Println(sort.Float64sAreSorted(s))

	s = []float64{5.2, 1.3, 0.7, 3.8, 2.6} // 未整列
	fmt.Println(sort.Float64sAreSorted(s))

	// Output: true
	// false
	// false
}

func ExampleReverse() {
	s := []int{5, 2, 6, 3, 1, 4} // ソートされていない
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
	// Output: [6 5 4 3 2 1]
}

func ExampleSlice() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
	// Output: By name: [{Alice 55} {Bob 75} {Gopher 7} {Vera 24}]
	// By age: [{Gopher 7} {Vera 24} {Alice 55} {Bob 75}]
}

func ExampleSliceStable() {

	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75},
		{"Bob", 75},
		{"Alice", 75},
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}

	// 名前でソートし、元の順序を保持します
	sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	// 名前の順序を保持しつつ年齢でソートする
	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age,name:", people)

	// Output: By name: [{Alice 25} {Alice 75} {Alice 75} {Bob 75} {Bob 25} {Colin 25} {Elizabeth 75} {Elizabeth 25}]
	// By age,name: [{Alice 25} {Bob 25} {Colin 25} {Elizabeth 25} {Alice 75} {Alice 75} {Bob 75} {Elizabeth 75}]
}

func ExampleStrings() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)
	// Output: [Alpha Bravo Delta Go Gopher Grin]
}
