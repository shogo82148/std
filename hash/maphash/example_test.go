// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maphash_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/hash/maphash"
)

func Example() {
	// ゼロのハッシュ値は有効で使用準備が整っています。初期シードの設定は必要ありません。
	var h maphash.Hash

	// 文字列をハッシュに追加し、現在のハッシュ値を表示します。
	h.WriteString("hello, ")
	fmt.Printf("%#x\n", h.Sum64())

	// 追加データをバイト配列の形式で追加します。
	h.Write([]byte{'w', 'o', 'r', 'l', 'd'})
	fmt.Printf("%#x\n", h.Sum64())

	// Resetは、以前にハッシュに追加されたすべてのデータを破棄しますが、
	// シードは変更しません。
	h.Reset()

	// 新しいハッシュh2を作成するために、SetSeedを使用し、hと完全に同じ動作をします。
	var h2 maphash.Hash
	h2.SetSeed(h.Seed())

	h.WriteString("same")
	h2.WriteString("same")
	fmt.Printf("%#x == %#x\n", h.Sum64(), h2.Sum64())
}
