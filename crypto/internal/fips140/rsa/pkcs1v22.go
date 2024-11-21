// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/crypto/internal/fips140"
	"github.com/shogo82148/std/io"
)

// PSSMaxSaltLength returns the maximum salt length for a given public key and
// hash function.
func PSSMaxSaltLength(pub *PublicKey, hash fips140.Hash) (int, error)

// SignPSS calculates the signature of hashed using RSASSA-PSS.
//
// In FIPS mode, rand is ignored and can be nil.
func SignPSS(rand io.Reader, priv *PrivateKey, hash fips140.Hash, hashed []byte, saltLength int) ([]byte, error)

// VerifyPSS verifies sig with RSASSA-PSS automatically detecting the salt length.
func VerifyPSS(pub *PublicKey, hash fips140.Hash, digest []byte, sig []byte) error

// VerifyPSS verifies sig with RSASSA-PSS and an expected salt length.
func VerifyPSSWithSaltLength(pub *PublicKey, hash fips140.Hash, digest []byte, sig []byte, saltLength int) error

// EncryptOAEP encrypts the given message with RSAES-OAEP.
//
// In FIPS mode, random is ignored and can be nil.
func EncryptOAEP(hash fips140.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)

// DecryptOAEP decrypts ciphertext using RSAES-OAEP.
func DecryptOAEP(hash, mgfHash fips140.Hash, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
