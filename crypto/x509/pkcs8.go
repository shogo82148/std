// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// ParsePKCS8PrivateKeyは、PKCS #8、ASN.1 DER形式の暗号化されていないプライベートキーを解析します。
//
// これは、 *[rsa.PrivateKey] 、 *[ecdsa.PrivateKey] 、 [ed25519.PrivateKey] (ポインタではなく)、または *[ecdh.PrivateKey] (X25519用)を返します。将来的にはさらに多くのタイプがサポートされる可能性があります。
//
// この種の鍵は、一般的に"PRIVATE KEY"タイプのPEMブロックでエンコードされます。
//
// Go 1.24以前では、RSA鍵のCRTパラメータは無視され、再計算されていました。
// 古い動作を復元するには、GODEBUG=x509rsacrt=0環境変数を使用してください。
func ParsePKCS8PrivateKey(der []byte) (key any, err error)

// MarshalPKCS8PrivateKeyは、プライベートキーをPKCS #8、ASN.1 DER形式に変換します。
//
// 現在、次のキータイプがサポートされています： *[rsa.PrivateKey] 、 *[ecdsa.PrivateKey] 、 [ed25519.PrivateKey] （ポインタでない）、および *[ecdh.PrivateKey] 。
// サポートされていないキータイプはエラーが発生します。
//
// この種の鍵は、一般的に"PRIVATE KEY"タイプのPEMブロックでエンコードされます。
//
// MarshalPKCS8PrivateKeyはRSA鍵に対して [rsa.PrivateKey.Precompute] を実行します。
func MarshalPKCS8PrivateKey(key any) ([]byte, error)
