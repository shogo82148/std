// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/encoding/gob"
	"github.com/shogo82148/std/fmt"
)

// この例では、インターフェース値のエンコード方法を示します。通常の型との主な違いは、
// インターフェースを実装する具体的な型を登録することです。
func Example_interface() {
	var network bytes.Buffer // ネットワークの代わり。

	// エンコーダとデコーダ（通常はエンコーダとは別のマシン上）に具体的な型を登録する必要があります。
	// それぞれの端では、これがどの具体的な型がインターフェースを実装して送信されているかをエンジンに伝えます。
	gob.Register(Point{})

	// エンコーダを作成し、いくつかの値を送信します。
	enc := gob.NewEncoder(&network)
	for i := 1; i <= 3; i++ {
		interfaceEncode(enc, Point{3 * i, 4 * i})
	}

	// デコーダを作成し、いくつかの値を受信します。
	dec := gob.NewDecoder(&network)
	for i := 1; i <= 3; i++ {
		result := interfaceDecode(dec)
		fmt.Println(result.Hypotenuse())
	}

	// Output:
	// 5
	// 10
	// 15
}
