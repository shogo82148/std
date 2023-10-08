// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ecdsa_test

import (
	"github.com/shogo82148/std/crypto/ecdsa"
	"github.com/shogo82148/std/crypto/elliptic"
	"github.com/shogo82148/std/crypto/rand"
	"github.com/shogo82148/std/crypto/sha256"
	"github.com/shogo82148/std/fmt"
)

func Example() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	msg := "hello, world"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", sig)

	valid := ecdsa.VerifyASN1(&privateKey.PublicKey, hash[:], sig)
	fmt.Println("signature verified:", valid)
}
