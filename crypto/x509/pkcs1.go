// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"github.com/shogo82148/std/crypto/rsa"
)

// ParsePKCS1PrivateKeyはPKCS #1形式、ASN.1 DER形式の [RSA] 秘密鍵を解析します。
//
// この種のキーは一般的に、"RSA PRIVATE KEY"というタイプのPEMブロックにエンコードされます。
func ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error)

// MarshalPKCS1PrivateKeyは [RSA] の秘密鍵をPKCS #1、ASN.1 DER形式に変換します。
//
// この種類の鍵は、一般的には「RSA PRIVATE KEY」というタイプのPEMブロックにエンコードされます。
// [RSA] 固有ではなく、より柔軟な鍵形式が必要な場合は、 [MarshalPKCS8PrivateKey] を使用してください。
func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte

// ParsePKCS1PublicKeyはPKCS＃1、ASN.1 DER形式の [RSA] 公開鍵を解析します。
//
// この種の鍵は、一般的に "RSA PUBLIC KEY"というタイプのPEMブロックでエンコードされています。
func ParsePKCS1PublicKey(der []byte) (*rsa.PublicKey, error)

// MarshalPKCS1PublicKeyは [RSA] 公開鍵をPKCS＃1のASN.1 DER形式に変換します。
//
// この種類の鍵は一般的に、"RSA PUBLIC KEY"タイプのPEMブロックにエンコードされます。
func MarshalPKCS1PublicKey(key *rsa.PublicKey) []byte
