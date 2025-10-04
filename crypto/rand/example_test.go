// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand_test

import (
	"github.com/shogo82148/std/crypto/rand"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/math/big"
)

<<<<<<< HEAD
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
=======
// ExampleInt prints a single cryptographically secure pseudorandom number between 0 and 99 inclusive.
func ExampleInt() {
	// Int cannot return an error when using rand.Reader.
	a, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println(a.Int64())
}

// ExamplePrime prints a cryptographically secure pseudorandom 64 bit prime number.
func ExamplePrime() {
	// Prime cannot return an error when using rand.Reader and bits >= 2.
	a, _ := rand.Prime(rand.Reader, 64)
	fmt.Println(a.Int64())
}

// ExampleRead prints a cryptographically secure pseudorandom 32 byte key.
func ExampleRead() {
	// Note that no error handling is necessary, as Read always succeeds.
	key := make([]byte, 32)
	rand.Read(key)
	// The key can contain any byte value, print the key in hex.
	fmt.Printf("% x\n", key)
}

// ExampleText prints a random key encoded in base32.
func ExampleText() {
	key := rand.Text()
	// The key is base32 and safe to display.
	fmt.Println(key)
>>>>>>> upstream/release-branch.go1.25
}
