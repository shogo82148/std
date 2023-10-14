// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"github.com/shogo82148/std/crypto/rsa"
)

<<<<<<< HEAD
// ParsePKCS1PrivateKeyはPKCS #1形式、ASN.1 DER形式のRSA秘密鍵を解析します。
=======
// ParsePKCS1PrivateKey parses an [RSA] private key in PKCS #1, ASN.1 DER form.
>>>>>>> upstream/master
//
// この種のキーは一般的に、"RSA PRIVATE KEY"というタイプのPEMブロックにエンコードされます。
func ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error)

<<<<<<< HEAD
// MarshalPKCS1PrivateKeyはRSAの秘密鍵をPKCS #1、ASN.1 DER形式に変換します。
//
// この種類の鍵は、一般的には「RSA PRIVATE KEY」というタイプのPEMブロックにエンコードされます。
// RSA固有ではなく、より柔軟な鍵形式が必要な場合は、MarshalPKCS8PrivateKeyを使用してください。
func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte

// ParsePKCS1PublicKeyはPKCS＃1、ASN.1 DER形式のRSA公開鍵を解析します。
=======
// MarshalPKCS1PrivateKey converts an [RSA] private key to PKCS #1, ASN.1 DER form.
//
// This kind of key is commonly encoded in PEM blocks of type "RSA PRIVATE KEY".
// For a more flexible key format which is not [RSA] specific, use
// [MarshalPKCS8PrivateKey].
func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte

// ParsePKCS1PublicKey parses an [RSA] public key in PKCS #1, ASN.1 DER form.
>>>>>>> upstream/master
//
// この種の鍵は、一般的に "RSA PUBLIC KEY"というタイプのPEMブロックでエンコードされています。
func ParsePKCS1PublicKey(der []byte) (*rsa.PublicKey, error)

<<<<<<< HEAD
// MarshalPKCS1PublicKeyはRSA公開鍵をPKCS＃1のASN.1 DER形式に変換します。
=======
// MarshalPKCS1PublicKey converts an [RSA] public key to PKCS #1, ASN.1 DER form.
>>>>>>> upstream/master
//
// この種類の鍵は一般的に、"RSA PUBLIC KEY"タイプのPEMブロックにエンコードされます。
func MarshalPKCS1PublicKey(key *rsa.PublicKey) []byte
