// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmp_test

import (
	"github.com/shogo82148/std/cmp"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/slices"
)

func ExampleOr() {
	// Suppose we have some user input
	// that may or may not be an empty string
	userInput1 := ""
	userInput2 := "some text"

	fmt.Println(cmp.Or(userInput1, "default"))
	fmt.Println(cmp.Or(userInput2, "default"))
	fmt.Println(cmp.Or(userInput1, userInput2, "default"))
	// Output:
	// default
	// some text
	// some text
}

func ExampleOr_sort() {
	type Order struct {
		Product  string
		Customer string
		Price    float64
	}
	orders := []Order{
		{"foo", "alice", 1.00},
		{"bar", "bob", 3.00},
		{"baz", "carol", 4.00},
		{"foo", "alice", 2.00},
		{"bar", "carol", 1.00},
		{"foo", "bob", 4.00},
	}
	// Sort by customer first, product second, and last by higher price
	slices.SortFunc(orders, func(a, b Order) int {
		return cmp.Or(
			cmp.Compare(a.Customer, b.Customer),
			cmp.Compare(a.Product, b.Product),
			cmp.Compare(b.Price, a.Price),
		)
	})
	for _, order := range orders {
		fmt.Printf("%s %s %.2f\n", order.Product, order.Customer, order.Price)
	}

	// Output:
	// foo alice 2.00
	// foo alice 1.00
	// bar bob 3.00
	// foo bob 4.00
	// bar carol 1.00
	// baz carol 4.00
}
