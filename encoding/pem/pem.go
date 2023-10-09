// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージpemは、プライバシー拡張メールで起源を持つPEMデータのエンコーディングを実装しています。現在最も一般的なPEMエンコーディングの使用法は、TLSキーと証明書です。RFC 1421を参照してください。
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
// Headersはキー: 値の行からなる、可能性のある空のシーケンスです。
type Block struct {
	Type    string
	Headers map[string]string
	Bytes   []byte
}

// Decodeは入力内で次のPEM形式のブロック（証明書、秘密鍵など）を見つけます。それはそのブロックと入力の残り部分を返します。PEMデータが見つからない場合は、pがnilであり、入力全体がrestとして返されます。
func Decode(data []byte) (p *Block, rest []byte)

// Encodeは、bのPEMエンコーディングをoutに書き込みます。
func Encode(out io.Writer, b *Block) error

// EncodeToMemoryはbのPEMエンコーディングを返します。
//
// bに無効なヘッダーがある場合や、エンコードできない場合、
// EncodeToMemoryはnilを返します。このエラーケースの詳細を報告することが重要な場合は、
// 代わりにEncodeを使用してください。
func EncodeToMemory(b *Block) []byte
