// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/crypto/rand"
	"github.com/shogo82148/std/fmt"
)

// この例は、rand.Readerから暗号的に安全な疑似乱数を10個読み込み、バイトスライスに書き込みます。
func ExampleRead() {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// スライスは、ゼロの代わりにランダムなバイトを含んでいるべきです。
	fmt.Println(bytes.Equal(b, make([]byte, c)))

	// Output:
	// false
}
