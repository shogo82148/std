// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt_test

import (
	"github.com/shogo82148/std/fmt"
)

func ExampleGoStringer() {
	p1 := Person{
		Name: "Warren",
		Age:  31,
		Addr: &Address{
			City:    "Denver",
			State:   "CO",
			Country: "U.S.A.",
		},
	}

	// GoString()が実装されていない場合、`fmt.Printf("%#v", p1)`の出力は次のようになります。
	// Person{Name:"Warren", Age:0x1f, Addr:(*main.Address)(0x10448240)}
	fmt.Printf("%#v\n", p1)

	p2 := Person{
		Name: "Theia",
		Age:  4,
	}

	// GoString()が実装されていなかった場合、`fmt.Printf("%#v", p2)`の出力は以下のようになります
	// Person{Name:"Theia", Age:0x4, Addr:(*main.Address)(nil)}
	fmt.Printf("%#v\n", p2)

	// Output:
	// Person{Name: "Warren", Age: 31, Addr: &Address{City: "Denver", State: "CO", Country: "U.S.A."}}
	// Person{Name: "Theia", Age: 4}
}
