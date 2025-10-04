// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand_test

import (
	"github.com/shogo82148/std/crypto/rand"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/math/big"
)

// ExampleIntは0から99までの範囲（境界値を含む）で暗号学的に安全な擬似乱数を1つ出力します。
func ExampleInt() {
	// rand.Readerを使用する場合、Intはエラーを返すことができません。
	a, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println(a.Int64())
}

// ExamplePrimeは暗号学的に安全な擬似乱数64ビット素数を出力します。
func ExamplePrime() {
	// rand.Readerを使用し、bits >= 2の場合、Primeはエラーを返すことができません。
	a, _ := rand.Prime(rand.Reader, 64)
	fmt.Println(a.Int64())
}

// ExampleReadは暗号学的に安全な擬似乱数32バイトキーを出力します。
func ExampleRead() {
	// Readは常に成功するため、エラーハンドリングは不要であることに注意してください。
	key := make([]byte, 32)
	rand.Read(key)
	// キーは任意のバイト値を含む可能性があるため、キーを16進数で出力します。
	fmt.Printf("% x\n", key)
}

// ExampleTextはbase32でエンコードされたランダムキーを出力します。
func ExampleText() {
	key := rand.Text()
	// キーはbase32で、表示しても安全です。
	fmt.Println(key)
}
