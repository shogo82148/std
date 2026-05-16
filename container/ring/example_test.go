// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ring_test

import (
	"github.com/shogo82148/std/container/ring"
	"github.com/shogo82148/std/fmt"
)

func ExampleRing_Len() {
	// サイズ4の新しいリングを作成します。
	r := ring.New(4)

	// その長さを出力します。
	fmt.Println(r.Len())

	// Output:
	// 4
}

func ExampleRing_Next() {
	// サイズ5の新しいリングを作成します。
	r := ring.New(5)

	// リングの長さを取得します。
	n := r.Len()

	// リングにいくつかの整数値を設定します。
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// リングを走査して内容を出力します。
	for j := 0; j < n; j++ {
		fmt.Println(r.Value)
		r = r.Next()
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleRing_Prev() {
	// サイズ5の新しいリングを作成します。
	r := ring.New(5)

	// リングの長さを取得します。
	n := r.Len()

	// リングにいくつかの整数値を設定します。
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// リングを逆方向に走査して内容を出力します。
	for j := 0; j < n; j++ {
		r = r.Prev()
		fmt.Println(r.Value)
	}

	// Output:
	// 4
	// 3
	// 2
	// 1
	// 0
}

func ExampleRing_Do() {
	// サイズ5の新しいリングを作成します。
	r := ring.New(5)

	// リングの長さを取得します。
	n := r.Len()

	// リングにいくつかの整数値を設定します。
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// リングを走査して内容を出力します。
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

func ExampleRing_Move() {
	// サイズ5の新しいリングを作成します。
	r := ring.New(5)

	// リングの長さを取得します。
	n := r.Len()

	// リングにいくつかの整数値を設定します。
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// ポインタを3ステップ進めます。
	r = r.Move(3)

	// リングを走査して内容を出力します。
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

	// Output:
	// 3
	// 4
	// 0
	// 1
	// 2
}

func ExampleRing_Link() {
	// サイズ2の2つのリングrとsを作成します。
	r := ring.New(2)
	s := ring.New(2)

	// リングの長さを取得します。
	lr := r.Len()
	ls := s.Len()

	// rを0で初期化します。
	for i := 0; i < lr; i++ {
		r.Value = 0
		r = r.Next()
	}

	// sを1で初期化します。
	for j := 0; j < ls; j++ {
		s.Value = 1
		s = s.Next()
	}

	// リングrとリングsを連結します。
	rs := r.Link(s)

	// 連結されたリングを走査して内容を出力します。
	rs.Do(func(p any) {
		fmt.Println(p.(int))
	})

	// Output:
	// 0
	// 0
	// 1
	// 1
}

func ExampleRing_Unlink() {
	// サイズ6の新しいリングを作成します。
	r := ring.New(6)

	// リングの長さを取得します。
	n := r.Len()

	// リングにいくつかの整数値を設定します。
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// rから3つの要素を切り離します。開始位置はr.Next()です。
	r.Unlink(3)

	// 残ったリングを走査して内容を出力します。
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

	// Output:
	// 0
	// 4
	// 5
}
