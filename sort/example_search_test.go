// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/sort"
)

// この例は昇順でソートされたリストの検索を示しています。
func ExampleSearch() {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
	// Output:
	// found 6 at index 2 in [1 3 6 10 15 21 28 36 45 55]
}

// この例では、降順でソートされたリストを検索する方法が示されています。
// アプローチは昇順でリストを検索するのと同じですが、条件が逆転しています。
func ExampleSearch_descendingOrder() {
	a := []int{55, 45, 36, 28, 21, 15, 10, 6, 3, 1}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] <= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
	// Output:
	// found 6 at index 7 in [55 45 36 28 21 15 10 6 3 1]
}

// この例は、昇順に並べられたリストで float64 を検索する方法を示しています。
func ExampleSearchFloat64s() {
	a := []float64{1.0, 2.0, 3.3, 4.6, 6.1, 7.2, 8.0}

	x := 2.0
	i := sort.SearchFloat64s(a, x)
	fmt.Printf("found %g at index %d in %v\n", x, i, a)

	x = 0.5
	i = sort.SearchFloat64s(a, x)
	fmt.Printf("%g not found, can be inserted at index %d in %v\n", x, i, a)
	// Output:
	// found 2 at index 1 in [1 2 3.3 4.6 6.1 7.2 8]
	// 0.5 not found, can be inserted at index 0 in [1 2 3.3 4.6 6.1 7.2 8]
}

// この例では、昇順に並べられたリスト内でintを検索する方法を示しています。
func ExampleSearchInts() {
	a := []int{1, 2, 3, 4, 6, 7, 8}

	x := 2
	i := sort.SearchInts(a, x)
	fmt.Printf("found %d at index %d in %v\n", x, i, a)

	x = 5
	i = sort.SearchInts(a, x)
	fmt.Printf("%d not found, can be inserted at index %d in %v\n", x, i, a)
	// Output:
	// found 2 at index 1 in [1 2 3 4 6 7 8]
	// 5 not found, can be inserted at index 4 in [1 2 3 4 6 7 8]
}
