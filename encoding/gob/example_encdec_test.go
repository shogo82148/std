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

// この例では、カスタムのエンコーディングとデコーディングメソッドを実装した値を伝送します。
func Example_encodeDecode() {
	var network bytes.Buffer // ネットワークの代わり。

	// エンコーダを作成し、値を送信します。
	enc := gob.NewEncoder(&network)
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatal("encode:", err)
	}

	// デコーダを作成し、値を受信します。
	dec := gob.NewDecoder(&network)
	var v Vector
	err = dec.Decode(&v)
	if err != nil {
		log.Fatal("decode:", err)
	}
	fmt.Println(v)

	// Output:
	// {3 4 5}
}
