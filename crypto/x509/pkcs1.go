// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"github.com/shogo82148/std/crypto/rsa"
)

// ParsePKCS1PrivateKeyはPKCS #1形式、ASN.1 DER形式の [RSA] 秘密鍵を解析します。
//
// この種の鍵は、一般的に"RSA PRIVATE KEY"タイプのPEMブロックでエンコードされます。
//
// Go 1.24以前では、CRTパラメータは無視され、再計算されていました。古い
// 動作を復元するには、GODEBUG=x509rsacrt=0環境変数を使用してください。
func ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error)

// MarshalPKCS1PrivateKeyは [RSA] の秘密鍵をPKCS #1、ASN.1 DER形式に変換します。
//
// この種の鍵は、一般的に"RSA PRIVATE KEY"タイプのPEMブロックでエンコードされます。
// [RSA] 固有でない、より柔軟な鍵形式については、
// [MarshalPKCS8PrivateKey] を使用してください。
//
// 鍵は最初に [rsa.PrivateKey.Validate] を呼び出して検証を通過している必要があります。
// MarshalPKCS1PrivateKeyは [rsa.PrivateKey.Precompute] を呼び出し、まだ事前計算されていない場合は
// 鍵を変更する可能性があります。
func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte

// ParsePKCS1PublicKeyはPKCS＃1、ASN.1 DER形式の [RSA] 公開鍵を解析します。
//
// この種の鍵は、一般的に "RSA PUBLIC KEY"というタイプのPEMブロックでエンコードされています。
func ParsePKCS1PublicKey(der []byte) (*rsa.PublicKey, error)

// MarshalPKCS1PublicKeyは [RSA] 公開鍵をPKCS＃1のASN.1 DER形式に変換します。
//
// この種類の鍵は一般的に、"RSA PUBLIC KEY"タイプのPEMブロックにエンコードされます。
func MarshalPKCS1PublicKey(key *rsa.PublicKey) []byte
