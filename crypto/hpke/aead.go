// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hpke

import (
	"github.com/shogo82148/std/crypto/cipher"
)

// AEAD は HPKE 暗号スイートの 3 つのコンポーネントの 1 つで、対秱暗号化を
// 実装します。
type AEAD interface {
	ID() uint16
	keySize() int
	nonceSize() int
	aead(key []byte) (cipher.AEAD, error)
}

// NewAEAD は与えられた AEAD ID の AEAD 実装を返します。
//
// アプリケーションは、ランタイム可変性が必要でない限り、
// [AES128GCM] や [ChaCha20Poly1305] などの特定の実装を使用することをお勧めします。
func NewAEAD(id uint16) (AEAD, error)

// AES128GCM は AES-128-GCM AEAD 実装を返します。
func AES128GCM() AEAD

// AES256GCM は AES-256-GCM AEAD 実装を返します。
func AES256GCM() AEAD

// ChaCha20Poly1305 は ChaCha20Poly1305 AEAD 実装を返します。
func ChaCha20Poly1305() AEAD

// ExportOnly は、暗号化または複号化ができない。
// ただし [Sender.Export] または [Recipient.Export] でのみ秘密をエクスポートできる
// プレースホルダー AEAD 実装を返します。
//
// これが使用される場合、[Sender.Seal] と [Recipient.Open] はエラーを返します。
func ExportOnly() AEAD
