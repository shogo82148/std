// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !fips140v1.0

package mldsa_test

import (
	"github.com/shogo82148/std/crypto/mldsa"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
)

func Example() {
	// 署名者は新しいML-DSA-44鍵ペアを生成します。
	sk, err := mldsa.GenerateKey(mldsa.MLDSA44())
	if err != nil {
		log.Fatal(err)
	}

	// 署名者は公開鍵のエンコーディングを公開します。
	publicKey := sk.PublicKey().Bytes()
	fmt.Printf("public key: %d bytes\n", len(publicKey))

	// 署名者はメッセージに署名し、その署名を公開します。
	msg := []byte("hello, world")
	sig, err := sk.Sign(nil, msg, &mldsa.Options{Context: "example"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("signature: %d bytes\n", len(sig))

	// 検証者は公開鍵を再構築し、署名を検証します。
	// context文字列は署名者が使用したものと一致している必要があります。
	pk, err := mldsa.NewPublicKey(mldsa.MLDSA44(), publicKey)
	if err != nil {
		log.Fatal(err)
	}
	if err := mldsa.Verify(pk, msg, sig, &mldsa.Options{Context: "example"}); err != nil {
		log.Fatal("invalid signature: ", err)
	}

	// Output:
	// public key: 1312 bytes
	// signature: 2420 bytes
}
