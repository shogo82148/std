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

<<<<<<< HEAD
// Decodeは入力内で次のPEM形式のブロック（証明書、秘密鍵など）を見つけます。それはそのブロックと入力の残り部分を返します。PEMデータが見つからない場合は、pがnilであり、入力全体がrestとして返されます。
=======
// Decode will find the next PEM formatted block (certificate, private key
// etc) in the input. It returns that block and the remainder of the input. If
// no PEM data is found, p is nil and the whole of the input is returned in
// rest. Blocks must start at the beginning of a line and end at the end of a line.
>>>>>>> upstream/release-branch.go1.25
func Decode(data []byte) (p *Block, rest []byte)

// Encodeは、bのPEMエンコーディングをoutに書き込みます。
func Encode(out io.Writer, b *Block) error

// EncodeToMemoryはbのPEMエンコーディングを返します。
//
// If b has invalid headers and cannot be encoded,
// EncodeToMemory returns nil. If it is important to
// report details about this error case, use [Encode] instead.
func EncodeToMemory(b *Block) []byte
