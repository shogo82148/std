// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cipher

// AEADは関連データを含めた認証暗号化を提供する暗号モードです。手法の説明については、以下を参照してください。
// https://en.wikipedia.org/wiki/Authenticated_encryption.
type AEAD interface {
	NonceSize() int

	Overhead() int

	Seal(dst, nonce, plaintext, additionalData []byte) []byte

	Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error)
}

// NewGCMは、標準のnonce長でラップされた128ビットのブロック暗号を返します。
//
// 一般的に、GCMのこの実装で実行されるGHASH操作は一定時間ではありません。
// aes.NewCipherで生成された基礎のブロックが、AESのハードウェアサポートを持つシステムである場合は例外です。詳細については、crypto/aesパッケージのドキュメントを参照してください。
func NewGCM(cipher Block) (AEAD, error)

// NewGCMWithNonceSize は、与えられた長さの非スタンダードなノンスを受け付ける、128-bitのブロック暗号をGalios Counter Modeでラップしたものを返します。長さはゼロであってはいけません。
// 他の暗号システムとの互換性が必要な場合にのみ、この関数を使用してください。他のユーザーは、より高速でミス使用に対してより抵抗力のあるNewGCMを使用すべきです。
func NewGCMWithNonceSize(cipher Block, size int) (AEAD, error)

// NewGCMWithTagSizeは、指定された128ビットのブロック暗号をGalois Counter Modeでラップし、指定された長さのタグを生成します。
// 12バイトから16バイトのタグサイズが許可されています。
// 非標準のタグ長を使用する既存の暗号システムとの互換性が必要な場合にのみ、この関数を使用してください。その他のユーザーは、誤用に対してより耐性があるNewGCMを使用するべきです。
func NewGCMWithTagSize(cipher Block, tagSize int) (AEAD, error)
