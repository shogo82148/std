// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/sort"
)

func Example() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)

	// スライスをソートする方法は2つあります。最初に、ByAgeなどのスライス型のためのメソッドセットを定義し、sort.Sortを呼び出すことができます。この最初の例では、その技術を使用しています。
	sort.Sort(ByAge(people))
	fmt.Println(people)

	// もう一つの方法は、カスタムのLess関数を使用してsort.Sliceを利用することです。
	// これはクロージャとして提供することができます。この場合、メソッドは必要ありません。
	// (存在する場合は無視されます。) ここでは逆順で再ソートします：クロージャとByAge.Lessを比較します。
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

	// Output:
	// [Bob: 31 John: 42 Michael: 17 Jenny: 26]
	// [Michael: 17 Jenny: 26 Bob: 31 John: 42]
	// [John: 42 Bob: 31 Jenny: 26 Michael: 17]
}
