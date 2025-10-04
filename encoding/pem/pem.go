// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// pemパッケージは、プライバシー拡張メールで起源を持つPEMデータのエンコーディングを実装しています。現在最も一般的なPEMエンコーディングの使用法は、TLSキーと証明書です。RFC 1421を参照してください。
package pem

import (
	"github.com/shogo82148/std/io"
)

// BlockはPEMエンコードされた構造体を表します。
//
// エンコードされた形式は次のようになります：
//
// -----BEGIN Type-----
// Headers
// Base64エンコードされたバイト
// -----END Type-----
//
// where [Block.Headers] is a possibly empty sequence of Key: Value lines.
type Block struct {
	Type    string
	Headers map[string]string
	Bytes   []byte
}

// Decodeは入力から次のPEM形式のブロック（証明書、秘密鍵
// など）を見つけます。そのブロックと入力の残りを返します。
// PEMデータが見つからない場合、pはnilで、入力全体がrestで返されます。
// ブロックは行の始まりから始まり、行の終わりで終わる必要があります。
func Decode(data []byte) (p *Block, rest []byte)

// Encodeは、bのPEMエンコーディングをoutに書き込みます。
func Encode(out io.Writer, b *Block) error

// EncodeToMemoryはbのPEMエンコーディングを返します。
//
// bが無効なヘッダーを持ちエンコードできない場合、
// EncodeToMemoryはnilを返します。このエラーケースの詳細を
// 報告することが重要な場合は、代わりに [Encode] を使用してください。
func EncodeToMemory(b *Block) []byte
