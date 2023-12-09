// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/encoding/gob"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
)

// この例は、パッケージの基本的な使用方法を示しています: エンコーダを作成し、
// いくつかの値を送信し、デコーダでそれらを受信します。
func Example_basic() {
	// エンコーダとデコーダを初期化します。通常、encとdecはネットワーク接続にバインドされ、
	// エンコーダとデコーダは別々のプロセスで実行されます。
	var network bytes.Buffer        // ネットワーク接続の代わり
	enc := gob.NewEncoder(&network) // ネットワークに書き込みます。
	dec := gob.NewDecoder(&network) // ネットワークから読み取ります。

	// いくつかの値をエンコード（送信）します。
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err = enc.Encode(P{1782, 1841, 1922, "Treehouse"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// 値をデコード（受信）し、印刷します。
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 2:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)

	// Output:
	// "Pythagoras": {3, 4}
	// "Treehouse": {1782, 1841}
}
