// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tabwriter_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/text/tabwriter"
)

func ExampleWriter_Init() {
	w := new(tabwriter.Writer)

	// タブストップ8でタブ区切りの列にフォーマットします。
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()

	// 最小幅5のスペースで区切られた列に右揃えでフォーマットします。
	// そして少なくとも1つの空白のパディング（よって、より広い列のエントリーが
	// お互いに触れないようにします）。
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "a\tb\tc\td\t.")
	fmt.Fprintln(w, "123\t12345\t1234567\t123456789\t.")
	fmt.Fprintln(w)
	w.Flush()

	// output:
	// a	b	c	d		.
	// 123	12345	1234567	123456789	.
	//
	//     a     b       c         d.
	//   123 12345 1234567 123456789.
}

func Example_elastic() {
	// bとdがどちらも各行の2番目のセルに表示されているにもかかわらず、
	// 異なる列に属していることに注目してください。
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\tc")
	fmt.Fprintln(w, "aa\tbb\tcc")
	fmt.Fprintln(w, "aaa\t") // trailing tab
	fmt.Fprintln(w, "aaaa\tdddd\teeee")
	w.Flush()

	// output:
	// ....a|..b|c
	// ...aa|.bb|cc
	// ..aaa|
	// .aaaa|.dddd|eeee
}

func Example_trailingTab() {
	// 3行目には末尾のタブがないことに注意してください。
	// したがって、その最終セルは整列した列の一部ではありません。
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, '-', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "a\tb\taligned\t")
	fmt.Fprintln(w, "aa\tbb\taligned\t")
	fmt.Fprintln(w, "aaa\tbbb\tunaligned") // 末尾にタブがない
	fmt.Fprintln(w, "aaaa\tbbbb\taligned\t")
	w.Flush()

	// output:
	// ------a|------b|---aligned|
	// -----aa|-----bb|---aligned|
	// ----aaa|----bbb|unaligned
	// ---aaaa|---bbbb|---aligned|
}
