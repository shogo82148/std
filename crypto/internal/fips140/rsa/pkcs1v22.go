// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/io"
)

// PSSMaxSaltLength returns the maximum salt length for a given public key and
// hash function.
func PSSMaxSaltLength(pub *PublicKey, hash hash.Hash) (int, error)

// SignPSS calculates the signature of hashed using RSASSA-PSS.
func SignPSS(rand io.Reader, priv *PrivateKey, hash hash.Hash, hashed []byte, saltLength int) ([]byte, error)

// VerifyPSS verifies sig with RSASSA-PSS automatically detecting the salt length.
func VerifyPSS(pub *PublicKey, hash hash.Hash, digest []byte, sig []byte) error

// VerifyPSSWithSaltLength verifies sig with RSASSA-PSS and an expected salt length.
func VerifyPSSWithSaltLength(pub *PublicKey, hash hash.Hash, digest []byte, sig []byte, saltLength int) error

// EncryptOAEP encrypts the given message with RSAES-OAEP.
func EncryptOAEP(hash, mgfHash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)

// DecryptOAEP decrypts ciphertext using RSAES-OAEP.
func DecryptOAEP(hash, mgfHash hash.Hash, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
