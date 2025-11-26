// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hpke

import (
	"github.com/shogo82148/std/fmt"
)

func Example() {
	// In this example, we use MLKEM768-X25519 as the KEM, HKDF-SHA256 as the
	// KDF, and AES-256-GCM as the AEAD to encrypt a single message from a
	// sender to a recipient using the one-shot API.

	kem, kdf, aead := MLKEM768X25519(), HKDFSHA256(), AES256GCM()

	// Recipient side
	var (
		recipientPrivateKey PrivateKey
		publicKeyBytes      []byte
	)
	{
		k, err := kem.GenerateKey()
		if err != nil {
			panic(err)
		}
		recipientPrivateKey = k
		publicKeyBytes = k.PublicKey().Bytes()
	}

	// Sender side
	var ciphertext []byte
	{
		publicKey, err := kem.NewPublicKey(publicKeyBytes)
		if err != nil {
			panic(err)
		}

		message := []byte("|-()-|")
		ct, err := Seal(publicKey, kdf, aead, []byte("example"), message)
		if err != nil {
			panic(err)
		}

		ciphertext = ct
	}

	// Recipient side
	{
		plaintext, err := Open(recipientPrivateKey, kdf, aead, []byte("example"), ciphertext)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Decrypted message: %s\n", plaintext)
	}

	// Output:
	// Decrypted message: |-()-|
}
