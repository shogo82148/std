// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"github.com/shogo82148/std/encoding/pem"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
)

type PEMCipher int

// EncryptPEMBlock暗号化アルゴリズムの可能な値。
const (
	_ PEMCipher = iota
	PEMCipherDES
	PEMCipher3DES
	PEMCipherAES128
	PEMCipherAES192
	PEMCipherAES256
)

// IsEncryptedPEMBlockは、PEMブロックがRFC 1423に従ってパスワードで暗号化されているかどうかを返します。
//
// Deprecated: RFC 1423で指定されている従来のPEM暗号化は、設計上の問題により安全ではありません。この方法では、暗号文を認証しないため、パディングオラクル攻撃に対して脆弱であり、攻撃者に平文を復元させる可能性があります。
func IsEncryptedPEMBlock(b *pem.Block) bool

// IncorrectPasswordError は、不正なパスワードが検出された場合に返されます。
var IncorrectPasswordError = errors.New("x509: decryption password incorrect")

// DecryptPEMBlockは、RFC 1423に従って暗号化されたPEMブロックと、それを暗号化するために使用されたパスワードを受け取り、復号化されたDER形式のバイトのスライスを返します。復号化に使用されるアルゴリズムは、DEK-Infoヘッダを調べて決定されます。DEK-Infoヘッダが存在しない場合、エラーが返されます。不正なパスワードが検出された場合、IncorrectPasswordErrorが返されます。フォーマットの不備のため、常に不正なパスワードを検出することはできません。これらの場合、エラーは返されませんが、復号化されたDERバイトはランダムなノイズになります。
// 廃止されました：RFC 1423で指定されたレガシーなPEM暗号化はセキュリティの設計上の問題があります。暗号文を認証しないため、パディングオラクル攻撃に対して脆弱であり、攻撃者が平文を回復することができます。
func DecryptPEMBlock(b *pem.Block, password []byte) ([]byte, error)

// EncryptPEMBlockは、指定されたアルゴリズムとパスワードで暗号化された指定されたDERエンコードされたデータを保持する指定されたタイプのPEMブロックを返します。
//
// 廃止予定: RFC 1423で指定されているレガシーPEM暗号化は、設計上安全ではありません。暗号文を認証しないため、パディングオラクル攻撃に対して脆弱であり、攻撃者が平文を復号化できる可能性があります。
func EncryptPEMBlock(rand io.Reader, blockType string, data, password []byte, alg PEMCipher) (*pem.Block, error)
