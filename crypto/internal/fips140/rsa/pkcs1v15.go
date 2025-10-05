// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rsa

// SignPKCS1v15 calculates an RSASSA-PKCS1-v1.5 signature.
//
// hash is the name of the hash function as returned by [crypto.Hash.String]
// or the empty string to indicate that the message is signed directly.
func SignPKCS1v15(priv *PrivateKey, hash string, hashed []byte) ([]byte, error)

// VerifyPKCS1v15 verifies an RSASSA-PKCS1-v1.5 signature.
//
// hash is the name of the hash function as returned by [crypto.Hash.String]
// or the empty string to indicate that the message is signed directly.
func VerifyPKCS1v15(pub *PublicKey, hash string, hashed []byte, sig []byte) error
