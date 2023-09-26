// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

import (
	"github.com/shogo82148/std/crypto"
	"github.com/shogo82148/std/io"
)

// EncryptPKCS1v15 encrypts the given message with RSA and the padding scheme from PKCS#1 v1.5.
// The message must be no longer than the length of the public modulus minus 11 bytes.
// WARNING: use of this function to encrypt plaintexts other than session keys
// is dangerous. Use RSA OAEP in new protocols.
func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) (out []byte, err error)

// DecryptPKCS1v15 decrypts a plaintext using RSA and the padding scheme from PKCS#1 v1.5.
// If rand != nil, it uses RSA blinding to avoid timing side-channel attacks.
func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) (out []byte, err error)

// DecryptPKCS1v15SessionKey decrypts a session key using RSA and the padding scheme from PKCS#1 v1.5.
// If rand != nil, it uses RSA blinding to avoid timing side-channel attacks.
// It returns an error if the ciphertext is the wrong length or if the
// ciphertext is greater than the public modulus. Otherwise, no error is
// returned. If the padding is valid, the resulting plaintext message is copied
// into key. Otherwise, key is unchanged. These alternatives occur in constant
// time. It is intended that the user of this function generate a random
// session key beforehand and continue the protocol with the resulting value.
// This will remove any possibility that an attacker can learn any information
// about the plaintext.
// See “Chosen Ciphertext Attacks Against Protocols Based on the RSA
// Encryption Standard PKCS #1”, Daniel Bleichenbacher, Advances in Cryptology
// (Crypto '98).
func DecryptPKCS1v15SessionKey(rand io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) (err error)

// These are ASN1 DER structures:
//   DigestInfo ::= SEQUENCE {
//     digestAlgorithm AlgorithmIdentifier,
//     digest OCTET STRING
//   }
// For performance, we don't use the generic ASN1 encoder. Rather, we
// precompute a prefix of the digest value that makes a valid ASN1 DER string
// with the correct contents.

// SignPKCS1v15 calculates the signature of hashed using RSASSA-PKCS1-V1_5-SIGN from RSA PKCS#1 v1.5.
// Note that hashed must be the result of hashing the input message using the
// given hash function.
func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) (s []byte, err error)

// VerifyPKCS1v15 verifies an RSA PKCS#1 v1.5 signature.
// hashed is the result of hashing the input message using the given hash
// function and sig is the signature. A valid signature is indicated by
// returning a nil error.
func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) (err error)
