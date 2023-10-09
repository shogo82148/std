// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"github.com/shogo82148/std/crypto/ecdsa"
)

// ParseECPrivateKeyはSEC 1、ASN.1 DER形式のECプライベートキーを解析します。
//
// この種類のキーは、一般的に "EC PRIVATE KEY" タイプのPEMブロックにエンコードされています。
func ParseECPrivateKey(der []byte) (*ecdsa.PrivateKey, error)

// MarshalECPrivateKeyは、ECの秘密鍵をSEC 1、ASN.1 DER形式に変換します。
//
// この種類の鍵は、一般的にはタイプ"EC PRIVATE KEY"のPEMブロックにエンコードされます。
// EC固有でないより柔軟な鍵形式を使用する場合は、MarshalPKCS8PrivateKeyを使用します。
func MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error)
