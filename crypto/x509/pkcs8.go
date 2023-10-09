// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// ParsePKCS8PrivateKeyは、PKCS #8、ASN.1 DER形式の暗号化されていないプライベートキーを解析します。
//
// これは、*rsa.PrivateKey、*ecdsa.PrivateKey、ed25519.PrivateKey (ポインタではなく)、または*ecdh.PrivateKey (X25519用)を返します。将来的にはさらに多くのタイプがサポートされる可能性があります。
//
// この種のキーは、一般的には「PRIVATE KEY」というタイプのPEMブロックにエンコードされています。
func ParsePKCS8PrivateKey(der []byte) (key any, err error)

// MarshalPKCS8PrivateKeyは、プライベートキーをPKCS #8、ASN.1 DER形式に変換します。
//
// 現在、次のキータイプがサポートされています：*rsa.PrivateKey、*ecdsa.PrivateKey、ed25519.PrivateKey（ポインタでない）、および*ecdh.PrivateKey。
// サポートされていないキータイプはエラーが発生します。
//
// この種のキーは一般的に、"PRIVATE KEY"というタイプのPEMブロックにエンコードされます。
func MarshalPKCS8PrivateKey(key any) ([]byte, error)
