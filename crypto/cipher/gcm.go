// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cipher

<<<<<<< HEAD
// AEADは関連データを含めた認証暗号化を提供する暗号モードです。手法の説明については、以下を参照してください。
// https://en.wikipedia.org/wiki/Authenticated_encryption.
type AEAD interface {
	NonceSize() int

	Overhead() int

	Seal(dst, nonce, plaintext, additionalData []byte) []byte

	Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error)
}

// NewGCMは、標準のnonce長でラップされた128ビットのブロック暗号を返します。
=======
// NewGCM returns the given 128-bit, block cipher wrapped in Galois Counter Mode
// with the standard nonce length.
>>>>>>> upstream/release-branch.go1.25
//
// 一般的に、GCMのこの実装で実行されるGHASH操作は一定時間ではありません。
// aes.NewCipherで生成された基礎の [Block] が、AESのハードウェアサポートを持つシステムである場合は例外です。詳細については、 [crypto/aes] パッケージのドキュメントを参照してください。
func NewGCM(cipher Block) (AEAD, error)

// NewGCMWithNonceSize は、与えられた長さの非スタンダードなノンスを受け付ける、128-bitのブロック暗号をGalios Counter Modeでラップしたものを返します。長さはゼロであってはいけません。
// 他の暗号システムとの互換性が必要な場合にのみ、この関数を使用してください。他のユーザーは、より高速でミス使用に対してより抵抗力のある [NewGCM] を使用すべきです。
func NewGCMWithNonceSize(cipher Block, size int) (AEAD, error)

// NewGCMWithTagSizeは、指定された128ビットのブロック暗号をGalois Counter Modeでラップし、指定された長さのタグを生成します。
// 12バイトから16バイトのタグサイズが許可されています。
// 非標準のタグ長を使用する既存の暗号システムとの互換性が必要な場合にのみ、この関数を使用してください。その他のユーザーは、誤用に対してより耐性がある [NewGCM] を使用するべきです。
func NewGCMWithTagSize(cipher Block, tagSize int) (AEAD, error)

// NewGCMWithRandomNonce returns the given cipher wrapped in Galois Counter
// Mode, with randomly-generated nonces. The cipher must have been created by
// [crypto/aes.NewCipher].
//
// It generates a random 96-bit nonce, which is prepended to the ciphertext by Seal,
// and is extracted from the ciphertext by Open. The NonceSize of the AEAD is zero,
// while the Overhead is 28 bytes (the combination of nonce size and tag size).
//
// A given key MUST NOT be used to encrypt more than 2^32 messages, to limit the
// risk of a random nonce collision to negligible levels.
func NewGCMWithRandomNonce(cipher Block) (AEAD, error)
